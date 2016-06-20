// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetNodeDescriptions_Pu.proto
// DO NOT EDIT!

/*
Package im_GetNodeDescriptions_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetNodeDescriptions_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetNodeDescriptions_Pu

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
	NodeIds                        *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=node_ids,json=nodeIds" json:"node_ids,omitempty"`
	NodeIdsNull                    bool                        `protobuf:"varint,1001,opt,name=node_ids_null,json=nodeIdsNull" json:"node_ids_null,omitempty"`
	IsTreeNodeId                   *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=is_tree_node_id,json=isTreeNodeId" json:"is_tree_node_id,omitempty"`
	IsTreeNodeIdNull               bool                        `protobuf:"varint,1002,opt,name=is_tree_node_id_null,json=isTreeNodeIdNull" json:"is_tree_node_id_null,omitempty"`
	LanguageId                     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                 bool                        `protobuf:"varint,1003,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	LookForProductDescriptions     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=look_for_product_descriptions,json=lookForProductDescriptions" json:"look_for_product_descriptions,omitempty"`
	LookForProductDescriptionsNull bool                        `protobuf:"varint,1004,opt,name=look_for_product_descriptions_null,json=lookForProductDescriptionsNull" json:"look_for_product_descriptions_null,omitempty"`
	StoreTreeNodeIdsInResult       *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=store_tree_node_ids_in_result,json=storeTreeNodeIdsInResult" json:"store_tree_node_ids_in_result,omitempty"`
	StoreTreeNodeIdsInResultNull   bool                        `protobuf:"varint,1005,opt,name=store_tree_node_ids_in_result_null,json=storeTreeNodeIdsInResultNull" json:"store_tree_node_ids_in_result_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *Parameters) GetIsTreeNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsTreeNodeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetLookForProductDescriptions() *dstore_values.BooleanValue {
	if m != nil {
		return m.LookForProductDescriptions
	}
	return nil
}

func (m *Parameters) GetStoreTreeNodeIdsInResult() *dstore_values.BooleanValue {
	if m != nil {
		return m.StoreTreeNodeIdsInResult
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
	RowId       int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Description *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=description" json:"description,omitempty"`
	NodeId      *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId  *dstore_values.IntegerValue `protobuf:"bytes,30002,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetNodeDescriptions_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetNodeDescriptions_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetNodeDescriptions_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x26, 0xc6, 0xfc, 0x70, 0x52, 0x6d, 0x18, 0x45, 0xd6, 0xc4, 0x16, 0x49, 0x2f, 0xd4, 0x9b,
	0x8d, 0x54, 0x0a, 0xbd, 0x50, 0x84, 0xa2, 0x96, 0x20, 0x5d, 0xe2, 0x22, 0x82, 0x22, 0x0e, 0xdb,
	0xec, 0x18, 0x06, 0x37, 0x33, 0x61, 0x66, 0xd6, 0xbe, 0x86, 0x7a, 0xed, 0x53, 0xf8, 0x2c, 0xbe,
	0x84, 0x55, 0x2f, 0x7c, 0x03, 0xcf, 0xec, 0x6c, 0xcc, 0xee, 0x6a, 0x9a, 0x7a, 0xd3, 0x72, 0xf6,
	0x7c, 0xdf, 0xf9, 0xbe, 0x9e, 0x7e, 0x67, 0x60, 0x3f, 0xd6, 0x46, 0x2a, 0x36, 0x64, 0x62, 0xca,
	0x05, 0x1b, 0xce, 0x95, 0x9c, 0xb0, 0x38, 0x55, 0x4c, 0x0f, 0xf9, 0x8c, 0x1e, 0x32, 0x13, 0xc8,
	0x98, 0x3d, 0x62, 0x7a, 0xa2, 0xf8, 0xdc, 0x70, 0x29, 0x34, 0x1d, 0xa7, 0x3e, 0x82, 0x8c, 0x24,
	0xb7, 0x1c, 0xd3, 0x77, 0x4c, 0x7f, 0x25, 0xbc, 0x77, 0x25, 0x97, 0x78, 0x1f, 0x25, 0x29, 0xd3,
	0x8e, 0xdd, 0xbb, 0x5e, 0xd6, 0x65, 0x4a, 0x49, 0x95, 0xb7, 0xfa, 0xe5, 0xd6, 0x8c, 0x69, 0x1d,
	0x4d, 0x59, 0xde, 0xdc, 0xa9, 0x36, 0x4d, 0xc4, 0xc5, 0x5b, 0xa9, 0x66, 0x91, 0xd5, 0x73, 0xa0,
	0xc1, 0xe7, 0x06, 0xc0, 0x38, 0x52, 0x11, 0x76, 0x99, 0xd2, 0x64, 0x0f, 0xda, 0x02, 0x7d, 0x51,
	0x1e, 0x6b, 0xaf, 0x76, 0xb3, 0x76, 0xbb, 0xb3, 0xdb, 0xf3, 0x73, 0xf3, 0xb9, 0x27, 0x6d, 0x14,
	0x17, 0xd3, 0x17, 0xb6, 0x08, 0x5b, 0x16, 0x3b, 0x8a, 0x35, 0xd9, 0x81, 0x4b, 0x0b, 0x1a, 0x15,
	0x69, 0x92, 0x78, 0xdf, 0x5a, 0x48, 0x6e, 0x87, 0x9d, 0x1c, 0x10, 0xe0, 0x37, 0x72, 0x00, 0x9b,
	0x5c, 0x53, 0xa3, 0x18, 0xa3, 0x39, 0xd8, 0xbb, 0x90, 0x49, 0xf4, 0x2b, 0x12, 0xc7, 0x52, 0x26,
	0x2c, 0x12, 0x4e, 0x63, 0x83, 0xeb, 0xe7, 0x48, 0x09, 0xb2, 0x41, 0x64, 0x08, 0x57, 0x2b, 0x33,
	0x9c, 0xde, 0xa9, 0xd3, 0xeb, 0x16, 0xc1, 0x99, 0xe8, 0x7d, 0xe8, 0x24, 0x91, 0x98, 0xa6, 0xb8,
	0x16, 0x2b, 0x58, 0xff, 0xa7, 0x20, 0x17, 0x86, 0x4d, 0x99, 0x72, 0x82, 0xb0, 0xc0, 0xa3, 0xdc,
	0x1d, 0xe8, 0x16, 0xd8, 0x4e, 0xea, 0xbb, 0x93, 0xba, 0xbc, 0x84, 0x65, 0x42, 0x6f, 0x60, 0x2b,
	0x91, 0xf2, 0x1d, 0xc5, 0x05, 0x53, 0x5c, 0x6d, 0x9c, 0x4e, 0x0c, 0x8d, 0x0b, 0xff, 0x5e, 0xef,
	0xe2, 0xfa, 0xbf, 0xb5, 0x67, 0x27, 0x3c, 0x91, 0x6a, 0xec, 0xf8, 0xc5, 0x74, 0x90, 0xa7, 0x30,
	0x38, 0x73, 0xbe, 0x33, 0xf7, 0xc3, 0x99, 0xdb, 0x5e, 0x3d, 0x28, 0x33, 0xfb, 0x1a, 0xb6, 0x32,
	0x17, 0xa5, 0x4d, 0x6a, 0xca, 0x05, 0xc5, 0x38, 0xa7, 0x89, 0xf1, 0x1a, 0xeb, 0xcd, 0x7a, 0x59,
	0x6b, 0xb9, 0x6e, 0x3d, 0x12, 0x61, 0x46, 0x26, 0x23, 0x18, 0x9c, 0x39, 0xdd, 0x59, 0xfd, 0xe9,
	0xac, 0xde, 0x58, 0x35, 0xc6, 0x1a, 0x1d, 0xfc, 0xaa, 0x43, 0x1b, 0xcb, 0x39, 0x1a, 0x67, 0xe4,
	0x2e, 0x34, 0xb2, 0xf0, 0x57, 0x93, 0x99, 0x9f, 0x95, 0x3b, 0x8c, 0xc7, 0xf6, 0x67, 0xe8, 0x80,
	0xe4, 0x25, 0x74, 0x6d, 0xec, 0x69, 0x21, 0xf7, 0x98, 0xb9, 0x3a, 0x92, 0xfd, 0x0a, 0xb9, 0x7a,
	0x1d, 0x47, 0x58, 0x8f, 0x96, 0x75, 0xb8, 0x39, 0x2b, 0x7f, 0x20, 0xfb, 0xd0, 0xca, 0xcf, 0x0d,
	0x43, 0x65, 0x27, 0x6e, 0xff, 0x35, 0xd1, 0x1d, 0xe3, 0x91, 0xfb, 0x1d, 0x2e, 0xe0, 0xe4, 0x10,
	0xea, 0x4a, 0x9e, 0x60, 0x1e, 0x2c, 0x6b, 0xcf, 0x3f, 0xe7, 0xdb, 0xe0, 0x2f, 0xd6, 0xe0, 0x87,
	0xf2, 0x24, 0xb4, 0x13, 0x7a, 0x5f, 0x6b, 0x50, 0xc7, 0x82, 0x5c, 0x83, 0x26, 0x96, 0x36, 0xde,
	0x1f, 0x02, 0xdc, 0x4c, 0x23, 0x6c, 0x60, 0x89, 0xe9, 0x7d, 0x00, 0x9d, 0x42, 0x42, 0xbc, 0x8f,
	0xc1, 0xda, 0x83, 0x2e, 0xe2, 0xf1, 0x2d, 0x68, 0x2d, 0xee, 0xf4, 0x53, 0xb0, 0xfe, 0x6e, 0x9a,
	0xee, 0xd6, 0xc9, 0x43, 0xd8, 0x28, 0xdd, 0xf8, 0x97, 0xd3, 0xda, 0x39, 0x8e, 0xce, 0xfc, 0x09,
	0xc0, 0xc1, 0x33, 0xe8, 0x73, 0x59, 0x59, 0xcb, 0xf2, 0xb1, 0x7d, 0xb5, 0xfb, 0xff, 0xcf, 0xf0,
	0x71, 0x33, 0x7b, 0xec, 0xee, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x29, 0xa7, 0x0b, 0x49, 0xc3,
	0x05, 0x00, 0x00,
}