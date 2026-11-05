// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package adapter

import (
	"testing"

	"github.com/qingdujian/abate/core"
)

func TestDegradations(t *testing.T) {
	mt := NewMappingTable(DefaultChaosMeshMappings())
	degs := mt.Degradations()
	if len(degs) == 0 {
		t.Fatal("expected degradations for unmapped sub-codes")
	}
	found := false
	for _, d := range degs {
		if d.SubCode == "F5.1" {
			found = true
		}
	}
	if !found {
		t.Error("F5.1 should be in degradations")
	}
}

func TestMappingTable_LookupSuccess(t *testing.T) {
	mt := NewMappingTable(DefaultChaosMeshMappings())
	fm, err := mt.Lookup("F1.1")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	if fm.UpstreamType != "IOChaos/delay" {
		t.Errorf("UpstreamType = %q", fm.UpstreamType)
	}
}

func TestMappingTable_LookupNotFound(t *testing.T) {
	mt := NewMappingTable(DefaultChaosMeshMappings())
	_, err := mt.Lookup("F5.1")
	if err == nil {
		t.Fatal("expected error for unmapped F5.1")
	}
}

func TestMappingTable_TranslatePassthrough(t *testing.T) {
	mt := NewMappingTable([]FaultMapping{
		{SubCode: "F1", UpstreamType: "Delay", ParamMap: map[string]string{}},
	})
	f := core.Fault{Class: core.FaultClass("F1"), Params: map[string]string{"x": "1"}}
	mapped, err := mt.Translate(f, core.Target{Kind: core.TargetInstance, Name: "i-1"})
	if err != nil {
		t.Fatalf("Translate: %v", err)
	}
	if mapped.Params["x"] != "1" {
		t.Errorf("passthrough param x = %q, want 1", mapped.Params["x"])
	}
}

func TestMappedClasses(t *testing.T) {
	mt := NewMappingTable(DefaultChaosMeshMappings())
	classes := mt.MappedClasses()
	if len(classes) == 0 {
		t.Fatal("expected non-empty mapped classes")
	}
	hasF1 := false
	for _, c := range classes {
		if c == core.F1Latency {
			hasF1 = true
		}
	}
	if !hasF1 {
		t.Error("F1 should be in mapped classes")
	}
}

func TestDefaultLitmusMappings(t *testing.T) {
	mt := NewMappingTable(DefaultLitmusMappings())
	fm, err := mt.Lookup("F2.1")
	if err != nil {
		t.Fatalf("Lookup F2.1: %v", err)
	}
	if fm.UpstreamType != "pod-delete" {
		t.Errorf("UpstreamType = %q, want pod-delete", fm.UpstreamType)
	}
	degs := mt.Degradations()
	if len(degs) != 10 {
		t.Errorf("degradations = %d, want 10", len(degs))
	}
}

func TestNewMappingTable_Empty(t *testing.T) {
	mt := NewMappingTable(nil)
	if len(mt.MappedClasses()) != 0 {
		t.Error("expected empty mapped classes")
	}
	if len(mt.Degradations()) != 20 {
		t.Errorf("degradations = %d, want 20", len(mt.Degradations()))
	}
}