// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyAppSettings_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyAppSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyAppSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyAppSettings_Ad

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
	ApplicationId             *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationIdNull         bool                        `protobuf:"varint,1001,opt,name=application_id_null,json=applicationIdNull" json:"application_id_null,omitempty"`
	UserId                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull                bool                        `protobuf:"varint,1002,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	KeyVariable               *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull           bool                        `protobuf:"varint,1003,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	Value                     *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	ValueNull                 bool                        `protobuf:"varint,1004,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
	Delete                    *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	DeleteNull                bool                        `protobuf:"varint,1005,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
	KeyValuePairsInTempdb     *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=key_value_pairs_in_tempdb,json=keyValuePairsInTempdb" json:"key_value_pairs_in_tempdb,omitempty"`
	KeyValuePairsInTempdbNull bool                        `protobuf:"varint,1006,opt,name=key_value_pairs_in_tempdb_null,json=keyValuePairsInTempdbNull" json:"key_value_pairs_in_tempdb_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplicationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
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

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Parameters) GetKeyValuePairsInTempdb() *dstore_values.BooleanValue {
	if m != nil {
		return m.KeyValuePairsInTempdb
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyAppSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyAppSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyAppSettings_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyAppSettings_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x1b, 0x92, 0x94, 0x49, 0xa1, 0x74, 0x2b, 0xd0, 0x26, 0x11, 0x51, 0x29, 0x42, 0x42,
	0x42, 0xda, 0x20, 0x82, 0x10, 0x1c, 0x38, 0xb4, 0x52, 0x0f, 0x39, 0xa4, 0xaa, 0x16, 0xa8, 0x80,
	0xcb, 0xca, 0xa9, 0xa7, 0x91, 0xd5, 0x8d, 0xbd, 0xb2, 0x37, 0xad, 0xfa, 0x0a, 0x9c, 0x78, 0x1b,
	0x9e, 0x89, 0xdf, 0x67, 0xc0, 0xf6, 0x38, 0x34, 0x09, 0xad, 0x02, 0x97, 0xc4, 0xe3, 0xf9, 0x7e,
	0xec, 0xd9, 0x19, 0xc3, 0x0b, 0x6e, 0x4a, 0xa5, 0xb1, 0x8b, 0x72, 0x24, 0x24, 0x76, 0x0b, 0xad,
	0x8e, 0x91, 0x4f, 0x34, 0x9a, 0xee, 0x58, 0x64, 0x03, 0xc5, 0xc5, 0xc9, 0xc5, 0x6e, 0x51, 0xbc,
	0xc1, 0xb2, 0x14, 0x72, 0x64, 0xb2, 0x5d, 0x9e, 0x58, 0x48, 0xa9, 0xa2, 0x47, 0xc4, 0x4b, 0x88,
	0x97, 0x5c, 0x03, 0x6e, 0x6d, 0x05, 0xf9, 0x33, 0x96, 0x4f, 0xd0, 0x10, 0xb7, 0xd5, 0x9c, 0xf7,
	0x44, 0xad, 0x95, 0x0e, 0xa9, 0xf6, 0x7c, 0x6a, 0x8c, 0xc6, 0xb0, 0x11, 0x86, 0xe4, 0xc3, 0xc5,
	0x64, 0xc9, 0x84, 0x3c, 0x51, 0x7a, 0xcc, 0x4a, 0xa1, 0x24, 0x81, 0x76, 0x3e, 0x55, 0x01, 0x0e,
	0x99, 0x66, 0x36, 0x8b, 0xda, 0x44, 0x7b, 0x70, 0x9b, 0x15, 0x45, 0x2e, 0x8e, 0x3d, 0x26, 0x13,
	0x3c, 0x5e, 0xd9, 0x5e, 0x79, 0xdc, 0x78, 0xd6, 0x4e, 0xc2, 0x05, 0xc2, 0xc9, 0x84, 0x2c, 0x71,
	0x84, 0xfa, 0xc8, 0x45, 0xe9, 0xad, 0x19, 0x4a, 0x9f, 0x47, 0x5d, 0xd8, 0x9a, 0xd7, 0xc8, 0xe4,
	0x24, 0xcf, 0xe3, 0xaf, 0x75, 0xab, 0xb4, 0x96, 0x6e, 0xce, 0x81, 0x0f, 0x6c, 0x26, 0x7a, 0x0e,
	0xf5, 0x89, 0x41, 0xed, 0xdc, 0x56, 0x97, 0xbb, 0xd5, 0x1c, 0xd6, 0xda, 0x3c, 0x80, 0xf5, 0xc0,
	0x22, 0xfd, 0x6f, 0xa4, 0x0f, 0x94, 0xf6, 0xc2, 0xaf, 0x61, 0xfd, 0x14, 0x2f, 0xb2, 0x33, 0xa6,
	0x05, 0x1b, 0xe6, 0x18, 0x57, 0xbc, 0x7a, 0x6b, 0x41, 0xdd, 0x94, 0xda, 0xd6, 0x9f, 0xc4, 0x1b,
	0x16, 0x7f, 0x14, 0xe0, 0xd1, 0x13, 0xd8, 0x9c, 0xa5, 0x93, 0xcd, 0x77, 0xb2, 0xd9, 0x98, 0x01,
	0x7a, 0xaf, 0xa7, 0x50, 0xf5, 0x7a, 0xf1, 0x8d, 0xa5, 0x26, 0x04, 0x8c, 0x3a, 0x00, 0x7e, 0x41,
	0xba, 0x3f, 0x48, 0xf7, 0xa6, 0xdf, 0xf2, 0x8a, 0x3d, 0xa8, 0x71, 0xcc, 0xed, 0x77, 0x89, 0xab,
	0x57, 0x56, 0x65, 0xa8, 0x54, 0x8e, 0x4c, 0x86, 0xaa, 0x10, 0x34, 0xda, 0x86, 0x06, 0xad, 0x48,
	0xf5, 0x67, 0x28, 0x0a, 0xed, 0x79, 0xd9, 0x77, 0xd0, 0xa4, 0x5b, 0x39, 0xeb, 0x82, 0x09, 0x6d,
	0x32, 0x21, 0xb3, 0x12, 0xc7, 0x05, 0x1f, 0xc6, 0xb5, 0xe5, 0x4e, 0x77, 0xfd, 0xcd, 0xed, 0xea,
	0xd0, 0x71, 0xfb, 0xf2, 0xad, 0x67, 0xda, 0xce, 0xe9, 0x5c, 0x2b, 0x4b, 0x67, 0xf9, 0x45, 0x67,
	0x69, 0x5e, 0xc9, 0x77, 0x47, 0xdb, 0xf9, 0xb2, 0x0a, 0x6b, 0x29, 0x9a, 0x42, 0x49, 0x83, 0xae,
	0xa0, 0xbe, 0xd5, 0x43, 0x07, 0xfe, 0x29, 0x68, 0x18, 0x21, 0x1a, 0x83, 0x7d, 0xf7, 0x9b, 0x12,
	0x30, 0xfa, 0x00, 0x77, 0x5c, 0x93, 0x67, 0x33, 0x5d, 0x6e, 0x1b, 0xaa, 0x62, 0xc9, 0xc9, 0x02,
	0x79, 0x71, 0x16, 0x06, 0x36, 0xee, 0x5f, 0xc6, 0xe9, 0xc6, 0x78, 0x7e, 0x23, 0x7a, 0x09, 0xf5,
	0x30, 0x5c, 0xb6, 0x89, 0x9c, 0x62, 0xe7, 0x2f, 0x45, 0x1a, 0xbd, 0x01, 0xfd, 0xa7, 0x53, 0x78,
	0xb4, 0x0f, 0x15, 0xad, 0xce, 0x6d, 0x57, 0x38, 0x56, 0x2f, 0xf9, 0xa7, 0x77, 0x20, 0x99, 0x16,
	0x21, 0x49, 0xd5, 0x79, 0xea, 0xf8, 0xad, 0xfb, 0x50, 0xb1, 0xeb, 0xe8, 0x1e, 0xd4, 0x6c, 0xe4,
	0x26, 0xe5, 0xf3, 0x81, 0x2d, 0x4b, 0x35, 0xad, 0xda, 0xb0, 0xcf, 0xf7, 0xde, 0x43, 0x5b, 0xa8,
	0x05, 0xf1, 0xcb, 0xc7, 0xe9, 0xe3, 0xab, 0x91, 0x32, 0xfc, 0x74, 0x9a, 0xe7, 0xff, 0xf1, 0x7e,
	0x0d, 0x6b, 0xfe, 0x9d, 0xe8, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x85, 0x12, 0x1d, 0x14, 0xfa,
	0x04, 0x00, 0x00,
}
