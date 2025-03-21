// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.26.1
// source: schedule/schedule.proto

package schedulev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_schedule_schedule_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{0}
}

type CreateGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TrainerId     int64                  `protobuf:"varint,2,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	mi := &file_schedule_schedule_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGroupRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

type GetGroupsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TrainerId     int64                  `protobuf:"varint,1,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	StudentId     int64                  `protobuf:"varint,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGroupsRequest) Reset() {
	*x = GetGroupsRequest{}
	mi := &file_schedule_schedule_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsRequest) ProtoMessage() {}

func (x *GetGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetGroupsRequest) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupsRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *GetGroupsRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

type GetGroupsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Groups        []*Group               `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGroupsResponse) Reset() {
	*x = GetGroupsResponse{}
	mi := &file_schedule_schedule_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsResponse) ProtoMessage() {}

func (x *GetGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetGroupsResponse) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupsResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type AddToGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     int64                  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Link          string                 `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddToGroupRequest) Reset() {
	*x = AddToGroupRequest{}
	mi := &file_schedule_schedule_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToGroupRequest) ProtoMessage() {}

func (x *AddToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddToGroupRequest) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *AddToGroupRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *AddToGroupRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type CreateScheduleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       int64                  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	StudentId     int64                  `protobuf:"varint,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId     int64                  `protobuf:"varint,3,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateScheduleRequest) Reset() {
	*x = CreateScheduleRequest{}
	mi := &file_schedule_schedule_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleRequest) ProtoMessage() {}

func (x *CreateScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleRequest.ProtoReflect.Descriptor instead.
func (*CreateScheduleRequest) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{5}
}

func (x *CreateScheduleRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *CreateScheduleRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *CreateScheduleRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *CreateScheduleRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type CreateScheduleForGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       int64                  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	TrainerId     int64                  `protobuf:"varint,3,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateScheduleForGroupRequest) Reset() {
	*x = CreateScheduleForGroupRequest{}
	mi := &file_schedule_schedule_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScheduleForGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleForGroupRequest) ProtoMessage() {}

func (x *CreateScheduleForGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleForGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateScheduleForGroupRequest) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{6}
}

func (x *CreateScheduleForGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *CreateScheduleForGroupRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *CreateScheduleForGroupRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetSchedulesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TrainerId     int64                  `protobuf:"varint,1,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	StudentId     int64                  `protobuf:"varint,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	GroupId       int64                  `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSchedulesRequest) Reset() {
	*x = GetSchedulesRequest{}
	mi := &file_schedule_schedule_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchedulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchedulesRequest) ProtoMessage() {}

func (x *GetSchedulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchedulesRequest.ProtoReflect.Descriptor instead.
func (*GetSchedulesRequest) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{7}
}

func (x *GetSchedulesRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *GetSchedulesRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *GetSchedulesRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GetSchedulesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schedules     []*Schedule            `protobuf:"bytes,1,rep,name=schedules,proto3" json:"schedules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSchedulesResponse) Reset() {
	*x = GetSchedulesResponse{}
	mi := &file_schedule_schedule_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchedulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchedulesResponse) ProtoMessage() {}

func (x *GetSchedulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchedulesResponse.ProtoReflect.Descriptor instead.
func (*GetSchedulesResponse) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{8}
}

func (x *GetSchedulesResponse) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TrainerId     int64                  `protobuf:"varint,3,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	Students      []int64                `protobuf:"varint,4,rep,packed,name=students,proto3" json:"students,omitempty"`
	Link          string                 `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_schedule_schedule_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{9}
}

func (x *Group) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *Group) GetStudents() []int64 {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *Group) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type Schedule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupName     string                 `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	GroupId       int64                  `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	StudentId     int64                  `protobuf:"varint,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId     int64                  `protobuf:"varint,4,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	mi := &file_schedule_schedule_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_schedule_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_schedule_schedule_proto_rawDescGZIP(), []int{10}
}

func (x *Schedule) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *Schedule) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Schedule) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *Schedule) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *Schedule) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_schedule_schedule_proto protoreflect.FileDescriptor

