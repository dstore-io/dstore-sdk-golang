// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyOrderInformation_Pu.proto
// DO NOT EDIT!

/*
Package om_ModifyOrderInformation_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyOrderInformation_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyOrderInformation_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	OrderId                        *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderIdNull                    bool                        `protobuf:"varint,1004,opt,name=order_id_null,json=orderIdNull" json:"order_id_null,omitempty"`
	InformationTypeId              *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull          bool                        `protobuf:"varint,1005,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	Information                    *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=information" json:"information,omitempty"`
	InformationNull                bool                        `protobuf:"varint,1006,opt,name=information_null,json=informationNull" json:"information_null,omitempty"`
	Country                        *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=country" json:"country,omitempty"`
	CountryNull                    bool                        `protobuf:"varint,1007,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1008,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Parameters) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	InformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	ErrorCode         *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	OrderId           *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Response_Row) GetErrorCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Response_Row) GetOrderId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyOrderInformation_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyOrderInformation_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyOrderInformation_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x56, 0x9b, 0xa6, 0x4e, 0xa7, 0xfd, 0xdb, 0xfe, 0x2e, 0x07, 0x93, 0x56, 0x55, 0x69, 0x85,
	0x84, 0x7a, 0xe1, 0x22, 0x8e, 0x95, 0x8a, 0x84, 0x00, 0x71, 0x11, 0xa0, 0xa1, 0x58, 0x15, 0x12,
	0x08, 0xc9, 0x32, 0xf1, 0x26, 0x5a, 0x29, 0xf1, 0x86, 0x5d, 0x9b, 0x2a, 0x97, 0xbc, 0x01, 0x87,
	0x67, 0xe2, 0x41, 0xb8, 0xe4, 0xcc, 0x23, 0x30, 0xde, 0xd9, 0x24, 0x8e, 0x9b, 0xd6, 0x51, 0x6f,
	0xe2, 0xac, 0xe7, 0x3b, 0x8c, 0xc7, 0xbb, 0x9f, 0x61, 0x2f, 0x54, 0xb1, 0x90, 0x6c, 0x87, 0x45,
	0x2d, 0x1e, 0xb1, 0x9d, 0xae, 0x14, 0x0d, 0x16, 0x26, 0x92, 0xa9, 0x1d, 0xd1, 0xf1, 0xf7, 0x45,
	0xc8, 0x9b, 0xbd, 0x67, 0x32, 0x64, 0xb2, 0x16, 0x35, 0x85, 0xec, 0x04, 0x31, 0x17, 0x91, 0x7f,
	0x90, 0xb8, 0x88, 0x8b, 0x85, 0xbd, 0x4d, 0x64, 0x97, 0xc8, 0xee, 0x69, 0x8c, 0xea, 0x8a, 0x31,
	0x7a, 0x17, 0xb4, 0x13, 0xa6, 0x48, 0xa0, 0x7a, 0x69, 0xd4, 0x9d, 0x49, 0x29, 0xa4, 0x29, 0xad,
	0x8e, 0x96, 0x3a, 0x4c, 0xa9, 0xa0, 0xc5, 0x4c, 0x71, 0x2b, 0x5f, 0x8c, 0x03, 0x3e, 0xb4, 0x23,
	0xd0, 0xe6, 0x17, 0x0b, 0xe0, 0x20, 0x90, 0x01, 0x56, 0x99, 0x54, 0xf6, 0x6b, 0x58, 0xeb, 0xe2,
	0x15, 0xbb, 0xe1, 0x21, 0x8b, 0x62, 0xde, 0xe4, 0x0d, 0x6a, 0x8e, 0x3a, 0x72, 0xa6, 0x36, 0xa6,
	0xae, 0xce, 0x5f, 0xaf, 0xba, 0xe6, 0x99, 0x4c, 0x9f, 0x2a, 0x96, 0x3c, 0x6a, 0xbd, 0x48, 0x17,
	0x5e, 0x95, 0xf8, 0xb5, 0x11, 0xba, 0x2e, 0x29, 0xfb, 0x31, 0x5c, 0x3e, 0x4d, 0xdd, 0x8f, 0x92,
	0x76, 0xdb, 0xf9, 0x66, 0xa1, 0x47, 0xc5, 0x5b, 0x3f, 0x59, 0xa7, 0x8e, 0x30, 0xfb, 0x3e, 0x2c,
	0x1a, 0xad, 0xb8, 0xd7, 0x65, 0x28, 0xe8, 0x4c, 0xeb, 0xde, 0x56, 0x73, 0xbd, 0xf1, 0x28, 0x66,
	0x2d, 0x26, 0xa9, 0xb9, 0x05, 0xa2, 0x1c, 0x22, 0xa3, 0x16, 0xda, 0x2e, 0xac, 0x8c, 0x4a, 0x50,
	0x03, 0xdf, 0xa9, 0x81, 0xe5, 0x2c, 0x56, 0x5b, 0xde, 0x81, 0xb9, 0x24, 0xe2, 0x6f, 0x13, 0xed,
	0x56, 0x2a, 0x9c, 0x44, 0x85, 0xc0, 0x68, 0x74, 0x05, 0x16, 0x07, 0x44, 0xf2, 0xf8, 0x41, 0x1e,
	0x0b, 0x7d, 0x88, 0xd6, 0xbf, 0x0d, 0x15, 0x91, 0x6e, 0x8a, 0x54, 0x7e, 0xa6, 0xf8, 0x61, 0x2c,
	0x0d, 0x46, 0xf9, 0x2d, 0xf8, 0xaf, 0xcf, 0x23, 0xf5, 0x9f, 0xa4, 0x3e, 0x6f, 0x00, 0x5a, 0xfc,
	0x09, 0xac, 0x64, 0xde, 0xfe, 0x60, 0x68, 0xe5, 0x62, 0x9f, 0xff, 0x33, 0x3c, 0x33, 0xb9, 0x5d,
	0x70, 0xc6, 0x88, 0x91, 0xf9, 0x2f, 0x32, 0x3f, 0x7f, 0x8c, 0xa5, 0xdb, 0xb8, 0x0b, 0xf3, 0x99,
	0x82, 0x33, 0x5b, 0x38, 0xc5, 0x2c, 0xdc, 0xde, 0x86, 0xe5, 0xac, 0xaf, 0xf6, 0xfb, 0x4d, 0x7e,
	0x4b, 0x99, 0x82, 0x76, 0xba, 0x09, 0x56, 0x43, 0x24, 0x51, 0x2c, 0x7b, 0x8e, 0x55, 0xe8, 0xd2,
	0x87, 0xda, 0x9b, 0xb0, 0x60, 0xfe, 0x92, 0xfa, 0x1f, 0x33, 0x4a, 0x73, 0x53, 0x2b, 0x3f, 0x87,
	0x8b, 0x8a, 0x75, 0xf1, 0xd0, 0xa0, 0x9a, 0xcf, 0xcd, 0x66, 0x4e, 0xf7, 0xb0, 0x72, 0x2a, 0x85,
	0x4e, 0xe7, 0x06, 0xd4, 0x1a, 0x6d, 0x6e, 0xbc, 0xad, 0xec, 0x7b, 0xb0, 0x76, 0x82, 0x24, 0xb5,
	0xf1, 0x97, 0xda, 0x70, 0xc6, 0x91, 0xd3, 0x9e, 0x36, 0xdf, 0xcf, 0x40, 0xc5, 0x63, 0xaa, 0x2b,
	0x22, 0xc5, 0xec, 0x6b, 0x50, 0xd6, 0x29, 0x91, 0x3f, 0xae, 0x26, 0x82, 0x28, 0x41, 0x1e, 0xa5,
	0xbf, 0x1e, 0x01, 0xed, 0x97, 0xb0, 0x9c, 0xe6, 0x83, 0x9f, 0x7d, 0x37, 0xd3, 0x1b, 0x25, 0x24,
	0xbb, 0x39, 0x72, 0x3e, 0x46, 0xf6, 0x71, 0x9d, 0x49, 0x31, 0x6f, 0xa9, 0x33, 0x7a, 0x03, 0xf7,
	0x8a, 0x65, 0x72, 0x09, 0xcf, 0x4c, 0xaa, 0xb8, 0x7e, 0x4c, 0x91, 0x52, 0x6b, 0x9f, 0xae, 0x5e,
	0x1f, 0x8e, 0x71, 0x51, 0x92, 0xe2, 0x08, 0x8f, 0x42, 0xca, 0xda, 0x75, 0x27, 0xcf, 0x51, 0xb7,
	0x3f, 0x09, 0xd7, 0x13, 0x47, 0x5e, 0x2a, 0x52, 0xfd, 0x3a, 0x05, 0x25, 0x5c, 0xd8, 0x17, 0x60,
	0x16, 0x97, 0xe9, 0xce, 0xff, 0x50, 0xc7, 0xe1, 0x94, 0xbd, 0x32, 0x2e, 0x71, 0x47, 0x3f, 0x1d,
	0x7f, 0x3c, 0x3e, 0xd6, 0xcf, 0x74, 0x3e, 0xf6, 0x00, 0xf4, 0x5c, 0xfd, 0x86, 0x08, 0x99, 0xf3,
	0x69, 0x02, 0x91, 0x39, 0x8d, 0x7f, 0x88, 0x70, 0x8c, 0x99, 0x61, 0x0c, 0x7c, 0xae, 0x4f, 0x9e,
	0x03, 0x0f, 0x0e, 0x61, 0x95, 0x8b, 0xdc, 0x98, 0x86, 0xdf, 0xaa, 0x57, 0xb7, 0xce, 0xf4, 0x15,
	0x7b, 0x33, 0xab, 0x3f, 0x14, 0x37, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x72, 0xfb, 0xeb,
	0x05, 0x07, 0x00, 0x00,
}
