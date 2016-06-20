// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetTrolleySurchargeInfo_Pu.proto
// DO NOT EDIT!

/*
Package om_GetTrolleySurchargeInfo_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetTrolleySurchargeInfo_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetTrolleySurchargeInfo_Pu

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
	UniqueId              *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull          bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PositionNo            *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=position_no,json=positionNo" json:"position_no,omitempty"`
	PositionNoNull        bool                        `protobuf:"varint,1002,opt,name=position_no_null,json=positionNoNull" json:"position_no_null,omitempty"`
	InformationTypeId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull bool                        `protobuf:"varint,1003,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPositionNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PositionNo
	}
	return nil
}

func (m *Parameters) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
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
	PositionNo        *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=position_no,json=positionNo" json:"position_no,omitempty"`
	InformationType   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	Information       *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=information" json:"information,omitempty"`
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

func (m *Response_Row) GetPositionNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.PositionNo
	}
	return nil
}

func (m *Response_Row) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func (m *Response_Row) GetInformation() *dstore_values.StringValue {
	if m != nil {
		return m.Information
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetTrolleySurchargeInfo_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetTrolleySurchargeInfo_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetTrolleySurchargeInfo_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetTrolleySurchargeInfo_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xef, 0x6a, 0x13, 0x41,
	0x10, 0xa7, 0x49, 0xd3, 0xc6, 0xa9, 0xb4, 0xf1, 0x8a, 0x72, 0x26, 0x20, 0x52, 0x11, 0x14, 0xe1,
	0x22, 0xfa, 0x41, 0x85, 0x8a, 0x20, 0xa8, 0x04, 0xed, 0x51, 0xd6, 0x22, 0xe8, 0x07, 0x8f, 0xb3,
	0x37, 0x9e, 0x8b, 0x97, 0xdd, 0xb8, 0x7b, 0x67, 0xe9, 0x5b, 0xf8, 0xef, 0x45, 0x7c, 0x15, 0x3f,
	0xfb, 0x02, 0xea, 0x4b, 0x38, 0x7b, 0x7b, 0x67, 0x2e, 0xdb, 0x60, 0x48, 0xbf, 0xb4, 0xcc, 0xcd,
	0xef, 0x37, 0xbf, 0x99, 0xf9, 0x4d, 0x16, 0x76, 0x13, 0x9d, 0x4b, 0x85, 0x43, 0x14, 0x29, 0x17,
	0x38, 0x9c, 0x28, 0x79, 0x88, 0x49, 0xa1, 0x50, 0x0f, 0xe5, 0x38, 0x7a, 0x82, 0xf9, 0x81, 0x92,
	0x59, 0x86, 0xc7, 0xcf, 0x0b, 0x75, 0xf8, 0x2e, 0x56, 0x29, 0x8e, 0xc4, 0x5b, 0x19, 0xed, 0x17,
	0x01, 0x01, 0x73, 0xe9, 0xdd, 0xb0, 0xec, 0xc0, 0xb2, 0x83, 0xff, 0x52, 0xfa, 0xdb, 0x95, 0xd4,
	0xc7, 0x38, 0x2b, 0x50, 0xdb, 0x0a, 0xfd, 0x8b, 0xb3, 0xfa, 0xa8, 0x94, 0x54, 0x55, 0x6a, 0x30,
	0x9b, 0x1a, 0xa3, 0xd6, 0x71, 0x8a, 0x55, 0xf2, 0x8a, 0x9b, 0xcc, 0x63, 0x4e, 0x3a, 0x6a, 0x1c,
	0xe7, 0x5c, 0x0a, 0x0b, 0xda, 0xf9, 0xd9, 0x02, 0xd8, 0x8f, 0x55, 0x4c, 0x59, 0x54, 0xda, 0xbb,
	0x03, 0x67, 0x0a, 0xc1, 0x3f, 0x14, 0x18, 0xf1, 0xc4, 0x5f, 0xb9, 0xbc, 0x72, 0x6d, 0xe3, 0x56,
	0x3f, 0xa8, 0x26, 0xa8, 0x9a, 0xd2, 0xb9, 0xe2, 0x22, 0x7d, 0x61, 0x02, 0xd6, 0xb5, 0xe0, 0x51,
	0xe2, 0x5d, 0x85, 0xcd, 0x7f, 0xc4, 0x48, 0x14, 0x59, 0xe6, 0xff, 0x5a, 0x27, 0x7a, 0x97, 0x9d,
	0xad, 0x21, 0x21, 0x7d, 0xf4, 0x76, 0x61, 0x63, 0x22, 0x35, 0x37, 0x0d, 0x44, 0x42, 0xfa, 0xad,
	0x52, 0x61, 0xe0, 0x28, 0x70, 0x91, 0x63, 0x8a, 0xca, 0x4a, 0x40, 0x8d, 0x0f, 0xa5, 0x77, 0x1d,
	0x7a, 0x0d, 0xb6, 0x95, 0xf9, 0x6d, 0x65, 0x36, 0xa7, 0xb0, 0x52, 0xe8, 0x29, 0x6c, 0x37, 0x86,
	0x8d, 0xf2, 0xe3, 0x49, 0x39, 0x52, 0x7b, 0xb1, 0xe0, 0xb9, 0x06, 0xef, 0x80, 0x68, 0x34, 0xdc,
	0x5d, 0xf0, 0xe7, 0x14, 0xb3, 0xfa, 0x7f, 0xac, 0xfe, 0xf9, 0x13, 0x2c, 0xd3, 0xc6, 0xce, 0x8f,
	0x55, 0xe8, 0x32, 0xd4, 0x13, 0x29, 0x34, 0x7a, 0x37, 0xa1, 0x53, 0x9a, 0xe7, 0x2e, 0xb6, 0x3a,
	0x0d, 0x6b, 0xec, 0x23, 0xf3, 0x97, 0x59, 0xa0, 0xf7, 0x12, 0x7a, 0xc6, 0xb6, 0xa8, 0x51, 0x9c,
	0x76, 0xd6, 0x26, 0x72, 0xe0, 0x90, 0x5d, 0x77, 0xf7, 0x28, 0x1e, 0x4d, 0x63, 0xb6, 0x35, 0x9e,
	0xfd, 0x40, 0x33, 0xad, 0x57, 0xe7, 0x42, 0x4b, 0x31, 0x15, 0x2f, 0x9d, 0xa8, 0x68, 0x8f, 0x69,
	0xcf, 0xfe, 0x67, 0x35, 0x9c, 0x56, 0xdb, 0x56, 0xf2, 0xc8, 0x5f, 0x2d, 0x59, 0xf7, 0x82, 0x25,
	0xee, 0x3b, 0xa8, 0x57, 0x11, 0x30, 0x79, 0xc4, 0x4c, 0x95, 0xfe, 0xf7, 0x16, 0xb4, 0x29, 0xf0,
	0x2e, 0xc0, 0x1a, 0x85, 0xc6, 0xa2, 0x4f, 0x21, 0x6d, 0xa7, 0xc3, 0x3a, 0x14, 0xd2, 0xea, 0x9f,
	0xcd, 0xf7, 0xf1, 0x73, 0x78, 0x2a, 0x23, 0xef, 0xcf, 0x9e, 0xdf, 0x97, 0x70, 0xb9, 0xfb, 0x7b,
	0x0c, 0x3d, 0xb7, 0x19, 0xff, 0x6b, 0xb8, 0xf0, 0x57, 0xb2, 0xe5, 0x34, 0x62, 0xda, 0x68, 0x3a,
	0xfa, 0x6d, 0x71, 0x89, 0x26, 0xfe, 0xe1, 0x6b, 0x18, 0x70, 0xe9, 0xec, 0x7d, 0xfa, 0x2a, 0xbd,
	0x7a, 0x90, 0x4a, 0x9d, 0xbc, 0xaf, 0xf3, 0xc9, 0xd2, 0x0f, 0xd7, 0x9b, 0xb5, 0xf2, 0x69, 0xb8,
	0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x92, 0x3c, 0x37, 0xa3, 0xf9, 0x04, 0x00, 0x00,
}