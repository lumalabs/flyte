// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/plugins/kubeflow/common.proto

package kubeflow

import (
	plugins "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/plugins"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of flyteidl/plugins/common.proto.

type RestartPolicy = plugins.RestartPolicy

const RestartPolicy_RESTART_POLICY_NEVER = plugins.RestartPolicy_RESTART_POLICY_NEVER
const RestartPolicy_RESTART_POLICY_ON_FAILURE = plugins.RestartPolicy_RESTART_POLICY_ON_FAILURE
const RestartPolicy_RESTART_POLICY_ALWAYS = plugins.RestartPolicy_RESTART_POLICY_ALWAYS

var RestartPolicy_name = plugins.RestartPolicy_name
var RestartPolicy_value = plugins.RestartPolicy_value

type CommonReplicaSpec = plugins.CommonReplicaSpec

type CleanPodPolicy int32

const (
	CleanPodPolicy_CLEANPOD_POLICY_NONE    CleanPodPolicy = 0
	CleanPodPolicy_CLEANPOD_POLICY_RUNNING CleanPodPolicy = 1
	CleanPodPolicy_CLEANPOD_POLICY_ALL     CleanPodPolicy = 2
)

// Enum value maps for CleanPodPolicy.
var (
	CleanPodPolicy_name = map[int32]string{
		0: "CLEANPOD_POLICY_NONE",
		1: "CLEANPOD_POLICY_RUNNING",
		2: "CLEANPOD_POLICY_ALL",
	}
	CleanPodPolicy_value = map[string]int32{
		"CLEANPOD_POLICY_NONE":    0,
		"CLEANPOD_POLICY_RUNNING": 1,
		"CLEANPOD_POLICY_ALL":     2,
	}
)

func (x CleanPodPolicy) Enum() *CleanPodPolicy {
	p := new(CleanPodPolicy)
	*p = x
	return p
}

func (x CleanPodPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CleanPodPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_plugins_kubeflow_common_proto_enumTypes[0].Descriptor()
}

func (CleanPodPolicy) Type() protoreflect.EnumType {
	return &file_flyteidl_plugins_kubeflow_common_proto_enumTypes[0]
}

func (x CleanPodPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CleanPodPolicy.Descriptor instead.
func (CleanPodPolicy) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP(), []int{0}
}

type SchedulingPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue         string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	PriorityClass string `protobuf:"bytes,2,opt,name=priority_class,json=priorityClass,proto3" json:"priority_class,omitempty"`
	MinAvailable  int32  `protobuf:"varint,3,opt,name=min_available,json=minAvailable,proto3" json:"min_available,omitempty"`
}

func (x *SchedulingPolicy) Reset() {
	*x = SchedulingPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_kubeflow_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulingPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulingPolicy) ProtoMessage() {}

func (x *SchedulingPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_kubeflow_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulingPolicy.ProtoReflect.Descriptor instead.
func (*SchedulingPolicy) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulingPolicy) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *SchedulingPolicy) GetPriorityClass() string {
	if x != nil {
		return x.PriorityClass
	}
	return ""
}

func (x *SchedulingPolicy) GetMinAvailable() int32 {
	if x != nil {
		return x.MinAvailable
	}
	return 0
}

type RunPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the policy to kill pods after the job completes. Default to None.
	CleanPodPolicy CleanPodPolicy `protobuf:"varint,1,opt,name=clean_pod_policy,json=cleanPodPolicy,proto3,enum=flyteidl.plugins.kubeflow.CleanPodPolicy" json:"clean_pod_policy,omitempty"`
	// TTL to clean up jobs. Default to infinite.
	TtlSecondsAfterFinished int32 `protobuf:"varint,2,opt,name=ttl_seconds_after_finished,json=ttlSecondsAfterFinished,proto3" json:"ttl_seconds_after_finished,omitempty"`
	// Specifies the duration in seconds relative to the startTime that the job may be active
	// before the system tries to terminate it; value must be positive integer.
	ActiveDeadlineSeconds int32 `protobuf:"varint,3,opt,name=active_deadline_seconds,json=activeDeadlineSeconds,proto3" json:"active_deadline_seconds,omitempty"`
	// Number of retries before marking this job failed.
	BackoffLimit int32 `protobuf:"varint,4,opt,name=backoff_limit,json=backoffLimit,proto3" json:"backoff_limit,omitempty"`
	// Scheduling policy to control priorities and queues
	SchedulingPolicy *SchedulingPolicy `protobuf:"bytes,5,opt,name=scheduling_policy,json=schedulingPolicy,proto3" json:"scheduling_policy,omitempty"`
	// Suspend job execution
	Suspend bool `protobuf:"varint,6,opt,name=suspend,proto3" json:"suspend,omitempty"`
}

func (x *RunPolicy) Reset() {
	*x = RunPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_kubeflow_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunPolicy) ProtoMessage() {}

