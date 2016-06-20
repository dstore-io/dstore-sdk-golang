// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetIndexDDL_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetIndexDDL_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetIndexDDL_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetIndexDDL_Ad

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
	TableName                      *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	TableNameNull                  bool                        `protobuf:"varint,1001,opt,name=table_name_null,json=tableNameNull" json:"table_name_null,omitempty"`
	IndexName                      *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	IndexNameNull                  bool                        `protobuf:"varint,1002,opt,name=index_name_null,json=indexNameNull" json:"index_name_null,omitempty"`
	DatabaseName                   *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=database_name,json=databaseName" json:"database_name,omitempty"`
	DatabaseNameNull               bool                        `protobuf:"varint,1003,opt,name=database_name_null,json=databaseNameNull" json:"database_name_null,omitempty"`
	FullyQualifiedNameInOutput     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=fully_qualified_name_in_output,json=fullyQualifiedNameInOutput" json:"fully_qualified_name_in_output,omitempty"`
	FullyQualifiedNameInOutputNull bool                        `protobuf:"varint,1004,opt,name=fully_qualified_name_in_output_null,json=fullyQualifiedNameInOutputNull" json:"fully_qualified_name_in_output_null,omitempty"`
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

func (m *Parameters) GetDatabaseName() *dstore_values.StringValue {
	if m != nil {
		return m.DatabaseName
	}
	return nil
}

