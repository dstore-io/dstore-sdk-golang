// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampCondCriteriaTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampCondCriteriaTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampCondCriteriaTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampCondCriteriaTypes_Ad

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
	CriteriaTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=criteria_type_id,json=criteriaTypeId" json:"criteria_type_id,omitempty"`
	CriteriaTypeIdNull bool                        `protobuf:"varint,1001,opt,name=criteria_type_id_null,json=criteriaTypeIdNull" json:"criteria_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCriteriaTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CriteriaTypeId
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
	CriteraiType   *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=criterai_type,json=criteraiType" json:"criterai_type,omitempty"`
	CriteriaTypeId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=criteria_type_id,json=criteriaTypeId" json:"criteria_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCriteraiType() *dstore_values.StringValue {
	if m != nil {
		return m.CriteraiType
	}
	return nil
}

func (m *Response_Row) GetCriteriaTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CriteriaTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampCondCriteriaTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampCondCriteriaTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampCondCriteriaTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xdb, 0xaa, 0x13, 0x31,
	0x14, 0x65, 0x1c, 0x7b, 0x21, 0xde, 0x4a, 0x44, 0x19, 0xa7, 0x20, 0xa5, 0xbe, 0xf8, 0x20, 0xa9,
	0xd4, 0x07, 0x45, 0x10, 0xd4, 0x52, 0xa5, 0x60, 0x07, 0x09, 0x52, 0xd0, 0x97, 0x21, 0x76, 0xb6,
	0x43, 0xa0, 0x93, 0x0c, 0x49, 0x6a, 0xf1, 0x0b, 0x7c, 0x55, 0x7f, 0xc2, 0x6f, 0xf1, 0x53, 0xfc,
	0x8b, 0x93, 0x99, 0xa4, 0xb4, 0x33, 0xe7, 0x70, 0xa0, 0xe7, 0xa5, 0x65, 0xcf, 0xda, 0x6b, 0x65,
	0xaf, 0xb5, 0x13, 0xf4, 0x2a, 0xd3, 0x46, 0x2a, 0x98, 0x80, 0xc8, 0xb9, 0x80, 0x49, 0xa9, 0xe4,
	0x1a, 0xb2, 0xad, 0x02, 0x3d, 0x91, 0x45, 0xfa, 0x1e, 0xcc, 0x8c, 0x15, 0xe5, 0x4c, 0x8a, 0x6c,
	0xa6, 0xb8, 0x01, 0xc5, 0xd9, 0xa7, 0x1f, 0x25, 0xe8, 0xf4, 0x4d, 0x46, 0x6c, 0xa7, 0x91, 0xf8,
	0x89, 0xa3, 0x13, 0x47, 0x27, 0x97, 0x73, 0xe2, 0xbb, 0xfe, 0xb0, 0xef, 0x6c, 0xb3, 0x05, 0xed,
	0x24, 0xe2, 0x07, 0xcd, 0x09, 0x40, 0x29, 0xa9, 0x3c, 0x34, 0x6c, 0x42, 0x05, 0x68, 0xcd, 0x72,
	0xf0, 0xe0, 0xa3, 0x36, 0x68, 0x18, 0x17, 0xdf, 0xa4, 0x2a, 0x98, 0xe1, 0x52, 0xb8, 0xa6, 0xf1,
	0xcf, 0x00, 0xa1, 0x8f, 0x4c, 0x31, 0x8b, 0x82, 0xd2, 0x78, 0x8e, 0x06, 0x6b, 0x3f, 0x54, 0x6a,
	0xec, 0x54, 0x29, 0xcf, 0xa2, 0x60, 0x14, 0x3c, 0xbe, 0x31, 0x1d, 0x12, 0xef, 0xc4, 0xcf, 0xc6,
	0x85, 0x81, 0x1c, 0xd4, 0xaa, 0xaa, 0xe8, 0xed, 0xf5, 0x91, 0x93, 0x45, 0x86, 0xa7, 0xe8, 0x5e,
	0x5b, 0x26, 0x15, 0xdb, 0xcd, 0x26, 0xfa, 0xdf, 0xb3, 0x62, 0x7d, 0x8a, 0x9b, 0xfd, 0x89, 0x85,
	0xc6, 0xff, 0x42, 0xd4, 0xa7, 0xa0, 0x4b, 0x29, 0x34, 0xe0, 0xa7, 0xa8, 0x53, 0xfb, 0xf4, 0x87,
	0xc7, 0xa4, 0x19, 0xa3, 0xcb, 0x60, 0x5e, 0xfd, 0x52, 0xd7, 0x88, 0x3f, 0xa3, 0x41, 0xe5, 0x30,
	0x3d, 0xb2, 0x18, 0x5d, 0x1b, 0x85, 0x96, 0x4c, 0x5a, 0xe4, 0x76, 0x10, 0x4b, 0x5b, 0x2f, 0x0e,
	0x35, 0xbd, 0x53, 0x34, 0x3f, 0xe0, 0x17, 0xa8, 0xe7, 0x93, 0x8d, 0xc2, 0x5a, 0xf1, 0xe1, 0x39,
	0x45, 0x97, 0xfb, 0xd2, 0xfd, 0xd3, 0x7d, 0x3b, 0xfe, 0x80, 0x42, 0x25, 0x77, 0xd1, 0xf5, 0x9a,
	0xf5, 0x92, 0x9c, 0x72, 0x17, 0xc8, 0x3e, 0x0b, 0x42, 0xe5, 0x8e, 0x56, 0x32, 0xf1, 0xdf, 0x00,
	0x85, 0xb6, 0xc0, 0xf7, 0x51, 0xd7, 0x96, 0xd5, 0x6a, 0x7e, 0x25, 0x36, 0x9e, 0x0e, 0xed, 0xd8,
	0xd2, 0xa6, 0xfe, 0x1a, 0xdd, 0x72, 0xb9, 0x32, 0x5e, 0xa7, 0x1e, 0xfd, 0x4e, 0x9a, 0xe9, 0xf9,
	0xd5, 0x69, 0xa3, 0xb8, 0xc8, 0xdd, 0xe6, 0x6e, 0xee, 0x19, 0xd5, 0xb9, 0xf8, 0xdd, 0x05, 0xeb,
	0xff, 0x93, 0x9c, 0xbc, 0xff, 0xb7, 0x2b, 0x34, 0xe4, 0xb2, 0x65, 0xf7, 0xf0, 0x72, 0xbe, 0x3c,
	0xbf, 0xe2, 0x9b, 0xfa, 0xda, 0xad, 0x2f, 0xed, 0xb3, 0xb3, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x4a, 0xc4, 0x19, 0x95, 0x03, 0x00, 0x00,
}