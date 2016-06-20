// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetForumAccessMatrix_Ad.proto
// DO NOT EDIT!

/*
Package fo_GetForumAccessMatrix_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetForumAccessMatrix_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetForumAccessMatrix_Ad

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
	ForumId               *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumIdNull           bool                        `protobuf:"varint,1001,opt,name=forum_id_null,json=forumIdNull" json:"forum_id_null,omitempty"`
	CheckPersonAccess     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=check_person_access,json=checkPersonAccess" json:"check_person_access,omitempty"`
	CheckPersonAccessNull bool                        `protobuf:"varint,1002,opt,name=check_person_access_null,json=checkPersonAccessNull" json:"check_person_access_null,omitempty"`
	GroupOrPersonId       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=group_or_person_id,json=groupOrPersonId" json:"group_or_person_id,omitempty"`
	GroupOrPersonIdNull   bool                        `protobuf:"varint,1003,opt,name=group_or_person_id_null,json=groupOrPersonIdNull" json:"group_or_person_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Parameters) GetCheckPersonAccess() *dstore_values.BooleanValue {
	if m != nil {
		return m.CheckPersonAccess
	}
	return nil
}

func (m *Parameters) GetGroupOrPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupOrPersonId
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
	RowId            int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	GroupDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=group_description,json=groupDescription" json:"group_description,omitempty"`
	AccessLevelId    *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=access_level_id,json=accessLevelId" json:"access_level_id,omitempty"`
	Description      *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=description" json:"description,omitempty"`
	ForumId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	Access           *dstore_values.BooleanValue `protobuf:"bytes,10005,opt,name=access" json:"access,omitempty"`
	GroupId          *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	PersonId         *dstore_values.IntegerValue `protobuf:"bytes,20003,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	AnonymousAccess  *dstore_values.BooleanValue `protobuf:"bytes,20004,opt,name=anonymous_access,json=anonymousAccess" json:"anonymous_access,omitempty"`
	GroupAccess      *dstore_values.BooleanValue `protobuf:"bytes,20007,opt,name=group_access,json=groupAccess" json:"group_access,omitempty"`
	AccessType       *dstore_values.IntegerValue `protobuf:"bytes,20008,opt,name=access_type,json=accessType" json:"access_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.GroupDescription
	}
	return nil
}

func (m *Response_Row) GetAccessLevelId() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevelId
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Response_Row) GetAccess() *dstore_values.BooleanValue {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *Response_Row) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetAnonymousAccess() *dstore_values.BooleanValue {
	if m != nil {
		return m.AnonymousAccess
	}
	return nil
}

func (m *Response_Row) GetGroupAccess() *dstore_values.BooleanValue {
	if m != nil {
		return m.GroupAccess
	}
	return nil
}

func (m *Response_Row) GetAccessType() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessType
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetForumAccessMatrix_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetForumAccessMatrix_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetForumAccessMatrix_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x1b, 0xd2, 0x84, 0x09, 0x55, 0x52, 0x47, 0x80, 0x49, 0x24, 0x84, 0xda, 0x4b, 0x4f,
	0x0e, 0xa2, 0x50, 0x7e, 0x24, 0x10, 0x45, 0x05, 0x52, 0x41, 0x4c, 0x65, 0x21, 0x24, 0xb8, 0x58,
	0xae, 0xbd, 0x09, 0x16, 0xce, 0xae, 0xb5, 0xb6, 0x5b, 0xfa, 0x16, 0xfc, 0x9f, 0x7a, 0x03, 0x09,
	0x9e, 0x81, 0x37, 0x02, 0x4e, 0xbc, 0x01, 0xb3, 0x3f, 0x26, 0x71, 0xd2, 0x2a, 0x81, 0x4b, 0xa2,
	0xf5, 0xce, 0x37, 0xdf, 0x37, 0xdf, 0xcc, 0xd8, 0x70, 0x33, 0x48, 0x52, 0xc6, 0x49, 0x87, 0xd0,
	0x41, 0x48, 0x49, 0x27, 0xe6, 0xcc, 0x27, 0x41, 0xc6, 0x49, 0xd2, 0xe9, 0x33, 0xf7, 0x21, 0x49,
	0x1f, 0x30, 0x9e, 0x0d, 0xb7, 0x7c, 0x9f, 0x24, 0x49, 0xcf, 0x4b, 0x79, 0xf8, 0xda, 0xdd, 0x0a,
	0x2c, 0x8c, 0x4a, 0x99, 0xb1, 0xae, 0xa0, 0x96, 0x82, 0x5a, 0x27, 0xc7, 0xb7, 0x9a, 0x9a, 0x64,
	0xdf, 0x8b, 0x32, 0x92, 0x28, 0x78, 0xeb, 0x42, 0x91, 0x99, 0x70, 0xce, 0xb8, 0xbe, 0x6a, 0x17,
	0xaf, 0x86, 0x98, 0xcc, 0x1b, 0x10, 0x7d, 0xb9, 0x36, 0x79, 0x99, 0x7a, 0x21, 0xed, 0x33, 0x3e,
	0xf4, 0xd2, 0x90, 0x51, 0x15, 0xb4, 0xfa, 0x7b, 0x11, 0x60, 0xd7, 0xe3, 0x1e, 0xde, 0x12, 0x9e,
	0x18, 0x9b, 0x50, 0xed, 0x0b, 0x61, 0x6e, 0x18, 0x98, 0x0b, 0x97, 0x16, 0xd6, 0x6b, 0x57, 0xda,
	0x96, 0x56, 0xaf, 0x35, 0x85, 0x34, 0x25, 0x03, 0xc2, 0x9f, 0x89, 0x93, 0x53, 0x91, 0xc1, 0x3b,
	0x81, 0xb1, 0x06, 0xcb, 0x39, 0xce, 0xa5, 0x59, 0x14, 0x99, 0x3f, 0x2a, 0x88, 0xae, 0x3a, 0x35,
	0x1d, 0x60, 0xe3, 0x33, 0xe3, 0x11, 0x34, 0xfd, 0x97, 0xc4, 0x7f, 0xe5, 0xc6, 0x48, 0xc5, 0xa8,
	0xeb, 0xc9, 0xea, 0xcd, 0xc5, 0x63, 0x79, 0xf6, 0x18, 0x8b, 0x88, 0x47, 0x15, 0xcf, 0x8a, 0xc4,
	0xed, 0x4a, 0x98, 0xf2, 0xcc, 0xb8, 0x01, 0xe6, 0x31, 0xc9, 0x14, 0xf9, 0x4f, 0x45, 0x7e, 0x76,
	0x0a, 0x25, 0x65, 0x74, 0xc1, 0x18, 0x70, 0x96, 0xc5, 0x2e, 0xe3, 0x39, 0x18, 0xab, 0x2d, 0xcd,
	0xae, 0xb6, 0x2e, 0x61, 0x4f, 0xb8, 0xca, 0x88, 0x55, 0x5f, 0x83, 0xf3, 0xd3, 0x99, 0x94, 0x84,
	0x5f, 0x4a, 0x42, 0x73, 0x02, 0x22, 0x04, 0xac, 0x7e, 0xaf, 0x40, 0xd5, 0x21, 0x49, 0xcc, 0x68,
	0x42, 0x8c, 0xcb, 0x50, 0x96, 0x1d, 0xd5, 0x76, 0xb7, 0xac, 0xe2, 0xb0, 0xa8, 0x6e, 0xdf, 0x17,
	0xbf, 0x8e, 0x0a, 0x34, 0x9e, 0x43, 0x43, 0xf4, 0xd2, 0x1d, 0x6b, 0x26, 0x7a, 0x58, 0x42, 0xb0,
	0x35, 0x01, 0x9e, 0x6c, 0x79, 0x0f, 0xcf, 0x3b, 0xa3, 0xb3, 0x53, 0x1f, 0x16, 0x1f, 0xa0, 0xa9,
	0x15, 0x3d, 0x43, 0xe8, 0x87, 0xc8, 0x78, 0x71, 0x2a, 0xa3, 0x9a, 0xb0, 0x9e, 0xfa, 0x77, 0xf2,
	0x70, 0x34, 0xb5, 0xc4, 0xd9, 0x81, 0x79, 0x4a, 0xa2, 0x36, 0xad, 0x79, 0x27, 0xde, 0xca, 0x7d,
	0xb0, 0x1c, 0x76, 0xe0, 0x88, 0x14, 0xad, 0xa3, 0x32, 0x94, 0xf0, 0x60, 0x9c, 0x83, 0x25, 0x3c,
	0x8a, 0xd6, 0xbc, 0xb1, 0xd1, 0x9a, 0xb2, 0x53, 0xc6, 0x23, 0x9a, 0xde, 0x85, 0x15, 0x65, 0x7a,
	0x40, 0x12, 0x9f, 0x87, 0xb1, 0xac, 0xff, 0xad, 0x5d, 0x74, 0x4f, 0xb7, 0x2f, 0x41, 0x1a, 0x3a,
	0x50, 0xdd, 0x6b, 0x48, 0xd4, 0xf6, 0x08, 0x64, 0x6c, 0x43, 0x5d, 0x4f, 0x4d, 0x44, 0xf6, 0x49,
	0x24, 0xa8, 0xde, 0xd9, 0xb3, 0xc7, 0x60, 0x59, 0x81, 0x1e, 0x0b, 0x0c, 0xea, 0xb9, 0x0d, 0xb5,
	0x71, 0x25, 0xef, 0x67, 0x2b, 0x19, 0x8f, 0x37, 0xae, 0x8f, 0x6d, 0xdc, 0x07, 0xfb, 0x1f, 0x56,
	0xee, 0x2a, 0x2c, 0xe9, 0x05, 0xfa, 0x68, 0xcf, 0xde, 0x20, 0x1d, 0x2b, 0xe8, 0x94, 0x7b, 0x48,
	0xf7, 0x69, 0x1e, 0x3a, 0x19, 0x8d, 0x74, 0xb7, 0xe0, 0xf4, 0x68, 0x59, 0x3e, 0x1f, 0xcd, 0xf1,
	0x72, 0xa8, 0xc6, 0xf9, 0x9e, 0x74, 0xa1, 0xe1, 0x51, 0x46, 0x0f, 0x87, 0x2c, 0x4b, 0xf2, 0xad,
	0xff, 0x72, 0x42, 0x8a, 0x82, 0xea, 0xfa, 0x5f, 0x98, 0xde, 0xfa, 0xbb, 0x70, 0x46, 0xc9, 0xd7,
	0x59, 0xbe, 0xce, 0x93, 0xa5, 0x26, 0x21, 0x3a, 0xc3, 0x1d, 0xa8, 0xe9, 0xa6, 0xa7, 0x87, 0x31,
	0x31, 0xbf, 0xcd, 0x53, 0x09, 0x28, 0xc4, 0x53, 0x04, 0xdc, 0x73, 0xa0, 0x1d, 0xb2, 0x89, 0xf9,
	0x1e, 0x7d, 0x0c, 0x5e, 0x6c, 0xfc, 0xc7, 0x67, 0x62, 0x6f, 0x49, 0xbe, 0x8b, 0x37, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x72, 0x64, 0x39, 0x13, 0x64, 0x06, 0x00, 0x00,
}