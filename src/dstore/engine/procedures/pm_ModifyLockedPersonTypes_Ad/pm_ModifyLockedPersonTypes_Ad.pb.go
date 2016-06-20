// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_ModifyLockedPersonTypes_Ad.proto
// DO NOT EDIT!

/*
Package pm_ModifyLockedPersonTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_ModifyLockedPersonTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_ModifyLockedPersonTypes_Ad

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
	Locked           *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=locked" json:"locked,omitempty"`
	LockedNull       bool                        `protobuf:"varint,1003,opt,name=locked_null,json=lockedNull" json:"locked_null,omitempty"`
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

func (m *Parameters) GetLocked() *dstore_values.BooleanValue {
	if m != nil {
		return m.Locked
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_ModifyLockedPersonTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_ModifyLockedPersonTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_ModifyLockedPersonTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x65, 0x1b, 0x9b, 0x2e, 0x77, 0x17, 0x5d, 0x66, 0x41, 0x62, 0x8a, 0xb2, 0xae, 0x2f, 0x82,
	0x30, 0x15, 0x57, 0x44, 0xc1, 0x97, 0x15, 0x7c, 0x28, 0xda, 0x52, 0x06, 0x15, 0xf4, 0x25, 0xa4,
	0x9d, 0xdb, 0x12, 0x4c, 0x66, 0xc2, 0x4c, 0x62, 0xe9, 0x5f, 0xf8, 0x27, 0x7e, 0x81, 0x1f, 0xa4,
	0xfe, 0x84, 0x33, 0x99, 0x69, 0x6d, 0xa2, 0x28, 0xee, 0x4b, 0x9b, 0x3b, 0xf7, 0x9e, 0x73, 0xe6,
	0xde, 0x73, 0x07, 0x9e, 0x73, 0x5d, 0x49, 0x85, 0x23, 0x14, 0xab, 0x4c, 0xe0, 0xa8, 0x54, 0x72,
	0x81, 0xbc, 0x56, 0xa8, 0x47, 0x65, 0x91, 0x4c, 0x24, 0xcf, 0x96, 0x9b, 0xd7, 0x72, 0xf1, 0x11,
	0xf9, 0x0c, 0x95, 0x96, 0xe2, 0xcd, 0xa6, 0x44, 0x9d, 0x5c, 0x72, 0x6a, 0x0a, 0x2b, 0x49, 0x1e,
	0x38, 0x34, 0x75, 0x68, 0xfa, 0x57, 0x48, 0x7c, 0xea, 0xa5, 0x3e, 0xa5, 0x79, 0x8d, 0xda, 0x31,
	0xc4, 0xb7, 0xda, 0xfa, 0xa8, 0x94, 0x54, 0x3e, 0x35, 0x6c, 0xa7, 0x0a, 0xd4, 0x3a, 0x5d, 0xa1,
	0x4f, 0xde, 0xeb, 0x26, 0xab, 0x34, 0x13, 0x4b, 0xa9, 0x8a, 0xb4, 0xca, 0xa4, 0x70, 0x45, 0xe7,
	0x5f, 0x7a, 0x00, 0xb3, 0x54, 0xa5, 0x26, 0x6b, 0xae, 0x42, 0x1e, 0xc3, 0xa0, 0xd6, 0xa8, 0x92,
	0x8c, 0x47, 0x07, 0x67, 0x07, 0xf7, 0x8f, 0x1e, 0x0d, 0xa9, 0xbf, 0xbf, 0xbf, 0x52, 0x26, 0x2a,
	0x5c, 0xa1, 0x7a, 0x67, 0x23, 0x16, 0xda, 0xda, 0x31, 0x27, 0x77, 0xe1, 0xd8, 0xa3, 0x12, 0x51,
	0xe7, 0x79, 0xf4, 0x6d, 0x60, 0xb0, 0x87, 0x0c, 0x5c, 0x7a, 0x6a, 0x8e, 0xc8, 0x25, 0x5c, 0x2f,
	0x9b, 0x5e, 0x93, 0xca, 0x34, 0x6b, 0xf9, 0x7b, 0xff, 0xe6, 0x3f, 0x2e, 0x77, 0xe3, 0x31, 0x2a,
	0x14, 0x4e, 0xdb, 0x14, 0x4e, 0xec, 0xbb, 0x13, 0x3b, 0xd9, 0xaf, 0x6d, 0x24, 0x2f, 0x20, 0xcc,
	0x9b, 0x21, 0x47, 0xc1, 0x1f, 0xa5, 0xe6, 0x52, 0xe6, 0x98, 0x0a, 0xdf, 0x8a, 0x2b, 0x25, 0x67,
	0x70, 0xe4, 0xbe, 0x1c, 0xf9, 0x0f, 0xdf, 0x89, 0x3b, 0xb3, 0xb4, 0xe7, 0x5f, 0x7b, 0x70, 0xc8,
	0x50, 0x97, 0x52, 0x68, 0x24, 0x0f, 0xa1, 0xdf, 0xf8, 0xe1, 0xa7, 0x15, 0xd3, 0xb6, 0xdb, 0xce,
	0xab, 0x97, 0xf6, 0x97, 0xb9, 0x42, 0xf2, 0x1e, 0x4e, 0xac, 0x13, 0xc9, 0x9e, 0x15, 0x66, 0x14,
	0x81, 0x01, 0xd3, 0x0e, 0xb8, 0x6b, 0xd8, 0xc4, 0xc4, 0xe3, 0x5f, 0x31, 0xbb, 0x51, 0xb4, 0x0f,
	0xc8, 0x53, 0x18, 0xf8, 0x0d, 0x30, 0x1d, 0x5b, 0xc6, 0x3b, 0xbf, 0x31, 0xba, 0xfd, 0x98, 0xb8,
	0x7f, 0xb6, 0x2d, 0x27, 0xaf, 0x20, 0x50, 0x72, 0x1d, 0x5d, 0x6b, 0x50, 0xcf, 0xe8, 0x7f, 0xac,
	0x2c, 0xdd, 0x8e, 0x82, 0x32, 0xb9, 0x66, 0x96, 0x25, 0xbe, 0x0d, 0x81, 0xf9, 0x26, 0x37, 0x21,
	0x34, 0x91, 0x75, 0xfa, 0xf3, 0xd4, 0x0c, 0xa7, 0xcf, 0xfa, 0x26, 0x1c, 0xf3, 0x17, 0x6f, 0x61,
	0x98, 0xc9, 0xae, 0xc4, 0xee, 0x4d, 0x7d, 0x78, 0x72, 0xb5, 0xd7, 0x36, 0x0f, 0x9b, 0x7d, 0xbe,
	0xf8, 0x19, 0x00, 0x00, 0xff, 0xff, 0x74, 0xca, 0x54, 0xd3, 0xae, 0x03, 0x00, 0x00,
}