// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetForumCategories_Ad.proto
// DO NOT EDIT!

/*
Package fo_GetForumCategories_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetForumCategories_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetForumCategories_Ad

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
	PredecessorForumCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=predecessor_forum_category_id,json=predecessorForumCategoryId" json:"predecessor_forum_category_id,omitempty"`
	PredecessorForumCategoryIdNull bool                        `protobuf:"varint,1001,opt,name=predecessor_forum_category_id_null,json=predecessorForumCategoryIdNull" json:"predecessor_forum_category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPredecessorForumCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredecessorForumCategoryId
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
	RowId                      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ForumCategoryId            *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=forum_category_id,json=forumCategoryId" json:"forum_category_id,omitempty"`
	CategoryDescription        *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	SortNo                     *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	ContainsCategoriesOrForums *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=contains_categories_or_forums,json=containsCategoriesOrForums" json:"contains_categories_or_forums,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetForumCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumCategoryId
	}
	return nil
}

func (m *Response_Row) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetContainsCategoriesOrForums() *dstore_values.IntegerValue {
	if m != nil {
		return m.ContainsCategoriesOrForums
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetForumCategories_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetForumCategories_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetForumCategories_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xd6, 0xd6, 0xb5, 0x9d, 0xcc, 0xc5, 0xc0, 0x43, 0x28, 0xa4, 0xda, 0x34, 0x15, 0x09, 0x71,
	0xe5, 0xf2, 0x2b, 0xb8, 0xe5, 0x6f, 0x30, 0xa1, 0x85, 0xc9, 0x17, 0x48, 0x70, 0x81, 0x15, 0x92,
	0xd3, 0xca, 0x52, 0xeb, 0x53, 0x1d, 0xa7, 0x4c, 0x7b, 0x0b, 0xfe, 0x5e, 0x86, 0xa7, 0xe0, 0x39,
	0x10, 0x2f, 0x81, 0x1d, 0x27, 0xb4, 0xc9, 0x34, 0x28, 0x37, 0x89, 0x4e, 0xce, 0xf7, 0x7d, 0x3e,
	0xfe, 0x3e, 0xc7, 0xec, 0x61, 0x6e, 0x0b, 0x24, 0x18, 0x81, 0x99, 0x68, 0x03, 0xa3, 0x39, 0x61,
	0x06, 0xf9, 0x82, 0xc0, 0x8e, 0xc6, 0xa8, 0x5e, 0x40, 0x71, 0x88, 0xb4, 0x98, 0x3d, 0x4d, 0x0b,
	0x98, 0x20, 0x69, 0xb0, 0xea, 0x71, 0x2e, 0x1c, 0xa6, 0x40, 0x7e, 0x33, 0x10, 0x45, 0x20, 0x8a,
	0x8b, 0xd0, 0xf1, 0x6e, 0xb5, 0xc0, 0xc7, 0x74, 0xba, 0x00, 0x1b, 0xc8, 0xf1, 0xf5, 0xe6, 0xaa,
	0x40, 0x84, 0x54, 0xb5, 0x06, 0xcd, 0xd6, 0x0c, 0xac, 0x4d, 0x27, 0x50, 0x35, 0x6f, 0xb4, 0x9b,
	0x45, 0xaa, 0xcd, 0x18, 0x69, 0x96, 0x16, 0x1a, 0x4d, 0x00, 0x0d, 0xbf, 0x6f, 0x30, 0x76, 0x92,
	0x52, 0xea, 0xba, 0x40, 0x96, 0xbf, 0x67, 0x7b, 0x73, 0x82, 0x1c, 0x32, 0xa7, 0x84, 0xa4, 0xc6,
	0x7e, 0x44, 0x95, 0x85, 0x19, 0xcf, 0x94, 0xce, 0xa3, 0x8d, 0x83, 0x8d, 0x5b, 0x97, 0xee, 0x0e,
	0x44, 0xb5, 0xa1, 0x6a, 0x50, 0x6d, 0x1c, 0x02, 0xe8, 0x8d, 0xaf, 0x64, 0xbc, 0xa2, 0xb0, 0xba,
	0xc7, 0xb3, 0xa3, 0x9c, 0xbf, 0x62, 0xc3, 0xbf, 0xea, 0x2b, 0xb3, 0x98, 0x4e, 0xa3, 0x9f, 0x7d,
	0xb7, 0xca, 0xb6, 0xdc, 0xbf, 0x58, 0x28, 0x71, 0xb0, 0xe1, 0xaf, 0x2d, 0xb6, 0x2d, 0xc1, 0xce,
	0xd1, 0x58, 0xe0, 0xb7, 0x59, 0xb7, 0x74, 0xa6, 0x9a, 0x30, 0x16, 0x4d, 0xcb, 0x83, 0x6b, 0xcf,
	0xfd, 0x53, 0x06, 0x20, 0x7f, 0xcb, 0x2e, 0x7b, 0x4f, 0xd4, 0x8a, 0x29, 0xd1, 0xe6, 0x41, 0xc7,
	0x91, 0x45, 0x8b, 0xdc, 0xb6, 0xee, 0xd8, 0xd5, 0x47, 0xcb, 0x5a, 0xee, 0xcc, 0x9a, 0x1f, 0xf8,
	0x23, 0xd6, 0xaf, 0xb2, 0x88, 0x3a, 0xa5, 0xe2, 0xfe, 0x39, 0xc5, 0x90, 0xd4, 0x71, 0x78, 0xcb,
	0x1a, 0xce, 0x0f, 0x59, 0x87, 0xf0, 0x34, 0xda, 0x2a, 0x59, 0xf7, 0xc5, 0x7a, 0xe7, 0x46, 0xd4,
	0x2e, 0x08, 0x89, 0xa7, 0xd2, 0x0b, 0xc4, 0x3f, 0x36, 0x59, 0xc7, 0x15, 0xfc, 0x1a, 0xeb, 0xb9,
	0xd2, 0x27, 0xf7, 0x29, 0x71, 0xc6, 0x74, 0x65, 0xd7, 0x95, 0x2e, 0x88, 0x97, 0xec, 0xca, 0xf9,
	0x70, 0x3f, 0x27, 0xff, 0x4e, 0x77, 0x67, 0xdc, 0x8a, 0x34, 0x61, 0x57, 0xff, 0x68, 0xe4, 0x60,
	0x33, 0xd2, 0xf3, 0xd2, 0xca, 0x2f, 0x49, 0x33, 0x88, 0x4a, 0xcc, 0x16, 0xa4, 0xcd, 0x24, 0x68,
	0xed, 0xd6, 0xc4, 0x67, 0x4b, 0x1e, 0x7f, 0xc0, 0xfa, 0x2e, 0xef, 0x42, 0x19, 0x8c, 0xbe, 0xae,
	0x31, 0x4f, 0xcf, 0x83, 0x13, 0xe4, 0x8a, 0xed, 0x65, 0x68, 0x7c, 0x4e, 0xb6, 0xde, 0x93, 0x37,
	0xa7, 0x3e, 0x64, 0x36, 0xfa, 0xb6, 0x86, 0x58, 0x5c, 0x4b, 0x2c, 0xed, 0x7d, 0x1d, 0x8e, 0x9e,
	0x7d, 0x72, 0xc2, 0x06, 0x1a, 0x5b, 0x81, 0x2c, 0x6f, 0x80, 0x77, 0x77, 0xfe, 0xfb, 0x6e, 0xf8,
	0xd0, 0x2b, 0x7f, 0xc1, 0x7b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xba, 0xb4, 0x4e, 0x71, 0x57,
	0x04, 0x00, 0x00,
}
