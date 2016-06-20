// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_ModifyLockedNodeCharacs_Ad.proto
// DO NOT EDIT!

/*
Package im_ModifyLockedNodeCharacs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_ModifyLockedNodeCharacs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_ModifyLockedNodeCharacs_Ad

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
	UserId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull               bool                        `protobuf:"varint,1001,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	NodeCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull bool                        `protobuf:"varint,1002,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
	LockStatus               *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=lock_status,json=lockStatus" json:"lock_status,omitempty"`
	LockStatusNull           bool                        `protobuf:"varint,1003,opt,name=lock_status_null,json=lockStatusNull" json:"lock_status_null,omitempty"`
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

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetLockStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.LockStatus
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_ModifyLockedNodeCharacs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_ModifyLockedNodeCharacs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_ModifyLockedNodeCharacs_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x51, 0x8b, 0x13, 0x31,
	0x10, 0xe6, 0x5a, 0xdb, 0x1e, 0x73, 0xa2, 0x47, 0x94, 0x63, 0x6d, 0x51, 0xf4, 0x7c, 0x51, 0x84,
	0x54, 0x54, 0x44, 0xe1, 0x10, 0x54, 0x7c, 0x28, 0xda, 0xa2, 0x11, 0x05, 0x7d, 0x59, 0xd6, 0xcd,
	0x5c, 0x0d, 0xb6, 0xc9, 0x91, 0x64, 0x3d, 0xfc, 0x17, 0xfe, 0x21, 0x7f, 0x8c, 0x8f, 0xea, 0x9f,
	0x70, 0xb2, 0x49, 0xed, 0xed, 0x7a, 0x78, 0x78, 0x2f, 0xed, 0x4e, 0xe6, 0xfb, 0xe6, 0x9b, 0xcc,
	0x37, 0x81, 0x3d, 0xe9, 0xbc, 0xb1, 0x38, 0x46, 0x3d, 0x57, 0x1a, 0xc7, 0x07, 0xd6, 0x94, 0x28,
	0x2b, 0x8b, 0x6e, 0xac, 0x96, 0xf9, 0xd4, 0x48, 0xb5, 0xff, 0xe5, 0x85, 0x29, 0x3f, 0xa1, 0x9c,
	0x19, 0x89, 0x4f, 0x3f, 0x16, 0xb6, 0x28, 0x5d, 0xfe, 0x58, 0x72, 0x02, 0x7a, 0xc3, 0x6e, 0x45,
	0x36, 0x8f, 0x6c, 0xfe, 0x4f, 0xca, 0xf0, 0x42, 0x92, 0xfa, 0x5c, 0x2c, 0x2a, 0x74, 0xb1, 0xc2,
	0xf0, 0x52, 0x53, 0x1f, 0xad, 0x35, 0x36, 0xa5, 0x46, 0xcd, 0xd4, 0x12, 0x9d, 0x2b, 0xe6, 0x98,
	0x92, 0xd7, 0xdb, 0x49, 0x5f, 0x28, 0xbd, 0x6f, 0xec, 0xb2, 0xf0, 0xca, 0xe8, 0x08, 0xda, 0xfd,
	0xde, 0x01, 0x78, 0x49, 0xfa, 0x94, 0x45, 0xeb, 0xd8, 0x3d, 0x18, 0x54, 0x0e, 0x6d, 0xae, 0x64,
	0xb6, 0x71, 0x75, 0xe3, 0xc6, 0xd6, 0x9d, 0x11, 0x4f, 0xfd, 0xa7, 0x96, 0x94, 0xf6, 0x38, 0x47,
	0xfb, 0x36, 0x44, 0xa2, 0x1f, 0xb0, 0x13, 0xc9, 0xae, 0xc1, 0xd9, 0xc4, 0xca, 0x75, 0xb5, 0x58,
	0x64, 0x3f, 0x06, 0xc4, 0xdd, 0x14, 0x10, 0xd3, 0x33, 0x3a, 0x62, 0xaf, 0x60, 0x47, 0xd3, 0x5d,
	0xf3, 0xb2, 0xbe, 0x2c, 0x69, 0x29, 0xe7, 0x55, 0x19, 0x74, 0x3a, 0x27, 0xeb, 0x5c, 0xd4, 0x7f,
	0xc6, 0xb4, 0x62, 0x92, 0xea, 0x23, 0x18, 0x1d, 0x5f, 0x32, 0x36, 0xf1, 0x33, 0x36, 0x91, 0x1d,
	0xc7, 0xad, 0x5b, 0xda, 0x83, 0xad, 0x05, 0x99, 0x90, 0x3b, 0x5f, 0xf8, 0xca, 0x65, 0xdd, 0x93,
	0xfb, 0x80, 0x80, 0x7f, 0x5d, 0xc3, 0xd9, 0x4d, 0xd8, 0x3e, 0xc2, 0x8e, 0x92, 0xbf, 0xa2, 0xe4,
	0xb9, 0x35, 0x2c, 0x08, 0xed, 0x7e, 0xeb, 0xc0, 0xa6, 0x40, 0x77, 0x60, 0xb4, 0x43, 0x76, 0x1b,
	0x7a, 0xb5, 0x83, 0x69, 0xbe, 0x43, 0xde, 0xdc, 0x8f, 0xe8, 0xee, 0xb3, 0xf0, 0x2b, 0x22, 0x90,
	0xbd, 0x83, 0xed, 0xe0, 0x5d, 0x7e, 0xc4, 0x3c, 0x1a, 0x5a, 0x97, 0xc8, 0xbc, 0x45, 0x6e, 0x5b,
	0x3c, 0xa5, 0x78, 0xb2, 0x8e, 0xc5, 0xf9, 0x65, 0xf3, 0x80, 0x3d, 0x80, 0x41, 0xda, 0x19, 0xba,
	0x7e, 0xa8, 0x78, 0xe5, 0xaf, 0x8a, 0x71, 0xa3, 0xa6, 0xf1, 0x5f, 0xac, 0xe0, 0xec, 0x39, 0x74,
	0xad, 0x39, 0xcc, 0xce, 0xd4, 0xac, 0x87, 0xfc, 0x3f, 0x96, 0x9c, 0xaf, 0x46, 0xc1, 0x85, 0x39,
	0x14, 0xa1, 0xca, 0xf0, 0x32, 0x74, 0xe9, 0x9b, 0xed, 0x40, 0x9f, 0xa2, 0xb0, 0x13, 0x5f, 0x67,
	0x34, 0x9c, 0x9e, 0xe8, 0x51, 0x38, 0x91, 0x4f, 0xde, 0xc0, 0x48, 0x99, 0x96, 0xc4, 0xfa, 0x15,
	0xbe, 0xbf, 0x7f, 0xba, 0xf7, 0xf9, 0xa1, 0x5f, 0xbf, 0x80, 0xbb, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x9c, 0xa6, 0xc1, 0xce, 0xe0, 0x03, 0x00, 0x00,
}
