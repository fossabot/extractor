// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extractor.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExtractRequest struct {
	Separator            string            `protobuf:"bytes,1,opt,name=separator,proto3" json:"separator,omitempty"`
	FileContents         map[string]string `protobuf:"bytes,2,rep,name=fileContents,proto3" json:"fileContents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExtractRequest) Reset()         { *m = ExtractRequest{} }
func (m *ExtractRequest) String() string { return proto.CompactTextString(m) }
func (*ExtractRequest) ProtoMessage()    {}
func (*ExtractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39fd19a8b349f81d, []int{0}
}
func (m *ExtractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractRequest.Unmarshal(m, b)
}
func (m *ExtractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractRequest.Marshal(b, m, deterministic)
}
func (m *ExtractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractRequest.Merge(m, src)
}
func (m *ExtractRequest) XXX_Size() int {
	return xxx_messageInfo_ExtractRequest.Size(m)
}
func (m *ExtractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractRequest proto.InternalMessageInfo

func (m *ExtractRequest) GetSeparator() string {
	if m != nil {
		return m.Separator
	}
	return ""
}

func (m *ExtractRequest) GetFileContents() map[string]string {
	if m != nil {
		return m.FileContents
	}
	return nil
}

type ExtractResponse struct {
	ManagementFiles      []*DependencyManagementFile `protobuf:"bytes,1,rep,name=managementFiles,proto3" json:"managementFiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtractResponse) Reset()         { *m = ExtractResponse{} }
