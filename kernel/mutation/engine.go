/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"context"
	"fmt"
	"slices"

	"abate/kernel/verify"
	"abate/sim/rand"
)

// Engine orchestrates one deterministic mutation experiment.
//
// 参考 PITest 的 Engine/Run 编排（发现 → 调度 → 判定 → 报告），改进点=每次 Run
// 从种子新建独立的 SIM-RNG 流、全流程全序化，因而可并发执行且同种子逐字节一致。
// The registry is immutable, so a single Engine is safe for concurrent runs.
type Engine struct {
	seed     uint64
	registry *Registry
}

// NewEngine builds an engine bound to one seed.
func NewEngine(seed uint64) *Engine {
	return &Engine{seed: seed, registry: NewRegistry()}
}

// Seed returns the engine seed.
func (e *Engine) Seed() uint64 { return e.seed }

// Registry returns the read-only operator registry of the engine.
func (e *Engine) Registry() *Registry { return e.registry }

// Run executes the mutation experiment and returns its deterministic report.
//
// The verdict of each mutant is delegated to the KRN-VER verifier: steady-state
// outcomes from the synthetic Target are folded into verify.Classify, and the
// score into verify.Summarize, so no verdict logic is duplicated in this kernel.
func (e *Engine) Run(ctx context.Context, target Target, suite Suite) (Report, error) {
	if ctx == nil {
		return Report{}, ErrCanceled
	}
	if err := ctx.Err(); err != nil {
		return Report{}, fmt.Errorf("mutation: %w", err)
	}
	stream := rand.New(e.seed)
	points, err := Discover(target, suite, e.registry, stream)
	if err != nil {
		return Report{}, err
	}
	ordered := Schedule(points)
	tests := append([]string(nil), suite.Tests()...)
	slices.Sort(tests)
	verdicts := make([]verify.MutationVerdict, 0, len(ordered))
	killing := make([]string, 0, len(ordered))
	used := make([]MutationPoint, 0, len(ordered))
	for _, point := range ordered {
		if err := ctx.Err(); err != nil {
			return Report{}, fmt.Errorf("mutation: %w", err)
		}
		verdict, killer, err := Judge(e.registry, target, tests, point)
		if err != nil {
			return Report{}, err
		}
		verdicts = append(verdicts, verdict)
		killing = append(killing, killer)
		used = append(used, point)
	}
	return NewReport(e.seed, target.Fingerprint(), suite.Fingerprint(), used, verdicts, killing), nil
}

// Judge evaluates one mutation point and returns its verdict and killing test.
//
// 参考 PITest 的 mutant 判定，改进点=判定完全委托 KRN-VER（Classify），击杀测试由
// 确定性二分最小化定位，且未注册算子以 sentinel error 显式拒绝。
func Judge(reg *Registry, target Target, tests []string,
	point MutationPoint,
) (verify.MutationVerdict, string, error) {
	if reg == nil {
		return "", "", ErrNilRegistry
	}
	if !reg.Valid(point.Operator) {
		return "", "", ErrUnknownOperator
	}
	if target.Equivalent(point) {
		return verify.MutantEquivalent, "", nil
	}
	observations := make([]verify.Verdict, 0, len(tests))
	for _, test := range tests {
		observations = append(observations, verify.Verdict{
			ProbeID: test, Outcome: target.Evaluate(point, test),
		})
	}
	index, err := Minimize(point, tests, func(p MutationPoint, test string) bool {
		return target.Evaluate(p, test) == verify.OutcomeViolated
	})
	if err != nil {
		return "", "", err
	}
	killer := ""
	if index >= 0 {
		killer = tests[index]
	}
	return verify.Classify(observations), killer, nil
}
