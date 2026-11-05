// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package core

import "testing"

// TestBlastRadiusValidate 验证三选二声明规则（CES-2.5.1）。
func TestBlastRadiusValidate(t *testing.T) {
	cases := []struct {
		name string
		b    BlastRadius
		ok   bool
	}{
		{"three", BlastRadius{MaxInstances: 1, MaxDurationMs: 1000, MaxTraffic: 500}, true},
		{"two-instance-duration", BlastRadius{MaxInstances: 1, MaxDurationMs: 1000}, true},
		{"two-instance-traffic", BlastRadius{MaxInstances: 1, MaxTraffic: 500}, true},
		{"two-duration-traffic", BlastRadius{MaxDurationMs: 1000, MaxTraffic: 500}, true},
		{"one", BlastRadius{MaxInstances: 1}, false},
		{"none", BlastRadius{}, false},
	}
	for _, c := range cases {
		err := c.b.Validate()
		if c.ok && err != nil {
			t.Errorf("%s: rejected: %v", c.name, err)
		}
		if !c.ok && err == nil {
			t.Errorf("%s: must be rejected", c.name)
		}
	}
}

// TestSLIAssertValidate 验证断言字段完整性（CES-2.6.1）。
func TestSLIAssertValidate(t *testing.T) {
	if err := (SLIAssert{SLI: "s", Assert: ">= 99%"}).Validate(); err != nil {
		t.Errorf("valid assert rejected: %v", err)
	}
	if err := (SLIAssert{Assert: ">= 99%"}).Validate(); err == nil {
		t.Error("empty sli must be rejected")
	}
	if err := (SLIAssert{SLI: "s"}).Validate(); err == nil {
		t.Error("empty assert must be rejected")
	}
}

// TestExperimentValidate 验证五元组整体验校与关键度边界（CES 第 2 章、CES-4.3.1）。
func TestExperimentValidate(t *testing.T) {
	base := goldenExperiment("exp-t")
	if err := base.Validate(); err != nil {
		t.Fatalf("golden experiment rejected: %v", err)
	}

	noID := base
	noID.ID = ""
	if err := noID.Validate(); err == nil {
		t.Error("empty id must be rejected")
	}

	badTarget := base
	badTarget.Target = Target{Kind: "x", Name: "n"}
	if err := badTarget.Validate(); err == nil {
		t.Error("invalid target must be rejected")
	}

	badFault := base
	badFault.Fault = Fault{Class: "F9", Params: map[string]string{"a": "b"}}
	if err := badFault.Validate(); err == nil {
		t.Error("invalid fault must be rejected")
	}

	badBlast := base
	badBlast.Blast = BlastRadius{MaxInstances: 1}
	if err := badBlast.Validate(); err == nil {
		t.Error("insufficient blastRadius must be rejected")
	}

	noSteady := base
	noSteady.Steady = nil
	if err := noSteady.Validate(); err == nil {
		t.Error("empty steadyState must be rejected (CES-2.6.1)")
	}

	badSteady := base
	badSteady.Steady = []SLIAssert{{SLI: "", Assert: "x"}}
	if err := badSteady.Validate(); err == nil {
		t.Error("invalid steady entry must be rejected")
	}

	badAbort := base
	badAbort.AbortOn = []SLIAssert{{SLI: "s", Assert: ""}}
	if err := badAbort.Validate(); err == nil {
		t.Error("invalid abortOn entry must be rejected")
	}

	for _, c := range []int{0, 6, -1} {
		bad := base
		bad.Criticality = c
		if err := bad.Validate(); err == nil {
			t.Errorf("criticality %d must be rejected", c)
		}
	}
}

// TestRecoveryBudget 验证恢复预算缺省值 300000ms（CES-4.4.2）。
func TestRecoveryBudget(t *testing.T) {
	e := goldenExperiment("exp-b")
	if got := e.RecoveryBudget(); got != 300000 {
		t.Errorf("default budget=%d, want 300000", got)
	}
	e.RecoveryBudgetMs = 60000
	if got := e.RecoveryBudget(); got != 60000 {
		t.Errorf("declared budget=%d, want 60000", got)
	}
}
