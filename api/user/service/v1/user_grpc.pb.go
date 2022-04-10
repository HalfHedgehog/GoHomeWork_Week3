// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	SignIn(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*SignInRes, error)
	LogIn(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*LogInRes, error)
	GetUserInfo(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*UserInfoRes, error)
	UpdateUserInfo(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*UpdateInfoRes, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SignIn(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*SignInRes, error) {
	out := new(SignInRes)
	err := c.cc.Invoke(ctx, "/user.server.v1.SearchService/signIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) LogIn(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*LogInRes, error) {
	out := new(LogInRes)
	err := c.cc.Invoke(ctx, "/user.server.v1.SearchService/logIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetUserInfo(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, "/user.server.v1.SearchService/getUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) UpdateUserInfo(ctx context.Context, in *UserDto, opts ...grpc.CallOption) (*UpdateInfoRes, error) {
	out := new(UpdateInfoRes)
	err := c.cc.Invoke(ctx, "/user.server.v1.SearchService/updateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	SignIn(context.Context, *UserDto) (*SignInRes, error)
	LogIn(context.Context, *UserDto) (*LogInRes, error)
	GetUserInfo(context.Context, *UserDto) (*UserInfoRes, error)
	UpdateUserInfo(context.Context, *UserDto) (*UpdateInfoRes, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SignIn(context.Context, *UserDto) (*SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedSearchServiceServer) LogIn(context.Context, *UserDto) (*LogInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedSearchServiceServer) GetUserInfo(context.Context, *UserDto) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedSearchServiceServer) UpdateUserInfo(context.Context, *UserDto) (*UpdateInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.server.v1.SearchService/signIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SignIn(ctx, req.(*UserDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.server.v1.SearchService/logIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).LogIn(ctx, req.(*UserDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.server.v1.SearchService/getUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetUserInfo(ctx, req.(*UserDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.server.v1.SearchService/updateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).UpdateUserInfo(ctx, req.(*UserDto))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.server.v1.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "signIn",
			Handler:    _SearchService_SignIn_Handler,
		},
		{
			MethodName: "logIn",
			Handler:    _SearchService_LogIn_Handler,
		},
		{
			MethodName: "getUserInfo",
			Handler:    _SearchService_GetUserInfo_Handler,
		},
		{
			MethodName: "updateUserInfo",
			Handler:    _SearchService_UpdateUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
