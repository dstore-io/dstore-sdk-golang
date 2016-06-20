// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonDetails_Pu.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonDetails_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonDetails_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonDetails_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	CaseSensitive                  *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull              bool                        `protobuf:"varint,1003,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	PersonCharacteristicId         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull     bool                        `protobuf:"varint,1004,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
	PersonId                       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                   bool                        `protobuf:"varint,1005,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
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

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonDetails_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonDetails_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonDetails_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonDetails_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x4f,
	0x14, 0xa7, 0xcd, 0x3f, 0x6d, 0xfe, 0x63, 0xad, 0x75, 0x02, 0x65, 0xdd, 0xd8, 0xa2, 0x15, 0xd1,
	0xab, 0x8d, 0x28, 0x42, 0xc1, 0x2b, 0xa3, 0x22, 0x11, 0x52, 0xc2, 0xf8, 0x01, 0x16, 0x61, 0xd9,
	0xee, 0x9e, 0xc4, 0xc1, 0x64, 0x26, 0xcc, 0x4c, 0x5a, 0xfa, 0x16, 0x5e, 0xfb, 0x26, 0x3e, 0x92,
	0x5f, 0xef, 0xe0, 0x99, 0x3d, 0xbb, 0xb1, 0xbb, 0x36, 0x56, 0x6f, 0x92, 0x4c, 0xce, 0xef, 0x6b,
	0xe6, 0x9c, 0x19, 0xb6, 0x9f, 0x59, 0xa7, 0x0d, 0x74, 0x41, 0x8d, 0xa5, 0x82, 0xee, 0xcc, 0xe8,
	0x14, 0xb2, 0xb9, 0x01, 0xdb, 0x9d, 0x4d, 0xe3, 0x81, 0xce, 0xe4, 0xe8, 0x74, 0x08, 0xc6, 0x6a,
	0xf5, 0x14, 0x5c, 0x22, 0x27, 0x36, 0x1e, 0xce, 0x23, 0x04, 0x39, 0xcd, 0xef, 0x10, 0x33, 0x22,
	0x66, 0xb4, 0x14, 0x1e, 0xb6, 0x0b, 0x8b, 0xe3, 0x64, 0x32, 0x07, 0x4b, 0xec, 0xf0, 0x5a, 0xd5,
	0x17, 0x8c, 0xd1, 0xa6, 0x28, 0x75, 0xaa, 0xa5, 0x29, 0x58, 0x9b, 0x8c, 0xa1, 0x28, 0xde, 0xaa,
	0x17, 0xd1, 0x46, 0x8d, 0xb4, 0x99, 0x26, 0x4e, 0x6a, 0x45, 0xa0, 0xbd, 0x4f, 0x4d, 0xc6, 0x86,
	0x89, 0x49, 0xb0, 0x8a, 0x61, 0xf8, 0x3b, 0x76, 0x7d, 0x96, 0x87, 0x8a, 0x65, 0x06, 0xca, 0xc9,
	0x91, 0x4c, 0x73, 0x74, 0x4c, 0x89, 0x82, 0x95, 0x1b, 0x2b, 0x77, 0x2f, 0xdd, 0x0f, 0xa3, 0x62,
	0x43, 0x45, 0x4e, 0xeb, 0x8c, 0x54, 0xe3, 0x37, 0x7e, 0x21, 0x42, 0xe2, 0xf7, 0x2b, 0xf4, 0xbc,
	0x64, 0xf9, 0x0b, 0x76, 0xf3, 0x4f, 0xea, 0xb1, 0x9a, 0x4f, 0x26, 0xc1, 0x97, 0x75, 0xf4, 0x68,
	0x89, 0xdd, 0xe5, 0x3a, 0x07, 0x08, 0xe3, 0x8f, 0xd9, 0x66, 0xa1, 0xe5, 0x4e, 0x67, 0x80, 0x82,
	0xc1, 0x6a, 0x9e, 0xad, 0x53, 0xcb, 0x26, 0x95, 0x83, 0x31, 0x18, 0x0a, 0xb7, 0x41, 0x94, 0x57,
	0xc8, 0xe8, 0x67, 0x3c, 0x62, 0xed, 0xaa, 0x04, 0x05, 0xf8, 0x4a, 0x01, 0xb6, 0xce, 0x62, 0x73,
	0xcb, 0x1e, 0xdb, 0x4c, 0x13, 0x0b, 0xb1, 0x05, 0x65, 0xa5, 0x93, 0xc7, 0x10, 0x34, 0xce, 0xb5,
	0x3c, 0xd2, 0x7a, 0x02, 0x09, 0x85, 0x15, 0x97, 0x3d, 0xe5, 0x65, 0xc9, 0xe0, 0x5d, 0xd6, 0xae,
	0x6a, 0x90, 0xe7, 0x37, 0xf2, 0xbc, 0x5a, 0x01, 0xe7, 0xa6, 0xaf, 0x59, 0x50, 0x84, 0x4c, 0xdf,
	0x63, 0x9f, 0x52, 0x6c, 0x93, 0xb4, 0x4e, 0xa6, 0x7e, 0xc7, 0xff, 0x5d, 0xbc, 0xe3, 0x6d, 0x22,
	0x3f, 0xa9, 0x70, 0x71, 0xef, 0x3d, 0xb6, 0xb3, 0x4c, 0x96, 0x12, 0x7d, 0xa7, 0x44, 0xe1, 0xf9,
	0xfc, 0x3c, 0xda, 0x3e, 0xfb, 0x7f, 0xd1, 0xce, 0xa0, 0x79, 0x71, 0x96, 0x56, 0xd9, 0x52, 0x7e,
	0x7b, 0xd1, 0xbc, 0xd2, 0xee, 0x07, 0xd9, 0x6d, 0x94, 0x10, 0x6f, 0xb0, 0xf7, 0x79, 0x95, 0xb5,
	0x04, 0xd8, 0x99, 0x56, 0x16, 0xf8, 0x3d, 0xd6, 0xcc, 0x47, 0xbf, 0x3e, 0x83, 0xc5, 0xa5, 0xa2,
	0x6b, 0xf1, 0xcc, 0x7f, 0x0a, 0x02, 0xf2, 0xb7, 0x6c, 0xcb, 0x0f, 0x7d, 0x7c, 0x66, 0xea, 0x71,
	0x48, 0x1a, 0x48, 0x8e, 0x6a, 0xe4, 0xfa, 0xdd, 0x18, 0xe0, 0xba, 0xff, 0x6b, 0x2d, 0xae, 0x4c,
	0xab, 0x7f, 0xe0, 0xd6, 0xd7, 0x8b, 0xcb, 0x86, 0x33, 0xe0, 0x15, 0x77, 0x7f, 0x53, 0xa4, 0xab,
	0x38, 0xa0, 0x6f, 0x51, 0xc2, 0xf9, 0x73, 0xd6, 0x30, 0xfa, 0x04, 0x5b, 0xe7, 0x59, 0x0f, 0xa3,
	0xbf, 0x7c, 0x19, 0xa2, 0xf2, 0x18, 0x22, 0xa1, 0x4f, 0x84, 0x57, 0x08, 0x77, 0x58, 0x03, 0x7f,
	0xf3, 0x6d, 0xb6, 0x86, 0x2b, 0xdf, 0x81, 0x8f, 0x07, 0x78, 0x30, 0x4d, 0xd1, 0xc4, 0x65, 0x3f,
	0xeb, 0x1d, 0xb2, 0x8e, 0xd4, 0x75, 0xf9, 0xc5, 0x93, 0x75, 0xf8, 0x68, 0xac, 0x6d, 0xf6, 0xa1,
	0xac, 0x67, 0xff, 0xf4, 0xaa, 0x1d, 0xad, 0xe5, 0x6f, 0xc7, 0x83, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x4c, 0x4b, 0x3f, 0x12, 0x05, 0x00, 0x00,
}