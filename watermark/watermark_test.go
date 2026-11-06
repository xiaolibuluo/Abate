// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package watermark

import (
	"strings"
	"testing"
)

// TestWatermarkConstants 验证代码级数字水印常量完整（任务一）。
// TestWatermarkConstants verifies the code-level watermark constants.
func TestWatermarkConstants(t *testing.T) {
	if Project != "Abate" {
		t.Errorf("Project = %q", Project)
	}
	if Vendor != "小黎的部落" {
		t.Errorf("Vendor = %q", Vendor)
	}
	if Website != "www.qingdujian.com" {
		t.Errorf("Website = %q", Website)
	}
	if Contact != "xiaolibuluo@163.com" {
		t.Errorf("Contact = %q", Contact)
	}
	if License != "AGPL-3.0-only + commercial" {
		t.Errorf("License = %q", License)
	}
	if Spec != "CES v0.1" {
		t.Errorf("Spec = %q", Spec)
	}
	if !strings.Contains(Copyright, "2026") || !strings.Contains(Copyright, Website) {
		t.Errorf("Copyright = %q", Copyright)
	}
}

// TestBanner 验证开机亮牌横幅包含全部水印要素且格式正确（任务二）。
// TestBanner verifies the boot banner carries all watermark elements.
func TestBanner(t *testing.T) {
	b := Banner("v0.1.0")
	for _, want := range []string{
		Project, Tagline, Copyright, License, Contact, Website, Spec, "v0.1.0",
	} {
		if !strings.Contains(b, want) {
			t.Errorf("banner missing %q:\n%s", want, b)
		}
	}
	if strings.HasSuffix(b, "\n") {
		t.Error("banner must not end with a newline")
	}
	if lines := strings.Count(b, "\n") + 1; lines != 5 {
		t.Errorf("banner lines = %d, want 5", lines)
	}
}
