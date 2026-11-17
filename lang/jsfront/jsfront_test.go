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

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsComprehensiveLines 覆盖 JS/TS 的声明、语句、表达式与专有构造。
var jsComprehensiveLines = []string{
	`"use strict";`,
	`import { readFile } from "fs";`,
	`export const VERSION: number = 2;`,
	`let count = 0;`,
	`const ratio = 3.5;`,
	`function add(a: number, b: number): number { return a + b; }`,
	`async function load(p: string): Promise<string> { return await fetch(p); }`,
	`class Box { private value: number; constructor(v: number) { this.value = v; } }`,
	`interface Shape { readonly id?: string; area(): number; }`,
	`enum Color { Red = 1, Green }`,
	`type Maybe = string | null;`,
	`type Pair<A, B> = [A, B];`,
	`namespace Util { export const ok = true; }`,
	`const fn = (x: number) => x * 2;`,
	`const obj = { a: 1, b: "two" };`,
	`const arr = [1, 2, 3];`,
	"const tpl = `hi ${name}!`;",
	`const re = /ab+c/gi;`,
	`const el = <div className="box">text</div>;`,
	`const opt = obj?.a?.b ?? 0;`,
	`const spread = [...arr];`,
	`const ternary = count ? 1 : 2;`,
	`const cast = value as string;`,
	`const sat = value satisfies number;`,
	`@decorator`,
	`class Decorated {}`,
	`if (count > 0) { count--; } else { count++; }`,
	`for (const item of arr) { use(item); }`,
	`while (count < 10) { count++; }`,
	`do { count--; } while (count > 0);`,
	`switch (count) { case 1: break; default: break; }`,
	`try { risky(); } catch (e) { throw e; } finally { cleanup(); }`,
	`const p = new Box(1);`,
	`const dyn = import("fs");`,
	`const inst = obj instanceof Object;`,
	`const div = 10 / 2;`,
	`const pow = 2 ** 3;`,
	`const assign = count ??= 5;`,
	`const group = (1 + 2) * 3;`,
	`const idx = arr[0];`,
	`label: for (;;) { break label; }`,
}

// jsComprehensiveSource 返回以换行拼接的综合源码。
func jsComprehensiveSource() string { return strings.Join(jsComprehensiveLines, "\n") }

// jsFrontend 构造一个测试用前端，失败即终止。
func jsFrontend(t *testing.T) air.Frontend {
	t.Helper()
	fe, err := NewJsFrontend()
	if err != nil || fe == nil {
		t.Fatalf("NewJsFrontend: fe=%v err=%v", fe, err)
	}
	return fe
}

// jsParse 解析单文件源码并返回单元与诊断。
func jsParse(t *testing.T, fe air.Frontend, path, dialect, src string) (*air.Unit,
	[]air.Diagnostic) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "js", Dialect: dialect,
		Files: []air.SourceFile{air.NewSourceFile(path, []byte(src))},
	}
	out, diags, err := fe.Parse(unit)
	if err != nil {
		t.Fatalf("Parse(%s): %v", path, err)
	}
	if out == nil {
		t.Fatalf("Parse(%s): nil unit", path)
	}
	return out, diags
}

// factsOf 收集单元里全部节点类与语义注记，形成可断言的事实集合。
func factsOf(u *air.Unit) map[string]bool {
	facts := make(map[string]bool)
	for _, n := range air.Collect(u) {
		facts["kind:"+n.Kind().String()] = true
		if v, ok := n.Sem(air.SemNote).(string); ok {
			facts["note:"+v] = true
		}
		if v, ok := n.Sem(air.SemOperator).(string); ok {
			facts["op:"+v] = true
		}
		if v, ok := n.Sem(air.SemTypeRef).(string); ok && v != "" {
			facts["tref:"+v] = true
		}
		if v, ok := n.Sem(air.SemTypeParam).(string); ok && v != "" {
			facts["tparam:"+v] = true
		}
		if v, ok := n.Sem(air.SemAsync).(string); ok {
			facts["async:"+v] = true
		}
		if v, ok := n.Sem(air.SemMutability).(string); ok {
			facts["mut:"+v] = true
		}
		if v, ok := n.Sem(air.SemDirective).(string); ok {
			facts["dir:"+v] = true
		}
		if v, ok := n.Sem(air.SemPartial).(string); ok {
			facts["partial:"+v] = true
		}
	}
	return facts
}

