// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderSurchInfoTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderSurchInfoTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderSurchInfoTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderSurchInfoTypes_Ad

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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	InformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	FieldTypeId       *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	InformationType   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderSurchInfoTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderSurchInfoTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderSurchInfoTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x51, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0xa9, 0xbd, 0xbb, 0xf7, 0x92, 0x21, 0x9b, 0x1d, 0x4a, 0xed, 0x40, 0xc6, 0x7c, 0x11,
	0x1f, 0x32, 0x51, 0x84, 0x81, 0x2f, 0x2a, 0xa8, 0x4c, 0x5d, 0x95, 0x38, 0x04, 0x7d, 0x29, 0x75,
	0x3d, 0xab, 0x81, 0x36, 0x19, 0x49, 0xea, 0xf0, 0x5b, 0xa8, 0xc3, 0x0f, 0xe9, 0x07, 0xf0, 0xdd,
	0xb4, 0xc9, 0xdc, 0xda, 0x8d, 0xcb, 0xd8, 0x4b, 0xcb, 0xe9, 0x39, 0xff, 0x5f, 0xce, 0xf9, 0x9f,
	0x06, 0x3d, 0x49, 0xa4, 0xe2, 0x02, 0x46, 0xc0, 0x52, 0xca, 0x60, 0xb4, 0x14, 0x7c, 0x0e, 0x49,
	0x21, 0x40, 0x8e, 0x78, 0x1e, 0xbd, 0x02, 0xf5, 0x4e, 0x24, 0x20, 0x3e, 0x14, 0x62, 0xfe, 0x75,
	0xc2, 0x16, 0x7c, 0xf6, 0x7d, 0x09, 0x32, 0x7a, 0x96, 0x60, 0x5d, 0xa7, 0xb8, 0x77, 0xdf, 0x88,
	0xb1, 0x11, 0xe3, 0xab, 0x14, 0x41, 0xcf, 0x1e, 0xf4, 0x2d, 0xce, 0x0a, 0x90, 0x06, 0x10, 0xdc,
	0xae, 0x9f, 0x0e, 0x42, 0x70, 0x61, 0x53, 0xfd, 0x7a, 0x2a, 0x07, 0x29, 0xe3, 0x14, 0x6c, 0xf2,
	0x6e, 0x33, 0xa9, 0x62, 0xaa, 0x0f, 0x13, 0x79, 0xac, 0x28, 0x67, 0xa6, 0x68, 0xb8, 0x76, 0x10,
	0x7a, 0x1f, 0x8b, 0x58, 0x67, 0x41, 0x48, 0xef, 0x0d, 0xea, 0xed, 0xd4, 0x44, 0x4a, 0x37, 0x16,
	0xd1, 0xc4, 0x77, 0x06, 0xce, 0xbd, 0xf6, 0xc3, 0x3e, 0xb6, 0xa3, 0xd8, 0xf6, 0x28, 0x53, 0x90,
	0x82, 0xf8, 0x58, 0x46, 0xe4, 0xc6, 0x8e, 0xae, 0x9c, 0x67, 0x92, 0x78, 0x63, 0xe4, 0x1f, 0x80,
	0x45, 0xac, 0xc8, 0x32, 0xff, 0xcf, 0x85, 0x46, 0x5e, 0x92, 0x9b, 0x7b, 0xaa, 0x50, 0x67, 0x87,
	0xbf, 0xcf, 0xd0, 0x25, 0x01, 0xb9, 0xe4, 0x4c, 0x82, 0xf7, 0x00, 0xb5, 0xaa, 0x99, 0x6d, 0x17,
	0x01, 0xae, 0x1b, 0x6a, 0xfc, 0x78, 0x51, 0x3e, 0x89, 0x29, 0xf4, 0x3e, 0xa1, 0x6e, 0x39, 0x6d,
	0xb4, 0x03, 0xf7, 0xaf, 0x0d, 0x5c, 0x2d, 0xc6, 0x0d, 0x71, 0xd3, 0x94, 0xa9, 0x8e, 0x27, 0xdb,
	0x98, 0x74, 0xf2, 0xfa, 0x07, 0x3d, 0xd3, 0x85, 0x75, 0xd9, 0x77, 0x2b, 0xe2, 0x9d, 0x3d, 0xa2,
	0xd9, 0xc1, 0xd4, 0xbc, 0xc9, 0xa6, 0xdc, 0x7b, 0x8d, 0x5c, 0xc1, 0x57, 0xfe, 0x59, 0xa5, 0x1a,
	0xe3, 0xe3, 0xff, 0x0a, 0xbc, 0x71, 0x02, 0x13, 0xbe, 0x22, 0x25, 0x24, 0xf8, 0xeb, 0x20, 0x57,
	0x07, 0xde, 0x2d, 0x74, 0xae, 0xc3, 0x72, 0x43, 0x3f, 0x42, 0x6d, 0x4e, 0x8b, 0xb4, 0x74, 0xa8,
	0x9d, 0x7f, 0x7b, 0x78, 0x8d, 0x3f, 0xc3, 0x93, 0xf6, 0xf8, 0x14, 0x5d, 0x5f, 0x50, 0xc8, 0x92,
	0xff, 0x9c, 0x5f, 0x47, 0x70, 0xda, 0x95, 0xc4, 0x12, 0x5e, 0xa2, 0x6e, 0xb3, 0x1f, 0x7f, 0x1d,
	0xd6, 0xd7, 0x69, 0x21, 0x52, 0x09, 0xca, 0x52, 0xc3, 0xe8, 0x34, 0x7a, 0x79, 0x3e, 0x43, 0x7d,
	0xca, 0x1b, 0xd6, 0x6d, 0x6f, 0xe3, 0xe7, 0xc7, 0x27, 0xdd, 0xd3, 0x2f, 0xe7, 0xd5, 0x55, 0x78,
	0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x13, 0xc2, 0x71, 0xd7, 0xe7, 0x03, 0x00, 0x00,
}
