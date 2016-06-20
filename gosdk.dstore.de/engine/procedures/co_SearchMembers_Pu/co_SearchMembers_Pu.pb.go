// Code generated by protoc-gen-go.
// source: dstore/engine/procedures/co_SearchMembers_Pu.proto
// DO NOT EDIT!

/*
Package co_SearchMembers_Pu is a generated protocol buffer package.

It is generated from these files:
	dstore/engine/procedures/co_SearchMembers_Pu.proto

It has these top-level messages:
	Parameters
	Response
*/
package co_SearchMembers_Pu

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
	Rowcount                          *dstore_values.IntegerValue `protobuf:"bytes,1,opt,name=rowcount" json:"rowcount,omitempty"`
	RowcountNull                      bool                        `protobuf:"varint,1001,opt,name=rowcount_null,json=rowcountNull" json:"rowcount_null,omitempty"`
	Result                            *dstore_values.IntegerValue `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	ResultNull                        bool                        `protobuf:"varint,1002,opt,name=result_null,json=resultNull" json:"result_null,omitempty"`
	UniqueId                          *dstore_values.StringValue  `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	UniqueIdNull                      bool                        `protobuf:"varint,1003,opt,name=unique_id_null,json=uniqueIdNull" json:"unique_id_null,omitempty"`
	PersonIdentificationValues        *dstore_values.StringValue  `protobuf:"bytes,4,opt,name=person_identification_values,json=personIdentificationValues" json:"person_identification_values,omitempty"`
	PersonIdentificationValuesNull    bool                        `protobuf:"varint,1004,opt,name=person_identification_values_null,json=personIdentificationValuesNull" json:"person_identification_values_null,omitempty"`
	CommunityId                       *dstore_values.IntegerValue `protobuf:"bytes,5,opt,name=community_id,json=communityId" json:"community_id,omitempty"`
	CommunityIdNull                   bool                        `protobuf:"varint,1005,opt,name=community_id_null,json=communityIdNull" json:"community_id_null,omitempty"`
	SearchString                      *dstore_values.StringValue  `protobuf:"bytes,6,opt,name=search_string,json=searchString" json:"search_string,omitempty"`
	SearchStringNull                  bool                        `protobuf:"varint,1006,opt,name=search_string_null,json=searchStringNull" json:"search_string_null,omitempty"`
	MaxNumberOfRows                   *dstore_values.IntegerValue `protobuf:"bytes,7,opt,name=max_number_of_rows,json=maxNumberOfRows" json:"max_number_of_rows,omitempty"`
	MaxNumberOfRowsNull               bool                        `protobuf:"varint,1007,opt,name=max_number_of_rows_null,json=maxNumberOfRowsNull" json:"max_number_of_rows_null,omitempty"`
	CaseSensitive                     *dstore_values.BooleanValue `protobuf:"bytes,8,opt,name=case_sensitive,json=caseSensitive" json:"case_sensitive,omitempty"`
	CaseSensitiveNull                 bool                        `protobuf:"varint,1008,opt,name=case_sensitive_null,json=caseSensitiveNull" json:"case_sensitive_null,omitempty"`
	OutputCharacteristicId1           *dstore_values.IntegerValue `protobuf:"bytes,9,opt,name=output_characteristic_id1,json=outputCharacteristicId1" json:"output_characteristic_id1,omitempty"`
	OutputCharacteristicId1Null       bool                        `protobuf:"varint,1009,opt,name=output_characteristic_id1_null,json=outputCharacteristicId1Null" json:"output_characteristic_id1_null,omitempty"`
	OutputCharacteristicId2           *dstore_values.IntegerValue `protobuf:"bytes,10,opt,name=output_characteristic_id2,json=outputCharacteristicId2" json:"output_characteristic_id2,omitempty"`
	OutputCharacteristicId2Null       bool                        `protobuf:"varint,1010,opt,name=output_characteristic_id2_null,json=outputCharacteristicId2Null" json:"output_characteristic_id2_null,omitempty"`
	OutputCharacteristicId3           *dstore_values.IntegerValue `protobuf:"bytes,11,opt,name=output_characteristic_id3,json=outputCharacteristicId3" json:"output_characteristic_id3,omitempty"`
	OutputCharacteristicId3Null       bool                        `protobuf:"varint,1011,opt,name=output_characteristic_id3_null,json=outputCharacteristicId3Null" json:"output_characteristic_id3_null,omitempty"`
	CommunityBinaryCategoryId         *dstore_values.IntegerValue `protobuf:"bytes,12,opt,name=community_binary_category_id,json=communityBinaryCategoryId" json:"community_binary_category_id,omitempty"`
	CommunityBinaryCategoryIdNull     bool                        `protobuf:"varint,1012,opt,name=community_binary_category_id_null,json=communityBinaryCategoryIdNull" json:"community_binary_category_id_null,omitempty"`
	ExactMatching                     *dstore_values.BooleanValue `protobuf:"bytes,13,opt,name=exact_matching,json=exactMatching" json:"exact_matching,omitempty"`
	ExactMatchingNull                 bool                        `protobuf:"varint,1013,opt,name=exact_matching_null,json=exactMatchingNull" json:"exact_matching_null,omitempty"`
	OnlineStatusInsteadOfIsOnline     *dstore_values.BooleanValue `protobuf:"bytes,14,opt,name=online_status_instead_of_is_online,json=onlineStatusInsteadOfIsOnline" json:"online_status_instead_of_is_online,omitempty"`
	OnlineStatusInsteadOfIsOnlineNull bool                        `protobuf:"varint,1014,opt,name=online_status_instead_of_is_online_null,json=onlineStatusInsteadOfIsOnlineNull" json:"online_status_instead_of_is_online_null,omitempty"`
	CharacteristicIdList              *dstore_values.StringValue  `protobuf:"bytes,15,opt,name=characteristic_id_list,json=characteristicIdList" json:"characteristic_id_list,omitempty"`
	CharacteristicIdListNull          bool                        `protobuf:"varint,1015,opt,name=characteristic_id_list_null,json=characteristicIdListNull" json:"characteristic_id_list_null,omitempty"`
	ConditionList                     *dstore_values.StringValue  `protobuf:"bytes,16,opt,name=condition_list,json=conditionList" json:"condition_list,omitempty"`
	ConditionListNull                 bool                        `protobuf:"varint,1016,opt,name=condition_list_null,json=conditionListNull" json:"condition_list_null,omitempty"`
	Country                           *dstore_values.StringValue  `protobuf:"bytes,17,opt,name=country" json:"country,omitempty"`
	CountryNull                       bool                        `protobuf:"varint,1017,opt,name=country_null,json=countryNull" json:"country_null,omitempty"`
	OutputIntoOneId                   *dstore_values.IntegerValue `protobuf:"bytes,18,opt,name=output_into_one_id,json=outputIntoOneId" json:"output_into_one_id,omitempty"`
	OutputIntoOneIdNull               bool                        `protobuf:"varint,1018,opt,name=output_into_one_id_null,json=outputIntoOneIdNull" json:"output_into_one_id_null,omitempty"`
	SearchOnlyMembersInOneId          *dstore_values.IntegerValue `protobuf:"bytes,19,opt,name=search_only_members_in_one_id,json=searchOnlyMembersInOneId" json:"search_only_members_in_one_id,omitempty"`
	SearchOnlyMembersInOneIdNull      bool                        `protobuf:"varint,1019,opt,name=search_only_members_in_one_id_null,json=searchOnlyMembersInOneIdNull" json:"search_only_members_in_one_id_null,omitempty"`
	FilterByBinaryCategoryId          *dstore_values.IntegerValue `protobuf:"bytes,20,opt,name=filter_by_binary_category_id,json=filterByBinaryCategoryId" json:"filter_by_binary_category_id,omitempty"`
	FilterByBinaryCategoryIdNull      bool                        `protobuf:"varint,1020,opt,name=filter_by_binary_category_id_null,json=filterByBinaryCategoryIdNull" json:"filter_by_binary_category_id_null,omitempty"`
	LastNickname                      *dstore_values.StringValue  `protobuf:"bytes,21,opt,name=last_nickname,json=lastNickname" json:"last_nickname,omitempty"`
	LastNicknameNull                  bool                        `protobuf:"varint,1021,opt,name=last_nickname_null,json=lastNicknameNull" json:"last_nickname_null,omitempty"`
	Next                              *dstore_values.BooleanValue `protobuf:"bytes,22,opt,name=next" json:"next,omitempty"`
	NextNull                          bool                        `protobuf:"varint,1022,opt,name=next_null,json=nextNull" json:"next_null,omitempty"`
	SeparatorInIdentVals              *dstore_values.StringValue  `protobuf:"bytes,23,opt,name=separator_in_ident_vals,json=separatorInIdentVals" json:"separator_in_ident_vals,omitempty"`
	SeparatorInIdentValsNull          bool                        `protobuf:"varint,1023,opt,name=separator_in_ident_vals_null,json=separatorInIdentValsNull" json:"separator_in_ident_vals_null,omitempty"`
	SeparatorInConditionList          *dstore_values.StringValue  `protobuf:"bytes,24,opt,name=separator_in_condition_list,json=separatorInConditionList" json:"separator_in_condition_list,omitempty"`
	SeparatorInConditionListNull      bool                        `protobuf:"varint,1024,opt,name=separator_in_condition_list_null,json=separatorInConditionListNull" json:"separator_in_condition_list_null,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parameters) GetRowcount() *dstore_values.IntegerValue {
	if m != nil {
		return m.Rowcount
	}
	return nil
}

func (m *Parameters) GetResult() *dstore_values.IntegerValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Parameters) GetUniqueId() *dstore_values.StringValue {
	if m != nil {
		return m.UniqueId
	}
	return nil
}

func (m *Parameters) GetPersonIdentificationValues() *dstore_values.StringValue {
	if m != nil {
		return m.PersonIdentificationValues
	}
	return nil
}

func (m *Parameters) GetCommunityId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *Parameters) GetSearchString() *dstore_values.StringValue {
	if m != nil {
		return m.SearchString
	}
	return nil
}

func (m *Parameters) GetMaxNumberOfRows() *dstore_values.IntegerValue {
	if m != nil {
		return m.MaxNumberOfRows
	}
	return nil
}

func (m *Parameters) GetCaseSensitive() *dstore_values.BooleanValue {
	if m != nil {
		return m.CaseSensitive
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId1() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId1
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId2() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId2
	}
	return nil
}

func (m *Parameters) GetOutputCharacteristicId3() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputCharacteristicId3
	}
	return nil
}

func (m *Parameters) GetCommunityBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityBinaryCategoryId
	}
	return nil
}

func (m *Parameters) GetExactMatching() *dstore_values.BooleanValue {
	if m != nil {
		return m.ExactMatching
	}
	return nil
}

func (m *Parameters) GetOnlineStatusInsteadOfIsOnline() *dstore_values.BooleanValue {
	if m != nil {
		return m.OnlineStatusInsteadOfIsOnline
	}
	return nil
}

func (m *Parameters) GetCharacteristicIdList() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicIdList
	}
	return nil
}

func (m *Parameters) GetConditionList() *dstore_values.StringValue {
	if m != nil {
		return m.ConditionList
	}
	return nil
}

func (m *Parameters) GetCountry() *dstore_values.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Parameters) GetOutputIntoOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.OutputIntoOneId
	}
	return nil
}

func (m *Parameters) GetSearchOnlyMembersInOneId() *dstore_values.IntegerValue {
	if m != nil {
		return m.SearchOnlyMembersInOneId
	}
	return nil
}

func (m *Parameters) GetFilterByBinaryCategoryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.FilterByBinaryCategoryId
	}
	return nil
}

func (m *Parameters) GetLastNickname() *dstore_values.StringValue {
	if m != nil {
		return m.LastNickname
	}
	return nil
}

func (m *Parameters) GetNext() *dstore_values.BooleanValue {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *Parameters) GetSeparatorInIdentVals() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInIdentVals
	}
	return nil
}

func (m *Parameters) GetSeparatorInConditionList() *dstore_values.StringValue {
	if m != nil {
		return m.SeparatorInConditionList
	}
	return nil
}

type Response struct {
	Error           *dstore_engine_error.Error                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	MetaInformation []*dstore_engine_metainformation.MetaInformation `protobuf:"bytes,2,rep,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Message         []*dstore_engine_message.Message                 `protobuf:"bytes,3,rep,name=message" json:"message,omitempty"`
	Row             []*Response_Row                                  `protobuf:"bytes,4,rep,name=row" json:"row,omitempty"`
	Count           *dstore_values.IntegerValue                      `protobuf:"bytes,101,opt,name=count" json:"count,omitempty"`
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

