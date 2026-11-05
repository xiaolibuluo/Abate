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

// NetworkInjector 是网络注入器，通过 tc/iptables 实现丢包、延迟、分区（F3）。
// NetworkInjector performs network chaos via tc/iptables (F3).
type NetworkInjector struct {
	cap    core.Capability
	mu     sync.Mutex
	active map[string]string
}

// tcFn 可在测试中替换以避免真实网络操作。
var tcFn = defaultTC

// defaultTC 使用 tc 命令配置网络规则。
func defaultTC(args ...string) error {
	return exec.Command("tc", args...).Run()
}

// NewNetworkInjector 构造网络注入器。
func NewNetworkInjector() *NetworkInjector {
	return &NetworkInjector{
		cap: core.Capability{
			Name:         "network-injector",
			Version:      "1.0.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F3Partition},
			TargetKinds:  []core.TargetKind{core.TargetInstance, core.TargetService, core.TargetDependency},
		},
		active: make(map[string]string),
	}
}

// Capability 返回能力声明。
func (n *NetworkInjector) Capability() core.Capability { return n.cap }

// Inject 配置网络故障规则（CES-2.4.1 幂等）。
func (n *NetworkInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if _, ok := n.active[e.ID]; ok {
		return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
	}

	iface := e.Fault.Params["interface"]
	if iface == "" {
		iface = "eth0"
	}
	loss := e.Fault.Params["loss"]
	if loss == "" {
		return core.Injection{}, fmt.Errorf("network-injector: loss parameter is required")
	}

	if err := tcFn("qdisc", "add", "dev", iface, "root", "netem", "loss", loss+"%"); err != nil {
		return core.Injection{}, fmt.Errorf("network-injector: tc qdisc add: %w", err)
	}

	n.active[e.ID] = iface
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 删除网络故障规则（CES-3.1.2 可重复调用）。
func (n *NetworkInjector) Recover(_ context.Context, inj core.Injection) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	iface, ok := n.active[inj.ID]
	if !ok {
		return nil
	}

	if err := tcFn("qdisc", "del", "dev", iface, "root"); err != nil {
		return fmt.Errorf("network-injector: tc qdisc del: %w", err)
	}

	delete(n.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (n *NetworkInjector) IsActive() bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	return len(n.active) > 0
}