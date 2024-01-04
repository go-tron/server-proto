// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: tableData.proto

package tableData_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	control "server-proto/pb/control"
	game "server-proto/pb/game"
	table "server-proto/pb/table"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TableDataTopic int32

const (
	TableDataTopic_TableDataTopicUnspecified TableDataTopic = 0
	TableDataTopic_TableDataTopicTableEnd    TableDataTopic = 1
)

// Enum value maps for TableDataTopic.
var (
	TableDataTopic_name = map[int32]string{
		0: "TableDataTopicUnspecified",
		1: "TableDataTopicTableEnd",
	}
	TableDataTopic_value = map[string]int32{
		"TableDataTopicUnspecified": 0,
		"TableDataTopicTableEnd":    1,
	}
)

func (x TableDataTopic) Enum() *TableDataTopic {
	p := new(TableDataTopic)
	*p = x
	return p
}

func (x TableDataTopic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TableDataTopic) Descriptor() protoreflect.EnumDescriptor {
	return file_tableData_proto_enumTypes[0].Descriptor()
}

func (TableDataTopic) Type() protoreflect.EnumType {
	return &file_tableData_proto_enumTypes[0]
}

func (x TableDataTopic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TableDataTopic.Descriptor instead.
func (TableDataTopic) EnumDescriptor() ([]byte, []int) {
	return file_tableData_proto_rawDescGZIP(), []int{0}
}

type TableEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config    *control.TableConfig `protobuf:"bytes,1,opt,name=Config,proto3" json:"Config,omitempty"`
	CreatedAt int64                `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	StartAt   int64                `protobuf:"varint,3,opt,name=StartAt,proto3" json:"StartAt,omitempty"`
	EndAt     int64                `protobuf:"varint,4,opt,name=EndAt,proto3" json:"EndAt,omitempty"`
	Users     []*User              `protobuf:"bytes,5,rep,name=Users,proto3" json:"Users,omitempty"`
	Decks     []*Deck              `protobuf:"bytes,6,rep,name=Decks,proto3" json:"Decks,omitempty"`
}

func (x *TableEnd) Reset() {
	*x = TableEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableEnd) ProtoMessage() {}

func (x *TableEnd) ProtoReflect() protoreflect.Message {
	mi := &file_tableData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableEnd.ProtoReflect.Descriptor instead.
func (*TableEnd) Descriptor() ([]byte, []int) {
	return file_tableData_proto_rawDescGZIP(), []int{0}
}

func (x *TableEnd) GetConfig() *control.TableConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *TableEnd) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TableEnd) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *TableEnd) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *TableEnd) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *TableEnd) GetDecks() []*Deck {
	if x != nil {
		return x.Decks
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Robot      bool   `protobuf:"varint,1,opt,name=Robot,proto3" json:"Robot,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserNo     string `protobuf:"bytes,3,opt,name=UserNo,proto3" json:"UserNo,omitempty"`
	Nickname   string `protobuf:"bytes,4,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Avatar     string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Seat       uint32 `protobuf:"varint,6,opt,name=Seat,proto3" json:"Seat,omitempty"`
	Score      int32  `protobuf:"varint,7,opt,name=Score,proto3" json:"Score,omitempty"`
	TableScore int32  `protobuf:"varint,8,opt,name=TableScore,proto3" json:"TableScore,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_tableData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_tableData_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetRobot() bool {
	if x != nil {
		return x.Robot
	}
	return false
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetUserNo() string {
	if x != nil {
		return x.UserNo
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *User) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *User) GetTableScore() int32 {
	if x != nil {
		return x.TableScore
	}
	return 0
}

type Deck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId      int64               `protobuf:"varint,1,opt,name=DeckId,proto3" json:"DeckId,omitempty"`
	DeckNo      uint32              `protobuf:"varint,2,opt,name=DeckNo,proto3" json:"DeckNo,omitempty"`
	StartAt     int64               `protobuf:"varint,3,opt,name=StartAt,proto3" json:"StartAt,omitempty"`
	EndAt       int64               `protobuf:"varint,4,opt,name=EndAt,proto3" json:"EndAt,omitempty"`
	Cards       string              `protobuf:"bytes,5,opt,name=Cards,proto3" json:"Cards,omitempty"`
	HoleCards   string              `protobuf:"bytes,6,opt,name=HoleCards,proto3" json:"HoleCards,omitempty"`
	RateDetails *game.RateDetails   `protobuf:"bytes,7,opt,name=RateDetails,proto3" json:"RateDetails,omitempty"`
	Records     []*table.DeckRecord `protobuf:"bytes,8,rep,name=Records,proto3" json:"Records,omitempty"`
}

