// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/im_InsertNodeBinary_Ad.proto
// DO NOT EDIT!

/*
Package im_InsertNodeBinary_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/im_InsertNodeBinary_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package im_InsertNodeBinary_Ad

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
	NodeIds                *dstore_values.StringValue    `protobuf:"bytes,1,opt,name=node_ids,json=nodeIds" json:"node_ids,omitempty"`
	NodeIdsNull            bool                          `protobuf:"varint,1001,opt,name=node_ids_null,json=nodeIdsNull" json:"node_ids_null,omitempty"`
	BinaryDescription      *dstore_values.StringValue    `protobuf:"bytes,2,opt,name=binary_description,json=binaryDescription" json:"binary_description,omitempty"`
	BinaryDescriptionNull  bool                          `protobuf:"varint,1002,opt,name=binary_description_null,json=binaryDescriptionNull" json:"binary_description_null,omitempty"`
	FromDate               *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	FromDateNull           bool                          `protobuf:"varint,1003,opt,name=from_date_null,json=fromDateNull" json:"from_date_null,omitempty"`
	ToDate                 *dstore_values.TimestampValue `protobuf:"bytes,4,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	ToDateNull             bool                          `protobuf:"varint,1004,opt,name=to_date_null,json=toDateNull" json:"to_date_null,omitempty"`
	FromTime               *dstore_values.TimestampValue `protobuf:"bytes,5,opt,name=from_time,json=fromTime" json:"from_time,omitempty"`
	FromTimeNull           bool                          `protobuf:"varint,1005,opt,name=from_time_null,json=fromTimeNull" json:"from_time_null,omitempty"`
	ToTime                 *dstore_values.TimestampValue `protobuf:"bytes,6,opt,name=to_time,json=toTime" json:"to_time,omitempty"`
	ToTimeNull             bool                          `protobuf:"varint,1006,opt,name=to_time_null,json=toTimeNull" json:"to_time_null,omitempty"`
	AutoContinue           *dstore_values.IntegerValue   `protobuf:"bytes,7,opt,name=auto_continue,json=autoContinue" json:"auto_continue,omitempty"`
	AutoContinueNull       bool                          `protobuf:"varint,1007,opt,name=auto_continue_null,json=autoContinueNull" json:"auto_continue_null,omitempty"`
	ReturnBinaryCodeId     *dstore_values.BooleanValue   `protobuf:"bytes,8,opt,name=return_binary_code_id,json=returnBinaryCodeId" json:"return_binary_code_id,omitempty"`
	ReturnBinaryCodeIdNull bool                          `protobuf:"varint,1008,opt,name=return_binary_code_id_null,json=returnBinaryCodeIdNull" json:"return_binary_code_id_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetNodeIds() *dstore_values.StringValue {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *Parameters) GetBinaryDescription() *dstore_values.StringValue {
	if m != nil {
		return m.BinaryDescription
	}
	return nil
}

func (m *Parameters) GetFromDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Parameters) GetToDate() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *Parameters) GetFromTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.FromTime
	}
	return nil
}

func (m *Parameters) GetToTime() *dstore_values.TimestampValue {
	if m != nil {
		return m.ToTime
	}
	return nil
}

func (m *Parameters) GetAutoContinue() *dstore_values.IntegerValue {
	if m != nil {
		return m.AutoContinue
	}
	return nil
}

func (m *Parameters) GetReturnBinaryCodeId() *dstore_values.BooleanValue {
	if m != nil {
		return m.ReturnBinaryCodeId
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
	RowId        int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	BinaryCodeId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=binary_code_id,json=binaryCodeId" json:"binary_code_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetBinaryCodeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryCodeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.im_InsertNodeBinary_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.im_InsertNodeBinary_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.im_InsertNodeBinary_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x49, 0x6f, 0x13, 0x3f,
	0x14, 0x57, 0x9b, 0x7f, 0x96, 0xff, 0xeb, 0x42, 0x31, 0x6a, 0x19, 0x12, 0x81, 0xa0, 0x05, 0x89,
	0x0b, 0x13, 0x54, 0x54, 0x40, 0x70, 0x81, 0x34, 0x1c, 0x72, 0xe8, 0x08, 0x8d, 0x10, 0x12, 0x5c,
	0x46, 0x4e, 0xc6, 0x8d, 0x2c, 0x65, 0xec, 0xc8, 0xf6, 0x50, 0xf1, 0x2d, 0xe0, 0x4b, 0x22, 0xb1,
	0x73, 0xe6, 0x84, 0x3d, 0xcf, 0xd9, 0x26, 0x91, 0x08, 0x97, 0x24, 0xf6, 0xfb, 0x6d, 0xb1, 0xfd,
	0x1e, 0x9c, 0xa4, 0xda, 0x48, 0xc5, 0xda, 0x4c, 0x0c, 0xb9, 0x60, 0xed, 0xb1, 0x92, 0x03, 0x96,
	0xe6, 0x8a, 0xe9, 0x36, 0xcf, 0x92, 0x9e, 0xd0, 0x4c, 0x99, 0x48, 0xa6, 0xac, 0xc3, 0x05, 0x55,
	0xef, 0x93, 0xe7, 0x69, 0x68, 0x11, 0x46, 0x92, 0xdb, 0x48, 0x0b, 0x91, 0x16, 0xae, 0xc6, 0x36,
	0xaf, 0x78, 0xf1, 0x77, 0x74, 0x94, 0x33, 0x8d, 0xd4, 0xe6, 0xb5, 0x45, 0x47, 0xa6, 0x94, 0x54,
	0xbe, 0xd4, 0x5a, 0x2c, 0x65, 0x4c, 0x6b, 0x3a, 0x64, 0xbe, 0x78, 0x54, 0x2e, 0x1a, 0xca, 0xc5,
	0xb9, 0x54, 0x19, 0x35, 0x5c, 0x0a, 0x04, 0x1d, 0x7e, 0xaa, 0x01, 0xbc, 0xa4, 0x8a, 0xda, 0x2a,
	0x53, 0x9a, 0x9c, 0x40, 0x43, 0xd8, 0x44, 0x09, 0x4f, 0x75, 0xb0, 0x71, 0x73, 0xe3, 0xee, 0xd6,
	0x71, 0x33, 0xf4, 0xc9, 0x7d, 0x26, 0x6d, 0x14, 0x17, 0xc3, 0xd7, 0x6e, 0x11, 0xd7, 0x1d, 0xb6,
	0x97, 0x6a, 0x72, 0x04, 0x3b, 0x13, 0x5a, 0x22, 0xf2, 0xd1, 0x28, 0xf8, 0x5c, 0xb7, 0xe4, 0x46,
	0xbc, 0xe5, 0x01, 0x91, 0xdd, 0x23, 0x3d, 0x20, 0x7d, 0xfc, 0xa7, 0x29, 0xd3, 0x03, 0xc5, 0xc7,
	0x2e, 0x46, 0xb0, 0xf9, 0x57, 0x97, 0xcb, 0xc8, 0xea, 0xce, 0x48, 0xe4, 0x11, 0x5c, 0x5d, 0x96,
	0x42, 0xe7, 0x2f, 0xe8, 0xbc, 0xbf, 0x44, 0x2a, 0x32, 0x3c, 0x81, 0xff, 0xcf, 0x95, 0xcc, 0x92,
	0x94, 0x1a, 0x16, 0x54, 0x0a, 0xeb, 0xeb, 0x25, 0x6b, 0xc3, 0xed, 0x31, 0x1a, 0x9a, 0x8d, 0xd1,
	0xbd, 0xe1, 0xf0, 0x5d, 0x0b, 0x27, 0x77, 0x60, 0x77, 0xca, 0x45, 0xaf, 0xaf, 0xe8, 0xb5, 0x3d,
	0x81, 0x14, 0x16, 0x0f, 0xa1, 0x6e, 0x24, 0x1a, 0xfc, 0xb7, 0x8e, 0x41, 0xcd, 0xc8, 0x42, 0xfe,
	0x16, 0x6c, 0x7b, 0x1e, 0x8a, 0x7f, 0x43, 0x71, 0xc0, 0xf2, 0x42, 0x7a, 0xa7, 0x10, 0x54, 0xd7,
	0x4e, 0xff, 0xca, 0xee, 0x4d, 0xd3, 0x3b, 0x00, 0x1a, 0x7c, 0x9f, 0x4b, 0xef, 0x20, 0x73, 0xe9,
	0x0b, 0x83, 0xda, 0x9a, 0xe9, 0x0b, 0x79, 0x4c, 0x3f, 0x13, 0xff, 0x31, 0x4d, 0x3f, 0x95, 0x7e,
	0x06, 0x3b, 0x34, 0xb7, 0xa0, 0x81, 0x14, 0x86, 0x8b, 0x9c, 0x05, 0xf5, 0xc2, 0xa0, 0x55, 0x32,
	0xe0, 0xc2, 0xb0, 0x21, 0x53, 0x28, 0xbf, 0xed, 0x18, 0xa7, 0x9e, 0x40, 0xee, 0x01, 0x59, 0x50,
	0x40, 0xab, 0x9f, 0x68, 0xb5, 0x37, 0x0f, 0x2d, 0x0c, 0x23, 0xd8, 0x57, 0xcc, 0xe4, 0x4a, 0x24,
	0xfe, 0xb1, 0x0c, 0xf0, 0x8d, 0x06, 0x8d, 0x95, 0xc6, 0x7d, 0x29, 0x47, 0x8c, 0x0a, 0x34, 0x26,
	0xc8, 0xc4, 0xd6, 0x3c, 0x2d, 0x5e, 0x31, 0x79, 0x0a, 0xcd, 0x95, 0x7a, 0x18, 0xe3, 0x17, 0xc6,
	0x38, 0x58, 0x26, 0xba, 0x30, 0x87, 0xbf, 0x37, 0xa1, 0x11, 0x33, 0x3d, 0x96, 0xb6, 0xed, 0xc9,
	0x7d, 0xa8, 0x16, 0x6d, 0x5c, 0xee, 0x31, 0x3f, 0x1d, 0xb0, 0xc5, 0x5f, 0xb8, 0xcf, 0x18, 0x81,
	0xe4, 0x0d, 0xec, 0xb9, 0x06, 0x4e, 0xe6, 0x3a, 0xd8, 0xb6, 0x4e, 0xc5, 0x92, 0xc3, 0x12, 0xb9,
	0xdc, 0xe7, 0x67, 0x76, 0xdd, 0x9b, 0xad, 0xe3, 0x4b, 0xd9, 0xe2, 0x06, 0x79, 0x0c, 0x75, 0x3f,
	0x38, 0x6c, 0x47, 0x38, 0xc5, 0x1b, 0x4b, 0x8a, 0x38, 0x56, 0xce, 0xf0, 0x3b, 0x9e, 0xc0, 0x49,
	0x17, 0x2a, 0x4a, 0x5e, 0xd8, 0x67, 0xee, 0x58, 0xc7, 0xe1, 0x3a, 0x23, 0x2e, 0x9c, 0x9c, 0x41,
	0x18, 0xcb, 0x8b, 0xd8, 0xd1, 0x9b, 0x14, 0x2a, 0xf6, 0x37, 0x39, 0x80, 0x9a, 0x5d, 0xb9, 0xeb,
	0xf9, 0x10, 0xd9, 0x53, 0xa9, 0xc6, 0x55, 0xbb, 0xb4, 0xa7, 0xde, 0x81, 0xdd, 0xd2, 0xf5, 0x7d,
	0x8c, 0xd6, 0x78, 0x38, 0xfd, 0xb9, 0x0b, 0xe8, 0x44, 0xd0, 0xe2, 0xb2, 0x94, 0x6f, 0x36, 0xb9,
	0xdf, 0xb6, 0xff, 0x71, 0xa6, 0xf7, 0x6b, 0xc5, 0xf0, 0x7c, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0x16, 0x33, 0x66, 0x0d, 0x06, 0x00, 0x00,
}
