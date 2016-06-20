// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyTRITrigger_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyTRITrigger_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyTRITrigger_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyTRITrigger_Ad

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
	TriggerTypeId            *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=trigger_type_id,json=triggerTypeId" json:"trigger_type_id,omitempty"`
	TriggerTypeIdNull        bool                        `protobuf:"varint,1001,opt,name=trigger_type_id_null,json=triggerTypeIdNull" json:"trigger_type_id_null,omitempty"`
	TriggerName              *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=trigger_name,json=triggerName" json:"trigger_name,omitempty"`
	TriggerNameNull          bool                        `protobuf:"varint,1002,opt,name=trigger_name_null,json=triggerNameNull" json:"trigger_name_null,omitempty"`
	TriggerDescription       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=trigger_description,json=triggerDescription" json:"trigger_description,omitempty"`
	TriggerDescriptionNull   bool                        `protobuf:"varint,1003,opt,name=trigger_description_null,json=triggerDescriptionNull" json:"trigger_description_null,omitempty"`
	ExecuteProcedureName     *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=execute_procedure_name,json=executeProcedureName" json:"execute_procedure_name,omitempty"`
	ExecuteProcedureNameNull bool                        `protobuf:"varint,1004,opt,name=execute_procedure_name_null,json=executeProcedureNameNull" json:"execute_procedure_name_null,omitempty"`
	DeleteTrigger            *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=delete_trigger,json=deleteTrigger" json:"delete_trigger,omitempty"`
	DeleteTriggerNull        bool                        `protobuf:"varint,1005,opt,name=delete_trigger_null,json=deleteTriggerNull" json:"delete_trigger_null,omitempty"`
	TRITriggerDefinition     *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=t_r_i_trigger_definition,json=tRITriggerDefinition" json:"t_r_i_trigger_definition,omitempty"`
	TRITriggerDefinitionNull bool                        `protobuf:"varint,1006,opt,name=t_r_i_trigger_definition_null,json=tRITriggerDefinitionNull" json:"t_r_i_trigger_definition_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTriggerTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TriggerTypeId
	}
	return nil
}

func (m *Parameters) GetTriggerName() *dstore_values.StringValue {
	if m != nil {
		return m.TriggerName
	}
	return nil
}

func (m *Parameters) GetTriggerDescription() *dstore_values.StringValue {
	if m != nil {
		return m.TriggerDescription
	}
	return nil
}

func (m *Parameters) GetExecuteProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ExecuteProcedureName
	}
	return nil
}

func (m *Parameters) GetDeleteTrigger() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteTrigger
	}
	return nil
}

func (m *Parameters) GetTRITriggerDefinition() *dstore_values.StringValue {
	if m != nil {
		return m.TRITriggerDefinition
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	TriggerId       *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=trigger_id,json=triggerId" json:"trigger_id,omitempty"`
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

func (m *Response) GetTriggerId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TriggerId
	}
	return nil
}

