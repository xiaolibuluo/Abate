/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// Fixed integer limits and binding powers. No clock, float or random state is
// ever consulted by a reduction, so identical bytes always produce identical AIR.
const (
	rsMaxDepth      = 48
	rsPostfixPower  = 95
	rsSelectorPower = 90
	rsPrefixPower   = 80
	rsRangePower    = 20
)

// rsContext selects the construct rule set for one token interval.
type rsContext int

// The context values distinguish translation-unit items, block statements,
// impl/trait members and match arms.
const (
	rsCtxTop rsContext = iota
	rsCtxBlock
	rsCtxImpl
	rsCtxTrait
	rsCtxMatchArm
	rsCtxStructLit
	rsCtxEnum
)

// rsBinding is one deterministic infix binding-power entry.
type rsBinding struct {
	text  string
	power int
	right bool
}

// rsNote is one named semantic note value; mutation tests replace the table.
type rsNote struct {
	name  string
	value string
}

// rsItemRule maps one leading keyword to its frozen AIR lowering.
type rsItemRule struct {
	keyword  string
	kind     air.NodeKind
	noteName string
	semKey   air.SemKey
	semValue string
}

// rsTokenKind binds one frozen frontkit token kind to an AIR node kind.
type rsTokenKind struct {
	token frontkit.TokenKind
	node  air.NodeKind
}

// rsDialectAlias maps one accepted spelling to one canonical edition.
type rsDialectAlias struct {
	alias     string
	canonical string
}

// RsKeywords lists Rust 2015-2024 reserved words, including macro_rules.
// A macro definition header then stays one keyword token instead of two.
var rsKeywords = []string{
	"as", "async", "await", "break", "const", "continue", "crate", "dyn", "else", "enum",
	"extern", "false", "fn", "for", "if", "impl", "in", "let", "loop", "macro_rules",
	"match", "mod", "move", "mut", "pub", "ref", "return", "self", "Self", "static",
	"struct", "super", "trait", "true", "type", "unsafe", "use", "where", "while",
	"abstract", "become", "box", "do", "final", "macro", "override", "priv", "typeof",
	"unsized", "virtual", "yield", "try", "union", "default",
}

// rsOperators holds Rust operator spellings. The single quote is an operator,
// not a string delimiter, so a lifetime such as 'a stays a two-token sequence;
// char literals are recombined by the reducer. NormalizeTables sorts longest
// first, so the compound spellings win over their single-character prefixes.
var rsOperators = []string{
	">>=", "<<=", "..=", "...", "::", "->", "=>", "==", "!=", "<=", ">=", "&&", "||",
	"<<", ">>", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "..", "'", "&", "|",
	"^", "+", "-", "*", "/", "%", "<", ">", "=", "!", "?", "@", "~",
}

// rsStringQuotes holds the ordinary double-quoted string delimiter. The single
// quote is deliberately absent so lifetimes are never swallowed by the lexer.
var rsStringQuotes = []string{`"`}

// rsRawPrefixes covers raw strings r"..", r#".."#, br".." and c".." plus their
// hash variants. Each prefix combines with the string quote in quoteMatch.
var rsRawPrefixes = []string{
	"br##", "cr##", "r###", "br#", "cr#", "r##", "br", "cr", "r#",
	"r", "b", "c",
}

var rsNumberSeparators = []string{"_"}

// rsVisibilityModifiers are skipped when resolving an item's leading keyword.
var rsVisibilityModifiers = []string{"pub", "default"}

// rsOpenBrackets and rsCloseBrackets pair macro and container delimiters.
var rsOpenBrackets = []string{"(", "[", "{"}
var rsCloseBrackets = []string{")", "]", "}"}

// rsAssignOps are the assignment spellings lowered to KindAssign.
var rsAssignOps = []string{
	"=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "<<=", ">>=",
}

// rsMacroNames lists well-known macros; a matching call receives SemTag.
var rsMacroNames = []string{
	"macro_rules", "println", "print", "eprintln", "format", "vec", "assert",
	"assert_eq", "panic", "write", "writeln", "matches", "dbg", "include",
	"stringify", "concat", "cfg",
}

// rsDirectivePrefixes marks comment spellings lowered to KindDirective.
var rsDirectivePrefixes = []string{
	"// clippy:", "// rustfmt:", "// allow:", "// deny:", "// warn:", "// cfg:",
}

// rsDocPrefixes marks doc comments lowered to KindDoc.
var rsDocPrefixes = []string{"///", "//!", "/**"}

