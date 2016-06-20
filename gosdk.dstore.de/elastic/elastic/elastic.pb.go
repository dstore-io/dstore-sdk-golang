// Code generated by protoc-gen-go.
// source: dstore/elastic/elastic.proto
// DO NOT EDIT!

/*
Package elastic is a generated protocol buffer package.

It is generated from these files:
	dstore/elastic/elastic.proto

It has these top-level messages:
	Field
	Range
	BoolQuery
	Query
	Sort
*/
package elastic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dstore_values "gosdk.dstore.de/values"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Sort_Order int32

const (
	Sort_ASCENDING  Sort_Order = 0
	Sort_DESCENDING Sort_Order = 1
)

var Sort_Order_name = map[int32]string{
	0: "ASCENDING",
	1: "DESCENDING",
}
var Sort_Order_value = map[string]int32{
	"ASCENDING":  0,
	"DESCENDING": 1,
}

func (x Sort_Order) String() string {
	return proto.EnumName(Sort_Order_name, int32(x))
}
func (Sort_Order) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Sort_Missing int32

const (
	Sort_LAST  Sort_Missing = 0
	Sort_FIRST Sort_Missing = 1
)

var Sort_Missing_name = map[int32]string{
	0: "LAST",
	1: "FIRST",
}
var Sort_Missing_value = map[string]int32{
	"LAST":  0,
	"FIRST": 1,
}

func (x Sort_Missing) String() string {
	return proto.EnumName(Sort_Missing_name, int32(x))
}
func (Sort_Missing) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

type Sort_Mode int32

const (
	Sort_UNDEFINED Sort_Mode = 0
	Sort_MIN       Sort_Mode = 1
	Sort_MAX       Sort_Mode = 2
	Sort_SUM       Sort_Mode = 3
	Sort_AVG       Sort_Mode = 4
	Sort_MEDIAN    Sort_Mode = 5
)

var Sort_Mode_name = map[int32]string{
	0: "UNDEFINED",
	1: "MIN",
	2: "MAX",
	3: "SUM",
	4: "AVG",
	5: "MEDIAN",
}
var Sort_Mode_value = map[string]int32{
	"UNDEFINED": 0,
	"MIN":       1,
	"MAX":       2,
	"SUM":       3,
	"AVG":       4,
	"MEDIAN":    5,
}

func (x Sort_Mode) String() string {
	return proto.EnumName(Sort_Mode_name, int32(x))
}
func (Sort_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 2} }

type Field struct {
	Name       string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MultiField bool                   `protobuf:"varint,2,opt,name=multi_field,json=multiField" json:"multi_field,omitempty"`
	Value      []*dstore_values.Value `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Field) GetValue() []*dstore_values.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Range struct {
	From string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BoolQuery struct {
	Filter             []*Query `protobuf:"bytes,1,rep,name=filter" json:"filter,omitempty"`
	Must               []*Query `protobuf:"bytes,2,rep,name=must" json:"must,omitempty"`
	Should             []*Query `protobuf:"bytes,3,rep,name=should" json:"should,omitempty"`
	MustNot            []*Query `protobuf:"bytes,4,rep,name=must_not,json=mustNot" json:"must_not,omitempty"`
	MinimumShouldMatch string   `protobuf:"bytes,5,opt,name=minimum_should_match,json=minimumShouldMatch" json:"minimum_should_match,omitempty"`
	Boost              float64  `protobuf:"fixed64,6,opt,name=boost" json:"boost,omitempty"`
}

func (m *BoolQuery) Reset()                    { *m = BoolQuery{} }
func (m *BoolQuery) String() string            { return proto.CompactTextString(m) }
func (*BoolQuery) ProtoMessage()               {}
func (*BoolQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BoolQuery) GetFilter() []*Query {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *BoolQuery) GetMust() []*Query {
	if m != nil {
		return m.Must
	}
	return nil
}

func (m *BoolQuery) GetShould() []*Query {
	if m != nil {
		return m.Should
	}
	return nil
}

func (m *BoolQuery) GetMustNot() []*Query {
	if m != nil {
		return m.MustNot
	}
	return nil
}

type Query struct {
	// Types that are valid to be assigned to Query:
	//	*Query_TermsQuery
	//	*Query_TypeQuery
	//	*Query_SimpleQueryStringQuery
	//	*Query_RangeQuery
	//	*Query_BoolQuery
	Query isQuery_Query `protobuf_oneof:"query"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isQuery_Query interface {
	isQuery_Query()
}

type Query_TermsQuery struct {
	TermsQuery *Query_Terms `protobuf:"bytes,1,opt,name=terms_query,json=termsQuery,oneof"`
}
type Query_TypeQuery struct {
	TypeQuery *Query_Type `protobuf:"bytes,2,opt,name=type_query,json=typeQuery,oneof"`
}
type Query_SimpleQueryStringQuery struct {
	SimpleQueryStringQuery *Query_SimpleQueryString `protobuf:"bytes,3,opt,name=simple_query_string_query,json=simpleQueryStringQuery,oneof"`
}
type Query_RangeQuery struct {
	RangeQuery *Query_Range `protobuf:"bytes,4,opt,name=range_query,json=rangeQuery,oneof"`
}
type Query_BoolQuery struct {
	BoolQuery *BoolQuery `protobuf:"bytes,20,opt,name=bool_query,json=boolQuery,oneof"`
}

func (*Query_TermsQuery) isQuery_Query()             {}
func (*Query_TypeQuery) isQuery_Query()              {}
func (*Query_SimpleQueryStringQuery) isQuery_Query() {}
func (*Query_RangeQuery) isQuery_Query()             {}
func (*Query_BoolQuery) isQuery_Query()              {}

func (m *Query) GetQuery() isQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Query) GetTermsQuery() *Query_Terms {
	if x, ok := m.GetQuery().(*Query_TermsQuery); ok {
		return x.TermsQuery
	}
	return nil
}

