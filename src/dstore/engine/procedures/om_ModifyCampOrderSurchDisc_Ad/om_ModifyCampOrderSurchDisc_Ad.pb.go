// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampOrderSurchDisc_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampOrderSurchDisc_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampOrderSurchDisc_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampOrderSurchDisc_Ad

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
	ApplyToSurchargeTypeId      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=apply_to_surcharge_type_id,json=applyToSurchargeTypeId" json:"apply_to_surcharge_type_id,omitempty"`
	ApplyToSurchargeTypeIdNull  bool                        `protobuf:"varint,1001,opt,name=apply_to_surcharge_type_id_null,json=applyToSurchargeTypeIdNull" json:"apply_to_surcharge_type_id_null,omitempty"`
	DiscountSurchargeTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=discount_surcharge_type_id,json=discountSurchargeTypeId" json:"discount_surcharge_type_id,omitempty"`
	DiscountSurchargeTypeIdNull bool                        `protobuf:"varint,1002,opt,name=discount_surcharge_type_id_null,json=discountSurchargeTypeIdNull" json:"discount_surcharge_type_id_null,omitempty"`
	DiscountValue               *dstore_values.DecimalValue `protobuf:"bytes,3,opt,name=discount_value,json=discountValue" json:"discount_value,omitempty"`
	DiscountValueNull           bool                        `protobuf:"varint,1003,opt,name=discount_value_null,json=discountValueNull" json:"discount_value_null,omitempty"`
	DeleteBenefit               *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=delete_benefit,json=deleteBenefit" json:"delete_benefit,omitempty"`
	DeleteBenefitNull           bool                        `protobuf:"varint,1004,opt,name=delete_benefit_null,json=deleteBenefitNull" json:"delete_benefit_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplyToSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplyToSurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetDiscountSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DiscountSurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetDiscountValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.DiscountValue
	}
	return nil
}

func (m *Parameters) GetDeleteBenefit() *dstore_values.IntegerValue {
	if m != nil {
		return m.DeleteBenefit
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	BenefitId       *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
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

func (m *Response) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampOrderSurchDisc_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampOrderSurchDisc_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampOrderSurchDisc_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x9b, 0xa6, 0x2d, 0x83, 0xca, 0x8f, 0x2b, 0x95, 0xb0, 0x11, 0x14, 0x95, 0x0b, 0x07,
	0xb4, 0x41, 0x70, 0x00, 0x55, 0xe2, 0x40, 0x68, 0x0f, 0x91, 0x48, 0x41, 0xdb, 0xaa, 0xfc, 0x5c,
	0x56, 0x6e, 0x3c, 0x09, 0x96, 0x76, 0xed, 0x95, 0xed, 0xa5, 0xca, 0x5b, 0xf0, 0x34, 0xbc, 0x13,
	0x3f, 0x4f, 0xc0, 0x09, 0x7b, 0xbd, 0xdb, 0x66, 0x97, 0x84, 0x88, 0x5e, 0x12, 0x4d, 0x66, 0xbe,
	0x9f, 0x8c, 0x3f, 0x1b, 0x5e, 0x32, 0x6d, 0xa4, 0xc2, 0x1e, 0x8a, 0x09, 0x17, 0xd8, 0xcb, 0x94,
	0x1c, 0x21, 0xcb, 0x15, 0xea, 0x9e, 0x4c, 0xe3, 0xa1, 0x64, 0x7c, 0x3c, 0x7d, 0x4d, 0xd3, 0xec,
	0xad, 0x62, 0xa8, 0x8e, 0x73, 0x35, 0xfa, 0x7c, 0xc0, 0xf5, 0x28, 0x7e, 0xc5, 0x42, 0x3b, 0x69,
	0x24, 0x79, 0xec, 0xe1, 0xa1, 0x87, 0x87, 0xff, 0xc6, 0x04, 0xdb, 0xa5, 0xd8, 0x17, 0x9a, 0xe4,
	0xa8, 0x3d, 0x45, 0x70, 0xb7, 0xee, 0x00, 0x95, 0x92, 0xaa, 0x6c, 0x75, 0xeb, 0xad, 0x14, 0xb5,
	0xa6, 0x13, 0x2c, 0x9b, 0x0f, 0x9b, 0x4d, 0x43, 0xb9, 0x18, 0x4b, 0x95, 0x52, 0xc3, 0xa5, 0xf0,
	0x43, 0x7b, 0xdf, 0xd6, 0x00, 0xde, 0x51, 0x45, 0x6d, 0x17, 0x95, 0x26, 0xef, 0x21, 0xa0, 0x59,
	0x96, 0x4c, 0x63, 0x23, 0x63, 0xed, 0x9c, 0x51, 0x35, 0xc1, 0xd8, 0x4c, 0x33, 0x8c, 0x39, 0xeb,
	0xac, 0x3c, 0x58, 0x79, 0x74, 0xfd, 0x69, 0x37, 0x2c, 0xff, 0x53, 0xe9, 0x92, 0x0b, 0x83, 0x13,
	0x54, 0xa7, 0xae, 0x8a, 0x76, 0x0a, 0xf8, 0x89, 0x3c, 0xae, 0xc0, 0x27, 0x16, 0x3b, 0x60, 0xe4,
	0x00, 0x76, 0x17, 0x13, 0xc7, 0x22, 0x4f, 0x92, 0xce, 0xf7, 0x0d, 0x4b, 0xbf, 0x19, 0x05, 0xf3,
	0x19, 0x8e, 0xec, 0x08, 0xf9, 0x00, 0x01, 0xb3, 0xab, 0x92, 0xb9, 0x30, 0x73, 0xec, 0xad, 0x2e,
	0xb7, 0x77, 0xa7, 0x82, 0x37, 0xfd, 0x1d, 0xc2, 0xee, 0x62, 0x66, 0xef, 0xef, 0x87, 0xf7, 0xd7,
	0x5d, 0x40, 0x51, 0x18, 0xec, 0xc3, 0x8d, 0x0b, 0x9a, 0x42, 0xbf, 0xd3, 0x9a, 0x6b, 0x8a, 0xe1,
	0x88, 0xa7, 0x34, 0xf1, 0xa6, 0xb6, 0x2a, 0x48, 0x51, 0x92, 0x1e, 0x6c, 0xd7, 0x39, 0xbc, 0xfc,
	0x4f, 0x2f, 0x7f, 0xbb, 0x36, 0x7c, 0x21, 0x8a, 0x89, 0x3d, 0xc0, 0xf8, 0x0c, 0x05, 0x8e, 0xb9,
	0xe9, 0xac, 0x2d, 0xdf, 0xc4, 0x96, 0x87, 0xf4, 0x3d, 0xa2, 0x10, 0xad, 0x71, 0x78, 0xd1, 0x5f,
	0x95, 0xe8, 0xec, 0xb0, 0x13, 0xdd, 0xfb, 0xbd, 0x0a, 0x9b, 0x11, 0xea, 0x4c, 0x0a, 0x8d, 0xe4,
	0x09, 0xb4, 0x8b, 0x58, 0x96, 0x09, 0x09, 0xc2, 0x7a, 0xea, 0x7d, 0x64, 0x0f, 0xdd, 0x67, 0xe4,
	0x07, 0xc9, 0x47, 0xb8, 0xe5, 0x02, 0x19, 0xcf, 0x24, 0xd2, 0x9e, 0x5f, 0xcb, 0x82, 0xc3, 0x06,
	0xb8, 0x99, 0xdb, 0xa1, 0xad, 0x07, 0x97, 0x75, 0x74, 0x33, 0xad, 0xff, 0x40, 0x5e, 0xc0, 0x46,
	0x79, 0x11, 0xec, 0xf2, 0x1d, 0xe3, 0xfd, 0xbf, 0x18, 0xfd, 0x35, 0x19, 0xfa, 0xef, 0xa8, 0x1a,
	0x27, 0x6f, 0xa0, 0xa5, 0xe4, 0xb9, 0xdd, 0x9e, 0x43, 0xed, 0x87, 0xff, 0x73, 0x75, 0xc3, 0x6a,
	0x17, 0x61, 0x24, 0xcf, 0x23, 0x47, 0x43, 0xf6, 0x01, 0xaa, 0x5d, 0xda, 0x70, 0xe2, 0xf2, 0x23,
	0xb9, 0x56, 0x8e, 0x0f, 0x58, 0x70, 0x0f, 0x5a, 0x96, 0x87, 0xec, 0xc0, 0xba, 0x65, 0x72, 0xf0,
	0xaf, 0x47, 0x16, 0xdf, 0x8e, 0xda, 0xb6, 0x1c, 0xb0, 0xfe, 0x29, 0x74, 0xb9, 0x6c, 0xf8, 0xbb,
	0x7c, 0x99, 0x3e, 0x3d, 0xbf, 0xe2, 0x9b, 0x75, 0xb6, 0x5e, 0x3c, 0x0a, 0xcf, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0xef, 0xbd, 0x14, 0xf5, 0x04, 0x00, 0x00,
}