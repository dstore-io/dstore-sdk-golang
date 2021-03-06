// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/ac_InsertAction_Ad.proto
// DO NOT EDIT!

/*
Package ac_InsertAction_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/ac_InsertAction_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package ac_InsertAction_Ad

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
	CommandId                     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=command_id,json=commandId" json:"command_id,omitempty"`
	CommandIdNull                 bool                        `protobuf:"varint,1001,opt,name=command_id_null,json=commandIdNull" json:"command_id_null,omitempty"`
	Status                        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	StatusNull                    bool                        `protobuf:"varint,1002,opt,name=status_null,json=statusNull" json:"status_null,omitempty"`
	ParameterNames                *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=parameter_names,json=parameterNames" json:"parameter_names,omitempty"`
	ParameterNamesNull            bool                        `protobuf:"varint,1003,opt,name=parameter_names_null,json=parameterNamesNull" json:"parameter_names_null,omitempty"`
	ParameterValues               *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=parameter_values,json=parameterValues" json:"parameter_values,omitempty"`
	ParameterValuesNull           bool                        `protobuf:"varint,1004,opt,name=parameter_values_null,json=parameterValuesNull" json:"parameter_values_null,omitempty"`
	SeparatorInParamValues        *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=separator_in_param_values,json=separatorInParamValues" json:"separator_in_param_values,omitempty"`
	SeparatorInParamValuesNull    bool                        `protobuf:"varint,1005,opt,name=separator_in_param_values_null,json=separatorInParamValuesNull" json:"separator_in_param_values_null,omitempty"`
	SepInAnyValuesOtherThanLF     *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=sep_in_any_values_other_than_l_f,json=sepInAnyValuesOtherThanLF" json:"sep_in_any_values_other_than_l_f,omitempty"`
	SepInAnyValuesOtherThanLFNull bool                        `protobuf:"varint,1006,opt,name=sep_in_any_values_other_than_l_f_null,json=sepInAnyValuesOtherThanLFNull" json:"sep_in_any_values_other_than_l_f_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommandId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommandId
	}
	return nil
}

func (m *Parameters) GetStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Parameters) GetParameterNames() *dstore_values.StringValue {
	if m != nil {
		return m.ParameterNames
	}
	return nil
}

func (m *Parameters) GetParameterValues() *dstore_values.StringValue {
	if m != nil {
		return m.ParameterValues
	}
	return nil
}

func (m *Parameters) GetSeparatorInParamValues() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInParamValues
	}
	return nil
}

func (m *Parameters) GetSepInAnyValuesOtherThanLF() *dstore_values.StringValue {
	if m != nil {
		return m.SepInAnyValuesOtherThanLF
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	NewActionId     *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=new_action_id,json=newActionId" json:"new_action_id,omitempty"`
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

func (m *Response) GetNewActionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NewActionId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.ac_InsertAction_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.ac_InsertAction_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.ac_InsertAction_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/ac_InsertAction_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xb5, 0x95, 0x76, 0xe3, 0x54, 0xa3, 0x93, 0x07, 0x53, 0xd6, 0x6a, 0x53, 0x35, 0x84,
	0xe0, 0x2a, 0x65, 0x8c, 0x0b, 0xc4, 0x0d, 0x94, 0x69, 0x48, 0x91, 0x68, 0x41, 0xe1, 0x8f, 0x04,
	0x37, 0x96, 0x69, 0xbc, 0x2e, 0xa2, 0xb5, 0x2b, 0x3b, 0xa5, 0xda, 0x1b, 0x70, 0xc9, 0x6b, 0xf2,
	0xf7, 0x86, 0x17, 0xe0, 0x38, 0x27, 0x69, 0x69, 0xd6, 0x29, 0xdc, 0x2c, 0xb3, 0xcf, 0xf9, 0x7e,
	0xdf, 0xe7, 0xca, 0xc7, 0x70, 0x14, 0xd9, 0x44, 0x1b, 0xd9, 0x91, 0x6a, 0x18, 0x2b, 0xd9, 0x99,
	0x18, 0x3d, 0x90, 0xd1, 0xd4, 0x48, 0xdb, 0x11, 0x03, 0x1e, 0x28, 0x2b, 0x4d, 0xd2, 0x1d, 0x24,
	0xb1, 0x56, 0xbc, 0x1b, 0xf9, 0x58, 0x4d, 0x34, 0x6b, 0x93, 0xc4, 0x27, 0x89, 0x7f, 0xb9, 0xaf,
	0xb9, 0x93, 0x41, 0x3f, 0x8b, 0xd1, 0x54, 0x5a, 0x92, 0x35, 0xf7, 0x96, 0x9d, 0xa4, 0x31, 0xda,
	0x64, 0xa5, 0xd6, 0x72, 0x69, 0x2c, 0xad, 0x15, 0x43, 0x99, 0x15, 0x6f, 0x17, 0x8b, 0x89, 0x88,
	0xd5, 0x99, 0x36, 0x63, 0xe1, 0xec, 0xa8, 0xe9, 0xf0, 0x4b, 0x0d, 0xe0, 0x95, 0x30, 0x02, 0xab,
	0xd2, 0x58, 0xf6, 0x18, 0x60, 0xa0, 0xc7, 0x63, 0xa1, 0x22, 0x1e, 0x47, 0xde, 0x5a, 0x7b, 0xed,
	0x5e, 0xfd, 0x41, 0xcb, 0xcf, 0x72, 0x67, 0xa9, 0x62, 0x95, 0xc8, 0xa1, 0x34, 0xef, 0xdc, 0x2a,
	0xbc, 0x9e, 0xb5, 0x07, 0x11, 0xbb, 0x0b, 0x8d, 0x85, 0x96, 0xab, 0xe9, 0x68, 0xe4, 0x7d, 0xdb,
	0x40, 0xc2, 0x66, 0xb8, 0x35, 0x6f, 0xea, 0xe3, 0x2e, 0x3b, 0x86, 0x9a, 0x4d, 0x44, 0x32, 0xb5,
	0xde, 0x7a, 0xb9, 0x41, 0xd6, 0xca, 0xda, 0x50, 0xa7, 0xff, 0x88, 0xfc, 0x9d, 0xc8, 0x40, 0x7b,
	0x29, 0xf6, 0x04, 0x1a, 0x93, 0xfc, 0x24, 0x5c, 0xe1, 0xd7, 0x7a, 0x95, 0x94, 0xdf, 0x2c, 0xf0,
	0x6d, 0x62, 0x62, 0x35, 0x24, 0xfc, 0x8d, 0xb9, 0xa4, 0xef, 0x14, 0xec, 0x08, 0x6e, 0x16, 0x20,
	0xe4, 0xf7, 0x83, 0xfc, 0xd8, 0x72, 0x7b, 0xea, 0x7b, 0x0a, 0xdb, 0x0b, 0x09, 0x59, 0x78, 0xd7,
	0x4a, 0x8d, 0x17, 0x59, 0xd3, 0xb5, 0xc5, 0x5f, 0xe5, 0x56, 0x11, 0x43, 0xd6, 0x3f, 0xc9, 0x7a,
	0xa7, 0x20, 0x48, 0xbd, 0xdf, 0xc2, 0x9e, 0x95, 0xae, 0x80, 0x36, 0x3c, 0x56, 0x3c, 0xed, 0xc9,
	0x43, 0x54, 0x4b, 0x43, 0xec, 0xce, 0xc5, 0x81, 0x4a, 0x6f, 0x41, 0x96, 0xe5, 0x04, 0x0e, 0xae,
	0xc4, 0x52, 0xa8, 0x5f, 0x14, 0xaa, 0xb9, 0x1a, 0x90, 0x66, 0xe3, 0xd0, 0xc6, 0xaa, 0x93, 0x0b,
	0x75, 0x91, 0x8b, 0x75, 0x72, 0x8e, 0xc7, 0x4b, 0xce, 0x85, 0xe2, 0x23, 0x7e, 0xe6, 0xd5, 0x4a,
	0x23, 0xba, 0xf3, 0x05, 0xaa, 0xab, 0x2e, 0x88, 0xfc, 0xd2, 0xe9, 0xdf, 0xa0, 0xfc, 0xc5, 0x73,
	0xd6, 0x83, 0x3b, 0x65, 0x06, 0x14, 0xf6, 0x37, 0x85, 0xdd, 0xbf, 0x12, 0xe5, 0xf2, 0x1e, 0xfe,
	0x59, 0x87, 0xcd, 0x50, 0xda, 0x89, 0xc6, 0xa1, 0x64, 0xf7, 0xa1, 0x9a, 0x0e, 0x5a, 0x36, 0x03,
	0xf3, 0x84, 0xd9, 0xec, 0xd2, 0x10, 0x9e, 0xba, 0xbf, 0x21, 0x35, 0xb2, 0xf7, 0xb0, 0xed, 0x46,
	0x8c, 0xff, 0x33, 0x63, 0x78, 0xbf, 0x2b, 0x28, 0xf6, 0x0b, 0xe2, 0xe2, 0x24, 0xf6, 0x70, 0x1d,
	0x2c, 0xd6, 0x61, 0x63, 0xbc, 0xbc, 0xc1, 0x1e, 0xc1, 0x46, 0x36, 0xda, 0x78, 0xa3, 0x1d, 0xf1,
	0xe0, 0x12, 0x91, 0x06, 0xbf, 0x47, 0xdf, 0x30, 0x6f, 0x67, 0x4f, 0xa1, 0x62, 0xf4, 0x0c, 0xaf,
	0xe3, 0xaa, 0x1c, 0x2b, 0x1e, 0xaa, 0xfc, 0xfc, 0x7e, 0xa8, 0x67, 0xa1, 0x93, 0xb2, 0x27, 0xb0,
	0xa5, 0xe4, 0x8c, 0x0b, 0x6a, 0xc1, 0x47, 0x41, 0x96, 0xcf, 0x6c, 0x1d, 0x15, 0xc4, 0x0c, 0xa2,
	0xe6, 0x3e, 0x54, 0x10, 0xc6, 0x76, 0xa1, 0x86, 0x38, 0x07, 0xf8, 0xda, 0x47, 0x42, 0x35, 0xac,
	0xe2, 0x32, 0x88, 0x9e, 0xbd, 0x86, 0x56, 0xac, 0x0b, 0xc1, 0x16, 0x8f, 0xe9, 0x87, 0x87, 0x43,
	0x6d, 0xa3, 0x4f, 0x79, 0x3d, 0xfa, 0xbf, 0xf7, 0xf6, 0x63, 0x2d, 0x7d, 0xdc, 0x8e, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x59, 0x66, 0x22, 0xa5, 0x05, 0x00, 0x00,
}