// requireFacts 断言每一个事实都出现在解析结果中。
func requireFacts(t *testing.T, facts map[string]bool, want ...string) {
	t.Helper()
	for _, key := range want {
		if !facts[key] {
			t.Errorf("missing mapping %q", key)
		}
	}
}

// TestNewJsFrontendSatisfiesFrontend 验证返回值满足 air.Frontend 且可真解析。
func TestNewJsFrontendSatisfiesFrontend(t *testing.T) {
	fe := jsFrontend(t)
	var _ = fe
	out, _ := jsParse(t, fe, "main.js", "", `const x = 1;`)
	if out.NodeCount() == 0 {
		t.Fatal("empty AIR tree")
	}
	if len(out.Roots) != 1 || out.Roots[0].Kind() != air.KindPackage {
		t.Fatalf("root kind = %v", out.Roots[0].Kind())
	}
}

// TestTagAndLexSpec 验证短标签、词法表与扩展名接线。
func TestTagAndLexSpec(t *testing.T) {
	spec := &jsSpec{}
	if spec.Tag() != "js" {
		t.Fatalf("Tag = %q", spec.Tag())
	}
	tag, ok := frontkit.CanonicalTag(spec.Tag())
	if !ok || tag != "js" {
		t.Fatalf("CanonicalTag = %q,%v", tag, ok)
	}
	if got := frontkit.ExtensionsFor("js"); len(got) != 6 {
		t.Fatalf("js extensions = %v", got)
	}
	ls := spec.LexSpec()
	if ls == nil || ls.LineComment() != "//" || ls.IndentSensitive() {
		t.Fatal("lex spec mismatch")
	}
	opening, closing := ls.BlockComment()
	if opening != "/*" || closing != "*/" {
		t.Fatalf("block comment = %q,%q", opening, closing)
	}
	if len(ls.Keywords()) < 40 || len(ls.Operators()) < 40 {
		t.Fatalf("tables too small: %d/%d", len(ls.Keywords()), len(ls.Operators()))
	}
	if len(ls.StringQuotes()) != 3 || len(ls.NumberSeparators()) != 1 {
		t.Fatal("string/separator table mismatch")
	}
	if len(ls.RawStringPrefixes()) != 0 {
		t.Fatal("JS has no raw string prefixes")
	}
}

// TestAirEnumsAreReal 证明降级只使用 lang/air 真实存在的枚举，AIR 零改动。
func TestAirEnumsAreReal(t *testing.T) {
	kinds := []air.NodeKind{
		air.KindPackage, air.KindImport, air.KindModuleDecl, air.KindConstDecl,
		air.KindVarDecl, air.KindFuncDecl, air.KindClassDecl, air.KindInterfaceDecl,
		air.KindEnumDecl, air.KindTypeDecl, air.KindArrow, air.KindAwait,
		air.KindBinary, air.KindAssign, air.KindCall, air.KindField,
		air.KindSelector, air.KindIndex, air.KindTernary, air.KindSpread,
		air.KindCompositeLit, air.KindGenericInst, air.KindGenericParam,
		air.KindUnion, air.KindOptional, air.KindTuple, air.KindNamed,
		air.KindTypeAssert, air.KindDecorator, air.KindDirective, air.KindRegion,
	}
	for _, k := range kinds {
		if !k.Valid() {
			t.Fatalf("invalid AIR kind %d", k)
		}
	}
	keys := []air.SemKey{
		air.SemTypeRef, air.SemScopeID, air.SemMutability, air.SemAsync,
		air.SemDirective, air.SemTag, air.SemNote, air.SemPartial,
		air.SemOperator, air.SemTypeParam,
	}
	for _, k := range keys {
		if !k.Valid() {
			t.Fatalf("invalid AIR sem key %d", k)
		}
	}
}

