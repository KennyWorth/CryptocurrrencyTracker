// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: grpc.proto

package pb

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

type CoinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinList []*CoinInfo `protobuf:"bytes,1,rep,name=coinList,proto3" json:"coinList,omitempty"`
}

func (x *CoinInfo) Reset() {
	*x = CoinInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinInfo) ProtoMessage() {}

func (x *CoinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinInfo.ProtoReflect.Descriptor instead.
func (*CoinInfo) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CoinInfo) GetCoinList() []*CoinInfo {
	if x != nil {
		return x.CoinList
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
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
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

type CoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request []string `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty"`
}

func (x *CoinRequest) Reset() {
	*x = CoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinRequest) ProtoMessage() {}

func (x *CoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinRequest.ProtoReflect.Descriptor instead.
func (*CoinRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *CoinRequest) GetRequest() []string {
	if x != nil {
		return x.Request
	}
	return nil
}

type CoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinList []*CoinIds `protobuf:"bytes,1,rep,name=coinList,proto3" json:"coinList,omitempty"`
}

func (x *CoinResponse) Reset() {
	*x = CoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinResponse) ProtoMessage() {}

func (x *CoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinResponse.ProtoReflect.Descriptor instead.
func (*CoinResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *CoinResponse) GetCoinList() []*CoinIds {
	if x != nil {
		return x.CoinList
	}
	return nil
}

type CoinIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CoinIds) Reset() {
	*x = CoinIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinIds) ProtoMessage() {}

func (x *CoinIds) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinIds.ProtoReflect.Descriptor instead.
func (*CoinIds) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *CoinIds) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoinIds) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CoinIds) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CoinInfo_CoinIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CoinInfo_CoinIds) Reset() {
	*x = CoinInfo_CoinIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinInfo_CoinIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinInfo_CoinIds) ProtoMessage() {}

func (x *CoinInfo_CoinIds) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinInfo_CoinIds.ProtoReflect.Descriptor instead.
func (*CoinInfo_CoinIds) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CoinInfo_CoinIds) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoinInfo_CoinIds) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CoinInfo_CoinIds) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x43, 0x0a, 0x07, 0x43,
	0x6f, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x0b, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x69,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x49, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x12, 0x13, 0x2e, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_goTypes = []interface{}{
	(*CoinInfo)(nil),         // 0: ctrackergrpc.CoinInfo
	(*Empty)(nil),            // 1: ctrackergrpc.Empty
	(*CoinRequest)(nil),      // 2: ctrackergrpc.CoinRequest
	(*CoinResponse)(nil),     // 3: ctrackergrpc.CoinResponse
	(*CoinIds)(nil),          // 4: ctrackergrpc.CoinIds
	(*CoinInfo_CoinIds)(nil), // 5: ctrackergrpc.CoinInfo.CoinIds
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: ctrackergrpc.CoinInfo.coinList:type_name -> ctrackergrpc.CoinInfo
	4, // 1: ctrackergrpc.CoinResponse.coinList:type_name -> ctrackergrpc.CoinIds
	1, // 2: ctrackergrpc.GetAllCoins.Coins:input_type -> ctrackergrpc.Empty
	3, // 3: ctrackergrpc.GetAllCoins.Coins:output_type -> ctrackergrpc.CoinResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinInfo); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinRequest); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinResponse); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinIds); i {
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
		file_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinInfo_CoinIds); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}