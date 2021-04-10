// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package models

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
//const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RequestData struct {
	MetricType   string         `protobuf:"bytes,2,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
	Metrics      []*MetricModel `protobuf:"bytes,3,rep,name=metrics,proto3" json:"metrics,omitempty"`
	CommonLabels []*LabelModel  `protobuf:"bytes,4,rep,name=common_labels,json=commonLabels,proto3" json:"common_labels,omitempty"`
}

func (m *RequestData) Reset()         { *m = RequestData{} }
func (m *RequestData) String() string { return proto.CompactTextString(m) }
func (*RequestData) ProtoMessage()    {}
func (*RequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}
func (m *RequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestData.Unmarshal(m, b)
}
func (m *RequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestData.Marshal(b, m, deterministic)
}
func (m *RequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestData.Merge(m, src)
}
func (m *RequestData) XXX_Size() int {
	return xxx_messageInfo_RequestData.Size(m)
}
func (m *RequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestData.DiscardUnknown(m)
}

var xxx_messageInfo_RequestData proto.InternalMessageInfo

func (m *RequestData) GetMetricType() string {
	if m != nil {
		return m.MetricType
	}
	return ""
}

func (m *RequestData) GetMetrics() []*MetricModel {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *RequestData) GetCommonLabels() []*LabelModel {
	if m != nil {
		return m.CommonLabels
	}
	return nil
}

type MetricModel struct {
	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value  float64       `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Labels []*LabelModel `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (m *MetricModel) Reset()         { *m = MetricModel{} }
func (m *MetricModel) String() string { return proto.CompactTextString(m) }
func (*MetricModel) ProtoMessage()    {}
func (*MetricModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}
func (m *MetricModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricModel.Unmarshal(m, b)
}
func (m *MetricModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricModel.Marshal(b, m, deterministic)
}
func (m *MetricModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricModel.Merge(m, src)
}
func (m *MetricModel) XXX_Size() int {
	return xxx_messageInfo_MetricModel.Size(m)
}
func (m *MetricModel) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricModel.DiscardUnknown(m)
}

var xxx_messageInfo_MetricModel proto.InternalMessageInfo

func (m *MetricModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricModel) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MetricModel) GetLabels() []*LabelModel {
	if m != nil {
		return m.Labels
	}
	return nil
}

type LabelModel struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *LabelModel) Reset()         { *m = LabelModel{} }
func (m *LabelModel) String() string { return proto.CompactTextString(m) }
func (*LabelModel) ProtoMessage()    {}
func (*LabelModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}
func (m *LabelModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelModel.Unmarshal(m, b)
}
func (m *LabelModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelModel.Marshal(b, m, deterministic)
}
func (m *LabelModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelModel.Merge(m, src)
}
func (m *LabelModel) XXX_Size() int {
	return xxx_messageInfo_LabelModel.Size(m)
}
func (m *LabelModel) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelModel.DiscardUnknown(m)
}

var xxx_messageInfo_LabelModel proto.InternalMessageInfo

func (m *LabelModel) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelModel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestData)(nil), "models.RequestData")
	proto.RegisterType((*MetricModel)(nil), "models.MetricModel")
	proto.RegisterType((*LabelModel)(nil), "models.LabelModel")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x1d, 0x52, 0x82, 0xfa, 0xa7, 0x48, 0xc8, 0x30, 0x44, 0x0c, 0xa6, 0xca, 0x54, 0x21,
	0x35, 0x95, 0x00, 0x89, 0x1d, 0x31, 0xd2, 0x25, 0x62, 0xaf, 0x1c, 0x63, 0xd2, 0x8a, 0xb8, 0x0e,
	0xb5, 0x83, 0x94, 0xb7, 0x60, 0xe2, 0x99, 0x18, 0x3b, 0x32, 0xa2, 0xe4, 0x45, 0x50, 0x7e, 0xa7,
	0x02, 0x86, 0x6e, 0x77, 0x97, 0xbb, 0x7c, 0xbf, 0x0c, 0x23, 0x23, 0x96, 0x52, 0xf1, 0xa4, 0xdc,
	0x68, 0xab, 0x69, 0xa0, 0xf4, 0x93, 0x2c, 0xcc, 0xf9, 0x34, 0x5f, 0xd9, 0x65, 0x95, 0x25, 0x42,
	0xab, 0x59, 0xae, 0x73, 0x3d, 0xc3, 0xcf, 0x59, 0xf5, 0x8c, 0x0e, 0x0d, 0x2a, 0x37, 0x8b, 0x3f,
	0x3c, 0x08, 0x53, 0xf9, 0x5a, 0x49, 0x63, 0xef, 0xb9, 0xe5, 0xf4, 0x02, 0x42, 0x25, 0xed, 0x66,
	0x25, 0x16, 0xb6, 0x2e, 0x65, 0x74, 0x30, 0xf6, 0x26, 0xc3, 0x14, 0x5c, 0xf4, 0x58, 0x97, 0x92,
	0x4e, 0xe1, 0xc8, 0x39, 0x13, 0xf9, 0x63, 0x7f, 0x12, 0x5e, 0x9d, 0x26, 0x8e, 0x9c, 0xcc, 0x31,
	0x9e, 0x77, 0x26, 0xdd, 0x75, 0xe8, 0x2d, 0x1c, 0x0b, 0xad, 0x94, 0x5e, 0x2f, 0x0a, 0x9e, 0xc9,
	0xc2, 0x44, 0x03, 0x1c, 0xd1, 0xdd, 0xe8, 0xa1, 0x4b, 0xdd, 0x66, 0xe4, 0x8a, 0x98, 0x98, 0x58,
	0x40, 0xf8, 0xe7, 0x87, 0x94, 0xc2, 0x60, 0xcd, 0x95, 0x8c, 0x3c, 0x3c, 0x08, 0x35, 0x3d, 0x83,
	0xc3, 0x37, 0x5e, 0x54, 0xee, 0x4a, 0x2f, 0x75, 0x86, 0x5e, 0x42, 0xd0, 0xa3, 0xfc, 0xbd, 0xa8,
	0xbe, 0x11, 0xdf, 0x00, 0xfc, 0xa6, 0xf4, 0x04, 0xfc, 0x17, 0x59, 0xf7, 0x88, 0x4e, 0xfe, 0x27,
	0x0c, 0x7b, 0xc2, 0x5d, 0xf4, 0xd9, 0x30, 0xb2, 0x6d, 0x18, 0xf9, 0x6e, 0x18, 0x79, 0x6f, 0x19,
	0xd9, 0xb6, 0x8c, 0x7c, 0xb5, 0x8c, 0x64, 0x01, 0x3e, 0xea, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xee, 0x79, 0x2e, 0x47, 0x9b, 0x01, 0x00, 0x00,
}
