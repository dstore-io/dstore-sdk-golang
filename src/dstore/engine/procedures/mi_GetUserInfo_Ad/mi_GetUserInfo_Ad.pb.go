// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetUserInfo_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetUserInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetUserInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetUserInfo_Ad

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
	UserId                 *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull             bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	IncludeSpecialUser     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=include_special_user,json=includeSpecialUser" json:"include_special_user,omitempty"`
	IncludeSpecialUserNull bool                        `protobuf:"varint,1002,opt,name=include_special_user_null,json=includeSpecialUserNull" json:"include_special_user_null,omitempty"`
	GetAssignedGroups      *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=get_assigned_groups,json=getAssignedGroups" json:"get_assigned_groups,omitempty"`
	GetAssignedGroupsNull  bool                        `protobuf:"varint,1003,opt,name=get_assigned_groups_null,json=getAssignedGroupsNull" json:"get_assigned_groups_null,omitempty"`
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

func (m *Parameters) GetIncludeSpecialUser() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeSpecialUser
	}
	return nil
}

func (m *Parameters) GetGetAssignedGroups() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetAssignedGroups
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
	UserName         *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserId           *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupName        *dstore_values.StringValue  `protobuf:"bytes,20001,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	GroupDescription *dstore_values.StringValue  `protobuf:"bytes,20002,opt,name=group_description,json=groupDescription" json:"group_description,omitempty"`
	UserGroupId      *dstore_values.IntegerValue `protobuf:"bytes,20005,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	SortNo           *dstore_values.IntegerValue `protobuf:"bytes,20006,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetGroupName() *dstore_values.StringValue {
	if m != nil {
		return m.GroupName
	}
	return nil
}

func (m *Response_Row) GetGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.GroupDescription
	}
	return nil
}

