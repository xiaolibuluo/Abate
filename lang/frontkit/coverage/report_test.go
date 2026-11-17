/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// AssertContainsAll fails when a rendered report lacks any expected fragment.
func assertContainsAll(t *testing.T, report string, fragments ...string) {
	t.Helper()
	for _, fragment := range fragments {
		if !strings.Contains(report, fragment) {
			t.Fatalf("report is missing %q\n%s", fragment, report)
		}
	}
}

// AssertErrorContains fails when an error is missing or lacks the expected fragment.
func assertErrorContains(t *testing.T, err error, fragment string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected an error containing %q, got nil", fragment)
	}
	if !strings.Contains(err.Error(), fragment) {
		t.Fatalf("error = %q, want a fragment %q", err.Error(), fragment)
	}
}

// LanguageTitleCases returns the tag and rendered title of every switch branch.
func languageTitleCases() []struct {
	name  string
	tag   string
	title string
} {
	return []struct {
		name  string
		tag   string
		title string
	}{
		{"c", "c", "C/C++"},
		{"cs", "cs", "C#"},
		{"js", "js", "JavaScript/TypeScript"},
		{"kt", "kt", "Kotlin"},
		{"py", "py", "Python"},
		{"rs", "rs", "Rust"},
		{"unknown", "swift", "SWIFT"},
		{"empty", "", ""},
		{"mixed-case", "Go", "GO"},
	}
}

// TestLanguageTitleRendersEveryBranch proves all six tags and the uppercase default.
func TestLanguageTitleRendersEveryBranch(t *testing.T) {
	t.Parallel()
	for _, tc := range languageTitleCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := languageTitle(tc.tag); got != tc.title {
				t.Fatalf("languageTitle(%q) = %q, want %q", tc.tag, got, tc.title)
			}
		})
	}
}

// TestMarkdownUsesLanguageTitle proves the heading renders every language label.
func TestMarkdownUsesLanguageTitle(t *testing.T) {
	t.Parallel()
	for _, tc := range languageTitleCases() {
		t.Run(tc.name, func(t *testing.T) {
			report := Markdown(Coverage{Language: tc.tag})
			if !strings.HasPrefix(report, "# "+tc.title+" 前端构造覆盖清单\n\n") {
				t.Fatalf("heading = %q, want the %q title", strings.SplitN(report, "\n", 2)[0], tc.title)
			}
		})
	}
}

// TestMarkdownWithoutGapsRendersEmptySections proves the no-uncovered and no-scope branches.
func TestMarkdownWithoutGapsRendersEmptySections(t *testing.T) {
	t.Parallel()
	data := Coverage{
		Language: "rs", Version: "0.21.0", Source: "tree-sitter-rust",
		Total: 3, Covered: 3, Exact: 2, Degraded: 1, ObservedKinds: 2,
		Entries: []Entry{
			{Node: "function_item", Kind: air.KindFuncDecl, Status: grammar.StatusExact, Observed: true, Count: 2, Rule: "rule.fn"},
			{Node: "call_expression", Kind: air.KindCall, Status: grammar.StatusExact, Observed: true, Count: 5, Rule: "rule.call"},
			{Node: "closure_expression", Kind: air.KindFuncLit, Status: grammar.StatusDegraded, Observed: true, Count: 1, Rule: "rule.lit"},
		},
	}
	report := Markdown(data)
	assertContainsAll(t, report,
		"# Rust 前端构造覆盖清单",
		"- 上游语法事实：`tree-sitter-rust` 0.21.0",
		"- 上游具体构造总数：**3**",
		"- 映射完备率：**100.0%**（3/3，无未映射空洞）",
		"- 实测覆盖：**3**（**100.0%**）",
		"  - 精确映射覆盖：2",
		"  - 降级映射覆盖：1",
		"- 能力域外（显式登记并给出理由）：0",
		"- 未映射：0",
		"- 认可覆盖率（覆盖 + 能力域外）：**100.0%**",
		"- 真实解析命中的 AIR 类型数：2",
		"统计口径：覆盖率来自对该语言样例语料的**真实解析**",
		"未覆盖条目的含义：",
		"## 未覆盖条目\n\n无：全部已映射构造的 AIR 类型都在真实解析产物中出现。",
		"## 能力域外条目\n\n无。",
		"## 全量构造映射清单",
		"| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |",
		"| `function_item` | FuncDecl | exact | 2 | `rule.fn` |",
		"| `closure_expression` | FuncLit | degraded | 1 | `rule.lit` |",
	)
}

