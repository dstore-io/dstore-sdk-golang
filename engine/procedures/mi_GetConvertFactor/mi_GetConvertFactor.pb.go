// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetConvertFactor.proto
// DO NOT EDIT!

/*
Package mi_GetConvertFactor is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetConvertFactor.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetConvertFactor

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
	FromUnitId     *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=from_unit_id,json=fromUnitId" json:"from_unit_id,omitempty"`
	FromUnitIdNull bool                          `protobuf:"varint,1001,opt,name=from_unit_id_null,json=fromUnitIdNull" json:"from_unit_id_null,omitempty"`
	ToUnitId       *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=to_unit_id,json=toUnitId" json:"to_unit_id,omitempty"`
	ToUnitIdNull   bool                          `protobuf:"varint,1002,opt,name=to_unit_id_null,json=toUnitIdNull" json:"to_unit_id_null,omitempty"`
	Date           *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
	DateNull       bool                          `protobuf:"varint,1003,opt,name=date_null,json=dateNull" json:"date_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetFromUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromUnitId
	}
	return nil
}

func (m *Parameters) GetToUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToUnitId
	}
	return nil
}

func (m *Parameters) GetDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.Date
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Factor1         *dstore_values.DecimalValue                      `protobuf:"bytes,101,opt,name=factor1" json:"factor1,omitempty"`
	Factor2         *dstore_values.DecimalValue                      `protobuf:"bytes,102,opt,name=factor2" json:"factor2,omitempty"`
	Divisor1        *dstore_values.DecimalValue                      `protobuf:"bytes,103,opt,name=divisor1" json:"divisor1,omitempty"`
	Divisor2        *dstore_values.DecimalValue                      `protobuf:"bytes,104,opt,name=divisor2" json:"divisor2,omitempty"`
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

func (m *Response) GetFactor1() *dstore_values.DecimalValue {
	if m != nil {
		return m.Factor1
	}
	return nil
}

func (m *Response) GetFactor2() *dstore_values.DecimalValue {
	if m != nil {
		return m.Factor2
	}
	return nil
}

func (m *Response) GetDivisor1() *dstore_values.DecimalValue {
	if m != nil {
		return m.Divisor1
	}
	return nil
}

func (m *Response) GetDivisor2() *dstore_values.DecimalValue {
	if m != nil {
		return m.Divisor2
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetConvertFactor.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetConvertFactor.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetConvertFactor.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetConvertFactor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xb5, 0xb5, 0x5d, 0xc3, 0x61, 0x62, 0x60, 0x24, 0x14, 0x52, 0x86, 0x60, 0x48, 0x08,
	0x71, 0x91, 0xb2, 0xa0, 0x09, 0xb8, 0xe0, 0x02, 0x10, 0xa0, 0x5e, 0x6c, 0x42, 0x11, 0x20, 0xc1,
	0x4d, 0x64, 0xea, 0xd3, 0x60, 0xd1, 0xd8, 0x95, 0xed, 0x76, 0xaf, 0x81, 0xc4, 0x63, 0xf0, 0x64,
	0xc0, 0x4b, 0x60, 0xc7, 0xc9, 0xd2, 0x66, 0x95, 0xda, 0x9b, 0x46, 0xf6, 0xf9, 0xbf, 0xff, 0xc4,
	0xe7, 0x77, 0x03, 0x09, 0xd3, 0x46, 0x2a, 0x1c, 0xa2, 0xc8, 0xb9, 0xc0, 0xe1, 0x4c, 0xc9, 0x31,
	0xb2, 0xb9, 0x42, 0x3d, 0x2c, 0x78, 0xf6, 0x1e, 0xcd, 0x1b, 0x29, 0x16, 0xa8, 0xcc, 0x3b, 0x3a,
	0xb6, 0xb2, 0xd8, 0x96, 0x8d, 0x24, 0xf7, 0x3d, 0x13, 0x7b, 0x26, 0x5e, 0x23, 0x8c, 0x6e, 0x56,
	0xb6, 0x0b, 0x3a, 0x9d, 0xa3, 0xf6, 0x5c, 0x74, 0x7b, 0xb5, 0x17, 0x2a, 0x55, 0x5b, 0x46, 0x83,
	0xd5, 0x52, 0x81, 0x5a, 0xd3, 0x1c, 0xab, 0xe2, 0x83, 0x76, 0xd1, 0x50, 0x2e, 0x26, 0x52, 0x15,
	0xd4, 0x70, 0x29, 0xbc, 0xe8, 0xe8, 0xf7, 0x2e, 0xc0, 0x07, 0xaa, 0xa8, 0xad, 0xa2, 0xd2, 0xe4,
	0x25, 0xec, 0x4f, 0x94, 0x2c, 0xb2, 0xb9, 0xe0, 0x26, 0xe3, 0x2c, 0xdc, 0xb9, 0xb7, 0xf3, 0xe8,
	0x6a, 0x32, 0x88, 0xab, 0x57, 0xaf, 0xde, 0x8b, 0x0b, 0x83, 0x39, 0xaa, 0xcf, 0x6e, 0x95, 0x82,
	0x03, 0x3e, 0x59, 0xfd, 0x88, 0x91, 0xc7, 0x70, 0x63, 0x19, 0xcf, 0xc4, 0x7c, 0x3a, 0x0d, 0xff,
	0xf4, 0xad, 0x49, 0x90, 0x5e, 0x6b, 0x74, 0x67, 0x76, 0x9b, 0xbc, 0x00, 0x30, 0xf2, 0xa2, 0xd1,
	0xee, 0xe6, 0x46, 0x81, 0x91, 0x55, 0x9b, 0x87, 0x70, 0xd0, 0xa0, 0xbe, 0xc9, 0x5f, 0xdf, 0x64,
	0xbf, 0xd6, 0x94, 0x2d, 0x8e, 0xa1, 0xcb, 0xa8, 0xc1, 0xb0, 0x53, 0x9a, 0x1f, 0xb6, 0xcc, 0x0d,
	0xb7, 0xf3, 0x32, 0xb4, 0x98, 0x79, 0xfb, 0x52, 0x4a, 0xee, 0xc0, 0x15, 0xf7, 0xf4, 0xa6, 0xff,
	0xbc, 0x69, 0xe0, 0x76, 0x9c, 0xe1, 0xd1, 0xaf, 0x2e, 0x04, 0x29, 0xea, 0x99, 0x14, 0x1a, 0xc9,
	0x13, 0xe8, 0x95, 0x59, 0x54, 0x43, 0x8a, 0xe2, 0xd5, 0x7c, 0x7d, 0x4e, 0x6f, 0xdd, 0x6f, 0xea,
	0x85, 0xe4, 0x0b, 0x5c, 0x77, 0x29, 0x64, 0x4b, 0x31, 0xd8, 0x83, 0x77, 0x2c, 0x1c, 0xb7, 0xe0,
	0x76, 0x58, 0xa7, 0x76, 0x3d, 0x6a, 0xd6, 0xe9, 0x41, 0xb1, 0xba, 0x41, 0x9e, 0x43, 0xbf, 0x4a,
	0xdf, 0x9e, 0xd6, 0x39, 0xde, 0xbd, 0xe4, 0xe8, 0xef, 0xc6, 0xa9, 0x7f, 0xa6, 0xb5, 0x9c, 0xbc,
	0x82, 0x8e, 0x92, 0xe7, 0x61, 0xb7, 0xa4, 0x86, 0xf1, 0xc6, 0x4b, 0x1a, 0xd7, 0x03, 0x88, 0x53,
	0x79, 0x9e, 0x3a, 0x96, 0x9c, 0x40, 0x7f, 0x52, 0xd6, 0x8e, 0x43, 0x5c, 0x9b, 0x23, 0xc3, 0x31,
	0x2f, 0xe8, 0xd4, 0x0f, 0xba, 0xd6, 0x36, 0x58, 0x12, 0x4e, 0xb6, 0xc6, 0x12, 0xf2, 0x0c, 0x02,
	0xc6, 0x17, 0x5c, 0xbb, 0x76, 0xf9, 0x66, 0xee, 0x42, 0xbc, 0x04, 0x26, 0xe1, 0xf7, 0xed, 0xc1,
	0x24, 0x3a, 0x84, 0x8e, 0x3d, 0x2b, 0xb9, 0x05, 0x7b, 0xf6, 0xb4, 0xee, 0xb6, 0xfe, 0x3c, 0xb3,
	0x78, 0x2f, 0xed, 0xd9, 0xe5, 0x88, 0xbd, 0xfe, 0x08, 0x03, 0x2e, 0x5b, 0x83, 0x6b, 0xbe, 0x08,
	0x5f, 0x4f, 0x72, 0xa9, 0xd9, 0x8f, 0xba, 0xce, 0xb6, 0xfc, 0x68, 0x7c, 0xdb, 0x2b, 0xff, 0xa0,
	0x4f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x68, 0x6a, 0xef, 0x5c, 0x6b, 0x04, 0x00, 0x00,
}