func (m *Response_Row) GetUserGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserGroupId
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetUserInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetUserInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetUserInfo_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x6b, 0x13, 0x5f,
	0x10, 0xc7, 0x49, 0xf3, 0xcb, 0xa5, 0xd3, 0x9f, 0xd8, 0x9e, 0x6a, 0xd9, 0x26, 0xe0, 0xa5, 0xbe,
	0x88, 0xe0, 0xa6, 0x78, 0xa3, 0xe8, 0x83, 0x44, 0x94, 0x12, 0x24, 0x8b, 0xac, 0x28, 0xe8, 0xcb,
	0xb2, 0xcd, 0x8e, 0xcb, 0x81, 0xcd, 0x39, 0xe1, 0x9c, 0x5d, 0xfb, 0x6f, 0x58, 0x9f, 0xfb, 0xa2,
	0x60, 0xff, 0x37, 0xf5, 0x0f, 0xf0, 0xd5, 0x39, 0x97, 0x12, 0x73, 0x81, 0xe4, 0x25, 0x61, 0xce,
	0xcc, 0x67, 0x66, 0xce, 0x7c, 0x67, 0x0f, 0x1c, 0x66, 0xba, 0x94, 0x0a, 0x7b, 0x28, 0x72, 0x2e,
	0xb0, 0x37, 0x51, 0x72, 0x84, 0x59, 0xa5, 0x50, 0xf7, 0xc6, 0x3c, 0x39, 0xc6, 0xf2, 0x9d, 0x46,
	0x35, 0x10, 0x9f, 0x64, 0xd2, 0xcf, 0x42, 0x72, 0x96, 0x92, 0xdd, 0x74, 0x44, 0xe8, 0x88, 0x70,
	0x21, 0xac, 0xb3, 0xeb, 0x53, 0x7e, 0x4e, 0x8b, 0x0a, 0xb5, 0xa3, 0x3a, 0xfb, 0xb3, 0x75, 0x50,
	0x29, 0xa9, 0xbc, 0xab, 0x3b, 0xeb, 0x1a, 0xa3, 0xd6, 0x69, 0x8e, 0xde, 0x79, 0x67, 0xde, 0x59,
	0xa6, 0x9c, 0xea, 0xa8, 0x71, 0x5a, 0x72, 0x29, 0x5c, 0xd0, 0xc1, 0x9f, 0x0d, 0x80, 0x37, 0xa9,
	0x4a, 0xc9, 0x8b, 0x4a, 0xb3, 0x47, 0xd0, 0xaa, 0xa8, 0x9f, 0x84, 0x67, 0x41, 0xed, 0x56, 0xed,
	0xee, 0xd6, 0x83, 0x6e, 0xe8, 0x7b, 0xf6, 0x2d, 0x71, 0x51, 0x62, 0x8e, 0xea, 0xbd, 0xb1, 0xe2,
	0xa6, 0x89, 0x1d, 0x64, 0xec, 0x36, 0xfc, 0xef, 0xa9, 0x44, 0x54, 0x45, 0x11, 0xfc, 0x6c, 0x11,
	0xdb, 0x8e, 0xc1, 0xb9, 0x23, 0x3a, 0x62, 0x43, 0xb8, 0xc6, 0xc5, 0xa8, 0xa8, 0x32, 0x4c, 0xf4,
	0x04, 0x47, 0x3c, 0x2d, 0x12, 0xe3, 0x0d, 0x36, 0x96, 0x56, 0x39, 0x91, 0xb2, 0xc0, 0x54, 0xb8,
	0x2a, 0xcc, 0x83, 0x6f, 0x1d, 0x67, 0xe6, 0xc5, 0x9e, 0xc2, 0xfe, 0xb2, 0x74, 0xae, 0xfc, 0x2f,
	0x57, 0x7e, 0x6f, 0x91, 0xb3, 0xad, 0xbc, 0x86, 0xdd, 0x1c, 0xcb, 0x24, 0xd5, 0x9a, 0xe7, 0x02,
	0xb3, 0x24, 0x57, 0xb2, 0x9a, 0xe8, 0xa0, 0xbe, 0xba, 0x93, 0x1d, 0xe2, 0xfa, 0x1e, 0x3b, 0xb6,
	0x14, 0x3b, 0x82, 0x60, 0x49, 0x32, 0xd7, 0xc7, 0x6f, 0xd7, 0xc7, 0xf5, 0x05, 0xca, 0xb4, 0x71,
	0x70, 0xd1, 0x80, 0x76, 0x8c, 0x7a, 0x22, 0x85, 0x46, 0x76, 0x08, 0x0d, 0xab, 0xab, 0x9f, 0x7a,
	0x27, 0x9c, 0xdd, 0x14, 0xa7, 0xf9, 0x2b, 0xf3, 0x1b, 0xbb, 0x40, 0xf6, 0x01, 0xb6, 0x8d, 0xa2,
	0xc9, 0x3f, 0x92, 0xd2, 0x30, 0xeb, 0x04, 0x87, 0x73, 0xf0, 0xbc, 0xf0, 0x43, 0xb2, 0x07, 0x53,
	0x3b, 0xbe, 0x3a, 0x9e, 0x3d, 0xa0, 0x3b, 0xb5, 0xfc, 0x26, 0xd1, 0x50, 0x4c, 0xc6, 0x1b, 0x0b,
	0x19, 0xdd, 0x9e, 0x0d, 0xdd, 0x7f, 0x7c, 0x19, 0xce, 0x9e, 0x43, 0x5d, 0xc9, 0xd3, 0xe0, 0x3f,
	0x4b, 0xdd, 0x0f, 0x57, 0xac, 0x7b, 0x78, 0x79, 0xfd, 0x30, 0x96, 0xa7, 0xb1, 0x21, 0x3b, 0x67,
	0x75, 0xa8, 0x93, 0xc1, 0xf6, 0xa0, 0x49, 0xa6, 0x59, 0xc3, 0x2f, 0x11, 0x4d, 0xa4, 0x11, 0x37,
	0xc8, 0xa4, 0x4d, 0x3b, 0x82, 0x4d, 0xa7, 0x33, 0x2d, 0x6c, 0x70, 0x16, 0xcd, 0x0e, 0xcb, 0x4b,
	0xa6, 0x4b, 0xc5, 0x45, 0xee, 0x14, 0x6b, 0x9b, 0xe8, 0x88, 0x82, 0xd9, 0xe3, 0xe9, 0x66, 0x7f,
	0x8d, 0xd6, 0x5f, 0xed, 0x67, 0x00, 0x56, 0x52, 0x57, 0xf1, 0xdb, 0x79, 0x6d, 0x65, 0xc9, 0x4d,
	0x1b, 0x6f, 0x6b, 0x0e, 0x60, 0xc7, 0xc1, 0x19, 0xea, 0x91, 0xe2, 0x13, 0x2b, 0xd2, 0xf7, 0x35,
	0x72, 0x6c, 0x5b, 0xec, 0xe5, 0x94, 0x62, 0x7d, 0xb8, 0x62, 0xdb, 0x77, 0xf9, 0xe8, 0x12, 0x3f,
	0xce, 0xd7, 0xf8, 0x40, 0xb7, 0x0c, 0x63, 0x77, 0x8e, 0xae, 0xf2, 0x04, 0x5a, 0x5a, 0xaa, 0x32,
	0x11, 0x32, 0xb8, 0x58, 0x07, 0x6e, 0x9a, 0xe8, 0x48, 0xbe, 0x18, 0x40, 0x97, 0xcb, 0x39, 0x2d,
	0xa7, 0x8f, 0xdd, 0xc7, 0x7b, 0xeb, 0x3f, 0x83, 0x27, 0x4d, 0xfb, 0xe8, 0x3c, 0xfc, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x92, 0xd7, 0x22, 0xdc, 0x3b, 0x05, 0x00, 0x00,
}
