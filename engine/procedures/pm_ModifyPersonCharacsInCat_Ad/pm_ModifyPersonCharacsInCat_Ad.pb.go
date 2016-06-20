// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonCharacsInCat_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonCharacsInCat_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonCharacsInCat_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonCharacsInCat_Ad

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
	PersonCharacCategoryId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_charac_category_id,json=personCharacCategoryId" json:"person_charac_category_id,omitempty"`
	PersonCharacCategoryIdNull      bool                        `protobuf:"varint,1001,opt,name=person_charac_category_id_null,json=personCharacCategoryIdNull" json:"person_charac_category_id_null,omitempty"`
	PersonCharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull      bool                        `protobuf:"varint,1002,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
	Required                        *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=required" json:"required,omitempty"`
	RequiredNull                    bool                        `protobuf:"varint,1003,opt,name=required_null,json=requiredNull" json:"required_null,omitempty"`
	Delete                          *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull                      bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
	KeepPropertiesHistoryInDays     *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=keep_properties_history_in_days,json=keepPropertiesHistoryInDays" json:"keep_properties_history_in_days,omitempty"`
	KeepPropertiesHistoryInDaysNull bool                        `protobuf:"varint,1005,opt,name=keep_properties_history_in_days_null,json=keepPropertiesHistoryInDaysNull" json:"keep_properties_history_in_days_null,omitempty"`
	DefaultValue                    *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	DefaultValueNull                bool                        `protobuf:"varint,1006,opt,name=default_value_null,json=defaultValueNull" json:"default_value_null,omitempty"`
	Country                         *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=country" json:"country,omitempty"`
	CountryNull                     bool                        `protobuf:"varint,1007,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
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

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Parameters) GetRequired() *dstore_values.BooleanValue {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Parameters) GetKeepPropertiesHistoryInDays() *dstore_values.IntegerValue {
	if m != nil {
		return m.KeepPropertiesHistoryInDays
	}
	return nil
}

func (m *Parameters) GetDefaultValue() *dstore_values.StringValue {
	if m != nil {
		return m.DefaultValue
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonCharacsInCat_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonCharacsInCat_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonCharacsInCat_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonCharacsInCat_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xd5, 0x86, 0x24, 0xd5, 0xb4, 0x15, 0x95, 0x91, 0xaa, 0x6d, 0xa2, 0x7e, 0xa8, 0xf4,
	0xc0, 0x01, 0xb6, 0x88, 0x22, 0x81, 0x90, 0x10, 0xd0, 0x80, 0x44, 0x24, 0x5a, 0x55, 0x2b, 0x51,
	0x09, 0x2e, 0x96, 0x1b, 0xbb, 0xa9, 0xd5, 0x8d, 0xbd, 0xd8, 0xbb, 0x54, 0x79, 0x0b, 0x9e, 0x88,
	0xf7, 0xe1, 0xfb, 0xc8, 0x15, 0xaf, 0xc7, 0x9b, 0x36, 0xa1, 0x6d, 0xd4, 0x4b, 0xe2, 0xf1, 0xcc,
	0xfc, 0xe6, 0x6f, 0xaf, 0x67, 0xe0, 0x39, 0xb7, 0xb9, 0x36, 0x62, 0x5b, 0xa8, 0xbe, 0x54, 0x62,
	0x3b, 0x33, 0xba, 0x27, 0x78, 0x61, 0x84, 0xdd, 0xce, 0x06, 0x74, 0x4f, 0x73, 0x79, 0x3c, 0x3c,
	0x10, 0xc6, 0x6a, 0xd5, 0x39, 0x61, 0x86, 0xf5, 0x6c, 0x57, 0x75, 0x58, 0x4e, 0x5f, 0xf1, 0xd8,
	0x45, 0xe6, 0x9a, 0xdc, 0xc7, 0xf4, 0x18, 0xd3, 0xe3, 0xeb, 0x73, 0x5a, 0x77, 0x42, 0xb1, 0xcf,
	0x2c, 0x2d, 0x84, 0x45, 0x44, 0x6b, 0x65, 0x5c, 0x81, 0x30, 0x46, 0x9b, 0xe0, 0x6a, 0x8f, 0xbb,
	0x06, 0xc2, 0x5a, 0xd6, 0x17, 0xc1, 0x79, 0x77, 0xd2, 0x99, 0x33, 0xa9, 0x8e, 0xb5, 0x19, 0xb0,
	0x5c, 0x6a, 0x85, 0x41, 0x9b, 0x7f, 0x1b, 0x00, 0x07, 0x4e, 0x84, 0xf3, 0x3a, 0x45, 0xe4, 0x10,
	0x56, 0x32, 0xaf, 0x8c, 0xf6, 0xbc, 0x34, 0xda, 0x63, 0xb9, 0xe8, 0x6b, 0x33, 0xa4, 0x92, 0x47,
	0x33, 0x1b, 0x33, 0xf7, 0xe6, 0x1f, 0xb5, 0xe3, 0x70, 0xa4, 0x20, 0x52, 0x2a, 0x17, 0x21, 0xcc,
	0x61, 0x69, 0x25, 0xcb, 0xd9, 0x85, 0x73, 0x75, 0x42, 0x6e, 0x97, 0x93, 0x0e, 0xac, 0x5d, 0xc9,
	0xa5, 0xaa, 0x48, 0xd3, 0xe8, 0x5b, 0xd3, 0xd1, 0xe7, 0x92, 0xd6, 0xe5, 0x80, 0x7d, 0x17, 0x42,
	0xde, 0x43, 0x34, 0x06, 0x71, 0x8a, 0xa5, 0xcd, 0x65, 0xaf, 0xd4, 0x36, 0x7b, 0x43, 0x6d, 0x55,
	0xae, 0xd3, 0xb6, 0x0b, 0xab, 0x57, 0x61, 0x51, 0xda, 0xf7, 0x4b, 0xa4, 0x9d, 0xe7, 0x7b, 0x69,
	0x4f, 0x60, 0xce, 0x88, 0x4f, 0x85, 0x34, 0x82, 0x47, 0xb5, 0x4b, 0xa5, 0x1c, 0x69, 0x9d, 0x0a,
	0xa6, 0x50, 0xca, 0x28, 0x98, 0x6c, 0xc1, 0x62, 0xb5, 0xc6, 0x62, 0x3f, 0xb0, 0xd8, 0x42, 0xb5,
	0xeb, 0xf1, 0x3b, 0xd0, 0xe0, 0x22, 0x75, 0x9f, 0x28, 0xba, 0x35, 0x1d, 0x1e, 0x42, 0xc9, 0x06,
	0xcc, 0xe3, 0x0a, 0xc1, 0x3f, 0x11, 0x0c, 0xb8, 0xe7, 0xb1, 0x0c, 0xd6, 0x4f, 0x85, 0xc8, 0xa8,
	0x7b, 0x0a, 0xee, 0x6c, 0xb9, 0x14, 0x96, 0x9e, 0xc8, 0x12, 0xec, 0x3e, 0x8b, 0xa2, 0x9c, 0x0d,
	0x6d, 0x54, 0x9f, 0x7e, 0xaf, 0xed, 0x92, 0x71, 0x30, 0x42, 0xbc, 0x45, 0x42, 0x57, 0xbd, 0x76,
	0xf9, 0x64, 0x1f, 0xb6, 0xa6, 0x94, 0x40, 0x75, 0xbf, 0x50, 0xdd, 0xfa, 0x35, 0x2c, 0x2f, 0xf9,
	0x05, 0x2c, 0x72, 0x71, 0xcc, 0x8a, 0x34, 0xa7, 0x5e, 0x4b, 0xd4, 0xf0, 0x02, 0x5b, 0x13, 0x02,
	0x6d, 0x6e, 0xa4, 0xea, 0xa3, 0xbe, 0x85, 0x90, 0xe0, 0x2d, 0xf2, 0x00, 0xc8, 0x18, 0x00, 0xcb,
	0xff, 0xc6, 0xf2, 0x4b, 0x17, 0x43, 0x7d, 0xbd, 0xc7, 0xd0, 0xec, 0xe9, 0x42, 0xe5, 0x66, 0x18,
	0x35, 0xa7, 0x56, 0xaa, 0x42, 0xc9, 0x26, 0x2c, 0x84, 0x25, 0xe2, 0xff, 0x20, 0x7e, 0x3e, 0x6c,
	0x96, 0xe4, 0xcd, 0xaf, 0xb3, 0x30, 0x97, 0x08, 0x9b, 0x69, 0x65, 0x05, 0x79, 0x08, 0x75, 0xdf,
	0xd7, 0xa1, 0xc7, 0x46, 0x45, 0xc2, 0xd8, 0xc0, 0x9e, 0x7f, 0x53, 0xfe, 0x26, 0x18, 0x48, 0x3e,
	0xc0, 0x52, 0xd9, 0xd1, 0xf4, 0x42, 0x4b, 0xbb, 0x26, 0xa8, 0xb9, 0xe4, 0x78, 0x22, 0x79, 0xb2,
	0xf1, 0xf7, 0x9c, 0xdd, 0x3d, 0xb7, 0x93, 0xdb, 0x83, 0xf1, 0x0d, 0xf2, 0x14, 0x9a, 0x61, 0x92,
	0xb8, 0xb7, 0x5c, 0x12, 0xd7, 0xfe, 0x23, 0xe2, 0x9c, 0xd9, 0xc3, 0xff, 0xa4, 0x0a, 0x27, 0xef,
	0xa0, 0x66, 0xf4, 0x99, 0x7b, 0xa4, 0x65, 0xd6, 0xb3, 0xf8, 0x26, 0xb3, 0x2f, 0xae, 0xee, 0x22,
	0x4e, 0xf4, 0x59, 0x52, 0x62, 0x5a, 0xab, 0x50, 0x73, 0x6b, 0xb2, 0x0c, 0x0d, 0x67, 0x95, 0x4d,
	0xfe, 0x65, 0xdf, 0xdd, 0x4e, 0x3d, 0xa9, 0x3b, 0xb3, 0xcb, 0x77, 0x29, 0xb4, 0xa5, 0x9e, 0xac,
	0x31, 0x1a, 0xcf, 0x1f, 0x5f, 0xf6, 0xb5, 0xe5, 0xa7, 0x95, 0x9f, 0xdf, 0x7c, 0x82, 0x1f, 0x35,
	0xfc, 0x88, 0xdc, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x07, 0x46, 0xc7, 0x51, 0x03, 0x06, 0x00,
	0x00,
}