/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import "abate/lang/air"

// 归约深度与绑定力常量，全部为整数，不涉及时钟、浮点与随机状态。
const (
	csMaxDepth        = 48
	csLinqPrefixPower = 6
	csDefaultDialect  = "cs12"
)

// csScope 区分翻译单元、类型体与语句体，决定成员归约类别。
type csScope int

const (
	csScopeTop csScope = iota
	csScopeType
	csScopeBody
)

// csKeywords 是 C# 保留字表；LINQ 上下文关键字刻意不列入，以便作为标识符
// 参与 token 保真投影，并在归约期由有序查询表识别。
var csKeywords = []string{
	"abstract", "add", "as", "async", "await", "base", "bool", "break", "byte",
	"case", "catch", "char", "checked", "class", "const", "continue", "decimal",
	"default", "delegate", "do", "double", "dynamic", "else", "enum", "event",
	"explicit", "extern", "false", "finally", "fixed", "float", "for", "foreach",
	"get", "goto", "if", "implicit", "in", "int", "interface", "internal", "is",
	"lock", "long", "nameof", "namespace", "new", "null", "object", "operator",
	"out", "override", "params", "partial", "private", "protected", "public",
	"readonly", "record", "ref", "remove", "return", "sbyte", "sealed", "set",
	"short", "sizeof", "stackalloc", "static", "string", "struct", "switch",
	"this", "throw", "true", "try", "typeof", "uint", "ulong", "unchecked",
	"unsafe", "ushort", "using", "value", "var", "virtual", "void", "volatile",
	"while", "yield",
}

// csOperators 是 C# 操作符表；frontkit.NormalizeTables 按长度降序归一化。
// "#" 与 "@" 不列入，井号指令与逐字标识符前缀交由归约期识别。
var csOperators = []string{
	"??=", "<<=", ">>=", ">>>=", "=>", "?.", "??", "++", "--", "->", "::",
	"==", "!=", "<=", ">=", "&&", "||", "+=", "-=", "*=", "/=", "%=", "&=",
	"|=", "^=", "<<", ">>", ">>>", "+", "-", "*", "/", "%", "&", "|", "^",
	"!", "~", "=", "<", ">", "?",
}

var csStringQuotes = []string{`"`, `'`}
var csRawStringPrefixes = []string{"$@", "@$", "$", "@"}
var csNumberSeparators = []string{"_"}

const (
	csLineComment       = "//"
	csBlockCommentOpen  = "/*"
	csBlockCommentClose = "*/"
)

// csLinqKeywords 是 LINQ 查询上下文关键字全集，是变异测试的可注入表。
var csLinqKeywords = []string{
	"from", "where", "select", "group", "orderby", "join", "let", "into",
	"in", "on", "equals", "by", "ascending", "descending",
}

// csLinqClauseKeywords 是查询子句起始关键字，用于切分 LINQ 区间。
var csLinqClauseKeywords = []string{
	"from", "where", "select", "group", "orderby", "join", "let", "into",
}

// csLinqStartKeywords 是能开启一个查询表达式的关键字。
var csLinqStartKeywords = []string{"from"}

// csModifierWords 是定位声明首关键字时跳过的修饰符表。
var csModifierWords = []string{
	"public", "private", "protected", "internal", "static", "readonly",
	"virtual", "override", "abstract", "sealed", "partial", "async", "new",
	"extern", "volatile", "ref", "out", "in", "params", "required",
}

// csTypeWords 是内建类型拼写，词法投影为 KindNamed。
var csTypeWords = []string{
	"bool", "byte", "sbyte", "char", "decimal", "double", "dynamic", "float",
	"int", "uint", "long", "ulong", "object", "short", "ushort", "string",
	"var", "void",
}

// csDeclRule 是一条「首关键字 → AIR 种类与语义」的确定性映射。
type csDeclRule struct {
	lead     string
	kind     air.NodeKind
	note     string
	semKey   air.SemKey
	semValue string
}

// csDeclRules 是 C# 构造到冻结 AIR 枚举的映射表，只用 kinds.go 真实常量。
var csDeclRules = []csDeclRule{
	{"namespace", air.KindModuleDecl, "namespace", 0, ""},
	{"using", air.KindImport, "using-directive", 0, ""},
	{"class", air.KindClassDecl, "", 0, ""},
	{"struct", air.KindStructDecl, "", 0, ""},
	{"interface", air.KindInterfaceDecl, "", 0, ""},
	{"enum", air.KindEnumDecl, "", 0, ""},
	{"record", air.KindClassDecl, "record", 0, ""},
	{"delegate", air.KindFuncDecl, "delegate", 0, ""},
	{"event", air.KindVarDecl, "event", 0, ""},
	{"const", air.KindConstDecl, "", 0, ""},
	{"return", air.KindReturn, "", 0, ""},
	{"yield", air.KindReturn, "yield", 0, ""},
	{"throw", air.KindThrow, "", 0, ""},
	{"if", air.KindIf, "", 0, ""},
	{"else", air.KindIf, "else", 0, ""},
	{"for", air.KindFor, "", 0, ""},
	{"foreach", air.KindRangeFor, "foreach", 0, ""},
	{"while", air.KindWhile, "", 0, ""},
	{"do", air.KindDoWhile, "", 0, ""},
	{"switch", air.KindSwitch, "", 0, ""},
	{"case", air.KindCase, "", 0, ""},
	{"default", air.KindCase, "default", 0, ""},
	{"try", air.KindTry, "", 0, ""},
	{"catch", air.KindCatch, "", 0, ""},
	{"finally", air.KindFinally, "", 0, ""},
	{"break", air.KindBranch, "break", 0, ""},
	{"continue", air.KindBranch, "continue", 0, ""},
	{"goto", air.KindBranch, "goto", 0, ""},
	{"lock", air.KindWith, "lock", 0, ""},
	{"checked", air.KindBlock, "checked", 0, ""},
	{"unchecked", air.KindBlock, "unchecked", 0, ""},
	{"unsafe", air.KindBlock, "unsafe", 0, ""},
	{"fixed", air.KindBlock, "fixed", 0, ""},
	{"await", air.KindAwait, "", air.SemAsync, "true"},
	{"new", air.KindCall, "object-creation", 0, ""},
	{"typeof", air.KindTypeAssert, "typeof", 0, ""},
	{"nameof", air.KindCall, "nameof", 0, ""},
}

