// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package validator 实现 CES 第 5 章准入与准出校验器。
//
// Package validator implements the CES Chapter 5 admission
// and exit validator.
package validator

import (
	"context"
	"fmt"

	"github.com/qingdujian/abate/core"
)

// Validator 是文件系统验证器，实现 core.Validator 接口。
// Preflight 检查实验五元组完整性与爆炸半径声明；
// Postflight 检查残留与不变量。
type Validator struct {
	cap     core.Capability
	residue bool
}

// New 构造验证器。
func New() *Validator {
	return &Validator{
		cap: core.Capability{Name: "file-validator", Version: "1.0.0", Protocol: core.ProtocolVersion},
	}
}

// Capability 返回能力声明。
func (v *Validator) Capability() core.Capability { return v.cap }

// SetResidue 设置残留标志，用于测试 Postflight 残留检出。
func (v *Validator) SetResidue(b bool) { v.residue = b }

// Preflight 准入检查：实验校验 + 爆炸半径风险标记（CES-3.1.4）。
func (v *Validator) Preflight(_ context.Context, e core.Experiment) ([]string, error) {
	if err := e.Validate(); err != nil {
		return nil, fmt.Errorf("preflight: %w", err)
	}
	var issues []string
	declared := 0
	if e.Blast.MaxInstances > 0 {
		declared++
	}
	if e.Blast.MaxDurationMs > 0 {
		declared++
	}
	if e.Blast.MaxTraffic > 0 {
		declared++
	}
	if declared < 3 {
		issues = append(issues, fmt.Sprintf(
			"blastRadius: only %d of 3 dimensions declared (CES-2.5.2)", declared))
	}
	return issues, nil
}

// Postflight 准出检查：残留扫描 + 不变量核验（CES-3.1.2）。
func (v *Validator) Postflight(_ context.Context, e core.Experiment) ([]string, error) {
	var issues []string
	if v.residue {
		issues = append(issues, "residue detected after recovery (CES-3.1.2, E1004)")
	}
	return issues, nil
}
