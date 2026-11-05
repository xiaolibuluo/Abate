// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package store

import (
	"fmt"
	"testing"

	"github.com/qingdujian/abate/core"
)

var errTestStore = fmt.Errorf("store test error")

func testRecord(id string) core.Record {
	return core.Record{
		Exp: core.Experiment{
			ID:          id,
			Seed:        42,
			Criticality: 3,
			Target:      core.Target{Kind: core.TargetService, Name: "svc"},
			Fault:       core.Fault{Class: core.F1Latency, Params: map[string]string{"delay": "100"}},
			Blast:       core.BlastRadius{MaxInstances: 1, MaxDurationMs: 60000},
			Steady:      []core.SLIAssert{{SLI: "latency", Assert: "< 500"}},
		},
		Evidence:       core.EvidenceA,
		HoldMs:         60000,
		WindowMs:       60000,
		RecoveryMs:     60000,
		InvariantPass:  5,
		InvariantTotal:  5,
		DeclaredRadius: 1,
	}
}

func testScore() core.Score {
	return core.Score{
		Final:    8900,
		Raw:      8900,
		Coverage: 9500,
		Factor:   10000,
		Dims: map[core.Dimension]core.BP{},
	}
}

func TestMemoryStoreSaveAndGetRecord(t *testing.T) {
	m := NewMemoryStore()
	r := testRecord("exp-a")
	if err := m.SaveRecord(r); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, ok, err := m.GetRecord("exp-a")
	if err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if got.Exp.ID != "exp-a" {
		t.Errorf("id = %q", got.Exp.ID)
	}
}

func TestMemoryStoreGetRecordNotFound(t *testing.T) {
	m := NewMemoryStore()
	_, ok, err := m.GetRecord("ghost")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if ok {
		t.Error("should not be found")
	}
}

func TestMemoryStoreListRecords(t *testing.T) {
	m := NewMemoryStore()
	_ = m.SaveRecord(testRecord("zeta"))
	_ = m.SaveRecord(testRecord("alpha"))
	list, err := m.ListRecords()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("len = %d", len(list))
	}
	if list[0].Exp.ID != "alpha" || list[1].Exp.ID != "zeta" {
		t.Errorf("order: %s, %s", list[0].Exp.ID, list[1].Exp.ID)
	}
}

func TestMemoryStoreSaveAndGetScore(t *testing.T) {
	m := NewMemoryStore()
	s := testScore()
	if err := m.SaveScore(s); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, ok, err := m.GetScore()
	if err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if got.Final != 8900 {
		t.Errorf("final = %d", got.Final)
	}
}

func TestMemoryStoreGetScoreEmpty(t *testing.T) {
	m := NewMemoryStore()
	_, ok, err := m.GetScore()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if ok {
		t.Error("should not be found")
	}
}
