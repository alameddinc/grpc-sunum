// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/fish.proto

package protoGo

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

type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X        uint64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y        uint64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_proto_fish_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMessage) GetX() uint64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *RequestMessage) GetY() uint64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *RequestMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type RequestRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RequestRegister) Reset() {
	*x = RequestRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRegister) ProtoMessage() {}

func (x *RequestRegister) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRegister.ProtoReflect.Descriptor instead.
func (*RequestRegister) Descriptor() ([]byte, []int) {
	return file_proto_fish_proto_rawDescGZIP(), []int{1}
}

func (x *RequestRegister) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ResponseRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseRegister) Reset() {
	*x = ResponseRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRegister) ProtoMessage() {}

func (x *ResponseRegister) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRegister.ProtoReflect.Descriptor instead.
func (*ResponseRegister) Descriptor() ([]byte, []int) {
	return file_proto_fish_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseRegister) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseRegister) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FishCount uint64 `protobuf:"varint,2,opt,name=fishCount,proto3" json:"fishCount,omitempty"`
	Status    bool   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_proto_fish_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ResponseMessage) GetFishCount() uint64 {
	if x != nil {
		return x.FishCount
	}
	return 0
}

func (x *ResponseMessage) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type RequestHighScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RequestHighScore) Reset() {
	*x = RequestHighScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHighScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHighScore) ProtoMessage() {}

func (x *RequestHighScore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fish_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHighScore.ProtoReflect.Descriptor instead.
func (*RequestHighScore) Descriptor() ([]byte, []int) {
	return file_proto_fish_proto_rawDescGZIP(), []int{4}
}

func (x *RequestHighScore) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ResponseHighScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users    map[string]int64 `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	YourRank int64            `protobuf:"varint,2,opt,name=yourRank,proto3" json:"yourRank,omitempty"`
	Status   bool             `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseHighScore) Reset() {
	*x = ResponseHighScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fish_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseHighScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHighScore) ProtoMessage() {}

func (x *ResponseHighScore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fish_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHighScore.ProtoReflect.Descriptor instead.
func (*ResponseHighScore) Descriptor() ([]byte, []int) {
	return file_proto_fish_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseHighScore) GetUsers() map[string]int64 {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ResponseHighScore) GetYourRank() int64 {
	if x != nil {
		return x.YourRank
	}
	return 0
}

func (x *ResponseHighScore) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_proto_fish_proto protoreflect.FileDescriptor

var file_proto_fish_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x73, 0x68, 0x22, 0x48, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x01, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x44, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x63, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x73, 0x68, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x69, 0x73, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x10,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbb, 0x01, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x79, 0x6f, 0x75, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x79, 0x6f, 0x75, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x1a, 0x38, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xcd, 0x01, 0x0a, 0x0b, 0x46,
	0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x66, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e,
	0x66, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0a, 0x54, 0x72, 0x79,
	0x54, 0x6f, 0x43, 0x61, 0x74, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x66, 0x69, 0x73, 0x68, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e,
	0x66, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x09, 0x48, 0x69,
	0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x66, 0x69, 0x73, 0x68, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x17, 0x2e, 0x66, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f,
	0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fish_proto_rawDescOnce sync.Once
	file_proto_fish_proto_rawDescData = file_proto_fish_proto_rawDesc
)

func file_proto_fish_proto_rawDescGZIP() []byte {
	file_proto_fish_proto_rawDescOnce.Do(func() {
		file_proto_fish_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fish_proto_rawDescData)
	})
	return file_proto_fish_proto_rawDescData
}

var file_proto_fish_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_fish_proto_goTypes = []interface{}{
	(*RequestMessage)(nil),    // 0: fish.RequestMessage
	(*RequestRegister)(nil),   // 1: fish.RequestRegister
	(*ResponseRegister)(nil),  // 2: fish.ResponseRegister
	(*ResponseMessage)(nil),   // 3: fish.ResponseMessage
	(*RequestHighScore)(nil),  // 4: fish.RequestHighScore
	(*ResponseHighScore)(nil), // 5: fish.ResponseHighScore
	nil,                       // 6: fish.ResponseHighScore.UsersEntry
}
var file_proto_fish_proto_depIdxs = []int32{
	6, // 0: fish.ResponseHighScore.Users:type_name -> fish.ResponseHighScore.UsersEntry
	1, // 1: fish.FishService.Register:input_type -> fish.RequestRegister
	0, // 2: fish.FishService.TryToCatch:input_type -> fish.RequestMessage
	4, // 3: fish.FishService.HighScore:input_type -> fish.RequestHighScore
	2, // 4: fish.FishService.Register:output_type -> fish.ResponseRegister
	3, // 5: fish.FishService.TryToCatch:output_type -> fish.ResponseMessage
	5, // 6: fish.FishService.HighScore:output_type -> fish.ResponseHighScore
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_fish_proto_init() }
func file_proto_fish_proto_init() {
	if File_proto_fish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_proto_fish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRegister); i {
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
		file_proto_fish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRegister); i {
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
		file_proto_fish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
		file_proto_fish_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHighScore); i {
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
		file_proto_fish_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseHighScore); i {
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
			RawDescriptor: file_proto_fish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fish_proto_goTypes,
		DependencyIndexes: file_proto_fish_proto_depIdxs,
		MessageInfos:      file_proto_fish_proto_msgTypes,
	}.Build()
	File_proto_fish_proto = out.File
	file_proto_fish_proto_rawDesc = nil
	file_proto_fish_proto_goTypes = nil
	file_proto_fish_proto_depIdxs = nil
}
