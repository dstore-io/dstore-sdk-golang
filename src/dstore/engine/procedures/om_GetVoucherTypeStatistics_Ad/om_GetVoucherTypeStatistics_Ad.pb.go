// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetVoucherTypeStatistics_Ad.proto
// DO NOT EDIT!

/*
Package om_GetVoucherTypeStatistics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetVoucherTypeStatistics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetVoucherTypeStatistics_Ad

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
	VoucherTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=voucher_type_id,json=voucherTypeId" json:"voucher_type_id,omitempty"`
	VoucherTypeIdNull bool                        `protobuf:"varint,1001,opt,name=voucher_type_id_null,json=voucherTypeIdNull" json:"voucher_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetVoucherTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VoucherTypeId
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
	RowId                     int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	LastCodeUsedAtDateAndTime *dstore_values.TimestampValue `protobuf:"bytes,10001,opt,name=last_code_used_at_date_and_time,json=lastCodeUsedAtDateAndTime" json:"last_code_used_at_date_and_time,omitempty"`
	NumberOfUnusedCodes       *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=number_of_unused_codes,json=numberOfUnusedCodes" json:"number_of_unused_codes,omitempty"`
	NumberOfCodes             *dstore_values.IntegerValue   `protobuf:"bytes,10003,opt,name=number_of_codes,json=numberOfCodes" json:"number_of_codes,omitempty"`
	NumberOfUsedCodes         *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=number_of_used_codes,json=numberOfUsedCodes" json:"number_of_used_codes,omitempty"`
	NumberOfMaxUsedCodes      *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=number_of_max_used_codes,json=numberOfMaxUsedCodes" json:"number_of_max_used_codes,omitempty"`
	MaxValidUntil             *dstore_values.TimestampValue `protobuf:"bytes,10006,opt,name=max_valid_until,json=maxValidUntil" json:"max_valid_until,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetLastCodeUsedAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.LastCodeUsedAtDateAndTime
	}
	return nil
}

func (m *Response_Row) GetNumberOfUnusedCodes() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfUnusedCodes
	}
	return nil
}

func (m *Response_Row) GetNumberOfCodes() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfCodes
	}
	return nil
}

func (m *Response_Row) GetNumberOfUsedCodes() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfUsedCodes
	}
	return nil
}

func (m *Response_Row) GetNumberOfMaxUsedCodes() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfMaxUsedCodes
	}
	return nil
}

func (m *Response_Row) GetMaxValidUntil() *dstore_values.TimestampValue {
	if m != nil {
		return m.MaxValidUntil
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetVoucherTypeStatistics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetVoucherTypeStatistics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetVoucherTypeStatistics_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x15, 0xd2, 0xa4, 0xd5, 0xa2, 0x28, 0xd4, 0x8d, 0x2a, 0x37, 0x11, 0x50, 0x95, 0x0b,
	0x07, 0xe4, 0x20, 0x38, 0x80, 0x90, 0x38, 0x84, 0x16, 0x50, 0x25, 0x12, 0xc0, 0x24, 0x91, 0xe0,
	0xb2, 0xda, 0x64, 0x27, 0xc1, 0x52, 0xbc, 0x1b, 0xed, 0xae, 0xd3, 0x72, 0xe6, 0xc0, 0x95, 0xef,
	0x77, 0xe0, 0xd1, 0x78, 0x0b, 0x66, 0xbd, 0x4e, 0x9d, 0x18, 0x09, 0x02, 0x17, 0x5b, 0xe3, 0x99,
	0xff, 0x6f, 0xbe, 0xd6, 0x4b, 0x1e, 0x72, 0x6d, 0xa4, 0x82, 0x36, 0x88, 0x69, 0x24, 0xa0, 0x3d,
	0x57, 0x72, 0x0c, 0x3c, 0x51, 0xa0, 0xdb, 0x32, 0xa6, 0x4f, 0xc1, 0x0c, 0x65, 0x32, 0x7e, 0x0b,
	0xaa, 0xff, 0x6e, 0x0e, 0xaf, 0x0c, 0x33, 0x91, 0x36, 0xd1, 0x58, 0xd3, 0x0e, 0x0f, 0x30, 0xd2,
	0x48, 0xef, 0x96, 0x93, 0x07, 0x4e, 0x1e, 0xfc, 0x59, 0xd3, 0xdc, 0xcb, 0x92, 0x2d, 0xd8, 0x2c,
	0x01, 0xed, 0x10, 0xcd, 0x83, 0xf5, 0x0a, 0x40, 0x29, 0xa9, 0x32, 0x57, 0x6b, 0xdd, 0x15, 0x83,
	0xd6, 0x6c, 0x0a, 0x99, 0xf3, 0x46, 0xd1, 0x69, 0x58, 0x24, 0x26, 0x52, 0xc5, 0x98, 0x51, 0x0a,
	0x17, 0x74, 0xf4, 0xbe, 0x44, 0xc8, 0x0b, 0xa6, 0x18, 0x7a, 0x41, 0x69, 0xef, 0x98, 0xd4, 0x17,
	0xae, 0x38, 0x6a, 0xb0, 0x3a, 0x1a, 0x71, 0xbf, 0x74, 0x58, 0xba, 0x79, 0xf9, 0x4e, 0x2b, 0xc8,
	0x1a, 0xc9, 0x4a, 0x8b, 0x84, 0x81, 0x29, 0xa8, 0xa1, 0xb5, 0xc2, 0xda, 0x22, 0x6f, 0xe8, 0x94,
	0x7b, 0xb7, 0x49, 0xa3, 0x00, 0xa1, 0x22, 0x99, 0xcd, 0xfc, 0x9f, 0xdb, 0x88, 0xda, 0x09, 0x77,
	0xd7, 0xa2, 0x7b, 0xe8, 0x39, 0xfa, 0x51, 0x25, 0x3b, 0x21, 0xe8, 0xb9, 0x14, 0x1a, 0x50, 0x5e,
	0x49, 0x7b, 0xcc, 0x32, 0x37, 0x83, 0xf5, 0x11, 0xba, 0xfe, 0x1f, 0xdb, 0x67, 0xe8, 0x02, 0xbd,
	0xd7, 0xe4, 0x8a, 0xed, 0x8e, 0xae, 0xb4, 0xe7, 0x5f, 0x3a, 0x2c, 0xa3, 0x38, 0x28, 0x88, 0x8b,
	0x43, 0xe8, 0xa2, 0x7d, 0x9a, 0xdb, 0x61, 0x3d, 0x5e, 0xff, 0xe0, 0xdd, 0x27, 0xdb, 0xd9, 0x54,
	0xfd, 0x72, 0x4a, 0xbc, 0xf6, 0x1b, 0xd1, 0xcd, 0xbc, 0xeb, 0xde, 0xe1, 0x32, 0xdc, 0x7b, 0x46,
	0xca, 0x4a, 0x9e, 0xf9, 0x5b, 0xa9, 0xea, 0x41, 0xf0, 0x2f, 0xe7, 0x20, 0x58, 0xce, 0x22, 0x08,
	0xe5, 0x59, 0x68, 0x31, 0xcd, 0x0f, 0x5b, 0xa4, 0x8c, 0x86, 0xb7, 0x4f, 0xaa, 0x68, 0xda, 0xbd,
	0x7c, 0xec, 0xe1, 0x78, 0x2a, 0x61, 0x05, 0x4d, 0x9c, 0xf9, 0x88, 0x5c, 0x9f, 0x31, 0x6d, 0xe8,
	0x58, 0x72, 0xa0, 0x89, 0x06, 0x4e, 0x99, 0xa1, 0x9c, 0x19, 0xa0, 0x4c, 0x70, 0x6a, 0xa2, 0x18,
	0xfc, 0x4f, 0xbd, 0x74, 0x9e, 0x57, 0x0b, 0x9b, 0xb4, 0x3e, 0x6d, 0x58, 0x3c, 0x77, 0xbb, 0x3c,
	0xb0, 0x98, 0x63, 0xa4, 0x0c, 0x10, 0xd2, 0x31, 0x27, 0x88, 0xe8, 0x08, 0xde, 0xc7, 0x20, 0xef,
	0x25, 0xd9, 0x17, 0x49, 0x3c, 0xc2, 0xb5, 0xca, 0x09, 0x4d, 0x44, 0x9a, 0xc5, 0xe6, 0xd3, 0xfe,
	0xe7, 0xde, 0xdf, 0x0f, 0xc9, 0x9e, 0xd3, 0x3e, 0x9f, 0x0c, 0x52, 0xa5, 0x4d, 0xa1, 0xbd, 0x13,
	0x52, 0xcf, 0x91, 0x8e, 0xf5, 0x65, 0x03, 0x56, 0x6d, 0xc9, 0x72, 0x94, 0x2e, 0x69, 0xac, 0x14,
	0x96, 0x97, 0xf5, 0x75, 0x03, 0xd4, 0xee, 0x45, 0x59, 0x17, 0x45, 0xf5, 0x89, 0x9f, 0xe3, 0x62,
	0x76, 0xbe, 0x8a, 0xfc, 0xb6, 0x01, 0xb2, 0xb1, 0x44, 0x76, 0xd9, 0x79, 0x4e, 0x7d, 0x42, 0xea,
	0x96, 0x85, 0x02, 0xfc, 0x1f, 0x12, 0x61, 0xa2, 0x99, 0xff, 0x7d, 0xa3, 0x8d, 0xd4, 0x50, 0x36,
	0xb4, 0xaa, 0x81, 0x15, 0x3d, 0x1a, 0x92, 0x56, 0x24, 0x0b, 0xc7, 0x29, 0xbf, 0x95, 0xde, 0xdc,
	0xfb, 0xcf, 0xfb, 0x6a, 0x54, 0x4d, 0x2f, 0x84, 0xbb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67,
	0x84, 0x5f, 0x35, 0xf1, 0x04, 0x00, 0x00,
}