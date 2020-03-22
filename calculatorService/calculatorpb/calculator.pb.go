// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculatorService/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Calculating struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculating) Reset()         { *m = Calculating{} }
func (m *Calculating) String() string { return proto.CompactTextString(m) }
func (*Calculating) ProtoMessage()    {}
func (*Calculating) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{0}
}

func (m *Calculating) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculating.Unmarshal(m, b)
}
func (m *Calculating) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculating.Marshal(b, m, deterministic)
}
func (m *Calculating) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculating.Merge(m, src)
}
func (m *Calculating) XXX_Size() int {
	return xxx_messageInfo_Calculating.Size(m)
}
func (m *Calculating) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculating.DiscardUnknown(m)
}

var xxx_messageInfo_Calculating proto.InternalMessageInfo

func (m *Calculating) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Calculating) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type CalculatorRequest struct {
	Calculating          *Calculating `protobuf:"bytes,1,opt,name=calculating,proto3" json:"calculating,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetCalculating() *Calculating {
	if m != nil {
		return m.Calculating
	}
	return nil
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatingManyTimes struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatingManyTimes) Reset()         { *m = CalculatingManyTimes{} }
func (m *CalculatingManyTimes) String() string { return proto.CompactTextString(m) }
func (*CalculatingManyTimes) ProtoMessage()    {}
func (*CalculatingManyTimes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{3}
}

func (m *CalculatingManyTimes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatingManyTimes.Unmarshal(m, b)
}
func (m *CalculatingManyTimes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatingManyTimes.Marshal(b, m, deterministic)
}
func (m *CalculatingManyTimes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatingManyTimes.Merge(m, src)
}
func (m *CalculatingManyTimes) XXX_Size() int {
	return xxx_messageInfo_CalculatingManyTimes.Size(m)
}
func (m *CalculatingManyTimes) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatingManyTimes.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatingManyTimes proto.InternalMessageInfo

func (m *CalculatingManyTimes) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CalculatorManyTimesRequest struct {
	CalculatingManyTimes *CalculatingManyTimes `protobuf:"bytes,1,opt,name=calculatingManyTimes,proto3" json:"calculatingManyTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CalculatorManyTimesRequest) Reset()         { *m = CalculatorManyTimesRequest{} }
func (m *CalculatorManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorManyTimesRequest) ProtoMessage()    {}
func (*CalculatorManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{4}
}

func (m *CalculatorManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorManyTimesRequest.Unmarshal(m, b)
}
func (m *CalculatorManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorManyTimesRequest.Merge(m, src)
}
func (m *CalculatorManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorManyTimesRequest.Size(m)
}
func (m *CalculatorManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorManyTimesRequest proto.InternalMessageInfo

func (m *CalculatorManyTimesRequest) GetCalculatingManyTimes() *CalculatingManyTimes {
	if m != nil {
		return m.CalculatingManyTimes
	}
	return nil
}

type CalculatorManyTimesResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorManyTimesResponse) Reset()         { *m = CalculatorManyTimesResponse{} }
func (m *CalculatorManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorManyTimesResponse) ProtoMessage()    {}
func (*CalculatorManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{5}
}

func (m *CalculatorManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorManyTimesResponse.Unmarshal(m, b)
}
func (m *CalculatorManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorManyTimesResponse.Merge(m, src)
}
func (m *CalculatorManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorManyTimesResponse.Size(m)
}
func (m *CalculatorManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorManyTimesResponse proto.InternalMessageInfo

func (m *CalculatorManyTimesResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type NumberAverage struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberAverage) Reset()         { *m = NumberAverage{} }
func (m *NumberAverage) String() string { return proto.CompactTextString(m) }
func (*NumberAverage) ProtoMessage()    {}
func (*NumberAverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{6}
}

func (m *NumberAverage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberAverage.Unmarshal(m, b)
}
func (m *NumberAverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberAverage.Marshal(b, m, deterministic)
}
func (m *NumberAverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberAverage.Merge(m, src)
}
func (m *NumberAverage) XXX_Size() int {
	return xxx_messageInfo_NumberAverage.Size(m)
}
func (m *NumberAverage) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberAverage.DiscardUnknown(m)
}

var xxx_messageInfo_NumberAverage proto.InternalMessageInfo

func (m *NumberAverage) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type NumberAverageRequest struct {
	NumberAverage        *NumberAverage `protobuf:"bytes,1,opt,name=NumberAverage,proto3" json:"NumberAverage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NumberAverageRequest) Reset()         { *m = NumberAverageRequest{} }
func (m *NumberAverageRequest) String() string { return proto.CompactTextString(m) }
func (*NumberAverageRequest) ProtoMessage()    {}
func (*NumberAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{7}
}

func (m *NumberAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberAverageRequest.Unmarshal(m, b)
}
func (m *NumberAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberAverageRequest.Marshal(b, m, deterministic)
}
func (m *NumberAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberAverageRequest.Merge(m, src)
}
func (m *NumberAverageRequest) XXX_Size() int {
	return xxx_messageInfo_NumberAverageRequest.Size(m)
}
func (m *NumberAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumberAverageRequest proto.InternalMessageInfo

func (m *NumberAverageRequest) GetNumberAverage() *NumberAverage {
	if m != nil {
		return m.NumberAverage
	}
	return nil
}

type NumberAverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberAverageResponse) Reset()         { *m = NumberAverageResponse{} }
func (m *NumberAverageResponse) String() string { return proto.CompactTextString(m) }
func (*NumberAverageResponse) ProtoMessage()    {}
func (*NumberAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{8}
}

func (m *NumberAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberAverageResponse.Unmarshal(m, b)
}
func (m *NumberAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberAverageResponse.Marshal(b, m, deterministic)
}
func (m *NumberAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberAverageResponse.Merge(m, src)
}
func (m *NumberAverageResponse) XXX_Size() int {
	return xxx_messageInfo_NumberAverageResponse.Size(m)
}
func (m *NumberAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NumberAverageResponse proto.InternalMessageInfo

func (m *NumberAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{9}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximum              int32    `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d6609153bff, []int{10}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculating)(nil), "calculator.Calculating")
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
	proto.RegisterType((*CalculatingManyTimes)(nil), "calculator.CalculatingManyTimes")
	proto.RegisterType((*CalculatorManyTimesRequest)(nil), "calculator.CalculatorManyTimesRequest")
	proto.RegisterType((*CalculatorManyTimesResponse)(nil), "calculator.CalculatorManyTimesResponse")
	proto.RegisterType((*NumberAverage)(nil), "calculator.NumberAverage")
	proto.RegisterType((*NumberAverageRequest)(nil), "calculator.NumberAverageRequest")
	proto.RegisterType((*NumberAverageResponse)(nil), "calculator.NumberAverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
}

