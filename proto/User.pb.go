// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User.proto

package User

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Gender int32

const (
	Gender_NOTDEFINED Gender = 0
	Gender_MALE       Gender = 1
	Gender_FEMALE     Gender = 2
)

var Gender_name = map[int32]string{
	0: "NOTDEFINED",
	1: "MALE",
	2: "FEMALE",
}

var Gender_value = map[string]int32{
	"NOTDEFINED": 0,
	"MALE":       1,
	"FEMALE":     2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{0}
}

type RemoveTvDeviceRequest struct {
	GoogleId             string   `protobuf:"bytes,1,opt,name=googleId,proto3" json:"googleId,omitempty"`
	TvEmac               string   `protobuf:"bytes,2,opt,name=tvEmac,proto3" json:"tvEmac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTvDeviceRequest) Reset()         { *m = RemoveTvDeviceRequest{} }
func (m *RemoveTvDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTvDeviceRequest) ProtoMessage()    {}
func (*RemoveTvDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{0}
}

func (m *RemoveTvDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTvDeviceRequest.Unmarshal(m, b)
}
func (m *RemoveTvDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTvDeviceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveTvDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTvDeviceRequest.Merge(m, src)
}
func (m *RemoveTvDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTvDeviceRequest.Size(m)
}
func (m *RemoveTvDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTvDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTvDeviceRequest proto.InternalMessageInfo

func (m *RemoveTvDeviceRequest) GetGoogleId() string {
	if m != nil {
		return m.GoogleId
	}
	return ""
}

func (m *RemoveTvDeviceRequest) GetTvEmac() string {
	if m != nil {
		return m.TvEmac
	}
	return ""
}

type RemoveTvDeviceResponse struct {
	IsTvDeviceRemoved    bool     `protobuf:"varint,1,opt,name=isTvDeviceRemoved,proto3" json:"isTvDeviceRemoved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTvDeviceResponse) Reset()         { *m = RemoveTvDeviceResponse{} }
func (m *RemoveTvDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveTvDeviceResponse) ProtoMessage()    {}
func (*RemoveTvDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{1}
}

func (m *RemoveTvDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTvDeviceResponse.Unmarshal(m, b)
}
func (m *RemoveTvDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTvDeviceResponse.Marshal(b, m, deterministic)
}
func (m *RemoveTvDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTvDeviceResponse.Merge(m, src)
}
func (m *RemoveTvDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveTvDeviceResponse.Size(m)
}
func (m *RemoveTvDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTvDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTvDeviceResponse proto.InternalMessageInfo

func (m *RemoveTvDeviceResponse) GetIsTvDeviceRemoved() bool {
	if m != nil {
		return m.IsTvDeviceRemoved
	}
	return false
}

type GetRequest struct {
	GoogleId             string   `protobuf:"bytes,1,opt,name=googleId,proto3" json:"googleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetGoogleId() string {
	if m != nil {
		return m.GoogleId
	}
	return ""
}

type DeleteRequest struct {
	GoogleId             string   `protobuf:"bytes,1,opt,name=googleId,proto3" json:"googleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetGoogleId() string {
	if m != nil {
		return m.GoogleId
	}
	return ""
}

type DeleteReponse struct {
	IsDeleted            bool     `protobuf:"varint,1,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReponse) Reset()         { *m = DeleteReponse{} }
func (m *DeleteReponse) String() string { return proto.CompactTextString(m) }
func (*DeleteReponse) ProtoMessage()    {}
func (*DeleteReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{4}
}

func (m *DeleteReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReponse.Unmarshal(m, b)
}
func (m *DeleteReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReponse.Marshal(b, m, deterministic)
}
func (m *DeleteReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReponse.Merge(m, src)
}
func (m *DeleteReponse) XXX_Size() int {
	return xxx_messageInfo_DeleteReponse.Size(m)
}
func (m *DeleteReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReponse proto.InternalMessageInfo

func (m *DeleteReponse) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

type LinkedDeviceResponse struct {
	IsLinkedDeviceFetched bool            `protobuf:"varint,1,opt,name=isLinkedDeviceFetched,proto3" json:"isLinkedDeviceFetched,omitempty"`
	LinkedDevices         []*LinkedDevice `protobuf:"bytes,2,rep,name=linkedDevices,proto3" json:"linkedDevices,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}        `json:"-"`
	XXX_unrecognized      []byte          `json:"-"`
	XXX_sizecache         int32           `json:"-"`
}

func (m *LinkedDeviceResponse) Reset()         { *m = LinkedDeviceResponse{} }
func (m *LinkedDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*LinkedDeviceResponse) ProtoMessage()    {}
func (*LinkedDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{5}
}

func (m *LinkedDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkedDeviceResponse.Unmarshal(m, b)
}
func (m *LinkedDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkedDeviceResponse.Marshal(b, m, deterministic)
}
func (m *LinkedDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkedDeviceResponse.Merge(m, src)
}
func (m *LinkedDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_LinkedDeviceResponse.Size(m)
}
func (m *LinkedDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkedDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LinkedDeviceResponse proto.InternalMessageInfo

func (m *LinkedDeviceResponse) GetIsLinkedDeviceFetched() bool {
	if m != nil {
		return m.IsLinkedDeviceFetched
	}
	return false
}

func (m *LinkedDeviceResponse) GetLinkedDevices() []*LinkedDevice {
	if m != nil {
		return m.LinkedDevices
	}
	return nil
}

type TvDevice struct {
	GoogleId             string        `protobuf:"bytes,1,opt,name=googleId,proto3" json:"googleId,omitempty"`
	LinkedDevice         *LinkedDevice `protobuf:"bytes,2,opt,name=linkedDevice,proto3" json:"linkedDevice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TvDevice) Reset()         { *m = TvDevice{} }
func (m *TvDevice) String() string { return proto.CompactTextString(m) }
func (*TvDevice) ProtoMessage()    {}
func (*TvDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{6}
}

func (m *TvDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TvDevice.Unmarshal(m, b)
}
func (m *TvDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TvDevice.Marshal(b, m, deterministic)
}
func (m *TvDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TvDevice.Merge(m, src)
}
func (m *TvDevice) XXX_Size() int {
	return xxx_messageInfo_TvDevice.Size(m)
}
func (m *TvDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_TvDevice.DiscardUnknown(m)
}

var xxx_messageInfo_TvDevice proto.InternalMessageInfo

func (m *TvDevice) GetGoogleId() string {
	if m != nil {
		return m.GoogleId
	}
	return ""
}

func (m *TvDevice) GetLinkedDevice() *LinkedDevice {
	if m != nil {
		return m.LinkedDevice
	}
	return nil
}

type UpdateResponse struct {
	IsUpdated            bool     `protobuf:"varint,1,opt,name=isUpdated,proto3" json:"isUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetIsUpdated() bool {
	if m != nil {
		return m.IsUpdated
	}
	return false
}

type CreateReponse struct {
	IsCreated            bool     `protobuf:"varint,1,opt,name=isCreated,proto3" json:"isCreated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReponse) Reset()         { *m = CreateReponse{} }
func (m *CreateReponse) String() string { return proto.CompactTextString(m) }
func (*CreateReponse) ProtoMessage()    {}
func (*CreateReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{8}
}

func (m *CreateReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReponse.Unmarshal(m, b)
}
func (m *CreateReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReponse.Marshal(b, m, deterministic)
}
func (m *CreateReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReponse.Merge(m, src)
}
func (m *CreateReponse) XXX_Size() int {
	return xxx_messageInfo_CreateReponse.Size(m)
}
func (m *CreateReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReponse proto.InternalMessageInfo

func (m *CreateReponse) GetIsCreated() bool {
	if m != nil {
		return m.IsCreated
	}
	return false
}

// Date dd/MM/YYYY
type User struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	GoogleId             string               `protobuf:"bytes,4,opt,name=googleId,proto3" json:"googleId,omitempty"`
	DateOfBirth          string               `protobuf:"bytes,5,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	Gender               Gender               `protobuf:"varint,12,opt,name=gender,proto3,enum=User.Gender" json:"gender,omitempty"`
	ImageUrl             string               `protobuf:"bytes,13,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Genre                []string             `protobuf:"bytes,6,rep,name=genre,proto3" json:"genre,omitempty"`
	Language             []string             `protobuf:"bytes,7,rep,name=language,proto3" json:"language,omitempty"`
	ContentType          []string             `protobuf:"bytes,8,rep,name=contentType,proto3" json:"contentType,omitempty"`
	LinkedDevices        []*LinkedDevice      `protobuf:"bytes,9,rep,name=linkedDevices,proto3" json:"linkedDevices,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{9}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetGoogleId() string {
	if m != nil {
		return m.GoogleId
	}
	return ""
}

func (m *User) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *User) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_NOTDEFINED
}

func (m *User) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *User) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *User) GetLanguage() []string {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *User) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *User) GetLinkedDevices() []*LinkedDevice {
	if m != nil {
		return m.LinkedDevices
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type LinkedDevice struct {
	TvEmac               string   `protobuf:"bytes,1,opt,name=tvEmac,proto3" json:"tvEmac,omitempty"`
	TvPanel              string   `protobuf:"bytes,2,opt,name=tvPanel,proto3" json:"tvPanel,omitempty"`
	TvBoard              string   `protobuf:"bytes,3,opt,name=tvBoard,proto3" json:"tvBoard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkedDevice) Reset()         { *m = LinkedDevice{} }
func (m *LinkedDevice) String() string { return proto.CompactTextString(m) }
func (*LinkedDevice) ProtoMessage()    {}
func (*LinkedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{10}
}

func (m *LinkedDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkedDevice.Unmarshal(m, b)
}
func (m *LinkedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkedDevice.Marshal(b, m, deterministic)
}
func (m *LinkedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkedDevice.Merge(m, src)
}
func (m *LinkedDevice) XXX_Size() int {
	return xxx_messageInfo_LinkedDevice.Size(m)
}
func (m *LinkedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_LinkedDevice proto.InternalMessageInfo

func (m *LinkedDevice) GetTvEmac() string {
	if m != nil {
		return m.TvEmac
	}
	return ""
}

func (m *LinkedDevice) GetTvPanel() string {
	if m != nil {
		return m.TvPanel
	}
	return ""
}

func (m *LinkedDevice) GetTvBoard() string {
	if m != nil {
		return m.TvBoard
	}
	return ""
}

func init() {
	proto.RegisterEnum("User.Gender", Gender_name, Gender_value)
	proto.RegisterType((*RemoveTvDeviceRequest)(nil), "User.RemoveTvDeviceRequest")
	proto.RegisterType((*RemoveTvDeviceResponse)(nil), "User.RemoveTvDeviceResponse")
	proto.RegisterType((*GetRequest)(nil), "User.GetRequest")
	proto.RegisterType((*DeleteRequest)(nil), "User.DeleteRequest")
	proto.RegisterType((*DeleteReponse)(nil), "User.DeleteReponse")
	proto.RegisterType((*LinkedDeviceResponse)(nil), "User.LinkedDeviceResponse")
	proto.RegisterType((*TvDevice)(nil), "User.TvDevice")
	proto.RegisterType((*UpdateResponse)(nil), "User.UpdateResponse")
	proto.RegisterType((*CreateReponse)(nil), "User.CreateReponse")
	proto.RegisterType((*User)(nil), "User.User")
	proto.RegisterType((*LinkedDevice)(nil), "User.LinkedDevice")
}

func init() {
	proto.RegisterFile("User.proto", fileDescriptor_979821478719c248)
}

var fileDescriptor_979821478719c248 = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6e, 0xda, 0x4e,
	0x10, 0xfe, 0x19, 0x08, 0x81, 0xe1, 0xcf, 0x8f, 0x6c, 0x08, 0xb2, 0x5c, 0xa4, 0xa2, 0x55, 0x0f,
	0x28, 0x6d, 0x41, 0x4a, 0xab, 0x2a, 0x8a, 0x7a, 0x49, 0xca, 0x1f, 0x45, 0x4d, 0x49, 0x45, 0x89,
	0x54, 0xf5, 0x10, 0xc9, 0x81, 0x89, 0x63, 0x15, 0xdb, 0xd4, 0x5e, 0x90, 0x7a, 0xed, 0xa1, 0x2f,
	0xd0, 0xe7, 0xea, 0xa9, 0x7d, 0x84, 0x3e, 0x48, 0xe5, 0x1d, 0x2f, 0xd8, 0x80, 0xd2, 0xdc, 0x76,
	0xbe, 0x99, 0xf9, 0x76, 0xf6, 0x9b, 0xf1, 0x18, 0xe0, 0x2a, 0x40, 0xbf, 0x35, 0xf3, 0x3d, 0xe1,
	0xb1, 0x4c, 0x78, 0x36, 0x1e, 0x5b, 0x9e, 0x67, 0x4d, 0xb1, 0x2d, 0xb1, 0x9b, 0xf9, 0x6d, 0x5b,
	0xd8, 0x0e, 0x06, 0xc2, 0x74, 0x66, 0x14, 0x66, 0xd4, 0xa3, 0x00, 0x73, 0x66, 0xb7, 0x4d, 0xd7,
	0xf5, 0x84, 0x29, 0x6c, 0xcf, 0x0d, 0xc8, 0xcb, 0xdf, 0xc2, 0xc1, 0x10, 0x1d, 0x6f, 0x81, 0xa3,
	0x45, 0x07, 0x17, 0xf6, 0x18, 0x87, 0xf8, 0x65, 0x8e, 0x81, 0x60, 0x06, 0xe4, 0x28, 0xf1, 0x7c,
	0xa2, 0x6b, 0x0d, 0xad, 0x99, 0x1f, 0x2e, 0x6d, 0x56, 0x83, 0xac, 0x58, 0x74, 0x1d, 0x73, 0xac,
	0xa7, 0xa4, 0x27, 0xb2, 0x78, 0x0f, 0x6a, 0xeb, 0x64, 0xc1, 0xcc, 0x73, 0x03, 0x64, 0xcf, 0x60,
	0xcf, 0x0e, 0x56, 0x68, 0x18, 0x43, 0xb4, 0xb9, 0xe1, 0xa6, 0x83, 0x37, 0x01, 0xfa, 0x28, 0x1e,
	0x50, 0x09, 0x7f, 0x0a, 0xa5, 0x0e, 0x4e, 0x51, 0x3c, 0xa4, 0x6c, 0xfe, 0x7c, 0x15, 0x4c, 0x55,
	0xd5, 0x21, 0x6f, 0x07, 0x04, 0xa9, 0x6a, 0x56, 0x00, 0xff, 0xae, 0x41, 0xf5, 0xc2, 0x76, 0x3f,
	0xe3, 0x64, 0xed, 0x31, 0x2f, 0xe1, 0xc0, 0x0e, 0xe2, 0x9e, 0x1e, 0x8a, 0xf1, 0xdd, 0x92, 0x62,
	0xbb, 0x93, 0x1d, 0x43, 0x69, 0x1a, 0x83, 0x03, 0x3d, 0xd5, 0x48, 0x37, 0x0b, 0x47, 0xac, 0x25,
	0x5b, 0x9a, 0xb8, 0x28, 0x19, 0xc8, 0xaf, 0x21, 0xa7, 0x14, 0xba, 0xb7, 0x2d, 0xaf, 0xa0, 0x18,
	0x4f, 0x94, 0xcd, 0xd9, 0x7e, 0x41, 0x22, 0x8e, 0xb7, 0xa0, 0x7c, 0x35, 0x9b, 0x98, 0x62, 0xf5,
	0x42, 0x29, 0x0c, 0x61, 0x31, 0x61, 0x22, 0x20, 0xd4, 0xf1, 0x8d, 0x8f, 0xe6, 0x9a, 0x8e, 0x04,
	0xc5, 0xc2, 0x23, 0x80, 0xff, 0x4e, 0x83, 0x1c, 0x55, 0xc6, 0x20, 0xe3, 0x9a, 0x0e, 0x46, 0x75,
	0xcb, 0x33, 0xab, 0xc2, 0x0e, 0x3a, 0xa6, 0x3d, 0x8d, 0x26, 0x89, 0x0c, 0xd6, 0x80, 0xc2, 0xec,
	0xce, 0x73, 0x71, 0x30, 0x77, 0x6e, 0xd0, 0xd7, 0xd3, 0xd2, 0x17, 0x87, 0x12, 0x3a, 0x64, 0xd6,
	0x74, 0x68, 0x40, 0x21, 0x2c, 0xf4, 0xf2, 0xf6, 0xcc, 0xf6, 0xc5, 0x9d, 0xbe, 0x43, 0xd9, 0x31,
	0x88, 0x3d, 0x81, 0xac, 0x85, 0xee, 0x04, 0x7d, 0xbd, 0xd8, 0xd0, 0x9a, 0xe5, 0xa3, 0x22, 0x69,
	0xd4, 0x97, 0xd8, 0x30, 0xf2, 0x85, 0x77, 0xd8, 0x8e, 0x69, 0xe1, 0x95, 0x3f, 0xd5, 0x4b, 0x74,
	0x87, 0xb2, 0xc3, 0xba, 0x2d, 0x74, 0x7d, 0xd4, 0xb3, 0x8d, 0x74, 0x58, 0xb7, 0x34, 0xc2, 0x8c,
	0xa9, 0xe9, 0x5a, 0x73, 0xd3, 0x42, 0x7d, 0x57, 0x3a, 0x96, 0x76, 0x58, 0xd5, 0xd8, 0x73, 0x05,
	0xba, 0x62, 0xf4, 0x75, 0x86, 0x7a, 0x4e, 0xba, 0xe3, 0xd0, 0xe6, 0x84, 0xe4, 0x1f, 0x38, 0x21,
	0xec, 0x18, 0xf2, 0x63, 0x52, 0xfb, 0x54, 0xe8, 0x20, 0xdb, 0x6e, 0xb4, 0x48, 0x8f, 0x96, 0x5a,
	0x0c, 0xad, 0x91, 0x5a, 0x0c, 0xc3, 0x55, 0x70, 0x98, 0x39, 0xa7, 0xb6, 0x9e, 0x0a, 0xbd, 0xf0,
	0xef, 0xcc, 0x65, 0x30, 0xff, 0x04, 0xc5, 0x78, 0x49, 0xb1, 0xa5, 0xa0, 0xc5, 0x97, 0x02, 0xd3,
	0x61, 0x57, 0x2c, 0xde, 0x9b, 0x2e, 0xaa, 0x1e, 0x2b, 0x93, 0x3c, 0x67, 0x9e, 0xe9, 0x4f, 0xa2,
	0x0e, 0x2b, 0xf3, 0xb0, 0x05, 0x59, 0xea, 0x05, 0x2b, 0x03, 0x0c, 0x2e, 0x47, 0x9d, 0x6e, 0xef,
	0x7c, 0xd0, 0xed, 0x54, 0xfe, 0x63, 0x39, 0xc8, 0xbc, 0x3b, 0xbd, 0xe8, 0x56, 0x34, 0x06, 0x90,
	0xed, 0x75, 0xe5, 0x39, 0x75, 0xf4, 0x33, 0x03, 0x85, 0x50, 0xa4, 0x0f, 0xe8, 0xcb, 0x5a, 0xba,
	0x00, 0x34, 0x7d, 0x72, 0xee, 0x80, 0x04, 0x94, 0xeb, 0x72, 0x9f, 0xce, 0x89, 0xf9, 0xe5, 0xb5,
	0x6f, 0xbf, 0xfe, 0xfc, 0x48, 0x55, 0x78, 0xa1, 0x4d, 0xc2, 0xcc, 0x03, 0xf4, 0x4f, 0xb4, 0x43,
	0xd6, 0x03, 0xa0, 0x99, 0xdf, 0xa0, 0xa9, 0x46, 0xe7, 0xc4, 0x67, 0x13, 0xe3, 0x21, 0x99, 0x14,
	0xcf, 0x00, 0x80, 0x96, 0x8a, 0xe4, 0x89, 0x4a, 0x48, 0xec, 0x2d, 0x63, 0x0d, 0x5c, 0xe7, 0x9b,
	0x48, 0x5c, 0xf1, 0xbd, 0x86, 0xdd, 0x3e, 0x0a, 0x49, 0x56, 0x51, 0x93, 0xab, 0xd6, 0xa5, 0x11,
	0x2b, 0x93, 0xef, 0x4b, 0x82, 0x12, 0xcf, 0xb5, 0x2d, 0x14, 0x2a, 0xfb, 0x23, 0x94, 0xa9, 0x71,
	0xcb, 0xa5, 0x52, 0xa6, 0x14, 0x65, 0x1b, 0xc6, 0x96, 0x89, 0x53, 0x6f, 0x34, 0x24, 0x65, 0x95,
	0xff, 0xdf, 0xa6, 0x11, 0x14, 0x8b, 0x89, 0x0c, 0x08, 0x99, 0x2d, 0x28, 0x27, 0xf7, 0x3f, 0x7b,
	0x44, 0x4c, 0x5b, 0x7f, 0x31, 0x46, 0x7d, 0xbb, 0x73, 0xe3, 0x22, 0x5f, 0x06, 0xc4, 0x2f, 0xba,
	0x86, 0x4a, 0x1f, 0xc5, 0x45, 0xe2, 0x1b, 0xd8, 0x54, 0xe2, 0xbe, 0x67, 0xd4, 0x25, 0x7b, 0x8d,
	0xef, 0x85, 0xca, 0xd0, 0x4b, 0x88, 0x3e, 0x38, 0xd1, 0x0e, 0x6f, 0xb2, 0x72, 0xf4, 0x5f, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x3b, 0xfe, 0x23, 0x6f, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*CreateReponse, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UpdateResponse, error)
	DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReponse, error)
	GetUser(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
	LinkedTvDevice(ctx context.Context, in *TvDevice, opts ...grpc.CallOption) (*LinkedDeviceResponse, error)
	RemoveTvDevice(ctx context.Context, in *RemoveTvDeviceRequest, opts ...grpc.CallOption) (*RemoveTvDeviceResponse, error)
	GetLinkedDevices(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*LinkedDeviceResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*CreateReponse, error) {
	out := new(CreateReponse)
	err := c.cc.Invoke(ctx, "/User.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReponse, error) {
	out := new(DeleteReponse)
	err := c.cc.Invoke(ctx, "/User.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/User.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LinkedTvDevice(ctx context.Context, in *TvDevice, opts ...grpc.CallOption) (*LinkedDeviceResponse, error) {
	out := new(LinkedDeviceResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/LinkedTvDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveTvDevice(ctx context.Context, in *RemoveTvDeviceRequest, opts ...grpc.CallOption) (*RemoveTvDeviceResponse, error) {
	out := new(RemoveTvDeviceResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/RemoveTvDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetLinkedDevices(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*LinkedDeviceResponse, error) {
	out := new(LinkedDeviceResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/GetLinkedDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*CreateReponse, error)
	UpdateUser(context.Context, *User) (*UpdateResponse, error)
	DeleteUser(context.Context, *DeleteRequest) (*DeleteReponse, error)
	GetUser(context.Context, *GetRequest) (*User, error)
	LinkedTvDevice(context.Context, *TvDevice) (*LinkedDeviceResponse, error)
	RemoveTvDevice(context.Context, *RemoveTvDeviceRequest) (*RemoveTvDeviceResponse, error)
	GetLinkedDevices(context.Context, *GetRequest) (*LinkedDeviceResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *User) (*CreateReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *User) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *DeleteRequest) (*DeleteReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) LinkedTvDevice(ctx context.Context, req *TvDevice) (*LinkedDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkedTvDevice not implemented")
}
func (*UnimplementedUserServiceServer) RemoveTvDevice(ctx context.Context, req *RemoveTvDeviceRequest) (*RemoveTvDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTvDevice not implemented")
}
func (*UnimplementedUserServiceServer) GetLinkedDevices(ctx context.Context, req *GetRequest) (*LinkedDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedDevices not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LinkedTvDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TvDevice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LinkedTvDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/LinkedTvDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LinkedTvDevice(ctx, req.(*TvDevice))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveTvDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTvDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveTvDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/RemoveTvDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveTvDevice(ctx, req.(*RemoveTvDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetLinkedDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetLinkedDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/GetLinkedDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetLinkedDevices(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "LinkedTvDevice",
			Handler:    _UserService_LinkedTvDevice_Handler,
		},
		{
			MethodName: "RemoveTvDevice",
			Handler:    _UserService_RemoveTvDevice_Handler,
		},
		{
			MethodName: "GetLinkedDevices",
			Handler:    _UserService_GetLinkedDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "User.proto",
}
