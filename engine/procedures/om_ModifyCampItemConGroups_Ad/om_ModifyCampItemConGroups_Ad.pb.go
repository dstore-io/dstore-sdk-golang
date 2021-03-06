// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_ModifyCampItemConGroups_Ad.proto
// DO NOT EDIT!

/*
Package om_ModifyCampItemConGroups_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_ModifyCampItemConGroups_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_ModifyCampItemConGroups_Ad

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
	ConditionId                   *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=condition_id,json=conditionId" json:"condition_id,omitempty"`
	ConditionIdNull               bool                        `protobuf:"varint,1001,opt,name=condition_id_null,json=conditionIdNull" json:"condition_id_null,omitempty"`
	ItemConditionGroupId          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=item_condition_group_id,json=itemConditionGroupId" json:"item_condition_group_id,omitempty"`
	ItemConditionGroupIdNull      bool                        `protobuf:"varint,1002,opt,name=item_condition_group_id_null,json=itemConditionGroupIdNull" json:"item_condition_group_id_null,omitempty"`
	ConditionGroupDescription     *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=condition_group_description,json=conditionGroupDescription" json:"condition_group_description,omitempty"`
	ConditionGroupDescriptionNull bool                        `protobuf:"varint,1003,opt,name=condition_group_description_null,json=conditionGroupDescriptionNull" json:"condition_group_description_null,omitempty"`
	MinNumberOfItems              *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=min_number_of_items,json=minNumberOfItems" json:"min_number_of_items,omitempty"`
	MinNumberOfItemsNull          bool                        `protobuf:"varint,1004,opt,name=min_number_of_items_null,json=minNumberOfItemsNull" json:"min_number_of_items_null,omitempty"`
	MaxNumberOfItems              *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=max_number_of_items,json=maxNumberOfItems" json:"max_number_of_items,omitempty"`
	MaxNumberOfItemsNull          bool                        `protobuf:"varint,1005,opt,name=max_number_of_items_null,json=maxNumberOfItemsNull" json:"max_number_of_items_null,omitempty"`
	FromQuantity                  *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=from_quantity,json=fromQuantity" json:"from_quantity,omitempty"`
	FromQuantityNull              bool                        `protobuf:"varint,1006,opt,name=from_quantity_null,json=fromQuantityNull" json:"from_quantity_null,omitempty"`
	ToQuantity                    *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=to_quantity,json=toQuantity" json:"to_quantity,omitempty"`
	ToQuantityNull                bool                        `protobuf:"varint,1007,opt,name=to_quantity_null,json=toQuantityNull" json:"to_quantity_null,omitempty"`
	FromItemBasicPrice            *dstore_values.DecimalValue `protobuf:"bytes,8,opt,name=from_item_basic_price,json=fromItemBasicPrice" json:"from_item_basic_price,omitempty"`
	FromItemBasicPriceNull        bool                        `protobuf:"varint,1008,opt,name=from_item_basic_price_null,json=fromItemBasicPriceNull" json:"from_item_basic_price_null,omitempty"`
	ToItemBasicPrice              *dstore_values.DecimalValue `protobuf:"bytes,9,opt,name=to_item_basic_price,json=toItemBasicPrice" json:"to_item_basic_price,omitempty"`
	ToItemBasicPriceNull          bool                        `protobuf:"varint,1009,opt,name=to_item_basic_price_null,json=toItemBasicPriceNull" json:"to_item_basic_price_null,omitempty"`
	FromBasicPriceSum             *dstore_values.DecimalValue `protobuf:"bytes,10,opt,name=from_basic_price_sum,json=fromBasicPriceSum" json:"from_basic_price_sum,omitempty"`
	FromBasicPriceSumNull         bool                        `protobuf:"varint,1010,opt,name=from_basic_price_sum_null,json=fromBasicPriceSumNull" json:"from_basic_price_sum_null,omitempty"`
	ToBasicPriceSum               *dstore_values.DecimalValue `protobuf:"bytes,11,opt,name=to_basic_price_sum,json=toBasicPriceSum" json:"to_basic_price_sum,omitempty"`
	ToBasicPriceSumNull           bool                        `protobuf:"varint,1011,opt,name=to_basic_price_sum_null,json=toBasicPriceSumNull" json:"to_basic_price_sum_null,omitempty"`
	DeleteConditionExtension      *dstore_values.BooleanValue `protobuf:"bytes,12,opt,name=delete_condition_extension,json=deleteConditionExtension" json:"delete_condition_extension,omitempty"`
	DeleteConditionExtensionNull  bool                        `protobuf:"varint,1012,opt,name=delete_condition_extension_null,json=deleteConditionExtensionNull" json:"delete_condition_extension_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetConditionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ConditionId
	}
	return nil
}

func (m *Parameters) GetItemConditionGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ItemConditionGroupId
	}
	return nil
}

func (m *Parameters) GetConditionGroupDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionGroupDescription
	}
	return nil
}

func (m *Parameters) GetMinNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.MinNumberOfItems
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfItems
	}
	return nil
}

func (m *Parameters) GetFromQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromQuantity
	}
	return nil
}

func (m *Parameters) GetToQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToQuantity
	}
	return nil
}

func (m *Parameters) GetFromItemBasicPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromItemBasicPrice
	}
	return nil
}

func (m *Parameters) GetToItemBasicPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToItemBasicPrice
	}
	return nil
}

func (m *Parameters) GetFromBasicPriceSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.FromBasicPriceSum
	}
	return nil
}

func (m *Parameters) GetToBasicPriceSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.ToBasicPriceSum
	}
	return nil
}

func (m *Parameters) GetDeleteConditionExtension() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteConditionExtension
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_ModifyCampItemConGroups_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_ModifyCampItemConGroups_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_ModifyCampItemConGroups_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_ModifyCampItemConGroups_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x96, 0xeb, 0x52, 0xdb, 0x46,
	0x14, 0xc7, 0x07, 0x5c, 0x03, 0x3d, 0x76, 0x8b, 0x59, 0x6e, 0xc2, 0x86, 0x96, 0xa1, 0x5f, 0xda,
	0x61, 0x2a, 0x3a, 0xed, 0x74, 0x5a, 0x26, 0x4c, 0x48, 0x20, 0x24, 0x38, 0x09, 0x0e, 0x51, 0x66,
	0x32, 0x13, 0x3e, 0x44, 0x23, 0x5b, 0x6b, 0xcf, 0x4e, 0x2c, 0xad, 0xb3, 0x2b, 0x07, 0x78, 0x8b,
	0xbc, 0x50, 0x9e, 0x23, 0xcf, 0x90, 0xfb, 0xf5, 0x01, 0xb2, 0xbb, 0x47, 0xbe, 0x49, 0x36, 0x0e,
	0x5f, 0x6c, 0xaf, 0x76, 0xff, 0xff, 0xff, 0x6f, 0x8f, 0xe4, 0x3d, 0x82, 0x1d, 0x5f, 0x46, 0x5c,
	0xd0, 0x2d, 0x1a, 0x36, 0x58, 0x48, 0xb7, 0x5a, 0x82, 0xd7, 0xa8, 0xdf, 0x16, 0x54, 0x6e, 0xf1,
	0xc0, 0x3d, 0xe2, 0x3e, 0xab, 0x9f, 0xef, 0x7b, 0x41, 0xab, 0x1c, 0xd1, 0x60, 0x9f, 0x87, 0xb7,
	0x04, 0x6f, 0xb7, 0xa4, 0x7b, 0xdd, 0xb7, 0xd5, 0xc2, 0x88, 0x93, 0x4d, 0x54, 0xdb, 0xa8, 0xb6,
	0x2f, 0x94, 0x14, 0xe7, 0xe3, 0xa8, 0x67, 0x5e, 0xb3, 0x4d, 0x25, 0x3a, 0x14, 0x57, 0x06, 0xf3,
	0xa9, 0x10, 0x5c, 0xc4, 0x53, 0xa5, 0xc1, 0xa9, 0x80, 0x4a, 0xe9, 0x35, 0x68, 0x3c, 0xf9, 0x5b,
	0x72, 0x32, 0xf2, 0x58, 0x58, 0xe7, 0x22, 0xf0, 0x22, 0xc6, 0x43, 0x5c, 0xb4, 0xf1, 0x32, 0x0f,
	0x70, 0xec, 0x09, 0x4f, 0xcd, 0x52, 0x21, 0xc9, 0x55, 0xc8, 0xd7, 0x78, 0xe8, 0x33, 0xbd, 0xc2,
	0x65, 0xbe, 0x35, 0xb1, 0x3e, 0xf1, 0x7b, 0xee, 0xef, 0x92, 0x1d, 0x6f, 0x22, 0xe6, 0x62, 0x61,
	0x44, 0x1b, 0x54, 0x3c, 0xd4, 0x23, 0x27, 0xd7, 0x15, 0x94, 0x7d, 0xb2, 0x09, 0x73, 0xfd, 0x7a,
	0x37, 0x6c, 0x37, 0x9b, 0xd6, 0xab, 0x69, 0xe5, 0x32, 0xe3, 0xcc, 0xf6, 0x2d, 0xac, 0xa8, 0xeb,
	0xc4, 0x81, 0x65, 0xa6, 0x2a, 0xe0, 0xf6, 0x14, 0x0d, 0x5d, 0x09, 0x9d, 0x3b, 0x39, 0x3e, 0x77,
	0x81, 0x61, 0xf5, 0x50, 0x6a, 0x6a, 0xa8, 0x00, 0x76, 0x61, 0x75, 0x84, 0x27, 0xb2, 0xbc, 0x46,
	0x16, 0x6b, 0x98, 0xd8, 0x40, 0x9d, 0x40, 0x29, 0xa9, 0xf5, 0xa9, 0xac, 0x09, 0xd6, 0xd2, 0x57,
	0xac, 0x8c, 0x01, 0x2b, 0x26, 0xc0, 0x64, 0x24, 0x58, 0xd8, 0x40, 0xae, 0x95, 0xda, 0x80, 0xed,
	0x8d, 0x9e, 0x98, 0x1c, 0xc2, 0xfa, 0x05, 0xde, 0x08, 0xf8, 0x06, 0x01, 0xd7, 0x46, 0xba, 0x18,
	0xca, 0xdb, 0x30, 0x1f, 0x30, 0xad, 0x08, 0xaa, 0x54, 0xb8, 0xbc, 0xee, 0xea, 0xfd, 0x48, 0xeb,
	0x87, 0xf1, 0x65, 0x2b, 0x28, 0x5d, 0xc5, 0xc8, 0xee, 0xd5, 0xf5, 0xf3, 0x27, 0xc9, 0x7f, 0x60,
	0x0d, 0xf1, 0x42, 0x9a, 0xb7, 0x48, 0xb3, 0x90, 0x14, 0x75, 0x21, 0xbc, 0xb3, 0x14, 0x44, 0xf6,
	0x7b, 0x20, 0xbc, 0xb3, 0x34, 0x44, 0xda, 0x0b, 0x21, 0xde, 0x75, 0x20, 0x12, 0x22, 0x03, 0x71,
	0x0d, 0x7e, 0xaa, 0x0b, 0xf5, 0xaf, 0x7a, 0xda, 0xf6, 0xc2, 0x88, 0x45, 0xe7, 0xd6, 0xd4, 0xf8,
	0xf8, 0xbc, 0x56, 0xdc, 0x8f, 0x05, 0xe4, 0x4f, 0x20, 0x03, 0x0e, 0x18, 0xfa, 0x1e, 0x43, 0x0b,
	0xfd, 0x4b, 0x4d, 0xe0, 0x0e, 0xe4, 0x22, 0xde, 0x8b, 0x9b, 0x1e, 0x1f, 0x07, 0x11, 0xef, 0x86,
	0xfd, 0x01, 0x85, 0x3e, 0x35, 0x46, 0x7d, 0xc0, 0xa8, 0x9f, 0x7b, 0xcb, 0x4c, 0x50, 0x05, 0x16,
	0x0d, 0x97, 0x79, 0x9e, 0xab, 0x9e, 0x64, 0x35, 0xb7, 0x25, 0x58, 0x8d, 0x5a, 0x33, 0x43, 0x23,
	0x7d, 0x5a, 0x63, 0x81, 0xd7, 0xc4, 0x48, 0xb3, 0x23, 0x5d, 0xa4, 0x3d, 0xad, 0x3b, 0xd6, 0x32,
	0x72, 0x05, 0x8a, 0x43, 0xfd, 0x10, 0xe2, 0x23, 0x42, 0x2c, 0xa5, 0x85, 0x9d, 0x7b, 0xad, 0xb8,
	0x53, 0x28, 0x3f, 0x8e, 0x47, 0x51, 0xfb, 0x4d, 0x80, 0xa8, 0x7b, 0x3d, 0xc4, 0x0b, 0x31, 0x3e,
	0xc5, 0xf7, 0x3a, 0x29, 0x32, 0x10, 0x77, 0x61, 0xc1, 0xec, 0xa0, 0x5f, 0x25, 0xdb, 0x81, 0x05,
	0xe3, 0x29, 0xe6, 0xb4, 0xb0, 0x67, 0xf7, 0xa0, 0x1d, 0x90, 0x6d, 0x58, 0x19, 0xe6, 0x86, 0x1c,
	0x9f, 0x91, 0x63, 0x31, 0x25, 0x33, 0x20, 0x87, 0x40, 0xd4, 0x0e, 0x92, 0x18, 0xb9, 0xf1, 0x18,
	0xb3, 0x11, 0x1f, 0x84, 0xf8, 0x17, 0x96, 0xd3, 0x4e, 0x88, 0xf0, 0x05, 0x11, 0xe6, 0x13, 0x12,
	0x03, 0xf0, 0x08, 0x8a, 0x3e, 0x6d, 0xaa, 0x33, 0xbb, 0xef, 0xa0, 0xa3, 0x67, 0x11, 0x0d, 0xa5,
	0x3e, 0xa4, 0xf2, 0x43, 0x41, 0xaa, 0x9c, 0x37, 0xa9, 0x17, 0x22, 0x88, 0x85, 0xf2, 0xee, 0x11,
	0x78, 0xd0, 0x11, 0x93, 0x9b, 0xf0, 0xeb, 0x68, 0x6b, 0x24, 0xfb, 0x8a, 0x64, 0xab, 0xa3, 0x3c,
	0x34, 0xe2, 0xc6, 0x8b, 0x49, 0x98, 0x71, 0xa8, 0x6c, 0xf1, 0x50, 0x52, 0xf2, 0x17, 0x64, 0x4d,
	0xdf, 0x8a, 0x1b, 0x4a, 0xf7, 0xfc, 0x8c, 0xbb, 0x22, 0xf6, 0xb4, 0x03, 0xfd, 0xe9, 0xe0, 0x42,
	0xb5, 0xc3, 0x82, 0xee, 0x58, 0x6e, 0x5f, 0xcb, 0x52, 0x5d, 0x21, 0xa3, 0xc4, 0x76, 0x42, 0x9c,
	0x6c, 0x6c, 0x47, 0x6a, 0x5c, 0xee, 0x8d, 0x9d, 0xd9, 0x60, 0xf0, 0x02, 0xf9, 0x1f, 0xa6, 0xe3,
	0x4e, 0xa9, 0x8e, 0x73, 0xed, 0xf8, 0x4b, 0xca, 0x11, 0xfb, 0xe8, 0x11, 0x7e, 0x3b, 0x9d, 0xe5,
	0xe4, 0x0e, 0x64, 0x04, 0x3f, 0x55, 0xc7, 0xac, 0x56, 0x6d, 0xdb, 0x97, 0x68, 0xed, 0x76, 0xa7,
	0x14, 0xb6, 0xc3, 0x4f, 0x1d, 0xed, 0x52, 0x5c, 0x83, 0x8c, 0xfa, 0x4d, 0x96, 0x60, 0x4a, 0x8d,
	0x74, 0xd3, 0x7b, 0x5e, 0x51, 0xc5, 0xc9, 0x3a, 0x59, 0x35, 0x2c, 0xfb, 0x7b, 0x8f, 0xa1, 0xc4,
	0x78, 0x22, 0xa2, 0xf7, 0xee, 0x71, 0xb2, 0xdb, 0xe0, 0xd2, 0x7f, 0xd2, 0x99, 0xf7, 0x2f, 0xfd,
	0x7a, 0x52, 0x9d, 0x32, 0x2f, 0x00, 0xff, 0x7c, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x2b, 0x80,
	0x0b, 0xdf, 0x08, 0x00, 0x00,
}
