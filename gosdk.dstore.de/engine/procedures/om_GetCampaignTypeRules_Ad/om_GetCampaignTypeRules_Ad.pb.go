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

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetCampaignTypeRules_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x9a, 0xe6, 0xc1, 0xa5, 0x40, 0x65, 0x04, 0x1a, 0x12, 0x09, 0x41, 0xd9, 0x74, 0xe5,
	0xa0, 0x02, 0x55, 0x91, 0x58, 0xf0, 0x50, 0x81, 0x2e, 0x3a, 0x42, 0x16, 0x42, 0x02, 0x16, 0x23,
	0x37, 0x73, 0x19, 0x8d, 0x98, 0xb1, 0x23, 0xdb, 0x43, 0xc4, 0x5f, 0xf0, 0xfc, 0x2f, 0x3e, 0x81,
	0x2d, 0x2c, 0xf8, 0x06, 0xec, 0xb1, 0xa3, 0x64, 0x86, 0xa2, 0xb6, 0x9b, 0x44, 0x37, 0xe7, 0x9c,
	0xfb, 0x38, 0xf7, 0xc6, 0x70, 0x3f, 0xd5, 0x46, 0x2a, 0x9c, 0xa0, 0xc8, 0x72, 0x81, 0x93, 0x99,
	0x92, 0x53, 0x4c, 0x2b, 0x85, 0x7a, 0x22, 0xcb, 0xe4, 0x19, 0x9a, 0x27, 0xbc, 0x9c, 0xf1, 0x3c,
	0x13, 0x2f, 0x3f, 0xce, 0x90, 0x55, 0x05, 0xea, 0xe4, 0x51, 0x4a, 0x2d, 0xcb, 0x48, 0xb2, 0xed,
	0xa5, 0xd4, 0x4b, 0xe9, 0xff, 0xf9, 0xa3, 0xcb, 0xa1, 0xc8, 0x07, 0x5e, 0x54, 0xa8, 0xbd, 0x7c,
	0x74, 0xad, 0x59, 0x19, 0x95, 0x92, 0x2a, 0x40, 0xe3, 0x26, 0x54, 0xa2, 0xd6, 0x3c, 0xc3, 0x00,
	0xde, 0x6a, 0x83, 0x86, 0xe7, 0xe2, 0x9d, 0x54, 0x25, 0x37, 0xb9, 0x14, 0x9e, 0xb4, 0xf5, 0xb3,
	0x03, 0xf0, 0x82, 0x2b, 0x6e, 0x51, 0x54, 0x9a, 0xec, 0xc3, 0xe6, 0x34, 0x34, 0x96, 0x18, 0xdb,
	0x59, 0x92, 0xa7, 0x51, 0xe7, 0x46, 0x67, 0xfb, 0xfc, 0xce, 0x98, 0x86, 0x29, 0x42, 0x6f, 0xb9,
	0x30, 0x98, 0xa1, 0x7a, 0xe5, 0x22, 0x76, 0x71, 0xba, 0x32, 0xcd, 0x41, 0x4a, 0x76, 0xe0, 0x4a,
	0x3b, 0x4d, 0x22, 0xaa, 0xa2, 0x88, 0x7e, 0x0d, 0x6c, 0xb2, 0x21, 0x23, 0x4d, 0x7e, 0x6c, 0x21,
	0x72, 0x17, 0x06, 0x95, 0x46, 0xe5, 0x2a, 0xae, 0x9d, 0x5c, 0xb1, 0xef, 0xb8, 0xb6, 0xd2, 0x4d,
	0xd8, 0x08, 0x2a, 0x5f, 0xe0, 0xb7, 0x2f, 0x00, 0x1e, 0x76, 0x89, 0xb7, 0xfe, 0xac, 0xc3, 0x90,
	0xa1, 0x9e, 0x49, 0xa1, 0x91, 0xdc, 0x86, 0x5e, 0x6d, 0x60, 0x98, 0x6a, 0x44, 0x9b, 0xbb, 0xf1,
	0xe6, 0xee, 0xbb, 0x4f, 0xe6, 0x89, 0xe4, 0x35, 0x6c, 0x3a, 0xeb, 0x92, 0x15, 0xef, 0x6c, 0x83,
	0x5d, 0x2b, 0xa6, 0x2d, 0x71, 0xdb, 0xe1, 0x43, 0x1b, 0x1f, 0x2c, 0x63, 0x76, 0xa9, 0x6c, 0xfe,
	0x40, 0xf6, 0x60, 0x10, 0x56, 0x16, 0x75, 0xeb, 0x8c, 0xd7, 0xff, 0xc9, 0xe8, 0x17, 0x7a, 0xe8,
	0xbf, 0xd9, 0x82, 0x4e, 0x9e, 0x43, 0x57, 0xc9, 0x79, 0xb4, 0x5e, 0xab, 0x76, 0xe9, 0x69, 0x0f,
	0x8c, 0x2e, 0x7c, 0xa0, 0x4c, 0xce, 0x99, 0x4b, 0x31, 0xfa, 0xb1, 0x06, 0x5d, 0x1b, 0x90, 0xab,
	0xd0, 0xb7, 0xa1, 0x73, 0xff, 0x53, 0x6c, 0xad, 0xe9, 0xb1, 0x9e, 0x0d, 0xad, 0xc1, 0x7b, 0x70,
	0xae, 0x36, 0x58, 0xd8, 0x13, 0x89, 0x3e, 0xc7, 0x4d, 0xd7, 0xc2, 0x66, 0xb4, 0x51, 0xb9, 0xc8,
	0xfc, 0x62, 0x86, 0x8e, 0x1d, 0x5b, 0x32, 0x79, 0x7a, 0xcc, 0x2d, 0x7d, 0x89, 0xcf, 0x7e, 0x4c,
	0xf7, 0x96, 0x87, 0xf1, 0x35, 0x3e, 0xfd, 0x65, 0xec, 0xc2, 0x80, 0x17, 0x85, 0x9c, 0x63, 0x1a,
	0x7d, 0x3b, 0x5e, 0x76, 0x24, 0x65, 0x81, 0x5c, 0x78, 0xd9, 0x82, 0x4c, 0x1e, 0xc2, 0x85, 0x46,
	0xdb, 0xd1, 0xf7, 0x93, 0x87, 0xde, 0x58, 0x6d, 0xf9, 0xf1, 0x5b, 0x18, 0xe7, 0xb2, 0xb5, 0x93,
	0xe5, 0x7b, 0xf1, 0xe6, 0x41, 0x26, 0x75, 0xfa, 0x7e, 0x81, 0xa7, 0x67, 0x7b, 0x52, 0x8e, 0xfa,
	0xf5, 0xff, 0xf6, 0xce, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x31, 0x2c, 0xf8, 0x90, 0x04,
	0x00, 0x00,
}