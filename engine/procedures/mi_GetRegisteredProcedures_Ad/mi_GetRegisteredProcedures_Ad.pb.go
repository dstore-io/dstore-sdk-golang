// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetRegisteredProcedures_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetRegisteredProcedures_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetRegisteredProcedures_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetRegisteredProcedures_Ad

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
	ProcedureId               *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=procedure_id,json=procedureId" json:"procedure_id,omitempty"`
	ProcedureIdNull           bool                        `protobuf:"varint,1001,opt,name=procedure_id_null,json=procedureIdNull" json:"procedure_id_null,omitempty"`
	ProcedureName             *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	ProcedureNameNull         bool                        `protobuf:"varint,1002,opt,name=procedure_name_null,json=procedureNameNull" json:"procedure_name_null,omitempty"`
	SearchWithLike            *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=search_with_like,json=searchWithLike" json:"search_with_like,omitempty"`
	SearchWithLikeNull        bool                        `protobuf:"varint,1003,opt,name=search_with_like_null,json=searchWithLikeNull" json:"search_with_like_null,omitempty"`
	ExecRightConfigurable     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=exec_right_configurable,json=execRightConfigurable" json:"exec_right_configurable,omitempty"`
	ExecRightConfigurableNull bool                        `protobuf:"varint,1004,opt,name=exec_right_configurable_null,json=execRightConfigurableNull" json:"exec_right_configurable_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetProcedureId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureId
	}
	return nil
}

func (m *Parameters) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func (m *Parameters) GetSearchWithLike() *dstore_values.BooleanValue {
	if m != nil {
		return m.SearchWithLike
	}
	return nil
}

func (m *Parameters) GetExecRightConfigurable() *dstore_values.BooleanValue {
	if m != nil {
		return m.ExecRightConfigurable
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
	RowId                         int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ProcedureId                   *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=procedure_id,json=procedureId" json:"procedure_id,omitempty"`
	CheckForExecutionRestrictions *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=check_for_execution_restrictions,json=checkForExecutionRestrictions" json:"check_for_execution_restrictions,omitempty"`
	LogExecutions                 *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=log_executions,json=logExecutions" json:"log_executions,omitempty"`
	ProcedureName                 *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetProcedureId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureId
	}
	return nil
}

func (m *Response_Row) GetCheckForExecutionRestrictions() *dstore_values.IntegerValue {
	if m != nil {
		return m.CheckForExecutionRestrictions
	}
	return nil
}

