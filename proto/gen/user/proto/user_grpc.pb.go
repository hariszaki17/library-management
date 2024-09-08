// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetUserDetails_FullMethodName      = "/user.UserService/GetUserDetails"
	UserService_Authenticate_FullMethodName        = "/user.UserService/Authenticate"
	UserService_VerifyJWT_FullMethodName           = "/user.UserService/VerifyJWT"
	UserService_UserBorrowBook_FullMethodName      = "/user.UserService/UserBorrowBook"
	UserService_UserReturnBook_FullMethodName      = "/user.UserService/UserReturnBook"
	UserService_GetBorrowingCount_FullMethodName   = "/user.UserService/GetBorrowingCount"
	UserService_GetBorrowingRecords_FullMethodName = "/user.UserService/GetBorrowingRecords"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The UserService service definition
type UserServiceClient interface {
	GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	VerifyJWT(ctx context.Context, in *VerifyJWTRequest, opts ...grpc.CallOption) (*VerifyJWTResponse, error)
	UserBorrowBook(ctx context.Context, in *UserBorrowBookRequest, opts ...grpc.CallOption) (*UserBorrowBookResponse, error)
	UserReturnBook(ctx context.Context, in *UserReturnBookRequest, opts ...grpc.CallOption) (*UserReturnBookResponse, error)
	GetBorrowingCount(ctx context.Context, in *GetBorrowingCountRequest, opts ...grpc.CallOption) (*GetBorrowingCountResponse, error)
	GetBorrowingRecords(ctx context.Context, in *GetBorrowingRecordsRequest, opts ...grpc.CallOption) (*GetBorrowingRecordsResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, UserService_Authenticate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyJWT(ctx context.Context, in *VerifyJWTRequest, opts ...grpc.CallOption) (*VerifyJWTResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyJWTResponse)
	err := c.cc.Invoke(ctx, UserService_VerifyJWT_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserBorrowBook(ctx context.Context, in *UserBorrowBookRequest, opts ...grpc.CallOption) (*UserBorrowBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserBorrowBookResponse)
	err := c.cc.Invoke(ctx, UserService_UserBorrowBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserReturnBook(ctx context.Context, in *UserReturnBookRequest, opts ...grpc.CallOption) (*UserReturnBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReturnBookResponse)
	err := c.cc.Invoke(ctx, UserService_UserReturnBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBorrowingCount(ctx context.Context, in *GetBorrowingCountRequest, opts ...grpc.CallOption) (*GetBorrowingCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBorrowingCountResponse)
	err := c.cc.Invoke(ctx, UserService_GetBorrowingCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBorrowingRecords(ctx context.Context, in *GetBorrowingRecordsRequest, opts ...grpc.CallOption) (*GetBorrowingRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBorrowingRecordsResponse)
	err := c.cc.Invoke(ctx, UserService_GetBorrowingRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
//
// The UserService service definition
type UserServiceServer interface {
	GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	VerifyJWT(context.Context, *VerifyJWTRequest) (*VerifyJWTResponse, error)
	UserBorrowBook(context.Context, *UserBorrowBookRequest) (*UserBorrowBookResponse, error)
	UserReturnBook(context.Context, *UserReturnBookRequest) (*UserReturnBookResponse, error)
	GetBorrowingCount(context.Context, *GetBorrowingCountRequest) (*GetBorrowingCountResponse, error)
	GetBorrowingRecords(context.Context, *GetBorrowingRecordsRequest) (*GetBorrowingRecordsResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedUserServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedUserServiceServer) VerifyJWT(context.Context, *VerifyJWTRequest) (*VerifyJWTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyJWT not implemented")
}
func (UnimplementedUserServiceServer) UserBorrowBook(context.Context, *UserBorrowBookRequest) (*UserBorrowBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBorrowBook not implemented")
}
func (UnimplementedUserServiceServer) UserReturnBook(context.Context, *UserReturnBookRequest) (*UserReturnBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserReturnBook not implemented")
}
func (UnimplementedUserServiceServer) GetBorrowingCount(context.Context, *GetBorrowingCountRequest) (*GetBorrowingCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBorrowingCount not implemented")
}
func (UnimplementedUserServiceServer) GetBorrowingRecords(context.Context, *GetBorrowingRecordsRequest) (*GetBorrowingRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBorrowingRecords not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_VerifyJWT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyJWT(ctx, req.(*VerifyJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserBorrowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBorrowBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserBorrowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserBorrowBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserBorrowBook(ctx, req.(*UserBorrowBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserReturnBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReturnBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserReturnBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserReturnBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserReturnBook(ctx, req.(*UserReturnBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBorrowingCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBorrowingCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBorrowingCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBorrowingCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBorrowingCount(ctx, req.(*GetBorrowingCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBorrowingRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBorrowingRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBorrowingRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBorrowingRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBorrowingRecords(ctx, req.(*GetBorrowingRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserDetails",
			Handler:    _UserService_GetUserDetails_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UserService_Authenticate_Handler,
		},
		{
			MethodName: "VerifyJWT",
			Handler:    _UserService_VerifyJWT_Handler,
		},
		{
			MethodName: "UserBorrowBook",
			Handler:    _UserService_UserBorrowBook_Handler,
		},
		{
			MethodName: "UserReturnBook",
			Handler:    _UserService_UserReturnBook_Handler,
		},
		{
			MethodName: "GetBorrowingCount",
			Handler:    _UserService_GetBorrowingCount_Handler,
		},
		{
			MethodName: "GetBorrowingRecords",
			Handler:    _UserService_GetBorrowingRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
