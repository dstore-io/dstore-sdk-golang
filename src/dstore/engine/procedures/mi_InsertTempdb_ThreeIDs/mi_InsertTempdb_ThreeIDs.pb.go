// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_InsertTempdb_ThreeIDs.proto
// DO NOT EDIT!

/*
Package mi_InsertTempdb_ThreeIDs is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_InsertTempdb_ThreeIDs.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_InsertTempdb_ThreeIDs

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
	IdList1     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=id_list1,json=idList1" json:"id_list1,omitempty"`
	IdList1Null bool                        `protobuf:"varint,1001,opt,name=id_list1_null,json=idList1Null" json:"id_list1_null,omitempty"`
	IdList2     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=id_list2,json=idList2" json:"id_list2,omitempty"`
	IdList2Null bool                        `protobuf:"varint,1002,opt,name=id_list2_null,json=idList2Null" json:"id_list2_null,omitempty"`
	IdList3     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=id_list3,json=idList3" json:"id_list3,omitempty"`
	IdList3Null bool                        `protobuf:"varint,1003,opt,name=id_list3_null,json=idList3Null" json:"id_list3_null,omitempty"`
	Delete      *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull  bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetIdList1() *dstore_values.StringValue {
	if m != nil {
		return m.IdList1
	}
	return nil
}

func (m *Parameters) GetIdList2() *dstore_values.StringValue {
	if m != nil {
		return m.IdList2
	}
	return nil
}

func (m *Parameters) GetIdList3() *dstore_values.StringValue {
	if m != nil {
		return m.IdList3
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_InsertTempdb_ThreeIDs.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_InsertTempdb_ThreeIDs.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_InsertTempdb_ThreeIDs.Response.Row")
}

var fileDescriptor0 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xdd, 0x8a, 0xd4, 0x30,
	0x14, 0xc7, 0xd9, 0xa9, 0xf3, 0xc1, 0x19, 0x44, 0x89, 0x20, 0xb5, 0x83, 0xb2, 0xec, 0x82, 0x78,
	0x95, 0x71, 0x5b, 0x45, 0xaf, 0x45, 0x85, 0x82, 0xbb, 0x2c, 0x61, 0x11, 0xf4, 0xa6, 0x74, 0xcc,
	0x71, 0x0c, 0xb4, 0xc9, 0x90, 0x74, 0xdc, 0xd7, 0xf0, 0x71, 0x7c, 0x25, 0x3f, 0xee, 0xbd, 0x35,
	0x69, 0x32, 0x4c, 0x5b, 0x11, 0x66, 0x6f, 0xda, 0x9c, 0x9c, 0xff, 0xff, 0x97, 0x70, 0xce, 0x09,
	0xbc, 0xe0, 0xa6, 0x51, 0x1a, 0x97, 0x28, 0xd7, 0x42, 0xe2, 0x72, 0xa3, 0xd5, 0x27, 0xe4, 0x5b,
	0x8d, 0x66, 0x59, 0x8b, 0x22, 0x97, 0x06, 0x75, 0x73, 0x85, 0xf5, 0x86, 0xaf, 0x8a, 0xab, 0x2f,
	0x1a, 0x31, 0x7f, 0x6d, 0xa8, 0xd5, 0x34, 0x8a, 0x3c, 0xf6, 0x46, 0xea, 0x8d, 0xf4, 0x7f, 0xea,
	0xe4, 0x5e, 0x38, 0xe0, 0x6b, 0x59, 0x6d, 0x31, 0x98, 0x93, 0x07, 0xfd, 0x53, 0x51, 0x6b, 0xa5,
	0x43, 0x6a, 0xd1, 0x4f, 0xd5, 0x68, 0x4c, 0xb9, 0xc6, 0x90, 0x3c, 0x1d, 0x26, 0x9b, 0x52, 0xc8,
	0xcf, 0x4a, 0xd7, 0x65, 0x23, 0x94, 0xf4, 0xa2, 0x93, 0x3f, 0x23, 0x80, 0xcb, 0x52, 0x97, 0x36,
	0x8b, 0xda, 0x90, 0xe7, 0x30, 0x13, 0xbc, 0xa8, 0x84, 0x69, 0xce, 0xe2, 0xa3, 0xe3, 0xa3, 0x27,
	0xf3, 0x34, 0xa1, 0xe1, 0xee, 0xe1, 0x4e, 0xa6, 0xd1, 0x42, 0xae, 0xdf, 0xbb, 0x80, 0x4d, 0x05,
	0x7f, 0xe7, 0xa4, 0xe4, 0x14, 0x6e, 0xef, 0x6c, 0x85, 0xdc, 0x56, 0x55, 0xfc, 0x63, 0x6a, 0xcd,
	0x33, 0x36, 0x0f, 0x82, 0x0b, 0xbb, 0xd7, 0x61, 0xa7, 0xf1, 0xe8, 0x50, 0x76, 0xda, 0x61, 0xa7,
	0x9e, 0xfd, 0xb3, 0xc7, 0x4e, 0x07, 0xec, 0x2c, 0x8e, 0x0e, 0x65, 0x67, 0x1d, 0x76, 0xe6, 0xd9,
	0xbf, 0x7a, 0xec, 0xac, 0x65, 0x67, 0x30, 0xe1, 0x58, 0xd9, 0xfa, 0xc4, 0xb7, 0x5a, 0xf2, 0x62,
	0x40, 0x5e, 0x29, 0x55, 0x61, 0x29, 0x3d, 0x3a, 0x48, 0xc9, 0x31, 0xcc, 0xfd, 0xca, 0x73, 0x7f,
	0x7b, 0x2e, 0xf8, 0x3d, 0x87, 0x3d, 0xf9, 0x3e, 0x82, 0x19, 0x43, 0xb3, 0x51, 0x76, 0x12, 0xc8,
	0x53, 0x18, 0xb7, 0x7d, 0x1d, 0x16, 0x3d, 0x0c, 0x8c, 0xef, 0xf9, 0x1b, 0xf7, 0x65, 0x5e, 0x48,
	0x3e, 0xc0, 0x5d, 0xd7, 0xd1, 0xa2, 0xd3, 0x52, 0x5b, 0xd5, 0xc8, 0x9a, 0xe9, 0xc0, 0x3c, 0x6c,
	0xfc, 0xb9, 0x8d, 0xf3, 0x7d, 0xcc, 0xee, 0xd4, 0xfd, 0x0d, 0xf2, 0x12, 0xa6, 0x61, 0x92, 0x6c,
	0x2d, 0x1d, 0xf1, 0xd1, 0x3f, 0x44, 0x3f, 0x67, 0xe7, 0xfe, 0xcf, 0x76, 0x72, 0xf2, 0x16, 0x22,
	0xad, 0xae, 0x6d, 0x9d, 0x9c, 0xeb, 0x19, 0x3d, 0x6c, 0xea, 0xe9, 0xae, 0x0a, 0x94, 0xa9, 0x6b,
	0xe6, 0x00, 0xc9, 0x43, 0x88, 0xec, 0x9a, 0xdc, 0x87, 0x89, 0x8d, 0x0a, 0xc1, 0xe3, 0x6f, 0x17,
	0xb6, 0x2e, 0x63, 0x36, 0xb6, 0x61, 0xce, 0x5f, 0x5d, 0xc2, 0x42, 0xa8, 0x01, 0x7d, 0xff, 0x18,
	0x3f, 0x9e, 0xdd, 0xf8, 0x99, 0xae, 0x26, 0xed, 0x6b, 0xc8, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0x3a, 0x12, 0xde, 0xe2, 0x03, 0x00, 0x00,
}
