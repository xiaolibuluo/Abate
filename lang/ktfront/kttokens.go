/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import "abate/lang/air"

// 归约深度、绑定力与默认方言常量，全部为整数或字符串，不涉及时钟、浮点与随机状态。
const (
	ktMaxDepth        = 48
	ktWhenPrefixPower = 6
	ktDefaultDialect  = "kt2.1"
)

// ktScope 区分翻译单元、类型体与语句体，决定成员归约类别。
type ktScope int

const (
	ktScopeTop ktScope = iota
	ktScopeType
	ktScopeBody
)

// ktKeywords 是 Kotlin 硬关键字与常用修饰符表；反引号标识符与字符串模板
// 由 stringQuotes 与归约期识别，不列入本表。
var ktKeywords = []string{
	"abstract", "actual", "annotation", "as", "break", "by", "catch", "class",
	"companion", "const", "constructor", "continue", "crossinline", "data",
	"delegate", "do", "dynamic", "else", "enum", "expect", "external", "false",
	"field", "file", "final", "finally", "for", "fun", "get", "if", "import",
	"in", "infix", "init", "inline", "inner", "interface", "internal", "is",
	"lateinit", "noinline", "null", "object", "open", "operator", "out",
	"override", "package", "param", "private", "property", "protected", "public",
	"receiver", "reified", "return", "sealed", "set", "setparam", "super",
	"suspend", "tailrec", "this", "throw", "true", "try", "typealias", "typeof",
	"val", "value", "var", "vararg", "when", "where", "while",
}

// ktOperators 是 Kotlin 操作符表；frontkit.NormalizeTables 按长度降序归一化，
// 保证 ?. ?: !! :: -> ..< 等复合操作符先于单字符匹配。
var ktOperators = []string{
	"..<", "===", "!==", "?.", "?:", "!!", "::", "->", "..", "++", "--",
	"+=", "-=", "*=", "/=", "%=", "==", "!=", "<=", ">=", "&&", "||",
	"+", "-", "*", "/", "%", "=", "<", ">", "!", "&", "|", "^", "~",
	"?", ".", ",", ":", "@", "$",
}

// ktStringQuotes 依次覆盖三引号原始字符串、普通字符串模板、字符字面量与反引号标识符。
var ktStringQuotes = []string{`"""`, `"`, `'`, "`"}
var ktRawStringPrefixes = []string{`"""`}
var ktNumberSeparators = []string{"_"}

const (
	ktLineComment       = "//"
	ktBlockCommentOpen  = "/*"
	ktBlockCommentClose = "*/"
)

// ktModifierWords 是定位声明主关键字时跳过的修饰符表。
var ktModifierWords = []string{
	"public", "private", "protected", "internal", "abstract", "final", "open",
	"override", "suspend", "inline", "noinline", "crossinline", "tailrec",
	"operator", "infix", "external", "lateinit", "const", "vararg", "data",
	"sealed", "inner", "enum", "annotation", "companion", "expect", "actual",
	"reified", "value",
}

// ktTypeWords 是内建类型拼写；仅用于挑选声明名时跳过，词法投影仍为 KindIdent。
var ktTypeWords = []string{
	"Int", "Long", "Short", "Byte", "Float", "Double", "Boolean", "Char",
	"String", "Unit", "Any", "Nothing", "Array", "List", "Map", "Set",
	"MutableList", "MutableMap", "Sequence", "Pair", "Triple", "Result",
	"Throwable",
}

// ktSoftContextWords 是挑选声明名时跳过的上下文关键字与隐式接收者。
var ktSoftContextWords = []string{
	"where", "by", "get", "set", "in", "is", "as", "it", "this", "field",
	"value", "out",
}

// ktDeclRule 是一条「主关键字 → AIR 种类与语义」的确定性映射。
type ktDeclRule struct {
	lead     string
	kind     air.NodeKind
	note     string
	semKey   air.SemKey
	semValue string
}

// ktDeclRules 是 Kotlin 构造到冻结 AIR 枚举的映射表，只用 kinds.go 真实常量。
var ktDeclRules = []ktDeclRule{
	{"package", air.KindModuleDecl, "package", 0, ""},
	{"import", air.KindImport, "import", 0, ""},
	{"class", air.KindClassDecl, "", 0, ""},
	{"interface", air.KindInterfaceDecl, "", 0, ""},
	{"object", air.KindClassDecl, "object", 0, ""},
	{"fun", air.KindFuncDecl, "", air.SemPurity, ""},
	{"val", air.KindVarDecl, "val-binding", 0, ""},
	{"var", air.KindVarDecl, "var-binding", 0, ""},
	{"when", air.KindMatch, "when", 0, ""},
	{"if", air.KindIf, "", 0, ""},
	{"else", air.KindIf, "else", 0, ""},
	{"for", air.KindRangeFor, "for-in", 0, ""},
	{"while", air.KindWhile, "", 0, ""},
	{"do", air.KindDoWhile, "", 0, ""},
	{"return", air.KindReturn, "", 0, ""},
	{"throw", air.KindThrow, "", 0, ""},
	{"break", air.KindBranch, "break", 0, ""},
	{"continue", air.KindBranch, "continue", 0, ""},
	{"try", air.KindTry, "", 0, ""},
	{"catch", air.KindCatch, "", 0, ""},
	{"finally", air.KindFinally, "", 0, ""},
	{"typealias", air.KindTypeDecl, "typealias", 0, ""},
	{"constructor", air.KindFuncDecl, "constructor", 0, ""},
	{"init", air.KindBlock, "init", 0, ""},
	{"get", air.KindFuncDecl, "getter", 0, ""},
	{"set", air.KindFuncDecl, "setter", 0, ""},
	{"is", air.KindTypeAssert, "is-check", 0, ""},
	{"as", air.KindTypeAssert, "as-cast", 0, ""},
	{"suspend", air.KindFuncDecl, "suspend", air.SemAsync, "true"},
}

