/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsKeywords 是 JS/TS 保留字与上下文关键字；true/false/null/undefined 保留为字面量标识符.
var jsKeywords = []string{
	"abstract", "any", "as", "asserts", "async", "await", "bigint", "boolean",
	"break", "case", "catch", "class", "const", "constructor", "continue",
	"debugger", "declare", "default", "delete", "do", "else", "enum", "export",
	"extends", "finally", "for", "from", "function", "get", "global", "if",
	"implements", "import", "in", "infer", "instanceof", "interface", "is",
	"keyof", "let", "module", "namespace", "never", "new", "number", "object",
	"of", "override", "private", "protected", "public", "readonly", "require",
	"return", "satisfies", "set", "static", "string", "super", "switch",
	"symbol", "this", "throw", "try", "type", "typeof", "unique", "unknown",
	"var", "void", "while", "with", "yield",
}

// jsOperators 是运算符表；frontkit.NormalizeTables 会按字节长度降序归一保证最长匹配.
var jsOperators = []string{
	">>>=", "===", "!==", "**=", "&&=", "||=", "??=",
	"...", "=>", "?.", "??", ">>>", "<<=", ">>=", "==", "!=", "<=", ">=",
	"&&", "||", "++", "--", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=",
	"<<", ">>", "**", "<", ">", "+", "-", "*", "/", "%", "&", "|", "^",
	"!", "=", "~",
}

// jsStringQuotes 覆盖双引号、单引号与反引号模板字符串.
var jsStringQuotes = []string{"\"", "'", "`"}

// jsNumberSeparators 是 ECMAScript 数字分组分隔符.
var jsNumberSeparators = []string{"_"}

// jsRestrictedWords 是 ASI 受限产生式关键字，其后换行必须插入分号.
var jsRestrictedWords = []string{"return", "throw", "break", "continue", "yield"}

// jsContinuationTokens 是禁止在其前插入分号的表达式续接符.
var jsContinuationTokens = []string{
	".", "?.", "(", "[", "+", "-", "*", "/", "%", "**", "=", "==", "===",
	"!=", "!==", "<", ">", "<=", ">=", "&&", "||", "??", "?", ":", ",",
	"=>", "++", "--", "+=", "-=", "*=", "/=", "%=", "&", "|", "^", "&=",
	"|=", "^=", "<<", ">>", ">>>", "<<=", ">>=", ">>>=", "&&=", "||=",
	"??=", "instanceof", "in", "as", "satisfies", "of",
}

// jsRegexOpeners 是允许紧随其后的 "/" 作为正则字面量起始的前置 token 文本.
var jsRegexOpeners = []string{
	"=", "(", "[", ",", ":", ";", "{", "}", "return", "throw", "case",
	"delete", "void", "typeof", "instanceof", "in", "of", "yield", "await",
	"=>", "??", "&&", "||", "!", "~", "+", "-", "*", "%", "&", "|", "^",
	"<", ">", "==", "===", "!=", "!==", "<=", ">=", "?.", "...", "default",
	"else", "do", "?", "[", "]",
}

// jsTypeKeywords 是 TS 类型归约识别的类型关键字与类型运算符.
var jsTypeKeywords = []string{
	"any", "unknown", "never", "void", "number", "string", "boolean", "symbol",
	"bigint", "object", "undefined", "null", "keyof", "infer", "readonly",
	"unique", "asserts", "is", "extends", "as", "satisfies", "abstract",
}

// jsTargets 是方言协商接受的 ECMAScript 目标版本.
var jsTargets = []string{
	"es5", "es6", "es2015", "es2016", "es2017", "es2018", "es2019",
	"es2020", "es2021", "es2022", "es2023", "es2024", "esnext",
}

// jsDeclRule 把一个语句首 token 文本映射到真实 AIR 节点类与降级注记.
type jsDeclRule struct {
	text string
	kind air.NodeKind
	note string
}

// jsDeclRules 是声明/语句关键字到 AIR 的有序映射表，线性查找保证确定性.
var jsDeclRules = []jsDeclRule{
	{"import", air.KindImport, "esm-import"},
	{"export", air.KindModuleDecl, "export"},
	{"const", air.KindConstDecl, "const-declaration"},
	{"let", air.KindVarDecl, "let-declaration"},
	{"var", air.KindVarDecl, "var-declaration"},
	{"function", air.KindFuncDecl, "function-declaration"},
	{"class", air.KindClassDecl, "class-declaration"},
	{"interface", air.KindInterfaceDecl, "ts-interface"},
	{"enum", air.KindEnumDecl, "ts-enum"},
	{"type", air.KindTypeDecl, "ts-type-alias"},
	{"namespace", air.KindModuleDecl, "ts-namespace"},
	{"module", air.KindModuleDecl, "ts-namespace"},
	{"declare", air.KindDirective, "ts-declare"},
	{"abstract", air.KindClassDecl, "ts-abstract"},
	{"return", air.KindReturn, "return"},
	{"throw", air.KindThrow, "throw"},
	{"if", air.KindIf, "if"},
	{"else", air.KindIf, "else-branch"},
	{"for", air.KindFor, "for"},
	{"while", air.KindWhile, "while"},
	{"do", air.KindDoWhile, "do-while"},
	{"switch", air.KindSwitch, "switch"},
	{"case", air.KindCase, "case"},
	{"default", air.KindCase, "default-case"},
	{"try", air.KindTry, "try"},
	{"catch", air.KindCatch, "catch"},
	{"finally", air.KindFinally, "finally"},
	{"break", air.KindBranch, "break"},
	{"continue", air.KindBranch, "continue"},
	{"with", air.KindWith, "with"},
	{"get", air.KindMethodDecl, "accessor-get"},
	{"set", air.KindMethodDecl, "accessor-set"},
	{"constructor", air.KindMethodDecl, "constructor"},
	{"static", air.KindMethodDecl, "static-member"},
	{"yield", air.KindUnary, "yield"},
	{"await", air.KindAwait, "await"},
	{"async", air.KindFuncDecl, "async-function"},
	{"@", air.KindDecorator, "decorator"},
	{"{", air.KindBlock, "block"},
}

