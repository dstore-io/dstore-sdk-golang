// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetForumSettings_Pu.proto
// DO NOT EDIT!

/*
Package fo_GetForumSettings_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetForumSettings_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetForumSettings_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	ForumId                        *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumIdNull                    bool                        `protobuf:"varint,1004,opt,name=forum_id_null,json=forumIdNull" json:"forum_id_null,omitempty"`
	KeyVariable                    *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull                bool                        `protobuf:"varint,1005,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	SelectResult                   *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=select_result,json=selectResult" json:"select_result,omitempty"`
	SelectResultNull               bool                        `protobuf:"varint,1006,opt,name=select_result_null,json=selectResultNull" json:"select_result_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1007,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetForumSettings_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetForumSettings_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetForumSettings_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_GetForumSettings_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0x26, 0xc6, 0x5c, 0x3c, 0xbd, 0x3a, 0x15, 0x5d, 0xd3, 0x52, 0xb4, 0x55, 0x10, 0xc4, 0xad,
	0x56, 0xd4, 0xbe, 0x88, 0x17, 0xbc, 0x10, 0xa1, 0xa5, 0xae, 0x52, 0x51, 0x84, 0x65, 0xdb, 0x9c,
	0x84, 0xa5, 0x9b, 0x9d, 0x38, 0x33, 0x6b, 0xe9, 0xbf, 0xb0, 0xef, 0x3e, 0xf9, 0xa7, 0x7c, 0xf2,
	0x07, 0x78, 0xff, 0x0b, 0x9e, 0x9d, 0x33, 0x69, 0xb3, 0xb1, 0x35, 0xf5, 0x25, 0x61, 0x72, 0xbe,
	0xdb, 0xcc, 0x9c, 0x33, 0x81, 0x5b, 0x2d, 0x6d, 0xa4, 0xc2, 0x25, 0x4c, 0x3b, 0x71, 0x8a, 0x4b,
	0x3d, 0x25, 0xb7, 0xb0, 0x95, 0x29, 0xd4, 0x4b, 0x6d, 0x19, 0x3e, 0x45, 0xf3, 0x44, 0xaa, 0xac,
	0xfb, 0x02, 0x8d, 0x89, 0xd3, 0x8e, 0x0e, 0xd7, 0x33, 0x9f, 0x10, 0x46, 0x8a, 0x4b, 0x4c, 0xf3,
	0x99, 0xe6, 0x1f, 0x8e, 0x6d, 0xcc, 0x38, 0xf1, 0xf7, 0x51, 0x92, 0xa1, 0x66, 0x6a, 0xe3, 0x7c,
	0xd1, 0x11, 0x95, 0x92, 0xca, 0x95, 0x66, 0x8b, 0xa5, 0x2e, 0x6a, 0x1d, 0x75, 0xd0, 0x15, 0x17,
	0x87, 0x8b, 0x26, 0x8a, 0xd3, 0xb6, 0x54, 0xdd, 0xc8, 0xc4, 0x32, 0x65, 0xd0, 0xc2, 0xe7, 0x2a,
	0xc0, 0x7a, 0xa4, 0x22, 0xaa, 0xa2, 0xd2, 0xe2, 0x2d, 0xcc, 0xf5, 0xe8, 0x5b, 0xa6, 0x61, 0xdc,
	0xc2, 0xd4, 0xc4, 0xed, 0x78, 0xcb, 0xa2, 0x43, 0x4e, 0xe4, 0x95, 0x2e, 0x94, 0xae, 0x8c, 0x2d,
	0x37, 0x7c, 0xb7, 0x1b, 0x97, 0x53, 0x1b, 0x45, 0x3b, 0xd8, 0xc8, 0x17, 0x41, 0x83, 0xf9, 0xcd,
	0x02, 0xdd, 0x96, 0xb4, 0x78, 0x06, 0x17, 0xff, 0xa5, 0x1e, 0xa6, 0x59, 0x92, 0x78, 0x5f, 0x6b,
	0xe4, 0x51, 0x0f, 0xe6, 0x8f, 0xd6, 0x59, 0x23, 0x98, 0x78, 0x00, 0x93, 0x4e, 0xcb, 0xec, 0xf6,
	0x90, 0x04, 0xbd, 0x13, 0x36, 0xdb, 0xec, 0x50, 0xb6, 0x38, 0x35, 0xd8, 0x41, 0xc5, 0xe1, 0xc6,
	0x99, 0xf2, 0x92, 0x18, 0xcd, 0x96, 0xf0, 0x61, 0xa6, 0x28, 0xc1, 0x01, 0xbe, 0x71, 0x80, 0xe9,
	0x41, 0xac, 0xb5, 0xbc, 0x03, 0xa7, 0xb2, 0x34, 0x7e, 0x97, 0x59, 0xb7, 0xf2, 0xc8, 0x93, 0xa8,
	0x33, 0x98, 0x8c, 0x2e, 0xc3, 0xe4, 0x3e, 0x91, 0x3d, 0xbe, 0xb3, 0xc7, 0x78, 0x1f, 0x62, 0xf5,
	0x6f, 0x43, 0xbd, 0x9d, 0x77, 0x44, 0x2e, 0x7f, 0x72, 0xf4, 0x66, 0x6a, 0x16, 0x4c, 0xf2, 0x8b,
	0x30, 0xd1, 0xe7, 0xb1, 0xfa, 0x0f, 0x56, 0x1f, 0x73, 0x00, 0x2b, 0x7e, 0x17, 0xc6, 0xb7, 0x71,
	0x97, 0x4e, 0x5a, 0xc5, 0xd1, 0x66, 0x82, 0x5e, 0x65, 0x64, 0xfe, 0x31, 0xc2, 0x6f, 0x38, 0xb8,
	0xb8, 0x0a, 0xa7, 0x07, 0xe9, 0xec, 0xf3, 0x93, 0x7d, 0xa6, 0x06, 0x80, 0xd6, 0xeb, 0x3e, 0x4c,
	0x68, 0x4c, 0x70, 0xcb, 0x84, 0x34, 0x17, 0x59, 0x62, 0xbc, 0xea, 0xa1, 0xbb, 0xd9, 0x94, 0x32,
	0xc1, 0x28, 0x75, 0x57, 0xc3, 0x8c, 0xc0, 0x12, 0xc4, 0x35, 0x10, 0x05, 0x05, 0xf6, 0xfb, 0xe5,
	0x6e, 0x66, 0x10, 0x6a, 0x0d, 0x9f, 0xc3, 0x39, 0x8d, 0x3d, 0x6a, 0x63, 0x92, 0x0f, 0x63, 0xd7,
	0x5e, 0x79, 0x57, 0x69, 0xaf, 0x36, 0x72, 0x9f, 0x67, 0xf6, 0xa9, 0x4d, 0x6e, 0x37, 0xfa, 0x59,
	0x8b, 0x7b, 0x30, 0x77, 0x84, 0x24, 0x67, 0xf9, 0xcd, 0x59, 0xbc, 0xc3, 0xc8, 0x79, 0xa6, 0x85,
	0x2f, 0x65, 0xa8, 0x53, 0xc4, 0x9e, 0x4c, 0x35, 0x8a, 0xeb, 0x50, 0xb1, 0x73, 0x3b, 0x3c, 0x40,
	0xee, 0x39, 0xe0, 0x99, 0x7e, 0x9c, 0x7f, 0x06, 0x0c, 0x14, 0xaf, 0x61, 0x3a, 0x9f, 0xd8, 0x70,
	0x60, 0x64, 0xa9, 0xc3, 0xcb, 0x44, 0xf6, 0x87, 0xc8, 0xc3, 0x83, 0xbd, 0x4a, 0xeb, 0xe6, 0xc1,
	0x3a, 0x98, 0xea, 0x16, 0x7f, 0x10, 0x2b, 0x50, 0x73, 0x2f, 0x05, 0x75, 0x71, 0xae, 0x38, 0xff,
	0x97, 0x22, 0xbf, 0x23, 0xab, 0xfc, 0x1d, 0xf4, 0xe1, 0xe2, 0x11, 0x94, 0x95, 0xdc, 0xa1, 0xe6,
	0xcc, 0x59, 0xcb, 0xfe, 0x71, 0xde, 0x34, 0xbf, 0x7f, 0x06, 0x7e, 0x20, 0x77, 0x82, 0x9c, 0x9e,
	0x1f, 0x86, 0xbd, 0x06, 0x0f, 0x47, 0xde, 0x0d, 0x03, 0x1b, 0x7b, 0x25, 0x28, 0x13, 0x5d, 0x9c,
	0x85, 0x2a, 0x09, 0xe4, 0xf3, 0xf1, 0x61, 0x8d, 0xb8, 0x95, 0xa0, 0x42, 0x4b, 0x9a, 0x80, 0x1b,
	0x7d, 0xc5, 0xbd, 0xb5, 0x63, 0x4a, 0xd2, 0xfd, 0x16, 0xe7, 0xe1, 0xd3, 0xc7, 0xd2, 0x7f, 0x4d,
	0xc4, 0xc3, 0x57, 0x30, 0x1b, 0xcb, 0xa1, 0x23, 0x38, 0xf8, 0x37, 0x78, 0xb3, 0xd2, 0x91, 0xba,
	0xb5, 0xdd, 0xaf, 0xb7, 0x8e, 0xff, 0x87, 0xb1, 0x59, 0xb5, 0x2f, 0xf3, 0xcd, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0xfb, 0x0a, 0x8d, 0x6a, 0x06, 0x00, 0x00,
}
