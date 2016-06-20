// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampaignTypeRules_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampaignTypeRules_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampaignTypeRules_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampaignTypeRules_Ad

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
	Allowed            *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=allowed" json:"allowed,omitempty"`
	AllowedNull        bool                        `protobuf:"varint,1003,opt,name=allowed_null,json=allowedNull" json:"allowed_null,omitempty"`
	DeleteRule         *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=delete_rule,json=deleteRule" json:"delete_rule,omitempty"`
	DeleteRuleNull     bool                        `protobuf:"varint,1004,opt,name=delete_rule_null,json=deleteRuleNull" json:"delete_rule_null,omitempty"`
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

func (m *Parameters) GetAllowed() *dstore_values.BooleanValue {
	if m != nil {
		return m.Allowed
	}
	return nil
}

func (m *Parameters) GetDeleteRule() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteRule
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampaignTypeRules_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampaignTypeRules_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampaignTypeRules_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x66, 0x77, 0x6c, 0xbb, 0x9c, 0x2e, 0x6b, 0x89, 0x28, 0xb5, 0x45, 0xd1, 0x7a, 0xa3, 0x08,
	0xa9, 0xac, 0x3f, 0x28, 0xec, 0x8d, 0xca, 0x5e, 0x14, 0xe9, 0x22, 0x41, 0x05, 0xbd, 0x19, 0xb2,
	0xcd, 0xd9, 0x32, 0x30, 0x93, 0x94, 0x64, 0xc6, 0xb2, 0x6f, 0xe1, 0x43, 0xf8, 0x1a, 0x3e, 0x90,
	0x3f, 0x0f, 0x61, 0xd2, 0x93, 0xd2, 0xce, 0x28, 0x2e, 0xee, 0x4d, 0x3b, 0x27, 0xe7, 0xfb, 0x49,
	0xce, 0x0f, 0x1c, 0x29, 0x57, 0x1a, 0x8b, 0x63, 0xd4, 0xf3, 0x4c, 0xe3, 0x78, 0x61, 0xcd, 0x0c,
	0x55, 0x65, 0xd1, 0x8d, 0x4d, 0x91, 0x4e, 0x8d, 0xca, 0xce, 0xce, 0x5f, 0xcb, 0x62, 0x21, 0xb3,
	0xb9, 0x7e, 0x77, 0xbe, 0x40, 0x51, 0xe5, 0xe8, 0xd2, 0x97, 0x8a, 0x7b, 0x60, 0x69, 0xd8, 0x43,
	0x62, 0x73, 0x62, 0xf3, 0x7f, 0x52, 0x06, 0xd7, 0xa2, 0xd5, 0x67, 0x99, 0x57, 0xe8, 0x48, 0x61,
	0x70, 0xb3, 0xee, 0x8f, 0xd6, 0x1a, 0x1b, 0x53, 0xc3, 0x7a, 0xaa, 0x40, 0xe7, 0xe4, 0x1c, 0x63,
	0xf2, 0x5e, 0x33, 0x59, 0xca, 0x4c, 0x9f, 0x19, 0x5b, 0xc8, 0x32, 0x33, 0x9a, 0x40, 0xa3, 0xaf,
	0x09, 0xc0, 0x5b, 0x69, 0xa5, 0xcf, 0xa2, 0x75, 0xec, 0x18, 0x7a, 0xb3, 0x78, 0xb1, 0xb4, 0xf4,
	0x37, 0x4b, 0x33, 0xd5, 0xdf, 0xb9, 0xb3, 0x73, 0xbf, 0x7b, 0x38, 0xe4, 0xf1, 0x21, 0xf1, 0x6e,
	0x99, 0x2e, 0x71, 0x8e, 0xf6, 0x43, 0x88, 0xc4, 0xc1, 0x6c, 0xeb, 0x35, 0x13, 0xc5, 0x0e, 0xe1,
	0x7a, 0x53, 0x26, 0xd5, 0x55, 0x9e, 0xf7, 0xbf, 0x77, 0xbc, 0xd8, 0x9e, 0x60, 0x75, 0xfc, 0x89,
	0x4f, 0xb1, 0x27, 0xd0, 0xa9, 0x1c, 0xda, 0xe0, 0xb8, 0x7b, 0xb1, 0x63, 0x3b, 0x60, 0xbd, 0xd3,
	0x5d, 0xd8, 0x8f, 0x2c, 0x32, 0xf8, 0x41, 0x06, 0x40, 0xe9, 0x95, 0xf0, 0x53, 0xe8, 0xc8, 0x3c,
	0x37, 0x4b, 0x54, 0xfd, 0xe4, 0xaf, 0xc2, 0xa7, 0xc6, 0xe4, 0x28, 0x35, 0x09, 0xaf, 0xb1, 0x6c,
	0x04, 0xfb, 0xf1, 0x93, 0x94, 0x7f, 0x92, 0x72, 0x37, 0x1e, 0xae, 0xa4, 0x8f, 0xa0, 0xab, 0x30,
	0xf7, 0xa5, 0x4b, 0xad, 0x6f, 0x61, 0xff, 0xca, 0xc5, 0xf2, 0x40, 0xf8, 0xd0, 0x71, 0xf6, 0x00,
	0x7a, 0x5b, 0x6c, 0x72, 0xf9, 0x45, 0x2e, 0x07, 0x1b, 0x58, 0x30, 0x1a, 0x7d, 0xdb, 0x85, 0x3d,
	0x81, 0x6e, 0x61, 0xb4, 0x43, 0xf6, 0x08, 0x5a, 0xab, 0x21, 0x88, 0x9d, 0x19, 0xf0, 0xfa, 0x88,
	0xd1, 0x80, 0x1c, 0x87, 0x5f, 0x41, 0x40, 0xf6, 0x11, 0x7a, 0xa1, 0xfd, 0xe9, 0x56, 0xff, 0x7d,
	0x91, 0x13, 0x4f, 0xe6, 0x0d, 0x72, 0x73, 0x4a, 0xa6, 0x3e, 0x9e, 0x6c, 0x62, 0x71, 0xb5, 0xa8,
	0x1f, 0xb0, 0xe7, 0xd0, 0x89, 0x63, 0xe7, 0xab, 0x1b, 0x14, 0x6f, 0xff, 0xa1, 0x48, 0x43, 0x39,
	0xa5, 0x7f, 0xb1, 0x86, 0xb3, 0x37, 0x90, 0x58, 0xb3, 0xf4, 0x45, 0x0b, 0xac, 0x17, 0xfc, 0x3f,
	0xf6, 0x84, 0xaf, 0x4b, 0xc1, 0x85, 0x59, 0x8a, 0xa0, 0x32, 0xb8, 0x05, 0x89, 0xff, 0x66, 0x37,
	0xa0, 0xed, 0xa3, 0x30, 0x43, 0x5f, 0x4e, 0x7c, 0x71, 0x5a, 0xa2, 0xe5, 0xc3, 0x89, 0x7a, 0xf5,
	0x1e, 0x86, 0x99, 0x69, 0x58, 0x6c, 0x16, 0xf9, 0xd3, 0xb3, 0xcb, 0xad, 0xf8, 0x69, 0x7b, 0xb5,
	0x44, 0x8f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x19, 0x9a, 0x6c, 0x23, 0x04, 0x00, 0x00,
}
