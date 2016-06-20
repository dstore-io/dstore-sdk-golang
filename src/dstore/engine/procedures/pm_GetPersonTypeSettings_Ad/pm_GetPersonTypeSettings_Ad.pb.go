// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonTypeSettings_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersonTypeSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonTypeSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonTypeSettings_Ad

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
	PersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	KeyVariable      *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull  bool                        `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
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
	RowId                 int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PersonTypeDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=person_type_description,json=personTypeDescription" json:"person_type_description,omitempty"`
	Value                 *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
	KeyVariable           *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	PersonTypeId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PersonTypeDescription
	}
	return nil
}

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

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonTypeSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonTypeSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonTypeSettings_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x6d, 0x6b, 0x13, 0x41,
	0x10, 0x26, 0x4d, 0xd3, 0x96, 0x69, 0xb1, 0x75, 0x8b, 0x7a, 0x26, 0x20, 0x52, 0xbf, 0x28, 0xc2,
	0xc6, 0x37, 0x50, 0x04, 0x85, 0x16, 0x45, 0xf2, 0xa1, 0x47, 0xd9, 0x48, 0x41, 0xbf, 0x1c, 0xd7,
	0xde, 0x18, 0x16, 0x93, 0xdd, 0x63, 0x77, 0xd3, 0xd2, 0x7f, 0xe1, 0xdb, 0x9f, 0xf0, 0x9f, 0xf8,
	0x57, 0xd4, 0x3f, 0xe1, 0xdc, 0xed, 0x36, 0xb9, 0xbb, 0x82, 0x35, 0x7e, 0xb9, 0x63, 0x6e, 0x9e,
	0x67, 0x9e, 0xd9, 0x67, 0xe6, 0x16, 0x9e, 0x67, 0xd6, 0x69, 0x83, 0x7d, 0x54, 0x23, 0xa9, 0xb0,
	0x9f, 0x1b, 0x7d, 0x8c, 0xd9, 0xd4, 0xa0, 0xed, 0xe7, 0x93, 0xe4, 0x0d, 0xba, 0x03, 0x34, 0x56,
	0xab, 0xb7, 0x67, 0x39, 0x0e, 0xd1, 0x39, 0xa9, 0x46, 0x36, 0xd9, 0xcd, 0x38, 0xc1, 0x9c, 0x66,
	0xf7, 0x3c, 0x97, 0x7b, 0x2e, 0xff, 0x0b, 0xa1, 0xbb, 0x1d, 0x64, 0x4e, 0xd2, 0xf1, 0x14, 0xad,
	0xe7, 0x77, 0x6f, 0xd6, 0xb5, 0xd1, 0x18, 0x6d, 0x42, 0xaa, 0x57, 0x4f, 0x4d, 0xd0, 0xda, 0x74,
	0x84, 0x21, 0x79, 0xa7, 0x99, 0x74, 0xa9, 0x54, 0x1f, 0xb4, 0x99, 0xa4, 0x4e, 0x6a, 0xe5, 0x41,
	0x3b, 0xbf, 0x5b, 0x00, 0x07, 0xa9, 0x49, 0x29, 0x4b, 0x1d, 0xb1, 0x5d, 0xb8, 0x92, 0x97, 0x9d,
	0x25, 0x8e, 0x5a, 0x4b, 0x64, 0x16, 0xb5, 0x6e, 0xb7, 0xee, 0xae, 0x3f, 0xea, 0xf1, 0x70, 0x88,
	0xd0, 0x99, 0x54, 0x0e, 0x47, 0x68, 0x0e, 0x8b, 0x48, 0x6c, 0xe4, 0xb3, 0xc3, 0x0c, 0x32, 0xc6,
	0x61, 0xbb, 0x5e, 0x22, 0x51, 0xd3, 0xf1, 0x38, 0xfa, 0xb9, 0x4a, 0x85, 0xd6, 0xc4, 0x56, 0x15,
	0x1b, 0x53, 0x82, 0xbd, 0x80, 0x8d, 0x8f, 0x78, 0x96, 0x9c, 0xa4, 0x46, 0xa6, 0x47, 0x63, 0x8c,
	0x96, 0x4a, 0xc1, 0x6e, 0x43, 0xd0, 0x3a, 0x43, 0x26, 0x79, 0xbd, 0x75, 0xc2, 0x1f, 0x06, 0x38,
	0xbb, 0x0f, 0x57, 0xab, 0x74, 0x2f, 0xf6, 0xcb, 0x8b, 0x6d, 0x56, 0x80, 0x85, 0xd6, 0xce, 0x8f,
	0x65, 0x58, 0x13, 0x68, 0x73, 0xad, 0x2c, 0xb2, 0x07, 0xd0, 0x29, 0xbd, 0x0c, 0x47, 0x9c, 0x29,
	0x86, 0x39, 0x79, 0x9f, 0x5f, 0x17, 0x4f, 0xe1, 0x81, 0xec, 0x1d, 0x6c, 0x15, 0x2e, 0x26, 0x15,
	0x1b, 0xa9, 0xdd, 0x36, 0x91, 0x79, 0x83, 0xdc, 0x34, 0x7b, 0x9f, 0xe2, 0xc1, 0x3c, 0x16, 0x9b,
	0x93, 0xfa, 0x07, 0xf6, 0x0c, 0x56, 0xc3, 0xf4, 0xa2, 0x76, 0x59, 0xf1, 0xd6, 0x85, 0x8a, 0x7e,
	0xb6, 0xfb, 0xfe, 0x2d, 0xce, 0xe1, 0x6c, 0x00, 0x6d, 0xa3, 0x4f, 0xa3, 0xe5, 0x92, 0xf5, 0x94,
	0xff, 0xf3, 0xb2, 0xf1, 0x73, 0x23, 0xb8, 0xd0, 0xa7, 0xa2, 0xa8, 0xd1, 0xfd, 0xbe, 0x04, 0x6d,
	0x0a, 0xd8, 0x75, 0x58, 0xa1, 0xb0, 0x98, 0xfe, 0xa7, 0x98, 0xbc, 0xe9, 0x88, 0x0e, 0x85, 0x34,
	0xda, 0x21, 0xdc, 0xa8, 0x8e, 0x36, 0x43, 0x7b, 0x6c, 0x64, 0x5e, 0xda, 0xf0, 0x39, 0xbe, 0x74,
	0x6c, 0xd7, 0xe6, 0xa3, 0x7f, 0x35, 0x67, 0xb2, 0x87, 0xd0, 0x29, 0xc1, 0xd1, 0x97, 0xcb, 0x4b,
	0x78, 0x24, 0x7b, 0xd9, 0x58, 0x99, 0xaf, 0xf1, 0x62, 0x3b, 0xb3, 0x77, 0x61, 0xcb, 0xbf, 0xc5,
	0x0b, 0xae, 0xf9, 0xde, 0x10, 0x7a, 0x52, 0x37, 0xdd, 0x9e, 0x5d, 0x0b, 0xef, 0x9f, 0xfc, 0xcf,
	0x85, 0x71, 0xb4, 0x52, 0xfe, 0x94, 0x8f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xa3, 0x84,
	0x7f, 0x6f, 0x04, 0x00, 0x00,
}
