// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ExportPersonData_Ad.proto
// DO NOT EDIT!

/*
Package pm_ExportPersonData_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ExportPersonData_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ExportPersonData_Ad

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
	PersonTypeId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull                bool                        `protobuf:"varint,1001,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	CharacteristicIdList            *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=characteristic_id_list,json=characteristicIdList" json:"characteristic_id_list,omitempty"`
	CharacteristicIdListNull        bool                        `protobuf:"varint,1002,opt,name=characteristic_id_list_null,json=characteristicIdListNull" json:"characteristic_id_list_null,omitempty"`
	Separator                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=separator" json:"separator,omitempty"`
	SeparatorNull                   bool                        `protobuf:"varint,1003,opt,name=separator_null,json=separatorNull" json:"separator_null,omitempty"`
	CharStringsInPropsToReplace     *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=char_strings_in_props_to_replace,json=charStringsInPropsToReplace" json:"char_strings_in_props_to_replace,omitempty"`
	CharStringsInPropsToReplaceNull bool                        `protobuf:"varint,1004,opt,name=char_strings_in_props_to_replace_null,json=charStringsInPropsToReplaceNull" json:"char_strings_in_props_to_replace_null,omitempty"`
	ReplaceCharStringsInPropsBy     *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=replace_char_strings_in_props_by,json=replaceCharStringsInPropsBy" json:"replace_char_strings_in_props_by,omitempty"`
	ReplaceCharStringsInPropsByNull bool                        `protobuf:"varint,1005,opt,name=replace_char_strings_in_props_by_null,json=replaceCharStringsInPropsByNull" json:"replace_char_strings_in_props_by_null,omitempty"`
	RowCount                        *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=row_count,json=rowCount" json:"row_count,omitempty"`
	RowCountNull                    bool                        `protobuf:"varint,1006,opt,name=row_count_null,json=rowCountNull" json:"row_count_null,omitempty"`
	IncludePersonId                 *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=include_person_id,json=includePersonId" json:"include_person_id,omitempty"`
	IncludePersonIdNull             bool                        `protobuf:"varint,1007,opt,name=include_person_id_null,json=includePersonIdNull" json:"include_person_id_null,omitempty"`
	StringForNoProperty             *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=string_for_no_property,json=stringForNoProperty" json:"string_for_no_property,omitempty"`
	StringForNoPropertyNull         bool                        `protobuf:"varint,1008,opt,name=string_for_no_property_null,json=stringForNoPropertyNull" json:"string_for_no_property_null,omitempty"`
	DateFormat                      *dstore_values.StringValue  `protobuf:"bytes,9,opt,name=date_format,json=dateFormat" json:"date_format,omitempty"`
	DateFormatNull                  bool                        `protobuf:"varint,1009,opt,name=date_format_null,json=dateFormatNull" json:"date_format_null,omitempty"`
	IncludeCreationDate             *dstore_values.BooleanValue `protobuf:"bytes,10,opt,name=include_creation_date,json=includeCreationDate" json:"include_creation_date,omitempty"`
	IncludeCreationDateNull         bool                        `protobuf:"varint,1010,opt,name=include_creation_date_null,json=includeCreationDateNull" json:"include_creation_date_null,omitempty"`
	IncludeLastEditedDate           *dstore_values.BooleanValue `protobuf:"bytes,11,opt,name=include_last_edited_date,json=includeLastEditedDate" json:"include_last_edited_date,omitempty"`
	IncludeLastEditedDateNull       bool                        `protobuf:"varint,1011,opt,name=include_last_edited_date_null,json=includeLastEditedDateNull" json:"include_last_edited_date_null,omitempty"`
	LastEditedInfoForCharacIds      *dstore_values.StringValue  `protobuf:"bytes,12,opt,name=last_edited_info_for_charac_ids,json=lastEditedInfoForCharacIds" json:"last_edited_info_for_charac_ids,omitempty"`
	LastEditedInfoForCharacIdsNull  bool                        `protobuf:"varint,1012,opt,name=last_edited_info_for_charac_ids_null,json=lastEditedInfoForCharacIdsNull" json:"last_edited_info_for_charac_ids_null,omitempty"`
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

func (m *Parameters) GetCharacteristicIdList() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList
	}
	return nil
}

func (m *Parameters) GetSeparator() *dstore_values.StringValue {
	if m != nil {
		return m.Separator
	}
	return nil
}

func (m *Parameters) GetCharStringsInPropsToReplace() *dstore_values.StringValue {
	if m != nil {
		return m.CharStringsInPropsToReplace
	}
	return nil
}

func (m *Parameters) GetReplaceCharStringsInPropsBy() *dstore_values.StringValue {
	if m != nil {
		return m.ReplaceCharStringsInPropsBy
	}
	return nil
}

func (m *Parameters) GetRowCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowCount
	}
	return nil
}

func (m *Parameters) GetIncludePersonId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludePersonId
	}
	return nil
}

func (m *Parameters) GetStringForNoProperty() *dstore_values.StringValue {
	if m != nil {
		return m.StringForNoProperty
	}
	return nil
}

func (m *Parameters) GetDateFormat() *dstore_values.StringValue {
	if m != nil {
		return m.DateFormat
	}
	return nil
}

func (m *Parameters) GetIncludeCreationDate() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeCreationDate
	}
	return nil
}

func (m *Parameters) GetIncludeLastEditedDate() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeLastEditedDate
	}
	return nil
}

func (m *Parameters) GetLastEditedInfoForCharacIds() *dstore_values.StringValue {
	if m != nil {
		return m.LastEditedInfoForCharacIds
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ExportPersonData_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ExportPersonData_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ExportPersonData_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ExportPersonData_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 834 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xeb, 0x6e, 0xe3, 0x44,
	0x14, 0xd6, 0x6e, 0x48, 0x9b, 0x3d, 0xed, 0x66, 0xcb, 0x04, 0x82, 0x37, 0xd1, 0x5e, 0xb4, 0xec,
	0x22, 0xf8, 0xe3, 0xa2, 0x05, 0xa4, 0x95, 0xb8, 0x48, 0xdd, 0xb4, 0x45, 0x91, 0x7a, 0x89, 0x4c,
	0x05, 0x82, 0x3f, 0x23, 0xd7, 0x9e, 0x06, 0x0b, 0xc7, 0x63, 0xcd, 0x38, 0x94, 0x3c, 0x03, 0x7f,
	0x78, 0x1a, 0xde, 0x89, 0xfb, 0xf5, 0x01, 0xf6, 0xcc, 0x9c, 0x71, 0x9a, 0xa4, 0x6e, 0xdd, 0x3f,
	0x89, 0xc7, 0xe7, 0x7c, 0x97, 0xb9, 0xf8, 0x9c, 0x81, 0x8f, 0x62, 0x5d, 0x48, 0x25, 0xb6, 0x45,
	0x36, 0x4e, 0x32, 0xb1, 0x9d, 0x2b, 0x19, 0x89, 0x78, 0xaa, 0x84, 0xde, 0xce, 0x27, 0x7c, 0xef,
	0x87, 0x5c, 0xaa, 0x62, 0x24, 0x94, 0x96, 0xd9, 0x6e, 0x58, 0x84, 0x7c, 0x27, 0xf6, 0x31, 0xa3,
	0x90, 0xec, 0x29, 0xc1, 0x7c, 0x82, 0xf9, 0xd5, 0xb9, 0xbd, 0x8e, 0x23, 0xff, 0x3e, 0x4c, 0xa7,
	0x42, 0x13, 0xb4, 0x77, 0x7f, 0x59, 0x51, 0x28, 0x25, 0x95, 0x0b, 0xf5, 0x97, 0x43, 0x13, 0xa1,
	0x75, 0x38, 0x16, 0x2e, 0xf8, 0xf6, 0x6a, 0xb0, 0x08, 0x93, 0xec, 0x4c, 0xaa, 0x49, 0x58, 0x24,
	0x32, 0xa3, 0xa4, 0x27, 0x3f, 0xde, 0x05, 0x18, 0x85, 0x2a, 0xc4, 0x28, 0x3a, 0x61, 0x3b, 0xd0,
	0xce, 0xad, 0x23, 0x5e, 0xcc, 0x72, 0xc1, 0x93, 0xd8, 0xbb, 0xf5, 0xf8, 0xd6, 0xbb, 0x1b, 0xcf,
	0xfb, 0xbe, 0xf3, 0xef, 0x9c, 0x25, 0x59, 0x21, 0xc6, 0x42, 0x7d, 0x69, 0x46, 0xc1, 0x26, 0x41,
	0x4e, 0x10, 0x31, 0x8c, 0x99, 0x0f, 0x9d, 0x65, 0x0a, 0x9e, 0x4d, 0xd3, 0xd4, 0xfb, 0x65, 0x1d,
	0x89, 0x5a, 0xc1, 0xd6, 0x62, 0xee, 0x11, 0x06, 0xd8, 0x08, 0xba, 0xd1, 0xb7, 0xe8, 0x20, 0x42,
	0x03, 0x89, 0x2e, 0x92, 0xc8, 0x40, 0x52, 0x7c, 0xf2, 0x6e, 0x5b, 0xe9, 0xde, 0x8a, 0xb4, 0x2e,
	0x54, 0x92, 0x8d, 0x49, 0xf9, 0x8d, 0x65, 0xe4, 0x30, 0x3e, 0xc0, 0x7f, 0xf6, 0x19, 0xf4, 0xab,
	0x19, 0xc9, 0xc9, 0xaf, 0xe4, 0xc4, 0xab, 0xc2, 0x5a, 0x47, 0x2f, 0xe0, 0x8e, 0x16, 0x39, 0x06,
	0x51, 0xd6, 0x6b, 0xd4, 0x9a, 0xb8, 0x48, 0x66, 0xef, 0x40, 0x7b, 0x3e, 0x20, 0xb1, 0xdf, 0x48,
	0xec, 0xee, 0xfc, 0xb5, 0x55, 0x08, 0xe1, 0xb1, 0x51, 0xe7, 0x44, 0xa3, 0x79, 0x92, 0x71, 0xdc,
	0x8e, 0x5c, 0xf3, 0x42, 0x72, 0x25, 0xf2, 0x34, 0x8c, 0x84, 0xf7, 0x5a, 0xad, 0xb0, 0x9d, 0xe5,
	0x17, 0x44, 0x31, 0xcc, 0x46, 0x86, 0xe0, 0x44, 0x06, 0x04, 0x67, 0xc7, 0xf0, 0xac, 0x4e, 0x82,
	0x1c, 0xfe, 0x4e, 0x0e, 0x1f, 0x5d, 0x43, 0x56, 0x7a, 0x2e, 0x71, 0xd5, 0xc4, 0xa7, 0x33, 0xaf,
	0x59, 0xef, 0xd9, 0x71, 0x0c, 0x2e, 0xa9, 0xbd, 0x9c, 0x19, 0xcf, 0x75, 0x12, 0xe4, 0xf9, 0x0f,
	0xe7, 0xf9, 0x1a, 0xb2, 0x72, 0x27, 0x95, 0x3c, 0xe7, 0x91, 0x9c, 0x66, 0x85, 0xb7, 0x56, 0x7f,
	0x92, 0x5b, 0x98, 0x3d, 0x30, 0xc9, 0xec, 0x19, 0xb4, 0xe7, 0x48, 0xd2, 0xfc, 0x93, 0x34, 0x37,
	0xcb, 0x14, 0x2b, 0xf0, 0x39, 0xbc, 0x9e, 0x64, 0x51, 0x3a, 0x8d, 0x05, 0x77, 0x87, 0x1e, 0x3f,
	0x99, 0xf5, 0x4a, 0xa1, 0x53, 0x29, 0x53, 0x11, 0x66, 0x24, 0x74, 0xcf, 0xa1, 0xe8, 0xf3, 0xc7,
	0xaf, 0xe6, 0x43, 0xe8, 0x5e, 0x22, 0x22, 0xdd, 0xbf, 0x48, 0xb7, 0xb3, 0x82, 0xb0, 0xf2, 0xc7,
	0xd0, 0xa5, 0x35, 0xe2, 0x67, 0xe6, 0xc0, 0x49, 0xbb, 0x4c, 0x42, 0x15, 0x33, 0xaf, 0x55, 0xbb,
	0x13, 0x1d, 0x1a, 0xec, 0xe3, 0x91, 0x94, 0x23, 0x07, 0x63, 0x9f, 0x42, 0xbf, 0x9a, 0x90, 0xbc,
	0xfc, 0x4d, 0x5e, 0xde, 0xaa, 0x80, 0x5a, 0x3f, 0x1f, 0xc3, 0x46, 0x1c, 0x16, 0x82, 0x53, 0x95,
	0xf1, 0xee, 0xd4, 0x9a, 0x00, 0x93, 0xbe, 0x6f, 0xb3, 0xd9, 0x7b, 0xb0, 0xb5, 0x00, 0x26, 0xc1,
	0x7f, 0x48, 0xb0, 0x7d, 0x91, 0xe6, 0xe6, 0xfd, 0x66, 0xb9, 0x5a, 0x91, 0x12, 0xb6, 0x9e, 0x71,
	0x93, 0xe2, 0x41, 0xfd, 0xd2, 0x97, 0x0b, 0x39, 0x70, 0x40, 0xac, 0xbd, 0x82, 0x7d, 0x02, 0xbd,
	0x4a, 0x42, 0x72, 0xf1, 0xaf, 0x9b, 0x76, 0x05, 0xd2, 0xda, 0x39, 0x01, 0xaf, 0x44, 0xa7, 0x21,
	0x96, 0x19, 0x11, 0x27, 0x85, 0x88, 0xc9, 0xd1, 0x46, 0xbd, 0xa3, 0x72, 0x2e, 0x07, 0x88, 0xdd,
	0xb3, 0x50, 0xeb, 0x69, 0x07, 0x1e, 0x5c, 0xc5, 0x4a, 0xb6, 0xfe, 0x23, 0x5b, 0xf7, 0x2b, 0xe1,
	0xd6, 0x18, 0x87, 0x47, 0x8b, 0x50, 0x53, 0xfe, 0xed, 0xc6, 0x52, 0xe9, 0xc3, 0x23, 0xa6, 0xbd,
	0xcd, 0xda, 0x3d, 0xea, 0xa5, 0x73, 0xe2, 0x21, 0x12, 0xe0, 0x36, 0x0c, 0x2c, 0x7c, 0x18, 0x6b,
	0x76, 0x08, 0x4f, 0x6b, 0x04, 0xc8, 0xea, 0xff, 0x64, 0xf5, 0xe1, 0xd5, 0x54, 0xc6, 0xef, 0x93,
	0x9f, 0x6f, 0x43, 0x2b, 0x10, 0x3a, 0x97, 0x99, 0x16, 0xec, 0x7d, 0x68, 0xda, 0x5e, 0xe7, 0x5a,
	0xd0, 0xdc, 0xa2, 0x6b, 0xa1, 0xd4, 0x07, 0xf7, 0xcc, 0x6f, 0x40, 0x89, 0xec, 0x6b, 0xd8, 0x32,
	0x5d, 0x8e, 0x2f, 0xb4, 0x39, 0x6c, 0x22, 0x0d, 0x04, 0xfb, 0x2b, 0xe0, 0xd5, 0x66, 0x78, 0x88,
	0xe3, 0xe1, 0xc5, 0x38, 0xb8, 0x37, 0x59, 0x7e, 0x81, 0x95, 0x64, 0xdd, 0x75, 0x57, 0xec, 0x08,
	0x86, 0xf1, 0xe1, 0x25, 0x46, 0xea, 0xbd, 0x87, 0xf4, 0x1f, 0x94, 0xe9, 0x6c, 0x17, 0x1a, 0x58,
	0x32, 0xb0, 0x9c, 0x1b, 0xd4, 0x73, 0xff, 0x26, 0xf7, 0x00, 0xbf, 0x5c, 0x03, 0x3f, 0x90, 0xe7,
	0x81, 0x81, 0xf7, 0x1e, 0x40, 0x03, 0x9f, 0x59, 0x17, 0xd6, 0x4c, 0x59, 0xc2, 0x22, 0xf3, 0xd3,
	0x11, 0xae, 0x4a, 0x33, 0x68, 0xe2, 0x70, 0x18, 0xbf, 0xfc, 0x0a, 0xfa, 0x89, 0x5c, 0xe5, 0x9e,
	0x5f, 0x4d, 0xbe, 0x79, 0x31, 0x96, 0x3a, 0xfe, 0xae, 0x8c, 0xc7, 0x37, 0xbf, 0xbd, 0x9c, 0xae,
	0xd9, 0x6b, 0xc2, 0x07, 0xaf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x1f, 0xbe, 0xff, 0xf7, 0x08,
	0x00, 0x00,
}