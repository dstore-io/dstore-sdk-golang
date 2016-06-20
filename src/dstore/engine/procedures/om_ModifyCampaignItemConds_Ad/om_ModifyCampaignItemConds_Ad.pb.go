// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampaignItemConds_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampaignItemConds_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampaignItemConds_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampaignItemConds_Ad

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
	ItemConditionId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	ItemConditionIdNull      bool                        `protobuf:"varint,1001,opt,name=item_condition_id_null,json=itemConditionIdNull" json:"item_condition_id_null,omitempty"`
	ConditionDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=condition_description,json=conditionDescription" json:"condition_description,omitempty"`
	ConditionDescriptionNull bool                        `protobuf:"varint,1002,opt,name=condition_description_null,json=conditionDescriptionNull" json:"condition_description_null,omitempty"`
	MinNumberOfItems         *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=min_number_of_items,json=minNumberOfItems" json:"min_number_of_items,omitempty"`
	MinNumberOfItemsNull     bool                        `protobuf:"varint,1003,opt,name=min_number_of_items_null,json=minNumberOfItemsNull" json:"min_number_of_items_null,omitempty"`
	MaxNumberOfItems         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=max_number_of_items,json=maxNumberOfItems" json:"max_number_of_items,omitempty"`
	MaxNumberOfItemsNull     bool                        `protobuf:"varint,1004,opt,name=max_number_of_items_null,json=maxNumberOfItemsNull" json:"max_number_of_items_null,omitempty"`
	FromQuantity             *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=from_quantity,json=fromQuantity" json:"from_quantity,omitempty"`
	FromQuantityNull         bool                        `protobuf:"varint,1005,opt,name=from_quantity_null,json=fromQuantityNull" json:"from_quantity_null,omitempty"`
	ToQuantity               *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=to_quantity,json=toQuantity" json:"to_quantity,omitempty"`
	ToQuantityNull           bool                        `protobuf:"varint,1006,opt,name=to_quantity_null,json=toQuantityNull" json:"to_quantity_null,omitempty"`
	FromItemBasicPrice       *dstore_values.DecimalValue `protobuf:"bytes,7,opt,name=from_item_basic_price,json=fromItemBasicPrice" json:"from_item_basic_price,omitempty"`
	FromItemBasicPriceNull   bool                        `protobuf:"varint,1007,opt,name=from_item_basic_price_null,json=fromItemBasicPriceNull" json:"from_item_basic_price_null,omitempty"`
	ToItemBasicPrice         *dstore_values.DecimalValue `protobuf:"bytes,8,opt,name=to_item_basic_price,json=toItemBasicPrice" json:"to_item_basic_price,omitempty"`
	ToItemBasicPriceNull     bool                        `protobuf:"varint,1008,opt,name=to_item_basic_price_null,json=toItemBasicPriceNull" json:"to_item_basic_price_null,omitempty"`
	FromBasicPriceSum        *dstore_values.DecimalValue `protobuf:"bytes,9,opt,name=from_basic_price_sum,json=fromBasicPriceSum" json:"from_basic_price_sum,omitempty"`
	FromBasicPriceSumNull    bool                        `protobuf:"varint,1009,opt,name=from_basic_price_sum_null,json=fromBasicPriceSumNull" json:"from_basic_price_sum_null,omitempty"`
	ToBasicPriceSum          *dstore_values.DecimalValue `protobuf:"bytes,10,opt,name=to_basic_price_sum,json=toBasicPriceSum" json:"to_basic_price_sum,omitempty"`
	ToBasicPriceSumNull      bool                        `protobuf:"varint,1010,opt,name=to_basic_price_sum_null,json=toBasicPriceSumNull" json:"to_basic_price_sum_null,omitempty"`
	DeleteCondition          *dstore_values.IntegerValue `protobuf:"bytes,11,opt,name=delete_condition,json=deleteCondition" json:"delete_condition,omitempty"`
	DeleteConditionNull      bool                        `protobuf:"varint,1011,opt,name=delete_condition_null,json=deleteConditionNull" json:"delete_condition_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Parameters) GetConditionDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionDescription
	}
	return nil
}

func (m *Parameters) GetMinNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.MinNumberOfItems
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfItems
	}
	return nil
}

func (m *Parameters) GetFromQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromQuantity
	}
	return nil
}

func (m *Parameters) GetToQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToQuantity
	}
	return nil
}

func (m *Parameters) GetFromItemBasicPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromItemBasicPrice
	}
	return nil
}

func (m *Parameters) GetToItemBasicPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToItemBasicPrice
	}
	return nil
}

func (m *Parameters) GetFromBasicPriceSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromBasicPriceSum
	}
	return nil
}

func (m *Parameters) GetToBasicPriceSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToBasicPriceSum
	}
	return nil
}

func (m *Parameters) GetDeleteCondition() *dstore_values.IntegerValue {
	if m != nil {
		return m.DeleteCondition
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	ConditionId     *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
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

func (m *Response) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampaignItemConds_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampaignItemConds_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampaignItemConds_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x96, 0x5b, 0x4f, 0x13, 0x4f,
	0x18, 0xc6, 0x03, 0xfd, 0x17, 0xf8, 0xbf, 0x45, 0x29, 0x03, 0xc5, 0xa5, 0x44, 0x63, 0xf0, 0x46,
	0x63, 0x5c, 0x8c, 0x78, 0x22, 0xa2, 0x51, 0xf0, 0x54, 0x95, 0x82, 0x6b, 0x34, 0xd1, 0x9b, 0xcd,
	0x76, 0x77, 0xda, 0x4c, 0xd2, 0xdd, 0xa9, 0x33, 0x5b, 0x81, 0x6f, 0xe1, 0xa5, 0x5f, 0xd1, 0xf3,
	0xe9, 0xda, 0xc4, 0x39, 0x6c, 0xbb, 0xed, 0xec, 0xc6, 0x45, 0x6f, 0xa0, 0xd3, 0x79, 0x9f, 0xe7,
	0xf9, 0xcd, 0xce, 0xe6, 0x7d, 0x0b, 0x9b, 0x01, 0x8f, 0x29, 0xc3, 0x6b, 0x38, 0xea, 0x90, 0x08,
	0xaf, 0xf5, 0x18, 0xf5, 0x71, 0xd0, 0x67, 0x98, 0xaf, 0xd1, 0xd0, 0xdd, 0xa1, 0x01, 0x69, 0x1f,
	0x6e, 0x7b, 0x61, 0xcf, 0x23, 0x9d, 0xa8, 0x11, 0xe3, 0x70, 0x9b, 0x46, 0x01, 0x77, 0xef, 0x04,
	0xb6, 0x28, 0x8c, 0x29, 0x3a, 0xaf, 0xd5, 0xb6, 0x56, 0xdb, 0x7f, 0x94, 0xd4, 0x17, 0x92, 0xa8,
	0x37, 0x5e, 0xb7, 0x8f, 0xb9, 0x76, 0xa8, 0x2f, 0x8f, 0xe7, 0x63, 0xc6, 0x28, 0x4b, 0xb6, 0x56,
	0xc6, 0xb7, 0x42, 0xcc, 0xb9, 0xd7, 0xc1, 0xc9, 0xe6, 0x19, 0x73, 0x33, 0xf6, 0x48, 0xd4, 0xa6,
	0x2c, 0xf4, 0x62, 0x42, 0x23, 0x5d, 0xb4, 0xfa, 0xae, 0x02, 0xb0, 0xe7, 0x31, 0x4f, 0xec, 0x62,
	0xc6, 0xd1, 0x03, 0x98, 0x27, 0x02, 0xc8, 0xf5, 0x05, 0x11, 0x91, 0x65, 0x2e, 0x09, 0xac, 0x89,
	0xd3, 0x13, 0x67, 0x2b, 0x97, 0x56, 0xec, 0xe4, 0x24, 0x09, 0x1c, 0x89, 0x62, 0xdc, 0xc1, 0xec,
	0x85, 0x5c, 0x39, 0x73, 0x24, 0x39, 0x86, 0x12, 0x35, 0x02, 0x74, 0x19, 0x96, 0x32, 0x46, 0x6e,
	0xd4, 0xef, 0x76, 0xad, 0xf7, 0xd3, 0xc2, 0x6e, 0xc6, 0x59, 0x30, 0x14, 0x4d, 0xb1, 0x87, 0x76,
	0xa1, 0x96, 0x0a, 0x02, 0xcc, 0x7d, 0x46, 0x7a, 0xf2, 0xb3, 0x35, 0xa9, 0x10, 0xea, 0x06, 0x02,
	0x8f, 0x19, 0x89, 0x3a, 0x9a, 0x60, 0x71, 0x28, 0xbc, 0x9b, 0xea, 0xd0, 0x4d, 0xa8, 0xe7, 0x1a,
	0x6a, 0x94, 0x0f, 0x1a, 0xc5, 0xca, 0x93, 0x2a, 0x9e, 0x47, 0xb0, 0x10, 0x12, 0x59, 0x1c, 0xb6,
	0x30, 0x73, 0x69, 0xdb, 0x95, 0xd0, 0xdc, 0x2a, 0x15, 0x3f, 0x90, 0xaa, 0xd0, 0x35, 0x95, 0x6c,
	0xb7, 0x2d, 0xaf, 0x98, 0xa3, 0x6b, 0x60, 0xe5, 0x78, 0x69, 0x90, 0x8f, 0x1a, 0x64, 0xd1, 0x14,
	0x0d, 0x21, 0xbc, 0x83, 0x0c, 0xc4, 0x7f, 0x47, 0x81, 0xf0, 0x0e, 0xb2, 0x10, 0x59, 0x2f, 0x0d,
	0xf1, 0x69, 0x00, 0x61, 0x88, 0x14, 0xc4, 0x6d, 0x38, 0xd6, 0x66, 0xe2, 0xe5, 0x7d, 0xdd, 0xf7,
	0xa2, 0x98, 0xc4, 0x87, 0x56, 0xb9, 0x38, 0x7e, 0x56, 0x2a, 0x9e, 0x26, 0x02, 0x74, 0x01, 0xd0,
	0x98, 0x83, 0x0e, 0xfd, 0xac, 0x43, 0xab, 0xa3, 0xa5, 0x2a, 0x70, 0x13, 0x2a, 0x31, 0x4d, 0xe3,
	0xa6, 0x8a, 0xe3, 0x20, 0xa6, 0xc3, 0xb0, 0x73, 0x50, 0x1d, 0x51, 0xeb, 0xa8, 0x2f, 0x3a, 0xea,
	0x78, 0x5a, 0xa6, 0x82, 0x9a, 0x50, 0x53, 0x5c, 0xea, 0x75, 0x6d, 0x79, 0x9c, 0xf8, 0x6e, 0x8f,
	0x11, 0x1f, 0x5b, 0xd3, 0xb9, 0x91, 0x01, 0xf6, 0x49, 0xe8, 0x75, 0x75, 0xa4, 0x3a, 0x91, 0x7c,
	0x48, 0x5b, 0x52, 0xb7, 0x27, 0x65, 0xe8, 0x06, 0xd4, 0x73, 0xfd, 0x34, 0xc4, 0x57, 0x0d, 0xb1,
	0x94, 0x15, 0x0e, 0xee, 0x5a, 0x70, 0x67, 0x50, 0x66, 0x8a, 0x51, 0xc4, 0x79, 0x0d, 0x10, 0x71,
	0xd7, 0x39, 0x5e, 0x1a, 0xe3, 0x5b, 0x72, 0xd7, 0xa6, 0x48, 0x41, 0x3c, 0x81, 0x45, 0x75, 0x82,
	0x51, 0x15, 0xef, 0x87, 0xd6, 0xff, 0xc5, 0x14, 0xf3, 0x52, 0x98, 0xda, 0x3d, 0xeb, 0x87, 0x68,
	0x03, 0x96, 0xf3, 0xdc, 0x34, 0xc7, 0x77, 0xcd, 0x51, 0xcb, 0xc8, 0x14, 0xc8, 0x43, 0x40, 0xe2,
	0x04, 0x26, 0x06, 0x14, 0x63, 0xcc, 0xc5, 0x74, 0x1c, 0xe2, 0x0a, 0x9c, 0xc8, 0x3a, 0x69, 0x84,
	0x1f, 0x49, 0x3f, 0x32, 0x24, 0x0a, 0xe0, 0x3e, 0x54, 0x03, 0xdc, 0x15, 0xad, 0x31, 0xed, 0x63,
	0x56, 0xe5, 0x08, 0xdd, 0x50, 0x8b, 0x86, 0xdd, 0x0d, 0xad, 0x43, 0xcd, 0xf4, 0xd1, 0xe1, 0x3f,
	0x93, 0x70, 0x43, 0x20, 0xc3, 0x57, 0x7f, 0x4d, 0xc2, 0x8c, 0x83, 0x79, 0x8f, 0x46, 0x1c, 0xa3,
	0x8b, 0x50, 0x56, 0x8d, 0x3f, 0x69, 0xc6, 0xc3, 0x4e, 0x98, 0x8c, 0x15, 0x3d, 0x14, 0xee, 0xc9,
	0xbf, 0x8e, 0x2e, 0x44, 0x2f, 0xa1, 0x2a, 0x5b, 0xbe, 0x3b, 0xd2, 0xf3, 0x45, 0x1b, 0x2d, 0x09,
	0xb1, 0x6d, 0x88, 0xcd, 0xc9, 0xb0, 0x23, 0xd6, 0x8d, 0x74, 0xed, 0xcc, 0x85, 0xe3, 0x5f, 0xa0,
	0xeb, 0x30, 0x9d, 0x8c, 0x1a, 0xd1, 0x0a, 0xa5, 0xe3, 0xa9, 0x8c, 0xa3, 0x1e, 0x44, 0x3b, 0xfa,
	0xbf, 0x33, 0x28, 0x47, 0x8f, 0xa1, 0xc4, 0xe8, 0xbe, 0xe8, 0x5d, 0x52, 0xb5, 0x61, 0xff, 0xc5,
	0x6c, 0xb4, 0x07, 0x8f, 0xc2, 0x76, 0xe8, 0xbe, 0x23, 0x5d, 0xd0, 0x2d, 0x98, 0x1d, 0x9b, 0x53,
	0xb8, 0xf8, 0x66, 0x2a, 0x7e, 0x3a, 0x71, 0xea, 0x27, 0xa1, 0x24, 0xbc, 0xd0, 0x12, 0x4c, 0x09,
	0x37, 0x69, 0xf0, 0xb6, 0x29, 0x1c, 0xca, 0x4e, 0x59, 0x2c, 0x1b, 0xc1, 0xd6, 0x73, 0x58, 0x21,
	0xd4, 0x40, 0x4c, 0x87, 0xff, 0xab, 0xab, 0xff, 0xf6, 0xb3, 0xa0, 0x35, 0xa5, 0x06, 0xef, 0xfa,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x42, 0x1b, 0xa1, 0x57, 0x08, 0x00, 0x00,
}
