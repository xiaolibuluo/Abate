/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"testing"

	"abate/lang/air"
	fkcov "abate/lang/frontkit/coverage"
)

// TestDedicatedKindsReallyEmitted proves that each dedicated kind is produced
// by a real parse of a construct that actually appears in the source. This is
// the executable anti-fraud guard: a mapping may only be marked exact when the
// frontend truly lowers the construct onto its dedicated kind.
func TestDedicatedKindsReallyEmitted(t *testing.T) {
	cases := []struct {
		name string
		src  string
		kind air.NodeKind
	}{
		{"export", `export const x = 1;`, air.KindJsExport},
		{"lexical-decl", `const x = 1;`, air.KindJsLexicalDecl},
		{"declarator", `const x = 1;`, air.KindJsDeclarator},
		{"this-super", `return this;`, air.KindJsThisSuper},
		{"generator-decl", `function* g(){}`, air.KindJsGeneratorFnDecl},
		{"generator-fn", `function* g(){}`, air.KindJsGeneratorFn},
		{"new", `const p = new Box(1);`, air.KindJsNew},
		{"paren", `const g = (1 + 2);`, air.KindJsParenExpr},
		{"sequence", `const g = (1, 2);`, air.KindJsSequence},
		{"incdec", `x++;`, air.KindJsIncDec},
		{"class-body", `class C {}`, air.KindJsClassBody},
		{"class-heritage", `class C extends B {}`, air.KindJsClassHeritage},
		{"switch-body", `switch (x) { case 1: break; }`, air.KindJsSwitchBody},
		{"switch-default", `switch (x) { default: break; }`, air.KindJsSwitchDefault},
		{"yield", `function* g(){ yield 1; }`, air.KindJsYield},
		{"template-sub", "const t = `a${x}b`;", air.KindJsTemplateSub},
		{"string-part", `const s = "abc";`, air.KindJsStringPart},
		{"property-ident", `const v = o.foo;`, air.KindJsPropertyIdent},
		{"arguments", `f(1, 2);`, air.KindJsArguments},
		{"formal-params", `function f(a, b) {}`, air.KindJsFormalParams},
	}
	fe := jsFrontend(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			unit, _ := jsParse(t, fe, "f.js", "", tc.src)
			observed := fkcov.Observe(unit)
			if observed[tc.kind] == 0 {
				t.Fatalf("kind %s (%d) not emitted for source %q; observed kinds=%v",
					tc.kind, tc.kind, tc.src, observed)
			}
		})
	}
}

// TestDedicatedMarkersZeroWidth proves that dedicated markers never widen a
// node: every emitted dedicated-kind node has an empty byte span, so lexical
// fidelity and existing AIR-IDs stay byte-exact.
func TestDedicatedMarkersZeroWidth(t *testing.T) {
	fe := jsFrontend(t)
	unit, _ := jsParse(t, fe, "f.js", "",
		`export function* g(a){ yield a.b; const p = new Box(a); }`)
	for _, node := range air.Collect(unit) {
		name := node.Kind().String()
		if len(name) > 2 && name[:2] == "Js" {
			pos, end := node.Pos(), node.End()
			if pos.Offset != end.Offset {
				t.Fatalf("dedicated kind %s is not zero-width: [%d,%d)",
					name, pos.Offset, end.Offset)
			}
		}
	}
}
