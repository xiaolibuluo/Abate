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

// TestTQSFullMarks 验证三层满分时总分为 10000 且门禁通过（CES-0.5.4）。
func TestTQSFullMarks(t *testing.T) {
	got, err := ComputeTQS(TQSInput{Mutation: 10000, Chaos: 10000, Coverage: 10000},
		DefaultThresholds())
	if err != nil {
		t.Fatalf("ComputeTQS: %v", err)
	}
	if got.Total != BPFull {
		t.Errorf("Total=%d, want 10000", got.Total)
	}
	if !got.Gate {
		t.Error("full marks must pass the gate")
	}
}

// TestTQSGeometric 验证几何聚合：三层 8000 时总分为 8000（立方根(8000³)=8000）。
func TestTQSGeometric(t *testing.T) {
	got, err := ComputeTQS(TQSInput{Mutation: 8000, Chaos: 8000, Coverage: 8000},
		DefaultThresholds())
	if err != nil {
		t.Fatalf("ComputeTQS: %v", err)
	}
	if got.Total != 8000 {
		t.Errorf("Total=%d, want 8000", got.Total)
	}
	if got.Gate {
		t.Error("8000 must fail 100%% thresholds")
	}
}

// TestTQSZeroLayerZeroTotal 验证任一层为零则总分为零（短板即 verdict）。
func TestTQSZeroLayerZeroTotal(t *testing.T) {
	for _, in := range []TQSInput{
		{Mutation: 0, Chaos: 10000, Coverage: 10000},
		{Mutation: 10000, Chaos: 0, Coverage: 10000},
		{Mutation: 10000, Chaos: 10000, Coverage: 0},
	} {
		got, err := ComputeTQS(in, DefaultThresholds())
		if err != nil {
			t.Fatalf("ComputeTQS(%+v): %v", in, err)
		}
		if got.Total != 0 {
			t.Errorf("%+v: Total=%d, want 0", in, got.Total)
		}
		if got.Gate {
			t.Errorf("%+v: gate must fail", in)
		}
	}
}

// TestTQSGateThresholds 验证门禁三层独立判定（混沌阈值 9500，变异与覆盖 10000）。
func TestTQSGateThresholds(t *testing.T) {
	// 混沌达 9500 但变异未满 10000：不通过
	in := TQSInput{Mutation: 9999, Chaos: 9500, Coverage: 10000}
	got, err := ComputeTQS(in, DefaultThresholds())
	if err != nil {
		t.Fatalf("ComputeTQS: %v", err)
	}
	if got.Gate {
		t.Error("mutation 9999 must fail the 10000 gate")
	}

	// 混沌 9499 低于 9500：不通过
	in = TQSInput{Mutation: 10000, Chaos: 9499, Coverage: 10000}
	got, err = ComputeTQS(in, DefaultThresholds())
	if err != nil {
		t.Fatalf("ComputeTQS: %v", err)
	}
	if got.Gate {
		t.Error("chaos 9499 must fail the 9500 gate")
	}

	// 全部恰好达阈值：通过
	in = TQSInput{Mutation: 10000, Chaos: 9500, Coverage: 10000}
	got, err = ComputeTQS(in, DefaultThresholds())
	if err != nil {
		t.Fatalf("ComputeTQS: %v", err)
	}
	if !got.Gate {
		t.Error("exact thresholds must pass the gate")
	}
}

// TestTQSValidate 验证输入范围校验。
func TestTQSValidate(t *testing.T) {
	if err := (TQSInput{Mutation: 10000, Chaos: 10000, Coverage: 10000}).Validate(); err != nil {
		t.Errorf("valid input rejected: %v", err)
	}
	if err := (TQSInput{Mutation: -1}).Validate(); err == nil {
		t.Error("negative mutation must be rejected")
	}
	if err := (TQSInput{Mutation: 10000, Chaos: 10001}).Validate(); err == nil {
		t.Error("chaos > 10000 must be rejected")
	}
	if _, err := ComputeTQS(TQSInput{Mutation: 10001}, DefaultThresholds()); err == nil {
		t.Error("invalid input must propagate error from ComputeTQS")
	}
}

// TestIcbrt 验证整数立方根的边界与精确值。
func TestIcbrt(t *testing.T) {
	cases := []struct {
		n, want int64
	}{
		{0, 0}, {-5, 0}, {1, 1}, {7, 1}, {8, 2},
		{26, 2}, {27, 3}, {1000, 10}, {1000000000000, 10000},
		{999, 9}, {999999999999, 9999},
	}
	for _, c := range cases {
		if got := icbrt(c.n); got != c.want {
			t.Errorf("icbrt(%d)=%d, want %d", c.n, got, c.want)
		}
	}
	// 最大输入防爆：10000³ = 1e12，icbrt 须精确回到 10000
	if got := icbrt(10000 * 10000 * 10000); got != 10000 {
		t.Errorf("icbrt(1e12)=%d, want 10000", got)
	}
}

// TestDefaultThresholds 锁定默认门禁常量（变异/覆盖 100%，混沌 95%）。
func TestDefaultThresholds(t *testing.T) {
	th := DefaultThresholds()
	if th.Mutation != 10000 || th.Coverage != 10000 || th.Chaos != 9500 {
		t.Errorf("DefaultThresholds=%+v, want {10000 9500 10000}", th)
	}
}
