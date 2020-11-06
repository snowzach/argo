import {Page, SlidingPanel, Tabs} from 'argo-ui/src/index';
import * as React from 'react';
import {RouteComponentProps} from 'react-router';
import {Condition, EventSource, eventTypes, kubernetes, Sensor, triggerTypes} from '../../../../models';
import {uiUrl} from '../../../shared/base';
import {BasePage} from '../../../shared/components/base-page';
import {ErrorNotice} from '../../../shared/components/error-notice';
import {NamespaceFilter} from '../../../shared/components/namespace-filter';
import {ResourceEditor} from '../../../shared/components/resource-editor/resource-editor';
import {ZeroState} from '../../../shared/components/zero-state';
import {services} from '../../../shared/services';
import {Utils} from '../../../shared/utils';
import {EventsPanel} from '../../../workflows/components/events-panel';
import {FullHeightLogsViewer} from '../../../workflows/components/workflow-logs-viewer/full-height-logs-viewer';
import {Edge, Graph, GraphPanel, Node} from './graph-panel';
import {icons} from './icons';
import {ID} from './id';

require('../../../workflows/components/workflow-details/workflow-details.scss');

interface State {
    namespace: string;
    selectedId?: string;
    error?: Error;
    resources: {[id: string]: {metadata: kubernetes.ObjectMeta; status?: {conditions?: Condition[]}}};
    touched: {[id: string]: any};
}

const phase = (r: {status?: {conditions?: Condition[]}}) => {
    if (!r.status || !r.status.conditions) {
        return '';
    }
    return r.status.conditions.find(c => c.status !== 'True') ? 'Warning' : 'Running';
};

export class NamespaceDetails extends BasePage<RouteComponentProps<any>, State> {
    private get namespace() {
        return this.state.namespace;
    }

    private set namespace(namespace: string) {
        this.fetch(namespace);
    }

    private get graph(): Graph {
        const nodes: Node[] = Object.entries(this.state.resources)
            .filter(([id]) => ID.split(id).type === 'Sensor')
            .map(([sensorId, sensor]) => ({
                id: sensorId,
                type: 'sensor',
                label: sensor.metadata.name,
                icon: icons.Sensor,
                phase: phase(sensor),
                touched: !!this.state.touched[sensorId]
            }));

        const edges: Edge[] = [];

        Object.entries(this.state.resources)
            .filter(([eventSourceId]) => ID.split(eventSourceId).type === 'EventSource')
            .forEach(([, eventSource]) => {
                const spec = (eventSource as EventSource).spec;
                Object.entries(spec)
                    .filter(([typeKey]) => ['template', 'service'].indexOf(typeKey) < 0)
                    .forEach(([typeKey, type]) => {
                        Object.keys(type).forEach(key => {
                            const eventId = ID.join({
                                type: 'EventType',
                                namespace: eventSource.metadata.namespace,
                                name: eventSource.metadata.name,
                                key
                            });
                            nodes.push({
                                id: eventId,
                                type: typeKey,
                                label: key,
                                phase: phase(eventSource),
                                icon: icons[eventTypes[typeKey] + 'EventType'],
                                touched: !!this.state.touched[eventId]
                            });
                        });
                    });
            });
        Object.entries(this.state.resources)
            .filter(([sensorId]) => ID.split(sensorId).type === 'Sensor')
            .forEach(([sensorId, sensor]) => {
                const spec = (sensor as Sensor).spec;
                spec.dependencies.forEach(d => {
                    const eventId = ID.join({
                        type: 'EventType',
                        namespace: sensor.metadata.namespace,
                        name: d.eventSourceName,
                        key: d.eventName
                    });
                    edges.push({x: eventId, y: sensorId, label: d.name});
                });
                spec.triggers
                    .map(t => t.template)
                    .filter(template => template)
                    .forEach(template => {
                        const triggerTypeKey = Object.keys(template).filter(t => ['name', 'conditions'].indexOf(t) === -1)[0];
                        const triggerId = ID.join({
                            type: 'Trigger',
                            namespace: sensor.metadata.namespace,
                            name: sensor.metadata.name,
                            key: template.name
                        });
                        nodes.push({
                            id: triggerId,
                            label: template.name,
                            type: triggerTypeKey,
                            phase: phase(sensor),
                            icon: icons[triggerTypes[triggerTypeKey] + 'Trigger'],
                            touched: !!this.state.touched[triggerId]
                        });
                        if (template.conditions) {
                            const conditionsId = ID.join({
                                type: 'Conditions',
                                namespace: sensor.metadata.namespace,
                                name: sensor.metadata.name,
                                key: template.conditions
                            });
                            nodes.push({
                                id: conditionsId,
                                label: template.conditions,
                                type: 'conditions',
                                icon: icons.Conditions,
                                touched: !!this.state.touched[conditionsId]
                            });
                            edges.push({x: sensorId, y: conditionsId});
                            edges.push({x: conditionsId, y: triggerId});
                        } else {
                            edges.push({x: sensorId, y: triggerId});
                        }
                    });
            });
        return {nodes, edges};
    }

    private get selected() {
        return this.resource(this.state.selectedId);
    }

    constructor(props: RouteComponentProps<any>, context: any) {
        super(props, context);
        this.state = {namespace: this.props.match.params.namespace || '', resources: {}, touched: {}};
    }

