// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_ModifyCommunities_Ad.proto
// DO NOT EDIT!

/*
Package co_ModifyCommunities_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_ModifyCommunities_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_ModifyCommunities_Ad

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
	CommunityId         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull     bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityName       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	CommunityNameNull   bool                        `protobuf:"varint,1002,opt,name=community_name_null,json=communityNameNull" json:"community_name_null,omitempty"`
	DeleteCommunity     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete_community,json=deleteCommunity" json:"delete_community,omitempty"`
	DeleteCommunityNull bool                        `protobuf:"varint,1003,opt,name=delete_community_null,json=deleteCommunityNull" json:"delete_community_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
	}
	return nil
}

func (m *Parameters) GetDeleteCommunity() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCommunity
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_ModifyCommunities_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_ModifyCommunities_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_ModifyCommunities_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x8b, 0xd3, 0x40,
	0x10, 0xe7, 0x1a, 0x7b, 0x77, 0xcc, 0xa9, 0xbd, 0xdb, 0xa2, 0xc4, 0x14, 0x45, 0x4e, 0x04, 0x41,
	0xd8, 0x1c, 0x16, 0xc4, 0x27, 0xe1, 0x94, 0x13, 0xfa, 0xd0, 0x2a, 0xfb, 0x20, 0xe8, 0x4b, 0xc8,
	0x75, 0xe7, 0xc2, 0x42, 0xb2, 0x7b, 0xec, 0x26, 0x1e, 0xfd, 0x16, 0x7e, 0x1b, 0x3f, 0x93, 0xf5,
	0x4b, 0xb8, 0xc9, 0x6e, 0xd3, 0x26, 0x2a, 0x78, 0x2f, 0x6d, 0x66, 0xe7, 0xf7, 0x67, 0x76, 0x66,
	0x16, 0x5e, 0x73, 0x53, 0x2a, 0x8d, 0x31, 0xca, 0x4c, 0x48, 0x8c, 0xaf, 0xb5, 0x5a, 0x22, 0xaf,
	0x34, 0x9a, 0x78, 0xa9, 0x92, 0xb9, 0xe2, 0xe2, 0x6a, 0xf5, 0x5e, 0x15, 0x45, 0x25, 0x45, 0x29,
	0xd0, 0x24, 0xe7, 0x9c, 0x5a, 0x48, 0xa9, 0xc8, 0x73, 0xc7, 0xa3, 0x8e, 0x47, 0xff, 0x01, 0x8e,
	0xc6, 0x5e, 0xfe, 0x5b, 0x9a, 0x57, 0x68, 0x1c, 0x37, 0x7a, 0xd4, 0xf5, 0x44, 0xad, 0x95, 0xf6,
	0xa9, 0x49, 0x37, 0x55, 0xa0, 0x31, 0x69, 0x86, 0x3e, 0xf9, 0xac, 0x9f, 0x2c, 0x53, 0x21, 0xaf,
	0x94, 0x2e, 0xd2, 0x52, 0x28, 0xe9, 0x40, 0xa7, 0xeb, 0x01, 0xc0, 0xa7, 0x54, 0xa7, 0x36, 0x8b,
	0xda, 0x90, 0xb7, 0x70, 0x77, 0xe9, 0x4b, 0x5a, 0x25, 0x82, 0x87, 0x7b, 0x4f, 0xf7, 0x5e, 0x1c,
	0xbd, 0x9a, 0x50, 0x5f, 0xbe, 0xaf, 0x4b, 0xc8, 0x12, 0x33, 0xd4, 0x9f, 0xeb, 0x88, 0x1d, 0xb5,
	0x84, 0x19, 0x27, 0x2f, 0xe1, 0x64, 0x97, 0x9f, 0xc8, 0x2a, 0xcf, 0xc3, 0x9f, 0x07, 0x56, 0xe5,
	0x90, 0x8d, 0x76, 0x80, 0x0b, 0x7b, 0x4e, 0xce, 0xe1, 0xfe, 0x16, 0x2c, 0x6d, 0x09, 0xe1, 0xa0,
	0xb1, 0x8b, 0x7a, 0x76, 0xa6, 0xd4, 0x42, 0x66, 0xce, 0xed, 0x5e, 0xcb, 0x58, 0x58, 0x02, 0x89,
	0x61, 0xdc, 0x95, 0x70, 0x8e, 0x6b, 0xe7, 0x78, 0xd2, 0x01, 0x37, 0x9e, 0x1f, 0xe0, 0x98, 0x63,
	0x6e, 0x2f, 0x9b, 0xb4, 0xb9, 0x30, 0xf8, 0xeb, 0x25, 0x2f, 0x95, 0xca, 0x31, 0x95, 0xce, 0x76,
	0xe4, 0x48, 0x9b, 0x71, 0xad, 0xc8, 0x14, 0x1e, 0xf4, 0x75, 0x9c, 0xf5, 0x2f, 0x67, 0x3d, 0xee,
	0x11, 0x6a, 0xf3, 0xd3, 0x1f, 0x03, 0x38, 0x64, 0x68, 0xae, 0x95, 0x34, 0x48, 0xce, 0x60, 0xd8,
	0x8c, 0xd2, 0xf7, 0xb8, 0xbd, 0xb4, 0x5f, 0x11, 0x37, 0xe6, 0x8b, 0xfa, 0x97, 0x39, 0x20, 0xf9,
	0x02, 0xc7, 0xf5, 0x10, 0x93, 0x9d, 0x29, 0xda, 0x8e, 0x05, 0x96, 0x4c, 0x7b, 0xe4, 0xfe, 0xac,
	0xe7, 0x36, 0x9e, 0x6d, 0x63, 0x36, 0x2a, 0xba, 0x07, 0xe4, 0x0d, 0x1c, 0xf8, 0xe5, 0xb1, 0xdd,
	0xa8, 0x15, 0x9f, 0xfc, 0xa1, 0xe8, 0x56, 0x6b, 0xee, 0xfe, 0xd9, 0x06, 0x4e, 0x2e, 0x20, 0xd0,
	0xea, 0x26, 0xbc, 0xd3, 0xb0, 0xa6, 0xf4, 0xbf, 0xf6, 0x9c, 0x6e, 0x9a, 0x40, 0x99, 0xba, 0x61,
	0x35, 0x3f, 0x7a, 0x0c, 0x81, 0xfd, 0x26, 0x0f, 0x61, 0xdf, 0x46, 0xf5, 0xe6, 0x7d, 0x5f, 0xd8,
	0xb6, 0x0c, 0xd9, 0xd0, 0x86, 0x33, 0xfe, 0xee, 0x23, 0x4c, 0x84, 0xea, 0x89, 0x6f, 0x1f, 0xdf,
	0xd7, 0xb3, 0xdb, 0x3e, 0xcb, 0xcb, 0xfd, 0x66, 0xfd, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0xac, 0xc1, 0xb1, 0xd1, 0x03, 0x00, 0x00,
}
