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

// EvidenceLevel 是 CES-5.1 定义的证据等级：实验结论严谨度的分级。
// EvidenceLevel is the evidence rigor grade defined by CES-5.1.
type EvidenceLevel string

const (
	// EvidenceA 控制组对比 + 显著性检验（如差分中的差分），生产与高利害判定专用。
	// EvidenceA uses control-group comparison with significance testing;
	// required for causal claims and high-stakes verdicts.
	EvidenceA EvidenceLevel = "A"
	// EvidenceB 单组时序基线对比，预发与常规演练适用。
	// EvidenceB uses single-group time-series baseline comparison.
	EvidenceB EvidenceLevel = "B"
	// EvidenceC 阈值判定，冒烟与开发自测适用。
	// EvidenceC uses threshold verdicts for smoke and developer self-tests.
	EvidenceC EvidenceLevel = "C"
)

// Valid 报告证据等级是否合法。
// Valid reports whether the evidence level is defined by the specification.
func (l EvidenceLevel) Valid() bool {
	switch l {
	case EvidenceA, EvidenceB, EvidenceC:
		return true
	}
	return false
}

// Coefficient 返回 CES-4.3.1 定义的证据等级系数 e_e（万分位）：
// A = 10000，B = 8000，C = 5000。低严谨度证据必须折抵权重（CES-4.3.2）。
//
// Coefficient returns the evidence-level coefficient e_e in basis points
// (CES-4.3.1): A = 10000, B = 8000, C = 5000. Lower-rigor evidence is
// discounted in scoring weight (CES-4.3.2).
func (l EvidenceLevel) Coefficient() (BP, error) {
	switch l {
	case EvidenceA:
		return BPFull, nil
	case EvidenceB:
		return 8000, nil
	case EvidenceC:
		return 5000, nil
	}
	return 0, fmt.Errorf("evidence level %q has no coefficient (CES-4.3.1)", l)
}
