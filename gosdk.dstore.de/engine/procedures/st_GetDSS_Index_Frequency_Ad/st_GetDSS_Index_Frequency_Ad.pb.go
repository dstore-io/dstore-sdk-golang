// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetDSS_Index_Frequency_Ad.proto
// DO NOT EDIT!

/*
Package st_GetDSS_Index_Frequency_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetDSS_Index_Frequency_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetDSS_Index_Frequency_Ad

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
	DomainTreeNodeId                  *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=domain_tree_node_id,json=domainTreeNodeId" json:"domain_tree_node_id,omitempty"`
	DomainTreeNodeIdNull              bool                          `protobuf:"varint,1001,opt,name=domain_tree_node_id_null,json=domainTreeNodeIdNull" json:"domain_tree_node_id_null,omitempty"`
	FromDate                          *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                      bool                          `protobuf:"varint,1002,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                            *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                        bool                          `protobuf:"varint,1003,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	GroupByNodesOnLevel               *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=group_by_nodes_on_level,json=groupByNodesOnLevel" json:"group_by_nodes_on_level,omitempty"`
	GroupByNodesOnLevelNull           bool                          `protobuf:"varint,1004,opt,name=group_by_nodes_on_level_null,json=groupByNodesOnLevelNull" json:"group_by_nodes_on_level_null,omitempty"`
	IsLevelId                         *dstore_values.BooleanValue   `protobuf:"bytes,5,opt,name=is_level_id,json=isLevelId" json:"is_level_id,omitempty"`
	IsLevelIdNull                     bool                          `protobuf:"varint,1005,opt,name=is_level_id_null,json=isLevelIdNull" json:"is_level_id_null,omitempty"`
	GroupByNodeCharacteristicId       *dstore_values.IntegerValue   `protobuf:"bytes,6,opt,name=group_by_node_characteristic_id,json=groupByNodeCharacteristicId" json:"group_by_node_characteristic_id,omitempty"`
	GroupByNodeCharacteristicIdNull   bool                          `protobuf:"varint,1006,opt,name=group_by_node_characteristic_id_null,json=groupByNodeCharacteristicIdNull" json:"group_by_node_characteristic_id_null,omitempty"`
	OnlyValuesInOneId                 *dstore_values.BooleanValue   `protobuf:"bytes,7,opt,name=only_values_in_one_id,json=onlyValuesInOneId" json:"only_values_in_one_id,omitempty"`
	OnlyValuesInOneIdNull             bool                          `protobuf:"varint,1007,opt,name=only_values_in_one_id_null,json=onlyValuesInOneIdNull" json:"only_values_in_one_id_null,omitempty"`
	WeightViews                       *dstore_values.DecimalValue   `protobuf:"bytes,8,opt,name=weight_views,json=weightViews" json:"weight_views,omitempty"`
	WeightViewsNull                   bool                          `protobuf:"varint,1008,opt,name=weight_views_null,json=weightViewsNull" json:"weight_views_null,omitempty"`
	WeightBuyingInterest              *dstore_values.DecimalValue   `protobuf:"bytes,9,opt,name=weight_buying_interest,json=weightBuyingInterest" json:"weight_buying_interest,omitempty"`
	WeightBuyingInterestNull          bool                          `protobuf:"varint,1009,opt,name=weight_buying_interest_null,json=weightBuyingInterestNull" json:"weight_buying_interest_null,omitempty"`
	WeightOrder                       *dstore_values.DecimalValue   `protobuf:"bytes,10,opt,name=weight_order,json=weightOrder" json:"weight_order,omitempty"`
	WeightOrderNull                   bool                          `protobuf:"varint,1010,opt,name=weight_order_null,json=weightOrderNull" json:"weight_order_null,omitempty"`
	IncludeDeactivatedNodes           *dstore_values.BooleanValue   `protobuf:"bytes,11,opt,name=include_deactivated_nodes,json=includeDeactivatedNodes" json:"include_deactivated_nodes,omitempty"`
	IncludeDeactivatedNodesNull       bool                          `protobuf:"varint,1011,opt,name=include_deactivated_nodes_null,json=includeDeactivatedNodesNull" json:"include_deactivated_nodes_null,omitempty"`
	IncludeNodesWithoutTreeNodeId     *dstore_values.BooleanValue   `protobuf:"bytes,12,opt,name=include_nodes_without_tree_node_id,json=includeNodesWithoutTreeNodeId" json:"include_nodes_without_tree_node_id,omitempty"`
	IncludeNodesWithoutTreeNodeIdNull bool                          `protobuf:"varint,1012,opt,name=include_nodes_without_tree_node_id_null,json=includeNodesWithoutTreeNodeIdNull" json:"include_nodes_without_tree_node_id_null,omitempty"`
	OrderResultByColumn               *dstore_values.IntegerValue   `protobuf:"bytes,13,opt,name=order_result_by_column,json=orderResultByColumn" json:"order_result_by_column,omitempty"`
	OrderResultByColumnNull           bool                          `protobuf:"varint,1013,opt,name=order_result_by_column_null,json=orderResultByColumnNull" json:"order_result_by_column_null,omitempty"`
	GetTopX                           *dstore_values.IntegerValue   `protobuf:"bytes,14,opt,name=get_top_x,json=getTopX" json:"get_top_x,omitempty"`
	GetTopXNull                       bool                          `protobuf:"varint,1014,opt,name=get_top_x_null,json=getTopXNull" json:"get_top_x_null,omitempty"`
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

func (m *Parameters) GetOrderResultByColumn() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderResultByColumn
	}
	return nil
}

func (m *Parameters) GetGetTopX() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetTopX
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
	Orders                 *dstore_values.DecimalValue `protobuf:"bytes,10001,opt,name=orders" json:"orders,omitempty"`
	VBOIndex               *dstore_values.DecimalValue `protobuf:"bytes,10002,opt,name=v_b_o_index,json=vBOIndex" json:"v_b_o_index,omitempty"`
	FrequencyOfId          *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=frequency_of_id,json=frequencyOfId" json:"frequency_of_id,omitempty"`
	Views                  *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=views" json:"views,omitempty"`
	FrequencyOfDescription *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=frequency_of_description,json=frequencyOfDescription" json:"frequency_of_description,omitempty"`
	BuyingInterests        *dstore_values.DecimalValue `protobuf:"bytes,10006,opt,name=buying_interests,json=buyingInterests" json:"buying_interests,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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

func (m *Response_Row) GetFrequencyOfId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FrequencyOfId
	}
	return nil
}

func (m *Response_Row) GetViews() *dstore_values.DecimalValue {
	if m != nil {
		return m.Views
	}
	return nil
}

func (m *Response_Row) GetFrequencyOfDescription() *dstore_values.StringValue {
	if m != nil {
		return m.FrequencyOfDescription
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetDSS_Index_Frequency_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetDSS_Index_Frequency_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetDSS_Index_Frequency_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/st_GetDSS_Index_Frequency_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1044 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x6e, 0x1b, 0xc5,
	0x17, 0x57, 0xff, 0xf9, 0x3b, 0x71, 0x8e, 0x73, 0xeb, 0xa6, 0x4d, 0xb6, 0x36, 0x6d, 0x69, 0x28,
	0xa2, 0x02, 0xc9, 0x41, 0x05, 0xd1, 0xaa, 0x15, 0x91, 0x70, 0x0d, 0x95, 0x11, 0xb1, 0x8b, 0x53,
	0x5a, 0x40, 0x42, 0xa3, 0xb5, 0x77, 0xec, 0xac, 0x58, 0xcf, 0x98, 0x9d, 0x59, 0xbb, 0x7e, 0x0b,
	0xee, 0x0f, 0xc1, 0x53, 0xf0, 0x99, 0x37, 0xe1, 0x7e, 0xe7, 0x33, 0x67, 0xe6, 0xac, 0x1d, 0xaf,
	0xeb, 0x64, 0xcd, 0x17, 0xdb, 0xe3, 0x39, 0xbf, 0xcb, 0xcc, 0x9c, 0x39, 0x67, 0xe0, 0xae, 0xaf,
	0xb4, 0x8c, 0xf8, 0x3e, 0x17, 0xdd, 0x40, 0xf0, 0xfd, 0x7e, 0x24, 0xdb, 0xdc, 0x8f, 0x23, 0xae,
	0xf6, 0x95, 0x66, 0xf7, 0xb9, 0xae, 0x1e, 0x1d, 0xb1, 0x9a, 0xf0, 0xf9, 0x13, 0xf6, 0x56, 0xc4,
	0x3f, 0x89, 0xb9, 0x68, 0x8f, 0xd8, 0x1b, 0x7e, 0x19, 0xe3, 0xb4, 0x74, 0x5e, 0x24, 0x70, 0x99,
	0xc0, 0xe5, 0xb3, 0x10, 0xc5, 0xed, 0x44, 0x68, 0xe0, 0x85, 0x31, 0x57, 0x44, 0x50, 0xbc, 0x94,
	0x56, 0xe7, 0x51, 0x24, 0xa3, 0x64, 0xaa, 0x94, 0x9e, 0xea, 0x71, 0xa5, 0xbc, 0x2e, 0x4f, 0x26,
	0x9f, 0x9b, 0x9d, 0xd4, 0x5e, 0x20, 0x3a, 0x32, 0xea, 0x79, 0x3a, 0x90, 0x82, 0x82, 0xf6, 0xbe,
	0xdd, 0x00, 0x78, 0xe0, 0x45, 0x1e, 0xce, 0xf2, 0x48, 0x39, 0x6f, 0xc3, 0xb6, 0x2f, 0x7b, 0x18,
	0xc7, 0x74, 0xc4, 0x39, 0x13, 0xd2, 0xe7, 0x2c, 0xf0, 0xdd, 0x73, 0xcf, 0x9e, 0xbb, 0x51, 0xb8,
	0x59, 0x2a, 0x27, 0x4b, 0x49, 0xec, 0x05, 0x42, 0xf3, 0x2e, 0x8f, 0x1e, 0x99, 0x51, 0x73, 0x8b,
	0x70, 0x0f, 0x11, 0x56, 0x47, 0x54, 0xcd, 0x77, 0x6e, 0x81, 0x3b, 0x87, 0x8b, 0x89, 0x38, 0x0c,
	0xdd, 0xef, 0x57, 0x90, 0x31, 0xdf, 0xbc, 0x30, 0x0b, 0xaa, 0xe3, 0xa4, 0x73, 0x07, 0x56, 0x3b,
	0x91, 0xec, 0x31, 0xdf, 0xd3, 0xdc, 0xfd, 0x9f, 0x95, 0xbe, 0x3c, 0x23, 0xad, 0x03, 0x5c, 0xab,
	0xf6, 0x7a, 0x7d, 0x12, 0xcf, 0x9b, 0xf8, 0x2a, 0x86, 0x3b, 0xcf, 0xc3, 0xc6, 0x04, 0x4b, 0x52,
	0x3f, 0x90, 0xd4, 0xda, 0x38, 0xc4, 0x4a, 0xbc, 0x06, 0x2b, 0x5a, 0x92, 0xc0, 0xd2, 0x22, 0x02,
	0xcb, 0x5a, 0x5a, 0xfa, 0x6b, 0xb0, 0x96, 0xe0, 0x88, 0xfc, 0x47, 0x22, 0x07, 0x9a, 0xb6, 0xd4,
	0xef, 0xc2, 0x6e, 0x37, 0x92, 0x71, 0x9f, 0xb5, 0x46, 0x76, 0xcd, 0x8a, 0x49, 0xc1, 0x42, 0x3e,
	0xe0, 0xa1, 0xfb, 0xff, 0xec, 0x6d, 0xdc, 0xb6, 0xd8, 0xca, 0xc8, 0x6c, 0x87, 0x6a, 0x88, 0x77,
	0x0c, 0xce, 0x39, 0x80, 0x67, 0x4e, 0xa1, 0x24, 0x17, 0x3f, 0x91, 0x8b, 0xdd, 0x39, 0x58, 0x6b,
	0xe9, 0x2e, 0x14, 0x02, 0x95, 0x00, 0xf0, 0x34, 0x73, 0x73, 0x6d, 0xb4, 0xa4, 0x0c, 0xb9, 0x27,
	0xc8, 0xc6, 0x6a, 0xa0, 0x2c, 0x1c, 0x8f, 0xf1, 0x06, 0x6c, 0x4d, 0x81, 0x49, 0xf0, 0x67, 0x12,
	0x5c, 0x9f, 0x44, 0x59, 0x19, 0x0f, 0xae, 0xa6, 0x6c, 0xb2, 0xf6, 0x31, 0x66, 0x56, 0x1b, 0x13,
	0x2b, 0x50, 0x3a, 0x68, 0x1b, 0xe9, 0xe5, 0xec, 0x1d, 0x28, 0x4d, 0xad, 0xe2, 0x5e, 0x8a, 0x00,
	0xcd, 0xd4, 0xe1, 0x7a, 0x86, 0x04, 0x19, 0xfc, 0x85, 0x0c, 0x5e, 0x3d, 0x83, 0xcb, 0x5a, 0x3e,
	0x84, 0x8b, 0x52, 0x84, 0x23, 0x46, 0x46, 0x18, 0xe6, 0xaa, 0x14, 0x36, 0xe3, 0x57, 0xb2, 0xf7,
	0xe8, 0xbc, 0x41, 0xda, 0x9f, 0xaa, 0x26, 0x1a, 0xc2, 0xa4, 0xfc, 0x1d, 0x28, 0xce, 0xa5, 0x23,
	0x53, 0xbf, 0x92, 0xa9, 0x8b, 0x4f, 0xe1, 0xac, 0x95, 0x03, 0x58, 0x1b, 0xf2, 0xa0, 0x7b, 0xac,
	0xd9, 0x20, 0xe0, 0x43, 0xe5, 0xe6, 0xe7, 0x3a, 0xf0, 0x79, 0x3b, 0xe8, 0x79, 0x21, 0x39, 0x28,
	0x10, 0xe0, 0x91, 0x89, 0x77, 0x5e, 0x82, 0xf3, 0xd3, 0x78, 0x92, 0xfc, 0x8d, 0x24, 0x37, 0xa7,
	0x02, 0x93, 0x24, 0xdd, 0x49, 0x82, 0x5b, 0xf1, 0x28, 0x10, 0x5d, 0x66, 0x8e, 0x00, 0x0b, 0x9a,
	0x76, 0x57, 0xb3, 0x65, 0x2f, 0x10, 0xb4, 0x62, 0x91, 0xb5, 0x04, 0x88, 0xfe, 0x4b, 0xf3, 0x29,
	0xc9, 0xc9, 0xef, 0xe4, 0xc4, 0x9d, 0x87, 0x9d, 0x59, 0xbf, 0x8c, 0x7c, 0x1e, 0xb9, 0xb0, 0xf0,
	0xfa, 0x1b, 0x26, 0x7e, 0x6a, 0xfd, 0x16, 0x4f, 0xaa, 0x7f, 0xa4, 0xd6, 0x6f, 0x03, 0xad, 0xd8,
	0x63, 0xb8, 0x14, 0x88, 0x76, 0x18, 0x63, 0x06, 0xf9, 0x1c, 0xb3, 0x22, 0x18, 0xe0, 0xf5, 0xf5,
	0xe9, 0x72, 0xb9, 0x85, 0xec, 0xb3, 0xdf, 0x4d, 0xd0, 0xd5, 0x13, 0xb0, 0xbd, 0x75, 0x4e, 0x15,
	0xae, 0x9c, 0x4a, 0x4c, 0x96, 0xfe, 0x24, 0x4b, 0xa5, 0x53, 0x18, 0xac, 0xbd, 0x0e, 0xec, 0x8d,
	0x59, 0x08, 0x39, 0x0c, 0xf4, 0xb1, 0x8c, 0x75, 0xba, 0x2a, 0xaf, 0x65, 0xfb, 0xbc, 0x9c, 0xd0,
	0x58, 0xea, 0xc7, 0x44, 0x32, 0x55, 0xa2, 0x8f, 0xe0, 0x85, 0x6c, 0x1d, 0xb2, 0xfd, 0x17, 0xd9,
	0xbe, 0x76, 0x26, 0xa1, 0x35, 0xff, 0x00, 0x76, 0xe8, 0x04, 0xf0, 0x68, 0xe3, 0x50, 0x9b, 0xab,
	0xda, 0x96, 0x61, 0xdc, 0x13, 0xee, 0xfa, 0x02, 0xf5, 0xcf, 0x42, 0x9b, 0x16, 0x59, 0x19, 0xdd,
	0xb3, 0x38, 0xe7, 0x75, 0x28, 0xcd, 0x67, 0x24, 0x6b, 0x7f, 0x27, 0xe5, 0x6f, 0x0e, 0xd4, 0x1a,
	0xba, 0x05, 0xab, 0x5d, 0x8e, 0x6b, 0x92, 0x7d, 0xf6, 0xc4, 0xdd, 0xc8, 0xf6, 0xb0, 0x82, 0xd1,
	0x0f, 0x65, 0xff, 0x7d, 0xe7, 0x3a, 0x6c, 0x4c, 0x80, 0x24, 0xf5, 0x0f, 0x49, 0x15, 0x92, 0x08,
	0x43, 0xbf, 0xf7, 0x5d, 0x0e, 0xf2, 0xa8, 0xda, 0x97, 0x42, 0x71, 0xe7, 0x65, 0xc8, 0xd9, 0x06,
	0x9d, 0xb4, 0xcc, 0x62, 0x39, 0xdd, 0xfd, 0xa9, 0x79, 0xbf, 0x69, 0x3e, 0x9b, 0x14, 0xe8, 0x7c,
	0x00, 0x5b, 0xa6, 0x35, 0xb3, 0xa9, 0xde, 0x8c, 0x4d, 0x6f, 0x09, 0xc1, 0xe5, 0x19, 0xf0, 0x6c,
	0x07, 0x3f, 0xc4, 0x71, 0xed, 0x64, 0xdc, 0xdc, 0xec, 0xa5, 0xff, 0x70, 0x6e, 0xc3, 0x4a, 0xf2,
	0x24, 0xc0, 0x2e, 0x67, 0x18, 0xaf, 0x3c, 0xc5, 0x48, 0x0f, 0x86, 0x43, 0xfa, 0x6e, 0x8e, 0xc3,
	0xf1, 0x1d, 0xb0, 0x14, 0xc9, 0x21, 0x36, 0x2c, 0x83, 0xba, 0x5d, 0x5e, 0xfc, 0x09, 0x53, 0x1e,
	0xef, 0x44, 0xb9, 0x29, 0x87, 0x4d, 0x43, 0x52, 0xfc, 0x66, 0x09, 0x96, 0x70, 0xe0, 0xec, 0xc0,
	0x32, 0x0e, 0x4d, 0xe2, 0x7e, 0x5a, 0xc7, 0xcd, 0xc9, 0x35, 0x73, 0x38, 0xc4, 0x24, 0x7c, 0x15,
	0x96, 0xed, 0xc9, 0x29, 0xf7, 0xb3, 0x7a, 0xf6, 0x9d, 0x4f, 0x62, 0x4d, 0x4f, 0x1b, 0xb0, 0x16,
	0x93, 0xb8, 0x6f, 0xe8, 0xc1, 0xfd, 0x7c, 0x01, 0x68, 0x7e, 0x50, 0x69, 0x58, 0xc7, 0x78, 0x4b,
	0x37, 0x3b, 0x13, 0xd3, 0xb2, 0x63, 0x3c, 0x7d, 0x51, 0xcf, 0x4e, 0x8c, 0xf5, 0x09, 0xa8, 0xd1,
	0x41, 0xe3, 0x37, 0x21, 0x47, 0xa5, 0xfa, 0xcb, 0x05, 0xc4, 0x29, 0xd4, 0x79, 0x0f, 0xdc, 0x94,
	0x32, 0xde, 0xa2, 0x76, 0x14, 0xf4, 0xed, 0xa9, 0x7f, 0x55, 0x4f, 0xe7, 0x4c, 0x42, 0xa3, 0x74,
	0x84, 0xb5, 0x92, 0x58, 0x76, 0xa6, 0x1c, 0x54, 0x4f, 0xa0, 0xce, 0x7d, 0xd8, 0x9a, 0xa9, 0xba,
	0xca, 0xfd, 0x7a, 0x01, 0x57, 0x9b, 0xad, 0x54, 0x21, 0x56, 0x95, 0x8f, 0xa0, 0x14, 0xc8, 0x99,
	0xf3, 0x3e, 0x79, 0xef, 0x7e, 0x78, 0xd0, 0x95, 0xca, 0xff, 0x78, 0x3c, 0xef, 0xff, 0xd7, 0x27,
	0x71, 0x6b, 0xd9, 0xbe, 0x3a, 0x5f, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x29, 0xc2, 0x1c, 0xd6,
	0x52, 0x0b, 0x00, 0x00,
}