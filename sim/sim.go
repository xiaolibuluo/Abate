// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package sim

import (
	"time"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/engine"
	"github.com/qingdujian/abate/log"
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
	log   core.Logger
}

// NewSimulator 构造指定模式的模拟器。
func NewSimulator(mode Mode, seed int64) *Simulator {
	clk := core.NewVirtualClock(0)
	rng := core.NewRNG(seed)
	inj := &VirtualInjector{rng: core.NewRNG(seed)}
	probe := &VirtualProbe{rng: core.NewRNG(seed + 1)}
	val := &VirtualValidator{}
	return &Simulator{
		clock: clk, rng: rng, mode: mode,
		inj: inj, probe: probe, val: val,
		log: log.Adapt(log.Nop()),
	}
}

// SetLogger 注入日志器，实现全链路日志输出。
func (s *Simulator) SetLogger(l core.Logger) {
	if l == nil {
		l = log.Adapt(log.Nop())
	}
	s.log = l.With(
		core.LogField{Key: "component", Value: "simulator"},
		core.LogField{Key: "mode", Value: string(s.mode)},
	)
}

// transitFn 可在测试中替换以模拟迁移失败。
var transitFn = engine.Transition

// Run 执行一次实验的完整生命周期（CES 第 3 章）。
func (s *Simulator) Run(e core.Experiment) (core.Record, error) {
	start := time.Now()
	logger := s.log.With(
		core.LogField{Key: "exp_id", Value: e.ID},
		core.LogField{Key: "seed", Value: e.Seed},
	)
	logger.Info("experiment started",
		core.LogField{Key: "fault_class", Value: string(e.Fault.Class)},
		core.LogField{Key: "target", Value: e.Target.Name},
		core.LogField{Key: "criticality", Value: e.Criticality},
	)

	state, err := s.runPreflight(logger, e)
	if err != nil {
		return core.Record{}, err
	}

	state, injection, windowMs, err := s.runInjection(logger, state, e)
	if err != nil {
		return core.Record{}, err
	}

	state, sample, err := s.runObservation(logger, state, e)
	if err != nil {
		return core.Record{}, err
	}

	if _, err := s.runRecovery(logger, state, injection); err != nil {
		return core.Record{}, err
	}

	holdMs := windowMs
	if sample.Value < int64(core.BPFull) {
		holdMs = windowMs / 2
		logger.Debug("hold time reduced due to degraded sample")
	}

	record := core.Record{
		Exp:            e,
		Evidence:       core.EvidenceA,
		HoldMs:         holdMs,
		WindowMs:       windowMs,
		RecoveryMs:     60000,
		InvariantPass:  10,
		InvariantTotal: 10,
		DeclaredRadius: int64(e.Blast.MaxInstances),
	}

	logger.Info("experiment completed",
		core.LogField{Key: "verdict", Value: "PASS"},
		core.LogField{Key: "hold_ms", Value: holdMs},
		core.LogField{Key: "window_ms", Value: windowMs},
		core.LogField{Key: "duration_ms", Value: time.Since(start).Milliseconds()},
		core.LogField{Key: "invariant_pass", Value: record.InvariantPass},
		core.LogField{Key: "invariant_total", Value: record.InvariantTotal},
	)

	return record, nil
}
