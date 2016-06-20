// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetMD5ForBinaryID_Pu.proto
// DO NOT EDIT!

/*
Package co_GetMD5ForBinaryID_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetMD5ForBinaryID_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetMD5ForBinaryID_Pu

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
	UniqueId                        *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                    bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues      *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull  bool                        `protobuf:"varint,1002,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                 bool                        `protobuf:"varint,1003,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunityMemberId               *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull           bool                        `protobuf:"varint,1004,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
	BinaryId                        *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=binary_id,json=binaryId" json:"binary_id,omitempty"`
	BinaryIdNull                    bool                        `protobuf:"varint,1005,opt,name=binary_id_null,json=binaryIdNull" json:"binary_id_null,omitempty"`
	AdditionalBinaryInformation     *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=additional_binary_information,json=additionalBinaryInformation" json:"additional_binary_information,omitempty"`
	AdditionalBinaryInformationNull bool                        `protobuf:"varint,1006,opt,name=additional_binary_information_null,json=additionalBinaryInformationNull" json:"additional_binary_information_null,omitempty"`
	SeparatorInIdentVals            *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull        bool                        `protobuf:"varint,1007,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
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

func (m *Parameters) GetBinaryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryId
	}
	return nil
}

func (m *Parameters) GetAdditionalBinaryInformation() *dstore_values.BooleanValue {
	if m != nil {
		return m.AdditionalBinaryInformation
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
	RowId           int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	MD5Hash         *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=m_d5_hash,json=mD5Hash" json:"m_d5_hash,omitempty"`
	ContentType     *dstore_values.StringValue    `protobuf:"bytes,20001,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	FileDateAndTime *dstore_values.TimestampValue `protobuf:"bytes,20002,opt,name=file_date_and_time,json=fileDateAndTime" json:"file_date_and_time,omitempty"`
	FilesizeInKB    *dstore_values.IntegerValue   `protobuf:"bytes,20004,opt,name=filesize_in_k_b,json=filesizeInKB" json:"filesize_in_k_b,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetMD5Hash() *dstore_values.StringValue {
	if m != nil {
		return m.MD5Hash
	}
	return nil
}

func (m *Response_Row) GetContentType() *dstore_values.StringValue {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Response_Row) GetFileDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.FileDateAndTime
	}
	return nil
}

func (m *Response_Row) GetFilesizeInKB() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilesizeInKB
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetMD5ForBinaryID_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetMD5ForBinaryID_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetMD5ForBinaryID_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/co_GetMD5ForBinaryID_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdf, 0x4e, 0x13, 0x4d,
	0x14, 0x0f, 0x94, 0xd2, 0x32, 0x90, 0x8f, 0xef, 0x5b, 0x3e, 0x75, 0x6d, 0x01, 0x11, 0x43, 0x62,
	0x62, 0xb2, 0x18, 0x09, 0x8a, 0x37, 0x1a, 0x49, 0x51, 0x2b, 0x96, 0xe0, 0x86, 0x18, 0x35, 0x26,
	0x9b, 0x69, 0x77, 0x28, 0x13, 0xba, 0x33, 0x75, 0x66, 0x2b, 0xc1, 0xa7, 0xd0, 0x7b, 0x6f, 0x34,
	0xde, 0x99, 0xf8, 0x18, 0xbe, 0x8b, 0xff, 0x5f, 0xc1, 0x33, 0x73, 0x76, 0x5b, 0xba, 0x40, 0xab,
	0x37, 0x6d, 0x66, 0xce, 0xf9, 0xfd, 0xc9, 0x39, 0x67, 0xce, 0x92, 0xeb, 0xa1, 0x8e, 0xa5, 0x62,
	0xcb, 0x4c, 0x34, 0xb9, 0x60, 0xcb, 0x6d, 0x25, 0x1b, 0x2c, 0xec, 0x28, 0xa6, 0x97, 0x1b, 0x32,
	0xb8, 0xc7, 0xe2, 0x5a, 0x65, 0xf5, 0xae, 0x54, 0xeb, 0x5c, 0x50, 0x75, 0x58, 0xad, 0x04, 0xdb,
	0x1d, 0x0f, 0x52, 0x62, 0xe9, 0x2c, 0x21, 0xce, 0x43, 0x9c, 0x77, 0x4a, 0x72, 0x69, 0x26, 0xa1,
	0x7f, 0x49, 0x5b, 0x1d, 0xa6, 0x11, 0x5b, 0x3a, 0xdf, 0xaf, 0xc9, 0x94, 0x92, 0x2a, 0x09, 0x95,
	0xfb, 0x43, 0x11, 0xd3, 0x9a, 0x36, 0x59, 0x12, 0xbc, 0x94, 0x0d, 0xc6, 0x94, 0x8b, 0x5d, 0xa9,
	0x22, 0x1a, 0x73, 0x29, 0x30, 0x69, 0xf1, 0x53, 0x81, 0x90, 0x6d, 0xaa, 0x28, 0x44, 0x99, 0xd2,
	0xce, 0x0d, 0x32, 0xd1, 0x11, 0xfc, 0x45, 0x87, 0x05, 0x3c, 0x74, 0x47, 0x16, 0x46, 0x2e, 0x4f,
	0x5e, 0x2b, 0x79, 0x89, 0xf7, 0xc4, 0x94, 0x8e, 0x15, 0x17, 0xcd, 0xc7, 0xe6, 0xe0, 0x17, 0x31,
	0xb9, 0x1a, 0x3a, 0x4b, 0xe4, 0x9f, 0x2e, 0x30, 0x10, 0x9d, 0x56, 0xcb, 0xfd, 0x52, 0x00, 0x78,
	0xd1, 0x9f, 0x4a, 0x53, 0xb6, 0xe0, 0xd2, 0x79, 0x4e, 0x66, 0xdb, 0xa0, 0x23, 0x05, 0xa4, 0x31,
	0x11, 0xf3, 0x5d, 0xde, 0xb0, 0x6e, 0x02, 0x24, 0x77, 0x47, 0x87, 0x4a, 0x96, 0x10, 0x5f, 0xed,
	0x83, 0xdb, 0x90, 0x76, 0x1e, 0x90, 0x8b, 0x83, 0xd8, 0xd1, 0xd7, 0x57, 0xf4, 0x35, 0x7f, 0x3a,
	0x8f, 0x75, 0x7a, 0x8b, 0x4c, 0x35, 0x64, 0x14, 0x81, 0xfb, 0xf8, 0xd0, 0x14, 0x23, 0x67, 0x9d,
	0x95, 0x33, 0xce, 0xb8, 0x88, 0x59, 0x93, 0x29, 0xb4, 0x36, 0xd9, 0x05, 0x40, 0x41, 0xae, 0x90,
	0xff, 0x8e, 0xe2, 0x51, 0xfb, 0x1b, 0x6a, 0x4f, 0x1f, 0x49, 0xb4, 0x62, 0x9b, 0x64, 0xa6, 0x97,
	0x1c, 0xb1, 0xa8, 0xce, 0x94, 0xd1, 0x1c, 0x1b, 0xae, 0xd9, 0x13, 0xa9, 0x59, 0x18, 0x28, 0xaf,
	0x11, 0xf7, 0x04, 0x32, 0x34, 0xf0, 0x1d, 0x0d, 0x9c, 0x39, 0x86, 0xb2, 0x36, 0xd6, 0xc8, 0x44,
	0xdd, 0x4e, 0xa3, 0x11, 0xcf, 0x0f, 0x17, 0x2f, 0x62, 0x36, 0xb6, 0xbf, 0x8b, 0x44, 0xa5, 0x1f,
	0x49, 0xfb, 0xd3, 0x14, 0x2b, 0x10, 0x90, 0x39, 0x1a, 0x86, 0xdc, 0x94, 0x9a, 0xb6, 0x82, 0x14,
	0xd1, 0x1b, 0x4a, 0x77, 0xfc, 0x44, 0xd1, 0xba, 0x94, 0x2d, 0x46, 0xb1, 0x3b, 0x7e, 0xb9, 0xc7,
	0x90, 0x3c, 0x9d, 0x1e, 0xde, 0x79, 0x48, 0x16, 0x07, 0x0a, 0xa0, 0xb7, 0x9f, 0xe8, 0xed, 0xc2,
	0x00, 0x26, 0x6b, 0xf7, 0x11, 0x39, 0xa7, 0x59, 0x1b, 0x5e, 0x07, 0x98, 0x01, 0x16, 0x9c, 0x2a,
	0x33, 0x4c, 0xda, 0x2d, 0x0c, 0x1d, 0xd4, 0xff, 0xbb, 0xd0, 0x2a, 0x4e, 0x19, 0x5c, 0x6b, 0xe7,
	0x36, 0x99, 0x3d, 0x85, 0x12, 0xad, 0xfd, 0x42, 0x6b, 0xee, 0x49, 0x60, 0xe3, 0x69, 0xf1, 0xf3,
	0x18, 0x29, 0xfa, 0x4c, 0xb7, 0xa5, 0xd0, 0xcc, 0xb9, 0x4a, 0xf2, 0x76, 0x1d, 0x64, 0x9f, 0x6a,
	0xb2, 0x66, 0x70, 0x55, 0x6c, 0x98, 0x5f, 0x1f, 0x13, 0x9d, 0xa7, 0xe4, 0x5f, 0xb3, 0x08, 0xfa,
	0x8a, 0x3e, 0xba, 0x90, 0x03, 0xb0, 0x97, 0x01, 0x67, 0xf7, 0x45, 0x0d, 0xce, 0x47, 0x0a, 0xe4,
	0x4f, 0x47, 0xfd, 0x17, 0x30, 0x3d, 0x85, 0x64, 0x01, 0xc1, 0x63, 0x31, 0x8c, 0xf3, 0xc7, 0x18,
	0x71, 0x3d, 0xd5, 0xf0, 0xdf, 0x4f, 0xd3, 0x9d, 0x0d, 0x92, 0x53, 0xf2, 0x00, 0xc6, 0xdd, 0xa0,
	0x56, 0xbc, 0x3f, 0xda, 0x95, 0x5e, 0x5a, 0x04, 0xcf, 0x97, 0x07, 0xbe, 0xc1, 0x97, 0x3e, 0x8e,
	0x92, 0x1c, 0x1c, 0x9c, 0xb3, 0x64, 0x1c, 0x8e, 0x66, 0x86, 0x5f, 0x6f, 0x41, 0x5d, 0xf2, 0x7e,
	0x1e, 0x8e, 0x30, 0xa4, 0xb0, 0xdc, 0xa2, 0x20, 0x5c, 0x0d, 0xf6, 0xa8, 0xde, 0x73, 0xdf, 0x6c,
	0x0d, 0xed, 0x60, 0x21, 0xaa, 0xac, 0xde, 0x87, 0x5c, 0x68, 0x1a, 0xec, 0x02, 0x18, 0x7c, 0xe8,
	0x54, 0x7c, 0xd8, 0x66, 0xee, 0xbb, 0xb7, 0xc3, 0x57, 0xe3, 0x64, 0x82, 0xd8, 0x01, 0x00, 0x8c,
	0xa5, 0xb3, 0xcb, 0x5b, 0x2c, 0x08, 0x69, 0xcc, 0x02, 0x2a, 0xc2, 0x20, 0xe6, 0x11, 0x73, 0xdf,
	0x27, 0x34, 0x73, 0x19, 0x1a, 0x13, 0xd4, 0x31, 0x8d, 0xda, 0xc8, 0x34, 0x6d, 0xa0, 0x15, 0x40,
	0xde, 0x11, 0xe1, 0x0e, 0x84, 0x9c, 0x0a, 0xb1, 0x57, 0x9a, 0xbf, 0x62, 0x66, 0x84, 0xf6, 0x83,
	0xba, 0xfb, 0x21, 0xa1, 0x1a, 0xf8, 0x5c, 0xa7, 0x52, 0x54, 0x55, 0x6c, 0xae, 0xaf, 0x3f, 0x21,
	0x65, 0x2e, 0x33, 0xb5, 0xee, 0x7d, 0xcf, 0x9e, 0xdd, 0x6c, 0x4a, 0x1d, 0xee, 0xa7, 0xf1, 0xf0,
	0x2f, 0x3e, 0x79, 0xf5, 0x71, 0xfb, 0x69, 0x59, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x32, 0x04,
	0x94, 0xa2, 0x2d, 0x07, 0x00, 0x00,
}