/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// Fixed integer limits and binding powers. No clock, float or random state is
// ever consulted by a reduction. Operator pyTargetPower binds tighter than the
// in operator, so a for-clause target list stops before the iterated
// expression. Operator pyAsPower binds tighter than the as operator, so a
// with-statement context manager stops before its as target.
const (
	pyMaxDepth     = 64
	pySuffixPower  = 90
	pyClausePower  = 9
	pyTernaryPower = 8
	pyTargetPower  = 31
	pyAsPower      = 27
)

// Container modes select one closing delimiter and one lowering note.
const (
	pyModeParen = iota
	pyModeList
	pyModeDict
	pyModeSet
)

// pyKeywords lists Python 3.8-3.12 reserved and soft keywords.
var pyKeywords = []string{
	"False", "None", "True", "and", "as", "assert", "async", "await",
	"break", "case", "class", "continue", "def", "del", "elif", "else",
	"except", "finally", "for", "from", "global", "if", "import", "in",
	"is", "lambda", "match", "nonlocal", "not", "or", "pass", "raise",
	"return", "try", "type", "while", "with", "yield",
}

// pyOperators lists operator spellings; frontkit.NormalizeTables sorts them by
// descending length and then lexicographically for deterministic longest match.
var pyOperators = []string{
	"**=", "//=", ">>=", "<<=", "...", "==", "!=", "<=", ">=", "->",
	"+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "@=", ":=", "**",
	"//", "<<", ">>", "+", "-", "*", "/", "%", "&", "|", "^", "~",
	"<", ">", "=",
}

// pyStringQuotes puts triple quotes first so docstrings are one string token.
var pyStringQuotes = []string{`"""`, `'''`, `"`, `'`}

// pyRawPrefixes covers raw, formatted, bytes and unicode literal prefixes.
var pyRawPrefixes = []string{
	"fr", "Fr", "fR", "FR", "rf", "rF", "Rf", "RF",
	"br", "Br", "bR", "BR", "rb", "rB", "Rb", "RB",
	"r", "R", "f", "F", "b", "B", "u", "U",
}

var pyNumberSeparators = []string{"_"}

// pyDialects lists the accepted canonical Python dialects.
var pyDialects = []string{"3.8", "3.9", "3.10", "3.11", "3.12", "3.13"}

const pyDefaultDialect = "3.12"

var pyDialectPrefixes = []string{"cpython", "python", "py"}

// pyDirectivePrefixes marks comment spellings lowered to KindDirective.
var pyDirectivePrefixes = []string{
	"# type:", "# noqa", "# pragma:", "# pylint:", "# mypy:", "# pyright:", "#!",
}

var pyFStringPrefixes = []string{
	"f", "F", "fr", "Fr", "fR", "FR", "rf", "rF", "Rf", "RF",
}

// pyGenericNames lists built-in generic types lowered to KindGenericInst.
var pyGenericNames = []string{
	"Optional", "Union", "list", "dict", "set", "tuple", "frozenset",
	"Sequence", "Mapping", "Iterable", "Callable", "Type", "Generic",
	"Literal", "Annotated", "Final", "ClassVar", "Iterator",
}

var pyAssignOps = []string{
	"=", ":=", "+=", "-=", "*=", "/=", "//=", "%=", "**=", "&=",
	"|=", "^=", ">>=", "<<=", "@=",
}

var pyClauseKeywords = []string{"for", "if", "async"}
var pySuffixTexts = []string{".", "(", "["}
var pyCloseTexts = []string{")", "]", "}"}
var pyCollectionNotes = []string{"tuple", "list", "dict", "set"}
var pyComprehensionNotes = []string{
	"generator", "comprehension", "dict-comprehension", "set-comprehension",
}

// pyBinding is one deterministic infix binding power entry.
type pyBinding struct {
	text  string
	power int
	right bool
}

// pyInfixBindings is the Python Pratt precedence table in declaration order.
var pyInfixBindings = []pyBinding{
	{"or", 10, false}, {"and", 20, false}, {"import", 25, false},
	{"as", 26, false}, {"in", 30, false}, {"is", 30, false},
	{"==", 30, false}, {"!=", 30, false}, {"<", 30, false},
	{"<=", 30, false}, {">", 30, false}, {">=", 30, false},
	{"|", 40, false}, {"^", 45, false}, {"&", 50, false},
	{"<<", 55, false}, {">>", 55, false}, {"+", 60, false},
	{"-", 60, false}, {"*", 70, false}, {"/", 70, false},
	{"//", 70, false}, {"%", 70, false}, {"@", 70, false},
	{"**", 80, true}, {":=", 5, true}, {"if", pyTernaryPower, true},
}

// pyStmtRule maps one leading simple-statement keyword to its AIR kind.
type pyStmtRule struct {
	keyword string
	kind    air.NodeKind
	note    string
}

