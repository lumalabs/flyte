// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/plugins/kubeflow/common.proto (package flyteidl.plugins.kubeflow, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum flyteidl.plugins.kubeflow.CleanPodPolicy
 */
export enum CleanPodPolicy {
  /**
   * @generated from enum value: CLEANPOD_POLICY_NONE = 0;
   */
  CLEANPOD_POLICY_NONE = 0,

  /**
   * @generated from enum value: CLEANPOD_POLICY_RUNNING = 1;
   */
  CLEANPOD_POLICY_RUNNING = 1,

  /**
   * @generated from enum value: CLEANPOD_POLICY_ALL = 2;
   */
  CLEANPOD_POLICY_ALL = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(CleanPodPolicy)
proto3.util.setEnumType(CleanPodPolicy, "flyteidl.plugins.kubeflow.CleanPodPolicy", [
  { no: 0, name: "CLEANPOD_POLICY_NONE" },
  { no: 1, name: "CLEANPOD_POLICY_RUNNING" },
  { no: 2, name: "CLEANPOD_POLICY_ALL" },
]);

/**
 * @generated from message flyteidl.plugins.kubeflow.SchedulingPolicy
 */
export class SchedulingPolicy extends Message<SchedulingPolicy> {
  /**
   * @generated from field: string queue = 1;
   */
  queue = "";

  /**
   * @generated from field: string priority_class = 2;
   */
  priorityClass = "";

  /**
   * @generated from field: int32 min_available = 3;
   */
  minAvailable = 0;

  constructor(data?: PartialMessage<SchedulingPolicy>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.plugins.kubeflow.SchedulingPolicy";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "queue", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "priority_class", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "min_available", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SchedulingPolicy {
    return new SchedulingPolicy().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SchedulingPolicy {
    return new SchedulingPolicy().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SchedulingPolicy {
    return new SchedulingPolicy().fromJsonString(jsonString, options);
  }

  static equals(a: SchedulingPolicy | PlainMessage<SchedulingPolicy> | undefined, b: SchedulingPolicy | PlainMessage<SchedulingPolicy> | undefined): boolean {
    return proto3.util.equals(SchedulingPolicy, a, b);
  }
}

/**
 * @generated from message flyteidl.plugins.kubeflow.RunPolicy
 */
export class RunPolicy extends Message<RunPolicy> {
  /**
   * Defines the policy to kill pods after the job completes. Default to None.
   *
   * @generated from field: flyteidl.plugins.kubeflow.CleanPodPolicy clean_pod_policy = 1;
   */
  cleanPodPolicy = CleanPodPolicy.CLEANPOD_POLICY_NONE;

  /**
   * TTL to clean up jobs. Default to infinite.
   *
   * @generated from field: int32 ttl_seconds_after_finished = 2;
   */
  ttlSecondsAfterFinished = 0;

  /**
   * Specifies the duration in seconds relative to the startTime that the job may be active
   * before the system tries to terminate it; value must be positive integer.
   *
   * @generated from field: int32 active_deadline_seconds = 3;
   */
  activeDeadlineSeconds = 0;

  /**
   * Number of retries before marking this job failed.
   *
   * @generated from field: int32 backoff_limit = 4;
   */
  backoffLimit = 0;

  /**
   * Scheduling policy to control priorities and queues
   *
   * @generated from field: flyteidl.plugins.kubeflow.SchedulingPolicy scheduling_policy = 5;
   */
  schedulingPolicy?: SchedulingPolicy;

  /**
   * Suspend job execution
   *
   * @generated from field: bool suspend = 6;
   */
  suspend = false;

  constructor(data?: PartialMessage<RunPolicy>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.plugins.kubeflow.RunPolicy";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "clean_pod_policy", kind: "enum", T: proto3.getEnumType(CleanPodPolicy) },
    { no: 2, name: "ttl_seconds_after_finished", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "active_deadline_seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "backoff_limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "scheduling_policy", kind: "message", T: SchedulingPolicy },
    { no: 6, name: "suspend", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RunPolicy {
    return new RunPolicy().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RunPolicy {
    return new RunPolicy().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RunPolicy {
    return new RunPolicy().fromJsonString(jsonString, options);
  }

  static equals(a: RunPolicy | PlainMessage<RunPolicy> | undefined, b: RunPolicy | PlainMessage<RunPolicy> | undefined): boolean {
    return proto3.util.equals(RunPolicy, a, b);
  }
}

