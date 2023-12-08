// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: control.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TableMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd     uint32     `protobuf:"varint,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Content *anypb.Any `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	UserId  string     `protobuf:"bytes,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TableId string     `protobuf:"bytes,4,opt,name=TableId,proto3" json:"TableId,omitempty"`
}

func (x *TableMessage) Reset() {
	*x = TableMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableMessage) ProtoMessage() {}

func (x *TableMessage) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableMessage.ProtoReflect.Descriptor instead.
func (*TableMessage) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0}
}

func (x *TableMessage) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *TableMessage) GetContent() *anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *TableMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TableMessage) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

type AssignUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserNo   string `protobuf:"bytes,2,opt,name=UserNo,proto3" json:"UserNo,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"` //昵称
	Avatar   string `protobuf:"bytes,4,opt,name=Avatar,proto3" json:"Avatar,omitempty"`     //头像
	Gender   uint32 `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`    //性别
	Seat     uint32 `protobuf:"varint,6,opt,name=Seat,proto3" json:"Seat,omitempty"`        //座位号0,1,2 逆时针方向排列
	Score    int32  `protobuf:"varint,7,opt,name=Score,proto3" json:"Score,omitempty"`      //分数
	Owner    bool   `protobuf:"varint,8,opt,name=Owner,proto3" json:"Owner,omitempty"`      //房主
	Rank     uint32 `protobuf:"varint,9,opt,name=Rank,proto3" json:"Rank,omitempty"`        //当前排名
}

func (x *AssignUser) Reset() {
	*x = AssignUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignUser) ProtoMessage() {}

func (x *AssignUser) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignUser.ProtoReflect.Descriptor instead.
func (*AssignUser) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{1}
}

func (x *AssignUser) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AssignUser) GetUserNo() string {
	if x != nil {
		return x.UserNo
	}
	return ""
}

func (x *AssignUser) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AssignUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AssignUser) GetGender() uint32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *AssignUser) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *AssignUser) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *AssignUser) GetOwner() bool {
	if x != nil {
		return x.Owner
	}
	return false
}

