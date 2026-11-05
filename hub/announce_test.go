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

// TestDiffIsEmpty 验证空差异判断。
func TestDiffIsEmpty(t *testing.T) {
	if !(Diff{}).IsEmpty() {
		t.Error("empty diff should be empty")
	}
	if (Diff{Added: []string{"F1"}}).IsEmpty() {
		t.Error("diff with added should not be empty")
	}
}

// TestComputeDiffNoChange 验证无差异。
func TestComputeDiffNoChange(t *testing.T) {
	old := NewSeedRegistry()
	new := NewSeedRegistry()
	diff := ComputeDiff(old, new)
	if !diff.IsEmpty() {
		t.Error("identical registries should have no diff")
	}
}

// TestComputeDiffAdded 验证新增。
func TestComputeDiffAdded(t *testing.T) {
	old := NewRegistry("v0.1")
	new := NewRegistry("v0.2")
	_ = new.Register(SubCode{Code: "F7.1", Name: "新"})
	diff := ComputeDiff(old, new)
	if len(diff.Added) != 1 || diff.Added[0] != "F7.1" {
		t.Errorf("added = %v, want [F7.1]", diff.Added)
	}
}

// TestComputeDiffUpdated 验证更新。
func TestComputeDiffUpdated(t *testing.T) {
	old := NewRegistry("v0.1")
	_ = old.Register(SubCode{Code: "F1.1", Name: "旧"})
	new := NewRegistry("v0.2")
	_ = new.Register(SubCode{Code: "F1.1", Name: "新"})
	diff := ComputeDiff(old, new)
	if len(diff.Updated) != 1 || diff.Updated[0] != "F1.1" {
		t.Errorf("updated = %v, want [F1.1]", diff.Updated)
	}
}

// TestComputeDiffRemoved 验证删除。
func TestComputeDiffRemoved(t *testing.T) {
	old := NewRegistry("v0.1")
	_ = old.Register(SubCode{Code: "F1.1", Name: "测试"})
	new := NewRegistry("v0.2")
	diff := ComputeDiff(old, new)
	if len(diff.Removed) != 1 || diff.Removed[0] != "F1.1" {
		t.Errorf("removed = %v, want [F1.1]", diff.Removed)
	}
}

// TestAnnouncementFormatNoChange 验证无差异公告格式。
func TestAnnouncementFormatNoChange(t *testing.T) {
	old := NewSeedRegistry()
	new := NewSeedRegistry()
	a := NewAnnouncement(old, new)
	out := a.Format()
	if !strings.Contains(out, "no changes") {
		t.Errorf("format = %q, want no changes", out)
	}
}

// TestAnnouncementFormatWithChanges 验证有差异公告格式。
func TestAnnouncementFormatWithChanges(t *testing.T) {
	old := NewRegistry("v0.1")
	_ = old.Register(SubCode{Code: "F1.1", Name: "旧"})
	_ = old.Register(SubCode{Code: "F1.2", Name: "不变"})

	new := NewRegistry("v0.2")
	_ = new.Register(SubCode{Code: "F1.1", Name: "新"})
	_ = new.Register(SubCode{Code: "F1.2", Name: "不变"})
	_ = new.Register(SubCode{Code: "F1.3", Name: "新增"})

	a := NewAnnouncement(old, new)
	out := a.Format()
	if !strings.Contains(out, "v0.1") || !strings.Contains(out, "v0.2") {
		t.Errorf("format missing versions: %q", out)
	}
	if !strings.Contains(out, "F1.3") {
		t.Errorf("format missing added F1.3: %q", out)
	}
	if !strings.Contains(out, "F1.1") {
		t.Errorf("format missing updated F1.1: %q", out)
	}
}

// TestAnnouncementFormatRemoved 验证删除项公告格式。
func TestAnnouncementFormatRemoved(t *testing.T) {
	old := NewRegistry("v0.1")
	_ = old.Register(SubCode{Code: "F9.9", Name: "废弃"})
	new := NewRegistry("v0.2")

	a := NewAnnouncement(old, new)
	out := a.Format()
	if !strings.Contains(out, "F9.9") {
		t.Errorf("format missing removed F9.9: %q", out)
	}
}
