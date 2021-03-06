// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyUnitConverts_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyUnitConverts_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyUnitConverts_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyUnitConverts_Ad

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
	FromUnitId          *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=from_unit_id,json=fromUnitId" json:"from_unit_id,omitempty"`
	FromUnitIdNull      bool                          `protobuf:"varint,1001,opt,name=from_unit_id_null,json=fromUnitIdNull" json:"from_unit_id_null,omitempty"`
	ToUnitId            *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=to_unit_id,json=toUnitId" json:"to_unit_id,omitempty"`
	ToUnitIdNull        bool                          `protobuf:"varint,1002,opt,name=to_unit_id_null,json=toUnitIdNull" json:"to_unit_id_null,omitempty"`
	ConvertFactor       *dstore_values.DecimalValue   `protobuf:"bytes,3,opt,name=convert_factor,json=convertFactor" json:"convert_factor,omitempty"`
	ConvertFactorNull   bool                          `protobuf:"varint,1003,opt,name=convert_factor_null,json=convertFactorNull" json:"convert_factor_null,omitempty"`
	FromDate            *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull        bool                          `protobuf:"varint,1004,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate              *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull          bool                          `protobuf:"varint,1005,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	DateIndependent     *dstore_values.BooleanValue   `protobuf:"bytes,6,opt,name=date_independent,json=dateIndependent" json:"date_independent,omitempty"`
	DateIndependentNull bool                          `protobuf:"varint,1006,opt,name=date_independent_null,json=dateIndependentNull" json:"date_independent_null,omitempty"`
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

