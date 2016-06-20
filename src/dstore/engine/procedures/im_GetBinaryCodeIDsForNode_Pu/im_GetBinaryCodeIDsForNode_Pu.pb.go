// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetBinaryCodeIDsForNode_Pu.proto
// DO NOT EDIT!

/*
Package im_GetBinaryCodeIDsForNode_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetBinaryCodeIDsForNode_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetBinaryCodeIDsForNode_Pu

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
	TreeNodeId                 *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull             bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	BinaryPropertyValueId      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=binary_property_value_id,json=binaryPropertyValueId" json:"binary_property_value_id,omitempty"`
	BinaryPropertyValueIdNull  bool                        `protobuf:"varint,1002,opt,name=binary_property_value_id_null,json=binaryPropertyValueIdNull" json:"binary_property_value_id_null,omitempty"`
	IsNodeId                   *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=is_node_id,json=isNodeId" json:"is_node_id,omitempty"`
	IsNodeIdNull               bool                        `protobuf:"varint,1003,opt,name=is_node_id_null,json=isNodeIdNull" json:"is_node_id_null,omitempty"`
	BinaryCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=binary_characteristic_id,json=binaryCharacteristicId" json:"binary_characteristic_id,omitempty"`
	BinaryCharacteristicIdNull bool                        `protobuf:"varint,1004,opt,name=binary_characteristic_id_null,json=binaryCharacteristicIdNull" json:"binary_characteristic_id_null,omitempty"`
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

func (m *Parameters) GetBinaryPropertyValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryPropertyValueId
	}
	return nil
}

func (m *Parameters) GetIsNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsNodeId
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicId
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
	RowId                 int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BinaryCodeId          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	TreeNodeId            *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	BinaryPropertyValueId *dstore_values.IntegerValue `protobuf:"bytes,20001,opt,name=binary_property_value_id,json=binaryPropertyValueId" json:"binary_property_value_id,omitempty"`
	BinaryPropertyValue   *dstore_values.StringValue  `protobuf:"bytes,20005,opt,name=binary_property_value,json=binaryPropertyValue" json:"binary_property_value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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

func (m *Response_Row) GetBinaryPropertyValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryPropertyValueId
	}
	return nil
}

func (m *Response_Row) GetBinaryPropertyValue() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryPropertyValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetBinaryCodeIDsForNode_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetBinaryCodeIDsForNode_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetBinaryCodeIDsForNode_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0x26, 0xa6, 0xb9, 0x70, 0x0c, 0xad, 0x4e, 0x69, 0x49, 0x13, 0x14, 0xa9, 0x20, 0xa2, 0xb0,
	0x11, 0x45, 0xb1, 0xa0, 0x82, 0xa9, 0x17, 0x82, 0x34, 0x84, 0xc5, 0x16, 0xf4, 0x65, 0xd9, 0x64,
	0x8f, 0x71, 0x20, 0x99, 0x09, 0x33, 0x13, 0x4b, 0x5f, 0xfd, 0x01, 0xe2, 0xe5, 0xb5, 0x2f, 0x22,
	0xfe, 0x38, 0x2f, 0x3f, 0xc2, 0x99, 0x9d, 0x49, 0xb6, 0xbb, 0x5d, 0x9b, 0xb6, 0x2f, 0x6d, 0xce,
	0xce, 0xf9, 0xbe, 0xf3, 0x9d, 0x73, 0xbe, 0x19, 0x78, 0x14, 0x49, 0xc5, 0x05, 0xb6, 0x90, 0x0d,
	0x29, 0xc3, 0xd6, 0x44, 0xf0, 0x01, 0x46, 0x53, 0x81, 0xb2, 0x45, 0xc7, 0xc1, 0x4b, 0x54, 0x6d,
	0xca, 0x42, 0x71, 0xb0, 0xcd, 0x23, 0xec, 0x3c, 0x93, 0x2f, 0xb8, 0xe8, 0xea, 0x5f, 0x41, 0x6f,
	0xea, 0xe9, 0x44, 0xc5, 0xc9, 0x6d, 0x8b, 0xf6, 0x2c, 0xda, 0x3b, 0x11, 0xd2, 0x58, 0x75, 0xa5,
	0x3e, 0x84, 0xa3, 0x29, 0x4a, 0xcb, 0xd0, 0xd8, 0x48, 0xd7, 0x47, 0x21, 0xb8, 0x70, 0x47, 0xcd,
	0xf4, 0xd1, 0x18, 0xa5, 0x0c, 0x87, 0xe8, 0x0e, 0xaf, 0x67, 0x0f, 0x55, 0x48, 0xd9, 0x3b, 0x2e,
	0xc6, 0xa1, 0xa2, 0x9c, 0xd9, 0xa4, 0xcd, 0x4f, 0x4b, 0x00, 0xbd, 0x50, 0x84, 0xfa, 0x14, 0x85,
	0x24, 0x8f, 0xa1, 0xa6, 0x04, 0x62, 0xc0, 0x8c, 0x20, 0x1a, 0xd5, 0x0b, 0xd7, 0x0a, 0x37, 0x2f,
	0xde, 0x6d, 0x7a, 0xae, 0x09, 0xa7, 0x8b, 0x32, 0x85, 0x43, 0x14, 0x7b, 0x26, 0xf2, 0xc1, 0x00,
	0x4c, 0x03, 0x9d, 0x88, 0xdc, 0x82, 0xcb, 0x47, 0xe1, 0x01, 0x9b, 0x8e, 0x46, 0xf5, 0x5f, 0x15,
	0x4d, 0x52, 0xf5, 0x97, 0x93, 0xbc, 0xae, 0xfe, 0x4c, 0x5e, 0x43, 0xbd, 0x1f, 0x8f, 0x21, 0xd0,
	0x4a, 0x26, 0x28, 0xd4, 0x41, 0x10, 0xd3, 0x9b, 0xb2, 0x17, 0x16, 0x97, 0x5d, 0xb3, 0xe0, 0x9e,
	0xc3, 0xc6, 0x1f, 0xb5, 0x82, 0xa7, 0x70, 0xe5, 0x7f, 0xac, 0x56, 0xcd, 0x6f, 0xab, 0x66, 0x23,
	0x17, 0x1e, 0x0b, 0xdb, 0x02, 0xa0, 0x72, 0x3e, 0x81, 0x62, 0xae, 0x94, 0x3e, 0xe7, 0x23, 0x0c,
	0x99, 0x95, 0x52, 0xa5, 0xd2, 0xf5, 0x7f, 0x03, 0x56, 0x12, 0xa8, 0xad, 0xf7, 0xc7, 0xd6, 0xab,
	0xcd, 0x72, 0xe2, 0x12, 0xbb, 0xf3, 0xde, 0x07, 0xef, 0xf5, 0xf0, 0x07, 0x7a, 0xf6, 0x54, 0x2a,
	0x3a, 0x30, 0x05, 0x97, 0x16, 0xf7, 0xbe, 0x6e, 0xc1, 0xdb, 0x29, 0xac, 0x2e, 0xdf, 0x9e, 0x37,
	0x7f, 0x8c, 0xd6, 0x8a, 0xf9, 0x6b, 0xc5, 0x34, 0xf2, 0xf1, 0x46, 0xda, 0xe6, 0x8f, 0x12, 0x54,
	0x7d, 0x94, 0x13, 0xce, 0x24, 0x92, 0x3b, 0x50, 0x8a, 0xed, 0xe6, 0x7c, 0xd0, 0xf0, 0xd2, 0x66,
	0xb6, 0x56, 0x7c, 0x6e, 0xfe, 0xfa, 0x36, 0x91, 0xbc, 0x81, 0x4b, 0xc6, 0x68, 0xc1, 0x11, 0xa7,
	0xe9, 0x6d, 0x16, 0x35, 0xd8, 0xcb, 0x80, 0xb3, 0x7e, 0xdc, 0xd1, 0x71, 0x27, 0x89, 0xfd, 0x95,
	0x71, 0xfa, 0x03, 0x79, 0x08, 0x15, 0x67, 0x70, 0xbd, 0x14, 0xc3, 0x78, 0xf5, 0x18, 0xa3, 0xb5,
	0xff, 0x8e, 0xfd, 0xef, 0xcf, 0xd2, 0xc9, 0x2b, 0x28, 0x0a, 0xbe, 0xaf, 0x27, 0x6b, 0x50, 0x5b,
	0xde, 0x19, 0x6e, 0xa4, 0x37, 0x1b, 0x85, 0xe7, 0xf3, 0x7d, 0xdf, 0xb0, 0x34, 0x3e, 0x16, 0xa1,
	0xa8, 0x03, 0xb2, 0x0e, 0x65, 0x1d, 0x9a, 0x8d, 0x7d, 0xee, 0xea, 0xe9, 0x94, 0xfc, 0x92, 0x0e,
	0xe3, 0x25, 0x2c, 0xcf, 0x96, 0xe0, 0x2c, 0xf4, 0xa5, 0xbb, 0x78, 0xa5, 0xb5, 0x7e, 0x22, 0x20,
	0x22, 0x4f, 0x32, 0xd7, 0xf0, 0x6b, 0xf7, 0x6c, 0xf7, 0xf0, 0x3e, 0x54, 0x66, 0xd0, 0x6f, 0xa7,
	0x80, 0x96, 0x99, 0x85, 0xed, 0x9d, 0x70, 0x25, 0xbf, 0x1f, 0x16, 0xce, 0x7d, 0x29, 0x7b, 0xb0,
	0x96, 0xcb, 0x5b, 0xff, 0x79, 0x98, 0x31, 0x96, 0x23, 0x95, 0x4a, 0x50, 0x36, 0xb4, 0x9c, 0xab,
	0x39, 0x9c, 0xed, 0x5d, 0x68, 0x52, 0x9e, 0x59, 0x64, 0xf2, 0x30, 0xbf, 0x7d, 0x70, 0xbe, 0x27,
	0xbb, 0x5f, 0x8e, 0x1f, 0xc5, 0x7b, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xf0, 0xf5, 0xab,
	0xf3, 0x05, 0x00, 0x00,
}
