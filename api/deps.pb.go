// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deps.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Dependency struct {
	Organization         *string  `protobuf:"bytes,1,opt,name=organization" json:"organization,omitempty"`
	Module               *string  `protobuf:"bytes,2,opt,name=module" json:"module,omitempty"`
	VersionConstraint    *string  `protobuf:"bytes,3,opt,name=versionConstraint" json:"versionConstraint,omitempty"`
	Scopes               []string `protobuf:"bytes,4,rep,name=scopes" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}
func (*Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_deps_9ed397ad379b8a34, []int{0}
}
func (m *Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dependency.Unmarshal(m, b)
}
func (m *Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dependency.Marshal(b, m, deterministic)
}
func (dst *Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dependency.Merge(dst, src)
}
func (m *Dependency) XXX_Size() int {
	return xxx_messageInfo_Dependency.Size(m)
}
func (m *Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Dependency proto.InternalMessageInfo

func (m *Dependency) GetOrganization() string {
	if m != nil && m.Organization != nil {
		return *m.Organization
	}
	return ""
}

func (m *Dependency) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *Dependency) GetVersionConstraint() string {
	if m != nil && m.VersionConstraint != nil {
		return *m.VersionConstraint
	}
	return ""
}

func (m *Dependency) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type DependencyManagementFile struct {
	Language             *string       `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	System               *string       `protobuf:"bytes,2,opt,name=system" json:"system,omitempty"`
	Organization         *string       `protobuf:"bytes,5,opt,name=organization" json:"organization,omitempty"`
	Module               *string       `protobuf:"bytes,6,opt,name=module" json:"module,omitempty"`
	Version              *string       `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	Dependencies         []*Dependency `protobuf:"bytes,8,rep,name=dependencies" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DependencyManagementFile) Reset()         { *m = DependencyManagementFile{} }
func (m *DependencyManagementFile) String() string { return proto.CompactTextString(m) }
func (*DependencyManagementFile) ProtoMessage()    {}
func (*DependencyManagementFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_deps_9ed397ad379b8a34, []int{1}
}
func (m *DependencyManagementFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DependencyManagementFile.Unmarshal(m, b)
}
func (m *DependencyManagementFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DependencyManagementFile.Marshal(b, m, deterministic)
}
func (dst *DependencyManagementFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DependencyManagementFile.Merge(dst, src)
}
func (m *DependencyManagementFile) XXX_Size() int {
	return xxx_messageInfo_DependencyManagementFile.Size(m)
}
func (m *DependencyManagementFile) XXX_DiscardUnknown() {
	xxx_messageInfo_DependencyManagementFile.DiscardUnknown(m)
}

var xxx_messageInfo_DependencyManagementFile proto.InternalMessageInfo

func (m *DependencyManagementFile) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *DependencyManagementFile) GetSystem() string {
	if m != nil && m.System != nil {
		return *m.System
	}
	return ""
}

func (m *DependencyManagementFile) GetOrganization() string {
	if m != nil && m.Organization != nil {
		return *m.Organization
	}
	return ""
}

func (m *DependencyManagementFile) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *DependencyManagementFile) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *DependencyManagementFile) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func init() {
	proto.RegisterType((*Dependency)(nil), "deps-cloud.des.api.Dependency")
	proto.RegisterType((*DependencyManagementFile)(nil), "deps-cloud.des.api.DependencyManagementFile")
}

func init() { proto.RegisterFile("deps.proto", fileDescriptor_deps_9ed397ad379b8a34) }

var fileDescriptor_deps_9ed397ad379b8a34 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0xb1, 0x7f, 0x1c, 0x8b, 0xe0, 0x1e, 0x64, 0xe9, 0x29, 0xe4, 0xd4, 0x83, 0xec,
	0xc1, 0x07, 0xf0, 0xa0, 0xe2, 0xcd, 0x4b, 0x8f, 0xde, 0x86, 0xee, 0x10, 0x46, 0x92, 0xd9, 0x25,
	0xb3, 0x15, 0xda, 0x87, 0xf0, 0x29, 0x7d, 0x10, 0x69, 0x9a, 0x58, 0x42, 0xa1, 0xc7, 0xdf, 0x7e,
	0xc3, 0xc7, 0xef, 0x5b, 0x00, 0x4f, 0x51, 0x5d, 0x6c, 0x43, 0x0a, 0xe6, 0xae, 0xf9, 0x8a, 0x9c,
	0xf6, 0xce, 0x93, 0x3a, 0x8c, 0x5c, 0xfe, 0x64, 0x00, 0x6f, 0x14, 0x49, 0x3c, 0xc9, 0x66, 0x67,
	0x4a, 0x58, 0x84, 0xb6, 0x42, 0xe1, 0x3d, 0x26, 0x0e, 0x62, 0xb3, 0x22, 0x5b, 0xdd, 0xac, 0x47,
	0x6f, 0xe6, 0x01, 0xa6, 0x4d, 0xf0, 0xdb, 0x9a, 0xec, 0x55, 0x97, 0xf6, 0x64, 0x1e, 0xe1, 0xfe,
	0x9b, 0x5a, 0xe5, 0x20, 0xaf, 0x41, 0x34, 0xb5, 0xc8, 0x92, 0x6c, 0xde, 0x9d, 0x9c, 0x07, 0x87,
	0x16, 0xdd, 0x84, 0x48, 0x6a, 0xaf, 0x8b, 0xfc, 0xd0, 0x72, 0xa4, 0xf2, 0x37, 0x03, 0x7b, 0x12,
	0xfa, 0x40, 0xc1, 0x8a, 0x1a, 0x92, 0xf4, 0xce, 0x35, 0x99, 0x25, 0xcc, 0x6b, 0x94, 0x6a, 0x8b,
	0x15, 0xf5, 0x6a, 0xff, 0xdc, 0x15, 0xee, 0x34, 0x51, 0x33, 0x68, 0x1d, 0xe9, 0x6c, 0xd2, 0xe4,
	0xe2, 0xa4, 0xe9, 0x68, 0x92, 0x85, 0x59, 0x6f, 0x6e, 0x67, 0x5d, 0x30, 0xa0, 0x79, 0x86, 0x85,
	0x1f, 0x2c, 0x99, 0xd4, 0xce, 0x8b, 0x7c, 0x75, 0xfb, 0xb4, 0x74, 0xe3, 0xef, 0x75, 0xa7, 0x25,
	0xeb, 0xd1, 0xfd, 0xcb, 0xe4, 0x33, 0xc7, 0xc8, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x67,
	0xa6, 0xbb, 0x9b, 0x01, 0x00, 0x00,
}
