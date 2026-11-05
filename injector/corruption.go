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
	"os"
	"strconv"
	"sync"

	"github.com/qingdujian/abate/core"
)

var defaultCorrupt = func(path string, offset int64, bit int) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if offset < 0 || offset >= int64(len(data)) {
		return nil, fmt.Errorf("offset %d out of range (size %d)", offset, len(data))
	}
	backup := make([]byte, len(data))
	copy(backup, data)
	data[offset] ^= 1 << bit
	return backup, os.WriteFile(path, data, 0644)
}

var defaultRestore = func(path string, backup []byte) error {
	return os.WriteFile(path, backup, 0644)
}

type corruptionState struct {
	path   string
	backup []byte
}

// CorruptionInjector 是数据损坏注入器，通过文件位翻转实现数据损坏（F5）。
// CorruptionInjector performs data corruption via file bit-flip (F5).
type CorruptionInjector struct {
	cap       core.Capability
	mu        sync.Mutex
	active    map[string]*corruptionState
	corruptFn func(path string, offset int64, bit int) ([]byte, error)
	restoreFn func(path string, backup []byte) error
}

// NewCorruptionInjector 构造数据损坏注入器。
func NewCorruptionInjector() *CorruptionInjector {
	return &CorruptionInjector{
		cap: core.Capability{
			Name:         "corruption-injector",
			Version:      "1.0.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F5Corruption},
			TargetKinds:  []core.TargetKind{core.TargetInstance, core.TargetService},
		},
		active:    make(map[string]*corruptionState),
		corruptFn: defaultCorrupt,
		restoreFn: defaultRestore,
	}
}

// Capability 返回能力声明。
func (c *CorruptionInjector) Capability() core.Capability { return c.cap }

// Inject 执行文件位翻转（CES-2.4.1 幂等）。
func (c *CorruptionInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.active[e.ID]; ok {
		return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
	}

	path := e.Fault.Params["file"]
	if path == "" {
		return core.Injection{}, fmt.Errorf("corruption-injector: file parameter is required")
	}

	offsetStr := e.Fault.Params["offset"]
	if offsetStr == "" {
		offsetStr = "0"
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return core.Injection{}, fmt.Errorf("corruption-injector: invalid offset %q: %w", offsetStr, err)
	}

	bitStr := e.Fault.Params["bit"]
	if bitStr == "" {
		bitStr = "0"
	}
	bit, err := strconv.Atoi(bitStr)
	if err != nil {
		return core.Injection{}, fmt.Errorf("corruption-injector: invalid bit %q: %w", bitStr, err)
	}

	backup, err := c.corruptFn(path, offset, bit)
	if err != nil {
		return core.Injection{}, fmt.Errorf("corruption-injector: %w", err)
	}

	c.active[e.ID] = &corruptionState{path: path, backup: backup}
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

// Recover 从备份恢复文件（CES-3.1.2 可重复调用）。
func (c *CorruptionInjector) Recover(_ context.Context, inj core.Injection) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	state, ok := c.active[inj.ID]
	if !ok {
		return nil
	}

	if err := c.restoreFn(state.path, state.backup); err != nil {
		return fmt.Errorf("corruption-injector: restore: %w", err)
	}

	delete(c.active, inj.ID)
	return nil
}

// IsActive 报告是否有活跃注入。
func (c *CorruptionInjector) IsActive() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.active) > 0
}
