// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/do_GetProcedureReturnCodes_Ad.proto
// DO NOT EDIT!

/*
Package do_GetProcedureReturnCodes_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/do_GetProcedureReturnCodes_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package do_GetProcedureReturnCodes_Ad

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
	ProcedureName                  *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
	ProcedureNameNull              bool                        `protobuf:"varint,1001,opt,name=procedure_name_null,json=procedureNameNull" json:"procedure_name_null,omitempty"`
	ProcedureCategoryId            *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=procedure_category_id,json=procedureCategoryId" json:"procedure_category_id,omitempty"`
	ProcedureCategoryIdNull        bool                        `protobuf:"varint,1002,opt,name=procedure_category_id_null,json=procedureCategoryIdNull" json:"procedure_category_id_null,omitempty"`
	ReturnCode                     *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=return_code,json=returnCode" json:"return_code,omitempty"`
	ReturnCodeNull                 bool                        `protobuf:"varint,1003,opt,name=return_code_null,json=returnCodeNull" json:"return_code_null,omitempty"`
	IncludeIndirectReturnCodes     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=include_indirect_return_codes,json=includeIndirectReturnCodes" json:"include_indirect_return_codes,omitempty"`
	IncludeIndirectReturnCodesNull bool                        `protobuf:"varint,1004,opt,name=include_indirect_return_codes_null,json=includeIndirectReturnCodesNull" json:"include_indirect_return_codes_null,omitempty"`
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

func (m *Parameters) GetReturnCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

func (m *Parameters) GetIncludeIndirectReturnCodes() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeIndirectReturnCodes
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
	RowId               int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	ReturnCode          *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=return_code,json=returnCode" json:"return_code,omitempty"`
	Description         *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=description" json:"description,omitempty"`
	ReturnCodeOccurence *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=return_code_occurence,json=returnCodeOccurence" json:"return_code_occurence,omitempty"`
	ProcedureName       *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=procedure_name,json=procedureName" json:"procedure_name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetReturnCode() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

func (m *Response_Row) GetDescription() *dstore_values.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Response_Row) GetReturnCodeOccurence() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnCodeOccurence
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.do_GetProcedureReturnCodes_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.do_GetProcedureReturnCodes_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.do_GetProcedureReturnCodes_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xa6, 0x4d, 0xd3, 0x96, 0x53, 0xac, 0x75, 0x42, 0x75, 0xdd, 0x60, 0x91, 0x78, 0xa3, 0x08,
	0x1b, 0x51, 0x10, 0x85, 0x78, 0xd1, 0x16, 0x91, 0x20, 0xdd, 0x86, 0x05, 0x05, 0xbd, 0x70, 0xd9,
	0xee, 0x1e, 0xc3, 0x40, 0x32, 0x13, 0x66, 0x36, 0x16, 0xdf, 0xc2, 0xbf, 0x57, 0xf0, 0x01, 0x7c,
	0x15, 0xdf, 0xc2, 0x9f, 0x87, 0xf0, 0xec, 0xce, 0x24, 0xfb, 0xd3, 0xd8, 0x1a, 0x6f, 0x12, 0xce,
	0x9e, 0xf3, 0x7d, 0xdf, 0x99, 0x33, 0xdf, 0x1c, 0xe8, 0x25, 0x3a, 0x95, 0x0a, 0xbb, 0x28, 0x86,
	0x5c, 0x60, 0x77, 0xa2, 0x64, 0x8c, 0xc9, 0x54, 0xa1, 0xee, 0x26, 0x32, 0x7c, 0x86, 0xe9, 0x60,
	0xf6, 0x21, 0xc0, 0x74, 0xaa, 0xc4, 0xa1, 0x4c, 0x50, 0x87, 0xfb, 0x89, 0x47, 0x85, 0xa9, 0x64,
	0x77, 0x0d, 0xda, 0x33, 0x68, 0xef, 0x5c, 0x88, 0xdb, 0xb2, 0x52, 0xef, 0xa2, 0xd1, 0x14, 0xb5,
	0x61, 0x70, 0xaf, 0x57, 0xf5, 0x51, 0x29, 0xa9, 0x6c, 0xaa, 0x5d, 0x4d, 0x8d, 0x51, 0xeb, 0x68,
	0x88, 0x36, 0x79, 0xab, 0x9e, 0x4c, 0x23, 0x2e, 0xde, 0x4a, 0x35, 0x8e, 0x52, 0x2e, 0x85, 0x29,
	0xea, 0x7c, 0x5d, 0x03, 0x18, 0x44, 0x2a, 0xa2, 0x2c, 0x2a, 0xcd, 0xf6, 0x61, 0x7b, 0x7e, 0xbe,
	0x50, 0xd0, 0x57, 0x67, 0xe5, 0xe6, 0xca, 0xed, 0xad, 0xfb, 0xae, 0x67, 0x8f, 0x61, 0x3b, 0xd3,
	0xa9, 0xe2, 0x62, 0xf8, 0x32, 0x0b, 0x82, 0x4b, 0x73, 0x84, 0x4f, 0x00, 0xd6, 0x85, 0x56, 0x95,
	0x22, 0x14, 0xd3, 0xd1, 0xc8, 0xf9, 0xb1, 0x41, 0x44, 0x9b, 0xc1, 0x95, 0x4a, 0xb1, 0x4f, 0x19,
	0x76, 0x0c, 0xbb, 0x05, 0x20, 0x8e, 0x52, 0x1c, 0x4a, 0xf5, 0x3e, 0xe4, 0x89, 0xb3, 0x9a, 0x4b,
	0xb7, 0x6b, 0xd2, 0x5c, 0x50, 0x05, 0x2a, 0xa3, 0x5d, 0x48, 0x1d, 0x5a, 0x60, 0x3f, 0x61, 0x3d,
	0x70, 0x17, 0x12, 0x9a, 0x46, 0x7e, 0x9a, 0x46, 0xae, 0x2d, 0x40, 0xe6, 0xed, 0xf4, 0x60, 0x4b,
	0xe5, 0xb7, 0x12, 0xc6, 0x74, 0x2d, 0x4e, 0xe3, 0xe2, 0x26, 0x40, 0xcd, 0x6f, 0x91, 0xdd, 0x81,
	0x9d, 0x12, 0xda, 0x28, 0xfe, 0x32, 0x8a, 0xdb, 0x45, 0x59, 0x2e, 0xf4, 0x06, 0x6e, 0x70, 0x11,
	0x8f, 0xa6, 0x54, 0xc6, 0x45, 0xc2, 0x15, 0xc6, 0x69, 0x58, 0xc2, 0x6a, 0x67, 0x6d, 0xa1, 0xf4,
	0x89, 0x94, 0x23, 0x8c, 0x84, 0x91, 0x76, 0x2d, 0x43, 0xdf, 0x12, 0x94, 0xfc, 0xc4, 0x9e, 0x43,
	0xe7, 0x5c, 0x7e, 0xd3, 0xdc, 0x6f, 0xd3, 0xdc, 0xde, 0xdf, 0x89, 0xb2, 0x66, 0x3b, 0xdf, 0xd7,
	0x60, 0x33, 0x40, 0x3d, 0x91, 0x42, 0x23, 0xbb, 0x07, 0xcd, 0xdc, 0x85, 0x75, 0x73, 0x58, 0x8f,
	0x1b, 0x87, 0x3e, 0xcd, 0x7e, 0x03, 0x53, 0xc8, 0x5e, 0xc1, 0x4e, 0xe6, 0xbf, 0xb0, 0x64, 0x40,
	0xba, 0xde, 0x06, 0x81, 0xbd, 0x1a, 0xb8, 0x6e, 0xd3, 0x23, 0x8a, 0xfb, 0x45, 0x1c, 0x5c, 0x1e,
	0x57, 0x3f, 0xb0, 0x47, 0xb0, 0x61, 0x7d, 0x4f, 0x77, 0x95, 0x31, 0xee, 0x9d, 0x61, 0x34, 0xaf,
	0xe2, 0xc8, 0xfc, 0x07, 0xb3, 0x72, 0x1a, 0x50, 0x43, 0xc9, 0x53, 0x1a, 0x73, 0x86, 0x7a, 0xec,
	0x2d, 0xf1, 0x50, 0xbd, 0xd9, 0x28, 0xbc, 0x40, 0x9e, 0x06, 0x19, 0x8b, 0xfb, 0x6d, 0x15, 0x1a,
	0x14, 0xb0, 0xab, 0xb0, 0x4e, 0x61, 0x66, 0xdf, 0x0f, 0x3e, 0x4d, 0xa7, 0x19, 0x34, 0x29, 0x24,
	0x53, 0x3e, 0xa9, 0xda, 0xea, 0xa3, 0xbf, 0x9c, 0xaf, 0x08, 0x4e, 0xd2, 0xb1, 0xe2, 0x93, 0x7c,
	0x76, 0x9f, 0xfc, 0x0b, 0x9f, 0x65, 0xb9, 0x9e, 0x0d, 0x60, 0xb7, 0x6c, 0x4b, 0x19, 0xc7, 0x74,
	0x26, 0x11, 0xa3, 0xf3, 0xf9, 0x1f, 0xfa, 0x68, 0x15, 0x7d, 0x1c, 0xcf, 0x80, 0xec, 0xe0, 0xcc,
	0xa6, 0xf8, 0xe2, 0x2f, 0xb9, 0x2a, 0x0e, 0x5e, 0x40, 0x9b, 0xcb, 0xda, 0xdc, 0x8b, 0xf5, 0xfa,
	0xfa, 0xe1, 0xff, 0x2d, 0xde, 0x93, 0xf5, 0x7c, 0xb5, 0x3d, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x42, 0xce, 0x99, 0x65, 0xb9, 0x05, 0x00, 0x00,
}
