// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetCorrespondingValues_Ad.proto
// DO NOT EDIT!

/*
Package im_GetCorrespondingValues_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetCorrespondingValues_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetCorrespondingValues_Ad

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
	GetCorrValsForValueId            *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=get_corr_vals_for_value_id,json=getCorrValsForValueId" json:"get_corr_vals_for_value_id,omitempty"`
	GetCorrValsForValueIdNull        bool                        `protobuf:"varint,1001,opt,name=get_corr_vals_for_value_id_null,json=getCorrValsForValueIdNull" json:"get_corr_vals_for_value_id_null,omitempty"`
	GetCorrValsForValsOfCharacId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=get_corr_vals_for_vals_of_charac_id,json=getCorrValsForValsOfCharacId" json:"get_corr_vals_for_vals_of_charac_id,omitempty"`
	GetCorrValsForValsOfCharacIdNull bool                        `protobuf:"varint,1002,opt,name=get_corr_vals_for_vals_of_charac_id_null,json=getCorrValsForValsOfCharacIdNull" json:"get_corr_vals_for_vals_of_charac_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetGetCorrValsForValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetCorrValsForValueId
	}
	return nil
}

func (m *Parameters) GetGetCorrValsForValsOfCharacId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetCorrValsForValsOfCharacId
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
	RowId                          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CorrespondingValue             *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=corresponding_value,json=correspondingValue" json:"corresponding_value,omitempty"`
	CorrespondingCharacId          *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=corresponding_charac_id,json=correspondingCharacId" json:"corresponding_charac_id,omitempty"`
	CorrespondingValueId           *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=corresponding_value_id,json=correspondingValueId" json:"corresponding_value_id,omitempty"`
	Value                          *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=value" json:"value,omitempty"`
	CorrespondingCharacDescription *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=corresponding_charac_description,json=correspondingCharacDescription" json:"corresponding_charac_description,omitempty"`
	ValueId                        *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCorrespondingValue() *dstore_values.StringValue {
	if m != nil {
		return m.CorrespondingValue
	}
	return nil
}

func (m *Response_Row) GetCorrespondingCharacId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CorrespondingCharacId
	}
	return nil
}

func (m *Response_Row) GetCorrespondingValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CorrespondingValueId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetCorrespondingCharacDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CorrespondingCharacDescription
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetCorrespondingValues_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetCorrespondingValues_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetCorrespondingValues_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x6f, 0xd3, 0x4c,
	0x10, 0x55, 0xeb, 0x2f, 0x17, 0xcd, 0xf7, 0x00, 0xda, 0x42, 0x71, 0x1d, 0x54, 0xa2, 0xf6, 0xa5,
	0xe2, 0x61, 0xc3, 0x45, 0x88, 0x4a, 0x3c, 0x41, 0xb8, 0xa8, 0x88, 0x06, 0xb4, 0x54, 0x95, 0xe0,
	0x65, 0x65, 0xec, 0x8d, 0x59, 0x29, 0xf1, 0x46, 0xbb, 0x4e, 0xfb, 0x17, 0x78, 0xe4, 0xfe, 0x13,
	0x11, 0xf0, 0x2b, 0x18, 0xef, 0x3a, 0x4d, 0x9d, 0x98, 0xc6, 0xe2, 0xc5, 0xd6, 0xec, 0xcc, 0x9c,
	0x39, 0x67, 0xe7, 0x68, 0xe1, 0x41, 0x6c, 0x32, 0xa5, 0x45, 0x4f, 0xa4, 0x89, 0x4c, 0x45, 0x6f,
	0xa2, 0x55, 0x24, 0xe2, 0xa9, 0x16, 0xa6, 0x27, 0xc7, 0xfc, 0x99, 0xc8, 0xfa, 0x4a, 0x63, 0x30,
	0x51, 0x69, 0x2c, 0xd3, 0xe4, 0x38, 0x1c, 0x4d, 0x85, 0xe1, 0x0f, 0x63, 0x8a, 0x75, 0x99, 0x22,
	0x37, 0x5d, 0x33, 0x75, 0xcd, 0xf4, 0xa2, 0x8e, 0x60, 0xa3, 0x18, 0x74, 0x62, 0x4f, 0x1c, 0x40,
	0xb0, 0x55, 0x9e, 0x2e, 0xb4, 0x56, 0xba, 0x48, 0x75, 0xca, 0xa9, 0xb1, 0x30, 0x26, 0x4c, 0x44,
	0x91, 0xdc, 0x5d, 0x4c, 0x66, 0xa1, 0x4c, 0x87, 0x4a, 0x8f, 0xc3, 0x4c, 0xaa, 0xd4, 0x15, 0xed,
	0xfc, 0x58, 0x07, 0x78, 0x15, 0xea, 0x10, 0xb3, 0x42, 0x1b, 0x72, 0x0c, 0x41, 0x22, 0x32, 0x1e,
	0x21, 0x3d, 0x8e, 0x24, 0x0c, 0xc7, 0x7a, 0x6e, 0xd9, 0x70, 0x19, 0xfb, 0x6b, 0xdd, 0xb5, 0xbd,
	0xff, 0xef, 0x74, 0x68, 0xa1, 0xa8, 0x60, 0x29, 0xd3, 0x4c, 0x24, 0x42, 0x5b, 0x15, 0xec, 0x6a,
	0xe2, 0xc4, 0x61, 0x64, 0x9e, 0x2a, 0x77, 0x78, 0x10, 0x93, 0x3e, 0xdc, 0xf8, 0x3b, 0x2e, 0x4f,
	0xa7, 0xa3, 0x91, 0xff, 0xab, 0x85, 0xe8, 0x6d, 0xb6, 0x55, 0x09, 0x30, 0xc0, 0x0a, 0x32, 0x84,
	0xdd, 0x4a, 0x10, 0xc3, 0xd5, 0x90, 0x47, 0xef, 0x51, 0x45, 0x94, 0xb3, 0x5c, 0x5f, 0xcd, 0xf2,
	0xfa, 0xd2, 0x10, 0xf3, 0x72, 0xd8, 0xb7, 0x08, 0x48, 0xf6, 0x35, 0xec, 0xd5, 0x98, 0xe3, 0x58,
	0xff, 0x76, 0xac, 0xbb, 0x17, 0x01, 0xe6, 0xe4, 0x77, 0x3e, 0x34, 0xa1, 0xcd, 0xec, 0xca, 0x8d,
	0x20, 0xb7, 0xa0, 0x61, 0xd7, 0x58, 0xdc, 0x68, 0x40, 0xcb, 0x1e, 0x71, 0x2b, 0x7e, 0x92, 0x7f,
	0x99, 0x2b, 0x24, 0x6f, 0xe0, 0x72, 0xbe, 0x40, 0x7e, 0x6e, 0x83, 0x28, 0xd4, 0xc3, 0x66, 0xba,
	0xd0, 0xbc, 0xb8, 0xe7, 0x43, 0x8c, 0x0f, 0xe6, 0x31, 0xbb, 0x34, 0x2e, 0x1f, 0x90, 0x7d, 0x68,
	0x15, 0xc6, 0xf1, 0x3d, 0x8b, 0xb8, 0xbd, 0x84, 0xe8, 0x6c, 0x75, 0xe8, 0xfe, 0x6c, 0x56, 0x4e,
	0x9e, 0x83, 0xa7, 0xd5, 0xa9, 0xff, 0x9f, 0xed, 0xda, 0xa7, 0xf5, 0x8d, 0x4e, 0x67, 0x37, 0x41,
	0x99, 0x3a, 0x65, 0x39, 0x48, 0xf0, 0xd3, 0x03, 0x0f, 0x03, 0xb2, 0x09, 0x4d, 0x0c, 0xf3, 0x3d,
	0x7e, 0x1c, 0xe0, 0xe5, 0x34, 0x58, 0x03, 0x43, 0x5c, 0xca, 0x0b, 0xd8, 0x88, 0xce, 0x63, 0x39,
	0xf7, 0xf8, 0x9f, 0x06, 0xe5, 0x1b, 0x2c, 0xb6, 0x6d, 0x32, 0x3d, 0x9b, 0xc7, 0x48, 0xb4, 0xc4,
	0x81, 0x1c, 0xc1, 0xb5, 0x32, 0xda, 0xdc, 0x3e, 0x9f, 0x07, 0x35, 0x5c, 0x5e, 0x6a, 0x3e, 0x33,
	0x0e, 0x83, 0xcd, 0x0a, 0x8e, 0x39, 0xe8, 0x97, 0x1a, 0xa0, 0x57, 0x96, 0x79, 0x22, 0xe6, 0x6d,
	0x68, 0x38, 0xa5, 0x5f, 0x57, 0x2b, 0x75, 0x95, 0x44, 0x40, 0xb7, 0x52, 0x5c, 0x2c, 0x4c, 0xa4,
	0xe5, 0xc4, 0x7a, 0xe7, 0xdb, 0x6a, 0xb4, 0xed, 0x0a, 0x91, 0x8f, 0xe7, 0x10, 0xe4, 0x3e, 0xb4,
	0xcf, 0xf4, 0x7d, 0xaf, 0xa1, 0xaf, 0x75, 0xe2, 0x24, 0x3d, 0x3a, 0x82, 0x8e, 0x54, 0x0b, 0x6e,
	0x99, 0xbf, 0xa9, 0x6f, 0xef, 0xfd, 0xd3, 0x6b, 0xfb, 0xae, 0x69, 0x1f, 0xb4, 0xbb, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x06, 0x71, 0x28, 0xad, 0x05, 0x00, 0x00,
}
