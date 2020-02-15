// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/service.proto

package service

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Account_AccountMode int32

const (
	Account_ACCOUNT_MODE_DEBT       Account_AccountMode = 0
	Account_ACCOUNT_MODE_INVESTMENT Account_AccountMode = 1
)

var Account_AccountMode_name = map[int32]string{
	0: "ACCOUNT_MODE_DEBT",
	1: "ACCOUNT_MODE_INVESTMENT",
}

var Account_AccountMode_value = map[string]int32{
	"ACCOUNT_MODE_DEBT":       0,
	"ACCOUNT_MODE_INVESTMENT": 1,
}

func (x Account_AccountMode) String() string {
	return proto.EnumName(Account_AccountMode_name, int32(x))
}

func (Account_AccountMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e51e679f9ae460e2, []int{0, 0}
}

type Account struct {
	// Not all accounts accrue interest at the same rate. Some may accrue interest once per year, once per month, etc. Since we use periods we need to know how many periods before interest is calculated. So, on a bi-weekly paycheck you might choose to view interest every 2 periods.
	AddInterestEveryNPeriods int64 `protobuf:"varint,1,opt,name=AddInterestEveryNPeriods,proto3" json:"AddInterestEveryNPeriods,omitempty"`
	// The Balance of the account. 45.05
	Balance float64 `protobuf:"fixed64,2,opt,name=Balance,proto3" json:"Balance,omitempty"`
	// The interest rate of hte account. %5.5 is 0.055.
	InterestRate float64 `protobuf:"fixed64,3,opt,name=InterestRate,proto3" json:"InterestRate,omitempty"`
	// The mode determines how contributions work for the account. A debt account contribution will remove money from the balance. An investment account contribution will add money to the balance. This value also affects how interest is calculated.
	Mode                 Account_AccountMode `protobuf:"varint,4,opt,name=Mode,proto3,enum=financialplanningcalculator.Account_AccountMode" json:"Mode,omitempty"`
	Name                 string              `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51e679f9ae460e2, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddInterestEveryNPeriods() int64 {
	if m != nil {
		return m.AddInterestEveryNPeriods
	}
	return 0
}

func (m *Account) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetInterestRate() float64 {
	if m != nil {
		return m.InterestRate
	}
	return 0
}

func (m *Account) GetMode() Account_AccountMode {
	if m != nil {
		return m.Mode
	}
	return Account_ACCOUNT_MODE_DEBT
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CalculateResponse struct {
	Periods              *Period  `protobuf:"bytes,1,opt,name=Periods,proto3" json:"Periods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateResponse) Reset()         { *m = CalculateResponse{} }
func (m *CalculateResponse) String() string { return proto.CompactTextString(m) }
func (*CalculateResponse) ProtoMessage()    {}
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51e679f9ae460e2, []int{1}
}