var pySimpleStmtRules = []pyStmtRule{
	{"return", air.KindReturn, ""},
	{"yield", air.KindUnary, "yield"},
	{"raise", air.KindThrow, ""},
	{"break", air.KindBranch, "break"},
	{"continue", air.KindBranch, "continue"},
	{"pass", air.KindDeclStmt, "pass"},
	{"del", air.KindUnary, "del"},
	{"assert", air.KindCall, "assert"},
	{"global", air.KindDeclStmt, "global"},
	{"nonlocal", air.KindDeclStmt, "nonlocal"},
	{"import", air.KindImport, ""},
	{"from", air.KindImport, "from-import"},
}

// pyHeaderRule maps one compound-statement keyword to its AIR kind.
type pyHeaderRule struct {
	keyword string
	kind    air.NodeKind
	note    string
}

var pyHeaderRules = []pyHeaderRule{
	{"if", air.KindIf, ""},
	{"elif", air.KindIf, "elif"},
	{"else", air.KindIf, "else"},
	{"while", air.KindWhile, ""},
	{"for", air.KindRangeFor, ""},
	{"with", air.KindWith, ""},
	{"try", air.KindTry, ""},
	{"except", air.KindCatch, ""},
	{"finally", air.KindFinally, ""},
	{"match", air.KindMatch, ""},
	{"case", air.KindCase, ""},
}

// pyTokenKind binds one frozen frontkit token kind to its AIR projection kind.
type pyTokenKind struct {
	token frontkit.TokenKind
	node  air.NodeKind
}

var pyTokenKinds = []pyTokenKind{
	{frontkit.TokenIdent, air.KindIdent},
	{frontkit.TokenNum, air.KindNumLit},
	{frontkit.TokenStr, air.KindStrLit},
	{frontkit.TokenChar, air.KindCharLit},
	{frontkit.TokenComment, air.KindComment},
}

// pyNote is one named semantic note value; mutation tests replace the table.
type pyNote struct {
	name  string
	value string
}

var pyNotes = []pyNote{
	{"comprehension", "comprehension"},
	{"comprehension-clause", "comprehension-clause"},
	{"generator", "generator-comprehension"},
	{"dict-comprehension", "dict-comprehension"},
	{"set-comprehension", "set-comprehension"},
	{"list", "list-literal"},
	{"tuple", "tuple-literal"},
	{"dict", "dict-literal"},
	{"set", "set-literal"},
	{"fstring", "fstring"},
	{"fstring-part", "fstring-interpolation"},
	{"indent", "indent-based"},
	{"walrus", "walrus"},
	{"decorator", "decorator"},
	{"yield", "yield"},
	{"del", "del"},
	{"assert", "assert"},
	{"pass", "pass"},
	{"global", "global"},
	{"nonlocal", "nonlocal"},
	{"slice", "slice"},
	{"optional", "optional"},
	{"from-import", "from-import"},
	{"elif", "elif"},
	{"else", "else"},
	{"break", "break"},
	{"continue", "continue"},
	{"dict-pair", "dict-pair"},
	{"ellipsis", "ellipsis"},
	{"type-annotation", "type-annotation"},
	{"docstring", "docstring"},
	{"star-args", "star-args"},
	{"star-list", "star-list"},
	{"kwargs", "kwargs"},
}

// pyInfixPower resolves one operator binding power from an injected table.
func pyInfixPower(table []pyBinding, text string) int {
	for i := range table {
		if table[i].text == text {
			return table[i].power
		}
	}
	return 0
}

// pyInfixRight reports whether one operator associates to the right.
func pyInfixRight(table []pyBinding, text string) bool {
	for i := range table {
		if table[i].text == text {
			return table[i].right
		}
	}
	return false
}

// pySimpleStmtRule resolves one simple-statement keyword.
func pySimpleStmtRule(keyword string) (pyStmtRule, bool) {
	for i := range pySimpleStmtRules {
		if pySimpleStmtRules[i].keyword == keyword {
			return pySimpleStmtRules[i], true
		}
	}
	return pyStmtRule{}, false
}

// pyHeaderKind resolves one compound-statement keyword.
func pyHeaderKind(keyword string) (pyHeaderRule, bool) {
	for i := range pyHeaderRules {
		if pyHeaderRules[i].keyword == keyword {
			return pyHeaderRules[i], true
		}
	}
	return pyHeaderRule{}, false
}

// pyTokenNodeKind projects one frozen token kind onto an AIR node kind.
func pyTokenNodeKind(kind frontkit.TokenKind) air.NodeKind {
	for i := range pyTokenKinds {
		if pyTokenKinds[i].token == kind {
			return pyTokenKinds[i].node
		}
	}
	return air.KindRegion
}

// pyIsLexicalKind reports whether one token has a byte-exact AIR projection.
func pyIsLexicalKind(kind frontkit.TokenKind) bool {
	return pyTokenNodeKind(kind) != air.KindRegion
}

// pyNoteFor resolves one named semantic note from an injected table.
func pyNoteFor(table []pyNote, name string) string {
	for i := range table {
		if table[i].name == name {
			return table[i].value
		}
	}
	return ""
}
