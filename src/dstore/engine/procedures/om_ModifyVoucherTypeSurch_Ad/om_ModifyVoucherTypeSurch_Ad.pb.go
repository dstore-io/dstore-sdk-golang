// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyVoucherTypeSurch_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyVoucherTypeSurch_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyVoucherTypeSurch_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyVoucherTypeSurch_Ad

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
	VoucherTypeId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=voucher_type_id,json=voucherTypeId" json:"voucher_type_id,omitempty"`
	VoucherTypeIdNull   bool                        `protobuf:"varint,1001,opt,name=voucher_type_id_null,json=voucherTypeIdNull" json:"voucher_type_id_null,omitempty"`
	SurchargeTypeId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull bool                        `protobuf:"varint,1002,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	SurchargeValue      *dstore_values.DecimalValue `protobuf:"bytes,3,opt,name=surcharge_value,json=surchargeValue" json:"surcharge_value,omitempty"`
	SurchargeValueNull  bool                        `protobuf:"varint,1003,opt,name=surcharge_value_null,json=surchargeValueNull" json:"surcharge_value_null,omitempty"`
	DeleteSurcharge     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_surcharge,json=deleteSurcharge" json:"delete_surcharge,omitempty"`
	DeleteSurchargeNull bool                        `protobuf:"varint,1004,opt,name=delete_surcharge_null,json=deleteSurchargeNull" json:"delete_surcharge_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetVoucherTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VoucherTypeId
	}
	return nil
}

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.SurchargeValue
	}
	return nil
}

func (m *Parameters) GetDeleteSurcharge() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteSurcharge
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyVoucherTypeSurch_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyVoucherTypeSurch_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyVoucherTypeSurch_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0xd5, 0x0e, 0x69, 0xab, 0x83, 0x20, 0xad, 0x5b, 0xaa, 0x90, 0x08, 0x84, 0xca, 0x06,
	0xb1, 0x70, 0x80, 0x82, 0x54, 0x89, 0x15, 0x77, 0x05, 0x29, 0x15, 0x72, 0xab, 0x4a, 0xb0, 0x19,
	0xb9, 0xf1, 0x69, 0x3a, 0xd2, 0x8c, 0x1d, 0xd9, 0x33, 0xa9, 0xfa, 0x16, 0x3c, 0x10, 0x2f, 0xc4,
	0xe5, 0x01, 0x58, 0x62, 0x8f, 0x9d, 0x84, 0x71, 0x23, 0x40, 0x6c, 0x92, 0x9c, 0x9c, 0xf3, 0xff,
	0x9f, 0x2f, 0xff, 0x0c, 0x3c, 0x17, 0xa6, 0x54, 0x1a, 0xfb, 0x28, 0xc7, 0x99, 0xc4, 0xfe, 0x44,
	0xab, 0x11, 0x8a, 0x4a, 0xa3, 0xe9, 0xab, 0x22, 0x1d, 0x2a, 0x91, 0x9d, 0x5d, 0x9e, 0xa8, 0x6a,
	0x74, 0x8e, 0xfa, 0xf8, 0x72, 0x82, 0x47, 0x95, 0x1e, 0x9d, 0xa7, 0x2f, 0x04, 0xb5, 0x73, 0xa5,
	0x22, 0x0f, 0xbd, 0x98, 0x7a, 0x31, 0xfd, 0x93, 0xa2, 0xbb, 0x1d, 0x40, 0x53, 0x9e, 0x57, 0x68,
	0xbc, 0x41, 0xf7, 0x76, 0x93, 0x8e, 0x5a, 0x2b, 0x1d, 0x5a, 0xbd, 0x66, 0xab, 0x40, 0x63, 0xf8,
	0x18, 0x43, 0xf3, 0x7e, 0xdc, 0x2c, 0x79, 0x26, 0xcf, 0x94, 0x2e, 0x78, 0x99, 0x29, 0xe9, 0x87,
	0xf6, 0x7e, 0x26, 0x00, 0x1f, 0xb8, 0xe6, 0xb6, 0x8b, 0xda, 0x90, 0x57, 0xd0, 0x9e, 0xfa, 0x75,
	0xa5, 0xa5, 0x5d, 0x58, 0x9a, 0x89, 0xce, 0xca, 0xbd, 0x95, 0x07, 0xd7, 0x9f, 0xf4, 0x68, 0xd8,
	0x46, 0x58, 0x5a, 0x26, 0x4b, 0x1c, 0xa3, 0x3e, 0x71, 0x15, 0xbb, 0x31, 0x5d, 0xec, 0x65, 0x20,
	0xc8, 0x23, 0xd8, 0x89, 0x4c, 0x52, 0x59, 0xe5, 0x79, 0xe7, 0xeb, 0xba, 0xb5, 0xda, 0x60, 0x5b,
	0x8d, 0xe9, 0x43, 0xdb, 0x21, 0xef, 0x60, 0xcb, 0xb8, 0x33, 0xe0, 0x7a, 0x8c, 0x73, 0xf0, 0xea,
	0xdf, 0xc1, 0xed, 0xb9, 0x2a, 0xa0, 0x9f, 0xc2, 0xee, 0x15, 0x23, 0x0f, 0xff, 0xe6, 0xe1, 0xdb,
	0x91, 0xa2, 0xc6, 0xbf, 0x86, 0x85, 0x51, 0x5a, 0x73, 0x3a, 0xc9, 0x52, 0xb8, 0xc0, 0x51, 0x56,
	0xf0, 0xdc, 0xc3, 0x6f, 0xce, 0x35, 0x75, 0x4d, 0x1e, 0xc3, 0x4e, 0xe4, 0xe2, 0xc9, 0xdf, 0x3d,
	0x99, 0x34, 0xc7, 0x6b, 0xf0, 0x5b, 0xd8, 0x14, 0x98, 0xdb, 0xa3, 0x4f, 0xe7, 0xcd, 0xce, 0xb5,
	0xa5, 0xe4, 0x53, 0xa5, 0x72, 0xe4, 0x32, 0x6c, 0xdb, 0x8b, 0x8e, 0x66, 0x1a, 0xb2, 0x0f, 0xb7,
	0x62, 0x1f, 0xcf, 0xfe, 0x11, 0x76, 0x1d, 0x09, 0x1c, 0x7c, 0xef, 0xcb, 0x2a, 0x6c, 0x30, 0x34,
	0x13, 0x25, 0x0d, 0xda, 0x3b, 0x6b, 0xd5, 0xc1, 0x0a, 0xd7, 0xdd, 0xa5, 0xcd, 0xd4, 0xfa, 0xd0,
	0xbd, 0x71, 0x9f, 0xcc, 0x0f, 0x92, 0x8f, 0xb0, 0xe9, 0x22, 0x95, 0xfe, 0x96, 0x29, 0x7b, 0x65,
	0x89, 0x15, 0xd3, 0x48, 0x1c, 0x27, 0x6f, 0x68, 0xeb, 0xc1, 0xa2, 0x66, 0xed, 0xa2, 0xf9, 0x07,
	0x39, 0x80, 0xf5, 0x10, 0x65, 0x7b, 0x0f, 0xce, 0xf1, 0xee, 0x15, 0x47, 0x1f, 0xf4, 0xa1, 0xff,
	0x66, 0xb3, 0x71, 0xf2, 0x1e, 0x12, 0xad, 0x2e, 0xec, 0x19, 0x3a, 0xd5, 0x01, 0xfd, 0xf7, 0x47,
	0x8f, 0xce, 0x4e, 0x82, 0x32, 0x75, 0xc1, 0x9c, 0x49, 0xf7, 0x0e, 0x24, 0xf6, 0x37, 0xd9, 0x85,
	0x35, 0x5b, 0xb9, 0x40, 0x7e, 0x3e, 0xb4, 0x67, 0xd3, 0x62, 0x2d, 0x5b, 0x0e, 0xc4, 0xcb, 0x63,
	0xe8, 0x65, 0x2a, 0x22, 0x2c, 0xde, 0x0c, 0x9f, 0x9e, 0xfd, 0xd7, 0x3b, 0xe3, 0x74, 0xad, 0x7e,
	0x2c, 0xf7, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xde, 0xb7, 0xaf, 0x44, 0x73, 0x04, 0x00, 0x00,
}
