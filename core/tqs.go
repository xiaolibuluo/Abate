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

// TQSInput 是统一质量分的三层输入（万分位分值）。
// TQSInput carries the three scoring layers of the Total Quality Score.
type TQSInput struct {
	// Mutation 变异杀死率：被杀死变异体占全部非等价变异体的万分位比率。
	// Mutation is the mutation kill ratio in basis points over non-equivalent mutants.
	Mutation BP `json:"mutation"`
	// Chaos 混沌韧性分：CES 第 4 章 Score_final（万分位）。
	// Chaos is the CES Chapter 4 reliability Score_final in basis points.
	Chaos BP `json:"chaos"`
	// Coverage 代码语句覆盖率（万分位）。
	// Coverage is the code statement coverage ratio in basis points.
	Coverage BP `json:"coverage"`
}

// Validate 校验三层输入均在 [0, 10000] 区间。
// Validate checks that all three layers are within [0, 10000].
func (in TQSInput) Validate() error {
	for name, v := range map[string]BP{
		"mutation": in.Mutation, "chaos": in.Chaos, "coverage": in.Coverage,
	} {
		if v < 0 || v > BPFull {
			return fmt.Errorf("tqs.%s: out of range [0,10000], got %d", name, v)
		}
	}
	return nil
}

// TQS 是统一质量分结果：几何聚合，任一层为零则总分为零——
// 100% 全量覆盖的硬语义：短板即 verdict，不允许长板补偿短板。
//
// TQS is the unified quality score: geometric aggregation — if any layer is
// zero the total is zero. This is the hard semantics of 100% total coverage:
// the weakest link is the verdict; no long-board compensation.
type TQS struct {
	// Total 总分（万分位；展示 ÷100 得 0.00–100.00）。
	Total BP `json:"total"`
	// Layers 三层原值。
	Layers TQSInput `json:"layers"`
	// Gate 准入准出门禁判定：三层均达各自阈值才通过。
	// Gate is the admission/exit verdict: all layers must pass thresholds.
	Gate bool `json:"gate"`
}

// TQSThresholds 是门禁阈值（万分位）：默认全量红线——变异 10000、覆盖 10000、混沌 9500。
// TQSThresholds are gate thresholds; defaults are the total red lines.
type TQSThresholds struct {
	Mutation BP
	Chaos    BP
	Coverage BP
}

// DefaultThresholds 返回默认门禁：变异与覆盖 100%，混沌 95%（CES-4.5 常量）。
// DefaultThresholds returns the default gates: mutation and coverage 100%,
// chaos 95% (the CES-4.5 constant).
func DefaultThresholds() TQSThresholds {
	return TQSThresholds{Mutation: BPFull, Chaos: BPTarget, Coverage: BPFull}
}

// ComputeTQS 计算统一质量分（万分位定点整数，CES-4.1.2 同源纪律）。
// 总分采用几何均值立方根的整数近似：Total = (M×C×V)^(1/3)，
// 用乘法-开方分解实现，无浮点参与。
//
// ComputeTQS computes the unified score in basis-point fixed-point integer
// arithmetic. The total is the geometric mean: Total = (M×C×V)^(1/3),
// realized by multiply-then-integer-cube-root, no floating point involved.
func ComputeTQS(in TQSInput, th TQSThresholds) (TQS, error) {
	if err := in.Validate(); err != nil {
		return TQS{}, err
	}
	prod := in.Mutation * in.Chaos * in.Coverage
	total := BP(icbrt(int64(prod)))
	gate := in.Mutation >= th.Mutation &&
		in.Chaos >= th.Chaos &&
		in.Coverage >= th.Coverage
	return TQS{Total: total, Layers: in, Gate: gate}, nil
}

// icbrt 计算非负整数的整数立方根（向零取整），牛顿迭代，纯整数。
// icbrt computes the integer cube root of a non-negative integer
// (truncating), via Newton iteration in pure integers.
func icbrt(n int64) int64 {
	if n <= 0 {
		return 0
	}
	// 初始上界：n^(1/3) <= n，用二分搜索，避免牛顿迭代溢出问题。
	// Binary search on [0, n], avoiding Newton overflow concerns.
	lo, hi := int64(0), n
	if hi > 1<<21 { // (1<<21)^3 > 2^63 上界收紧防爆
		hi = 1 << 21
	}
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if mid*mid*mid <= n {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}
