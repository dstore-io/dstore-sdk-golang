// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetDSS_Index_Trend_Ad.proto
// DO NOT EDIT!

/*
Package st_GetDSS_Index_Trend_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetDSS_Index_Trend_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetDSS_Index_Trend_Ad

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
	DomainTreeNodeId                  *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=domain_tree_node_id,json=domainTreeNodeId" json:"domain_tree_node_id,omitempty"`
	DomainTreeNodeIdNull              bool                        `protobuf:"varint,1001,opt,name=domain_tree_node_id_null,json=domainTreeNodeIdNull" json:"domain_tree_node_id_null,omitempty"`
	Intervalls                        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=intervalls" json:"intervalls,omitempty"`
	IntervallsNull                    bool                        `protobuf:"varint,1002,opt,name=intervalls_null,json=intervallsNull" json:"intervalls_null,omitempty"`
	MinutesPerIntervall               *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=minutes_per_intervall,json=minutesPerIntervall" json:"minutes_per_intervall,omitempty"`
	MinutesPerIntervallNull           bool                        `protobuf:"varint,1003,opt,name=minutes_per_intervall_null,json=minutesPerIntervallNull" json:"minutes_per_intervall_null,omitempty"`
	GroupByNodesOnLevel               *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=group_by_nodes_on_level,json=groupByNodesOnLevel" json:"group_by_nodes_on_level,omitempty"`
	GroupByNodesOnLevelNull           bool                        `protobuf:"varint,1004,opt,name=group_by_nodes_on_level_null,json=groupByNodesOnLevelNull" json:"group_by_nodes_on_level_null,omitempty"`
	IsLevelId                         *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=is_level_id,json=isLevelId" json:"is_level_id,omitempty"`
	IsLevelIdNull                     bool                        `protobuf:"varint,1005,opt,name=is_level_id_null,json=isLevelIdNull" json:"is_level_id_null,omitempty"`
	GroupByNodeCharacteristicId       *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=group_by_node_characteristic_id,json=groupByNodeCharacteristicId" json:"group_by_node_characteristic_id,omitempty"`
	GroupByNodeCharacteristicIdNull   bool                        `protobuf:"varint,1006,opt,name=group_by_node_characteristic_id_null,json=groupByNodeCharacteristicIdNull" json:"group_by_node_characteristic_id_null,omitempty"`
	OnlyValuesInOneId                 *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=only_values_in_one_id,json=onlyValuesInOneId" json:"only_values_in_one_id,omitempty"`
	OnlyValuesInOneIdNull             bool                        `protobuf:"varint,1007,opt,name=only_values_in_one_id_null,json=onlyValuesInOneIdNull" json:"only_values_in_one_id_null,omitempty"`
	IdsInOneIdAreTreeNodeIds          *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=ids_in_one_id_are_tree_node_ids,json=idsInOneIdAreTreeNodeIds" json:"ids_in_one_id_are_tree_node_ids,omitempty"`
	IdsInOneIdAreTreeNodeIdsNull      bool                        `protobuf:"varint,1008,opt,name=ids_in_one_id_are_tree_node_ids_null,json=idsInOneIdAreTreeNodeIdsNull" json:"ids_in_one_id_are_tree_node_ids_null,omitempty"`
	WeightViews                       *dstore_values.DecimalValue `protobuf:"bytes,9,opt,name=weight_views,json=weightViews" json:"weight_views,omitempty"`
	WeightViewsNull                   bool                        `protobuf:"varint,1009,opt,name=weight_views_null,json=weightViewsNull" json:"weight_views_null,omitempty"`
	WeightBuyingInterest              *dstore_values.DecimalValue `protobuf:"bytes,10,opt,name=weight_buying_interest,json=weightBuyingInterest" json:"weight_buying_interest,omitempty"`
	WeightBuyingInterestNull          bool                        `protobuf:"varint,1010,opt,name=weight_buying_interest_null,json=weightBuyingInterestNull" json:"weight_buying_interest_null,omitempty"`
	WeightOrder                       *dstore_values.DecimalValue `protobuf:"bytes,11,opt,name=weight_order,json=weightOrder" json:"weight_order,omitempty"`
	WeightOrderNull                   bool                        `protobuf:"varint,1011,opt,name=weight_order_null,json=weightOrderNull" json:"weight_order_null,omitempty"`
	IncludeDeactivatedNodes           *dstore_values.BooleanValue `protobuf:"bytes,12,opt,name=include_deactivated_nodes,json=includeDeactivatedNodes" json:"include_deactivated_nodes,omitempty"`
	IncludeDeactivatedNodesNull       bool                        `protobuf:"varint,1012,opt,name=include_deactivated_nodes_null,json=includeDeactivatedNodesNull" json:"include_deactivated_nodes_null,omitempty"`
	IncludeNodesWithoutTreeNodeId     *dstore_values.BooleanValue `protobuf:"bytes,13,opt,name=include_nodes_without_tree_node_id,json=includeNodesWithoutTreeNodeId" json:"include_nodes_without_tree_node_id,omitempty"`
	IncludeNodesWithoutTreeNodeIdNull bool                        `protobuf:"varint,1013,opt,name=include_nodes_without_tree_node_id_null,json=includeNodesWithoutTreeNodeIdNull" json:"include_nodes_without_tree_node_id_null,omitempty"`
	OrderResultByIntervall            *dstore_values.IntegerValue `protobuf:"bytes,14,opt,name=order_result_by_intervall,json=orderResultByIntervall" json:"order_result_by_intervall,omitempty"`
	OrderResultByIntervallNull        bool                        `protobuf:"varint,1014,opt,name=order_result_by_intervall_null,json=orderResultByIntervallNull" json:"order_result_by_intervall_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetDomainTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DomainTreeNodeId
	}
	return nil
}

func (m *Parameters) GetIntervalls() *dstore_values.IntegerValue {
	if m != nil {
		return m.Intervalls
	}
	return nil
}

func (m *Parameters) GetMinutesPerIntervall() *dstore_values.IntegerValue {
	if m != nil {
		return m.MinutesPerIntervall
	}
	return nil
}

func (m *Parameters) GetGroupByNodesOnLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupByNodesOnLevel
	}
	return nil
}

func (m *Parameters) GetIsLevelId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsLevelId
	}
	return nil
}

func (m *Parameters) GetGroupByNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupByNodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetOnlyValuesInOneId() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyValuesInOneId
	}
	return nil
}

func (m *Parameters) GetIdsInOneIdAreTreeNodeIds() *dstore_values.BooleanValue {
	if m != nil {
		return m.IdsInOneIdAreTreeNodeIds
	}
	return nil
}

func (m *Parameters) GetWeightViews() *dstore_values.DecimalValue {
	if m != nil {
		return m.WeightViews
	}
	return nil
}

func (m *Parameters) GetWeightBuyingInterest() *dstore_values.DecimalValue {
	if m != nil {
		return m.WeightBuyingInterest
	}
	return nil
}

func (m *Parameters) GetWeightOrder() *dstore_values.DecimalValue {
	if m != nil {
		return m.WeightOrder
	}
	return nil
}

func (m *Parameters) GetIncludeDeactivatedNodes() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeDeactivatedNodes
	}
	return nil
}

func (m *Parameters) GetIncludeNodesWithoutTreeNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeNodesWithoutTreeNodeId
	}
	return nil
}

func (m *Parameters) GetOrderResultByIntervall() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderResultByIntervall
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
	RowId              int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	IntervallStart     *dstore_values.TimestampValue `protobuf:"bytes,10001,opt,name=intervall_start,json=intervallStart" json:"intervall_start,omitempty"`
	Orders             *dstore_values.DecimalValue   `protobuf:"bytes,10002,opt,name=orders" json:"orders,omitempty"`
	VBOIndex           *dstore_values.DecimalValue   `protobuf:"bytes,10003,opt,name=v_b_o_index,json=vBOIndex" json:"v_b_o_index,omitempty"`
	TrendOfId          *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=trend_of_id,json=trendOfId" json:"trend_of_id,omitempty"`
	Views              *dstore_values.DecimalValue   `protobuf:"bytes,10005,opt,name=views" json:"views,omitempty"`
	TrendOfDescription *dstore_values.StringValue    `protobuf:"bytes,10006,opt,name=trend_of_description,json=trendOfDescription" json:"trend_of_description,omitempty"`
	Intervall          *dstore_values.IntegerValue   `protobuf:"bytes,10007,opt,name=intervall" json:"intervall,omitempty"`
	BuyingInterests    *dstore_values.DecimalValue   `protobuf:"bytes,10008,opt,name=buying_interests,json=buyingInterests" json:"buying_interests,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetIntervallStart() *dstore_values.TimestampValue {
	if m != nil {
		return m.IntervallStart
	}
	return nil
}

func (m *Response_Row) GetOrders() *dstore_values.DecimalValue {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *Response_Row) GetVBOIndex() *dstore_values.DecimalValue {
	if m != nil {
		return m.VBOIndex
	}
	return nil
}

func (m *Response_Row) GetTrendOfId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TrendOfId
	}
	return nil
}

func (m *Response_Row) GetViews() *dstore_values.DecimalValue {
	if m != nil {
		return m.Views
	}
	return nil
}

func (m *Response_Row) GetTrendOfDescription() *dstore_values.StringValue {
	if m != nil {
		return m.TrendOfDescription
	}
	return nil
}

func (m *Response_Row) GetIntervall() *dstore_values.IntegerValue {
	if m != nil {
		return m.Intervall
	}
	return nil
}

func (m *Response_Row) GetBuyingInterests() *dstore_values.DecimalValue {
	if m != nil {
		return m.BuyingInterests
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetDSS_Index_Trend_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetDSS_Index_Trend_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetDSS_Index_Trend_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetDSS_Index_Trend_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1079 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x57, 0xeb, 0x6e, 0x1b, 0x45,
	0x14, 0x56, 0x49, 0x73, 0x9b, 0xb4, 0x4d, 0xba, 0xb9, 0x6d, 0xec, 0x36, 0x81, 0x50, 0x41, 0x24,
	0x24, 0x07, 0x95, 0x4a, 0x45, 0x05, 0x55, 0x6a, 0x1a, 0xa8, 0x0c, 0xc4, 0x2e, 0x0e, 0x4a, 0x55,
	0x7e, 0x30, 0x5a, 0x7b, 0x27, 0xce, 0x88, 0xf5, 0x8e, 0x35, 0x33, 0xb6, 0xf1, 0x5b, 0x70, 0x87,
	0x7f, 0x3c, 0x08, 0x6f, 0xc1, 0x5b, 0x70, 0xbf, 0xf3, 0x9b, 0x33, 0x73, 0xc6, 0xd9, 0x5d, 0x77,
	0xdd, 0x75, 0xff, 0x24, 0x99, 0x9d, 0xf3, 0x5d, 0xce, 0xdc, 0xce, 0x09, 0xb9, 0x1d, 0x2a, 0x2d,
	0x24, 0xdb, 0x67, 0x71, 0x9b, 0xc7, 0x6c, 0xbf, 0x2b, 0x45, 0x8b, 0x85, 0x3d, 0xc9, 0xd4, 0xbe,
	0xd2, 0xf4, 0x01, 0xd3, 0x87, 0xc7, 0xc7, 0xb4, 0x1a, 0x87, 0xec, 0x13, 0xfa, 0x81, 0x64, 0x71,
	0x48, 0xef, 0x85, 0x15, 0x88, 0xd1, 0xc2, 0x7b, 0x09, 0x81, 0x15, 0x04, 0x56, 0x26, 0x45, 0x97,
	0x56, 0x9d, 0x40, 0x3f, 0x88, 0x7a, 0x4c, 0x21, 0xb8, 0xb4, 0x95, 0x55, 0x65, 0x52, 0x0a, 0xe9,
	0xa6, 0xca, 0xd9, 0xa9, 0x0e, 0x53, 0x2a, 0x68, 0x33, 0x37, 0xf9, 0xe2, 0xf8, 0xa4, 0x0e, 0x78,
	0x7c, 0x2a, 0x64, 0x27, 0xd0, 0x5c, 0xc4, 0x18, 0xb4, 0xfb, 0xc3, 0x32, 0x21, 0x0f, 0x03, 0x19,
	0xc0, 0x2c, 0x93, 0xca, 0x7b, 0x87, 0xac, 0x86, 0xa2, 0x03, 0x71, 0x54, 0x4b, 0xc6, 0x68, 0x2c,
	0x42, 0x46, 0x79, 0xe8, 0x5f, 0x78, 0xfe, 0xc2, 0xde, 0xd2, 0xcd, 0x72, 0xc5, 0xa5, 0xe1, 0xec,
	0xf1, 0x58, 0xb3, 0x36, 0x93, 0x27, 0x66, 0xd4, 0x58, 0x41, 0x1c, 0xa4, 0xc2, 0x6a, 0x80, 0xaa,
	0x86, 0xde, 0x6d, 0xe2, 0xe7, 0x70, 0xd1, 0xb8, 0x17, 0x45, 0xfe, 0x8f, 0xf3, 0xc0, 0xb8, 0xd0,
	0x58, 0x1b, 0x07, 0xd5, 0x60, 0xd2, 0x7b, 0x83, 0x10, 0x43, 0x2d, 0x41, 0x26, 0x52, 0xfe, 0x73,
	0xc5, 0xda, 0xa9, 0x70, 0x6f, 0x8f, 0x2c, 0x27, 0x23, 0x14, 0xfb, 0x09, 0xc5, 0xae, 0x24, 0xdf,
	0xad, 0x4c, 0x9d, 0xac, 0x77, 0x78, 0xdc, 0xd3, 0x4c, 0xd1, 0x2e, 0x93, 0xf4, 0x7c, 0xd6, 0x9f,
	0x29, 0x56, 0x5c, 0x75, 0xc8, 0x87, 0x4c, 0x56, 0x47, 0x38, 0xef, 0x4d, 0x52, 0xca, 0x25, 0x44,
	0x17, 0x3f, 0xa3, 0x8b, 0xcd, 0x1c, 0xa4, 0xb5, 0xf3, 0x3e, 0xd9, 0x6c, 0x4b, 0xd1, 0xeb, 0xd2,
	0xe6, 0xd0, 0xae, 0x95, 0xa2, 0x22, 0xa6, 0x11, 0xeb, 0xb3, 0xc8, 0xbf, 0x38, 0x85, 0x21, 0x8b,
	0x3d, 0x18, 0x9a, 0x65, 0x54, 0xf5, 0xf8, 0x3d, 0x83, 0xf3, 0xee, 0x92, 0x6b, 0x13, 0x28, 0xd1,
	0xd2, 0x2f, 0xce, 0x52, 0x0e, 0xd6, 0x6d, 0xc4, 0x12, 0x57, 0x0e, 0x00, 0xa7, 0x60, 0x36, 0xd7,
	0x46, 0x53, 0x88, 0x88, 0x05, 0x31, 0xda, 0x58, 0xe4, 0xca, 0xc2, 0x61, 0xfb, 0xf7, 0xc8, 0x4a,
	0x0a, 0x8c, 0x82, 0xbf, 0xa2, 0xe0, 0xe5, 0xf3, 0x28, 0x2b, 0x13, 0x90, 0x9d, 0x8c, 0x4d, 0xda,
	0x3a, 0x83, 0x13, 0xd9, 0x82, 0xc5, 0xe1, 0x4a, 0xf3, 0x96, 0x91, 0x9e, 0x2b, 0x5e, 0x81, 0x72,
	0x2a, 0x8b, 0xfb, 0x19, 0x02, 0x30, 0x53, 0x23, 0x37, 0x0a, 0x24, 0xd0, 0xe0, 0x6f, 0x68, 0x70,
	0xe7, 0x29, 0x5c, 0xd6, 0xf2, 0x11, 0x59, 0x17, 0x71, 0x34, 0xa4, 0x68, 0x04, 0xb6, 0x1a, 0x56,
	0xd6, 0xde, 0x94, 0xf9, 0xe2, 0x35, 0xba, 0x6a, 0x90, 0xf6, 0x4f, 0x55, 0x8d, 0xeb, 0xb1, 0xb9,
	0x2a, 0x77, 0x48, 0x29, 0x97, 0x0e, 0x4d, 0xfd, 0x8e, 0xa6, 0xd6, 0x9f, 0xc0, 0x59, 0x2b, 0x1f,
	0x91, 0x1d, 0x1e, 0xa6, 0x31, 0x81, 0x64, 0x99, 0x1b, 0xa7, 0xfc, 0x85, 0x62, 0x53, 0x3e, 0xc4,
	0x39, 0xd6, 0x7b, 0x92, 0x25, 0x17, 0x52, 0x79, 0xef, 0x92, 0x1b, 0x05, 0xfc, 0xe8, 0xf2, 0x0f,
	0x74, 0x79, 0x6d, 0x12, 0x91, 0x35, 0x7b, 0x97, 0x5c, 0x1a, 0x30, 0xde, 0x3e, 0xd3, 0xb4, 0xcf,
	0xd9, 0x40, 0xf9, 0x8b, 0xb9, 0xce, 0x42, 0xd6, 0xe2, 0x9d, 0x20, 0x42, 0x67, 0x4b, 0x08, 0x38,
	0x31, 0xf1, 0xde, 0x2b, 0xe4, 0x6a, 0x1a, 0x8f, 0xca, 0x7f, 0xa2, 0xf2, 0x72, 0x2a, 0xd0, 0xdd,
	0xa8, 0x0d, 0x17, 0xdc, 0xec, 0x0d, 0x79, 0xdc, 0xc6, 0x1b, 0xc9, 0x94, 0xf6, 0x49, 0xb1, 0xec,
	0x1a, 0x42, 0x0f, 0x2c, 0xb2, 0xea, 0x80, 0xe0, 0xbf, 0x9c, 0x4f, 0x89, 0x4e, 0xfe, 0x42, 0x27,
	0x7e, 0x1e, 0x76, 0x2c, 0x7f, 0x21, 0x43, 0x26, 0xfd, 0xa5, 0xa9, 0xf3, 0xaf, 0x9b, 0xf8, 0x54,
	0xfe, 0x16, 0x8f, 0xaa, 0x7f, 0x67, 0xf2, 0xb7, 0x81, 0x56, 0xec, 0x11, 0xd9, 0xe2, 0x71, 0x2b,
	0xea, 0xc1, 0x2e, 0x85, 0x0c, 0x8e, 0x30, 0xef, 0x07, 0x9a, 0x85, 0xf8, 0x12, 0xf8, 0x97, 0x8a,
	0xcf, 0xc4, 0xa6, 0x43, 0x1f, 0x26, 0x60, 0xfb, 0x44, 0x78, 0x87, 0x64, 0x7b, 0x22, 0x31, 0x5a,
	0xfa, 0x07, 0x2d, 0x95, 0x27, 0x30, 0x58, 0x7b, 0xa7, 0x64, 0x77, 0xc4, 0x82, 0xc8, 0x01, 0xd7,
	0x67, 0xa2, 0xa7, 0xb3, 0xa5, 0xe7, 0x72, 0xb1, 0xcf, 0xeb, 0x8e, 0xc6, 0x52, 0x3f, 0x42, 0x92,
	0x54, 0x1d, 0x3a, 0x26, 0x2f, 0x17, 0xeb, 0xa0, 0xed, 0x7f, 0xd1, 0xf6, 0x0b, 0x4f, 0x25, 0xb4,
	0xe6, 0x4f, 0xc8, 0x16, 0xee, 0x00, 0x6c, 0x6d, 0x2f, 0xd2, 0xe6, 0x5d, 0x49, 0x0a, 0xc8, 0x95,
	0xe2, 0xd7, 0x6a, 0xc3, 0xa2, 0x1b, 0x16, 0x7c, 0x30, 0x4c, 0x6a, 0xc8, 0x7d, 0xb2, 0x3d, 0x91,
	0x17, 0x3d, 0xfe, 0x87, 0x1e, 0x4b, 0xf9, 0x04, 0xc6, 0xdc, 0xee, 0xf7, 0x73, 0x64, 0x01, 0x66,
	0xba, 0x22, 0x56, 0xcc, 0x7b, 0x95, 0xcc, 0xda, 0x96, 0xc1, 0x15, 0xf1, 0x52, 0x25, 0xdb, 0x8b,
	0x60, 0x3b, 0xf1, 0x96, 0xf9, 0xd9, 0xc0, 0x40, 0xef, 0x31, 0x59, 0x31, 0xcd, 0x02, 0x4d, 0x75,
	0x0b, 0x50, 0x85, 0x67, 0x00, 0x5c, 0x19, 0x03, 0x8f, 0xf7, 0x14, 0x47, 0x30, 0xae, 0x26, 0xe3,
	0xc6, 0x72, 0x27, 0xfb, 0xc1, 0x7b, 0x9d, 0xcc, 0xbb, 0x26, 0x05, 0xaa, 0xac, 0x61, 0xdc, 0x7e,
	0x82, 0x11, 0x5b, 0x98, 0x23, 0xfc, 0xdd, 0x18, 0x85, 0x7b, 0x6f, 0x93, 0x19, 0x29, 0x06, 0x50,
	0x0a, 0x0d, 0xea, 0x56, 0x65, 0xba, 0x86, 0xaa, 0x32, 0x5a, 0x85, 0x4a, 0x43, 0x0c, 0x1a, 0x86,
	0xa0, 0xf4, 0xdd, 0x45, 0x32, 0x03, 0x03, 0x6f, 0x83, 0xcc, 0xc1, 0xd0, 0x9c, 0xb0, 0x4f, 0x6b,
	0xb0, 0x30, 0xb3, 0x8d, 0x59, 0x18, 0xc2, 0x69, 0x79, 0x90, 0xea, 0x1f, 0xa8, 0xd2, 0x81, 0xd4,
	0xfe, 0x67, 0x35, 0xbb, 0x72, 0xd7, 0xc7, 0xf6, 0x53, 0x73, 0xf0, 0xa6, 0x83, 0x4e, 0x17, 0x77,
	0x34, 0x69, 0x2f, 0x8e, 0x0d, 0xca, 0xbb, 0x45, 0xe6, 0xec, 0x16, 0x29, 0xff, 0xf3, 0x5a, 0xf1,
	0x2d, 0x77, 0xb1, 0xa6, 0xe4, 0xf6, 0x69, 0x93, 0x0a, 0x58, 0x7c, 0x48, 0xc4, 0xff, 0x62, 0x0a,
	0xe8, 0x42, 0xff, 0xa0, 0x6e, 0xd3, 0x86, 0x06, 0x64, 0x49, 0xdb, 0xcc, 0xc5, 0xa9, 0x49, 0xec,
	0xcb, 0x5a, 0xf1, 0x39, 0x5c, 0xb4, 0x80, 0xfa, 0x29, 0x64, 0x7e, 0x93, 0xcc, 0xe2, 0xa3, 0xfc,
	0xd5, 0x14, 0xa2, 0x18, 0x0a, 0x75, 0x70, 0xed, 0x5c, 0x11, 0xee, 0x4a, 0x4b, 0xf2, 0xae, 0x3d,
	0x2e, 0x5f, 0xd7, 0xb2, 0x87, 0xcd, 0x51, 0x28, 0x2d, 0xe1, 0x45, 0x44, 0x06, 0xcf, 0x29, 0x1f,
	0x26, 0x30, 0xa8, 0x83, 0x8b, 0xc9, 0x2d, 0xfa, 0x66, 0x1a, 0xfb, 0xe7, 0xe1, 0xb0, 0x71, 0x2b,
	0x63, 0x6f, 0xb2, 0xf2, 0xbf, 0x9d, 0x22, 0x93, 0xe5, 0x66, 0xe6, 0x99, 0x56, 0x07, 0x8f, 0x49,
	0x99, 0x8b, 0xb1, 0x03, 0x96, 0xb4, 0xfa, 0x1f, 0xde, 0x69, 0x0b, 0x15, 0x7e, 0x3c, 0x9a, 0x0f,
	0x9f, 0xe5, 0xbf, 0x81, 0xe6, 0x9c, 0x6d, 0xba, 0x5f, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x37,
	0x05, 0xc6, 0xbe, 0x49, 0x0c, 0x00, 0x00,
}
