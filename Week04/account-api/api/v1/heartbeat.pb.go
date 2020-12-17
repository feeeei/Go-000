// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package account_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HeartbeatType int32

const (
	HeartbeatType_Ping HeartbeatType = 0
	HeartbeatType_Pong HeartbeatType = 1
)

var HeartbeatType_name = map[int32]string{
	0: "Ping",
	1: "Pong",
}

var HeartbeatType_value = map[string]int32{
	"Ping": 0,
	"Pong": 1,
}

func (x HeartbeatType) String() string {
	return proto.EnumName(HeartbeatType_name, int32(x))
}

func (HeartbeatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{0}
}

type Heart struct {
	Type                 HeartbeatType `protobuf:"varint,1,opt,name=type,proto3,enum=account.v1.HeartbeatType" json:"type,omitempty"`
	Ts                   int64         `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Heart) Reset()         { *m = Heart{} }
func (m *Heart) String() string { return proto.CompactTextString(m) }
func (*Heart) ProtoMessage()    {}
func (*Heart) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{0}
}

func (m *Heart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heart.Unmarshal(m, b)
}
func (m *Heart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heart.Marshal(b, m, deterministic)
}
func (m *Heart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heart.Merge(m, src)
}
func (m *Heart) XXX_Size() int {
	return xxx_messageInfo_Heart.Size(m)
}
func (m *Heart) XXX_DiscardUnknown() {
	xxx_messageInfo_Heart.DiscardUnknown(m)
}

var xxx_messageInfo_Heart proto.InternalMessageInfo

func (m *Heart) GetType() HeartbeatType {
	if m != nil {
		return m.Type
	}
	return HeartbeatType_Ping
}

func (m *Heart) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func init() {
	proto.RegisterEnum("account.v1.HeartbeatType", HeartbeatType_name, HeartbeatType_value)
	proto.RegisterType((*Heart)(nil), "account.v1.Heart")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_3c667767fb9826a9) }

var fileDescriptor_3c667767fb9826a9 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4c, 0x4e,
	0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x33, 0x54, 0x72, 0xe3, 0x62, 0xf5, 0x00, 0x49, 0x0b, 0xe9,
	0x72, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x49, 0xea, 0x21,
	0xd4, 0xe8, 0x79, 0xc0, 0xf4, 0x87, 0x54, 0x16, 0xa4, 0x06, 0x81, 0x95, 0x09, 0xf1, 0x71, 0x31,
	0x95, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x95, 0x14, 0x6b, 0x29, 0x73, 0xf1,
	0xa2, 0x28, 0x13, 0xe2, 0xe0, 0x62, 0x09, 0xc8, 0xcc, 0x4b, 0x17, 0x60, 0x00, 0xb3, 0xf2, 0xf3,
	0xd2, 0x05, 0x18, 0x8d, 0x1c, 0xb8, 0x38, 0xe1, 0x8a, 0x84, 0x8c, 0x91, 0x39, 0x82, 0x18, 0xf6,
	0x49, 0x61, 0x0a, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x7d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x69, 0xc9, 0x66, 0x08, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	Heartbeat(ctx context.Context, in *Heart, opts ...grpc.CallOption) (*Heart, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Heartbeat(ctx context.Context, in *Heart, opts ...grpc.CallOption) (*Heart, error) {
	out := new(Heart)
	err := c.cc.Invoke(ctx, "/account.v1.Heartbeat/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	Heartbeat(context.Context, *Heart) (*Heart, error)
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Heart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1.Heartbeat/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Heartbeat(ctx, req.(*Heart))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.v1.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _Heartbeat_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}