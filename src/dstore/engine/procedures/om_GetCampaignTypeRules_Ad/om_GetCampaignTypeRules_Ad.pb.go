// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignTypeRules_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaignTypeRules_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignTypeRules_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignTypeRules_Ad

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
	CampaignTypeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignTypeIdNull bool                        `protobuf:"varint,1001,opt,name=campaign_type_id_null,json=campaignTypeIdNull" json:"campaign_type_id_null,omitempty"`
	UserId             *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull         bool                        `protobuf:"varint,1002,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
	}
	return nil
}

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
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
	RowId          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	UserName       *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	CampaignTypeId *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	UserId         *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Allowed        *dstore_values.BooleanValue `protobuf:"bytes,10004,opt,name=allowed" json:"allowed,omitempty"`
	CampaignType   *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=campaign_type,json=campaignType" json:"campaign_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetUserName() *dstore_values.StringValue {
	if m != nil {
		return m.UserName
	}
	return nil
}

func (m *Response_Row) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetAllowed() *dstore_values.BooleanValue {
	if m != nil {
		return m.Allowed
	}
	return nil
}

func (m *Response_Row) GetCampaignType() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignType
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignTypeRules_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignTypeRules_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignTypeRules_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x9a, 0xe6, 0x87, 0xa1, 0x40, 0x65, 0x04, 0x5a, 0x12, 0x09, 0x41, 0xb9, 0xf4, 0xe4,
	0xa0, 0x16, 0xaa, 0x72, 0xe3, 0x47, 0x05, 0x7a, 0xe8, 0x0a, 0x59, 0x08, 0x09, 0x2e, 0x2b, 0x37,
	0x3b, 0xac, 0x56, 0xda, 0xb5, 0x23, 0xdb, 0x4b, 0xc4, 0x5b, 0xf0, 0xfb, 0x5e, 0x3c, 0x02, 0x57,
	0x38, 0xf0, 0x0c, 0xd8, 0x6b, 0x47, 0xc9, 0x2e, 0x45, 0x0d, 0x5c, 0x12, 0x4d, 0xe6, 0xfb, 0xe6,
	0x9b, 0xf9, 0x66, 0x62, 0x78, 0x90, 0x6a, 0x23, 0x15, 0x4e, 0x50, 0x64, 0xb9, 0xc0, 0xc9, 0x4c,
	0xc9, 0x29, 0xa6, 0x95, 0x42, 0x3d, 0x91, 0x65, 0xf2, 0x0c, 0xcd, 0x13, 0x5e, 0xce, 0x78, 0x9e,
	0x89, 0x97, 0xef, 0x67, 0xc8, 0xaa, 0x02, 0x75, 0xf2, 0x28, 0xa5, 0x16, 0x65, 0x24, 0xd9, 0xf5,
	0x54, 0xea, 0xa9, 0xf4, 0xef, 0xf8, 0xd1, 0xd5, 0x20, 0xf2, 0x8e, 0x17, 0x15, 0x6a, 0x4f, 0x1f,
	0xdd, 0x68, 0x2a, 0xa3, 0x52, 0x52, 0x85, 0xd4, 0xb8, 0x99, 0x2a, 0x51, 0x6b, 0x9e, 0x61, 0x48,
	0xde, 0x69, 0x27, 0x0d, 0xcf, 0xc5, 0x5b, 0xa9, 0x4a, 0x6e, 0x72, 0x29, 0x3c, 0x68, 0xe7, 0x7b,
	0x07, 0xe0, 0x05, 0x57, 0xdc, 0x66, 0x51, 0x69, 0x72, 0x04, 0xdb, 0xd3, 0xd0, 0x58, 0x62, 0x6c,
	0x67, 0x49, 0x9e, 0x46, 0x9d, 0x5b, 0x9d, 0xdd, 0x8b, 0x7b, 0x63, 0x1a, 0xa6, 0x08, 0xbd, 0xe5,
	0xc2, 0x60, 0x86, 0xea, 0x95, 0x8b, 0xd8, 0xe5, 0xe9, 0xca, 0x34, 0xc7, 0x29, 0xd9, 0x83, 0x6b,
	0xed, 0x32, 0x89, 0xa8, 0x8a, 0x22, 0xfa, 0x31, 0xb0, 0xc5, 0x86, 0x8c, 0x34, 0xf1, 0xb1, 0x4d,
	0x91, 0x7b, 0x30, 0xa8, 0x34, 0x2a, 0xa7, 0xb8, 0x71, 0xbe, 0x62, 0xdf, 0x61, 0xad, 0xd2, 0x6d,
	0xd8, 0x0a, 0x2c, 0x2f, 0xf0, 0xd3, 0x0b, 0x80, 0x4f, 0xbb, 0xc2, 0x3b, 0xbf, 0x36, 0x61, 0xc8,
	0x50, 0xcf, 0xa4, 0xd0, 0x48, 0xee, 0x42, 0xaf, 0x36, 0x30, 0x4c, 0x35, 0xa2, 0xcd, 0xdd, 0x78,
	0x73, 0x8f, 0xdc, 0x27, 0xf3, 0x40, 0xf2, 0x1a, 0xb6, 0x9d, 0x75, 0xc9, 0x8a, 0x77, 0xb6, 0xc1,
	0xae, 0x25, 0xd3, 0x16, 0xb9, 0xed, 0xf0, 0x89, 0x8d, 0x8f, 0x97, 0x31, 0xbb, 0x52, 0x36, 0x7f,
	0x20, 0x87, 0x30, 0x08, 0x2b, 0x8b, 0xba, 0x75, 0xc5, 0x9b, 0x7f, 0x54, 0xf4, 0x0b, 0x3d, 0xf1,
	0xdf, 0x6c, 0x01, 0x27, 0xcf, 0xa1, 0xab, 0xe4, 0x3c, 0xda, 0xac, 0x59, 0x07, 0x74, 0xdd, 0x03,
	0xa3, 0x0b, 0x1f, 0x28, 0x93, 0x73, 0xe6, 0x4a, 0x8c, 0xbe, 0x6d, 0x40, 0xd7, 0x06, 0xe4, 0x3a,
	0xf4, 0x6d, 0xe8, 0xdc, 0xff, 0x10, 0x5b, 0x6b, 0x7a, 0xac, 0x67, 0x43, 0x6b, 0xf0, 0x21, 0x5c,
	0xa8, 0x0d, 0x16, 0xf6, 0x44, 0xa2, 0x8f, 0x71, 0xd3, 0xb5, 0xb0, 0x19, 0x6d, 0x54, 0x2e, 0x32,
	0xbf, 0x98, 0xa1, 0x43, 0xc7, 0x16, 0x4c, 0x9e, 0x9e, 0x71, 0x4b, 0x9f, 0xe2, 0x7f, 0x3f, 0xa6,
	0xfb, 0xcb, 0xc3, 0xf8, 0x1c, 0xaf, 0x7f, 0x19, 0x07, 0x30, 0xe0, 0x45, 0x21, 0xe7, 0x98, 0x46,
	0x5f, 0xce, 0xa6, 0x9d, 0x4a, 0x59, 0x20, 0x17, 0x9e, 0xb6, 0x00, 0x93, 0x87, 0x70, 0xa9, 0xd1,
	0x76, 0xf4, 0xf5, 0xfc, 0xa1, 0xb7, 0x56, 0x5b, 0x7e, 0xcc, 0x60, 0x9c, 0xcb, 0xd6, 0x4e, 0x96,
	0xef, 0xc5, 0x9b, 0xfd, 0xff, 0x78, 0x49, 0x4e, 0xfb, 0xf5, 0xdf, 0x75, 0xff, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x25, 0x61, 0xf6, 0x6d, 0x87, 0x04, 0x00, 0x00,
}