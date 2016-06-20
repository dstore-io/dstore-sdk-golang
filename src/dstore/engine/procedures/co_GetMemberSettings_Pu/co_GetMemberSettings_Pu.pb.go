// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetMemberSettings_Pu.proto
// DO NOT EDIT!

/*
Package co_GetMemberSettings_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetMemberSettings_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetMemberSettings_Pu

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
	UniqueId                         *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                     bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull   bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                      *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                  bool                        `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	GetEntryForCommunityMemberId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=get_entry_for_community_member_id,json=getEntryForCommunityMemberId" json:"get_entry_for_community_member_id,omitempty"`
	GetEntryForCommunityMemberIdNull bool                        `protobuf:"varint,1004,opt,name=get_entry_for_community_member_id_null,json=getEntryForCommunityMemberIdNull" json:"get_entry_for_community_member_id_null,omitempty"`
	KeyVariable                      *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull                  bool                        `protobuf:"varint,1005,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	SeparatorInIdentVals             *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull         bool                        `protobuf:"varint,1006,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetGetEntryForCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetEntryForCommunityMemberId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
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
	RowId       int32                      `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value       *dstore_values.StringValue `protobuf:"bytes,10001,opt,name=value" json:"value,omitempty"`
	KeyVariable *dstore_values.StringValue `protobuf:"bytes,10002,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetMemberSettings_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetMemberSettings_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetMemberSettings_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xed, 0x6a, 0x13, 0x41,
	0x14, 0xa5, 0xc6, 0xa4, 0xf5, 0xb6, 0x58, 0x5d, 0x45, 0xd7, 0xb4, 0x94, 0xb6, 0x52, 0x11, 0x84,
	0x4d, 0xb5, 0xa0, 0xfe, 0x51, 0x41, 0x89, 0x12, 0x21, 0xb1, 0x5d, 0xa1, 0xa0, 0x08, 0xcb, 0x26,
	0xb9, 0x0d, 0x83, 0xd9, 0x9d, 0x38, 0x33, 0x6b, 0xc9, 0x5b, 0x54, 0x5f, 0xc3, 0x57, 0xf0, 0x85,
	0xfc, 0x7c, 0x06, 0xef, 0xec, 0xdd, 0x7c, 0xad, 0xa6, 0x6b, 0xff, 0xb4, 0xcc, 0xde, 0x73, 0xce,
	0x3d, 0xb9, 0x5f, 0x70, 0xbf, 0xab, 0x8d, 0x54, 0x58, 0xc3, 0xb8, 0x27, 0x62, 0xac, 0x0d, 0x94,
	0xec, 0x60, 0x37, 0x51, 0xa8, 0x6b, 0x1d, 0x19, 0xbc, 0x40, 0xd3, 0xc4, 0xa8, 0x8d, 0xea, 0x35,
	0x1a, 0x23, 0xe2, 0x9e, 0x0e, 0xf6, 0x13, 0x8f, 0x20, 0x46, 0x3a, 0x3b, 0xcc, 0xf3, 0x98, 0xe7,
	0xcd, 0x01, 0x57, 0xaf, 0x64, 0xf2, 0x1f, 0xc3, 0x7e, 0x82, 0x9a, 0xb9, 0xd5, 0x1b, 0xb3, 0x39,
	0x51, 0x29, 0xa9, 0xb2, 0xd0, 0xda, 0x6c, 0x28, 0x42, 0xad, 0xc3, 0x1e, 0x66, 0xc1, 0x9b, 0xf9,
	0xa0, 0x09, 0x45, 0x7c, 0x24, 0x55, 0x14, 0x1a, 0x21, 0x63, 0x06, 0x6d, 0x7f, 0xad, 0x00, 0xec,
	0x87, 0x2a, 0xa4, 0x28, 0x2a, 0xed, 0x3c, 0x80, 0x0b, 0x49, 0x2c, 0x3e, 0x24, 0x18, 0x88, 0xae,
	0xbb, 0xb0, 0xb9, 0x70, 0x7b, 0xf9, 0x5e, 0xd5, 0xcb, 0xbc, 0x67, 0xa6, 0xb4, 0x51, 0x64, 0xf7,
	0xd0, 0x3e, 0xfc, 0x25, 0x06, 0x37, 0xba, 0xce, 0x0e, 0x5c, 0x1c, 0x13, 0x83, 0x38, 0xe9, 0xf7,
	0xdd, 0x6f, 0x8b, 0x44, 0x5f, 0xf2, 0x57, 0x46, 0x90, 0x16, 0x7d, 0x74, 0xde, 0xc1, 0xfa, 0x80,
	0xf2, 0xc8, 0x98, 0x60, 0x18, 0x1b, 0x71, 0x24, 0x3a, 0xa9, 0x9b, 0x80, 0xc5, 0xdd, 0x73, 0x85,
	0x29, 0xab, 0xcc, 0x6f, 0xcc, 0xd0, 0xd3, 0x90, 0x76, 0x5e, 0xc2, 0xd6, 0x69, 0xea, 0xec, 0xeb,
	0x3b, 0xfb, 0xda, 0x98, 0xaf, 0x93, 0x3a, 0x7d, 0x0c, 0x2b, 0x1d, 0x19, 0x45, 0xe4, 0xde, 0x0c,
	0x6d, 0x31, 0x4a, 0xa9, 0xb3, 0xb5, 0x9c, 0x33, 0x11, 0x1b, 0xec, 0xa1, 0x62, 0x6b, 0xcb, 0x63,
	0x02, 0x15, 0xe4, 0x0e, 0x5c, 0x9e, 0xe6, 0x73, 0xee, 0x1f, 0x9c, 0x7b, 0x75, 0x0a, 0x98, 0x26,
	0xeb, 0xc2, 0x56, 0x0f, 0x4d, 0x40, 0x56, 0xd4, 0x30, 0xa0, 0x16, 0x05, 0x13, 0x6a, 0x94, 0x0e,
	0x89, 0x75, 0x70, 0xbe, 0xd8, 0xc1, 0x3a, 0xa9, 0xd4, 0xad, 0xc8, 0x73, 0xa9, 0x9e, 0x8d, 0x24,
	0x78, 0xcc, 0xc8, 0xd2, 0x01, 0xdc, 0x2a, 0xcc, 0xc2, 0x3e, 0x7f, 0xb2, 0xcf, 0xcd, 0xd3, 0xe4,
	0x52, 0xe3, 0x8f, 0x60, 0xe5, 0x3d, 0x0e, 0xa9, 0xbe, 0x4a, 0x84, 0xed, 0x3e, 0xba, 0xe5, 0xc2,
	0xfe, 0x2d, 0x13, 0xfe, 0x30, 0x83, 0xdb, 0x22, 0x4d, 0xd3, 0x39, 0xf9, 0xaf, 0xac, 0x48, 0x53,
	0xc0, 0x34, 0xd7, 0x01, 0x5c, 0xd7, 0x38, 0xa0, 0x59, 0x25, 0xe9, 0x40, 0x64, 0x3d, 0xb6, 0xad,
	0xd5, 0x6e, 0xa5, 0x30, 0xed, 0xd5, 0x31, 0xb5, 0xc1, 0x3d, 0xa7, 0xcf, 0xda, 0x79, 0x02, 0xeb,
	0x73, 0x24, 0xd9, 0xca, 0x6f, 0xb6, 0xe2, 0xfe, 0x8b, 0x6c, 0x3d, 0x6d, 0x7f, 0x29, 0xc1, 0x92,
	0x8f, 0x7a, 0x20, 0x63, 0x8d, 0xce, 0x2e, 0x94, 0xd3, 0xe5, 0xcc, 0x2f, 0x4e, 0xb6, 0xf4, 0xbc,
	0xb8, 0x75, 0xfb, 0xd7, 0x67, 0xa0, 0xf3, 0x06, 0x2e, 0xd9, 0xb5, 0x0c, 0xa6, 0xf6, 0x92, 0x56,
	0xa0, 0x44, 0x64, 0x2f, 0x47, 0xce, 0x6f, 0x6f, 0x93, 0xde, 0x8d, 0xc9, 0xdb, 0x5f, 0x8d, 0x66,
	0x3f, 0x38, 0x0f, 0x61, 0x31, 0x3b, 0x07, 0x34, 0xba, 0x56, 0x71, 0xe3, 0x2f, 0x45, 0x3e, 0x16,
	0x4d, 0xfe, 0xef, 0x8f, 0xe0, 0x4e, 0x1d, 0x4a, 0x4a, 0x1e, 0xd3, 0xb8, 0x59, 0xd6, 0x9e, 0xf7,
	0x5f, 0x97, 0xcb, 0x1b, 0x15, 0xc1, 0xf3, 0xe5, 0xb1, 0x6f, 0xf9, 0xd5, 0x93, 0x05, 0x28, 0xd1,
	0xc3, 0xb9, 0x06, 0x15, 0x7a, 0xda, 0x01, 0x3e, 0x69, 0x51, 0x5d, 0xca, 0x7e, 0x99, 0x9e, 0x34,
	0x8d, 0x77, 0xa1, 0x9c, 0xf6, 0xc9, 0xfd, 0xd4, 0x2a, 0xec, 0x1e, 0x23, 0xed, 0x4e, 0xce, 0x4c,
	0xdb, 0xe7, 0xd6, 0x99, 0xc6, 0xed, 0xe9, 0x2b, 0x58, 0x13, 0x32, 0xf7, 0x83, 0x26, 0x27, 0xfc,
	0xed, 0xee, 0x59, 0x8f, 0x7b, 0xbb, 0x92, 0x1e, 0xd1, 0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0x0a, 0xe4, 0xa0, 0x17, 0x06, 0x00, 0x00,
}
