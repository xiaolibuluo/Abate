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

import (
	"reflect"
	"testing"
)

// TestTargetValidate 验证三类寻址模式与非法定义（CES-2.2.1）。
func TestTargetValidate(t *testing.T) {
	for _, k := range []TargetKind{TargetService, TargetInstance, TargetDependency} {
		tgt := Target{Kind: k, Name: "order"}
		if err := tgt.Validate(); err != nil {
			t.Errorf("kind %q rejected: %v", k, err)
		}
	}
	if err := (Target{Kind: "pod", Name: "x"}).Validate(); err == nil {
		t.Error("invalid kind must be rejected")
	}
	if err := (Target{Kind: TargetService}).Validate(); err == nil {
		t.Error("empty name must be rejected")
	}
}

// TestSelectorKeys 验证选择器键字典序升序（确定性要求）。
func TestSelectorKeys(t *testing.T) {
	tgt := Target{
		Kind: TargetService, Name: "order",
		Selector: map[string]string{"zone": "gz", "env": "staging", "app": "order"},
	}
	got := tgt.SelectorKeys()
	want := []string{"app", "env", "zone"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SelectorKeys=%v, want %v", got, want)
	}
	empty := Target{Kind: TargetService, Name: "order"}
	if keys := empty.SelectorKeys(); len(keys) != 0 {
		t.Errorf("empty selector keys=%v, want []", keys)
	}
}

// TestTargetMatches 验证全键等值匹配（逻辑与）。
func TestTargetMatches(t *testing.T) {
	tgt := Target{
		Kind: TargetService, Name: "order",
		Selector: map[string]string{"env": "staging"},
	}
	if !tgt.Matches(map[string]string{"env": "staging", "zone": "gz"}) {
		t.Error("must match superset attrs")
	}
	if tgt.Matches(map[string]string{"env": "prod"}) {
		t.Error("must not match differing value")
	}
	if tgt.Matches(map[string]string{}) {
		t.Error("must not match missing key")
	}
	noSel := Target{Kind: TargetService, Name: "order"}
	if !noSel.Matches(map[string]string{"any": "thing"}) {
		t.Error("empty selector must match everything")
	}
}
