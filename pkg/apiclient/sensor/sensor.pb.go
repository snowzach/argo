// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apiclient/sensor/sensor.proto

package sensor

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	v11 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListSensorsRequest struct {
	Namespace            string          `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ListOptions          *v1.ListOptions `protobuf:"bytes,2,opt,name=listOptions,proto3" json:"listOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListSensorsRequest) Reset()         { *m = ListSensorsRequest{} }
func (m *ListSensorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSensorsRequest) ProtoMessage()    {}
func (*ListSensorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ba963e1c6b5b55, []int{0}
}
func (m *ListSensorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListSensorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListSensorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListSensorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSensorsRequest.Merge(m, src)
}
func (m *ListSensorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListSensorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSensorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSensorsRequest proto.InternalMessageInfo

func (m *ListSensorsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListSensorsRequest) GetListOptions() *v1.ListOptions {
	if m != nil {
		return m.ListOptions
	}
	return nil
}

type SensorsLogsRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// optional
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TriggerName          string             `protobuf:"bytes,3,opt,name=triggerName,proto3" json:"triggerName,omitempty"`
	PodLogOptions        *v11.PodLogOptions `protobuf:"bytes,4,opt,name=podLogOptions,proto3" json:"podLogOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SensorsLogsRequest) Reset()         { *m = SensorsLogsRequest{} }
func (m *SensorsLogsRequest) String() string { return proto.CompactTextString(m) }
func (*SensorsLogsRequest) ProtoMessage()    {}
func (*SensorsLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ba963e1c6b5b55, []int{1}
}
func (m *SensorsLogsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SensorsLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SensorsLogsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SensorsLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorsLogsRequest.Merge(m, src)
}
func (m *SensorsLogsRequest) XXX_Size() int {
	return m.Size()
}
func (m *SensorsLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorsLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SensorsLogsRequest proto.InternalMessageInfo

func (m *SensorsLogsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SensorsLogsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SensorsLogsRequest) GetTriggerName() string {
	if m != nil {
		return m.TriggerName
	}
	return ""
}

func (m *SensorsLogsRequest) GetPodLogOptions() *v11.PodLogOptions {
	if m != nil {
		return m.PodLogOptions
	}
	return nil
}

type LogEntry struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	SensorName           string   `protobuf:"bytes,2,opt,name=sensorName,proto3" json:"sensorName,omitempty"`
	TriggerName          string   `protobuf:"bytes,3,opt,name=triggerName,proto3" json:"triggerName,omitempty"`
	Level                string   `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	Time                 *v1.Time `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	Msg                  string   `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ba963e1c6b5b55, []int{2}
}
func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return m.Size()
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LogEntry) GetSensorName() string {
	if m != nil {
		return m.SensorName
	}
	return ""
}

func (m *LogEntry) GetTriggerName() string {
	if m != nil {
		return m.TriggerName
	}
	return ""
}

func (m *LogEntry) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LogEntry) GetTime() *v1.Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *LogEntry) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ListSensorsRequest)(nil), "sensor.ListSensorsRequest")
	proto.RegisterType((*SensorsLogsRequest)(nil), "sensor.SensorsLogsRequest")
	proto.RegisterType((*LogEntry)(nil), "sensor.LogEntry")
}

func init() { proto.RegisterFile("pkg/apiclient/sensor/sensor.proto", fileDescriptor_78ba963e1c6b5b55) }

