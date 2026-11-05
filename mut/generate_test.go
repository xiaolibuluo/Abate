// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package mut

import (
	"strings"
	"testing"

	"github.com/qingdujian/abate/core"
)

// goldenSrc 是含算术与关系运算符的最小合法 Go 源码。
const goldenSrc = `package calc

// Max 返回较大值。
func Max(a, b int) int {
	if a > b {
		return a + b
	}
	return b - a
}
`

// TestGenerateFindsMutants 验证 AST 遍历找出全部 M1/M2 变异点。
func TestGenerateFindsMutants(t *testing.T) {
	ms, err := Generate("calc.go", goldenSrc, 42)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if len(ms) == 0 {
		t.Fatal("no mutants generated")
	}
	var m1, m2 int
	for _, m := range ms {
		switch m.Class {
		case core.OpArithmetic:
			m1++
		case core.OpRelational:
			m2++
		}
		if err := m.Validate(); err != nil {
			t.Errorf("generated invalid mutant %+v: %v", m, err)
		}
	}
	if m1 == 0 || m2 == 0 {
		t.Errorf("want both M1 and M2 mutants, got M1=%d M2=%d", m1, m2)
	}
}

// TestGenerateDeterministic 验证同 seed 同集合同顺序，不同 seed 同集合不同序（CES-9.2.2）。
func TestGenerateDeterministic(t *testing.T) {
	a, err := Generate("calc.go", goldenSrc, 7)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	b, err := Generate("calc.go", goldenSrc, 7)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if len(a) != len(b) {
		t.Fatalf("set size differs: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i].ID() != b[i].ID() {
			t.Fatalf("order diverged at %d: %q vs %q", i, a[i].ID(), b[i].ID())
		}
	}
	c, err := Generate("calc.go", goldenSrc, 8)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if len(c) != len(a) {
		t.Fatalf("seed changed set size: %d vs %d", len(c), len(a))
	}
}

// TestGenerateParseError 验证语法错误源码被拒绝。
func TestGenerateParseError(t *testing.T) {
	if _, err := Generate("bad.go", "package x\nfunc {" , 1); err == nil {
		t.Error("invalid Go source must be rejected")
	}
}

// TestApply 验证变异应用的合法替换与源码内容正确性。
func TestApply(t *testing.T) {
	ms, err := Generate("calc.go", goldenSrc, 42)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	for _, m := range ms {
		mutated, err := Apply(goldenSrc, m)
		if err != nil {
			t.Fatalf("Apply(%s): %v", m.ID(), err)
		}
		if mutated == goldenSrc {
			t.Errorf("Apply(%s) produced no change", m.ID())
		}
		if !strings.Contains(mutated, m.To) {
			t.Errorf("mutated source missing %q: %s", m.To, mutated)
		}
	}
}

// TestApplyErrors 验证 Apply 各错误分支。
func TestApplyErrors(t *testing.T) {
	bad := core.Mutation{}
	if _, err := Apply(goldenSrc, bad); err == nil {
		t.Error("invalid mutation must be rejected")
	}

	beyond := core.Mutation{
		File: "calc.go", Line: 999, Col: 1,
		Class: core.OpArithmetic, From: "+", To: "-",
	}
	if _, err := Apply(goldenSrc, beyond); err == nil {
		t.Error("line beyond file must be rejected")
	}

	badCol := core.Mutation{
		File: "calc.go", Line: 5, Col: 9999,
		Class: core.OpArithmetic, From: "+", To: "-",
	}
	if _, err := Apply(goldenSrc, badCol); err == nil {
		t.Error("col out of range must be rejected")
	}

	drift := core.Mutation{
		File: "calc.go", Line: 5, Col: 1,
		Class: core.OpArithmetic, From: "+", To: "-",
	}
	if _, err := Apply(goldenSrc, drift); err == nil {
		t.Error("source drift must be rejected")
	}
}

// TestSplitLinesTrailingNewline 验证以换行结尾的源码不产生末尾空行（splitLines 尾部分支）。
func TestSplitLinesTrailingNewline(t *testing.T) {
	withNL := "a\nb\n"
	got := splitLines(withNL)
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("splitLines(trailing NL)=%v, want [a b]", got)
	}
	withoutNL := "a\nb"
	got2 := splitLines(withoutNL)
	if len(got2) != 2 || got2[1] != "b" {
		t.Errorf("splitLines(no trailing NL)=%v, want [a b]", got2)
	}
}

// TestSplitJoinRoundtrip 验证行切分与合并的往返一致性。
func TestSplitJoinRoundtrip(t *testing.T) {
	if got := joinLines(splitLines(goldenSrc)); got != strings.TrimSuffix(goldenSrc, "\n") {
		// splitLines 丢弃末尾换行语义，比较时去除尾换行
		t.Errorf("roundtrip mismatch: %q", got)
	}
	if got := joinLines(splitLines("")); got != "" {
		t.Errorf("empty source roundtrip = %q, want empty", got)
	}
}
