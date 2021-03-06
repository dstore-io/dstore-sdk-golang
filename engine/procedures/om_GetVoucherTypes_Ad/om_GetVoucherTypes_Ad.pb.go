// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetVoucherTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetVoucherTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetVoucherTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetVoucherTypes_Ad

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
	VoucherTypeId              *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=voucher_type_id,json=voucherTypeId" json:"voucher_type_id,omitempty"`
	VoucherTypeIdNull          bool                        `protobuf:"varint,1001,opt,name=voucher_type_id_null,json=voucherTypeIdNull" json:"voucher_type_id_null,omitempty"`
	VCodeOriginTypeId          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=v_code_origin_type_id,json=vCodeOriginTypeId" json:"v_code_origin_type_id,omitempty"`
	VCodeOriginTypeIdNull      bool                        `protobuf:"varint,1002,opt,name=v_code_origin_type_id_null,json=vCodeOriginTypeIdNull" json:"v_code_origin_type_id_null,omitempty"`
	CodeStatus                 *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=code_status,json=codeStatus" json:"code_status,omitempty"`
	CodeStatusNull             bool                        `protobuf:"varint,1003,opt,name=code_status_null,json=codeStatusNull" json:"code_status_null,omitempty"`
	SortByCodeCreationDate     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=sort_by_code_creation_date,json=sortByCodeCreationDate" json:"sort_by_code_creation_date,omitempty"`
	SortByCodeCreationDateNull bool                        `protobuf:"varint,1004,opt,name=sort_by_code_creation_date_null,json=sortByCodeCreationDateNull" json:"sort_by_code_creation_date_null,omitempty"`
	OutputIntoOneId            *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull        bool                        `protobuf:"varint,1005,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
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

func (m *Parameters) GetVCodeOriginTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VCodeOriginTypeId
	}
	return nil
}

func (m *Parameters) GetCodeStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.CodeStatus
	}
	return nil
}

func (m *Parameters) GetSortByCodeCreationDate() *dstore_values.BooleanValue {
	if m != nil {
		return m.SortByCodeCreationDate
	}
	return nil
}

func (m *Parameters) GetOutputIntoOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputIntoOneId
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
	RowId                  int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	VCodeOriginType        *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=v_code_origin_type,json=vCodeOriginType" json:"v_code_origin_type,omitempty"`
	DefaultValidUntil      *dstore_values.TimestampValue `protobuf:"bytes,10002,opt,name=default_valid_until,json=defaultValidUntil" json:"default_valid_until,omitempty"`
	GenerationPattern      *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=generation_pattern,json=generationPattern" json:"generation_pattern,omitempty"`
	ValidForXDays          *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=valid_for_x_days,json=validForXDays" json:"valid_for_x_days,omitempty"`
	VoucherTypeId          *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=voucher_type_id,json=voucherTypeId" json:"voucher_type_id,omitempty"`
	XTimesUsablePerPerson  *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=x_times_usable_per_person,json=xTimesUsablePerPerson" json:"x_times_usable_per_person,omitempty"`
	VoucherTypeDescription *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=voucher_type_description,json=voucherTypeDescription" json:"voucher_type_description,omitempty"`
	XTimesUsable           *dstore_values.IntegerValue   `protobuf:"bytes,10008,opt,name=x_times_usable,json=xTimesUsable" json:"x_times_usable,omitempty"`
	VCodeOriginTypeId      *dstore_values.IntegerValue   `protobuf:"bytes,10009,opt,name=v_code_origin_type_id,json=vCodeOriginTypeId" json:"v_code_origin_type_id,omitempty"`
	CodeStatus             *dstore_values.IntegerValue   `protobuf:"bytes,10010,opt,name=code_status,json=codeStatus" json:"code_status,omitempty"`
	BenefitTypeDescription *dstore_values.StringValue    `protobuf:"bytes,10011,opt,name=benefit_type_description,json=benefitTypeDescription" json:"benefit_type_description,omitempty"`
	BenefitTypeId          *dstore_values.IntegerValue   `protobuf:"bytes,10012,opt,name=benefit_type_id,json=benefitTypeId" json:"benefit_type_id,omitempty"`
	LastCodeCreatedAt      *dstore_values.TimestampValue `protobuf:"bytes,20010,opt,name=last_code_created_at,json=lastCodeCreatedAt" json:"last_code_created_at,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetVCodeOriginType() *dstore_values.StringValue {
	if m != nil {
		return m.VCodeOriginType
	}
	return nil
}

func (m *Response_Row) GetDefaultValidUntil() *dstore_values.TimestampValue {
	if m != nil {
		return m.DefaultValidUntil
	}
	return nil
}

func (m *Response_Row) GetGenerationPattern() *dstore_values.StringValue {
	if m != nil {
		return m.GenerationPattern
	}
	return nil
}

func (m *Response_Row) GetValidForXDays() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValidForXDays
	}
	return nil
}

func (m *Response_Row) GetVoucherTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VoucherTypeId
	}
	return nil
}

func (m *Response_Row) GetXTimesUsablePerPerson() *dstore_values.IntegerValue {
	if m != nil {
		return m.XTimesUsablePerPerson
	}
	return nil
}

func (m *Response_Row) GetVoucherTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.VoucherTypeDescription
	}
	return nil
}

func (m *Response_Row) GetXTimesUsable() *dstore_values.IntegerValue {
	if m != nil {
		return m.XTimesUsable
	}
	return nil
}

func (m *Response_Row) GetVCodeOriginTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.VCodeOriginTypeId
	}
	return nil
}

func (m *Response_Row) GetCodeStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.CodeStatus
	}
	return nil
}

func (m *Response_Row) GetBenefitTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.BenefitTypeDescription
	}
	return nil
}

func (m *Response_Row) GetBenefitTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BenefitTypeId
	}
	return nil
}

func (m *Response_Row) GetLastCodeCreatedAt() *dstore_values.TimestampValue {
	if m != nil {
		return m.LastCodeCreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetVoucherTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetVoucherTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetVoucherTypes_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetVoucherTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x96, 0xd9, 0x4e, 0xdb, 0x4c,
	0x14, 0xc7, 0xc5, 0x97, 0x2f, 0x80, 0x86, 0xef, 0x63, 0x19, 0x96, 0x9a, 0xa0, 0x2e, 0x82, 0x9b,
	0xf6, 0xc6, 0x74, 0x55, 0xab, 0xaa, 0xbd, 0x00, 0x42, 0x5b, 0x2a, 0x61, 0x22, 0x17, 0xd2, 0xe5,
	0x66, 0xe4, 0xc4, 0x43, 0x6a, 0xd5, 0xf1, 0x58, 0x33, 0xe3, 0x40, 0x2e, 0xfb, 0x06, 0xdd, 0x69,
	0x29, 0x4f, 0xd0, 0x37, 0xeb, 0xf6, 0x0e, 0x3d, 0xe3, 0x31, 0x24, 0x76, 0x82, 0x62, 0x7a, 0x01,
	0x68, 0x98, 0xf3, 0x3b, 0xe7, 0x3f, 0x67, 0x93, 0xd1, 0x4d, 0x57, 0x48, 0xc6, 0xe9, 0x32, 0x0d,
	0x1a, 0x5e, 0x40, 0x97, 0x43, 0xce, 0xea, 0xd4, 0x8d, 0x38, 0x15, 0xcb, 0xac, 0x49, 0x1e, 0x52,
	0x59, 0x65, 0x51, 0xfd, 0x25, 0xe5, 0xdb, 0xed, 0x90, 0x0a, 0xb2, 0xe2, 0x9a, 0x60, 0x20, 0x19,
	0x5e, 0xd2, 0x94, 0xa9, 0x29, 0xb3, 0xaf, 0x69, 0x69, 0x3a, 0x71, 0xdd, 0x72, 0xfc, 0x88, 0x0a,
	0x4d, 0x96, 0xe6, 0xd3, 0xf1, 0x28, 0xe7, 0x8c, 0x27, 0x57, 0x0b, 0xe9, 0xab, 0x26, 0x15, 0xc2,
	0x69, 0xd0, 0xe4, 0x72, 0x29, 0x7b, 0x29, 0x1d, 0x2f, 0xd8, 0x65, 0xbc, 0xe9, 0x48, 0x8f, 0x05,
	0xda, 0x68, 0xf1, 0xa0, 0x88, 0x50, 0xc5, 0xe1, 0x0e, 0xdc, 0x52, 0x2e, 0xf0, 0x1a, 0x9a, 0x68,
	0x69, 0x4d, 0x44, 0x82, 0x28, 0xe2, 0xb9, 0xc6, 0xd0, 0xa5, 0xa1, 0xcb, 0x63, 0xd7, 0x17, 0xcc,
	0x44, 0x7f, 0x22, 0xcd, 0x0b, 0x24, 0x6d, 0x50, 0x5e, 0x55, 0x27, 0xfb, 0xff, 0x56, 0xe7, 0x1d,
	0x1b, 0x2e, 0xbe, 0x8a, 0x66, 0x32, 0x4e, 0x48, 0x10, 0xf9, 0xbe, 0xf1, 0x7d, 0x04, 0x5c, 0x8d,
	0xda, 0x53, 0x29, 0x6b, 0x0b, 0x6e, 0xf0, 0x26, 0x9a, 0x6d, 0x91, 0x3a, 0x73, 0x29, 0x61, 0xdc,
	0x03, 0xb5, 0x27, 0xc1, 0xff, 0x19, 0x1c, 0x7c, 0xaa, 0xb5, 0x06, 0xe0, 0x56, 0xcc, 0x25, 0x02,
	0xee, 0xa2, 0x52, 0x5f, 0x77, 0x5a, 0xc6, 0x0f, 0x2d, 0x63, 0xb6, 0x87, 0x8b, 0xa5, 0xdc, 0x43,
	0x63, 0x31, 0x29, 0xa4, 0x23, 0x23, 0x61, 0x14, 0x06, 0x0b, 0x40, 0xca, 0xfe, 0x49, 0x6c, 0x8e,
	0xaf, 0xa0, 0xc9, 0x2e, 0x5a, 0xc7, 0xfb, 0xa9, 0xe3, 0x8d, 0x77, 0xcc, 0xe2, 0x40, 0x4f, 0x51,
	0x49, 0x30, 0x2e, 0x49, 0xad, 0xad, 0xa5, 0xd6, 0x39, 0x8d, 0x0b, 0x43, 0x5c, 0x47, 0x52, 0xe3,
	0xdf, 0xbe, 0x71, 0x6b, 0x8c, 0xf9, 0xd4, 0x09, 0x74, 0xdc, 0x39, 0x85, 0xaf, 0xb6, 0xd5, 0x2b,
	0xd6, 0x12, 0xb6, 0x0c, 0x28, 0x2e, 0xa3, 0x8b, 0xa7, 0x3b, 0xd6, 0x92, 0x7e, 0x69, 0x49, 0xa5,
	0xfe, 0x1e, 0x62, 0x79, 0x8f, 0x10, 0x66, 0x91, 0x0c, 0x23, 0x49, 0xe0, 0xb1, 0x8c, 0xb0, 0x20,
	0xae, 0x47, 0x71, 0x70, 0x3a, 0x26, 0x34, 0xb6, 0x01, 0xd4, 0x56, 0xa0, 0xaa, 0x71, 0x0b, 0x9d,
	0xeb, 0xf5, 0xa4, 0x75, 0xfc, 0xd6, 0x3a, 0xa6, 0x33, 0x88, 0x12, 0xb0, 0x78, 0x88, 0xd0, 0xa8,
	0x4d, 0x45, 0xc8, 0x02, 0x41, 0xa1, 0xa5, 0x8a, 0x71, 0xdf, 0x27, 0xdd, 0x58, 0x32, 0xd3, 0xd3,
	0xa4, 0x67, 0x62, 0x5d, 0xfd, 0xb6, 0xb5, 0x21, 0x7e, 0x8e, 0x26, 0x55, 0xc7, 0x93, 0xae, 0x96,
	0x87, 0x6e, 0x2a, 0x00, 0x6c, 0x66, 0xe0, 0xec, 0x60, 0x6c, 0xc2, 0x79, 0xa3, 0x73, 0xb6, 0x27,
	0x9a, 0xe9, 0x7f, 0xe0, 0x3b, 0x68, 0x24, 0x99, 0x34, 0x68, 0x0f, 0xe5, 0xf1, 0x42, 0x8f, 0x47,
	0x3d, 0x87, 0x9b, 0xfa, 0xaf, 0x7d, 0x6c, 0x0e, 0xe3, 0x55, 0xe0, 0x6c, 0x0f, 0x8a, 0xab, 0xa8,
	0x6b, 0x66, 0x8e, 0x95, 0x60, 0x1e, 0xa7, 0xc0, 0xb4, 0xd9, 0x9e, 0xad, 0xe8, 0xd2, 0xeb, 0x51,
	0x54, 0x80, 0x03, 0x9e, 0x43, 0xc3, 0x70, 0x54, 0x55, 0x79, 0x63, 0x41, 0x56, 0x8a, 0x76, 0x11,
	0x8e, 0x90, 0x6f, 0xa8, 0x5c, 0x6f, 0xf7, 0x1b, 0x6f, 0xad, 0x74, 0xe6, 0x92, 0xd2, 0x09, 0xc9,
	0xbd, 0xa0, 0x91, 0x54, 0x2e, 0x33, 0x11, 0xd8, 0x42, 0xd3, 0x2e, 0xdd, 0x75, 0x22, 0x5f, 0x12,
	0x30, 0x87, 0xa2, 0x45, 0x81, 0xf4, 0x7c, 0xe3, 0x9d, 0x76, 0x75, 0x3e, 0xe3, 0x4a, 0x7a, 0xf0,
	0x4c, 0xe9, 0x34, 0xc3, 0x64, 0x2e, 0x13, 0xb4, 0xaa, 0xc8, 0x1d, 0x05, 0xe2, 0xc7, 0x08, 0x37,
	0x68, 0x40, 0xb9, 0x6e, 0xc7, 0xd0, 0x91, 0xb0, 0x73, 0x02, 0xe3, 0xfd, 0x60, 0x65, 0x53, 0x1d,
	0xac, 0xa2, 0x29, 0xbc, 0x8e, 0x26, 0xb5, 0x26, 0xa8, 0x0b, 0xd9, 0x87, 0xde, 0x6e, 0x0b, 0xe3,
	0x83, 0x95, 0x67, 0x57, 0x29, 0xea, 0x01, 0xe3, 0xcf, 0xca, 0x80, 0xc0, 0xb0, 0xf4, 0x2c, 0xbc,
	0x8f, 0xd6, 0x99, 0x37, 0x5e, 0x15, 0xcd, 0xef, 0x93, 0xf8, 0xfd, 0x24, 0x12, 0x4e, 0xcd, 0xa7,
	0x24, 0x04, 0x87, 0xf0, 0x23, 0xa0, 0xeb, 0x3e, 0xe5, 0xf0, 0x37, 0xbb, 0xbf, 0xad, 0xe8, 0x9d,
	0x18, 0xae, 0x50, 0x5e, 0x89, 0x51, 0xbc, 0x83, 0x8c, 0x94, 0x3a, 0x97, 0x8a, 0x3a, 0xf7, 0xc2,
	0xb8, 0x99, 0x0f, 0x06, 0xa7, 0x6d, 0xae, 0x4b, 0x65, 0xb9, 0x83, 0xe2, 0x55, 0x34, 0x9e, 0x96,
	0x6b, 0x7c, 0xce, 0xa1, 0xf1, 0xbf, 0x6e, 0x8d, 0xd0, 0x1b, 0xa7, 0xac, 0xec, 0x2f, 0xd6, 0x5f,
	0xed, 0xec, 0xfb, 0xe9, 0xbd, 0x7b, 0x68, 0x9d, 0x6d, 0xf1, 0x42, 0xa6, 0x6a, 0xd0, 0x23, 0xbb,
	0x9e, 0xec, 0xcd, 0xd4, 0xd7, 0x1c, 0x99, 0x4a, 0xe0, 0x6c, 0xa6, 0xa0, 0x3d, 0x52, 0x6e, 0xe1,
	0x7d, 0x47, 0x79, 0xda, 0xa3, 0xcb, 0x1d, 0xbc, 0xad, 0x82, 0x66, 0x7c, 0x47, 0xc8, 0xae, 0x75,
	0x4c, 0x5d, 0xe2, 0x48, 0xe3, 0xdb, 0xd1, 0x50, 0xae, 0x49, 0x52, 0xf0, 0xc9, 0x92, 0xa6, 0xee,
	0x8a, 0x5c, 0xad, 0xa2, 0x05, 0x8f, 0x65, 0xf6, 0x47, 0xe7, 0x43, 0xe4, 0xc5, 0xed, 0x06, 0x13,
	0xee, 0xab, 0xe3, 0x7b, 0x37, 0xf7, 0xb7, 0x4a, 0x6d, 0x38, 0xfe, 0x2a, 0xb8, 0xf1, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0x84, 0x78, 0x4d, 0xe4, 0x08, 0x00, 0x00,
}
