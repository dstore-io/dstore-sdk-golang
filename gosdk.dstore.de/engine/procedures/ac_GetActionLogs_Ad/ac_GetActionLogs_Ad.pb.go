// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/ac_GetActionLogs_Ad.proto
// DO NOT EDIT!

/*
Package ac_GetActionLogs_Ad is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/ac_GetActionLogs_Ad.proto

It has these top-level messages:
	Parameters
	Response
*/
package ac_GetActionLogs_Ad

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
	ActionId       *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=action_id,json=actionId" json:"action_id,omitempty"`
	ActionIdNull   bool                        `protobuf:"varint,1001,opt,name=action_id_null,json=actionIdNull" json:"action_id_null,omitempty"`
	OnlyPartNo     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=only_part_no,json=onlyPartNo" json:"only_part_no,omitempty"`
	OnlyPartNoNull bool                        `protobuf:"varint,1002,opt,name=only_part_no_null,json=onlyPartNoNull" json:"only_part_no_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetActionId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ActionId
	}
	return nil
}

func (m *Parameters) GetOnlyPartNo() *dstore_values.IntegerValue {
	if m != nil {
		return m.OnlyPartNo
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
	RowId            int32                         `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	SmallBody        *dstore_values.StringValue    `protobuf:"bytes,10001,opt,name=small_body,json=smallBody" json:"small_body,omitempty"`
	AlreadyRead      *dstore_values.BooleanValue   `protobuf:"bytes,10002,opt,name=already_read,json=alreadyRead" json:"already_read,omitempty"`
	PostingId        *dstore_values.IntegerValue   `protobuf:"bytes,10003,opt,name=posting_id,json=postingId" json:"posting_id,omitempty"`
	ReplyToPostingId *dstore_values.IntegerValue   `protobuf:"bytes,10004,opt,name=reply_to_posting_id,json=replyToPostingId" json:"reply_to_posting_id,omitempty"`
	Subject          *dstore_values.StringValue    `protobuf:"bytes,10005,opt,name=subject" json:"subject,omitempty"`
	EMailOfAuthor    *dstore_values.StringValue    `protobuf:"bytes,10006,opt,name=e_mail_of_author,json=eMailOfAuthor" json:"e_mail_of_author,omitempty"`
	AuthorPersonId   *dstore_values.IntegerValue   `protobuf:"bytes,10007,opt,name=author_person_id,json=authorPersonId" json:"author_person_id,omitempty"`
	PostDateChar     *dstore_values.StringValue    `protobuf:"bytes,10008,opt,name=post_date_char,json=postDateChar" json:"post_date_char,omitempty"`
	Visible          *dstore_values.IntegerValue   `protobuf:"bytes,10009,opt,name=visible" json:"visible,omitempty"`
	HasBinaries      *dstore_values.BooleanValue   `protobuf:"bytes,10010,opt,name=has_binaries,json=hasBinaries" json:"has_binaries,omitempty"`
	Author           *dstore_values.StringValue    `protobuf:"bytes,10011,opt,name=author" json:"author,omitempty"`
	PostDate         *dstore_values.TimestampValue `protobuf:"bytes,10012,opt,name=post_date,json=postDate" json:"post_date,omitempty"`
	Body             *dstore_values.StringValue    `protobuf:"bytes,10013,opt,name=body" json:"body,omitempty"`
	ForumId          *dstore_values.IntegerValue   `protobuf:"bytes,20002,opt,name=forum_id,json=forumId" json:"forum_id,omitempty"`
	Value3           *dstore_values.StringValue    `protobuf:"bytes,20003,opt,name=value3" json:"value3,omitempty"`
	Value1           *dstore_values.StringValue    `protobuf:"bytes,20005,opt,name=value1" json:"value1,omitempty"`
	Value2           *dstore_values.StringValue    `protobuf:"bytes,20006,opt,name=value2" json:"value2,omitempty"`
	MainPostingId    *dstore_values.IntegerValue   `protobuf:"bytes,20010,opt,name=main_posting_id,json=mainPostingId" json:"main_posting_id,omitempty"`
	HasSuccessors    *dstore_values.BooleanValue   `protobuf:"bytes,20014,opt,name=has_successors,json=hasSuccessors" json:"has_successors,omitempty"`
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

