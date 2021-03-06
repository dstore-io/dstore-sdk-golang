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

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetCampaignTypes_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xcb, 0xaa, 0x13, 0x31,
	0x18, 0x66, 0x1c, 0x7b, 0x21, 0xde, 0x4a, 0x44, 0x19, 0xa7, 0x20, 0xa5, 0xba, 0x70, 0x95, 0x4a,
	0x45, 0xe8, 0xd2, 0x5b, 0x95, 0x2e, 0x3a, 0x48, 0x10, 0x45, 0x37, 0x43, 0xec, 0xfc, 0x0e, 0xc1,
	0x4e, 0x32, 0x24, 0xa9, 0xc5, 0x27, 0x70, 0xab, 0xbe, 0xc4, 0x79, 0x8e, 0xf3, 0x38, 0xe7, 0x2d,
	0x4e, 0x66, 0x92, 0xd2, 0xce, 0x9c, 0x2e, 0x7a, 0x36, 0x2d, 0xff, 0x7c, 0x97, 0xe4, 0xfb, 0x92,
	0xa0, 0x97, 0x99, 0x36, 0x52, 0xc1, 0x04, 0x44, 0xce, 0x05, 0x4c, 0x4a, 0x25, 0x57, 0x90, 0x6d,
	0x14, 0xe8, 0x89, 0x2c, 0xd2, 0x0f, 0x60, 0xde, 0xb2, 0xa2, 0x64, 0x3c, 0x17, 0x9f, 0x7e, 0x97,
	0xa0, 0xd3, 0xd7, 0x19, 0xb1, 0x0c, 0x23, 0xf1, 0x53, 0x27, 0x23, 0x4e, 0x46, 0x8e, 0x73, 0xe3,
	0xfb, 0xde, 0xfc, 0x17, 0x5b, 0x6f, 0x40, 0x3b, 0x69, 0xfc, 0xa8, 0xb9, 0x22, 0x28, 0x25, 0x95,
	0x87, 0x86, 0x4d, 0xa8, 0x00, 0xad, 0x59, 0x0e, 0x1e, 0x7c, 0xd2, 0x06, 0x0d, 0xe3, 0xe2, 0x87,
	0x54, 0x05, 0x33, 0x5c, 0x0a, 0x47, 0x1a, 0xff, 0x09, 0x10, 0xfa, 0xc8, 0x14, 0xb3, 0x28, 0x28,
	0x8d, 0xe7, 0x68, 0xb0, 0xf2, 0x9b, 0x4a, 0x8d, 0xdd, 0x55, 0xca, 0xb3, 0x28, 0x18, 0x05, 0xcf,
	0x6e, 0x4d, 0x87, 0xc4, 0x27, 0xf0, 0x7b, 0xe3, 0xc2, 0x40, 0x0e, 0xea, 0x73, 0x35, 0xd1, 0xbb,
	0xab, 0x83, 0x24, 0x8b, 0x0c, 0x4f, 0xd1, 0x83, 0xb6, 0x4d, 0x2a, 0x36, 0xeb, 0x75, 0x74, 0xd1,
	0xb3, 0x66, 0x7d, 0x8a, 0x9b, 0xfc, 0xc4, 0x42, 0xe3, 0xf3, 0x10, 0xf5, 0x29, 0xe8, 0x52, 0x0a,
	0x0d, 0xf8, 0x39, 0xea, 0xd4, 0x39, 0xfd, 0xe2, 0x31, 0x69, 0xd6, 0xe7, 0x3a, 0x98, 0x57, 0xbf,
	0xd4, 0x11, 0xf1, 0x57, 0x34, 0xa8, 0x12, 0xa6, 0x07, 0x11, 0xa3, 0x1b, 0xa3, 0xd0, 0x8a, 0x49,
	0x4b, 0xdc, 0x2e, 0x62, 0x69, 0xe7, 0xc5, 0x7e, 0xa6, 0xf7, 0x8a, 0xe6, 0x07, 0x3c, 0x43, 0x3d,
	0xdf, 0x6c, 0x14, 0xd6, 0x8e, 0x8f, 0xaf, 0x38, 0xba, 0xde, 0x97, 0xee, 0x9f, 0xee, 0xe8, 0xf8,
	0x1d, 0x0a, 0x95, 0xdc, 0x46, 0x37, 0x6b, 0xd5, 0x94, 0x9c, 0x72, 0x07, 0xc8, 0xae, 0x03, 0x42,
	0xe5, 0x96, 0x56, 0xf2, 0xf8, 0x2c, 0x40, 0xa1, 0x1d, 0xf0, 0x43, 0xd4, 0xb5, 0x63, 0x75, 0x24,
	0x7f, 0x13, 0x5b, 0x4b, 0x87, 0x76, 0xec, 0x68, 0xdb, 0x7e, 0x7f, 0xe4, 0xd0, 0xfe, 0x25, 0xd7,
	0x3f, 0xb5, 0x57, 0xe8, 0x4e, 0xc3, 0x27, 0xfa, 0x9f, 0x34, 0xdb, 0xf7, 0x26, 0xda, 0x28, 0x2e,
	0x72, 0xe7, 0x71, 0xfb, 0xd0, 0xe3, 0xcd, 0x17, 0x34, 0xe4, 0xb2, 0x15, 0x73, 0xff, 0x42, 0xbe,
	0xcd, 0x72, 0xa9, 0xb3, 0x9f, 0x3b, 0x3c, 0x3b, 0xfd, 0x11, 0x7d, 0xef, 0xd6, 0xb7, 0xf5, 0xc5,
	0x65, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x05, 0x4b, 0x17, 0x7e, 0x03, 0x00, 0x00,
}
