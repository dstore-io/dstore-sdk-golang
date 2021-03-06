// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonTypeMetaInfo_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersonTypeMetaInfo_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonTypeMetaInfo_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonTypeMetaInfo_Ad

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
	PersonTypeId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull              bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	MetaInformationTypeIdList     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=meta_information_type_id_list,json=metaInformationTypeIdList" json:"meta_information_type_id_list,omitempty"`
	MetaInformationTypeIdListNull bool                        `protobuf:"varint,1002,opt,name=meta_information_type_id_list_null,json=metaInformationTypeIdListNull" json:"meta_information_type_id_list_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetMetaInformationTypeIdList() *dstore_values.StringValue {
	if m != nil {
		return m.MetaInformationTypeIdList
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	MetaInformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=meta_information_type_id,json=metaInformationTypeId" json:"meta_information_type_id,omitempty"`
	MetaInformationType       *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=meta_information_type,json=metaInformationType" json:"meta_information_type,omitempty"`
	MetaInformation           *dstore_values.DecimalValue `protobuf:"bytes,10003,opt,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	PersonTypeId              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	TranslatedMetaInformation *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=translated_meta_information,json=translatedMetaInformation" json:"translated_meta_information,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetMetaInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.MetaInformationTypeId
	}
	return nil
}

func (m *Response_Row) GetMetaInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.MetaInformationType
	}
	return nil
}

func (m *Response_Row) GetMetaInformation() *dstore_values.DecimalValue {
	if m != nil {
		return m.MetaInformation
	}
	return nil
}

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Response_Row) GetTranslatedMetaInformation() *dstore_values.StringValue {
	if m != nil {
		return m.TranslatedMetaInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonTypeMetaInfo_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonTypeMetaInfo_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonTypeMetaInfo_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_GetPersonTypeMetaInfo_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0xb5, 0xe5, 0xd7, 0x6d, 0x3a, 0x3f, 0x04, 0x93, 0xab, 0xa1, 0xb4, 0xd5, 0x10, 0x2a,
	0x37, 0x70, 0x93, 0x22, 0xb8, 0x00, 0x21, 0x71, 0xb1, 0x49, 0x68, 0x2a, 0xa2, 0x61, 0x8a, 0x00,
	0x09, 0x98, 0x14, 0x99, 0xf9, 0x10, 0x59, 0x24, 0x76, 0x64, 0xbb, 0x4c, 0xbc, 0x05, 0x7f, 0x1f,
	0x82, 0x2b, 0xde, 0x84, 0xf7, 0x80, 0xa7, 0xc0, 0x89, 0xdd, 0xb2, 0xa6, 0xd9, 0xca, 0x6e, 0x5a,
	0x59, 0xe7, 0x7c, 0xbf, 0xfe, 0xd8, 0xdf, 0x13, 0xc3, 0x03, 0xa6, 0x8d, 0x54, 0x38, 0x42, 0x91,
	0x71, 0x81, 0xa3, 0x52, 0xc9, 0x63, 0x64, 0x53, 0x85, 0x7a, 0x54, 0x16, 0xe9, 0x01, 0x9a, 0x43,
	0x54, 0x5a, 0x8a, 0x67, 0x1f, 0x4a, 0x9c, 0xa0, 0xa1, 0x63, 0xf1, 0x56, 0xa6, 0x7b, 0x2c, 0xb2,
	0x6d, 0x46, 0x92, 0x5b, 0x4e, 0x1b, 0x39, 0x6d, 0x74, 0x8e, 0xa0, 0xdf, 0xf5, 0xdb, 0xbc, 0xa7,
	0xf9, 0x14, 0xb5, 0xd3, 0xf7, 0x7b, 0x8b, 0x7b, 0xa3, 0x52, 0x52, 0xf9, 0xd2, 0x60, 0xb1, 0x54,
	0xa0, 0xd6, 0x34, 0x43, 0x5f, 0xbc, 0xd1, 0x2c, 0x1a, 0xca, 0xed, 0x3e, 0xaa, 0xa0, 0x86, 0x4b,
	0xe1, 0x9a, 0x86, 0xdf, 0xd7, 0x01, 0x0e, 0xa9, 0xa2, 0xb6, 0x6a, 0x89, 0xc8, 0x1e, 0x5c, 0x2e,
	0x6b, 0xb2, 0xd4, 0x58, 0xb4, 0x94, 0xb3, 0x70, 0xed, 0xfa, 0xda, 0xcd, 0xff, 0xef, 0x0c, 0x22,
	0x7f, 0x08, 0x4f, 0xc6, 0x85, 0xc1, 0x0c, 0xd5, 0x8b, 0x6a, 0x95, 0x5c, 0x2a, 0xe7, 0x87, 0x19,
	0x33, 0x12, 0x41, 0x77, 0xd1, 0x22, 0x15, 0xd3, 0x3c, 0x0f, 0x7f, 0x6d, 0x5a, 0xa3, 0xad, 0x64,
	0xfb, 0x74, 0x6f, 0x6c, 0x0b, 0xe4, 0x08, 0x76, 0x2b, 0xb4, 0xf4, 0x14, 0xdb, 0x5c, 0x99, 0x73,
	0x6d, 0xc2, 0xf5, 0x9a, 0xa0, 0xdf, 0x20, 0xd0, 0x46, 0x71, 0x91, 0x39, 0x80, 0x5e, 0xe1, 0xef,
	0xd0, 0xeb, 0x9d, 0xfb, 0x13, 0x2b, 0x26, 0x8f, 0x61, 0x78, 0xae, 0xbb, 0x83, 0xfb, 0xed, 0xe0,
	0x76, 0xcf, 0xf4, 0xa9, 0x48, 0x87, 0x3f, 0x3b, 0xb0, 0x95, 0xa0, 0x2e, 0xa5, 0xd0, 0x48, 0x6e,
	0x43, 0xa7, 0x4e, 0xc2, 0x5f, 0xd0, 0x1c, 0xcf, 0xa7, 0xec, 0x52, 0x7a, 0x54, 0xfd, 0x26, 0xae,
	0x91, 0xbc, 0x84, 0xed, 0x26, 0x8a, 0x3d, 0x5b, 0x60, 0xc5, 0x51, 0x43, 0xdc, 0x8c, 0x6a, 0xb2,
	0x88, 0x95, 0x5c, 0x69, 0x70, 0x92, 0xfb, 0xb0, 0xe9, 0xb3, 0x0f, 0x83, 0xda, 0xf1, 0xda, 0x92,
	0xa3, 0x9b, 0x8c, 0x89, 0xfb, 0x4f, 0x66, 0xed, 0x64, 0x0c, 0x81, 0x92, 0x27, 0xe1, 0x7f, 0xb5,
	0xea, 0x5e, 0xf4, 0xcf, 0xa3, 0x1a, 0xcd, 0x2e, 0x22, 0x4a, 0xe4, 0x49, 0x52, 0x79, 0xf4, 0x7f,
	0x04, 0x10, 0xd8, 0x05, 0xb9, 0x0a, 0x1b, 0x76, 0x59, 0xcd, 0xce, 0xc7, 0xd8, 0xde, 0x4d, 0x27,
	0xe9, 0xd8, 0xa5, 0x1d, 0x8c, 0xe7, 0x10, 0x9e, 0x15, 0x45, 0xf8, 0x29, 0x5e, 0x3d, 0x66, 0x3b,
	0xad, 0xe9, 0x90, 0xa7, 0xb0, 0xd3, 0x6a, 0x1b, 0x7e, 0x8e, 0x57, 0x0e, 0x4e, 0xb7, 0xc5, 0x92,
	0x1c, 0xb4, 0xe4, 0xf4, 0xa5, 0x9d, 0x8f, 0xe1, 0x31, 0x2f, 0x68, 0xee, 0xcc, 0x96, 0x52, 0xd9,
	0x5f, 0xfa, 0x98, 0xbe, 0xc6, 0x17, 0xfd, 0x9a, 0x5e, 0xc3, 0xc0, 0x28, 0x2a, 0x74, 0x4e, 0x0d,
	0xb2, 0x74, 0x89, 0xeb, 0xdb, 0xea, 0x33, 0xf6, 0xfe, 0xea, 0x1b, 0x73, 0xb4, 0x7f, 0x04, 0x03,
	0x2e, 0x9b, 0x99, 0xcf, 0x9f, 0xb6, 0x57, 0x0f, 0x33, 0xa9, 0xd9, 0xbb, 0x59, 0x9d, 0x5d, 0xf0,
	0xf5, 0x7b, 0xb3, 0x51, 0xbf, 0x30, 0x77, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xf3, 0x31,
	0xa0, 0x3c, 0x05, 0x00, 0x00,
}
