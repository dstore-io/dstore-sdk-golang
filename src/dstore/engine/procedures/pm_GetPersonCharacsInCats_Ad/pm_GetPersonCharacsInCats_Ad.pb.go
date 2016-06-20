// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonCharacsInCats_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersonCharacsInCats_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonCharacsInCats_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonCharacsInCats_Ad

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
	PersonCharacCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
	PersonCharacCategoryIdNull bool                        `protobuf:"varint,1001,opt,name=person_charac_category_id_null,json=personCharacCategoryIdNull" json:"person_charac_category_id_null,omitempty"`
	Country                    *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=country" json:"country,omitempty"`
	CountryNull                bool                        `protobuf:"varint,1002,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacCategoryId
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
	RowId                       int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	DefaultValue                *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	Required                    *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=required" json:"required,omitempty"`
	CharacteristicDescription   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	PersonCharacteristicId      *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	DescriptionForAdmin         *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=description_for_admin,json=descriptionForAdmin" json:"description_for_admin,omitempty"`
	CategoryDescription         *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	KeepPropertiesHistoryInDays *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=keep_properties_history_in_days,json=keepPropertiesHistoryInDays" json:"keep_properties_history_in_days,omitempty"`
	DefValRestrictedByPattern   *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=def_val_restricted_by_pattern,json=defValRestrictedByPattern" json:"def_val_restricted_by_pattern,omitempty"`
	PersonCharacCategoryId      *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDefaultValue() *dstore_values.StringValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *Response_Row) GetRequired() *dstore_values.BooleanValue {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetDescriptionForAdmin() *dstore_values.StringValue {
	if m != nil {
		return m.DescriptionForAdmin
	}
	return nil
}

func (m *Response_Row) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Response_Row) GetKeepPropertiesHistoryInDays() *dstore_values.IntegerValue {
	if m != nil {
		return m.KeepPropertiesHistoryInDays
	}
	return nil
}

