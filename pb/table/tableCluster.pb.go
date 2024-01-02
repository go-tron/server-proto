// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: tableCluster.proto

package table_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeVO `protobuf:"bytes,1,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *NodeList) Reset() {
	*x = NodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableCluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_tableCluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeList.ProtoReflect.Descriptor instead.
func (*NodeList) Descriptor() ([]byte, []int) {
	return file_tableCluster_proto_rawDescGZIP(), []int{0}
}

func (x *NodeList) GetNodes() []*NodeVO {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Addr       string `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
	TableCount uint64 `protobuf:"varint,3,opt,name=TableCount,proto3" json:"TableCount,omitempty"`
}

func (x *NodeVO) Reset() {
	*x = NodeVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableCluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeVO) ProtoMessage() {}

func (x *NodeVO) ProtoReflect() protoreflect.Message {
	mi := &file_tableCluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeVO.ProtoReflect.Descriptor instead.
func (*NodeVO) Descriptor() ([]byte, []int) {
	return file_tableCluster_proto_rawDescGZIP(), []int{1}
}

func (x *NodeVO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeVO) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *NodeVO) GetTableCount() uint64 {
	if x != nil {
		return x.TableCount
	}
	return 0
}

type NodeNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *NodeNameReq) Reset() {
	*x = NodeNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableCluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNameReq) ProtoMessage() {}

func (x *NodeNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_tableCluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNameReq.ProtoReflect.Descriptor instead.
func (*NodeNameReq) Descriptor() ([]byte, []int) {
	return file_tableCluster_proto_rawDescGZIP(), []int{2}
}

func (x *NodeNameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tableCluster_proto protoreflect.FileDescriptor

var file_tableCluster_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x56, 0x4f, 0x52, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x06,
	0x4e, 0x6f, 0x64, 0x65, 0x56, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x21,
	0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0xcc, 0x03, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x72,
	0x76, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x72, 0x76, 0x12, 0x15,
	0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x72, 0x76, 0x12,
	0x15, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x72,
	0x76, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x72, 0x76, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x53, 0x72,
	0x76, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x70, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x56, 0x4f, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x72, 0x76,
	0x12, 0x14, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x20, 0x5a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tableCluster_proto_rawDescOnce sync.Once
	file_tableCluster_proto_rawDescData = file_tableCluster_proto_rawDesc
)

func file_tableCluster_proto_rawDescGZIP() []byte {
	file_tableCluster_proto_rawDescOnce.Do(func() {
		file_tableCluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableCluster_proto_rawDescData)
	})
	return file_tableCluster_proto_rawDescData
}

var file_tableCluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tableCluster_proto_goTypes = []interface{}{
	(*NodeList)(nil),      // 0: table_pb.NodeList
	(*NodeVO)(nil),        // 1: table_pb.NodeVO
	(*NodeNameReq)(nil),   // 2: table_pb.NodeNameReq
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
	(*TableQueryReq)(nil), // 4: table_pb.TableQueryReq
	(*TableIdReq)(nil),    // 5: table_pb.TableIdReq
	(*TableCount)(nil),    // 6: table_pb.TableCount
	(*TableList)(nil),     // 7: table_pb.TableList
	(*TableGetVO)(nil),    // 8: table_pb.TableGetVO
}
var file_tableCluster_proto_depIdxs = []int32{
	1, // 0: table_pb.NodeList.Nodes:type_name -> table_pb.NodeVO
	3, // 1: table_pb.TableCluster.NodeListSrv:input_type -> google.protobuf.Empty
	2, // 2: table_pb.TableCluster.NodeOpenSrv:input_type -> table_pb.NodeNameReq
	2, // 3: table_pb.TableCluster.NodeCloseSrv:input_type -> table_pb.NodeNameReq
	4, // 4: table_pb.TableCluster.TableCountSrv:input_type -> table_pb.TableQueryReq
	4, // 5: table_pb.TableCluster.TableListSrv:input_type -> table_pb.TableQueryReq
	5, // 6: table_pb.TableCluster.TableGetSrv:input_type -> table_pb.TableIdReq
	5, // 7: table_pb.TableCluster.TableCloseSrv:input_type -> table_pb.TableIdReq
	0, // 8: table_pb.TableCluster.NodeListSrv:output_type -> table_pb.NodeList
	3, // 9: table_pb.TableCluster.NodeOpenSrv:output_type -> google.protobuf.Empty
	3, // 10: table_pb.TableCluster.NodeCloseSrv:output_type -> google.protobuf.Empty
	6, // 11: table_pb.TableCluster.TableCountSrv:output_type -> table_pb.TableCount
	7, // 12: table_pb.TableCluster.TableListSrv:output_type -> table_pb.TableList
	8, // 13: table_pb.TableCluster.TableGetSrv:output_type -> table_pb.TableGetVO
	3, // 14: table_pb.TableCluster.TableCloseSrv:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tableCluster_proto_init() }
func file_tableCluster_proto_init() {
	if File_tableCluster_proto != nil {
		return
	}
	file_table_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tableCluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeList); i {
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
		file_tableCluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeVO); i {
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
		file_tableCluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNameReq); i {
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
			RawDescriptor: file_tableCluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tableCluster_proto_goTypes,
		DependencyIndexes: file_tableCluster_proto_depIdxs,
		MessageInfos:      file_tableCluster_proto_msgTypes,
	}.Build()
	File_tableCluster_proto = out.File
	file_tableCluster_proto_rawDesc = nil
	file_tableCluster_proto_goTypes = nil
	file_tableCluster_proto_depIdxs = nil
}