func (m *Parameters) GetConvertFactor() *dstore_values.DecimalValue {
	if m != nil {
		return m.ConvertFactor
	}
	return nil
}

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetDateIndependent() *dstore_values.BooleanValue {
	if m != nil {
		return m.DateIndependent
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyUnitConverts_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyUnitConverts_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyUnitConverts_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_ModifyUnitConverts_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xd5, 0x86, 0xa4, 0x61, 0x08, 0x49, 0xeb, 0x08, 0x14, 0x12, 0x15, 0x41, 0x11, 0x15,
	0xe2, 0xb0, 0x41, 0x14, 0xf1, 0x51, 0x89, 0x03, 0x05, 0x2a, 0xe5, 0x90, 0x0a, 0xad, 0x04, 0x52,
	0xb9, 0xac, 0xdc, 0xd8, 0x89, 0x2c, 0xb2, 0x76, 0x64, 0x3b, 0xad, 0x78, 0x05, 0x4e, 0x3c, 0x0e,
	0xaf, 0xc4, 0xe7, 0x33, 0x60, 0x7b, 0x9c, 0xa6, 0xbb, 0x14, 0x51, 0x2e, 0x9b, 0x78, 0x67, 0xfe,
	0xff, 0x9f, 0x3d, 0x3b, 0x63, 0x78, 0xc2, 0x8c, 0x55, 0x9a, 0xf7, 0xb9, 0x9c, 0x08, 0xc9, 0xfb,
	0x33, 0xad, 0x46, 0x9c, 0xcd, 0x35, 0x37, 0xfd, 0x5c, 0x64, 0x43, 0xc5, 0xc4, 0xf8, 0xe3, 0x5b,
	0x29, 0xec, 0x4b, 0x25, 0x8f, 0xb9, 0xb6, 0x26, 0x7b, 0xc1, 0x12, 0x97, 0x63, 0x15, 0xd9, 0x46,
	0x61, 0x82, 0xc2, 0xe4, 0x6f, 0xd9, 0xdd, 0x76, 0x04, 0x1c, 0xd3, 0xe9, 0x9c, 0x1b, 0x14, 0x77,
	0x6f, 0x14, 0xa9, 0x5c, 0x6b, 0xa5, 0x63, 0xa8, 0x57, 0x0c, 0xe5, 0xdc, 0x18, 0x3a, 0xe1, 0x31,
	0x78, 0xa7, 0x1c, 0xb4, 0x54, 0xc8, 0xb1, 0xd2, 0x39, 0xb5, 0x42, 0x49, 0x4c, 0xda, 0xfa, 0x54,
	0x05, 0x78, 0x43, 0x35, 0x75, 0x51, 0xae, 0x0d, 0x79, 0x0e, 0x8d, 0xb1, 0x56, 0x79, 0x36, 0x77,
	0x1b, 0xcb, 0x04, 0xeb, 0xac, 0xdc, 0x5a, 0xb9, 0x77, 0xe5, 0x61, 0x2f, 0x89, 0xfb, 0x8f, 0xfb,
	0x12, 0xd2, 0xf2, 0x09, 0xd7, 0xef, 0xfc, 0x2a, 0x05, 0x2f, 0xf0, 0x07, 0x19, 0x30, 0x72, 0x1f,
	0x36, 0xce, 0xca, 0x33, 0x39, 0x9f, 0x4e, 0x3b, 0x5f, 0xd7, 0x9c, 0x49, 0x3d, 0x6d, 0x2e, 0xf3,
	0x0e, 0xdc, 0x6b, 0xf2, 0x0c, 0xc0, 0xaa, 0x53, 0xd0, 0xea, 0xbf, 0x41, 0x75, 0xab, 0x22, 0x66,
	0x1b, 0x5a, 0x4b, 0x29, 0x42, 0xbe, 0x21, 0xa4, 0xb1, 0xc8, 0x09, 0x88, 0x3d, 0x68, 0x8e, 0xb0,
	0xba, 0xd9, 0x98, 0x8e, 0x9c, 0x71, 0xa7, 0x72, 0x2e, 0x86, 0xf1, 0x91, 0xc8, 0xe9, 0x14, 0x31,
	0x57, 0xa3, 0x64, 0x3f, 0x28, 0x48, 0x1f, 0xda, 0x45, 0x0f, 0xe4, 0x7d, 0x47, 0xde, 0x46, 0x21,
	0x39, 0x40, 0x77, 0xe1, 0x72, 0xa8, 0x01, 0xa3, 0x96, 0x77, 0x2e, 0x05, 0xde, 0x66, 0x89, 0x67,
	0x85, 0xfb, 0x52, 0x96, 0xe6, 0xb3, 0x78, 0x30, 0x9f, 0xff, 0xca, 0xa5, 0x93, 0xbb, 0xd0, 0x3c,
	0xd5, 0x22, 0xe7, 0x47, 0x3c, 0xd7, 0x22, 0x25, 0x20, 0x1e, 0xc3, 0x9a, 0x3b, 0x7f, 0x00, 0x54,
	0x2f, 0x02, 0xa8, 0x59, 0x15, 0xec, 0x6f, 0x43, 0x23, 0xea, 0xd0, 0xfc, 0x27, 0x9a, 0x03, 0x86,
	0x83, 0xf5, 0x3e, 0xac, 0x87, 0xb8, 0x90, 0x8c, 0xcf, 0xb8, 0x7b, 0x48, 0xdb, 0xa9, 0x9d, 0x5b,
	0xb4, 0x23, 0xa5, 0xa6, 0x9c, 0x4a, 0x24, 0xb4, 0xbc, 0x68, 0xb0, 0xd4, 0x90, 0x1d, 0xb8, 0x56,
	0xf6, 0x41, 0xe6, 0x2f, 0x64, 0xb6, 0x4b, 0x02, 0x0f, 0xdf, 0xfa, 0xb2, 0x0a, 0xf5, 0x94, 0x9b,
	0x99, 0x92, 0x86, 0x93, 0x07, 0x50, 0x0d, 0xad, 0x1e, 0x7b, 0xb0, 0x9b, 0x14, 0x67, 0x08, 0xc7,
	0xe0, 0xb5, 0x7f, 0xa6, 0x98, 0x48, 0x0e, 0x61, 0xdd, 0x37, 0x79, 0x76, 0xa6, 0xcb, 0x5d, 0x5f,
	0x55, 0x9c, 0x38, 0x29, 0x89, 0xcb, 0xb3, 0x30, 0x74, 0xeb, 0xc1, 0x72, 0x9d, 0xb6, 0xf2, 0xe2,
	0x0b, 0xf2, 0x14, 0xd6, 0xe2, 0x70, 0xb9, 0x16, 0xf2, 0x8e, 0x37, 0xff, 0x70, 0xc4, 0xd1, 0x1b,
	0xe2, 0x6f, 0xba, 0x48, 0x77, 0x05, 0xad, 0x68, 0x75, 0xe2, 0x1a, 0xc1, 0xab, 0x1e, 0x25, 0x17,
	0xbb, 0x08, 0x92, 0x45, 0x15, 0x92, 0x54, 0x9d, 0xa4, 0xde, 0xa0, 0xbb, 0x09, 0x15, 0xf7, 0x9f,
	0x5c, 0x87, 0x9a, 0x5b, 0xf9, 0x89, 0xf9, 0x7c, 0xe0, 0xea, 0x52, 0x4d, 0xab, 0x6e, 0x39, 0x60,
	0x7b, 0x87, 0xd0, 0x13, 0xaa, 0xe4, 0xbe, 0xbc, 0x9f, 0xde, 0xef, 0x4e, 0x94, 0x61, 0x1f, 0x16,
	0x71, 0xf6, 0x3f, 0x57, 0xd8, 0x51, 0x2d, 0xdc, 0x14, 0x3b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0x82, 0xe1, 0x62, 0xfe, 0x04, 0x00, 0x00,
}