func (m *Response) GetCount() *dstore_values.IntegerValue {
	if m != nil {
		return m.Count
	}
	return nil
}

type Response_Row struct {
	RowId                     int32                       `protobuf:"varint,10000,opt,name=row_id,json=rowId" json:"row_id,omitempty"`
	Value1RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10001,opt,name=value1_restricted_by_pattern,json=value1RestrictedByPattern" json:"value1_restricted_by_pattern,omitempty"`
	BinaryId                  *dstore_values.IntegerValue `protobuf:"bytes,10002,opt,name=binary_id,json=binaryId" json:"binary_id,omitempty"`
	Value2RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10003,opt,name=value2_restricted_by_pattern,json=value2RestrictedByPattern" json:"value2_restricted_by_pattern,omitempty"`
	CommunityMemberId         *dstore_values.IntegerValue `protobuf:"bytes,10004,opt,name=community_member_id,json=communityMemberId" json:"community_member_id,omitempty"`
	Value3                    *dstore_values.StringValue  `protobuf:"bytes,10005,opt,name=value3" json:"value3,omitempty"`
	OnlineStatus              *dstore_values.IntegerValue `protobuf:"bytes,10006,opt,name=online_status,json=onlineStatus" json:"online_status,omitempty"`
	Value1                    *dstore_values.StringValue  `protobuf:"bytes,10007,opt,name=value1" json:"value1,omitempty"`
	Value2                    *dstore_values.StringValue  `protobuf:"bytes,10008,opt,name=value2" json:"value2,omitempty"`
	Value3RestrictedByPattern *dstore_values.StringValue  `protobuf:"bytes,10009,opt,name=value3_restricted_by_pattern,json=value3RestrictedByPattern" json:"value3_restricted_by_pattern,omitempty"`
	Nickname                  *dstore_values.StringValue  `protobuf:"bytes,10010,opt,name=nickname" json:"nickname,omitempty"`
	IsOnline                  *dstore_values.BooleanValue `protobuf:"bytes,10011,opt,name=is_online,json=isOnline" json:"is_online,omitempty"`
	Operator1                 *dstore_values.StringValue  `protobuf:"bytes,20001,opt,name=operator1" json:"operator1,omitempty"`
	Condition1                *dstore_values.StringValue  `protobuf:"bytes,20002,opt,name=condition1" json:"condition1,omitempty"`
	Condition2                *dstore_values.StringValue  `protobuf:"bytes,20003,opt,name=condition2" json:"condition2,omitempty"`
	Operator2                 *dstore_values.StringValue  `protobuf:"bytes,20004,opt,name=operator2" json:"operator2,omitempty"`
	CharacteristicDescription *dstore_values.StringValue  `protobuf:"bytes,20005,opt,name=characteristic_description,json=characteristicDescription" json:"characteristic_description,omitempty"`
	EstimatedRowsAffected     *dstore_values.IntegerValue `protobuf:"bytes,20006,opt,name=estimated_rows_affected,json=estimatedRowsAffected" json:"estimated_rows_affected,omitempty"`
	CharacteristicId          *dstore_values.IntegerValue `protobuf:"bytes,20007,opt,name=characteristic_id,json=characteristicId" json:"characteristic_id,omitempty"`
}

