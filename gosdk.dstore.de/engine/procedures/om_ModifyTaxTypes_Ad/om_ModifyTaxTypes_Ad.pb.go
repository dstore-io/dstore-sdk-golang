// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyTaxTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyTaxTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyTaxTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyTaxTypes_Ad

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
	TaxTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tax_type_id,json=taxTypeId" json:"tax_type_id,omitempty"`
	TaxTypeIdNull bool                        `protobuf:"varint,1001,opt,name=tax_type_id_null,json=taxTypeIdNull" json:"tax_type_id_null,omitempty"`
	TaxType       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=tax_type,json=taxType" json:"tax_type,omitempty"`
	TaxTypeNull   bool                        `protobuf:"varint,1002,opt,name=tax_type_null,json=taxTypeNull" json:"tax_type_null,omitempty"`
	Delete        *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete" json:"delete,omitempty"`
	DeleteNull    bool                        `protobuf:"varint,1003,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTaxTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TaxTypeId
	}
	return nil
}

func (m *Parameters) GetTaxType() *dstore_values.StringValue {
	if m != nil {
		return m.TaxType
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyTaxTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyTaxTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyTaxTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyTaxTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x13, 0xf2, 0xc3, 0x44, 0x15, 0x95, 0x91, 0xd0, 0x92, 0x08, 0x54, 0xa5, 0x97, 0x9e,
	0x9c, 0x8a, 0x08, 0x84, 0xc4, 0x89, 0x4a, 0x1c, 0x72, 0x48, 0x85, 0xac, 0x82, 0x04, 0x97, 0x95,
	0x8b, 0xa7, 0x2b, 0x8b, 0x8d, 0x1d, 0xd9, 0x0e, 0x6d, 0xdf, 0x82, 0xa7, 0xe0, 0x09, 0x78, 0x29,
	0xe0, 0x25, 0xf0, 0xee, 0x38, 0x69, 0xb3, 0x44, 0xa2, 0x97, 0x5d, 0x8f, 0xe7, 0xfb, 0xbe, 0x19,
	0xfb, 0x1b, 0xc3, 0x54, 0xf9, 0x60, 0x1d, 0x4e, 0xd0, 0x14, 0xda, 0xe0, 0x64, 0xe9, 0xec, 0x17,
	0x54, 0x2b, 0x87, 0x7e, 0x62, 0x17, 0xf9, 0xdc, 0x2a, 0x7d, 0x79, 0x73, 0x2e, 0xaf, 0xcf, 0x6f,
	0x96, 0xe8, 0xf3, 0xb7, 0x8a, 0xc7, 0x7c, 0xb0, 0x6c, 0x4c, 0x24, 0x4e, 0x24, 0xbe, 0x0b, 0x39,
	0x7c, 0x9c, 0x84, 0xbf, 0xc9, 0x72, 0x85, 0x9e, 0x88, 0xc3, 0xa7, 0xdb, 0xd5, 0xd0, 0x39, 0xeb,
	0x52, 0x6a, 0xb4, 0x9d, 0x5a, 0xa0, 0xf7, 0xb2, 0xc0, 0x94, 0x3c, 0x6a, 0x26, 0x83, 0xd4, 0xe6,
	0xd2, 0xba, 0x85, 0x0c, 0xda, 0x1a, 0x02, 0x8d, 0x7f, 0xb4, 0x00, 0xde, 0x4b, 0x27, 0x63, 0x16,
	0x9d, 0x67, 0x6f, 0x60, 0x10, 0xe4, 0x75, 0x1e, 0x62, 0x43, 0xb9, 0x56, 0xd9, 0xde, 0xe1, 0xde,
	0xf1, 0xe0, 0xc5, 0x88, 0xa7, 0xd6, 0x53, 0x5b, 0xda, 0x04, 0x2c, 0xd0, 0x7d, 0xac, 0x22, 0xf1,
	0x30, 0x50, 0xff, 0x33, 0xc5, 0x8e, 0xe1, 0xe0, 0x0e, 0x39, 0x37, 0xab, 0xb2, 0xcc, 0x7e, 0xf5,
	0xa2, 0x44, 0x5f, 0xec, 0x6f, 0x50, 0x67, 0x71, 0x97, 0xbd, 0x84, 0xfe, 0x1a, 0x99, 0xb5, 0xea,
	0x1a, 0xc3, 0x46, 0x0d, 0x1f, 0x9c, 0x36, 0x05, 0x95, 0xe8, 0x25, 0x32, 0x3b, 0x82, 0xfd, 0x4d,
	0x81, 0x5a, 0xfd, 0x37, 0xa9, 0x0f, 0x12, 0xa0, 0xd6, 0x9e, 0x42, 0x57, 0x61, 0x19, 0x8f, 0x93,
	0xb5, 0x77, 0x76, 0x7f, 0x61, 0x6d, 0x89, 0xd2, 0x90, 0x74, 0x82, 0xb2, 0x43, 0x18, 0xd0, 0x8a,
	0x74, 0xff, 0x90, 0x2e, 0xd0, 0x5e, 0x25, 0x3b, 0xfe, 0xd9, 0x82, 0xbe, 0x40, 0xbf, 0xb4, 0xc6,
	0x23, 0x3b, 0x81, 0x4e, 0x6d, 0x43, 0xba, 0xa0, 0x4d, 0xf3, 0xc9, 0x5b, 0xb2, 0xe8, 0x5d, 0xf5,
	0x15, 0x04, 0x64, 0x9f, 0xe0, 0xa0, 0x32, 0x20, 0xbf, 0xe3, 0x40, 0x3c, 0x79, 0x3b, 0x92, 0x79,
	0x83, 0xdc, 0xf4, 0x69, 0x1e, 0xe3, 0xd9, 0x6d, 0x2c, 0x1e, 0x2d, 0xb6, 0x37, 0xd8, 0x6b, 0xe8,
	0x25, 0xe3, 0xe3, 0x89, 0x2b, 0xc5, 0xe7, 0xff, 0x28, 0xd2, 0x58, 0xcc, 0xe9, 0x2f, 0xd6, 0x70,
	0x76, 0x0a, 0x6d, 0x67, 0xaf, 0xb2, 0x07, 0x35, 0xeb, 0x84, 0xff, 0x7f, 0x40, 0xf9, 0xfa, 0x06,
	0xb8, 0xb0, 0x57, 0xa2, 0x22, 0x0f, 0x9f, 0x41, 0x3b, 0xae, 0xd9, 0x13, 0xe8, 0xc6, 0xa8, 0x9a,
	0x99, 0xef, 0x67, 0xf1, 0x4e, 0x3a, 0xa2, 0x13, 0xc3, 0x99, 0x3a, 0xfd, 0x00, 0x23, 0x6d, 0x1b,
	0xca, 0xb7, 0xef, 0xe5, 0xf3, 0xab, 0xc2, 0x7a, 0xf5, 0x75, 0x9d, 0x57, 0xf7, 0x7d, 0x52, 0x17,
	0xdd, 0x7a, 0x7a, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x1c, 0xf5, 0xe9, 0x8a, 0x03,
	0x00, 0x00,
}