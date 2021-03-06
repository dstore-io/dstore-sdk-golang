// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyPaymentTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyPaymentTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyPaymentTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyPaymentTypes_Ad

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
	PaymentTypeId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PaymentTypeIdNull                bool                        `protobuf:"varint,1001,opt,name=payment_type_id_null,json=paymentTypeIdNull" json:"payment_type_id_null,omitempty"`
	PaymentTypeDescription           *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=payment_type_description,json=paymentTypeDescription" json:"payment_type_description,omitempty"`
	PaymentTypeDescriptionNull       bool                        `protobuf:"varint,1002,opt,name=payment_type_description_null,json=paymentTypeDescriptionNull" json:"payment_type_description_null,omitempty"`
	RegionId                         *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	RegionIdNull                     bool                        `protobuf:"varint,1003,opt,name=region_id_null,json=regionIdNull" json:"region_id_null,omitempty"`
	BruttoSumFrom                    *dstore_values.DecimalValue `protobuf:"bytes,4,opt,name=brutto_sum_from,json=bruttoSumFrom" json:"brutto_sum_from,omitempty"`
	BruttoSumFromNull                bool                        `protobuf:"varint,1004,opt,name=brutto_sum_from_null,json=bruttoSumFromNull" json:"brutto_sum_from_null,omitempty"`
	BruttoSumTo                      *dstore_values.DecimalValue `protobuf:"bytes,5,opt,name=brutto_sum_to,json=bruttoSumTo" json:"brutto_sum_to,omitempty"`
	BruttoSumToNull                  bool                        `protobuf:"varint,1005,opt,name=brutto_sum_to_null,json=bruttoSumToNull" json:"brutto_sum_to_null,omitempty"`
	CurrencyId                       *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencyIdNull                   bool                        `protobuf:"varint,1006,opt,name=currency_id_null,json=currencyIdNull" json:"currency_id_null,omitempty"`
	SurchargeTypeId                  *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull              bool                        `protobuf:"varint,1007,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	Cost                             *dstore_values.DecimalValue `protobuf:"bytes,8,opt,name=cost" json:"cost,omitempty"`
	CostNull                         bool                        `protobuf:"varint,1008,opt,name=cost_null,json=costNull" json:"cost_null,omitempty"`
	DeletePaymentType                *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=delete_payment_type,json=deletePaymentType" json:"delete_payment_type,omitempty"`
	DeletePaymentTypeNull            bool                        `protobuf:"varint,1009,opt,name=delete_payment_type_null,json=deletePaymentTypeNull" json:"delete_payment_type_null,omitempty"`
	Active                           *dstore_values.BooleanValue `protobuf:"bytes,10,opt,name=active" json:"active,omitempty"`
	ActiveNull                       bool                        `protobuf:"varint,1010,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
	PredefBillContentDescription     *dstore_values.StringValue  `protobuf:"bytes,11,opt,name=predef_bill_content_description,json=predefBillContentDescription" json:"predef_bill_content_description,omitempty"`
	PredefBillContentDescriptionNull bool                        `protobuf:"varint,1011,opt,name=predef_bill_content_description_null,json=predefBillContentDescriptionNull" json:"predef_bill_content_description_null,omitempty"`
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

func (m *Parameters) GetPaymentTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentTypeDescription
	}
	return nil
}

func (m *Parameters) GetRegionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RegionId
	}
	return nil
}

func (m *Parameters) GetBruttoSumFrom() *dstore_values.DecimalValue {
	if m != nil {
		return m.BruttoSumFrom
	}
	return nil
}

func (m *Parameters) GetBruttoSumTo() *dstore_values.DecimalValue {
	if m != nil {
		return m.BruttoSumTo
	}
	return nil
}

func (m *Parameters) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetCost() *dstore_values.DecimalValue {
	if m != nil {
		return m.Cost
	}
	return nil
}

func (m *Parameters) GetDeletePaymentType() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeletePaymentType
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Parameters) GetPredefBillContentDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PredefBillContentDescription
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyPaymentTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyPaymentTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyPaymentTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyPaymentTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0xc7, 0x05, 0x21, 0x21, 0x4c, 0x0e, 0x04, 0xcc, 0x39, 0xc8, 0x27, 0xc0, 0x39, 0x88, 0x7e,
	0xa8, 0x95, 0x2a, 0xa7, 0x2a, 0x48, 0xad, 0xaa, 0x4a, 0x55, 0xa1, 0xa5, 0x42, 0x15, 0x14, 0xb9,
	0xa8, 0x12, 0xbd, 0xb1, 0x1c, 0x7b, 0xe3, 0x5a, 0xb5, 0xbd, 0xd6, 0xae, 0x0d, 0xca, 0x5b, 0xf4,
	0x45, 0x7a, 0xdf, 0x57, 0xea, 0xf7, 0xd7, 0x0b, 0x74, 0x77, 0xc7, 0x49, 0x6c, 0x13, 0xea, 0xf6,
	0x26, 0xc9, 0x7a, 0xe6, 0xff, 0xff, 0x4d, 0xc6, 0xbb, 0xb3, 0x70, 0xdb, 0xe5, 0x09, 0x65, 0xa4,
	0x4b, 0x22, 0xcf, 0x8f, 0x48, 0x37, 0x66, 0xd4, 0x21, 0x6e, 0xca, 0x08, 0xef, 0xd2, 0xd0, 0x3a,
	0xa0, 0xae, 0xdf, 0x1f, 0x1c, 0xd9, 0x83, 0x90, 0x44, 0xc9, 0xf1, 0x20, 0x26, 0xdc, 0x7a, 0xe0,
	0x1a, 0x22, 0x27, 0xa1, 0xda, 0x55, 0x14, 0x1a, 0x28, 0x34, 0x2e, 0xca, 0xee, 0x2c, 0x67, 0x80,
	0x53, 0x3b, 0x48, 0x09, 0x47, 0x71, 0xe7, 0xdf, 0x22, 0x95, 0x30, 0x46, 0x59, 0x16, 0x5a, 0x2d,
	0x86, 0x42, 0xc2, 0xb9, 0xed, 0x91, 0x2c, 0x78, 0xa9, 0x1c, 0x4c, 0x6c, 0x3f, 0xea, 0x53, 0x16,
	0xda, 0x89, 0x4f, 0x23, 0x4c, 0xda, 0x7c, 0x03, 0x00, 0x47, 0x36, 0xb3, 0x45, 0x94, 0x30, 0xae,
	0xed, 0x42, 0x3b, 0xc6, 0x9a, 0xac, 0x44, 0x14, 0x65, 0xf9, 0xae, 0x3e, 0xb5, 0x31, 0x75, 0xad,
	0x75, 0x6b, 0xd5, 0xc8, 0xfe, 0x42, 0x56, 0x9a, 0x1f, 0x25, 0xc4, 0x23, 0xec, 0xb9, 0x5c, 0x99,
	0xf3, 0xf1, 0xf8, 0x7f, 0xec, 0xbb, 0xda, 0x4d, 0xf8, 0xbb, 0x64, 0x62, 0x45, 0x69, 0x10, 0xe8,
	0xef, 0x66, 0x85, 0x55, 0xd3, 0x5c, 0x2a, 0x64, 0x1f, 0x8a, 0x88, 0x76, 0x0c, 0x7a, 0x41, 0xe1,
	0x12, 0xee, 0x30, 0x3f, 0x96, 0x75, 0xea, 0xd3, 0x8a, 0xdf, 0x29, 0xf1, 0x79, 0xc2, 0xfc, 0xc8,
	0x43, 0xfc, 0x4a, 0xce, 0xf0, 0xe1, 0x58, 0xa9, 0xed, 0xc0, 0xfa, 0x45, 0xae, 0x58, 0xd0, 0x7b,
	0x2c, 0xa8, 0x33, 0x59, 0xaf, 0x2a, 0xbb, 0x03, 0x73, 0x8c, 0x78, 0x52, 0x21, 0x5a, 0x51, 0xab,
	0x6e, 0x45, 0x13, 0xb3, 0x45, 0x17, 0xae, 0xc0, 0xc2, 0x48, 0x89, 0xb8, 0x0f, 0x88, 0xfb, 0x6b,
	0x98, 0xa2, 0x00, 0xa2, 0xe3, 0x3d, 0x96, 0x26, 0x09, 0xb5, 0x78, 0x1a, 0x5a, 0x7d, 0x46, 0x43,
	0x7d, 0x66, 0x22, 0xc6, 0x25, 0x8e, 0x1f, 0xda, 0x41, 0xd6, 0x71, 0xd4, 0x3c, 0x4b, 0xc3, 0x3d,
	0xa1, 0x90, 0x1d, 0x2f, 0x99, 0x20, 0xf1, 0x63, 0xd6, 0xf1, 0x42, 0xb6, 0xc2, 0xde, 0x87, 0xf9,
	0x9c, 0x22, 0xa1, 0x7a, 0xbd, 0x1a, 0xda, 0x1a, 0xd9, 0x1c, 0x53, 0xed, 0x06, 0x68, 0x05, 0x03,
	0x04, 0x7e, 0x42, 0x60, 0x3b, 0x97, 0xa9, 0x70, 0xf7, 0xa0, 0xe5, 0xa4, 0x8c, 0x91, 0xc8, 0x19,
	0xc8, 0x46, 0x36, 0xaa, 0x1b, 0x09, 0xc3, 0x7c, 0xd1, 0xca, 0xeb, 0xb0, 0x98, 0x53, 0x23, 0xe9,
	0x33, 0x92, 0x16, 0xc6, 0x69, 0x0a, 0xf4, 0x18, 0x96, 0x78, 0xca, 0x9c, 0x97, 0x36, 0xf3, 0xc8,
	0x68, 0x0b, 0xcf, 0x56, 0xe3, 0xda, 0x23, 0x55, 0xb6, 0x89, 0xb7, 0x61, 0xe5, 0x9c, 0x11, 0x92,
	0xbf, 0x20, 0x79, 0xb9, 0xa4, 0x50, 0xf8, 0x2e, 0xcc, 0x38, 0x94, 0x27, 0x7a, 0xb3, 0xba, 0x9b,
	0x2a, 0x51, 0x5b, 0x83, 0x39, 0xf9, 0x8d, 0xce, 0x5f, 0xd1, 0xb9, 0x29, 0x9f, 0x28, 0xbb, 0x27,
	0xb0, 0xec, 0x92, 0x40, 0x1c, 0x4d, 0x2b, 0xbf, 0x91, 0xf5, 0xb9, 0x89, 0xee, 0x3d, 0x4a, 0x03,
	0x62, 0x47, 0xe8, 0xbe, 0x84, 0xba, 0xdc, 0x80, 0x11, 0x5b, 0x59, 0x9f, 0x60, 0x86, 0xe4, 0x6f,
	0x48, 0xfe, 0xe7, 0x9c, 0x4a, 0x95, 0xb1, 0x05, 0x0d, 0xdb, 0x49, 0xfc, 0x53, 0xa2, 0x43, 0x35,
	0x39, 0x4b, 0xd5, 0x36, 0xa0, 0x85, 0xbf, 0x90, 0xf0, 0x1d, 0x09, 0x80, 0xcf, 0x94, 0xad, 0x0d,
	0xff, 0xc7, 0x8c, 0xb8, 0xa4, 0x6f, 0xf5, 0xfc, 0x20, 0xb0, 0x1c, 0x2a, 0x5e, 0x88, 0xa8, 0x2a,
	0x7f, 0xf8, 0x5b, 0x95, 0x87, 0x7f, 0x0d, 0x2d, 0x76, 0x84, 0xc3, 0x2e, 0x1a, 0xe4, 0x47, 0xc0,
	0x53, 0xb8, 0x5c, 0x81, 0xc0, 0xea, 0x7e, 0x60, 0x75, 0x1b, 0xbf, 0x32, 0x93, 0x35, 0x6f, 0xbe,
	0x9d, 0x86, 0xa6, 0x49, 0x78, 0x4c, 0x23, 0x4e, 0xc4, 0xb1, 0xab, 0xab, 0x69, 0x9c, 0xcd, 0xc8,
	0x51, 0x99, 0xd9, 0x98, 0xc7, 0x49, 0xfd, 0x48, 0x7e, 0x9a, 0x98, 0xa8, 0x9d, 0xc0, 0xa2, 0x9c,
	0xc3, 0x56, 0x6e, 0x10, 0x8b, 0x01, 0x57, 0x13, 0x62, 0xa3, 0x24, 0x2e, 0x8f, 0xeb, 0x03, 0xb1,
	0xde, 0x1f, 0xaf, 0xcd, 0x76, 0x58, 0x7c, 0x20, 0x5e, 0xef, 0x6c, 0x36, 0xff, 0xc5, 0x9c, 0x92,
	0x8e, 0xff, 0x9d, 0x73, 0xc4, 0xdb, 0xe1, 0x00, 0xbf, 0xcd, 0x61, 0xba, 0xb6, 0x07, 0x35, 0x46,
	0xcf, 0xc4, 0xd8, 0x91, 0xaa, 0x6d, 0xe3, 0xf7, 0xee, 0x2a, 0x63, 0xd8, 0x05, 0xc3, 0xa4, 0x67,
	0xa6, 0x34, 0xe8, 0xac, 0x43, 0x4d, 0xfc, 0xd6, 0x56, 0xa0, 0x21, 0x56, 0xf2, 0xdc, 0xbd, 0x3e,
	0x14, 0x7d, 0xa9, 0x9b, 0x75, 0xb1, 0xdc, 0x77, 0x77, 0x4e, 0x60, 0xd5, 0xa7, 0x25, 0xf7, 0xf1,
	0x15, 0xfa, 0xe2, 0xae, 0x47, 0xb9, 0xfb, 0x6a, 0x18, 0x77, 0xff, 0xe4, 0x96, 0xed, 0x35, 0xd4,
	0x65, 0xb6, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0x00, 0x69, 0x7a, 0x9f, 0xa1, 0x07, 0x00, 0x00,
}
