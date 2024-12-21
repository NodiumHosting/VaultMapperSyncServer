// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.28.3
// source: VaultMapperProtocol/vaultmapper.proto

package proto

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

type CellType int32

const (
	CellType_CELLTYPE_UNKNOWN  CellType = 0
	CellType_CELLTYPE_ROOM     CellType = 1
	CellType_CELLTYPE_TUNNEL_X CellType = 2
	CellType_CELLTYPE_TUNNEL_Z CellType = 3
)

// Enum value maps for CellType.
var (
	CellType_name = map[int32]string{
		0: "CELLTYPE_UNKNOWN",
		1: "CELLTYPE_ROOM",
		2: "CELLTYPE_TUNNEL_X",
		3: "CELLTYPE_TUNNEL_Z",
	}
	CellType_value = map[string]int32{
		"CELLTYPE_UNKNOWN":  0,
		"CELLTYPE_ROOM":     1,
		"CELLTYPE_TUNNEL_X": 2,
		"CELLTYPE_TUNNEL_Z": 3,
	}
)

func (x CellType) Enum() *CellType {
	p := new(CellType)
	*p = x
	return p
}

func (x CellType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CellType) Descriptor() protoreflect.EnumDescriptor {
	return file_VaultMapperProtocol_vaultmapper_proto_enumTypes[0].Descriptor()
}

func (CellType) Type() protoreflect.EnumType {
	return &file_VaultMapperProtocol_vaultmapper_proto_enumTypes[0]
}

func (x CellType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CellType.Descriptor instead.
func (CellType) EnumDescriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{0}
}

type RoomType int32

const (
	RoomType_ROOMTYPE_UNKNOWN   RoomType = 0
	RoomType_ROOMTYPE_START     RoomType = 1
	RoomType_ROOMTYPE_BASIC     RoomType = 2
	RoomType_ROOMTYPE_ORE       RoomType = 3
	RoomType_ROOMTYPE_CHALLENGE RoomType = 4
	RoomType_ROOMTYPE_OMEGA     RoomType = 5
)

// Enum value maps for RoomType.
var (
	RoomType_name = map[int32]string{
		0: "ROOMTYPE_UNKNOWN",
		1: "ROOMTYPE_START",
		2: "ROOMTYPE_BASIC",
		3: "ROOMTYPE_ORE",
		4: "ROOMTYPE_CHALLENGE",
		5: "ROOMTYPE_OMEGA",
	}
	RoomType_value = map[string]int32{
		"ROOMTYPE_UNKNOWN":   0,
		"ROOMTYPE_START":     1,
		"ROOMTYPE_BASIC":     2,
		"ROOMTYPE_ORE":       3,
		"ROOMTYPE_CHALLENGE": 4,
		"ROOMTYPE_OMEGA":     5,
	}
)

func (x RoomType) Enum() *RoomType {
	p := new(RoomType)
	*p = x
	return p
}

func (x RoomType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomType) Descriptor() protoreflect.EnumDescriptor {
	return file_VaultMapperProtocol_vaultmapper_proto_enumTypes[1].Descriptor()
}

func (RoomType) Type() protoreflect.EnumType {
	return &file_VaultMapperProtocol_vaultmapper_proto_enumTypes[1]
}

func (x RoomType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomType.Descriptor instead.
func (RoomType) EnumDescriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{1}
}

type RoomName int32

const (
	RoomName_ROOMNAME_UNKNOWN       RoomName = 0
	RoomName_ROOMNAME_BLACKSMITH    RoomName = 1
	RoomName_ROOMNAME_COVE          RoomName = 2
	RoomName_ROOMNAME_CRYSTAL_CAVES RoomName = 3
	RoomName_ROOMNAME_DIG_SITE      RoomName = 4
	RoomName_ROOMNAME_DRAGON        RoomName = 5
	RoomName_ROOMNAME_FACTORY       RoomName = 6
	RoomName_ROOMNAME_LIBRARY       RoomName = 7
	RoomName_ROOMNAME_MINE          RoomName = 8
	RoomName_ROOMNAME_MUSH_ROOM     RoomName = 9
	RoomName_ROOMNAME_PAINTING      RoomName = 10
	RoomName_ROOMNAME_VENDOR        RoomName = 11
	RoomName_ROOMNAME_VILLAGE       RoomName = 12
	RoomName_ROOMNAME_WILD_WEST     RoomName = 13
	RoomName_ROOMNAME_X_MARK        RoomName = 14
	RoomName_ROOMNAME_CUBE          RoomName = 15
	RoomName_ROOMNAME_LAB           RoomName = 16
	RoomName_ROOMNAME_RAID          RoomName = 17
)

