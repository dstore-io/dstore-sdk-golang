// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyOrderStatesInCats_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyOrderStatesInCats_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyOrderStatesInCats_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyOrderStatesInCats_Ad

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
	OrderStateCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_state_category_id,json=orderStateCategoryId" json:"order_state_category_id,omitempty"`
	OrderStateCategoryIdNull bool                        `protobuf:"varint,1001,opt,name=order_state_category_id_null,json=orderStateCategoryIdNull" json:"order_state_category_id_null,omitempty"`
	OrderStateId             *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=order_state_id,json=orderStateId" json:"order_state_id,omitempty"`
	OrderStateIdNull         bool                        `protobuf:"varint,1002,opt,name=order_state_id_null,json=orderStateIdNull" json:"order_state_id_null,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull               bool                        `protobuf:"varint,1003,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	Delete                   *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull               bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderStateCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderStateCategoryId
	}
	return nil
}

func (m *Parameters) GetOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderStateId
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyOrderStatesInCats_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyOrderStatesInCats_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyOrderStatesInCats_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x6f, 0x6b, 0x13, 0x4f,
	0x10, 0xc7, 0x69, 0xef, 0x97, 0xa4, 0x4c, 0xcb, 0xcf, 0xb2, 0x15, 0x3d, 0x13, 0x95, 0x5a, 0x9f,
	0x08, 0xc2, 0x46, 0xac, 0x88, 0x82, 0x20, 0xb5, 0xf8, 0x20, 0x48, 0xa2, 0xac, 0x28, 0xe8, 0x93,
	0x63, 0xdb, 0xdd, 0x86, 0x83, 0xcb, 0x4e, 0xd9, 0xbd, 0x58, 0xfa, 0x2e, 0x7c, 0x43, 0x3e, 0xf3,
	0xcd, 0xf8, 0xe7, 0x45, 0x38, 0x9b, 0xd9, 0x9a, 0x5c, 0xac, 0x16, 0x7d, 0x72, 0xb9, 0xb9, 0x99,
	0xef, 0x67, 0xe6, 0x26, 0xdf, 0x3d, 0x78, 0x62, 0x42, 0x8d, 0xde, 0xf6, 0xad, 0x1b, 0x97, 0xce,
	0xf6, 0x8f, 0x3d, 0x1e, 0x5a, 0x33, 0xf5, 0x36, 0xf4, 0x71, 0x52, 0x0c, 0xd1, 0x94, 0x47, 0xa7,
	0x2f, 0xbd, 0xb1, 0xfe, 0x75, 0xad, 0x6b, 0x1b, 0x06, 0x6e, 0x5f, 0xd7, 0xa1, 0xd8, 0x33, 0x92,
	0x0a, 0x6b, 0x14, 0x77, 0x59, 0x2d, 0x59, 0x2d, 0xff, 0x28, 0xe9, 0x6e, 0xa5, 0x56, 0x1f, 0x74,
	0x35, 0xb5, 0x81, 0x09, 0xdd, 0x6b, 0xcd, 0xfe, 0xd6, 0x7b, 0xf4, 0x29, 0xd5, 0x6b, 0xa6, 0x26,
	0x36, 0x04, 0x3d, 0xb6, 0x29, 0x79, 0x7b, 0x39, 0x59, 0xeb, 0xd2, 0x1d, 0xa1, 0x9f, 0xe8, 0xba,
	0x44, 0xc7, 0x45, 0x3b, 0x9f, 0x33, 0x80, 0x57, 0xda, 0x6b, 0xca, 0x5a, 0x1f, 0x84, 0x82, 0xab,
	0x18, 0x07, 0x2b, 0x42, 0x9c, 0xac, 0x38, 0xa4, 0xcb, 0x18, 0xfd, 0x69, 0x51, 0x9a, 0x7c, 0x65,
	0x7b, 0xe5, 0xce, 0xfa, 0xfd, 0x9e, 0x4c, 0xef, 0x93, 0x46, 0x2c, 0x1d, 0x55, 0x58, 0xff, 0x36,
	0x46, 0xea, 0x32, 0xfe, 0x7c, 0xa9, 0xfd, 0xa4, 0x1c, 0x18, 0xf1, 0x14, 0xae, 0xff, 0x86, 0x59,
	0xb8, 0x69, 0x55, 0xe5, 0x5f, 0x3a, 0x44, 0x5e, 0x53, 0xf9, 0x79, 0xe2, 0x11, 0x15, 0x88, 0x3d,
	0xf8, 0x7f, 0x11, 0x40, 0xb3, 0xac, 0x5e, 0x3c, 0xcb, 0xc6, 0x1c, 0x47, 0x33, 0x48, 0xd8, 0x6a,
	0x22, 0xb8, 0xf5, 0x57, 0x6e, 0xbd, 0xb9, 0x58, 0x3b, 0x6b, 0xf9, 0x00, 0x3a, 0x01, 0x7d, 0x5d,
	0x38, 0xcc, 0xb3, 0x8b, 0x7b, 0xb5, 0x63, 0xed, 0x08, 0xc5, 0x2d, 0xd8, 0x48, 0x2a, 0xc6, 0x7f,
	0x63, 0x3c, 0x70, 0x7a, 0x06, 0xde, 0x85, 0xb6, 0xb1, 0x15, 0x2d, 0x3b, 0xff, 0xef, 0x5c, 0xee,
	0x01, 0x62, 0x65, 0xb5, 0x4b, 0x5c, 0x2e, 0x15, 0xdb, 0xb0, 0xce, 0x77, 0x8c, 0xfd, 0x9e, 0xb0,
	0xfc, 0x2c, 0x62, 0x77, 0x3e, 0xad, 0xc2, 0x9a, 0xb2, 0xe1, 0x18, 0x5d, 0xb0, 0xe2, 0x1e, 0xb4,
	0x66, 0x26, 0x49, 0x7f, 0x59, 0x57, 0x36, 0x2d, 0xc8, 0x06, 0x7a, 0x1e, 0xaf, 0x8a, 0x0b, 0xc5,
	0x3b, 0xd8, 0x8c, 0xf6, 0x28, 0x16, 0xfc, 0x41, 0x3b, 0xce, 0x48, 0x2c, 0x97, 0xc4, 0xcb, 0x2e,
	0x1a, 0x52, 0x3c, 0x98, 0xc7, 0xea, 0xd2, 0xa4, 0xf9, 0x40, 0x3c, 0x82, 0x4e, 0xb2, 0x25, 0x6d,
	0x32, 0x12, 0x6f, 0xfe, 0x42, 0x64, 0xd3, 0x0e, 0xf9, 0x57, 0x9d, 0x95, 0x8b, 0x17, 0x90, 0x79,
	0x3c, 0xa1, 0x3d, 0x45, 0xd5, 0x63, 0xf9, 0x17, 0xe7, 0x48, 0x9e, 0xad, 0x42, 0x2a, 0x3c, 0x51,
	0x91, 0xd2, 0xbd, 0x01, 0x19, 0xdd, 0x8b, 0x2b, 0xd0, 0xa6, 0x28, 0x5a, 0xe8, 0xe3, 0x88, 0x96,
	0xd3, 0x52, 0x2d, 0x0a, 0x07, 0xe6, 0xd9, 0x1b, 0xe8, 0x95, 0xb8, 0xd4, 0x62, 0x7e, 0xd0, 0xdf,
	0x3f, 0xfc, 0xb7, 0x4f, 0xc0, 0x41, 0x7b, 0x76, 0xc8, 0x76, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0xa4, 0xee, 0x33, 0x43, 0x04, 0x00, 0x00,
}
