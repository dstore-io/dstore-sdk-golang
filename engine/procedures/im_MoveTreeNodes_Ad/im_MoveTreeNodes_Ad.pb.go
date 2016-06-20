// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_MoveTreeNodes_Ad.proto
// DO NOT EDIT!

/*
Package im_MoveTreeNodes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_MoveTreeNodes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_MoveTreeNodes_Ad

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
	TargetTreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=target_tree_node_id,json=targetTreeNodeId" json:"target_tree_node_id,omitempty"`
	TargetTreeNodeIdNull bool                        `protobuf:"varint,1001,opt,name=target_tree_node_id_null,json=targetTreeNodeIdNull" json:"target_tree_node_id_null,omitempty"`
	Resort               *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=resort" json:"resort,omitempty"`
	ResortNull           bool                        `protobuf:"varint,1002,opt,name=resort_null,json=resortNull" json:"resort_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTargetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TargetTreeNodeId
	}
	return nil
}

func (m *Parameters) GetResort() *dstore_values.BooleanValue {
	if m != nil {
		return m.Resort
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_MoveTreeNodes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_MoveTreeNodes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_MoveTreeNodes_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/im_MoveTreeNodes_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0xaa, 0x13, 0x31,
	0x10, 0xa5, 0x77, 0xed, 0xbd, 0x97, 0xdc, 0x07, 0x4b, 0x2a, 0xb2, 0x6e, 0x51, 0x6a, 0x7d, 0xf1,
	0x29, 0x2b, 0x2d, 0xa2, 0xaf, 0x15, 0x7c, 0xa8, 0xd0, 0x22, 0xa1, 0x08, 0xfa, 0xb2, 0xa4, 0x66,
	0x5c, 0x16, 0x77, 0x33, 0x25, 0x49, 0xdb, 0xdf, 0xf0, 0x5f, 0xfc, 0x29, 0xf5, 0x2b, 0xcc, 0x6e,
	0xb2, 0xd4, 0x5d, 0x0b, 0xfa, 0xd2, 0x66, 0x72, 0xe6, 0x9c, 0x99, 0x9c, 0x99, 0x25, 0x73, 0x69,
	0x2c, 0x6a, 0x48, 0x41, 0xe5, 0x85, 0x82, 0x74, 0xaf, 0xf1, 0x33, 0xc8, 0x83, 0x06, 0x93, 0x16,
	0x55, 0xb6, 0xc6, 0x23, 0x6c, 0x35, 0xc0, 0x06, 0x25, 0x98, 0x6c, 0x29, 0x99, 0x83, 0x2d, 0xd2,
	0xa7, 0x9e, 0xc3, 0x3c, 0x87, 0x5d, 0x48, 0x4c, 0xc6, 0x41, 0xf6, 0x28, 0xca, 0x03, 0x18, 0xcf,
	0x4b, 0x1e, 0x75, 0x6b, 0x81, 0xd6, 0xa8, 0x03, 0x34, 0xe9, 0x42, 0x15, 0x18, 0x23, 0x72, 0x08,
	0xe0, 0xb3, 0x3e, 0x68, 0x45, 0xa1, 0xbe, 0xa0, 0xae, 0x84, 0x2d, 0x50, 0xf9, 0xa4, 0xd9, 0x8f,
	0x01, 0x21, 0xef, 0x85, 0x16, 0x0e, 0x05, 0x6d, 0xe8, 0x3b, 0x32, 0xb6, 0x42, 0xe7, 0x60, 0x33,
	0xeb, 0xfa, 0xca, 0x94, 0x6b, 0x2c, 0x2b, 0x64, 0x3c, 0x98, 0x0e, 0x9e, 0xdf, 0xcd, 0x27, 0x2c,
	0xbc, 0x20, 0xb4, 0x57, 0x28, 0x0b, 0x39, 0xe8, 0x0f, 0x75, 0xc4, 0x47, 0x9e, 0xd7, 0x3e, 0x67,
	0x25, 0xe9, 0x2b, 0x12, 0x5f, 0xd0, 0xca, 0xd4, 0xa1, 0x2c, 0xe3, 0x9f, 0x37, 0x4e, 0xf1, 0x96,
	0x3f, 0xe8, 0x93, 0x36, 0x0e, 0xa4, 0x0b, 0x72, 0xed, 0x9c, 0x44, 0x6d, 0xe3, 0xab, 0x8b, 0x75,
	0x77, 0x88, 0x25, 0x08, 0xe5, 0xeb, 0x86, 0x54, 0x3a, 0x25, 0x77, 0xfe, 0xe4, 0x0b, 0xfc, 0xf2,
	0x05, 0x88, 0xbf, 0xab, 0x65, 0x67, 0xdf, 0xaf, 0xc8, 0x2d, 0x07, 0xb3, 0x47, 0x65, 0x80, 0xbe,
	0x20, 0xc3, 0xc6, 0xc8, 0xf0, 0xb4, 0x84, 0x75, 0x87, 0xe3, 0x4d, 0x7e, 0x5b, 0xff, 0x72, 0x9f,
	0x48, 0x3f, 0x92, 0x51, 0x6d, 0x61, 0xf6, 0x87, 0x87, 0xae, 0xbf, 0xc8, 0x91, 0x59, 0x8f, 0xdc,
	0x77, 0x7a, 0xed, 0xe2, 0xd5, 0x39, 0xe6, 0xf7, 0xab, 0xee, 0x05, 0x7d, 0x4d, 0x6e, 0xc2, 0xe8,
	0xe2, 0xa8, 0x51, 0x7c, 0xf2, 0x97, 0xa2, 0x1f, 0xec, 0xda, 0xff, 0xf3, 0x36, 0x9d, 0x2e, 0x49,
	0xa4, 0xf1, 0x14, 0xdf, 0x6b, 0x58, 0x29, 0xfb, 0xe7, 0x86, 0xb1, 0xd6, 0x00, 0xc6, 0xf1, 0xc4,
	0x6b, 0x6e, 0xf2, 0x98, 0x44, 0xee, 0x4c, 0x1f, 0x3a, 0xd3, 0xf1, 0x54, 0x0f, 0xfb, 0xdb, 0xc6,
	0x59, 0x32, 0xe4, 0x43, 0x17, 0xae, 0xe4, 0x9b, 0x2d, 0x99, 0x14, 0xd8, 0x13, 0x3e, 0xaf, 0xfb,
	0xa7, 0x97, 0x39, 0x1a, 0xf9, 0xb5, 0xc5, 0xe5, 0x7f, 0x7e, 0x11, 0xbb, 0xeb, 0x66, 0xfb, 0x16,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x48, 0x37, 0xe1, 0x48, 0x03, 0x00, 0x00,
}