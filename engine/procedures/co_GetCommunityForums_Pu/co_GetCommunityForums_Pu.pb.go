// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunityForums_Pu.proto
// DO NOT EDIT!

/*
Package co_GetCommunityForums_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunityForums_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunityForums_Pu

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
	CommunityId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1002,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1003,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1004,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
	}
	return nil
}

type Response struct {
	Error             *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation   []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message           []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row               []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	CommunityMemberId *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
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

func (m *Response) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

type Response_Row struct {
	RowId         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CommunityId   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	ForumId       *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumName     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=forum_name,json=forumName" json:"forum_name,omitempty"`
	CommunityName *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	LanguageId    *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	AccessLevel   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=access_level,json=accessLevel" json:"access_level,omitempty"`
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

func (m *Response_Row) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Response_Row) GetForumName() *dstore_values.StringValue {
	if m != nil {
		return m.ForumName
	}
	return nil
}

func (m *Response_Row) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Response_Row) GetAccessLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevel
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunityForums_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunityForums_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunityForums_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_GetCommunityForums_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0xed, 0x6a, 0xd4, 0x4e,
	0x14, 0xc6, 0xe9, 0x7f, 0xdb, 0xee, 0xf6, 0x6c, 0xff, 0xad, 0x4e, 0x45, 0xe3, 0x6e, 0x29, 0x5a,
	0x51, 0x04, 0x21, 0x2b, 0x2a, 0xac, 0x14, 0xb4, 0x50, 0xb1, 0xb2, 0xea, 0x2e, 0x35, 0x1f, 0x84,
	0x8a, 0x10, 0xd2, 0x64, 0x36, 0x04, 0x93, 0x99, 0x75, 0x26, 0x69, 0xf1, 0x9b, 0x97, 0xe0, 0x7b,
	0xaf, 0xd1, 0x97, 0x8b, 0xf0, 0xcc, 0x4c, 0xb2, 0x2f, 0xd1, 0x75, 0xb7, 0x5f, 0x5a, 0x26, 0xe7,
	0xfc, 0x9e, 0xf3, 0x64, 0xe6, 0x99, 0x2c, 0xb4, 0x03, 0x99, 0x72, 0x41, 0x5b, 0x94, 0x85, 0x11,
	0xa3, 0xad, 0x81, 0xe0, 0x3e, 0x0d, 0x32, 0x41, 0x65, 0xcb, 0xe7, 0xee, 0x13, 0x9a, 0x3e, 0xe2,
	0x49, 0x92, 0xb1, 0x28, 0x7d, 0xb7, 0xcf, 0x45, 0x96, 0x48, 0xf7, 0x20, 0xb3, 0xb1, 0x27, 0xe5,
	0xe4, 0x86, 0x01, 0x6d, 0x03, 0xda, 0xd3, 0xba, 0x1b, 0x1b, 0xf9, 0x80, 0x63, 0x2f, 0xce, 0xa8,
	0x34, 0x70, 0xe3, 0xf2, 0xe4, 0x54, 0x2a, 0x04, 0x17, 0x79, 0xa9, 0x39, 0x59, 0x4a, 0xa8, 0x94,
	0x5e, 0x48, 0xf3, 0xe2, 0xb5, 0x72, 0x31, 0xf5, 0x22, 0xd6, 0xe7, 0x22, 0xf1, 0xd2, 0x88, 0x33,
	0xd3, 0xb4, 0x7d, 0xba, 0x08, 0x70, 0xe0, 0x09, 0x0f, 0xab, 0x54, 0x48, 0xf2, 0x10, 0x56, 0xfd,
	0xc2, 0x96, 0x1b, 0x05, 0xd6, 0xc2, 0x95, 0x85, 0x9b, 0xf5, 0x3b, 0x4d, 0x3b, 0xf7, 0x9f, 0xfb,
	0x8a, 0x58, 0x4a, 0x43, 0x2a, 0x5e, 0xaa, 0x95, 0x53, 0x1f, 0x02, 0x9d, 0x80, 0xdc, 0x82, 0xf3,
	0xe3, 0xbc, 0xcb, 0xb2, 0x38, 0xb6, 0xbe, 0x57, 0x51, 0xa5, 0xe6, 0xac, 0x8f, 0x35, 0xf6, 0xf0,
	0x39, 0x69, 0xc3, 0x0a, 0x2e, 0xdf, 0x66, 0x54, 0x4d, 0xfa, 0x4f, 0x4f, 0x6a, 0x94, 0x26, 0xc9,
	0x54, 0x44, 0x2c, 0x34, 0x83, 0x6a, 0xa6, 0x19, 0xa7, 0x5c, 0x87, 0xb5, 0x21, 0x68, 0x46, 0xfc,
	0x30, 0x23, 0x56, 0x8b, 0x16, 0xad, 0xff, 0x1a, 0x36, 0x07, 0xf8, 0x52, 0x9c, 0x61, 0x1b, 0x65,
	0x69, 0xd4, 0x8f, 0x7c, 0xfd, 0xea, 0xae, 0x11, 0xb7, 0x2a, 0x33, 0x47, 0x36, 0x0c, 0xdf, 0x99,
	0xc0, 0x75, 0x49, 0x92, 0xa7, 0x70, 0xf5, 0x5f, 0xea, 0xc6, 0xd7, 0x4f, 0xe3, 0x6b, 0x6b, 0xba,
	0x8e, 0x76, 0xfa, 0x02, 0x2e, 0x49, 0x3a, 0xc0, 0x63, 0x40, 0x23, 0x6e, 0x94, 0x2b, 0x2a, 0x21,
	0x69, 0x2d, 0xce, 0x34, 0x79, 0x61, 0x88, 0x76, 0xcc, 0x04, 0x7c, 0x2c, 0xc9, 0x2e, 0x6c, 0x4e,
	0x91, 0x34, 0xce, 0x7e, 0x19, 0x67, 0xd6, 0xdf, 0x60, 0xe5, 0x69, 0xfb, 0xfd, 0x32, 0xd4, 0x1c,
	0x2a, 0x07, 0x9c, 0x49, 0x4a, 0x6e, 0xc3, 0x92, 0xce, 0x5d, 0x1e, 0x88, 0xa1, 0x9d, 0x3c, 0xd0,
	0x26, 0x93, 0x8f, 0xd5, 0x5f, 0xc7, 0x34, 0x92, 0x43, 0x38, 0xa7, 0x12, 0xe7, 0x8e, 0x45, 0x0e,
	0xcf, 0xb8, 0x82, 0xb0, 0x5d, 0x82, 0xcb, 0xc1, 0xec, 0xe2, 0xba, 0x33, 0x5a, 0x3b, 0xeb, 0xc9,
	0xe4, 0x03, 0x72, 0x1f, 0xaa, 0x79, 0xd2, 0xf1, 0x08, 0x95, 0xe2, 0xd6, 0x1f, 0x8a, 0xe6, 0x1e,
	0x74, 0xcd, 0x7f, 0xa7, 0x68, 0x27, 0xfb, 0x50, 0x11, 0xfc, 0x04, 0xf7, 0x54, 0x51, 0xf7, 0xec,
	0xf9, 0x6e, 0xa5, 0x5d, 0xec, 0x82, 0xed, 0xf0, 0x13, 0x47, 0x09, 0x90, 0x67, 0xb0, 0x31, 0x8a,
	0x79, 0x42, 0x93, 0x23, 0x2a, 0x54, 0x86, 0xe9, 0xec, 0xdb, 0x32, 0xba, 0x1e, 0x5d, 0x8d, 0x75,
	0x82, 0xc6, 0x69, 0x05, 0x2a, 0xa8, 0x4c, 0x2e, 0xc2, 0x32, 0x6a, 0x2b, 0x9d, 0x0f, 0x3d, 0x14,
	0x5a, 0x72, 0x96, 0x70, 0x89, 0x69, 0xdf, 0x2d, 0xdd, 0xc9, 0x8f, 0xbd, 0x33, 0x5e, 0xca, 0x36,
	0xd4, 0xfa, 0xea, 0x65, 0x14, 0xfc, 0x69, 0x0e, 0xb8, 0xaa, 0xbb, 0x11, 0xdc, 0x01, 0x30, 0x20,
	0xc3, 0xef, 0x83, 0xf5, 0xb9, 0x37, 0x33, 0x8a, 0x2b, 0xba, 0xbd, 0x87, 0xdd, 0x64, 0x0f, 0xd6,
	0x46, 0xae, 0x35, 0xff, 0x65, 0x36, 0xff, 0xff, 0x10, 0xd1, 0x1a, 0x0f, 0xa0, 0x1e, 0x7b, 0x2c,
	0xcc, 0xf0, 0xe8, 0x94, 0xf7, 0xaf, 0x73, 0x78, 0x87, 0x02, 0x30, 0x1b, 0xe7, 0xf9, 0x3e, 0x1e,
	0xbd, 0x1b, 0xd3, 0x63, 0x1a, 0x5b, 0xdf, 0xe6, 0xd9, 0x38, 0x43, 0x3c, 0x57, 0xc0, 0xde, 0x21,
	0x34, 0x23, 0x5e, 0x4a, 0xc9, 0xe8, 0xa3, 0xff, 0x6a, 0x27, 0xe4, 0x32, 0x78, 0x53, 0xd4, 0x83,
	0xb3, 0xfc, 0x2e, 0x1c, 0x2d, 0xeb, 0xcf, 0xef, 0xdd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70,
	0x28, 0xac, 0x54, 0x53, 0x06, 0x00, 0x00,
}
