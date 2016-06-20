// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetPersonPMStatistics_Ad.proto
// DO NOT EDIT!

/*
Package st_GetPersonPMStatistics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetPersonPMStatistics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetPersonPMStatistics_Ad

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
	PersonIds                      *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_ids,json=personIds" json:"person_ids,omitempty"`
	PersonIdsNull                  bool                        `protobuf:"varint,1001,opt,name=person_ids_null,json=personIdsNull" json:"person_ids_null,omitempty"`
	FromMonth                      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=from_month,json=fromMonth" json:"from_month,omitempty"`
	FromMonthNull                  bool                        `protobuf:"varint,1002,opt,name=from_month_null,json=fromMonthNull" json:"from_month_null,omitempty"`
	FromYear                       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=from_year,json=fromYear" json:"from_year,omitempty"`
	FromYearNull                   bool                        `protobuf:"varint,1003,opt,name=from_year_null,json=fromYearNull" json:"from_year_null,omitempty"`
	ToMonth                        *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=to_month,json=toMonth" json:"to_month,omitempty"`
	ToMonthNull                    bool                        `protobuf:"varint,1004,opt,name=to_month_null,json=toMonthNull" json:"to_month_null,omitempty"`
	ToYear                         *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=to_year,json=toYear" json:"to_year,omitempty"`
	ToYearNull                     bool                        `protobuf:"varint,1005,opt,name=to_year_null,json=toYearNull" json:"to_year_null,omitempty"`
	BasicCharacteristicNumbers     *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=basic_characteristic_numbers,json=basicCharacteristicNumbers" json:"basic_characteristic_numbers,omitempty"`
	BasicCharacteristicNumbersNull bool                        `protobuf:"varint,1006,opt,name=basic_characteristic_numbers_null,json=basicCharacteristicNumbersNull" json:"basic_characteristic_numbers_null,omitempty"`
	HTreeNodeIds                   *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=h_tree_node_ids,json=hTreeNodeIds" json:"h_tree_node_ids,omitempty"`
	HTreeNodeIdsNull               bool                        `protobuf:"varint,1007,opt,name=h_tree_node_ids_null,json=hTreeNodeIdsNull" json:"h_tree_node_ids_null,omitempty"`
	SummarizeMonths                *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=summarize_months,json=summarizeMonths" json:"summarize_months,omitempty"`
	SummarizeMonthsNull            bool                        `protobuf:"varint,1008,opt,name=summarize_months_null,json=summarizeMonthsNull" json:"summarize_months_null,omitempty"`
	GetTopX                        *dstore_values.IntegerValue `protobuf:"bytes,9,opt,name=get_top_x,json=getTopX" json:"get_top_x,omitempty"`
	GetTopXNull                    bool                        `protobuf:"varint,1009,opt,name=get_top_x_null,json=getTopXNull" json:"get_top_x_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIds() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIds
	}
	return nil
}

func (m *Parameters) GetFromMonth() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromMonth
	}
	return nil
}

func (m *Parameters) GetFromYear() *dstore_values.IntegerValue {
	if m != nil {
		return m.FromYear
	}
	return nil
}

func (m *Parameters) GetToMonth() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToMonth
	}
	return nil
}

func (m *Parameters) GetToYear() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToYear
	}
	return nil
}

func (m *Parameters) GetBasicCharacteristicNumbers() *dstore_values.StringValue {
	if m != nil {
		return m.BasicCharacteristicNumbers
	}
	return nil
}

func (m *Parameters) GetHTreeNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.HTreeNodeIds
	}
	return nil
}

func (m *Parameters) GetSummarizeMonths() *dstore_values.BooleanValue {
	if m != nil {
		return m.SummarizeMonths
	}
	return nil
}

func (m *Parameters) GetGetTopX() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetTopX
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
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value1RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value1_restricted_by_pattern,json=value1RestrictedByPattern" json:"value1_restricted_by_pattern,omitempty"`
	Month                     *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=month" json:"month,omitempty"`
	Value2RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value2_restricted_by_pattern,json=value2RestrictedByPattern" json:"value2_restricted_by_pattern,omitempty"`
	TotalValue                *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=total_value,json=totalValue" json:"total_value,omitempty"`
	Year                      *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=year" json:"year,omitempty"`
	PersonId                  *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=person_id,json=personId" json:"person_id,omitempty"`
	HTreeNodeId               *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	Value1                    *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=value1" json:"value1,omitempty"`
	Value2                    *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=value2" json:"value2,omitempty"`
	DirectValue               *dstore_values.DecimalValue `protobuf:"bytes,10010,opt,name=direct_value,json=directValue" json:"direct_value,omitempty"`
	BasicCharacteristicNumber *dstore_values.IntegerValue `protobuf:"bytes,10011,opt,name=basic_characteristic_number,json=basicCharacteristicNumber" json:"basic_characteristic_number,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue1RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value1RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetMonth() *dstore_values.IntegerValue {
	if m != nil {
		return m.Month
	}
	return nil
}

func (m *Response_Row) GetValue2RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value2RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetTotalValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalValue
	}
	return nil
}

func (m *Response_Row) GetYear() *dstore_values.IntegerValue {
	if m != nil {
		return m.Year
	}
	return nil
}

func (m *Response_Row) GetPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetDirectValue() *dstore_values.DecimalValue {
	if m != nil {
		return m.DirectValue
	}
	return nil
}

func (m *Response_Row) GetBasicCharacteristicNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNumber
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetPersonPMStatistics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetPersonPMStatistics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetPersonPMStatistics_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 849 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x5b, 0x8f, 0xdb, 0x44,
	0x14, 0x56, 0x49, 0x73, 0x3b, 0x49, 0x37, 0x2b, 0x17, 0x90, 0x37, 0x41, 0x15, 0x6d, 0x41, 0xc0,
	0x8b, 0x03, 0x69, 0x45, 0xa1, 0x12, 0x42, 0x2d, 0x02, 0x14, 0xa4, 0x44, 0x91, 0x5b, 0x21, 0x16,
	0xad, 0x34, 0x72, 0xe2, 0xd9, 0xc4, 0x52, 0xec, 0x89, 0x66, 0x26, 0x2c, 0xcb, 0x3b, 0xef, 0xdc,
	0xaf, 0xaf, 0xfc, 0x36, 0x24, 0xee, 0xf0, 0x0f, 0x38, 0x9e, 0x33, 0x71, 0x36, 0x66, 0x77, 0x6d,
	0xf1, 0xb2, 0x9b, 0xf1, 0x7c, 0x97, 0x33, 0x73, 0x2e, 0x36, 0xdc, 0x0f, 0x95, 0x16, 0x92, 0xf7,
	0x79, 0x32, 0x8f, 0x12, 0xde, 0x5f, 0x49, 0x31, 0xe3, 0xe1, 0x5a, 0x72, 0xd5, 0x57, 0x9a, 0xbd,
	0xcb, 0xf5, 0x84, 0x4b, 0x25, 0x92, 0xc9, 0xe8, 0x91, 0x0e, 0x74, 0xa4, 0x74, 0x34, 0x53, 0xec,
	0x41, 0xe8, 0x21, 0x4c, 0x0b, 0xe7, 0x25, 0xe2, 0x7a, 0xc4, 0xf5, 0x2e, 0x21, 0x74, 0xaf, 0x5b,
	0x9b, 0x8f, 0x82, 0xe5, 0x9a, 0x2b, 0xe2, 0x77, 0x0f, 0x76, 0xbd, 0xb9, 0x94, 0x42, 0xda, 0xad,
	0xde, 0xee, 0x56, 0xcc, 0x95, 0x0a, 0xe6, 0xdc, 0x6e, 0xde, 0xce, 0x6f, 0xea, 0x20, 0x4a, 0x8e,
	0x85, 0x8c, 0xd1, 0x51, 0x24, 0x04, 0xba, 0xf5, 0x73, 0x1d, 0x60, 0x12, 0xc8, 0x00, 0x77, 0x31,
	0x22, 0xe7, 0x75, 0x80, 0x95, 0x89, 0x8c, 0x45, 0xa1, 0x72, 0xaf, 0x3c, 0x7b, 0xe5, 0xc5, 0xd6,
	0xa0, 0xeb, 0xd9, 0x03, 0xd8, 0xa8, 0x94, 0x96, 0x51, 0x32, 0x7f, 0x3f, 0x5d, 0xf8, 0x4d, 0x42,
	0x0f, 0x43, 0xe5, 0xbc, 0x00, 0x9d, 0x2d, 0x95, 0x25, 0xeb, 0xe5, 0xd2, 0xfd, 0xa5, 0x8e, 0x02,
	0x0d, 0xff, 0x5a, 0x06, 0x1a, 0xe3, 0x53, 0xe7, 0x3e, 0xc0, 0xb1, 0x14, 0x31, 0x8b, 0x45, 0xa2,
	0x17, 0xee, 0x13, 0xc6, 0xa3, 0x97, 0xf3, 0x88, 0x12, 0xcd, 0xe7, 0x5c, 0x5a, 0x93, 0x14, 0x3e,
	0x4a, 0xd1, 0xa9, 0xc9, 0x96, 0x4b, 0x26, 0xbf, 0x5a, 0x93, 0x0c, 0x64, 0x4c, 0x5e, 0x03, 0xc3,
	0x62, 0xa7, 0x3c, 0x90, 0x6e, 0xa5, 0xd8, 0xa3, 0x91, 0xa2, 0x0f, 0x11, 0xec, 0x3c, 0x0f, 0x7b,
	0x19, 0x93, 0x1c, 0x7e, 0x23, 0x87, 0xf6, 0x06, 0x62, 0x0c, 0x5e, 0x85, 0x86, 0x16, 0xf6, 0x0c,
	0x57, 0x8b, 0xf5, 0xeb, 0x5a, 0xd0, 0x09, 0x6e, 0xc3, 0xb5, 0x0d, 0x8f, 0xd4, 0x7f, 0x27, 0xf5,
	0x96, 0x05, 0x18, 0xf1, 0xbb, 0x80, 0x78, 0x8a, 0xbd, 0x5a, 0xac, 0x5d, 0xd3, 0xc2, 0x44, 0x7e,
	0x13, 0xda, 0x96, 0x45, 0xca, 0x7f, 0x90, 0x32, 0xd0, 0xb6, 0x11, 0x3e, 0x82, 0x67, 0xa6, 0x81,
	0x8a, 0x66, 0x6c, 0xb6, 0xc0, 0xa4, 0xcf, 0x30, 0xe7, 0xa6, 0xf8, 0x10, 0x1f, 0x4f, 0x31, 0x49,
	0x6e, 0xad, 0x30, 0xe3, 0x5d, 0xc3, 0x7f, 0x6b, 0x87, 0x3e, 0x26, 0xb6, 0xf3, 0x1e, 0xdc, 0xbc,
	0x4c, 0x9d, 0xa2, 0xfa, 0x93, 0xa2, 0xba, 0x71, 0xb1, 0x8e, 0x89, 0xf4, 0x01, 0x74, 0x16, 0x4c,
	0x4b, 0xce, 0x59, 0x22, 0x42, 0x6e, 0xca, 0xb1, 0x5e, 0x18, 0x5c, 0x7b, 0xf1, 0x18, 0x19, 0x63,
	0x24, 0xa4, 0x15, 0xd9, 0x87, 0x27, 0x73, 0x12, 0x14, 0xc1, 0x5f, 0x14, 0xc1, 0xfe, 0x59, 0xb0,
	0xf1, 0x7c, 0x07, 0xf6, 0xd5, 0x3a, 0x8e, 0x03, 0x19, 0x7d, 0xc2, 0x29, 0x45, 0xca, 0x6d, 0x9c,
	0x7b, 0xff, 0x53, 0x21, 0x96, 0x3c, 0x48, 0xc8, 0xb5, 0x93, 0x91, 0x4c, 0x06, 0x95, 0x73, 0x07,
	0x9e, 0xca, 0xeb, 0x90, 0xf3, 0xdf, 0xe4, 0x7c, 0x3d, 0x47, 0x30, 0xe6, 0xf7, 0xa0, 0x39, 0xe7,
	0x9a, 0x69, 0xb1, 0x62, 0x1f, 0xbb, 0xcd, 0x12, 0x15, 0x85, 0xe8, 0xc7, 0x62, 0xf5, 0x81, 0xf3,
	0x1c, 0xec, 0x65, 0x44, 0xb2, 0xf9, 0xc7, 0x96, 0x94, 0x45, 0xa4, 0xf2, 0xb7, 0x7e, 0x6a, 0x40,
	0xc3, 0xe7, 0x6a, 0x25, 0x12, 0xc5, 0x9d, 0x97, 0xa1, 0x6a, 0xc6, 0x48, 0xbe, 0xc3, 0xed, 0x88,
	0xa2, 0x11, 0xf3, 0x76, 0xfa, 0xd7, 0x27, 0xa0, 0x73, 0x08, 0xfb, 0xe9, 0x00, 0x61, 0x67, 0x26,
	0x08, 0xb6, 0x6e, 0x05, 0xc9, 0x5e, 0x8e, 0x9c, 0x9f, 0x33, 0x23, 0x5c, 0x0f, 0xb7, 0x6b, 0xbf,
	0x13, 0xef, 0x3e, 0xc0, 0x56, 0xad, 0xdb, 0xc1, 0x85, 0x8d, 0x9a, 0x2a, 0xde, 0xf8, 0x8f, 0x22,
	0x8d, 0xb5, 0x11, 0xfd, 0xf7, 0x37, 0x70, 0x67, 0x08, 0x15, 0x29, 0x4e, 0xb0, 0xfd, 0x52, 0xd6,
	0x3d, 0xaf, 0xf4, 0x9c, 0xf5, 0x36, 0x17, 0xe1, 0xf9, 0xe2, 0xc4, 0x4f, 0x35, 0xba, 0x9f, 0xd6,
	0xa0, 0x82, 0x0b, 0xe7, 0x69, 0xa8, 0xe1, 0x12, 0x6b, 0xc5, 0xfd, 0x6c, 0x8c, 0x77, 0x53, 0xf5,
	0xab, 0xb8, 0x1c, 0x86, 0x69, 0xe3, 0x98, 0x24, 0xbc, 0xc2, 0x70, 0xe8, 0x63, 0xc5, 0x61, 0xcd,
	0x86, 0x6c, 0x7a, 0xca, 0x56, 0x81, 0xc6, 0xea, 0x4d, 0xdc, 0xcf, 0xc7, 0x85, 0xc5, 0x79, 0x40,
	0x02, 0x7e, 0xc6, 0x7f, 0x78, 0x3a, 0x21, 0xb6, 0x33, 0x80, 0x2a, 0x4d, 0x92, 0x2f, 0xc6, 0xc5,
	0x89, 0x27, 0x68, 0x16, 0xd1, 0xe0, 0x82, 0x88, 0xbe, 0x2c, 0x1b, 0xd1, 0xe0, 0xbc, 0x88, 0xde,
	0x00, 0x1c, 0x48, 0x3a, 0x58, 0x32, 0x03, 0x71, 0xbf, 0x3a, 0x3f, 0xae, 0x90, 0xcf, 0xa2, 0x38,
	0x58, 0x92, 0x1a, 0x18, 0x82, 0xf9, 0x8d, 0x05, 0x76, 0xd5, 0x4c, 0xaf, 0xaf, 0x4b, 0x9c, 0xc7,
	0x20, 0xf1, 0xcd, 0xd3, 0xcc, 0x5e, 0x1f, 0xee, 0x37, 0x25, 0x68, 0x8d, 0xcd, 0x5b, 0x05, 0x47,
	0xc5, 0xde, 0x6e, 0x9f, 0xbb, 0xdf, 0x96, 0xe0, 0xb7, 0xce, 0xb4, 0x3f, 0x76, 0x6c, 0x8d, 0xb2,
	0xe3, 0x7e, 0x57, 0x7c, 0x6d, 0x16, 0x9a, 0x91, 0x06, 0xee, 0xf7, 0x65, 0x49, 0x03, 0xe7, 0x4d,
	0x68, 0x87, 0x91, 0xe4, 0x33, 0x6d, 0x6f, 0xf6, 0x87, 0x12, 0x37, 0xdb, 0x22, 0x06, 0x5d, 0xed,
	0x11, 0xf4, 0x2e, 0x19, 0xb2, 0xee, 0x8f, 0x25, 0x8e, 0x7e, 0x70, 0xe1, 0xec, 0x7d, 0xf8, 0x08,
	0x7a, 0x91, 0xc8, 0x75, 0xd2, 0xf6, 0x6b, 0xe7, 0xc3, 0xbb, 0xff, 0xe7, 0x3b, 0x68, 0x5a, 0x33,
	0xdf, 0x1a, 0x77, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x29, 0xf4, 0x73, 0x93, 0x46, 0x09, 0x00,
	0x00,
}
