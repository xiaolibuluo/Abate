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

// TestEvidenceLevelValid 验证三个法定等级与非法等级（CES-5.1）。
func TestEvidenceLevelValid(t *testing.T) {
	for _, l := range []EvidenceLevel{EvidenceA, EvidenceB, EvidenceC} {
		if !l.Valid() {
			t.Errorf("%q must be valid", l)
		}
	}
	for _, l := range []EvidenceLevel{"", "D", "a", "AA"} {
		if l.Valid() {
			t.Errorf("%q must be invalid", l)
		}
	}
}

// TestEvidenceCoefficient 验证等级系数 A=10000, B=8000, C=5000（CES-4.3.1）。
func TestEvidenceCoefficient(t *testing.T) {
	cases := []struct {
		l    EvidenceLevel
		want BP
	}{
		{EvidenceA, 10000},
		{EvidenceB, 8000},
		{EvidenceC, 5000},
	}
	for _, c := range cases {
		got, err := c.l.Coefficient()
		if err != nil {
			t.Fatalf("Coefficient(%q): %v", c.l, err)
		}
		if got != c.want {
			t.Errorf("Coefficient(%q)=%d, want %d", c.l, got, c.want)
		}
	}
	if _, err := EvidenceLevel("X").Coefficient(); err == nil {
		t.Error("invalid level must return error")
	}
}

// TestEvidenceDiscountOrder 验证低严谨度证据权重严格更低（CES-4.3.2）。
func TestEvidenceDiscountOrder(t *testing.T) {
	a, _ := EvidenceA.Coefficient()
	b, _ := EvidenceB.Coefficient()
	c, _ := EvidenceC.Coefficient()
	if !(a > b && b > c) {
		t.Errorf("coefficients must strictly decrease: A=%d B=%d C=%d", a, b, c)
	}
}
