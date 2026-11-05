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

func TestLatencyInjector_Capability(t *testing.T) {
	l := NewLatencyInjector()
	cap := l.Capability()
	if cap.Name != "latency-injector" {
		t.Errorf("name = %s", cap.Name)
	}
	if len(cap.FaultClasses) != 1 || cap.FaultClasses[0] != core.F1Latency {
		t.Errorf("fault classes = %v", cap.FaultClasses)
	}
	if l.IsActive() {
		t.Error("IsActive should be false on new injector")
	}
}

func TestLatencyInjector_Inject_Success(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	var called []string
	tcFn = func(args ...string) error {
		called = args
		return nil
	}

	l := NewLatencyInjector()
	exp := core.Experiment{ID: "lat1", Fault: core.Fault{Params: map[string]string{
		"interface": "eth1",
		"delay":     "100ms",
		"jitter":    "10ms",
	}}}

	inj, err := l.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "lat1" {
		t.Errorf("injection ID = %s", inj.ID)
	}
	if !l.IsActive() {
		t.Error("IsActive should be true after inject")
	}
	if called[0] != "qdisc" || called[3] != "eth1" || called[6] != "delay" || called[7] != "100ms" || called[8] != "10ms" {
		t.Errorf("tc args = %v", called)
	}
}

func TestLatencyInjector_Inject_DefaultInterface(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	var called []string
	tcFn = func(args ...string) error {
		called = args
		return nil
	}

	l := NewLatencyInjector()
	exp := core.Experiment{ID: "lat2", Fault: core.Fault{Params: map[string]string{
		"delay": "50ms",
	}}}

	_, err := l.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if called[3] != "eth0" {
		t.Errorf("default interface = %s, want eth0", called[3])
	}
	if len(called) != 8 || called[7] != "50ms" {
		t.Errorf("tc args without jitter = %v", called)
	}
}

func TestLatencyInjector_Inject_Idempotent(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	callCount := 0
	tcFn = func(args ...string) error {
		callCount++
		return nil
	}

	l := NewLatencyInjector()
	exp := core.Experiment{ID: "lat3", Fault: core.Fault{Params: map[string]string{"delay": "10ms"}}}

	_, _ = l.Inject(context.Background(), exp)
	_, err := l.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("idempotent Inject: %v", err)
	}
	if callCount != 1 {
		t.Errorf("tc called %d times, want 1", callCount)
	}
}

func TestLatencyInjector_Inject_MissingDelay(t *testing.T) {
	l := NewLatencyInjector()
	_, err := l.Inject(context.Background(), core.Experiment{ID: "lat4", Fault: core.Fault{Params: map[string]string{}}})
	if err == nil {
		t.Fatal("expected error for missing delay")
	}
}

func TestLatencyInjector_Inject_TCError(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	tcFn = func(args ...string) error { return errors.New("tc failed") }

	l := NewLatencyInjector()
	_, err := l.Inject(context.Background(), core.Experiment{ID: "lat5", Fault: core.Fault{Params: map[string]string{"delay": "10ms"}}})
	if err == nil {
		t.Fatal("expected error for tc failure")
	}
}

func TestLatencyInjector_Recover_Success(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	tcFn = func(args ...string) error { return nil }

	l := NewLatencyInjector()
	exp := core.Experiment{ID: "lat6", Fault: core.Fault{Params: map[string]string{"delay": "10ms"}}}
	inj, _ := l.Inject(context.Background(), exp)

	if err := l.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if l.IsActive() {
		t.Error("IsActive should be false after recover")
	}
}

func TestLatencyInjector_Recover_Idempotent(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	tcFn = func(args ...string) error { return nil }

	l := NewLatencyInjector()
	exp := core.Experiment{ID: "lat7", Fault: core.Fault{Params: map[string]string{"delay": "10ms"}}}
	inj, _ := l.Inject(context.Background(), exp)

	_ = l.Recover(context.Background(), inj)
	if err := l.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover again: %v", err)
	}
}

func TestLatencyInjector_Recover_TCError(t *testing.T) {
	orig := tcFn
	defer func() { tcFn = orig }()
	tcFn = func(args ...string) error { return nil }

	l := NewLatencyInjector()
	exp := core.Experiment{ID: "lat8", Fault: core.Fault{Params: map[string]string{"delay": "10ms"}}}
	inj, _ := l.Inject(context.Background(), exp)

	tcFn = func(args ...string) error { return errors.New("tc del failed") }
	if err := l.Recover(context.Background(), inj); err == nil {
		t.Fatal("expected error for tc del failure")
	}
}
