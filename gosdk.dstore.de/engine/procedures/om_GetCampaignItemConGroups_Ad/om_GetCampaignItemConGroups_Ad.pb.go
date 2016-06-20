// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignItemConGroups_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaignItemConGroups_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignItemConGroups_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignItemConGroups_Ad

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
	ConditionId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	ConditionIdNull          bool                        `protobuf:"varint,1001,opt,name=condition_id_null,json=conditionIdNull" json:"condition_id_null,omitempty"`
	ItemConditionGroupId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=item_condition_group_id,json=itemConditionGroupId" json:"item_condition_group_id,omitempty"`
	ItemConditionGroupIdNull bool                        `protobuf:"varint,1002,opt,name=item_condition_group_id_null,json=itemConditionGroupIdNull" json:"item_condition_group_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
	}
	return nil
}

func (m *Parameters) GetItemConditionGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionGroupId
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
	RowId                        int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Condition1                   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=condition1" json:"condition1,omitempty"`
	Condition2                   *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=condition2" json:"condition2,omitempty"`
	RecursiveEvaluation          *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=recursive_evaluation,json=recursiveEvaluation" json:"recursive_evaluation,omitempty"`
	ToBasicPriceSumPart          *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=to_basic_price_sum_part,json=toBasicPriceSumPart" json:"to_basic_price_sum_part,omitempty"`
	DomainTreeNodeIds            *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=domain_tree_node_ids,json=domainTreeNodeIds" json:"domain_tree_node_ids,omitempty"`
	NodeCharacteristicId         *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	ItemConditionPartId          *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=item_condition_part_id,json=itemConditionPartId" json:"item_condition_part_id,omitempty"`
	FromBasicPriceSumPart        *dstore_values.DecimalValue `protobuf:"bytes,10008,opt,name=from_basic_price_sum_part,json=fromBasicPriceSumPart" json:"from_basic_price_sum_part,omitempty"`
	ToQuantity                   *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=to_quantity,json=toQuantity" json:"to_quantity,omitempty"`
	MaxNumberOfItems             *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=max_number_of_items,json=maxNumberOfItems" json:"max_number_of_items,omitempty"`
	ToItemBasicPricePart         *dstore_values.DecimalValue `protobuf:"bytes,10011,opt,name=to_item_basic_price_part,json=toItemBasicPricePart" json:"to_item_basic_price_part,omitempty"`
	ItemConditionPartDescription *dstore_values.StringValue  `protobuf:"bytes,10012,opt,name=item_condition_part_description,json=itemConditionPartDescription" json:"item_condition_part_description,omitempty"`
	LevelIds                     *dstore_values.StringValue  `protobuf:"bytes,10013,opt,name=level_ids,json=levelIds" json:"level_ids,omitempty"`
	ItemConditionGroupId         *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=item_condition_group_id,json=itemConditionGroupId" json:"item_condition_group_id,omitempty"`
	ExtendedItemCondGroupDescr   *dstore_values.StringValue  `protobuf:"bytes,10015,opt,name=extended_item_cond_group_descr,json=extendedItemCondGroupDescr" json:"extended_item_cond_group_descr,omitempty"`
	FromQuantity                 *dstore_values.IntegerValue `protobuf:"bytes,10016,opt,name=from_quantity,json=fromQuantity" json:"from_quantity,omitempty"`
	FromItemBasicPrice           *dstore_values.DecimalValue `protobuf:"bytes,10017,opt,name=from_item_basic_price,json=fromItemBasicPrice" json:"from_item_basic_price,omitempty"`
	ItemGroupSortNo              *dstore_values.IntegerValue `protobuf:"bytes,10018,opt,name=item_group_sort_no,json=itemGroupSortNo" json:"item_group_sort_no,omitempty"`
	Operator1                    *dstore_values.StringValue  `protobuf:"bytes,10019,opt,name=operator1" json:"operator1,omitempty"`
	Operator2                    *dstore_values.StringValue  `protobuf:"bytes,10020,opt,name=operator2" json:"operator2,omitempty"`
	CombinePartsWithANDOperator  *dstore_values.BooleanValue `protobuf:"bytes,10021,opt,name=combine_parts_with_a_n_d_operator,json=combinePartsWithANDOperator" json:"combine_parts_with_a_n_d_operator,omitempty"`
	MinNumberOfItemsPart         *dstore_values.IntegerValue `protobuf:"bytes,10022,opt,name=min_number_of_items_part,json=minNumberOfItemsPart" json:"min_number_of_items_part,omitempty"`
	ItemPartSortNo               *dstore_values.IntegerValue `protobuf:"bytes,10023,opt,name=item_part_sort_no,json=itemPartSortNo" json:"item_part_sort_no,omitempty"`
	ExtendedItemCondPartDescr    *dstore_values.StringValue  `protobuf:"bytes,10024,opt,name=extended_item_cond_part_descr,json=extendedItemCondPartDescr" json:"extended_item_cond_part_descr,omitempty"`
	MaxNumberOfItemsPart         *dstore_values.IntegerValue `protobuf:"bytes,10025,opt,name=max_number_of_items_part,json=maxNumberOfItemsPart" json:"max_number_of_items_part,omitempty"`
	FromItemBasicPricePart       *dstore_values.DecimalValue `protobuf:"bytes,10026,opt,name=from_item_basic_price_part,json=fromItemBasicPricePart" json:"from_item_basic_price_part,omitempty"`
	MinNumberOfItems             *dstore_values.IntegerValue `protobuf:"bytes,10027,opt,name=min_number_of_items,json=minNumberOfItems" json:"min_number_of_items,omitempty"`
	FromBasicPriceSum            *dstore_values.DecimalValue `protobuf:"bytes,10028,opt,name=from_basic_price_sum,json=fromBasicPriceSum" json:"from_basic_price_sum,omitempty"`
	ToBasicPriceSum              *dstore_values.DecimalValue `protobuf:"bytes,10029,opt,name=to_basic_price_sum,json=toBasicPriceSum" json:"to_basic_price_sum,omitempty"`
	ConditionId                  *dstore_values.IntegerValue `protobuf:"bytes,10030,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	ToItemBasicPrice             *dstore_values.DecimalValue `protobuf:"bytes,10031,opt,name=to_item_basic_price,json=toItemBasicPrice" json:"to_item_basic_price,omitempty"`
	ToQuantityPart               *dstore_values.IntegerValue `protobuf:"bytes,10032,opt,name=to_quantity_part,json=toQuantityPart" json:"to_quantity_part,omitempty"`
	FromQuantityPart             *dstore_values.IntegerValue `protobuf:"bytes,10033,opt,name=from_quantity_part,json=fromQuantityPart" json:"from_quantity_part,omitempty"`
	InheritDepth                 *dstore_values.IntegerValue `protobuf:"bytes,10034,opt,name=inherit_depth,json=inheritDepth" json:"inherit_depth,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCondition1() *dstore_values.StringValue {
	if m != nil {
		return m.Condition1
	}
	return nil
}