    public render() {
        const selected = this.selected;
        return (
            <Page
                title='Namespace'
                toolbar={{
                    breadcrumbs: [{title: 'Namespaces', path: uiUrl('namespaces')}],
                    tools: [<NamespaceFilter key='namespace-filter' value={this.namespace} onChange={namespace => (this.namespace = namespace)} />]
                }}>
                <div className='argo-container'>{this.renderGraph()}</div>
                <SlidingPanel isShown={!!selected} onClose={() => this.setState({selectedId: null})}>
                    {!!selected && (
                        <div>
                            <h4>
                                {selected.kind}/{selected.name} {selected.key}
                            </h4>
                            <Tabs
                                navTransparent={true}
                                tabs={[
                                    {
                                        title: 'SUMMARY',
                                        key: 'summary',
                                        content: <ResourceEditor readonly={true} kind={selected.kind} value={selected.value} />
                                    },
                                    {
                                        title: 'LOGS',
                                        key: 'logs',
                                        content: (
                                            <div className='white-box' style={{height: 600}}>
                                                <FullHeightLogsViewer
                                                    source={{
                                                        key: 'logs',
                                                        loadLogs: () =>
                                                            selected.kind === 'Sensor'
                                                                ? services.sensor
                                                                      .sensorsLogs(this.namespace, selected.name, selected.key, 10)
                                                                      .map(e => e.time + ' ' + e.level + ': ' + e.msg + '\n')
                                                                : services.eventSource
                                                                      .eventSourcesLogs(this.namespace, selected.name, '', selected.key, 10)
                                                                      .map(e => e.time + ' ' + e.level + ': ' + e.msg + '\n'),
                                                        shouldRepeat: () => false
                                                    }}
                                                />
                                            </div>
                                        )
                                    },
                                    {
                                        title: 'EVENTS',
                                        key: 'events',
                                        content: <EventsPanel kind={selected.kind} namespace={selected.namespace} name={selected.name} />
                                    }
                                ]}
                            />
                        </div>
                    )}
                </SlidingPanel>
            </Page>
        );
    }

    public componentDidMount(): void {
        this.fetch(this.namespace);
    }

    private resource(i: string) {
        if (!i) {
            return;
        }
        const {type, namespace, name, key} = ID.split(i);
        const kind = ({EventType: 'EventSource', Trigger: 'Sensor'} as {[key: string]: string})[type] || type;
        return {namespace, kind, name, key, value: this.state.resources[ID.join({type: kind, namespace, name})]};
    }

    private renderGraph() {
        if (this.state.error) {
            return <ErrorNotice error={this.state.error} onReload={() => this.fetch(this.namespace)} />;
        }
        const g = this.graph;
        if (g.nodes.length === 0) {
            return <ZeroState title='No entities found'>No Argo Events resources found.</ZeroState>;
        }
        return (
            <div style={{textAlign: 'center'}}>
                <GraphPanel graph={g} onSelect={selectedId => this.setState({selectedId})} />
            </div>
        );
    }

    private fetch(namespace: string) {
        const updateResources = (s: State, type: string, list: {items: {metadata: kubernetes.ObjectMeta}[]}) => {
            (list.items || []).forEach(v => {
                s.resources[ID.join({type, namespace: v.metadata.namespace, name: v.metadata.name})] = v;
            });
            return {resources: s.resources};
        };
        this.setState({resources: {}}, () => {
            Promise.all([
                services.eventSource.list(namespace).then(list => this.setState(s => updateResources(s, 'EventSource', list))),
                services.sensor.list(namespace).then(list => this.setState(s => updateResources(s, 'Sensor', list)))
            ])
                .then(() =>
                    this.setState({error: null, namespace}, () => {
                        this.appContext.router.history.push(uiUrl(`namespaces/${namespace}`));
                        Utils.setCurrentNamespace(namespace);
                    })
                )
                .then(() =>
                    Promise.all([
                        services.eventSource
                            .eventSourcesLogs(namespace, '', '', '', 0)
                            .filter(e => !!e.eventSourceName)
                            .subscribe(
                                e =>
                                    this.touch(
                                        ID.join({
                                            type: 'EventType',
                                            namespace: e.namespace,
                                            name: e.eventSourceName,
                                            key: e.eventSourceType + '.' + e.eventName
                                        })
                                    ),
                                error => this.setState({error})
                            ),
                        services.sensor
                            .sensorsLogs(namespace, '', '', 0)
                            .filter(e => !!e.triggerName)
                            .subscribe(
                                e => {
                                    this.touch(
                                        ID.join({
                                            type: 'Sensor',
                                            namespace: e.namespace,
                                            name: e.sensorName
                                        })
                                    );
                                    this.touch(
                                        ID.join({
                                            type: 'Trigger',
                                            namespace: e.namespace,
                                            name: e.sensorName,
                                            key: e.triggerName
                                        })
                                    );
                                },
                                error => this.setState({error})
                            )
                    ])
                )
                .catch(error => this.setState({error}));
        });
    }

    private touch(id: string) {
        this.setState(state => {
            clearTimeout(state.touched[id]);
            state.touched[id] = setTimeout(() => {
                this.setState(s => {
                    delete s.touched[id];
                    return {touched: s.touched};
                });
            }, 3000);
            return {touched: state.touched};
        });
    }
}