func (m *ExtractResponse) String() string { return proto.CompactTextString(m) }
func (*ExtractResponse) ProtoMessage()    {}
func (*ExtractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39fd19a8b349f81d, []int{1}
}
func (m *ExtractResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractResponse.Unmarshal(m, b)
}
func (m *ExtractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractResponse.Marshal(b, m, deterministic)
}
func (m *ExtractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractResponse.Merge(m, src)
}
func (m *ExtractResponse) XXX_Size() int {
	return xxx_messageInfo_ExtractResponse.Size(m)
}
func (m *ExtractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractResponse proto.InternalMessageInfo

func (m *ExtractResponse) GetManagementFiles() []*DependencyManagementFile {
	if m != nil {
		return m.ManagementFiles
	}
	return nil
}

type MatchRequest struct {
	Separator            string   `protobuf:"bytes,1,opt,name=separator,proto3" json:"separator,omitempty"`
	Paths                []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchRequest) Reset()         { *m = MatchRequest{} }
func (m *MatchRequest) String() string { return proto.CompactTextString(m) }
func (*MatchRequest) ProtoMessage()    {}
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39fd19a8b349f81d, []int{2}
}
func (m *MatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchRequest.Unmarshal(m, b)
}
func (m *MatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchRequest.Marshal(b, m, deterministic)
}
func (m *MatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchRequest.Merge(m, src)
}
func (m *MatchRequest) XXX_Size() int {
	return xxx_messageInfo_MatchRequest.Size(m)
}
func (m *MatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MatchRequest proto.InternalMessageInfo

func (m *MatchRequest) GetSeparator() string {
	if m != nil {
		return m.Separator
	}
	return ""
}

func (m *MatchRequest) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type MatchResponse struct {
	MatchedPaths         []string `protobuf:"bytes,1,rep,name=matchedPaths,proto3" json:"matchedPaths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchResponse) Reset()         { *m = MatchResponse{} }
func (m *MatchResponse) String() string { return proto.CompactTextString(m) }
func (*MatchResponse) ProtoMessage()    {}
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39fd19a8b349f81d, []int{3}
}
func (m *MatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchResponse.Unmarshal(m, b)
}
func (m *MatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchResponse.Marshal(b, m, deterministic)
}
func (m *MatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchResponse.Merge(m, src)
}
func (m *MatchResponse) XXX_Size() int {
	return xxx_messageInfo_MatchResponse.Size(m)
}
func (m *MatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MatchResponse proto.InternalMessageInfo

func (m *MatchResponse) GetMatchedPaths() []string {
	if m != nil {
		return m.MatchedPaths
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtractRequest)(nil), "cloud.deps.dependency_extractor.api.ExtractRequest")
	proto.RegisterMapType((map[string]string)(nil), "cloud.deps.dependency_extractor.api.ExtractRequest.FileContentsEntry")
	proto.RegisterType((*ExtractResponse)(nil), "cloud.deps.dependency_extractor.api.ExtractResponse")
	proto.RegisterType((*MatchRequest)(nil), "cloud.deps.dependency_extractor.api.MatchRequest")
	proto.RegisterType((*MatchResponse)(nil), "cloud.deps.dependency_extractor.api.MatchResponse")
}

func init() { proto.RegisterFile("extractor.proto", fileDescriptor_39fd19a8b349f81d) }

var fileDescriptor_39fd19a8b349f81d = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x8e, 0x02, 0xca, 0x12, 0x08, 0x2c, 0x15, 0xb2, 0x0c, 0x48, 0xc5, 0xbd, 0x54,
	0x95, 0xb0, 0xd5, 0x84, 0x03, 0x8a, 0x84, 0x90, 0x0a, 0xe1, 0x56, 0x09, 0xf9, 0xd8, 0x4b, 0x35,
	0xb5, 0x07, 0x67, 0x45, 0xb2, 0xbb, 0xec, 0x4e, 0x0a, 0xe6, 0xc8, 0x1b, 0x20, 0x4e, 0x3c, 0x10,
	0x4f, 0xc0, 0x8d, 0x33, 0x0f, 0x82, 0xb2, 0x5e, 0x92, 0x86, 0x5e, 0x12, 0x2e, 0xd6, 0xce, 0xec,
	0xfc, 0xff, 0x7c, 0x33, 0xf2, 0xf2, 0x01, 0x7e, 0x22, 0x0b, 0x25, 0x69, 0x9b, 0x19, 0xab, 0x49,
	0x8b, 0x83, 0x72, 0xa6, 0x17, 0x55, 0x56, 0xa1, 0x71, 0xcb, 0x0f, 0xaa, 0x0a, 0x55, 0xd9, 0x9c,
	0xaf, 0xcb, 0xc0, 0xc8, 0x84, 0xfb, 0x6b, 0x2f, 0x48, 0xce, 0x6a, 0x49, 0xd3, 0xc5, 0x45, 0x56,
	0xea, 0x79, 0x5e, 0x5b, 0x53, 0x3e, 0xc5, 0x52, 0xbb, 0xc6, 0x11, 0x86, 0xb0, 0x06, 0xc2, 0x8f,
	0xd0, 0xe4, 0x34, 0x95, 0xb6, 0x3a, 0x37, 0x60, 0xa9, 0xc9, 0x6b, 0xad, 0xeb, 0x19, 0x82, 0x91,
	0x2e, 0x1c, 0x73, 0x30, 0x32, 0x07, 0xa5, 0x34, 0x01, 0x49, 0xad, 0x82, 0x77, 0xfa, 0x8b, 0xf1,
	0x3b, 0x93, 0xb6, 0x73, 0x81, 0x1f, 0x16, 0xe8, 0x48, 0x3c, 0xe2, 0x3d, 0x87, 0x06, 0x2c, 0x90,
	0xb6, 0x31, 0xdb, 0x67, 0x87, 0xbd, 0x62, 0x9d, 0x10, 0x92, 0xf7, 0xdf, 0xc9, 0x19, 0xbe, 0xd2,
	0x8a, 0x50, 0x91, 0x8b, 0xa3, 0xfd, 0xce, 0xe1, 0xad, 0xe1, 0x24, 0xdb, 0x62, 0xa8, 0x6c, 0xb3,
	0x51, 0xf6, 0xe6, 0x8a, 0xcf, 0x44, 0x91, 0x6d, 0x8a, 0x0d, 0xeb, 0xe4, 0x25, 0xbf, 0x77, 0xad,
	0x44, 0xdc, 0xe5, 0x9d, 0xf7, 0xd8, 0x04, 0xae, 0xe5, 0x51, 0xec, 0xf1, 0xee, 0x25, 0xcc, 0x16,
	0x18, 0x47, 0x3e, 0xd7, 0x06, 0xe3, 0xe8, 0x39, 0x4b, 0x3f, 0xf3, 0xc1, 0xaa, 0xa5, 0x33, 0x5a,
	0x39, 0x14, 0x35, 0x1f, 0xcc, 0x41, 0x41, 0x8d, 0x73, 0x54, 0xb4, 0x74, 0x77, 0x31, 0xf3, 0x13,
	0xbc, 0xd8, 0x6a, 0x82, 0xd7, 0xab, 0x8b, 0xd3, 0x0d, 0x97, 0xe2, 0x5f, 0xd7, 0xf4, 0x84, 0xf7,
	0x4f, 0x81, 0xca, 0xe9, 0x76, 0x5b, 0xdd, 0xe3, 0x5d, 0x03, 0x34, 0x6d, 0xd7, 0xd9, 0x2b, 0xda,
	0x20, 0x1d, 0xf1, 0xdb, 0xc1, 0x23, 0xd0, 0xa7, 0xbc, 0x3f, 0x5f, 0x26, 0xb0, 0x7a, 0xeb, 0xab,
	0x99, 0xaf, 0xde, 0xc8, 0x0d, 0x7f, 0x44, 0xfc, 0xfe, 0x1a, 0x73, 0xf2, 0x17, 0x5f, 0x7c, 0x65,
	0xbc, 0xeb, 0xdd, 0xc4, 0xf1, 0x56, 0xa3, 0x5e, 0xa5, 0x4f, 0x86, 0xbb, 0x48, 0x5a, 0xd8, 0xf4,
	0xc9, 0x97, 0x9f, 0xbf, 0xbf, 0x45, 0x0f, 0xd3, 0x07, 0xf9, 0xe5, 0x71, 0xbe, 0xd2, 0x48, 0x74,
	0xb9, 0xe7, 0x1d, 0xb3, 0x23, 0xf1, 0x9d, 0xf1, 0x9b, 0x81, 0x50, 0x8c, 0xfe, 0xe3, 0x17, 0x4a,
	0x9e, 0xed, 0x26, 0x0a, 0x64, 0x07, 0x9e, 0xec, 0x71, 0x1a, 0x5f, 0x23, 0x0b, 0xba, 0x31, 0x3b,
	0x3a, 0xe9, 0x9e, 0x75, 0xc0, 0xc8, 0x8b, 0x1b, 0xfe, 0x9d, 0x8c, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x46, 0x9d, 0x84, 0xdc, 0xc7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DependencyExtractorClient is the client API for DependencyExtractor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DependencyExtractorClient interface {
	Match(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (*MatchResponse, error)
	Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error)
}

type dependencyExtractorClient struct {
	cc *grpc.ClientConn
}

func NewDependencyExtractorClient(cc *grpc.ClientConn) DependencyExtractorClient {
	return &dependencyExtractorClient{cc}
}

func (c *dependencyExtractorClient) Match(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (*MatchResponse, error) {
	out := new(MatchResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.dependency_extractor.api.DependencyExtractor/Match", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dependencyExtractorClient) Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error) {
	out := new(ExtractResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.dependency_extractor.api.DependencyExtractor/Extract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DependencyExtractorServer is the server API for DependencyExtractor service.
type DependencyExtractorServer interface {
	Match(context.Context, *MatchRequest) (*MatchResponse, error)
	Extract(context.Context, *ExtractRequest) (*ExtractResponse, error)
}

// UnimplementedDependencyExtractorServer can be embedded to have forward compatible implementations.
type UnimplementedDependencyExtractorServer struct {
}

func (*UnimplementedDependencyExtractorServer) Match(ctx context.Context, req *MatchRequest) (*MatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (*UnimplementedDependencyExtractorServer) Extract(ctx context.Context, req *ExtractRequest) (*ExtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extract not implemented")
}

func RegisterDependencyExtractorServer(s *grpc.Server, srv DependencyExtractorServer) {
	s.RegisterService(&_DependencyExtractor_serviceDesc, srv)
}

func _DependencyExtractor_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependencyExtractorServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.dependency_extractor.api.DependencyExtractor/Match",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependencyExtractorServer).Match(ctx, req.(*MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DependencyExtractor_Extract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependencyExtractorServer).Extract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.dependency_extractor.api.DependencyExtractor/Extract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependencyExtractorServer).Extract(ctx, req.(*ExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DependencyExtractor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.deps.dependency_extractor.api.DependencyExtractor",
	HandlerType: (*DependencyExtractorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Match",
			Handler:    _DependencyExtractor_Match_Handler,
		},
		{
			MethodName: "Extract",
			Handler:    _DependencyExtractor_Extract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extractor.proto",
}
