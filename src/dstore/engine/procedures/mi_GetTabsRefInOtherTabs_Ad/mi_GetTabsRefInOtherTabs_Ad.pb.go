// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetTabsRefInOtherTabs_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetTabsRefInOtherTabs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetTabsRefInOtherTabs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetTabsRefInOtherTabs_Ad

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
	TableId                         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	TableIdNull                     bool                        `protobuf:"varint,1001,opt,name=table_id_null,json=tableIdNull" json:"table_id_null,omitempty"`
	UsedInSearchItems               *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=used_in_search_items,json=usedInSearchItems" json:"used_in_search_items,omitempty"`
	UsedInSearchItemsNull           bool                        `protobuf:"varint,1002,opt,name=used_in_search_items_null,json=usedInSearchItemsNull" json:"used_in_search_items_null,omitempty"`
	UsedInSearchItemLacks           *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=used_in_search_item_lacks,json=usedInSearchItemLacks" json:"used_in_search_item_lacks,omitempty"`
	UsedInSearchItemLacksNull       bool                        `protobuf:"varint,1003,opt,name=used_in_search_item_lacks_null,json=usedInSearchItemLacksNull" json:"used_in_search_item_lacks_null,omitempty"`
	UsedInRatingValues              *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=used_in_rating_values,json=usedInRatingValues" json:"used_in_rating_values,omitempty"`
	UsedInRatingValuesNull          bool                        `protobuf:"varint,1004,opt,name=used_in_rating_values_null,json=usedInRatingValuesNull" json:"used_in_rating_values_null,omitempty"`
	UsedInRatingSubjects            *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=used_in_rating_subjects,json=usedInRatingSubjects" json:"used_in_rating_subjects,omitempty"`
	UsedInRatingSubjectsNull        bool                        `protobuf:"varint,1005,opt,name=used_in_rating_subjects_null,json=usedInRatingSubjectsNull" json:"used_in_rating_subjects_null,omitempty"`
	UsedInMetaInformationTypes      *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=used_in_meta_information_types,json=usedInMetaInformationTypes" json:"used_in_meta_information_types,omitempty"`
	UsedInMetaInformationTypesNull  bool                        `protobuf:"varint,1006,opt,name=used_in_meta_information_types_null,json=usedInMetaInformationTypesNull" json:"used_in_meta_information_types_null,omitempty"`
	UsedInPersonCharacValueRefs     *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=used_in_person_charac_value_refs,json=usedInPersonCharacValueRefs" json:"used_in_person_charac_value_refs,omitempty"`
	UsedInPersonCharacValueRefsNull bool                        `protobuf:"varint,1007,opt,name=used_in_person_charac_value_refs_null,json=usedInPersonCharacValueRefsNull" json:"used_in_person_charac_value_refs_null,omitempty"`
	UsedInPredefinedBillContent     *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=used_in_predefined_bill_content,json=usedInPredefinedBillContent" json:"used_in_predefined_bill_content,omitempty"`
	UsedInPredefinedBillContentNull bool                        `protobuf:"varint,1008,opt,name=used_in_predefined_bill_content_null,json=usedInPredefinedBillContentNull" json:"used_in_predefined_bill_content_null,omitempty"`
	UsedInCashAccTransactions       *dstore_values.IntegerValue `protobuf:"bytes,9,opt,name=used_in_cash_acc_transactions,json=usedInCashAccTransactions" json:"used_in_cash_acc_transactions,omitempty"`
	UsedInCashAccTransactionsNull   bool                        `protobuf:"varint,1009,opt,name=used_in_cash_acc_transactions_null,json=usedInCashAccTransactionsNull" json:"used_in_cash_acc_transactions_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableId
	}
	return nil
}

func (m *Parameters) GetUsedInSearchItems() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInSearchItems
	}
	return nil
}

func (m *Parameters) GetUsedInSearchItemLacks() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInSearchItemLacks
	}
	return nil
}

func (m *Parameters) GetUsedInRatingValues() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInRatingValues
	}
	return nil
}

func (m *Parameters) GetUsedInRatingSubjects() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInRatingSubjects
	}
	return nil
}

func (m *Parameters) GetUsedInMetaInformationTypes() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInMetaInformationTypes
	}
	return nil
}

func (m *Parameters) GetUsedInPersonCharacValueRefs() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInPersonCharacValueRefs
	}
	return nil
}

func (m *Parameters) GetUsedInPredefinedBillContent() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInPredefinedBillContent
	}
	return nil
}

func (m *Parameters) GetUsedInCashAccTransactions() *dstore_values.IntegerValue {
	if m != nil {
		return m.UsedInCashAccTransactions
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
	RowId                       int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TableId                     *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	UsedInSearchItemLacks       *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=used_in_search_item_lacks,json=usedInSearchItemLacks" json:"used_in_search_item_lacks,omitempty"`
	UsedInRatingValues          *dstore_values.BooleanValue `protobuf:"bytes,10003,opt,name=used_in_rating_values,json=usedInRatingValues" json:"used_in_rating_values,omitempty"`
	UsedInPredefinedBillContent *dstore_values.BooleanValue `protobuf:"bytes,10004,opt,name=used_in_predefined_bill_content,json=usedInPredefinedBillContent" json:"used_in_predefined_bill_content,omitempty"`
	TableName                   *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	PrimaryKeyColumnName        *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=primary_key_column_name,json=primaryKeyColumnName" json:"primary_key_column_name,omitempty"`
	UsedInSearchItems           *dstore_values.BooleanValue `protobuf:"bytes,10007,opt,name=used_in_search_items,json=usedInSearchItems" json:"used_in_search_items,omitempty"`
	UsedInPersonCharacValueRefs *dstore_values.BooleanValue `protobuf:"bytes,10008,opt,name=used_in_person_charac_value_refs,json=usedInPersonCharacValueRefs" json:"used_in_person_charac_value_refs,omitempty"`
	UsedInMetaInformationTypes  *dstore_values.BooleanValue `protobuf:"bytes,10009,opt,name=used_in_meta_information_types,json=usedInMetaInformationTypes" json:"used_in_meta_information_types,omitempty"`
	UsedInRatingSubjects        *dstore_values.BooleanValue `protobuf:"bytes,10010,opt,name=used_in_rating_subjects,json=usedInRatingSubjects" json:"used_in_rating_subjects,omitempty"`
	UsedInCashAccTransactions   *dstore_values.BooleanValue `protobuf:"bytes,10011,opt,name=used_in_cash_acc_transactions,json=usedInCashAccTransactions" json:"used_in_cash_acc_transactions,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableId
	}
	return nil
}

func (m *Response_Row) GetUsedInSearchItemLacks() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInSearchItemLacks
	}
	return nil
}

func (m *Response_Row) GetUsedInRatingValues() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInRatingValues
	}
	return nil
}

func (m *Response_Row) GetUsedInPredefinedBillContent() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInPredefinedBillContent
	}
	return nil
}

func (m *Response_Row) GetTableName() *dstore_values.StringValue {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *Response_Row) GetPrimaryKeyColumnName() *dstore_values.StringValue {
	if m != nil {
		return m.PrimaryKeyColumnName
	}
	return nil
}

func (m *Response_Row) GetUsedInSearchItems() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInSearchItems
	}
	return nil
}

func (m *Response_Row) GetUsedInPersonCharacValueRefs() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInPersonCharacValueRefs
	}
	return nil
}

func (m *Response_Row) GetUsedInMetaInformationTypes() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInMetaInformationTypes
	}
	return nil
}

func (m *Response_Row) GetUsedInRatingSubjects() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInRatingSubjects
	}
	return nil
}

func (m *Response_Row) GetUsedInCashAccTransactions() *dstore_values.BooleanValue {
	if m != nil {
		return m.UsedInCashAccTransactions
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetTabsRefInOtherTabs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetTabsRefInOtherTabs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetTabsRefInOtherTabs_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 875 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x96, 0xeb, 0x6e, 0xe3, 0x44,
	0x18, 0x86, 0x55, 0x42, 0x9a, 0x74, 0x2a, 0x04, 0x0c, 0x6d, 0x71, 0x13, 0x7a, 0x50, 0x0b, 0x12,
	0xfc, 0x71, 0x11, 0x20, 0x0a, 0xe5, 0x07, 0x6a, 0x23, 0x84, 0x02, 0x6d, 0x5a, 0x39, 0xa5, 0x12,
	0x48, 0x30, 0x8c, 0xed, 0x49, 0xe2, 0xad, 0x63, 0x47, 0x33, 0xce, 0x56, 0xb9, 0x8b, 0x3d, 0x1f,
	0xef, 0x65, 0xef, 0x64, 0x2f, 0x60, 0xcf, 0xdb, 0x3b, 0xd8, 0xf1, 0x7c, 0x4e, 0x9b, 0x83, 0x4f,
	0xdd, 0x3f, 0xad, 0xc6, 0x33, 0xef, 0xfb, 0x3d, 0x19, 0x8f, 0xdf, 0x6f, 0xd0, 0x8e, 0x2d, 0x02,
	0x9f, 0xb3, 0x2d, 0xe6, 0xb5, 0x1d, 0x8f, 0x6d, 0xf5, 0xb8, 0x6f, 0x31, 0xbb, 0xcf, 0x99, 0xd8,
	0xea, 0x3a, 0xe4, 0x77, 0x16, 0x1c, 0x53, 0x53, 0x18, 0xac, 0x55, 0xf7, 0x0e, 0x83, 0x0e, 0xe3,
	0xe1, 0x88, 0xec, 0xda, 0xba, 0x5c, 0x16, 0xf8, 0xf8, 0x1b, 0xd0, 0xea, 0xa0, 0xd5, 0x53, 0x04,
	0x95, 0xcf, 0xa2, 0x32, 0xd7, 0xa9, 0xdb, 0x67, 0x02, 0xf4, 0x95, 0xe5, 0xf1, 0xda, 0x8c, 0x73,
	0x9f, 0x47, 0x53, 0xd5, 0xf1, 0xa9, 0x2e, 0x13, 0x82, 0xb6, 0x59, 0x34, 0xb9, 0x39, 0x39, 0x19,
	0x50, 0xc7, 0x6b, 0xf9, 0xbc, 0x4b, 0x03, 0xc7, 0xf7, 0x60, 0xd1, 0xc6, 0x39, 0x42, 0xe8, 0x88,
	0x72, 0x2a, 0x67, 0x19, 0x17, 0xf8, 0x47, 0x54, 0x0e, 0xa8, 0xe9, 0x32, 0xe2, 0xd8, 0xda, 0xcc,
	0xfa, 0xcc, 0xd7, 0xf3, 0xdf, 0x55, 0xf5, 0x08, 0x3f, 0x62, 0x72, 0xbc, 0x80, 0xb5, 0x19, 0x3f,
	0x09, 0x47, 0x46, 0x49, 0x2d, 0xae, 0xdb, 0x78, 0x13, 0x7d, 0x34, 0xd4, 0x11, 0xaf, 0xef, 0xba,
	0xda, 0xb3, 0x92, 0x54, 0x97, 0x8d, 0xf9, 0x68, 0x41, 0x43, 0x3e, 0xc3, 0xfb, 0x68, 0xa1, 0x2f,
	0x98, 0x4d, 0x1c, 0x8f, 0x08, 0x46, 0xb9, 0xd5, 0x21, 0x4e, 0xc0, 0xba, 0x42, 0xfb, 0x20, 0xbb,
	0xd0, 0xa7, 0xa1, 0xb0, 0xee, 0x35, 0x95, 0xac, 0x1e, 0xaa, 0xf0, 0xcf, 0x68, 0x39, 0xce, 0x0d,
	0xca, 0x3f, 0x87, 0xf2, 0x8b, 0x53, 0x32, 0x05, 0xf2, 0x57, 0xac, 0x94, 0xb8, 0xd4, 0x3a, 0x15,
	0x5a, 0x21, 0x9b, 0x66, 0xca, 0x76, 0x3f, 0x54, 0xe2, 0x3d, 0xb4, 0x9a, 0x68, 0x0b, 0x58, 0x2f,
	0x00, 0x6b, 0x39, 0x56, 0xaf, 0xd0, 0x1a, 0x68, 0x71, 0xe8, 0xc1, 0xe5, 0x7b, 0xf2, 0xda, 0x04,
	0x00, 0xb4, 0x0f, 0xb3, 0xb1, 0x30, 0xd8, 0x1a, 0x4a, 0xa7, 0x1e, 0x09, 0xfc, 0x0b, 0xaa, 0xc4,
	0xfa, 0x01, 0xcf, 0x4b, 0xe0, 0x59, 0x9a, 0x16, 0x2a, 0x18, 0x03, 0x7d, 0x3e, 0x21, 0x16, 0x7d,
	0xf3, 0x1a, 0xb3, 0x02, 0xa1, 0x15, 0xb3, 0x71, 0x16, 0x46, 0x5d, 0x9b, 0x91, 0x10, 0xff, 0x8a,
	0xbe, 0x48, 0xf0, 0x04, 0xa4, 0x57, 0x80, 0xa4, 0xc5, 0x89, 0x15, 0x14, 0xb9, 0xdc, 0xe5, 0xf0,
	0x48, 0x93, 0x91, 0x33, 0x4d, 0x82, 0x41, 0x4f, 0x6e, 0xd5, 0x6c, 0x36, 0x5b, 0x05, 0xec, 0x0f,
	0xa4, 0x41, 0xfd, 0x52, 0x7f, 0x1c, 0xca, 0xe5, 0x31, 0xdd, 0x4c, 0x2f, 0x00, 0xa0, 0xaf, 0x01,
	0x74, 0x35, 0xd9, 0x49, 0xe1, 0x9a, 0x68, 0x7d, 0xe8, 0xd6, 0x93, 0x5f, 0x98, 0xf4, 0xb0, 0x3a,
	0xf2, 0x7b, 0xb3, 0xe0, 0x3d, 0x10, 0xce, 0x5a, 0x42, 0x2b, 0x65, 0x03, 0x57, 0xa1, 0xcc, 0x91,
	0xb2, 0xa8, 0x29, 0x07, 0x98, 0x90, 0x7a, 0x7c, 0x88, 0xbe, 0xca, 0xaa, 0x01, 0xcc, 0x6f, 0x80,
	0x79, 0x2d, 0xc5, 0x4c, 0x41, 0x53, 0xb4, 0x76, 0x61, 0xc8, 0x99, 0xcd, 0x5a, 0x32, 0x40, 0x6c,
	0x62, 0x3a, 0xae, 0x4b, 0x2c, 0x5f, 0x62, 0x79, 0x81, 0x56, 0xce, 0xcf, 0x7c, 0xe1, 0xb0, 0x27,
	0x0d, 0x6a, 0xa0, 0x97, 0x07, 0xfd, 0xcb, 0x8c, 0x12, 0x80, 0xfc, 0x76, 0x1c, 0x39, 0xce, 0x4b,
	0x21, 0xff, 0x8b, 0x56, 0x86, 0x7e, 0x16, 0x15, 0x1d, 0x42, 0x2d, 0x8b, 0x04, 0x9c, 0x7a, 0x82,
	0x5a, 0xe1, 0x0b, 0x11, 0xda, 0x5c, 0x36, 0x70, 0xf4, 0x5d, 0xd6, 0xa4, 0x7e, 0xd7, 0xb2, 0x8e,
	0x47, 0xd4, 0xf8, 0x0f, 0xb4, 0x91, 0x6a, 0x0f, 0xb0, 0xe7, 0x00, 0xbb, 0x92, 0xe8, 0x13, 0xa2,
	0x6e, 0x3c, 0x9d, 0x43, 0x65, 0x83, 0x89, 0x9e, 0x1c, 0x33, 0xfc, 0x2d, 0x2a, 0xaa, 0x44, 0x8f,
	0xe2, 0xb6, 0xa2, 0x8f, 0x77, 0x0b, 0x48, 0xfb, 0xdf, 0xc2, 0xbf, 0x06, 0x2c, 0xc4, 0x7f, 0xa3,
	0x4f, 0x26, 0xcf, 0xa5, 0x8c, 0xd0, 0x82, 0x14, 0xeb, 0x13, 0xe2, 0xc9, 0xc8, 0x9f, 0x38, 0xa4,
	0xc6, 0xc7, 0xdd, 0xf1, 0x07, 0xf8, 0x27, 0x54, 0x8a, 0x7a, 0x88, 0x8c, 0xc1, 0xd0, 0x71, 0x75,
	0xca, 0x11, 0x3a, 0xcc, 0x01, 0xfc, 0x37, 0x86, 0xcb, 0x71, 0x1d, 0x15, 0xb8, 0x7f, 0x26, 0x53,
	0x2a, 0x54, 0x6d, 0xeb, 0xb9, 0x5b, 0x9e, 0x3e, 0xdc, 0x08, 0xdd, 0xf0, 0xcf, 0x8c, 0xd0, 0xa3,
	0xf2, 0xa4, 0x84, 0x0a, 0x72, 0x80, 0x97, 0xd0, 0xac, 0x1c, 0x86, 0x9d, 0xe8, 0x46, 0x43, 0xee,
	0x4d, 0xd1, 0x28, 0xca, 0xa1, 0xec, 0x35, 0xdb, 0x23, 0x3d, 0xea, 0x66, 0xe3, 0x0a, 0x4d, 0xea,
	0x24, 0x2d, 0xf6, 0x6f, 0xc5, 0x3b, 0x99, 0xbe, 0xef, 0x32, 0xea, 0xa5, 0xe6, 0xfe, 0x61, 0x52,
	0x66, 0xdf, 0xce, 0xe1, 0x19, 0x17, 0xda, 0x66, 0xf6, 0xe7, 0x77, 0x27, 0x87, 0x75, 0xea, 0xf7,
	0xb7, 0x83, 0x10, 0xec, 0xa2, 0x27, 0x7b, 0xbf, 0x76, 0xb7, 0x31, 0x7e, 0xfa, 0x22, 0x3b, 0x11,
	0xf0, 0x21, 0x95, 0x31, 0xa7, 0x96, 0x37, 0xe4, 0xea, 0xb0, 0x2f, 0xf4, 0xb8, 0xd3, 0xa5, 0x7c,
	0x40, 0x4e, 0xd9, 0x40, 0x32, 0xb9, 0xfd, 0xae, 0x07, 0x46, 0xf7, 0xb2, 0x8d, 0x16, 0x22, 0xed,
	0x9f, 0x6c, 0x50, 0x53, 0x4a, 0xe5, 0x79, 0x90, 0x70, 0x39, 0xb8, 0x9f, 0xe3, 0x87, 0xc6, 0xdc,
	0x0e, 0xac, 0x1c, 0xb1, 0xfb, 0xe0, 0x0a, 0x7b, 0x18, 0x9b, 0xbb, 0xff, 0x67, 0xb6, 0xa2, 0x87,
	0x39, 0x4a, 0xa4, 0xf5, 0xa2, 0x66, 0x72, 0x07, 0x7e, 0x94, 0xc3, 0x3a, 0xbe, 0x05, 0xff, 0x97,
	0x15, 0x95, 0x8f, 0x73, 0x58, 0x27, 0x67, 0xe5, 0x5e, 0x13, 0x55, 0x1d, 0x7f, 0x22, 0x02, 0x2e,
	0x6f, 0xcc, 0xff, 0xfc, 0xf0, 0x3e, 0x77, 0x69, 0x73, 0x56, 0xdd, 0x57, 0xbf, 0x7f, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x3a, 0x09, 0x4b, 0x83, 0x8a, 0x0b, 0x00, 0x00,
}
