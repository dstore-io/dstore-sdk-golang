// Code generated by protoc-gen-go.
// source: dstore/engine/procedure.proto
// DO NOT EDIT!

/*
Package procedure is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedure.proto

It has these top-level messages:
	Call
	Parameter
	Response
*/
package procedure

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

type Call struct {
	ProcedureName string       `protobuf:"bytes,1,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	Parameter     []*Parameter `protobuf:"bytes,2,rep,name=parameter" json:"parameter,omitempty"`
	CallId        int32        `protobuf:"varint,3,opt,name=call_id,json=callId" json:"call_id,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Call) GetParameter() []*Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

type Parameter struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	IsNull bool   `protobuf:"varint,3,opt,name=is_null,json=isNull" json:"is_null,omitempty"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	Error            *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Message          []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	CallId           int32                                            `protobuf:"varint,4,opt,name=call_id,json=callId" json:"call_id,omitempty"`
	Row              []*Response_Row                                  `protobuf:"bytes,5,rep,name=row" json:"row,omitempty"`
	OutputParameters map[string]*dstore_values.Value                  `protobuf:"bytes,6,rep,name=output_parameters,json=outputParameters" json:"output_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MetaInformation  []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,7,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetError() *dstore_engine_error.Error {
	if m != nil {
		return m.Error
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

func (m *Response) GetOutputParameters() map[string]*dstore_values.Value {
	if m != nil {
		return m.OutputParameters
	}
	return nil
}

func (m *Response) GetMetaInformation() []*dstore_engine_metainformation.MetaInformation {
	if m != nil {
		return m.MetaInformation
	}
	return nil
}

type Response_Row struct {
	RowId   int32                           `protobuf:"varint,1,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Columns map[string]*dstore_values.Value `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *Response_Row) GetColumns() map[string]*dstore_values.Value {
	if m != nil {
		return m.Columns
	}
	return nil
}

func init() {
	proto.RegisterType((*Call)(nil), "dstore.engine.procedure.Call")
	proto.RegisterType((*Parameter)(nil), "dstore.engine.procedure.Parameter")
	proto.RegisterType((*Response)(nil), "dstore.engine.procedure.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.procedure.Response.Row")
}

var fileDescriptor0 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0xcd, 0x26, 0xdb, 0x3d, 0xeb, 0xcf, 0x3a, 0xb6, 0x6c, 0x8c, 0x28, 0x65, 0xa5, 0x50,
	0x14, 0x52, 0x59, 0x2f, 0x5a, 0xbc, 0x12, 0x4b, 0x2f, 0x0a, 0x76, 0x2d, 0x73, 0x21, 0xd4, 0x9b,
	0x10, 0x9b, 0xb1, 0x04, 0x93, 0x99, 0x65, 0x66, 0x62, 0xe9, 0x1b, 0xf8, 0x1e, 0xbe, 0x86, 0x0f,
	0xe7, 0xfc, 0x25, 0x9b, 0xac, 0x2e, 0x0a, 0xbd, 0xd9, 0x9c, 0xdf, 0xef, 0x9c, 0xf3, 0x9d, 0x33,
	0x0b, 0xcf, 0x72, 0x21, 0x19, 0x27, 0x87, 0x84, 0x5e, 0x17, 0x94, 0x1c, 0x2e, 0x39, 0xbb, 0x22,
	0x79, 0xcd, 0x49, 0xa2, 0x24, 0xc9, 0xd0, 0xd4, 0xba, 0x13, 0xeb, 0x4e, 0x5a, 0x77, 0xfc, 0xd8,
	0xe5, 0x7d, 0xcf, 0xca, 0x9a, 0x08, 0x1b, 0x1d, 0x3f, 0xe9, 0x83, 0x11, 0xce, 0x19, 0x77, 0xae,
	0xa7, 0x7d, 0x57, 0x45, 0x84, 0xc8, 0xae, 0x5d, 0x95, 0xf8, 0xc5, 0xba, 0x53, 0x66, 0x05, 0xfd,
	0xca, 0x78, 0x95, 0xc9, 0x82, 0x51, 0x1b, 0x34, 0xfb, 0xe1, 0xc1, 0xe0, 0x24, 0x2b, 0x4b, 0xb4,
	0x0f, 0x0f, 0xda, 0x3e, 0x52, 0x9a, 0x55, 0x24, 0xf2, 0xf6, 0xbc, 0x83, 0x11, 0xbe, 0xdf, 0x5a,
	0x17, 0xca, 0x88, 0xde, 0xc1, 0x68, 0x99, 0x71, 0x25, 0x49, 0xc2, 0xa3, 0xad, 0x3d, 0xff, 0x60,
	0x3c, 0x9f, 0x25, 0x1b, 0xc6, 0x49, 0x2e, 0x9a, 0x48, 0xbc, 0x4a, 0x42, 0x53, 0x18, 0x5e, 0xa9,
	0x82, 0x69, 0x91, 0x47, 0xbe, 0xaa, 0x10, 0xe0, 0x50, 0xab, 0x67, 0xf9, 0x6c, 0x01, 0xa3, 0x36,
	0x01, 0x21, 0x18, 0x74, 0x9a, 0x30, 0x32, 0xda, 0x81, 0xc0, 0x10, 0xa3, 0xea, 0x6a, 0xa3, 0x55,
	0x34, 0x5e, 0x21, 0x52, 0x5a, 0x97, 0xa5, 0xc1, 0xdb, 0xc6, 0x61, 0x21, 0x16, 0x4a, 0x9b, 0xfd,
	0x0c, 0x60, 0x1b, 0x13, 0xb1, 0x64, 0x54, 0x10, 0xf4, 0x1a, 0x02, 0x43, 0x9c, 0x01, 0x1c, 0xcf,
	0xe3, 0xb5, 0x9e, 0x2d, 0xa9, 0xa7, 0xfa, 0x17, 0xdb, 0x40, 0x74, 0x0c, 0x43, 0xc7, 0xa7, 0xc2,
	0xd5, 0x73, 0x3e, 0x5f, 0xcb, 0x69, 0xd8, 0x3e, 0xb7, 0x5f, 0xdc, 0x84, 0x77, 0x27, 0x1c, 0x74,
	0x27, 0x44, 0x47, 0xe0, 0x73, 0x76, 0x13, 0x05, 0x06, 0x6e, 0x7f, 0x23, 0x6d, 0x4d, 0xd3, 0x09,
	0x66, 0x37, 0x58, 0x67, 0xa0, 0x1c, 0x1e, 0xb1, 0x5a, 0x2e, 0x6b, 0x99, 0xb6, 0x3c, 0x8a, 0x28,
	0x34, 0x30, 0x47, 0xff, 0x86, 0xf9, 0x68, 0x52, 0x5b, 0x6e, 0xc5, 0x29, 0x95, 0xfc, 0x16, 0x4f,
	0xd8, 0x9a, 0x19, 0x5d, 0xc2, 0x44, 0x1f, 0x49, 0xda, 0xb9, 0x92, 0x68, 0x68, 0x8a, 0x24, 0x7f,
	0x8c, 0xde, 0xbf, 0xa5, 0x73, 0xa5, 0x9f, 0xad, 0x74, 0xfc, 0xb0, 0xea, 0x1b, 0xe2, 0x4b, 0xd8,
	0xfd, 0x6b, 0x17, 0x68, 0x02, 0xfe, 0x37, 0x72, 0xeb, 0xd6, 0xac, 0x45, 0xf4, 0xb2, 0xbb, 0xe5,
	0xf1, 0x7c, 0xa7, 0x29, 0xed, 0xde, 0xc4, 0x27, 0xfd, 0x71, 0xbb, 0x7f, 0xbb, 0x75, 0xec, 0xc5,
	0xbf, 0x3c, 0xf0, 0x15, 0x51, 0x68, 0x17, 0x42, 0x45, 0x95, 0x26, 0xdd, 0x33, 0xa4, 0x07, 0x4a,
	0x53, 0x9c, 0x7f, 0x50, 0xcb, 0x60, 0x65, 0x5d, 0x51, 0xe1, 0xce, 0x75, 0xfe, 0x5f, 0xbc, 0x27,
	0x27, 0x36, 0xc9, 0x72, 0xd5, 0x40, 0xc4, 0x17, 0x70, 0xaf, 0xeb, 0xb8, 0x7b, 0xfb, 0xef, 0x5f,
	0xc1, 0xa4, 0x60, 0xfd, 0x96, 0x3e, 0x4f, 0x37, 0xfc, 0x7d, 0x7c, 0x09, 0xcd, 0xa3, 0x7d, 0xf3,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x5d, 0x71, 0x18, 0x60, 0x04, 0x00, 0x00,
}