func (m *Response_Row) GetDefValRestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.DefValRestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetPersonCharacCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacCategoryId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonCharacsInCats_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonCharacsInCats_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonCharacsInCats_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x6f, 0xd3, 0x3e,
	0x14, 0xd7, 0xd6, 0xff, 0xb6, 0xca, 0xdb, 0x5f, 0x20, 0x0f, 0xa6, 0x2c, 0x15, 0x03, 0x8d, 0x17,
	0xc4, 0x43, 0x86, 0xb8, 0x48, 0x93, 0x78, 0x61, 0xeb, 0xb8, 0x14, 0x69, 0xa1, 0x8a, 0xd0, 0x10,
	0x93, 0x90, 0xe5, 0x26, 0x5e, 0x67, 0xd1, 0xda, 0xc1, 0x76, 0x98, 0xfa, 0xcc, 0x1b, 0x4f, 0xdc,
	0x2f, 0x5f, 0x81, 0x6f, 0x06, 0x9f, 0x82, 0x93, 0xd8, 0xcd, 0xd2, 0x32, 0xd6, 0xb2, 0x97, 0x36,
	0x27, 0x3e, 0xbf, 0xdf, 0xf9, 0x9d, 0x8b, 0x4f, 0xd0, 0xdd, 0x44, 0x1b, 0xa9, 0xd8, 0x06, 0x13,
	0x5d, 0x2e, 0xd8, 0x46, 0xaa, 0x64, 0xcc, 0x92, 0x4c, 0x31, 0xbd, 0x91, 0xf6, 0xc9, 0x43, 0x66,
	0xda, 0x4c, 0x69, 0x29, 0x9a, 0x87, 0x54, 0xd1, 0x58, 0xb7, 0x44, 0x93, 0x1a, 0x4d, 0xb6, 0x92,
	0x00, 0xfc, 0x8c, 0xc4, 0xd7, 0x2d, 0x38, 0xb0, 0xe0, 0xe0, 0x34, 0x84, 0xbf, 0xec, 0x02, 0xbd,
	0xa6, 0xbd, 0x8c, 0x69, 0x4b, 0xe0, 0xaf, 0x8e, 0x46, 0x67, 0x4a, 0x49, 0xe5, 0x8e, 0x1a, 0xa3,
	0x47, 0x7d, 0xa6, 0x35, 0xed, 0x32, 0x77, 0x78, 0x75, 0xfc, 0xd0, 0x50, 0x2e, 0x0e, 0xa4, 0xea,
	0x53, 0xc3, 0xa5, 0xb0, 0x4e, 0xeb, 0x6f, 0x66, 0x11, 0x6a, 0x83, 0x0a, 0x38, 0x05, 0x49, 0x78,
	0x0f, 0xad, 0xa6, 0x85, 0x34, 0x12, 0x17, 0xda, 0x48, 0x4c, 0x0d, 0xeb, 0x4a, 0x35, 0x20, 0x3c,
	0xf1, 0x66, 0xae, 0xcc, 0x5c, 0x5b, 0xbc, 0xd9, 0x08, 0x5c, 0x42, 0x4e, 0x24, 0x17, 0xe0, 0xc1,
	0xd4, 0x5e, 0x6e, 0x45, 0x2b, 0x69, 0x25, 0xb1, 0xa6, 0xc3, 0xb6, 0x12, 0xdc, 0x44, 0x6b, 0x7f,
	0xe5, 0x25, 0x22, 0xeb, 0xf5, 0xbc, 0x9f, 0x0b, 0xc0, 0x5e, 0x8f, 0xfc, 0x93, 0x09, 0x42, 0x70,
	0xc1, 0xb7, 0xd1, 0x42, 0x2c, 0x33, 0x61, 0xd4, 0xc0, 0x9b, 0x2d, 0xa4, 0xf8, 0x63, 0x52, 0xb4,
	0x51, 0x5c, 0x74, 0xad, 0x92, 0xa1, 0x2b, 0x5e, 0x47, 0x4b, 0xee, 0xd1, 0x06, 0xfa, 0x65, 0x03,
	0x2d, 0xba, 0x97, 0x39, 0xf3, 0xfa, 0x8f, 0x3a, 0xaa, 0x47, 0x4c, 0xa7, 0x52, 0x68, 0x86, 0x6f,
	0xa0, 0xb9, 0xa2, 0xc6, 0x2e, 0xdf, 0x32, 0x88, 0x6b, 0xa0, 0xad, 0xff, 0xfd, 0xfc, 0x37, 0xb2,
	0x8e, 0xf8, 0x39, 0x3a, 0x9f, 0x57, 0x97, 0x54, 0xca, 0x0b, 0x0a, 0x6b, 0x00, 0x0e, 0xc6, 0xc0,
	0xe3, 0x4d, 0xd8, 0x05, 0xbb, 0x75, 0x6c, 0x47, 0xe7, 0xfa, 0xa3, 0x2f, 0xf0, 0x26, 0x5a, 0x70,
	0x5d, 0xf5, 0x6a, 0x05, 0xe3, 0xda, 0x1f, 0x8c, 0xb6, 0xe7, 0xbb, 0xf6, 0x3f, 0x1a, 0xba, 0xe3,
	0xc7, 0xa8, 0xa6, 0xe4, 0x91, 0xf7, 0x5f, 0x81, 0xda, 0x0c, 0xa6, 0x9f, 0xc2, 0x60, 0x58, 0x89,
	0x20, 0x92, 0x47, 0x51, 0x4e, 0xe2, 0xbf, 0x9d, 0x47, 0x35, 0x30, 0xf0, 0x0a, 0x9a, 0x07, 0x33,
	0x9f, 0x85, 0x77, 0x21, 0x14, 0x67, 0x2e, 0x9a, 0x03, 0x13, 0xda, 0x7b, 0x0f, 0xfd, 0x9f, 0xb0,
	0x03, 0x9a, 0xf5, 0x0c, 0x29, 0x5a, 0xe1, 0xbd, 0x0f, 0x27, 0x36, 0x68, 0xc9, 0x21, 0x0a, 0x0b,
	0xf2, 0xac, 0x2b, 0xf6, 0x2a, 0xe3, 0x8a, 0x25, 0xde, 0x87, 0xf0, 0xc4, 0x41, 0xeb, 0x48, 0xd9,
	0x63, 0x54, 0x58, 0x74, 0xe9, 0x8d, 0xf7, 0x91, 0x6f, 0x67, 0x0a, 0x06, 0x98, 0x6b, 0xc3, 0x63,
	0x92, 0x30, 0x1d, 0x2b, 0x9e, 0x16, 0x6d, 0xf8, 0x38, 0x59, 0xc8, 0xea, 0x28, 0x7c, 0xe7, 0x18,
	0x0d, 0xd7, 0xc1, 0x1b, 0x19, 0xdb, 0x32, 0x04, 0x54, 0xe0, 0x53, 0xf8, 0x8f, 0xd7, 0x61, 0x08,
	0x86, 0x7a, 0x3d, 0x41, 0x17, 0x2b, 0x22, 0x09, 0xb4, 0x9b, 0xd0, 0xa4, 0xcf, 0x85, 0xf7, 0x79,
	0xb2, 0xdc, 0xe5, 0x0a, 0xf2, 0x81, 0x54, 0x5b, 0x39, 0x0e, 0x87, 0xe8, 0x42, 0x79, 0xa3, 0xaa,
	0xe9, 0x7f, 0x99, 0x82, 0x6f, 0x08, 0xac, 0x26, 0xde, 0x41, 0x97, 0x5f, 0x32, 0x96, 0x12, 0x58,
	0x12, 0x90, 0x82, 0xe1, 0x4c, 0x93, 0x43, 0x9e, 0x53, 0xc0, 0x85, 0x15, 0x24, 0xa1, 0x03, 0xed,
	0x7d, 0x9d, 0x22, 0xff, 0x46, 0x4e, 0xd2, 0x2e, 0x39, 0x1e, 0x59, 0x8a, 0x96, 0xd8, 0x01, 0x02,
	0xfc, 0x02, 0x5d, 0x82, 0x11, 0xc8, 0x07, 0x86, 0xc0, 0x2a, 0x05, 0x45, 0x50, 0xa1, 0x84, 0x74,
	0x06, 0x24, 0xa5, 0x06, 0x6a, 0x25, 0xbc, 0x6f, 0x53, 0xf4, 0x0e, 0x18, 0xe0, 0x29, 0x2a, 0xf1,
	0xdb, 0x83, 0xb6, 0x45, 0xe3, 0x67, 0xa7, 0xad, 0xb2, 0xef, 0xe1, 0x99, 0x77, 0xd9, 0xf6, 0x53,
	0xd4, 0xe0, 0x72, 0xfc, 0x3e, 0x95, 0x9f, 0x84, 0xfd, 0x3b, 0x67, 0xfa, 0x58, 0x74, 0xe6, 0x8b,
	0x7d, 0x7c, 0xeb, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xe1, 0xd6, 0x23, 0x6c, 0x06, 0x00,
	0x00,
}
