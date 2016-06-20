// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetUnits_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetUnits_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetUnits_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetUnits_Ad

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
	UnitId             *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	UnitIdNull         bool                        `protobuf:"varint,1001,opt,name=unit_id_null,json=unitIdNull" json:"unit_id_null,omitempty"`
	UnitCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=unit_category_id,json=unitCategoryId" json:"unit_category_id,omitempty"`
	UnitCategoryIdNull bool                        `protobuf:"varint,1002,opt,name=unit_category_id_null,json=unitCategoryIdNull" json:"unit_category_id_null,omitempty"`
	Active             *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=active" json:"active,omitempty"`
	ActiveNull         bool                        `protobuf:"varint,1003,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitId
	}
	return nil
}

func (m *Parameters) GetUnitCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitCategoryId
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
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
	RowId                   int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Active                  *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=active" json:"active,omitempty"`
	UnitId                  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	UnitSymbol              *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=unit_symbol,json=unitSymbol" json:"unit_symbol,omitempty"`
	UnitDescription         *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=unit_description,json=unitDescription" json:"unit_description,omitempty"`
	UnitCategoryId          *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=unit_category_id,json=unitCategoryId" json:"unit_category_id,omitempty"`
	TriangleConvert         *dstore_values.BooleanValue `protobuf:"bytes,10006,opt,name=triangle_convert,json=triangleConvert" json:"triangle_convert,omitempty"`
	UnitCategoryDescription *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=unit_category_description,json=unitCategoryDescription" json:"unit_category_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Response_Row) GetUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitId
	}
	return nil
}

func (m *Response_Row) GetUnitSymbol() *dstore_values.StringValue {
	if m != nil {
		return m.UnitSymbol
	}
	return nil
}

func (m *Response_Row) GetUnitDescription() *dstore_values.StringValue {
	if m != nil {
		return m.UnitDescription
	}
	return nil
}

func (m *Response_Row) GetUnitCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitCategoryId
	}
	return nil
}

func (m *Response_Row) GetTriangleConvert() *dstore_values.BooleanValue {
	if m != nil {
		return m.TriangleConvert
	}
	return nil
}

