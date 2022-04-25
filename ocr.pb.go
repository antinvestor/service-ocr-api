// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: v1/ocr.proto

package ocr_v1

import (
	common "github.com/antinvestor/apis/common"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OCRFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId     string            `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Language   string            `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Status     common.STATUS     `protobuf:"varint,3,opt,name=status,proto3,enum=apis.STATUS" json:"status,omitempty"`
	Text       string            `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Properties map[string]string `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *OCRFile) Reset() {
	*x = OCRFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ocr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCRFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRFile) ProtoMessage() {}

func (x *OCRFile) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ocr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRFile.ProtoReflect.Descriptor instead.
func (*OCRFile) Descriptor() ([]byte, []int) {
	return file_v1_ocr_proto_rawDescGZIP(), []int{0}
}

func (x *OCRFile) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *OCRFile) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *OCRFile) GetStatus() common.STATUS {
	if x != nil {
		return x.Status
	}
	return common.STATUS_UNKNOWN
}

func (x *OCRFile) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *OCRFile) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

//Request to determine text found in a file
type OcrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId string            `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	LanguageId  string            `protobuf:"bytes,2,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
	Properties  map[string]string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Async       bool              `protobuf:"varint,4,opt,name=async,proto3" json:"async,omitempty"`
	FileId      []string          `protobuf:"bytes,5,rep,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *OcrRequest) Reset() {
	*x = OcrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ocr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcrRequest) ProtoMessage() {}

func (x *OcrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ocr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcrRequest.ProtoReflect.Descriptor instead.
func (*OcrRequest) Descriptor() ([]byte, []int) {
	return file_v1_ocr_proto_rawDescGZIP(), []int{1}
}

func (x *OcrRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *OcrRequest) GetLanguageId() string {
	if x != nil {
		return x.LanguageId
	}
	return ""
}

func (x *OcrRequest) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *OcrRequest) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *OcrRequest) GetFileId() []string {
	if x != nil {
		return x.FileId
	}
	return nil
}

type OcrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId string     `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Result      []*OCRFile `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *OcrResponse) Reset() {
	*x = OcrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ocr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcrResponse) ProtoMessage() {}

func (x *OcrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ocr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcrResponse.ProtoReflect.Descriptor instead.
func (*OcrResponse) Descriptor() ([]byte, []int) {
	return file_v1_ocr_proto_rawDescGZIP(), []int{2}
}

func (x *OcrResponse) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *OcrResponse) GetResult() []*OCRFile {
	if x != nil {
		return x.Result
	}
	return nil
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ocr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ocr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_v1_ocr_proto_rawDescGZIP(), []int{3}
}

func (x *StatusRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

var File_v1_ocr_proto protoreflect.FileDescriptor

var file_v1_ocr_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x6f, 0x63, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x70, 0x69, 0x73, 0x1a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x07, 0x4f, 0x43, 0x52, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x4f, 0x43, 0x52, 0x46, 0x69, 0x6c, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xda, 0x02, 0x0a, 0x0a,
	0x4f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x10, 0x03, 0x18, 0x28, 0x32, 0x10, 0x5b, 0x30, 0x2d,
	0x39, 0x61, 0x2d, 0x7a, 0x5f, 0x2d, 0x5d, 0x7b, 0x33, 0x2c, 0x32, 0x30, 0x7d, 0x52, 0x0b, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0b, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x15, 0xfa, 0x42, 0x12, 0x72, 0x10, 0x10, 0x02, 0x18, 0x03, 0x32, 0x0a, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x7b, 0x32, 0x2c, 0x33, 0x7d, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x4f, 0x63,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x3d, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x21,
	0x92, 0x01, 0x1e, 0x08, 0x01, 0x10, 0x05, 0x22, 0x18, 0x72, 0x16, 0x10, 0x03, 0x18, 0x28, 0x32,
	0x10, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x5f, 0x2d, 0x5d, 0x7b, 0x33, 0x2c, 0x32, 0x30,
	0x7d, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x0b, 0x4f, 0x63, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x4f, 0x43, 0x52, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x10,
	0x03, 0x18, 0x28, 0x32, 0x10, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x5f, 0x2d, 0x5d, 0x7b,
	0x33, 0x2c, 0x32, 0x30, 0x7d, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x32, 0x70, 0x0a, 0x0a, 0x4f, 0x43, 0x52, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6f, 0x63, 0x72, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ocr_proto_rawDescOnce sync.Once
	file_v1_ocr_proto_rawDescData = file_v1_ocr_proto_rawDesc
)

func file_v1_ocr_proto_rawDescGZIP() []byte {
	file_v1_ocr_proto_rawDescOnce.Do(func() {
		file_v1_ocr_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ocr_proto_rawDescData)
	})
	return file_v1_ocr_proto_rawDescData
}

var file_v1_ocr_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_ocr_proto_goTypes = []interface{}{
	(*OCRFile)(nil),       // 0: apis.OCRFile
	(*OcrRequest)(nil),    // 1: apis.OcrRequest
	(*OcrResponse)(nil),   // 2: apis.OcrResponse
	(*StatusRequest)(nil), // 3: apis.StatusRequest
	nil,                   // 4: apis.OCRFile.PropertiesEntry
	nil,                   // 5: apis.OcrRequest.PropertiesEntry
	(common.STATUS)(0),    // 6: apis.STATUS
}
var file_v1_ocr_proto_depIdxs = []int32{
	6, // 0: apis.OCRFile.status:type_name -> apis.STATUS
	4, // 1: apis.OCRFile.properties:type_name -> apis.OCRFile.PropertiesEntry
	5, // 2: apis.OcrRequest.properties:type_name -> apis.OcrRequest.PropertiesEntry
	0, // 3: apis.OcrResponse.result:type_name -> apis.OCRFile
	1, // 4: apis.OCRService.Recognize:input_type -> apis.OcrRequest
	3, // 5: apis.OCRService.Status:input_type -> apis.StatusRequest
	2, // 6: apis.OCRService.Recognize:output_type -> apis.OcrResponse
	2, // 7: apis.OCRService.Status:output_type -> apis.OcrResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_ocr_proto_init() }
func file_v1_ocr_proto_init() {
	if File_v1_ocr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_ocr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCRFile); i {
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
		file_v1_ocr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcrRequest); i {
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
		file_v1_ocr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcrResponse); i {
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
		file_v1_ocr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
			RawDescriptor: file_v1_ocr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_ocr_proto_goTypes,
		DependencyIndexes: file_v1_ocr_proto_depIdxs,
		MessageInfos:      file_v1_ocr_proto_msgTypes,
	}.Build()
	File_v1_ocr_proto = out.File
	file_v1_ocr_proto_rawDesc = nil
	file_v1_ocr_proto_goTypes = nil
	file_v1_ocr_proto_depIdxs = nil
}