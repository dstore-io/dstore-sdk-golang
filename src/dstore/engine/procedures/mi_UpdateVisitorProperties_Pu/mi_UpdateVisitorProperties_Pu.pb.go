// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_UpdateVisitorProperties_Pu.proto
// DO NOT EDIT!

/*
Package mi_UpdateVisitorProperties_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_UpdateVisitorProperties_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_UpdateVisitorProperties_Pu

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
	UniqueId       *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull   bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	CountryId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryIdNull  bool                        `protobuf:"varint,1002,opt,name=country_id_null,json=countryIdNull" json:"country_id_null,omitempty"`
	LanguageId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull bool                        `protobuf:"varint,1003,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	CurrencyId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencyIdNull bool                        `protobuf:"varint,1004,opt,name=currency_id_null,json=currencyIdNull" json:"currency_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetCountryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountryId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_UpdateVisitorProperties_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_UpdateVisitorProperties_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_UpdateVisitorProperties_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0x69, 0xd7, 0xb4, 0xe9, 0xa9, 0xb6, 0x65, 0x05, 0xa9, 0x09, 0x8a, 0x54, 0x44, 0x45,
	0xd8, 0x88, 0x82, 0x5f, 0xf4, 0x4a, 0xf0, 0x22, 0x48, 0x4b, 0x18, 0x68, 0x41, 0x6f, 0x96, 0x35,
	0x7b, 0x5c, 0x06, 0x92, 0x99, 0x78, 0x66, 0xc6, 0xe2, 0x5b, 0xf8, 0x0e, 0x3e, 0x87, 0x0f, 0xe4,
	0xc7, 0x43, 0x78, 0x66, 0x67, 0xb6, 0x9b, 0xac, 0xe0, 0x47, 0x6f, 0x92, 0x9c, 0x3d, 0xff, 0xff,
	0xf9, 0x65, 0xcf, 0x07, 0x1c, 0x96, 0xc6, 0x6a, 0xc2, 0x11, 0xaa, 0x4a, 0x2a, 0x1c, 0x2d, 0x48,
	0x4f, 0xb1, 0x74, 0x84, 0x66, 0x34, 0x97, 0xf9, 0xc9, 0xa2, 0x2c, 0x2c, 0x9e, 0x4a, 0x23, 0x59,
	0x33, 0x21, 0xbd, 0x40, 0xb2, 0x12, 0x4d, 0x3e, 0x71, 0x19, 0x0b, 0xad, 0x4e, 0x1f, 0x04, 0x77,
	0x16, 0xdc, 0xd9, 0x1f, 0x2d, 0x83, 0xab, 0x11, 0xf5, 0xb1, 0x98, 0x39, 0x34, 0xa1, 0xc2, 0xe0,
	0xfa, 0x2a, 0x1f, 0x89, 0x34, 0xc5, 0xd4, 0x70, 0x35, 0x35, 0x47, 0x63, 0x8a, 0x0a, 0x63, 0xf2,
	0x76, 0x37, 0x69, 0x0b, 0xa9, 0xde, 0x6b, 0x9a, 0x17, 0x56, 0x6a, 0x15, 0x44, 0x07, 0x5f, 0x12,
	0x80, 0x49, 0x41, 0x05, 0x67, 0x91, 0x4c, 0xfa, 0x14, 0xb6, 0x9c, 0x92, 0x1f, 0x1c, 0xe6, 0xb2,
	0xdc, 0x5f, 0xbb, 0xb5, 0x76, 0x6f, 0xfb, 0xd1, 0x20, 0x8b, 0x6f, 0x10, 0xff, 0x94, 0xb1, 0x24,
	0x55, 0x75, 0xea, 0x03, 0xd1, 0x0f, 0xe2, 0x71, 0x99, 0xde, 0x81, 0x9d, 0x73, 0x63, 0xae, 0xdc,
	0x6c, 0xb6, 0xff, 0x6d, 0x93, 0xed, 0x7d, 0x71, 0xb9, 0x91, 0x1c, 0xf3, 0xc3, 0xf4, 0x05, 0xc0,
	0x54, 0x3b, 0x65, 0xe9, 0x93, 0x07, 0xac, 0xd7, 0x80, 0x61, 0x07, 0x20, 0x95, 0xc5, 0x0a, 0x29,
	0x10, 0xb6, 0xa2, 0x9c, 0x11, 0x77, 0x61, 0xb7, 0xf5, 0x06, 0xc6, 0xf7, 0xc0, 0xb8, 0x72, 0x2e,
	0xaa, 0x21, 0x87, 0xb0, 0x3d, 0x2b, 0x54, 0xe5, 0xb8, 0x15, 0x9e, 0x92, 0xfc, 0x9d, 0x02, 0x8d,
	0x9e, 0x31, 0xf7, 0x61, 0x6f, 0xc9, 0x1d, 0x38, 0x3f, 0x02, 0x67, 0xa7, 0x95, 0x35, 0xa0, 0xa9,
	0x23, 0x42, 0x35, 0xad, 0x5f, 0xe7, 0xd2, 0x3f, 0x80, 0x1a, 0x7d, 0x00, 0x2d, 0xb9, 0x03, 0xe8,
	0x67, 0x04, 0xb5, 0x32, 0x0f, 0x3a, 0xf8, 0xba, 0x0e, 0x7d, 0x81, 0x66, 0xa1, 0x95, 0xc1, 0xf4,
	0x21, 0xf4, 0xea, 0x1d, 0xe8, 0xce, 0x27, 0x6e, 0x58, 0xd8, 0x8f, 0x57, 0xfe, 0x53, 0x04, 0x61,
	0xfa, 0x06, 0xf6, 0xfc, 0xf4, 0xf3, 0xa5, 0xf1, 0x73, 0xef, 0x13, 0x36, 0x67, 0x1d, 0x73, 0x77,
	0x49, 0x8e, 0x38, 0x1e, 0xb7, 0xb1, 0xd8, 0x9d, 0xaf, 0x3e, 0x48, 0x9f, 0xc1, 0x66, 0xdc, 0x3a,
	0xee, 0xb3, 0xaf, 0x78, 0xf3, 0xb7, 0x8a, 0x61, 0x27, 0x8f, 0xc2, 0xb7, 0x68, 0xe4, 0xe9, 0x6b,
	0x48, 0x48, 0x9f, 0x71, 0xd3, 0xbc, 0xeb, 0x79, 0xf6, 0x1f, 0x67, 0x92, 0x35, 0xad, 0xc8, 0x84,
	0x3e, 0x13, 0xbe, 0xca, 0xe0, 0x06, 0x24, 0xfc, 0x3b, 0xbd, 0x06, 0x1b, 0x1c, 0xf9, 0x59, 0x7c,
	0x3e, 0xe6, 0xe6, 0xf4, 0x44, 0x8f, 0xc3, 0x71, 0xf9, 0xf2, 0x04, 0x86, 0x52, 0x77, 0x10, 0xed,
	0x1d, 0xbf, 0x7d, 0x72, 0xb1, 0x0b, 0x7f, 0xb7, 0x51, 0xdf, 0xd0, 0xe3, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x77, 0xef, 0x44, 0x22, 0x04, 0x00, 0x00,
}
