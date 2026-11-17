/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"strings"
	"testing"

	"abate/lang/air"
)

// jsDialectExtensions 是方言矩阵的六种扩展名。
var jsDialectExtensions = []string{".js", ".jsx", ".mjs", ".cjs", ".ts", ".tsx"}

// jsDialectRows 是每种扩展名期望的协商结果。
var jsDialectRows = []struct {
	ext    string
	flavor string
	jsx    bool
	ts     bool
	module string
}{
	{".js", "js", false, false, "auto"},
	{".jsx", "jsx", true, false, "auto"},
	{".mjs", "mjs", false, false, "esm"},
	{".cjs", "cjs", false, false, "commonjs"},
	{".ts", "ts", false, true, "esm"},
	{".tsx", "tsx", true, true, "esm"},
}

// TestResolveDialectSixExtensions 验证六种扩展名的 JSX/TS/模块制式协商与大小写归一。
func TestResolveDialectSixExtensions(t *testing.T) {
	for _, row := range jsDialectRows {
		d := resolveDialect("module" + row.ext)
		if d.flavor != row.flavor || d.jsx != row.jsx || d.ts != row.ts ||
			d.module != row.module {
			t.Errorf("%s: got %+v, want %+v", row.ext, d, row)
		}
		if d.target != "es2022" {
			t.Errorf("%s: target = %q", row.ext, d.target)
		}
	}
	upper := resolveDialect("MODULE.TSX")
	if upper.flavor != "tsx" || !upper.jsx || !upper.ts {
		t.Fatalf("upper-case extension: %+v", upper)
	}
	unknown := resolveDialect("module.unknown")
	if unknown.flavor != "js" || unknown.jsx || unknown.ts {
		t.Fatalf("unknown extension: %+v", unknown)
	}
	seen := make(map[string]bool)
	for _, row := range jsDialectRows {
		seen[resolveDialect("m"+row.ext).canonical()] = true
	}
	if len(seen) != 6 {
		t.Fatalf("expected 6 distinct dialect keys, got %d", len(seen))
	}
}

// TestNormalizeDialectMatrix 遍历六扩展名 × strict/sloppy × target 共 36 种组合。
func TestNormalizeDialectMatrix(t *testing.T) {
	spec := &jsSpec{}
	fe := jsFrontend(t)
	stricts := []string{"strict", "sloppy"}
	targets := []string{"es2020", "es2022", "esnext"}
	src := `const x = 1; export default x;`
	seen := make(map[string]bool)
	for _, ext := range jsDialectExtensions {
		for _, strict := range stricts {
			for _, target := range targets {
				dialect := ext + " " + strict + " target=" + target
				canonical, err := spec.NormalizeDialect(dialect)
				if err != nil {
					t.Fatalf("NormalizeDialect(%q): %v", dialect, err)
				}
				again, err := spec.NormalizeDialect(dialect)
				if err != nil || again != canonical {
					t.Fatalf("non-deterministic dialect: %q vs %q", canonical, again)
				}
				wantStrict := "strict=false"
				if strict == "strict" {
					wantStrict = "strict=true"
				}
				if !strings.Contains(canonical, "js/"+strings.TrimPrefix(ext, ".")+"/") ||
					!strings.Contains(canonical, wantStrict) ||
					!strings.Contains(canonical, "target="+target) {
					t.Errorf("dialect %q -> %q", dialect, canonical)
				}
				seen[canonical] = true
				out, _ := jsParse(t, fe, "matrix"+ext, dialect, src)
				if out.Annotations["lang.js.dialect"] != canonical {
					t.Errorf("unit dialect = %q, want %q",
						out.Annotations["lang.js.dialect"], canonical)
				}
				if out.NodeCount() == 0 {
					t.Errorf("empty tree for %q", dialect)
				}
			}
		}
	}
	if len(seen) != 36 {
		t.Fatalf("expected 36 distinct canonical dialects, got %d", len(seen))
	}
}

