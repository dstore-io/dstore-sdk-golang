// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderStateDescr_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderStateDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderStateDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderStateDescr_Ad

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
	OrderStateId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=order_state_id,json=orderStateId" json:"order_state_id,omitempty"`
	OrderStateIdNull bool                        `protobuf:"varint,1001,opt,name=order_state_id_null,json=orderStateIdNull" json:"order_state_id_null,omitempty"`
	LanguageId       *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull   bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderStateId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	RowId                 int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	LanguageSymbol        *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=language_symbol,json=languageSymbol" json:"language_symbol,omitempty"`
	Language              *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=language" json:"language,omitempty"`
	TranslatedDescription *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=translated_description,json=translatedDescription" json:"translated_description,omitempty"`
	OrderStateId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=order_state_id,json=orderStateId" json:"order_state_id,omitempty"`
	PublicDescription     *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=public_description,json=publicDescription" json:"public_description,omitempty"`
	OrderState            *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=order_state,json=orderState" json:"order_state,omitempty"`
	LanguageId            *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
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

func (m *Response_Row) GetTranslatedDescription() *dstore_values.StringValue {
	if m != nil {
		return m.TranslatedDescription
	}
	return nil
}

func (m *Response_Row) GetOrderStateId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderStateId
	}
	return nil
}

func (m *Response_Row) GetPublicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PublicDescription
	}
	return nil
}

func (m *Response_Row) GetOrderState() *dstore_values.StringValue {
	if m != nil {
		return m.OrderState
	}
	return nil
}

