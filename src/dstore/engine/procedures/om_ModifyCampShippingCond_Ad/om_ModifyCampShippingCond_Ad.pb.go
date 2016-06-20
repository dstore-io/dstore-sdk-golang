// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampShippingCond_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampShippingCond_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampShippingCond_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampShippingCond_Ad

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
	FilterByShippingTypeIds         *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=filter_by_shipping_type_ids,json=filterByShippingTypeIds" json:"filter_by_shipping_type_ids,omitempty"`
	FilterByShippingTypeIdsNull     bool                        `protobuf:"varint,1001,opt,name=filter_by_shipping_type_ids_null,json=filterByShippingTypeIdsNull" json:"filter_by_shipping_type_ids_null,omitempty"`
	NegateFilter                    *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=negate_filter,json=negateFilter" json:"negate_filter,omitempty"`
	NegateFilterNull                bool                        `protobuf:"varint,1002,opt,name=negate_filter_null,json=negateFilterNull" json:"negate_filter_null,omitempty"`
	DeleteShippingTypeCondition     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete_shipping_type_condition,json=deleteShippingTypeCondition" json:"delete_shipping_type_condition,omitempty"`
	DeleteShippingTypeConditionNull bool                        `protobuf:"varint,1003,opt,name=delete_shipping_type_condition_null,json=deleteShippingTypeConditionNull" json:"delete_shipping_type_condition_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetFilterByShippingTypeIds() *dstore_values.StringValue {
	if m != nil {
		return m.FilterByShippingTypeIds
	}
	return nil
}

func (m *Parameters) GetNegateFilter() *dstore_values.BooleanValue {
	if m != nil {
		return m.NegateFilter
	}
	return nil
}

func (m *Parameters) GetDeleteShippingTypeCondition() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteShippingTypeCondition
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ConditionId     *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
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

func (m *Response) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampShippingCond_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampShippingCond_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampShippingCond_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0xa5, 0x5d, 0xd3, 0x96, 0xdb, 0x8a, 0x65, 0x04, 0x5d, 0xb3, 0x58, 0x4b, 0xfb, 0x22, 0x82,
	0x1b, 0x51, 0x84, 0x82, 0x20, 0xda, 0x60, 0x21, 0x42, 0x8a, 0xac, 0x45, 0xd4, 0x97, 0x71, 0xd2,
	0xb9, 0x59, 0x07, 0x76, 0x67, 0x96, 0x99, 0x8d, 0x25, 0x7f, 0xe1, 0xa7, 0xf9, 0x1b, 0xea, 0x2f,
	0xf8, 0xe0, 0xec, 0xde, 0x8d, 0x49, 0x56, 0x8d, 0xd2, 0x97, 0x24, 0x93, 0x7b, 0xce, 0xb9, 0x67,
	0xee, 0x3d, 0xbb, 0xf0, 0x44, 0xba, 0xd2, 0x58, 0xec, 0xa1, 0x4e, 0x95, 0xc6, 0x5e, 0x61, 0xcd,
	0x39, 0xca, 0x89, 0x45, 0xd7, 0x33, 0x39, 0x1f, 0x1a, 0xa9, 0xc6, 0xd3, 0xbe, 0xc8, 0x8b, 0xd7,
	0x1f, 0x55, 0x51, 0x28, 0x9d, 0xf6, 0x8d, 0x96, 0xfc, 0xb9, 0x8c, 0x3d, 0xae, 0x34, 0xec, 0x1e,
	0x91, 0x63, 0x22, 0xc7, 0xab, 0x18, 0xdd, 0xeb, 0x4d, 0xa3, 0x4f, 0x22, 0x9b, 0xa0, 0x23, 0x81,
	0xee, 0xad, 0xe5, 0xee, 0x68, 0xad, 0xb1, 0x4d, 0x29, 0x5a, 0x2e, 0xe5, 0xe8, 0x9c, 0x48, 0xb1,
	0x29, 0x1e, 0xb6, 0x8b, 0xa5, 0x50, 0x7a, 0x6c, 0x6c, 0x2e, 0x4a, 0x65, 0x34, 0x81, 0x0e, 0xbe,
	0x04, 0x00, 0xaf, 0x84, 0x15, 0xbe, 0x8a, 0xd6, 0xb1, 0xb7, 0x10, 0x8d, 0x55, 0xe6, 0x7f, 0xf2,
	0xd1, 0x94, 0xbb, 0xc6, 0x1d, 0x2f, 0xa7, 0x05, 0x72, 0x25, 0x5d, 0xb8, 0xb6, 0xbf, 0x76, 0x77,
	0xfb, 0x61, 0x37, 0x6e, 0xae, 0xd4, 0xd8, 0x74, 0xa5, 0xf5, 0xa8, 0x37, 0xd5, 0x21, 0xb9, 0x49,
	0xf4, 0xe3, 0xe9, 0xec, 0x6a, 0x67, 0x9e, 0x3b, 0x90, 0x8e, 0x9d, 0xc0, 0xfe, 0x0a, 0x65, 0xae,
	0x27, 0x59, 0x16, 0x7e, 0xdd, 0xf4, 0xfa, 0x5b, 0x49, 0xf4, 0x17, 0x8d, 0x53, 0x8f, 0x61, 0xcf,
	0xe0, 0xaa, 0xc6, 0x54, 0x94, 0xc8, 0x09, 0x15, 0xae, 0xd7, 0x9e, 0xa2, 0x96, 0xa7, 0x91, 0x31,
	0x19, 0x0a, 0x4d, 0xa6, 0x76, 0x88, 0x71, 0x52, 0x13, 0xd8, 0x7d, 0x60, 0x4b, 0x0a, 0xd4, 0xfb,
	0x1b, 0xf5, 0xde, 0x5d, 0x84, 0xd6, 0x0d, 0x3f, 0xc0, 0x9e, 0xc4, 0xcc, 0x8f, 0xa7, 0xe5, 0xfa,
	0xdc, 0xef, 0x4c, 0x55, 0x93, 0x0c, 0x83, 0x7f, 0x3b, 0x88, 0x48, 0x62, 0xf1, 0x42, 0xfd, 0x19,
	0x9f, 0x0d, 0xe1, 0x70, 0x75, 0x07, 0x72, 0xf8, 0x9d, 0x1c, 0xde, 0x59, 0x21, 0x55, 0x19, 0x3e,
	0xf8, 0xb1, 0x0e, 0x5b, 0x09, 0xba, 0xc2, 0x68, 0x87, 0xec, 0x01, 0x74, 0xea, 0xc0, 0xb4, 0x57,
	0xd7, 0xa4, 0x91, 0xc2, 0xf4, 0xa2, 0xfa, 0x4c, 0x08, 0xc8, 0xde, 0xc1, 0x6e, 0x15, 0x15, 0xbe,
	0x90, 0x15, 0x3f, 0xe3, 0xc0, 0x93, 0xe3, 0x16, 0xb9, 0x9d, 0xa8, 0xa1, 0x3f, 0x0f, 0xe6, 0xe7,
	0xe4, 0x5a, 0xbe, 0xfc, 0x07, 0x3b, 0x82, 0xcd, 0x26, 0xa2, 0x7e, 0x66, 0x95, 0xe2, 0xde, 0x6f,
	0x8a, 0x14, 0xe0, 0x21, 0x7d, 0x27, 0x33, 0x38, 0x7b, 0x09, 0x81, 0x35, 0x17, 0xe1, 0x95, 0x9a,
	0x75, 0x14, 0xff, 0xff, 0x23, 0x15, 0xcf, 0x26, 0x11, 0x27, 0xe6, 0x22, 0xa9, 0x44, 0xd8, 0x53,
	0xd8, 0x99, 0x4f, 0x56, 0xc9, 0x10, 0xff, 0xb8, 0x3e, 0xa5, 0x4b, 0x4c, 0xd1, 0xd2, 0xfa, 0xb6,
	0x7f, 0x11, 0x06, 0xb2, 0x7b, 0x1b, 0x02, 0xaf, 0xc5, 0x6e, 0xc0, 0x86, 0x57, 0xab, 0x04, 0x3e,
	0x9f, 0x7a, 0x85, 0x4e, 0xd2, 0xf1, 0xc7, 0x81, 0x3c, 0x3e, 0x83, 0x48, 0x99, 0x96, 0xc3, 0xf9,
	0x1b, 0xe3, 0xfd, 0xe3, 0x4b, 0xbd, 0x4b, 0x46, 0x1b, 0xf5, 0xe3, 0xfa, 0xe8, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x3a, 0x73, 0x94, 0xd9, 0x8b, 0x04, 0x00, 0x00,
}