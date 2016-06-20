// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_DeleteTreeNodes_Ad.proto
// DO NOT EDIT!

/*
Package im_DeleteTreeNodes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_DeleteTreeNodes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_DeleteTreeNodes_Ad

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
	TreeNodeList      *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=tree_node_list,json=treeNodeList" json:"tree_node_list,omitempty"`
	TreeNodeListNull  bool                        `protobuf:"varint,1001,opt,name=tree_node_list_null,json=treeNodeListNull" json:"tree_node_list_null,omitempty"`
	CascadeClones     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=cascade_clones,json=cascadeClones" json:"cascade_clones,omitempty"`
	CascadeClonesNull bool                        `protobuf:"varint,1002,opt,name=cascade_clones_null,json=cascadeClonesNull" json:"cascade_clones_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeList() *dstore_values.StringValue {
	if m != nil {
		return m.TreeNodeList
	}
	return nil
}

func (m *Parameters) GetCascadeClones() *dstore_values.BooleanValue {
	if m != nil {
		return m.CascadeClones
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_DeleteTreeNodes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_DeleteTreeNodes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_DeleteTreeNodes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x1b, 0xfa, 0x23, 0x03, 0xa5, 0xb8, 0x12, 0x0a, 0xa9, 0x40, 0xa8, 0xbd, 0x70, 0x72,
	0xf8, 0x3b, 0x70, 0x84, 0x16, 0x0e, 0x95, 0x68, 0x85, 0x2c, 0x84, 0x04, 0x97, 0xc8, 0x4d, 0x66,
	0x2b, 0x4b, 0x89, 0x5d, 0xd9, 0xee, 0xf6, 0x35, 0xf6, 0x65, 0xf6, 0xa1, 0x76, 0xaf, 0xfb, 0x02,
	0xeb, 0xc4, 0xae, 0xda, 0x64, 0xf7, 0xb0, 0x7b, 0x69, 0x33, 0x99, 0xef, 0x67, 0xfc, 0x79, 0x82,
	0xbe, 0x64, 0xda, 0x48, 0x05, 0x31, 0x88, 0x0d, 0x17, 0x10, 0x6f, 0x95, 0x4c, 0x21, 0xdb, 0x29,
	0xd0, 0x31, 0x2f, 0x92, 0x1f, 0x90, 0x83, 0x81, 0x3f, 0x0a, 0x60, 0x25, 0x33, 0xd0, 0xc9, 0xf7,
	0x8c, 0x58, 0x80, 0x91, 0x78, 0xea, 0x58, 0xc4, 0xb1, 0xc8, 0xbd, 0xd0, 0x68, 0xe4, 0xa5, 0xcf,
	0x59, 0xbe, 0x03, 0xed, 0x98, 0xd1, 0xeb, 0xba, 0x1f, 0x28, 0x25, 0x95, 0x6f, 0x8d, 0xeb, 0xad,
	0x02, 0xb4, 0x66, 0x1b, 0xf0, 0xcd, 0x69, 0xb3, 0x69, 0x18, 0x17, 0x67, 0x52, 0x15, 0xcc, 0x70,
	0x29, 0x1c, 0x68, 0x72, 0xd3, 0x42, 0xe8, 0x37, 0x53, 0xcc, 0x76, 0x41, 0x69, 0xfc, 0x0d, 0x0d,
	0x8c, 0x1d, 0x28, 0x11, 0x76, 0xa2, 0x24, 0xe7, 0xda, 0x84, 0xad, 0x77, 0xad, 0xf7, 0x4f, 0x3f,
	0x45, 0xc4, 0x8f, 0xef, 0x27, 0xd3, 0x46, 0x71, 0xb1, 0xf9, 0x5b, 0x16, 0xf4, 0x99, 0xf1, 0x47,
	0xf8, 0x65, 0xf1, 0x98, 0xa0, 0x51, 0x5d, 0x21, 0x11, 0xbb, 0x3c, 0x0f, 0xaf, 0x7a, 0x56, 0xa7,
	0x4f, 0x87, 0xa7, 0xd8, 0x95, 0x6d, 0xe0, 0x19, 0x1a, 0xa4, 0x4c, 0xa7, 0xcc, 0xa2, 0xd3, 0x5c,
	0x0a, 0xd0, 0x61, 0xbb, 0x72, 0x1c, 0x37, 0x1c, 0xd7, 0x52, 0xe6, 0xc0, 0x84, 0xb3, 0x7c, 0xee,
	0x29, 0xf3, 0x8a, 0x81, 0x63, 0x34, 0xaa, 0x6b, 0x38, 0xcf, 0x6b, 0xe7, 0xf9, 0xb2, 0x06, 0x2e,
	0x4d, 0x27, 0x97, 0x6d, 0xd4, 0xa7, 0xa0, 0xb7, 0x52, 0x68, 0xc0, 0x1f, 0x50, 0xa7, 0xca, 0xb4,
	0x79, 0x54, 0x7f, 0x53, 0x2e, 0xef, 0x9f, 0xe5, 0x2f, 0x75, 0x40, 0xfc, 0x0f, 0x0d, 0xcb, 0x34,
	0x93, 0x93, 0x38, 0xed, 0xd4, 0x81, 0x25, 0x93, 0x06, 0xb9, 0x19, 0xfa, 0xd2, 0xd6, 0x8b, 0x63,
	0x4d, 0x5f, 0x14, 0xf5, 0x17, 0xf8, 0x2b, 0xea, 0xf9, 0x5b, 0x0c, 0x83, 0x4a, 0xf1, 0xed, 0x1d,
	0x45, 0x77, 0xc7, 0x4b, 0xf7, 0x4f, 0x0f, 0x70, 0x3c, 0x47, 0x81, 0x92, 0xfb, 0xf0, 0x49, 0xc5,
	0xfa, 0x48, 0x1e, 0xb0, 0x6e, 0xe4, 0x10, 0x01, 0xa1, 0x72, 0x4f, 0x4b, 0x76, 0xf4, 0x06, 0x05,
	0xf6, 0x19, 0xbf, 0x42, 0x5d, 0x5b, 0x25, 0x3c, 0x0b, 0x2f, 0x56, 0x36, 0x94, 0x0e, 0xed, 0xd8,
	0x72, 0x91, 0xcd, 0x96, 0x68, 0xcc, 0x65, 0x43, 0xfa, 0xb8, 0xff, 0xff, 0xc9, 0xe3, 0xbe, 0x8c,
	0x75, 0xb7, 0xda, 0xc1, 0xcf, 0xb7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x41, 0x8b, 0x10, 0x52,
	0x03, 0x00, 0x00,
}