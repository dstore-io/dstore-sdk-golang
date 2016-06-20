// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyPersonRelationship_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyPersonRelationship_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyPersonRelationship_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyPersonRelationship_Ad

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
	PersonId                     *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                 bool                          `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	RelatedPersonId              *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=related_person_id,json=relatedPersonId" json:"related_person_id,omitempty"`
	RelatedPersonIdNull          bool                          `protobuf:"varint,1002,opt,name=related_person_id_null,json=relatedPersonIdNull" json:"related_person_id_null,omitempty"`
	RelationshipId               *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull           bool                          `protobuf:"varint,1003,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	SortNo                       *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull                   bool                          `protobuf:"varint,1004,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	ValidFrom                    *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ValidFromNull                bool                          `protobuf:"varint,1005,opt,name=valid_from_null,json=validFromNull" json:"valid_from_null,omitempty"`
	ValidTo                      *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=valid_to,json=validTo" json:"valid_to,omitempty"`
	ValidToNull                  bool                          `protobuf:"varint,1006,opt,name=valid_to_null,json=validToNull" json:"valid_to_null,omitempty"`
	OldValidFrom                 *dstore_values.TimestampValue `protobuf:"bytes,7,opt,name=old_valid_from,json=oldValidFrom" json:"old_valid_from,omitempty"`
	OldValidFromNull             bool                          `protobuf:"varint,1007,opt,name=old_valid_from_null,json=oldValidFromNull" json:"old_valid_from_null,omitempty"`
	DeletePersonRelationship     *dstore_values.BooleanValue   `protobuf:"bytes,8,opt,name=delete_person_relationship,json=deletePersonRelationship" json:"delete_person_relationship,omitempty"`
	DeletePersonRelationshipNull bool                          `protobuf:"varint,1008,opt,name=delete_person_relationship_null,json=deletePersonRelationshipNull" json:"delete_person_relationship_null,omitempty"`
	AccessLevel                  *dstore_values.IntegerValue   `protobuf:"bytes,9,opt,name=access_level,json=accessLevel" json:"access_level,omitempty"`
	AccessLevelNull              bool                          `protobuf:"varint,1009,opt,name=access_level_null,json=accessLevelNull" json:"access_level_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
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

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetValidFrom() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *Parameters) GetValidTo() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidTo
	}
	return nil
}

func (m *Parameters) GetOldValidFrom() *dstore_values.TimestampValue {
	if m != nil {
		return m.OldValidFrom
	}
	return nil
}

func (m *Parameters) GetDeletePersonRelationship() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeletePersonRelationship
	}
	return nil
}

func (m *Parameters) GetAccessLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevel
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyPersonRelationship_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyPersonRelationship_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyPersonRelationship_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd5, 0x86, 0x24, 0xee, 0x34, 0x6d, 0xda, 0x2d, 0xaa, 0x4c, 0x4a, 0x39, 0xb4, 0x42,
	0x20, 0x81, 0x5c, 0x0e, 0x95, 0xa8, 0x10, 0x20, 0x71, 0x2a, 0xaa, 0xd4, 0x44, 0x91, 0x85, 0x22,
	0xc1, 0x8d, 0xe5, 0xc6, 0x9b, 0x60, 0xc9, 0xf6, 0x5a, 0x6b, 0xa7, 0x15, 0x8f, 0xc0, 0x1d, 0x4f,
	0xc4, 0xfb, 0x70, 0x86, 0x37, 0x60, 0xbc, 0x63, 0x27, 0xb6, 0x43, 0x49, 0xe1, 0xa6, 0xcd, 0x7a,
	0xff, 0xff, 0xff, 0x66, 0xd7, 0xbb, 0x63, 0x78, 0xe4, 0x44, 0xb1, 0x90, 0x7c, 0x87, 0x07, 0x43,
	0x37, 0xe0, 0x3b, 0xa1, 0x14, 0x7d, 0xee, 0x8c, 0x24, 0x8f, 0x76, 0x42, 0xdf, 0x6a, 0x0b, 0xc7,
	0x1d, 0xbc, 0xeb, 0x72, 0x19, 0x89, 0xc0, 0xe4, 0x9e, 0x1d, 0xbb, 0x22, 0x88, 0xde, 0xba, 0xa1,
	0xf5, 0xc4, 0x31, 0x50, 0x19, 0x0b, 0x76, 0x8b, 0xec, 0x06, 0xd9, 0x8d, 0xbf, 0x7b, 0x5a, 0x6b,
	0x29, 0xec, 0xd8, 0xf6, 0x46, 0x3c, 0xa2, 0x88, 0xd6, 0x85, 0x62, 0x05, 0x5c, 0x4a, 0x21, 0xd3,
	0xa9, 0x8d, 0xe2, 0x94, 0xcf, 0xa3, 0xc8, 0x1e, 0xf2, 0x74, 0x72, 0xbb, 0x3c, 0x19, 0xdb, 0x6e,
	0x30, 0x10, 0xd2, 0x57, 0x48, 0x12, 0x6d, 0xbd, 0xd7, 0x00, 0xba, 0xb6, 0xb4, 0x71, 0x16, 0x2b,
	0x62, 0x7b, 0xb0, 0x10, 0xaa, 0xca, 0x2c, 0xd7, 0xd1, 0xe7, 0xae, 0xcc, 0xdd, 0x58, 0xbc, 0xbb,
	0x61, 0xa4, 0x4b, 0x48, 0x8b, 0x72, 0x83, 0x98, 0x0f, 0xb9, 0xec, 0x25, 0x23, 0x53, 0x23, 0xf5,
	0x81, 0xc3, 0xae, 0xc1, 0xf2, 0xd8, 0x69, 0x05, 0x23, 0xcf, 0xd3, 0x3f, 0xd5, 0xd1, 0xaf, 0x99,
	0x8d, 0x4c, 0xd2, 0xc1, 0x87, 0xec, 0x25, 0xac, 0xca, 0x64, 0xd1, 0xdc, 0xb1, 0x26, 0xa0, 0xf9,
	0xd9, 0xa0, 0x66, 0xea, 0xea, 0x66, 0xbc, 0x5d, 0x58, 0x9f, 0x0a, 0x22, 0xee, 0x67, 0xe2, 0xae,
	0x95, 0x1c, 0x0a, 0xff, 0x1c, 0x28, 0x28, 0xdb, 0x73, 0x84, 0x57, 0x66, 0xc3, 0x97, 0xf3, 0x1e,
	0x64, 0xdf, 0x81, 0xf3, 0xa5, 0x14, 0x22, 0x7f, 0x21, 0x32, 0x2b, 0xca, 0x15, 0x78, 0x17, 0xea,
	0x91, 0x90, 0xb1, 0x15, 0x08, 0xfd, 0xdc, 0x6c, 0x60, 0x2d, 0xd1, 0x76, 0x04, 0xbb, 0x0a, 0x8d,
	0xd4, 0x45, 0x80, 0xaf, 0x04, 0x00, 0x9a, 0x56, 0xc1, 0x0f, 0x01, 0x30, 0x01, 0x2b, 0x18, 0x48,
	0xe1, 0xeb, 0x55, 0x95, 0xbd, 0x59, 0xca, 0x8e, 0x5d, 0x3c, 0x19, 0xb1, 0xed, 0x87, 0x94, 0xbe,
	0xa0, 0x0c, 0xfb, 0xa8, 0x67, 0xd7, 0xa1, 0x39, 0x71, 0x13, 0xe3, 0x1b, 0x31, 0x96, 0xc6, 0x22,
	0x85, 0xd9, 0x03, 0x8d, 0x84, 0xb1, 0xd0, 0x6b, 0x67, 0x81, 0xd4, 0x95, 0xfc, 0x95, 0x60, 0xdb,
	0xb0, 0x94, 0x39, 0x09, 0xf0, 0x9d, 0x00, 0x8b, 0xa9, 0x40, 0xc5, 0x3f, 0x83, 0x65, 0xe1, 0x39,
	0x56, 0x6e, 0x25, 0xf5, 0xb3, 0x40, 0x1a, 0x68, 0xea, 0x8d, 0x17, 0x63, 0xc0, 0x5a, 0x31, 0x84,
	0x78, 0x3f, 0x88, 0xb7, 0x92, 0xd7, 0x2a, 0xe8, 0x6b, 0x68, 0x39, 0xdc, 0xc3, 0x83, 0x9f, 0x9d,
	0xa0, 0xfc, 0x7b, 0xd3, 0xb5, 0x3f, 0xbe, 0xa6, 0x23, 0x21, 0x3c, 0x6e, 0x07, 0x84, 0xd7, 0xc9,
	0x3e, 0x7d, 0x97, 0xd9, 0x3e, 0x5c, 0x3e, 0x3d, 0x9a, 0xca, 0xfa, 0x49, 0x65, 0x5d, 0x3c, 0x2d,
	0x43, 0x95, 0xf8, 0x18, 0x1a, 0x76, 0xbf, 0x8f, 0xd7, 0xda, 0xf2, 0xf8, 0x31, 0xf7, 0xf4, 0x85,
	0xd9, 0x67, 0x67, 0x91, 0x0c, 0x87, 0x89, 0x9e, 0xdd, 0x84, 0xd5, 0xbc, 0x9f, 0xc8, 0xbf, 0x88,
	0xdc, 0xcc, 0x09, 0x13, 0xd8, 0xd6, 0xc7, 0x79, 0xd0, 0x4c, 0x1e, 0x85, 0xc8, 0xe7, 0xec, 0x36,
	0x54, 0x55, 0xa7, 0x49, 0xbb, 0x40, 0xcb, 0x28, 0x36, 0x32, 0xea, 0x42, 0x2f, 0x92, 0xbf, 0x26,
	0x09, 0x71, 0x3b, 0x57, 0x92, 0x1e, 0x63, 0xe5, 0x9a, 0x0c, 0xde, 0xec, 0x0a, 0x9a, 0x8d, 0x92,
	0xb9, 0xdc, 0x8a, 0xda, 0x38, 0x3e, 0x98, 0x8c, 0xcd, 0xa6, 0x5f, 0x7c, 0x80, 0xa7, 0xaf, 0x9e,
	0xf6, 0x36, 0xbc, 0xae, 0x49, 0xe2, 0xa5, 0xa9, 0x44, 0xea, 0x7c, 0x6d, 0xfa, 0x6f, 0x66, 0x72,
	0x76, 0x08, 0x15, 0x29, 0x4e, 0xf0, 0xce, 0x25, 0xae, 0x07, 0xc6, 0xbf, 0x74, 0x63, 0x23, 0xdb,
	0x0b, 0xc3, 0x14, 0x27, 0x66, 0x12, 0xd3, 0xda, 0x84, 0x0a, 0xfe, 0x66, 0xeb, 0x50, 0xc3, 0x51,
	0xd2, 0x3c, 0x3e, 0x74, 0x70, 0x77, 0xaa, 0x66, 0x15, 0x87, 0x07, 0xce, 0xd3, 0x1e, 0x6c, 0xb8,
	0xa2, 0xcc, 0x18, 0x7f, 0x30, 0xde, 0xdc, 0xff, 0xcf, 0x4f, 0xc9, 0x51, 0x4d, 0xf5, 0xea, 0x7b,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x25, 0x14, 0x5e, 0x01, 0x8c, 0x06, 0x00, 0x00,
}
