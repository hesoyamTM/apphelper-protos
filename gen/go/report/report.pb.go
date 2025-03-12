// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.26.1
// source: report/report.proto

package reportv1

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
	mi := &file_report_report_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[0]
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
	return file_report_report_proto_rawDescGZIP(), []int{0}
}

type GetReportsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     int64                  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId     int64                  `protobuf:"varint,2,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReportsRequest) Reset() {
	*x = GetReportsRequest{}
	mi := &file_report_report_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReportsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportsRequest) ProtoMessage() {}

func (x *GetReportsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportsRequest.ProtoReflect.Descriptor instead.
func (*GetReportsRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{1}
}

func (x *GetReportsRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *GetReportsRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

type GetReportsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reports       []*Report              `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReportsResponse) Reset() {
	*x = GetReportsResponse{}
	mi := &file_report_report_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReportsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportsResponse) ProtoMessage() {}

func (x *GetReportsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportsResponse.ProtoReflect.Descriptor instead.
func (*GetReportsResponse) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetReportsResponse) GetReports() []*Report {
	if x != nil {
		return x.Reports
	}
	return nil
}

type CreateReportRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	StudentId int64                  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId int64                  `protobuf:"varint,2,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	// repeated image images = 3;
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReportRequest) Reset() {
	*x = CreateReportRequest{}
	mi := &file_report_report_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportRequest) ProtoMessage() {}

func (x *CreateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReportRequest.ProtoReflect.Descriptor instead.
func (*CreateReportRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReportRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *CreateReportRequest) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *CreateReportRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Report struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     int64                  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId     int64                  `protobuf:"varint,2,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	Images        []*Image               `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Report) Reset() {
	*x = Report{}
	mi := &file_report_report_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{4}
}

func (x *Report) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *Report) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *Report) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Report) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Report) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Width         int32                  `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32                  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ImageData     []byte                 `protobuf:"bytes,3,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_report_report_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{5}
}

func (x *Image) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Image) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Image) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

var File_report_report_proto protoreflect.FileDescriptor

var file_report_report_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x51, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x75, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x54, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0x89, 0x01, 0x0a, 0x06, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x14, 0x5a, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_report_report_proto_rawDescOnce sync.Once
	file_report_report_proto_rawDescData []byte
)

func file_report_report_proto_rawDescGZIP() []byte {
	file_report_report_proto_rawDescOnce.Do(func() {
		file_report_report_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_report_report_proto_rawDesc), len(file_report_report_proto_rawDesc)))
	})
	return file_report_report_proto_rawDescData
}

var file_report_report_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_report_report_proto_goTypes = []any{
	(*Empty)(nil),                 // 0: report.Empty
	(*GetReportsRequest)(nil),     // 1: report.GetReportsRequest
	(*GetReportsResponse)(nil),    // 2: report.GetReportsResponse
	(*CreateReportRequest)(nil),   // 3: report.CreateReportRequest
	(*Report)(nil),                // 4: report.report
	(*Image)(nil),                 // 5: report.image
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_report_report_proto_depIdxs = []int32{
	4, // 0: report.GetReportsResponse.reports:type_name -> report.report
	5, // 1: report.report.images:type_name -> report.image
	6, // 2: report.report.date:type_name -> google.protobuf.Timestamp
	1, // 3: report.Report.GetReports:input_type -> report.GetReportsRequest
	3, // 4: report.Report.CreateReport:input_type -> report.CreateReportRequest
	2, // 5: report.Report.GetReports:output_type -> report.GetReportsResponse
	0, // 6: report.Report.CreateReport:output_type -> report.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_report_report_proto_init() }
func file_report_report_proto_init() {
	if File_report_report_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_report_report_proto_rawDesc), len(file_report_report_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_report_proto_goTypes,
		DependencyIndexes: file_report_report_proto_depIdxs,
		MessageInfos:      file_report_report_proto_msgTypes,
	}.Build()
	File_report_report_proto = out.File
	file_report_report_proto_goTypes = nil
	file_report_report_proto_depIdxs = nil
}
