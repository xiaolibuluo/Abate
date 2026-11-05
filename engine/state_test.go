// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package engine

import "testing"

// TestLegalTransitions 验证规范主路径：Pending→Preflight→Injecting→Observing→Recovering→Verdict（CES-3.1）。
func TestLegalTransitions(t *testing.T) {
	path := []State{
		StatePending, StatePreflight, StateInjecting,
		StateObserving, StateRecovering, StateVerdict,
	}
	for i := 0; i+1 < len(path); i++ {
		if !CanTransit(path[i], path[i+1]) {
			t.Errorf("%s -> %s must be legal", path[i], path[i+1])
		}
		next, err := Transition(path[i], path[i+1])
		if err != nil || next != path[i+1] {
			t.Errorf("Transition(%s,%s) = (%s,%v)", path[i], path[i+1], next, err)
		}
	}
}

// TestIllegalTransitions 验证非法迁移被拒绝并视为实现缺陷（CES-3.1）。
func TestIllegalTransitions(t *testing.T) {
	cases := [][2]State{
		{StatePending, StateInjecting},   // 跳过准入
		{StatePending, StateObserving},
		{StateInjecting, StateVerdict},   // 跳过恢复
		{StateObserving, StateVerdict},   // 跳过恢复
		{StateVerdict, StatePending},     // 终态不可迁出
		{StateRecovering, StateInjecting}, // 恢复中不可回注
	}
	for _, c := range cases {
		if CanTransit(c[0], c[1]) {
			t.Errorf("%s -> %s must be illegal", c[0], c[1])
		}
		if _, err := Transition(c[0], c[1]); err == nil {
			t.Errorf("Transition(%s,%s) must fail", c[0], c[1])
		}
	}
}

// TestMustRecover 验证红线触发的强制恢复不变量（CES-3.1.1）。
func TestMustRecover(t *testing.T) {
	for _, s := range []State{StatePreflight, StateInjecting, StateObserving} {
		if !MustRecover(s) {
			t.Errorf("red line in %s must force Recovering", s)
		}
		if !CanTransit(s, StateRecovering) {
			t.Errorf("%s -> Recovering must be legal", s)
		}
	}
	for _, s := range []State{StatePending, StateRecovering, StateVerdict} {
		if MustRecover(s) {
			t.Errorf("red line in %s must not require Recovering", s)
		}
	}
}

// TestTerminal 验证终态判定与状态全集。
func TestTerminal(t *testing.T) {
	if !StateVerdict.Terminal() {
		t.Error("Verdict must be terminal")
	}
	for _, s := range States() {
		if s != StateVerdict && s.Terminal() {
			t.Errorf("%s must not be terminal", s)
		}
	}
	if len(States()) != 6 {
		t.Errorf("len(States())=%d, want 6", len(States()))
	}
}
