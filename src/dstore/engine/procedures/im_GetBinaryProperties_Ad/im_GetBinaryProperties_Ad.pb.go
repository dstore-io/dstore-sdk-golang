// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinaryProperties_Ad.proto
// DO NOT EDIT!

/*
Package im_GetBinaryProperties_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinaryProperties_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinaryProperties_Ad

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
	BinaryCodeId                *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	BinaryCodeIdNull            bool                        `protobuf:"varint,1001,opt,name=binary_code_id_null,json=binaryCodeIdNull" json:"binary_code_id_null,omitempty"`
	BinaryCharacteristicIds     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=binary_characteristic_ids,json=binaryCharacteristicIds" json:"binary_characteristic_ids,omitempty"`
	BinaryCharacteristicIdsNull bool                        `protobuf:"varint,1002,opt,name=binary_characteristic_ids_null,json=binaryCharacteristicIdsNull" json:"binary_characteristic_ids_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicIds() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryCharacteristicIds
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CharacteristicDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	Value                     *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
	ValueId                   *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinaryProperties_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinaryProperties_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinaryProperties_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0x9a, 0xa6, 0x89, 0x06, 0x04, 0x95, 0x2b, 0xc1, 0x66, 0x23, 0x55, 0xa8, 0x1c, 0xe0,
	0xe4, 0x40, 0x11, 0xa2, 0xd7, 0x96, 0xa2, 0x2a, 0x87, 0x46, 0xc5, 0x87, 0x4a, 0xf4, 0xb2, 0xda,
	0xae, 0x87, 0x60, 0x29, 0xb1, 0x57, 0xb6, 0x43, 0xc5, 0x95, 0x5f, 0xc0, 0xc7, 0x99, 0x9f, 0xc7,
	0x01, 0x7e, 0x05, 0x5e, 0xdb, 0x21, 0xd9, 0x45, 0x51, 0xe8, 0x25, 0xc9, 0x78, 0xe6, 0xbd, 0x79,
	0x7e, 0x33, 0x0e, 0x1c, 0x71, 0x63, 0x95, 0xc6, 0x21, 0xca, 0x89, 0x90, 0x38, 0x2c, 0xb5, 0x2a,
	0x90, 0xcf, 0x35, 0x9a, 0xa1, 0x98, 0x65, 0x67, 0x68, 0x4f, 0x84, 0xcc, 0xf5, 0xa7, 0x0b, 0xad,
	0x4a, 0xd4, 0x56, 0xa0, 0xc9, 0x8e, 0x39, 0x75, 0x45, 0x56, 0x91, 0x27, 0x01, 0x49, 0x03, 0x92,
	0xae, 0x2d, 0x4f, 0xf7, 0x62, 0x8b, 0x8f, 0xf9, 0x74, 0x8e, 0x26, 0xa0, 0xd3, 0x7e, 0xbd, 0x2f,
	0x6a, 0xad, 0x74, 0x4c, 0x0d, 0xea, 0xa9, 0x19, 0x1a, 0x93, 0x4f, 0x30, 0x26, 0x1f, 0x37, 0x93,
	0x36, 0x17, 0xf2, 0xbd, 0xd2, 0xb3, 0xdc, 0x0a, 0x25, 0x43, 0xd1, 0xc1, 0x8f, 0x2d, 0x80, 0x8b,
	0x5c, 0xe7, 0x2e, 0x8b, 0xda, 0x90, 0x63, 0xb8, 0x77, 0xed, 0x75, 0x65, 0x85, 0xe2, 0x98, 0x09,
	0x9e, 0xb4, 0x1e, 0xb5, 0x9e, 0xde, 0x39, 0x1c, 0xd0, 0x78, 0x85, 0xa8, 0x4c, 0x48, 0x8b, 0x13,
	0xd4, 0x97, 0x55, 0xc4, 0xee, 0x06, 0xc8, 0x6b, 0x87, 0x18, 0x71, 0x42, 0x61, 0xaf, 0x4e, 0x91,
	0xc9, 0xf9, 0x74, 0x9a, 0xfc, 0xea, 0x3a, 0xa2, 0x1e, 0xdb, 0x5d, 0xad, 0x1d, 0xbb, 0x04, 0xb9,
	0x84, 0xfe, 0xa2, 0xfe, 0x83, 0x13, 0x52, 0x38, 0x1d, 0xc2, 0x58, 0x51, 0x38, 0xa4, 0x49, 0xb6,
	0x7c, 0xf7, 0xb4, 0xd1, 0xdd, 0x58, 0x2d, 0xe4, 0x24, 0x34, 0x7f, 0x18, 0x09, 0x6b, 0xd8, 0x11,
	0x37, 0xe4, 0x14, 0xf6, 0xd7, 0xf2, 0x06, 0x49, 0xbf, 0x83, 0xa4, 0xc1, 0x1a, 0x86, 0x4a, 0xdd,
	0xc1, 0xe7, 0x6d, 0xe8, 0x31, 0x34, 0xa5, 0x92, 0x06, 0xc9, 0x33, 0xe8, 0x78, 0xf7, 0xa3, 0x29,
	0x7f, 0x65, 0xc5, 0xb9, 0x86, 0xc9, 0xbc, 0xa9, 0x3e, 0x59, 0x28, 0x24, 0xef, 0x60, 0xb7, 0xf2,
	0x3d, 0x5b, 0x31, 0xde, 0xdd, 0xa9, 0xed, 0xc0, 0xb4, 0x01, 0x6e, 0x8e, 0xe7, 0xdc, 0xc5, 0xa3,
	0x65, 0xcc, 0xee, 0xcf, 0xea, 0x07, 0xe4, 0x08, 0xba, 0x71, 0xde, 0x49, 0xdb, 0x33, 0xee, 0xff,
	0xc3, 0x18, 0xb6, 0xe1, 0x3c, 0x7c, 0xb3, 0x45, 0x39, 0x39, 0x83, 0xb6, 0x56, 0x37, 0xc9, 0xb6,
	0x47, 0xbd, 0xa4, 0xff, 0xb9, 0x9c, 0x74, 0x61, 0x03, 0x65, 0xea, 0x86, 0x55, 0x0c, 0xe9, 0xcf,
	0x16, 0xb4, 0x5d, 0x40, 0x1e, 0xc0, 0x8e, 0x0b, 0xab, 0x6d, 0xf9, 0x32, 0x76, 0xce, 0x74, 0x58,
	0xc7, 0x85, 0x6e, 0x15, 0xae, 0x20, 0x6d, 0x78, 0xcf, 0xd1, 0x14, 0x5a, 0x94, 0xde, 0x87, 0xaf,
	0xe3, 0x8d, 0xc3, 0xed, 0xd7, 0xe1, 0xa7, 0x4b, 0x34, 0x79, 0x0e, 0x1d, 0x0f, 0x48, 0xbe, 0x6d,
	0xa6, 0x09, 0x95, 0xe4, 0x15, 0xf4, 0xfc, 0x8f, 0x4a, 0xe8, 0xf7, 0xf1, 0xe6, 0xbd, 0xee, 0xfa,
	0xc3, 0x11, 0x3f, 0x79, 0x0b, 0x03, 0xa1, 0x1a, 0x3e, 0x2d, 0x9f, 0xff, 0xd5, 0xe1, 0xed, 0xff,
	0x18, 0xae, 0x77, 0xfc, 0xf3, 0x7b, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x85, 0x57, 0x7a, 0xcb,
	0x55, 0x04, 0x00, 0x00,
}
