/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// ktFuzzSeeds 是原生 Go fuzz 的确定性种子语料。
func ktFuzzSeeds() []string {
	return []string{
		"",
		"package demo\n",
		"data class User(val name: String)\n",
		"suspend fun fetch(url: String): Int { return url.length }\n",
		"fun String.shout() = uppercase()\n",
		"val s = \"hi $name and ${name.len}\"\n",
		"val r = when (x) { 1 -> a() else -> b() }\n",
		"val n = a?.len ?: 0\n",
		"val raw = \"\"\"raw text\"\"\"\n",
		"val `weird name` = 42\n",
		"/** KDoc */\n",
		"class C { fun f( {\n",
		"\xff\xfe\n",
		"val z = a!!\n",
		"sealed class R { companion object { } }\n",
	}
}

// FuzzKtFrontend 用 Go 原生 fuzz 归约任意字节缓冲。适配器必须不 panic，
// 未报告分类错误时必须返回合法 unit 与确定的指纹。
func FuzzKtFrontend(f *testing.F) {
	for _, seed := range ktFuzzSeeds() {
		f.Add([]byte(seed))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		fe, err := NewKtFrontend()
		if err != nil {
			t.Fatalf("NewKtFrontend: %v", err)
		}
		unit := air.SourceUnit{
			Language: "kt", Dialect: "kt2.1",
			Files: []air.SourceFile{air.NewSourceFile("fuzz.kt", data)},
		}
		out, diags, parseErr := fe.Parse(unit)
		if parseErr != nil {
			return
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("nil or malformed unit for %q", data[:ktMinInt(len(data), 32)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("invalid unit for %q: %v", data[:ktMinInt(len(data), 32)], err)
		}
		if out.NodeCount() < 1 {
			t.Fatal("unit has no nodes")
		}
		if out.Roots[0].Kind() != air.KindPackage {
			t.Fatal("root is not KindPackage")
		}
		_ = out.Fingerprint()
		_ = diags
	})
}

// TestKtFuzzSeedsReproducible 把种子语料各解析两次，断言指纹逐字节一致。
func TestKtFuzzSeedsReproducible(t *testing.T) {
	fe := ktFrontend(t)
	for _, src := range ktFuzzSeeds() {
		request := air.SourceUnit{
			Language: "kt", Dialect: "kt2.1",
			Files: []air.SourceFile{air.NewSourceFile("seed.kt", []byte(src))},
		}
		a, _, errA := fe.Parse(request)
		b, _, errB := fe.Parse(request)
		if (errA == nil) != (errB == nil) || a == nil || b == nil {
			t.Fatalf("inconsistent seed parse for %q", src)
		}
		if a.Fingerprint() != b.Fingerprint() {
			t.Fatalf("seed %q is not deterministic", src)
		}
		if a.NodeCount() != b.NodeCount() {
			t.Fatalf("seed %q has non-deterministic node count", src)
		}
	}
}

// TestKtFuzzSeedsIntervals 断言每个种子都能通过位置区间校验。
func TestKtFuzzSeedsIntervals(t *testing.T) {
	for _, src := range ktFuzzSeeds() {
		out, _, err := ktFrontend(t).Parse(air.SourceUnit{
			Language: "kt", Dialect: "kt2.1",
			Files: []air.SourceFile{air.NewSourceFile("seed.kt", []byte(src))},
		})
		if err != nil {
			t.Fatalf("Parse(%q): %v", src, err)
		}
		if err := air.CheckIntervals(out.Roots[0], len(src)); err != nil {
			t.Fatalf("CheckIntervals for %q: %v", src, err)
		}
	}
}

// TestKtTableLookups 覆盖所有有序表查找（关键字/标注/中缀/声明/方言/空安全）的命中与未命中分支。
func TestKtTableLookups(t *testing.T) {
	if !ktContains(ktKeywords, "class") || ktContains(ktKeywords, "nope") {
		t.Fatal("ktContains broken")
	}
	if ktNoteValue(ktNotes, "null-safe") != "null-safe" ||
		ktNoteValue(ktNotes, "unknown-note") != "unknown-note" {
		t.Fatal("ktNoteValue broken")
	}
	if ktInfixPower(ktInfixBindings, "?:") != 15 ||
		ktInfixPower(ktInfixBindings, ",") != 3 ||
		ktInfixPower(ktInfixBindings, "nope") != 0 {
		t.Fatal("ktInfixPower broken")
	}
	if !ktInfixRight(ktInfixBindings, "=") ||
		ktInfixRight(ktInfixBindings, "+") || ktInfixRight(ktInfixBindings, "nope") {
		t.Fatal("ktInfixRight broken")
	}
	if rule, ok := ktDeclRuleFor("package"); !ok || rule.kind != air.KindModuleDecl {
		t.Fatal("ktDeclRuleFor broken")
	}
	if _, ok := ktDeclRuleFor("nope"); ok {
		t.Fatal("ktDeclRuleFor accepted unknown keyword")
	}
	if got, ok := ktDialectFor("2.1"); !ok || got != "kt2.1" {
		t.Fatal("ktDialectFor broken")
	}
	if _, ok := ktDialectFor("nope"); ok {
		t.Fatal("ktDialectFor accepted unknown dialect")
	}
	if rule, ok := ktNullSafeFor(ktNullSafeRules, "!!"); !ok ||
		rule.kind != air.KindUnary {
		t.Fatal("ktNullSafeFor broken")
	}
	if _, ok := ktNullSafeFor(ktNullSafeRules, "nope"); ok {
		t.Fatal("ktNullSafeFor accepted unknown operator")
	}
}

// TestKtTokenLookups 覆盖 tokenizer 派生辅助函数（文本/名称/扩展点/字符串风味/标识符字节/接收者提示）的命中与未命中分支。
func TestKtTokenLookups(t *testing.T) {
	tokens := ktTokenize(t, "data class C { }\n")
	if !ktHasText(tokens, 0, len(tokens), "class") ||
		ktHasText(tokens, 0, len(tokens), "nope") {
		t.Fatal("ktHasText broken")
	}
	if ktLeadIndex(tokens, 0, len(tokens)) != 1 {
		t.Fatal("ktLeadIndex must skip the data modifier")
	}
	if ktNameOf(tokens, 0, len(tokens)) != "C" {
		t.Fatal("ktNameOf broken")
	}
	if !ktHasExtensionDot(ktTokenize(t, "fun String.shout() { }\n"), 0, 6) ||
		ktHasExtensionDot(tokens, 0, 1) {
		t.Fatal("ktHasExtensionDot broken")
	}
	if start, end := ktSpanTokens(tokens, 5, 2); start.Offset != 0 || end.Offset != 0 {
		t.Fatal("ktSpanTokens empty interval broken")
	}
	for _, tc := range []struct{ text, want string }{
		{`"""a"""`, "raw-string"}, {"`id`", "backtick-ident"},
		{`"hi $x"`, "string-template"}, {`"plain"`, "plain"},
	} {
		if got := ktStringFlavor(tc.text); got != tc.want {
			t.Fatalf("ktStringFlavor(%q)=%q, want %q", tc.text, got, tc.want)
		}
	}
	if !ktIsDocComment("/** doc */") || ktIsDocComment("/* doc */") {
		t.Fatal("ktIsDocComment broken")
	}
	if !ktIdentStartByte('_') || !ktIdentStartByte('a') || ktIdentStartByte('1') {
		t.Fatal("ktIdentStartByte broken")
	}
	if !ktIdentContinueByte('9') || !ktIdentContinueByte(0x80) ||
		ktIdentContinueByte('$') {
		t.Fatal("ktIdentContinueByte broken")
	}
	if !ktHasReceiverHint(ktTokenize(t, "this + 1\n"), 0, 2) {
		t.Fatal("ktHasReceiverHint broken")
	}
}

// TestKtMapperEdgeCases 覆盖节点构造器的空值与表外标注分支。
func TestKtMapperEdgeCases(t *testing.T) {
	mapper := newKtMap(nil, nil)
	if mapper.note("unknown-note") != "unknown-note" {
		t.Fatal("note fallback broken")
	}
	mapper.apply(nil, ktMapping{kind: air.KindVarDecl, note: "lambda"})
	tok := frontkit.Token{
		Kind: frontkit.TokenIdent, Text: "x", Offset: 0, EndOffset: 1,
		Line: 1, Column: 1,
	}
	if node := mapper.infixNode(air.KindBinary, nil, tok,
		mapper.literal(tok)); node.Kind() != air.KindBinary {
		t.Fatal("infixNode nil-left fallback broken")
	}
	if node := mapper.unaryNode(air.KindMatchExpr, tok, nil); node.Kind() != air.KindMatchExpr {
		t.Fatal("unaryNode nil-operand fallback broken")
	}
	for _, kind := range []frontkit.TokenKind{
		frontkit.TokenIdent, frontkit.TokenKeyword, frontkit.TokenNum,
		frontkit.TokenStr, frontkit.TokenChar, frontkit.TokenComment,
		frontkit.TokenOp, frontkit.TokenPunct, frontkit.TokenIllegal,
	} {
		probe := tok
		probe.Kind = kind
		if mapper.literal(probe) == nil {
			t.Fatalf("literal returned nil for %s", kind)
		}
	}
	if mapper.comment(frontkit.Token{
		Kind: frontkit.TokenComment, Text: "/* block */",
		Offset: 0, EndOffset: 11, Line: 1, Column: 1,
	}).Kind() != air.KindComment {
		t.Fatal("block comment lowering broken")
	}
	mapper.apply(mapper.literal(tok), ktMapping{
		kind: air.KindVarDecl, note: "lambda", semKey: air.SemOperator,
		semValue: "?", async: true, partial: true, mutability: "lateinit",
	})
}
