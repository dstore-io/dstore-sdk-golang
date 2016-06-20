// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetApplicationSettings_Ad.proto
// DO NOT EDIT!

/*
Package mi_GetApplicationSettings_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetApplicationSettings_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetApplicationSettings_Ad

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
	ApplicationId                    *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationIdNull                bool                        `protobuf:"varint,1001,opt,name=application_id_null,json=applicationIdNull" json:"application_id_null,omitempty"`
	UserId                           *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIdNull                       bool                        `protobuf:"varint,1002,opt,name=user_id_null,json=userIdNull" json:"user_id_null,omitempty"`
	KeyVariable                      *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	KeyVariableNull                  bool                        `protobuf:"varint,1003,opt,name=key_variable_null,json=keyVariableNull" json:"key_variable_null,omitempty"`
	SearchForKeyVariableWithLike     *dstore_values.BooleanValue `protobuf:"bytes,4,opt,name=search_for_key_variable_with_like,json=searchForKeyVariableWithLike" json:"search_for_key_variable_with_like,omitempty"`
	SearchForKeyVariableWithLikeNull bool                        `protobuf:"varint,1004,opt,name=search_for_key_variable_with_like_null,json=searchForKeyVariableWithLikeNull" json:"search_for_key_variable_with_like_null,omitempty"`
	IncludeValuesForGlobalUser       *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=include_values_for_global_user,json=includeValuesForGlobalUser" json:"include_values_for_global_user,omitempty"`
	IncludeValuesForGlobalUserNull   bool                        `protobuf:"varint,1005,opt,name=include_values_for_global_user_null,json=includeValuesForGlobalUserNull" json:"include_values_for_global_user_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplicationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *Parameters) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Parameters) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Parameters) GetSearchForKeyVariableWithLike() *dstore_values.BooleanValue {
	if m != nil {
		return m.SearchForKeyVariableWithLike
	}
	return nil
}

func (m *Parameters) GetIncludeValuesForGlobalUser() *dstore_values.BooleanValue {
	if m != nil {
		return m.IncludeValuesForGlobalUser
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
	RowId                      int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value                      *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value" json:"value,omitempty"`
	KeyVariable                *dstore_values.StringValue  `protobuf:"bytes,10002,opt,name=key_variable,json=keyVariable" json:"key_variable,omitempty"`
	ApplicationId              *dstore_values.IntegerValue `protobuf:"bytes,10003,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ValueDerivedFromGlobalUser *dstore_values.BooleanValue `protobuf:"bytes,10004,opt,name=value_derived_from_global_user,json=valueDerivedFromGlobalUser" json:"value_derived_from_global_user,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue() *dstore_values.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response_Row) GetKeyVariable() *dstore_values.StringValue {
	if m != nil {
		return m.KeyVariable
	}
	return nil
}

func (m *Response_Row) GetApplicationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *Response_Row) GetValueDerivedFromGlobalUser() *dstore_values.BooleanValue {
	if m != nil {
		return m.ValueDerivedFromGlobalUser
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetApplicationSettings_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetApplicationSettings_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetApplicationSettings_Ad.Response.Row")
}

func init() {
	proto.RegisterFile("dstore/engine/procedures/mi_GetApplicationSettings_Ad.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x4e, 0x13, 0x4f,
	0x14, 0x0f, 0x94, 0x05, 0x72, 0xe0, 0xff, 0x47, 0x86, 0xc4, 0xac, 0xc5, 0x34, 0x5c, 0x12, 0x63,
	0x34, 0xd9, 0x7a, 0xfb, 0x40, 0x62, 0x24, 0x01, 0x15, 0x82, 0xc2, 0x46, 0xd7, 0x88, 0xd1, 0xc4,
	0x8c, 0xd3, 0xee, 0x61, 0x99, 0xb0, 0xdd, 0x69, 0x66, 0xb6, 0x34, 0xbe, 0x85, 0xb7, 0xe7, 0xf0,
	0xab, 0x8f, 0xe2, 0x33, 0x78, 0x7b, 0x07, 0x67, 0x67, 0xb6, 0xb6, 0xbb, 0x68, 0x4b, 0xbf, 0x40,
	0xce, 0x9e, 0xf3, 0xbb, 0xec, 0xec, 0xf9, 0x4d, 0xe1, 0x6e, 0xa8, 0x52, 0x21, 0xb1, 0x8e, 0x49,
	0xc4, 0x13, 0xac, 0xb7, 0xa5, 0x68, 0x62, 0xd8, 0x91, 0xa8, 0xea, 0x2d, 0x4e, 0x77, 0x31, 0xdd,
	0x6a, 0xb7, 0x63, 0xde, 0x64, 0x29, 0x17, 0xc9, 0x33, 0x4c, 0x53, 0x9e, 0x44, 0x8a, 0x6e, 0x85,
	0x9e, 0x9e, 0x4b, 0x05, 0xb9, 0x66, 0xc1, 0x9e, 0x05, 0x7b, 0xc3, 0x10, 0xd5, 0xa5, 0x5c, 0xe8,
	0x94, 0xc5, 0x1d, 0x54, 0x96, 0xa0, 0x7a, 0xa9, 0xa8, 0x8e, 0x52, 0x0a, 0x99, 0xb7, 0x96, 0x8b,
	0xad, 0x16, 0x2a, 0xc5, 0x22, 0xcc, 0x9b, 0xeb, 0xe5, 0x66, 0xca, 0x78, 0x72, 0x24, 0x64, 0xcb,
	0xc8, 0xda, 0xa1, 0xb5, 0xcf, 0x0e, 0xc0, 0x13, 0x26, 0x99, 0xee, 0xa2, 0x54, 0x64, 0x1b, 0xfe,
	0x67, 0x7d, 0x6b, 0x94, 0x87, 0xee, 0xc4, 0xca, 0xc4, 0xd5, 0xb9, 0x5b, 0xcb, 0x5e, 0xfe, 0x16,
	0xb9, 0x33, 0x9e, 0xa4, 0x18, 0xa1, 0x3c, 0xcc, 0xaa, 0xe0, 0xbf, 0x01, 0xc8, 0x5e, 0x48, 0xea,
	0xb0, 0x54, 0xe4, 0xa0, 0x49, 0x27, 0x8e, 0xdd, 0x6f, 0x33, 0x9a, 0x69, 0x36, 0x58, 0x2c, 0x0c,
	0xfb, 0xba, 0x43, 0xee, 0xc0, 0x4c, 0x47, 0xa1, 0xcc, 0xd4, 0x26, 0x47, 0xab, 0x4d, 0x67, 0xb3,
	0x5a, 0x66, 0x15, 0xe6, 0x73, 0x94, 0xe5, 0xff, 0x6e, 0xf9, 0xc1, 0xb6, 0x0d, 0xf1, 0x3d, 0x98,
	0x3f, 0xc1, 0xb7, 0xf4, 0x94, 0x49, 0xce, 0x1a, 0x31, 0xba, 0x15, 0xc3, 0x5e, 0x2d, 0xb1, 0xab,
	0x54, 0xea, 0xf3, 0xb7, 0xe4, 0x73, 0x7a, 0xfe, 0x30, 0x1f, 0x27, 0xd7, 0x61, 0x71, 0x10, 0x6e,
	0x65, 0x7e, 0x58, 0x99, 0x85, 0x81, 0x41, 0xa3, 0x15, 0xc2, 0xaa, 0x42, 0x26, 0x9b, 0xc7, 0x54,
	0x1f, 0x31, 0x2d, 0xe0, 0xba, 0x3c, 0x3d, 0xa6, 0x31, 0x3f, 0x41, 0x77, 0xea, 0xaf, 0xaf, 0xd7,
	0x10, 0x22, 0x46, 0x96, 0x58, 0x07, 0x97, 0x2d, 0xcb, 0x8e, 0x90, 0x8f, 0xfb, 0x0a, 0x2f, 0x34,
	0xc3, 0xbe, 0x26, 0x20, 0x4f, 0xe1, 0xca, 0x48, 0x15, 0xeb, 0xf3, 0xa7, 0xf5, 0xb9, 0x32, 0x8c,
	0xce, 0x18, 0xa7, 0x50, 0xe3, 0x49, 0x33, 0xee, 0x84, 0x48, 0xad, 0x1f, 0x43, 0x1d, 0xc5, 0xa2,
	0xc1, 0x62, 0x9a, 0x1d, 0xa6, 0xeb, 0x8c, 0x76, 0x5d, 0xcd, 0x29, 0x4c, 0xa5, 0xb4, 0xda, 0xae,
	0xc1, 0x3f, 0xd7, 0x70, 0xb2, 0x0f, 0xeb, 0xc3, 0x05, 0xac, 0xe1, 0x5f, 0xd6, 0x70, 0xed, 0xdf,
	0x4c, 0x99, 0xdd, 0xb5, 0xaf, 0x53, 0x30, 0x1b, 0xa0, 0x6a, 0x8b, 0x44, 0x21, 0xb9, 0x01, 0x8e,
	0x89, 0x43, 0xbe, 0xa5, 0x7f, 0xbe, 0x6c, 0x9e, 0x35, 0x1b, 0x95, 0x87, 0xd9, 0xdf, 0xc0, 0x0e,
	0x92, 0x97, 0x70, 0x21, 0x0b, 0x02, 0x1d, 0x48, 0x82, 0x5e, 0xba, 0x8a, 0x06, 0x7b, 0x25, 0x70,
	0x39, 0x2f, 0x07, 0xba, 0xde, 0xeb, 0xd7, 0xc1, 0x42, 0xab, 0xf8, 0x80, 0x6c, 0xc0, 0x4c, 0x1e,
	0x40, 0xbd, 0x68, 0x19, 0x63, 0xed, 0x0c, 0xa3, 0x8d, 0xe7, 0x81, 0xfd, 0x1f, 0xf4, 0xc6, 0xc9,
	0x23, 0xa8, 0x48, 0xd1, 0xd5, 0xdb, 0x91, 0xa1, 0x36, 0xbc, 0xf3, 0x5f, 0x18, 0x5e, 0xef, 0x24,
	0xbc, 0x40, 0x74, 0x83, 0x8c, 0xa4, 0xfa, 0x65, 0x12, 0x2a, 0xba, 0x20, 0x17, 0x61, 0x5a, 0x97,
	0x59, 0xa6, 0xde, 0xf9, 0xfa, 0x70, 0x9c, 0xc0, 0xd1, 0xa5, 0x8e, 0xcd, 0x4d, 0x70, 0xcc, 0x57,
	0x70, 0xdf, 0xfb, 0x23, 0xd3, 0x60, 0x27, 0xc9, 0x66, 0x29, 0x46, 0x1f, 0xfc, 0xf1, 0x72, 0x74,
	0xff, 0xcc, 0xa5, 0xf2, 0xd1, 0x1f, 0xfb, 0x56, 0x79, 0x03, 0x35, 0x33, 0x45, 0x43, 0x94, 0xfc,
	0x14, 0x43, 0x7a, 0x24, 0x45, 0xab, 0xb0, 0xa6, 0x9f, 0xfc, 0x73, 0xec, 0xa9, 0x79, 0xf8, 0xc0,
	0x52, 0xec, 0x68, 0x86, 0xfe, 0x76, 0x6d, 0xbf, 0x86, 0x65, 0x2e, 0x4a, 0x87, 0xdf, 0xbf, 0xea,
	0x5f, 0x6d, 0x46, 0x42, 0x85, 0x27, 0xbd, 0x7e, 0x38, 0xee, 0xaf, 0x41, 0x63, 0xda, 0x5c, 0xb8,
	0xb7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x91, 0x12, 0xcc, 0x8f, 0x4d, 0x06, 0x00, 0x00,
}