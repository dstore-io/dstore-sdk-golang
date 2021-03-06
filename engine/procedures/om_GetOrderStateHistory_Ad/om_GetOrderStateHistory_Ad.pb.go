// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderStateHistory_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderStateHistory_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderStateHistory_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderStateHistory_Ad

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
	OrderOrOrderContentId     *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=order_or_order_content_id,json=orderOrOrderContentId" json:"order_or_order_content_id,omitempty"`
	OrderOrOrderContentIdNull bool                          `protobuf:"varint,1001,opt,name=order_or_order_content_id_null,json=orderOrOrderContentIdNull" json:"order_or_order_content_id_null,omitempty"`
	IsOrderId                 *dstore_values.BooleanValue   `protobuf:"bytes,2,opt,name=is_order_id,json=isOrderId" json:"is_order_id,omitempty"`
	IsOrderIdNull             bool                          `protobuf:"varint,1002,opt,name=is_order_id_null,json=isOrderIdNull" json:"is_order_id_null,omitempty"`
	FromOrderStateId          *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=from_order_state_id,json=fromOrderStateId" json:"from_order_state_id,omitempty"`
	FromOrderStateIdNull      bool                          `protobuf:"varint,1003,opt,name=from_order_state_id_null,json=fromOrderStateIdNull" json:"from_order_state_id_null,omitempty"`
	ToOrderStateId            *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=to_order_state_id,json=toOrderStateId" json:"to_order_state_id,omitempty"`
	ToOrderStateIdNull        bool                          `protobuf:"varint,1004,opt,name=to_order_state_id_null,json=toOrderStateIdNull" json:"to_order_state_id_null,omitempty"`
	UserId                    *dstore_values.IntegerValue   `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull                bool                          `protobuf:"varint,1005,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	FromDateAndTime           *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=from_date_and_time,json=fromDateAndTime" json:"from_date_and_time,omitempty"`
	FromDateAndTimeNull       bool                          `protobuf:"varint,1006,opt,name=from_date_and_time_null,json=fromDateAndTimeNull" json:"from_date_and_time_null,omitempty"`
	ToDateAndTime             *dstore_values.TimestampValue `protobuf:"bytes,7,opt,name=to_date_and_time,json=toDateAndTime" json:"to_date_and_time,omitempty"`
	ToDateAndTimeNull         bool                          `protobuf:"varint,1007,opt,name=to_date_and_time_null,json=toDateAndTimeNull" json:"to_date_and_time_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderOrOrderContentId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderOrOrderContentId
	}
	return nil
}

func (m *Parameters) GetIsOrderId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsOrderId
	}
	return nil
}

func (m *Parameters) GetFromOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromOrderStateId
	}
	return nil
}

func (m *Parameters) GetToOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToOrderStateId
	}
	return nil
}

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Parameters) GetFromDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDateAndTime
	}
	return nil
}

