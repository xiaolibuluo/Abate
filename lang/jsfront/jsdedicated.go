/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file carries the genuine dedicated-kind reductions of jsfront. Every
// marker below is emitted for a construct the parser actually recognized, and
// is anchored on a real token (or a reduced node) of that construct. A marker
// is a zero-width node: it never rewrites an existing projection, so lexical
// fidelity and every existing AIR-ID stay byte-exact. This mirrors the
// dedicated-marker pattern already proven in csfront/rsfront/ktfront.

// dedicated anchors a zero-width dedicated-kind marker and tags it with the
// tree-sitter construct name that produced it, so the emission is auditable.
func (m *jsMapper) dedicated(kind air.NodeKind, pos air.Position,
	text, note string) air.Node {
	zero := air.NewPosition(pos.Line, pos.Column, pos.Offset, pos.Offset)
	node := air.NewNodeAt(kind, zero, zero, text)
	return m.annotate(node, air.SemNote, note)
}

// tokenPos converts a real token into a zero-width anchor position.
func tokenPos(tok frontkit.Token) air.Position {
	return air.NewPosition(tok.Line, tok.Column, tok.Offset, tok.Offset)
}

// jsIsAugAssign reports whether text is a compound-assignment operator that
// jsfront binds as an infix rule (every multi-character operator ending in
// '=' except equality, relational and arrow tokens).
func jsIsAugAssign(text string) bool {
	if len(text) < 2 || text[len(text)-1] != '=' {
		return false
	}
	switch text {
	case "==", "!=", "<=", ">=", "===", "!==", "=>":
		return false
	default:
		return true
	}
}

// jsIsFormalParamsOpen reports whether the '(' at index i opens a function's
// formal-parameter list. It looks back over an optional 'function' keyword and
// the function name (or a '*' generator marker), so a call's '(' is excluded.
func jsIsFormalParamsOpen(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Text != "(" {
		return false
	}
	prev := toks[i-1]
	if prev.Kind != frontkit.TokenIdent && prev.Text != "*" {
		return false
	}
	// Walk back to find the nearest 'function' keyword before the name.
	j := i - 1
	for j >= 0 && j >= i-4 {
		if toks[j].Text == "function" {
			return true
		}
		if toks[j].Kind == frontkit.TokenNewline ||
			toks[j].Text == ";" || toks[j].Text == "{" || toks[j].Text == "}" {
			break
		}
		j--
	}
	return false
}

// jsIsDeclaratorIdent reports whether the identifier at index i is the name of
// a variable declarator introduced by let/const/var on the same statement.
func jsIsDeclaratorIdent(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Kind != frontkit.TokenIdent {
		return false
	}
	prev := toks[i-1]
	if prev.Text != "const" && prev.Text != "let" && prev.Text != "var" {
		return false
	}
	return prev.Kind == frontkit.TokenKeyword
}

// keywordMarker emits the dedicated marker for one keyword token. It returns
// false when the token is not a dedicated-construct keyword, so the caller can
// fall through to the symbol/context dispatcher.
func (p *jsParser) keywordMarker(out *[]air.Node, i int,
	tok frontkit.Token) bool {
	pos := tokenPos(tok)
	switch tok.Text {
	case "export":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsExport, pos, tok.Text, "export_statement"))
	case "yield":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsYield, pos, tok.Text, "yield_expression"))
	case "using":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsUsingDecl, pos, tok.Text, "using_declaration"))
	case "let", "const":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsLexicalDecl, pos, tok.Text, "lexical_declaration"))
	case "this", "super":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsThisSuper, pos, tok.Text, "this-super"))
	case "new":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsNew, pos, tok.Text, "new_expression"))
	case "debugger":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsDebugger, pos, tok.Text, "debugger_statement"))
	case "class":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsClassBody, pos, tok.Text, "class_body"))
		if jsIsClassExprContext(p.toks, i) {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsClassExpr, pos, tok.Text, "class_expression"))
		}
	case "switch":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsSwitchBody, pos, tok.Text, "switch_body"))
	case "extends":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsClassHeritage, pos, tok.Text, "class_heritage"))
	case "function":
		if next := p.significant(i+1, 1); next != nil && next.Text == "*" {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsGeneratorFnDecl, pos, tok.Text, "generator_function_declaration"))
			*out = append(*out, p.mapper.dedicated(
				air.KindJsGeneratorFn, pos, tok.Text, "generator_function"))
		}
	case "default":
		if next := p.significant(i+1, 1); next != nil && next.Text == ":" {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsSwitchDefault, pos, tok.Text, "switch_default"))
		}
	case "static":
		if next := p.significant(i+1, 1); next != nil && next.Text == "{" {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsStaticBlock, pos, tok.Text, "class_static_block"))
		}
	default:
		return false
	}
	return true
}