func init() {
	proto.RegisterFile("calculatorService/calculatorpb/calculator.proto", fileDescriptor_5c8c7d6609153bff)
}

var fileDescriptor_5c8c7d6609153bff = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x49, 0xc4, 0x64, 0x0a, 0x26, 0x0e, 0xa8, 0x58, 0xa3, 0x42, 0x4d, 0x84, 0x03, 0x01,
	0x82, 0xf1, 0xe0, 0xc9, 0x28, 0x89, 0x17, 0x03, 0x87, 0x8a, 0x9a, 0x78, 0x31, 0xa5, 0xac, 0xd8,
	0x84, 0xb6, 0xb8, 0xdb, 0x12, 0xfd, 0x71, 0xcf, 0xc6, 0x76, 0xdb, 0x6e, 0xeb, 0x56, 0x8e, 0xf3,
	0xfa, 0xf6, 0xcd, 0x7b, 0x3b, 0xd3, 0x85, 0xbe, 0x65, 0x2e, 0xad, 0x60, 0x69, 0xfa, 0x1e, 0x7d,
	0x20, 0x74, 0x6d, 0x5b, 0x44, 0x40, 0x56, 0x33, 0xa1, 0xe8, 0xad, 0xa8, 0xe7, 0x7b, 0x08, 0x29,
	0xa2, 0x3f, 0x82, 0x3a, 0xe2, 0x95, 0xed, 0x2e, 0xb0, 0x05, 0x95, 0x37, 0x9b, 0x32, 0xff, 0xd5,
	0x0d, 0x9c, 0x19, 0xa1, 0x0d, 0xa5, 0xa9, 0x74, 0xb6, 0x0c, 0x35, 0xc4, 0x26, 0x21, 0x84, 0x67,
	0x50, 0x65, 0xc4, 0xf2, 0xdc, 0x79, 0xcc, 0x29, 0x85, 0x9c, 0x4a, 0x04, 0x46, 0x24, 0x7d, 0x02,
	0xbb, 0xa3, 0xa4, 0x89, 0x41, 0x3e, 0x02, 0xc2, 0x7c, 0xbc, 0x02, 0xd5, 0x4a, 0x7b, 0x85, 0xda,
	0xea, 0xf0, 0xa0, 0x27, 0xf8, 0x13, 0xac, 0x18, 0x22, 0x57, 0xef, 0x02, 0x8a, 0x7a, 0x6c, 0xe5,
	0xb9, 0x8c, 0xe0, 0x3e, 0x94, 0x29, 0x61, 0xc1, 0xd2, 0xe7, 0x3e, 0x79, 0xa5, 0xf7, 0xa0, 0x2e,
	0x28, 0x8d, 0x4d, 0xf7, 0x6b, 0x6a, 0x3b, 0x84, 0xfd, 0xf2, 0x33, 0xb9, 0x78, 0xa5, 0x53, 0xd0,
	0x52, 0xf5, 0x84, 0x1e, 0xdb, 0x9e, 0x42, 0xdd, 0x92, 0xa8, 0x71, 0xff, 0xcd, 0x02, 0xff, 0xa9,
	0x8c, 0xf4, 0xb4, 0x7e, 0x09, 0x47, 0xd2, 0x9e, 0x1b, 0xa2, 0xb5, 0xa1, 0x1a, 0x5d, 0xf1, 0xcd,
	0x9a, 0x50, 0x73, 0x41, 0x0a, 0x33, 0x3d, 0x43, 0x3d, 0x43, 0x8c, 0xd3, 0x5c, 0xe7, 0x04, 0x78,
	0x8c, 0x43, 0x31, 0x46, 0xf6, 0x60, 0x96, 0xaf, 0xf7, 0x61, 0x2f, 0x27, 0x2c, 0xb5, 0x5c, 0x4a,
	0x2c, 0x77, 0x01, 0xef, 0x6c, 0x77, 0x3e, 0x36, 0x3f, 0x6d, 0x27, 0x70, 0x62, 0x1f, 0x45, 0xbe,
	0xfb, 0x50, 0xcb, 0xb0, 0xb9, 0x78, 0x03, 0xb6, 0x9d, 0x08, 0xe2, 0xfc, 0xb8, 0x1c, 0x7e, 0x97,
	0xc4, 0x5d, 0xe3, 0x7f, 0x00, 0xde, 0x03, 0xa4, 0x20, 0x1e, 0xcb, 0x86, 0x94, 0x2c, 0xa6, 0x76,
	0x52, 0xf4, 0x99, 0x37, 0x7f, 0x87, 0x9a, 0x64, 0x56, 0x78, 0x2e, 0x3f, 0x96, 0x5f, 0x20, 0xad,
	0xbd, 0x91, 0x17, 0xf5, 0x19, 0x28, 0xf8, 0x94, 0x1f, 0x6f, 0xb3, 0x78, 0x2e, 0x5c, 0xbd, 0xf5,
	0x0f, 0x23, 0xd2, 0xed, 0x28, 0x68, 0x80, 0x2a, 0xdc, 0x2a, 0x66, 0x02, 0xff, 0x1d, 0x8e, 0x76,
	0x5a, 0xf8, 0x3d, 0x56, 0x1c, 0x28, 0xb7, 0x3b, 0x2f, 0x15, 0xf1, 0x9d, 0x99, 0x95, 0xc3, 0xd7,
	0xe5, 0xe2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x0d, 0x21, 0xaa, 0x90, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	CalculatorManyTimes(ctx context.Context, in *CalculatorManyTimesRequest, opts ...grpc.CallOption) (CalculatorService_CalculatorManyTimesClient, error)
	NumberAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_NumberAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculatorManyTimes(ctx context.Context, in *CalculatorManyTimesRequest, opts ...grpc.CallOption) (CalculatorService_CalculatorManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/CalculatorManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculatorManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_CalculatorManyTimesClient interface {
	Recv() (*CalculatorManyTimesResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculatorManyTimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculatorManyTimesClient) Recv() (*CalculatorManyTimesResponse, error) {
	m := new(CalculatorManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) NumberAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_NumberAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/NumberAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceNumberAverageClient{stream}
	return x, nil
}

type CalculatorService_NumberAverageClient interface {
	Send(*NumberAverageRequest) error
	CloseAndRecv() (*NumberAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceNumberAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceNumberAverageClient) Send(m *NumberAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceNumberAverageClient) CloseAndRecv() (*NumberAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(NumberAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Calculator(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	CalculatorManyTimes(*CalculatorManyTimesRequest, CalculatorService_CalculatorManyTimesServer) error
	NumberAverage(CalculatorService_NumberAverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Calculator(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculator not implemented")
}
func (*UnimplementedCalculatorServiceServer) CalculatorManyTimes(req *CalculatorManyTimesRequest, srv CalculatorService_CalculatorManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculatorManyTimes not implemented")
}
func (*UnimplementedCalculatorServiceServer) NumberAverage(srv CalculatorService_NumberAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method NumberAverage not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(srv CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculator(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculatorManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalculatorManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).CalculatorManyTimes(m, &calculatorServiceCalculatorManyTimesServer{stream})
}

type CalculatorService_CalculatorManyTimesServer interface {
	Send(*CalculatorManyTimesResponse) error
	grpc.ServerStream
}

type calculatorServiceCalculatorManyTimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculatorManyTimesServer) Send(m *CalculatorManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_NumberAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).NumberAverage(&calculatorServiceNumberAverageServer{stream})
}

type CalculatorService_NumberAverageServer interface {
	SendAndClose(*NumberAverageResponse) error
	Recv() (*NumberAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceNumberAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceNumberAverageServer) SendAndClose(m *NumberAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceNumberAverageServer) Recv() (*NumberAverageRequest, error) {
	m := new(NumberAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculator",
			Handler:    _CalculatorService_Calculator_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalculatorManyTimes",
			Handler:       _CalculatorService_CalculatorManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NumberAverage",
			Handler:       _CalculatorService_NumberAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculatorService/calculatorpb/calculator.proto",
}