func (x *AssignUser) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type CreateTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users             []*AssignUser `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	GameId            int64         `protobuf:"varint,2,opt,name=GameId,proto3" json:"GameId,omitempty"`                               //游戏ID
	TableId           int64         `protobuf:"varint,3,opt,name=TableId,proto3" json:"TableId,omitempty"`                             //牌桌ID
	ForceJoinTimeout  uint32        `protobuf:"varint,4,opt,name=ForceJoinTimeout,proto3" json:"ForceJoinTimeout,omitempty"`           //强制加入时间
	RoundId           int64         `protobuf:"varint,5,opt,name=RoundId,proto3" json:"RoundId,omitempty"`                             //轮次ID
	TableNo           string        `protobuf:"bytes,6,opt,name=TableNo,proto3" json:"TableNo,omitempty"`                              //牌桌编号
	GameType          GameType      `protobuf:"varint,7,opt,name=GameType,proto3,enum=GameType" json:"GameType,omitempty"`             //游戏类型
	GameCategory      GameCategory  `protobuf:"varint,8,opt,name=GameCategory,proto3,enum=GameCategory" json:"GameCategory,omitempty"` //游戏类别
	GameRuleId        uint32        `protobuf:"varint,9,opt,name=GameRuleId,proto3" json:"GameRuleId,omitempty"`                       //游戏规则ID
	GameConfigId      uint32        `protobuf:"varint,10,opt,name=GameConfigId,proto3" json:"GameConfigId,omitempty"`                  //游戏配置ID
	Title             string        `protobuf:"bytes,11,opt,name=Title,proto3" json:"Title,omitempty"`                                 //游戏标题
	Image             string        `protobuf:"bytes,12,opt,name=Image,proto3" json:"Image,omitempty"`                                 //游戏图标
	BaseScore         uint32        `protobuf:"varint,13,opt,name=BaseScore,proto3" json:"BaseScore,omitempty"`                        //底分
	InitialScore      uint32        `protobuf:"varint,14,opt,name=InitialScore,proto3" json:"InitialScore,omitempty"`                  //初始积分
	CallRateTimeout   uint32        `protobuf:"varint,15,opt,name=CallRateTimeout,proto3" json:"CallRateTimeout,omitempty"`            //叫分超时时间
	CallDoubleTimeout uint32        `protobuf:"varint,16,opt,name=CallDoubleTimeout,proto3" json:"CallDoubleTimeout,omitempty"`        //叫双倍超时时间
	PlayTimeout       uint32        `protobuf:"varint,17,opt,name=PlayTimeout,proto3" json:"PlayTimeout,omitempty"`                    //打牌超时时间
	FollowTimeout     uint32        `protobuf:"varint,18,opt,name=FollowTimeout,proto3" json:"FollowTimeout,omitempty"`                //跟牌超时时间
	CantFollowTimeout uint32        `protobuf:"varint,19,opt,name=CantFollowTimeout,proto3" json:"CantFollowTimeout,omitempty"`        //要不起超时时间
	People            uint32        `protobuf:"varint,20,opt,name=People,proto3" json:"People,omitempty"`                              //参加总人数
	CurrentPeople     uint32        `protobuf:"varint,21,opt,name=CurrentPeople,proto3" json:"CurrentPeople,omitempty"`                //当前剩余人数
	Rounds            uint32        `protobuf:"varint,22,opt,name=Rounds,proto3" json:"Rounds,omitempty"`                              //游戏总轮次
	CurrentRound      uint32        `protobuf:"varint,23,opt,name=CurrentRound,proto3" json:"CurrentRound,omitempty"`                  //游戏当前轮次
	PromotePeople     uint32        `protobuf:"varint,24,opt,name=PromotePeople,proto3" json:"PromotePeople,omitempty"`                //当前轮次晋级人数
	Decks             uint32        `protobuf:"varint,25,opt,name=Decks,proto3" json:"Decks,omitempty"`                                //打几副牌
	WithBomb          bool          `protobuf:"varint,26,opt,name=WithBomb,proto3" json:"WithBomb,omitempty"`                          //是否带炸弹
	WithDouble        bool          `protobuf:"varint,27,opt,name=WithDouble,proto3" json:"WithDouble,omitempty"`                      //是否带加倍
	MaxRate           uint32        `protobuf:"varint,28,opt,name=MaxRate,proto3" json:"MaxRate,omitempty"`                            //最大倍率
	MinEntryGold      uint32        `protobuf:"varint,29,opt,name=MinEntryGold,proto3" json:"MinEntryGold,omitempty"`                  //最少携带金币
	MaxEntryGold      uint32        `protobuf:"varint,30,opt,name=MaxEntryGold,proto3" json:"MaxEntryGold,omitempty"`                  //最多携带金币
}

func (x *CreateTable) Reset() {
	*x = CreateTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTable) ProtoMessage() {}

func (x *CreateTable) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTable.ProtoReflect.Descriptor instead.
func (*CreateTable) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTable) GetUsers() []*AssignUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *CreateTable) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *CreateTable) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *CreateTable) GetForceJoinTimeout() uint32 {
	if x != nil {
		return x.ForceJoinTimeout
	}
	return 0
}

func (x *CreateTable) GetRoundId() int64 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

func (x *CreateTable) GetTableNo() string {
	if x != nil {
		return x.TableNo
	}
	return ""
}

func (x *CreateTable) GetGameType() GameType {
	if x != nil {
		return x.GameType
	}
	return GameType_GameTypeUnspecified
}

func (x *CreateTable) GetGameCategory() GameCategory {
	if x != nil {
		return x.GameCategory
	}
	return GameCategory_GameCategoryUnspecified
}

func (x *CreateTable) GetGameRuleId() uint32 {
	if x != nil {
		return x.GameRuleId
	}
	return 0
}

func (x *CreateTable) GetGameConfigId() uint32 {
	if x != nil {
		return x.GameConfigId
	}
	return 0
}

func (x *CreateTable) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTable) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateTable) GetBaseScore() uint32 {
	if x != nil {
		return x.BaseScore
	}
	return 0
}

func (x *CreateTable) GetInitialScore() uint32 {
	if x != nil {
		return x.InitialScore
	}
	return 0
}

func (x *CreateTable) GetCallRateTimeout() uint32 {
	if x != nil {
		return x.CallRateTimeout
	}
	return 0
}

func (x *CreateTable) GetCallDoubleTimeout() uint32 {
	if x != nil {
		return x.CallDoubleTimeout
	}
	return 0
}

func (x *CreateTable) GetPlayTimeout() uint32 {
	if x != nil {
		return x.PlayTimeout
	}
	return 0
}

func (x *CreateTable) GetFollowTimeout() uint32 {
	if x != nil {
		return x.FollowTimeout
	}
	return 0
}

func (x *CreateTable) GetCantFollowTimeout() uint32 {
	if x != nil {
		return x.CantFollowTimeout
	}
	return 0
}

func (x *CreateTable) GetPeople() uint32 {
	if x != nil {
		return x.People
	}
	return 0
}

func (x *CreateTable) GetCurrentPeople() uint32 {
	if x != nil {
		return x.CurrentPeople
	}
	return 0
}

func (x *CreateTable) GetRounds() uint32 {
	if x != nil {
		return x.Rounds
	}
	return 0
}

func (x *CreateTable) GetCurrentRound() uint32 {
	if x != nil {
		return x.CurrentRound
	}
	return 0
}

func (x *CreateTable) GetPromotePeople() uint32 {
	if x != nil {
		return x.PromotePeople
	}
	return 0
}

func (x *CreateTable) GetDecks() uint32 {
	if x != nil {
		return x.Decks
	}
	return 0
}

func (x *CreateTable) GetWithBomb() bool {
	if x != nil {
		return x.WithBomb
	}
	return false
}

func (x *CreateTable) GetWithDouble() bool {
	if x != nil {
		return x.WithDouble
	}
	return false
}

func (x *CreateTable) GetMaxRate() uint32 {
	if x != nil {
		return x.MaxRate
	}
	return 0
}

func (x *CreateTable) GetMinEntryGold() uint32 {
	if x != nil {
		return x.MinEntryGold
	}
	return 0
}

func (x *CreateTable) GetMaxEntryGold() uint32 {
	if x != nil {
		return x.MaxEntryGold
	}
	return 0
}

type CreateTableResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId  int64 `protobuf:"varint,1,opt,name=GameId,proto3" json:"GameId,omitempty"`
	TableId int64 `protobuf:"varint,2,opt,name=TableId,proto3" json:"TableId,omitempty"`
	Succeed bool  `protobuf:"varint,3,opt,name=Succeed,proto3" json:"Succeed,omitempty"`
}

func (x *CreateTableResult) Reset() {
	*x = CreateTableResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableResult) ProtoMessage() {}

func (x *CreateTableResult) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableResult.ProtoReflect.Descriptor instead.
func (*CreateTableResult) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTableResult) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *CreateTableResult) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *CreateTableResult) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

type CoverRobot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameConfigId uint32 `protobuf:"varint,1,opt,name=GameConfigId,proto3" json:"GameConfigId,omitempty"`
	TableId      int64  `protobuf:"varint,2,opt,name=TableId,proto3" json:"TableId,omitempty"`
}

func (x *CoverRobot) Reset() {
	*x = CoverRobot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoverRobot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoverRobot) ProtoMessage() {}

func (x *CoverRobot) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoverRobot.ProtoReflect.Descriptor instead.
func (*CoverRobot) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{4}
}

func (x *CoverRobot) GetGameConfigId() uint32 {
	if x != nil {
		return x.GameConfigId
	}
	return 0
}

func (x *CoverRobot) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

type AssignTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users            []*AssignUser `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	GameId           int64         `protobuf:"varint,2,opt,name=GameId,proto3" json:"GameId,omitempty"`                     //游戏ID
	TableId          int64         `protobuf:"varint,3,opt,name=TableId,proto3" json:"TableId,omitempty"`                   //牌桌ID
	ForceJoinTimeout uint32        `protobuf:"varint,4,opt,name=ForceJoinTimeout,proto3" json:"ForceJoinTimeout,omitempty"` //强制加入时间
}

