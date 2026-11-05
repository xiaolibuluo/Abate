// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package probe 实现 CES 第 7 章 Probe 角色的文件系统适配器。
//
// Package probe implements the CES Chapter 7 Probe role
// as a file-system adapter.
package probe

import (
	"context"
	"fmt"

	"github.com/qingdujian/abate/core"
)

// Probe 是文件系统探针，实现 core.Probe 接口。
// 使用确定性 RNG 生成可重放的 SLI 样本（CES-5.3.2）。
type Probe struct {
	cap core.Capability
	rng core.RNG
}

// New 构造指定种子的探针。
func New(seed int64) *Probe {
	return &Probe{
		cap: core.Capability{
			Name:        "file-probe",
			Version:     "1.0.0",
			Protocol:    core.ProtocolVersion,
			TargetKinds: []core.TargetKind{core.TargetService, core.TargetInstance, core.TargetDependency},
		},
		rng: core.NewRNG(seed),
	}
}

// Capability 返回能力声明。
func (p *Probe) Capability() core.Capability { return p.cap }

// Observe 对目标执行一次确定性采样（CES-5.3.2）。
func (p *Probe) Observe(_ context.Context, t core.Target) (core.Sample, error) {
	if t.Name == "" {
		return core.Sample{}, fmt.Errorf("probe: target name is empty")
	}
	return core.Sample{
		SLI:         t.Name,
		Value:       int64(p.rng.Intn(10000)),
		TimestampMs: int64(p.rng.Intn(1000000)),
	}, nil
}
