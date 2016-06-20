// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_ModifyTRITriggerToDos_Ad.proto
// DO NOT EDIT!

/*
Package mi_ModifyTRITriggerToDos_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_ModifyTRITriggerToDos_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_ModifyTRITriggerToDos_Ad

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
	TriggerId         *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=trigger_id,json=triggerId" json:"trigger_id,omitempty"`
	TriggerIdNull     bool                        `protobuf:"varint,1001,opt,name=trigger_id_null,json=triggerIdNull" json:"trigger_id_null,omitempty"`
	WorkStepNo        *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=work_step_no,json=workStepNo" json:"work_step_no,omitempty"`
	WorkStepNoNull    bool                        `protobuf:"varint,1002,opt,name=work_step_no_null,json=workStepNoNull" json:"work_step_no_null,omitempty"`
	ToDoNo            *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=to_do_no,json=toDoNo" json:"to_do_no,omitempty"`
	ToDoNoNull        bool                        `protobuf:"varint,1003,opt,name=to_do_no_null,json=toDoNoNull" json:"to_do_no_null,omitempty"`
	ProcCall          *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=proc_call,json=procCall" json:"proc_call,omitempty"`
	ProcCallNull      bool                        `protobuf:"varint,1004,opt,name=proc_call_null,json=procCallNull" json:"proc_call_null,omitempty"`
	IterationList     *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=iteration_list,json=iterationList" json:"iteration_list,omitempty"`
	IterationListNull bool                        `protobuf:"varint,1005,opt,name=iteration_list_null,json=iterationListNull" json:"iteration_list_null,omitempty"`
	DeleteToDo        *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=delete_to_do,json=deleteToDo" json:"delete_to_do,omitempty"`
	DeleteToDoNull    bool                        `protobuf:"varint,1006,opt,name=delete_to_do_null,json=deleteToDoNull" json:"delete_to_do_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetTriggerId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TriggerId
	}
	return nil
}

func (m *Parameters) GetWorkStepNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.WorkStepNo
	}
	return nil
}

func (m *Parameters) GetToDoNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.ToDoNo
	}
	return nil
}

func (m *Parameters) GetProcCall() *dstore_values.StringValue {
	if m != nil {
		return m.ProcCall
	}
	return nil
}

func (m *Parameters) GetIterationList() *dstore_values.StringValue {
	if m != nil {
		return m.IterationList
	}
	return nil
}

func (m *Parameters) GetDeleteToDo() *dstore_values.BooleanValue {
	if m != nil {
		return m.DeleteToDo
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
	RowId int32 `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_ModifyTRITriggerToDos_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_ModifyTRITriggerToDos_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_ModifyTRITriggerToDos_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0xa6, 0x4d, 0x93, 0xa6, 0xc7, 0x26, 0xb5, 0x53, 0x90, 0x98, 0xa0, 0x48, 0x44, 0xbc, 0x3c,
	0x6c, 0xc4, 0x0b, 0x95, 0x82, 0x0f, 0xf5, 0xf2, 0x10, 0x30, 0x41, 0xb6, 0x41, 0xd0, 0x97, 0x61,
	0xdb, 0x9d, 0x86, 0xc1, 0xcd, 0x9e, 0x30, 0xb3, 0x31, 0xf8, 0x2f, 0xfc, 0x3f, 0xfe, 0x22, 0xaf,
	0x8f, 0x3e, 0x7b, 0x66, 0xcf, 0xe4, 0xb6, 0x8a, 0x91, 0xbe, 0x24, 0x3b, 0x7b, 0xbe, 0xcb, 0xce,
	0x9c, 0xef, 0x0c, 0x1c, 0xc5, 0x36, 0x43, 0xa3, 0x3a, 0x2a, 0x1d, 0xea, 0x54, 0x75, 0xc6, 0x06,
	0xcf, 0x54, 0x3c, 0x31, 0xca, 0x76, 0x46, 0x5a, 0xf6, 0x30, 0xd6, 0xe7, 0x1f, 0x07, 0x61, 0x77,
	0x60, 0xf4, 0x70, 0xa8, 0xcc, 0x00, 0x5f, 0xa0, 0x95, 0xc7, 0x71, 0x40, 0xb0, 0x0c, 0xc5, 0x5d,
	0xe6, 0x06, 0xcc, 0x0d, 0xfe, 0x41, 0x68, 0x1e, 0x78, 0x9b, 0x0f, 0x51, 0x32, 0x51, 0x96, 0xf9,
	0xcd, 0xab, 0xab, 0xde, 0xca, 0x18, 0x34, 0xbe, 0xd4, 0x5a, 0x2d, 0x8d, 0x94, 0xb5, 0xd1, 0x50,
	0xf9, 0xe2, 0xcd, 0x62, 0x31, 0x8b, 0x74, 0x7a, 0x8e, 0x66, 0x14, 0x65, 0x1a, 0x53, 0x06, 0xb5,
	0x7f, 0x6d, 0x01, 0xbc, 0x8e, 0x4c, 0x44, 0x55, 0x65, 0xac, 0x38, 0x02, 0xc8, 0xf8, 0x9b, 0xa4,
	0x8e, 0x1b, 0x1b, 0x37, 0x36, 0xee, 0x5c, 0x7a, 0xd0, 0x0a, 0xfc, 0x06, 0xfc, 0x57, 0xe9, 0x34,
	0x53, 0x04, 0x78, 0xe3, 0x56, 0xe1, 0x8e, 0x87, 0x77, 0x63, 0x71, 0x1b, 0xf6, 0x16, 0x5c, 0x99,
	0x4e, 0x92, 0xa4, 0xf1, 0x65, 0x9b, 0x14, 0xaa, 0x61, 0x6d, 0x0e, 0xea, 0xd3, 0x5b, 0xf1, 0x14,
	0x76, 0xa7, 0x68, 0xde, 0x4b, 0x9b, 0xa9, 0xb1, 0x4c, 0xb1, 0xb1, 0xb9, 0xde, 0x06, 0x1c, 0xe1,
	0x84, 0xf0, 0x7d, 0x14, 0xf7, 0x60, 0x7f, 0x99, 0xce, 0x4e, 0x5f, 0xd9, 0xa9, 0xbe, 0xc0, 0xe5,
	0x56, 0x8f, 0xa1, 0x9a, 0xa1, 0x8c, 0xd1, 0xd9, 0x94, 0xd6, 0xdb, 0x54, 0x32, 0xea, 0x04, 0x59,
	0xb4, 0xa1, 0x36, 0xa3, 0xb1, 0xfc, 0x37, 0x96, 0x07, 0xae, 0xe7, 0xd2, 0x87, 0xb0, 0xe3, 0x62,
	0x20, 0xcf, 0x22, 0xaa, 0x6f, 0xe5, 0xda, 0xcd, 0x82, 0xb6, 0xa5, 0x7d, 0xa7, 0x43, 0x96, 0xae,
	0x3a, 0xf0, 0x73, 0xc2, 0x8a, 0x5b, 0x50, 0x9f, 0x13, 0x59, 0xfd, 0x3b, 0xab, 0xef, 0xce, 0x20,
	0xb9, 0xfe, 0x31, 0xd4, 0x35, 0xf5, 0x24, 0x6f, 0x96, 0x4c, 0xb4, 0xcd, 0x1a, 0xe5, 0xb5, 0x26,
	0xb5, 0x39, 0xe3, 0x15, 0x11, 0x44, 0x07, 0x0e, 0x56, 0x25, 0xd8, 0xee, 0x07, 0xdb, 0xed, 0xaf,
	0x80, 0x67, 0x9d, 0x89, 0x55, 0x42, 0x51, 0x90, 0xf9, 0xf6, 0x1b, 0x95, 0xbf, 0x1e, 0xd9, 0x29,
	0x62, 0xa2, 0xa2, 0xd4, 0x77, 0x86, 0x09, 0x2e, 0xc2, 0xae, 0x33, 0xcb, 0x74, 0x76, 0xfb, 0xe9,
	0x3b, 0xb3, 0xc0, 0x39, 0xab, 0xf6, 0xe7, 0x4d, 0xa8, 0x86, 0xca, 0x8e, 0x31, 0xb5, 0x4a, 0xdc,
	0x87, 0x72, 0x1e, 0x6b, 0x9f, 0xb8, 0xf9, 0x16, 0xfd, 0xc8, 0x70, 0xe4, 0x5f, 0xba, 0xdf, 0x90,
	0x81, 0xe2, 0x2d, 0x5c, 0x76, 0x81, 0x96, 0x4b, 0x89, 0xa6, 0x1c, 0x95, 0x88, 0x1c, 0x14, 0xc8,
	0xc5, 0xdc, 0xf7, 0x68, 0xdd, 0x5d, 0xac, 0xc3, 0xbd, 0xd1, 0xea, 0x0b, 0xf1, 0x04, 0xb6, 0xfd,
	0x20, 0x51, 0x64, 0x9c, 0xe2, 0xf5, 0x3f, 0x14, 0x79, 0xcc, 0x7a, 0xfc, 0x1f, 0xce, 0xe0, 0xa2,
	0x0b, 0x25, 0x83, 0x53, 0x0a, 0x83, 0x63, 0x1d, 0x06, 0xff, 0x3d, 0xf7, 0xc1, 0xec, 0x20, 0x82,
	0x10, 0xa7, 0xa1, 0xd3, 0x68, 0x5e, 0x83, 0x12, 0x3d, 0x8b, 0x2b, 0x50, 0xa1, 0x95, 0x9b, 0xc5,
	0x4f, 0x7d, 0x3a, 0x9a, 0x72, 0x58, 0xa6, 0x65, 0x37, 0x7e, 0x76, 0x02, 0x2d, 0x8d, 0x05, 0x83,
	0xc5, 0xa5, 0xf4, 0xee, 0xd1, 0x45, 0xae, 0xab, 0xd3, 0x4a, 0x7e, 0x25, 0x3c, 0xfc, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x05, 0x7d, 0xb8, 0xed, 0x04, 0x00, 0x00,
}
