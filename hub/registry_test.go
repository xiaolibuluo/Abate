// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package hub

import (
	"strings"
	"testing"
)

// TestNewSeedRegistry 验证种子表 20 项。
func TestNewSeedRegistry(t *testing.T) {
	r := NewSeedRegistry()
	if r.Count() != 20 {
		t.Errorf("seed count = %d, want 20", r.Count())
	}
	if r.Version() != "v0.1" {
		t.Errorf("version = %q, want v0.1", r.Version())
	}
}

// TestRegistryRegisterAndLookup 验证注册和查找。
func TestRegistryRegisterAndLookup(t *testing.T) {
	r := NewRegistry("v0.2")
	sc := SubCode{Code: "F7.1", Name: "测试", Params: "x"}
	if err := r.Register(sc); err != nil {
		t.Fatalf("Register: %v", err)
	}
	got, err := r.Lookup("F7.1")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	if got.Name != "测试" {
		t.Errorf("name = %q", got.Name)
	}
}

// TestRegistryRegisterEmpty 验证空编码注册失败。
func TestRegistryRegisterEmpty(t *testing.T) {
	r := NewRegistry("v0.2")
	err := r.Register(SubCode{Code: ""})
	if err == nil {
		t.Error("expected error for empty code")
	}
}

// TestRegistryLookupNotFound 验证查找不存在编码。
func TestRegistryLookupNotFound(t *testing.T) {
	r := NewRegistry("v0.2")
	_, err := r.Lookup("F9.9")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Errorf("err = %v, want not found", err)
	}
}

// TestRegistryAllCodesSorted 验证字典序。
func TestRegistryAllCodesSorted(t *testing.T) {
	r := NewRegistry("v0.2")
	_ = r.Register(SubCode{Code: "F2.1"})
	_ = r.Register(SubCode{Code: "F1.1"})
	codes := r.AllCodeStrings()
	if codes[0] != "F1.1" || codes[1] != "F2.1" {
		t.Errorf("codes = %v, want sorted", codes)
	}
}

// TestRegistrySetVersion 验证版本更新。
func TestRegistrySetVersion(t *testing.T) {
	r := NewRegistry("v0.1")
	r.SetVersion("v0.2")
	if r.Version() != "v0.2" {
		t.Errorf("version = %q, want v0.2", r.Version())
	}
}

// TestRegistryMerge 验证合并。
func TestRegistryMerge(t *testing.T) {
	old := NewRegistry("v0.1")
	_ = old.Register(SubCode{Code: "F1.1", Name: "旧"})
	_ = old.Register(SubCode{Code: "F1.2", Name: "不变"})

	new := NewRegistry("v0.2")
	_ = new.Register(SubCode{Code: "F1.1", Name: "新"})
	_ = new.Register(SubCode{Code: "F1.2", Name: "不变"})
	_ = new.Register(SubCode{Code: "F1.3", Name: "新增"})

	added, updated := old.Merge(new)
	if len(added) != 1 || added[0] != "F1.3" {
		t.Errorf("added = %v, want [F1.3]", added)
	}
	if len(updated) != 1 || updated[0] != "F1.1" {
		t.Errorf("updated = %v, want [F1.1]", updated)
	}
	sc, _ := old.Lookup("F1.1")
	if sc.Name != "新" {
		t.Errorf("merged name = %q, want 新", sc.Name)
	}
}
