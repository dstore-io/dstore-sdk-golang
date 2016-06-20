// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderSurcharges_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderSurcharges_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderSurcharges_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderSurcharges_Ad

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
	OrderId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderIdNull      bool                        `protobuf:"varint,1001,opt,name=order_id_null,json=orderIdNull" json:"order_id_null,omitempty"`
	SplitByTaxes     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=split_by_taxes,json=splitByTaxes" json:"split_by_taxes,omitempty"`
	SplitByTaxesNull bool                        `protobuf:"varint,1002,opt,name=split_by_taxes_null,json=splitByTaxesNull" json:"split_by_taxes_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *Parameters) GetSplitByTaxes() *dstore_values.IntegerValue {
	if m != nil {
		return m.SplitByTaxes
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
	RowId                    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	OriginalSurchargeValue   *dstore_values.DecimalValue `protobuf:"bytes,10001,opt,name=original_surcharge_value,json=originalSurchargeValue" json:"original_surcharge_value,omitempty"`
	OrigSurchValUnitId       *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=orig_surch_val_unit_id,json=origSurchValUnitId" json:"orig_surch_val_unit_id,omitempty"`
	OrigSurchValUnitSymbol   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=orig_surch_val_unit_symbol,json=origSurchValUnitSymbol" json:"orig_surch_val_unit_symbol,omitempty"`
	AppliedOnNetSum          *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=applied_on_net_sum,json=appliedOnNetSum" json:"applied_on_net_sum,omitempty"`
	TaxesMultiplier          *dstore_values.DecimalValue `protobuf:"bytes,10005,opt,name=taxes_multiplier,json=taxesMultiplier" json:"taxes_multiplier,omitempty"`
	AbsoluteNetSurcharge     *dstore_values.DecimalValue `protobuf:"bytes,10006,opt,name=absolute_net_surcharge,json=absoluteNetSurcharge" json:"absolute_net_surcharge,omitempty"`
	PositionNo               *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=position_no,json=positionNo" json:"position_no,omitempty"`
	AppliedOnGrossSum        *dstore_values.DecimalValue `protobuf:"bytes,10008,opt,name=applied_on_gross_sum,json=appliedOnGrossSum" json:"applied_on_gross_sum,omitempty"`
	CurrencyId               *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencySymbol           *dstore_values.StringValue  `protobuf:"bytes,10010,opt,name=currency_symbol,json=currencySymbol" json:"currency_symbol,omitempty"`
	OrigSurchValIsAbsolute   *dstore_values.IntegerValue `protobuf:"bytes,10011,opt,name=orig_surch_val_is_absolute,json=origSurchValIsAbsolute" json:"orig_surch_val_is_absolute,omitempty"`
	AbsoluteGrossSurcharge   *dstore_values.DecimalValue `protobuf:"bytes,10012,opt,name=absolute_gross_surcharge,json=absoluteGrossSurcharge" json:"absolute_gross_surcharge,omitempty"`
	CategoryDescription      *dstore_values.StringValue  `protobuf:"bytes,10013,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	SurchargeTypeDescription *dstore_values.StringValue  `protobuf:"bytes,10014,opt,name=surcharge_type_description,json=surchargeTypeDescription" json:"surcharge_type_description,omitempty"`
	SurchargeTypeId          *dstore_values.IntegerValue `protobuf:"bytes,10015,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeCategoryId  *dstore_values.IntegerValue `protobuf:"bytes,10016,opt,name=surcharge_type_category_id,json=surchargeTypeCategoryId" json:"surcharge_type_category_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetOriginalSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.OriginalSurchargeValue
	}
	return nil
}

func (m *Response_Row) GetOrigSurchValUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrigSurchValUnitId
	}
	return nil
}

func (m *Response_Row) GetOrigSurchValUnitSymbol() *dstore_values.StringValue {
	if m != nil {
		return m.OrigSurchValUnitSymbol
	}
	return nil
}

