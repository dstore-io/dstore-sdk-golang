// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/ac_GetCommandSettingEntry.proto
// DO NOT EDIT!

/*
Package ac_GetCommandSettingEntry is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/ac_GetCommandSettingEntry.proto

It has these top-level messages:
	Parameters
	Response
*/
package ac_GetCommandSettingEntry

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
	CommandId        *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=command_id,json=commandId" json:"command_id,omitempty"`
	CommandIdNull    bool                        `protobuf:"varint,1001,opt,name=command_id_null,json=commandIdNull" json:"command_id_null,omitempty"`
	KeyVariable      *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull  bool                        `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	SelectResult     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=select_result,json=selectResult" json:"select_result,omitempty"`
	SelectResultNull bool                        `protobuf:"varint,1003,opt,name=select_result_null,json=selectResultNull" json:"select_result_null,omitempty"`
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

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Parameters) GetSelectResult() *dstore_values.BooleanValue {
	if m != nil {
		return m.SelectResult
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Value           *dstore_values.StringValue                       `protobuf:"bytes,101,opt,name=value" json:"value,omitempty"`
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

func (m *Response) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type Response_Row struct {
	RowId int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=value" json:"value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.ac_GetCommandSettingEntry.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.ac_GetCommandSettingEntry.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.ac_GetCommandSettingEntry.Response.Row")
}

var fileDescriptor0 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0xa5, 0x5d, 0xd3, 0xd6, 0xdb, 0x96, 0xd4, 0x11, 0x64, 0x4d, 0x40, 0xa4, 0x3e, 0x54, 0x10,
	0x37, 0x5a, 0x11, 0x8a, 0x20, 0x88, 0x52, 0x4a, 0x1f, 0x1a, 0xea, 0x08, 0x05, 0x7d, 0x59, 0x26,
	0xd9, 0x6b, 0x58, 0xdc, 0x9d, 0x29, 0x33, 0x93, 0x96, 0xfe, 0x85, 0xfe, 0x80, 0x9f, 0xe4, 0x7f,
	0xa8, 0x3f, 0xe1, 0xdd, 0xbd, 0x93, 0x26, 0xbb, 0x2a, 0xc6, 0x97, 0x84, 0xbb, 0xf7, 0x9c, 0x7b,
	0x66, 0xce, 0xb9, 0x03, 0x07, 0x99, 0xf3, 0xc6, 0xe2, 0x00, 0xf5, 0x24, 0xd7, 0x38, 0x38, 0xb7,
	0x66, 0x8c, 0xd9, 0xd4, 0xa2, 0x1b, 0xa8, 0x71, 0x7a, 0x84, 0xfe, 0x8d, 0x29, 0x4b, 0xa5, 0xb3,
	0x77, 0xe8, 0x7d, 0xae, 0x27, 0x87, 0xda, 0xdb, 0xab, 0x84, 0x40, 0xde, 0x88, 0x3d, 0x66, 0x26,
	0xcc, 0x4c, 0xfe, 0x0a, 0xef, 0xdd, 0x0e, 0x12, 0x17, 0xaa, 0x98, 0xa2, 0x63, 0x76, 0xef, 0x6e,
	0x53, 0x17, 0xad, 0x35, 0x36, 0xb4, 0xfa, 0xcd, 0x56, 0x89, 0xce, 0xa9, 0x09, 0x86, 0xe6, 0x83,
	0x76, 0xd3, 0xab, 0x5c, 0x7f, 0x34, 0xb6, 0x54, 0x3e, 0x37, 0x9a, 0x41, 0xbb, 0xdf, 0x56, 0x01,
	0x4e, 0x95, 0x55, 0xd4, 0x45, 0xeb, 0xc4, 0x0b, 0x80, 0x31, 0x9f, 0x2b, 0xcd, 0xb3, 0x78, 0xe5,
	0xfe, 0xca, 0xc3, 0xcd, 0xfd, 0x7e, 0x12, 0x8e, 0x1f, 0x4e, 0x95, 0x6b, 0x8f, 0x13, 0xb4, 0x67,
	0x55, 0x25, 0x6f, 0x06, 0xf8, 0x71, 0x26, 0xf6, 0xa0, 0x3b, 0xe7, 0xa6, 0x7a, 0x5a, 0x14, 0xf1,
	0xf7, 0x75, 0x9a, 0xb0, 0x21, 0xb7, 0xaf, 0x41, 0x43, 0xfa, 0x2a, 0x5e, 0xc2, 0xd6, 0x27, 0xbc,
	0x4a, 0x2f, 0x94, 0xcd, 0xd5, 0xa8, 0xc0, 0x78, 0xb5, 0x96, 0xe9, 0xb5, 0x64, 0x9c, 0xb7, 0xe4,
	0x0b, 0xab, 0x6c, 0x12, 0xfe, 0x2c, 0xc0, 0xc5, 0x23, 0xb8, 0xb5, 0x48, 0x67, 0xa5, 0x1f, 0xac,
	0xd4, 0x5d, 0x00, 0xd6, 0x5a, 0xaf, 0x60, 0xdb, 0x61, 0x81, 0x63, 0x9f, 0x52, 0x50, 0xd3, 0xc2,
	0xc7, 0xd1, 0x1f, 0xef, 0x34, 0x32, 0xa6, 0x40, 0xa5, 0x59, 0x6d, 0x8b, 0x19, 0xb2, 0x26, 0x88,
	0xc7, 0x20, 0x1a, 0x13, 0x58, 0xef, 0x27, 0xeb, 0xed, 0x2c, 0x42, 0x2b, 0xc1, 0xdd, 0xaf, 0x11,
	0x6c, 0x50, 0x79, 0x6e, 0xb4, 0x43, 0xf1, 0x04, 0x3a, 0x75, 0x5c, 0xc1, 0xc9, 0xeb, 0x2b, 0x86,
	0x45, 0xe0, 0x28, 0x0f, 0xab, 0x5f, 0xc9, 0x40, 0xf1, 0x1e, 0x76, 0xaa, 0xa0, 0xd2, 0x85, 0xa4,
	0xc8, 0x9f, 0x88, 0xc8, 0x49, 0x8b, 0xdc, 0xce, 0xf3, 0x84, 0xea, 0xe3, 0x79, 0x2d, 0xbb, 0x65,
	0xf3, 0x83, 0x38, 0x80, 0xf5, 0xb0, 0x20, 0x64, 0x42, 0x35, 0xf1, 0xde, 0x6f, 0x13, 0x79, 0x7d,
	0x4e, 0xf8, 0x5f, 0xce, 0xe0, 0xe2, 0x08, 0x22, 0x6b, 0x2e, 0xe3, 0x1b, 0x35, 0xeb, 0x79, 0xb2,
	0xe4, 0x36, 0x27, 0x33, 0x1b, 0x12, 0x69, 0x2e, 0x65, 0x35, 0xa1, 0xf2, 0xa3, 0x36, 0x3c, 0xc6,
	0x7f, 0x46, 0xce, 0xc0, 0xde, 0x29, 0x44, 0xc4, 0x16, 0x77, 0x60, 0x8d, 0xf8, 0xd5, 0x4e, 0x7e,
	0x1e, 0x12, 0xb5, 0x23, 0x3b, 0x54, 0xd2, 0xce, 0x3d, 0x9d, 0x0d, 0xfc, 0x32, 0x5c, 0x72, 0xe2,
	0xeb, 0xb7, 0xd0, 0xcf, 0x4d, 0xeb, 0x0e, 0xf3, 0xb7, 0xfc, 0x61, 0xff, 0xff, 0x5f, 0xf9, 0x68,
	0xad, 0x7e, 0x4b, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x65, 0x35, 0x85, 0xb3, 0x22, 0x04,
	0x00, 0x00,
}
