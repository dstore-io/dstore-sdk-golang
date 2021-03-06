// Code generated by protoc-gen-go.
// source: dstore/elastic/item/item.proto
// DO NOT EDIT!

/*
Package elastic_item is a generated protocol buffer package.

It is generated from these files:
	dstore/elastic/item/item.proto

It has these top-level messages:
	Item
	Facet
*/
package elastic_item

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gosdk.dstore.de/elastic/elastic"
import dstore_elastic_node "gosdk.dstore.de/elastic/item/elastic_node"
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

type Item struct {
	Node        *dstore_elastic_node.Node     `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	VariantNode []*dstore_elastic_node.Node   `protobuf:"bytes,2,rep,name=variant_node,json=variantNode" json:"variant_node,omitempty"`
	LastUpdated *dstore_values.TimestampValue `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
	Score       float64                       `protobuf:"fixed64,10,opt,name=score" json:"score,omitempty"`
	IndexName   string                        `protobuf:"bytes,11,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	DocumentId  string                        `protobuf:"bytes,12,opt,name=document_id,json=documentId" json:"document_id,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Item) GetNode() *dstore_elastic_node.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Item) GetVariantNode() []*dstore_elastic_node.Node {
	if m != nil {
		return m.VariantNode
	}
	return nil
}

func (m *Item) GetLastUpdated() *dstore_values.TimestampValue {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Facet struct {
	FieldName  string              `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	FacetValue []*Facet_FacetValue `protobuf:"bytes,2,rep,name=facet_value,json=facetValue" json:"facet_value,omitempty"`
	RangeFrom  float64             `protobuf:"fixed64,10,opt,name=range_from,json=rangeFrom" json:"range_from,omitempty"`
	RangeTo    float64             `protobuf:"fixed64,11,opt,name=range_to,json=rangeTo" json:"range_to,omitempty"`
}

func (m *Facet) Reset()                    { *m = Facet{} }
func (m *Facet) String() string            { return proto.CompactTextString(m) }
func (*Facet) ProtoMessage()               {}
func (*Facet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Facet) GetFacetValue() []*Facet_FacetValue {
	if m != nil {
		return m.FacetValue
	}
	return nil
}

type Facet_FacetValue struct {
	Value             *dstore_values.Value `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	TotalItemCount    int32                `protobuf:"varint,2,opt,name=total_item_count,json=totalItemCount" json:"total_item_count,omitempty"`
	MatchingItemCount int32                `protobuf:"varint,3,opt,name=matching_item_count,json=matchingItemCount" json:"matching_item_count,omitempty"`
	Active            bool                 `protobuf:"varint,4,opt,name=active" json:"active,omitempty"`
	MinValue          float64              `protobuf:"fixed64,5,opt,name=min_value,json=minValue" json:"min_value,omitempty"`
	MaxValue          float64              `protobuf:"fixed64,6,opt,name=max_value,json=maxValue" json:"max_value,omitempty"`
}

func (m *Facet_FacetValue) Reset()                    { *m = Facet_FacetValue{} }
func (m *Facet_FacetValue) String() string            { return proto.CompactTextString(m) }
func (*Facet_FacetValue) ProtoMessage()               {}
func (*Facet_FacetValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Facet_FacetValue) GetValue() *dstore_values.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "dstore.elastic.item.Item")
	proto.RegisterType((*Facet)(nil), "dstore.elastic.item.Facet")
	proto.RegisterType((*Facet_FacetValue)(nil), "dstore.elastic.item.Facet.FacetValue")
}

