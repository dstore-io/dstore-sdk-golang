// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_PrioritizeARelationship_Pu.proto
// DO NOT EDIT!

/*
Package pm_PrioritizeARelationship_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_PrioritizeARelationship_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_PrioritizeARelationship_Pu

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
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	RelatedPersonId                *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=related_person_id,json=relatedPersonId" json:"related_person_id,omitempty"`
	RelatedPersonIdNull            bool                        `protobuf:"varint,1004,opt,name=related_person_id_null,json=relatedPersonIdNull" json:"related_person_id_null,omitempty"`
	RelationshipId                 *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull             bool                        `protobuf:"varint,1005,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	MovePriorityByX                *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=move_priority_by_x,json=movePriorityByX" json:"move_priority_by_x,omitempty"`
	MovePriorityByXNull            bool                        `protobuf:"varint,1006,opt,name=move_priority_by_x_null,json=movePriorityByXNull" json:"move_priority_by_x_null,omitempty"`
	IncreasePriority               *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=increase_priority,json=increasePriority" json:"increase_priority,omitempty"`
	IncreasePriorityNull           bool                        `protobuf:"varint,1007,opt,name=increase_priority_null,json=increasePriorityNull" json:"increase_priority_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1008,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetMovePriorityByX() *dstore_values.IntegerValue {
	if m != nil {
		return m.MovePriorityByX
	}
	return nil
}

func (m *Parameters) GetIncreasePriority() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncreasePriority
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_PrioritizeARelationship_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_PrioritizeARelationship_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_PrioritizeARelationship_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_PrioritizeARelationship_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0xff, 0x4e, 0xd4, 0x40,
	0x10, 0xc7, 0x03, 0xe7, 0xc1, 0x39, 0x12, 0x84, 0x85, 0x40, 0x3d, 0x90, 0x28, 0xc6, 0xc4, 0xc4,
	0xa4, 0xf8, 0x33, 0x6a, 0x62, 0x42, 0x20, 0x1a, 0x3d, 0x0d, 0xe4, 0x6c, 0x8c, 0x51, 0x63, 0x6c,
	0xca, 0x75, 0x39, 0x37, 0xde, 0x75, 0xeb, 0x6e, 0x0b, 0x9e, 0x4f, 0xe1, 0x7f, 0x3e, 0x8d, 0x0f,
	0xe4, 0x6f, 0x1f, 0xc1, 0xe9, 0xce, 0xf6, 0xb8, 0xf6, 0x80, 0x93, 0x7f, 0x38, 0xb6, 0x33, 0xdf,
	0xef, 0x67, 0x76, 0xbb, 0x33, 0x85, 0xfb, 0xa1, 0x4e, 0xa4, 0xe2, 0x6b, 0x3c, 0x6a, 0x8b, 0x88,
	0xaf, 0xc5, 0x4a, 0xb6, 0x78, 0x98, 0x2a, 0xae, 0xd7, 0xe2, 0xae, 0xdf, 0x54, 0x42, 0x2a, 0x91,
	0x88, 0x4f, 0x7c, 0xc3, 0xe3, 0x9d, 0x20, 0x11, 0x32, 0xd2, 0xef, 0x44, 0xec, 0x37, 0x53, 0x17,
	0x13, 0x13, 0xc9, 0xae, 0x92, 0xda, 0x25, 0xb5, 0x7b, 0xac, 0xa4, 0x3e, 0x67, 0x51, 0x7b, 0x41,
	0x27, 0xe5, 0x9a, 0x1c, 0xea, 0xe7, 0x8a, 0x7c, 0xae, 0x94, 0x54, 0x36, 0xb4, 0x54, 0x0c, 0x75,
	0xb9, 0xd6, 0x41, 0x9b, 0xdb, 0xe0, 0xa5, 0x72, 0x30, 0x09, 0x44, 0xb4, 0x2b, 0x55, 0xd7, 0x20,
	0x29, 0x69, 0xf5, 0x4b, 0x0d, 0xa0, 0x19, 0xa8, 0x00, 0xa3, 0x5c, 0x69, 0xf6, 0x06, 0x96, 0x63,
	0xfc, 0x95, 0x91, 0x2f, 0x42, 0x1e, 0x25, 0x62, 0x57, 0xb4, 0x4c, 0xb6, 0x4f, 0x15, 0x39, 0x63,
	0x17, 0xc6, 0xae, 0x9c, 0xb9, 0x51, 0x77, 0xed, 0xa6, 0x6c, 0x9d, 0x3a, 0x51, 0x22, 0x6a, 0xbf,
	0xc8, 0x16, 0x5e, 0x9d, 0xf4, 0x8d, 0x82, 0xdc, 0x84, 0x34, 0x7b, 0x02, 0x17, 0x8f, 0x73, 0xf7,
	0xa3, 0xb4, 0xd3, 0x71, 0xbe, 0x4d, 0x22, 0xa3, 0xe6, 0xad, 0x1c, 0xed, 0xb3, 0x8d, 0x69, 0x6c,
	0x03, 0xa6, 0xad, 0x57, 0xd2, 0x8b, 0x39, 0x1a, 0x3a, 0xe3, 0xa6, 0xb6, 0xa5, 0x52, 0x6d, 0x22,
	0x4a, 0x78, 0x9b, 0x2b, 0x2a, 0x6e, 0x8a, 0x24, 0xcf, 0x51, 0xd1, 0x08, 0x99, 0x0b, 0x73, 0x45,
	0x0b, 0x2a, 0xe0, 0x3b, 0x15, 0x30, 0x33, 0x98, 0x6b, 0x90, 0x77, 0xe0, 0x74, 0x1a, 0x89, 0x0f,
	0xa9, 0xa1, 0x55, 0x46, 0x9e, 0x44, 0x8d, 0x92, 0x11, 0x74, 0x19, 0xa6, 0xfb, 0x42, 0x62, 0xfc,
	0x20, 0xc6, 0x54, 0x9e, 0x62, 0xfc, 0x1f, 0xc1, 0xac, 0xca, 0x2e, 0x04, 0x0f, 0xfd, 0xfe, 0x31,
	0x39, 0xa7, 0x46, 0xef, 0xea, 0xac, 0x55, 0x35, 0xed, 0x89, 0xb1, 0x5b, 0xb0, 0x30, 0x64, 0x44,
	0xdc, 0x9f, 0xc4, 0x9d, 0x2b, 0x29, 0x0c, 0xfe, 0x01, 0x90, 0x51, 0x7e, 0x1f, 0x11, 0x5e, 0x1d,
	0x0d, 0x9f, 0x1e, 0xd4, 0x20, 0xfb, 0x3a, 0xcc, 0x97, 0x5c, 0x88, 0xfc, 0x8b, 0xc8, 0xac, 0x98,
	0x6e, 0xc0, 0x8f, 0x81, 0x75, 0xe5, 0x1e, 0xf7, 0x63, 0x6a, 0x8c, 0x9e, 0xbf, 0xd3, 0xf3, 0x3f,
	0x3a, 0x13, 0xff, 0xb1, 0xf1, 0x4c, 0x66, 0xdb, 0xa9, 0xb7, 0xd9, 0x7b, 0xc9, 0x6e, 0xc3, 0xe2,
	0xb0, 0x13, 0xf1, 0x7f, 0xdb, 0x9d, 0x97, 0x24, 0xb6, 0x80, 0x59, 0x11, 0xb5, 0x14, 0x0f, 0xf4,
	0x81, 0xd4, 0x99, 0x3c, 0x94, 0xbf, 0x23, 0x65, 0x87, 0x07, 0x74, 0x11, 0xbd, 0x99, 0x5c, 0x95,
	0x1b, 0x62, 0x01, 0x0b, 0x43, 0x4e, 0xc4, 0xff, 0x43, 0xfc, 0xf9, 0xb2, 0xc4, 0x14, 0xf0, 0x0c,
	0x16, 0x35, 0x8f, 0xb1, 0x0d, 0x11, 0xe5, 0x0b, 0xdb, 0x1e, 0x59, 0x57, 0x68, 0xa7, 0x36, 0xf2,
	0x9e, 0xcd, 0xf7, 0xa5, 0x0d, 0x6a, 0x17, 0x7c, 0xac, 0xd9, 0x3a, 0x2c, 0x1f, 0x61, 0x49, 0xf5,
	0xfc, 0xa5, 0x7a, 0x9c, 0xc3, 0xc4, 0x59, 0x4d, 0xab, 0x5f, 0xc7, 0xa1, 0xe6, 0x71, 0x1d, 0xe3,
	0xab, 0xe2, 0xec, 0x1a, 0x54, 0xcd, 0xdc, 0x29, 0x0f, 0x00, 0x3b, 0xd5, 0x68, 0x26, 0x3d, 0xcc,
	0xfe, 0x7a, 0x94, 0xc8, 0x5e, 0xc1, 0x4c, 0x36, 0x71, 0xfc, 0x81, 0x91, 0x83, 0x1d, 0x5a, 0x41,
	0xb1, 0x5b, 0x12, 0x97, 0x07, 0xd3, 0x16, 0xae, 0x1b, 0x07, 0x6b, 0x7c, 0xcb, 0xc5, 0x07, 0xec,
	0x2e, 0x4c, 0xda, 0x49, 0x87, 0x5d, 0x98, 0x39, 0xae, 0x0c, 0x39, 0xd2, 0x1c, 0xdc, 0xa2, 0x5f,
	0x2f, 0x4f, 0x67, 0x4f, 0xa1, 0xa2, 0xe4, 0x3e, 0xf6, 0x54, 0xa6, 0xba, 0xe7, 0x9e, 0x60, 0x34,
	0xbb, 0xf9, 0x51, 0xb8, 0x9e, 0xdc, 0xf7, 0x32, 0x97, 0xfa, 0x79, 0xa8, 0xe0, 0xff, 0x6c, 0x01,
	0x26, 0x70, 0x95, 0x75, 0xcb, 0xe7, 0x6d, 0x3c, 0x9c, 0xaa, 0x57, 0xc5, 0x65, 0x23, 0xdc, 0x7c,
	0x0b, 0x4b, 0x42, 0x96, 0x11, 0xfd, 0x6f, 0xc7, 0xeb, 0xf5, 0xb6, 0xd4, 0xe1, 0xfb, 0x3c, 0x1e,
	0x9e, 0xf8, 0xf3, 0xb2, 0x33, 0x61, 0x06, 0xf8, 0xcd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d,
	0xd1, 0xc9, 0x08, 0x9f, 0x06, 0x00, 0x00,
}