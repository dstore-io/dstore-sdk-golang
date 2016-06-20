// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinary_Ad.proto
// DO NOT EDIT!

/*
Package im_GetBinary_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinary_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinary_Ad

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
	BinaryCodeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	BinaryCodeIdNull bool                        `protobuf:"varint,1001,opt,name=binary_code_id_null,json=binaryCodeIdNull" json:"binary_code_id_null,omitempty"`
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
	RowId      int32                     `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BinaryCode *dstore_values.BytesValue `protobuf:"bytes,10001,opt,name=binary_code,json=binaryCode" json:"binary_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBinaryCode() *dstore_values.BytesValue {
	if m != nil {
		return m.BinaryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinary_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinary_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinary_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x6b, 0xe2, 0x40,
	0x14, 0x45, 0xb3, 0x7e, 0x70, 0x5d, 0x76, 0x65, 0x84, 0x25, 0x46, 0x76, 0x11, 0xf7, 0x45, 0x58,
	0x18, 0x97, 0xf6, 0xa5, 0x50, 0xfa, 0xa0, 0xa5, 0x88, 0x0f, 0x4a, 0x99, 0x87, 0x42, 0x7d, 0x09,
	0xd1, 0xdc, 0xca, 0x80, 0x99, 0x91, 0x99, 0x58, 0xf1, 0xa9, 0x7f, 0xa1, 0xfd, 0x99, 0x85, 0xfe,
	0x88, 0x4e, 0x32, 0x11, 0x4d, 0x0a, 0xa5, 0x2f, 0x49, 0x6e, 0xce, 0x39, 0x77, 0xe6, 0x9c, 0x7b,
	0x81, 0x86, 0x3a, 0x96, 0x0a, 0x07, 0x28, 0x56, 0x5c, 0xe0, 0x60, 0xa3, 0xe4, 0x12, 0xc3, 0xad,
	0x42, 0x3d, 0xe0, 0x91, 0x3f, 0xc6, 0x78, 0xc4, 0x45, 0xa0, 0xf6, 0xfe, 0x30, 0xa4, 0x06, 0x8a,
	0x25, 0xf9, 0x6d, 0xf9, 0xd4, 0xf2, 0x69, 0x81, 0xe4, 0xb5, 0xb2, 0x76, 0x8f, 0xc1, 0x7a, 0x8b,
	0xda, 0x6a, 0xbc, 0x76, 0xfe, 0x0c, 0x54, 0x4a, 0xaa, 0x0c, 0xea, 0xe4, 0xa1, 0x08, 0xb5, 0x0e,
	0x56, 0x98, 0x81, 0x7f, 0x8b, 0x60, 0x1c, 0x70, 0xf1, 0x20, 0x55, 0x14, 0xc4, 0x5c, 0x0a, 0x4b,
	0xea, 0x3d, 0x01, 0xdc, 0x06, 0x2a, 0x30, 0x20, 0x2a, 0x4d, 0x86, 0xf0, 0x63, 0x61, 0x2f, 0xb3,
	0x94, 0x21, 0xfa, 0x3c, 0x74, 0x4b, 0xdd, 0x52, 0xbf, 0x71, 0xd6, 0xc9, 0x7c, 0xd2, 0xec, 0x62,
	0x5c, 0xc4, 0xb8, 0x42, 0x75, 0x97, 0x54, 0xec, 0xbb, 0x95, 0x5c, 0x1b, 0xc5, 0x24, 0x24, 0x14,
	0x5a, 0xf9, 0x16, 0xbe, 0xd8, 0xae, 0xd7, 0xee, 0x6b, 0xcd, 0x34, 0xaa, 0xb3, 0xe6, 0x29, 0x77,
	0x66, 0x80, 0xde, 0x5b, 0x19, 0xea, 0x0c, 0xf5, 0x46, 0x0a, 0x8d, 0xe4, 0x3f, 0x54, 0x52, 0x7b,
	0xd9, 0xb1, 0x1e, 0xcd, 0xc7, 0x65, 0xad, 0xdf, 0x24, 0x4f, 0x66, 0x89, 0xe4, 0x1e, 0x9a, 0x89,
	0x31, 0xff, 0xc4, 0x99, 0x5b, 0xee, 0x3a, 0x46, 0x4c, 0x0b, 0xe2, 0xa2, 0xff, 0xa9, 0xa9, 0x27,
	0xc7, 0x9a, 0xfd, 0x8c, 0xf2, 0x3f, 0xc8, 0x05, 0xd4, 0xb2, 0x40, 0x5d, 0x27, 0xed, 0xf8, 0xe7,
	0x43, 0x47, 0x1b, 0xf7, 0xd4, 0xbe, 0xd9, 0x81, 0x4e, 0xae, 0xc0, 0x51, 0x72, 0xe7, 0x7e, 0x4b,
	0x55, 0xff, 0xe8, 0xa7, 0x33, 0xa7, 0x07, 0xf3, 0x94, 0xc9, 0x1d, 0x4b, 0x74, 0xde, 0x1c, 0x1c,
	0xf3, 0x4d, 0x7e, 0x41, 0xd5, 0x54, 0xc9, 0x10, 0x9e, 0x67, 0x26, 0x8e, 0x0a, 0xab, 0x98, 0xd2,
	0x24, 0x7c, 0x09, 0x8d, 0x93, 0x84, 0xdd, 0x97, 0x59, 0x9a, 0x55, 0xbb, 0x30, 0xa2, 0xc5, 0x3e,
	0x46, 0x6d, 0x07, 0x04, 0xc7, 0xd0, 0x47, 0x63, 0xe8, 0x70, 0x59, 0xb8, 0xd1, 0x71, 0x6b, 0xe7,
	0xfd, 0xaf, 0xee, 0xf3, 0xa2, 0x9a, 0xee, 0xcf, 0xf9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0xb4, 0x5c, 0x76, 0x02, 0x03, 0x00, 0x00,
}
