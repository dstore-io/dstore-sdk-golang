// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampaigns_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampaigns_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampaigns_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampaigns_Ad

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
	CampaignName            *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=campaign_name,json=campaignName" json:"campaign_name,omitempty"`
	CampaignNameNull        bool                        `protobuf:"varint,1001,opt,name=campaign_name_null,json=campaignNameNull" json:"campaign_name_null,omitempty"`
	CampaignDescription     *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=campaign_description,json=campaignDescription" json:"campaign_description,omitempty"`
	CampaignDescriptionNull bool                        `protobuf:"varint,1002,opt,name=campaign_description_null,json=campaignDescriptionNull" json:"campaign_description_null,omitempty"`
	CampaignTypeId          *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=campaign_type_id,json=campaignTypeId" json:"campaign_type_id,omitempty"`
	CampaignTypeIdNull      bool                        `protobuf:"varint,1003,opt,name=campaign_type_id_null,json=campaignTypeIdNull" json:"campaign_type_id_null,omitempty"`
	Active                  *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=active" json:"active,omitempty"`
	ActiveNull              bool                        `protobuf:"varint,1004,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
	DeleteCampaign          *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_campaign,json=deleteCampaign" json:"delete_campaign,omitempty"`
	DeleteCampaignNull      bool                        `protobuf:"varint,1005,opt,name=delete_campaign_null,json=deleteCampaignNull" json:"delete_campaign_null,omitempty"`
	ForceDelete             *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=force_delete,json=forceDelete" json:"force_delete,omitempty"`
	ForceDeleteNull         bool                        `protobuf:"varint,1006,opt,name=force_delete_null,json=forceDeleteNull" json:"force_delete_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetCampaignName() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignName
	}
	return nil
}

func (m *Parameters) GetCampaignDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CampaignDescription
	}
	return nil
}

func (m *Parameters) GetCampaignTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignTypeId
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Parameters) GetDeleteCampaign() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteCampaign
	}
	return nil
}

func (m *Parameters) GetForceDelete() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForceDelete
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	CampaignId      *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=campaign_id,json=campaignId" json:"campaign_id,omitempty"`
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

