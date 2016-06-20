// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetPostingCharacs_Pu.proto
// DO NOT EDIT!

/*
Package fo_GetPostingCharacs_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetPostingCharacs_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetPostingCharacs_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	ForumId                        *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumIdNull                    bool                        `protobuf:"varint,1004,opt,name=forum_id_null,json=forumIdNull" json:"forum_id_null,omitempty"`
	CharacteristicIdList           *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=characteristic_id_list,json=characteristicIdList" json:"characteristic_id_list,omitempty"`
	CharacteristicIdListNull       bool                        `protobuf:"varint,1005,opt,name=characteristic_id_list_null,json=characteristicIdListNull" json:"characteristic_id_list_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1006,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Parameters) GetCharacteristicIdList() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
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
	CharacteristicDescription   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	FieldTypeId                 *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	PostingCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=posting_characteristic_id,json=postingCharacteristicId" json:"posting_characteristic_id,omitempty"`
	PrecisionValue              *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=precision_value,json=precisionValue" json:"precision_value,omitempty"`
	CommonCharacteristic        *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=common_characteristic,json=commonCharacteristic" json:"common_characteristic,omitempty"`
	MaxLength                   *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=max_length,json=maxLength" json:"max_length,omitempty"`
	PredefinedValues            *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=predefined_values,json=predefinedValues" json:"predefined_values,omitempty"`
	CheckPostingVisibility      *dstore_values.BooleanValue `protobuf:"bytes,10008,opt,name=check_posting_visibility,json=checkPostingVisibility" json:"check_posting_visibility,omitempty"`
	Format                      *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=format" json:"format,omitempty"`
	BasicFieldType              *dstore_values.StringValue  `protobuf:"bytes,10010,opt,name=basic_field_type,json=basicFieldType" json:"basic_field_type,omitempty"`
	FieldTypeDescription        *dstore_values.StringValue  `protobuf:"bytes,10011,opt,name=field_type_description,json=fieldTypeDescription" json:"field_type_description,omitempty"`
	BasicFieldTypeId            *dstore_values.IntegerValue `protobuf:"bytes,10012,opt,name=basic_field_type_id,json=basicFieldTypeId" json:"basic_field_type_id,omitempty"`
	PropertyModificationAllowed *dstore_values.IntegerValue `protobuf:"bytes,10013,opt,name=property_modification_allowed,json=propertyModificationAllowed" json:"property_modification_allowed,omitempty"`
	MaxNumberOfProperties       *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=max_number_of_properties,json=maxNumberOfProperties" json:"max_number_of_properties,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
	}
	return nil
}

func (m *Response_Row) GetPostingCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetPrecisionValue() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrecisionValue
	}
	return nil
}

func (m *Response_Row) GetCommonCharacteristic() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommonCharacteristic
	}
	return nil
}

func (m *Response_Row) GetMaxLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxLength
	}
	return nil
}

func (m *Response_Row) GetPredefinedValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.PredefinedValues
	}
	return nil
}

func (m *Response_Row) GetCheckPostingVisibility() *dstore_values.BooleanValue {
	if m != nil {
		return m.CheckPostingVisibility
	}
	return nil
}

func (m *Response_Row) GetFormat() *dstore_values.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *Response_Row) GetBasicFieldType() *dstore_values.StringValue {
	if m != nil {
		return m.BasicFieldType
	}
	return nil
}

func (m *Response_Row) GetFieldTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.FieldTypeDescription
	}
	return nil
}

func (m *Response_Row) GetBasicFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicFieldTypeId
	}
	return nil
}

func (m *Response_Row) GetPropertyModificationAllowed() *dstore_values.IntegerValue {
	if m != nil {
		return m.PropertyModificationAllowed
	}
	return nil
}

func (m *Response_Row) GetMaxNumberOfProperties() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetPostingCharacs_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetPostingCharacs_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetPostingCharacs_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 890 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0xd5, 0x12, 0x92, 0x66, 0x6f, 0x97, 0x6c, 0x98, 0x76, 0xb3, 0x6e, 0xb2, 0xac, 0x60, 0x57,
	0x2b, 0xf1, 0xe4, 0xae, 0xa8, 0xb4, 0x20, 0x21, 0x01, 0x0b, 0x2c, 0x28, 0xb0, 0x4d, 0x83, 0x05,
	0x15, 0x54, 0x48, 0xc6, 0xb1, 0x27, 0xe9, 0x08, 0xdb, 0x63, 0x3c, 0x76, 0x4b, 0xff, 0x03, 0x0f,
	0x7c, 0x7f, 0xc3, 0xaf, 0xe3, 0x85, 0xcf, 0xbf, 0x00, 0xd7, 0x33, 0x63, 0x27, 0x76, 0x53, 0xec,
	0x7d, 0xa9, 0x3b, 0x99, 0x73, 0xce, 0x9c, 0x99, 0x7b, 0xee, 0xd8, 0x70, 0xcf, 0x13, 0x09, 0x8f,
	0xe9, 0x2e, 0x0d, 0x17, 0x2c, 0xa4, 0xbb, 0x51, 0xcc, 0x5d, 0xea, 0xa5, 0x31, 0x15, 0xbb, 0x73,
	0x6e, 0xbf, 0x49, 0x93, 0x29, 0x17, 0x09, 0x0b, 0x17, 0xaf, 0x1d, 0x3b, 0xb1, 0xe3, 0x0a, 0x7b,
	0x9a, 0x9a, 0x08, 0x49, 0x38, 0xb9, 0xa3, 0x78, 0xa6, 0xe2, 0x99, 0x17, 0x80, 0x87, 0x5b, 0x5a,
	0xfe, 0xc4, 0xf1, 0x53, 0x2a, 0x14, 0x77, 0xb8, 0x53, 0x5e, 0x93, 0xc6, 0x31, 0x8f, 0xf5, 0xd4,
	0xa8, 0x3c, 0x15, 0x50, 0x21, 0x9c, 0x05, 0xd5, 0x93, 0xb7, 0xab, 0x93, 0x89, 0xc3, 0xc2, 0x39,
	0x8f, 0x03, 0x27, 0x61, 0x3c, 0x54, 0xa0, 0x5b, 0x9f, 0x75, 0x00, 0xa6, 0xb8, 0x3e, 0xce, 0xd2,
	0x58, 0x90, 0x0f, 0xe1, 0x46, 0x84, 0x4f, 0x1e, 0xda, 0xcc, 0xa3, 0x61, 0xc2, 0xe6, 0xcc, 0x95,
	0x68, 0x5b, 0x39, 0x32, 0x2e, 0x3d, 0x7d, 0xe9, 0xd9, 0xcd, 0xe7, 0x86, 0xa6, 0xde, 0x8e, 0xf6,
	0x29, 0x92, 0x18, 0xf7, 0x71, 0x98, 0x0d, 0xac, 0xa1, 0xe2, 0x8f, 0x4b, 0x74, 0x39, 0x25, 0xc8,
	0x5b, 0xf0, 0xcc, 0xff, 0xa9, 0xdb, 0x61, 0xea, 0xfb, 0xc6, 0xef, 0x1b, 0xb8, 0x46, 0xd7, 0xba,
	0x79, 0xb1, 0xce, 0x04, 0x61, 0xe4, 0x3e, 0xf4, 0xb4, 0x56, 0x72, 0x16, 0x51, 0x14, 0x34, 0x1e,
	0x93, 0xde, 0x46, 0x15, 0x6f, 0x2c, 0x4c, 0xe8, 0x82, 0xc6, 0xca, 0xdc, 0x15, 0x45, 0x79, 0x17,
	0x19, 0x63, 0x8f, 0x98, 0xb0, 0x55, 0x96, 0x50, 0x06, 0xfe, 0x50, 0x06, 0xfa, 0xab, 0x58, 0xb9,
	0xe4, 0xf3, 0x70, 0x39, 0x0d, 0xd9, 0x27, 0xa9, 0x5c, 0xad, 0x55, 0x7b, 0x12, 0x5d, 0x05, 0xc6,
	0x85, 0xee, 0x40, 0xaf, 0x20, 0xaa, 0x35, 0xfe, 0x54, 0x6b, 0x5c, 0xc9, 0x21, 0x52, 0xff, 0x1e,
	0x74, 0xb1, 0x3c, 0x69, 0x90, 0xc9, 0x3f, 0x5e, 0xbf, 0x99, 0x0d, 0x09, 0x46, 0xf9, 0xdb, 0xf0,
	0x44, 0xce, 0x53, 0xea, 0x7f, 0x29, 0xf5, 0x4d, 0x0d, 0x90, 0xe2, 0x53, 0x18, 0xb8, 0x32, 0x68,
	0x58, 0x67, 0x86, 0xb9, 0x73, 0x33, 0xb4, 0x8f, 0xff, 0x19, 0xed, 0xda, 0x9d, 0x6c, 0x97, 0x99,
	0x63, 0xef, 0x21, 0x3e, 0xc9, 0x4b, 0x30, 0x5a, 0xaf, 0xa8, 0x4c, 0xfc, 0xad, 0x4c, 0x18, 0xeb,
	0xb8, 0xd2, 0xd1, 0x3b, 0x70, 0x5d, 0xd0, 0x08, 0x27, 0x71, 0x59, 0x9b, 0xe9, 0x4c, 0x64, 0x51,
	0x10, 0x46, 0xa7, 0xde, 0x52, 0x41, 0x1d, 0xab, 0x8c, 0xe0, 0xcf, 0x82, 0xbc, 0x0c, 0x37, 0x2e,
	0x90, 0x54, 0x9e, 0xfe, 0xd1, 0x9e, 0xd6, 0x91, 0x33, 0x4f, 0xb7, 0xfe, 0x05, 0xe8, 0x5a, 0x54,
	0x44, 0x3c, 0x14, 0x94, 0xdc, 0x85, 0xb6, 0x6c, 0xb6, 0x6a, 0xea, 0x75, 0x13, 0xab, 0x46, 0x7c,
	0x90, 0xfd, 0xb5, 0x14, 0x90, 0x7c, 0x00, 0xfd, 0xac, 0xcd, 0xec, 0x95, 0x3e, 0xc3, 0x58, 0xb6,
	0x90, 0x6c, 0x56, 0xc8, 0xd5, 0x6e, 0xdc, 0xc7, 0xf1, 0x78, 0x39, 0xb6, 0xae, 0x06, 0xe5, 0x1f,
	0xc8, 0x0b, 0xb0, 0xa1, 0xdb, 0x1b, 0xa3, 0x97, 0x29, 0xde, 0x3c, 0xa7, 0xa8, 0x9a, 0x7f, 0x5f,
	0x3d, 0xad, 0x1c, 0x4e, 0x1e, 0x40, 0x2b, 0xe6, 0xa7, 0x98, 0xa8, 0x8c, 0xb5, 0x67, 0x36, 0xba,
	0x89, 0xcc, 0xfc, 0x10, 0x4c, 0x8b, 0x9f, 0x5a, 0x19, 0x7f, 0xf8, 0x5b, 0x17, 0x5a, 0x38, 0x20,
	0x03, 0xe8, 0xe0, 0x30, 0xcb, 0xe8, 0xe7, 0x13, 0x3c, 0x97, 0xb6, 0xd5, 0xc6, 0x21, 0xa6, 0xf0,
	0x08, 0x86, 0x95, 0x38, 0x78, 0x54, 0xb8, 0x31, 0x8b, 0xe4, 0x29, 0x7c, 0x31, 0xa9, 0x2d, 0xe9,
	0x4e, 0x99, 0xfe, 0xfa, 0x92, 0x4d, 0x5e, 0xc1, 0x84, 0x33, 0xea, 0x7b, 0x45, 0xaf, 0x7f, 0x39,
	0xa9, 0xef, 0x8f, 0x4d, 0x49, 0xd1, 0xbd, 0xfe, 0x3e, 0xec, 0x44, 0x6a, 0x93, 0xf6, 0xb9, 0xd0,
	0x1a, 0x5f, 0x35, 0x50, 0xbb, 0x1e, 0xad, 0x9e, 0xd1, 0x32, 0xce, 0x78, 0xbc, 0x57, 0xa3, 0x98,
	0xba, 0x4c, 0x14, 0x17, 0x99, 0xf1, 0x75, 0x03, 0xbd, 0x5e, 0x41, 0x92, 0x63, 0xec, 0x86, 0x6b,
	0x2e, 0x0f, 0x02, 0xd4, 0x28, 0xfb, 0x33, 0xbe, 0x69, 0x20, 0xb6, 0xad, 0xa8, 0x65, 0x6f, 0xe4,
	0x45, 0x80, 0xc0, 0xf9, 0xd4, 0xf6, 0xb1, 0xd4, 0xc9, 0xb1, 0xf1, 0x6d, 0x03, 0x9d, 0xcb, 0x88,
	0x7f, 0x28, 0xe1, 0x64, 0x0c, 0x4f, 0xa2, 0x43, 0x8f, 0xce, 0x31, 0x25, 0x5e, 0x7e, 0xfd, 0x7f,
	0xb7, 0x5e, 0x63, 0xc6, 0xb9, 0x4f, 0x1d, 0xb5, 0x11, 0xbc, 0x37, 0x0b, 0x9a, 0xbe, 0xf6, 0x0f,
	0x01, 0x2f, 0x01, 0xea, 0x7e, 0x6c, 0xe7, 0x15, 0x38, 0xc1, 0x7d, 0xcf, 0x98, 0xcf, 0x92, 0x33,
	0xe3, 0xfb, 0x06, 0x8a, 0x03, 0xc9, 0xd6, 0x19, 0x3d, 0x2c, 0xb8, 0x64, 0x0f, 0x3a, 0xaa, 0x3f,
	0x8c, 0x1f, 0xea, 0xd3, 0xa5, 0xa1, 0x58, 0xae, 0xfe, 0xcc, 0x11, 0x58, 0xf7, 0x65, 0xa0, 0x8c,
	0x1f, 0xeb, 0xe9, 0x3d, 0x49, 0x7a, 0x23, 0x4f, 0x14, 0x96, 0x6b, 0xb0, 0x92, 0xc8, 0xd5, 0xa4,
	0xff, 0x54, 0x2f, 0xb6, 0x5d, 0x24, 0x73, 0x35, 0xe4, 0x6f, 0xc3, 0x56, 0xd5, 0x59, 0x16, 0xce,
	0x9f, 0x1b, 0xd4, 0xad, 0x5f, 0x76, 0x87, 0xa9, 0xfc, 0x08, 0x9e, 0xc2, 0x17, 0x3c, 0xbe, 0xc2,
	0x92, 0x33, 0x3b, 0xe0, 0xde, 0xf2, 0x55, 0xeb, 0xf8, 0x3e, 0x3f, 0xa5, 0x9e, 0xf1, 0x4b, 0x03,
	0xd9, 0x51, 0x2e, 0xb1, 0xbf, 0xa2, 0x70, 0x5f, 0x09, 0x90, 0xf7, 0xc0, 0xc8, 0xd2, 0x15, 0xa6,
	0xc1, 0x8c, 0xc6, 0x36, 0x9f, 0xdb, 0x1a, 0xcc, 0x30, 0x27, 0xbf, 0x36, 0x10, 0xbf, 0x86, 0xec,
	0x89, 0x24, 0x1f, 0xcc, 0xa7, 0x05, 0xf5, 0xd5, 0x03, 0x18, 0x31, 0x5e, 0xb9, 0xa4, 0x96, 0x9f,
	0x59, 0x47, 0x77, 0x1f, 0xf5, 0x03, 0x6c, 0xd6, 0x91, 0x1f, 0x3a, 0x7b, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0xab, 0xb8, 0x64, 0xbb, 0x09, 0x00, 0x00,
}
