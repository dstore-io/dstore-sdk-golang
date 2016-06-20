// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_LoadDatabase_Ad.proto
// DO NOT EDIT!

/*
Package mi_LoadDatabase_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_LoadDatabase_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_LoadDatabase_Ad

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
	DatabaseName                *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=database_name,json=databaseName" json:"database_name,omitempty"`
	DatabaseNameNull            bool                        `protobuf:"varint,1001,opt,name=database_name_null,json=databaseNameNull" json:"database_name_null,omitempty"`
	DeviceNames                 *dstore_values.StringValue  `protobuf:"bytes,2,opt,name=device_names,json=deviceNames" json:"device_names,omitempty"`
	DeviceNamesNull             bool                        `protobuf:"varint,1002,opt,name=device_names_null,json=deviceNamesNull" json:"device_names_null,omitempty"`
	DeviceNamesSeparatedBy      *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=device_names_separated_by,json=deviceNamesSeparatedBy" json:"device_names_separated_by,omitempty"`
	DeviceNamesSeparatedByNull  bool                        `protobuf:"varint,1003,opt,name=device_names_separated_by_null,json=deviceNamesSeparatedByNull" json:"device_names_separated_by_null,omitempty"`
	DevicePath                  *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=device_path,json=devicePath" json:"device_path,omitempty"`
	DevicePathNull              bool                        `protobuf:"varint,1004,opt,name=device_path_null,json=devicePathNull" json:"device_path_null,omitempty"`
	KillProcessesOnDBFirst      *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=kill_processes_on_d_b_first,json=killProcessesOnDBFirst" json:"kill_processes_on_d_b_first,omitempty"`
	KillProcessesOnDBFirstNull  bool                        `protobuf:"varint,1005,opt,name=kill_processes_on_d_b_first_null,json=killProcessesOnDBFirstNull" json:"kill_processes_on_d_b_first_null,omitempty"`
	OnlineDatabaseAfterLoad     *dstore_values.BooleanValue `protobuf:"bytes,6,opt,name=online_database_after_load,json=onlineDatabaseAfterLoad" json:"online_database_after_load,omitempty"`
	OnlineDatabaseAfterLoadNull bool                        `protobuf:"varint,1006,opt,name=online_database_after_load_null,json=onlineDatabaseAfterLoadNull" json:"online_database_after_load_null,omitempty"`
	AdjustDBUserToLogins        *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=adjust_d_b_user_to_logins,json=adjustDBUserToLogins" json:"adjust_d_b_user_to_logins,omitempty"`
	AdjustDBUserToLoginsNull    bool                        `protobuf:"varint,1007,opt,name=adjust_d_b_user_to_logins_null,json=adjustDBUserToLoginsNull" json:"adjust_d_b_user_to_logins_null,omitempty"`
	CompressLevel               *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=compress_level,json=compressLevel" json:"compress_level,omitempty"`
	CompressLevelNull           bool                        `protobuf:"varint,1008,opt,name=compress_level_null,json=compressLevelNull" json:"compress_level_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetDatabaseName() *dstore_values.StringValue {
	if m != nil {
		return m.DatabaseName
	}
	return nil
}

func (m *Parameters) GetDeviceNames() *dstore_values.StringValue {
	if m != nil {
		return m.DeviceNames
	}
	return nil
}

func (m *Parameters) GetDeviceNamesSeparatedBy() *dstore_values.StringValue {
	if m != nil {
		return m.DeviceNamesSeparatedBy
	}
	return nil
}

func (m *Parameters) GetDevicePath() *dstore_values.StringValue {
	if m != nil {
		return m.DevicePath
	}
	return nil
}

func (m *Parameters) GetKillProcessesOnDBFirst() *dstore_values.BooleanValue {
	if m != nil {
		return m.KillProcessesOnDBFirst
	}
	return nil
}

func (m *Parameters) GetOnlineDatabaseAfterLoad() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlineDatabaseAfterLoad
	}
	return nil
}

func (m *Parameters) GetAdjustDBUserToLogins() *dstore_values.BooleanValue {
	if m != nil {
		return m.AdjustDBUserToLogins
	}
	return nil
}

func (m *Parameters) GetCompressLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.CompressLevel
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
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_LoadDatabase_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_LoadDatabase_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_LoadDatabase_Ad.Response.Row")
}

var fileDescriptor0 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x4e, 0x13, 0x4f,
	0x18, 0x0f, 0xf4, 0x5f, 0x68, 0x3e, 0xce, 0xc3, 0x3f, 0xb8, 0xb4, 0x11, 0x09, 0xde, 0x68, 0x88,
	0xad, 0x87, 0x1b, 0x13, 0x63, 0x94, 0x0a, 0x26, 0x18, 0x40, 0xb2, 0x82, 0x51, 0x6f, 0x26, 0x53,
	0x76, 0xa8, 0xab, 0xdb, 0x99, 0x66, 0x66, 0x0b, 0xe1, 0x2d, 0x7c, 0x16, 0xaf, 0x7d, 0x21, 0xcf,
	0x3e, 0x82, 0xdf, 0xec, 0xb7, 0x2b, 0xdd, 0x6d, 0x6b, 0xe3, 0x0d, 0x74, 0xf6, 0xfb, 0x9d, 0x66,
	0x33, 0xf3, 0x5b, 0xb8, 0x13, 0xd8, 0x58, 0x1b, 0xd9, 0x90, 0xaa, 0x1d, 0x2a, 0xd9, 0xe8, 0x1a,
	0x7d, 0x22, 0x83, 0x9e, 0x91, 0xb6, 0xd1, 0x09, 0xf9, 0x9e, 0x16, 0xc1, 0xb6, 0x88, 0x45, 0x4b,
	0x58, 0xc9, 0xb7, 0x82, 0x3a, 0x4e, 0x63, 0xcd, 0xd6, 0x89, 0x52, 0x27, 0x4a, 0x7d, 0x10, 0x57,
	0x5d, 0x4e, 0x45, 0xcf, 0x44, 0xd4, 0x93, 0x96, 0x68, 0xd5, 0xd5, 0xbc, 0x93, 0x34, 0x46, 0x9b,
	0x74, 0x54, 0xcb, 0x8f, 0x3a, 0xd2, 0x5a, 0xd1, 0x96, 0xe9, 0xf0, 0x7a, 0x71, 0x18, 0x8b, 0x50,
	0x9d, 0x6a, 0xd3, 0x11, 0x71, 0xa8, 0x15, 0x81, 0x36, 0x3e, 0x55, 0x00, 0x0e, 0x85, 0x11, 0x38,
	0x95, 0xc6, 0xb2, 0x47, 0x30, 0x17, 0x64, 0x79, 0x14, 0x3e, 0xf4, 0x26, 0xd6, 0x27, 0x6e, 0xcc,
	0xdc, 0xad, 0xd6, 0xd3, 0xe8, 0x69, 0x30, 0x1b, 0x9b, 0x50, 0xb5, 0x5f, 0xba, 0x85, 0x3f, 0x9b,
	0x11, 0x0e, 0x10, 0xcf, 0x6e, 0x01, 0xcb, 0x09, 0x70, 0xd5, 0x8b, 0x22, 0xef, 0xf3, 0x34, 0xca,
	0x54, 0xfc, 0xc5, 0x7e, 0xe8, 0x01, 0x0e, 0xd8, 0x43, 0x98, 0x0d, 0xe4, 0x59, 0x78, 0x42, 0x60,
	0xeb, 0x4d, 0x8e, 0xb5, 0x9b, 0x21, 0xbc, 0x53, 0xb0, 0x6c, 0x13, 0x96, 0xfa, 0xe9, 0x64, 0xf6,
	0x85, 0xcc, 0x16, 0xfa, 0x80, 0x89, 0xd7, 0x31, 0xac, 0xe6, 0xc0, 0x56, 0x76, 0x71, 0xe3, 0xb1,
	0x0c, 0x78, 0xeb, 0xc2, 0x2b, 0x8d, 0x35, 0x5e, 0xe9, 0xd3, 0x7b, 0x91, 0x51, 0x9b, 0x17, 0xec,
	0x09, 0xac, 0x8d, 0x94, 0xa5, 0x40, 0x5f, 0x29, 0x50, 0x75, 0xb8, 0x40, 0x92, 0xed, 0x01, 0xa4,
	0xfb, 0xe2, 0x5d, 0x11, 0xbf, 0xf5, 0xfe, 0x1b, 0x9b, 0x06, 0x08, 0x7e, 0x88, 0x68, 0x76, 0x13,
	0x16, 0xfb, 0xc8, 0xe4, 0xf9, 0x8d, 0x3c, 0xe7, 0x2f, 0x61, 0x89, 0xcf, 0x2b, 0xa8, 0xbd, 0x0f,
	0xa3, 0x88, 0x27, 0xc7, 0xd5, 0x5a, 0x8c, 0xab, 0x15, 0xc7, 0xa4, 0xfc, 0x34, 0x34, 0x36, 0xf6,
	0xca, 0x89, 0x6f, 0xad, 0xe0, 0xdb, 0xd2, 0x3a, 0x92, 0x42, 0xa5, 0xaf, 0xc1, 0xf1, 0x0f, 0x33,
	0xfa, 0x73, 0xb5, 0xdd, 0x7c, 0xea, 0xa8, 0x6c, 0x07, 0xd6, 0xff, 0xa2, 0x4c, 0xa1, 0xbe, 0xa7,
	0x2f, 0x62, 0xb8, 0x44, 0x1a, 0xb0, 0xaa, 0x55, 0x84, 0xe7, 0x95, 0xff, 0x39, 0x46, 0xe2, 0x14,
	0x8f, 0x26, 0x8f, 0xf0, 0xaa, 0x78, 0x53, 0xe3, 0xf3, 0x5d, 0x21, 0x7a, 0x76, 0xab, 0xb6, 0x1c,
	0xd9, 0x5d, 0x33, 0x0c, 0x78, 0x6d, 0xb4, 0x32, 0xe5, 0xfb, 0x41, 0xf9, 0x6a, 0x23, 0x24, 0x92,
	0x80, 0x47, 0xb0, 0x2a, 0x82, 0x77, 0x3d, 0xdc, 0x92, 0xdb, 0x5c, 0xcf, 0xa2, 0x40, 0xac, 0x51,
	0x03, 0xef, 0x98, 0xf5, 0xa6, 0xc7, 0xe7, 0xfb, 0x9f, 0xd8, 0xdb, 0xcd, 0x63, 0xa4, 0x1e, 0xe9,
	0xbd, 0x84, 0xc8, 0xb6, 0x60, 0x6d, 0xa4, 0x2a, 0x65, 0xfb, 0x49, 0xd9, 0xbc, 0x61, 0xf4, 0x24,
	0x58, 0x13, 0xe6, 0x4f, 0x74, 0xa7, 0x8b, 0x0d, 0x64, 0x79, 0x24, 0xcf, 0x64, 0xe4, 0x55, 0x86,
	0xa6, 0x09, 0x55, 0x2c, 0xdb, 0xd2, 0x50, 0x9a, 0xb9, 0x8c, 0xb2, 0xe7, 0x18, 0xac, 0x01, 0xcb,
	0x79, 0x0d, 0xf2, 0xfe, 0x45, 0xde, 0x4b, 0x39, 0xb0, 0x33, 0xdd, 0xf8, 0x38, 0x09, 0x15, 0x5f,
	0xda, 0xae, 0x56, 0x56, 0xb2, 0xdb, 0x50, 0x4e, 0xca, 0xa9, 0x58, 0x1a, 0x69, 0xdf, 0x51, 0x71,
	0xed, 0xb8, 0xbf, 0x3e, 0x01, 0xd9, 0x6b, 0x58, 0x74, 0xb5, 0xc4, 0xfb, 0x7a, 0x09, 0x2b, 0xa0,
	0x84, 0xe4, 0x7a, 0x81, 0x5c, 0x6c, 0xaf, 0x7d, 0x5c, 0xef, 0x5e, 0xae, 0xfd, 0x85, 0x4e, 0xfe,
	0x01, 0xbb, 0x0f, 0xd3, 0x69, 0x1d, 0xe2, 0xdd, 0x76, 0x8a, 0x6b, 0x03, 0x8a, 0x54, 0x96, 0xfb,
	0xf4, 0xdf, 0xcf, 0xe0, 0xec, 0x31, 0x94, 0x8c, 0x3e, 0xc7, 0x3b, 0x38, 0x34, 0xc7, 0x60, 0xb9,
	0x67, 0xfb, 0xaf, 0xfb, 0xfa, 0xdc, 0x77, 0xd4, 0xea, 0x55, 0x28, 0xe1, 0x6f, 0xb6, 0x02, 0x53,
	0xb8, 0xe2, 0x61, 0xe0, 0x7d, 0x38, 0xc0, 0x37, 0x52, 0xf6, 0xcb, 0xb8, 0xdc, 0x0d, 0x9a, 0xcf,
	0xa0, 0x16, 0xea, 0x82, 0xee, 0xe5, 0xf7, 0xe3, 0xcd, 0xe6, 0x3f, 0x7c, 0x59, 0x5a, 0x53, 0x49,
	0x8d, 0xdf, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x11, 0x65, 0xce, 0xff, 0x8f, 0x06, 0x00, 0x00,
}
