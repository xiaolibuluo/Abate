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

// goldenCase 是附录 C 金标准数据集的一个用例（CES-4.8.2）。
// goldenCase is one case of the Appendix C golden dataset (CES-4.8.2).
type goldenCase struct {
	name     string
	records  []Record
	cov      CoverageInput
	wantDims map[Dimension]BP
	wantRaw  BP
	wantC    BP
	wantK    BP
	wantFin  BP
}

// goldenRecordsGS1 构造 GS-1 输入：单实验满覆盖。
// goldenRecordsGS1 builds the GS-1 input: one experiment, full coverage.
func goldenRecordsGS1() []Record {
	return []Record{{
		Exp: Experiment{
			ID: "exp-a", Criticality: 4, RecoveryBudgetMs: 300000,
		},
		Evidence:       EvidenceA,
		HoldMs:         90000,
		WindowMs:       100000,
		RecoveryMs:     60000,
		Degraded:       true,
		DegradedPass:   90,
		DegradedTotal:  100,
		InvariantPass:  9,
		InvariantTotal: 10,
		DeclaredRadius: 10,
		ObservedRadius: 8,
	}}
}

// goldenRecordsGS2 构造 GS-2 输入：双实验，数据丢失封顶 + 部分覆盖。
// goldenRecordsGS2 builds the GS-2 input: two experiments, data-loss cap
// and partial coverage.
func goldenRecordsGS2() []Record {
	return []Record{
		{
			Exp: Experiment{ID: "exp-a", Criticality: 5},
			Evidence:       EvidenceB,
			HoldMs:         60000,
			WindowMs:       100000,
			RecoveryMs:     -1,
			InvariantPass:  10,
			InvariantTotal: 10,
			DataLoss:       true,
			DeclaredRadius: 5,
		},
		{
			Exp: Experiment{ID: "exp-b", Criticality: 3},
			Evidence:       EvidenceC,
			HoldMs:         100000,
			WindowMs:       100000,
			RecoveryMs:     90000,
			Degraded:       true,
			DegradedPass:   100,
			DegradedTotal:  100,
			InvariantPass:  8,
			InvariantTotal: 10,
			DeclaredRadius: 10,
			ObservedRadius: 20,
		},
	}
}

// TestGoldenDataset 逐值复算附录 C 金标准数据集（CES-4.8.2）。
// TestGoldenDataset reproduces the Appendix C golden dataset value by value.
func TestGoldenDataset(t *testing.T) {
	cases := []goldenCase{
		{
			name:    "GS-1 单实验满覆盖",
			records: goldenRecordsGS1(),
			cov:     CoverageInput{Exercised: 95, Total: 100},
			wantDims: map[Dimension]BP{
				DimAR: 9000, DimREC: 8000, DimDC: 9000, DimDI: 9000, DimBRC: 10000,
			},
			wantRaw: 8900, wantC: 9500, wantK: 10000, wantFin: 8900,
		},
		{
			name:    "GS-2 双实验数据丢失封顶",
			records: goldenRecordsGS2(),
			cov:     CoverageInput{Exercised: 47, Total: 100},
			wantDims: map[Dimension]BP{
				DimAR: 7090, DimREC: 7000, DimDC: 2727, DimDI: 2000, DimBRC: 8636,
			},
			wantRaw: 5336, wantC: 4700, wantK: 8989, wantFin: 4796,
		},
		{
			name:    "GS-3 零覆盖下限",
			records: goldenRecordsGS1(),
			cov:     CoverageInput{Exercised: 0, Total: 100},
			wantDims: map[Dimension]BP{
				DimAR: 9000, DimREC: 8000, DimDC: 9000, DimDI: 9000, DimBRC: 10000,
			},
			wantRaw: 8900, wantC: 0, wantK: 8000, wantFin: 7120,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewScorer().Score(tc.records, tc.cov)
			if err != nil {
				t.Fatalf("score: %v", err)
			}
			if s.Raw != tc.wantRaw {
				t.Errorf("raw = %d, want %d", s.Raw, tc.wantRaw)
			}
			if s.Coverage != tc.wantC {
				t.Errorf("C = %d, want %d", s.Coverage, tc.wantC)
			}
			if s.Factor != tc.wantK {
				t.Errorf("k = %d, want %d", s.Factor, tc.wantK)
			}
			if s.Final != tc.wantFin {
				t.Errorf("final = %d, want %d", s.Final, tc.wantFin)
			}
			for _, d := range Dimensions() {
				if s.Dims[d] != tc.wantDims[d] {
					t.Errorf("dim %s = %d, want %d", d, s.Dims[d], tc.wantDims[d])
				}
			}
		})
	}
}
