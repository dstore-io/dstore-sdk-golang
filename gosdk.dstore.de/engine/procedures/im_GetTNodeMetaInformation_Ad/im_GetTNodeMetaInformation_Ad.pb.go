// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetTNodeMetaInformation_Ad.proto
// DO NOT EDIT!

/*
Package im_GetTNodeMetaInformation_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetTNodeMetaInformation_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetTNodeMetaInformation_Ad

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
	TreeNodeId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull                bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	MetaInformationTypeIdList     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=meta_information_type_id_list,json=metaInformationTypeIdList" json:"meta_information_type_id_list,omitempty"`
	MetaInformationTypeIdListNull bool                        `protobuf:"varint,1002,opt,name=meta_information_type_id_list_null,json=metaInformationTypeIdListNull" json:"meta_information_type_id_list_null,omitempty"`
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

func (m *Parameters) GetMetaInformationTypeIdList() *dstore_values.StringValue {
	if m != nil {
		return m.MetaInformationTypeIdList
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	MetaInformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=meta_information_type_id,json=metaInformationTypeId" json:"meta_information_type_id,omitempty"`
	MetaInformationType       *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=meta_information_type,json=metaInformationType" json:"meta_information_type,omitempty"`
	TreeNodeId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	MetaInformation           *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	TranslatedMetaInformation *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=translated_meta_information,json=translatedMetaInformation" json:"translated_meta_information,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetMetaInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.MetaInformationTypeId
	}
	return nil
}

func (m *Response_Row) GetMetaInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.MetaInformationType
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetMetaInformation() *dstore_values.DecimalValue {
	if m != nil {
		return m.MetaInformation
	}
	return nil
}

func (m *Response_Row) GetTranslatedMetaInformation() *dstore_values.StringValue {
	if m != nil {
		return m.TranslatedMetaInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetTNodeMetaInformation_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetTNodeMetaInformation_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetTNodeMetaInformation_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetTNodeMetaInformation_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xed, 0x6a, 0x13, 0x41,
	0x14, 0xa5, 0x8d, 0x69, 0xcb, 0x55, 0xfc, 0x98, 0x52, 0xd9, 0x24, 0x54, 0x24, 0xfe, 0x11, 0x85,
	0x8d, 0xe8, 0x1f, 0x05, 0x3f, 0x50, 0x90, 0x12, 0xb5, 0xab, 0x0c, 0x55, 0xf0, 0x03, 0x97, 0xb1,
	0x73, 0x5d, 0x06, 0x77, 0x67, 0xc2, 0xcc, 0xc4, 0xe2, 0x5b, 0xf8, 0xf9, 0x00, 0xfe, 0xf6, 0x55,
	0x7c, 0x11, 0x7d, 0x0a, 0x67, 0x77, 0x36, 0xa6, 0x99, 0x6c, 0x1b, 0xf2, 0x27, 0x61, 0xf6, 0xde,
	0x73, 0xee, 0x99, 0x7b, 0x0e, 0x03, 0xb7, 0xb9, 0xb1, 0x4a, 0xe3, 0x00, 0x65, 0x26, 0x24, 0x0e,
	0x46, 0x5a, 0xed, 0x23, 0x1f, 0x6b, 0x34, 0x03, 0x51, 0xa4, 0x3b, 0x68, 0xf7, 0x12, 0xc5, 0x71,
	0x17, 0x2d, 0x1b, 0xca, 0xf7, 0x4a, 0x17, 0xcc, 0x0a, 0x25, 0xd3, 0xfb, 0x3c, 0x76, 0x8d, 0x56,
	0x91, 0xab, 0x1e, 0x1d, 0x7b, 0x74, 0x7c, 0x2c, 0xa4, 0xbb, 0x59, 0x8f, 0xfa, 0xc8, 0xf2, 0x31,
	0x1a, 0xcf, 0xd0, 0xed, 0xcc, 0xce, 0x47, 0xad, 0x95, 0xae, 0x4b, 0xbd, 0xd9, 0x52, 0x81, 0xc6,
	0xb0, 0x0c, 0xeb, 0xe2, 0xa5, 0xb0, 0x68, 0x99, 0x98, 0x8e, 0xf3, 0x4d, 0xfd, 0x9f, 0xab, 0x00,
	0xcf, 0x98, 0x66, 0xae, 0x8a, 0xda, 0x90, 0x3b, 0x70, 0xca, 0x6a, 0xc4, 0x54, 0x3a, 0x7d, 0xa9,
	0xe0, 0xd1, 0xca, 0xc5, 0x95, 0xcb, 0x27, 0xaf, 0xf7, 0xe2, 0xfa, 0x12, 0xb5, 0x2e, 0x21, 0x2d,
	0x66, 0xa8, 0x5f, 0x94, 0x27, 0x0a, 0x25, 0xa0, 0xbc, 0xcf, 0x90, 0x93, 0x2b, 0x70, 0xee, 0x30,
	0x3c, 0x95, 0xe3, 0x3c, 0x8f, 0xfe, 0xac, 0x3b, 0x92, 0x0d, 0x7a, 0x7a, 0xda, 0x97, 0xb8, 0xcf,
	0xe4, 0x0d, 0x6c, 0x97, 0x92, 0xd2, 0x43, 0x9a, 0x52, 0xfb, 0x69, 0x54, 0xe1, 0x72, 0x61, 0x6c,
	0xb4, 0x5a, 0xcd, 0xee, 0x06, 0xb3, 0x8d, 0xd5, 0x42, 0x66, 0x7e, 0x74, 0xa7, 0x98, 0x5d, 0xe1,
	0x9e, 0x83, 0x0f, 0xf9, 0x13, 0x07, 0x26, 0x8f, 0xa0, 0x7f, 0x2c, 0xbb, 0x97, 0xf6, 0xd7, 0x4b,
	0xdb, 0x3e, 0x92, 0xa7, 0x54, 0xda, 0xff, 0xdd, 0x86, 0x0d, 0x8a, 0x66, 0xa4, 0xa4, 0x41, 0x72,
	0x0d, 0xda, 0x95, 0x03, 0xf5, 0x6a, 0xfe, 0xcb, 0xab, 0xfd, 0xf5, 0xee, 0x3c, 0x2c, 0x7f, 0xa9,
	0x6f, 0x24, 0x2f, 0xe1, 0x6c, 0x28, 0xc5, 0xdd, 0xad, 0xe5, 0xc0, 0x71, 0x00, 0x0e, 0x2d, 0x0a,
	0x12, 0x42, 0xcf, 0x04, 0x3a, 0xc9, 0x4d, 0x58, 0xaf, 0x3d, 0x8f, 0x5a, 0x15, 0xe3, 0x85, 0x39,
	0x46, 0x9f, 0x88, 0x5d, 0xff, 0x4f, 0x27, 0xed, 0xe4, 0x31, 0xb4, 0xb4, 0x3a, 0x88, 0x4e, 0x54,
	0xa8, 0x5b, 0xf1, 0x12, 0x21, 0x8d, 0x27, 0xab, 0x88, 0xa9, 0x3a, 0xa0, 0x25, 0x4b, 0xf7, 0x57,
	0x0b, 0x5a, 0xee, 0x40, 0xce, 0xc3, 0x9a, 0x3b, 0x96, 0xb9, 0xf9, 0x9c, 0xb8, 0xed, 0xb4, 0x69,
	0xdb, 0x1d, 0x5d, 0x2c, 0x9e, 0x43, 0x74, 0x94, 0x19, 0xd1, 0x97, 0x64, 0x71, 0xc4, 0xb6, 0x1a,
	0xfd, 0x21, 0x4f, 0x61, 0xab, 0x91, 0x36, 0xfa, 0x9a, 0x2c, 0x8c, 0xce, 0x66, 0x03, 0x25, 0xb9,
	0x1b, 0xa4, 0xff, 0x5b, 0xb2, 0x5c, 0xfc, 0x77, 0x1a, 0x9c, 0xfe, 0xde, 0xcc, 0xc1, 0x71, 0x5f,
	0x14, 0x2c, 0xf7, 0x1c, 0x73, 0xbe, 0xbe, 0x86, 0x9e, 0xd5, 0x4c, 0x9a, 0x9c, 0x59, 0xe4, 0xe9,
	0x1c, 0xe7, 0x8f, 0xc5, 0xf7, 0xeb, 0x4c, 0xf1, 0x81, 0x85, 0x0f, 0xde, 0x42, 0x4f, 0xa8, 0xc0,
	0xf1, 0xe9, 0xa3, 0xf6, 0xea, 0x5e, 0xa6, 0x0c, 0xff, 0x30, 0xa9, 0xf3, 0xa5, 0xdf, 0xbd, 0x77,
	0x6b, 0xd5, 0xcb, 0x72, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x55, 0xc8, 0x96, 0x38,
	0x05, 0x00, 0x00,
}