// symbolMarker emits the dedicated marker for one operator, punctuation or
// literal token, using the surrounding context to anchor a real construct.
func (p *jsParser) symbolMarker(out *[]air.Node, i int,
	tok frontkit.Token) {
	pos := tokenPos(tok)
	switch {
	case jsIsAugAssign(tok.Text):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsAugAssign, pos, tok.Text, "augmented_assignment_expression"))
	case tok.Text == "." || tok.Text == "?.":
		p.dotMarker(out, i, tok)
	case tok.Kind == frontkit.TokenStr && jsIsTemplate(tok):
		if strings.Contains(tok.Text, "${") {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsTemplateSub, pos, "${", "template_substitution"))
		}
	case tok.Kind == frontkit.TokenStr:
		*out = append(*out, p.mapper.dedicated(
			air.KindJsStringPart, pos, tok.Text, "string_fragment"))
	case tok.Text == "/" && jsRegexAllowed(p.significant(i-1, -1)):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsRegex, pos, tok.Text, "regex"))
	case tok.Text == "++" || tok.Text == "--":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsIncDec, pos, tok.Text, "update_expression"))
	case tok.Text == ",":
		*out = append(*out, p.mapper.dedicated(
			air.KindJsSequence, pos, tok.Text, "sequence_expression"))
	case tok.Text == "(" && jsIsFormalParamsOpen(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsFormalParams, pos, tok.Text, "formal_parameters"))
	case tok.Kind == frontkit.TokenIdent && jsIsDeclaratorIdent(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsDeclarator, pos, tok.Text, "variable_declarator"))
	}
	p.constructMarker(out, i, tok)
}

// dotMarker emits the meta-property marker (new.target / import.meta) and the
// property-identifier marker for one dot-family token.
func (p *jsParser) dotMarker(out *[]air.Node, i int, tok frontkit.Token) {
	if prev := p.significant(i-1, -1); prev != nil &&
		(prev.Text == "new" || prev.Text == "import") {
		if next := p.significant(i+1, 1); next != nil &&
			(next.Text == "target" || next.Text == "meta") {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsMetaProperty, tokenPos(tok), tok.Text, "meta_property"))
		}
	}
	if next := p.significant(i+1, 1); next != nil &&
		jsIsLexicalToken(next.Kind) && next.Text != "" {
		*out = append(*out, p.mapper.dedicated(
			air.KindJsPropertyIdent, tokenPos(*next), next.Text,
			"property_identifier"))
	}
}

// constructMarker emits the dedicated markers for the bracket, colon,
// assignment and spread constructs whose recognition needs the token-interval
// context predicates of jsdedicated_context.go. Its token texts never overlap
// symbolMarker's cases, so the two dispatchers stay independent.
func (p *jsParser) constructMarker(out *[]air.Node, i int,
	tok frontkit.Token) {
	pos := tokenPos(tok)
	switch {
	case tok.Text == "[" && jsIsComputedName(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsComputedName, pos, tok.Text, "computed_property_name"))
	case tok.Text == ":" && jsIsPairColon(p.toks, i):
		if jsInPatternBraces(p.toks, i) {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsPairPattern, pos, tok.Text, "pair_pattern"))
		} else {
			*out = append(*out, p.mapper.dedicated(
				air.KindJsPair, pos, tok.Text, "pair"))
		}
	case tok.Text == "=" && jsIsObjectAssignDefault(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsObjectAssignPat, pos, tok.Text, "object_assignment_pattern"))
	case tok.Text == "..." && jsIsRestInParams(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsRestPattern, pos, tok.Text, "rest_pattern"))
	case tok.Text == "{" && jsIsPatternBrace(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsObjectPattern, pos, tok.Text, "object_pattern"))
	case tok.Text == "[" && !jsIsComputedName(p.toks, i) &&
		jsIsPatternBracket(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsArrayPattern, pos, tok.Text, "array_pattern"))
	case tok.Text == "=" && !jsIsObjectAssignDefault(p.toks, i) &&
		jsIsAssignPattern(p.toks, i):
		*out = append(*out, p.mapper.dedicated(
			air.KindJsAssignPattern, pos, tok.Text, "assignment_pattern"))
	}
}
