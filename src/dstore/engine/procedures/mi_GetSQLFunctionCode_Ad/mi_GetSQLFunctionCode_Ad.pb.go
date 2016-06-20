// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetSQLFunctionCode_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetSQLFunctionCode_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetSQLFunctionCode_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetSQLFunctionCode_Ad

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
	SQLFunctionName                  *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=s_q_l_function_name,json=sQLFunctionName" json:"s_q_l_function_name,omitempty"`
	SQLFunctionNameNull              bool                        `protobuf:"varint,1001,opt,name=s_q_l_function_name_null,json=sQLFunctionNameNull" json:"s_q_l_function_name_null,omitempty"`
	IgnoreComments                   *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=ignore_comments,json=ignoreComments" json:"ignore_comments,omitempty"`
	IgnoreCommentsNull               bool                        `protobuf:"varint,1002,opt,name=ignore_comments_null,json=ignoreCommentsNull" json:"ignore_comments_null,omitempty"`
	GetCodeLinesAsResultSet          *dstore_values.BooleanValue `protobuf:"bytes,3,opt,name=get_code_lines_as_result_set,json=getCodeLinesAsResultSet" json:"get_code_lines_as_result_set,omitempty"`
	GetCodeLinesAsResultSetNull      bool                        `protobuf:"varint,1003,opt,name=get_code_lines_as_result_set_null,json=getCodeLinesAsResultSetNull" json:"get_code_lines_as_result_set_null,omitempty"`
	IncludeStatementsForCreation     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=include_statements_for_creation,json=includeStatementsForCreation" json:"include_statements_for_creation,omitempty"`
	IncludeStatementsForCreationNull bool                        `protobuf:"varint,1004,opt,name=include_statements_for_creation_null,json=includeStatementsForCreationNull" json:"include_statements_for_creation_null,omitempty"`
	OnlyFunctionHeader               *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=only_function_header,json=onlyFunctionHeader" json:"only_function_header,omitempty"`
	OnlyFunctionHeaderNull           bool                        `protobuf:"varint,1005,opt,name=only_function_header_null,json=onlyFunctionHeaderNull" json:"only_function_header_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSQLFunctionName() *dstore_values.StringValue {
	if m != nil {
		return m.SQLFunctionName
	}
	return nil
}

func (m *Parameters) GetIgnoreComments() *dstore_values.BooleanValue {
	if m != nil {
		return m.IgnoreComments
	}
	return nil
}

func (m *Parameters) GetGetCodeLinesAsResultSet() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetCodeLinesAsResultSet
	}
	return nil
}

func (m *Parameters) GetIncludeStatementsForCreation() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeStatementsForCreation
	}
	return nil
}

func (m *Parameters) GetOnlyFunctionHeader() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlyFunctionHeader
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
	RowId              int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	CodeLineNumber     *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=code_line_number,json=codeLineNumber" json:"code_line_number,omitempty"`
	CodePartLineNumber *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=code_part_line_number,json=codePartLineNumber" json:"code_part_line_number,omitempty"`
	CodeLine           *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=code_line,json=codeLine" json:"code_line,omitempty"`
	CommentPart        *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=comment_part,json=commentPart" json:"comment_part,omitempty"`
	CodePart           *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=code_part,json=codePart" json:"code_part,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetCodeLineNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.CodeLineNumber
	}
	return nil
}

func (m *Response_Row) GetCodePartLineNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.CodePartLineNumber
	}
	return nil
}

func (m *Response_Row) GetCodeLine() *dstore_values.StringValue {
	if m != nil {
		return m.CodeLine
	}
	return nil
}

func (m *Response_Row) GetCommentPart() *dstore_values.StringValue {
	if m != nil {
		return m.CommentPart
	}
	return nil
}

