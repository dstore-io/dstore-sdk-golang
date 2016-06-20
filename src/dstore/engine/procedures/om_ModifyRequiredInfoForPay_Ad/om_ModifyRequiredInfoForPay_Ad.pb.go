// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyRequiredInfoForPay_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyRequiredInfoForPay_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyRequiredInfoForPay_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyRequiredInfoForPay_Ad

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
	PaymentForShippingId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	PaymentForShippingIdNull   bool                        `protobuf:"varint,1001,opt,name=payment_for_shipping_id_null,json=paymentForShippingIdNull" json:"payment_for_shipping_id_null,omitempty"`
	PersonCharacCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
	PersonCharacCategoryIdNull bool                        `protobuf:"varint,1002,opt,name=person_charac_category_id_null,json=personCharacCategoryIdNull" json:"person_charac_category_id_null,omitempty"`
	Delete                     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete" json:"delete,omitempty"`
	DeleteNull                 bool                        `protobuf:"varint,1003,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPaymentForShippingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentForShippingId
	}
	return nil
}

func (m *Parameters) GetPersonCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacCategoryId
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyRequiredInfoForPay_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyRequiredInfoForPay_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyRequiredInfoForPay_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0xa7, 0x3d, 0x93, 0x96, 0xe9, 0x83, 0xb2, 0x4a, 0xbd, 0x5e, 0xb4, 0x94, 0xfa, 0xe2, 0x83,
	0x5c, 0xc4, 0x3e, 0x28, 0x82, 0x88, 0x06, 0x85, 0x80, 0x29, 0x65, 0x85, 0x80, 0xbe, 0x2c, 0xdb,
	0xec, 0x34, 0x5d, 0xb8, 0xdb, 0x3d, 0x77, 0x2f, 0x96, 0x7c, 0x0b, 0x5f, 0xfc, 0x3a, 0x7e, 0x1f,
	0xf5, 0x4b, 0x38, 0x77, 0xbb, 0xa5, 0x26, 0x26, 0x8a, 0x7d, 0xb9, 0xdb, 0xd9, 0x99, 0xdf, 0x6f,
	0x7e, 0x3b, 0x7f, 0xe0, 0x85, 0xf2, 0xb5, 0x75, 0xd8, 0x47, 0x33, 0xd5, 0x06, 0xfb, 0x95, 0xb3,
	0x13, 0x54, 0x33, 0x87, 0xbe, 0x6f, 0x4b, 0x31, 0xb2, 0x4a, 0x9f, 0xcd, 0x39, 0x7e, 0x9a, 0x69,
	0x87, 0x6a, 0x68, 0xce, 0xec, 0x5b, 0xeb, 0x4e, 0xe4, 0x5c, 0xbc, 0x52, 0x39, 0x45, 0xd6, 0x96,
	0x3d, 0x0a, 0xf0, 0x3c, 0xc0, 0xf3, 0xbf, 0x63, 0xb2, 0xdb, 0x31, 0xd9, 0x67, 0x59, 0xcc, 0xd0,
	0x07, 0x8a, 0x6c, 0x6f, 0x51, 0x01, 0x3a, 0x67, 0x5d, 0x74, 0xf5, 0x16, 0x5d, 0x25, 0x7a, 0x2f,
	0xa7, 0x18, 0x9d, 0x0f, 0x96, 0x9d, 0xb5, 0xd4, 0x94, 0xcd, 0x95, 0xb2, 0xd6, 0xd6, 0x84, 0xa0,
	0xc3, 0xaf, 0x09, 0xc0, 0x89, 0x74, 0x92, 0xbc, 0xe8, 0x3c, 0xe3, 0x70, 0xb7, 0x92, 0xf3, 0x12,
	0x4d, 0x2d, 0x28, 0x52, 0xf8, 0x73, 0x5d, 0x55, 0xda, 0x4c, 0x85, 0x56, 0xe9, 0xc6, 0xc1, 0xc6,
	0xc3, 0x9d, 0x27, 0xbd, 0x3c, 0x3e, 0x28, 0x4a, 0xd4, 0xa6, 0xc6, 0x29, 0xba, 0x71, 0x63, 0xf1,
	0x3b, 0x11, 0x4b, 0x0f, 0x7a, 0x1f, 0x91, 0x43, 0xc5, 0x5e, 0xc2, 0xbd, 0x35, 0x9c, 0xc2, 0xcc,
	0x8a, 0x22, 0xfd, 0xbe, 0x45, 0xcc, 0xdb, 0x3c, 0x5d, 0x05, 0x3e, 0xa6, 0x00, 0x36, 0x86, 0xbd,
	0x8a, 0xc4, 0x59, 0x23, 0x26, 0xe7, 0x24, 0x75, 0x22, 0x26, 0x92, 0x92, 0x5a, 0x37, 0x6f, 0x64,
	0x6d, 0xfe, 0x5b, 0xd6, 0x6e, 0x40, 0x0f, 0x5a, 0xf0, 0x20, 0x62, 0x49, 0xd8, 0x00, 0xf6, 0xd7,
	0xf2, 0x06, 0x69, 0x3f, 0x82, 0xb4, 0x6c, 0x35, 0x41, 0x2b, 0xee, 0x08, 0xba, 0x0a, 0x0b, 0xaa,
	0x5e, 0x9a, 0xac, 0x54, 0x72, 0x6a, 0x6d, 0x81, 0xd2, 0x04, 0x25, 0x31, 0x94, 0x1d, 0xc0, 0x4e,
	0x38, 0x85, 0x34, 0x3f, 0x43, 0x1a, 0x08, 0x77, 0x0d, 0xed, 0xe1, 0xb7, 0x4d, 0xd8, 0xe6, 0xe8,
	0x2b, 0x6b, 0x3c, 0xb2, 0xc7, 0xd0, 0x69, 0xbb, 0x1e, 0x7b, 0x90, 0xe5, 0x8b, 0x43, 0x15, 0x26,
	0xe2, 0x4d, 0xf3, 0xe5, 0x21, 0x90, 0x7d, 0x80, 0x5b, 0x4d, 0xbf, 0xc5, 0x6f, 0x0d, 0xa7, 0x4a,
	0x25, 0x04, 0xce, 0x97, 0xc0, 0xcb, 0x63, 0x31, 0x22, 0x7b, 0x78, 0x65, 0xf3, 0x9b, 0xe5, 0xe2,
	0x05, 0x7b, 0x06, 0x5b, 0x71, 0xce, 0xe8, 0xc5, 0x0d, 0xe3, 0xfe, 0x1f, 0x8c, 0x61, 0x0a, 0x47,
	0xe1, 0xcf, 0x2f, 0xc3, 0xd9, 0x3b, 0x48, 0x9c, 0xbd, 0x48, 0x6f, 0xb4, 0xa8, 0xe7, 0xf9, 0xff,
	0x6c, 0x46, 0x7e, 0x59, 0x8b, 0x9c, 0xdb, 0x0b, 0xde, 0xd0, 0x64, 0xf7, 0x21, 0xa1, 0x33, 0xdb,
	0x85, 0x2e, 0x59, 0xcd, 0x24, 0x7c, 0x39, 0xa6, 0xea, 0x74, 0x78, 0x87, 0xcc, 0xa1, 0x7a, 0x3d,
	0x86, 0x9e, 0xb6, 0x4b, 0x39, 0xae, 0x96, 0xf7, 0xe3, 0xd3, 0x6b, 0xae, 0xf5, 0x69, 0xb7, 0xdd,
	0x9b, 0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0xbf, 0xf9, 0x61, 0x18, 0x04, 0x00, 0x00,
}
