// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyOrderSurchInfTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyOrderSurchInfTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyOrderSurchInfTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyOrderSurchInfTypes_Ad

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
	InformationType       *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	InformationTypeNull   bool                        `protobuf:"varint,1002,opt,name=information_type_null,json=informationTypeNull" json:"information_type_null,omitempty"`
	FieldTypeId           *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	FieldTypeIdNull       bool                        `protobuf:"varint,1003,opt,name=field_type_id_null,json=fieldTypeIdNull" json:"field_type_id_null,omitempty"`
	Delete                *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull            bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
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

func (m *Parameters) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyOrderSurchInfTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyOrderSurchInfTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyOrderSurchInfTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x5d, 0x93, 0x96, 0x13, 0x24, 0x75, 0x42, 0x65, 0x4d, 0x50, 0x4a, 0xbd, 0xf1, 0xa2,
	0x6c, 0xc4, 0x5c, 0x28, 0x82, 0x88, 0x42, 0x2f, 0x82, 0xa6, 0xca, 0x28, 0x05, 0xbd, 0x59, 0xb6,
	0x9d, 0x93, 0x38, 0xb0, 0x99, 0x09, 0x33, 0x1b, 0x4b, 0xdf, 0xc2, 0x27, 0xf2, 0x2d, 0x7c, 0x08,
	0x7f, 0x1e, 0xc2, 0x99, 0x3d, 0xd3, 0x36, 0x3b, 0x29, 0x8a, 0xde, 0x24, 0x7b, 0xf6, 0x9c, 0xef,
	0xfb, 0xce, 0x39, 0xf3, 0xcd, 0xc2, 0x33, 0x61, 0x2b, 0x6d, 0x70, 0x88, 0x6a, 0x26, 0x15, 0x0e,
	0x17, 0x46, 0x9f, 0xa2, 0x58, 0x1a, 0xb4, 0x43, 0x3d, 0xcf, 0x27, 0x5a, 0xc8, 0xe9, 0xf9, 0x1b,
	0x23, 0xd0, 0xbc, 0x5b, 0x9a, 0xd3, 0x4f, 0x63, 0x35, 0x7d, 0x7f, 0xbe, 0x40, 0x9b, 0xbf, 0x10,
	0x99, 0xab, 0xac, 0x34, 0x3b, 0x20, 0x78, 0x46, 0xf0, 0xec, 0xcf, 0x98, 0x7e, 0x2f, 0x88, 0x7d,
	0x2e, 0xca, 0x25, 0x5a, 0xa2, 0xe8, 0xdf, 0x69, 0x76, 0x80, 0xc6, 0x68, 0x13, 0x52, 0x83, 0x66,
	0x6a, 0x8e, 0xd6, 0x16, 0x33, 0x0c, 0xc9, 0xfb, 0x71, 0xb2, 0x2a, 0xa4, 0x9a, 0x6a, 0x33, 0x2f,
	0x2a, 0xa9, 0x15, 0x15, 0xed, 0x7f, 0x4b, 0x00, 0xde, 0x16, 0xa6, 0x70, 0x59, 0x34, 0x96, 0xbd,
	0x82, 0xde, 0x4a, 0x4d, 0x5e, 0xb9, 0xc6, 0x72, 0x29, 0xd2, 0x8d, 0xbd, 0x8d, 0x07, 0x9d, 0x47,
	0x83, 0x2c, 0x0c, 0x13, 0xda, 0x93, 0xaa, 0xc2, 0x19, 0x9a, 0x63, 0x1f, 0xf1, 0x5b, 0x2b, 0x38,
	0x3f, 0xcf, 0x58, 0xb0, 0x27, 0x90, 0x5e, 0x43, 0x96, 0xab, 0x65, 0x59, 0xa6, 0xdf, 0xb7, 0x1c,
	0xe5, 0x36, 0xdf, 0x5d, 0x43, 0x1d, 0xb9, 0x2c, 0x3b, 0x84, 0x9d, 0x18, 0x99, 0x6e, 0xd6, 0x3d,
	0xf4, 0xa3, 0x1e, 0x6c, 0x65, 0xa4, 0x9a, 0x51, 0x0b, 0xdd, 0x88, 0x8c, 0x8d, 0x60, 0x77, 0xad,
	0x81, 0x5a, 0xfd, 0x07, 0xa9, 0xf7, 0x22, 0x40, 0xad, 0xfd, 0x1c, 0x6e, 0x4e, 0x25, 0x96, 0xe2,
	0x72, 0xf8, 0xe4, 0xef, 0xc3, 0x77, 0x6a, 0x44, 0x18, 0xfb, 0x00, 0x58, 0x83, 0x80, 0x24, 0x7f,
	0x92, 0x64, 0x77, 0xa5, 0xb2, 0x96, 0x1b, 0x41, 0x5b, 0x60, 0xe9, 0xb6, 0x9f, 0xde, 0xb8, 0x56,
	0xe7, 0x44, 0xeb, 0x12, 0x0b, 0x45, 0x3a, 0xa1, 0x94, 0xed, 0x41, 0x87, 0x9e, 0x88, 0xfb, 0x17,
	0x71, 0x03, 0xbd, 0xf3, 0xb4, 0xfb, 0x5f, 0x37, 0x61, 0x9b, 0xa3, 0x5d, 0x68, 0x65, 0x91, 0x3d,
	0x84, 0x56, 0xed, 0x9a, 0x70, 0x8e, 0x97, 0x3b, 0x0c, 0xa6, 0x24, 0x47, 0x1d, 0xfa, 0x5f, 0x4e,
	0x85, 0xec, 0x03, 0xec, 0x78, 0xbf, 0xe4, 0x2b, 0x0b, 0x72, 0x07, 0x90, 0x38, 0x70, 0x16, 0x81,
	0x63, 0x5b, 0x4d, 0x5c, 0x3c, 0xbe, 0x8a, 0x79, 0x77, 0xde, 0x7c, 0xe1, 0x5c, 0xb1, 0x15, 0x7c,
	0xea, 0x36, 0xeb, 0x19, 0xef, 0xad, 0x31, 0x92, 0x8b, 0x27, 0xf4, 0xcf, 0x2f, 0xca, 0xd9, 0x6b,
	0x48, 0x8c, 0x3e, 0x73, 0x7b, 0xf2, 0xa8, 0xa7, 0xd9, 0xbf, 0xdc, 0xac, 0xec, 0x62, 0x17, 0x19,
	0xd7, 0x67, 0xdc, 0xd3, 0xf4, 0xef, 0x42, 0xe2, 0x9e, 0xd9, 0x6d, 0x68, 0xbb, 0xc8, 0x9f, 0xf3,
	0x97, 0x23, 0xb7, 0x9d, 0x16, 0x6f, 0xb9, 0x70, 0x2c, 0x5e, 0x1e, 0xc3, 0x40, 0xea, 0x48, 0xe3,
	0xea, 0xf2, 0x7f, 0x7c, 0xfc, 0x9f, 0x9f, 0x85, 0x93, 0x76, 0x7d, 0xef, 0x46, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0x6b, 0x63, 0x81, 0x58, 0x04, 0x00, 0x00,
}
