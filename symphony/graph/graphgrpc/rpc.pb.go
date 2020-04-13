// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package graphgrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Tenant struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tenant) Reset()         { *m = Tenant{} }
func (m *Tenant) String() string { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()    {}
func (*Tenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Tenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenant.Unmarshal(m, b)
}
func (m *Tenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenant.Marshal(b, m, deterministic)
}
func (m *Tenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenant.Merge(m, src)
}
func (m *Tenant) XXX_Size() int {
	return xxx_messageInfo_Tenant.Size(m)
}
func (m *Tenant) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenant.DiscardUnknown(m)
}

var xxx_messageInfo_Tenant proto.InternalMessageInfo

func (m *Tenant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tenant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TenantList struct {
	Tenants              []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TenantList) Reset()         { *m = TenantList{} }
func (m *TenantList) String() string { return proto.CompactTextString(m) }
func (*TenantList) ProtoMessage()    {}
func (*TenantList) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *TenantList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantList.Unmarshal(m, b)
}
func (m *TenantList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantList.Marshal(b, m, deterministic)
}
func (m *TenantList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantList.Merge(m, src)
}
func (m *TenantList) XXX_Size() int {
	return xxx_messageInfo_TenantList.Size(m)
}
func (m *TenantList) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantList.DiscardUnknown(m)
}

var xxx_messageInfo_TenantList proto.InternalMessageInfo

func (m *TenantList) GetTenants() []*Tenant {
	if m != nil {
		return m.Tenants
	}
	return nil
}

