/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"abate/kernel/algebra"
	"abate/sim/clock"
	"abate/sim/fs"
	"abate/sim/net"
	"abate/sim/vsched"
)

func rawOperator(class, target, value string) algebra.Operator {
	return algebra.NewBasic(formatID("network-latency"), algebra.FamilyChaos,
		algebra.Params{"class": class, "target": target, "value": value}, nil, true)
}

// fakeDualOperator wraps a real operator but returns a chosen Inverse result,
// so the rollback dual-validation branches stay reachable.
type fakeDualOperator struct {
	algebra.Operator
	dual algebra.Operator
}

func (fakeDualOperator) Invertible() bool { return true }

func (f fakeDualOperator) Inverse() (algebra.Operator, error) {
	if f.dual == nil {
		return nil, algebra.ErrNotInvertible
	}
	return f.dual, nil
}

func TestParameterVariantsDetected(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	for _, op := range registry.All() {
		original, err := strconv.ParseInt(op.Params()["value"], 10, 64)
		if err != nil {
			t.Fatalf("%s has a malformed value: %v", op.ID(), err)
		}
		variant, err := registry.Variant(op.ID(), algebra.Params{
			"value": strconv.FormatInt(original+13, 10),
		})
		if err != nil {
			t.Fatalf("%s variant: %v", op.ID(), err)
		}
		if variant.Fingerprint() == op.Fingerprint() {
			t.Fatalf("%s parameter mutation was not detected", op.ID())
		}
		world := mustWorld(t, 17)
		injector := mustInjector(t, world)
		ctx := context.Background()
		baseline := world.Snapshot()
		if err := injector.Inject(ctx, variant); err != nil {
			t.Fatalf("%s variant inject: %v", op.ID(), err)
		}
		if err := injector.Rollback(ctx, variant); err != nil {
			t.Fatalf("%s variant rollback: %v", op.ID(), err)
		}
		if world.Snapshot() != baseline {
			t.Fatalf("%s variant did not restore the baseline", op.ID())
		}
	}
}

func TestSuspendModeVariant(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	variant, err := registry.Variant(formatID("process-kill"),
		algebra.Params{"mode": "suspend"})
	if err != nil {
		t.Fatalf("Variant: %v", err)
	}
	world := mustWorld(t, 4)
	injector := mustInjector(t, world)
	ctx := context.Background()
	baseline := world.Snapshot()
	if err := injector.Inject(ctx, variant); err != nil {
		t.Fatalf("suspend inject: %v", err)
	}
	task, ok := world.Scheduler().Task("task-1")
	if !ok || task.State() != 3 {
		t.Fatalf("suspend left task state %d", task.State())
	}
	if err := injector.Rollback(ctx, variant); err != nil {
		t.Fatalf("suspend rollback: %v", err)
	}
	if world.Snapshot() != baseline {
		t.Fatal("suspend mode did not restore the baseline")
	}
}

func TestMalformedVariantAndUnknownID(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	op := registry.All()[0]
	malformed, err := registry.Variant(op.ID(), algebra.Params{"value": "not-an-int"})
	if err != nil {
		t.Fatalf("Variant: %v", err)
	}
	if err := mustInjector(t, mustWorld(t, 1)).Inject(context.Background(), malformed); !errors.Is(err, ErrInvalidParams) {
		t.Fatalf("malformed params error = %v, want ErrInvalidParams", err)
	}
	unknownID := algebra.OperatorID("c-absent-v1")
	if _, err := registry.Variant(unknownID, algebra.Params{}); !errors.Is(err, ErrUnknownOperator) {
		t.Fatalf("unknown variant id error = %v", err)
	}
	if registry.Contains(unknownID) {
		t.Fatal("registry accepted an unknown identifier")
	}
	if _, ok := registry.specByClass("absent"); ok {
		t.Fatal("specByClass accepted an unknown class")
	}
	if _, err := paramsOf(rawOperator("x", "y", "bad")); !errors.Is(err, ErrInvalidParams) {
		t.Fatalf("paramsOf error = %v", err)
	}
}

func TestRollbackDualValidation(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	op := NewRegistry().All()[0]
	world := mustWorld(t, 6)
	injector := mustInjector(t, world)
	notDual := fakeDualOperator{Operator: op, dual: op}
	if err := injector.Rollback(ctx, notDual); !errors.Is(err, ErrNotDual) {
		t.Fatalf("non-dual rollback error = %v, want ErrNotDual", err)
	}
	noInverse := fakeDualOperator{Operator: op}
	if err := injector.Rollback(ctx, noInverse); !errors.Is(err, algebra.ErrNotInvertible) {
		t.Fatalf("inverse failure error = %v", err)
	}
	if err := injector.Inject(ctx, op); err != nil {
		t.Fatalf("inject: %v", err)
	}
	world.baseline = "tampered"
	if err := injector.Rollback(ctx, op); !errors.Is(err, ErrRollbackIncomplete) {
		t.Fatalf("incomplete rollback error = %v, want ErrRollbackIncomplete", err)
	}
}

