// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

package proto

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

// ValidationBearerTokenRequest validates auth token
type ValidationBearerTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationBearerTokenRequest) Reset()         { *m = ValidationBearerTokenRequest{} }
func (m *ValidationBearerTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidationBearerTokenRequest) ProtoMessage()    {}
func (*ValidationBearerTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

func (m *ValidationBearerTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationBearerTokenRequest.Unmarshal(m, b)
}
func (m *ValidationBearerTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationBearerTokenRequest.Marshal(b, m, deterministic)
}
func (m *ValidationBearerTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationBearerTokenRequest.Merge(m, src)
}
func (m *ValidationBearerTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidationBearerTokenRequest.Size(m)
}
func (m *ValidationBearerTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationBearerTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationBearerTokenRequest proto.InternalMessageInfo

func (m *ValidationBearerTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// ValidationBearerTokenResponse return auth token information after successful verification
type ValidationBearerTokenResponse struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Scope                string   `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationBearerTokenResponse) Reset()         { *m = ValidationBearerTokenResponse{} }
func (m *ValidationBearerTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidationBearerTokenResponse) ProtoMessage()    {}
func (*ValidationBearerTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1}
}

func (m *ValidationBearerTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationBearerTokenResponse.Unmarshal(m, b)
}
func (m *ValidationBearerTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationBearerTokenResponse.Marshal(b, m, deterministic)
}
func (m *ValidationBearerTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationBearerTokenResponse.Merge(m, src)
}
func (m *ValidationBearerTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidationBearerTokenResponse.Size(m)
}
func (m *ValidationBearerTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationBearerTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationBearerTokenResponse proto.InternalMessageInfo

func (m *ValidationBearerTokenResponse) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ValidationBearerTokenResponse) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ValidationBearerTokenResponse) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

// CreateClientRequest creates client credentials
type CreateClientRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClientRequest) Reset()         { *m = CreateClientRequest{} }
func (m *CreateClientRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClientRequest) ProtoMessage()    {}
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{2}
}

func (m *CreateClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientRequest.Unmarshal(m, b)
}
func (m *CreateClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientRequest.Marshal(b, m, deterministic)
}
func (m *CreateClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientRequest.Merge(m, src)
}
func (m *CreateClientRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClientRequest.Size(m)
}
func (m *CreateClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientRequest proto.InternalMessageInfo

func (m *CreateClientRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CreateClientRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

// CreateClientResponse new client credentials
type CreateClientResponse struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
	UserID               string   `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	Domain               string   `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClientResponse) Reset()         { *m = CreateClientResponse{} }
func (m *CreateClientResponse) String() string { return proto.CompactTextString(m) }
func (*CreateClientResponse) ProtoMessage()    {}
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{3}
}

func (m *CreateClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientResponse.Unmarshal(m, b)
}
func (m *CreateClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientResponse.Marshal(b, m, deterministic)
}
func (m *CreateClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientResponse.Merge(m, src)
}
func (m *CreateClientResponse) XXX_Size() int {
	return xxx_messageInfo_CreateClientResponse.Size(m)
}
func (m *CreateClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientResponse proto.InternalMessageInfo

func (m *CreateClientResponse) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *CreateClientResponse) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *CreateClientResponse) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CreateClientResponse) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidationBearerTokenRequest)(nil), "proto.ValidationBearerTokenRequest")
	proto.RegisterType((*ValidationBearerTokenResponse)(nil), "proto.ValidationBearerTokenResponse")
	proto.RegisterType((*CreateClientRequest)(nil), "proto.CreateClientRequest")
	proto.RegisterType((*CreateClientResponse)(nil), "proto.CreateClientResponse")
}

func init() {
	proto.RegisterFile("authentication.proto", fileDescriptor_d0dbc99083440df2)
}

var fileDescriptor_d0dbc99083440df2 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0x65, 0x4a, 0x2b, 0x3a, 0xea, 0xca, 0xb4, 0x55, 0x14, 0x40, 0x42, 0xa1, 0x0b, 0x56,
	0x59, 0x00, 0x17, 0x20, 0x2d, 0x42, 0xdd, 0x55, 0x2d, 0x62, 0xef, 0x3a, 0x23, 0xb0, 0x08, 0xb6,
	0xb1, 0x1d, 0x8e, 0xc0, 0xc9, 0x38, 0x18, 0x8a, 0x9d, 0x42, 0x22, 0xa5, 0x15, 0xab, 0xe4, 0x67,
	0xf2, 0x9f, 0xff, 0xcc, 0x18, 0xc6, 0xac, 0x74, 0xaf, 0x28, 0x9d, 0xe0, 0xcc, 0x09, 0x25, 0x53,
	0x6d, 0x94, 0x53, 0xb4, 0xef, 0x1f, 0xc9, 0x1d, 0x9c, 0x3f, 0xb3, 0x42, 0xe4, 0xbe, 0x94, 0x21,
	0x33, 0x68, 0x9e, 0xd4, 0x1b, 0xca, 0x35, 0x7e, 0x94, 0x68, 0x1d, 0x1d, 0x43, 0xdf, 0x55, 0x3a,
	0x22, 0x97, 0xe4, 0x7a, 0xb8, 0x0e, 0x22, 0x11, 0x70, 0xb1, 0xc7, 0x65, 0xb5, 0x92, 0x16, 0x69,
	0x0c, 0x27, 0xbc, 0x10, 0x28, 0xdd, 0x72, 0x51, 0x3b, 0x7f, 0x35, 0x9d, 0xc2, 0xa0, 0xb4, 0x68,
	0x96, 0x8b, 0xe8, 0xc8, 0x57, 0x6a, 0x55, 0x1d, 0x65, 0xb9, 0xd2, 0x18, 0xf5, 0xc2, 0x51, 0x5e,
	0x24, 0x0f, 0x70, 0x3a, 0x37, 0xc8, 0x1c, 0xce, 0xbd, 0x7f, 0x97, 0xeb, 0x0f, 0x42, 0x5a, 0x90,
	0x29, 0x0c, 0x72, 0xf5, 0xce, 0x84, 0xdc, 0xc1, 0x83, 0x4a, 0xbe, 0x08, 0x8c, 0xdb, 0x9c, 0x7f,
	0x24, 0x4d, 0x60, 0x14, 0xde, 0x37, 0xc8, 0x0d, 0xba, 0x1a, 0xd9, 0xfa, 0xd6, 0x08, 0xd2, 0xdb,
	0x13, 0xe4, 0xb8, 0x19, 0xe4, 0xe6, 0x9b, 0xc0, 0xe4, 0xbe, 0xb5, 0x90, 0x0d, 0x9a, 0x4f, 0xc1,
	0x91, 0x6e, 0x61, 0xd2, 0x39, 0x54, 0x7a, 0x15, 0x56, 0x96, 0x1e, 0x5a, 0x54, 0x3c, 0x3b, 0xfc,
	0x53, 0xdd, 0xed, 0x23, 0x8c, 0x9a, 0x53, 0xa0, 0x71, 0xed, 0xea, 0x18, 0x71, 0x7c, 0xd6, 0x59,
	0x0b, 0xa0, 0x6c, 0x06, 0x13, 0xce, 0x2c, 0xcb, 0x51, 0x33, 0x8d, 0x45, 0xfa, 0x62, 0x34, 0x4f,
	0xab, 0x7b, 0x96, 0x0d, 0xab, 0xe6, 0x56, 0x95, 0x71, 0x45, 0xb6, 0x03, 0x4f, 0xb8, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xd6, 0x13, 0xbf, 0x83, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	ValidationBearerToken(ctx context.Context, in *ValidationBearerTokenRequest, opts ...grpc.CallOption) (*ValidationBearerTokenResponse, error)
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) ValidationBearerToken(ctx context.Context, in *ValidationBearerTokenRequest, opts ...grpc.CallOption) (*ValidationBearerTokenResponse, error) {
	out := new(ValidationBearerTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthenticationService/ValidationBearerToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthenticationService/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	ValidationBearerToken(context.Context, *ValidationBearerTokenRequest) (*ValidationBearerTokenResponse, error)
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) ValidationBearerToken(ctx context.Context, req *ValidationBearerTokenRequest) (*ValidationBearerTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidationBearerToken not implemented")
}
func (*UnimplementedAuthenticationServiceServer) CreateClient(ctx context.Context, req *CreateClientRequest) (*CreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_ValidationBearerToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationBearerTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ValidationBearerToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthenticationService/ValidationBearerToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ValidationBearerToken(ctx, req.(*ValidationBearerTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthenticationService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidationBearerToken",
			Handler:    _AuthenticationService_ValidationBearerToken_Handler,
		},
		{
			MethodName: "CreateClient",
			Handler:    _AuthenticationService_CreateClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
