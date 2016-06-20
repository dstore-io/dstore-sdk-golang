// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyShippingTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyShippingTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyShippingTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyShippingTypes_Ad

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
	ShippingTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=shipping_type_id,json=shippingTypeId" json:"shipping_type_id,omitempty"`
	ShippingTypeIdNull               bool                        `protobuf:"varint,1001,opt,name=shipping_type_id_null,json=shippingTypeIdNull" json:"shipping_type_id_null,omitempty"`
	ShippingTypeDescription          *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=shipping_type_description,json=shippingTypeDescription" json:"shipping_type_description,omitempty"`
	ShippingTypeDescriptionNull      bool                        `protobuf:"varint,1002,opt,name=shipping_type_description_null,json=shippingTypeDescriptionNull" json:"shipping_type_description_null,omitempty"`
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
	DeleteShippingType               *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=delete_shipping_type,json=deleteShippingType" json:"delete_shipping_type,omitempty"`
	DeleteShippingTypeNull           bool                        `protobuf:"varint,1009,opt,name=delete_shipping_type_null,json=deleteShippingTypeNull" json:"delete_shipping_type_null,omitempty"`
	Active                           *dstore_values.BooleanValue `protobuf:"bytes,10,opt,name=active" json:"active,omitempty"`
	ActiveNull                       bool                        `protobuf:"varint,1010,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
	PredefBillContentDescription     *dstore_values.StringValue  `protobuf:"bytes,11,opt,name=predef_bill_content_description,json=predefBillContentDescription" json:"predef_bill_content_description,omitempty"`
	PredefBillContentDescriptionNull bool                        `protobuf:"varint,1011,opt,name=predef_bill_content_description_null,json=predefBillContentDescriptionNull" json:"predef_bill_content_description_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetShippingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ShippingTypeId
	}
	return nil
}

func (m *Parameters) GetShippingTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ShippingTypeDescription
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

