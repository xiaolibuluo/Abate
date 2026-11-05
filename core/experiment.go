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

// BlastRadius 声明爆炸半径上限（CES-2.5.1：三者至少声明其二）。
// BlastRadius declares blast-radius limits (CES-2.5.1: at least two of three).
type BlastRadius struct {
	// MaxInstances 实例上限；0 表示未声明。
	MaxInstances int `yaml:"maxInstances,omitempty" json:"maxInstances,omitempty"`
	// MaxDurationMs 时长上限（毫秒）；0 表示未声明。
	MaxDurationMs int64 `yaml:"maxDurationMs,omitempty" json:"maxDurationMs,omitempty"`
	// MaxTraffic 流量比例上限（万分位）；0 表示未声明。
	MaxTraffic BP `yaml:"maxTraffic,omitempty" json:"maxTraffic,omitempty"`
}

// Validate 校验爆炸半径：声明维度不少于二（CES-2.5.1）。
// Validate checks that at least two of the three limits are declared.
func (b BlastRadius) Validate() error {
	declared := 0
	if b.MaxInstances > 0 {
		declared++
	}
	if b.MaxDurationMs > 0 {
		declared++
	}
	if b.MaxTraffic > 0 {
		declared++
	}
	if declared < 2 {
		return fmt.Errorf("blastRadius: at least two of maxInstances/maxDurationMs/maxTraffic required (CES-2.5.1)")
	}
	return nil
}

// SLIAssert 是一条稳态断言（CES-2.6.1）：指标源、窗口、阈值、统计方法。
// SLIAssert is one steady-state assertion (CES-2.6.1): metric source, window,
// threshold, and statistical method.
type SLIAssert struct {
	// SLI 指标名（指标源标识）。
	SLI string `yaml:"sli" json:"sli"`
	// Assert 阈值表达式，如 ">= 99.9%"、"< 800ms"。
	Assert string `yaml:"assert" json:"assert"`
	// WindowMs 统计窗口（毫秒）。
	WindowMs int64 `yaml:"windowMs,omitempty" json:"windowMs,omitempty"`
	// Method 统计方法标识（如 mean、p99、rate）。
	Method string `yaml:"method,omitempty" json:"method,omitempty"`
}

// Validate 校验断言字段完整性（CES-2.6.1）。
// Validate checks assertion field completeness (CES-2.6.1).
func (s SLIAssert) Validate() error {
	if s.SLI == "" {
		return fmt.Errorf("sli: must not be empty (CES-2.6.1)")
	}
	if s.Assert == "" {
		return fmt.Errorf("assert: must not be empty (CES-2.6.1)")
	}
	return nil
}

// Experiment 是 CES 第 2 章定义的实验五元组聚合根。
// Experiment is the aggregate root of the CES Chapter 2 five-tuple.
type Experiment struct {
	// ID 实验唯一标识，聚合排序键（CES-4.1.3）。
	ID string `yaml:"id" json:"id"`
	// Seed 确定性重放种子（CES-3.1.3、CES-6.1.1）。
	Seed int64 `yaml:"seed" json:"seed"`
	// Target 作用目标（CES-2.2.1）。
	Target Target `yaml:"target" json:"target"`
	// Fault 注入故障（CES-2.3.1）。
	Fault Fault `yaml:"fault" json:"fault"`
	// Blast 爆炸半径（CES-2.5.1）。
	Blast BlastRadius `yaml:"blastRadius" json:"blastRadius"`
	// Steady 稳态断言集（CES-2.6.1，至少一条）。
	Steady []SLIAssert `yaml:"steadyState" json:"steadyState"`
	// AbortOn 红线集（CES-2.7.1）。
	AbortOn []SLIAssert `yaml:"abortOn" json:"abortOn"`
	// Criticality 目标关键度 1–5（CES-4.3.1）。
	Criticality int `yaml:"criticality" json:"criticality"`
	// RecoveryBudgetMs 恢复预算（CES-4.4.2，缺省 300000ms）。
	RecoveryBudgetMs int64 `yaml:"recoveryBudgetMs,omitempty" json:"recoveryBudgetMs,omitempty"`
}

// Validate 依次校验五元组与关键度（CES 第 2 章强制结构）。
// Validate checks the five-tuple and criticality in order.
func (e Experiment) Validate() error {
	if e.ID == "" {
		return fmt.Errorf("id: must not be empty")
	}
	if err := e.Target.Validate(); err != nil {
		return err
	}
	if err := e.Fault.Validate(); err != nil {
		return err
	}
	if err := e.Blast.Validate(); err != nil {
		return err
	}
	if len(e.Steady) == 0 {
		return fmt.Errorf("steadyState: at least one assertion required (CES-2.6.1)")
	}
	for i, a := range e.Steady {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("steadyState[%d]: %w", i, err)
		}
	}
	for i, a := range e.AbortOn {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("abortOn[%d]: %w", i, err)
		}
	}
	if e.Criticality < 1 || e.Criticality > 5 {
		return fmt.Errorf("criticality: must be 1-5 (CES-4.3.1)")
	}
	return nil
}

// RecoveryBudget 返回恢复预算，未声明时取规范缺省值 300000ms（CES-4.4.2）。
// RecoveryBudget returns the declared budget or the CES-4.4.2 default.
func (e Experiment) RecoveryBudget() int64 {
	if e.RecoveryBudgetMs > 0 {
		return e.RecoveryBudgetMs
	}
	return 300000
}
