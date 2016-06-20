// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifySurchargeTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifySurchargeTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifySurchargeTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifySurchargeTypes_Ad

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
	SurchargeTypeId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull          bool                        `protobuf:"varint,1001,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	SurchargeTypeDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=surcharge_type_description,json=surchargeTypeDescription" json:"surcharge_type_description,omitempty"`
	SurchargeTypeDescriptionNull bool                        `protobuf:"varint,1002,opt,name=surcharge_type_description_null,json=surchargeTypeDescriptionNull" json:"surcharge_type_description_null,omitempty"`
	Relative                     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=relative" json:"relative,omitempty"`
	RelativeNull                 bool                        `protobuf:"varint,1003,opt,name=relative_null,json=relativeNull" json:"relative_null,omitempty"`
	Brutto                       *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=brutto" json:"brutto,omitempty"`
	BruttoNull                   bool                        `protobuf:"varint,1004,opt,name=brutto_null,json=bruttoNull" json:"brutto_null,omitempty"`
	UnitId                       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	UnitIdNull                   bool                        `protobuf:"varint,1005,opt,name=unit_id_null,json=unitIdNull" json:"unit_id_null,omitempty"`
	DeleteSurchargeType          *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=delete_surcharge_type,json=deleteSurchargeType" json:"delete_surcharge_type,omitempty"`
	DeleteSurchargeTypeNull      bool                        `protobuf:"varint,1006,opt,name=delete_surcharge_type_null,json=deleteSurchargeTypeNull" json:"delete_surcharge_type_null,omitempty"`
	CategoryId                   *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	CategoryIdNull               bool                        `protobuf:"varint,1007,opt,name=category_id_null,json=categoryIdNull" json:"category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetSurchargeTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeTypeDescription
	}
	return nil
}

func (m *Parameters) GetRelative() *dstore_values.BooleanValue {
	if m != nil {
		return m.Relative
	}
	return nil
}

func (m *Parameters) GetBrutto() *dstore_values.BooleanValue {
	if m != nil {
		return m.Brutto
	}
	return nil
}

func (m *Parameters) GetUnitId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UnitId
	}
	return nil
}

func (m *Parameters) GetDeleteSurchargeType() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteSurchargeType
	}
	return nil
}

func (m *Parameters) GetCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CategoryId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifySurchargeTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifySurchargeTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifySurchargeTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x1b, 0xe2, 0x54, 0xd3, 0x42, 0xcb, 0x46, 0x14, 0xe3, 0x70, 0x29, 0x85, 0x87, 0xf2,
	0xe2, 0x20, 0x52, 0x71, 0x91, 0xfa, 0x02, 0xe2, 0x96, 0x87, 0x14, 0x64, 0x10, 0x02, 0x5e, 0x2c,
	0x27, 0x9e, 0x06, 0x4b, 0x8e, 0x37, 0xda, 0xb5, 0x5b, 0xe5, 0x2f, 0xf8, 0x18, 0x9e, 0xf8, 0x23,
	0xee, 0xbf, 0xc0, 0xae, 0x67, 0xdd, 0xc4, 0x6e, 0xa2, 0x20, 0x5e, 0x12, 0xef, 0xce, 0x39, 0x67,
	0xce, 0xec, 0xee, 0x0c, 0x3c, 0x0a, 0x65, 0xca, 0x05, 0xb6, 0x31, 0x19, 0x46, 0x09, 0xb6, 0xc7,
	0x82, 0x0f, 0x30, 0xcc, 0x04, 0xca, 0x36, 0x1f, 0xf9, 0x3d, 0x1e, 0x46, 0x47, 0x93, 0x37, 0x99,
	0x18, 0x7c, 0x0a, 0xc4, 0x10, 0xdf, 0x4e, 0xc6, 0x28, 0xfd, 0xc7, 0xa1, 0xab, 0x50, 0x29, 0x67,
	0x7b, 0x44, 0x75, 0x89, 0xea, 0x2e, 0xc6, 0x3b, 0x4d, 0x93, 0xe4, 0x38, 0x88, 0x33, 0x94, 0x44,
	0x77, 0xae, 0x94, 0x33, 0xa3, 0x10, 0x5c, 0x98, 0x50, 0xab, 0x1c, 0x1a, 0xa1, 0x94, 0xc1, 0x10,
	0x4d, 0xf0, 0x56, 0x35, 0x98, 0x06, 0x51, 0x72, 0xc4, 0xc5, 0x28, 0x48, 0x23, 0x9e, 0x10, 0x68,
	0xf7, 0x8b, 0x05, 0xf0, 0x3a, 0x10, 0x81, 0x8a, 0xa2, 0x90, 0xec, 0x05, 0x5c, 0x94, 0x85, 0x2b,
	0x3f, 0x55, 0xb6, 0xfc, 0x28, 0xb4, 0x57, 0x76, 0x56, 0xf6, 0xd6, 0xef, 0xb5, 0x5c, 0x53, 0x86,
	0x31, 0x17, 0x25, 0x29, 0x0e, 0x51, 0xbc, 0xd3, 0x2b, 0x6f, 0x53, 0xce, 0xd6, 0xd2, 0x0d, 0xd9,
	0x3e, 0x6c, 0x9f, 0x11, 0xf2, 0x93, 0x2c, 0x8e, 0xed, 0x6f, 0x0d, 0x25, 0xb7, 0xe6, 0x35, 0x2b,
	0x8c, 0x43, 0x15, 0x63, 0xef, 0xc1, 0xa9, 0xb0, 0x42, 0x94, 0x03, 0x11, 0x8d, 0xb5, 0x63, 0x7b,
	0x35, 0xf7, 0xe1, 0x54, 0x7c, 0xc8, 0x54, 0x44, 0xc9, 0x90, 0x6c, 0xd8, 0x25, 0xd1, 0xa7, 0x53,
	0x2e, 0x7b, 0x0e, 0x37, 0x16, 0x2b, 0x93, 0xb1, 0xef, 0x64, 0xec, 0xea, 0x22, 0x8d, 0xdc, 0xe1,
	0x03, 0x58, 0x13, 0x18, 0xab, 0x13, 0x3c, 0x46, 0xbb, 0x36, 0xf7, 0x5c, 0xfa, 0x9c, 0xc7, 0x18,
	0x24, 0x64, 0xe8, 0x14, 0xcc, 0x6e, 0xc3, 0xf9, 0xe2, 0x9b, 0xd2, 0xfd, 0xa0, 0x74, 0x1b, 0xc5,
	0x6e, 0x2e, 0xdf, 0x01, 0xab, 0x2f, 0xb2, 0x34, 0xe5, 0xf6, 0xb9, 0xe5, 0xe2, 0x06, 0xca, 0x76,
	0x60, 0x9d, 0xbe, 0x48, 0xf8, 0x27, 0x09, 0x03, 0xed, 0xe5, 0xb2, 0xfb, 0xd0, 0xc8, 0x92, 0x28,
	0xd5, 0x97, 0x59, 0x5f, 0x7e, 0x99, 0x96, 0xc6, 0xaa, 0x3b, 0xbc, 0x09, 0x1b, 0x86, 0x45, 0xc2,
	0xbf, 0x8c, 0x30, 0x85, 0x73, 0xe1, 0x57, 0x70, 0x29, 0xc4, 0x58, 0xbd, 0x1d, 0xbf, 0x7c, 0xba,
	0xb6, 0xb5, 0xdc, 0x7e, 0x93, 0x98, 0xa5, 0x2e, 0x60, 0x07, 0xe0, 0xcc, 0x15, 0x24, 0x07, 0xbf,
	0xc9, 0xc1, 0xe5, 0x39, 0xcc, 0xdc, 0xce, 0x01, 0xac, 0x0f, 0x02, 0x55, 0x09, 0x17, 0x13, 0x5d,
	0x6b, 0x63, 0x79, 0xad, 0x50, 0xe0, 0x55, 0xbd, 0x77, 0x60, 0x6b, 0x86, 0x4d, 0x19, 0xff, 0x50,
	0xc6, 0x0b, 0x53, 0x98, 0x4e, 0xb4, 0xfb, 0x75, 0x15, 0xd6, 0x3c, 0x94, 0x63, 0x9e, 0x48, 0x64,
	0x77, 0xa1, 0x9e, 0x37, 0xa5, 0x69, 0x94, 0xd3, 0x07, 0x6a, 0xfa, 0x9d, 0x1a, 0xf6, 0x99, 0xfe,
	0xf5, 0x08, 0xc8, 0x3e, 0xc0, 0x96, 0x6e, 0x47, 0x7f, 0xa6, 0x1f, 0xd5, 0xeb, 0xae, 0x29, 0xb2,
	0x5b, 0x21, 0x57, 0xbb, 0xb6, 0xa7, 0xd6, 0xdd, 0xe9, 0xda, 0xdb, 0x1c, 0x95, 0x37, 0xd8, 0x43,
	0x68, 0x98, 0x31, 0xa0, 0xde, 0xa7, 0x56, 0xbc, 0x7e, 0x46, 0x91, 0x86, 0x44, 0x8f, 0xfe, 0xbd,
	0x02, 0xce, 0x5e, 0x42, 0x4d, 0xf0, 0x13, 0xf5, 0xf0, 0x34, 0xeb, 0xbe, 0xfb, 0xaf, 0x43, 0xcb,
	0x2d, 0xce, 0xc1, 0xf5, 0xf8, 0x89, 0xa7, 0x25, 0x9c, 0x6b, 0x50, 0x53, 0xdf, 0x6c, 0x1b, 0x2c,
	0xb5, 0xd2, 0x17, 0xf1, 0xf9, 0x50, 0x9d, 0x4c, 0xdd, 0xab, 0xab, 0x65, 0x37, 0x7c, 0xe2, 0x41,
	0x2b, 0xe2, 0x15, 0xfd, 0xe9, 0x3c, 0xfd, 0xd8, 0xf9, 0x8f, 0x49, 0xdb, 0xb7, 0xf2, 0x71, 0xd6,
	0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x35, 0x31, 0x75, 0xa7, 0x05, 0x00, 0x00,
}