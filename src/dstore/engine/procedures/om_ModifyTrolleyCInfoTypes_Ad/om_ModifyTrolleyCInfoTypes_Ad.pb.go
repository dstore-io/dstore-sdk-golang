// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyTrolleyCInfoTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyTrolleyCInfoTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyTrolleyCInfoTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyTrolleyCInfoTypes_Ad

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
	InformationTypeId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull          bool                        `protobuf:"varint,1001,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	InformationType                *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	InformationTypeNull            bool                        `protobuf:"varint,1002,opt,name=information_type_null,json=informationTypeNull" json:"information_type_null,omitempty"`
	FieldTypeId                    *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	FieldTypeIdNull                bool                        `protobuf:"varint,1003,opt,name=field_type_id_null,json=fieldTypeIdNull" json:"field_type_id_null,omitempty"`
	Delete                         *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
	DeleteNull                     bool                        `protobuf:"varint,1004,opt,name=delete_null,json=deleteNull" json:"delete_null,omitempty"`
	CorrespOrderContInfoTypeId     *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=corresp_order_cont_info_type_id,json=correspOrderContInfoTypeId" json:"corresp_order_cont_info_type_id,omitempty"`
	CorrespOrderContInfoTypeIdNull bool                        `protobuf:"varint,1005,opt,name=corresp_order_cont_info_type_id_null,json=correspOrderContInfoTypeIdNull" json:"corresp_order_cont_info_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Parameters) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func (m *Parameters) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
	}
	return nil
}

func (m *Parameters) GetDelete() *dstore_values.BooleanValue {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Parameters) GetCorrespOrderContInfoTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CorrespOrderContInfoTypeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyTrolleyCInfoTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyTrolleyCInfoTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyTrolleyCInfoTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x6f, 0x6b, 0x13, 0x4f,
	0x10, 0xa6, 0xcd, 0xef, 0xd2, 0x32, 0xe1, 0x47, 0xea, 0x86, 0xca, 0x79, 0xc1, 0x5a, 0xaa, 0x2f,
	0x04, 0xe5, 0x22, 0x06, 0x44, 0x41, 0x10, 0x2d, 0x7d, 0x11, 0x24, 0x55, 0x8e, 0x2a, 0xe8, 0x9b,
	0xf3, 0x9a, 0x9d, 0x84, 0x83, 0xcb, 0x6e, 0xd8, 0xbb, 0x58, 0xf2, 0x05, 0x7c, 0xed, 0x17, 0xf2,
	0x03, 0xf9, 0xef, 0x3b, 0x38, 0x7b, 0xb3, 0x49, 0x93, 0x4b, 0x35, 0xe8, 0x9b, 0xe4, 0xe6, 0x66,
	0x9e, 0xe7, 0x99, 0xe7, 0x76, 0x66, 0xe1, 0xa9, 0xcc, 0x0b, 0x6d, 0xb0, 0x83, 0x6a, 0x94, 0x2a,
	0xec, 0x4c, 0x8c, 0x1e, 0xa0, 0x9c, 0x1a, 0xcc, 0x3b, 0x7a, 0x1c, 0xf7, 0xb5, 0x4c, 0x87, 0xb3,
	0x33, 0xa3, 0xb3, 0x0c, 0x67, 0xc7, 0x3d, 0x35, 0xd4, 0x67, 0xb3, 0x09, 0xe6, 0xf1, 0x73, 0x19,
	0x52, 0x61, 0xa1, 0xc5, 0x3d, 0x46, 0x87, 0x8c, 0x0e, 0xff, 0x08, 0x09, 0x5a, 0x4e, 0xea, 0x63,
	0x92, 0x4d, 0x31, 0x67, 0x86, 0xe0, 0xc6, 0xaa, 0x3e, 0x1a, 0xa3, 0x8d, 0x4b, 0xb5, 0x57, 0x53,
	0x63, 0xcc, 0xf3, 0x64, 0x84, 0x2e, 0x79, 0xbb, 0x9a, 0x2c, 0x92, 0x94, 0xc4, 0xcc, 0x38, 0x29,
	0x52, 0xad, 0xb8, 0xe8, 0xe8, 0x93, 0x07, 0xf0, 0x3a, 0x31, 0x09, 0x65, 0xd1, 0xe4, 0xe2, 0x25,
	0xb4, 0x96, 0x6a, 0xe2, 0x82, 0x1a, 0x8b, 0x53, 0xe9, 0x6f, 0x1d, 0x6e, 0xdd, 0x6d, 0x3c, 0x6c,
	0x87, 0xce, 0x8b, 0x6b, 0x2f, 0x55, 0x05, 0x8e, 0xd0, 0xbc, 0xb5, 0x51, 0x74, 0x6d, 0x09, 0x67,
	0xfd, 0xf4, 0xa4, 0x78, 0x0c, 0xfe, 0x15, 0x64, 0xb1, 0x9a, 0x66, 0x99, 0xff, 0x75, 0x87, 0x28,
	0x77, 0xa3, 0xfd, 0x35, 0xd4, 0x29, 0x65, 0xc5, 0x09, 0xec, 0x55, 0x91, 0xfe, 0x76, 0xd9, 0x43,
	0x50, 0xe9, 0x21, 0x2f, 0x4c, 0xaa, 0x46, 0xdc, 0x42, 0xb3, 0x42, 0x26, 0xba, 0xb0, 0xbf, 0xd6,
	0x40, 0xa9, 0xfe, 0x8d, 0xd5, 0x5b, 0x15, 0x40, 0xa9, 0xfd, 0x0c, 0xfe, 0x1f, 0xa6, 0x98, 0xc9,
	0x85, 0xf9, 0xda, 0x66, 0xf3, 0x8d, 0x12, 0xe1, 0x6c, 0xdf, 0x07, 0xb1, 0x42, 0xc0, 0x92, 0xdf,
	0x59, 0xb2, 0xb9, 0x54, 0x59, 0xca, 0x75, 0xa1, 0x2e, 0x31, 0xa3, 0xaf, 0xef, 0xff, 0x77, 0xa5,
	0xce, 0xb9, 0xd6, 0x19, 0x26, 0x8a, 0x75, 0x5c, 0xa9, 0x38, 0x84, 0x06, 0x3f, 0x31, 0xf7, 0x0f,
	0xe6, 0x06, 0x7e, 0x57, 0xd2, 0x7e, 0x80, 0x5b, 0x03, 0x6d, 0x68, 0x48, 0x27, 0xb1, 0x36, 0x12,
	0x4d, 0x3c, 0xd0, 0xaa, 0x88, 0xad, 0xdf, 0x85, 0x2f, 0x6f, 0xb3, 0xaf, 0xc0, 0x71, 0xbc, 0xb2,
	0x14, 0xc7, 0xc4, 0x30, 0x9f, 0x56, 0xb2, 0xd9, 0x87, 0x3b, 0x1b, 0x14, 0xb8, 0xb9, 0x9f, 0xdc,
	0xdc, 0xc1, 0xef, 0xa9, 0x6c, 0xc3, 0x47, 0x5f, 0xb6, 0x61, 0x37, 0xa2, 0xbc, 0x56, 0x39, 0x8a,
	0x07, 0xe0, 0x95, 0x63, 0xee, 0x06, 0x6f, 0x71, 0xe8, 0x6e, 0x89, 0x78, 0x05, 0x4e, 0xec, 0x6f,
	0xc4, 0x85, 0xe2, 0x1d, 0xec, 0xd9, 0x01, 0x8f, 0x97, 0x4e, 0x94, 0x26, 0xa6, 0x46, 0xe0, 0xb0,
	0x02, 0xae, 0xee, 0x41, 0x9f, 0xe2, 0xde, 0x65, 0x1c, 0x35, 0xc7, 0xab, 0x2f, 0x68, 0x8c, 0x77,
	0xdc, 0x62, 0xd1, 0x28, 0x58, 0xc6, 0x83, 0x35, 0x46, 0x5e, 0xbb, 0x3e, 0xff, 0x47, 0xf3, 0x72,
	0xda, 0xa6, 0x9a, 0xd1, 0x17, 0x74, 0xb0, 0x16, 0xf5, 0x24, 0xfc, 0x8b, 0x9b, 0x20, 0x9c, 0x7f,
	0x8a, 0x30, 0xd2, 0x17, 0x91, 0x65, 0x09, 0x6e, 0x42, 0x8d, 0x9e, 0xc5, 0x75, 0xa8, 0x53, 0x64,
	0xcf, 0xef, 0xf3, 0x29, 0x7d, 0x1c, 0x2f, 0xf2, 0x28, 0xec, 0xc9, 0x17, 0x6f, 0xa0, 0x9d, 0xea,
	0x8a, 0xc4, 0xe5, 0x55, 0xf5, 0xfe, 0xd1, 0xbf, 0x5d, 0x62, 0xe7, 0xf5, 0xf2, 0x9a, 0xe8, 0xfe,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x66, 0x69, 0x5f, 0x05, 0x05, 0x00, 0x00,
}
