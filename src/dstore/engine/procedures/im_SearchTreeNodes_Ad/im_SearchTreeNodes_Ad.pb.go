// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_SearchTreeNodes_Ad.proto
// DO NOT EDIT!

/*
Package im_SearchTreeNodes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_SearchTreeNodes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_SearchTreeNodes_Ad

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
	LevelId                           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	LevelIdNull                       bool                        `protobuf:"varint,1001,opt,name=level_id_null,json=levelIdNull" json:"level_id_null,omitempty"`
	DomainTreeNodeId                  *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=domain_tree_node_id,json=domainTreeNodeId" json:"domain_tree_node_id,omitempty"`
	DomainTreeNodeIdNull              bool                        `protobuf:"varint,1002,opt,name=domain_tree_node_id_null,json=domainTreeNodeIdNull" json:"domain_tree_node_id_null,omitempty"`
	StartAtRowNo                      *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=start_at_row_no,json=startAtRowNo" json:"start_at_row_no,omitempty"`
	StartAtRowNoNull                  bool                        `protobuf:"varint,1003,opt,name=start_at_row_no_null,json=startAtRowNoNull" json:"start_at_row_no_null,omitempty"`
	RowCount                          *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=row_count,json=rowCount" json:"row_count,omitempty"`
	RowCountNull                      bool                        `protobuf:"varint,1004,opt,name=row_count_null,json=rowCountNull" json:"row_count_null,omitempty"`
	IncludeInactiveNodes              *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=include_inactive_nodes,json=includeInactiveNodes" json:"include_inactive_nodes,omitempty"`
	IncludeInactiveNodesNull          bool                        `protobuf:"varint,1005,opt,name=include_inactive_nodes_null,json=includeInactiveNodesNull" json:"include_inactive_nodes_null,omitempty"`
	Country                           *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=country" json:"country,omitempty"`
	CountryNull                       bool                        `protobuf:"varint,1006,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	BinaryCharacteristicValueId       *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=binary_characteristic_value_id,json=binaryCharacteristicValueId" json:"binary_characteristic_value_id,omitempty"`
	BinaryCharacteristicValueIdNull   bool                        `protobuf:"varint,1007,opt,name=binary_characteristic_value_id_null,json=binaryCharacteristicValueIdNull" json:"binary_characteristic_value_id_null,omitempty"`
	SortByCharacteristicIdList        *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=sort_by_characteristic_id_list,json=sortByCharacteristicIdList" json:"sort_by_characteristic_id_list,omitempty"`
	SortByCharacteristicIdListNull    bool                        `protobuf:"varint,1008,opt,name=sort_by_characteristic_id_list_null,json=sortByCharacteristicIdListNull" json:"sort_by_characteristic_id_list_null,omitempty"`
	SortOptionList                    *dstore_values.StringValue  `protobuf:"bytes,9,opt,name=sort_option_list,json=sortOptionList" json:"sort_option_list,omitempty"`
	SortOptionListNull                bool                        `protobuf:"varint,1009,opt,name=sort_option_list_null,json=sortOptionListNull" json:"sort_option_list_null,omitempty"`
	InheritDepthOptionList            *dstore_values.StringValue  `protobuf:"bytes,10,opt,name=inherit_depth_option_list,json=inheritDepthOptionList" json:"inherit_depth_option_list,omitempty"`
	InheritDepthOptionListNull        bool                        `protobuf:"varint,1010,opt,name=inherit_depth_option_list_null,json=inheritDepthOptionListNull" json:"inherit_depth_option_list_null,omitempty"`
	RecursiveEvaluationOptionList     *dstore_values.StringValue  `protobuf:"bytes,11,opt,name=recursive_evaluation_option_list,json=recursiveEvaluationOptionList" json:"recursive_evaluation_option_list,omitempty"`
	RecursiveEvaluationOptionListNull bool                        `protobuf:"varint,1011,opt,name=recursive_evaluation_option_list_null,json=recursiveEvaluationOptionListNull" json:"recursive_evaluation_option_list_null,omitempty"`
	GetValuesForSortByCharacs         *dstore_values.BooleanValue `protobuf:"bytes,12,opt,name=get_values_for_sort_by_characs,json=getValuesForSortByCharacs" json:"get_values_for_sort_by_characs,omitempty"`
	GetValuesForSortByCharacsNull     bool                        `protobuf:"varint,1012,opt,name=get_values_for_sort_by_characs_null,json=getValuesForSortByCharacsNull" json:"get_values_for_sort_by_characs_null,omitempty"`
	OutputIntoOneId                   *dstore_values.IntegerValue `protobuf:"bytes,13,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull               bool                        `protobuf:"varint,1013,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
	FilterTreeNodesInOneId            *dstore_values.BooleanValue `protobuf:"bytes,14,opt,name=filter_tree_nodes_in_one_id,json=filterTreeNodesInOneId" json:"filter_tree_nodes_in_one_id,omitempty"`
	FilterTreeNodesInOneIdNull        bool                        `protobuf:"varint,1014,opt,name=filter_tree_nodes_in_one_id_null,json=filterTreeNodesInOneIdNull" json:"filter_tree_nodes_in_one_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetLevelId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelId
	}
	return nil
}

func (m *Parameters) GetDomainTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DomainTreeNodeId
	}
	return nil
}

func (m *Parameters) GetStartAtRowNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.StartAtRowNo
	}
	return nil
}

func (m *Parameters) GetRowCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowCount
	}
	return nil
}

func (m *Parameters) GetIncludeInactiveNodes() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeInactiveNodes
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicValueId
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

func (m *Parameters) GetOutputIntoOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputIntoOneId
	}
	return nil
}

func (m *Parameters) GetFilterTreeNodesInOneId() *dstore_values.BooleanValue {
	if m != nil {
		return m.FilterTreeNodesInOneId
	}
	return nil
}

type Response struct {
	Error                    *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation          []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message                  []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row                      []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	NumberOfElementsInResult *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=number_of_elements_in_result,json=numberOfElementsInResult" json:"number_of_elements_in_result,omitempty"`
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

func (m *Response) GetNumberOfElementsInResult() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfElementsInResult
	}
	return nil
}

type Response_Row struct {
	RowId                      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Active                     *dstore_values.BooleanValue `protobuf:"bytes,10001,opt,name=active" json:"active,omitempty"`
	NodeDescription            *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Value2                     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value2" json:"value2,omitempty"`
	Value3                     *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=value3" json:"value3,omitempty"`
	Value1                     *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=value1" json:"value1,omitempty"`
	BinaryCodeId               *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	TreeNodeId                 *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                     *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	PrePredecessorsDescription *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=pre_predecessors_description,json=prePredecessorsDescription" json:"pre_predecessors_description,omitempty"`
	PrePredecessorsTreeNodeId  *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=pre_predecessors_tree_node_id,json=prePredecessorsTreeNodeId" json:"pre_predecessors_tree_node_id,omitempty"`
	PredecessorsTreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,10011,opt,name=predecessors_tree_node_id,json=predecessorsTreeNodeId" json:"predecessors_tree_node_id,omitempty"`
	PredecessorsDescription    *dstore_values.StringValue  `protobuf:"bytes,10012,opt,name=predecessors_description,json=predecessorsDescription" json:"predecessors_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
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

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
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

func (m *Response_Row) GetPrePredecessorsDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PrePredecessorsDescription
	}
	return nil
}

func (m *Response_Row) GetPrePredecessorsTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrePredecessorsTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetPredecessorsTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredecessorsTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetPredecessorsDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PredecessorsDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_SearchTreeNodes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_SearchTreeNodes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_SearchTreeNodes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 1129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x57, 0x59, 0x53, 0x1b, 0x47,
	0x10, 0x2e, 0x87, 0x20, 0x41, 0x83, 0x81, 0x5a, 0x08, 0x5e, 0x84, 0x21, 0x04, 0x95, 0xab, 0xf2,
	0x24, 0xc2, 0xe1, 0x24, 0x4f, 0xae, 0x32, 0x87, 0x2b, 0x4a, 0x8c, 0xb0, 0xe5, 0xc4, 0xb9, 0x2a,
	0xde, 0x2c, 0xab, 0x01, 0xb6, 0x6a, 0x35, 0xa3, 0x9a, 0x19, 0xe1, 0xca, 0xbf, 0xc8, 0x7d, 0xff,
	0x85, 0xfc, 0xa2, 0x3c, 0xe7, 0x21, 0xf7, 0x9d, 0xe7, 0xf4, 0x4c, 0x8f, 0x84, 0x76, 0x91, 0xb4,
	0xf2, 0x03, 0x48, 0xab, 0xe9, 0xef, 0xd8, 0xde, 0x9e, 0xe9, 0x5e, 0xd8, 0x69, 0x28, 0x2d, 0x24,
	0xdb, 0x60, 0xfc, 0x34, 0xe6, 0x6c, 0xa3, 0x25, 0x45, 0xc4, 0x1a, 0x6d, 0xc9, 0xd4, 0x46, 0xdc,
	0x0c, 0x1e, 0xb0, 0x50, 0x46, 0x67, 0xaf, 0x4b, 0xc6, 0x6a, 0xa2, 0xc1, 0x54, 0x70, 0xbb, 0x51,
	0xc1, 0x00, 0x2d, 0xbc, 0x32, 0xa1, 0x2a, 0x84, 0xaa, 0xf4, 0x0d, 0x2d, 0xcd, 0x3b, 0xea, 0xf3,
	0x30, 0x69, 0x33, 0x45, 0xc8, 0xd2, 0x52, 0x5a, 0x8f, 0x49, 0x29, 0xa4, 0x5b, 0x5a, 0x4e, 0x2f,
	0x35, 0x99, 0x52, 0xe1, 0x29, 0x73, 0x8b, 0xe5, 0xec, 0xa2, 0x0e, 0x63, 0x7e, 0x22, 0x64, 0x33,
	0xd4, 0xb1, 0xe0, 0x14, 0xb4, 0xfe, 0xdd, 0x2c, 0xc0, 0xbd, 0x50, 0x86, 0xb8, 0xca, 0xa4, 0xf2,
	0x5e, 0x84, 0x89, 0x84, 0x9d, 0xb3, 0x24, 0x88, 0x1b, 0xfe, 0x95, 0xb5, 0x2b, 0xcf, 0x4f, 0x6d,
	0x2d, 0x57, 0x9c, 0x71, 0xe7, 0x29, 0xe6, 0x9a, 0x9d, 0x32, 0xf9, 0xd0, 0x5c, 0xd5, 0x8b, 0x36,
	0xb8, 0xda, 0xf0, 0xca, 0x70, 0xb5, 0x83, 0x0b, 0x78, 0x3b, 0x49, 0xfc, 0x1f, 0x8b, 0x88, 0x9e,
	0xa8, 0x4f, 0xb9, 0x80, 0x1a, 0xfe, 0xe6, 0xbd, 0x0a, 0xf3, 0x0d, 0xd1, 0x44, 0x13, 0x81, 0xc6,
	0x9b, 0x0e, 0x38, 0xde, 0xb5, 0xd1, 0x79, 0x2a, 0x5f, 0x67, 0x8e, 0x70, 0x9d, 0x5c, 0xa1, 0xe0,
	0x4b, 0xe0, 0xf7, 0xe1, 0x22, 0xed, 0x9f, 0x48, 0x7b, 0x21, 0x0b, 0xb2, 0x26, 0x76, 0x61, 0x56,
	0xe9, 0x50, 0xea, 0x20, 0xd4, 0x81, 0x14, 0x8f, 0x11, 0xe9, 0x8f, 0xe5, 0x1b, 0x98, 0xb6, 0x98,
	0xdb, 0xba, 0x2e, 0x1e, 0xd7, 0x84, 0xb7, 0x01, 0x0b, 0x19, 0x0e, 0x12, 0xfe, 0x99, 0x84, 0xe7,
	0x7a, 0x83, 0xad, 0xe8, 0xcb, 0x30, 0x69, 0xe2, 0x22, 0xd1, 0xe6, 0xda, 0x7f, 0x3a, 0x5f, 0x6e,
	0x02, 0xa3, 0xf7, 0x4c, 0xb0, 0x77, 0x03, 0x66, 0xba, 0x48, 0x12, 0xf9, 0x85, 0x44, 0xa6, 0x3b,
	0x21, 0x56, 0xe0, 0x3e, 0x2c, 0xc6, 0x3c, 0x4a, 0xda, 0x26, 0x0b, 0x3c, 0x8c, 0x74, 0x7c, 0x4e,
	0x39, 0x51, 0xfe, 0x78, 0x5f, 0xb5, 0x63, 0x21, 0x12, 0x16, 0x72, 0x52, 0x5b, 0x70, 0xd0, 0xaa,
	0x43, 0xda, 0x72, 0xf4, 0x6e, 0xc1, 0x72, 0x7f, 0x4a, 0xb2, 0xf1, 0x2b, 0xd9, 0xf0, 0xfb, 0x61,
	0xad, 0xa5, 0x1d, 0x28, 0x5a, 0xd7, 0xf2, 0x03, 0xbf, 0x60, 0x3d, 0x94, 0x32, 0x1e, 0x94, 0x96,
	0x31, 0x3f, 0x75, 0x85, 0xe4, 0x42, 0xbd, 0x75, 0x98, 0x76, 0x5f, 0x49, 0xe6, 0x37, 0x57, 0x47,
	0xee, 0x47, 0xcb, 0xfc, 0x3e, 0xac, 0x1e, 0xa3, 0x25, 0x0c, 0x89, 0xce, 0xb0, 0x74, 0x23, 0xac,
	0xdc, 0x58, 0xe9, 0x38, 0x0a, 0x2c, 0xb1, 0x29, 0xa9, 0x62, 0x7e, 0x8a, 0x97, 0x89, 0x62, 0x2f,
	0xc5, 0x60, 0x97, 0xb0, 0xba, 0x0e, 0xa1, 0x3c, 0x5c, 0x81, 0xcc, 0xfd, 0x4e, 0xe6, 0x9e, 0x1d,
	0x42, 0x65, 0x0d, 0x3f, 0x82, 0x55, 0x25, 0xb0, 0x5c, 0x8e, 0x2f, 0xf1, 0x21, 0x53, 0x82, 0xdf,
	0xfc, 0x89, 0xdc, 0x0c, 0x95, 0x0c, 0xc3, 0x6e, 0x46, 0xa4, 0xda, 0xb8, 0x8b, 0x9f, 0xde, 0x5d,
	0x28, 0x0f, 0xe7, 0x27, 0xbb, 0x7f, 0x90, 0xdd, 0xd5, 0xc1, 0x4c, 0xd6, 0xed, 0x3e, 0xcc, 0x59,
	0x36, 0xd1, 0x32, 0xe7, 0x04, 0xf9, 0x9b, 0xcc, 0xf5, 0x37, 0x63, 0x30, 0x47, 0x16, 0x62, 0x3d,
	0x6d, 0xc1, 0x33, 0x59, 0x16, 0x72, 0xf1, 0x27, 0xb9, 0xf0, 0xd2, 0xf1, 0x56, 0xf9, 0x0d, 0x58,
	0x8a, 0xf9, 0x19, 0x3a, 0xd2, 0x41, 0x83, 0xb5, 0xf4, 0x59, 0xca, 0x02, 0xe4, 0x5a, 0x58, 0x74,
	0xe0, 0x7d, 0x83, 0xed, 0xb1, 0xb2, 0x07, 0xab, 0x03, 0x69, 0xc9, 0xd3, 0x5f, 0xe4, 0xa9, 0xd4,
	0x9f, 0xc0, 0x7a, 0x8b, 0x60, 0x4d, 0xb2, 0xa8, 0x2d, 0x95, 0xd9, 0x07, 0xcc, 0xa8, 0xdb, 0x63,
	0x34, 0x65, 0x71, 0x2a, 0xd7, 0xe2, 0x4a, 0x97, 0xe3, 0xa0, 0x4b, 0xd1, 0xe3, 0xf4, 0x3e, 0xdc,
	0xc8, 0x13, 0x21, 0xc3, 0x7f, 0x93, 0xe1, 0xe7, 0x86, 0xd2, 0x75, 0x6a, 0xef, 0x94, 0x69, 0x2a,
	0x5c, 0x15, 0xe0, 0xf1, 0x1f, 0xa4, 0x4b, 0x45, 0xf9, 0xd3, 0xf9, 0x27, 0xc4, 0x12, 0x52, 0xd8,
	0x6f, 0xea, 0x8e, 0x90, 0x0f, 0x7a, 0xca, 0x47, 0x79, 0xaf, 0x41, 0x79, 0x38, 0x3f, 0x19, 0xfe,
	0x87, 0x0c, 0xaf, 0x0c, 0x24, 0xb2, 0x66, 0x5f, 0x01, 0x4f, 0xb4, 0x75, 0xab, 0xad, 0xf1, 0xc8,
	0xd1, 0x22, 0x10, 0xdc, 0xee, 0xe6, 0xab, 0xf9, 0xbb, 0x79, 0x96, 0x60, 0x55, 0x44, 0x1d, 0x71,
	0xb3, 0x83, 0x6f, 0xc2, 0xb5, 0xcb, 0x4c, 0x64, 0xe5, 0x5f, 0xb2, 0x32, 0x9f, 0x81, 0x58, 0x03,
	0x6f, 0xc1, 0xf2, 0x49, 0x9c, 0xe0, 0xa6, 0xb8, 0x68, 0x2b, 0x0a, 0x19, 0x3a, 0x4e, 0x66, 0xf2,
	0x53, 0xb5, 0x48, 0xf8, 0x6e, 0x5b, 0xaf, 0x72, 0x32, 0x74, 0x00, 0x6b, 0x43, 0x98, 0xc9, 0xd9,
	0x7f, 0xae, 0x0c, 0xfb, 0x53, 0x18, 0x83, 0xeb, 0xdf, 0x4f, 0xc2, 0x44, 0x9d, 0xa9, 0x96, 0xe0,
	0x8a, 0x79, 0x2f, 0xc0, 0xb8, 0x9d, 0x06, 0x5c, 0xab, 0xee, 0x16, 0x9e, 0x9b, 0x31, 0x68, 0x52,
	0x38, 0x30, 0xff, 0xeb, 0x14, 0xe8, 0xbd, 0x0d, 0x73, 0x66, 0x0e, 0x08, 0x7a, 0x06, 0x01, 0xec,
	0xbf, 0x63, 0x08, 0xae, 0x64, 0xc0, 0xd9, 0x71, 0xe1, 0x10, 0xaf, 0xab, 0x17, 0xd7, 0xf5, 0xd9,
	0x66, 0xfa, 0x07, 0xec, 0x71, 0x45, 0x37, 0x7f, 0x60, 0x43, 0x35, 0x8c, 0xab, 0x97, 0x18, 0x69,
	0x3a, 0x39, 0xa4, 0xcf, 0x7a, 0x27, 0x1c, 0xf7, 0xe7, 0x18, 0x36, 0x33, 0xec, 0x8b, 0x06, 0xb5,
	0x59, 0x19, 0x61, 0x50, 0xaa, 0x74, 0x52, 0x50, 0xc1, 0x1e, 0x5b, 0x37, 0x68, 0xef, 0x5d, 0xb8,
	0xce, 0xdb, 0xcd, 0x63, 0xcc, 0xaf, 0x38, 0x09, 0x58, 0xc2, 0x9a, 0x8c, 0x6b, 0x9b, 0x60, 0x1c,
	0xcc, 0xda, 0x89, 0xf6, 0x59, 0x7e, 0x11, 0xf9, 0x44, 0x70, 0x74, 0x72, 0xe0, 0xe0, 0x55, 0x5e,
	0xb7, 0xe0, 0xd2, 0x0f, 0x05, 0x18, 0x43, 0x25, 0x6f, 0x11, 0x0a, 0xa6, 0x1b, 0x63, 0x25, 0x7c,
	0x58, 0x43, 0xbe, 0xf1, 0xfa, 0x38, 0x5e, 0xe2, 0xc3, 0xdd, 0x81, 0x02, 0xb5, 0x3f, 0xff, 0xa3,
	0x5a, 0x7e, 0x89, 0xb8, 0x58, 0xef, 0x0e, 0xcc, 0xd9, 0xb9, 0x05, 0xef, 0x29, 0x92, 0xb1, 0xdd,
	0xb6, 0xfe, 0xc7, 0xb5, 0xdc, 0x33, 0x64, 0xd6, 0x80, 0xf6, 0x2f, 0x30, 0xde, 0x36, 0x14, 0x6c,
	0xd8, 0x96, 0xff, 0x49, 0x3e, 0xda, 0x85, 0x76, 0x41, 0xdb, 0xfe, 0xa7, 0xa3, 0x82, 0xb6, 0xbb,
	0xa0, 0x4d, 0xff, 0xb3, 0x51, 0x41, 0x9b, 0x38, 0x71, 0xcd, 0x74, 0x9a, 0xa9, 0x9b, 0xf8, 0x3e,
	0xaf, 0x8d, 0x30, 0x71, 0xb9, 0xa6, 0x4a, 0xe3, 0xde, 0x2d, 0x98, 0x4e, 0xcd, 0x8c, 0x5f, 0x8c,
	0xc0, 0x00, 0xfa, 0x62, 0x5c, 0xbc, 0x09, 0xc5, 0x0e, 0xf4, 0xcb, 0x11, 0xa0, 0x05, 0x4e, 0xb0,
	0xf7, 0xe0, 0x7a, 0x4b, 0xb2, 0x00, 0xff, 0x1a, 0x2c, 0xc2, 0x6a, 0x15, 0x52, 0xa5, 0x9e, 0xd6,
	0x57, 0xf9, 0x59, 0x28, 0x21, 0xf8, 0x5e, 0x0f, 0xbe, 0xf7, 0xc1, 0x3d, 0x82, 0x95, 0x4b, 0xf4,
	0xa9, 0xdb, 0xfc, 0x7a, 0x04, 0xaf, 0x4b, 0x19, 0x81, 0x9e, 0x21, 0xf9, 0x4d, 0x58, 0x1a, 0xcc,
	0xfd, 0xcd, 0x08, 0xdc, 0x8b, 0xad, 0xfe, 0xc4, 0x0f, 0xc1, 0x1f, 0x98, 0x93, 0x6f, 0xf3, 0x73,
	0x72, 0xad, 0xd5, 0x3f, 0x21, 0xbb, 0x87, 0x38, 0x73, 0x8a, 0xcc, 0x01, 0x70, 0xf1, 0x7e, 0xf5,
	0x4e, 0xe5, 0xc9, 0xde, 0xbc, 0x8e, 0x0b, 0xf6, 0x1d, 0x67, 0xfb, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0xcd, 0x5d, 0x1f, 0xb2, 0x0d, 0x00, 0x00,
}
