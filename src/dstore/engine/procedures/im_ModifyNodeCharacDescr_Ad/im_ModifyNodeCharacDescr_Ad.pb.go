// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyNodeCharacDescr_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyNodeCharacDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyNodeCharacDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyNodeCharacDescr_Ad

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
	NodeCharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull      bool                        `protobuf:"varint,1001,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	LanguageId                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	CharacteristicDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	CharacteristicDescriptionNull bool                        `protobuf:"varint,1003,opt,name=characteristic_description_null,json=characteristicDescriptionNull" json:"characteristic_description_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyNodeCharacDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyNodeCharacDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyNodeCharacDescr_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xe1, 0x6a, 0xd4, 0x40,
	0x10, 0xa6, 0x8d, 0xd7, 0x96, 0x29, 0x68, 0x59, 0xa5, 0xa4, 0x39, 0xaa, 0x52, 0xff, 0xd8, 0x3f,
	0x7b, 0xa2, 0x82, 0x22, 0x22, 0xa8, 0x15, 0xb9, 0x1f, 0x77, 0xe8, 0x0a, 0x82, 0xfe, 0x09, 0x31,
	0x3b, 0x8d, 0x0b, 0x77, 0xbb, 0xc7, 0x6e, 0xce, 0xe2, 0x5b, 0xf8, 0x20, 0xbe, 0x81, 0x4f, 0xa4,
	0xbe, 0x84, 0xb3, 0xd9, 0xbd, 0xa6, 0x89, 0xad, 0x8a, 0x7f, 0xee, 0x32, 0x99, 0xf9, 0xe6, 0xfb,
	0x66, 0xf2, 0x0d, 0x3c, 0x92, 0xae, 0x36, 0x16, 0x47, 0xa8, 0x2b, 0xa5, 0x71, 0xb4, 0xb0, 0xa6,
	0x44, 0xb9, 0xb4, 0xe8, 0x46, 0x6a, 0x9e, 0x4f, 0x8c, 0x54, 0xc7, 0x9f, 0xa7, 0x46, 0xe2, 0xf3,
	0x8f, 0x85, 0x2d, 0xca, 0x23, 0x74, 0xa5, 0xcd, 0x9f, 0x4a, 0x4e, 0x65, 0xb5, 0x61, 0x87, 0x01,
	0xcb, 0x03, 0x96, 0xff, 0x01, 0x90, 0x5d, 0x8d, 0x34, 0x9f, 0x8a, 0xd9, 0x12, 0x5d, 0xc0, 0x67,
	0x7b, 0x5d, 0x6e, 0xb4, 0xd6, 0xd8, 0x98, 0x1a, 0x76, 0x53, 0x73, 0x74, 0xae, 0xa8, 0x30, 0x26,
	0x6f, 0xf5, 0x93, 0x75, 0xa1, 0xf4, 0xb1, 0xb1, 0xf3, 0xa2, 0x56, 0x46, 0x87, 0xa2, 0x83, 0xaf,
	0x09, 0xc0, 0x2b, 0xd2, 0x40, 0x59, 0xb4, 0x8e, 0xbd, 0x86, 0x5d, 0x4d, 0xb2, 0xf2, 0xb2, 0xd1,
	0x45, 0xaf, 0x94, 0xab, 0x55, 0x99, 0x2b, 0x99, 0xae, 0xdd, 0x5c, 0xbb, 0xbd, 0x7d, 0x77, 0xc8,
	0xe3, 0x30, 0x51, 0xa1, 0xd2, 0x35, 0x56, 0x68, 0xdf, 0xfa, 0x48, 0x5c, 0xd3, 0xa7, 0x13, 0xad,
	0x90, 0x63, 0xc9, 0x9e, 0xc0, 0xf0, 0xfc, 0x96, 0xb9, 0x5e, 0xce, 0x66, 0xe9, 0xf7, 0x4d, 0x6a,
	0xbc, 0x25, 0xd2, 0xf3, 0xb0, 0x53, 0x2a, 0x60, 0x8f, 0x61, 0x7b, 0x56, 0xe8, 0x6a, 0x49, 0x83,
	0x79, 0x1d, 0xeb, 0x7f, 0xd7, 0x01, 0xab, 0x7a, 0x62, 0x3f, 0x84, 0x9d, 0x33, 0xe8, 0x40, 0xf9,
	0x23, 0x50, 0x5e, 0x6e, 0xcb, 0x1a, 0xa2, 0x77, 0x90, 0xf5, 0x34, 0x4a, 0xff, 0x5d, 0xd4, 0xc2,
	0xaf, 0x2b, 0x4d, 0x1a, 0xde, 0xac, 0xc7, 0xeb, 0x6a, 0xab, 0x74, 0x15, 0x68, 0xf7, 0xba, 0xe8,
	0xa3, 0x16, 0xcc, 0x5e, 0xc2, 0x8d, 0x8b, 0x5b, 0x07, 0x51, 0x3f, 0x83, 0xa8, 0xfd, 0x0b, 0x9b,
	0x78, 0x8d, 0x07, 0xdf, 0xd6, 0x61, 0x4b, 0xa0, 0x5b, 0x18, 0xed, 0x90, 0xdd, 0x81, 0x41, 0x63,
	0x86, 0xf8, 0x6d, 0x4e, 0xb5, 0x45, 0xa3, 0x05, 0xa3, 0xbc, 0xf0, 0xbf, 0x22, 0x14, 0xd2, 0x88,
	0x3b, 0xde, 0x06, 0xf9, 0x19, 0x1f, 0xd0, 0x42, 0x13, 0x02, 0xf3, 0x1e, 0xb8, 0xef, 0x96, 0x09,
	0xc5, 0xe3, 0x36, 0x16, 0x57, 0xe6, 0xdd, 0x17, 0xec, 0x21, 0x6c, 0x46, 0xfb, 0xd1, 0xaa, 0x7c,
	0xc7, 0xeb, 0xbf, 0x75, 0x0c, 0xe6, 0x9c, 0x84, 0x7f, 0xb1, 0x2a, 0x67, 0x63, 0x48, 0xac, 0x39,
	0x49, 0x2f, 0x35, 0xa8, 0x07, 0xfc, 0x9f, 0xaf, 0x85, 0xaf, 0x16, 0xc1, 0x85, 0x39, 0x11, 0xbe,
	0x47, 0xb6, 0x0f, 0x09, 0x3d, 0xb3, 0x5d, 0xd8, 0xa0, 0xc8, 0xbb, 0xe5, 0xcb, 0x94, 0x56, 0x33,
	0x10, 0x03, 0x0a, 0xc7, 0xf2, 0xd9, 0x1b, 0x18, 0x2a, 0xd3, 0x23, 0x68, 0x4f, 0xf9, 0xfd, 0xfd,
	0xff, 0x39, 0xf2, 0x0f, 0x1b, 0xcd, 0x21, 0xdd, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xd0,
	0x34, 0xfe, 0x23, 0x04, 0x00, 0x00,
}