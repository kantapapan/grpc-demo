// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sphere.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Sphere ...
type Sphere struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Job                  string   `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	License              string   `protobuf:"bytes,3,opt,name=license,proto3" json:"license,omitempty"`
	Skill                string   `protobuf:"bytes,4,opt,name=skill,proto3" json:"skill,omitempty"`
	Time                 int32    `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sphere) Reset()         { *m = Sphere{} }
func (m *Sphere) String() string { return proto.CompactTextString(m) }
func (*Sphere) ProtoMessage()    {}
func (*Sphere) Descriptor() ([]byte, []int) {
	return fileDescriptor_969a4ec4166c6933, []int{0}
}

func (m *Sphere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sphere.Unmarshal(m, b)
}
func (m *Sphere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sphere.Marshal(b, m, deterministic)
}
func (m *Sphere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sphere.Merge(m, src)
}
func (m *Sphere) XXX_Size() int {
	return xxx_messageInfo_Sphere.Size(m)
}
func (m *Sphere) XXX_DiscardUnknown() {
	xxx_messageInfo_Sphere.DiscardUnknown(m)
}

var xxx_messageInfo_Sphere proto.InternalMessageInfo

func (m *Sphere) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Sphere) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *Sphere) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

func (m *Sphere) GetSkill() string {
	if m != nil {
		return m.Skill
	}
	return ""
}

func (m *Sphere) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

