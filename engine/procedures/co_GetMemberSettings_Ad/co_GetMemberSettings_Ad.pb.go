// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetMemberSettings_Ad.proto
// DO NOT EDIT!

/*
Package co_GetMemberSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetMemberSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetMemberSettings_Ad

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
	CommunityId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull       bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityMemberId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull bool                        `protobuf:"varint,1002,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
	KeyVariable           *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull       bool                        `protobuf:"varint,1003,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
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

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CommunityId       *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityMemberId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	Value             *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value" json:"value,omitempty"`
	CommunityName     *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	KeyVariable       *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Response_Row) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
	}
	return nil
}

func (m *Response_Row) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetMemberSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetMemberSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetMemberSettings_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_GetMemberSettings_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x6d, 0x6b, 0x13, 0x41,
	0x10, 0xa6, 0x49, 0xd3, 0x96, 0x6d, 0xb5, 0xba, 0x45, 0x39, 0x13, 0x10, 0xa9, 0x08, 0x82, 0x70,
	0x51, 0x0b, 0xa2, 0x1f, 0x54, 0x2c, 0x14, 0x29, 0x9a, 0x43, 0x56, 0x28, 0xea, 0x97, 0xe3, 0x92,
	0x1b, 0x8f, 0xa5, 0xb9, 0xdd, 0xb2, 0xbb, 0x69, 0xe9, 0xbf, 0xf0, 0xf5, 0x17, 0xfa, 0xc9, 0x97,
	0xdf, 0x20, 0xce, 0xde, 0x6e, 0x72, 0x2f, 0xad, 0xc4, 0xf8, 0x25, 0x61, 0x6e, 0x9e, 0x67, 0x9e,
	0xd9, 0x99, 0x67, 0x97, 0x3c, 0x48, 0xb5, 0x91, 0x0a, 0xfa, 0x20, 0x32, 0x2e, 0xa0, 0x7f, 0xa4,
	0xe4, 0x08, 0xd2, 0x89, 0x02, 0xdd, 0x1f, 0xc9, 0xf8, 0x39, 0x98, 0x01, 0xe4, 0x43, 0x50, 0xaf,
	0xc1, 0x18, 0x2e, 0x32, 0x1d, 0x3f, 0x4b, 0x43, 0x84, 0x18, 0x49, 0x6f, 0x39, 0x5e, 0xe8, 0x78,
	0xe1, 0x5f, 0xc0, 0xdd, 0x2d, 0x5f, 0xfe, 0x38, 0x19, 0x4f, 0x40, 0x3b, 0x6e, 0xf7, 0x5a, 0x5d,
	0x13, 0x94, 0x92, 0xca, 0xa7, 0x7a, 0xf5, 0x54, 0x0e, 0x5a, 0x27, 0x19, 0xf8, 0xe4, 0xcd, 0x66,
	0xd2, 0x24, 0x5c, 0xbc, 0x97, 0x2a, 0x4f, 0x0c, 0x97, 0xc2, 0x81, 0xb6, 0x7f, 0xb5, 0x08, 0x79,
	0x95, 0xa8, 0x04, 0xb3, 0xa0, 0x34, 0x7d, 0x42, 0x36, 0x46, 0x32, 0xcf, 0x27, 0x82, 0x9b, 0xd3,
	0x98, 0xa7, 0xc1, 0xd2, 0x8d, 0xa5, 0xdb, 0xeb, 0xf7, 0x7b, 0xa1, 0x6f, 0xdf, 0xf7, 0xc5, 0x85,
	0x81, 0x0c, 0xd4, 0x81, 0x8d, 0xd8, 0xfa, 0x8c, 0xb0, 0x9f, 0xd2, 0x3b, 0xe4, 0x72, 0x95, 0x1f,
	0x8b, 0xc9, 0x78, 0x1c, 0x7c, 0x5f, 0xc5, 0x2a, 0x6b, 0x6c, 0xb3, 0x02, 0x8c, 0xf0, 0x3b, 0x7d,
	0x41, 0xb6, 0x4a, 0x70, 0x5e, 0x0c, 0xc3, 0x6a, 0xb6, 0xe6, 0x6b, 0x96, 0x22, 0x6e, 0x86, 0xa8,
	0xfc, 0x90, 0x04, 0xe7, 0x14, 0x73, 0x0d, 0xfc, 0x70, 0x0d, 0x5c, 0x39, 0xc3, 0x2a, 0xda, 0x78,
	0x4c, 0x36, 0x0e, 0xe1, 0x34, 0x3e, 0x4e, 0x14, 0x4f, 0x86, 0x63, 0x08, 0xda, 0x85, 0x7e, 0xb7,
	0xa1, 0xaf, 0x8d, 0xc2, 0x2d, 0xf9, 0x23, 0x23, 0xfe, 0xc0, 0xc3, 0xed, 0x91, 0xab, 0x74, 0xa7,
	0xf8, 0xd3, 0x1f, 0xb9, 0x02, 0xb4, 0x5a, 0xdb, 0xbf, 0x97, 0xc9, 0x1a, 0x03, 0x7d, 0x24, 0x85,
	0x06, 0x7a, 0x97, 0x74, 0x8a, 0x65, 0xfa, 0x29, 0xcf, 0x14, 0xbd, 0x49, 0xdc, 0xa2, 0xf7, 0xec,
	0x2f, 0x73, 0x40, 0xfa, 0x96, 0x5c, 0xb2, 0x6b, 0x8c, 0x2b, 0x7b, 0xc4, 0x71, 0xb5, 0x91, 0x1c,
	0x36, 0xc8, 0xcd, 0x6d, 0x0f, 0x30, 0xde, 0x2f, 0x63, 0xb6, 0x99, 0xd7, 0x3f, 0xe0, 0xfc, 0x56,
	0xbd, 0x7d, 0x70, 0x00, 0xb6, 0xe2, 0xf5, 0x33, 0x15, 0x9d, 0xb9, 0x06, 0xee, 0x9f, 0x4d, 0xe1,
	0x74, 0x8f, 0xb4, 0x95, 0x3c, 0x09, 0x96, 0x0b, 0xd6, 0x4e, 0xf8, 0x4f, 0x4e, 0x0f, 0xa7, 0x43,
	0x08, 0x99, 0x3c, 0x61, 0x96, 0xdf, 0xfd, 0xd6, 0x22, 0x6d, 0x0c, 0xe8, 0x55, 0xb2, 0x82, 0xa1,
	0x35, 0xc2, 0x87, 0x08, 0xe7, 0xd2, 0x61, 0x1d, 0x0c, 0x71, 0xc1, 0x4f, 0x1b, 0xd6, 0xfc, 0x18,
	0x2d, 0xe8, 0xcd, 0x97, 0xe7, 0xdb, 0xed, 0x53, 0xf4, 0x5f, 0x7e, 0xbb, 0x47, 0x3a, 0x05, 0x32,
	0xf8, 0x1c, 0xcd, 0xf5, 0x8b, 0x43, 0xd2, 0x5d, 0x72, 0xb1, 0x6c, 0x40, 0xe0, 0x95, 0x0b, 0xbe,
	0xcc, 0xe7, 0x5e, 0x98, 0x51, 0x22, 0x64, 0xd8, 0x0b, 0x5a, 0x33, 0xeb, 0xd7, 0x68, 0x21, 0xb7,
	0xee, 0xbe, 0x21, 0x3d, 0x2e, 0x1b, 0x3b, 0x2a, 0x5f, 0xb1, 0x77, 0x8f, 0x32, 0xa9, 0xd3, 0xc3,
	0x69, 0x3e, 0x5d, 0xe0, 0xa1, 0x1b, 0xae, 0x14, 0x0f, 0xca, 0xce, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0xe0, 0xd5, 0x94, 0x23, 0x05, 0x00, 0x00,
}
