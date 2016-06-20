// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyValuesInCategories_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyValuesInCategories_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyValuesInCategories_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyValuesInCategories_Ad

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
	ValueId             *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	ValueIdNull         bool                        `protobuf:"varint,1001,opt,name=value_id_null,json=valueIdNull" json:"value_id_null,omitempty"`
	ValueCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=value_category_id,json=valueCategoryId" json:"value_category_id,omitempty"`
	ValueCategoryIdNull bool                        `protobuf:"varint,1002,opt,name=value_category_id_null,json=valueCategoryIdNull" json:"value_category_id_null,omitempty"`
	Delete              *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete" json:"delete,omitempty"`
	DeleteNull          bool                        `protobuf:"varint,1003,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Parameters) GetValueCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueCategoryId
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyValuesInCategories_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyValuesInCategories_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyValuesInCategories_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_ModifyValuesInCategories_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x89, 0xf9, 0xe1, 0x04, 0xa9, 0x4e, 0xa1, 0xc4, 0x04, 0xa5, 0xb4, 0x37, 0x5e, 0xc8,
	0x44, 0xac, 0x88, 0x08, 0x82, 0x3f, 0x88, 0x04, 0x4c, 0x91, 0xb9, 0x10, 0xf4, 0x66, 0x99, 0x76,
	0x4e, 0x97, 0xc1, 0xdd, 0x99, 0x32, 0xb3, 0xb1, 0xe4, 0x2d, 0x7c, 0x22, 0x7d, 0x1e, 0xf5, 0x25,
	0x9c, 0x9d, 0x33, 0x4b, 0xdd, 0xad, 0x28, 0xb9, 0x49, 0xf6, 0xfc, 0x7c, 0xdf, 0x77, 0xfe, 0x06,
	0x9e, 0x2b, 0x5f, 0x59, 0x87, 0x0b, 0x34, 0xb9, 0x36, 0xb8, 0xb8, 0x70, 0xf6, 0x0c, 0xd5, 0xda,
	0xa1, 0x5f, 0xe8, 0x32, 0x5b, 0x59, 0xa5, 0xcf, 0x37, 0x1f, 0x64, 0xb1, 0x46, 0xbf, 0x34, 0xaf,
	0x65, 0x85, 0xb9, 0x75, 0x1a, 0x7d, 0xf6, 0x52, 0xf1, 0x90, 0x59, 0x59, 0xf6, 0x80, 0xe0, 0x9c,
	0xe0, 0xfc, 0xdf, 0x98, 0xd9, 0x5e, 0x12, 0xfb, 0x12, 0xa3, 0x44, 0x31, 0xbb, 0xd3, 0xae, 0x00,
	0x9d, 0xb3, 0x2e, 0x85, 0xe6, 0xed, 0x50, 0x89, 0xde, 0xcb, 0x1c, 0x53, 0xf0, 0xa8, 0x1b, 0xac,
	0xa4, 0x36, 0xe7, 0xd6, 0x95, 0xb2, 0xd2, 0xd6, 0x50, 0xd2, 0xe1, 0xf7, 0x1e, 0xc0, 0x7b, 0xe9,
	0x64, 0x88, 0xa2, 0xf3, 0xec, 0x09, 0x8c, 0xa3, 0x76, 0xa6, 0xd5, 0x74, 0xe7, 0x60, 0xe7, 0xfe,
	0xe4, 0xd1, 0x9c, 0xa7, 0x0e, 0x52, 0x4d, 0xda, 0x84, 0x82, 0xd1, 0xc5, 0xfa, 0xc5, 0x28, 0x3a,
	0x97, 0x8a, 0x1d, 0xc1, 0xcd, 0x06, 0x97, 0x99, 0x75, 0x51, 0x4c, 0x7f, 0x8c, 0x02, 0x7a, 0x2c,
	0x26, 0x29, 0xe1, 0x24, 0xf8, 0xd8, 0x5b, 0xb8, 0x4d, 0x49, 0x67, 0xd4, 0xf4, 0xa6, 0x56, 0xe9,
	0xfd, 0x5f, 0x65, 0x37, 0x3a, 0xd3, 0xa4, 0x36, 0x41, 0xed, 0x31, 0xec, 0x5f, 0x23, 0x22, 0xd9,
	0x9f, 0x24, 0xbb, 0xd7, 0x41, 0x44, 0xf9, 0x63, 0x18, 0x2a, 0x2c, 0x42, 0x9f, 0xd3, 0xfe, 0x5f,
	0x35, 0x4f, 0xad, 0x2d, 0x50, 0x1a, 0xd2, 0x4c, 0xa9, 0xec, 0x00, 0x26, 0xf4, 0x45, 0xfc, 0xbf,
	0x88, 0x1f, 0xc8, 0x57, 0xd3, 0x1e, 0x7e, 0xeb, 0xc1, 0x58, 0xa0, 0xbf, 0xb0, 0xc6, 0x23, 0x7b,
	0x08, 0x83, 0xb8, 0x9f, 0x34, 0xbc, 0x19, 0x6f, 0xaf, 0x9f, 0x76, 0xf7, 0xa6, 0xfe, 0x15, 0x94,
	0xc8, 0x3e, 0xc2, 0xad, 0x7a, 0x33, 0xd9, 0x1f, 0xab, 0x09, 0x33, 0xe9, 0x07, 0x30, 0xef, 0x80,
	0xbb, 0x0b, 0x5c, 0x05, 0x7b, 0x79, 0x65, 0x8b, 0xdd, 0xb2, 0xed, 0x60, 0x4f, 0x61, 0x94, 0x2e,
	0x22, 0x74, 0x5c, 0x33, 0xde, 0xbb, 0xc6, 0x48, 0xf7, 0xb2, 0xa2, 0x7f, 0xd1, 0xa4, 0xb3, 0x77,
	0xd0, 0x77, 0xf6, 0x72, 0x7a, 0x23, 0xa2, 0x9e, 0xf1, 0x6d, 0x6e, 0x98, 0x37, 0xb3, 0xe0, 0xc2,
	0x5e, 0x8a, 0x9a, 0x66, 0x76, 0x17, 0xfa, 0xe1, 0x9b, 0xed, 0xc3, 0x30, 0x58, 0xf5, 0xce, 0xbf,
	0x9e, 0x84, 0xe9, 0x0c, 0xc4, 0x20, 0x98, 0x4b, 0xf5, 0x2a, 0x83, 0xb9, 0xb6, 0x1d, 0x8d, 0xab,
	0x67, 0xf6, 0xe9, 0x45, 0x6e, 0xbd, 0xfa, 0xdc, 0xc4, 0xd5, 0xf6, 0x2f, 0xf1, 0x74, 0x18, 0x4f,
	0xfd, 0xf8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x19, 0x7c, 0xd6, 0xcb, 0x03, 0x00, 0x00,
}