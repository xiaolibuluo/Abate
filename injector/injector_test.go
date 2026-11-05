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
	"os"
	"path/filepath"
	"testing"

	"github.com/qingdujian/abate/core"
)

func testExperiment() core.Experiment {
	return core.Experiment{
		ID:          "inj-test",
		Seed:        42,
		Criticality: 3,
		Target:      core.Target{Kind: core.TargetService, Name: "svc"},
		Fault:       core.Fault{Class: core.F1Latency, Params: map[string]string{"delay": "100"}},
		Blast:       core.BlastRadius{MaxInstances: 1, MaxDurationMs: 60000, MaxTraffic: 1000},
		Steady:       []core.SLIAssert{{SLI: "latency", Assert: "< 500"}},
	}
}

func TestNew(t *testing.T) {
	dir := t.TempDir()
	in, err := New(dir)
	if err != nil {
		t.Fatalf("new: %v", err)
	}
	if in == nil {
		t.Fatal("injector is nil")
	}
	cap := in.Capability()
	if cap.Name != "file-injector" {
		t.Errorf("name = %q", cap.Name)
	}
	if cap.Protocol != core.ProtocolVersion {
		t.Errorf("protocol = %q", cap.Protocol)
	}
	if len(cap.FaultClasses) != 6 {
		t.Errorf("faultClasses = %d", len(cap.FaultClasses))
	}
}

func TestNewMkdirError(t *testing.T) {
	_, err := New("/proc/cannot/create")
	if err == nil {
		t.Error("expected mkdir error")
	}
}

func TestInjectAndRecover(t *testing.T) {
	dir := t.TempDir()
	in, _ := New(dir)
	e := testExperiment()
	inj, err := in.Inject(context.Background(), e)
	if err != nil {
		t.Fatalf("inject: %v", err)
	}
	if inj.ID != "inj-test" {
		t.Errorf("id = %q", inj.ID)
	}
	if !in.IsActive() {
		t.Error("should be active")
	}
	if err := in.Recover(context.Background(), inj); err != nil {
		t.Fatalf("recover: %v", err)
	}
	if in.IsActive() {
		t.Error("should be inactive")
	}
}

func TestInjectIdempotent(t *testing.T) {
	dir := t.TempDir()
	in, _ := New(dir)
	e := testExperiment()
	i1, _ := in.Inject(context.Background(), e)
	i2, _ := in.Inject(context.Background(), e)
	if i1.ID != i2.ID {
		t.Errorf("idempotent: %q vs %q", i1.ID, i2.ID)
	}
}

func TestRecoverNoFile(t *testing.T) {
	dir := t.TempDir()
	in, _ := New(dir)
	if err := in.Recover(context.Background(), core.Injection{ID: "ghost"}); err != nil {
		t.Errorf("recover missing: %v", err)
	}
}

func TestInjectWriteError(t *testing.T) {
	in := &Injector{dir: "/proc/cannot"}
	_, err := in.Inject(context.Background(), testExperiment())
	if err == nil {
		t.Error("expected write error")
	}
}

func TestRecoverRemoveError(t *testing.T) {
	dir := t.TempDir()
	in, _ := New(dir)
	blockedDir := filepath.Join(dir, "blocked.inj")
	_ = os.MkdirAll(blockedDir, 0755)
	_ = os.WriteFile(filepath.Join(blockedDir, "inner"), []byte("x"), 0644)
	err := in.Recover(context.Background(), core.Injection{ID: "blocked"})
	if err == nil {
		t.Error("expected remove error")
	}
}

func TestCapability(t *testing.T) {
	in := &Injector{cap: core.Capability{Name: "test", Protocol: core.ProtocolVersion}}
	cap := in.Capability()
	if cap.Name != "test" {
		t.Errorf("name = %q", cap.Name)
	}
}