func (m *Query) GetTypeQuery() *Query_Type {
	if x, ok := m.GetQuery().(*Query_TypeQuery); ok {
		return x.TypeQuery
	}
	return nil
}

func (m *Query) GetSimpleQueryStringQuery() *Query_SimpleQueryString {
	if x, ok := m.GetQuery().(*Query_SimpleQueryStringQuery); ok {
		return x.SimpleQueryStringQuery
	}
	return nil
}

func (m *Query) GetRangeQuery() *Query_Range {
	if x, ok := m.GetQuery().(*Query_RangeQuery); ok {
		return x.RangeQuery
	}
	return nil
}

func (m *Query) GetBoolQuery() *BoolQuery {
	if x, ok := m.GetQuery().(*Query_BoolQuery); ok {
		return x.BoolQuery
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Query) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Query_OneofMarshaler, _Query_OneofUnmarshaler, _Query_OneofSizer, []interface{}{
		(*Query_TermsQuery)(nil),
		(*Query_TypeQuery)(nil),
		(*Query_SimpleQueryStringQuery)(nil),
		(*Query_RangeQuery)(nil),
		(*Query_BoolQuery)(nil),
	}
}

func _Query_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Query)
	// query
	switch x := m.Query.(type) {
	case *Query_TermsQuery:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TermsQuery); err != nil {
			return err
		}
	case *Query_TypeQuery:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypeQuery); err != nil {
			return err
		}
	case *Query_SimpleQueryStringQuery:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SimpleQueryStringQuery); err != nil {
			return err
		}
	case *Query_RangeQuery:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RangeQuery); err != nil {
			return err
		}
	case *Query_BoolQuery:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BoolQuery); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Query.Query has unexpected type %T", x)
	}
	return nil
}

