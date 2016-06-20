// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersons_Conditions_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersons_Conditions_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersons_Conditions_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersons_Conditions_Ad

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
	SortOrder1                   *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=sort_order1,json=sortOrder1" json:"sort_order1,omitempty"`
	SortOrder1Null               bool                        `protobuf:"varint,1001,opt,name=sort_order1_null,json=sortOrder1Null" json:"sort_order1_null,omitempty"`
	SortOrder2                   *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=sort_order2,json=sortOrder2" json:"sort_order2,omitempty"`
	SortOrder2Null               bool                        `protobuf:"varint,1002,opt,name=sort_order2_null,json=sortOrder2Null" json:"sort_order2_null,omitempty"`
	SortOrder3                   *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=sort_order3,json=sortOrder3" json:"sort_order3,omitempty"`
	SortOrder3Null               bool                        `protobuf:"varint,1003,opt,name=sort_order3_null,json=sortOrder3Null" json:"sort_order3_null,omitempty"`
	OutputCharacteristicIds      *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=output_characteristic_ids,json=outputCharacteristicIds" json:"output_characteristic_ids,omitempty"`
	OutputCharacteristicIdsNull  bool                        `protobuf:"varint,1004,opt,name=output_characteristic_ids_null,json=outputCharacteristicIdsNull" json:"output_characteristic_ids_null,omitempty"`
	PersonTypeId                 *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull             bool                        `protobuf:"varint,1005,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	LastValues                   *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=last_values,json=lastValues" json:"last_values,omitempty"`
	LastValuesNull               bool                        `protobuf:"varint,1006,opt,name=last_values_null,json=lastValuesNull" json:"last_values_null,omitempty"`
	LastPersonId                 *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=last_person_id,json=lastPersonId" json:"last_person_id,omitempty"`
	LastPersonIdNull             bool                        `protobuf:"varint,1007,opt,name=last_person_id_null,json=lastPersonIdNull" json:"last_person_id_null,omitempty"`
	Next                         *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=next" json:"next,omitempty"`
	NextNull                     bool                        `protobuf:"varint,1008,opt,name=next_null,json=nextNull" json:"next_null,omitempty"`
	RowNumber                    *dstore_values.IntegerValue `protobuf:"bytes,9,opt,name=row_number,json=rowNumber" json:"row_number,omitempty"`
	RowNumberNull                bool                        `protobuf:"varint,1009,opt,name=row_number_null,json=rowNumberNull" json:"row_number_null,omitempty"`
	GroupId                      *dstore_values.IntegerValue `protobuf:"bytes,10,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	GroupIdNull                  bool                        `protobuf:"varint,1010,opt,name=group_id_null,json=groupIdNull" json:"group_id_null,omitempty"`
	CharacteristicIdList         *dstore_values.StringValue  `protobuf:"bytes,11,opt,name=characteristic_id_list,json=characteristicIdList" json:"characteristic_id_list,omitempty"`
	CharacteristicIdListNull     bool                        `protobuf:"varint,1011,opt,name=characteristic_id_list_null,json=characteristicIdListNull" json:"characteristic_id_list_null,omitempty"`
	ConditionList                *dstore_values.StringValue  `protobuf:"bytes,12,opt,name=condition_list,json=conditionList" json:"condition_list,omitempty"`
	ConditionListNull            bool                        `protobuf:"varint,1012,opt,name=condition_list_null,json=conditionListNull" json:"condition_list_null,omitempty"`
	Country                      *dstore_values.StringValue  `protobuf:"bytes,13,opt,name=country" json:"country,omitempty"`
	CountryNull                  bool                        `protobuf:"varint,1013,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	OutputIntoOneId              *dstore_values.IntegerValue `protobuf:"bytes,14,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull          bool                        `protobuf:"varint,1014,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
	SeparatorInLastValues        *dstore_values.StringValue  `protobuf:"bytes,15,opt,name=separator_in_last_values,json=separatorInLastValues" json:"separator_in_last_values,omitempty"`
	SeparatorInLastValuesNull    bool                        `protobuf:"varint,1015,opt,name=separator_in_last_values_null,json=separatorInLastValuesNull" json:"separator_in_last_values_null,omitempty"`
	SeparatorInConditionList     *dstore_values.StringValue  `protobuf:"bytes,16,opt,name=separator_in_condition_list,json=separatorInConditionList" json:"separator_in_condition_list,omitempty"`
	SeparatorInConditionListNull bool                        `protobuf:"varint,1016,opt,name=separator_in_condition_list_null,json=separatorInConditionListNull" json:"separator_in_condition_list_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSortOrder1() *dstore_values.BooleanValue {
	if m != nil {
		return m.SortOrder1
	}
	return nil
}