func (m *Response_Row) GetAlreadyRead() *dstore_values.BooleanValue {
	if m != nil {
		return m.AlreadyRead
	}
	return nil
}

func (m *Response_Row) GetPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.PostingId
	}
	return nil
}

func (m *Response_Row) GetReplyToPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReplyToPostingId
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

func (m *Response_Row) GetPostDateChar() *dstore_values.StringValue {
	if m != nil {
		return m.PostDateChar
	}
	return nil
}

func (m *Response_Row) GetVisible() *dstore_values.IntegerValue {
	if m != nil {
		return m.Visible
	}
	return nil
}

func (m *Response_Row) GetHasBinaries() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasBinaries
	}
	return nil
}

func (m *Response_Row) GetAuthor() *dstore_values.StringValue {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Response_Row) GetPostDate() *dstore_values.TimestampValue {
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

func (m *Response_Row) GetMainPostingId() *dstore_values.IntegerValue {
	if m != nil {
		return m.MainPostingId
	}
	return nil
}

func (m *Response_Row) GetHasSuccessors() *dstore_values.BooleanValue {
	if m != nil {
		return m.HasSuccessors
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.ac_GetActionLogs_Ad.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.ac_GetActionLogs_Ad.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.ac_GetActionLogs_Ad.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/ac_GetActionLogs_Ad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x95, 0xdb, 0x6e, 0x13, 0x39,
	0x18, 0xc7, 0x95, 0x4d, 0x9b, 0x83, 0x9b, 0xa6, 0x59, 0x57, 0x5a, 0x79, 0x53, 0xed, 0x6a, 0xb7,
	0xab, 0x95, 0x10, 0x17, 0x13, 0x9a, 0x52, 0x84, 0x40, 0x08, 0xa5, 0xb4, 0xa0, 0x0a, 0x1a, 0xaa,
	0xa1, 0x42, 0x82, 0x1b, 0xcb, 0xc9, 0xb8, 0xe9, 0xc0, 0x64, 0x1c, 0xd9, 0x4e, 0xab, 0xbe, 0x05,
	0xe7, 0x63, 0xb9, 0x01, 0xc1, 0x05, 0x17, 0x48, 0xbc, 0x05, 0x8f, 0xc0, 0x2b, 0xc0, 0x53, 0xf0,
	0x79, 0x3c, 0x39, 0x52, 0x34, 0x73, 0x33, 0x91, 0xf3, 0xfd, 0x7f, 0x9f, 0xbf, 0xe3, 0x0c, 0xaa,
	0x7b, 0x4a, 0x0b, 0xc9, 0x6b, 0x3c, 0xec, 0xf8, 0x21, 0xaf, 0xf5, 0xa4, 0x68, 0x73, 0xaf, 0x2f,
	0xb9, 0xaa, 0xb1, 0x36, 0xbd, 0xc6, 0x75, 0xa3, 0xad, 0x7d, 0x11, 0xde, 0x10, 0x1d, 0x45, 0x1b,
	0x9e, 0x03, 0x66, 0x2d, 0xf0, 0xbf, 0x96, 0x71, 0x2c, 0xe3, 0x9c, 0x20, 0xac, 0x2e, 0xc6, 0x6e,
	0x0f, 0x58, 0xd0, 0xe7, 0xca, 0x72, 0xd5, 0x3f, 0x27, 0xef, 0xe2, 0x52, 0x0a, 0x19, 0x9b, 0x96,
	0x26, 0x4d, 0x5d, 0xae, 0x14, 0xeb, 0xf0, 0xd8, 0xf8, 0xdf, 0xb4, 0x51, 0x33, 0x3f, 0xdc, 0x13,
	0xb2, 0xcb, 0xcc, 0x95, 0x56, 0xb4, 0xfc, 0x35, 0x83, 0xd0, 0x0e, 0x93, 0x0c, 0xac, 0x5c, 0x2a,
	0x7c, 0x1e, 0x15, 0x59, 0x14, 0x11, 0xf5, 0x3d, 0x92, 0xf9, 0x27, 0x73, 0x6a, 0xae, 0xbe, 0xe4,
	0xc4, 0x71, 0xc7, 0x41, 0xf9, 0xa1, 0xe6, 0x1d, 0x2e, 0x6f, 0x9b, 0x93, 0x5b, 0xb0, 0xea, 0x2d,
	0x0f, 0xff, 0x8f, 0xca, 0x43, 0x92, 0x86, 0xfd, 0x20, 0x20, 0xdf, 0xf2, 0xc0, 0x17, 0xdc, 0xd2,
	0x40, 0xd2, 0x84, 0x3f, 0xf1, 0x25, 0x54, 0x12, 0x61, 0x70, 0x44, 0x7b, 0x4c, 0x6a, 0x1a, 0x0a,
	0xf2, 0x5b, 0xf2, 0x1d, 0xc8, 0x00, 0x10, 0xa3, 0x6e, 0x0a, 0x7c, 0x1a, 0xfd, 0x3e, 0x8e, 0xdb,
	0x8b, 0xbe, 0xdb, 0x8b, 0xca, 0x23, 0x9d, 0xb9, 0x6a, 0xf9, 0xcb, 0x1c, 0x2a, 0xb8, 0x5c, 0xf5,
	0x44, 0xa8, 0x38, 0x3e, 0x83, 0x66, 0xa3, 0xc2, 0xc5, 0x49, 0x55, 0x9d, 0xc9, 0x66, 0xd8, 0xa2,
	0x6e, 0x9a, 0xa7, 0x6b, 0x85, 0xf8, 0x0e, 0xaa, 0x98, 0x92, 0xd1, 0xb1, 0x9a, 0x41, 0xb4, 0x59,
	0x80, 0x9d, 0x29, 0x78, 0xba, 0xb2, 0xdb, 0x70, 0xde, 0x1a, 0x9d, 0xdd, 0x85, 0xee, 0xe4, 0x1f,
	0x50, 0xe5, 0x7c, 0xdc, 0x2a, 0x92, 0x8d, 0x3c, 0xfe, 0xfd, 0x93, 0x47, 0xdb, 0xc8, 0x6d, 0xfb,
	0xeb, 0x0e, 0xe4, 0xb8, 0x81, 0xb2, 0x52, 0x1c, 0x92, 0x99, 0x88, 0xaa, 0x39, 0x89, 0x13, 0xe5,
	0x0c, 0x0a, 0xe0, 0xb8, 0xe2, 0xd0, 0x35, 0x6c, 0xf5, 0x73, 0x11, 0x65, 0xe1, 0x80, 0xff, 0x40,
	0x39, 0x38, 0x9a, 0x3e, 0x3f, 0x68, 0x42, 0x4d, 0x66, 0xdd, 0x59, 0x38, 0x42, 0x23, 0x2f, 0x20,
	0xa4, 0xba, 0x2c, 0x08, 0x68, 0x4b, 0x78, 0x47, 0xe4, 0x61, 0x73, 0xb2, 0x5e, 0x71, 0x83, 0x94,
	0x96, 0x7e, 0xd8, 0xb1, 0xfd, 0x29, 0x46, 0xf2, 0x75, 0x50, 0xe3, 0xcb, 0xa8, 0xc4, 0x02, 0xc9,
	0x99, 0x77, 0x44, 0xcd, 0x93, 0x3c, 0x6a, 0x9e, 0xd8, 0xde, 0x96, 0x10, 0x01, 0x67, 0xa1, 0xc5,
	0xe7, 0x62, 0xc2, 0x85, 0x07, 0xbe, 0x88, 0x50, 0x4f, 0x28, 0x0d, 0xbe, 0x4d, 0x60, 0x8f, 0x9b,
	0xc9, 0xd3, 0x51, 0x8c, 0xf5, 0x10, 0xf9, 0x75, 0xb4, 0x28, 0x79, 0x0f, 0xa6, 0x43, 0x0b, 0x3a,
	0xe6, 0xe5, 0x49, 0x0a, 0x2f, 0x95, 0x08, 0xdc, 0x15, 0x3b, 0x43, 0x67, 0x6b, 0x28, 0xaf, 0xfa,
	0xad, 0x7b, 0xbc, 0xad, 0xc9, 0xd3, 0xe4, 0x1a, 0x0c, 0xb4, 0x78, 0x03, 0x55, 0x38, 0xed, 0x32,
	0x3f, 0xa0, 0x62, 0x8f, 0xb2, 0xbe, 0xde, 0x87, 0x91, 0x7b, 0x96, 0xcc, 0xcf, 0xf3, 0x6d, 0x60,
	0x6e, 0xee, 0x35, 0x22, 0x02, 0x5f, 0x45, 0x15, 0xcb, 0xd2, 0x1e, 0x6c, 0xa5, 0xdd, 0xc6, 0xe7,
	0x29, 0xd2, 0x28, 0x5b, 0x6a, 0x27, 0x82, 0x20, 0x89, 0x06, 0x2a, 0x9b, 0x42, 0x50, 0x8f, 0x69,
	0x4e, 0xdb, 0xfb, 0x4c, 0x92, 0x17, 0xc9, 0xb1, 0x94, 0x0c, 0xb2, 0x01, 0xc4, 0x15, 0x00, 0xf0,
	0x39, 0x94, 0x3f, 0xf0, 0x95, 0xdf, 0x0a, 0x38, 0x79, 0x99, 0x22, 0x82, 0x81, 0xd8, 0x8c, 0xc2,
	0x3e, 0x53, 0xb4, 0xe5, 0x87, 0x4c, 0xfa, 0x5c, 0x91, 0x57, 0x69, 0x46, 0x01, 0x88, 0xf5, 0x18,
	0xc0, 0xab, 0x28, 0x17, 0xd7, 0xef, 0x75, 0x72, 0xcc, 0xb1, 0x14, 0xe6, 0xa7, 0x38, 0x4c, 0x98,
	0x1c, 0x5b, 0xee, 0xaf, 0x29, 0x4e, 0xfb, 0xb0, 0x4d, 0x9a, 0x75, 0x7b, 0xf1, 0x2b, 0x6c, 0x90,
	0x2e, 0xae, 0xa1, 0x99, 0x68, 0xe6, 0xdf, 0x24, 0xdf, 0x17, 0x09, 0x61, 0x8f, 0x0b, 0xb0, 0xd4,
	0xfd, 0xae, 0x69, 0xcf, 0xdb, 0xe3, 0x14, 0xaf, 0xcb, 0x7c, 0x24, 0x87, 0xc6, 0x9c, 0x45, 0xb9,
	0xc8, 0xbe, 0x4a, 0xde, 0x1d, 0x67, 0x92, 0xb3, 0xb3, 0xda, 0x21, 0xb5, 0x42, 0xde, 0xa7, 0xa6,
	0x56, 0x86, 0x54, 0x9d, 0x7c, 0x48, 0x4d, 0xd5, 0xf1, 0x26, 0x5a, 0x80, 0x31, 0x0e, 0xc7, 0x17,
	0xe9, 0x63, 0x9a, 0x14, 0xe7, 0x0d, 0x35, 0x5a, 0xa3, 0x0d, 0x54, 0x36, 0x63, 0xa0, 0xfa, 0xed,
	0x36, 0xbc, 0xc2, 0x84, 0x54, 0xe4, 0xd3, 0x2f, 0xbc, 0x4c, 0x4c, 0xc2, 0x3c, 0x40, 0xb7, 0x86,
	0xcc, 0xfa, 0x2e, 0x5a, 0xf2, 0xc5, 0xd4, 0xdb, 0x6e, 0xf4, 0xcd, 0xbd, 0xbb, 0xd6, 0x11, 0xca,
	0xbb, 0x3f, 0xb0, 0x7b, 0x29, 0x3f, 0xcb, 0xad, 0x5c, 0xf4, 0x09, 0x5c, 0xfd, 0x11, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0x29, 0x83, 0xc3, 0xcd, 0x07, 0x00, 0x00,
}