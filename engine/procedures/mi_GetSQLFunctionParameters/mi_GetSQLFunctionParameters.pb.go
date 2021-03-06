// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetSQLFunctionParameters.proto
// DO NOT EDIT!

/*
Package mi_GetSQLFunctionParameters is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetSQLFunctionParameters.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetSQLFunctionParameters

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
	SQLFunctionName     *dstore_values.StringValue `protobuf:"bytes,1,opt,name=s_q_l_function_name,json=sQLFunctionName" json:"s_q_l_function_name,omitempty"`
	SQLFunctionNameNull bool                       `protobuf:"varint,1001,opt,name=s_q_l_function_name_null,json=sQLFunctionNameNull" json:"s_q_l_function_name_null,omitempty"`
	ParameterName       *dstore_values.StringValue `protobuf:"bytes,2,opt,name=parameter_name,json=parameterName" json:"parameter_name,omitempty"`
	ParameterNameNull   bool                       `protobuf:"varint,1002,opt,name=parameter_name_null,json=parameterNameNull" json:"parameter_name_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSQLFunctionName() *dstore_values.StringValue {
	if m != nil {
		return m.SQLFunctionName
	}
	return nil
}

func (m *Parameters) GetParameterName() *dstore_values.StringValue {
	if m != nil {
		return m.ParameterName
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
	RowId          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Length         *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=length" json:"length,omitempty"`
	PrecisionValue *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=precision_value,json=precisionValue" json:"precision_value,omitempty"`
	ParameterName  *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=parameter_name,json=parameterName" json:"parameter_name,omitempty"`
	Scale          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=scale" json:"scale,omitempty"`
	Datatype       *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=datatype" json:"datatype,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Response_Row) GetPrecisionValue() *dstore_values.IntegerValue {
	if m != nil {
		return m.PrecisionValue
	}
	return nil
}

func (m *Response_Row) GetParameterName() *dstore_values.StringValue {
	if m != nil {
		return m.ParameterName
	}
	return nil
}

func (m *Response_Row) GetScale() *dstore_values.IntegerValue {
	if m != nil {
		return m.Scale
	}
	return nil
}

func (m *Response_Row) GetDatatype() *dstore_values.StringValue {
	if m != nil {
		return m.Datatype
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetSQLFunctionParameters.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetSQLFunctionParameters.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetSQLFunctionParameters.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetSQLFunctionParameters.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x66, 0xdb, 0x6d, 0xb7, 0x1c, 0x71, 0xab, 0x53, 0x90, 0xb1, 0x05, 0x91, 0xf5, 0x46, 0x6f,
	0x52, 0x59, 0xc5, 0x15, 0xc1, 0x0b, 0x17, 0x56, 0x2d, 0xb8, 0x41, 0x23, 0x08, 0x8a, 0x10, 0xc6,
	0xe6, 0x6c, 0x1c, 0x4c, 0x66, 0xe2, 0xcc, 0xd4, 0xc5, 0x5b, 0x9f, 0xc0, 0xdf, 0xa7, 0xf2, 0x49,
	0xdc, 0xa7, 0x70, 0x92, 0x99, 0xb6, 0x26, 0x2b, 0xbb, 0xed, 0x4d, 0xc2, 0xc9, 0xf9, 0xbe, 0xf3,
	0x9d, 0x39, 0xdf, 0xc9, 0xc0, 0x83, 0x44, 0x1b, 0xa9, 0x70, 0x8c, 0x22, 0xe5, 0x02, 0xc7, 0x85,
	0x92, 0x53, 0x4c, 0x66, 0x0a, 0xf5, 0x38, 0xe7, 0xf1, 0x13, 0x34, 0x2f, 0x5f, 0x3c, 0x7b, 0x3c,
	0x13, 0x53, 0xc3, 0xa5, 0x78, 0xce, 0x14, 0xcb, 0xd1, 0xa0, 0xd2, 0x81, 0x85, 0x19, 0x49, 0x6e,
	0x39, 0x6e, 0xe0, 0xb8, 0xc1, 0x19, 0x84, 0xe1, 0xc0, 0xcb, 0x7c, 0x62, 0xd9, 0x0c, 0x3d, 0x7f,
	0x78, 0xb5, 0xae, 0x8d, 0x4a, 0x49, 0xe5, 0x53, 0xa3, 0x7a, 0x2a, 0x47, 0xad, 0x59, 0x8a, 0x3e,
	0x79, 0xa3, 0x99, 0x34, 0x8c, 0x8b, 0x23, 0xa9, 0x72, 0x56, 0xaa, 0x3a, 0xd0, 0xce, 0x97, 0x16,
	0xc0, 0xb2, 0x01, 0xf2, 0x14, 0x06, 0x3a, 0xfe, 0x18, 0x67, 0xf1, 0x91, 0x6f, 0x2e, 0x16, 0x36,
	0x45, 0x37, 0xae, 0x6f, 0xdc, 0xbc, 0xb0, 0x3b, 0x0c, 0xfc, 0x49, 0x7c, 0x7b, 0xda, 0x28, 0x2e,
	0xd2, 0x57, 0x65, 0x10, 0xf5, 0xf5, 0xf2, 0x40, 0xa1, 0xa5, 0x90, 0x7b, 0x40, 0xff, 0x53, 0x29,
	0x16, 0xb3, 0x2c, 0xa3, 0x7f, 0xb6, 0x6c, 0xbd, 0x5e, 0x34, 0x68, 0x70, 0x42, 0x9b, 0x23, 0x8f,
	0x60, 0xbb, 0x98, 0xf7, 0xe3, 0xc4, 0x5b, 0xe7, 0x8a, 0x5f, 0x5c, 0x30, 0x2a, 0xe9, 0x31, 0x0c,
	0xea, 0x25, 0x9c, 0xea, 0x89, 0x53, 0xbd, 0x5c, 0x03, 0x97, 0x9a, 0x3b, 0x27, 0x9b, 0xd0, 0x8b,
	0x50, 0x17, 0x52, 0x68, 0x24, 0xb7, 0xa1, 0x53, 0x8d, 0xb8, 0x79, 0x68, 0x6f, 0x9f, 0x1b, 0xff,
	0x41, 0xf9, 0x8c, 0x1c, 0x90, 0xbc, 0x86, 0x4b, 0xe5, 0x70, 0xe3, 0x7f, 0xa6, 0x6b, 0x9b, 0x6e,
	0x5b, 0x72, 0xd0, 0x20, 0x37, 0x3d, 0x38, 0xb4, 0xf1, 0x64, 0x19, 0x47, 0xfd, 0xbc, 0xfe, 0x81,
	0xdc, 0x87, 0x2d, 0x6f, 0x2a, 0x6d, 0x57, 0x15, 0xaf, 0x9d, 0xaa, 0xe8, 0x2c, 0x3f, 0x74, 0xef,
	0x68, 0x0e, 0x27, 0x13, 0x68, 0x2b, 0x79, 0x4c, 0x37, 0x2b, 0xd6, 0x5e, 0xb0, 0xf2, 0x0e, 0x06,
	0xf3, 0x41, 0x04, 0x91, 0x3c, 0x8e, 0xca, 0x1a, 0xc3, 0xdf, 0x2d, 0x68, 0xdb, 0x80, 0x5c, 0x81,
	0xae, 0x0d, 0x63, 0x9e, 0xd0, 0xaf, 0xa1, 0x9d, 0x4d, 0x27, 0xea, 0xd8, 0x70, 0x92, 0x90, 0xbb,
	0xd0, 0xcd, 0x6c, 0x61, 0xf3, 0x9e, 0x7e, 0x0b, 0xab, 0x99, 0x8d, 0x1a, 0x5e, 0x71, 0x61, 0x30,
	0x45, 0xe5, 0xcc, 0xf2, 0x58, 0x72, 0x00, 0xfd, 0x42, 0xe1, 0x94, 0xeb, 0x72, 0x37, 0x2a, 0x20,
	0xfd, 0xbe, 0x02, 0x7d, 0x7b, 0x41, 0xaa, 0x62, 0xb2, 0x7f, 0x6a, 0x5f, 0x7e, 0x84, 0xeb, 0x2e,
	0xcc, 0x2e, 0x74, 0xf4, 0x94, 0x65, 0x48, 0x7f, 0xae, 0xd0, 0x80, 0x83, 0x92, 0x3d, 0xe8, 0x25,
	0xcc, 0x30, 0xf3, 0xb9, 0x40, 0xfa, 0xeb, 0x7c, 0xc5, 0x05, 0x78, 0xff, 0x2d, 0x8c, 0xb8, 0x6c,
	0xf8, 0xb1, 0xbc, 0x4f, 0xde, 0x3c, 0x4c, 0xa5, 0x4e, 0x3e, 0xcc, 0xf3, 0xc9, 0x9a, 0x57, 0xce,
	0xbb, 0x6e, 0xf5, 0x5b, 0xdf, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x08, 0x76, 0xed, 0x7f, 0xb1,
	0x04, 0x00, 0x00,
}
