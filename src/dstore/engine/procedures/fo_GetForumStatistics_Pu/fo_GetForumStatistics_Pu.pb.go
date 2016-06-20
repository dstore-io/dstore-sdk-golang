// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetForumStatistics_Pu.proto
// DO NOT EDIT!

/*
Package fo_GetForumStatistics_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetForumStatistics_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetForumStatistics_Pu

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
	PersonIdentificationValues        *dstore_values.StringValue    `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull    bool                          `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                      *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull                  bool                          `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                          *dstore_values.StringValue    `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                      bool                          `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	ForumId                           *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	ForumIdNull                       bool                          `protobuf:"varint,1004,opt,name=forum_id_null,json=forumIdNull" json:"forum_id_null,omitempty"`
	ListOfStatisticInformationIds     *dstore_values.StringValue    `protobuf:"bytes,5,opt,name=list_of_statistic_information_ids,json=listOfStatisticInformationIds" json:"list_of_statistic_information_ids,omitempty"`
	ListOfStatisticInformationIdsNull bool                          `protobuf:"varint,1005,opt,name=list_of_statistic_information_ids_null,json=listOfStatisticInformationIdsNull" json:"list_of_statistic_information_ids_null,omitempty"`
	FromDate                          *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                      bool                          `protobuf:"varint,1006,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                            *dstore_values.TimestampValue `protobuf:"bytes,7,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                        bool                          `protobuf:"varint,1007,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	DateFormat                        *dstore_values.StringValue    `protobuf:"bytes,8,opt,name=date_format,json=dateFormat" json:"date_format,omitempty"`
	DateFormatNull                    bool                          `protobuf:"varint,1008,opt,name=date_format_null,json=dateFormatNull" json:"date_format_null,omitempty"`
	Visibility                        *dstore_values.IntegerValue   `protobuf:"bytes,9,opt,name=visibility" json:"visibility,omitempty"`
	VisibilityNull                    bool                          `protobuf:"varint,1009,opt,name=visibility_null,json=visibilityNull" json:"visibility_null,omitempty"`
	SeparatorInIdentVals              *dstore_values.StringValue    `protobuf:"bytes,10,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull          bool                          `protobuf:"varint,1010,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Parameters) GetListOfStatisticInformationIds() *dstore_values.StringValue {
	if m != nil {
		return m.ListOfStatisticInformationIds
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

func (m *Parameters) GetDateFormat() *dstore_values.StringValue {
	if m != nil {
		return m.DateFormat
	}
	return nil
}

func (m *Parameters) GetVisibility() *dstore_values.IntegerValue {
	if m != nil {
		return m.Visibility
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
	ForumId                *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	Value                  *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
	StatisticInformationId *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=statistic_information_id,json=statisticInformationId" json:"statistic_information_id,omitempty"`
	StatisticInformation   *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=statistic_information,json=statisticInformation" json:"statistic_information,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetForumStatistics_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetForumStatistics_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetForumStatistics_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x4e, 0x13, 0x4f,
	0x14, 0x0f, 0x94, 0xde, 0x0e, 0xfc, 0x81, 0x0c, 0x7f, 0x71, 0x2d, 0x97, 0x58, 0x88, 0x04, 0xbf,
	0x2c, 0xa2, 0x46, 0x8c, 0x7e, 0x30, 0x1a, 0xc5, 0xd4, 0x84, 0x5a, 0x57, 0x43, 0xa2, 0x31, 0xd9,
	0x2c, 0x74, 0xda, 0x4c, 0xd2, 0xee, 0xd4, 0x9d, 0x29, 0x84, 0xb7, 0xf0, 0xf6, 0x26, 0x3e, 0x81,
	0x4f, 0xe0, 0x73, 0x78, 0xbf, 0x3c, 0x81, 0x33, 0x73, 0xa6, 0xb7, 0xb5, 0x65, 0xcb, 0x17, 0x9a,
	0xe9, 0xf9, 0x5d, 0xce, 0x9c, 0x9e, 0xfd, 0x2d, 0xb0, 0x53, 0x15, 0x92, 0x47, 0x74, 0x8b, 0x86,
	0x75, 0x16, 0xd2, 0xad, 0x56, 0xc4, 0x0f, 0x69, 0xb5, 0x1d, 0x51, 0xb1, 0x55, 0xe3, 0xfe, 0x43,
	0x2a, 0x77, 0x79, 0xd4, 0x6e, 0x3e, 0x95, 0x81, 0x64, 0x42, 0xb2, 0x43, 0xe1, 0x57, 0xda, 0xae,
	0xc2, 0x48, 0x4e, 0x36, 0x90, 0xe8, 0x22, 0xd1, 0x1d, 0x85, 0x2e, 0x2c, 0x58, 0x83, 0xa3, 0xa0,
	0xd1, 0xa6, 0x02, 0xc9, 0x85, 0x0b, 0x83, 0xae, 0x34, 0x8a, 0x78, 0x64, 0x4b, 0x4b, 0x83, 0xa5,
	0x26, 0x15, 0x22, 0xa8, 0x53, 0x5b, 0x5c, 0x8f, 0x17, 0x65, 0xc0, 0xc2, 0x1a, 0x8f, 0x9a, 0xca,
	0x91, 0x87, 0x08, 0x5a, 0xfb, 0x98, 0x07, 0xa8, 0x04, 0x51, 0xa0, 0xaa, 0x34, 0x12, 0xe4, 0x25,
	0x2c, 0xb7, 0xd4, 0x27, 0x0f, 0x7d, 0x56, 0xa5, 0xa1, 0x64, 0x35, 0x76, 0x68, 0xd0, 0x3e, 0x76,
	0xe4, 0x4c, 0x5c, 0x9c, 0xd8, 0x9c, 0xbe, 0x5a, 0x70, 0xed, 0x7d, 0x6c, 0x9f, 0x42, 0x46, 0x2c,
	0xac, 0xef, 0xeb, 0x83, 0x57, 0x40, 0x7e, 0x69, 0x80, 0x6e, 0x4a, 0x82, 0x3c, 0x82, 0xe2, 0x69,
	0xea, 0x7e, 0xd8, 0x6e, 0x34, 0x9c, 0xcf, 0x59, 0xe5, 0x91, 0xf3, 0x56, 0x47, 0xeb, 0x94, 0x15,
	0x8c, 0xdc, 0x85, 0x59, 0xab, 0x25, 0x4f, 0x5a, 0x54, 0x09, 0x3a, 0x93, 0xa6, 0xb7, 0xa5, 0x58,
	0x6f, 0x2c, 0x94, 0xb4, 0x4e, 0x23, 0x6c, 0x6e, 0x06, 0x29, 0xcf, 0x14, 0xa3, 0x54, 0x25, 0x2e,
	0x2c, 0x0c, 0x4a, 0x60, 0x03, 0x5f, 0xb0, 0x81, 0xf9, 0x7e, 0xac, 0xb1, 0xdc, 0x81, 0x7c, 0x3b,
	0x64, 0xaf, 0xda, 0xc6, 0x2d, 0x95, 0x38, 0x89, 0x1c, 0x82, 0x95, 0xd1, 0x25, 0x98, 0xed, 0x12,
	0xd1, 0xe3, 0x2b, 0x7a, 0xcc, 0x74, 0x20, 0x46, 0xff, 0x06, 0xe4, 0x6a, 0x7a, 0x27, 0xb4, 0xfc,
	0x54, 0xf2, 0x65, 0xb2, 0x06, 0xac, 0xe4, 0xd7, 0xe1, 0xbf, 0x0e, 0x0f, 0xd5, 0xbf, 0xa1, 0xfa,
	0xb4, 0x05, 0x18, 0xf1, 0x2a, 0x14, 0x1b, 0x6a, 0xd1, 0x7c, 0x5e, 0xf3, 0x45, 0x67, 0xe7, 0xfc,
	0xbe, 0x7d, 0x50, 0x64, 0xe1, 0xa4, 0x13, 0x2f, 0xb5, 0xa2, 0x45, 0x1e, 0xd7, 0xba, 0x6b, 0x5b,
	0xea, 0x29, 0x94, 0xaa, 0x82, 0x78, 0xb0, 0x91, 0xe8, 0x82, 0x3d, 0x7e, 0xc7, 0x1e, 0x8b, 0xa7,
	0xea, 0x99, 0xce, 0x6f, 0x41, 0xbe, 0x16, 0xf1, 0xa6, 0x5f, 0x0d, 0x24, 0x75, 0x32, 0xa6, 0xc3,
	0x95, 0x58, 0x87, 0x92, 0xa9, 0xd5, 0x97, 0x41, 0xb3, 0x65, 0x27, 0xaf, 0xf1, 0xf7, 0x15, 0x5c,
	0x4f, 0xbe, 0xcb, 0x45, 0xdf, 0x1f, 0x76, 0xf2, 0x1d, 0x88, 0x9d, 0x7c, 0x56, 0x72, 0x34, 0xc8,
	0x8e, 0x63, 0x90, 0x91, 0xdc, 0xc8, 0x17, 0x61, 0xc6, 0xf2, 0x50, 0xfc, 0x27, 0x8a, 0x03, 0x96,
	0x8d, 0xf4, 0x6d, 0x98, 0x36, 0x75, 0xbc, 0x96, 0x93, 0x4b, 0x9c, 0x30, 0x68, 0xf8, 0xae, 0x41,
	0x93, 0xcb, 0x30, 0xdf, 0x47, 0x46, 0x8f, 0x5f, 0xe8, 0x31, 0xdb, 0x83, 0x59, 0x1f, 0x38, 0x62,
	0x82, 0x1d, 0xb0, 0x06, 0x93, 0x27, 0x4e, 0x3e, 0x79, 0x7d, 0xfa, 0xe0, 0x64, 0x13, 0xe6, 0x7a,
	0x27, 0xb4, 0xf9, 0x6d, 0x6d, 0x7a, 0xdf, 0x1b, 0x9b, 0x27, 0x70, 0x5e, 0xd0, 0x96, 0x0a, 0x0c,
	0xa5, 0xab, 0x7e, 0x58, 0x7c, 0x90, 0xf5, 0xf3, 0x2b, 0x1c, 0x48, 0xbc, 0xda, 0xff, 0x5d, 0x6a,
	0x09, 0x1f, 0x6c, 0xf5, 0xb5, 0x20, 0x77, 0x60, 0x79, 0x84, 0x24, 0x76, 0xf2, 0x07, 0x3b, 0x71,
	0x86, 0x91, 0x75, 0x4f, 0x6b, 0x9f, 0xa6, 0x20, 0xe7, 0x51, 0xd1, 0xe2, 0xa1, 0xa0, 0xe4, 0x0a,
	0xa4, 0x4d, 0x42, 0xc6, 0xa3, 0xca, 0x46, 0x2f, 0xa6, 0xe7, 0x03, 0xfd, 0xd7, 0x43, 0x20, 0x79,
	0x0e, 0xf3, 0x3a, 0x1b, 0xfb, 0xd7, 0x54, 0x65, 0x49, 0x4a, 0x91, 0xdd, 0x18, 0x39, 0x1e, 0xa1,
	0x7b, 0xea, 0xdc, 0xb7, 0xb0, 0xde, 0x5c, 0x73, 0xf0, 0x0b, 0x72, 0x13, 0xb2, 0x36, 0x93, 0x55,
	0x5e, 0x68, 0xc5, 0xd5, 0x7f, 0x14, 0x31, 0xb1, 0xf7, 0xf0, 0xd3, 0xeb, 0xc0, 0xc9, 0x2e, 0xa4,
	0x22, 0x7e, 0xac, 0x62, 0x40, 0xb3, 0xae, 0xbb, 0xe3, 0xbd, 0x3f, 0xdc, 0xce, 0x14, 0x5c, 0x8f,
	0x1f, 0x7b, 0x5a, 0xa0, 0xf0, 0x61, 0x12, 0x52, 0xea, 0x40, 0x16, 0x21, 0xa3, 0x8e, 0x3a, 0x59,
	0x5e, 0x97, 0xd5, 0x60, 0xd2, 0x5e, 0x5a, 0x1d, 0x55, 0x76, 0xec, 0xf4, 0x65, 0xce, 0x9b, 0xf2,
	0x19, 0x42, 0x67, 0x1b, 0xd2, 0xa6, 0xee, 0xbc, 0x2d, 0x27, 0xfe, 0xee, 0x88, 0x24, 0xfb, 0xe0,
	0x8c, 0x0a, 0x05, 0xe7, 0xdd, 0x18, 0xde, 0x8b, 0x62, 0x68, 0x4a, 0x90, 0x0a, 0x9c, 0x1b, 0xaa,
	0xeb, 0xbc, 0x2f, 0x8f, 0xb1, 0x92, 0x43, 0x34, 0xef, 0x55, 0x60, 0x89, 0xf1, 0xd8, 0xd0, 0x7b,
	0x6f, 0xfb, 0x17, 0xdb, 0x67, 0xfe, 0x3f, 0xe0, 0x20, 0x63, 0x5e, 0xb7, 0xd7, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x3a, 0x55, 0x08, 0x3c, 0x43, 0x08, 0x00, 0x00,
}