// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_InsertPageHit_Pu.proto
// DO NOT EDIT!

/*
Package st_InsertPageHit_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_InsertPageHit_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_InsertPageHit_Pu

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
	PageDescription     *dstore_values.StringValue `protobuf:"bytes,1,opt,name=page_description,json=pageDescription" json:"page_description,omitempty"`
	PageDescriptionNull bool                       `protobuf:"varint,1001,opt,name=page_description_null,json=pageDescriptionNull" json:"page_description_null,omitempty"`
	UniqueId            *dstore_values.StringValue `protobuf:"bytes,2,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull        bool                       `protobuf:"varint,1002,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPageDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PageDescription
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_InsertPageHit_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_InsertPageHit_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_InsertPageHit_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x4a, 0xeb, 0x30,
	0x18, 0x67, 0xeb, 0xd9, 0x9f, 0x93, 0x73, 0x70, 0x23, 0x43, 0xa9, 0x1d, 0x8a, 0x4e, 0x04, 0x2f,
	0x24, 0x95, 0xed, 0x42, 0x6f, 0x15, 0x07, 0x16, 0xd9, 0x18, 0xb9, 0x10, 0xf4, 0xa6, 0xd4, 0x35,
	0x8e, 0xc0, 0x96, 0xd4, 0x24, 0x75, 0xaf, 0xe1, 0xbb, 0xf8, 0x52, 0xea, 0x53, 0x98, 0x34, 0x9d,
	0x73, 0x55, 0x18, 0xde, 0xb4, 0xfd, 0xfa, 0xfb, 0xf3, 0xa5, 0xbf, 0xef, 0x2b, 0xe8, 0xc6, 0x52,
	0x71, 0x41, 0x7c, 0xc2, 0x26, 0x94, 0x11, 0x3f, 0x11, 0x7c, 0x4c, 0xe2, 0x54, 0x10, 0xe9, 0x4b,
	0x15, 0x06, 0x4c, 0x12, 0xa1, 0x46, 0xd1, 0x84, 0x5c, 0x51, 0x15, 0x8e, 0x52, 0xa4, 0x61, 0xc5,
	0xe1, 0xbe, 0xd5, 0x20, 0xab, 0x41, 0x3f, 0x10, 0xbd, 0x56, 0x6e, 0xfb, 0x14, 0x4d, 0x53, 0x22,
	0xad, 0xce, 0xdb, 0x5e, 0xed, 0x45, 0x84, 0xe0, 0x22, 0x87, 0xda, 0xab, 0xd0, 0x8c, 0x48, 0xa9,
	0xcd, 0x72, 0xf0, 0xa0, 0x08, 0xaa, 0x88, 0xb2, 0x07, 0x2e, 0x66, 0x91, 0xa2, 0x9c, 0x59, 0x52,
	0xe7, 0xb5, 0x04, 0xc0, 0x28, 0x12, 0x91, 0x46, 0x89, 0x90, 0xb0, 0x0f, 0x9a, 0x89, 0x76, 0x08,
	0x63, 0x22, 0xc7, 0x82, 0x26, 0x86, 0xe8, 0x96, 0xf6, 0x4a, 0x47, 0xff, 0xba, 0x1e, 0xca, 0x8f,
	0x9f, 0x9f, 0x4d, 0x2a, 0x41, 0xd9, 0xe4, 0xc6, 0x14, 0xb8, 0x61, 0x34, 0x97, 0x4b, 0x09, 0xec,
	0x81, 0xcd, 0xa2, 0x4d, 0xc8, 0xd2, 0xe9, 0xd4, 0x7d, 0xab, 0x69, 0xb3, 0x3a, 0x6e, 0x15, 0x04,
	0x43, 0x8d, 0xc1, 0x53, 0xf0, 0x37, 0x65, 0xf4, 0x31, 0x25, 0x21, 0x8d, 0xdd, 0xf2, 0xda, 0xa6,
	0x75, 0x4b, 0x0e, 0x62, 0x78, 0x08, 0x36, 0x3e, 0x85, 0xb6, 0xcd, 0xbb, 0x6d, 0xf3, 0x7f, 0x41,
	0x31, 0xfe, 0x9d, 0x97, 0x32, 0xa8, 0x63, 0x22, 0x13, 0xae, 0x53, 0x87, 0x27, 0xa0, 0x92, 0x05,
	0x59, 0xfc, 0xba, 0x7c, 0x38, 0x36, 0xe4, 0xbe, 0xb9, 0x62, 0x4b, 0x84, 0xb7, 0xa0, 0x69, 0x22,
	0x0c, 0xbf, 0x64, 0xa8, 0x4f, 0xe9, 0x68, 0x31, 0x2a, 0x88, 0x8b, 0x49, 0x0f, 0x74, 0x1d, 0x2c,
	0x6b, 0xdc, 0x98, 0xad, 0xbe, 0x80, 0x67, 0xa0, 0x96, 0x8f, 0xce, 0x75, 0x32, 0xc7, 0xdd, 0x6f,
	0x8e, 0x76, 0xb0, 0x03, 0x7b, 0xc7, 0x0b, 0x3a, 0x3c, 0x07, 0x8e, 0xe0, 0x73, 0xf7, 0x4f, 0xa6,
	0xf2, 0xd1, 0xda, 0x0d, 0x43, 0x8b, 0x00, 0x10, 0xe6, 0x73, 0x6c, 0xb4, 0xde, 0x0e, 0x70, 0xf4,
	0x33, 0xdc, 0x02, 0x55, 0x5d, 0x99, 0xe8, 0x9f, 0x87, 0x3a, 0x92, 0x0a, 0xae, 0xe8, 0x32, 0x88,
	0x2f, 0xae, 0x41, 0x9b, 0xf2, 0x82, 0xf1, 0x72, 0xdd, 0xef, 0x8e, 0x7f, 0xf3, 0x23, 0xdc, 0x57,
	0xb3, 0xa5, 0xeb, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xc7, 0x52, 0xaa, 0x3f, 0x03, 0x00,
	0x00,
}