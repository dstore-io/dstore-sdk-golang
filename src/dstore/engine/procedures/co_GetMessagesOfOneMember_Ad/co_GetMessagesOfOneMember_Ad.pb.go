// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_GetMessagesOfOneMember_Ad.proto
// DO NOT EDIT!

/*
Package co_GetMessagesOfOneMember_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_GetMessagesOfOneMember_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_GetMessagesOfOneMember_Ad

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
	CommunityMemberId             *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	CommunityMemberIdNull         bool                          `protobuf:"varint,1001,opt,name=community_member_id_null,json=communityMemberIdNull" json:"community_member_id_null,omitempty"`
	CommunityId                   *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull               bool                          `protobuf:"varint,1002,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	CommunicatingMemberId         *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=communicating_member_id,json=communicatingMemberId" json:"communicating_member_id,omitempty"`
	CommunicatingMemberIdNull     bool                          `protobuf:"varint,1003,opt,name=communicating_member_id_null,json=communicatingMemberIdNull" json:"communicating_member_id_null,omitempty"`
	SentOrReceivedMessages        *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=sent_or_received_messages,json=sentOrReceivedMessages" json:"sent_or_received_messages,omitempty"`
	SentOrReceivedMessagesNull    bool                          `protobuf:"varint,1004,opt,name=sent_or_received_messages_null,json=sentOrReceivedMessagesNull" json:"sent_or_received_messages_null,omitempty"`
	FromDate                      *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                  bool                          `protobuf:"varint,1005,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                        *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                    bool                          `protobuf:"varint,1006,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	OrderByMessageDateAndTime     *dstore_values.IntegerValue   `protobuf:"bytes,7,opt,name=order_by_message_date_and_time,json=orderByMessageDateAndTime" json:"order_by_message_date_and_time,omitempty"`
	OrderByMessageDateAndTimeNull bool                          `protobuf:"varint,1007,opt,name=order_by_message_date_and_time_null,json=orderByMessageDateAndTimeNull" json:"order_by_message_date_and_time_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetCommunicatingMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunicatingMemberId
	}
	return nil
}

func (m *Parameters) GetSentOrReceivedMessages() *dstore_values.IntegerValue {
	if m != nil {
		return m.SentOrReceivedMessages
	}
	return nil
}

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetOrderByMessageDateAndTime() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderByMessageDateAndTime
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
	ToCommunityMember      *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=to_community_member,json=toCommunityMember" json:"to_community_member,omitempty"`
	MessageStatus          *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=message_status,json=messageStatus" json:"message_status,omitempty"`
	ToCommunityMemberId    *dstore_values.IntegerValue   `protobuf:"bytes,10003,opt,name=to_community_member_id,json=toCommunityMemberId" json:"to_community_member_id,omitempty"`
	Message                *dstore_values.StringValue    `protobuf:"bytes,10004,opt,name=message" json:"message,omitempty"`
	FromCommunityMember    *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=from_community_member,json=fromCommunityMember" json:"from_community_member,omitempty"`
	CommunityId            *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityName          *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=community_name,json=communityName" json:"community_name,omitempty"`
	MessageDateAndTime     *dstore_values.TimestampValue `protobuf:"bytes,10008,opt,name=message_date_and_time,json=messageDateAndTime" json:"message_date_and_time,omitempty"`
	FromCommunityMemberId  *dstore_values.IntegerValue   `protobuf:"bytes,10009,opt,name=from_community_member_id,json=fromCommunityMemberId" json:"from_community_member_id,omitempty"`
	MessageDateAndTimeChar *dstore_values.StringValue    `protobuf:"bytes,10010,opt,name=message_date_and_time_char,json=messageDateAndTimeChar" json:"message_date_and_time_char,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetToCommunityMember() *dstore_values.StringValue {
	if m != nil {
		return m.ToCommunityMember
	}
	return nil
}

func (m *Response_Row) GetMessageStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.MessageStatus
	}
	return nil
}

func (m *Response_Row) GetToCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToCommunityMemberId
	}
	return nil
}

func (m *Response_Row) GetMessage() *dstore_values.StringValue {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Response_Row) GetFromCommunityMember() *dstore_values.StringValue {
	if m != nil {
		return m.FromCommunityMember
	}
	return nil
}

func (m *Response_Row) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Response_Row) GetCommunityName() *dstore_values.StringValue {
	if m != nil {
		return m.CommunityName
	}
	return nil
}

func (m *Response_Row) GetMessageDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.MessageDateAndTime
	}
	return nil
}

func (m *Response_Row) GetFromCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromCommunityMemberId
	}
	return nil
}

func (m *Response_Row) GetMessageDateAndTimeChar() *dstore_values.StringValue {
	if m != nil {
		return m.MessageDateAndTimeChar
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_GetMessagesOfOneMember_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_GetMessagesOfOneMember_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_GetMessagesOfOneMember_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x6b, 0x4f, 0x13, 0x4d,
	0x14, 0x0e, 0x6f, 0xdf, 0x5e, 0xde, 0xc3, 0xed, 0x65, 0x9b, 0xd6, 0x6d, 0x11, 0xa2, 0x10, 0x13,
	0xa3, 0xc9, 0x62, 0x34, 0x18, 0xa2, 0x89, 0x0a, 0x68, 0x0c, 0x1a, 0x5a, 0x5d, 0x10, 0xa3, 0x1f,
	0xdc, 0x2c, 0xdd, 0xa1, 0x6e, 0xc2, 0xce, 0x90, 0xd9, 0x29, 0x84, 0x7f, 0xa1, 0x78, 0xf7, 0x8b,
	0xff, 0xc1, 0x7f, 0xe5, 0xfd, 0x2f, 0x78, 0x66, 0x67, 0x7a, 0xdb, 0x2e, 0x6c, 0xf5, 0x0b, 0xc9,
	0xf4, 0x9c, 0xe7, 0x39, 0xcf, 0x39, 0x33, 0xcf, 0x61, 0xe1, 0xba, 0x17, 0x0a, 0xc6, 0xc9, 0x02,
	0xa1, 0x4d, 0x9f, 0x92, 0x85, 0x3d, 0xce, 0x1a, 0xc4, 0x6b, 0x71, 0x12, 0x2e, 0x34, 0x98, 0x73,
	0x97, 0x88, 0x75, 0x12, 0x86, 0x6e, 0x93, 0x84, 0xf5, 0x9d, 0x3a, 0x25, 0xeb, 0x24, 0xd8, 0x26,
	0xdc, 0x59, 0xf6, 0x2c, 0xcc, 0x13, 0xcc, 0xb8, 0xa0, 0xc0, 0x96, 0x02, 0x5b, 0x27, 0x21, 0xaa,
	0x45, 0x5d, 0x68, 0xdf, 0xdd, 0x6d, 0x91, 0x50, 0x11, 0x54, 0x2b, 0xfd, 0xd5, 0x09, 0xe7, 0x8c,
	0xeb, 0xd0, 0x74, 0x7f, 0x28, 0x50, 0xac, 0x3a, 0x38, 0x1f, 0x0f, 0x0a, 0xd7, 0xa7, 0x3b, 0x8c,
	0x07, 0xae, 0xf0, 0x19, 0x55, 0x49, 0x73, 0x9f, 0xf2, 0x00, 0x0f, 0x5c, 0xee, 0x62, 0x94, 0xf0,
	0xd0, 0xb8, 0x0f, 0xc5, 0x06, 0x0b, 0x82, 0x16, 0xf5, 0xc5, 0xa1, 0x13, 0x28, 0x5d, 0xbe, 0x67,
	0x8e, 0x9c, 0x19, 0x39, 0x3f, 0x7a, 0x79, 0xda, 0xd2, 0xad, 0x68, 0x79, 0x3e, 0x15, 0xa4, 0x49,
	0xf8, 0x96, 0x3c, 0xd9, 0x53, 0x1d, 0x9c, 0x6a, 0x67, 0xcd, 0x33, 0x96, 0xc0, 0x4c, 0x20, 0x73,
	0x68, 0x6b, 0x77, 0xd7, 0xfc, 0x92, 0x47, 0xca, 0x82, 0x5d, 0x1a, 0x40, 0xd5, 0x30, 0x6a, 0xdc,
	0x80, 0xb1, 0x2e, 0x12, 0xeb, 0xff, 0x93, 0x5e, 0x7f, 0xb4, 0x03, 0xc0, 0xca, 0x17, 0x61, 0xaa,
	0x17, 0xaf, 0x4a, 0x7e, 0x55, 0x25, 0x27, 0x7b, 0x12, 0xa3, 0x62, 0x1b, 0x70, 0x4a, 0xff, 0xd4,
	0xc0, 0xc9, 0xd0, 0x66, 0x4f, 0xdf, 0x99, 0xf4, 0xba, 0xa5, 0x3e, 0x6c, 0xa7, 0xf7, 0x5b, 0x70,
	0xfa, 0x18, 0x52, 0x25, 0xe6, 0x9b, 0x12, 0x53, 0x49, 0x44, 0x47, 0xb2, 0xb6, 0xa0, 0x12, 0x12,
	0x2a, 0x1c, 0xc6, 0x1d, 0x4e, 0x1a, 0xc4, 0xdf, 0x27, 0x9e, 0xa3, 0x2f, 0x38, 0x34, 0xff, 0x4d,
	0x17, 0x56, 0x96, 0xe8, 0x3a, 0xb7, 0x35, 0xb6, 0xfd, 0xe2, 0x8c, 0x55, 0x98, 0x3d, 0x96, 0x57,
	0x69, 0xfb, 0xae, 0xb4, 0x55, 0x93, 0x09, 0x22, 0x71, 0xd7, 0xe0, 0xbf, 0x1d, 0xce, 0x02, 0xc7,
	0x73, 0x05, 0x31, 0xb3, 0x91, 0x98, 0x99, 0x98, 0x18, 0xe1, 0x23, 0xab, 0x70, 0x83, 0x3d, 0x25,
	0xa7, 0x20, 0xf3, 0x6f, 0x63, 0xba, 0x71, 0x0e, 0x26, 0x3a, 0x58, 0x55, 0xf0, 0x87, 0x2a, 0x38,
	0xd6, 0x4e, 0x89, 0x4a, 0x5c, 0x85, 0xbc, 0x60, 0xaa, 0x40, 0x6e, 0x98, 0x02, 0x39, 0xc1, 0x22,
	0xfa, 0xb3, 0x30, 0xa6, 0x71, 0x8a, 0xfc, 0xa7, 0x22, 0x07, 0x15, 0x8e, 0xa8, 0x9f, 0xc1, 0x2c,
	0xe3, 0x1e, 0xde, 0xc6, 0xf6, 0x61, 0xbb, 0x75, 0x05, 0x70, 0xa9, 0xe7, 0x48, 0x5a, 0x33, 0x9f,
	0x3e, 0xdf, 0x4a, 0x44, 0xb1, 0x72, 0xa8, 0xe7, 0x22, 0xc9, 0x97, 0xa9, 0xb7, 0x89, 0x68, 0x74,
	0xd1, 0xfc, 0xc9, 0xfc, 0x4a, 0xd9, 0x2f, 0xa5, 0x6c, 0xe6, 0x58, 0x22, 0x29, 0x76, 0xee, 0x73,
	0x01, 0x0a, 0x36, 0x09, 0xf7, 0x18, 0x0d, 0x89, 0x71, 0x09, 0xb2, 0x91, 0xff, 0xb5, 0x23, 0xab,
	0x56, 0xff, 0x72, 0x51, 0xbb, 0xe1, 0x8e, 0xfc, 0x6b, 0xab, 0x44, 0xe3, 0x09, 0xfc, 0x2f, 0x9d,
	0xef, 0xf4, 0x58, 0x1f, 0xed, 0x94, 0x41, 0xb0, 0x15, 0x03, 0xc7, 0x17, 0xc4, 0x3a, 0x9e, 0xd7,
	0xba, 0x67, 0x7b, 0x32, 0xe8, 0xff, 0x01, 0xfd, 0x9d, 0xd7, 0xdd, 0xa1, 0x51, 0x24, 0xe3, 0xec,
	0x00, 0xa3, 0xda, 0x47, 0xba, 0x33, 0xbb, 0x9d, 0x6e, 0xdc, 0x83, 0x0c, 0x67, 0x07, 0xf8, 0x8a,
	0x25, 0x6a, 0xc9, 0x1a, 0x7e, 0x43, 0x5a, 0xed, 0x49, 0x58, 0x36, 0x3b, 0xb0, 0x25, 0x49, 0xf5,
	0x28, 0x07, 0x19, 0x3c, 0x18, 0x65, 0xc8, 0xe1, 0x51, 0xba, 0xf6, 0x45, 0x0d, 0x87, 0x93, 0xb5,
	0xb3, 0x78, 0x44, 0x27, 0xe2, 0x4a, 0xc3, 0xf7, 0x10, 0x5f, 0x44, 0xe6, 0xcb, 0x5a, 0xff, 0x04,
	0xf5, 0x15, 0x87, 0x82, 0xa3, 0x11, 0xf5, 0x4a, 0x13, 0x6c, 0xb5, 0x7f, 0x3d, 0xa1, 0x79, 0x26,
	0xda, 0x17, 0x8a, 0x4f, 0x4f, 0xb4, 0x42, 0xf3, 0xa8, 0x96, 0xfe, 0x54, 0xc6, 0x35, 0x66, 0x23,
	0x82, 0x18, 0x0f, 0xa1, 0x9c, 0xa0, 0x48, 0x2a, 0x7f, 0x35, 0x04, 0x59, 0x71, 0x40, 0x15, 0x36,
	0xb9, 0xd8, 0xbd, 0x8a, 0xd7, 0xe9, 0x8d, 0x75, 0xee, 0xa1, 0x0e, 0xa5, 0xc8, 0x8a, 0x03, 0xd3,
	0x79, 0x93, 0x4e, 0x52, 0x94, 0xc8, 0xf8, 0x7c, 0x6e, 0xc6, 0x16, 0xf7, 0xdb, 0xda, 0x1f, 0x6e,
	0xee, 0x15, 0x98, 0xe8, 0x12, 0x50, 0xfc, 0xb7, 0x64, 0xbe, 0x4b, 0x97, 0x32, 0xde, 0x81, 0xd4,
	0x10, 0x81, 0xf3, 0x2d, 0x25, 0xbb, 0xfa, 0x7d, 0x6d, 0x98, 0x45, 0x62, 0x04, 0x83, 0x8e, 0x7e,
	0x04, 0x66, 0xe2, 0xa0, 0x64, 0x8f, 0x1f, 0x86, 0xe8, 0xb1, 0x94, 0x30, 0x2c, 0xec, 0xf6, 0x31,
	0x54, 0x93, 0xf7, 0x43, 0xe3, 0xb9, 0xcb, 0xcd, 0x8f, 0xe9, 0x9d, 0x97, 0x07, 0xb5, 0xae, 0x22,
	0x74, 0x65, 0x13, 0xa6, 0x7d, 0x16, 0xf3, 0x55, 0xf7, 0xb3, 0xe5, 0xe9, 0xe2, 0x5f, 0x7d, 0xd0,
	0x6c, 0xe7, 0xa2, 0x6f, 0x86, 0x2b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xaa, 0xe1, 0x0a,
	0x10, 0x09, 0x00, 0x00,
}