func (x *RunPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_kubeflow_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunPolicy.ProtoReflect.Descriptor instead.
func (*RunPolicy) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP(), []int{1}
}

func (x *RunPolicy) GetCleanPodPolicy() CleanPodPolicy {
	if x != nil {
		return x.CleanPodPolicy
	}
	return CleanPodPolicy_CLEANPOD_POLICY_NONE
}

func (x *RunPolicy) GetTtlSecondsAfterFinished() int32 {
	if x != nil {
		return x.TtlSecondsAfterFinished
	}
	return 0
}

func (x *RunPolicy) GetActiveDeadlineSeconds() int32 {
	if x != nil {
		return x.ActiveDeadlineSeconds
	}
	return 0
}

func (x *RunPolicy) GetBackoffLimit() int32 {
	if x != nil {
		return x.BackoffLimit
	}
	return 0
}

func (x *RunPolicy) GetSchedulingPolicy() *SchedulingPolicy {
	if x != nil {
		return x.SchedulingPolicy
	}
	return nil
}

func (x *RunPolicy) GetSuspend() bool {
	if x != nil {
		return x.Suspend
	}
	return false
}

var File_flyteidl_plugins_kubeflow_common_proto protoreflect.FileDescriptor

var file_flyteidl_plugins_kubeflow_common_proto_rawDesc = []byte{
	0x0a, 0x26, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66,
	0x6c, 0x6f, 0x77, 0x1a, 0x1d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xee, 0x02, 0x0a, 0x09, 0x52, 0x75, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x53, 0x0a, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f,
	0x70, 0x6f, 0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x29, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x65,
	0x61, 0x6e, 0x50, 0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x63, 0x6c, 0x65,
	0x61, 0x6e, 0x50, 0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x1a, 0x74,
	0x74, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x17, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x58, 0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x10, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x2a, 0x60, 0x0a, 0x0e, 0x43, 0x6c, 0x65,
	0x61, 0x6e, 0x50, 0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x4c, 0x45, 0x41, 0x4e, 0x50, 0x4f, 0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x50, 0x4f,
	0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x50, 0x4f, 0x44, 0x5f, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0xfa, 0x01, 0x0a, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72,
	0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x66, 0x6c, 0x6f, 0x77, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x4b, 0xaa, 0x02, 0x19, 0x46, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x66, 0x6c,
	0x6f, 0x77, 0xe2, 0x02, 0x25, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x46, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x3a, 0x3a,
	0x4b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_flyteidl_plugins_kubeflow_common_proto_rawDescOnce sync.Once
	file_flyteidl_plugins_kubeflow_common_proto_rawDescData = file_flyteidl_plugins_kubeflow_common_proto_rawDesc
)

func file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP() []byte {
	file_flyteidl_plugins_kubeflow_common_proto_rawDescOnce.Do(func() {
		file_flyteidl_plugins_kubeflow_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_plugins_kubeflow_common_proto_rawDescData)
	})
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescData
}

var file_flyteidl_plugins_kubeflow_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyteidl_plugins_kubeflow_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flyteidl_plugins_kubeflow_common_proto_goTypes = []interface{}{
	(CleanPodPolicy)(0),      // 0: flyteidl.plugins.kubeflow.CleanPodPolicy
	(*SchedulingPolicy)(nil), // 1: flyteidl.plugins.kubeflow.SchedulingPolicy
	(*RunPolicy)(nil),        // 2: flyteidl.plugins.kubeflow.RunPolicy
}
var file_flyteidl_plugins_kubeflow_common_proto_depIdxs = []int32{
	0, // 0: flyteidl.plugins.kubeflow.RunPolicy.clean_pod_policy:type_name -> flyteidl.plugins.kubeflow.CleanPodPolicy
	1, // 1: flyteidl.plugins.kubeflow.RunPolicy.scheduling_policy:type_name -> flyteidl.plugins.kubeflow.SchedulingPolicy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flyteidl_plugins_kubeflow_common_proto_init() }
func file_flyteidl_plugins_kubeflow_common_proto_init() {
	if File_flyteidl_plugins_kubeflow_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_plugins_kubeflow_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulingPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_plugins_kubeflow_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_plugins_kubeflow_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_plugins_kubeflow_common_proto_goTypes,
		DependencyIndexes: file_flyteidl_plugins_kubeflow_common_proto_depIdxs,
		EnumInfos:         file_flyteidl_plugins_kubeflow_common_proto_enumTypes,
		MessageInfos:      file_flyteidl_plugins_kubeflow_common_proto_msgTypes,
	}.Build()
	File_flyteidl_plugins_kubeflow_common_proto = out.File
	file_flyteidl_plugins_kubeflow_common_proto_rawDesc = nil
	file_flyteidl_plugins_kubeflow_common_proto_goTypes = nil
	file_flyteidl_plugins_kubeflow_common_proto_depIdxs = nil
}
