// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetNewCharacsForNode_Ad.proto
// DO NOT EDIT!

/*
Package im_GetNewCharacsForNode_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetNewCharacsForNode_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetNewCharacsForNode_Ad

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
	NodeId                       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	NodeIdNull                   bool                        `protobuf:"varint,1001,opt,name=node_id_null,json=nodeIdNull" json:"node_id_null,omitempty"`
	OrderByCharacDescription     *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=order_by_charac_description,json=orderByCharacDescription" json:"order_by_charac_description,omitempty"`
	OrderByCharacDescriptionNull bool                        `protobuf:"varint,1002,opt,name=order_by_charac_description_null,json=orderByCharacDescriptionNull" json:"order_by_charac_description_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Parameters) GetOrderByCharacDescription() *dstore_values.BooleanValue {
	if m != nil {
		return m.OrderByCharacDescription
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CharacteristicDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	NodeCharacteristicId      *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetNewCharacsForNode_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetNewCharacsForNode_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetNewCharacsForNode_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetNewCharacsForNode_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x55, 0xba, 0x24, 0xa9, 0x06, 0x24, 0x90, 0x41, 0xd5, 0x76, 0x83, 0x50, 0x29, 0x97, 0x9e,
	0x1c, 0x04, 0x08, 0x81, 0xc4, 0x85, 0x16, 0x28, 0x39, 0x74, 0x85, 0x7c, 0x40, 0x22, 0x1c, 0x56,
	0x9b, 0x78, 0x58, 0x2c, 0x12, 0x3b, 0xb2, 0x37, 0x44, 0xfd, 0x05, 0x4e, 0xc0, 0x7f, 0xf1, 0x21,
	0xc0, 0x4f, 0x60, 0xaf, 0x1d, 0xa5, 0x5e, 0x28, 0xd0, 0xcb, 0xae, 0xc6, 0x33, 0x6f, 0xe6, 0xf9,
	0xbd, 0x31, 0x3c, 0xe6, 0xa6, 0x56, 0x1a, 0x87, 0x28, 0x2b, 0x21, 0x71, 0xb8, 0xd0, 0x6a, 0x8a,
	0x7c, 0xa9, 0xd1, 0x0c, 0xc5, 0xbc, 0x38, 0xc6, 0x3a, 0xc7, 0xd5, 0xd1, 0xfb, 0x52, 0x97, 0x53,
	0xf3, 0x42, 0xe9, 0x5c, 0x71, 0x2c, 0x9e, 0x72, 0x6a, 0xab, 0x6a, 0x45, 0x0e, 0x3c, 0x94, 0x7a,
	0x28, 0x3d, 0xbf, 0x3e, 0xbb, 0x1e, 0x86, 0x7c, 0x2c, 0x67, 0x4b, 0x34, 0x1e, 0x9e, 0xed, 0xc6,
	0x93, 0x51, 0x6b, 0xa5, 0x43, 0x6a, 0x10, 0xa7, 0xe6, 0x68, 0x4c, 0x59, 0x61, 0x48, 0xde, 0x69,
	0x27, 0xeb, 0x52, 0xc8, 0x77, 0x4a, 0xcf, 0xcb, 0x5a, 0x28, 0xe9, 0x8b, 0xf6, 0x3f, 0x6d, 0x01,
	0xbc, 0xb2, 0x34, 0x6c, 0x16, 0xb5, 0x21, 0x0f, 0xa0, 0x2f, 0x1d, 0x17, 0xc1, 0xd3, 0xce, 0x5e,
	0xe7, 0xe0, 0xf2, 0xbd, 0x01, 0x0d, 0xe4, 0x03, 0x25, 0x21, 0x6b, 0xac, 0x50, 0xbf, 0x76, 0x11,
	0xeb, 0xb9, 0xda, 0x11, 0x27, 0xb7, 0xe1, 0x4a, 0x40, 0x15, 0x72, 0x39, 0x9b, 0xa5, 0xdf, 0xfb,
	0x16, 0xbb, 0xcd, 0xc0, 0xa7, 0x73, 0x7b, 0x44, 0xc6, 0x30, 0x50, 0x9a, 0xa3, 0x2e, 0x26, 0xa7,
	0xc5, 0xb4, 0xb9, 0x77, 0xc1, 0xd1, 0x4c, 0xb5, 0x58, 0x38, 0x32, 0xe9, 0xd6, 0x1f, 0x87, 0x4d,
	0x94, 0x9a, 0x61, 0x29, 0xfd, 0xb0, 0xb4, 0xc1, 0x1f, 0x9e, 0x7a, 0xd5, 0x9e, 0x6d, 0xc0, 0xe4,
	0x18, 0xf6, 0xfe, 0xd2, 0xdb, 0x53, 0xfa, 0xe1, 0x29, 0xdd, 0x3c, 0xaf, 0x89, 0x23, 0xb9, 0xff,
	0x33, 0x81, 0x6d, 0x86, 0x66, 0xa1, 0xa4, 0x41, 0x72, 0x17, 0xba, 0x8d, 0xd4, 0x41, 0x88, 0x8c,
	0xc6, 0x2e, 0x7a, 0x1b, 0x9e, 0xbb, 0x2f, 0xf3, 0x85, 0xe4, 0x0d, 0x5c, 0x73, 0x22, 0x17, 0x67,
	0x54, 0xb6, 0x17, 0x4b, 0x2c, 0x98, 0xb6, 0xc0, 0x6d, 0x2f, 0x4e, 0x6c, 0x3c, 0xda, 0xc4, 0xec,
	0xea, 0x3c, 0x3e, 0x20, 0x8f, 0xa0, 0x1f, 0xcc, 0x4d, 0x93, 0xa6, 0xe3, 0xad, 0xdf, 0x3a, 0x7a,
	0xeb, 0x4f, 0xfc, 0x9f, 0xad, 0xcb, 0xc9, 0x4b, 0x48, 0xb4, 0x5a, 0xa5, 0x97, 0x1a, 0xd4, 0x43,
	0xfa, 0xbf, 0xab, 0x48, 0xd7, 0x3a, 0x50, 0xa6, 0x56, 0xcc, 0xb5, 0xc8, 0xbe, 0x75, 0x20, 0xb1,
	0x01, 0xd9, 0x81, 0x9e, 0x0d, 0xdd, 0x8a, 0x7c, 0xce, 0xad, 0x34, 0x5d, 0xd6, 0xb5, 0xa1, 0xdd,
	0x82, 0x31, 0x64, 0x5e, 0x7d, 0xbb, 0x49, 0xc2, 0xd4, 0x22, 0x76, 0xf8, 0x4b, 0x1e, 0xcb, 0x18,
	0x2c, 0x36, 0xb5, 0x16, 0xb2, 0xf2, 0x0e, 0xef, 0xc6, 0xf0, 0xb3, 0x16, 0x33, 0xd8, 0x69, 0x36,
	0xac, 0x35, 0xc0, 0x72, 0xf8, 0x9a, 0xff, 0x7b, 0x4f, 0x6f, 0x38, 0xec, 0x51, 0x04, 0x1d, 0xf1,
	0xc3, 0xb7, 0x30, 0x10, 0xaa, 0x25, 0xc8, 0xe6, 0x59, 0x8f, 0x9f, 0x54, 0xca, 0xf0, 0x0f, 0xeb,
	0x3c, 0xbf, 0xd8, 0xcb, 0x9f, 0xf4, 0x9a, 0xe7, 0x75, 0xff, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0xad, 0x40, 0xf3, 0x37, 0x04, 0x00, 0x00,
}