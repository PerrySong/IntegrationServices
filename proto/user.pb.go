// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package userpb

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

type UserGithubBasicResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Bio                  string   `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	Company              string   `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	ReposUrl             string   `protobuf:"bytes,7,opt,name=repos_url,json=reposUrl,proto3" json:"repos_url,omitempty"`
	Email                string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Location             string   `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	PublicRepos          int32    `protobuf:"varint,10,opt,name=public_repos,json=publicRepos,proto3" json:"public_repos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGithubBasicResponse) Reset()         { *m = UserGithubBasicResponse{} }
func (m *UserGithubBasicResponse) String() string { return proto.CompactTextString(m) }
func (*UserGithubBasicResponse) ProtoMessage()    {}
func (*UserGithubBasicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *UserGithubBasicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGithubBasicResponse.Unmarshal(m, b)
}
func (m *UserGithubBasicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGithubBasicResponse.Marshal(b, m, deterministic)
}
func (m *UserGithubBasicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGithubBasicResponse.Merge(m, src)
}
func (m *UserGithubBasicResponse) XXX_Size() int {
	return xxx_messageInfo_UserGithubBasicResponse.Size(m)
}
func (m *UserGithubBasicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGithubBasicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGithubBasicResponse proto.InternalMessageInfo

func (m *UserGithubBasicResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserGithubBasicResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserGithubBasicResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UserGithubBasicResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *UserGithubBasicResponse) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *UserGithubBasicResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *UserGithubBasicResponse) GetReposUrl() string {
	if m != nil {
		return m.ReposUrl
	}
	return ""
}

func (m *UserGithubBasicResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserGithubBasicResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UserGithubBasicResponse) GetPublicRepos() int32 {
	if m != nil {
		return m.PublicRepos
	}
	return 0
}

type UserGithubBasicRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGithubBasicRequest) Reset()         { *m = UserGithubBasicRequest{} }
func (m *UserGithubBasicRequest) String() string { return proto.CompactTextString(m) }
func (*UserGithubBasicRequest) ProtoMessage()    {}
func (*UserGithubBasicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *UserGithubBasicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGithubBasicRequest.Unmarshal(m, b)
}
func (m *UserGithubBasicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGithubBasicRequest.Marshal(b, m, deterministic)
}
func (m *UserGithubBasicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGithubBasicRequest.Merge(m, src)
}
func (m *UserGithubBasicRequest) XXX_Size() int {
	return xxx_messageInfo_UserGithubBasicRequest.Size(m)
}
func (m *UserGithubBasicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGithubBasicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGithubBasicRequest proto.InternalMessageInfo

func (m *UserGithubBasicRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*UserGithubBasicResponse)(nil), "user.UserGithubBasicResponse")
	proto.RegisterType((*UserGithubBasicRequest)(nil), "user.UserGithubBasicRequest")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5) }

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0xfd, 0xda, 0xb5, 0x5d, 0x7b, 0xf7, 0x29, 0x23, 0x88, 0x86, 0xe9, 0xa0, 0xf6, 0xa9, 0x4f,
	0x13, 0xf4, 0x1f, 0xec, 0x65, 0xf8, 0xe2, 0x43, 0x65, 0x2f, 0x82, 0x8c, 0xa4, 0x8b, 0x18, 0x68,
	0x9b, 0x98, 0xa4, 0x03, 0xff, 0x87, 0x3f, 0x58, 0x6e, 0xc2, 0x44, 0x50, 0xdf, 0xee, 0x39, 0xe7,
	0x9e, 0x93, 0x70, 0x2e, 0xcc, 0xb5, 0x51, 0x4e, 0xdd, 0x8c, 0x56, 0x98, 0x95, 0x1f, 0x49, 0x82,
	0x73, 0xf5, 0x11, 0xc3, 0xc5, 0xd6, 0x0a, 0xb3, 0x91, 0xee, 0x75, 0xe4, 0x6b, 0x66, 0x65, 0xdb,
	0x08, 0xab, 0xd5, 0x60, 0x05, 0x39, 0x85, 0x58, 0xee, 0x69, 0x54, 0x46, 0xf5, 0xa4, 0x89, 0xe5,
	0x9e, 0x2c, 0x20, 0x47, 0xcf, 0xc0, 0x7a, 0x41, 0xe3, 0x32, 0xaa, 0x8b, 0xe6, 0x0b, 0x93, 0x39,
	0x4c, 0x46, 0xd3, 0xd1, 0x89, 0xa7, 0x71, 0x24, 0x4b, 0x00, 0x76, 0x60, 0x8e, 0x99, 0x1d, 0x0a,
	0x89, 0x17, 0x8a, 0xc0, 0x6c, 0x4d, 0x87, 0x06, 0x2e, 0x15, 0x4d, 0x83, 0x81, 0x4b, 0x45, 0x28,
	0x4c, 0x5b, 0xd5, 0x6b, 0x36, 0xbc, 0xd3, 0xcc, 0xb3, 0x47, 0x48, 0x2e, 0xa1, 0x30, 0x42, 0x2b,
	0xeb, 0x93, 0xa6, 0xe1, 0x65, 0x4f, 0x60, 0xd0, 0x19, 0xa4, 0xa2, 0x67, 0xb2, 0xa3, 0xb9, 0x17,
	0x02, 0xc0, 0xbf, 0x76, 0xaa, 0x65, 0x4e, 0xaa, 0x81, 0x16, 0xc1, 0x71, 0xc4, 0xe4, 0x1a, 0xfe,
	0xeb, 0x91, 0x77, 0xb2, 0xdd, 0xf9, 0x10, 0x0a, 0x65, 0x54, 0xa7, 0xcd, 0x2c, 0x70, 0x0d, 0x52,
	0x55, 0x0d, 0xe7, 0x3f, 0x5a, 0x79, 0x1b, 0x85, 0x75, 0xdf, 0x4a, 0x49, 0xb0, 0x94, 0xdb, 0x67,
	0x98, 0xe1, 0xe6, 0xa3, 0x30, 0x07, 0xd9, 0x0a, 0xf2, 0x00, 0x27, 0x1b, 0xe1, 0x82, 0xef, 0x7e,
	0x78, 0x51, 0xe4, 0x6a, 0xe5, 0x3b, 0xff, 0x3d, 0x6d, 0xb1, 0xfc, 0x43, 0x0d, 0x17, 0xa8, 0xfe,
	0xad, 0xf3, 0xa7, 0x0c, 0x37, 0x34, 0xe7, 0x99, 0x3f, 0xdb, 0xdd, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x86, 0xe9, 0x02, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetGithubInfo(ctx context.Context, in *UserGithubBasicRequest, opts ...grpc.CallOption) (*UserGithubBasicResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetGithubInfo(ctx context.Context, in *UserGithubBasicRequest, opts ...grpc.CallOption) (*UserGithubBasicResponse, error) {
	out := new(UserGithubBasicResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetGithubInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetGithubInfo(context.Context, *UserGithubBasicRequest) (*UserGithubBasicResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetGithubInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGithubBasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetGithubInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetGithubInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetGithubInfo(ctx, req.(*UserGithubBasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGithubInfo",
			Handler:    _UserService_GetGithubInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}