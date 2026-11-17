/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"encoding/json"
	"strconv"
)

// reportBufferCap is the initial capacity (16 KiB) of the Markdown render
// buffer; the document is small and deterministic, so a single upfront
// allocation avoids repeated growth.
const reportBufferCap = 16384

// Report is the aggregated self-developed golden regression result.
//
// The report is deterministic: it holds no timestamp and no map iteration
// order, so two runs over the same golden are byte identical.
type Report struct {
	// Source identifies the self-developed regression source.
	Source string `json:"source"`

	// Seed is the deterministic seed of the Abate engine run.
	Seed int `json:"seed"`

	// Total is the number of cases in the frozen golden.
	Total int `json:"total"`

	// Comparable is the number of cases taking part in the rate.
	Comparable int `json:"comparable"`

	// Consistent is the number of comparable cases that agreed.
	Consistent int `json:"consistent"`

	// Rate is the consistency rate over the comparable cases.
	Rate string `json:"rate"`

	// Outcomes lists every case in golden order.
	Outcomes []Outcome `json:"outcomes"`

	// AbateOnly lists the Abate nodes no golden case maps to.
	AbateOnly []Outcome `json:"abateOnly"`
}

// JSON renders the regression report as canonical machine readable JSON.
func (r Report) JSON() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// Markdown renders the regression report as the human readable review document.
//
// The document states the self-developed source, the per case detail, the
// excluded local operators and the honest boundaries of the comparison.
func (r Report) Markdown() []byte {
	doc := make([]byte, 0, reportBufferCap)
	doc = append(doc, "<!--\nCopyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.\n"+
		"Abate is dual-licensed under AGPLv3 and a commercial license.\n"+
		"For commercial licensing, contact xiaolibuluo@163.com.\n"+
		"Source: https://github.com/xiaolibuluo/Abate\n-->\n\n"...)
	doc = append(doc, "# Abate 自研变异语义 golden 回归报告\n\n"...)
	doc = append(doc, "> 本报告由 `scripts/pit-parity.sh` 离线复跑生成，对比本仓引擎输出与冻结的\n"...)
	doc = append(doc, "> 自研 golden，非任何上游工具的官方对拍；不含时间戳，逐字节可复现。\n\n"...)
	doc = append(doc, "## 1. 结论\n\n"...)
	doc = append(doc, "| 指标 | 数值 |\n|---|---|\n"...)
	doc = append(doc, "| 数据源 | 自研 golden 语义回归（非上游官方对拍） |\n"...)
	doc = append(doc, "| 引擎确定性种子 | "+strconv.Itoa(r.Seed)+" |\n"...)
	doc = append(doc, "| golden 判例总数 | "+strconv.Itoa(r.Total)+" |\n"...)
	doc = append(doc, "| 可比对点数 | "+strconv.Itoa(r.Comparable)+" |\n"...)
	doc = append(doc, "| 一致点数 | "+strconv.Itoa(r.Consistent)+" |\n"...)
	doc = append(doc, "| 一致率 | "+r.Rate+" |\n"...)
	doc = append(doc, "| 排除点数（无 Abate 对应语义） | "+strconv.Itoa(r.excludedCount())+" |\n"...)
	doc = append(doc, "| 竞品对拍分数 | NOT_MEASURED（已删除全部第三方对拍产物） |\n\n"...)
	doc = append(doc, "冻结 golden 由 `kernel/mutation/pitparity/gen_golden.go` 用本仓引擎对同构夹具生成，\n"...)
	doc = append(doc, "权威数据源为 `"+GoldenRelativePath+"`。\n\n"...)
	doc = append(doc, "## 2. 回归口径\n\n"...)
	doc = append(doc, "三方独立判定逐点互校，任一方不一致即记为差异并定位根因：\n\n"...)
	doc = append(doc, "1. **冻结期望分类**：自研 golden 中由语义见证派生的判定，非外部引擎实跑结果；\n"...)
	doc = append(doc, "2. **Abate 引擎判定**：`kernel/mutation` 的真实引擎 "...)
	doc = append(doc, "（`NewRegistryWith` + `NewEngineWithRegistry` + `Engine.Run`）在同构夹具上产出的判定。\n"...)
	doc = append(doc, "3. **语义见证**：按自研基线语义独立判定该变异在套件内是否可区分，"...)
	doc = append(doc, "该判定不看冻结期望分类，用于反查映射表是否被结果反向拟合。\n\n"...)
	doc = append(doc, "状态映射口径：\n\n"...)
	doc = append(doc, "- 期望 `KILLED` ↔ Abate `KILLED`；\n"...)
	doc = append(doc, "- 期望 `SURVIVED` 且语义见证判定为套件内不可区分 ↔ Abate `EQUIVALENT`"...)
	doc = append(doc, "（自研夹具用 `h-equivalent` 节点显式承载语义等价变异）；\n"...)
	doc = append(doc, "- 期望 `NO_COVERAGE` ↔ Abate 覆盖计数为 0 且未被杀死，即 `SURVIVED`"...)
	doc = append(doc, "（自研夹具用 `j-survived` 节点显式承载未覆盖变异）。\n\n"...)
	doc = append(doc, "## 3. 逐点明细\n\n"...)
	doc = append(doc, "表中 # 为 golden 中的判例序号，与第 4 节的排除项编号一致。\n\n"...)
	doc = append(doc, "| # | 方法 | 行 | 算子标识 | 期望分类 | Abate 节点 | Abate 算子类 | "...)
	doc = append(doc, "Abate 判定 | 语义见证 | 比对口径 | 一致 |\n"...)
	doc = append(doc, "|---|---|---|---|---|---|---|---|---|---|---|\n"...)
	for i, outcome := range r.Outcomes {
		if outcome.Disposition != DispositionComparable {
			continue
		}
		consistent := markerInconsistent
		if outcome.Consistent {
			consistent = markerConsistent
		}
		doc = append(doc, "| "+strconv.Itoa(i+1)+
			" | "+outcome.Method+
			" | "+strconv.Itoa(outcome.Line)+
			" | "+outcome.Mutator+
			" | "+outcome.Expect+
			" | "+outcome.Node+
			" | "+outcome.Class+
			" | "+outcome.AbateVerdict+
			" | "+outcome.Witness+
			" | "+outcome.VerdictMode+
			" | "+consistent+" |\n"...)
	}
	doc = append(doc, "\n映射说明与差异明细：\n\n"...)
	for i, outcome := range r.Outcomes {
		if outcome.Disposition != DispositionComparable {
			continue
		}
		doc = append(doc, "- **#"+strconv.Itoa(i+1)+" "+outcome.Method+":"+
			strconv.Itoa(outcome.Line)+" "+outcome.Mutator+"** — "+
			outcome.Reason...)
		if outcome.Difference != "" {
			doc = append(doc, " 差异："+outcome.Difference...)
		}
		doc = append(doc, "\n"...)
	}
	doc = append(doc, "\n## 4. 被排除的无对应语义判例\n\n"...)
	doc = append(doc, "下列变异**不参与一致率计算**，逐条给出排除原因，不静默丢弃、不凑数。\n\n"...)
	doc = append(doc, "| # | 方法 | 行 | 算子标识 | 期望分类 | 语义见证 | 排除原因 |\n"...)
	doc = append(doc, "|---|---|---|---|---|---|---|\n"...)
	for i, outcome := range r.Outcomes {
		if outcome.Disposition != DispositionExcluded {
			continue
		}
		doc = append(doc, "| "+strconv.Itoa(i+1)+
			" | "+outcome.Method+
			" | "+strconv.Itoa(outcome.Line)+
			" | "+outcome.Mutator+
			" | "+outcome.Expect+
			" | "+outcome.Witness+
			" | "+outcome.Reason+" |\n"...)
	}
	doc = append(doc, "\n## 5. Abate 侧无 golden 对应点的节点\n\n"...)
	if len(r.AbateOnly) == 0 {
		doc = append(doc, "无。\n\n"...)
	} else {
		doc = append(doc, "| Abate 节点 | Abate 判定 | 说明 |\n|---|---|---|\n"...)
		for _, node := range r.AbateOnly {
			doc = append(doc, "| "+node.Node+" | "+node.AbateVerdict+" | "+
				node.Reason+" |\n"...)
		}
		doc = append(doc, "\n"...)
	}
	doc = append(doc, "## 6. 如实边界声明\n\n"...)
	doc = append(doc, "1. **未启动任何外部引擎**：本轮不回放第三方变异工具，也不读取任何上游产物，"...)
	doc = append(doc, "只用本仓引擎输出与冻结自研 golden 做逐点回归；\n"...)
	doc = append(doc, "2. **竞品对拍分数 NOT_MEASURED**：上游产物已彻底删除，不再宣称任何官方对拍一致率；\n"...)
	doc = append(doc, "3. **映射表非一一覆盖**：7 条判例落在 Abate 七类算子没有语义的构造上"...)
	doc = append(doc, "（顺序判定删除、字符串返回置空），已逐条登记为排除项；\n"...)
	doc = append(doc, "4. **语义等价由见证给出**：本仓引擎识别语义等价变异，等价结论由独立见证交叉校验；\n"...)
	doc = append(doc, "5. **夹具同构**：Abate 侧夹具与 `kernel/mutation/helpers_test.go` 的夹具"...)
	doc = append(doc, "同构（十节点、七类算子、等价/超时/未覆盖三条支路），两侧夹具若发生漂移，"...)
	doc = append(doc, "两侧各自的单元测试会先行失败。\n"...)
	return doc
}

// excludedCount returns the number of excluded points.
func (r Report) excludedCount() int {
	count := 0
	for _, outcome := range r.Outcomes {
		if outcome.Disposition == DispositionExcluded {
			count++
		}
	}
	return count
}

// markerConsistent renders a consistent (passing) verdict as a review marker.
const markerConsistent = "是"

// markerInconsistent renders an inconsistent verdict as a review marker.
const markerInconsistent = "否"