func _Query_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Query)
	switch tag {
	case 1: // query.terms_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Query_Terms)
		err := b.DecodeMessage(msg)
		m.Query = &Query_TermsQuery{msg}
		return true, err
	case 2: // query.type_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Query_Type)
		err := b.DecodeMessage(msg)
		m.Query = &Query_TypeQuery{msg}
		return true, err
	case 3: // query.simple_query_string_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Query_SimpleQueryString)
		err := b.DecodeMessage(msg)
		m.Query = &Query_SimpleQueryStringQuery{msg}
		return true, err
	case 4: // query.range_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Query_Range)
		err := b.DecodeMessage(msg)
		m.Query = &Query_RangeQuery{msg}
		return true, err
	case 20: // query.bool_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BoolQuery)
		err := b.DecodeMessage(msg)
		m.Query = &Query_BoolQuery{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Query_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Query)
	// query
	switch x := m.Query.(type) {
	case *Query_TermsQuery:
		s := proto.Size(x.TermsQuery)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_TypeQuery:
		s := proto.Size(x.TypeQuery)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_SimpleQueryStringQuery:
		s := proto.Size(x.SimpleQueryStringQuery)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_RangeQuery:
		s := proto.Size(x.RangeQuery)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_BoolQuery:
		s := proto.Size(x.BoolQuery)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Query_Terms struct {
	FieldName           string   `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	Value               []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
	AllowPartialMatches bool     `protobuf:"varint,3,opt,name=allow_partial_matches,json=allowPartialMatches" json:"allow_partial_matches,omitempty"`
}

func (m *Query_Terms) Reset()                    { *m = Query_Terms{} }
func (m *Query_Terms) String() string            { return proto.CompactTextString(m) }
func (*Query_Terms) ProtoMessage()               {}
func (*Query_Terms) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Query_Type struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Query_Type) Reset()                    { *m = Query_Type{} }
func (m *Query_Type) String() string            { return proto.CompactTextString(m) }
func (*Query_Type) ProtoMessage()               {}
func (*Query_Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

type Query_Range struct {
	FieldName         string `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	GreaterThanOrEqal string `protobuf:"bytes,2,opt,name=greater_than_or_eqal,json=greaterThanOrEqal" json:"greater_than_or_eqal,omitempty"`
	GreaterThan       string `protobuf:"bytes,3,opt,name=greater_than,json=greaterThan" json:"greater_than,omitempty"`
	LessThanOrEqual   string `protobuf:"bytes,4,opt,name=less_than_or_equal,json=lessThanOrEqual" json:"less_than_or_equal,omitempty"`
	LessThan          string `protobuf:"bytes,5,opt,name=less_than,json=lessThan" json:"less_than,omitempty"`
}

func (m *Query_Range) Reset()                    { *m = Query_Range{} }
func (m *Query_Range) String() string            { return proto.CompactTextString(m) }
func (*Query_Range) ProtoMessage()               {}
func (*Query_Range) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 2} }

