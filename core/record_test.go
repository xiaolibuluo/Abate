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

// TestRecordWeight 验证 w_e = c_e × e_e（CES-4.3.1）：关键度 4 × A 级 = 40000。
func TestRecordWeight(t *testing.T) {
	r := goldenRecord("exp-w")
	w, err := r.Weight()
	if err != nil {
		t.Fatalf("Weight: %v", err)
	}
	if w != 40000 {
		t.Errorf("Weight=%d, want 40000 (4 × 10000)", w)
	}

	r.Evidence = EvidenceB
	w, err = r.Weight()
	if err != nil {
		t.Fatalf("Weight: %v", err)
	}
	if w != 32000 {
		t.Errorf("Weight=%d, want 32000 (4 × 8000)", w)
	}

	r.Evidence = "X"
	if _, err := r.Weight(); err == nil {
		t.Error("invalid evidence must return error")
	}
}

// TestRecoveryScoreBranches 验证恢复分各分支（CES-4.4.2）。
func TestRecoveryScoreBranches(t *testing.T) {
	r := goldenRecord("exp-r")

	s, ok := r.RecoveryScore()
	if !ok || s != BPFull {
		t.Errorf("zero recovery = (%d,%v), want (10000,true)", s, ok)
	}

	r.RecoveryMs = 150000
	s, ok = r.RecoveryScore()
	if !ok || s != 5000 {
		t.Errorf("half budget = (%d,%v), want (5000,true)", s, ok)
	}

	r.RecoveryMs = 600000
	s, ok = r.RecoveryScore()
	if !ok || s != 0 {
		t.Errorf("over budget = (%d,%v), want (0,true)", s, ok)
	}

	r.RecoveryMs = -1
	if _, ok = r.RecoveryScore(); ok {
		t.Error("no recovery event must be unmeasured")
	}
}

// TestDegradationScore 验证降级分（CES-4.4.3）：开关未激活归零，激活后按达标率折算。
func TestDegradationScore(t *testing.T) {
	r := goldenRecord("exp-d")
	if got := r.DegradationScore(); got != BPFull {
		t.Errorf("full = %d, want 10000", got)
	}

	r.Degraded = false
	if got := r.DegradationScore(); got != 0 {
		t.Errorf("switch off = %d, want 0", got)
	}

	r.Degraded = true
	r.DegradedPass, r.DegradedTotal = 800, 1000
	if got := r.DegradationScore(); got != 8000 {
		t.Errorf("80%% quality = %d, want 8000", got)
	}

	r.DegradedPass, r.DegradedTotal = 0, 0
	if got := r.DegradationScore(); got != 0 {
		t.Errorf("zero requests = %d, want 0", got)
	}
}

// TestRadiusScoreBranches 验证半径控制分（CES-4.4.5）。
func TestRadiusScoreBranches(t *testing.T) {
	r := goldenRecord("exp-br")
	if got := r.RadiusScore(); got != BPFull {
		t.Errorf("declared==observed = %d, want 10000", got)
	}

	r.ObservedRadius = 0
	if got := r.RadiusScore(); got != BPFull {
		t.Errorf("zero observed = %d, want 10000", got)
	}

	r.ObservedRadius = 2
	if got := r.RadiusScore(); got != 5000 {
		t.Errorf("double observed = %d, want 5000", got)
	}

	r.ObservedRadius = 100
	if got := r.RadiusScore(); got != 100 {
		t.Errorf("hundredfold observed = %d, want 100", got)
	}

	r.RedlineLate = true
	if got := r.RadiusScore(); got != 0 {
		t.Errorf("redline late = %d, want 0", got)
	}
}

// TestIntegrityScore 验证不变量通过率（CES-4.4.4）。
func TestIntegrityScore(t *testing.T) {
	r := goldenRecord("exp-i")
	if got := r.IntegrityScore(); got != BPFull {
		t.Errorf("full = %d, want 10000", got)
	}
	r.InvariantPass = 7
	if got := r.IntegrityScore(); got != 7000 {
		t.Errorf("7/10 = %d, want 7000", got)
	}
	r.InvariantTotal = 0
	if got := r.IntegrityScore(); got != 0 {
		t.Errorf("zero invariants = %d, want 0", got)
	}
}
