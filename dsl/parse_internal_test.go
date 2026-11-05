// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package dsl

import (
	"strings"
	"testing"

	"github.com/qingdujian/abate/core"
)

// TestErrorString 验证结构化错误的字符串格式（CES-2.1.2）。
func TestErrorString(t *testing.T) {
	e := Error{Path: "seed", Reason: "expected integer", Clause: "CES-2.1.2"}
	if got := e.Error(); got != "seed: expected integer (CES-2.1.2)" {
		t.Errorf("Error()=%q", got)
	}
}

// TestParseCommentsAndBlankLines 验证注释与空行被跳过。
func TestParseCommentsAndBlankLines(t *testing.T) {
	text := "# 头部注释\n\n" + goldenDSL + "\n# 尾部注释\n"
	e, err := Parse(text)
	if err != nil {
		t.Fatalf("Parse with comments: %v", err)
	}
	if e.ID != "exp-001" {
		t.Errorf("ID=%q", e.ID)
	}
}

// TestParseMissingColon 验证无冒号行被拒绝并包装为行号错误。
func TestParseMissingColon(t *testing.T) {
	_, err := Parse("plain text without colon\n")
	if err == nil {
		t.Fatal("line without colon must be rejected")
	}
	if !strings.Contains(err.Error(), "line 1") {
		t.Errorf("error must carry line number: %v", err)
	}
}

// TestParseEmptyKey 验证空键行被拒绝。
func TestParseEmptyKey(t *testing.T) {
	if _, err := Parse(": v\n"); err == nil {
		t.Error("empty key must be rejected")
	}
}

