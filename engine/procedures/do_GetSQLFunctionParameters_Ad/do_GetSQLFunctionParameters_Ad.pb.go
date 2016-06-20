// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetSQLFunctionParameters_Ad.proto
// DO NOT EDIT!

/*
Package do_GetSQLFunctionParameters_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetSQLFunctionParameters_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetSQLFunctionParameters_Ad

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
	DefaultValue              *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	Length                    *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=length" json:"length,omitempty"`
	IntroducedIndstoreVersion *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=introduced_indstore_version,json=introducedIndstoreVersion" json:"introduced_indstore_version,omitempty"`
	DataType                  *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=data_type,json=dataType" json:"data_type,omitempty"`
	PrecisionValue            *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=precision_value,json=precisionValue" json:"precision_value,omitempty"`
	ParameterName             *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=parameter_name,json=parameterName" json:"parameter_name,omitempty"`
	Scale                     *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=scale" json:"scale,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetDefaultValue() *dstore_values.StringValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *Response_Row) GetLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Response_Row) GetIntroducedIndstoreVersion() *dstore_values.StringValue {
	if m != nil {
		return m.IntroducedIndstoreVersion
	}
	return nil
}

func (m *Response_Row) GetDataType() *dstore_values.StringValue {
	if m != nil {
		return m.DataType
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

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetSQLFunctionParameters_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetSQLFunctionParameters_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetSQLFunctionParameters_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/do_GetSQLFunctionParameters_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x6f, 0x6b, 0x13, 0x41,
	0x10, 0xc6, 0x89, 0x31, 0x69, 0x5c, 0x6d, 0x23, 0x1b, 0x90, 0x6d, 0x02, 0x52, 0xea, 0x1b, 0x5f,
	0xc8, 0x45, 0xaa, 0x88, 0x08, 0x42, 0x2d, 0x54, 0x0d, 0xb4, 0x41, 0x4f, 0x29, 0xa8, 0x2f, 0x96,
	0x33, 0x3b, 0x39, 0x0f, 0x2f, 0xbb, 0xe7, 0xee, 0x5e, 0x8b, 0x1f, 0x42, 0xf1, 0xbf, 0x9f, 0x51,
	0xfc, 0x12, 0xce, 0xdd, 0x6e, 0x7a, 0xe6, 0x94, 0xa6, 0x7d, 0x93, 0x30, 0x37, 0xf3, 0x9b, 0x7b,
	0xe6, 0x99, 0xdb, 0x25, 0xf7, 0x85, 0xb1, 0x4a, 0xc3, 0x10, 0x64, 0x9c, 0x48, 0x18, 0x66, 0x5a,
	0x4d, 0x40, 0xe4, 0x1a, 0xcc, 0x50, 0x28, 0xfe, 0x08, 0xec, 0xb3, 0xa7, 0x7b, 0x0f, 0x73, 0x39,
	0xb1, 0x89, 0x92, 0x4f, 0x22, 0x1d, 0xcd, 0xc0, 0x82, 0x36, 0xfc, 0x81, 0x08, 0xb0, 0xd2, 0x2a,
	0x7a, 0xc3, 0xe1, 0x81, 0xc3, 0x83, 0x93, 0x99, 0x7e, 0xcf, 0xbf, 0xec, 0x30, 0x4a, 0x73, 0x30,
	0xae, 0x45, 0x7f, 0x7d, 0x51, 0x01, 0x68, 0xad, 0xb4, 0x4f, 0x0d, 0x16, 0x53, 0x33, 0x30, 0x26,
	0x8a, 0xc1, 0x27, 0xaf, 0xd5, 0x93, 0x36, 0x4a, 0xe4, 0x54, 0xe9, 0x59, 0x54, 0xbc, 0xd8, 0x15,
	0x6d, 0x7e, 0x6c, 0x10, 0x52, 0x69, 0xa0, 0x8f, 0x49, 0xcf, 0xf0, 0x77, 0x3c, 0xe5, 0x53, 0xaf,
	0x8f, 0x4b, 0x4c, 0xb1, 0xc6, 0x46, 0xe3, 0xfa, 0xc5, 0xad, 0x7e, 0xe0, 0x87, 0xf1, 0xf2, 0x8c,
	0xd5, 0x89, 0x8c, 0x0f, 0x8a, 0x20, 0xec, 0x9a, 0x6a, 0xa6, 0x31, 0x22, 0xf4, 0x0e, 0x61, 0xff,
	0xe9, 0xc4, 0x65, 0x9e, 0xa6, 0xec, 0xd7, 0x0a, 0xf6, 0xeb, 0x84, 0xbd, 0x1a, 0x33, 0xc6, 0xdc,
	0xe6, 0x87, 0x36, 0xe9, 0x84, 0x60, 0x32, 0x25, 0x0d, 0xd0, 0x9b, 0xa4, 0x55, 0x8e, 0x5b, 0x17,
	0xe0, 0xdd, 0x74, 0x56, 0xec, 0x16, 0xbf, 0xa1, 0x2b, 0xa4, 0x2f, 0xc8, 0xe5, 0x62, 0x50, 0xfe,
	0xd7, 0xa4, 0xec, 0xdc, 0x46, 0x13, 0xe1, 0xa0, 0x06, 0xd7, 0xfd, 0xd8, 0xc7, 0x78, 0x54, 0xc5,
	0x61, 0x77, 0xb6, 0xf8, 0x80, 0xde, 0x25, 0x2b, 0xde, 0x60, 0xd6, 0x2c, 0x3b, 0x5e, 0xfd, 0xa7,
	0xa3, 0xb3, 0x7f, 0xdf, 0xfd, 0x87, 0xf3, 0x72, 0xba, 0x47, 0x9a, 0x5a, 0x1d, 0xb1, 0xf3, 0x25,
	0x75, 0x2f, 0x38, 0xcb, 0x27, 0x11, 0xcc, 0xbd, 0x08, 0x42, 0x75, 0x14, 0x16, 0x6d, 0xfa, 0xbf,
	0x9b, 0xa4, 0x89, 0x01, 0xbd, 0x42, 0xda, 0x18, 0xf2, 0x44, 0xb0, 0x4f, 0x63, 0xb4, 0xa7, 0x15,
	0xb6, 0x30, 0x1c, 0x09, 0xba, 0x4d, 0x56, 0x05, 0x4c, 0xa3, 0x3c, 0xb5, 0xbc, 0x5c, 0x14, 0xfb,
	0x3c, 0x5e, 0xba, 0xbe, 0x4b, 0x9e, 0x28, 0x23, 0x7a, 0x9b, 0xb4, 0x53, 0x54, 0x67, 0xdf, 0xb0,
	0x2f, 0x0e, 0x1d, 0xd4, 0xd0, 0x44, 0x5a, 0x88, 0x41, 0x3b, 0xd6, 0xd7, 0xd2, 0x57, 0x64, 0x80,
	0xcf, 0xb5, 0x12, 0x39, 0x9e, 0x10, 0x5c, 0x80, 0x43, 0xf8, 0x21, 0x0e, 0x52, 0x6c, 0xe1, 0xeb,
	0x72, 0x15, 0xeb, 0x15, 0x3f, 0xf2, 0xf8, 0x81, 0xa3, 0xd1, 0xfc, 0x0b, 0x22, 0xc2, 0xbd, 0xda,
	0xf7, 0x19, 0xb0, 0x6f, 0xcb, 0x5b, 0x75, 0x8a, 0xea, 0xe7, 0x58, 0x4c, 0x77, 0x49, 0x37, 0xd3,
	0x30, 0x49, 0x8a, 0x36, 0xde, 0x90, 0xef, 0xa7, 0x98, 0x6a, 0xed, 0x18, 0x72, 0x9e, 0xec, 0x90,
	0xb5, 0x6c, 0xbe, 0x18, 0x77, 0x28, 0x7e, 0x2c, 0x57, 0xb1, 0x7a, 0x8c, 0x94, 0x67, 0x62, 0x8b,
	0xb4, 0xcc, 0x24, 0x4a, 0x81, 0xfd, 0x3c, 0x85, 0x00, 0x57, 0xba, 0xc3, 0xd1, 0x55, 0x55, 0xfb,
	0x64, 0xaa, 0x4b, 0xe8, 0xe5, 0x76, 0xac, 0x8c, 0x78, 0x3b, 0xcf, 0x8b, 0xb3, 0xdf, 0x53, 0xaf,
	0xdb, 0xe5, 0x45, 0x70, 0xeb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x07, 0x31, 0xcb, 0xe9,
	0x04, 0x00, 0x00,
}