// rsDialects lists the accepted canonical Rust editions.
var rsDialects = []string{"2015", "2018", "2021", "2024"}

const rsDefaultDialect = "2021"

var rsDialectPrefixes = []string{"rust", "edition", "rs"}

var rsDialectAliases = []rsDialectAlias{
	{"2015", "2015"}, {"2018", "2018"}, {"2021", "2021"}, {"2024", "2024"}, {"stable", "2021"}, {"beta", "2021"}, {"nightly", "2021"},
}

// rsInfixBindings is the Rust Pratt precedence table in declaration order.
var rsInfixBindings = []rsBinding{
	{"=", 5, true}, {"+=", 5, true}, {"-=", 5, true}, {"*=", 5, true}, {"/=", 5, true}, {"%=", 5, true}, {"&=", 5, true}, {"|=", 5, true},
	{"^=", 5, true}, {"<<=", 5, true}, {">>=", 5, true}, {"=>", 1, false}, {"->", 1, false}, {"..", rsRangePower, false}, {"..=", rsRangePower, false},
	{"||", 25, false}, {"&&", 30, false}, {"as", 34, false},
	{"==", 35, false}, {"!=", 35, false}, {"<", 35, false}, {"<=", 35, false}, {">", 35, false}, {">=", 35, false},
	{"|", 40, false}, {"^", 42, false}, {"&", 45, false}, {"<<", 50, false}, {">>", 50, false},
	{"+", 60, false}, {"-", 60, false}, {"*", 70, false}, {"/", 70, false}, {"%", 70, false},
}

// rsItemRules is the item and statement keyword to AIR lowering table.
var rsItemRules = []rsItemRule{
	{"use", air.KindImport, "", air.SemInvalid, ""}, {"mod", air.KindModuleDecl, "", air.SemInvalid, ""},
	{"extern", air.KindImport, "extern", air.SemInvalid, ""}, {"fn", air.KindFuncDecl, "", air.SemInvalid, ""},
	{"struct", air.KindStructDecl, "", air.SemInvalid, ""}, {"enum", air.KindEnumDecl, "", air.SemInvalid, ""},
	{"trait", air.KindTraitDecl, "", air.SemInvalid, ""}, {"impl", air.KindBlock, "impl", air.SemInvalid, ""},
	{"type", air.KindTypeDecl, "", air.SemInvalid, ""}, {"const", air.KindConstDecl, "", air.SemInvalid, ""},
	{"static", air.KindVarDecl, "static", air.SemInvalid, ""}, {"macro_rules", air.KindCall, "macro", air.SemInvalid, ""},
	{"unsafe", air.KindBlock, "unsafe", air.SemInvalid, ""}, {"async", air.KindFuncDecl, "", air.SemAsync, "true"},
	{"let", air.KindVarDecl, "", air.SemInvalid, ""}, {"return", air.KindReturn, "", air.SemInvalid, ""},
	{"break", air.KindBranch, "break", air.SemInvalid, ""}, {"continue", air.KindBranch, "continue", air.SemInvalid, ""},
	{"if", air.KindIf, "", air.SemInvalid, ""}, {"while", air.KindWhile, "", air.SemInvalid, ""},
	{"for", air.KindRangeFor, "", air.SemInvalid, ""}, {"loop", air.KindWhile, "loop", air.SemInvalid, ""},
	{"match", air.KindMatch, "", air.SemInvalid, ""}, {"move", air.KindLambda, "move", air.SemInvalid, ""},
	{"dyn", air.KindNamed, "dyn", air.SemInvalid, ""}, {"where", air.KindGenericParam, "where", air.SemInvalid, ""},
	{"box", air.KindUnary, "box", air.SemInvalid, ""}, {"await", air.KindAwait, "", air.SemAsync, "true"},
	{"crate", air.KindImport, "extern-crate", air.SemInvalid, ""}, {"union", air.KindStructDecl, "union", air.SemInvalid, ""},
	{"yield", air.KindUnary, "yield", air.SemInvalid, ""}, {"try", air.KindTry, "", air.SemInvalid, ""},
}

