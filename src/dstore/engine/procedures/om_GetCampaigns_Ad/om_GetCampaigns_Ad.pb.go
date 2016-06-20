// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaigns_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaigns_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaigns_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaigns_Ad

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
	CampaignId                     *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
	CampaignIdNull                 bool                          `protobuf:"varint,1001,opt,name=campaign_id_null,json=campaignIdNull" json:"campaign_id_null,omitempty"`
	CampaignTypeId                 *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignTypeIdNull             bool                          `protobuf:"varint,1002,opt,name=campaign_type_id_null,json=campaignTypeIdNull" json:"campaign_type_id_null,omitempty"`
	Active                         *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=active" json:"active,omitempty"`
	ActiveNull                     bool                          `protobuf:"varint,1003,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
	FilterByPeriodStartMinDate     *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=filter_by_period_start_min_date,json=filterByPeriodStartMinDate" json:"filter_by_period_start_min_date,omitempty"`
	FilterByPeriodStartMinDateNull bool                          `protobuf:"varint,1004,opt,name=filter_by_period_start_min_date_null,json=filterByPeriodStartMinDateNull" json:"filter_by_period_start_min_date_null,omitempty"`
	FilterByPeriodStartMaxDate     *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=filter_by_period_start_max_date,json=filterByPeriodStartMaxDate" json:"filter_by_period_start_max_date,omitempty"`
	FilterByPeriodStartMaxDateNull bool                          `protobuf:"varint,1005,opt,name=filter_by_period_start_max_date_null,json=filterByPeriodStartMaxDateNull" json:"filter_by_period_start_max_date_null,omitempty"`
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

func (m *Parameters) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Parameters) GetFilterByPeriodStartMinDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FilterByPeriodStartMinDate
	}
	return nil
}

func (m *Parameters) GetFilterByPeriodStartMaxDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FilterByPeriodStartMaxDate
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
	RowId               int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Active              *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=active" json:"active,omitempty"`
	CampaignId          *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
	CampaignDescription *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=campaign_description,json=campaignDescription" json:"campaign_description,omitempty"`
	CampaignTypeId      *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignName        *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=campaign_name,json=campaignName" json:"campaign_name,omitempty"`
	CampaignType        *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=campaign_type,json=campaignType" json:"campaign_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Response_Row) GetCampaignId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignId
	}
	return nil
}

func (m *Response_Row) GetCampaignDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignDescription
	}
	return nil
}

func (m *Response_Row) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
	}
	return nil
}

func (m *Response_Row) GetCampaignName() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignName
	}
	return nil
}

