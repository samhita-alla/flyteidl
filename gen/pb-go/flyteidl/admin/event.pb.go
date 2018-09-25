// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/event.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import event "github.com/lyft/flyteidl/gen/pb-go/flyteidl/event"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// the id of the originator (Propeller) of the even
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Workflow execution id
	ExecutionId string `protobuf:"bytes,3,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// This represents the timestamp of when the event occured, not the request timestamp
	OccurredAt           *timestamp.Timestamp          `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	Event                *event.WorkflowExecutionEvent `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *WorkflowExecutionEventRequest) Reset()         { *m = WorkflowExecutionEventRequest{} }
func (m *WorkflowExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventRequest) ProtoMessage()    {}
func (*WorkflowExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_cdf749342e91a184, []int{0}
}
func (m *WorkflowExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventRequest.Merge(dst, src)
}
func (m *WorkflowExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Size(m)
}
func (m *WorkflowExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventRequest proto.InternalMessageInfo

func (m *WorkflowExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *WorkflowExecutionEventRequest) GetEvent() *event.WorkflowExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type WorkflowExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionEventResponse) Reset()         { *m = WorkflowExecutionEventResponse{} }
func (m *WorkflowExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventResponse) ProtoMessage()    {}
func (*WorkflowExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_cdf749342e91a184, []int{1}
}
func (m *WorkflowExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventResponse.Merge(dst, src)
}
func (m *WorkflowExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Size(m)
}
func (m *WorkflowExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventResponse proto.InternalMessageInfo

type NodeExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// the id of the originator (Propeller) of the even
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Workflow execution id
	ExecutionId string `protobuf:"bytes,3,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// This represents the timestamp of when the event occured, not the request timestamp
	OccurredAt           *timestamp.Timestamp      `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	Event                *event.NodeExecutionEvent `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *NodeExecutionEventRequest) Reset()         { *m = NodeExecutionEventRequest{} }
