// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonTypeSettings_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonTypeSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonTypeSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonTypeSettings_Ad

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
	PersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	KeyVariable      *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull  bool                        `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	Value            *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull        bool                        `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
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

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonTypeSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonTypeSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonTypeSettings_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonTypeSettings_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0xa7, 0x89, 0x69, 0xeb, 0xb4, 0xd8, 0xba, 0x05, 0x39, 0x13, 0x2c, 0x52, 0x5f, 0x04, 0x65,
	0x23, 0xfa, 0x22, 0x42, 0xc1, 0x0a, 0x3e, 0x04, 0x4c, 0x29, 0xab, 0x14, 0xf4, 0xe5, 0xb8, 0xba,
	0xd3, 0x63, 0x69, 0xb2, 0x7b, 0xec, 0x6e, 0x5a, 0xf2, 0x2d, 0xfc, 0x42, 0xfa, 0x7d, 0xd4, 0x2f,
	0xe1, 0xdc, 0xee, 0x26, 0xe6, 0x4e, 0xb0, 0xe4, 0x25, 0xb9, 0xd9, 0xf9, 0xfd, 0x99, 0x9d, 0x99,
	0x85, 0x63, 0xe9, 0xbc, 0xb1, 0x38, 0x44, 0x5d, 0x2a, 0x8d, 0xc3, 0xca, 0x9a, 0xaf, 0x28, 0x67,
	0x16, 0xdd, 0xb0, 0x9a, 0xe6, 0x63, 0x23, 0xd5, 0xe5, 0xfc, 0x0c, 0xad, 0x33, 0xfa, 0xd3, 0xbc,
	0xc2, 0x8f, 0xe8, 0xbd, 0xd2, 0xa5, 0xcb, 0x4f, 0x24, 0x27, 0xa4, 0x37, 0xec, 0x79, 0xa4, 0xf3,
	0x48, 0xe7, 0xff, 0xe7, 0xf4, 0x0f, 0x92, 0xd9, 0x75, 0x31, 0x99, 0xa1, 0x8b, 0x12, 0xfd, 0x87,
	0xcd, 0x0a, 0xd0, 0x5a, 0x63, 0x53, 0x6a, 0xd0, 0x4c, 0x4d, 0xd1, 0xb9, 0xa2, 0xc4, 0x94, 0x7c,
	0xd2, 0x4e, 0xfa, 0x42, 0xe9, 0x4b, 0x63, 0xa7, 0x85, 0x57, 0x46, 0x47, 0xd0, 0xd1, 0xf7, 0x0e,
	0xc0, 0x59, 0x61, 0x0b, 0xca, 0x52, 0x45, 0xec, 0x04, 0xee, 0x55, 0xa1, 0xb2, 0xdc, 0x53, 0x69,
	0xb9, 0x92, 0xd9, 0xc6, 0xe3, 0x8d, 0xa7, 0x3b, 0x2f, 0x07, 0x3c, 0xdd, 0x23, 0x55, 0xa6, 0xb4,
	0xc7, 0x12, 0xed, 0x79, 0x1d, 0x89, 0xdd, 0x6a, 0x79, 0x99, 0x91, 0x64, 0x1c, 0x0e, 0x9a, 0x12,
	0xb9, 0x9e, 0x4d, 0x26, 0xd9, 0xcf, 0x2d, 0x12, 0xda, 0x16, 0xfb, 0xab, 0xd8, 0x53, 0x4a, 0xb0,
	0x63, 0xd8, 0xbd, 0xc2, 0x79, 0x7e, 0x5d, 0x58, 0x55, 0x5c, 0x4c, 0x30, 0xeb, 0x04, 0xc3, 0x7e,
	0xcb, 0xd0, 0x79, 0x4b, 0x4d, 0x8a, 0x7e, 0x3b, 0x84, 0x3f, 0x4f, 0x70, 0xf6, 0x0c, 0xee, 0xaf,
	0xd2, 0xa3, 0xd9, 0xaf, 0x68, 0xb6, 0xb7, 0x02, 0x0c, 0x5e, 0x2f, 0xa0, 0x17, 0xf4, 0xb2, 0xee,
	0xad, 0x26, 0x11, 0xc8, 0x0e, 0x01, 0xc2, 0x47, 0xd4, 0xfd, 0x1d, 0x75, 0xef, 0x86, 0xa3, 0x5a,
	0xf1, 0xe8, 0x47, 0x07, 0xb6, 0x05, 0xba, 0xca, 0x68, 0x87, 0xb5, 0x7c, 0x98, 0x4e, 0x6a, 0xda,
	0x52, 0x3e, 0x0d, 0x3f, 0x4e, 0xee, 0x7d, 0xfd, 0x2b, 0x22, 0x90, 0x7d, 0x86, 0xfd, 0x7a, 0x2e,
	0xf9, 0xca, 0x60, 0xa8, 0x01, 0x5d, 0x22, 0xf3, 0x16, 0xb9, 0x3d, 0xbe, 0x31, 0xc5, 0xa3, 0xbf,
	0xb1, 0xd8, 0x9b, 0x36, 0x0f, 0xd8, 0x6b, 0xd8, 0x4a, 0xfb, 0x40, 0xb7, 0xad, 0x15, 0x0f, 0xff,
	0x51, 0x8c, 0xdb, 0x32, 0x8e, 0xff, 0x62, 0x01, 0x67, 0x1f, 0xa0, 0x6b, 0xcd, 0x4d, 0x76, 0x27,
	0xb0, 0xde, 0xf0, 0x75, 0x36, 0x98, 0x2f, 0x7a, 0xc1, 0x85, 0xb9, 0x11, 0xb5, 0x4c, 0xff, 0x11,
	0x74, 0xe9, 0x9b, 0x3d, 0x80, 0x4d, 0x8a, 0xea, 0x8d, 0xfa, 0x76, 0x4a, 0xdd, 0xe9, 0x89, 0x1e,
	0x85, 0x23, 0xf9, 0x2e, 0x87, 0x81, 0x32, 0x6d, 0x8f, 0xe5, 0x23, 0xfb, 0xf2, 0xb6, 0x34, 0x4e,
	0x5e, 0x2d, 0xf2, 0x72, 0xfd, 0x77, 0x78, 0xb1, 0x19, 0x16, 0xfd, 0xd5, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x20, 0x13, 0x68, 0xe0, 0xc9, 0x03, 0x00, 0x00,
}
