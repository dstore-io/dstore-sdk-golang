// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_SearchTreeNodes_Pu.proto
// DO NOT EDIT!

/*
Package im_SearchTreeNodes_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_SearchTreeNodes_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_SearchTreeNodes_Pu

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
	Country                           *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	CountryNull                       bool                        `protobuf:"varint,1005,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	BinaryCharacteristicValueId       *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=binary_characteristic_value_id,json=binaryCharacteristicValueId" json:"binary_characteristic_value_id,omitempty"`
	BinaryCharacteristicValueIdNull   bool                        `protobuf:"varint,1006,opt,name=binary_characteristic_value_id_null,json=binaryCharacteristicValueIdNull" json:"binary_characteristic_value_id_null,omitempty"`
	SortByCharacteristicIdList        *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=sort_by_characteristic_id_list,json=sortByCharacteristicIdList" json:"sort_by_characteristic_id_list,omitempty"`
	SortByCharacteristicIdListNull    bool                        `protobuf:"varint,1007,opt,name=sort_by_characteristic_id_list_null,json=sortByCharacteristicIdListNull" json:"sort_by_characteristic_id_list_null,omitempty"`
	SortOptionList                    *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=sort_option_list,json=sortOptionList" json:"sort_option_list,omitempty"`
	SortOptionListNull                bool                        `protobuf:"varint,1008,opt,name=sort_option_list_null,json=sortOptionListNull" json:"sort_option_list_null,omitempty"`
	InheritDepthOptionList            *dstore_values.StringValue  `protobuf:"bytes,9,opt,name=inherit_depth_option_list,json=inheritDepthOptionList" json:"inherit_depth_option_list,omitempty"`
	InheritDepthOptionListNull        bool                        `protobuf:"varint,1009,opt,name=inherit_depth_option_list_null,json=inheritDepthOptionListNull" json:"inherit_depth_option_list_null,omitempty"`
	RecursiveEvaluationOptionList     *dstore_values.StringValue  `protobuf:"bytes,10,opt,name=recursive_evaluation_option_list,json=recursiveEvaluationOptionList" json:"recursive_evaluation_option_list,omitempty"`
	RecursiveEvaluationOptionListNull bool                        `protobuf:"varint,1010,opt,name=recursive_evaluation_option_list_null,json=recursiveEvaluationOptionListNull" json:"recursive_evaluation_option_list_null,omitempty"`
	GetValuesForSortByCharacs         *dstore_values.BooleanValue `protobuf:"bytes,11,opt,name=get_values_for_sort_by_characs,json=getValuesForSortByCharacs" json:"get_values_for_sort_by_characs,omitempty"`
	GetValuesForSortByCharacsNull     bool                        `protobuf:"varint,1011,opt,name=get_values_for_sort_by_characs_null,json=getValuesForSortByCharacsNull" json:"get_values_for_sort_by_characs_null,omitempty"`
	OutputIntoOneId                   *dstore_values.IntegerValue `protobuf:"bytes,12,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull               bool                        `protobuf:"varint,1012,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
	FilterTreeNodesInOneId            *dstore_values.BooleanValue `protobuf:"bytes,13,opt,name=filter_tree_nodes_in_one_id,json=filterTreeNodesInOneId" json:"filter_tree_nodes_in_one_id,omitempty"`
	FilterTreeNodesInOneIdNull        bool                        `protobuf:"varint,1013,opt,name=filter_tree_nodes_in_one_id_null,json=filterTreeNodesInOneIdNull" json:"filter_tree_nodes_in_one_id_null,omitempty"`
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
	NodeDescription            *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Value2                     *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value2" json:"value2,omitempty"`
	Value3                     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value3" json:"value3,omitempty"`
	Value1                     *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=value1" json:"value1,omitempty"`
	BinaryCodeId               *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	TreeNodeId                 *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                     *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	PrePredecessorsDescription *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=pre_predecessors_description,json=prePredecessorsDescription" json:"pre_predecessors_description,omitempty"`
	PrePredecessorsTreeNodeId  *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=pre_predecessors_tree_node_id,json=prePredecessorsTreeNodeId" json:"pre_predecessors_tree_node_id,omitempty"`
	PredecessorsTreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=predecessors_tree_node_id,json=predecessorsTreeNodeId" json:"predecessors_tree_node_id,omitempty"`
	PredecessorsDescription    *dstore_values.StringValue  `protobuf:"bytes,10011,opt,name=predecessors_description,json=predecessorsDescription" json:"predecessors_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_SearchTreeNodes_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_SearchTreeNodes_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_SearchTreeNodes_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 1075 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x97, 0x5b, 0x53, 0xdb, 0x46,
	0x14, 0xc7, 0x27, 0xa5, 0xd8, 0x70, 0x20, 0xc0, 0x88, 0x94, 0x08, 0x13, 0x28, 0xc5, 0x93, 0x99,
	0x3e, 0x99, 0x02, 0x49, 0xdb, 0xa7, 0xce, 0x94, 0x4b, 0xa6, 0x6e, 0x83, 0xa1, 0x4a, 0x9b, 0xde,
	0xa6, 0x51, 0x85, 0xbc, 0x18, 0xcd, 0xc8, 0xbb, 0x9e, 0xdd, 0x15, 0x99, 0x7e, 0x8b, 0xde, 0x2f,
	0xe9, 0x73, 0x9f, 0xfb, 0xb9, 0x7a, 0xbf, 0xbf, 0xf7, 0x68, 0xcf, 0xfa, 0x22, 0x61, 0x5b, 0xee,
	0x03, 0xd8, 0xf2, 0x9e, 0xff, 0xff, 0xfc, 0xb4, 0x7b, 0x56, 0x67, 0x05, 0x77, 0x9a, 0x4a, 0x0b,
	0xc9, 0xb6, 0x19, 0x6f, 0x45, 0x9c, 0x6d, 0x77, 0xa4, 0x08, 0x59, 0x33, 0x91, 0x4c, 0x6d, 0x47,
	0x6d, 0xff, 0x01, 0x0b, 0x64, 0x78, 0xf1, 0x96, 0x64, 0xac, 0x21, 0x9a, 0x4c, 0xf9, 0xa7, 0x49,
	0x0d, 0x03, 0xb4, 0x70, 0xaa, 0xa4, 0xaa, 0x91, 0xaa, 0x36, 0x34, 0xb4, 0xb2, 0x6c, 0xad, 0x2f,
	0x83, 0x38, 0x61, 0x8a, 0x94, 0x95, 0xd5, 0x6c, 0x3e, 0x26, 0xa5, 0x90, 0x76, 0x68, 0x2d, 0x3b,
	0xd4, 0x66, 0x4a, 0x05, 0x2d, 0x66, 0x07, 0xab, 0xf9, 0x41, 0x1d, 0x44, 0xfc, 0x5c, 0xc8, 0x76,
	0xa0, 0x23, 0xc1, 0x29, 0x68, 0xeb, 0xc9, 0x02, 0xc0, 0x69, 0x20, 0x03, 0x1c, 0x65, 0x52, 0x39,
	0x2f, 0xc2, 0x4c, 0xcc, 0x2e, 0x59, 0xec, 0x47, 0x4d, 0xf7, 0xda, 0xe6, 0xb5, 0xe7, 0xe7, 0x76,
	0xd7, 0x6a, 0x16, 0xdc, 0x32, 0x45, 0x5c, 0xb3, 0x16, 0x93, 0x0f, 0xd3, 0x2b, 0xaf, 0x6c, 0x82,
	0xeb, 0x4d, 0xa7, 0x0a, 0xd7, 0xbb, 0x3a, 0x9f, 0x27, 0x71, 0xec, 0xfe, 0x58, 0x46, 0xf5, 0x8c,
	0x37, 0x67, 0x03, 0x1a, 0xf8, 0x9b, 0xf3, 0x3a, 0x2c, 0x37, 0x45, 0x1b, 0x21, 0x7c, 0x8d, 0x37,
	0xed, 0x73, 0xbc, 0xeb, 0x34, 0xcf, 0x53, 0xc5, 0x79, 0x96, 0x48, 0xd7, 0x9d, 0x2b, 0x4c, 0xf8,
	0x12, 0xb8, 0x43, 0xbc, 0x28, 0xf7, 0x4f, 0x94, 0xfb, 0x46, 0x5e, 0x64, 0x20, 0xf6, 0x61, 0x51,
	0xe9, 0x40, 0x6a, 0x3f, 0xd0, 0xbe, 0x14, 0x8f, 0x51, 0xe9, 0x4e, 0x15, 0x03, 0xcc, 0x1b, 0xcd,
	0xab, 0xda, 0x13, 0x8f, 0x1b, 0xc2, 0xd9, 0x86, 0x1b, 0x39, 0x0f, 0x4a, 0xfc, 0x33, 0x25, 0x5e,
	0x1a, 0x0c, 0x36, 0x49, 0x5f, 0x86, 0xd9, 0x34, 0x2e, 0x14, 0x09, 0xd7, 0xee, 0xd3, 0xc5, 0xe9,
	0x66, 0x30, 0xfa, 0x20, 0x0d, 0x76, 0x6e, 0xc3, 0x42, 0x4f, 0x49, 0x49, 0x7e, 0xa1, 0x24, 0xf3,
	0xdd, 0x10, 0x93, 0xe0, 0x0e, 0x94, 0x4d, 0x88, 0xfc, 0xd8, 0x9d, 0x36, 0xf6, 0x95, 0x9c, 0xbd,
	0xd2, 0x32, 0xe2, 0x2d, 0xbb, 0x6a, 0x36, 0xd4, 0xd9, 0x82, 0x79, 0xfb, 0x95, 0xac, 0x7f, 0xb5,
	0x8b, 0x66, 0x7f, 0x34, 0xce, 0x1f, 0xc1, 0xc6, 0x59, 0xc4, 0x03, 0x0c, 0x09, 0x2f, 0xb0, 0x4e,
	0x42, 0x2c, 0x93, 0x48, 0xe9, 0x28, 0xf4, 0x8d, 0x71, 0xba, 0x7e, 0xa5, 0xe2, 0xfb, 0x59, 0x23,
	0x8b, 0x83, 0x8c, 0x83, 0x19, 0xc2, 0xa5, 0x3c, 0x86, 0xea, 0xf8, 0x0c, 0x04, 0xf7, 0x1b, 0xc1,
	0x3d, 0x3b, 0xc6, 0xca, 0x00, 0x3f, 0x82, 0x0d, 0x25, 0x70, 0x6d, 0xce, 0xae, 0xf8, 0xa1, 0x53,
	0x8c, 0xdf, 0xdc, 0x72, 0xe1, 0x0c, 0x55, 0x52, 0x87, 0xfd, 0x5c, 0x92, 0x7a, 0xf3, 0x3e, 0x7e,
	0x3a, 0xf7, 0xa1, 0x3a, 0xde, 0x9f, 0x70, 0x7f, 0x27, 0xdc, 0x8d, 0xd1, 0x4e, 0x86, 0xf6, 0x10,
	0x96, 0x8c, 0x9b, 0xe8, 0xa4, 0x9b, 0x92, 0xf8, 0x66, 0x0a, 0xf9, 0x16, 0x52, 0xcd, 0x89, 0x91,
	0x18, 0xa6, 0x5d, 0x78, 0x26, 0xef, 0x42, 0x14, 0x7f, 0x10, 0x85, 0x93, 0x8d, 0x37, 0x99, 0xdf,
	0x86, 0xd5, 0x88, 0x5f, 0x20, 0x91, 0xf6, 0x9b, 0xac, 0xa3, 0x2f, 0x32, 0x08, 0xb3, 0x85, 0x08,
	0x2b, 0x56, 0x7c, 0x98, 0x6a, 0x07, 0x50, 0x0e, 0x60, 0x63, 0xa4, 0x2d, 0x31, 0xfd, 0x49, 0x4c,
	0x95, 0xe1, 0x06, 0x86, 0x2d, 0x84, 0x4d, 0xc9, 0xc2, 0x44, 0xaa, 0xe8, 0x92, 0xf9, 0x2c, 0xcd,
	0x6e, 0x9e, 0x59, 0x19, 0x44, 0x28, 0x44, 0x5c, 0xef, 0x79, 0x1c, 0xf5, 0x2c, 0x06, 0x48, 0xdf,
	0x84, 0xdb, 0x45, 0x49, 0x08, 0xf8, 0x2f, 0x02, 0x7e, 0x6e, 0xac, 0x5d, 0xb7, 0xf6, 0x5a, 0x4c,
	0x53, 0xe1, 0x2a, 0x1f, 0x9f, 0xb5, 0x7e, 0xb6, 0x54, 0x94, 0x3b, 0x37, 0x74, 0xb3, 0x9c, 0x09,
	0x11, 0xb3, 0x80, 0x13, 0xf6, 0x2a, 0x5a, 0x98, 0x6f, 0xea, 0x9e, 0x90, 0x0f, 0x06, 0xca, 0x47,
	0x39, 0x6f, 0x40, 0x75, 0xbc, 0x3f, 0x01, 0xff, 0x4d, 0xc0, 0xeb, 0x23, 0x8d, 0x0c, 0xec, 0x6b,
	0xe0, 0x88, 0x44, 0x77, 0x12, 0xed, 0xe3, 0x5e, 0x15, 0xbe, 0xe0, 0x66, 0x37, 0xcf, 0x17, 0xef,
	0xe6, 0x45, 0x92, 0xd5, 0x51, 0x75, 0xc2, 0xd3, 0x1d, 0x7c, 0x17, 0x6e, 0x5e, 0x75, 0x22, 0x94,
	0x7f, 0x08, 0x65, 0x39, 0x27, 0x31, 0x00, 0xef, 0xc2, 0xda, 0x79, 0x14, 0xe3, 0xa6, 0xe8, 0x3f,
	0xc3, 0x15, 0x3a, 0x74, 0x49, 0xae, 0x17, 0x4f, 0xd5, 0x0a, 0xe9, 0x7b, 0x3d, 0xb4, 0xce, 0x09,
	0xe8, 0x08, 0x36, 0xc7, 0x38, 0x13, 0xd9, 0xbf, 0xb6, 0x0c, 0x87, 0x5b, 0xa4, 0x80, 0x5b, 0xdf,
	0xcf, 0xc2, 0x8c, 0xc7, 0x54, 0x47, 0x70, 0xc5, 0x9c, 0x17, 0x60, 0xda, 0xb4, 0x5e, 0xdb, 0x17,
	0x7b, 0x85, 0x67, 0x1b, 0x3a, 0xb5, 0xe5, 0xa3, 0xf4, 0xbf, 0x47, 0x81, 0xce, 0x7b, 0xb0, 0x94,
	0x36, 0x5d, 0x7f, 0xa0, 0xeb, 0x62, 0xb3, 0x9b, 0x42, 0x71, 0x2d, 0x27, 0xce, 0xf7, 0xe6, 0x63,
	0xbc, 0xae, 0xf7, 0xaf, 0xbd, 0xc5, 0x76, 0xf6, 0x07, 0x6c, 0x28, 0x65, 0xdb, 0xec, 0xb1, 0x7b,
	0xa5, 0x8e, 0x1b, 0x57, 0x1c, 0xe9, 0x28, 0x70, 0x4c, 0x9f, 0x5e, 0x37, 0x1c, 0xf7, 0xe7, 0x14,
	0x76, 0x0e, 0x6c, 0x42, 0xa9, 0x6a, 0xa7, 0x36, 0xc1, 0xa9, 0xa4, 0xd6, 0x9d, 0x82, 0x1a, 0x36,
	0x34, 0x2f, 0x55, 0x3b, 0x1f, 0xc0, 0x2d, 0x9e, 0xb4, 0xcf, 0x70, 0x7e, 0xc5, 0xb9, 0xcf, 0x62,
	0xd6, 0x66, 0x5c, 0x9b, 0x09, 0xc6, 0x53, 0x50, 0x12, 0x6b, 0x97, 0x15, 0x17, 0x91, 0x4b, 0x06,
	0x27, 0xe7, 0x47, 0x56, 0x5e, 0xe7, 0x9e, 0x11, 0x57, 0x7e, 0x28, 0xc1, 0x14, 0x66, 0x72, 0x56,
	0xa0, 0x94, 0xb6, 0x3e, 0xac, 0x84, 0x4f, 0x1a, 0xe8, 0x37, 0xed, 0x4d, 0xe3, 0x25, 0x2e, 0xee,
	0x3d, 0x58, 0x32, 0xed, 0x1e, 0xe9, 0x42, 0x19, 0x99, 0x0d, 0xe8, 0x7e, 0xda, 0x28, 0x7c, 0x1a,
	0x2c, 0xa6, 0xa2, 0xc3, 0xbe, 0xc6, 0xd9, 0x83, 0x92, 0x09, 0xdb, 0x75, 0x3f, 0x2b, 0x56, 0xdb,
	0xd0, 0x9e, 0x68, 0xcf, 0xfd, 0x7c, 0x52, 0xd1, 0x5e, 0x4f, 0xb4, 0xe3, 0x7e, 0x31, 0xa9, 0x68,
	0x07, 0x0f, 0x2a, 0x0b, 0xdd, 0xb6, 0x68, 0x0f, 0x4a, 0x5f, 0x36, 0x26, 0x38, 0xa8, 0xd8, 0xf6,
	0x48, 0xa7, 0xa4, 0x57, 0x60, 0x3e, 0x73, 0xd4, 0xfa, 0x6a, 0x02, 0x07, 0xd0, 0xfd, 0x53, 0xd6,
	0x5d, 0x28, 0x77, 0xa5, 0x5f, 0x4f, 0x20, 0x2d, 0x71, 0x92, 0x7d, 0x08, 0xb7, 0x3a, 0x92, 0xf9,
	0xf8, 0xd7, 0x64, 0x21, 0xd6, 0x9d, 0x90, 0x2a, 0xb3, 0x5a, 0xdf, 0x14, 0xcf, 0x42, 0x05, 0xc5,
	0xa7, 0x03, 0xfa, 0xc1, 0x85, 0x7b, 0x04, 0xeb, 0x57, 0xec, 0x33, 0xb7, 0xf9, 0xed, 0x04, 0xac,
	0xab, 0xb9, 0x04, 0x03, 0x67, 0xcb, 0x77, 0x60, 0x75, 0xb4, 0xf7, 0x93, 0x09, 0xbc, 0x57, 0x3a,
	0xc3, 0x8d, 0x1f, 0x82, 0x3b, 0x72, 0x4e, 0xbe, 0x2b, 0x9e, 0x93, 0x9b, 0x9d, 0xe1, 0x13, 0xb2,
	0x7f, 0x0c, 0x6b, 0x91, 0xc8, 0x6d, 0xe5, 0xfe, 0x6b, 0xc9, 0xfb, 0xb5, 0xff, 0xf7, 0xc2, 0x72,
	0x56, 0x32, 0xaf, 0x06, 0x7b, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0x91, 0x62, 0x8d, 0xe9,
	0x0c, 0x00, 0x00,
}
