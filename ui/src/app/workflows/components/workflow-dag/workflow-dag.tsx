import * as React from 'react';

import {NODE_PHASE, NodeStatus} from '../../../../models';
import {GraphPanel} from '../../../shared/components/graph/graph-panel';
import {Graph} from '../../../shared/components/graph/types';
import {Loading} from '../../../shared/components/loading';
import {Utils} from '../../../shared/utils';
import {getCollapsedNodeName, getMessage, getNodeParent, isCollapsedNode} from './graph/collapsible-node';
import {icons} from './icons';
import {WorkflowDagRenderOptionsPanel} from './workflow-dag-render-options-panel';

export interface WorkflowDagRenderOptions {
    expandNodes: Set<string>;
}

export interface WorkflowDagProps {
    workflowName: string;
    nodes: {[nodeId: string]: NodeStatus};
    selectedNodeId?: string;
    nodeClicked?: (nodeId: string) => any;
}

const types = {
    Pod: true,
    Steps: true,
    DAG: true,
    Retry: true,
    Skipped: true,
    Suspend: true,
    TaskGroup: false,
    StepGroup: false,
    Collapsed: true
};

function progress(n: NodeStatus) {
    if (!n || !n.estimatedDuration) {
        return null;
    }
    return (new Date().getTime() - new Date(n.startedAt).getTime()) / 1000 / n.estimatedDuration;
}

function nodeLabel(n: NodeStatus) {
    const phase = n.type === 'Suspend' && n.phase === 'Running' ? 'Suspended' : n.phase;
    return {
        label: Utils.shortNodeName(n),
        type: n.type,
        icon: icons[phase],
        progress: phase === 'Running' && progress(n),
        classNames: phase
    };
}

const classNames = () => {
    const v: {[label: string]: boolean} = {
        Suspended: true
    };
    Object.entries(NODE_PHASE).forEach(([, label]) => (v[label] = true));
    return v;
};

export class WorkflowDag extends React.Component<WorkflowDagProps, WorkflowDagRenderOptions> {
    private graph: Graph;

    constructor(props: Readonly<WorkflowDagProps>) {
        super(props);
        this.state = {
            expandNodes: new Set()
        };
    }

    public render() {
        if (!this.props.nodes) {
            return <Loading />;
        }
        this.prepareGraph();
        return (
            <GraphPanel
                graph={this.graph}
                types={types}
                classNames={classNames()}
                selectedNode={this.props.selectedNodeId}
                onNodeSelect={id => this.selectNode(id)}
                nodeSize={40}
                options={<WorkflowDagRenderOptionsPanel {...this.state} onChange={workflowDagRenderOptions => this.saveOptions(workflowDagRenderOptions)} />}
                hideTypes={true}
            />
        );
    }

    private saveOptions(newChanges: WorkflowDagRenderOptions) {
        this.setState(newChanges);
    }

    private getNode(nodeId: string): NodeStatus {
        const node: NodeStatus = this.props.nodes[nodeId];
        if (!node) {
            return null;
        }
        return node;
    }

