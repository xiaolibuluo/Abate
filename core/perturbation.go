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

import "fmt"

// PerturbationKind 区分两种扰动形态：变异与混沌在数学上同构，
// 都是对正常状态的受控偏离，仅在作用层与判定对象上不同。
// PerturbationKind distinguishes the two perturbation forms; mutation and
// chaos are mathematically isomorphic — controlled deviations from the
// normal state, differing only in layer and verdict subject.
type PerturbationKind string

const (
	// PerturbStatic 静态扰动（变异测试）：改动代码，跑测试，证明测试有效。
	// PerturbStatic is static perturbation (mutation testing): change the
	// code, run the tests, prove the tests are effective.
	PerturbStatic PerturbationKind = "static"
	// PerturbDynamic 动态扰动（混沌测试）：改动运行时，测 SLI，证明系统可靠。
	// PerturbDynamic is dynamic perturbation (chaos testing): change the
	// runtime, measure SLIs, prove the system is reliable.
	PerturbDynamic PerturbationKind = "dynamic"
)

// Valid 报告扰动形态是否合法。
// Valid reports whether the perturbation kind is defined.
func (k PerturbationKind) Valid() bool {
	return k == PerturbStatic || k == PerturbDynamic
}

// Outcome 是扰动判定的统一语义（双结合理论核心）：
// 变异体被测试杀死 = Contained；变异体存活逃逸 = Escaped。
// 混沌下稳态保持 = Contained；稳态被破坏逃逸成事故 = Escaped。
// 100% 全量覆盖的数学定义即 Escaped = 0。
//
// Outcome is the unified verdict semantics (the core of the dual theory):
// a mutant killed by tests is Contained, a surviving mutant is Escaped;
// steady state held under chaos is Contained, a violated steady state is
// Escaped. The mathematical definition of 100% total coverage is Escaped = 0.
type Outcome string

const (
	// OutcomeContained 扰动被控制在观测网内：质量目标达成。
	OutcomeContained Outcome = "contained"
	// OutcomeEscaped 扰动逃逸出观测网：暴露测试盲区或系统弱点。
	OutcomeEscaped Outcome = "escaped"
)

// Perturbation 是 HPOV 范式的扰动抽象：一切变异体与一切混沌故障
// 都是 Perturbation 的实现。内核只依赖此接口，不感知具体形态。
//
// Perturbation is the HPOV perturbation abstraction: every mutant and every
// chaos fault implements it. The kernel depends only on this interface and
// is agnostic to concrete forms.
type Perturbation interface {
	// ID 扰动全局唯一标识（确定性生成，同输入同 ID）。
	ID() string
	// Kind 扰动形态：static 或 dynamic。
	Kind() PerturbationKind
	// Describe 人类可读的扰动描述（用于报告与审计）。
	Describe() string
}

// Hypothesis 是 HPOV 第一步：对扰动后果的先验假设。
// 变异假设：『现有测试集会杀死此变异体』；混沌假设：『系统稳态成立』。
// Hypothesis is the first HPOV step: a prior assumption about the
// perturbation consequence.
type Hypothesis struct {
	// Statement 假设陈述（自然语言，审计可读）。
	Statement string `json:"statement"`
	// Subject 假设作用对象标识（包路径/服务名）。
	Subject string `json:"subject"`
}

// Validate 校验假设完整性。
// Validate checks hypothesis completeness.
func (h Hypothesis) Validate() error {
	if h.Statement == "" {
		return fmt.Errorf("hypothesis.statement: must not be empty")
	}
	if h.Subject == "" {
		return fmt.Errorf("hypothesis.subject: must not be empty")
	}
	return nil
}

// Observation 是 HPOV 第三步：一次扰动执行后的结构化观测。
// Observation is the third HPOV step: structured observation after one
// perturbation run.
type Observation struct {
	// PerturbationID 被观测扰动的 ID。
	PerturbationID string `json:"perturbationId"`
	// Outcome 统一判定语义。
	Outcome Outcome `json:"outcome"`
	// Detail 判定依据（测试失败摘要 / SLI 违例明细）。
	Detail string `json:"detail,omitempty"`
}

// Validate 校验观测完整性。
// Validate checks observation completeness.
func (o Observation) Validate() error {
	if o.PerturbationID == "" {
		return fmt.Errorf("observation.perturbationId: must not be empty")
	}
	if o.Outcome != OutcomeContained && o.Outcome != OutcomeEscaped {
		return fmt.Errorf("observation.outcome: invalid %q", o.Outcome)
	}
	return nil
}
