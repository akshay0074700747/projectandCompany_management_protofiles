// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: project.proto

package projectpb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description     string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	ProjectUsername string `protobuf:"bytes,3,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
	Aim             string `protobuf:"bytes,4,opt,name=Aim,proto3" json:"Aim,omitempty"`
	IsCompanybased  bool   `protobuf:"varint,5,opt,name=IsCompanybased,proto3" json:"IsCompanybased,omitempty"`
	ComapanyId      string `protobuf:"bytes,6,opt,name=ComapanyId,proto3" json:"ComapanyId,omitempty"`
	OwnerID         string `protobuf:"bytes,7,opt,name=OwnerID,proto3" json:"OwnerID,omitempty"`
	IsPublic        bool   `protobuf:"varint,8,opt,name=IsPublic,proto3" json:"IsPublic,omitempty"`
}

func (x *CreateProjectReq) Reset() {
	*x = CreateProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectReq) ProtoMessage() {}

func (x *CreateProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectReq.ProtoReflect.Descriptor instead.
func (*CreateProjectReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProjectReq) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

func (x *CreateProjectReq) GetAim() string {
	if x != nil {
		return x.Aim
	}
	return ""
}

func (x *CreateProjectReq) GetIsCompanybased() bool {
	if x != nil {
		return x.IsCompanybased
	}
	return false
}

func (x *CreateProjectReq) GetComapanyId() string {
	if x != nil {
		return x.ComapanyId
	}
	return ""
}

func (x *CreateProjectReq) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *CreateProjectReq) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type CreateProjectRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID       string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description     string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ProjectUsername string `protobuf:"bytes,4,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
	Aim             string `protobuf:"bytes,5,opt,name=Aim,proto3" json:"Aim,omitempty"`
	IsCompanybased  bool   `protobuf:"varint,6,opt,name=IsCompanybased,proto3" json:"IsCompanybased,omitempty"`
	ComapanyID      string `protobuf:"bytes,7,opt,name=ComapanyID,proto3" json:"ComapanyID,omitempty"`
}

func (x *CreateProjectRes) Reset() {
	*x = CreateProjectRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRes) ProtoMessage() {}

func (x *CreateProjectRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRes.ProtoReflect.Descriptor instead.
func (*CreateProjectRes) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectRes) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *CreateProjectRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProjectRes) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

func (x *CreateProjectRes) GetAim() string {
	if x != nil {
		return x.Aim
	}
	return ""
}

func (x *CreateProjectRes) GetIsCompanybased() bool {
	if x != nil {
		return x.IsCompanybased
	}
	return false
}

func (x *CreateProjectRes) GetComapanyID() string {
	if x != nil {
		return x.ComapanyID
	}
	return ""
}

type AddMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	PermissionID uint32 `protobuf:"varint,2,opt,name=PermissionID,proto3" json:"PermissionID,omitempty"`
	RoleID       uint32 `protobuf:"varint,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	ProjectID    string `protobuf:"bytes,4,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	RecruiterID  string `protobuf:"bytes,5,opt,name=RecruiterID,proto3" json:"RecruiterID,omitempty"`
}

func (x *AddMemberReq) Reset() {
	*x = AddMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMemberReq) ProtoMessage() {}

