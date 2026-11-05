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
	"fmt"
	"strings"
)

// Diff 描述两个注册表版本之间的差异。
type Diff struct {
	Added   []string
	Updated []string
	Removed []string
}

// IsEmpty 判断是否有差异。
func (d Diff) IsEmpty() bool {
	return len(d.Added) == 0 && len(d.Updated) == 0 && len(d.Removed) == 0
}

// ComputeDiff 比较新旧注册表，返回差异。
func ComputeDiff(old, new *Registry) Diff {
	oldCodes := make(map[string]SubCode)
	for _, sc := range old.AllCodes() {
		oldCodes[sc.Code] = sc
	}
	newCodes := make(map[string]SubCode)
	for _, sc := range new.AllCodes() {
		newCodes[sc.Code] = sc
	}
	var diff Diff
	for code, newSC := range newCodes {
		if oldSC, exists := oldCodes[code]; exists {
			if oldSC != newSC {
				diff.Updated = append(diff.Updated, code)
			}
		} else {
			diff.Added = append(diff.Added, code)
		}
	}
	for code := range oldCodes {
		if _, exists := newCodes[code]; !exists {
			diff.Removed = append(diff.Removed, code)
		}
	}
	return diff
}

// Announcement 是 ChaosHub 子编码表版本公告。
type Announcement struct {
	OldVersion string
	NewVersion string
	Diff       Diff
}

// NewAnnouncement 从差异构造公告。
func NewAnnouncement(old, new *Registry) *Announcement {
	return &Announcement{
		OldVersion: old.Version(),
		NewVersion: new.Version(),
		Diff:       ComputeDiff(old, new),
	}
}

// Format 将公告格式化为人类可读文本。
func (a *Announcement) Format() string {
	if a.Diff.IsEmpty() {
		return fmt.Sprintf("ChaosHub %s → %s: no changes", a.OldVersion, a.NewVersion)
	}
	var parts []string
	parts = append(parts, fmt.Sprintf("ChaosHub %s → %s:", a.OldVersion, a.NewVersion))
	if len(a.Diff.Added) > 0 {
		parts = append(parts, fmt.Sprintf("  + added: %s", strings.Join(a.Diff.Added, ", ")))
	}
	if len(a.Diff.Updated) > 0 {
		parts = append(parts, fmt.Sprintf("  ~ updated: %s", strings.Join(a.Diff.Updated, ", ")))
	}
	if len(a.Diff.Removed) > 0 {
		parts = append(parts, fmt.Sprintf("  - removed: %s", strings.Join(a.Diff.Removed, ", ")))
	}
	return strings.Join(parts, "\n")
}