type Query_SimpleQueryString struct {
	Query                   string   `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Field                   []string `protobuf:"bytes,2,rep,name=field" json:"field,omitempty"`
	UseAndAsDefaultOperator bool     `protobuf:"varint,3,opt,name=use_and_as_default_operator,json=useAndAsDefaultOperator" json:"use_and_as_default_operator,omitempty"`
	MinimumShouldMatch      string   `protobuf:"bytes,4,opt,name=minimum_should_match,json=minimumShouldMatch" json:"minimum_should_match,omitempty"`
}

func (m *Query_SimpleQueryString) Reset()                    { *m = Query_SimpleQueryString{} }
func (m *Query_SimpleQueryString) String() string            { return proto.CompactTextString(m) }
func (*Query_SimpleQueryString) ProtoMessage()               {}
func (*Query_SimpleQueryString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 3} }

// see: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-sort.html
type Sort struct {
	// Types that are valid to be assigned to SortBy:
	//	*Sort_FieldSort_
	//	*Sort_ScoreSort_
	SortBy isSort_SortBy `protobuf_oneof:"sort_by"`
}

func (m *Sort) Reset()                    { *m = Sort{} }
func (m *Sort) String() string            { return proto.CompactTextString(m) }
func (*Sort) ProtoMessage()               {}
func (*Sort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isSort_SortBy interface {
	isSort_SortBy()
}

type Sort_FieldSort_ struct {
	FieldSort *Sort_FieldSort `protobuf:"bytes,1,opt,name=field_sort,json=fieldSort,oneof"`
}
type Sort_ScoreSort_ struct {
	ScoreSort *Sort_ScoreSort `protobuf:"bytes,2,opt,name=score_sort,json=scoreSort,oneof"`
}

func (*Sort_FieldSort_) isSort_SortBy() {}
func (*Sort_ScoreSort_) isSort_SortBy() {}

func (m *Sort) GetSortBy() isSort_SortBy {
	if m != nil {
		return m.SortBy
	}
	return nil
}

func (m *Sort) GetFieldSort() *Sort_FieldSort {
	if x, ok := m.GetSortBy().(*Sort_FieldSort_); ok {
		return x.FieldSort
	}
	return nil
}

func (m *Sort) GetScoreSort() *Sort_ScoreSort {
	if x, ok := m.GetSortBy().(*Sort_ScoreSort_); ok {
		return x.ScoreSort
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Sort) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Sort_OneofMarshaler, _Sort_OneofUnmarshaler, _Sort_OneofSizer, []interface{}{
		(*Sort_FieldSort_)(nil),
		(*Sort_ScoreSort_)(nil),
	}
}

func _Sort_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Sort)
	// sort_by
	switch x := m.SortBy.(type) {
	case *Sort_FieldSort_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FieldSort); err != nil {
			return err
		}
	case *Sort_ScoreSort_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ScoreSort); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Sort.SortBy has unexpected type %T", x)
	}
	return nil
}

func _Sort_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Sort)
	switch tag {
	case 1: // sort_by.field_sort
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Sort_FieldSort)
		err := b.DecodeMessage(msg)
		m.SortBy = &Sort_FieldSort_{msg}
		return true, err
	case 2: // sort_by.score_sort
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Sort_ScoreSort)
		err := b.DecodeMessage(msg)
		m.SortBy = &Sort_ScoreSort_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Sort_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Sort)
	// sort_by
	switch x := m.SortBy.(type) {
	case *Sort_FieldSort_:
		s := proto.Size(x.FieldSort)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Sort_ScoreSort_:
		s := proto.Size(x.ScoreSort)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Sort_ScoreSort struct {
}

func (m *Sort_ScoreSort) Reset()                    { *m = Sort_ScoreSort{} }
func (m *Sort_ScoreSort) String() string            { return proto.CompactTextString(m) }
func (*Sort_ScoreSort) ProtoMessage()               {}
func (*Sort_ScoreSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Sort_FieldSort struct {
	FieldName        string       `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	SortOrder        Sort_Order   `protobuf:"varint,2,opt,name=sort_order,json=sortOrder,enum=dstore.elastic.Sort_Order" json:"sort_order,omitempty"`
	MissingTreatment Sort_Missing `protobuf:"varint,3,opt,name=missing_treatment,json=missingTreatment,enum=dstore.elastic.Sort_Missing" json:"missing_treatment,omitempty"`
	SortMode         Sort_Mode    `protobuf:"varint,4,opt,name=sort_mode,json=sortMode,enum=dstore.elastic.Sort_Mode" json:"sort_mode,omitempty"`
	NestedFilter     []*Query     `protobuf:"bytes,5,rep,name=nested_filter,json=nestedFilter" json:"nested_filter,omitempty"`
}

func (m *Sort_FieldSort) Reset()                    { *m = Sort_FieldSort{} }
func (m *Sort_FieldSort) String() string            { return proto.CompactTextString(m) }
func (*Sort_FieldSort) ProtoMessage()               {}
func (*Sort_FieldSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *Sort_FieldSort) GetNestedFilter() []*Query {
	if m != nil {
		return m.NestedFilter
	}
	return nil
}

func init() {
	proto.RegisterType((*Field)(nil), "dstore.elastic.Field")
	proto.RegisterType((*Range)(nil), "dstore.elastic.Range")
	proto.RegisterType((*BoolQuery)(nil), "dstore.elastic.BoolQuery")
	proto.RegisterType((*Query)(nil), "dstore.elastic.Query")
	proto.RegisterType((*Query_Terms)(nil), "dstore.elastic.Query.Terms")
	proto.RegisterType((*Query_Type)(nil), "dstore.elastic.Query.Type")
	proto.RegisterType((*Query_Range)(nil), "dstore.elastic.Query.Range")
	proto.RegisterType((*Query_SimpleQueryString)(nil), "dstore.elastic.Query.SimpleQueryString")
	proto.RegisterType((*Sort)(nil), "dstore.elastic.Sort")
	proto.RegisterType((*Sort_ScoreSort)(nil), "dstore.elastic.Sort.ScoreSort")
	proto.RegisterType((*Sort_FieldSort)(nil), "dstore.elastic.Sort.FieldSort")
	proto.RegisterEnum("dstore.elastic.Sort_Order", Sort_Order_name, Sort_Order_value)
	proto.RegisterEnum("dstore.elastic.Sort_Missing", Sort_Missing_name, Sort_Missing_value)
	proto.RegisterEnum("dstore.elastic.Sort_Mode", Sort_Mode_name, Sort_Mode_value)
}

