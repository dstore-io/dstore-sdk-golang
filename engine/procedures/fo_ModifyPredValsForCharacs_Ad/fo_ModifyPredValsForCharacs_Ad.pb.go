// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_ModifyPredValsForCharacs_Ad.proto
// DO NOT EDIT!

/*
Package fo_ModifyPredValsForCharacs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_ModifyPredValsForCharacs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_ModifyPredValsForCharacs_Ad

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
	PostingCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=posting_characteristic_id,json=postingCharacteristicId" json:"posting_characteristic_id,omitempty"`
	PostingCharacteristicIdNull bool                        `protobuf:"varint,1001,opt,name=posting_characteristic_id_null,json=postingCharacteristicIdNull" json:"posting_characteristic_id_null,omitempty"`
	OldValue                    *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=old_value,json=oldValue" json:"old_value,omitempty"`
	OldValueNull                bool                        `protobuf:"varint,1002,opt,name=old_value_null,json=oldValueNull" json:"old_value_null,omitempty"`
	Value                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ValueNull                   bool                        `protobuf:"varint,1003,opt,name=value_null,json=valueNull" json:"value_null,omitempty"`
	OnlyValidForForumId         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=only_valid_for_forum_id,json=onlyValidForForumId" json:"only_valid_for_forum_id,omitempty"`
	OnlyValidForForumIdNull     bool                        `protobuf:"varint,1004,opt,name=only_valid_for_forum_id_null,json=onlyValidForForumIdNull" json:"only_valid_for_forum_id_null,omitempty"`
	Country                     *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	CountryNull                 bool                        `protobuf:"varint,1005,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPostingCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingCharacteristicId
	}
	return nil
}

func (m *Parameters) GetOldValue() *dstore_values.StringValue {
	if m != nil {
		return m.OldValue
	}
	return nil
}

func (m *Parameters) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Parameters) GetOnlyValidForForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OnlyValidForForumId
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_ModifyPredValsForCharacs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_ModifyPredValsForCharacs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_ModifyPredValsForCharacs_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_ModifyPredValsForCharacs_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x6f, 0x6b, 0x13, 0x4f,
	0x10, 0xa6, 0x4d, 0xd3, 0xa4, 0xd3, 0xf2, 0xfb, 0xc9, 0x16, 0x4c, 0x9a, 0x68, 0x90, 0x88, 0xe0,
	0x0b, 0xb9, 0x88, 0x0a, 0x8a, 0xa0, 0xa8, 0xd5, 0x42, 0xc0, 0x94, 0x7a, 0x2f, 0x22, 0xfa, 0xe6,
	0x38, 0x6f, 0x37, 0xe7, 0xe1, 0x65, 0x37, 0xec, 0xde, 0xb5, 0xe4, 0x5b, 0xf8, 0x89, 0xfc, 0x00,
	0x7e, 0x13, 0xff, 0x7d, 0x07, 0x67, 0x6f, 0xf6, 0xd2, 0xe6, 0x34, 0x86, 0xbe, 0x48, 0x6e, 0x67,
	0x67, 0x9e, 0xe7, 0x99, 0x9d, 0x9d, 0x59, 0x78, 0xc2, 0x4d, 0xa6, 0xb4, 0x18, 0x08, 0x19, 0x27,
	0x52, 0x0c, 0x66, 0x5a, 0x45, 0x82, 0xe7, 0x5a, 0x98, 0xc1, 0x44, 0x05, 0x23, 0xc5, 0x93, 0xc9,
	0xfc, 0x44, 0x0b, 0x3e, 0x0e, 0x53, 0x73, 0xa4, 0xf4, 0xe1, 0xc7, 0x50, 0x87, 0x91, 0x09, 0x9e,
	0x73, 0x0f, 0x23, 0x33, 0xc5, 0xee, 0x10, 0xdc, 0x23, 0xb8, 0xf7, 0x6f, 0x4c, 0x67, 0xdf, 0x89,
	0x9d, 0x86, 0x69, 0x2e, 0x0c, 0x51, 0x74, 0x0e, 0x96, 0x33, 0x10, 0x5a, 0x2b, 0xed, 0x5c, 0xdd,
	0x65, 0xd7, 0x54, 0x18, 0x13, 0xc6, 0xc2, 0x39, 0x6f, 0x56, 0x9d, 0x59, 0x98, 0xc8, 0x89, 0xd2,
	0xd3, 0x30, 0x4b, 0x94, 0xa4, 0xa0, 0xfe, 0xd7, 0x2d, 0x80, 0x13, 0xd4, 0x47, 0xaf, 0xd0, 0x86,
	0xbd, 0x85, 0x83, 0x99, 0x32, 0x59, 0x22, 0xe3, 0x20, 0x2a, 0xd2, 0xc2, 0xdd, 0x04, 0xed, 0x28,
	0x48, 0x78, 0x7b, 0xe3, 0xc6, 0xc6, 0xed, 0xdd, 0x7b, 0x5d, 0xcf, 0x1d, 0xc9, 0x25, 0x99, 0xc8,
	0x4c, 0xc4, 0x42, 0x8f, 0xad, 0xe5, 0xb7, 0x1c, 0xfa, 0x70, 0x09, 0x3c, 0xe4, 0xec, 0x25, 0xf4,
	0x56, 0x12, 0x07, 0x32, 0x4f, 0xd3, 0xf6, 0xb7, 0x06, 0xd2, 0x37, 0xfd, 0xee, 0x0a, 0x86, 0x63,
	0x8c, 0x61, 0x0f, 0x61, 0x47, 0xa5, 0x3c, 0x28, 0x94, 0xdb, 0x9b, 0x45, 0x3a, 0x9d, 0x4a, 0x3a,
	0x26, 0xd3, 0x88, 0xa6, 0x6c, 0x9a, 0x18, 0x5c, 0xac, 0xd8, 0x2d, 0xf8, 0x6f, 0x01, 0x24, 0xb9,
	0xef, 0x24, 0xb7, 0x57, 0x86, 0x14, 0xfc, 0x77, 0xa1, 0x4e, 0xdc, 0xb5, 0xb5, 0xdc, 0x14, 0xc8,
	0x7a, 0x00, 0x17, 0x48, 0x7f, 0x10, 0xe9, 0xce, 0xe9, 0x82, 0xf1, 0x0d, 0xb4, 0x94, 0x4c, 0xe7,
	0x56, 0x19, 0x0f, 0x8a, 0xd5, 0xb7, 0xbf, 0x7c, 0x6a, 0xcb, 0xb9, 0xb5, 0xbe, 0x9c, 0xfb, 0x16,
	0x3b, 0xb6, 0x50, 0xec, 0x93, 0x23, 0x0b, 0xc4, 0x52, 0x3e, 0x85, 0x6b, 0x2b, 0x28, 0x29, 0x89,
	0x9f, 0x94, 0x44, 0xeb, 0x2f, 0xd8, 0x22, 0xa5, 0x07, 0xd0, 0x88, 0x54, 0x2e, 0x33, 0x3d, 0x6f,
	0xd7, 0xd7, 0x1e, 0xb3, 0x0c, 0x65, 0x7d, 0xd8, 0x73, 0x4b, 0x52, 0xf9, 0x45, 0x2a, 0xbb, 0x6e,
	0xd3, 0x32, 0xf7, 0xbf, 0x6c, 0x42, 0xd3, 0x17, 0x66, 0xa6, 0xa4, 0x11, 0xb6, 0x96, 0x45, 0xab,
	0xba, 0xb6, 0x59, 0x88, 0xb8, 0x49, 0xa0, 0x36, 0x7e, 0x65, 0xff, 0x7d, 0x0a, 0x64, 0xef, 0xe0,
	0x8a, 0x6d, 0xd2, 0xe0, 0x42, 0x97, 0xe2, 0x25, 0xd7, 0x10, 0xec, 0x55, 0xc0, 0xd5, 0x5e, 0x1e,
	0xa1, 0x3d, 0x3c, 0xb7, 0xfd, 0xff, 0xa7, 0xcb, 0x1b, 0xec, 0x11, 0x34, 0xdc, 0x70, 0xe0, 0xd5,
	0x5a, 0xc6, 0xde, 0x1f, 0x8c, 0x34, 0x3a, 0x23, 0xfa, 0xfa, 0x65, 0x38, 0x7b, 0x0d, 0x35, 0xad,
	0xce, 0xf0, 0xb2, 0x2c, 0xea, 0xb1, 0x77, 0x99, 0x71, 0xf6, 0xca, 0x5a, 0x78, 0xbe, 0x3a, 0xf3,
	0x2d, 0x4d, 0xe7, 0x3a, 0xd4, 0x70, 0xcd, 0xae, 0xc2, 0x36, 0x5a, 0xb6, 0x09, 0x3e, 0x1f, 0x63,
	0x75, 0xea, 0x7e, 0x1d, 0xcd, 0x21, 0x7f, 0x11, 0x40, 0x37, 0x51, 0x15, 0x8d, 0xf3, 0x17, 0xe7,
	0xfd, 0xb3, 0x58, 0x19, 0xfe, 0xa9, 0xf4, 0xf3, 0xcb, 0x3f, 0x4a, 0x1f, 0xb6, 0x8b, 0xa9, 0xbf,
	0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x05, 0xbb, 0x94, 0xd6, 0x04, 0x00, 0x00,
}
