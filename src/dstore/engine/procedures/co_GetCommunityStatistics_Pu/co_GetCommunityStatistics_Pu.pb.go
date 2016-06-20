// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunityStatistics_Pu.proto
// DO NOT EDIT!

/*
Package co_GetCommunityStatistics_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunityStatistics_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunityStatistics_Pu

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
	UniqueId                          *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                      bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues        *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull    bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                   bool                        `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	ListOfStatisticInformationIds     *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=list_of_statistic_information_ids,json=listOfStatisticInformationIds" json:"list_of_statistic_information_ids,omitempty"`
	ListOfStatisticInformationIdsNull bool                        `protobuf:"varint,1004,opt,name=list_of_statistic_information_ids_null,json=listOfStatisticInformationIdsNull" json:"list_of_statistic_information_ids_null,omitempty"`
	SeparatorInIdentVals              *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull          bool                        `protobuf:"varint,1005,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetListOfStatisticInformationIds() *dstore_values.StringValue {
	if m != nil {
		return m.ListOfStatisticInformationIds
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
	RowId                  int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value                  *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value" json:"value,omitempty"`
	StatisticInformationId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=statistic_information_id,json=statisticInformationId" json:"statistic_information_id,omitempty"`
	StatisticInformation   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=statistic_information,json=statisticInformation" json:"statistic_information,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetStatisticInformationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.StatisticInformationId
	}
	return nil
}

func (m *Response_Row) GetStatisticInformation() *dstore_values.StringValue {
	if m != nil {
		return m.StatisticInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunityStatistics_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunityStatistics_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunityStatistics_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x6d, 0x6b, 0x13, 0x41,
	0x10, 0xa6, 0xa6, 0xd7, 0xd6, 0x69, 0xb1, 0xba, 0x6a, 0x3d, 0xd3, 0x5a, 0x4c, 0x45, 0x11, 0x85,
	0x8b, 0x2f, 0x88, 0x05, 0x41, 0x41, 0x11, 0x89, 0xd0, 0x18, 0x4f, 0x29, 0x28, 0xc2, 0x71, 0xe6,
	0x36, 0x61, 0x21, 0xd9, 0x8d, 0xbb, 0x7b, 0x16, 0xff, 0x85, 0x2f, 0xdf, 0xfc, 0x57, 0xfe, 0x0c,
	0xdf, 0xfe, 0x80, 0x9f, 0x9c, 0xbd, 0xb9, 0xa4, 0xc9, 0x99, 0xe4, 0x8a, 0x5f, 0x12, 0x76, 0x67,
	0x9e, 0x67, 0x9e, 0x9b, 0x79, 0x66, 0xe1, 0x5e, 0x62, 0xac, 0xd2, 0xbc, 0xce, 0x65, 0x57, 0x48,
	0x5e, 0x1f, 0x68, 0xd5, 0xe6, 0x49, 0xaa, 0xb9, 0xa9, 0xb7, 0x55, 0xf4, 0x84, 0xdb, 0x47, 0xaa,
	0xdf, 0x4f, 0xa5, 0xb0, 0x1f, 0x5e, 0xd8, 0xd8, 0x0a, 0x63, 0x45, 0xdb, 0x44, 0xad, 0x34, 0xc0,
	0x3c, 0xab, 0xd8, 0x35, 0x02, 0x07, 0x04, 0x0e, 0xe6, 0x21, 0xaa, 0xa7, 0xf3, 0x42, 0xef, 0xe3,
	0x5e, 0xca, 0x0d, 0x11, 0x54, 0xcf, 0x4f, 0x56, 0xe7, 0x5a, 0x2b, 0x9d, 0x87, 0x36, 0x27, 0x43,
	0x7d, 0x6e, 0x4c, 0xdc, 0xe5, 0x79, 0xf0, 0x52, 0x31, 0x68, 0x63, 0x21, 0x3b, 0x4a, 0xf7, 0xb1,
	0xa2, 0x92, 0x94, 0xb4, 0xf3, 0xcd, 0x03, 0x68, 0xc5, 0x3a, 0xc6, 0x28, 0xd7, 0x86, 0xdd, 0x85,
	0xe3, 0x28, 0xeb, 0x5d, 0xca, 0x23, 0x91, 0xf8, 0x0b, 0x17, 0x17, 0xae, 0xae, 0xde, 0xaa, 0x06,
	0xf9, 0x07, 0xe4, 0xa2, 0x8c, 0xd5, 0x42, 0x76, 0xf7, 0xdd, 0x21, 0x5c, 0xa1, 0xe4, 0x46, 0xc2,
	0x2e, 0xc3, 0x89, 0x11, 0x30, 0x92, 0x69, 0xaf, 0xe7, 0x7f, 0x5f, 0x46, 0xf8, 0x4a, 0xb8, 0x36,
	0x4c, 0x69, 0xe2, 0x25, 0x7b, 0x03, 0x5b, 0x03, 0xac, 0xa3, 0x24, 0xa6, 0x71, 0x69, 0x45, 0x47,
	0xb4, 0x33, 0x35, 0x11, 0x91, 0xfb, 0xc7, 0x4a, 0x4b, 0x56, 0x09, 0xdf, 0x98, 0x80, 0x67, 0x21,
	0xc3, 0x9e, 0x42, 0x6d, 0x1e, 0x3b, 0xe9, 0xfa, 0x41, 0xba, 0xb6, 0x67, 0xf3, 0x64, 0x4a, 0xef,
	0xc3, 0x5a, 0x7b, 0x38, 0x24, 0xd7, 0x8c, 0x4a, 0xa6, 0x6c, 0xb3, 0xa0, 0x4c, 0x48, 0xcb, 0xbb,
	0x5c, 0x93, 0xb4, 0xd5, 0x11, 0x00, 0x1b, 0x72, 0x1d, 0x4e, 0x8d, 0xe3, 0xa9, 0xf6, 0x4f, 0xaa,
	0xbd, 0x3e, 0x96, 0x98, 0x15, 0x4b, 0xa0, 0xd6, 0x43, 0x17, 0x44, 0xaa, 0x13, 0x99, 0xa1, 0x21,
	0xa2, 0xb1, 0x61, 0x21, 0x81, 0xf1, 0x17, 0x4b, 0x7b, 0x73, 0xc1, 0x91, 0x3c, 0xeb, 0x8c, 0x3c,
	0xd5, 0x38, 0x64, 0x68, 0x24, 0x86, 0x85, 0x70, 0xa5, 0xb4, 0x0a, 0xe9, 0xfc, 0x45, 0x3a, 0x6b,
	0x73, 0xf9, 0x32, 0xe5, 0xcf, 0xe1, 0x9c, 0xe1, 0x03, 0x34, 0x10, 0x6a, 0x42, 0x2e, 0x6a, 0xbc,
	0xeb, 0xb7, 0xf1, 0xbd, 0x52, 0xbd, 0x67, 0x46, 0xd0, 0x06, 0x0d, 0x02, 0xaf, 0x0d, 0x7b, 0x00,
	0x5b, 0x33, 0x28, 0x49, 0xdc, 0x6f, 0x12, 0xe7, 0x4f, 0x03, 0x3b, 0x4d, 0x3b, 0x5f, 0x17, 0x61,
	0x25, 0xe4, 0x66, 0xa0, 0xa4, 0xe1, 0xec, 0x06, 0x78, 0xd9, 0xc6, 0x14, 0xdd, 0x9c, 0xaf, 0x23,
	0x6d, 0xd3, 0x63, 0xf7, 0x1b, 0x52, 0x22, 0x7b, 0x05, 0x27, 0xdd, 0xae, 0x8c, 0x77, 0x06, 0x7d,
	0x59, 0x41, 0x70, 0x50, 0x00, 0x17, 0x57, 0x6a, 0x0f, 0xcf, 0x63, 0x3d, 0x0a, 0xd7, 0xfb, 0x93,
	0x17, 0x6c, 0x17, 0x96, 0xf3, 0x1d, 0x45, 0x3f, 0x39, 0xc6, 0xed, 0x7f, 0x18, 0x69, 0x83, 0xf7,
	0xe8, 0x3f, 0x1c, 0xa6, 0xa3, 0xb5, 0x2b, 0x5a, 0x1d, 0xa0, 0x07, 0x1c, 0x6a, 0x37, 0x38, 0xfa,
	0x9b, 0x12, 0x0c, 0x3b, 0x11, 0x84, 0xea, 0x20, 0x74, 0x24, 0xd5, 0x3f, 0x0b, 0x50, 0xc1, 0x03,
	0xdb, 0x80, 0x25, 0x3c, 0x3a, 0x73, 0x7f, 0x6c, 0x62, 0x73, 0xbc, 0xd0, 0xc3, 0x23, 0x5a, 0xf7,
	0x26, 0x78, 0xd9, 0xb0, 0xfc, 0x4f, 0xcd, 0xd2, 0x11, 0x52, 0x26, 0xdb, 0x07, 0x7f, 0x96, 0xa5,
	0xfc, 0xcf, 0xcd, 0xf2, 0xd5, 0xd9, 0x30, 0x53, 0x3d, 0xc6, 0x5a, 0x70, 0x76, 0x2a, 0xaf, 0xff,
	0xa5, 0x79, 0x04, 0x77, 0x4d, 0xe1, 0x7c, 0xf8, 0x12, 0x36, 0x85, 0x2a, 0xf4, 0xef, 0xf0, 0x41,
	0x7f, 0x7d, 0xe7, 0xbf, 0x9e, 0xfa, 0xb7, 0x4b, 0xd9, 0x6b, 0x7a, 0xfb, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x42, 0x76, 0xf7, 0x2a, 0x06, 0x00, 0x00,
}