func TestClockOffsetDualState(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	op, err := registry.Variant(formatID("clock-offset"), algebra.Params{})
	if err != nil {
		t.Fatalf("Variant: %v", err)
	}
	world := mustWorld(t, 9)
	injector := mustInjector(t, world)
	ctx := context.Background()
	baseline := world.Snapshot()
	if world.Clock().State() != clock.StateRunning {
		t.Fatal("fresh world clock is not running")
	}
	if err := injector.Inject(ctx, op); err != nil {
		t.Fatalf("clock inject: %v", err)
	}
	if world.Clock().State() != clock.StateOffset {
		t.Fatalf("clock state = %s, want OFFSET", world.Clock().State())
	}
	if err := injector.Rollback(ctx, op); err != nil {
		t.Fatalf("clock rollback: %v", err)
	}
	if world.Clock().State() != clock.StateRunning || world.Snapshot() != baseline {
		t.Fatal("clock offset dual did not restore the baseline state")
	}
}

func TestDependencyFaultPrecedence(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	world := mustWorld(t, 10)
	injector := mustInjector(t, world)
	ctx := context.Background()
	for _, name := range []string{"dependency-error", "dependency-timeout"} {
		op, err := registry.Variant(formatID(name), algebra.Params{})
		if err != nil {
			t.Fatalf("Variant %s: %v", name, err)
		}
		if err := injector.Inject(ctx, op); err != nil {
			t.Fatalf("inject %s: %v", name, err)
		}
	}
	if !errors.Is(world.Dependency("svc"), ErrDependencyTimeout) {
		t.Fatal("timeout should take precedence over the error fault")
	}
	world.RestoreBaseline()
	if world.Snapshot() != world.Baseline() {
		t.Fatal("RestoreBaseline did not restore a clean world")
	}
}

func TestDispatchAndPrimitiveErrorPaths(t *testing.T) {
	t.Parallel()
	world := &World{
		clk: clock.NewVirtualClock(), network: net.NewNetworkWithSeed(clock.NewVirtualClock(), 1),
		filesystem: fs.NewFileSystem(clock.NewVirtualClock()), scheduler: vsched.NewScheduler(clock.NewVirtualClock()),
		quotaBytes: fs.MaxTotalBytes,
	}
	world.link = world.network.Link("x", "y", net.LinkConfig{})
	bogus := rawOperator("bogus", "n1-n2", "1")
	malformed := rawOperator("network-latency", "n1-n2", "bad")
	checks := []struct {
		name string
		err  error
	}{
		{"apply-dispatch", dispatchApply(world, bogus)},
		{"recover-dispatch", dispatchRecover(world, bogus)},
		{"apply-network-class", applyNetwork(world, bogus)},
		{"recover-network-class", recoverNetwork(world, bogus)},
		{"apply-network-params", applyNetwork(world, malformed)},
		{"recover-network-params", recoverNetwork(world, malformed)},
		{"apply-network-sim", applyNetwork(world, rawOperator("network-latency", "n1-n2", "1"))},
		{"recover-network-sim", recoverNetwork(world, rawOperator("network-latency", "n1-n2", "1"))},
		{"apply-resource-class", applyResource(world, bogus)},
		{"recover-resource-class", recoverResource(world, bogus)},
		{"apply-resource-params", applyResource(world, malformed)},
		{"recover-resource-params", recoverResource(world, malformed)},
		{"apply-resource-oom", applyResource(world, rawOperator("memory-pressure", "task-1", "9999"))},
		{"apply-clock-params", applyClock(world, malformed)},
		{"recover-clock-params", recoverClock(world, malformed)},
		{"apply-process-class", applyProcess(world, bogus)},
		{"recover-process-class", recoverProcess(world, bogus)},
		{"apply-process-params", applyProcess(world, malformed)},
		{"recover-process-params", recoverProcess(world, malformed)},
		{"apply-process-sim", applyProcess(world, rawOperator("process-kill", "task-1", "0"))},
		{"recover-process-sim", recoverProcess(world, rawOperator("process-kill", "task-1", "0"))},
		{"apply-dependency-class", applyDependency(world, bogus)},
		{"recover-dependency-class", recoverDependency(world, bogus)},
		{"apply-dependency-params", applyDependency(world, malformed)},
		{"recover-dependency-params", recoverDependency(world, malformed)},
	}
	for _, check := range checks {
		if check.err == nil {
			t.Errorf("%s accepted an invalid primitive call", check.name)
		}
	}
	suspendOp := algebra.NewBasic(formatID("process-kill"), algebra.FamilyChaos,
		algebra.Params{"class": "process-kill", "target": "task-1", "value": "0", "mode": "suspend"}, nil, true)
	if err := recoverProcess(world, suspendOp); err == nil {
		t.Error("suspend-mode recovery accepted a missing task")
	}
	if state := world.fsState(); state == "" {
		t.Error("fsState returned an empty snapshot for a missing file")
	}
}
