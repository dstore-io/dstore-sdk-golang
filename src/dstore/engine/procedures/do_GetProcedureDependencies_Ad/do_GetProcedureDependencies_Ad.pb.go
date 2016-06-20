// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetProcedureDependencies_Ad.proto
// DO NOT EDIT!

/*
Package do_GetProcedureDependencies_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetProcedureDependencies_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetProcedureDependencies_Ad

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
	ProcedureName                *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	ProcedureNameNull            bool                        `protobuf:"varint,1001,opt,name=procedure_name_null,json=procedureNameNull" json:"procedure_name_null,omitempty"`
	ProcedureCategoryId          *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=procedure_category_id,json=procedureCategoryId" json:"procedure_category_id,omitempty"`
	ProcedureCategoryIdNull      bool                        `protobuf:"varint,1002,opt,name=procedure_category_id_null,json=procedureCategoryIdNull" json:"procedure_category_id_null,omitempty"`
	AscertainUsedProcedures      *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=ascertain_used_procedures,json=ascertainUsedProcedures" json:"ascertain_used_procedures,omitempty"`
	AscertainUsedProceduresNull  bool                        `protobuf:"varint,1003,opt,name=ascertain_used_procedures_null,json=ascertainUsedProceduresNull" json:"ascertain_used_procedures_null,omitempty"`
	OrderByLevelOfDependence     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=order_by_level_of_dependence,json=orderByLevelOfDependence" json:"order_by_level_of_dependence,omitempty"`
	OrderByLevelOfDependenceNull bool                        `protobuf:"varint,1004,opt,name=order_by_level_of_dependence_null,json=orderByLevelOfDependenceNull" json:"order_by_level_of_dependence_null,omitempty"`
	OnlyUpToDependenceLevel      *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=only_up_to_dependence_level,json=onlyUpToDependenceLevel" json:"only_up_to_dependence_level,omitempty"`
	OnlyUpToDependenceLevelNull  bool                        `protobuf:"varint,1005,opt,name=only_up_to_dependence_level_null,json=onlyUpToDependenceLevelNull" json:"only_up_to_dependence_level_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func (m *Parameters) GetProcedureCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureCategoryId
	}
	return nil
}

func (m *Parameters) GetAscertainUsedProcedures() *dstore_values.BooleanValue {
	if m != nil {
		return m.AscertainUsedProcedures
	}
	return nil
}

func (m *Parameters) GetOrderByLevelOfDependence() *dstore_values.BooleanValue {
	if m != nil {
		return m.OrderByLevelOfDependence
	}
	return nil
}

func (m *Parameters) GetOnlyUpToDependenceLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.OnlyUpToDependenceLevel
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
	RowId             int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	LevelOfDependence *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=level_of_dependence,json=levelOfDependence" json:"level_of_dependence,omitempty"`
	UsedProcedure     *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=used_procedure,json=usedProcedure" json:"used_procedure,omitempty"`
	ProcedureName     *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetLevelOfDependence() *dstore_values.IntegerValue {
	if m != nil {
		return m.LevelOfDependence
	}
	return nil
}

func (m *Response_Row) GetUsedProcedure() *dstore_values.StringValue {
	if m != nil {
		return m.UsedProcedure
	}
	return nil
}

func (m *Response_Row) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetProcedureDependencies_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetProcedureDependencies_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetProcedureDependencies_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x4b, 0x6f, 0xd3, 0x4c,
	0x14, 0x55, 0xbe, 0x34, 0x6d, 0x75, 0x3f, 0x51, 0xe8, 0x44, 0x10, 0x37, 0xa9, 0xaa, 0x52, 0x36,
	0x2c, 0x90, 0x83, 0x60, 0x01, 0x42, 0xb0, 0x68, 0x28, 0x8f, 0x4a, 0xa9, 0x5b, 0x59, 0xb4, 0x08,
	0x58, 0x8c, 0x9c, 0xf8, 0x26, 0xb2, 0xe4, 0xcc, 0x44, 0x33, 0x76, 0xab, 0xfc, 0x0b, 0x1e, 0xe2,
	0x47, 0xf2, 0x5a, 0xb1, 0x61, 0xc9, 0x8c, 0xc7, 0x89, 0x63, 0x37, 0x0f, 0x95, 0x4d, 0xa2, 0x3b,
	0xf7, 0x9e, 0x73, 0xcf, 0xcc, 0xdc, 0x33, 0x86, 0x67, 0xbe, 0x8c, 0xb8, 0xc0, 0x26, 0xb2, 0x7e,
	0xc0, 0xb0, 0x39, 0x14, 0xbc, 0x8b, 0x7e, 0x2c, 0x50, 0x36, 0x7d, 0x4e, 0x5f, 0x61, 0x74, 0x32,
	0x5e, 0x38, 0xc0, 0x21, 0x32, 0x1f, 0x59, 0x37, 0x40, 0x49, 0xf7, 0x7d, 0x5b, 0x55, 0x46, 0x9c,
	0xdc, 0x33, 0x70, 0xdb, 0xc0, 0xed, 0xc5, 0x98, 0x7a, 0x35, 0x6d, 0x76, 0xee, 0x85, 0x31, 0x4a,
	0x43, 0x51, 0xdf, 0xca, 0x2b, 0x40, 0x21, 0xb8, 0x48, 0x53, 0x8d, 0x7c, 0x6a, 0x80, 0x52, 0x7a,
	0x7d, 0x4c, 0x93, 0x77, 0x8a, 0xc9, 0xc8, 0x0b, 0x58, 0x8f, 0x8b, 0x81, 0x17, 0x05, 0x9c, 0x99,
	0xa2, 0xbd, 0x3f, 0x15, 0x80, 0x13, 0x4f, 0x78, 0x2a, 0x8b, 0x42, 0x92, 0x7d, 0xd8, 0x98, 0xec,
	0x90, 0x32, 0xb5, 0x6a, 0x95, 0x76, 0x4b, 0x77, 0xff, 0x7f, 0x50, 0xb7, 0xd3, 0x7d, 0xa4, 0xca,
	0x64, 0x24, 0x02, 0xd6, 0x3f, 0xd3, 0x81, 0x7b, 0x6d, 0x82, 0x70, 0x14, 0x80, 0x34, 0xa1, 0x9a,
	0xa7, 0xa0, 0x2c, 0x0e, 0x43, 0xeb, 0xdb, 0x9a, 0x22, 0x5a, 0x77, 0x37, 0x73, 0xc5, 0x8e, 0xca,
	0x90, 0x63, 0xb8, 0x99, 0x01, 0xba, 0x5e, 0x84, 0x7d, 0x2e, 0x46, 0x34, 0xf0, 0xad, 0xff, 0x92,
	0xd6, 0x8d, 0x42, 0xeb, 0x80, 0xa9, 0x0a, 0x14, 0xa6, 0x77, 0xd6, 0xea, 0x79, 0x0a, 0x3c, 0xf4,
	0xc9, 0x53, 0xa8, 0xcf, 0x24, 0x34, 0x42, 0xbe, 0x1b, 0x21, 0xb5, 0x19, 0xc8, 0x44, 0xce, 0x5b,
	0xd8, 0xf2, 0x64, 0x17, 0x85, 0x3e, 0x2d, 0x1a, 0x4b, 0xf4, 0x69, 0x76, 0xe7, 0x56, 0x79, 0xa6,
	0xa4, 0x0e, 0xe7, 0x21, 0x7a, 0xcc, 0x48, 0xaa, 0x4d, 0xd0, 0xa7, 0x0a, 0x3c, 0xb9, 0x6a, 0x49,
	0x0e, 0x60, 0x67, 0x2e, 0xb1, 0x91, 0xf6, 0xc3, 0x48, 0x6b, 0xcc, 0x61, 0x48, 0xe4, 0x7d, 0x80,
	0x6d, 0x2e, 0x7c, 0x14, 0xb4, 0x33, 0xa2, 0x21, 0x9e, 0x63, 0x48, 0x79, 0x8f, 0xfa, 0xe3, 0x39,
	0x42, 0x6b, 0x65, 0xb9, 0x42, 0x2b, 0x21, 0x68, 0x8d, 0xda, 0x1a, 0x7e, 0xdc, 0x9b, 0x0c, 0x21,
	0x92, 0xd7, 0x70, 0x7b, 0x11, 0xb9, 0x51, 0xf9, 0xd3, 0xa8, 0xdc, 0x9e, 0xc7, 0x92, 0xc8, 0x7c,
	0x07, 0x0d, 0xce, 0xc2, 0x11, 0x8d, 0x87, 0x34, 0xe2, 0xd3, 0x14, 0x09, 0xad, 0x55, 0x59, 0x7e,
	0xb5, 0x35, 0x8d, 0x3f, 0x1d, 0xbe, 0xe1, 0x19, 0x73, 0xd2, 0x8a, 0xbc, 0x84, 0xdd, 0x05, 0xd4,
	0x46, 0xe3, 0xaf, 0xf4, 0x24, 0xe7, 0x70, 0x68, 0x89, 0x7b, 0x5f, 0x57, 0x60, 0xdd, 0x45, 0x39,
	0xe4, 0x4c, 0x22, 0xb9, 0x0f, 0x95, 0xc4, 0x58, 0xc5, 0x79, 0x4f, 0x7d, 0x6b, 0x4c, 0xf7, 0x42,
	0xff, 0xba, 0xa6, 0x50, 0xed, 0xf0, 0x86, 0xb6, 0x14, 0x9d, 0xf2, 0x94, 0x9a, 0xd8, 0xb2, 0x02,
	0xdb, 0x05, 0x70, 0xd1, 0x79, 0x47, 0x2a, 0x3e, 0xcc, 0x62, 0xf7, 0xfa, 0x20, 0xbf, 0x40, 0x1e,
	0xc3, 0x5a, 0x6a, 0x65, 0x35, 0x70, 0x9a, 0x71, 0xe7, 0x12, 0xa3, 0x31, 0xfa, 0x91, 0xf9, 0x77,
	0xc7, 0xe5, 0xa4, 0x0d, 0x65, 0xc1, 0x2f, 0xd4, 0x10, 0x68, 0xd4, 0x13, 0xfb, 0x2a, 0x8f, 0x8f,
	0x3d, 0x3e, 0x0b, 0xdb, 0xe5, 0x17, 0xae, 0xa6, 0xa9, 0xff, 0x2e, 0x41, 0x59, 0x05, 0xe4, 0x16,
	0xac, 0xaa, 0x50, 0x5b, 0xf2, 0xa3, 0xa3, 0x8e, 0xa7, 0xe2, 0x56, 0x54, 0xa8, 0x8c, 0xd6, 0x86,
	0xea, 0xac, 0x11, 0xfc, 0xe4, 0x2c, 0xbf, 0xdd, 0xcd, 0xf0, 0xd2, 0xf0, 0xb5, 0x60, 0x23, 0xef,
	0x0a, 0xeb, 0xb3, 0xb3, 0xfc, 0xf1, 0x89, 0xa7, 0x2d, 0xa2, 0x39, 0x0a, 0xef, 0xd7, 0x17, 0xe7,
	0x8a, 0x0f, 0x58, 0xeb, 0x0c, 0x1a, 0x01, 0x2f, 0x1c, 0x5d, 0xe6, 0xd4, 0xf7, 0x8f, 0xfe, 0xf1,
	0x83, 0xd0, 0x59, 0x4d, 0x5e, 0xdc, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x02, 0x73,
	0x95, 0x52, 0x06, 0x00, 0x00,
}
