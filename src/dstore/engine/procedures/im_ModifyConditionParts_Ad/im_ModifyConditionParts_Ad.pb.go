// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyConditionParts_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyConditionParts_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyConditionParts_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyConditionParts_Ad

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
	ConditionPartDescription     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=condition_part_description,json=conditionPartDescription" json:"condition_part_description,omitempty"`
	ConditionPartDescriptionNull bool                        `protobuf:"varint,1001,opt,name=condition_part_description_null,json=conditionPartDescriptionNull" json:"condition_part_description_null,omitempty"`
	LevelIds                     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=level_ids,json=levelIds" json:"level_ids,omitempty"`
	LevelIdsNull                 bool                        `protobuf:"varint,1002,opt,name=level_ids_null,json=levelIdsNull" json:"level_ids_null,omitempty"`
	DomainTreeNodeIds            *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=domain_tree_node_ids,json=domainTreeNodeIds" json:"domain_tree_node_ids,omitempty"`
	DomainTreeNodeIdsNull        bool                        `protobuf:"varint,1003,opt,name=domain_tree_node_ids_null,json=domainTreeNodeIdsNull" json:"domain_tree_node_ids_null,omitempty"`
	NodeCharacteristicId         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull     bool                        `protobuf:"varint,1004,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	Operator1                    *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=operator1" json:"operator1,omitempty"`
	Operator1Null                bool                        `protobuf:"varint,1005,opt,name=operator1_null,json=operator1Null" json:"operator1_null,omitempty"`
	Condition1                   *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=condition1" json:"condition1,omitempty"`
	Condition1Null               bool                        `protobuf:"varint,1006,opt,name=condition1_null,json=condition1Null" json:"condition1_null,omitempty"`
	Operator2                    *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=operator2" json:"operator2,omitempty"`
	Operator2Null                bool                        `protobuf:"varint,1007,opt,name=operator2_null,json=operator2Null" json:"operator2_null,omitempty"`
	Condition2                   *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=condition2" json:"condition2,omitempty"`
	Condition2Null               bool                        `protobuf:"varint,1008,opt,name=condition2_null,json=condition2Null" json:"condition2_null,omitempty"`
	InheritDepth                 *dstore_values.IntegerValue `protobuf:"bytes,9,opt,name=inherit_depth,json=inheritDepth" json:"inherit_depth,omitempty"`
	InheritDepthNull             bool                        `protobuf:"varint,1009,opt,name=inherit_depth_null,json=inheritDepthNull" json:"inherit_depth_null,omitempty"`
	RecursiveEvaluation          *dstore_values.IntegerValue `protobuf:"bytes,10,opt,name=recursive_evaluation,json=recursiveEvaluation" json:"recursive_evaluation,omitempty"`
	RecursiveEvaluationNull      bool                        `protobuf:"varint,1010,opt,name=recursive_evaluation_null,json=recursiveEvaluationNull" json:"recursive_evaluation_null,omitempty"`
	DeleteConditionPart          *dstore_values.BooleanValue `protobuf:"bytes,11,opt,name=delete_condition_part,json=deleteConditionPart" json:"delete_condition_part,omitempty"`
	DeleteConditionPartNull      bool                        `protobuf:"varint,1011,opt,name=delete_condition_part_null,json=deleteConditionPartNull" json:"delete_condition_part_null,omitempty"`
	Country                      *dstore_values.StringValue  `protobuf:"bytes,12,opt,name=country" json:"country,omitempty"`
	CountryNull                  bool                        `protobuf:"varint,1012,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetConditionPartDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionPartDescription
	}
	return nil
}

func (m *Parameters) GetLevelIds() *dstore_values.StringValue {
	if m != nil {
		return m.LevelIds
	}
	return nil
}

func (m *Parameters) GetDomainTreeNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.DomainTreeNodeIds
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetOperator1() *dstore_values.StringValue {
	if m != nil {
		return m.Operator1
	}
	return nil
}

func (m *Parameters) GetCondition1() *dstore_values.StringValue {
	if m != nil {
		return m.Condition1
	}
	return nil
}

func (m *Parameters) GetOperator2() *dstore_values.StringValue {
	if m != nil {
		return m.Operator2
	}
	return nil
}

func (m *Parameters) GetCondition2() *dstore_values.StringValue {
	if m != nil {
		return m.Condition2
	}
	return nil
}

func (m *Parameters) GetInheritDepth() *dstore_values.IntegerValue {
	if m != nil {
		return m.InheritDepth
	}
	return nil
}

func (m *Parameters) GetRecursiveEvaluation() *dstore_values.IntegerValue {
	if m != nil {
		return m.RecursiveEvaluation
	}
	return nil
}

func (m *Parameters) GetDeleteConditionPart() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteConditionPart
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

type Response struct {
	Error                         *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation               []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message                       []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row                           []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ConditionPartId               *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=condition_part_id,json=conditionPartId" json:"condition_part_id,omitempty"`
	ConditionPartIdWithSameConfig *dstore_values.IntegerValue                      `protobuf:"bytes,102,opt,name=condition_part_id_with_same_config,json=conditionPartIdWithSameConfig" json:"condition_part_id_with_same_config,omitempty"`
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