func (m *Parameters) GetSortOrder2() *dstore_values.BooleanValue {
	if m != nil {
		return m.SortOrder2
	}
	return nil
}

func (m *Parameters) GetSortOrder3() *dstore_values.BooleanValue {
	if m != nil {
		return m.SortOrder3
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicIds() *dstore_values.StringValue {
	if m != nil {
		return m.OutputCharacteristicIds
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetLastValues() *dstore_values.StringValue {
	if m != nil {
		return m.LastValues
	}
	return nil
}

func (m *Parameters) GetLastPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LastPersonId
	}
	return nil
}

func (m *Parameters) GetNext() *dstore_values.BooleanValue {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *Parameters) GetRowNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowNumber
	}
	return nil
}

func (m *Parameters) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *Parameters) GetCharacteristicIdList() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList
	}
	return nil
}

func (m *Parameters) GetConditionList() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionList
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Parameters) GetOutputIntoOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputIntoOneId
	}
	return nil
}

func (m *Parameters) GetSeparatorInLastValues() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInLastValues
	}
	return nil
}

func (m *Parameters) GetSeparatorInConditionList() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInConditionList
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Count           *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=count" json:"count,omitempty"`
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

func (m *Response) GetCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.Count
	}
	return nil
}

type Response_Row struct {
	RowId    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PersonId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	Value1   *dstore_values.StringValue  `protobuf:"bytes,20002,opt,name=value1" json:"value1,omitempty"`
	Value2   *dstore_values.StringValue  `protobuf:"bytes,30003,opt,name=value2" json:"value2,omitempty"`
	Value3   *dstore_values.StringValue  `protobuf:"bytes,40002,opt,name=value3" json:"value3,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersons_Conditions_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersons_Conditions_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersons_Conditions_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 932 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcb, 0x8f, 0xdb, 0x44,
	0x18, 0xd7, 0x76, 0x37, 0x9b, 0xec, 0x97, 0x7d, 0xe1, 0x85, 0x76, 0x76, 0x53, 0xaa, 0x6a, 0x7b,
	0x80, 0x5e, 0x1c, 0x36, 0x59, 0x5e, 0x05, 0x21, 0x2d, 0x0b, 0x2a, 0x91, 0xda, 0x74, 0xe5, 0xa2,
	0x4a, 0x70, 0xb1, 0xbc, 0xf1, 0x10, 0x2c, 0x25, 0x1e, 0x6b, 0xc6, 0xa1, 0xf4, 0xbf, 0x80, 0x3b,
	0x27, 0xfe, 0x05, 0x2e, 0x9c, 0x39, 0xf3, 0x4f, 0x70, 0xa4, 0xbc, 0xdf, 0x88, 0x13, 0xdf, 0xcc,
	0x37, 0x4e, 0x62, 0x27, 0xa9, 0xb3, 0xbd, 0x24, 0xb6, 0xe7, 0xf7, 0x9a, 0xcf, 0x33, 0xfe, 0x06,
	0x6e, 0x85, 0x2a, 0x15, 0x92, 0x37, 0x79, 0xdc, 0x8f, 0x62, 0xde, 0x4c, 0xa4, 0xe8, 0xf1, 0x70,
	0x24, 0xb9, 0x6a, 0x26, 0x43, 0xff, 0x36, 0x4f, 0xcf, 0xb8, 0x54, 0x22, 0x56, 0xfe, 0xa9, 0x88,
	0xc3, 0x28, 0x8d, 0xf4, 0xe5, 0x49, 0xe8, 0x22, 0x2c, 0x15, 0xce, 0x4d, 0xe2, 0xba, 0xc4, 0x75,
	0x9f, 0x40, 0x38, 0xd8, 0xb3, 0x36, 0x9f, 0x04, 0x83, 0x11, 0x57, 0xc4, 0x3f, 0xd8, 0xcf, 0x7b,
	0x73, 0x29, 0x85, 0xb4, 0x43, 0x8d, 0xfc, 0xd0, 0x90, 0x2b, 0x15, 0xf4, 0xb9, 0x1d, 0xbc, 0x51,
	0x1c, 0x4c, 0x83, 0x28, 0xfe, 0x48, 0xc8, 0x61, 0xa0, 0xfd, 0x08, 0x74, 0xf8, 0xdd, 0x36, 0xc0,
	0x59, 0x20, 0x03, 0x1c, 0xc5, 0x44, 0xce, 0x9b, 0x50, 0x57, 0x42, 0xa6, 0xbe, 0x90, 0x21, 0x97,
	0x47, 0x6c, 0xe5, 0xfa, 0xca, 0x8b, 0xf5, 0x56, 0xc3, 0xb5, 0x33, 0xb0, 0xb1, 0xce, 0x85, 0x18,
	0xf0, 0x20, 0x7e, 0xa0, 0xef, 0x3c, 0xd0, 0xf8, 0x7b, 0x06, 0xee, 0xdc, 0x84, 0xdd, 0x29, 0xb6,
	0x1f, 0x8f, 0x06, 0x03, 0xf6, 0x7d, 0x15, 0x35, 0x6a, 0xde, 0xf6, 0x04, 0xd6, 0xc5, 0xc7, 0x79,
	0xa3, 0x16, 0xbb, 0x74, 0x11, 0xa3, 0x56, 0xde, 0xa8, 0x45, 0x46, 0x8f, 0x8b, 0x46, 0xad, 0x59,
	0xa3, 0x36, 0x5b, 0xbd, 0x88, 0x51, 0x3b, 0x6f, 0xd4, 0x26, 0xa3, 0x1f, 0x8a, 0x46, 0x6d, 0x63,
	0xf4, 0x00, 0xf6, 0xc5, 0x28, 0x4d, 0x46, 0xa9, 0xdf, 0xfb, 0x18, 0x0b, 0xda, 0xc3, 0x7a, 0x46,
	0x2a, 0x8d, 0x7a, 0x7e, 0x14, 0x2a, 0xb6, 0x66, 0x6c, 0x0f, 0x0a, 0xb6, 0x2a, 0x95, 0x51, 0xdc,
	0x27, 0xd7, 0x2b, 0x44, 0x3e, 0xcd, 0x71, 0x3b, 0xa1, 0x72, 0xde, 0x81, 0x6b, 0x0b, 0x75, 0x29,
	0xd0, 0x8f, 0x14, 0xa8, 0xb1, 0x40, 0xc1, 0xa4, 0x3b, 0x81, 0xed, 0xc4, 0x2c, 0x39, 0x3f, 0x7d,
	0x94, 0x70, 0xe4, 0xb2, 0xca, 0xdc, 0x4a, 0x44, 0x71, 0xca, 0xfb, 0x5c, 0x52, 0xa6, 0x4d, 0xa2,
	0xbc, 0x8f, 0x8c, 0x4e, 0xe8, 0xb8, 0xb0, 0x97, 0x97, 0x20, 0xf7, 0x9f, 0xc8, 0x7d, 0x77, 0x1a,
	0x6b, 0x2c, 0xdf, 0x80, 0xfa, 0x20, 0x50, 0xa9, 0x4f, 0xca, 0x6c, 0xbd, 0xb4, 0x04, 0xa0, 0xe1,
	0xe6, 0x52, 0xe9, 0xc2, 0x4f, 0x91, 0xc9, 0xe9, 0x67, 0x5b, 0xf8, 0x09, 0x2c, 0x9b, 0x9a, 0x81,
	0xda, 0x70, 0x38, 0xb5, 0xea, 0x12, 0x53, 0xd3, 0x14, 0xda, 0x84, 0x34, 0xb5, 0xbc, 0x04, 0x19,
	0xfe, 0x62, 0xa7, 0x36, 0x8d, 0x35, 0x96, 0x4d, 0x58, 0x8b, 0xf9, 0xa7, 0x29, 0xab, 0x95, 0xaf,
	0x26, 0x03, 0x74, 0xae, 0xc2, 0x86, 0xfe, 0x27, 0xd9, 0x5f, 0x49, 0xb6, 0xa6, 0x9f, 0x18, 0xb9,
	0x5b, 0x00, 0x52, 0x3c, 0xc4, 0xc1, 0xe1, 0x39, 0x97, 0x6c, 0xa3, 0x3c, 0xfd, 0x06, 0xc2, 0xbb,
	0x06, 0xed, 0xbc, 0x00, 0x3b, 0x13, 0x2e, 0xe9, 0xff, 0x46, 0xfa, 0x5b, 0x63, 0x90, 0x31, 0x79,
	0x05, 0x6a, 0x7d, 0x29, 0x46, 0x89, 0x2e, 0x10, 0x94, 0x5b, 0x54, 0x0d, 0x18, 0x6b, 0x73, 0x03,
	0xb6, 0x32, 0x1e, 0xc9, 0xff, 0x4e, 0xf2, 0x75, 0x0b, 0x30, 0xe2, 0x67, 0x70, 0x79, 0x66, 0x75,
	0xfa, 0x03, 0xbc, 0x62, 0xf5, 0xd2, 0xd7, 0xfe, 0x6c, 0xaf, 0xb0, 0x62, 0xef, 0xe0, 0xbf, 0xf3,
	0x16, 0x34, 0xe6, 0x2b, 0x52, 0x88, 0x3f, 0x28, 0x04, 0x9b, 0xc7, 0xcd, 0x56, 0x45, 0x2f, 0xfb,
	0xb6, 0x52, 0x92, 0xcd, 0xd2, 0x24, 0x5b, 0x63, 0x86, 0x89, 0xd0, 0x84, 0xbd, 0xbc, 0x04, 0x59,
	0xff, 0x49, 0xd6, 0xcf, 0xe4, 0xc0, 0xc6, 0xf3, 0x18, 0xaa, 0x3d, 0x31, 0x8a, 0x53, 0xf9, 0x88,
	0x6d, 0x95, 0x9a, 0x65, 0x50, 0xe7, 0x10, 0x36, 0xed, 0x25, 0xe9, 0xff, 0x65, 0xeb, 0x6b, 0x1f,
	0x1a, 0xe5, 0xf7, 0xc0, 0xb1, 0x1f, 0x01, 0x7c, 0x49, 0xc2, 0x17, 0xb1, 0xd9, 0xc2, 0xdb, 0xe5,
	0xaf, 0x71, 0x87, 0x68, 0x1d, 0x64, 0xdd, 0x8b, 0xf5, 0x2e, 0x7e, 0x19, 0xae, 0xcc, 0x2a, 0x91,
	0xf1, 0xdf, 0x64, 0xbc, 0x57, 0xa0, 0x98, 0x00, 0xf7, 0x81, 0x29, 0x9e, 0x60, 0xad, 0xd1, 0x09,
	0x99, 0xfe, 0xf4, 0xce, 0xde, 0x29, 0x9d, 0xeb, 0x73, 0x63, 0x6e, 0x27, 0xbe, 0x33, 0xd9, 0xe4,
	0x27, 0xf0, 0xfc, 0x22, 0x51, 0x4a, 0xf4, 0x0f, 0x25, 0xda, 0x9f, 0x4b, 0x37, 0xb9, 0x3e, 0x80,
	0x46, 0x4e, 0xa2, 0xf0, 0xce, 0x77, 0x4b, 0xa3, 0xb1, 0x29, 0xed, 0xd3, 0xdc, 0xeb, 0xbf, 0x0d,
	0xd7, 0x9f, 0x20, 0x4d, 0x01, 0xff, 0xa5, 0x80, 0x57, 0x17, 0x89, 0xe8, 0x8c, 0x87, 0xdf, 0xae,
	0x41, 0xcd, 0xe3, 0x2a, 0xc1, 0x26, 0xcf, 0x9d, 0x97, 0xa0, 0x62, 0x3a, 0xb8, 0xed, 0xad, 0xe3,
	0x68, 0xf6, 0x74, 0x40, 0xdd, 0xfd, 0x5d, 0xfd, 0xeb, 0x11, 0x10, 0xa7, 0xb8, 0xab, 0x7b, 0xb7,
	0x3f, 0xd5, 0xbc, 0xb1, 0x5f, 0xae, 0x22, 0xd9, 0x2d, 0x90, 0x8b, 0x2d, 0xfe, 0x2e, 0xde, 0x77,
	0x26, 0xf7, 0xde, 0xce, 0x30, 0xff, 0xc0, 0x79, 0x0d, 0xaa, 0xf6, 0xcc, 0x80, 0x8d, 0x51, 0x2b,
	0x5e, 0x9b, 0x51, 0xa4, 0x13, 0xc5, 0x5d, 0xfa, 0xf7, 0x32, 0xb8, 0xd3, 0x81, 0x55, 0xfc, 0xbc,
	0x60, 0x5f, 0xd3, 0xac, 0x57, 0xdd, 0xa5, 0x8f, 0x38, 0x6e, 0x56, 0x08, 0xd7, 0x13, 0x0f, 0x3d,
	0xad, 0xe1, 0x1c, 0x41, 0xc5, 0x2c, 0x75, 0xc6, 0xcb, 0x97, 0x33, 0x21, 0x0f, 0xfe, 0x5b, 0x81,
	0x55, 0xe4, 0x3b, 0x97, 0x61, 0x5d, 0x7f, 0xfc, 0x70, 0x2b, 0x7c, 0xd6, 0x45, 0x72, 0xc5, 0xab,
	0xe0, 0x2d, 0x2e, 0xf2, 0xd7, 0x61, 0x63, 0xd2, 0x0d, 0x3e, 0xef, 0x96, 0xeb, 0xd6, 0x92, 0xac,
	0x15, 0x1c, 0xc3, 0xba, 0x01, 0x1c, 0xb1, 0x2f, 0xbf, 0x58, 0x29, 0x5d, 0x3d, 0x16, 0x3b, 0x66,
	0xb5, 0xd8, 0x57, 0x8f, 0x97, 0x65, 0xb5, 0xc6, 0xac, 0x36, 0xfb, 0xe6, 0xeb, 0x4b, 0x4b, 0xb2,
	0xda, 0x6f, 0xdf, 0x87, 0x46, 0x24, 0x8a, 0x15, 0x1f, 0x1f, 0x48, 0x3f, 0x3c, 0x7e, 0x9a, 0xa3,
	0xea, 0xf9, 0xba, 0x39, 0x0e, 0xb6, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x27, 0xaf, 0x0c,
	0xe9, 0x0a, 0x00, 0x00,
}