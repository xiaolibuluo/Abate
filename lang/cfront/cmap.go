/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// cMapping 是一个 C/C++ 构造归约后的 AIR 种类与语义标注。
type cMapping struct {
	kind     air.NodeKind
	semKey   air.SemKey
	semValue string
	note     string
	name     string
}

// cNesting 区分顶层构造与嵌套语句：用具名布尔类型承载语义，保留
// true/false 字面量的可读性，同时避免把裸 bool 作为控制耦合参数暴露。
type cNesting bool

const cTopLevel cNesting = true

// cMapper 把顶层/语句构造映射到冻结的 air 枚举，只用真实存在的 Kind/SemKey。
type cMapper struct{}

func newCMapper() *cMapper { return &cMapper{} }

// cLeadWord 返回一个区间内第一个以字母或下划线开头的 token 文本。
//
// 属性前缀会被跳过，因此 [[nodiscard]] int f(void) 的引导词是 int 而不是
// nodiscard。
func cLeadWord(tokens []frontkit.Token, lo, hi int) string {
	for j := cAttributeExtent(tokens, lo, hi); j < hi && j < len(tokens); j++ {
		text := tokens[j].Text
		if text == "" {
			continue
		}
		ch := text[0]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
			return text
		}
	}
	return ""
}

// cHasText 报告区间内是否存在指定拼写的 token。
func cHasText(tokens []frontkit.Token, lo, hi int, text string) bool {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text {
			return true
		}
	}
	return false
}

// cNameOf 挑出构造名：跳过声明说明符后的第一个标识符。
//
// 属性前缀会被跳过，因此 [[nodiscard]] int f(void) 的构造名是 f。
func cNameOf(tokens []frontkit.Token, lo, hi int) string {
	fallback := ""
	for j := cAttributeExtent(tokens, lo, hi); j < hi && j < len(tokens); j++ {
		token := tokens[j]
		if token.Text == "" ||
			(token.Kind != frontkit.TokenIdent && token.Kind != frontkit.TokenKeyword) {
			continue
		}
		if fallback == "" {
			fallback = token.Text
		}
		if !cIsSpecifier(token.Text) && token.Text[0] != '_' {
			return token.Text
		}
	}
	return fallback
}

// match 依据首关键字与区间特征，确定性地选择一个构造映射。
//
// 优先级从专有构造到通用声明：预处理/模板/标签 → 声明关键字 → 语句关键字 →
// 函数/变量声明 → 说明符与兜底。全部使用 kinds.go 真实常量。
func (*cMapper) match(tokens []frontkit.Token, lo, hi int, top cNesting) cMapping {
	result := cMapping{name: cNameOf(tokens, lo, hi)}
	if matchCPreproc(tokens, lo, hi, &result) {
		return result
	}
	if matchCDecls(tokens, lo, hi, &result) {
		return result
	}
	if matchCStmts(tokens, lo, hi, &result) {
		return result
	}
	if matchCCallOrDecl(tokens, lo, hi, top, &result) {
		return result
	}
	matchCSpecifiers(tokens, lo, hi, top, &result)
	return result
}

// matchCPreproc 处理预处理行、模板与标签单元，返回是否已命中专属构造。
func matchCPreproc(tokens []frontkit.Token, lo, hi int, result *cMapping) bool {
	if cHasText(tokens, lo, hi, "#") {
		// 预处理行由 cpreproc 归约，映射表在此给出权威种类与语义标注。
		result.kind = air.KindDirective
		result.semKey, result.semValue = air.SemDirective, cLeadWord(tokens, lo+1, hi)
		return true
	}
	lead := cLeadWord(tokens, lo, hi)
	if lead == "template" {
		result.kind, result.semKey, result.semValue = air.KindGenericInst, air.SemTypeParam, "template"
		return true
	}
	if cIsLabelUnit(tokens, lo, hi) {
		result.kind = air.KindLabel
		return true
	}
	return false
}

