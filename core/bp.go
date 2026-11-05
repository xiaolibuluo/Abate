// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package core 定义 CES 规范的核心领域类型与插件接口。
// Package core defines the CES domain types and plugin interfaces.
package core

import "sort"

// BP 是万分位（basis point, ‱）定点整数：10000 表示 100.00 分 / 100%。
// CES-4.1.2：评分链路禁止使用浮点数，一律以 BP 整数运算，
// 消除跨语言、跨平台的舍入漂移。
//
// BP is a basis-point (‱) fixed-point integer: 10000 means 100.00 / 100%.
// CES-4.1.2: floating point is forbidden in the scoring path; BP integer
// arithmetic only, eliminating cross-language and cross-platform rounding drift.
type BP int64

const (
	// BPFull 满分/全额：10000‱ = 100.00。
	// BPFull is full marks: 10000‱ = 100.00.
	BPFull BP = 10000
	// BPTarget 覆盖目标比率 M：9500‱ = 95%（CES-4.5）。
	// BPTarget is the coverage target ratio M: 9500‱ = 95% (CES-4.5).
	BPTarget BP = 9500
)

// MulDiv 计算 a × b ÷ d，乘法先于除法，整数除法向零取整（CES-4.1.3）。
// d 为 0 时返回 0：除零在评分语义中等价于『无有效样本』，由调用方按缺测处理。
//
// MulDiv computes a × b ÷ d with multiplication before division, truncating
// toward zero (CES-4.1.3). A zero divisor yields 0, which the caller treats
// as the unmeasured case.
func (a BP) MulDiv(b BP, d BP) BP {
	if d == 0 {
		return 0
	}
	return a * b / d
}

// Clamp 将 v 收敛到 [lo, hi] 区间。
// Clamp constrains v to the inclusive range [lo, hi].
func (v BP) Clamp(lo, hi BP) BP {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// Min 返回两值中的较小者。
// Min returns the smaller of the two values.
func (v BP) Min(o BP) BP {
	if v < o {
		return v
	}
	return o
}

// Max 返回两值中的较大者。
// Max returns the greater of the two values.
func (v BP) Max(o BP) BP {
	if v > o {
		return v
	}
	return o
}

// Ratio 计算 part × 10000 ÷ whole 的万分位比率；whole 为 0 时返回 0。
// Ratio computes part × 10000 ÷ whole as basis points; a zero whole yields 0.
func Ratio(part, whole int64) BP {
	if whole == 0 {
		return 0
	}
	return BP(part * int64(BPFull) / whole)
}

// Weighted 计算加权聚合：Σ(w×v) ÷ Σ(w)；权重和为 0 返回 0 与 false（缺测）。
// Weighted computes Σ(w×v) ÷ Σ(w); a zero weight sum returns 0 and false,
// signalling the unmeasured case per CES-4.6.
func Weighted(weights, values []BP) (BP, bool) {
	if len(weights) != len(values) {
		return 0, false
	}
	var num, den BP
	for i := range weights {
		num += weights[i] * values[i]
		den += weights[i]
	}
	if den == 0 {
		return 0, false
	}
	return num / den, true
}

// SortedIDs 返回字典序升序的 ID 副本（CES-4.1.3 聚合前排序要求）。
// SortedIDs returns a lexicographically ascending copy of ids, as required
// before aggregation by CES-4.1.3.
func SortedIDs(ids []string) []string {
	out := make([]string, len(ids))
	copy(out, ids)
	sort.Strings(out)
	return out
}
