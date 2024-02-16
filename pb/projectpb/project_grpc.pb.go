// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: project.proto

package projectpb

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
	ProjectService_CreateProject_FullMethodName       = "/project.ProjectService/CreateProject"
	ProjectService_AddMembers_FullMethodName          = "/project.ProjectService/AddMembers"
	ProjectService_ProjectInvites_FullMethodName      = "/project.ProjectService/ProjectInvites"
	ProjectService_AcceptProjectInvite_FullMethodName = "/project.ProjectService/AcceptProjectInvite"
	ProjectService_GetProjectDetailes_FullMethodName  = "/project.ProjectService/GetProjectDetailes"
	ProjectService_GetProjectMembers_FullMethodName   = "/project.ProjectService/GetProjectMembers"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	CreateProject(ctx context.Context, in *CreateProjectReq, opts ...grpc.CallOption) (*CreateProjectRes, error)
	AddMembers(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*empty.Empty, error)
	ProjectInvites(ctx context.Context, in *ProjectInvitesReq, opts ...grpc.CallOption) (ProjectService_ProjectInvitesClient, error)
	AcceptProjectInvite(ctx context.Context, in *AcceptProjectInviteReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetProjectDetailes(ctx context.Context, in *GetProjectReq, opts ...grpc.CallOption) (*GetProjectDetailesRes, error)
	GetProjectMembers(ctx context.Context, in *GetProjectReq, opts ...grpc.CallOption) (ProjectService_GetProjectMembersClient, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectReq, opts ...grpc.CallOption) (*CreateProjectRes, error) {
	out := new(CreateProjectRes)
	err := c.cc.Invoke(ctx, ProjectService_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) AddMembers(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectService_AddMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ProjectInvites(ctx context.Context, in *ProjectInvitesReq, opts ...grpc.CallOption) (ProjectService_ProjectInvitesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[0], ProjectService_ProjectInvites_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceProjectInvitesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_ProjectInvitesClient interface {
	Recv() (*ProjectInvitesRes, error)
	grpc.ClientStream
}

type projectServiceProjectInvitesClient struct {
	grpc.ClientStream
}

func (x *projectServiceProjectInvitesClient) Recv() (*ProjectInvitesRes, error) {
	m := new(ProjectInvitesRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) AcceptProjectInvite(ctx context.Context, in *AcceptProjectInviteReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectService_AcceptProjectInvite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjectDetailes(ctx context.Context, in *GetProjectReq, opts ...grpc.CallOption) (*GetProjectDetailesRes, error) {
	out := new(GetProjectDetailesRes)
	err := c.cc.Invoke(ctx, ProjectService_GetProjectDetailes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjectMembers(ctx context.Context, in *GetProjectReq, opts ...grpc.CallOption) (ProjectService_GetProjectMembersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[1], ProjectService_GetProjectMembers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetProjectMembersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetProjectMembersClient interface {
	Recv() (*GetProjectMembersRes, error)
	grpc.ClientStream
}

type projectServiceGetProjectMembersClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetProjectMembersClient) Recv() (*GetProjectMembersRes, error) {
	m := new(GetProjectMembersRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	CreateProject(context.Context, *CreateProjectReq) (*CreateProjectRes, error)
	AddMembers(context.Context, *AddMemberReq) (*empty.Empty, error)
	ProjectInvites(*ProjectInvitesReq, ProjectService_ProjectInvitesServer) error
	AcceptProjectInvite(context.Context, *AcceptProjectInviteReq) (*empty.Empty, error)
	GetProjectDetailes(context.Context, *GetProjectReq) (*GetProjectDetailesRes, error)
	GetProjectMembers(*GetProjectReq, ProjectService_GetProjectMembersServer) error
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) CreateProject(context.Context, *CreateProjectReq) (*CreateProjectRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) AddMembers(context.Context, *AddMemberReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMembers not implemented")
}
func (UnimplementedProjectServiceServer) ProjectInvites(*ProjectInvitesReq, ProjectService_ProjectInvitesServer) error {
	return status.Errorf(codes.Unimplemented, "method ProjectInvites not implemented")
}
func (UnimplementedProjectServiceServer) AcceptProjectInvite(context.Context, *AcceptProjectInviteReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptProjectInvite not implemented")
}
func (UnimplementedProjectServiceServer) GetProjectDetailes(context.Context, *GetProjectReq) (*GetProjectDetailesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectDetailes not implemented")
}
func (UnimplementedProjectServiceServer) GetProjectMembers(*GetProjectReq, ProjectService_GetProjectMembersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProjectMembers not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_AddMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_AddMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddMembers(ctx, req.(*AddMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ProjectInvites_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProjectInvitesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).ProjectInvites(m, &projectServiceProjectInvitesServer{stream})
}

type ProjectService_ProjectInvitesServer interface {
	Send(*ProjectInvitesRes) error
	grpc.ServerStream
}

type projectServiceProjectInvitesServer struct {
	grpc.ServerStream
}

func (x *projectServiceProjectInvitesServer) Send(m *ProjectInvitesRes) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_AcceptProjectInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptProjectInviteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AcceptProjectInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_AcceptProjectInvite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AcceptProjectInvite(ctx, req.(*AcceptProjectInviteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjectDetailes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProjectDetailes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProjectDetailes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProjectDetailes(ctx, req.(*GetProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjectMembers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetProjectReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetProjectMembers(m, &projectServiceGetProjectMembersServer{stream})
}

type ProjectService_GetProjectMembersServer interface {
	Send(*GetProjectMembersRes) error
	grpc.ServerStream
}

type projectServiceGetProjectMembersServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetProjectMembersServer) Send(m *GetProjectMembersRes) error {
	return x.ServerStream.SendMsg(m)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "AddMembers",
			Handler:    _ProjectService_AddMembers_Handler,
		},
		{
			MethodName: "AcceptProjectInvite",
			Handler:    _ProjectService_AcceptProjectInvite_Handler,
		},
		{
			MethodName: "GetProjectDetailes",
			Handler:    _ProjectService_GetProjectDetailes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProjectInvites",
			Handler:       _ProjectService_ProjectInvites_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetProjectMembers",
			Handler:       _ProjectService_GetProjectMembers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "project.proto",
}