func (x *Deck) Reset() {
	*x = Deck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deck) ProtoMessage() {}

func (x *Deck) ProtoReflect() protoreflect.Message {
	mi := &file_tableData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deck.ProtoReflect.Descriptor instead.
func (*Deck) Descriptor() ([]byte, []int) {
	return file_tableData_proto_rawDescGZIP(), []int{2}
}

func (x *Deck) GetDeckId() int64 {
	if x != nil {
		return x.DeckId
	}
	return 0
}

func (x *Deck) GetDeckNo() uint32 {
	if x != nil {
		return x.DeckNo
	}
	return 0
}

func (x *Deck) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *Deck) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *Deck) GetCards() string {
	if x != nil {
		return x.Cards
	}
	return ""
}

func (x *Deck) GetHoleCards() string {
	if x != nil {
		return x.HoleCards
	}
	return ""
}

func (x *Deck) GetRateDetails() *game.RateDetails {
	if x != nil {
		return x.RateDetails
	}
	return nil
}

func (x *Deck) GetRecords() []*table.DeckRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_tableData_proto protoreflect.FileDescriptor

var file_tableData_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x62, 0x1a,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x45, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70,
	0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x45, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a,
	0x05, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x63, 0x6b,
	0x52, 0x05, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x65,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x04, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x44,
	0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x63, 0x6b, 0x4e, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x44, 0x65, 0x63, 0x6b, 0x4e, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x48, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x2a, 0x4b, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x10, 0x01, 0x42, 0x28,
	0x5a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x3b, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tableData_proto_rawDescOnce sync.Once
	file_tableData_proto_rawDescData = file_tableData_proto_rawDesc
)

func file_tableData_proto_rawDescGZIP() []byte {
	file_tableData_proto_rawDescOnce.Do(func() {
		file_tableData_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableData_proto_rawDescData)
	})
	return file_tableData_proto_rawDescData
}

var file_tableData_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tableData_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tableData_proto_goTypes = []interface{}{
	(TableDataTopic)(0),         // 0: tableData_pb.TableDataTopic
	(*TableEnd)(nil),            // 1: tableData_pb.TableEnd
	(*User)(nil),                // 2: tableData_pb.User
	(*Deck)(nil),                // 3: tableData_pb.Deck
	(*control.TableConfig)(nil), // 4: control_pb.TableConfig
	(*game.RateDetails)(nil),    // 5: RateDetails
	(*table.DeckRecord)(nil),    // 6: table_pb.DeckRecord
}
var file_tableData_proto_depIdxs = []int32{
	4, // 0: tableData_pb.TableEnd.Config:type_name -> control_pb.TableConfig
	2, // 1: tableData_pb.TableEnd.Users:type_name -> tableData_pb.User
	3, // 2: tableData_pb.TableEnd.Decks:type_name -> tableData_pb.Deck
	5, // 3: tableData_pb.Deck.RateDetails:type_name -> RateDetails
	6, // 4: tableData_pb.Deck.Records:type_name -> table_pb.DeckRecord
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tableData_proto_init() }
func file_tableData_proto_init() {
	if File_tableData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tableData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableEnd); i {
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
		file_tableData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_tableData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deck); i {
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
			RawDescriptor: file_tableData_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tableData_proto_goTypes,
		DependencyIndexes: file_tableData_proto_depIdxs,
		EnumInfos:         file_tableData_proto_enumTypes,
		MessageInfos:      file_tableData_proto_msgTypes,
	}.Build()
	File_tableData_proto = out.File
	file_tableData_proto_rawDesc = nil
	file_tableData_proto_goTypes = nil
	file_tableData_proto_depIdxs = nil
}