func (m *NodeExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventRequest) ProtoMessage()    {}
func (*NodeExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_cdf749342e91a184, []int{2}
}
func (m *NodeExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventRequest.Unmarshal(m, b)
}
func (m *NodeExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventRequest.Merge(dst, src)
}
func (m *NodeExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventRequest.Size(m)
}
func (m *NodeExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventRequest proto.InternalMessageInfo

func (m *NodeExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *NodeExecutionEventRequest) GetEvent() *event.NodeExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type NodeExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionEventResponse) Reset()         { *m = NodeExecutionEventResponse{} }
func (m *NodeExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventResponse) ProtoMessage()    {}
func (*NodeExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_cdf749342e91a184, []int{3}
}
func (m *NodeExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventResponse.Unmarshal(m, b)
}
func (m *NodeExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventResponse.Merge(dst, src)
}
func (m *NodeExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventResponse.Size(m)
}
func (m *NodeExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventResponse proto.InternalMessageInfo

type TaskExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// the id of the originator (Propeller) of the even
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Workflow execution id
	ExecutionId string `protobuf:"bytes,3,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// This represents the timestamp of when the event occured, not the request timestamp
	OccurredAt           *timestamp.Timestamp      `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	Event                *event.TaskExecutionEvent `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TaskExecutionEventRequest) Reset()         { *m = TaskExecutionEventRequest{} }
func (m *TaskExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventRequest) ProtoMessage()    {}
func (*TaskExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_cdf749342e91a184, []int{4}
}
func (m *TaskExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventRequest.Unmarshal(m, b)
}
func (m *TaskExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventRequest.Merge(dst, src)
}
func (m *TaskExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventRequest.Size(m)
}
func (m *TaskExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventRequest proto.InternalMessageInfo

func (m *TaskExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *TaskExecutionEventRequest) GetEvent() *event.TaskExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type TaskExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEventResponse) Reset()         { *m = TaskExecutionEventResponse{} }
func (m *TaskExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventResponse) ProtoMessage()    {}
func (*TaskExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_cdf749342e91a184, []int{5}
}
func (m *TaskExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventResponse.Unmarshal(m, b)
}
func (m *TaskExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventResponse.Merge(dst, src)
}
func (m *TaskExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventResponse.Size(m)
}
func (m *TaskExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WorkflowExecutionEventRequest)(nil), "flyteidl.admin.WorkflowExecutionEventRequest")
	proto.RegisterType((*WorkflowExecutionEventResponse)(nil), "flyteidl.admin.WorkflowExecutionEventResponse")
	proto.RegisterType((*NodeExecutionEventRequest)(nil), "flyteidl.admin.NodeExecutionEventRequest")
	proto.RegisterType((*NodeExecutionEventResponse)(nil), "flyteidl.admin.NodeExecutionEventResponse")
	proto.RegisterType((*TaskExecutionEventRequest)(nil), "flyteidl.admin.TaskExecutionEventRequest")
	proto.RegisterType((*TaskExecutionEventResponse)(nil), "flyteidl.admin.TaskExecutionEventResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/event.proto", fileDescriptor_event_cdf749342e91a184) }

var fileDescriptor_event_cdf749342e91a184 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x53, 0x15, 0x95, 0x87, 0x71, 0xe8, 0x04, 0x8d, 0x28, 0x76, 0x30, 0x2c, 0xde, 0x45,
	0x59, 0x4c, 0x74, 0xd1, 0x84, 0x81, 0xc5, 0x81, 0x90, 0x98, 0xb8, 0x90, 0xb6, 0xf7, 0x5a, 0x1b,
	0xda, 0xbe, 0xda, 0x5e, 0x45, 0x3e, 0xb5, 0xa3, 0xab, 0xe1, 0x8e, 0x16, 0x30, 0xe5, 0x03, 0xe8,
	0xd6, 0xde, 0xff, 0xde, 0x3f, 0xf7, 0xfb, 0x25, 0x0f, 0x2c, 0x3f, 0x5a, 0x48, 0x0c, 0x45, 0xc4,
	0x1d, 0x11, 0x87, 0x09, 0xc7, 0x0f, 0x4c, 0x24, 0x4b, 0x33, 0x92, 0x64, 0x9e, 0x96, 0x19, 0x53,
	0x99, 0x75, 0x11, 0x10, 0x05, 0x11, 0x72, 0x95, 0xba, 0x85, 0xcf, 0x65, 0x18, 0x63, 0x2e, 0x9d,
	0x38, 0xd5, 0x03, 0xd6, 0xba, 0x4c, 0xd5, 0x6c, 0x96, 0xd9, 0xdf, 0x06, 0x74, 0x5f, 0x28, 0x9b,
	0xf9, 0x11, 0xcd, 0x87, 0x9f, 0xe8, 0x15, 0x32, 0xa4, 0x64, 0xb8, 0xbc, 0x30, 0xc6, 0xf7, 0x02,
	0x73, 0x69, 0x76, 0x01, 0x32, 0xfd, 0x39, 0x0d, 0x45, 0xdb, 0xe8, 0x19, 0xfd, 0xe6, 0xb8, 0xb9,
	0x3a, 0x19, 0x09, 0xb3, 0x03, 0xc7, 0x34, 0x4f, 0x30, 0x5b, 0x86, 0x7b, 0x2a, 0x3c, 0x52, 0xff,
	0x23, 0x61, 0x5e, 0xc2, 0x09, 0x96, 0x95, 0xcb, 0x78, 0x5f, 0xc5, 0xad, 0xea, 0x6c, 0x24, 0xcc,
	0x7b, 0x68, 0x91, 0xe7, 0x15, 0x59, 0x86, 0x62, 0xea, 0xc8, 0xf6, 0x41, 0xcf, 0xe8, 0xb7, 0x6e,
	0x2d, 0xa6, 0x89, 0x58, 0x49, 0xc4, 0x26, 0x25, 0xd1, 0x18, 0xca, 0xeb, 0x8f, 0xd2, 0x7c, 0x80,
	0x86, 0x42, 0x69, 0x37, 0xd4, 0xd8, 0x15, 0xab, 0xc4, 0x68, 0xc2, 0x1d, 0x5c, 0x7a, 0xc8, 0xee,
	0xc1, 0xf9, 0x2e, 0xf0, 0x3c, 0xa5, 0x24, 0x47, 0xfb, 0xcb, 0x80, 0xce, 0x33, 0x09, 0xfc, 0x6b,
	0x5e, 0xee, 0xb6, 0xbd, 0xd8, 0xbf, 0xbd, 0xd4, 0x30, 0xad, 0x9c, 0x9c, 0x81, 0x55, 0x07, 0xbc,
	0xe1, 0x63, 0xe2, 0xe4, 0xb3, 0xff, 0xe6, 0xa3, 0x86, 0x69, 0xed, 0xa3, 0x0e, 0x58, 0xfb, 0x78,
	0x1a, 0xbc, 0xde, 0x04, 0xa1, 0x7c, 0x2b, 0x5c, 0xe6, 0x51, 0xcc, 0xa3, 0x85, 0x2f, 0x79, 0xb5,
	0x69, 0x01, 0x26, 0x3c, 0x75, 0xaf, 0x03, 0xe2, 0xdb, 0x9b, 0xec, 0x1e, 0xaa, 0xc7, 0x0e, 0x7e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xa4, 0x96, 0x79, 0xe2, 0x03, 0x00, 0x00,
}
