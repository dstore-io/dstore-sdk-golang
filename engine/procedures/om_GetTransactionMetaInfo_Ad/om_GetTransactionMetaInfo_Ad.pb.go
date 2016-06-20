// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetTransactionMetaInfo_Ad.proto
// DO NOT EDIT!

/*
Package om_GetTransactionMetaInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetTransactionMetaInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetTransactionMetaInfo_Ad

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
	TransactionId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	TransactionIdNull bool                        `protobuf:"varint,1001,opt,name=transaction_id_null,json=transactionIdNull" json:"transaction_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTransactionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TransactionId
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
	RowId                      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TransactionByUserId        *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=transaction_by_user_id,json=transactionByUserId" json:"transaction_by_user_id,omitempty"`
	TransactionByProcedureName *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=transaction_by_procedure_name,json=transactionByProcedureName" json:"transaction_by_procedure_name,omitempty"`
	InternalTransactionComment *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=internal_transaction_comment,json=internalTransactionComment" json:"internal_transaction_comment,omitempty"`
	TransactionId              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTransactionByUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TransactionByUserId
	}
	return nil
}

func (m *Response_Row) GetTransactionByProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.TransactionByProcedureName
	}
	return nil
}

func (m *Response_Row) GetInternalTransactionComment() *dstore_values.StringValue {
	if m != nil {
		return m.InternalTransactionComment
	}
	return nil
}

func (m *Response_Row) GetTransactionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetTransactionMetaInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetTransactionMetaInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetTransactionMetaInfo_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetTransactionMetaInfo_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x65, 0xbb, 0x6e, 0x5b, 0xae, 0xf8, 0x35, 0x85, 0x12, 0xb3, 0x2a, 0xa5, 0xbe, 0x88, 0x0f,
	0xb3, 0xa2, 0x2f, 0x05, 0x41, 0x70, 0x8b, 0x48, 0x85, 0x86, 0x3a, 0xa8, 0xa0, 0x50, 0xc3, 0x74,
	0x73, 0x0d, 0xc1, 0xcd, 0x4c, 0x99, 0x99, 0xb5, 0xf4, 0xd1, 0x7f, 0xe0, 0x47, 0xff, 0xa4, 0xef,
	0xfe, 0x00, 0x67, 0x32, 0x13, 0x9b, 0x49, 0x41, 0xbb, 0x2f, 0x09, 0x37, 0xe7, 0x9e, 0x73, 0xe7,
	0x9e, 0x7b, 0x27, 0xf0, 0xb4, 0xd0, 0x46, 0x2a, 0x9c, 0xa0, 0x28, 0x2b, 0x81, 0x93, 0x63, 0x25,
	0x67, 0x58, 0x2c, 0x14, 0xea, 0x89, 0xac, 0xf3, 0x97, 0x68, 0xde, 0x28, 0x2e, 0x34, 0x9f, 0x99,
	0x4a, 0x8a, 0x7d, 0x34, 0x7c, 0x4f, 0x7c, 0x92, 0xf9, 0xf3, 0x82, 0xda, 0x3c, 0x23, 0xc9, 0x43,
	0x4f, 0xa6, 0x9e, 0x4c, 0xff, 0xc5, 0x48, 0x37, 0x42, 0xa1, 0x2f, 0x7c, 0xbe, 0x40, 0xed, 0x05,
	0xd2, 0xdb, 0x71, 0x75, 0x54, 0x4a, 0xaa, 0x00, 0x8d, 0x63, 0xa8, 0x46, 0xad, 0x79, 0x89, 0x01,
	0xbc, 0xdf, 0x07, 0x0d, 0xaf, 0x6c, 0x1d, 0x55, 0x73, 0x57, 0xd6, 0x27, 0x6d, 0x7f, 0x1d, 0x00,
	0x1c, 0x70, 0xc5, 0x2d, 0x8a, 0x4a, 0x93, 0x29, 0x5c, 0x37, 0xe7, 0x47, 0xcb, 0xab, 0x22, 0x19,
	0x6c, 0x0d, 0x1e, 0x5c, 0x7d, 0x3c, 0xa6, 0xa1, 0x8b, 0x70, 0xb2, 0x4a, 0x18, 0x2c, 0x51, 0xbd,
	0x73, 0x11, 0xbb, 0xd6, 0xa1, 0xec, 0x15, 0x64, 0x02, 0x1b, 0xb1, 0x46, 0x2e, 0x16, 0xf3, 0x79,
	0xf2, 0x6b, 0xcd, 0x2a, 0xad, 0xb3, 0x5b, 0x51, 0x72, 0x66, 0x91, 0xed, 0xb3, 0x11, 0xac, 0x33,
	0xd4, 0xc7, 0x52, 0x68, 0x24, 0x8f, 0x60, 0xd4, 0x74, 0x18, 0x0a, 0xa7, 0x34, 0xb6, 0xcf, 0x77,
	0xff, 0xc2, 0x3d, 0x99, 0x4f, 0x24, 0xef, 0xe1, 0xa6, 0xeb, 0x2d, 0xef, 0x34, 0x97, 0xac, 0x6c,
	0x0d, 0x2d, 0x99, 0xf6, 0xc8, 0x7d, 0x0b, 0x5a, 0xeb, 0x43, 0xcc, 0x6e, 0xd4, 0xf1, 0x07, 0xb2,
	0x03, 0x6b, 0xc1, 0xd3, 0x64, 0xd8, 0x28, 0xde, 0xbb, 0xa0, 0xe8, 0x1d, 0xdf, 0xf7, 0x6f, 0xd6,
	0xa6, 0x93, 0x57, 0x30, 0x54, 0xf2, 0x24, 0xb9, 0xd2, 0xb0, 0x76, 0xe8, 0xe5, 0x77, 0x80, 0xb6,
	0x4e, 0x50, 0x26, 0x4f, 0x98, 0x13, 0x49, 0x7f, 0xaf, 0xc0, 0xd0, 0x06, 0x64, 0x13, 0x56, 0x6d,
	0xe8, 0x86, 0xf2, 0x2d, 0xb3, 0xe6, 0x8c, 0xd8, 0xc8, 0x86, 0xd6, 0xf0, 0xd7, 0xb0, 0xd9, 0x35,
	0xfc, 0xe8, 0x34, 0x5f, 0x68, 0x54, 0x2e, 0xef, 0x7b, 0xf6, 0xff, 0xe9, 0x75, 0x87, 0x35, 0x3d,
	0x7d, 0x6b, 0x99, 0x56, 0xf2, 0x23, 0xdc, 0xed, 0x49, 0xfe, 0x5d, 0xfa, 0x5c, 0xd8, 0x65, 0x49,
	0x7e, 0x64, 0xf1, 0x78, 0x82, 0xb2, 0x36, 0xaa, 0x12, 0xa5, 0x17, 0x4e, 0x23, 0xe1, 0x83, 0x96,
	0x9f, 0x59, 0x3a, 0x39, 0x84, 0x3b, 0xee, 0x10, 0x4a, 0xf0, 0x79, 0xde, 0x2d, 0x34, 0x93, 0x75,
	0x8d, 0xc2, 0x24, 0x3f, 0x2f, 0x21, 0xdf, 0x0a, 0x74, 0x7c, 0xdc, 0xf5, 0x74, 0xb2, 0x7b, 0x61,
	0x8d, 0xcf, 0xb2, 0x65, 0xf7, 0x78, 0x7a, 0x08, 0xe3, 0x4a, 0xf6, 0x26, 0x77, 0x7e, 0xf5, 0x3f,
	0x3c, 0x2b, 0xa5, 0x2e, 0x3e, 0xb7, 0x78, 0xb1, 0xec, 0xdf, 0xe1, 0x68, 0xb5, 0xb9, 0x80, 0x4f,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xec, 0xb9, 0x79, 0xe3, 0x5d, 0x04, 0x00, 0x00,
}