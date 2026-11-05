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

import "testing"

// TestPerturbationKindValid 验证两种扰动形态的合法性判定。
func TestPerturbationKindValid(t *testing.T) {
	if !PerturbStatic.Valid() || !PerturbDynamic.Valid() {
		t.Error("static/dynamic must be valid")
	}
	for _, k := range []PerturbationKind{"", "hybrid", "STATIC"} {
		if k.Valid() {
			t.Errorf("%q must be invalid", k)
		}
	}
}

// TestHypothesisValidate 验证 HPOV 第一步假设的完整性校验。
func TestHypothesisValidate(t *testing.T) {
	ok := Hypothesis{Statement: "tests kill this mutant", Subject: "core"}
	if err := ok.Validate(); err != nil {
		t.Errorf("valid hypothesis rejected: %v", err)
	}
	if err := (Hypothesis{Subject: "core"}).Validate(); err == nil {
		t.Error("empty statement must be rejected")
	}
	if err := (Hypothesis{Statement: "s"}).Validate(); err == nil {
		t.Error("empty subject must be rejected")
	}
}

// TestObservationValidate 验证 HPOV 第三步观测的完整性与判定语义。
func TestObservationValidate(t *testing.T) {
	ok := Observation{PerturbationID: "mut:a.go:1:1:M1", Outcome: OutcomeContained}
	if err := ok.Validate(); err != nil {
		t.Errorf("valid observation rejected: %v", err)
	}
	esc := Observation{PerturbationID: "mut:a.go:1:1:M1", Outcome: OutcomeEscaped}
	if err := esc.Validate(); err != nil {
		t.Errorf("escaped observation rejected: %v", err)
	}
	if err := (Observation{Outcome: OutcomeContained}).Validate(); err == nil {
		t.Error("empty perturbationId must be rejected")
	}
	bad := Observation{PerturbationID: "x", Outcome: "unknown"}
	if err := bad.Validate(); err == nil {
		t.Error("invalid outcome must be rejected")
	}
}

// TestMutationAsPerturbation 验证变异体实现扰动接口（双结合统一语义）。
func TestMutationAsPerturbation(t *testing.T) {
	var p Perturbation = Mutation{
		File: "core/bp.go", Line: 10, Col: 5,
		Class: OpArithmetic, From: "+", To: "-",
	}
	if p.Kind() != PerturbStatic {
		t.Errorf("mutation kind = %q, want static", p.Kind())
	}
	if p.ID() == "" || p.Describe() == "" {
		t.Error("mutation ID/Describe must be non-empty")
	}
}
