// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetProcExecRights_Group_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetProcExecRights_Group_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetProcExecRights_Group_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetProcExecRights_Group_Ad

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
	UserGroupId             *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	UserGroupIdNull         bool                        `protobuf:"varint,1001,opt,name=user_group_id_null,json=userGroupIdNull" json:"user_group_id_null,omitempty"`
	ProcedureId             *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=procedure_id,json=procedureId" json:"procedure_id,omitempty"`
	ProcedureIdNull         bool                        `protobuf:"varint,1002,opt,name=procedure_id_null,json=procedureIdNull" json:"procedure_id_null,omitempty"`
	ProcedureCategoryId     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=procedure_category_id,json=procedureCategoryId" json:"procedure_category_id,omitempty"`
	ProcedureCategoryIdNull bool                        `protobuf:"varint,1003,opt,name=procedure_category_id_null,json=procedureCategoryIdNull" json:"procedure_category_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetUserGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserGroupId
	}
	return nil
}

func (m *Parameters) GetProcedureId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureId
	}
	return nil
}

func (m *Parameters) GetProcedureCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureCategoryId
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
	RowId          int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	GroupName      *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	ExecutionRight *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=execution_right,json=executionRight" json:"execution_right,omitempty"`
	ProcedureId    *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=procedure_id,json=procedureId" json:"procedure_id,omitempty"`
	ProcedureName  *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	UserGroupId    *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetGroupName() *dstore_values.StringValue {
	if m != nil {
		return m.GroupName
	}
	return nil
}

func (m *Response_Row) GetExecutionRight() *dstore_values.IntegerValue {
	if m != nil {
		return m.ExecutionRight
	}
	return nil
}

func (m *Response_Row) GetProcedureId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ProcedureId
	}
	return nil
}

func (m *Response_Row) GetProcedureName() *dstore_values.StringValue {
	if m != nil {
		return m.ProcedureName
	}
	return nil
}

