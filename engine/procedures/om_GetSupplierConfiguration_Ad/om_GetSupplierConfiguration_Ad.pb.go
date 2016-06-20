// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetSupplierConfiguration_Ad.proto
// DO NOT EDIT!

/*
Package om_GetSupplierConfiguration_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetSupplierConfiguration_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetSupplierConfiguration_Ad

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
	SupplierId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=supplier_id,json=supplierId" json:"supplier_id,omitempty"`
	SupplierIdNull bool                        `protobuf:"varint,1001,opt,name=supplier_id_null,json=supplierIdNull" json:"supplier_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSupplierId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SupplierId
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
	RowId                          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PaymentTypeDescription         *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=payment_type_description,json=paymentTypeDescription" json:"payment_type_description,omitempty"`
	ShippingTypeDescription        *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=shipping_type_description,json=shippingTypeDescription" json:"shipping_type_description,omitempty"`
	PurchasePriceCurrencyId        *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=purchase_price_currency_id,json=purchasePriceCurrencyId" json:"purchase_price_currency_id,omitempty"`
	PurchasePriceCurrencySymbol    *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=purchase_price_currency_symbol,json=purchasePriceCurrencySymbol" json:"purchase_price_currency_symbol,omitempty"`
	PurchasePriceCharacteristicId  *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=purchase_price_characteristic_id,json=purchasePriceCharacteristicId" json:"purchase_price_characteristic_id,omitempty"`
	ShippingTypeId                 *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=shipping_type_id,json=shippingTypeId" json:"shipping_type_id,omitempty"`
	PaymentTypeId                  *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PurchasePriceCharacDescription *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=purchase_price_charac_description,json=purchasePriceCharacDescription" json:"purchase_price_charac_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPaymentTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentTypeDescription
	}
	return nil
}

func (m *Response_Row) GetShippingTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ShippingTypeDescription
	}
	return nil
}

func (m *Response_Row) GetPurchasePriceCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PurchasePriceCurrencyId
	}
	return nil
}

func (m *Response_Row) GetPurchasePriceCurrencySymbol() *dstore_values.StringValue {
	if m != nil {
		return m.PurchasePriceCurrencySymbol
	}
	return nil
}

func (m *Response_Row) GetPurchasePriceCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PurchasePriceCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetShippingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ShippingTypeId
	}
	return nil
}

func (m *Response_Row) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

func (m *Response_Row) GetPurchasePriceCharacDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PurchasePriceCharacDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetSupplierConfiguration_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetSupplierConfiguration_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetSupplierConfiguration_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetSupplierConfiguration_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x55, 0xbf, 0x7c, 0x69, 0xab, 0xa9, 0x68, 0x2b, 0x23, 0xb5, 0x6e, 0x22, 0xa2, 0x50, 0x6e,
	0x40, 0x42, 0x0e, 0x82, 0x1b, 0x84, 0x40, 0x02, 0x5a, 0x40, 0x95, 0x68, 0x54, 0xb9, 0xfc, 0xa8,
	0xdc, 0x18, 0xd7, 0x9e, 0x38, 0x2b, 0xe2, 0xdd, 0xd5, 0xae, 0x4d, 0x95, 0x47, 0xe0, 0x8e, 0x7f,
	0x78, 0x43, 0xc4, 0x5b, 0x30, 0xb6, 0xd7, 0x4a, 0xec, 0x06, 0x92, 0xde, 0x24, 0x1a, 0xcf, 0x9c,
	0x33, 0x7b, 0xce, 0xce, 0x2c, 0x3c, 0x08, 0x75, 0x22, 0x14, 0xf6, 0x90, 0x47, 0x8c, 0x63, 0x4f,
	0x2a, 0x11, 0x60, 0x98, 0x2a, 0xd4, 0x3d, 0x11, 0x7b, 0xcf, 0x30, 0x39, 0x4e, 0xa5, 0x1c, 0x31,
	0x54, 0x7b, 0x82, 0x0f, 0x58, 0x94, 0x2a, 0x3f, 0x61, 0x82, 0x7b, 0x8f, 0x42, 0x87, 0x2a, 0x13,
	0x61, 0xdd, 0x2c, 0xe0, 0x4e, 0x01, 0x77, 0xfe, 0x8d, 0x69, 0x5d, 0x36, 0xcd, 0xde, 0xfb, 0xa3,
	0x14, 0x75, 0x41, 0xd1, 0xda, 0xa9, 0x9e, 0x00, 0x95, 0x12, 0xca, 0xa4, 0xda, 0xd5, 0x54, 0x8c,
	0x5a, 0xfb, 0x11, 0x9a, 0xe4, 0xb5, 0x7a, 0x32, 0xf1, 0x19, 0x1f, 0x08, 0x15, 0xe7, 0xed, 0x8a,
	0xa2, 0xdd, 0x14, 0xe0, 0xc8, 0x57, 0x3e, 0x25, 0x51, 0x69, 0xeb, 0x3e, 0xac, 0x69, 0x73, 0x36,
	0x8f, 0x85, 0xf6, 0x52, 0x77, 0xe9, 0xfa, 0xda, 0xed, 0xb6, 0x63, 0x34, 0x98, 0x53, 0x31, 0x9e,
	0x60, 0x84, 0xea, 0x55, 0x16, 0xb9, 0x50, 0xd6, 0x1f, 0x84, 0xd6, 0x0d, 0xd8, 0x9c, 0x42, 0x7b,
	0x3c, 0x1d, 0x8d, 0xec, 0xdf, 0x2b, 0xc4, 0xb1, 0xea, 0xae, 0x4f, 0xca, 0xfa, 0xf4, 0x79, 0xf7,
	0xc3, 0x2a, 0xac, 0xba, 0xa8, 0xa5, 0xe0, 0x1a, 0xad, 0x5b, 0xd0, 0xcc, 0x45, 0x99, 0x7e, 0x2d,
	0xa7, 0xea, 0x59, 0x21, 0xf8, 0x49, 0xf6, 0xeb, 0x16, 0x85, 0xd6, 0x09, 0x6c, 0x66, 0x72, 0xbc,
	0x29, 0x3d, 0xf6, 0x7f, 0xdd, 0x06, 0x81, 0x9d, 0x1a, 0xb8, 0xae, 0xfa, 0x90, 0xe2, 0x83, 0x49,
	0xec, 0x6e, 0xc4, 0xd5, 0x0f, 0xd6, 0x5d, 0x58, 0x31, 0x36, 0xda, 0x8d, 0x9c, 0xb1, 0x73, 0x8e,
	0xb1, 0x30, 0xf9, 0xb0, 0xf8, 0x77, 0xcb, 0x72, 0xeb, 0x39, 0x34, 0x94, 0x38, 0xb3, 0xff, 0xcf,
	0x51, 0xf7, 0x9c, 0x8b, 0x5c, 0xbc, 0x53, 0x7a, 0xe1, 0xb8, 0xe2, 0xcc, 0xcd, 0x68, 0x5a, 0xbf,
	0x9a, 0xd0, 0xa0, 0xc0, 0xda, 0x82, 0x65, 0x0a, 0xb3, 0xdb, 0xf8, 0xd8, 0x27, 0x7b, 0x9a, 0x6e,
	0x93, 0x42, 0x32, 0xfb, 0x25, 0xd8, 0xd2, 0x1f, 0xc7, 0xc8, 0x13, 0x2f, 0x19, 0x4b, 0xf4, 0x42,
	0xd4, 0x81, 0x62, 0x32, 0xb7, 0xe2, 0x53, 0xbf, 0x6a, 0xa4, 0xb9, 0x38, 0x9d, 0x28, 0xc6, 0xa3,
	0xe2, 0xde, 0xb6, 0x0c, 0xf8, 0x05, 0x61, 0xf7, 0x27, 0x50, 0xeb, 0x35, 0xec, 0xe8, 0x21, 0x93,
	0x92, 0x0a, 0xcf, 0xf3, 0x7e, 0x9e, 0xcf, 0xbb, 0x5d, 0xa2, 0xeb, 0xc4, 0x27, 0xd0, 0x92, 0xa9,
	0x0a, 0x86, 0xbe, 0x46, 0x4f, 0x2a, 0x16, 0xa0, 0x17, 0xa4, 0x4a, 0x21, 0x0f, 0xc6, 0x99, 0xb6,
	0x2f, 0xfd, 0xf9, 0xa3, 0xb6, 0x5d, 0xe2, 0x8f, 0x32, 0xf8, 0x9e, 0x41, 0x93, 0x15, 0x6f, 0xa1,
	0xf3, 0x37, 0x6a, 0x3d, 0x8e, 0x4f, 0xc5, 0xc8, 0xfe, 0x3a, 0xff, 0xe0, 0xed, 0x99, 0xec, 0xc7,
	0x39, 0xde, 0x42, 0xe8, 0xd6, 0x3b, 0x0c, 0x69, 0x6b, 0x02, 0x5a, 0x1a, 0xa6, 0x13, 0x16, 0x64,
	0x12, 0xbe, 0x2d, 0x20, 0xe1, 0x4a, 0xb5, 0x49, 0x85, 0x83, 0x84, 0x3c, 0xa5, 0x05, 0xaa, 0x98,
	0x4f, 0xb4, 0xdf, 0x17, 0xa0, 0x5d, 0x9f, 0x36, 0x9d, 0x78, 0xf6, 0x61, 0xa3, 0x32, 0x1b, 0x44,
	0xf3, 0x63, 0x01, 0x9a, 0x4b, 0x53, 0x33, 0x41, 0x2c, 0x03, 0xb8, 0x3a, 0x53, 0x74, 0x65, 0x24,
	0x7e, 0xce, 0x77, 0xb6, 0x33, 0x43, 0xf4, 0xd4, 0x64, 0x3c, 0xf6, 0xa0, 0xcd, 0x44, 0x6d, 0x5d,
	0x26, 0xcf, 0xec, 0x9b, 0x87, 0x91, 0xd0, 0xe1, 0xbb, 0x32, 0x1f, 0x5e, 0xfc, 0x25, 0x3e, 0x5d,
	0xce, 0x9f, 0xba, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xda, 0x9a, 0x9f, 0xcb, 0x05,
	0x00, 0x00,
}