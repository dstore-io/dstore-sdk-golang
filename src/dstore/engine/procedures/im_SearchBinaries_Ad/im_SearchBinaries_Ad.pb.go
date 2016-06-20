// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_SearchBinaries_Ad.proto
// DO NOT EDIT!

/*
Package im_SearchBinaries_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_SearchBinaries_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_SearchBinaries_Ad

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
	AssignedToTable              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=assigned_to_table,json=assignedToTable" json:"assigned_to_table,omitempty"`
	AssignedToTableNull          bool                        `protobuf:"varint,1001,opt,name=assigned_to_table_null,json=assignedToTableNull" json:"assigned_to_table_null,omitempty"`
	InputNestLevelConds          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=input_nest_level_conds,json=inputNestLevelConds" json:"input_nest_level_conds,omitempty"`
	InputNestLevelCondsNull      bool                        `protobuf:"varint,1002,opt,name=input_nest_level_conds_null,json=inputNestLevelCondsNull" json:"input_nest_level_conds_null,omitempty"`
	ValueSeparatorINOperator     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value_separator_i_n_operator,json=valueSeparatorINOperator" json:"value_separator_i_n_operator,omitempty"`
	ValueSeparatorINOperatorNull bool                        `protobuf:"varint,1003,opt,name=value_separator_i_n_operator_null,json=valueSeparatorINOperatorNull" json:"value_separator_i_n_operator_null,omitempty"`
	StartAtRowNo                 *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=start_at_row_no,json=startAtRowNo" json:"start_at_row_no,omitempty"`
	StartAtRowNoNull             bool                        `protobuf:"varint,1004,opt,name=start_at_row_no_null,json=startAtRowNoNull" json:"start_at_row_no_null,omitempty"`
	RowCount                     *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=row_count,json=rowCount" json:"row_count,omitempty"`
	RowCountNull                 bool                        `protobuf:"varint,1005,opt,name=row_count_null,json=rowCountNull" json:"row_count_null,omitempty"`
	OutputBinaryCharacId1        *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=output_binary_charac_id1,json=outputBinaryCharacId1" json:"output_binary_charac_id1,omitempty"`
	OutputBinaryCharacId1Null    bool                        `protobuf:"varint,1006,opt,name=output_binary_charac_id1_null,json=outputBinaryCharacId1Null" json:"output_binary_charac_id1_null,omitempty"`
	OutputBinaryCharacId2        *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=output_binary_charac_id2,json=outputBinaryCharacId2" json:"output_binary_charac_id2,omitempty"`
	OutputBinaryCharacId2Null    bool                        `protobuf:"varint,1007,opt,name=output_binary_charac_id2_null,json=outputBinaryCharacId2Null" json:"output_binary_charac_id2_null,omitempty"`
	OutputBinaryCharacId3        *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=output_binary_charac_id3,json=outputBinaryCharacId3" json:"output_binary_charac_id3,omitempty"`
	OutputBinaryCharacId3Null    bool                        `protobuf:"varint,1008,opt,name=output_binary_charac_id3_null,json=outputBinaryCharacId3Null" json:"output_binary_charac_id3_null,omitempty"`
	IncludeBinaryCode            *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=include_binary_code,json=includeBinaryCode" json:"include_binary_code,omitempty"`
	IncludeBinaryCodeNull        bool                        `protobuf:"varint,1009,opt,name=include_binary_code_null,json=includeBinaryCodeNull" json:"include_binary_code_null,omitempty"`
	NodeCharacteristicId1        *dstore_values.IntegerValue `protobuf:"bytes,10,opt,name=node_characteristic_id1,json=nodeCharacteristicId1" json:"node_characteristic_id1,omitempty"`
	NodeCharacteristicId1Null    bool                        `protobuf:"varint,1010,opt,name=node_characteristic_id1_null,json=nodeCharacteristicId1Null" json:"node_characteristic_id1_null,omitempty"`
	NodeCharacteristicId2        *dstore_values.IntegerValue `protobuf:"bytes,11,opt,name=node_characteristic_id2,json=nodeCharacteristicId2" json:"node_characteristic_id2,omitempty"`
	NodeCharacteristicId2Null    bool                        `protobuf:"varint,1011,opt,name=node_characteristic_id2_null,json=nodeCharacteristicId2Null" json:"node_characteristic_id2_null,omitempty"`
	NodeCharacteristicId3        *dstore_values.IntegerValue `protobuf:"bytes,12,opt,name=node_characteristic_id3,json=nodeCharacteristicId3" json:"node_characteristic_id3,omitempty"`
	NodeCharacteristicId3Null    bool                        `protobuf:"varint,1012,opt,name=node_characteristic_id3_null,json=nodeCharacteristicId3Null" json:"node_characteristic_id3_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetAssignedToTable() *dstore_values.IntegerValue {
	if m != nil {
		return m.AssignedToTable
	}
	return nil
}

func (m *Parameters) GetInputNestLevelConds() *dstore_values.IntegerValue {
	if m != nil {
		return m.InputNestLevelConds
	}
	return nil
}

func (m *Parameters) GetValueSeparatorINOperator() *dstore_values.StringValue {
	if m != nil {
		return m.ValueSeparatorINOperator
	}
	return nil
}

func (m *Parameters) GetStartAtRowNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.StartAtRowNo
	}
	return nil
}

func (m *Parameters) GetRowCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowCount
	}
	return nil
}

func (m *Parameters) GetOutputBinaryCharacId1() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputBinaryCharacId1
	}
	return nil
}

func (m *Parameters) GetOutputBinaryCharacId2() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputBinaryCharacId2
	}
	return nil
}

func (m *Parameters) GetOutputBinaryCharacId3() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputBinaryCharacId3
	}
	return nil
}

func (m *Parameters) GetIncludeBinaryCode() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeBinaryCode
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId1() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId1
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId2() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId2
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId3() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId3
	}
	return nil
}

type Response struct {
	Error                    *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation          []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message                  []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row                      []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	NumberOfElementsInResult *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=number_of_elements_in_result,json=numberOfElementsInResult" json:"number_of_elements_in_result,omitempty"`
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

func (m *Response) GetNumberOfElementsInResult() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfElementsInResult
	}
	return nil
}

type Response_Row struct {
	RowId                int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BinaryCodeId         *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	NodeId               *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	ChosenTreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=chosen_tree_node_id,json=chosenTreeNodeId" json:"chosen_tree_node_id,omitempty"`
	Value3               *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=value3" json:"value3,omitempty"`
	BinaryValue1         *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=binary_value1,json=binaryValue1" json:"binary_value1,omitempty"`
	Value1               *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=value1" json:"value1,omitempty"`
	BinaryValue3         *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=binary_value3,json=binaryValue3" json:"binary_value3,omitempty"`
	Value2               *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=value2" json:"value2,omitempty"`
	BinaryValue2         *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=binary_value2,json=binaryValue2" json:"binary_value2,omitempty"`
	ThumbnailCode        *dstore_values.BytesValue   `protobuf:"bytes,20008,opt,name=thumbnail_code,json=thumbnailCode" json:"thumbnail_code,omitempty"`
	BinaryCode           *dstore_values.BytesValue   `protobuf:"bytes,20011,opt,name=binary_code,json=binaryCode" json:"binary_code,omitempty"`
	NodeCharacteristicId *dstore_values.IntegerValue `protobuf:"bytes,30002,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	Value                *dstore_values.StringValue  `protobuf:"bytes,30003,opt,name=value" json:"value,omitempty"`
	ValueId              *dstore_values.IntegerValue `protobuf:"bytes,30005,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetChosenTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ChosenTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func (m *Response_Row) GetBinaryValue1() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryValue1
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetBinaryValue3() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryValue3
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetBinaryValue2() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryValue2
	}
	return nil
}

func (m *Response_Row) GetThumbnailCode() *dstore_values.BytesValue {
	if m != nil {
		return m.ThumbnailCode
	}
	return nil
}

func (m *Response_Row) GetBinaryCode() *dstore_values.BytesValue {
	if m != nil {
		return m.BinaryCode
	}
	return nil
}

func (m *Response_Row) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_SearchBinaries_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_SearchBinaries_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_SearchBinaries_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 1021 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x97, 0x5b, 0x6f, 0x23, 0x35,
	0x14, 0xc7, 0x55, 0x4a, 0x93, 0xac, 0xdb, 0x6e, 0xbb, 0xd3, 0xdd, 0xae, 0x7b, 0x01, 0x41, 0x11,
	0x12, 0x2f, 0xa4, 0x9b, 0x19, 0x90, 0x78, 0x00, 0x89, 0xb6, 0x5a, 0x41, 0xc4, 0x6e, 0x76, 0x35,
	0xad, 0x90, 0x58, 0x1e, 0xac, 0x49, 0xc6, 0x4d, 0x47, 0x9a, 0xd8, 0xd1, 0xd8, 0xd9, 0x55, 0xdf,
	0x79, 0xe7, 0x7e, 0x93, 0xf8, 0x00, 0x48, 0x3c, 0x20, 0x81, 0xf8, 0x50, 0x94, 0xfb, 0xe5, 0x03,
	0xec, 0xb1, 0x8f, 0x93, 0xb6, 0xd3, 0x64, 0x26, 0xed, 0x4b, 0x2b, 0x8f, 0xcf, 0xff, 0x7f, 0x7e,
	0x76, 0xec, 0x73, 0x66, 0x48, 0x10, 0x2b, 0x2d, 0x33, 0xbe, 0xcd, 0x45, 0x37, 0x11, 0x7c, 0xbb,
	0x9f, 0xc9, 0x0e, 0x8f, 0x07, 0x19, 0x57, 0xdb, 0x49, 0x8f, 0xed, 0xf3, 0x28, 0xeb, 0x1c, 0xed,
	0x26, 0x22, 0xca, 0x12, 0xae, 0xd8, 0x4e, 0x5c, 0x87, 0x79, 0x2d, 0xbd, 0x2d, 0x14, 0xd5, 0x51,
	0x54, 0x1f, 0x17, 0xb9, 0xbe, 0xe2, 0x8c, 0x1f, 0x47, 0xe9, 0x80, 0x2b, 0x14, 0xae, 0xaf, 0x9d,
	0xcf, 0xc6, 0xb3, 0x4c, 0x66, 0x6e, 0x6a, 0xe3, 0xfc, 0x54, 0x8f, 0x2b, 0x15, 0x75, 0xb9, 0x9b,
	0x7c, 0x29, 0x3f, 0xa9, 0xa3, 0x44, 0x1c, 0xca, 0xac, 0x17, 0xe9, 0x44, 0x0a, 0x0c, 0xda, 0xfa,
	0x71, 0x91, 0x90, 0x87, 0x51, 0x16, 0xc1, 0x2c, 0xcf, 0x94, 0xf7, 0x0e, 0xb9, 0x11, 0x29, 0x95,
	0x74, 0x05, 0x8f, 0x99, 0x96, 0x4c, 0x47, 0xed, 0x94, 0xd3, 0x99, 0x17, 0x66, 0x5e, 0x99, 0xf7,
	0x37, 0xea, 0x6e, 0x01, 0x0e, 0x2e, 0x11, 0x9a, 0x77, 0x79, 0xf6, 0xbe, 0x19, 0x85, 0x4b, 0x43,
	0xd5, 0x81, 0x3c, 0x30, 0x1a, 0xef, 0x35, 0xb2, 0x7a, 0xc1, 0x88, 0x89, 0x41, 0x9a, 0xd2, 0x5f,
	0xab, 0x60, 0x57, 0x0b, 0x57, 0x72, 0x8a, 0x16, 0xcc, 0x79, 0x0f, 0xc9, 0x6a, 0x22, 0xfa, 0x03,
	0xcd, 0x04, 0x57, 0x9a, 0xa5, 0xfc, 0x31, 0x4f, 0x59, 0x47, 0x8a, 0x58, 0xd1, 0x67, 0xca, 0x19,
	0x56, 0xac, 0xb4, 0x05, 0xca, 0x7b, 0x46, 0xb8, 0x67, 0x74, 0xde, 0x5b, 0x64, 0x63, 0xbc, 0x23,
	0xc2, 0x9c, 0x20, 0xcc, 0xed, 0x31, 0x52, 0x0b, 0xf4, 0x88, 0x6c, 0xda, 0x54, 0x4c, 0xf1, 0x3e,
	0xec, 0x12, 0xe4, 0x66, 0x09, 0x13, 0x4c, 0xf6, 0xb9, 0x1d, 0xd0, 0x59, 0x8b, 0xb5, 0x9e, 0xc3,
	0x52, 0x3a, 0x4b, 0x44, 0x17, 0xa9, 0xa8, 0x7d, 0xb6, 0x3f, 0x94, 0x37, 0x5b, 0x0f, 0x9c, 0xd6,
	0x7b, 0x97, 0xbc, 0x58, 0xe4, 0x8d, 0x80, 0xbf, 0x21, 0xe0, 0xe6, 0x24, 0x17, 0x4b, 0xb9, 0x4b,
	0x96, 0x94, 0x8e, 0x32, 0xcd, 0x22, 0xcd, 0x32, 0xf9, 0x84, 0x09, 0x49, 0x9f, 0x2d, 0xdf, 0xaf,
	0x05, 0xab, 0xd9, 0xd1, 0xa1, 0x7c, 0xd2, 0x92, 0xde, 0x36, 0xb9, 0x99, 0xf3, 0x40, 0x80, 0xdf,
	0x11, 0x60, 0xf9, 0x6c, 0xb0, 0x4d, 0xfa, 0x06, 0xb9, 0x66, 0xe2, 0x3a, 0x72, 0x20, 0x34, 0x9d,
	0x2b, 0x4f, 0x57, 0x83, 0xe8, 0x3d, 0x13, 0xec, 0xbd, 0x4c, 0xae, 0x8f, 0x94, 0x98, 0xe4, 0x0f,
	0x4c, 0xb2, 0x30, 0x0c, 0xb1, 0x09, 0x0e, 0x08, 0x95, 0x03, 0x6d, 0x7e, 0xbb, 0xb6, 0xb9, 0x22,
	0xc7, 0xac, 0x73, 0x04, 0xab, 0xef, 0xb0, 0x24, 0x6e, 0xd0, 0x4a, 0x79, 0xbe, 0x5b, 0x28, 0xb6,
	0xd7, 0xeb, 0x78, 0xcf, 0x4a, 0x9b, 0x71, 0xc3, 0xdb, 0x21, 0xcf, 0x4d, 0x72, 0x45, 0x96, 0x3f,
	0x91, 0x65, 0x6d, 0xac, 0xbc, 0x04, 0xcc, 0xa7, 0xd5, 0x2b, 0x82, 0xf9, 0x05, 0x60, 0x3e, 0x82,
	0xfd, 0x55, 0x00, 0xe6, 0x97, 0x80, 0x05, 0xb4, 0x76, 0x45, 0xb0, 0xa0, 0x00, 0x2c, 0x40, 0xb0,
	0xbf, 0x0b, 0xc0, 0x02, 0x0b, 0xf6, 0x1e, 0x81, 0xcb, 0xd9, 0x49, 0x07, 0x31, 0x1f, 0x79, 0xc8,
	0x98, 0xd3, 0x6b, 0x63, 0x99, 0xda, 0x52, 0xa6, 0x3c, 0x12, 0xc8, 0x74, 0xc3, 0xe9, 0x9c, 0x2b,
	0xa8, 0xe0, 0xe0, 0xd1, 0x31, 0x66, 0x88, 0xf2, 0x0f, 0xa2, 0xdc, 0xba, 0xa0, 0xb2, 0x18, 0xfb,
	0xe4, 0xb6, 0x30, 0xa1, 0xb8, 0x00, 0x28, 0x78, 0x89, 0xd2, 0x09, 0x1e, 0x28, 0x32, 0xc5, 0xf6,
	0x18, 0xed, 0xde, 0x39, 0xa9, 0x39, 0x50, 0x6f, 0x93, 0xcd, 0x09, 0xa6, 0x88, 0xf4, 0xaf, 0xdb,
	0x9d, 0xb1, 0xea, 0x62, 0x2c, 0x9f, 0xce, 0x5f, 0x11, 0xcb, 0x9f, 0x8c, 0xe5, 0x4e, 0xd3, 0x7f,
	0x05, 0x58, 0x7e, 0x31, 0x56, 0x40, 0x17, 0xae, 0x88, 0x15, 0x4c, 0xc6, 0x72, 0x67, 0xe9, 0xff,
	0x02, 0x2c, 0x7b, 0x96, 0xb6, 0x3e, 0x22, 0xa4, 0x16, 0x72, 0xd5, 0x97, 0x42, 0x71, 0xef, 0x0e,
	0x99, 0xb3, 0xfd, 0xd0, 0xf5, 0xa8, 0x51, 0x21, 0x76, 0x4d, 0x16, 0x7b, 0xe5, 0x5d, 0xf3, 0x37,
	0xc4, 0x40, 0xef, 0x03, 0xb2, 0x6c, 0x3a, 0x21, 0x3b, 0xd3, 0x0a, 0xa1, 0xb9, 0xcc, 0x82, 0xb8,
	0x9e, 0x13, 0xe7, 0x1b, 0xe6, 0x7d, 0x18, 0x37, 0x4f, 0xc7, 0xe1, 0x52, 0xef, 0xfc, 0x03, 0x38,
	0x98, 0x55, 0xd7, 0x81, 0xa1, 0x2f, 0x18, 0xc7, 0xe7, 0x2f, 0x38, 0x62, 0x7f, 0xbe, 0x8f, 0xff,
	0xc3, 0x61, 0x38, 0x14, 0xf0, 0x59, 0x28, 0x7d, 0x50, 0xb4, 0x8d, 0xea, 0x4e, 0xbd, 0xfc, 0x4d,
	0xa1, 0x3e, 0xdc, 0x81, 0x3a, 0x14, 0xe4, 0xd0, 0x88, 0xbd, 0x0f, 0x61, 0x67, 0x07, 0xbd, 0x36,
	0xcf, 0x98, 0x3c, 0x64, 0x3c, 0xe5, 0x3d, 0x2e, 0xb4, 0x82, 0x65, 0x32, 0x78, 0x2f, 0x19, 0xa4,
	0x9a, 0xf2, 0xf2, 0xdf, 0x8c, 0xa2, 0xc1, 0x83, 0xc3, 0xbb, 0x4e, 0xde, 0x14, 0xa1, 0x15, 0xaf,
	0x7f, 0x5c, 0x25, 0xb3, 0x90, 0xc9, 0x5b, 0x25, 0x15, 0x53, 0xba, 0x93, 0x98, 0x7e, 0xd2, 0x02,
	0xbf, 0xb9, 0x70, 0x0e, 0x86, 0xcd, 0x18, 0x16, 0x70, 0xfd, 0xec, 0x5d, 0x84, 0xf9, 0x4f, 0x5b,
	0x53, 0x74, 0xa0, 0xf6, 0xe8, 0x7e, 0x82, 0xc7, 0xeb, 0xa4, 0x2a, 0x9c, 0xf8, 0xb3, 0x29, 0xc4,
	0x15, 0x81, 0x32, 0xa8, 0x2d, 0x9d, 0x23, 0xa9, 0xb8, 0x60, 0x3a, 0xe3, 0x50, 0x06, 0x9c, 0xc5,
	0xe7, 0x53, 0x58, 0x2c, 0xa3, 0xf0, 0x00, 0x74, 0x2d, 0x34, 0x0b, 0x48, 0xc5, 0x06, 0x06, 0xf4,
	0x8b, 0x56, 0x69, 0x6b, 0x77, 0xa1, 0x70, 0xa6, 0x17, 0xdd, 0xe2, 0xed, 0x83, 0x06, 0xfd, 0xb2,
	0x5c, 0xeb, 0x96, 0x6e, 0x07, 0x8d, 0x51, 0xda, 0x06, 0xfd, 0x6a, 0xda, 0xb4, 0x8d, 0x7c, 0xda,
	0x80, 0x7e, 0x7d, 0xb9, 0xb4, 0xc1, 0x28, 0xad, 0x4f, 0xbf, 0x99, 0x36, 0xad, 0x9f, 0x4f, 0xeb,
	0xd3, 0x6f, 0x2f, 0x97, 0xd6, 0x37, 0x87, 0x45, 0x1f, 0xc1, 0x49, 0x13, 0x51, 0x92, 0x62, 0x23,
	0xf8, 0xfe, 0x3b, 0xbc, 0xbf, 0x6b, 0xf9, 0x56, 0x70, 0xac, 0xb9, 0x42, 0x8b, 0xc5, 0x91, 0xc4,
	0x36, 0x81, 0x37, 0xc9, 0xfc, 0xd9, 0x4e, 0xf2, 0x43, 0xb9, 0x01, 0x39, 0x3d, 0x6c, 0x50, 0xda,
	0x56, 0xc7, 0x57, 0x21, 0xfa, 0xd3, 0xc9, 0x14, 0x6f, 0xbb, 0x37, 0xc7, 0x55, 0x27, 0xcf, 0x27,
	0x73, 0x36, 0x98, 0xfe, 0x7c, 0x32, 0x53, 0xba, 0x23, 0x18, 0x0a, 0x25, 0xa3, 0x86, 0xef, 0x80,
	0x90, 0xfa, 0x97, 0x69, 0x52, 0x57, 0xed, 0xc3, 0x66, 0xbc, 0x7b, 0x0f, 0x5e, 0x6c, 0x65, 0xae,
	0x52, 0x9c, 0x7e, 0x88, 0x3c, 0x7a, 0xf5, 0x52, 0x9f, 0x28, 0xed, 0x8a, 0xfd, 0x1a, 0x08, 0x9e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x3d, 0x74, 0x1d, 0xda, 0x0c, 0x00, 0x00,
}
