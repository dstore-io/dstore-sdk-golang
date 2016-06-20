// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunitySettings.proto
// DO NOT EDIT!

/*
Package co_GetCommunitySettings is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunitySettings.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunitySettings

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
	CommunityId                 *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull             bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	KeyVariable                 *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull             bool                        `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	ResultOnlyInOutputParam     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=result_only_in_output_param,json=resultOnlyInOutputParam" json:"result_only_in_output_param,omitempty"`
	ResultOnlyInOutputParamNull bool                        `protobuf:"varint,1003,opt,name=result_only_in_output_param_null,json=resultOnlyInOutputParamNull" json:"result_only_in_output_param_null,omitempty"`
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

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Parameters) GetResultOnlyInOutputParam() *dstore_values.BooleanValue {
	if m != nil {
		return m.ResultOnlyInOutputParam
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
	RowId         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CommunityId   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	Value         *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
	CommunityName *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	KeyVariable   *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunitySettings.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunitySettings.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunitySettings.Response.Row")
}

var fileDescriptor0 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0xa6, 0xbd, 0xa6, 0x2d, 0xdb, 0x6a, 0x75, 0x05, 0x3d, 0x73, 0x20, 0xa5, 0x22, 0x08, 0xc2,
	0xa5, 0x5a, 0x10, 0x5f, 0x54, 0xa8, 0x54, 0xc9, 0x43, 0x2f, 0x72, 0x42, 0x41, 0x5f, 0x8e, 0x4d,
	0x32, 0x86, 0xc5, 0xbb, 0xdd, 0xb0, 0xbb, 0xd7, 0xd2, 0xff, 0xc0, 0x47, 0x7f, 0xfe, 0x83, 0xbe,
	0xa9, 0xff, 0x84, 0xb3, 0xb7, 0x9b, 0xe4, 0x72, 0x5a, 0x9b, 0xbc, 0x24, 0xec, 0xcd, 0xf7, 0xcd,
	0x37, 0x3b, 0xf3, 0xcd, 0x92, 0xc7, 0x43, 0x6d, 0xa4, 0x82, 0x0e, 0x88, 0x11, 0x17, 0xd0, 0x19,
	0x2b, 0x39, 0x80, 0x61, 0xa9, 0x40, 0x77, 0x06, 0x32, 0x7b, 0x05, 0xe6, 0x85, 0x2c, 0x8a, 0x52,
	0x70, 0x73, 0xfe, 0x06, 0x8c, 0xe1, 0x62, 0xa4, 0x63, 0x84, 0x18, 0x49, 0xef, 0x39, 0x5e, 0xec,
	0x78, 0xf1, 0x05, 0xe0, 0xf6, 0x0d, 0x9f, 0xfe, 0x94, 0xe5, 0x25, 0x78, 0x6e, 0xfb, 0xf6, 0xbc,
	0x26, 0x28, 0x25, 0x95, 0x0f, 0x45, 0xf3, 0xa1, 0x02, 0xb4, 0x66, 0x23, 0xf0, 0xc1, 0xbb, 0xcd,
	0xa0, 0x61, 0x5c, 0xbc, 0x97, 0xaa, 0x60, 0x86, 0x4b, 0xe1, 0x40, 0x7b, 0x1f, 0x03, 0x42, 0x5e,
	0x33, 0xc5, 0x30, 0x0a, 0x4a, 0xd3, 0x67, 0x64, 0x7b, 0x30, 0xa9, 0x2a, 0xe3, 0xc3, 0x70, 0x65,
	0x77, 0xe5, 0xfe, 0xd6, 0xa3, 0x28, 0xf6, 0xe5, 0xfb, 0xba, 0xb8, 0x30, 0x30, 0x02, 0x75, 0x62,
	0x4f, 0xe9, 0xd6, 0x94, 0xd0, 0x1d, 0xd2, 0x07, 0xe4, 0x7a, 0x9d, 0x9f, 0x89, 0x32, 0xcf, 0xc3,
	0x9f, 0x1b, 0x98, 0x65, 0x33, 0xdd, 0xa9, 0x01, 0x13, 0xfc, 0x4e, 0x9f, 0x92, 0xed, 0x0f, 0x70,
	0x9e, 0x9d, 0x32, 0xc5, 0x59, 0x3f, 0x87, 0x70, 0xb5, 0x12, 0x6b, 0x37, 0xc4, 0xb4, 0x51, 0xd8,
	0x1b, 0xaf, 0x85, 0xf8, 0x13, 0x0f, 0xb7, 0x5a, 0x75, 0xba, 0xd3, 0xfa, 0xe5, 0xb5, 0x6a, 0xc0,
	0x4a, 0xeb, 0x2d, 0x89, 0x70, 0x4a, 0x65, 0x6e, 0x32, 0x29, 0x72, 0x2c, 0x4d, 0x64, 0xb2, 0x34,
	0xe3, 0xd2, 0x64, 0x63, 0x7b, 0xf9, 0x30, 0xf8, 0xe7, 0x3d, 0xfb, 0x52, 0xe6, 0xc0, 0x84, 0xd3,
	0xbe, 0xe5, 0xf8, 0x3d, 0xa4, 0x77, 0x45, 0xaf, 0x22, 0x57, 0x8d, 0xa3, 0x2f, 0xc9, 0xee, 0x7f,
	0x52, 0xbb, 0xb2, 0x7e, 0xbb, 0xb2, 0xa2, 0x0b, 0x72, 0xd8, 0x12, 0xf7, 0x7e, 0xac, 0x91, 0xcd,
	0x14, 0xf4, 0x58, 0x0a, 0x0d, 0x74, 0x9f, 0xb4, 0xaa, 0x41, 0xfb, 0x09, 0x4c, 0x9b, 0xe2, 0x0d,
	0xe4, 0x4c, 0x70, 0x64, 0x7f, 0x53, 0x07, 0xc4, 0x1b, 0x5e, 0xb3, 0x23, 0xce, 0x6a, 0x33, 0xc6,
	0x8e, 0x06, 0x48, 0x8e, 0x1b, 0xe4, 0xa6, 0x13, 0x8e, 0xf1, 0xdc, 0x9d, 0x9d, 0xd3, 0x9d, 0x62,
	0xfe, 0x03, 0x7d, 0x42, 0x36, 0xbc, 0xb5, 0xb0, 0x51, 0x36, 0xe3, 0x9d, 0xbf, 0x32, 0x3a, 0xe3,
	0x1d, 0xbb, 0xff, 0x74, 0x02, 0xa7, 0x47, 0x24, 0x50, 0xf2, 0x2c, 0x5c, 0xab, 0x58, 0x07, 0xf1,
	0x42, 0x5b, 0x10, 0x4f, 0x9a, 0x10, 0xa7, 0xf2, 0x2c, 0xb5, 0x7c, 0xdb, 0x8d, 0x6a, 0x24, 0x21,
	0x5c, 0x6a, 0x11, 0x07, 0x6c, 0x7f, 0x5f, 0x25, 0x01, 0xd2, 0xe9, 0x4d, 0xb2, 0x8e, 0x09, 0xac,
	0x95, 0x3f, 0x25, 0xc8, 0x6d, 0xa5, 0x2d, 0x3c, 0xa2, 0x51, 0x9f, 0x37, 0x8c, 0xfe, 0x39, 0x59,
	0xd2, 0xe9, 0x0f, 0x27, 0x25, 0x7d, 0x49, 0x16, 0xac, 0x89, 0x1e, 0x92, 0xab, 0x33, 0x4d, 0x81,
	0x2b, 0x17, 0x7e, 0xbd, 0x9c, 0x7b, 0x65, 0x4a, 0x49, 0x90, 0x61, 0x17, 0x74, 0x6e, 0x67, 0xbe,
	0x25, 0x4b, 0x2d, 0xcd, 0x61, 0x8f, 0x44, 0x5c, 0x36, 0xe6, 0x30, 0x7b, 0xc5, 0xde, 0xed, 0x2f,
	0xfb, 0xbe, 0xf5, 0xd7, 0xab, 0x77, 0xe4, 0xe0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x44,
	0xf8, 0xed, 0x1a, 0x05, 0x00, 0x00,
}
