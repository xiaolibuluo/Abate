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

import "sort"

// 维度权重（万分位，CES-4.2）。
// Dimension weights in basis points (CES-4.2).
const (
	weightAR  BP = 3000
	weightREC BP = 2000
	weightDC  BP = 2000
	weightDI  BP = 2000
	weightBRC BP = 1000

	// diCapOnDataLoss 是 CES-4.4.4 硬规则的 DI 封顶值（2000‱ = 20.00 分）。
	diCapOnDataLoss BP = 2000
	// coverageBase 与 coverageSpan 定义修正系数 k 的下界与可变区间（CES-4.5）。
	coverageBase BP = 8000
	coverageSpan BP = 2000
)

// Dimension 标识五个评分维度。
// Dimension identifies one of the five scoring dimensions.
type Dimension string

const (
	DimAR  Dimension = "AR"
	DimREC Dimension = "REC"
	DimDC  Dimension = "DC"
	DimDI  Dimension = "DI"
	DimBRC Dimension = "BRC"
)

// Dimensions 返回固定顺序的维度列表（确定性输出）。
// Dimensions returns the dimensions in a fixed order.
func Dimensions() []Dimension {
	return []Dimension{DimAR, DimREC, DimDC, DimDI, DimBRC}
}

// Score 是一次评分运行的完整结果（CES-4.7.2 报告必备字段）。
// Score is the complete result of one scoring run (CES-4.7.2 report fields).
type Score struct {
	// Final 修正后总分（万分位；展示时 ÷100 得 0.00–100.00，CES-4.7.1）。
	Final BP `json:"final"`
	// Raw 覆盖修正前的加权总分（万分位，CES-4.2.1）。
	Raw BP `json:"raw"`
	// Dims 各维度原值（万分位）。
	Dims map[Dimension]BP `json:"dims"`
	// Coverage 覆盖矩阵达成率 C（万分位）。
	Coverage BP `json:"coverage"`
	// Factor 覆盖修正系数 k（万分位）。
	Factor BP `json:"factor"`
	// Unmeasured 缺测维度清单（CES-4.6.1，按规范固定顺序输出）。
	Unmeasured []Dimension `json:"unmeasured,omitempty"`
}

// Display 返回展示分（百分制字符串由调用方格式化；此处返回整数部分与两位小数）。
// Display returns the integer and two-decimal parts of the display score.
func (s Score) Display() (int64, int64) {
	return int64(s.Final) / 100, int64(s.Final) % 100
}

// CoverageInput 描述覆盖矩阵的达成情况（CES-4.5）。
// CoverageInput describes coverage-matrix attainment (CES-4.5).
type CoverageInput struct {
	// Exercised 已演练组合数（含 ChaosHub TOP 20 中已演练项）。
	Exercised int64
	// Total 覆盖矩阵总组合数。
	Total int64
}

// Coverage 计算达成率 C 与修正系数 k（CES-4.5）。
// Coverage computes attainment C and correction factor k (CES-4.5).
func (c CoverageInput) Coverage() (BP, BP) {
	cov := Ratio(c.Exercised, c.Total)
	k := coverageBase + coverageSpan.MulDiv(cov.Min(BPTarget), BPTarget)
	return cov, k
}

// Scorer 是第 4 章评分算法的无状态实现；全部计算为万分位定点整数（CES-4.1.2）。
// Scorer is the stateless Chapter 4 scoring algorithm, basis-point fixed-point
// integer arithmetic only (CES-4.1.2).
type Scorer struct{}

// NewScorer 构造评分器。
// NewScorer constructs a scorer.
func NewScorer() Scorer { return Scorer{} }

// Score 对实验记录集合执行完整评分（CES 第 4 章）。
// 输入记录先按实验 ID 字典序升序排序（CES-4.1.3），再逐维度加权聚合。
//
// Score runs the full scoring algorithm over experiment records.
// Records are sorted ascending by experiment ID first (CES-4.1.3).
func (Scorer) Score(records []Record, cov CoverageInput) (Score, error) {
	sorted := make([]Record, len(records))
	copy(sorted, records)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Exp.ID < sorted[j].Exp.ID
	})

	var arW, arV, recW, recV, dcW, dcV, diW, diV, brcW, brcV []BP
	dataLoss := false

	for _, r := range sorted {
		w, err := r.Weight()
		if err != nil {
			return Score{}, err
		}
		arW = append(arW, w)
		arV = append(arV, r.HoldRatio())
		if s, ok := r.RecoveryScore(); ok {
			recW = append(recW, w)
			recV = append(recV, s)
		}
		dcW = append(dcW, w)
		dcV = append(dcV, r.DegradationScore())
		diW = append(diW, w)
		diV = append(diV, r.IntegrityScore())
		brcW = append(brcW, w)
		brcV = append(brcV, r.RadiusScore())
		dataLoss = dataLoss || r.DataLoss
	}

	ar, arOK := Weighted(arW, arV)
	rec, recOK := Weighted(recW, recV)
	dc, dcOK := Weighted(dcW, dcV)
	di, diOK := Weighted(diW, diV)
	brc, brcOK := Weighted(brcW, brcV)
	if dataLoss {
		di = di.Min(diCapOnDataLoss)
	}

	dims := map[Dimension]BP{
		DimAR: ar, DimREC: rec, DimDC: dc, DimDI: di, DimBRC: brc,
	}
	var unmeasured []Dimension
	for i, ok := range []bool{arOK, recOK, dcOK, diOK, brcOK} {
		if !ok {
			d := Dimensions()[i]
			dims[d] = 0
			unmeasured = append(unmeasured, d)
		}
	}

	raw := (ar*weightAR + rec*weightREC + dc*weightDC +
		di*weightDI + brc*weightBRC) / BPFull
	covC, k := cov.Coverage()
	final := raw.MulDiv(k, BPFull)

	return Score{
		Final: final, Raw: raw, Dims: dims,
		Coverage: covC, Factor: k, Unmeasured: unmeasured,
	}, nil
}
