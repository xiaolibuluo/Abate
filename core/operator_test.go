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
	"strings"
	"testing"
)

// TestOperatorClasses 验证五大算子类编码集合与固定顺序（CES-9.2）。
func TestOperatorClasses(t *testing.T) {
	got := OperatorClasses()
	want := []OperatorClass{
		OpArithmetic, OpRelational, OpConditional, OpReturnValue, OpStatement,
	}
	if len(got) != len(want) {
		t.Fatalf("len=%d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d = %q, want %q", i, got[i], want[i])
		}
	}
}

// TestOperatorClassValid 验证算子类合法与非法编码判定。
func TestOperatorClassValid(t *testing.T) {
	for _, c := range OperatorClasses() {
		if !c.Valid() {
			t.Errorf("%q must be valid", c)
		}
	}
	for _, c := range []OperatorClass{"", "M0", "M6", "m1"} {
		if c.Valid() {
			t.Errorf("%q must be invalid", c)
		}
	}
}

// TestMutationID 验证变异体全局标识格式 mut:{file}:{line}:{col}:{class}（CES-9.2.3）。
func TestMutationID(t *testing.T) {
	m := Mutation{File: "core/bp.go", Line: 42, Col: 7, Class: OpArithmetic}
	want := "mut:core/bp.go:42:7:M1"
	if m.ID() != want {
		t.Errorf("ID=%q, want %q", m.ID(), want)
	}
}

// TestMutationKind 验证变异体是静态扰动（CES-0.5.1）。
func TestMutationKind(t *testing.T) {
	m := Mutation{}
	if m.Kind() != PerturbStatic {
		t.Errorf("Kind=%q, want static", m.Kind())
	}
}

// TestMutationDescribe 验证人类可读描述包含关键字段。
func TestMutationDescribe(t *testing.T) {
	m := Mutation{
		File: "a.go", Line: 1, Col: 2,
		Class: OpRelational, From: "<", To: "<=",
	}
	d := m.Describe()
	for _, frag := range []string{"M2", "a.go:1:2", "<", "<="} {
		if !strings.Contains(d, frag) {
			t.Errorf("Describe %q missing %q", d, frag)
		}
	}
}

// TestMutationValidate 验证变异体定义校验的各分支。
func TestMutationValidate(t *testing.T) {
	ok := Mutation{
		File: "a.go", Line: 1, Col: 1,
		Class: OpStatement, From: "return x", To: "",
	}
	if err := ok.Validate(); err != nil {
		t.Errorf("valid mutation rejected: %v", err)
	}

	noFile := ok
	noFile.File = ""
	if err := noFile.Validate(); err == nil {
		t.Error("empty file must be rejected")
	}

	badPos := ok
	badPos.Line = 0
	if err := badPos.Validate(); err == nil {
		t.Error("line 0 must be rejected")
	}
	badCol := ok
	badCol.Col = 0
	if err := badCol.Validate(); err == nil {
		t.Error("col 0 must be rejected")
	}

	badClass := ok
	badClass.Class = "M9"
	if err := badClass.Validate(); err == nil {
		t.Error("invalid class must be rejected")
	}

	noop := ok
	noop.To = noop.From
	if err := noop.Validate(); err == nil {
		t.Error("from==to must be rejected (no perturbation)")
	}
}
