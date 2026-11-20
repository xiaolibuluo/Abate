/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// multilangFixtures maps each of the seven language tags to a minimal source
// snippet plus the extensions that must route to it.
type multilangFixture struct {
	tag    string
	ext    string
	source string
}

var multilangSources = []multilangFixture{
	{"go", ".go", "package p\n\nvar foo_bar int\n"},
	{"py", ".py", "def greet(name):\n    result = 1\n    return result\n"},
	{"js", ".js", "function greet(name) {\n  var result = 1;\n  return result;\n}\n"},
	{"c", ".c", "int add(int a, int b) {\n    int foo_bar = a + b;\n    return foo_bar;\n}\n"},
	{"cs", ".cs", "public class G {\n    public int Foo() { var foo_bar = 1; return foo_bar; }\n}\n"},
	{"rs", ".rs", "fn add(a: i32) -> i32 {\n    let foo_bar = a;\n    foo_bar\n}\n"},
	{"kt", ".kt", "fun add(a: Int): Int {\n    val foo_bar = a\n    return foo_bar\n}\n"},
}

// writeMultilangDir writes all seven fixtures into one directory.
func writeMultilangDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	for _, fixture := range multilangSources {
		path := filepath.Join(dir, "sample"+fixture.ext)
		if err := os.WriteFile(path, []byte(fixture.source), 0o600); err != nil {
			t.Fatalf("write %s fixture: %v", fixture.tag, err)
		}
	}
	return dir
}

// parseCheckJSON decodes a check JSON report into a generic document.
func parseCheckJSON(t *testing.T, out string) map[string]any {
	t.Helper()
	var document map[string]any
	if err := json.Unmarshal([]byte(out), &document); err != nil {
		t.Fatalf("check output is not valid JSON: %v", err)
	}
	return document
}

// TestCheckSevenLanguagesDispatch proves each language source is parsed by its
// own frontend, loads a non-empty rule set and produces findings.
func TestCheckSevenLanguagesDispatch(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	for _, fixture := range multilangSources {
		path := filepath.Join(dir, "sample"+fixture.ext)
		if err := os.WriteFile(path, []byte(fixture.source), 0o600); err != nil {
			t.Fatalf("write %s: %v", fixture.tag, err)
		}
	}
	for _, fixture := range multilangSources {
		t.Run(fixture.tag, func(t *testing.T) {
			path := filepath.Join(dir, "sample"+fixture.ext)
			code, out, errs := invoke("check", path, "--languages", fixture.tag,
				"--format", "json")
			if code != exitBusinessFailure && code != exitOK {
				t.Fatalf("%s check exit = %d (stderr=%s)", fixture.tag, code, errs)
			}
			document := parseCheckJSON(t, out)
			summary, ok := document["summary"].(map[string]any)
			if !ok {
				t.Fatalf("%s report has no summary", fixture.tag)
			}
			if summary["files"] != float64(1) {
				t.Fatalf("%s files = %v, want 1", fixture.tag, summary["files"])
			}
			rules, ok := summary["rules"].(float64)
			if !ok || rules <= 0 {
				t.Fatalf("%s loaded no rules", fixture.tag)
			}
			findings, ok := document["findings"].([]any)
			if !ok || len(findings) == 0 {
				t.Fatalf("%s produced no findings (frontend may be a stub)", fixture.tag)
			}
		})
	}
}

// TestCheckMixedDirectoryDispatch proves one invocation against a directory
// dispatches all seven languages and attributes findings to each rule family.
func TestCheckMixedDirectoryDispatch(t *testing.T) {
	t.Parallel()
	dir := writeMultilangDir(t)
	code, out, errs := invoke("check", dir, "--format", "json")
	if code != exitBusinessFailure && code != exitOK {
		t.Fatalf("mixed check exit = %d (stderr=%s)", code, errs)
	}
	document := parseCheckJSON(t, out)
	summary, ok := document["summary"].(map[string]any)
	if !ok {
		t.Fatal("report has no summary")
	}
	if summary["files"] != float64(len(multilangSources)) {
		t.Fatalf("files = %v, want %d", summary["files"], len(multilangSources))
	}
	seen := make(map[string]bool, len(multilangSources))
	findings, ok := document["findings"].([]any)
	if !ok {
		t.Fatalf("findings is not an array")
	}
	for _, raw := range findings {
		finding, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		ruleID, ok := finding["ruleId"].(string)
		if !ok {
			continue
		}
		parts := strings.Split(ruleID, "-")
		if len(parts) == 3 {
			seen[strings.ToLower(parts[0])] = true
		}
	}
	for _, fixture := range multilangSources {
		if !seen[fixture.tag] {
			t.Fatalf("no finding attributed to language %q", fixture.tag)
		}
	}
}

