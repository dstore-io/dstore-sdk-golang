// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetTemplates_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetTemplates_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetTemplates_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetTemplates_Ad

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
	TemplateId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=template_id,json=templateId" json:"template_id,omitempty"`
	TemplateIdNull bool                        `protobuf:"varint,1001,opt,name=template_id_null,json=templateIdNull" json:"template_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTemplateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TemplateId
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
	RowId               int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description         *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	FrameName           *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=frame_name,json=frameName" json:"frame_name,omitempty"`
	FileName            *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	CanBeSourceTemplate *dstore_values.BooleanValue `protobuf:"bytes,10004,opt,name=can_be_source_template,json=canBeSourceTemplate" json:"can_be_source_template,omitempty"`
	IsRealTimeTemplate  *dstore_values.BooleanValue `protobuf:"bytes,10005,opt,name=is_real_time_template,json=isRealTimeTemplate" json:"is_real_time_template,omitempty"`
	TemplateId          *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=template_id,json=templateId" json:"template_id,omitempty"`
	LanguageId          *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
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

func (m *Response_Row) GetFrameName() *dstore_values.StringValue {
	if m != nil {
		return m.FrameName
	}
	return nil
}

func (m *Response_Row) GetFileName() *dstore_values.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *Response_Row) GetCanBeSourceTemplate() *dstore_values.BooleanValue {
	if m != nil {
		return m.CanBeSourceTemplate
	}
	return nil
}

func (m *Response_Row) GetIsRealTimeTemplate() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsRealTimeTemplate
	}
	return nil
}