func init() { proto.RegisterFile("dstore/elastic/elastic.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 934 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0x5d, 0x92, 0xdb, 0x44,
	0x10, 0x8e, 0x6d, 0x69, 0x6d, 0xb5, 0x13, 0xe3, 0x9d, 0x38, 0xe0, 0x78, 0x43, 0x12, 0xfc, 0x00,
	0x81, 0x14, 0x0e, 0xb5, 0x54, 0x51, 0xc5, 0x7f, 0xd9, 0xd8, 0x1b, 0x5c, 0x85, 0xbd, 0x20, 0x6f,
	0x52, 0x14, 0x2f, 0x2a, 0xed, 0x6a, 0xbc, 0xab, 0x42, 0xd2, 0x38, 0x1a, 0x19, 0x6a, 0x2f, 0xc1,
	0x29, 0x78, 0xe7, 0x12, 0x9c, 0x81, 0xa3, 0xf0, 0x4c, 0x77, 0xcf, 0xc8, 0x6c, 0xf6, 0xc7, 0x79,
	0x52, 0xcf, 0xf4, 0xf7, 0x7d, 0x3d, 0xdd, 0xea, 0x9e, 0x81, 0x07, 0x91, 0x2e, 0x54, 0x2e, 0x9f,
	0xc9, 0x24, 0xd4, 0x45, 0x7c, 0x52, 0x7e, 0x07, 0xab, 0x5c, 0x15, 0x4a, 0xb4, 0x8c, 0x77, 0x60,
	0x77, 0x7b, 0x77, 0x2d, 0xfa, 0xb7, 0x30, 0x59, 0x4b, 0x6d, 0x40, 0xfd, 0x33, 0x70, 0x0f, 0x62,
	0x99, 0x44, 0x42, 0x80, 0x93, 0x85, 0xa9, 0xec, 0x56, 0x1e, 0x57, 0x9e, 0x78, 0x3e, 0xdb, 0xe2,
	0x11, 0x34, 0xd3, 0x75, 0x52, 0xc4, 0xc1, 0x92, 0x20, 0xdd, 0x2a, 0xba, 0x1a, 0x3e, 0xf0, 0x96,
	0x21, 0x7d, 0x04, 0x2e, 0xab, 0x75, 0x6b, 0x8f, 0x6b, 0x4f, 0x9a, 0xfb, 0x9d, 0x81, 0x0d, 0x69,
	0x43, 0xbc, 0xa4, 0x8f, 0x6f, 0x20, 0xfd, 0xa7, 0xe0, 0xfa, 0x61, 0x76, 0x2a, 0x29, 0xd2, 0x32,
	0x57, 0x69, 0x19, 0x89, 0x6c, 0xd1, 0x82, 0x6a, 0xa1, 0x38, 0x80, 0xe7, 0xa3, 0xd5, 0xff, 0xa3,
	0x0a, 0xde, 0x48, 0xa9, 0xe4, 0xa7, 0xb5, 0xcc, 0xcf, 0xc5, 0xc7, 0xb0, 0xb3, 0x8c, 0x93, 0x42,
	0xe6, 0xc8, 0xa1, 0x38, 0xf7, 0x06, 0xaf, 0xa7, 0x36, 0x60, 0x98, 0x6f, 0x41, 0xe2, 0x43, 0x70,
	0xd2, 0xb5, 0x2e, 0x50, 0x6e, 0x0b, 0x98, 0x21, 0xa4, 0xac, 0xcf, 0xd4, 0x1a, 0x93, 0xab, 0x6d,
	0x55, 0x36, 0x20, 0xf1, 0x09, 0x34, 0x88, 0x16, 0x64, 0xaa, 0xe8, 0x3a, 0xdb, 0x08, 0x75, 0x82,
	0xcd, 0x55, 0x81, 0x8c, 0x4e, 0x1a, 0x67, 0x71, 0xba, 0x4e, 0x03, 0xa3, 0x11, 0xa4, 0x61, 0x71,
	0x72, 0xd6, 0x75, 0x39, 0x55, 0x61, 0x7d, 0x0b, 0x76, 0xcd, 0xc8, 0x23, 0x3a, 0xe0, 0x1e, 0x2b,
	0x85, 0xc7, 0xdf, 0x41, 0x48, 0xc5, 0x37, 0x8b, 0xfe, 0xbf, 0x3b, 0xe0, 0x9a, 0x62, 0x7c, 0x03,
	0x4d, 0x4c, 0x32, 0xd5, 0xc1, 0x2b, 0x5a, 0x72, 0x15, 0x9b, 0xfb, 0x7b, 0xd7, 0x1e, 0x63, 0x70,
	0x44, 0xc0, 0xef, 0x6f, 0xf9, 0xc0, 0x0c, 0xc3, 0xff, 0x12, 0xa0, 0x38, 0x5f, 0x49, 0x4b, 0xaf,
	0x32, 0xbd, 0x77, 0x03, 0x1d, 0x71, 0xc8, 0xf6, 0x08, 0x6f, 0xc8, 0x11, 0xdc, 0xd7, 0x71, 0xba,
	0x4a, 0x2c, 0x3d, 0xd0, 0x45, 0x1e, 0x67, 0xa7, 0x56, 0xab, 0xc6, 0x5a, 0x1f, 0x5c, 0xaf, 0xb5,
	0x60, 0x1a, 0xdb, 0x0b, 0x26, 0xa1, 0xf0, 0xdb, 0xfa, 0xf2, 0xe6, 0x26, 0xc5, 0x9c, 0x5a, 0xc5,
	0xea, 0x3a, 0xdb, 0x52, 0xe4, 0x9e, 0xa2, 0x14, 0x99, 0x61, 0xf8, 0x5f, 0x00, 0x60, 0xd5, 0x12,
	0x4b, 0xef, 0x30, 0xfd, 0xfe, 0x65, 0xfa, 0xa6, 0xbd, 0x28, 0xc3, 0xe3, 0x72, 0xd1, 0x5b, 0x81,
	0xcb, 0x55, 0x13, 0xef, 0x02, 0x70, 0xdb, 0x07, 0x17, 0xc6, 0xc2, 0xe3, 0x9d, 0x39, 0xcd, 0x46,
	0xa7, 0x6c, 0x7d, 0xea, 0x32, 0xcf, 0x36, 0xb9, 0xd8, 0x87, 0x7b, 0x61, 0x92, 0xa8, 0xdf, 0x83,
	0x55, 0x98, 0x17, 0x71, 0x98, 0x98, 0xbf, 0x2d, 0x35, 0xd7, 0xa6, 0xe1, 0xdf, 0x65, 0xe7, 0x8f,
	0xc6, 0x37, 0x33, 0xae, 0x5e, 0x0f, 0x1c, 0x2a, 0x34, 0xcd, 0x05, 0x15, 0xba, 0x9c, 0x0b, 0xb2,
	0x7b, 0x7f, 0x57, 0xca, 0xa9, 0x79, 0xc3, 0x71, 0x9e, 0x41, 0xe7, 0x34, 0x97, 0x21, 0xfe, 0xe7,
	0xa0, 0x38, 0x0b, 0xb3, 0x40, 0xe5, 0x81, 0x7c, 0x15, 0x26, 0x76, 0xa4, 0x76, 0xad, 0xef, 0x08,
	0x5d, 0x87, 0xf9, 0x04, 0x1d, 0xe2, 0x3d, 0xb8, 0x7d, 0x91, 0xc0, 0x07, 0xf4, 0xfc, 0xe6, 0x05,
	0xa0, 0x78, 0x0a, 0x22, 0x91, 0x5a, 0x5f, 0x10, 0x5c, 0xa3, 0xa2, 0xc3, 0xc0, 0xb7, 0xc8, 0x53,
	0xca, 0xe1, 0xb6, 0xd8, 0x03, 0x6f, 0x03, 0xb6, 0xdd, 0xdd, 0x28, 0x31, 0xbd, 0xbf, 0x2a, 0xb0,
	0x7b, 0xa5, 0x01, 0xa8, 0x84, 0xff, 0xf7, 0x30, 0x96, 0x90, 0x17, 0xb4, 0x5b, 0x5e, 0x37, 0x5c,
	0x58, 0x5e, 0x88, 0xaf, 0x60, 0x6f, 0xad, 0x65, 0x10, 0x66, 0x51, 0x10, 0xea, 0x20, 0x92, 0xcb,
	0x10, 0x6f, 0xa1, 0x40, 0xad, 0x64, 0x1e, 0xe2, 0x5f, 0xb5, 0xe5, 0x7d, 0x07, 0x21, 0xc3, 0x2c,
	0x1a, 0xea, 0xb1, 0xf1, 0x1f, 0x5a, 0xf7, 0x8d, 0x53, 0xe8, 0xdc, 0x34, 0x85, 0xa3, 0xba, 0x3d,
	0x5b, 0xff, 0x1f, 0x07, 0x9c, 0x85, 0xca, 0x0b, 0xf1, 0x6d, 0xf9, 0x03, 0x34, 0xae, 0xec, 0xd8,
	0x3d, 0xbc, 0xdc, 0x54, 0x84, 0x1c, 0xf0, 0xdd, 0x48, 0x16, 0x75, 0xd6, 0xb2, 0x5c, 0x90, 0x80,
	0x3e, 0x41, 0xb0, 0x11, 0xa8, 0x6e, 0x11, 0x58, 0x10, 0xac, 0x14, 0xd0, 0xe5, 0xa2, 0xd7, 0x04,
	0x6f, 0xe3, 0xe9, 0xfd, 0x89, 0x37, 0xe4, 0x26, 0xd0, 0x9b, 0xba, 0xe3, 0x73, 0x0c, 0x8d, 0x30,
	0xfc, 0x89, 0x11, 0x5e, 0xa2, 0x14, 0xba, 0x75, 0x75, 0xe6, 0x39, 0xf4, 0x21, 0x21, 0x30, 0x28,
	0xda, 0x6c, 0x8a, 0x29, 0xec, 0xa6, 0xb1, 0xd6, 0x34, 0xe5, 0x05, 0xf5, 0x46, 0x2a, 0xb3, 0x82,
	0xcb, 0xdd, 0xda, 0x7f, 0x70, 0xad, 0xc2, 0xcc, 0xa0, 0xfd, 0xb6, 0xa5, 0x1d, 0x95, 0x2c, 0xf1,
	0x19, 0xb0, 0x6e, 0x90, 0xaa, 0x48, 0x72, 0xe9, 0x5b, 0x57, 0xa7, 0xd2, 0x48, 0x20, 0xc0, 0x6f,
	0x10, 0x96, 0x2c, 0x1c, 0xe7, 0x3b, 0x99, 0xd4, 0x85, 0x8c, 0x02, 0xfb, 0x0a, 0xb8, 0xdb, 0xae,
	0xde, 0xdb, 0x06, 0x7b, 0xc0, 0xd0, 0xfe, 0xfb, 0xe0, 0x9a, 0x3c, 0xee, 0x80, 0x37, 0x5c, 0x7c,
	0x37, 0x99, 0x8f, 0xa7, 0xf3, 0xe7, 0xed, 0x5b, 0xf8, 0xe0, 0xc0, 0x78, 0xb2, 0x59, 0x57, 0xfa,
	0x0f, 0xa1, 0x6e, 0x0f, 0x2e, 0x1a, 0xe0, 0xfc, 0x30, 0x5c, 0x1c, 0x21, 0xc8, 0xc3, 0xc7, 0x71,
	0xea, 0xa3, 0x59, 0xe9, 0x4f, 0xc0, 0xe1, 0xb3, 0xa0, 0xcc, 0x8b, 0xf9, 0x78, 0x72, 0x30, 0x9d,
	0x4f, 0xc6, 0x88, 0xa8, 0x43, 0x6d, 0x36, 0x9d, 0xb7, 0x2b, 0x6c, 0x0c, 0x7f, 0x6e, 0x57, 0xc9,
	0x58, 0xbc, 0x98, 0xb5, 0x6b, 0x64, 0x0c, 0x5f, 0x3e, 0x6f, 0x3b, 0x02, 0x60, 0x67, 0x36, 0x19,
	0x4f, 0x87, 0xf3, 0xb6, 0x3b, 0xf2, 0xa0, 0xce, 0x25, 0x38, 0x3e, 0x1f, 0x7d, 0x0d, 0xbb, 0xb1,
	0xba, 0x94, 0xc2, 0xa8, 0x3e, 0x31, 0xc6, 0x2f, 0x8f, 0x4e, 0x95, 0x8e, 0x7e, 0x2d, 0xdd, 0xd1,
	0x95, 0x37, 0xfe, 0x78, 0x87, 0xdf, 0xef, 0x4f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x94, 0x2d,
	0x21, 0x8d, 0x04, 0x08, 0x00, 0x00,
}