func (m *Parameters) GetFullyQualifiedNameInOutput() *dstore_values.BooleanValue {
	if m != nil {
		return m.FullyQualifiedNameInOutput
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetIndexDDL_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetIndexDDL_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetIndexDDL_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetIndexDDL_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0xbb, 0xa4, 0x2d, 0x86, 0xaa, 0xd5, 0x22, 0xa1, 0xb0, 0x11, 0x01, 0xb5, 0x07, 0xb8,
	0x74, 0x53, 0xd1, 0x0b, 0x9c, 0x10, 0xa8, 0x08, 0x45, 0x6a, 0x03, 0xf8, 0x80, 0x04, 0x17, 0xcb,
	0xc1, 0x93, 0xc8, 0x62, 0x63, 0x07, 0xdb, 0x4b, 0xe1, 0x2d, 0x78, 0x15, 0x1e, 0x8b, 0x9f, 0x1b,
	0x2f, 0xc0, 0xf8, 0x67, 0x9b, 0x66, 0x2b, 0x48, 0x2f, 0xd9, 0x8c, 0xe7, 0xfb, 0x19, 0xaf, 0x3f,
	0x2f, 0x39, 0x14, 0xd6, 0x69, 0x03, 0x03, 0x50, 0x53, 0xa9, 0x60, 0x30, 0x37, 0xfa, 0x03, 0x88,
	0xda, 0x80, 0x1d, 0xcc, 0x24, 0x7b, 0x09, 0x6e, 0xa8, 0x04, 0x7c, 0x39, 0x3e, 0x3e, 0x61, 0xcf,
	0x44, 0x89, 0x4d, 0xa7, 0xf3, 0x7b, 0x91, 0x51, 0x46, 0x46, 0x79, 0x09, 0x56, 0xdc, 0x4a, 0x92,
	0x9f, 0x79, 0x55, 0x83, 0x8d, 0xac, 0xe2, 0xce, 0xb2, 0x0f, 0x18, 0xa3, 0x4d, 0x6a, 0xf5, 0x96,
	0x5b, 0x33, 0xb0, 0x96, 0x4f, 0x21, 0x35, 0xf7, 0xdb, 0x4d, 0xc7, 0xa5, 0x9a, 0x68, 0x33, 0xe3,
	0x4e, 0x6a, 0x15, 0x41, 0x7b, 0x7f, 0x32, 0x42, 0x5e, 0x73, 0xc3, 0xb1, 0x0b, 0xc6, 0xe6, 0x4f,
	0x08, 0x71, 0x7c, 0x5c, 0x01, 0x53, 0xb8, 0xd2, 0x5d, 0xbb, 0xbf, 0xf6, 0xf0, 0xc6, 0xa3, 0xa2,
	0x4c, 0x63, 0xa7, 0xa9, 0xac, 0x33, 0x52, 0x4d, 0xdf, 0xfa, 0x82, 0x5e, 0x0f, 0xe8, 0x11, 0x82,
	0xf3, 0x07, 0x64, 0x67, 0x41, 0x65, 0xaa, 0xae, 0xaa, 0xee, 0x8f, 0x4d, 0x14, 0xd8, 0xa2, 0xdb,
	0xe7, 0xa0, 0x11, 0xae, 0x7a, 0x0f, 0xe9, 0xf7, 0x1c, 0x3d, 0xd6, 0x57, 0x7b, 0x04, 0x74, 0xe3,
	0xb1, 0xa0, 0x46, 0x8f, 0x9f, 0xc9, 0xe3, 0x1c, 0x14, 0x3c, 0x9e, 0x92, 0x6d, 0xc1, 0xd1, 0x96,
	0xdb, 0xb4, 0x95, 0x6c, 0xa5, 0xcd, 0xcd, 0x86, 0x10, 0x9c, 0x0e, 0x48, 0xbe, 0x24, 0x10, 0xcd,
	0x7e, 0x45, 0xb3, 0xdd, 0x8b, 0xd0, 0xe0, 0xc7, 0x48, 0x7f, 0x82, 0xcf, 0xaf, 0xec, 0x53, 0xcd,
	0x2b, 0x39, 0x91, 0x20, 0x22, 0x4b, 0x2a, 0xa6, 0x6b, 0x37, 0xaf, 0x5d, 0xf7, 0x5a, 0x18, 0xa0,
	0xd7, 0x1a, 0x60, 0xac, 0x75, 0x05, 0x5c, 0xc5, 0x09, 0x8a, 0x20, 0xf1, 0xa6, 0x51, 0xf0, 0xe2,
	0x43, 0xf5, 0x2a, 0xd0, 0xf3, 0x13, 0xb2, 0xff, 0x7f, 0x83, 0x38, 0xe0, 0xef, 0x38, 0x60, 0xff,
	0xdf, 0x4a, 0x7e, 0xdc, 0xbd, 0xef, 0xeb, 0x64, 0x8b, 0x82, 0x9d, 0x6b, 0x65, 0x21, 0x3f, 0x24,
	0x9d, 0x90, 0xa9, 0xf6, 0x71, 0xa7, 0x94, 0xc6, 0xbc, 0xbd, 0xf0, 0xbf, 0x34, 0x02, 0xf3, 0x77,
	0x64, 0xd7, 0xa7, 0x89, 0x5d, 0x88, 0x13, 0x9e, 0x63, 0x86, 0xe4, 0xb2, 0x45, 0x6e, 0x87, 0xee,
	0x14, 0xeb, 0xe1, 0xa2, 0xa6, 0x3b, 0xb3, 0xe5, 0x85, 0xfc, 0x31, 0xd9, 0x4c, 0x29, 0xc6, 0x23,
	0xf3, 0x8a, 0xfd, 0x4b, 0x8a, 0x31, 0xe3, 0xa7, 0xf1, 0x49, 0x1b, 0x38, 0x1e, 0x79, 0x66, 0xf4,
	0x19, 0xbe, 0x67, 0xcf, 0x3a, 0x28, 0x57, 0x5c, 0xb5, 0xb2, 0xd9, 0x7e, 0x49, 0xf5, 0x19, 0xf5,
	0xcc, 0xe2, 0x2e, 0xc9, 0xf0, 0x7f, 0x7e, 0x9b, 0x6c, 0x60, 0xc5, 0xa4, 0xe8, 0x7e, 0x1b, 0xe1,
	0x0b, 0xe9, 0xd0, 0x0e, 0x96, 0x43, 0xf1, 0x9c, 0x92, 0x9e, 0xd4, 0x2d, 0xd9, 0xc5, 0x9d, 0x7f,
	0x7f, 0x34, 0xd5, 0x56, 0x7c, 0x6c, 0xfa, 0xe2, 0x4a, 0x9f, 0x85, 0xf1, 0x46, 0xb8, 0x84, 0x47,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x61, 0x6d, 0x4e, 0x2a, 0x4b, 0x04, 0x00, 0x00,
}