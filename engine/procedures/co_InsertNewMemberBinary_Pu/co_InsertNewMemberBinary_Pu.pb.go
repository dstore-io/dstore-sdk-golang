// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_InsertNewMemberBinary_Pu.proto
// DO NOT EDIT!

/*
Package co_InsertNewMemberBinary_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_InsertNewMemberBinary_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_InsertNewMemberBinary_Pu

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
	UniqueId                       *dstore_values.StringValue    `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                          `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues     *dstore_values.StringValue    `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                          `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                    *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                bool                          `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityMemberId              *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull          bool                          `protobuf:"varint,1004,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
	CommunityBinaryCategoryId      *dstore_values.IntegerValue   `protobuf:"bytes,5,opt,name=community_binary_category_id,json=communityBinaryCategoryId" json:"community_binary_category_id,omitempty"`
	CommunityBinaryCategoryIdNull  bool                          `protobuf:"varint,1005,opt,name=community_binary_category_id_null,json=communityBinaryCategoryIdNull" json:"community_binary_category_id_null,omitempty"`
	BinaryDescription              *dstore_values.StringValue    `protobuf:"bytes,6,opt,name=binary_description,json=binaryDescription" json:"binary_description,omitempty"`
	BinaryDescriptionNull          bool                          `protobuf:"varint,1006,opt,name=binary_description_null,json=binaryDescriptionNull" json:"binary_description_null,omitempty"`
	MD5Hash                        *dstore_values.StringValue    `protobuf:"bytes,7,opt,name=m_d5_hash,json=mD5Hash" json:"m_d5_hash,omitempty"`
	MD5HashNull                    bool                          `protobuf:"varint,1007,opt,name=m_d5_hash_null,json=mD5HashNull" json:"m_d5_hash_null,omitempty"`
	FileDateAndTime                *dstore_values.TimestampValue `protobuf:"bytes,8,opt,name=file_date_and_time,json=fileDateAndTime" json:"file_date_and_time,omitempty"`
	FileDateAndTimeNull            bool                          `protobuf:"varint,1008,opt,name=file_date_and_time_null,json=fileDateAndTimeNull" json:"file_date_and_time_null,omitempty"`
	FilesizeInKB                   *dstore_values.IntegerValue   `protobuf:"bytes,9,opt,name=filesize_in_k_b,json=filesizeInKB" json:"filesize_in_k_b,omitempty"`
	FilesizeInKBNull               bool                          `protobuf:"varint,1009,opt,name=filesize_in_k_b_null,json=filesizeInKBNull" json:"filesize_in_k_b_null,omitempty"`
	ContentType                    *dstore_values.StringValue    `protobuf:"bytes,10,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	ContentTypeNull                bool                          `protobuf:"varint,1010,opt,name=content_type_null,json=contentTypeNull" json:"content_type_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue    `protobuf:"bytes,11,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                          `protobuf:"varint,1011,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Parameters) GetCommunityBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityBinaryCategoryId
	}
	return nil
}

func (m *Parameters) GetBinaryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryDescription
	}
	return nil
}

func (m *Parameters) GetMD5Hash() *dstore_values.StringValue {
	if m != nil {
		return m.MD5Hash
	}
	return nil
}

func (m *Parameters) GetFileDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.FileDateAndTime
	}
	return nil
}

func (m *Parameters) GetFilesizeInKB() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilesizeInKB
	}
	return nil
}

func (m *Parameters) GetContentType() *dstore_values.StringValue {
	if m != nil {
		return m.ContentType
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_InsertNewMemberBinary_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_InsertNewMemberBinary_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_InsertNewMemberBinary_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_InsertNewMemberBinary_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xfb, 0x4e, 0x13, 0x4d,
	0x14, 0x0f, 0xf4, 0x2b, 0x2d, 0x53, 0xc2, 0x65, 0xe1, 0xfb, 0x58, 0xca, 0x25, 0x1f, 0xa8, 0x89,
	0xc6, 0x64, 0x6b, 0x34, 0x88, 0x31, 0x41, 0x63, 0xc5, 0xc4, 0x85, 0x40, 0x70, 0x43, 0x4c, 0x34,
	0x24, 0x9b, 0x6d, 0x77, 0x28, 0x13, 0xba, 0x33, 0x75, 0x66, 0x2b, 0xc1, 0x77, 0x30, 0xf1, 0x7d,
	0x7c, 0x22, 0xef, 0xb7, 0x17, 0xf0, 0xcc, 0x9c, 0xed, 0x6d, 0x0b, 0xac, 0xfe, 0xd3, 0x76, 0xf6,
	0xfc, 0x6e, 0x9d, 0x9d, 0x39, 0x87, 0xdc, 0x0f, 0x55, 0x2c, 0x24, 0xad, 0x50, 0xde, 0x60, 0x9c,
	0x56, 0x5a, 0x52, 0xd4, 0x69, 0xd8, 0x96, 0x54, 0x55, 0xea, 0xc2, 0x77, 0xb9, 0xa2, 0x32, 0xde,
	0xa3, 0xa7, 0xbb, 0x34, 0xaa, 0x51, 0x59, 0x65, 0x3c, 0x90, 0x67, 0xfe, 0x7e, 0xdb, 0x01, 0x58,
	0x2c, 0xac, 0x1b, 0xc8, 0x75, 0x90, 0xeb, 0x5c, 0x42, 0x28, 0xcf, 0x26, 0x36, 0xaf, 0x83, 0x66,
	0x9b, 0x2a, 0xe4, 0x97, 0x17, 0x06, 0xbd, 0xa9, 0x94, 0x42, 0x26, 0xa5, 0xc5, 0xc1, 0x52, 0x44,
	0x95, 0x0a, 0x1a, 0x34, 0x29, 0x5e, 0x49, 0x17, 0xe3, 0x80, 0xf1, 0x23, 0x21, 0xa3, 0x20, 0x66,
	0x82, 0x23, 0x68, 0xed, 0x6d, 0x89, 0x90, 0xfd, 0x40, 0x06, 0x50, 0xa5, 0x52, 0x59, 0x1b, 0x64,
	0xbc, 0xcd, 0xd9, 0xab, 0x36, 0xf5, 0x59, 0x68, 0x8f, 0xfc, 0x3f, 0x72, 0xbd, 0x74, 0xbb, 0xec,
	0x24, 0xf9, 0x93, 0x50, 0x2a, 0x96, 0x8c, 0x37, 0x9e, 0xeb, 0x85, 0x57, 0x44, 0xb0, 0x1b, 0x5a,
	0xd7, 0xc8, 0x64, 0x97, 0xe8, 0xf3, 0x76, 0xb3, 0x69, 0x7f, 0x28, 0x00, 0xbd, 0xe8, 0x4d, 0x74,
	0x20, 0x7b, 0xf0, 0xd0, 0x3a, 0x24, 0x4b, 0x2d, 0xf0, 0x11, 0x1c, 0x60, 0x94, 0xc7, 0xec, 0x88,
	0xd5, 0x4d, 0x1a, 0x1f, 0xc5, 0xed, 0xd1, 0x4c, 0xcb, 0x32, 0xf2, 0xdd, 0x01, 0xba, 0x29, 0x29,
	0x6b, 0x9b, 0xac, 0x5e, 0xa6, 0x8e, 0xb9, 0x3e, 0x62, 0xae, 0x95, 0x8b, 0x75, 0x4c, 0xd2, 0x07,
	0x64, 0xa2, 0x2e, 0xa2, 0x08, 0xd2, 0xc7, 0x67, 0x7a, 0x33, 0x72, 0x26, 0xd9, 0x62, 0x2a, 0x19,
	0xe3, 0x31, 0x6d, 0x50, 0x89, 0xd1, 0x4a, 0x5d, 0x02, 0x6c, 0xc8, 0x4d, 0x32, 0xd3, 0xcf, 0x47,
	0xef, 0x4f, 0xe8, 0x3d, 0xd5, 0x07, 0x34, 0x66, 0x3b, 0x64, 0xb6, 0x07, 0x8e, 0xcc, 0xa1, 0xd0,
	0x9e, 0xff, 0x64, 0x7b, 0xf6, 0x4c, 0xf0, 0x2c, 0x81, 0xf3, 0x3d, 0x62, 0x9f, 0x23, 0x86, 0x01,
	0x3e, 0x63, 0x80, 0x7f, 0x87, 0x58, 0x9d, 0xb7, 0xd3, 0x63, 0xd6, 0xf0, 0x54, 0xc2, 0xc6, 0xd0,
	0x86, 0x90, 0x66, 0x0f, 0xf2, 0xd9, 0x79, 0x16, 0xba, 0x02, 0x78, 0xaa, 0x1f, 0x27, 0x74, 0xc8,
	0xe5, 0x92, 0xd5, 0xcb, 0xd4, 0x31, 0xe0, 0x17, 0x0c, 0xb8, 0x7c, 0xa1, 0x8c, 0x09, 0xea, 0x12,
	0x2b, 0x11, 0x08, 0xa9, 0xaa, 0x4b, 0xd6, 0xd2, 0x2f, 0xcf, 0x1e, 0xcb, 0x3c, 0x3c, 0x33, 0xc8,
	0xda, 0xea, 0x91, 0xe0, 0xc4, 0xcf, 0x0f, 0x4b, 0x61, 0x96, 0xaf, 0xc9, 0x66, 0x0d, 0x91, 0x4c,
	0x86, 0xbb, 0x64, 0x3c, 0xf2, 0xc3, 0x75, 0xff, 0x38, 0x50, 0xc7, 0x76, 0x21, 0xd3, 0xba, 0x10,
	0x6d, 0xad, 0x3f, 0x05, 0xa8, 0x75, 0x95, 0x4c, 0x76, 0x79, 0xe8, 0xf3, 0x0d, 0x7d, 0x4a, 0x09,
	0xc2, 0xa8, 0x6f, 0x13, 0xeb, 0x88, 0x35, 0xa9, 0x1f, 0xc2, 0x1f, 0xf7, 0x03, 0x1e, 0xfa, 0x31,
	0x8b, 0xa8, 0x5d, 0x34, 0x36, 0xcb, 0x29, 0x1b, 0x5d, 0x52, 0x71, 0x10, 0xb5, 0xd0, 0x69, 0x4a,
	0x13, 0xb7, 0x80, 0xf7, 0x88, 0x87, 0x07, 0x50, 0xb2, 0xd6, 0xc9, 0xfc, 0xb0, 0x16, 0x5a, 0x7f,
	0x47, 0xeb, 0xd9, 0x14, 0xc5, 0x44, 0xa8, 0x12, 0xa3, 0xa4, 0xd8, 0x1b, 0xb8, 0xd4, 0xdc, 0x3f,
	0xf1, 0x6b, 0xf6, 0x78, 0xf6, 0x01, 0x98, 0xe8, 0x70, 0x5c, 0xbe, 0x53, 0xb5, 0x2a, 0x64, 0x2e,
	0xa5, 0x81, 0xbe, 0x3f, 0xd0, 0x77, 0xba, 0x1f, 0x6c, 0x4c, 0x37, 0xf5, 0xb5, 0x03, 0x3d, 0x1e,
	0xfb, 0xf1, 0x59, 0x8b, 0xda, 0x24, 0x73, 0x63, 0x4b, 0x09, 0xfe, 0x00, 0xe0, 0x78, 0xeb, 0x7a,
	0x74, 0x34, 0xfb, 0xd9, 0xbd, 0x75, 0x5d, 0xa0, 0xf1, 0x7a, 0x46, 0xe6, 0x15, 0x6d, 0x41, 0xf3,
	0x03, 0x69, 0x9d, 0xce, 0x34, 0x0d, 0xdd, 0x2b, 0x94, 0x5d, 0xca, 0xb4, 0x9d, 0xeb, 0x52, 0x5d,
	0x6c, 0x22, 0xf0, 0x58, 0x59, 0x0f, 0xc9, 0xd2, 0x05, 0x92, 0x18, 0xe5, 0x17, 0x46, 0xb1, 0xcf,
	0x23, 0xeb, 0x4c, 0x6b, 0xef, 0x47, 0x49, 0xd1, 0xa3, 0xaa, 0x25, 0x60, 0x44, 0x58, 0xb7, 0x48,
	0xde, 0x74, 0xfb, 0x74, 0x27, 0x4e, 0x26, 0x09, 0x4e, 0x82, 0x27, 0xfa, 0xd3, 0x43, 0xa0, 0xf5,
	0x82, 0x4c, 0xeb, 0x3e, 0xef, 0xf7, 0x35, 0x7a, 0xe8, 0xa9, 0x39, 0x20, 0x3b, 0x29, 0x72, 0x7a,
	0x1c, 0xec, 0xc2, 0xda, 0xed, 0xad, 0xbd, 0xa9, 0x68, 0xf0, 0x01, 0xb4, 0x95, 0x42, 0x32, 0x5f,
	0xa0, 0x17, 0x6a, 0xc5, 0x95, 0x21, 0x45, 0x9c, 0x3e, 0xbb, 0xf8, 0xed, 0x75, 0xe0, 0x70, 0x5b,
	0x73, 0x52, 0x9c, 0x42, 0x37, 0xd3, 0xac, 0x0d, 0xe7, 0x8f, 0xc7, 0xa1, 0xd3, 0xd9, 0x08, 0xc7,
	0x13, 0xa7, 0x9e, 0xd6, 0x28, 0x2f, 0x93, 0x1c, 0xfc, 0xb6, 0xfe, 0x23, 0x63, 0xb0, 0xd2, 0x2d,
	0xe9, 0xdd, 0x1e, 0x6c, 0x4d, 0xde, 0xcb, 0xc3, 0xd2, 0x0d, 0xab, 0x87, 0x64, 0x91, 0x89, 0x94,
	0x41, 0x6f, 0x56, 0xbf, 0xdc, 0x6c, 0x08, 0x15, 0x9e, 0x74, 0xea, 0xe1, 0x5f, 0x8e, 0xf3, 0xda,
	0x98, 0x19, 0x99, 0x77, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x21, 0xd8, 0xbe, 0x0d, 0x08,
	0x00, 0x00,
}