func (m *Parameters) GetToDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDateAndTime
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
	RowId                   int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ChangingDateAndTime     *dstore_values.TimestampValue `protobuf:"bytes,10001,opt,name=changing_date_and_time,json=changingDateAndTime" json:"changing_date_and_time,omitempty"`
	ToOrderStateId          *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=to_order_state_id,json=toOrderStateId" json:"to_order_state_id,omitempty"`
	UserName                *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserId                  *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	CompleteOrder           *dstore_values.BooleanValue   `protobuf:"bytes,10005,opt,name=complete_order,json=completeOrder" json:"complete_order,omitempty"`
	FromOrderStateId        *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=from_order_state_id,json=fromOrderStateId" json:"from_order_state_id,omitempty"`
	ChangingDateAndTimeChar *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=changing_date_and_time_char,json=changingDateAndTimeChar" json:"changing_date_and_time_char,omitempty"`
	OrderId                 *dstore_values.IntegerValue   `protobuf:"bytes,10008,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	ToOrderState            *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=to_order_state,json=toOrderState" json:"to_order_state,omitempty"`
	OrderContentId          *dstore_values.IntegerValue   `protobuf:"bytes,10010,opt,name=order_content_id,json=orderContentId" json:"order_content_id,omitempty"`
	FromOrderState          *dstore_values.StringValue    `protobuf:"bytes,10011,opt,name=from_order_state,json=fromOrderState" json:"from_order_state,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetChangingDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.ChangingDateAndTime
	}
	return nil
}

func (m *Response_Row) GetToOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToOrderStateId
	}
	return nil
}

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetCompleteOrder() *dstore_values.BooleanValue {
	if m != nil {
		return m.CompleteOrder
	}
	return nil
}

func (m *Response_Row) GetFromOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromOrderStateId
	}
	return nil
}

func (m *Response_Row) GetChangingDateAndTimeChar() *dstore_values.StringValue {
	if m != nil {
		return m.ChangingDateAndTimeChar
	}
	return nil
}

func (m *Response_Row) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *Response_Row) GetToOrderState() *dstore_values.StringValue {
	if m != nil {
		return m.ToOrderState
	}
	return nil
}

func (m *Response_Row) GetOrderContentId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderContentId
	}
	return nil
}

func (m *Response_Row) GetFromOrderState() *dstore_values.StringValue {
	if m != nil {
		return m.FromOrderState
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderStateHistory_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderStateHistory_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderStateHistory_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetOrderStateHistory_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x52, 0xd3, 0x5c,
	0x14, 0x1e, 0xfe, 0xfe, 0x3d, 0xb0, 0x80, 0x52, 0x52, 0x81, 0xd0, 0x8e, 0x8c, 0xe2, 0x0d, 0x57,
	0x41, 0xc5, 0x03, 0x8e, 0xde, 0x00, 0x22, 0xa0, 0x43, 0x70, 0x82, 0x3a, 0x83, 0x5e, 0x64, 0x42,
	0xb3, 0xa9, 0x19, 0xdb, 0xec, 0xce, 0xce, 0xae, 0x8c, 0x6f, 0xe1, 0xf9, 0x78, 0xeb, 0x93, 0xf8,
	0x32, 0x8e, 0xe7, 0x57, 0x70, 0xed, 0xac, 0x14, 0x9a, 0xb4, 0x4c, 0x8a, 0x33, 0x9d, 0xb6, 0x3b,
	0x6b, 0x7d, 0xeb, 0xfb, 0xb2, 0x4e, 0x09, 0x5c, 0x73, 0x03, 0xc9, 0x05, 0x5b, 0x60, 0x7e, 0xdd,
	0xf3, 0xd9, 0x42, 0x4b, 0xf0, 0x1a, 0x73, 0xdb, 0x82, 0x05, 0x0b, 0xbc, 0x69, 0xaf, 0x33, 0xb9,
	0x2d, 0x5c, 0x26, 0x76, 0xa4, 0x23, 0xd9, 0x86, 0xa7, 0x3c, 0x9f, 0xd9, 0xcb, 0xae, 0x81, 0x5e,
	0x92, 0x6b, 0xf3, 0x04, 0x35, 0x08, 0x6a, 0x1c, 0xef, 0x5f, 0x29, 0x47, 0x24, 0x4f, 0x9d, 0x46,
	0x9b, 0x05, 0x04, 0xaf, 0xcc, 0xc4, 0x99, 0x99, 0x10, 0x5c, 0x44, 0xa6, 0x6a, 0xdc, 0xd4, 0x64,
	0x41, 0xe0, 0xd4, 0x59, 0x64, 0x3c, 0x97, 0x34, 0x4a, 0xc7, 0xf3, 0xf7, 0xb9, 0x68, 0x3a, 0xd2,
	0xe3, 0x3e, 0x39, 0xcd, 0x7d, 0xcd, 0x01, 0xdc, 0x75, 0x84, 0x83, 0x56, 0x26, 0x02, 0xed, 0x3e,
	0xcc, 0x70, 0x25, 0xcc, 0xe6, 0xea, 0xa3, 0xfe, 0xd4, 0xb8, 0x2f, 0x99, 0x2f, 0x6d, 0xcf, 0xd5,
	0x87, 0xce, 0x0c, 0xcd, 0x8f, 0x5c, 0xac, 0x1a, 0xd1, 0xed, 0x44, 0x22, 0x3d, 0xb4, 0xd7, 0x99,
	0x78, 0xa0, 0x4e, 0xd6, 0x64, 0x08, 0xda, 0x16, 0xe1, 0xdd, 0xad, 0x12, 0x74, 0xd3, 0xd5, 0x56,
	0x60, 0xf6, 0xd8, 0xb0, 0xb6, 0xdf, 0x6e, 0x34, 0xf4, 0x6f, 0x79, 0x0c, 0x5e, 0xb0, 0x66, 0xfa,
	0xe2, 0x4d, 0xf4, 0xd0, 0xae, 0xc3, 0x88, 0x17, 0x44, 0x68, 0x14, 0xf3, 0x5f, 0x5f, 0x31, 0x7b,
	0x9c, 0x37, 0x98, 0xe3, 0x93, 0x98, 0x61, 0x2f, 0x08, 0xe3, 0xa0, 0x80, 0x79, 0x28, 0x75, 0x81,
	0x89, 0xf2, 0x3b, 0x51, 0x8e, 0x1d, 0x7a, 0x85, 0x34, 0xb7, 0xa1, 0xbc, 0x2f, 0xb0, 0x44, 0xe4,
	0x1b, 0xa8, 0x02, 0x29, 0xba, 0x4c, 0xfa, 0xbd, 0x97, 0x14, 0xee, 0xa8, 0xac, 0xc8, 0x7a, 0x15,
	0xf4, 0x3e, 0xb1, 0x88, 0xfd, 0x07, 0xb1, 0x9f, 0x4a, 0x82, 0x42, 0x11, 0xb7, 0x60, 0x42, 0xf2,
	0xa4, 0x84, 0xff, 0xd3, 0x25, 0x14, 0x25, 0x8f, 0x09, 0x58, 0x84, 0xa9, 0x9e, 0x38, 0x44, 0xff,
	0x93, 0xe8, 0xb5, 0x38, 0x20, 0x24, 0xbf, 0x04, 0xf9, 0x76, 0x40, 0x49, 0xce, 0xa6, 0x53, 0xe6,
	0x94, 0x2f, 0x52, 0x9d, 0x85, 0xd1, 0x08, 0x45, 0x04, 0xbf, 0x88, 0x00, 0xc8, 0x1c, 0xa5, 0x56,
	0x0b, 0xd3, 0xe1, 0x2a, 0x21, 0x8e, 0xef, 0xda, 0xd2, 0x6b, 0x32, 0x3d, 0x17, 0x72, 0x9c, 0x4e,
	0x70, 0x28, 0x13, 0x2a, 0x6e, 0xb6, 0x88, 0x65, 0x5c, 0x01, 0x6f, 0x22, 0x6e, 0xd9, 0x77, 0xef,
	0xa1, 0x49, 0xbb, 0x0c, 0xd3, 0xbd, 0xb1, 0x88, 0xf9, 0x37, 0x31, 0x97, 0x13, 0x90, 0x28, 0xb1,
	0x25, 0x4c, 0x48, 0x5c, 0x40, 0x7e, 0x10, 0x01, 0x63, 0x92, 0x77, 0xd3, 0x5f, 0x80, 0xc9, 0x64,
	0x1c, 0x22, 0xff, 0x43, 0xe4, 0x13, 0x31, 0x77, 0x45, 0x3d, 0xf7, 0xa5, 0x00, 0x05, 0x8b, 0x05,
	0x2d, 0xee, 0x07, 0x4c, 0x3b, 0x0f, 0xd9, 0x70, 0x8e, 0xa3, 0x99, 0xaa, 0x18, 0xf1, 0x15, 0x41,
	0x33, 0xbe, 0xa6, 0xbe, 0x2d, 0x72, 0xd4, 0x76, 0xa1, 0xa4, 0x26, 0xd8, 0xee, 0x1a, 0x61, 0x9c,
	0x81, 0x0c, 0x82, 0x8d, 0x04, 0x38, 0x39, 0xe8, 0x5b, 0x78, 0xde, 0x3c, 0x3a, 0x5b, 0xe3, 0xcd,
	0xf8, 0x05, 0x6d, 0x09, 0xf2, 0xd1, 0xe6, 0xc0, 0x36, 0x57, 0x11, 0x67, 0x7b, 0x22, 0xd2, 0x5e,
	0xd9, 0xa2, 0x5f, 0xab, 0xe3, 0xae, 0x6d, 0x40, 0x46, 0xf0, 0x03, 0xec, 0x4c, 0x85, 0xba, 0x62,
	0x0c, 0xba, 0xe7, 0x8c, 0x4e, 0x1e, 0x0c, 0x8b, 0x1f, 0x58, 0x2a, 0x44, 0xe5, 0x73, 0x0e, 0x32,
	0x78, 0xd0, 0xa6, 0x20, 0x87, 0x47, 0xd5, 0x7b, 0xcf, 0x4d, 0x4c, 0x4d, 0xd6, 0xca, 0xe2, 0x11,
	0xdb, 0x6b, 0x07, 0xa6, 0x6a, 0x8f, 0x1d, 0x15, 0xb8, 0x9e, 0x28, 0xdf, 0x0b, 0x73, 0x90, 0xfa,
	0x95, 0x3b, 0xe8, 0xee, 0x2a, 0xae, 0xf7, 0x1b, 0xb3, 0x97, 0xe6, 0xc9, 0xe7, 0x6c, 0x09, 0x86,
	0xc3, 0xe6, 0xf7, 0x71, 0x8f, 0xea, 0xaf, 0xcc, 0x78, 0x4d, 0xa3, 0x00, 0x81, 0x14, 0x48, 0x4f,
	0xf8, 0x82, 0xf2, 0x36, 0x9d, 0xb0, 0x8f, 0x0f, 0x87, 0xed, 0xb5, 0x39, 0xf8, 0xb4, 0xad, 0x42,
	0xb1, 0xc6, 0x9b, 0xad, 0x06, 0x6e, 0x6d, 0xd2, 0xaf, 0xbf, 0x31, 0xd3, 0x17, 0xe2, 0x58, 0x07,
	0x13, 0x8a, 0xd7, 0xee, 0xf4, 0x5f, 0x75, 0x6f, 0xcd, 0x7f, 0xd9, 0x75, 0xbb, 0x50, 0xed, 0x5f,
	0x20, 0x1b, 0x2f, 0x0b, 0xfd, 0x5d, 0x7a, 0x52, 0xa6, 0xfb, 0x94, 0x68, 0x15, 0xb1, 0xb8, 0x46,
	0x0b, 0x87, 0x6b, 0xff, 0xfd, 0x00, 0xe2, 0xf2, 0x3c, 0xda, 0xfa, 0xcb, 0x50, 0x8c, 0xd7, 0x57,
	0xff, 0x90, 0x2e, 0x63, 0xb4, 0xbb, 0xb6, 0x6a, 0x61, 0xf4, 0x3c, 0x07, 0x3f, 0x0e, 0xd2, 0x21,
	0x3c, 0xfe, 0x04, 0x5c, 0x83, 0x52, 0x32, 0xd7, 0xfa, 0xa7, 0x74, 0x31, 0xc5, 0x78, 0x9e, 0x57,
	0x1e, 0x41, 0xd5, 0xe3, 0x89, 0x39, 0x3b, 0x7a, 0x15, 0x79, 0x78, 0xa3, 0xce, 0x03, 0xf7, 0x49,
	0xc7, 0xee, 0x9e, 0xec, 0x6d, 0x65, 0x2f, 0x17, 0xbe, 0x12, 0x2c, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0xf3, 0x7e, 0x33, 0xeb, 0x08, 0x00, 0x00,
}
