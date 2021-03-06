// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPersonRelationships_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPersonRelationships_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPersonRelationships_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPersonRelationships_Ad

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
	PersonId                        *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                    bool                          `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	RelatedPersonId                 *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=related_person_id,json=relatedPersonId" json:"related_person_id,omitempty"`
	RelatedPersonIdNull             bool                          `protobuf:"varint,1002,opt,name=related_person_id_null,json=relatedPersonIdNull" json:"related_person_id_null,omitempty"`
	RelationshipId                  *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	RelationshipIdNull              bool                          `protobuf:"varint,1003,opt,name=relationship_id_null,json=relationshipIdNull" json:"relationship_id_null,omitempty"`
	FromDate                        *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                    bool                          `protobuf:"varint,1004,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                          *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                      bool                          `protobuf:"varint,1005,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	OnlyRelationsCurrentlyValid     *dstore_values.BooleanValue   `protobuf:"bytes,6,opt,name=only_relations_currently_valid,json=onlyRelationsCurrentlyValid" json:"only_relations_currently_valid,omitempty"`
	OnlyRelationsCurrentlyValidNull bool                          `protobuf:"varint,1006,opt,name=only_relations_currently_valid_null,json=onlyRelationsCurrentlyValidNull" json:"only_relations_currently_valid_null,omitempty"`
	RelatedPersonTypeId             *dstore_values.IntegerValue   `protobuf:"bytes,7,opt,name=related_person_type_id,json=relatedPersonTypeId" json:"related_person_type_id,omitempty"`
	RelatedPersonTypeIdNull         bool                          `protobuf:"varint,1007,opt,name=related_person_type_id_null,json=relatedPersonTypeIdNull" json:"related_person_type_id_null,omitempty"`
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

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetOnlyRelationsCurrentlyValid() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyRelationsCurrentlyValid
	}
	return nil
}

