// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetNodePaymentForShip_Ad.proto
// DO NOT EDIT!

/*
Package om_GetNodePaymentForShip_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetNodePaymentForShip_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetNodePaymentForShip_Ad

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
	TreeNodeId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
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
	RowId                         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	HideWhenOrderedAlone          *dstore_values.BooleanValue `protobuf:"bytes,10001,opt,name=hide_when_ordered_alone,json=hideWhenOrderedAlone" json:"hide_when_ordered_alone,omitempty"`
	NodeDescription               *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	PaymentForShippingId          *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=payment_for_shipping_id,json=paymentForShippingId" json:"payment_for_shipping_id,omitempty"`
	DescriptionForAdmin           *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=description_for_admin,json=descriptionForAdmin" json:"description_for_admin,omitempty"`
	TreeNodeId                    *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                        *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Always                        *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=always" json:"always,omitempty"`
	PaymentForShippingDescription *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=payment_for_shipping_description,json=paymentForShippingDescription" json:"payment_for_shipping_description,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetHideWhenOrderedAlone() *dstore_values.BooleanValue {
	if m != nil {
		return m.HideWhenOrderedAlone
	}
	return nil
}

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetPaymentForShippingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentForShippingId
	}
	return nil
}

func (m *Response_Row) GetDescriptionForAdmin() *dstore_values.StringValue {
	if m != nil {
		return m.DescriptionForAdmin
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetAlways() *dstore_values.BooleanValue {
	if m != nil {
		return m.Always
	}
	return nil
}

func (m *Response_Row) GetPaymentForShippingDescription() *dstore_values.StringValue {
	if m != nil {
		return m.PaymentForShippingDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetNodePaymentForShip_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetNodePaymentForShip_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetNodePaymentForShip_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetNodePaymentForShip_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0xa9, 0x69, 0x93, 0x32, 0x8a, 0xad, 0xdb, 0xaa, 0x67, 0x82, 0x12, 0xea, 0x8b, 0xfa,
	0x70, 0x11, 0x7f, 0xa0, 0x08, 0x15, 0x22, 0x5a, 0xc9, 0x43, 0xaf, 0xe5, 0x0a, 0x8a, 0x22, 0x1c,
	0xdb, 0xee, 0x98, 0x2c, 0xde, 0xed, 0x1e, 0xbb, 0x17, 0x43, 0xff, 0x0b, 0xad, 0x3f, 0xff, 0x45,
	0x9f, 0xfd, 0x07, 0xdc, 0xbd, 0xbd, 0x90, 0xbb, 0x6b, 0x31, 0xe9, 0x4b, 0xc2, 0xdc, 0xcc, 0x67,
	0x76, 0xbe, 0xdf, 0x59, 0x16, 0x9e, 0x31, 0x9d, 0x49, 0x85, 0x3d, 0x14, 0x43, 0x2e, 0xb0, 0x97,
	0x2a, 0x79, 0x84, 0x6c, 0xac, 0x50, 0xf7, 0x64, 0x12, 0xbd, 0xc6, 0x2c, 0x90, 0x0c, 0xf7, 0xe9,
	0x71, 0x82, 0x22, 0xdb, 0x91, 0xea, 0x60, 0xc4, 0xd3, 0xa8, 0xcf, 0x7c, 0x53, 0x96, 0x49, 0x72,
	0xd7, 0xb1, 0xbe, 0x63, 0xfd, 0xff, 0x00, 0xed, 0x8d, 0xe2, 0x98, 0xcf, 0x34, 0x1e, 0xa3, 0x76,
	0x7c, 0xfb, 0x46, 0xf5, 0x6c, 0x54, 0x4a, 0xaa, 0x22, 0xd5, 0xa9, 0xa6, 0x12, 0xd4, 0x9a, 0x0e,
	0xb1, 0x48, 0xde, 0xae, 0x27, 0x33, 0xca, 0xc5, 0x47, 0xa9, 0x12, 0x9a, 0x71, 0x29, 0x5c, 0xd1,
	0xd6, 0x04, 0x60, 0x9f, 0x2a, 0x6a, 0x92, 0xa8, 0x34, 0xd9, 0x86, 0x4b, 0x99, 0x42, 0x8c, 0x84,
	0x99, 0x2e, 0xe2, 0xcc, 0x5b, 0xea, 0x2e, 0xdd, 0xb9, 0xf8, 0xa0, 0xe3, 0x17, 0x0a, 0x8a, 0xb1,
	0xb8, 0xc8, 0x70, 0x88, 0xea, 0x8d, 0x8d, 0x42, 0xb0, 0x80, 0x55, 0x33, 0x60, 0xe4, 0x1e, 0x5c,
	0x29, 0xe3, 0x91, 0x18, 0xc7, 0xb1, 0xf7, 0xa7, 0x65, 0x9a, 0xac, 0x86, 0x97, 0x67, 0x75, 0x81,
	0xf9, 0xbc, 0x75, 0xd2, 0x82, 0xd5, 0x10, 0x75, 0x2a, 0x85, 0x46, 0x72, 0x1f, 0x56, 0x72, 0x59,
	0xc5, 0x81, 0x6d, 0xbf, 0x6a, 0x99, 0x93, 0xfc, 0xca, 0xfe, 0x86, 0xae, 0x90, 0xbc, 0x83, 0x75,
	0x2b, 0x28, 0x2a, 0x29, 0xf2, 0x2e, 0x74, 0x1b, 0x06, 0xf6, 0x6b, 0x70, 0x5d, 0xf7, 0xae, 0x89,
	0x07, 0xb3, 0x38, 0x5c, 0x4b, 0xaa, 0x1f, 0xc8, 0x53, 0x68, 0x15, 0x46, 0x7a, 0x8d, 0xbc, 0xe3,
	0xad, 0x53, 0x1d, 0x9d, 0xcd, 0xbb, 0xee, 0x3f, 0x9c, 0x96, 0x93, 0x01, 0x34, 0x94, 0x9c, 0x78,
	0xcb, 0x39, 0xf5, 0xc4, 0x5f, 0x78, 0xef, 0xfe, 0xd4, 0x08, 0x3f, 0x94, 0x93, 0xd0, 0xf6, 0x68,
	0xff, 0x5d, 0x86, 0x86, 0x09, 0xc8, 0x35, 0x68, 0x9a, 0xd0, 0xee, 0xe2, 0x4b, 0x60, 0xbc, 0x59,
	0x09, 0x57, 0x4c, 0x68, 0xac, 0x3e, 0x80, 0xeb, 0x23, 0x6e, 0x5c, 0x9e, 0x8c, 0x50, 0x44, 0x52,
	0x31, 0x54, 0xc8, 0x22, 0x1a, 0x4b, 0x81, 0xde, 0xd7, 0xe0, 0xcc, 0xad, 0x1d, 0x4a, 0x19, 0x23,
	0x15, 0x6e, 0x6b, 0x9b, 0x16, 0x7e, 0x6b, 0xd8, 0x3d, 0x87, 0xf6, 0x2d, 0x49, 0x76, 0x60, 0x3d,
	0x5f, 0x1d, 0x43, 0x7d, 0xa4, 0x78, 0x9a, 0x9b, 0x7a, 0x12, 0x54, 0x57, 0x52, 0x74, 0xd3, 0x99,
	0xe2, 0x62, 0xe8, 0x9a, 0xad, 0x59, 0xe8, 0xe5, 0x8c, 0xb1, 0xc3, 0xa5, 0x4e, 0x63, 0x64, 0x6c,
	0x8d, 0xb4, 0x51, 0x99, 0x9a, 0x6a, 0xab, 0xe2, 0x5b, 0x30, 0xff, 0x4a, 0x6d, 0xa6, 0x15, 0x83,
	0x2c, 0x6a, 0x14, 0xef, 0xc1, 0xd5, 0xd2, 0x5c, 0x79, 0x63, 0xca, 0x12, 0x2e, 0xbc, 0xef, 0xf3,
	0x27, 0xdc, 0x28, 0x91, 0xa6, 0x6b, 0xdf, 0x72, 0xe4, 0x79, 0xed, 0xb2, 0xff, 0x08, 0xce, 0x77,
	0xdb, 0x1f, 0x43, 0x6b, 0x8a, 0xfe, 0x5c, 0x00, 0x6d, 0x0a, 0x87, 0x3d, 0x82, 0x26, 0x8d, 0x27,
	0xf4, 0x58, 0x7b, 0xbf, 0x16, 0x58, 0x54, 0x51, 0x4b, 0x18, 0x74, 0xcf, 0xb4, 0xb4, 0xbc, 0xaa,
	0xdf, 0xf3, 0x8d, 0xb8, 0x79, 0xda, 0xda, 0xd2, 0xe2, 0x5e, 0x7c, 0x80, 0x0e, 0x97, 0xb5, 0x7b,
	0x3b, 0x7b, 0xeb, 0xde, 0x6f, 0x0f, 0xa5, 0x66, 0x9f, 0xa6, 0x79, 0x76, 0xce, 0xe7, 0xf0, 0xb0,
	0x99, 0x3f, 0x39, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x45, 0x4b, 0xcf, 0x4d, 0x05,
	0x00, 0x00,
}