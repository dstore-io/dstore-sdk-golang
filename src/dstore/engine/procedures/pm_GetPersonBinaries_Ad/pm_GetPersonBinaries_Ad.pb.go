// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonBinaries_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersonBinaries_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonBinaries_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonBinaries_Ad

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
	PersonId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull               bool                        `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	IncludeBinaryCode          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=include_binary_code,json=includeBinaryCode" json:"include_binary_code,omitempty"`
	IncludeBinaryCodeNull      bool                        `protobuf:"varint,1002,opt,name=include_binary_code_null,json=includeBinaryCodeNull" json:"include_binary_code_null,omitempty"`
	BinaryCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=binary_characteristic_id,json=binaryCharacteristicId" json:"binary_characteristic_id,omitempty"`
	BinaryCharacteristicIdNull bool                        `protobuf:"varint,1003,opt,name=binary_characteristic_id_null,json=binaryCharacteristicIdNull" json:"binary_characteristic_id_null,omitempty"`
	Value                      *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	ValueNull                  bool                        `protobuf:"varint,1004,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetIncludeBinaryCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.IncludeBinaryCode
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicId
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PersonId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	BinaryDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=binary_description,json=binaryDescription" json:"binary_description,omitempty"`
	BinaryCodeId      *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	ThumbnailCode     *dstore_values.BytesValue   `protobuf:"bytes,10004,opt,name=thumbnail_code,json=thumbnailCode" json:"thumbnail_code,omitempty"`
	BinaryCode        *dstore_values.BytesValue   `protobuf:"bytes,10005,opt,name=binary_code,json=binaryCode" json:"binary_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetBinaryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryDescription
	}
	return nil
}

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Response_Row) GetThumbnailCode() *dstore_values.BytesValue {
	if m != nil {
		return m.ThumbnailCode
	}
	return nil
}

func (m *Response_Row) GetBinaryCode() *dstore_values.BytesValue {
	if m != nil {
		return m.BinaryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonBinaries_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonBinaries_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonBinaries_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x15, 0xe4, 0x0b, 0x84, 0x81, 0x2f, 0x2d, 0x46, 0x45, 0x26, 0x51, 0x51, 0x45, 0x85, 0xd4,
	0x2b, 0x07, 0x15, 0xa9, 0x6a, 0xd5, 0x2b, 0x42, 0x51, 0x45, 0x2b, 0x52, 0x64, 0xa9, 0x95, 0xda,
	0x1b, 0xcb, 0xb1, 0xa7, 0x61, 0xa5, 0x64, 0x37, 0xda, 0x75, 0x8a, 0x72, 0xdd, 0x17, 0xe8, 0xef,
	0x03, 0xf6, 0xae, 0x7f, 0x0f, 0xd1, 0xd9, 0x1d, 0xe7, 0xcf, 0x80, 0x02, 0x37, 0x90, 0xf1, 0xce,
	0x39, 0x73, 0x76, 0xe6, 0xcc, 0xc2, 0xa3, 0xd4, 0x64, 0x4a, 0x63, 0x03, 0x65, 0x47, 0x48, 0x6c,
	0xf4, 0xb5, 0x4a, 0x30, 0x1d, 0x68, 0x34, 0x8d, 0x7e, 0x2f, 0x7a, 0x8e, 0xd9, 0x29, 0x6a, 0xa3,
	0x64, 0x53, 0xc8, 0x58, 0x0b, 0x34, 0xd1, 0x41, 0x1a, 0x50, 0x4a, 0xa6, 0xbc, 0x5d, 0xc6, 0x05,
	0x8c, 0x0b, 0xae, 0x48, 0xae, 0x6d, 0xe4, 0xf4, 0x1f, 0xe2, 0xee, 0x00, 0x0d, 0x63, 0x6b, 0x5b,
	0xb3, 0x35, 0x51, 0x6b, 0xa5, 0xf3, 0xa3, 0xfa, 0xec, 0x51, 0x0f, 0x8d, 0x89, 0x3b, 0x98, 0x1f,
	0xde, 0x2f, 0x1e, 0x66, 0xb1, 0x90, 0xef, 0x95, 0xee, 0xc5, 0x99, 0x50, 0x92, 0x93, 0x76, 0x7e,
	0x94, 0x00, 0x4e, 0x63, 0x1d, 0xd3, 0x29, 0xa9, 0xf1, 0x1e, 0xc3, 0x4a, 0xdf, 0xa9, 0x8a, 0x44,
	0xea, 0x2f, 0xdc, 0x5b, 0x78, 0xb0, 0xfa, 0xb0, 0x1e, 0xe4, 0xda, 0x73, 0x51, 0x42, 0x66, 0xd8,
	0x41, 0xfd, 0xc6, 0x46, 0x61, 0x85, 0xb3, 0x8f, 0x53, 0x6f, 0x17, 0xaa, 0x63, 0x64, 0x24, 0x07,
	0xdd, 0xae, 0xff, 0x6b, 0x99, 0xf0, 0x95, 0x70, 0x6d, 0x94, 0xd2, 0xa2, 0x8f, 0xde, 0x4b, 0xd8,
	0x10, 0x32, 0xe9, 0x0e, 0x52, 0x8c, 0xda, 0xf6, 0xe2, 0xc3, 0x28, 0x51, 0x29, 0xfa, 0x8b, 0xf3,
	0x4b, 0xad, 0xe7, 0x38, 0xd7, 0xaf, 0xe1, 0x21, 0xa1, 0x48, 0xad, 0x7f, 0x09, 0x19, 0x57, 0xff,
	0xcd, 0xd5, 0xef, 0x5c, 0x40, 0x39, 0x19, 0xaf, 0xc1, 0x1f, 0x21, 0xce, 0xe8, 0xf6, 0x09, 0x5d,
	0x5e, 0x98, 0x4c, 0x24, 0xf6, 0xda, 0xa5, 0xf9, 0x5a, 0x36, 0x19, 0x7c, 0x38, 0x83, 0xa5, 0x26,
	0x34, 0xe1, 0xee, 0x55, 0xb4, 0xac, 0xea, 0x0f, 0xab, 0xaa, 0x5d, 0x8e, 0x77, 0xd2, 0xf6, 0xa0,
	0xec, 0x4a, 0xfa, 0xff, 0x39, 0x1d, 0xb5, 0x82, 0x0e, 0x93, 0x69, 0x21, 0x3b, 0x2c, 0x83, 0x13,
	0xbd, 0x6d, 0x00, 0xf7, 0x83, 0x4b, 0xfc, 0xe5, 0x12, 0x2b, 0xee, 0x93, 0x65, 0xdc, 0xf9, 0x58,
	0x86, 0x4a, 0x88, 0xa6, 0xaf, 0xa4, 0x41, 0x4b, 0xef, 0x1c, 0x94, 0x4f, 0x77, 0x4c, 0x9f, 0x3b,
	0x93, 0xdd, 0x75, 0x64, 0xff, 0x86, 0x9c, 0xe8, 0xbd, 0x85, 0xdb, 0xd6, 0x3b, 0xd1, 0x94, 0x79,
	0x68, 0x5e, 0x25, 0x02, 0x07, 0x05, 0x70, 0xd1, 0x62, 0x27, 0x14, 0x1f, 0x4f, 0xe2, 0xf0, 0x56,
	0x6f, 0xf6, 0x03, 0x0d, 0x70, 0x39, 0xf7, 0x2c, 0x75, 0xdd, 0x32, 0x6e, 0x5f, 0x60, 0x64, 0x47,
	0x9f, 0xf0, 0xff, 0x70, 0x94, 0xee, 0x1d, 0x41, 0x49, 0xab, 0x73, 0xea, 0x91, 0x45, 0xed, 0x07,
	0xd7, 0x5a, 0xaf, 0x60, 0xd4, 0x84, 0x20, 0x54, 0xe7, 0xa1, 0xc5, 0xd7, 0x7e, 0x2e, 0x42, 0x89,
	0x02, 0x6f, 0x13, 0x96, 0x28, 0xb4, 0xd3, 0xff, 0xd4, 0xa2, 0xbe, 0x94, 0xc3, 0x32, 0x85, 0x34,
	0xd0, 0x27, 0xd3, 0xfb, 0xf0, 0xb9, 0x75, 0x93, 0x85, 0x78, 0x01, 0x5e, 0xee, 0x85, 0x14, 0x4d,
	0xa2, 0x45, 0xdf, 0x35, 0xee, 0x4b, 0x6b, 0xee, 0x54, 0xd7, 0x19, 0xf6, 0x6c, 0x82, 0x22, 0x5f,
	0x55, 0xa7, 0x0d, 0x4e, 0x5a, 0xbe, 0x5e, 0x43, 0xcb, 0x5a, 0x7b, 0x6c, 0x7a, 0xd2, 0x73, 0x00,
	0xd5, 0xec, 0x6c, 0xd0, 0x6b, 0xcb, 0x58, 0x74, 0x79, 0xe9, 0xbe, 0x31, 0xc7, 0x56, 0x81, 0xa3,
	0x3d, 0xcc, 0xd0, 0x30, 0xc3, 0xff, 0x63, 0x84, 0xdb, 0xb7, 0xa7, 0xb0, 0x3a, 0xbd, 0xb4, 0xdf,
	0xe7, 0xe2, 0x61, 0xa2, 0xa0, 0xf9, 0x0a, 0xea, 0x42, 0x15, 0x07, 0x35, 0x7e, 0x3f, 0xdf, 0xed,
	0xdd, 0xf4, 0x65, 0x6d, 0x2f, 0xb9, 0x17, 0x6c, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0x01, 0x88, 0xf9, 0x94, 0x05, 0x00, 0x00,
}