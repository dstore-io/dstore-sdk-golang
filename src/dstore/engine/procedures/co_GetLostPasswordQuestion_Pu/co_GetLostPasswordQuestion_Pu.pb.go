// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetLostPasswordQuestion_Pu.proto
// DO NOT EDIT!

/*
Package co_GetLostPasswordQuestion_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetLostPasswordQuestion_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetLostPasswordQuestion_Pu

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
	CommunityId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull   bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	Nickname          *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	NicknameNull      bool                        `protobuf:"varint,1002,opt,name=nickname_null,json=nicknameNull" json:"nickname_null,omitempty"`
	CaseSensitive     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull bool                        `protobuf:"varint,1003,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
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

func (m *Parameters) GetNickname() *dstore_values.StringValue {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Question        *dstore_values.StringValue                       `protobuf:"bytes,101,opt,name=question" json:"question,omitempty"`
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

func (m *Response) GetQuestion() *dstore_values.StringValue {
	if m != nil {
		return m.Question
	}
	return nil
}

type Response_Row struct {
	RowId                 int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Question              *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=question" json:"question,omitempty"`
	NicknameCaseSensitive *dstore_values.StringValue `protobuf:"bytes,10002,opt,name=nickname_case_sensitive,json=nicknameCaseSensitive" json:"nickname_case_sensitive,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetQuestion() *dstore_values.StringValue {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *Response_Row) GetNicknameCaseSensitive() *dstore_values.StringValue {
	if m != nil {
		return m.NicknameCaseSensitive
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetLostPasswordQuestion_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetLostPasswordQuestion_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetLostPasswordQuestion_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x6a, 0x13, 0x4d,
	0x14, 0x27, 0x4d, 0xd3, 0x86, 0xd3, 0xf6, 0xeb, 0xd7, 0x29, 0xea, 0x9a, 0x80, 0x48, 0xf5, 0x42,
	0x28, 0x6c, 0x44, 0xa1, 0x2a, 0x88, 0x17, 0x15, 0x91, 0xa2, 0x0d, 0x71, 0x8a, 0x82, 0xde, 0x2c,
	0xd3, 0xdd, 0x63, 0x18, 0xcc, 0xce, 0xd4, 0x99, 0xd9, 0x06, 0xef, 0x7d, 0x00, 0xf5, 0x35, 0x7c,
	0x15, 0x5f, 0x44, 0x7d, 0x09, 0x67, 0x76, 0x66, 0xd3, 0xdd, 0x55, 0xac, 0x7a, 0x93, 0x64, 0xce,
	0xf9, 0xfd, 0xc9, 0x9c, 0xf3, 0xdb, 0x85, 0xfb, 0x99, 0x36, 0x52, 0xe1, 0x08, 0xc5, 0x94, 0x0b,
	0x1c, 0x9d, 0x28, 0x99, 0x62, 0x56, 0x28, 0xd4, 0xa3, 0x54, 0x26, 0x8f, 0xd1, 0x3c, 0x95, 0xda,
	0x4c, 0x98, 0xd6, 0x73, 0xa9, 0xb2, 0x67, 0x05, 0x6a, 0xc3, 0xa5, 0x48, 0x26, 0x45, 0x6c, 0x81,
	0x46, 0x92, 0x5d, 0xcf, 0x8e, 0x3d, 0x3b, 0xfe, 0x2d, 0x65, 0xb0, 0x1d, 0xac, 0x4e, 0xd9, 0xcc,
	0x56, 0xbd, 0xc2, 0xe0, 0x72, 0xd3, 0x1f, 0x95, 0x92, 0x2a, 0xb4, 0x86, 0xcd, 0x56, 0x8e, 0x5a,
	0xb3, 0x29, 0x86, 0xe6, 0xb5, 0x76, 0xd3, 0x30, 0x2e, 0x5e, 0x4b, 0x95, 0x33, 0x67, 0xe7, 0x41,
	0x3b, 0x5f, 0x96, 0x00, 0x26, 0x4c, 0x31, 0xdb, 0x45, 0xa5, 0xc9, 0x03, 0x58, 0x4f, 0x65, 0x9e,
	0x17, 0x82, 0x9b, 0x77, 0x09, 0xcf, 0xa2, 0xce, 0xd5, 0xce, 0x8d, 0xb5, 0x5b, 0xc3, 0x38, 0x5c,
	0x22, 0xfc, 0x2f, 0x2e, 0x0c, 0x4e, 0x51, 0xbd, 0x70, 0x27, 0xba, 0xb6, 0x20, 0x1c, 0x64, 0x64,
	0x17, 0xb6, 0xea, 0xfc, 0x44, 0x14, 0xb3, 0x59, 0xf4, 0x75, 0xd5, 0xaa, 0xf4, 0xe9, 0x66, 0x0d,
	0x38, 0xb6, 0x75, 0xb2, 0x07, 0x7d, 0xc1, 0xd3, 0x37, 0xc2, 0x9a, 0x47, 0x4b, 0xa5, 0xd1, 0xa0,
	0x65, 0xa4, 0x8d, 0xe2, 0x62, 0xea, 0x7d, 0x16, 0x58, 0x72, 0x1d, 0x36, 0xaa, 0xdf, 0xde, 0xe0,
	0x9b, 0x37, 0x58, 0xaf, 0xaa, 0xa5, 0xfa, 0x3e, 0xfc, 0x97, 0x32, 0x8d, 0x89, 0x46, 0xa1, 0xb9,
	0xe1, 0xa7, 0x18, 0x75, 0x7f, 0x79, 0x99, 0x63, 0x29, 0x67, 0xc8, 0x84, 0x37, 0xd9, 0x70, 0x94,
	0xa3, 0x8a, 0x41, 0x46, 0xb0, 0xdd, 0xd4, 0xf0, 0x7e, 0xdf, 0xbd, 0xdf, 0x56, 0x03, 0xec, 0x4c,
	0x77, 0xde, 0x2f, 0x43, 0x9f, 0xa2, 0x3e, 0x91, 0x42, 0x23, 0xb9, 0x09, 0xbd, 0x72, 0x59, 0x61,
	0x8a, 0x8b, 0xcb, 0x85, 0x28, 0xf8, 0x45, 0x3e, 0x72, 0x9f, 0xd4, 0x03, 0xc9, 0x4b, 0xf8, 0xdf,
	0xad, 0x29, 0xa9, 0xed, 0xc9, 0x4e, 0xa6, 0x6b, 0xc9, 0x71, 0x8b, 0xdc, 0xde, 0xe6, 0xa1, 0x3d,
	0x1f, 0x9c, 0x9d, 0xe9, 0x66, 0xde, 0x2c, 0x90, 0xbb, 0xb0, 0x1a, 0xe2, 0x61, 0xe7, 0xe0, 0x14,
	0xaf, 0xfc, 0xa4, 0xe8, 0xc3, 0x73, 0xe8, 0xbf, 0x69, 0x05, 0x27, 0x4f, 0xa0, 0xab, 0xe4, 0x3c,
	0x5a, 0x2e, 0x59, 0xf7, 0xe2, 0xbf, 0xc8, 0x73, 0x5c, 0x8d, 0x22, 0xa6, 0x72, 0x4e, 0x9d, 0x8a,
	0xdb, 0xf9, 0xdb, 0x00, 0x88, 0xf0, 0xfc, 0x9d, 0x57, 0xd8, 0xc1, 0xe7, 0x0e, 0x74, 0xad, 0x08,
	0xb9, 0x08, 0x2b, 0x56, 0xc6, 0x45, 0xf3, 0xc3, 0xd8, 0xd2, 0x7b, 0xb4, 0x67, 0x8f, 0x36, 0x78,
	0x77, 0x6a, 0xba, 0x1f, 0xc7, 0x7f, 0x2e, 0x4c, 0x8e, 0xe0, 0xd2, 0x22, 0x4c, 0xad, 0xbc, 0x7c,
	0x3a, 0x5f, 0xe7, 0x42, 0xc5, 0x7d, 0x58, 0x8f, 0xc2, 0xfe, 0x73, 0x18, 0x72, 0xd9, 0x9a, 0xd4,
	0xd9, 0x7b, 0xe3, 0xd5, 0xde, 0xbf, 0xbd, 0x51, 0x8e, 0x57, 0xca, 0x67, 0xf6, 0xf6, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x09, 0x68, 0x45, 0x92, 0x04, 0x00, 0x00,
}
