// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetTRITriggerReplFuncts_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetTRITriggerReplFuncts_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetTRITriggerReplFuncts_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetTRITriggerReplFuncts_Ad

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
	TriggerId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=trigger_id,json=triggerId" json:"trigger_id,omitempty"`
	TriggerIdNull              bool                        `protobuf:"varint,1001,opt,name=trigger_id_null,json=triggerIdNull" json:"trigger_id_null,omitempty"`
	FunctionDefinitionLIKE     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=function_definition_l_i_k_e,json=functionDefinitionLIKE" json:"function_definition_l_i_k_e,omitempty"`
	FunctionDefinitionLIKENull bool                        `protobuf:"varint,1002,opt,name=function_definition_l_i_k_e_null,json=functionDefinitionLIKENull" json:"function_definition_l_i_k_e_null,omitempty"`
	ReplacementFunction        *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=replacement_function,json=replacementFunction" json:"replacement_function,omitempty"`
	ReplacementFunctionNull    bool                        `protobuf:"varint,1003,opt,name=replacement_function_null,json=replacementFunctionNull" json:"replacement_function_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTriggerId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TriggerId
	}
	return nil
}

func (m *Parameters) GetFunctionDefinitionLIKE() *dstore_values.StringValue {
	if m != nil {
		return m.FunctionDefinitionLIKE
	}
	return nil
}

func (m *Parameters) GetReplacementFunction() *dstore_values.StringValue {
	if m != nil {
		return m.ReplacementFunction
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
	RowId              int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description        *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	FunctionDefinition *dstore_values.StringValue `protobuf:"bytes,10002,opt,name=function_definition,json=functionDefinition" json:"function_definition,omitempty"`
	ReplaceString      *dstore_values.StringValue `protobuf:"bytes,10003,opt,name=replace_string,json=replaceString" json:"replace_string,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetFunctionDefinition() *dstore_values.StringValue {
	if m != nil {
		return m.FunctionDefinition
	}
	return nil
}