// TestNormalizeDialectAliases 验证长别名与 module 显式覆盖的确定性归一。
func TestNormalizeDialectAliases(t *testing.T) {
	spec := &jsSpec{}
	cases := []struct {
		raw    string
		want   string
		suffix string
	}{
		{"", "js/js/auto", "strict=false/target=es2022"},
		{"javascript", "js/js/auto", "strict=false/target=es2022"},
		{"typescript", "js/ts/esm", "strict=true"},
		{"commonjs", "js/cjs/commonjs", "strict=false"},
		{"mjs", "js/mjs/esm", "strict=true"},
		{"jsx strict es6", "js/jsx/auto", "strict=true/target=es6"},
		{"tsx|sloppy|target=es2015", "js/tsx/esm", "strict=false"},
		{"module=cjs js", "js/js/commonjs", "strict=false"},
		{"use strict", "js/js/auto", "strict=true"},
	}
	for _, tc := range cases {
		got, err := spec.NormalizeDialect(tc.raw)
		if err != nil {
			t.Fatalf("NormalizeDialect(%q): %v", tc.raw, err)
		}
		if !strings.Contains(got, tc.want) || !strings.Contains(got, tc.suffix) {
			t.Errorf("NormalizeDialect(%q) = %q, want %q + %q",
				tc.raw, got, tc.want, tc.suffix)
		}
	}
}

// TestDialectAffectsJSXAndTS 证明扩展名协商真正改变归约树与指纹。
func TestDialectAffectsJSXAndTS(t *testing.T) {
	fe := jsFrontend(t)
	src := `const x: number = 1; const el = <div a="1">t</div>; const y = x ?? 0;`
	fingerprints := make(map[[32]byte]bool)
	for _, row := range jsDialectRows {
		out, _ := jsParse(t, fe, "feature"+row.ext, "", src)
		facts := factsOf(out)
		hasJSX := facts["note:jsx"]
		hasType := facts["note:type-annotation"]
		if hasJSX != row.jsx {
			t.Errorf("%s: jsx=%v want %v", row.ext, hasJSX, row.jsx)
		}
		if hasType != row.ts {
			t.Errorf("%s: ts-type=%v want %v", row.ext, hasType, row.ts)
		}
		if out.Roots[0].Sem(air.SemTag) != resolveDialect("f"+row.ext).canonical() {
			t.Errorf("%s: root SemTag mismatch", row.ext)
		}
		fingerprints[out.Fingerprint()] = true
	}
	if len(fingerprints) != 6 {
		t.Fatalf("expected 6 distinct fingerprints, got %d", len(fingerprints))
	}
}

// TestNormalizeDialectInvalid 验证非法描述串返回确定性错误且不 panic。
func TestNormalizeDialectInvalid(t *testing.T) {
	spec := &jsSpec{}
	bad := []string{
		string([]byte{0x00}),
		string([]byte{0x7f, 0x01}),
		strings.Repeat("a", 257),
		"tsx\x01strict",
	}
	for _, d := range bad {
		if _, err := spec.NormalizeDialect(d); err == nil {
			t.Errorf("expected error for dialect %q", d)
		}
	}
	fe := jsFrontend(t)
	unit := air.SourceUnit{
		Language: "js", Dialect: "\x00",
		Files: []air.SourceFile{air.NewSourceFile("bad.js", []byte("let x=1;"))},
	}
	if _, _, err := fe.Parse(unit); err == nil {
		t.Fatal("Parse should reject an invalid dialect")
	}
}

// TestDialectKeywordMatrix 验证 strict/target 组合对单元注解的确定性影响。
func TestDialectKeywordMatrix(t *testing.T) {
	fe := jsFrontend(t)
	src := `async function f() { return await g(); }`
	left, _ := jsParse(t, fe, "m.ts", "ts strict target=es2020", src)
	right, _ := jsParse(t, fe, "m.ts", "ts sloppy target=es2022", src)
	if left.Annotations["lang.js.dialect"] ==
		right.Annotations["lang.js.dialect"] {
		t.Fatal("strict/target did not change the dialect annotation")
	}
	if left.Fingerprint() == right.Fingerprint() {
		t.Fatal("dialect change did not change the fingerprint")
	}
}