func (m *Response_Row) GetCodePart() *dstore_values.StringValue {
	if m != nil {
		return m.CodePart
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetSQLFunctionCode_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetSQLFunctionCode_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetSQLFunctionCode_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x4f, 0x13, 0x41,
	0x14, 0x0e, 0xd4, 0x02, 0x0e, 0x46, 0xc8, 0x80, 0xb8, 0xb4, 0x44, 0x11, 0x8d, 0xf1, 0x69, 0x11,
	0x35, 0x4a, 0x7c, 0x30, 0x41, 0x94, 0x4b, 0x02, 0x15, 0x97, 0xc4, 0x44, 0x5e, 0x26, 0xdb, 0xee,
	0xa1, 0x6e, 0xb2, 0x3b, 0x83, 0x33, 0x53, 0x89, 0x7f, 0xc0, 0x67, 0xaf, 0x3f, 0x52, 0xc4, 0xff,
	0xe0, 0x99, 0x4b, 0x5b, 0xba, 0x5c, 0x16, 0x5f, 0x20, 0xa7, 0xe7, 0x7c, 0x97, 0x99, 0x73, 0xce,
	0x2c, 0x79, 0x96, 0x28, 0x2d, 0x24, 0x2c, 0x02, 0x6f, 0xa7, 0x1c, 0x16, 0x0f, 0xa4, 0x68, 0x41,
	0xd2, 0x91, 0xa0, 0x16, 0xf3, 0x94, 0xad, 0x83, 0xde, 0x7d, 0xbb, 0xb5, 0xd6, 0xe1, 0x2d, 0x9d,
	0x0a, 0xbe, 0x2a, 0x12, 0x60, 0x2b, 0x49, 0x88, 0x35, 0x5a, 0xd0, 0xfb, 0x0e, 0x18, 0x3a, 0x60,
	0x78, 0x5e, 0x75, 0x6d, 0xca, 0x0b, 0x7c, 0x8a, 0xb3, 0x0e, 0x28, 0x07, 0xae, 0xcd, 0x0e, 0xaa,
	0x82, 0x94, 0x42, 0xfa, 0x54, 0x7d, 0x30, 0x95, 0x83, 0x52, 0x71, 0x1b, 0x7c, 0xf2, 0x6e, 0x31,
	0xa9, 0xe3, 0x94, 0xef, 0x0b, 0x99, 0xc7, 0x46, 0xd2, 0x15, 0x2d, 0x1c, 0x57, 0x09, 0xd9, 0x89,
	0x65, 0x8c, 0x59, 0x90, 0x8a, 0x6e, 0x90, 0x29, 0xc5, 0x3e, 0xb2, 0x8c, 0xed, 0x7b, 0x67, 0x8c,
	0x63, 0x2a, 0x18, 0x9a, 0x1f, 0x7a, 0x30, 0xfe, 0xa8, 0x16, 0xfa, 0x63, 0x78, 0x7b, 0x4a, 0xcb,
	0x94, 0xb7, 0xdf, 0x99, 0x20, 0x9a, 0x50, 0xfd, 0xd3, 0x34, 0x10, 0x42, 0x9f, 0x92, 0xe0, 0x0c,
	0x26, 0xc6, 0x3b, 0x59, 0x16, 0xfc, 0x1e, 0x45, 0xbe, 0xb1, 0x68, 0xaa, 0x80, 0x69, 0x60, 0x8e,
	0xbe, 0x22, 0x13, 0x69, 0x9b, 0xa3, 0x0a, 0x6b, 0x89, 0x3c, 0x07, 0xae, 0x55, 0x30, 0x6c, 0xd5,
	0xeb, 0x05, 0xf5, 0xa6, 0x10, 0x19, 0xc4, 0xdc, 0xc9, 0x5f, 0x77, 0x98, 0x55, 0x0f, 0xa1, 0x4b,
	0x64, 0xba, 0xc0, 0xe2, 0x94, 0x8f, 0x9c, 0x32, 0x1d, 0x2c, 0xb7, 0xc2, 0x7b, 0x64, 0xae, 0x0d,
	0x1a, 0xeb, 0xb1, 0x15, 0x19, 0x5e, 0x98, 0x62, 0xb1, 0x62, 0xd8, 0xdb, 0x4e, 0xa6, 0x99, 0x02,
	0x1d, 0x54, 0xca, 0x5d, 0xdc, 0x44, 0x02, 0xd3, 0xca, 0x2d, 0x03, 0x5f, 0x51, 0x91, 0x05, 0xef,
	0x82, 0xa6, 0xeb, 0xe4, 0xce, 0x45, 0xdc, 0xce, 0xdb, 0x1f, 0xe7, 0xad, 0x7e, 0x0e, 0x89, 0x35,
	0xd9, 0x24, 0xb7, 0x53, 0xde, 0xca, 0x3a, 0xc8, 0xa3, 0x74, 0xac, 0xc1, 0x1d, 0x0d, 0x9b, 0xca,
	0x5a, 0x12, 0x6c, 0x5f, 0x83, 0x2b, 0xe5, 0x3e, 0xe7, 0x3c, 0xc7, 0x6e, 0x8f, 0x62, 0x4d, 0xc8,
	0x55, 0x4f, 0x40, 0xdf, 0x90, 0x7b, 0x25, 0x1a, 0xce, 0xef, 0xb1, 0xf3, 0x3b, 0x7f, 0x11, 0x99,
	0x35, 0xbd, 0x4d, 0xa6, 0x05, 0xcf, 0x3e, 0xf7, 0x27, 0xe1, 0x03, 0xc4, 0x09, 0xc8, 0xa0, 0x5a,
	0xee, 0x94, 0x1a, 0x60, 0x77, 0x46, 0x36, 0x2c, 0x8c, 0x3e, 0x27, 0xb3, 0x67, 0xd1, 0x39, 0x53,
	0x7f, 0x9d, 0xa9, 0x99, 0xd3, 0x38, 0x63, 0x65, 0xe1, 0x4b, 0x95, 0x8c, 0xe1, 0x8d, 0x1e, 0x08,
	0xae, 0x80, 0x3e, 0x24, 0x55, 0xbb, 0x4c, 0xc5, 0xf1, 0xf6, 0x5b, 0xea, 0x16, 0xed, 0xb5, 0xf9,
	0x1b, 0xb9, 0x42, 0xfa, 0x9e, 0x4c, 0x9a, 0x35, 0x62, 0x27, 0xf6, 0x08, 0xa7, 0xb3, 0x82, 0xe0,
	0xb0, 0x00, 0x2e, 0x6e, 0xdb, 0x36, 0xc6, 0x9b, 0xfd, 0x38, 0x9a, 0xc8, 0x07, 0x7f, 0xa0, 0xcb,
	0x64, 0xd4, 0xaf, 0x2f, 0x4e, 0x9a, 0x61, 0xbc, 0x75, 0x8a, 0xd1, 0x2d, 0xf7, 0xb6, 0xfb, 0x1f,
	0x75, 0xcb, 0xe9, 0x1a, 0xa9, 0x48, 0x71, 0x88, 0x7d, 0x37, 0xa8, 0x27, 0xe1, 0xe5, 0x9e, 0x9a,
	0xb0, 0x7b, 0x0b, 0x61, 0x24, 0x0e, 0x23, 0x43, 0x50, 0x3b, 0x1a, 0x26, 0x15, 0x0c, 0xe8, 0x0c,
	0x19, 0xc1, 0x90, 0xa5, 0x49, 0xf0, 0xb5, 0x81, 0x17, 0x53, 0x8d, 0xaa, 0x18, 0x6e, 0x26, 0xa8,
	0x33, 0xd9, 0x1b, 0x60, 0xbc, 0xec, 0xbc, 0x89, 0x2d, 0xfc, 0xd6, 0x38, 0xb3, 0x87, 0x29, 0xd7,
	0xd0, 0x06, 0xe9, 0x77, 0xb3, 0xe5, 0xa7, 0xb9, 0x61, 0x31, 0x38, 0x5f, 0x37, 0x2c, 0xcf, 0x41,
	0x2c, 0xf5, 0x00, 0xd9, 0xf7, 0x4b, 0x90, 0x51, 0x03, 0xc5, 0x07, 0x4b, 0x9f, 0x20, 0x5c, 0x26,
	0x57, 0x7b, 0xc6, 0x82, 0x1f, 0x8d, 0xd2, 0xb7, 0x6a, 0xac, 0x6b, 0x88, 0xbe, 0x20, 0xd7, 0xfc,
	0xfb, 0x60, 0xdd, 0x04, 0x3f, 0xcb, 0xc1, 0xe3, 0x1e, 0x60, 0x3c, 0xf4, 0x94, 0x2d, 0xf8, 0xd7,
	0x25, 0x95, 0x0d, 0xf2, 0xe5, 0x0e, 0xa9, 0xa7, 0xa2, 0xd0, 0xab, 0xfe, 0xf7, 0x64, 0x6f, 0xe9,
	0xbf, 0xbf, 0x34, 0xcd, 0x11, 0xfb, 0xa0, 0x3f, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xc6,
	0x10, 0x2f, 0xa5, 0x06, 0x00, 0x00,
}
