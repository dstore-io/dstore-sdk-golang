// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampPersGroupCond_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampPersGroupCond_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampPersGroupCond_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampPersGroupCond_Ad

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
	FilterByPersonGroupIds         *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=filter_by_person_group_ids,json=filterByPersonGroupIds" json:"filter_by_person_group_ids,omitempty"`
	FilterByPersonGroupIdsNull     bool                        `protobuf:"varint,1001,opt,name=filter_by_person_group_ids_null,json=filterByPersonGroupIdsNull" json:"filter_by_person_group_ids_null,omitempty"`
	AnyIdInFilter                  *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=any_id_in_filter,json=anyIdInFilter" json:"any_id_in_filter,omitempty"`
	AnyIdInFilterNull              bool                        `protobuf:"varint,1002,opt,name=any_id_in_filter_null,json=anyIdInFilterNull" json:"any_id_in_filter_null,omitempty"`
	NegateFilter                   *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=negate_filter,json=negateFilter" json:"negate_filter,omitempty"`
	NegateFilterNull               bool                        `protobuf:"varint,1003,opt,name=negate_filter_null,json=negateFilterNull" json:"negate_filter_null,omitempty"`
	DeletePersonGroupCondition     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_person_group_condition,json=deletePersonGroupCondition" json:"delete_person_group_condition,omitempty"`
	DeletePersonGroupConditionNull bool                        `protobuf:"varint,1004,opt,name=delete_person_group_condition_null,json=deletePersonGroupConditionNull" json:"delete_person_group_condition_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetFilterByPersonGroupIds() *dstore_values.StringValue {
	if m != nil {
		return m.FilterByPersonGroupIds
	}
	return nil
}

func (m *Parameters) GetAnyIdInFilter() *dstore_values.BooleanValue {
	if m != nil {
		return m.AnyIdInFilter
	}
	return nil
}

func (m *Parameters) GetNegateFilter() *dstore_values.BooleanValue {
	if m != nil {
		return m.NegateFilter
	}
	return nil
}

func (m *Parameters) GetDeletePersonGroupCondition() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeletePersonGroupCondition
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampPersGroupCond_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampPersGroupCond_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampPersGroupCond_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x7b, 0x8b, 0xd3, 0x40,
	0x10, 0xe7, 0xae, 0xf7, 0x62, 0xef, 0x0e, 0xeb, 0x8a, 0x47, 0x4d, 0xb9, 0x53, 0xea, 0x3f, 0x82,
	0x98, 0xfa, 0x00, 0x51, 0x10, 0xd1, 0xab, 0x0f, 0x82, 0xf4, 0x28, 0x01, 0x0f, 0xf4, 0x0f, 0x97,
	0xed, 0xed, 0x36, 0x2c, 0x24, 0xbb, 0x61, 0x37, 0xf5, 0xe8, 0xb7, 0xf0, 0x2b, 0xf9, 0x71, 0x7c,
	0x7c, 0x05, 0xc1, 0xd9, 0x4c, 0x6a, 0x9b, 0x78, 0x0f, 0xf4, 0x9f, 0x36, 0x93, 0x99, 0xdf, 0x23,
	0x33, 0xb3, 0x4b, 0x9e, 0x09, 0x57, 0x18, 0x2b, 0xfb, 0x52, 0x27, 0x4a, 0xcb, 0x7e, 0x6e, 0xcd,
	0x89, 0x14, 0x53, 0x2b, 0x5d, 0xdf, 0x64, 0x6c, 0x68, 0x84, 0x9a, 0xcc, 0x06, 0x3c, 0xcb, 0x47,
	0xd2, 0xba, 0xb7, 0xd6, 0x4c, 0xf3, 0x81, 0xd1, 0x82, 0xbd, 0x14, 0x21, 0x14, 0x16, 0x86, 0xde,
	0x45, 0x74, 0x88, 0xe8, 0xf0, 0x42, 0x48, 0x70, 0xad, 0x92, 0xfa, 0xcc, 0xd3, 0xa9, 0x74, 0xc8,
	0x10, 0xdc, 0xa8, 0xeb, 0x4b, 0x6b, 0x8d, 0xad, 0x52, 0xdd, 0x7a, 0x2a, 0x93, 0xce, 0xf1, 0x44,
	0x56, 0xc9, 0xdb, 0xcd, 0x64, 0xc1, 0x95, 0x9e, 0x18, 0x9b, 0xf1, 0x42, 0x19, 0x8d, 0x45, 0xbd,
	0xaf, 0x6b, 0x84, 0x8c, 0xb8, 0xe5, 0x90, 0x05, 0x33, 0xf4, 0x98, 0x04, 0x13, 0x95, 0xc2, 0x23,
	0x1b, 0xcf, 0x58, 0x0e, 0x6f, 0x8c, 0x66, 0x89, 0x77, 0xc8, 0x94, 0x70, 0x9d, 0x95, 0x5b, 0x2b,
	0x77, 0xb6, 0x1f, 0x06, 0x61, 0xf5, 0x49, 0x95, 0x4b, 0x57, 0x58, 0xa5, 0x93, 0x63, 0x1f, 0xc4,
	0x7b, 0x88, 0x3e, 0x9c, 0x8d, 0x4a, 0x6c, 0xf9, 0x71, 0x91, 0x70, 0xf4, 0x15, 0xb9, 0x79, 0x3e,
	0x2f, 0xd3, 0xd3, 0x34, 0xed, 0x7c, 0xdb, 0x04, 0xf6, 0xad, 0x38, 0x38, 0x9b, 0xe1, 0x08, 0x4a,
	0x80, 0xa5, 0xcd, 0xf5, 0x0c, 0x20, 0x4c, 0x69, 0x86, 0x75, 0x9d, 0xd5, 0xd2, 0x53, 0xb7, 0xe1,
	0x69, 0x6c, 0x4c, 0x2a, 0xb9, 0x46, 0x53, 0xbb, 0x00, 0x8a, 0x44, 0xa4, 0xdf, 0x94, 0x08, 0xfa,
	0x80, 0x5c, 0x6f, 0xb2, 0xa0, 0x83, 0xef, 0xe8, 0xe0, 0x6a, 0xad, 0xbc, 0x14, 0x7e, 0x41, 0x76,
	0xb5, 0x4c, 0x78, 0x21, 0xe7, 0xaa, 0xad, 0xcb, 0x55, 0x77, 0x10, 0x51, 0x89, 0xde, 0x23, 0xb4,
	0xc6, 0x80, 0x8a, 0x3f, 0x50, 0xb1, 0xbd, 0x5c, 0x5a, 0x0a, 0x7e, 0x22, 0xfb, 0x42, 0xa6, 0x30,
	0x93, 0x7a, 0xb3, 0x4e, 0x60, 0x4f, 0x94, 0x9f, 0x5e, 0x67, 0xed, 0x72, 0x03, 0x01, 0x32, 0x2c,
	0xf5, 0x71, 0x30, 0x87, 0xd3, 0x77, 0xa4, 0x77, 0x21, 0x3f, 0xda, 0xfb, 0x89, 0xf6, 0x0e, 0xce,
	0x27, 0xf2, 0x66, 0x7b, 0xbf, 0x56, 0xc9, 0x56, 0x2c, 0x5d, 0x6e, 0xb4, 0x93, 0xf4, 0x3e, 0x59,
	0x2f, 0x37, 0xb4, 0xb9, 0x2c, 0xd5, 0xfe, 0xe3, 0xf6, 0xbe, 0xf6, 0xbf, 0x31, 0x16, 0xd2, 0x0f,
	0xa4, 0xed, 0x77, 0x93, 0x2d, 0x2d, 0x27, 0x4c, 0xb5, 0x05, 0xe0, 0xb0, 0x01, 0x6e, 0xae, 0xf0,
	0x10, 0xe2, 0x68, 0x11, 0xc7, 0x57, 0xb2, 0xfa, 0x0b, 0xfa, 0x84, 0x6c, 0x56, 0x67, 0x02, 0x26,
	0xe6, 0x19, 0x0f, 0xfe, 0x62, 0xc4, 0x13, 0x33, 0xc4, 0xff, 0x78, 0x5e, 0x0e, 0x0d, 0x6a, 0x59,
	0x73, 0x0a, 0x6d, 0xf6, 0xa8, 0xa7, 0xe1, 0x3f, 0x1c, 0xe2, 0x70, 0xde, 0x8a, 0x30, 0x36, 0xa7,
	0xb1, 0x67, 0xa1, 0xcf, 0xc9, 0xce, 0xa2, 0xb3, 0x4a, 0x74, 0xe4, 0x99, 0xc3, 0x53, 0xba, 0x90,
	0x89, 0xb4, 0x38, 0xbc, 0xed, 0x3f, 0x80, 0x48, 0x04, 0xfb, 0xa4, 0x05, 0x5c, 0x74, 0x8f, 0x6c,
	0x00, 0x9b, 0x27, 0xf8, 0x72, 0x04, 0x0c, 0xeb, 0xf1, 0x3a, 0x84, 0x91, 0x38, 0x7c, 0x4f, 0xba,
	0xca, 0x34, 0x2c, 0x2e, 0x6e, 0xa9, 0x8f, 0x8f, 0xff, 0xef, 0xfe, 0x1a, 0x6f, 0x94, 0x37, 0xc4,
	0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x6b, 0x7f, 0x2e, 0x00, 0x05, 0x00, 0x00,
}