// TestParseProducesRealTree 断言综合源码归约出真实层级节点树与 ≥20 个构造映射。
func TestParseProducesRealTree(t *testing.T) {
	fe := jsFrontend(t)
	out, _ := jsParse(t, fe, "all.tsx", "tsx/strict/es2022",
		jsComprehensiveSource())
	if out.NodeCount() < 60 {
		t.Fatalf("node count too small: %d", out.NodeCount())
	}
	facts := factsOf(out)
	t.Logf("FACTS(%d)=%v", len(facts), sortedFactKeys(facts))
	requireFacts(t, facts,
		"kind:Package", "kind:Import", "kind:ModuleDecl", "kind:ConstDecl",
		"kind:VarDecl", "kind:FuncDecl", "kind:ClassDecl", "kind:InterfaceDecl",
		"kind:EnumDecl", "kind:TypeDecl", "kind:Arrow", "kind:Await",
		"kind:Binary", "kind:Assign", "kind:Call", "kind:Field",
		"kind:Selector", "kind:Index", "kind:Ternary", "kind:Spread",
		"kind:CompositeLit", "kind:GenericInst", "kind:GenericParam",
		"kind:Union", "kind:Optional", "kind:Tuple", "kind:Named",
		"kind:TypeAssert", "kind:Decorator", "kind:Directive",
		"kind:Return", "kind:If", "kind:RangeFor", "kind:While",
		"kind:DoWhile", "kind:Switch", "kind:Case", "kind:Try",
		"kind:Catch", "kind:Finally", "kind:Throw", "kind:Branch",
		"note:jsx", "note:optional-chain", "note:nullish",
		"note:template-string", "note:template-interpolation",
		"note:regex-literal", "note:arrow-function", "note:decorator",
		"note:use-strict", "note:type-annotation", "note:union-type",
		"note:new-expression", "note:dynamic-import", "note:jsx-attribute",
		"op:optional-chain", "op:nullish", "op:division", "op:exponentiation",
		"op:conditional", "op:regex", "mut:immutable", "mut:mutable",
		"async:true", "dir:import", "dir:export", "dir:use strict",
	)
}

// sortedFactKeys 返回排序后的事实键，仅用于测试日志。
func sortedFactKeys(facts map[string]bool) []string {
	keys := make([]string, 0, len(facts))
	for k := range facts {
		keys = append(keys, k)
	}
	for i := 1; i < len(keys); i++ {
		for j := i; j > 0 && keys[j] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
		}
	}
	return keys
}

// TestDeterminism100 验证同一输入 100 次指纹逐字节一致。
func TestDeterminism100(t *testing.T) {
	fe := jsFrontend(t)
	unit := air.SourceUnit{
		Language: "js", Dialect: "tsx/strict/es2022",
		Files: []air.SourceFile{
			air.NewSourceFile("all.tsx", []byte(jsComprehensiveSource())),
		},
	}
	determinism.MustMatch100(t, 20260521, func() []byte {
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestParseTokensDirect 直接调用 LanguageSpec.ParseTokens 验证真归约。
func TestParseTokensDirect(t *testing.T) {
	spec := &jsSpec{}
	src := []byte("const x = obj?.a ?? 1;\n")
	lexer, err := frontkit.NewLexer(spec.LexSpec(), src)
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize: %v", err)
	}
	root, diags, err := spec.ParseTokens(tokens, "direct.ts")
	if err != nil || root == nil {
		t.Fatalf("ParseTokens: root=%v err=%v", root, err)
	}
	if root.Kind() != air.KindPackage || len(root.Children()) == 0 {
		t.Fatalf("root = %s with %d children", root.Kind(), len(root.Children()))
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
}
