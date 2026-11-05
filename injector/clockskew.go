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
	"strconv"
	"sync"
	"time"

	"github.com/qingdujian/abate/core"
)

var dateCmdFn = func(newTime string) error {
	return exec.Command("date", "-s", newTime).Run()
}

var defaultSetClock = func(offset time.Duration) error {
	newTime := time.Now().Add(offset).Format("2006-01-02 15:04:05")
	return dateCmdFn(newTime)
}

type clockState struct {
	offset time.Duration
}

// ClockSkewInjector 是时钟偏移注入器，通过 date -s 实现系统时钟偏移（F6）。
// ClockSkewInjector performs clock skew via date -s (F6).
type ClockSkewInjector struct {
	cap      core.Capability
	mu       sync.Mutex
	active   map[string]*clockState
	setClock func(time.Duration) error
}

// NewClockSkewInjector 构造时钟偏移注入器。
func NewClockSkewInjector() *ClockSkewInjector {
	return &ClockSkewInjector{
		cap: core.Capability{
			Name:         "clockskew-injector",
			Version:      "1.0.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F6ClockSkew},
			TargetKinds:  []core.TargetKind{core.TargetInstance, core.TargetService},
		},
		active:   make(map[string]*clockState),
		setClock: defaultSetClock,
	}
}

// Capability 返回能力声明。
func (c *ClockSkewInjector) Capability() core.Capability { return c.cap }

// Inject 偏移系统时钟（CES-2.4.1 幂等）。
func (c *ClockSkewInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.active[e.ID]; ok {
		return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
	}

	offsetStr := e.Fault.Params["offset"]
	if offsetStr == "" {
		return core.Injection{}, fmt.Errorf("clockskew-injector: offset parameter is required")
	}

	offsetMs, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return core.Injection{}, fmt.Errorf("clockskew-injector: invalid offset %q: %w", offsetStr, err)
	}

	offset := time.Duration(offsetMs) * time.Millisecond
	if err := c.setClock(offset); err != nil {
		return core.Injection{}, fmt.Errorf("clockskew-injector: set clock: %w", err)
	}

	c.active[e.ID] = &clockState{offset: offset}
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 反向偏移恢复时钟（CES-3.1.2 可重复调用）。
func (c *ClockSkewInjector) Recover(_ context.Context, inj core.Injection) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	state, ok := c.active[inj.ID]
	if !ok {
		return nil
	}

	if err := c.setClock(-state.offset); err != nil {
		return fmt.Errorf("clockskew-injector: restore clock: %w", err)
	}

	delete(c.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (c *ClockSkewInjector) IsActive() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.active) > 0
}