// matchCDecls 处理命名空间/类/结构体/联合/枚举/using/typedef 声明。
func matchCDecls(tokens []frontkit.Token, lo, hi int, result *cMapping) bool {
	switch cLeadWord(tokens, lo, hi) {
	case "namespace":
		result.kind = air.KindModuleDecl
		return true
	case "class":
		result.kind = air.KindClassDecl
		return true
	case "struct":
		result.kind = air.KindStructDecl
		return true
	case "union":
		result.kind, result.note = air.KindStructDecl, "union"
		return true
	case "enum":
		result.kind = air.KindEnumDecl
		if cHasText(tokens, lo, hi, "class") {
			result.note = "enum-class"
		}
		return true
	case "using":
		result.kind = air.KindUsingDecl
		return true
	case "typedef":
		result.kind = air.KindTypeDecl
		return true
	}
	return false
}

// matchCStmts 处理语句与控制流关键字。
func matchCStmts(tokens []frontkit.Token, lo, hi int, result *cMapping) bool {
	lead := cLeadWord(tokens, lo, hi)
	switch lead {
	case "return":
		result.kind = air.KindReturn
		return true
	case "if", "else":
		result.kind = air.KindIf
		return true
	case "for":
		result.kind = air.KindFor
		return true
	case "while":
		result.kind = air.KindWhile
		return true
	case "do":
		result.kind = air.KindDoWhile
		return true
	case "switch":
		result.kind = air.KindSwitch
		return true
	case "case", "default":
		result.kind = air.KindCase
		return true
	case "break", "continue", "goto":
		result.kind, result.note = air.KindBranch, lead
		return true
	case "throw":
		result.kind = air.KindThrow
		return true
	case "try":
		result.kind = air.KindTry
		return true
	case "catch":
		result.kind = air.KindCatch
		return true
	case "sizeof":
		result.kind, result.note = air.KindCast, "sizeof"
		return true
	case "new":
		result.kind, result.note = air.KindUnary, "new"
		return true
	case "delete":
		result.kind, result.note = air.KindUnary, "delete"
		return true
	}
	return false
}

// matchCCallOrDecl 处理函数定义/调用与类型引导的变量声明。
func matchCCallOrDecl(tokens []frontkit.Token, lo, hi int, top cNesting, result *cMapping) bool {
	lead := cLeadWord(tokens, lo, hi)
	hasParen := cHasText(tokens, lo, hi, "(")
	hasBody := cHasText(tokens, lo, hi, "{")
	if hasParen && (hasBody || top == cTopLevel) {
		result.kind = air.KindFuncDecl
		return true
	}
	if cIsTypeWord(lead) || lead == "auto" {
		result.kind = air.KindVarDecl
		return true
	}
	if hasParen {
		result.kind = air.KindCall
		return true
	}
	return false
}

// matchCSpecifiers 处理说明符与兜底表达式语句。
func matchCSpecifiers(tokens []frontkit.Token, lo, hi int, top cNesting, result *cMapping) {
	lead := cLeadWord(tokens, lo, hi)
	switch lead {
	case "const", "constexpr", "consteval":
		result.kind = air.KindConstDecl
	case "static", "extern":
		result.kind, result.note = air.KindVarDecl, lead
	default:
		if top == cTopLevel {
			result.kind = air.KindVarDecl
		} else {
			result.kind = air.KindExprStmt
		}
	}
}

// cBoundOperator routes one C/C++ operator token through frontkit.BindOperator.
func cBoundOperator(kind air.NodeKind, start, end air.Position,
	tok frontkit.Token, children ...air.Node) air.Node {
	return frontkit.BindOperator(kind, start, end, tok, children...)
}

// annotate 把映射的语义标注写到一个节点上，使用真实 SemKey。
func (*cMapper) annotate(node air.Node, mapping cMapping) {
	if node == nil {
		return
	}
	if mapping.semKey.Valid() && mapping.semValue != "" {
		_ = air.Annotate(node, mapping.semKey, mapping.semValue)
	}
	if mapping.note != "" {
		_ = air.Annotate(node, air.SemNote, mapping.note)
	}
}
