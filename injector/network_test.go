// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package injector

import (
	"context"
	"errors"
	"testing"

	"github.com/qingdujian/abate/core"
)

func validNetworkExperiment() core.Experiment {
	return core.Experiment{
		ID:          "exp-net",
		Seed:        1,
		Criticality: 3,
		Target:      core.Target{Kind: core.TargetInstance, Name: "node-1"},
		Fault:       core.Fault{Class: core.F3Partition, Params: map[string]string{"loss": "10", "interface": "eth0"}},
		Blast:       core.BlastRadius{MaxInstances: 1, MaxDurationMs: 30000},
		Steady:      []core.SLIAssert{{SLI: "latency", Assert: "< 500ms"}},
	}
}

func TestNetworkInjector_Capability(t *testing.T) {
	n := NewNetworkInjector()
	cap := n.Capability()
	if cap.Name != "network-injector" {
		t.Errorf("name = %q", cap.Name)
	}
}

func TestNetworkInjector_Inject_Success(t *testing.T) {
	old := tcFn
	tcFn = func(args ...string) error { return nil }
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	inj, err := n.Inject(context.Background(), validNetworkExperiment())
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "exp-net" {
		t.Errorf("inj.ID = %q", inj.ID)
	}
	if !n.IsActive() {
		t.Error("should be active")
	}
}

func TestNetworkInjector_Inject_Idempotent(t *testing.T) {
	calls := 0
	old := tcFn
	tcFn = func(args ...string) error { calls++; return nil }
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	exp := validNetworkExperiment()
	_, _ = n.Inject(context.Background(), exp)
	_, err := n.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("second Inject: %v", err)
	}
	if calls != 1 {
		t.Errorf("tc called %d times, want 1", calls)
	}
}

func TestNetworkInjector_Inject_DefaultInterface(t *testing.T) {
	var captured []string
	old := tcFn
	tcFn = func(args ...string) error {
		captured = args
		return nil
	}
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	exp := validNetworkExperiment()
	exp.Fault.Params = map[string]string{"loss": "5"}
	_, err := n.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	found := false
	for _, a := range captured {
		if a == "eth0" {
			found = true
		}
	}
	if !found {
		t.Error("default interface eth0 not found in tc args")
	}
}

func TestNetworkInjector_Inject_MissingLoss(t *testing.T) {
	old := tcFn
	tcFn = func(args ...string) error { return nil }
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	exp := validNetworkExperiment()
	exp.Fault.Params = map[string]string{"interface": "eth0"}
	_, err := n.Inject(context.Background(), exp)
	if err == nil {
		t.Fatal("expected error for missing loss")
	}
}

func TestNetworkInjector_Inject_TCError(t *testing.T) {
	old := tcFn
	tcFn = func(args ...string) error { return errors.New("RTNETLINK: Operation not permitted") }
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	_, err := n.Inject(context.Background(), validNetworkExperiment())
	if err == nil {
		t.Fatal("expected tc error")
	}
}

func TestNetworkInjector_Recover_Success(t *testing.T) {
	old := tcFn
	tcFn = func(args ...string) error { return nil }
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	inj, _ := n.Inject(context.Background(), validNetworkExperiment())
	if err := n.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if n.IsActive() {
		t.Error("should not be active after recover")
	}
}

func TestNetworkInjector_Recover_Idempotent(t *testing.T) {
	n := NewNetworkInjector()
	if err := n.Recover(context.Background(), core.Injection{ID: "none"}); err != nil {
		t.Errorf("Recover with no active: %v", err)
	}
}

func TestNetworkInjector_Recover_TCError(t *testing.T) {
	old := tcFn
	tcFn = func(args ...string) error {
		if len(args) > 0 && args[0] == "qdisc" && len(args) > 1 && args[1] == "del" {
			return errors.New("delete failed")
		}
		return nil
	}
	defer func() { tcFn = old }()

	n := NewNetworkInjector()
	inj, _ := n.Inject(context.Background(), validNetworkExperiment())
	err := n.Recover(context.Background(), inj)
	if err == nil {
		t.Fatal("expected tc delete error")
	}
}

func TestDefaultTC(t *testing.T) {
	err := defaultTC("qdisc", "show")
	if err != nil {
		t.Log("tc returned error (expected if tc not installed)")
	}
}