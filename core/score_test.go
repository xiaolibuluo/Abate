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

// goldenExperiment 构造一个合法实验（关键度 4）。
func goldenExperiment(id string) Experiment {
	return Experiment{
		ID:          id,
		Seed:        42,
		Target:      Target{Kind: TargetService, Name: "order"},
		Fault:       Fault{Class: F1Latency, Params: map[string]string{"delay": "500ms"}},
		Blast:       BlastRadius{MaxInstances: 1, MaxDurationMs: 300000},
		Steady:      []SLIAssert{{SLI: "http_success_rate", Assert: ">= 99.9%"}},
		Criticality: 4,
	}
}

// goldenRecord 构造一条满分执行记录（A 级证据）。
func goldenRecord(id string) Record {
	return Record{
		Exp:            goldenExperiment(id),
		Evidence:       EvidenceA,
		HoldMs:         300000, WindowMs: 300000,
		RecoveryMs:     0,
		Degraded:       true, DegradedPass: 1000, DegradedTotal: 1000,
		InvariantPass:  10, InvariantTotal: 10,
		DeclaredRadius: 1, ObservedRadius: 1,
	}
}

// TestScoreGolden 金标准用例（附录 C 雏形）：满分记录 + 满覆盖 ⇒ 总分 100.00。
func TestScoreGolden(t *testing.T) {
	s := NewScorer()
	got, err := s.Score([]Record{goldenRecord("exp-001")},
		CoverageInput{Exercised: 95, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	if got.Final != BPFull {
		t.Errorf("Final=%d, want 10000 (100.00)", got.Final)
	}
	if got.Raw != BPFull || got.Factor != BPFull || got.Coverage != 9500 {
		t.Errorf("Raw=%d Factor=%d Coverage=%d, want 10000/10000/9500",
			got.Raw, got.Factor, got.Coverage)
	}
	if len(got.Unmeasured) != 0 {
		t.Errorf("Unmeasured=%v, want empty", got.Unmeasured)
	}
	whole, frac := got.Display()
	if whole != 100 || frac != 0 {
		t.Errorf("Display=(%d,%d), want (100,0)", whole, frac)
	}
}

// TestScoreCoveragePenalty 验证覆盖修正：零覆盖时 k=8000，总分保留 80%（CES-4.5.1）。
func TestScoreCoveragePenalty(t *testing.T) {
	s := NewScorer()
	got, err := s.Score([]Record{goldenRecord("exp-001")},
		CoverageInput{Exercised: 0, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	if got.Factor != 8000 {
		t.Errorf("Factor=%d, want 8000", got.Factor)
	}
	if got.Final != 8000 {
		t.Errorf("Final=%d, want 8000 (80.00)", got.Final)
	}
}

// TestScoreDataLossCap 验证 DI 硬规则：不可调和数据丢失时 DI 封顶 2000（CES-4.4.4）。
func TestScoreDataLossCap(t *testing.T) {
	r := goldenRecord("exp-001")
	r.DataLoss = true
	s := NewScorer()
	got, err := s.Score([]Record{r}, CoverageInput{Exercised: 95, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	if got.Dims[DimDI] != 2000 {
		t.Errorf("DI=%d, want 2000 (capped)", got.Dims[DimDI])
	}
	// 其余维度满分 ⇒ Raw = (10000×3000 + 10000×2000 + 10000×2000 + 2000×2000 + 10000×1000) ÷ 10000 = 8400
	if got.Raw != 8400 {
		t.Errorf("Raw=%d, want 8400", got.Raw)
	}
}

// TestScoreRecoveryUnmeasured 验证无恢复事件时 REC 缺测记 0（CES-4.6.1）。
func TestScoreRecoveryUnmeasured(t *testing.T) {
	r := goldenRecord("exp-001")
	r.RecoveryMs = -1
	s := NewScorer()
	got, err := s.Score([]Record{r}, CoverageInput{Exercised: 95, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	if got.Dims[DimREC] != 0 {
		t.Errorf("REC=%d, want 0 (unmeasured)", got.Dims[DimREC])
	}
	if len(got.Unmeasured) != 1 || got.Unmeasured[0] != DimREC {
		t.Errorf("Unmeasured=%v, want [REC]", got.Unmeasured)
	}
}

// TestScoreDeterministic 验证输入乱序不影响结果（CES-4.1.3 排序保证）。
func TestScoreDeterministic(t *testing.T) {
	a, b := goldenRecord("exp-001"), goldenRecord("exp-002")
	s := NewScorer()
	r1, err := s.Score([]Record{a, b}, CoverageInput{Exercised: 50, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	r2, err := s.Score([]Record{b, a}, CoverageInput{Exercised: 50, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	if r1.Final != r2.Final || r1.Raw != r2.Raw {
		t.Errorf("nondeterministic: %+v vs %+v", r1, r2)
	}
}

// TestScoreInvalidEvidence 验证非法证据等级的错误从 Score 透传（CES-4.3.1）。
func TestScoreInvalidEvidence(t *testing.T) {
	r := goldenRecord("exp-bad")
	r.Evidence = "X"
	s := NewScorer()
	if _, err := s.Score([]Record{r}, CoverageInput{Exercised: 1, Total: 1}); err == nil {
		t.Error("invalid evidence level must propagate an error from Score")
	}
}

// TestScoreHalfHold 验证半稳态记录的 AR 折算（CES-4.4.1）。
func TestScoreHalfHold(t *testing.T) {
	r := goldenRecord("exp-001")
	r.HoldMs = 150000
	s := NewScorer()
	got, err := s.Score([]Record{r}, CoverageInput{Exercised: 95, Total: 100})
	if err != nil {
		t.Fatalf("Score: %v", err)
	}
	if got.Dims[DimAR] != 5000 {
		t.Errorf("AR=%d, want 5000", got.Dims[DimAR])
	}
}
