// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/fo_GetPosting_Pu.proto
// DO NOT EDIT!

/*
Package fo_GetPosting_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/fo_GetPosting_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package fo_GetPosting_Pu

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
	PersonIdentificationValues     *dstore_values.StringValue  `protobuf:"bytes,1,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull bool                        `protobuf:"varint,1001,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	PersonTypeId                   *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=person_type_id,json=personTypeId" json:"person_type_id,omitempty"`
	PersonTypeIdNull               bool                        `protobuf:"varint,1002,opt,name=person_type_id_null,json=personTypeIdNull" json:"person_type_id_null,omitempty"`
	UniqueId                       *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                   bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PostingId                      *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=posting_id,json=postingId" json:"posting_id,omitempty"`
	PostingIdNull                  bool                        `protobuf:"varint,1004,opt,name=posting_id_null,json=postingIdNull" json:"posting_id_null,omitempty"`
	AdditionalInfos                *dstore_values.BooleanValue `protobuf:"bytes,5,opt,name=additional_infos,json=additionalInfos" json:"additional_infos,omitempty"`
	AdditionalInfosNull            bool                        `protobuf:"varint,1005,opt,name=additional_infos_null,json=additionalInfosNull" json:"additional_infos_null,omitempty"`
	SeparatorInIdentVals           *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull       bool                        `protobuf:"varint,1006,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetPersonTypeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PersonTypeId
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingId
	}
	return nil
}

func (m *Parameters) GetAdditionalInfos() *dstore_values.BooleanValue {
	if m != nil {
		return m.AdditionalInfos
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
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
	RowId            int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	SmallBody        *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=small_body,json=smallBody" json:"small_body,omitempty"`
	ForumId          *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	Value3           *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value3" json:"value3,omitempty"`
	ReplyToPostingId *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=reply_to_posting_id,json=replyToPostingId" json:"reply_to_posting_id,omitempty"`
	Value1           *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=value1" json:"value1,omitempty"`
	Value2           *dstore_values.StringValue  `protobuf:"bytes,10006,opt,name=value2" json:"value2,omitempty"`
	Subject          *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=subject" json:"subject,omitempty"`
	EMailOfAuthor    *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=e_mail_of_author,json=eMailOfAuthor" json:"e_mail_of_author,omitempty"`
	AuthorPersonId   *dstore_values.IntegerValue `protobuf:"bytes,10009,opt,name=author_person_id,json=authorPersonId" json:"author_person_id,omitempty"`
	MainPostingId    *dstore_values.IntegerValue `protobuf:"bytes,10010,opt,name=main_posting_id,json=mainPostingId" json:"main_posting_id,omitempty"`
	HasBinaries      *dstore_values.BooleanValue `protobuf:"bytes,10011,opt,name=has_binaries,json=hasBinaries" json:"has_binaries,omitempty"`
	Visible          *dstore_values.IntegerValue `protobuf:"bytes,10012,opt,name=visible" json:"visible,omitempty"`
	Author           *dstore_values.StringValue  `protobuf:"bytes,10013,opt,name=author" json:"author,omitempty"`
	HasSuccessors    *dstore_values.BooleanValue `protobuf:"bytes,10014,opt,name=has_successors,json=hasSuccessors" json:"has_successors,omitempty"`
	PostDate         *dstore_values.StringValue  `protobuf:"bytes,10015,opt,name=post_date,json=postDate" json:"post_date,omitempty"`
	Body             *dstore_values.StringValue  `protobuf:"bytes,10016,opt,name=body" json:"body,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetSmallBody() *dstore_values.StringValue {
	if m != nil {
		return m.SmallBody
	}
	return nil
}

func (m *Response_Row) GetForumId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ForumId
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func (m *Response_Row) GetReplyToPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReplyToPostingId
	}
	return nil
}

func (m *Response_Row) GetValue1() *dstore_values.StringValue {
	if m != nil {
		return m.Value1
	}
	return nil
}

func (m *Response_Row) GetValue2() *dstore_values.StringValue {
	if m != nil {
		return m.Value2
	}
	return nil
}

func (m *Response_Row) GetSubject() *dstore_values.StringValue {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Response_Row) GetEMailOfAuthor() *dstore_values.StringValue {
	if m != nil {
		return m.EMailOfAuthor
	}
	return nil
}

func (m *Response_Row) GetAuthorPersonId() *dstore_values.IntegerValue {
	if m != nil {
		return m.AuthorPersonId
	}
	return nil
}

func (m *Response_Row) GetMainPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.MainPostingId
	}
	return nil
}

func (m *Response_Row) GetHasBinaries() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasBinaries
	}
	return nil
}

func (m *Response_Row) GetVisible() *dstore_values.IntegerValue {
	if m != nil {
		return m.Visible
	}
	return nil
}

func (m *Response_Row) GetAuthor() *dstore_values.StringValue {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Response_Row) GetHasSuccessors() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasSuccessors
	}
	return nil
}

func (m *Response_Row) GetPostDate() *dstore_values.StringValue {
	if m != nil {
		return m.PostDate
	}
	return nil
}

func (m *Response_Row) GetBody() *dstore_values.StringValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.fo_GetPosting_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.fo_GetPosting_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.fo_GetPosting_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x96, 0x69, 0x6f, 0xf3, 0x44,
	0x10, 0xc7, 0x55, 0xd2, 0x5c, 0xd3, 0xe6, 0xd0, 0x86, 0xc3, 0xa4, 0x55, 0x05, 0x45, 0x08, 0x90,
	0x90, 0x43, 0x1b, 0x71, 0x88, 0x17, 0x54, 0x2d, 0xe5, 0x08, 0xa8, 0x21, 0x98, 0x0a, 0x09, 0x84,
	0xb4, 0x72, 0xe2, 0x4d, 0xba, 0xc8, 0xf1, 0x06, 0xaf, 0xdd, 0xaa, 0xdf, 0x82, 0xfb, 0xbe, 0xbe,
	0x1d, 0xe2, 0xfc, 0x02, 0xbc, 0x78, 0x9e, 0xd9, 0x5d, 0xe7, 0xf2, 0xd3, 0xd6, 0x7e, 0xd3, 0x74,
	0xb3, 0xff, 0xdf, 0xfc, 0x67, 0xc7, 0x33, 0x1b, 0x43, 0xc7, 0x93, 0x91, 0x08, 0x59, 0x87, 0x05,
	0x13, 0x1e, 0xb0, 0xce, 0x2c, 0x14, 0x23, 0xe6, 0xc5, 0x21, 0x93, 0x9d, 0xb1, 0xa0, 0x6f, 0xb1,
	0x68, 0x20, 0x64, 0xc4, 0x83, 0x09, 0x1d, 0xc4, 0x36, 0xee, 0x45, 0x82, 0xec, 0x19, 0xc0, 0x36,
	0x80, 0x9d, 0x56, 0xb5, 0x5b, 0x49, 0xc0, 0x4b, 0xd7, 0x8f, 0x99, 0x34, 0x50, 0xfb, 0xf1, 0x75,
	0x17, 0x16, 0x86, 0x22, 0x4c, 0xb6, 0x76, 0xd6, 0xb7, 0xa6, 0x4c, 0x4a, 0x77, 0xc2, 0x92, 0xcd,
	0xa7, 0xd2, 0x9b, 0x91, 0xcb, 0x83, 0xb1, 0x08, 0xa7, 0x6e, 0xc4, 0x45, 0x60, 0x44, 0xfb, 0xf7,
	0x8a, 0x00, 0x03, 0x37, 0x74, 0x71, 0x97, 0x85, 0x92, 0x7c, 0x02, 0xbb, 0x33, 0xfc, 0x14, 0x01,
	0xe5, 0x1e, 0x0b, 0x22, 0x3e, 0xe6, 0x23, 0xad, 0xa6, 0x26, 0x23, 0x6b, 0xe3, 0x89, 0x8d, 0x67,
	0xb7, 0x0e, 0xdb, 0x76, 0x72, 0x8e, 0x24, 0x4f, 0x19, 0x85, 0x78, 0x80, 0x0f, 0xd5, 0xc2, 0x69,
	0x1b, 0xbe, 0xb7, 0x86, 0xeb, 0x2d, 0x49, 0xde, 0x81, 0x27, 0xef, 0x8a, 0x4e, 0x83, 0xd8, 0xf7,
	0xad, 0x3f, 0xcb, 0xe8, 0x51, 0x71, 0xf6, 0x6e, 0x8f, 0xd3, 0x47, 0x19, 0x39, 0x86, 0x7a, 0x12,
	0x2b, 0xba, 0x9e, 0x31, 0x0c, 0x68, 0x3d, 0xa4, 0x73, 0xdb, 0x49, 0xe5, 0xc6, 0x83, 0x88, 0x4d,
	0x58, 0x68, 0x92, 0xdb, 0x36, 0xc8, 0x39, 0x12, 0x3d, 0x8f, 0xd8, 0xd0, 0x5a, 0x0f, 0x61, 0x12,
	0xf8, 0xcb, 0x24, 0xd0, 0x5c, 0xd5, 0x6a, 0xcb, 0x97, 0xa1, 0x1a, 0x07, 0xfc, 0xb3, 0x58, 0xbb,
	0x15, 0x32, 0x2b, 0x51, 0x31, 0x62, 0x34, 0x7a, 0x1a, 0xea, 0x0b, 0xd0, 0x78, 0xfc, 0x6d, 0x3c,
	0xb6, 0xe7, 0x12, 0x1d, 0xff, 0x55, 0x80, 0x59, 0xd2, 0x0b, 0x68, 0xb0, 0x99, 0x7d, 0x9c, 0x6a,
	0x22, 0x47, 0x8b, 0x67, 0xa0, 0xb1, 0x64, 0x8d, 0xc7, 0x3f, 0xc6, 0xa3, 0xb6, 0x10, 0x69, 0x93,
	0x37, 0xa1, 0xe9, 0x7a, 0x1e, 0x57, 0xd5, 0x74, 0x7d, 0xaa, 0x1a, 0x42, 0x5a, 0xc5, 0x1b, 0xad,
	0x86, 0x42, 0xf8, 0xcc, 0x35, 0x35, 0x77, 0x1a, 0x4b, 0xa8, 0xa7, 0x18, 0xd2, 0x85, 0x47, 0xd2,
	0x71, 0x8c, 0xed, 0xbf, 0xc6, 0xb6, 0x95, 0x02, 0xb4, 0xf9, 0xfb, 0xf0, 0x98, 0x64, 0x33, 0x6c,
	0x37, 0xf4, 0x41, 0xc6, 0xb4, 0x81, 0x7a, 0xfa, 0xd2, 0x2a, 0x65, 0xd6, 0xf3, 0xe1, 0x05, 0xda,
	0x33, 0x6d, 0x81, 0x5f, 0x4b, 0x72, 0x04, 0xbb, 0xb7, 0x84, 0x34, 0xe9, 0xfc, 0x67, 0xd2, 0xb1,
	0x6e, 0x82, 0x55, 0x4e, 0xfb, 0xff, 0x57, 0xa1, 0xe2, 0x30, 0x39, 0x13, 0x81, 0x64, 0xe4, 0x05,
	0x28, 0xea, 0xf9, 0x4a, 0x37, 0x7a, 0x32, 0xb0, 0x66, 0xf6, 0xde, 0x50, 0x7f, 0x1d, 0x23, 0x24,
	0x1f, 0x41, 0x53, 0x4d, 0x16, 0x5d, 0x19, 0x2d, 0xec, 0xc4, 0x02, 0xc2, 0x76, 0x0a, 0x4e, 0x0f,
	0xe0, 0x19, 0xae, 0x7b, 0xcb, 0xb5, 0xd3, 0x98, 0xae, 0x7f, 0x41, 0x5e, 0x81, 0x72, 0x32, 0xd1,
	0xd8, 0x6d, 0x2a, 0xe2, 0xde, 0x03, 0x11, 0xcd, 0xbc, 0x9f, 0x99, 0x4f, 0x67, 0x2e, 0x27, 0xaf,
	0x41, 0x21, 0x14, 0x57, 0xd8, 0x42, 0x8a, 0x7a, 0xde, 0xbe, 0xfb, 0xd6, 0xb1, 0xe7, 0xa7, 0xb7,
	0x1d, 0x71, 0xe5, 0x28, 0xb0, 0xfd, 0x47, 0x19, 0x0a, 0xb8, 0x20, 0x8f, 0x42, 0x09, 0x97, 0xaa,
	0x1b, 0x3f, 0xef, 0x63, 0x41, 0x8a, 0x4e, 0x11, 0x97, 0xd8, 0x6d, 0xd8, 0xa9, 0x72, 0xea, 0xfa,
	0x3e, 0x1d, 0x0a, 0xef, 0xda, 0xfa, 0xa2, 0x9f, 0xf9, 0xec, 0xaa, 0x5a, 0x7e, 0x82, 0x6a, 0x9c,
	0xa2, 0x0a, 0x1e, 0x31, 0x9e, 0xaa, 0xa8, 0x5f, 0xf6, 0xb3, 0x9b, 0xbc, 0xac, 0xd5, 0x68, 0xda,
	0x85, 0x92, 0xde, 0xef, 0x5a, 0x5f, 0x65, 0x1b, 0x26, 0x52, 0xf2, 0x2e, 0xb4, 0x42, 0x36, 0xf3,
	0xaf, 0x69, 0x24, 0xe8, 0xca, 0x70, 0x7d, 0x9d, 0xc3, 0xb8, 0xa9, 0xc1, 0x73, 0x31, 0x58, 0x0c,
	0xd9, 0x3c, 0x83, 0x03, 0xeb, 0x9b, 0xbc, 0x19, 0x1c, 0x2c, 0xa0, 0x43, 0xeb, 0xdb, 0xbc, 0xd0,
	0x21, 0x79, 0x11, 0xca, 0x32, 0x1e, 0x7e, 0xca, 0x46, 0x91, 0xf5, 0x5d, 0x36, 0x35, 0xd7, 0x92,
	0x53, 0x68, 0x32, 0x3a, 0x75, 0xb9, 0x4f, 0xc5, 0x98, 0xba, 0x71, 0x74, 0x81, 0x9d, 0xfc, 0x7d,
	0x36, 0x5f, 0x63, 0x67, 0xc8, 0xbc, 0x37, 0x3e, 0xd6, 0x84, 0xbe, 0x22, 0xf4, 0x7f, 0x74, 0x71,
	0x5b, 0x5b, 0x3f, 0xe4, 0x28, 0x58, 0xdd, 0x50, 0x83, 0xe4, 0xe2, 0xc6, 0x6c, 0x1a, 0x98, 0x4b,
	0xb0, 0x5a, 0xf7, 0x1f, 0x73, 0x84, 0xa9, 0x29, 0x68, 0x59, 0xf4, 0x23, 0xd8, 0xbe, 0x70, 0x25,
	0x1d, 0xf2, 0xc0, 0x0d, 0x39, 0xfe, 0x04, 0xfd, 0xd4, 0xcf, 0xbe, 0xad, 0xb6, 0x90, 0x38, 0x49,
	0x00, 0xf2, 0x12, 0x94, 0x2f, 0xb9, 0xe4, 0x43, 0x9f, 0x59, 0x3f, 0xe7, 0xe9, 0xb7, 0x44, 0xac,
	0x1e, 0x5c, 0x52, 0xc2, 0x5f, 0x72, 0x3c, 0x38, 0x23, 0x25, 0xaf, 0x43, 0x5d, 0x65, 0x2b, 0xe3,
	0xd1, 0x08, 0x67, 0x51, 0x84, 0xd2, 0xfa, 0x35, 0x47, 0xbe, 0x35, 0x64, 0x3e, 0x58, 0x20, 0x38,
	0xf8, 0xfa, 0x66, 0xa7, 0x9e, 0x1b, 0x31, 0xeb, 0xb7, 0x6c, 0xf3, 0x8a, 0x52, 0x9f, 0xa2, 0x98,
	0x74, 0x60, 0x53, 0x8f, 0xe4, 0xef, 0xd9, 0x90, 0x16, 0x9e, 0xbc, 0x0d, 0x3b, 0x5c, 0xa4, 0x2e,
	0x88, 0xe5, 0x7b, 0xcc, 0xc7, 0xcf, 0xe5, 0x7e, 0xc3, 0x19, 0x96, 0xf4, 0x0b, 0x45, 0xf7, 0x7e,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x1b, 0x4a, 0xa7, 0x15, 0x09, 0x00, 0x00,
}
