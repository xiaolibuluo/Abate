// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package adapter

import (
	"context"
	"fmt"
	"sync"

	"github.com/qingdujian/abate/core"
)

// Adapter 是第三方混沌平台适配器，实现 core.Injector（CES-7.2.3）。
// 适配器必须实现 Injector 角色；Probe/Validator/Reporter 允许由核心兜底。
// Adapter is a third-party platform adapter implementing core.Injector.
type Adapter struct {
	cap     core.Capability
	backend Backend
	table   *MappingTable
	mu      sync.Mutex
	active  map[string]string // experimentID → upstreamID
}

// New 构造适配器。
// New constructs an adapter.
func New(name, version string, backend Backend, table *MappingTable) (*Adapter, error) {
	if backend == nil {
		return nil, fmt.Errorf("adapter: backend must not be nil")
	}
	if table == nil {
		return nil, fmt.Errorf("adapter: mapping table must not be nil")
	}
	return &Adapter{
		cap: core.Capability{
			Name:         name,
			Version:      version,
			Protocol:     core.ProtocolVersion,
			FaultClasses: table.MappedClasses(),
			TargetKinds: []core.TargetKind{
				core.TargetService,
				core.TargetInstance,
				core.TargetDependency,
			},
		},
		backend: backend,
		table:   table,
		active:  make(map[string]string),
	}, nil
}

// Capability 返回能力声明（CES-7.1.2）。
func (a *Adapter) Capability() core.Capability { return a.cap }

// Inject 翻译 CES 故障为上游格式并调用 Backend.Apply（CES-2.4.1 幂等）。
func (a *Adapter) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 幂等：已有活跃注入则直接返回已有句柄。
	if upstreamID, ok := a.active[e.ID]; ok {
		return core.Injection{ID: upstreamID, ExperimentID: e.ID}, nil
	}

	mapped, err := a.table.Translate(e.Fault, e.Target)
	if err != nil {
		return core.Injection{}, fmt.Errorf("adapter: translate fault: %w", err)
	}

	upstreamID, err := a.backend.Apply(mapped)
	if err != nil {
		return core.Injection{}, fmt.Errorf("adapter: backend apply: %w", err)
	}

	a.active[e.ID] = upstreamID
	return core.Injection{ID: upstreamID, ExperimentID: e.ID}, nil
}

// Recover 调用 Backend.Cleanup 清除上游故障（CES-3.1.2 可重复调用）。
func (a *Adapter) Recover(_ context.Context, in core.Injection) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	upstreamID, ok := a.active[in.ExperimentID]
	if !ok {
		// 无活跃注入，幂等返回成功。
		return nil
	}

	if err := a.backend.Cleanup(upstreamID); err != nil {
		return fmt.Errorf("adapter: backend cleanup: %w", err)
	}

	delete(a.active, in.ExperimentID)
	return nil
}

// Degradations 返回适配器的降级记录（CES-7.2.2）。
func (a *Adapter) Degradations() []Degradation {
	return a.table.Degradations()
}