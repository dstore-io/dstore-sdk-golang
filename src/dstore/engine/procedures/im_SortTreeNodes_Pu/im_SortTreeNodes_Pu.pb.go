// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_SortTreeNodes_Pu.proto
// DO NOT EDIT!

/*
Package im_SortTreeNodes_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_SortTreeNodes_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_SortTreeNodes_Pu

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
	SortByCharacteristicIdList        *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=sort_by_characteristic_id_list,json=sortByCharacteristicIdList" json:"sort_by_characteristic_id_list,omitempty"`
	SortByCharacteristicIdListNull    bool                        `protobuf:"varint,1001,opt,name=sort_by_characteristic_id_list_null,json=sortByCharacteristicIdListNull" json:"sort_by_characteristic_id_list_null,omitempty"`
	SortOptionList                    *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=sort_option_list,json=sortOptionList" json:"sort_option_list,omitempty"`
	SortOptionListNull                bool                        `protobuf:"varint,1002,opt,name=sort_option_list_null,json=sortOptionListNull" json:"sort_option_list_null,omitempty"`
	InheritDepthOptionList            *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=inherit_depth_option_list,json=inheritDepthOptionList" json:"inherit_depth_option_list,omitempty"`
	InheritDepthOptionListNull        bool                        `protobuf:"varint,1003,opt,name=inherit_depth_option_list_null,json=inheritDepthOptionListNull" json:"inherit_depth_option_list_null,omitempty"`
	RecursiveEvaluationOptionList     *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=recursive_evaluation_option_list,json=recursiveEvaluationOptionList" json:"recursive_evaluation_option_list,omitempty"`
	RecursiveEvaluationOptionListNull bool                        `protobuf:"varint,1004,opt,name=recursive_evaluation_option_list_null,json=recursiveEvaluationOptionListNull" json:"recursive_evaluation_option_list_null,omitempty"`
	BinaryCharacteristicValueId       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=binary_characteristic_value_id,json=binaryCharacteristicValueId" json:"binary_characteristic_value_id,omitempty"`
	BinaryCharacteristicValueIdNull   bool                        `protobuf:"varint,1005,opt,name=binary_characteristic_value_id_null,json=binaryCharacteristicValueIdNull" json:"binary_characteristic_value_id_null,omitempty"`
	GetValuesForSortByCharacs         *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=get_values_for_sort_by_characs,json=getValuesForSortByCharacs" json:"get_values_for_sort_by_characs,omitempty"`
	GetValuesForSortByCharacsNull     bool                        `protobuf:"varint,1006,opt,name=get_values_for_sort_by_characs_null,json=getValuesForSortByCharacsNull" json:"get_values_for_sort_by_characs_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Parameters) GetBinaryCharacteristicValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicValueId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_SortTreeNodes_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_SortTreeNodes_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_SortTreeNodes_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x96, 0xeb, 0x4e, 0xdb, 0x4a,
	0x10, 0xc7, 0xc5, 0xc9, 0x49, 0x82, 0xe6, 0x20, 0x40, 0x3e, 0x2a, 0x75, 0x92, 0x92, 0x02, 0x51,
	0xa5, 0x7e, 0xa8, 0x9c, 0x92, 0xa8, 0x52, 0x3f, 0x55, 0x2a, 0x37, 0x09, 0x01, 0x29, 0x35, 0x94,
	0xaa, 0x95, 0x8a, 0xeb, 0xd8, 0xdb, 0x60, 0x29, 0xf1, 0x46, 0xbb, 0x0e, 0x88, 0xb7, 0xe8, 0xfd,
	0xfa, 0x10, 0x7d, 0x8b, 0x3e, 0x4a, 0xa5, 0x5e, 0x9f, 0xa1, 0xb3, 0xbb, 0xce, 0xc5, 0x26, 0x89,
	0xd3, 0x0f, 0x10, 0xaf, 0x77, 0xff, 0xff, 0xfd, 0xed, 0xec, 0xcc, 0xc8, 0x50, 0x71, 0x79, 0x40,
	0x19, 0x29, 0x13, 0xbf, 0xe1, 0xf9, 0xa4, 0xdc, 0x66, 0xd4, 0x21, 0x6e, 0x87, 0x11, 0x5e, 0xf6,
	0x5a, 0xd6, 0x01, 0x65, 0xc1, 0x21, 0x23, 0xa4, 0x46, 0x5d, 0xc2, 0xad, 0xfd, 0x8e, 0x81, 0xd3,
	0x01, 0xd5, 0x96, 0x95, 0xc6, 0x50, 0x1a, 0x63, 0xc8, 0xc2, 0xfc, 0xff, 0xa1, 0xed, 0xa9, 0xdd,
	0xec, 0x10, 0xae, 0x74, 0xf9, 0x5c, 0x74, 0x2f, 0xc2, 0x18, 0x65, 0xe1, 0x54, 0x21, 0x3a, 0xd5,
	0x22, 0x9c, 0xdb, 0x0d, 0x12, 0x4e, 0x96, 0xe2, 0x93, 0x81, 0xed, 0xf9, 0xcf, 0x28, 0x6b, 0xd9,
	0x81, 0x47, 0x7d, 0xb5, 0x68, 0xe5, 0x6b, 0x16, 0x60, 0xdf, 0x66, 0x36, 0xce, 0x12, 0xc6, 0xb5,
	0x63, 0x28, 0x72, 0x84, 0xb2, 0xea, 0xe7, 0x96, 0x73, 0x82, 0xaf, 0x1d, 0x7c, 0xeb, 0xf1, 0xc0,
	0x73, 0x2c, 0xcf, 0xb5, 0x9a, 0xf8, 0xa4, 0x4f, 0x2d, 0x4d, 0x5d, 0xff, 0xaf, 0x92, 0x37, 0xc2,
	0xc3, 0x84, 0xa4, 0x3c, 0x60, 0x9e, 0xdf, 0x38, 0x12, 0x03, 0x33, 0x2f, 0x1c, 0xd6, 0xce, 0xd7,
	0x23, 0xfa, 0x6d, 0x77, 0x17, 0x7f, 0xb5, 0x5d, 0x28, 0x8d, 0xf7, 0xb7, 0xfc, 0x4e, 0xb3, 0xa9,
	0x7f, 0xcb, 0xe2, 0x2e, 0xd3, 0x66, 0x71, 0xb4, 0x53, 0x0d, 0x97, 0x69, 0x1b, 0x30, 0x2f, 0xdd,
	0x68, 0x5b, 0x9c, 0x48, 0xf1, 0xfd, 0x93, 0xc8, 0x37, 0x2b, 0x34, 0xf7, 0xa4, 0x44, 0x32, 0x55,
	0xe0, 0x52, 0xdc, 0x45, 0x51, 0x7c, 0x57, 0x14, 0x5a, 0x74, 0xbd, 0xdc, 0xf9, 0x01, 0xe4, 0x3c,
	0xff, 0x04, 0x89, 0x02, 0xcb, 0x25, 0xed, 0xe0, 0x24, 0x82, 0x90, 0x4a, 0x44, 0x58, 0x08, 0xc5,
	0x1b, 0x42, 0x3b, 0x80, 0xb2, 0x0e, 0xc5, 0x91, 0xb6, 0x8a, 0xe9, 0x87, 0x62, 0xca, 0x0f, 0x37,
	0x90, 0x6c, 0x0e, 0x2c, 0x31, 0xe2, 0x74, 0x18, 0xf7, 0x4e, 0x89, 0x45, 0xc4, 0xee, 0xf2, 0xc2,
	0x23, 0x88, 0xff, 0x26, 0x22, 0x2e, 0xf6, 0x3c, 0x36, 0x7b, 0x16, 0x03, 0xa4, 0xf7, 0xe1, 0x5a,
	0xd2, 0x26, 0x0a, 0xf8, 0xa7, 0x02, 0x5e, 0x1e, 0x6b, 0x27, 0xb9, 0x9f, 0x42, 0xb1, 0xee, 0xf9,
	0x36, 0xbb, 0x90, 0x1a, 0x92, 0x0e, 0x13, 0x44, 0x4f, 0x4b, 0xea, 0x42, 0x8c, 0xda, 0xf3, 0x03,
	0xd2, 0x20, 0x4c, 0x61, 0x17, 0x94, 0x45, 0x34, 0x65, 0xe4, 0xd4, 0xb6, 0xab, 0xed, 0x41, 0x69,
	0xfc, 0x0e, 0x0a, 0xf9, 0x97, 0x42, 0xbe, 0x3a, 0xc6, 0x4a, 0x02, 0x63, 0xb1, 0x34, 0x48, 0xa0,
	0xc4, 0xdc, 0xc2, 0xca, 0xb2, 0xa2, 0xb9, 0xcd, 0xf5, 0xcc, 0x50, 0xe0, 0x3a, 0xa5, 0x4d, 0x62,
	0xfb, 0x0a, 0x38, 0x87, 0x16, 0xf2, 0x89, 0x6f, 0x51, 0x76, 0x30, 0x90, 0xef, 0x5c, 0xdb, 0x81,
	0xd2, 0x78, 0x7f, 0x85, 0xfb, 0x5b, 0xe1, 0x2e, 0x8e, 0x34, 0x12, 0xb0, 0x2b, 0x5f, 0xa6, 0x61,
	0xda, 0x24, 0xbc, 0x4d, 0x7d, 0x4e, 0xb4, 0x9b, 0x90, 0x96, 0x6d, 0x24, 0x5e, 0xcd, 0x61, 0x6b,
	0x52, 0x2d, 0x66, 0x53, 0xfc, 0x37, 0xd5, 0x42, 0xed, 0x11, 0xcc, 0x8b, 0x06, 0x62, 0x0d, 0x74,
	0x10, 0x2c, 0xb5, 0x14, 0x8a, 0x8d, 0x98, 0x38, 0xde, 0x67, 0xf6, 0x70, 0xbc, 0xdd, 0x1f, 0x9b,
	0x73, 0xad, 0xe8, 0x0b, 0xed, 0x36, 0x64, 0xc3, 0xc6, 0x85, 0x95, 0x23, 0x1c, 0x8b, 0x17, 0x1c,
	0x55, 0x5b, 0xdb, 0x53, 0xbf, 0x66, 0x77, 0xb9, 0x76, 0x17, 0x52, 0x8c, 0x9e, 0x61, 0x32, 0x0b,
	0x55, 0xd9, 0x48, 0xec, 0xaf, 0x46, 0x37, 0x00, 0x86, 0x49, 0xcf, 0x4c, 0xa1, 0xcd, 0x7f, 0xce,
	0x40, 0x0a, 0x07, 0xda, 0x02, 0x64, 0x70, 0x28, 0x92, 0xec, 0x79, 0x0d, 0x63, 0x92, 0x36, 0xd3,
	0x38, 0xc4, 0x94, 0xd9, 0x82, 0x79, 0x1f, 0xd5, 0x58, 0x8e, 0xdc, 0x61, 0x9e, 0x4c, 0x58, 0xfd,
	0x45, 0x2d, 0xb1, 0x7a, 0xe6, 0x84, 0x68, 0xa3, 0xaf, 0xd1, 0xaa, 0x90, 0x91, 0xcb, 0x2a, 0xfa,
	0xcb, 0x64, 0x75, 0xb8, 0xb4, 0x27, 0xaa, 0xea, 0xaf, 0x26, 0x15, 0x55, 0x7b, 0xa2, 0x55, 0xfd,
	0xf5, 0xa4, 0xa2, 0x55, 0x6d, 0x0d, 0x66, 0xbb, 0x95, 0x21, 0x4e, 0x8b, 0x61, 0x78, 0x53, 0x4b,
	0x2e, 0xb6, 0x99, 0xb0, 0x42, 0x50, 0x82, 0xa1, 0xba, 0x03, 0x33, 0x01, 0x06, 0xdb, 0xf2, 0x43,
	0x87, 0xb7, 0x13, 0x38, 0x40, 0x10, 0x5e, 0x0f, 0xea, 0x6f, 0x41, 0xb6, 0x2b, 0x7d, 0x37, 0x81,
	0x34, 0xe3, 0x2b, 0xd9, 0x13, 0xb8, 0xd2, 0x66, 0xc4, 0xc2, 0x3f, 0x97, 0x38, 0x98, 0x18, 0x94,
	0xf1, 0xc8, 0x6d, 0xbd, 0x4f, 0x8e, 0x42, 0x1e, 0xc5, 0xfb, 0x03, 0xfa, 0xc1, 0x8b, 0x3b, 0x86,
	0xc5, 0x0b, 0xf6, 0x91, 0x63, 0x7e, 0x98, 0x80, 0x35, 0x17, 0xdb, 0xe0, 0xb0, 0x7f, 0xea, 0x87,
	0x90, 0x1b, 0xed, 0xfd, 0x71, 0x02, 0xef, 0x85, 0xf6, 0x70, 0xe3, 0x23, 0xd0, 0x47, 0xc6, 0xe4,
	0x53, 0x72, 0x4c, 0x2e, 0xb7, 0x87, 0x07, 0x64, 0x6d, 0x07, 0x0a, 0x1e, 0x8d, 0xd5, 0x5a, 0xff,
	0xfb, 0xe7, 0xf1, 0x8d, 0xbf, 0xf9, 0x32, 0xaa, 0x67, 0xe4, 0x57, 0x48, 0xf5, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0xc9, 0x49, 0xc4, 0x50, 0x09, 0x00, 0x00,
}