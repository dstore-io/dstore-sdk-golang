// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_InsertIntoTrolley_Pu.proto
// DO NOT EDIT!

/*
Package om_InsertIntoTrolley_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_InsertIntoTrolley_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_InsertIntoTrolley_Pu

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
	UniqueId                      *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                  bool                        `protobuf:"varint,1001,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	NodeId                        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	NodeIdNull                    bool                        `protobuf:"varint,1002,opt,name=node_id_null,json=nodeIdNull" json:"node_id_null,omitempty"`
	IsTreeNodeId                  *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=is_tree_node_id,json=isTreeNodeId" json:"is_tree_node_id,omitempty"`
	IsTreeNodeIdNull              bool                        `protobuf:"varint,1003,opt,name=is_tree_node_id_null,json=isTreeNodeIdNull" json:"is_tree_node_id_null,omitempty"`
	Quantity                      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=quantity" json:"quantity,omitempty"`
	QuantityNull                  bool                        `protobuf:"varint,1004,opt,name=quantity_null,json=quantityNull" json:"quantity_null,omitempty"`
	InputNestLevelForContInfo     *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=input_nest_level_for_cont_info,json=inputNestLevelForContInfo" json:"input_nest_level_for_cont_info,omitempty"`
	InputNestLevelForContInfoNull bool                        `protobuf:"varint,1005,opt,name=input_nest_level_for_cont_info_null,json=inputNestLevelForContInfoNull" json:"input_nest_level_for_cont_info_null,omitempty"`
	Country                       *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=country" json:"country,omitempty"`
	CountryNull                   bool                        `protobuf:"varint,1006,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	BonusItemForItemSetId         *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=bonus_item_for_item_set_id,json=bonusItemForItemSetId" json:"bonus_item_for_item_set_id,omitempty"`
	BonusItemForItemSetIdNull     bool                        `protobuf:"varint,1007,opt,name=bonus_item_for_item_set_id_null,json=bonusItemForItemSetIdNull" json:"bonus_item_for_item_set_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Parameters) GetIsTreeNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsTreeNodeId
	}
	return nil
}

func (m *Parameters) GetQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Parameters) GetInputNestLevelForContInfo() *dstore_values.IntegerValue {
	if m != nil {
		return m.InputNestLevelForContInfo
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Parameters) GetBonusItemForItemSetId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BonusItemForItemSetId
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	InformationTypeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=information_type_id,json=informationTypeId" json:"information_type_id,omitempty"`
	HTreeNodeId       *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	NodeId            *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId        *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	ErrorCode         *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetInformationTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.InformationTypeId
	}
	return nil
}

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetErrorCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_InsertIntoTrolley_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_InsertIntoTrolley_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_InsertIntoTrolley_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x4e, 0x13, 0x4f,
	0x14, 0x0f, 0x94, 0x7e, 0x70, 0xe8, 0x1f, 0xf8, 0x2f, 0x6a, 0x4a, 0x89, 0xa8, 0x20, 0x89, 0x57,
	0x5b, 0x22, 0x2a, 0x26, 0x26, 0x26, 0x42, 0x30, 0x69, 0x84, 0x4a, 0x56, 0x42, 0xa2, 0x17, 0x4e,
	0x16, 0x7a, 0x2c, 0x9b, 0x6c, 0x67, 0xca, 0xcc, 0x2c, 0xa4, 0x6f, 0xe1, 0xe7, 0xe3, 0xf8, 0x32,
	0xde, 0xf9, 0xfd, 0x04, 0x26, 0x9e, 0x9d, 0xd9, 0xa5, 0xed, 0x0a, 0xb6, 0xde, 0xb4, 0x3b, 0x33,
	0xbf, 0x8f, 0xb3, 0x73, 0xf6, 0x37, 0x03, 0xf7, 0x9a, 0x4a, 0x0b, 0x89, 0x35, 0xe4, 0xad, 0x80,
	0x63, 0xad, 0x23, 0xc5, 0x21, 0x36, 0x23, 0x89, 0xaa, 0x26, 0xda, 0xac, 0xce, 0x15, 0x4a, 0x5d,
	0xe7, 0x5a, 0xec, 0x49, 0x11, 0x86, 0xd8, 0x65, 0xbb, 0x91, 0x4b, 0x10, 0x2d, 0x9c, 0x15, 0xcb,
	0x73, 0x2d, 0xcf, 0xbd, 0x00, 0x5c, 0x9d, 0x4b, 0xe4, 0x4f, 0xfc, 0x30, 0x42, 0x65, 0xb9, 0xd5,
	0xf9, 0x41, 0x4f, 0x94, 0x52, 0xc8, 0x64, 0x69, 0x61, 0x70, 0xa9, 0x8d, 0x4a, 0xf9, 0x2d, 0x4c,
	0x16, 0x97, 0xb3, 0x8b, 0xda, 0x0f, 0xf8, 0x2b, 0x21, 0xdb, 0xbe, 0x0e, 0x04, 0xb7, 0xa0, 0xa5,
	0x8f, 0x05, 0x80, 0x5d, 0x5f, 0xfa, 0xb4, 0x8a, 0x52, 0x39, 0xeb, 0x30, 0x19, 0xf1, 0xe0, 0x38,
	0x42, 0x16, 0x34, 0x2b, 0x63, 0xd7, 0xc7, 0x6e, 0x4d, 0xdd, 0xae, 0xba, 0x49, 0xed, 0x49, 0x51,
	0x4a, 0xcb, 0x80, 0xb7, 0xf6, 0xe3, 0x81, 0x57, 0xb2, 0xe0, 0x7a, 0xd3, 0x59, 0x81, 0xe9, 0x33,
	0x22, 0xe3, 0x51, 0x18, 0x56, 0x3e, 0x17, 0x89, 0x5e, 0xf2, 0xca, 0x29, 0xa4, 0x41, 0x93, 0xce,
	0x1d, 0x28, 0x72, 0xd1, 0x34, 0xea, 0xe3, 0x46, 0x7d, 0x21, 0xa3, 0x1e, 0x70, 0x8d, 0x2d, 0x94,
	0x56, 0xbe, 0x10, 0x63, 0x49, 0xfc, 0x06, 0x94, 0x13, 0x96, 0x95, 0xfe, 0x62, 0xa5, 0xc1, 0x2e,
	0x1b, 0xe1, 0x0d, 0x98, 0x09, 0x14, 0xd3, 0x12, 0x91, 0xa5, 0x06, 0xb9, 0x73, 0x0d, 0x0e, 0x84,
	0x08, 0xd1, 0xe7, 0xd6, 0xa0, 0x1c, 0xa8, 0x3d, 0xa2, 0x34, 0xac, 0x4d, 0x0d, 0x2e, 0x65, 0x34,
	0xac, 0xdd, 0x57, 0x6b, 0x37, 0xdb, 0x0f, 0x36, 0xa6, 0xeb, 0x50, 0x3a, 0x8e, 0x7c, 0xae, 0x03,
	0xdd, 0xad, 0x4c, 0x0c, 0x7f, 0x9d, 0x33, 0xb0, 0x73, 0x13, 0xfe, 0x4b, 0x9f, 0xad, 0xc5, 0xb7,
	0x64, 0xb3, 0xd2, 0x59, 0x23, 0xff, 0x12, 0x16, 0x03, 0xde, 0x89, 0x34, 0xe3, 0xa8, 0x34, 0x0b,
	0xf1, 0x04, 0x43, 0x46, 0xfd, 0x63, 0x87, 0x82, 0x6b, 0x16, 0xb7, 0xb2, 0x92, 0x1f, 0x6e, 0x3a,
	0x6f, 0x24, 0x1a, 0xa4, 0xb0, 0x1d, 0x0b, 0x3c, 0x16, 0x72, 0x93, 0xe8, 0x75, 0x62, 0x3b, 0x4f,
	0x60, 0xf9, 0xef, 0xfa, 0xb6, 0xb6, 0xef, 0xb6, 0xb6, 0xab, 0x17, 0x0a, 0xa5, 0x9d, 0x3d, 0x14,
	0x11, 0xd7, 0xb2, 0x5b, 0x29, 0x0c, 0xfd, 0x6e, 0x52, 0xa8, 0xb3, 0x04, 0xe5, 0xe4, 0xd1, 0x7a,
	0xfd, 0xb0, 0x5e, 0x53, 0xc9, 0xa4, 0x51, 0xde, 0x87, 0xea, 0x81, 0xe0, 0x91, 0x62, 0x81, 0xc6,
	0xb6, 0x29, 0xd0, 0x3c, 0x28, 0xd4, 0x71, 0x97, 0x8b, 0xc3, 0xb7, 0xe0, 0xb2, 0xa1, 0xd7, 0x89,
	0x44, 0x45, 0xc7, 0x7f, 0xcf, 0x50, 0x53, 0xbb, 0x37, 0xe1, 0xda, 0xc5, 0xba, 0xb6, 0x9c, 0x9f,
	0xb6, 0x9c, 0xf9, 0x73, 0x05, 0xe2, 0xe2, 0x96, 0x7e, 0x4d, 0x40, 0xc9, 0x43, 0xd5, 0x11, 0x14,
	0x67, 0x67, 0x15, 0xf2, 0x26, 0x9d, 0xd9, 0xe4, 0x24, 0xa9, 0xb7, 0xc9, 0xdd, 0x8a, 0x7f, 0x3d,
	0x0b, 0x74, 0x9e, 0xc3, 0x6c, 0x9c, 0x4b, 0xd6, 0x17, 0x4c, 0x0a, 0x46, 0x8e, 0xc8, 0x6e, 0x86,
	0x9c, 0x8d, 0xef, 0x0e, 0x8d, 0xeb, 0xbd, 0xb1, 0x37, 0xd3, 0x1e, 0x9c, 0x70, 0xee, 0x43, 0x31,
	0x39, 0x0f, 0x28, 0x09, 0xb1, 0xe2, 0xe2, 0x1f, 0x8a, 0xf6, 0xb4, 0xd8, 0xb1, 0xff, 0x5e, 0x0a,
	0x77, 0xb6, 0x20, 0x27, 0xc5, 0x29, 0x7d, 0xd1, 0x31, 0x6b, 0xcd, 0x1d, 0xe9, 0xe8, 0x72, 0xd3,
	0x4d, 0x70, 0x3d, 0x71, 0xea, 0xc5, 0xfc, 0xea, 0xa7, 0x71, 0xc8, 0xd1, 0xc0, 0xb9, 0x02, 0x05,
	0x1a, 0xc6, 0xbd, 0x7a, 0xdd, 0xa0, 0x7d, 0xc9, 0x7b, 0x79, 0x1a, 0xd2, 0xfe, 0x6f, 0xc3, 0x5c,
	0xdf, 0x0b, 0x31, 0xdd, 0xed, 0x98, 0xd8, 0xbe, 0x69, 0x0c, 0xef, 0xe8, 0xff, 0x7d, 0xc4, 0x3d,
	0xe2, 0x91, 0xda, 0x23, 0x98, 0x3e, 0x1a, 0xcc, 0xff, 0xdb, 0x11, 0x84, 0xa6, 0x8e, 0xfa, 0xf2,
	0x7f, 0xb7, 0x77, 0x38, 0xbd, 0x6b, 0x8c, 0x7e, 0x3a, 0x3d, 0x84, 0xf2, 0x80, 0xef, 0xfb, 0x11,
	0xb8, 0xa0, 0x7b, 0xb6, 0x0f, 0x00, 0xcc, 0xc7, 0x40, 0xb9, 0x6b, 0x62, 0xe5, 0xc3, 0x08, 0xec,
	0x49, 0x83, 0xdf, 0x24, 0xf8, 0xc6, 0x53, 0x58, 0x08, 0x44, 0xa6, 0x45, 0xbd, 0x5b, 0xe9, 0xc5,
	0xea, 0xbf, 0xde, 0x57, 0x07, 0x05, 0x73, 0x2f, 0xac, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0x3f, 0xb2, 0xe2, 0xea, 0x06, 0x00, 0x00,
}