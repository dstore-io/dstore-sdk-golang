// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_DeletePostingBinary_Pu.proto
// DO NOT EDIT!

/*
Package fo_DeletePostingBinary_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_DeletePostingBinary_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_DeletePostingBinary_Pu

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
	PostingId                      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=posting_id,json=postingId" json:"posting_id,omitempty"`
	PostingIdNull                  bool                        `protobuf:"varint,1004,opt,name=posting_id_null,json=postingIdNull" json:"posting_id_null,omitempty"`
	SortNo                         *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	SortNoNull                     bool                        `protobuf:"varint,1005,opt,name=sort_no_null,json=sortNoNull" json:"sort_no_null,omitempty"`
	PostingBinaryIdentifier        *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=posting_binary_identifier,json=postingBinaryIdentifier" json:"posting_binary_identifier,omitempty"`
	PostingBinaryIdentifierNull    bool                        `protobuf:"varint,1006,opt,name=posting_binary_identifier_null,json=postingBinaryIdentifierNull" json:"posting_binary_identifier_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1007,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingId
	}
	return nil
}

func (m *Parameters) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Parameters) GetPostingBinaryIdentifier() *dstore_values.StringValue {
	if m != nil {
		return m.PostingBinaryIdentifier
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_DeletePostingBinary_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_DeletePostingBinary_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_DeletePostingBinary_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x4f,
	0x14, 0xa7, 0xcd, 0x3f, 0x49, 0x7b, 0xfe, 0xb1, 0x2d, 0x53, 0xb1, 0xdb, 0xa4, 0x06, 0x5b, 0x91,
	0x7a, 0xb5, 0x91, 0xaa, 0x58, 0xbc, 0x11, 0x4b, 0x45, 0x56, 0x68, 0x48, 0x17, 0x29, 0x28, 0xc2,
	0xb2, 0x6d, 0x26, 0x61, 0x20, 0x99, 0x59, 0x67, 0x76, 0x2d, 0x7d, 0x0b, 0x5f, 0xc7, 0x47, 0xf2,
	0xfb, 0xda, 0x3b, 0xcf, 0xce, 0x99, 0x7c, 0xad, 0xc6, 0xe8, 0x4d, 0xb7, 0xb3, 0xe7, 0xf7, 0xb5,
	0x93, 0x73, 0x0e, 0x1c, 0x76, 0x4d, 0xaa, 0x34, 0x6f, 0x71, 0xd9, 0x17, 0x92, 0xb7, 0x12, 0xad,
	0x2e, 0x78, 0x37, 0xd3, 0xdc, 0xb4, 0x7a, 0x2a, 0x3a, 0xe6, 0x03, 0x9e, 0xf2, 0x8e, 0x32, 0xa9,
	0x90, 0xfd, 0x23, 0x21, 0x63, 0x7d, 0x15, 0x75, 0x32, 0x1f, 0x41, 0xa9, 0x62, 0xfb, 0xc4, 0xf4,
	0x89, 0xe9, 0xcf, 0x85, 0xd7, 0x37, 0x9d, 0xc5, 0xbb, 0x78, 0x90, 0x71, 0x43, 0xec, 0xfa, 0xf6,
	0xac, 0x2f, 0xd7, 0x5a, 0x69, 0x57, 0x6a, 0xcc, 0x96, 0x86, 0xdc, 0x98, 0xb8, 0xcf, 0x5d, 0xf1,
	0x76, 0xb1, 0x98, 0xc6, 0x42, 0xf6, 0x94, 0x1e, 0xc6, 0xa9, 0x50, 0x92, 0x40, 0x7b, 0x3f, 0x2a,
	0x00, 0x9d, 0x58, 0xc7, 0x58, 0xe5, 0xda, 0xb0, 0x37, 0xb0, 0x93, 0xe0, 0x53, 0xc9, 0x48, 0x74,
	0xb9, 0x4c, 0x45, 0x4f, 0x5c, 0x58, 0x74, 0x44, 0x89, 0xbc, 0xa5, 0x5b, 0x4b, 0x77, 0xff, 0x3f,
	0xa8, 0xfb, 0xee, 0x83, 0x5c, 0x4e, 0x93, 0x6a, 0xfc, 0x8c, 0xb3, 0xfc, 0x10, 0xd6, 0x89, 0x1f,
	0xcc, 0xd0, 0x6d, 0xc9, 0xb0, 0x17, 0xb0, 0xfb, 0x27, 0xf5, 0x48, 0x66, 0x83, 0x81, 0xf7, 0xb1,
	0x8a, 0x1e, 0x2b, 0x61, 0x73, 0xbe, 0x4e, 0x1b, 0x61, 0xec, 0x29, 0xac, 0x39, 0xad, 0xf4, 0x2a,
	0xe1, 0x28, 0xe8, 0x2d, 0xdb, 0x6c, 0x8d, 0x42, 0x36, 0x21, 0x53, 0xde, 0xe7, 0x9a, 0xc2, 0xd5,
	0x88, 0xf2, 0x12, 0x19, 0x41, 0x97, 0xf9, 0xb0, 0x39, 0x2b, 0x41, 0x01, 0x3e, 0x51, 0x80, 0x8d,
	0x69, 0xac, 0xb5, 0x7c, 0x04, 0xab, 0x99, 0x14, 0x6f, 0x33, 0xeb, 0x56, 0x5a, 0x78, 0x13, 0x2b,
	0x04, 0x46, 0xa3, 0x3b, 0xb0, 0x36, 0x26, 0x92, 0xc7, 0x67, 0xf2, 0xa8, 0x8d, 0x20, 0x56, 0xff,
	0x31, 0x40, 0x42, 0x1d, 0x91, 0x1b, 0xfc, 0xb7, 0xf8, 0x73, 0x56, 0x1d, 0x1c, 0x2d, 0xf6, 0x61,
	0x7d, 0xc2, 0x25, 0x8f, 0x2f, 0xe4, 0x71, 0x6d, 0x0c, 0xb2, 0x26, 0x0f, 0xa0, 0x6a, 0x94, 0x4e,
	0x23, 0xa9, 0xbc, 0xf2, 0x62, 0x87, 0x4a, 0x8e, 0x6d, 0x2b, 0xb6, 0x0b, 0x35, 0xc7, 0x22, 0xed,
	0xaf, 0xa4, 0x0d, 0x54, 0xb6, 0xc2, 0x67, 0xb0, 0x3d, 0x4a, 0x70, 0x4e, 0x0d, 0x3d, 0xfa, 0x91,
	0xb9, 0xf6, 0x2a, 0x0b, 0x6f, 0x6b, 0x2b, 0x99, 0x1e, 0x86, 0x60, 0x4c, 0x65, 0xc7, 0xd0, 0x9c,
	0xab, 0x4b, 0x61, 0xbe, 0x51, 0x98, 0xc6, 0x1c, 0x05, 0x9b, 0xee, 0x14, 0xb6, 0x0c, 0x4f, 0xb0,
	0xd1, 0xd1, 0x3f, 0x12, 0xae, 0x01, 0xf3, 0xbe, 0x33, 0x5e, 0x75, 0x61, 0xb6, 0xeb, 0x63, 0x6a,
	0x40, 0x0d, 0x89, 0xaf, 0x0d, 0x7b, 0x02, 0x3b, 0x73, 0x24, 0x29, 0xd6, 0x77, 0x8a, 0xe5, 0xfd,
	0x8e, 0x9c, 0x67, 0xda, 0xfb, 0xb0, 0x0c, 0x2b, 0x21, 0x37, 0x89, 0x92, 0x86, 0xb3, 0x7b, 0x50,
	0xb6, 0x93, 0x5d, 0x1c, 0x31, 0xb7, 0x33, 0x68, 0xea, 0x9f, 0xe5, 0x7f, 0x43, 0x02, 0xb2, 0x57,
	0xb0, 0x91, 0xcf, 0x74, 0x34, 0x35, 0xd4, 0x38, 0x03, 0x25, 0x24, 0xfb, 0x05, 0x72, 0x71, 0xf4,
	0x4f, 0xf0, 0x1c, 0x4c, 0xce, 0xe1, 0xfa, 0x70, 0xf6, 0x05, 0x3b, 0x84, 0xaa, 0xdb, 0x25, 0xd8,
	0xe7, 0xb9, 0x62, 0xf3, 0x17, 0x45, 0xda, 0x34, 0x27, 0xf4, 0x0c, 0x47, 0x70, 0xf6, 0x1c, 0x4a,
	0x5a, 0x5d, 0x62, 0xf3, 0xe6, 0xac, 0x87, 0xfe, 0x5f, 0x2e, 0x3e, 0x7f, 0x74, 0x0d, 0x7e, 0xa8,
	0x2e, 0xc3, 0x5c, 0xa1, 0x7e, 0x13, 0x4a, 0xf8, 0x3f, 0xbb, 0x01, 0x15, 0x3c, 0xe5, 0xf3, 0xf0,
	0xbe, 0x8d, 0x17, 0x53, 0x0e, 0xcb, 0x78, 0x0c, 0xba, 0x47, 0xa7, 0xd0, 0x10, 0xaa, 0x20, 0x3f,
	0xd9, 0xc8, 0xaf, 0x0f, 0xfe, 0x7d, 0x57, 0x9f, 0x57, 0xec, 0x46, 0xbc, 0xff, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x9a, 0x0f, 0xa1, 0xe4, 0xe8, 0x05, 0x00, 0x00,
}
