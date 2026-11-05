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

func TestNewDiskInjector(t *testing.T) {
	d, err := NewDiskInjector(t.TempDir())
	if err != nil {
		t.Fatalf("NewDiskInjector: %v", err)
	}
	if d.Capability().Name != "disk-injector" {
		t.Errorf("name = %s", d.Capability().Name)
	}
	if d.IsActive() {
		t.Error("IsActive should be false on new injector")
	}
}

func TestNewDiskInjector_MkdirFail(t *testing.T) {
	blocker := filepath.Join(t.TempDir(), "blocker")
	os.WriteFile(blocker, []byte("x"), 0644)
	_, err := NewDiskInjector(blocker)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDiskInject_Recover(t *testing.T) {
	d, _ := NewDiskInjector(t.TempDir())
	ctx := context.Background()
	exp := core.Experiment{ID: "exp1", Fault: core.Fault{Params: map[string]string{"size": "1M"}}}

	inj, err := d.Inject(ctx, exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if !d.IsActive() {
		t.Error("IsActive should be true after inject")
	}

	inj2, err := d.Inject(ctx, exp)
	if err != nil {
		t.Fatalf("Idempotent Inject: %v", err)
	}
	if inj2.ID != inj.ID {
		t.Error("idempotent inject should return same ID")
	}

	if err := d.Recover(ctx, inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if d.IsActive() {
		t.Error("IsActive should be false after recover")
	}

	if err := d.Recover(ctx, inj); err != nil {
		t.Fatalf("Recover again: %v", err)
	}
}

func TestDiskInject_MissingSize(t *testing.T) {
	d, _ := NewDiskInjector(t.TempDir())
	_, err := d.Inject(context.Background(), core.Experiment{ID: "exp2", Fault: core.Fault{Params: map[string]string{}}})
	if err == nil {
		t.Fatal("expected error for missing size")
	}
}

func TestDiskInject_CreateFail(t *testing.T) {
	d, _ := NewDiskInjector(t.TempDir())
	blocker := filepath.Join(t.TempDir(), "blocker")
	os.WriteFile(blocker, []byte("x"), 0644)
	d.dir = blocker
	_, err := d.Inject(context.Background(), core.Experiment{ID: "exp3", Fault: core.Fault{Params: map[string]string{"size": "1M"}}})
	if err == nil {
		t.Fatal("expected error for create fail")
	}
}

func TestDiskInject_TruncateFail(t *testing.T) {
	d, _ := NewDiskInjector(t.TempDir())
	d.truncateFn = func(f *os.File, size int64) error { return os.ErrClosed }
	_, err := d.Inject(context.Background(), core.Experiment{ID: "exp4", Fault: core.Fault{Params: map[string]string{"size": "1M"}}})
	if err == nil {
		t.Fatal("expected error for truncate fail")
	}
}

func TestDiskRecover_RemoveFail(t *testing.T) {
	d, _ := NewDiskInjector(t.TempDir())
	path := filepath.Join(d.dir, "exp5.fill")
	os.MkdirAll(filepath.Join(path, "sub"), 0755)
	d.active["exp5"] = path
	err := d.Recover(context.Background(), core.Injection{ID: "exp5"})
	if err == nil {
		t.Fatal("expected error for remove fail")
	}
}

func TestDiskRecover_NotExist(t *testing.T) {
	d, _ := NewDiskInjector(t.TempDir())
	d.Inject(context.Background(), core.Experiment{ID: "exp6", Fault: core.Fault{Params: map[string]string{"size": "1M"}}})
	os.Remove(filepath.Join(d.dir, "exp6.fill"))
	if err := d.Recover(context.Background(), core.Injection{ID: "exp6"}); err != nil {
		t.Fatalf("Recover with not-exist: %v", err)
	}
}

func TestDefaultTruncate(t *testing.T) {
	f, _ := os.CreateTemp(t.TempDir(), "test")
	defer f.Close()
	if err := defaultTruncate(f, 100); err != nil {
		t.Fatalf("defaultTruncate: %v", err)
	}
	info, _ := f.Stat()
	if info.Size() != 100 {
		t.Errorf("size = %d, want 100", info.Size())
	}
}

func TestParseSize(t *testing.T) {
	tests := []struct {
		input string
		want  int64
	}{
		{"", 0},
		{"100", 100},
		{"1K", 1024},
		{"1k", 1024},
		{"1M", 1024 * 1024},
		{"1m", 1024 * 1024},
		{"1G", 1024 * 1024 * 1024},
		{"1g", 1024 * 1024 * 1024},
		{"0", 0},
		{"abc", 0},
		{"1xM", 0},
		{"1x", 0},
		{"K", 0},
	}
	for _, tt := range tests {
		if got := parseSize(tt.input); got != tt.want {
			t.Errorf("parseSize(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