func (m *Parameters) GetRelatedPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelatedPersonTypeId
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
	RowId                       int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	RelatedPersonId             *dstore_values.IntegerValue   `protobuf:"bytes,10001,opt,name=related_person_id,json=relatedPersonId" json:"related_person_id,omitempty"`
	PersonId                    *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	RelationshipId              *dstore_values.IntegerValue   `protobuf:"bytes,10003,opt,name=relationship_id,json=relationshipId" json:"relationship_id,omitempty"`
	PersonTypeIdOfRelatedPerson *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=person_type_id_of_related_person,json=personTypeIdOfRelatedPerson" json:"person_type_id_of_related_person,omitempty"`
	ValidTo                     *dstore_values.TimestampValue `protobuf:"bytes,10005,opt,name=valid_to,json=validTo" json:"valid_to,omitempty"`
	ValidToChar                 *dstore_values.StringValue    `protobuf:"bytes,10006,opt,name=valid_to_char,json=validToChar" json:"valid_to_char,omitempty"`
	ValidFrom                   *dstore_values.TimestampValue `protobuf:"bytes,10007,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	SortNo                      *dstore_values.IntegerValue   `protobuf:"bytes,10008,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	AccessLevel                 *dstore_values.IntegerValue   `protobuf:"bytes,10009,opt,name=access_level,json=accessLevel" json:"access_level,omitempty"`
	ValidFromChar               *dstore_values.StringValue    `protobuf:"bytes,10010,opt,name=valid_from_char,json=validFromChar" json:"valid_from_char,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetRelatedPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelatedPersonId
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetRelationshipId() *dstore_values.IntegerValue {
	if m != nil {
		return m.RelationshipId
	}
	return nil
}

func (m *Response_Row) GetPersonTypeIdOfRelatedPerson() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeIdOfRelatedPerson
	}
	return nil
}

func (m *Response_Row) GetValidTo() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidTo
	}
	return nil
}

func (m *Response_Row) GetValidToChar() *dstore_values.StringValue {
	if m != nil {
		return m.ValidToChar
	}
	return nil
}

func (m *Response_Row) GetValidFrom() *dstore_values.TimestampValue {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetAccessLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.AccessLevel
	}
	return nil
}

func (m *Response_Row) GetValidFromChar() *dstore_values.StringValue {
	if m != nil {
		return m.ValidFromChar
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPersonRelationships_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPersonRelationships_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPersonRelationships_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_GetPersonRelationships_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xeb, 0x6e, 0xd3, 0x4a,
	0x10, 0x56, 0x4f, 0x4e, 0x2e, 0xdd, 0xde, 0xce, 0x71, 0x8f, 0x7a, 0x4c, 0x22, 0x4a, 0x69, 0x85,
	0x84, 0xf8, 0xe1, 0x72, 0x57, 0x01, 0x01, 0x82, 0xb6, 0x94, 0x22, 0x6a, 0x22, 0xab, 0xaa, 0x04,
	0x12, 0x32, 0x6e, 0xbc, 0x49, 0x2d, 0x1c, 0xaf, 0xb5, 0xde, 0xb4, 0xea, 0x5b, 0x70, 0xbf, 0x3d,
	0x10, 0x6f, 0x02, 0x12, 0xf7, 0x57, 0x60, 0xd6, 0xb3, 0x4e, 0x63, 0x27, 0x6a, 0x5c, 0xfe, 0xb4,
	0xda, 0x9d, 0xf9, 0xbe, 0x6f, 0x66, 0x76, 0x66, 0x62, 0x72, 0xcd, 0x8d, 0x04, 0xe3, 0x74, 0x91,
	0x06, 0x2d, 0x2f, 0xa0, 0x8b, 0x21, 0x67, 0x0d, 0xea, 0x76, 0x38, 0x8d, 0x16, 0xc3, 0xb6, 0xbd,
	0x46, 0x45, 0x9d, 0xf2, 0x88, 0x05, 0x16, 0xf5, 0x1d, 0xe1, 0xb1, 0x20, 0xda, 0xf1, 0xc2, 0xc8,
	0xbe, 0xe5, 0x1a, 0xe0, 0x27, 0x98, 0x76, 0x06, 0xc1, 0x06, 0x82, 0x8d, 0xc3, 0x10, 0xd5, 0x69,
	0x25, 0xb4, 0xeb, 0xf8, 0x1d, 0x1a, 0x21, 0x41, 0xf5, 0x58, 0x5a, 0x9d, 0x72, 0xce, 0xb8, 0x32,
	0xd5, 0xd2, 0xa6, 0x36, 0x8d, 0x22, 0xa7, 0x45, 0x95, 0x71, 0x21, 0x6b, 0x14, 0x8e, 0x17, 0x34,
	0x19, 0x6f, 0xc7, 0x9a, 0xe8, 0x34, 0xff, 0xa9, 0x44, 0x48, 0xdd, 0xe1, 0x0e, 0x58, 0x21, 0x24,
	0x6d, 0x89, 0x8c, 0x86, 0x71, 0x68, 0xb6, 0xe7, 0xea, 0x23, 0x73, 0x23, 0xa7, 0xc7, 0xce, 0xd7,
	0x0c, 0x95, 0x80, 0x0a, 0xca, 0x0b, 0x04, 0x6d, 0x51, 0xbe, 0x25, 0x4f, 0x56, 0x05, 0xbd, 0xd7,
	0x5d, 0xed, 0x14, 0x99, 0xec, 0x22, 0xed, 0xa0, 0xe3, 0xfb, 0xfa, 0x97, 0x32, 0xe0, 0x2b, 0xd6,
	0x78, 0xe2, 0x62, 0xc2, 0xa5, 0xb6, 0x46, 0xfe, 0xe5, 0x32, 0x6b, 0xea, 0xda, 0x07, 0x42, 0x7f,
	0x0d, 0x17, 0x9a, 0x52, 0xa8, 0x7a, 0xa2, 0x77, 0x91, 0xcc, 0xf4, 0x11, 0xa1, 0xee, 0x57, 0xd4,
	0x9d, 0xce, 0x20, 0x62, 0xf9, 0x15, 0x82, 0x44, 0xaa, 0xe8, 0x52, 0xbc, 0x30, 0x5c, 0x7c, 0xb2,
	0x17, 0x03, 0xda, 0xe7, 0xc8, 0x7f, 0x19, 0x16, 0x54, 0xfe, 0x86, 0xca, 0x5a, 0xda, 0x3d, 0x16,
	0xbe, 0x4a, 0x46, 0x9b, 0x9c, 0xb5, 0x6d, 0x17, 0x42, 0xd2, 0xff, 0x8e, 0x25, 0x8f, 0x67, 0x24,
	0x85, 0x07, 0xef, 0x27, 0x9c, 0x76, 0xa8, 0x4a, 0x2b, 0xfd, 0x57, 0xc0, 0x5d, 0x96, 0xb6, 0x8b,
	0x45, 0xa1, 0xef, 0xaa, 0xb4, 0x89, 0x4b, 0x2c, 0x71, 0x99, 0x94, 0x05, 0x43, 0x81, 0x62, 0x1e,
	0x81, 0x92, 0x60, 0x31, 0xfd, 0x49, 0x32, 0xae, 0x70, 0x48, 0xfe, 0x03, 0xc9, 0x09, 0x9a, 0x63,
	0xea, 0x27, 0x64, 0x96, 0x05, 0xfe, 0xbe, 0xdd, 0x4d, 0xcc, 0x6e, 0x74, 0x38, 0xa7, 0x81, 0x80,
	0x3b, 0x20, 0x87, 0x2a, 0x96, 0x06, 0x56, 0x71, 0x9b, 0x31, 0x9f, 0x3a, 0x01, 0xea, 0xd5, 0x24,
	0x45, 0xb7, 0xe5, 0x97, 0x13, 0x82, 0x2d, 0x89, 0xd7, 0x36, 0xc8, 0xc2, 0xe1, 0x0a, 0x18, 0xdb,
	0x4f, 0x8c, 0xed, 0xc4, 0x21, 0x54, 0x71, 0xc0, 0xf5, 0xbe, 0xee, 0x10, 0xfb, 0x21, 0x95, 0xcf,
	0x5d, 0x1e, 0xfe, 0xdc, 0xe9, 0xce, 0xd9, 0x04, 0x20, 0xbc, 0xf9, 0x75, 0x52, 0x1b, 0xcc, 0x88,
	0x81, 0xfd, 0xc2, 0xc0, 0xfe, 0x1f, 0x00, 0x95, 0x01, 0xcd, 0x7f, 0x2e, 0x93, 0x8a, 0x45, 0xa3,
	0x10, 0xc2, 0xa5, 0xda, 0x59, 0x52, 0x8c, 0xa7, 0x58, 0x4d, 0x58, 0xd5, 0x48, 0xaf, 0x08, 0x9c,
	0xf0, 0x55, 0xf9, 0xd7, 0x42, 0x47, 0xed, 0x21, 0xf9, 0x47, 0xce, 0xaf, 0xdd, 0x33, 0xc0, 0x30,
	0x35, 0x05, 0x00, 0x1b, 0x19, 0x70, 0x76, 0xcc, 0x37, 0xe0, 0xbc, 0x7e, 0x70, 0xb6, 0xa6, 0xda,
	0xe9, 0x0b, 0x18, 0xf9, 0xb2, 0xda, 0x1b, 0x30, 0x0a, 0x92, 0x71, 0xb6, 0x8f, 0x11, 0xb7, 0xca,
	0x06, 0xfe, 0xb7, 0x12, 0x77, 0xed, 0x1e, 0x29, 0x70, 0xb6, 0x07, 0xdd, 0x2c, 0x51, 0x4b, 0x46,
	0xfe, 0x3d, 0x67, 0x24, 0x95, 0x30, 0x2c, 0xb6, 0x67, 0x49, 0x92, 0xea, 0xc7, 0x22, 0x29, 0xc0,
	0x41, 0x9b, 0x21, 0x25, 0x38, 0xca, 0x87, 0x7a, 0x66, 0x42, 0x71, 0x8a, 0x56, 0x11, 0x8e, 0x50,
	0xfe, 0xbb, 0x83, 0xf6, 0xc6, 0x73, 0xf3, 0x0f, 0x16, 0xc7, 0x95, 0xde, 0x15, 0xf7, 0xc2, 0x3c,
	0xca, 0x8e, 0x5b, 0xed, 0xdf, 0x1e, 0x2f, 0xcd, 0xa3, 0xaf, 0x8f, 0x06, 0x99, 0xcb, 0xb4, 0x10,
	0x6b, 0xda, 0xe9, 0xec, 0xf4, 0x57, 0x39, 0x78, 0x6b, 0x61, 0x4f, 0x97, 0x3d, 0x68, 0x5a, 0xbd,
	0x89, 0x42, 0x9a, 0x15, 0x9c, 0x1b, 0xc1, 0xf4, 0xd7, 0x66, 0x9e, 0x7d, 0x50, 0x8e, 0xfd, 0x37,
	0x99, 0x76, 0x93, 0x4c, 0x24, 0x50, 0xbb, 0xb1, 0xe3, 0x70, 0xfd, 0x8d, 0x99, 0xee, 0x53, 0x85,
	0x8f, 0x04, 0xf7, 0x82, 0x16, 0x82, 0xc7, 0x14, 0x78, 0x19, 0xfc, 0x61, 0x56, 0x08, 0x12, 0xc8,
	0xfd, 0xa4, 0xbf, 0xcd, 0xa5, 0x3e, 0x1a, 0x23, 0xee, 0x00, 0x40, 0xbb, 0x44, 0xca, 0x11, 0xe3,
	0xc2, 0x0e, 0x98, 0xfe, 0x2e, 0x47, 0x19, 0x4a, 0xd2, 0xd9, 0x94, 0x61, 0x8f, 0x3b, 0x8d, 0x06,
	0xf4, 0xa6, 0xed, 0xd3, 0x5d, 0xea, 0xeb, 0xef, 0x73, 0x60, 0xc7, 0x10, 0x71, 0x5f, 0x02, 0xb4,
	0x65, 0x32, 0x75, 0x10, 0x36, 0x66, 0xfe, 0x61, 0x78, 0xe6, 0x13, 0xdd, 0xc0, 0x65, 0xee, 0xb7,
	0x1f, 0x93, 0x9a, 0xc7, 0xb2, 0xb3, 0xd0, 0xfd, 0x60, 0x78, 0x74, 0xa3, 0xc5, 0x22, 0xf7, 0x69,
	0x62, 0x77, 0x8f, 0xfa, 0x4d, 0xb1, 0x5d, 0x8a, 0x7f, 0xb6, 0x2f, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x77, 0x98, 0x68, 0x89, 0x93, 0x08, 0x00, 0x00,
}