func (m *CalculateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateResponse.Unmarshal(m, b)
}
func (m *CalculateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateResponse.Marshal(b, m, deterministic)
}
func (m *CalculateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateResponse.Merge(m, src)
}
func (m *CalculateResponse) XXX_Size() int {
	return xxx_messageInfo_CalculateResponse.Size(m)
}
func (m *CalculateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateResponse proto.InternalMessageInfo

func (m *CalculateResponse) GetPeriods() *Period {
	if m != nil {
		return m.Periods
	}
	return nil
}

type Period struct {
	Accounts             *Account `protobuf:"bytes,1,opt,name=Accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51e679f9ae460e2, []int{2}
}

func (m *Period) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Period.Unmarshal(m, b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Period.Marshal(b, m, deterministic)
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return xxx_messageInfo_Period.Size(m)
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetAccounts() *Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterEnum("financialplanningcalculator.Account_AccountMode", Account_AccountMode_name, Account_AccountMode_value)
	proto.RegisterType((*Account)(nil), "financialplanningcalculator.Account")
	proto.RegisterType((*CalculateResponse)(nil), "financialplanningcalculator.CalculateResponse")
	proto.RegisterType((*Period)(nil), "financialplanningcalculator.Period")
}

func init() { proto.RegisterFile("service/service.proto", fileDescriptor_e51e679f9ae460e2) }

var fileDescriptor_e51e679f9ae460e2 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xab, 0xda, 0x40,
	0x10, 0xc7, 0xdf, 0xbe, 0x67, 0x4d, 0x1d, 0x4b, 0xd1, 0x05, 0xdb, 0xa0, 0x97, 0x90, 0xf6, 0x90,
	0x53, 0x2c, 0xf6, 0x56, 0x28, 0x34, 0xc6, 0xb4, 0x58, 0x30, 0xca, 0x1a, 0x7b, 0x28, 0x05, 0x59,
	0x93, 0x51, 0x02, 0x71, 0x37, 0x24, 0xab, 0xe0, 0xb5, 0x7f, 0x49, 0xff, 0xd4, 0x62, 0x7e, 0x88,
	0x52, 0x2a, 0xef, 0xb4, 0x93, 0x99, 0xf9, 0x4c, 0xbe, 0x33, 0x7c, 0xa1, 0x97, 0x63, 0x76, 0x8c,
	0x43, 0x1c, 0x56, 0xaf, 0x9d, 0x66, 0x52, 0x49, 0x3a, 0xd8, 0xc6, 0x82, 0x8b, 0x30, 0xe6, 0x49,
	0x9a, 0x70, 0x21, 0x62, 0xb1, 0x0b, 0x79, 0x12, 0x1e, 0x12, 0xae, 0x64, 0xd6, 0x1f, 0xec, 0xa4,
	0xdc, 0x25, 0x38, 0x2c, 0x5a, 0x37, 0x87, 0xed, 0x10, 0xf7, 0xa9, 0x3a, 0x95, 0xa4, 0xf9, 0xe7,
	0x11, 0x34, 0x27, 0x0c, 0xe5, 0x41, 0x28, 0xfa, 0x09, 0x74, 0x27, 0x8a, 0xa6, 0x42, 0x61, 0x86,
	0xb9, 0xf2, 0x8e, 0x98, 0x9d, 0xfc, 0x05, 0x66, 0xb1, 0x8c, 0x72, 0x9d, 0x18, 0xc4, 0x7a, 0x62,
	0xff, 0xad, 0x53, 0x1d, 0xb4, 0x31, 0x4f, 0xb8, 0x08, 0x51, 0x7f, 0x34, 0x88, 0x45, 0x58, 0xfd,
	0x49, 0x4d, 0x78, 0x55, 0x23, 0x8c, 0x2b, 0xd4, 0x9f, 0x8a, 0xf2, 0x4d, 0x8e, 0x4e, 0xa0, 0x31,
	0x93, 0x11, 0xea, 0x0d, 0x83, 0x58, 0xaf, 0x47, 0x1f, 0xec, 0x3b, 0xeb, 0xd8, 0x95, 0xda, 0xfa,
	0x3d, 0x73, 0xac, 0xa0, 0x29, 0x85, 0x86, 0xcf, 0xf7, 0xa8, 0xbf, 0x30, 0x88, 0xd5, 0x62, 0x45,
	0x6c, 0x3a, 0xd0, 0xbe, 0x6a, 0xa4, 0x3d, 0xe8, 0x3a, 0xae, 0x3b, 0x5f, 0xf9, 0xc1, 0x7a, 0x36,
	0x9f, 0x78, 0xeb, 0x89, 0x37, 0x0e, 0x3a, 0x0f, 0x74, 0x00, 0x6f, 0x6f, 0xd2, 0x53, 0xff, 0x87,
	0xb7, 0x0c, 0x66, 0x9e, 0x1f, 0x74, 0x88, 0xc9, 0xa0, 0xeb, 0x56, 0xbf, 0x47, 0x86, 0x79, 0x2a,
	0x45, 0x8e, 0xf4, 0x33, 0x68, 0xd7, 0xa7, 0x69, 0x8f, 0xde, 0xdd, 0x15, 0x5d, 0xf6, 0xb2, 0x9a,
	0x31, 0xbf, 0x43, 0xb3, 0x0c, 0xe9, 0x17, 0x78, 0x59, 0x09, 0xac, 0x27, 0xbd, 0x7f, 0xce, 0xfa,
	0xec, 0x42, 0x8d, 0x7e, 0x13, 0x30, 0xbf, 0xd6, 0xc4, 0xa2, 0x22, 0xdc, 0x0b, 0xb1, 0x2c, 0x9d,
	0x42, 0x7f, 0x41, 0xe7, 0x1b, 0xaa, 0x55, 0x8e, 0xd9, 0x65, 0x1b, 0xfa, 0xc6, 0x2e, 0xbd, 0x61,
	0xd7, 0xde, 0xb0, 0xbd, 0xb3, 0x37, 0xfa, 0xf6, 0x5d, 0x09, 0xff, 0x5c, 0xc3, 0x7c, 0x18, 0xb7,
	0x7e, 0x6a, 0x95, 0x25, 0x37, 0xcd, 0x62, 0xd8, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0x99, 0x7e, 0xe3, 0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinancialPlanningCalculatorServiceClient is the client API for FinancialPlanningCalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinancialPlanningCalculatorServiceClient interface {
	// Calculate for the current user.
	GetUserCalculate(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type financialPlanningCalculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewFinancialPlanningCalculatorServiceClient(cc *grpc.ClientConn) FinancialPlanningCalculatorServiceClient {
	return &financialPlanningCalculatorServiceClient{cc}
}

func (c *financialPlanningCalculatorServiceClient) GetUserCalculate(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/financialplanningcalculator.FinancialPlanningCalculatorService/GetUserCalculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancialPlanningCalculatorServiceServer is the server API for FinancialPlanningCalculatorService service.
type FinancialPlanningCalculatorServiceServer interface {
	// Calculate for the current user.
	GetUserCalculate(context.Context, *empty.Empty) (*CalculateResponse, error)
}

// UnimplementedFinancialPlanningCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFinancialPlanningCalculatorServiceServer struct {
}

func (*UnimplementedFinancialPlanningCalculatorServiceServer) GetUserCalculate(ctx context.Context, req *empty.Empty) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCalculate not implemented")
}

func RegisterFinancialPlanningCalculatorServiceServer(s *grpc.Server, srv FinancialPlanningCalculatorServiceServer) {
	s.RegisterService(&_FinancialPlanningCalculatorService_serviceDesc, srv)
}

func _FinancialPlanningCalculatorService_GetUserCalculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialPlanningCalculatorServiceServer).GetUserCalculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financialplanningcalculator.FinancialPlanningCalculatorService/GetUserCalculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialPlanningCalculatorServiceServer).GetUserCalculate(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinancialPlanningCalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "financialplanningcalculator.FinancialPlanningCalculatorService",
	HandlerType: (*FinancialPlanningCalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserCalculate",
			Handler:    _FinancialPlanningCalculatorService_GetUserCalculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}