func (x *AddMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMemberReq.ProtoReflect.Descriptor instead.
func (*AddMemberReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

func (x *AddMemberReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddMemberReq) GetPermissionID() uint32 {
	if x != nil {
		return x.PermissionID
	}
	return 0
}

func (x *AddMemberReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *AddMemberReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *AddMemberReq) GetRecruiterID() string {
	if x != nil {
		return x.RecruiterID
	}
	return ""
}

type ProjectInvitesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberID string `protobuf:"bytes,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
}

func (x *ProjectInvitesReq) Reset() {
	*x = ProjectInvitesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInvitesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInvitesReq) ProtoMessage() {}

func (x *ProjectInvitesReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInvitesReq.ProtoReflect.Descriptor instead.
func (*ProjectInvitesReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectInvitesReq) GetMemberID() string {
	if x != nil {
		return x.MemberID
	}
	return ""
}

type ProjectInvitesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Members   uint32 `protobuf:"varint,2,opt,name=Members,proto3" json:"Members,omitempty"`
}

func (x *ProjectInvitesRes) Reset() {
	*x = ProjectInvitesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInvitesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInvitesRes) ProtoMessage() {}

func (x *ProjectInvitesRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInvitesRes.ProtoReflect.Descriptor instead.
func (*ProjectInvitesRes) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectInvitesRes) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *ProjectInvitesRes) GetMembers() uint32 {
	if x != nil {
		return x.Members
	}
	return 0
}

type AcceptProjectInviteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Accept    bool   `protobuf:"varint,2,opt,name=Accept,proto3" json:"Accept,omitempty"`
	UserID    string `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *AcceptProjectInviteReq) Reset() {
	*x = AcceptProjectInviteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptProjectInviteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptProjectInviteReq) ProtoMessage() {}

func (x *AcceptProjectInviteReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptProjectInviteReq.ProtoReflect.Descriptor instead.
func (*AcceptProjectInviteReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{5}
}

func (x *AcceptProjectInviteReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *AcceptProjectInviteReq) GetAccept() bool {
	if x != nil {
		return x.Accept
	}
	return false
}

func (x *AcceptProjectInviteReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID       string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	ProjectUsername string `protobuf:"bytes,2,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
	MemberID        string `protobuf:"bytes,3,opt,name=MemberID,proto3" json:"MemberID,omitempty"`
}

func (x *GetProjectReq) Reset() {
	*x = GetProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectReq) ProtoMessage() {}

func (x *GetProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectReq.ProtoReflect.Descriptor instead.
func (*GetProjectReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{6}
}

func (x *GetProjectReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *GetProjectReq) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

func (x *GetProjectReq) GetMemberID() string {
	if x != nil {
		return x.MemberID
	}
	return ""
}

type GetProjectMembersRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Role       string `protobuf:"bytes,3,opt,name=Role,proto3" json:"Role,omitempty"`
	Permission string `protobuf:"bytes,4,opt,name=Permission,proto3" json:"Permission,omitempty"`
	Reputation uint32 `protobuf:"varint,5,opt,name=Reputation,proto3" json:"Reputation,omitempty"`
}

func (x *GetProjectMembersRes) Reset() {
	*x = GetProjectMembersRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectMembersRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectMembersRes) ProtoMessage() {}

func (x *GetProjectMembersRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectMembersRes.ProtoReflect.Descriptor instead.
func (*GetProjectMembersRes) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{7}
}

func (x *GetProjectMembersRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetProjectMembersRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProjectMembersRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetProjectMembersRes) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *GetProjectMembersRes) GetReputation() uint32 {
	if x != nil {
		return x.Reputation
	}
	return 0
}

type GetProjectDetailesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID         string                  `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	ProjectUsername   string                  `protobuf:"bytes,2,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
	Aim               string                  `protobuf:"bytes,3,opt,name=Aim,proto3" json:"Aim,omitempty"`
	Members           uint32                  `protobuf:"varint,4,opt,name=Members,proto3" json:"Members,omitempty"`
	MembersDetails    []*GetProjectMembersRes `protobuf:"bytes,5,rep,name=MembersDetails,proto3" json:"MembersDetails,omitempty"`
	MembersTerminated uint32                  `protobuf:"varint,6,opt,name=MembersTerminated,proto3" json:"MembersTerminated,omitempty"`
	IssuesRaised      uint32                  `protobuf:"varint,7,opt,name=IssuesRaised,proto3" json:"IssuesRaised,omitempty"`
	WorkflowsRunning  uint32                  `protobuf:"varint,8,opt,name=WorkflowsRunning,proto3" json:"WorkflowsRunning,omitempty"`
	IssuesPending     uint32                  `protobuf:"varint,9,opt,name=IssuesPending,proto3" json:"IssuesPending,omitempty"`
}

func (x *GetProjectDetailesRes) Reset() {
	*x = GetProjectDetailesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectDetailesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectDetailesRes) ProtoMessage() {}

func (x *GetProjectDetailesRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectDetailesRes.ProtoReflect.Descriptor instead.
func (*GetProjectDetailesRes) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{8}
}

func (x *GetProjectDetailesRes) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *GetProjectDetailesRes) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

func (x *GetProjectDetailesRes) GetAim() string {
	if x != nil {
		return x.Aim
	}
	return ""
}

func (x *GetProjectDetailesRes) GetMembers() uint32 {
	if x != nil {
		return x.Members
	}
	return 0
}

func (x *GetProjectDetailesRes) GetMembersDetails() []*GetProjectMembersRes {
	if x != nil {
		return x.MembersDetails
	}
	return nil
}

func (x *GetProjectDetailesRes) GetMembersTerminated() uint32 {
	if x != nil {
		return x.MembersTerminated
	}
	return 0
}

func (x *GetProjectDetailesRes) GetIssuesRaised() uint32 {
	if x != nil {
		return x.IssuesRaised
	}
	return 0
}

func (x *GetProjectDetailesRes) GetWorkflowsRunning() uint32 {
	if x != nil {
		return x.WorkflowsRunning
	}
	return 0
}

func (x *GetProjectDetailesRes) GetIssuesPending() uint32 {
	if x != nil {
		return x.IssuesPending
	}
	return 0
}

type LogintoProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectUsername string `protobuf:"bytes,1,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
}

func (x *LogintoProjectReq) Reset() {
	*x = LogintoProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogintoProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogintoProjectReq) ProtoMessage() {}

func (x *LogintoProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogintoProjectReq.ProtoReflect.Descriptor instead.
func (*LogintoProjectReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{9}
}

func (x *LogintoProjectReq) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x69,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x69, 0x6d, 0x12, 0x26, 0x0a, 0x0e,
	0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x61, 0x73, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62,
	0x61, 0x73, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x49, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0xea, 0x01, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x41, 0x69, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x69, 0x6d, 0x12,
	0x26, 0x0a, 0x0e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x61, 0x73, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x62, 0x61, 0x73, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6d,
	0x61, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0xa0, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22,
	0x0a, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x72,
	0x75, 0x69, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52,
	0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x73, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x44, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf6, 0x02, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x41, 0x69, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x69, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x52, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x61, 0x69, 0x73, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x61, 0x69,
	0x73, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x24, 0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x6f,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x92, 0x04, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3b,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x6f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData = file_project_proto_rawDesc
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_proto_rawDescData)
	})
	return file_project_proto_rawDescData
}

