// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetOrderInformationTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetOrderInformationTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetOrderInformationTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetOrderInformationTypes_Ad

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
	InformationTypeId          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	InformationTypeIdNull      bool                        `protobuf:"varint,1001,opt,name=information_type_id_null,json=informationTypeIdNull" json:"information_type_id_null,omitempty"`
	PersonCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	PersonCharacteristicIdNull bool                        `protobuf:"varint,1002,opt,name=person_characteristic_id_null,json=personCharacteristicIdNull" json:"person_characteristic_id_null,omitempty"`
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

func (m *Parameters) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CharacteristicDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	FieldTypeId               *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=field_type_id,json=fieldTypeId" json:"field_type_id,omitempty"`
	PersonCharacteristicId    *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=person_characteristic_id,json=personCharacteristicId" json:"person_characteristic_id,omitempty"`
	InformationType           *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=information_type,json=informationType" json:"information_type,omitempty"`
	OrderInformationTypeId    *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=order_information_type_id,json=orderInformationTypeId" json:"order_information_type_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetFieldTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FieldTypeId
	}
	return nil
}

func (m *Response_Row) GetPersonCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetInformationType() *dstore_values.StringValue {
	if m != nil {
		return m.InformationType
	}
	return nil
}

func (m *Response_Row) GetOrderInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderInformationTypeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetOrderInformationTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetOrderInformationTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetOrderInformationTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x1b, 0xd2, 0x56, 0x53, 0xa1, 0x96, 0xad, 0xa8, 0x1c, 0x47, 0x20, 0x54, 0x2e, 0x1c,
	0x90, 0x83, 0xe0, 0x00, 0x42, 0x42, 0x82, 0xf2, 0xa7, 0x08, 0x6a, 0xd0, 0x0a, 0x82, 0xe8, 0xc5,
	0x32, 0xde, 0x69, 0x58, 0x29, 0xf1, 0x5a, 0xbb, 0x1b, 0x2a, 0xde, 0x82, 0xdf, 0xd7, 0xe0, 0xc4,
	0xb3, 0xf0, 0x0e, 0xf0, 0x14, 0x8c, 0xbd, 0x5b, 0x12, 0xbb, 0xa1, 0x81, 0x5e, 0x12, 0x8d, 0x67,
	0xbe, 0xd9, 0x6f, 0xe6, 0xfb, 0x76, 0xe1, 0x8e, 0x30, 0x56, 0x69, 0xec, 0x61, 0x3e, 0x94, 0x39,
	0xf6, 0x0a, 0xad, 0x32, 0x14, 0x13, 0x8d, 0xa6, 0xa7, 0xc6, 0xc9, 0x63, 0xb4, 0xcf, 0xb4, 0x40,
	0xdd, 0xcf, 0x0f, 0x94, 0x1e, 0xa7, 0x56, 0xaa, 0xfc, 0xc5, 0xfb, 0x02, 0x4d, 0x72, 0x4f, 0x44,
	0x54, 0x69, 0x15, 0xbb, 0xea, 0xe0, 0x91, 0x83, 0x47, 0x27, 0x63, 0xc2, 0x2d, 0x7f, 0xd8, 0xbb,
	0x74, 0x34, 0x41, 0xe3, 0x5a, 0x84, 0x9d, 0x3a, 0x03, 0xd4, 0x5a, 0x69, 0x9f, 0xea, 0xd6, 0x53,
	0x63, 0x34, 0x26, 0x1d, 0xa2, 0x4f, 0x5e, 0x6e, 0x26, 0x6d, 0x2a, 0xa7, 0x67, 0xba, 0xa2, 0x9d,
	0x6f, 0xcb, 0x00, 0xcf, 0x53, 0x9d, 0x52, 0x16, 0xb5, 0x61, 0x4f, 0x60, 0x6b, 0xa6, 0x26, 0xb1,
	0x44, 0x2c, 0x91, 0x22, 0x58, 0xba, 0xb4, 0x74, 0x65, 0xfd, 0x7a, 0x37, 0xf2, 0xc3, 0x78, 0x7a,
	0x32, 0xb7, 0x38, 0x44, 0x3d, 0x28, 0x23, 0x7e, 0x4e, 0xd6, 0xe7, 0xe9, 0x0b, 0x76, 0x0b, 0x82,
	0x39, 0xcd, 0x92, 0x7c, 0x32, 0x1a, 0x05, 0x3f, 0x57, 0xa9, 0xe5, 0x1a, 0x3f, 0x7f, 0x0c, 0x15,
	0x53, 0x96, 0xbd, 0x84, 0xa0, 0x20, 0x3a, 0x04, 0xca, 0xde, 0x12, 0xb9, 0x8c, 0xb8, 0x49, 0x63,
	0x65, 0x56, 0x72, 0x59, 0x5e, 0xcc, 0x65, 0xdb, 0x81, 0xef, 0xd7, 0xb0, 0x44, 0x68, 0x17, 0x2e,
	0xfc, 0xad, 0xad, 0x63, 0xf5, 0xcb, 0xb1, 0x0a, 0xe7, 0xe3, 0x4b, 0x6a, 0x3b, 0x3f, 0xda, 0xb0,
	0xc6, 0xd1, 0x14, 0x2a, 0x37, 0xc8, 0xae, 0x41, 0xbb, 0x92, 0xc3, 0x2f, 0x28, 0x8c, 0xea, 0x6a,
	0x3b, 0xa9, 0x1e, 0x96, 0xbf, 0xdc, 0x15, 0xb2, 0xd7, 0xb0, 0x59, 0x0a, 0x91, 0xcc, 0xcc, 0x4d,
	0x13, 0xb5, 0x08, 0x1c, 0x35, 0xc0, 0x4d, 0xbd, 0xf6, 0x28, 0x9e, 0xf1, 0x0c, 0xdf, 0x18, 0xd7,
	0x3f, 0xd0, 0xba, 0x57, 0xbd, 0x01, 0x82, 0x56, 0xd5, 0xf1, 0xe2, 0xb1, 0x8e, 0xce, 0x1e, 0x7b,
	0xee, 0x9f, 0x1f, 0x95, 0xb3, 0xa7, 0xd0, 0xd2, 0xea, 0x30, 0x38, 0x53, 0xa1, 0x6e, 0x47, 0xff,
	0x63, 0xd9, 0xe8, 0x68, 0x17, 0x11, 0x57, 0x87, 0xbc, 0x6c, 0x13, 0x7e, 0x6f, 0x41, 0x8b, 0x02,
	0xb6, 0x0d, 0x2b, 0x14, 0x96, 0x92, 0x7d, 0x88, 0x69, 0x3d, 0x6d, 0xde, 0xa6, 0x90, 0x54, 0xd8,
	0x87, 0xb0, 0xb1, 0x7e, 0x81, 0x26, 0xd3, 0xb2, 0xa8, 0x96, 0xf1, 0x31, 0xae, 0xaf, 0xd2, 0xeb,
	0x6b, 0xac, 0x96, 0xf9, 0xd0, 0xc9, 0xdb, 0xa9, 0xc3, 0x1f, 0x4c, 0xd1, 0xec, 0x2e, 0x9c, 0x3d,
	0x90, 0x38, 0x12, 0x7f, 0x9c, 0xfb, 0x29, 0x5e, 0x6c, 0x97, 0xf5, 0x0a, 0xe2, 0x4d, 0x3b, 0x38,
	0xc1, 0x7a, 0x9f, 0xe3, 0xd3, 0x7b, 0xef, 0x11, 0x6c, 0x36, 0x2f, 0x43, 0xf0, 0x65, 0xf1, 0xac,
	0x1b, 0x8d, 0x0b, 0xc2, 0x5e, 0x41, 0x47, 0x95, 0x5a, 0x24, 0xf3, 0xee, 0xe9, 0xd7, 0x7f, 0x21,
	0xa8, 0xe6, 0x48, 0xd9, 0x17, 0xbb, 0x03, 0xe8, 0x4a, 0xd5, 0xd0, 0x7e, 0xfa, 0xda, 0xed, 0xdf,
	0x3c, 0xe5, 0x3b, 0xf8, 0x66, 0xa5, 0x7a, 0x68, 0x6e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xdc,
	0x33, 0x0d, 0x9a, 0x49, 0x05, 0x00, 0x00,
}
