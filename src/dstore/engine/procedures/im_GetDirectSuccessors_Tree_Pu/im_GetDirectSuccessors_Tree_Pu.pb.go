// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetDirectSuccessors_Tree_Pu.proto
// DO NOT EDIT!

/*
Package im_GetDirectSuccessors_Tree_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetDirectSuccessors_Tree_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetDirectSuccessors_Tree_Pu

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
	TreeNodeId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull               bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	LanguageId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull               bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	FilterByCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=filter_by_characteristic_id,json=filterByCharacteristicId" json:"filter_by_characteristic_id,omitempty"`
	FilterByCharacteristicIdNull bool                        `protobuf:"varint,1003,opt,name=filter_by_characteristic_id_null,json=filterByCharacteristicIdNull" json:"filter_by_characteristic_id_null,omitempty"`
	FilterByCharacValue          *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=filter_by_charac_value,json=filterByCharacValue" json:"filter_by_charac_value,omitempty"`
	FilterByCharacValueNull      bool                        `protobuf:"varint,1004,opt,name=filter_by_charac_value_null,json=filterByCharacValueNull" json:"filter_by_charac_value_null,omitempty"`
	ApplyFilterForAllNodes       *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=apply_filter_for_all_nodes,json=applyFilterForAllNodes" json:"apply_filter_for_all_nodes,omitempty"`
	ApplyFilterForAllNodesNull   bool                        `protobuf:"varint,1005,opt,name=apply_filter_for_all_nodes_null,json=applyFilterForAllNodesNull" json:"apply_filter_for_all_nodes_null,omitempty"`
	NegateFilterByParams         *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=negate_filter_by_params,json=negateFilterByParams" json:"negate_filter_by_params,omitempty"`
	NegateFilterByParamsNull     bool                        `protobuf:"varint,1006,opt,name=negate_filter_by_params_null,json=negateFilterByParamsNull" json:"negate_filter_by_params_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetFilterByCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilterByCharacteristicId
	}
	return nil
}

func (m *Parameters) GetFilterByCharacValue() *dstore_values.StringValue {
	if m != nil {
		return m.FilterByCharacValue
	}
	return nil
}

func (m *Parameters) GetApplyFilterForAllNodes() *dstore_values.BooleanValue {
	if m != nil {
		return m.ApplyFilterForAllNodes
	}
	return nil
}

func (m *Parameters) GetNegateFilterByParams() *dstore_values.BooleanValue {
	if m != nil {
		return m.NegateFilterByParams
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
	RowId           int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TreeLevel       *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=tree_level,json=treeLevel" json:"tree_level,omitempty"`
	NodeDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Predecessor     *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=predecessor" json:"predecessor,omitempty"`
	LevelNo         *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=level_no,json=levelNo" json:"level_no,omitempty"`
	HasNextSibling  *dstore_values.BooleanValue `protobuf:"bytes,10005,opt,name=has_next_sibling,json=hasNextSibling" json:"has_next_sibling,omitempty"`
	TreeNodeId      *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	HasSuccessors   *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=has_successors,json=hasSuccessors" json:"has_successors,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTreeLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeLevel
	}
	return nil
}

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetPredecessor() *dstore_values.IntegerValue {
	if m != nil {
		return m.Predecessor
	}
	return nil
}

func (m *Response_Row) GetLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelNo
	}
	return nil
}

func (m *Response_Row) GetHasNextSibling() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasNextSibling
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetHasSuccessors() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasSuccessors
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetDirectSuccessors_Tree_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetDirectSuccessors_Tree_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetDirectSuccessors_Tree_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xc7, 0xc3, 0xd3, 0xa7, 0x05, 0x0f, 0x0a, 0x38, 0x18, 0x58, 0x5b, 0xa2, 0x04, 0x6f, 0xd4,
	0x98, 0xc5, 0xe8, 0x05, 0x46, 0x45, 0x23, 0x60, 0x09, 0x09, 0xac, 0x64, 0x31, 0x18, 0xb9, 0x99,
	0x6c, 0xb7, 0x43, 0x99, 0x64, 0xba, 0xd3, 0xcc, 0x6c, 0x41, 0xbe, 0x82, 0x57, 0xbe, 0xfb, 0x99,
	0xfc, 0x26, 0xbe, 0x7f, 0x06, 0xcf, 0xcc, 0x2c, 0xf4, 0xc5, 0x42, 0x8b, 0x37, 0x6d, 0x66, 0xe7,
	0xfc, 0xff, 0xe7, 0x37, 0x3b, 0xe7, 0x9c, 0x85, 0xc5, 0xaa, 0x4e, 0xa5, 0x62, 0xf3, 0x2c, 0xa9,
	0xf1, 0x84, 0xcd, 0x37, 0x94, 0x8c, 0x59, 0xb5, 0xa9, 0x98, 0x9e, 0xe7, 0x75, 0xba, 0xca, 0xd2,
	0x15, 0xae, 0x58, 0x9c, 0x6e, 0x35, 0xe3, 0x98, 0x69, 0x2d, 0x95, 0xa6, 0xcf, 0x15, 0x63, 0x74,
	0xb3, 0xe9, 0x63, 0x64, 0x2a, 0xc9, 0x2d, 0x27, 0xf7, 0x9d, 0xdc, 0x3f, 0x5d, 0x53, 0x9c, 0xcc,
	0x92, 0xed, 0x47, 0xa2, 0xc9, 0xb4, 0xb3, 0x28, 0x5e, 0xee, 0x24, 0x60, 0x4a, 0x49, 0x95, 0x6d,
	0x95, 0x3a, 0xb7, 0xea, 0x68, 0x17, 0xd5, 0x58, 0xb6, 0x79, 0xad, 0x7b, 0x33, 0x8d, 0x78, 0xb2,
	0x2b, 0x55, 0x3d, 0x4a, 0xb9, 0x4c, 0x5c, 0xd0, 0xdc, 0x97, 0x02, 0xc0, 0x66, 0xa4, 0x22, 0xdc,
	0x65, 0x4a, 0x93, 0x45, 0x38, 0x9f, 0x1a, 0x96, 0x44, 0x56, 0x19, 0xe5, 0x55, 0x6f, 0x68, 0x76,
	0xe8, 0xfa, 0xe8, 0x9d, 0x92, 0x9f, 0x9d, 0x22, 0xe3, 0xe2, 0x49, 0xca, 0x6a, 0x4c, 0x6d, 0x9b,
	0x55, 0x08, 0x46, 0x10, 0x60, 0xfc, 0x5a, 0x95, 0xdc, 0x84, 0x8b, 0xed, 0x72, 0x9a, 0x34, 0x85,
	0xf0, 0xbe, 0x0d, 0xa3, 0xc9, 0x48, 0x38, 0xd6, 0x8a, 0x0b, 0xf0, 0x31, 0x79, 0x08, 0xa3, 0x22,
	0x4a, 0x6a, 0x4d, 0x04, 0x36, 0x99, 0xfe, 0x1b, 0x20, 0xd3, 0x51, 0x3c, 0x66, 0xba, 0x01, 0x13,
	0x6d, 0x6a, 0x97, 0xe8, 0x7b, 0x96, 0xa8, 0x15, 0x66, 0x13, 0xed, 0x40, 0x69, 0x97, 0x0b, 0x3c,
	0x1e, 0xad, 0x1c, 0xd2, 0x78, 0x0f, 0x0f, 0x1b, 0xe3, 0x82, 0xeb, 0x94, 0xc7, 0x26, 0x71, 0xae,
	0x7f, 0x62, 0xcf, 0xe9, 0x97, 0x0e, 0x97, 0x3b, 0xd4, 0x88, 0xb1, 0x0a, 0xb3, 0xa7, 0x78, 0x3b,
	0xac, 0x1f, 0x0e, 0x6b, 0xe6, 0x24, 0x13, 0x0b, 0xf9, 0x0c, 0xa6, 0xba, 0x8d, 0xa8, 0x45, 0xf1,
	0xfe, 0xb7, 0x7c, 0xc5, 0x2e, 0x3e, 0x9d, 0x2a, 0x9e, 0xd4, 0x1c, 0xde, 0x64, 0xa7, 0xb3, 0x7d,
	0x88, 0x37, 0x59, 0xea, 0x6d, 0xe8, 0xa0, 0x7e, 0x3a, 0xa8, 0xe9, 0x1e, 0x52, 0xcb, 0xf3, 0x02,
	0x8a, 0x51, 0xa3, 0x21, 0x0e, 0x69, 0x66, 0x82, 0x75, 0x43, 0x23, 0x21, 0xec, 0xcd, 0x6a, 0x2f,
	0xdf, 0xf3, 0x9d, 0x55, 0xa4, 0x14, 0x2c, 0x4a, 0x1c, 0xd4, 0x94, 0x95, 0x97, 0xad, 0xba, 0x2c,
	0xd5, 0x13, 0x21, 0xcc, 0xdd, 0x6b, 0xb2, 0x02, 0x57, 0x4f, 0x36, 0x76, 0x6c, 0xbf, 0x1c, 0x5b,
	0xb1, 0xb7, 0x83, 0xc5, 0x0b, 0x61, 0x3a, 0x61, 0xb5, 0x28, 0x65, 0xb4, 0x75, 0xc8, 0x86, 0x29,
	0x63, 0xed, 0x15, 0xfa, 0xb3, 0x5d, 0x72, 0xda, 0x72, 0x76, 0x76, 0x5b, 0xff, 0x9a, 0x3c, 0x86,
	0x99, 0x13, 0x3c, 0x1d, 0xd6, 0x6f, 0x87, 0xe5, 0xf5, 0x12, 0x1b, 0xa8, 0xb9, 0xd7, 0x05, 0x18,
	0x09, 0x99, 0x6e, 0xc8, 0x44, 0x33, 0x72, 0x1b, 0xf2, 0xb6, 0x53, 0xb3, 0x16, 0x3a, 0xbe, 0xbf,
	0x6c, 0x10, 0xb8, 0x2e, 0x7e, 0x6a, 0x7e, 0x43, 0x17, 0x48, 0x5e, 0xc2, 0x84, 0xe9, 0x51, 0xda,
	0xd6, 0xa4, 0xd8, 0x15, 0x39, 0x14, 0xfb, 0x5d, 0xe2, 0xee, 0x56, 0xde, 0xc0, 0xf5, 0x5a, 0x6b,
	0x1d, 0x8e, 0xd7, 0x3b, 0x1f, 0x90, 0x7b, 0x30, 0x9c, 0xcd, 0x06, 0x2c, 0x77, 0xe3, 0x78, 0xe5,
	0x2f, 0x47, 0x37, 0x39, 0x36, 0xdc, 0x7f, 0x78, 0x14, 0x4e, 0xd6, 0x21, 0xa7, 0xe4, 0x01, 0x16,
	0xa1, 0x51, 0xdd, 0xf7, 0xcf, 0x32, 0xcd, 0xfc, 0xa3, 0x77, 0xe1, 0x87, 0xf2, 0x20, 0x34, 0x36,
	0xc5, 0xaf, 0x39, 0xc8, 0xe1, 0x82, 0x4c, 0x41, 0x01, 0x97, 0xa6, 0xfb, 0xde, 0x04, 0xf8, 0x7a,
	0xf2, 0x61, 0x1e, 0x97, 0xd8, 0x4e, 0x0f, 0xc0, 0x4e, 0x13, 0x2a, 0xd8, 0x3e, 0x13, 0xde, 0xdb,
	0xa0, 0x7f, 0x6b, 0x9e, 0x33, 0xf1, 0xeb, 0x26, 0x9c, 0x94, 0x61, 0xc2, 0xce, 0x1d, 0xac, 0x91,
	0x58, 0xf1, 0x86, 0x7d, 0x7f, 0xef, 0x82, 0xbe, 0xdd, 0x33, 0x6e, 0x44, 0x2b, 0x2d, 0x0d, 0x79,
	0x04, 0xa3, 0x0d, 0xc5, 0xaa, 0xcc, 0x1d, 0xc7, 0x7b, 0x3f, 0x00, 0x45, 0xbb, 0x80, 0x2c, 0xc0,
	0x88, 0xe5, 0xc7, 0x92, 0xf6, 0x3e, 0x0c, 0x20, 0x1e, 0xb6, 0xd1, 0x81, 0x34, 0x07, 0xd8, 0x8b,
	0xb0, 0xd8, 0xd8, 0xab, 0x94, 0x6a, 0x5e, 0x11, 0x88, 0xe8, 0x7d, 0x0c, 0xfa, 0x97, 0xf3, 0x18,
	0xaa, 0x02, 0x14, 0x6d, 0x39, 0x0d, 0x1e, 0xa0, 0x73, 0x88, 0x7f, 0x0a, 0xce, 0x36, 0xc5, 0x97,
	0xc1, 0x38, 0x52, 0x7d, 0x7c, 0xa3, 0xde, 0xe7, 0x01, 0x28, 0x2e, 0xa0, 0xa6, 0x55, 0x04, 0x4b,
	0xdb, 0x50, 0xe2, 0xb2, 0xab, 0x5e, 0x5a, 0x1f, 0xcf, 0x9d, 0x85, 0x7f, 0xfc, 0xac, 0x56, 0x0a,
	0xf6, 0xbb, 0x75, 0xf7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x49, 0x49, 0xda, 0x98, 0x07,
	0x00, 0x00,
}