// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_GetDirectSuccessors_Ad.proto
// DO NOT EDIT!

/*
Package im_GetDirectSuccessors_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_GetDirectSuccessors_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_GetDirectSuccessors_Ad

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
	TreeNodeId                      *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	TreeNodeIdNull                  bool                        `protobuf:"varint,1001,opt,name=tree_node_id_null,json=treeNodeIdNull" json:"tree_node_id_null,omitempty"`
	LanguageId                      *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	LanguageIdNull                  bool                        `protobuf:"varint,1002,opt,name=language_id_null,json=languageIdNull" json:"language_id_null,omitempty"`
	OnlyActive                      *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=only_active,json=onlyActive" json:"only_active,omitempty"`
	OnlyActiveNull                  bool                        `protobuf:"varint,1003,opt,name=only_active_null,json=onlyActiveNull" json:"only_active_null,omitempty"`
	OnlyNotDeleted                  *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=only_not_deleted,json=onlyNotDeleted" json:"only_not_deleted,omitempty"`
	OnlyNotDeletedNull              bool                        `protobuf:"varint,1004,opt,name=only_not_deleted_null,json=onlyNotDeletedNull" json:"only_not_deleted_null,omitempty"`
	GroupByCharacteristicId         *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=group_by_characteristic_id,json=groupByCharacteristicId" json:"group_by_characteristic_id,omitempty"`
	GroupByCharacteristicIdNull     bool                        `protobuf:"varint,1005,opt,name=group_by_characteristic_id_null,json=groupByCharacteristicIdNull" json:"group_by_characteristic_id_null,omitempty"`
	BinaryCharacteristicValueId     *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=binary_characteristic_value_id,json=binaryCharacteristicValueId" json:"binary_characteristic_value_id,omitempty"`
	BinaryCharacteristicValueIdNull bool                        `protobuf:"varint,1006,opt,name=binary_characteristic_value_id_null,json=binaryCharacteristicValueIdNull" json:"binary_characteristic_value_id_null,omitempty"`
	FilterByCharacteristicId        *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=filter_by_characteristic_id,json=filterByCharacteristicId" json:"filter_by_characteristic_id,omitempty"`
	FilterByCharacteristicIdNull    bool                        `protobuf:"varint,1007,opt,name=filter_by_characteristic_id_null,json=filterByCharacteristicIdNull" json:"filter_by_characteristic_id_null,omitempty"`
	FilterByCharacValue             *dstore_values.StringValue  `protobuf:"bytes,8,opt,name=filter_by_charac_value,json=filterByCharacValue" json:"filter_by_charac_value,omitempty"`
	FilterByCharacValueNull         bool                        `protobuf:"varint,1008,opt,name=filter_by_charac_value_null,json=filterByCharacValueNull" json:"filter_by_charac_value_null,omitempty"`
	NegateFilterByParams            *dstore_values.BooleanValue `protobuf:"bytes,9,opt,name=negate_filter_by_params,json=negateFilterByParams" json:"negate_filter_by_params,omitempty"`
	NegateFilterByParamsNull        bool                        `protobuf:"varint,1009,opt,name=negate_filter_by_params_null,json=negateFilterByParamsNull" json:"negate_filter_by_params_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTreeNodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeNodeId
	}
	return nil
}

func (m *Parameters) GetLanguageId() *dstore_values.IntegerValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *Parameters) GetOnlyActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyActive
	}
	return nil
}

func (m *Parameters) GetOnlyNotDeleted() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyNotDeleted
	}
	return nil
}

func (m *Parameters) GetGroupByCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GroupByCharacteristicId
	}
	return nil
}

func (m *Parameters) GetBinaryCharacteristicValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCharacteristicValueId
	}
	return nil
}

func (m *Parameters) GetFilterByCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilterByCharacteristicId
	}
	return nil
}

func (m *Parameters) GetFilterByCharacValue() *dstore_values.StringValue {
	if m != nil {
		return m.FilterByCharacValue
	}
	return nil
}

func (m *Parameters) GetNegateFilterByParams() *dstore_values.BooleanValue {
	if m != nil {
		return m.NegateFilterByParams
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
	RowId           int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	NodeDescription *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=node_description,json=nodeDescription" json:"node_description,omitempty"`
	Active          *dstore_values.BooleanValue `protobuf:"bytes,10002,opt,name=active" json:"active,omitempty"`
	LevelNo         *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=level_no,json=levelNo" json:"level_no,omitempty"`
	BinaryCodeId    *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
	NodeId          *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TreeNodeId      *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=tree_node_id,json=treeNodeId" json:"tree_node_id,omitempty"`
	CountSuccessors *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=count_successors,json=countSuccessors" json:"count_successors,omitempty"`
	SortNo          *dstore_values.IntegerValue `protobuf:"bytes,10008,opt,name=sort_no,json=sortNo" json:"sort_no,omitempty"`
	InheritsFrom    *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=inherits_from,json=inheritsFrom" json:"inherits_from,omitempty"`
	SymbolId        *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=symbol_id,json=symbolId" json:"symbol_id,omitempty"`
	ValueSortNo     *dstore_values.IntegerValue `protobuf:"bytes,20003,opt,name=value_sort_no,json=valueSortNo" json:"value_sort_no,omitempty"`
	ValueId         *dstore_values.IntegerValue `protobuf:"bytes,20004,opt,name=value_id,json=valueId" json:"value_id,omitempty"`
	Value           *dstore_values.StringValue  `protobuf:"bytes,20010,opt,name=value" json:"value,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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

func (m *Response_Row) GetLevelNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelNo
	}
	return nil
}

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
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

func (m *Response_Row) GetCountSuccessors() *dstore_values.IntegerValue {
	if m != nil {
		return m.CountSuccessors
	}
	return nil
}

func (m *Response_Row) GetSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.SortNo
	}
	return nil
}

func (m *Response_Row) GetInheritsFrom() *dstore_values.IntegerValue {
	if m != nil {
		return m.InheritsFrom
	}
	return nil
}

func (m *Response_Row) GetSymbolId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SymbolId
	}
	return nil
}

func (m *Response_Row) GetValueSortNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueSortNo
	}
	return nil
}

func (m *Response_Row) GetValueId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ValueId
	}
	return nil
}

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_GetDirectSuccessors_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_GetDirectSuccessors_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_GetDirectSuccessors_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 908 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x49, 0x6f, 0x1c, 0x45,
	0x14, 0x96, 0x31, 0xb3, 0xf8, 0x39, 0xc4, 0xa6, 0x03, 0x71, 0x33, 0x13, 0x25, 0x51, 0x72, 0x60,
	0x39, 0x8c, 0x91, 0x21, 0xc2, 0x48, 0x2c, 0xb2, 0xf1, 0xa2, 0x1c, 0xdc, 0x40, 0x47, 0x42, 0x90,
	0x4b, 0xd3, 0xd3, 0x5d, 0x9e, 0x94, 0xd4, 0x53, 0x35, 0xaa, 0xaa, 0x71, 0xe4, 0x7f, 0xc1, 0xbe,
	0x29, 0x37, 0x38, 0xf1, 0x23, 0xf8, 0x23, 0x5c, 0xb8, 0xb2, 0xc3, 0x3f, 0xc8, 0xab, 0x7a, 0x35,
	0x9e, 0x25, 0x33, 0xee, 0xf6, 0x65, 0x46, 0xd5, 0xf5, 0xbe, 0xa5, 0x5e, 0x57, 0x7f, 0x55, 0xb0,
	0x9d, 0x6b, 0x23, 0x15, 0xdb, 0x64, 0xa2, 0xc7, 0x05, 0xdb, 0x1c, 0x28, 0x99, 0xb1, 0x7c, 0xa8,
	0x98, 0xde, 0xe4, 0xfd, 0xe4, 0x90, 0x99, 0x3d, 0xae, 0x58, 0x66, 0xee, 0x0d, 0xb3, 0x8c, 0x69,
	0x2d, 0x95, 0x4e, 0x76, 0xf2, 0x0e, 0x16, 0x19, 0x19, 0xbc, 0x48, 0xc8, 0x0e, 0x21, 0x3b, 0x0b,
	0xcb, 0x5b, 0x57, 0xbc, 0xc4, 0x49, 0x5a, 0x0c, 0x99, 0x26, 0x74, 0xeb, 0x85, 0x69, 0x5d, 0xa6,
	0x94, 0x54, 0x7e, 0xaa, 0x3d, 0x3d, 0xd5, 0x47, 0xa6, 0xb4, 0xc7, 0xfc, 0xe4, 0xed, 0xd9, 0x49,
	0x93, 0x72, 0x71, 0x2c, 0x55, 0x3f, 0x35, 0x5c, 0x0a, 0x2a, 0xba, 0xf5, 0xdb, 0x0a, 0xc0, 0x07,
	0xa9, 0x4a, 0x71, 0x96, 0x29, 0x1d, 0xbc, 0x0d, 0x97, 0x8c, 0x62, 0x2c, 0x11, 0x32, 0x67, 0x09,
	0xcf, 0xc3, 0xa5, 0x9b, 0x4b, 0x2f, 0xad, 0x6e, 0xb5, 0x3b, 0x7e, 0x01, 0xde, 0x17, 0x17, 0x86,
	0xf5, 0x98, 0xfa, 0xc8, 0x8e, 0x62, 0xb0, 0x80, 0x08, 0xeb, 0xef, 0xe6, 0xc1, 0x2b, 0xf0, 0xec,
	0x24, 0x3c, 0x11, 0xc3, 0xa2, 0x08, 0x7f, 0x6f, 0x20, 0x49, 0x33, 0xbe, 0x3c, 0xae, 0x8b, 0xf0,
	0x71, 0xf0, 0x16, 0xac, 0x16, 0xa9, 0xe8, 0x0d, 0xd1, 0xb0, 0x55, 0x7a, 0xaa, 0x82, 0xd2, 0xa8,
	0x1e, 0x95, 0x5e, 0x86, 0xf5, 0x09, 0x34, 0x09, 0xfd, 0xe1, 0x85, 0xc6, 0x65, 0x23, 0x21, 0x29,
	0x8a, 0xd3, 0x24, 0xcd, 0x0c, 0x3f, 0x61, 0xe1, 0xf2, 0x5c, 0xa1, 0xae, 0x94, 0x05, 0x4b, 0x85,
	0x17, 0xb2, 0xf5, 0x3b, 0xae, 0xdc, 0x0a, 0x4d, 0xa0, 0x49, 0xe8, 0x4f, 0x2f, 0x34, 0x2e, 0x73,
	0x42, 0xfb, 0xbe, 0x54, 0x48, 0x93, 0xe4, 0xac, 0xc0, 0x8e, 0xe6, 0xe1, 0xd3, 0xe5, 0x6a, 0x8e,
	0x26, 0x92, 0x66, 0x8f, 0x20, 0xc1, 0x16, 0x3c, 0x3f, 0x4b, 0x43, 0xb2, 0x7f, 0x91, 0x6c, 0x30,
	0x5d, 0xef, 0xa4, 0x3f, 0x86, 0x56, 0x4f, 0xc9, 0xe1, 0x20, 0xe9, 0x9e, 0x26, 0xd9, 0x03, 0x7c,
	0x9f, 0x19, 0xbe, 0x4e, 0xae, 0x0d, 0xcf, 0x6c, 0x6f, 0x6b, 0xe5, 0xbd, 0xdd, 0x70, 0xf0, 0xdd,
	0xd3, 0xf7, 0xa6, 0xc0, 0xd8, 0xe8, 0x7d, 0xb8, 0xb1, 0x98, 0x99, 0x7c, 0xfd, 0x4d, 0xbe, 0xda,
	0x0b, 0x28, 0x9c, 0xc1, 0x4f, 0xe1, 0x7a, 0x97, 0x8b, 0x54, 0x3d, 0x41, 0xe2, 0xcc, 0x58, 0x93,
	0xf5, 0x72, 0x93, 0x6d, 0xa2, 0x98, 0x16, 0x70, 0x53, 0x68, 0xf4, 0x08, 0x6e, 0x9f, 0xaf, 0x40,
	0x66, 0xff, 0x21, 0xb3, 0x37, 0xce, 0xa1, 0x72, 0x86, 0xef, 0x43, 0xfb, 0x98, 0x17, 0x38, 0x33,
	0xbf, 0xa5, 0x8d, 0x72, 0xb7, 0x21, 0xe1, 0xe7, 0xf4, 0xf4, 0x10, 0x6e, 0x9e, 0xc3, 0x4d, 0x3e,
	0xff, 0x25, 0x9f, 0xd7, 0x16, 0x91, 0x38, 0x93, 0xef, 0xc3, 0xd5, 0x59, 0x22, 0x5a, 0x6e, 0xd8,
	0x74, 0xfe, 0x5a, 0x33, 0xfe, 0xb4, 0x51, 0x5c, 0xf4, 0xc8, 0xde, 0x95, 0x69, 0x66, 0xf7, 0x10,
	0xbf, 0xff, 0xf6, 0x7c, 0x42, 0x32, 0xf5, 0x1f, 0x99, 0xda, 0x98, 0x03, 0x75, 0x7e, 0x62, 0xd8,
	0x10, 0xac, 0x97, 0x1a, 0x96, 0x8c, 0x59, 0x06, 0x36, 0x5d, 0x74, 0xb8, 0x52, 0xfe, 0x21, 0x3c,
	0x47, 0xd8, 0x03, 0x4f, 0xee, 0x62, 0x49, 0x07, 0xef, 0xc2, 0xb5, 0x05, 0x9c, 0xe4, 0xe9, 0x7f,
	0xf2, 0x14, 0xce, 0x03, 0x5b, 0x53, 0xb7, 0x7e, 0x6d, 0x42, 0x33, 0x66, 0x7a, 0x20, 0x85, 0x66,
	0xc1, 0xab, 0x50, 0x73, 0x01, 0xea, 0x93, 0xed, 0xac, 0x41, 0x3e, 0x9a, 0x29, 0x5c, 0xf7, 0xed,
	0x6f, 0x4c, 0x85, 0xc1, 0x27, 0xb0, 0x6e, 0xa3, 0x33, 0x99, 0xc8, 0x4e, 0x0c, 0xab, 0x65, 0x04,
	0x77, 0x66, 0xc0, 0xb3, 0x09, 0x7b, 0x84, 0xe3, 0xbb, 0xe3, 0x71, 0xbc, 0xd6, 0x9f, 0x7e, 0x10,
	0x6c, 0x43, 0xc3, 0x47, 0x36, 0xa6, 0x92, 0x65, 0xbc, 0xfe, 0x04, 0x23, 0x05, 0xfa, 0x11, 0xfd,
	0xc7, 0xa3, 0x72, 0xdc, 0x41, 0xcb, 0x4a, 0x3e, 0xc4, 0x74, 0xb1, 0xa8, 0x3b, 0x9d, 0x8a, 0xe7,
	0x4b, 0x67, 0xd4, 0x86, 0x4e, 0x2c, 0x1f, 0xc6, 0x96, 0xa1, 0xf5, 0x4b, 0x1d, 0x96, 0x71, 0x10,
	0x5c, 0x85, 0x3a, 0x0e, 0xed, 0xce, 0xfe, 0x2c, 0xc2, 0xce, 0xd4, 0xe2, 0x1a, 0x0e, 0x71, 0xab,
	0x1e, 0xc0, 0xba, 0x0b, 0xf3, 0x9c, 0xe9, 0x4c, 0xf1, 0x81, 0x5b, 0xfd, 0xe7, 0x51, 0xe9, 0xe6,
	0x5a, 0xb3, 0xa0, 0xbd, 0x31, 0x26, 0x78, 0x1d, 0xea, 0x3e, 0x7f, 0xbf, 0x88, 0xca, 0x77, 0x82,
	0xaf, 0x0d, 0xde, 0x80, 0x66, 0xc1, 0x4e, 0x58, 0x81, 0x59, 0x18, 0x7e, 0x19, 0x95, 0x7f, 0x72,
	0x0d, 0x57, 0x1d, 0xc9, 0x60, 0x17, 0x2e, 0x8f, 0xc2, 0xc0, 0x9f, 0x64, 0x5f, 0x55, 0x80, 0x5f,
	0xf2, 0xa1, 0x40, 0x87, 0xd9, 0x1d, 0x68, 0x8c, 0x8e, 0xc1, 0xaf, 0x2b, 0x80, 0xeb, 0x82, 0x60,
	0xef, 0xcc, 0x1c, 0xa1, 0xdf, 0x44, 0x17, 0x3b, 0x43, 0x0f, 0x61, 0x3d, 0x93, 0x43, 0x61, 0x12,
	0x7d, 0xf6, 0xea, 0xc2, 0x6f, 0x2b, 0x70, 0xac, 0x39, 0xd4, 0xf8, 0x7d, 0x5b, 0xff, 0xf8, 0x6f,
	0x6c, 0xef, 0xbe, 0xab, 0xe2, 0xdf, 0x16, 0x63, 0xeb, 0x76, 0xe0, 0x19, 0x2e, 0x1e, 0x60, 0xce,
	0x18, 0x9d, 0x1c, 0x2b, 0xd9, 0x0f, 0xbf, 0xaf, 0xd2, 0xb9, 0x11, 0xe4, 0x00, 0x11, 0xc1, 0x9b,
	0xb0, 0xa2, 0x4f, 0xfb, 0x5d, 0x59, 0xd8, 0xf5, 0xff, 0x50, 0x01, 0xde, 0xa4, 0x72, 0x5c, 0x3d,
	0xaa, 0x53, 0xde, 0x8c, 0xac, 0xff, 0xf8, 0xa8, 0xc2, 0x1d, 0x64, 0xd5, 0x3d, 0xbc, 0x47, 0x0b,
	0xd8, 0x86, 0xe6, 0xd9, 0xa1, 0xf2, 0x53, 0x15, 0x74, 0xe3, 0xc4, 0x1f, 0x21, 0x5b, 0x50, 0xa3,
	0xf4, 0xfc, 0xf9, 0xd1, 0x52, 0xe9, 0x16, 0xa7, 0xd2, 0xdd, 0x0f, 0xa1, 0xcd, 0xe5, 0xcc, 0x07,
	0x38, 0xbe, 0x1a, 0xde, 0xdf, 0xba, 0xf8, 0xa5, 0xb1, 0x5b, 0x77, 0x57, 0xb3, 0xd7, 0x1e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x84, 0xec, 0xdf, 0x61, 0x71, 0x0a, 0x00, 0x00,
}
