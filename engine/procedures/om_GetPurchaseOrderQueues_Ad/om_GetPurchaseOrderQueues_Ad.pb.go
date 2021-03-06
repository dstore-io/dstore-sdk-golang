// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPurchaseOrderQueues_Ad.proto
// DO NOT EDIT!

/*
Package om_GetPurchaseOrderQueues_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPurchaseOrderQueues_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPurchaseOrderQueues_Ad

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
	SupplierId               *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=supplier_id,json=supplierId" json:"supplier_id,omitempty"`
	SupplierIdNull           bool                          `protobuf:"varint,1001,opt,name=supplier_id_null,json=supplierIdNull" json:"supplier_id_null,omitempty"`
	FromOrderDeadline        *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=from_order_deadline,json=fromOrderDeadline" json:"from_order_deadline,omitempty"`
	FromOrderDeadlineNull    bool                          `protobuf:"varint,1002,opt,name=from_order_deadline_null,json=fromOrderDeadlineNull" json:"from_order_deadline_null,omitempty"`
	ToOrderDeadline          *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=to_order_deadline,json=toOrderDeadline" json:"to_order_deadline,omitempty"`
	ToOrderDeadlineNull      bool                          `protobuf:"varint,1003,opt,name=to_order_deadline_null,json=toOrderDeadlineNull" json:"to_order_deadline_null,omitempty"`
	OrderTypeId              *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=order_type_id,json=orderTypeId" json:"order_type_id,omitempty"`
	OrderTypeIdNull          bool                          `protobuf:"varint,1004,opt,name=order_type_id_null,json=orderTypeIdNull" json:"order_type_id_null,omitempty"`
	NodeId                   *dstore_values.IntegerValue   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	NodeIdNull               bool                          `protobuf:"varint,1005,opt,name=node_id_null,json=nodeIdNull" json:"node_id_null,omitempty"`
	GetSummaryOnly           *dstore_values.IntegerValue   `protobuf:"bytes,6,opt,name=get_summary_only,json=getSummaryOnly" json:"get_summary_only,omitempty"`
	GetSummaryOnlyNull       bool                          `protobuf:"varint,1006,opt,name=get_summary_only_null,json=getSummaryOnlyNull" json:"get_summary_only_null,omitempty"`
	NodeCharacteristicId     *dstore_values.IntegerValue   `protobuf:"bytes,7,opt,name=node_characteristic_id,json=nodeCharacteristicId" json:"node_characteristic_id,omitempty"`
	NodeCharacteristicIdNull bool                          `protobuf:"varint,1007,opt,name=node_characteristic_id_null,json=nodeCharacteristicIdNull" json:"node_characteristic_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSupplierId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SupplierId
	}
	return nil
}

func (m *Parameters) GetFromOrderDeadline() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromOrderDeadline
	}
	return nil
}

func (m *Parameters) GetToOrderDeadline() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToOrderDeadline
	}
	return nil
}

func (m *Parameters) GetOrderTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderTypeId
	}
	return nil
}

func (m *Parameters) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Parameters) GetGetSummaryOnly() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetSummaryOnly
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
	RowId                         int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	SupplCharacVal2RestrByPattern *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=suppl_charac_val2_restr_by_pattern,json=supplCharacVal2RestrByPattern" json:"suppl_charac_val2_restr_by_pattern,omitempty"`
	LastEditedByUserId            *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=last_edited_by_user_id,json=lastEditedByUserId" json:"last_edited_by_user_id,omitempty"`
	CurrencySymbol                *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=currency_symbol,json=currencySymbol" json:"currency_symbol,omitempty"`
	LastEditedAtDateAndTime       *dstore_values.TimestampValue `protobuf:"bytes,10004,opt,name=last_edited_at_date_and_time,json=lastEditedAtDateAndTime" json:"last_edited_at_date_and_time,omitempty"`
	SupplierId                    *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=supplier_id,json=supplierId" json:"supplier_id,omitempty"`
	NetSum                        *dstore_values.DecimalValue   `protobuf:"bytes,10006,opt,name=net_sum,json=netSum" json:"net_sum,omitempty"`
	NodeDescription               *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	OrderDeadline                 *dstore_values.TimestampValue `protobuf:"bytes,10008,opt,name=order_deadline,json=orderDeadline" json:"order_deadline,omitempty"`
	NodeId                        *dstore_values.IntegerValue   `protobuf:"bytes,10009,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId                    *dstore_values.IntegerValue   `protobuf:"bytes,10010,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	CreatedByUserName             *dstore_values.StringValue    `protobuf:"bytes,10011,opt,name=created_by_user_name,json=createdByUserName" json:"created_by_user_name,omitempty"`
	Comment                       *dstore_values.StringValue    `protobuf:"bytes,10012,opt,name=comment" json:"comment,omitempty"`
	SupplierCharacteristicValue2  *dstore_values.StringValue    `protobuf:"bytes,10013,opt,name=supplier_characteristic_value2,json=supplierCharacteristicValue2" json:"supplier_characteristic_value2,omitempty"`
	ItemProperty                  *dstore_values.StringValue    `protobuf:"bytes,10014,opt,name=item_property,json=itemProperty" json:"item_property,omitempty"`
	SupplierCharacteristicValue1  *dstore_values.StringValue    `protobuf:"bytes,10015,opt,name=supplier_characteristic_value1,json=supplierCharacteristicValue1" json:"supplier_characteristic_value1,omitempty"`
	Quantity                      *dstore_values.IntegerValue   `protobuf:"bytes,10016,opt,name=quantity" json:"quantity,omitempty"`
	LastEditedByUserName          *dstore_values.StringValue    `protobuf:"bytes,10017,opt,name=last_edited_by_user_name,json=lastEditedByUserName" json:"last_edited_by_user_name,omitempty"`
	CurrencyId                    *dstore_values.IntegerValue   `protobuf:"bytes,10018,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CreatedAtDateAndTime          *dstore_values.TimestampValue `protobuf:"bytes,10019,opt,name=created_at_date_and_time,json=createdAtDateAndTime" json:"created_at_date_and_time,omitempty"`
	CreatedByUserId               *dstore_values.IntegerValue   `protobuf:"bytes,10020,opt,name=created_by_user_id,json=createdByUserId" json:"created_by_user_id,omitempty"`
	OrderTypeDescription          *dstore_values.StringValue    `protobuf:"bytes,10021,opt,name=order_type_description,json=orderTypeDescription" json:"order_type_description,omitempty"`
	HTreeNodeId                   *dstore_values.IntegerValue   `protobuf:"bytes,10022,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	OrderTypeId                   *dstore_values.IntegerValue   `protobuf:"bytes,10023,opt,name=order_type_id,json=orderTypeId" json:"order_type_id,omitempty"`
	ItemNo                        *dstore_values.StringValue    `protobuf:"bytes,10024,opt,name=item_no,json=itemNo" json:"item_no,omitempty"`
	SupplCharacVal1RestrByPattern *dstore_values.StringValue    `protobuf:"bytes,10025,opt,name=suppl_charac_val1_restr_by_pattern,json=supplCharacVal1RestrByPattern" json:"suppl_charac_val1_restr_by_pattern,omitempty"`
	NumberOfItems                 *dstore_values.IntegerValue   `protobuf:"bytes,20001,opt,name=number_of_items,json=numberOfItems" json:"number_of_items,omitempty"`
	NumberOfExpiredDeadlines      *dstore_values.IntegerValue   `protobuf:"bytes,20002,opt,name=number_of_expired_deadlines,json=numberOfExpiredDeadlines" json:"number_of_expired_deadlines,omitempty"`
	MinOrderDeadline              *dstore_values.DecimalValue   `protobuf:"bytes,20003,opt,name=min_order_deadline,json=minOrderDeadline" json:"min_order_deadline,omitempty"`
	NumberOfQueues                *dstore_values.IntegerValue   `protobuf:"bytes,20004,opt,name=number_of_queues,json=numberOfQueues" json:"number_of_queues,omitempty"`
	MaxOrderDeadline              *dstore_values.DecimalValue   `protobuf:"bytes,20005,opt,name=max_order_deadline,json=maxOrderDeadline" json:"max_order_deadline,omitempty"`
	NumberOfSuppliers             *dstore_values.IntegerValue   `protobuf:"bytes,20006,opt,name=number_of_suppliers,json=numberOfSuppliers" json:"number_of_suppliers,omitempty"`
	TotalNetSum                   *dstore_values.DecimalValue   `protobuf:"bytes,30002,opt,name=total_net_sum,json=totalNetSum" json:"total_net_sum,omitempty"`
	PurchasePriceCharacteristicId *dstore_values.IntegerValue   `protobuf:"bytes,30006,opt,name=purchase_price_characteristic_id,json=purchasePriceCharacteristicId" json:"purchase_price_characteristic_id,omitempty"`
	TotalQuantity                 *dstore_values.IntegerValue   `protobuf:"bytes,30012,opt,name=total_quantity,json=totalQuantity" json:"total_quantity,omitempty"`
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

func (m *Response_Row) GetLastEditedByUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LastEditedByUserId
	}
	return nil
}

func (m *Response_Row) GetCurrencySymbol() *dstore_values.StringValue {
	if m != nil {
		return m.CurrencySymbol
	}
	return nil
}

func (m *Response_Row) GetLastEditedAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.LastEditedAtDateAndTime
	}
	return nil
}

func (m *Response_Row) GetSupplierId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SupplierId
	}
	return nil
}

func (m *Response_Row) GetNetSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.NetSum
	}
	return nil
}

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetOrderDeadline() *dstore_values.TimestampValue {
	if m != nil {
		return m.OrderDeadline
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

func (m *Response_Row) GetCreatedByUserName() *dstore_values.StringValue {
	if m != nil {
		return m.CreatedByUserName
	}
	return nil
}

func (m *Response_Row) GetComment() *dstore_values.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *Response_Row) GetSupplierCharacteristicValue2() *dstore_values.StringValue {
	if m != nil {
		return m.SupplierCharacteristicValue2
	}
	return nil
}

func (m *Response_Row) GetItemProperty() *dstore_values.StringValue {
	if m != nil {
		return m.ItemProperty
	}
	return nil
}

func (m *Response_Row) GetSupplierCharacteristicValue1() *dstore_values.StringValue {
	if m != nil {
		return m.SupplierCharacteristicValue1
	}
	return nil
}

func (m *Response_Row) GetQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Response_Row) GetLastEditedByUserName() *dstore_values.StringValue {
	if m != nil {
		return m.LastEditedByUserName
	}
	return nil
}

func (m *Response_Row) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Response_Row) GetCreatedAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.CreatedAtDateAndTime
	}
	return nil
}

func (m *Response_Row) GetCreatedByUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CreatedByUserId
	}
	return nil
}

func (m *Response_Row) GetOrderTypeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.OrderTypeDescription
	}
	return nil
}

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetOrderTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OrderTypeId
	}
	return nil
}

func (m *Response_Row) GetItemNo() *dstore_values.StringValue {
	if m != nil {
		return m.ItemNo
	}
	return nil
}

func (m *Response_Row) GetSupplCharacVal1RestrByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.SupplCharacVal1RestrByPattern
	}
	return nil
}

func (m *Response_Row) GetNumberOfItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfItems
	}
	return nil
}

func (m *Response_Row) GetNumberOfExpiredDeadlines() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfExpiredDeadlines
	}
	return nil
}

func (m *Response_Row) GetMinOrderDeadline() *dstore_values.DecimalValue {
	if m != nil {
		return m.MinOrderDeadline
	}
	return nil
}

func (m *Response_Row) GetNumberOfQueues() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfQueues
	}
	return nil
}

func (m *Response_Row) GetMaxOrderDeadline() *dstore_values.DecimalValue {
	if m != nil {
		return m.MaxOrderDeadline
	}
	return nil
}

func (m *Response_Row) GetNumberOfSuppliers() *dstore_values.IntegerValue {
	if m != nil {
		return m.NumberOfSuppliers
	}
	return nil
}

func (m *Response_Row) GetTotalNetSum() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalNetSum
	}
	return nil
}

func (m *Response_Row) GetPurchasePriceCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PurchasePriceCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetTotalQuantity() *dstore_values.IntegerValue {
	if m != nil {
		return m.TotalQuantity
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPurchaseOrderQueues_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPurchaseOrderQueues_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPurchaseOrderQueues_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetPurchaseOrderQueues_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x57, 0xeb, 0x4e, 0x24, 0x45,
	0x14, 0xce, 0x8a, 0x0c, 0xa4, 0x80, 0x19, 0x28, 0x58, 0x2c, 0x61, 0xd9, 0xac, 0xf8, 0x47, 0x8d,
	0x19, 0x64, 0xdc, 0x4d, 0x36, 0x51, 0x71, 0x41, 0x70, 0x83, 0x4a, 0x33, 0x0c, 0xb8, 0x89, 0xc6,
	0x4d, 0xa7, 0x98, 0xae, 0x19, 0x3a, 0x4e, 0x77, 0xcd, 0x56, 0x57, 0xbb, 0xcc, 0x5b, 0x78, 0xbf,
	0xae, 0x97, 0xbd, 0xaa, 0x7f, 0x8d, 0xf1, 0x97, 0x8f, 0xa1, 0x0f, 0x21, 0xde, 0x5e, 0xc1, 0x53,
	0x55, 0xdd, 0x33, 0xd3, 0x3d, 0xa3, 0xdd, 0xc4, 0x3f, 0x90, 0x9e, 0xae, 0xef, 0x3b, 0x5f, 0x9d,
	0x3a, 0xe7, 0xab, 0xd3, 0xe8, 0x39, 0x27, 0x90, 0x5c, 0xb0, 0x15, 0xe6, 0x37, 0x5d, 0x9f, 0xad,
	0xb4, 0x05, 0xaf, 0x33, 0x27, 0x14, 0x2c, 0x58, 0xe1, 0x9e, 0x7d, 0x95, 0xc9, 0x6a, 0x28, 0xea,
	0x47, 0x34, 0x60, 0xbb, 0xc2, 0x61, 0x62, 0x2f, 0x64, 0x21, 0x0b, 0xec, 0x75, 0xa7, 0x0c, 0xeb,
	0x24, 0xc7, 0x4f, 0x19, 0x70, 0xd9, 0x80, 0xcb, 0xff, 0x85, 0x58, 0x98, 0x8d, 0x02, 0xbd, 0x43,
	0x5b, 0xf0, 0x8b, 0x21, 0x58, 0x78, 0x34, 0x19, 0x9d, 0x09, 0xc1, 0x45, 0xf4, 0x6a, 0x31, 0xf9,
	0xca, 0x63, 0x41, 0x40, 0x9b, 0x2c, 0x7a, 0xf9, 0x78, 0xfa, 0xa5, 0xa4, 0xae, 0xdf, 0xe0, 0xc2,
	0xa3, 0xd2, 0xe5, 0xbe, 0x59, 0xb4, 0xfc, 0x6b, 0x01, 0xa1, 0x2a, 0x15, 0x14, 0xde, 0x32, 0x11,
	0xe0, 0xe7, 0xd1, 0x44, 0x10, 0xb6, 0xdb, 0x2d, 0x97, 0x09, 0xdb, 0x75, 0xc8, 0x99, 0x0b, 0x67,
	0x9e, 0x98, 0xa8, 0x2c, 0x96, 0xa3, 0x2d, 0x44, 0xb2, 0x5c, 0x5f, 0xb2, 0x26, 0x13, 0xd7, 0xd4,
	0x53, 0x0d, 0xc5, 0xeb, 0xb7, 0x1d, 0xfc, 0x24, 0x9a, 0xee, 0x43, 0xdb, 0x7e, 0xd8, 0x6a, 0x91,
	0xdf, 0xc6, 0x80, 0x63, 0xbc, 0x56, 0xec, 0x2d, 0xb3, 0xe0, 0x67, 0xbc, 0x83, 0x66, 0x1b, 0x02,
	0x72, 0xc1, 0x55, 0x02, 0x6c, 0x87, 0x51, 0xa7, 0x05, 0x22, 0xc9, 0x43, 0x3a, 0xe0, 0x52, 0x2a,
	0xa0, 0x74, 0x61, 0x67, 0x92, 0x7a, 0x6d, 0x13, 0x72, 0x46, 0x21, 0x75, 0xe6, 0x36, 0x23, 0x1c,
	0xbe, 0x8c, 0xc8, 0x10, 0x3a, 0xa3, 0xe0, 0xc4, 0x28, 0x38, 0x3b, 0x80, 0xd2, 0x42, 0xb6, 0xd1,
	0x8c, 0xe4, 0x69, 0x19, 0x23, 0x79, 0x64, 0x94, 0x24, 0x4f, 0x8a, 0xb8, 0x88, 0xe6, 0x07, 0xa8,
	0x8c, 0x84, 0xdf, 0x8d, 0x84, 0xd9, 0x14, 0x42, 0x0b, 0x78, 0x11, 0x4d, 0x19, 0x88, 0xec, 0xb4,
	0x99, 0x4a, 0xfa, 0xc3, 0xd9, 0x49, 0x9f, 0xd0, 0x88, 0x03, 0x00, 0x40, 0xd6, 0x9f, 0x46, 0x38,
	0x41, 0x60, 0x42, 0xfe, 0x61, 0x42, 0x96, 0xfa, 0x56, 0xea, 0x70, 0x17, 0xd1, 0x98, 0xcf, 0x1d,
	0x1d, 0x68, 0x34, 0x3b, 0x50, 0x41, 0xad, 0x85, 0x18, 0x8f, 0xa1, 0xc9, 0x08, 0x65, 0xd8, 0xff,
	0x34, 0xec, 0xc8, 0xbc, 0xd6, 0xc4, 0x5b, 0x68, 0xba, 0xc9, 0xa4, 0x1d, 0x84, 0x9e, 0x47, 0x45,
	0xc7, 0xe6, 0x7e, 0xab, 0x43, 0x0a, 0xd9, 0x11, 0x8a, 0x00, 0xda, 0x37, 0x98, 0x5d, 0x80, 0xe0,
	0x0a, 0x3a, 0x9b, 0xa6, 0x31, 0x21, 0xff, 0x32, 0x21, 0x71, 0x72, 0xbd, 0x0e, 0xbd, 0x87, 0xe6,
	0xb5, 0x3a, 0x68, 0x29, 0x41, 0xeb, 0x50, 0xc8, 0x6e, 0x20, 0xdd, 0xba, 0xda, 0xe2, 0x58, 0xb6,
	0x80, 0x39, 0x05, 0x7d, 0x29, 0x81, 0x84, 0x0d, 0xaf, 0xa1, 0xc5, 0xe1, 0x94, 0x46, 0xcc, 0xdf,
	0x46, 0x0c, 0x19, 0x86, 0x55, 0x92, 0x96, 0x7f, 0x99, 0x47, 0xe3, 0x35, 0x16, 0xb4, 0xb9, 0x1f,
	0x30, 0xfc, 0x0c, 0x1a, 0xd5, 0x5d, 0x1b, 0xf5, 0xd3, 0x42, 0x39, 0x69, 0x09, 0xa6, 0xa3, 0xb7,
	0xd4, 0xdf, 0x9a, 0x59, 0x88, 0xdf, 0x40, 0xd3, 0xaa, 0x5f, 0xed, 0xbe, 0x86, 0x85, 0xde, 0x18,
	0x01, 0x70, 0x39, 0x05, 0x4e, 0xb7, 0xf5, 0x0e, 0x3c, 0x6f, 0xf7, 0x9e, 0x6b, 0x25, 0x2f, 0xf9,
	0x03, 0xb4, 0xca, 0x58, 0xe4, 0x13, 0x50, 0xe6, 0x8a, 0xf1, 0xfc, 0x00, 0xa3, 0x71, 0x91, 0x1d,
	0xf3, 0xbf, 0x16, 0x2f, 0xc7, 0xaf, 0xa0, 0x11, 0xc1, 0x6f, 0x42, 0x7d, 0x2a, 0xd4, 0xe5, 0x72,
	0x7e, 0x5f, 0x2b, 0xc7, 0x99, 0x28, 0xd7, 0xf8, 0xcd, 0x9a, 0x22, 0x59, 0xf8, 0x71, 0x0e, 0x8d,
	0xc0, 0x03, 0x9e, 0x47, 0x05, 0x78, 0x54, 0x47, 0xf5, 0xae, 0x05, 0xc9, 0x19, 0xad, 0x8d, 0xc2,
	0x23, 0xe4, 0xbf, 0x81, 0x96, 0xb5, 0x63, 0x44, 0x07, 0x60, 0xc3, 0xc9, 0x55, 0x6c, 0x70, 0x5c,
	0x29, 0xec, 0xc3, 0x8e, 0xdd, 0xa6, 0x12, 0xf2, 0xed, 0x93, 0xf7, 0xac, 0x64, 0x42, 0xa3, 0xf3,
	0x85, 0x55, 0xae, 0xdf, 0x34, 0xc7, 0xbb, 0xa4, 0x69, 0xcc, 0x19, 0xc1, 0x2f, 0x95, 0x9a, 0xe2,
	0xd8, 0xe8, 0x54, 0x0d, 0x03, 0xae, 0xa2, 0xf9, 0x16, 0x0d, 0xa4, 0xcd, 0x1c, 0x57, 0x32, 0x47,
	0x71, 0x87, 0x81, 0xf1, 0xbe, 0xf7, 0xad, 0xec, 0xda, 0xc1, 0x0a, 0xbb, 0xa5, 0xa1, 0x1b, 0x9d,
	0xd7, 0x03, 0x6d, 0x82, 0x9b, 0xa8, 0x54, 0x0f, 0x85, 0x60, 0x7e, 0xbd, 0x63, 0x07, 0x1d, 0xef,
	0x90, 0xb7, 0xc8, 0x07, 0xd9, 0x32, 0x8b, 0x31, 0x66, 0x5f, 0x43, 0xf0, 0x75, 0x74, 0xae, 0x5f,
	0x17, 0x95, 0xb6, 0x43, 0x25, 0xb3, 0xa9, 0xef, 0xd8, 0xca, 0x86, 0xc8, 0x87, 0x56, 0x1e, 0x8b,
	0x7a, 0xa4, 0xa7, 0x6f, 0x5d, 0x6e, 0x02, 0xc1, 0xba, 0xef, 0x1c, 0xc0, 0x12, 0xfc, 0x42, 0xd2,
	0xe7, 0x3f, 0xb2, 0x4e, 0x67, 0xf4, 0x97, 0xc0, 0x44, 0x4c, 0x93, 0x92, 0x8f, 0x87, 0x43, 0x1d,
	0x56, 0x77, 0x3d, 0xda, 0x8a, 0x5d, 0x44, 0xf7, 0x2c, 0x7e, 0x19, 0x4d, 0xeb, 0xa6, 0x72, 0x58,
	0x50, 0x17, 0x6e, 0x5b, 0x57, 0xf5, 0x27, 0xd9, 0xb9, 0x29, 0x29, 0xd0, 0x66, 0x0f, 0x03, 0x56,
	0x53, 0x4c, 0x19, 0xf6, 0xa7, 0xb9, 0xd2, 0x61, 0x8c, 0xb6, 0xeb, 0xd7, 0x97, 0x7a, 0x56, 0xf8,
	0x99, 0x95, 0xdf, 0x0b, 0xd7, 0xd0, 0xa4, 0x14, 0x0c, 0x9c, 0x3d, 0xc2, 0x7e, 0x9e, 0x27, 0x79,
	0x0a, 0x61, 0x19, 0xfc, 0x6b, 0x68, 0xae, 0x2e, 0x18, 0xed, 0x2f, 0x37, 0x1f, 0xee, 0x5f, 0xf2,
	0x45, 0x76, 0x26, 0x66, 0x22, 0xa0, 0x29, 0x36, 0x0b, 0x50, 0x6a, 0x13, 0x75, 0xee, 0x79, 0xcc,
	0x97, 0xe4, 0x56, 0x36, 0x41, 0xbc, 0x16, 0x53, 0x74, 0xbe, 0x5b, 0x00, 0x29, 0x8f, 0xd3, 0xb0,
	0x0a, 0xf9, 0x32, 0x9b, 0xed, 0x5c, 0x4c, 0x91, 0xb4, 0x40, 0xfd, 0xb2, 0x82, 0xaf, 0xa0, 0x29,
	0x28, 0x3c, 0xcf, 0x86, 0x41, 0xa3, 0xcd, 0x84, 0xec, 0x90, 0xaf, 0xb2, 0x19, 0x27, 0x15, 0xa2,
	0x1a, 0x01, 0x32, 0x45, 0xae, 0x92, 0xaf, 0xff, 0x9f, 0xc8, 0x55, 0x70, 0xc3, 0xf1, 0x1b, 0x21,
	0xf5, 0xa5, 0x0b, 0xfa, 0xbe, 0xc9, 0x71, 0x90, 0xdd, 0xd5, 0x78, 0x1f, 0x91, 0x61, 0xce, 0xa1,
	0x8f, 0xf2, 0x76, 0xb6, 0xac, 0xb9, 0xb4, 0x75, 0xe8, 0xd3, 0x84, 0xbe, 0xec, 0x9a, 0x07, 0x94,
	0xd6, 0x9d, 0x3c, 0xa5, 0x15, 0x03, 0xa0, 0xb4, 0xae, 0x21, 0x12, 0x97, 0xd6, 0x80, 0x63, 0xdc,
	0xcd, 0xd5, 0x22, 0x71, 0x69, 0x26, 0xed, 0x62, 0x1b, 0xe1, 0x74, 0xc9, 0x82, 0xba, 0x7b, 0x39,
	0xd4, 0x95, 0x12, 0x15, 0x0b, 0x12, 0xe1, 0xae, 0xee, 0x9b, 0x56, 0xfa, 0x9d, 0xe0, 0x7e, 0x8e,
	0xa4, 0x75, 0xa7, 0x99, 0x7e, 0x3b, 0x58, 0x47, 0xc5, 0x23, 0x3b, 0xd1, 0x92, 0x0f, 0x72, 0x28,
	0x9b, 0x38, 0x3a, 0xe8, 0xf5, 0xe4, 0x95, 0xf4, 0x10, 0xf6, 0xad, 0x75, 0xca, 0x29, 0x0c, 0xe6,
	0x2a, 0x5d, 0xed, 0x3e, 0x27, 0xdf, 0x65, 0x6f, 0xa4, 0xa0, 0xd6, 0x5a, 0x7c, 0xd8, 0x35, 0xb7,
	0x3a, 0x78, 0xcd, 0x7d, 0x7f, 0xea, 0x6b, 0x6e, 0x35, 0x75, 0xcd, 0x6d, 0xa1, 0x92, 0x1f, 0x7a,
	0x87, 0xb0, 0x41, 0xde, 0xb0, 0x55, 0xec, 0x80, 0xdc, 0xbe, 0x95, 0x63, 0xba, 0x9f, 0x32, 0xa8,
	0xdd, 0xc6, 0xb6, 0xc2, 0xe0, 0xb7, 0x60, 0x2a, 0xea, 0xd2, 0xb0, 0xe3, 0xb6, 0x2b, 0xa0, 0x22,
	0x62, 0x13, 0x0e, 0xc8, 0x9d, 0x3c, 0x94, 0x24, 0xa6, 0xdc, 0x32, 0xf8, 0xd8, 0x8e, 0x03, 0xfc,
	0x2a, 0xc2, 0x9e, 0xeb, 0xa7, 0x67, 0xf1, 0xbb, 0xff, 0x42, 0x9a, 0xb8, 0x61, 0xa6, 0x01, 0x98,
	0x1c, 0xc6, 0xaf, 0xc2, 0x5d, 0xd3, 0x95, 0x7a, 0x43, 0x4f, 0x22, 0xe4, 0x5e, 0x1e, 0x7d, 0xc5,
	0x58, 0x9f, 0x19, 0x5f, 0xb4, 0x2a, 0x7a, 0x9c, 0x56, 0x75, 0x3f, 0x9f, 0x2a, 0x7a, 0x9c, 0x54,
	0x05, 0x9f, 0x3d, 0x3d, 0x55, 0xb1, 0x31, 0x05, 0xe4, 0x41, 0x1e, 0x61, 0x33, 0xb1, 0xb0, 0xfd,
	0x18, 0x07, 0x95, 0x3f, 0x25, 0xb9, 0xa4, 0x2d, 0x3b, 0xbe, 0x8d, 0x7f, 0x38, 0xc9, 0x21, 0x6b,
	0x42, 0x63, 0x2c, 0x73, 0x27, 0x37, 0xd0, 0x85, 0x76, 0x34, 0xb4, 0x81, 0x53, 0xbb, 0xf5, 0x61,
	0x53, 0xf4, 0x4f, 0x27, 0x39, 0xe4, 0x2d, 0xc5, 0x34, 0x55, 0xc5, 0x32, 0x30, 0x50, 0x6f, 0xa2,
	0xa2, 0x91, 0xda, 0xb5, 0xdb, 0x9f, 0xf3, 0xb0, 0x9a, 0xfd, 0xed, 0x45, 0x98, 0x8d, 0xeb, 0x68,
	0xd1, 0xe5, 0xa9, 0xc9, 0xb3, 0xf7, 0x39, 0xfe, 0xe6, 0x5a, 0x93, 0x07, 0xce, 0xdb, 0xf1, 0x7b,
	0xe7, 0xb4, 0x5f, 0xec, 0x87, 0x05, 0xfd, 0x51, 0xfc, 0xec, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3a, 0x00, 0x3a, 0x8a, 0xf1, 0x0f, 0x00, 0x00,
}
