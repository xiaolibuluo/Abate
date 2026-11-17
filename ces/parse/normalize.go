/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"abate/ces/schema"
)

// Normalize fills defaults and verifies cross-field references.
//
// The returned copy is safe to execute: defaults are applied, probe
// references resolve, the schedule has exactly one mode, and a missing seed
// is derived deterministically from the document bytes.
func Normalize(exp *schema.Experiment) (*schema.Experiment, error) {
	if exp == nil {
		return nil, newError(RuleEmptyDocument, "$", "nil experiment")
	}
	out := *exp
	if out.Metadata.Namespace == "" {
		out.Metadata.Namespace = "default"
	}
	if out.Spec.Rollback.Strategy == "" {
		out.Spec.Rollback.Strategy = schema.StrategyAuto
	}
	probeIDs := make(map[string]bool, len(out.Spec.Probes))
	for i := range out.Spec.Probes {
		p := out.Spec.Probes[i]
		if p.ID == "" || probeIDs[p.ID] {
			return nil, newError(RuleMalformedDocument,
				fmt.Sprintf("$.spec.probes[%d].id", i), "empty or duplicate probe ID")
		}
		probeIDs[p.ID] = true
		if p.Interval <= 0 {
			return nil, newError(RuleMalformedDocument,
				fmt.Sprintf("$.spec.probes[%d].interval", i), "interval must be positive")
		}
	}
	for i, ss := range out.Spec.SteadyState {
		if !probeIDs[ss.ProbeRef] {
			return nil, newError(RuleDanglingProbeRef,
				fmt.Sprintf("$.spec.steadyState[%d].probeRef", i),
				"steady state references an undeclared probe")
		}
	}
	if count := scheduleModes(out.Spec.Schedule); count != 1 {
		return nil, newError(RuleScheduleExclusive, "$.spec.schedule",
			"exactly one of once, cron or continuous must be set")
	}
	if out.Spec.Seed == 0 {
		seed, err := deriveSeed(exp)
		if err != nil {
			return nil, newError(RuleMalformedDocument, "$.spec.seed", err.Error())
		}
		out.Spec.Seed = seed
	}
	return &out, nil
}

// scheduleModes counts the non-nil scheduling modes.
func scheduleModes(s schema.Schedule) int {
	count := 0
	if s.Once != nil {
		count++
	}
	if s.Cron != nil {
		count++
	}
	if s.Continuous != nil {
		count++
	}
	return count
}

// deriveSeed computes a stable seed from the canonical document encoding.
func deriveSeed(exp *schema.Experiment) (uint64, error) {
	data, err := json.Marshal(exp)
	if err != nil {
		return 0, err
	}
	sum := sha256.Sum256(data)
	return binary.BigEndian.Uint64(sum[:8]), nil
}
