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

// PITVersion is the PITest release that produced the reference report.
const PITVersion = "1.19.6"

// PITCommand is the command that produced the reference report.
const PITCommand = "mvn -q org.pitest:pitest-maven:1.19.6:mutationCoverage " +
	"-Dmutators=DEFAULTS (cwd: kernel/mutation/testdata/pit/baseline)"

// reportBufferCap is the initial capacity (16 KiB) of the Markdown render
// buffer; the document is small and deterministic, so a single upfront
// allocation avoids repeated growth.
const reportBufferCap = 16384

// Report is the aggregated point by point parity result.
//
// The report is deterministic: it holds no timestamp and no map iteration
// order, so two runs over the same reference report are byte identical.
type Report struct {
	// PITVersion is the PITest release of the reference report.
	PITVersion string `json:"pitVersion"`

	// Command is the command that produced the reference report.
	Command string `json:"command"`

	// Total is the number of PIT mutations in the reference report.
	Total int `json:"total"`

	// Comparable is the number of mutations taking part in the rate.
	Comparable int `json:"comparable"`

	// Consistent is the number of comparable mutations that agreed.
	Consistent int `json:"consistent"`

	// Rate is the consistency rate over the comparable mutations.
	Rate string `json:"rate"`

	// Outcomes lists every mutation in reference report order.
	Outcomes []Outcome `json:"outcomes"`

	// AbateOnly lists the Abate nodes no PIT mutation maps to.
	AbateOnly []Outcome `json:"abateOnly"`
}