func (m *Parameters) GetDeleteShippingType() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteShippingType
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyShippingTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyShippingTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyShippingTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x6d, 0x4f, 0xd3, 0x40,
	0x1c, 0x0f, 0x8c, 0x8d, 0xf1, 0x9f, 0x30, 0x38, 0x10, 0xcb, 0x86, 0x48, 0x50, 0xa3, 0x26, 0xa6,
	0x33, 0xa0, 0x09, 0x31, 0x26, 0x46, 0x1e, 0x34, 0xbc, 0x18, 0x6a, 0x21, 0x24, 0xfa, 0xa6, 0xe9,
	0xda, 0xdb, 0x68, 0xd2, 0xf6, 0x96, 0xbb, 0x16, 0xc2, 0xb7, 0xf0, 0xeb, 0xf0, 0x91, 0x7c, 0x7e,
	0xfa, 0x02, 0x5e, 0xef, 0xdf, 0x6d, 0x6d, 0x01, 0xab, 0x6f, 0xd8, 0xae, 0xf7, 0x7b, 0xda, 0x9f,
	0xde, 0xef, 0x60, 0xd3, 0x11, 0x21, 0xe3, 0xb4, 0x45, 0x83, 0x9e, 0x1b, 0xd0, 0x56, 0x9f, 0x33,
	0x9b, 0x3a, 0x11, 0xa7, 0xa2, 0xc5, 0x7c, 0xb3, 0xcd, 0x1c, 0xb7, 0x7b, 0x76, 0x70, 0xec, 0xf6,
	0xfb, 0x6e, 0xd0, 0x3b, 0x3c, 0xeb, 0x53, 0x61, 0xbe, 0x70, 0x74, 0x09, 0x0a, 0x19, 0xb9, 0x87,
	0x4c, 0x1d, 0x99, 0xfa, 0x95, 0xf0, 0xc6, 0x7c, 0x62, 0x71, 0x62, 0x79, 0x11, 0x15, 0xc8, 0x6e,
	0x2c, 0x65, 0x7d, 0x29, 0xe7, 0x8c, 0x27, 0x5b, 0xcd, 0xec, 0x96, 0x4f, 0x85, 0xb0, 0x7a, 0x34,
	0xd9, 0xbc, 0x9d, 0xdf, 0x0c, 0x2d, 0x37, 0xe8, 0x32, 0xee, 0x5b, 0xa1, 0xcb, 0x02, 0x04, 0xad,
	0x9d, 0x03, 0xc0, 0x1b, 0x8b, 0x5b, 0x72, 0x97, 0x72, 0x41, 0x76, 0x61, 0x56, 0x24, 0xa1, 0xcc,
	0x50, 0xa6, 0x32, 0x5d, 0x47, 0x1b, 0x5b, 0x1d, 0xbb, 0x5f, 0x5b, 0x6f, 0xea, 0xc9, 0x8f, 0x48,
	0xb2, 0xb9, 0x41, 0x48, 0x7b, 0x94, 0x1f, 0xc5, 0x2b, 0x63, 0x46, 0xa4, 0x7e, 0xc9, 0x9e, 0x43,
	0xd6, 0xe1, 0x7a, 0x5e, 0xc6, 0x0c, 0x22, 0xcf, 0xd3, 0x3e, 0x4e, 0x4a, 0xb1, 0xaa, 0x41, 0xb2,
	0xf8, 0x7d, 0xb9, 0x45, 0x8e, 0x60, 0x29, 0xcb, 0x71, 0xa8, 0xb0, 0xb9, 0xdb, 0x8f, 0xc3, 0x6a,
	0xe3, 0x2a, 0x43, 0x23, 0x97, 0x41, 0x84, 0x5c, 0xa2, 0x31, 0xc2, 0x8d, 0xb4, 0xe4, 0xce, 0x88,
	0x4a, 0x76, 0x60, 0xe5, 0x4a, 0x5d, 0x0c, 0xf5, 0x09, 0x43, 0x35, 0xaf, 0x50, 0x50, 0xe9, 0x36,
	0x61, 0x8a, 0xd3, 0x5e, 0x4c, 0x91, 0x13, 0x29, 0x15, 0x4f, 0xa4, 0x8a, 0x68, 0x39, 0x8b, 0xbb,
	0x30, 0x33, 0x64, 0xa2, 0xdf, 0x67, 0xf4, 0xbb, 0x36, 0x80, 0x28, 0x83, 0x6d, 0xa8, 0x77, 0x78,
	0x14, 0x86, 0xcc, 0x14, 0x91, 0x6f, 0x76, 0x39, 0xf3, 0xb5, 0x89, 0x4b, 0x6d, 0x1c, 0x6a, 0xbb,
	0xbe, 0xe5, 0xa1, 0xcd, 0x34, 0x72, 0x0e, 0x22, 0xff, 0xa5, 0x64, 0x90, 0x47, 0xb0, 0x90, 0x13,
	0x41, 0xc7, 0x2f, 0xe8, 0x38, 0x97, 0x41, 0x2b, 0xdb, 0xe7, 0x30, 0x9d, 0x62, 0x84, 0x4c, 0x2b,
	0x17, 0x9b, 0xd6, 0x86, 0x32, 0x87, 0x8c, 0x3c, 0x04, 0x92, 0x11, 0x40, 0xc3, 0xaf, 0x68, 0x58,
	0x4f, 0x21, 0x95, 0xdd, 0x33, 0xa8, 0xd9, 0x11, 0xe7, 0x34, 0xb0, 0xcf, 0xe2, 0x41, 0x56, 0x8a,
	0x07, 0x09, 0x03, 0xbc, 0x1c, 0xe5, 0x03, 0x98, 0x4d, 0xb1, 0xd1, 0xe9, 0x1b, 0x3a, 0xcd, 0x8c,
	0x60, 0xca, 0xe8, 0x15, 0xcc, 0x89, 0x88, 0xdb, 0xc7, 0x16, 0xef, 0xd1, 0xe1, 0x9b, 0x3c, 0x59,
	0x6c, 0x57, 0x1f, 0xb2, 0x92, 0x57, 0xf9, 0x31, 0x2c, 0x5e, 0x10, 0x42, 0xe7, 0xef, 0xe8, 0x3c,
	0x9f, 0x63, 0x28, 0xfb, 0x16, 0x4c, 0xd8, 0x4c, 0x84, 0x5a, 0xb5, 0x78, 0x9a, 0x0a, 0x48, 0x96,
	0x61, 0x2a, 0xfe, 0x44, 0xe5, 0x1f, 0xa8, 0x5c, 0x8d, 0x9f, 0x28, 0xb9, 0x36, 0x2c, 0x38, 0xd4,
	0x93, 0x47, 0xd4, 0xcc, 0xbc, 0xca, 0xda, 0xd4, 0xa5, 0xf2, 0x1d, 0xc6, 0x3c, 0x6a, 0x05, 0x28,
	0x4f, 0x90, 0x98, 0xae, 0x1a, 0xf2, 0x14, 0x96, 0x2e, 0x93, 0x43, 0xf3, 0x9f, 0x68, 0xbe, 0x78,
	0x91, 0xa7, 0xa2, 0x6c, 0x40, 0xc5, 0xb2, 0x43, 0xf7, 0x84, 0x6a, 0x50, 0x6c, 0x9e, 0x40, 0xc9,
	0x2a, 0xd4, 0xf0, 0x1b, 0x5a, 0xfc, 0x42, 0x0b, 0xc0, 0x67, 0x4a, 0xd6, 0x82, 0x5b, 0x7d, 0x4e,
	0x1d, 0xda, 0x35, 0x3b, 0xae, 0xe7, 0x99, 0x36, 0x93, 0xff, 0x94, 0x20, 0xcc, 0x74, 0x40, 0xad,
	0xb0, 0x03, 0x96, 0x51, 0x62, 0x4b, 0x2a, 0x6c, 0xa3, 0x40, 0xba, 0x08, 0x5e, 0xc3, 0x9d, 0x02,
	0x0b, 0x4c, 0xf7, 0x1b, 0xd3, 0xad, 0xfe, 0x4d, 0x2c, 0xce, 0xbc, 0x76, 0x3e, 0x0e, 0x55, 0x83,
	0x8a, 0x3e, 0x0b, 0x04, 0x95, 0x47, 0xaf, 0xac, 0x9a, 0x39, 0xa9, 0xcb, 0x61, 0xcc, 0xa4, 0xf3,
	0xb1, 0xb5, 0x77, 0xe3, 0xbf, 0x06, 0x02, 0xc9, 0x3b, 0x98, 0x8d, 0x3b, 0xd9, 0x4c, 0x95, 0xb2,
	0xec, 0xb9, 0x92, 0x24, 0xeb, 0x39, 0x72, 0xbe, 0xba, 0xdb, 0x72, 0xbd, 0x37, 0x5a, 0x1b, 0x75,
	0x3f, 0xfb, 0x40, 0xb6, 0xd5, 0x64, 0x72, 0x17, 0xc8, 0xae, 0x8a, 0x15, 0x57, 0x2e, 0x28, 0xe2,
	0x4d, 0xd1, 0xc6, 0x4f, 0x63, 0x00, 0x97, 0xe7, 0xa6, 0xc4, 0xd9, 0xa9, 0xac, 0x9e, 0x98, 0xf5,
	0x44, 0xff, 0xc7, 0x8b, 0x4b, 0x1f, 0x8c, 0x41, 0x37, 0xd8, 0xa9, 0x11, 0x2b, 0x34, 0x6e, 0x42,
	0x49, 0x7e, 0x27, 0x8b, 0x50, 0x91, 0xab, 0xf8, 0xf0, 0x7d, 0xd8, 0x97, 0x83, 0x29, 0x1b, 0x65,
	0xb9, 0xdc, 0x73, 0xb6, 0xde, 0x42, 0xd3, 0x65, 0x39, 0xf9, 0xd1, 0x8d, 0xfa, 0x7e, 0xfd, 0xff,
	0xef, 0xda, 0x4e, 0x45, 0xdd, 0x68, 0x1b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xeb, 0xc8,
	0xaf, 0xa8, 0x07, 0x00, 0x00,
}