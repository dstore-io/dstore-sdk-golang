// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetRelationshipSettings_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetRelationshipSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetRelationshipSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetRelationshipSettings_Ad

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
	RelationshipId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull   bool                        `protobuf:"varint,1001,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	FromPersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=from_person_type_id,json=fromPersonTypeId" json:"from_person_type_id,omitempty"`
	FromPersonTypeIdNull bool                        `protobuf:"varint,1002,opt,name=from_person_type_id_null,json=fromPersonTypeIdNull" json:"from_person_type_id_null,omitempty"`
	ToPersonTypeId       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=to_person_type_id,json=toPersonTypeId" json:"to_person_type_id,omitempty"`
	ToPersonTypeIdNull   bool                        `protobuf:"varint,1003,opt,name=to_person_type_id_null,json=toPersonTypeIdNull" json:"to_person_type_id_null,omitempty"`
	KeyVariable          *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull      bool                        `protobuf:"varint,1004,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetRelationshipId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelationshipId
	}
	return nil
}

func (m *Parameters) GetFromPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromPersonTypeId
	}
	return nil
}

func (m *Parameters) GetToPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToPersonTypeId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ToPersonTypeDescription   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=to_person_type_description,json=toPersonTypeDescription" json:"to_person_type_description,omitempty"`
	FromPersonTypeDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=from_person_type_description,json=fromPersonTypeDescription" json:"from_person_type_description,omitempty"`
	RelationshipId            *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	FromPersonTypeId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=from_person_type_id,json=fromPersonTypeId" json:"from_person_type_id,omitempty"`
	Relationship              *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=relationship" json:"relationship,omitempty"`
	Value                     *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=value" json:"value,omitempty"`
	KeyVariable               *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	ToPersonTypeId            *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=to_person_type_id,json=toPersonTypeId" json:"to_person_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetToPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ToPersonTypeDescription
	}
	return nil
}

func (m *Response_Row) GetFromPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.FromPersonTypeDescription
	}
	return nil
}

func (m *Response_Row) GetRelationshipId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelationshipId
	}
	return nil
}

func (m *Response_Row) GetFromPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromPersonTypeId
	}
	return nil
}

