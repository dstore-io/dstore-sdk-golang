// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderSurchargeInfo_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderSurchargeInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderSurchargeInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderSurchargeInfo_Ad

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
	OrderId               *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderIdNull           bool                        `protobuf:"varint,1001,opt,name=order_id_null,json=orderIdNull" json:"order_id_null,omitempty"`
	PositionNo            *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=position_no,json=positionNo" json:"position_no,omitempty"`
	PositionNoNull        bool                        `protobuf:"varint,1002,opt,name=position_no_null,json=positionNoNull" json:"position_no_null,omitempty"`
	InformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull bool                        `protobuf:"varint,1003,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
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

func (m *Parameters) GetPositionNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PositionNo
	}
	return nil
}

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
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
	RowId                int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CreatedByUserId      *dstore_values.IntegerValue   `protobuf:"bytes,10001,opt,name=created_by_user_id,json=createdByUserId" json:"created_by_user_id,omitempty"`
	InformationTypeId    *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	PositionNo           *dstore_values.IntegerValue   `protobuf:"bytes,10003,opt,name=position_no,json=positionNo" json:"position_no,omitempty"`
	InformationType      *dstore_values.StringValue    `protobuf:"bytes,10004,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	Information          *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=information" json:"information,omitempty"`
	CreatedAtDateAndTime *dstore_values.TimestampValue `protobuf:"bytes,10006,opt,name=created_at_date_and_time,json=createdAtDateAndTime" json:"created_at_date_and_time,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCreatedByUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CreatedByUserId
	}
	return nil
}

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Response_Row) GetPositionNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PositionNo
	}
	return nil
}

