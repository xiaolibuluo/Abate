/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

	"abate/kernel/injection"
	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/sim/record"
)

// mixedKillTest is the first synthetic test identifier that kills a mutant.
const mixedKillTest = "t2"

// mixedTarget is the synthetic code-space port that the CLI wires into the L6
// mutation engine. Its seven nodes cover all seven M-family operator classes,
// so the derived mutation stream enumerates every registered family.
//
// 参考 PITest 的合成 Target 端口（M2 里程碑先行，替代 M3 的 LNG-AIR），改进点=
// 节点与记号全部固定声明、指纹只遍历有序切片，杜绝 map 迭代序泄漏。
type mixedTarget struct {
	nodes  []string
	tokens map[string]string
}

// newMixedTarget builds the fixed seven-node synthetic target.
func newMixedTarget() *mixedTarget {
	return &mixedTarget{
		nodes: []string{
			"n-arith", "n-boundary", "n-call", "n-constant",
			"n-control", "n-negate", "n-return",
		},
		tokens: map[string]string{
			"n-arith": "+", "n-boundary": "<", "n-call": "call",
			"n-constant": "1", "n-control": "break", "n-negate": "==",
			"n-return": "true",
		},
	}
}

// Fingerprint returns the stable digest of the synthetic code space.
func (t *mixedTarget) Fingerprint() [32]byte {
	var b strings.Builder
	for _, node := range t.nodes {
		_, _ = b.WriteString(node)
		_ = b.WriteByte('=')
		_, _ = b.WriteString(t.tokens[node])
		_ = b.WriteByte(';')
	}
	return sha256.Sum256([]byte(b.String()))
}

// Nodes returns the mutable node references in deterministic order.
func (t *mixedTarget) Nodes() []string {
	return append([]string(nil), t.nodes...)
}

// Token returns the original lexical token of one node.
func (t *mixedTarget) Token(node string) string {
	return t.tokens[node]
}

// Equivalent reports whether a mutant is semantically equivalent.
func (*mixedTarget) Equivalent(point mutation.MutationPoint) bool {
	return point.Mutated == point.Original
}

// Evaluate reports the steady-state outcome of one test on one mutant.
func (*mixedTarget) Evaluate(point mutation.MutationPoint, test string) verify.Outcome {
	if point.Mutated == point.Original || test < mixedKillTest {
		return verify.OutcomeHeld
	}
	return verify.OutcomeViolated
}

// mixedSuite is the synthetic test-selection port wired into the L6 engine.
//
// 参考 PITest 覆盖率引导，改进点=覆盖度用整数计数表达且测试标识固定有序，
// 无浮点权重、无线程优先级、无 map 迭代序泄漏。
type mixedSuite struct{}

// Fingerprint returns the stable digest of the synthetic test suite.
func (mixedSuite) Fingerprint() [32]byte {
	return sha256.Sum256([]byte("abate-m2-cli-suite"))
}

// Tests returns the test identifiers in deterministic order.
func (mixedSuite) Tests() []string {
	return []string{"t1", "t2", "t3"}
}

// Coverage returns the integer number of tests covering one node.
func (mixedSuite) Coverage(_ string) uint64 {
	return 1
}

// mixedPorts groups the L6/L7 entry points of the M⊗C orchestration. Keeping
// them as per-call fields (never package globals) makes every error branch
// unit-testable without weakening the real error propagation.
type mixedPorts struct {
	setup  func(seed uint64) (*injection.Injector, error)
	apply  func(ctx context.Context, injector *injection.Injector) error
	engine func(ctx context.Context, seed uint64) (mutation.Report, error)
}

// defaultMixedPorts binds the orchestration to the real L6 and L7 kernels.
func defaultMixedPorts() mixedPorts {
	return mixedPorts{
		setup: func(seed uint64) (*injection.Injector, error) {
			// NewWorld is infallible for the fixed identifiers; a nil world is
			// still rejected by NewInjector, so no failure is ever swallowed.
			world, _ := injection.NewWorld(seed)
			return injection.NewInjector(world)
		},
		apply: func(ctx context.Context, injector *injection.Injector) error {
			for _, op := range injector.Registry().All() {
				if err := injector.Inject(ctx, op); err != nil {
					return fmt.Errorf("inject %s: %w", op.ID(), err)
				}
				if err := injector.Rollback(ctx, op); err != nil {
					return fmt.Errorf("rollback %s: %w", op.ID(), err)
				}
			}
			return nil
		},
		engine: func(ctx context.Context, seed uint64) (mutation.Report, error) {
			return mutation.NewEngine(seed).Run(ctx, newMixedTarget(), mixedSuite{})
		},
	}
}

// runMixedFrames orchestrates one M⊗C perturbation over the real L6 and L7
// kernels and returns the SIM-REC frames they emit.
//
// 参考 PITest 四条算法线与 chaosblade 故障参数模型的 M2 集成，改进点=把两个内核
// 的事件流统一收敛到 CLI 的单一 recorder，且每帧 payload 均取自真实 SIM 状态或
// Engine 判定，不做任何合成占位。 The C-family frames are the injector's own
// NET/FS/VTASK/FAULT records (virtual link, in-memory FS, virtual tasks and
// virtual clock), while the M-family frames carry each mutant's minimized
// verdict and deterministic point fingerprint derived by mutation.Engine. The
// caller stamps every frame with the scheduler instant before recording, so the
// archive keeps per-kind virtual-time monotonicity.
func runMixedFrames(ctx context.Context, seed uint64,
	ports mixedPorts,
) ([]record.Frame, error) {
	injector, err := ports.setup(seed)
	if err != nil {
		return nil, fmt.Errorf("cli: mixed setup: %w", err)
	}
	if err := ports.apply(ctx, injector); err != nil {
		return nil, fmt.Errorf("cli: mixed apply: %w", err)
	}
	frames := injector.Frames()
	report, err := ports.engine(ctx, seed)
	if err != nil {
		return nil, fmt.Errorf("cli: mixed engine: %w", err)
	}
	for index, verdict := range report.Verdicts {
		point := report.Points[index]
		frames = append(frames, record.Frame{
			Kind: record.KindFault,
			Payload: []byte(fmt.Sprintf("M/%d/%s/%s/%x", index,
				point.Operator.ID(), verdict, point.Fingerprint)),
		})
	}
	return frames, nil
}

// mixedFrames runs the M⊗C orchestration against the real kernels.
func mixedFrames(seed uint64) ([]record.Frame, error) {
	return runMixedFrames(context.Background(), seed, defaultMixedPorts())
}
