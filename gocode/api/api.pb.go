// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.1
// source: protos/api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Person struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_protos_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PersonDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonDetails) Reset() {
	*x = PersonDetails{}
	mi := &file_protos_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonDetails) ProtoMessage() {}

func (x *PersonDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonDetails.ProtoReflect.Descriptor instead.
func (*PersonDetails) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{1}
}

func (x *PersonDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_protos_api_proto protoreflect.FileDescriptor

var file_protos_api_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x49, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x3f, 0x0a, 0x0d,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protos_api_proto_rawDescOnce sync.Once
	file_protos_api_proto_rawDescData []byte
)

func file_protos_api_proto_rawDescGZIP() []byte {
	file_protos_api_proto_rawDescOnce.Do(func() {
		file_protos_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_api_proto_rawDesc), len(file_protos_api_proto_rawDesc)))
	})
	return file_protos_api_proto_rawDescData
}

var file_protos_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_api_proto_goTypes = []any{
	(*Person)(nil),        // 0: api.Person
	(*PersonDetails)(nil), // 1: api.PersonDetails
}
var file_protos_api_proto_depIdxs = []int32{
	0, // 0: api.PersonService.GetPerson:input_type -> api.Person
	1, // 1: api.PersonService.GetPerson:output_type -> api.PersonDetails
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_api_proto_init() }
func file_protos_api_proto_init() {
	if File_protos_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_api_proto_rawDesc), len(file_protos_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_api_proto_goTypes,
		DependencyIndexes: file_protos_api_proto_depIdxs,
		MessageInfos:      file_protos_api_proto_msgTypes,
	}.Build()
	File_protos_api_proto = out.File
	file_protos_api_proto_goTypes = nil
	file_protos_api_proto_depIdxs = nil
}
