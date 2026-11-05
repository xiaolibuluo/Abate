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
	"os"
	"path/filepath"
	"testing"

	"github.com/qingdujian/abate/core"
)

func TestCorruptionInjector_Capability(t *testing.T) {
	c := NewCorruptionInjector()
	cap := c.Capability()
	if cap.Name != "corruption-injector" {
		t.Errorf("name = %s", cap.Name)
	}
	if len(cap.FaultClasses) != 1 || cap.FaultClasses[0] != core.F5Corruption {
		t.Errorf("fault classes = %v", cap.FaultClasses)
	}
	if c.IsActive() {
		t.Error("IsActive should be false on new injector")
	}
}

func TestCorruptionInjector_Inject_Success(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00, 0xFF, 0x55}, 0644)

	c := NewCorruptionInjector()
	exp := core.Experiment{ID: "cor1", Fault: core.Fault{Params: map[string]string{
		"file":   fpath,
		"offset": "1",
		"bit":    "0",
	}}}

	inj, err := c.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "cor1" {
		t.Errorf("injection ID = %s", inj.ID)
	}
	if !c.IsActive() {
		t.Error("IsActive should be true after inject")
	}

	data, _ := os.ReadFile(fpath)
	if data[1] != 0xFE {
		t.Errorf("byte after flip = 0x%02X, want 0xFE", data[1])
	}
}

func TestCorruptionInjector_Inject_Defaults(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00, 0x01, 0x02}, 0644)

	c := NewCorruptionInjector()
	exp := core.Experiment{ID: "cor2", Fault: core.Fault{Params: map[string]string{
		"file": fpath,
	}}}

	_, err := c.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject with defaults: %v", err)
	}

	data, _ := os.ReadFile(fpath)
	if data[0] != 0x01 {
		t.Errorf("byte after flip = 0x%02X, want 0x01", data[0])
	}
}

func TestCorruptionInjector_Inject_Idempotent(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00}, 0644)

	c := NewCorruptionInjector()
	exp := core.Experiment{ID: "cor3", Fault: core.Fault{Params: map[string]string{"file": fpath}}}

	_, _ = c.Inject(context.Background(), exp)
	data1, _ := os.ReadFile(fpath)
	_, err := c.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("idempotent Inject: %v", err)
	}
	data2, _ := os.ReadFile(fpath)
	if data2[0] != data1[0] {
		t.Error("idempotent inject should not modify file again")
	}
}

func TestCorruptionInjector_Inject_MissingFile(t *testing.T) {
	c := NewCorruptionInjector()
	_, err := c.Inject(context.Background(), core.Experiment{ID: "cor4", Fault: core.Fault{Params: map[string]string{}}})
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestCorruptionInjector_Inject_InvalidOffset(t *testing.T) {
	c := NewCorruptionInjector()
	_, err := c.Inject(context.Background(), core.Experiment{ID: "cor5", Fault: core.Fault{Params: map[string]string{
		"file":   "/tmp/x",
		"offset": "abc",
	}}})
	if err == nil {
		t.Fatal("expected error for invalid offset")
	}
}

func TestCorruptionInjector_Inject_InvalidBit(t *testing.T) {
	c := NewCorruptionInjector()
	_, err := c.Inject(context.Background(), core.Experiment{ID: "cor6", Fault: core.Fault{Params: map[string]string{
		"file": "/tmp/x",
		"bit":  "abc",
	}}})
	if err == nil {
		t.Fatal("expected error for invalid bit")
	}
}

func TestCorruptionInjector_Inject_CorruptError(t *testing.T) {
	c := NewCorruptionInjector()
	c.corruptFn = func(path string, offset int64, bit int) ([]byte, error) {
		return nil, errors.New("corrupt failed")
	}
	_, err := c.Inject(context.Background(), core.Experiment{ID: "cor7", Fault: core.Fault{Params: map[string]string{"file": "/tmp/x"}}})
	if err == nil {
		t.Fatal("expected error for corrupt failure")
	}
}

func TestCorruptionInjector_Recover_Success(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	orig := []byte{0x00, 0xFF, 0x55}
	os.WriteFile(fpath, orig, 0644)

	c := NewCorruptionInjector()
	exp := core.Experiment{ID: "cor8", Fault: core.Fault{Params: map[string]string{"file": fpath, "offset": "0", "bit": "3"}}}
	inj, _ := c.Inject(context.Background(), exp)

	if err := c.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if c.IsActive() {
		t.Error("IsActive should be false after recover")
	}
	data, _ := os.ReadFile(fpath)
	if data[0] != 0x00 {
		t.Errorf("byte after recover = 0x%02X, want 0x00", data[0])
	}
}

func TestCorruptionInjector_Recover_Idempotent(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00}, 0644)

	c := NewCorruptionInjector()
	exp := core.Experiment{ID: "cor9", Fault: core.Fault{Params: map[string]string{"file": fpath}}}
	inj, _ := c.Inject(context.Background(), exp)

	_ = c.Recover(context.Background(), inj)
	if err := c.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover again: %v", err)
	}
}

func TestCorruptionInjector_Recover_RestoreError(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00}, 0644)

	c := NewCorruptionInjector()
	exp := core.Experiment{ID: "cor10", Fault: core.Fault{Params: map[string]string{"file": fpath}}}
	inj, _ := c.Inject(context.Background(), exp)

	c.restoreFn = func(path string, backup []byte) error { return errors.New("restore failed") }
	if err := c.Recover(context.Background(), inj); err == nil {
		t.Fatal("expected error for restore failure")
	}
}

func TestDefaultCorrupt_ReadError(t *testing.T) {
	_, err := defaultCorrupt("/nonexistent/path/file", 0, 0)
	if err == nil {
		t.Fatal("expected error for read failure")
	}
}

func TestDefaultCorrupt_OffsetOutOfRange(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00}, 0644)

	_, err := defaultCorrupt(fpath, 10, 0)
	if err == nil {
		t.Fatal("expected error for offset out of range")
	}
}

func TestDefaultCorrupt_NegativeOffset(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	os.WriteFile(fpath, []byte{0x00}, 0644)

	_, err := defaultCorrupt(fpath, -1, 0)
	if err == nil {
		t.Fatal("expected error for negative offset")
	}
}

func TestDefaultRestore(t *testing.T) {
	tmp := t.TempDir()
	fpath := filepath.Join(tmp, "data.bin")
	backup := []byte{0x42}
	if err := defaultRestore(fpath, backup); err != nil {
		t.Fatalf("defaultRestore: %v", err)
	}
	data, _ := os.ReadFile(fpath)
	if data[0] != 0x42 {
		t.Errorf("restored byte = 0x%02X, want 0x42", data[0])
	}
}
