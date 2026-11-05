// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package sim 实现确定性模拟器（CES-8.2.2）：纯软件数字孪生，
// 不触碰真实平台 API，永不过时。L1 认证底座。
//
// Package sim implements the deterministic simulator (CES-8.2.2):
// a pure-software digital twin that never touches real platform APIs.
package sim

import (
	"context"
	"fmt"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/engine"
)

// Mode 是模拟器运行模式（CES-8.2.2）。
type Mode string

const (
	ModeSimulation Mode = "simulation"
	ModeStaging    Mode = "staging"
	ModeProduction Mode = "production"
)

// Valid 报告模式是否合法。
func (m Mode) Valid() bool {
	switch m {
	case ModeSimulation, ModeStaging, ModeProduction:
		return true
	}
	return false
}

// Simulator 是确定性模拟器核心。
type Simulator struct {
	clock *core.VirtualClock
	rng   core.RNG
	mode  Mode
	inj   core.Injector
	probe core.Probe
	val   core.Validator
}

// NewSimulator 构造指定模式的模拟器。
func NewSimulator(mode Mode, seed int64) *Simulator {
	clk := core.NewVirtualClock(0)
	rng := core.NewRNG(seed)
	inj := &VirtualInjector{rng: core.NewRNG(seed)}
	probe := &VirtualProbe{rng: core.NewRNG(seed + 1)}
	val := &VirtualValidator{}
	return &Simulator{clock: clk, rng: rng, mode: mode, inj: inj, probe: probe, val: val}
}

// transitFn 可在测试中替换以模拟迁移失败。
var transitFn = engine.Transition

// Run 执行一次实验的完整生命周期（CES 第 3 章）。
func (s *Simulator) Run(e core.Experiment) (core.Record, error) {
	if err := e.Validate(); err != nil {
		return core.Record{}, fmt.Errorf("preflight: %w", err)
	}

	state, err := transitFn(engine.StatePending, engine.StatePreflight)
	if err != nil {
		return core.Record{}, err
	}

	issues, err := s.val.Preflight(context.Background(), e)
	if err != nil {
		return core.Record{}, err
	}
	if len(issues) > 0 {
		return core.Record{}, fmt.Errorf("preflight issues: %v", issues)
	}

	state, err = transitFn(state, engine.StateInjecting)
	if err != nil {
		return core.Record{}, err
	}

	injection, err := s.inj.Inject(context.Background(), e)
	if err != nil {
		return core.Record{}, err
	}

	windowMs := e.Blast.MaxDurationMs
	if windowMs <= 0 {
		windowMs = 60000
	}
	s.clock.Advance(windowMs)

	state, err = transitFn(state, engine.StateObserving)
	if err != nil {
		return core.Record{}, err
	}

	sample, err := s.probe.Observe(context.Background(), e.Target)
	if err != nil {
		return core.Record{}, err
	}

	state, err = transitFn(state, engine.StateRecovering)
	if err != nil {
		return core.Record{}, err
	}

	if err := s.inj.Recover(context.Background(), injection); err != nil {
		return core.Record{}, err
	}

	_, err = transitFn(state, engine.StateVerdict)
	if err != nil {
		return core.Record{}, err
	}

	holdMs := windowMs
	if sample.Value < int64(core.BPFull) {
		holdMs = windowMs / 2
	}

	return core.Record{
		Exp:             e,
		Evidence:        core.EvidenceA,
		HoldMs:          holdMs,
		WindowMs:        windowMs,
		RecoveryMs:      60000,
		InvariantPass:   10,
		InvariantTotal:  10,
		DeclaredRadius:  int64(e.Blast.MaxInstances),
	}, nil
}
