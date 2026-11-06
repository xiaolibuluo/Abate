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
	"go/parser"
	"go/token"
	"sync"

	"github.com/qingdujian/abate/core"
)

// LLMClient 是大语言模型客户端接口，可由外部实现注入。
// LLMClient is the LLM client interface for test generation.
type LLMClient interface {
	// GenerateTest 根据存活变异体信息生成补测代码。
	// feedback 携带上一轮尝试的失败原因，供 LLM 针对性修正（CES-9.5.5）。
	GenerateTest(mutant core.Mutation, source, feedback string) (testCode string, err error)
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
		feedback := ""
		for attempt := 1; attempt <= f.config.MaxRetries; attempt++ {
			sm.Attempts = attempt
			testCode, err := f.llm.GenerateTest(m, source, feedback)
			if err != nil {
				feedback = fmt.Sprintf("LLM call failed: %v", err)
				continue
			}
			if err := syntaxCheck(testCode); err != nil {
				feedback = fmt.Sprintf("The generated test has a syntax error: %v", err)
				continue
			}
			passed, err := f.runner.RunTest(source, testCode)
			if err != nil {
				feedback = fmt.Sprintf("Test execution failed: %v", err)
				continue
			}
			if !passed {
				sm.Killed = true
				sm.TestCode = testCode
				break
			}
			feedback = "The test passed against the mutant: the mutant survived and the test did not kill it. Write a stronger test."
		}
		f.results = append(f.results, sm)
	}
	return f.results, nil
}

// syntaxCheck 对 LLM 生成的补测代码做本地语法预检（CES-9.5.5 第一道门禁）。
// syntaxCheck runs a local syntax pre-check on LLM-generated test code (CES-9.5.5).
func syntaxCheck(testCode string) error {
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "generated_test.go", testCode, 0); err != nil {
		return fmt.Errorf("go/parser: %w", err)
	}
	return nil
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
