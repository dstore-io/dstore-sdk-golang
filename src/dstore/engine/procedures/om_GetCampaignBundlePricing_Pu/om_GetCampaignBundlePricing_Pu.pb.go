// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignBundlePricing_Pu.proto
// DO NOT EDIT!

/*
Package om_GetCampaignBundlePricing_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignBundlePricing_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignBundlePricing_Pu

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
	BenefitId                         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=benefit_id,json=benefitId" json:"benefit_id,omitempty"`
	BenefitIdNull                     bool                        `protobuf:"varint,1001,opt,name=benefit_id_null,json=benefitIdNull" json:"benefit_id_null,omitempty"`
	OnlyDefinition                    *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=only_definition,json=onlyDefinition" json:"only_definition,omitempty"`
	OnlyDefinitionNull                bool                        `protobuf:"varint,1002,opt,name=only_definition_null,json=onlyDefinitionNull" json:"only_definition_null,omitempty"`
	FilterByItemSetId                 *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=filter_by_item_set_id,json=filterByItemSetId" json:"filter_by_item_set_id,omitempty"`
	FilterByItemSetIdNull             bool                        `protobuf:"varint,1003,opt,name=filter_by_item_set_id_null,json=filterByItemSetIdNull" json:"filter_by_item_set_id_null,omitempty"`
	SortByCharacteristicIdList        *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=sort_by_characteristic_id_list,json=sortByCharacteristicIdList" json:"sort_by_characteristic_id_list,omitempty"`
	SortByCharacteristicIdListNull    bool                        `protobuf:"varint,1004,opt,name=sort_by_characteristic_id_list_null,json=sortByCharacteristicIdListNull" json:"sort_by_characteristic_id_list_null,omitempty"`
	SortOptionList                    *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=sort_option_list,json=sortOptionList" json:"sort_option_list,omitempty"`
	SortOptionListNull                bool                        `protobuf:"varint,1005,opt,name=sort_option_list_null,json=sortOptionListNull" json:"sort_option_list_null,omitempty"`
	InheritDepthOptionList            *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=inherit_depth_option_list,json=inheritDepthOptionList" json:"inherit_depth_option_list,omitempty"`
	InheritDepthOptionListNull        bool                        `protobuf:"varint,1006,opt,name=inherit_depth_option_list_null,json=inheritDepthOptionListNull" json:"inherit_depth_option_list_null,omitempty"`
	RecursiveEvaluationOptionList     *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=recursive_evaluation_option_list,json=recursiveEvaluationOptionList" json:"recursive_evaluation_option_list,omitempty"`
	RecursiveEvaluationOptionListNull bool                        `protobuf:"varint,1007,opt,name=recursive_evaluation_option_list_null,json=recursiveEvaluationOptionListNull" json:"recursive_evaluation_option_list_null,omitempty"`
	GetValuesForSortByCharacs         *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=get_values_for_sort_by_characs,json=getValuesForSortByCharacs" json:"get_values_for_sort_by_characs,omitempty"`
	GetValuesForSortByCharacsNull     bool                        `protobuf:"varint,1008,opt,name=get_values_for_sort_by_characs_null,json=getValuesForSortByCharacsNull" json:"get_values_for_sort_by_characs_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetBenefitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitId
	}
	return nil
}

func (m *Parameters) GetOnlyDefinition() *dstore_values.IntegerValue {
	if m != nil {
		return m.OnlyDefinition
	}
	return nil
}

func (m *Parameters) GetFilterByItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilterByItemSetId
	}
	return nil
}

func (m *Parameters) GetSortByCharacteristicIdList() *dstore_values.StringValue {
	if m != nil {
		return m.SortByCharacteristicIdList
	}
	return nil
}

func (m *Parameters) GetSortOptionList() *dstore_values.StringValue {
	if m != nil {
		return m.SortOptionList
	}
	return nil
}

func (m *Parameters) GetInheritDepthOptionList() *dstore_values.StringValue {
	if m != nil {
		return m.InheritDepthOptionList
	}
	return nil
}

func (m *Parameters) GetRecursiveEvaluationOptionList() *dstore_values.StringValue {
	if m != nil {
		return m.RecursiveEvaluationOptionList
	}
	return nil
}

func (m *Parameters) GetGetValuesForSortByCharacs() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetValuesForSortByCharacs
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
	RowId                         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	NetBasedPricing               *dstore_values.BooleanValue `protobuf:"bytes,10001,opt,name=net_based_pricing,json=netBasedPricing" json:"net_based_pricing,omitempty"`
	ItemConditionId               *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=item_condition_id,json=itemConditionId" json:"item_condition_id,omitempty"`
	BundlePricingTypeId           *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=bundle_pricing_type_id,json=bundlePricingTypeId" json:"bundle_pricing_type_id,omitempty"`
	BundlePricingType             *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=bundle_pricing_type,json=bundlePricingType" json:"bundle_pricing_type,omitempty"`
	ItemSetId                     *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=item_set_id,json=itemSetId" json:"item_set_id,omitempty"`
	Quantity                      *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=quantity" json:"quantity,omitempty"`
	DistinctItemsOnly             *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=distinct_items_only,json=distinctItemsOnly" json:"distinct_items_only,omitempty"`
	ItemConditionDescription      *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=item_condition_description,json=itemConditionDescription" json:"item_condition_description,omitempty"`
	SortNo                        *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	BundlePriceOrDiscount         *dstore_values.DecimalValue `protobuf:"bytes,10010,opt,name=bundle_price_or_discount,json=bundlePriceOrDiscount" json:"bundle_price_or_discount,omitempty"`
	Condition1                    *dstore_values.StringValue  `protobuf:"bytes,20001,opt,name=condition1" json:"condition1,omitempty"`
	Condition2                    *dstore_values.StringValue  `protobuf:"bytes,20002,opt,name=condition2" json:"condition2,omitempty"`
	RecursiveEvaluation           *dstore_values.IntegerValue `protobuf:"bytes,20003,opt,name=recursive_evaluation,json=recursiveEvaluation" json:"recursive_evaluation,omitempty"`
	DomainTreeNodeIds             *dstore_values.StringValue  `protobuf:"bytes,20004,opt,name=domain_tree_node_ids,json=domainTreeNodeIds" json:"domain_tree_node_ids,omitempty"`
	NodeCharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,20005,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	ItemConditionPartId           *dstore_values.IntegerValue `protobuf:"bytes,20006,opt,name=item_condition_part_id,json=itemConditionPartId" json:"item_condition_part_id,omitempty"`
	CombineGroupsWithANDOperator  *dstore_values.BooleanValue `protobuf:"bytes,20008,opt,name=combine_groups_with_a_n_d_operator,json=combineGroupsWithANDOperator" json:"combine_groups_with_a_n_d_operator,omitempty"`
	ItemConditionPartDescription  *dstore_values.StringValue  `protobuf:"bytes,20009,opt,name=item_condition_part_description,json=itemConditionPartDescription" json:"item_condition_part_description,omitempty"`
	LevelIds                      *dstore_values.StringValue  `protobuf:"bytes,20011,opt,name=level_ids,json=levelIds" json:"level_ids,omitempty"`
	ItemConditionGroupId          *dstore_values.IntegerValue `protobuf:"bytes,20012,opt,name=item_condition_group_id,json=itemConditionGroupId" json:"item_condition_group_id,omitempty"`
	ItemGroupSortNo               *dstore_values.IntegerValue `protobuf:"bytes,20013,opt,name=item_group_sort_no,json=itemGroupSortNo" json:"item_group_sort_no,omitempty"`
	Operator1                     *dstore_values.StringValue  `protobuf:"bytes,20014,opt,name=operator1" json:"operator1,omitempty"`
	Operator2                     *dstore_values.StringValue  `protobuf:"bytes,20016,opt,name=operator2" json:"operator2,omitempty"`
	ItemConditionGroupDescription *dstore_values.StringValue  `protobuf:"bytes,20017,opt,name=item_condition_group_description,json=itemConditionGroupDescription" json:"item_condition_group_description,omitempty"`
	CombinePartsWithANDOperator   *dstore_values.BooleanValue `protobuf:"bytes,20018,opt,name=combine_parts_with_a_n_d_operator,json=combinePartsWithANDOperator" json:"combine_parts_with_a_n_d_operator,omitempty"`
	ItemPartSortNo                *dstore_values.IntegerValue `protobuf:"bytes,20019,opt,name=item_part_sort_no,json=itemPartSortNo" json:"item_part_sort_no,omitempty"`
	InheritDepth                  *dstore_values.IntegerValue `protobuf:"bytes,20027,opt,name=inherit_depth,json=inheritDepth" json:"inherit_depth,omitempty"`
	NodeDescription               *dstore_values.StringValue  `protobuf:"bytes,30001,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Value2                        *dstore_values.StringValue  `protobuf:"bytes,30002,opt,name=value2" json:"value2,omitempty"`
	Value3                        *dstore_values.StringValue  `protobuf:"bytes,30003,opt,name=value3" json:"value3,omitempty"`
	Value1                        *dstore_values.StringValue  `protobuf:"bytes,30004,opt,name=value1" json:"value1,omitempty"`
	TreeNodeId                    *dstore_values.IntegerValue `protobuf:"bytes,30006,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                        *dstore_values.IntegerValue `protobuf:"bytes,30007,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
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

func (m *Response_Row) GetItemConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionId
	}
	return nil
}

func (m *Response_Row) GetBundlePricingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BundlePricingTypeId
	}
	return nil
}

func (m *Response_Row) GetBundlePricingType() *dstore_values.StringValue {
	if m != nil {
		return m.BundlePricingType
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

func (m *Response_Row) GetBundlePriceOrDiscount() *dstore_values.DecimalValue {
	if m != nil {
		return m.BundlePriceOrDiscount
	}
	return nil
}

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

func (m *Response_Row) GetCombineGroupsWithANDOperator() *dstore_values.BooleanValue {
	if m != nil {
		return m.CombineGroupsWithANDOperator
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

func (m *Response_Row) GetItemConditionGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ItemConditionGroupDescription
	}
	return nil
}

func (m *Response_Row) GetCombinePartsWithANDOperator() *dstore_values.BooleanValue {
	if m != nil {
		return m.CombinePartsWithANDOperator
	}
	return nil
}

func (m *Response_Row) GetItemPartSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemPartSortNo
	}
	return nil
}

func (m *Response_Row) GetInheritDepth() *dstore_values.IntegerValue {
	if m != nil {
		return m.InheritDepth
	}
	return nil
}

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignBundlePricing_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignBundlePricing_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignBundlePricing_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 1321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x58, 0x69, 0x6f, 0x1b, 0x45,
	0x18, 0x56, 0x09, 0x71, 0x92, 0xb7, 0x87, 0x9b, 0x75, 0x12, 0xa6, 0x4e, 0x1b, 0xda, 0x44, 0x08,
	0x3e, 0x20, 0x87, 0x38, 0x1c, 0x25, 0x02, 0x21, 0x1c, 0x97, 0x10, 0xda, 0x38, 0xa9, 0x13, 0x8a,
	0xe0, 0x03, 0xab, 0xf5, 0xee, 0xc4, 0x59, 0xc9, 0xde, 0x31, 0x3b, 0xeb, 0x44, 0xfe, 0x17, 0xdc,
	0x97, 0x2a, 0x24, 0x4e, 0x51, 0x71, 0x08, 0x5a, 0x04, 0x1f, 0xf8, 0x09, 0xfc, 0x12, 0xc2, 0xf9,
	0x13, 0x78, 0x67, 0x66, 0xd7, 0xde, 0xb5, 0x9d, 0xec, 0x84, 0x2f, 0x8d, 0xd7, 0xfb, 0x3e, 0xcf,
	0x3c, 0xf3, 0x1e, 0xcf, 0x4c, 0x0d, 0xcf, 0x3a, 0x3c, 0x60, 0x3e, 0x5d, 0xa4, 0x5e, 0xdd, 0xf5,
	0xe8, 0x62, 0xcb, 0x67, 0x36, 0x75, 0xda, 0x3e, 0xe5, 0x8b, 0xac, 0x69, 0xae, 0xd1, 0x60, 0xd5,
	0x6a, 0xb6, 0x2c, 0xb7, 0xee, 0x95, 0xda, 0x9e, 0xd3, 0xa0, 0x5b, 0xbe, 0x6b, 0xbb, 0x5e, 0xdd,
	0xdc, 0x6a, 0x17, 0x30, 0x32, 0x60, 0xc6, 0xa3, 0x0a, 0x5e, 0x50, 0xf0, 0xc2, 0xf1, 0x98, 0x7c,
	0x2e, 0x5c, 0x6c, 0xdf, 0x6a, 0xb4, 0x29, 0x57, 0x14, 0xf9, 0x0b, 0x49, 0x05, 0xd4, 0xf7, 0x99,
	0x1f, 0xbe, 0x9a, 0x4d, 0xbe, 0x6a, 0x52, 0xce, 0xad, 0x3a, 0x0d, 0x5f, 0x2e, 0xf4, 0xbf, 0x0c,
	0x2c, 0xd7, 0xdb, 0x65, 0x7e, 0xd3, 0x0a, 0x5c, 0xe6, 0xa9, 0xa0, 0xf9, 0x4f, 0x26, 0x00, 0xb6,
	0x2c, 0xdf, 0xc2, 0xb7, 0xd4, 0xe7, 0xc6, 0x0a, 0x40, 0x8d, 0x7a, 0x74, 0xd7, 0x0d, 0x4c, 0xd7,
	0x21, 0xa7, 0x2e, 0x9f, 0x7a, 0xe4, 0x74, 0x71, 0xb6, 0x10, 0xee, 0x21, 0x54, 0xe5, 0x7a, 0x01,
	0xad, 0x53, 0xff, 0x96, 0x78, 0xaa, 0x4e, 0x84, 0xe1, 0xeb, 0x8e, 0xf1, 0x30, 0x64, 0x7b, 0x58,
	0xd3, 0x6b, 0x37, 0x1a, 0xe4, 0xf7, 0x31, 0x64, 0x18, 0xaf, 0x9e, 0xed, 0x06, 0x55, 0xf0, 0x5b,
	0xa3, 0x0c, 0x59, 0xe6, 0x35, 0x3a, 0xa6, 0x83, 0xdf, 0x79, 0xae, 0x10, 0x43, 0xee, 0x4b, 0x5f,
	0xe9, 0x9c, 0xc0, 0x94, 0xbb, 0x10, 0x63, 0x09, 0xa6, 0xfa, 0x58, 0xd4, 0x9a, 0x87, 0x6a, 0x4d,
	0x23, 0x19, 0x2e, 0x17, 0xde, 0x80, 0xe9, 0x5d, 0xb7, 0x81, 0x1b, 0x35, 0x6b, 0x1d, 0xd3, 0x0d,
	0x68, 0xd3, 0xe4, 0x54, 0x6e, 0x74, 0x24, 0x7d, 0xf9, 0x49, 0x85, 0x2c, 0x75, 0xd6, 0x11, 0xb7,
	0x4d, 0xc5, 0x86, 0x57, 0x20, 0x3f, 0x94, 0x4e, 0xe9, 0xf8, 0x43, 0xe9, 0x98, 0x1e, 0xc0, 0x49,
	0x29, 0xaf, 0xc3, 0x1c, 0x67, 0x7e, 0x20, 0x90, 0xf6, 0x1e, 0xe6, 0xdf, 0xc6, 0x18, 0x97, 0x07,
	0xae, 0x2d, 0xf0, 0x0d, 0xfc, 0x44, 0xee, 0x97, 0x9a, 0xf2, 0x7d, 0x9a, 0x78, 0xe0, 0x63, 0xaf,
	0x28, 0x49, 0x79, 0xc1, 0x50, 0xea, 0xac, 0x26, 0xf0, 0xeb, 0xce, 0x0d, 0xfc, 0x6b, 0xdc, 0x80,
	0x85, 0xe3, 0xf9, 0x95, 0xc8, 0x3f, 0x95, 0xc8, 0xb9, 0xa3, 0x99, 0xc2, 0x8a, 0x9d, 0x97, 0x6c,
	0xac, 0x25, 0xf3, 0x2c, 0xf5, 0x8d, 0xa6, 0xea, 0x3b, 0x27, 0x30, 0x9b, 0x12, 0x22, 0x35, 0x15,
	0x61, 0xba, 0x9f, 0x45, 0xa9, 0xf8, 0x2b, 0x2c, 0x59, 0x32, 0x5e, 0xae, 0xfc, 0x32, 0x5c, 0x70,
	0xbd, 0x3d, 0x54, 0x14, 0x60, 0xa1, 0x5b, 0xc1, 0x5e, 0x42, 0x42, 0x26, 0x55, 0xc2, 0x4c, 0x08,
	0x2e, 0x0b, 0x6c, 0x4c, 0xca, 0x2a, 0xcc, 0x1d, 0x49, 0xab, 0x34, 0xfd, 0xad, 0x34, 0xe5, 0x87,
	0x13, 0x48, 0x6d, 0x36, 0x5c, 0xf6, 0xa9, 0xdd, 0xf6, 0xb9, 0xbb, 0x4f, 0x4d, 0x2a, 0x56, 0x97,
	0x93, 0x95, 0x90, 0x38, 0x96, 0x2a, 0xf1, 0x52, 0x97, 0xe3, 0x5a, 0x97, 0x22, 0xa6, 0xf4, 0x26,
	0x3c, 0x94, 0xb6, 0x88, 0x12, 0xfc, 0x8f, 0x12, 0x7c, 0xe5, 0x58, 0xba, 0xa8, 0xf7, 0xea, 0xd8,
	0xa8, 0x4a, 0x8b, 0x89, 0x8e, 0x60, 0x26, 0x5b, 0x85, 0x93, 0xf1, 0xa1, 0xf3, 0x50, 0x63, 0xac,
	0x41, 0x2d, 0x4f, 0xc9, 0xbe, 0x80, 0x14, 0xf2, 0x13, 0x7f, 0x81, 0xf9, 0xdb, 0xb1, 0xf6, 0xe1,
	0xc6, 0x75, 0x58, 0x38, 0x9e, 0x5f, 0x09, 0xfe, 0x57, 0x09, 0xbe, 0x74, 0x24, 0x91, 0x10, 0x3b,
	0xff, 0xdb, 0x34, 0x8c, 0x57, 0x29, 0x6f, 0x31, 0x8f, 0x53, 0xe3, 0x31, 0x18, 0x95, 0xf6, 0x17,
	0x3a, 0x53, 0x37, 0xad, 0xa1, 0xbb, 0x2a, 0x6b, 0xbc, 0x26, 0xfe, 0xad, 0xaa, 0x40, 0xe3, 0x55,
	0x38, 0x2f, 0x8c, 0xcf, 0x8c, 0x39, 0x1f, 0x9a, 0xcd, 0x08, 0x82, 0x0b, 0x7d, 0xe0, 0x7e, 0x7f,
	0xdc, 0xc0, 0xe7, 0xf5, 0xde, 0x73, 0x35, 0xdb, 0x4c, 0x7e, 0x61, 0x5c, 0x85, 0xb1, 0xd0, 0x70,
	0xd1, 0x3f, 0x04, 0xe3, 0xdc, 0x00, 0xa3, 0xb2, 0xe3, 0x0d, 0xf5, 0xb7, 0x1a, 0x85, 0xe3, 0x70,
	0x8e, 0xf8, 0xec, 0x00, 0x27, 0x5c, 0xa0, 0x56, 0x0a, 0x27, 0x39, 0x22, 0x0a, 0x51, 0x2e, 0x0a,
	0x55, 0x76, 0x50, 0x15, 0x34, 0xf9, 0x3b, 0x39, 0x18, 0xc1, 0x07, 0x63, 0x06, 0x32, 0xf8, 0x28,
	0xec, 0xec, 0xcd, 0x0a, 0xa6, 0x67, 0xb4, 0x3a, 0x8a, 0x8f, 0x68, 0x53, 0x2f, 0xc2, 0xa4, 0x87,
	0xe5, 0xa8, 0x59, 0x9c, 0x3a, 0x66, 0x4b, 0x31, 0x91, 0xb7, 0x2a, 0xe9, 0x25, 0xce, 0x22, 0xac,
	0x24, 0x50, 0xe1, 0xf2, 0x82, 0x49, 0xda, 0x9c, 0xcd, 0x3c, 0x47, 0x39, 0x2e, 0x2e, 0xf6, 0x76,
	0x25, 0xdd, 0x3c, 0xb3, 0x02, 0xb6, 0x1a, 0xa1, 0x50, 0xd3, 0x4d, 0x98, 0xa9, 0xc9, 0x9d, 0x45,
	0x82, 0xcc, 0xa0, 0xd3, 0xa2, 0x82, 0xee, 0x1d, 0x0d, 0xba, 0x5c, 0x2d, 0x9e, 0x95, 0x1d, 0x44,
	0x22, 0xe5, 0x75, 0xc8, 0x0d, 0xa1, 0x24, 0xef, 0x56, 0x52, 0x27, 0x70, 0x72, 0x80, 0xce, 0x78,
	0x06, 0x4e, 0xc7, 0xcf, 0x87, 0xf7, 0x34, 0x44, 0x4d, 0xb8, 0xdd, 0x83, 0xe1, 0x2a, 0x8c, 0xbf,
	0xd1, 0xb6, 0xbc, 0xc0, 0x0d, 0x3a, 0xe4, 0x7d, 0x0d, 0x68, 0x37, 0x1a, 0x3b, 0x23, 0xe7, 0x08,
	0xf3, 0xf5, 0xec, 0x40, 0x9e, 0x28, 0xdc, 0x14, 0xc7, 0x18, 0xf9, 0x40, 0xa3, 0x5a, 0x93, 0x11,
	0x50, 0x1c, 0x34, 0x7c, 0x13, 0x61, 0xd8, 0xfc, 0xf9, 0xbe, 0x7a, 0x39, 0x94, 0xdb, 0xbe, 0x2b,
	0xed, 0x80, 0x7c, 0x98, 0x9e, 0x19, 0x92, 0xa8, 0x5b, 0xb9, 0x07, 0x36, 0x9e, 0x80, 0x31, 0x39,
	0xd4, 0x1e, 0x23, 0x1f, 0x69, 0xec, 0x30, 0x23, 0x82, 0x2b, 0x0c, 0xed, 0x9c, 0xc4, 0x8a, 0x44,
	0x4d, 0xf4, 0x06, 0x94, 0x6d, 0xb3, 0xb6, 0x17, 0x90, 0x8f, 0x87, 0xf3, 0x38, 0xd4, 0x76, 0x9b,
	0x56, 0x43, 0xf1, 0x4c, 0xf7, 0x4a, 0x45, 0x37, 0xfd, 0x72, 0x08, 0xc5, 0x72, 0x41, 0x77, 0x8f,
	0x4b, 0xe4, 0xd3, 0xdb, 0xa7, 0x52, 0x77, 0x16, 0x8b, 0x4f, 0xa0, 0x8b, 0xe4, 0xb3, 0x13, 0xa1,
	0x8b, 0xc6, 0x16, 0x4c, 0x0d, 0x33, 0x68, 0xf2, 0xf9, 0x6d, 0x8d, 0xeb, 0x53, 0x6e, 0x88, 0x5b,
	0xe3, 0x35, 0x65, 0xca, 0x61, 0x4d, 0x34, 0x23, 0x33, 0xf0, 0x29, 0xc5, 0x14, 0x3b, 0x62, 0x32,
	0x38, 0xf9, 0x42, 0x43, 0xd9, 0xa4, 0x42, 0xee, 0x20, 0xb0, 0x82, 0xb8, 0x75, 0x87, 0x1b, 0xdb,
	0x30, 0x23, 0x29, 0x06, 0xee, 0x01, 0xe4, 0x4b, 0x1d, 0x89, 0x53, 0x02, 0xdc, 0x7f, 0x33, 0x30,
	0xaa, 0x30, 0xd3, 0xd7, 0x5a, 0x2d, 0xcb, 0x97, 0xb3, 0xf2, 0x95, 0xd6, 0xbe, 0x13, 0x8d, 0x85,
	0x17, 0x50, 0x31, 0x36, 0x75, 0x98, 0xb7, 0x59, 0xb3, 0x86, 0x26, 0x68, 0xd6, 0x7d, 0xd6, 0x6e,
	0x71, 0xf3, 0xc0, 0xc5, 0xa3, 0xd9, 0x32, 0xb1, 0x6f, 0xf1, 0xbc, 0xa3, 0xbe, 0x85, 0x7c, 0xe4,
	0xeb, 0x23, 0xf8, 0x13, 0xc3, 0x70, 0x31, 0x24, 0x5a, 0x93, 0x3c, 0xaf, 0x20, 0xcd, 0xf3, 0x95,
	0xf2, 0x66, 0x48, 0x81, 0x07, 0xf7, 0x83, 0xc3, 0xc4, 0xc7, 0x87, 0xe3, 0x8e, 0x46, 0xae, 0x2f,
	0x0e, 0x6c, 0x22, 0x3e, 0x21, 0x4f, 0xc3, 0x44, 0x83, 0xee, 0xd3, 0x86, 0x2c, 0xdd, 0x37, 0x1a,
	0x74, 0xe3, 0x32, 0x5c, 0x54, 0x6c, 0x07, 0x1e, 0xe8, 0xd3, 0x27, 0xf3, 0x21, 0xb2, 0xfb, 0xad,
	0x56, 0xc9, 0x12, 0xc2, 0x64, 0x0e, 0x30, 0xbd, 0x2f, 0x81, 0x21, 0x59, 0x15, 0x57, 0x34, 0xbd,
	0xdf, 0xe9, 0x10, 0x4a, 0xff, 0x96, 0x3c, 0xdb, 0x6a, 0x8e, 0x57, 0x60, 0x22, 0x2a, 0xc8, 0x12,
	0xf9, 0x5e, 0x63, 0x73, 0xbd, 0xf0, 0x38, 0xb6, 0x48, 0x7e, 0x38, 0x09, 0xb6, 0x68, 0x50, 0xb8,
	0x3c, 0x34, 0x33, 0xf1, 0xd2, 0xfd, 0xa8, 0x41, 0x79, 0x69, 0x30, 0x43, 0xf1, 0xda, 0x51, 0xb8,
	0x12, 0x75, 0xa2, 0xe8, 0x8c, 0xe1, 0x8d, 0x78, 0x57, 0xa7, 0x11, 0x67, 0x43, 0x1e, 0xd1, 0x1d,
	0x03, 0x7d, 0x18, 0x9d, 0xa7, 0xb2, 0xfb, 0xa2, 0x82, 0xdc, 0xd3, 0x29, 0xc8, 0x39, 0x81, 0x13,
	0x9c, 0x61, 0x3d, 0x4a, 0x70, 0x36, 0x71, 0x9f, 0x25, 0xbf, 0xea, 0xb0, 0x9c, 0x89, 0x5f, 0x6e,
	0x8d, 0x35, 0x38, 0x2f, 0x7d, 0x22, 0x91, 0xcb, 0xc3, 0xf4, 0x5c, 0x66, 0x05, 0x2a, 0x9e, 0xbd,
	0xc7, 0x21, 0x23, 0xc3, 0x8a, 0xe4, 0xae, 0x06, 0x3c, 0x8c, 0xed, 0xa2, 0x96, 0xc9, 0x3d, 0x6d,
	0xd4, 0x72, 0x17, 0xb5, 0x44, 0x7e, 0xd2, 0x46, 0x2d, 0x19, 0xcf, 0xc1, 0x99, 0xb8, 0xb5, 0x92,
	0x9f, 0x0f, 0x35, 0xb2, 0x05, 0x41, 0xd7, 0x54, 0x8d, 0x27, 0x61, 0x2c, 0xc2, 0xfe, 0xa2, 0x83,
	0xcd, 0x78, 0x12, 0x57, 0xba, 0x05, 0xb3, 0x2e, 0xeb, 0xbb, 0xf0, 0xf5, 0x7e, 0x52, 0x78, 0xed,
	0xa9, 0xff, 0xf9, 0x63, 0x43, 0x2d, 0x23, 0xff, 0x37, 0xbf, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0xd7, 0x30, 0xe7, 0xae, 0x10, 0x00, 0x00,
}
