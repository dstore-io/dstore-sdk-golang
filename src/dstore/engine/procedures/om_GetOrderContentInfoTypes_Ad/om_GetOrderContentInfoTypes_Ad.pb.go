// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderContentInfoTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderContentInfoTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderContentInfoTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderContentInfoTypes_Ad

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
	InformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull bool                        `protobuf:"varint,1001,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	RowId                         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	OrderContentInformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=order_content_information_type_id,json=orderContentInformationTypeId" json:"order_content_information_type_id,omitempty"`
	FieldTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	InformationType               *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetOrderContentInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderContentInformationTypeId
	}
	return nil
}

func (m *Response_Row) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
	}
	return nil
}

func (m *Response_Row) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderContentInfoTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderContentInfoTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderContentInfoTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x5d, 0xcb, 0xd3, 0x30,
	0x14, 0xc7, 0xd9, 0xd3, 0x67, 0x2f, 0x64, 0xc8, 0x66, 0x87, 0x52, 0x3b, 0x94, 0x39, 0x6f, 0xbc,
	0x90, 0x4c, 0xf4, 0x42, 0x11, 0x04, 0x5f, 0x50, 0x19, 0xba, 0x2a, 0x41, 0x06, 0x7a, 0x13, 0xea,
	0x7a, 0x56, 0x02, 0x6d, 0x32, 0x92, 0xcc, 0xe1, 0xa7, 0xf0, 0x65, 0x9f, 0xc4, 0x8f, 0xe5, 0xb7,
	0x30, 0x6d, 0x32, 0xb7, 0x76, 0xc3, 0xb7, 0x9b, 0x8d, 0xd3, 0x73, 0xfe, 0xbf, 0x9c, 0xf3, 0x3f,
	0x09, 0x7a, 0x98, 0x28, 0x2d, 0x24, 0x4c, 0x80, 0xa7, 0x8c, 0xc3, 0x64, 0x25, 0xc5, 0x02, 0x92,
	0xb5, 0x04, 0x35, 0x11, 0x39, 0x7d, 0x01, 0xfa, 0xb5, 0x4c, 0x40, 0x3e, 0x15, 0x5c, 0x03, 0xd7,
	0x53, 0xbe, 0x14, 0x6f, 0x3f, 0xad, 0x40, 0xd1, 0xc7, 0x09, 0x36, 0x95, 0x5a, 0xf8, 0xb7, 0xac,
	0x1c, 0x5b, 0x39, 0xfe, 0xbd, 0x26, 0x1c, 0xb8, 0xc3, 0x3e, 0xc6, 0xd9, 0x1a, 0x94, 0x45, 0x84,
	0x57, 0xaa, 0x1d, 0x80, 0x94, 0x42, 0xba, 0xd4, 0xb0, 0x9a, 0xca, 0x41, 0xa9, 0x38, 0x05, 0x97,
	0xbc, 0x51, 0x4f, 0xea, 0x98, 0x99, 0xc3, 0x64, 0x1e, 0x6b, 0x26, 0xb8, 0x2d, 0x1a, 0x6f, 0x1b,
	0x08, 0xbd, 0x89, 0x65, 0x6c, 0xb2, 0x20, 0x95, 0xff, 0x12, 0x0d, 0x0e, 0x6a, 0xa8, 0x36, 0x8d,
	0x51, 0x96, 0x04, 0x8d, 0x51, 0xe3, 0x66, 0xf7, 0xce, 0x10, 0xbb, 0x61, 0x5c, 0x7b, 0xcc, 0x0c,
	0x90, 0x82, 0x9c, 0x17, 0x11, 0xb9, 0x78, 0xa0, 0x2b, 0xe6, 0x99, 0x26, 0xfe, 0x7d, 0x14, 0x9c,
	0x80, 0x51, 0xbe, 0xce, 0xb2, 0xe0, 0x47, 0xdb, 0x20, 0x3b, 0xe4, 0xd2, 0x91, 0x2a, 0x32, 0xd9,
	0xf1, 0xf7, 0x73, 0xd4, 0x21, 0xa0, 0x56, 0x82, 0x2b, 0xf0, 0x6f, 0xa3, 0x66, 0x39, 0xb3, 0xeb,
	0x22, 0xc4, 0x55, 0x4b, 0xad, 0x1f, 0xcf, 0x8a, 0x5f, 0x62, 0x0b, 0xfd, 0x77, 0xa8, 0x5f, 0x4c,
	0x4b, 0x0f, 0xe0, 0xc1, 0xd9, 0xc8, 0x33, 0x62, 0x5c, 0x13, 0xd7, 0x4d, 0x99, 0x99, 0x78, 0xba,
	0x8f, 0x49, 0x2f, 0xaf, 0x7e, 0x30, 0x33, 0xb5, 0x9d, 0xcb, 0x81, 0x57, 0x12, 0xaf, 0x1d, 0x11,
	0xed, 0x0e, 0x66, 0xf6, 0x9f, 0xec, 0xca, 0xfd, 0x57, 0xc8, 0x93, 0x62, 0x13, 0x9c, 0x97, 0xaa,
	0x07, 0xf8, 0x5f, 0xee, 0x05, 0xde, 0x79, 0x81, 0x89, 0xd8, 0x90, 0x02, 0x13, 0x7e, 0x3e, 0x43,
	0x9e, 0x09, 0xfc, 0xcb, 0xa8, 0x65, 0xc2, 0x62, 0x47, 0x5f, 0x22, 0x63, 0x4f, 0x93, 0x34, 0x4d,
	0x68, 0xbc, 0x5f, 0xa2, 0xeb, 0xa2, 0xa0, 0xd1, 0x85, 0xc5, 0xd1, 0x53, 0x6b, 0xfd, 0x1a, 0xfd,
	0x79, 0xaf, 0x57, 0x45, 0xad, 0xa9, 0xea, 0x8e, 0x1f, 0xa1, 0x0b, 0x4b, 0x06, 0x59, 0xf2, 0x8b,
	0xf9, 0xed, 0x2f, 0x98, 0xdd, 0x52, 0xe2, 0x08, 0xcf, 0x51, 0xbf, 0xde, 0x5b, 0xb0, 0x8d, 0xaa,
	0xab, 0x76, 0x10, 0xa5, 0x25, 0xe3, 0xa9, 0x65, 0xf4, 0x6a, 0x37, 0xe7, 0xc9, 0x1c, 0x0d, 0x99,
	0xa8, 0xd9, 0xba, 0x7f, 0xad, 0xef, 0xef, 0xfd, 0xe7, 0x3b, 0xfe, 0xd0, 0x2a, 0x1f, 0xca, 0xdd,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0xf9, 0x6f, 0x5d, 0x09, 0x04, 0x00, 0x00,
}