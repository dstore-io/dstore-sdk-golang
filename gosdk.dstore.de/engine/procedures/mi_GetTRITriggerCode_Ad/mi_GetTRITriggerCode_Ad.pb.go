// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetTRITriggerCode_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetTRITriggerCode_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetTRITriggerCode_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetTRITriggerCode_Ad

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
	TriggerId               *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=trigger_id,json=triggerId" json:"trigger_id,omitempty"`
	TriggerIdNull           bool                        `protobuf:"varint,1001,opt,name=trigger_id_null,json=triggerIdNull" json:"trigger_id_null,omitempty"`
	MaxCharsPerCodeLine     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=max_chars_per_code_line,json=maxCharsPerCodeLine" json:"max_chars_per_code_line,omitempty"`
	MaxCharsPerCodeLineNull bool                        `protobuf:"varint,1002,opt,name=max_chars_per_code_line_null,json=maxCharsPerCodeLineNull" json:"max_chars_per_code_line_null,omitempty"`
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

func (m *Parameters) GetMaxCharsPerCodeLine() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxCharsPerCodeLine
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetTRITriggerCode_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetTRITriggerCode_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetTRITriggerCode_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetTRITriggerCode_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xe1, 0x6a, 0x14, 0x31,
	0x10, 0xc7, 0xe9, 0xad, 0xd7, 0xd6, 0x11, 0xa9, 0xa4, 0x60, 0xd7, 0x3d, 0x15, 0xa9, 0x88, 0x7e,
	0xca, 0x89, 0x05, 0x51, 0x3f, 0x08, 0x5a, 0x8a, 0x2c, 0xd8, 0x52, 0x43, 0x11, 0xf5, 0xcb, 0x12,
	0x2f, 0xe3, 0x1a, 0xdc, 0xdd, 0x1c, 0xc9, 0x9e, 0xf5, 0x01, 0x7c, 0x00, 0xdf, 0xc6, 0x67, 0xd2,
	0xa7, 0x70, 0xb2, 0xc9, 0x79, 0xee, 0x6a, 0x51, 0xbf, 0xdc, 0xed, 0x64, 0xfe, 0xbf, 0xc9, 0x3f,
	0x33, 0x09, 0xdc, 0x57, 0xae, 0x35, 0x16, 0xa7, 0xd8, 0x94, 0xba, 0xc1, 0xe9, 0xdc, 0x9a, 0x19,
	0xaa, 0x85, 0x45, 0x37, 0xad, 0x75, 0xf1, 0x0c, 0xdb, 0x13, 0x91, 0x9f, 0x58, 0x5d, 0x96, 0x68,
	0xf7, 0x8d, 0xc2, 0xe2, 0x89, 0xe2, 0x24, 0x69, 0x0d, 0xbb, 0x15, 0x38, 0x1e, 0x38, 0x7e, 0x86,
	0x38, 0xdb, 0x8e, 0xe5, 0x3f, 0xca, 0x6a, 0x81, 0x2e, 0xb0, 0xd9, 0x95, 0xfe, 0x9e, 0x68, 0xad,
	0xb1, 0x31, 0x35, 0xe9, 0xa7, 0x6a, 0x74, 0x4e, 0x96, 0x18, 0x93, 0x37, 0x87, 0xc9, 0x56, 0xea,
	0xe6, 0x9d, 0xb1, 0xb5, 0x6c, 0xb5, 0x69, 0x82, 0x68, 0xf7, 0xf3, 0x08, 0xe0, 0x58, 0x5a, 0x49,
	0x59, 0xb4, 0x8e, 0x3d, 0x02, 0x68, 0x83, 0xa5, 0x42, 0xab, 0x74, 0xed, 0xc6, 0xda, 0x9d, 0x0b,
	0xf7, 0x26, 0x3c, 0x9a, 0x8f, 0xae, 0x74, 0xd3, 0x22, 0x09, 0x5e, 0xfa, 0x48, 0x9c, 0x8f, 0xf2,
	0x5c, 0xb1, 0xdb, 0xb0, 0xb5, 0x62, 0x8b, 0x66, 0x51, 0x55, 0xe9, 0xb7, 0x0d, 0xaa, 0xb0, 0x29,
	0x2e, 0xfe, 0x14, 0x1d, 0xd1, 0x2a, 0x7b, 0x01, 0x3b, 0xb5, 0xfc, 0x54, 0xcc, 0xde, 0x4b, 0xeb,
	0x8a, 0x39, 0xc9, 0x67, 0xfe, 0xf8, 0x15, 0x99, 0x4c, 0x47, 0x7f, 0xdf, 0x71, 0x9b, 0xd8, 0x7d,
	0x8f, 0x1e, 0x87, 0xbe, 0x3d, 0x27, 0x8e, 0x3d, 0x86, 0xab, 0x67, 0x94, 0x0c, 0x46, 0xbe, 0x07,
	0x23, 0x3b, 0x7f, 0x60, 0xbd, 0xa5, 0xdd, 0xaf, 0x23, 0xd8, 0x14, 0xe8, 0xe6, 0xa6, 0x71, 0xc8,
	0xee, 0xc2, 0xb8, 0x6b, 0x72, 0x3c, 0x7f, 0xc6, 0xfb, 0xc3, 0x0b, 0x03, 0x38, 0xf0, 0xbf, 0x22,
	0x08, 0xd9, 0x6b, 0xb8, 0xe4, 0xdb, 0x5b, 0xfc, 0xd2, 0x5f, 0x3a, 0x4a, 0x42, 0x30, 0x1f, 0xc0,
	0xc3, 0x29, 0x1c, 0x52, 0x9c, 0xaf, 0x62, 0xb1, 0x55, 0xf7, 0x17, 0xd8, 0x03, 0xd8, 0x88, 0x63,
	0x4d, 0x93, 0xae, 0xe2, 0xf5, 0xdf, 0x2a, 0x86, 0xa1, 0x1f, 0x86, 0x7f, 0xb1, 0x94, 0xb3, 0x03,
	0x48, 0xac, 0x39, 0x4d, 0xcf, 0x75, 0xd4, 0x1e, 0xff, 0xa7, 0x1b, 0xc8, 0x97, 0x4d, 0xe0, 0xc2,
	0x9c, 0x0a, 0xcf, 0x67, 0xd7, 0x20, 0xa1, 0x6f, 0x76, 0x19, 0xd6, 0x29, 0xf2, 0xb7, 0xe2, 0xcb,
	0x11, 0xb5, 0x65, 0x2c, 0xc6, 0x14, 0xe6, 0xea, 0xe9, 0x2b, 0x98, 0x68, 0x33, 0x28, 0xbe, 0x7a,
	0x16, 0x6f, 0x1e, 0x96, 0xc6, 0xa9, 0x0f, 0xcb, 0xbc, 0xfa, 0x8f, 0x97, 0xf3, 0x76, 0xbd, 0xbb,
	0xa1, 0x7b, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x16, 0x04, 0x3d, 0x74, 0x03, 0x00, 0x00,
}