// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_discovery.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	context "golang.org/x/net/context"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Operation defines the set of possible operations that can be performed in
// relation to a resource.
type StreamResponse_Operation int32

const (
	// Specifies an unknown operation.
	StreamResponse_UNKNOWN StreamResponse_Operation = 0
	// Specifies a create operation.
	StreamResponse_CREATE StreamResponse_Operation = 1
	// Specifies an update operation.
	StreamResponse_UPDATE StreamResponse_Operation = 2
	// Specifies a delete operation.
	StreamResponse_DELETE StreamResponse_Operation = 3
)

var StreamResponse_Operation_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
}

var StreamResponse_Operation_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATE":  1,
	"UPDATE":  2,
	"DELETE":  3,
}

func (x StreamResponse_Operation) String() string {
	return proto.EnumName(StreamResponse_Operation_name, int32(x))
}

func (StreamResponse_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{1, 0}
}

// StreamRequest represents a streaming request sent by a federated service mesh
// consumer to a federated service mesh owner.
//
// Following are some of the valid field combinations for a StreamRequest.
//
// * Initial request:
//     resource_url   != ""
//     response_nonce == ""
//     status         == nil
//
// * ACK, request:
//     resource_url   != ""
//     response_nonce != ""
//     status         != nil
//       code         == OK
//       message      == ""
//
// * NACK, request:
//     resource_url   != ""
//     response_nonce != ""
//     status         != nil
//       code         != OK
//       message      != ""
type StreamRequest struct {
	// REQUIRED. The type URL of the resource being requested.
	// Example: "type.googleapis.com/federation.types.v1alpha1.FederatedService"
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// The nonce of a consumed StreamResponse to ACK/NACK. This field should be
	// omitted in the first StreamRequest sent by the federated service mesh
	// consumer.
	ResponseNonce string `protobuf:"bytes,2,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// The message consumption status of the client. This field should be omitted
	// in the first StreamRequest sent by the federated service mesh consumer. If
	// the message is consumed successfully, the status code should be set to OK.
	// If there is a failure in consuming a message, an appropriate status code
	// must be set along with the error details in the status' message field.
	Status               *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *StreamRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *StreamRequest) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// StreamResponse represents a streaming response sent by a federated service
// mesh owner to a federated service mesh consumer.
type StreamResponse struct {
	// REQUIRED. The unique identifier for a StreamResponse.
	Nonce string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// REQUIRED. The type URL of the resource being returned in the response.
	// Example: "type.googleapis.com/federation.types.v1alpha1.FederatedService"
	ResourceUrl string `protobuf:"bytes,2,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// REQUIRED. The typed resource in relation to which an operation is to be
	// performed.
	Resource *any.Any `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// REQUIRED. The operation to be performed in relation to the resource.
	Operation            StreamResponse_Operation `protobuf:"varint,4,opt,name=operation,proto3,enum=federation.resourcediscovery.v1alpha1.StreamResponse_Operation" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *StreamResponse) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *StreamResponse) GetResource() *any.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *StreamResponse) GetOperation() StreamResponse_Operation {
	if m != nil {
		return m.Operation
	}
	return StreamResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("federation.resourcediscovery.v1alpha1.StreamResponse_Operation", StreamResponse_Operation_name, StreamResponse_Operation_value)
	proto.RegisterType((*StreamRequest)(nil), "federation.resourcediscovery.v1alpha1.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "federation.resourcediscovery.v1alpha1.StreamResponse")
}

func init() { proto.RegisterFile("resource_discovery.proto", fileDescriptor_af67c44042f17abc) }

var fileDescriptor_af67c44042f17abc = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xd1, 0x4b, 0xc2, 0x50,
	0x14, 0xc6, 0xbb, 0xb3, 0x2c, 0xaf, 0x69, 0xe3, 0x22, 0xb4, 0x46, 0x0f, 0x26, 0x08, 0xa3, 0x87,
	0x3b, 0xb5, 0x7a, 0x0b, 0xc2, 0x72, 0x4f, 0xc5, 0x8c, 0xa9, 0x04, 0x41, 0xc8, 0x9c, 0x57, 0x1d,
	0xac, 0xdd, 0x75, 0xef, 0x26, 0xf8, 0x1a, 0xf4, 0x37, 0xd4, 0x9f, 0x1b, 0xdb, 0xdd, 0x55, 0xc2,
	0x97, 0xea, 0xed, 0xdc, 0x73, 0xbe, 0xf3, 0x9d, 0xdf, 0xbe, 0x41, 0x8d, 0x11, 0x4e, 0x13, 0xe6,
	0x91, 0xf1, 0xd4, 0xe7, 0x1e, 0x5d, 0x12, 0xb6, 0xc2, 0x11, 0xa3, 0x31, 0x45, 0xcd, 0x19, 0x99,
	0x12, 0xe6, 0xc6, 0x3e, 0x0d, 0xb1, 0x14, 0x6d, 0x34, 0xcb, 0xb6, 0x1b, 0x44, 0x0b, 0xb7, 0xad,
	0x9f, 0xcc, 0x29, 0x9d, 0x07, 0xc4, 0xcc, 0x96, 0x26, 0xc9, 0xcc, 0x74, 0xc3, 0xdc, 0x41, 0x3f,
	0xce, 0x47, 0x2c, 0xf2, 0x4c, 0x1e, 0xbb, 0x71, 0xc2, 0xc5, 0xa0, 0xf1, 0x01, 0x60, 0x65, 0x10,
	0x33, 0xe2, 0xbe, 0x3a, 0xe4, 0x2d, 0x21, 0x3c, 0x46, 0x67, 0xf0, 0x70, 0x0d, 0x92, 0xb0, 0x40,
	0x03, 0x75, 0x60, 0x94, 0x9c, 0xb2, 0xec, 0x8d, 0x58, 0x80, 0x9a, 0xb0, 0xca, 0x08, 0x8f, 0x68,
	0xc8, 0xc9, 0x38, 0xa4, 0xa1, 0x47, 0x34, 0x25, 0x13, 0x55, 0x64, 0xd7, 0x4e, 0x9b, 0xe8, 0x1c,
	0x16, 0xc5, 0x2d, 0xad, 0x50, 0x07, 0x46, 0xb9, 0x83, 0xb0, 0xa0, 0xc0, 0x2c, 0xf2, 0xf0, 0x20,
	0x9b, 0x38, 0xb9, 0xa2, 0xf1, 0xa5, 0xc0, 0xaa, 0xe4, 0x10, 0x1e, 0xa8, 0x06, 0xf7, 0x84, 0xb9,
	0x20, 0x10, 0x8f, 0x2d, 0x3c, 0x65, 0x1b, 0xaf, 0x05, 0x0f, 0xe4, 0x33, 0xbf, 0x5c, 0x93, 0x97,
	0x65, 0x34, 0xb8, 0x1b, 0xae, 0x9c, 0xb5, 0x0a, 0xbd, 0xc0, 0x12, 0x8d, 0xf2, 0x84, 0xb5, 0xdd,
	0x3a, 0x30, 0xaa, 0x9d, 0x1b, 0xfc, 0xab, 0xd0, 0xf1, 0x4f, 0x68, 0xdc, 0x97, 0x36, 0xce, 0xc6,
	0xb1, 0x71, 0x0d, 0x4b, 0xeb, 0x3e, 0x2a, 0xc3, 0xfd, 0x91, 0x7d, 0x6f, 0xf7, 0x9f, 0x6c, 0x75,
	0x07, 0x41, 0x58, 0xbc, 0x73, 0xac, 0xee, 0xd0, 0x52, 0x41, 0x5a, 0x8f, 0x1e, 0x7b, 0x69, 0xad,
	0xa4, 0x75, 0xcf, 0x7a, 0xb0, 0x86, 0x96, 0x5a, 0xe8, 0x7c, 0x02, 0xa8, 0xf6, 0xe4, 0xe1, 0x01,
	0x61, 0x4b, 0xdf, 0x23, 0xe8, 0x1d, 0xc0, 0x23, 0x8b, 0xc7, 0xee, 0x24, 0xf0, 0xf9, 0x42, 0x30,
	0xa0, 0xcb, 0x3f, 0x22, 0x67, 0xff, 0x5b, 0xbf, 0xfa, 0xd7, 0x87, 0x1a, 0xa0, 0x05, 0x6e, 0x4f,
	0x9f, 0xf5, 0xad, 0x05, 0x53, 0x2e, 0x4c, 0x8a, 0x59, 0xd8, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0x70, 0x30, 0xd0, 0xd8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	// Establish a new stream with a federated service mesh owner. The federated
	// service mesh owner can then send StreamResponse messages, and the federated
	// service mesh consumer can ACK/NACK these via new StreamRequest messages.
	EstablishStream(ctx context.Context, opts ...grpc.CallOption) (DiscoveryService_EstablishStreamClient, error)
}

type discoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryServiceClient(cc *grpc.ClientConn) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) EstablishStream(ctx context.Context, opts ...grpc.CallOption) (DiscoveryService_EstablishStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DiscoveryService_serviceDesc.Streams[0], "/federation.resourcediscovery.v1alpha1.DiscoveryService/EstablishStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &discoveryServiceEstablishStreamClient{stream}
	return x, nil
}

type DiscoveryService_EstablishStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type discoveryServiceEstablishStreamClient struct {
	grpc.ClientStream
}

func (x *discoveryServiceEstablishStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *discoveryServiceEstablishStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
type DiscoveryServiceServer interface {
	// Establish a new stream with a federated service mesh owner. The federated
	// service mesh owner can then send StreamResponse messages, and the federated
	// service mesh consumer can ACK/NACK these via new StreamRequest messages.
	EstablishStream(DiscoveryService_EstablishStreamServer) error
}

func RegisterDiscoveryServiceServer(s *grpc.Server, srv DiscoveryServiceServer) {
	s.RegisterService(&_DiscoveryService_serviceDesc, srv)
}

func _DiscoveryService_EstablishStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DiscoveryServiceServer).EstablishStream(&discoveryServiceEstablishStreamServer{stream})
}

type DiscoveryService_EstablishStreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type discoveryServiceEstablishStreamServer struct {
	grpc.ServerStream
}

func (x *discoveryServiceEstablishStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *discoveryServiceEstablishStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "federation.resourcediscovery.v1alpha1.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishStream",
			Handler:       _DiscoveryService_EstablishStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "resource_discovery.proto",
}
