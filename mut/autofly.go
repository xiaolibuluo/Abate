// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package mut

import (
	"fmt"
	"sync"

	"github.com/qingdujian/abate/core"
)

// LLMClient 是大语言模型客户端接口，可由外部实现注入。
// LLMClient is the LLM client interface for test generation.
type LLMClient interface {
	// GenerateTest 根据存活变异体信息生成补测代码。
	GenerateTest(mutant core.Mutation, source string) (testCode string, err error)
}

// TestRunner 是测试执行器接口，用于验证补测是否杀死变异体。
type TestRunner interface {
	// RunTest 执行测试代码，返回是否通过（通过=变异体存活，失败=变异体被杀死）。
	RunTest(source, testCode string) (passed bool, err error)
}

// FlywheelConfig 是自生长飞轮配置。
type FlywheelConfig struct {
	MaxRetries     int
	MaxMutants     int
}

// DefaultFlywheelConfig 返回默认配置。
func DefaultFlywheelConfig() FlywheelConfig {
	return FlywheelConfig{MaxRetries: 3, MaxMutants: 100}
}

// SurvivedMutant 记录一个存活变异体及其补测尝试。
type SurvivedMutant struct {
	Mutant    core.Mutation
	Attempts  int
	Killed    bool
	TestCode  string
}

// Flywheel 是 AI 自生长飞轮：存活变异体→LLM 生成补测→杀死率验证。
// Flywheel is the AI self-growth flywheel.
type Flywheel struct {
	llm    LLMClient
	runner TestRunner
	config FlywheelConfig
	mu     sync.Mutex
	results []SurvivedMutant
}

// NewFlywheel 构造自生长飞轮。
func NewFlywheel(llm LLMClient, runner TestRunner, config FlywheelConfig) (*Flywheel, error) {
	if llm == nil {
		return nil, fmt.Errorf("flywheel: LLM client must not be nil")
	}
	if runner == nil {
		return nil, fmt.Errorf("flywheel: test runner must not be nil")
	}
	return &Flywheel{llm: llm, runner: runner, config: config}, nil
}

// Process 处理存活变异体列表，为每个变异体生成补测并验证。
func (f *Flywheel) Process(mutants []core.Mutation, source string) ([]SurvivedMutant, error) {
	if len(mutants) == 0 {
		return nil, nil
	}
	if f.config.MaxMutants > 0 && len(mutants) > f.config.MaxMutants {
		mutants = mutants[:f.config.MaxMutants]
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	f.results = make([]SurvivedMutant, 0, len(mutants))

	for _, m := range mutants {
		sm := SurvivedMutant{Mutant: m}
		for attempt := 1; attempt <= f.config.MaxRetries; attempt++ {
			sm.Attempts = attempt
			testCode, err := f.llm.GenerateTest(m, source)
			if err != nil {
				continue
			}
			passed, err := f.runner.RunTest(source, testCode)
			if err != nil {
				continue
			}
			if !passed {
				sm.Killed = true
				sm.TestCode = testCode
				break
			}
		}
		f.results = append(f.results, sm)
	}
	return f.results, nil
}

// KillRate 返回杀死率（0-1）。
func (f *Flywheel) KillRate() float64 {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.results) == 0 {
		return 0
	}
	killed := 0
	for _, sm := range f.results {
		if sm.Killed {
			killed++
		}
	}
	return float64(killed) / float64(len(f.results))
}

// Results 返回飞轮处理结果。
func (f *Flywheel) Results() []SurvivedMutant {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]SurvivedMutant, len(f.results))
	copy(out, f.results)
	return out
}

// SurvivedCount 返回未被杀死的变异体数量。
func (f *Flywheel) SurvivedCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	count := 0
	for _, sm := range f.results {
		if !sm.Killed {
			count++
		}
	}
	return count
}
