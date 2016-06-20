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

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xed, 0x6a, 0x14, 0x31,
	0x14, 0xa5, 0x1d, 0xb7, 0x2d, 0x57, 0xf1, 0x23, 0xa5, 0x32, 0xbb, 0x4b, 0x45, 0xd6, 0x3f, 0xa2,
	0x30, 0x2b, 0x0a, 0xa2, 0xa0, 0x82, 0x82, 0x94, 0x55, 0x3b, 0x4a, 0x68, 0x05, 0x3f, 0x60, 0x88,
	0xcd, 0x75, 0x09, 0xcc, 0x24, 0x4b, 0x92, 0xb5, 0xf8, 0x16, 0x7e, 0x3e, 0x80, 0xbf, 0x7d, 0x15,
	0x5f, 0x44, 0x9f, 0xc2, 0xcc, 0x64, 0xea, 0x74, 0xb3, 0xd3, 0x2e, 0xf5, 0xcf, 0x2e, 0x99, 0x7b,
	0xcf, 0xb9, 0x27, 0xf7, 0x1c, 0x02, 0xf7, 0xb8, 0xb1, 0x4a, 0xe3, 0x10, 0xe5, 0x58, 0x48, 0x1c,
	0x4e, 0xb4, 0xda, 0x43, 0x3e, 0xd5, 0x68, 0x86, 0xa2, 0xc8, 0xb6, 0xd0, 0xee, 0xa4, 0x8a, 0xe3,
	0x36, 0x5a, 0x36, 0x92, 0xef, 0x95, 0x2e, 0x98, 0x15, 0x4a, 0x66, 0x0f, 0x79, 0xe2, 0x1a, 0xad,
	0x22, 0xd7, 0x3d, 0x3a, 0xf1, 0xe8, 0xe4, 0x58, 0x48, 0x6f, 0xbd, 0x1e, 0xf5, 0x81, 0xe5, 0x53,
	0x34, 0x9e, 0xa1, 0xd7, 0x9d, 0x9d, 0x8f, 0x5a, 0x2b, 0x5d, 0x97, 0xfa, 0xb3, 0xa5, 0x02, 0x8d,
	0x61, 0x63, 0xac, 0x8b, 0x57, 0xc2, 0xa2, 0x65, 0xa2, 0x19, 0xe7, 0x9b, 0x06, 0x3f, 0x96, 0x01,
	0x5e, 0x30, 0xcd, 0x5c, 0x15, 0xb5, 0x21, 0xf7, 0xe1, 0x8c, 0xd5, 0x88, 0x99, 0x74, 0xfa, 0x32,
	0xc1, 0xe3, 0xa5, 0xcb, 0x4b, 0x57, 0x4f, 0xdf, 0xec, 0x27, 0xf5, 0x25, 0x6a, 0x5d, 0x42, 0x5a,
	0x1c, 0xa3, 0x7e, 0x59, 0x9e, 0x28, 0x94, 0x80, 0xf2, 0x3e, 0x23, 0x4e, 0xae, 0xc1, 0x85, 0xc3,
	0xf0, 0x4c, 0x4e, 0xf3, 0x3c, 0xfe, 0xbd, 0xea, 0x48, 0xd6, 0xe8, 0xd9, 0xa6, 0x2f, 0x75, 0x9f,
	0xc9, 0x5b, 0xd8, 0x2c, 0x25, 0x65, 0x87, 0x34, 0x65, 0xf6, 0xe3, 0xa4, 0xc2, 0xe5, 0xc2, 0xd8,
	0x78, 0xb9, 0x9a, 0xdd, 0x0b, 0x66, 0x1b, 0xab, 0x85, 0x1c, 0xfb, 0xd1, 0xdd, 0x62, 0x76, 0x85,
	0x3b, 0x0e, 0x3e, 0xe2, 0xcf, 0x1c, 0x98, 0x3c, 0x81, 0xc1, 0xb1, 0xec, 0x5e, 0xda, 0x1f, 0x2f,
	0x6d, 0xf3, 0x48, 0x9e, 0x52, 0xe9, 0xe0, 0x57, 0x07, 0xd6, 0x28, 0x9a, 0x89, 0x92, 0x06, 0xc9,
	0x0d, 0xe8, 0x54, 0x0e, 0xd4, 0xab, 0xf9, 0x27, 0xaf, 0xf6, 0xd7, 0xbb, 0xf3, 0xb8, 0xfc, 0xa5,
	0xbe, 0x91, 0xbc, 0x82, 0xf3, 0xa1, 0x14, 0x77, 0xb7, 0xc8, 0x81, 0x93, 0x00, 0x1c, 0x5a, 0x14,
	0x24, 0x84, 0x9e, 0x0b, 0x74, 0x92, 0x3b, 0xb0, 0x5a, 0x7b, 0x1e, 0x47, 0x15, 0xe3, 0xa5, 0x39,
	0x46, 0x9f, 0x88, 0x6d, 0xff, 0x4f, 0x0f, 0xda, 0xc9, 0x53, 0x88, 0xb4, 0xda, 0x8f, 0x4f, 0x55,
	0xa8, 0xbb, 0xc9, 0x09, 0x42, 0x9a, 0x1c, 0xac, 0x22, 0xa1, 0x6a, 0x9f, 0x96, 0x2c, 0xbd, 0x9f,
	0x11, 0x44, 0xee, 0x40, 0x2e, 0xc2, 0x8a, 0x3b, 0x96, 0xb9, 0xf9, 0x94, 0xba, 0xed, 0x74, 0x68,
	0xc7, 0x1d, 0x5d, 0x2c, 0x76, 0x21, 0x3e, 0xca, 0x8c, 0xf8, 0x73, 0xba, 0x38, 0x62, 0x1b, 0xad,
	0xfe, 0x90, 0xe7, 0xb0, 0xd1, 0x4a, 0x1b, 0x7f, 0x49, 0x17, 0x46, 0x67, 0xbd, 0x85, 0x92, 0x3c,
	0x08, 0xd2, 0xff, 0x35, 0x3d, 0x59, 0xfc, 0xb7, 0x5a, 0x9c, 0xfe, 0xd6, 0xce, 0xc1, 0x71, 0x4f,
	0x14, 0x2c, 0xf7, 0x1c, 0x73, 0xbe, 0xbe, 0x81, 0xbe, 0xd5, 0x4c, 0x9a, 0x9c, 0x59, 0xe4, 0xd9,
	0x1c, 0xe7, 0xf7, 0xc5, 0xf7, 0xeb, 0x36, 0xf8, 0xc0, 0xc2, 0x47, 0xbb, 0xd0, 0x17, 0x2a, 0x70,
	0xbc, 0x79, 0xd4, 0x5e, 0xdf, 0xfe, 0xbf, 0xe7, 0xee, 0xdd, 0x4a, 0xf5, 0xa0, 0xdc, 0xfa, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x97, 0x02, 0xb6, 0x2f, 0x05, 0x00, 0x00,
}