// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_SearchSourceCode_Ad.proto
// DO NOT EDIT!

/*
Package mi_SearchSourceCode_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_SearchSourceCode_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_SearchSourceCode_Ad

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
	SearchPattern                  *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=search_pattern,json=searchPattern" json:"search_pattern,omitempty"`
	SearchPatternNull              bool                        `protobuf:"varint,1001,opt,name=search_pattern_null,json=searchPatternNull" json:"search_pattern_null,omitempty"`
	CaseSensitive                  *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull              bool                        `protobuf:"varint,1002,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	DoNotSearchInCommentPart       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=do_not_search_in_comment_part,json=doNotSearchInCommentPart" json:"do_not_search_in_comment_part,omitempty"`
	DoNotSearchInCommentPartNull   bool                        `protobuf:"varint,1003,opt,name=do_not_search_in_comment_part_null,json=doNotSearchInCommentPartNull" json:"do_not_search_in_comment_part_null,omitempty"`
	SearchOnlyThisObjectName       *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=search_only_this_object_name,json=searchOnlyThisObjectName" json:"search_only_this_object_name,omitempty"`
	SearchOnlyThisObjectNameNull   bool                        `protobuf:"varint,1004,opt,name=search_only_this_object_name_null,json=searchOnlyThisObjectNameNull" json:"search_only_this_object_name_null,omitempty"`
	SearchOnlyObjectType           *dstore_values.StringValue  `protobuf:"bytes,5,opt,name=search_only_object_type,json=searchOnlyObjectType" json:"search_only_object_type,omitempty"`
	SearchOnlyObjectTypeNull       bool                        `protobuf:"varint,1005,opt,name=search_only_object_type_null,json=searchOnlyObjectTypeNull" json:"search_only_object_type_null,omitempty"`
	GetDistinctObjectNamesOnly     *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=get_distinct_object_names_only,json=getDistinctObjectNamesOnly" json:"get_distinct_object_names_only,omitempty"`
	GetDistinctObjectNamesOnlyNull bool                        `protobuf:"varint,1006,opt,name=get_distinct_object_names_only_null,json=getDistinctObjectNamesOnlyNull" json:"get_distinct_object_names_only_null,omitempty"`
	SearchObjectsWithNameLike      *dstore_values.StringValue  `protobuf:"bytes,7,opt,name=search_objects_with_name_like,json=searchObjectsWithNameLike" json:"search_objects_with_name_like,omitempty"`
	SearchObjectsWithNameLikeNull  bool                        `protobuf:"varint,1007,opt,name=search_objects_with_name_like_null,json=searchObjectsWithNameLikeNull" json:"search_objects_with_name_like_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSearchPattern() *dstore_values.StringValue {
	if m != nil {
		return m.SearchPattern
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

func (m *Parameters) GetDoNotSearchInCommentPart() *dstore_values.IntegerValue {
	if m != nil {
		return m.DoNotSearchInCommentPart
	}
	return nil
}

func (m *Parameters) GetSearchOnlyThisObjectName() *dstore_values.StringValue {
	if m != nil {
		return m.SearchOnlyThisObjectName
	}
	return nil
}

func (m *Parameters) GetSearchOnlyObjectType() *dstore_values.StringValue {
	if m != nil {
		return m.SearchOnlyObjectType
	}
	return nil
}

func (m *Parameters) GetGetDistinctObjectNamesOnly() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetDistinctObjectNamesOnly
	}
	return nil
}

func (m *Parameters) GetSearchObjectsWithNameLike() *dstore_values.StringValue {
	if m != nil {
		return m.SearchObjectsWithNameLike
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
	ObjectType     *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=object_type,json=objectType" json:"object_type,omitempty"`
	CodeLineNumber *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=code_line_number,json=codeLineNumber" json:"code_line_number,omitempty"`
	CodeLine       *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=code_line,json=codeLine" json:"code_line,omitempty"`
	ObjectName     *dstore_values.StringValue  `protobuf:"bytes,10004,opt,name=object_name,json=objectName" json:"object_name,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetObjectType() *dstore_values.StringValue {
	if m != nil {
		return m.ObjectType
	}
	return nil
}

func (m *Response_Row) GetCodeLineNumber() *dstore_values.IntegerValue {
	if m != nil {
		return m.CodeLineNumber
	}
	return nil
}

func (m *Response_Row) GetCodeLine() *dstore_values.StringValue {
	if m != nil {
		return m.CodeLine
	}
	return nil
}

func (m *Response_Row) GetObjectName() *dstore_values.StringValue {
	if m != nil {
		return m.ObjectName
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_SearchSourceCode_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_SearchSourceCode_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_SearchSourceCode_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_SearchSourceCode_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x95, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0xd5, 0xa6, 0x69, 0x83, 0x2b, 0x4a, 0x71, 0x11, 0xb8, 0xe9, 0x45, 0x90, 0xb2, 0x60,
	0x35, 0x41, 0x45, 0x48, 0x5d, 0x20, 0xa1, 0x5e, 0x40, 0x04, 0xb5, 0x69, 0x99, 0x56, 0x54, 0x54,
	0x48, 0xd6, 0x24, 0x63, 0x52, 0xd3, 0xc4, 0x8e, 0x6c, 0xa7, 0x55, 0xdf, 0x82, 0xdb, 0x8e, 0xa7,
	0x60, 0xc1, 0x43, 0x71, 0x7f, 0x05, 0x8e, 0xc7, 0x0e, 0xc9, 0x84, 0x26, 0xd3, 0x4d, 0x22, 0xcf,
	0x39, 0xdf, 0x7f, 0x7e, 0x1f, 0x8f, 0xcf, 0xa0, 0x87, 0xb1, 0x36, 0x52, 0xb1, 0x32, 0x13, 0x0d,
	0x2e, 0x58, 0xb9, 0xad, 0x64, 0x9d, 0xc5, 0x1d, 0xc5, 0x74, 0xb9, 0xc5, 0xe9, 0x3e, 0x8b, 0x54,
	0xfd, 0x78, 0x5f, 0x76, 0x54, 0x9d, 0x6d, 0xca, 0x98, 0xd1, 0xf5, 0x38, 0x80, 0x0c, 0x23, 0xf1,
	0x5d, 0x87, 0x05, 0x0e, 0x0b, 0x2e, 0xce, 0x2d, 0xce, 0x79, 0xf1, 0xd3, 0xa8, 0xd9, 0x61, 0xda,
	0xa1, 0xc5, 0xf9, 0x74, 0x45, 0xa6, 0x94, 0x54, 0x3e, 0xb4, 0x90, 0x0e, 0xb5, 0x98, 0xd6, 0x51,
	0x83, 0xf9, 0xe0, 0xca, 0x60, 0xd0, 0x44, 0x5c, 0xbc, 0x91, 0xaa, 0x15, 0x19, 0x2e, 0x85, 0x4b,
	0x2a, 0x7d, 0x29, 0x20, 0xb4, 0x17, 0xa9, 0x08, 0xa2, 0x4c, 0x69, 0xbc, 0x8e, 0x66, 0x74, 0xe2,
	0x8b, 0xb6, 0x23, 0x03, 0x4f, 0x04, 0x19, 0xbb, 0x3d, 0x76, 0x6f, 0x7a, 0xb5, 0x18, 0x78, 0xff,
	0xde, 0x99, 0x36, 0x8a, 0x8b, 0xc6, 0x4b, 0xbb, 0x08, 0xaf, 0x3a, 0x62, 0xcf, 0x01, 0xb8, 0x8c,
	0xe6, 0xd2, 0x12, 0x54, 0x74, 0x9a, 0x4d, 0xf2, 0x6d, 0x0a, 0x84, 0x0a, 0xe1, 0xf5, 0x54, 0x72,
	0x15, 0x22, 0x78, 0x03, 0xcd, 0xd4, 0x23, 0xcd, 0xa8, 0x66, 0x42, 0x73, 0xc3, 0x4f, 0x19, 0x19,
	0x4f, 0x6a, 0x2e, 0x0c, 0xd4, 0xac, 0x49, 0xd9, 0x64, 0x91, 0xf0, 0x45, 0x2d, 0xb2, 0xdf, 0x25,
	0x6c, 0xd1, 0xb4, 0x86, 0x2b, 0xfa, 0xdd, 0x17, 0x4d, 0x25, 0x27, 0x45, 0x5f, 0xa3, 0xa5, 0x58,
	0x52, 0x21, 0x0d, 0xf5, 0x66, 0xb9, 0xa0, 0x75, 0xd9, 0x6a, 0x31, 0x61, 0xc0, 0xb7, 0x32, 0x24,
	0x77, 0xa1, 0x07, 0x2e, 0x0c, 0x6b, 0x30, 0xe5, 0x3c, 0x90, 0x58, 0x56, 0xa5, 0x71, 0xe7, 0x58,
	0x11, 0x9b, 0x8e, 0x86, 0x66, 0x1a, 0x5c, 0x41, 0xa5, 0x91, 0xea, 0xce, 0xdd, 0x0f, 0xe7, 0x6e,
	0x71, 0x98, 0x4c, 0x62, 0xf4, 0x08, 0x2d, 0x7a, 0x0d, 0x29, 0x9a, 0xe7, 0xd4, 0x1c, 0x73, 0x4d,
	0x65, 0xed, 0x2d, 0xab, 0x83, 0x08, 0x9c, 0x1a, 0x99, 0xc8, 0x3c, 0x1f, 0xe2, 0xf8, 0x5d, 0xc0,
	0x0f, 0x80, 0xde, 0x4d, 0xe0, 0x2a, 0xb0, 0xf8, 0x19, 0xba, 0x33, 0x4a, 0xdb, 0xb9, 0xfc, 0xe9,
	0x5d, 0x0e, 0x53, 0x49, 0x5c, 0xbe, 0x40, 0xb7, 0xfa, 0x95, 0xbc, 0x88, 0x39, 0x6f, 0x33, 0x92,
	0xcf, 0x34, 0x78, 0xa3, 0x27, 0xed, 0x64, 0x0f, 0x80, 0xc3, 0x8f, 0xd3, 0x1b, 0xef, 0x93, 0x74,
	0xbe, 0x7e, 0x39, 0x5f, 0xe4, 0x22, 0x38, 0xf1, 0x44, 0xd1, 0x72, 0x83, 0x19, 0x1a, 0x73, 0x6d,
	0xb8, 0x00, 0xb4, 0x6f, 0x67, 0x3a, 0xd1, 0x24, 0x93, 0xd9, 0xef, 0x59, 0x11, 0x24, 0xb6, 0xbc,
	0x42, 0x6f, 0xcf, 0xda, 0x96, 0xc3, 0xdb, 0x68, 0x65, 0x74, 0x01, 0x67, 0xf4, 0xb7, 0x33, 0xba,
	0x3c, 0x5c, 0xa9, 0xfb, 0x46, 0x76, 0xf7, 0x9b, 0x44, 0x35, 0x3d, 0xe3, 0xe6, 0xd8, 0x1d, 0x44,
	0x93, 0x9f, 0x30, 0x32, 0x95, 0xd9, 0xc8, 0x79, 0xdf, 0x0b, 0xc7, 0x1f, 0x02, 0x6e, 0x4b, 0x6c,
	0x03, 0x8c, 0x9f, 0xa3, 0xd2, 0x48, 0x75, 0x67, 0xf5, 0x8f, 0xb3, 0xba, 0x34, 0x54, 0xc7, 0x3a,
	0x2d, 0x7d, 0x9d, 0x40, 0x85, 0x90, 0xe9, 0xb6, 0x14, 0x9a, 0xe1, 0xfb, 0x28, 0x9f, 0x4c, 0xa4,
	0xc1, 0x41, 0xe1, 0x07, 0x9d, 0x9b, 0x56, 0x4f, 0xec, 0x6f, 0xe8, 0x12, 0xf1, 0x2b, 0x34, 0x6b,
	0x67, 0x11, 0xed, 0x1b, 0x46, 0x70, 0xe3, 0x73, 0x00, 0x07, 0x03, 0xf0, 0xe0, 0xc8, 0xda, 0x81,
	0x75, 0xa5, 0xb7, 0x0e, 0xaf, 0xb5, 0xd2, 0x0f, 0xf0, 0x1a, 0x9a, 0xf2, 0x33, 0x10, 0xee, 0xaf,
	0x55, 0x5c, 0xfe, 0x4f, 0xd1, 0x4d, 0xc8, 0x1d, 0xf7, 0x1f, 0x76, 0xd3, 0xf1, 0x16, 0xca, 0x29,
	0x79, 0x06, 0xb7, 0xc9, 0x52, 0xab, 0xc1, 0x65, 0xa6, 0x75, 0xd0, 0xed, 0x41, 0x10, 0xca, 0xb3,
	0xd0, 0xe2, 0xc5, 0xcf, 0xe3, 0x28, 0x07, 0x0b, 0x7c, 0x13, 0x4d, 0xc2, 0x92, 0xf2, 0x98, 0xbc,
	0xab, 0x42, 0x5b, 0xf2, 0x61, 0x1e, 0x96, 0x95, 0x18, 0x3f, 0x42, 0xd3, 0xfd, 0x57, 0xe3, 0x7d,
	0x35, 0xf3, 0x48, 0x91, 0xec, 0xdd, 0x88, 0xa7, 0x68, 0xb6, 0x6e, 0x4b, 0x37, 0xc1, 0x13, 0x9c,
	0x57, 0xab, 0xc6, 0x14, 0xf9, 0x50, 0xcd, 0x9e, 0x53, 0x33, 0x96, 0xda, 0x06, 0xa8, 0x9a, 0x30,
	0xd0, 0xa5, 0x2b, 0xff, 0x74, 0xc8, 0xc7, 0x6c, 0x0f, 0x85, 0x2e, 0xdf, 0xe7, 0x3f, 0x99, 0x3d,
	0x9f, 0x2e, 0xed, 0xdf, 0xbe, 0x3f, 0x1b, 0x87, 0x68, 0x81, 0xcb, 0x81, 0xd6, 0xf6, 0xbe, 0x9f,
	0x47, 0x6b, 0x0d, 0xa9, 0xe3, 0x93, 0x6e, 0x3c, 0xbe, 0xfc, 0x27, 0xb6, 0x36, 0x99, 0x7c, 0xcb,
	0x1e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x93, 0x4e, 0xc2, 0x52, 0x9c, 0x07, 0x00, 0x00,
}
