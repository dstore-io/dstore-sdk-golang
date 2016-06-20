// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignBundlePricing_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaignBundlePricing_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignBundlePricing_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignBundlePricing_Ad

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
	CampaignId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
	CampaignIdNull      bool                        `protobuf:"varint,1001,opt,name=campaign_id_null,json=campaignIdNull" json:"campaign_id_null,omitempty"`
	BenefitId           *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	BenefitIdNull       bool                        `protobuf:"varint,1002,opt,name=benefit_id_null,json=benefitIdNull" json:"benefit_id_null,omitempty"`
	GetAssignedSets     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=get_assigned_sets,json=getAssignedSets" json:"get_assigned_sets,omitempty"`
	GetAssignedSetsNull bool                        `protobuf:"varint,1003,opt,name=get_assigned_sets_null,json=getAssignedSetsNull" json:"get_assigned_sets_null,omitempty"`
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

func (m *Parameters) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Parameters) GetGetAssignedSets() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetAssignedSets
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
	NetBasedPricing          *dstore_values.BooleanValue `protobuf:"bytes,10001,opt,name=net_based_pricing,json=netBasedPricing" json:"net_based_pricing,omitempty"`
	BenefitId                *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	BundlePricingTypeId      *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=bundle_pricing_type_id,json=bundlePricingTypeId" json:"bundle_pricing_type_id,omitempty"`
	TotalQuantity            *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=total_quantity,json=totalQuantity" json:"total_quantity,omitempty"`
	BundlePriceOrDiscount    *dstore_values.DecimalValue `protobuf:"bytes,10005,opt,name=bundle_price_or_discount,json=bundlePriceOrDiscount" json:"bundle_price_or_discount,omitempty"`
	ItemConditionId          *dstore_values.IntegerValue `protobuf:"bytes,20002,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	ItemSetId                *dstore_values.IntegerValue `protobuf:"bytes,20006,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	Quantity                 *dstore_values.IntegerValue `protobuf:"bytes,20007,opt,name=quantity" json:"quantity,omitempty"`
	DistinctItemsOnly        *dstore_values.BooleanValue `protobuf:"bytes,20008,opt,name=distinct_items_only,json=distinctItemsOnly" json:"distinct_items_only,omitempty"`
	ItemConditionDescription *dstore_values.StringValue  `protobuf:"bytes,20009,opt,name=item_condition_description,json=itemConditionDescription" json:"item_condition_description,omitempty"`
	SortNo                   *dstore_values.IntegerValue `protobuf:"bytes,20010,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetNetBasedPricing() *dstore_values.BooleanValue {
	if m != nil {
		return m.NetBasedPricing
	}
	return nil
}

func (m *Response_Row) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Response_Row) GetBundlePricingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BundlePricingTypeId
	}
	return nil
}

func (m *Response_Row) GetTotalQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.TotalQuantity
	}
	return nil
}

func (m *Response_Row) GetBundlePriceOrDiscount() *dstore_values.DecimalValue {
	if m != nil {
		return m.BundlePriceOrDiscount
	}
	return nil
}

func (m *Response_Row) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Response_Row) GetItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemSetId
	}
	return nil
}

func (m *Response_Row) GetQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Response_Row) GetDistinctItemsOnly() *dstore_values.BooleanValue {
	if m != nil {
		return m.DistinctItemsOnly
	}
	return nil
}

func (m *Response_Row) GetItemConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionDescription
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignBundlePricing_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignBundlePricing_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignBundlePricing_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xc7, 0x03, 0x95, 0x52, 0x0e, 0x81, 0xd2, 0x69, 0x24, 0x6b, 0x49, 0x8c, 0xc1, 0x0b, 0x35,
	0x31, 0x8b, 0x51, 0xe3, 0x07, 0xca, 0x05, 0x1f, 0x06, 0x9b, 0x48, 0x81, 0x55, 0x49, 0xe4, 0x66,
	0xb3, 0xdd, 0x1d, 0x9a, 0x49, 0xb6, 0x33, 0x75, 0x66, 0x2a, 0xe1, 0x2d, 0xfc, 0xbc, 0xe3, 0xca,
	0x0b, 0xbf, 0x5e, 0xc1, 0x27, 0xf1, 0xc2, 0x07, 0x50, 0x5f, 0xc2, 0x33, 0x3b, 0x53, 0xfa, 0x01,
	0x6a, 0xf1, 0x06, 0x32, 0x3b, 0xe7, 0xf7, 0xff, 0x9f, 0x39, 0xe7, 0x4c, 0x07, 0x96, 0x12, 0xa5,
	0x85, 0xa4, 0x0b, 0x94, 0x37, 0x18, 0xa7, 0x0b, 0x2d, 0x29, 0x62, 0x9a, 0xb4, 0x25, 0x55, 0x0b,
	0xa2, 0x19, 0xae, 0x53, 0xbd, 0x1a, 0x35, 0x5b, 0x11, 0x6b, 0xf0, 0x95, 0x36, 0x4f, 0x52, 0xba,
	0x25, 0x59, 0xcc, 0x78, 0x23, 0x5c, 0x4e, 0x7c, 0x8c, 0xd4, 0x82, 0x5c, 0xb5, 0xb8, 0x6f, 0x71,
	0xff, 0xef, 0x4c, 0xa5, 0xec, 0xcc, 0x5e, 0x44, 0x69, 0x9b, 0x2a, 0x2b, 0x51, 0x39, 0xd7, 0x9f,
	0x01, 0x95, 0x52, 0x48, 0xb7, 0x35, 0xd7, 0xbf, 0xd5, 0xa4, 0x4a, 0x45, 0x0d, 0xea, 0x36, 0x2f,
	0x0e, 0x6e, 0xea, 0x88, 0xf1, 0x3d, 0x21, 0x9b, 0x91, 0x66, 0x82, 0xdb, 0xa0, 0xf9, 0x6f, 0xa3,
	0x00, 0x5b, 0x91, 0x8c, 0x70, 0x97, 0x4a, 0x45, 0xee, 0xc3, 0x64, 0xec, 0x92, 0x0b, 0x59, 0xe2,
	0x8d, 0x5c, 0x18, 0xb9, 0x3c, 0x79, 0x7d, 0xce, 0x77, 0x87, 0x70, 0x69, 0x31, 0xae, 0x69, 0x83,
	0xca, 0x1d, 0xb3, 0x0a, 0xa0, 0x13, 0x5f, 0x4d, 0xc8, 0x15, 0x98, 0xe9, 0xa1, 0x43, 0xde, 0x4e,
	0x53, 0xef, 0xc7, 0x38, 0x6a, 0x14, 0x82, 0xe9, 0x6e, 0x58, 0x0d, 0x3f, 0x93, 0x45, 0x80, 0x3a,
	0xe5, 0x74, 0x8f, 0x69, 0xe3, 0x33, 0xfa, 0x6f, 0x9f, 0x09, 0x17, 0x8e, 0x36, 0x97, 0xa0, 0xd8,
	0x65, 0xad, 0xcb, 0x4f, 0xeb, 0x32, 0x75, 0x14, 0x94, 0x99, 0xac, 0x43, 0xa9, 0x41, 0x75, 0x18,
	0x29, 0x85, 0xc6, 0x34, 0x09, 0x15, 0xd5, 0xca, 0xcb, 0x9d, 0xe8, 0x55, 0x17, 0x22, 0xa5, 0x11,
	0xb7, 0x5e, 0x45, 0xa4, 0x96, 0x1d, 0xf4, 0x18, 0x19, 0x72, 0x13, 0x66, 0x8f, 0x09, 0x59, 0xe3,
	0x5f, 0xd6, 0xb8, 0x3c, 0x40, 0x18, 0xfb, 0xf9, 0xef, 0x05, 0x28, 0x04, 0x54, 0xb5, 0x04, 0x57,
	0x94, 0x5c, 0x83, 0xb1, 0xac, 0x73, 0xae, 0xa6, 0x15, 0xbf, 0x7f, 0x30, 0x6c, 0x57, 0x1f, 0x98,
	0xbf, 0x81, 0x0d, 0x24, 0xcf, 0x60, 0xc6, 0xf4, 0x2c, 0xec, 0x69, 0x1a, 0x16, 0x2a, 0x87, 0xb0,
	0x3f, 0x00, 0x0f, 0xb6, 0x76, 0x03, 0xd7, 0xd5, 0xee, 0x3a, 0x28, 0x36, 0xfb, 0x3f, 0x90, 0x3b,
	0x30, 0xee, 0x66, 0x05, 0xcb, 0x61, 0x14, 0xcf, 0x1f, 0x53, 0xb4, 0x93, 0xb4, 0x61, 0xff, 0x07,
	0x9d, 0x70, 0xf2, 0x08, 0x72, 0x52, 0xec, 0x7b, 0x67, 0x32, 0x6a, 0xd1, 0x3f, 0xcd, 0x74, 0xfb,
	0x9d, 0x5a, 0xf8, 0x81, 0xd8, 0x0f, 0x8c, 0x4c, 0xe5, 0x6b, 0x1e, 0x72, 0xb8, 0x20, 0xb3, 0x90,
	0xc7, 0xa5, 0x99, 0x84, 0x97, 0x35, 0x2c, 0xcf, 0x58, 0x30, 0x86, 0x4b, 0xec, 0xf4, 0x43, 0x28,
	0x71, 0xac, 0x7b, 0x3d, 0x52, 0x58, 0xf4, 0x96, 0x55, 0xf2, 0x5e, 0xd5, 0x86, 0xe8, 0x20, 0x62,
	0x2b, 0x86, 0x72, 0xf6, 0xe4, 0x5e, 0xdf, 0xbc, 0xbd, 0xae, 0x9d, 0x6a, 0xe0, 0xb6, 0x61, 0xb6,
	0x9e, 0x1d, 0xa6, 0x93, 0x43, 0xa8, 0x0f, 0x5a, 0xd4, 0x08, 0xbd, 0x19, 0x42, 0xa8, 0x5c, 0xef,
	0x2d, 0xc4, 0x13, 0x24, 0x51, 0x72, 0x15, 0xa6, 0xb5, 0xd0, 0x51, 0x1a, 0x3e, 0x6f, 0x47, 0x5c,
	0x33, 0x7d, 0xe0, 0xbd, 0x1d, 0x42, 0x6a, 0x2a, 0x63, 0xb6, 0x1d, 0x42, 0x9e, 0x82, 0xd7, 0x93,
	0x17, 0x0d, 0x85, 0x0c, 0x13, 0xa6, 0x62, 0xd1, 0xe6, 0xda, 0x7b, 0x77, 0xb2, 0x5c, 0x42, 0x63,
	0xd6, 0x8c, 0x52, 0x2b, 0x77, 0xb6, 0x9b, 0x19, 0xdd, 0x94, 0x6b, 0x0e, 0x25, 0x55, 0x28, 0x31,
	0x4d, 0x9b, 0x61, 0x2c, 0x78, 0xc2, 0xcc, 0xbc, 0x98, 0x93, 0xbe, 0x3f, 0x1c, 0xe2, 0xc7, 0xa0,
	0x68, 0xb8, 0xd5, 0x0e, 0x86, 0xc7, 0x5c, 0x82, 0xc9, 0x4c, 0x0a, 0x2f, 0x8c, 0x11, 0xf9, 0x30,
	0x8c, 0xc8, 0x84, 0x21, 0xf0, 0x12, 0x21, 0x7e, 0x17, 0x0a, 0x47, 0xf5, 0xf9, 0x38, 0x0c, 0x7b,
	0x14, 0x4e, 0x36, 0xa0, 0x8c, 0xb5, 0xd0, 0x8c, 0xc7, 0xe8, 0x8c, 0x82, 0x2a, 0x14, 0x3c, 0x3d,
	0xf0, 0x3e, 0xfd, 0x41, 0xa5, 0x6f, 0x7a, 0x4a, 0x1d, 0xb2, 0x6a, 0xc0, 0x4d, 0xe4, 0xc8, 0x2e,
	0x54, 0x06, 0x6a, 0x92, 0x50, 0x15, 0x4b, 0xd6, 0xca, 0xae, 0xe5, 0xe7, 0xc3, 0x81, 0x5b, 0xed,
	0x54, 0x95, 0x96, 0xd8, 0x72, 0x2b, 0xea, 0xf5, 0xd5, 0x66, 0xad, 0x4b, 0x93, 0x5b, 0x30, 0xae,
	0x84, 0xd4, 0x21, 0x17, 0xde, 0x97, 0x61, 0x0e, 0x99, 0x37, 0xd1, 0x35, 0xb1, 0xb2, 0x03, 0x73,
	0x4c, 0x0c, 0x5c, 0xc1, 0xee, 0xfb, 0xb4, 0x7b, 0xfb, 0x3f, 0x5f, 0xae, 0x7a, 0x3e, 0x7b, 0x1a,
	0x6e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xa1, 0x06, 0x0b, 0xfb, 0x06, 0x00, 0x00,
}