type UserInput struct {
	Tenant               string   `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInput) Reset()         { *m = UserInput{} }
func (m *UserInput) String() string { return proto.CompactTextString(m) }
func (*UserInput) ProtoMessage()    {}
func (*UserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *UserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInput.Unmarshal(m, b)
}
func (m *UserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInput.Marshal(b, m, deterministic)
}
func (m *UserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInput.Merge(m, src)
}
func (m *UserInput) XXX_Size() int {
	return xxx_messageInfo_UserInput.Size(m)
}
func (m *UserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInput.DiscardUnknown(m)
}

var xxx_messageInfo_UserInput proto.InternalMessageInfo

func (m *UserInput) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *UserInput) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddUserInput struct {
	Tenant               string   `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	IsOwner              bool     `protobuf:"varint,3,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserInput) Reset()         { *m = AddUserInput{} }
func (m *AddUserInput) String() string { return proto.CompactTextString(m) }
func (*AddUserInput) ProtoMessage()    {}
func (*AddUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *AddUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserInput.Unmarshal(m, b)
}
func (m *AddUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserInput.Marshal(b, m, deterministic)
}
func (m *AddUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserInput.Merge(m, src)
}
func (m *AddUserInput) XXX_Size() int {
	return xxx_messageInfo_AddUserInput.Size(m)
}
func (m *AddUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserInput proto.InternalMessageInfo

func (m *AddUserInput) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *AddUserInput) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddUserInput) GetIsOwner() bool {
	if m != nil {
		return m.IsOwner
	}
	return false
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AlertPayload struct {
	TenantID             string            `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Alertname            string            `protobuf:"bytes,2,opt,name=alertname,proto3" json:"alertname,omitempty"`
	NetworkID            string            `protobuf:"bytes,3,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AlertPayload) Reset()         { *m = AlertPayload{} }
func (m *AlertPayload) String() string { return proto.CompactTextString(m) }
func (*AlertPayload) ProtoMessage()    {}
func (*AlertPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *AlertPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPayload.Unmarshal(m, b)
}
func (m *AlertPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPayload.Marshal(b, m, deterministic)
}
func (m *AlertPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPayload.Merge(m, src)
}
func (m *AlertPayload) XXX_Size() int {
	return xxx_messageInfo_AlertPayload.Size(m)
}
func (m *AlertPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPayload.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPayload proto.InternalMessageInfo

func (m *AlertPayload) GetTenantID() string {
	if m != nil {
		return m.TenantID
	}
	return ""
}

func (m *AlertPayload) GetAlertname() string {
	if m != nil {
		return m.Alertname
	}
	return ""
}

func (m *AlertPayload) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *AlertPayload) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ExecutionResult struct {
	Successes            []string          `protobuf:"bytes,1,rep,name=successes,proto3" json:"successes,omitempty"`
	Errors               []*ExecutionError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExecutionResult) Reset()         { *m = ExecutionResult{} }
func (m *ExecutionResult) String() string { return proto.CompactTextString(m) }
func (*ExecutionResult) ProtoMessage()    {}
func (*ExecutionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *ExecutionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionResult.Unmarshal(m, b)
}
func (m *ExecutionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionResult.Marshal(b, m, deterministic)
}
func (m *ExecutionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionResult.Merge(m, src)
}
func (m *ExecutionResult) XXX_Size() int {
	return xxx_messageInfo_ExecutionResult.Size(m)
}
func (m *ExecutionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionResult proto.InternalMessageInfo

func (m *ExecutionResult) GetSuccesses() []string {
	if m != nil {
		return m.Successes
	}
	return nil
}

func (m *ExecutionResult) GetErrors() []*ExecutionError {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ExecutionError struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionError) Reset()         { *m = ExecutionError{} }
func (m *ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionError) ProtoMessage()    {}
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *ExecutionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionError.Unmarshal(m, b)
}
func (m *ExecutionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionError.Marshal(b, m, deterministic)
}
func (m *ExecutionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionError.Merge(m, src)
}
func (m *ExecutionError) XXX_Size() int {
	return xxx_messageInfo_ExecutionError.Size(m)
}
func (m *ExecutionError) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionError.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionError proto.InternalMessageInfo

func (m *ExecutionError) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExecutionError) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Tenant)(nil), "graph.Tenant")
	proto.RegisterType((*TenantList)(nil), "graph.TenantList")
	proto.RegisterType((*UserInput)(nil), "graph.UserInput")
	proto.RegisterType((*AddUserInput)(nil), "graph.AddUserInput")
	proto.RegisterType((*User)(nil), "graph.User")
	proto.RegisterType((*AlertPayload)(nil), "graph.AlertPayload")
	proto.RegisterMapType((map[string]string)(nil), "graph.AlertPayload.LabelsEntry")
	proto.RegisterType((*ExecutionResult)(nil), "graph.ExecutionResult")
	proto.RegisterType((*ExecutionError)(nil), "graph.executionError")
}

func init() {
	proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1)
}

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5f, 0x6f, 0x94, 0x4e,
	0x14, 0x0d, 0xb0, 0xa5, 0xe5, 0x6e, 0xdb, 0x5f, 0x7f, 0x53, 0x6d, 0x08, 0x36, 0xb5, 0xe1, 0xc5,
	0xc6, 0x28, 0x1b, 0x69, 0xea, 0xbf, 0x07, 0xe3, 0xea, 0x6e, 0xcc, 0x26, 0x35, 0x36, 0x74, 0xf5,
	0xc1, 0x07, 0x13, 0x96, 0xbd, 0xa5, 0xa4, 0x2c, 0x90, 0x99, 0xa1, 0x2b, 0x9f, 0xca, 0xef, 0xe3,
	0xa7, 0x31, 0x33, 0x03, 0xbb, 0x6c, 0x9b, 0xc6, 0xd4, 0x17, 0x32, 0xf7, 0xdc, 0x33, 0xf7, 0x1c,
	0x2e, 0x07, 0xb0, 0x68, 0x11, 0x79, 0x05, 0xcd, 0x79, 0x4e, 0xd6, 0x62, 0x1a, 0x16, 0x97, 0xce,
	0xa3, 0x38, 0xcf, 0xe3, 0x14, 0x7b, 0x12, 0x9c, 0x94, 0x17, 0x3d, 0x9c, 0x15, 0xbc, 0x52, 0x1c,
	0xe7, 0xe0, 0x66, 0x73, 0x4e, 0xc3, 0xa2, 0x40, 0xca, 0x54, 0xdf, 0x7d, 0x06, 0xe6, 0x18, 0xb3,
	0x30, 0xe3, 0x64, 0x1b, 0xf4, 0x64, 0x6a, 0x6b, 0x87, 0xda, 0x91, 0x15, 0xe8, 0xc9, 0x94, 0x10,
	0xe8, 0x64, 0xe1, 0x0c, 0x6d, 0x5d, 0x22, 0xf2, 0xec, 0x9e, 0x00, 0x28, 0xf6, 0x69, 0xc2, 0x38,
	0x79, 0x02, 0xeb, 0x5c, 0x56, 0xcc, 0xd6, 0x0e, 0x8d, 0xa3, 0xae, 0xbf, 0xe5, 0x49, 0x47, 0x9e,
	0xe2, 0x04, 0x4d, 0xd7, 0x3d, 0x06, 0xeb, 0x2b, 0x43, 0x3a, 0xca, 0x8a, 0x92, 0x93, 0x3d, 0x30,
	0x15, 0x5e, 0x6b, 0xd5, 0x55, 0xad, 0xaf, 0x37, 0xfa, 0xee, 0x19, 0x6c, 0xf6, 0xa7, 0xd3, 0x7b,
	0xdf, 0x23, 0x36, 0xac, 0x27, 0xec, 0xcb, 0x3c, 0x43, 0x6a, 0x1b, 0x87, 0xda, 0xd1, 0x46, 0xd0,
	0x94, 0xee, 0x1e, 0x74, 0xc4, 0xb8, 0xd6, 0x9b, 0x1a, 0x52, 0xe9, 0xb7, 0x06, 0x9b, 0xfd, 0x14,
	0x29, 0x3f, 0x0b, 0xab, 0x34, 0x0f, 0xa7, 0xc4, 0x81, 0x0d, 0x35, 0x7c, 0x34, 0xa8, 0xc5, 0x16,
	0x35, 0xd9, 0x07, 0x2b, 0x14, 0xdc, 0xd6, 0x6e, 0x96, 0x80, 0xe8, 0x66, 0xc8, 0xe7, 0x39, 0xbd,
	0x1a, 0x0d, 0xa4, 0xbc, 0x15, 0x2c, 0x01, 0xf2, 0x0a, 0xcc, 0x34, 0x9c, 0x60, 0xca, 0xec, 0x8e,
	0xdc, 0xd7, 0xe3, 0x7a, 0x5f, 0x6d, 0x71, 0xef, 0x54, 0x32, 0x86, 0x19, 0xa7, 0x55, 0x50, 0xd3,
	0x9d, 0x37, 0xd0, 0x6d, 0xc1, 0x64, 0x07, 0x8c, 0x2b, 0xac, 0x6a, 0x6b, 0xe2, 0x48, 0x1e, 0xc0,
	0xda, 0x75, 0x98, 0x96, 0x8d, 0x23, 0x55, 0xbc, 0xd5, 0x5f, 0x6b, 0xee, 0x0f, 0xf8, 0x6f, 0xf8,
	0x13, 0xa3, 0x92, 0x27, 0x79, 0x16, 0x20, 0x2b, 0x53, 0x2e, 0x4c, 0xb2, 0x32, 0x8a, 0x90, 0x31,
	0x54, 0x5f, 0xce, 0x0a, 0x96, 0x00, 0x79, 0x0e, 0x26, 0x52, 0x9a, 0x53, 0x66, 0xeb, 0xd2, 0xe4,
	0xc3, 0xda, 0x24, 0x36, 0x53, 0x86, 0xa2, 0x1b, 0xd4, 0x24, 0xd7, 0x87, 0xed, 0xd5, 0xce, 0xad,
	0x20, 0xed, 0x80, 0x81, 0x94, 0xd6, 0xce, 0xc4, 0xd1, 0xff, 0xa5, 0xc3, 0x96, 0xca, 0xc8, 0x39,
	0xd2, 0xeb, 0x24, 0x42, 0x72, 0x02, 0xe6, 0x47, 0x8a, 0x21, 0x47, 0xb2, 0xef, 0xa9, 0xc4, 0x7a,
	0x4d, 0x62, 0xbd, 0x73, 0x4e, 0x93, 0x2c, 0xfe, 0x26, 0xde, 0xc8, 0x59, 0x4d, 0x18, 0x79, 0x01,
	0x1d, 0x99, 0xc4, 0xbd, 0x5b, 0x97, 0x86, 0xe2, 0x1f, 0x70, 0xfe, 0x5f, 0xa1, 0x4b, 0xaa, 0x0f,
	0xc6, 0x27, 0xe4, 0xf7, 0x93, 0x79, 0x0f, 0x1b, 0x63, 0x5a, 0x66, 0xd1, 0xdf, 0xfd, 0xdd, 0x61,
	0x84, 0xbc, 0x03, 0x73, 0x80, 0x29, 0xfe, 0xeb, 0x7d, 0x7f, 0x06, 0x5d, 0x11, 0xdd, 0x66, 0x5d,
	0x4f, 0x17, 0xeb, 0xda, 0x6d, 0x22, 0xd4, 0xfa, 0x55, 0x9c, 0x6e, 0x0d, 0xca, 0xb4, 0xfb, 0x0b,
	0xe9, 0x9d, 0x16, 0xac, 0x88, 0x77, 0xc9, 0x7d, 0x86, 0xdd, 0x7e, 0x24, 0xbe, 0x28, 0x93, 0xd1,
	0x6c, 0x64, 0x5f, 0xc2, 0xfa, 0x98, 0x26, 0x71, 0x8c, 0x74, 0xa9, 0xdb, 0x8a, 0xae, 0x18, 0x27,
	0xc1, 0x1b, 0x81, 0xfb, 0x70, 0xf0, 0x7d, 0xff, 0x62, 0x12, 0xf5, 0x58, 0x35, 0x2b, 0x2e, 0xf3,
	0xac, 0xea, 0x49, 0x96, 0x7a, 0xc6, 0xb4, 0x88, 0x26, 0xa6, 0x94, 0x3f, 0xfe, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x1c, 0x2c, 0xb9, 0xdc, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantServiceClient interface {
	Create(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TenantList, error)
	Get(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error)
	Truncate(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) Create(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TenantList, error) {
	out := new(TenantList)
	err := c.cc.Invoke(ctx, "/graph.TenantService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Get(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Truncate(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Delete(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
type TenantServiceServer interface {
	Create(context.Context, *wrappers.StringValue) (*Tenant, error)
	List(context.Context, *empty.Empty) (*TenantList, error)
	Get(context.Context, *wrappers.StringValue) (*Tenant, error)
	Truncate(context.Context, *wrappers.StringValue) (*empty.Empty, error)
	Delete(context.Context, *wrappers.StringValue) (*empty.Empty, error)
}

// UnimplementedTenantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (*UnimplementedTenantServiceServer) Create(ctx context.Context, req *wrappers.StringValue) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTenantServiceServer) List(ctx context.Context, req *empty.Empty) (*TenantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTenantServiceServer) Get(ctx context.Context, req *wrappers.StringValue) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTenantServiceServer) Truncate(ctx context.Context, req *wrappers.StringValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Truncate not implemented")
}
func (*UnimplementedTenantServiceServer) Delete(ctx context.Context, req *wrappers.StringValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTenantServiceServer(s *grpc.Server, srv TenantServiceServer) {
	s.RegisterService(&_TenantService_serviceDesc, srv)
}

func _TenantService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Create(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Get(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Truncate(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Delete(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TenantService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TenantService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TenantService_Get_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _TenantService_Truncate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TenantService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *AddUserInput, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *AddUserInput, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/graph.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *AddUserInput) (*User, error)
	Delete(context.Context, *UserInput) (*empty.Empty, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *AddUserInput) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Delete(ctx context.Context, req *UserInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*AddUserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// ActionsAlertServiceClient is the client API for ActionsAlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActionsAlertServiceClient interface {
	Trigger(ctx context.Context, in *AlertPayload, opts ...grpc.CallOption) (*ExecutionResult, error)
}

type actionsAlertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsAlertServiceClient(cc grpc.ClientConnInterface) ActionsAlertServiceClient {
	return &actionsAlertServiceClient{cc}
}

func (c *actionsAlertServiceClient) Trigger(ctx context.Context, in *AlertPayload, opts ...grpc.CallOption) (*ExecutionResult, error) {
	out := new(ExecutionResult)
	err := c.cc.Invoke(ctx, "/graph.ActionsAlertService/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsAlertServiceServer is the server API for ActionsAlertService service.
type ActionsAlertServiceServer interface {
	Trigger(context.Context, *AlertPayload) (*ExecutionResult, error)
}

// UnimplementedActionsAlertServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActionsAlertServiceServer struct {
}

func (*UnimplementedActionsAlertServiceServer) Trigger(ctx context.Context, req *AlertPayload) (*ExecutionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}

func RegisterActionsAlertServiceServer(s *grpc.Server, srv ActionsAlertServiceServer) {
	s.RegisterService(&_ActionsAlertService_serviceDesc, srv)
}

func _ActionsAlertService_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAlertServiceServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.ActionsAlertService/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAlertServiceServer).Trigger(ctx, req.(*AlertPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionsAlertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.ActionsAlertService",
	HandlerType: (*ActionsAlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _ActionsAlertService_Trigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}