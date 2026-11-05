// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package sim

import (
	"context"

	"github.com/qingdujian/abate/core"
)

// VirtualValidator 是模拟器的虚拟校验器：纯软件准入与准出检查。
//
// VirtualValidator is the simulator's virtual validator.
type VirtualValidator struct {
	residue bool
}

// Capability 返回能力声明。
func (v *VirtualValidator) Capability() core.Capability {
	return core.Capability{
		Name:     "sim-validator",
		Version:  "0.1",
		Protocol: core.ProtocolVersion,
	}
}

// Preflight 执行准入检查（CES-3.1.4）。
func (v *VirtualValidator) Preflight(_ context.Context, e core.Experiment) ([]string, error) {
	var issues []string
	if e.ID == "" {
		issues = append(issues, "empty experiment ID")
	}
	return issues, nil
}

// Postflight 执行准出检查：恢复验证与残留扫描（CES-3.1.2）。
func (v *VirtualValidator) Postflight(_ context.Context, _ core.Experiment) ([]string, error) {
	var issues []string
	if v.residue {
		issues = append(issues, "injection residue detected")
	}
	return issues, nil
}

// SetResidue 设置残留标志（测试用）。
func (v *VirtualValidator) SetResidue(b bool) {
	v.residue = b
}
