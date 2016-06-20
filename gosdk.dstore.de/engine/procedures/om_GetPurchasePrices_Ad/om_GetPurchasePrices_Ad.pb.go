// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPurchasePrices_Ad.proto
// DO NOT EDIT!

/*
Package om_GetPurchasePrices_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPurchasePrices_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPurchasePrices_Ad

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
	TreeNodeOrNodeId                 *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=tree_node_or_node_id,json=treeNodeOrNodeId" json:"tree_node_or_node_id,omitempty"`
	TreeNodeOrNodeIdNull             bool                          `protobuf:"varint,1001,opt,name=tree_node_or_node_id_null,json=treeNodeOrNodeIdNull" json:"tree_node_or_node_id_null,omitempty"`
	IsTreeNodeId                     *dstore_values.BooleanValue   `protobuf:"bytes,2,opt,name=is_tree_node_id,json=isTreeNodeId" json:"is_tree_node_id,omitempty"`
	IsTreeNodeIdNull                 bool                          `protobuf:"varint,1002,opt,name=is_tree_node_id_null,json=isTreeNodeIdNull" json:"is_tree_node_id_null,omitempty"`
	Quantity                         *dstore_values.IntegerValue   `protobuf:"bytes,3,opt,name=quantity" json:"quantity,omitempty"`
	QuantityNull                     bool                          `protobuf:"varint,1003,opt,name=quantity_null,json=quantityNull" json:"quantity_null,omitempty"`
	SupplierId                       *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=supplier_id,json=supplierId" json:"supplier_id,omitempty"`
	SupplierIdNull                   bool                          `protobuf:"varint,1004,opt,name=supplier_id_null,json=supplierIdNull" json:"supplier_id_null,omitempty"`
	DateForPropertyDetermination     *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=date_for_property_determination,json=dateForPropertyDetermination" json:"date_for_property_determination,omitempty"`
	DateForPropertyDeterminationNull bool                          `protobuf:"varint,1005,opt,name=date_for_property_determination_null,json=dateForPropertyDeterminationNull" json:"date_for_property_determination_null,omitempty"`
	NodeCharacteristicId             *dstore_values.IntegerValue   `protobuf:"bytes,6,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull         bool                          `protobuf:"varint,1006,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeOrNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeOrNodeId
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

func (m *Parameters) GetSupplierId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SupplierId
	}
	return nil
}

func (m *Parameters) GetDateForPropertyDetermination() *dstore_values.TimestampValue {
	if m != nil {
		return m.DateForPropertyDetermination
	}
	return nil
}

func (m *Parameters) GetNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeCharacteristicId
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
	SupplCharacVal2RestrByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=suppl_charac_val2_restr_by_pattern,json=supplCharacVal2RestrByPattern" json:"suppl_charac_val2_restr_by_pattern,omitempty"`
	ItemProperty                  *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=item_property,json=itemProperty" json:"item_property,omitempty"`
	SupplierCharacteristicValue2  *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=supplier_characteristic_value2,json=supplierCharacteristicValue2" json:"supplier_characteristic_value2,omitempty"`
	TotalNetPrice                 *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=total_net_price,json=totalNetPrice" json:"total_net_price,omitempty"`
	Quantity                      *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=quantity" json:"quantity,omitempty"`
	SupplierCharacteristicValue1  *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=supplier_characteristic_value1,json=supplierCharacteristicValue1" json:"supplier_characteristic_value1,omitempty"`
	CurrencyId                    *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencySymbol                *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=currency_symbol,json=currencySymbol" json:"currency_symbol,omitempty"`
	SupplierId                    *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=supplier_id,json=supplierId" json:"supplier_id,omitempty"`
	NodeDescription               *dstore_values.StringValue  `protobuf:"bytes,10010,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	UnitNetPrice                  *dstore_values.DecimalValue `protobuf:"bytes,10011,opt,name=unit_net_price,json=unitNetPrice" json:"unit_net_price,omitempty"`
	HTreeNodeId                   *dstore_values.IntegerValue `protobuf:"bytes,10012,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	ItemNo                        *dstore_values.StringValue  `protobuf:"bytes,10013,opt,name=item_no,json=itemNo" json:"item_no,omitempty"`
	NodeId                        *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId                    *dstore_values.IntegerValue `protobuf:"bytes,10015,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	SupplCharacVal1RestrByPattern *dstore_values.StringValue  `protobuf:"bytes,10016,opt,name=suppl_charac_val1_restr_by_pattern,json=supplCharacVal1RestrByPattern" json:"suppl_charac_val1_restr_by_pattern,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetSupplCharacVal2RestrByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.SupplCharacVal2RestrByPattern
	}
	return nil
}

func (m *Response_Row) GetItemProperty() *dstore_values.StringValue {
	if m != nil {
		return m.ItemProperty
	}
	return nil
}

func (m *Response_Row) GetSupplierCharacteristicValue2() *dstore_values.StringValue {
	if m != nil {
		return m.SupplierCharacteristicValue2
	}
	return nil
}

func (m *Response_Row) GetTotalNetPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalNetPrice
	}
	return nil
}

func (m *Response_Row) GetQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Response_Row) GetSupplierCharacteristicValue1() *dstore_values.StringValue {
	if m != nil {
		return m.SupplierCharacteristicValue1
	}
	return nil
}

func (m *Response_Row) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Response_Row) GetCurrencySymbol() *dstore_values.StringValue {
	if m != nil {
		return m.CurrencySymbol
	}
	return nil
}

func (m *Response_Row) GetSupplierId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SupplierId
	}
	return nil
}

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetUnitNetPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.UnitNetPrice
	}
	return nil
}

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetItemNo() *dstore_values.StringValue {
	if m != nil {
		return m.ItemNo
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

func (m *Response_Row) GetSupplCharacVal1RestrByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.SupplCharacVal1RestrByPattern
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPurchasePrices_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPurchasePrices_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPurchasePrices_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetPurchasePrices_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xeb, 0x6e, 0xdc, 0x44,
	0x14, 0x56, 0x09, 0xd9, 0x84, 0xb3, 0x9b, 0x64, 0x65, 0xa2, 0xca, 0x4d, 0xda, 0x52, 0x85, 0x22,
	0xc1, 0x1f, 0x87, 0x6c, 0xb9, 0x4a, 0x50, 0xd1, 0x90, 0x16, 0x55, 0xa8, 0xee, 0x62, 0x10, 0x02,
	0xfe, 0x58, 0x13, 0x7b, 0xba, 0x19, 0xb1, 0x9e, 0x31, 0x33, 0x63, 0xaa, 0x7d, 0x0b, 0xee, 0x97,
	0x72, 0x7d, 0x09, 0xde, 0x83, 0xd7, 0xe0, 0xfa, 0x0c, 0x9c, 0x99, 0xb1, 0x77, 0xd7, 0xee, 0x82,
	0x77, 0xc5, 0x9f, 0x6c, 0xc6, 0xe7, 0x7c, 0xdf, 0xf9, 0xe6, 0xdc, 0x6c, 0x78, 0x21, 0x55, 0x5a,
	0x48, 0x7a, 0x48, 0xf9, 0x88, 0x71, 0x7a, 0x98, 0x4b, 0x91, 0xd0, 0xb4, 0x90, 0x54, 0x1d, 0x8a,
	0x2c, 0x7e, 0x83, 0xea, 0x61, 0x21, 0x93, 0x33, 0xa2, 0xe8, 0x50, 0xb2, 0x84, 0xaa, 0xf8, 0x46,
	0x1a, 0xa0, 0x8b, 0x16, 0xde, 0x53, 0x0e, 0x17, 0x38, 0x5c, 0xf0, 0x2f, 0xce, 0x7b, 0x8f, 0x97,
	0xf4, 0x1f, 0x93, 0x71, 0x41, 0x95, 0xc3, 0xee, 0x5d, 0xa8, 0xc7, 0xa4, 0x52, 0x0a, 0x59, 0x9a,
	0xf6, 0xeb, 0xa6, 0x8c, 0x2a, 0x45, 0x46, 0xb4, 0x34, 0x3e, 0xd9, 0x34, 0x6a, 0xc2, 0xf8, 0x3d,
	0x21, 0x33, 0xa2, 0x99, 0xe0, 0xce, 0xe9, 0xe0, 0x41, 0x07, 0x60, 0x48, 0x24, 0x41, 0x2b, 0x95,
	0xca, 0x7b, 0x13, 0x76, 0xb5, 0xa4, 0x34, 0xe6, 0x22, 0xa5, 0xb1, 0x90, 0xee, 0x97, 0xa5, 0xfe,
	0xb9, 0x2b, 0xe7, 0x9e, 0xee, 0x0e, 0xf6, 0x83, 0xf2, 0x1a, 0xa5, 0x3e, 0xc6, 0x35, 0x1d, 0x51,
	0xf9, 0xae, 0x39, 0x45, 0x7d, 0x03, 0x0c, 0xd1, 0xff, 0xae, 0x34, 0x7f, 0x6f, 0xa7, 0xde, 0x4b,
	0x70, 0x61, 0x11, 0x59, 0xcc, 0x8b, 0xf1, 0xd8, 0xff, 0x6d, 0x03, 0x29, 0x37, 0xa3, 0xdd, 0x26,
	0x2a, 0x44, 0xa3, 0x77, 0x0c, 0x3b, 0x4c, 0xc5, 0x33, 0x30, 0x2a, 0x78, 0x64, 0xa1, 0x82, 0x53,
	0x21, 0xc6, 0x94, 0x70, 0xa7, 0xa0, 0xc7, 0xd4, 0x3b, 0x25, 0x1b, 0x46, 0x3f, 0x84, 0xdd, 0x06,
	0x87, 0x0b, 0xfc, 0xbb, 0x0b, 0xdc, 0x9f, 0x77, 0xb6, 0x41, 0x5f, 0x84, 0xcd, 0x8f, 0x0a, 0xc2,
	0x35, 0xd3, 0x13, 0x7f, 0xad, 0xfd, 0xbe, 0x53, 0x67, 0xef, 0x2a, 0x6c, 0x55, 0xff, 0xbb, 0x10,
	0x7f, 0xb8, 0x10, 0xbd, 0xea, 0xa9, 0xa5, 0x7f, 0x05, 0xba, 0xaa, 0xc8, 0xf3, 0x31, 0xa3, 0xd2,
	0xdc, 0xe7, 0xd1, 0xf6, 0x08, 0x50, 0xf9, 0xe3, 0x6d, 0x9e, 0x81, 0xfe, 0x1c, 0xda, 0x85, 0xf9,
	0xd3, 0x85, 0xd9, 0x9e, 0xb9, 0xd9, 0x40, 0x29, 0x3c, 0x91, 0x12, 0x4d, 0x63, 0x2c, 0x75, 0x8c,
	0x45, 0xce, 0xa9, 0x44, 0x5d, 0xa9, 0xa9, 0x6f, 0xc6, 0xb8, 0xad, 0xbd, 0xbf, 0x6e, 0x83, 0x5f,
	0x6a, 0x04, 0xd7, 0x0c, 0x1b, 0x48, 0x93, 0x2c, 0x77, 0xe1, 0x2f, 0x1a, 0x96, 0x5b, 0x42, 0x0e,
	0x4b, 0x8e, 0x93, 0x79, 0x0a, 0xef, 0x2e, 0x5c, 0x6d, 0x89, 0xe2, 0x44, 0xfe, 0xe5, 0x44, 0x5e,
	0xf9, 0x2f, 0x32, 0x2b, 0xfb, 0x2d, 0x38, 0x6f, 0xeb, 0x84, 0x23, 0x21, 0x49, 0x82, 0x56, 0xa6,
	0x34, 0x4b, 0x4c, 0xaa, 0x3a, 0xed, 0xa9, 0xda, 0x35, 0xd0, 0xd7, 0x6b, 0x48, 0x4c, 0xda, 0x75,
	0xd8, 0x5f, 0x4c, 0xe9, 0xa4, 0xfd, 0xed, 0xa4, 0xf9, 0x8b, 0xb0, 0x46, 0xd2, 0xc1, 0xaf, 0x5d,
	0xd8, 0x8c, 0xa8, 0xca, 0x05, 0x57, 0xd4, 0x7b, 0x16, 0xd6, 0xed, 0xe8, 0x95, 0xb3, 0xb0, 0x17,
	0xd4, 0x47, 0xda, 0x8d, 0xe5, 0x4d, 0xf3, 0x37, 0x72, 0x8e, 0xde, 0xfb, 0xd0, 0x37, 0x43, 0x17,
	0xcf, 0x4d, 0x1d, 0xb6, 0xf1, 0x1a, 0x82, 0x83, 0x06, 0xb8, 0x39, 0x9b, 0x77, 0xf0, 0x7c, 0x7b,
	0x76, 0x8e, 0x76, 0xb2, 0xfa, 0x03, 0x1c, 0xad, 0x8d, 0x72, 0xd8, 0xb1, 0x55, 0x0d, 0xe3, 0xe5,
	0x87, 0x18, 0xdd, 0x2a, 0xb8, 0xe3, 0x7e, 0xa3, 0xca, 0xdd, 0xbb, 0x09, 0x6b, 0x52, 0xdc, 0xc7,
	0xf6, 0x33, 0xa8, 0x6b, 0xc1, 0x52, 0x7b, 0x29, 0xa8, 0x92, 0x10, 0x44, 0xe2, 0x7e, 0x64, 0xf0,
	0x7b, 0xbf, 0x3c, 0x06, 0x6b, 0x78, 0xf0, 0xce, 0x43, 0x07, 0x8f, 0xa6, 0x4a, 0x9f, 0x84, 0x98,
	0x97, 0xf5, 0x68, 0x1d, 0x8f, 0x98, 0xfa, 0x7b, 0x70, 0x60, 0xdb, 0xb2, 0xcc, 0x7d, 0x8c, 0x45,
	0x1b, 0xc4, 0xb8, 0x27, 0xb5, 0x8c, 0x4f, 0x27, 0x71, 0x4e, 0x34, 0xa6, 0x9a, 0xfb, 0x9f, 0x86,
	0xf5, 0x5c, 0x96, 0xa5, 0x45, 0x2f, 0xc6, 0x47, 0xae, 0xb2, 0x97, 0x2c, 0x8d, 0x2b, 0x0f, 0x3e,
	0x19, 0x44, 0x86, 0xe3, 0x78, 0x32, 0x74, 0x0c, 0xde, 0x6b, 0xb0, 0xc5, 0x34, 0xcd, 0xa6, 0x2d,
	0xe8, 0x7f, 0xd6, 0x4e, 0xd9, 0x33, 0x88, 0xaa, 0x11, 0x3d, 0x02, 0x97, 0xa7, 0x93, 0xd5, 0x68,
	0x14, 0x8b, 0x1d, 0xf8, 0x9f, 0xb7, 0x53, 0x5e, 0xac, 0x28, 0xea, 0x7d, 0x64, 0x8d, 0x03, 0xef,
	0x04, 0x76, 0xb4, 0xd0, 0x64, 0x1c, 0x73, 0xaa, 0x51, 0x29, 0x66, 0xd5, 0xff, 0x22, 0x5c, 0xd8,
	0xd4, 0x29, 0x4d, 0x58, 0x46, 0xc6, 0x8e, 0x74, 0xcb, 0x82, 0x42, 0x2c, 0x88, 0x81, 0x60, 0xcd,
	0x67, 0xfb, 0xe9, 0xcb, 0x70, 0x95, 0x05, 0xd5, 0x76, 0xc5, 0x23, 0xff, 0xab, 0xff, 0x77, 0xc5,
	0x23, 0xef, 0x55, 0xe8, 0x26, 0x85, 0x94, 0x94, 0x27, 0x13, 0xd3, 0x0c, 0x5f, 0x2f, 0xa1, 0x0f,
	0x2a, 0x00, 0xb6, 0x0b, 0x66, 0x68, 0x0a, 0x57, 0x93, 0xec, 0x54, 0x8c, 0xfd, 0x6f, 0xda, 0x25,
	0x6d, 0x57, 0x98, 0xb7, 0x2d, 0xc4, 0x88, 0x98, 0x5f, 0xb1, 0xdf, 0x86, 0xab, 0xed, 0xd8, 0x5b,
	0xd0, 0xb7, 0xeb, 0x22, 0xa5, 0x2a, 0x91, 0x2c, 0xb7, 0xf3, 0xfa, 0xa0, 0x5d, 0xc5, 0x8e, 0x01,
	0x9d, 0xcc, 0x30, 0xf8, 0xf6, 0xda, 0x2e, 0x38, 0xd3, 0x73, 0xd5, 0xfe, 0x6e, 0x89, 0x6a, 0xf7,
	0x0c, 0x66, 0x5a, 0xec, 0x1b, 0xb0, 0x7d, 0x56, 0x7f, 0x01, 0x7e, 0xbf, 0xc4, 0x6d, 0xba, 0x67,
	0x73, 0x2f, 0xc0, 0xe7, 0x60, 0xc3, 0x8e, 0x06, 0x17, 0xfe, 0x0f, 0xed, 0xb7, 0xe8, 0x18, 0xdf,
	0x50, 0x78, 0xcf, 0xc3, 0x46, 0x15, 0xf1, 0xc7, 0x25, 0x22, 0x76, 0xb8, 0x0b, 0x76, 0x1d, 0x7a,
	0x35, 0xb5, 0x3f, 0x2d, 0x93, 0x7b, 0x3d, 0x13, 0xbb, 0x60, 0x5f, 0x1c, 0x3d, 0xbc, 0x2f, 0x7e,
	0x5e, 0x79, 0x5f, 0x1c, 0xd5, 0xf7, 0xc5, 0xf1, 0x7b, 0xb0, 0xcf, 0x44, 0x63, 0xeb, 0xcd, 0xbe,
	0xe2, 0x3e, 0x78, 0x79, 0x24, 0x54, 0xfa, 0x61, 0x65, 0x4f, 0x57, 0xf8, 0xd0, 0x3b, 0xed, 0xd8,
	0x0f, 0xaa, 0x6b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x81, 0x3e, 0x70, 0x23, 0x0a, 0x00,
	0x00,
}