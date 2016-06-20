// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyLanguageDescr_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyLanguageDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyLanguageDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyLanguageDescr_Ad

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
	LanguageId                *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull            bool                        `protobuf:"varint,1001,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	TranslationLanguageId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=translation_language_id,json=translationLanguageId" json:"translation_language_id,omitempty"`
	TranslationLanguageIdNull bool                        `protobuf:"varint,1002,opt,name=translation_language_id_null,json=translationLanguageIdNull" json:"translation_language_id_null,omitempty"`
	LanguageName              *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=language_name,json=languageName" json:"language_name,omitempty"`
	LanguageNameNull          bool                        `protobuf:"varint,1003,opt,name=language_name_null,json=languageNameNull" json:"language_name_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetTranslationLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TranslationLanguageId
	}
	return nil
}

func (m *Parameters) GetLanguageName() *dstore_values.StringValue {
	if m != nil {
		return m.LanguageName
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyLanguageDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyLanguageDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyLanguageDescr_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x66, 0x1b, 0xbb, 0xbb, 0x9c, 0xf5, 0xa7, 0x8c, 0xa8, 0x69, 0xaa, 0x22, 0xf5, 0x42, 0xbd,
	0x70, 0x2a, 0x15, 0xa1, 0x17, 0x82, 0x3f, 0x28, 0x52, 0x68, 0x8b, 0x8e, 0x20, 0xe8, 0x4d, 0x18,
	0x9b, 0x69, 0x18, 0x48, 0x66, 0xca, 0x4c, 0x62, 0xf1, 0x2d, 0x7c, 0x1d, 0x1f, 0x49, 0x7d, 0x05,
	0x2f, 0x9c, 0x64, 0x26, 0xcd, 0x8f, 0x16, 0xf5, 0x66, 0xb7, 0x27, 0xe7, 0xfb, 0xe3, 0x9c, 0x33,
	0x30, 0x8b, 0x74, 0x26, 0x15, 0x9b, 0x30, 0x11, 0x73, 0xc1, 0x26, 0x5b, 0x25, 0xd7, 0x2c, 0xca,
	0x15, 0xd3, 0x93, 0x94, 0x87, 0x4b, 0x19, 0xf1, 0xcd, 0xe7, 0x05, 0x15, 0x71, 0x4e, 0x63, 0xf6,
	0x82, 0xe9, 0xb5, 0x0a, 0x9f, 0x45, 0xd8, 0x80, 0x32, 0x89, 0xee, 0x58, 0x26, 0xb6, 0x4c, 0x7c,
	0x10, 0x1e, 0x5c, 0x76, 0x16, 0x9f, 0x68, 0x92, 0x33, 0x6d, 0xd9, 0xc1, 0xb0, 0xed, 0xcb, 0x94,
	0x92, 0xca, 0xb5, 0x46, 0xed, 0x56, 0xca, 0xb4, 0x36, 0x72, 0xae, 0x79, 0xbb, 0xdb, 0xcc, 0x28,
	0x17, 0x1b, 0xa9, 0x52, 0x9a, 0x71, 0x29, 0x2c, 0x68, 0xfc, 0xb3, 0x07, 0xf0, 0x9a, 0x2a, 0x6a,
	0xba, 0x4c, 0x69, 0xf4, 0x18, 0xce, 0x12, 0x17, 0x2a, 0xe4, 0x91, 0x7f, 0x74, 0xeb, 0xe8, 0xee,
	0xd9, 0x74, 0x84, 0x5d, 0x7e, 0x17, 0x8b, 0x8b, 0x8c, 0xc5, 0x4c, 0xbd, 0x2b, 0x2a, 0x02, 0x15,
	0x7e, 0x1e, 0xa1, 0x7b, 0x30, 0x68, 0xb0, 0x43, 0x91, 0x27, 0x89, 0xff, 0xed, 0xc4, 0x68, 0x9c,
	0x92, 0x8b, 0x35, 0x6c, 0x65, 0x3e, 0xa3, 0xb7, 0x70, 0x2d, 0x53, 0x54, 0xe8, 0xa4, 0x0c, 0x13,
	0x36, 0x4d, 0x7b, 0x7f, 0x37, 0xbd, 0xd2, 0xe0, 0x2e, 0x6a, 0xff, 0xa7, 0x70, 0xfd, 0x80, 0xa8,
	0xcd, 0xf2, 0xdd, 0x66, 0x19, 0xfe, 0x91, 0x5d, 0xc6, 0x7a, 0x02, 0x17, 0xf6, 0x2c, 0x61, 0x86,
	0xe2, 0x7b, 0x65, 0x98, 0xa0, 0x13, 0x46, 0x67, 0x8a, 0x8b, 0xd8, 0x66, 0x39, 0x5f, 0x11, 0x56,
	0x06, 0x8f, 0xee, 0x03, 0x6a, 0x09, 0x58, 0xe3, 0x1f, 0xd6, 0x78, 0xd0, 0x84, 0x16, 0x7e, 0xe3,
	0xaf, 0x3d, 0x38, 0x25, 0x4c, 0x6f, 0xa5, 0xd0, 0x0c, 0x3d, 0x80, 0x7e, 0xb9, 0x5c, 0x37, 0xf6,
	0xbd, 0xa9, 0x3b, 0x1b, 0xbb, 0xf8, 0x97, 0xc5, 0x5f, 0x62, 0x81, 0xe8, 0x3d, 0x0c, 0x8a, 0xb5,
	0x86, 0x8d, 0xbd, 0x9a, 0xf1, 0x79, 0x86, 0x8c, 0x3b, 0xe4, 0xee, 0xf6, 0x97, 0xa6, 0x9e, 0xd7,
	0x35, 0xb9, 0x94, 0xb6, 0x3f, 0xa0, 0x19, 0x9c, 0xb8, 0x73, 0x32, 0x33, 0x28, 0x14, 0x6f, 0xfe,
	0xa6, 0x68, 0x8f, 0x6d, 0x69, 0xff, 0x93, 0x0a, 0x8e, 0x5e, 0x81, 0xa7, 0xe4, 0xce, 0x3f, 0x57,
	0xb2, 0x1e, 0xe1, 0x7f, 0xbc, 0x7d, 0x5c, 0x8d, 0x01, 0x13, 0xb9, 0x23, 0x85, 0x42, 0x70, 0x03,
	0x3c, 0xf3, 0x1b, 0x5d, 0x85, 0x63, 0x53, 0x15, 0x97, 0xf1, 0x65, 0x65, 0x06, 0xd3, 0x27, 0x7d,
	0x53, 0xce, 0xa3, 0xe7, 0x6f, 0x60, 0xc4, 0x65, 0x47, 0xbe, 0x7e, 0x94, 0x1f, 0xa6, 0xff, 0xff,
	0x5c, 0x3f, 0x1e, 0x97, 0x8f, 0xe2, 0xe1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x96, 0x4d,
	0x29, 0xeb, 0x03, 0x00, 0x00,
}
