// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetSetsForBonItBenefits_Ad.proto
// DO NOT EDIT!

/*
Package om_GetSetsForBonItBenefits_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetSetsForBonItBenefits_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetSetsForBonItBenefits_Ad

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
	BenefitId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	BenefitIdNull bool                        `protobuf:"varint,1001,opt,name=benefit_id_null,json=benefitIdNull" json:"benefit_id_null,omitempty"`
	ItemSetId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	ItemSetIdNull bool                        `protobuf:"varint,1002,opt,name=item_set_id_null,json=itemSetIdNull" json:"item_set_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Parameters) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
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
	ItemConditionId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	BenefitId                *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	ItemSetId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	MaxQuantity              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=max_quantity,json=maxQuantity" json:"max_quantity,omitempty"`
	ItemConditionDescription *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=item_condition_description,json=itemConditionDescription" json:"item_condition_description,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Response_Row) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Response_Row) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
	}
	return nil
}

func (m *Response_Row) GetMaxQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxQuantity
	}
	return nil
}

func (m *Response_Row) GetItemConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionDescription
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetSetsForBonItBenefits_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetSetsForBonItBenefits_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetSetsForBonItBenefits_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetSetsForBonItBenefits_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xed, 0x8a, 0x13, 0x31,
	0x14, 0xa5, 0xd6, 0xb6, 0xeb, 0xad, 0xd2, 0x35, 0x82, 0x8c, 0x53, 0x10, 0x59, 0x7f, 0xb8, 0x20,
	0x4c, 0x45, 0x11, 0xd4, 0x5d, 0x58, 0xac, 0x9f, 0x45, 0x76, 0xd0, 0x11, 0x04, 0xfd, 0xe1, 0x30,
	0xdb, 0xdc, 0x1d, 0x82, 0x9d, 0xa4, 0x26, 0xa9, 0xab, 0x6f, 0xe1, 0xf7, 0x9b, 0xf8, 0x2a, 0xe2,
	0x2b, 0xe8, 0x53, 0x98, 0x4c, 0xd2, 0xed, 0xce, 0x08, 0x6e, 0xfb, 0x67, 0x86, 0x9b, 0x7b, 0xce,
	0xbd, 0x37, 0xe7, 0x1e, 0x02, 0xdb, 0x54, 0x69, 0x21, 0x71, 0x80, 0x3c, 0x67, 0x1c, 0x07, 0x53,
	0x29, 0xc6, 0x48, 0x67, 0x12, 0xd5, 0x40, 0x14, 0xe9, 0x23, 0xd4, 0xcf, 0x51, 0xab, 0x87, 0x42,
	0x0e, 0x05, 0x1f, 0xe9, 0x21, 0x72, 0xdc, 0x67, 0x5a, 0xa5, 0x77, 0x69, 0x64, 0x80, 0x5a, 0x90,
	0xab, 0x8e, 0x1d, 0x39, 0x76, 0xf4, 0x5f, 0x4a, 0x78, 0xce, 0xb7, 0x7a, 0x97, 0x4d, 0x66, 0xa8,
	0x5c, 0x85, 0xf0, 0x42, 0xb5, 0x3f, 0x4a, 0x29, 0xa4, 0x4f, 0xf5, 0xab, 0xa9, 0x02, 0x95, 0xca,
	0x72, 0xf4, 0xc9, 0xcb, 0xf5, 0xa4, 0xce, 0x18, 0xdf, 0x17, 0xb2, 0xc8, 0x34, 0x13, 0xdc, 0x81,
	0x36, 0x7e, 0x35, 0x00, 0x9e, 0x66, 0x32, 0x33, 0x59, 0x94, 0x8a, 0xdc, 0x01, 0xd8, 0x73, 0xf3,
	0xa4, 0x8c, 0x06, 0x8d, 0x4b, 0x8d, 0xcd, 0xee, 0xf5, 0x7e, 0xe4, 0xaf, 0xe0, 0xa7, 0x62, 0x5c,
	0x63, 0x8e, 0xf2, 0x85, 0x8d, 0x92, 0x53, 0x1e, 0x3e, 0xa2, 0xe4, 0x0a, 0xf4, 0x16, 0xdc, 0x94,
	0xcf, 0x26, 0x93, 0xe0, 0x77, 0xc7, 0x54, 0x58, 0x4b, 0xce, 0x1c, 0x82, 0x62, 0x73, 0x4a, 0xb6,
	0xa0, 0xcb, 0x34, 0x16, 0xa9, 0xc2, 0xb2, 0xcb, 0x89, 0x25, 0xba, 0x58, 0xbc, 0x11, 0xcd, 0x74,
	0xd9, 0x84, 0xf5, 0x23, 0x64, 0xd7, 0xe6, 0x8f, 0x6f, 0x73, 0x88, 0xb2, 0x6d, 0x36, 0x7e, 0xb6,
	0x60, 0x2d, 0x41, 0x35, 0x15, 0x5c, 0x21, 0xb9, 0x06, 0xad, 0x52, 0x38, 0x7f, 0xa7, 0x30, 0xaa,
	0xae, 0xc5, 0x89, 0xfa, 0xc0, 0x7e, 0x13, 0x07, 0x24, 0x2f, 0x61, 0xdd, 0x4a, 0x96, 0x1e, 0xd1,
	0xcc, 0x8c, 0xda, 0x34, 0xe4, 0xa8, 0x46, 0xae, 0x2b, 0xbb, 0x6b, 0xe2, 0xd1, 0x22, 0x4e, 0x7a,
	0x45, 0xf5, 0x80, 0xdc, 0x82, 0x8e, 0x5f, 0x55, 0xd0, 0x2c, 0x2b, 0x5e, 0xfc, 0xa7, 0xa2, 0x5b,
	0xe4, 0xae, 0xfb, 0x27, 0x73, 0x38, 0x79, 0x02, 0x4d, 0x29, 0x0e, 0x82, 0x93, 0x25, 0xeb, 0x76,
	0xb4, 0x82, 0xb7, 0xa2, 0xb9, 0x14, 0x51, 0x22, 0x0e, 0x12, 0x5b, 0x25, 0xfc, 0xd1, 0x84, 0xa6,
	0x09, 0xc8, 0x79, 0x68, 0x9b, 0xd0, 0xae, 0xe2, 0x63, 0x6c, 0xd4, 0x69, 0x25, 0x2d, 0x13, 0x1a,
	0xa9, 0x1f, 0xc3, 0xd9, 0x52, 0xea, 0xb1, 0xe0, 0x94, 0xd9, 0xc1, 0x2d, 0xe4, 0x53, 0x7c, 0xfc,
	0xba, 0x7a, 0x96, 0x76, 0x6f, 0xce, 0x32, 0x95, 0xb6, 0x2a, 0xb6, 0xfa, 0x1c, 0xaf, 0xe4, 0xab,
	0xed, 0xaa, 0x5d, 0xbe, 0xc4, 0x2b, 0xf9, 0x65, 0x07, 0x4e, 0x17, 0xd9, 0xfb, 0xf4, 0xed, 0x2c,
	0xe3, 0x9a, 0xe9, 0x0f, 0xc1, 0xd7, 0x25, 0xe8, 0x5d, 0xc3, 0x78, 0xe6, 0x09, 0xc6, 0x07, 0x61,
	0x4d, 0x05, 0x8a, 0x6a, 0x2c, 0xd9, 0xb4, 0x74, 0xc4, 0xb7, 0xb8, 0xea, 0x27, 0x5f, 0x4e, 0x69,
	0xc9, 0x78, 0xee, 0xaa, 0x05, 0x15, 0x35, 0xee, 0x2f, 0xc8, 0xe4, 0x26, 0x74, 0x94, 0x90, 0x3a,
	0xe5, 0x22, 0xf8, 0xbe, 0xc4, 0x58, 0x6d, 0x0b, 0x8e, 0xc5, 0xf0, 0x35, 0xf4, 0x99, 0xa8, 0xed,
	0x7e, 0xf1, 0x2a, 0xbd, 0xda, 0xc9, 0x85, 0xa2, 0x6f, 0xe6, 0x79, 0xba, 0xf2, 0xc3, 0xb5, 0xd7,
	0x2e, 0x9f, 0x86, 0x1b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x15, 0x2d, 0x83, 0x97, 0xf9, 0x04,
	0x00, 0x00,
}
