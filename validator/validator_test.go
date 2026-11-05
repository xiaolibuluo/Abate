// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package validator

import (
	"context"
	"testing"

	"github.com/qingdujian/abate/core"
)

func testExperiment() core.Experiment {
	return core.Experiment{
		ID:          "val-test",
		Seed:        42,
		Criticality: 3,
		Target:      core.Target{Kind: core.TargetService, Name: "svc"},
		Fault:       core.Fault{Class: core.F1Latency, Params: map[string]string{"delay": "100"}},
		Blast:       core.BlastRadius{MaxInstances: 1, MaxDurationMs: 60000, MaxTraffic: 1000},
		Steady:       []core.SLIAssert{{SLI: "latency", Assert: "< 500"}},
	}
}

func TestCapability(t *testing.T) {
	v := New()
	cap := v.Capability()
	if cap.Name != "file-validator" {
		t.Errorf("name = %q", cap.Name)
	}
	if cap.Protocol != core.ProtocolVersion {
		t.Errorf("protocol = %q", cap.Protocol)
	}
}

func TestPreflightOK(t *testing.T) {
	v := New()
	issues, err := v.Preflight(context.Background(), testExperiment())
	if err != nil {
		t.Fatalf("preflight: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("issues = %v", issues)
	}
}

func TestPreflightPartialBlast(t *testing.T) {
	v := New()
	e := testExperiment()
	e.Blast.MaxTraffic = 0
	issues, err := v.Preflight(context.Background(), e)
	if err != nil {
		t.Fatalf("preflight: %v", err)
	}
	if len(issues) == 0 {
		t.Error("expected partial blast issue")
	}
}

func TestPreflightValidationError(t *testing.T) {
	v := New()
	e := testExperiment()
	e.ID = ""
	_, err := v.Preflight(context.Background(), e)
	if err == nil {
		t.Error("expected validation error")
	}
}

func TestPostflightOK(t *testing.T) {
	v := New()
	issues, err := v.Postflight(context.Background(), testExperiment())
	if err != nil {
		t.Fatalf("postflight: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("issues = %v", issues)
	}
}

func TestPostflightResidue(t *testing.T) {
	v := New()
	v.SetResidue(true)
	issues, err := v.Postflight(context.Background(), testExperiment())
	if err != nil {
		t.Fatalf("postflight: %v", err)
	}
	if len(issues) == 0 {
		t.Error("expected residue issue")
	}
}

func TestPreflightFullBlast(t *testing.T) {
	v := New()
	e := testExperiment()
	e.Blast.MaxTraffic = 1000
	issues, err := v.Preflight(context.Background(), e)
	if err != nil {
		t.Fatalf("preflight: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("expected no issues, got %v", issues)
	}
}
