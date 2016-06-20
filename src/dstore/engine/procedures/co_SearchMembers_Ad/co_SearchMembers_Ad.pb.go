// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_SearchMembers_Ad.proto
// DO NOT EDIT!

/*
Package co_SearchMembers_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_SearchMembers_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_SearchMembers_Ad

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
	CommunityId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull               bool                        `protobuf:"varint,1001,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	SearchString                  *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=search_string,json=searchString" json:"search_string,omitempty"`
	SearchStringNull              bool                        `protobuf:"varint,1002,opt,name=search_string_null,json=searchStringNull" json:"search_string_null,omitempty"`
	MaxNumberOfRows               *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=max_number_of_rows,json=maxNumberOfRows" json:"max_number_of_rows,omitempty"`
	MaxNumberOfRowsNull           bool                        `protobuf:"varint,1003,opt,name=max_number_of_rows_null,json=maxNumberOfRowsNull" json:"max_number_of_rows_null,omitempty"`
	CaseSensitive                 *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull             bool                        `protobuf:"varint,1004,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	OutputCharacteristicId1       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=output_characteristic_id1,json=outputCharacteristicId1" json:"output_characteristic_id1,omitempty"`
	OutputCharacteristicId1Null   bool                        `protobuf:"varint,1005,opt,name=output_characteristic_id1_null,json=outputCharacteristicId1Null" json:"output_characteristic_id1_null,omitempty"`
	OutputCharacteristicId2       *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=output_characteristic_id2,json=outputCharacteristicId2" json:"output_characteristic_id2,omitempty"`
	OutputCharacteristicId2Null   bool                        `protobuf:"varint,1006,opt,name=output_characteristic_id2_null,json=outputCharacteristicId2Null" json:"output_characteristic_id2_null,omitempty"`
	OutputCharacteristicId3       *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=output_characteristic_id3,json=outputCharacteristicId3" json:"output_characteristic_id3,omitempty"`
	OutputCharacteristicId3Null   bool                        `protobuf:"varint,1007,opt,name=output_characteristic_id3_null,json=outputCharacteristicId3Null" json:"output_characteristic_id3_null,omitempty"`
	CommunityBinaryCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=community_binary_category_id,json=communityBinaryCategoryId" json:"community_binary_category_id,omitempty"`
	CommunityBinaryCategoryIdNull bool                        `protobuf:"varint,1008,opt,name=community_binary_category_id_null,json=communityBinaryCategoryIdNull" json:"community_binary_category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetSearchString() *dstore_values.StringValue {
	if m != nil {
		return m.SearchString
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfRows() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfRows
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId1() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId1
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId2() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId2
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId3() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId3
	}
	return nil
}

func (m *Parameters) GetCommunityBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityBinaryCategoryId
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
	Value1RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value1_restricted_by_pattern,json=value1RestrictedByPattern" json:"value1_restricted_by_pattern,omitempty"`
	BinaryId                  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=binary_id,json=binaryId" json:"binary_id,omitempty"`
	Value2RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value2_restricted_by_pattern,json=value2RestrictedByPattern" json:"value2_restricted_by_pattern,omitempty"`
	CommunityMemberId         *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	Value3                    *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=value3" json:"value3,omitempty"`
	Value1                    *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=value1" json:"value1,omitempty"`
	Value2                    *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=value2" json:"value2,omitempty"`
	Value3RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=value3_restricted_by_pattern,json=value3RestrictedByPattern" json:"value3_restricted_by_pattern,omitempty"`
	Nickname                  *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=nickname" json:"nickname,omitempty"`
	IsOnline                  *dstore_values.BooleanValue `protobuf:"bytes,10010,opt,name=is_online,json=isOnline" json:"is_online,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue1RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value1RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetBinaryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryId
	}
	return nil
}

func (m *Response_Row) GetValue2RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value2RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetValue3RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value3RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetNickname() *dstore_values.StringValue {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *Response_Row) GetIsOnline() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsOnline
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_SearchMembers_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_SearchMembers_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_SearchMembers_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xeb, 0x6e, 0xd3, 0x48,
	0x14, 0x56, 0x37, 0x9b, 0xcb, 0x4e, 0xdb, 0x6d, 0x3b, 0x95, 0xb6, 0x4e, 0xb2, 0x5b, 0x6d, 0xbb,
	0x7f, 0x56, 0x02, 0x1c, 0x62, 0x0b, 0x01, 0x7f, 0x40, 0x4d, 0x41, 0x22, 0x82, 0xa6, 0x95, 0x2b,
	0x81, 0x40, 0x48, 0x96, 0xe3, 0x4c, 0xd3, 0x11, 0xc9, 0x4c, 0x34, 0xe3, 0x14, 0xf2, 0x16, 0xdc,
	0x6f, 0xaf, 0xc5, 0x33, 0xf0, 0x83, 0x72, 0x7d, 0x04, 0x8e, 0x67, 0x5c, 0xe7, 0xd2, 0x24, 0x0e,
	0xfd, 0x93, 0x68, 0x3c, 0xe7, 0x3b, 0xdf, 0xe7, 0x73, 0xe6, 0x7c, 0x63, 0x64, 0x35, 0x64, 0xc0,
	0x05, 0x29, 0x11, 0xd6, 0xa4, 0x8c, 0x94, 0x3a, 0x82, 0xfb, 0xa4, 0xd1, 0x15, 0x44, 0x96, 0x7c,
	0xee, 0xee, 0x13, 0x4f, 0xf8, 0x87, 0x3b, 0xa4, 0x5d, 0x27, 0x42, 0xba, 0x5b, 0x0d, 0x13, 0xb6,
	0x03, 0x8e, 0x37, 0x34, 0xc6, 0xd4, 0x18, 0x73, 0x4c, 0x60, 0x61, 0x35, 0x4a, 0x7b, 0xe4, 0xb5,
	0xba, 0x44, 0x6a, 0x5c, 0x21, 0x3f, 0xcc, 0x45, 0x84, 0xe0, 0x22, 0xda, 0x2a, 0x0e, 0x6f, 0xb5,
	0x89, 0x94, 0x5e, 0x93, 0x44, 0x9b, 0xff, 0x8d, 0x6e, 0x06, 0x1e, 0x65, 0x07, 0x5c, 0xb4, 0xbd,
	0x80, 0x72, 0xa6, 0x83, 0x36, 0x3f, 0xe6, 0x10, 0xda, 0xf3, 0x84, 0x07, 0xbb, 0xa0, 0x01, 0x5f,
	0x43, 0x0b, 0x3e, 0x6f, 0xb7, 0xbb, 0x8c, 0x06, 0x3d, 0x97, 0x36, 0x8c, 0xb9, 0x7f, 0xe7, 0xfe,
	0x9f, 0xb7, 0x8a, 0x66, 0x24, 0x3d, 0xd2, 0x45, 0x59, 0x40, 0x9a, 0x44, 0xdc, 0x0d, 0x57, 0xce,
	0x7c, 0x0c, 0xa8, 0x36, 0xf0, 0x39, 0xb4, 0x32, 0x88, 0x77, 0x59, 0xb7, 0xd5, 0x32, 0x3e, 0x65,
	0x21, 0x4b, 0xce, 0x59, 0x1a, 0x08, 0xac, 0xc1, 0x73, 0x7c, 0x1d, 0x2d, 0x4a, 0x55, 0x01, 0x57,
	0x06, 0x82, 0xb2, 0xa6, 0xf1, 0x9b, 0x62, 0x2b, 0x8c, 0xb0, 0xe9, 0x4d, 0x4d, 0xb6, 0xa0, 0x01,
	0xfb, 0xea, 0x11, 0xbe, 0x80, 0xf0, 0x50, 0x02, 0x4d, 0x77, 0xac, 0xe9, 0x96, 0x07, 0x43, 0x15,
	0xdf, 0x2d, 0x84, 0xdb, 0xde, 0x13, 0x08, 0x0a, 0xeb, 0xed, 0xf2, 0x03, 0x57, 0xf0, 0xc7, 0xd2,
	0x48, 0x25, 0xbf, 0xe2, 0x12, 0xc0, 0x6a, 0x0a, 0xb5, 0x7b, 0xe0, 0x00, 0x06, 0x5f, 0x42, 0x6b,
	0xa7, 0x33, 0x69, 0xf6, 0xcf, 0x9a, 0x7d, 0x75, 0x04, 0xa2, 0x04, 0x54, 0xd0, 0x9f, 0xbe, 0x27,
	0x89, 0x2b, 0x09, 0x93, 0x34, 0xa0, 0x47, 0xc4, 0xf8, 0x7d, 0x2c, 0x79, 0x9d, 0xf3, 0x16, 0xf1,
	0x98, 0x26, 0x5f, 0x0c, 0x21, 0xfb, 0x27, 0x08, 0x5c, 0x42, 0xab, 0xc3, 0x39, 0x34, 0xed, 0x17,
	0x4d, 0xbb, 0x32, 0x14, 0xac, 0x48, 0xef, 0xa1, 0x3c, 0xef, 0x06, 0x9d, 0x6e, 0xe0, 0xfa, 0x87,
	0xd0, 0x68, 0x1f, 0xfa, 0x4c, 0x65, 0x40, 0x7d, 0x68, 0x4f, 0xd9, 0x48, 0x27, 0xbf, 0xfc, 0x9a,
	0x46, 0x6f, 0x0f, 0x81, 0xab, 0x8d, 0x32, 0xbe, 0x81, 0xd6, 0x27, 0x26, 0xd6, 0xa2, 0xbe, 0x6a,
	0x51, 0xc5, 0x09, 0x19, 0x92, 0xe4, 0x59, 0x46, 0xe6, 0xcc, 0xf2, 0xac, 0x69, 0xf2, 0x2c, 0x2d,
	0xef, 0xdb, 0x54, 0x79, 0x56, 0x92, 0x3c, 0xdb, 0xc8, 0x9e, 0x59, 0x9e, 0x3d, 0x4d, 0x9e, 0xad,
	0xe5, 0x7d, 0x9f, 0x2a, 0xcf, 0x56, 0xf2, 0x1e, 0xa2, 0xbf, 0xfb, 0xf3, 0x56, 0xa7, 0xcc, 0x13,
	0x3d, 0xd7, 0xf7, 0x80, 0x9f, 0x0b, 0x35, 0xbf, 0xb9, 0x64, 0x85, 0xf9, 0x38, 0x41, 0x45, 0xe1,
	0xb7, 0x23, 0x38, 0x4c, 0x73, 0x15, 0x6d, 0x4c, 0xcb, 0xae, 0x65, 0xfe, 0xd0, 0x32, 0xff, 0x99,
	0x98, 0x26, 0x14, 0xba, 0x79, 0x9c, 0x45, 0x39, 0x87, 0xc8, 0x0e, 0x67, 0x92, 0xe0, 0x8b, 0x28,
	0xad, 0x5c, 0x2c, 0xb2, 0x97, 0x78, 0xe0, 0x23, 0x67, 0xd4, 0x0e, 0x77, 0x33, 0xfc, 0x75, 0x74,
	0x20, 0xbe, 0x8f, 0x96, 0x43, 0xff, 0x72, 0x07, 0x0c, 0x0c, 0xdc, 0x22, 0x05, 0x60, 0x73, 0x04,
	0x3c, 0x6a, 0x73, 0x3b, 0xb0, 0xae, 0xf6, 0xd7, 0x30, 0xcb, 0xc3, 0x0f, 0xf0, 0x15, 0x94, 0x8d,
	0x7c, 0x13, 0xac, 0x20, 0xcc, 0xb8, 0x7e, 0x2a, 0xa3, 0x76, 0xd5, 0x1d, 0xfd, 0xef, 0x9c, 0x84,
	0xe3, 0x2d, 0x94, 0x82, 0xb9, 0x87, 0x19, 0x0e, 0x51, 0x25, 0x33, 0xd1, 0xde, 0xcd, 0x93, 0x02,
	0x98, 0xe0, 0x08, 0x4e, 0x88, 0x2d, 0x7c, 0x48, 0xa3, 0x14, 0x2c, 0xf0, 0x5f, 0x28, 0x03, 0xcb,
	0xb0, 0x63, 0x4f, 0x6b, 0x50, 0x93, 0xb4, 0x93, 0x86, 0x25, 0x74, 0x00, 0xfa, 0xab, 0x7a, 0x56,
	0x76, 0xe1, 0x6e, 0x01, 0x27, 0x83, 0x03, 0xd0, 0x70, 0xeb, 0x3d, 0xb7, 0xe3, 0x05, 0x70, 0x14,
	0x98, 0xf1, 0xac, 0x96, 0x68, 0x99, 0x79, 0x9d, 0xc0, 0x89, 0xf1, 0x95, 0xde, 0x9e, 0x46, 0xe3,
	0xab, 0xe8, 0x8f, 0xa8, 0xab, 0x40, 0xfc, 0xbc, 0x96, 0x7c, 0x56, 0x72, 0x3a, 0x7c, 0x40, 0x98,
	0x35, 0x41, 0xd8, 0x8b, 0x59, 0x85, 0x59, 0xe3, 0x84, 0xdd, 0x01, 0x93, 0x8b, 0x0f, 0x5e, 0x5b,
	0x15, 0x30, 0x94, 0xf8, 0x72, 0x06, 0x89, 0xfd, 0xfb, 0x47, 0x17, 0x1e, 0xb4, 0xda, 0x28, 0xa3,
	0x22, 0x6d, 0xe3, 0x55, 0xb2, 0xaa, 0x28, 0x34, 0x06, 0x95, 0x8d, 0xd7, 0xb3, 0x82, 0xca, 0x31,
	0xc8, 0x32, 0xde, 0xcc, 0x0a, 0xb2, 0xe2, 0x52, 0xda, 0x13, 0x4a, 0xf9, 0x76, 0xd6, 0x52, 0xda,
	0xe3, 0x4a, 0x79, 0x19, 0xe5, 0x18, 0xf5, 0x1f, 0x31, 0xb8, 0xe1, 0x8d, 0x77, 0xc9, 0x99, 0xe2,
	0xe0, 0xf0, 0x70, 0x50, 0xe9, 0x72, 0xd6, 0x82, 0xd3, 0x6c, 0xbc, 0xaf, 0x25, 0x5f, 0x54, 0x39,
	0x2a, 0x77, 0x55, 0x74, 0xe5, 0x36, 0x2a, 0x52, 0x3e, 0x32, 0x0f, 0xfd, 0x4f, 0xa4, 0x07, 0xe7,
	0x7f, 0xe5, 0xe3, 0xa9, 0x9e, 0x51, 0x1f, 0x2a, 0xf6, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x14, 0x9e, 0x40, 0x73, 0x09, 0x00, 0x00,
}