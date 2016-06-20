// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_DeletedStoreUser_Ad.proto
// DO NOT EDIT!

/*
Package mi_DeletedStoreUser_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_DeletedStoreUser_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_DeletedStoreUser_Ad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dstore_values "dstore/values"
import dstore_engine_error "dstore/engine/error"
import dstore_engine_message "dstore/engine/message"
import dstore_engine_metainformation "dstore/engine/metainformation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Parameters struct {
	UserName                          *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserNameNull                      bool                        `protobuf:"varint,1001,opt,name=user_name_null,json=userNameNull" json:"user_name_null,omitempty"`
	InvalidateUserInUserInfo          *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=invalidate_user_in_user_info,json=invalidateUserInUserInfo" json:"invalidate_user_in_user_info,omitempty"`
	InvalidateUserInUserInfoNull      bool                        `protobuf:"varint,1002,opt,name=invalidate_user_in_user_info_null,json=invalidateUserInUserInfoNull" json:"invalidate_user_in_user_info_null,omitempty"`
	OnlyDeleteUserFromThisEngine      *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=only_delete_user_from_this_engine,json=onlyDeleteUserFromThisEngine" json:"only_delete_user_from_this_engine,omitempty"`
	OnlyDeleteUserFromThisEngineNull  bool                        `protobuf:"varint,1003,opt,name=only_delete_user_from_this_engine_null,json=onlyDeleteUserFromThisEngineNull" json:"only_delete_user_from_this_engine_null,omitempty"`
	DropLoginWhenDBIsNotAvailable     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=drop_login_when_d_b_is_not_available,json=dropLoginWhenDBIsNotAvailable" json:"drop_login_when_d_b_is_not_available,omitempty"`
	DropLoginWhenDBIsNotAvailableNull bool                        `protobuf:"varint,1004,opt,name=drop_login_when_d_b_is_not_available_null,json=dropLoginWhenDBIsNotAvailableNull" json:"drop_login_when_d_b_is_not_available_null,omitempty"`
	KillProcessesForUserFirst         *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=kill_processes_for_user_first,json=killProcessesForUserFirst" json:"kill_processes_for_user_first,omitempty"`
	KillProcessesForUserFirstNull     bool                        `protobuf:"varint,1005,opt,name=kill_processes_for_user_first_null,json=killProcessesForUserFirstNull" json:"kill_processes_for_user_first_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Parameters) GetInvalidateUserInUserInfo() *dstore_values.BooleanValue {
	if m != nil {
		return m.InvalidateUserInUserInfo
	}
	return nil
}

func (m *Parameters) GetOnlyDeleteUserFromThisEngine() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyDeleteUserFromThisEngine
	}
	return nil
}

func (m *Parameters) GetDropLoginWhenDBIsNotAvailable() *dstore_values.BooleanValue {
	if m != nil {
		return m.DropLoginWhenDBIsNotAvailable
	}
	return nil
}

func (m *Parameters) GetKillProcessesForUserFirst() *dstore_values.BooleanValue {
	if m != nil {
		return m.KillProcessesForUserFirst
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetError() *dstore_engine_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetMetaInformation() []*dstore_engine_metainformation.MetaInformation {
	if m != nil {
		return m.MetaInformation
	}
	return nil
}

func (m *Response) GetMessage() []*dstore_engine_message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Response) GetRow() []*Response_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

type Response_Row struct {
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_DeletedStoreUser_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_DeletedStoreUser_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_DeletedStoreUser_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0x86, 0xd5, 0xe6, 0x4b, 0x9b, 0xcf, 0x20, 0x40, 0x46, 0x42, 0xdb, 0xfc, 0xa0, 0x26, 0x14,
	0x04, 0x27, 0x1b, 0x54, 0x84, 0xe0, 0xb4, 0x55, 0x5a, 0x11, 0x44, 0xa3, 0xb2, 0x50, 0x10, 0x20,
	0x64, 0x6d, 0xb0, 0x93, 0x58, 0x38, 0x76, 0x64, 0x3b, 0x8d, 0xb8, 0x0b, 0xae, 0x06, 0x71, 0x4b,
	0xfc, 0xdd, 0x03, 0x63, 0x7b, 0x97, 0x90, 0x85, 0x26, 0xe5, 0x64, 0x37, 0xb6, 0xe7, 0x7d, 0xe6,
	0xcd, 0xac, 0x67, 0xd0, 0x7d, 0x6a, 0xac, 0xd2, 0xac, 0xcd, 0xe4, 0x90, 0x4b, 0xd6, 0x9e, 0x68,
	0xf5, 0x8e, 0xd1, 0xa9, 0x66, 0xa6, 0x3d, 0xe6, 0xa4, 0xc3, 0x04, 0xb3, 0x8c, 0x3e, 0x73, 0x11,
	0x27, 0x86, 0x69, 0xb2, 0x47, 0x63, 0x88, 0xb0, 0x0a, 0xef, 0x04, 0x59, 0x1c, 0x64, 0xf1, 0xdf,
	0x63, 0xab, 0x57, 0x33, 0xf8, 0x69, 0x2a, 0xa6, 0xcc, 0x04, 0x69, 0x75, 0x6b, 0x31, 0x23, 0xd3,
	0x5a, 0xe9, 0xec, 0xa8, 0xb6, 0x78, 0x34, 0x66, 0xc6, 0xa4, 0x43, 0x96, 0x1d, 0xde, 0x28, 0x1e,
	0xda, 0x94, 0xcb, 0x81, 0xd2, 0xe3, 0xd4, 0x72, 0x25, 0x43, 0x50, 0xeb, 0xf3, 0x06, 0x42, 0xc7,
	0xa9, 0x4e, 0xe1, 0x94, 0x69, 0x83, 0x1f, 0xa0, 0xff, 0xa7, 0xce, 0x8b, 0x84, 0x8d, 0x68, 0x6d,
	0x7b, 0xed, 0xf6, 0x85, 0xdd, 0x6a, 0x9c, 0x59, 0xcf, 0x4c, 0x19, 0xab, 0xb9, 0x1c, 0xbe, 0x70,
	0x8b, 0xa4, 0xe2, 0x82, 0x7b, 0x10, 0x8b, 0x6f, 0xa2, 0x4b, 0xbf, 0x84, 0x44, 0x4e, 0x85, 0x88,
	0xbe, 0x6c, 0x82, 0xbc, 0x92, 0x5c, 0xcc, 0x43, 0x7a, 0xb0, 0x89, 0xdf, 0xa0, 0x3a, 0x97, 0x00,
	0xe2, 0x34, 0xb5, 0x8c, 0x78, 0x05, 0x97, 0xf9, 0x7b, 0xa0, 0xa2, 0x75, 0x9f, 0xb2, 0x56, 0x48,
	0xd9, 0x57, 0x4a, 0xb0, 0x54, 0x86, 0x9c, 0xd1, 0x1c, 0xe0, 0xca, 0xd6, 0x95, 0xe1, 0x39, 0x50,
	0xf8, 0x11, 0x6a, 0x2e, 0x83, 0x07, 0x5b, 0x5f, 0x83, 0xad, 0xfa, 0x59, 0x14, 0x6f, 0x93, 0xa2,
	0xa6, 0x92, 0xe2, 0x03, 0xa1, 0xfe, 0x1b, 0x05, 0xc4, 0x40, 0xab, 0x31, 0xb1, 0x23, 0x6e, 0x48,
	0x28, 0x68, 0x54, 0x5a, 0xed, 0xb5, 0xee, 0x28, 0xe1, 0x43, 0x3b, 0xfe, 0x21, 0x20, 0x9e, 0x03,
	0xe1, 0xc0, 0x03, 0xf0, 0x53, 0x74, 0x6b, 0x65, 0x96, 0x60, 0xfa, 0x5b, 0x30, 0xbd, 0xbd, 0x0c,
	0xe7, 0x8d, 0x8f, 0xd0, 0x0e, 0xd5, 0x6a, 0x42, 0x84, 0x82, 0x2d, 0x32, 0x1b, 0x31, 0x49, 0x28,
	0xe9, 0x13, 0xc0, 0x49, 0x65, 0x49, 0x7a, 0x9a, 0x72, 0x91, 0xf6, 0x05, 0x8b, 0xfe, 0x5b, 0xed,
	0xbd, 0xe1, 0x40, 0x4f, 0x1c, 0xe7, 0x25, 0x60, 0x3a, 0xfb, 0x5d, 0xd3, 0x53, 0x76, 0x2f, 0x27,
	0xe0, 0x13, 0x74, 0xe7, 0x3c, 0x99, 0x82, 0xff, 0xef, 0xc1, 0x7f, 0x73, 0x29, 0xd2, 0xff, 0x81,
	0xb7, 0xa8, 0xf1, 0x9e, 0x0b, 0x41, 0x7c, 0x5f, 0x19, 0xc3, 0x0c, 0x81, 0x1b, 0x9b, 0x95, 0x86,
	0x6b, 0x63, 0xa3, 0xf2, 0x6a, 0xe7, 0x5b, 0x8e, 0x70, 0x9c, 0x03, 0x0e, 0x95, 0xf6, 0xc5, 0x72,
	0x6a, 0xfc, 0x18, 0xb5, 0x96, 0xe2, 0x83, 0xdd, 0x1f, 0xc1, 0x6e, 0xe3, 0x4c, 0x8e, 0xb3, 0xda,
	0xfa, 0xb4, 0x8e, 0x2a, 0x09, 0x33, 0x13, 0x25, 0x0d, 0xc3, 0x77, 0x51, 0xd9, 0x37, 0x66, 0xb1,
	0x69, 0xb2, 0x7e, 0x0f, 0x4d, 0x7b, 0xe0, 0x9e, 0x49, 0x08, 0xc4, 0xaf, 0xd0, 0x15, 0xd7, 0x92,
	0xe4, 0xb7, 0x9e, 0x84, 0xeb, 0x5f, 0x02, 0x71, 0x5c, 0x10, 0x17, 0x3b, 0xf7, 0x08, 0xd6, 0xdd,
	0xf9, 0x3a, 0xb9, 0x3c, 0x5e, 0xdc, 0xc0, 0x0f, 0xd1, 0x66, 0x36, 0x0a, 0xe0, 0x92, 0x3a, 0xe2,
	0xf5, 0x3f, 0x88, 0x61, 0x50, 0x1c, 0x85, 0x77, 0x92, 0x87, 0xe3, 0x0e, 0x2a, 0x69, 0x35, 0x83,
	0xeb, 0xe1, 0x54, 0xbb, 0xf1, 0x79, 0x86, 0x56, 0x9c, 0xd7, 0x20, 0x4e, 0xd4, 0x2c, 0x71, 0xf2,
	0x6a, 0x03, 0x95, 0xe0, 0x37, 0xbe, 0x86, 0x36, 0x60, 0x45, 0x38, 0x8d, 0x3e, 0xf6, 0xa0, 0x2a,
	0xe5, 0xa4, 0x0c, 0xcb, 0x2e, 0xdd, 0xef, 0xa1, 0x1a, 0x57, 0x05, 0xf6, 0x7c, 0x8e, 0xbe, 0x6e,
	0xff, 0xe3, 0x84, 0xed, 0x6f, 0xf8, 0x51, 0x76, 0xef, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14,
	0xf2, 0xa8, 0x67, 0x9b, 0x05, 0x00, 0x00,
}
