// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/launch_plan.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LaunchPlanCreateRequest struct {
	Id                   *Identifier     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Spec                 *LaunchPlanSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LaunchPlanCreateRequest) Reset()         { *m = LaunchPlanCreateRequest{} }
func (m *LaunchPlanCreateRequest) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanCreateRequest) ProtoMessage()    {}
func (*LaunchPlanCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_launch_plan_24a639c47e09cc65, []int{0}
}
func (m *LaunchPlanCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanCreateRequest.Unmarshal(m, b)
}
func (m *LaunchPlanCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanCreateRequest.Marshal(b, m, deterministic)
}
func (dst *LaunchPlanCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanCreateRequest.Merge(dst, src)
}
func (m *LaunchPlanCreateRequest) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanCreateRequest.Size(m)
}
func (m *LaunchPlanCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanCreateRequest proto.InternalMessageInfo

func (m *LaunchPlanCreateRequest) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LaunchPlanCreateRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LaunchPlanCreateRequest) GetSpec() *LaunchPlanSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type LaunchPlan struct {
	Id                   *Identifier     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Urn                  string          `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
	Spec                 *LaunchPlanSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LaunchPlan) Reset()         { *m = LaunchPlan{} }
func (m *LaunchPlan) String() string { return proto.CompactTextString(m) }
func (*LaunchPlan) ProtoMessage()    {}
func (*LaunchPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_launch_plan_24a639c47e09cc65, []int{1}
}
func (m *LaunchPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlan.Unmarshal(m, b)
}
func (m *LaunchPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlan.Marshal(b, m, deterministic)
}
func (dst *LaunchPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlan.Merge(dst, src)
}
func (m *LaunchPlan) XXX_Size() int {
	return xxx_messageInfo_LaunchPlan.Size(m)
}
func (m *LaunchPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlan.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlan proto.InternalMessageInfo

func (m *LaunchPlan) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LaunchPlan) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LaunchPlan) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *LaunchPlan) GetSpec() *LaunchPlanSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type LaunchPlanList struct {
	LaunchPlans          []*LaunchPlan `protobuf:"bytes,1,rep,name=launch_plans,json=launchPlans,proto3" json:"launch_plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LaunchPlanList) Reset()         { *m = LaunchPlanList{} }
func (m *LaunchPlanList) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanList) ProtoMessage()    {}
func (*LaunchPlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_launch_plan_24a639c47e09cc65, []int{2}
}
func (m *LaunchPlanList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanList.Unmarshal(m, b)
}
func (m *LaunchPlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanList.Marshal(b, m, deterministic)
}
func (dst *LaunchPlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanList.Merge(dst, src)
}
func (m *LaunchPlanList) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanList.Size(m)
}
func (m *LaunchPlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanList.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanList proto.InternalMessageInfo

func (m *LaunchPlanList) GetLaunchPlans() []*LaunchPlan {
	if m != nil {
		return m.LaunchPlans
	}
	return nil
}

type LaunchPlanSpec struct {
	// Reference to the Workflow template that the launch plan references
	WorkflowUrn string `protobuf:"bytes,1,opt,name=workflow_urn,json=workflowUrn,proto3" json:"workflow_urn,omitempty"`
	// Metadata for the Launch Plan
	EntityMetadata *LaunchPlanMetadata `protobuf:"bytes,2,opt,name=entity_metadata,json=entityMetadata,proto3" json:"entity_metadata,omitempty"`
	// Input values to be passed for the execution
	DefaultInputs []*Parameter `protobuf:"bytes,3,rep,name=default_inputs,json=defaultInputs,proto3" json:"default_inputs,omitempty"`
	// Fixed, non overridable inputs for the Launch Plan
	FixedInputs          *core.NamedValueCollection `protobuf:"bytes,4,opt,name=fixed_inputs,json=fixedInputs,proto3" json:"fixed_inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LaunchPlanSpec) Reset()         { *m = LaunchPlanSpec{} }
func (m *LaunchPlanSpec) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanSpec) ProtoMessage()    {}
func (*LaunchPlanSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_launch_plan_24a639c47e09cc65, []int{3}
}
func (m *LaunchPlanSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanSpec.Unmarshal(m, b)
}
func (m *LaunchPlanSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanSpec.Marshal(b, m, deterministic)
}
func (dst *LaunchPlanSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanSpec.Merge(dst, src)
}
func (m *LaunchPlanSpec) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanSpec.Size(m)
}
func (m *LaunchPlanSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanSpec.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanSpec proto.InternalMessageInfo

func (m *LaunchPlanSpec) GetWorkflowUrn() string {
	if m != nil {
		return m.WorkflowUrn
	}
	return ""
}

func (m *LaunchPlanSpec) GetEntityMetadata() *LaunchPlanMetadata {
	if m != nil {
		return m.EntityMetadata
	}
	return nil
}

func (m *LaunchPlanSpec) GetDefaultInputs() []*Parameter {
	if m != nil {
		return m.DefaultInputs
	}
	return nil
}

func (m *LaunchPlanSpec) GetFixedInputs() *core.NamedValueCollection {
	if m != nil {
		return m.FixedInputs
	}
	return nil
}

type LaunchPlanMetadata struct {
	// Schedule to execute the Launch Plan
	Schedule *Schedule `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// List of notifications based on Execution status transitions
	Notifications        []*Notification `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LaunchPlanMetadata) Reset()         { *m = LaunchPlanMetadata{} }
func (m *LaunchPlanMetadata) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanMetadata) ProtoMessage()    {}
func (*LaunchPlanMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_launch_plan_24a639c47e09cc65, []int{4}
}
func (m *LaunchPlanMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanMetadata.Unmarshal(m, b)
}
func (m *LaunchPlanMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanMetadata.Marshal(b, m, deterministic)
}
func (dst *LaunchPlanMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanMetadata.Merge(dst, src)
}
func (m *LaunchPlanMetadata) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanMetadata.Size(m)
}
func (m *LaunchPlanMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanMetadata proto.InternalMessageInfo

func (m *LaunchPlanMetadata) GetSchedule() *Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (m *LaunchPlanMetadata) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func init() {
	proto.RegisterType((*LaunchPlanCreateRequest)(nil), "flyteidl.admin.LaunchPlanCreateRequest")
	proto.RegisterType((*LaunchPlan)(nil), "flyteidl.admin.LaunchPlan")
	proto.RegisterType((*LaunchPlanList)(nil), "flyteidl.admin.LaunchPlanList")
	proto.RegisterType((*LaunchPlanSpec)(nil), "flyteidl.admin.LaunchPlanSpec")
	proto.RegisterType((*LaunchPlanMetadata)(nil), "flyteidl.admin.LaunchPlanMetadata")
}

func init() {
	proto.RegisterFile("flyteidl/admin/launch_plan.proto", fileDescriptor_launch_plan_24a639c47e09cc65)
}

var fileDescriptor_launch_plan_24a639c47e09cc65 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0xdd, 0x0a, 0xa8, 0xb3, 0x5d, 0x90, 0x2f, 0x84, 0xf2, 0xa1, 0x25, 0x5c, 0x2a,
	0x24, 0x12, 0xb1, 0xe5, 0x8a, 0x84, 0x5a, 0x09, 0xa9, 0xa2, 0x94, 0xca, 0x15, 0x1c, 0xb8, 0xac,
	0xbc, 0xf6, 0x64, 0xd7, 0xc2, 0xb1, 0x83, 0xed, 0x50, 0xf6, 0x2f, 0x70, 0xe0, 0xc2, 0x85, 0x9f,
	0x8b, 0xe2, 0x7c, 0xed, 0x06, 0x38, 0x20, 0x71, 0xb3, 0x3d, 0xef, 0x3b, 0xf3, 0x4c, 0x32, 0x83,
	0x66, 0x99, 0xdc, 0x38, 0x10, 0x5c, 0xa6, 0x94, 0xe7, 0x42, 0xa5, 0x92, 0x96, 0x8a, 0xad, 0x17,
	0x85, 0xa4, 0x2a, 0x29, 0x8c, 0x76, 0x1a, 0x4f, 0x5b, 0x45, 0xe2, 0x15, 0x87, 0x0f, 0x3b, 0x07,
	0xd3, 0x06, 0x52, 0xa1, 0x1c, 0x98, 0x8c, 0x32, 0xa8, 0xe5, 0x5b, 0xe1, 0x3a, 0xa1, 0x65, 0x6b,
	0xe0, 0xa5, 0x6c, 0xc3, 0xf7, 0x07, 0x61, 0xa6, 0xf3, 0x5c, 0x37, 0xa5, 0xe2, 0x1f, 0x01, 0xba,
	0x7b, 0xee, 0x01, 0x2e, 0x25, 0x55, 0xa7, 0x06, 0xa8, 0x03, 0x02, 0x9f, 0x4b, 0xb0, 0x0e, 0x3f,
	0x45, 0x23, 0xc1, 0xa3, 0x60, 0x16, 0x1c, 0x85, 0xf3, 0xc3, 0x64, 0x97, 0x29, 0x39, 0xe3, 0xa0,
	0x9c, 0xc8, 0x04, 0x18, 0x32, 0x12, 0x1c, 0x47, 0xe8, 0xe6, 0x17, 0x30, 0x56, 0x68, 0x15, 0x8d,
	0x66, 0xc1, 0xd1, 0x3e, 0x69, 0xaf, 0x78, 0x8e, 0xf6, 0x6c, 0x01, 0x2c, 0x1a, 0xfb, 0x3c, 0x8f,
	0x86, 0x79, 0xfa, 0xe2, 0x57, 0x05, 0x30, 0xe2, 0xb5, 0xf1, 0xcf, 0x00, 0xa1, 0x3e, 0xf0, 0x9f,
	0x40, 0xee, 0xa0, 0x71, 0x69, 0x94, 0xe7, 0xd8, 0x27, 0xd5, 0xb1, 0x43, 0xdb, 0xfb, 0x07, 0xb4,
	0x77, 0x68, 0xda, 0xbf, 0x9f, 0x0b, 0xeb, 0xf0, 0x4b, 0x34, 0xd9, 0xfa, 0x85, 0x36, 0x0a, 0x66,
	0xe3, 0x3f, 0x71, 0xf6, 0x2e, 0x12, 0xca, 0xee, 0x6c, 0xe3, 0x6f, 0xa3, 0xed, 0x8c, 0x55, 0x25,
	0xfc, 0x18, 0x4d, 0xae, 0xb5, 0xf9, 0x94, 0x49, 0x7d, 0xbd, 0xa8, 0x90, 0x03, 0x8f, 0x1c, 0xb6,
	0x6f, 0xef, 0x8d, 0xc2, 0x6f, 0xd0, 0xed, 0xaa, 0x6d, 0xb7, 0x59, 0xe4, 0xe0, 0x28, 0xa7, 0x8e,
	0xfa, 0x76, 0xc3, 0x79, 0xfc, 0xf7, 0xba, 0x6f, 0x1b, 0x25, 0x99, 0xd6, 0xd6, 0xf6, 0x8e, 0x5f,
	0xa1, 0x29, 0x87, 0x8c, 0x96, 0xd2, 0x2d, 0x84, 0x2a, 0x4a, 0x67, 0xa3, 0xb1, 0xef, 0xe1, 0xde,
	0x30, 0xd7, 0x25, 0x35, 0x34, 0x07, 0x07, 0x86, 0x1c, 0x34, 0x86, 0x33, 0xaf, 0xc7, 0xaf, 0xd1,
	0x24, 0x13, 0x5f, 0x81, 0xb7, 0xfe, 0xfa, 0x8b, 0x3e, 0xe9, 0xfd, 0xd5, 0xe0, 0x26, 0x17, 0x34,
	0x07, 0xfe, 0x81, 0xca, 0x12, 0x4e, 0xb5, 0x94, 0xc0, 0x9c, 0xd0, 0x8a, 0x84, 0xde, 0x58, 0xe7,
	0x89, 0xbf, 0x07, 0x08, 0xff, 0x0e, 0x8c, 0x5f, 0xa0, 0x5b, 0xed, 0x50, 0x37, 0x63, 0x10, 0x0d,
	0xd1, 0xae, 0x9a, 0x38, 0xe9, 0x94, 0xf8, 0x04, 0x1d, 0x28, 0x5d, 0x8d, 0x06, 0xa3, 0x55, 0x25,
	0x1b, 0x8d, 0x7c, 0x57, 0x0f, 0x86, 0xd6, 0x8b, 0x2d, 0x11, 0xd9, 0xb5, 0x9c, 0x1c, 0x7f, 0x7c,
	0xbe, 0x12, 0x6e, 0x5d, 0x2e, 0x13, 0xa6, 0xf3, 0x54, 0x6e, 0x32, 0x97, 0x76, 0xeb, 0xb4, 0x02,
	0x95, 0x16, 0xcb, 0x67, 0x2b, 0x9d, 0xee, 0x6e, 0xd8, 0xf2, 0x86, 0xdf, 0xad, 0xe3, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa6, 0x79, 0x43, 0x29, 0xea, 0x03, 0x00, 0x00,
}