func init() { proto.RegisterFile("dstore/elastic/item/item.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x27, 0xdb, 0xa6, 0xb6, 0x2f, 0x45, 0x74, 0x76, 0x59, 0xb2, 0xd5, 0xd5, 0x65, 0x41, 0xa8,
	0x82, 0x29, 0xe8, 0xd5, 0x83, 0xac, 0x58, 0xd8, 0xcb, 0x22, 0x41, 0x3d, 0x78, 0x09, 0x63, 0x66,
	0x52, 0x07, 0x33, 0x33, 0x25, 0x99, 0x94, 0x7e, 0x07, 0x0f, 0x7e, 0x49, 0x3f, 0x88, 0x33, 0x6f,
	0x26, 0x56, 0x42, 0xd9, 0xcb, 0xd0, 0xf7, 0xfb, 0xf7, 0xde, 0x9b, 0x4c, 0xe1, 0x19, 0x6b, 0x8d,
	0x6e, 0xf8, 0x8a, 0xd7, 0xb4, 0x35, 0xa2, 0x5c, 0x09, 0xc3, 0x25, 0x1e, 0xd9, 0xb6, 0xd1, 0x46,
	0x93, 0x53, 0xcf, 0x67, 0x81, 0xcf, 0x1c, 0xb5, 0x78, 0x3a, 0x30, 0xf5, 0x24, 0x5a, 0x16, 0x47,
	0x23, 0x95, 0x66, 0x3c, 0xf0, 0x21, 0x72, 0xb5, 0xa3, 0x75, 0xc7, 0x5b, 0x0f, 0x5e, 0xff, 0x3a,
	0x81, 0xf1, 0xad, 0x15, 0x92, 0xd7, 0x30, 0x76, 0xda, 0x34, 0xba, 0x8a, 0x96, 0xc9, 0x9b, 0x8b,
	0x6c, 0xd0, 0x1f, 0x73, 0xee, 0xec, 0x91, 0xa3, 0x8c, 0xbc, 0x83, 0xf9, 0x8e, 0x36, 0x82, 0x2a,
	0x53, 0xa0, 0xed, 0xe4, 0x6a, 0x74, 0xbf, 0x2d, 0x09, 0x72, 0x57, 0x90, 0xf7, 0x30, 0x77, 0x82,
	0xa2, 0xdb, 0x32, 0x6a, 0x38, 0x4b, 0x47, 0xd8, 0xf4, 0xb2, 0x77, 0x87, 0x09, 0x8d, 0x90, 0xbc,
	0x35, 0x54, 0x6e, 0xbf, 0xba, 0x3a, 0x4f, 0x9c, 0xe5, 0x8b, 0x77, 0x90, 0x33, 0x88, 0xdb, 0xd2,
	0x6a, 0x53, 0xb0, 0xd6, 0x28, 0xf7, 0x05, 0xb9, 0x04, 0x10, 0x8a, 0xf1, 0x7d, 0xa1, 0xa8, 0xe4,
	0x69, 0x62, 0xa9, 0x59, 0x3e, 0x43, 0xe4, 0xce, 0x02, 0xe4, 0x39, 0x24, 0x4c, 0x97, 0x9d, 0xe4,
	0x76, 0x6a, 0xc1, 0xd2, 0x39, 0xf2, 0xd0, 0x43, 0xb7, 0xec, 0xfa, 0xf7, 0x08, 0xe2, 0x35, 0x2d,
	0xb9, 0x71, 0x49, 0x95, 0xe0, 0x35, 0xf3, 0x49, 0x91, 0x4f, 0x42, 0x04, 0x93, 0xd6, 0x90, 0x54,
	0x4e, 0x57, 0xe0, 0xa8, 0x61, 0xfb, 0x17, 0xd9, 0x91, 0x8f, 0x96, 0x61, 0x9e, 0x3f, 0xfd, 0x1e,
	0x50, 0xfd, 0xfb, 0xed, 0xda, 0x34, 0x54, 0x6d, 0x78, 0x51, 0x35, 0x5a, 0x86, 0x5d, 0x66, 0x88,
	0xac, 0x2d, 0x40, 0x2e, 0x60, 0xea, 0x69, 0xa3, 0x71, 0x9b, 0x28, 0x7f, 0x80, 0xf5, 0x67, 0xbd,
	0xf8, 0x13, 0x01, 0x1c, 0x42, 0xc9, 0x2b, 0x88, 0xfd, 0x28, 0xfe, 0xfb, 0x9d, 0x0d, 0xae, 0xd2,
	0x77, 0xf6, 0x12, 0xb2, 0x84, 0x47, 0x46, 0x1b, 0x5a, 0x17, 0x6e, 0xbe, 0xa2, 0xd4, 0x9d, 0x32,
	0x76, 0x83, 0x68, 0x19, 0xe7, 0x0f, 0x11, 0x77, 0xef, 0xe1, 0x83, 0x43, 0x49, 0x06, 0xa7, 0x92,
	0x9a, 0xf2, 0x87, 0x50, 0x9b, 0xff, 0xc5, 0x23, 0x14, 0x3f, 0xee, 0xa9, 0x83, 0xfe, 0x1c, 0x26,
	0xb4, 0x34, 0x62, 0xc7, 0xd3, 0xb1, 0x95, 0x4c, 0xf3, 0x50, 0x91, 0x27, 0x30, 0x93, 0x42, 0x85,
	0xcb, 0x8a, 0x71, 0x91, 0xa9, 0x05, 0xfc, 0xe8, 0x8e, 0xa4, 0xfb, 0x40, 0x4e, 0x02, 0x49, 0xf7,
	0x48, 0xde, 0x7c, 0x82, 0x73, 0xa1, 0x8f, 0xdd, 0xeb, 0x4d, 0xf2, 0xd1, 0x57, 0xae, 0xfb, 0xb7,
	0x97, 0x1b, 0xdd, 0xb2, 0x9f, 0xbd, 0x8e, 0x0d, 0xfe, 0x04, 0xa1, 0xc0, 0x15, 0xbe, 0x4f, 0xf0,
	0xe1, 0xbf, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xaa, 0x53, 0x1e, 0x82, 0x03, 0x00, 0x00,
}
