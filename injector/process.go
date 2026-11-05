// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package injector

import (
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/qingdujian/abate/core"
)

// ProcessInjector 是进程注入器，通过 kill/pkill 实现进程终止与重启（F2）。
// ProcessInjector performs process kill and restart via system commands (F2).
type ProcessInjector struct {
	cap    core.Capability
	mu     sync.Mutex
	active map[string]string
}

// killFn 可在测试中替换以避免真实进程操作。
var killFn = defaultKill

// defaultKill 使用 pkill 终止进程。
func defaultKill(signal, pattern string) error {
	return exec.Command("pkill", "-"+signal, pattern).Run()
}

// NewProcessInjector 构造进程注入器。
func NewProcessInjector() *ProcessInjector {
	return &ProcessInjector{
		cap: core.Capability{
			Name:         "process-injector",
			Version:      "1.0.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F2Crash},
			TargetKinds:  []core.TargetKind{core.TargetInstance, core.TargetService},
		},
		active: make(map[string]string),
	}
}

// Capability 返回能力声明。
func (p *ProcessInjector) Capability() core.Capability { return p.cap }

// Inject 终止目标进程（CES-2.4.1 幂等）。
func (p *ProcessInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.active[e.ID]; ok {
		return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
	}

	signal := e.Fault.Params["signal"]
	if signal == "" {
		signal = "TERM"
	}
	target := e.Target.Name
	if target == "" {
		return core.Injection{}, fmt.Errorf("process-injector: target name is empty")
	}

	if err := killFn(signal, target); err != nil {
		return core.Injection{}, fmt.Errorf("process-injector: kill %s with %s: %w", target, signal, err)
	}

	p.active[e.ID] = target
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 进程注入无需恢复操作（进程已被终止），幂等返回 nil（CES-3.1.2）。
func (p *ProcessInjector) Recover(_ context.Context, inj core.Injection) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (p *ProcessInjector) IsActive() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.active) > 0
}