// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPrices_Pu.proto
// DO NOT EDIT!

/*
Package om_GetPrices_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPrices_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPrices_Pu

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
	NodeIds                       *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=node_ids,json=nodeIds" json:"node_ids,omitempty"`
	NodeIdsNull                   bool                        `protobuf:"varint,1001,opt,name=node_ids_null,json=nodeIdsNull" json:"node_ids_null,omitempty"`
	Quantities                    *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=quantities" json:"quantities,omitempty"`
	QuantitiesNull                bool                        `protobuf:"varint,1002,opt,name=quantities_null,json=quantitiesNull" json:"quantities_null,omitempty"`
	PersonId                      *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	PersonIdNull                  bool                        `protobuf:"varint,1003,opt,name=person_id_null,json=personIdNull" json:"person_id_null,omitempty"`
	CurrencyId                    *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=currency_id,json=currencyId" json:"currency_id,omitempty"`
	CurrencyIdNull                bool                        `protobuf:"varint,1004,opt,name=currency_id_null,json=currencyIdNull" json:"currency_id_null,omitempty"`
	IsTreeNodeId                  *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=is_tree_node_id,json=isTreeNodeId" json:"is_tree_node_id,omitempty"`
	IsTreeNodeIdNull              bool                        `protobuf:"varint,1005,opt,name=is_tree_node_id_null,json=isTreeNodeIdNull" json:"is_tree_node_id_null,omitempty"`
	PriceNodeCharacteristicId     *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=price_node_characteristic_id,json=priceNodeCharacteristicId" json:"price_node_characteristic_id,omitempty"`
	PriceNodeCharacteristicIdNull bool                        `protobuf:"varint,1006,opt,name=price_node_characteristic_id_null,json=priceNodeCharacteristicIdNull" json:"price_node_characteristic_id_null,omitempty"`
	ComputeSum                    *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=compute_sum,json=computeSum" json:"compute_sum,omitempty"`
	ComputeSumNull                bool                        `protobuf:"varint,1007,opt,name=compute_sum_null,json=computeSumNull" json:"compute_sum_null,omitempty"`
	UniqueId                      *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                  bool                        `protobuf:"varint,1008,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	GetAdditionalPriceInfo        *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=get_additional_price_info,json=getAdditionalPriceInfo" json:"get_additional_price_info,omitempty"`
	GetAdditionalPriceInfoNull    bool                        `protobuf:"varint,1009,opt,name=get_additional_price_info_null,json=getAdditionalPriceInfoNull" json:"get_additional_price_info_null,omitempty"`
	DeliveryPersonId              *dstore_values.IntegerValue `protobuf:"bytes,10,opt,name=delivery_person_id,json=deliveryPersonId" json:"delivery_person_id,omitempty"`
	DeliveryPersonIdNull          bool                        `protobuf:"varint,1010,opt,name=delivery_person_id_null,json=deliveryPersonIdNull" json:"delivery_person_id_null,omitempty"`
	GetPricePerSingleNodeId       *dstore_values.BooleanValue `protobuf:"bytes,11,opt,name=get_price_per_single_node_id,json=getPricePerSingleNodeId" json:"get_price_per_single_node_id,omitempty"`
	GetPricePerSingleNodeIdNull   bool                        `protobuf:"varint,1011,opt,name=get_price_per_single_node_id_null,json=getPricePerSingleNodeIdNull" json:"get_price_per_single_node_id_null,omitempty"`
	PaymentTypeId                 *dstore_values.IntegerValue `protobuf:"bytes,12,opt,name=payment_type_id,json=paymentTypeId" json:"payment_type_id,omitempty"`
	PaymentTypeIdNull             bool                        `protobuf:"varint,1012,opt,name=payment_type_id_null,json=paymentTypeIdNull" json:"payment_type_id_null,omitempty"`
	ShippingTypeId                *dstore_values.IntegerValue `protobuf:"bytes,13,opt,name=shipping_type_id,json=shippingTypeId" json:"shipping_type_id,omitempty"`
	ShippingTypeIdNull            bool                        `protobuf:"varint,1013,opt,name=shipping_type_id_null,json=shippingTypeIdNull" json:"shipping_type_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *Parameters) GetQuantities() *dstore_values.StringValue {
	if m != nil {
		return m.Quantities
	}
	return nil
}

func (m *Parameters) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Parameters) GetCurrencyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CurrencyId
	}
	return nil
}

func (m *Parameters) GetIsTreeNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsTreeNodeId
	}
	return nil
}

func (m *Parameters) GetPriceNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PriceNodeCharacteristicId
	}
	return nil
}

func (m *Parameters) GetComputeSum() *dstore_values.BooleanValue {
	if m != nil {
		return m.ComputeSum
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetGetAdditionalPriceInfo() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetAdditionalPriceInfo
	}
	return nil
}

func (m *Parameters) GetDeliveryPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.DeliveryPersonId
	}
	return nil
}

func (m *Parameters) GetGetPricePerSingleNodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetPricePerSingleNodeId
	}
	return nil
}

func (m *Parameters) GetPaymentTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PaymentTypeId
	}
	return nil
}

func (m *Parameters) GetShippingTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ShippingTypeId
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
	RowId                          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TotalNettoPrice                *dstore_values.DecimalValue `protobuf:"bytes,10001,opt,name=total_netto_price,json=totalNettoPrice" json:"total_netto_price,omitempty"`
	PreciseAbsUnitGrossSurcharge   *dstore_values.DecimalValue `protobuf:"bytes,10002,opt,name=precise_abs_unit_gross_surcharge,json=preciseAbsUnitGrossSurcharge" json:"precise_abs_unit_gross_surcharge,omitempty"`
	PriceNodeCharacteristicId      *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=price_node_characteristic_id,json=priceNodeCharacteristicId" json:"price_node_characteristic_id,omitempty"`
	AbsoluteUnitNettoSurcharge     *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=absolute_unit_netto_surcharge,json=absoluteUnitNettoSurcharge" json:"absolute_unit_netto_surcharge,omitempty"`
	UnitGrossPrice                 *dstore_values.DecimalValue `protobuf:"bytes,10005,opt,name=unit_gross_price,json=unitGrossPrice" json:"unit_gross_price,omitempty"`
	AbsoluteTotalNetSurcharge      *dstore_values.DecimalValue `protobuf:"bytes,10006,opt,name=absolute_total_net_surcharge,json=absoluteTotalNetSurcharge" json:"absolute_total_net_surcharge,omitempty"`
	AbsoluteTotalGrossSurcharge    *dstore_values.DecimalValue `protobuf:"bytes,10007,opt,name=absolute_total_gross_surcharge,json=absoluteTotalGrossSurcharge" json:"absolute_total_gross_surcharge,omitempty"`
	UnitBruttoPrice                *dstore_values.DecimalValue `protobuf:"bytes,10008,opt,name=unit_brutto_price,json=unitBruttoPrice" json:"unit_brutto_price,omitempty"`
	PreciseUnitGrossPrice          *dstore_values.DecimalValue `protobuf:"bytes,10009,opt,name=precise_unit_gross_price,json=preciseUnitGrossPrice" json:"precise_unit_gross_price,omitempty"`
	UnitNetPrice                   *dstore_values.DecimalValue `protobuf:"bytes,10010,opt,name=unit_net_price,json=unitNetPrice" json:"unit_net_price,omitempty"`
	TotalBruttoPrice               *dstore_values.DecimalValue `protobuf:"bytes,10011,opt,name=total_brutto_price,json=totalBruttoPrice" json:"total_brutto_price,omitempty"`
	PreciseAbsUnitNetSurcharge     *dstore_values.DecimalValue `protobuf:"bytes,10012,opt,name=precise_abs_unit_net_surcharge,json=preciseAbsUnitNetSurcharge" json:"precise_abs_unit_net_surcharge,omitempty"`
	SurchargeReason                *dstore_values.StringValue  `protobuf:"bytes,10013,opt,name=surcharge_reason,json=surchargeReason" json:"surcharge_reason,omitempty"`
	NodeId                         *dstore_values.IntegerValue `protobuf:"bytes,10014,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId                     *dstore_values.IntegerValue `protobuf:"bytes,10015,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	AbsoluteTotalNettoSurcharge    *dstore_values.DecimalValue `protobuf:"bytes,10016,opt,name=absolute_total_netto_surcharge,json=absoluteTotalNettoSurcharge" json:"absolute_total_netto_surcharge,omitempty"`
	RelativeSurcharge              *dstore_values.DecimalValue `protobuf:"bytes,10017,opt,name=relative_surcharge,json=relativeSurcharge" json:"relative_surcharge,omitempty"`
	AbsoluteTotalBruttoSurcharge   *dstore_values.DecimalValue `protobuf:"bytes,10018,opt,name=absolute_total_brutto_surcharge,json=absoluteTotalBruttoSurcharge" json:"absolute_total_brutto_surcharge,omitempty"`
	SurchargeGeneratedByCampIds    *dstore_values.StringValue  `protobuf:"bytes,10019,opt,name=surcharge_generated_by_camp_ids,json=surchargeGeneratedByCampIds" json:"surcharge_generated_by_camp_ids,omitempty"`
	PreciseTotalGrossPrice         *dstore_values.DecimalValue `protobuf:"bytes,10020,opt,name=precise_total_gross_price,json=preciseTotalGrossPrice" json:"precise_total_gross_price,omitempty"`
	UnitNettoPrice                 *dstore_values.DecimalValue `protobuf:"bytes,10021,opt,name=unit_netto_price,json=unitNettoPrice" json:"unit_netto_price,omitempty"`
	TaxesMultiplier                *dstore_values.DecimalValue `protobuf:"bytes,10022,opt,name=taxes_multiplier,json=taxesMultiplier" json:"taxes_multiplier,omitempty"`
	PreciseAbsTotalGrossSurcharge  *dstore_values.DecimalValue `protobuf:"bytes,10023,opt,name=precise_abs_total_gross_surcharge,json=preciseAbsTotalGrossSurcharge" json:"precise_abs_total_gross_surcharge,omitempty"`
	PreciseAbsTotalNetSurcharge    *dstore_values.DecimalValue `protobuf:"bytes,10024,opt,name=precise_abs_total_net_surcharge,json=preciseAbsTotalNetSurcharge" json:"precise_abs_total_net_surcharge,omitempty"`
	TotalNetPrice                  *dstore_values.DecimalValue `protobuf:"bytes,10025,opt,name=total_net_price,json=totalNetPrice" json:"total_net_price,omitempty"`
	Quantity                       *dstore_values.IntegerValue `protobuf:"bytes,10026,opt,name=quantity" json:"quantity,omitempty"`
	QuantityPerBundleItemSetIdList *dstore_values.StringValue  `protobuf:"bytes,10027,opt,name=quantity_per_bundle_item_set_id_list,json=quantityPerBundleItemSetIdList" json:"quantity_per_bundle_item_set_id_list,omitempty"`
	SurchargeValue                 *dstore_values.DecimalValue `protobuf:"bytes,10028,opt,name=surcharge_value,json=surchargeValue" json:"surcharge_value,omitempty"`
	AbsoluteUnitNetSurcharge       *dstore_values.DecimalValue `protobuf:"bytes,10029,opt,name=absolute_unit_net_surcharge,json=absoluteUnitNetSurcharge" json:"absolute_unit_net_surcharge,omitempty"`
	AbsoluteUnitBruttoSurcharge    *dstore_values.DecimalValue `protobuf:"bytes,10030,opt,name=absolute_unit_brutto_surcharge,json=absoluteUnitBruttoSurcharge" json:"absolute_unit_brutto_surcharge,omitempty"`
	AbsoluteUnitGrossSurcharge     *dstore_values.DecimalValue `protobuf:"bytes,10031,opt,name=absolute_unit_gross_surcharge,json=absoluteUnitGrossSurcharge" json:"absolute_unit_gross_surcharge,omitempty"`
	PreciseTotalNetPrice           *dstore_values.DecimalValue `protobuf:"bytes,10032,opt,name=precise_total_net_price,json=preciseTotalNetPrice" json:"precise_total_net_price,omitempty"`
	SurchargeTypeId                *dstore_values.IntegerValue `protobuf:"bytes,10033,opt,name=surcharge_type_id,json=surchargeTypeId" json:"surcharge_type_id,omitempty"`
	PreciseUnitNetPrice            *dstore_values.DecimalValue `protobuf:"bytes,10034,opt,name=precise_unit_net_price,json=preciseUnitNetPrice" json:"precise_unit_net_price,omitempty"`
	TotalGrossPrice                *dstore_values.DecimalValue `protobuf:"bytes,10035,opt,name=total_gross_price,json=totalGrossPrice" json:"total_gross_price,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTotalNettoPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalNettoPrice
	}
	return nil
}

func (m *Response_Row) GetPreciseAbsUnitGrossSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseAbsUnitGrossSurcharge
	}
	return nil
}

func (m *Response_Row) GetPriceNodeCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PriceNodeCharacteristicId
	}
	return nil
}

func (m *Response_Row) GetAbsoluteUnitNettoSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteUnitNettoSurcharge
	}
	return nil
}

func (m *Response_Row) GetUnitGrossPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.UnitGrossPrice
	}
	return nil
}

func (m *Response_Row) GetAbsoluteTotalNetSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteTotalNetSurcharge
	}
	return nil
}

func (m *Response_Row) GetAbsoluteTotalGrossSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteTotalGrossSurcharge
	}
	return nil
}

func (m *Response_Row) GetUnitBruttoPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.UnitBruttoPrice
	}
	return nil
}

func (m *Response_Row) GetPreciseUnitGrossPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseUnitGrossPrice
	}
	return nil
}

func (m *Response_Row) GetUnitNetPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.UnitNetPrice
	}
	return nil
}

func (m *Response_Row) GetTotalBruttoPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalBruttoPrice
	}
	return nil
}

func (m *Response_Row) GetPreciseAbsUnitNetSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseAbsUnitNetSurcharge
	}
	return nil
}

func (m *Response_Row) GetSurchargeReason() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeReason
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

func (m *Response_Row) GetAbsoluteTotalNettoSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteTotalNettoSurcharge
	}
	return nil
}

func (m *Response_Row) GetRelativeSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.RelativeSurcharge
	}
	return nil
}

func (m *Response_Row) GetAbsoluteTotalBruttoSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteTotalBruttoSurcharge
	}
	return nil
}

func (m *Response_Row) GetSurchargeGeneratedByCampIds() *dstore_values.StringValue {
	if m != nil {
		return m.SurchargeGeneratedByCampIds
	}
	return nil
}

func (m *Response_Row) GetPreciseTotalGrossPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseTotalGrossPrice
	}
	return nil
}

func (m *Response_Row) GetUnitNettoPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.UnitNettoPrice
	}
	return nil
}

func (m *Response_Row) GetTaxesMultiplier() *dstore_values.DecimalValue {
	if m != nil {
		return m.TaxesMultiplier
	}
	return nil
}

func (m *Response_Row) GetPreciseAbsTotalGrossSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseAbsTotalGrossSurcharge
	}
	return nil
}

func (m *Response_Row) GetPreciseAbsTotalNetSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseAbsTotalNetSurcharge
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

func (m *Response_Row) GetQuantityPerBundleItemSetIdList() *dstore_values.StringValue {
	if m != nil {
		return m.QuantityPerBundleItemSetIdList
	}
	return nil
}

func (m *Response_Row) GetSurchargeValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.SurchargeValue
	}
	return nil
}

func (m *Response_Row) GetAbsoluteUnitNetSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteUnitNetSurcharge
	}
	return nil
}

func (m *Response_Row) GetAbsoluteUnitBruttoSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteUnitBruttoSurcharge
	}
	return nil
}

func (m *Response_Row) GetAbsoluteUnitGrossSurcharge() *dstore_values.DecimalValue {
	if m != nil {
		return m.AbsoluteUnitGrossSurcharge
	}
	return nil
}

func (m *Response_Row) GetPreciseTotalNetPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseTotalNetPrice
	}
	return nil
}

func (m *Response_Row) GetSurchargeTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SurchargeTypeId
	}
	return nil
}

func (m *Response_Row) GetPreciseUnitNetPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.PreciseUnitNetPrice
	}
	return nil
}

func (m *Response_Row) GetTotalGrossPrice() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalGrossPrice
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPrices_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPrices_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPrices_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 1497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x58, 0xf9, 0x6f, 0xdc, 0xd4,
	0x13, 0x57, 0xbf, 0xf9, 0xe6, 0xe8, 0x24, 0xe9, 0x6e, 0xdd, 0x34, 0x71, 0xb2, 0x4d, 0x7a, 0x81,
	0x54, 0x84, 0xb4, 0x41, 0x45, 0x40, 0x85, 0x00, 0xa9, 0x09, 0x25, 0x2c, 0xd0, 0x55, 0xd8, 0x34,
	0x45, 0x14, 0x90, 0xf1, 0xae, 0x5f, 0xb7, 0x96, 0x76, 0xed, 0xad, 0xfd, 0xdc, 0xb2, 0xff, 0x05,
	0xf7, 0x7d, 0x1f, 0xe5, 0xbe, 0x25, 0x24, 0xfe, 0x24, 0xee, 0xfb, 0x77, 0xe6, 0xbd, 0xb1, 0xfd,
	0x6c, 0x6f, 0x5a, 0xbf, 0x8a, 0x5f, 0xa0, 0x8e, 0xe7, 0xf3, 0xf9, 0xcc, 0x9b, 0x99, 0x37, 0x33,
	0x5e, 0xa8, 0x3b, 0x21, 0xf7, 0x03, 0xb6, 0xca, 0xbc, 0xae, 0xeb, 0xb1, 0xd5, 0x41, 0xe0, 0x77,
	0x98, 0x13, 0x05, 0x2c, 0x5c, 0xf5, 0xfb, 0xd6, 0x06, 0xe3, 0x9b, 0x81, 0xdb, 0x61, 0xa1, 0xb5,
	0x19, 0xd5, 0xf1, 0x15, 0xf7, 0x8d, 0x65, 0xb2, 0xaf, 0x93, 0x7d, 0xbd, 0x60, 0xb4, 0xb4, 0x2f,
	0xa6, 0xbb, 0x64, 0xf7, 0x22, 0x16, 0x12, 0x66, 0x69, 0x31, 0xaf, 0xc1, 0x82, 0xc0, 0x0f, 0xe2,
	0x57, 0xb5, 0xfc, 0xab, 0x3e, 0x0b, 0x43, 0xbb, 0xcb, 0xe2, 0x97, 0x47, 0x8b, 0x2f, 0xb9, 0xed,
	0x7a, 0xe7, 0xfd, 0xa0, 0x6f, 0x73, 0xd7, 0xf7, 0xc8, 0xe8, 0xc8, 0x77, 0x33, 0x00, 0x9b, 0x76,
	0x60, 0xe3, 0x5b, 0x16, 0x84, 0xc6, 0x6d, 0x30, 0xe5, 0xf9, 0x0e, 0xb3, 0x5c, 0x27, 0x34, 0x77,
	0x1d, 0xda, 0x75, 0x6c, 0xfa, 0xf8, 0x52, 0x7c, 0xc4, 0x7a, 0xec, 0x53, 0xc8, 0x03, 0xd7, 0xeb,
	0x9e, 0x15, 0x0f, 0xad, 0x49, 0x61, 0xdb, 0x70, 0x42, 0xe3, 0x28, 0xcc, 0x26, 0x30, 0xcb, 0x8b,
	0x7a, 0x3d, 0xf3, 0x87, 0x49, 0x04, 0x4f, 0xb5, 0xa6, 0x63, 0x83, 0x26, 0xfe, 0xcd, 0xb8, 0x13,
	0xe0, 0x62, 0x64, 0x7b, 0xdc, 0xe5, 0x2e, 0x0b, 0xcd, 0xff, 0x95, 0xb2, 0x67, 0xac, 0x8d, 0x63,
	0x50, 0x51, 0x4f, 0x24, 0xf1, 0x23, 0x49, 0xec, 0x51, 0x7f, 0x97, 0x2a, 0x27, 0x60, 0xf7, 0x00,
	0x4f, 0xe2, 0x7b, 0xe8, 0x8c, 0x39, 0x26, 0x45, 0x6a, 0x05, 0x11, 0xd7, 0xe3, 0xac, 0xcb, 0x02,
	0x52, 0x99, 0x22, 0xeb, 0x86, 0x63, 0xdc, 0x08, 0x7b, 0x52, 0x24, 0x49, 0xfc, 0x44, 0x12, 0x33,
	0x89, 0x89, 0x14, 0xb8, 0x0b, 0xa6, 0x3b, 0x51, 0x10, 0x30, 0xaf, 0x33, 0x14, 0x12, 0xff, 0x2f,
	0x97, 0x80, 0xc4, 0x1e, 0x45, 0x6e, 0x82, 0x6a, 0x06, 0x4d, 0x32, 0x3f, 0xc7, 0x27, 0x51, 0x66,
	0x52, 0x68, 0x0d, 0x2a, 0x6e, 0x68, 0xf1, 0x80, 0x31, 0x2b, 0x0e, 0xae, 0x39, 0xbe, 0xa3, 0x58,
	0xdb, 0xf7, 0x7b, 0xcc, 0xf6, 0x48, 0x6c, 0xc6, 0x0d, 0xcf, 0x20, 0xa4, 0x29, 0x03, 0x6f, 0xac,
	0xc2, 0x5c, 0x81, 0x83, 0x24, 0x7f, 0x21, 0xc9, 0x6a, 0xd6, 0x58, 0x8a, 0x3e, 0x0e, 0x07, 0x06,
	0xa2, 0x1c, 0xc9, 0xbc, 0x73, 0x01, 0x4b, 0xa3, 0x83, 0x95, 0xe1, 0x86, 0xdc, 0xed, 0x08, 0x0f,
	0x26, 0xca, 0x8f, 0xbb, 0x28, 0x09, 0x04, 0xe7, 0x7a, 0x0e, 0x8e, 0xee, 0x34, 0xe0, 0xf0, 0xb5,
	0xd8, 0xc9, 0xb7, 0x5f, 0xc9, 0xb7, 0xe5, 0xab, 0xd2, 0xa4, 0x69, 0xf0, 0xfb, 0x83, 0x88, 0x33,
	0x2b, 0x8c, 0xfa, 0xe6, 0x64, 0x79, 0x64, 0x20, 0xb6, 0xdf, 0x8a, 0xfa, 0x32, 0x0d, 0x0a, 0x4d,
	0xba, 0xbf, 0x25, 0x69, 0x48, 0xcd, 0xa4, 0xd0, 0x1d, 0xb0, 0x3b, 0xf2, 0xdc, 0x8b, 0x91, 0x4c,
	0xc0, 0x54, 0x69, 0xd5, 0x4e, 0x91, 0x31, 0xd5, 0x53, 0x0a, 0x24, 0x85, 0xdf, 0xe3, 0x7a, 0x4a,
	0x4c, 0x24, 0xff, 0x59, 0x58, 0xec, 0x32, 0x6e, 0xd9, 0x8e, 0xe3, 0x8a, 0x7b, 0x69, 0xf7, 0x2c,
	0x0a, 0x91, 0xb8, 0xab, 0xe6, 0xee, 0xf2, 0x63, 0xcd, 0x23, 0xfa, 0x64, 0x0a, 0x96, 0xbd, 0xa4,
	0x81, 0x50, 0x63, 0x1d, 0x56, 0xae, 0xca, 0x4b, 0xee, 0xfc, 0x41, 0xee, 0x2c, 0xed, 0x4c, 0x20,
	0x9d, 0x6b, 0x80, 0xe1, 0xb0, 0x9e, 0x7b, 0x89, 0x05, 0x43, 0x4b, 0x5d, 0x2b, 0x28, 0x2f, 0x82,
	0x6a, 0x02, 0xdb, 0x4c, 0xae, 0xd7, 0xed, 0xb0, 0x30, 0x4a, 0x45, 0x8e, 0xfc, 0x49, 0x8e, 0xcc,
	0x15, 0x31, 0xd2, 0x85, 0x73, 0x70, 0x40, 0x9c, 0x83, 0x9c, 0x47, 0xa0, 0x15, 0x62, 0xac, 0x7b,
	0xea, 0x4e, 0x4c, 0x97, 0x87, 0x68, 0xa1, 0x1b, 0x37, 0x59, 0x24, 0xde, 0x92, 0xe8, 0xf8, 0x7a,
	0x6c, 0xc0, 0xe1, 0x6b, 0x71, 0x93, 0x77, 0x7f, 0x91, 0x77, 0xb5, 0xab, 0x90, 0x48, 0x27, 0xd7,
	0xa1, 0x32, 0xb0, 0x87, 0x7d, 0xe6, 0x71, 0x8b, 0x0f, 0x07, 0xd2, 0xaf, 0x99, 0xf2, 0x20, 0xcd,
	0xc6, 0x98, 0x33, 0x08, 0x41, 0x6f, 0x6e, 0x81, 0xb9, 0x02, 0x09, 0x39, 0xf0, 0x37, 0x39, 0xb0,
	0x37, 0x67, 0x2d, 0x65, 0x4f, 0x41, 0x35, 0xbc, 0xe0, 0x0e, 0x06, 0xe8, 0x4e, 0xaa, 0x3b, 0x5b,
	0xae, 0xbb, 0x27, 0x01, 0xc5, 0xc2, 0xc7, 0x61, 0x7f, 0x91, 0x86, 0x94, 0xff, 0x21, 0x65, 0x23,
	0x6f, 0x2f, 0xa4, 0x8f, 0x7c, 0xbf, 0x08, 0x53, 0x2d, 0x16, 0x0e, 0x7c, 0x2f, 0x64, 0xe8, 0xf9,
	0xb8, 0x1c, 0x4b, 0xc5, 0x99, 0x11, 0x8f, 0x39, 0x1a, 0x59, 0xa7, 0xc4, 0x7f, 0x5b, 0x64, 0x68,
	0x3c, 0x0a, 0x55, 0x31, 0x90, 0xac, 0xcc, 0x44, 0xc2, 0x91, 0x30, 0x86, 0xe0, 0x7a, 0x01, 0x5c,
	0x9c, 0x5b, 0xa7, 0xf1, 0xb9, 0xa1, 0x9e, 0x5b, 0x95, 0x7e, 0xfe, 0x0f, 0x38, 0x01, 0x26, 0xe3,
	0x41, 0x88, 0xfd, 0x5f, 0x30, 0xae, 0x8c, 0x30, 0xd2, 0x98, 0x3c, 0x4d, 0xff, 0x6f, 0x25, 0xe6,
	0xc6, 0xdd, 0x30, 0x16, 0xf8, 0x97, 0xb1, 0xa5, 0x0b, 0xd4, 0xcd, 0xf5, 0x6b, 0xce, 0xea, 0x7a,
	0x72, 0xf8, 0x7a, 0xcb, 0xbf, 0xdc, 0x12, 0xb8, 0xa5, 0x2b, 0x0b, 0x30, 0x86, 0x0f, 0xc6, 0x3c,
	0x4c, 0xe0, 0xa3, 0xc8, 0xc5, 0xd3, 0x4d, 0x8c, 0xc7, 0x78, 0x6b, 0x1c, 0x1f, 0x31, 0xcc, 0xf7,
	0xc3, 0x5e, 0xee, 0x73, 0xbc, 0x88, 0x1e, 0xe3, 0xdc, 0xa7, 0xaa, 0x33, 0x9f, 0x69, 0xee, 0x98,
	0x2f, 0x87, 0x75, 0xdc, 0xbe, 0xdd, 0xa3, 0x7c, 0x55, 0x24, 0xac, 0x29, 0x50, 0x52, 0xdf, 0x70,
	0xe0, 0xd0, 0x20, 0x40, 0x8b, 0x90, 0x59, 0x76, 0x3b, 0xb4, 0xb0, 0x9f, 0x70, 0xab, 0x1b, 0xf8,
	0x61, 0x88, 0xdd, 0x2c, 0x10, 0x6d, 0x15, 0xcf, 0xfe, 0xac, 0x06, 0xf1, 0x81, 0x98, 0xe5, 0x64,
	0x3b, 0xdc, 0x46, 0x8e, 0x0d, 0x41, 0xb1, 0x95, 0x30, 0x18, 0x4f, 0x94, 0xcc, 0x82, 0xe7, 0x9a,
	0xff, 0x69, 0x18, 0x58, 0xb0, 0x8c, 0xce, 0xfb, 0x3d, 0xd1, 0x84, 0xe5, 0x09, 0x28, 0x2c, 0xea,
	0x04, 0xcf, 0x6b, 0x9c, 0x60, 0x29, 0xa1, 0x10, 0xfe, 0xcb, 0x08, 0x29, 0xff, 0xef, 0x83, 0x6a,
	0x26, 0x32, 0x14, 0xee, 0x17, 0x34, 0x38, 0x45, 0xdb, 0xa6, 0x58, 0x50, 0xb4, 0x31, 0x0e, 0xa9,
	0xa3, 0x69, 0x02, 0x33, 0x7e, 0xbe, 0xa8, 0xc1, 0xb9, 0x98, 0x30, 0x9c, 0x89, 0x53, 0xa9, 0xdc,
	0xb4, 0x61, 0xa5, 0x40, 0x5f, 0x4c, 0xe5, 0x4b, 0x1a, 0x02, 0xb5, 0x9c, 0x40, 0x21, 0x93, 0x58,
	0x79, 0x32, 0x12, 0xed, 0x20, 0x52, 0x95, 0xf7, 0xb2, 0x4e, 0xe5, 0x09, 0xd8, 0x9a, 0x44, 0x51,
	0x2c, 0xb6, 0xc1, 0x4c, 0x2a, 0x6f, 0x24, 0xb6, 0xaf, 0x68, 0x10, 0xee, 0x8f, 0xd1, 0xdb, 0xf9,
	0x10, 0xaf, 0xc9, 0x59, 0x29, 0x4b, 0x20, 0x26, 0x7b, 0x55, 0x83, 0x4c, 0x0c, 0x52, 0x91, 0x74,
	0xe2, 0x78, 0x00, 0x0c, 0x0a, 0x5f, 0xee, 0x94, 0xaf, 0x69, 0xf0, 0x54, 0x25, 0x2e, 0x7b, 0xcc,
	0x27, 0x61, 0x65, 0xe4, 0x82, 0xe5, 0x93, 0xfe, 0xba, 0x4e, 0x71, 0xe6, 0xaf, 0x57, 0x2e, 0xeb,
	0x58, 0x9c, 0x29, 0x99, 0x15, 0x30, 0x1b, 0x07, 0x9e, 0xf9, 0x46, 0xb3, 0x74, 0xbd, 0xa8, 0xa4,
	0xa0, 0x96, 0xc4, 0xe0, 0xc6, 0x3e, 0x99, 0x4c, 0xc2, 0x37, 0x35, 0xee, 0xe3, 0x04, 0x6d, 0xe4,
	0xc6, 0x3d, 0x30, 0x93, 0xdb, 0x2c, 0xdf, 0xd2, 0xc0, 0x02, 0x57, 0x8b, 0xe5, 0x68, 0xd1, 0x16,
	0x6f, 0xef, 0xdb, 0xd7, 0x5d, 0xb4, 0x85, 0xeb, 0xfb, 0x20, 0x18, 0x01, 0xeb, 0x61, 0x4f, 0xbf,
	0xc4, 0x32, 0xb4, 0xef, 0x68, 0xd0, 0xee, 0x4d, 0x70, 0x8a, 0xac, 0x03, 0x07, 0x0b, 0xfe, 0xc6,
	0x55, 0xa2, 0x98, 0xdf, 0xd5, 0x69, 0x98, 0x39, 0x87, 0xa9, 0x62, 0xb2, 0x37, 0xf9, 0xa0, 0xca,
	0x69, 0x97, 0x79, 0x2c, 0xb0, 0x39, 0x73, 0xac, 0xf6, 0xd0, 0xea, 0xd8, 0xfd, 0x81, 0xfc, 0xa8,
	0x7a, 0xaf, 0x3c, 0xc5, 0xb5, 0x94, 0x63, 0x23, 0xa1, 0x58, 0x1b, 0xae, 0x23, 0x81, 0xf8, 0xd2,
	0x7a, 0x04, 0x16, 0x93, 0xc2, 0xcc, 0xf6, 0x0a, 0xaa, 0xf5, 0xf7, 0x35, 0x4e, 0x30, 0x1f, 0xc3,
	0x55, 0x9b, 0xa0, 0x8a, 0x4f, 0x9a, 0x65, 0x76, 0x36, 0x7d, 0xa0, 0xdb, 0x2c, 0x33, 0xa3, 0x69,
	0x03, 0xaa, 0xdc, 0x7e, 0x0a, 0x47, 0x64, 0x3f, 0xea, 0x71, 0x77, 0xd0, 0x73, 0x59, 0x60, 0x5e,
	0xd1, 0x9a, 0x71, 0x02, 0x75, 0x3a, 0x05, 0x19, 0xe7, 0xc5, 0xb7, 0x82, 0xba, 0x82, 0x3b, 0x77,
	0xc6, 0x0f, 0x35, 0x98, 0x97, 0xd5, 0x2d, 0xdc, 0xa9, 0x37, 0xb6, 0xe1, 0xe0, 0xa8, 0x4e, 0xfe,
	0xae, 0x7f, 0xa4, 0x53, 0xca, 0x05, 0x95, 0xdc, 0x65, 0xbf, 0x17, 0x2a, 0x8a, 0x97, 0x62, 0xfb,
	0xb1, 0x06, 0xe7, 0x6c, 0x32, 0xf7, 0x29, 0xb4, 0x27, 0x60, 0x2a, 0xfe, 0xd8, 0x1d, 0x9a, 0x9f,
	0x68, 0xdc, 0xd7, 0xd4, 0xda, 0x70, 0xe1, 0x86, 0xe4, 0xdf, 0x72, 0xcd, 0x6d, 0x47, 0x9e, 0x83,
	0x6b, 0xae, 0xcb, 0x59, 0xdf, 0x0a, 0xd1, 0x25, 0xdc, 0xf7, 0x7a, 0x38, 0x96, 0xcd, 0x4f, 0xcb,
	0xab, 0x73, 0x25, 0x21, 0xc2, 0x35, 0x78, 0x4d, 0xd2, 0x34, 0x90, 0x65, 0x8b, 0xf1, 0x86, 0xf3,
	0x10, 0x52, 0xe0, 0x4a, 0xaa, 0x5a, 0x94, 0x25, 0xf1, 0xe6, 0x67, 0x3a, 0x65, 0x94, 0x82, 0xe4,
	0xb3, 0xf1, 0x18, 0xd4, 0x46, 0x96, 0x83, 0x4c, 0x46, 0x3e, 0xd7, 0xa0, 0x34, 0x0b, 0xab, 0xc1,
	0xce, 0x13, 0x37, 0x3b, 0x17, 0x15, 0xff, 0x17, 0xd7, 0xd3, 0xbc, 0xb6, 0xd3, 0x19, 0xa9, 0x24,
	0x46, 0x96, 0x9b, 0x62, 0xe5, 0x7e, 0x79, 0xbd, 0xcb, 0x4d, 0xa1, 0x6c, 0xb7, 0x60, 0x21, 0xdf,
	0x08, 0x54, 0x69, 0x7d, 0xa5, 0x41, 0x3d, 0x97, 0x6d, 0x03, 0x69, 0x85, 0xe1, 0x9e, 0xa0, 0x92,
	0x97, 0x7c, 0x50, 0x7c, 0xad, 0x51, 0x6a, 0x2a, 0xe7, 0xf1, 0x27, 0xc5, 0xc3, 0x30, 0x9f, 0xdb,
	0x13, 0x94, 0x77, 0xdf, 0x68, 0x78, 0xb7, 0x2f, 0xb3, 0x25, 0x64, 0x9d, 0x1b, 0x6d, 0x79, 0xdf,
	0x6a, 0xaf, 0xcf, 0xaa, 0xd7, 0xad, 0x6d, 0x40, 0xcd, 0xf5, 0x0b, 0xeb, 0xbd, 0xfa, 0xe9, 0xee,
	0xdc, 0x31, 0xdd, 0x1f, 0xf5, 0xda, 0x13, 0xf2, 0x47, 0xb4, 0x5b, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xde, 0xce, 0x22, 0xbf, 0x07, 0x14, 0x00, 0x00,
}