// TestParseDeepNesting 验证超过三层的缩进被拒绝（decode default 分支）。
func TestParseDeepNesting(t *testing.T) {
	text := "target:\n  name: x\n      deep: y\n"
	_, err := Parse(text)
	if err == nil {
		t.Fatal("6-space nesting must be rejected")
	}
	if !strings.Contains(err.Error(), "deeper than three levels") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestParseNestedOutsideList 验证无列表项先导的缩进字段被拒绝（listKey 为空分支）。
func TestParseNestedOutsideList(t *testing.T) {
	text := "steadyState:\n    sli: x\n"
	_, err := Parse(text)
	if err == nil {
		t.Fatal("nested field without list item must be rejected")
	}
	if !strings.Contains(err.Error(), "outside list item") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestApplyTopBranches 直测顶层字段分派的全部分支。
func TestApplyTopBranches(t *testing.T) {
	var e core.Experiment
	if err := applyTop(&e, "seed", "abc"); err == nil {
		t.Error("non-integer seed must error")
	}
	if err := applyTop(&e, "seed", "9"); err != nil || e.Seed != 9 {
		t.Errorf("seed=9: err=%v got=%d", err, e.Seed)
	}
	if err := applyTop(&e, "recoveryBudgetMs", "x"); err == nil {
		t.Error("non-integer recoveryBudgetMs must error")
	}
	if err := applyTop(&e, "recoveryBudgetMs", "60000"); err != nil || e.RecoveryBudgetMs != 60000 {
		t.Errorf("recoveryBudgetMs=60000: err=%v got=%d", err, e.RecoveryBudgetMs)
	}
	for _, section := range []string{"target", "fault", "blastRadius", "steadyState", "abortOn", "params"} {
		if err := applyTop(&e, section, ""); err != nil {
			t.Errorf("section header %q must be a no-op: %v", section, err)
		}
	}
	if err := applyTop(&e, "futureField", "v"); err != nil {
		t.Errorf("unknown field must be tolerated (CES-2.1.1): %v", err)
	}
}

// TestApplyBlastBranches 直测爆炸半径字段的全部分支（CES-2.5.1）。
func TestApplyBlastBranches(t *testing.T) {
	var e core.Experiment
	if err := applyBlast(&e, "maxTraffic", "500"); err != nil || e.Blast.MaxTraffic != 500 {
		t.Errorf("maxTraffic=500: err=%v got=%d", err, e.Blast.MaxTraffic)
	}
	if err := applyBlast(&e, "maxInstances", "abc"); err == nil {
		t.Error("non-integer maxInstances must error")
	}
	if err := applyBlast(&e, "unknownLimit", "10"); err != nil {
		t.Errorf("unknown blast key parsed but ignored: %v", err)
	}
}

// TestApplyListFieldEmptyList 直测空列表填充的拒绝分支与 abortOn 列表选择分支。
func TestApplyListFieldEmptyList(t *testing.T) {
	var e core.Experiment
	if err := applyListField(&e, "steadyState", "sli", "x"); err == nil {
		t.Error("steadyState empty list must error")
	}
	if err := applyListField(&e, "abortOn", "sli", "x"); err == nil {
		t.Error("abortOn empty list must error")
	}
}

// TestAssignAssertBranches 直测断言字段写入的全部分支（严格语义：windowMs 非整数报错）。
func TestAssignAssertBranches(t *testing.T) {
	var a core.SLIAssert
	if err := assignAssert(&a, "method", "p99"); err != nil {
		t.Errorf("method: %v", err)
	}
	if a.Method != "p99" {
		t.Errorf("method=%q", a.Method)
	}
	if err := assignAssert(&a, "windowMs", "60000"); err != nil {
		t.Errorf("windowMs: %v", err)
	}
	if a.WindowMs != 60000 {
		t.Errorf("windowMs=%d", a.WindowMs)
	}
	if err := assignAssert(&a, "windowMs", "abc"); err == nil {
		t.Error("non-integer windowMs must return structured error (CES-2.1.2)")
	}
	if a.WindowMs != 60000 {
		t.Errorf("failed windowMs must not alter value, got %d", a.WindowMs)
	}
	if err := assignAssert(&a, "unknown", "v"); err != nil {
		t.Errorf("unknown key must be a no-op: %v", err)
	}
	if a.SLI != "" || a.Assert != "" {
		t.Errorf("unknown key must not alter assert: %+v", a)
	}
}

// TestParseInvalidWindowMsViaDecode 验证 decode 路径透传 windowMs 结构化错误（覆盖 decode 的 applyListField 错误分支）。
func TestParseInvalidWindowMsViaDecode(t *testing.T) {
	text := strings.Replace(goldenDSL, "windowMs: 60000", "windowMs: abc", 1)
	_, err := Parse(text)
	if err == nil {
		t.Fatal("non-integer windowMs in list must be rejected via decode")
	}
	if !strings.Contains(err.Error(), "windowMs") {
		t.Errorf("error must name windowMs: %v", err)
	}
}

// TestParseInvalidListItemFirstFieldViaDecode 验证列表项首字段非法时 applyListItem 透传错误（覆盖其 assignAssert 错误分支）。
func TestParseInvalidListItemFirstFieldViaDecode(t *testing.T) {
	text := `id: exp-bad-window
seed: 1
criticality: 2
target:
  kind: service
  name: order
fault:
  type: F1
  delay: 100ms
blastRadius:
  maxInstances: 1
  maxDurationMs: 60000
steadyState:
  - windowMs: abc
`
	_, err := Parse(text)
	if err == nil {
		t.Fatal("non-integer windowMs as first list field must be rejected via decode")
	}
	if !strings.Contains(err.Error(), "windowMs") {
		t.Errorf("error must name windowMs: %v", err)
	}
}

// TestParseInvalidBlastViaDecode 验证 decode 路径透传 blastRadius 结构化错误（覆盖 decode 的 applySection 错误分支）。
func TestParseInvalidBlastViaDecode(t *testing.T) {
	text := strings.Replace(goldenDSL, "maxInstances: 1", "maxInstances: x", 1)
	_, err := Parse(text)
	if err == nil {
		t.Fatal("non-integer maxInstances must be rejected via decode")
	}
	if !strings.Contains(err.Error(), "maxInstances") {
		t.Errorf("error must name maxInstances: %v", err)
	}
}
