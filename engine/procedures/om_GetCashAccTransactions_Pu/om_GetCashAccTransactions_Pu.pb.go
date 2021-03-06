// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetCashAccTransactions_Pu.proto
// DO NOT EDIT!

/*
Package om_GetCashAccTransactions_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetCashAccTransactions_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetCashAccTransactions_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue    `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                          `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue   `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                          `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue    `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                          `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	CashAccountTypeId              *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=cash_account_type_id,json=cashAccountTypeId" json:"cash_account_type_id,omitempty"`
	CashAccountTypeIdNull          bool                          `protobuf:"varint,1004,opt,name=cash_account_type_id_null,json=cashAccountTypeIdNull" json:"cash_account_type_id_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue    `protobuf:"bytes,5,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                          `protobuf:"varint,1005,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
	OnlyTransactionsSince          *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=only_transactions_since,json=onlyTransactionsSince" json:"only_transactions_since,omitempty"`
	OnlyTransactionsSinceNull      bool                          `protobuf:"varint,1006,opt,name=only_transactions_since_null,json=onlyTransactionsSinceNull" json:"only_transactions_since_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetCashAccountTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CashAccountTypeId
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
	}
	return nil
}

func (m *Parameters) GetOnlyTransactionsSince() *dstore_values.TimestampValue {
	if m != nil {
		return m.OnlyTransactionsSince
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	AccountBalance  *dstore_values.DecimalValue                      `protobuf:"bytes,101,opt,name=account_balance,json=accountBalance" json:"account_balance,omitempty"`
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

func (m *Response) GetAccountBalance() *dstore_values.DecimalValue {
	if m != nil {
		return m.AccountBalance
	}
	return nil
}

type Response_Row struct {
	RowId                  int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TransactionType        *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=transaction_type,json=transactionType" json:"transaction_type,omitempty"`
	PersonId               *dstore_values.IntegerValue   `protobuf:"bytes,10002,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	TransactionDateAndTime *dstore_values.TimestampValue `protobuf:"bytes,10003,opt,name=transaction_date_and_time,json=transactionDateAndTime" json:"transaction_date_and_time,omitempty"`
	TransactionValue       *dstore_values.DecimalValue   `protobuf:"bytes,10004,opt,name=transaction_value,json=transactionValue" json:"transaction_value,omitempty"`
	TransactionComment     *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=transaction_comment,json=transactionComment" json:"transaction_comment,omitempty"`
	TransactionTypeId      *dstore_values.IntegerValue   `protobuf:"bytes,10006,opt,name=transaction_type_id,json=transactionTypeId" json:"transaction_type_id,omitempty"`
	CashAccountTypeId      *dstore_values.IntegerValue   `protobuf:"bytes,10007,opt,name=cash_account_type_id,json=cashAccountTypeId" json:"cash_account_type_id,omitempty"`
	SuspendedUntil         *dstore_values.TimestampValue `protobuf:"bytes,10008,opt,name=suspended_until,json=suspendedUntil" json:"suspended_until,omitempty"`
	TransactionId          *dstore_values.IntegerValue   `protobuf:"bytes,10009,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTransactionType() *dstore_values.StringValue {
	if m != nil {
		return m.TransactionType
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetTransactionDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.TransactionDateAndTime
	}
	return nil
}

func (m *Response_Row) GetTransactionValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.TransactionValue
	}
	return nil
}

func (m *Response_Row) GetTransactionComment() *dstore_values.StringValue {
	if m != nil {
		return m.TransactionComment
	}
	return nil
}

func (m *Response_Row) GetTransactionTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TransactionTypeId
	}
	return nil
}

func (m *Response_Row) GetCashAccountTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CashAccountTypeId
	}
	return nil
}

func (m *Response_Row) GetSuspendedUntil() *dstore_values.TimestampValue {
	if m != nil {
		return m.SuspendedUntil
	}
	return nil
}

func (m *Response_Row) GetTransactionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetCashAccTransactions_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetCashAccTransactions_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetCashAccTransactions_Pu.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/om_GetCashAccTransactions_Pu.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xeb, 0x6e, 0xd3, 0x4a,
	0x10, 0x56, 0x4f, 0x4e, 0xd2, 0x74, 0x4f, 0x4f, 0xd2, 0xba, 0x37, 0x27, 0xed, 0xa9, 0x0e, 0x45,
	0x48, 0x88, 0x1f, 0x2e, 0x82, 0x1f, 0x14, 0x21, 0x01, 0xbd, 0x40, 0x15, 0x44, 0xa3, 0x62, 0x5a,
	0x24, 0x10, 0xc8, 0xda, 0xda, 0xdb, 0xb0, 0x22, 0xde, 0x0d, 0x5e, 0x9b, 0xaa, 0x6f, 0xc1, 0xfd,
	0xf2, 0x00, 0x3c, 0x47, 0x9f, 0x87, 0xeb, 0x33, 0x30, 0xeb, 0x71, 0x52, 0xc7, 0x4d, 0x9a, 0xc0,
	0x9f, 0x46, 0xbb, 0x3b, 0xdf, 0x37, 0xdf, 0xee, 0xcc, 0x37, 0x35, 0xb9, 0xe6, 0xa9, 0x50, 0x06,
	0x6c, 0x99, 0x89, 0x06, 0x17, 0x6c, 0xb9, 0x15, 0x48, 0x97, 0x79, 0x51, 0xc0, 0xd4, 0xb2, 0xf4,
	0x9d, 0x4d, 0x16, 0xae, 0x53, 0xf5, 0x74, 0xd5, 0x75, 0x77, 0x02, 0x2a, 0x14, 0x75, 0x43, 0x2e,
	0x85, 0x72, 0xb6, 0x23, 0x0b, 0xe2, 0x42, 0x69, 0x5c, 0x40, 0xb0, 0x85, 0x60, 0xeb, 0x34, 0x44,
	0x75, 0x2a, 0x49, 0xf4, 0x82, 0x36, 0x23, 0xa6, 0x90, 0xa0, 0x5a, 0xe9, 0xce, 0xce, 0x82, 0x40,
	0x06, 0xc9, 0xd1, 0x7c, 0xf7, 0x91, 0xcf, 0x94, 0xa2, 0x0d, 0x96, 0x1c, 0x9e, 0xcd, 0x1e, 0x86,
	0x94, 0x8b, 0x7d, 0x19, 0xf8, 0x54, 0xe7, 0xc3, 0xa0, 0xa5, 0xa3, 0x02, 0x21, 0xdb, 0x34, 0xa0,
	0x70, 0xca, 0x02, 0x65, 0x3c, 0x26, 0x0b, 0x2d, 0xf8, 0x95, 0xc2, 0xe1, 0x1e, 0x13, 0x21, 0xdf,
	0xe7, 0x6e, 0x1c, 0xed, 0xa0, 0x22, 0x73, 0xe4, 0xff, 0x91, 0xf3, 0xff, 0x5c, 0xaa, 0x5a, 0xc9,
	0x9d, 0x12, 0x9d, 0x2a, 0x0c, 0xb8, 0x68, 0x3c, 0xd0, 0x0b, 0xbb, 0x8a, 0xf8, 0x5a, 0x17, 0x3c,
	0x3e, 0x52, 0xc6, 0x1d, 0x72, 0xe6, 0x34, 0x76, 0x47, 0x44, 0xcd, 0xa6, 0xf9, 0x65, 0x14, 0x72,
	0x14, 0xed, 0xc5, 0xfe, 0x3c, 0x75, 0x08, 0x33, 0x56, 0x49, 0x29, 0xe1, 0x0a, 0x0f, 0x5b, 0x0c,
	0x08, 0xcd, 0xbf, 0x62, 0x6d, 0xf3, 0x19, 0x6d, 0x5c, 0x84, 0xac, 0xc1, 0x02, 0x14, 0x37, 0x8e,
	0x90, 0x1d, 0x40, 0xd4, 0x3c, 0xc3, 0x22, 0x53, 0xdd, 0x14, 0x28, 0xe0, 0x2b, 0x0a, 0x98, 0x48,
	0xc7, 0xc6, 0x29, 0xaf, 0x90, 0xb1, 0x48, 0xf0, 0xe7, 0x51, 0x9c, 0x2d, 0x37, 0xf0, 0x25, 0x8a,
	0x18, 0x0c, 0x89, 0xce, 0x91, 0x52, 0x07, 0x88, 0x39, 0xbe, 0x61, 0x8e, 0xf1, 0x76, 0x48, 0xcc,
	0x7f, 0x97, 0x4c, 0xbb, 0xd0, 0x17, 0x0e, 0x75, 0x5d, 0x19, 0x89, 0xb0, 0x73, 0xb1, 0xbf, 0x07,
	0x5f, 0x6c, 0xd2, 0xc5, 0x86, 0xd2, 0xb8, 0xe4, 0x76, 0x57, 0x49, 0xa5, 0x17, 0x1b, 0xe6, 0xff,
	0x8e, 0xf9, 0x67, 0x4e, 0xc0, 0x62, 0x21, 0xf7, 0xc8, 0x9c, 0x62, 0x2d, 0xe8, 0x0a, 0xc8, 0xe7,
	0xf0, 0xa4, 0x5a, 0xba, 0x48, 0xca, 0xcc, 0x0f, 0xbc, 0xf6, 0x74, 0x07, 0x5a, 0xc3, 0xea, 0xc1,
	0xb6, 0x32, 0x6e, 0x90, 0x85, 0x3e, 0x94, 0x28, 0xe8, 0x07, 0x0a, 0x32, 0x7b, 0x81, 0x63, 0x4d,
	0xbb, 0x64, 0x4e, 0x8a, 0xe6, 0xa1, 0x13, 0xa6, 0x2d, 0xa3, 0xb8, 0x70, 0x99, 0x59, 0x88, 0x35,
	0xfd, 0x97, 0xd1, 0x14, 0x72, 0xb0, 0x43, 0x48, 0xfd, 0x16, 0xca, 0x9a, 0xd1, 0xe8, 0xb4, 0xdf,
	0xee, 0x6b, 0xac, 0x71, 0x93, 0x2c, 0xf4, 0xa1, 0x45, 0x5d, 0x3f, 0x51, 0x57, 0xa5, 0x27, 0x5a,
	0x0b, 0x5b, 0xfa, 0x5c, 0x24, 0x45, 0x9b, 0xa9, 0x16, 0x6c, 0x32, 0xe3, 0x22, 0xc9, 0xc7, 0xfe,
	0xcc, 0x1a, 0x25, 0x31, 0x3f, 0x7a, 0xf7, 0x96, 0xfe, 0x6b, 0x63, 0xa0, 0xf1, 0x90, 0x4c, 0x68,
	0x67, 0x3a, 0x29, 0x6b, 0x42, 0x27, 0xe7, 0x00, 0x6c, 0x65, 0xc0, 0x59, 0x03, 0x6f, 0xc1, 0xba,
	0x76, 0xbc, 0xb6, 0xcb, 0x7e, 0xf7, 0x86, 0xb1, 0x42, 0x46, 0x93, 0x89, 0x00, 0xdd, 0xaa, 0x19,
	0x17, 0x4f, 0x30, 0xe2, 0xbc, 0xd8, 0xc2, 0x5f, 0xbb, 0x1d, 0x0e, 0x46, 0xcd, 0x05, 0xf2, 0x00,
	0x1a, 0x4f, 0xa3, 0x56, 0xac, 0xe1, 0x27, 0x98, 0xd5, 0x7e, 0x09, 0xcb, 0x96, 0x07, 0xb6, 0x26,
	0x31, 0x36, 0x48, 0xb9, 0xdd, 0x82, 0x7b, 0xb4, 0x49, 0x75, 0xc1, 0x58, 0xcf, 0x86, 0xf6, 0x98,
	0xcb, 0x7d, 0xda, 0xc4, 0x72, 0x95, 0x12, 0xcc, 0x1a, 0x42, 0xaa, 0x47, 0x79, 0x92, 0x03, 0x4a,
	0x63, 0x96, 0x14, 0x80, 0x54, 0xbb, 0xe2, 0x65, 0x1d, 0x58, 0xf2, 0x76, 0x1e, 0x96, 0xd0, 0xed,
	0xb7, 0xc9, 0x44, 0xaa, 0x84, 0x71, 0xb3, 0x9b, 0xaf, 0xea, 0x03, 0x9b, 0xb5, 0x9c, 0x02, 0x69,
	0x03, 0x80, 0x6b, 0xc6, 0x3a, 0x23, 0xca, 0x7c, 0x5d, 0x1f, 0xec, 0xbc, 0x62, 0x7b, 0x4e, 0x41,
	0x25, 0x2b, 0x69, 0x09, 0x1e, 0x0d, 0x99, 0x43, 0x85, 0xe7, 0xe8, 0x46, 0x34, 0xdf, 0xd4, 0x87,
	0x69, 0xd2, 0xd9, 0x14, 0xc1, 0x06, 0xe0, 0x57, 0x85, 0xb7, 0x03, 0x11, 0x46, 0x8d, 0x4c, 0xa6,
	0xa9, 0x63, 0xb0, 0xf9, 0xb6, 0x3e, 0xf8, 0x19, 0xd3, 0x8f, 0x12, 0xef, 0xc0, 0x90, 0x99, 0x4a,
	0x53, 0xb9, 0xd2, 0xf7, 0xc1, 0x66, 0xe6, 0xbb, 0xc1, 0x6f, 0x65, 0xa4, 0x70, 0xeb, 0x08, 0xcb,
	0xb2, 0xb5, 0x27, 0xd6, 0xfb, 0x21, 0x1e, 0x6e, 0x32, 0xf3, 0xf4, 0xf0, 0x82, 0x5b, 0x7d, 0x06,
	0xe0, 0x87, 0xfa, 0x1f, 0x4d, 0xc0, 0x4d, 0x52, 0x56, 0x91, 0x6a, 0x31, 0xe1, 0x31, 0xcf, 0x81,
	0x6d, 0xde, 0x34, 0x3f, 0x0e, 0x55, 0x86, 0x52, 0x07, 0xb6, 0xab, 0x51, 0xc6, 0x3a, 0x29, 0xa5,
	0x6f, 0x09, 0x8a, 0x3e, 0x0d, 0xa1, 0xe8, 0xdf, 0x14, 0xa6, 0xe6, 0xad, 0x3d, 0x21, 0xf3, 0x5c,
	0x66, 0xac, 0x74, 0xfc, 0x25, 0xf1, 0xe8, 0x7a, 0x43, 0x2a, 0xef, 0x59, 0xfb, 0xdc, 0xfb, 0xdd,
	0x8f, 0x8d, 0xbd, 0x42, 0xfc, 0xff, 0xfc, 0xf2, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x64,
	0x17, 0xda, 0xac, 0x08, 0x00, 0x00,
}
