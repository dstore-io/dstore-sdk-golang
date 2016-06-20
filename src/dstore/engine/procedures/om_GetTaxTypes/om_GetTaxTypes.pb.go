// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetTaxTypes.proto
// DO NOT EDIT!

/*
Package om_GetTaxTypes is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetTaxTypes.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetTaxTypes

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
	TaxTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tax_type_id,json=taxTypeId" json:"tax_type_id,omitempty"`
	TaxTypeIdNull bool                        `protobuf:"varint,1001,opt,name=tax_type_id_null,json=taxTypeIdNull" json:"tax_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTaxTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TaxTypeId
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
	RowId     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TaxType   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=tax_type,json=taxType" json:"tax_type,omitempty"`
	TaxTypeId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=tax_type_id,json=taxTypeId" json:"tax_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTaxType() *dstore_values.StringValue {
	if m != nil {
		return m.TaxType
	}
	return nil
}

func (m *Response_Row) GetTaxTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TaxTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetTaxTypes.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetTaxTypes.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetTaxTypes.Response.Row")
}

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x4b, 0xf3, 0x30,
	0x18, 0x65, 0x6f, 0xdf, 0x7d, 0xbc, 0x19, 0x2f, 0x8e, 0x08, 0x52, 0x3b, 0x91, 0x31, 0x2f, 0x1c,
	0x82, 0x99, 0x28, 0x88, 0xa0, 0x57, 0x82, 0xca, 0x2e, 0x56, 0x24, 0x0c, 0x41, 0x6f, 0x4a, 0xb4,
	0xb1, 0x14, 0xd6, 0x64, 0x24, 0x99, 0x9b, 0xff, 0xc2, 0x8f, 0xdf, 0xe0, 0x8f, 0xf3, 0x5f, 0x98,
	0x36, 0xa9, 0x5b, 0x2b, 0x88, 0xde, 0xb4, 0x3c, 0x39, 0xe7, 0x3c, 0x79, 0xce, 0x79, 0x02, 0x76,
	0x43, 0xa9, 0xb8, 0xa0, 0x7d, 0xca, 0xa2, 0x98, 0xd1, 0xfe, 0x44, 0xf0, 0x3b, 0x1a, 0x4e, 0x05,
	0x95, 0x7d, 0x9e, 0x04, 0x17, 0x54, 0x8d, 0xc8, 0x7c, 0xf4, 0x38, 0xa1, 0x12, 0x69, 0x44, 0x71,
	0xb8, 0x61, 0xe8, 0xc8, 0xd0, 0x51, 0x91, 0xe3, 0xad, 0xda, 0x66, 0x0f, 0x64, 0x3c, 0xcd, 0x25,
	0xde, 0x7a, 0xf1, 0x06, 0x2a, 0x04, 0x17, 0x16, 0x6a, 0x17, 0xa1, 0x84, 0x4a, 0x49, 0x22, 0x6a,
	0xc1, 0xad, 0x32, 0xa8, 0x48, 0xcc, 0xee, 0xb9, 0x48, 0x88, 0x8a, 0x39, 0x33, 0xa4, 0xae, 0x04,
	0xe0, 0x92, 0x08, 0xa2, 0x41, 0x2a, 0x24, 0x3c, 0x06, 0x4d, 0x45, 0xe6, 0x81, 0xd2, 0xc3, 0x04,
	0x71, 0xe8, 0x56, 0x3a, 0x95, 0x5e, 0x73, 0xbf, 0x8d, 0xec, 0xcc, 0x76, 0xaa, 0x98, 0x29, 0x1a,
	0x51, 0x71, 0x95, 0x56, 0xf8, 0x9f, 0x32, 0xb3, 0x0f, 0x42, 0xd8, 0x03, 0xad, 0x25, 0x71, 0xc0,
	0xa6, 0xe3, 0xb1, 0xfb, 0x5e, 0xd7, 0x2d, 0x1a, 0xf8, 0xff, 0x27, 0xcb, 0xd7, 0xa7, 0xdd, 0x37,
	0x07, 0x34, 0x30, 0x95, 0x13, 0xce, 0x24, 0x85, 0x7b, 0xa0, 0x9a, 0x59, 0xb2, 0xb7, 0x79, 0xa8,
	0x98, 0x90, 0xb1, 0x7b, 0x96, 0x7e, 0xb1, 0x21, 0xc2, 0x6b, 0xd0, 0x4a, 0xcd, 0x04, 0x4b, 0x6e,
	0xdc, 0x3f, 0x1d, 0x47, 0x8b, 0x51, 0x49, 0x5c, 0xf6, 0x3c, 0xd4, 0xf5, 0x60, 0x51, 0xe3, 0x95,
	0xa4, 0x78, 0x00, 0x8f, 0x40, 0xdd, 0x86, 0xe8, 0x3a, 0x59, 0xc7, 0xcd, 0x2f, 0x1d, 0x4d, 0xc4,
	0x43, 0xf3, 0xc7, 0x39, 0x1d, 0x9e, 0x00, 0x47, 0xf0, 0x99, 0xfb, 0x37, 0x53, 0xed, 0xa0, 0xef,
	0xd6, 0x8c, 0x72, 0xef, 0x08, 0xf3, 0x19, 0x4e, 0x65, 0xde, 0x6b, 0x05, 0x38, 0xba, 0x80, 0x6b,
	0xa0, 0xa6, 0xcb, 0x34, 0xfb, 0x27, 0x5f, 0xc7, 0x51, 0xc5, 0x55, 0x5d, 0xea, 0x6c, 0x0f, 0x41,
	0x23, 0xcf, 0xd6, 0x7d, 0xf6, 0x8b, 0x41, 0xd9, 0xb5, 0x48, 0x25, 0x62, 0x16, 0x99, 0xad, 0xd4,
	0x6d, 0xde, 0x7a, 0xaa, 0xc2, 0x42, 0x5f, 0xfc, 0xdf, 0x6c, 0xf4, 0xf4, 0x1c, 0xb4, 0x63, 0x5e,
	0xb2, 0xb2, 0x78, 0xe0, 0x37, 0xdb, 0x3f, 0x7c, 0xfa, 0xb7, 0xb5, 0xec, 0xad, 0x1d, 0x7c, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xff, 0x45, 0x9d, 0x63, 0x2c, 0x03, 0x00, 0x00,
}