func (m *Response_Row) GetUserGroupId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserGroupId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetProcExecRights_Group_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetProcExecRights_Group_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetProcExecRights_Group_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetProcExecRights_Group_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x13, 0xd2, 0x96, 0x29, 0x6d, 0x60, 0x2b, 0x20, 0x38, 0x12, 0x42, 0xe5, 0x82, 0x54,
	0xe4, 0x20, 0xb8, 0x00, 0xaa, 0x28, 0x14, 0x45, 0x55, 0x85, 0x6a, 0xaa, 0x3d, 0x20, 0xc1, 0x01,
	0xcb, 0xc4, 0x83, 0xb1, 0x48, 0xbc, 0xd1, 0xae, 0x4d, 0xe1, 0x0d, 0x38, 0xf2, 0xfb, 0x8c, 0x08,
	0x78, 0x00, 0xae, 0xcc, 0xfe, 0xc4, 0x89, 0x4d, 0x45, 0x9a, 0x8b, 0xad, 0xd9, 0x99, 0x6f, 0xbe,
	0xf9, 0xf9, 0x76, 0x61, 0x27, 0x56, 0xb9, 0x90, 0xd8, 0xc3, 0x2c, 0x49, 0x33, 0xec, 0x8d, 0xa5,
	0x18, 0x60, 0x5c, 0x48, 0x54, 0xbd, 0x51, 0x1a, 0xee, 0x63, 0x7e, 0x44, 0x07, 0xfd, 0xf7, 0x38,
	0xe0, 0x69, 0xf2, 0x26, 0x57, 0xe1, 0xbe, 0x14, 0xc5, 0x38, 0x7c, 0x14, 0xfb, 0x14, 0x98, 0x0b,
	0xb6, 0x6d, 0xd1, 0xbe, 0x45, 0xfb, 0xff, 0x85, 0x78, 0x9b, 0x8e, 0xea, 0x5d, 0x34, 0x2c, 0x50,
	0xd9, 0x0c, 0xde, 0x95, 0x2a, 0x3f, 0x4a, 0x29, 0xa4, 0x73, 0x75, 0xab, 0xae, 0x11, 0x2a, 0x15,
	0x25, 0xe8, 0x9c, 0xd7, 0xeb, 0xce, 0x3c, 0x4a, 0xb3, 0xd7, 0x42, 0x8e, 0xa2, 0x3c, 0x15, 0x99,
	0x0d, 0xda, 0xfa, 0xd3, 0x00, 0x38, 0x8a, 0x64, 0x44, 0x5e, 0x94, 0x8a, 0xed, 0xc2, 0x7a, 0xa1,
	0x50, 0x86, 0x89, 0xa9, 0x28, 0x8d, 0x3b, 0x4b, 0xd7, 0x96, 0x6e, 0xac, 0xdd, 0xee, 0xfa, 0xae,
	0x0b, 0x57, 0x58, 0x9a, 0xe5, 0x98, 0xa0, 0x7c, 0xa6, 0x2d, 0xbe, 0xa6, 0x11, 0xa6, 0x85, 0x83,
	0x98, 0xdd, 0x04, 0x56, 0x49, 0x10, 0x66, 0xc5, 0x70, 0xd8, 0xf9, 0xb9, 0x42, 0x69, 0x56, 0x79,
	0x7b, 0x26, 0x32, 0xa0, 0x73, 0xf6, 0x00, 0xce, 0x95, 0xe3, 0xd4, 0x6c, 0x8d, 0x53, 0xb0, 0x95,
	0x00, 0x62, 0xdb, 0x86, 0x0b, 0xb3, 0x78, 0x4b, 0xf6, 0xcb, 0x91, 0xcd, 0x04, 0x1a, 0xb2, 0xa7,
	0x70, 0x71, 0x1a, 0x3c, 0x88, 0x28, 0xa7, 0x90, 0x1f, 0x34, 0x6b, 0x73, 0x3e, 0xeb, 0x66, 0x89,
	0x7c, 0xec, 0x80, 0xc4, 0xbe, 0x03, 0xde, 0x89, 0x09, 0x6d, 0x19, 0xbf, 0x6d, 0x19, 0x97, 0x4f,
	0x40, 0xea, 0x72, 0xb6, 0x3e, 0xb6, 0x60, 0x95, 0xa3, 0x1a, 0x8b, 0x4c, 0x21, 0xbb, 0x05, 0x2d,
	0xb3, 0x57, 0x37, 0x6f, 0xcf, 0xaf, 0xaa, 0xc6, 0xee, 0xbc, 0xaf, 0xbf, 0xdc, 0x06, 0xb2, 0xe7,
	0x70, 0x5e, 0x6f, 0x34, 0x9c, 0x59, 0x29, 0x8d, 0xaf, 0x49, 0x60, 0xbf, 0x06, 0xae, 0x2f, 0xfe,
	0x90, 0xec, 0x83, 0xa9, 0xcd, 0xdb, 0xa3, 0xea, 0x01, 0xbb, 0x0b, 0x2b, 0x4e, 0x49, 0x34, 0x1a,
	0x9d, 0xf1, 0xea, 0x3f, 0x19, 0xad, 0xce, 0x0e, 0xed, 0x9f, 0x4f, 0xc2, 0xd9, 0x13, 0x68, 0x4a,
	0x71, 0xdc, 0x39, 0x63, 0x50, 0xf7, 0xfc, 0x05, 0xa4, 0xef, 0x4f, 0x46, 0xe1, 0x73, 0x71, 0xcc,
	0x75, 0x16, 0xef, 0x47, 0x03, 0x9a, 0x64, 0xb0, 0x4b, 0xb0, 0x4c, 0xa6, 0x5e, 0xd4, 0xa7, 0x80,
	0xa6, 0xd3, 0xe2, 0x2d, 0x32, 0x69, 0xfc, 0xf7, 0x01, 0xac, 0xca, 0x32, 0x52, 0x6f, 0xe7, 0x73,
	0x50, 0x9d, 0x9c, 0xdb, 0xa2, 0xca, 0x65, 0x9a, 0x25, 0x76, 0x89, 0x67, 0x4d, 0x78, 0x40, 0xd1,
	0xac, 0x0f, 0x6d, 0xa4, 0x22, 0x0a, 0xdd, 0x6f, 0x28, 0x75, 0x29, 0x9d, 0x2f, 0xc1, 0x7c, 0x19,
	0x6c, 0x94, 0x20, 0x53, 0x3e, 0x5d, 0x97, 0xaa, 0x7e, 0xbf, 0x06, 0x0b, 0x0a, 0x78, 0x0f, 0x36,
	0xa6, 0x09, 0x4c, 0x1f, 0xdf, 0xe6, 0xf7, 0xb1, 0x5e, 0x42, 0x4c, 0x2f, 0x0f, 0xeb, 0x77, 0xf6,
	0x7b, 0xb0, 0xd8, 0xa5, 0xdd, 0x7b, 0x09, 0xdd, 0x54, 0xd4, 0xb6, 0x35, 0x7d, 0xe6, 0x5e, 0xec,
	0x26, 0x42, 0xc5, 0x6f, 0x27, 0xfe, 0x78, 0xe1, 0x97, 0xf0, 0xd5, 0xb2, 0x79, 0x6b, 0xee, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x07, 0x8e, 0xab, 0xaa, 0x4a, 0x05, 0x00, 0x00,
}