// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello_message.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d008acd07f1cf888, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d008acd07f1cf888, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.message.example.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "hello.message.example.HelloResponse")
}

func init() {
	proto.RegisterFile("hello_message.proto", fileDescriptor_d008acd07f1cf888)
}

var fileDescriptor_d008acd07f1cf888 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x05, 0x0b, 0xea, 0xc1, 0x04, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x95, 0x94, 0xb8,
	0x78, 0x3c, 0x40, 0x12, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x22, 0x17, 0x2f,
	0x54, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x71, 0x3a, 0x54,
	0x0d, 0x88, 0x69, 0x94, 0x02, 0x35, 0x26, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x84,
	0x8b, 0x15, 0xcc, 0x17, 0x52, 0xd6, 0xc3, 0x6a, 0xaf, 0x1e, 0xb2, 0xa5, 0x52, 0x2a, 0xf8, 0x15,
	0x41, 0x6c, 0x55, 0x62, 0x70, 0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x7b, 0xc8, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x4c, 0x4c, 0xe3, 0xe7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.message.example.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.message.example.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.message.example.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_message.proto",
}
