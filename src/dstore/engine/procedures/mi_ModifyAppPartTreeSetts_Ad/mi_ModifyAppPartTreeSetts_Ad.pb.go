// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyAppPartTreeSetts_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyAppPartTreeSetts_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyAppPartTreeSetts_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyAppPartTreeSetts_Ad

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
	ApplicationPartTreeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=application_part_tree_id,json=applicationPartTreeId" json:"application_part_tree_id,omitempty"`
	ApplicationPartTreeIdNull bool                        `protobuf:"varint,1001,opt,name=application_part_tree_id_null,json=applicationPartTreeIdNull" json:"application_part_tree_id_null,omitempty"`
	UserId                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull                bool                        `protobuf:"varint,1002,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	KeyVariable               *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull           bool                        `protobuf:"varint,1003,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	Value                     *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	ValueNull                 bool                        `protobuf:"varint,1004,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
	Delete                    *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	DeleteNull                bool                        `protobuf:"varint,1005,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
	DataInTempdb              *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=data_in_tempdb,json=dataInTempdb" json:"data_in_tempdb,omitempty"`
	DataInTempdbNull          bool                        `protobuf:"varint,1006,opt,name=data_in_tempdb_null,json=dataInTempdbNull" json:"data_in_tempdb_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplicationPartTreeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationPartTreeId
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

func (m *Parameters) GetDataInTempdb() *dstore_values.BooleanValue {
	if m != nil {
		return m.DataInTempdb
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyAppPartTreeSetts_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyAppPartTreeSetts_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyAppPartTreeSetts_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0x13, 0x3f,
	0x10, 0x57, 0x9b, 0x7f, 0xd2, 0xfe, 0xa7, 0x11, 0x2d, 0xae, 0x40, 0xdb, 0x44, 0xad, 0x4a, 0xb9,
	0x20, 0x90, 0x36, 0x88, 0x82, 0x54, 0x09, 0x71, 0x08, 0x12, 0x87, 0x20, 0xa5, 0x42, 0x4b, 0x54,
	0x09, 0x2e, 0x2b, 0xa7, 0x3b, 0x8d, 0x2c, 0x36, 0xf6, 0xca, 0x76, 0x5a, 0xf5, 0x05, 0x38, 0xf3,
	0x40, 0xbc, 0x10, 0x9f, 0xcf, 0x80, 0xed, 0x71, 0x4a, 0x12, 0x0a, 0x41, 0x5c, 0x12, 0xdb, 0xf3,
	0xfb, 0x98, 0x1d, 0xcf, 0x18, 0x9e, 0x16, 0xc6, 0x2a, 0x8d, 0x1d, 0x94, 0x23, 0x21, 0xb1, 0x53,
	0x69, 0x75, 0x8a, 0xc5, 0x44, 0xa3, 0xe9, 0x8c, 0x45, 0xde, 0x57, 0x85, 0x38, 0xbb, 0xec, 0x56,
	0xd5, 0x2b, 0xae, 0xed, 0x40, 0x23, 0xbe, 0x46, 0x6b, 0x4d, 0xde, 0x2d, 0x52, 0x87, 0xb3, 0x8a,
	0xdd, 0x27, 0x72, 0x4a, 0xe4, 0xf4, 0x4f, 0x8c, 0xd6, 0x76, 0x34, 0x3a, 0xe7, 0xe5, 0x04, 0x0d,
	0x09, 0xb4, 0x76, 0xe6, 0xdd, 0x51, 0x6b, 0xa5, 0x63, 0xa8, 0x3d, 0x1f, 0x1a, 0xa3, 0x31, 0x7c,
	0x84, 0x31, 0x78, 0x77, 0x31, 0x68, 0xb9, 0x90, 0x67, 0x4a, 0x8f, 0xb9, 0x15, 0x4a, 0x12, 0xe8,
	0xe0, 0x7d, 0x1d, 0xc0, 0xa5, 0xc1, 0x5d, 0x14, 0xb5, 0x61, 0x03, 0x48, 0x78, 0x55, 0x95, 0xe2,
	0x34, 0x60, 0xf2, 0xca, 0x25, 0x98, 0x5b, 0x97, 0x61, 0x2e, 0x8a, 0x64, 0x65, 0x7f, 0xe5, 0xde,
	0xc6, 0xa3, 0x76, 0x1a, 0xbf, 0x27, 0xe6, 0x28, 0xa4, 0xc5, 0x11, 0xea, 0x13, 0xbf, 0xcb, 0x6e,
	0xcd, 0x90, 0xa7, 0x1f, 0xd7, 0x2b, 0x58, 0x17, 0x76, 0x7f, 0xa7, 0x9a, 0xcb, 0x49, 0x59, 0x26,
	0x9f, 0xd6, 0x9c, 0xf6, 0x7a, 0xb6, 0x73, 0x2d, 0xfd, 0xd8, 0x21, 0xd8, 0x63, 0x58, 0x9b, 0x18,
	0xd4, 0x3e, 0x8f, 0xd5, 0xe5, 0x79, 0x34, 0x3c, 0xd6, 0x19, 0xdf, 0x81, 0x66, 0x64, 0x91, 0xcf,
	0x67, 0xf2, 0x01, 0x0a, 0x07, 0xe1, 0x67, 0xd0, 0x7c, 0x87, 0x97, 0xf9, 0x39, 0xd7, 0x82, 0x0f,
	0x4b, 0x4c, 0x6a, 0x41, 0xbd, 0xb5, 0xa0, 0x6e, 0xac, 0x16, 0x72, 0x44, 0xe2, 0x1b, 0x0e, 0x7f,
	0x12, 0xe1, 0xec, 0x01, 0xdc, 0x9c, 0xa5, 0x93, 0xcd, 0x17, 0xb2, 0xd9, 0x9c, 0x01, 0x06, 0xaf,
	0x87, 0x50, 0x0f, 0x7a, 0xc9, 0x7f, 0x4b, 0x4d, 0x08, 0xc8, 0xf6, 0x00, 0xc2, 0x82, 0x74, 0xbf,
	0x92, 0xee, 0xff, 0xe1, 0x28, 0x28, 0x1e, 0x42, 0xa3, 0xc0, 0xd2, 0xdd, 0x5d, 0x52, 0xbf, 0xb6,
	0x2a, 0x43, 0xa5, 0x4a, 0xe4, 0x32, 0x56, 0x85, 0xa0, 0x6c, 0x1f, 0x36, 0x68, 0x45, 0xaa, 0xdf,
	0x62, 0x51, 0xe8, 0x2c, 0xc8, 0x76, 0xe1, 0x46, 0xc1, 0x2d, 0xcf, 0x85, 0xcc, 0x2d, 0x8e, 0xab,
	0x62, 0x98, 0x34, 0x96, 0xcb, 0x37, 0x3d, 0xa5, 0x27, 0x07, 0x81, 0xc0, 0x52, 0xd8, 0x9e, 0x97,
	0x20, 0xb3, 0xef, 0x64, 0xb6, 0x35, 0x8b, 0xf5, 0x96, 0x07, 0x1f, 0x57, 0x61, 0x3d, 0x43, 0x53,
	0x29, 0x69, 0xd0, 0x17, 0x2a, 0xb4, 0x79, 0xec, 0xb9, 0xab, 0x42, 0xc5, 0x19, 0xa2, 0x11, 0x78,
	0xe1, 0x7f, 0x33, 0x02, 0xb2, 0x37, 0xb0, 0xe5, 0x1b, 0x3c, 0x9f, 0xe9, 0x70, 0xd7, 0x28, 0x35,
	0x47, 0x4e, 0x17, 0xc8, 0x8b, 0x73, 0xd0, 0x47, 0x9f, 0xc9, 0xd5, 0x3e, 0xdb, 0x1c, 0xcf, 0x1f,
	0xb0, 0x23, 0x58, 0x8b, 0x83, 0xe5, 0x9a, 0xc3, 0x2b, 0xee, 0xfd, 0xa2, 0x48, 0x63, 0xd7, 0xa7,
	0xff, 0x6c, 0x0a, 0x67, 0x2f, 0xa1, 0xa6, 0xd5, 0x85, 0xbb, 0x6d, 0xcf, 0x3a, 0x4a, 0xff, 0xfe,
	0x21, 0x48, 0xa7, 0x95, 0x48, 0x33, 0x75, 0x91, 0x79, 0x91, 0xd6, 0x2e, 0xd4, 0xdc, 0x9a, 0xdd,
	0x86, 0x86, 0xdb, 0xf9, 0x31, 0xf8, 0x70, 0xec, 0x6a, 0x53, 0xcf, 0xea, 0x6e, 0xdb, 0x2b, 0x9e,
	0x0f, 0xa0, 0x2d, 0xd4, 0x82, 0xc3, 0xcf, 0x77, 0xea, 0xed, 0x93, 0x7f, 0x7a, 0xc1, 0x86, 0x8d,
	0xf0, 0x48, 0x1c, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37, 0xe6, 0x98, 0xa4, 0x01, 0x05, 0x00,
	0x00,
}
