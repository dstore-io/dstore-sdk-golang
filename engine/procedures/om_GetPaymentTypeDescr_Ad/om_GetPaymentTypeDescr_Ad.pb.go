// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPaymentTypeDescr_Ad.proto
// DO NOT EDIT!

/*
Package om_GetPaymentTypeDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPaymentTypeDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPaymentTypeDescr_Ad

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
	PaymentTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PaymentTypeIdNull bool                        `protobuf:"varint,1001,opt,name=payment_type_id_null,json=paymentTypeIdNull" json:"payment_type_id_null,omitempty"`
	LanguageId        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull    bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Language               *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=language" json:"language,omitempty"`
	PaymentTypeDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=payment_type_description,json=paymentTypeDescription" json:"payment_type_description,omitempty"`
	LanguageId             *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetLanguage() *dstore_values.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *Response_Row) GetPaymentTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentTypeDescription
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPaymentTypeDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPaymentTypeDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPaymentTypeDescr_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetPaymentTypeDescr_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x65, 0xdd, 0x6e, 0xbb, 0xdc, 0xa2, 0xad, 0x51, 0xca, 0xb8, 0x0b, 0x22, 0xf5, 0x41, 0x7d,
	0xc9, 0x8a, 0x22, 0x16, 0xd4, 0x07, 0xbf, 0x28, 0x7d, 0xe8, 0x52, 0x82, 0x0a, 0xf6, 0x65, 0x18,
	0x3b, 0xd7, 0x61, 0x70, 0x27, 0x19, 0x92, 0xac, 0xa5, 0x3f, 0xc1, 0x37, 0x3f, 0x7e, 0xa5, 0xe2,
	0x8f, 0xf0, 0x66, 0x92, 0x71, 0x9a, 0x11, 0x69, 0xfb, 0xb2, 0xcb, 0x9d, 0x7b, 0xce, 0x3d, 0x27,
	0xe7, 0x26, 0xb0, 0x93, 0x1b, 0xab, 0x34, 0xce, 0x50, 0x16, 0xa5, 0xc4, 0x59, 0xad, 0xd5, 0x11,
	0xe6, 0x4b, 0x8d, 0x66, 0xa6, 0xaa, 0x74, 0x17, 0xed, 0x41, 0x76, 0x52, 0xa1, 0xb4, 0x6f, 0x4e,
	0x6a, 0x7c, 0x85, 0xe6, 0x48, 0xa7, 0xcf, 0x73, 0x4e, 0x20, 0xab, 0xd8, 0x1d, 0xcf, 0xe4, 0x9e,
	0xc9, 0xff, 0x0b, 0x9f, 0x5c, 0x0b, 0x12, 0x9f, 0xb3, 0xc5, 0x12, 0x8d, 0x67, 0x4f, 0x6e, 0xc4,
	0xba, 0xa8, 0xb5, 0xd2, 0xa1, 0x35, 0x8d, 0x5b, 0x15, 0x1a, 0x93, 0x15, 0x18, 0x9a, 0xb7, 0xfb,
	0x4d, 0x9b, 0x95, 0xf2, 0xa3, 0xd2, 0x55, 0x66, 0x4b, 0x25, 0x3d, 0x68, 0xfb, 0xf7, 0x00, 0xe0,
	0x20, 0xd3, 0x19, 0x75, 0x51, 0x1b, 0xf6, 0x12, 0x36, 0x6a, 0xef, 0x2b, 0xb5, 0x64, 0x2c, 0x2d,
	0xf3, 0x64, 0x70, 0x6b, 0x70, 0x77, 0xfd, 0xc1, 0x94, 0x87, 0x33, 0x04, 0x6b, 0xa5, 0xb4, 0x58,
	0xa0, 0x7e, 0xe7, 0x2a, 0x71, 0xb9, 0xee, 0xce, 0xb2, 0x97, 0xb3, 0xfb, 0x70, 0xbd, 0x37, 0x24,
	0x95, 0xcb, 0xc5, 0x22, 0xf9, 0xb9, 0x46, 0xa3, 0xc6, 0xe2, 0x6a, 0x84, 0x9e, 0x53, 0x87, 0x3d,
	0x85, 0xf5, 0x45, 0x26, 0x8b, 0x25, 0x99, 0x77, 0x92, 0x97, 0xce, 0x96, 0x84, 0x16, 0x4f, 0x7a,
	0xf7, 0x60, 0xf3, 0x14, 0xdb, 0x6b, 0xfd, 0xf2, 0x5a, 0x57, 0x3a, 0x98, 0x13, 0xda, 0xfe, 0xb2,
	0x02, 0x63, 0x81, 0xa6, 0x56, 0xd2, 0x20, 0xf9, 0x1c, 0x35, 0x61, 0x86, 0x23, 0x4e, 0x78, 0xbc,
	0x26, 0x1f, 0xf4, 0x6b, 0xf7, 0x2b, 0x3c, 0x90, 0xbd, 0x87, 0x4d, 0x17, 0x63, 0x7a, 0x2a, 0x47,
	0x32, 0x3b, 0x24, 0x32, 0xef, 0x91, 0xfb, 0x69, 0xef, 0x53, 0xbd, 0xd7, 0xd5, 0x62, 0xa3, 0x8a,
	0x3f, 0xb0, 0x1d, 0x58, 0x0b, 0xeb, 0x4b, 0x86, 0xcd, 0xc4, 0x9b, 0xff, 0x4c, 0xf4, 0xcb, 0xdd,
	0xf7, 0xff, 0xa2, 0x85, 0xb3, 0x5d, 0x18, 0x6a, 0x75, 0x9c, 0xac, 0x34, 0xac, 0x47, 0xfc, 0x9c,
	0x77, 0x8d, 0xb7, 0x31, 0x70, 0xa1, 0x8e, 0x85, 0x9b, 0x30, 0xa1, 0xbb, 0x30, 0xa4, 0x82, 0x6d,
	0xc1, 0x2a, 0x95, 0x6e, 0x11, 0x5f, 0xe7, 0x94, 0xcc, 0x48, 0x8c, 0xa8, 0xa4, 0x9c, 0x1f, 0xc3,
	0xb8, 0x8d, 0x33, 0xf9, 0x36, 0x8f, 0x33, 0x0b, 0x3b, 0x32, 0x56, 0x97, 0xb2, 0xf0, 0x2b, 0xfa,
	0x0b, 0x66, 0x6f, 0x21, 0x89, 0x2e, 0x44, 0xee, 0x3c, 0x94, 0x75, 0x13, 0xdf, 0xf7, 0xb3, 0x07,
	0x6d, 0xd5, 0x3d, 0xfb, 0x9e, 0xca, 0x9e, 0xc5, 0xb7, 0xe6, 0xc7, 0xfc, 0x42, 0xd7, 0xe6, 0xc5,
	0x21, 0x4c, 0x4b, 0xd5, 0x8b, 0xab, 0x7b, 0xd4, 0x87, 0x4f, 0x0a, 0x65, 0xf2, 0x4f, 0x6d, 0x3f,
	0xbf, 0xd0, 0xbb, 0xff, 0xb0, 0xda, 0xbc, 0xae, 0x87, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x00,
	0xaf, 0x9d, 0xb9, 0x34, 0x04, 0x00, 0x00,
}