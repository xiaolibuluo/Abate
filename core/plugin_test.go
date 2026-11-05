// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package core

import (
	"strings"
	"testing"
)

// TestCapabilityCheckProtocol 验证协议版本协商（CES-7.1）：一致放行，不一致拒绝且报错含错误码。
func TestCapabilityCheckProtocol(t *testing.T) {
	ok := Capability{
		Name: "injector-host", Version: "0.1.0", Protocol: ProtocolVersion,
		FaultClasses: []FaultClass{F2Crash, F3Partition},
		TargetKinds:  []TargetKind{TargetInstance},
	}
	if err := ok.CheckProtocol(); err != nil {
		t.Errorf("compatible plugin rejected: %v", err)
	}

	bad := Capability{Name: "legacy", Protocol: "v0"}
	err := bad.CheckProtocol()
	if err == nil {
		t.Fatal("incompatible plugin must be rejected")
	}
	if !strings.Contains(err.Error(), "E1001") {
		t.Errorf("error must carry code E1001, got: %v", err)
	}
	if !strings.Contains(err.Error(), "legacy") {
		t.Errorf("error must name the plugin, got: %v", err)
	}
}

// TestProtocolVersionPinned 锁定当前协议主版本，防止意外变更破坏兼容承诺。
func TestProtocolVersionPinned(t *testing.T) {
	if ProtocolVersion != "v1" {
		t.Errorf("ProtocolVersion=%q, want v1 (CES-7.1)", ProtocolVersion)
	}
}
