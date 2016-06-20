// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetLockedPersonTypes_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetLockedPersonTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetLockedPersonTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetLockedPersonTypes_Ad

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
	UserId           *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull       bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	PersonTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
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
	RowId                 int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	UserName              *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	PersonTypeDescription *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=person_type_description,json=personTypeDescription" json:"person_type_description,omitempty"`
	UserId                *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	PersonTypeId          *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetPersonTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PersonTypeDescription
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetLockedPersonTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetLockedPersonTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetLockedPersonTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x9a, 0xe6, 0x87, 0xa1, 0x82, 0xca, 0x15, 0xb0, 0x24, 0x12, 0x82, 0xf2, 0xd2, 0x27,
	0x07, 0x51, 0x40, 0xe5, 0xb1, 0x15, 0x08, 0x2a, 0xd1, 0x55, 0x65, 0x10, 0x12, 0xbc, 0xac, 0x96,
	0xec, 0x10, 0x59, 0x24, 0xf6, 0xca, 0x76, 0xa8, 0xb8, 0x05, 0x7f, 0x77, 0xe0, 0x24, 0xdc, 0x81,
	0x2b, 0xc0, 0x29, 0x18, 0xaf, 0xdd, 0x26, 0xbb, 0x15, 0x6a, 0xe9, 0x4b, 0xa2, 0xf1, 0x7c, 0xdf,
	0xcc, 0xf8, 0xfb, 0xc6, 0x0b, 0x8f, 0x0b, 0xeb, 0xb4, 0xc1, 0x11, 0xaa, 0x89, 0x54, 0x38, 0x2a,
	0x8d, 0x1e, 0x63, 0x31, 0x37, 0x68, 0x47, 0xe5, 0x2c, 0x7b, 0x86, 0xee, 0x85, 0x1e, 0x7f, 0xc0,
	0xe2, 0x10, 0x8d, 0xd5, 0xea, 0xd5, 0xa7, 0x12, 0x6d, 0xb6, 0x5b, 0x70, 0x42, 0x39, 0xcd, 0xb6,
	0x02, 0x95, 0x07, 0x2a, 0xff, 0x37, 0x7e, 0xb0, 0x11, 0x9b, 0x7c, 0xcc, 0xa7, 0x73, 0xb4, 0x81,
	0x3e, 0xb8, 0x59, 0xef, 0x8c, 0xc6, 0x68, 0x13, 0x53, 0xc3, 0x7a, 0x6a, 0x86, 0xd6, 0xe6, 0x13,
	0x8c, 0xc9, 0xbb, 0xcd, 0xa4, 0xcb, 0xa5, 0x7a, 0xaf, 0xcd, 0x2c, 0x77, 0x52, 0xab, 0x00, 0xda,
	0xfc, 0xd5, 0x02, 0x38, 0xcc, 0x4d, 0x4e, 0x59, 0x1a, 0x85, 0x3d, 0x80, 0xde, 0xdc, 0xa2, 0xc9,
	0x64, 0x91, 0xb4, 0x6e, 0xb7, 0xb6, 0x2e, 0xdf, 0x1f, 0xf2, 0x38, 0x7c, 0x1c, 0x49, 0x2a, 0x87,
	0x13, 0x34, 0xaf, 0x7d, 0x24, 0xba, 0x1e, 0xbb, 0x5f, 0xb0, 0x3b, 0xb0, 0x16, 0x59, 0x99, 0x9a,
	0x4f, 0xa7, 0xc9, 0xef, 0x1e, 0x71, 0xfb, 0x02, 0x42, 0x3a, 0xa5, 0x23, 0xb6, 0x0b, 0x57, 0xca,
	0xea, 0xae, 0x99, 0xa3, 0xcb, 0xfa, 0xfa, 0x2b, 0x67, 0xd7, 0x5f, 0x2b, 0x4f, 0xe4, 0xa1, 0x2e,
	0x1c, 0x36, 0xea, 0x25, 0x42, 0xb3, 0x3f, 0xa1, 0xd9, 0xfa, 0x32, 0xd6, 0xb7, 0xdc, 0xfc, 0xb9,
	0x0a, 0x7d, 0x81, 0xb6, 0xd4, 0xca, 0x22, 0xbb, 0x07, 0x9d, 0x4a, 0xb8, 0x78, 0xad, 0x01, 0xaf,
	0x7b, 0x12, 0x44, 0x7d, 0xea, 0x7f, 0x45, 0x00, 0xb2, 0x37, 0xb0, 0xee, 0x25, 0xcb, 0x96, 0x34,
	0xa3, 0x99, 0xdb, 0x44, 0xe6, 0x0d, 0x72, 0x53, 0xd9, 0x03, 0x8a, 0xf7, 0x17, 0xb1, 0xb8, 0x3a,
	0xab, 0x1f, 0xb0, 0x1d, 0xe8, 0x45, 0xab, 0x92, 0x76, 0x55, 0xf1, 0xd6, 0xa9, 0x8a, 0xc1, 0xc8,
	0x83, 0xf0, 0x2f, 0x8e, 0xe1, 0xec, 0x39, 0xb4, 0x8d, 0x3e, 0x4a, 0x56, 0x2b, 0xd6, 0x23, 0x7e,
	0xde, 0xc5, 0xe2, 0xc7, 0x3a, 0x70, 0xa1, 0x8f, 0x84, 0x2f, 0x31, 0xf8, 0xb1, 0x02, 0x6d, 0x0a,
	0xd8, 0x75, 0xe8, 0x52, 0xe8, 0x0d, 0xf9, 0x9c, 0x92, 0x34, 0x1d, 0xd1, 0xa1, 0x90, 0xd4, 0xde,
	0x81, 0x4b, 0x95, 0xa7, 0x8a, 0x56, 0x23, 0xf9, 0x92, 0xd6, 0x55, 0x8b, 0x66, 0x59, 0x67, 0xa4,
	0x9a, 0x04, 0xaf, 0xfa, 0x1e, 0x9d, 0x12, 0x98, 0xbd, 0x84, 0x1b, 0xcb, 0x3e, 0x15, 0x68, 0xc7,
	0x46, 0x96, 0x95, 0x7e, 0x5f, 0xcf, 0xae, 0x73, 0x6d, 0xe1, 0xe3, 0x93, 0x05, 0x93, 0x3d, 0x5c,
	0x2c, 0xe6, 0xb7, 0xf4, 0xfc, 0x9b, 0xb9, 0x77, 0x6a, 0xed, 0xbe, 0xa7, 0xff, 0xb9, 0x77, 0x7b,
	0x02, 0x86, 0x52, 0x37, 0xa5, 0x3e, 0x79, 0xfe, 0x6f, 0xb7, 0x2f, 0xf0, 0x61, 0x78, 0xd7, 0xad,
	0x5e, 0xdf, 0xf6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x92, 0xa4, 0xd9, 0x56, 0x04, 0x00,
	0x00,
}
