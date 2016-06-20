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

var fileDescriptor0 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x6f, 0x13, 0x3d,
	0x10, 0x55, 0xbf, 0x34, 0x49, 0x35, 0xfd, 0x04, 0x95, 0x91, 0xca, 0x92, 0x00, 0xaa, 0xca, 0x0b,
	0x08, 0x69, 0xcb, 0x45, 0xdc, 0xa4, 0x8a, 0x87, 0x56, 0xb4, 0x2a, 0x55, 0xa3, 0x68, 0x45, 0x2b,
	0xd1, 0x17, 0x6b, 0x93, 0x9d, 0x86, 0x15, 0x1b, 0x3b, 0xb2, 0x1d, 0xaa, 0xbc, 0xf2, 0x0b, 0xb8,
	0xf3, 0xd7, 0x78, 0xe5, 0x91, 0xcb, 0x8f, 0x60, 0x6c, 0x6f, 0xd2, 0x6e, 0x7a, 0x49, 0xe0, 0x25,
	0x89, 0x3d, 0x73, 0xce, 0x9c, 0x3d, 0x33, 0x93, 0x85, 0xd5, 0x44, 0x1b, 0xa9, 0x70, 0x05, 0x45,
	0x27, 0x15, 0xb8, 0xd2, 0x53, 0xb2, 0x8d, 0x49, 0x5f, 0xa1, 0x5e, 0xe9, 0x75, 0xf9, 0x26, 0x9a,
	0x26, 0x2a, 0x2d, 0x45, 0x53, 0x61, 0x82, 0x07, 0x94, 0x90, 0xec, 0xc5, 0x99, 0xe6, 0xcd, 0x7e,
	0x48, 0x89, 0x46, 0xb2, 0xdb, 0x1e, 0x1d, 0x7a, 0x74, 0x78, 0x2e, 0xa4, 0x76, 0x29, 0x2f, 0xf5,
	0x26, 0xce, 0xfa, 0xa8, 0x3d, 0x43, 0xed, 0x4a, 0xb1, 0x3e, 0x2a, 0x25, 0x55, 0x1e, 0xaa, 0x17,
	0x43, 0x5d, 0xd4, 0x3a, 0xee, 0x60, 0x1e, 0xbc, 0x31, 0x1e, 0x34, 0x71, 0x2a, 0x0e, 0xa4, 0xea,
	0xc6, 0x26, 0x95, 0xc2, 0x27, 0x2d, 0x7f, 0x9b, 0x05, 0x68, 0xc6, 0x2a, 0xa6, 0x28, 0x69, 0x62,
	0xbb, 0x10, 0xf4, 0x9c, 0x36, 0xde, 0x7e, 0x45, 0xb7, 0x6d, 0xba, 0x4c, 0xb5, 0x49, 0xdb, 0x3c,
	0x4d, 0x82, 0x99, 0xa5, 0x99, 0x9b, 0xf3, 0xf7, 0xea, 0x61, 0xfe, 0x40, 0xb9, 0xc6, 0x54, 0x18,
	0xec, 0xa0, 0xda, 0xb3, 0xa7, 0x68, 0xd1, 0x83, 0xd7, 0x0b, 0xd8, 0xad, 0x84, 0xad, 0xc1, 0xb5,
	0xb3, 0x68, 0xb9, 0xe8, 0x67, 0x59, 0xf0, 0xa3, 0x4a, 0xe4, 0x73, 0x51, 0xed, 0x74, 0x7c, 0x83,
	0x52, 0xd8, 0x2a, 0xcc, 0x67, 0xb1, 0xe8, 0xf4, 0xe9, 0x01, 0xad, 0x9a, 0xff, 0x26, 0xab, 0x81,
	0x61, 0x3e, 0x29, 0xb8, 0x05, 0x0b, 0xc7, 0xd0, 0xbe, 0xe8, 0x4f, 0x5f, 0xf4, 0xc2, 0x51, 0x9a,
	0x2b, 0xb4, 0x0f, 0x75, 0x29, 0xb2, 0x01, 0xf7, 0x94, 0xbc, 0xdd, 0x57, 0x0a, 0x85, 0xf1, 0x17,
	0x54, 0xb8, 0x74, 0x6a, 0xe1, 0x96, 0x94, 0x19, 0xc6, 0xc2, 0x17, 0x0e, 0x2c, 0xde, 0xfd, 0xd4,
	0xeb, 0x43, 0xf4, 0x9e, 0x05, 0xb3, 0x4d, 0x58, 0x3a, 0x87, 0xdb, 0xcb, 0xfa, 0xe5, 0x65, 0x5d,
	0x3d, 0x8b, 0xc4, 0x89, 0x7c, 0x09, 0xf5, 0x0e, 0x1a, 0xae, 0xf0, 0x00, 0x29, 0xd4, 0x26, 0x2e,
	0xea, 0x2c, 0x37, 0x71, 0x2b, 0x73, 0xee, 0xcc, 0x4e, 0x76, 0xe7, 0x32, 0xe1, 0xa3, 0x11, 0x7c,
	0x43, 0xaa, 0x17, 0x16, 0x4c, 0x56, 0x6d, 0xc0, 0xd2, 0x39, 0xd4, 0x5e, 0xe3, 0x6f, 0xaf, 0xb1,
	0x7e, 0x06, 0x87, 0x95, 0xb8, 0xfc, 0xb6, 0x02, 0x73, 0x11, 0xea, 0x9e, 0x14, 0x1a, 0xd9, 0x1d,
	0x28, 0xbb, 0xc1, 0xcd, 0xa7, 0xa8, 0x16, 0x16, 0xd7, 0xc2, 0x0f, 0xf5, 0x33, 0xfb, 0x19, 0xf9,
	0x44, 0x7a, 0xc2, 0x05, 0x3b, 0xb2, 0xfc, 0xd8, 0xcc, 0x52, 0xd3, 0x4b, 0x04, 0x0e, 0xc7, 0xc0,
	0xe3, 0x93, 0xbd, 0x43, 0xe7, 0xad, 0xa3, 0x73, 0x74, 0xb1, 0x5b, 0xbc, 0x60, 0x8f, 0xa1, 0x9a,
	0xaf, 0x0a, 0x75, 0xd3, 0x32, 0x5e, 0x3f, 0xc1, 0xe8, 0x17, 0x69, 0xc7, 0x7f, 0x47, 0xc3, 0x74,
	0xb6, 0x0d, 0x25, 0x25, 0x0f, 0xc9, 0x5e, 0x8b, 0x7a, 0x12, 0xfe, 0xc5, 0x6e, 0x87, 0x43, 0x2b,
	0xc2, 0x48, 0x1e, 0x46, 0x96, 0xa5, 0xf6, 0xbd, 0x04, 0x25, 0x3a, 0xb0, 0x45, 0xa8, 0xd0, 0xd1,
	0xb6, 0xed, 0x5d, 0x83, 0xdc, 0x29, 0x47, 0x65, 0x3a, 0x52, 0x23, 0x1e, 0xc1, 0xdc, 0xa8, 0xa1,
	0xef, 0x1b, 0x93, 0x3b, 0x5a, 0x35, 0x79, 0x07, 0x69, 0x82, 0x5d, 0x9c, 0x7a, 0xa8, 0x8d, 0x4a,
	0x69, 0x93, 0x12, 0xde, 0x1a, 0xf0, 0x5e, 0x6c, 0x68, 0xa7, 0x44, 0xf0, 0xa1, 0x51, 0xec, 0x41,
	0xce, 0x65, 0x53, 0x45, 0x27, 0x9f, 0x60, 0x77, 0x17, 0x8d, 0xe0, 0x6b, 0x83, 0xa6, 0x07, 0xb3,
	0xbb, 0x50, 0x76, 0xb1, 0xe0, 0xe3, 0x64, 0x16, 0x9f, 0xc9, 0x9e, 0x03, 0x4b, 0x4f, 0xee, 0xd1,
	0xa7, 0xc6, 0xe4, 0x45, 0x5a, 0x48, 0xc7, 0x17, 0xe8, 0x29, 0xfc, 0xef, 0x3d, 0x79, 0x8d, 0x03,
	0xeb, 0xcb, 0xe7, 0x29, 0x7c, 0x01, 0x87, 0xd8, 0xc6, 0x81, 0xf7, 0xd4, 0x5b, 0x43, 0xd8, 0x2f,
	0xd3, 0x78, 0xea, 0x2e, 0x09, 0xf8, 0x00, 0xaa, 0x5a, 0x2a, 0xc3, 0x85, 0x0c, 0xbe, 0x4e, 0x81,
	0xab, 0xd8, 0xe4, 0x86, 0x5c, 0xdb, 0x85, 0x7a, 0x2a, 0xc7, 0xe7, 0x64, 0xf4, 0x06, 0xd9, 0x7f,
	0xf8, 0x6f, 0xef, 0x96, 0x56, 0xc5, 0xfd, 0x7b, 0xdf, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x85, 0xc8, 0x13, 0x9c, 0x06, 0x00, 0x00,
}
