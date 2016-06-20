// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetChangeOStateTrigger_Ad.proto
// DO NOT EDIT!

/*
Package om_GetChangeOStateTrigger_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetChangeOStateTrigger_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetChangeOStateTrigger_Ad

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
	OrderStateId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_state_id,json=orderStateId" json:"order_state_id,omitempty"`
	OrderStateIdNull bool                        `protobuf:"varint,1001,opt,name=order_state_id_null,json=orderStateIdNull" json:"order_state_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderStateId
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
	RowId                        int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	IgnorePosAlreadyInOrderState *dstore_values.BooleanValue `protobuf:"bytes,10001,opt,name=ignore_pos_already_in_order_state,json=ignorePosAlreadyInOrderState" json:"ignore_pos_already_in_order_state,omitempty"`
	SuitableForOrderStateIdList  *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=suitable_for_order_state_id_list,json=suitableForOrderStateIdList" json:"suitable_for_order_state_id_list,omitempty"`
	ChangeOrderStateTriggerId    *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=change_order_state_trigger_id,json=changeOrderStateTriggerId" json:"change_order_state_trigger_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetIgnorePosAlreadyInOrderState() *dstore_values.BooleanValue {
	if m != nil {
		return m.IgnorePosAlreadyInOrderState
	}
	return nil
}

func (m *Response_Row) GetSuitableForOrderStateIdList() *dstore_values.StringValue {
	if m != nil {
		return m.SuitableForOrderStateIdList
	}
	return nil
}

func (m *Response_Row) GetChangeOrderStateTriggerId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ChangeOrderStateTriggerId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetChangeOStateTrigger_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetChangeOStateTrigger_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetChangeOStateTrigger_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0x65, 0xb7, 0xdb, 0xdd, 0x25, 0x8a, 0x2e, 0x59, 0x90, 0xd9, 0x56, 0xa5, 0xae, 0x2f, 0xe2,
	0x43, 0x2a, 0x8a, 0xb0, 0xe0, 0x53, 0x15, 0x95, 0x8a, 0x3b, 0x5d, 0xe2, 0x22, 0xe8, 0x83, 0x21,
	0xdd, 0xb9, 0x8e, 0x81, 0x69, 0x52, 0x92, 0xd4, 0xe2, 0x93, 0x7f, 0xc1, 0x8f, 0x3f, 0xe5, 0x4f,
	0xd1, 0x7f, 0xe1, 0x9d, 0x49, 0x6a, 0x67, 0x46, 0x51, 0xd9, 0x97, 0x96, 0x3b, 0xf7, 0x9e, 0x73,
	0x73, 0xce, 0x49, 0xc8, 0x83, 0xcc, 0x79, 0x63, 0x61, 0x08, 0x3a, 0x57, 0x1a, 0x86, 0x73, 0x6b,
	0xce, 0x20, 0x5b, 0x58, 0x70, 0x43, 0x33, 0x13, 0x4f, 0xc1, 0x3f, 0x7a, 0x27, 0x75, 0x0e, 0x93,
	0x17, 0x5e, 0x7a, 0x38, 0xb5, 0x2a, 0xcf, 0xc1, 0x8a, 0x51, 0xc6, 0x70, 0xce, 0x1b, 0x7a, 0x3b,
	0x80, 0x59, 0x00, 0xb3, 0xbf, 0x21, 0x7a, 0xfb, 0x71, 0xd1, 0x7b, 0x59, 0x2c, 0xc0, 0x05, 0x82,
	0xde, 0x41, 0x73, 0x3b, 0x58, 0x6b, 0x6c, 0x6c, 0xf5, 0x9b, 0xad, 0x19, 0x38, 0x27, 0x73, 0x88,
	0xcd, 0x9b, 0xed, 0xa6, 0x97, 0x4a, 0xbf, 0x35, 0x76, 0x26, 0xbd, 0x32, 0x3a, 0x0c, 0x1d, 0x7e,
	0x24, 0xe4, 0x44, 0x5a, 0x89, 0x4d, 0xb0, 0x8e, 0x8e, 0xc8, 0x25, 0x63, 0x33, 0x3c, 0x8b, 0x2b,
	0x0f, 0x26, 0x54, 0x96, 0x6c, 0x0c, 0x36, 0x6e, 0x5d, 0xb8, 0xdb, 0x67, 0x51, 0x44, 0x3c, 0x98,
	0xd2, 0x1e, 0xf0, 0xc8, 0x2f, 0xcb, 0x8a, 0x5f, 0xac, 0x20, 0x95, 0x94, 0x71, 0x46, 0x19, 0xd9,
	0x6f, 0x52, 0x08, 0xbd, 0x28, 0x8a, 0xe4, 0xc7, 0x0e, 0x12, 0xed, 0xf2, 0xbd, 0xfa, 0x6c, 0x8a,
	0x8d, 0xc3, 0xef, 0x5b, 0x64, 0x97, 0x83, 0x9b, 0x1b, 0xed, 0x80, 0xde, 0x21, 0xdd, 0x4a, 0x5e,
	0x5c, 0xdb, 0x63, 0x4d, 0xef, 0x82, 0xf4, 0xc7, 0xe5, 0x2f, 0x0f, 0x83, 0xf4, 0x15, 0xd9, 0x2b,
	0x85, 0x89, 0x9a, 0xb2, 0x64, 0x73, 0xd0, 0x41, 0x30, 0x6b, 0x81, 0xdb, 0xfa, 0x8f, 0xb1, 0x1e,
	0xaf, 0x6b, 0x7e, 0x79, 0xd6, 0xfc, 0x40, 0x8f, 0xc8, 0x4e, 0x34, 0x34, 0xe9, 0x54, 0x8c, 0xd7,
	0x7f, 0x63, 0x0c, 0x76, 0x1f, 0x87, 0x7f, 0xbe, 0x1a, 0xa7, 0xcf, 0x48, 0xc7, 0x9a, 0x65, 0xb2,
	0x55, 0xa1, 0x8e, 0xd8, 0xff, 0x5f, 0x00, 0xb6, 0x72, 0x82, 0x71, 0xb3, 0xe4, 0x25, 0x49, 0xef,
	0xdb, 0x26, 0xe9, 0x60, 0x41, 0xaf, 0x90, 0x6d, 0x2c, 0xcb, 0x48, 0x3e, 0xa5, 0x68, 0x4e, 0x97,
	0x77, 0xb1, 0x44, 0xbf, 0x81, 0xdc, 0x50, 0xb9, 0x46, 0x7e, 0x31, 0x37, 0x4e, 0xc8, 0xc2, 0x82,
	0xcc, 0x3e, 0xa0, 0x1d, 0xa2, 0x96, 0x42, 0xf2, 0x39, 0xfd, 0x63, 0x8c, 0x53, 0x63, 0x0a, 0x90,
	0x3a, 0xc4, 0x78, 0x35, 0xd0, 0x9c, 0x18, 0x37, 0x0a, 0x24, 0x63, 0x3d, 0xf9, 0x15, 0x16, 0x9d,
	0x92, 0x81, 0x5b, 0x28, 0x2f, 0xa7, 0x05, 0x08, 0xb4, 0x48, 0xb4, 0x32, 0x2e, 0x94, 0xf3, 0xc9,
	0x97, 0xb4, 0x99, 0x5a, 0xdc, 0xe2, 0xbc, 0x55, 0x3a, 0x0f, 0x4b, 0xfa, 0x2b, 0x92, 0x27, 0xc6,
	0x4e, 0x6a, 0x57, 0xe1, 0x39, 0xe2, 0xe9, 0x1b, 0x72, 0xed, 0xac, 0xb2, 0xa5, 0xc1, 0xee, 0xa3,
	0x3b, 0xa8, 0xfc, 0x6b, 0xfa, 0xef, 0xdb, 0x78, 0x10, 0x28, 0xd6, 0xe4, 0xd1, 0xdd, 0x71, 0xf6,
	0xf0, 0x94, 0xf4, 0x95, 0x69, 0xa5, 0xb1, 0x7e, 0xcb, 0xaf, 0xef, 0x9f, 0xeb, 0x95, 0x4f, 0xb7,
	0xab, 0x87, 0x74, 0xef, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x33, 0x4d, 0x3b, 0x25, 0x04,
	0x00, 0x00,
}
