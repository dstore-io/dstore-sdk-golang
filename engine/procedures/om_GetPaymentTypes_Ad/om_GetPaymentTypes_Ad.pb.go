// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPaymentTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetPaymentTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPaymentTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPaymentTypes_Ad

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
	LanguageId             *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull         bool                          `protobuf:"varint,1001,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	PaymentTypeId          *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PaymentTypeIdNull      bool                          `protobuf:"varint,1002,opt,name=payment_type_id_null,json=paymentTypeIdNull" json:"payment_type_id_null,omitempty"`
	OnlyActive             *dstore_values.BooleanValue   `protobuf:"bytes,3,opt,name=only_active,json=onlyActive" json:"only_active,omitempty"`
	OnlyActiveNull         bool                          `protobuf:"varint,1003,opt,name=only_active_null,json=onlyActiveNull" json:"only_active_null,omitempty"`
	ValidAtDateAndTime     *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=valid_at_date_and_time,json=validAtDateAndTime" json:"valid_at_date_and_time,omitempty"`
	ValidAtDateAndTimeNull bool                          `protobuf:"varint,1004,opt,name=valid_at_date_and_time_null,json=validAtDateAndTimeNull" json:"valid_at_date_and_time_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

func (m *Parameters) GetOnlyActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyActive
	}
	return nil
}

func (m *Parameters) GetValidAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidAtDateAndTime
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
	RowId                    int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BruttoSumTo              *dstore_values.DecimalValue   `protobuf:"bytes,10001,opt,name=brutto_sum_to,json=bruttoSumTo" json:"brutto_sum_to,omitempty"`
	PriorityNo               *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=priority_no,json=priorityNo" json:"priority_no,omitempty"`
	SurchargeUnitSymbol      *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=surcharge_unit_symbol,json=surchargeUnitSymbol" json:"surcharge_unit_symbol,omitempty"`
	BruttoSumFrom            *dstore_values.DecimalValue   `protobuf:"bytes,10004,opt,name=brutto_sum_from,json=bruttoSumFrom" json:"brutto_sum_from,omitempty"`
	SurchargeValue           *dstore_values.DecimalValue   `protobuf:"bytes,10005,opt,name=surcharge_value,json=surchargeValue" json:"surcharge_value,omitempty"`
	Cost                     *dstore_values.DecimalValue   `protobuf:"bytes,10006,opt,name=cost" json:"cost,omitempty"`
	GrossSumTo               *dstore_values.DecimalValue   `protobuf:"bytes,10007,opt,name=gross_sum_to,json=grossSumTo" json:"gross_sum_to,omitempty"`
	CurrencyId               *dstore_values.IntegerValue   `protobuf:"bytes,10008,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencySymbol           *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=currency_symbol,json=currencySymbol" json:"currency_symbol,omitempty"`
	CreatedAtDateAndTime     *dstore_values.TimestampValue `protobuf:"bytes,10010,opt,name=created_at_date_and_time,json=createdAtDateAndTime" json:"created_at_date_and_time,omitempty"`
	Active                   *dstore_values.BooleanValue   `protobuf:"bytes,10011,opt,name=active" json:"active,omitempty"`
	CostCurrencyId           *dstore_values.IntegerValue   `protobuf:"bytes,10012,opt,name=cost_currency_id,json=costCurrencyId" json:"cost_currency_id,omitempty"`
	PaymentTypeDescription   *dstore_values.StringValue    `protobuf:"bytes,10013,opt,name=payment_type_description,json=paymentTypeDescription" json:"payment_type_description,omitempty"`
	TranslatedDescription    *dstore_values.StringValue    `protobuf:"bytes,10014,opt,name=translated_description,json=translatedDescription" json:"translated_description,omitempty"`
	GrossSumFrom             *dstore_values.DecimalValue   `protobuf:"bytes,10015,opt,name=gross_sum_from,json=grossSumFrom" json:"gross_sum_from,omitempty"`
	Region                   *dstore_values.StringValue    `protobuf:"bytes,10016,opt,name=region" json:"region,omitempty"`
	SurchargeIsAbsoluteValue *dstore_values.IntegerValue   `protobuf:"bytes,10017,opt,name=surcharge_is_absolute_value,json=surchargeIsAbsoluteValue" json:"surcharge_is_absolute_value,omitempty"`
	CostCurrencySymbol       *dstore_values.StringValue    `protobuf:"bytes,10018,opt,name=cost_currency_symbol,json=costCurrencySymbol" json:"cost_currency_symbol,omitempty"`
	RegionId                 *dstore_values.IntegerValue   `protobuf:"bytes,10019,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	SurchargeTypeDescription *dstore_values.StringValue    `protobuf:"bytes,10020,opt,name=surcharge_type_description,json=surchargeTypeDescription" json:"surcharge_type_description,omitempty"`
	PaymentTypeId            *dstore_values.IntegerValue   `protobuf:"bytes,10021,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	SurchargeTypeId          *dstore_values.IntegerValue   `protobuf:"bytes,10022,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeUnitId          *dstore_values.IntegerValue   `protobuf:"bytes,10023,opt,name=surcharge_unit_id,json=surchargeUnitId" json:"surcharge_unit_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBruttoSumTo() *dstore_values.DecimalValue {
	if m != nil {
		return m.BruttoSumTo
	}
	return nil
}

func (m *Response_Row) GetPriorityNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PriorityNo
	}
	return nil
}

func (m *Response_Row) GetSurchargeUnitSymbol() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeUnitSymbol
	}
	return nil
}

func (m *Response_Row) GetBruttoSumFrom() *dstore_values.DecimalValue {
	if m != nil {
		return m.BruttoSumFrom
	}
	return nil
}

func (m *Response_Row) GetSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.SurchargeValue
	}
	return nil
}

func (m *Response_Row) GetCost() *dstore_values.DecimalValue {
	if m != nil {
		return m.Cost
	}
	return nil
}

func (m *Response_Row) GetGrossSumTo() *dstore_values.DecimalValue {
	if m != nil {
		return m.GrossSumTo
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

func (m *Response_Row) GetCreatedAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.CreatedAtDateAndTime
	}
	return nil
}

func (m *Response_Row) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Response_Row) GetCostCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CostCurrencyId
	}
	return nil
}

func (m *Response_Row) GetPaymentTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentTypeDescription
	}
	return nil
}

func (m *Response_Row) GetTranslatedDescription() *dstore_values.StringValue {
	if m != nil {
		return m.TranslatedDescription
	}
	return nil
}

func (m *Response_Row) GetGrossSumFrom() *dstore_values.DecimalValue {
	if m != nil {
		return m.GrossSumFrom
	}
	return nil
}

func (m *Response_Row) GetRegion() *dstore_values.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Response_Row) GetSurchargeIsAbsoluteValue() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeIsAbsoluteValue
	}
	return nil
}

func (m *Response_Row) GetCostCurrencySymbol() *dstore_values.StringValue {
	if m != nil {
		return m.CostCurrencySymbol
	}
	return nil
}

func (m *Response_Row) GetRegionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RegionId
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeTypeDescription
	}
	return nil
}

func (m *Response_Row) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Response_Row) GetSurchargeUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeUnitId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPaymentTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPaymentTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPaymentTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetPaymentTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 940 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0x56, 0x71, 0x93, 0x86, 0x49, 0x63, 0xa7, 0xd3, 0xd4, 0x5a, 0x6c, 0x81, 0x50, 0xfb, 0x02,
	0x2f, 0x1b, 0xa0, 0x95, 0x00, 0x09, 0x10, 0x6e, 0xd3, 0x42, 0x1e, 0xb2, 0x94, 0x6d, 0x1a, 0xa9,
	0xf0, 0xb0, 0x1a, 0xef, 0x4e, 0xcd, 0x8a, 0xdd, 0x19, 0x6b, 0x66, 0xb6, 0x95, 0x7f, 0x02, 0x6f,
	0xdc, 0xef, 0xd7, 0x72, 0xfb, 0x67, 0x48, 0x5c, 0x7e, 0x04, 0x67, 0x2e, 0xeb, 0xbd, 0xc4, 0x92,
	0xb7, 0x2f, 0x89, 0x66, 0xe7, 0x7c, 0xdf, 0xb9, 0x7d, 0xe7, 0x78, 0xd0, 0xb5, 0x44, 0x2a, 0x2e,
	0xe8, 0x3e, 0x65, 0xb3, 0x94, 0xd1, 0xfd, 0xb9, 0xe0, 0x31, 0x4d, 0x0a, 0x41, 0xe5, 0x3e, 0xcf,
	0xa3, 0xb7, 0xa8, 0xba, 0x4d, 0x16, 0x39, 0x65, 0xea, 0x78, 0x31, 0xa7, 0x32, 0x9a, 0x24, 0x3e,
	0x18, 0x28, 0x8e, 0xaf, 0x58, 0x94, 0x6f, 0x51, 0xfe, 0x4a, 0xd3, 0xd1, 0x45, 0x47, 0xfd, 0x80,
	0x64, 0x05, 0x95, 0x16, 0x39, 0x7a, 0xaa, 0xe9, 0x8f, 0x0a, 0xc1, 0x85, 0xbb, 0x1a, 0x37, 0xaf,
	0x72, 0x2a, 0x25, 0x99, 0x51, 0x77, 0x79, 0xa5, 0x7d, 0xa9, 0x48, 0xca, 0xee, 0x73, 0x91, 0x13,
	0x95, 0x72, 0x66, 0x8d, 0x2e, 0xff, 0xd5, 0x43, 0xe8, 0x36, 0x11, 0x04, 0x6e, 0xa9, 0x90, 0xf8,
	0x35, 0xb4, 0x9d, 0x11, 0x36, 0x2b, 0x80, 0x25, 0x4a, 0x13, 0xef, 0xcc, 0xb3, 0x67, 0x9e, 0xdb,
	0x7e, 0x69, 0xec, 0xbb, 0xd8, 0x5d, 0x58, 0x29, 0x53, 0x74, 0x46, 0xc5, 0x89, 0x3e, 0x85, 0xa8,
	0xb4, 0x3f, 0x4c, 0xf0, 0xf3, 0x68, 0xb7, 0x86, 0x8e, 0x58, 0x91, 0x65, 0xde, 0xdf, 0xe7, 0x80,
	0x63, 0x2b, 0xec, 0x57, 0x66, 0x01, 0x7c, 0xc6, 0x37, 0xd0, 0x60, 0x6e, 0x93, 0x8f, 0x14, 0x64,
	0xaf, 0x9d, 0x3d, 0xb1, 0xde, 0xd9, 0xce, 0xbc, 0x2a, 0x18, 0xf8, 0x7b, 0x01, 0xed, 0xb5, 0x48,
	0xac, 0xcf, 0x7f, 0xac, 0xcf, 0x0b, 0x0d, 0x6b, 0xe3, 0x16, 0xf2, 0xe3, 0x2c, 0x5b, 0x44, 0x24,
	0x56, 0xe9, 0x03, 0xea, 0xf5, 0x56, 0xba, 0x9c, 0x72, 0x9e, 0x51, 0xc2, 0x5c, 0x7e, 0xda, 0x7e,
	0x62, 0xcc, 0x75, 0x7e, 0x35, 0xb4, 0xf5, 0xf5, 0xaf, 0xcb, 0xaf, 0x32, 0x33, 0x8e, 0xde, 0x45,
	0x43, 0x60, 0x83, 0x80, 0x88, 0x8a, 0x12, 0xa2, 0x68, 0x44, 0x58, 0x12, 0xa9, 0x34, 0xa7, 0xde,
	0x59, 0xe3, 0xf3, 0xe9, 0x96, 0x4f, 0x7d, 0x25, 0x15, 0xc9, 0xe7, 0xd6, 0x2b, 0x36, 0xe0, 0x89,
	0x3a, 0x00, 0xe8, 0x84, 0x25, 0xc7, 0x70, 0x0b, 0xb1, 0x8f, 0x57, 0x53, 0xda, 0x40, 0xfe, 0xb3,
	0x81, 0x0c, 0x4f, 0x23, 0x75, 0x40, 0x97, 0x1f, 0x0d, 0xd0, 0x56, 0x48, 0xe5, 0x9c, 0x33, 0x49,
	0xa1, 0x70, 0x1b, 0x46, 0x46, 0xae, 0xc1, 0x23, 0xbf, 0x29, 0x4e, 0x2b, 0xb1, 0x9b, 0xfa, 0x6f,
	0x68, 0x0d, 0xf1, 0x3d, 0xb4, 0xab, 0x05, 0x14, 0xd5, 0x14, 0x04, 0x0d, 0xeb, 0x01, 0xd8, 0x6f,
	0x81, 0xdb, 0x3a, 0x3b, 0x82, 0xf3, 0x61, 0x75, 0x0e, 0x07, 0x79, 0xf3, 0x03, 0x7e, 0x05, 0x9d,
	0x73, 0xc2, 0x85, 0x7e, 0x68, 0xc6, 0x67, 0x4e, 0x31, 0x5a, 0x59, 0x1f, 0xd9, 0xff, 0x61, 0x69,
	0x0e, 0x22, 0xea, 0x09, 0xfe, 0x10, 0x2a, 0xaa, 0x51, 0x2f, 0xfa, 0x1d, 0x26, 0xcc, 0x2f, 0x4b,
	0xe0, 0x87, 0xfc, 0x61, 0xa8, 0xd1, 0xa3, 0x8f, 0x76, 0x50, 0x0f, 0x0e, 0x78, 0x88, 0x36, 0xe1,
	0xa8, 0x85, 0xf8, 0x71, 0x00, 0x55, 0xd9, 0x08, 0x37, 0xe0, 0x08, 0x22, 0x7b, 0x13, 0xed, 0x4c,
	0x45, 0xa1, 0x14, 0x8f, 0x64, 0x91, 0x47, 0x8a, 0x7b, 0x9f, 0x04, 0x2b, 0x55, 0x93, 0xd0, 0x38,
	0xcd, 0x49, 0x66, 0xfb, 0xb7, 0x6d, 0x21, 0x77, 0x8a, 0xfc, 0x98, 0xe3, 0xd7, 0xd1, 0xf6, 0x5c,
	0xa4, 0x5c, 0xa4, 0x6a, 0x11, 0x31, 0xee, 0x7d, 0x1a, 0x74, 0x98, 0xaa, 0x12, 0x10, 0x70, 0xfc,
	0x0e, 0xba, 0x24, 0x0b, 0x11, 0x7f, 0x40, 0x04, 0x8c, 0x55, 0xc1, 0x52, 0x15, 0xc9, 0x45, 0x3e,
	0xe5, 0x99, 0xf7, 0x59, 0xd0, 0xec, 0x9e, 0x23, 0x92, 0x4a, 0xa4, 0x6c, 0x66, 0x79, 0x2e, 0x2e,
	0x91, 0x77, 0x01, 0x78, 0xc7, 0xe0, 0xf0, 0x01, 0x1a, 0xd4, 0x32, 0xba, 0x2f, 0x78, 0xee, 0x7d,
	0xde, 0x21, 0xa7, 0x9d, 0x65, 0x4e, 0xb7, 0x00, 0x82, 0x6f, 0xa2, 0x41, 0x15, 0x96, 0xb1, 0xf7,
	0xbe, 0xe8, 0xc0, 0xd2, 0x5f, 0x82, 0xcc, 0x19, 0xa4, 0x78, 0x36, 0xe6, 0x52, 0x79, 0x5f, 0x76,
	0xc0, 0x1a, 0x4b, 0xfc, 0x06, 0x3a, 0x3f, 0x13, 0x5c, 0xca, 0xb2, 0x1f, 0x5f, 0x75, 0x40, 0x22,
	0x83, 0x58, 0xb6, 0x23, 0x2e, 0x84, 0xa0, 0x2c, 0x5e, 0xe8, 0x6e, 0x7f, 0xdd, 0xa5, 0x1d, 0x25,
	0x00, 0xf4, 0x00, 0xd5, 0x5b, 0xc2, 0x5d, 0x23, 0xbe, 0x59, 0xdf, 0x88, 0x7e, 0x89, 0x71, 0x3d,
	0x38, 0x41, 0x5e, 0x2c, 0x28, 0xcc, 0xe8, 0x8a, 0x0d, 0xf1, 0x6d, 0xd0, 0x65, 0x45, 0xec, 0x39,
	0x7c, 0x73, 0x49, 0x5c, 0x43, 0x9b, 0x6e, 0xb7, 0x7d, 0x17, 0xac, 0x5f, 0x6e, 0xce, 0x16, 0xdf,
	0x42, 0xbb, 0xba, 0xb4, 0x51, 0xbd, 0x2e, 0xdf, 0x77, 0xa8, 0x4b, 0x5f, 0xa3, 0x6e, 0x54, 0xb5,
	0xb9, 0x8b, 0xbc, 0xc6, 0x42, 0x4e, 0xa8, 0x8c, 0x45, 0x3a, 0x37, 0xdb, 0xe2, 0x87, 0xf5, 0x45,
	0x1a, 0xd6, 0x16, 0xf6, 0x41, 0x05, 0xc5, 0x21, 0x1a, 0x2a, 0x41, 0x98, 0xcc, 0x4c, 0xbd, 0xea,
	0xa4, 0x3f, 0xae, 0x27, 0xbd, 0x54, 0x41, 0xeb, 0x9c, 0xd7, 0x51, 0xbf, 0x52, 0x91, 0x99, 0x81,
	0x9f, 0x3a, 0xe8, 0xe8, 0x7c, 0xa9, 0x23, 0x33, 0x02, 0x57, 0x61, 0x65, 0xd0, 0x99, 0x8e, 0xe3,
	0xe7, 0xf5, 0x71, 0x38, 0x53, 0xfc, 0x3e, 0x1a, 0x57, 0x73, 0x93, 0xca, 0x88, 0x4c, 0x25, 0xcf,
	0x0a, 0x55, 0xce, 0xd0, 0x2f, 0x1d, 0xca, 0xee, 0x2d, 0x09, 0x0e, 0xe5, 0xc4, 0xc1, 0xed, 0x34,
	0x1d, 0xa1, 0xbd, 0x66, 0x23, 0x9d, 0x42, 0x1f, 0xad, 0x8f, 0x0f, 0xd7, 0x7b, 0xe9, 0x54, 0xfa,
	0x2a, 0x7a, 0xd2, 0x46, 0xad, 0x05, 0xf1, 0x6b, 0x87, 0xc8, 0xb6, 0xac, 0x39, 0x48, 0xe1, 0x1e,
	0x1a, 0x55, 0x69, 0x9e, 0x12, 0xc3, 0x6f, 0xeb, 0xe3, 0xa9, 0x92, 0x6c, 0xcb, 0xe1, 0xe0, 0xf4,
	0xdb, 0xe1, 0xf7, 0xe0, 0xb1, 0x1f, 0x0f, 0x6f, 0xa3, 0x0b, 0xad, 0x00, 0x81, 0xe7, 0x8f, 0x0e,
	0x3c, 0x83, 0x46, 0x60, 0x6d, 0x26, 0xb3, 0xa0, 0x81, 0xe9, 0xcf, 0xc7, 0x62, 0xd2, 0xdb, 0xf9,
	0x30, 0xb9, 0x7e, 0x82, 0xc6, 0x29, 0x6f, 0xfd, 0x8e, 0x55, 0xef, 0xcb, 0xf7, 0x5e, 0x9e, 0x71,
	0x99, 0x7c, 0x58, 0xde, 0x27, 0x9d, 0x9f, 0xa0, 0xd3, 0x4d, 0xf3, 0xd8, 0xbb, 0xfa, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0x4c, 0x20, 0x6b, 0xbb, 0x0a, 0x00, 0x00,
}
