// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package reporter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/qingdujian/abate/core"
)

var errTestReporter = fmt.Errorf("reporter test error")

func testScore() core.Score {
	return core.Score{
		Final:    8900,
		Raw:      8900,
		Coverage: 9500,
		Factor:   10000,
		Dims: map[core.Dimension]core.BP{
			core.DimAR:  9000,
			core.DimREC: 8000,
			core.DimDC:  9000,
			core.DimDI:  9000,
			core.DimBRC: 10000,
		},
	}
}

func TestNew(t *testing.T) {
	dir := t.TempDir()
	r, err := New(dir)
	if err != nil {
		t.Fatalf("new: %v", err)
	}
	if r == nil {
		t.Fatal("reporter is nil")
	}
	cap := r.Capability()
	if cap.Name != "file-reporter" {
		t.Errorf("cap name = %q", cap.Name)
	}
	if cap.Protocol != core.ProtocolVersion {
		t.Errorf("protocol = %q", cap.Protocol)
	}
}

func TestNewMkdirError(t *testing.T) {
	_, err := New("/proc/cannot/create")
	if err == nil {
		t.Error("expected mkdir error")
	}
}

func TestEmit(t *testing.T) {
	dir := t.TempDir()
	r, _ := New(dir)
	path, err := r.Emit(context.Background(), testScore())
	if err != nil {
		t.Fatalf("emit: %v", err)
	}
	if path == "" {
		t.Error("path is empty")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	var rep Report
	if err := json.Unmarshal(data, &rep); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if rep.SchemaVersion != "ces/v0.1" {
		t.Errorf("schema = %q", rep.SchemaVersion)
	}
	if rep.ContentHash == "" {
		t.Error("hash is empty")
	}
}

func TestEmitWriteError(t *testing.T) {
	r := &Reporter{dir: "/proc/cannot"}
	_, err := r.Emit(context.Background(), testScore())
	if err == nil {
		t.Error("expected write error")
	}
}

func TestReportContentHash(t *testing.T) {
	dir := t.TempDir()
	r, _ := New(dir)
	_, _ = r.Emit(context.Background(), testScore())
	path := filepath.Join(dir, "report.json")
	data, _ := os.ReadFile(path)
	var rep Report
	_ = json.Unmarshal(data, &rep)
	if rep.ContentHash == "" {
		t.Error("hash should not be empty")
	}
	if len(rep.ContentHash) != 64 {
		t.Errorf("hash len = %d, want 64", len(rep.ContentHash))
	}
}

func TestEmitMarshalErrorFirst(t *testing.T) {
	dir := t.TempDir()
	r, _ := New(dir)
	orig := marshalIndent
	callCount := 0
	marshalIndent = func(_ interface{}, _, _ string) ([]byte, error) {
		callCount++
		if callCount == 1 {
			return nil, errTestReporter
		}
		return orig(nil, "", "  ")
	}
	defer func() { marshalIndent = orig }()
	_, err := r.Emit(context.Background(), testScore())
	if err == nil {
		t.Error("expected marshal error")
	}
}

func TestEmitMarshalErrorSecond(t *testing.T) {
	dir := t.TempDir()
	r, _ := New(dir)
	orig := marshalIndent
	callCount := 0
	marshalIndent = func(_ interface{}, _, _ string) ([]byte, error) {
		callCount++
		if callCount == 2 {
			return nil, errTestReporter
		}
		return orig(nil, "", "  ")
	}
	defer func() { marshalIndent = orig }()
	_, err := r.Emit(context.Background(), testScore())
	if err == nil {
		t.Error("expected marshal error")
	}
}
