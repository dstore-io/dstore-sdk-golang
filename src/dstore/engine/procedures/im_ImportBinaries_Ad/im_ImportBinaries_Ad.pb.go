// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ImportBinaries_Ad.proto
// DO NOT EDIT!

/*
Package im_ImportBinaries_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ImportBinaries_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ImportBinaries_Ad

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
	ValueOrItemBinaries     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=value_or_item_binaries,json=valueOrItemBinaries" json:"value_or_item_binaries,omitempty"`
	ValueOrItemBinariesNull bool                        `protobuf:"varint,1001,opt,name=value_or_item_binaries_null,json=valueOrItemBinariesNull" json:"value_or_item_binaries_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetValueOrItemBinaries() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueOrItemBinaries
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ImportBinaries_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ImportBinaries_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ImportBinaries_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xa5, 0xcd, 0xd7, 0x1f, 0xe6, 0x5b, 0x28, 0x53, 0xa8, 0x31, 0x41, 0x29, 0x75, 0xe3, 0xc6,
	0x69, 0xb1, 0x1b, 0x37, 0x2e, 0x2c, 0xb8, 0x08, 0xd8, 0x5a, 0x66, 0x21, 0xe8, 0x66, 0x48, 0xcd,
	0xb5, 0x0c, 0x34, 0x33, 0x65, 0x26, 0xb5, 0xaf, 0xe1, 0x0b, 0xf8, 0x16, 0xbe, 0x94, 0x6f, 0xe1,
	0x24, 0x33, 0xa5, 0x26, 0x16, 0xc4, 0x4d, 0x92, 0x3b, 0xe7, 0x9e, 0x33, 0xf7, 0x9c, 0x5c, 0x34,
	0x4a, 0x74, 0x26, 0x15, 0x0c, 0x40, 0x2c, 0xb8, 0x80, 0xc1, 0x4a, 0xc9, 0x67, 0x48, 0xd6, 0x0a,
	0xf4, 0x80, 0xa7, 0x2c, 0x4a, 0x57, 0x52, 0x65, 0x63, 0x2e, 0x62, 0xc5, 0x41, 0xb3, 0x9b, 0x84,
	0x18, 0x3c, 0x93, 0xb8, 0x6f, 0x49, 0xc4, 0x92, 0xc8, 0xbe, 0xce, 0xa0, 0xe3, 0x84, 0x5f, 0xe3,
	0xe5, 0x1a, 0xb4, 0x25, 0x06, 0xc7, 0xe5, 0xdb, 0x40, 0x29, 0xa9, 0x1c, 0x14, 0x96, 0xa1, 0x14,
	0xb4, 0x8e, 0x17, 0xe0, 0xc0, 0xb3, 0x2a, 0x98, 0xc5, 0x5c, 0xbc, 0x48, 0x95, 0xc6, 0x19, 0x97,
	0xc2, 0x36, 0xf5, 0xdf, 0x6b, 0x08, 0xcd, 0x62, 0x15, 0x1b, 0x14, 0x94, 0xc6, 0x33, 0xd4, 0x2d,
	0xee, 0x66, 0x52, 0x31, 0x9e, 0x41, 0xca, 0xe6, 0x6e, 0x3a, 0xbf, 0xd6, 0xab, 0x9d, 0xff, 0xbf,
	0x0c, 0x89, 0x73, 0xe1, 0x26, 0xe4, 0x22, 0x83, 0x05, 0xa8, 0x87, 0xbc, 0xa2, 0x9d, 0xe2, 0xf0,
	0x5e, 0x45, 0x86, 0xb8, 0x75, 0x85, 0xaf, 0x51, 0xb8, 0x5f, 0x91, 0x89, 0xf5, 0x72, 0xe9, 0x7f,
	0xb6, 0x8c, 0x6e, 0x9b, 0x1e, 0xed, 0xa1, 0x4e, 0x0d, 0xde, 0xff, 0xa8, 0xa3, 0x36, 0x05, 0xbd,
	0x92, 0x42, 0x03, 0x1e, 0xa2, 0x46, 0xe1, 0xde, 0x0d, 0x13, 0x90, 0x72, 0xa4, 0x36, 0x99, 0xdb,
	0xfc, 0x49, 0x6d, 0x23, 0x7e, 0x44, 0x87, 0xb9, 0x6f, 0xf6, 0xcd, 0xb8, 0x5f, 0xef, 0x79, 0x86,
	0x4c, 0x2a, 0xe4, 0x6a, 0x3c, 0x13, 0x53, 0x47, 0xbb, 0x9a, 0x1e, 0xa4, 0xe5, 0x03, 0x7c, 0x85,
	0x5a, 0x2e, 0x6f, 0xdf, 0x2b, 0x14, 0x4f, 0x7f, 0x28, 0xda, 0xbf, 0x31, 0xb1, 0x6f, 0xba, 0x6d,
	0xc7, 0x63, 0xe4, 0x29, 0xb9, 0xf1, 0xff, 0x15, 0xac, 0x21, 0xf9, 0x7d, 0x2f, 0xc8, 0x36, 0x01,
	0x42, 0xe5, 0x86, 0xe6, 0xe4, 0xe0, 0x04, 0x79, 0xe6, 0x1b, 0x77, 0x51, 0xd3, 0x54, 0x8c, 0x27,
	0xfe, 0xdb, 0xd4, 0x64, 0xd2, 0xa0, 0x0d, 0x53, 0x46, 0xc9, 0xf8, 0x0e, 0x85, 0x5c, 0x56, 0x94,
	0x77, 0x6b, 0xfa, 0x74, 0xf1, 0xa7, 0x05, 0x9e, 0x37, 0x8b, 0x5d, 0x19, 0x7d, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xca, 0xf0, 0x8c, 0xed, 0xf8, 0x02, 0x00, 0x00,
}
