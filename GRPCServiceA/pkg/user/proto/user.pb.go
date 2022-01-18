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

func init() {
	proto.RegisterType((*GetUserReq)(nil), "proto.GetUserReq")
	proto.RegisterType((*GetUserResp)(nil), "proto.GetUserResp")
	proto.RegisterType((*CreateUserReq)(nil), "proto.CreateUserReq")
	proto.RegisterType((*CreateUserResp)(nil), "proto.CreateUserResp")
	proto.RegisterType((*Status)(nil), "proto.Status")
	proto.RegisterType((*DeleteUserReq)(nil), "proto.DeleteUserReq")
	proto.RegisterType((*DeleteUserResp)(nil), "proto.DeleteUserResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xd7, 0xc9, 0xa6, 0xa5, 0x53, 0x5a, 0x2d, 0xd6, 0x22, 0x59, 0x2b, 0x24, 0xaa, 0x70,
	0xe9, 0x81, 0x6d, 0xa5, 0x22, 0x71, 0xe1, 0x04, 0x45, 0xf4, 0xc8, 0xca, 0x2b, 0x2e, 0x5c, 0x90,
	0x9b, 0x0c, 0xc8, 0xab, 0x34, 0x0e, 0xb6, 0x97, 0x8f, 0xc7, 0xe0, 0x7d, 0x78, 0x04, 0x1e, 0x0a,
	0xd9, 0x89, 0xf3, 0x51, 0x16, 0x0e, 0x9c, 0x32, 0xf3, 0x9f, 0x0f, 0xcd, 0xef, 0x1f, 0x03, 0xdc,
	0x1a, 0xd4, 0xab, 0x4a, 0x2b, 0xab, 0x68, 0xe2, 0x3f, 0xe9, 0x23, 0x80, 0x1d, 0xda, 0x77, 0x06,
	0x35, 0xc7, 0xcf, 0x74, 0x0e, 0x91, 0xcc, 0x19, 0x59, 0x90, 0xe5, 0x84, 0x47, 0x32, 0x4f, 0x7f,
	0x11, 0x98, 0xb6, 0x65, 0x53, 0x1d, 0xd7, 0x29, 0x85, 0xd3, 0x52, 0x1c, 0x90, 0x45, 0x5e, 0xf1,
	0x31, 0x3d, 0x83, 0x58, 0x7c, 0x42, 0x16, 0x2f, 0xc8, 0x32, 0xe6, 0x2e, 0xa4, 0x0b, 0x98, 0x96,
	0xc2, 0x4a, 0x55, 0x8a, 0x42, 0xda, 0xef, 0xec, 0xd4, 0x37, 0xf7, 0x25, 0x37, 0x73, 0xa3, 0xf6,
	0x2c, 0xf1, 0x15, 0x17, 0x52, 0x06, 0xe3, 0x4c, 0xa3, 0xb0, 0x98, 0xb3, 0x91, 0x57, 0x43, 0x4a,
	0xcf, 0x21, 0xc1, 0x83, 0x90, 0x05, 0x1b, 0x7b, 0xbd, 0x4e, 0xe8, 0x13, 0x48, 0x50, 0x6b, 0xa5,
	0xd9, 0xbd, 0x05, 0x59, 0x4e, 0x37, 0xb3, 0x9a, 0x72, 0x75, 0x6d, 0x85, 0xbd, 0x35, 0xbc, 0xae,
	0xa5, 0x3f, 0x08, 0xcc, 0xb6, 0x7e, 0x4d, 0x00, 0x0e, 0x00, 0xe4, 0x4f, 0x80, 0xa8, 0x03, 0x38,
	0x83, 0xb8, 0xfa, 0x9a, 0x7b, 0xa4, 0x09, 0x77, 0xe1, 0x7f, 0x21, 0xb5, 0x87, 0x8f, 0x7a, 0x87,
	0xa7, 0x1f, 0x60, 0xde, 0x3f, 0xe9, 0x0e, 0x93, 0x7b, 0x56, 0x44, 0x43, 0x2b, 0x5a, 0xe8, 0xf8,
	0x1f, 0xd0, 0xcf, 0x61, 0x54, 0x0b, 0x0e, 0x36, 0x53, 0x79, 0x0d, 0x9b, 0x70, 0x1f, 0xbb, 0xe5,
	0x07, 0x34, 0x26, 0x00, 0x4f, 0x78, 0x48, 0xd3, 0xc7, 0x30, 0x7b, 0x8d, 0x05, 0x76, 0x5e, 0x1d,
	0x3f, 0x8e, 0xb7, 0x30, 0xef, 0x37, 0x98, 0xca, 0x2d, 0xcb, 0xbd, 0x12, 0xda, 0x42, 0xda, 0x5d,
	0x1a, 0xfd, 0xfd, 0xd2, 0xcd, 0x4f, 0x02, 0xf7, 0xdd, 0xae, 0x6b, 0xd4, 0x5f, 0x64, 0x86, 0x86,
	0x6e, 0x60, 0xdc, 0xbc, 0x3e, 0xfa, 0xa0, 0x99, 0xe8, 0x1e, 0xeb, 0x05, 0x3d, 0x96, 0x4c, 0x95,
	0x9e, 0xd0, 0x17, 0x00, 0x9d, 0x9f, 0xf4, 0xbc, 0xe9, 0x19, 0xfc, 0xf5, 0x8b, 0x87, 0x77, 0xa8,
	0x61, 0xb8, 0x43, 0x6a, 0x87, 0x07, 0x36, 0xb4, 0xc3, 0x43, 0xf6, 0xf4, 0xe4, 0xd5, 0xea, 0xfd,
	0xd3, 0x8f, 0x5b, 0x51, 0xc8, 0x6f, 0x56, 0x5d, 0xee, 0xf6, 0xeb, 0x37, 0xb2, 0x14, 0xc5, 0xe5,
	0x95, 0x56, 0x37, 0x98, 0xd9, 0xf5, 0x8e, 0x5f, 0x6d, 0x1b, 0xae, 0x97, 0x6b, 0xbf, 0x60, 0x3f,
	0xf2, 0x9f, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x62, 0xee, 0x42, 0x7d, 0x96, 0x03, 0x00,
	0x00,
}
