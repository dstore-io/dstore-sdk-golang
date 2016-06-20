// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_LogoutCommunityMember_Ad.proto
// DO NOT EDIT!

/*
Package co_LogoutCommunityMember_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_LogoutCommunityMember_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_LogoutCommunityMember_Ad

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
	CommunityId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull       bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityMemberId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull bool                        `protobuf:"varint,1002,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
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

func (m *Parameters) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_LogoutCommunityMember_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_LogoutCommunityMember_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_LogoutCommunityMember_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x4a, 0xe3, 0x40,
	0x14, 0xa6, 0xcd, 0xf6, 0x87, 0xe9, 0x42, 0x77, 0xa7, 0xec, 0x92, 0x4d, 0x59, 0x91, 0x7a, 0xa3,
	0x08, 0x53, 0x51, 0x41, 0xf1, 0x42, 0x50, 0xf1, 0x22, 0x68, 0x8b, 0x8c, 0x20, 0xe8, 0x4d, 0x48,
	0x9b, 0x63, 0x08, 0x34, 0x33, 0x65, 0x26, 0xb1, 0x78, 0xe7, 0x23, 0xf8, 0x3e, 0x3e, 0x91, 0x3e,
	0x85, 0x93, 0x4c, 0x6a, 0x9a, 0x54, 0x54, 0xbc, 0x69, 0x73, 0xe6, 0x7c, 0xdf, 0xf9, 0xce, 0xf9,
	0xce, 0x0c, 0x3a, 0xf0, 0x64, 0xc4, 0x05, 0xf4, 0x81, 0xf9, 0x01, 0x83, 0xfe, 0x54, 0xf0, 0x31,
	0x78, 0xb1, 0x00, 0xd9, 0x1f, 0x73, 0xe7, 0x9c, 0xfb, 0x3c, 0x8e, 0x4e, 0x78, 0x18, 0xc6, 0x2c,
	0x88, 0xee, 0x07, 0x10, 0x8e, 0x40, 0x38, 0x47, 0x1e, 0x51, 0xb0, 0x88, 0xe3, 0x0d, 0xcd, 0x25,
	0x9a, 0x4b, 0x3e, 0x20, 0x58, 0x9d, 0x4c, 0xe6, 0xce, 0x9d, 0xc4, 0x20, 0x35, 0xdf, 0xfa, 0x57,
	0xd4, 0x06, 0x21, 0xb8, 0xc8, 0x52, 0xdd, 0x62, 0x2a, 0x04, 0x29, 0x5d, 0x1f, 0xb2, 0xe4, 0x5a,
	0x39, 0x19, 0xb9, 0x01, 0xbb, 0xe5, 0x22, 0x74, 0xa3, 0x80, 0x33, 0x0d, 0xea, 0x3d, 0x54, 0x11,
	0xba, 0x70, 0x85, 0xab, 0xb2, 0x20, 0x24, 0x3e, 0x44, 0x3f, 0xc7, 0xf3, 0xb6, 0x9c, 0xc0, 0x33,
	0x2b, 0xab, 0x95, 0xf5, 0xd6, 0x76, 0x97, 0x64, 0x23, 0x64, 0x7d, 0x05, 0x2c, 0x02, 0x1f, 0xc4,
	0x55, 0x12, 0xd1, 0xd6, 0x1b, 0xc1, 0xf6, 0xf0, 0x26, 0xfa, 0xbd, 0xc8, 0x77, 0x58, 0x3c, 0x99,
	0x98, 0xcf, 0x0d, 0x55, 0xa5, 0x49, 0xdb, 0x0b, 0xc0, 0xa1, 0x3a, 0xc7, 0x67, 0xa8, 0x93, 0x83,
	0x43, 0x6d, 0x82, 0xd2, 0xac, 0x7e, 0xae, 0x99, 0x8b, 0x68, 0xef, 0x94, 0xf2, 0x3e, 0x32, 0xdf,
	0x29, 0xa6, 0x1b, 0x78, 0xd1, 0x0d, 0xfc, 0x59, 0x62, 0x25, 0x6d, 0xf4, 0x9e, 0xaa, 0xa8, 0x49,
	0x41, 0x4e, 0x39, 0x93, 0x80, 0xb7, 0x50, 0x2d, 0x35, 0x38, 0x9b, 0xdc, 0x22, 0xc5, 0xe5, 0x69,
	0xf3, 0x4f, 0x93, 0x5f, 0xaa, 0x81, 0xf8, 0x1a, 0xfd, 0x4a, 0xac, 0x75, 0x16, 0xbc, 0x55, 0x23,
	0x18, 0x8a, 0x4c, 0x4a, 0xe4, 0xf2, 0x06, 0x06, 0x2a, 0xb6, 0xf3, 0x98, 0xb6, 0xc3, 0xe2, 0x81,
	0x9a, 0xa9, 0x91, 0xad, 0xd4, 0x34, 0xd2, 0x8a, 0x2b, 0x4b, 0x15, 0xf5, 0xc2, 0x07, 0xfa, 0x9f,
	0xce, 0xe1, 0xd8, 0x46, 0x86, 0xe0, 0x33, 0xf3, 0x47, 0xca, 0xda, 0x23, 0x5f, 0xbe, 0x81, 0x64,
	0x6e, 0x04, 0xa1, 0x7c, 0x46, 0x93, 0x1a, 0xd6, 0x7f, 0x64, 0xa8, 0x6f, 0xfc, 0x17, 0xd5, 0x55,
	0x94, 0xec, 0xe7, 0x71, 0xa8, 0xac, 0xa9, 0xd1, 0x9a, 0x0a, 0x6d, 0xef, 0xf8, 0x12, 0x75, 0x03,
	0x5e, 0x12, 0xc8, 0x9f, 0xc7, 0xcd, 0xee, 0x77, 0x1e, 0xce, 0xa8, 0x9e, 0x5e, 0xce, 0x9d, 0xd7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x27, 0x7c, 0x8d, 0x77, 0x03, 0x00, 0x00,
}