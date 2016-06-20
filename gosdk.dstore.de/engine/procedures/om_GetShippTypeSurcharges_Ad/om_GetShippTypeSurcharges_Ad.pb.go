// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetShippTypeSurcharges_Ad.proto
// DO NOT EDIT!

/*
Package om_GetShippTypeSurcharges_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetShippTypeSurcharges_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetShippTypeSurcharges_Ad

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
	ShippingTypeId      *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=shipping_type_id,json=shippingTypeId" json:"shipping_type_id,omitempty"`
	ShippingTypeIdNull  bool                          `protobuf:"varint,1001,opt,name=shipping_type_id_null,json=shippingTypeIdNull" json:"shipping_type_id_null,omitempty"`
	SurchargeTypeId     *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull bool                          `protobuf:"varint,1002,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	ValidFrom           *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ValidFromNull       bool                          `protobuf:"varint,1003,opt,name=valid_from_null,json=validFromNull" json:"valid_from_null,omitempty"`
	ValidTo             *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=valid_to,json=validTo" json:"valid_to,omitempty"`
	ValidToNull         bool                          `protobuf:"varint,1004,opt,name=valid_to_null,json=validToNull" json:"valid_to_null,omitempty"`
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

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetValidFrom() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *Parameters) GetValidTo() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidTo
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
	UnitId                   *dstore_values.IntegerValue   `protobuf:"bytes,10001,opt,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	ValidTo                  *dstore_values.TimestampValue `protobuf:"bytes,10002,opt,name=valid_to,json=validTo" json:"valid_to,omitempty"`
	PriorityNo               *dstore_values.IntegerValue   `protobuf:"bytes,10003,opt,name=priority_no,json=priorityNo" json:"priority_no,omitempty"`
	ShippingTypeDescription  *dstore_values.StringValue    `protobuf:"bytes,10004,opt,name=shipping_type_description,json=shippingTypeDescription" json:"shipping_type_description,omitempty"`
	ValidFrom                *dstore_values.TimestampValue `protobuf:"bytes,10005,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	SurchargeIsAbsoluteValue *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=surcharge_is_absolute_value,json=surchargeIsAbsoluteValue" json:"surcharge_is_absolute_value,omitempty"`
	UnitSymbol               *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=unit_symbol,json=unitSymbol" json:"unit_symbol,omitempty"`
	ShippingTypeId           *dstore_values.IntegerValue   `protobuf:"bytes,10008,opt,name=shipping_type_id,json=shippingTypeId" json:"shipping_type_id,omitempty"`
	SurchargeTypeDescription *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=surcharge_type_description,json=surchargeTypeDescription" json:"surcharge_type_description,omitempty"`
	SurchargeValue           *dstore_values.DecimalValue   `protobuf:"bytes,10010,opt,name=surcharge_value,json=surchargeValue" json:"surcharge_value,omitempty"`
	SurchargeTypeId          *dstore_values.IntegerValue   `protobuf:"bytes,10011,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitId
	}
	return nil
}

func (m *Response_Row) GetValidTo() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidTo
	}
	return nil
}

func (m *Response_Row) GetPriorityNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PriorityNo
	}
	return nil
}

func (m *Response_Row) GetShippingTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ShippingTypeDescription
	}
	return nil
}

func (m *Response_Row) GetValidFrom() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *Response_Row) GetSurchargeIsAbsoluteValue() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeIsAbsoluteValue
	}
	return nil
}

func (m *Response_Row) GetUnitSymbol() *dstore_values.StringValue {
	if m != nil {
		return m.UnitSymbol
	}
	return nil
}

func (m *Response_Row) GetShippingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ShippingTypeId
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeTypeDescription
	}
	return nil
}

func (m *Response_Row) GetSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.SurchargeValue
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetShippTypeSurcharges_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetShippTypeSurcharges_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetShippTypeSurcharges_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetShippTypeSurcharges_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x56, 0xff, 0xfc, 0x4d, 0xc2, 0x44, 0xa5, 0x65, 0x2b, 0x8a, 0x9b, 0x08, 0x84, 0xda, 0x0b,
	0x10, 0x17, 0x2e, 0x2a, 0x20, 0x15, 0x71, 0x90, 0x8a, 0x68, 0x4b, 0x91, 0x6a, 0x21, 0xb7, 0x02,
	0x15, 0x84, 0x2c, 0x37, 0xde, 0xa6, 0x2b, 0x62, 0xaf, 0xb5, 0xbb, 0xa1, 0xea, 0x5b, 0x70, 0x3e,
	0x3e, 0x07, 0x6f, 0xc1, 0x2d, 0x0f, 0xc1, 0xe1, 0x21, 0x98, 0xf5, 0xda, 0x89, 0xed, 0x56, 0x24,
	0xb9, 0x69, 0x35, 0x99, 0xef, 0x9b, 0x6f, 0x76, 0x4e, 0x86, 0x5b, 0x81, 0x54, 0x5c, 0xd0, 0x25,
	0x1a, 0x75, 0x58, 0x44, 0x97, 0x62, 0xc1, 0xdb, 0x34, 0xe8, 0x09, 0x2a, 0x97, 0x78, 0xe8, 0x6d,
	0x50, 0xb5, 0x7d, 0xc0, 0xe2, 0x78, 0xe7, 0x28, 0xa6, 0xdb, 0x3d, 0xd1, 0x3e, 0xf0, 0x45, 0x87,
	0x4a, 0x6f, 0x35, 0xb0, 0x11, 0xa7, 0x38, 0xb9, 0x62, 0xc8, 0xb6, 0x21, 0xdb, 0xff, 0x62, 0x34,
	0x67, 0x53, 0xa1, 0x97, 0x7e, 0xb7, 0x47, 0xa5, 0x09, 0xd0, 0x9c, 0x2f, 0xaa, 0x53, 0x21, 0xb8,
	0x48, 0x5d, 0xad, 0xa2, 0x2b, 0xa4, 0x52, 0xfa, 0x1d, 0x9a, 0x3a, 0x17, 0xcb, 0x4e, 0xe5, 0xb3,
	0x68, 0x9f, 0x8b, 0xd0, 0x57, 0x8c, 0x47, 0x06, 0xb4, 0xf0, 0xbd, 0x02, 0xf0, 0xc8, 0x17, 0x3e,
	0x7a, 0xa9, 0x90, 0x64, 0x0d, 0x66, 0xa4, 0x4e, 0x8d, 0x45, 0x1d, 0x4f, 0x61, 0x7a, 0x1e, 0x0b,
	0xac, 0x89, 0x8b, 0x13, 0x97, 0x1b, 0xcb, 0x2d, 0x3b, 0x7d, 0x47, 0x9a, 0x1b, 0x8b, 0x14, 0xed,
	0x50, 0xf1, 0x58, 0x5b, 0xee, 0xe9, 0x8c, 0xa4, 0x9f, 0xb4, 0x19, 0x90, 0x65, 0x38, 0x5b, 0x0e,
	0xe3, 0x45, 0xbd, 0x6e, 0xd7, 0xfa, 0x59, 0xc3, 0x60, 0x75, 0x97, 0x14, 0xf1, 0x0e, 0xba, 0xc8,
	0x06, 0x9c, 0x91, 0x59, 0x31, 0xfa, 0xda, 0xff, 0x0d, 0xd7, 0x9e, 0xee, 0xb3, 0x52, 0xf1, 0xeb,
	0x30, 0x77, 0x2c, 0x90, 0x51, 0xff, 0x65, 0xd4, 0x67, 0x4b, 0x8c, 0x44, 0xfe, 0x36, 0x00, 0x46,
	0x47, 0xe4, 0xbe, 0xe0, 0xa1, 0x55, 0x49, 0x74, 0xcf, 0x97, 0x74, 0x15, 0xc3, 0x0a, 0x2b, 0x3f,
	0x8c, 0x8d, 0xf2, 0xa9, 0x84, 0xb0, 0x8e, 0x78, 0x72, 0x09, 0xa6, 0x07, 0x6c, 0x23, 0xf6, 0xdb,
	0x88, 0x4d, 0xf5, 0x41, 0x89, 0xcc, 0x0a, 0xd4, 0x0d, 0x50, 0x71, 0xeb, 0xff, 0x51, 0x44, 0x6a,
	0x09, 0x7c, 0x87, 0x93, 0x45, 0x98, 0xca, 0x98, 0x46, 0xe0, 0x8f, 0x11, 0x68, 0xa4, 0x00, 0x1d,
	0x7e, 0xe1, 0x47, 0x1d, 0xea, 0x2e, 0x95, 0x31, 0x8f, 0x24, 0x25, 0x57, 0x61, 0x32, 0x19, 0x96,
	0xb4, 0x83, 0x4d, 0xbb, 0x38, 0x89, 0x66, 0x90, 0xd6, 0xf4, 0x5f, 0xd7, 0x00, 0xc9, 0x2e, 0xcc,
	0xe8, 0x31, 0xf1, 0x72, 0x73, 0x82, 0x2d, 0xa8, 0x20, 0xd9, 0x2e, 0x91, 0xcb, 0xd3, 0xb4, 0x85,
	0xf6, 0xe6, 0xc0, 0x76, 0xa7, 0xc3, 0xe2, 0x0f, 0xf8, 0xf0, 0x5a, 0x3a, 0x9e, 0x58, 0x5c, 0x1d,
	0xf1, 0xc2, 0xb1, 0x88, 0x66, 0x78, 0xb7, 0xcc, 0x7f, 0x37, 0x83, 0x93, 0x87, 0x50, 0x11, 0xfc,
	0x10, 0xab, 0xa5, 0x59, 0x2b, 0xf6, 0xe8, 0xeb, 0x64, 0x67, 0x95, 0xb0, 0x5d, 0x7e, 0xe8, 0xea,
	0x20, 0xcd, 0x6f, 0x55, 0xa8, 0xa0, 0x41, 0xe6, 0xa0, 0x8a, 0xa6, 0x9e, 0xb0, 0x57, 0x0e, 0x16,
	0x67, 0xd2, 0x9d, 0x44, 0x13, 0x67, 0xe7, 0x06, 0xd4, 0x7a, 0x11, 0x53, 0xda, 0xf1, 0xda, 0x19,
	0x3e, 0x7b, 0x55, 0x0d, 0x46, 0xda, 0xcd, 0x5c, 0x57, 0xdf, 0x38, 0x63, 0xb5, 0xf5, 0x0e, 0x34,
	0x62, 0xc1, 0xb8, 0x60, 0xea, 0xc8, 0x8b, 0xb8, 0xf5, 0x76, 0x04, 0x55, 0xc8, 0x08, 0x0e, 0x27,
	0x4f, 0x60, 0xbe, 0xb8, 0x69, 0x01, 0x95, 0x6d, 0xc1, 0xe2, 0xa4, 0x75, 0xef, 0x9c, 0x62, 0xe3,
	0xd3, 0x60, 0x52, 0x09, 0x84, 0x9b, 0x58, 0xe7, 0xf2, 0x9b, 0x78, 0x7f, 0xc0, 0xc5, 0xbc, 0xf2,
	0xfb, 0xf0, 0xde, 0x19, 0x73, 0x21, 0x9e, 0x41, 0x6b, 0xb0, 0x84, 0x4c, 0x7a, 0xfe, 0x9e, 0xe4,
	0xdd, 0x9e, 0xa2, 0x5e, 0xc2, 0xb4, 0x3e, 0x8c, 0xf0, 0x4c, 0xab, 0x1f, 0x60, 0x53, 0xae, 0xa6,
	0xf4, 0xc4, 0x83, 0xbb, 0xda, 0x48, 0xba, 0x24, 0x8f, 0xc2, 0x3d, 0xde, 0xb5, 0x3e, 0x0e, 0x7f,
	0x26, 0x68, 0xfc, 0x76, 0x02, 0x27, 0xeb, 0x27, 0xdc, 0xb8, 0x4f, 0xce, 0xf8, 0x47, 0x6e, 0x17,
	0x9a, 0xa5, 0x3b, 0x93, 0xaf, 0xfd, 0xe7, 0xe1, 0x49, 0x59, 0x85, 0x3b, 0x94, 0x2f, 0xfe, 0x1a,
	0x0c, 0xae, 0x5a, 0x5a, 0xb1, 0x2f, 0x27, 0x67, 0x18, 0xd0, 0x36, 0x0b, 0xfd, 0x6e, 0x96, 0x61,
	0x46, 0x32, 0x75, 0x7a, 0x70, 0xd2, 0x49, 0xfd, 0xea, 0x8c, 0x7f, 0x53, 0xef, 0x3d, 0x87, 0x16,
	0xe3, 0xa5, 0xd5, 0x1b, 0x7c, 0x06, 0x9f, 0xde, 0xed, 0x70, 0x19, 0xbc, 0xc8, 0xfc, 0xc1, 0xb8,
	0x5f, 0xca, 0xbd, 0x6a, 0xf2, 0x31, 0xba, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x58, 0xde,
	0x9e, 0x69, 0x07, 0x00, 0x00,
}