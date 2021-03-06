// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_UpdatePossibleDuplicates_Ad.proto
// DO NOT EDIT!

/*
Package pm_UpdatePossibleDuplicates_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_UpdatePossibleDuplicates_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_UpdatePossibleDuplicates_Ad

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
	PersonId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull bool                        `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_UpdatePossibleDuplicates_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_UpdatePossibleDuplicates_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_UpdatePossibleDuplicates_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_UpdatePossibleDuplicates_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0x66, 0xeb, 0x6f, 0x7f, 0x7e, 0x51, 0x54, 0x2a, 0x48, 0xed, 0x50, 0xc6, 0x44, 0xf0, 0x42,
	0x32, 0xd1, 0x1b, 0x11, 0x04, 0x15, 0xbd, 0x18, 0xb8, 0x31, 0x02, 0x0a, 0x7a, 0x53, 0xba, 0xe5,
	0x58, 0x82, 0x6d, 0x52, 0x92, 0xd6, 0xbd, 0x86, 0x4f, 0xe4, 0xfb, 0xf8, 0x16, 0xa6, 0x4d, 0x3b,
	0x6d, 0x05, 0x61, 0x37, 0x4d, 0xcf, 0xf9, 0xbe, 0xef, 0x9c, 0x7c, 0x27, 0x07, 0x5d, 0x52, 0x95,
	0x08, 0x09, 0x43, 0xe0, 0x01, 0xe3, 0x30, 0x8c, 0xa5, 0x98, 0x03, 0x4d, 0x25, 0xa8, 0x61, 0x1c,
	0x79, 0x0f, 0x31, 0xf5, 0x13, 0x98, 0x0a, 0xa5, 0xd8, 0x2c, 0x84, 0xdb, 0x34, 0x0e, 0xd9, 0x5c,
	0x27, 0x94, 0x77, 0x4d, 0xb1, 0x66, 0x26, 0xc2, 0x3e, 0x36, 0x72, 0x6c, 0xe4, 0xf8, 0x6f, 0x8d,
	0xbb, 0x5d, 0x34, 0x7b, 0xf3, 0xc3, 0x14, 0x94, 0x29, 0xe1, 0xee, 0x56, 0x6f, 0x00, 0x52, 0x0a,
	0x59, 0x40, 0xbd, 0x2a, 0x14, 0x81, 0x52, 0x7e, 0x00, 0x05, 0x78, 0x50, 0x07, 0x13, 0x9f, 0xf1,
	0x17, 0x21, 0x23, 0x3f, 0x61, 0x82, 0x1b, 0xd2, 0x20, 0x42, 0x68, 0xea, 0x4b, 0x5f, 0x83, 0x20,
	0x95, 0x7d, 0x8e, 0xfe, 0xc7, 0xfa, 0x14, 0xdc, 0x63, 0xd4, 0x69, 0xf4, 0x1b, 0x47, 0x6b, 0xa7,
	0x3d, 0x5c, 0x38, 0x28, 0xee, 0xc4, 0x78, 0x02, 0x01, 0xc8, 0xc7, 0x2c, 0x22, 0x5d, 0xc3, 0x1e,
	0x51, 0xfb, 0x10, 0x6d, 0x2c, 0x95, 0x1e, 0x4f, 0xc3, 0xd0, 0xf9, 0xec, 0x68, 0x7d, 0x97, 0xac,
	0x97, 0x94, 0x89, 0x4e, 0x0e, 0x3e, 0x9a, 0xa8, 0x4b, 0x40, 0xc5, 0x82, 0x2b, 0xb0, 0x4f, 0x50,
	0x2b, 0x37, 0x53, 0x74, 0x72, 0x71, 0x75, 0x56, 0xc6, 0xe8, 0x5d, 0xf6, 0x25, 0x86, 0x68, 0x3f,
	0xa1, 0xad, 0xcc, 0x86, 0xf7, 0xc3, 0x87, 0xd3, 0xec, 0x5b, 0x5a, 0x8c, 0x6b, 0xe2, 0xba, 0xdb,
	0xb1, 0x8e, 0x47, 0xdf, 0x31, 0xd9, 0x8c, 0xaa, 0x09, 0x6d, 0xbd, 0x53, 0x8c, 0xcf, 0xb1, 0xf2,
	0x8a, 0xfb, 0xbf, 0x2a, 0x9a, 0xe1, 0x8e, 0xcd, 0x49, 0x4a, 0xba, 0x7d, 0x8f, 0x2c, 0x29, 0x16,
	0xce, 0xbf, 0x5c, 0x75, 0x81, 0x57, 0x79, 0x70, 0x5c, 0xce, 0x02, 0x13, 0xb1, 0x20, 0x59, 0x19,
	0x77, 0x0f, 0x59, 0xfa, 0xdf, 0xde, 0x41, 0x6d, 0x1d, 0x65, 0xcf, 0xf0, 0x3e, 0xd1, 0xd3, 0x69,
	0x91, 0x96, 0x0e, 0x47, 0xf4, 0xc6, 0x43, 0x3d, 0x26, 0xea, 0x3d, 0x96, 0x3b, 0xf9, 0x7c, 0x15,
	0x08, 0x45, 0x5f, 0x4b, 0x9c, 0xae, 0xbe, 0xb6, 0xb3, 0x76, 0xbe, 0x17, 0x67, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x8f, 0xaa, 0x86, 0xde, 0xf8, 0x02, 0x00, 0x00,
}
