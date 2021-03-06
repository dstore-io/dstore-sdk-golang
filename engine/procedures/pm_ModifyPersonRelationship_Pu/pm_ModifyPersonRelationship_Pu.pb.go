// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonRelationship_Pu.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonRelationship_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonRelationship_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonRelationship_Pu

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
	PersonIdentificationValues      *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull  bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull                bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                        *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                    bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	RelatedPersonId                 *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=related_person_id,json=relatedPersonId" json:"related_person_id,omitempty"`
	RelatedPersonIdNull             bool                        `protobuf:"varint,1004,opt,name=related_person_id_null,json=relatedPersonIdNull" json:"related_person_id_null,omitempty"`
	RelationshipId                  *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull              bool                        `protobuf:"varint,1005,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	AccessLevel                     *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=access_level,json=accessLevel" json:"access_level,omitempty"`
	AccessLevelNull                 bool                        `protobuf:"varint,1006,opt,name=access_level_null,json=accessLevelNull" json:"access_level_null,omitempty"`
	RelatedIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=related_identification_values,json=relatedIdentificationValues" json:"related_identification_values,omitempty"`
	RelatedIdentificationValuesNull bool                        `protobuf:"varint,1007,opt,name=related_identification_values_null,json=relatedIdentificationValuesNull" json:"related_identification_values_null,omitempty"`
	SeparatorInIdentVals            *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull        bool                        `protobuf:"varint,1008,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetRelatedPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelatedPersonId
	}
	return nil
}

func (m *Parameters) GetRelationshipId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelationshipId
	}
	return nil
}

func (m *Parameters) GetAccessLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevel
	}
	return nil
}

func (m *Parameters) GetRelatedIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.RelatedIdentificationValues
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonRelationship_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonRelationship_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonRelationship_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_ModifyPersonRelationship_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x6a, 0xd4, 0x40,
	0x14, 0xa6, 0x5d, 0xb7, 0x5d, 0x4f, 0x4b, 0x2f, 0xd3, 0x52, 0xe3, 0x6e, 0x5b, 0xb5, 0x22, 0x08,
	0x4a, 0xea, 0x0d, 0x14, 0xc1, 0x2b, 0x8a, 0xac, 0xb4, 0x65, 0x0d, 0x22, 0x28, 0x62, 0x88, 0x9b,
	0xe9, 0x3a, 0xb8, 0x9b, 0x89, 0x33, 0x49, 0x4b, 0xdf, 0xc2, 0x47, 0xf1, 0x09, 0x7c, 0x1f, 0xef,
	0x3e, 0x82, 0x27, 0x73, 0x26, 0xeb, 0x26, 0x6d, 0x37, 0xf4, 0x4f, 0xbb, 0x93, 0xf3, 0xdd, 0x32,
	0x33, 0xe7, 0x04, 0xee, 0x85, 0x3a, 0x91, 0x8a, 0x6f, 0xf2, 0xa8, 0x27, 0x22, 0xbe, 0x19, 0x2b,
	0xd9, 0xe5, 0x61, 0xaa, 0xb8, 0xde, 0x8c, 0x07, 0xfe, 0xb6, 0x0c, 0xc5, 0xee, 0x41, 0x87, 0x2b,
	0x2d, 0x23, 0x8f, 0xf7, 0x83, 0x44, 0xc8, 0x48, 0x7f, 0x10, 0xb1, 0xdf, 0x49, 0x5d, 0x44, 0x26,
	0x92, 0x5d, 0x25, 0xba, 0x4b, 0x74, 0x77, 0x3c, 0xa7, 0xb9, 0x64, 0xcd, 0xf6, 0x82, 0x7e, 0xca,
	0x35, 0x49, 0x34, 0xcf, 0x16, 0x13, 0x70, 0xa5, 0xa4, 0xb2, 0xa5, 0x56, 0xb1, 0x34, 0xe0, 0x5a,
	0x07, 0x3d, 0x6e, 0x8b, 0x17, 0xcb, 0xc5, 0x24, 0x10, 0xd1, 0xae, 0x54, 0x03, 0x63, 0x49, 0xa0,
	0x8d, 0x2f, 0x0d, 0x80, 0x4e, 0xa0, 0x02, 0xac, 0x62, 0x22, 0xf6, 0x16, 0x56, 0x63, 0x93, 0xcc,
	0x17, 0x21, 0x8f, 0x12, 0xb1, 0x2b, 0xba, 0x06, 0xed, 0x53, 0x22, 0x67, 0xe2, 0xfc, 0xc4, 0xe5,
	0x99, 0x1b, 0x4d, 0xd7, 0xbe, 0x95, 0xcd, 0xa9, 0x13, 0x25, 0xa2, 0xde, 0xab, 0x6c, 0xe1, 0x35,
	0x89, 0xdf, 0x2e, 0xd0, 0x4d, 0x49, 0xb3, 0xe7, 0x70, 0x61, 0x9c, 0xba, 0x1f, 0xa5, 0xfd, 0xbe,
	0xf3, 0x6d, 0x1a, 0x3d, 0x1a, 0xde, 0xfa, 0xf1, 0x3a, 0x3b, 0x08, 0x63, 0x8f, 0x60, 0xce, 0x6a,
	0x25, 0x07, 0x31, 0x47, 0x41, 0x67, 0xd2, 0x64, 0x6b, 0x95, 0xb2, 0x89, 0x28, 0xe1, 0x3d, 0xae,
	0x28, 0xdc, 0x2c, 0x51, 0x5e, 0x22, 0xa3, 0x1d, 0x32, 0x17, 0x96, 0x8a, 0x12, 0x14, 0xe0, 0x3b,
	0x05, 0x58, 0x18, 0xc5, 0x1a, 0xcb, 0xdb, 0x70, 0x3a, 0x8d, 0xc4, 0xa7, 0xd4, 0xb8, 0xd5, 0x2a,
	0x77, 0xa2, 0x41, 0x60, 0x34, 0xba, 0x04, 0x73, 0x43, 0x22, 0x79, 0xfc, 0x20, 0x8f, 0xd9, 0x1c,
	0x62, 0xf4, 0x9f, 0xc1, 0xa2, 0xca, 0x2e, 0x04, 0x0f, 0xfd, 0xe1, 0x36, 0x39, 0xa7, 0xaa, 0xdf,
	0x6a, 0xde, 0xb2, 0x3a, 0x76, 0xc7, 0xd8, 0x2d, 0x58, 0x39, 0x24, 0x44, 0xbe, 0x3f, 0xc9, 0x77,
	0xa9, 0xc4, 0x30, 0xf6, 0x4f, 0x80, 0x84, 0xf2, 0xfb, 0x88, 0xe6, 0xf5, 0x6a, 0xf3, 0xb9, 0x51,
	0x0e, 0x7a, 0x5f, 0x87, 0xe5, 0x92, 0x0a, 0x39, 0xff, 0x22, 0x67, 0x56, 0x84, 0x1b, 0xe3, 0xfb,
	0x30, 0x1b, 0x74, 0xbb, 0x78, 0x77, 0xfd, 0x3e, 0xdf, 0xe3, 0x7d, 0x67, 0xaa, 0xda, 0x75, 0x86,
	0x08, 0x5b, 0x19, 0x9e, 0x5d, 0x81, 0xc5, 0x51, 0x3e, 0xf9, 0xfd, 0x26, 0xbf, 0xf9, 0x11, 0xa0,
	0x31, 0x7b, 0x07, 0x6b, 0xf9, 0xde, 0x1c, 0x7d, 0xc5, 0xa7, 0x2b, 0x0f, 0xb6, 0x65, 0x05, 0x8e,
	0xbc, 0xe3, 0x5b, 0xb0, 0x31, 0x56, 0x9f, 0xd2, 0xfd, 0xa1, 0x74, 0xe7, 0xc6, 0x28, 0x99, 0xb4,
	0x2f, 0xe0, 0x8c, 0xe6, 0x31, 0xf6, 0x27, 0x66, 0xf1, 0x85, 0xed, 0x9b, 0x4c, 0x49, 0x3b, 0x8d,
	0xca, 0x9c, 0xcb, 0x43, 0x6a, 0x9b, 0xfa, 0x08, 0x1f, 0x6b, 0xf6, 0x00, 0x56, 0x8f, 0x91, 0xa4,
	0x68, 0x7f, 0x29, 0x9a, 0x73, 0x14, 0x39, 0xcb, 0xb4, 0xf1, 0x75, 0x12, 0x1a, 0x1e, 0xd7, 0x31,
	0x9e, 0x21, 0x67, 0xd7, 0xa0, 0x6e, 0x06, 0x52, 0x79, 0x32, 0xd8, 0x79, 0x47, 0xc3, 0xea, 0x69,
	0xf6, 0xd7, 0x23, 0x20, 0x7b, 0x0d, 0x0b, 0xd9, 0x28, 0xf2, 0x47, 0x66, 0x11, 0xb6, 0x6e, 0x0d,
	0xc9, 0x6e, 0x89, 0x5c, 0x9e, 0x58, 0xdb, 0xb8, 0x6e, 0xff, 0x5f, 0x7b, 0xf3, 0x83, 0xe2, 0x03,
	0x76, 0x07, 0xa6, 0xed, 0x08, 0xc4, 0xf6, 0xcc, 0x14, 0xd7, 0x0f, 0x29, 0xd2, 0x80, 0xdc, 0xa6,
	0xff, 0x5e, 0x0e, 0xc7, 0x53, 0xab, 0x29, 0xb9, 0x8f, 0xcd, 0x96, 0xb1, 0xee, 0xba, 0x27, 0x19,
	0xda, 0x6e, 0xbe, 0x17, 0xae, 0x27, 0xf7, 0xbd, 0x4c, 0xa6, 0xb9, 0x06, 0x35, 0xfc, 0xcd, 0x56,
	0x60, 0x0a, 0x57, 0x59, 0x1f, 0x7d, 0xde, 0xc1, 0xdd, 0xa9, 0x7b, 0x75, 0x5c, 0xb6, 0xc3, 0xc7,
	0x3e, 0xb4, 0x84, 0x2c, 0x7b, 0x0c, 0xbf, 0x2b, 0x6f, 0x1e, 0xf6, 0xa4, 0x0e, 0x3f, 0xe6, 0xf5,
	0xf0, 0xe4, 0x9f, 0x9e, 0xf7, 0x53, 0x66, 0xb6, 0xdf, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x2f,
	0x17, 0x02, 0xba, 0xbc, 0x06, 0x00, 0x00,
}