// rsNotes is the named semantic-note table injected into every reducer.
var rsNotes = []rsNote{
	{"macro", "macro"}, {"macro-rules", "macro-rules"}, {"impl", "impl"}, {"unsafe", "unsafe"}, {"extern", "extern"}, {"extern-abi", "extern-abi"},
	{"extern-crate", "extern-crate"}, {"static", "static"}, {"loop", "loop"}, {"move", "move"}, {"dyn", "dyn"}, {"where", "where"}, {"box", "box"},
	{"break", "break"}, {"continue", "continue"}, {"if-let", "if-let"}, {"while-let", "while-let"}, {"match-arm", "match-arm"},
	{"try-op", "try-op"}, {"range", "range"}, {"lifetime", "lifetime"}, {"array", "array"}, {"struct-literal", "struct-literal"},
	{"tuple", "tuple"}, {"pattern", "pattern"}, {"inner-attribute", "inner-attribute"}, {"generic", "generic"},
	{"closure", "closure"}, {"await", "await"}, {"ref", "ref"}, {"known-macro", "known-macro"}, {"directive", "directive"},
	{"mutable", "mutable"}, {"const-fn", "const-fn"}, {"union", "union"}, {"yield", "yield"}, {"method", "method"},
	{"char-literal", "char-literal"}, {"doc-comment", "doc-comment"},
}

// rsTokenKinds projects the frozen lexical kinds onto frozen AIR kinds.
var rsTokenKinds = []rsTokenKind{
	{frontkit.TokenIdent, air.KindIdent},
	{frontkit.TokenNum, air.KindNumLit},
	{frontkit.TokenStr, air.KindStrLit},
	{frontkit.TokenChar, air.KindCharLit},
	{frontkit.TokenComment, air.KindComment},
}

// rsContains reports whether text is one entry of an ordered table.
func rsContains(table []string, text string) bool {
	for i := range table {
		if table[i] == text {
			return true
		}
	}
	return false
}

// rsHasPrefix reports whether text starts with one entry of an ordered table.
func rsHasPrefix(text string, prefixes []string) bool {
	for i := range prefixes {
		if strings.HasPrefix(text, prefixes[i]) {
			return true
		}
	}
	return false
}

// rsNoteFor resolves one named semantic note from an injected table.
func rsNoteFor(table []rsNote, name string) string {
	for i := range table {
		if table[i].name == name {
			return table[i].value
		}
	}
	return ""
}

// rsInfixPower resolves one operator binding power from an injected table.
func rsInfixPower(table []rsBinding, text string) int {
	for i := range table {
		if table[i].text == text {
			return table[i].power
		}
	}
	return 0
}

// rsInfixRight reports whether one operator associates to the right.
func rsInfixRight(table []rsBinding, text string) bool {
	for i := range table {
		if table[i].text == text {
			return table[i].right
		}
	}
	return false
}

// rsItemRuleFor resolves one item or statement keyword.
func rsItemRuleFor(table []rsItemRule, keyword string) (rsItemRule, bool) {
	for i := range table {
		if table[i].keyword == keyword {
			return table[i], true
		}
	}
	return rsItemRule{}, false
}

// rsTokenNodeKind projects one frozen token kind onto an AIR node kind.
func rsTokenNodeKind(kind frontkit.TokenKind) air.NodeKind {
	for i := range rsTokenKinds {
		if rsTokenKinds[i].token == kind {
			return rsTokenKinds[i].node
		}
	}
	return air.KindRegion
}

// rsIsLexicalKind reports whether one token has a byte-exact AIR projection.
func rsIsLexicalKind(kind frontkit.TokenKind) bool {
	return rsTokenNodeKind(kind) != air.KindRegion
}

// rsIsTrivia reports whether one token is a layout marker the reducer skips.
func rsIsTrivia(kind frontkit.TokenKind) bool {
	return kind == frontkit.TokenNewline || kind == frontkit.TokenIndent ||
		kind == frontkit.TokenDedent
}

// rsIsOpenBracket reports whether text opens a macro or container delimiter.
func rsIsOpenBracket(text string) bool { return rsContains(rsOpenBrackets, text) }

// rsCloseOf returns the matching closing delimiter for one opening bracket.
func rsCloseOf(open string) string {
	for i := range rsOpenBrackets {
		if rsOpenBrackets[i] == open {
			return rsCloseBrackets[i]
		}
	}
	return ")"
}

// rsNormalizeDialectKey lowercases a dialect and strips its edition prefixes.
func rsNormalizeDialectKey(d string) string {
	key := strings.ToLower(strings.TrimSpace(d))
	for _, prefix := range rsDialectPrefixes {
		key = strings.TrimSpace(strings.TrimPrefix(key, prefix))
	}
	return key
}