func (m *Response_Row) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func (m *Response_Row) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *Response_Row) GetCreatedAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.CreatedAtDateAndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderSurchargeInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderSurchargeInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderSurchargeInfo_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetOrderSurchargeInfo_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x55, 0x9b, 0x2f, 0x4d, 0x34, 0xd1, 0x47, 0x83, 0x0b, 0xc8, 0x24, 0x02, 0xa1, 0x76, 0x03,
	0x1b, 0x07, 0x81, 0x04, 0x08, 0xd1, 0x45, 0x2a, 0x7e, 0x14, 0x41, 0x0d, 0x32, 0xa5, 0x12, 0x08,
	0x69, 0xe4, 0x66, 0x2e, 0xc6, 0x22, 0x9e, 0x89, 0x66, 0x26, 0x54, 0x79, 0x0b, 0xfe, 0x9f, 0x8c,
	0x0d, 0x8f, 0x00, 0x2c, 0x79, 0x01, 0xee, 0x78, 0xc6, 0x8a, 0xed, 0x56, 0x84, 0xb2, 0x89, 0x73,
	0x7d, 0xef, 0x39, 0xe7, 0xfe, 0x1c, 0x99, 0xdc, 0x66, 0x4a, 0x0b, 0x09, 0x03, 0xe0, 0x49, 0xca,
	0x61, 0x30, 0x95, 0x62, 0x0c, 0x6c, 0x26, 0x41, 0x0d, 0x44, 0x46, 0x1f, 0x80, 0x7e, 0x2c, 0x19,
	0xc8, 0xa7, 0x33, 0x39, 0x7e, 0x1d, 0xcb, 0x04, 0x46, 0xfc, 0x95, 0xa0, 0x43, 0x16, 0x60, 0x99,
	0x16, 0xde, 0x15, 0x8b, 0x0d, 0x2c, 0x36, 0xf8, 0x03, 0xa0, 0xb7, 0xe1, 0x64, 0xde, 0xc6, 0x93,
	0x19, 0x28, 0x8b, 0xef, 0x9d, 0xaf, 0x6a, 0x83, 0x94, 0x42, 0xba, 0x54, 0xbf, 0x9a, 0xca, 0x40,
	0xa9, 0x38, 0x01, 0x97, 0xdc, 0xaa, 0x27, 0x75, 0x9c, 0xa2, 0x8e, 0xcc, 0x62, 0x9d, 0x0a, 0x6e,
	0x8b, 0x36, 0xbf, 0xad, 0x12, 0xf2, 0x24, 0x96, 0x31, 0x66, 0x41, 0x2a, 0xef, 0x06, 0x69, 0x0b,
	0xd3, 0x19, 0x4d, 0x99, 0xbf, 0x72, 0x69, 0xe5, 0x72, 0xe7, 0x5a, 0x3f, 0x70, 0xed, 0xbb, 0x9e,
	0x52, 0xae, 0x21, 0x01, 0xb9, 0x6f, 0xa2, 0xa8, 0x95, 0x17, 0x8f, 0x98, 0xb7, 0x45, 0xfe, 0x2f,
	0x70, 0x94, 0xcf, 0x26, 0x13, 0xff, 0x7b, 0x0b, 0xd1, 0xed, 0xa8, 0xe3, 0x0a, 0x42, 0x7c, 0xe7,
	0xdd, 0x21, 0x9d, 0xa9, 0x50, 0xa9, 0x51, 0xa7, 0x5c, 0xf8, 0xab, 0xcb, 0xf9, 0x49, 0x51, 0x1f,
	0xe2, 0x1a, 0x49, 0xb7, 0x84, 0xb6, 0x2a, 0x3f, 0xac, 0xca, 0xa9, 0x45, 0x59, 0x2e, 0xf4, 0x90,
	0x6c, 0x94, 0x26, 0xa5, 0x7a, 0x3e, 0x05, 0x33, 0x50, 0x63, 0xb9, 0xe0, 0xe9, 0x12, 0x6e, 0x0f,
	0x61, 0x38, 0xda, 0x2d, 0xe2, 0x1f, 0x43, 0x66, 0xf5, 0x7f, 0x5a, 0xfd, 0xb3, 0x47, 0x50, 0xa6,
	0x8d, 0xcd, 0x5f, 0x4d, 0xd2, 0x8e, 0x40, 0x4d, 0x05, 0x57, 0xe0, 0x5d, 0x25, 0xcd, 0xfc, 0x72,
	0x6e, 0xad, 0xbd, 0xa0, 0xea, 0x0a, 0x7b, 0xd5, 0x7b, 0xe6, 0x37, 0xb2, 0x85, 0xde, 0x73, 0xd2,
	0x35, 0x37, 0xa3, 0x25, 0x72, 0xdc, 0x59, 0x03, 0xc1, 0x41, 0x0d, 0x5c, 0x3f, 0xed, 0x2e, 0xc6,
	0xa3, 0x45, 0x1c, 0xad, 0x67, 0xd5, 0x17, 0x38, 0x53, 0xcb, 0x79, 0x05, 0x97, 0x62, 0x18, 0x2f,
	0x1e, 0x61, 0xb4, 0x4e, 0xda, 0xb5, 0xcf, 0xa8, 0x28, 0xf7, 0x46, 0xa4, 0x21, 0xc5, 0xa1, 0xff,
	0x5f, 0x8e, 0xba, 0x19, 0xfc, 0xb5, 0xb5, 0x83, 0x62, 0x11, 0x41, 0x24, 0x0e, 0x23, 0xc3, 0xd1,
	0xfb, 0xda, 0x20, 0x0d, 0x0c, 0xbc, 0x73, 0x64, 0x0d, 0x43, 0x73, 0xa0, 0x77, 0x21, 0xee, 0xa6,
	0x19, 0x35, 0x31, 0xc4, 0xc5, 0x8f, 0x88, 0x37, 0x96, 0x10, 0x6b, 0x60, 0xf4, 0x60, 0x4e, 0x67,
	0xca, 0xba, 0xf2, 0x7d, 0xb8, 0xfc, 0x8a, 0xeb, 0x0e, 0xb7, 0x33, 0x7f, 0xa6, 0x72, 0x7b, 0x3e,
	0x3a, 0xde, 0x10, 0x1f, 0xc2, 0x7f, 0x72, 0xc4, 0x76, 0xd5, 0xc7, 0x1f, 0xc3, 0x93, 0x19, 0xf9,
	0x3e, 0xe9, 0xd6, 0x9b, 0xf1, 0x3f, 0x85, 0x55, 0x57, 0x38, 0x0e, 0xa5, 0x65, 0xca, 0x13, 0x37,
	0x54, 0xad, 0x11, 0xd3, 0x46, 0xd9, 0x1a, 0x9f, 0x97, 0x53, 0x94, 0xeb, 0xbd, 0x7d, 0xe2, 0x17,
	0xeb, 0x8d, 0x35, 0x65, 0xf8, 0x87, 0xc6, 0x9c, 0x51, 0x9d, 0x66, 0xe0, 0x7f, 0xb1, 0x5c, 0x17,
	0x6a, 0x5c, 0x26, 0xa7, 0x74, 0x9c, 0x4d, 0x2d, 0xdd, 0x19, 0x87, 0x1f, 0xea, 0xbb, 0xf8, 0x18,
	0x72, 0xb6, 0x87, 0xf9, 0x9d, 0x97, 0xa4, 0x9f, 0x8a, 0x9a, 0x31, 0x16, 0xdf, 0xcb, 0x17, 0xdb,
	0x89, 0x50, 0xec, 0x4d, 0x91, 0x67, 0x27, 0xfc, 0xa4, 0x1e, 0xac, 0xe5, 0x9f, 0xad, 0xeb, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x36, 0xe5, 0xf3, 0x91, 0x05, 0x00, 0x00,
}