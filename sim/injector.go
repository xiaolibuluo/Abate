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
	"fmt"

	"github.com/qingdujian/abate/core"
)

// VirtualInjector 是模拟器的虚拟注入器：不触碰真实平台，
// 仅记录注入与恢复操作，保证幂等与无残留（CES-2.4.1、CES-3.1.2）。
//
// VirtualInjector is the simulator's virtual injector: never touches real
// platforms, only records injection and recovery operations.
type VirtualInjector struct {
	rng      core.RNG
	active   bool
	injectID string
}

// Capability 返回能力声明。
// Capability returns the capability declaration.
func (v *VirtualInjector) Capability() core.Capability {
	return core.Capability{
		Name:        "sim-injector",
		Version:     "0.1",
		Protocol:    core.ProtocolVersion,
		FaultClasses: core.FaultClasses(),
		TargetKinds:  []core.TargetKind{core.TargetService, core.TargetInstance, core.TargetDependency},
	}
}

// Inject 模拟故障注入：幂等，重复调用不叠加。
// Inject simulates fault injection: idempotent.
func (v *VirtualInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	if v.active {
		return core.Injection{ID: v.injectID, ExperimentID: e.ID}, nil
	}
	v.active = true
	v.injectID = fmt.Sprintf("sim-%s-%d", e.ID, v.rng.Uint64())
	return core.Injection{ID: v.injectID, ExperimentID: e.ID}, nil
}

// Recover 模拟故障恢复：可重复调用，最终无残留。
// Recover simulates fault recovery: repeatable, ultimately residue-free.
func (v *VirtualInjector) Recover(_ context.Context, _ core.Injection) error {
	v.active = false
	return nil
}

// IsActive 报告当前是否有活跃注入。
// IsActive reports whether there is an active injection.
func (v *VirtualInjector) IsActive() bool {
	return v.active
}
