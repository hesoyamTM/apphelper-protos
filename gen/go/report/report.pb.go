// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: report/report.proto

package reportv1

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

type SetPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SetPublicKeyRequest) Reset() {
	*x = SetPublicKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPublicKeyRequest) ProtoMessage() {}

func (x *SetPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*SetPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{1}
}

func (x *SetPublicKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetReportRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId   int64    `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId   int64    `protobuf:"varint,2,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	Images      []*Image `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetReportResponse) Reset() {
	*x = GetReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportResponse) ProtoMessage() {}

func (x *GetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportResponse.ProtoReflect.Descriptor instead.
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{3}
}

func (x *GetReportResponse) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *GetReportResponse) GetTrainerId() int64 {
	if x != nil {
		return x.TrainerId
	}
	return 0
}

func (x *GetReportResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *GetReportResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId int64 `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TrainerId int64 `protobuf:"varint,2,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
	// repeated image images = 3;
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateReportRequest) Reset() {
	*x = CreateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportRequest) ProtoMessage() {}

func (x *CreateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_report_report_proto_rawDescGZIP(), []int{4}
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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width     int32  `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height    int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ImageData []byte `protobuf:"bytes,3,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_report_report_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x54, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0xc3, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x14, 0x5a, 0x12,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_report_proto_rawDescOnce sync.Once
	file_report_report_proto_rawDescData = file_report_report_proto_rawDesc
)

func file_report_report_proto_rawDescGZIP() []byte {
	file_report_report_proto_rawDescOnce.Do(func() {
		file_report_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_report_proto_rawDescData)
	})
	return file_report_report_proto_rawDescData
}

var file_report_report_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_report_report_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: report.Empty
	(*SetPublicKeyRequest)(nil), // 1: report.SetPublicKeyRequest
	(*GetReportRequest)(nil),    // 2: report.GetReportRequest
	(*GetReportResponse)(nil),   // 3: report.GetReportResponse
	(*CreateReportRequest)(nil), // 4: report.CreateReportRequest
	(*Image)(nil),               // 5: report.image
}
var file_report_report_proto_depIdxs = []int32{
	5, // 0: report.GetReportResponse.images:type_name -> report.image
	1, // 1: report.Results.SetPublicKey:input_type -> report.SetPublicKeyRequest
	2, // 2: report.Results.GetReport:input_type -> report.GetReportRequest
	4, // 3: report.Results.CreateReport:input_type -> report.CreateReportRequest
	0, // 4: report.Results.SetPublicKey:output_type -> report.Empty
	3, // 5: report.Results.GetReport:output_type -> report.GetReportResponse
	0, // 6: report.Results.CreateReport:output_type -> report.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_report_report_proto_init() }
func file_report_report_proto_init() {
	if File_report_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_report_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPublicKeyRequest); i {
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
		file_report_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportRequest); i {
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
		file_report_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportResponse); i {
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
		file_report_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReportRequest); i {
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
		file_report_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
			RawDescriptor: file_report_report_proto_rawDesc,
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
	file_report_report_proto_rawDesc = nil
	file_report_report_proto_goTypes = nil
	file_report_report_proto_depIdxs = nil
}