func (m *Response_Row) GetTemplateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetTemplates_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetTemplates_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetTemplates_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetTemplates_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x55, 0x4c, 0xd3, 0x74, 0x23, 0x41, 0xb5, 0x15, 0x95, 0x49, 0x24, 0x14, 0x95, 0x0b,
	0x5c, 0x1c, 0xbe, 0x0e, 0x08, 0x51, 0x09, 0x2a, 0x21, 0x94, 0x43, 0x0d, 0xb8, 0x15, 0x12, 0x5c,
	0x56, 0x9b, 0x78, 0x62, 0xad, 0xb0, 0x77, 0xa3, 0xdd, 0x35, 0x7d, 0x07, 0x4e, 0x7c, 0xf3, 0x8c,
	0xdc, 0x79, 0x00, 0xc6, 0xde, 0x35, 0x89, 0x93, 0x4a, 0x0d, 0x87, 0x38, 0x1a, 0xcf, 0xff, 0x37,
	0xde, 0xf9, 0xcf, 0xd8, 0xe4, 0x7e, 0x6a, 0xac, 0xd2, 0x30, 0x02, 0x99, 0x09, 0x09, 0xa3, 0xb9,
	0x56, 0x53, 0x48, 0x4b, 0x0d, 0x66, 0x54, 0x08, 0xf6, 0x12, 0xec, 0x19, 0x14, 0xf3, 0x9c, 0x5b,
	0x30, 0xec, 0x79, 0x1a, 0x61, 0xd6, 0x2a, 0x3a, 0x74, 0x48, 0xe4, 0x90, 0x68, 0x5d, 0xd7, 0xdf,
	0xf7, 0x45, 0x3f, 0xf2, 0xbc, 0x04, 0xe3, 0xb0, 0xfe, 0xcd, 0xf6, 0x93, 0x40, 0x6b, 0xa5, 0x7d,
	0x6a, 0xd0, 0x4e, 0x15, 0x60, 0x0c, 0xcf, 0xc0, 0x27, 0x6f, 0xaf, 0x26, 0x2d, 0x17, 0x72, 0xa6,
	0x74, 0xc1, 0xad, 0x50, 0xd2, 0x89, 0x0e, 0x4b, 0x42, 0x5e, 0x73, 0xcd, 0x31, 0x09, 0xda, 0xd0,
	0xa7, 0xa4, 0x67, 0xfd, 0x79, 0x98, 0x48, 0xc3, 0xad, 0xe1, 0xd6, 0x9d, 0xde, 0x83, 0x41, 0xe4,
	0xcf, 0xed, 0x4f, 0x25, 0xa4, 0x85, 0x0c, 0xf4, 0xdb, 0x2a, 0x4a, 0x48, 0xa3, 0x1f, 0xa7, 0xf4,
	0x2e, 0xd9, 0x5b, 0xa2, 0x99, 0x2c, 0xf3, 0x3c, 0xfc, 0xbd, 0x83, 0x35, 0xba, 0xc9, 0xb5, 0x85,
	0x2c, 0xc6, 0xdb, 0x87, 0x9f, 0x3a, 0xa4, 0x9b, 0x80, 0x99, 0x2b, 0x69, 0x80, 0xde, 0x23, 0xdb,
	0x75, 0x53, 0xfe, 0x79, 0xfd, 0xa8, 0xed, 0x93, 0x6b, 0xf8, 0x45, 0x75, 0x4d, 0x9c, 0x90, 0xbe,
	0x23, 0x7b, 0x55, 0x3b, 0x6c, 0xa9, 0x9f, 0xf0, 0xca, 0x30, 0x40, 0x38, 0x5a, 0x81, 0x57, 0xbb,
	0x3e, 0xc1, 0x78, 0xbc, 0x88, 0x93, 0xeb, 0x45, 0xfb, 0x06, 0x7d, 0x4c, 0x76, 0xbc, 0x8d, 0x61,
	0x50, 0x57, 0xbc, 0xb5, 0x56, 0xd1, 0x99, 0x7c, 0xe2, 0xfe, 0x93, 0x46, 0x4e, 0x9f, 0x91, 0x40,
	0xab, 0xf3, 0xf0, 0xea, 0xc5, 0xe7, 0x58, 0x5f, 0x8a, 0xa6, 0xff, 0x28, 0x51, 0xe7, 0x49, 0x85,
	0xf6, 0xff, 0x04, 0x24, 0xc0, 0x80, 0x1e, 0x90, 0x0e, 0x86, 0xd5, 0x04, 0x3e, 0xc7, 0x68, 0xc9,
	0x76, 0xb2, 0x8d, 0x21, 0x1a, 0x7c, 0x44, 0x7a, 0x29, 0x98, 0xa9, 0x16, 0xf3, 0xba, 0xe3, 0x2f,
	0x71, 0xdb, 0x2f, 0x3f, 0x1f, 0x63, 0xb5, 0x90, 0x99, 0x1b, 0xcf, 0xb2, 0x9e, 0x3e, 0x21, 0x64,
	0x56, 0x8d, 0x9a, 0x49, 0xbc, 0x84, 0x5f, 0x2f, 0xa7, 0x77, 0x6b, 0x79, 0x8c, 0x3f, 0xb4, 0x65,
	0x77, 0x26, 0x72, 0x8f, 0x7e, 0xbb, 0x1c, 0xed, 0x56, 0xea, 0x9a, 0x7c, 0x43, 0x0e, 0xa6, 0x5c,
	0xb2, 0x09, 0x30, 0xa3, 0x4a, 0x3d, 0x05, 0xd6, 0xac, 0x42, 0xf8, 0x3d, 0xbe, 0x70, 0xbf, 0x26,
	0x4a, 0xe5, 0xc0, 0xa5, 0xab, 0xb3, 0x8f, 0xec, 0x31, 0x9c, 0xd6, 0x64, 0xe3, 0x1e, 0x7d, 0x45,
	0x6e, 0x08, 0xc3, 0x34, 0xf0, 0x9c, 0x59, 0x51, 0x2c, 0x55, 0xfc, 0xb1, 0x41, 0x45, 0x2a, 0x4c,
	0x82, 0xe4, 0x19, 0x82, 0xff, 0x0a, 0x1e, 0xb5, 0xf7, 0xfe, 0x67, 0xfc, 0x7f, 0x8b, 0x8f, 0x78,
	0xce, 0x65, 0x56, 0xe2, 0x16, 0x54, 0xf8, 0xaf, 0x4d, 0xf0, 0x06, 0x18, 0xa7, 0xc7, 0xa7, 0x64,
	0x20, 0xd4, 0xca, 0xbe, 0x2c, 0xbe, 0x27, 0xef, 0x1f, 0x65, 0xca, 0xa4, 0x1f, 0x9a, 0x7c, 0xba,
	0xd9, 0x27, 0x67, 0xd2, 0xa9, 0xdf, 0xef, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab, 0x2f,
	0xbc, 0x0c, 0xa8, 0x04, 0x00, 0x00,
}