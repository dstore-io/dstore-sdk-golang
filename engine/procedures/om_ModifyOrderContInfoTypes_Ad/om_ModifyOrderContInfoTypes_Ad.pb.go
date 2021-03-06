// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyOrderContInfoTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyOrderContInfoTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyOrderContInfoTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyOrderContInfoTypes_Ad

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
	InformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull bool                        `protobuf:"varint,1001,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	InformationType       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	InformationTypeNull   bool                        `protobuf:"varint,1002,opt,name=information_type_null,json=informationTypeNull" json:"information_type_null,omitempty"`
	Delete                *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete" json:"delete,omitempty"`
	DeleteNull            bool                        `protobuf:"varint,1003,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
	FieldTypeId           *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	FieldTypeIdNull       bool                        `protobuf:"varint,1004,opt,name=field_type_id_null,json=fieldTypeIdNull" json:"field_type_id_null,omitempty"`
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

func (m *Parameters) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Parameters) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyOrderContInfoTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyOrderContInfoTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyOrderContInfoTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyOrderContInfoTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x5d, 0x93, 0x96, 0x13, 0x24, 0x75, 0x42, 0x65, 0x4d, 0x50, 0x4a, 0xbd, 0xf1, 0xa2,
	0x6c, 0xc4, 0xdc, 0x88, 0x20, 0xfe, 0xd1, 0x8b, 0xa0, 0xa9, 0xb2, 0x88, 0xa0, 0x37, 0xcb, 0xd6,
	0x39, 0x59, 0x06, 0x37, 0x33, 0x61, 0x66, 0x63, 0xe9, 0x5b, 0xf8, 0x44, 0xbe, 0x85, 0x0f, 0xe1,
	0xcf, 0x43, 0x78, 0x66, 0xcf, 0xa4, 0x4d, 0x36, 0x45, 0xe9, 0x4d, 0x32, 0x67, 0xce, 0xf9, 0xbe,
	0xef, 0xcc, 0x37, 0x67, 0x16, 0x9e, 0x4a, 0x57, 0x19, 0x8b, 0x43, 0xd4, 0x85, 0xd2, 0x38, 0x9c,
	0x5b, 0xf3, 0x19, 0xe5, 0xc2, 0xa2, 0x1b, 0x9a, 0x59, 0x36, 0x31, 0x52, 0x4d, 0xcf, 0xdf, 0x5a,
	0x89, 0xf6, 0x95, 0xd1, 0xd5, 0x58, 0x4f, 0xcd, 0xfb, 0xf3, 0x39, 0xba, 0xec, 0x85, 0x4c, 0xa8,
	0xb2, 0x32, 0xe2, 0x88, 0xe1, 0x09, 0xc3, 0x93, 0x7f, 0x63, 0xfa, 0xbd, 0x20, 0xf6, 0x35, 0x2f,
	0x17, 0xe8, 0x98, 0xa2, 0x7f, 0x67, 0xbd, 0x03, 0xb4, 0xd6, 0xd8, 0x90, 0x1a, 0xac, 0xa7, 0x66,
	0xe8, 0x5c, 0x5e, 0x60, 0x48, 0xde, 0x6f, 0x26, 0xab, 0x5c, 0x91, 0x98, 0x9d, 0xe5, 0x95, 0x32,
	0x9a, 0x8b, 0x0e, 0x7f, 0x44, 0x00, 0xef, 0x72, 0x9b, 0x53, 0x16, 0xad, 0x13, 0xaf, 0xa1, 0xb7,
	0x52, 0x93, 0x55, 0xd4, 0x58, 0xa6, 0x64, 0xbc, 0x75, 0xb0, 0xf5, 0xa0, 0xf3, 0x68, 0x90, 0x84,
	0xc3, 0x84, 0xf6, 0x94, 0xae, 0xb0, 0x40, 0xfb, 0xc1, 0x47, 0xe9, 0xad, 0x15, 0x9c, 0x3f, 0xcf,
	0x58, 0x8a, 0xc7, 0x10, 0x5f, 0x41, 0x96, 0xe9, 0x45, 0x59, 0xc6, 0x3f, 0x77, 0x88, 0x72, 0x37,
	0xdd, 0xdf, 0x40, 0x9d, 0x50, 0x56, 0x1c, 0xc3, 0x5e, 0x13, 0x19, 0x6f, 0xd7, 0x3d, 0xf4, 0x1b,
	0x3d, 0xb8, 0xca, 0x2a, 0x5d, 0x70, 0x0b, 0xdd, 0x06, 0x99, 0x18, 0xc1, 0xfe, 0x46, 0x03, 0xb5,
	0xfa, 0x2f, 0x56, 0xef, 0x35, 0x00, 0xb5, 0xf6, 0x08, 0xda, 0x12, 0x4b, 0xb2, 0x23, 0x8e, 0xae,
	0x3c, 0xf5, 0xa9, 0x31, 0x25, 0xe6, 0x9a, 0x25, 0x43, 0xa9, 0x38, 0x80, 0x0e, 0xaf, 0x98, 0xff,
	0x37, 0xf3, 0x03, 0xef, 0xd5, 0xb4, 0xcf, 0xe0, 0xe6, 0x54, 0x61, 0x29, 0x2f, 0x3c, 0xbd, 0xf1,
	0x7f, 0x4f, 0x3b, 0x35, 0x22, 0xb8, 0x79, 0x04, 0x62, 0x8d, 0x80, 0x95, 0xfe, 0xb0, 0x52, 0x77,
	0xa5, 0xd2, 0xcb, 0x1d, 0x7e, 0xdf, 0x86, 0xdd, 0x14, 0xdd, 0xdc, 0x68, 0x87, 0xe2, 0x21, 0xb4,
	0xea, 0xa9, 0x09, 0xf7, 0x78, 0xe1, 0x61, 0x18, 0x4a, 0x9e, 0xa8, 0x63, 0xff, 0x9b, 0x72, 0xa1,
	0xf8, 0x08, 0x7b, 0x7e, 0x5e, 0xb2, 0x15, 0x83, 0xe8, 0x02, 0x22, 0x02, 0x27, 0x0d, 0x70, 0x73,
	0xac, 0x26, 0x14, 0x8f, 0x2f, 0xe3, 0xb4, 0x3b, 0x5b, 0xdf, 0xa0, 0xa9, 0xd8, 0x09, 0x73, 0x4a,
	0x06, 0x7b, 0xc6, 0x7b, 0x1b, 0x8c, 0x3c, 0xc5, 0x13, 0xfe, 0x4f, 0x97, 0xe5, 0xe2, 0x0d, 0x44,
	0xd6, 0x9c, 0x91, 0x71, 0x1e, 0xf5, 0x24, 0xb9, 0xce, 0xcb, 0x4a, 0x96, 0x5e, 0x24, 0xa9, 0x39,
	0x4b, 0x3d, 0x4d, 0xff, 0x2e, 0x44, 0xb4, 0x16, 0xb7, 0xa1, 0x4d, 0x91, 0xbf, 0x90, 0x6f, 0x27,
	0xe4, 0x4e, 0x2b, 0x6d, 0x51, 0x38, 0x96, 0x2f, 0x33, 0x18, 0x28, 0xd3, 0xd0, 0xb8, 0x7c, 0xfc,
	0x9f, 0x9e, 0x17, 0xc6, 0xc9, 0x2f, 0xcb, 0xbc, 0xbc, 0xfe, 0xf7, 0xe1, 0xb4, 0x5d, 0x3f, 0xc0,
	0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x91, 0x9b, 0x0b, 0x61, 0x04, 0x00, 0x00,
}
