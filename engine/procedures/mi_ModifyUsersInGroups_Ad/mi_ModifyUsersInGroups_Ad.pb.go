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
import dstore_values "gosdk.dstore.de/values"
import dstore_engine_error "gosdk.dstore.de/engine/error"
import dstore_engine_message "gosdk.dstore.de/engine/message"
import dstore_engine_metainformation "gosdk.dstore.de/engine/metainformation"

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

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyUsersInGroups_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x66, 0x1b, 0xdb, 0x2e, 0x67, 0x77, 0xd9, 0x65, 0x04, 0x89, 0x29, 0x8a, 0xd6, 0x0b, 0xbd,
	0xd0, 0xa9, 0xf8, 0x03, 0x05, 0x2f, 0xfc, 0x01, 0x29, 0xb9, 0x68, 0x29, 0x03, 0x0a, 0xf6, 0x26,
	0x44, 0x67, 0x1a, 0x06, 0x9b, 0x99, 0x30, 0x93, 0xb4, 0xf4, 0x2d, 0x7c, 0x1d, 0x1f, 0xc0, 0x87,
	0x51, 0x5f, 0xc2, 0x99, 0xcc, 0x94, 0x36, 0x11, 0xa9, 0x7b, 0x93, 0xe4, 0xcc, 0xf9, 0x7e, 0x4e,
	0xce, 0x39, 0x03, 0x63, 0xaa, 0x4b, 0xa9, 0xd8, 0x88, 0x89, 0x8c, 0x0b, 0x36, 0x2a, 0x94, 0xfc,
	0xc2, 0x68, 0xa5, 0x98, 0x1e, 0xe5, 0x3c, 0x99, 0x4a, 0xca, 0x97, 0xdb, 0x0f, 0x9a, 0x29, 0x1d,
	0x8b, 0x89, 0x92, 0x55, 0xa1, 0x93, 0xb7, 0x14, 0x1b, 0x50, 0x29, 0xd1, 0x43, 0xc7, 0xc4, 0x8e,
	0x89, 0xff, 0x09, 0x8f, 0x6e, 0x7a, 0x8b, 0x75, 0xba, 0xaa, 0x98, 0x76, 0xec, 0xe8, 0x76, 0xd3,
	0x97, 0x29, 0x25, 0x95, 0x4f, 0x0d, 0x9a, 0xa9, 0x9c, 0x69, 0x9d, 0x66, 0xcc, 0x27, 0x1f, 0xb4,
	0x93, 0x65, 0xca, 0xc5, 0x52, 0xaa, 0x3c, 0x2d, 0xb9, 0x14, 0x0e, 0x34, 0xfc, 0xd1, 0x01, 0x98,
	0xa7, 0x2a, 0x35, 0x59, 0x53, 0x0b, 0x7a, 0x01, 0xfd, 0xca, 0x14, 0x95, 0x70, 0x1a, 0x9e, 0xdc,
	0x3b, 0x79, 0x74, 0xf6, 0x6c, 0x80, 0x7d, 0xed, 0xbe, 0x24, 0x2e, 0x4a, 0x96, 0x31, 0xf5, 0xd1,
	0x46, 0xa4, 0x67, 0xb1, 0x31, 0x45, 0xf7, 0xe1, 0xdc, 0xb3, 0x12, 0x51, 0xad, 0x56, 0xe1, 0xcf,
	0xbe, 0xe1, 0x9e, 0x12, 0x70, 0xe9, 0x99, 0x39, 0x42, 0xaf, 0xe1, 0xa2, 0x86, 0x64, 0xf6, 0x5f,
	0xad, 0x7c, 0xe7, 0xb8, 0xfc, 0x99, 0x65, 0xd4, 0xcd, 0x31, 0x1e, 0x8f, 0x01, 0x35, 0x04, 0x9c,
	0xd3, 0x2f, 0xe7, 0x74, 0x79, 0x80, 0xac, 0xed, 0xde, 0xc0, 0x45, 0x2e, 0xd7, 0x2c, 0x29, 0x14,
	0x97, 0x8a, 0x97, 0xdb, 0x30, 0x38, 0x6e, 0x77, 0x6e, 0x19, 0x73, 0x4f, 0x40, 0x4f, 0x00, 0x35,
	0x14, 0x9c, 0xdf, 0x6f, 0xe7, 0x77, 0x75, 0x08, 0xb5, 0x86, 0xc3, 0xef, 0x1d, 0x38, 0x25, 0x4c,
	0x17, 0x52, 0x68, 0x86, 0x9e, 0x42, 0xb7, 0x9e, 0x92, 0xef, 0x61, 0x84, 0x9b, 0xf3, 0x77, 0x13,
	0x7c, 0x6f, 0x9f, 0xc4, 0x01, 0xd1, 0x27, 0xb8, 0xb2, 0xf3, 0x49, 0x0e, 0x06, 0x64, 0x3a, 0x14,
	0x18, 0x32, 0x6e, 0x91, 0xdb, 0x63, 0x9c, 0x9a, 0x38, 0xde, 0xc7, 0xe4, 0x32, 0x6f, 0x1e, 0xa0,
	0x31, 0xf4, 0xfd, 0x5e, 0x98, 0x26, 0x58, 0xc5, 0xbb, 0x7f, 0x29, 0xba, 0xad, 0x99, 0xba, 0x37,
	0xd9, 0xc1, 0xd1, 0x04, 0x02, 0x25, 0x37, 0xe1, 0x8d, 0x9a, 0xf5, 0x12, 0xff, 0xe7, 0x12, 0xe3,
	0x5d, 0x1b, 0x30, 0x91, 0x1b, 0x62, 0x15, 0xa2, 0x3b, 0x10, 0x98, 0x6f, 0x74, 0x0b, 0x7a, 0x26,
	0xb2, 0xc3, 0xff, 0x36, 0x33, 0x8d, 0xe9, 0x92, 0xae, 0x09, 0x63, 0xfa, 0x6e, 0x01, 0x03, 0x2e,
	0x5b, 0xf2, 0xfb, 0xdb, 0xb5, 0x78, 0x95, 0x49, 0x4d, 0xbf, 0xee, 0xf2, 0xf4, 0x5a, 0x17, 0xf0,
	0x73, 0xaf, 0x5e, 0xf3, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x5a, 0x2c, 0x2b, 0xbd,
	0x03, 0x00, 0x00,
}
