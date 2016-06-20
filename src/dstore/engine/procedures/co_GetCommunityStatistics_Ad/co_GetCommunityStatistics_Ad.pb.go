// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetCommunityStatistics_Ad.proto
// DO NOT EDIT!

/*
Package co_GetCommunityStatistics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetCommunityStatistics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetCommunityStatistics_Ad

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
	CommunityId                    *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                bool                          `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	StatisticInformationIdList     *dstore_values.StringValue    `protobuf:"bytes,2,opt,name=statistic_information_id_list,json=statisticInformationIdList" json:"statistic_information_id_list,omitempty"`
	StatisticInformationIdListNull bool                          `protobuf:"varint,1002,opt,name=statistic_information_id_list_null,json=statisticInformationIdListNull" json:"statistic_information_id_list_null,omitempty"`
	FromDate                       *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                   bool                          `protobuf:"varint,1003,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                         *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                     bool                          `protobuf:"varint,1004,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	CalculateSum                   *dstore_values.BooleanValue   `protobuf:"bytes,5,opt,name=calculate_sum,json=calculateSum" json:"calculate_sum,omitempty"`
	CalculateSumNull               bool                          `protobuf:"varint,1005,opt,name=calculate_sum_null,json=calculateSumNull" json:"calculate_sum_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetStatisticInformationIdList() *dstore_values.StringValue {
	if m != nil {
		return m.StatisticInformationIdList
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

func (m *Parameters) GetCalculateSum() *dstore_values.BooleanValue {
	if m != nil {
		return m.CalculateSum
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
	RowId                  int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description            *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	CommunityId            *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	IsPublic               *dstore_values.BooleanValue   `protobuf:"bytes,10003,opt,name=is_public,json=isPublic" json:"is_public,omitempty"`
	Value                  *dstore_values.DecimalValue   `protobuf:"bytes,10004,opt,name=value" json:"value,omitempty"`
	CommunityName          *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	StatisticInformationId *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=statistic_information_id,json=statisticInformationId" json:"statistic_information_id,omitempty"`
	StatisticInformation   *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=statistic_information,json=statisticInformation" json:"statistic_information,omitempty"`
	Day                    *dstore_values.TimestampValue `protobuf:"bytes,10008,opt,name=day" json:"day,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Response_Row) GetIsPublic() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsPublic
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
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

func (m *Response_Row) GetDay() *dstore_values.TimestampValue {
	if m != nil {
		return m.Day
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetCommunityStatistics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetCommunityStatistics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetCommunityStatistics_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xeb, 0x6e, 0xd3, 0x3e,
	0x14, 0xd7, 0xfe, 0x5d, 0xbb, 0xee, 0xac, 0xbb, 0xfc, 0x0d, 0x4c, 0xa6, 0xd3, 0x26, 0x18, 0x42,
	0x42, 0x20, 0x32, 0x34, 0x04, 0x1a, 0x20, 0x6e, 0x03, 0x84, 0x06, 0x2c, 0x9a, 0x32, 0x34, 0x09,
	0x24, 0x14, 0x79, 0x89, 0x57, 0x59, 0x4a, 0xe2, 0xca, 0x76, 0x98, 0xf6, 0x16, 0xdc, 0xe1, 0x21,
	0x78, 0x14, 0x3e, 0xf2, 0x12, 0xdc, 0xde, 0x01, 0x3b, 0x4e, 0xdb, 0x34, 0x74, 0x6b, 0xe1, 0x4b,
	0x2b, 0xdb, 0xbf, 0xcb, 0x71, 0xce, 0xf9, 0x25, 0x70, 0x23, 0x94, 0x8a, 0x0b, 0xba, 0x42, 0x93,
	0x16, 0x4b, 0xe8, 0x4a, 0x5b, 0xf0, 0x80, 0x86, 0xa9, 0xa0, 0x72, 0x25, 0xe0, 0xfe, 0x43, 0xaa,
	0xee, 0xf1, 0x38, 0x4e, 0x13, 0xa6, 0x0e, 0xb6, 0x15, 0x51, 0x4c, 0x2a, 0x16, 0x48, 0xff, 0x6e,
	0xe8, 0x68, 0x9c, 0xe2, 0xe8, 0xbc, 0x25, 0x3b, 0x96, 0xec, 0x1c, 0xc5, 0x68, 0x1e, 0xcb, 0x8d,
	0x5e, 0x92, 0x28, 0xa5, 0xd2, 0x0a, 0x34, 0x4f, 0xf6, 0xbb, 0x53, 0x21, 0xb8, 0xc8, 0x8f, 0x16,
	0xfa, 0x8f, 0x62, 0x2a, 0x25, 0x69, 0xd1, 0xfc, 0xf0, 0x4c, 0xf9, 0x50, 0x11, 0x96, 0xec, 0x71,
	0x11, 0x6b, 0x47, 0x9e, 0x58, 0xd0, 0xf2, 0xd7, 0x71, 0x80, 0x2d, 0x22, 0x88, 0x3e, 0xa5, 0x42,
	0xa2, 0x5b, 0xd0, 0x08, 0x3a, 0xa5, 0xf9, 0x2c, 0xc4, 0x63, 0xa7, 0xc6, 0xce, 0x4d, 0xad, 0x2e,
	0x38, 0xf9, 0x1d, 0xf2, 0xba, 0x58, 0xa2, 0x68, 0x8b, 0x8a, 0x1d, 0xb3, 0xf2, 0xa6, 0xba, 0x84,
	0x8d, 0x10, 0x5d, 0x80, 0xff, 0x8b, 0x7c, 0x3f, 0x49, 0xa3, 0x08, 0x7f, 0x9b, 0xd0, 0x2a, 0x75,
	0x6f, 0xb6, 0x00, 0x74, 0xf5, 0x3e, 0x7a, 0x01, 0x8b, 0xb2, 0x73, 0x7d, 0xbf, 0x50, 0x9a, 0x21,
	0x46, 0x7a, 0x17, 0xff, 0x97, 0xb9, 0x37, 0x4b, 0xee, 0x52, 0x09, 0x96, 0xb4, 0xac, 0x79, 0xb3,
	0x2b, 0xb0, 0xd1, 0xe3, 0x6f, 0x84, 0x4f, 0xf4, 0x16, 0x7a, 0x0c, 0xcb, 0x47, 0xca, 0xdb, 0xe2,
	0xbe, 0xdb, 0xe2, 0x96, 0x0e, 0x17, 0xca, 0x6a, 0xbd, 0x0e, 0x93, 0x7b, 0x82, 0xc7, 0x7e, 0x48,
	0x14, 0xc5, 0x95, 0xac, 0xae, 0xc5, 0x52, 0x5d, 0x8a, 0xe9, 0xe7, 0xaf, 0x48, 0xdc, 0xb6, 0xa5,
	0xd5, 0x0d, 0xfe, 0xbe, 0x86, 0xa3, 0xb3, 0x30, 0xd3, 0xe5, 0x5a, 0xd3, 0x1f, 0xd6, 0xb4, 0xd1,
	0x81, 0x64, 0x16, 0x57, 0x61, 0x42, 0x71, 0x6b, 0x30, 0x3e, 0x8a, 0x41, 0x4d, 0xf1, 0x4c, 0xfe,
	0x34, 0x34, 0x72, 0x9e, 0x15, 0xff, 0x69, 0xc5, 0xc1, 0x1e, 0x67, 0xd2, 0x77, 0x60, 0x3a, 0x20,
	0x51, 0x90, 0x46, 0x06, 0x24, 0xd3, 0x18, 0x57, 0x07, 0xf6, 0x75, 0x97, 0xf3, 0x88, 0x92, 0xc4,
	0xca, 0x37, 0xba, 0x8c, 0xed, 0x34, 0x46, 0x17, 0x01, 0xf5, 0x29, 0x58, 0xab, 0x5f, 0xd6, 0x6a,
	0xae, 0x08, 0x35, 0x86, 0xcb, 0x5f, 0x6a, 0x50, 0xf7, 0xa8, 0x6c, 0xf3, 0x44, 0x52, 0x74, 0x09,
	0xaa, 0xd9, 0xd0, 0xe6, 0xd3, 0xd4, 0xed, 0x67, 0x9e, 0x08, 0x3b, 0xd0, 0x0f, 0xcc, 0xaf, 0x67,
	0x81, 0xe8, 0x19, 0xcc, 0x99, 0x71, 0x2d, 0x76, 0x4d, 0x0f, 0x43, 0x45, 0x93, 0x9d, 0x12, 0xb9,
	0x3c, 0xd5, 0x9b, 0x7a, 0x5d, 0xe8, 0xa0, 0x37, 0x1b, 0xf7, 0x6f, 0xa0, 0x35, 0x98, 0xc8, 0x63,
	0xa2, 0xdb, 0x68, 0x14, 0x97, 0xfe, 0x50, 0xb4, 0x21, 0xda, 0xb4, 0xff, 0x5e, 0x07, 0x8e, 0x1e,
	0x41, 0x45, 0xf0, 0x7d, 0xdd, 0x1b, 0xc3, 0x5a, 0x73, 0x46, 0x8f, 0xb5, 0xd3, 0x79, 0x12, 0x8e,
	0xc7, 0xf7, 0x3d, 0x23, 0xd2, 0xfc, 0x3c, 0x0e, 0x15, 0xbd, 0x40, 0xf3, 0x50, 0xd3, 0x4b, 0x93,
	0xb4, 0x57, 0xae, 0x7e, 0x38, 0x55, 0xaf, 0xaa, 0x97, 0x3a, 0x47, 0x37, 0x61, 0x2a, 0xa4, 0x32,
	0x10, 0xac, 0x9d, 0xdd, 0xfd, 0xb5, 0x3b, 0x34, 0x09, 0x45, 0x3c, 0xba, 0x5d, 0x8a, 0xf1, 0x1b,
	0xf7, 0x2f, 0x73, 0x7c, 0x0d, 0x26, 0x99, 0xf4, 0xdb, 0xe9, 0x6e, 0xc4, 0x02, 0xfc, 0xd6, 0x1d,
	0x3e, 0x2d, 0x75, 0x26, 0xb7, 0x32, 0x34, 0x5a, 0x85, 0x6a, 0x06, 0xc0, 0xef, 0x06, 0xd3, 0x42,
	0x1a, 0xb0, 0x98, 0x44, 0x96, 0x66, 0xa1, 0x68, 0x1d, 0x66, 0x7a, 0xf5, 0x26, 0xfa, 0x65, 0x84,
	0xdf, 0x0f, 0xbf, 0xf1, 0x74, 0x97, 0xe2, 0x6a, 0x06, 0xda, 0x01, 0x7c, 0x58, 0xdc, 0xf1, 0x87,
	0x11, 0xee, 0x3f, 0x3f, 0xf8, 0x0d, 0x80, 0xb6, 0xe0, 0xc4, 0x40, 0x5d, 0xfc, 0x71, 0x78, 0x89,
	0xc7, 0x07, 0x69, 0xea, 0x3c, 0x54, 0x42, 0x72, 0x80, 0x3f, 0xb9, 0xa3, 0xa4, 0xdc, 0x40, 0xd7,
	0x9f, 0xc2, 0x02, 0xe3, 0xa5, 0x89, 0xeb, 0x7d, 0x85, 0x9e, 0x5f, 0xf9, 0xa7, 0xef, 0xd3, 0x6e,
	0x2d, 0xfb, 0x04, 0x5c, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x6a, 0x43, 0xe5, 0xdf, 0x06,
	0x00, 0x00,
}