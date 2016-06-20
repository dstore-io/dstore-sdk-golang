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

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x89, 0xf9, 0xe1, 0x04, 0xa9, 0x4e, 0xa1, 0xc4, 0x0d, 0x4a, 0x69, 0x6f, 0xbc, 0x90,
	0x89, 0x58, 0x51, 0x11, 0xbc, 0x50, 0x11, 0x59, 0x30, 0x45, 0xe6, 0xa2, 0xa0, 0x37, 0xcb, 0xb4,
	0x73, 0x1a, 0x06, 0x76, 0x67, 0xca, 0xcc, 0xc6, 0x92, 0xb7, 0xf0, 0x89, 0xf4, 0x79, 0xd4, 0x97,
	0x70, 0x76, 0xce, 0x84, 0xba, 0x1b, 0x51, 0xf4, 0x26, 0xd9, 0xf3, 0xf3, 0x7d, 0xdf, 0xf9, 0x1b,
	0x78, 0xa1, 0x7c, 0x6d, 0x1d, 0xce, 0xd1, 0x2c, 0xb5, 0xc1, 0xf9, 0xa5, 0xb3, 0xe7, 0xa8, 0x56,
	0x0e, 0xfd, 0x5c, 0x57, 0xc5, 0xc2, 0x2a, 0x7d, 0xb1, 0x3e, 0x95, 0xe5, 0x0a, 0x7d, 0x6e, 0x5e,
	0xcb, 0x1a, 0x97, 0xd6, 0x69, 0xf4, 0xc5, 0x4b, 0xc5, 0x43, 0x66, 0x6d, 0xd9, 0x03, 0x82, 0x73,
	0x82, 0xf3, 0x3f, 0x63, 0xb2, 0xbd, 0x24, 0xf6, 0x29, 0x46, 0x89, 0x22, 0xbb, 0xd3, 0xae, 0x00,
	0x9d, 0xb3, 0x2e, 0x85, 0x66, 0xed, 0x50, 0x85, 0xde, 0xcb, 0x25, 0xa6, 0xe0, 0x51, 0x37, 0x58,
	0x4b, 0x6d, 0x2e, 0xac, 0xab, 0x64, 0xad, 0xad, 0xa1, 0xa4, 0xc3, 0xaf, 0x3d, 0x80, 0xf7, 0xd2,
	0xc9, 0x10, 0x45, 0xe7, 0xd9, 0x13, 0x18, 0x47, 0xed, 0x42, 0xab, 0xe9, 0xce, 0xc1, 0xce, 0xfd,
	0xc9, 0xa3, 0x19, 0x4f, 0x1d, 0xa4, 0x9a, 0xb4, 0x09, 0x05, 0xa3, 0x8b, 0xf5, 0x8b, 0x51, 0x74,
	0xe6, 0x8a, 0x1d, 0xc1, 0xcd, 0x0d, 0xae, 0x30, 0xab, 0xb2, 0x9c, 0x7e, 0x1b, 0x05, 0xf4, 0x58,
	0x4c, 0x52, 0xc2, 0x49, 0xf0, 0xb1, 0xb7, 0x70, 0x9b, 0x92, 0xce, 0xa9, 0xe9, 0x75, 0xa3, 0xd2,
	0xfb, 0xbb, 0xca, 0x6e, 0x74, 0xa6, 0x49, 0xad, 0x83, 0xda, 0x63, 0xd8, 0xdf, 0x22, 0x22, 0xd9,
	0xef, 0x24, 0xbb, 0xd7, 0x41, 0x44, 0xf9, 0x63, 0x18, 0x2a, 0x2c, 0x43, 0x9f, 0xd3, 0xfe, 0x6f,
	0x35, 0xcf, 0xac, 0x2d, 0x51, 0x1a, 0xd2, 0x4c, 0xa9, 0xec, 0x00, 0x26, 0xf4, 0x45, 0xfc, 0x3f,
	0x88, 0x1f, 0xc8, 0xd7, 0xd0, 0x1e, 0x7e, 0xe9, 0xc1, 0x58, 0xa0, 0xbf, 0xb4, 0xc6, 0x23, 0x7b,
	0x08, 0x83, 0xb8, 0x9f, 0x34, 0xbc, 0x8c, 0xb7, 0xd7, 0x4f, 0xbb, 0x7b, 0xd3, 0xfc, 0x0a, 0x4a,
	0x64, 0x1f, 0xe0, 0x56, 0xb3, 0x99, 0xe2, 0x97, 0xd5, 0x84, 0x99, 0xf4, 0x03, 0x98, 0x77, 0xc0,
	0xdd, 0x05, 0x2e, 0x82, 0x9d, 0x5f, 0xdb, 0x62, 0xb7, 0x6a, 0x3b, 0xd8, 0x33, 0x18, 0xa5, 0x8b,
	0x08, 0x1d, 0x37, 0x8c, 0xf7, 0xb6, 0x18, 0xe9, 0x5e, 0x16, 0xf4, 0x2f, 0x36, 0xe9, 0xec, 0x1d,
	0xf4, 0x9d, 0xbd, 0x9a, 0xde, 0x88, 0xa8, 0xe7, 0xfc, 0x5f, 0x6e, 0x98, 0x6f, 0x66, 0xc1, 0x85,
	0xbd, 0x12, 0x0d, 0x4d, 0x76, 0x17, 0xfa, 0xe1, 0x9b, 0xed, 0xc3, 0x30, 0x58, 0xcd, 0xce, 0x3f,
	0x9f, 0x84, 0xe9, 0x0c, 0xc4, 0x20, 0x98, 0xb9, 0x7a, 0x75, 0x0a, 0x33, 0x6d, 0x3b, 0x1a, 0xd7,
	0xcf, 0xec, 0xe3, 0xd3, 0xff, 0x7c, 0x80, 0x67, 0xc3, 0x78, 0xe1, 0xc7, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0x3f, 0xc9, 0x6a, 0xc2, 0x03, 0x00, 0x00,
}