func (m *Response_Row) GetAppliedOnNetSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.AppliedOnNetSum
	}
	return nil
}

func (m *Response_Row) GetTaxesMultiplier() *dstore_values.DecimalValue {
	if m != nil {
		return m.TaxesMultiplier
	}
	return nil
}

func (m *Response_Row) GetAbsoluteNetSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteNetSurcharge
	}
	return nil
}

func (m *Response_Row) GetPositionNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PositionNo
	}
	return nil
}

func (m *Response_Row) GetAppliedOnGrossSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.AppliedOnGrossSum
	}
	return nil
}

func (m *Response_Row) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Response_Row) GetCurrencySymbol() *dstore_values.StringValue {
	if m != nil {
		return m.CurrencySymbol
	}
	return nil
}

func (m *Response_Row) GetOrigSurchValIsAbsolute() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrigSurchValIsAbsolute
	}
	return nil
}

func (m *Response_Row) GetAbsoluteGrossSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteGrossSurcharge
	}
	return nil
}

func (m *Response_Row) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeTypeDescription
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeCategoryId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderSurcharges_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderSurcharges_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderSurcharges_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x96, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x86, 0xd5, 0x9d, 0xdd, 0x83, 0xa6, 0xdd, 0x4d, 0x3b, 0xad, 0xba, 0x4d, 0x2a, 0x21, 0xd4,
	0x4a, 0x88, 0x2b, 0x97, 0x93, 0x80, 0x1b, 0x2e, 0x5a, 0x0a, 0x25, 0x17, 0x71, 0x2b, 0xb7, 0x94,
	0xc3, 0xcd, 0xc8, 0x89, 0x87, 0x60, 0xc9, 0x9e, 0x89, 0x66, 0x6c, 0x4a, 0xde, 0x82, 0xf3, 0xb9,
	0xc0, 0x6b, 0xf0, 0x2e, 0xbc, 0x00, 0x3c, 0x05, 0x6b, 0xc6, 0x33, 0x4e, 0x9c, 0x16, 0xc5, 0xdc,
	0xb4, 0x1a, 0xcf, 0x7c, 0xff, 0x5a, 0xfe, 0xff, 0x65, 0x3b, 0xe8, 0x7a, 0x28, 0x53, 0x2e, 0xe8,
	0x06, 0x65, 0xdd, 0x88, 0xd1, 0x8d, 0x9e, 0xe0, 0x1d, 0x1a, 0x66, 0x82, 0xca, 0x0d, 0x9e, 0x90,
	0x1d, 0x9a, 0xee, 0x8a, 0x90, 0x8a, 0xfd, 0x4c, 0x74, 0x9e, 0x04, 0xa2, 0x4b, 0x25, 0xd9, 0x0c,
	0x5d, 0x38, 0x93, 0x72, 0x7c, 0x3e, 0x07, 0xdd, 0x1c, 0x74, 0xff, 0x74, 0xba, 0xb1, 0x64, 0x0a,
	0x3c, 0x0d, 0xe2, 0x8c, 0xca, 0x1c, 0x6e, 0x9c, 0x29, 0x57, 0xa5, 0x42, 0x70, 0x61, 0xb6, 0x56,
	0xcb, 0x5b, 0x09, 0x95, 0x32, 0xe8, 0x52, 0xb3, 0xb9, 0x3e, 0xba, 0x99, 0x06, 0x11, 0x7b, 0xcc,
	0x45, 0x12, 0xa4, 0x11, 0x67, 0xf9, 0xa1, 0xb5, 0x1f, 0x13, 0x08, 0xed, 0x05, 0x22, 0x80, 0x5d,
	0x2a, 0x24, 0xbe, 0x86, 0x66, 0xb8, 0x6a, 0x8b, 0x44, 0xa1, 0x33, 0x71, 0x6e, 0xe2, 0xc2, 0xec,
	0xe5, 0x55, 0xd7, 0xf4, 0x6e, 0x7a, 0x8a, 0x58, 0x4a, 0xbb, 0x54, 0x1c, 0xaa, 0x95, 0x3f, 0xad,
	0x0f, 0x37, 0x43, 0xbc, 0x8e, 0xfe, 0xb3, 0x1c, 0x61, 0x59, 0x1c, 0x3b, 0x3f, 0xa7, 0x81, 0x9e,
	0xf1, 0x67, 0xcd, 0x01, 0x0f, 0xae, 0xe1, 0x4d, 0x34, 0x2f, 0x7b, 0x71, 0x94, 0x92, 0x76, 0x9f,
	0xa4, 0xc1, 0x33, 0x2a, 0x9d, 0x7f, 0xc6, 0x97, 0x98, 0xd3, 0xc8, 0x56, 0xff, 0x40, 0x01, 0xd8,
	0x45, 0x4b, 0x65, 0x89, 0xbc, 0xda, 0xaf, 0xbc, 0xda, 0xc2, 0xf0, 0x59, 0x55, 0x72, 0xed, 0xfb,
	0x1c, 0x9a, 0xf1, 0xa9, 0xec, 0x71, 0x26, 0x29, 0xbe, 0x88, 0x26, 0xb5, 0x79, 0xe6, 0xce, 0x1a,
	0x6e, 0x39, 0x95, 0xdc, 0xd8, 0xdb, 0xea, 0xaf, 0x9f, 0x1f, 0xc4, 0x0f, 0xd1, 0x82, 0xb2, 0x8d,
	0x0c, 0xf9, 0x06, 0x3d, 0xd7, 0x00, 0x76, 0x47, 0xe0, 0x51, 0x77, 0x5b, 0xb0, 0x6e, 0x0e, 0xd6,
	0x7e, 0x3d, 0x29, 0x5f, 0xc0, 0x37, 0xd0, 0xb4, 0x89, 0xcb, 0xa9, 0x69, 0xc5, 0xb3, 0x27, 0x14,
	0xf3, 0x30, 0x5b, 0xf9, 0x7f, 0xdf, 0x1e, 0xc7, 0x77, 0x50, 0x4d, 0xf0, 0x23, 0xe7, 0x5f, 0x4d,
	0x5d, 0x75, 0xab, 0x8d, 0x96, 0x6b, 0x5d, 0x70, 0x7d, 0x7e, 0xe4, 0x2b, 0x81, 0xc6, 0x31, 0x42,
	0x35, 0x58, 0xe0, 0x15, 0x34, 0x05, 0x4b, 0x95, 0xf8, 0x73, 0x0f, 0x8c, 0x99, 0xf4, 0x27, 0x61,
	0x09, 0x99, 0x1e, 0x22, 0x87, 0x8b, 0x08, 0x54, 0x83, 0x98, 0x48, 0xab, 0x45, 0x74, 0x46, 0xce,
	0x0b, 0xef, 0xd4, 0xe4, 0x42, 0xda, 0x89, 0x92, 0x20, 0xce, 0x93, 0x5b, 0xb1, 0x74, 0xd1, 0x88,
	0xbe, 0x8e, 0xf7, 0x90, 0xde, 0xc9, 0x35, 0x95, 0x1e, 0xc9, 0x18, 0x24, 0x0a, 0xf5, 0x5f, 0x7a,
	0xe3, 0xe7, 0x01, 0x2b, 0x56, 0x2b, 0xc2, 0xfa, 0x1e, 0x80, 0xd0, 0xe9, 0x7d, 0xd4, 0x38, 0x4d,
	0x51, 0xf6, 0x93, 0x36, 0x8f, 0x9d, 0x57, 0x5e, 0x39, 0x6e, 0xa3, 0x2a, 0x53, 0x11, 0xb1, 0xee,
	0x50, 0xab, 0xc3, 0xa2, 0xfb, 0x1a, 0xc5, 0x4d, 0x84, 0x83, 0x1e, 0xcc, 0x14, 0x0d, 0x09, 0x67,
	0x84, 0x51, 0xd0, 0xcc, 0x12, 0xe7, 0x75, 0x85, 0x9b, 0xaf, 0x1b, 0x6e, 0x97, 0x79, 0x34, 0xdd,
	0xcf, 0x12, 0xbc, 0x83, 0x16, 0xf2, 0x81, 0x4d, 0xb2, 0x38, 0x8d, 0xd4, 0x9e, 0x70, 0xde, 0x54,
	0x11, 0xd2, 0x54, 0xab, 0x80, 0xb0, 0x8f, 0x56, 0x82, 0xb6, 0xe4, 0x71, 0x96, 0x52, 0xd3, 0x91,
	0x71, 0xd7, 0x79, 0x5b, 0x41, 0x6e, 0xd9, 0xb2, 0xba, 0x2d, 0x43, 0xe2, 0x9b, 0x68, 0xb6, 0xc7,
	0x65, 0xa4, 0x06, 0x93, 0x30, 0xee, 0xbc, 0xab, 0x90, 0x03, 0xb2, 0x80, 0xc7, 0x71, 0x0b, 0x2d,
	0x0f, 0xd9, 0xd4, 0x15, 0x5c, 0x4a, 0x6d, 0xd4, 0xfb, 0x0a, 0x0d, 0x2d, 0x16, 0x46, 0xed, 0x28,
	0x4e, 0x59, 0x05, 0xdd, 0x74, 0x32, 0x21, 0x28, 0xeb, 0xf4, 0xd5, 0x54, 0x7c, 0xa8, 0xd2, 0x8d,
	0x05, 0x60, 0x1a, 0xb6, 0x51, 0xbd, 0xc0, 0xcd, 0x08, 0x7c, 0x1c, 0x3f, 0x02, 0xf3, 0x96, 0x31,
	0xd1, 0x3f, 0x38, 0x31, 0x53, 0x91, 0x24, 0xd6, 0x3c, 0xe7, 0x53, 0x85, 0x9e, 0x4a, 0x43, 0xd5,
	0x94, 0x9b, 0x86, 0x55, 0xcf, 0x55, 0x11, 0xa0, 0xf5, 0xca, 0x46, 0xf8, 0xb9, 0xca, 0x73, 0x65,
	0x69, 0x63, 0x98, 0x0d, 0xd1, 0x43, 0xcb, 0x9d, 0x00, 0xea, 0x73, 0xd1, 0x27, 0x21, 0x95, 0x1d,
	0x11, 0xf5, 0xf4, 0x0b, 0xeb, 0x78, 0xfc, 0xcd, 0x2f, 0x59, 0x70, 0x7b, 0xc0, 0xc1, 0xcb, 0xaf,
	0x31, 0x78, 0xec, 0xd3, 0x7e, 0x8f, 0x96, 0x54, 0xbf, 0x8c, 0x57, 0x75, 0x0a, 0xfc, 0x00, 0xe8,
	0x61, 0xe9, 0xbb, 0x68, 0x71, 0x44, 0x1a, 0x72, 0xfe, 0x5a, 0xc1, 0xd3, 0x7a, 0x49, 0x12, 0xc2,
	0x3e, 0xd9, 0x64, 0xe1, 0x01, 0x48, 0x7e, 0xab, 0x20, 0xf9, 0x7f, 0x49, 0xf2, 0x96, 0xa1, 0x9b,
	0xe1, 0xd6, 0x1e, 0x5a, 0x8d, 0xf8, 0xc8, 0xeb, 0x75, 0xf0, 0xc9, 0x7f, 0x74, 0xe9, 0xaf, 0x7f,
	0x0c, 0xb4, 0xa7, 0xf4, 0x37, 0xf7, 0xca, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x38, 0x8b,
	0xc5, 0x48, 0x08, 0x00, 0x00,
}