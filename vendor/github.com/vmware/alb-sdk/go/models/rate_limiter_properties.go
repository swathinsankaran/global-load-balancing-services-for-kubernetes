// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// RateLimiterProperties rate limiter properties
// swagger:model RateLimiterProperties
type RateLimiterProperties struct {

	// Number of stages in msf rate limiter. Allowed values are 1-2. Field introduced in 20.1.1.
	MsfNumStages *int32 `json:"msf_num_stages,omitempty"`

	// Each stage size in msf rate limiter. Field introduced in 20.1.1.
	MsfStageSize *int64 `json:"msf_stage_size,omitempty"`
}