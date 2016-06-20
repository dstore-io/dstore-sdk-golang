// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifySurchargeTypeDescr_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifySurchargeTypeDescr_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifySurchargeTypeDescr_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifySurchargeTypeDescr_Ad

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
	SurchargeTypeId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull          bool                        `protobuf:"varint,1001,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	LanguageId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull               bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	SurchargeTypeDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=surcharge_type_description,json=surchargeTypeDescription" json:"surcharge_type_description,omitempty"`
	SurchargeTypeDescriptionNull bool                        `protobuf:"varint,1003,opt,name=surcharge_type_description_null,json=surchargeTypeDescriptionNull" json:"surcharge_type_description_null,omitempty"`
	DeleteTranslation            *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_translation,json=deleteTranslation" json:"delete_translation,omitempty"`
	DeleteTranslationNull        bool                        `protobuf:"varint,1004,opt,name=delete_translation_null,json=deleteTranslationNull" json:"delete_translation_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetSurchargeTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeTypeDescription
	}
	return nil
}

func (m *Parameters) GetDeleteTranslation() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteTranslation
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifySurchargeTypeDescr_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifySurchargeTypeDescr_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifySurchargeTypeDescr_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0x13, 0x31,
	0x10, 0x55, 0x9b, 0xde, 0x34, 0x95, 0x68, 0xeb, 0x8a, 0xb2, 0x6c, 0xb8, 0xa9, 0xbc, 0x80, 0x84,
	0x36, 0x08, 0x90, 0x8a, 0x10, 0x3c, 0x80, 0xb8, 0x28, 0x88, 0x54, 0xc8, 0x54, 0x15, 0xf0, 0xb2,
	0x72, 0xe3, 0xe9, 0x62, 0x69, 0x63, 0x47, 0xf6, 0x2e, 0x55, 0xff, 0x82, 0x77, 0xfe, 0x85, 0xff,
	0xe1, 0xf2, 0x11, 0xd8, 0x3b, 0x1b, 0x92, 0xec, 0x52, 0x10, 0xbc, 0x24, 0x99, 0xcc, 0x9c, 0x39,
	0xc7, 0x33, 0xc7, 0x86, 0x47, 0xd2, 0x15, 0xc6, 0x62, 0x0f, 0x75, 0xa6, 0x34, 0xf6, 0xc6, 0xd6,
	0x0c, 0x51, 0x96, 0x16, 0x5d, 0xcf, 0x8c, 0xd2, 0x81, 0x91, 0xea, 0xf8, 0xf4, 0x4d, 0x69, 0x87,
	0x1f, 0x84, 0xcd, 0xf0, 0xe0, 0x74, 0x8c, 0x4f, 0xd1, 0x0d, 0x6d, 0xfa, 0x58, 0x26, 0xbe, 0xb2,
	0x30, 0xec, 0x16, 0xc1, 0x13, 0x82, 0x27, 0x7f, 0xc6, 0xc4, 0xdb, 0x35, 0xd9, 0x47, 0x91, 0x97,
	0xe8, 0xa8, 0x45, 0x7c, 0x71, 0x5e, 0x01, 0x5a, 0x6b, 0x6c, 0x9d, 0xea, 0xce, 0xa7, 0x46, 0xe8,
	0x9c, 0xc8, 0xb0, 0x4e, 0x5e, 0x6f, 0x26, 0x0b, 0xa1, 0xf4, 0xb1, 0xb1, 0x23, 0x51, 0x28, 0xa3,
	0xa9, 0x68, 0xf7, 0xf3, 0x12, 0xc0, 0x6b, 0x61, 0x85, 0xcf, 0xa2, 0x75, 0xec, 0x05, 0x6c, 0xb9,
	0x89, 0xb2, 0xb4, 0xf0, 0xd2, 0x52, 0x25, 0xa3, 0x85, 0x6b, 0x0b, 0x37, 0xd6, 0xef, 0x74, 0x93,
	0xfa, 0x28, 0xb5, 0x38, 0xa5, 0x0b, 0xcc, 0xd0, 0x1e, 0x86, 0x88, 0x6f, 0xb8, 0xd9, 0xf3, 0xf4,
	0x25, 0xbb, 0x07, 0x3b, 0xad, 0x46, 0xa9, 0x2e, 0xf3, 0x3c, 0xfa, 0xba, 0xea, 0xdb, 0xad, 0xf1,
	0xed, 0x06, 0x62, 0xdf, 0xe7, 0xd8, 0x43, 0x58, 0xcf, 0x85, 0xce, 0x4a, 0x7f, 0x88, 0x40, 0xbc,
	0xf8, 0x77, 0x62, 0x98, 0xd4, 0x7b, 0xce, 0x9b, 0xb0, 0x39, 0x83, 0x26, 0xb6, 0x6f, 0xc4, 0x76,
	0x6e, 0x5a, 0x56, 0x11, 0xbd, 0x85, 0xb8, 0x21, 0x4f, 0x86, 0x1d, 0xa8, 0x71, 0x18, 0x4d, 0xd4,
	0xa9, 0x78, 0xe3, 0x06, 0xaf, 0x2b, 0xac, 0xd2, 0x19, 0xd1, 0x46, 0xae, 0xb5, 0x3f, 0xc2, 0xb2,
	0xe7, 0x70, 0xf5, 0xec, 0xce, 0xa4, 0xe9, 0x3b, 0x69, 0xba, 0x74, 0x56, 0x8f, 0x4a, 0xe1, 0x4b,
	0x60, 0x12, 0x73, 0xbf, 0x95, 0xb4, 0xb0, 0x42, 0xbb, 0xbc, 0x5a, 0x5a, 0xb4, 0xf4, 0xdb, 0x89,
	0x1c, 0x19, 0x93, 0xa3, 0xd0, 0x24, 0x6d, 0x8b, 0x60, 0x07, 0x53, 0x14, 0xdb, 0x83, 0x0b, 0xed,
	0x5e, 0xa4, 0xe5, 0x07, 0x69, 0x39, 0xdf, 0x02, 0x05, 0x11, 0xbb, 0x5f, 0x16, 0x61, 0x8d, 0xa3,
	0x1b, 0x1b, 0xed, 0x90, 0xdd, 0x86, 0xe5, 0xca, 0x7b, 0xb5, 0x1f, 0x7e, 0x8d, 0xa7, 0xb6, 0x36,
	0xf9, 0xf2, 0x59, 0xf8, 0xe4, 0x54, 0xc8, 0xde, 0xc1, 0x66, 0x70, 0x5d, 0x3a, 0x63, 0x3b, 0xbf,
	0xd3, 0x8e, 0x07, 0x27, 0x0d, 0x70, 0xd3, 0x9c, 0x03, 0x1f, 0xf7, 0xa7, 0x31, 0xdf, 0x18, 0xcd,
	0xff, 0xc1, 0xee, 0xc3, 0x6a, 0xed, 0x76, 0xbf, 0xad, 0xd0, 0xf1, 0x4a, 0xab, 0x23, 0xdd, 0x85,
	0x01, 0x7d, 0xf3, 0x49, 0x39, 0x7b, 0x05, 0x1d, 0x6b, 0x4e, 0xfc, 0x24, 0x03, 0xea, 0x41, 0xf2,
	0x2f, 0xf7, 0x33, 0x99, 0xcc, 0x22, 0xe1, 0xe6, 0x84, 0x87, 0x36, 0xf1, 0x65, 0xe8, 0xf8, 0xdf,
	0x6c, 0x07, 0x56, 0x7c, 0x14, 0x3c, 0xfb, 0x69, 0xdf, 0x4f, 0x67, 0x99, 0x2f, 0xfb, 0xb0, 0x2f,
	0x9f, 0x1c, 0x42, 0x57, 0x99, 0x06, 0xc7, 0xf4, 0x09, 0x79, 0xbf, 0xf7, 0x9f, 0x8f, 0xcb, 0xd1,
	0x4a, 0x75, 0x7b, 0xef, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x92, 0xc4, 0xad, 0x9e, 0x04,
	0x00, 0x00,
}
