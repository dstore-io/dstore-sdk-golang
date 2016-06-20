// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/mi_GetApplicationPartsTree_Pu.proto
// DO NOT EDIT!

/*
Package mi_GetApplicationPartsTree_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/mi_GetApplicationPartsTree_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package mi_GetApplicationPartsTree_Pu

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
	ApplicationPartTreeId            *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=application_part_tree_id,json=applicationPartTreeId" json:"application_part_tree_id,omitempty"`
	ApplicationPartTreeIdNull        bool                        `protobuf:"varint,1001,opt,name=application_part_tree_id_null,json=applicationPartTreeIdNull" json:"application_part_tree_id_null,omitempty"`
	GetRootPartsForApplicationId     *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=get_root_parts_for_application_id,json=getRootPartsForApplicationId" json:"get_root_parts_for_application_id,omitempty"`
	GetRootPartsForApplicationIdNull bool                        `protobuf:"varint,1002,opt,name=get_root_parts_for_application_id_null,json=getRootPartsForApplicationIdNull" json:"get_root_parts_for_application_id_null,omitempty"`
	IdsInOneId                       *dstore_values.IntegerValue `protobuf:"bytes,3,opt,name=ids_in_one_id,json=idsInOneId" json:"ids_in_one_id,omitempty"`
	IdsInOneIdNull                   bool                        `protobuf:"varint,1003,opt,name=ids_in_one_id_null,json=idsInOneIdNull" json:"ids_in_one_id_null,omitempty"`
	ReturnResultSet                  *dstore_values.IntegerValue `protobuf:"bytes,4,opt,name=return_result_set,json=returnResultSet" json:"return_result_set,omitempty"`
	ReturnResultSetNull              bool                        `protobuf:"varint,1004,opt,name=return_result_set_null,json=returnResultSetNull" json:"return_result_set_null,omitempty"`
	GetPartsTreeForGlobalUser        *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=get_parts_tree_for_global_user,json=getPartsTreeForGlobalUser" json:"get_parts_tree_for_global_user,omitempty"`
	GetPartsTreeForGlobalUserNull    bool                        `protobuf:"varint,1005,opt,name=get_parts_tree_for_global_user_null,json=getPartsTreeForGlobalUserNull" json:"get_parts_tree_for_global_user_null,omitempty"`
	OutputIntoOneId                  *dstore_values.IntegerValue `protobuf:"bytes,6,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull              bool                        `protobuf:"varint,1006,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
	OutputIntoTwoIds                 *dstore_values.BooleanValue `protobuf:"bytes,7,opt,name=output_into_two_ids,json=outputIntoTwoIds" json:"output_into_two_ids,omitempty"`
	OutputIntoTwoIdsNull             bool                        `protobuf:"varint,1007,opt,name=output_into_two_ids_null,json=outputIntoTwoIdsNull" json:"output_into_two_ids_null,omitempty"`
	MaxTreeLevel                     *dstore_values.IntegerValue `protobuf:"bytes,8,opt,name=max_tree_level,json=maxTreeLevel" json:"max_tree_level,omitempty"`
	MaxTreeLevelNull                 bool                        `protobuf:"varint,1008,opt,name=max_tree_level_null,json=maxTreeLevelNull" json:"max_tree_level_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetApplicationPartTreeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationPartTreeId
	}
	return nil
}

func (m *Parameters) GetGetRootPartsForApplicationId() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetRootPartsForApplicationId
	}
	return nil
}

func (m *Parameters) GetIdsInOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.IdsInOneId
	}
	return nil
}

func (m *Parameters) GetReturnResultSet() *dstore_values.IntegerValue {
	if m != nil {
		return m.ReturnResultSet
	}
	return nil
}

func (m *Parameters) GetGetPartsTreeForGlobalUser() *dstore_values.IntegerValue {
	if m != nil {
		return m.GetPartsTreeForGlobalUser
	}
	return nil
}

func (m *Parameters) GetOutputIntoOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputIntoOneId
	}
	return nil
}

func (m *Parameters) GetOutputIntoTwoIds() *dstore_values.BooleanValue {
	if m != nil {
		return m.OutputIntoTwoIds
	}
	return nil
}

func (m *Parameters) GetMaxTreeLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxTreeLevel
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
	SuccessorApplicationPartId *dstore_values.IntegerValue `protobuf:"bytes,10001,opt,name=successor_application_part_id,json=successorApplicationPartId" json:"successor_application_part_id,omitempty"`
	TreeLevel                  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=tree_level,json=treeLevel" json:"tree_level,omitempty"`
	SuccessorApplicationPart   *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=successor_application_part,json=successorApplicationPart" json:"successor_application_part,omitempty"`
	UserId                     *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	HasSuccessors              *dstore_values.IntegerValue `protobuf:"bytes,10005,opt,name=has_successors,json=hasSuccessors" json:"has_successors,omitempty"`
	SuccessorApplicPartTreeId  *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=successor_applic_part_tree_id,json=successorApplicPartTreeId" json:"successor_applic_part_tree_id,omitempty"`
	ApplicationPartTreeId      *dstore_values.IntegerValue `protobuf:"bytes,10007,opt,name=application_part_tree_id,json=applicationPartTreeId" json:"application_part_tree_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetSuccessorApplicationPartId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SuccessorApplicationPartId
	}
	return nil
}

func (m *Response_Row) GetTreeLevel() *dstore_values.IntegerValue {
	if m != nil {
		return m.TreeLevel
	}
	return nil
}

func (m *Response_Row) GetSuccessorApplicationPart() *dstore_values.StringValue {
	if m != nil {
		return m.SuccessorApplicationPart
	}
	return nil
}

func (m *Response_Row) GetUserId() *dstore_values.IntegerValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *Response_Row) GetHasSuccessors() *dstore_values.IntegerValue {
	if m != nil {
		return m.HasSuccessors
	}
	return nil
}

func (m *Response_Row) GetSuccessorApplicPartTreeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SuccessorApplicPartTreeId
	}
	return nil
}

func (m *Response_Row) GetApplicationPartTreeId() *dstore_values.IntegerValue {
	if m != nil {
		return m.ApplicationPartTreeId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.mi_GetApplicationPartsTree_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.mi_GetApplicationPartsTree_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.mi_GetApplicationPartsTree_Pu.Response.Row")
}

var fileDescriptor0 = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xeb, 0x6e, 0x13, 0x39,
	0x14, 0x56, 0x37, 0xcd, 0x65, 0xcf, 0x6e, 0x2f, 0x3b, 0xd9, 0xed, 0x4e, 0xd2, 0x6d, 0xd5, 0x6d,
	0x25, 0x84, 0x54, 0x69, 0x8a, 0xb8, 0x23, 0x10, 0x52, 0x41, 0x50, 0x02, 0x34, 0x94, 0x69, 0x8b,
	0x54, 0x7e, 0x60, 0x4d, 0x13, 0x13, 0x46, 0x9a, 0x8c, 0x23, 0xdb, 0xd3, 0xf0, 0x18, 0xdc, 0x79,
	0x26, 0x1e, 0x85, 0x72, 0xeb, 0x23, 0x70, 0x6c, 0x4f, 0x93, 0xcc, 0xa4, 0xb9, 0xd0, 0x3f, 0xad,
	0x1c, 0x9f, 0xef, 0xe2, 0xcf, 0xf6, 0xf1, 0xc0, 0x8d, 0xba, 0x90, 0x8c, 0xd3, 0x35, 0x1a, 0x36,
	0xfc, 0x90, 0xae, 0xb5, 0x38, 0xab, 0xd1, 0x7a, 0xc4, 0xa9, 0x58, 0x6b, 0xfa, 0x64, 0x83, 0xca,
	0xf5, 0x56, 0x2b, 0xf0, 0x6b, 0x9e, 0xf4, 0x59, 0xb8, 0xe5, 0x71, 0x29, 0x76, 0x38, 0xa5, 0x64,
	0x2b, 0x72, 0xb0, 0x50, 0x32, 0x6b, 0xd5, 0xa0, 0x1d, 0x83, 0x76, 0x86, 0x42, 0xca, 0xc5, 0x58,
	0xea, 0xc0, 0x0b, 0x22, 0x2a, 0x0c, 0x43, 0xb9, 0x94, 0xd4, 0xa7, 0x9c, 0x33, 0x1e, 0x4f, 0xcd,
	0x27, 0xa7, 0x9a, 0x54, 0x08, 0xaf, 0x41, 0xe3, 0xc9, 0x95, 0xf4, 0xa4, 0xf4, 0xfc, 0xf0, 0x39,
	0xe3, 0x4d, 0xad, 0x6b, 0x8a, 0x96, 0x8f, 0x0a, 0x00, 0x68, 0xc1, 0xc3, 0x59, 0xca, 0x85, 0xb5,
	0x03, 0xb6, 0xd7, 0xf5, 0x46, 0x5a, 0x68, 0x8e, 0x48, 0x65, 0xce, 0xaf, 0xdb, 0x13, 0x4b, 0x13,
	0x67, 0xff, 0x38, 0x3f, 0xef, 0xc4, 0x0b, 0x8a, 0x3d, 0xfa, 0xa1, 0xa4, 0x0d, 0xca, 0x9f, 0xa8,
	0x91, 0xfb, 0x8f, 0x97, 0x5c, 0x98, 0x5a, 0x57, 0xa5, 0x6e, 0xad, 0xc3, 0xc2, 0x20, 0x56, 0x12,
	0x46, 0x41, 0x60, 0x7f, 0xce, 0x23, 0x77, 0xc1, 0x2d, 0x9d, 0x08, 0xaf, 0x62, 0x85, 0x55, 0x87,
	0xff, 0x1b, 0x54, 0x12, 0xce, 0x98, 0xd4, 0x78, 0x41, 0x70, 0x25, 0xa4, 0x97, 0x15, 0x1d, 0xfe,
	0x36, 0xda, 0xe1, 0x7f, 0xc8, 0xe2, 0x22, 0x89, 0x8e, 0xfd, 0x2e, 0xe3, 0x3d, 0x3b, 0x81, 0x46,
	0x1f, 0xc3, 0x99, 0x91, 0x2a, 0xc6, 0xf1, 0xa1, 0x71, 0xbc, 0x34, 0x8c, 0x4e, 0x1b, 0xbf, 0x09,
	0x53, 0x7e, 0x5d, 0x10, 0x3f, 0x24, 0x2c, 0xd4, 0x31, 0x66, 0x46, 0x9b, 0x04, 0x44, 0x54, 0xc2,
	0x47, 0xa1, 0xca, 0x6e, 0x15, 0xac, 0x04, 0xde, 0xc8, 0x7f, 0x31, 0xf2, 0xd3, 0xdd, 0x42, 0x2d,
	0xb6, 0x01, 0x7f, 0x71, 0x2a, 0x23, 0x1e, 0x12, 0x3c, 0x9e, 0x51, 0x20, 0x89, 0xa0, 0xd2, 0x9e,
	0x1c, 0x2d, 0x38, 0x63, 0x50, 0xae, 0x06, 0x6d, 0x53, 0x69, 0x5d, 0x84, 0xb9, 0x3e, 0x22, 0xa3,
	0xfc, 0xd5, 0x28, 0x17, 0x53, 0x08, 0x2d, 0xff, 0x0c, 0x16, 0x55, 0x7c, 0x26, 0x39, 0xbd, 0xc1,
	0x2a, 0xbe, 0x46, 0xc0, 0xf6, 0xbd, 0x80, 0x44, 0x82, 0x72, 0x3b, 0x3b, 0xda, 0x4b, 0x09, 0x29,
	0x3a, 0x97, 0x02, 0x23, 0xdd, 0xd0, 0xf0, 0x5d, 0x44, 0x5b, 0x0f, 0x60, 0x65, 0x38, 0xbf, 0xb1,
	0xf8, 0xcd, 0x58, 0x5c, 0x18, 0x48, 0xa4, 0xcd, 0xde, 0x03, 0x8b, 0x45, 0xb2, 0x15, 0x49, 0xcc,
	0x56, 0xb2, 0xe3, 0xdd, 0xc9, 0x8d, 0x11, 0x96, 0x81, 0x55, 0x10, 0x65, 0xb6, 0xe8, 0x12, 0xfc,
	0xdb, 0xcf, 0x64, 0xac, 0x7c, 0x8f, 0xd3, 0x4a, 0x41, 0xb4, 0x81, 0xfb, 0x50, 0xec, 0x85, 0xc9,
	0x36, 0x43, 0x98, 0xb0, 0xf3, 0x27, 0x3a, 0xd8, 0x67, 0x2c, 0xa0, 0x5e, 0x68, 0x1c, 0xcc, 0x76,
	0xe9, 0x76, 0xda, 0xac, 0x52, 0x17, 0xd6, 0x15, 0xb0, 0x4f, 0xe0, 0x32, 0x1e, 0x7e, 0x18, 0x0f,
	0x7f, 0xa7, 0x41, 0xda, 0xc4, 0x3a, 0x4c, 0x37, 0xbd, 0x97, 0x26, 0xcc, 0x80, 0x1e, 0xd0, 0xc0,
	0x2e, 0x8c, 0x4e, 0xe0, 0x4f, 0x84, 0xa8, 0x50, 0x1f, 0x2a, 0x80, 0xe5, 0x40, 0x31, 0x49, 0x61,
	0x64, 0x8f, 0x8c, 0xec, 0x6c, 0x6f, 0xad, 0x92, 0x5c, 0x3e, 0xcc, 0x41, 0x01, 0xcf, 0x4d, 0x8b,
	0x85, 0x82, 0x5a, 0xe7, 0x20, 0xab, 0x1b, 0x5a, 0xdc, 0x5d, 0xca, 0x4e, 0xb2, 0x5d, 0x9a, 0x66,
	0x77, 0x47, 0xfd, 0x75, 0x4d, 0xa1, 0xb5, 0x07, 0xb3, 0xaa, 0x95, 0x91, 0x9e, 0x5e, 0x86, 0x17,
	0x3f, 0x83, 0x60, 0x27, 0x05, 0x4e, 0x77, 0xbc, 0x4d, 0x1c, 0x57, 0xba, 0x63, 0x77, 0xa6, 0x99,
	0xfc, 0xc1, 0xba, 0x0a, 0xf9, 0xb8, 0x85, 0xe2, 0x2d, 0x55, 0x8c, 0x8b, 0x7d, 0x8c, 0xa6, 0xc1,
	0x6e, 0x9a, 0xff, 0xee, 0x71, 0x39, 0x9e, 0xcc, 0x0c, 0x67, 0x6d, 0xbc, 0x6a, 0x0a, 0x75, 0xcd,
	0xf9, 0x85, 0x9e, 0xef, 0x1c, 0x47, 0xe1, 0xb8, 0xac, 0xed, 0x2a, 0x96, 0xf2, 0xa7, 0x49, 0xc8,
	0xe0, 0xc0, 0x9a, 0x83, 0x1c, 0x0e, 0xd5, 0xa9, 0x7c, 0x55, 0xc5, 0x74, 0xb2, 0x6e, 0x16, 0x87,
	0x78, 0xde, 0x08, 0x2c, 0x88, 0xa8, 0x56, 0x43, 0xe9, 0x54, 0x73, 0xd2, 0x8d, 0x15, 0xcb, 0x5f,
	0x57, 0x47, 0xef, 0x61, 0xb9, 0x43, 0x91, 0x32, 0x86, 0x02, 0xd7, 0x01, 0x7a, 0x0e, 0xc4, 0x9b,
	0x31, 0xd8, 0x7e, 0x97, 0x9d, 0xe3, 0xb0, 0x07, 0xe5, 0xc1, 0xee, 0xec, 0xb7, 0xd5, 0xe4, 0x3e,
	0xc7, 0x64, 0x42, 0x72, 0x3f, 0x6c, 0x18, 0x2e, 0x7b, 0x90, 0x33, 0xbc, 0x68, 0x79, 0x7d, 0xcb,
	0x71, 0x89, 0xef, 0xc6, 0x30, 0x95, 0x53, 0xc5, 0xb8, 0x9c, 0xdb, 0x30, 0xfd, 0xc2, 0x13, 0xa4,
	0x43, 0x2b, 0xec, 0xf7, 0x63, 0xa0, 0xa7, 0x10, 0xb3, 0xdd, 0x81, 0x60, 0x6f, 0xeb, 0x0b, 0x3d,
	0xf9, 0x3c, 0x7e, 0x18, 0x83, 0xb3, 0x94, 0x5a, 0x5a, 0xcf, 0x1b, 0xb9, 0x3b, 0xe4, 0xe5, 0xfd,
	0x58, 0x3d, 0xed, 0xd3, 0x7b, 0x6b, 0x17, 0xe6, 0x7d, 0x96, 0x3a, 0x8f, 0xdd, 0x2f, 0x98, 0xa7,
	0x97, 0x4f, 0xf7, 0x6d, 0xb3, 0x9f, 0xd3, 0x5f, 0x0f, 0x17, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x8c, 0x6c, 0xb3, 0x8b, 0x1c, 0x09, 0x00, 0x00,
}