func (m *Response) GetCampaignId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CampaignId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampaigns_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampaigns_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampaigns_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x1b, 0x92, 0x56, 0x93, 0x42, 0x8a, 0x5b, 0x60, 0x9b, 0x08, 0x54, 0xb5, 0x17, 0x24,
	0xc4, 0x86, 0xb6, 0x1c, 0x90, 0x40, 0x20, 0x68, 0x7a, 0xc8, 0x21, 0x11, 0x5a, 0x21, 0x24, 0xb8,
	0xac, 0xdc, 0x78, 0x1a, 0x59, 0x4a, 0xec, 0xc8, 0xde, 0xb4, 0xea, 0x5b, 0xf0, 0x48, 0xbc, 0x0e,
	0xbf, 0x17, 0x5e, 0x00, 0xef, 0x8e, 0x37, 0xc9, 0x2e, 0x91, 0x42, 0x2f, 0x49, 0xec, 0xf9, 0xfe,
	0x1c, 0xcf, 0x18, 0x9e, 0x0b, 0x9b, 0x68, 0x83, 0x6d, 0x54, 0x43, 0xa9, 0xb0, 0x3d, 0x31, 0x7a,
	0x80, 0x62, 0x6a, 0xd0, 0xb6, 0xf5, 0x38, 0xee, 0x69, 0x21, 0x2f, 0xae, 0x4f, 0xf9, 0x78, 0xc2,
	0xe5, 0x50, 0xd9, 0xf8, 0xad, 0x08, 0x1d, 0x20, 0xd1, 0xec, 0x90, 0x58, 0x21, 0xb1, 0xc2, 0xa5,
	0xd0, 0xe6, 0x8e, 0x97, 0xbe, 0xe4, 0xa3, 0x29, 0x5a, 0x62, 0x36, 0xf7, 0x8a, 0x7e, 0x68, 0x8c,
	0x36, 0xbe, 0xd4, 0x2a, 0x96, 0xc6, 0x68, 0x2d, 0x1f, 0xa2, 0x2f, 0x1e, 0x96, 0x8b, 0x09, 0x97,
	0xea, 0x42, 0x9b, 0x31, 0x4f, 0xa4, 0x56, 0x04, 0x3a, 0xf8, 0x5a, 0x05, 0x78, 0xcf, 0x0d, 0x77,
	0x55, 0x34, 0x96, 0xbd, 0x81, 0xdb, 0x03, 0x1f, 0x28, 0x56, 0x6e, 0x33, 0x58, 0xdb, 0x5f, 0x7b,
	0x5c, 0x3f, 0x6e, 0x86, 0x3e, 0xbd, 0x0f, 0x66, 0x13, 0x23, 0xd5, 0xf0, 0x63, 0xba, 0x88, 0xb6,
	0x72, 0x42, 0xdf, 0xe1, 0xd9, 0x53, 0x60, 0x05, 0x81, 0x58, 0x4d, 0x47, 0xa3, 0xe0, 0xdb, 0x86,
	0x93, 0xd9, 0x8c, 0xb6, 0x17, 0xa1, 0x7d, 0x57, 0x60, 0x3d, 0xd8, 0x9d, 0xc1, 0x05, 0xda, 0x81,
	0x91, 0x93, 0x34, 0x5c, 0xb0, 0xbe, 0xd2, 0x76, 0x27, 0xe7, 0x75, 0xe6, 0x34, 0xf6, 0x12, 0xf6,
	0x96, 0xc9, 0x51, 0x88, 0xef, 0x14, 0xe2, 0xc1, 0x12, 0x62, 0x96, 0xe5, 0x0c, 0x66, 0xf9, 0xe2,
	0xe4, 0x7a, 0x82, 0xb1, 0x14, 0x41, 0x25, 0xcb, 0xd1, 0x2a, 0xe5, 0x90, 0x2a, 0xc1, 0x21, 0x1a,
	0x0a, 0x72, 0x27, 0x27, 0x7d, 0x70, 0x9c, 0xae, 0x60, 0xc7, 0x70, 0xaf, 0x2c, 0x43, 0xfe, 0x3f,
	0xc8, 0x9f, 0x15, 0xf1, 0x99, 0xf5, 0x09, 0xd4, 0xf8, 0x20, 0x91, 0x97, 0x18, 0xdc, 0x5a, 0x6d,
	0xe8, 0xa1, 0x6c, 0x1f, 0xea, 0xf4, 0x8b, 0xe4, 0x7f, 0x92, 0x3c, 0xd0, 0x5e, 0x26, 0xdb, 0x81,
	0x86, 0xc0, 0x91, 0xbb, 0xd9, 0x38, 0xf7, 0x0c, 0xaa, 0x4b, 0xf5, 0xcf, 0xb5, 0x1e, 0x21, 0x57,
	0xfe, 0x40, 0xc4, 0xc9, 0x1b, 0x93, 0x1d, 0xc1, 0x6e, 0x49, 0x85, 0x0c, 0x7f, 0xf9, 0xf3, 0x14,
	0xe1, 0x99, 0xf1, 0x6b, 0xd8, 0x72, 0x8d, 0x36, 0xc0, 0x98, 0x6a, 0x41, 0x6d, 0xf5, 0xa9, 0xea,
	0x19, 0xa1, 0x93, 0xe1, 0xd9, 0x13, 0xb8, 0xbb, 0xc8, 0x27, 0xbf, 0xdf, 0xe4, 0xd7, 0x58, 0x00,
	0xa6, 0x66, 0x07, 0x7f, 0xd6, 0x61, 0x33, 0x42, 0x3b, 0xd1, 0xca, 0x22, 0x7b, 0x06, 0xd5, 0x6c,
	0x40, 0xca, 0x8d, 0xeb, 0xc7, 0x8e, 0x86, 0xe7, 0x2c, 0xfd, 0x8c, 0x08, 0xc8, 0x3e, 0xc1, 0x76,
	0x3a, 0x1a, 0xf1, 0xc2, 0x6c, 0xb8, 0xf6, 0xab, 0x38, 0x72, 0x58, 0x22, 0x97, 0x27, 0xa8, 0xe7,
	0xd6, 0xdd, 0xf9, 0x3a, 0x6a, 0x8c, 0x8b, 0x1b, 0xec, 0x05, 0x6c, 0xf8, 0x91, 0x74, 0x8d, 0x94,
	0x2a, 0x3e, 0xfa, 0x47, 0x91, 0x06, 0xb6, 0x47, 0xdf, 0x51, 0x0e, 0x67, 0xa7, 0x50, 0x31, 0xfa,
	0xca, 0x75, 0x43, 0xca, 0x3a, 0x0a, 0xff, 0xe3, 0xed, 0x08, 0xf3, 0xbf, 0x20, 0x8c, 0xf4, 0x55,
	0x94, 0xb2, 0xd9, 0x2b, 0xa8, 0xcf, 0x6e, 0xcc, 0xf5, 0x32, 0xae, 0xbe, 0x04, 0xc8, 0xf1, 0x5d,
	0xd1, 0x7c, 0x08, 0x15, 0xa7, 0xc4, 0xee, 0x43, 0xcd, 0x69, 0xa5, 0xfc, 0x2f, 0x7d, 0x27, 0x50,
	0x8d, 0xaa, 0x6e, 0xd9, 0x15, 0xef, 0x7a, 0xd0, 0x92, 0xba, 0x14, 0x6c, 0xfe, 0x14, 0x7e, 0x0e,
	0x6f, 0xf6, 0x48, 0x9e, 0xd7, 0xb2, 0xe7, 0xe8, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x5e, 0x02, 0x9b, 0x5d, 0x05, 0x00, 0x00,
}