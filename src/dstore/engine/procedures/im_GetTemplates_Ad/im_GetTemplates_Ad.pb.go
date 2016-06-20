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

var fileDescriptor0 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xeb, 0x8a, 0x13, 0x4d,
	0x10, 0x65, 0x37, 0x7b, 0xad, 0x7c, 0x9f, 0xbb, 0xb6, 0xa2, 0x31, 0x41, 0x59, 0xd6, 0x3f, 0xde,
	0x98, 0x78, 0x01, 0x11, 0x71, 0xc1, 0x08, 0x22, 0x11, 0x76, 0x90, 0x61, 0x11, 0xf4, 0xcf, 0x30,
	0x9b, 0xae, 0x8d, 0x0d, 0x33, 0xdd, 0xa1, 0xbb, 0xe3, 0xbe, 0x86, 0xd7, 0x17, 0x13, 0x1f, 0xc2,
	0xdb, 0x43, 0x58, 0x3d, 0x3d, 0x71, 0x2e, 0x09, 0xec, 0xee, 0x8f, 0x4c, 0xa8, 0xa9, 0x73, 0xea,
	0xd4, 0x74, 0x9d, 0x6a, 0xb8, 0xc7, 0x8d, 0x55, 0x1a, 0xfb, 0x28, 0xc7, 0x42, 0x62, 0x7f, 0xa2,
	0xd5, 0x08, 0xf9, 0x54, 0xa3, 0xe9, 0x8b, 0x2c, 0x7e, 0x81, 0xf6, 0x00, 0xb3, 0x49, 0x9a, 0x58,
	0x34, 0xf1, 0x80, 0x07, 0x94, 0xb5, 0x8a, 0xed, 0x78, 0x4a, 0xe0, 0x29, 0xc1, 0x3c, 0xae, 0x7b,
	0xa1, 0x28, 0xfa, 0x3e, 0x49, 0xa7, 0x68, 0x3c, 0xad, 0x7b, 0xa5, 0xae, 0x84, 0x5a, 0x2b, 0x5d,
	0xa4, 0x7a, 0xf5, 0x54, 0x86, 0xc6, 0x24, 0x63, 0x2c, 0x92, 0xd7, 0x9b, 0x49, 0x9b, 0x08, 0x79,
	0xa4, 0x74, 0x96, 0x58, 0xa1, 0xa4, 0x07, 0xed, 0x7e, 0x6f, 0x01, 0xbc, 0x4a, 0x74, 0x42, 0x59,
	0xd4, 0x86, 0xed, 0xc1, 0x7f, 0x56, 0x23, 0xc6, 0x52, 0x71, 0x8c, 0x05, 0xef, 0x2c, 0xed, 0x2c,
	0xdd, 0x68, 0xdf, 0xef, 0x05, 0x45, 0xe7, 0x45, 0x5f, 0x42, 0x5a, 0x1c, 0xa3, 0x7e, 0xed, 0xa2,
	0x08, 0x1c, 0x21, 0x24, 0xfc, 0x90, 0xb3, 0x5b, 0x70, 0xbe, 0x4a, 0x8f, 0xe5, 0x34, 0x4d, 0x3b,
	0x3f, 0xd7, 0xa9, 0xc8, 0x46, 0x74, 0xae, 0xc4, 0x85, 0xf4, 0x9a, 0x3d, 0x81, 0x76, 0x9a, 0xc8,
	0xf1, 0x94, 0x1a, 0x76, 0x4a, 0xcb, 0xa7, 0x50, 0x9a, 0xe1, 0x49, 0xe9, 0x26, 0x6c, 0x57, 0xd8,
	0x5e, 0xe8, 0x57, 0x21, 0x54, 0xc2, 0x72, 0xa1, 0x21, 0x30, 0xa3, 0xa6, 0x7a, 0x84, 0xb1, 0x2d,
	0xce, 0xda, 0xe9, 0xb5, 0x4e, 0xd6, 0xdb, 0xf6, 0xb4, 0xd9, 0x84, 0x48, 0xf5, 0x21, 0x5c, 0x9e,
	0x2f, 0xe5, 0xc5, 0x7f, 0x7b, 0xf1, 0x8b, 0x4d, 0x4e, 0xde, 0xc2, 0x00, 0xfe, 0xe7, 0x0e, 0x9c,
	0x48, 0x1e, 0x5b, 0x91, 0x61, 0x67, 0x25, 0x57, 0xbf, 0xda, 0x50, 0x77, 0x29, 0x63, 0x93, 0x6c,
	0xe2, 0xf5, 0xdb, 0x8e, 0x33, 0x90, 0xfc, 0x80, 0x5e, 0xb3, 0x3b, 0xc0, 0x6a, 0x25, 0xbc, 0xea,
	0x1f, 0xaf, 0xba, 0x55, 0x41, 0x3a, 0xc1, 0xdd, 0x1f, 0x2b, 0xb0, 0x11, 0xa1, 0x99, 0x28, 0x69,
	0x90, 0xdd, 0x85, 0xd5, 0xdc, 0x34, 0xc5, 0x34, 0xbb, 0x41, 0xdd, 0x87, 0xde, 0x50, 0xcf, 0xdd,
	0x33, 0xf2, 0x40, 0xf6, 0x06, 0xb6, 0x9d, 0x5d, 0xe2, 0x8a, 0x5f, 0x68, 0x40, 0x2d, 0x22, 0x07,
	0x0d, 0x72, 0xd3, 0x55, 0xfb, 0x14, 0x0f, 0xcb, 0x38, 0xda, 0xca, 0xea, 0x2f, 0xd8, 0x23, 0x58,
	0x2f, 0x6c, 0x4a, 0x23, 0x70, 0x15, 0xaf, 0xcd, 0x55, 0xf4, 0x26, 0xde, 0xf7, 0xff, 0xd1, 0x0c,
	0xce, 0x9e, 0x42, 0x4b, 0xab, 0x63, 0x3a, 0xba, 0x45, 0x7d, 0x2c, 0x58, 0xba, 0xd9, 0xf7, 0x07,
	0x91, 0x3a, 0x8e, 0x1c, 0xb5, 0xfb, 0x6d, 0x19, 0x5a, 0x14, 0xb0, 0x4b, 0xb0, 0x46, 0xa1, 0x73,
	0xc1, 0x87, 0x90, 0x8e, 0x64, 0x35, 0x5a, 0xa5, 0x90, 0xc6, 0xbb, 0x07, 0x6d, 0x8e, 0x66, 0xa4,
	0xc5, 0x24, 0xff, 0xe2, 0x8f, 0x61, 0xfd, 0xbc, 0x8a, 0x29, 0x19, 0xab, 0x85, 0x1c, 0xcf, 0x46,
	0x54, 0xe2, 0xd9, 0x63, 0xd8, 0x14, 0xf2, 0x1d, 0x6a, 0x61, 0x91, 0x77, 0x3e, 0x85, 0x0b, 0x0d,
	0x76, 0xa8, 0x54, 0x8a, 0x89, 0xf4, 0xec, 0x12, 0x4e, 0x5c, 0x38, 0x72, 0x5b, 0x18, 0x4b, 0x7a,
	0x74, 0x3e, 0x9f, 0xac, 0xbc, 0x99, 0xc3, 0x43, 0xfa, 0xb9, 0xb6, 0xab, 0x9b, 0xf4, 0x25, 0x3c,
	0xdb, 0x2a, 0x11, 0xbd, 0xba, 0x18, 0x5f, 0x4f, 0x43, 0xb7, 0xff, 0xfc, 0xfd, 0xec, 0x25, 0xf4,
	0x84, 0x6a, 0x4c, 0xa3, 0xbc, 0x0d, 0xdf, 0xde, 0x3e, 0xc3, 0x3d, 0x79, 0xb8, 0x96, 0x5f, 0x4a,
	0x0f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x3f, 0xc6, 0x89, 0x5d, 0x05, 0x00, 0x00,
}