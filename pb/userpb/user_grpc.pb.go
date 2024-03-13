// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user.proto

package userpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_SignupUser_FullMethodName             = "/user.UserService/SignupUser"
	UserService_GetRoles_FullMethodName               = "/user.UserService/GetRoles"
	UserService_SetStatus_FullMethodName              = "/user.UserService/SetStatus"
	UserService_GetByEmail_FullMethodName             = "/user.UserService/GetByEmail"
	UserService_SearchforMembers_FullMethodName       = "/user.UserService/SearchforMembers"
	UserService_AddRoles_FullMethodName               = "/user.UserService/AddRoles"
	UserService_GetUserDetails_FullMethodName         = "/user.UserService/GetUserDetails"
	UserService_GetStreamofUserDetails_FullMethodName = "/user.UserService/GetStreamofUserDetails"
	UserService_GetStreamofRoles_FullMethodName       = "/user.UserService/GetStreamofRoles"
	UserService_GetRole_FullMethodName                = "/user.UserService/GetRole"
	UserService_EditStatus_FullMethodName             = "/user.UserService/EditStatus"
	UserService_UpdateUserDetails_FullMethodName      = "/user.UserService/UpdateUserDetails"
	UserService_DropUserAccount_FullMethodName        = "/user.UserService/DropUserAccount"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	SignupUser(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*UserResponce, error)
	GetRoles(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (UserService_GetRolesClient, error)
	SetStatus(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetByEmail(ctx context.Context, in *GetByEmailReq, opts ...grpc.CallOption) (*GetByEmailRes, error)
	SearchforMembers(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (UserService_SearchforMembersClient, error)
	AddRoles(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserDetails(ctx context.Context, in *GetUserDetailsReq, opts ...grpc.CallOption) (*GetUserDetailsRes, error)
	GetStreamofUserDetails(ctx context.Context, opts ...grpc.CallOption) (UserService_GetStreamofUserDetailsClient, error)
	GetStreamofRoles(ctx context.Context, opts ...grpc.CallOption) (UserService_GetStreamofRolesClient, error)
	GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleRes, error)
	EditStatus(ctx context.Context, in *EditStatusReq, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateUserDetails(ctx context.Context, in *UpdateUserDetailsReq, opts ...grpc.CallOption) (*empty.Empty, error)
	DropUserAccount(ctx context.Context, in *DropUserAccountReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SignupUser(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, UserService_SignupUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetRoles(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (UserService_GetRolesClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], UserService_GetRoles_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetRolesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetRolesClient interface {
	Recv() (*Role, error)
	grpc.ClientStream
}

type userServiceGetRolesClient struct {
	grpc.ClientStream
}

func (x *userServiceGetRolesClient) Recv() (*Role, error) {
	m := new(Role)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) SetStatus(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_SetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByEmail(ctx context.Context, in *GetByEmailReq, opts ...grpc.CallOption) (*GetByEmailRes, error) {
	out := new(GetByEmailRes)
	err := c.cc.Invoke(ctx, UserService_GetByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SearchforMembers(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (UserService_SearchforMembersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], UserService_SearchforMembers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceSearchforMembersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_SearchforMembersClient interface {
	Recv() (*SearchRes, error)
	grpc.ClientStream
}

type userServiceSearchforMembersClient struct {
	grpc.ClientStream
}

func (x *userServiceSearchforMembersClient) Recv() (*SearchRes, error) {
	m := new(SearchRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) AddRoles(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_AddRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsReq, opts ...grpc.CallOption) (*GetUserDetailsRes, error) {
	out := new(GetUserDetailsRes)
	err := c.cc.Invoke(ctx, UserService_GetUserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetStreamofUserDetails(ctx context.Context, opts ...grpc.CallOption) (UserService_GetStreamofUserDetailsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[2], UserService_GetStreamofUserDetails_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetStreamofUserDetailsClient{stream}
	return x, nil
}

type UserService_GetStreamofUserDetailsClient interface {
	Send(*GetUserDetailsReq) error
	Recv() (*GetUserDetailsRes, error)
	grpc.ClientStream
}

type userServiceGetStreamofUserDetailsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetStreamofUserDetailsClient) Send(m *GetUserDetailsReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetStreamofUserDetailsClient) Recv() (*GetUserDetailsRes, error) {
	m := new(GetUserDetailsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetStreamofRoles(ctx context.Context, opts ...grpc.CallOption) (UserService_GetStreamofRolesClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[3], UserService_GetStreamofRoles_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetStreamofRolesClient{stream}
	return x, nil
}

type UserService_GetStreamofRolesClient interface {
	Send(*GetStreamofRolesReq) error
	CloseAndRecv() (*GetStreamofRolesRes, error)
	grpc.ClientStream
}

type userServiceGetStreamofRolesClient struct {
	grpc.ClientStream
}

func (x *userServiceGetStreamofRolesClient) Send(m *GetStreamofRolesReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetStreamofRolesClient) CloseAndRecv() (*GetStreamofRolesRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetStreamofRolesRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleRes, error) {
	out := new(GetRoleRes)
	err := c.cc.Invoke(ctx, UserService_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditStatus(ctx context.Context, in *EditStatusReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_EditStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserDetails(ctx context.Context, in *UpdateUserDetailsReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateUserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DropUserAccount(ctx context.Context, in *DropUserAccountReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_DropUserAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	SignupUser(context.Context, *SignupUserRequest) (*UserResponce, error)
	GetRoles(*empty.Empty, UserService_GetRolesServer) error
	SetStatus(context.Context, *StatusReq) (*empty.Empty, error)
	GetByEmail(context.Context, *GetByEmailReq) (*GetByEmailRes, error)
	SearchforMembers(*SearchReq, UserService_SearchforMembersServer) error
	AddRoles(context.Context, *AddRoleReq) (*empty.Empty, error)
	GetUserDetails(context.Context, *GetUserDetailsReq) (*GetUserDetailsRes, error)
	GetStreamofUserDetails(UserService_GetStreamofUserDetailsServer) error
	GetStreamofRoles(UserService_GetStreamofRolesServer) error
	GetRole(context.Context, *GetRoleReq) (*GetRoleRes, error)
	EditStatus(context.Context, *EditStatusReq) (*empty.Empty, error)
	UpdateUserDetails(context.Context, *UpdateUserDetailsReq) (*empty.Empty, error)
	DropUserAccount(context.Context, *DropUserAccountReq) (*empty.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) SignupUser(context.Context, *SignupUserRequest) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupUser not implemented")
}
func (UnimplementedUserServiceServer) GetRoles(*empty.Empty, UserService_GetRolesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedUserServiceServer) SetStatus(context.Context, *StatusReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedUserServiceServer) GetByEmail(context.Context, *GetByEmailReq) (*GetByEmailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedUserServiceServer) SearchforMembers(*SearchReq, UserService_SearchforMembersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchforMembers not implemented")
}
func (UnimplementedUserServiceServer) AddRoles(context.Context, *AddRoleReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoles not implemented")
}
func (UnimplementedUserServiceServer) GetUserDetails(context.Context, *GetUserDetailsReq) (*GetUserDetailsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedUserServiceServer) GetStreamofUserDetails(UserService_GetStreamofUserDetailsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamofUserDetails not implemented")
}
func (UnimplementedUserServiceServer) GetStreamofRoles(UserService_GetStreamofRolesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamofRoles not implemented")
}
func (UnimplementedUserServiceServer) GetRole(context.Context, *GetRoleReq) (*GetRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedUserServiceServer) EditStatus(context.Context, *EditStatusReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditStatus not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserDetails(context.Context, *UpdateUserDetailsReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDetails not implemented")
}
func (UnimplementedUserServiceServer) DropUserAccount(context.Context, *DropUserAccountReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropUserAccount not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_SignupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SignupUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignupUser(ctx, req.(*SignupUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetRoles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetRoles(m, &userServiceGetRolesServer{stream})
}

type UserService_GetRolesServer interface {
	Send(*Role) error
	grpc.ServerStream
}

type userServiceGetRolesServer struct {
	grpc.ServerStream
}

func (x *userServiceGetRolesServer) Send(m *Role) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetStatus(ctx, req.(*StatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByEmail(ctx, req.(*GetByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SearchforMembers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).SearchforMembers(m, &userServiceSearchforMembersServer{stream})
}

type UserService_SearchforMembersServer interface {
	Send(*SearchRes) error
	grpc.ServerStream
}

type userServiceSearchforMembersServer struct {
	grpc.ServerStream
}

func (x *userServiceSearchforMembersServer) Send(m *SearchRes) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_AddRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddRoles(ctx, req.(*AddRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsReq)
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
		return srv.(UserServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetStreamofUserDetails_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetStreamofUserDetails(&userServiceGetStreamofUserDetailsServer{stream})
}

type UserService_GetStreamofUserDetailsServer interface {
	Send(*GetUserDetailsRes) error
	Recv() (*GetUserDetailsReq, error)
	grpc.ServerStream
}

type userServiceGetStreamofUserDetailsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetStreamofUserDetailsServer) Send(m *GetUserDetailsRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetStreamofUserDetailsServer) Recv() (*GetUserDetailsReq, error) {
	m := new(GetUserDetailsReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GetStreamofRoles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetStreamofRoles(&userServiceGetStreamofRolesServer{stream})
}

type UserService_GetStreamofRolesServer interface {
	SendAndClose(*GetStreamofRolesRes) error
	Recv() (*GetStreamofRolesReq, error)
	grpc.ServerStream
}

type userServiceGetStreamofRolesServer struct {
	grpc.ServerStream
}

func (x *userServiceGetStreamofRolesServer) SendAndClose(m *GetStreamofRolesRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetStreamofRolesServer) Recv() (*GetStreamofRolesReq, error) {
	m := new(GetStreamofRolesReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetRole(ctx, req.(*GetRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_EditStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditStatus(ctx, req.(*EditStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDetailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserDetails(ctx, req.(*UpdateUserDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DropUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropUserAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DropUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DropUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DropUserAccount(ctx, req.(*DropUserAccountReq))
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
			MethodName: "SignupUser",
			Handler:    _UserService_SignupUser_Handler,
		},
		{
			MethodName: "SetStatus",
			Handler:    _UserService_SetStatus_Handler,
		},
		{
			MethodName: "GetByEmail",
			Handler:    _UserService_GetByEmail_Handler,
		},
		{
			MethodName: "AddRoles",
			Handler:    _UserService_AddRoles_Handler,
		},
		{
			MethodName: "GetUserDetails",
			Handler:    _UserService_GetUserDetails_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _UserService_GetRole_Handler,
		},
		{
			MethodName: "EditStatus",
			Handler:    _UserService_EditStatus_Handler,
		},
		{
			MethodName: "UpdateUserDetails",
			Handler:    _UserService_UpdateUserDetails_Handler,
		},
		{
			MethodName: "DropUserAccount",
			Handler:    _UserService_DropUserAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRoles",
			Handler:       _UserService_GetRoles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchforMembers",
			Handler:       _UserService_SearchforMembers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStreamofUserDetails",
			Handler:       _UserService_GetStreamofUserDetails_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetStreamofRoles",
			Handler:       _UserService_GetStreamofRoles_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}