    private prepareGraph() {
        this.graph = new Graph();
        const edges = this.graph.edges;
        const nodes = this.graph.nodes;

        interface PrepareNode {
            nodeName: string;
            children: string[];
            parent: string;
        }

        const allNodes = this.props.nodes;
        const getChildren = (nodeId: string): string[] => {
            if (!allNodes[nodeId] || !allNodes[nodeId].children) {
                return [];
            }
            return allNodes[nodeId].children.filter(child => allNodes[child]);
        };
        const pushChildren = (nodeId: string, isExpanded: boolean, queue: PrepareNode[]): void => {
            const children: string[] = getChildren(nodeId);
            if (!children) {
                return;
            }

            if (children.length > 3 && !isExpanded) {
                // Node will be collapsed
                queue.push({
                    nodeName: children[0],
                    parent: nodeId,
                    children: getChildren(children[0])
                });
                const newChildren: string[] = children
                    .slice(1, children.length - 1)
                    .map(v => [v])
                    .reduce((a, b) => a.concat(b), []);
                queue.push({
                    nodeName: getCollapsedNodeName(nodeId, children.length - 2 + ' hidden nodes', allNodes[children[0]].type),
                    parent: nodeId,
                    children: newChildren
                });
                queue.push({
                    nodeName: children[children.length - 1],
                    parent: nodeId,
                    children: getChildren(children[children.length - 1])
                });
            } else {
                // Node will not be collapsed
                children.map(child =>
                    queue.push({
                        nodeName: child,
                        parent: nodeId,
                        children: getChildren(child)
                    })
                );
            }
        };

        const traverse = (root: PrepareNode): void => {
            const queue: PrepareNode[] = [root];
            const consideredChildren: Set<string> = new Set<string>();
            let previousCollapsed: string = '';

            while (queue.length > 0) {
                const item = queue.pop();

                if (isCollapsedNode(item.nodeName)) {
                    if (item.nodeName !== previousCollapsed) {
                        nodes.set(item.nodeName, {
                            label: getMessage(item.nodeName),
                            type: 'Collapsed',
                            icon: icons.Collapsed
                        });
                        edges.set({v: item.parent, w: item.nodeName}, {});
                        previousCollapsed = item.nodeName;
                    }
                    continue;
                }
                const child = allNodes[item.nodeName];
                const isExpanded: boolean = this.state.expandNodes.has('*') || this.state.expandNodes.has(item.nodeName);
                nodes.set(item.nodeName, nodeLabel(child));
                edges.set({v: item.parent, w: item.nodeName}, {});

                // If we have already considered the children of this node, don't consider them again
                if (consideredChildren.has(item.nodeName)) {
                    continue;
                }
                consideredChildren.add(item.nodeName);

                const node: NodeStatus = this.props.nodes[item.nodeName];
                if (!node || node.phase === NODE_PHASE.OMITTED) {
                    continue;
                }

                pushChildren(node.id, isExpanded, queue);
            }
        };

        const workflowRoot: PrepareNode = {
            nodeName: this.props.workflowName,
            parent: '',
            children: getChildren(this.props.workflowName)
        };

        // Traverse the workflow from the root node
        traverse(workflowRoot);

        const onExitHandlerNodeId = Object.values(allNodes).find(nodeId => nodeId.name === `${this.props.workflowName}.onExit`);
        if (onExitHandlerNodeId) {
            this.getOutboundNodes(this.props.workflowName).forEach(v => {
                const exitHandler = allNodes[onExitHandlerNodeId.id];
                nodes.set(onExitHandlerNodeId.id, nodeLabel(exitHandler));
                if (nodes.has(v)) {
                    edges.set({v, w: onExitHandlerNodeId.id}, {});
                }
            });
            const onExitRoot: PrepareNode = {
                nodeName: onExitHandlerNodeId.id,
                parent: '',
                children: getChildren(onExitHandlerNodeId.id)
            };
            // Traverse the onExit tree starting from the onExit node itself
            traverse(onExitRoot);
        }
    }

    private selectNode(nodeId: string) {
        if (isCollapsedNode(nodeId)) {
            this.expandNode(nodeId);
        } else {
            return this.props.nodeClicked && this.props.nodeClicked(nodeId);
        }
    }

    private expandNode(nodeId: string) {
        if (isCollapsedNode(getNodeParent(nodeId))) {
            this.expandNode(getNodeParent(nodeId));
        } else {
            this.setState({expandNodes: new Set(this.state.expandNodes).add(getNodeParent(nodeId))});
        }
    }

    private getOutboundNodes(nodeID: string): string[] {
        const node = this.getNode(nodeID);
        if (node.type === 'Pod' || node.type === 'Skipped') {
            return [node.id];
        }
        let outbound = Array<string>();
        for (const outboundNodeID of node.outboundNodes || []) {
            const outNode = this.getNode(outboundNodeID);
            if (outNode.type === 'Pod') {
                outbound.push(outboundNodeID);
            } else {
                outbound = outbound.concat(this.getOutboundNodes(outboundNodeID));
            }
        }
        return outbound;
    }
}
