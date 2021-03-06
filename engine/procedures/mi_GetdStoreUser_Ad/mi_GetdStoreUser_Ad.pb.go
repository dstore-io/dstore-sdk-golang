// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetdStoreUser_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetdStoreUser_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetdStoreUser_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetdStoreUser_Ad

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
	OnlyAdmins     *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=only_admins,json=onlyAdmins" json:"only_admins,omitempty"`
	OnlyAdminsNull bool                        `protobuf:"varint,1001,opt,name=only_admins_null,json=onlyAdminsNull" json:"only_admins_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetOnlyAdmins() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyAdmins
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
	RowId    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	UserName *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserId   *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=user_id,json=userId" json:"user_id,omitempty"`
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

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetdStoreUser_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetdStoreUser_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetdStoreUser_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_GetdStoreUser_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x5f, 0xeb, 0xd3, 0x30,
	0x14, 0x65, 0xd6, 0xfd, 0x31, 0x3f, 0xd0, 0x11, 0x41, 0x6a, 0x07, 0x32, 0xe7, 0x8b, 0xbe, 0xa4,
	0x32, 0x19, 0xf8, 0xe0, 0xcb, 0x04, 0x91, 0x3d, 0xac, 0x48, 0xfc, 0x03, 0xfa, 0x52, 0x32, 0x73,
	0x2d, 0xc1, 0x36, 0x19, 0x49, 0xeb, 0xf0, 0x4b, 0x88, 0xfa, 0x1d, 0xfc, 0x70, 0x7e, 0x0b, 0x93,
	0x26, 0x63, 0x6b, 0x1d, 0xe8, 0xcb, 0xc6, 0xcd, 0x39, 0xe7, 0xe6, 0x9e, 0x73, 0x53, 0xb4, 0xe4,
	0xa6, 0x56, 0x1a, 0x52, 0x90, 0x85, 0x90, 0x90, 0xee, 0xb5, 0xfa, 0x08, 0xbc, 0xd1, 0x60, 0xd2,
	0x4a, 0xe4, 0x2f, 0xa1, 0xe6, 0xaf, 0x1d, 0xfc, 0xd6, 0x80, 0xce, 0xd7, 0x9c, 0x58, 0xb8, 0x56,
	0xf8, 0xbe, 0xd7, 0x10, 0xaf, 0x21, 0x17, 0x88, 0xc9, 0xed, 0xd0, 0xf6, 0x0b, 0x2b, 0x1b, 0x30,
	0x5e, 0x97, 0xdc, 0xed, 0xde, 0x05, 0x5a, 0x2b, 0x1d, 0xa0, 0x59, 0x17, 0xaa, 0xc0, 0x18, 0x56,
	0x40, 0x00, 0x1f, 0xf4, 0xc1, 0x9a, 0x09, 0xf9, 0x49, 0xe9, 0x8a, 0xd5, 0x42, 0x49, 0x4f, 0x5a,
	0x34, 0x08, 0xbd, 0x62, 0x9a, 0x59, 0x10, 0xb4, 0xc1, 0xcf, 0xd0, 0x95, 0x92, 0xe5, 0xd7, 0x9c,
	0xf1, 0x4a, 0x48, 0x13, 0x0f, 0xe6, 0x83, 0x87, 0x57, 0xcb, 0x19, 0x09, 0x83, 0x87, 0xa9, 0x76,
	0x4a, 0x95, 0xc0, 0xe4, 0x3b, 0x57, 0x51, 0xe4, 0xf8, 0xeb, 0x96, 0x8e, 0x1f, 0xa1, 0xe9, 0x99,
	0x3a, 0x97, 0x4d, 0x59, 0xc6, 0xbf, 0xc7, 0xb6, 0xc7, 0x84, 0xde, 0x3c, 0xd1, 0x32, 0x7b, 0xbc,
	0xf8, 0x15, 0xa1, 0x09, 0x05, 0xb3, 0x57, 0xd2, 0x00, 0x7e, 0x8c, 0x86, 0xad, 0xa9, 0x70, 0x5f,
	0x42, 0xba, 0x41, 0x79, 0xc3, 0x2f, 0xdc, 0x2f, 0xf5, 0x44, 0xfc, 0x1e, 0x4d, 0x9d, 0x9d, 0xfc,
	0xcc, 0x4f, 0x7c, 0x6d, 0x1e, 0x59, 0x31, 0xe9, 0x89, 0xfb, 0xae, 0xb7, 0xb6, 0xde, 0x9c, 0x6a,
	0x7a, 0xab, 0xea, 0x1e, 0xe0, 0xa7, 0x68, 0x1c, 0x62, 0x8c, 0xa3, 0xb6, 0xe3, 0xbd, 0xbf, 0x3a,
	0xfa, 0x90, 0xb7, 0xfe, 0x9f, 0x1e, 0xe9, 0x78, 0x8d, 0x22, 0xad, 0x0e, 0xf1, 0xf5, 0x56, 0x95,
	0x92, 0x7f, 0x6e, 0x9b, 0x1c, 0x03, 0x20, 0x54, 0x1d, 0xa8, 0xd3, 0x26, 0xdf, 0x06, 0x28, 0xb2,
	0x05, 0xbe, 0x83, 0x46, 0xb6, 0xcc, 0x05, 0x8f, 0xbf, 0x67, 0x36, 0x93, 0x21, 0x1d, 0xda, 0x72,
	0xc3, 0xed, 0x70, 0x37, 0x1a, 0x27, 0x96, 0x76, 0x61, 0xf1, 0x8f, 0xac, 0x1b, 0x57, 0x58, 0x8f,
	0xa9, 0xb5, 0x90, 0x85, 0xdf, 0xce, 0xc4, 0xb1, 0x33, 0x4b, 0xc6, 0x2b, 0x34, 0x6e, 0x95, 0xb6,
	0xe5, 0xcf, 0xec, 0xe2, 0x5a, 0x85, 0xac, 0xa1, 0x00, 0xed, 0x85, 0x23, 0x47, 0xde, 0xf0, 0xe7,
	0x6f, 0xd0, 0x4c, 0xa8, 0x9e, 0x95, 0xd3, 0x63, 0xff, 0xb0, 0x2a, 0x94, 0xe1, 0x9f, 0x8f, 0x38,
	0xff, 0xcf, 0xef, 0x61, 0x37, 0x6a, 0xdf, 0xde, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xde, 0x0f, 0x16, 0x46, 0x03, 0x00, 0x00,
}
