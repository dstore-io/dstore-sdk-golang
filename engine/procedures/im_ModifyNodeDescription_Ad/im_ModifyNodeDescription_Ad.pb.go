// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyNodeDescription_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyNodeDescription_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyNodeDescription_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyNodeDescription_Ad

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
	NodeId                        *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	NodeIdNull                    bool                        `protobuf:"varint,1001,opt,name=node_id_null,json=nodeIdNull" json:"node_id_null,omitempty"`
	LanguageId                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	NodeDescription               *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	NodeDescriptionNull           bool                        `protobuf:"varint,1003,opt,name=node_description_null,json=nodeDescriptionNull" json:"node_description_null,omitempty"`
	DeleteNodeDescription         *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_node_description,json=deleteNodeDescription" json:"delete_node_description,omitempty"`
	DeleteNodeDescriptionNull     bool                        `protobuf:"varint,1004,opt,name=delete_node_description_null,json=deleteNodeDescriptionNull" json:"delete_node_description_null,omitempty"`
	UpdateProductDescriptions     *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=update_product_descriptions,json=updateProductDescriptions" json:"update_product_descriptions,omitempty"`
	UpdateProductDescriptionsNull bool                        `protobuf:"varint,1005,opt,name=update_product_descriptions_null,json=updateProductDescriptionsNull" json:"update_product_descriptions_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Parameters) GetDeleteNodeDescription() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteNodeDescription
	}
	return nil
}

func (m *Parameters) GetUpdateProductDescriptions() *dstore_values.BooleanValue {
	if m != nil {
		return m.UpdateProductDescriptions
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyNodeDescription_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyNodeDescription_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyNodeDescription_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_ModifyNodeDescription_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6b, 0x13, 0x41,
	0x10, 0xa7, 0x8d, 0x49, 0xca, 0x54, 0x6c, 0xd8, 0x52, 0xcd, 0x87, 0x95, 0x5a, 0x5f, 0xec, 0xcb,
	0x45, 0xac, 0xa0, 0x88, 0x82, 0x8a, 0x05, 0xf3, 0x90, 0x50, 0x4e, 0x10, 0xfc, 0x80, 0xe3, 0x9a,
	0x9d, 0x1e, 0x8b, 0x97, 0xdd, 0xb0, 0x7b, 0x67, 0xf1, 0x3f, 0xf0, 0xd1, 0xff, 0xc7, 0xbf, 0xc8,
	0xaf, 0xff, 0xc1, 0xd9, 0x9b, 0x8b, 0x49, 0xae, 0x5f, 0xfa, 0x92, 0x64, 0x76, 0xe6, 0xf7, 0x91,
	0xdd, 0x99, 0x81, 0xc7, 0xd2, 0x65, 0xc6, 0x62, 0x1f, 0x75, 0xa2, 0x34, 0xf6, 0xa7, 0xd6, 0x8c,
	0x51, 0xe6, 0x16, 0x5d, 0x5f, 0x4d, 0xa2, 0xa1, 0x91, 0xea, 0xf8, 0xf3, 0xc8, 0x48, 0x7c, 0x89,
	0x6e, 0x6c, 0xd5, 0x34, 0x53, 0x46, 0x47, 0xcf, 0x65, 0x40, 0x65, 0x99, 0x11, 0x7b, 0x8c, 0x0d,
	0x18, 0x1b, 0x5c, 0x00, 0xe8, 0x6e, 0x96, 0x32, 0x9f, 0xe2, 0x34, 0x47, 0xc7, 0xf8, 0x6e, 0x67,
	0x59, 0x1b, 0xad, 0x35, 0xb6, 0x4c, 0xf5, 0x96, 0x53, 0x13, 0x74, 0x2e, 0x4e, 0xb0, 0x4c, 0xde,
	0xa9, 0x26, 0xb3, 0x58, 0xe9, 0x63, 0x63, 0x27, 0xb1, 0x97, 0xe3, 0xa2, 0xdd, 0x2f, 0x75, 0x80,
	0xc3, 0xd8, 0xc6, 0x94, 0x45, 0xeb, 0xc4, 0x03, 0x68, 0x6a, 0xb2, 0x15, 0x29, 0xd9, 0x5e, 0xd9,
	0x59, 0xb9, 0xbb, 0x7e, 0xbf, 0x17, 0x94, 0xee, 0x4b, 0x4b, 0x4a, 0x67, 0x98, 0xa0, 0x7d, 0xe3,
	0xa3, 0xb0, 0xe1, 0x6b, 0x07, 0x52, 0xdc, 0x86, 0xab, 0x25, 0x2a, 0xd2, 0x79, 0x9a, 0xb6, 0xbf,
	0x37, 0x09, 0xbb, 0x16, 0x02, 0xa7, 0x47, 0x74, 0x24, 0x9e, 0xc0, 0x7a, 0x1a, 0xeb, 0x24, 0x27,
	0x7b, 0x9e, 0x7c, 0xf5, 0x72, 0x72, 0x98, 0xd5, 0x93, 0xc0, 0x1e, 0xb4, 0x16, 0xd0, 0x2c, 0xf2,
	0x83, 0x45, 0xae, 0xcd, 0xcb, 0x0a, 0xa1, 0x03, 0x68, 0x15, 0x5e, 0xe4, 0xfc, 0x66, 0xdb, 0xb5,
	0x42, 0xad, 0x5b, 0x51, 0x73, 0x99, 0x55, 0x3a, 0x61, 0xb1, 0x0d, 0xbd, 0xfc, 0x18, 0x62, 0x1f,
	0xb6, 0xaa, 0x34, 0x2c, 0xfb, 0x93, 0x65, 0x37, 0x2b, 0x80, 0x42, 0xfb, 0x35, 0xdc, 0x90, 0x98,
	0xd2, 0x4d, 0x46, 0xa7, 0x2c, 0x5c, 0x39, 0xf3, 0x0f, 0x1f, 0x19, 0x93, 0x62, 0xac, 0xd9, 0xc3,
	0x16, 0x63, 0x2b, 0x6d, 0x21, 0x9e, 0xc1, 0xcd, 0x73, 0x48, 0xd9, 0xd0, 0x2f, 0x36, 0xd4, 0x39,
	0x13, 0x5d, 0xd8, 0x7a, 0x0f, 0xbd, 0x7c, 0x2a, 0x63, 0x62, 0xa0, 0x37, 0x97, 0xf9, 0x38, 0x5b,
	0x24, 0x71, 0xed, 0xfa, 0xe5, 0xd6, 0x3a, 0x8c, 0x3f, 0x64, 0xf8, 0x02, 0xbf, 0x13, 0xaf, 0x60,
	0xe7, 0x02, 0x72, 0xb6, 0xf8, 0x9b, 0x2d, 0x6e, 0x9f, 0xcb, 0xe2, 0x6d, 0xee, 0x7e, 0x5b, 0x85,
	0xb5, 0x10, 0xdd, 0x94, 0x62, 0x14, 0xf7, 0xa0, 0x5e, 0x34, 0x7a, 0xd9, 0x86, 0x7f, 0xdf, 0xae,
	0x1c, 0x22, 0x1e, 0x82, 0x03, 0xff, 0x19, 0x72, 0xa1, 0x78, 0x0b, 0x2d, 0xdf, 0xe2, 0xd1, 0x42,
	0x8f, 0x53, 0x9b, 0xd5, 0x08, 0x1c, 0x54, 0xc0, 0xd5, 0x49, 0x18, 0x52, 0x3c, 0x98, 0xc7, 0xe1,
	0xc6, 0x64, 0xf9, 0x40, 0x3c, 0x82, 0x66, 0x39, 0x5a, 0xd4, 0x4a, 0x9e, 0xf1, 0xd6, 0x29, 0x46,
	0x1e, 0xbc, 0x21, 0x7f, 0x87, 0xb3, 0x72, 0x31, 0x80, 0x9a, 0x35, 0x27, 0xf4, 0xfa, 0x1e, 0xf5,
	0x30, 0xf8, 0xe7, 0x4d, 0x10, 0xcc, 0x2e, 0x22, 0x08, 0xcd, 0x49, 0xe8, 0x39, 0xba, 0xdb, 0x50,
	0xa3, 0xdf, 0xe2, 0x3a, 0x34, 0x28, 0xf2, 0x33, 0xf4, 0x75, 0x44, 0x57, 0x53, 0x0f, 0xeb, 0x14,
	0x0e, 0xe4, 0x8b, 0x0f, 0xd0, 0x53, 0xa6, 0x22, 0x30, 0x5f, 0x53, 0xef, 0x9e, 0x26, 0xc6, 0xc9,
	0x8f, 0xb3, 0xbc, 0xfc, 0xcf, 0x4d, 0x76, 0xd4, 0x28, 0xb6, 0xc5, 0xfe, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x56, 0xc5, 0x82, 0xc0, 0x08, 0x05, 0x00, 0x00,
}
