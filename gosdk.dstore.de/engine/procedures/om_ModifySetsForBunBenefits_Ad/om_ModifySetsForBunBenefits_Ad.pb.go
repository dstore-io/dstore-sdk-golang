// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifySetsForBunBenefits_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifySetsForBunBenefits_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifySetsForBunBenefits_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifySetsForBunBenefits_Ad

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
	BenefitId             *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	BenefitIdNull         bool                        `protobuf:"varint,1001,opt,name=benefit_id_null,json=benefitIdNull" json:"benefit_id_null,omitempty"`
	ItemSetId             *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	ItemSetIdNull         bool                        `protobuf:"varint,1002,opt,name=item_set_id_null,json=itemSetIdNull" json:"item_set_id_null,omitempty"`
	SortNo                *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull            bool                        `protobuf:"varint,1003,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	DeleteCombination     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_combination,json=deleteCombination" json:"delete_combination,omitempty"`
	DeleteCombinationNull bool                        `protobuf:"varint,1004,opt,name=delete_combination_null,json=deleteCombinationNull" json:"delete_combination_null,omitempty"`
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

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetDeleteCombination() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCombination
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifySetsForBunBenefits_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifySetsForBunBenefits_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifySetsForBunBenefits_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifySetsForBunBenefits_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x5d, 0x9b, 0xd4, 0x53, 0xa5, 0x75, 0x44, 0x8d, 0x09, 0x8a, 0xd6, 0x0b, 0x7b, 0x21,
	0x1b, 0x51, 0x41, 0xa9, 0x08, 0x1a, 0x51, 0x88, 0x98, 0x20, 0x23, 0x08, 0x7a, 0x33, 0x6c, 0x3a,
	0x27, 0x61, 0x70, 0x77, 0x4e, 0x99, 0x99, 0x58, 0x7c, 0x0b, 0x9f, 0xc6, 0x4b, 0xdf, 0xc7, 0x9f,
	0x87, 0x70, 0x66, 0x67, 0xd2, 0x34, 0x89, 0xa8, 0xbd, 0x49, 0x72, 0xf6, 0x7c, 0x3f, 0x27, 0x67,
	0xbe, 0x59, 0x78, 0x22, 0xad, 0x23, 0x83, 0x5d, 0xd4, 0x13, 0xa5, 0xb1, 0x7b, 0x68, 0xe8, 0x00,
	0xe5, 0xd4, 0xa0, 0xed, 0x52, 0x25, 0x06, 0x24, 0xd5, 0xf8, 0xf3, 0x5b, 0x74, 0xf6, 0x25, 0x99,
	0xde, 0x54, 0xf7, 0x50, 0xe3, 0x58, 0x39, 0x2b, 0x9e, 0xc9, 0xdc, 0x23, 0x1d, 0xb1, 0x3b, 0x91,
	0x9e, 0x47, 0x7a, 0xfe, 0x77, 0x4e, 0xfb, 0x62, 0x32, 0xfb, 0x54, 0x94, 0x53, 0xb4, 0x51, 0xa2,
	0x7d, 0x75, 0x71, 0x02, 0x34, 0x86, 0x4c, 0x6a, 0x75, 0x16, 0x5b, 0x15, 0x5a, 0x5b, 0x4c, 0x30,
	0x35, 0x6f, 0x2d, 0x37, 0x5d, 0xa1, 0xf4, 0x98, 0x4c, 0x55, 0x38, 0x45, 0x3a, 0x82, 0x76, 0xbf,
	0x66, 0x00, 0x6f, 0x0a, 0x53, 0xf8, 0x2e, 0x1a, 0xcb, 0xf6, 0x01, 0x46, 0x71, 0x1e, 0xa1, 0x64,
	0x6b, 0xed, 0xc6, 0xda, 0xde, 0xd6, 0xbd, 0x4e, 0x9e, 0xfe, 0x43, 0x9a, 0x4a, 0x69, 0x87, 0x13,
	0x34, 0xef, 0x42, 0xc5, 0xcf, 0x26, 0x78, 0x5f, 0xb2, 0xdb, 0xb0, 0x3d, 0xe7, 0x0a, 0x3d, 0x2d,
	0xcb, 0xd6, 0xf7, 0xa6, 0x57, 0xd8, 0xe4, 0xe7, 0x8f, 0x41, 0x43, 0xff, 0x94, 0x3d, 0x86, 0x2d,
	0xe5, 0xb0, 0x12, 0x16, 0x6b, 0x97, 0xf5, 0xff, 0x70, 0x09, 0x78, 0xbf, 0x32, 0xef, 0xb2, 0x07,
	0x3b, 0x27, 0xc8, 0xd1, 0xe6, 0x47, 0xb2, 0x39, 0x46, 0xd5, 0x36, 0x0f, 0xa0, 0x69, 0xc9, 0x38,
	0xa1, 0xa9, 0x95, 0xfd, 0xdb, 0xa2, 0x11, 0xb0, 0x43, 0x62, 0x37, 0xe1, 0x5c, 0x62, 0x45, 0xed,
	0x9f, 0x51, 0x1b, 0x62, 0xbb, 0x16, 0x7e, 0x05, 0x4c, 0x62, 0xe9, 0x17, 0x26, 0x0e, 0xa8, 0x1a,
	0x29, 0x5d, 0xef, 0xb3, 0x75, 0xe6, 0x8f, 0x1e, 0x23, 0xa2, 0x12, 0x0b, 0x1d, 0x3d, 0x2e, 0x44,
	0xda, 0xf3, 0x39, 0x8b, 0x3d, 0x84, 0x2b, 0xab, 0x5a, 0xd1, 0xf9, 0x57, 0x74, 0xbe, 0xb4, 0x42,
	0x0a, 0x43, 0xec, 0x7e, 0x5b, 0x87, 0x4d, 0x8e, 0xf6, 0x90, 0xb4, 0x45, 0x76, 0x17, 0x36, 0xea,
	0x58, 0xa4, 0x13, 0x6b, 0xe7, 0x8b, 0xa9, 0x8b, 0x91, 0x79, 0x11, 0x3e, 0x79, 0x04, 0xb2, 0xf7,
	0xb0, 0x13, 0x02, 0x21, 0x4e, 0x24, 0xc2, 0x1f, 0x44, 0xe6, 0xc9, 0xf9, 0x12, 0x79, 0x39, 0x37,
	0x03, 0x5f, 0xf7, 0xe7, 0x35, 0xdf, 0xae, 0x16, 0x1f, 0xb0, 0x47, 0xd0, 0x4c, 0x41, 0xf4, 0x7b,
	0x0f, 0x8a, 0xd7, 0x57, 0x14, 0x63, 0x4c, 0x07, 0xf1, 0x9b, 0xcf, 0xe0, 0xec, 0x35, 0x64, 0x86,
	0x8e, 0xfc, 0x26, 0x03, 0x6b, 0x3f, 0x3f, 0xcd, 0xd5, 0xc9, 0x67, 0xbb, 0xc8, 0x39, 0x1d, 0xf1,
	0x20, 0xd3, 0xbe, 0x06, 0x99, 0xff, 0xcd, 0x2e, 0x43, 0xc3, 0x57, 0x21, 0x68, 0x5f, 0x86, 0x7e,
	0x3b, 0x1b, 0x7c, 0xc3, 0x97, 0x7d, 0xd9, 0x13, 0xd0, 0x51, 0xb4, 0xe4, 0x31, 0xbf, 0xdd, 0x1f,
	0x9e, 0x4e, 0xc8, 0xca, 0x8f, 0xb3, 0xbe, 0x3c, 0xfd, 0x0b, 0x60, 0xd4, 0xa8, 0x6f, 0xd8, 0xfd,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x9e, 0xe5, 0x86, 0x42, 0x04, 0x00, 0x00,
}