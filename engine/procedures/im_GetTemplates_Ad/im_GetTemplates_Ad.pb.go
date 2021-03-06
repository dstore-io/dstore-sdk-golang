// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetTemplates_Ad.proto
// DO NOT EDIT!

/*
Package im_GetTemplates_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetTemplates_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetTemplates_Ad

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
	TreeNodeId           *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull       bool                          `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	LanguageId           *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull       bool                          `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	SourceTemplateId     *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=source_template_id,json=sourceTemplateId" json:"source_template_id,omitempty"`
	SourceTemplateIdNull bool                          `protobuf:"varint,1003,opt,name=source_template_id_null,json=sourceTemplateIdNull" json:"source_template_id_null,omitempty"`
	DateAndTime          *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=date_and_time,json=dateAndTime" json:"date_and_time,omitempty"`
	DateAndTimeNull      bool                          `protobuf:"varint,1004,opt,name=date_and_time_null,json=dateAndTimeNull" json:"date_and_time_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetSourceTemplateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SourceTemplateId
	}
	return nil
}

func (m *Parameters) GetDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.DateAndTime
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
	RowId       int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	Inherited   *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=inherited" json:"inherited,omitempty"`
	FrameName   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=frame_name,json=frameName" json:"frame_name,omitempty"`
	LanguageId  *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	TemplateId  *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=template_id,json=templateId" json:"template_id,omitempty"`
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

func (m *Response_Row) GetInherited() *dstore_values.BooleanValue {
	if m != nil {
		return m.Inherited
	}
	return nil
}

func (m *Response_Row) GetFrameName() *dstore_values.StringValue {
	if m != nil {
		return m.FrameName
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetTemplates_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetTemplates_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetTemplates_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/im_GetTemplates_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x9b, 0xa6, 0x97, 0x0d, 0xd0, 0xb0, 0x20, 0x08, 0x89, 0x40, 0x55, 0x79, 0x01, 0x84,
	0x1c, 0x6e, 0x42, 0x08, 0x51, 0x89, 0x20, 0x21, 0x94, 0x87, 0x5a, 0xc8, 0x54, 0x48, 0xf0, 0x62,
	0xb9, 0xd9, 0xa9, 0x59, 0x61, 0xef, 0x46, 0xbb, 0x1b, 0xfa, 0x1b, 0x5c, 0x7f, 0x0c, 0xf1, 0x11,
	0xdc, 0x3e, 0x82, 0x59, 0xaf, 0x8d, 0x2f, 0x89, 0xd4, 0xf4, 0x21, 0x8e, 0xc6, 0x73, 0xce, 0x9c,
	0xf1, 0xce, 0x99, 0x25, 0x77, 0x99, 0x36, 0x52, 0xc1, 0x10, 0x44, 0xcc, 0x05, 0x0c, 0xa7, 0x4a,
	0x4e, 0x80, 0xcd, 0x14, 0xe8, 0x21, 0x4f, 0xc3, 0x17, 0x60, 0x0e, 0x20, 0x9d, 0x26, 0x91, 0x01,
	0x1d, 0x8e, 0x98, 0x87, 0x59, 0x23, 0xe9, 0x8e, 0xa3, 0x78, 0x8e, 0xe2, 0xcd, 0xe3, 0xfa, 0x17,
	0xf2, 0xa2, 0x1f, 0xa2, 0x64, 0x06, 0xda, 0xd1, 0xfa, 0x57, 0xea, 0x4a, 0xa0, 0x94, 0x54, 0x79,
	0x6a, 0x50, 0x4f, 0xa5, 0xa0, 0x75, 0x14, 0x43, 0x9e, 0xbc, 0xde, 0x4c, 0x9a, 0x88, 0x8b, 0x23,
	0xa9, 0xd2, 0xc8, 0x70, 0x29, 0x1c, 0x68, 0xf7, 0x47, 0x8b, 0x90, 0x97, 0x91, 0x8a, 0x30, 0x0b,
	0x4a, 0xd3, 0x3d, 0x72, 0xc6, 0x28, 0x80, 0x50, 0x48, 0x06, 0x21, 0x67, 0xbd, 0x95, 0x9d, 0x95,
	0x1b, 0x9d, 0x7b, 0x03, 0x2f, 0xef, 0x3c, 0xef, 0x8b, 0x0b, 0x03, 0x31, 0xa8, 0xd7, 0x36, 0x0a,
	0x88, 0x25, 0xf8, 0x88, 0x1f, 0x33, 0x7a, 0x8b, 0x9c, 0xaf, 0xd2, 0x43, 0x31, 0x4b, 0x92, 0xde,
	0xaf, 0x0d, 0x2c, 0xb2, 0x19, 0x9c, 0x2b, 0x71, 0x3e, 0xbe, 0xa6, 0x4f, 0x48, 0x27, 0x89, 0x44,
	0x3c, 0xc3, 0x86, 0xad, 0xd2, 0xea, 0x12, 0x4a, 0x05, 0x1e, 0x95, 0x6e, 0x92, 0x6e, 0x85, 0xed,
	0x84, 0x7e, 0xe7, 0x42, 0x25, 0x2c, 0x13, 0x1a, 0x13, 0xaa, 0xe5, 0x4c, 0x4d, 0x20, 0x34, 0xf9,
	0x59, 0x5b, 0xbd, 0xd6, 0xc9, 0x7a, 0x5d, 0x47, 0x2b, 0x26, 0x84, 0xaa, 0x0f, 0xc9, 0xe5, 0xf9,
	0x52, 0x4e, 0xfc, 0x8f, 0x13, 0xbf, 0xd8, 0xe4, 0x64, 0x2d, 0x8c, 0xc8, 0x59, 0x66, 0xc1, 0x91,
	0x60, 0xa1, 0xe1, 0x29, 0xf4, 0xd6, 0x32, 0xf5, 0xab, 0x0d, 0x75, 0x9b, 0xd2, 0x26, 0x4a, 0xa7,
	0x4e, 0xbf, 0x63, 0x39, 0x23, 0xc1, 0x0e, 0xf0, 0x35, 0xbd, 0x4d, 0x68, 0xad, 0x84, 0x53, 0xfd,
	0xeb, 0x54, 0xb7, 0x2b, 0x48, 0x2b, 0xb8, 0xfb, 0x73, 0x8d, 0x6c, 0x06, 0xa0, 0xa7, 0x52, 0x68,
	0xa0, 0x77, 0x48, 0x3b, 0x33, 0x4d, 0x3e, 0xcd, 0xbe, 0x57, 0xf7, 0xa1, 0x33, 0xd4, 0x73, 0xfb,
	0x0c, 0x1c, 0x90, 0xbe, 0x21, 0x5d, 0x6b, 0x97, 0xb0, 0xe2, 0x17, 0x1c, 0x50, 0x0b, 0xc9, 0x5e,
	0x83, 0xdc, 0x74, 0xd5, 0x3e, 0xc6, 0xe3, 0x32, 0x0e, 0xb6, 0xd3, 0xfa, 0x0b, 0xfa, 0x88, 0x6c,
	0xe4, 0x36, 0xc5, 0x11, 0xd8, 0x8a, 0xd7, 0xe6, 0x2a, 0x3a, 0x13, 0xef, 0xbb, 0xff, 0xa0, 0x80,
	0xd3, 0xa7, 0xa4, 0xa5, 0xe4, 0x31, 0x1e, 0xdd, 0xa2, 0x3e, 0x16, 0x2c, 0x5d, 0xf1, 0xfd, 0x5e,
	0x20, 0x8f, 0x03, 0x4b, 0xed, 0x7f, 0x5f, 0x25, 0x2d, 0x0c, 0xe8, 0x25, 0xb2, 0x8e, 0xa1, 0x75,
	0xc1, 0x47, 0x1f, 0x8f, 0xa4, 0x1d, 0xb4, 0x31, 0xc4, 0xf1, 0xee, 0x91, 0x0e, 0x03, 0x3d, 0x51,
	0x7c, 0x9a, 0x7d, 0xf1, 0x27, 0xbf, 0x7e, 0x5e, 0xf9, 0x94, 0xb4, 0x51, 0x5c, 0xc4, 0xc5, 0x88,
	0x4a, 0x3c, 0x7d, 0x4c, 0xb6, 0xb8, 0x78, 0x07, 0x8a, 0x1b, 0x60, 0xbd, 0xcf, 0xfe, 0x42, 0x83,
	0x1d, 0x4a, 0x99, 0x40, 0x24, 0x1c, 0xbb, 0x84, 0x23, 0x97, 0x1c, 0xd9, 0x2d, 0x0c, 0x05, 0x3e,
	0x7a, 0x5f, 0x4e, 0x56, 0xde, 0xca, 0xe0, 0x3e, 0xfe, 0x6c, 0xdb, 0xd5, 0x4d, 0xfa, 0xea, 0x9f,
	0x6e, 0x95, 0x90, 0x5e, 0x5d, 0x8c, 0x6f, 0xcb, 0xd0, 0xcd, 0x7f, 0x7f, 0x3f, 0x7b, 0x45, 0x06,
	0x5c, 0x36, 0xa6, 0x51, 0xde, 0x86, 0x6f, 0x1f, 0xc4, 0x52, 0xb3, 0xf7, 0x45, 0x9e, 0x2d, 0x77,
	0x61, 0x1e, 0xae, 0x67, 0xb7, 0xd3, 0xfd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xef, 0xca,
	0xb5, 0x66, 0x05, 0x00, 0x00,
}