// TestCheckMultiLanguageDeterminism proves a mixed-language run is
// byte-for-byte reproducible across repeated invocations.
func TestCheckMultiLanguageDeterminism(t *testing.T) {
	t.Parallel()
	dir := writeMultilangDir(t)
	first := ""
	for i := 0; i < 5; i++ {
		_, out, errs := invoke("check", dir, "--format", "sarif")
		if errs != "" {
			t.Fatalf("run %d stderr=%s", i, errs)
		}
		if i == 0 {
			first = out
			continue
		}
		if out != first {
			t.Fatalf("run %d output differs from run 0", i)
		}
	}
}

// TestCheckLanguageFilter proves --languages restricts dispatch to the named
// subset and leaves other language files unchecked.
func TestCheckLanguageFilter(t *testing.T) {
	t.Parallel()
	dir := writeMultilangDir(t)
	_, out, errs := invoke("check", dir, "--languages", "py,js", "--format", "json")
	if errs != "" {
		t.Fatalf("filtered check stderr=%s", errs)
	}
	document := parseCheckJSON(t, out)
	summary, ok := document["summary"].(map[string]any)
	if !ok {
		t.Fatalf("summary is not an object: %v", document["summary"])
	}
	if summary["files"] != float64(2) {
		t.Fatalf("files = %v, want 2 (py+js only)", summary["files"])
	}
	findings, ok := document["findings"].([]any)
	if !ok {
		t.Fatalf("findings is not an array")
	}
	for _, raw := range findings {
		m, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		ruleID, ok := m["ruleId"].(string)
		if !ok {
			continue
		}
		prefix := strings.ToLower(strings.Split(ruleID, "-")[0])
		if prefix != "py" && prefix != "js" {
			t.Fatalf("language filter leaked finding %s", ruleID)
		}
	}
}

// TestCheckExtensionRouting proves alternate extensions route to the right
// language (e.g. .jsx to js, .kts to kt, .h to c).
func TestCheckExtensionRouting(t *testing.T) {
	t.Parallel()
	cases := []struct {
		ext  string
		lang string
	}{
		{".js", "js"}, {".jsx", "js"}, {".mjs", "js"},
		{".c", "c"}, {".h", "c"},
		{".kt", "kt"}, {".kts", "kt"},
	}
	for _, c := range cases {
		t.Run(c.ext, func(t *testing.T) {
			t.Parallel()
			path := filepath.Join(t.TempDir(), "sample"+c.ext)
			if err := os.WriteFile(path, []byte(multilangSourceFor(c.lang)), 0o600); err != nil {
				t.Fatalf("write fixture: %v", err)
			}
			_, out, errs := invoke("check", path, "--languages", c.lang, "--format", "json")
			if errs != "" {
				t.Fatalf("stderr=%s", errs)
			}
			summary, ok := parseCheckJSON(t, out)["summary"].(map[string]any)
			if !ok {
				t.Fatalf("ext %s: summary is not an object", c.ext)
			}
			if summary["files"] != float64(1) {
				t.Fatalf("ext %s routed to %s: files=%v, want 1", c.ext, c.lang, summary["files"])
			}
		})
	}
}

// multilangSourceFor returns a source snippet for a language tag.
func multilangSourceFor(tag string) string {
	for _, fixture := range multilangSources {
		if fixture.tag == tag {
			return fixture.source
		}
	}
	return ""
}

// TestCheckHelpDocumentsLanguages proves the check help lists all seven
// languages and the new --languages flag.
