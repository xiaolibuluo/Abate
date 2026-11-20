/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckHelpDocumentsLanguages(t *testing.T) {
	t.Parallel()
	_, out, errs := invoke("check", "--help")
	if errs != "" {
		t.Fatalf("stderr=%s", errs)
	}
	for _, want := range []string{
		"--languages", ".go", ".py", ".js", ".c", ".cs", ".rs", ".kt",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("check help missing %q", want)
		}
	}
}

// TestCheckToleratesBrokenSyntax is the executable counterpart of the
// "error-tolerant, not a syntax gate" capability boundary. A syntactically
// broken Python source must still produce a report (partial AIR + findings)
// rather than being rejected as a usage error. This locks in the documented
// behavior so a future "strict parse" change cannot silently invalidate it.
func TestCheckToleratesBrokenSyntax(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "broken.py")
	if err := os.WriteFile(path, []byte("def f(:\n  x =\n"), 0o600); err != nil {
		t.Fatalf("write broken fixture: %v", err)
	}
	code, out, errs := invoke("check", path, "--format", "json")
	if code == exitUsageError {
		t.Fatalf("broken syntax was rejected as usage error (rc=%d, stderr=%s)", code, errs)
	}
	document := parseCheckJSON(t, out)
	if _, ok := document["summary"].(map[string]any); !ok {
		t.Fatalf("broken syntax produced no report summary (out=%s)", out)
	}
}

// TestCheckNestedStructureMatchesAST is the executable evidence that rules
// match on the parsed AST, not on flat text: a four-level nested Python file
// must trigger several distinct structural PY-AIR-* rule classes. A pure
// line/regex scanner would not attach multiple structural rules to nested nodes.
func TestCheckNestedStructureMatchesAST(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "nested.py")
	src := "def outer(x):\n    if x:\n        for i in range(3):\n" +
		"            y = i\n            return y\n    return x\n"
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatalf("write nested fixture: %v", err)
	}
	code, out, errs := invoke("check", path, "--format", "json")
	if code != exitBusinessFailure && code != exitOK {
		t.Fatalf("nested check exit = %d (stderr=%s)", code, errs)
	}
	document := parseCheckJSON(t, out)
	findings, ok := document["findings"].([]any)
	if !ok || len(findings) == 0 {
		t.Fatal("nested source produced no findings")
	}
	classes := map[string]bool{}
	for _, raw := range findings {
		f, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		id, ok := f["ruleId"].(string)
		if !ok {
			continue
		}
		if strings.HasPrefix(id, "PY-AIR-") {
			classes[id] = true
		}
	}
	if len(classes) < 3 {
		t.Fatalf("nested source hit only %d PY-AIR-* structural classes: %v", len(classes), classes)
	}
}
