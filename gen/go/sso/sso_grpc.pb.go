//protoc -I proto proto/sso/sso.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: sso/sso.proto

package ssov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Auth_Register_FullMethodName = "/Auth.Auth/Register"
	Auth_Login_FullMethodName    = "/Auth.Auth/Login"
	Auth_IsAdimin_FullMethodName = "/Auth.Auth/IsAdimin"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	IsAdimin(ctx context.Context, in *IsAdiminRequest, opts ...grpc.CallOption) (*IsAdiminResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Auth_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) IsAdimin(ctx context.Context, in *IsAdiminRequest, opts ...grpc.CallOption) (*IsAdiminResponse, error) {
	out := new(IsAdiminResponse)
	err := c.cc.Invoke(ctx, Auth_IsAdimin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	IsAdimin(context.Context, *IsAdiminRequest) (*IsAdiminResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) IsAdimin(context.Context, *IsAdiminRequest) (*IsAdiminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAdimin not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_IsAdimin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAdiminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsAdimin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_IsAdimin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsAdimin(ctx, req.(*IsAdiminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "IsAdimin",
			Handler:    _Auth_IsAdimin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}

const (
	BannerService_GetBanner_FullMethodName    = "/Auth.BannerService/GetBanner"
	BannerService_AddBanner_FullMethodName    = "/Auth.BannerService/AddBanner"
	BannerService_UpdateBanner_FullMethodName = "/Auth.BannerService/UpdateBanner"
	BannerService_DeleteBanner_FullMethodName = "/Auth.BannerService/DeleteBanner"
	BannerService_AddTag_FullMethodName       = "/Auth.BannerService/AddTag"
	BannerService_UpdateTag_FullMethodName    = "/Auth.BannerService/UpdateTag"
	BannerService_DeleteTag_FullMethodName    = "/Auth.BannerService/DeleteTag"
)

// BannerServiceClient is the client API for BannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BannerServiceClient interface {
	GetBanner(ctx context.Context, in *GetBannerRequest, opts ...grpc.CallOption) (*GetBannerResponse, error)
	AddBanner(ctx context.Context, in *Banner, opts ...grpc.CallOption) (*Banner, error)
	UpdateBanner(ctx context.Context, in *Banner, opts ...grpc.CallOption) (*Banner, error)
	DeleteBanner(ctx context.Context, in *DeleteBannerRequest, opts ...grpc.CallOption) (*Banner, error)
	// rpc GetTags (google.protobuf.Empty) returns (GetTagsResponse);
	AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
	UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*Tag, error)
}

type bannerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBannerServiceClient(cc grpc.ClientConnInterface) BannerServiceClient {
	return &bannerServiceClient{cc}
}

func (c *bannerServiceClient) GetBanner(ctx context.Context, in *GetBannerRequest, opts ...grpc.CallOption) (*GetBannerResponse, error) {
	out := new(GetBannerResponse)
	err := c.cc.Invoke(ctx, BannerService_GetBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) AddBanner(ctx context.Context, in *Banner, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, BannerService_AddBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) UpdateBanner(ctx context.Context, in *Banner, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, BannerService_UpdateBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) DeleteBanner(ctx context.Context, in *DeleteBannerRequest, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, BannerService_DeleteBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, BannerService_AddTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, BannerService_UpdateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, BannerService_DeleteTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BannerServiceServer is the server API for BannerService service.
// All implementations must embed UnimplementedBannerServiceServer
// for forward compatibility
type BannerServiceServer interface {
	GetBanner(context.Context, *GetBannerRequest) (*GetBannerResponse, error)
	AddBanner(context.Context, *Banner) (*Banner, error)
	UpdateBanner(context.Context, *Banner) (*Banner, error)
	DeleteBanner(context.Context, *DeleteBannerRequest) (*Banner, error)
	// rpc GetTags (google.protobuf.Empty) returns (GetTagsResponse);
	AddTag(context.Context, *Tag) (*Tag, error)
	UpdateTag(context.Context, *Tag) (*Tag, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*Tag, error)
	mustEmbedUnimplementedBannerServiceServer()
}

// UnimplementedBannerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBannerServiceServer struct {
}

func (UnimplementedBannerServiceServer) GetBanner(context.Context, *GetBannerRequest) (*GetBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanner not implemented")
}
func (UnimplementedBannerServiceServer) AddBanner(context.Context, *Banner) (*Banner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBanner not implemented")
}
func (UnimplementedBannerServiceServer) UpdateBanner(context.Context, *Banner) (*Banner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
func (UnimplementedBannerServiceServer) DeleteBanner(context.Context, *DeleteBannerRequest) (*Banner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (UnimplementedBannerServiceServer) AddTag(context.Context, *Tag) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedBannerServiceServer) UpdateTag(context.Context, *Tag) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedBannerServiceServer) DeleteTag(context.Context, *DeleteTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedBannerServiceServer) mustEmbedUnimplementedBannerServiceServer() {}

// UnsafeBannerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BannerServiceServer will
// result in compilation errors.
type UnsafeBannerServiceServer interface {
	mustEmbedUnimplementedBannerServiceServer()
}

func RegisterBannerServiceServer(s grpc.ServiceRegistrar, srv BannerServiceServer) {
	s.RegisterService(&BannerService_ServiceDesc, srv)
}

func _BannerService_GetBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).GetBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_GetBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).GetBanner(ctx, req.(*GetBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_AddBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Banner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).AddBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_AddBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).AddBanner(ctx, req.(*Banner))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_UpdateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Banner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).UpdateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_UpdateBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).UpdateBanner(ctx, req.(*Banner))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_DeleteBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).DeleteBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_DeleteBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).DeleteBanner(ctx, req.(*DeleteBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_AddTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).AddTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).UpdateTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BannerService_ServiceDesc is the grpc.ServiceDesc for BannerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BannerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auth.BannerService",
	HandlerType: (*BannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBanner",
			Handler:    _BannerService_GetBanner_Handler,
		},
		{
			MethodName: "AddBanner",
			Handler:    _BannerService_AddBanner_Handler,
		},
		{
			MethodName: "UpdateBanner",
			Handler:    _BannerService_UpdateBanner_Handler,
		},
		{
			MethodName: "DeleteBanner",
			Handler:    _BannerService_DeleteBanner_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _BannerService_AddTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _BannerService_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _BannerService_DeleteTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