func (m *Response_Row) GetCondition2() *dstore_values.StringValue {
	if m != nil {
		return m.Condition2
	}
	return nil
}

func (m *Response_Row) GetRecursiveEvaluation() *dstore_values.IntegerValue {
	if m != nil {
		return m.RecursiveEvaluation
	}
	return nil
}

func (m *Response_Row) GetToBasicPriceSumPart() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToBasicPriceSumPart
	}
	return nil
}

func (m *Response_Row) GetDomainTreeNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.DomainTreeNodeIds
	}
	return nil
}

func (m *Response_Row) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetItemConditionPartId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionPartId
	}
	return nil
}

func (m *Response_Row) GetFromBasicPriceSumPart() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromBasicPriceSumPart
	}
	return nil
}

func (m *Response_Row) GetToQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToQuantity
	}
	return nil
}

func (m *Response_Row) GetMaxNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfItems
	}
	return nil
}

func (m *Response_Row) GetToItemBasicPricePart() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToItemBasicPricePart
	}
	return nil
}

func (m *Response_Row) GetItemConditionPartDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionPartDescription
	}
	return nil
}

func (m *Response_Row) GetLevelIds() *dstore_values.StringValue {
	if m != nil {
		return m.LevelIds
	}
	return nil
}

func (m *Response_Row) GetItemConditionGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionGroupId
	}
	return nil
}

func (m *Response_Row) GetExtendedItemCondGroupDescr() *dstore_values.StringValue {
	if m != nil {
		return m.ExtendedItemCondGroupDescr
	}
	return nil
}

func (m *Response_Row) GetFromQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromQuantity
	}
	return nil
}

func (m *Response_Row) GetFromItemBasicPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromItemBasicPrice
	}
	return nil
}

func (m *Response_Row) GetItemGroupSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemGroupSortNo
	}
	return nil
}

