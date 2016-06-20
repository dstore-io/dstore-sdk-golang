// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetBinaryCatAccessLevels.proto
// DO NOT EDIT!

/*
Package co_GetBinaryCatAccessLevels is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetBinaryCatAccessLevels.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetBinaryCatAccessLevels

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
	CommunityBinaryCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_binary_category_id,json=communityBinaryCategoryId" json:"community_binary_category_id,omitempty"`
	CommunityBinaryCategoryIdNull bool                        `protobuf:"varint,1001,opt,name=community_binary_category_id_null,json=communityBinaryCategoryIdNull" json:"community_binary_category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityBinaryCategoryId
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
	RowId         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	AccessLevelId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=access_level_id,json=accessLevelId" json:"access_level_id,omitempty"`
	Description   *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=description" json:"description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetAccessLevelId() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevelId
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetBinaryCatAccessLevels.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetBinaryCatAccessLevels.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetBinaryCatAccessLevels.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_GetBinaryCatAccessLevels.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x6a, 0xd5, 0x30,
	0x18, 0xa7, 0xd6, 0xb3, 0x8d, 0x0c, 0x99, 0x54, 0x90, 0xae, 0x47, 0x65, 0xce, 0x1b, 0xbd, 0xc9,
	0x11, 0xbd, 0x50, 0x84, 0x5d, 0x38, 0x15, 0x39, 0xe0, 0x8a, 0xe4, 0x42, 0x50, 0x06, 0x25, 0x6b,
	0x3f, 0x4b, 0xb0, 0x4d, 0x0e, 0x49, 0xba, 0xb1, 0xb7, 0x50, 0x1f, 0xc1, 0x77, 0xf0, 0x01, 0x7c,
	0x1c, 0xdf, 0xc2, 0xaf, 0x49, 0xce, 0xce, 0x5a, 0xf1, 0xa8, 0x37, 0x6d, 0xbf, 0xfe, 0xfe, 0x7c,
	0xc9, 0xef, 0x4b, 0xc8, 0xb3, 0xca, 0x58, 0xa5, 0x61, 0x06, 0xb2, 0x16, 0x12, 0x66, 0x0b, 0xad,
	0x4a, 0xa8, 0x3a, 0x0d, 0x66, 0x56, 0xaa, 0xe2, 0x35, 0xd8, 0x43, 0x21, 0xb9, 0x3e, 0x7f, 0xc1,
	0xed, 0xf3, 0xb2, 0x04, 0x63, 0xde, 0xc0, 0x29, 0x34, 0x86, 0x22, 0xcd, 0xaa, 0xe4, 0x81, 0xd7,
	0x52, 0xaf, 0xa5, 0x6b, 0x04, 0xd9, 0x8d, 0xd0, 0xe6, 0x94, 0x37, 0x1d, 0x04, 0x7d, 0xb6, 0x3b,
	0xec, 0x0d, 0x5a, 0x2b, 0x1d, 0xa0, 0xe9, 0x10, 0x6a, 0xd1, 0x89, 0xd7, 0x10, 0xc0, 0x7b, 0x63,
	0xd0, 0x72, 0x21, 0x3f, 0x2a, 0xdd, 0x72, 0x2b, 0x94, 0xf4, 0xa4, 0xfd, 0xef, 0x11, 0x21, 0x6f,
	0xb9, 0xe6, 0x88, 0x82, 0x36, 0xc9, 0x31, 0xb9, 0x55, 0xaa, 0xb6, 0xed, 0xa4, 0xb0, 0xe7, 0xc5,
	0x89, 0x5b, 0x63, 0x51, 0x72, 0x0b, 0xb5, 0xc2, 0x0f, 0x51, 0xa5, 0xd1, 0x5e, 0x74, 0x7f, 0xfb,
	0xd1, 0x94, 0x86, 0x2d, 0x85, 0x75, 0x0a, 0x89, 0x0c, 0xd0, 0xef, 0xfa, 0x8a, 0xed, 0x5e, 0x18,
	0x5c, 0xec, 0xd1, 0xc9, 0xe7, 0x55, 0x32, 0x27, 0x77, 0xd7, 0xb9, 0x17, 0xb2, 0x6b, 0x9a, 0xf4,
	0xe7, 0x26, 0xf6, 0xd8, 0x62, 0xb7, 0xff, 0x68, 0x93, 0x23, 0x6b, 0xff, 0x47, 0x4c, 0xb6, 0x18,
	0x98, 0x85, 0x92, 0x06, 0x92, 0x87, 0x64, 0xe2, 0x52, 0x09, 0xcb, 0xcb, 0xe8, 0x30, 0x71, 0x9f,
	0xd8, 0xab, 0xfe, 0xc9, 0x3c, 0x31, 0x79, 0x4f, 0xae, 0xf7, 0x79, 0x14, 0x97, 0x02, 0x49, 0xaf,
	0xec, 0xc5, 0x28, 0xa6, 0x23, 0xf1, 0x38, 0xb6, 0x23, 0xac, 0xe7, 0xab, 0x9a, 0xed, 0xb4, 0xc3,
	0x1f, 0xc9, 0x53, 0xb2, 0x19, 0xe6, 0x90, 0xc6, 0xce, 0xf1, 0xce, 0x6f, 0x8e, 0x7e, 0x4a, 0x47,
	0xfe, 0xcd, 0x96, 0x74, 0x8c, 0x27, 0xd6, 0xea, 0x2c, 0xbd, 0xea, 0x54, 0x4f, 0xe8, 0x3f, 0x1f,
	0x1b, 0xba, 0x0c, 0x82, 0x32, 0x75, 0xc6, 0x7a, 0x8f, 0xec, 0x5b, 0x44, 0x62, 0x2c, 0x92, 0x9b,
	0x64, 0x03, 0xcb, 0x7e, 0x72, 0x9f, 0x73, 0xcc, 0x66, 0xc2, 0x26, 0x58, 0xe2, 0x24, 0x5e, 0x92,
	0x1d, 0xee, 0x1c, 0x8a, 0xa6, 0xb7, 0xe8, 0x09, 0x5f, 0xf2, 0xbf, 0xcf, 0xf6, 0x1a, 0x5f, 0xb5,
	0x45, 0x97, 0x03, 0xb2, 0x5d, 0x81, 0x29, 0xb5, 0x58, 0xb8, 0x00, 0xbf, 0xe6, 0xc3, 0xf8, 0x83,
	0x83, 0xb1, 0x5a, 0xc8, 0xda, 0x1b, 0x5c, 0xe6, 0x1f, 0x1e, 0x93, 0xa9, 0x50, 0xa3, 0x6d, 0xae,
	0x6e, 0xd6, 0x87, 0x83, 0x5a, 0x99, 0xea, 0xd3, 0x12, 0xaf, 0xfe, 0xf3, 0xf2, 0x9d, 0x6c, 0xb8,
	0x03, 0xfe, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x22, 0x29, 0x14, 0xbb, 0x03, 0x00,
	0x00,
}