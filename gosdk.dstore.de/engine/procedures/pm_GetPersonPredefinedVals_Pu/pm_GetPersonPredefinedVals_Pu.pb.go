// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonPredefinedVals_Pu.proto
// DO NOT EDIT!

/*
Package pm_GetPersonPredefinedVals_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonPredefinedVals_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonPredefinedVals_Pu

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
	PersonCharacteristicId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull   bool                        `protobuf:"varint,1001,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
	LanguageId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull               bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	OnlyValuesCurrentlyValid     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=only_values_currently_valid,json=onlyValuesCurrentlyValid" json:"only_values_currently_valid,omitempty"`
	OnlyValuesCurrentlyValidNull bool                        `protobuf:"varint,1003,opt,name=only_values_currently_valid_null,json=onlyValuesCurrentlyValidNull" json:"only_values_currently_valid_null,omitempty"`
	GetReferencesForTableId      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=get_references_for_table_id,json=getReferencesForTableId" json:"get_references_for_table_id,omitempty"`
	GetReferencesForTableIdNull  bool                        `protobuf:"varint,1004,opt,name=get_references_for_table_id_null,json=getReferencesForTableIdNull" json:"get_references_for_table_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetOnlyValuesCurrentlyValid() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyValuesCurrentlyValid
	}
	return nil
}

func (m *Parameters) GetGetReferencesForTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetReferencesForTableId
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
	RowId                    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TableId                  *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	ValueRestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value_restricted_by_pattern,json=valueRestrictedByPattern" json:"value_restricted_by_pattern,omitempty"`
	Value                    *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value" json:"value,omitempty"`
	IsCurrentlyValid         *dstore_values.BooleanValue `protobuf:"bytes,10004,opt,name=is_currently_valid,json=isCurrentlyValid" json:"is_currently_valid,omitempty"`
	TableKeyId               *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=table_key_id,json=tableKeyId" json:"table_key_id,omitempty"`
	ValueId                  *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableId
	}
	return nil
}

func (m *Response_Row) GetValueRestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.ValueRestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetIsCurrentlyValid() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsCurrentlyValid
	}
	return nil
}

func (m *Response_Row) GetTableKeyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableKeyId
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonPredefinedVals_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonPredefinedVals_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonPredefinedVals_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_GetPersonPredefinedVals_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0xa6, 0xa6, 0x49, 0xca, 0xa9, 0x68, 0x19, 0xa1, 0xae, 0x89, 0x4a, 0xa9, 0x2f, 0x8a, 0xb0,
	0xf5, 0x82, 0xa8, 0x50, 0x14, 0x5a, 0x6c, 0xa9, 0xa5, 0x21, 0x2c, 0x5a, 0xb0, 0x0f, 0x0e, 0x9b,
	0xec, 0xe9, 0xba, 0x74, 0x33, 0x13, 0x66, 0x26, 0x96, 0xbc, 0xfa, 0x0b, 0xbc, 0xfb, 0xd7, 0x7c,
	0xf5, 0xd1, 0xcb, 0x8f, 0xf0, 0xcc, 0xcc, 0x26, 0x6d, 0xd2, 0x4b, 0xda, 0x97, 0x24, 0x33, 0xe7,
	0x7c, 0xdf, 0xf9, 0xf6, 0x3b, 0xe7, 0x64, 0x61, 0x39, 0xd1, 0x46, 0x2a, 0x5c, 0x42, 0x91, 0x66,
	0x02, 0x97, 0xba, 0x4a, 0xb6, 0x31, 0xe9, 0x29, 0xd4, 0x4b, 0xdd, 0x0e, 0x5f, 0x47, 0xd3, 0x44,
	0xa5, 0xa5, 0x68, 0x2a, 0x4c, 0x70, 0x97, 0x12, 0x92, 0xed, 0x38, 0xd7, 0xbc, 0xd9, 0x0b, 0x29,
	0xd1, 0x48, 0x76, 0xd7, 0xa3, 0x43, 0x8f, 0x0e, 0x4f, 0x85, 0xd4, 0xae, 0x14, 0xa5, 0xde, 0xc7,
	0x79, 0x0f, 0xb5, 0x67, 0xa8, 0x5d, 0x1b, 0xad, 0x8f, 0x4a, 0x49, 0x55, 0x84, 0xea, 0xa3, 0xa1,
	0x0e, 0x6a, 0x1d, 0xa7, 0x58, 0x04, 0x6f, 0x8d, 0x07, 0x4d, 0x9c, 0x89, 0x5d, 0xa9, 0x3a, 0xb1,
	0xc9, 0xa4, 0xf0, 0x49, 0x8b, 0x3f, 0xa7, 0x01, 0x9a, 0xb1, 0x8a, 0x29, 0x4a, 0x9a, 0xd8, 0x6b,
	0x08, 0xba, 0x4e, 0x1b, 0x6f, 0xbf, 0xa3, 0xdb, 0x36, 0x5d, 0x66, 0xda, 0x64, 0x6d, 0x9e, 0x25,
	0xc1, 0xd4, 0xc2, 0xd4, 0xed, 0xd9, 0x07, 0xf5, 0xb0, 0x78, 0xa0, 0x42, 0x63, 0x26, 0x0c, 0xa6,
	0xa8, 0xb6, 0xed, 0x29, 0x9a, 0xf7, 0xe0, 0xd5, 0x11, 0xec, 0x46, 0xc2, 0x56, 0xe0, 0xc6, 0x49,
	0xb4, 0x5c, 0xf4, 0xf2, 0x3c, 0xf8, 0x5d, 0x25, 0xf2, 0x99, 0xa8, 0x76, 0x3c, 0xbe, 0x41, 0x29,
	0x6c, 0x19, 0x66, 0xf3, 0x58, 0xa4, 0x3d, 0x7a, 0x40, 0xab, 0xe6, 0xc2, 0x64, 0x35, 0x30, 0xc8,
	0x27, 0x05, 0x77, 0x60, 0xee, 0x10, 0xda, 0x17, 0xfd, 0xe3, 0x8b, 0x5e, 0x3a, 0x48, 0x73, 0x85,
	0x76, 0xa0, 0x2e, 0x45, 0xde, 0xe7, 0x9e, 0x92, 0xb7, 0x7b, 0x4a, 0xa1, 0x30, 0xfe, 0x82, 0x0a,
	0x97, 0x8e, 0x2d, 0xdc, 0x92, 0x32, 0xc7, 0x58, 0xf8, 0xc2, 0x81, 0xc5, 0xbb, 0x9f, 0x7a, 0x75,
	0x80, 0xde, 0xb6, 0x60, 0xb6, 0x0e, 0x0b, 0xa7, 0x70, 0x7b, 0x59, 0x7f, 0xbd, 0xac, 0xeb, 0x27,
	0x91, 0x38, 0x91, 0x6f, 0xa0, 0x9e, 0xa2, 0xe1, 0x0a, 0x77, 0x91, 0x42, 0x6d, 0xe2, 0xa2, 0xce,
	0x72, 0x13, 0xb7, 0x72, 0xe7, 0xce, 0xf4, 0x64, 0x77, 0xae, 0x12, 0x3e, 0x1a, 0xc2, 0xd7, 0xa4,
	0x7a, 0x65, 0xc1, 0x64, 0xd5, 0x1a, 0x2c, 0x9c, 0x42, 0xed, 0x35, 0xfe, 0xf3, 0x1a, 0xeb, 0x27,
	0x70, 0x58, 0x89, 0x8b, 0x1f, 0x2a, 0x30, 0x13, 0xa1, 0xee, 0x4a, 0xa1, 0x91, 0xdd, 0x83, 0xb2,
	0x1b, 0xdc, 0x62, 0x8a, 0x6a, 0xe1, 0xe8, 0x5a, 0xf8, 0xa1, 0x7e, 0x61, 0x3f, 0x23, 0x9f, 0x48,
	0x4f, 0x38, 0x67, 0x47, 0x96, 0x1f, 0x9a, 0x59, 0x6a, 0x7a, 0x89, 0xc0, 0xe1, 0x18, 0x78, 0x7c,
	0xb2, 0xb7, 0xe8, 0xbc, 0x71, 0x70, 0x8e, 0x2e, 0x77, 0x46, 0x2f, 0xd8, 0x13, 0xa8, 0x16, 0xab,
	0x42, 0xdd, 0xb4, 0x8c, 0x37, 0x8f, 0x30, 0xfa, 0x45, 0xda, 0xf2, 0xdf, 0xd1, 0x20, 0x9d, 0x6d,
	0x42, 0x49, 0xc9, 0x7d, 0xb2, 0xd7, 0xa2, 0x9e, 0x86, 0xe7, 0xd8, 0xed, 0x70, 0x60, 0x45, 0x18,
	0xc9, 0xfd, 0xc8, 0xb2, 0xd4, 0x7e, 0x95, 0xa0, 0x44, 0x07, 0x36, 0x0f, 0x15, 0x3a, 0xda, 0xb6,
	0x7d, 0x6c, 0x90, 0x3b, 0xe5, 0xa8, 0x4c, 0x47, 0x6a, 0xc4, 0x63, 0x98, 0x19, 0x36, 0xf4, 0x53,
	0x63, 0x72, 0x47, 0xab, 0xa6, 0xe8, 0x20, 0x4d, 0xb0, 0x8b, 0x53, 0x0f, 0xb5, 0x51, 0x19, 0x6d,
	0x52, 0xc2, 0x5b, 0x7d, 0xde, 0x8d, 0x0d, 0xed, 0x94, 0x08, 0x3e, 0x37, 0x46, 0x7b, 0x50, 0x70,
	0xd9, 0x54, 0x91, 0x16, 0x13, 0xec, 0xee, 0xa2, 0x21, 0x7c, 0xa5, 0xdf, 0xf4, 0x60, 0x76, 0x1f,
	0xca, 0x2e, 0x16, 0x7c, 0x99, 0xcc, 0xe2, 0x33, 0xd9, 0x4b, 0x60, 0xd9, 0xd1, 0x3d, 0xfa, 0xda,
	0x98, 0xbc, 0x48, 0x73, 0xd9, 0xf8, 0x02, 0x3d, 0x83, 0x8b, 0xde, 0x93, 0x3d, 0xec, 0x5b, 0x5f,
	0xbe, 0x9d, 0xc1, 0x17, 0x70, 0x88, 0x4d, 0xec, 0x7b, 0x4f, 0xbd, 0x35, 0x84, 0xfd, 0x7e, 0x16,
	0x4f, 0xdd, 0x25, 0x01, 0x1f, 0x41, 0x55, 0x4b, 0x65, 0xb8, 0x90, 0xc1, 0x8f, 0x33, 0xe0, 0x2a,
	0x36, 0xb9, 0x21, 0x57, 0xde, 0x42, 0x3d, 0x93, 0xe3, 0x73, 0x32, 0x7c, 0x83, 0xec, 0x3c, 0x4f,
	0xa5, 0x4e, 0xf6, 0x06, 0xf1, 0xe4, 0xdc, 0x2f, 0x99, 0x56, 0xc5, 0xfd, 0x8d, 0x3f, 0xfc, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0xdb, 0xaf, 0x85, 0xa5, 0x06, 0x00, 0x00,
}