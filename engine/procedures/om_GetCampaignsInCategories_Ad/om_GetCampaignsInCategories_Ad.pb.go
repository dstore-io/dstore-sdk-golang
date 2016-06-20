// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignsInCategories_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaignsInCategories_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignsInCategories_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignsInCategories_Ad

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
	CampaignCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=campaign_category_id,json=campaignCategoryId" json:"campaign_category_id,omitempty"`
	CampaignCategoryIdNull bool                        `protobuf:"varint,1001,opt,name=campaign_category_id_null,json=campaignCategoryIdNull" json:"campaign_category_id_null,omitempty"`
	CampaignId             *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
	CampaignIdNull         bool                        `protobuf:"varint,1002,opt,name=campaign_id_null,json=campaignIdNull" json:"campaign_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCampaignCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignCategoryId
	}
	return nil
}

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
	RowId               int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Active              *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=active" json:"active,omitempty"`
	CampaignId          *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
	CampaignDescription *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=campaign_description,json=campaignDescription" json:"campaign_description,omitempty"`
	CampaignTypeId      *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignName        *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=campaign_name,json=campaignName" json:"campaign_name,omitempty"`
	CategoryDescription *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
	SortNo              *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	CampaignCategoryId  *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=campaign_category_id,json=campaignCategoryId" json:"campaign_category_id,omitempty"`
	CampaignType        *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=campaign_type,json=campaignType" json:"campaign_type,omitempty"`
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