// TestMarkdownWithGapsRendersPopulatedSections proves the uncovered and scope tables plus inventory rows.
func TestMarkdownWithGapsRendersPopulatedSections(t *testing.T) {
	t.Parallel()
	data := Coverage{
		Language: "py", Version: "1.0", Source: "tree-sitter-python",
		Total: 4, Covered: 1, Exact: 1, OutOfScope: 1, Unmapped: 1, ObservedKinds: 1,
		Entries: []Entry{
			{Node: "function_definition", Kind: air.KindFuncDecl, Status: grammar.StatusExact, Reason: "exact-node", Observed: true, Count: 3, Rule: "rule.fn"},
			{Node: "lambda", Kind: air.KindFuncLit, Status: grammar.StatusDegraded, Reason: "no closure node", Observed: false, Rule: "rule.lambda"},
			{Node: "decorator", Kind: air.KindAttribute, Status: grammar.StatusOutOfScope, Reason: "metadata is out of scope", Observed: false, Rule: "rule.decorator"},
			{Node: "type_comment", Kind: air.KindInvalid, Status: grammar.StatusUnmapped, Reason: "no rule", Observed: false},
		},
	}
	report := Markdown(data)
	assertContainsAll(t, report,
		"# Python 前端构造覆盖清单",
		"- 映射完备率：**75.0%**（3/4，无未映射空洞）",
		"- 实测覆盖：**1**（**25.0%**）",
		"- 认可覆盖率（覆盖 + 能力域外）：**50.0%**",
		"## 未覆盖条目\n\n| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |\n|---|---|---|",
		"| `lambda` | FuncLit | degraded | no closure node |",
		"## 能力域外条目\n\n| 上游构造 | 原因 |\n|---|---|",
		"| `decorator` | metadata is out of scope |",
		"| `function_definition` | FuncDecl | exact | 3 | `rule.fn` |",
		"| `decorator` | Attribute | out-of-scope | 0 | `rule.decorator` |",
		"| `type_comment` | Invalid | unmapped | 0 |",
	)
	if strings.Contains(report, "无：全部已映射构造") {
		t.Fatal("a report with gaps must not claim there are none")
	}
	if strings.Contains(report, "## 能力域外条目\n\n无。") {
		t.Fatal("a report with scope rows must not claim there are none")
	}
}

// TestReconcileUpdateWritesNestedReport proves update mode creates directories and exact bytes.
func TestReconcileUpdateWritesNestedReport(t *testing.T) {
	t.Chdir(t.TempDir())
	target := filepath.Join("reports", "nested", "coverage.md")
	body := "# Rust coverage\n"
	if err := Reconcile(target, body, true); err != nil {
		t.Fatalf("update mode failed: %v", err)
	}
	info, err := os.Stat(filepath.Join("reports", "nested"))
	if err != nil {
		t.Fatalf("nested directory was not created: %v", err)
	}
	if !info.IsDir() {
		t.Fatal("the nested path is not a directory")
	}
	if err := Reconcile(target, body, false); err != nil {
		t.Fatalf("the written bytes do not match the report: %v", err)
	}
}

// TestReconcileUpdateOverwritesExistingReport proves a second update replaces stale content.
func TestReconcileUpdateOverwritesExistingReport(t *testing.T) {
	t.Chdir(t.TempDir())
	target := "coverage.md"
	if err := Reconcile(target, "old report\n", true); err != nil {
		t.Fatalf("first update failed: %v", err)
	}
	if err := Reconcile(target, "new report\n", true); err != nil {
		t.Fatalf("second update failed: %v", err)
	}
	if err := Reconcile(target, "new report\n", false); err != nil {
		t.Fatalf("check mode rejected the rewritten report: %v", err)
	}
	err := Reconcile(target, "old report\n", false)
	assertErrorContains(t, err, "is stale")
}

