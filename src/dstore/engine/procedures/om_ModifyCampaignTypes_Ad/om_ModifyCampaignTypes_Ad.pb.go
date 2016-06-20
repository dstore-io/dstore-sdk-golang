// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampaignTypes_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampaignTypes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampaignTypes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampaignTypes_Ad

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
	CampaignTypeId         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignTypeIdNull     bool                        `protobuf:"varint,1001,opt,name=campaign_type_id_null,json=campaignTypeIdNull" json:"campaign_type_id_null,omitempty"`
	CampaignType           *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=campaign_type,json=campaignType" json:"campaign_type,omitempty"`
	CampaignTypeNull       bool                        `protobuf:"varint,1002,opt,name=campaign_type_null,json=campaignTypeNull" json:"campaign_type_null,omitempty"`
	DeleteCampaignType     *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=delete_campaign_type,json=deleteCampaignType" json:"delete_campaign_type,omitempty"`
	DeleteCampaignTypeNull bool                        `protobuf:"varint,1003,opt,name=delete_campaign_type_null,json=deleteCampaignTypeNull" json:"delete_campaign_type_null,omitempty"`
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

func (m *Parameters) GetCampaignType() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignType
	}
	return nil
}

func (m *Parameters) GetDeleteCampaignType() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCampaignType
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampaignTypes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampaignTypes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampaignTypes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xa6, 0x8d, 0xdb, 0x96, 0xe3, 0xdf, 0x32, 0x6a, 0x49, 0xb3, 0x28, 0x52, 0x2f, 0xf4, 0xc6,
	0x59, 0x59, 0x11, 0x8a, 0x37, 0xa2, 0x52, 0x64, 0x2f, 0x52, 0x74, 0x10, 0x41, 0x6f, 0xc2, 0x74,
	0xe7, 0x34, 0x0c, 0x24, 0x33, 0x61, 0x26, 0x6b, 0xe9, 0xad, 0x4f, 0xe0, 0xeb, 0xf8, 0x48, 0xea,
	0x4b, 0x38, 0xc9, 0x4c, 0xdd, 0x24, 0xbb, 0x82, 0xbd, 0x49, 0x72, 0x72, 0xce, 0xf7, 0x7d, 0xe7,
	0x17, 0x8e, 0x84, 0xad, 0xb5, 0xc1, 0x29, 0xaa, 0x5c, 0x2a, 0x9c, 0x56, 0x46, 0x2f, 0x50, 0x2c,
	0x0d, 0xda, 0xa9, 0x2e, 0xb3, 0x54, 0x0b, 0x79, 0x76, 0xf1, 0x96, 0x97, 0x15, 0x97, 0xb9, 0xfa,
	0x78, 0x51, 0xa1, 0xcd, 0x5e, 0x0b, 0xea, 0x82, 0x6a, 0x4d, 0x1e, 0x7b, 0x24, 0xf5, 0x48, 0xfa,
	0xcf, 0xf0, 0xe4, 0x4e, 0x90, 0xf8, 0xca, 0x8b, 0x25, 0x5a, 0x8f, 0x4e, 0x0e, 0xfa, 0xba, 0x68,
	0x8c, 0x36, 0xc1, 0x35, 0xe9, 0xbb, 0x4a, 0xb4, 0x96, 0xe7, 0x18, 0x9c, 0x8f, 0x86, 0xce, 0x9a,
	0x4b, 0x75, 0xa6, 0x4d, 0xc9, 0x6b, 0xa9, 0x95, 0x0f, 0x3a, 0xfc, 0x16, 0x01, 0xbc, 0xe7, 0x86,
	0x3b, 0x2f, 0x1a, 0x4b, 0x8e, 0x61, 0xbc, 0x08, 0x49, 0x65, 0xb5, 0xcb, 0x2a, 0x93, 0x22, 0xde,
	0x7a, 0xb8, 0xf5, 0xe4, 0xfa, 0x6c, 0x42, 0x43, 0x11, 0x21, 0x37, 0xa9, 0x6a, 0xcc, 0xd1, 0x7c,
	0x6a, 0x2c, 0x76, 0x6b, 0xd1, 0xa9, 0x64, 0x2e, 0xc8, 0x0c, 0xee, 0x0d, 0x69, 0x32, 0xb5, 0x2c,
	0x8a, 0xf8, 0xe7, 0xae, 0x23, 0xdb, 0x63, 0xa4, 0x1f, 0x7f, 0xe2, 0x5c, 0xe4, 0x15, 0xdc, 0xec,
	0x61, 0xe2, 0xed, 0x56, 0x37, 0x19, 0xe8, 0xda, 0xda, 0x48, 0x95, 0x7b, 0xd9, 0x1b, 0x5d, 0x1a,
	0xf2, 0x14, 0x48, 0x5f, 0xb4, 0x55, 0xfc, 0xe5, 0x15, 0xc7, 0xdd, 0xd0, 0x56, 0x2f, 0x85, 0xbb,
	0x02, 0x0b, 0x57, 0x76, 0xd6, 0x97, 0x8d, 0x36, 0x96, 0x7b, 0xaa, 0x75, 0x81, 0x5c, 0x79, 0x5d,
	0xe2, 0x81, 0xdd, 0xf1, 0x91, 0x97, 0x70, 0xb0, 0x89, 0xce, 0x27, 0xf1, 0xdb, 0x27, 0xb1, 0xbf,
	0x8e, 0x6b, 0x52, 0x39, 0xfc, 0xb1, 0x0d, 0x7b, 0x0c, 0x6d, 0xa5, 0x95, 0x45, 0xf2, 0x0c, 0x46,
	0xed, 0x88, 0x43, 0xdf, 0xff, 0xd6, 0x1f, 0x96, 0xc7, 0x8f, 0xff, 0xb8, 0x79, 0x32, 0x1f, 0x48,
	0x3e, 0xc3, 0xb8, 0x19, 0x6e, 0xd6, 0x99, 0xae, 0x6b, 0x5e, 0xe4, 0xc0, 0x74, 0x00, 0x1e, 0xee,
	0x40, 0xea, 0xec, 0xf9, 0xca, 0x66, 0xb7, 0xcb, 0xfe, 0x0f, 0x72, 0x04, 0xbb, 0x61, 0xa9, 0x5c,
	0x5f, 0x1a, 0xc6, 0x07, 0x6b, 0x8c, 0x7e, 0xe5, 0x52, 0xff, 0x66, 0x97, 0xe1, 0xe4, 0x1d, 0x44,
	0x46, 0x9f, 0xc7, 0xd7, 0x5a, 0xd4, 0x0b, 0xfa, 0x9f, 0x17, 0x40, 0x2f, 0xdb, 0x40, 0x99, 0x3e,
	0x67, 0x0d, 0x43, 0x72, 0x1f, 0x22, 0xf7, 0x4d, 0xf6, 0x61, 0xc7, 0x59, 0xcd, 0x3e, 0x7e, 0x3f,
	0x71, 0x8d, 0x19, 0xb1, 0x91, 0x33, 0xe7, 0xe2, 0xcd, 0x07, 0x98, 0x48, 0x3d, 0xa0, 0x5f, 0x9d,
	0xe6, 0x97, 0xd9, 0xd5, 0x8f, 0xf6, 0x74, 0xa7, 0x3d, 0x8d, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x6b, 0xeb, 0x24, 0xf1, 0x03, 0x00, 0x00,
}