var file_schedule_schedule_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x47, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xa0,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x89, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x6e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x48, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x22, 0xb2, 0x01, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0xb0, 0x03, 0x0a, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x27, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x6f, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_schedule_schedule_proto_rawDescOnce sync.Once
	file_schedule_schedule_proto_rawDescData []byte
)

func file_schedule_schedule_proto_rawDescGZIP() []byte {
	file_schedule_schedule_proto_rawDescOnce.Do(func() {
		file_schedule_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_schedule_schedule_proto_rawDesc), len(file_schedule_schedule_proto_rawDesc)))
	})
	return file_schedule_schedule_proto_rawDescData
}

var file_schedule_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_schedule_schedule_proto_goTypes = []any{
	(*Empty)(nil),                         // 0: schedule.Empty
	(*CreateGroupRequest)(nil),            // 1: schedule.CreateGroupRequest
	(*GetGroupsRequest)(nil),              // 2: schedule.GetGroupsRequest
	(*GetGroupsResponse)(nil),             // 3: schedule.GetGroupsResponse
	(*AddToGroupRequest)(nil),             // 4: schedule.AddToGroupRequest
	(*CreateScheduleRequest)(nil),         // 5: schedule.CreateScheduleRequest
	(*CreateScheduleForGroupRequest)(nil), // 6: schedule.CreateScheduleForGroupRequest
	(*GetSchedulesRequest)(nil),           // 7: schedule.GetSchedulesRequest
	(*GetSchedulesResponse)(nil),          // 8: schedule.GetSchedulesResponse
	(*Group)(nil),                         // 9: schedule.Group
	(*Schedule)(nil),                      // 10: schedule.schedule
	(*timestamppb.Timestamp)(nil),         // 11: google.protobuf.Timestamp
}
var file_schedule_schedule_proto_depIdxs = []int32{
	9,  // 0: schedule.GetGroupsResponse.groups:type_name -> schedule.Group
	11, // 1: schedule.CreateScheduleRequest.date:type_name -> google.protobuf.Timestamp
	11, // 2: schedule.CreateScheduleForGroupRequest.date:type_name -> google.protobuf.Timestamp
	10, // 3: schedule.GetSchedulesResponse.schedules:type_name -> schedule.schedule
	11, // 4: schedule.schedule.date:type_name -> google.protobuf.Timestamp
	1,  // 5: schedule.Schedule.CreateGroup:input_type -> schedule.CreateGroupRequest
	2,  // 6: schedule.Schedule.GetGroups:input_type -> schedule.GetGroupsRequest
	4,  // 7: schedule.Schedule.AddToGroup:input_type -> schedule.AddToGroupRequest
	5,  // 8: schedule.Schedule.CreateSchedule:input_type -> schedule.CreateScheduleRequest
	6,  // 9: schedule.Schedule.CreateScheduleForGroup:input_type -> schedule.CreateScheduleForGroupRequest
	7,  // 10: schedule.Schedule.GetSchedule:input_type -> schedule.GetSchedulesRequest
	0,  // 11: schedule.Schedule.CreateGroup:output_type -> schedule.Empty
	3,  // 12: schedule.Schedule.GetGroups:output_type -> schedule.GetGroupsResponse
	0,  // 13: schedule.Schedule.AddToGroup:output_type -> schedule.Empty
	0,  // 14: schedule.Schedule.CreateSchedule:output_type -> schedule.Empty
	0,  // 15: schedule.Schedule.CreateScheduleForGroup:output_type -> schedule.Empty
	8,  // 16: schedule.Schedule.GetSchedule:output_type -> schedule.GetSchedulesResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_schedule_schedule_proto_init() }
func file_schedule_schedule_proto_init() {
	if File_schedule_schedule_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_schedule_schedule_proto_rawDesc), len(file_schedule_schedule_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schedule_schedule_proto_goTypes,
		DependencyIndexes: file_schedule_schedule_proto_depIdxs,
		MessageInfos:      file_schedule_schedule_proto_msgTypes,
	}.Build()
	File_schedule_schedule_proto = out.File
	file_schedule_schedule_proto_goTypes = nil
	file_schedule_schedule_proto_depIdxs = nil
}
