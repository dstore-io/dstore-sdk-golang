// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetSortOrderForResultSet_Ad.proto
// DO NOT EDIT!

/*
Package do_GetSortOrderForResultSet_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetSortOrderForResultSet_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetSortOrderForResultSet_Ad

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
	ProcedureName                  *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	ProcedureNameNull              bool                        `protobuf:"varint,1001,opt,name=procedure_name_null,json=procedureNameNull" json:"procedure_name_null,omitempty"`
	ProcedureResultConditionId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=procedure_result_condition_id,json=procedureResultConditionId" json:"procedure_result_condition_id,omitempty"`
	ProcedureResultConditionIdNull bool                        `protobuf:"varint,1002,opt,name=procedure_result_condition_id_null,json=procedureResultConditionIdNull" json:"procedure_result_condition_id_null,omitempty"`
	SortConditionId                *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=sort_condition_id,json=sortConditionId" json:"sort_condition_id,omitempty"`
	SortConditionIdNull            bool                        `protobuf:"varint,1003,opt,name=sort_condition_id_null,json=sortConditionIdNull" json:"sort_condition_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func (m *Parameters) GetProcedureResultConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureResultConditionId
	}
	return nil
}

func (m *Parameters) GetSortConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortConditionId
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
	RowId                    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Condition                *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=condition" json:"condition,omitempty"`
	SortConditionId          *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=sort_condition_id,json=sortConditionId" json:"sort_condition_id,omitempty"`
	ParameterName            *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=parameter_name,json=parameterName" json:"parameter_name,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	AscendingSortOrder       *dstore_values.BooleanValue `protobuf:"bytes,10005,opt,name=ascending_sort_order,json=ascendingSortOrder" json:"ascending_sort_order,omitempty"`
	ResultColumnName         *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=result_column_name,json=resultColumnName" json:"result_column_name,omitempty"`
	SortConditionDescription *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=sort_condition_description,json=sortConditionDescription" json:"sort_condition_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCondition() *dstore_values.StringValue {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *Response_Row) GetSortConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortConditionId
	}
	return nil
}

func (m *Response_Row) GetParameterName() *dstore_values.StringValue {
	if m != nil {
		return m.ParameterName
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetAscendingSortOrder() *dstore_values.BooleanValue {
	if m != nil {
		return m.AscendingSortOrder
	}
	return nil
}

func (m *Response_Row) GetResultColumnName() *dstore_values.StringValue {
	if m != nil {
		return m.ResultColumnName
	}
	return nil
}

func (m *Response_Row) GetSortConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.SortConditionDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetSortOrderForResultSet_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetSortOrderForResultSet_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetSortOrderForResultSet_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/do_GetSortOrderForResultSet_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x49, 0x9b, 0x94, 0x41, 0xd0, 0x76, 0x83, 0x2a, 0xe3, 0x88, 0x0a, 0x85, 0x17, 0x1e,
	0x90, 0x83, 0xb8, 0x48, 0x80, 0x84, 0x44, 0xcb, 0xa5, 0x44, 0x50, 0x83, 0x5c, 0x09, 0xa9, 0x3c,
	0x60, 0xb9, 0xd9, 0xc5, 0xb2, 0x70, 0x76, 0xa2, 0x5d, 0x87, 0xfe, 0x06, 0x50, 0x2e, 0x3f, 0xc1,
	0x2b, 0x1f, 0x05, 0xfc, 0x04, 0x7b, 0xf1, 0xa5, 0x76, 0x4b, 0xd3, 0xf2, 0x92, 0x68, 0x3c, 0x73,
	0x8e, 0xcf, 0x9c, 0x99, 0x31, 0x3c, 0xa0, 0x32, 0x43, 0xc1, 0x06, 0x8c, 0xc7, 0x09, 0x67, 0x83,
	0x89, 0xc0, 0x11, 0xa3, 0x53, 0xc1, 0xe4, 0x80, 0x62, 0xb8, 0xc9, 0xb2, 0x6d, 0x14, 0xd9, 0x4b,
	0x41, 0x99, 0x78, 0x8a, 0x22, 0x60, 0x72, 0x9a, 0x66, 0xdb, 0x2c, 0x0b, 0xd7, 0xa9, 0xa7, 0x2a,
	0x33, 0x24, 0xd7, 0x2d, 0xdc, 0xb3, 0x70, 0xef, 0x78, 0x8c, 0xdb, 0xcd, 0x5f, 0xf6, 0x21, 0x4a,
	0xa7, 0x4c, 0x5a, 0x0a, 0xf7, 0x52, 0x5d, 0x01, 0x13, 0x02, 0x45, 0x9e, 0xea, 0xd5, 0x53, 0x63,
	0x26, 0x65, 0x14, 0xb3, 0x3c, 0x79, 0xb5, 0x99, 0xcc, 0xa2, 0x84, 0xbf, 0x43, 0x31, 0x8e, 0xb2,
	0x04, 0xb9, 0x2d, 0xea, 0xff, 0x68, 0x01, 0xbc, 0x8a, 0x44, 0xa4, 0xb2, 0x4c, 0x48, 0xb2, 0x0e,
	0x17, 0xca, 0x0e, 0x43, 0xae, 0x9e, 0x3a, 0x73, 0x57, 0xe6, 0xae, 0x9d, 0xbb, 0xe9, 0x7a, 0x79,
	0x1f, 0xb9, 0x32, 0x99, 0x89, 0x84, 0xc7, 0xaf, 0x75, 0x10, 0x9c, 0x2f, 0x11, 0xbe, 0x02, 0x90,
	0x01, 0x74, 0xeb, 0x14, 0x21, 0x9f, 0xa6, 0xa9, 0xf3, 0xab, 0xa3, 0x88, 0x16, 0x83, 0x95, 0x5a,
	0xb1, 0xaf, 0x32, 0xe4, 0x2d, 0x5c, 0xae, 0x00, 0xc2, 0xd8, 0x11, 0x8e, 0x90, 0xd3, 0x44, 0xcb,
	0x0c, 0x13, 0xea, 0x9c, 0x31, 0x12, 0x7a, 0x0d, 0x09, 0x09, 0xcf, 0x58, 0xcc, 0x84, 0xd5, 0xe0,
	0x96, 0x0c, 0xd6, 0xcf, 0x47, 0x05, 0x7e, 0x48, 0xc9, 0x73, 0xe8, 0x1f, 0xcb, 0x6f, 0xf5, 0xfd,
	0xb6, 0xfa, 0xd6, 0xfe, 0x4d, 0x64, 0xc4, 0x6e, 0xc2, 0x8a, 0x54, 0xd3, 0xab, 0x0b, 0x6c, 0xcd,
	0x16, 0xb8, 0xa4, 0x51, 0x07, 0x55, 0xdd, 0x86, 0xd5, 0x43, 0x44, 0x56, 0xc9, 0x1f, 0xab, 0xa4,
	0xdb, 0x40, 0xe8, 0xd7, 0xf7, 0x7f, 0xb6, 0x61, 0x51, 0x09, 0x9b, 0x20, 0x97, 0x8c, 0xdc, 0x80,
	0x05, 0xb3, 0x0c, 0xcd, 0x19, 0xe5, 0xbb, 0x66, 0x17, 0xe5, 0x89, 0xfe, 0x0d, 0x6c, 0x21, 0xd9,
	0x81, 0x65, 0xbd, 0x06, 0xe1, 0x81, 0x3d, 0x50, 0xee, 0xb6, 0x14, 0xd8, 0x6b, 0x80, 0x9b, 0xdb,
	0xb2, 0xa5, 0xe2, 0x61, 0x15, 0x07, 0x4b, 0xe3, 0xfa, 0x03, 0x72, 0x17, 0x3a, 0xf9, 0xfa, 0x29,
	0x3b, 0x34, 0xe3, 0xda, 0x21, 0x46, 0xbb, 0x9c, 0x5b, 0xf6, 0x3f, 0x28, 0xca, 0xc9, 0x0b, 0x68,
	0x09, 0xdc, 0x73, 0xe6, 0x0d, 0xea, 0xbe, 0x77, 0x9a, 0x83, 0xf1, 0x0a, 0x2f, 0xbc, 0x00, 0xf7,
	0x02, 0x4d, 0xe3, 0xee, 0xcf, 0x43, 0x4b, 0x05, 0x64, 0x15, 0xda, 0x2a, 0xd4, 0xd3, 0xf9, 0xe8,
	0x2b, 0x7b, 0x16, 0x82, 0x05, 0x15, 0x2a, 0xdf, 0xef, 0xc1, 0xd9, 0xd2, 0x72, 0xe7, 0x93, 0x3f,
	0x73, 0xbb, 0xab, 0x6a, 0xf2, 0xec, 0xa8, 0xd9, 0x7f, 0xf6, 0xff, 0x63, 0xf8, 0x1b, 0xea, 0xcc,
	0x8a, 0xa3, 0xb3, 0x67, 0xb6, 0xef, 0x9f, 0xe0, 0xce, 0x0a, 0x88, 0xb9, 0xb3, 0x3b, 0xd0, 0x31,
	0x6a, 0x38, 0x3a, 0x5f, 0x4e, 0xa0, 0xa1, 0xad, 0x8b, 0x7d, 0x24, 0x3e, 0x5c, 0x8c, 0xe4, 0x88,
	0x29, 0x2d, 0x3c, 0x0e, 0x0d, 0x01, 0x6a, 0x63, 0x9d, 0xaf, 0x47, 0x73, 0xec, 0x22, 0xa6, 0x2c,
	0xe2, 0x96, 0x83, 0x94, 0xc8, 0x72, 0x20, 0x64, 0x08, 0xa4, 0xbc, 0xa9, 0x74, 0x3a, 0xe6, 0xb6,
	0x9d, 0x6f, 0xb3, 0xdb, 0x59, 0x16, 0xf9, 0x81, 0x69, 0x94, 0xe9, 0x68, 0x07, 0xdc, 0x86, 0xbf,
	0x94, 0xc9, 0x91, 0x48, 0x26, 0x66, 0x56, 0xdf, 0x67, 0x53, 0x3a, 0x35, 0x9f, 0x1f, 0x57, 0xe0,
	0x8d, 0x10, 0x7a, 0x09, 0x36, 0x56, 0xab, 0xfa, 0x94, 0xbf, 0x79, 0x18, 0xa3, 0xa4, 0xef, 0x8b,
	0x3c, 0x3d, 0xfd, 0xd7, 0x7e, 0xb7, 0x6d, 0x3e, 0xa7, 0xb7, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x85, 0x16, 0x04, 0x2f, 0x06, 0x00, 0x00,
}
