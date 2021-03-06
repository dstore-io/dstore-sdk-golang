// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/pm_GetPossibleDuplicates_Ad.proto
// DO NOT EDIT!

/*
Package pm_GetPossibleDuplicates_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/pm_GetPossibleDuplicates_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package pm_GetPossibleDuplicates_Ad

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
	PersonId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull               bool                        `protobuf:"varint,1001,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	OutputCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=output_characteristic_id,json=outputCharacteristicId" json:"output_characteristic_id,omitempty"`
	OutputCharacteristicIdNull bool                        `protobuf:"varint,1002,opt,name=output_characteristic_id_null,json=outputCharacteristicIdNull" json:"output_characteristic_id_null,omitempty"`
	RowCount                   *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=row_count,json=rowCount" json:"row_count,omitempty"`
	RowCountNull               bool                        `protobuf:"varint,1003,opt,name=row_count_null,json=rowCountNull" json:"row_count_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId
	}
	return nil
}

func (m *Parameters) GetRowCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.RowCount
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
	RowId    int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	PersonId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	Value    *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=value" json:"value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.pm_GetPossibleDuplicates_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.pm_GetPossibleDuplicates_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.pm_GetPossibleDuplicates_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/pm_GetPossibleDuplicates_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x8a, 0xd4, 0x40,
	0x10, 0x65, 0x36, 0xce, 0xee, 0xd8, 0x8a, 0x4a, 0x0b, 0x4b, 0xcc, 0xa0, 0xc8, 0x8a, 0xa0, 0x2f,
	0x3d, 0x5e, 0x1e, 0xbc, 0x80, 0x0f, 0xee, 0x2a, 0x32, 0x0f, 0x3b, 0x2c, 0x0d, 0x0a, 0x8a, 0x10,
	0xb2, 0x49, 0x19, 0x1b, 0x33, 0xdd, 0xa1, 0xbb, 0xe3, 0x7e, 0x84, 0x2f, 0xea, 0x67, 0xf8, 0x67,
	0xab, 0x3f, 0x61, 0x75, 0xba, 0x33, 0x6e, 0xe2, 0x6d, 0xe7, 0x65, 0x42, 0xa5, 0xce, 0xa9, 0x53,
	0xa7, 0xaa, 0x32, 0xe4, 0x71, 0x61, 0xac, 0xd2, 0x30, 0x03, 0x59, 0x0a, 0x09, 0xb3, 0x5a, 0xab,
	0x1c, 0x8a, 0x46, 0x83, 0x99, 0xd5, 0xcb, 0xf4, 0x05, 0xd8, 0x03, 0x65, 0x8c, 0x38, 0xac, 0xe0,
	0x59, 0x53, 0x57, 0x22, 0xcf, 0x2c, 0x98, 0xf4, 0x69, 0xc1, 0x10, 0x66, 0x15, 0xbd, 0xed, 0xb9,
	0xcc, 0x73, 0xd9, 0x3f, 0x08, 0xc9, 0xe5, 0x20, 0xf3, 0x31, 0xab, 0x1a, 0x30, 0x9e, 0x9f, 0x5c,
	0xe9, 0x6b, 0x83, 0xd6, 0x4a, 0x87, 0xd4, 0xb4, 0x9f, 0x5a, 0x82, 0x31, 0x59, 0x09, 0x21, 0x79,
	0x63, 0x98, 0xb4, 0x99, 0x90, 0xef, 0x94, 0x5e, 0x66, 0x56, 0x28, 0xe9, 0x41, 0x3b, 0xc7, 0x1b,
	0x84, 0x1c, 0x64, 0x3a, 0xc3, 0x2c, 0x68, 0x43, 0x1f, 0x92, 0xb3, 0x35, 0x3e, 0x95, 0x4c, 0x45,
	0x11, 0x8f, 0xae, 0x8f, 0x6e, 0x9d, 0xbb, 0x37, 0x65, 0xa1, 0xff, 0xd0, 0x94, 0x90, 0x16, 0x4a,
	0xd0, 0xaf, 0x5c, 0xc4, 0x27, 0x1e, 0x3d, 0x2f, 0xe8, 0x4d, 0x72, 0x61, 0xc5, 0x4c, 0x65, 0x53,
	0x55, 0xf1, 0xf1, 0x16, 0xf2, 0x27, 0xfc, 0x7c, 0x07, 0x59, 0xe0, 0x4b, 0xfa, 0x92, 0xc4, 0xaa,
	0xb1, 0x75, 0x63, 0xd3, 0xfc, 0x3d, 0xca, 0xe6, 0xa8, 0x2a, 0x8c, 0x15, 0xb9, 0xd3, 0xdb, 0xf8,
	0xbf, 0xde, 0xb6, 0x27, 0xef, 0xf5, 0xb8, 0xa8, 0xbe, 0x4b, 0xae, 0xfe, 0xad, 0xac, 0x6f, 0xe6,
	0xbb, 0x6f, 0x26, 0xf9, 0x33, 0xbf, 0x6d, 0x0d, 0xbd, 0x6b, 0x75, 0x94, 0xe6, 0xaa, 0x91, 0x36,
	0x8e, 0x4e, 0xe1, 0x1d, 0xd1, 0x7b, 0x0e, 0xec, 0xbc, 0xaf, 0x98, 0x5e, 0xee, 0x47, 0xf0, 0xde,
	0x41, 0x9c, 0xc0, 0xce, 0xb7, 0x88, 0x4c, 0x38, 0x98, 0x5a, 0x49, 0x03, 0xf4, 0x0e, 0x19, 0xb7,
	0x9b, 0x0c, 0x53, 0x4e, 0x58, 0xff, 0x4a, 0xfc, 0x96, 0x9f, 0xbb, 0x5f, 0xee, 0x81, 0xf4, 0x35,
	0xb9, 0xe4, 0x76, 0x98, 0x9e, 0x58, 0x22, 0x8e, 0x2c, 0x42, 0x32, 0x1b, 0x90, 0x87, 0xab, 0xde,
	0xc7, 0x78, 0xfe, 0x2b, 0xe6, 0x17, 0x97, 0xfd, 0x17, 0x68, 0x7d, 0x2b, 0xdc, 0x0e, 0x1a, 0x77,
	0x15, 0xaf, 0xfd, 0x56, 0xd1, 0x5f, 0xd6, 0xbe, 0x7f, 0xf2, 0x0e, 0x4e, 0xe7, 0x24, 0x42, 0x8f,
	0xf1, 0x99, 0x96, 0xf5, 0x80, 0x9d, 0xfa, 0xd4, 0x59, 0x37, 0x08, 0xc6, 0xd5, 0x11, 0x77, 0x35,
	0x92, 0x4f, 0x23, 0x12, 0x61, 0x40, 0xb7, 0xc9, 0xa6, 0x9b, 0x26, 0x1e, 0xc4, 0xe7, 0x05, 0xce,
	0x66, 0xcc, 0xc7, 0x18, 0xe2, 0x8e, 0x1f, 0x9d, 0xbc, 0xcd, 0x2f, 0x8b, 0x75, 0x8e, 0xf3, 0x2e,
	0x19, 0xb7, 0x80, 0xf8, 0xeb, 0xa2, 0x3f, 0xed, 0x40, 0x33, 0x56, 0x0b, 0x59, 0x7a, 0x96, 0x47,
	0xee, 0xbe, 0x25, 0x53, 0xa1, 0x86, 0x7e, 0x56, 0x9f, 0xfd, 0x9b, 0x27, 0xa5, 0x32, 0xc5, 0x87,
	0x2e, 0x5f, 0xac, 0xf9, 0xcf, 0x70, 0xb8, 0xd9, 0x7e, 0x7d, 0xf7, 0x7f, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x9a, 0x87, 0x81, 0xe1, 0x58, 0x04, 0x00, 0x00,
}