// JSON renders the parity report as canonical machine readable JSON.
func (r Report) JSON() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// Markdown renders the parity report as the human readable review document.
//
// The document states the mapping rules, the per point detail, the excluded
// PIT specific operators and the honest boundaries of the comparison.
func (r Report) Markdown() []byte {
	doc := make([]byte, 0, reportBufferCap)
	doc = append(doc, "<!--\nCopyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.\n"+
		"Abate is dual-licensed under AGPLv3 and a commercial license.\n"+
		"For commercial licensing, contact xiaolibuluo@163.com.\n"+
		"Source: https://github.com/xiaolibuluo/Abate\n-->\n\n"...)
	doc = append(doc, "# PIT 对拍报告（KRN-MUT 变异核 ↔ PITest "+r.PITVersion+"）\n\n"...)
	doc = append(doc, "> 本报告由 `scripts/pit-parity.sh` 离线复跑生成，内容全部来自入库工件与\n"...)
	doc = append(doc, "> Abate 变异核的真实引擎输出，不含时间戳，逐字节可复现。\n\n"...)
	doc = append(doc, "## 1. 结论\n\n"...)
	doc = append(doc, "| 指标 | 数值 |\n|---|---|\n"...)
	doc = append(doc, "| PIT 版本 | "+r.PITVersion+" |\n"...)
	doc = append(doc, "| PIT 变异总数 | "+strconv.Itoa(r.Total)+" |\n"...)
	doc = append(doc, "| 可比对点数 | "+strconv.Itoa(r.Comparable)+" |\n"...)
	doc = append(doc, "| 一致点数 | "+strconv.Itoa(r.Consistent)+" |\n"...)
	doc = append(doc, "| 一致率 | "+r.Rate+" |\n"...)
	doc = append(doc, "| 排除点数（PIT 专属算子） | "+strconv.Itoa(r.excludedCount())+" |\n\n"...)
	doc = append(doc, "PIT 参考产物生成命令：`"+r.Command+"`\n\n"...)
	doc = append(doc, "## 2. 对拍口径\n\n"...)
	doc = append(doc, "三方独立判定逐点互校，任一方不一致即记为差异并定位根因：\n\n"...)
	doc = append(doc, "1. **PIT 官方状态**：入库工件 `kernel/mutation/testdata/pit/pit-mutations.json` "...)
	doc = append(doc, "中 PITest "+r.PITVersion+" 的真实逐点结论。\n"...)
	doc = append(doc, "2. **Abate 引擎判定**：`kernel/mutation` 的真实引擎 "...)
	doc = append(doc, "（`NewRegistryWith` + `NewEngineWithRegistry` + `Engine.Run`）在同构夹具上产出的判定。\n"...)
	doc = append(doc, "3. **语义见证**：按 Java 基准工程语义独立判定该变异在套件内是否可区分，"...)
	doc = append(doc, "该判定不看 PIT 状态，用于反查映射表是否被结果反向拟合。\n\n"...)
	doc = append(doc, "状态映射口径：\n\n"...)
	doc = append(doc, "- PIT `KILLED` ↔ Abate `KILLED`；\n"...)
	doc = append(doc, "- PIT `SURVIVED` 且语义见证判定为套件内不可区分 ↔ Abate `EQUIVALENT`"...)
	doc = append(doc, "（Abate 夹具用 `h-equivalent` 节点显式承载语义等价变异）；\n"...)
	doc = append(doc, "- PIT `NO_COVERAGE` ↔ Abate 覆盖计数为 0 且未被杀死，即 `SURVIVED`"...)
	doc = append(doc, "（Abate 夹具用 `j-survived` 节点显式承载未覆盖变异）。\n\n"...)
	doc = append(doc, "## 3. 逐点明细\n\n"...)
	doc = append(doc, "表中 # 为参考报告中的行序号，与第 4 节的排除项编号一致。\n\n"...)
	doc = append(doc, "| # | 方法 | 行 | PIT 算子 | PIT 状态 | Abate 节点 | Abate 算子类 | "...)
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
			" | "+outcome.PITStatus+
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
	doc = append(doc, "\n## 4. 被排除的 PIT 专属算子\n\n"...)
	doc = append(doc, "下列变异**不参与一致率计算**，逐条给出排除原因，不静默丢弃、不凑数。\n\n"...)
	doc = append(doc, "| # | 方法 | 行 | PIT 算子 | PIT 状态 | 语义见证 | 排除原因 |\n"...)
	doc = append(doc, "|---|---|---|---|---|---|---|\n"...)
	for i, outcome := range r.Outcomes {
		if outcome.Disposition != DispositionExcluded {
			continue
		}
		doc = append(doc, "| "+strconv.Itoa(i+1)+
			" | "+outcome.Method+
			" | "+strconv.Itoa(outcome.Line)+
			" | "+outcome.Mutator+
			" | "+outcome.PITStatus+
			" | "+outcome.Witness+
			" | "+outcome.Reason+" |\n"...)
	}
	doc = append(doc, "\n## 5. Abate 侧无 PIT 对应点的节点\n\n"...)
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
	doc = append(doc, "1. **未启动 JVM**：本轮对拍不回放 PIT 进程，而是使用入库的官方实跑产物"...)
	doc = append(doc, "（`pit-mutations.json` / `pit-mutations.xml` / `pit-console-report.txt`）；"...)
	doc = append(doc, "`scripts/pit-parity.sh --rerun-pit` 才会真实重跑 PITest。\n"...)
	doc = append(doc, "2. **基准工程覆盖面有限**：PIT 基准工程只有一个类、三个测试、"...)
	doc = append(doc, "十个方法，28 条变异覆盖七类算子族，不等于 PITest 全量算子族。\n"...)
	doc = append(doc, "3. **映射表非一一覆盖**：PIT 的 7 条变异落在 Abate 七类算子没有语义的算子族上"...)
	doc = append(doc, "（顺序判定删除、字符串返回置空），已逐条登记为排除项。\n"...)
	doc = append(doc, "4. **语义等价的口径**：PITest 不识别语义等价变异，只报告 `SURVIVED`；"...)
	doc = append(doc, "本报告的等价结论由语义见证给出，并在第 2 节写明映射口径。\n"...)
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
