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

// OperatorClass 是变异算子分类编码，属守恒层：
// 五大经典算子类自 1970 年代变异测试学科定型以来未变，编码永久固定，
// 只可标记废弃，不得回收复用——与故障六分类（fault.go）同级冻结。
//
// OperatorClass is the mutation operator classification, part of the
// conserved layer: the five classic operator classes have been stable since
// the discipline formed in the 1970s. Codes are permanent.
type OperatorClass string

const (
	// OpArithmetic 算术变异：+/-/*/÷/% 互换（AOR）。
	OpArithmetic OperatorClass = "M1"
	// OpRelational 关系变异：< <= > >= == != 互换与边界偏移（ROR）。
	OpRelational OperatorClass = "M2"
	// OpConditional 条件变异：逻辑与/或互换、条件取反（COR/LOR）。
	OpConditional OperatorClass = "M3"
	// OpReturnValue 返回值变异：返回零值/nil/反值（RVR）。
	OpReturnValue OperatorClass = "M4"
	// OpStatement 语句变异：删除或复制语句（SDL）。
	OpStatement OperatorClass = "M5"
)

// OperatorClasses 返回全部合法算子类，顺序固定（确定性要求）。
// OperatorClasses returns all valid operator classes in a fixed order.
func OperatorClasses() []OperatorClass {
	return []OperatorClass{
		OpArithmetic, OpRelational, OpConditional,
		OpReturnValue, OpStatement,
	}
}

// Valid 报告算子类编码是否合法。
// Valid reports whether the operator class code is defined.
func (c OperatorClass) Valid() bool {
	switch c {
	case OpArithmetic, OpRelational, OpConditional,
		OpReturnValue, OpStatement:
		return true
	}
	return false
}

// Mutation 是一个具体变异体：在某个源位置应用某类算子产生的扰动。
// Mutation is one concrete mutant: the perturbation produced by applying an
// operator class at a source location.
type Mutation struct {
	// File 源文件路径（仓库相对）。
	File string `json:"file"`
	// Line 变异点起始行（1-based）。
	Line int `json:"line"`
	// Col 变异点起始列（1-based）。
	Col int `json:"col"`
	// Class 应用的算子类。
	Class OperatorClass `json:"class"`
	// From 原代码片段。
	From string `json:"from"`
	// To 变异后代码片段。
	To string `json:"to"`
}

// ID 实现 Perturbation 接口：确定性全局唯一标识。
// ID implements Perturbation: a deterministic globally unique identifier.
func (m Mutation) ID() string {
	return fmt.Sprintf("mut:%s:%d:%d:%s", m.File, m.Line, m.Col, m.Class)
}

// Kind 实现 Perturbation 接口：变异体是静态扰动。
// Kind implements Perturbation: a mutant is a static perturbation.
func (m Mutation) Kind() PerturbationKind { return PerturbStatic }

// Describe 实现 Perturbation 接口：人类可读描述。
// Describe implements Perturbation with a human-readable description.
func (m Mutation) Describe() string {
	return fmt.Sprintf("%s %s:%d:%d %q -> %q",
		m.Class, m.File, m.Line, m.Col, m.From, m.To)
}

// Validate 校验变异体定义完整性。
// Validate checks mutant definition completeness.
func (m Mutation) Validate() error {
	if m.File == "" {
		return fmt.Errorf("mutation.file: must not be empty")
	}
	if m.Line < 1 || m.Col < 1 {
		return fmt.Errorf("mutation position: line/col must be >= 1, got %d:%d",
			m.Line, m.Col)
	}
	if !m.Class.Valid() {
		return fmt.Errorf("mutation.class: invalid %q", m.Class)
	}
	if m.From == m.To {
		return fmt.Errorf("mutation: from equals to, no perturbation produced")
	}
	return nil
}
