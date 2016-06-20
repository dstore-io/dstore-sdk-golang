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

var fileDescriptor0 = []byte{
	// 1282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x57, 0x6b, 0x4f, 0x24, 0x45,
	0x17, 0xce, 0xbe, 0xbc, 0x0c, 0x9b, 0x02, 0x66, 0xa0, 0x60, 0xb1, 0x84, 0x65, 0xb3, 0xe2, 0x17,
	0x35, 0x66, 0x90, 0x71, 0x49, 0x36, 0x51, 0xd7, 0x05, 0xc1, 0x0d, 0x2a, 0xcd, 0x30, 0xe0, 0x26,
	0x1a, 0x4d, 0xa7, 0x98, 0xae, 0x81, 0x4e, 0xa6, 0xbb, 0x7a, 0xab, 0xab, 0x5d, 0xe6, 0x5f, 0x78,
	0xbf, 0xae, 0x97, 0xbd, 0xaa, 0x5f, 0x8d, 0xf1, 0x93, 0x3f, 0x43, 0x7f, 0x84, 0x78, 0xfb, 0x0b,
	0x9e, 0xaa, 0xea, 0x9e, 0x99, 0xee, 0x19, 0xed, 0x46, 0xbf, 0x40, 0x7a, 0xba, 0x9e, 0xe7, 0x3c,
	0x75, 0xea, 0x9c, 0xa7, 0x4e, 0xa3, 0x67, 0x9c, 0x50, 0x72, 0xc1, 0x96, 0x99, 0x7f, 0xe8, 0xfa,
	0x6c, 0x39, 0x10, 0xbc, 0xc9, 0x9c, 0x48, 0xb0, 0x70, 0x99, 0x7b, 0xf6, 0x35, 0x26, 0xeb, 0x91,
	0x68, 0x1e, 0xd1, 0x90, 0xed, 0x08, 0x87, 0x89, 0xdd, 0x88, 0x45, 0x2c, 0xb4, 0xd7, 0x9c, 0x2a,
	0xac, 0x93, 0x1c, 0x3f, 0x61, 0xc0, 0x55, 0x03, 0xae, 0xfe, 0x13, 0x62, 0x7e, 0x26, 0x0e, 0xf4,
	0x16, 0x6d, 0xc3, 0x2f, 0x86, 0x60, 0xfe, 0xe1, 0x74, 0x74, 0x26, 0x04, 0x17, 0xf1, 0xab, 0x85,
	0xf4, 0x2b, 0x8f, 0x85, 0x21, 0x3d, 0x64, 0xf1, 0xcb, 0x47, 0xb3, 0x2f, 0x25, 0x75, 0xfd, 0x16,
	0x17, 0x1e, 0x95, 0x2e, 0xf7, 0xcd, 0xa2, 0xa5, 0x9f, 0x4b, 0x08, 0xd5, 0xa9, 0xa0, 0xf0, 0x96,
	0x89, 0x10, 0x3f, 0x8b, 0xc6, 0xc3, 0x28, 0x08, 0xda, 0x2e, 0x13, 0xb6, 0xeb, 0x90, 0x33, 0x17,
	0xcf, 0x3c, 0x36, 0x5e, 0x5b, 0xa8, 0xc6, 0x5b, 0x88, 0x65, 0xb9, 0xbe, 0x64, 0x87, 0x4c, 0x5c,
	0x57, 0x4f, 0x0d, 0x94, 0xac, 0xdf, 0x72, 0xf0, 0xe3, 0x68, 0xaa, 0x0f, 0x6d, 0xfb, 0x51, 0xbb,
	0x4d, 0x7e, 0x19, 0x03, 0x8e, 0xb3, 0x8d, 0x72, 0x6f, 0x99, 0x05, 0x3f, 0xe3, 0x6d, 0x34, 0xd3,
	0x12, 0x90, 0x0b, 0xae, 0x12, 0x60, 0x3b, 0x8c, 0x3a, 0x6d, 0x10, 0x49, 0xfe, 0xa7, 0x03, 0x2e,
	0x66, 0x02, 0x4a, 0x17, 0x76, 0x26, 0xa9, 0x17, 0x98, 0x90, 0xd3, 0x0a, 0xa9, 0x33, 0xb7, 0x11,
	0xe3, 0xf0, 0x65, 0x44, 0x86, 0xd0, 0x19, 0x05, 0x27, 0x46, 0xc1, 0xb9, 0x01, 0x94, 0x16, 0xb2,
	0x85, 0xa6, 0x25, 0xcf, 0xca, 0x18, 0x29, 0x22, 0xa3, 0x22, 0x79, 0x5a, 0xc4, 0x25, 0x34, 0x37,
	0x40, 0x65, 0x24, 0xfc, 0x6a, 0x24, 0xcc, 0x64, 0x10, 0x5a, 0xc0, 0xf3, 0x68, 0xd2, 0x40, 0x64,
	0x27, 0x60, 0x2a, 0xe9, 0xff, 0xcf, 0x4f, 0xfa, 0xb8, 0x46, 0xec, 0x03, 0x00, 0xb2, 0xfe, 0x24,
	0xc2, 0x29, 0x02, 0x13, 0xf2, 0x37, 0x13, 0xb2, 0xd2, 0xb7, 0x52, 0x87, 0xbb, 0x84, 0xc6, 0x7c,
	0xee, 0xe8, 0x40, 0xa3, 0xf9, 0x81, 0x4a, 0x6a, 0x2d, 0xc4, 0x78, 0x04, 0x4d, 0xc4, 0x28, 0xc3,
	0xfe, 0xbb, 0x61, 0x47, 0xe6, 0xb5, 0x26, 0xde, 0x44, 0x53, 0x87, 0x4c, 0xda, 0x61, 0xe4, 0x79,
	0x54, 0x74, 0x6c, 0xee, 0xb7, 0x3b, 0xa4, 0x94, 0x1f, 0xa1, 0x0c, 0xa0, 0x3d, 0x83, 0xd9, 0x01,
	0x08, 0xae, 0xa1, 0x73, 0x59, 0x1a, 0x13, 0xf2, 0x0f, 0x13, 0x12, 0xa7, 0xd7, 0xeb, 0xd0, 0xbb,
	0x68, 0x4e, 0xab, 0x83, 0x96, 0x12, 0xb4, 0x09, 0x85, 0xec, 0x86, 0xd2, 0x6d, 0xaa, 0x2d, 0x8e,
	0xe5, 0x0b, 0x98, 0x55, 0xd0, 0x17, 0x52, 0x48, 0xd8, 0xf0, 0x15, 0xb4, 0x30, 0x9c, 0xd2, 0x88,
	0xf9, 0xd3, 0x88, 0x21, 0xc3, 0xb0, 0x4a, 0xd2, 0xd2, 0x4f, 0x73, 0xe8, 0x6c, 0x83, 0x85, 0x01,
	0xf7, 0x43, 0x86, 0x9f, 0x42, 0xa3, 0xba, 0x6b, 0xe3, 0x7e, 0x9a, 0xaf, 0xa6, 0x2d, 0xc1, 0x74,
	0xf4, 0xa6, 0xfa, 0xdb, 0x30, 0x0b, 0xf1, 0x6b, 0x68, 0x4a, 0xf5, 0xab, 0xdd, 0xd7, 0xb0, 0xd0,
	0x1b, 0x23, 0x00, 0xae, 0x66, 0xc0, 0xd9, 0xb6, 0xde, 0x86, 0xe7, 0xad, 0xde, 0x73, 0xa3, 0xe2,
	0xa5, 0x7f, 0x80, 0x56, 0x19, 0x8b, 0x7d, 0x02, 0xca, 0x5c, 0x31, 0x5e, 0x18, 0x60, 0x34, 0x2e,
	0xb2, 0x6d, 0xfe, 0x37, 0x92, 0xe5, 0xf8, 0x25, 0x34, 0x22, 0xf8, 0x4d, 0xa8, 0x4f, 0x85, 0xba,
	0x5c, 0x2d, 0xee, 0x6b, 0xd5, 0x24, 0x13, 0xd5, 0x06, 0xbf, 0xd9, 0x50, 0x24, 0xf3, 0xdf, 0xcf,
	0xa2, 0x11, 0x78, 0xc0, 0x73, 0xa8, 0x04, 0x8f, 0xea, 0xa8, 0xde, 0xb6, 0x20, 0x39, 0xa3, 0x8d,
	0x51, 0x78, 0x84, 0xfc, 0xb7, 0xd0, 0x92, 0x76, 0x8c, 0xf8, 0x00, 0x6c, 0x38, 0xb9, 0x9a, 0x0d,
	0x8e, 0x2b, 0x85, 0x7d, 0xd0, 0xb1, 0x03, 0x2a, 0x21, 0xdf, 0x3e, 0x79, 0xc7, 0x4a, 0x27, 0x34,
	0x3e, 0x5f, 0x58, 0xe5, 0xfa, 0x87, 0xe6, 0x78, 0x17, 0x35, 0x8d, 0x39, 0x23, 0xf8, 0xa5, 0xd6,
	0x50, 0x1c, 0xeb, 0x9d, 0xba, 0x61, 0xc0, 0x75, 0x34, 0xd7, 0xa6, 0xa1, 0xb4, 0x99, 0xe3, 0x4a,
	0xe6, 0x28, 0xee, 0x28, 0x34, 0xde, 0xf7, 0xae, 0x95, 0x5f, 0x3b, 0x58, 0x61, 0x37, 0x35, 0x74,
	0xbd, 0xf3, 0x6a, 0xa8, 0x4d, 0x70, 0x03, 0x55, 0x9a, 0x91, 0x10, 0xcc, 0x6f, 0x76, 0xec, 0xb0,
	0xe3, 0x1d, 0xf0, 0x36, 0x79, 0x2f, 0x5f, 0x66, 0x39, 0xc1, 0xec, 0x69, 0x08, 0x7e, 0x13, 0x9d,
	0xef, 0xd7, 0x45, 0xa5, 0xed, 0x50, 0xc9, 0x6c, 0xea, 0x3b, 0xb6, 0xb2, 0x21, 0xf2, 0xbe, 0x55,
	0xc4, 0xa2, 0x1e, 0xea, 0xe9, 0x5b, 0x93, 0x1b, 0x40, 0xb0, 0xe6, 0x3b, 0xfb, 0xb0, 0x04, 0x3f,
	0x97, 0xf6, 0xf9, 0x0f, 0xac, 0xd3, 0x19, 0xfd, 0x2a, 0x98, 0x88, 0x69, 0x52, 0xf2, 0xe1, 0x70,
	0xa8, 0xc3, 0x9a, 0xae, 0x47, 0xdb, 0x89, 0x8b, 0xe8, 0x9e, 0xc5, 0x2f, 0xa2, 0x29, 0xdd, 0x54,
	0x0e, 0x0b, 0x9b, 0xc2, 0x0d, 0x74, 0x55, 0x7f, 0x94, 0x9f, 0x9b, 0x8a, 0x02, 0x6d, 0xf4, 0x30,
	0x60, 0x35, 0xe5, 0x8c, 0x61, 0x7f, 0x5c, 0x28, 0x1d, 0xc6, 0x68, 0xbb, 0x7e, 0xbd, 0xda, 0xb3,
	0xc2, 0x4f, 0xac, 0xe2, 0x5e, 0x78, 0x05, 0x4d, 0x48, 0xc1, 0xc0, 0xd9, 0x63, 0xec, 0xa7, 0x45,
	0x92, 0xa7, 0x10, 0x96, 0xc1, 0xbf, 0x82, 0x66, 0x9b, 0x82, 0xd1, 0xfe, 0x72, 0xf3, 0xe1, 0xfe,
	0x25, 0x9f, 0xe5, 0x67, 0x62, 0x3a, 0x06, 0x9a, 0x62, 0xb3, 0x00, 0xa5, 0x36, 0xd1, 0xe4, 0x9e,
	0xc7, 0x7c, 0x49, 0x6e, 0xe5, 0x13, 0x24, 0x6b, 0x31, 0x45, 0x17, 0xba, 0x05, 0x90, 0xf1, 0x38,
	0x0d, 0xab, 0x91, 0xcf, 0xf3, 0xd9, 0xce, 0x27, 0x14, 0x69, 0x0b, 0xd4, 0x2f, 0x6b, 0xf8, 0x2a,
	0x9a, 0x84, 0xc2, 0xf3, 0x6c, 0x18, 0x34, 0x02, 0x26, 0x64, 0x87, 0x7c, 0x91, 0xcf, 0x38, 0xa1,
	0x10, 0xf5, 0x18, 0x90, 0x2b, 0x72, 0x85, 0x7c, 0xf9, 0xdf, 0x44, 0xae, 0x80, 0x1b, 0x9e, 0xbd,
	0x11, 0x51, 0x5f, 0xba, 0xa0, 0xef, 0xab, 0x02, 0x07, 0xd9, 0x5d, 0x8d, 0xf7, 0x10, 0x19, 0xe6,
	0x1c, 0xfa, 0x28, 0x6f, 0xe7, 0xcb, 0x9a, 0xcd, 0x5a, 0x87, 0x3e, 0x4d, 0xe8, 0xcb, 0xae, 0x79,
	0x40, 0x69, 0xdd, 0x29, 0x52, 0x5a, 0x09, 0x00, 0x4a, 0xeb, 0x3a, 0x22, 0x49, 0x69, 0x0d, 0x38,
	0xc6, 0xdd, 0x42, 0x2d, 0x92, 0x94, 0x66, 0xda, 0x2e, 0xb6, 0x10, 0xce, 0x96, 0x2c, 0xa8, 0xbb,
	0x57, 0x40, 0x5d, 0x25, 0x55, 0xb1, 0x20, 0x11, 0xee, 0xea, 0xbe, 0x69, 0xa5, 0xdf, 0x09, 0xee,
	0x17, 0x48, 0x5a, 0x77, 0x9a, 0xe9, 0xb7, 0x83, 0x35, 0x54, 0x3e, 0xb2, 0x53, 0x2d, 0xf9, 0xa0,
	0x80, 0xb2, 0xf1, 0xa3, 0xfd, 0x5e, 0x4f, 0x5e, 0xcd, 0x0e, 0x61, 0x5f, 0x5b, 0xa7, 0x9c, 0xc2,
	0x60, 0xae, 0xd2, 0xd5, 0xee, 0x73, 0xf2, 0x4d, 0xfe, 0x46, 0x4a, 0x6a, 0xad, 0xc5, 0x87, 0x5d,
	0x73, 0x2b, 0x83, 0xd7, 0xdc, 0xb7, 0xa7, 0xbe, 0xe6, 0x56, 0x32, 0xd7, 0xdc, 0x26, 0xaa, 0xf8,
	0x91, 0x77, 0x00, 0x1b, 0xe4, 0x2d, 0x5b, 0xc5, 0x0e, 0xc9, 0xed, 0x5b, 0x05, 0xa6, 0xfb, 0x49,
	0x83, 0xda, 0x69, 0x6d, 0x29, 0x0c, 0x7e, 0x03, 0xa6, 0xa2, 0x2e, 0x0d, 0x3b, 0x0e, 0x5c, 0x01,
	0x15, 0x91, 0x98, 0x70, 0x48, 0xee, 0x14, 0xa1, 0x24, 0x09, 0xe5, 0xa6, 0xc1, 0x27, 0x76, 0x1c,
	0xe2, 0x97, 0x11, 0xf6, 0x5c, 0x3f, 0x3b, 0x8b, 0xdf, 0xfd, 0x1b, 0xd2, 0xd4, 0x0d, 0x33, 0x05,
	0xc0, 0xf4, 0x30, 0x7e, 0x0d, 0xee, 0x9a, 0xae, 0xd4, 0x1b, 0x7a, 0x12, 0x21, 0xf7, 0x8a, 0xe8,
	0x2b, 0x27, 0xfa, 0xcc, 0xf8, 0xa2, 0x55, 0xd1, 0xe3, 0xac, 0xaa, 0xfb, 0xc5, 0x54, 0xd1, 0xe3,
	0xb4, 0x2a, 0xf8, 0xec, 0xe9, 0xa9, 0x4a, 0x8c, 0x29, 0x24, 0x0f, 0x8a, 0x08, 0x9b, 0x4e, 0x84,
	0xed, 0x25, 0x38, 0xa8, 0xfc, 0x49, 0xc9, 0x25, 0x6d, 0xdb, 0xc9, 0x6d, 0xfc, 0xdd, 0x49, 0x01,
	0x59, 0xe3, 0x1a, 0x63, 0x99, 0x3b, 0xb9, 0x85, 0x2e, 0x06, 0xf1, 0xd0, 0x06, 0x4e, 0xed, 0x36,
	0x87, 0x4d, 0xd1, 0x3f, 0x9c, 0x14, 0x90, 0xb7, 0x98, 0xd0, 0xd4, 0x15, 0xcb, 0xc0, 0x40, 0xbd,
	0x81, 0xca, 0x46, 0x6a, 0xd7, 0x6e, 0x7f, 0x2c, 0xc2, 0x6a, 0xf6, 0xb7, 0x1b, 0x63, 0xd6, 0xf7,
	0xd1, 0x82, 0xcb, 0x33, 0x93, 0x67, 0xef, 0x73, 0xfc, 0xf5, 0xd5, 0x7f, 0xf5, 0xa1, 0x7e, 0x50,
	0xd2, 0xdf, 0xc2, 0x4f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x7e, 0xff, 0xbc, 0xe8, 0x0f,
	0x00, 0x00,
}
