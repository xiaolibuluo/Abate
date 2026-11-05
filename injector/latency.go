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
	"sync"

	"github.com/qingdujian/abate/core"
)

// LatencyInjector 是延迟注入器，通过 tc netem delay 实现网络延迟（F1）。
// LatencyInjector performs network latency injection via tc netem delay (F1).
type LatencyInjector struct {
	cap    core.Capability
	mu     sync.Mutex
	active map[string]string
}

// NewLatencyInjector 构造延迟注入器。
func NewLatencyInjector() *LatencyInjector {
	return &LatencyInjector{
		cap: core.Capability{
			Name:         "latency-injector",
			Version:      "1.0.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F1Latency},
			TargetKinds:  []core.TargetKind{core.TargetInstance, core.TargetService, core.TargetDependency},
		},
		active: make(map[string]string),
	}
}

// Capability 返回能力声明。
func (l *LatencyInjector) Capability() core.Capability { return l.cap }

// Inject 配置网络延迟规则（CES-2.4.1 幂等）。
func (l *LatencyInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.active[e.ID]; ok {
		return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
	}

	iface := e.Fault.Params["interface"]
	if iface == "" {
		iface = "eth0"
	}
	delay := e.Fault.Params["delay"]
	if delay == "" {
		return core.Injection{}, fmt.Errorf("latency-injector: delay parameter is required")
	}

	jitter := e.Fault.Params["jitter"]
	args := []string{"qdisc", "add", "dev", iface, "root", "netem", "delay", delay}
	if jitter != "" {
		args = append(args, jitter)
	}

	if err := tcFn(args...); err != nil {
		return core.Injection{}, fmt.Errorf("latency-injector: tc qdisc add: %w", err)
	}

	l.active[e.ID] = iface
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 删除网络延迟规则（CES-3.1.2 可重复调用）。
func (l *LatencyInjector) Recover(_ context.Context, inj core.Injection) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	iface, ok := l.active[inj.ID]
	if !ok {
		return nil
	}

	if err := tcFn("qdisc", "del", "dev", iface, "root"); err != nil {
		return fmt.Errorf("latency-injector: tc qdisc del: %w", err)
	}

	delete(l.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (l *LatencyInjector) IsActive() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.active) > 0
}