func (m *Response_Row) GetCampaignType() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignType
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaigns_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaigns_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaigns_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x56, 0xff, 0xfc, 0x6e, 0xab, 0x2d, 0x87, 0x6a, 0x0b, 0xc8, 0x24, 0xa2, 0x44, 0x85, 0x0b,
	0x10, 0x92, 0x03, 0x2d, 0x17, 0x5c, 0x80, 0x54, 0x4a, 0x0b, 0x2a, 0x52, 0xac, 0xca, 0x20, 0x24,
	0xb8, 0xb1, 0x36, 0xf1, 0x36, 0x5a, 0x29, 0xde, 0xb5, 0x76, 0x37, 0x6d, 0xf3, 0x16, 0x9c, 0xdf,
	0x02, 0x9e, 0x82, 0x97, 0xe1, 0xf4, 0x0e, 0x8c, 0x3d, 0xce, 0xc1, 0x6e, 0xc0, 0x84, 0x9b, 0x24,
	0xe3, 0x99, 0xef, 0x90, 0xdd, 0x99, 0x31, 0xb9, 0x13, 0x19, 0xab, 0x34, 0x6f, 0x71, 0xd9, 0x13,
	0x92, 0xb7, 0x12, 0xad, 0xba, 0x3c, 0x1a, 0x68, 0x6e, 0x5a, 0x2a, 0x0e, 0x9f, 0x70, 0xfb, 0x88,
	0xc5, 0x09, 0x13, 0x3d, 0x69, 0xc2, 0x87, 0x91, 0x07, 0x59, 0xab, 0x68, 0x13, 0x21, 0x1e, 0x42,
	0xbc, 0xd3, 0x75, 0xf5, 0xb5, 0x9c, 0xf4, 0x88, 0xf5, 0x07, 0xdc, 0x20, 0xac, 0x7e, 0xb9, 0xa8,
	0xc4, 0xb5, 0x56, 0x3a, 0x4f, 0x35, 0x8a, 0xa9, 0x98, 0x1b, 0xc3, 0x7a, 0x3c, 0x4f, 0x5e, 0x2b,
	0x27, 0x2d, 0x13, 0xf2, 0x50, 0xe9, 0x98, 0x59, 0xa1, 0x24, 0x16, 0x6d, 0x7c, 0x72, 0x08, 0x39,
	0x60, 0x9a, 0x41, 0x96, 0x6b, 0x43, 0xef, 0x93, 0x95, 0x6e, 0x6e, 0x28, 0x14, 0x91, 0xbb, 0xd0,
	0x5c, 0xb8, 0xb1, 0xb2, 0xd9, 0xf0, 0x72, 0xe3, 0xb9, 0x2d, 0x21, 0x2d, 0xef, 0x71, 0xfd, 0x22,
	0x8d, 0x02, 0x32, 0xaa, 0xdf, 0x8f, 0xe8, 0x4d, 0xb2, 0x3a, 0x85, 0x0e, 0xe5, 0xa0, 0xdf, 0x77,
	0xbf, 0x2e, 0x01, 0xc7, 0x72, 0x70, 0x6e, 0x52, 0xe6, 0xc3, 0x63, 0xba, 0x37, 0x55, 0x6a, 0x87,
	0x09, 0x4f, 0xd5, 0xfe, 0xab, 0x56, 0x1b, 0xd3, 0x3c, 0x07, 0x0c, 0x28, 0x6e, 0x92, 0x8b, 0x65,
	0x1a, 0x94, 0xfd, 0x86, 0xb2, 0xb4, 0x58, 0x9f, 0x49, 0x6f, 0x91, 0x45, 0xd6, 0xb5, 0xe2, 0x88,
	0xbb, 0xb5, 0x6a, 0xc1, 0xbc, 0x94, 0x36, 0xc9, 0x0a, 0xfe, 0x42, 0xfa, 0xef, 0x48, 0x4f, 0xf0,
	0x59, 0x46, 0xdb, 0x21, 0x57, 0x0f, 0x45, 0x1f, 0x4e, 0x31, 0xec, 0x0c, 0xc3, 0x84, 0x6b, 0xa1,
	0xa2, 0xd0, 0x58, 0xa6, 0x6d, 0x18, 0x0b, 0x19, 0x46, 0xcc, 0x72, 0xf7, 0xff, 0x4c, 0xef, 0x4a,
	0x49, 0xcf, 0x0a, 0xb8, 0x37, 0x0b, 0x3e, 0x51, 0xb1, 0x8e, 0x2c, 0x3b, 0xc3, 0x83, 0x8c, 0xe3,
	0x59, 0x4a, 0xd1, 0x16, 0x72, 0x17, 0x08, 0x68, 0x9b, 0x5c, 0xaf, 0xd0, 0x40, 0x7b, 0x3f, 0xd0,
	0xde, 0xfa, 0xef, 0xa9, 0xaa, 0x2c, 0xb3, 0x13, 0xb4, 0xec, 0xfc, 0xb3, 0x65, 0x76, 0x52, 0x65,
	0x39, 0xd7, 0x40, 0xcb, 0x3f, 0xff, 0x60, 0x19, 0xa9, 0x52, 0xcb, 0x1b, 0x5f, 0x1c, 0xb2, 0x1c,
	0x70, 0x93, 0x28, 0x69, 0x38, 0xbd, 0x4d, 0x9c, 0x6c, 0x1a, 0xf2, 0x3e, 0xad, 0x7b, 0xc5, 0x01,
	0xc3, 0x49, 0xd9, 0x4b, 0x3f, 0x03, 0x2c, 0xa4, 0x2f, 0xc9, 0x6a, 0x3a, 0x07, 0xe1, 0xd4, 0x20,
	0x40, 0xdb, 0xd5, 0x00, 0xec, 0x95, 0xc0, 0xe5, 0x71, 0x69, 0x43, 0xbc, 0x3f, 0x89, 0x83, 0xf3,
	0x71, 0xf1, 0x01, 0xbd, 0x47, 0x96, 0xf2, 0xf9, 0x83, 0xbe, 0x4a, 0x19, 0xd7, 0x4f, 0x31, 0xe2,
	0x74, 0xb6, 0xf1, 0x3b, 0x18, 0x95, 0xd3, 0x6d, 0x52, 0xd3, 0xea, 0x18, 0xba, 0x63, 0x96, 0x8f,
	0x19, 0xdb, 0x64, 0xf4, 0xff, 0xbd, 0x40, 0x1d, 0x07, 0x29, 0xb4, 0xfe, 0xb9, 0x46, 0x6a, 0x10,
	0xd0, 0x4b, 0x64, 0x11, 0xc2, 0x74, 0x96, 0x5e, 0xfb, 0x70, 0x24, 0x4e, 0xe0, 0x40, 0x08, 0x63,
	0x72, 0x77, 0xdc, 0xf2, 0x6f, 0xfc, 0xbf, 0xef, 0xf9, 0x07, 0xc5, 0x65, 0xf0, 0xd6, 0x9f, 0x6f,
	0x1b, 0xf8, 0xe4, 0xc2, 0x18, 0x1e, 0x71, 0xd3, 0xd5, 0x22, 0xc9, 0xce, 0xfb, 0x9d, 0x5f, 0xbc,
	0xad, 0x9c, 0xc7, 0x58, 0x2d, 0x64, 0x0f, 0x69, 0xd6, 0x46, 0xc0, 0xdd, 0x09, 0x8e, 0x3e, 0x9e,
	0xb1, 0x32, 0xde, 0xfb, 0xf3, 0xef, 0x8c, 0x6d, 0x72, 0x76, 0xcc, 0x23, 0x61, 0xf1, 0xb9, 0x1f,
	0xaa, 0x0d, 0x9d, 0x19, 0x21, 0x7c, 0x00, 0x14, 0x18, 0x52, 0x27, 0xee, 0xc7, 0x39, 0x18, 0x52,
	0x17, 0x3b, 0x4f, 0x49, 0x43, 0xa8, 0xd2, 0x4d, 0x4f, 0x5e, 0x21, 0xaf, 0x6e, 0xcd, 0xf1, 0x72,
	0xe9, 0x2c, 0x66, 0x9b, 0x7c, 0xeb, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xe4, 0x42, 0x94,
	0x92, 0x06, 0x00, 0x00,
}