func (m *Response_Row) GetCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CategoryDescription
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetCampaignCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignCategoryId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignsInCategories_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignsInCategories_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignsInCategories_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetCampaignsInCategories_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0xa6, 0x8d, 0x49, 0xca, 0xd4, 0x4b, 0x99, 0x4a, 0xd9, 0x26, 0x20, 0x52, 0x5f, 0x14, 0x64,
	0x23, 0x5e, 0x40, 0x8a, 0x85, 0x6a, 0xbd, 0x10, 0x30, 0x8b, 0x2c, 0x22, 0xe8, 0xcb, 0x32, 0xee,
	0x8e, 0xcb, 0x60, 0x76, 0x66, 0x99, 0x99, 0xb4, 0xe4, 0xcd, 0x17, 0xdf, 0xbd, 0x5f, 0x7e, 0x81,
	0xbf, 0x4d, 0x7f, 0x85, 0x67, 0x76, 0x66, 0x93, 0xdd, 0x44, 0x4c, 0xd2, 0x97, 0x96, 0x93, 0x73,
	0xbe, 0x33, 0xdf, 0xf9, 0xbe, 0xb3, 0x07, 0x1d, 0x24, 0x4a, 0x0b, 0x49, 0x7b, 0x94, 0xa7, 0x8c,
	0xd3, 0x5e, 0x2e, 0x45, 0x4c, 0x93, 0x91, 0xa4, 0xaa, 0x27, 0xb2, 0xe8, 0x09, 0xd5, 0x47, 0x24,
	0xcb, 0x09, 0x4b, 0xb9, 0xea, 0xf3, 0x23, 0xa2, 0x69, 0x2a, 0x24, 0xa3, 0x2a, 0xba, 0x9f, 0xf8,
	0x50, 0xa9, 0x05, 0xbe, 0x6e, 0xe1, 0xbe, 0x85, 0xfb, 0xff, 0xc7, 0x74, 0xb6, 0xdd, 0x63, 0xc7,
	0x64, 0x38, 0xa2, 0xca, 0xb6, 0xe8, 0xec, 0xd6, 0x19, 0x50, 0x29, 0x85, 0x74, 0xa9, 0x6e, 0x3d,
	0x95, 0x51, 0xa5, 0x48, 0x4a, 0x5d, 0xf2, 0xca, 0x6c, 0x52, 0x13, 0xc6, 0xdf, 0x08, 0x99, 0x11,
	0xcd, 0x04, 0xb7, 0x45, 0x7b, 0xef, 0xd6, 0x11, 0x7a, 0x46, 0x24, 0x81, 0x2c, 0x95, 0x0a, 0x0f,
	0xd0, 0xc5, 0xd8, 0x91, 0x8b, 0x62, 0x4b, 0x6d, 0x1c, 0xb1, 0xc4, 0x5b, 0xbb, 0xbc, 0x76, 0x75,
	0xf3, 0x66, 0xd7, 0x77, 0xd3, 0x38, 0x7e, 0x8c, 0x43, 0x05, 0x95, 0x2f, 0x4c, 0x14, 0xe2, 0x12,
	0xe8, 0x46, 0x1a, 0xf7, 0x13, 0xbc, 0x8f, 0x76, 0xff, 0xd5, 0x2e, 0xe2, 0xa3, 0xe1, 0xd0, 0xfb,
	0xdd, 0x86, 0xa6, 0x1b, 0xe1, 0xce, 0x3c, 0x2e, 0x80, 0x34, 0xbe, 0x87, 0x36, 0x27, 0x58, 0x60,
	0xb0, 0xbe, 0x98, 0x01, 0x2a, 0xeb, 0xe1, 0xe5, 0x6b, 0x68, 0xab, 0x82, 0xb6, 0x0f, 0xfe, 0xb1,
	0x0f, 0x9e, 0x9f, 0x96, 0x99, 0x87, 0xf6, 0x7e, 0xb5, 0xd1, 0x46, 0x48, 0x55, 0x2e, 0xb8, 0xa2,
	0xf8, 0x06, 0x6a, 0x16, 0x02, 0xbb, 0x89, 0x3b, 0x7e, 0xdd, 0x3f, 0x2b, 0xfe, 0x23, 0xf3, 0x37,
	0xb4, 0x85, 0xf8, 0x25, 0xda, 0x32, 0xd2, 0x46, 0x15, 0x6d, 0x81, 0x6c, 0x03, 0xc0, 0xfe, 0x0c,
	0x78, 0xd6, 0x81, 0x01, 0xc4, 0xfd, 0x69, 0x1c, 0x5e, 0xc8, 0xea, 0x3f, 0xe0, 0xbb, 0xa8, 0xed,
	0x2c, 0xf5, 0x1a, 0x45, 0xc7, 0x4b, 0x73, 0x1d, 0xad, 0xe1, 0x03, 0xfb, 0x3f, 0x2c, 0xcb, 0xf1,
	0x53, 0xd4, 0x90, 0xe2, 0xc4, 0x3b, 0x53, 0xa0, 0xf6, 0xfd, 0x55, 0x96, 0xd0, 0x2f, 0xb5, 0xf0,
	0x43, 0x71, 0x12, 0x9a, 0x36, 0x9d, 0xf7, 0x4d, 0xd4, 0x80, 0x00, 0xef, 0xa0, 0x16, 0x84, 0xc6,
	0x8d, 0x0f, 0x01, 0xc8, 0xd3, 0x0c, 0x9b, 0x10, 0x82, 0xd8, 0xb7, 0x51, 0x8b, 0xc4, 0x9a, 0x1d,
	0x53, 0xef, 0x63, 0xb0, 0xd8, 0x26, 0x57, 0x8b, 0x0f, 0xea, 0x06, 0x7f, 0x0a, 0x56, 0x73, 0x38,
	0xa8, 0xac, 0x6a, 0x42, 0x55, 0x2c, 0x59, 0x5e, 0x68, 0xff, 0x39, 0xa8, 0x3b, 0xe7, 0xfa, 0x28,
	0x2d, 0x19, 0x4f, 0x6d, 0x9b, 0xed, 0x12, 0xf8, 0x70, 0x8a, 0xc3, 0x8f, 0x2b, 0x1b, 0xa3, 0xc7,
	0x39, 0x35, 0x9c, 0xbe, 0x2c, 0xc1, 0x69, 0xb2, 0x4e, 0xcf, 0x01, 0x04, 0xbc, 0x0e, 0xd1, 0xb9,
	0x49, 0x1f, 0x0e, 0xdf, 0x95, 0xf7, 0x75, 0x31, 0xa1, 0xb3, 0x25, 0x22, 0x00, 0x80, 0x9d, 0xcc,
	0x7d, 0x2c, 0xd5, 0xc9, 0xbe, 0x2d, 0x35, 0x99, 0x05, 0x56, 0x27, 0xbb, 0x83, 0xda, 0x4a, 0x48,
	0x1d, 0x71, 0xe1, 0x7d, 0x5f, 0xc6, 0x1f, 0x53, 0x1c, 0x88, 0x9a, 0xc0, 0xd5, 0x5b, 0xf0, 0x23,
	0x38, 0xdd, 0x31, 0xa8, 0x0a, 0x63, 0x04, 0xf6, 0x7e, 0xae, 0x20, 0x8c, 0x11, 0xf7, 0x41, 0x84,
	0xba, 0x4c, 0xcc, 0x2c, 0xf3, 0xf4, 0x20, 0xbf, 0x3a, 0x4c, 0x85, 0x4a, 0xde, 0x96, 0xf9, 0x64,
	0xf5, 0x9b, 0xfd, 0xba, 0x55, 0x1c, 0xc5, 0x5b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xcf,
	0xac, 0x86, 0xf5, 0x05, 0x00, 0x00,
}