func (m *Response_Row) GetUnitCategoryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.UnitCategoryDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetUnits_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetUnits_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetUnits_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x1b, 0x92, 0x54, 0x37, 0x88, 0x46, 0x46, 0xc0, 0x34, 0x41, 0xa8, 0x94, 0x05, 0x08,
	0x89, 0x09, 0x6a, 0x8b, 0xc4, 0xa2, 0x1b, 0x28, 0x6d, 0xd5, 0x45, 0x47, 0x68, 0x10, 0x20, 0xd8,
	0x8c, 0x9c, 0xcc, 0x25, 0xb2, 0x34, 0xb1, 0x23, 0xdb, 0x49, 0xd5, 0x7f, 0x60, 0xc1, 0x9b, 0x4f,
	0xe1, 0x9b, 0x80, 0x3f, 0x60, 0x85, 0x3d, 0x76, 0x9a, 0x99, 0x01, 0x91, 0xb0, 0x49, 0xe6, 0xfa,
	0x9e, 0x73, 0x1f, 0x3e, 0x47, 0x86, 0x7b, 0xa9, 0xd2, 0x42, 0x62, 0x0f, 0xf9, 0x90, 0x71, 0xec,
	0x8d, 0xa5, 0x18, 0x60, 0x3a, 0x91, 0xa8, 0x7a, 0x23, 0x96, 0x1c, 0xa1, 0x7e, 0xce, 0x99, 0x56,
	0xc9, 0xa3, 0x34, 0x34, 0x19, 0x2d, 0xc8, 0x75, 0x07, 0x0f, 0x1d, 0x3c, 0x2c, 0x63, 0x3a, 0x97,
	0x7d, 0xb1, 0x29, 0xcd, 0x26, 0xa8, 0x1c, 0xa5, 0xb3, 0x51, 0xee, 0x80, 0x52, 0x0a, 0xe9, 0x53,
	0xdd, 0x72, 0x6a, 0x84, 0x4a, 0xd1, 0x21, 0xfa, 0xe4, 0xad, 0x6a, 0x52, 0x53, 0xc6, 0xdf, 0x08,
	0x39, 0xa2, 0x9a, 0x09, 0xee, 0x40, 0x5b, 0xdf, 0x56, 0x01, 0x9e, 0x52, 0x49, 0x4d, 0x16, 0xa5,
	0x22, 0xbb, 0xd0, 0x9c, 0x98, 0x61, 0x12, 0x96, 0x06, 0x2b, 0x9b, 0x2b, 0x77, 0x5a, 0xdb, 0xdd,
	0xd0, 0x0f, 0xec, 0x47, 0x62, 0x5c, 0xe3, 0x10, 0xe5, 0x0b, 0x1b, 0xc5, 0x0d, 0x8b, 0x3d, 0x4e,
	0xc9, 0x4d, 0xb8, 0xe8, 0x59, 0x09, 0x9f, 0x64, 0x59, 0xf0, 0xbd, 0x69, 0xb8, 0x6b, 0x31, 0xb8,
	0x74, 0x64, 0x8e, 0xc8, 0x01, 0xb4, 0x73, 0xc8, 0x80, 0x1a, 0xbe, 0x90, 0x67, 0xb6, 0xc3, 0xea,
	0xe2, 0x0e, 0x97, 0x2c, 0x69, 0xdf, 0x73, 0x4c, 0xa7, 0x6d, 0xb8, 0x52, 0x2d, 0xe3, 0x5a, 0xfe,
	0x70, 0x2d, 0x49, 0x19, 0x9f, 0xb7, 0xde, 0x81, 0x06, 0x1d, 0x68, 0x36, 0xc5, 0xa0, 0xb6, 0xc4,
	0x4a, 0x0e, 0x4a, 0x36, 0xa1, 0xe5, 0xbe, 0x5c, 0xf9, 0x9f, 0x7e, 0x23, 0x77, 0x66, 0xcb, 0x6e,
	0xbd, 0x6d, 0xc0, 0x5a, 0x8c, 0x6a, 0x2c, 0xb8, 0x42, 0x72, 0x1f, 0xea, 0xb9, 0x2e, 0xfe, 0xd6,
	0x3a, 0x61, 0x59, 0x66, 0xa7, 0xd9, 0x81, 0xfd, 0x8d, 0x1d, 0x90, 0xbc, 0x82, 0xb6, 0x55, 0x24,
	0x29, 0x48, 0x62, 0x2e, 0xa4, 0x66, 0xc8, 0x61, 0x85, 0x5c, 0x15, 0xee, 0xc4, 0xc4, 0xc7, 0xf3,
	0x38, 0x5e, 0x1f, 0x95, 0x0f, 0xc8, 0x43, 0x68, 0x7a, 0x27, 0x98, 0x8d, 0x6d, 0xc5, 0x1b, 0x7f,
	0x54, 0x74, 0x3e, 0x39, 0x71, 0xff, 0xf1, 0x0c, 0x4e, 0xf6, 0xa0, 0x26, 0xc5, 0x69, 0x70, 0x21,
	0x67, 0xdd, 0x0d, 0xff, 0xe5, 0xd5, 0x70, 0xb6, 0x7b, 0x18, 0x8b, 0xd3, 0xd8, 0xd2, 0x3a, 0xbf,
	0x6a, 0x50, 0x33, 0x01, 0xb9, 0x0a, 0x0d, 0x13, 0x5a, 0x85, 0xdf, 0x45, 0xe6, 0x3a, 0xea, 0x71,
	0xdd, 0x84, 0x46, 0xbc, 0xdd, 0x73, 0x21, 0xde, 0x47, 0xcb, 0x2b, 0xf1, 0x60, 0x6e, 0xc9, 0x0f,
	0xd1, 0xf2, 0x9e, 0xdc, 0x83, 0x56, 0x4e, 0x53, 0x67, 0xa3, 0xbe, 0xc8, 0x82, 0x8f, 0x51, 0x59,
	0x18, 0x4f, 0x55, 0x5a, 0x32, 0x3e, 0x74, 0xcc, 0xdc, 0xae, 0xcf, 0x72, 0x38, 0x39, 0xf4, 0x76,
	0x4d, 0x51, 0x0d, 0x24, 0x1b, 0xe7, 0xea, 0x7c, 0x5a, 0x5c, 0x62, 0xdd, 0x92, 0x9e, 0xcc, 0x39,
	0xe7, 0x75, 0x8a, 0xb6, 0xff, 0x1c, 0xfd, 0xbf, 0xef, 0x8f, 0xa0, 0x6d, 0xda, 0x50, 0x3e, 0xcc,
	0x30, 0x19, 0x08, 0x3e, 0x45, 0xa9, 0x83, 0x2f, 0x7f, 0xaf, 0xd3, 0x17, 0x22, 0x43, 0xca, 0xfd,
	0x40, 0x33, 0xd6, 0xbe, 0x23, 0x91, 0x97, 0xb0, 0x51, 0x1e, 0xa8, 0xb8, 0xe1, 0xd7, 0xc5, 0x1b,
	0x5e, 0x2b, 0x0e, 0x56, 0xd8, 0xf4, 0xf1, 0x21, 0x74, 0x99, 0xa8, 0x38, 0x66, 0xfe, 0x18, 0xbe,
	0xbe, 0xbd, 0xe4, 0x33, 0xd9, 0x6f, 0xe4, 0xef, 0xd2, 0xce, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd1, 0xcc, 0x69, 0xd4, 0x58, 0x05, 0x00, 0x00,
}