// SphereList a collection of speries
type SphereList struct {
	Items                []*Sphere `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SphereList) Reset()         { *m = SphereList{} }
func (m *SphereList) String() string { return proto.CompactTextString(m) }
func (*SphereList) ProtoMessage()    {}
func (*SphereList) Descriptor() ([]byte, []int) {
	return fileDescriptor_969a4ec4166c6933, []int{1}
}

func (m *SphereList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SphereList.Unmarshal(m, b)
}
func (m *SphereList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SphereList.Marshal(b, m, deterministic)
}
func (m *SphereList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SphereList.Merge(m, src)
}
func (m *SphereList) XXX_Size() int {
	return xxx_messageInfo_SphereList.Size(m)
}
func (m *SphereList) XXX_DiscardUnknown() {
	xxx_messageInfo_SphereList.DiscardUnknown(m)
}

var xxx_messageInfo_SphereList proto.InternalMessageInfo

func (m *SphereList) GetItems() []*Sphere {
	if m != nil {
		return m.Items
	}
	return nil
}

type SphereJobMessage struct {
	TargetJob            string   `protobuf:"bytes,1,opt,name=target_job,json=targetJob,proto3" json:"target_job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SphereJobMessage) Reset()         { *m = SphereJobMessage{} }
func (m *SphereJobMessage) String() string { return proto.CompactTextString(m) }
func (*SphereJobMessage) ProtoMessage()    {}
func (*SphereJobMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_969a4ec4166c6933, []int{2}
}

func (m *SphereJobMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SphereJobMessage.Unmarshal(m, b)
}
func (m *SphereJobMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SphereJobMessage.Marshal(b, m, deterministic)
}
func (m *SphereJobMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SphereJobMessage.Merge(m, src)
}
func (m *SphereJobMessage) XXX_Size() int {
	return xxx_messageInfo_SphereJobMessage.Size(m)
}
func (m *SphereJobMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SphereJobMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SphereJobMessage proto.InternalMessageInfo

func (m *SphereJobMessage) GetTargetJob() string {
	if m != nil {
		return m.TargetJob
	}
	return ""
}

func init() {
	proto.RegisterType((*Sphere)(nil), "proto.Sphere")
	proto.RegisterType((*SphereList)(nil), "proto.SphereList")
	proto.RegisterType((*SphereJobMessage)(nil), "proto.SphereJobMessage")
}

func init() { proto.RegisterFile("sphere.proto", fileDescriptor_969a4ec4166c6933) }

var fileDescriptor_969a4ec4166c6933 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0xa4, 0x5b, 0xe9, 0x68, 0x4b, 0x1d, 0x04, 0x17, 0x41, 0x08, 0xf1, 0x92, 0x53,
	0xa1, 0xf5, 0xee, 0xc1, 0x8b, 0x10, 0x14, 0x21, 0xfd, 0x01, 0xd2, 0x6d, 0x87, 0x3a, 0x9a, 0x76,
	0x4b, 0x76, 0x11, 0xfc, 0xf7, 0x92, 0x19, 0x44, 0x72, 0xda, 0x99, 0xef, 0xed, 0xbc, 0xf7, 0xe0,
	0x32, 0x9e, 0x3e, 0xa8, 0xa3, 0xc5, 0xa9, 0x0b, 0x29, 0xa0, 0x95, 0xa7, 0x3c, 0xc2, 0x78, 0x2d,
	0x18, 0x67, 0x90, 0xf1, 0xce, 0x99, 0xc2, 0x54, 0xb6, 0xc9, 0x78, 0x87, 0x73, 0xc8, 0x3f, 0x83,
	0x77, 0x59, 0x61, 0xaa, 0x49, 0xd3, 0x8f, 0xe8, 0xe0, 0xbc, 0xe5, 0x2d, 0x1d, 0x23, 0xb9, 0x5c,
	0xe8, 0xdf, 0x8a, 0xd7, 0x60, 0xe3, 0x17, 0xb7, 0xad, 0x1b, 0x09, 0xd7, 0x05, 0x11, 0x46, 0x89,
	0x0f, 0xe4, 0xac, 0x78, 0xca, 0x5c, 0x2e, 0x01, 0x34, 0xef, 0x85, 0x63, 0xc2, 0x7b, 0xb0, 0x9c,
	0xe8, 0x10, 0x9d, 0x29, 0xf2, 0xea, 0x62, 0x35, 0xd5, 0x6e, 0x0b, 0xfd, 0xd1, 0xa8, 0x56, 0x2e,
	0x61, 0xae, 0xa0, 0x0e, 0xfe, 0x95, 0x62, 0xdc, 0xec, 0x09, 0xef, 0x00, 0xd2, 0xa6, 0xdb, 0x53,
	0x7a, 0xef, 0x3b, 0x1a, 0x49, 0x9d, 0x28, 0xa9, 0x83, 0x5f, 0xbd, 0xc1, 0x54, 0x4f, 0xd6, 0xd4,
	0x7d, 0xf3, 0x96, 0xf0, 0x11, 0x66, 0xcf, 0x94, 0x94, 0x3d, 0xfd, 0xd4, 0xc1, 0xe3, 0xcd, 0x20,
	0xeb, 0xdf, 0xfa, 0xf6, 0x6a, 0x20, 0xf4, 0x35, 0xcb, 0x33, 0x3f, 0x16, 0xf6, 0xf0, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x32, 0xef, 0x7f, 0x80, 0x44, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SphereServiceClient is the client API for SphereService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SphereServiceClient interface {
	GetSphereByJob(ctx context.Context, in *SphereJobMessage, opts ...grpc.CallOption) (*SphereList, error)
}

type sphereServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSphereServiceClient(cc grpc.ClientConnInterface) SphereServiceClient {
	return &sphereServiceClient{cc}
}

func (c *sphereServiceClient) GetSphereByJob(ctx context.Context, in *SphereJobMessage, opts ...grpc.CallOption) (*SphereList, error) {
	out := new(SphereList)
	err := c.cc.Invoke(ctx, "/proto.SphereService/GetSphereByJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SphereServiceServer is the server API for SphereService service.
type SphereServiceServer interface {
	GetSphereByJob(context.Context, *SphereJobMessage) (*SphereList, error)
}

// UnimplementedSphereServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSphereServiceServer struct {
}

func (*UnimplementedSphereServiceServer) GetSphereByJob(ctx context.Context, req *SphereJobMessage) (*SphereList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSphereByJob not implemented")
}

func RegisterSphereServiceServer(s *grpc.Server, srv SphereServiceServer) {
	s.RegisterService(&_SphereService_serviceDesc, srv)
}

func _SphereService_GetSphereByJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SphereJobMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphereServiceServer).GetSphereByJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SphereService/GetSphereByJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphereServiceServer).GetSphereByJob(ctx, req.(*SphereJobMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _SphereService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SphereService",
	HandlerType: (*SphereServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSphereByJob",
			Handler:    _SphereService_GetSphereByJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sphere.proto",
}