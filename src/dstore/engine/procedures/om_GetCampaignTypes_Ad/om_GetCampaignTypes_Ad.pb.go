// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCampaignTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_GetCampaignTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCampaignTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCampaignTypes_Ad

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
	CampaignTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignType   *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=campaign_type,json=campaignType" json:"campaign_type,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCampaignTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCampaignTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCampaignTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xcb, 0xaa, 0x13, 0x31,
	0x18, 0x66, 0x1c, 0x7b, 0x21, 0xde, 0x4a, 0x44, 0x19, 0xa7, 0x20, 0xa5, 0xba, 0x70, 0x95, 0x91,
	0x8a, 0xe0, 0xd2, 0x5b, 0x95, 0x2e, 0x3a, 0x48, 0x10, 0x41, 0x37, 0x43, 0xec, 0xfc, 0x0e, 0x81,
	0x4e, 0x32, 0x24, 0xa9, 0xc5, 0x27, 0x70, 0xab, 0xe7, 0x25, 0xce, 0x73, 0x9c, 0xc7, 0x39, 0x6f,
	0x71, 0x32, 0x93, 0x94, 0x76, 0xe6, 0x74, 0x71, 0xba, 0x69, 0xf9, 0xe7, 0xbb, 0x24, 0xdf, 0x97,
	0x04, 0xbd, 0xce, 0xb5, 0x91, 0x0a, 0x12, 0x10, 0x05, 0x17, 0x90, 0x54, 0x4a, 0xae, 0x20, 0xdf,
	0x28, 0xd0, 0x89, 0x2c, 0xb3, 0xcf, 0x60, 0x3e, 0xb0, 0xb2, 0x62, 0xbc, 0x10, 0x5f, 0xff, 0x54,
	0xa0, 0xb3, 0x77, 0x39, 0xb1, 0x0c, 0x23, 0xf1, 0x73, 0x27, 0x23, 0x4e, 0x46, 0x8e, 0x73, 0xe3,
	0x87, 0xde, 0xfc, 0x37, 0x5b, 0x6f, 0x40, 0x3b, 0x69, 0xfc, 0xa4, 0xbd, 0x22, 0x28, 0x25, 0x95,
	0x87, 0xc6, 0x6d, 0xa8, 0x04, 0xad, 0x59, 0x01, 0x1e, 0x7c, 0xd6, 0x05, 0x0d, 0xe3, 0xe2, 0x97,
	0x54, 0x25, 0x33, 0x5c, 0x0a, 0x47, 0x9a, 0xfe, 0x0d, 0x10, 0xfa, 0xc2, 0x14, 0xb3, 0x28, 0x28,
	0x8d, 0xe7, 0x68, 0xb4, 0xf2, 0x9b, 0xca, 0x8c, 0xdd, 0x55, 0xc6, 0xf3, 0x28, 0x98, 0x04, 0x2f,
	0xee, 0xcc, 0xc6, 0xc4, 0x27, 0xf0, 0x7b, 0xe3, 0xc2, 0x40, 0x01, 0xea, 0x5b, 0x3d, 0xd1, 0xfb,
	0xab, 0x83, 0x24, 0x8b, 0x1c, 0xcf, 0xd0, 0xa3, 0xae, 0x4d, 0x26, 0x36, 0xeb, 0x75, 0x74, 0x39,
	0xb0, 0x66, 0x43, 0x8a, 0xdb, 0xfc, 0xd4, 0x42, 0xd3, 0x8b, 0x10, 0x0d, 0x29, 0xe8, 0x4a, 0x0a,
	0x0d, 0xf8, 0x25, 0xea, 0x35, 0x39, 0xfd, 0xe2, 0x31, 0x69, 0xd7, 0xe7, 0x3a, 0x98, 0xd7, 0xbf,
	0xd4, 0x11, 0xf1, 0x77, 0x34, 0xaa, 0x13, 0x66, 0x07, 0x11, 0xa3, 0x5b, 0x93, 0xd0, 0x8a, 0x49,
	0x47, 0xdc, 0x2d, 0x62, 0x69, 0xe7, 0xc5, 0x7e, 0xa6, 0x0f, 0xca, 0xf6, 0x07, 0xfc, 0x06, 0x0d,
	0x7c, 0xb3, 0x51, 0xd8, 0x38, 0x3e, 0xbd, 0xe6, 0xe8, 0x7a, 0x5f, 0xba, 0x7f, 0xba, 0xa3, 0xe3,
	0x8f, 0x28, 0x54, 0x72, 0x1b, 0xdd, 0x6e, 0x54, 0x33, 0x72, 0x93, 0x3b, 0x40, 0x76, 0x1d, 0x10,
	0x2a, 0xb7, 0xb4, 0x96, 0xc7, 0xe7, 0x01, 0x0a, 0xed, 0x80, 0x1f, 0xa3, 0xbe, 0x1d, 0xeb, 0x23,
	0xf9, 0x97, 0xda, 0x5a, 0x7a, 0xb4, 0x67, 0x47, 0xdb, 0xf6, 0xa7, 0x23, 0x87, 0xf6, 0x3f, 0x3d,
	0xfd, 0xd4, 0xde, 0xa2, 0x7b, 0x2d, 0x9f, 0xe8, 0x2c, 0x6d, 0xb7, 0xef, 0x4d, 0xb4, 0x51, 0x5c,
	0x14, 0xce, 0xe3, 0xee, 0xa1, 0xc7, 0xfb, 0x14, 0x8d, 0xb9, 0xec, 0xc4, 0xdc, 0xbf, 0x90, 0x1f,
	0xc9, 0x89, 0x6f, 0xe7, 0x67, 0xbf, 0xb9, 0xa4, 0xaf, 0xae, 0x02, 0x00, 0x00, 0xff, 0xff, 0x06,
	0x22, 0x8a, 0xbe, 0x75, 0x03, 0x00, 0x00,
}