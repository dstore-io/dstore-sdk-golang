// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetTemplates_Pu.proto
// DO NOT EDIT!

/*
Package im_GetTemplates_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetTemplates_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetTemplates_Pu

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
	Parameter            *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=parameter" json:"parameter,omitempty"`
	ParameterNull        bool                        `protobuf:"varint,1001,opt,name=parameter_null,json=parameterNull" json:"parameter_null,omitempty"`
	Type                 *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	TypeNull             bool                        `protobuf:"varint,1002,opt,name=type_null,json=typeNull" json:"type_null,omitempty"`
	LanguageId           *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull       bool                        `protobuf:"varint,1003,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	SourceTemplateId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=source_template_id,json=sourceTemplateId" json:"source_template_id,omitempty"`
	SourceTemplateIdNull bool                        `protobuf:"varint,1004,opt,name=source_template_id_null,json=sourceTemplateIdNull" json:"source_template_id_null,omitempty"`
	FrameName            *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=frame_name,json=frameName" json:"frame_name,omitempty"`
	FrameNameNull        bool                        `protobuf:"varint,1005,opt,name=frame_name_null,json=frameNameNull" json:"frame_name_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetParameter() *dstore_values.StringValue {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *Parameters) GetType() *dstore_values.StringValue {
	if m != nil {
		return m.Type
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

func (m *Parameters) GetFrameName() *dstore_values.StringValue {
	if m != nil {
		return m.FrameName
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
	RowId              int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description        *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	FrameName          *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=frame_name,json=frameName" json:"frame_name,omitempty"`
	FileName           *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	TreeNodeId         *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId             *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	IsRealTimeTemplate *dstore_values.BooleanValue `protobuf:"bytes,10006,opt,name=is_real_time_template,json=isRealTimeTemplate" json:"is_real_time_template,omitempty"`
	TemplateId         *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=template_id,json=templateId" json:"template_id,omitempty"`
	LanguageId         *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
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

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetTemplates_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetTemplates_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetTemplates_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x6b, 0x6b, 0xd4, 0x4c,
	0x18, 0xa5, 0xef, 0x5e, 0xba, 0x7d, 0xfa, 0xda, 0x96, 0xf1, 0x16, 0xb7, 0x22, 0xa5, 0x82, 0x17,
	0x84, 0xd4, 0x0b, 0x8a, 0x8a, 0x15, 0x11, 0x44, 0x2a, 0x34, 0x96, 0x50, 0x04, 0xfd, 0x12, 0xd2,
	0xcd, 0xec, 0x32, 0x90, 0xcc, 0x84, 0x99, 0x89, 0xc5, 0xff, 0xe0, 0x07, 0xef, 0xf6, 0x57, 0xf8,
	0xbf, 0xbc, 0xfd, 0x07, 0x9f, 0xc9, 0x24, 0x9b, 0xdd, 0xec, 0xc2, 0xee, 0x7e, 0xd8, 0x0d, 0xcf,
	0xcc, 0x39, 0xe7, 0x99, 0xc9, 0x39, 0x33, 0x81, 0x5b, 0x91, 0xd2, 0x42, 0xd2, 0x1d, 0xca, 0x07,
	0x8c, 0xd3, 0x9d, 0x54, 0x8a, 0x1e, 0x8d, 0x32, 0x49, 0xd5, 0x0e, 0x4b, 0x82, 0xe7, 0x54, 0x1f,
	0xd2, 0x24, 0x8d, 0x43, 0x4d, 0x55, 0x70, 0x90, 0xb9, 0x38, 0xab, 0x05, 0xd9, 0xb2, 0x14, 0xd7,
	0x52, 0xdc, 0x49, 0x5c, 0xf7, 0x74, 0x21, 0xfa, 0x36, 0x8c, 0x33, 0xaa, 0x2c, 0xad, 0x7b, 0x61,
	0xbc, 0x13, 0x95, 0x52, 0xc8, 0x62, 0x6a, 0x73, 0x7c, 0x2a, 0xa1, 0x4a, 0x85, 0x03, 0x5a, 0x4c,
	0x5e, 0xae, 0x4f, 0xea, 0x90, 0xf1, 0xbe, 0x90, 0x49, 0xa8, 0x99, 0xe0, 0x16, 0xb4, 0xfd, 0xbe,
	0x09, 0x70, 0x10, 0xca, 0x10, 0x67, 0xa9, 0x54, 0xe4, 0x3e, 0xac, 0xa4, 0x65, 0xe5, 0x2c, 0x6d,
	0x2d, 0x5d, 0x5b, 0xbd, 0xdd, 0x75, 0x8b, 0x65, 0x17, 0x8b, 0x52, 0x5a, 0x32, 0x3e, 0x78, 0x65,
	0x0a, 0xbf, 0x02, 0x93, 0x2b, 0xb0, 0x36, 0x2c, 0x02, 0x9e, 0xc5, 0xb1, 0xf3, 0x73, 0x19, 0xf9,
	0x1d, 0xff, 0xd4, 0x70, 0xd8, 0xc3, 0x51, 0xe2, 0x42, 0x53, 0xbf, 0x4b, 0xa9, 0xf3, 0xdf, 0x4c,
	0xf1, 0x1c, 0x47, 0x2e, 0xc2, 0x8a, 0x79, 0x5a, 0xc9, 0x5f, 0x56, 0xb2, 0x63, 0x46, 0x72, 0xb5,
	0x47, 0xb0, 0x1a, 0x87, 0x7c, 0x90, 0xe1, 0xae, 0x03, 0x16, 0x39, 0x8d, 0x5c, 0x74, 0xb3, 0x26,
	0xca, 0xb8, 0xa6, 0x03, 0x2a, 0xad, 0x2a, 0x94, 0xf8, 0xbd, 0x88, 0x5c, 0x87, 0x8d, 0x11, 0xb6,
	0x6d, 0xf1, 0xdb, 0xb6, 0x58, 0xab, 0x60, 0x79, 0xa3, 0x3d, 0x20, 0x4a, 0x64, 0xb2, 0x47, 0x03,
	0x5d, 0x18, 0x66, 0xfa, 0x35, 0x67, 0xf7, 0xdb, 0xb0, 0xb4, 0xd2, 0x66, 0xec, 0x7a, 0x0f, 0xce,
	0x4f, 0x4a, 0xd9, 0xe6, 0x7f, 0x6c, 0xf3, 0x33, 0x75, 0x4e, 0xbe, 0x84, 0x07, 0x00, 0x7d, 0xf3,
	0x26, 0x03, 0x8e, 0x7f, 0x4e, 0x6b, 0xb6, 0x39, 0x39, 0xda, 0xc3, 0x1f, 0xb9, 0x0a, 0xeb, 0x15,
	0xd5, 0xb6, 0xfa, 0x5b, 0xb8, 0x33, 0x04, 0x99, 0x1e, 0xdb, 0x3f, 0xda, 0xd0, 0xf1, 0xa9, 0x4a,
	0x05, 0x57, 0x94, 0xdc, 0x84, 0x56, 0x1e, 0xb6, 0x7a, 0x10, 0x8a, 0xfc, 0xda, 0x20, 0x3e, 0x33,
	0xff, 0xbe, 0x05, 0x92, 0xd7, 0xb0, 0x61, 0x62, 0x16, 0x8c, 0xe4, 0x0c, 0x8d, 0x6e, 0x20, 0xd9,
	0xad, 0x91, 0xeb, 0x69, 0xdc, 0xc7, 0x7a, 0xaf, 0xaa, 0xfd, 0xf5, 0x64, 0x7c, 0x00, 0x93, 0xb9,
	0x5c, 0xc4, 0x1b, 0x5d, 0x36, 0x8a, 0x97, 0x26, 0x14, 0x6d, 0xf8, 0xf7, 0xed, 0xd3, 0x2f, 0xe1,
	0xe4, 0x09, 0x34, 0xa4, 0x38, 0x46, 0xaf, 0xa6, 0xad, 0x63, 0xca, 0x61, 0x2d, 0xf7, 0xef, 0xfa,
	0xe2, 0xd8, 0x37, 0xd4, 0xee, 0x49, 0x13, 0x1a, 0x58, 0x90, 0x73, 0xd0, 0xc6, 0xd2, 0x18, 0xff,
	0xc1, 0xc3, 0x57, 0xd2, 0xf2, 0x5b, 0x58, 0xa2, 0xa3, 0xbb, 0xb0, 0x1a, 0x51, 0xd5, 0x93, 0x2c,
	0xcd, 0x77, 0xfc, 0xd1, 0x9b, 0xe9, 0xcd, 0x28, 0x9e, 0x3c, 0x1c, 0x33, 0xf6, 0x93, 0xb7, 0x88,
	0xb3, 0x78, 0x60, 0xfb, 0x2c, 0x2e, 0xa8, 0x9f, 0x67, 0x53, 0x3b, 0x06, 0x9d, 0x33, 0x1f, 0xc3,
	0xff, 0x5a, 0x52, 0x64, 0x8a, 0x28, 0xcf, 0xf2, 0x17, 0x6f, 0x8e, 0xc3, 0x63, 0x18, 0x1e, 0x12,
	0x70, 0xd3, 0x77, 0x61, 0xb9, 0xa4, 0x7e, 0x9d, 0x83, 0xda, 0xe6, 0x96, 0xf6, 0x12, 0xce, 0x32,
	0x15, 0x48, 0x1a, 0xc6, 0x81, 0x66, 0x49, 0x75, 0x06, 0x9c, 0x6f, 0xd3, 0x45, 0x8e, 0x84, 0x88,
	0x69, 0xc8, 0xad, 0x08, 0x61, 0xca, 0x47, 0xe6, 0x21, 0x12, 0x4b, 0xcb, 0xcc, 0xcb, 0x1f, 0x3d,
	0x92, 0xdf, 0xe7, 0xda, 0x46, 0x75, 0x1a, 0x77, 0xc7, 0x6f, 0x90, 0x13, 0x6f, 0xa1, 0x2b, 0xe4,
	0xe9, 0x0b, 0xd8, 0x64, 0xa2, 0x96, 0xa9, 0xea, 0x5b, 0xf0, 0xe6, 0xc6, 0x02, 0x5f, 0x89, 0xa3,
	0x76, 0x7e, 0x25, 0xdf, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x47, 0x9e, 0x9e, 0xc6, 0x5b, 0x06,
	0x00, 0x00,
}
