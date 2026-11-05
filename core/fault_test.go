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

// TestFaultClassCodes 验证六分类编码集合与固定顺序（CES 第 1 章）。
func TestFaultClassCodes(t *testing.T) {
	got := FaultClasses()
	want := []FaultClass{F1Latency, F2Crash, F3Partition, F4Exhaustion, F5Corruption, F6ClockSkew}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d = %q, want %q", i, got[i], want[i])
		}
	}
}

// TestFaultClassValid 验证合法编码与非法编码的判定。
func TestFaultClassValid(t *testing.T) {
	for _, c := range FaultClasses() {
		if !c.Valid() {
			t.Errorf("%q must be valid", c)
		}
	}
	for _, c := range []FaultClass{"", "F0", "F7", "f1"} {
		if c.Valid() {
			t.Errorf("%q must be invalid", c)
		}
	}
}

// TestFaultValidate 验证故障定义校验（CES-2.3.1）。
func TestFaultValidate(t *testing.T) {
	ok := Fault{Class: F1Latency, Params: map[string]string{"delay": "500ms"}}
	if err := ok.Validate(); err != nil {
		t.Errorf("valid fault rejected: %v", err)
	}
	badClass := Fault{Class: "F9", Params: map[string]string{"a": "b"}}
	if err := badClass.Validate(); err == nil {
		t.Error("invalid class must be rejected")
	}
	noParams := Fault{Class: F2Crash}
	if err := noParams.Validate(); err == nil {
		t.Error("empty params must be rejected")
	}
}
