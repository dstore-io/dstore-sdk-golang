// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetLanguages.proto
// DO NOT EDIT!

/*
Package mi_GetLanguages is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetLanguages.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetLanguages

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
	LanguageId               *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull           bool                        `protobuf:"varint,1001,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	GetLocalLanguageName     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=get_local_language_name,json=getLocalLanguageName" json:"get_local_language_name,omitempty"`
	GetLocalLanguageNameNull bool                        `protobuf:"varint,1002,opt,name=get_local_language_name_null,json=getLocalLanguageNameNull" json:"get_local_language_name_null,omitempty"`
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

func (m *Parameters) GetGetLocalLanguageName() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetLocalLanguageName
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
	LanguageSymbol    *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=language_symbol,json=languageSymbol" json:"language_symbol,omitempty"`
	Language          *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=language" json:"language,omitempty"`
	LanguageId        *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LocalLanguageName *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=local_language_name,json=localLanguageName" json:"local_language_name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetLanguageSymbol() *dstore_values.StringValue {
	if m != nil {
		return m.LanguageSymbol
	}
	return nil
}

func (m *Response_Row) GetLanguage() *dstore_values.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Response_Row) GetLocalLanguageName() *dstore_values.StringValue {
	if m != nil {
		return m.LocalLanguageName
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetLanguages.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetLanguages.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetLanguages.Response.Row")
}

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6a, 0x14, 0x41,
	0x10, 0x26, 0xbb, 0xd9, 0x64, 0xa9, 0x80, 0x89, 0x1d, 0xd1, 0x71, 0x57, 0x45, 0xe2, 0x25, 0x22,
	0xf4, 0x8a, 0x1e, 0xf4, 0xa0, 0x08, 0xa2, 0x84, 0x60, 0x32, 0x48, 0x0b, 0x82, 0x5e, 0x86, 0xde,
	0x6c, 0x39, 0x0c, 0xf4, 0x74, 0x87, 0xee, 0x59, 0x83, 0xaf, 0xe0, 0xc9, 0xbf, 0x87, 0xf0, 0x1d,
	0x7c, 0x22, 0x7d, 0x0a, 0x6b, 0xa6, 0x7b, 0x66, 0x33, 0xe3, 0x6a, 0xcc, 0x65, 0x97, 0xea, 0xfa,
	0xbe, 0xea, 0xaa, 0xef, 0xab, 0x69, 0xe0, 0x33, 0x57, 0x18, 0x8b, 0x13, 0xd4, 0x69, 0xa6, 0x71,
	0x72, 0x6c, 0xcd, 0x11, 0xce, 0xe6, 0x16, 0xdd, 0x24, 0xcf, 0x92, 0x3d, 0x2c, 0x0e, 0xa4, 0x4e,
	0xe7, 0x32, 0x45, 0xc7, 0x29, 0x55, 0x18, 0x76, 0xdd, 0xe3, 0xb9, 0xc7, 0xf3, 0x0e, 0x68, 0xb4,
	0x1d, 0xca, 0xbd, 0x97, 0x6a, 0x5e, 0x73, 0x46, 0x57, 0xdb, 0x77, 0xa0, 0xb5, 0xc6, 0x86, 0xd4,
	0xb8, 0x9d, 0xca, 0xd1, 0x39, 0xaa, 0x13, 0x92, 0xb7, 0xba, 0xc9, 0x42, 0x66, 0xfa, 0x9d, 0xb1,
	0xb9, 0x2c, 0x32, 0xa3, 0x3d, 0x68, 0xe7, 0x63, 0x0f, 0xe0, 0xa5, 0xb4, 0x92, 0xb2, 0x68, 0x1d,
	0x7b, 0x04, 0x1b, 0x2a, 0x74, 0x93, 0x64, 0xb3, 0x68, 0xe5, 0xe6, 0xca, 0xee, 0xc6, 0xbd, 0x71,
	0x98, 0x92, 0x87, 0xb6, 0x32, 0x5d, 0x60, 0x8a, 0xf6, 0x75, 0x19, 0x09, 0xa8, 0xf1, 0xfb, 0x33,
	0x76, 0x1b, 0xb6, 0x4e, 0xb1, 0x13, 0x3d, 0x57, 0x2a, 0xfa, 0xb9, 0x4e, 0x35, 0x86, 0xe2, 0xc2,
	0x02, 0x16, 0xd3, 0x31, 0x13, 0x70, 0x25, 0xc5, 0x22, 0x51, 0xe6, 0x48, 0xaa, 0xa4, 0x21, 0x69,
	0xea, 0x23, 0xea, 0x2d, 0xbd, 0x74, 0x6a, 0x8c, 0x42, 0xa9, 0xfd, 0xa5, 0x97, 0x88, 0x7b, 0x50,
	0x52, 0x6b, 0xe9, 0x62, 0x22, 0xb2, 0x27, 0x70, 0xed, 0x2f, 0x35, 0x7d, 0x2b, 0xbf, 0x7c, 0x2b,
	0xd1, 0x32, 0x72, 0xd9, 0xd4, 0xce, 0x8f, 0x55, 0x18, 0x0a, 0x74, 0xc7, 0x46, 0x3b, 0x64, 0x77,
	0x61, 0x50, 0x49, 0x1d, 0x44, 0x18, 0xf1, 0xb6, 0x75, 0xde, 0x86, 0xe7, 0xe5, 0xaf, 0xf0, 0x40,
	0xf6, 0x06, 0xb6, 0x4a, 0x91, 0x93, 0x53, 0x2a, 0xd3, 0x30, 0x7d, 0x22, 0xf3, 0x0e, 0xb9, 0xeb,
	0xc5, 0x21, 0xc5, 0xfb, 0x8b, 0x58, 0x6c, 0xe6, 0xed, 0x03, 0xf6, 0x10, 0xd6, 0x83, 0xb9, 0x51,
	0xbf, 0xaa, 0x78, 0xe3, 0x8f, 0x8a, 0xde, 0xfa, 0x43, 0xff, 0x2f, 0x6a, 0x38, 0x7b, 0x0c, 0x7d,
	0x6b, 0x4e, 0xa2, 0xd5, 0x8a, 0x75, 0x87, 0xff, 0x73, 0xff, 0x78, 0x3d, 0x3c, 0x17, 0xe6, 0x44,
	0x94, 0xbc, 0xd1, 0xf7, 0x1e, 0xf4, 0x29, 0x60, 0x97, 0x61, 0x8d, 0xc2, 0x72, 0x27, 0x3e, 0xc5,
	0xa4, 0xc7, 0x40, 0x0c, 0x28, 0x24, 0xcb, 0x9f, 0xc1, 0x66, 0xa3, 0xb4, 0xfb, 0x90, 0x4f, 0x8d,
	0x8a, 0x3e, 0xc7, 0x6d, 0xc1, 0x82, 0x81, 0xae, 0xb0, 0x99, 0x4e, 0xbd, 0x7f, 0xcd, 0x36, 0xbc,
	0xaa, 0x28, 0xec, 0x01, 0x0c, 0xeb, 0x93, 0xe8, 0xcb, 0xd9, 0xf4, 0x06, 0x4c, 0xd3, 0xb5, 0xf6,
	0xf5, 0x6b, 0x7c, 0xbe, 0x85, 0x7d, 0x01, 0xdb, 0xcb, 0x36, 0xf0, 0xdb, 0xd9, 0x2d, 0x5c, 0x54,
	0xdd, 0x0d, 0x7a, 0xba, 0x07, 0xe3, 0xcc, 0x74, 0x04, 0x5e, 0x3c, 0x08, 0x6f, 0x77, 0xff, 0xf7,
	0xa9, 0x98, 0xae, 0x55, 0x9f, 0xe6, 0xfd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0xb8, 0x4f,
	0x19, 0x5d, 0x04, 0x00, 0x00,
}