func (m *Response_Row) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderStateDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderStateDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderStateDescr_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x25, 0xae, 0x49, 0xc3, 0xad, 0xb4, 0x71, 0x8a, 0x65, 0x4d, 0x40, 0xa4, 0x82, 0xe8, 0xcb,
	0xc4, 0x2f, 0xa8, 0x0f, 0xfa, 0xd0, 0x12, 0x95, 0x0a, 0x5d, 0xcb, 0x14, 0x04, 0x7d, 0x59, 0x36,
	0xd9, 0x71, 0x19, 0xd8, 0x9d, 0x09, 0x33, 0xb3, 0x16, 0xff, 0x85, 0xdf, 0xe2, 0x3f, 0x54, 0xff,
	0x83, 0xe0, 0xcc, 0xce, 0x26, 0xbb, 0x9b, 0x22, 0xdb, 0xbc, 0x24, 0xdc, 0xbd, 0xe7, 0xdc, 0x7b,
	0xee, 0xb9, 0x33, 0x03, 0xfb, 0xb1, 0xd2, 0x42, 0xd2, 0x31, 0xe5, 0x09, 0xe3, 0x74, 0x3c, 0x97,
	0x62, 0x46, 0xe3, 0x5c, 0x52, 0x35, 0x16, 0x59, 0xf8, 0x82, 0xea, 0x57, 0x32, 0xa6, 0xf2, 0x54,
	0x47, 0x9a, 0x4e, 0xa8, 0x9a, 0xc9, 0xf0, 0x20, 0xc6, 0x06, 0xa3, 0x05, 0xba, 0xed, 0x88, 0xd8,
	0x11, 0xf1, 0xff, 0xd0, 0xc3, 0x9d, 0xb2, 0xc1, 0xfb, 0x28, 0xcd, 0xa9, 0x72, 0xe4, 0xe1, 0xf5,
	0x66, 0x57, 0x2a, 0xa5, 0x90, 0x65, 0x6a, 0xd4, 0x4c, 0x65, 0x54, 0xa9, 0x28, 0xa1, 0x65, 0xf2,
	0xd6, 0x6a, 0x52, 0x47, 0x8c, 0xbf, 0x13, 0x32, 0x8b, 0x34, 0x13, 0xdc, 0x81, 0xf6, 0x7e, 0x75,
	0x00, 0x4e, 0x22, 0x19, 0x99, 0x2c, 0x95, 0x0a, 0x1d, 0xc0, 0x96, 0xb0, 0xb2, 0x42, 0x65, 0x75,
	0x85, 0x2c, 0xf6, 0x3b, 0x37, 0x3b, 0x77, 0x36, 0x1f, 0x8c, 0x70, 0x39, 0x41, 0xa9, 0x8c, 0x71,
	0x4d, 0x13, 0x2a, 0x5f, 0xdb, 0x88, 0x5c, 0x11, 0xcb, 0x49, 0x8e, 0x62, 0x84, 0x61, 0xa7, 0x59,
	0x22, 0xe4, 0x79, 0x9a, 0xfa, 0xbf, 0x37, 0x4c, 0xa1, 0x3e, 0x19, 0xd4, 0xb1, 0x81, 0x49, 0xa0,
	0x27, 0xb0, 0x99, 0x46, 0x3c, 0xc9, 0x8d, 0x70, 0xdb, 0xef, 0x52, 0x7b, 0x3f, 0x58, 0xe0, 0x4d,
	0xb7, 0xbb, 0x30, 0xa8, 0xb1, 0x5d, 0xab, 0x3f, 0xae, 0xd5, 0x56, 0x05, 0xb3, 0x8d, 0xf6, 0x7e,
	0xf6, 0xa0, 0x4f, 0xa8, 0x9a, 0x0b, 0xae, 0x28, 0xba, 0x07, 0xdd, 0xc2, 0xc8, 0x72, 0xbe, 0x21,
	0x6e, 0x6e, 0xc8, 0x99, 0xfc, 0xcc, 0xfe, 0x12, 0x07, 0x44, 0x6f, 0x60, 0x60, 0x2d, 0x0c, 0x6b,
	0x1e, 0x1a, 0xb1, 0x9e, 0x21, 0xe3, 0x15, 0xf2, 0xaa, 0xd3, 0xc7, 0x26, 0x3e, 0xaa, 0x62, 0xb2,
	0x9d, 0x35, 0x3f, 0xa0, 0xc7, 0xb0, 0x51, 0xae, 0xce, 0xf7, 0x8a, 0x8a, 0x37, 0xce, 0x55, 0x74,
	0x8b, 0x3d, 0x76, 0xff, 0x64, 0x01, 0x47, 0xcf, 0xc1, 0x93, 0xe2, 0xcc, 0xbf, 0x5c, 0xb0, 0x1e,
	0xe1, 0x8b, 0x1d, 0x33, 0xbc, 0x70, 0x01, 0x13, 0x71, 0x46, 0x6c, 0x81, 0xe1, 0x5f, 0x0f, 0x3c,
	0x13, 0xa0, 0x5d, 0xe8, 0x99, 0xd0, 0xee, 0xe1, 0x63, 0x60, 0x8c, 0xe9, 0x92, 0xae, 0x09, 0x8d,
	0xcd, 0x13, 0xd8, 0x5e, 0xda, 0xac, 0x3e, 0x64, 0x53, 0x91, 0xfa, 0x9f, 0x82, 0xa6, 0x73, 0xe5,
	0xa6, 0x94, 0x96, 0x8c, 0x27, 0x6e, 0x51, 0xcb, 0x0d, 0x9c, 0x16, 0x14, 0xb4, 0x0f, 0xfd, 0xc5,
	0x17, 0xff, 0x73, 0x3b, 0x7d, 0x09, 0x46, 0x04, 0x76, 0xb5, 0x8c, 0xb8, 0x4a, 0xcd, 0x14, 0x71,
	0x18, 0xdb, 0x39, 0xd8, 0xbc, 0xd8, 0xc0, 0x97, 0xf6, 0x32, 0xd7, 0x2a, 0xea, 0xa4, 0x62, 0xa2,
	0xc3, 0x73, 0x47, 0xfd, 0x6b, 0xb0, 0xee, 0x59, 0x7f, 0x09, 0x68, 0x9e, 0x4f, 0x53, 0x36, 0x6b,
	0x68, 0xfa, 0xd6, 0xae, 0xe9, 0xaa, 0xa3, 0xd5, 0xf5, 0x98, 0x7b, 0x50, 0xd3, 0xe3, 0x7f, 0x6f,
	0x2f, 0x02, 0x95, 0x16, 0xf4, 0xb4, 0x79, 0x8b, 0x7e, 0x04, 0x6b, 0x5d, 0xa3, 0xc3, 0x13, 0x18,
	0x31, 0xb1, 0x72, 0x7c, 0xaa, 0xe7, 0xed, 0xed, 0xfd, 0xb5, 0x1f, 0xbe, 0x69, 0xaf, 0x78, 0x5f,
	0x1e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x39, 0x09, 0x2a, 0x34, 0x05, 0x00, 0x00,
}
