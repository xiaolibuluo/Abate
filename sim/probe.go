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

// VirtualProbe 是模拟器的虚拟采样器：不连接真实指标源，
// 由确定性 RNG 驱动产出 SLI 样本（CES-5.3.2）。
//
// VirtualProbe is the simulator's virtual probe: never connects to real
// metric sources; SLI samples are driven by a deterministic RNG.
type VirtualProbe struct {
	rng core.RNG
}

// Capability 返回能力声明。
func (p *VirtualProbe) Capability() core.Capability {
	return core.Capability{
		Name:     "sim-probe",
		Version:  "0.1",
		Protocol: core.ProtocolVersion,
	}
}

// Observe 产出一次确定性 SLI 样本。
func (p *VirtualProbe) Observe(_ context.Context, t core.Target) (core.Sample, error) {
	val := int64(p.rng.Uint64() % uint64(core.BPFull))
	return core.Sample{
		SLI:         t.Name,
		Value:       val,
		TimestampMs: 0,
	}, nil
}
