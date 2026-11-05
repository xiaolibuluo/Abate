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

// Record 是一次实验执行的完整记录：第 4 章评分算法的输入载体。
// Record is the full execution record of one experiment: the input carrier
// for the Chapter 4 scoring algorithm.
type Record struct {
	// Exp 产生本记录的实验定义。
	Exp Experiment `json:"experiment"`
	// Evidence 本记录的证据等级（CES-5.1）。
	Evidence EvidenceLevel `json:"evidence"`
	// HoldMs 注入窗口内全部稳态 SLI 同时成立的累计毫秒数（CES-4.4.1）。
	HoldMs int64 `json:"holdMs"`
	// WindowMs 注入窗口总毫秒数（CES-4.4.1）。
	WindowMs int64 `json:"windowMs"`
	// RecoveryMs 红线触发到全部 SLI 回落基线并持续 60s 的毫秒数；
	// -1 表示本实验无恢复事件（CES-4.4.2 缺测判定）。
	// RecoveryMs is the recovery duration; -1 means no recovery event
	// occurred (the unmeasured case of CES-4.4.2).
	RecoveryMs int64 `json:"recoveryMs"`
	// Degraded 降级开关是否在注入后 10s 内激活（CES-4.4.3，仅 F3 依赖面实验）。
	Degraded bool `json:"degraded"`
	// DegradedPass 降级模式达标请求数（CES-4.4.3）。
	DegradedPass int64 `json:"degradedPass"`
	// DegradedTotal 降级模式总请求数（CES-4.4.3）。
	DegradedTotal int64 `json:"degradedTotal"`
	// InvariantPass 通过的不变量数（CES-4.4.4）。
	InvariantPass int64 `json:"invariantPass"`
	// InvariantTotal 不变量总数（CES-4.4.4）。
	InvariantTotal int64 `json:"invariantTotal"`
	// DataLoss 是否检出不可调和数据丢失（CES-4.4.4 硬规则）。
	DataLoss bool `json:"dataLoss"`
	// DeclaredRadius 声明爆炸半径（CES-4.4.5）。
	DeclaredRadius int64 `json:"declaredRadius"`
	// ObservedRadius 实际观测半径；0 表示未越界（CES-4.4.5）。
	ObservedRadius int64 `json:"observedRadius"`
	// RedlineLate 红线响应时延是否超过声明上限（CES-4.4.5）。
	RedlineLate bool `json:"redlineLate"`
}

// Weight 计算实验权重 w_e = c_e × e_e（CES-4.3.1）。
// Weight computes w_e = c_e × e_e per CES-4.3.1.
func (r Record) Weight() (BP, error) {
	coef, err := r.Evidence.Coefficient()
	if err != nil {
		return 0, err
	}
	return BP(r.Exp.Criticality) * coef, nil
}

// HoldRatio 计算 hold_e（CES-4.4.1）：稳态成立时长占注入窗口的万分位比率。
// HoldRatio computes hold_e per CES-4.4.1.
func (r Record) HoldRatio() BP {
	return Ratio(r.HoldMs, r.WindowMs)
}

// RecoveryScore 计算 s_e（CES-4.4.2）；无恢复事件返回 0 与 false。
// RecoveryScore computes s_e per CES-4.4.2; the boolean is false when no
// recovery event occurred (unmeasured).
func (r Record) RecoveryScore() (BP, bool) {
	if r.RecoveryMs < 0 {
		return 0, false
	}
	budget := r.Exp.RecoveryBudget()
	s := BPFull - Ratio(r.RecoveryMs, budget)
	return s.Clamp(0, BPFull), true
}

// DegradationScore 计算 DC_e（CES-4.4.3）。
// DegradationScore computes DC_e per CES-4.4.3.
func (r Record) DegradationScore() BP {
	var f BP
	if r.Degraded {
		f = BPFull
	}
	q := Ratio(r.DegradedPass, r.DegradedTotal)
	return f.MulDiv(q, BPFull)
}

// IntegrityScore 计算 d_e（CES-4.4.4）。
// IntegrityScore computes d_e per CES-4.4.4.
func (r Record) IntegrityScore() BP {
	return Ratio(r.InvariantPass, r.InvariantTotal)
}

// RadiusScore 计算 r_e（CES-4.4.5）：红线超时直接归零；观测半径为 0 取满分。
// RadiusScore computes r_e per CES-4.4.5.
func (r Record) RadiusScore() BP {
	if r.RedlineLate {
		return 0
	}
	if r.ObservedRadius == 0 {
		return BPFull
	}
	ratio := Ratio(r.DeclaredRadius, r.ObservedRadius)
	return ratio.Min(BPFull)
}