func (m *Response_Row) GetLogExecutions() *dstore_values.IntegerValue {
	if m != nil {
		return m.LogExecutions
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetRegisteredProcedures_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetRegisteredProcedures_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetRegisteredProcedures_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetRegisteredProcedures_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x4f, 0x13, 0x41,
	0x10, 0x0e, 0x94, 0x02, 0x59, 0xa4, 0xc0, 0x12, 0xf4, 0x68, 0xd5, 0x10, 0xfc, 0x62, 0x42, 0x72,
	0x35, 0xf8, 0x45, 0x13, 0x23, 0x02, 0xa9, 0x86, 0x28, 0x0d, 0x59, 0x13, 0x8d, 0x7e, 0x70, 0x73,
	0xbd, 0x1b, 0xae, 0x9b, 0x5e, 0x6f, 0xc9, 0xee, 0xd5, 0xfa, 0x33, 0x7c, 0xfb, 0x5d, 0xfe, 0x0f,
	0x5f, 0x12, 0x7f, 0x82, 0xce, 0xde, 0x5e, 0xaf, 0xbd, 0x03, 0x85, 0x7e, 0x81, 0xce, 0xcd, 0x3c,
	0xcf, 0x33, 0xbb, 0xf3, 0xcc, 0x92, 0x47, 0x81, 0x4e, 0xa4, 0x82, 0x26, 0xc4, 0xa1, 0x88, 0xa1,
	0x79, 0xa6, 0xa4, 0x0f, 0xc1, 0x40, 0x81, 0x6e, 0xf6, 0x05, 0x7f, 0x06, 0x09, 0x83, 0x50, 0xe8,
	0x04, 0x14, 0x04, 0x27, 0x79, 0x8a, 0xef, 0x07, 0x2e, 0x16, 0x26, 0x92, 0xee, 0x58, 0xb4, 0x6b,
	0xd1, 0xee, 0x7f, 0x21, 0xf5, 0xf5, 0x4c, 0xea, 0xbd, 0x17, 0x0d, 0x40, 0x5b, 0x86, 0xfa, 0x66,
	0x51, 0x1f, 0x94, 0x92, 0x2a, 0x4b, 0x35, 0x8a, 0xa9, 0x3e, 0x68, 0xed, 0x85, 0x90, 0x25, 0xef,
	0x94, 0x93, 0x89, 0x27, 0xe2, 0x53, 0xa9, 0xfa, 0x5e, 0x22, 0x64, 0x6c, 0x8b, 0xb6, 0xff, 0x54,
	0x08, 0x39, 0xf1, 0x94, 0x87, 0x59, 0x50, 0x9a, 0x3e, 0x26, 0xd7, 0xf2, 0xf3, 0x71, 0x11, 0x38,
	0x33, 0x5b, 0x33, 0x77, 0x97, 0x76, 0x1b, 0x6e, 0x76, 0x88, 0xac, 0x2f, 0x11, 0x27, 0x10, 0x82,
	0x7a, 0x65, 0x22, 0xb6, 0x94, 0x03, 0x8e, 0x02, 0xba, 0x43, 0xd6, 0x26, 0xf1, 0x3c, 0x1e, 0x44,
	0x91, 0xf3, 0x7d, 0x01, 0x59, 0x16, 0xd9, 0xca, 0x44, 0x61, 0x1b, 0xbf, 0xd3, 0x7d, 0x52, 0x1b,
	0x17, 0xc7, 0xd8, 0x82, 0x33, 0x9b, 0xca, 0xd5, 0x4b, 0x72, 0x3a, 0x51, 0x22, 0x0e, 0xad, 0xda,
	0x72, 0x8e, 0x68, 0x23, 0x80, 0x36, 0xc9, 0x7a, 0x91, 0xc2, 0x2a, 0xfe, 0xb0, 0x8a, 0x6b, 0x85,
	0xe2, 0x54, 0xb3, 0x45, 0x56, 0x35, 0x78, 0xca, 0xef, 0xf2, 0xa1, 0x48, 0xba, 0x3c, 0x12, 0x3d,
	0x70, 0x2a, 0x17, 0x1e, 0xb2, 0x23, 0x65, 0x04, 0x5e, 0x6c, 0x65, 0x6b, 0x16, 0xf4, 0x1a, 0x31,
	0x2f, 0x10, 0x42, 0x77, 0xc9, 0x46, 0x99, 0xc6, 0x2a, 0xff, 0xb4, 0xca, 0xb4, 0x58, 0x9f, 0x4a,
	0xbf, 0x24, 0x37, 0xe0, 0x03, 0xf8, 0x5c, 0x89, 0xb0, 0x9b, 0x70, 0x5f, 0xc6, 0xa7, 0x22, 0x1c,
	0x28, 0xaf, 0x13, 0x81, 0x33, 0x77, 0x79, 0x07, 0x1b, 0x06, 0xcb, 0x0c, 0xf4, 0x70, 0x02, 0x49,
	0x9f, 0x90, 0x9b, 0xff, 0x20, 0xb5, 0xfd, 0xfc, 0xb2, 0xfd, 0x6c, 0x5e, 0x88, 0x36, 0x6d, 0x6d,
	0xff, 0x9e, 0x23, 0x8b, 0x0c, 0xf4, 0x99, 0x8c, 0x35, 0xd0, 0x7b, 0xa4, 0x9a, 0xfa, 0x2b, 0x1b,
	0x7c, 0x3e, 0x89, 0xcc, 0xbd, 0xd6, 0x7b, 0x2d, 0xf3, 0x97, 0xd9, 0x42, 0xfa, 0x86, 0xac, 0x1a,
	0x67, 0xf1, 0x09, 0x6b, 0xe1, 0x18, 0x2b, 0x08, 0x76, 0x4b, 0xe0, 0xb2, 0x01, 0x8f, 0x31, 0x3e,
	0x1a, 0xc7, 0x6c, 0xa5, 0x5f, 0xfc, 0x40, 0x1f, 0x90, 0x85, 0xcc, 0xd1, 0x38, 0x22, 0xc3, 0x78,
	0xfb, 0x1c, 0xa3, 0xf5, 0xfb, 0xb1, 0xfd, 0xcf, 0x46, 0xe5, 0xf4, 0x39, 0xa9, 0x28, 0x39, 0xc4,
	0x6b, 0x35, 0xa8, 0x87, 0xee, 0x14, 0x2b, 0xe8, 0x8e, 0xae, 0xc2, 0x65, 0x72, 0xc8, 0x0c, 0x4b,
	0xfd, 0xdb, 0x2c, 0xa9, 0x60, 0x40, 0xaf, 0x93, 0x79, 0x0c, 0xcd, 0x56, 0x7c, 0x6c, 0xe3, 0xed,
	0x54, 0x59, 0x15, 0x43, 0xf4, 0xfc, 0x5e, 0x69, 0x67, 0x3e, 0xb5, 0xa7, 0x5c, 0x1a, 0x20, 0x5b,
	0x7e, 0x17, 0xfc, 0x1e, 0xc7, 0xa3, 0x73, 0x33, 0xa8, 0x81, 0x39, 0x3e, 0xc7, 0x86, 0xd0, 0xf8,
	0xbe, 0xf9, 0xad, 0x9d, 0xcf, 0x57, 0x20, 0xbd, 0x95, 0xb2, 0x3c, 0x95, 0xaa, 0x35, 0xe2, 0x60,
	0x13, 0x14, 0xf4, 0x90, 0xd4, 0x22, 0x19, 0x8e, 0x05, 0xb4, 0xf3, 0xe5, 0x0a, 0xa4, 0xcb, 0x88,
	0xc9, 0xf9, 0x34, 0x3d, 0x38, 0xb7, 0xb3, 0x5f, 0xdb, 0x53, 0x2e, 0xed, 0xc1, 0x3b, 0xd2, 0x10,
	0xb2, 0x34, 0x94, 0xf1, 0xab, 0xfa, 0x76, 0x2f, 0x94, 0x3a, 0xe8, 0x8d, 0xf2, 0xc1, 0xd4, 0x0f,
	0x6f, 0x67, 0x3e, 0x7d, 0xda, 0xee, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x27, 0x60, 0x35, 0x14,
	0xb9, 0x05, 0x00, 0x00,
}