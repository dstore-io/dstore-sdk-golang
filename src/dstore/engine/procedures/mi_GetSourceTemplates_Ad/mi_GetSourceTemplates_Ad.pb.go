// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetSourceTemplates_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetSourceTemplates_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetSourceTemplates_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetSourceTemplates_Ad

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
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	RowId       int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	TemplateId  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=template_id,json=templateId" json:"template_id,omitempty"`
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

func (m *Response_Row) GetTemplateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetSourceTemplates_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetSourceTemplates_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetSourceTemplates_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xc7, 0xe9, 0xcd, 0x6d, 0xef, 0x65, 0x7a, 0xe1, 0x4a, 0x04, 0x89, 0x29, 0x48, 0xa9, 0x20,
	0xae, 0xa6, 0x7e, 0x81, 0x6e, 0x5c, 0x28, 0xa8, 0x74, 0xd1, 0x52, 0x46, 0x11, 0x74, 0x13, 0x62,
	0xe7, 0x18, 0x06, 0x9a, 0x99, 0x70, 0x66, 0x6a, 0x5f, 0x43, 0xdd, 0xfa, 0x20, 0xbe, 0x9e, 0x93,
	0x4c, 0x4a, 0x9b, 0x88, 0xa0, 0x9b, 0x96, 0x7f, 0xce, 0xff, 0x77, 0xbe, 0xe6, 0x90, 0x63, 0xae,
	0x8d, 0x42, 0xe8, 0x83, 0x4c, 0x84, 0x84, 0x7e, 0x86, 0x6a, 0x02, 0x7c, 0x86, 0xa0, 0xfb, 0xa9,
	0x88, 0xae, 0xc0, 0x5c, 0xab, 0x19, 0x4e, 0xe0, 0x06, 0xd2, 0x6c, 0x1a, 0x1b, 0xd0, 0xd1, 0x19,
	0xa7, 0xd6, 0x63, 0x94, 0xbf, 0xe3, 0x40, 0xea, 0x40, 0xfa, 0x95, 0x3b, 0x5c, 0x2f, 0x0b, 0x3c,
	0xc5, 0xd3, 0x19, 0x68, 0x07, 0x87, 0x9b, 0xd5, 0xaa, 0x80, 0xa8, 0xb0, 0x0c, 0x75, 0xaa, 0xa1,
	0x14, 0xb4, 0x8e, 0x13, 0x28, 0x83, 0xdb, 0xf5, 0xa0, 0x89, 0x85, 0x7c, 0x54, 0x98, 0xc6, 0x46,
	0x28, 0xe9, 0x4c, 0xbd, 0x7f, 0x84, 0x8c, 0x63, 0x8c, 0x6d, 0x10, 0x50, 0xf7, 0xde, 0x3d, 0xf2,
	0x97, 0x81, 0xce, 0x94, 0xd4, 0xe0, 0xef, 0x91, 0x66, 0x51, 0x2b, 0x68, 0x74, 0x1b, 0xbb, 0xed,
	0x83, 0x90, 0x56, 0x87, 0x70, 0x7d, 0x5c, 0xe4, 0xbf, 0xcc, 0x19, 0xfd, 0x3b, 0xb2, 0x96, 0x57,
	0x89, 0x56, 0xca, 0x04, 0xbf, 0xba, 0x9e, 0x85, 0x69, 0x0d, 0xae, 0x37, 0x33, 0xb4, 0x7a, 0xb0,
	0xd4, 0xec, 0x7f, 0x5a, 0xfd, 0xe0, 0x9f, 0x90, 0x3f, 0xe5, 0x74, 0x81, 0x57, 0x64, 0xdc, 0xfa,
	0x94, 0xd1, 0xcd, 0x3e, 0x74, 0xff, 0x6c, 0x61, 0xf7, 0x2f, 0x89, 0x87, 0x6a, 0x1e, 0xfc, 0x2e,
	0xa8, 0x23, 0xfa, 0xbd, 0x97, 0xa0, 0x8b, 0x2d, 0x50, 0xa6, 0xe6, 0x2c, 0x4f, 0x10, 0xbe, 0x35,
	0x88, 0x67, 0x85, 0xbf, 0x41, 0x5a, 0x56, 0x46, 0x82, 0x07, 0xcf, 0x23, 0xbb, 0x98, 0x26, 0x6b,
	0x5a, 0x39, 0xe0, 0xfe, 0x29, 0x69, 0x73, 0xd0, 0x13, 0x14, 0x59, 0x31, 0xf7, 0xcb, 0xa8, 0xba,
	0xb5, 0xf2, 0x49, 0xb5, 0x41, 0x21, 0x93, 0xdb, 0x5c, 0xb0, 0x55, 0x7f, 0x8e, 0x9b, 0xb2, 0x81,
	0x3c, 0xf7, 0xab, 0xc3, 0x3b, 0x35, 0x5c, 0x48, 0x03, 0x09, 0xa0, 0xe3, 0xc9, 0x02, 0x18, 0xf0,
	0xf3, 0x31, 0xe9, 0x08, 0x55, 0x1b, 0x6e, 0x79, 0x9f, 0xf7, 0xfb, 0x3f, 0xbe, 0xdc, 0x87, 0x56,
	0x71, 0x20, 0x87, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x10, 0x5d, 0xd2, 0xf5, 0x02, 0x00,
	0x00,
}
