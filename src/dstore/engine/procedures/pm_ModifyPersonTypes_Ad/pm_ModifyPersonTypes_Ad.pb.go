// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonTypes_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonTypes_Ad

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
	PersonTypeId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull          bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	PersonTypeDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_type_description,json=personTypeDescription" json:"person_type_description,omitempty"`
	PersonTypeDescriptionNull bool                        `protobuf:"varint,1002,opt,name=person_type_description_null,json=personTypeDescriptionNull" json:"person_type_description_null,omitempty"`
	CountryId                 *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	CountryIdNull             bool                        `protobuf:"varint,1003,opt,name=country_id_null,json=countryIdNull" json:"country_id_null,omitempty"`
	SortNo                    *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull                bool                        `protobuf:"varint,1004,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	DeletePersonType          *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_person_type,json=deletePersonType" json:"delete_person_type,omitempty"`
	DeletePersonTypeNull      bool                        `protobuf:"varint,1005,opt,name=delete_person_type_null,json=deletePersonTypeNull" json:"delete_person_type_null,omitempty"`
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

func (m *Parameters) GetPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PersonTypeDescription
	}
	return nil
}

func (m *Parameters) GetCountryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountryId
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetDeletePersonType() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeletePersonType
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x9b, 0x26, 0x29, 0x43, 0xa1, 0xd1, 0x16, 0x68, 0x9a, 0x00, 0x82, 0x22, 0x04, 0x27,
	0xa7, 0xa2, 0xa8, 0x42, 0x9c, 0x28, 0xa2, 0x87, 0x1c, 0x12, 0xaa, 0x15, 0x42, 0x82, 0x8b, 0xe5,
	0xc6, 0xd3, 0xc8, 0x92, 0xb3, 0x63, 0xed, 0x3a, 0x54, 0x79, 0x0b, 0xde, 0x86, 0x37, 0xe1, 0x1d,
	0xf8, 0x7b, 0x07, 0x76, 0x3d, 0xdb, 0x38, 0x71, 0x5b, 0x15, 0x2e, 0x49, 0x36, 0xf3, 0xfd, 0x8c,
	0x67, 0xbf, 0x31, 0x1c, 0xc4, 0x26, 0x27, 0x8d, 0x3d, 0x54, 0xe3, 0x44, 0x61, 0x2f, 0xd3, 0x34,
	0xc2, 0x78, 0xaa, 0xd1, 0xf4, 0xb2, 0x49, 0x38, 0xa0, 0x38, 0x39, 0x9d, 0x1d, 0xa3, 0x36, 0xa4,
	0x3e, 0xcc, 0x32, 0x34, 0xe1, 0x61, 0x1c, 0x58, 0x48, 0x4e, 0xe2, 0x29, 0xf3, 0x02, 0xe6, 0x05,
	0x57, 0x80, 0x3b, 0x5b, 0x5e, 0xfe, 0x4b, 0x94, 0x4e, 0xd1, 0x30, 0xb7, 0xb3, 0xb3, 0xec, 0x89,
	0x5a, 0x93, 0xf6, 0xa5, 0xee, 0x72, 0x69, 0x82, 0xc6, 0x44, 0x63, 0xf4, 0xc5, 0x27, 0xd5, 0x62,
	0x1e, 0x25, 0xea, 0x94, 0xf4, 0x24, 0xca, 0x13, 0x52, 0x0c, 0xda, 0xfd, 0xbe, 0x06, 0x70, 0x1c,
	0xe9, 0xc8, 0x56, 0x6d, 0x2b, 0xe2, 0x10, 0x6e, 0x67, 0x45, 0x4b, 0x61, 0x6e, 0x7b, 0x0a, 0x93,
	0xb8, 0xbd, 0xf2, 0x68, 0xe5, 0xf9, 0xcd, 0x17, 0xdd, 0xc0, 0x3f, 0x80, 0xef, 0x2c, 0x51, 0x39,
	0x8e, 0x51, 0x7f, 0x74, 0x27, 0xb9, 0x91, 0xcd, 0x9f, 0xa2, 0x1f, 0x8b, 0x00, 0xb6, 0x96, 0x25,
	0x42, 0x35, 0x4d, 0xd3, 0xf6, 0x8f, 0xa6, 0x15, 0x5a, 0x97, 0xad, 0x45, 0xec, 0xd0, 0x16, 0x84,
	0x84, 0xed, 0x45, 0x7c, 0x8c, 0x66, 0xa4, 0x93, 0xcc, 0xb5, 0xd8, 0x5e, 0x2d, 0xbc, 0x3b, 0x15,
	0x6f, 0x93, 0xeb, 0x44, 0x8d, 0xd9, 0xfa, 0x6e, 0x29, 0xf7, 0xae, 0x24, 0x8a, 0x37, 0x70, 0xff,
	0x0a, 0x4d, 0x6e, 0xe6, 0x27, 0x37, 0xb3, 0x73, 0x29, 0xbb, 0xe8, 0xea, 0x35, 0xc0, 0x88, 0xa6,
	0x2a, 0xd7, 0x33, 0x37, 0x84, 0xda, 0xf5, 0x43, 0xb8, 0xe1, 0xe1, 0x76, 0x02, 0xcf, 0x60, 0xb3,
	0xe4, 0xb2, 0xe1, 0x2f, 0x36, 0xbc, 0x35, 0x07, 0x15, 0x26, 0x2f, 0xa1, 0x69, 0x48, 0xe7, 0xa1,
	0xa2, 0xf6, 0xda, 0xf5, 0x0e, 0x0d, 0x87, 0x1d, 0x92, 0x78, 0x0c, 0x1b, 0x9e, 0xc5, 0xda, 0xbf,
	0x59, 0x1b, 0xb8, 0x5c, 0x08, 0xf7, 0x41, 0xc4, 0x98, 0xda, 0x2b, 0x0d, 0x17, 0xc6, 0xd0, 0xae,
	0x5f, 0xea, 0x71, 0x42, 0x94, 0x62, 0xa4, 0xd8, 0xa3, 0xc5, 0xb4, 0x32, 0x96, 0xe2, 0x00, 0xb6,
	0x2f, 0x4a, 0xb1, 0xf1, 0x1f, 0x36, 0xbe, 0x53, 0xe5, 0xb8, 0x16, 0x76, 0xbf, 0xad, 0xc2, 0xba,
	0x44, 0x93, 0x91, 0x32, 0x28, 0xf6, 0xa0, 0x5e, 0xc4, 0xd6, 0xa7, 0x69, 0x7e, 0xa3, 0x7e, 0x1d,
	0x38, 0xd2, 0x47, 0xee, 0x53, 0x32, 0x50, 0x7c, 0x82, 0x96, 0x0b, 0x6c, 0xb8, 0x90, 0x58, 0x1b,
	0x87, 0x9a, 0x25, 0x07, 0x15, 0x72, 0x35, 0xd7, 0x03, 0x7b, 0xee, 0x97, 0x67, 0xb9, 0x39, 0x59,
	0xfe, 0x43, 0xbc, 0x82, 0xa6, 0x5f, 0x14, 0x7b, 0xaf, 0x4e, 0xf1, 0xe1, 0x05, 0x45, 0x5e, 0xa3,
	0x01, 0x7f, 0xcb, 0x73, 0xb8, 0x38, 0x82, 0x9a, 0xa6, 0x33, 0x7b, 0x57, 0x8e, 0xb5, 0x1f, 0xfc,
	0xd3, 0x4e, 0x07, 0xe7, 0x43, 0x08, 0x24, 0x9d, 0x49, 0xc7, 0xef, 0x3c, 0x80, 0x9a, 0xfd, 0x2d,
	0xee, 0x41, 0xc3, 0x9e, 0x5c, 0xbc, 0xbe, 0x0e, 0xed, 0x58, 0xea, 0xb2, 0x6e, 0x8f, 0xfd, 0xf8,
	0xed, 0x7b, 0xe8, 0x26, 0x54, 0x15, 0x9f, 0xbf, 0x68, 0x3e, 0xef, 0xfd, 0xef, 0x2b, 0xe8, 0xa4,
	0x51, 0xac, 0xfa, 0xfe, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xf4, 0x26, 0x16, 0xbd, 0x04,
	0x00, 0x00,
}
