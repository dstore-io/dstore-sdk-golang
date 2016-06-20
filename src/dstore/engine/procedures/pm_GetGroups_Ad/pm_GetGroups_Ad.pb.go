// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetGroups_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetGroups_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetGroups_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetGroups_Ad

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
	PersonId                     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                 bool                        `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	PersonTypeId                 *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull             bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	IncludeIndependentGroups     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=include_independent_groups,json=includeIndependentGroups" json:"include_independent_groups,omitempty"`
	IncludeIndependentGroupsNull bool                        `protobuf:"varint,1003,opt,name=include_independent_groups_null,json=includeIndependentGroupsNull" json:"include_independent_groups_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetIncludeIndependentGroups() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeIndependentGroups
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
	ConditionId      *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	SortNo           *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	PersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	GroupId          *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
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

func (m *Response_Row) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Response_Row) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetGroups_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetGroups_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetGroups_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x4d, 0xf3, 0xe3, 0xb4, 0x68, 0x9c, 0x82, 0xac, 0x89, 0x7f, 0x54, 0x84, 0x82, 0x30,
	0x11, 0x45, 0xf4, 0x46, 0xa4, 0x45, 0x8d, 0xb9, 0xe8, 0x22, 0x8b, 0x08, 0x7a, 0xb3, 0x6c, 0x33,
	0xc7, 0x30, 0x90, 0xcc, 0x2c, 0x33, 0xb3, 0x16, 0xaf, 0x7d, 0x01, 0x7f, 0xdf, 0xce, 0x17, 0x50,
	0x2f, 0x7c, 0x05, 0xcf, 0xec, 0x4c, 0x92, 0xee, 0xd6, 0xda, 0x78, 0xb3, 0xcb, 0xd9, 0xf3, 0x7d,
	0xe7, 0x3b, 0x73, 0xce, 0xb7, 0x43, 0x18, 0x37, 0x56, 0x69, 0x18, 0x80, 0x9c, 0x08, 0x09, 0x83,
	0x5c, 0xab, 0x31, 0xf0, 0x42, 0x83, 0x19, 0xe4, 0xb3, 0x74, 0x08, 0x76, 0xa8, 0x55, 0x91, 0x9b,
	0x74, 0x8f, 0x33, 0x4c, 0x59, 0x45, 0xaf, 0x7a, 0x3c, 0xf3, 0x78, 0x56, 0x03, 0xf5, 0xb6, 0x43,
	0xb9, 0x77, 0xd9, 0xb4, 0x00, 0xe3, 0x39, 0xbd, 0xcb, 0x55, 0x0d, 0xd0, 0x5a, 0xe9, 0x90, 0xea,
	0x57, 0x53, 0x33, 0x30, 0x26, 0x9b, 0x40, 0x48, 0xde, 0xac, 0x27, 0x6d, 0x26, 0xe4, 0x5b, 0xa5,
	0x67, 0x99, 0x15, 0x4a, 0x7a, 0xd0, 0xce, 0x87, 0x06, 0x21, 0x2f, 0x32, 0x9d, 0x61, 0x16, 0xb4,
	0xa1, 0x0f, 0xc9, 0xb9, 0x1c, 0xdf, 0x4a, 0xa6, 0x82, 0x47, 0x6b, 0x37, 0xd6, 0x76, 0x37, 0xef,
	0xf6, 0xc3, 0x19, 0x59, 0x68, 0x4a, 0x48, 0x0b, 0x13, 0xd0, 0xaf, 0x5c, 0x94, 0x74, 0x3c, 0x7a,
	0xc4, 0xe9, 0x2d, 0x72, 0x7e, 0xc1, 0x4c, 0x65, 0x31, 0x9d, 0x46, 0x3f, 0xda, 0xc8, 0xef, 0x24,
	0x5b, 0x73, 0x48, 0x8c, 0x1f, 0xe9, 0xde, 0x02, 0x66, 0xdf, 0xe7, 0xe0, 0x54, 0xd6, 0xcf, 0x56,
	0x09, 0x25, 0x5e, 0x22, 0x03, 0x95, 0x18, 0xd9, 0xae, 0x96, 0xf0, 0x72, 0x3f, 0xbd, 0x5c, 0xf7,
	0x38, 0xb6, 0x94, 0x7c, 0x4d, 0x7a, 0x42, 0x8e, 0xa7, 0x05, 0x47, 0xac, 0xe4, 0x90, 0x03, 0x3e,
	0xa4, 0x4d, 0x27, 0xe5, 0xd4, 0xa3, 0xc6, 0x5f, 0xe5, 0x0f, 0x95, 0x9a, 0x42, 0x26, 0xbd, 0x7c,
	0x14, 0xe8, 0xa3, 0x25, 0xdb, 0xaf, 0x8c, 0x3e, 0x23, 0xd7, 0x4f, 0x2f, 0xed, 0xdb, 0xfa, 0xe5,
	0xdb, 0xba, 0x72, 0x5a, 0x0d, 0xd7, 0xe2, 0xce, 0xef, 0x0d, 0xd2, 0x49, 0xc0, 0xe4, 0x4a, 0x1a,
	0xa0, 0x77, 0x48, 0xb3, 0xdc, 0x71, 0x98, 0x7f, 0x8f, 0x55, 0x3d, 0xe3, 0xf7, 0xff, 0xd4, 0x3d,
	0x13, 0x0f, 0xc4, 0x13, 0x76, 0xdd, 0x76, 0xd3, 0x63, 0xeb, 0xc5, 0xb1, 0x36, 0x90, 0xcc, 0x6a,
	0xe4, 0xba, 0x09, 0x0e, 0x30, 0x1e, 0x2d, 0xe3, 0xe4, 0xc2, 0xac, 0xfa, 0x01, 0x0d, 0xd1, 0x0e,
	0xae, 0xc2, 0x49, 0xb9, 0x8a, 0xd7, 0x4e, 0x54, 0xf4, 0x9e, 0x3b, 0xf0, 0xef, 0x64, 0x0e, 0xa7,
	0x8f, 0x48, 0x43, 0xab, 0xa3, 0x68, 0xa3, 0x64, 0xdd, 0x66, 0xff, 0x34, 0x3e, 0x9b, 0x1f, 0x9e,
	0x25, 0xea, 0x28, 0x71, 0xbc, 0xde, 0xf7, 0x75, 0xd2, 0xc0, 0x80, 0x5e, 0x22, 0x2d, 0x0c, 0x9d,
	0x51, 0x3e, 0xc6, 0x38, 0x8f, 0x66, 0xd2, 0xc4, 0x10, 0x5d, 0xf0, 0x9c, 0x5c, 0x2c, 0xc7, 0x9c,
	0x72, 0x30, 0x63, 0x2d, 0xf2, 0xf2, 0xd0, 0x9f, 0xe2, 0xea, 0xc8, 0xc2, 0x36, 0x8d, 0xd5, 0x42,
	0x4e, 0xfc, 0x32, 0xbb, 0x25, 0xeb, 0xc9, 0x92, 0x44, 0x1f, 0x93, 0xad, 0xb1, 0x92, 0x5c, 0xb8,
	0xc0, 0xe9, 0x7c, 0x8e, 0xcf, 0x76, 0xe4, 0xe6, 0x82, 0x81, 0xad, 0xdc, 0x27, 0x6d, 0xa3, 0xb4,
	0x4d, 0xa5, 0x8a, 0xbe, 0xac, 0xc0, 0x6d, 0x39, 0x70, 0xac, 0xe8, 0xfe, 0x89, 0x5f, 0xe1, 0x6b,
	0xfc, 0xbf, 0xff, 0xc2, 0x03, 0xd2, 0xf1, 0x53, 0x40, 0xf6, 0xb7, 0x15, 0xd8, 0xed, 0x12, 0x3d,
	0xe2, 0xfb, 0x43, 0xd2, 0x17, 0xaa, 0xbe, 0x94, 0xc5, 0xed, 0xf5, 0x66, 0x77, 0xd5, 0x7b, 0xed,
	0xb0, 0x55, 0xde, 0x23, 0xf7, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x56, 0x2a, 0xc2, 0x00, 0x0a,
	0x05, 0x00, 0x00,
}
