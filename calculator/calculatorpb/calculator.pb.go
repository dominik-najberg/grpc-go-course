// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

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

type Sum struct {
	FirstNumber          float32  `protobuf:"fixed32,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         float32  `protobuf:"fixed32,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetFirstNumber() float32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Sum) GetSecondNumber() float32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type PrimeNumberDecomposition struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecomposition) Reset()         { *m = PrimeNumberDecomposition{} }
func (m *PrimeNumberDecomposition) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposition) ProtoMessage()    {}
func (*PrimeNumberDecomposition) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumberDecomposition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecomposition.Unmarshal(m, b)
}
func (m *PrimeNumberDecomposition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecomposition.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecomposition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecomposition.Merge(m, src)
}
func (m *PrimeNumberDecomposition) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecomposition.Size(m)
}
func (m *PrimeNumberDecomposition) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecomposition.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecomposition proto.InternalMessageInfo

func (m *PrimeNumberDecomposition) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	PrimeNumberDecomposition *PrimeNumberDecomposition `protobuf:"bytes,1,opt,name=primeNumberDecomposition,proto3" json:"primeNumberDecomposition,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetPrimeNumberDecomposition() *PrimeNumberDecomposition {
	if m != nil {
		return m.PrimeNumberDecomposition
	}
	return nil
}

type PrimeNumberDecompositionResponse struct {
	Response             int32    `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetResponse() int32 {
	if m != nil {
		return m.Response
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "calculator.Sum")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecomposition)(nil), "calculator.PrimeNumberDecomposition")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x4f, 0x3a, 0x31,
	0x10, 0xfd, 0x2d, 0xe4, 0x47, 0x74, 0x40, 0x8c, 0x8d, 0x22, 0xe1, 0x22, 0x54, 0x4d, 0x48, 0x34,
	0x40, 0xf0, 0xe2, 0xc9, 0x44, 0xf1, 0xaa, 0x31, 0xcb, 0x4d, 0x0f, 0x0a, 0xeb, 0x68, 0x36, 0xa1,
	0xdb, 0xb5, 0x7f, 0xf8, 0x02, 0x7e, 0x12, 0xbf, 0xa9, 0xd9, 0x6e, 0x0b, 0x15, 0x59, 0xc3, 0xad,
	0x33, 0xfb, 0xde, 0x9b, 0xf7, 0x66, 0x5b, 0xe8, 0x46, 0x93, 0x59, 0xa4, 0x67, 0x13, 0xc5, 0x45,
	0x7f, 0x79, 0x4c, 0xa7, 0x5e, 0xd1, 0x4b, 0x05, 0x57, 0x9c, 0xc0, 0xb2, 0x43, 0xef, 0xa0, 0x3c,
	0xd6, 0x8c, 0x74, 0xa0, 0xf6, 0x16, 0x0b, 0xa9, 0x9e, 0x13, 0xcd, 0xa6, 0x28, 0x9a, 0x41, 0x3b,
	0xe8, 0x96, 0xc2, 0xaa, 0xe9, 0xdd, 0x9b, 0x16, 0x39, 0x86, 0x1d, 0x89, 0x11, 0x4f, 0x5e, 0x1d,
	0xa6, 0x64, 0x30, 0xb5, 0xbc, 0x99, 0x83, 0x68, 0x1f, 0x60, 0xac, 0x59, 0x88, 0x1f, 0x1a, 0xa5,
	0x22, 0x1d, 0x28, 0x4b, 0xcd, 0x8c, 0x58, 0x75, 0xb8, 0xdb, 0xf3, 0x8c, 0x64, 0xa0, 0xec, 0x1b,
	0x3d, 0x85, 0xaa, 0x21, 0xc8, 0x94, 0x27, 0x12, 0x49, 0x03, 0x2a, 0x02, 0xa5, 0x9e, 0x29, 0x43,
	0xda, 0x0e, 0x6d, 0x45, 0x87, 0xd0, 0x7c, 0x10, 0x31, 0xc3, 0x7c, 0xcc, 0x2d, 0x46, 0x9c, 0xa5,
	0x5c, 0xc6, 0x2a, 0xe6, 0x49, 0xc6, 0xf1, 0x5c, 0xff, 0x0f, 0x6d, 0x45, 0x3f, 0x03, 0x38, 0x2a,
	0x22, 0x39, 0x87, 0x2f, 0xd0, 0x4c, 0x0b, 0x20, 0xd6, 0xf6, 0x89, 0x6f, 0xbb, 0x50, 0xae, 0x50,
	0x85, 0x5e, 0x41, 0xbb, 0xd8, 0x84, 0x4d, 0xdd, 0x82, 0x2d, 0x61, 0xcf, 0x36, 0xc3, 0xa2, 0xa6,
	0x7d, 0x38, 0x18, 0x71, 0x96, 0x6a, 0x85, 0xd7, 0x73, 0x14, 0x93, 0x77, 0x74, 0xd6, 0x8b, 0x62,
	0x0f, 0xa0, 0xb1, 0x4a, 0x58, 0xbb, 0xdc, 0xc0, 0x2d, 0x77, 0xf8, 0x55, 0x82, 0xbd, 0xd1, 0x22,
	0xe4, 0x18, 0xc5, 0x3c, 0x8e, 0x90, 0x5c, 0xe6, 0x37, 0xa3, 0xb1, 0xfa, 0xdb, 0xf2, 0xf1, 0xad,
	0xc3, 0x5f, 0x7d, 0x6b, 0xf8, 0x1f, 0xd1, 0xb0, 0xef, 0x72, 0xa2, 0x97, 0x9d, 0x9c, 0x6d, 0xb4,
	0x4a, 0xab, 0x7f, 0xbe, 0x19, 0xd8, 0x0d, 0x1d, 0x04, 0xe4, 0x09, 0xea, 0x3f, 0x83, 0x93, 0x8e,
	0xaf, 0xb1, 0x76, 0x8b, 0x2d, 0xfa, 0x17, 0xc4, 0x89, 0x77, 0x83, 0x9b, 0xfa, 0x63, 0xcd, 0x7f,
	0x54, 0xd3, 0x8a, 0x79, 0x4a, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x92, 0x25, 0x35,
	0x76, 0x03, 0x00, 0x00,
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
	// Unary API
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server Streaming API
	DecomposePrimeNumber(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_DecomposePrimeNumberClient, error)
	// Client Streaming API
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) DecomposePrimeNumber(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_DecomposePrimeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/DecomposePrimeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceDecomposePrimeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_DecomposePrimeNumberClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServiceDecomposePrimeNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceDecomposePrimeNumberClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary API
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server Streaming API
	DecomposePrimeNumber(*PrimeNumberDecompositionRequest, CalculatorService_DecomposePrimeNumberServer) error
	// Client Streaming API
	ComputeAverage(CalculatorService_ComputeAverageServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) DecomposePrimeNumber(req *PrimeNumberDecompositionRequest, srv CalculatorService_DecomposePrimeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method DecomposePrimeNumber not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_DecomposePrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).DecomposePrimeNumber(m, &calculatorServiceDecomposePrimeNumberServer{stream})
}

type CalculatorService_DecomposePrimeNumberServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServiceDecomposePrimeNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceDecomposePrimeNumberServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
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
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DecomposePrimeNumber",
			Handler:       _CalculatorService_DecomposePrimeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
