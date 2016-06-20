// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetReturnCodeMessage.proto
// DO NOT EDIT!

/*
Package mi_GetReturnCodeMessage is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetReturnCodeMessage.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetReturnCodeMessage

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
	ReturnCode     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=return_code,json=returnCode" json:"return_code,omitempty"`
	ReturnCodeNull bool                        `protobuf:"varint,1001,opt,name=return_code_null,json=returnCodeNull" json:"return_code_null,omitempty"`
	LanguageId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetReturnCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
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
	RowId               int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ReturnCode          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=return_code,json=returnCode" json:"return_code,omitempty"`
	Description         *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=description" json:"description,omitempty"`
	DetailedDescription *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=detailed_description,json=detailedDescription" json:"detailed_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetReturnCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetDetailedDescription() *dstore_values.StringValue {
	if m != nil {
		return m.DetailedDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetReturnCodeMessage.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetReturnCodeMessage.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetReturnCodeMessage.Response.Row")
}

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0x14, 0x41,
	0x10, 0x65, 0xdd, 0x6c, 0x12, 0x3a, 0xa0, 0xa1, 0x23, 0x32, 0xce, 0x82, 0x48, 0x44, 0xd0, 0x4b,
	0x6f, 0x30, 0x20, 0x1e, 0xf4, 0xa2, 0x06, 0xc9, 0x21, 0xa3, 0xf4, 0x41, 0xd0, 0xcb, 0x30, 0x6e,
	0x97, 0x43, 0xc3, 0x4c, 0xf7, 0x52, 0xdd, 0x63, 0xfe, 0x81, 0x67, 0x3f, 0xfe, 0x9f, 0x77, 0xf5,
	0x4f, 0xd8, 0x3d, 0xdd, 0xe3, 0x7c, 0x88, 0x24, 0xb9, 0x0c, 0xd4, 0xd4, 0x7b, 0xf5, 0xaa, 0x5f,
	0x55, 0x91, 0xc7, 0xc2, 0x58, 0x8d, 0xb0, 0x02, 0x55, 0x4a, 0x05, 0xab, 0x0d, 0xea, 0x35, 0x88,
	0x06, 0xc1, 0xac, 0x6a, 0x99, 0xbf, 0x02, 0xcb, 0xc1, 0x36, 0xa8, 0x5e, 0x68, 0x01, 0x67, 0x60,
	0x4c, 0x51, 0x02, 0x73, 0x10, 0xab, 0xe9, 0xfd, 0xc0, 0x63, 0x81, 0xc7, 0xfe, 0x03, 0x4e, 0x0f,
	0x62, 0xf9, 0x4f, 0x45, 0xd5, 0x80, 0x09, 0xdc, 0xf4, 0xf6, 0x58, 0x13, 0x10, 0x35, 0xc6, 0xd4,
	0x72, 0x9c, 0xaa, 0x87, 0x9a, 0xe9, 0xbd, 0x69, 0xd2, 0x16, 0x52, 0x7d, 0xd4, 0x58, 0x17, 0x56,
	0x6a, 0x15, 0x40, 0x87, 0x3f, 0x66, 0x84, 0xbc, 0x29, 0xb0, 0x70, 0x59, 0x40, 0x43, 0x9f, 0x92,
	0x3d, 0x6c, 0xbb, 0xca, 0xd7, 0xae, 0xad, 0x64, 0x76, 0x77, 0xf6, 0x60, 0xef, 0xd1, 0x92, 0xc5,
	0xee, 0x63, 0x5b, 0x52, 0x59, 0x28, 0x01, 0xdf, 0xfa, 0x88, 0x13, 0xfc, 0xfb, 0x0a, 0xfa, 0x90,
	0xec, 0x0f, 0xd8, 0xb9, 0x6a, 0xaa, 0x2a, 0xf9, 0xb9, 0xe3, 0x6a, 0xec, 0xf2, 0xeb, 0x3d, 0x2c,
	0x73, 0xbf, 0xbd, 0x50, 0x55, 0xa8, 0xb2, 0x71, 0xed, 0xe6, 0x52, 0x24, 0xd7, 0x2e, 0x21, 0xd4,
	0xe1, 0x4f, 0x85, 0x17, 0x1a, 0xb0, 0x83, 0xd0, 0xaf, 0x28, 0xd4, 0xc3, 0xbc, 0xd0, 0xe1, 0xe7,
	0x2d, 0xb2, 0xcb, 0xc1, 0x6c, 0xb4, 0x32, 0x40, 0x8f, 0xc8, 0xa2, 0xb5, 0x2f, 0x3e, 0x2c, 0x65,
	0xe3, 0xb1, 0x04, 0x6b, 0x4f, 0xfc, 0x97, 0x07, 0x20, 0x7d, 0x47, 0xf6, 0xbd, 0x71, 0xf9, 0xc0,
	0x39, 0xd7, 0xec, 0xdc, 0x91, 0xd9, 0x84, 0x3c, 0xf5, 0xf7, 0xcc, 0xc5, 0xa7, 0x7d, 0xcc, 0x6f,
	0xd4, 0xe3, 0x1f, 0xf4, 0x09, 0xd9, 0x89, 0x03, 0x4b, 0xe6, 0x6d, 0xc5, 0x3b, 0xff, 0x54, 0x0c,
	0xe3, 0x8c, 0xdb, 0xc1, 0x3b, 0x38, 0x3d, 0x21, 0x73, 0xd4, 0xe7, 0xc9, 0x56, 0xcb, 0x3a, 0x66,
	0x97, 0xda, 0x2d, 0xd6, 0x99, 0xc0, 0xb8, 0x3e, 0xe7, 0x9e, 0x9f, 0xfe, 0x9e, 0x91, 0xb9, 0x0b,
	0xe8, 0x2d, 0xb2, 0xed, 0x42, 0x3f, 0x86, 0x2f, 0x99, 0xf3, 0x65, 0xc1, 0x17, 0x2e, 0x74, 0x2e,
	0x3f, 0x1b, 0x2f, 0xc3, 0xd7, 0xec, 0x6a, 0xdb, 0xe0, 0xe8, 0x02, 0xcc, 0x1a, 0xe5, 0xa6, 0x75,
	0xed, 0x5b, 0x36, 0xf6, 0x3c, 0xd2, 0x8d, 0x45, 0xa9, 0xca, 0xc0, 0x1e, 0xe2, 0x69, 0x46, 0x6e,
	0x0a, 0x6f, 0x69, 0x05, 0x22, 0x1f, 0xd6, 0xf9, 0x7e, 0x71, 0x9d, 0x83, 0x8e, 0xf8, 0xb2, 0xe7,
	0x3d, 0x7f, 0x4d, 0x96, 0x52, 0x4f, 0xbc, 0xea, 0xef, 0xf7, 0xfd, 0xd1, 0x55, 0x2f, 0xfb, 0xc3,
	0x76, 0x7b, 0x41, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xe3, 0xcc, 0x01, 0x14, 0x04,
	0x00, 0x00,
}
