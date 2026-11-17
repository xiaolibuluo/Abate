/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"context"
	"errors"
	"testing"

	"abate/kernel/injection"
	"abate/kernel/mutation"
	"abate/sim/record"
)

// TestRunMixedFramesErrorBranches drives each defensive error branch of the
// orchestration through its injectable ports, proving real error propagation
// instead of swallowing failures.
func TestRunMixedFramesErrorBranches(t *testing.T) {
	t.Parallel()
	sentinel := errMixedTest
	ports := defaultMixedPorts()
	t.Run("setup", func(t *testing.T) {
		broken := ports
		broken.setup = func(uint64) (*injection.Injector, error) {
			return nil, sentinel
		}
		if _, err := runMixedFrames(context.Background(), 42, broken); !errors.Is(err, sentinel) {
			t.Fatalf("setup error = %v, want sentinel", err)
		}
	})
	t.Run("apply", func(t *testing.T) {
		broken := ports
		broken.apply = func(context.Context, *injection.Injector) error {
			return sentinel
		}
		if _, err := runMixedFrames(context.Background(), 42, broken); !errors.Is(err, sentinel) {
			t.Fatalf("apply error = %v, want sentinel", err)
		}
	})
	t.Run("engine", func(t *testing.T) {
		broken := ports
		broken.engine = func(context.Context, uint64) (mutation.Report, error) {
			return mutation.Report{}, sentinel
		}
		if _, err := runMixedFrames(context.Background(), 42, broken); !errors.Is(err, sentinel) {
			t.Fatalf("engine error = %v, want sentinel", err)
		}
	})
}

// stepContext is a context test double whose Err() starts failing at a chosen
// call index, used to drive the real injector's Inject/Rollback error branches.
type stepContext struct {
	context.Context
	calls    int
	failFrom int
}

// Err returns nil before failFrom and context.Canceled from failFrom onward.
func (c *stepContext) Err() error {
	c.calls++
	if c.calls >= c.failFrom {
		return context.Canceled
	}
	return nil
}

// TestDefaultMixedPortsApplyErrors proves the default apply closure propagates
// real Inject and Rollback failures through the genuine injector.
func TestDefaultMixedPortsApplyErrors(t *testing.T) {
	t.Parallel()
	injector, err := defaultMixedPorts().setup(42)
	if err != nil {
		t.Fatalf("setup: %v", err)
	}
	apply := defaultMixedPorts().apply
	if err := apply(&stepContext{Context: context.Background(), failFrom: 1}, injector); !errors.Is(err, context.Canceled) {
		t.Fatalf("inject error = %v, want context.Canceled", err)
	}
	if err := apply(&stepContext{Context: context.Background(), failFrom: 2}, injector); !errors.Is(err, context.Canceled) {
		t.Fatalf("rollback error = %v, want context.Canceled", err)
	}
}

// TestDispatchMixedError proves the handler surfaces a mixed-perturbation
// failure instead of silently recording a partial stream.
func TestDispatchMixedError(t *testing.T) {
	t.Parallel()
	exp := loadFixture(t, "ok.yaml")
	h := newExperimentHandler(exp, probeSchedule(exp), nil)
	h.mixed = func(uint64) ([]record.Frame, error) {
		return nil, errMixedTest
	}
	if _, err := h.handler()(applyEventForTest()); !errors.Is(err, errMixedTest) {
		t.Fatalf("dispatch error = %v, want sentinel", err)
	}
}
