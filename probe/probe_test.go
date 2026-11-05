// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package probe

import (
	"context"
	"testing"

	"github.com/qingdujian/abate/core"
)

func TestNew(t *testing.T) {
	p := New(42)
	if p == nil {
		t.Fatal("probe is nil")
	}
	cap := p.Capability()
	if cap.Name != "file-probe" {
		t.Errorf("name = %q", cap.Name)
	}
	if cap.Protocol != core.ProtocolVersion {
		t.Errorf("protocol = %q", cap.Protocol)
	}
}

func TestObserve(t *testing.T) {
	p := New(42)
	tgt := core.Target{Kind: core.TargetService, Name: "svc"}
	s1, err := p.Observe(context.Background(), tgt)
	if err != nil {
		t.Fatalf("observe: %v", err)
	}
	if s1.SLI != "svc" {
		t.Errorf("sli = %q", s1.SLI)
	}
	p2 := New(42)
	s2, _ := p2.Observe(context.Background(), tgt)
	if s1.Value != s2.Value {
		t.Errorf("deterministic: %d vs %d", s1.Value, s2.Value)
	}
}

func TestObserveEmptyName(t *testing.T) {
	p := New(42)
	_, err := p.Observe(context.Background(), core.Target{Kind: core.TargetService, Name: ""})
	if err == nil {
		t.Error("expected error for empty name")
	}
}

func TestCapability(t *testing.T) {
	p := &Probe{cap: core.Capability{Name: "test", Protocol: core.ProtocolVersion}}
	cap := p.Capability()
	if cap.Name != "test" {
		t.Errorf("name = %q", cap.Name)
	}
}
