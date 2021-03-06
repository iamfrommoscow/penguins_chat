// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package microChat

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

type User struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	ID                   uint64   `protobuf:"varint,2,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "microChat.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcd, 0x4c, 0x2e, 0xca, 0x77, 0xce, 0x48,
	0x2c, 0x51, 0xd2, 0xe1, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0x12, 0xe1, 0x62, 0xcd, 0xc9, 0x4f,
	0xcf, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xf8, 0xb8, 0x98, 0x3c,
	0x5d, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x5d, 0x8c, 0xac, 0xb8, 0xb8, 0x41,
	0xaa, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x53, 0x8b, 0x84, 0xb4, 0xb9, 0x58, 0xc1, 0x4c, 0x21, 0x7e,
	0x3d, 0xb8, 0x89, 0x7a, 0x20, 0x05, 0x52, 0xe8, 0x02, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xbb, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x88, 0xcd, 0x82, 0x89, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserCheckerClient is the client API for UserChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserCheckerClient interface {
	Check(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type userCheckerClient struct {
	cc *grpc.ClientConn
}

func NewUserCheckerClient(cc *grpc.ClientConn) UserCheckerClient {
	return &userCheckerClient{cc}
}

func (c *userCheckerClient) Check(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/microChat.UserChecker/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCheckerServer is the server API for UserChecker service.
type UserCheckerServer interface {
	Check(context.Context, *User) (*User, error)
}

func RegisterUserCheckerServer(s *grpc.Server, srv UserCheckerServer) {
	s.RegisterService(&_UserChecker_serviceDesc, srv)
}

func _UserChecker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microChat.UserChecker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCheckerServer).Check(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "microChat.UserChecker",
	HandlerType: (*UserCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _UserChecker_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
