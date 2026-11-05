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

// goldenDSL 是一个合法实验的完整 DSL 文本。
const goldenDSL = `id: exp-001
seed: 42
criticality: 4
target:
  kind: service
  name: order
  env: staging
fault:
  type: F1
  delay: 500ms
blastRadius:
  maxInstances: 1
  maxDurationMs: 300000
steadyState:
  - sli: http_success_rate
    assert: ">= 99.9%"
    windowMs: 60000
abortOn:
  - sli: error_rate
    assert: "> 5%"
`

// TestParseGolden 验证合法实验完整解析（CES 第 2 章五元组）。
func TestParseGolden(t *testing.T) {
	e, err := Parse(goldenDSL)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if e.ID != "exp-001" || e.Seed != 42 || e.Criticality != 4 {
		t.Errorf("top fields: %+v", e)
	}
	if e.Target.Kind != core.TargetService || e.Target.Name != "order" {
		t.Errorf("target: %+v", e.Target)
	}
	if e.Target.Selector["env"] != "staging" {
		t.Errorf("selector: %v", e.Target.Selector)
	}
	if e.Fault.Class != core.F1Latency || e.Fault.Params["delay"] != "500ms" {
		t.Errorf("fault: %+v", e.Fault)
	}
	if e.Blast.MaxInstances != 1 || e.Blast.MaxDurationMs != 300000 {
		t.Errorf("blast: %+v", e.Blast)
	}
	if len(e.Steady) != 1 || e.Steady[0].SLI != "http_success_rate" {
		t.Errorf("steady: %+v", e.Steady)
	}
	if e.Steady[0].WindowMs != 60000 {
		t.Errorf("steady window: %+v", e.Steady[0])
	}
	if len(e.AbortOn) != 1 || e.AbortOn[0].Assert != "> 5%" {
		t.Errorf("abortOn: %+v", e.AbortOn)
	}
}

// TestParseQuotedScalars 验证带引号标量的首尾引号被剥离（YAML 标量语义）。
func TestParseQuotedScalars(t *testing.T) {
	e, err := Parse(goldenDSL)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if e.AbortOn[0].Assert != "> 5%" {
		t.Errorf("quoted assert = %q, want %q", e.AbortOn[0].Assert, "> 5%")
	}
	text := strings.Replace(goldenDSL, "assert: \">= 99.9%\"", "assert: '>= 99.9%'", 1)
	e2, err := Parse(text)
	if err != nil {
		t.Fatalf("Parse single-quoted: %v", err)
	}
	if e2.Steady[0].Assert != ">= 99.9%" {
		t.Errorf("single-quoted assert = %q, want %q", e2.Steady[0].Assert, ">= 99.9%")
	}
}

// TestParseUnknownFieldTolerated 验证未知字段不拒绝（CES-2.1.1）。
func TestParseUnknownFieldTolerated(t *testing.T) {
	text := "futureField: hello\n" + goldenDSL
	if _, err := Parse(text); err != nil {
		t.Errorf("unknown top-level field must be tolerated: %v", err)
	}
}

// TestParseStructuredError 验证校验失败输出结构化错误（CES-2.1.2）。
func TestParseStructuredError(t *testing.T) {
	text := strings.Replace(goldenDSL, "criticality: 4", "criticality: abc", 1)
	_, err := Parse(text)
	if err == nil {
		t.Fatal("non-integer criticality must be rejected")
	}
	dslErr, ok := err.(Error)
	if !ok {
		t.Fatalf("error type = %T, want dsl.Error", err)
	}
	if dslErr.Path != "criticality" || dslErr.Clause != "CES-2.1.2" {
		t.Errorf("structured error: %+v", dslErr)
	}
}

// TestParseIndentError 验证奇数缩进被拒绝。
func TestParseIndentError(t *testing.T) {
	text := "id: exp-x\n criticality: 4\n"
	if _, err := Parse(text); err == nil {
		t.Error("odd indentation must be rejected")
	}
}

// TestParseSpecValidation 验证解析后执行规范校验（缺稳态断言必须拒绝，CES-2.6.1）。
func TestParseSpecValidation(t *testing.T) {
	text := `id: exp-y
criticality: 3
target:
  kind: service
  name: order
fault:
  type: F2
  signal: SIGKILL
blastRadius:
  maxInstances: 1
  maxDurationMs: 60000
`
	if _, err := Parse(text); err == nil {
		t.Error("missing steadyState must be rejected by spec validation")
	}
}

// TestParseListOutsideSection 验证列表项出现在非法节被拒绝。
func TestParseListOutsideSection(t *testing.T) {
	text := `id: exp-z
fault:
  - sli: x
`
	if _, err := Parse(text); err == nil {
		t.Error("list item under fault must be rejected")
	}
}
