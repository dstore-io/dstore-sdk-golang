// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinaryPredefinedVals_Pu.proto
// DO NOT EDIT!

/*
Package im_GetBinaryPredefinedVals_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinaryPredefinedVals_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinaryPredefinedVals_Pu

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
	BinaryCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=binary_characteristic_id,json=binaryCharacteristicId" json:"binary_characteristic_id,omitempty"`
	BinaryCharacteristicIdNull bool                        `protobuf:"varint,1001,opt,name=binary_characteristic_id_null,json=binaryCharacteristicIdNull" json:"binary_characteristic_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBinaryCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicId
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
	RowId   int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value" json:"value,omitempty"`
	ValueId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
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

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinaryPredefinedVals_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinaryPredefinedVals_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinaryPredefinedVals_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetBinaryPredefinedVals_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd4, 0x30,
	0x18, 0xa5, 0xd6, 0xee, 0x0c, 0xf1, 0x42, 0x89, 0xb0, 0xd4, 0x0e, 0xca, 0xb2, 0xde, 0x08, 0x42,
	0xc6, 0x9f, 0x0b, 0x15, 0x04, 0x61, 0x44, 0x64, 0x90, 0x2d, 0x43, 0xc0, 0x05, 0xbd, 0xb0, 0x64,
	0x9b, 0x6f, 0x6b, 0xb0, 0x4d, 0x96, 0x24, 0x75, 0xf1, 0x11, 0xbc, 0x53, 0x5f, 0x42, 0xf0, 0xcd,
	0x7c, 0x0b, 0xd3, 0x24, 0xc3, 0xda, 0xfa, 0xb3, 0xec, 0xcd, 0x4c, 0xbf, 0xef, 0x3b, 0xe7, 0x24,
	0x39, 0x27, 0x41, 0x4f, 0xb9, 0xb1, 0x4a, 0xc3, 0x12, 0x64, 0x23, 0x24, 0x2c, 0x4f, 0xb4, 0xaa,
	0x81, 0xf7, 0x1a, 0xcc, 0x52, 0x74, 0xd5, 0x4b, 0xb0, 0x2b, 0x21, 0x99, 0xfe, 0xb4, 0xd1, 0xc0,
	0xe1, 0xd8, 0x01, 0xf8, 0x21, 0x6b, 0x4d, 0xb5, 0xe9, 0x89, 0x03, 0x5a, 0x85, 0xef, 0x06, 0x36,
	0x09, 0x6c, 0xf2, 0x5f, 0x4a, 0x71, 0x3d, 0x2e, 0xf5, 0x91, 0xb5, 0x3d, 0x98, 0xa0, 0x50, 0xdc,
	0x18, 0xaf, 0x0f, 0x5a, 0x2b, 0x1d, 0x47, 0x8b, 0xf1, 0xa8, 0x03, 0x63, 0x58, 0x03, 0x71, 0x78,
	0x7b, 0x3a, 0xb4, 0x4c, 0xc8, 0x63, 0xa5, 0x3b, 0x66, 0x85, 0x92, 0x01, 0xb4, 0xff, 0x3d, 0x41,
	0x68, 0xc3, 0x34, 0x73, 0x53, 0xd0, 0x06, 0xbf, 0x46, 0xf9, 0x91, 0xdf, 0x5b, 0x55, 0xbf, 0x77,
	0xdd, 0xda, 0x35, 0x85, 0xb1, 0xa2, 0xae, 0x04, 0xcf, 0x93, 0xbd, 0xe4, 0xce, 0x95, 0x07, 0x0b,
	0x12, 0x0f, 0x14, 0xf7, 0x28, 0xa4, 0x85, 0x06, 0xf4, 0xe1, 0x50, 0xd1, 0xdd, 0x40, 0x7e, 0x3e,
	0xe2, 0xae, 0x39, 0x5e, 0xa1, 0x9b, 0xff, 0x92, 0xad, 0x64, 0xdf, 0xb6, 0xf9, 0xcf, 0x99, 0x13,
	0x9f, 0xd3, 0xe2, 0xef, 0xfc, 0xd2, 0x41, 0xf6, 0x7f, 0xa4, 0x68, 0x4e, 0xc1, 0x9c, 0x28, 0x69,
	0x00, 0xdf, 0x43, 0x99, 0xf7, 0x21, 0x6e, 0xaa, 0x20, 0x63, 0x97, 0x83, 0x47, 0x2f, 0x86, 0x5f,
	0x1a, 0x80, 0xf8, 0x0d, 0xba, 0x36, 0x38, 0x50, 0xfd, 0x66, 0x41, 0x7e, 0x69, 0x2f, 0x75, 0x64,
	0x32, 0x21, 0x4f, 0x8d, 0x3a, 0x70, 0xf5, 0xfa, 0xac, 0xa6, 0x57, 0xbb, 0x71, 0x03, 0x3f, 0x46,
	0xb3, 0xe8, 0x7c, 0x9e, 0x7a, 0xc5, 0x5b, 0x7f, 0x28, 0x86, 0x5c, 0x0e, 0xc2, 0x3f, 0xdd, 0xc2,
	0xf1, 0x2b, 0x94, 0x6a, 0x75, 0x9a, 0x5f, 0xf6, 0xac, 0x27, 0xe4, 0x02, 0x57, 0x85, 0x6c, 0xad,
	0x20, 0x54, 0x9d, 0xd2, 0x41, 0xa5, 0xf8, 0x9c, 0xa0, 0xd4, 0x15, 0x78, 0x17, 0xed, 0xb8, 0x72,
	0x48, 0xec, 0x4b, 0xe9, 0xdc, 0xc9, 0x68, 0xe6, 0x4a, 0x17, 0xc2, 0x7d, 0x94, 0xf9, 0xcc, 0xf2,
	0xaf, 0xe5, 0xd8, 0xb4, 0x98, 0xa4, 0xb1, 0x5a, 0xc8, 0x26, 0x04, 0x19, 0x90, 0xf8, 0x11, 0x9a,
	0xfb, 0x8f, 0x41, 0xec, 0x5b, 0x79, 0x7e, 0xfe, 0x33, 0xdf, 0x5c, 0xf3, 0xd5, 0x3b, 0xb4, 0x10,
	0x6a, 0x72, 0x9e, 0xb3, 0x87, 0xf3, 0xf6, 0x59, 0xa3, 0x0c, 0xff, 0xb0, 0x9d, 0xf3, 0x0b, 0xbf,
	0xad, 0xa3, 0x1d, 0x7f, 0x7b, 0x1f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x65, 0x09, 0x19, 0x07,
	0x9c, 0x03, 0x00, 0x00,
}