func (m *Response_Row) Reset()                    { *m = Response_Row{} }
func (m *Response_Row) String() string            { return proto.CompactTextString(m) }
func (*Response_Row) ProtoMessage()               {}
func (*Response_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Row) GetValue1RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value1RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetBinaryId() *dstore_values.IntegerValue {
	if m != nil {
		return m.BinaryId
	}
	return nil
}

func (m *Response_Row) GetValue2RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value2RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetCommunityMemberId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CommunityMemberId
	}
	return nil
}

func (m *Response_Row) GetValue3() *dstore_values.StringValue {
	if m != nil {
		return m.Value3
	}
	return nil
}

func (m *Response_Row) GetOnlineStatus() *dstore_values.IntegerValue {
	if m != nil {
		return m.OnlineStatus
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

func (m *Response_Row) GetValue3RestrictedByPattern() *dstore_values.StringValue {
	if m != nil {
		return m.Value3RestrictedByPattern
	}
	return nil
}

func (m *Response_Row) GetNickname() *dstore_values.StringValue {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *Response_Row) GetIsOnline() *dstore_values.BooleanValue {
	if m != nil {
		return m.IsOnline
	}
	return nil
}

func (m *Response_Row) GetOperator1() *dstore_values.StringValue {
	if m != nil {
		return m.Operator1
	}
	return nil
}

func (m *Response_Row) GetCondition1() *dstore_values.StringValue {
	if m != nil {
		return m.Condition1
	}
	return nil
}

func (m *Response_Row) GetCondition2() *dstore_values.StringValue {
	if m != nil {
		return m.Condition2
	}
	return nil
}

func (m *Response_Row) GetOperator2() *dstore_values.StringValue {
	if m != nil {
		return m.Operator2
	}
	return nil
}

func (m *Response_Row) GetCharacteristicDescription() *dstore_values.StringValue {
	if m != nil {
		return m.CharacteristicDescription
	}
	return nil
}

func (m *Response_Row) GetEstimatedRowsAffected() *dstore_values.IntegerValue {
	if m != nil {
		return m.EstimatedRowsAffected
	}
	return nil
}

func (m *Response_Row) GetCharacteristicId() *dstore_values.IntegerValue {
	if m != nil {
		return m.CharacteristicId
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "dstore.engine.co_SearchMembers_Pu.Parameters")
	proto.RegisterType((*Response)(nil), "dstore.engine.co_SearchMembers_Pu.Response")
	proto.RegisterType((*Response_Row)(nil), "dstore.engine.co_SearchMembers_Pu.Response.Row")
}

func init() { proto.RegisterFile("dstore/engine/procedures/co_SearchMembers_Pu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x58, 0xd9, 0x6f, 0xdb, 0xc6,
	0x13, 0x46, 0x7e, 0x8e, 0x63, 0x79, 0x6c, 0x27, 0x36, 0x9d, 0xc4, 0xeb, 0x23, 0x41, 0xec, 0x5f,
	0x8b, 0x16, 0x28, 0x2a, 0x57, 0x54, 0x83, 0xb4, 0x45, 0xd1, 0x22, 0x4e, 0x8a, 0x46, 0x41, 0x7c,
	0x54, 0xee, 0x81, 0x1e, 0x00, 0x41, 0x53, 0x2b, 0x87, 0x88, 0x44, 0xaa, 0x5c, 0x2a, 0x89, 0xdf,
	0xfa, 0xdc, 0xa7, 0xde, 0x67, 0x5e, 0x7a, 0xff, 0x43, 0xfd, 0x27, 0x7a, 0xdf, 0xf7, 0x91, 0xce,
	0xee, 0x2c, 0x29, 0x51, 0x96, 0xb4, 0x4c, 0x5e, 0x6c, 0x89, 0x3b, 0xdf, 0x37, 0x9f, 0x76, 0x67,
	0x66, 0x67, 0x08, 0x76, 0x4d, 0xc4, 0x61, 0xc4, 0x57, 0x79, 0xb0, 0xeb, 0x07, 0x7c, 0xb5, 0x15,
	0x85, 0x1e, 0xaf, 0xb5, 0x23, 0x2e, 0x56, 0xbd, 0xd0, 0xd9, 0xe6, 0x6e, 0xe4, 0x5d, 0x5e, 0xe7,
	0xcd, 0x1d, 0x1e, 0x09, 0x67, 0xab, 0x5d, 0xc4, 0xe5, 0x38, 0xb4, 0x96, 0x09, 0x53, 0x24, 0x4c,
	0xb1, 0x8f, 0xe1, 0xc2, 0xac, 0xa6, 0xbd, 0xea, 0x36, 0xda, 0x5c, 0x10, 0x6e, 0x61, 0x3e, 0xeb,
	0x8b, 0x47, 0x51, 0x18, 0xe9, 0xa5, 0xc5, 0xec, 0x52, 0x93, 0x0b, 0xe1, 0xee, 0x72, 0xbd, 0xf8,
	0xff, 0xde, 0xc5, 0xd8, 0xf5, 0x83, 0x7a, 0x18, 0x35, 0xdd, 0xd8, 0x0f, 0x03, 0x32, 0x5a, 0xb9,
	0xc9, 0x00, 0xb6, 0xdc, 0xc8, 0xc5, 0x55, 0xd4, 0x60, 0x9d, 0x81, 0x42, 0x14, 0x5e, 0xf3, 0xc2,
	0x76, 0x10, 0xb3, 0x03, 0xa7, 0x0e, 0xdc, 0x3d, 0x61, 0x2f, 0x16, 0xb5, 0x6c, 0xad, 0xc9, 0x0f,
	0x62, 0xbe, 0xcb, 0xa3, 0xa7, 0xe5, 0xb7, 0x6a, 0x6a, 0x6c, 0xdd, 0x01, 0x53, 0xc9, 0x67, 0x27,
	0x68, 0x37, 0x1a, 0xec, 0x8b, 0x31, 0x84, 0x17, 0xaa, 0x93, 0xc9, 0xd3, 0x0d, 0x7c, 0x68, 0x95,
	0xe1, 0x10, 0xee, 0x51, 0xbb, 0x11, 0xb3, 0xff, 0x99, 0xc9, 0xb5, 0xa9, 0x75, 0x0a, 0x26, 0xe8,
	0x13, 0x11, 0x7f, 0x49, 0xc4, 0x40, 0xcf, 0x14, 0xed, 0x19, 0x18, 0x6f, 0x07, 0xfe, 0x8b, 0x6d,
	0xee, 0xf8, 0x35, 0x36, 0xa2, 0x98, 0x17, 0x7a, 0x98, 0x45, 0x1c, 0xf9, 0xc1, 0xae, 0x56, 0x4d,
	0xc6, 0x95, 0x9a, 0x75, 0x27, 0x1c, 0x4e, 0x81, 0xc4, 0xfe, 0x95, 0x96, 0x9d, 0x98, 0x28, 0xfe,
	0x17, 0x60, 0xa9, 0x85, 0xbb, 0x13, 0x06, 0x68, 0xc6, 0x83, 0xd8, 0xaf, 0xfb, 0x9e, 0xda, 0x43,
	0x87, 0xc8, 0xd9, 0x41, 0xa3, 0xcb, 0x05, 0xc2, 0x57, 0x32, 0x70, 0xb5, 0x24, 0xac, 0x8b, 0xb0,
	0x3c, 0x8c, 0x9d, 0x74, 0x7d, 0x4d, 0xba, 0x4e, 0x0e, 0xe6, 0x51, 0x4a, 0x1f, 0x81, 0x49, 0x2f,
	0x6c, 0x36, 0x51, 0x7d, 0xbc, 0x27, 0x37, 0x63, 0xd4, 0xbc, 0xcd, 0x13, 0x29, 0x00, 0x37, 0xe4,
	0x1e, 0x98, 0xe9, 0xc6, 0x93, 0xef, 0x6f, 0xc8, 0xf7, 0x91, 0x2e, 0x43, 0xe5, 0xec, 0x51, 0x98,
	0x12, 0x2a, 0x82, 0x1d, 0xfa, 0xa9, 0xec, 0x90, 0x71, 0x1f, 0x26, 0x09, 0xb0, 0xad, 0x1e, 0x59,
	0xf7, 0x82, 0x95, 0x21, 0x20, 0x77, 0xdf, 0x92, 0xbb, 0xe9, 0x6e, 0x53, 0xe5, 0xef, 0x02, 0x58,
	0x4d, 0xf7, 0x3a, 0x1a, 0xc9, 0x7c, 0x71, 0xc2, 0xba, 0x83, 0xb1, 0x25, 0xd8, 0x98, 0xf9, 0x27,
	0x1e, 0x41, 0xd8, 0x86, 0x42, 0x6d, 0xd6, 0xab, 0x88, 0xb1, 0x4e, 0xc3, 0xdc, 0x7e, 0x26, 0xf2,
	0xfe, 0x1d, 0x79, 0x9f, 0xed, 0x81, 0x28, 0x01, 0x6b, 0x70, 0xd8, 0x73, 0x05, 0x77, 0x04, 0x0f,
	0x84, 0x1f, 0xfb, 0x57, 0x39, 0x2b, 0xf4, 0x75, 0xbe, 0x13, 0x86, 0x0d, 0xee, 0xd2, 0xb9, 0x54,
	0xa7, 0x24, 0x64, 0x3b, 0x41, 0x58, 0xab, 0x30, 0x9b, 0xe5, 0x20, 0xb7, 0xdf, 0x93, 0xdb, 0x99,
	0x8c, 0xb1, 0x72, 0xfa, 0x0c, 0xcc, 0x87, 0xed, 0xb8, 0xd5, 0x8e, 0x1d, 0xef, 0x32, 0x26, 0xaa,
	0x87, 0x79, 0xea, 0x8b, 0xd8, 0xf7, 0xf0, 0x78, 0x4a, 0x6c, 0xdc, 0xfc, 0xe3, 0xe7, 0x08, 0x7d,
	0x2e, 0x03, 0xae, 0xd4, 0x4a, 0xd6, 0x79, 0x38, 0x39, 0x90, 0x98, 0x44, 0xfd, 0x40, 0xa2, 0x16,
	0x07, 0x30, 0x98, 0xe4, 0xd9, 0x0c, 0x6e, 0x5b, 0x9e, 0x3d, 0x4c, 0x9e, 0x4d, 0xf2, 0x7e, 0x1c,
	0x2a, 0xcf, 0x36, 0xc9, 0x2b, 0xb3, 0x89, 0xdb, 0x96, 0x57, 0x1e, 0x26, 0xaf, 0x4c, 0xf2, 0x7e,
	0x1a, 0x2a, 0xaf, 0x9c, 0x54, 0x96, 0x4e, 0xbe, 0xed, 0xf8, 0x81, 0x1b, 0xed, 0x39, 0x98, 0xd4,
	0x7c, 0x37, 0x8c, 0x54, 0xfe, 0x4e, 0x9a, 0x15, 0xce, 0xa7, 0x04, 0x6b, 0x0a, 0x7f, 0x4e, 0xc3,
	0x31, 0x9b, 0x2b, 0xb0, 0x3c, 0x8c, 0x9d, 0x64, 0xfe, 0x4c, 0x32, 0x4f, 0x0c, 0xa4, 0x49, 0x42,
	0x9f, 0x5f, 0x47, 0xfd, 0x0e, 0xde, 0x1e, 0xde, 0x65, 0x99, 0xec, 0x53, 0x39, 0x42, 0x5f, 0x41,
	0xd6, 0x35, 0x42, 0x86, 0x7e, 0x96, 0x83, 0x04, 0xfc, 0xa2, 0x43, 0x3f, 0x63, 0xac, 0x9c, 0xd6,
	0x61, 0x25, 0x0c, 0x1a, 0x78, 0x79, 0x61, 0x7d, 0x70, 0xe3, 0xb6, 0x70, 0xfc, 0x40, 0xc4, 0xdc,
	0xad, 0xc9, 0x8c, 0xf5, 0x85, 0x43, 0x6b, 0xec, 0xb0, 0x59, 0xc8, 0x09, 0x32, 0xdd, 0x56, 0x2c,
	0x15, 0x22, 0xd9, 0xac, 0x57, 0xc4, 0xa6, 0x7a, 0x6c, 0x6d, 0xc3, 0x5d, 0x66, 0x3f, 0x24, 0xf6,
	0x57, 0x12, 0xbb, 0x3c, 0x94, 0x50, 0x89, 0xdf, 0x82, 0xe3, 0xfb, 0x22, 0xc3, 0x69, 0xe0, 0x27,
	0x76, 0xc4, 0x58, 0x26, 0x8f, 0x7a, 0x3d, 0xd1, 0x72, 0x09, 0xff, 0x63, 0x71, 0x5f, 0xec, 0xcf,
	0x48, 0xd2, 0x7e, 0x23, 0x69, 0xac, 0x1f, 0x56, 0x29, 0x3a, 0x8b, 0xe5, 0x2b, 0x0c, 0x6a, 0xbe,
	0xba, 0x5c, 0x94, 0x92, 0x69, 0xa3, 0x92, 0xa9, 0x14, 0xa1, 0x24, 0xc8, 0xea, 0x95, 0xa1, 0x20,
	0xd7, 0xbf, 0x27, 0xd5, 0xab, 0xdb, 0x58, 0xf9, 0xbc, 0x1f, 0xc6, 0xd4, 0xf5, 0x1f, 0xed, 0xb1,
	0x19, 0xa3, 0xb3, 0xc4, 0xd4, 0x5a, 0x91, 0xd7, 0x98, 0xfa, 0x48, 0xfc, 0x7f, 0x10, 0xff, 0x84,
	0x7e, 0x98, 0xdc, 0x06, 0x3a, 0x01, 0x31, 0x1d, 0x42, 0x3c, 0x24, 0x75, 0xfb, 0x5b, 0x39, 0x6e,
	0x03, 0x82, 0x55, 0x10, 0xb5, 0x19, 0xc8, 0x2e, 0x00, 0x6f, 0x83, 0xfd, 0x4c, 0xe4, 0xf8, 0x4f,
	0x7d, 0x1b, 0xf4, 0x40, 0x74, 0xee, 0x9e, 0xd0, 0xb7, 0x17, 0x06, 0xc3, 0x9e, 0xd3, 0xd4, 0x6d,
	0x9c, 0x1f, 0x24, 0x5a, 0x66, 0xcd, 0x5a, 0x18, 0x31, 0x60, 0xd8, 0xec, 0xe9, 0x36, 0xb0, 0x12,
	0x90, 0xa8, 0x0a, 0xac, 0x0c, 0x65, 0x27, 0x7d, 0x7f, 0x91, 0xbe, 0xa5, 0x41, 0x34, 0x4a, 0xe8,
	0xf3, 0xb0, 0x54, 0xf7, 0x1b, 0x18, 0x10, 0xce, 0x4e, 0xdf, 0x22, 0x73, 0x34, 0x87, 0x4e, 0x22,
	0x58, 0xdb, 0x5f, 0x63, 0x2e, 0xc0, 0xf2, 0x30, 0x72, 0x92, 0xf9, 0xb7, 0x96, 0x39, 0x88, 0x25,
	0x69, 0x27, 0x1a, 0xae, 0x8c, 0x28, 0xdf, 0xbb, 0x12, 0x60, 0x43, 0xca, 0x8e, 0x99, 0xdb, 0x09,
	0x09, 0xd8, 0xd0, 0xf6, 0xb2, 0x9d, 0xc8, 0x10, 0x90, 0xef, 0x7f, 0x74, 0x3b, 0xd1, 0x6d, 0xaa,
	0xfc, 0xad, 0xc2, 0xc1, 0x80, 0x5f, 0x8f, 0xd9, 0x71, 0x73, 0xfd, 0x50, 0x86, 0xd6, 0x12, 0x8c,
	0xcb, 0xff, 0x44, 0xfb, 0x2f, 0xd1, 0x16, 0xe4, 0x13, 0x45, 0xf7, 0x04, 0xcc, 0x09, 0xde, 0xc2,
	0xd4, 0x43, 0x16, 0x79, 0x4c, 0xaa, 0x99, 0x93, 0x3d, 0x9c, 0x60, 0x73, 0xe6, 0x84, 0x4f, 0xa1,
	0x15, 0x6a, 0xee, 0xf0, 0xb1, 0xc0, 0x1d, 0x59, 0x1a, 0x40, 0x49, 0x1a, 0x6e, 0xea, 0x8c, 0xef,
	0x07, 0x56, 0x9a, 0x9e, 0x85, 0xc5, 0x0c, 0x41, 0x4f, 0xfa, 0x33, 0xa3, 0xae, 0x6e, 0xea, 0x73,
	0x99, 0x4a, 0xf0, 0x38, 0x9c, 0x1a, 0x42, 0x4d, 0xfa, 0x5e, 0x2a, 0x24, 0xd1, 0xd9, 0x9f, 0x44,
	0x6a, 0x5c, 0xf9, 0x7c, 0x12, 0x0a, 0x55, 0x2e, 0x5a, 0x61, 0x20, 0xb8, 0x75, 0x1f, 0x8c, 0xaa,
	0xf9, 0x46, 0x0f, 0x1f, 0xa9, 0x34, 0x3d, 0x33, 0xd1, 0xec, 0xf3, 0x98, 0xfc, 0x5b, 0x25, 0x43,
	0xfc, 0x89, 0xd3, 0x72, 0xb2, 0x71, 0xba, 0x46, 0x1b, 0x1c, 0x2e, 0x46, 0x10, 0x5c, 0xec, 0x01,
	0xf7, 0x0e, 0x40, 0xeb, 0xf8, 0xbd, 0xd2, 0xf9, 0x8e, 0x5d, 0x62, 0xf6, 0x81, 0xf5, 0x00, 0x8c,
	0xe9, 0x89, 0x0a, 0x87, 0x0a, 0xc9, 0x78, 0x72, 0x1f, 0x23, 0xcd, 0x5b, 0xeb, 0xf4, 0xbf, 0x9a,
	0x98, 0x63, 0xa5, 0x1d, 0xc1, 0x8e, 0x12, 0xe7, 0x02, 0x89, 0x5a, 0x2d, 0x1a, 0x07, 0xbf, 0x62,
	0xb2, 0x01, 0x45, 0xec, 0x35, 0xab, 0x12, 0x6b, 0x95, 0x60, 0x94, 0xc6, 0x30, 0x6e, 0xce, 0x4e,
	0xb2, 0x5c, 0x78, 0x19, 0x60, 0x04, 0xf1, 0xd6, 0x71, 0x9c, 0xb2, 0xc2, 0x6b, 0x32, 0xb3, 0x5f,
	0xd9, 0x40, 0xf0, 0x68, 0x75, 0x14, 0xbf, 0x62, 0xaa, 0x62, 0xb3, 0xa1, 0xd0, 0x25, 0x07, 0x67,
	0x27, 0x3c, 0x64, 0xbc, 0x23, 0x6a, 0x32, 0x6b, 0x5b, 0x6e, 0x8c, 0x59, 0x19, 0xb0, 0x57, 0x37,
	0x8c, 0xf1, 0x30, 0x4f, 0x04, 0xd5, 0x14, 0xbf, 0xb6, 0xb7, 0x45, 0x68, 0xeb, 0x41, 0x18, 0xd7,
	0xe9, 0x8f, 0x8e, 0x5f, 0xdb, 0xc8, 0x31, 0x3c, 0x92, 0x79, 0x97, 0x30, 0x7b, 0x80, 0xb0, 0xd7,
	0xf3, 0x0a, 0xb3, 0xfb, 0x09, 0xbb, 0x24, 0xef, 0xac, 0xa4, 0x0b, 0xa2, 0x3a, 0x2a, 0x25, 0xbe,
	0x91, 0x43, 0x62, 0x67, 0x18, 0xa2, 0xb3, 0x42, 0xad, 0x38, 0xc2, 0x2a, 0xcb, 0x32, 0x7b, 0xd3,
	0xac, 0x4a, 0x9b, 0x62, 0x3c, 0x4c, 0x65, 0x1a, 0x0c, 0xf6, 0x56, 0x0e, 0xe7, 0x93, 0xdd, 0x3d,
	0x46, 0xea, 0xb7, 0xc4, 0xde, 0xce, 0xeb, 0xb7, 0x94, 0x82, 0x6c, 0xf6, 0x4e, 0x5e, 0x90, 0x9d,
	0x9e, 0x46, 0x79, 0xc0, 0x69, 0xbc, 0x9b, 0xf7, 0x34, 0xca, 0xfd, 0x4e, 0xe3, 0x0c, 0x14, 0xd2,
	0x02, 0xff, 0x9e, 0x99, 0x29, 0x35, 0x96, 0xf1, 0xd5, 0xe9, 0xf9, 0xde, 0xdf, 0x30, 0x17, 0xed,
	0x82, 0x9f, 0xf4, 0x77, 0x0f, 0xc1, 0x78, 0x88, 0x83, 0xb3, 0xac, 0x40, 0x25, 0xf6, 0xc1, 0x8d,
	0x03, 0x46, 0xaf, 0x1d, 0x73, 0xeb, 0x61, 0x80, 0xb4, 0xb4, 0x95, 0xd8, 0x87, 0x39, 0xc0, 0x5d,
	0xf6, 0x19, 0xb4, 0xcd, 0x3e, 0xba, 0x25, 0xb4, 0xdd, 0xad, 0xdb, 0x66, 0x1f, 0xdf, 0x8a, 0x6e,
	0x1b, 0x2f, 0xfd, 0x85, 0x9e, 0x66, 0xb1, 0xc6, 0x85, 0x17, 0xf9, 0x2d, 0x55, 0x21, 0x3f, 0xc9,
	0x41, 0x36, 0x9f, 0xc5, 0x9f, 0xef, 0xc0, 0xad, 0xa7, 0x60, 0x0e, 0x8f, 0xd6, 0xc7, 0x3a, 0x89,
	0xa1, 0xa1, 0x66, 0x67, 0xb7, 0x5e, 0xe7, 0xf2, 0xa0, 0xd9, 0xa7, 0x37, 0x72, 0xbc, 0x37, 0x3a,
	0x96, 0xa2, 0xe5, 0x6c, 0x7d, 0x56, 0x63, 0xad, 0x8b, 0x30, 0xb3, 0xaf, 0xc1, 0x65, 0x9f, 0xe5,
	0x21, 0x9c, 0xee, 0x6d, 0x7a, 0xd7, 0x9e, 0x84, 0x45, 0x3f, 0xec, 0xa9, 0xbc, 0x9d, 0xd7, 0x74,
	0xcf, 0x9d, 0xde, 0x0d, 0x45, 0xed, 0x4a, 0xb2, 0x5e, 0xcb, 0xf9, 0x26, 0x6f, 0xe7, 0x90, 0x7a,
	0x6b, 0x56, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x67, 0xe8, 0xf5, 0x5b, 0x00, 0x14, 0x00, 0x00,
}