func (m *Response_Row) GetReplaceString() *dstore_values.StringValue {
	if m != nil {
		return m.ReplaceString
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetTRITriggerReplFuncts_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetTRITriggerReplFuncts_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetTRITriggerReplFuncts_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x25, 0x6e, 0xd3, 0xd6, 0x5b, 0x6a, 0x65, 0x22, 0x75, 0xbb, 0x01, 0x29, 0xf5, 0x41, 0x41,
	0xd8, 0x88, 0x82, 0xf8, 0xf9, 0x60, 0x31, 0x95, 0xd0, 0x26, 0xc8, 0x58, 0x15, 0x7d, 0x19, 0xd6,
	0xec, 0xcd, 0x32, 0xb8, 0x3b, 0x13, 0x66, 0x36, 0xf6, 0x67, 0xf8, 0xf5, 0xec, 0x0f, 0x54, 0xc1,
	0xdf, 0xe0, 0xec, 0xce, 0xe4, 0x6b, 0x1b, 0x12, 0xec, 0x4b, 0x92, 0x3b, 0xf7, 0x9c, 0x73, 0xcf,
	0xdc, 0x7b, 0x33, 0xf0, 0x34, 0xd6, 0xb9, 0x54, 0xd8, 0x42, 0x91, 0x70, 0x81, 0xad, 0xa1, 0x92,
	0x7d, 0x8c, 0x47, 0x0a, 0x75, 0x2b, 0xe3, 0xec, 0x25, 0xe6, 0xa7, 0xb4, 0x73, 0xaa, 0x78, 0x92,
	0xa0, 0xa2, 0x38, 0x4c, 0x8f, 0x46, 0xa2, 0x9f, 0x6b, 0xf6, 0x3c, 0x0e, 0x0d, 0x30, 0x97, 0xe4,
	0x8e, 0x65, 0x87, 0x96, 0x1d, 0x2e, 0xa5, 0x04, 0x0d, 0x57, 0xea, 0x73, 0x94, 0x8e, 0x50, 0x5b,
	0x85, 0x60, 0x6f, 0xbe, 0x3e, 0x2a, 0x25, 0x95, 0x4b, 0x35, 0xe7, 0x53, 0x19, 0x6a, 0x1d, 0x25,
	0xe8, 0x92, 0x37, 0xab, 0xc9, 0x3c, 0xe2, 0x62, 0x20, 0x55, 0x16, 0xe5, 0x5c, 0x0a, 0x0b, 0x3a,
	0xf8, 0xe9, 0x01, 0xbc, 0x8a, 0x54, 0x64, 0xb2, 0xa8, 0x34, 0x79, 0x0c, 0x90, 0x5b, 0x63, 0x8c,
	0xc7, 0x7e, 0x6d, 0xbf, 0x76, 0x7b, 0xeb, 0x5e, 0x33, 0x74, 0x57, 0x70, 0xae, 0xb8, 0xc8, 0xd1,
	0x00, 0xde, 0x16, 0x11, 0xbd, 0xec, 0xe0, 0x9d, 0x98, 0xdc, 0x82, 0x9d, 0x29, 0x97, 0x89, 0x51,
	0x9a, 0xfa, 0xbf, 0x36, 0x8c, 0xc2, 0x26, 0xdd, 0x9e, 0x80, 0x7a, 0xe6, 0x94, 0xbc, 0x83, 0xe6,
	0xa0, 0xb8, 0xb2, 0x71, 0xc1, 0x62, 0x1c, 0x70, 0xc1, 0xcb, 0x9f, 0x29, 0xe3, 0xec, 0x13, 0x43,
	0xff, 0x52, 0x59, 0x35, 0xa8, 0x54, 0xd5, 0x46, 0x43, 0x24, 0xb6, 0xe8, 0xee, 0x98, 0xfe, 0x62,
	0xc2, 0x3e, 0xe9, 0x1c, 0xb7, 0x49, 0x1b, 0xf6, 0x97, 0x08, 0x5b, 0x4b, 0xbf, 0xad, 0xa5, 0x60,
	0xb1, 0x44, 0xe9, 0xaf, 0x0b, 0xd7, 0x94, 0x19, 0x4b, 0xd4, 0xc7, 0x0c, 0x45, 0xce, 0xc6, 0x48,
	0xdf, 0x5b, 0x69, 0xac, 0x31, 0xc3, 0x3b, 0x72, 0x34, 0xf2, 0x04, 0xf6, 0x16, 0xc9, 0x59, 0x3b,
	0x7f, 0xac, 0x9d, 0xeb, 0x0b, 0x88, 0x85, 0x97, 0x83, 0x2f, 0x6b, 0xb0, 0x49, 0x51, 0x0f, 0xa5,
	0xd0, 0x48, 0xee, 0x42, 0xbd, 0x9c, 0xbe, 0x1b, 0xcc, 0xc4, 0x89, 0xdb, 0x2d, 0xbb, 0x19, 0xed,
	0xe2, 0x93, 0x5a, 0x20, 0x79, 0x0f, 0x57, 0x8b, 0xb9, 0xb3, 0x99, 0xc1, 0x9b, 0xfe, 0x7a, 0x86,
	0x1c, 0x56, 0xc8, 0xd5, 0xf5, 0xe8, 0x9a, 0xb8, 0x33, 0x8d, 0xe9, 0x4e, 0x36, 0x7f, 0x40, 0x1e,
	0xc2, 0x86, 0xdb, 0x37, 0xd3, 0x98, 0x42, 0xf1, 0xc6, 0x39, 0x45, 0xbb, 0x8d, 0x5d, 0xfb, 0x4d,
	0xc7, 0x70, 0x72, 0x0c, 0x9e, 0x92, 0x67, 0xfe, 0x5a, 0xc9, 0x7a, 0x14, 0xfe, 0xc7, 0x1f, 0x24,
	0x1c, 0xb7, 0x22, 0xa4, 0xf2, 0x8c, 0x16, 0x2a, 0xc1, 0xdf, 0x1a, 0x78, 0x26, 0x20, 0xbb, 0xb0,
	0x6e, 0xc2, 0x62, 0x6b, 0xbf, 0xf6, 0x4c, 0x77, 0xea, 0xb4, 0x6e, 0x42, 0xb3, 0x95, 0xcf, 0x60,
	0x2b, 0x46, 0xdd, 0x57, 0x7c, 0x58, 0x5e, 0xfe, 0x5b, 0x6f, 0xe5, 0x10, 0x67, 0xf1, 0xe4, 0x04,
	0x1a, 0x0b, 0x56, 0xca, 0xff, 0xbe, 0x5a, 0x86, 0x9c, 0xdf, 0x30, 0x72, 0x08, 0x57, 0xdc, 0xa0,
	0x99, 0x85, 0xfa, 0x3f, 0x56, 0x0b, 0x6d, 0x3b, 0xca, 0xeb, 0xf2, 0xec, 0xf0, 0x0d, 0x34, 0xb9,
	0xac, 0x34, 0x6d, 0xfa, 0x26, 0x7d, 0x78, 0x70, 0xb1, 0xd7, 0xea, 0xe3, 0x7a, 0xf9, 0x1e, 0xdc,
	0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x62, 0x5e, 0x1a, 0xee, 0x04, 0x00, 0x00,
}