// jsDeclMapping 返回语句首 token 对应的 AIR 节点类与注记.
func jsDeclMapping(text string) (air.NodeKind, string, bool) {
	for i := range jsDeclRules {
		if jsDeclRules[i].text == text {
			return jsDeclRules[i].kind, jsDeclRules[i].note, true
		}
	}
	return air.KindInvalid, "", false
}

// jsIsLexicalToken 判定 token 是否属于保真子集，必须投影到 AIR lexical 节点.
func jsIsLexicalToken(kind frontkit.TokenKind) bool {
	switch kind {
	case frontkit.TokenIdent, frontkit.TokenNum, frontkit.TokenStr,
		frontkit.TokenChar, frontkit.TokenComment, frontkit.TokenIllegal:
		return true
	default:
		return false
	}
}

// jsInStringSlice 以线性扫描判定成员，避免 map 迭代序参与任何输出.
func jsInStringSlice(list []string, value string) bool {
	for i := range list {
		if list[i] == value {
			return true
		}
	}
	return false
}

// jsIsRestricted 判定是否为 ASI 受限产生式关键字.
func jsIsRestricted(text string) bool {
	return jsInStringSlice(jsRestrictedWords, text)
}

// jsIsContinuation 判定是否为禁止插入分号的续接符.
func jsIsContinuation(text string) bool {
	return jsInStringSlice(jsContinuationTokens, text)
}

// jsIsRegexOpener 判定前置 token 是否允许 "/" 起始正则字面量.
func jsIsRegexOpener(text string) bool {
	return jsInStringSlice(jsRegexOpeners, text)
}

// jsIsTypeKeyword 判定 token 文本是否为 TS 类型关键字.
func jsIsTypeKeyword(text string) bool {
	return jsInStringSlice(jsTypeKeywords, text)
}

// jsBindRule 是一个中缀运算符的绑定力与 AIR 降级映射.
type jsBindRule struct {
	text     string
	power    int
	right    bool
	kind     air.NodeKind
	operator string
}

// jsInfixRules 是 JS/TS 中缀绑定表，线性查找保证无 map 迭代序参与归约.
var jsInfixRules = []jsBindRule{
	{"??", 3, false, air.KindBinary, "nullish"},
	{"||", 3, false, air.KindBinary, "logical-or"},
	{"&&", 4, false, air.KindBinary, "logical-and"},
	{"|", 5, false, air.KindBinary, "bitwise-or"},
	{"^", 6, false, air.KindBinary, "bitwise-xor"},
	{"&", 7, false, air.KindBinary, "bitwise-and"},
	{"==", 8, false, air.KindBinary, "equality"},
	{"!=", 8, false, air.KindBinary, "inequality"},
	{"===", 8, false, air.KindBinary, "strict-equality"},
	{"!==", 8, false, air.KindBinary, "strict-inequality"},
	{"<", 9, false, air.KindBinary, "less-than"},
	{">", 9, false, air.KindBinary, "greater-than"},
	{"<=", 9, false, air.KindBinary, "less-or-equal"},
	{">=", 9, false, air.KindBinary, "greater-or-equal"},
	{"instanceof", 9, false, air.KindBinary, "instanceof"},
	{"in", 9, false, air.KindBinary, "in"},
	{"<<", 10, false, air.KindBinary, "left-shift"},
	{">>", 10, false, air.KindBinary, "right-shift"},
	{">>>", 10, false, air.KindBinary, "unsigned-right-shift"},
	{"+", 11, false, air.KindBinary, "addition"},
	{"-", 11, false, air.KindBinary, "subtraction"},
	{"*", 12, false, air.KindBinary, "multiplication"},
	{"/", 12, false, air.KindBinary, "division"},
	{"%", 12, false, air.KindBinary, "remainder"},
	{"**", 13, true, air.KindBinary, "exponentiation"},
	{"=", 1, true, air.KindAssign, "assignment"},
	{"+=", 1, true, air.KindAssign, "add-assignment"},
	{"-=", 1, true, air.KindAssign, "subtract-assignment"},
	{"*=", 1, true, air.KindAssign, "multiply-assignment"},
	{"/=", 1, true, air.KindAssign, "divide-assignment"},
	{"%=", 1, true, air.KindAssign, "remainder-assignment"},
	{"&=", 1, true, air.KindAssign, "and-assignment"},
	{"|=", 1, true, air.KindAssign, "or-assignment"},
	{"^=", 1, true, air.KindAssign, "xor-assignment"},
	{"<<=", 1, true, air.KindAssign, "shift-assignment"},
	{">>=", 1, true, air.KindAssign, "shift-assignment"},
	{">>>=", 1, true, air.KindAssign, "shift-assignment"},
	{"&&=", 1, true, air.KindAssign, "logical-and-assignment"},
	{"||=", 1, true, air.KindAssign, "logical-or-assignment"},
	{"??=", 1, true, air.KindAssign, "nullish-assignment"},
}

// jsLookupInfix 按精确 token 文本查找中缀规则.
func jsLookupInfix(text string) (jsBindRule, bool) {
	for i := range jsInfixRules {
		if jsInfixRules[i].text == text {
			return jsInfixRules[i], true
		}
	}
	return jsBindRule{}, false
}