var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_project_proto_goTypes = []interface{}{
	(*CreateProjectReq)(nil),       // 0: project.CreateProjectReq
	(*CreateProjectRes)(nil),       // 1: project.CreateProjectRes
	(*AddMemberReq)(nil),           // 2: project.AddMemberReq
	(*ProjectInvitesReq)(nil),      // 3: project.ProjectInvitesReq
	(*ProjectInvitesRes)(nil),      // 4: project.ProjectInvitesRes
	(*AcceptProjectInviteReq)(nil), // 5: project.AcceptProjectInviteReq
	(*GetProjectReq)(nil),          // 6: project.GetProjectReq
	(*GetProjectMembersRes)(nil),   // 7: project.GetProjectMembersRes
	(*GetProjectDetailesRes)(nil),  // 8: project.GetProjectDetailesRes
	(*LogintoProjectReq)(nil),      // 9: project.LogintoProjectReq
	(*empty.Empty)(nil),            // 10: google.protobuf.Empty
}
var file_project_proto_depIdxs = []int32{
	7,  // 0: project.GetProjectDetailesRes.MembersDetails:type_name -> project.GetProjectMembersRes
	0,  // 1: project.ProjectService.CreateProject:input_type -> project.CreateProjectReq
	2,  // 2: project.ProjectService.AddMembers:input_type -> project.AddMemberReq
	3,  // 3: project.ProjectService.ProjectInvites:input_type -> project.ProjectInvitesReq
	5,  // 4: project.ProjectService.AcceptProjectInvite:input_type -> project.AcceptProjectInviteReq
	6,  // 5: project.ProjectService.GetProjectDetailes:input_type -> project.GetProjectReq
	6,  // 6: project.ProjectService.GetProjectMembers:input_type -> project.GetProjectReq
	9,  // 7: project.ProjectService.LogintoProject:input_type -> project.LogintoProjectReq
	1,  // 8: project.ProjectService.CreateProject:output_type -> project.CreateProjectRes
	10, // 9: project.ProjectService.AddMembers:output_type -> google.protobuf.Empty
	4,  // 10: project.ProjectService.ProjectInvites:output_type -> project.ProjectInvitesRes
	10, // 11: project.ProjectService.AcceptProjectInvite:output_type -> google.protobuf.Empty
	8,  // 12: project.ProjectService.GetProjectDetailes:output_type -> project.GetProjectDetailesRes
	7,  // 13: project.ProjectService.GetProjectMembers:output_type -> project.GetProjectMembersRes
	10, // 14: project.ProjectService.LogintoProject:output_type -> google.protobuf.Empty
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMemberReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitesReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitesRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptProjectInviteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectMembersRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectDetailesRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogintoProjectReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_rawDesc = nil
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
