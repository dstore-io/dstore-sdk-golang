// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_InsertBinaryForPosting_Pu.proto
// DO NOT EDIT!

/*
Package fo_InsertBinaryForPosting_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_InsertBinaryForPosting_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_InsertBinaryForPosting_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PostingId                      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=posting_id,json=postingId" json:"posting_id,omitempty"`
	PostingIdNull                  bool                        `protobuf:"varint,1004,opt,name=posting_id_null,json=postingIdNull" json:"posting_id_null,omitempty"`
	ContentType                    *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	ContentTypeNull                bool                        `protobuf:"varint,1005,opt,name=content_type_null,json=contentTypeNull" json:"content_type_null,omitempty"`
	Filename                       *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=filename" json:"filename,omitempty"`
	FilenameNull                   bool                        `protobuf:"varint,1006,opt,name=filename_null,json=filenameNull" json:"filename_null,omitempty"`
	Description                    *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	DescriptionNull                bool                        `protobuf:"varint,1007,opt,name=description_null,json=descriptionNull" json:"description_null,omitempty"`
	PostingBinaryIdentifier        *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=posting_binary_identifier,json=postingBinaryIdentifier" json:"posting_binary_identifier,omitempty"`
	PostingBinaryIdentifierNull    bool                        `protobuf:"varint,1008,opt,name=posting_binary_identifier_null,json=postingBinaryIdentifierNull" json:"posting_binary_identifier_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,9,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1009,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetContentType() *dstore_values.StringValue {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Parameters) GetFilename() *dstore_values.StringValue {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *Parameters) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_InsertBinaryForPosting_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_InsertBinaryForPosting_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_InsertBinaryForPosting_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/fo_InsertBinaryForPosting_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xfb, 0x6a, 0x13, 0x4d,
	0x14, 0xa7, 0xcd, 0x97, 0x36, 0x3d, 0xbd, 0x7e, 0x53, 0xb1, 0xdb, 0xa4, 0x16, 0xad, 0x8a, 0x52,
	0x61, 0x2b, 0x0a, 0x5a, 0xbc, 0x62, 0x51, 0x21, 0x42, 0x4b, 0x5d, 0xa4, 0xa0, 0x28, 0xcb, 0x36,
	0x3b, 0x09, 0x83, 0xe9, 0xcc, 0x3a, 0xb3, 0xb1, 0xf4, 0x2d, 0x7c, 0x20, 0x1f, 0xc3, 0x97, 0xf0,
	0xae, 0x6f, 0xe0, 0xd9, 0x39, 0xb3, 0xe9, 0x26, 0x9a, 0xae, 0xfe, 0x93, 0xec, 0xec, 0xf9, 0xdd,
	0x66, 0x98, 0x73, 0x16, 0x6e, 0xc7, 0x26, 0x55, 0x9a, 0x6f, 0x70, 0xd9, 0x11, 0x92, 0x6f, 0x24,
	0x5a, 0xb5, 0x78, 0xdc, 0xd3, 0xdc, 0x6c, 0xb4, 0x55, 0xd8, 0x94, 0x86, 0xeb, 0x74, 0x4b, 0xc8,
	0x48, 0x1f, 0x3d, 0x56, 0x7a, 0x57, 0x99, 0x54, 0xc8, 0x4e, 0xb8, 0xdb, 0xf3, 0x11, 0x97, 0x2a,
	0xb6, 0x4e, 0x64, 0x9f, 0xc8, 0xfe, 0x49, 0x8c, 0xfa, 0xa2, 0x33, 0x7a, 0x1b, 0x75, 0x7b, 0xdc,
	0x90, 0x40, 0x7d, 0x79, 0xd0, 0x9d, 0x6b, 0xad, 0xb4, 0x2b, 0x35, 0x06, 0x4b, 0x07, 0xdc, 0x98,
	0xa8, 0xc3, 0x5d, 0xf1, 0xfc, 0x70, 0x31, 0x8d, 0x84, 0x6c, 0x2b, 0x7d, 0x10, 0xa5, 0x42, 0x49,
	0x02, 0xad, 0x7d, 0xa8, 0x01, 0xec, 0x46, 0x3a, 0xc2, 0x2a, 0xd7, 0x86, 0xbd, 0x84, 0x95, 0x04,
	0xff, 0x95, 0x0c, 0x45, 0xcc, 0x65, 0x2a, 0xda, 0xa2, 0x65, 0xd1, 0x21, 0x25, 0xf2, 0xc6, 0xce,
	0x8e, 0x5d, 0x9e, 0xbe, 0x56, 0xf7, 0xdd, 0x9e, 0x5c, 0x4e, 0x93, 0x6a, 0xdc, 0xc0, 0x5e, 0xb6,
	0x08, 0xea, 0xc4, 0x6f, 0x0e, 0xd0, 0x6d, 0xc9, 0xb0, 0x27, 0x70, 0xee, 0x24, 0xf5, 0x50, 0xf6,
	0xba, 0x5d, 0xef, 0xe3, 0x24, 0x7a, 0xd4, 0x82, 0xd5, 0xd1, 0x3a, 0x3b, 0x08, 0x63, 0x0f, 0x60,
	0xce, 0x69, 0xa5, 0x47, 0x09, 0x47, 0x41, 0x6f, 0xdc, 0x66, 0x6b, 0x0c, 0x65, 0x13, 0x32, 0xe5,
	0x1d, 0xae, 0x29, 0xdc, 0x0c, 0x51, 0x9e, 0x21, 0xa3, 0x19, 0x33, 0x1f, 0x16, 0x07, 0x25, 0x28,
	0xc0, 0x27, 0x0a, 0xb0, 0x50, 0xc4, 0x5a, 0xcb, 0x9b, 0x30, 0xd5, 0x93, 0xe2, 0x4d, 0xcf, 0xba,
	0x55, 0x4a, 0x4f, 0xa2, 0x46, 0x60, 0x34, 0xba, 0x08, 0x73, 0x7d, 0x22, 0x79, 0x7c, 0x26, 0x8f,
	0x99, 0x1c, 0x62, 0xf5, 0x6f, 0x01, 0x24, 0xee, 0x2e, 0xa0, 0xc1, 0x7f, 0xe5, 0xdb, 0x99, 0x72,
	0x70, 0xb4, 0xb8, 0x04, 0xf3, 0xc7, 0x5c, 0xf2, 0xf8, 0x42, 0x1e, 0xb3, 0x7d, 0x90, 0x35, 0xb9,
	0x0b, 0x33, 0x2d, 0x85, 0x22, 0x32, 0xb5, 0xbb, 0xf6, 0xaa, 0xa5, 0xfb, 0x98, 0x76, 0xf8, 0xec,
	0x24, 0xd8, 0x15, 0xf8, 0xbf, 0x48, 0x27, 0xa7, 0xaf, 0xe4, 0x34, 0x5f, 0x00, 0x5a, 0xaf, 0x1b,
	0x50, 0x6b, 0x8b, 0x2e, 0x97, 0x78, 0xbb, 0xbc, 0x89, 0xf2, 0xf3, 0xca, 0xb1, 0xec, 0x02, 0xcc,
	0xe6, 0xcf, 0x64, 0xf0, 0xcd, 0x1d, 0x57, 0xfe, 0xd6, 0xaa, 0xdf, 0x81, 0xe9, 0x98, 0x9b, 0x96,
	0x16, 0x49, 0x76, 0x35, 0xbc, 0xc9, 0xf2, 0x8d, 0x14, 0xe0, 0x6c, 0x1d, 0x16, 0x0a, 0x4b, 0xb2,
	0xf9, 0xee, 0xf6, 0x51, 0x28, 0x58, 0xa7, 0x3d, 0x58, 0xce, 0x0f, 0x77, 0xdf, 0x76, 0x6d, 0xff,
	0xfe, 0x72, 0xed, 0xd5, 0x4a, 0x7d, 0x97, 0x1c, 0x99, 0x3a, 0xbe, 0xd9, 0xa7, 0xb2, 0x87, 0xb0,
	0x3a, 0x52, 0x97, 0x12, 0xfd, 0xa0, 0x44, 0x8d, 0x11, 0x0a, 0x36, 0xdd, 0x53, 0x58, 0x32, 0x3c,
	0xc1, 0x1e, 0x46, 0xff, 0x50, 0xb8, 0xde, 0xca, 0x5a, 0xca, 0x78, 0x53, 0xa5, 0xd9, 0x4e, 0xf5,
	0xa9, 0x4d, 0xea, 0x35, 0x7c, 0x6d, 0xd8, 0x7d, 0x58, 0x19, 0x21, 0x49, 0xb1, 0x7e, 0x52, 0x2c,
	0xef, 0x4f, 0xe4, 0x2c, 0xd3, 0xda, 0xfb, 0x71, 0xa8, 0x05, 0xdc, 0x24, 0x0a, 0x47, 0x1d, 0xbb,
	0x0a, 0x55, 0x3b, 0xb4, 0x86, 0xa7, 0x87, 0x9b, 0x88, 0x34, 0xd0, 0x1e, 0x65, 0xbf, 0x01, 0x01,
	0xd9, 0x73, 0x58, 0xc8, 0xc6, 0x55, 0x58, 0x98, 0x57, 0xd8, 0xde, 0x15, 0x24, 0xfb, 0x43, 0xe4,
	0xe1, 0xa9, 0xb6, 0x8d, 0xeb, 0xe6, 0xf1, 0x3a, 0x98, 0x3f, 0x18, 0x7c, 0xc1, 0x36, 0x61, 0xd2,
	0x8d, 0x49, 0x6c, 0xe1, 0x4c, 0x71, 0xf5, 0x37, 0x45, 0x1a, 0xa2, 0xdb, 0xf4, 0x1f, 0xe4, 0x70,
	0x9c, 0x5e, 0x15, 0xad, 0x0e, 0xb1, 0x2f, 0x33, 0xd6, 0xa6, 0xff, 0xf7, 0x63, 0xdd, 0xcf, 0x4f,
	0xc2, 0x0f, 0xd4, 0x61, 0x90, 0x89, 0xd4, 0xcf, 0x40, 0x05, 0x9f, 0xd9, 0x69, 0x98, 0xc0, 0x55,
	0xd6, 0xed, 0xef, 0x76, 0xf0, 0x6c, 0xaa, 0x41, 0x15, 0x97, 0xcd, 0x78, 0xeb, 0x15, 0x34, 0x84,
	0x1a, 0x72, 0x38, 0xfe, 0xea, 0xbc, 0xb8, 0xd7, 0x51, 0x26, 0x7e, 0x9d, 0xd7, 0xe3, 0x7f, 0xfd,
	0x30, 0xed, 0x4f, 0xd8, 0xd9, 0x7f, 0xfd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x7e, 0x14,
	0x19, 0xd8, 0x06, 0x00, 0x00,
}
