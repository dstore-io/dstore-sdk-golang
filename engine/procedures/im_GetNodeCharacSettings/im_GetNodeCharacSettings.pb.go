// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetNodeCharacSettings.proto
// DO NOT EDIT!

/*
Package im_GetNodeCharacSettings is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetNodeCharacSettings.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetNodeCharacSettings

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
	NodeCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull bool                        `protobuf:"varint,1001,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	KeyVariable              *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull          bool                        `protobuf:"varint,1002,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
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
	Value                     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value" json:"value,omitempty"`
	KeyVariable               *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
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

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetNodeCharacSettings.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetNodeCharacSettings.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetNodeCharacSettings.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/im_GetNodeCharacSettings.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6a, 0x14, 0x41,
	0x10, 0x25, 0xbb, 0xd9, 0x24, 0x54, 0x84, 0x68, 0x2b, 0x61, 0x32, 0x0b, 0x22, 0x11, 0x44, 0x10,
	0x66, 0xbd, 0x81, 0x22, 0xe8, 0x83, 0x57, 0xf2, 0x90, 0x41, 0x5b, 0x08, 0x24, 0x2f, 0xc3, 0x64,
	0xa6, 0x1c, 0x9b, 0xec, 0x76, 0x2f, 0xdd, 0xbd, 0x09, 0x7e, 0x83, 0x2f, 0xde, 0x3e, 0xc8, 0x0f,
	0xf1, 0x07, 0xf4, 0x2b, 0xac, 0xe9, 0xee, 0xcd, 0xee, 0x8c, 0x09, 0x31, 0x79, 0xd9, 0xa5, 0xba,
	0xce, 0xa9, 0x3a, 0x5d, 0xa7, 0xa6, 0xe1, 0x51, 0x69, 0xac, 0xd2, 0x38, 0x40, 0x59, 0x09, 0x89,
	0x83, 0xb1, 0x56, 0x05, 0x96, 0x13, 0x8d, 0x66, 0x20, 0x46, 0xd9, 0x1b, 0xb4, 0xa9, 0x2a, 0xf1,
	0xc5, 0xc7, 0x5c, 0xe7, 0xc5, 0x7b, 0xb4, 0x56, 0xc8, 0xca, 0x24, 0x84, 0xb1, 0x8a, 0xdd, 0xf2,
	0xc4, 0xc4, 0x13, 0x93, 0xd3, 0xd0, 0xf1, 0xd5, 0xd0, 0xe0, 0x30, 0x1f, 0x4e, 0x30, 0x90, 0xe3,
	0x8d, 0x66, 0x57, 0xd4, 0x5a, 0xe9, 0x90, 0xea, 0x37, 0x53, 0x23, 0x34, 0x26, 0xaf, 0x30, 0x24,
	0x6f, 0xb6, 0x93, 0x36, 0x17, 0xf2, 0x83, 0xd2, 0xa3, 0xdc, 0x0a, 0x25, 0x3d, 0x68, 0xf3, 0x73,
	0x07, 0xe0, 0x2d, 0x69, 0xa0, 0x2c, 0x6a, 0xc3, 0xde, 0xc1, 0xba, 0x24, 0x59, 0x59, 0xe1, 0x74,
	0xd1, 0x91, 0x30, 0x56, 0x14, 0x99, 0x28, 0xa3, 0x85, 0x1b, 0x0b, 0xb7, 0x57, 0xef, 0xf7, 0x93,
	0x70, 0x93, 0xa0, 0x50, 0x48, 0x8b, 0x15, 0xea, 0x9d, 0x3a, 0xe2, 0xd7, 0xe4, 0xf1, 0x8d, 0xa6,
	0xcc, 0xad, 0x92, 0x3d, 0x83, 0xfe, 0xc9, 0x25, 0x33, 0x39, 0x19, 0x0e, 0xa3, 0xdf, 0xcb, 0x54,
	0x78, 0x85, 0x47, 0x27, 0x71, 0x53, 0x02, 0xb0, 0xa7, 0x70, 0xe9, 0x00, 0x3f, 0x65, 0x87, 0xb9,
	0x16, 0xf9, 0xfe, 0x10, 0xa3, 0x8e, 0x13, 0x12, 0xb7, 0x84, 0x18, 0xab, 0x69, 0x82, 0x5e, 0xc7,
	0x2a, 0xe1, 0x77, 0x02, 0x9c, 0xdd, 0x81, 0x2b, 0xf3, 0x74, 0xdf, 0xf4, 0x8f, 0x6f, 0xba, 0x36,
	0x07, 0xac, 0x7b, 0x6d, 0xfe, 0x5a, 0x84, 0x15, 0x8e, 0x66, 0xac, 0xa4, 0x41, 0x76, 0x17, 0x7a,
	0x6e, 0xd6, 0xe1, 0xea, 0xc7, 0x1d, 0x83, 0x89, 0xde, 0x87, 0x57, 0xf5, 0x2f, 0xf7, 0x40, 0xb6,
	0x0b, 0x97, 0xeb, 0x29, 0x67, 0x73, 0x63, 0x26, 0xb9, 0x5d, 0x22, 0x27, 0x2d, 0x72, 0xdb, 0x8c,
	0x6d, 0x8a, 0xb7, 0x66, 0x31, 0x5f, 0x1b, 0x35, 0x0f, 0xd8, 0x63, 0x58, 0x0e, 0xee, 0x46, 0x5d,
	0x57, 0xf1, 0xfa, 0x3f, 0x15, 0xbd, 0xf7, 0xdb, 0xfe, 0x9f, 0x4f, 0xe1, 0xec, 0x35, 0x74, 0xb5,
	0x3a, 0x8a, 0x16, 0x1d, 0xeb, 0x61, 0xf2, 0x7f, 0x9b, 0x98, 0x4c, 0xa7, 0x90, 0x70, 0x75, 0xc4,
	0xeb, 0x02, 0xf1, 0xcf, 0x0e, 0x74, 0x29, 0x60, 0xeb, 0xb0, 0x44, 0x61, 0xbd, 0x12, 0x5f, 0x52,
	0x1a, 0x4c, 0x8f, 0xf7, 0x28, 0x24, 0x9f, 0xf7, 0x20, 0x6e, 0x59, 0x5c, 0xa2, 0x29, 0xb4, 0x18,
	0xbb, 0x31, 0x7c, 0x4d, 0xcf, 0xb4, 0x6d, 0xa3, 0x49, 0x7f, 0x39, 0x63, 0x33, 0x7e, 0xea, 0x5a,
	0x7e, 0x4b, 0x2f, 0xba, 0x97, 0xf7, 0xa0, 0xe7, 0xc0, 0xd1, 0xf7, 0xb3, 0xa5, 0x79, 0x24, 0xad,
	0x72, 0x73, 0x15, 0x7f, 0xa4, 0xe7, 0xda, 0xc5, 0xe7, 0xbb, 0xd0, 0x17, 0xaa, 0xe5, 0xc0, 0xec,
	0x11, 0xd9, 0x7b, 0x52, 0x29, 0x53, 0x1e, 0x4c, 0xf3, 0xe5, 0x79, 0xde, 0x99, 0xfd, 0x25, 0xf7,
	0x39, 0x3f, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x85, 0x7f, 0x25, 0x35, 0xa3, 0x04, 0x00, 0x00,
}