// Enum value maps for RoomName.
var (
	RoomName_name = map[int32]string{
		0:  "ROOMNAME_UNKNOWN",
		1:  "ROOMNAME_BLACKSMITH",
		2:  "ROOMNAME_COVE",
		3:  "ROOMNAME_CRYSTAL_CAVES",
		4:  "ROOMNAME_DIG_SITE",
		5:  "ROOMNAME_DRAGON",
		6:  "ROOMNAME_FACTORY",
		7:  "ROOMNAME_LIBRARY",
		8:  "ROOMNAME_MINE",
		9:  "ROOMNAME_MUSH_ROOM",
		10: "ROOMNAME_PAINTING",
		11: "ROOMNAME_VENDOR",
		12: "ROOMNAME_VILLAGE",
		13: "ROOMNAME_WILD_WEST",
		14: "ROOMNAME_X_MARK",
		15: "ROOMNAME_CUBE",
		16: "ROOMNAME_LAB",
		17: "ROOMNAME_RAID",
	}
	RoomName_value = map[string]int32{
		"ROOMNAME_UNKNOWN":       0,
		"ROOMNAME_BLACKSMITH":    1,
		"ROOMNAME_COVE":          2,
		"ROOMNAME_CRYSTAL_CAVES": 3,
		"ROOMNAME_DIG_SITE":      4,
		"ROOMNAME_DRAGON":        5,
		"ROOMNAME_FACTORY":       6,
		"ROOMNAME_LIBRARY":       7,
		"ROOMNAME_MINE":          8,
		"ROOMNAME_MUSH_ROOM":     9,
		"ROOMNAME_PAINTING":      10,
		"ROOMNAME_VENDOR":        11,
		"ROOMNAME_VILLAGE":       12,
		"ROOMNAME_WILD_WEST":     13,
		"ROOMNAME_X_MARK":        14,
		"ROOMNAME_CUBE":          15,
		"ROOMNAME_LAB":           16,
		"ROOMNAME_RAID":          17,
	}
)

func (x RoomName) Enum() *RoomName {
	p := new(RoomName)
	*p = x
	return p
}

func (x RoomName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomName) Descriptor() protoreflect.EnumDescriptor {
	return file_VaultMapperProtocol_vaultmapper_proto_enumTypes[2].Descriptor()
}

func (RoomName) Type() protoreflect.EnumType {
	return &file_VaultMapperProtocol_vaultmapper_proto_enumTypes[2]
}

func (x RoomName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomName.Descriptor instead.
func (RoomName) EnumDescriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{2}
}

type Vault struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cells         []*VaultCell           `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
	Players       []*VaultPlayer         `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vault) Reset() {
	*x = Vault{}
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vault) ProtoMessage() {}

func (x *Vault) ProtoReflect() protoreflect.Message {
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vault.ProtoReflect.Descriptor instead.
func (*Vault) Descriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{0}
}

func (x *Vault) GetCells() []*VaultCell {
	if x != nil {
		return x.Cells
	}
	return nil
}

func (x *Vault) GetPlayers() []*VaultPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

type VaultPlayer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             int32                  `protobuf:"zigzag32,1,opt,name=x,proto3" json:"x,omitempty"`
	Z             int32                  `protobuf:"zigzag32,2,opt,name=z,proto3" json:"z,omitempty"`
	Uuid          string                 `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Color         *Color                 `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Yaw           float32                `protobuf:"fixed32,5,opt,name=yaw,proto3" json:"yaw,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VaultPlayer) Reset() {
	*x = VaultPlayer{}
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VaultPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultPlayer) ProtoMessage() {}

func (x *VaultPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultPlayer.ProtoReflect.Descriptor instead.
func (*VaultPlayer) Descriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{1}
}

func (x *VaultPlayer) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *VaultPlayer) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *VaultPlayer) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *VaultPlayer) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *VaultPlayer) GetYaw() float32 {
	if x != nil {
		return x.Yaw
	}
	return 0
}

type Color struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	R             uint32                 `protobuf:"varint,1,opt,name=r,proto3" json:"r,omitempty"`
	G             uint32                 `protobuf:"varint,2,opt,name=g,proto3" json:"g,omitempty"`
	B             uint32                 `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Color) Reset() {
	*x = Color{}
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{2}
}

