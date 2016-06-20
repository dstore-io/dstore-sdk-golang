// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/om_GetPredefinedBillContent_Ad.proto
// DO NOT EDIT!

/*
Package om_GetPredefinedBillContent_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/om_GetPredefinedBillContent_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package om_GetPredefinedBillContent_Ad

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
	PredefinedBillContentId     *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=predefined_bill_content_id,json=predefinedBillContentId" json:"predefined_bill_content_id,omitempty"`
	PredefinedBillContentIdNull bool                        `protobuf:"varint,1001,opt,name=predefined_bill_content_id_null,json=predefinedBillContentIdNull" json:"predefined_bill_content_id_null,omitempty"`
	TableId                     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	TableIdNull                 bool                        `protobuf:"varint,1002,opt,name=table_id_null,json=tableIdNull" json:"table_id_null,omitempty"`
	TableKeyId                  *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=table_key_id,json=tableKeyId" json:"table_key_id,omitempty"`
	TableKeyIdNull              bool                        `protobuf:"varint,1003,opt,name=table_key_id_null,json=tableKeyIdNull" json:"table_key_id_null,omitempty"`
	Active                      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=active" json:"active,omitempty"`
	ActiveNull                  bool                        `protobuf:"varint,1004,opt,name=active_null,json=activeNull" json:"active_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPredefinedBillContentId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredefinedBillContentId
	}
	return nil
}

func (m *Parameters) GetTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableId
	}
	return nil
}

func (m *Parameters) GetTableKeyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableKeyId
	}
	return nil
}

func (m *Parameters) GetActive() *dstore_values.IntegerValue {
	if m != nil {
		return m.Active
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
	RowId                   int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TableId                 *dstore_values.IntegerValue   `protobuf:"bytes,10001,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Active                  *dstore_values.BooleanValue   `protobuf:"bytes,10002,opt,name=active" json:"active,omitempty"`
	ContentDescription      *dstore_values.StringValue    `protobuf:"bytes,10003,opt,name=content_description,json=contentDescription" json:"content_description,omitempty"`
	TableKeyId              *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=table_key_id,json=tableKeyId" json:"table_key_id,omitempty"`
	PredefinedBillContentId *dstore_values.IntegerValue   `protobuf:"bytes,10005,opt,name=predefined_bill_content_id,json=predefinedBillContentId" json:"predefined_bill_content_id,omitempty"`
	CreatedAtDateAndTime    *dstore_values.TimestampValue `protobuf:"bytes,10006,opt,name=created_at_date_and_time,json=createdAtDateAndTime" json:"created_at_date_and_time,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTableId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableId
	}
	return nil
}

func (m *Response_Row) GetActive() *dstore_values.BooleanValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Response_Row) GetContentDescription() *dstore_values.StringValue {
	if m != nil {
		return m.ContentDescription
	}
	return nil
}

func (m *Response_Row) GetTableKeyId() *dstore_values.IntegerValue {
	if m != nil {
		return m.TableKeyId
	}
	return nil
}

func (m *Response_Row) GetPredefinedBillContentId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PredefinedBillContentId
	}
	return nil
}

func (m *Response_Row) GetCreatedAtDateAndTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.CreatedAtDateAndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.om_GetPredefinedBillContent_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.om_GetPredefinedBillContent_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.om_GetPredefinedBillContent_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x6d, 0x6b, 0xd4, 0x40,
	0x10, 0xa6, 0xc6, 0x6b, 0xcb, 0xd4, 0xd7, 0xad, 0xe8, 0x79, 0x87, 0x5a, 0xda, 0x2f, 0x22, 0x92,
	0x8a, 0x15, 0x2b, 0x42, 0x85, 0xab, 0x2d, 0x52, 0x6c, 0x43, 0x09, 0x72, 0x58, 0xbf, 0x84, 0xbd,
	0xcb, 0xf4, 0x08, 0x26, 0xbb, 0x61, 0xb3, 0xd7, 0xd2, 0x7f, 0xe1, 0xfb, 0xaf, 0xd2, 0x1f, 0xe2,
	0x0b, 0xf8, 0x17, 0x9c, 0x64, 0xf7, 0x7a, 0x97, 0xd8, 0xf6, 0x38, 0xfd, 0x72, 0x97, 0xc9, 0xcc,
	0xf3, 0x3c, 0x93, 0x99, 0x67, 0x17, 0xd6, 0xc2, 0x4c, 0x4b, 0x85, 0xcb, 0x28, 0x7a, 0x91, 0xc0,
	0xe5, 0x54, 0xc9, 0x2e, 0x86, 0x7d, 0x85, 0xd9, 0xb2, 0x4c, 0x82, 0x17, 0xa8, 0x77, 0x15, 0x86,
	0xb8, 0x4f, 0xa9, 0x70, 0x3d, 0x8a, 0xe3, 0xe7, 0x52, 0x68, 0x14, 0x3a, 0x68, 0x85, 0x2e, 0x55,
	0x6a, 0xc9, 0xee, 0x1b, 0xb8, 0x6b, 0xe0, 0xee, 0xd9, 0x98, 0xc6, 0xbc, 0x15, 0x3b, 0xe0, 0x71,
	0x1f, 0x33, 0x43, 0xd1, 0xb8, 0x59, 0xee, 0x00, 0x95, 0x92, 0xca, 0xa6, 0x9a, 0xe5, 0x54, 0x82,
	0x59, 0xc6, 0x7b, 0x68, 0x93, 0x4b, 0xd5, 0xa4, 0xe6, 0x91, 0xd8, 0x97, 0x2a, 0xe1, 0x3a, 0x92,
	0xc2, 0x14, 0x2d, 0x7e, 0x75, 0x00, 0x76, 0xb9, 0xe2, 0x94, 0x45, 0x95, 0xb1, 0xd7, 0xd0, 0x48,
	0x8f, 0x9b, 0x0b, 0x3a, 0xd4, 0x5d, 0xd0, 0xb5, 0xed, 0x45, 0x61, 0x7d, 0x6a, 0x61, 0xea, 0xee,
	0xdc, 0xc3, 0xa6, 0x6b, 0xbf, 0xc9, 0x76, 0x19, 0x51, 0xbe, 0x87, 0xaa, 0x9d, 0x47, 0xfe, 0x8d,
	0xf4, 0xa4, 0x6f, 0xdb, 0x0a, 0xd9, 0x26, 0xdc, 0x39, 0x9d, 0x39, 0x10, 0xfd, 0x38, 0xae, 0x7f,
	0x9f, 0x21, 0xfe, 0x59, 0xbf, 0x79, 0x0a, 0x85, 0x47, 0x35, 0xec, 0x31, 0xcc, 0x6a, 0xde, 0x89,
	0x31, 0x6f, 0xe7, 0xdc, 0xf8, 0x76, 0x66, 0x8a, 0x62, 0x92, 0x5f, 0x82, 0x8b, 0x03, 0x9c, 0x11,
	0xfb, 0x61, 0xc4, 0xe6, 0x6c, 0x41, 0x41, 0xbe, 0x06, 0x17, 0x4c, 0xd1, 0x5b, 0x3c, 0xca, 0x05,
	0x9c, 0xf1, 0x02, 0x50, 0x00, 0x5e, 0xe2, 0x11, 0x69, 0xdc, 0x83, 0xab, 0xa3, 0x70, 0xa3, 0xf3,
	0xd3, 0xe8, 0x5c, 0x1a, 0xd6, 0x15, 0x52, 0x2b, 0x30, 0xcd, 0xbb, 0x3a, 0x3a, 0xc0, 0xfa, 0xf9,
	0xf1, 0x22, 0xb6, 0x94, 0x2d, 0xc0, 0x9c, 0x79, 0x32, 0xd4, 0xbf, 0x0c, 0x35, 0x98, 0x77, 0x39,
	0xed, 0xe2, 0xef, 0x1a, 0xcc, 0xfa, 0x98, 0xa5, 0x52, 0x64, 0xc8, 0x1e, 0x40, 0xad, 0x30, 0x8b,
	0xdd, 0x5b, 0xc3, 0x2d, 0x7b, 0xd1, 0x18, 0x69, 0x33, 0xff, 0xf5, 0x4d, 0x21, 0xdb, 0x83, 0x2b,
	0xb9, 0x4d, 0x82, 0x11, 0x9f, 0xd0, 0x94, 0x1d, 0x02, 0xbb, 0x15, 0x70, 0xd5, 0x4d, 0x3b, 0x14,
	0x6f, 0x0d, 0x63, 0xff, 0x72, 0x52, 0x7e, 0xc1, 0x9e, 0xc0, 0x8c, 0xb5, 0x27, 0x8d, 0x35, 0x67,
	0xbc, 0xfd, 0x17, 0xa3, 0x31, 0xef, 0x8e, 0xf9, 0xf7, 0x07, 0xe5, 0x6c, 0x1b, 0x1c, 0x25, 0x0f,
	0x69, 0x4e, 0x39, 0xea, 0xa9, 0x3b, 0xc9, 0x81, 0x72, 0x07, 0xb3, 0x70, 0x7d, 0x79, 0xe8, 0xe7,
	0x34, 0x8d, 0x6f, 0x0e, 0x38, 0x14, 0xb0, 0xeb, 0x30, 0x4d, 0x61, 0xbe, 0xe5, 0x77, 0x1e, 0x8d,
	0xa7, 0xe6, 0xd7, 0x28, 0xa4, 0x25, 0xae, 0x8e, 0x18, 0xec, 0xbd, 0x37, 0x81, 0xc3, 0x1e, 0x1d,
	0x6f, 0xf4, 0xc3, 0xc9, 0xb0, 0x8e, 0x94, 0x31, 0x72, 0x51, 0x5e, 0xe9, 0x36, 0xcc, 0x0f, 0x8e,
	0x41, 0x88, 0x59, 0x57, 0x45, 0x69, 0x31, 0xf4, 0x8f, 0x5e, 0x79, 0x65, 0x96, 0x22, 0xd3, 0x2a,
	0x12, 0x3d, 0xc3, 0xc0, 0x2c, 0x6e, 0x63, 0x08, 0x63, 0xcf, 0x2a, 0x06, 0xfe, 0xe4, 0x4d, 0xe6,
	0xe0, 0xbd, 0x33, 0x8f, 0xff, 0x67, 0xef, 0x3f, 0xce, 0x7f, 0x1b, 0xea, 0x5d, 0x85, 0x5c, 0x13,
	0x2f, 0xa7, 0x6f, 0xa5, 0x87, 0x80, 0x8b, 0x30, 0xd0, 0x51, 0x82, 0xf5, 0x2f, 0x86, 0xf8, 0x56,
	0x85, 0x38, 0xcf, 0x65, 0x9a, 0x27, 0xa9, 0xa1, 0xbe, 0x66, 0xf1, 0x2d, 0xbd, 0x41, 0x7f, 0x2d,
	0x11, 0xbe, 0xa2, 0xfc, 0x7a, 0x1b, 0x9a, 0x91, 0xac, 0x98, 0x62, 0x78, 0x49, 0xbf, 0x59, 0xfd,
	0xc7, 0xeb, 0xbb, 0x33, 0x5d, 0xdc, 0x8f, 0x2b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xad, 0xfd,
	0xc9, 0x7f, 0x00, 0x06, 0x00, 0x00,
}