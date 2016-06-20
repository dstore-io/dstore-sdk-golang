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

var fileDescriptor0 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xd6, 0xda, 0xb5, 0xab, 0x0e, 0x62, 0x05, 0x57, 0x42, 0xa6, 0x95, 0x10, 0x1a, 0x37, 0x70,
	0x93, 0xa2, 0x31, 0x31, 0xc4, 0x1d, 0x93, 0x06, 0x54, 0x62, 0x11, 0x78, 0x12, 0x12, 0xdc, 0x44,
	0xa6, 0x39, 0x2b, 0x96, 0x12, 0x3b, 0xd8, 0x2e, 0x13, 0xb7, 0x3c, 0x01, 0xbf, 0x4f, 0xc5, 0x93,
	0xb0, 0xa7, 0xc0, 0x89, 0xdd, 0x96, 0x64, 0x88, 0x15, 0x6e, 0x12, 0x9d, 0x9c, 0xef, 0x3b, 0xdf,
	0xf1, 0xf9, 0x4e, 0x0c, 0x0f, 0x53, 0x63, 0x95, 0xc6, 0x31, 0xca, 0x99, 0x90, 0x38, 0x2e, 0xb4,
	0x9a, 0x62, 0x3a, 0xd7, 0x68, 0xc6, 0xb9, 0x48, 0x9e, 0xa0, 0x3d, 0x7e, 0xf1, 0xec, 0xf1, 0x5c,
	0x4e, 0xad, 0x50, 0xf2, 0x39, 0xd7, 0x3c, 0x47, 0x8b, 0xda, 0x44, 0x0e, 0x66, 0x15, 0xb9, 0xe3,
	0xb9, 0x91, 0xe7, 0x46, 0x7f, 0x21, 0x0c, 0x07, 0x41, 0xe6, 0x3d, 0xcf, 0xe6, 0x18, 0xf8, 0xc3,
	0xeb, 0x75, 0x6d, 0xd4, 0x5a, 0xe9, 0x90, 0x1a, 0xd5, 0x53, 0x39, 0x1a, 0xc3, 0x67, 0x18, 0x92,
	0xb7, 0x9a, 0x49, 0xcb, 0x85, 0x3c, 0x51, 0x3a, 0xe7, 0xa5, 0xaa, 0x07, 0xed, 0x7c, 0x6c, 0x01,
	0xac, 0x1a, 0x20, 0x4f, 0x61, 0x60, 0x92, 0x77, 0x49, 0x96, 0x9c, 0x84, 0xe6, 0x12, 0xe9, 0x52,
	0x74, 0xe3, 0xe6, 0xc6, 0xed, 0x4b, 0xbb, 0xc3, 0x28, 0x9c, 0x24, 0xb4, 0x67, 0xac, 0x16, 0x72,
	0xf6, 0xb2, 0x0c, 0x58, 0xdf, 0xac, 0x0e, 0x14, 0x3b, 0x0a, 0xb9, 0x0f, 0xf4, 0x0f, 0x95, 0x12,
	0x39, 0xcf, 0x32, 0xfa, 0x73, 0xcb, 0xd5, 0xeb, 0xb1, 0x41, 0x83, 0x13, 0xbb, 0x1c, 0x79, 0x04,
	0xdb, 0xc5, 0xa2, 0x1f, 0x2f, 0xde, 0xba, 0x50, 0xfc, 0xf2, 0x92, 0x51, 0x49, 0x8f, 0x61, 0x50,
	0x2f, 0xe1, 0x55, 0xcf, 0xbc, 0xea, 0xd5, 0x1a, 0xb8, 0xd4, 0xdc, 0x39, 0xdb, 0x84, 0x1e, 0x43,
	0x53, 0x28, 0x69, 0x90, 0xdc, 0x85, 0x4e, 0x35, 0xe2, 0xe6, 0xa1, 0x83, 0x7d, 0x7e, 0xfc, 0x87,
	0xe5, 0x93, 0x79, 0x20, 0x79, 0x05, 0x57, 0xca, 0xe1, 0x26, 0xbf, 0x4d, 0xd7, 0x35, 0xdd, 0x76,
	0xe4, 0xa8, 0x41, 0x6e, 0x7a, 0x70, 0xe4, 0xe2, 0xc9, 0x2a, 0x66, 0xfd, 0xbc, 0xfe, 0x81, 0x3c,
	0x80, 0xad, 0x60, 0x2a, 0x6d, 0x57, 0x15, 0x6f, 0x9c, 0xab, 0xe8, 0x2d, 0x3f, 0xf2, 0x6f, 0xb6,
	0x80, 0x93, 0x09, 0xb4, 0xb5, 0x3a, 0xa5, 0x9b, 0x15, 0x6b, 0x3f, 0x5a, 0x7b, 0x07, 0xa3, 0xc5,
	0x20, 0x22, 0xa6, 0x4e, 0x59, 0x59, 0x63, 0xf8, 0xa3, 0x05, 0x6d, 0x17, 0x90, 0x6b, 0xd0, 0x75,
	0x61, 0x22, 0x52, 0xfa, 0x29, 0x76, 0xb3, 0xe9, 0xb0, 0x8e, 0x0b, 0x27, 0x29, 0xd9, 0x83, 0x6e,
	0xe6, 0x0a, 0xdb, 0xb7, 0xf4, 0x73, 0x5c, 0xcd, 0x6c, 0xd4, 0xf0, 0x4a, 0x48, 0x8b, 0x33, 0xd4,
	0xde, 0xac, 0x80, 0x25, 0x87, 0xd0, 0x2f, 0x34, 0x4e, 0x85, 0x29, 0x77, 0xa3, 0x02, 0xd2, 0x2f,
	0x6b, 0xd0, 0xb7, 0x97, 0xa4, 0x2a, 0x26, 0x07, 0xe7, 0xf6, 0xe5, 0x6b, 0xfc, 0xaf, 0x0b, 0xb3,
	0x0b, 0x1d, 0x33, 0xe5, 0x19, 0xd2, 0x6f, 0x6b, 0x34, 0xe0, 0xa1, 0x64, 0x1f, 0x7a, 0x29, 0xb7,
	0xdc, 0x7e, 0x28, 0x90, 0x7e, 0xbf, 0x58, 0x71, 0x09, 0x3e, 0x38, 0x86, 0x91, 0x50, 0x0d, 0x3f,
	0x56, 0xf7, 0xc9, 0xeb, 0xbd, 0xff, 0xb9, 0x69, 0xde, 0x74, 0xab, 0xbf, 0xf9, 0xde, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x46, 0x7a, 0xf1, 0x0d, 0xa8, 0x04, 0x00, 0x00,
}