func (m *Response_Row) GetOperator1() *dstore_values.StringValue {
	if m != nil {
		return m.Operator1
	}
	return nil
}

func (m *Response_Row) GetOperator2() *dstore_values.StringValue {
	if m != nil {
		return m.Operator2
	}
	return nil
}

func (m *Response_Row) GetCombinePartsWithANDOperator() *dstore_values.BooleanValue {
	if m != nil {
		return m.CombinePartsWithANDOperator
	}
	return nil
}

func (m *Response_Row) GetMinNumberOfItemsPart() *dstore_values.IntegerValue {
	if m != nil {
		return m.MinNumberOfItemsPart
	}
	return nil
}

func (m *Response_Row) GetItemPartSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemPartSortNo
	}
	return nil
}

func (m *Response_Row) GetExtendedItemCondPartDescr() *dstore_values.StringValue {
	if m != nil {
		return m.ExtendedItemCondPartDescr
	}
	return nil
}

func (m *Response_Row) GetMaxNumberOfItemsPart() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfItemsPart
	}
	return nil
}

func (m *Response_Row) GetFromItemBasicPricePart() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromItemBasicPricePart
	}
	return nil
}

func (m *Response_Row) GetMinNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.MinNumberOfItems
	}
	return nil
}

func (m *Response_Row) GetFromBasicPriceSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromBasicPriceSum
	}
	return nil
}

func (m *Response_Row) GetToBasicPriceSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToBasicPriceSum
	}
	return nil
}

func (m *Response_Row) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
	}
	return nil
}

func (m *Response_Row) GetToItemBasicPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToItemBasicPrice
	}
	return nil
}

func (m *Response_Row) GetToQuantityPart() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToQuantityPart
	}
	return nil
}

func (m *Response_Row) GetFromQuantityPart() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromQuantityPart
	}
	return nil
}