// TestReconcileCheckAcceptsMatchingReport proves an unchanged report passes the gate.
func TestReconcileCheckAcceptsMatchingReport(t *testing.T) {
	t.Chdir(t.TempDir())
	target := "coverage.md"
	if err := Reconcile(target, "byte for byte\n", true); err != nil {
		t.Fatalf("update mode failed: %v", err)
	}
	if err := Reconcile(target, "byte for byte\n", false); err != nil {
		t.Fatalf("check mode rejected a matching report: %v", err)
	}
}

// TestReconcileCheckRejectsStaleReport proves a differing report is refused.
func TestReconcileCheckRejectsStaleReport(t *testing.T) {
	t.Chdir(t.TempDir())
	target := "coverage.md"
	if err := Reconcile(target, "committed\n", true); err != nil {
		t.Fatalf("update mode failed: %v", err)
	}
	err := Reconcile(target, "generated\n", false)
	assertErrorContains(t, err, "is stale")
}

// TestReconcileCheckRejectsMissingFile proves an absent committed report is a read failure.
func TestReconcileCheckRejectsMissingFile(t *testing.T) {
	t.Chdir(t.TempDir())
	err := Reconcile("absent.md", "report\n", false)
	assertErrorContains(t, err, "read report")
}

// TestReconcileCheckRejectsUnsafePaths proves only absolute paths are refused, while relative parent paths are read.
func TestReconcileCheckRejectsUnsafePaths(t *testing.T) {
	t.Chdir(t.TempDir())
	t.Run("absolute", func(t *testing.T) {
		err := Reconcile(filepath.Join(t.TempDir(), "coverage.md"), "report\n", false)
		assertErrorContains(t, err, "unsafe report path")
	})
	relative := []struct {
		name   string
		target string
	}{
		{"parent-traversal", "../coverage.md"},
		{"deep-traversal", filepath.Join("reports", "..", "..", "coverage.md")},
		{"dot-dot-in-name", "a..b.md"},
	}
	for _, tc := range relative {
		t.Run(tc.name, func(t *testing.T) {
			err := Reconcile(tc.target, "report\n", false)
			assertErrorContains(t, err, "read report")
			if strings.Contains(err.Error(), "unsafe report path") {
				t.Fatalf("relative path %q was wrongly rejected as unsafe: %v", tc.target, err)
			}
		})
	}
}

// TestReconcileCheckAcceptsRelativeParentPath proves an in-repo parent path is read normally.
func TestReconcileCheckAcceptsRelativeParentPath(t *testing.T) {
	root := t.TempDir()
	child := filepath.Join(root, "lang", "cfront")
	if err := os.MkdirAll(child, 0o750); err != nil {
		t.Fatalf("creating the child directory failed: %v", err)
	}
	if err := Reconcile(filepath.Join(root, "coverage.md"), "committed\n", true); err != nil {
		t.Fatalf("writing the parent report failed: %v", err)
	}
	t.Chdir(child)
	if err := Reconcile(filepath.Join("..", "..", "coverage.md"), "committed\n", false); err != nil {
		t.Fatalf("a relative parent path was rejected: %v", err)
	}
}

// TestReconcileUpdateReportsDirectoryFailure proves a blocked parent directory is reported.
func TestReconcileUpdateReportsDirectoryFailure(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := Reconcile("leaf", "blocked\n", true); err != nil {
		t.Fatalf("creating the blocking file failed: %v", err)
	}
	err := Reconcile(filepath.Join("leaf", "child.md"), "report\n", true)
	assertErrorContains(t, err, "prepare report directory")
}

// TestReconcileUpdateReportsWriteFailure proves a directory target is reported.
func TestReconcileUpdateReportsWriteFailure(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := Reconcile(filepath.Join("nested", "report.md"), "report\n", true); err != nil {
		t.Fatalf("creating the nested directory failed: %v", err)
	}
	err := Reconcile("nested", "report\n", true)
	assertErrorContains(t, err, "write report")
}
