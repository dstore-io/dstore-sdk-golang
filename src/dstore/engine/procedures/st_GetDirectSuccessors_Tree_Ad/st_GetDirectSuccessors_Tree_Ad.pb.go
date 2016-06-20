// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/st_GetDirectSuccessors_Tree_Ad.proto
// DO NOT EDIT!

/*
Package st_GetDirectSuccessors_Tree_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/st_GetDirectSuccessors_Tree_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package st_GetDirectSuccessors_Tree_Ad

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
	HTreeNodeId                      *dstore_values.IntegerValue   `protobuf:"bytes,1,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	HTreeNodeIdNull                  bool                          `protobuf:"varint,1001,opt,name=h_tree_node_id_null,json=hTreeNodeIdNull" json:"h_tree_node_id_null,omitempty"`
	FromDate                         *dstore_values.TimestampValue `protobuf:"bytes,2,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull                     bool                          `protobuf:"varint,1002,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                           *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull                       bool                          `protobuf:"varint,1003,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	BasicCharacteristicNo1           *dstore_values.IntegerValue   `protobuf:"bytes,4,opt,name=basic_characteristic_no1,json=basicCharacteristicNo1" json:"basic_characteristic_no1,omitempty"`
	BasicCharacteristicNo1Null       bool                          `protobuf:"varint,1004,opt,name=basic_characteristic_no1_null,json=basicCharacteristicNo1Null" json:"basic_characteristic_no1_null,omitempty"`
	WeightBasicCharacteristicNo1     *dstore_values.DecimalValue   `protobuf:"bytes,5,opt,name=weight_basic_characteristic_no1,json=weightBasicCharacteristicNo1" json:"weight_basic_characteristic_no1,omitempty"`
	WeightBasicCharacteristicNo1Null bool                          `protobuf:"varint,1005,opt,name=weight_basic_characteristic_no1_null,json=weightBasicCharacteristicNo1Null" json:"weight_basic_characteristic_no1_null,omitempty"`
	BasicCharacteristicNo2           *dstore_values.IntegerValue   `protobuf:"bytes,6,opt,name=basic_characteristic_no2,json=basicCharacteristicNo2" json:"basic_characteristic_no2,omitempty"`
	BasicCharacteristicNo2Null       bool                          `protobuf:"varint,1006,opt,name=basic_characteristic_no2_null,json=basicCharacteristicNo2Null" json:"basic_characteristic_no2_null,omitempty"`
	WeightBasicCharacteristicNo2     *dstore_values.DecimalValue   `protobuf:"bytes,7,opt,name=weight_basic_characteristic_no2,json=weightBasicCharacteristicNo2" json:"weight_basic_characteristic_no2,omitempty"`
	WeightBasicCharacteristicNo2Null bool                          `protobuf:"varint,1007,opt,name=weight_basic_characteristic_no2_null,json=weightBasicCharacteristicNo2Null" json:"weight_basic_characteristic_no2_null,omitempty"`
	BasicCharacteristicNo3           *dstore_values.IntegerValue   `protobuf:"bytes,8,opt,name=basic_characteristic_no3,json=basicCharacteristicNo3" json:"basic_characteristic_no3,omitempty"`
	BasicCharacteristicNo3Null       bool                          `protobuf:"varint,1008,opt,name=basic_characteristic_no3_null,json=basicCharacteristicNo3Null" json:"basic_characteristic_no3_null,omitempty"`
	WeightBasicCharacteristicNo3     *dstore_values.DecimalValue   `protobuf:"bytes,9,opt,name=weight_basic_characteristic_no3,json=weightBasicCharacteristicNo3" json:"weight_basic_characteristic_no3,omitempty"`
	WeightBasicCharacteristicNo3Null bool                          `protobuf:"varint,1009,opt,name=weight_basic_characteristic_no3_null,json=weightBasicCharacteristicNo3Null" json:"weight_basic_characteristic_no3_null,omitempty"`
	SourceTableForStatistics         *dstore_values.IntegerValue   `protobuf:"bytes,10,opt,name=source_table_for_statistics,json=sourceTableForStatistics" json:"source_table_for_statistics,omitempty"`
	SourceTableForStatisticsNull     bool                          `protobuf:"varint,1010,opt,name=source_table_for_statistics_null,json=sourceTableForStatisticsNull" json:"source_table_for_statistics_null,omitempty"`
	DisplayOnlyActive                *dstore_values.BooleanValue   `protobuf:"bytes,11,opt,name=display_only_active,json=displayOnlyActive" json:"display_only_active,omitempty"`
	DisplayOnlyActiveNull            bool                          `protobuf:"varint,1011,opt,name=display_only_active_null,json=displayOnlyActiveNull" json:"display_only_active_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetBasicCharacteristicNo1() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNo1
	}
	return nil
}

func (m *Parameters) GetWeightBasicCharacteristicNo1() *dstore_values.DecimalValue {
	if m != nil {
		return m.WeightBasicCharacteristicNo1
	}
	return nil
}

func (m *Parameters) GetBasicCharacteristicNo2() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNo2
	}
	return nil
}

func (m *Parameters) GetWeightBasicCharacteristicNo2() *dstore_values.DecimalValue {
	if m != nil {
		return m.WeightBasicCharacteristicNo2
	}
	return nil
}

func (m *Parameters) GetBasicCharacteristicNo3() *dstore_values.IntegerValue {
	if m != nil {
		return m.BasicCharacteristicNo3
	}
	return nil
}

func (m *Parameters) GetWeightBasicCharacteristicNo3() *dstore_values.DecimalValue {
	if m != nil {
		return m.WeightBasicCharacteristicNo3
	}
	return nil
}

func (m *Parameters) GetSourceTableForStatistics() *dstore_values.IntegerValue {
	if m != nil {
		return m.SourceTableForStatistics
	}
	return nil
}

func (m *Parameters) GetDisplayOnlyActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.DisplayOnlyActive
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
	TreeLevel                      *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=tree_level,json=treeLevel" json:"tree_level,omitempty"`
	Predecessor                    *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=predecessor" json:"predecessor,omitempty"`
	LevelNo                        *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=level_no,json=levelNo" json:"level_no,omitempty"`
	RelativeValueBasicCharacNo3    *dstore_values.DecimalValue `protobuf:"bytes,10004,opt,name=relative_value_basic_charac_no3,json=relativeValueBasicCharacNo3" json:"relative_value_basic_charac_no3,omitempty"`
	RelativeValueBasicCharacNo2    *dstore_values.DecimalValue `protobuf:"bytes,10005,opt,name=relative_value_basic_charac_no2,json=relativeValueBasicCharacNo2" json:"relative_value_basic_charac_no2,omitempty"`
	RelativeValueBasicCharacNo1    *dstore_values.DecimalValue `protobuf:"bytes,10006,opt,name=relative_value_basic_charac_no1,json=relativeValueBasicCharacNo1" json:"relative_value_basic_charac_no1,omitempty"`
	TotalValueIndex                *dstore_values.DecimalValue `protobuf:"bytes,10007,opt,name=total_value_index,json=totalValueIndex" json:"total_value_index,omitempty"`
	RelativeValueIndex             *dstore_values.DecimalValue `protobuf:"bytes,10008,opt,name=relative_value_index,json=relativeValueIndex" json:"relative_value_index,omitempty"`
	NodeDescription                *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Active                         *dstore_values.BooleanValue `protobuf:"bytes,10010,opt,name=active" json:"active,omitempty"`
	HasNextSibling                 *dstore_values.BooleanValue `protobuf:"bytes,10011,opt,name=has_next_sibling,json=hasNextSibling" json:"has_next_sibling,omitempty"`
	TotalValueBasicCharacNo1       *dstore_values.DecimalValue `protobuf:"bytes,10012,opt,name=total_value_basic_charac_no1,json=totalValueBasicCharacNo1" json:"total_value_basic_charac_no1,omitempty"`
	HTreeNodeId                    *dstore_values.IntegerValue `protobuf:"bytes,10013,opt,name=h_tree_node_id,json=hTreeNodeId" json:"h_tree_node_id,omitempty"`
	TotalValueBasicCharacNo2       *dstore_values.DecimalValue `protobuf:"bytes,10014,opt,name=total_value_basic_charac_no2,json=totalValueBasicCharacNo2" json:"total_value_basic_charac_no2,omitempty"`
	TotalValueBasicCharacNo3       *dstore_values.DecimalValue `protobuf:"bytes,10015,opt,name=total_value_basic_charac_no3,json=totalValueBasicCharacNo3" json:"total_value_basic_charac_no3,omitempty"`
	MaxRelativeValuePerPredecessor *dstore_values.DecimalValue `protobuf:"bytes,10016,opt,name=max_relative_value_per_predecessor,json=maxRelativeValuePerPredecessor" json:"max_relative_value_per_predecessor,omitempty"`
	TreeNodeId                     *dstore_values.IntegerValue `protobuf:"bytes,10017,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	NodeId                         *dstore_values.IntegerValue `protobuf:"bytes,10018,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	MaxRelativeValueIndexPerPred   *dstore_values.DecimalValue `protobuf:"bytes,10019,opt,name=max_relative_value_index_per_pred,json=maxRelativeValueIndexPerPred" json:"max_relative_value_index_per_pred,omitempty"`
	FromDate                       *dstore_values.StringValue  `protobuf:"bytes,10020,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	ToDate                         *dstore_values.StringValue  `protobuf:"bytes,10021,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	HasSuccessors                  *dstore_values.BooleanValue `protobuf:"bytes,10022,opt,name=has_successors,json=hasSuccessors" json:"has_successors,omitempty"`
	Deleted                        *dstore_values.BooleanValue `protobuf:"bytes,10023,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTreeLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeLevel
	}
	return nil
}

func (m *Response_Row) GetPredecessor() *dstore_values.IntegerValue {
	if m != nil {
		return m.Predecessor
	}
	return nil
}

func (m *Response_Row) GetLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelNo
	}
	return nil
}

func (m *Response_Row) GetRelativeValueBasicCharacNo3() *dstore_values.DecimalValue {
	if m != nil {
		return m.RelativeValueBasicCharacNo3
	}
	return nil
}

func (m *Response_Row) GetRelativeValueBasicCharacNo2() *dstore_values.DecimalValue {
	if m != nil {
		return m.RelativeValueBasicCharacNo2
	}
	return nil
}

func (m *Response_Row) GetRelativeValueBasicCharacNo1() *dstore_values.DecimalValue {
	if m != nil {
		return m.RelativeValueBasicCharacNo1
	}
	return nil
}

func (m *Response_Row) GetTotalValueIndex() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalValueIndex
	}
	return nil
}

func (m *Response_Row) GetRelativeValueIndex() *dstore_values.DecimalValue {
	if m != nil {
		return m.RelativeValueIndex
	}
	return nil
}

func (m *Response_Row) GetNodeDescription() *dstore_values.StringValue {
	if m != nil {
		return m.NodeDescription
	}
	return nil
}

func (m *Response_Row) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Response_Row) GetHasNextSibling() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasNextSibling
	}
	return nil
}

func (m *Response_Row) GetTotalValueBasicCharacNo1() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalValueBasicCharacNo1
	}
	return nil
}

func (m *Response_Row) GetHTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.HTreeNodeId
	}
	return nil
}

func (m *Response_Row) GetTotalValueBasicCharacNo2() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalValueBasicCharacNo2
	}
	return nil
}

func (m *Response_Row) GetTotalValueBasicCharacNo3() *dstore_values.DecimalValue {
	if m != nil {
		return m.TotalValueBasicCharacNo3
	}
	return nil
}

func (m *Response_Row) GetMaxRelativeValuePerPredecessor() *dstore_values.DecimalValue {
	if m != nil {
		return m.MaxRelativeValuePerPredecessor
	}
	return nil
}

func (m *Response_Row) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Response_Row) GetNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Response_Row) GetMaxRelativeValueIndexPerPred() *dstore_values.DecimalValue {
	if m != nil {
		return m.MaxRelativeValueIndexPerPred
	}
	return nil
}

func (m *Response_Row) GetFromDate() *dstore_values.StringValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Response_Row) GetToDate() *dstore_values.StringValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Response_Row) GetHasSuccessors() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasSuccessors
	}
	return nil
}

func (m *Response_Row) GetDeleted() *dstore_values.BooleanValue {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.st_GetDirectSuccessors_Tree_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.st_GetDirectSuccessors_Tree_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.st_GetDirectSuccessors_Tree_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 1114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x57, 0xeb, 0x6e, 0x1b, 0x45,
	0x14, 0x56, 0x09, 0xb1, 0xd3, 0x93, 0x34, 0x49, 0x27, 0x50, 0x0d, 0x71, 0x5a, 0xd2, 0x00, 0x12,
	0x3f, 0xc0, 0x21, 0xeb, 0x42, 0x51, 0x11, 0x15, 0x49, 0x43, 0x4b, 0x44, 0xbb, 0x8d, 0x36, 0xa5,
	0x12, 0x15, 0xd2, 0x68, 0xbd, 0x3b, 0xb5, 0x57, 0x5a, 0xef, 0x58, 0x33, 0xe3, 0x24, 0x7d, 0x0b,
	0xa0, 0xdc, 0xef, 0x77, 0x1e, 0x89, 0x67, 0xe0, 0x4e, 0xe1, 0x05, 0x98, 0xcb, 0x3a, 0x5e, 0x3b,
	0x8e, 0x77, 0x1b, 0xff, 0x49, 0x3c, 0x9e, 0xf3, 0x7d, 0xe7, 0x9b, 0xcf, 0x7b, 0xe6, 0xec, 0x81,
	0x57, 0x43, 0x21, 0x19, 0xa7, 0xab, 0x34, 0x69, 0x44, 0x09, 0x5d, 0x6d, 0x73, 0x16, 0xd0, 0xb0,
	0xc3, 0xa9, 0x58, 0x15, 0x92, 0x5c, 0xa3, 0x72, 0x33, 0xe2, 0x34, 0x90, 0x3b, 0x9d, 0x20, 0xa0,
	0x42, 0x30, 0x2e, 0xc8, 0x2d, 0x4e, 0x29, 0x59, 0x0f, 0xab, 0x2a, 0x52, 0x32, 0xf4, 0x9c, 0x85,
	0x57, 0x2d, 0xbc, 0x3a, 0x1a, 0xb3, 0xb8, 0x90, 0x26, 0xdb, 0xf5, 0xe3, 0x0e, 0x15, 0x96, 0x62,
	0xf1, 0x89, 0x7e, 0x05, 0x94, 0x73, 0xc6, 0xd3, 0xad, 0x4a, 0xff, 0x56, 0x4b, 0xd1, 0xf9, 0x0d,
	0x9a, 0x6e, 0x3e, 0x35, 0xb8, 0x29, 0xfd, 0x28, 0xb9, 0xcb, 0x78, 0xcb, 0x97, 0x11, 0x4b, 0x6c,
	0xd0, 0xca, 0x2f, 0x33, 0x00, 0xdb, 0x3e, 0xf7, 0xd5, 0x2e, 0xe5, 0x02, 0xbd, 0x06, 0xb3, 0x4d,
	0x22, 0xb5, 0x9a, 0x84, 0x85, 0x94, 0x44, 0x21, 0x3e, 0xb1, 0x7c, 0xe2, 0xd9, 0x69, 0xa7, 0x52,
	0x4d, 0xcf, 0x91, 0x2a, 0x8b, 0x12, 0x49, 0x1b, 0x94, 0xdf, 0xd6, 0x2b, 0x6f, 0xba, 0xa9, 0xf5,
	0xbb, 0x0a, 0xb0, 0x15, 0xa2, 0xe7, 0x61, 0xa1, 0x9f, 0x81, 0x24, 0x9d, 0x38, 0xc6, 0xbf, 0x96,
	0x15, 0xcf, 0x94, 0x37, 0x97, 0x09, 0x75, 0xd5, 0xf7, 0xe8, 0x12, 0x9c, 0xbc, 0xcb, 0x59, 0x8b,
	0x84, 0xbe, 0xa4, 0xf8, 0x11, 0x93, 0xeb, 0xec, 0x40, 0x2e, 0x19, 0xa9, 0x73, 0x49, 0xbf, 0xd5,
	0xb6, 0xd9, 0xa6, 0x74, 0xfc, 0xa6, 0x0a, 0x47, 0xcf, 0xc0, 0xec, 0x01, 0xd6, 0x66, 0xf9, 0xcd,
	0x66, 0x99, 0xe9, 0x86, 0x98, 0x14, 0x2f, 0x41, 0x59, 0x32, 0x9b, 0x60, 0xa2, 0x48, 0x82, 0x92,
	0x64, 0x86, 0xfe, 0x3c, 0xcc, 0xa4, 0x38, 0x4b, 0xfe, 0xbb, 0x25, 0x07, 0xbb, 0x6d, 0xa8, 0xdf,
	0x02, 0x5c, 0xf7, 0x45, 0x14, 0x90, 0xa0, 0xa9, 0x3c, 0x0c, 0x94, 0x85, 0x91, 0x90, 0x6a, 0x99,
	0xb0, 0x35, 0xfc, 0x68, 0xbe, 0x71, 0x67, 0x0c, 0xf8, 0x4a, 0x1f, 0xd6, 0x65, 0x6b, 0x68, 0x03,
	0xce, 0x1e, 0x45, 0x6b, 0xa5, 0xfc, 0x61, 0xa5, 0x2c, 0x0e, 0xc7, 0x1b, 0x69, 0x75, 0x78, 0x72,
	0x8f, 0x46, 0x8d, 0xa6, 0x24, 0x47, 0x2a, 0x9c, 0x1c, 0xaa, 0x30, 0xa4, 0x41, 0xd4, 0xf2, 0x63,
	0xab, 0x70, 0xc9, 0x72, 0x6c, 0x0c, 0xd7, 0x79, 0x13, 0x9e, 0xce, 0xc9, 0x61, 0xe5, 0xfe, 0x69,
	0xe5, 0x2e, 0x8f, 0x22, 0xcb, 0xf1, 0xd3, 0xc1, 0xa5, 0xe3, 0xfa, 0xe9, 0x8c, 0xf0, 0xd3, 0xb1,
	0x02, 0xff, 0x1a, 0xe5, 0xa7, 0x53, 0xcc, 0x4f, 0x07, 0x97, 0xc7, 0xf3, 0xd3, 0xc9, 0xf7, 0x33,
	0x95, 0xfb, 0x77, 0xbe, 0x9f, 0x4e, 0x8e, 0x9f, 0x35, 0x3c, 0x75, 0x5c, 0x3f, 0x6b, 0x23, 0xfc,
	0xac, 0x59, 0x81, 0xff, 0x8c, 0xf2, 0xb3, 0x56, 0xcc, 0xcf, 0x1a, 0x3e, 0x39, 0x9e, 0x9f, 0xb5,
	0x7c, 0x3f, 0x53, 0xb9, 0x0f, 0xf2, 0xfd, 0xb4, 0xa2, 0xef, 0x40, 0x45, 0xb0, 0x0e, 0x0f, 0x28,
	0x91, 0x7e, 0x3d, 0xa6, 0x44, 0xdd, 0xa6, 0x44, 0xdd, 0x1c, 0xd2, 0x04, 0x09, 0x0c, 0xf9, 0x96,
	0x62, 0x8b, 0xbf, 0xa5, 0xe1, 0x57, 0x19, 0xdf, 0x39, 0x00, 0xa3, 0x6b, 0xb0, 0x3c, 0x82, 0xdb,
	0x0a, 0xfd, 0xd7, 0x0a, 0x5d, 0x3a, 0x8a, 0xc4, 0x88, 0x7c, 0x13, 0x16, 0xc2, 0x48, 0xb4, 0x63,
	0xff, 0x1e, 0x61, 0x49, 0x7c, 0x8f, 0xa8, 0x63, 0x44, 0xbb, 0x14, 0x4f, 0x0f, 0x15, 0x57, 0x67,
	0x2c, 0xa6, 0x7e, 0x62, 0xc5, 0x9d, 0x4e, 0x71, 0x37, 0x15, 0x6c, 0xdd, 0xa0, 0xd0, 0xcb, 0x80,
	0x87, 0x90, 0x59, 0x35, 0xff, 0x59, 0x35, 0x8f, 0x1f, 0x42, 0x69, 0x19, 0x2b, 0xf7, 0xe7, 0x61,
	0xca, 0xa3, 0xa2, 0xcd, 0x12, 0x41, 0xd1, 0x0b, 0x30, 0x69, 0xfa, 0x56, 0xda, 0x4e, 0x16, 0xab,
	0xfd, 0x6d, 0xd1, 0xf6, 0xb4, 0xd7, 0xf5, 0x5f, 0xcf, 0x06, 0xa2, 0xb7, 0x61, 0x5e, 0x77, 0x2c,
	0x92, 0x69, 0x59, 0xaa, 0x3f, 0x4c, 0x28, 0x70, 0x75, 0x00, 0x3c, 0xd8, 0xd8, 0x6e, 0xa8, 0xf5,
	0x56, 0x6f, 0xed, 0xcd, 0xb5, 0xfa, 0xbf, 0x50, 0x67, 0x2a, 0xa7, 0x9d, 0x52, 0x35, 0x04, 0xcd,
	0x78, 0xee, 0x10, 0xa3, 0xed, 0xa3, 0x37, 0xec, 0x7f, 0xaf, 0x1b, 0x8e, 0xae, 0xc3, 0x04, 0x67,
	0x7b, 0xea, 0x6a, 0xd7, 0xa8, 0x4b, 0xd5, 0x87, 0xe9, 0xed, 0xd5, 0xae, 0x17, 0x55, 0x8f, 0xed,
	0x79, 0x9a, 0x66, 0xf1, 0xc1, 0x29, 0x98, 0x50, 0x0b, 0x74, 0x06, 0x4a, 0x6a, 0xa9, 0x9b, 0xed,
	0xbb, 0xae, 0xb2, 0x67, 0xd2, 0x9b, 0x54, 0x4b, 0xd5, 0x4a, 0x5f, 0x01, 0x30, 0x8d, 0x34, 0xa6,
	0xbb, 0x34, 0xc6, 0xef, 0xb9, 0xf9, 0x4f, 0xd7, 0x49, 0x1d, 0x7f, 0x5d, 0x87, 0xa3, 0xcb, 0x30,
	0xdd, 0xe6, 0x54, 0x15, 0x8b, 0x91, 0x81, 0xdf, 0x2f, 0x80, 0xce, 0x02, 0xd0, 0x45, 0x98, 0x32,
	0x79, 0x55, 0x95, 0xe0, 0xfb, 0x05, 0xc0, 0x65, 0x13, 0xed, 0x32, 0x5d, 0xd8, 0x9c, 0xc6, 0xbe,
	0x79, 0x4c, 0x4c, 0x60, 0x5f, 0xf1, 0x99, 0xc2, 0xfe, 0xc0, 0xcd, 0xaf, 0xec, 0x4a, 0x97, 0xc4,
	0x2c, 0x33, 0x35, 0xa9, 0x0b, 0x3b, 0x37, 0x87, 0x83, 0x3f, 0x1c, 0x2f, 0x87, 0x93, 0x9f, 0x63,
	0x0d, 0x7f, 0x34, 0x5e, 0x8e, 0x35, 0xf4, 0x06, 0x9c, 0x96, 0x4c, 0xfa, 0x71, 0x9a, 0x20, 0x4a,
	0x42, 0xba, 0x8f, 0x3f, 0x2e, 0xc0, 0x3a, 0x67, 0x60, 0xe6, 0xf3, 0x96, 0x06, 0x21, 0x17, 0x1e,
	0x1b, 0x50, 0x6b, 0xc9, 0x3e, 0x29, 0x40, 0x86, 0xfa, 0x24, 0x5a, 0xbe, 0xab, 0x30, 0x6f, 0xde,
	0xdf, 0x42, 0x2a, 0x02, 0x1e, 0xb5, 0x4d, 0xf9, 0x7d, 0xea, 0xf6, 0x17, 0x6f, 0xca, 0x25, 0x24,
	0x8f, 0x92, 0x46, 0xaa, 0x4b, 0x83, 0x36, 0x7b, 0x18, 0x74, 0x01, 0x4a, 0xe9, 0xfd, 0xf3, 0x99,
	0x9b, 0x7f, 0x01, 0xa5, 0xb1, 0x3a, 0x7b, 0xd3, 0x57, 0x77, 0x1e, 0xdd, 0x97, 0x44, 0x44, 0xf5,
	0x58, 0xf1, 0xe3, 0xcf, 0x0b, 0xe0, 0x67, 0x15, 0xca, 0x55, 0xa0, 0x1d, 0x8b, 0x41, 0xef, 0xc0,
	0x52, 0xd6, 0xdf, 0x43, 0x3f, 0xe0, 0x17, 0x05, 0xdc, 0xc1, 0x3d, 0xab, 0x07, 0x7e, 0xbd, 0xf5,
	0x43, 0x2f, 0xcb, 0x5f, 0xba, 0x0f, 0xf9, 0xb6, 0x3c, 0x5a, 0xa0, 0x83, 0xbf, 0x1a, 0x43, 0xa0,
	0x93, 0xc3, 0x5e, 0xc3, 0x5f, 0x8f, 0xc1, 0x5e, 0x43, 0x4d, 0x58, 0x69, 0xf9, 0xfb, 0x64, 0xe0,
	0xb1, 0x6b, 0x53, 0x4e, 0xb2, 0x17, 0xcf, 0x37, 0x05, 0x72, 0x9c, 0x53, 0x3c, 0x5e, 0xf6, 0x19,
	0xdc, 0xa6, 0x7c, 0x3b, 0x73, 0x17, 0x5d, 0x56, 0x6f, 0xe2, 0x59, 0x9b, 0xbf, 0x2d, 0x60, 0xb3,
	0xb9, 0x3a, 0x53, 0x97, 0x5f, 0x84, 0x72, 0x17, 0xfa, 0x5d, 0x01, 0x68, 0x29, 0xb1, 0x30, 0x0a,
	0xe7, 0x87, 0x1c, 0xd0, 0xd4, 0xd5, 0xc1, 0x31, 0xf1, 0xf7, 0x05, 0xce, 0xb7, 0x34, 0x78, 0x3e,
	0x53, 0x63, 0xe9, 0x21, 0x55, 0x3b, 0xca, 0x8c, 0x40, 0x3f, 0xe4, 0xd7, 0x58, 0x6f, 0x00, 0xba,
	0xd0, 0x9b, 0x6c, 0x7e, 0xcc, 0xc7, 0x75, 0xe7, 0x9a, 0x2b, 0xa0, 0xcb, 0x84, 0x88, 0x83, 0x16,
	0x85, 0x7f, 0x2a, 0x50, 0x5a, 0xa7, 0x14, 0xa6, 0xd7, 0xd5, 0xf4, 0x50, 0x15, 0xd2, 0x58, 0x0d,
	0x8d, 0x21, 0xfe, 0xb9, 0x00, 0xba, 0x1b, 0xbc, 0x71, 0x1b, 0x2a, 0x11, 0x1b, 0x68, 0x9c, 0xbd,
	0x99, 0xfa, 0xce, 0xc5, 0x63, 0x4e, 0xdb, 0xf5, 0x92, 0x19, 0x67, 0x6b, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0x77, 0xcc, 0x5b, 0xaf, 0x0f, 0x00, 0x00,
}
