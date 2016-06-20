// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetTreeNodePMStatistics_Ad.proto
// DO NOT EDIT!

/*
Package st_GetTreeNodePMStatistics_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetTreeNodePMStatistics_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetTreeNodePMStatistics_Ad

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
	FromMonth                      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=from_month,json=fromMonth" json:"from_month,omitempty"`
	FromMonthNull                  bool                        `protobuf:"varint,1001,opt,name=from_month_null,json=fromMonthNull" json:"from_month_null,omitempty"`
	FromYear                       *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=from_year,json=fromYear" json:"from_year,omitempty"`
	FromYearNull                   bool                        `protobuf:"varint,1002,opt,name=from_year_null,json=fromYearNull" json:"from_year_null,omitempty"`
	ToMonth                        *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=to_month,json=toMonth" json:"to_month,omitempty"`
	ToMonthNull                    bool                        `protobuf:"varint,1003,opt,name=to_month_null,json=toMonthNull" json:"to_month_null,omitempty"`
	ToYear                         *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=to_year,json=toYear" json:"to_year,omitempty"`
	ToYearNull                     bool                        `protobuf:"varint,1004,opt,name=to_year_null,json=toYearNull" json:"to_year_null,omitempty"`
	BasicCharacteristicNumbers     *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=basic_characteristic_numbers,json=basicCharacteristicNumbers" json:"basic_characteristic_numbers,omitempty"`
	BasicCharacteristicNumbersNull bool                        `protobuf:"varint,1005,opt,name=basic_characteristic_numbers_null,json=basicCharacteristicNumbersNull" json:"basic_characteristic_numbers_null,omitempty"`
	HTreeNodeIds                   *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=h_tree_node_ids,json=hTreeNodeIds" json:"h_tree_node_ids,omitempty"`
	HTreeNodeIdsNull               bool                        `protobuf:"varint,1006,opt,name=h_tree_node_ids_null,json=hTreeNodeIdsNull" json:"h_tree_node_ids_null,omitempty"`
	SummarizeMonths                *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=summarize_months,json=summarizeMonths" json:"summarize_months,omitempty"`
	SummarizeMonthsNull            bool                        `protobuf:"varint,1007,opt,name=summarize_months_null,json=summarizeMonthsNull" json:"summarize_months_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Month                     *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=month" json:"month,omitempty"`
	TotalValue                *dstore_values.DecimalValue `protobuf:"bytes,10002,opt,name=total_value,json=totalValue" json:"total_value,omitempty"`
	Year                      *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=year" json:"year,omitempty"`
	HTreeNodeId               *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	DirectValue               *dstore_values.DecimalValue `protobuf:"bytes,10005,opt,name=direct_value,json=directValue" json:"direct_value,omitempty"`
	BasicCharacteristicNumber *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=basic_characteristic_number,json=basicCharacteristicNumber" json:"basic_characteristic_number,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetMonth() *dstore_values.IntegerValue {
	if m != nil {
		return m.Month
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

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetTreeNodePMStatistics_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetTreeNodePMStatistics_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetTreeNodePMStatistics_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xfb, 0x6a, 0x13, 0x4d,
	0x14, 0xa7, 0x5f, 0x9a, 0xcb, 0x77, 0x92, 0x36, 0x65, 0xaa, 0xb2, 0x4d, 0xa4, 0xd8, 0x16, 0x51,
	0x10, 0x36, 0xd2, 0x4a, 0x51, 0x51, 0xa4, 0x8a, 0x4a, 0x95, 0x84, 0xb2, 0x5e, 0xa0, 0x52, 0x58,
	0x36, 0xbb, 0xd3, 0x64, 0x21, 0xbb, 0x53, 0x66, 0x26, 0x16, 0x7d, 0x0a, 0xef, 0x2f, 0xa0, 0xaf,
	0xe2, 0xc3, 0x78, 0x7f, 0x05, 0xcf, 0xce, 0xd9, 0xa4, 0xcd, 0xda, 0x36, 0xd1, 0x7f, 0xda, 0xcc,
	0xce, 0xef, 0x72, 0x66, 0xcf, 0xf9, 0xed, 0xc0, 0x8d, 0x40, 0x69, 0x21, 0x79, 0x83, 0xc7, 0x9d,
	0x30, 0xe6, 0x8d, 0x3d, 0x29, 0x7c, 0x1e, 0xf4, 0x25, 0x57, 0x0d, 0xa5, 0xdd, 0xfb, 0x5c, 0x3f,
	0x96, 0x9c, 0xb7, 0x44, 0xc0, 0xb7, 0x9a, 0x8f, 0xb4, 0xa7, 0x43, 0xa5, 0x43, 0x5f, 0xb9, 0x1b,
	0x81, 0x8d, 0x40, 0x2d, 0xd8, 0x25, 0x62, 0xdb, 0xc4, 0xb6, 0x4f, 0xa4, 0xd4, 0xe6, 0x53, 0xab,
	0xe7, 0x5e, 0xaf, 0xcf, 0x15, 0x29, 0xd4, 0x16, 0x46, 0xfd, 0xb9, 0x94, 0x42, 0xa6, 0x5b, 0xf5,
	0xd1, 0xad, 0x88, 0x2b, 0xe5, 0x75, 0x78, 0xba, 0xb9, 0x92, 0xdd, 0xd4, 0x5e, 0x18, 0xef, 0x0a,
	0x19, 0xa1, 0xa3, 0x88, 0x09, 0xb4, 0xfc, 0xb1, 0x00, 0xb0, 0xe5, 0x49, 0x0f, 0x77, 0xb9, 0x54,
	0xec, 0x3a, 0xc0, 0xae, 0x14, 0x91, 0x1b, 0x89, 0x58, 0x77, 0xad, 0xa9, 0x73, 0x53, 0x17, 0xcb,
	0xab, 0x75, 0x3b, 0x3d, 0x42, 0x5a, 0x55, 0x18, 0x6b, 0xde, 0xe1, 0xf2, 0x69, 0xb2, 0x72, 0xfe,
	0x4f, 0xe0, 0xcd, 0x04, 0xcd, 0x2e, 0x40, 0xf5, 0x80, 0xeb, 0xc6, 0xfd, 0x5e, 0xcf, 0xfa, 0x52,
	0x44, 0x85, 0x92, 0x33, 0x33, 0x04, 0xb5, 0xf0, 0x29, 0xbb, 0x0a, 0x86, 0xe5, 0xbe, 0xe0, 0x9e,
	0xb4, 0xfe, 0x1b, 0xef, 0x51, 0x4a, 0xd0, 0xdb, 0x08, 0x66, 0xe7, 0x61, 0x76, 0xc8, 0x24, 0x87,
	0xaf, 0xe4, 0x50, 0x19, 0x40, 0x8c, 0xc1, 0x3a, 0x94, 0xb4, 0x48, 0xcf, 0x90, 0x1b, 0xaf, 0x5f,
	0xd4, 0x82, 0x4e, 0xb0, 0x02, 0x33, 0x03, 0x1e, 0xa9, 0x7f, 0x23, 0xf5, 0x72, 0x0a, 0x30, 0xe2,
	0x57, 0x00, 0xf1, 0x54, 0xfb, 0xf4, 0x78, 0xed, 0x82, 0x16, 0xa6, 0xf2, 0x25, 0xa8, 0xa4, 0x2c,
	0x52, 0xfe, 0x4e, 0xca, 0x40, 0xdb, 0x46, 0x78, 0x07, 0xce, 0xb6, 0x3d, 0x15, 0xfa, 0xae, 0xdf,
	0xc5, 0x86, 0xf8, 0xd8, 0x0f, 0x33, 0x18, 0x88, 0x8f, 0xda, 0xd8, 0x1b, 0x2b, 0x6f, 0xdc, 0x6a,
	0x19, 0x37, 0xa5, 0x65, 0x18, 0x77, 0xc8, 0xac, 0x66, 0xf8, 0x77, 0x46, 0xe8, 0x2d, 0x62, 0xb3,
	0x07, 0xb0, 0x74, 0x92, 0x3a, 0x55, 0xf5, 0x83, 0xaa, 0x5a, 0x3c, 0x5e, 0xc7, 0x54, 0xba, 0x01,
	0xd5, 0xae, 0xab, 0x71, 0x86, 0xdd, 0x18, 0x87, 0xd8, 0x0d, 0x03, 0x65, 0x15, 0xc6, 0x16, 0x57,
	0xe9, 0x0e, 0xa6, 0x7e, 0x33, 0x50, 0xac, 0x01, 0xa7, 0x32, 0x12, 0x54, 0xc1, 0x4f, 0xaa, 0x60,
	0xee, 0x30, 0xd8, 0x78, 0xde, 0x83, 0x39, 0xd5, 0x8f, 0x22, 0x4f, 0x86, 0x2f, 0x39, 0xb5, 0x48,
	0x59, 0xc5, 0x23, 0xdf, 0x7f, 0x5b, 0x88, 0x1e, 0xf7, 0x62, 0x72, 0xad, 0x0e, 0x49, 0xa6, 0x83,
	0x8a, 0xad, 0xc1, 0xe9, 0xac, 0x0e, 0x39, 0xff, 0x22, 0xe7, 0xf9, 0x0c, 0x21, 0x31, 0x5f, 0xfe,
	0x9c, 0x87, 0x92, 0xc3, 0xd5, 0x9e, 0x88, 0x15, 0x67, 0x97, 0x21, 0x6f, 0x32, 0x98, 0xc6, 0x63,
	0x78, 0xe6, 0x34, 0xe1, 0x94, 0xcf, 0xbb, 0xc9, 0x5f, 0x87, 0x80, 0x6c, 0x1b, 0xe6, 0x92, 0xf4,
	0xb9, 0x87, 0xe2, 0x87, 0x73, 0x9f, 0x43, 0xb2, 0x9d, 0x21, 0x67, 0x43, 0xda, 0xc4, 0xf5, 0xe6,
	0xc1, 0xda, 0xa9, 0x46, 0xa3, 0x0f, 0x30, 0x4b, 0xc5, 0x34, 0xf5, 0x38, 0xe9, 0x89, 0xe2, 0xe2,
	0x1f, 0x8a, 0xf4, 0x4d, 0x68, 0xd2, 0x7f, 0x67, 0x00, 0x67, 0x0f, 0x21, 0x27, 0xc5, 0x3e, 0xce,
	0x70, 0xc2, 0xba, 0x66, 0xff, 0xc5, 0x67, 0xca, 0x1e, 0xbc, 0x0a, 0xdb, 0x11, 0xfb, 0x4e, 0xa2,
	0x52, 0xfb, 0x94, 0x83, 0x1c, 0x2e, 0xd8, 0x19, 0x28, 0xe0, 0x12, 0xdb, 0x69, 0xbd, 0x6a, 0xe1,
	0xdb, 0xc9, 0x3b, 0x79, 0x5c, 0x6e, 0x06, 0x6c, 0x15, 0xf2, 0x14, 0xc7, 0xd7, 0xad, 0xf1, 0x99,
	0x21, 0x28, 0xbb, 0x09, 0x98, 0x3b, 0xed, 0xf5, 0x5c, 0x83, 0xb1, 0xde, 0x1c, 0xcd, 0x0c, 0xb8,
	0x1f, 0x46, 0x5e, 0x8f, 0x98, 0x60, 0x08, 0xe6, 0x37, 0xb6, 0x69, 0xda, 0x84, 0xf4, 0xed, 0x04,
	0x8e, 0x06, 0x89, 0x63, 0x3d, 0x3b, 0x3a, 0x93, 0xd6, 0xbb, 0x09, 0xb8, 0xe5, 0x43, 0xa3, 0xca,
	0x6e, 0x41, 0x25, 0x08, 0x25, 0xf7, 0x75, 0x5a, 0xf4, 0xfb, 0x09, 0x8a, 0x2e, 0x13, 0x83, 0xaa,
	0xde, 0x81, 0xfa, 0x09, 0x31, 0xb5, 0x3e, 0x4c, 0x50, 0xd0, 0xc2, 0xb1, 0xe9, 0xbd, 0xfd, 0x04,
	0xea, 0xa1, 0xc8, 0xb4, 0xfa, 0xe0, 0x3e, 0x7b, 0xb6, 0xfe, 0x6f, 0x37, 0x5d, 0xbb, 0x60, 0xee,
	0x92, 0xb5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0x16, 0x62, 0x5f, 0x2a, 0x07, 0x00, 0x00,
}