var fileDescriptor_78ba963e1c6b5b55 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x8b, 0x13, 0x3f,
	0x18, 0x26, 0xdd, 0x6e, 0x7f, 0xbf, 0xa6, 0x2c, 0x2c, 0xc1, 0x43, 0xa9, 0x6b, 0xe9, 0x8e, 0x07,
	0x57, 0x61, 0x13, 0xbb, 0x7a, 0x10, 0x04, 0x0f, 0xa2, 0xe8, 0xa1, 0xac, 0x32, 0xf5, 0xe4, 0x45,
	0xb2, 0xb3, 0x2f, 0x69, 0xec, 0x4c, 0x32, 0x26, 0xd9, 0x81, 0x45, 0xbc, 0x78, 0xf1, 0x03, 0x28,
	0x7e, 0x0f, 0xbf, 0x85, 0x47, 0xd1, 0x2f, 0x20, 0xc5, 0x0f, 0x22, 0xc9, 0x4c, 0x3b, 0x23, 0xad,
	0xd8, 0x53, 0xdf, 0xbc, 0x7f, 0x9f, 0x67, 0xde, 0xe7, 0x2d, 0x3e, 0xcc, 0xe7, 0x82, 0xf1, 0x5c,
	0x26, 0xa9, 0x04, 0xe5, 0x98, 0x05, 0x65, 0xb5, 0xa9, 0x7e, 0x68, 0x6e, 0xb4, 0xd3, 0xa4, 0x53,
	0xbe, 0x06, 0x07, 0x42, 0x6b, 0x91, 0x82, 0xcf, 0x66, 0x5c, 0x29, 0xed, 0xb8, 0x93, 0x5a, 0xd9,
	0x32, 0x6b, 0x70, 0x77, 0x7e, 0xcf, 0x52, 0xa9, 0x7d, 0x34, 0xe3, 0xc9, 0x4c, 0x2a, 0x30, 0x97,
	0xac, 0x6a, 0x6e, 0x59, 0x06, 0x8e, 0xb3, 0x62, 0xcc, 0x04, 0x28, 0x30, 0xdc, 0xc1, 0x79, 0x55,
	0x15, 0xd5, 0x55, 0x2c, 0xd1, 0x06, 0x36, 0xe5, 0x3c, 0x15, 0xd2, 0xcd, 0x2e, 0xce, 0x68, 0xa2,
	0x33, 0xc6, 0x8d, 0xd0, 0xb9, 0xd1, 0xaf, 0x83, 0x71, 0x0c, 0x05, 0x28, 0x67, 0xeb, 0x29, 0x15,
	0xfa, 0x62, 0xcc, 0xd3, 0x7c, 0xc6, 0xd7, 0x3a, 0x45, 0x1f, 0x10, 0x26, 0x13, 0x69, 0xdd, 0x34,
	0xe4, 0xd9, 0x18, 0xde, 0x5c, 0x80, 0x75, 0xe4, 0x00, 0x77, 0x15, 0xcf, 0xc0, 0xe6, 0x3c, 0x81,
	0x3e, 0x1a, 0xa1, 0xa3, 0x6e, 0x5c, 0x3b, 0xc8, 0x14, 0xf7, 0x52, 0x69, 0xdd, 0xb3, 0x3c, 0xb0,
	0xed, 0xb7, 0x46, 0xe8, 0xa8, 0x77, 0x32, 0xa6, 0x25, 0x70, 0xda, 0xa4, 0x4b, 0xf3, 0xb9, 0xf0,
	0x0e, 0x4b, 0x3d, 0x5d, 0x5a, 0x8c, 0xe9, 0xa4, 0x2e, 0x8c, 0x9b, 0x5d, 0xa2, 0x2f, 0x08, 0x93,
	0x0a, 0xc5, 0x44, 0x8b, 0x2d, 0x91, 0x10, 0xdc, 0xf6, 0x8f, 0x00, 0xa1, 0x1b, 0x07, 0x9b, 0x8c,
	0x70, 0xcf, 0x19, 0x29, 0x04, 0x98, 0x53, 0x1f, 0xda, 0x09, 0xa1, 0xa6, 0x8b, 0x3c, 0xc1, 0x7b,
	0xb9, 0x3e, 0x9f, 0x68, 0xb1, 0x64, 0xd0, 0x0e, 0x0c, 0x0e, 0x1b, 0x0c, 0xa8, 0xff, 0xf4, 0x1e,
	0xef, 0xf3, 0x66, 0x62, 0xfc, 0x67, 0x5d, 0xf4, 0x1d, 0xe1, 0xff, 0x27, 0x5a, 0x3c, 0x56, 0xce,
	0x5c, 0xfe, 0x03, 0xe9, 0x10, 0xe3, 0x72, 0x17, 0xa7, 0x35, 0xde, 0x86, 0x67, 0x0b, 0xd4, 0x57,
	0xf0, 0x6e, 0x0a, 0x05, 0xa4, 0xfd, 0xdd, 0x10, 0x2b, 0x1f, 0xe4, 0x01, 0x6e, 0x3b, 0x99, 0x41,
	0xbf, 0x13, 0x28, 0xdc, 0xda, 0x6e, 0x09, 0x2f, 0x64, 0x06, 0x71, 0xa8, 0x23, 0xfb, 0x78, 0x27,
	0xb3, 0xa2, 0xff, 0x5f, 0xe8, 0xe9, 0xcd, 0x93, 0xcf, 0x2d, 0xbc, 0x57, 0x2e, 0x62, 0x0a, 0xa6,
	0x90, 0x09, 0x90, 0x4f, 0x08, 0xf7, 0x1a, 0x22, 0x21, 0x03, 0x5a, 0x5d, 0xc3, 0xba, 0x72, 0x06,
	0x8f, 0x68, 0xad, 0x4d, 0xba, 0xd4, 0x66, 0x30, 0x5e, 0x95, 0xda, 0xac, 0xd1, 0x54, 0x4d, 0x96,
	0xda, 0xa4, 0x65, 0x27, 0xdf, 0x33, 0xba, 0xfe, 0xfe, 0xc7, 0xaf, 0x8f, 0xad, 0x6b, 0xe4, 0x6a,
	0x38, 0x83, 0x62, 0x5c, 0xe9, 0xd8, 0xb2, 0xb7, 0xab, 0x2f, 0xfa, 0x8e, 0x28, 0xdc, 0x6b, 0x08,
	0xa6, 0x46, 0xb5, 0xae, 0xa2, 0xc1, 0xfe, 0x0a, 0x71, 0xb5, 0xad, 0x88, 0x85, 0x09, 0x37, 0xc9,
	0x8d, 0xd5, 0x04, 0x67, 0x80, 0x67, 0x9b, 0x06, 0xb1, 0x54, 0x0b, 0x7b, 0x1b, 0x3d, 0xbc, 0xff,
	0x75, 0x31, 0x44, 0xdf, 0x16, 0x43, 0xf4, 0x73, 0x31, 0x44, 0x2f, 0x8f, 0xff, 0x76, 0x83, 0x6c,
	0xd3, 0xff, 0xc7, 0x59, 0x27, 0xdc, 0xdb, 0x9d, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0xd9,
	0xf0, 0x67, 0x5e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SensorServiceClient is the client API for SensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SensorServiceClient interface {
	ListSensors(ctx context.Context, in *ListSensorsRequest, opts ...grpc.CallOption) (*v1alpha1.SensorList, error)
	SensorsLogs(ctx context.Context, in *SensorsLogsRequest, opts ...grpc.CallOption) (SensorService_SensorsLogsClient, error)
}