func (m *Response) GetConditionPartId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionPartId
	}
	return nil
}

func (m *Response) GetConditionPartIdWithSameConfig() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionPartIdWithSameConfig
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyConditionParts_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyConditionParts_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyConditionParts_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xeb, 0x6e, 0xd3, 0x48,
	0x14, 0x56, 0x37, 0x9b, 0x26, 0x3d, 0x49, 0x6f, 0xd3, 0xcb, 0xba, 0xc9, 0x76, 0x77, 0x95, 0xd5,
	0xae, 0xfa, 0x67, 0xdd, 0x6d, 0xba, 0x5a, 0x28, 0x20, 0x04, 0xb4, 0x05, 0x22, 0xd4, 0x00, 0x06,
	0x71, 0xfb, 0x63, 0xb9, 0xf6, 0x24, 0x1d, 0x29, 0xf1, 0x44, 0x33, 0x4e, 0xab, 0xbe, 0x05, 0x8f,
	0xc0, 0xeb, 0x41, 0xb9, 0xc3, 0x03, 0x70, 0xc6, 0xe3, 0x5c, 0xec, 0xb8, 0x72, 0xe1, 0x4f, 0x9b,
	0x99, 0x39, 0xdf, 0xe5, 0xcc, 0xb1, 0xcf, 0x31, 0xec, 0x78, 0x32, 0xe0, 0x82, 0x6e, 0x52, 0xbf,
	0xcd, 0x7c, 0xba, 0xd9, 0x13, 0xdc, 0xa5, 0x5e, 0x5f, 0x50, 0xb9, 0xc9, 0xba, 0xf6, 0x01, 0xf7,
	0x58, 0xeb, 0x74, 0x97, 0xfb, 0x1e, 0x0b, 0x18, 0xf7, 0x1f, 0x38, 0x22, 0x90, 0xf6, 0x4d, 0xcf,
	0xc4, 0xa8, 0x80, 0x93, 0x0d, 0x0d, 0x35, 0x35, 0xd4, 0x3c, 0x3f, 0xbe, 0xb2, 0x14, 0x89, 0x1c,
	0x3b, 0x9d, 0x3e, 0x95, 0x1a, 0x5e, 0x59, 0x8b, 0x2b, 0x53, 0x21, 0xb8, 0x88, 0x8e, 0xaa, 0xf1,
	0xa3, 0x2e, 0x95, 0xd2, 0x69, 0xd3, 0xe8, 0xf0, 0xcf, 0xe4, 0x61, 0xe0, 0x30, 0xbf, 0xc5, 0x45,
	0xd7, 0x51, 0x9a, 0x3a, 0xa8, 0xf6, 0xaa, 0x04, 0x80, 0xf2, 0x0e, 0x9e, 0x52, 0x21, 0xc9, 0x33,
	0xa8, 0xb8, 0x03, 0x57, 0x76, 0x0f, 0x6d, 0xd9, 0x1e, 0x95, 0xae, 0x60, 0x3d, 0xb5, 0x61, 0x4c,
	0xfd, 0x31, 0xb5, 0x51, 0xaa, 0x57, 0xcc, 0x28, 0x9f, 0xc8, 0xa5, 0x0c, 0x04, 0xf3, 0xdb, 0x4f,
	0xd4, 0xc2, 0x32, 0xdc, 0xf1, 0x9c, 0xf6, 0x46, 0x58, 0x72, 0x1b, 0x7e, 0x3f, 0x9f, 0xd9, 0xf6,
	0xfb, 0x9d, 0x8e, 0xf1, 0xba, 0x80, 0xfc, 0x45, 0xeb, 0xd7, 0xf3, 0x38, 0x9a, 0x18, 0x44, 0x2e,
	0xc1, 0x4c, 0x87, 0x1e, 0xd3, 0x8e, 0xcd, 0x3c, 0x69, 0xfc, 0x94, 0x69, 0xa8, 0x18, 0x06, 0x37,
	0x3c, 0x49, 0xfe, 0x82, 0xb9, 0x21, 0x50, 0xeb, 0xbd, 0xd1, 0x7a, 0xe5, 0x41, 0x48, 0xc8, 0x7f,
	0x0f, 0x96, 0x3d, 0xde, 0xc5, 0x9b, 0xb2, 0x03, 0x41, 0xa9, 0xed, 0x73, 0x8f, 0x86, 0x52, 0xb9,
	0x4c, 0xa9, 0x45, 0x8d, 0x7b, 0x8c, 0xb0, 0x26, 0xa2, 0x94, 0xe6, 0x0e, 0xac, 0xa5, 0x91, 0x69,
	0xf9, 0x33, 0x2d, 0xbf, 0x32, 0x01, 0x0b, 0x7d, 0x3c, 0x84, 0xd5, 0x30, 0xdc, 0x3d, 0xc2, 0xea,
	0xb8, 0x58, 0x1c, 0x26, 0x03, 0xe6, 0x22, 0xda, 0xf8, 0x39, 0x74, 0x52, 0x4d, 0x38, 0x61, 0x7e,
	0x40, 0xdb, 0x54, 0x68, 0x2b, 0xcb, 0x0a, 0xba, 0x1b, 0x43, 0x36, 0x3c, 0x72, 0x1d, 0xaa, 0xe9,
	0x94, 0xda, 0xcf, 0x5b, 0xed, 0xc7, 0x48, 0xc3, 0x86, 0x96, 0x2e, 0xc3, 0x0c, 0xef, 0x51, 0xe1,
	0xa0, 0xec, 0x96, 0x91, 0xcf, 0xbc, 0x8f, 0x51, 0x30, 0xf9, 0x1b, 0xe6, 0x86, 0x0b, 0x2d, 0xf6,
	0x4e, 0x8b, 0xcd, 0x0e, 0xb7, 0x43, 0x85, 0x2b, 0x00, 0xc3, 0xe2, 0x6f, 0x19, 0xd3, 0x99, 0x12,
	0x63, 0xd1, 0x64, 0x03, 0xe6, 0x47, 0x2b, 0x2d, 0xf2, 0x5e, 0x8b, 0xcc, 0x8d, 0xf6, 0x93, 0x79,
	0xd4, 0x8d, 0xc2, 0xc5, 0xf3, 0xa8, 0x8f, 0xe7, 0x51, 0xd7, 0x12, 0x1f, 0x12, 0x79, 0xd4, 0x27,
	0xf2, 0xa8, 0x1b, 0xc5, 0xef, 0xc8, 0xa3, 0x1e, 0xcb, 0x23, 0x12, 0xf9, 0x98, 0xcc, 0x43, 0xab,
	0xdc, 0x80, 0x59, 0xe6, 0x1f, 0x61, 0x8d, 0xd4, 0xbb, 0xd4, 0x0b, 0x8e, 0x8c, 0x99, 0xec, 0x27,
	0xa3, 0x1c, 0x21, 0xf6, 0x14, 0x80, 0xfc, 0x03, 0x24, 0xc6, 0xa0, 0xe5, 0x3e, 0x69, 0xb9, 0x85,
	0xf1, 0xd0, 0x50, 0xb0, 0x09, 0xcb, 0x82, 0xba, 0x7d, 0x21, 0xd9, 0x31, 0xb5, 0xa9, 0xa2, 0x0f,
	0x5b, 0x89, 0x01, 0xd9, 0xba, 0x4b, 0x43, 0xe0, 0xfe, 0x10, 0x47, 0xae, 0xc2, 0x5a, 0x1a, 0x9f,
	0x76, 0xf1, 0x59, 0xbb, 0xf8, 0x25, 0x05, 0x18, 0x9a, 0xb9, 0x0f, 0x2b, 0x1e, 0xed, 0x60, 0xdb,
	0xb2, 0xe3, 0x7d, 0xc5, 0x28, 0xa5, 0xba, 0x39, 0xe4, 0xbc, 0x43, 0x1d, 0x3f, 0x72, 0xa3, 0x91,
	0xb1, 0x06, 0x4c, 0xae, 0x41, 0x25, 0x95, 0x50, 0xdb, 0xf9, 0x12, 0xd9, 0x49, 0x41, 0x86, 0x76,
	0xfe, 0x83, 0x82, 0xcb, 0xfb, 0x7e, 0x20, 0x4e, 0x8d, 0x72, 0x66, 0xbd, 0x07, 0xa1, 0xa4, 0x06,
	0xe5, 0xe8, 0xa7, 0x56, 0xf9, 0xaa, 0x55, 0x4a, 0xd1, 0xa6, 0x62, 0xae, 0x9d, 0xe5, 0xa0, 0x68,
	0x51, 0xd9, 0xe3, 0xbe, 0xa4, 0xe4, 0x5f, 0xc8, 0x87, 0x03, 0x20, 0xd9, 0x8b, 0xa3, 0xd9, 0xa2,
	0x87, 0xc3, 0xbe, 0xfa, 0x6b, 0xe9, 0x40, 0xf2, 0x1c, 0x16, 0x54, 0xeb, 0xb7, 0xc7, 0x7a, 0x3f,
	0xf6, 0xcd, 0x1c, 0x82, 0xcd, 0x04, 0x38, 0x39, 0x21, 0x0e, 0x70, 0xdd, 0x18, 0xad, 0xad, 0xf9,
	0x6e, 0x7c, 0x03, 0x5f, 0xa4, 0x42, 0x34, 0x72, 0xb0, 0x3d, 0x2a, 0xc6, 0xdf, 0x26, 0x18, 0xf5,
	0x40, 0x3a, 0xd0, 0xff, 0xad, 0x41, 0x38, 0xb9, 0x0b, 0x39, 0xc1, 0x4f, 0xb0, 0x95, 0x29, 0xd4,
	0xff, 0xe6, 0x45, 0x07, 0xa4, 0x39, 0xb8, 0x07, 0xd3, 0xe2, 0x27, 0x96, 0xa2, 0x20, 0x77, 0x60,
	0x31, 0x51, 0x2e, 0x6c, 0x91, 0x34, 0xfb, 0x81, 0x9c, 0x8f, 0x4d, 0x19, 0xec, 0x8e, 0x2d, 0xa8,
	0x4d, 0x10, 0xd9, 0x27, 0x0c, 0x5f, 0x0a, 0x89, 0xd3, 0x51, 0x3d, 0x13, 0x2d, 0xd6, 0x36, 0x5a,
	0xd9, 0xcc, 0xeb, 0x09, 0xe6, 0xa7, 0xc8, 0xf1, 0x08, 0x29, 0x76, 0x43, 0x86, 0xca, 0x3a, 0xe4,
	0xd0, 0x3c, 0x59, 0x85, 0x69, 0xb4, 0xaf, 0xcc, 0xbe, 0x6c, 0x22, 0x67, 0xde, 0xca, 0xe3, 0xb2,
	0xe1, 0xdd, 0xb2, 0xa0, 0xca, 0x78, 0xe2, 0x42, 0x46, 0x1f, 0x1b, 0x2f, 0xb6, 0x7f, 0xe0, 0x33,
	0xe4, 0x70, 0x3a, 0x9c, 0xf5, 0xdb, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x03, 0xb6, 0x85, 0x3b,
	0xc4, 0x08, 0x00, 0x00,
}