type Response_Row struct {
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyTRITrigger_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyTRITrigger_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyTRITrigger_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0x69, 0x63, 0xd2, 0x3a, 0x55, 0x6b, 0x37, 0x25, 0x9c, 0x09, 0x15, 0xa9, 0x3e, 0x08,
	0xc2, 0x45, 0x2a, 0x82, 0x0a, 0x8a, 0xd6, 0xf8, 0x10, 0x24, 0x21, 0x2c, 0x41, 0xd0, 0x97, 0xe3,
	0x9a, 0x9b, 0x84, 0x85, 0x64, 0x37, 0xec, 0x5e, 0xac, 0xfd, 0x2f, 0x7c, 0xf0, 0x9f, 0xf4, 0xe7,
	0x83, 0x7f, 0x81, 0x73, 0xb7, 0x7b, 0x49, 0xee, 0x7a, 0x12, 0x7d, 0x69, 0xd9, 0xec, 0xf7, 0x33,
	0xf3, 0xbd, 0xd9, 0x99, 0x81, 0xc7, 0x91, 0x89, 0x95, 0xc6, 0x36, 0xca, 0x89, 0x90, 0xd8, 0x9e,
	0x6b, 0x35, 0xc2, 0x68, 0xa1, 0xd1, 0xb4, 0x67, 0x22, 0xe8, 0xa9, 0x48, 0x8c, 0x2f, 0x86, 0xbc,
	0x3b, 0xd4, 0x62, 0x32, 0x41, 0x1d, 0xbc, 0x8a, 0x7c, 0x52, 0xc4, 0x8a, 0xdd, 0xb3, 0x98, 0x6f,
	0x31, 0xbf, 0x5c, 0xdb, 0xac, 0xbb, 0xe0, 0x1f, 0xc3, 0xe9, 0x02, 0x8d, 0x45, 0x9b, 0xb7, 0xf2,
	0x19, 0x51, 0x6b, 0xa5, 0xdd, 0x55, 0x2b, 0x7f, 0x35, 0x43, 0x63, 0xc2, 0x09, 0xba, 0xcb, 0xbb,
	0xc5, 0xcb, 0x38, 0x14, 0x72, 0xac, 0xf4, 0x2c, 0x8c, 0x85, 0x92, 0x56, 0x74, 0xfc, 0xa5, 0x06,
	0x30, 0x08, 0x75, 0x48, 0xb7, 0xa8, 0x0d, 0x7b, 0x0d, 0xfb, 0xb1, 0xb3, 0x13, 0x5f, 0xcc, 0x31,
	0x10, 0x91, 0xb7, 0x75, 0x67, 0xeb, 0xfe, 0xde, 0x49, 0xcb, 0x77, 0x1f, 0xe0, 0xac, 0x09, 0x19,
	0x23, 0xa9, 0xde, 0x25, 0x27, 0x7e, 0xdd, 0x31, 0x43, 0x42, 0xba, 0x11, 0x7b, 0x08, 0x87, 0x85,
	0x20, 0x81, 0x5c, 0x4c, 0xa7, 0xde, 0xd7, 0x1d, 0x0a, 0xb5, 0xcb, 0x0f, 0x72, 0xea, 0x3e, 0xdd,
	0xb0, 0xe7, 0x70, 0x2d, 0x23, 0x24, 0x59, 0xf1, 0xb6, 0xd3, 0x9c, 0xcd, 0x42, 0x4e, 0x43, 0x1a,
	0x39, 0xb1, 0x29, 0xf7, 0x9c, 0xbe, 0x4f, 0x72, 0xf6, 0x00, 0x0e, 0xd6, 0x71, 0x9b, 0xed, 0x9b,
	0xcd, 0xb6, 0xbf, 0x26, 0x4c, 0x73, 0xbd, 0x85, 0x7a, 0x26, 0x8e, 0xd0, 0x8c, 0xb4, 0x98, 0x27,
	0xe5, 0xf0, 0x2a, 0x1b, 0x53, 0x32, 0x87, 0x75, 0x56, 0x14, 0x7b, 0x0a, 0x5e, 0x49, 0x30, 0x6b,
	0xe0, 0xbb, 0x35, 0xd0, 0xb8, 0x8c, 0xa5, 0x3e, 0x06, 0xd0, 0xc0, 0x4f, 0x38, 0x5a, 0xc4, 0x18,
	0x2c, 0x9b, 0xc8, 0x7e, 0xfd, 0x95, 0x8d, 0x56, 0x0e, 0x1d, 0x39, 0xc8, 0xc0, 0xb4, 0x0c, 0x2f,
	0xa0, 0x55, 0x1e, 0xd1, 0xfa, 0xf9, 0x61, 0xfd, 0x78, 0x65, 0x6c, 0xea, 0xe8, 0x14, 0x6e, 0x44,
	0x38, 0xa5, 0x46, 0x08, 0x9c, 0x65, 0xaf, 0x5a, 0xfa, 0xf6, 0x67, 0x4a, 0x4d, 0x31, 0x94, 0xee,
	0xed, 0x2d, 0xe2, 0x9a, 0x98, 0xb5, 0xa1, 0x9e, 0x8f, 0x61, 0x73, 0xff, 0x74, 0x4f, 0x9f, 0x13,
	0xa7, 0x49, 0x39, 0x55, 0x30, 0xd0, 0x81, 0x08, 0x56, 0x75, 0x1c, 0x0b, 0x29, 0xd2, 0x37, 0xa9,
	0x6d, 0x2e, 0x44, 0xbc, 0x9c, 0x9f, 0xce, 0x92, 0x63, 0x2f, 0xe1, 0xe8, 0x6f, 0x31, 0xad, 0x9d,
	0x5f, 0xae, 0x14, 0x65, 0x74, 0xe2, 0xea, 0xf8, 0xf7, 0x36, 0xec, 0x72, 0x34, 0x73, 0x25, 0x0d,
	0x52, 0x3f, 0x57, 0xd3, 0xa1, 0x73, 0xa3, 0xb0, 0xf4, 0xe3, 0x66, 0xd9, 0x0e, 0xe4, 0x9b, 0xe4,
	0x2f, 0xb7, 0x42, 0xf6, 0x1e, 0x6e, 0x26, 0xe3, 0x16, 0xac, 0xcd, 0x1b, 0xf5, 0x74, 0x85, 0x60,
	0xbf, 0x00, 0x17, 0xa7, 0xb2, 0x47, 0xe7, 0xee, 0xea, 0xcc, 0xf7, 0x67, 0xf9, 0x1f, 0xd8, 0x13,
	0xd8, 0x71, 0x63, 0x4e, 0x2d, 0x9b, 0x44, 0xbc, 0x7d, 0x29, 0xa2, 0x5d, 0x02, 0x3d, 0xfb, 0x9f,
	0x67, 0x72, 0xd6, 0x81, 0x8a, 0x56, 0xe7, 0xd4, 0x5d, 0x09, 0x75, 0xe2, 0xff, 0xcb, 0x42, 0xf2,
	0xb3, 0x1a, 0xf8, 0x5c, 0x9d, 0xf3, 0x04, 0x67, 0xcf, 0x00, 0xb2, 0xaa, 0xd2, 0x72, 0xc0, 0xcd,
	0xcb, 0xe1, 0xaa, 0x93, 0x77, 0xa3, 0xe6, 0x11, 0x54, 0x28, 0x0e, 0x6b, 0x40, 0x8d, 0x22, 0x25,
	0xf8, 0xe7, 0x3e, 0xf1, 0x55, 0x5e, 0xa5, 0x63, 0x37, 0x3a, 0xed, 0x43, 0x4b, 0xa8, 0x82, 0xaf,
	0xd5, 0x7e, 0xfd, 0xd0, 0xfe, 0xcf, 0xcd, 0x7b, 0x56, 0x4b, 0x57, 0xdc, 0xa3, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0xc4, 0xd7, 0x57, 0xb3, 0x05, 0x00, 0x00,
}
