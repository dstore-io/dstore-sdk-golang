// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyUsersInGroups_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyUsersInGroups_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyUsersInGroups_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyUsersInGroups_Ad

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
	UserId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull       bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	UserGroupId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	UserGroupIdNull  bool                        `protobuf:"varint,1002,opt,name=user_group_id_null,json=userGroupIdNull" json:"user_group_id_null,omitempty"`
	MovePriority     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=move_priority,json=movePriority" json:"move_priority,omitempty"`
	MovePriorityNull bool                        `protobuf:"varint,1003,opt,name=move_priority_null,json=movePriorityNull" json:"move_priority_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Parameters) GetUserGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserGroupId
	}
	return nil
}

func (m *Parameters) GetMovePriority() *dstore_values.IntegerValue {
	if m != nil {
		return m.MovePriority
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyUsersInGroups_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyUsersInGroups_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyUsersInGroups_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0xea, 0xd3, 0x30,
	0x14, 0xe7, 0xbf, 0xba, 0xed, 0xcf, 0xd9, 0xc6, 0x46, 0x04, 0xa9, 0x1d, 0x8a, 0xce, 0x0b, 0xbd,
	0xd0, 0x4c, 0xa6, 0xc2, 0xee, 0xfc, 0x00, 0x19, 0xbd, 0xd8, 0x98, 0x01, 0x05, 0xbd, 0x29, 0xd5,
	0x66, 0x25, 0xb0, 0x26, 0x25, 0x69, 0x37, 0xf6, 0x16, 0xbe, 0x8e, 0x0f, 0xe0, 0xc3, 0xa8, 0x2f,
	0x61, 0xd2, 0x64, 0x6c, 0xad, 0xc8, 0xf4, 0xa6, 0xed, 0xc9, 0xf9, 0x7d, 0x9c, 0x9e, 0x73, 0x02,
	0xf3, 0x44, 0x15, 0x42, 0xd2, 0x29, 0xe5, 0x29, 0xe3, 0x74, 0x9a, 0x4b, 0xf1, 0x85, 0x26, 0xa5,
	0xa4, 0x6a, 0x9a, 0xb1, 0x68, 0x29, 0x12, 0xb6, 0x39, 0xbc, 0x57, 0x54, 0xaa, 0x90, 0x2f, 0xa4,
	0x28, 0x73, 0x15, 0xbd, 0x4e, 0xb0, 0x06, 0x15, 0x02, 0x3d, 0xb4, 0x4c, 0x6c, 0x99, 0xf8, 0xaf,
	0xf0, 0xe0, 0xa6, 0xb3, 0xd8, 0xc5, 0xdb, 0x92, 0x2a, 0xcb, 0x0e, 0x6e, 0xd7, 0x7d, 0xa9, 0x94,
	0x42, 0xba, 0xd4, 0xb8, 0x9e, 0xca, 0xa8, 0x52, 0x71, 0x4a, 0x5d, 0xf2, 0x41, 0x33, 0x59, 0xc4,
	0x8c, 0x6f, 0x84, 0xcc, 0xe2, 0x82, 0x09, 0x6e, 0x41, 0x93, 0xef, 0x2d, 0x80, 0x75, 0x2c, 0x63,
	0x9d, 0xd5, 0xb5, 0xa0, 0xe7, 0xd0, 0x2d, 0x75, 0x51, 0x11, 0x4b, 0xfc, 0xab, 0x7b, 0x57, 0x8f,
	0x7a, 0xb3, 0x31, 0x76, 0xb5, 0xbb, 0x92, 0x18, 0x2f, 0x68, 0x4a, 0xe5, 0x07, 0x13, 0x91, 0x8e,
	0xc1, 0x86, 0x09, 0xba, 0x0f, 0x7d, 0xc7, 0x8a, 0x78, 0xb9, 0xdd, 0xfa, 0x3f, 0xba, 0x9a, 0x7b,
	0x4d, 0xc0, 0xa6, 0x57, 0xfa, 0x08, 0xbd, 0x84, 0x41, 0x05, 0x49, 0xcd, 0xbf, 0x1a, 0xf9, 0xd6,
	0x65, 0xf9, 0x9e, 0x61, 0x54, 0xcd, 0xd1, 0x1e, 0x8f, 0x01, 0xd5, 0x04, 0xac, 0xd3, 0x4f, 0xeb,
	0x34, 0x3c, 0x43, 0x56, 0x76, 0xaf, 0x60, 0x90, 0x89, 0x1d, 0x8d, 0x72, 0xc9, 0x84, 0x64, 0xc5,
	0xc1, 0xf7, 0x2e, 0xdb, 0xf5, 0x0d, 0x63, 0xed, 0x08, 0xe8, 0x09, 0xa0, 0x9a, 0x82, 0xf5, 0xfb,
	0x65, 0xfd, 0x46, 0xe7, 0x50, 0x63, 0x38, 0xf9, 0xd6, 0x82, 0x6b, 0x42, 0x55, 0x2e, 0xb8, 0xa2,
	0xe8, 0x29, 0xb4, 0xab, 0x29, 0xb9, 0x1e, 0x06, 0xb8, 0x3e, 0x7f, 0x3b, 0xc1, 0xb7, 0xe6, 0x49,
	0x2c, 0x10, 0x7d, 0x84, 0x91, 0x99, 0x4f, 0x74, 0x36, 0x20, 0xdd, 0x21, 0x4f, 0x93, 0x71, 0x83,
	0xdc, 0x1c, 0xe3, 0x52, 0xc7, 0xe1, 0x29, 0x26, 0xc3, 0xac, 0x7e, 0x80, 0xe6, 0xd0, 0x75, 0x7b,
	0xa1, 0x9b, 0x60, 0x14, 0xef, 0xfe, 0xa1, 0x68, 0xb7, 0x66, 0x69, 0xdf, 0xe4, 0x08, 0x47, 0x0b,
	0xf0, 0xa4, 0xd8, 0xfb, 0x37, 0x2a, 0xd6, 0x0b, 0xfc, 0x8f, 0x4b, 0x8c, 0x8f, 0x6d, 0xc0, 0x44,
	0xec, 0x89, 0x51, 0x08, 0xee, 0x80, 0xa7, 0xbf, 0xd1, 0x2d, 0xe8, 0xe8, 0xc8, 0x0c, 0xff, 0xeb,
	0x4a, 0x37, 0xa6, 0x4d, 0xda, 0x3a, 0x0c, 0x93, 0x37, 0xef, 0x60, 0xcc, 0x44, 0x43, 0xfe, 0x74,
	0xbb, 0x3e, 0xcd, 0xfe, 0xff, 0xde, 0x7d, 0xee, 0x54, 0xdb, 0xfd, 0xec, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xda, 0x55, 0xbe, 0xa8, 0xb4, 0x03, 0x00, 0x00,
}