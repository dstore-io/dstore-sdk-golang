// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunityMemberSettings.proto
// DO NOT EDIT!

/*
Package co_GetCommunityMemberSettings is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunityMemberSettings.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunityMemberSettings

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
	CommunityMemberId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull bool                        `protobuf:"varint,1001,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
	CommunityId           *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull       bool                        `protobuf:"varint,1002,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	KeyVariable           *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull       bool                        `protobuf:"varint,1003,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	SelectResult          *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=select_result,json=selectResult" json:"select_result,omitempty"`
	SelectResultNull      bool                        `protobuf:"varint,1004,opt,name=select_result_null,json=selectResultNull" json:"select_result_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Parameters) GetSelectResult() *dstore_values.BooleanValue {
	if m != nil {
		return m.SelectResult
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Value           *dstore_values.StringValue                       `protobuf:"bytes,101,opt,name=value" json:"value,omitempty"`
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

func (m *Response) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type Response_Row struct {
	RowId       int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value       *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=value" json:"value,omitempty"`
	KeyVariable *dstore_values.StringValue `protobuf:"bytes,20002,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunityMemberSettings.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunityMemberSettings.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunityMemberSettings.Response.Row")
}

var fileDescriptor0 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x86, 0x59, 0xc7, 0x6d, 0x4b, 0xba, 0xd2, 0x36, 0x45, 0x19, 0x77, 0x41, 0xa4, 0xde, 0x08,
	0xc5, 0x59, 0x3f, 0x40, 0x14, 0xfc, 0x42, 0x11, 0x29, 0xb2, 0x8b, 0x44, 0x2c, 0xe8, 0xcd, 0x90,
	0x9d, 0x3d, 0x2e, 0xc1, 0x99, 0xa4, 0x24, 0x99, 0x96, 0xfe, 0x0b, 0x7b, 0xef, 0x95, 0xff, 0xcb,
	0x3b, 0x6f, 0xfc, 0xf8, 0x11, 0x9e, 0x4c, 0xb2, 0xdd, 0x99, 0x69, 0xb1, 0xd6, 0x9b, 0x5d, 0xb2,
	0xe7, 0xbc, 0xef, 0x93, 0x9c, 0xbc, 0x1b, 0xf2, 0x68, 0x6a, 0xac, 0xd2, 0x30, 0x04, 0x39, 0x13,
	0x12, 0x86, 0x7b, 0x5a, 0x65, 0x30, 0x2d, 0x35, 0x98, 0x61, 0xa6, 0xd2, 0x57, 0x60, 0x5f, 0xa8,
	0xa2, 0x28, 0xa5, 0xb0, 0x87, 0x23, 0x28, 0x26, 0xa0, 0xdf, 0x82, 0xb5, 0x42, 0xce, 0x4c, 0x82,
	0x8d, 0x56, 0xd1, 0x6d, 0xaf, 0x4e, 0xbc, 0x3a, 0xf9, 0xab, 0xa4, 0xbf, 0x19, 0x50, 0xfb, 0x3c,
	0x2f, 0x21, 0x38, 0xf4, 0xaf, 0x36, 0xf9, 0xa0, 0xb5, 0xd2, 0xa1, 0x34, 0x68, 0x96, 0x0a, 0x30,
	0x86, 0xcf, 0x20, 0x14, 0x6f, 0xb4, 0x8b, 0x96, 0x0b, 0xf9, 0x51, 0xe9, 0x82, 0x5b, 0xa1, 0xa4,
	0x6f, 0xda, 0xfa, 0x16, 0x11, 0xf2, 0x86, 0x6b, 0x8e, 0x55, 0xd0, 0x86, 0xbe, 0x26, 0x9b, 0xd9,
	0x7c, 0x6f, 0x69, 0x51, 0x6d, 0x2e, 0x15, 0xd3, 0xb8, 0x73, 0xbd, 0x73, 0x73, 0xf5, 0xee, 0x20,
	0x09, 0x67, 0x09, 0xdb, 0x13, 0xd2, 0xc2, 0x0c, 0xf4, 0xae, 0x5b, 0xb1, 0x8d, 0xac, 0x79, 0xa6,
	0x9d, 0x29, 0x7d, 0x40, 0xe2, 0x53, 0xcc, 0x52, 0x59, 0xe6, 0x79, 0xfc, 0x63, 0x19, 0x2d, 0x57,
	0xd8, 0xe5, 0x13, 0xaa, 0x31, 0x56, 0xe9, 0x13, 0xd2, 0x5b, 0x28, 0x91, 0x7f, 0xe1, 0x6c, 0xfe,
	0xea, 0xb1, 0x00, 0xc9, 0xdb, 0x64, 0xa3, 0xae, 0xf7, 0xc8, 0x9f, 0x1e, 0xb9, 0x56, 0x6b, 0xac,
	0x60, 0x8f, 0x49, 0xef, 0x13, 0x1c, 0xa6, 0xfb, 0x5c, 0x0b, 0x3e, 0xc9, 0x21, 0x8e, 0x2a, 0x58,
	0xbf, 0x05, 0x33, 0x56, 0xe3, 0x15, 0x05, 0x16, 0xf6, 0xef, 0x86, 0x76, 0xc7, 0xaa, 0xcb, 0x3d,
	0xeb, 0x57, 0x60, 0xd5, 0x1a, 0x2b, 0xd6, 0x33, 0x72, 0xc9, 0x40, 0x0e, 0x99, 0x4d, 0x31, 0x3f,
	0x65, 0x6e, 0xe3, 0x8b, 0xa7, 0x9e, 0x6c, 0xa2, 0x54, 0x0e, 0x5c, 0x7a, 0x5a, 0xcf, 0x2b, 0x58,
	0x25, 0xa0, 0xb7, 0x08, 0x6d, 0x38, 0x78, 0xde, 0x6f, 0xcf, 0x5b, 0xaf, 0xb7, 0x3a, 0xe0, 0xd6,
	0xf7, 0x88, 0xac, 0xe0, 0x72, 0x4f, 0x49, 0x03, 0xf4, 0x36, 0xe9, 0x56, 0xe9, 0x09, 0xf7, 0x79,
	0x7c, 0xc4, 0x90, 0x4d, 0x9f, 0xac, 0x97, 0xee, 0x93, 0xf9, 0x46, 0xfa, 0x9e, 0xac, 0xbb, 0xdc,
	0xa4, 0xb5, 0xe0, 0xe0, 0x65, 0x44, 0x28, 0x4e, 0x5a, 0xe2, 0x76, 0xbc, 0x46, 0xb8, 0xde, 0x59,
	0xac, 0xd9, 0x5a, 0xd1, 0xfc, 0x01, 0xd3, 0xb1, 0x1c, 0xf2, 0x8a, 0x13, 0x77, 0x8e, 0xd7, 0x4e,
	0x38, 0xfa, 0x34, 0x8f, 0xfc, 0x37, 0x9b, 0xb7, 0x63, 0x48, 0x23, 0xad, 0x0e, 0x70, 0x74, 0x4e,
	0xf5, 0x30, 0x39, 0xc7, 0x1f, 0x2c, 0x99, 0x8f, 0x22, 0x61, 0xea, 0x80, 0x39, 0x17, 0x37, 0x93,
	0x6a, 0xe8, 0x31, 0x9c, 0x79, 0xed, 0xbe, 0xb1, 0x7f, 0xd4, 0x21, 0x11, 0xca, 0xe9, 0x15, 0xb2,
	0x84, 0x06, 0x2e, 0x9e, 0x9f, 0xc7, 0xa8, 0xed, 0xb2, 0x2e, 0x2e, 0x31, 0x7c, 0x77, 0xe6, 0x8e,
	0x47, 0xe3, 0x7f, 0xb4, 0xa4, 0x4f, 0x5b, 0x11, 0xfc, 0xfa, 0xa5, 0x73, 0xae, 0x10, 0x3e, 0x7f,
	0x47, 0x06, 0x42, 0xb5, 0x26, 0xb1, 0x78, 0xa8, 0x3e, 0xdc, 0xff, 0xbf, 0x27, 0x6c, 0xb2, 0x54,
	0x3d, 0x12, 0xf7, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xf7, 0x54, 0x24, 0x03, 0x05, 0x00,
	0x00,
}
