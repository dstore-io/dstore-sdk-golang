// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetPageCategories_Ad.proto
// DO NOT EDIT!

/*
Package st_GetPageCategories_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetPageCategories_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetPageCategories_Ad

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
	PageCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=page_category_id,json=pageCategoryId" json:"page_category_id,omitempty"`
	PageCategoryIdNull bool                        `protobuf:"varint,1001,opt,name=page_category_id_null,json=pageCategoryIdNull" json:"page_category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPageCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageCategoryId
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
	RowId                   int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PageCategoryId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=page_category_id,json=pageCategoryId" json:"page_category_id,omitempty"`
	PageCategoryDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=page_category_description,json=pageCategoryDescription" json:"page_category_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPageCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PageCategoryId
	}
	return nil
}

func (m *Response_Row) GetPageCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PageCategoryDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetPageCategories_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetPageCategories_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetPageCategories_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetPageCategories_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdb, 0x0a, 0xd3, 0x30,
	0x18, 0xa6, 0xd6, 0x1d, 0x88, 0xa0, 0x23, 0xa2, 0x76, 0x1d, 0xc8, 0x98, 0x08, 0x5e, 0x65, 0xb2,
	0x81, 0xe8, 0xa5, 0x87, 0x29, 0xbb, 0x58, 0x19, 0xb9, 0xf0, 0x74, 0x53, 0xea, 0xf2, 0x5b, 0x82,
	0x5b, 0x52, 0x92, 0xcc, 0xb1, 0x27, 0xf0, 0x56, 0x7d, 0x20, 0xdf, 0xc3, 0x47, 0xf0, 0x2d, 0x4c,
	0x9b, 0x8c, 0xae, 0x55, 0xc1, 0xdd, 0xb4, 0xfc, 0xf9, 0x0e, 0xc9, 0xff, 0xfd, 0x09, 0x7a, 0xc4,
	0xb4, 0x91, 0x0a, 0xa6, 0x20, 0x72, 0x2e, 0x60, 0x5a, 0x28, 0xb9, 0x01, 0xb6, 0x57, 0xa0, 0xa7,
	0xda, 0xa4, 0xaf, 0xc0, 0xac, 0xb3, 0x1c, 0x9e, 0x67, 0x06, 0x72, 0xa9, 0x38, 0xe8, 0xf4, 0x29,
	0x23, 0x96, 0x62, 0x24, 0xbe, 0xef, 0x74, 0xc4, 0xe9, 0xc8, 0x3f, 0xc8, 0xf1, 0x4d, 0x6f, 0xff,
	0x39, 0xdb, 0xee, 0x41, 0x3b, 0x6d, 0x3c, 0x6c, 0xee, 0x09, 0x4a, 0x49, 0xe5, 0xa1, 0x51, 0x13,
	0xda, 0x81, 0xd6, 0xd6, 0xd3, 0x83, 0xf7, 0xda, 0xa0, 0xc9, 0xb8, 0xf8, 0x28, 0xd5, 0x2e, 0x33,
	0x5c, 0x0a, 0x47, 0x9a, 0x7c, 0x09, 0x10, 0x5a, 0x67, 0x2a, 0xb3, 0x28, 0x28, 0x8d, 0x17, 0x68,
	0x50, 0x58, 0x87, 0x74, 0xe3, 0x8e, 0x75, 0x4c, 0x39, 0x8b, 0x82, 0x71, 0xf0, 0xe0, 0xda, 0x6c,
	0x44, 0x7c, 0x0b, 0xfe, 0x6c, 0x5c, 0x58, 0x06, 0xa8, 0xd7, 0x65, 0x45, 0xaf, 0x17, 0x75, 0x2b,
	0xc7, 0x25, 0xc3, 0x33, 0x74, 0xab, 0x6d, 0x93, 0x8a, 0xfd, 0x76, 0x1b, 0xfd, 0xea, 0x59, 0xb3,
	0x3e, 0xc5, 0x4d, 0x7e, 0x62, 0xa1, 0xc9, 0xcf, 0x10, 0xf5, 0x29, 0xe8, 0x42, 0x0a, 0x0d, 0xf8,
	0x21, 0xea, 0x54, 0x7d, 0xfa, 0xcd, 0x63, 0xd2, 0xcc, 0xcf, 0x65, 0xb0, 0x28, 0xbf, 0xd4, 0x11,
	0xf1, 0x3b, 0x34, 0x28, 0x3b, 0x4c, 0xcf, 0x5a, 0x8c, 0xae, 0x8c, 0x43, 0x2b, 0x26, 0x2d, 0x71,
	0x3b, 0x88, 0x95, 0xad, 0x97, 0x75, 0x4d, 0x6f, 0xec, 0x9a, 0x0b, 0xf8, 0x31, 0xea, 0xf9, 0x64,
	0xa3, 0xb0, 0x72, 0xbc, 0xfb, 0x87, 0xa3, 0xcb, 0x7d, 0xe5, 0xfe, 0xf4, 0x44, 0xb7, 0x71, 0x86,
	0x4a, 0x1e, 0xa2, 0xab, 0x95, 0x6a, 0x4e, 0xfe, 0xeb, 0x12, 0x90, 0x53, 0x08, 0x84, 0xca, 0x03,
	0x2d, 0xf5, 0xf1, 0x8f, 0x00, 0x85, 0xb6, 0xc0, 0xb7, 0x51, 0xd7, 0x96, 0xe5, 0x4c, 0xbe, 0x26,
	0x36, 0x97, 0x0e, 0xed, 0xd8, 0xd2, 0xc6, 0xfd, 0xf2, 0x2f, 0x53, 0xfb, 0x96, 0x5c, 0x3e, 0xb6,
	0x37, 0x68, 0xd8, 0xf4, 0x61, 0xa0, 0x37, 0x8a, 0x17, 0x55, 0x98, 0xdf, 0x93, 0xe6, 0x28, 0xbc,
	0xa1, 0x36, 0x8a, 0x8b, 0xdc, 0xf9, 0xdd, 0x39, 0xf7, 0x7b, 0x51, 0x6b, 0x9f, 0xbd, 0x45, 0x23,
	0x2e, 0x5b, 0xed, 0xd7, 0x6f, 0xe7, 0xfd, 0x93, 0x5c, 0x6a, 0xf6, 0xe9, 0x84, 0xb3, 0x0b, 0x9e,
	0xd7, 0x87, 0x6e, 0x75, 0x8d, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x59, 0x99, 0xeb,
	0x99, 0x03, 0x00, 0x00,
}
