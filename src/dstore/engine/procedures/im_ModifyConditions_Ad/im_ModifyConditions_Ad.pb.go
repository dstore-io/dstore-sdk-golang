// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyConditions_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyConditions_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyConditions_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyConditions_Ad

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
	ConditionDescription             *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=condition_description,json=conditionDescription" json:"condition_description,omitempty"`
	ConditionDescriptionNull         bool                        `protobuf:"varint,1001,opt,name=condition_description_null,json=conditionDescriptionNull" json:"condition_description_null,omitempty"`
	CombineGroupsWithANDOperator     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=combine_groups_with_a_n_d_operator,json=combineGroupsWithANDOperator" json:"combine_groups_with_a_n_d_operator,omitempty"`
	CombineGroupsWithANDOperatorNull bool                        `protobuf:"varint,1002,opt,name=combine_groups_with_a_n_d_operator_null,json=combineGroupsWithANDOperatorNull" json:"combine_groups_with_a_n_d_operator_null,omitempty"`
	DeleteCondition                  *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete_condition,json=deleteCondition" json:"delete_condition,omitempty"`
	DeleteConditionNull              bool                        `protobuf:"varint,1003,opt,name=delete_condition_null,json=deleteConditionNull" json:"delete_condition_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionDescription
	}
	return nil
}

func (m *Parameters) GetCombineGroupsWithANDOperator() *dstore_values.BooleanValue {
	if m != nil {
		return m.CombineGroupsWithANDOperator
	}
	return nil
}

func (m *Parameters) GetDeleteCondition() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCondition
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ConditionId     *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
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

func (m *Response) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyConditions_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyConditions_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyConditions_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xb5, 0x85, 0x6e, 0xd3, 0x1b, 0xd2, 0x26, 0x8f, 0xa1, 0x90, 0x02, 0x9a, 0x0a, 0x12,
	0x9c, 0x52, 0xb4, 0x09, 0x89, 0x0b, 0x48, 0x83, 0x02, 0xea, 0xa1, 0x1d, 0xf2, 0x01, 0x04, 0x17,
	0x2b, 0xad, 0xdf, 0x82, 0xa5, 0x24, 0x8e, 0xec, 0x94, 0x8a, 0x6f, 0xc1, 0x87, 0xe3, 0x4b, 0x00,
	0x77, 0xce, 0xbc, 0xc4, 0x6e, 0x4b, 0x43, 0xc5, 0xe0, 0xd2, 0xd6, 0x7d, 0xef, 0xff, 0xf3, 0x3f,
	0xcf, 0xff, 0x18, 0x1e, 0x4b, 0x5b, 0x69, 0x83, 0x7d, 0x2c, 0x52, 0x55, 0x60, 0xbf, 0x34, 0x7a,
	0x8a, 0x72, 0x66, 0xd0, 0xf6, 0x55, 0x2e, 0x46, 0x5a, 0xaa, 0xcb, 0xcf, 0x2f, 0x74, 0x21, 0x55,
	0xa5, 0x74, 0x61, 0xc5, 0xb9, 0x8c, 0xa9, 0xa3, 0xd2, 0xec, 0xbe, 0x93, 0xc5, 0x4e, 0x16, 0x6f,
	0xee, 0x8d, 0x8e, 0x3c, 0xfc, 0x53, 0x92, 0xcd, 0xd0, 0x3a, 0x69, 0x74, 0x6b, 0x7d, 0x47, 0x34,
	0x46, 0x1b, 0x5f, 0xea, 0xae, 0x97, 0x72, 0xb4, 0x36, 0x49, 0xd1, 0x17, 0xef, 0xb5, 0x8b, 0x55,
	0xa2, 0x8a, 0x4b, 0x6d, 0xf2, 0xa4, 0xde, 0xcf, 0x35, 0xf5, 0xbe, 0x06, 0x00, 0x6f, 0x12, 0x93,
	0x50, 0x15, 0x8d, 0x65, 0x17, 0x70, 0x3c, 0x5d, 0x38, 0x12, 0x12, 0xed, 0xd4, 0xa8, 0xb2, 0xfe,
	0x1d, 0x6e, 0x9d, 0x6c, 0x3d, 0xdc, 0x3f, 0x8d, 0x62, 0xff, 0x18, 0xde, 0xa0, 0xad, 0x8c, 0x2a,
	0xd2, 0xb7, 0xf5, 0x82, 0xdf, 0x58, 0x0a, 0x07, 0x2b, 0x1d, 0x7b, 0x0a, 0xd1, 0x46, 0xa0, 0x28,
	0x66, 0x59, 0x16, 0x7e, 0xdb, 0x25, 0xec, 0x1e, 0x0f, 0x37, 0x49, 0xc7, 0xd4, 0xc0, 0x10, 0x7a,
	0x53, 0x9d, 0x4f, 0xc8, 0xbf, 0x48, 0x8d, 0x9e, 0x95, 0x56, 0xcc, 0x55, 0xf5, 0x51, 0x24, 0x82,
	0x60, 0x42, 0x97, 0x68, 0x12, 0xb2, 0x13, 0x6e, 0x37, 0xe6, 0xba, 0x2d, 0x73, 0x13, 0xad, 0x33,
	0x4c, 0x0a, 0xe7, 0xee, 0xb6, 0xc7, 0xbc, 0x6e, 0x28, 0xef, 0x08, 0x72, 0x3e, 0x1e, 0x5c, 0x78,
	0x00, 0xe3, 0xf0, 0xe0, 0xea, 0x6d, 0x9c, 0xe5, 0xef, 0xce, 0xf2, 0xc9, 0xdf, 0x78, 0x8d, 0xf5,
	0x57, 0x70, 0x28, 0x31, 0xa3, 0xb1, 0x8a, 0xe5, 0xd3, 0x85, 0xc1, 0xd5, 0x46, 0x0f, 0x9c, 0x68,
	0x99, 0x0b, 0x76, 0x06, 0xc7, 0x6d, 0x8e, 0x73, 0xf2, 0xc3, 0x39, 0x39, 0x6a, 0x09, 0xea, 0xcd,
	0x7b, 0x3f, 0xb7, 0x61, 0x8f, 0xa3, 0x2d, 0x29, 0x57, 0xc8, 0x1e, 0x41, 0xa7, 0x09, 0x4d, 0xfb,
	0x10, 0x7d, 0x16, 0x5d, 0xa0, 0x5e, 0xd6, 0x9f, 0xdc, 0x35, 0xb2, 0xf7, 0x70, 0x58, 0xc7, 0x45,
	0xfc, 0x96, 0x17, 0x1a, 0x72, 0x40, 0xe2, 0xb8, 0x25, 0x6e, 0xa7, 0x6a, 0x44, 0xeb, 0xe1, 0x6a,
	0xcd, 0x0f, 0xf2, 0xf5, 0x3f, 0xd8, 0x13, 0xd8, 0xf5, 0x31, 0xa5, 0x69, 0xd4, 0xc4, 0xbb, 0x7f,
	0x10, 0x5d, 0x88, 0x47, 0xee, 0x9b, 0x2f, 0xda, 0xd9, 0x00, 0x02, 0xa3, 0xe7, 0xe1, 0xb5, 0x46,
	0x75, 0x1a, 0xff, 0xcb, 0x0b, 0x15, 0x2f, 0x66, 0x10, 0x73, 0x3d, 0xe7, 0xb5, 0x9c, 0x3d, 0x83,
	0xeb, 0xab, 0x39, 0x2a, 0x19, 0xe2, 0xc6, 0x23, 0x51, 0x45, 0x85, 0x29, 0x1a, 0x77, 0x24, 0xfb,
	0x4b, 0xc1, 0x50, 0x46, 0x77, 0x20, 0x20, 0x16, 0xbb, 0x09, 0x3b, 0x44, 0xab, 0x01, 0x5f, 0xc6,
	0x44, 0xe8, 0xf0, 0x0e, 0x2d, 0x87, 0xf2, 0xf9, 0x18, 0xba, 0x4a, 0xb7, 0xbc, 0xad, 0xee, 0x88,
	0x0f, 0xfd, 0xff, 0xbc, 0x3d, 0x26, 0x3b, 0xcd, 0x6b, 0x7a, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x01, 0xf2, 0x19, 0xa4, 0x77, 0x04, 0x00, 0x00,
}
