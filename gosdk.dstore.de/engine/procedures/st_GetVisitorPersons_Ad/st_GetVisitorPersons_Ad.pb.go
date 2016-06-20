// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetVisitorPersons_Ad.proto
// DO NOT EDIT!

/*
Package st_GetVisitorPersons_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetVisitorPersons_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetVisitorPersons_Ad

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
	VisitorId                     *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=visitor_id,json=visitorId" json:"visitor_id,omitempty"`
	VisitorIdNull                 bool                          `protobuf:"varint,1001,opt,name=visitor_id_null,json=visitorIdNull" json:"visitor_id_null,omitempty"`
	FromDate                      *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                  bool                          `protobuf:"varint,1002,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                        *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                    bool                          `protobuf:"varint,1003,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	PersonId                      *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                  bool                          `protobuf:"varint,1004,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	VisitorOrPersonIdsInOneId     *dstore_values.IntegerValue   `protobuf:"bytes,5,opt,name=visitor_or_person_ids_in_one_id,json=visitorOrPersonIdsInOneId" json:"visitor_or_person_ids_in_one_id,omitempty"`
	VisitorOrPersonIdsInOneIdNull bool                          `protobuf:"varint,1005,opt,name=visitor_or_person_ids_in_one_id_null,json=visitorOrPersonIdsInOneIdNull" json:"visitor_or_person_ids_in_one_id_null,omitempty"`
	OutputIntoOneId               *dstore_values.BooleanValue   `protobuf:"bytes,6,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull           bool                          `protobuf:"varint,1006,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetVisitorId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VisitorId
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

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetVisitorOrPersonIdsInOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VisitorOrPersonIdsInOneId
	}
	return nil
}

func (m *Parameters) GetOutputIntoOneId() *dstore_values.BooleanValue {
	if m != nil {
		return m.OutputIntoOneId
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
	CharacVal2RestrByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=charac_val2_restr_by_pattern,json=characVal2RestrByPattern" json:"charac_val2_restr_by_pattern,omitempty"`
	VisitorId                *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=visitor_id,json=visitorId" json:"visitor_id,omitempty"`
	PersonId                 *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	ValidTo                  *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=valid_to,json=validTo" json:"valid_to,omitempty"`
	ValidFrom                *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	PersonType               *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=person_type,json=personType" json:"person_type,omitempty"`
	CharacteristicValue2     *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=characteristic_value2,json=characteristicValue2" json:"characteristic_value2,omitempty"`
	CharacVal1RestrByPattern *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=charac_val1_restr_by_pattern,json=characVal1RestrByPattern" json:"charac_val1_restr_by_pattern,omitempty"`
	CharacteristicValue1     *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=characteristic_value1,json=characteristicValue1" json:"characteristic_value1,omitempty"`
	PersonTypeId             *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCharacVal2RestrByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.CharacVal2RestrByPattern
	}
	return nil
}

func (m *Response_Row) GetVisitorId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VisitorId
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetValidTo() *dstore_values.StringValue {
	if m != nil {
		return m.ValidTo
	}
	return nil
}

func (m *Response_Row) GetValidFrom() *dstore_values.StringValue {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *Response_Row) GetPersonType() *dstore_values.StringValue {
	if m != nil {
		return m.PersonType
	}
	return nil
}

func (m *Response_Row) GetCharacteristicValue2() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicValue2
	}
	return nil
}

func (m *Response_Row) GetCharacVal1RestrByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.CharacVal1RestrByPattern
	}
	return nil
}

func (m *Response_Row) GetCharacteristicValue1() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicValue1
	}
	return nil
}

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetVisitorPersons_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetVisitorPersons_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetVisitorPersons_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetVisitorPersons_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xeb, 0x4e, 0x13, 0x4d,
	0x18, 0x0e, 0x5f, 0xbf, 0x9e, 0x06, 0x14, 0xb3, 0x78, 0x58, 0x8a, 0x78, 0x40, 0x89, 0xfe, 0x5a,
	0x6c, 0x89, 0x44, 0xd4, 0x3f, 0x12, 0x51, 0x9b, 0x48, 0x69, 0x36, 0x84, 0x78, 0xf8, 0xb1, 0x59,
	0xd8, 0xa1, 0x6e, 0x6c, 0x67, 0x36, 0x3b, 0x53, 0x08, 0xff, 0xbd, 0x00, 0xcf, 0xa7, 0xab, 0xf1,
	0x76, 0xf0, 0x74, 0x0d, 0xbe, 0x33, 0xef, 0xb4, 0x65, 0x17, 0xc9, 0xb6, 0x26, 0x84, 0x66, 0x76,
	0x9e, 0xe7, 0x7d, 0x9e, 0x7d, 0x4f, 0x4b, 0x96, 0x02, 0x21, 0x79, 0x4c, 0x17, 0x28, 0x6b, 0x85,
	0x8c, 0x2e, 0x44, 0x31, 0xdf, 0xa6, 0x41, 0x37, 0xa6, 0x62, 0x41, 0x48, 0xef, 0x21, 0x95, 0x9b,
	0xa1, 0x08, 0x01, 0xd0, 0xa4, 0xb1, 0xe0, 0x4c, 0x78, 0xf7, 0x02, 0x07, 0x20, 0x92, 0x5b, 0xf3,
	0xc8, 0x73, 0x90, 0xe7, 0x1c, 0x03, 0xae, 0x4c, 0x99, 0xf0, 0xbb, 0x7e, 0xbb, 0x4b, 0x05, 0x72,
	0x2b, 0xd3, 0x49, 0x4d, 0x1a, 0xc7, 0x3c, 0x36, 0x57, 0x33, 0xc9, 0xab, 0x0e, 0x15, 0xc2, 0x6f,
	0x51, 0x73, 0x79, 0x25, 0x7d, 0x29, 0xfd, 0x90, 0xed, 0xf0, 0xb8, 0xe3, 0xcb, 0x90, 0x33, 0x04,
	0xcd, 0x7d, 0xcb, 0x13, 0xd2, 0xf4, 0x63, 0x1f, 0x6e, 0xc1, 0x86, 0x75, 0x9b, 0x90, 0x5d, 0x74,
	0xe5, 0x85, 0x81, 0x3d, 0x76, 0x69, 0xec, 0xfa, 0x78, 0x6d, 0xc6, 0x31, 0xe6, 0x8d, 0xab, 0x90,
	0x49, 0xda, 0xa2, 0xf1, 0xa6, 0x3a, 0xb9, 0x65, 0x03, 0xaf, 0x07, 0xd6, 0x35, 0x32, 0x39, 0xe0,
	0x7a, 0xac, 0xdb, 0x6e, 0xdb, 0x07, 0x45, 0x88, 0x50, 0x72, 0x4f, 0xf4, 0x41, 0x0d, 0x78, 0x0a,
	0x22, 0xe5, 0x9d, 0x98, 0x77, 0xbc, 0xc0, 0x97, 0xd4, 0xfe, 0x4f, 0x6b, 0xcc, 0xa6, 0x34, 0x64,
	0x08, 0xef, 0x22, 0xfd, 0x4e, 0x84, 0x2a, 0x25, 0x85, 0xbf, 0x0f, 0x70, 0x6b, 0x9e, 0x9c, 0xec,
	0x73, 0x51, 0xe3, 0x3b, 0x6a, 0x4c, 0xf4, 0x20, 0x5a, 0x62, 0x89, 0x14, 0x25, 0x47, 0x81, 0xdc,
	0x30, 0x02, 0x05, 0xc9, 0x75, 0xf8, 0xcb, 0x64, 0xc2, 0xf0, 0x30, 0xf8, 0x0f, 0x0c, 0x4e, 0xf0,
	0x5a, 0x87, 0xbe, 0x45, 0xca, 0x91, 0xae, 0x98, 0xca, 0xd0, 0xff, 0xd9, 0x19, 0x2a, 0x21, 0x1a,
	0x12, 0x04, 0xde, 0xfb, 0x4c, 0x0c, 0xff, 0xd3, 0x78, 0xef, 0x41, 0xb4, 0x80, 0x47, 0x2e, 0xf6,
	0xf2, 0x08, 0x7f, 0x7d, 0x86, 0xf0, 0x42, 0xe6, 0x71, 0x46, 0x95, 0x6c, 0x3e, 0x5b, 0x76, 0xda,
	0xc4, 0x58, 0x37, 0xfd, 0x55, 0x0f, 0x44, 0x9d, 0xad, 0x33, 0x0a, 0x3e, 0x1e, 0x93, 0xab, 0x19,
	0x02, 0xe8, 0xee, 0x17, 0xba, 0x9b, 0x3d, 0x36, 0x92, 0xb6, 0xfb, 0x88, 0x58, 0xbc, 0x2b, 0xa3,
	0xae, 0x04, 0x3a, 0xe4, 0xce, 0x38, 0x2c, 0xfc, 0xd5, 0xe1, 0x16, 0xe7, 0x6d, 0xea, 0x33, 0x74,
	0x38, 0x89, 0xb4, 0x3a, 0xb0, 0xd0, 0xd7, 0x4d, 0x72, 0xee, 0x68, 0x24, 0xb4, 0xf2, 0x1b, 0xad,
	0x4c, 0xa5, 0x28, 0xca, 0xc0, 0xdc, 0xab, 0x12, 0x29, 0xb9, 0x54, 0x44, 0x30, 0x43, 0xd4, 0xba,
	0x41, 0xf2, 0x7a, 0x40, 0x4c, 0xef, 0x56, 0x9c, 0xe4, 0xe0, 0xe1, 0xf0, 0xac, 0xaa, 0xff, 0x2e,
	0x02, 0xad, 0xa7, 0xe4, 0x94, 0x1a, 0x0d, 0xef, 0xd0, 0x6c, 0x40, 0x53, 0xe6, 0x80, 0xec, 0xa4,
	0xc8, 0xe9, 0x09, 0x5a, 0x83, 0x73, 0x7d, 0x70, 0x76, 0x27, 0x3b, 0xc9, 0x07, 0xd0, 0x2a, 0x45,
	0x33, 0x92, 0xd0, 0x85, 0x2a, 0xe2, 0x85, 0x23, 0x11, 0x71, 0x60, 0xd7, 0xf0, 0xd7, 0xed, 0xc1,
	0xad, 0x55, 0x92, 0x8b, 0xf9, 0x1e, 0xb4, 0x97, 0x62, 0x2d, 0x3a, 0x43, 0x6d, 0x0f, 0xa7, 0x97,
	0x04, 0xc7, 0xe5, 0x7b, 0xae, 0xe2, 0x57, 0x0e, 0xf2, 0x24, 0x07, 0x07, 0xeb, 0x2c, 0x29, 0xc0,
	0x51, 0xd5, 0xe5, 0x75, 0x03, 0xf2, 0x92, 0x77, 0xf3, 0x70, 0x84, 0x8c, 0x3f, 0x27, 0xe7, 0xb7,
	0x5f, 0xc0, 0xf4, 0x6f, 0x7b, 0x50, 0xa0, 0x9a, 0x07, 0x7b, 0x4c, 0xc6, 0xde, 0xd6, 0xbe, 0x17,
	0xf9, 0x12, 0xd6, 0x01, 0xb3, 0xdf, 0x34, 0x92, 0x59, 0x34, 0x65, 0x04, 0x54, 0xc8, 0x5a, 0x58,
	0x45, 0x1b, 0x03, 0xc0, 0xa1, 0xe6, 0x2a, 0xfa, 0xca, 0x7e, 0x13, 0xc9, 0xd6, 0x9d, 0xc4, 0x2e,
	0x79, 0xdb, 0x18, 0x69, 0x99, 0x2c, 0x1f, 0x9e, 0xb2, 0x77, 0x8d, 0x51, 0xc6, 0x6c, 0x89, 0x94,
	0x00, 0x00, 0x9d, 0x23, 0xb9, 0xfd, 0x3e, 0xfb, 0x05, 0x8a, 0x1a, 0xbc, 0xc1, 0xf5, 0xee, 0xd3,
	0x3c, 0xb5, 0x49, 0xec, 0x0f, 0xd9, 0xcc, 0xb2, 0x86, 0x3f, 0x00, 0xb4, 0x75, 0x97, 0x8c, 0x1b,
	0xbb, 0x72, 0x3f, 0xa2, 0xf6, 0xc7, 0x6c, 0x32, 0x41, 0xfc, 0x06, 0xc0, 0xad, 0x26, 0x39, 0x83,
	0x59, 0x84, 0xbc, 0x85, 0x42, 0x86, 0xba, 0x1c, 0x5d, 0x5a, 0xb3, 0x3f, 0x65, 0xc7, 0x39, 0x9d,
	0x64, 0xea, 0x87, 0xb5, 0x64, 0x61, 0xab, 0x47, 0x0b, 0xfb, 0x79, 0x94, 0xc2, 0x56, 0x53, 0x85,
	0x3d, 0xc6, 0x6e, 0xd5, 0xfe, 0xf2, 0x6f, 0x76, 0xab, 0xd6, 0x4a, 0x7f, 0x33, 0xaa, 0xf4, 0xa9,
	0x92, 0x7f, 0x1d, 0xa2, 0xe4, 0x13, 0x83, 0x14, 0xd6, 0x83, 0x95, 0x27, 0x64, 0x26, 0xe4, 0xa9,
	0x49, 0x19, 0x7c, 0x9f, 0x9f, 0x2d, 0xb7, 0xb8, 0x08, 0x5e, 0xf6, 0xee, 0x83, 0x11, 0x3e, 0xe1,
	0x5b, 0x05, 0xfd, 0xa9, 0x5c, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x06, 0x4a, 0x20, 0xfd,
	0x07, 0x00, 0x00,
}