func (x *AssignTable) Reset() {
	*x = AssignTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignTable) ProtoMessage() {}

func (x *AssignTable) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignTable.ProtoReflect.Descriptor instead.
func (*AssignTable) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{5}
}

func (x *AssignTable) GetUsers() []*AssignUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *AssignTable) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *AssignTable) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *AssignTable) GetForceJoinTimeout() uint32 {
	if x != nil {
		return x.ForceJoinTimeout
	}
	return 0
}

type ForceStartTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId  int64 `protobuf:"varint,1,opt,name=GameId,proto3" json:"GameId,omitempty"`
	TableId int64 `protobuf:"varint,2,opt,name=TableId,proto3" json:"TableId,omitempty"`
}

func (x *ForceStartTable) Reset() {
	*x = ForceStartTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceStartTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceStartTable) ProtoMessage() {}

func (x *ForceStartTable) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceStartTable.ProtoReflect.Descriptor instead.
func (*ForceStartTable) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{6}
}

func (x *ForceStartTable) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *ForceStartTable) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

type TableEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId  int64 `protobuf:"varint,1,opt,name=GameId,proto3" json:"GameId,omitempty"`
	TableId int64 `protobuf:"varint,2,opt,name=TableId,proto3" json:"TableId,omitempty"`
}