func (m *Response_Row) GetInheritDepth() *dstore_values.IntegerValue {
	if m != nil {
		return m.InheritDepth
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignItemConGroups_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignItemConGroups_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignItemConGroups_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetCampaignItemConGroups_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x57, 0x5b, 0x73, 0x1b, 0x35,
	0x14, 0x9e, 0x36, 0xa4, 0x4d, 0x95, 0x42, 0x12, 0xc5, 0xa4, 0x8a, 0x53, 0x0a, 0x94, 0x17, 0x66,
	0x60, 0x1c, 0x6a, 0x5e, 0xb8, 0x4c, 0x81, 0x34, 0x29, 0x19, 0x03, 0xd9, 0xb4, 0x9b, 0x4e, 0xb9,
	0xcc, 0x30, 0x9a, 0xf5, 0xae, 0xe2, 0x68, 0xf0, 0xae, 0x8c, 0x24, 0x27, 0xe5, 0x4f, 0x30, 0x5c,
	0xcb, 0xfd, 0x7e, 0xbf, 0x5f, 0xfe, 0x15, 0x3c, 0xf3, 0x03, 0x38, 0x92, 0xd6, 0x97, 0x5d, 0xbb,
	0x58, 0xe6, 0x25, 0x99, 0xb5, 0xce, 0xf7, 0xf9, 0x9c, 0xef, 0x3b, 0xe7, 0x68, 0x8d, 0x2e, 0x26,
	0x4a, 0x0b, 0xc9, 0xd6, 0x59, 0xd6, 0xe2, 0x19, 0x5b, 0xef, 0x48, 0x11, 0xb3, 0xa4, 0x2b, 0x99,
	0x5a, 0x17, 0x29, 0xdd, 0x66, 0x7a, 0x33, 0x4a, 0x3b, 0x11, 0x6f, 0x65, 0x0d, 0xcd, 0xd2, 0x4d,
	0x91, 0x6d, 0x4b, 0xd1, 0xed, 0x28, 0xba, 0x91, 0xd4, 0x20, 0x52, 0x0b, 0xfc, 0xa0, 0x83, 0xd7,
	0x1c, 0xbc, 0xf6, 0xdf, 0x98, 0xea, 0x72, 0xfe, 0x65, 0x87, 0x51, 0xbb, 0xcb, 0x94, 0xa3, 0xa8,
	0xae, 0x16, 0x33, 0x60, 0x52, 0x0a, 0x99, 0x1f, 0xad, 0x15, 0x8f, 0x52, 0xa6, 0x54, 0xd4, 0x62,
	0xf9, 0xe1, 0x7d, 0xe5, 0x43, 0x1d, 0xf1, 0x6c, 0x5f, 0xc8, 0x34, 0xd2, 0x5c, 0x64, 0x2e, 0xe8,
	0xfc, 0xeb, 0xc7, 0x11, 0xba, 0x12, 0xc9, 0x08, 0x4e, 0x99, 0x54, 0xf8, 0x09, 0x74, 0x3a, 0x16,
	0x59, 0xc2, 0x4d, 0x04, 0xe5, 0x09, 0x39, 0x76, 0xcf, 0xb1, 0xfb, 0xe7, 0xeb, 0x6b, 0xb5, 0xbc,
	0x8a, 0x3c, 0x2f, 0x9e, 0x69, 0xd6, 0x62, 0xf2, 0xba, 0x79, 0x0a, 0xe7, 0xfb, 0x80, 0x46, 0x82,
	0x1f, 0x40, 0x4b, 0xc3, 0x78, 0x9a, 0x75, 0xdb, 0x6d, 0xf2, 0xd7, 0x49, 0x60, 0x99, 0x0b, 0x17,
	0x86, 0x02, 0x03, 0xf8, 0x1c, 0x87, 0xe8, 0x0c, 0x07, 0x05, 0xe8, 0x00, 0xd1, 0x32, 0x4a, 0x98,
	0xef, 0x3d, 0x3e, 0xf9, 0x7b, 0x2b, 0xdc, 0xa9, 0xe7, 0xa0, 0x56, 0x43, 0x48, 0xe0, 0x49, 0x74,
	0xf6, 0x16, 0x9c, 0x2e, 0x97, 0xbf, 0x5d, 0x2e, 0x64, 0x1c, 0xd8, 0x24, 0x75, 0xfe, 0xe6, 0x19,
	0x34, 0x17, 0x32, 0xd5, 0x11, 0x99, 0x62, 0xf8, 0x21, 0x34, 0x6b, 0xe5, 0xce, 0x75, 0xa8, 0xd6,
	0x8a, 0x6e, 0x3a, 0x2b, 0x2e, 0x9b, 0xbf, 0xa1, 0x0b, 0xc4, 0x2f, 0xa2, 0x45, 0x23, 0x34, 0x1d,
	0x52, 0x1a, 0x8a, 0x99, 0x01, 0x70, 0xad, 0x04, 0x2e, 0xfb, 0xb1, 0x03, 0xcf, 0x8d, 0xc1, 0x73,
	0xb8, 0x90, 0x16, 0x3f, 0xc0, 0x8f, 0xa0, 0x93, 0xb9, 0xc1, 0x64, 0xc6, 0x32, 0x9e, 0x1b, 0x61,
	0x74, 0xf6, 0xef, 0xb8, 0xff, 0x61, 0x2f, 0x1c, 0x3f, 0x87, 0x66, 0xa4, 0x38, 0x22, 0xb7, 0x59,
	0xd4, 0x63, 0xb5, 0x69, 0x5a, 0xb2, 0xd6, 0xd3, 0xa2, 0x16, 0x8a, 0xa3, 0xd0, 0xd0, 0x54, 0xff,
	0xa9, 0xa0, 0x19, 0x78, 0xc0, 0x2b, 0xe8, 0x04, 0x3c, 0x1a, 0xb7, 0xde, 0x08, 0x40, 0x9e, 0xd9,
	0x70, 0x16, 0x1e, 0xc1, 0x82, 0xc7, 0x11, 0xea, 0xab, 0x7f, 0x81, 0xbc, 0x19, 0x14, 0xa5, 0xcb,
	0xad, 0x54, 0x5a, 0xf2, 0xac, 0xe5, 0x9c, 0x1c, 0x0a, 0x2f, 0x80, 0xeb, 0xe4, 0xad, 0x69, 0xc0,
	0x75, 0xbc, 0x8b, 0x2a, 0x92, 0xc5, 0x5d, 0xa9, 0xf8, 0x21, 0xa3, 0xcc, 0x04, 0x3b, 0x03, 0xde,
	0x0e, 0x26, 0xb7, 0xd3, 0x72, 0x1f, 0x79, 0xb9, 0x0f, 0x34, 0x1d, 0xaa, 0x05, 0x6d, 0x46, 0x8a,
	0xc7, 0xb4, 0x23, 0x79, 0xcc, 0xa8, 0xea, 0xa6, 0xb4, 0x13, 0x49, 0x4d, 0xde, 0x19, 0xcf, 0x99,
	0xb0, 0x98, 0xa7, 0x51, 0x3b, 0xe7, 0xd4, 0xe2, 0x92, 0xc1, 0x5e, 0x31, 0xd0, 0xbd, 0x6e, 0x0a,
	0x83, 0xa6, 0xc1, 0x8c, 0x4a, 0x22, 0x52, 0xb0, 0x9e, 0x6a, 0xc9, 0x18, 0xcd, 0x44, 0xc2, 0x40,
	0x43, 0x45, 0xde, 0x9d, 0x5c, 0xeb, 0x92, 0x03, 0x5e, 0x03, 0x5c, 0x00, 0xb0, 0x46, 0xa2, 0x20,
	0xc3, 0x15, 0xcb, 0x10, 0x1f, 0xc0, 0x10, 0xc7, 0x30, 0xc3, 0x5c, 0x69, 0x48, 0x16, 0x4c, 0xb9,
	0xe9, 0x51, 0x74, 0xc5, 0x60, 0x37, 0x0b, 0x50, 0x30, 0xf0, 0x2a, 0x5a, 0x29, 0xcd, 0x90, 0xa9,
	0xd8, 0x70, 0xbe, 0xe7, 0x23, 0x64, 0x61, 0xb4, 0x4c, 0xc9, 0x40, 0x79, 0x1d, 0xad, 0xee, 0x4b,
	0xe8, 0xb4, 0xb1, 0x52, 0xbe, 0xef, 0x21, 0xe5, 0x9d, 0x06, 0x3e, 0x2a, 0xe6, 0x45, 0x34, 0x0f,
	0x06, 0xbd, 0xda, 0x8d, 0x32, 0xcd, 0xf5, 0x6b, 0xe4, 0x03, 0x8f, 0xfc, 0x90, 0x16, 0x57, 0xf3,
	0x78, 0xfc, 0x2c, 0x5a, 0x4e, 0xa3, 0x1b, 0xb0, 0x19, 0xd2, 0x26, 0x93, 0x54, 0xec, 0x53, 0x93,
	0xbb, 0x22, 0x1f, 0x7a, 0xd0, 0x2c, 0x02, 0x30, 0xb0, 0xb8, 0xdd, 0x7d, 0x33, 0x34, 0x0a, 0x5f,
	0x43, 0x04, 0x72, 0xb1, 0xca, 0x0d, 0x97, 0x69, 0x4b, 0xfc, 0xc8, 0xa3, 0xc4, 0x8a, 0x16, 0x86,
	0x68, 0x50, 0xa4, 0xad, 0xb0, 0x89, 0xee, 0x1e, 0x67, 0x46, 0xc2, 0x54, 0x2c, 0x79, 0xc7, 0xb6,
	0xf7, 0xc7, 0x93, 0x3b, 0xe7, 0xec, 0x88, 0x29, 0x5b, 0x03, 0x02, 0xd8, 0x2c, 0xa7, 0xda, 0xec,
	0x90, 0xb5, 0x6d, 0x1f, 0x7e, 0x32, 0x99, 0x6d, 0xce, 0x46, 0x9b, 0xf6, 0xdb, 0xbb, 0xf5, 0x0a,
	0xff, 0x34, 0xf8, 0xbf, 0x3b, 0x9c, 0xa2, 0x73, 0xec, 0x86, 0x66, 0x59, 0xc2, 0x12, 0xda, 0x67,
	0xcf, 0x89, 0x6d, 0xdd, 0xe4, 0xb3, 0xc9, 0x39, 0x56, 0x7b, 0x14, 0xf9, 0x46, 0x4b, 0x2c, 0xbb,
	0xad, 0x1a, 0x6f, 0xa0, 0xdb, 0x6d, 0x37, 0xf6, 0xfb, 0xe6, 0x73, 0x8f, 0x5c, 0x4f, 0x1b, 0x48,
	0xbf, 0x73, 0x76, 0x91, 0xed, 0xc8, 0x11, 0xbb, 0xc9, 0x17, 0x1e, 0x4e, 0x63, 0x03, 0x2d, 0x7a,
	0x8d, 0x1b, 0x08, 0x5b, 0x2e, 0x57, 0xa6, 0x12, 0xe0, 0x71, 0x26, 0xc8, 0x97, 0x1e, 0x89, 0x2d,
	0x18, 0x9c, 0xad, 0x6e, 0x0f, 0x50, 0x81, 0xc0, 0x8f, 0xa2, 0x53, 0xa2, 0xc3, 0x64, 0x04, 0x88,
	0x0b, 0xe4, 0xab, 0xc9, 0x52, 0x0d, 0xa2, 0x87, 0xa1, 0x75, 0xf2, 0xf5, 0x14, 0xd0, 0x3a, 0x4e,
	0xd0, 0xbd, 0xb1, 0x48, 0x9b, 0x70, 0xa5, 0xd8, 0x0e, 0x55, 0xf4, 0x88, 0xeb, 0x03, 0x1a, 0xd1,
	0x8c, 0x26, 0xb4, 0x17, 0x45, 0xbe, 0x19, 0x5f, 0x4f, 0x53, 0x88, 0x36, 0x8b, 0x32, 0xc7, 0xb9,
	0x96, 0xd3, 0x98, 0x2e, 0x55, 0xcf, 0x03, 0xc9, 0x46, 0xb0, 0xb5, 0x9b, 0x13, 0x98, 0x21, 0x4b,
	0x61, 0x75, 0x96, 0x26, 0xd6, 0x0d, 0xd9, 0xb7, 0x3e, 0x1d, 0x07, 0xe8, 0xc2, 0xd8, 0xda, 0x21,
	0xdb, 0x46, 0x4b, 0x56, 0x7c, 0x3b, 0x5a, 0x3d, 0xed, 0xbf, 0xf3, 0xa0, 0xbb, 0xc3, 0xc0, 0x0c,
	0x45, 0x2e, 0xfd, 0xcb, 0xe8, 0xae, 0x31, 0xad, 0x3b, 0x98, 0x58, 0xf2, 0xfd, 0x64, 0x4d, 0x57,
	0xcb, 0x9d, 0xdb, 0x1f, 0x57, 0x5b, 0xfd, 0xe8, 0xbe, 0x72, 0xd5, 0xff, 0xe0, 0x55, 0x7d, 0x69,
	0x69, 0xd9, 0xea, 0x5f, 0x40, 0xd5, 0xb1, 0xbd, 0xec, 0x78, 0x7f, 0xf4, 0x68, 0xe8, 0x95, 0xd1,
	0x86, 0xb6, 0xcc, 0x66, 0xbf, 0x8e, 0xba, 0x45, 0x7e, 0xf2, 0xda, 0xaf, 0x25, 0xa3, 0xf0, 0x0e,
	0xaa, 0x8c, 0xbb, 0x43, 0xc8, 0xcf, 0x1e, 0x09, 0x2e, 0x8d, 0x5c, 0x1f, 0x66, 0xe0, 0x46, 0xef,
	0x76, 0xf2, 0x8b, 0x07, 0xd9, 0x42, 0xe9, 0x5a, 0x87, 0x97, 0xce, 0xe2, 0x5b, 0xf3, 0xaf, 0xc1,
	0x94, 0xaf, 0xcd, 0xa0, 0xd3, 0x98, 0xab, 0x83, 0xfc, 0xe6, 0x91, 0xcc, 0x62, 0xf9, 0xd6, 0xc0,
	0x4f, 0xa3, 0xc5, 0xa1, 0x3b, 0xd1, 0x99, 0xf8, 0xbb, 0x4f, 0x2f, 0x0f, 0x2e, 0x46, 0x6b, 0xde,
	0x33, 0x08, 0x17, 0xb6, 0xa4, 0x63, 0xfa, 0xc3, 0xc7, 0xbb, 0xe1, 0x55, 0x69, 0xb9, 0x60, 0xe3,
	0xf2, 0xec, 0x00, 0x5e, 0x31, 0xcc, 0x1c, 0x74, 0xf4, 0x01, 0xf9, 0xd3, 0x67, 0xe3, 0xe6, 0x90,
	0x2d, 0x83, 0xb8, 0x44, 0xd1, 0x1a, 0x17, 0xa5, 0x77, 0xd7, 0xc1, 0xaf, 0xb1, 0x97, 0x9e, 0x6a,
	0x09, 0x95, 0xbc, 0xd2, 0x3b, 0x4f, 0xa6, 0xff, 0xc1, 0xd6, 0x3c, 0x61, 0x7f, 0x11, 0x3d, 0xfc,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x0d, 0xd4, 0xb5, 0xf2, 0x0d, 0x00, 0x00,
}