type sensorServiceClient struct {
	cc *grpc.ClientConn
}

func NewSensorServiceClient(cc *grpc.ClientConn) SensorServiceClient {
	return &sensorServiceClient{cc}
}

func (c *sensorServiceClient) ListSensors(ctx context.Context, in *ListSensorsRequest, opts ...grpc.CallOption) (*v1alpha1.SensorList, error) {
	out := new(v1alpha1.SensorList)
	err := c.cc.Invoke(ctx, "/sensor.SensorService/ListSensors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) SensorsLogs(ctx context.Context, in *SensorsLogsRequest, opts ...grpc.CallOption) (SensorService_SensorsLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SensorService_serviceDesc.Streams[0], "/sensor.SensorService/SensorsLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorServiceSensorsLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SensorService_SensorsLogsClient interface {
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type sensorServiceSensorsLogsClient struct {
	grpc.ClientStream
}

func (x *sensorServiceSensorsLogsClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SensorServiceServer is the server API for SensorService service.
type SensorServiceServer interface {
	ListSensors(context.Context, *ListSensorsRequest) (*v1alpha1.SensorList, error)
	SensorsLogs(*SensorsLogsRequest, SensorService_SensorsLogsServer) error
}

// UnimplementedSensorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSensorServiceServer struct {
}

func (*UnimplementedSensorServiceServer) ListSensors(ctx context.Context, req *ListSensorsRequest) (*v1alpha1.SensorList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSensors not implemented")
}
func (*UnimplementedSensorServiceServer) SensorsLogs(req *SensorsLogsRequest, srv SensorService_SensorsLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method SensorsLogs not implemented")
}

func RegisterSensorServiceServer(s *grpc.Server, srv SensorServiceServer) {
	s.RegisterService(&_SensorService_serviceDesc, srv)
}

func _SensorService_ListSensors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSensorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).ListSensors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.SensorService/ListSensors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).ListSensors(ctx, req.(*ListSensorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_SensorsLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SensorsLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorServiceServer).SensorsLogs(m, &sensorServiceSensorsLogsServer{stream})
}

type SensorService_SensorsLogsServer interface {
	Send(*LogEntry) error
	grpc.ServerStream
}

type sensorServiceSensorsLogsServer struct {
	grpc.ServerStream
}

func (x *sensorServiceSensorsLogsServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

var _SensorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.SensorService",
	HandlerType: (*SensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSensors",
			Handler:    _SensorService_ListSensors_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SensorsLogs",
			Handler:       _SensorService_SensorsLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/apiclient/sensor/sensor.proto",
}

func (m *ListSensorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListSensorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListSensorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ListOptions != nil {
		{
			size, err := m.ListOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSensor(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SensorsLogsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SensorsLogsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SensorsLogsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PodLogOptions != nil {
		{
			size, err := m.PodLogOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSensor(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.TriggerName) > 0 {
		i -= len(m.TriggerName)
		copy(dAtA[i:], m.TriggerName)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.TriggerName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Time != nil {
		{
			size, err := m.Time.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSensor(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Level) > 0 {
		i -= len(m.Level)
		copy(dAtA[i:], m.Level)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.Level)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TriggerName) > 0 {
		i -= len(m.TriggerName)
		copy(dAtA[i:], m.TriggerName)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.TriggerName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SensorName) > 0 {
		i -= len(m.SensorName)
		copy(dAtA[i:], m.SensorName)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.SensorName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintSensor(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSensor(dAtA []byte, offset int, v uint64) int {
	offset -= sovSensor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListSensorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	if m.ListOptions != nil {
		l = m.ListOptions.Size()
		n += 1 + l + sovSensor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SensorsLogsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	l = len(m.TriggerName)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	if m.PodLogOptions != nil {
		l = m.PodLogOptions.Size()
		n += 1 + l + sovSensor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	l = len(m.SensorName)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	l = len(m.TriggerName)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	if m.Time != nil {
		l = m.Time.Size()
		n += 1 + l + sovSensor(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovSensor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSensor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSensor(x uint64) (n int) {
	return sovSensor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListSensorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSensor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListSensorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListSensorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListOptions == nil {
				m.ListOptions = &v1.ListOptions{}
			}
			if err := m.ListOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSensor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSensor
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSensor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SensorsLogsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSensor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SensorsLogsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SensorsLogsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TriggerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodLogOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PodLogOptions == nil {
				m.PodLogOptions = &v11.PodLogOptions{}
			}
			if err := m.PodLogOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSensor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSensor
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSensor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSensor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TriggerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = &v1.Time{}
			}
			if err := m.Time.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSensor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSensor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSensor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSensor
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSensor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSensor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSensor
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSensor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSensor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSensor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSensor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSensor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSensor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSensor = fmt.Errorf("proto: unexpected end of group")
)
