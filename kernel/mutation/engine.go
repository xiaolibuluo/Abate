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
	"runtime"
	"slices"
	"sync"

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
	// workers is the mutation-judging concurrency. A value below two runs the
	// sequential loop, preserving the original single-goroutine behavior.
	workers int
}

// NewEngine builds an engine bound to one seed.
func NewEngine(seed uint64) *Engine {
	return &Engine{seed: seed, registry: NewRegistry(), workers: 1}
}

// NewEngineParallel builds an engine bound to one seed that judges mutation
// points concurrently across workers goroutines. workers below two keeps the
// sequential loop. The outcome is byte-identical to the sequential engine:
// discovery order, scheduling order and report assembly are unchanged because
// worker results are written back by mutation index, never by completion order.
func NewEngineParallel(seed uint64, workers int) *Engine {
	if workers < 1 {
		workers = 1
	}
	return &Engine{seed: seed, registry: NewRegistry(), workers: workers}
}

// Seed returns the engine seed.
func (e *Engine) Seed() uint64 { return e.seed }

// Registry returns the read-only operator registry of the engine.
func (e *Engine) Registry() *Registry { return e.registry }

// effectiveWorkers resolves the concurrency, clamping to the machine when the
// engine was configured with an unbounded value.
func (e *Engine) effectiveWorkers() int {
	if e.workers < 2 {
		return 1
	}
	if n := runtime.NumCPU(); e.workers > n {
		return n
	}
	return e.workers
}

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
	verdicts, killing, err := e.judgeAll(ctx, target, ordered, tests)
	if err != nil {
		return Report{}, err
	}
	return NewReport(e.seed, target.Fingerprint(), suite.Fingerprint(), ordered, verdicts, killing), nil
}

// judgeAll evaluates every scheduled point, sequentially when the engine is
// single-worker or concurrently across a fixed worker pool. Concurrent results
// are placed back into their scheduled index, so the returned slices keep the
// exact total order of ordered regardless of goroutine completion order.
func (e *Engine) judgeAll(ctx context.Context, target Target,
	ordered []MutationPoint, tests []string,
) ([]verify.MutationVerdict, []string, error) {
	verdicts := make([]verify.MutationVerdict, len(ordered))
	killing := make([]string, len(ordered))
	workers := e.effectiveWorkers()
	if workers < 2 || len(ordered) < 2 {
		for i, point := range ordered {
			if err := ctx.Err(); err != nil {
				return nil, nil, fmt.Errorf("mutation: %w", err)
			}
			verdict, killer, err := e.judgeOne(target, tests, point)
			if err != nil {
				return nil, nil, err
			}
			verdicts[i] = verdict
			killing[i] = killer
		}
		return verdicts, killing, nil
	}
	return e.judgeAllConcurrent(ctx, target, ordered, tests, verdicts, killing, workers)
}

// judgeAllConcurrent is the worker-pool path of judgeAll. Each worker consumes
// scheduled indexes, judges that point, and writes verdict and killing test back
// at the same index. The first error is recorded once and stops further dispatch;
// index-addressed writes keep the partial result deterministic for the caller.
func (e *Engine) judgeAllConcurrent(ctx context.Context, target Target,
	ordered []MutationPoint, tests []string,
	verdicts []verify.MutationVerdict, killing []string, workers int,
) ([]verify.MutationVerdict, []string, error) {
	type job struct{ index int }
	jobs := make(chan job)
	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error
	fail := func(err error) { once.Do(func() { firstErr = err }) }
	worker := func() {
		defer wg.Done()
		for j := range jobs {
			if err := ctx.Err(); err != nil {
				fail(fmt.Errorf("mutation: %w", err))
				continue
			}
			verdict, killer, err := e.judgeOne(target, tests, ordered[j.index])
			if err != nil {
				fail(err)
				continue
			}
			verdicts[j.index] = verdict
			killing[j.index] = killer
		}
	}
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go worker()
	}
	go func() {
		defer close(jobs)
		for i := range ordered {
			select {
			case <-ctx.Done():
				fail(fmt.Errorf("mutation: %w", ctx.Err()))
				return
			case jobs <- job{index: i}:
			}
		}
	}()
	wg.Wait()
	if firstErr != nil {
		return nil, nil, firstErr
	}
	return verdicts, killing, nil
}

// judgeOne evaluates one mutation point and returns its verdict and killing test.
// It is the per-point body shared by the sequential loop and the worker pool.
func (e *Engine) judgeOne(target Target, tests []string,
	point MutationPoint,
) (verify.MutationVerdict, string, error) {
	return Judge(e.registry, target, tests, point)
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
