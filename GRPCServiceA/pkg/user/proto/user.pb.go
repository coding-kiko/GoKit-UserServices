// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetUserReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (m *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(m, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

func (m *GetUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Nationality          string   `protobuf:"bytes,4,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Job                  string   `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	Created              string   `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Error                *Status  `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetUserResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserResp) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *GetUserResp) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *GetUserResp) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *GetUserResp) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *GetUserResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetUserResp) GetError() *Status {
	if m != nil {
		return m.Error
	}
	return nil
}

type CreateUserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Nationality          string   `protobuf:"bytes,4,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Job                  string   `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserReq) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *CreateUserReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *CreateUserReq) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *CreateUserReq) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *CreateUserReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created              string   `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Error                *Status  `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (m *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(m, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

func (m *CreateUserResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateUserResp) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *CreateUserResp) GetError() *Status {
	if m != nil {
		return m.Error
	}
	return nil
}

type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteUserReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserReq) Reset()         { *m = DeleteUserReq{} }
func (m *DeleteUserReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUserReq) ProtoMessage()    {}
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *DeleteUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserReq.Unmarshal(m, b)
}
func (m *DeleteUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserReq.Marshal(b, m, deterministic)
}
func (m *DeleteUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserReq.Merge(m, src)
}
func (m *DeleteUserReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUserReq.Size(m)
}
func (m *DeleteUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserReq proto.InternalMessageInfo

func (m *DeleteUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteUserResp struct {
	Deleted              string   `protobuf:"bytes,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Error                *Status  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResp) Reset()         { *m = DeleteUserResp{} }
func (m *DeleteUserResp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResp) ProtoMessage()    {}
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *DeleteUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResp.Unmarshal(m, b)
}
func (m *DeleteUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResp.Marshal(b, m, deterministic)
}
func (m *DeleteUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResp.Merge(m, src)
}
func (m *DeleteUserResp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResp.Size(m)
}
func (m *DeleteUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResp proto.InternalMessageInfo

func (m *DeleteUserResp) GetDeleted() string {
	if m != nil {
		return m.Deleted
	}
	return ""
}

func (m *DeleteUserResp) GetError() *Status {
	if m != nil {
		return m.Error
	}
	return nil
}

type UpdateUserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Nationality          string   `protobuf:"bytes,4,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Job                  string   `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserReq) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UpdateUserReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *UpdateUserReq) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *UpdateUserReq) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *UpdateUserReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UpdateUserResp struct {
	Updated              string   `protobuf:"bytes,1,opt,name=updated,proto3" json:"updated,omitempty"`
	Error                *Status  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResp) Reset()         { *m = UpdateUserResp{} }
func (m *UpdateUserResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResp) ProtoMessage()    {}
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UpdateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResp.Unmarshal(m, b)
}
func (m *UpdateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResp.Merge(m, src)
}
func (m *UpdateUserResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResp.Size(m)
}
func (m *UpdateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResp proto.InternalMessageInfo

func (m *UpdateUserResp) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *UpdateUserResp) GetError() *Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserReq)(nil), "proto.GetUserReq")
	proto.RegisterType((*GetUserResp)(nil), "proto.GetUserResp")
	proto.RegisterType((*CreateUserReq)(nil), "proto.CreateUserReq")
	proto.RegisterType((*CreateUserResp)(nil), "proto.CreateUserResp")
	proto.RegisterType((*Status)(nil), "proto.Status")
	proto.RegisterType((*DeleteUserReq)(nil), "proto.DeleteUserReq")
	proto.RegisterType((*DeleteUserResp)(nil), "proto.DeleteUserResp")
	proto.RegisterType((*UpdateUserReq)(nil), "proto.UpdateUserReq")
	proto.RegisterType((*UpdateUserResp)(nil), "proto.UpdateUserResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x37, 0xc9, 0xa6, 0xa5, 0x53, 0x5a, 0x2d, 0xd6, 0x22, 0x59, 0x2b, 0x24, 0xaa, 0x70,
	0xe9, 0x81, 0x6d, 0xa5, 0x22, 0x71, 0xe1, 0x04, 0x45, 0xf4, 0xb8, 0x2b, 0xaf, 0xf6, 0xc2, 0x05,
	0xb9, 0xcd, 0x80, 0xbc, 0x4a, 0xeb, 0x60, 0xbb, 0xfc, 0x79, 0x0c, 0xde, 0x8b, 0x17, 0xe2, 0x86,
	0xec, 0xc4, 0x4e, 0x52, 0x2a, 0x40, 0x9c, 0x38, 0x65, 0xe6, 0xb3, 0xbf, 0xd1, 0xfc, 0x66, 0x1c,
	0x80, 0xbd, 0x46, 0x35, 0x2b, 0x95, 0x34, 0x92, 0xa4, 0xee, 0x93, 0x3d, 0x02, 0x58, 0xa1, 0xb9,
	0xd5, 0xa8, 0x18, 0x7e, 0x24, 0x63, 0x88, 0x45, 0x4e, 0xa3, 0x49, 0x34, 0x1d, 0xb0, 0x58, 0xe4,
	0xd9, 0xf7, 0x08, 0x86, 0xe1, 0x58, 0x97, 0x87, 0xe7, 0x84, 0xc0, 0xe9, 0x8e, 0x6f, 0x91, 0xc6,
	0x4e, 0x71, 0x31, 0x39, 0x83, 0x84, 0x7f, 0x40, 0x9a, 0x4c, 0xa2, 0x69, 0xc2, 0x6c, 0x48, 0x26,
	0x30, 0xdc, 0x71, 0x23, 0xe4, 0x8e, 0x17, 0xc2, 0x7c, 0xa5, 0xa7, 0xee, 0x72, 0x5b, 0xb2, 0x9e,
	0x3b, 0xb9, 0xa6, 0xa9, 0x3b, 0xb1, 0x21, 0xa1, 0xd0, 0xdf, 0x28, 0xe4, 0x06, 0x73, 0xda, 0x73,
	0xaa, 0x4f, 0xc9, 0x39, 0xa4, 0xb8, 0xe5, 0xa2, 0xa0, 0x7d, 0xa7, 0x57, 0x09, 0x79, 0x02, 0x29,
	0x2a, 0x25, 0x15, 0xbd, 0x37, 0x89, 0xa6, 0xc3, 0xc5, 0xa8, 0xa2, 0x9c, 0xdd, 0x18, 0x6e, 0xf6,
	0x9a, 0x55, 0x67, 0xd9, 0xb7, 0x08, 0x46, 0x4b, 0x57, 0xc6, 0x03, 0x7b, 0x80, 0xe8, 0x57, 0x80,
	0xb8, 0x01, 0x38, 0x83, 0xa4, 0xfc, 0x9c, 0x3b, 0xa4, 0x01, 0xb3, 0xe1, 0x3f, 0x21, 0x85, 0xc6,
	0x7b, 0xad, 0xc6, 0xb3, 0x77, 0x30, 0x6e, 0xb7, 0x74, 0x64, 0xc8, 0xad, 0x51, 0xc4, 0xdd, 0x51,
	0x04, 0xe8, 0xe4, 0x37, 0xd0, 0xcf, 0xa1, 0x57, 0x09, 0x16, 0x76, 0x23, 0xf3, 0x0a, 0x36, 0x65,
	0x2e, 0xb6, 0xc5, 0xb7, 0xa8, 0xb5, 0x07, 0x1e, 0x30, 0x9f, 0x66, 0x8f, 0x61, 0xf4, 0x1a, 0x0b,
	0x6c, 0x66, 0x75, 0xf8, 0x38, 0xae, 0x60, 0xdc, 0xbe, 0xa0, 0x4b, 0x5b, 0x2c, 0x77, 0x8a, 0xbf,
	0xe6, 0xd3, 0xa6, 0xd3, 0xf8, 0x0f, 0xeb, 0xb9, 0x2d, 0xf3, 0xff, 0x6a, 0x3d, 0x57, 0x30, 0x6e,
	0xb7, 0x54, 0x41, 0xee, 0x9d, 0x12, 0x20, 0xeb, 0xf4, 0xaf, 0x20, 0x17, 0x3f, 0x22, 0xb8, 0x6f,
	0x6b, 0xdd, 0xa0, 0xfa, 0x24, 0x36, 0xa8, 0xc9, 0x02, 0xfa, 0xf5, 0x2f, 0x46, 0x1e, 0xd4, 0x8e,
	0xe6, 0x8f, 0xbc, 0x20, 0x87, 0x92, 0x2e, 0xb3, 0x13, 0xf2, 0x02, 0xa0, 0x79, 0x34, 0xe4, 0xbc,
	0xbe, 0xd3, 0x79, 0xda, 0x17, 0x0f, 0x8f, 0xa8, 0xde, 0xdc, 0xec, 0x2d, 0x98, 0x3b, 0xbb, 0x0e,
	0xe6, 0xee, 0x82, 0x2b, 0x73, 0x33, 0x8f, 0x60, 0xee, 0x6c, 0x2d, 0x98, 0xbb, 0x83, 0xcb, 0x4e,
	0x5e, 0xcd, 0xde, 0x3e, 0x7d, 0xbf, 0xe4, 0x85, 0xf8, 0x62, 0xe4, 0xe5, 0x6a, 0x3d, 0x7f, 0x23,
	0x76, 0xbc, 0xb8, 0xbc, 0x56, 0xf2, 0x0e, 0x37, 0x66, 0xbe, 0x62, 0xd7, 0xcb, 0x7a, 0x28, 0x2f,
	0xe7, 0xae, 0xc0, 0xba, 0xe7, 0x3e, 0xcf, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x46, 0xcb, 0xc8,
	0xc0, 0xb8, 0x04, 0x00, 0x00,
}
