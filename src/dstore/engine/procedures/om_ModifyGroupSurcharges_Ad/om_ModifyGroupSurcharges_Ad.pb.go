// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyGroupSurcharges_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyGroupSurcharges_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyGroupSurcharges_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyGroupSurcharges_Ad

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
	GroupId             *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	GroupIdNull         bool                        `protobuf:"varint,1001,opt,name=group_id_null,json=groupIdNull" json:"group_id_null,omitempty"`
	TreeNodeId          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull      bool                        `protobuf:"varint,1002,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	SurchargeTypeId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	SurchargeTypeIdNull bool                        `protobuf:"varint,1003,opt,name=surcharge_type_id_null,json=surchargeTypeIdNull" json:"surcharge_type_id_null,omitempty"`
	SurchargeValue      *dstore_values.DecimalValue `protobuf:"bytes,4,opt,name=surcharge_value,json=surchargeValue" json:"surcharge_value,omitempty"`
	SurchargeValueNull  bool                        `protobuf:"varint,1004,opt,name=surcharge_value_null,json=surchargeValueNull" json:"surcharge_value_null,omitempty"`
	DeleteSurcharge     *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_surcharge,json=deleteSurcharge" json:"delete_surcharge,omitempty"`
	DeleteSurchargeNull bool                        `protobuf:"varint,1005,opt,name=delete_surcharge_null,json=deleteSurchargeNull" json:"delete_surcharge_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *Parameters) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Parameters) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Parameters) GetSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.SurchargeValue
	}
	return nil
}

func (m *Parameters) GetDeleteSurcharge() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteSurcharge
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyGroupSurcharges_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyGroupSurcharges_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyGroupSurcharges_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xd5, 0xa6, 0x69, 0xa2, 0x53, 0xe8, 0x65, 0x0a, 0x55, 0x48, 0x04, 0x42, 0xed, 0x06,
	0x58, 0x38, 0x40, 0x2b, 0x40, 0x48, 0x2c, 0x40, 0x40, 0x95, 0x45, 0x22, 0xe4, 0x22, 0x24, 0xd8,
	0x58, 0x6e, 0xe6, 0xd4, 0x58, 0xb2, 0x67, 0xa2, 0x19, 0x9b, 0xaa, 0x6f, 0xc1, 0xfb, 0xf0, 0x04,
	0x3c, 0x0a, 0xb7, 0x77, 0xe0, 0xcc, 0x25, 0x49, 0xed, 0x56, 0x04, 0x75, 0x93, 0xf8, 0xf8, 0xfc,
	0xff, 0xf9, 0xc6, 0xe3, 0x7f, 0x0c, 0xcf, 0xb9, 0x2e, 0xa4, 0xc2, 0x3e, 0x8a, 0x24, 0x15, 0xd8,
	0x9f, 0x28, 0x39, 0x46, 0x5e, 0x2a, 0xd4, 0x7d, 0x99, 0x47, 0x43, 0xc9, 0xd3, 0x93, 0xb3, 0x43,
	0x25, 0xcb, 0xc9, 0x51, 0xa9, 0xc6, 0x9f, 0x63, 0x95, 0xa0, 0x8e, 0x5e, 0xf2, 0x80, 0x64, 0x85,
	0x64, 0xf7, 0x9d, 0x37, 0x70, 0xde, 0xe0, 0x1f, 0x86, 0xee, 0xb6, 0xc7, 0x7c, 0x89, 0xb3, 0x12,
	0xb5, 0xf3, 0x77, 0x6f, 0x55, 0xd9, 0xa8, 0x94, 0x54, 0xbe, 0xd5, 0xab, 0xb6, 0x72, 0xd4, 0x3a,
	0x4e, 0xd0, 0x37, 0xf7, 0xea, 0xcd, 0x22, 0x4e, 0xc5, 0x89, 0x54, 0x79, 0x5c, 0xa4, 0x52, 0x38,
	0xd1, 0xee, 0xf7, 0x15, 0x80, 0x77, 0xb1, 0x8a, 0xa9, 0x8b, 0x4a, 0xb3, 0x27, 0xd0, 0x4e, 0xcc,
	0xb2, 0xa2, 0x94, 0x77, 0x96, 0xee, 0x2e, 0xdd, 0x5b, 0x7b, 0xdc, 0x0b, 0xfc, 0xf2, 0xfd, 0x9a,
	0x52, 0x51, 0x60, 0x82, 0xea, 0x83, 0xa9, 0xc2, 0x96, 0x15, 0x0f, 0x38, 0xdb, 0x83, 0xeb, 0x53,
	0x5f, 0x24, 0xca, 0x2c, 0xeb, 0xfc, 0x68, 0x91, 0xbb, 0x1d, 0xae, 0x79, 0xc1, 0x88, 0xee, 0xb1,
	0x17, 0x70, 0xad, 0x50, 0x88, 0x91, 0x90, 0x1c, 0x0d, 0x60, 0x79, 0x31, 0x00, 0x8c, 0x61, 0x44,
	0x7a, 0x62, 0x3c, 0x80, 0xad, 0xf3, 0x76, 0xc7, 0xf9, 0xe9, 0x38, 0xeb, 0x73, 0x9d, 0x45, 0x1d,
	0xc2, 0x96, 0x9e, 0xee, 0x6c, 0x54, 0x9c, 0x4d, 0x2c, 0xaf, 0xb1, 0x98, 0xb7, 0x31, 0x73, 0xbd,
	0x27, 0x13, 0x41, 0x0f, 0x60, 0xe7, 0xc2, 0x20, 0x47, 0xfe, 0xe5, 0xc8, 0xdb, 0x35, 0x87, 0xc5,
	0xbf, 0x86, 0xf9, 0xa0, 0xc8, 0x72, 0x3a, 0x2b, 0x97, 0xc2, 0x39, 0x8e, 0xd3, 0x3c, 0xce, 0x1c,
	0x7c, 0x7d, 0xe6, 0xb1, 0x35, 0x7b, 0x04, 0x37, 0x6a, 0x53, 0x1c, 0xf9, 0xb7, 0x23, 0xb3, 0xaa,
	0xdc, 0x82, 0xdf, 0xc2, 0x26, 0xc7, 0x8c, 0xde, 0x65, 0x34, 0x6b, 0x76, 0x9a, 0x97, 0x92, 0x8f,
	0xa5, 0xcc, 0x30, 0x16, 0xfe, 0xb1, 0x9d, 0x69, 0x16, 0x46, 0xb6, 0x0f, 0x37, 0xeb, 0x73, 0x1c,
	0xfb, 0x8f, 0x7f, 0xea, 0x9a, 0xc1, 0xc0, 0x77, 0xbf, 0x2d, 0x43, 0x3b, 0x44, 0x3d, 0x91, 0x42,
	0x23, 0x7b, 0x08, 0x4d, 0x9b, 0x54, 0x1f, 0xa3, 0x6e, 0x50, 0x3d, 0x05, 0x2e, 0xc5, 0x6f, 0xcc,
	0x6f, 0xe8, 0x84, 0xec, 0x23, 0x6c, 0x9a, 0x8c, 0x46, 0xe7, 0x42, 0x4a, 0x11, 0x69, 0x90, 0x39,
	0xa8, 0x99, 0xeb, 0x51, 0x1e, 0x52, 0x3d, 0x98, 0xd7, 0xe1, 0x46, 0x5e, 0xbd, 0xc1, 0x9e, 0x41,
	0xcb, 0x9f, 0x0d, 0x0a, 0x81, 0x99, 0x78, 0xe7, 0xc2, 0x44, 0x77, 0x72, 0x86, 0xee, 0x3f, 0x9c,
	0xca, 0xd9, 0x00, 0x1a, 0x4a, 0x9e, 0xd2, 0xdb, 0x33, 0xae, 0xa7, 0xc1, 0x7f, 0x1f, 0xe5, 0x60,
	0xba, 0x11, 0x41, 0x28, 0x4f, 0x43, 0x33, 0xa3, 0x7b, 0x1b, 0x1a, 0x74, 0xcd, 0x76, 0x60, 0x95,
	0x2a, 0x93, 0xc7, 0xaf, 0x23, 0xda, 0x9a, 0x66, 0xd8, 0xa4, 0x72, 0xc0, 0x5f, 0x1d, 0x41, 0x2f,
	0x95, 0x35, 0xc0, 0xfc, 0x3b, 0xf3, 0xe9, 0xe0, 0x2a, 0x5f, 0xa0, 0xe3, 0x55, 0x7b, 0xca, 0xf7,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x16, 0x7c, 0x75, 0x90, 0xc0, 0x04, 0x00, 0x00,
}
