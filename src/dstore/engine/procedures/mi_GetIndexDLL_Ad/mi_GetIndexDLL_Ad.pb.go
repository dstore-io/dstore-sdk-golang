// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetIndexDLL_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetIndexDLL_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetIndexDLL_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetIndexDLL_Ad

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
	TableName     *dstore_values.StringValue `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	TableNameNull bool                       `protobuf:"varint,1001,opt,name=table_name_null,json=tableNameNull" json:"table_name_null,omitempty"`
	IndexName     *dstore_values.StringValue `protobuf:"bytes,2,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	IndexNameNull bool                       `protobuf:"varint,1002,opt,name=index_name_null,json=indexNameNull" json:"index_name_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTableName() *dstore_values.StringValue {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *Parameters) GetIndexName() *dstore_values.StringValue {
	if m != nil {
		return m.IndexName
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetIndexDLL_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetIndexDLL_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetIndexDLL_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xa6, 0xcd, 0xed, 0xcf, 0x9d, 0xcb, 0xa5, 0x97, 0xb9, 0x20, 0x31, 0xc5, 0x1f, 0xea, 0x42,
	0x11, 0x9c, 0x16, 0xdd, 0xe8, 0x4a, 0x14, 0x45, 0x02, 0x6d, 0x91, 0x59, 0x08, 0xba, 0x09, 0xa9,
	0x39, 0x96, 0x81, 0x64, 0xa6, 0xcc, 0xa4, 0xd6, 0xc7, 0xf0, 0x55, 0x7c, 0x17, 0x5f, 0x42, 0x9f,
	0xc2, 0x49, 0x66, 0xda, 0xd8, 0x74, 0x51, 0x37, 0x6d, 0xce, 0x7c, 0xdf, 0x77, 0xbe, 0x99, 0xef,
	0x1c, 0xd4, 0x8b, 0x54, 0x2a, 0x24, 0x74, 0x81, 0x8f, 0x19, 0x87, 0xee, 0x44, 0x8a, 0x47, 0x88,
	0xa6, 0x12, 0x54, 0x37, 0x61, 0xc1, 0x0d, 0xa4, 0x3e, 0x8f, 0xe0, 0xe5, 0xaa, 0xdf, 0x0f, 0x2e,
	0x22, 0xa2, 0xc1, 0x54, 0xe0, 0x1d, 0xa3, 0x20, 0x46, 0x41, 0x56, 0x68, 0xde, 0x7f, 0xdb, 0xf2,
	0x39, 0x8c, 0xa7, 0xa0, 0x8c, 0xca, 0xdb, 0x5c, 0xf6, 0x01, 0x29, 0x85, 0xb4, 0x50, 0x7b, 0x19,
	0x4a, 0x40, 0xa9, 0x70, 0x0c, 0x16, 0xdc, 0x2b, 0x83, 0x69, 0xc8, 0xf8, 0x93, 0x90, 0x49, 0x98,
	0x32, 0xc1, 0x0d, 0xa9, 0xf3, 0x5e, 0x41, 0xe8, 0x36, 0x94, 0xa1, 0x46, 0x41, 0x2a, 0x7c, 0x86,
	0x50, 0x1a, 0x8e, 0x62, 0x08, 0xb8, 0x3e, 0x71, 0x2b, 0xbb, 0x95, 0x83, 0x3f, 0xc7, 0x1e, 0xb1,
	0xd7, 0xb6, 0xb7, 0x52, 0xa9, 0x64, 0x7c, 0x7c, 0x97, 0x15, 0xf4, 0x77, 0xce, 0x1e, 0x6a, 0x32,
	0xde, 0x47, 0xad, 0x42, 0x1a, 0xf0, 0x69, 0x1c, 0xbb, 0x1f, 0x0d, 0xdd, 0xa0, 0x49, 0xff, 0x2e,
	0x48, 0x43, 0x7d, 0x9a, 0x79, 0xb0, 0xec, 0xcd, 0xc6, 0xa3, 0xba, 0xde, 0x23, 0x67, 0xcf, 0x3d,
	0x0a, 0xa9, 0xf1, 0xf8, 0xb4, 0x1e, 0x0b, 0x52, 0xe6, 0xd1, 0x79, 0xab, 0xa2, 0x26, 0x05, 0x35,
	0x11, 0x5c, 0x01, 0xee, 0xa1, 0x5a, 0x1e, 0x5a, 0xf9, 0x3d, 0x76, 0x0c, 0x26, 0xd0, 0xeb, 0xec,
	0x97, 0x1a, 0x22, 0xbe, 0x47, 0xff, 0xb2, 0xb8, 0x82, 0x6f, 0x79, 0xe9, 0x8b, 0x3a, 0x5a, 0x4c,
	0x4a, 0xe2, 0x72, 0xaa, 0x03, 0x5d, 0xfb, 0x45, 0x4d, 0x5b, 0xc9, 0xf2, 0x01, 0x3e, 0x45, 0x0d,
	0x3b, 0x26, 0xd7, 0xc9, 0x3b, 0x6e, 0xaf, 0x74, 0x34, 0x43, 0x1c, 0x98, 0x7f, 0x3a, 0xa7, 0xe3,
	0x73, 0xe4, 0x48, 0x31, 0x73, 0x7f, 0xe5, 0xaa, 0x23, 0xb2, 0x66, 0x97, 0xc8, 0xfc, 0xf9, 0x84,
	0x8a, 0x19, 0xcd, 0x94, 0xde, 0x16, 0x72, 0xf4, 0x37, 0xde, 0x40, 0x75, 0x5d, 0x05, 0x2c, 0x72,
	0x5f, 0x87, 0x3a, 0x90, 0x1a, 0xad, 0xe9, 0xd2, 0x8f, 0x2e, 0x7d, 0xd4, 0x66, 0xa2, 0xd4, 0xb6,
	0x58, 0xea, 0x87, 0xc3, 0x9f, 0xaf, 0xfb, 0xa8, 0x9e, 0x2f, 0xd7, 0xc9, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xed, 0xf3, 0x74, 0xd3, 0x23, 0x03, 0x00, 0x00,
}
