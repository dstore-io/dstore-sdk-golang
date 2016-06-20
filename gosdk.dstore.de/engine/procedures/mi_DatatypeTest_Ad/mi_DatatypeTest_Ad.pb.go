// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_DatatypeTest_Ad.proto
// DO NOT EDIT!

/*
Package mi_DatatypeTest_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_DatatypeTest_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_DatatypeTest_Ad

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
	SetOutputParams     *dstore_values.BooleanValue `protobuf:"bytes,1,opt,name=set_output_params,json=setOutputParams" json:"set_output_params,omitempty"`
	SetOutputParamsNull bool                        `protobuf:"varint,1001,opt,name=set_output_params_null,json=setOutputParamsNull" json:"set_output_params_null,omitempty"`
	GetResultSet        *dstore_values.BooleanValue `protobuf:"bytes,2,opt,name=get_result_set,json=getResultSet" json:"get_result_set,omitempty"`
	GetResultSetNull    bool                        `protobuf:"varint,1002,opt,name=get_result_set_null,json=getResultSetNull" json:"get_result_set_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetSetOutputParams() *dstore_values.BooleanValue {
	if m != nil {
		return m.SetOutputParams
	}
	return nil
}

func (m *Parameters) GetGetResultSet() *dstore_values.BooleanValue {
	if m != nil {
		return m.GetResultSet
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	TestBit         *dstore_values.BooleanValue                      `protobuf:"bytes,101,opt,name=test_bit,json=testBit" json:"test_bit,omitempty"`
	TestChar        *dstore_values.StringValue                       `protobuf:"bytes,102,opt,name=test_char,json=testChar" json:"test_char,omitempty"`
	TestDatetime    *dstore_values.TimestampValue                    `protobuf:"bytes,103,opt,name=test_datetime,json=testDatetime" json:"test_datetime,omitempty"`
	TestDecimal     *dstore_values.DecimalValue                      `protobuf:"bytes,104,opt,name=test_decimal,json=testDecimal" json:"test_decimal,omitempty"`
	TestImage       *dstore_values.BytesValue                        `protobuf:"bytes,105,opt,name=test_image,json=testImage" json:"test_image,omitempty"`
	TestInteger     *dstore_values.IntegerValue                      `protobuf:"bytes,106,opt,name=test_integer,json=testInteger" json:"test_integer,omitempty"`
	TestMoney       *dstore_values.DecimalValue                      `protobuf:"bytes,107,opt,name=test_money,json=testMoney" json:"test_money,omitempty"`
	TestSmallint    *dstore_values.IntegerValue                      `protobuf:"bytes,108,opt,name=test_smallint,json=testSmallint" json:"test_smallint,omitempty"`
	TestText        *dstore_values.StringValue                       `protobuf:"bytes,109,opt,name=test_text,json=testText" json:"test_text,omitempty"`
	TestTinyint     *dstore_values.IntegerValue                      `protobuf:"bytes,110,opt,name=test_tinyint,json=testTinyint" json:"test_tinyint,omitempty"`
	TestVarchar     *dstore_values.StringValue                       `protobuf:"bytes,111,opt,name=test_varchar,json=testVarchar" json:"test_varchar,omitempty"`
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

func (m *Response) GetTestBit() *dstore_values.BooleanValue {
	if m != nil {
		return m.TestBit
	}
	return nil
}

func (m *Response) GetTestChar() *dstore_values.StringValue {
	if m != nil {
		return m.TestChar
	}
	return nil
}

func (m *Response) GetTestDatetime() *dstore_values.TimestampValue {
	if m != nil {
		return m.TestDatetime
	}
	return nil
}

func (m *Response) GetTestDecimal() *dstore_values.DecimalValue {
	if m != nil {
		return m.TestDecimal
	}
	return nil
}

func (m *Response) GetTestImage() *dstore_values.BytesValue {
	if m != nil {
		return m.TestImage
	}
	return nil
}

func (m *Response) GetTestInteger() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestInteger
	}
	return nil
}

func (m *Response) GetTestMoney() *dstore_values.DecimalValue {
	if m != nil {
		return m.TestMoney
	}
	return nil
}

func (m *Response) GetTestSmallint() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestSmallint
	}
	return nil
}

func (m *Response) GetTestText() *dstore_values.StringValue {
	if m != nil {
		return m.TestText
	}
	return nil
}

func (m *Response) GetTestTinyint() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestTinyint
	}
	return nil
}

func (m *Response) GetTestVarchar() *dstore_values.StringValue {
	if m != nil {
		return m.TestVarchar
	}
	return nil
}

type Response_Row struct {
	RowId                         int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	TestText                      *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=test_text,json=testText" json:"test_text,omitempty"`
	TestDecimalParamInput         *dstore_values.StringValue    `protobuf:"bytes,10002,opt,name=test_decimal_param_input,json=testDecimalParamInput" json:"test_decimal_param_input,omitempty"`
	TestDatetime                  *dstore_values.TimestampValue `protobuf:"bytes,10003,opt,name=test_datetime,json=testDatetime" json:"test_datetime,omitempty"`
	TestTextParamInputByteLength  *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=test_text_param_input_byte_length,json=testTextParamInputByteLength" json:"test_text_param_input_byte_length,omitempty"`
	TestTextParamInput            *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=test_text_param_input,json=testTextParamInput" json:"test_text_param_input,omitempty"`
	TestDecimal                   *dstore_values.DecimalValue   `protobuf:"bytes,10006,opt,name=test_decimal,json=testDecimal" json:"test_decimal,omitempty"`
	TestSmallintParamInput        *dstore_values.StringValue    `protobuf:"bytes,10007,opt,name=test_smallint_param_input,json=testSmallintParamInput" json:"test_smallint_param_input,omitempty"`
	TestChar                      *dstore_values.StringValue    `protobuf:"bytes,10008,opt,name=test_char,json=testChar" json:"test_char,omitempty"`
	TestTinyintParamInput         *dstore_values.StringValue    `protobuf:"bytes,10009,opt,name=test_tinyint_param_input,json=testTinyintParamInput" json:"test_tinyint_param_input,omitempty"`
	TestBit                       *dstore_values.BooleanValue   `protobuf:"bytes,10010,opt,name=test_bit,json=testBit" json:"test_bit,omitempty"`
	TestInteger                   *dstore_values.IntegerValue   `protobuf:"bytes,10011,opt,name=test_integer,json=testInteger" json:"test_integer,omitempty"`
	TestTinyint                   *dstore_values.IntegerValue   `protobuf:"bytes,10012,opt,name=test_tinyint,json=testTinyint" json:"test_tinyint,omitempty"`
	TestMoneyParamInput           *dstore_values.StringValue    `protobuf:"bytes,10013,opt,name=test_money_param_input,json=testMoneyParamInput" json:"test_money_param_input,omitempty"`
	TestImageParamInputByteLength *dstore_values.IntegerValue   `protobuf:"bytes,10014,opt,name=test_image_param_input_byte_length,json=testImageParamInputByteLength" json:"test_image_param_input_byte_length,omitempty"`
	TestTextParamInputMD5         *dstore_values.StringValue    `protobuf:"bytes,10015,opt,name=test_text_param_input_m_d5,json=testTextParamInputMD5" json:"test_text_param_input_m_d5,omitempty"`
	TestCharParamInput            *dstore_values.StringValue    `protobuf:"bytes,10016,opt,name=test_char_param_input,json=testCharParamInput" json:"test_char_param_input,omitempty"`
	TestSmallint                  *dstore_values.IntegerValue   `protobuf:"bytes,10017,opt,name=test_smallint,json=testSmallint" json:"test_smallint,omitempty"`
	TestBitParamInput             *dstore_values.StringValue    `protobuf:"bytes,10018,opt,name=test_bit_param_input,json=testBitParamInput" json:"test_bit_param_input,omitempty"`
	TestImage                     *dstore_values.BytesValue     `protobuf:"bytes,10019,opt,name=test_image,json=testImage" json:"test_image,omitempty"`
	TestBigint                    *dstore_values.LongValue      `protobuf:"bytes,10020,opt,name=test_bigint,json=testBigint" json:"test_bigint,omitempty"`
	TestImageParamInputMD5        *dstore_values.StringValue    `protobuf:"bytes,10021,opt,name=test_image_param_input_m_d5,json=testImageParamInputMD5" json:"test_image_param_input_m_d5,omitempty"`
	TestIntegerParamInput         *dstore_values.StringValue    `protobuf:"bytes,10022,opt,name=test_integer_param_input,json=testIntegerParamInput" json:"test_integer_param_input,omitempty"`
	TestSysname                   *dstore_values.StringValue    `protobuf:"bytes,10023,opt,name=test_sysname,json=testSysname" json:"test_sysname,omitempty"`
	TestVarcharParamInput         *dstore_values.StringValue    `protobuf:"bytes,10024,opt,name=test_varchar_param_input,json=testVarcharParamInput" json:"test_varchar_param_input,omitempty"`
	TestDatetimeParamInput        *dstore_values.StringValue    `protobuf:"bytes,10025,opt,name=test_datetime_param_input,json=testDatetimeParamInput" json:"test_datetime_param_input,omitempty"`
	TestImageParamInput           *dstore_values.StringValue    `protobuf:"bytes,10026,opt,name=test_image_param_input,json=testImageParamInput" json:"test_image_param_input,omitempty"`
	TestMoney                     *dstore_values.DecimalValue   `protobuf:"bytes,10027,opt,name=test_money,json=testMoney" json:"test_money,omitempty"`
	TestVarchar                   *dstore_values.StringValue    `protobuf:"bytes,10028,opt,name=test_varchar,json=testVarchar" json:"test_varchar,omitempty"`
	TestNumeric                   *dstore_values.DecimalValue   `protobuf:"bytes,10029,opt,name=test_numeric,json=testNumeric" json:"test_numeric,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetTestText() *dstore_values.StringValue {
	if m != nil {
		return m.TestText
	}
	return nil
}

func (m *Response_Row) GetTestDecimalParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestDecimalParamInput
	}
	return nil
}

func (m *Response_Row) GetTestDatetime() *dstore_values.TimestampValue {
	if m != nil {
		return m.TestDatetime
	}
	return nil
}

func (m *Response_Row) GetTestTextParamInputByteLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestTextParamInputByteLength
	}
	return nil
}

func (m *Response_Row) GetTestTextParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestTextParamInput
	}
	return nil
}

func (m *Response_Row) GetTestDecimal() *dstore_values.DecimalValue {
	if m != nil {
		return m.TestDecimal
	}
	return nil
}

func (m *Response_Row) GetTestSmallintParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestSmallintParamInput
	}
	return nil
}

func (m *Response_Row) GetTestChar() *dstore_values.StringValue {
	if m != nil {
		return m.TestChar
	}
	return nil
}

func (m *Response_Row) GetTestTinyintParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestTinyintParamInput
	}
	return nil
}

func (m *Response_Row) GetTestBit() *dstore_values.BooleanValue {
	if m != nil {
		return m.TestBit
	}
	return nil
}

func (m *Response_Row) GetTestInteger() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestInteger
	}
	return nil
}

func (m *Response_Row) GetTestTinyint() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestTinyint
	}
	return nil
}

func (m *Response_Row) GetTestMoneyParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestMoneyParamInput
	}
	return nil
}

func (m *Response_Row) GetTestImageParamInputByteLength() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestImageParamInputByteLength
	}
	return nil
}

func (m *Response_Row) GetTestTextParamInputMD5() *dstore_values.StringValue {
	if m != nil {
		return m.TestTextParamInputMD5
	}
	return nil
}

func (m *Response_Row) GetTestCharParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestCharParamInput
	}
	return nil
}

func (m *Response_Row) GetTestSmallint() *dstore_values.IntegerValue {
	if m != nil {
		return m.TestSmallint
	}
	return nil
}

func (m *Response_Row) GetTestBitParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestBitParamInput
	}
	return nil
}

func (m *Response_Row) GetTestImage() *dstore_values.BytesValue {
	if m != nil {
		return m.TestImage
	}
	return nil
}

func (m *Response_Row) GetTestBigint() *dstore_values.LongValue {
	if m != nil {
		return m.TestBigint
	}
	return nil
}

func (m *Response_Row) GetTestImageParamInputMD5() *dstore_values.StringValue {
	if m != nil {
		return m.TestImageParamInputMD5
	}
	return nil
}

func (m *Response_Row) GetTestIntegerParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestIntegerParamInput
	}
	return nil
}

func (m *Response_Row) GetTestSysname() *dstore_values.StringValue {
	if m != nil {
		return m.TestSysname
	}
	return nil
}

func (m *Response_Row) GetTestVarcharParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestVarcharParamInput
	}
	return nil
}

func (m *Response_Row) GetTestDatetimeParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestDatetimeParamInput
	}
	return nil
}

func (m *Response_Row) GetTestImageParamInput() *dstore_values.StringValue {
	if m != nil {
		return m.TestImageParamInput
	}
	return nil
}

func (m *Response_Row) GetTestMoney() *dstore_values.DecimalValue {
	if m != nil {
		return m.TestMoney
	}
	return nil
}

func (m *Response_Row) GetTestVarchar() *dstore_values.StringValue {
	if m != nil {
		return m.TestVarchar
	}
	return nil
}

func (m *Response_Row) GetTestNumeric() *dstore_values.DecimalValue {
	if m != nil {
		return m.TestNumeric
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_DatatypeTest_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_DatatypeTest_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_DatatypeTest_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/mi_DatatypeTest_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1043 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x57, 0x5b, 0x6f, 0xdc, 0x44,
	0x14, 0x56, 0x08, 0x69, 0xc2, 0xa4, 0xd0, 0x66, 0x42, 0xa3, 0xc9, 0x86, 0xa2, 0x12, 0x5e, 0x78,
	0x72, 0xb8, 0xb4, 0xb4, 0x14, 0xd1, 0x90, 0x34, 0x08, 0xad, 0xd4, 0x0c, 0x95, 0x13, 0x45, 0xc0,
	0x8b, 0xe5, 0xac, 0xa7, 0x8e, 0xa9, 0xed, 0x59, 0x79, 0x66, 0x5b, 0xf6, 0x5f, 0x70, 0xbf, 0xdf,
	0x2f, 0xe5, 0xfe, 0xa3, 0x78, 0x84, 0x1f, 0x81, 0x98, 0x9b, 0xbd, 0x33, 0xde, 0x8d, 0x6c, 0xbf,
	0x24, 0xb2, 0xcf, 0xf9, 0xce, 0xf9, 0xce, 0x39, 0xdf, 0x9c, 0xf1, 0x82, 0xe7, 0x22, 0xc6, 0x69,
	0x41, 0xb6, 0x48, 0x1e, 0x27, 0x39, 0xd9, 0x1a, 0x16, 0x74, 0x40, 0xa2, 0x51, 0x41, 0xd8, 0x56,
	0x96, 0x04, 0x7b, 0x21, 0x0f, 0xf9, 0x78, 0x48, 0x0e, 0x09, 0xe3, 0xc1, 0x4e, 0xe4, 0x09, 0x2b,
	0xa7, 0xf0, 0x92, 0x86, 0x78, 0x1a, 0xe2, 0x4d, 0xfb, 0xf5, 0x56, 0x4d, 0xd0, 0x7b, 0x61, 0x3a,
	0x22, 0x4c, 0xc3, 0x7a, 0xeb, 0x6e, 0x26, 0x52, 0x14, 0xb4, 0x30, 0xa6, 0x0d, 0xd7, 0x94, 0x11,
	0xc6, 0xc2, 0x98, 0x18, 0xe3, 0xd3, 0x75, 0x23, 0x0f, 0x93, 0xfc, 0x0e, 0x2d, 0xb2, 0x90, 0x27,
	0x34, 0xd7, 0x4e, 0x9b, 0xff, 0xcd, 0x01, 0x70, 0x3b, 0x2c, 0x42, 0x61, 0x25, 0x05, 0x83, 0xaf,
	0x83, 0x15, 0x46, 0x78, 0x40, 0x47, 0x7c, 0x38, 0xe2, 0xc1, 0x50, 0x1a, 0x18, 0x9a, 0xbb, 0x34,
	0xf7, 0xcc, 0xf2, 0xf3, 0x1b, 0x9e, 0xa1, 0x6f, 0xc8, 0x1d, 0x53, 0x9a, 0x92, 0x30, 0x3f, 0x92,
	0x4f, 0xfe, 0x39, 0x81, 0x7a, 0x43, 0x81, 0x54, 0x30, 0x06, 0x2f, 0x83, 0xb5, 0xa9, 0x40, 0x41,
	0x3e, 0x4a, 0x53, 0xf4, 0xcf, 0xa2, 0x08, 0xb7, 0xe4, 0xaf, 0xd6, 0x10, 0x58, 0xd8, 0xe0, 0x0e,
	0x78, 0x2c, 0x16, 0x28, 0xd1, 0xc5, 0x51, 0xca, 0x03, 0xe1, 0x81, 0x1e, 0x6a, 0xce, 0x7d, 0x56,
	0x40, 0x7c, 0x85, 0x38, 0x20, 0x1c, 0x7a, 0x60, 0xd5, 0x0d, 0xa1, 0xb3, 0xfe, 0xab, 0xb3, 0x9e,
	0xb7, 0x7d, 0x65, 0xca, 0xcd, 0xbf, 0x11, 0x58, 0x12, 0x6f, 0x86, 0x34, 0x67, 0x04, 0x3e, 0x0b,
	0x16, 0x54, 0x7b, 0x4d, 0xc9, 0x3d, 0xcf, 0x9d, 0x98, 0x6e, 0xfd, 0x6b, 0xf2, 0xaf, 0xaf, 0x1d,
	0xe1, 0x5b, 0xe0, 0xbc, 0x6c, 0x6c, 0x60, 0x75, 0x56, 0x70, 0x9e, 0x17, 0x60, 0xaf, 0x06, 0xae,
	0xf7, 0x7f, 0x5f, 0x3c, 0xf7, 0x27, 0xcf, 0xfe, 0xb9, 0xcc, 0x7d, 0x01, 0xaf, 0x81, 0x45, 0x33,
	0x50, 0x34, 0xaf, 0x22, 0x3e, 0x39, 0x15, 0x51, 0x8f, 0x7b, 0x5f, 0xff, 0xf7, 0x4b, 0x77, 0xf8,
	0x2a, 0x98, 0x2f, 0xe8, 0x7d, 0xf4, 0xf0, 0x6c, 0x1e, 0xd3, 0xf2, 0x2c, 0xeb, 0xf7, 0x7c, 0x7a,
	0xdf, 0x97, 0x50, 0xf8, 0x22, 0x58, 0xe2, 0xd2, 0x78, 0x9c, 0x70, 0x44, 0x9a, 0x47, 0xb0, 0x28,
	0x9d, 0x77, 0x13, 0x0e, 0xaf, 0x82, 0x47, 0x14, 0x6e, 0x70, 0x12, 0x16, 0xe8, 0x8e, 0xdb, 0x44,
	0x03, 0x64, 0xbc, 0x48, 0xf2, 0x58, 0xe3, 0x54, 0x92, 0x9b, 0xc2, 0x17, 0xee, 0x82, 0x47, 0x15,
	0x30, 0x0a, 0x39, 0xe1, 0x49, 0x46, 0x50, 0xac, 0xc0, 0x17, 0x6b, 0x60, 0x69, 0x62, 0x3c, 0xcc,
	0x86, 0x66, 0xf4, 0x12, 0xb3, 0x67, 0x20, 0xf0, 0x06, 0x38, 0xab, 0x63, 0x90, 0x41, 0x92, 0x85,
	0x29, 0x3a, 0x99, 0x49, 0xdc, 0x58, 0x75, 0x80, 0x65, 0x15, 0x40, 0xbf, 0x11, 0x0d, 0x07, 0x0a,
	0x2f, 0x1e, 0x44, 0xcf, 0x13, 0x85, 0x5e, 0xaf, 0x97, 0x3d, 0x16, 0x2e, 0x1a, 0xab, 0x2a, 0xed,
	0x4b, 0xdf, 0x2a, 0x73, 0x92, 0x73, 0x12, 0x93, 0x02, 0xbd, 0x33, 0x33, 0xb3, 0xb1, 0x5a, 0x99,
	0xfb, 0xfa, 0x0d, 0xbc, 0x6e, 0x32, 0x67, 0x34, 0x27, 0x63, 0x74, 0xb7, 0x99, 0xb7, 0xca, 0xbd,
	0x2f, 0xbd, 0xc5, 0xb0, 0x75, 0xe7, 0x98, 0x30, 0xa6, 0x22, 0x05, 0x4a, 0x9b, 0x93, 0x2b, 0xb6,
	0x07, 0x06, 0x50, 0x0d, 0x8d, 0x93, 0x77, 0x39, 0xca, 0xda, 0x0d, 0xed, 0x50, 0xf8, 0x56, 0x65,
	0xf3, 0x24, 0x1f, 0xcb, 0xcc, 0x79, 0xcb, 0xb2, 0x0f, 0xb5, 0x3f, 0x7c, 0xc5, 0xe0, 0xef, 0x85,
	0x85, 0x12, 0x0c, 0x6d, 0xcc, 0xad, 0xe0, 0x47, 0xda, 0xbd, 0xf7, 0x60, 0x05, 0xcc, 0x0b, 0xc5,
	0xc2, 0x35, 0x70, 0x46, 0x68, 0x36, 0x48, 0x22, 0xf4, 0x1e, 0x16, 0x11, 0x16, 0xfc, 0x05, 0xf1,
	0xd8, 0x8f, 0xc4, 0x3c, 0xad, 0xba, 0xde, 0xc7, 0x1d, 0x0a, 0x3b, 0x04, 0xc8, 0x56, 0x92, 0xde,
	0x5f, 0x62, 0xba, 0x62, 0x5b, 0xa1, 0x0f, 0x9a, 0x03, 0x5d, 0xb0, 0x54, 0xa5, 0xb6, 0x5b, 0x5f,
	0x22, 0xe1, 0xcd, 0xba, 0xc6, 0x3f, 0xc4, 0xdd, 0x45, 0x4e, 0xc0, 0x53, 0x55, 0x51, 0x36, 0xaf,
	0x40, 0x2a, 0x33, 0x48, 0xc5, 0x21, 0xe7, 0x27, 0xe8, 0x23, 0xdc, 0x3c, 0x89, 0x27, 0xca, 0x6a,
	0x27, 0x0c, 0x77, 0x45, 0x88, 0x5b, 0x2a, 0x02, 0xc4, 0xe0, 0xc2, 0xcc, 0x34, 0xe8, 0xe3, 0xe6,
	0xf2, 0xe1, 0x74, 0x64, 0xb8, 0x5d, 0x3b, 0x9b, 0x9f, 0xe0, 0x8e, 0x87, 0xf3, 0x08, 0xac, 0x3b,
	0x32, 0x77, 0x48, 0x7d, 0xda, 0x4c, 0x6a, 0xcd, 0x96, 0xbc, 0x45, 0xec, 0x9a, 0xbd, 0xb1, 0x3e,
	0xc3, 0x1d, 0x56, 0x56, 0x29, 0x12, 0xa3, 0x7e, 0x87, 0xd0, 0xe7, 0x2d, 0x45, 0x62, 0x4e, 0x82,
	0xc5, 0xe7, 0xaa, 0xb5, 0x79, 0xbf, 0xc0, 0x1d, 0x56, 0xef, 0x76, 0x6d, 0x07, 0x7d, 0x89, 0x3b,
	0x2e, 0xa1, 0xed, 0xda, 0x69, 0xfe, 0x0a, 0x77, 0x3c, 0xce, 0xb7, 0xc1, 0xda, 0x64, 0x8b, 0x39,
	0xed, 0xf8, 0xba, 0xb9, 0x1d, 0xab, 0xd5, 0x46, 0xb3, 0x9a, 0x11, 0x83, 0xcd, 0xc9, 0x46, 0x3e,
	0x55, 0xed, 0xdf, 0xb4, 0x20, 0x7a, 0xb1, 0x5a, 0xd6, 0x33, 0xe5, 0x7e, 0x04, 0x7a, 0xb3, 0x4f,
	0x55, 0x16, 0x44, 0x57, 0xd0, 0xb7, 0x6d, 0xa7, 0xe9, 0x68, 0x7e, 0x7f, 0xef, 0x4a, 0x75, 0x8c,
	0xa4, 0xba, 0x9c, 0x8e, 0x7c, 0xd7, 0xf2, 0x18, 0x49, 0xa5, 0x59, 0x0d, 0xd9, 0xa9, 0x2f, 0xfb,
	0xef, 0x71, 0xd7, 0x6d, 0x7f, 0x0b, 0x3c, 0x5e, 0x0a, 0xcc, 0x61, 0xf4, 0x43, 0x33, 0xa3, 0x15,
	0xa3, 0x35, 0x8b, 0xd0, 0x4b, 0xce, 0x9d, 0xf9, 0x23, 0xee, 0x70, 0x69, 0x5e, 0x07, 0xcb, 0x86,
	0x48, 0x2c, 0x2b, 0xf9, 0x49, 0x63, 0x51, 0x0d, 0x9b, 0xd2, 0x32, 0x3b, 0xd0, 0xd9, 0xa5, 0x33,
	0x7c, 0x13, 0x6c, 0x9c, 0x22, 0x0c, 0x35, 0xb0, 0x9f, 0x5b, 0xee, 0x83, 0x9a, 0x20, 0xe4, 0xc4,
	0xca, 0x53, 0x6d, 0x3a, 0xe8, 0xb4, 0xe8, 0x41, 0x4b, 0x1d, 0x98, 0x13, 0x65, 0xb5, 0xa9, 0xbc,
	0x29, 0xd9, 0x98, 0xe5, 0xe2, 0x5b, 0x1b, 0xfd, 0x82, 0xdb, 0x5d, 0x75, 0x07, 0xda, 0xbf, 0x62,
	0x65, 0x6e, 0x4a, 0x87, 0xd5, 0xaf, 0x2d, 0x59, 0x99, 0x6b, 0xd3, 0x62, 0x55, 0xee, 0xd4, 0xf2,
	0x42, 0x72, 0xc2, 0xfe, 0xd6, 0xb2, 0x87, 0xe5, 0xcd, 0x64, 0xc5, 0x2d, 0x17, 0xc1, 0xd4, 0x74,
	0xd0, 0xef, 0x2d, 0x17, 0x41, 0x6d, 0x30, 0xf0, 0x65, 0xe7, 0x03, 0xe9, 0x0f, 0xdc, 0xe9, 0x0b,
	0xe9, 0x46, 0xed, 0x33, 0xe3, 0x4f, 0xdc, 0xe9, 0x3b, 0xa3, 0x5a, 0x8c, 0xf9, 0x28, 0x23, 0x45,
	0x32, 0x40, 0x7f, 0xb5, 0xbd, 0xbb, 0xb0, 0x06, 0xec, 0x1e, 0x80, 0x8d, 0x84, 0xd6, 0x3e, 0xc3,
	0x27, 0x3f, 0x18, 0xdf, 0xbe, 0x1c, 0x53, 0x16, 0xdd, 0x2d, 0xed, 0x51, 0xbb, 0xdf, 0x94, 0xc7,
	0x67, 0xd4, 0x0f, 0xb8, 0x17, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x28, 0xca, 0xef, 0x27, 0x89,
	0x0e, 0x00, 0x00,
}
