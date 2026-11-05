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
	"time"

	"github.com/qingdujian/abate/core"
)

func TestClockSkewInjector_Capability(t *testing.T) {
	c := NewClockSkewInjector()
	cap := c.Capability()
	if cap.Name != "clockskew-injector" {
		t.Errorf("name = %s", cap.Name)
	}
	if len(cap.FaultClasses) != 1 || cap.FaultClasses[0] != core.F6ClockSkew {
		t.Errorf("fault classes = %v", cap.FaultClasses)
	}
	if c.IsActive() {
		t.Error("IsActive should be false on new injector")
	}
}

func TestClockSkewInjector_Inject_Success(t *testing.T) {
	var gotOffset time.Duration
	c := NewClockSkewInjector()
	c.setClock = func(offset time.Duration) error {
		gotOffset = offset
		return nil
	}

	exp := core.Experiment{ID: "clk1", Fault: core.Fault{Params: map[string]string{"offset": "5000"}}}
	inj, err := c.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "clk1" {
		t.Errorf("injection ID = %s", inj.ID)
	}
	if !c.IsActive() {
		t.Error("IsActive should be true after inject")
	}
	if gotOffset != 5*time.Second {
		t.Errorf("offset = %v, want 5s", gotOffset)
	}
}

func TestClockSkewInjector_Inject_Idempotent(t *testing.T) {
	callCount := 0
	c := NewClockSkewInjector()
	c.setClock = func(offset time.Duration) error {
		callCount++
		return nil
	}

	exp := core.Experiment{ID: "clk2", Fault: core.Fault{Params: map[string]string{"offset": "1000"}}}
	_, _ = c.Inject(context.Background(), exp)
	_, err := c.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("idempotent Inject: %v", err)
	}
	if callCount != 1 {
		t.Errorf("setClock called %d times, want 1", callCount)
	}
}

func TestClockSkewInjector_Inject_MissingOffset(t *testing.T) {
	c := NewClockSkewInjector()
	_, err := c.Inject(context.Background(), core.Experiment{ID: "clk3", Fault: core.Fault{Params: map[string]string{}}})
	if err == nil {
		t.Fatal("expected error for missing offset")
	}
}

func TestClockSkewInjector_Inject_InvalidOffset(t *testing.T) {
	c := NewClockSkewInjector()
	_, err := c.Inject(context.Background(), core.Experiment{ID: "clk4", Fault: core.Fault{Params: map[string]string{"offset": "abc"}}})
	if err == nil {
		t.Fatal("expected error for invalid offset")
	}
}

func TestClockSkewInjector_Inject_SetClockError(t *testing.T) {
	c := NewClockSkewInjector()
	c.setClock = func(offset time.Duration) error { return errors.New("date failed") }
	_, err := c.Inject(context.Background(), core.Experiment{ID: "clk5", Fault: core.Fault{Params: map[string]string{"offset": "100"}}})
	if err == nil {
		t.Fatal("expected error for setClock failure")
	}
}

func TestClockSkewInjector_Recover_Success(t *testing.T) {
	var recoverOffset time.Duration
	c := NewClockSkewInjector()
	c.setClock = func(offset time.Duration) error {
		recoverOffset = offset
		return nil
	}

	exp := core.Experiment{ID: "clk6", Fault: core.Fault{Params: map[string]string{"offset": "3000"}}}
	inj, _ := c.Inject(context.Background(), exp)

	recoverOffset = 0
	if err := c.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if c.IsActive() {
		t.Error("IsActive should be false after recover")
	}
	if recoverOffset != -3*time.Second {
		t.Errorf("recover offset = %v, want -3s", recoverOffset)
	}
}

func TestClockSkewInjector_Recover_Idempotent(t *testing.T) {
	c := NewClockSkewInjector()
	c.setClock = func(offset time.Duration) error { return nil }

	exp := core.Experiment{ID: "clk7", Fault: core.Fault{Params: map[string]string{"offset": "1000"}}}
	inj, _ := c.Inject(context.Background(), exp)

	_ = c.Recover(context.Background(), inj)
	if err := c.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover again: %v", err)
	}
}

func TestClockSkewInjector_Recover_SetClockError(t *testing.T) {
	c := NewClockSkewInjector()
	c.setClock = func(offset time.Duration) error { return nil }

	exp := core.Experiment{ID: "clk8", Fault: core.Fault{Params: map[string]string{"offset": "1000"}}}
	inj, _ := c.Inject(context.Background(), exp)

	c.setClock = func(offset time.Duration) error { return errors.New("date failed") }
	if err := c.Recover(context.Background(), inj); err == nil {
		t.Fatal("expected error for recover setClock failure")
	}
}

func TestDateCmdFn(t *testing.T) {
	// Use invalid date string to exercise the function body without changing system clock.
	if err := dateCmdFn("invalid-date"); err == nil {
		t.Fatal("expected error for invalid date string")
	}
}

func TestDefaultSetClock(t *testing.T) {
	orig := dateCmdFn
	defer func() { dateCmdFn = orig }()
	var gotTime string
	dateCmdFn = func(newTime string) error {
		gotTime = newTime
		return nil
	}
	if err := defaultSetClock(5 * time.Second); err != nil {
		t.Fatalf("defaultSetClock: %v", err)
	}
	if gotTime == "" {
		t.Error("dateCmdFn was not called")
	}
}

func TestDefaultSetClock_Error(t *testing.T) {
	orig := dateCmdFn
	defer func() { dateCmdFn = orig }()
	dateCmdFn = func(newTime string) error { return errors.New("date failed") }
	if err := defaultSetClock(1 * time.Second); err == nil {
		t.Fatal("expected error from dateCmdFn")
	}
}
