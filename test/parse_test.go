// Abate —— 止息混沌 · 多语言输出解析器测试（准层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import "testing"

// TestParseIssuesGoVet 验证 go vet 输出格式解析。
func TestParseIssuesGoVet(t *testing.T) {
	out := "main.go:12:3: printf: Printf format %d has arg wrong type string\n" +
	"util/util.go:45:5: unreachable code\n"
	issues := ParseIssues(out, "go vet", SevError)
	if len(issues) != 2 {
		t.Fatalf("want 2 issues, got %d", len(issues))
	}
	if issues[0].File != "main.go" || issues[0].Line != 12 {
		t.Fatalf("bad location: %+v", issues[0])
	}
	if issues[0].Severity != SevError {
		t.Fatalf("want error, got %s", issues[0].Severity)
	}
	if issues[1].Line != 45 {
		t.Fatalf("want line 45, got %d", issues[1].Line)
	}
}

// TestParseIssuesKotlinLint 验证 Android lint 输出格式解析。
func TestParseIssuesKotlinLint(t *testing.T) {
	out := "app/src/main/java/com/x/Main.kt:30: Error: MissingPermission [SomeId]\n" +
	"app/src/main/java/com/x/Util.kt:88: Warning: unused import\n"
	issues := ParseIssues(out, "lint", SevError)
	if len(issues) != 2 {
		t.Fatalf("want 2 issues, got %d", len(issues))
	}
	if issues[0].Severity != SevError || issues[0].Line != 30 {
		t.Fatalf("bad first issue: %+v", issues[0])
	}
	if issues[1].Severity != SevWarning || issues[1].Line != 88 {
		t.Fatalf("bad second issue: %+v", issues[1])
	}
}

// TestParseIssuesJsTs 验证 eslint/tsc 输出格式解析。
func TestParseIssuesJsTs(t *testing.T) {
	out := "src/app.js:10:5: error: 'x' is not defined\n" +
	"src/util.ts:22:1: warning: unexpected any\n"
	issues := ParseIssues(out, "eslint", SevError)
	if len(issues) != 2 {
		t.Fatalf("want 2 issues, got %d", len(issues))
	}
	if issues[0].Severity != SevError || issues[1].Severity != SevWarning {
		t.Fatalf("bad severity: %+v", issues)
	}
}

// TestParseIssuesPython 验证 pyflakes 输出格式解析。
func TestParseIssuesPython(t *testing.T) {
	out := "app.py:7:5: E999 syntax error\n" +
	"util.py:20:1: W605 invalid escape sequence\n"
	issues := ParseIssues(out, "pyflakes", SevError)
	if len(issues) != 2 {
		t.Fatalf("want 2 issues, got %d", len(issues))
	}
	if issues[0].Severity != SevError {
		t.Fatalf("E999 should be error, got %s", issues[0].Severity)
	}
	if issues[1].Severity != SevWarning {
		t.Fatalf("W605 should be warning, got %s", issues[1].Severity)
	}
}

// TestParseIssuesNoiseIgnored 验证非代码文件与无关行被忽略。
func TestParseIssuesNoiseIgnored(t *testing.T) {
	out := "some random line\n" +
	"build.gradle:5: not a code file\n" +
	"README.md:3:1: doc noise\n" +
	"main.go:0: zero line ignored\n"
	issues := ParseIssues(out, "go vet", SevError)
	if len(issues) != 0 {
		t.Fatalf("want 0 issues, got %d: %+v", len(issues), issues)
	}
}

// TestCountAndFormatIssues 验证统计与摘要格式化。
func TestCountAndFormatIssues(t *testing.T) {
	issues := []Issue{
		{File: "a.go", Line: 1, Severity: SevError, Message: "boom"},
		{File: "a.go", Line: 2, Severity: SevWarning, Message: "careful"},
		{File: "b.go", Line: 3, Severity: SevInfo, Message: "note"},
	}
	st := CountIssues(issues)
	if st.Total != 3 || st.Errors != 1 || st.Warnings != 1 || st.Infos != 1 {
		t.Fatalf("bad stats: %+v", st)
	}
	s := FormatIssues(issues, 2)
	if s == "" {
		t.Fatal("format should not be empty")
	}
}