// csBinding 是一条 Pratt 中缀绑定力。
type csBinding struct {
	text  string
	power int
	right bool
}

// csInfixBindings 是 C# 表达式的确定序 Pratt 绑定表。
var csInfixBindings = []csBinding{
	{"=", 10, true}, {"+=", 10, true}, {"-=", 10, true}, {"*=", 10, true},
	{"/=", 10, true}, {"%=", 10, true}, {"&=", 10, true}, {"|=", 10, true},
	{"^=", 10, true}, {"<<=", 10, true}, {">>=", 10, true}, {"??=", 10, true},
	{"=>", 8, true}, {"??", 15, true}, {"||", 20, false}, {"&&", 25, false},
	{"==", 30, false}, {"!=", 30, false}, {"<", 30, false}, {">", 30, false},
	{"<=", 30, false}, {">=", 30, false}, {"|", 35, false}, {"^", 38, false},
	{"&", 40, false}, {"<<", 45, false}, {">>", 45, false}, {"+", 50, false},
	{"-", 50, false}, {"*", 55, false}, {"/", 55, false}, {"%", 55, false},
	{"is", 58, false}, {"as", 58, false}, {"?", 9, true}, {":", 9, true},
	{"in", 12, false}, {"?.", 92, false}, {".", 92, false},
	{"(", 94, false}, {"[", 94, false}, {"::", 95, false}, {",", 3, false},
}

// csNote 是一条命名语义标注值。
type csNote struct {
	name  string
	value string
}

// csNotes 是 C# 专有构造的降级标注表，全部落在 air.SemNote 真实枚举上。
var csNotes = []csNote{
	{"linq", "linq"}, {"linq-clause", "linq-clause"},
	{"property", "property"}, {"property-accessors", "property-accessors"},
	{"record", "record"}, {"indexer", "indexer"}, {"event", "event"},
	{"delegate", "delegate"}, {"attribute", "attribute"},
	{"nullable", "nullable"}, {"null-conditional", "null-conditional"},
	{"null-coalescing", "null-coalescing"},
	{"null-coalescing-assign", "null-coalescing-assign"},
	{"lambda", "lambda"}, {"interpolated", "interpolated"},
	{"verbatim", "verbatim"}, {"verbatim-interpolated", "verbatim-interpolated"},
	{"interpolation", "interpolation"}, {"region", "region"},
	{"endregion", "endregion"}, {"directive", "directive"},
	{"xml-doc", "xml-doc"}, {"namespace", "namespace"},
	{"using-directive", "using-directive"}, {"using-statement", "using-statement"},
	{"foreach", "foreach"}, {"else", "else"}, {"default", "default"},
	{"yield", "yield"}, {"break", "break"}, {"continue", "continue"},
	{"goto", "goto"}, {"lock", "lock"}, {"checked", "checked"},
	{"unchecked", "unchecked"}, {"unsafe", "unsafe"}, {"fixed", "fixed"},
	{"object-creation", "object-creation"}, {"typeof", "typeof"},
	{"nameof", "nameof"},
}

// csDialect 是一条方言别名映射。
type csDialect struct {
	alias     string
	canonical string
}

// csDialectAliases 是有序别名表，绝不依赖 map 迭代序。
var csDialectAliases = []csDialect{
	{"", "cs12"}, {"cs", "cs12"}, {"csharp", "cs12"}, {"c#", "cs12"},
	{"cs8", "cs8"}, {"8.0", "cs8"}, {"c#8", "cs8"}, {"c# 8", "cs8"},
	{"cs9", "cs9"}, {"9.0", "cs9"}, {"c#9", "cs9"}, {"c# 9", "cs9"},
	{"cs10", "cs10"}, {"10.0", "cs10"}, {"c#10", "cs10"}, {"c# 10", "cs10"},
	{"cs11", "cs11"}, {"11.0", "cs11"}, {"c#11", "cs11"}, {"c# 11", "cs11"},
	{"cs12", "cs12"}, {"12.0", "cs12"}, {"c#12", "cs12"}, {"c# 12", "cs12"},
	{"latest", "cs12"}, {"net6", "cs10"}, {"net7", "cs11"}, {"net8", "cs12"},
}

var csDialects = []string{"cs8", "cs9", "cs10", "cs11", "cs12"}
