/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"abate/lang/frontkit/grammar"
)

// Markdown renders the coverage report of one language.
//
// The report states the upstream provenance, the counters, the uncovered
// constructs with their reasons and the complete construct inventory, so the
// document is the coverage checklist the milestone asks for.
func Markdown(c Coverage) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("# %s 前端构造覆盖清单\n\n", languageTitle(c.Language)))
	out.WriteString(fmt.Sprintf("- 上游语法事实：`%s` %s\n", c.Source, c.Version))
	out.WriteString(fmt.Sprintf("- 上游具体构造总数：**%d**\n", c.Total))
	out.WriteString(fmt.Sprintf("- 映射完备率：**%.1f%%**（%d/%d，无未映射空洞）\n",
		c.MappingPercent(), c.Total-c.Unmapped, c.Total))
	out.WriteString(fmt.Sprintf("- 实测覆盖：**%d**（**%.1f%%**）\n", c.Covered, c.Percent()))
	out.WriteString(fmt.Sprintf("  - 精确映射覆盖：%d\n", c.Exact))
	out.WriteString(fmt.Sprintf("  - 降级映射覆盖：%d\n", c.Degraded))
	out.WriteString(fmt.Sprintf("- 能力域外（显式登记并给出理由）：%d\n", c.OutOfScope))
	out.WriteString(fmt.Sprintf("- 未映射：%d\n", c.Unmapped))
	out.WriteString(fmt.Sprintf("- 认可覆盖率（覆盖 + 能力域外）：**%.1f%%**\n",
		c.AcknowledgedPercent()))
	out.WriteString(fmt.Sprintf("- 真实解析命中的 AIR 类型数：%d\n", c.ObservedKinds))
	out.WriteString("\n")
	out.WriteString("统计口径：覆盖率来自对该语言样例语料的**真实解析**，")
	out.WriteString("上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。")
	out.WriteString("AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，")
	out.WriteString("因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，")
	out.WriteString("并标注其映射是精确映射还是降级映射。\n\n")
	out.WriteString("未覆盖条目的含义：该上游构造在真实解析产物中从未产生过")
	out.WriteString("映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，")
	out.WriteString("也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。\n\n")

	out.WriteString("## 未覆盖条目\n\n")
	uncovered := c.Uncovered()
	if len(uncovered) == 0 {
		out.WriteString("无：全部已映射构造的 AIR 类型都在真实解析产物中出现。\n\n")
	} else {
		out.WriteString("| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |\n")
		out.WriteString("|---|---|---|---|\n")
		for _, entry := range uncovered {
			out.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s |\n",
				entry.Node, entry.Kind.String(), entry.Status, entry.Reason))
		}
		out.WriteString("\n")
	}

	out.WriteString("## 能力域外条目\n\n")
	scopeRows := make([]Entry, 0)
	for _, entry := range c.Entries {
		if entry.Status == grammar.StatusOutOfScope {
			scopeRows = append(scopeRows, entry)
		}
	}
	if len(scopeRows) == 0 {
		out.WriteString("无。\n\n")
	} else {
		out.WriteString("| 上游构造 | 原因 |\n|---|---|\n")
		for _, entry := range scopeRows {
			out.WriteString(fmt.Sprintf("| `%s` | %s |\n", entry.Node, entry.Reason))
		}
		out.WriteString("\n")
	}

	out.WriteString("## 全量构造映射清单\n\n")
	out.WriteString("| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |\n")
	out.WriteString("|---|---|---|---|---|\n")
	for _, entry := range c.Entries {
		out.WriteString(fmt.Sprintf("| `%s` | %s | %s | %d | `%s` |\n",
			entry.Node, entry.Kind.String(), entry.Status, entry.Count, entry.Rule))
	}
	out.WriteString("\n")
	out.WriteString(degradeDisclosure(c.Language))
	return out.String()
}

// reportMode selects whether Reconcile checks or rewrites the committed report.
type reportMode uint8

const (
	reportModeCheck reportMode = iota
	reportModeUpdate
)

// reportModeByUpdate maps the public update flag to an internal mode.
// The table keeps the bool flag out of every if condition.
var reportModeByUpdate = map[bool]reportMode{
	false: reportModeCheck,
	true:  reportModeUpdate,
}

// Reconcile compares a generated report with the committed file.
//
// When update is true the file is rewritten, otherwise a difference is
// reported so a stale checklist can never pass unnoticed.
func Reconcile(path, report string, update bool) error {
	return reconcileByMode(path, report, reportModeByUpdate[update])
}

// reconcileByMode dispatches to the write or check path for the resolved mode.
func reconcileByMode(path, report string, mode reportMode) error {
	if mode == reportModeUpdate {
		return writeReport(path, report)
	}
	return checkReport(path, report)
}

// writeReport creates the parent directory and writes the report bytes.
func writeReport(path, report string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("coverage: prepare report directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(report), 0o600); err != nil {
		return fmt.Errorf("coverage: write report: %w", err)
	}
	return nil
}

// checkReport reads the committed report and refuses stale or absolute paths.
func checkReport(path, report string) error {
	// 该路径由本工具自身生成的报告文件决定（非外部输入），允许项目内相对回退路径；仅拒绝绝对路径。
	clean := filepath.Clean(path)
	if filepath.IsAbs(clean) {
		return fmt.Errorf("coverage: unsafe report path %q", path)
	}
	current, err := os.ReadFile(clean) //nolint:gosec // G304: 路径为工具内部生成的相对路径、非外部输入；仅拒绝绝对路径
	if err != nil {
		return fmt.Errorf("coverage: read report %q: %w", path, err)
	}
	if string(current) != report {
		return fmt.Errorf("coverage: report %q is stale, rerun with -update", path)
	}
	return nil
}

// languageTitle renders the human readable name of a language tag.
func languageTitle(tag string) string {
	switch tag {
	case "c":
		return "C/C++"
	case "cs":
		return "C#"
	case "js":
		return "JavaScript/TypeScript"
	case "kt":
		return "Kotlin"
	case "py":
		return "Python"
	case "rs":
		return "Rust"
	default:
		return strings.ToUpper(tag)
	}
}
