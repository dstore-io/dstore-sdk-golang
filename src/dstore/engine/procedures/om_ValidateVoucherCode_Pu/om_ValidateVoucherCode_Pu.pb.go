// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ValidateVoucherCode_Pu.proto
// DO NOT EDIT!

/*
Package om_ValidateVoucherCode_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ValidateVoucherCode_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ValidateVoucherCode_Pu

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
	UniqueId        *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull    bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	VoucherCode     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=voucher_code,json=voucherCode" json:"voucher_code,omitempty"`
	VoucherCodeNull bool                        `protobuf:"varint,1002,opt,name=voucher_code_null,json=voucherCodeNull" json:"voucher_code_null,omitempty"`
	PersonId        *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull    bool                        `protobuf:"varint,1003,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetVoucherCode() *dstore_values.StringValue {
	if m != nil {
		return m.VoucherCode
	}
	return nil
}

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ValidateVoucherCode_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ValidateVoucherCode_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ValidateVoucherCode_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdf, 0xea, 0xd3, 0x30,
	0x14, 0x66, 0xab, 0xfb, 0x97, 0x0d, 0xa7, 0x11, 0xa4, 0x76, 0x28, 0xa2, 0x88, 0x82, 0x90, 0xc9,
	0x44, 0xdc, 0x8d, 0x37, 0x8a, 0xc8, 0x2e, 0x36, 0x66, 0x2e, 0x06, 0x7a, 0x53, 0xe2, 0x7a, 0xac,
	0x81, 0x36, 0x99, 0x49, 0xbb, 0xbd, 0x86, 0xaf, 0xe3, 0x23, 0xa9, 0x2f, 0x61, 0xd2, 0xa4, 0xfb,
	0x53, 0x11, 0xfd, 0xdd, 0x6c, 0x3d, 0x39, 0xdf, 0x77, 0xbe, 0x7c, 0xe7, 0x9c, 0xa0, 0x79, 0xa2,
	0x0b, 0xa9, 0x60, 0x0a, 0x22, 0xe5, 0x02, 0xa6, 0x3b, 0x25, 0xb7, 0x90, 0x94, 0x0a, 0xf4, 0x54,
	0xe6, 0xf1, 0x86, 0x65, 0x3c, 0x61, 0x05, 0x6c, 0x64, 0xb9, 0xfd, 0x02, 0xea, 0x8d, 0x4c, 0x20,
	0x5e, 0x97, 0xc4, 0x80, 0x0a, 0x89, 0x1f, 0x3b, 0x26, 0x71, 0x4c, 0xf2, 0x57, 0x78, 0x74, 0xcb,
	0x4b, 0xec, 0x59, 0x56, 0x82, 0x76, 0xec, 0xe8, 0xce, 0xa5, 0x2e, 0x28, 0x25, 0x95, 0x4f, 0x4d,
	0x2e, 0x53, 0x39, 0x68, 0xcd, 0x52, 0xf0, 0xc9, 0x87, 0xcd, 0x64, 0xc1, 0xb8, 0xf8, 0x2c, 0x55,
	0xce, 0x0a, 0x2e, 0x85, 0x03, 0x3d, 0xf8, 0xde, 0x46, 0x68, 0xcd, 0x14, 0x33, 0x59, 0x50, 0x1a,
	0xbf, 0x44, 0x83, 0x52, 0xf0, 0xaf, 0x25, 0xc4, 0x3c, 0x09, 0x5b, 0xf7, 0x5b, 0x4f, 0x86, 0xb3,
	0x88, 0xf8, 0xdb, 0xfb, 0x4b, 0xe9, 0x42, 0x71, 0x91, 0x6e, 0x6c, 0x40, 0xfb, 0x0e, 0xbc, 0x48,
	0xf0, 0x23, 0x74, 0xfd, 0x48, 0x8c, 0x45, 0x99, 0x65, 0xe1, 0x8f, 0x9e, 0xa1, 0xf7, 0xe9, 0xa8,
	0x86, 0xac, 0xcc, 0x21, 0x7e, 0x85, 0x46, 0x7b, 0x67, 0x39, 0xde, 0x1a, 0xcf, 0x61, 0xfb, 0x9f,
	0x12, 0xc3, 0xfd, 0xa9, 0x45, 0xf8, 0x29, 0xba, 0x79, 0x4e, 0x77, 0x42, 0x3f, 0x9d, 0xd0, 0xf8,
	0x0c, 0x58, 0x69, 0xcd, 0xd1, 0x60, 0x67, 0x3c, 0x49, 0x61, 0xbd, 0x04, 0x95, 0xd0, 0xa4, 0x21,
	0xc4, 0x45, 0x01, 0x29, 0x28, 0x6f, 0xc6, 0xa1, 0x9d, 0x99, 0x23, 0xd3, 0x69, 0xfc, 0xf2, 0x66,
	0x6a, 0x88, 0x15, 0xb0, 0xbd, 0xeb, 0x53, 0xd0, 0x3b, 0x29, 0x34, 0xe0, 0x67, 0xa8, 0x53, 0x4d,
	0xa6, 0xd9, 0x35, 0x3f, 0x73, 0x37, 0xb5, 0xb7, 0xf6, 0x97, 0x3a, 0x20, 0xfe, 0x80, 0x6e, 0xd8,
	0x99, 0xc4, 0x67, 0x43, 0x31, 0xfd, 0x08, 0x0c, 0x99, 0x34, 0xc8, 0xcd, 0xd1, 0x2d, 0x4d, 0xbc,
	0x38, 0xc5, 0x74, 0x9c, 0x5f, 0x1e, 0x18, 0xeb, 0x3d, 0xbf, 0x0b, 0xc6, 0xb8, 0xad, 0x78, 0xef,
	0x8f, 0x8a, 0x6e, 0x53, 0x96, 0xee, 0x9f, 0xd6, 0x70, 0xfc, 0x0e, 0x05, 0x4a, 0x1e, 0xc2, 0x6b,
	0x15, 0xeb, 0x05, 0xf9, 0xcf, 0xc5, 0x25, 0x75, 0x1b, 0x08, 0x95, 0x07, 0x6a, 0x2b, 0x44, 0x77,
	0x51, 0x60, 0xbe, 0xf1, 0x6d, 0xd4, 0x35, 0x91, 0x9d, 0xc0, 0xb7, 0x95, 0x69, 0x4c, 0x87, 0x76,
	0x4c, 0xb8, 0x48, 0x5e, 0xbf, 0x47, 0x13, 0x2e, 0x1b, 0xe5, 0x4f, 0x2f, 0xea, 0xe3, 0xec, 0xea,
	0x6f, 0xed, 0x53, 0xb7, 0xda, 0xe8, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x36, 0x9d,
	0xec, 0xa8, 0x03, 0x00, 0x00,
}