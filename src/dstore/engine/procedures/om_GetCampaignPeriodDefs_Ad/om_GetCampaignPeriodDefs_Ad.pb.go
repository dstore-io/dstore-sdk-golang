// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignPeriodDefs_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaignPeriodDefs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignPeriodDefs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignPeriodDefs_Ad

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
	CampaignId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
	CampaignIdNull bool                        `protobuf:"varint,1001,opt,name=campaign_id_null,json=campaignIdNull" json:"campaign_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCampaignId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignId
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
	RowId                 int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	RepeatEveryXTimeUnits *dstore_values.IntegerValue   `protobuf:"bytes,10001,opt,name=repeat_every_x_time_units,json=repeatEveryXTimeUnits" json:"repeat_every_x_time_units,omitempty"`
	RepeatUntil           *dstore_values.TimestampValue `protobuf:"bytes,10002,opt,name=repeat_until,json=repeatUntil" json:"repeat_until,omitempty"`
	StartDateAndTime      *dstore_values.TimestampValue `protobuf:"bytes,10003,opt,name=start_date_and_time,json=startDateAndTime" json:"start_date_and_time,omitempty"`
	EndDateAndTime        *dstore_values.TimestampValue `protobuf:"bytes,10004,opt,name=end_date_and_time,json=endDateAndTime" json:"end_date_and_time,omitempty"`
	StartDateAndTimeChar  *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=start_date_and_time_char,json=startDateAndTimeChar" json:"start_date_and_time_char,omitempty"`
	TimeUnitId            *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=time_unit_id,json=timeUnitId" json:"time_unit_id,omitempty"`
	PeriodDefinitionName  *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=period_definition_name,json=periodDefinitionName" json:"period_definition_name,omitempty"`
	EndDateAndTimeChar    *dstore_values.StringValue    `protobuf:"bytes,10008,opt,name=end_date_and_time_char,json=endDateAndTimeChar" json:"end_date_and_time_char,omitempty"`
	TimeUnit              *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=time_unit,json=timeUnit" json:"time_unit,omitempty"`
	PeriodDefinitionId    *dstore_values.IntegerValue   `protobuf:"bytes,10010,opt,name=period_definition_id,json=periodDefinitionId" json:"period_definition_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetRepeatEveryXTimeUnits() *dstore_values.IntegerValue {
	if m != nil {
		return m.RepeatEveryXTimeUnits
	}
	return nil
}

func (m *Response_Row) GetRepeatUntil() *dstore_values.TimestampValue {
	if m != nil {
		return m.RepeatUntil
	}
	return nil
}

func (m *Response_Row) GetStartDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.StartDateAndTime
	}
	return nil
}

func (m *Response_Row) GetEndDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.EndDateAndTime
	}
	return nil
}

func (m *Response_Row) GetStartDateAndTimeChar() *dstore_values.StringValue {
	if m != nil {
		return m.StartDateAndTimeChar
	}
	return nil
}

func (m *Response_Row) GetTimeUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TimeUnitId
	}
	return nil
}

func (m *Response_Row) GetPeriodDefinitionName() *dstore_values.StringValue {
	if m != nil {
		return m.PeriodDefinitionName
	}
	return nil
}

func (m *Response_Row) GetEndDateAndTimeChar() *dstore_values.StringValue {
	if m != nil {
		return m.EndDateAndTimeChar
	}
	return nil
}

func (m *Response_Row) GetTimeUnit() *dstore_values.StringValue {
	if m != nil {
		return m.TimeUnit
	}
	return nil
}

func (m *Response_Row) GetPeriodDefinitionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PeriodDefinitionId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignPeriodDefs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignPeriodDefs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignPeriodDefs_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x4e, 0x54, 0x31,
	0x10, 0x0e, 0xae, 0xc0, 0x3a, 0x10, 0xc4, 0x82, 0xa4, 0x2c, 0xd1, 0x10, 0xbc, 0x91, 0x9b, 0x83,
	0x51, 0x13, 0x89, 0x31, 0x26, 0x08, 0xc4, 0xec, 0x05, 0x2b, 0x1e, 0x84, 0xa8, 0x37, 0x4d, 0xe5,
	0x0c, 0x6b, 0x93, 0xdd, 0x76, 0xd3, 0x76, 0x41, 0xdf, 0xc2, 0x3f, 0xfc, 0x7b, 0x0e, 0x5f, 0xca,
	0xb7, 0x70, 0xce, 0x69, 0x17, 0x38, 0x07, 0x23, 0x1b, 0x6f, 0x76, 0x33, 0x9d, 0xf9, 0xbe, 0xce,
	0xf7, 0xcd, 0x9c, 0xc2, 0xc3, 0xcc, 0x79, 0x63, 0x71, 0x05, 0x75, 0x5b, 0x69, 0x5c, 0xe9, 0x59,
	0xb3, 0x8f, 0x59, 0xdf, 0xa2, 0x5b, 0x31, 0x5d, 0xf1, 0x14, 0xfd, 0xba, 0xec, 0xf6, 0xa4, 0x6a,
	0xeb, 0x6d, 0xb4, 0xca, 0x64, 0x1b, 0x78, 0xe0, 0xc4, 0x5a, 0x96, 0x50, 0x99, 0x37, 0x6c, 0x39,
	0x60, 0x93, 0x80, 0x4d, 0xfe, 0x01, 0x68, 0xcc, 0xc4, 0x6b, 0x0e, 0x65, 0xa7, 0x8f, 0x2e, 0xe0,
	0x1b, 0xf3, 0xe5, 0xbb, 0xd1, 0x5a, 0x63, 0x63, 0x6a, 0xa1, 0x9c, 0xea, 0xa2, 0x73, 0xb2, 0x8d,
	0x31, 0x79, 0xab, 0x9a, 0xf4, 0x52, 0xe9, 0x03, 0x63, 0xbb, 0xd2, 0x2b, 0xa3, 0x43, 0xd1, 0x52,
	0x1f, 0x60, 0x5b, 0x5a, 0x49, 0x49, 0xb4, 0x8e, 0x3d, 0x82, 0x89, 0xfd, 0xd8, 0x98, 0x50, 0x19,
	0x1f, 0x59, 0x1c, 0xb9, 0x3d, 0x71, 0x77, 0x21, 0x89, 0x02, 0x62, 0x57, 0x4a, 0x7b, 0x6c, 0xa3,
	0xdd, 0xcb, 0xa3, 0x14, 0x06, 0xf5, 0xcd, 0x8c, 0x2d, 0xc3, 0xf4, 0x19, 0xb4, 0xd0, 0xfd, 0x4e,
	0x87, 0xff, 0x1e, 0x27, 0x8e, 0x7a, 0x3a, 0x75, 0x5a, 0xd6, 0xa2, 0xe3, 0xa5, 0x5f, 0x75, 0xa8,
	0xa7, 0xe8, 0x7a, 0x46, 0x3b, 0x64, 0x77, 0x60, 0xb4, 0x10, 0x15, 0xef, 0x6b, 0x24, 0x65, 0xc3,
	0x82, 0xe0, 0xcd, 0xfc, 0x37, 0x0d, 0x85, 0xec, 0x15, 0x4c, 0xe7, 0x72, 0xc4, 0x19, 0x3d, 0xfc,
	0xd2, 0x62, 0x8d, 0xc0, 0x49, 0x05, 0x5c, 0x55, 0xbd, 0x45, 0x71, 0xf3, 0x34, 0x4e, 0xaf, 0x76,
	0xcb, 0x07, 0x6c, 0x15, 0xc6, 0xa3, 0x8d, 0xbc, 0x56, 0x30, 0xde, 0x3c, 0xc7, 0x18, 0x4c, 0xde,
	0x0a, 0xff, 0xe9, 0xa0, 0x9c, 0x35, 0xa1, 0x66, 0xcd, 0x11, 0xbf, 0x5c, 0xa0, 0x1e, 0x24, 0x43,
	0x4f, 0x3d, 0x19, 0x18, 0x91, 0xa4, 0xe6, 0x28, 0xcd, 0x39, 0x1a, 0xc7, 0x63, 0x50, 0xa3, 0x80,
	0xcd, 0xc1, 0x18, 0x85, 0xf9, 0x28, 0x3e, 0xb4, 0xc8, 0x9b, 0xd1, 0x74, 0x94, 0x42, 0x72, 0x7a,
	0x0f, 0xe6, 0x2d, 0xf6, 0x50, 0x7a, 0x81, 0x87, 0x68, 0xdf, 0x8b, 0x77, 0xc2, 0xab, 0x2e, 0x8a,
	0xbe, 0x56, 0xde, 0xf1, 0x8f, 0xad, 0x8b, 0xc7, 0x76, 0x3d, 0xc0, 0x37, 0x73, 0xf4, 0xcb, 0x17,
	0x84, 0xdd, 0xcd, 0xa1, 0x6c, 0x0d, 0x26, 0x23, 0x6f, 0x5f, 0x7b, 0xd5, 0xe1, 0x9f, 0x02, 0xd5,
	0x8d, 0x0a, 0x55, 0x7e, 0x99, 0xf3, 0xa4, 0x25, 0x90, 0x4d, 0x04, 0xcc, 0x6e, 0x0e, 0x61, 0x5b,
	0x30, 0x43, 0x29, 0xeb, 0x45, 0x26, 0x3d, 0x0a, 0xa9, 0xb3, 0xa2, 0x35, 0xfe, 0x79, 0x28, 0xa6,
	0xe9, 0x02, 0xba, 0x41, 0xc8, 0x35, 0x9d, 0xe5, 0x6d, 0x91, 0xa9, 0xd7, 0x90, 0x38, 0xca, 0x64,
	0x5f, 0x86, 0x22, 0x9b, 0x22, 0xe0, 0x59, 0xaa, 0x1d, 0xe0, 0x7f, 0xe9, 0x4c, 0xec, 0xbf, 0x95,
	0x96, 0x1f, 0xb7, 0xca, 0xab, 0x17, 0x19, 0x9d, 0xb7, 0x4a, 0xb7, 0x03, 0xdd, 0x6c, 0xb5, 0xb7,
	0x75, 0x02, 0xb2, 0xc7, 0x30, 0x79, 0x62, 0x7d, 0x3e, 0xa7, 0xaf, 0x43, 0x98, 0x0f, 0x3e, 0x1a,
	0x4e, 0x93, 0x7c, 0x0e, 0x73, 0xbd, 0x62, 0x19, 0x44, 0x86, 0x07, 0x8a, 0xce, 0x68, 0x07, 0x85,
	0xa6, 0xef, 0x91, 0x7f, 0x1b, 0xa2, 0xa5, 0xde, 0x60, 0x8f, 0x22, 0xb2, 0x45, 0x40, 0xf6, 0x0c,
	0xe6, 0xce, 0x59, 0x16, 0x54, 0x7e, 0xbf, 0x98, 0x92, 0x95, 0x4d, 0x2b, 0x34, 0xae, 0xc2, 0x95,
	0x13, 0x8d, 0xfc, 0xc7, 0xc5, 0x1c, 0xf5, 0x81, 0x3e, 0xd6, 0x82, 0xd9, 0xf3, 0xea, 0xc8, 0xa5,
	0x9f, 0x43, 0xb8, 0xc4, 0xaa, 0xe2, 0x9a, 0xd9, 0x93, 0x1d, 0x58, 0x50, 0xa6, 0xf2, 0x65, 0x9d,
	0xbe, 0xc5, 0xaf, 0xef, 0xff, 0xcf, 0x2b, 0xfd, 0x66, 0xac, 0x78, 0x09, 0xef, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x02, 0x0e, 0x57, 0xe4, 0x05, 0x00, 0x00,
}