func (x *Color) GetR() uint32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *Color) GetG() uint32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *Color) GetB() uint32 {
	if x != nil {
		return x.B
	}
	return 0
}

type VaultCell struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             int32                  `protobuf:"zigzag32,1,opt,name=x,proto3" json:"x,omitempty"`
	Z             int32                  `protobuf:"zigzag32,2,opt,name=z,proto3" json:"z,omitempty"`
	CellType      CellType               `protobuf:"varint,3,opt,name=cellType,proto3,enum=CellType" json:"cellType,omitempty"`
	RoomType      RoomType               `protobuf:"varint,4,opt,name=roomType,proto3,enum=RoomType" json:"roomType,omitempty"`
	RoomName      RoomName               `protobuf:"varint,5,opt,name=roomName,proto3,enum=RoomName" json:"roomName,omitempty"`
	Explored      bool                   `protobuf:"varint,6,opt,name=explored,proto3" json:"explored,omitempty"`
	Inscribed     bool                   `protobuf:"varint,7,opt,name=inscribed,proto3" json:"inscribed,omitempty"`
	Marked        bool                   `protobuf:"varint,8,opt,name=marked,proto3" json:"marked,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VaultCell) Reset() {
	*x = VaultCell{}
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VaultCell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultCell) ProtoMessage() {}

func (x *VaultCell) ProtoReflect() protoreflect.Message {
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultCell.ProtoReflect.Descriptor instead.
func (*VaultCell) Descriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{3}
}

func (x *VaultCell) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *VaultCell) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *VaultCell) GetCellType() CellType {
	if x != nil {
		return x.CellType
	}
	return CellType_CELLTYPE_UNKNOWN
}

func (x *VaultCell) GetRoomType() RoomType {
	if x != nil {
		return x.RoomType
	}
	return RoomType_ROOMTYPE_UNKNOWN
}

func (x *VaultCell) GetRoomName() RoomName {
	if x != nil {
		return x.RoomName
	}
	return RoomName_ROOMNAME_UNKNOWN
}

func (x *VaultCell) GetExplored() bool {
	if x != nil {
		return x.Explored
	}
	return false
}

func (x *VaultCell) GetInscribed() bool {
	if x != nil {
		return x.Inscribed
	}
	return false
}

func (x *VaultCell) GetMarked() bool {
	if x != nil {
		return x.Marked
	}
	return false
}

type PlayerUUID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerUUID) Reset() {
	*x = PlayerUUID{}
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerUUID) ProtoMessage() {}

func (x *PlayerUUID) ProtoReflect() protoreflect.Message {
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerUUID.ProtoReflect.Descriptor instead.
func (*PlayerUUID) Descriptor() ([]byte, []int) {
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerUUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_VaultMapperProtocol_vaultmapper_proto_msgTypes[5]
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
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP(), []int{5}
}

var File_VaultMapperProtocol_vaultmapper_proto protoreflect.FileDescriptor

var file_VaultMapperProtocol_vaultmapper_proto_rawDesc = []byte{
	0x0a, 0x25, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x6d, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x05, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x20, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x65, 0x6c,
	0x6c, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x6d, 0x0a, 0x0b, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x01, 0x7a, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x61, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x79, 0x61, 0x77, 0x22, 0x31, 0x0a, 0x05, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x72,
	0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x67, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x62, 0x22, 0xee, 0x01, 0x0a,
	0x09, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x01, 0x7a, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x22, 0x20, 0x0a,
	0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x61, 0x0a, 0x08, 0x43, 0x65, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x45, 0x4c, 0x4c, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x45,
	0x4c, 0x4c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x45, 0x4c, 0x4c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x58, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x45, 0x4c, 0x4c, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x5a, 0x10, 0x03, 0x2a, 0x86, 0x01, 0x0a, 0x08,
	0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x4f, 0x4f, 0x4d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4f, 0x4d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x41, 0x53, 0x49, 0x43, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4f, 0x4d, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x4f, 0x4d,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4f, 0x4d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4d, 0x45,
	0x47, 0x41, 0x10, 0x05, 0x2a, 0x92, 0x03, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x4f, 0x4d, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x53, 0x4d, 0x49, 0x54, 0x48, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x56,
	0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x43, 0x52, 0x59, 0x53, 0x54, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x56, 0x45, 0x53, 0x10, 0x03, 0x12,
	0x15, 0x0a, 0x11, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x44, 0x49, 0x47, 0x5f,
	0x53, 0x49, 0x54, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x47, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x52,
	0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x46, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x59, 0x10,
	0x06, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4c, 0x49,
	0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x4f, 0x4d, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f,
	0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x55, 0x53, 0x48, 0x5f, 0x52, 0x4f, 0x4f, 0x4d,
	0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x50,
	0x41, 0x49, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x4f,
	0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x10, 0x0b, 0x12, 0x14,
	0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x56, 0x49, 0x4c, 0x4c, 0x41,
	0x47, 0x45, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x57, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x45, 0x53, 0x54, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f,
	0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x58, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x10,
	0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x55,
	0x42, 0x45, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x4c, 0x41, 0x42, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x4f, 0x4d, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x52, 0x41, 0x49, 0x44, 0x10, 0x11, 0x32, 0x92, 0x01, 0x0a, 0x09, 0x56, 0x4d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x0a, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x65, 0x6c, 0x6c, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x08, 0x4d,
	0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x06, 0x2e, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VaultMapperProtocol_vaultmapper_proto_rawDescOnce sync.Once
	file_VaultMapperProtocol_vaultmapper_proto_rawDescData = file_VaultMapperProtocol_vaultmapper_proto_rawDesc
)

func file_VaultMapperProtocol_vaultmapper_proto_rawDescGZIP() []byte {
	file_VaultMapperProtocol_vaultmapper_proto_rawDescOnce.Do(func() {
		file_VaultMapperProtocol_vaultmapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_VaultMapperProtocol_vaultmapper_proto_rawDescData)
	})
	return file_VaultMapperProtocol_vaultmapper_proto_rawDescData
}

var file_VaultMapperProtocol_vaultmapper_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_VaultMapperProtocol_vaultmapper_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_VaultMapperProtocol_vaultmapper_proto_goTypes = []any{
	(CellType)(0),       // 0: CellType
	(RoomType)(0),       // 1: RoomType
	(RoomName)(0),       // 2: RoomName
	(*Vault)(nil),       // 3: Vault
	(*VaultPlayer)(nil), // 4: VaultPlayer
	(*Color)(nil),       // 5: Color
	(*VaultCell)(nil),   // 6: VaultCell
	(*PlayerUUID)(nil),  // 7: PlayerUUID
	(*Empty)(nil),       // 8: Empty
}
var file_VaultMapperProtocol_vaultmapper_proto_depIdxs = []int32{
	6,  // 0: Vault.cells:type_name -> VaultCell
	4,  // 1: Vault.players:type_name -> VaultPlayer
	5,  // 2: VaultPlayer.color:type_name -> Color
	0,  // 3: VaultCell.cellType:type_name -> CellType
	1,  // 4: VaultCell.roomType:type_name -> RoomType
	2,  // 5: VaultCell.roomName:type_name -> RoomName
	6,  // 6: VMService.NewVaultCell:input_type -> VaultCell
	4,  // 7: VMService.Movement:input_type -> VaultPlayer
	3,  // 8: VMService.VaultSync:input_type -> Vault
	7,  // 9: VMService.PlayerLeave:input_type -> PlayerUUID
	8,  // 10: VMService.NewVaultCell:output_type -> Empty
	8,  // 11: VMService.Movement:output_type -> Empty
	8,  // 12: VMService.VaultSync:output_type -> Empty
	8,  // 13: VMService.PlayerLeave:output_type -> Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_VaultMapperProtocol_vaultmapper_proto_init() }
func file_VaultMapperProtocol_vaultmapper_proto_init() {
	if File_VaultMapperProtocol_vaultmapper_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_VaultMapperProtocol_vaultmapper_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_VaultMapperProtocol_vaultmapper_proto_goTypes,
		DependencyIndexes: file_VaultMapperProtocol_vaultmapper_proto_depIdxs,
		EnumInfos:         file_VaultMapperProtocol_vaultmapper_proto_enumTypes,
		MessageInfos:      file_VaultMapperProtocol_vaultmapper_proto_msgTypes,
	}.Build()
	File_VaultMapperProtocol_vaultmapper_proto = out.File
	file_VaultMapperProtocol_vaultmapper_proto_rawDesc = nil
	file_VaultMapperProtocol_vaultmapper_proto_goTypes = nil
	file_VaultMapperProtocol_vaultmapper_proto_depIdxs = nil
}