func (m *Response_Row) GetRelationship() *dstore_values.StringValue {
	if m != nil {
		return m.Relationship
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Response_Row) GetToPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToPersonTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetRelationshipSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetRelationshipSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetRelationshipSettings_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x4f, 0x14, 0x4d,
	0x10, 0x0d, 0xdf, 0xb2, 0x2c, 0x29, 0xc8, 0xb7, 0xd0, 0x10, 0x1c, 0x76, 0x8d, 0x31, 0xf8, 0x62,
	0x42, 0x32, 0xab, 0x92, 0x78, 0x49, 0xbc, 0x44, 0x03, 0x12, 0x24, 0x4c, 0xc8, 0xa8, 0x44, 0x8d,
	0xc9, 0x64, 0x60, 0x8a, 0x75, 0xe2, 0xee, 0xf4, 0xa4, 0xbb, 0x17, 0xc2, 0x2f, 0xf0, 0xd5, 0xbb,
	0xfe, 0x42, 0x13, 0x2f, 0x6f, 0xfe, 0x01, 0x6b, 0xa6, 0x1b, 0x98, 0xe9, 0xdd, 0xb8, 0x2e, 0x2f,
	0x90, 0xde, 0xaa, 0x73, 0x4e, 0x75, 0xd5, 0xa9, 0x1e, 0xb8, 0x1d, 0x49, 0xc5, 0x05, 0xb6, 0x30,
	0x69, 0xc7, 0x09, 0xb6, 0x52, 0xc1, 0xf7, 0x30, 0xea, 0x09, 0x94, 0xad, 0xb4, 0x1b, 0xac, 0xa3,
	0xf2, 0xb1, 0x13, 0xaa, 0x98, 0x27, 0xf2, 0x55, 0x9c, 0x3e, 0x46, 0xa5, 0xe2, 0xa4, 0x2d, 0x83,
	0xfb, 0x91, 0x4b, 0x89, 0x8a, 0xb3, 0x65, 0x8d, 0x76, 0x35, 0xda, 0xfd, 0x2b, 0xa4, 0x31, 0x67,
	0xa4, 0x0e, 0xc2, 0x4e, 0x0f, 0xa5, 0x66, 0x68, 0x2c, 0x96, 0xf5, 0x51, 0x08, 0x2e, 0x4c, 0xa8,
	0x59, 0x0e, 0x75, 0x51, 0xca, 0xb0, 0x8d, 0x26, 0x78, 0xc9, 0x0e, 0xaa, 0x30, 0x4e, 0xf6, 0xb9,
	0xe8, 0xe6, 0xd2, 0x3a, 0x69, 0xe9, 0x77, 0x05, 0x60, 0x3b, 0x14, 0x21, 0x45, 0x51, 0x48, 0xb6,
	0x0a, 0x75, 0x51, 0xa8, 0x2d, 0x88, 0x23, 0x67, 0xec, 0xe2, 0xd8, 0xe5, 0xa9, 0x6b, 0x4d, 0xd7,
	0xdc, 0xc3, 0x94, 0x16, 0x27, 0x0a, 0xdb, 0x28, 0x76, 0xb2, 0x93, 0xff, 0x7f, 0x11, 0xb3, 0x11,
	0xb1, 0xab, 0x30, 0x6f, 0xb1, 0x04, 0x49, 0xaf, 0xd3, 0x71, 0xbe, 0xd7, 0x88, 0x6b, 0xd2, 0x67,
	0xe5, 0x74, 0x8f, 0x42, 0xec, 0x11, 0xcc, 0xed, 0x0b, 0xde, 0x0d, 0x52, 0xaa, 0x82, 0x27, 0x81,
	0x3a, 0x4a, 0x31, 0x13, 0xff, 0x6f, 0xb8, 0xf8, 0x4c, 0x86, 0xdb, 0xce, 0x61, 0x4f, 0x08, 0x45,
	0xf2, 0x37, 0xc0, 0x19, 0xc0, 0xa5, 0x4b, 0xf8, 0xa1, 0x4b, 0x98, 0xb7, 0x41, 0x79, 0x11, 0x0f,
	0x61, 0x56, 0x71, 0xbb, 0x84, 0xca, 0x3f, 0xdc, 0x5f, 0xf1, 0x52, 0x01, 0x2b, 0xb0, 0xd0, 0xc7,
	0xa3, 0xe5, 0x7f, 0x9a, 0x0e, 0x94, 0x01, 0xb9, 0xf8, 0x1d, 0x98, 0x7e, 0x8d, 0x47, 0xc1, 0x41,
	0x28, 0xe2, 0x70, 0xb7, 0x83, 0xce, 0x78, 0xae, 0xdb, 0xb0, 0x74, 0xa5, 0x12, 0x64, 0x16, 0x2d,
	0x3b, 0x45, 0xf9, 0x3b, 0x26, 0x9d, 0x2d, 0xc3, 0x6c, 0x11, 0xae, 0xe5, 0x7e, 0x69, 0xb9, 0x7a,
	0x21, 0x31, 0xd3, 0x5a, 0x7a, 0x53, 0x83, 0x49, 0x1f, 0x65, 0x4a, 0x23, 0x40, 0x76, 0x05, 0xaa,
	0xb9, 0xa7, 0xcc, 0xa4, 0x4f, 0x14, 0x8d, 0x63, 0xb5, 0xdf, 0xd6, 0xb2, 0xbf, 0xbe, 0x4e, 0x64,
	0xcf, 0x61, 0x26, 0x73, 0x53, 0x50, 0xb0, 0x13, 0x4d, 0xaa, 0x42, 0x60, 0xd7, 0x02, 0xdb, 0xa6,
	0xdb, 0xa2, 0xf3, 0xc6, 0xe9, 0xd9, 0xaf, 0x77, 0xcb, 0x3f, 0xb0, 0x9b, 0x50, 0x33, 0x2e, 0xa6,
	0xc6, 0x67, 0x8c, 0x17, 0xfa, 0x18, 0xb5, 0xc7, 0xb7, 0xf4, 0x7f, 0xff, 0x38, 0x9d, 0x6d, 0x42,
	0x45, 0xf0, 0x43, 0x6a, 0x5b, 0x86, 0xba, 0xe5, 0x8e, 0xb0, 0x76, 0xee, 0x71, 0x2b, 0x5c, 0x9f,
	0x1f, 0xfa, 0x19, 0x4b, 0xe3, 0xdb, 0x38, 0x54, 0xe8, 0xc0, 0x16, 0x60, 0x82, 0x8e, 0x99, 0x0d,
	0xde, 0x7a, 0xd4, 0x9d, 0xaa, 0x5f, 0xa5, 0x23, 0x4d, 0xf8, 0x19, 0x34, 0xac, 0x09, 0x47, 0x28,
	0xf7, 0x44, 0x9c, 0xe6, 0xbd, 0x78, 0xe7, 0x0d, 0x9d, 0xdd, 0xb9, 0xa2, 0x03, 0x56, 0x4f, 0xb1,
	0xec, 0x25, 0x9c, 0xef, 0x33, 0x6f, 0x91, 0xfb, 0xfd, 0x70, 0xee, 0xc5, 0xb2, 0xb9, 0x8b, 0xec,
	0x6b, 0xfd, 0xfb, 0xfd, 0xc1, 0x1b, 0x7d, 0xc1, 0x37, 0x07, 0x6f, 0xeb, 0x47, 0xef, 0x2c, 0xeb,
	0x7a, 0x0f, 0xa6, 0x8b, 0xf4, 0xce, 0xa7, 0xe1, 0x37, 0x2c, 0x01, 0xe8, 0xb9, 0xa9, 0xe6, 0x39,
	0xce, 0xe7, 0xe1, 0x48, 0x9d, 0xc9, 0xee, 0x5a, 0xcb, 0xf6, 0xc5, 0x1b, 0x6d, 0xdb, 0xd6, 0x07,
	0xbd, 0x14, 0x5f, 0xbd, 0x91, 0x9f, 0x8a, 0x07, 0x4f, 0xa1, 0x19, 0x73, 0xdb, 0xac, 0x27, 0x5f,
	0x98, 0x17, 0xd7, 0xcf, 0xf6, 0xed, 0xd9, 0x9d, 0xc8, 0x5f, 0xf7, 0x95, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0xcd, 0x29, 0x15, 0xbc, 0x06, 0x00, 0x00,
}