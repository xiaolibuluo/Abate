// Abate —— 止息混沌 · 实验生命周期阶段
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package sim

import (
	"context"
	"fmt"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/engine"
)

// runPreflight 执行实验配置校验和预检。
func (s *Simulator) runPreflight(logger core.Logger, e core.Experiment) (engine.State, error) {
	logger.Debug("validating experiment config")
	if err := e.Validate(); err != nil {
		logger.Error("experiment validation failed", core.LogField{Key: "err", Value: err.Error()})
		return "", fmt.Errorf("preflight: %w", err)
	}
	logger.Debug("experiment config valid")

	logger.Info("state transition", core.LogField{Key: "from", Value: "Pending"}, core.LogField{Key: "to", Value: "Preflight"})
	state, err := transitFn(engine.StatePending, engine.StatePreflight)
	if err != nil {
		logger.Error("state transition failed", core.LogField{Key: "err", Value: err.Error()})
		return "", err
	}

	logger.Debug("running preflight checks")
	issues, err := s.val.Preflight(context.Background(), e)
	if err != nil {
		logger.Error("preflight error", core.LogField{Key: "err", Value: err.Error()})
		return "", err
	}
	if len(issues) > 0 {
		logger.Warn("preflight issues found", core.LogField{Key: "issues", Value: fmt.Sprintf("%v", issues)}, core.LogField{Key: "count", Value: len(issues)})
		return "", fmt.Errorf("preflight issues: %v", issues)
	}
	logger.Info("preflight passed")
	return state, nil
}

// runInjection 执行故障注入并等待持有窗口结束。
func (s *Simulator) runInjection(logger core.Logger, state engine.State, e core.Experiment) (engine.State, core.Injection, int64, error) {
	logger.Info("state transition", core.LogField{Key: "from", Value: "Preflight"}, core.LogField{Key: "to", Value: "Injecting"})
	state, err := transitFn(state, engine.StateInjecting)
	if err != nil {
		logger.Error("state transition failed", core.LogField{Key: "err", Value: err.Error()})
		return "", core.Injection{}, 0, err
	}

	logger.Info("injecting fault",
		core.LogField{Key: "fault_class", Value: string(e.Fault.Class)},
		core.LogField{Key: "target", Value: e.Target.Name},
	)
	injection, err := s.inj.Inject(context.Background(), e)
	if err != nil {
		logger.Error("injection failed", core.LogField{Key: "err", Value: err.Error()})
		return "", core.Injection{}, 0, err
	}
	logger.Info("injection completed", core.LogField{Key: "injection_id", Value: injection.ID})

	windowMs := e.Blast.MaxDurationMs
	if windowMs <= 0 {
		windowMs = 60000
		logger.Debug("using default window", core.LogField{Key: "window_ms", Value: windowMs})
	}
	s.clock.Advance(windowMs)
	logger.Debug("injection hold window elapsed", core.LogField{Key: "window_ms", Value: windowMs})
	return state, injection, windowMs, nil
}

// runObservation 观测稳态并返回样本。
func (s *Simulator) runObservation(logger core.Logger, state engine.State, e core.Experiment) (engine.State, core.Sample, error) {
	logger.Info("state transition", core.LogField{Key: "from", Value: "Injecting"}, core.LogField{Key: "to", Value: "Observing"})
	state, err := transitFn(state, engine.StateObserving)
	if err != nil {
		logger.Error("state transition failed", core.LogField{Key: "err", Value: err.Error()})
		return "", core.Sample{}, err
	}

	logger.Debug("observing steady state")
	sample, err := s.probe.Observe(context.Background(), e.Target)
	if err != nil {
		logger.Error("observation failed", core.LogField{Key: "err", Value: err.Error()})
		return "", core.Sample{}, err
	}
	logger.Info("observation completed",
		core.LogField{Key: "sli", Value: sample.SLI},
		core.LogField{Key: "value", Value: sample.Value},
	)
	return state, sample, nil
}

// runRecovery 恢复注入的故障。
func (s *Simulator) runRecovery(logger core.Logger, state engine.State, injection core.Injection) (engine.State, error) {
	logger.Info("state transition", core.LogField{Key: "from", Value: "Observing"}, core.LogField{Key: "to", Value: "Recovering"})
	state, err := transitFn(state, engine.StateRecovering)
	if err != nil {
		logger.Error("state transition failed", core.LogField{Key: "err", Value: err.Error()})
		return "", err
	}

	logger.Debug("recovering injected fault", core.LogField{Key: "injection_id", Value: injection.ID})
	if err := s.inj.Recover(context.Background(), injection); err != nil {
		logger.Error("recovery failed", core.LogField{Key: "err", Value: err.Error()})
		return "", err
	}
	logger.Info("recovery completed")

	logger.Info("state transition", core.LogField{Key: "from", Value: "Recovering"}, core.LogField{Key: "to", Value: "Verdict"})
	_, err = transitFn(state, engine.StateVerdict)
	if err != nil {
		logger.Error("state transition failed", core.LogField{Key: "err", Value: err.Error()})
		return "", err
	}
	return state, nil
}