// ktBinding 是一条 Pratt 中缀绑定力。
type ktBinding struct {
	text  string
	power int
	right bool
}

// ktInfixBindings 是 Kotlin 表达式的确定序 Pratt 绑定表；"?" 与 "!!" 是后缀
// 构造，交由词法叶子的专有投影处理，不注册为需要右操作数的中缀规则。
var ktInfixBindings = []ktBinding{
	{"=", 10, true}, {"+=", 10, true}, {"-=", 10, true}, {"*=", 10, true},
	{"/=", 10, true}, {"%=", 10, true},
	{"->", 8, true}, {"?:", 15, true}, {"||", 20, false}, {"&&", 25, false},
	{"===", 30, false}, {"!==", 30, false}, {"==", 30, false}, {"!=", 30, false},
	{"<", 30, false}, {">", 30, false}, {"<=", 30, false}, {">=", 30, false},
	{"|", 35, false}, {"^", 38, false}, {"&", 40, false},
	{"<<", 45, false}, {">>", 45, false}, {">>>", 45, false},
	{"..", 48, false}, {"..<", 48, false},
	{"+", 50, false}, {"-", 50, false}, {"*", 55, false}, {"/", 55, false},
	{"%", 55, false}, {"is", 58, false}, {"as", 58, false}, {"in", 12, false},
	{"?.", 92, false}, {".", 92, false}, {"::", 95, false},
	{"(", 94, false}, {"[", 94, false}, {",", 3, false},
}

// ktNote 是一条命名语义标注值。
type ktNote struct {
	name  string
	value string
}

// ktNotes 是 Kotlin 专有构造的降级标注表，全部落在 air.SemNote 真实枚举上；
// 表外名字由 ktNoteValue 回退到名字本身，保持确定性。
var ktNotes = []ktNote{
	{"null-safe", "null-safe"}, {"not-null-assert", "not-null-assert"},
	{"elvis", "elvis"}, {"nullable", "nullable"},
	{"when", "when"}, {"when-expr", "when-expr"}, {"when-branch", "when-branch"},
	{"suspend", "suspend"}, {"extension", "extension"},
	{"companion-object", "companion-object"}, {"data-class", "data-class"},
	{"sealed-class", "sealed-class"}, {"enum-class", "enum-class"},
	{"annotation-class", "annotation-class"}, {"const-val", "const-val"},
	{"val-binding", "val-binding"}, {"var-binding", "var-binding"},
	{"lambda", "lambda"}, {"lambda-with-receiver", "lambda-with-receiver"},
	{"string-template", "string-template"}, {"raw-string", "raw-string"},
	{"backtick-ident", "backtick-ident"}, {"template-entry", "template-entry"},
	{"package", "package"}, {"import", "import"}, {"object", "object"},
	{"for-in", "for-in"}, {"typealias", "typealias"},
	{"constructor", "constructor"}, {"init", "init"}, {"kdoc", "kdoc"},
	{"getter", "getter"}, {"setter", "setter"}, {"is-check", "is-check"},
	{"as-cast", "as-cast"}, {"operator", "operator"}, {"infix", "infix"},
	{"else", "else"}, {"break", "break"}, {"continue", "continue"},
}

// ktNullSafeRule 是一条「空安全操作符 → AIR 种类与标注名」的可注入映射。
type ktNullSafeRule struct {
	text string
	kind air.NodeKind
	note string
}

// ktNullSafeRules 是空安全映射关键字表，是变异测试的可注入表；删去条目后
// 对应 token 不再获得专有种类与 SemNote，AIR 哈希随之改变。
var ktNullSafeRules = []ktNullSafeRule{
	{"?.", air.KindField, "null-safe"},
	{"!!", air.KindUnary, "not-null-assert"},
	{"?:", air.KindBinary, "elvis"},
	{"?", air.KindOptional, "nullable"},
}

// ktWhenBranchWords 是 when 分支的切分标记，是变异测试的可注入表。
var ktWhenBranchWords = []string{"->", "else"}

// ktLambdaReceiverWords 是 lambda 隐式接收者的提示拼写。
var ktLambdaReceiverWords = []string{"this", "it"}

// ktDialect 是一条方言别名映射。
type ktDialect struct {
	alias     string
	canonical string
}

// ktDialectAliases 是有序别名表，绝不依赖 map 迭代序。
var ktDialectAliases = []ktDialect{
	{"", ktDefaultDialect}, {"kt", ktDefaultDialect},
	{"kotlin", ktDefaultDialect},
	{"1.6", "kt1.6"}, {"kt1.6", "kt1.6"},
	{"1.7", "kt1.7"}, {"kt1.7", "kt1.7"},
	{"1.8", "kt1.8"}, {"kt1.8", "kt1.8"},
	{"1.9", "kt1.9"}, {"kt1.9", "kt1.9"},
	{"2.0", "kt2.0"}, {"kt2.0", "kt2.0"},
	{"2.1", "kt2.1"}, {"kt2.1", "kt2.1"},
	{"latest", "kt2.1"},
}

var ktDialects = []string{
	"kt1.6", "kt1.7", "kt1.8", "kt1.9", "kt2.0", "kt2.1",
}