func (x *TableEnd) Reset() {
	*x = TableEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableEnd) ProtoMessage() {}

func (x *TableEnd) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[7]
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
	return file_control_proto_rawDescGZIP(), []int{7}
}

func (x *TableEnd) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *TableEnd) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

type LeaveTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users   []int64 `protobuf:"varint,1,rep,packed,name=Users,proto3" json:"Users,omitempty"`
	GameId  int64   `protobuf:"varint,2,opt,name=GameId,proto3" json:"GameId,omitempty"`   //游戏ID
	TableId int64   `protobuf:"varint,3,opt,name=TableId,proto3" json:"TableId,omitempty"` //牌桌ID
}

func (x *LeaveTable) Reset() {
	*x = LeaveTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveTable) ProtoMessage() {}

func (x *LeaveTable) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveTable.ProtoReflect.Descriptor instead.
func (*LeaveTable) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{8}
}

func (x *LeaveTable) GetUsers() []int64 {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *LeaveTable) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *LeaveTable) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

var File_control_proto protoreflect.FileDescriptor

var file_control_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x0a,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x22, 0xf0, 0x07, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x12,
	0x25, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x09, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x61,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x47, 0x61, 0x6d,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x61, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x61, 0x73,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x42, 0x61,
	0x73, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x11, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x43,
	0x61, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x43, 0x61, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x6f,
	0x70, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x65, 0x63,
	0x6b, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6d, 0x62, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6d, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4d, 0x61,
	0x78, 0x52, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x4d, 0x69, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x78,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x4d, 0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x6f, 0x6c, 0x64, 0x22, 0x5f, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x22, 0x4a,
	0x0a, 0x0a, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x43, 0x0a, 0x0f, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_proto_rawDescOnce sync.Once
	file_control_proto_rawDescData = file_control_proto_rawDesc
)

func file_control_proto_rawDescGZIP() []byte {
	file_control_proto_rawDescOnce.Do(func() {
		file_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_proto_rawDescData)
	})
	return file_control_proto_rawDescData
}

var file_control_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_control_proto_goTypes = []interface{}{
	(*TableMessage)(nil),      // 0: TableMessage
	(*AssignUser)(nil),        // 1: AssignUser
	(*CreateTable)(nil),       // 2: CreateTable
	(*CreateTableResult)(nil), // 3: CreateTableResult
	(*CoverRobot)(nil),        // 4: CoverRobot
	(*AssignTable)(nil),       // 5: AssignTable
	(*ForceStartTable)(nil),   // 6: ForceStartTable
	(*TableEnd)(nil),          // 7: TableEnd
	(*LeaveTable)(nil),        // 8: LeaveTable
	(*anypb.Any)(nil),         // 9: google.protobuf.Any
	(GameType)(0),             // 10: GameType
	(GameCategory)(0),         // 11: GameCategory
}
var file_control_proto_depIdxs = []int32{
	9,  // 0: TableMessage.Content:type_name -> google.protobuf.Any
	1,  // 1: CreateTable.Users:type_name -> AssignUser
	10, // 2: CreateTable.GameType:type_name -> GameType
	11, // 3: CreateTable.GameCategory:type_name -> GameCategory
	1,  // 4: AssignTable.Users:type_name -> AssignUser
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_control_proto_init() }
func file_control_proto_init() {
	if File_control_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableMessage); i {
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
		file_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignUser); i {
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
		file_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTable); i {
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
		file_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableResult); i {
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
		file_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoverRobot); i {
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
		file_control_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignTable); i {
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
		file_control_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceStartTable); i {
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
		file_control_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_control_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveTable); i {
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
			RawDescriptor: file_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_control_proto_goTypes,
		DependencyIndexes: file_control_proto_depIdxs,
		MessageInfos:      file_control_proto_msgTypes,
	}.Build()
	File_control_proto = out.File
	file_control_proto_rawDesc = nil
	file_control_proto_goTypes = nil
	file_control_proto_depIdxs = nil
}
