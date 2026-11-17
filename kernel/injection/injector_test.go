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
	"testing"

	"abate/kernel/algebra"
	"abate/sim/record"
)

func mustWorld(t *testing.T, seed uint64) *World {
	t.Helper()
	world, err := NewWorld(seed)
	if err != nil {
		t.Fatalf("NewWorld(%d): %v", seed, err)
	}
	return world
}

func mustInjector(t *testing.T, world *World) *Injector {
	t.Helper()
	injector, err := NewInjector(world)
	if err != nil {
		t.Fatalf("NewInjector: %v", err)
	}
	return injector
}

// TestTwelveOperatorsInjectObserveRollback proves all twelve C-family faults
// change real SIM state and that Inverse + baseline restore it exactly.
func TestTwelveOperatorsInjectObserveRollback(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	operators := registry.All()
	if len(operators) != 12 || len(registry.Classes()) != 12 {
		t.Fatalf("registry enumerated %d operators / %d classes, want 12/12",
			len(operators), len(registry.Classes()))
	}
	restored := 0
	for _, op := range operators {
		world := mustWorld(t, 42)
		injector := mustInjector(t, world)
		ctx := context.Background()
		baseline := world.Snapshot()
		if baseline != world.Baseline() {
			t.Fatalf("%s: fresh world differs from baseline", op.ID())
		}
		if err := injector.Inject(ctx, op); err != nil {
			t.Fatalf("%s inject: %v", op.ID(), err)
		}
		if world.Snapshot() == baseline {
			t.Fatalf("%s did not change any observable SIM state", op.ID())
		}
		frames := injector.Frames()
		if len(frames) != 4 {
			t.Fatalf("%s recorded %d frames, want 4", op.ID(), len(frames))
		}
		kinds := make(map[record.FrameKind]bool)
		for _, frame := range frames {
			kinds[frame.Kind] = true
			if len(frame.Payload) == 0 {
				t.Fatalf("%s produced an empty %s frame", op.ID(), frame.Kind)
			}
		}
		for _, kind := range []record.FrameKind{
			record.KindFault, record.KindNet, record.KindFS, record.KindVTask,
		} {
			if !kinds[kind] {
				t.Fatalf("%s produced no %s frame: %v", op.ID(), kind, kinds)
			}
		}
		class := op.Params()["class"]
		if class == "dependency-timeout" &&
			!errors.Is(world.Dependency("svc"), ErrDependencyTimeout) {
			t.Fatal("dependency timeout fault is not observable")
		}
		if class == "dependency-error" &&
			!errors.Is(world.Dependency("svc"), ErrDependencyFailure) {
			t.Fatal("dependency error fault is not observable")
		}
		if err := injector.Rollback(ctx, op); err != nil {
			t.Fatalf("%s rollback: %v", op.ID(), err)
		}
		if world.Snapshot() != baseline {
			t.Fatalf("%s baseline not restored:\n got %s\nwant %s",
				op.ID(), world.Snapshot(), baseline)
		}
		if world.Dependency("svc") != nil {
			t.Fatalf("%s left an active dependency fault", op.ID())
		}
		restored++
	}
	if restored != 12 {
		t.Fatalf("dual compensation restored %d/12", restored)
	}
	t.Logf("dual compensation restored %d/12 = 100%%", restored)
}

func TestRegistryAndWorldAccessors(t *testing.T) {
	t.Parallel()
	world := mustWorld(t, 7)
	injector := mustInjector(t, world)
	if injector.World() != world || injector.Registry() == nil {
		t.Fatal("injector accessors disagree with construction")
	}
	if world.Clock() == nil || world.Network() == nil ||
		world.FileSystem() == nil || world.Scheduler() == nil {
		t.Fatal("world exposes a nil SIM component")
	}
	if len(world.KnownTargets()) != 6 {
		t.Fatalf("known targets = %v", world.KnownTargets())
	}
	if !errors.Is(world.Dependency(""), ErrUnknownTarget) {
		t.Fatal("empty dependency name should be rejected")
	}
	if err := world.Dependency("svc"); err != nil {
		t.Fatalf("clean dependency call failed: %v", err)
	}
	registry := NewRegistry()
	for _, op := range registry.All() {
		if op.Family() != algebra.FamilyChaos || !registry.Contains(op.ID()) {
			t.Fatalf("operator %s is not a registered chaos operator", op.ID())
		}
		if !op.Invertible() {
			t.Fatalf("operator %s lacks a dual", op.ID())
		}
		if _, err := op.Inverse(); err != nil {
			t.Fatalf("operator %s inverse failed: %v", op.ID(), err)
		}
		if op.Params()["compensation"] == "" {
			t.Fatalf("operator %s lacks a compensation strategy", op.ID())
		}
	}
}

func TestInjectorValidationErrors(t *testing.T) {
	t.Parallel()
	world := mustWorld(t, 3)
	injector := mustInjector(t, world)
	valid := injector.Registry().All()[0]
	ctx := context.Background()
	if _, err := NewInjector(nil); !errors.Is(err, ErrNilWorld) {
		t.Fatalf("nil world error = %v", err)
	}
	var missing context.Context
	if err := injector.Inject(missing, valid); !errors.Is(err, ErrCanceled) {
		t.Fatalf("nil context error = %v", err)
	}
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if err := injector.Inject(canceled, valid); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled context error = %v", err)
	}
	if err := injector.Inject(ctx, nil); !errors.Is(err, ErrNilOperator) {
		t.Fatalf("nil operator error = %v", err)
	}
	wrongFamily := algebra.NewBasic(algebra.FormatOperatorID(algebra.FamilyMutation, "x", "1"),
		algebra.FamilyMutation, algebra.Params{"class": "network-latency", "target": "n1-n2"}, nil, true)
	if err := injector.Inject(ctx, wrongFamily); !errors.Is(err, ErrWrongFamily) {
		t.Fatalf("wrong family error = %v", err)
	}
	unknown := algebra.NewBasic(algebra.FormatOperatorID(algebra.FamilyChaos, "unknown", "1"),
		algebra.FamilyChaos, algebra.Params{"class": "network-latency", "target": "n1-n2"}, nil, true)
	if err := injector.Inject(ctx, unknown); !errors.Is(err, ErrUnknownOperator) {
		t.Fatalf("unknown operator error = %v", err)
	}
	noDual := algebra.NewBasic(valid.ID(), algebra.FamilyChaos,
		algebra.Params{"class": "network-latency", "target": "n1-n2", "value": "1"}, nil, false)
	if err := injector.Inject(ctx, noDual); !errors.Is(err, ErrNotInvertible) {
		t.Fatalf("non-invertible error = %v", err)
	}
	badTarget, err := injector.Registry().Variant(valid.ID(),
		algebra.Params{"target": "absent"})
	if err != nil {
		t.Fatalf("Variant: %v", err)
	}
	if err := injector.Inject(ctx, badTarget); !errors.Is(err, ErrUnknownTarget) {
		t.Fatalf("unknown target error = %v", err)
	}
	badClass, err := injector.Registry().Variant(valid.ID(),
		algebra.Params{"class": "not-registered"})
	if err != nil {
		t.Fatalf("Variant: %v", err)
	}
	if err := injector.Inject(ctx, badClass); !errors.Is(err, ErrUnknownOperator) {
		t.Fatalf("unknown class error = %v", err)
	}
}
