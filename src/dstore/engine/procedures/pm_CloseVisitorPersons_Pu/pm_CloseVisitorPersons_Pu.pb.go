// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_CloseVisitorPersons_Pu.proto
// DO NOT EDIT!

/*
Package pm_CloseVisitorPersons_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_CloseVisitorPersons_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_CloseVisitorPersons_Pu

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
	UniqueId     *dstore_values.StringValue `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull bool                       `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_CloseVisitorPersons_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_CloseVisitorPersons_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_CloseVisitorPersons_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xa6, 0x8d, 0xfd, 0x71, 0x15, 0x95, 0x15, 0x24, 0xa6, 0x28, 0xa5, 0x22, 0x7a, 0xda, 0x4a,
	0x45, 0xec, 0x59, 0x11, 0xe9, 0xa1, 0xa5, 0xee, 0xa1, 0xa0, 0x97, 0x10, 0xcd, 0x58, 0x16, 0x92,
	0xdd, 0xba, 0x9b, 0xd8, 0xd7, 0xf0, 0x75, 0x7c, 0x24, 0xdf, 0xc2, 0x49, 0x36, 0xad, 0x36, 0x22,
	0xe8, 0x25, 0xd9, 0xd9, 0xf9, 0xbe, 0x99, 0xf9, 0xbe, 0x1d, 0xd2, 0x0f, 0x4d, 0xa2, 0x34, 0x74,
	0x41, 0x4e, 0x85, 0x84, 0xee, 0x4c, 0xab, 0x27, 0x08, 0x53, 0x0d, 0xa6, 0x3b, 0x8b, 0xfd, 0xeb,
	0x48, 0x19, 0x98, 0x08, 0x23, 0x10, 0x32, 0x06, 0x6d, 0x94, 0x34, 0xfe, 0x38, 0x65, 0x08, 0x4a,
	0x14, 0x3d, 0xb1, 0x4c, 0x66, 0x99, 0xec, 0x57, 0xb8, 0xb7, 0x5b, 0xb4, 0x78, 0x0d, 0xa2, 0x14,
	0x8c, 0x65, 0x7b, 0xfb, 0xab, 0x7d, 0x41, 0x6b, 0xa5, 0x8b, 0x54, 0x6b, 0x35, 0x15, 0x83, 0x31,
	0xc1, 0x14, 0x8a, 0xe4, 0x51, 0x39, 0x99, 0x04, 0x42, 0x3e, 0x2b, 0x1d, 0x07, 0x89, 0x50, 0xd2,
	0x82, 0x3a, 0x11, 0x21, 0xe3, 0x40, 0x07, 0x98, 0xc4, 0x29, 0xe8, 0x25, 0x59, 0x4f, 0xa5, 0x78,
	0x49, 0xc1, 0x17, 0xa1, 0x5b, 0x69, 0x57, 0x4e, 0x37, 0x7a, 0x1e, 0x2b, 0x86, 0x2f, 0x66, 0x32,
	0x89, 0x16, 0x72, 0x3a, 0xc9, 0x02, 0xde, 0xb4, 0xe0, 0x41, 0x48, 0x8f, 0xc9, 0xd6, 0x92, 0xe8,
	0xcb, 0x34, 0x8a, 0xdc, 0x8f, 0x06, 0xd2, 0x9b, 0x7c, 0x73, 0x01, 0x19, 0xe1, 0x65, 0xe7, 0xbd,
	0x4a, 0x9a, 0x1c, 0xcc, 0x0c, 0xe5, 0x02, 0x3d, 0x23, 0xb5, 0x5c, 0x4b, 0xb9, 0x51, 0xe1, 0x92,
	0xd5, 0x79, 0x93, 0x7d, 0xb9, 0x05, 0xd2, 0x7b, 0xb2, 0x93, 0xa9, 0xf0, 0xbf, 0xc9, 0x70, 0xab,
	0x6d, 0x07, 0xc9, 0xac, 0x44, 0x2e, 0x8b, 0x1d, 0x62, 0x3c, 0xf8, 0x8a, 0xf9, 0x76, 0xbc, 0x7a,
	0x41, 0xfb, 0xa4, 0x51, 0xb8, 0xe7, 0x3a, 0x79, 0xc5, 0xc3, 0x1f, 0x15, 0xad, 0xb7, 0x43, 0xfb,
	0xe7, 0x0b, 0x38, 0xbd, 0x25, 0x8e, 0x56, 0x73, 0x77, 0x2d, 0x67, 0x5d, 0xb0, 0x3f, 0x3e, 0x35,
	0x5b, 0xd8, 0xc0, 0xb8, 0x9a, 0xf3, 0xac, 0x82, 0x77, 0x40, 0x1c, 0x3c, 0xd3, 0x3d, 0x52, 0xc7,
	0x28, 0x7b, 0x80, 0xb7, 0x11, 0x1a, 0x53, 0xe3, 0x35, 0x0c, 0x07, 0xe1, 0xd5, 0x1d, 0x69, 0x09,
	0x55, 0x2e, 0xbf, 0xdc, 0xc1, 0x87, 0xde, 0xff, 0xb7, 0xf3, 0xb1, 0x9e, 0xef, 0xc0, 0xf9, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x05, 0x09, 0x7c, 0xda, 0x02, 0x00, 0x00,
}
