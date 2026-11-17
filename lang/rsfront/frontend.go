/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

// This file holds the public air.Frontend wiring and the Rust LanguageSpec.

import (
	"fmt"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// RsFrontend is the Rust frontend handle, an alias of the shared frontkit base
// bound to rsSpec by NewRsFrontend.
type RsFrontend = frontkit.FrontendBase

// Compile-time proofs of the two contracts this package implements.
var (
	_ frontkit.LanguageSpec = (*rsSpec)(nil)
	_ air.Frontend          = (*RsFrontend)(nil)
)

// NewRsFrontend returns an air.Frontend that reduces Rust sources to AIR.
func NewRsFrontend() (air.Frontend, error) {
	return frontkit.NewFrontendBase(defaultRsSpec())
}

// rsSpec implements frontkit.LanguageSpec for Rust. The injected tables are
// copied per instance so a mutated spec never changes another frontend.
type rsSpec struct {
	cfg      rsConfig
	keywords []string
}

// defaultRsSpec builds the production Rust language specification.
func defaultRsSpec() *rsSpec {
	return &rsSpec{cfg: defaultRsConfig()}
}

// Tag returns the canonical short Rust language tag.
func (*rsSpec) Tag() string { return "rs" }

// LexSpec returns the Rust lexical tables.
func (s *rsSpec) LexSpec() frontkit.LexSpec {
	return rsLexSpec{keywords: s.keywords}
}

// NormalizeDialect maps one edition spelling onto a canonical 2015-2024 string.
func (*rsSpec) NormalizeDialect(d string) (string, error) {
	key := rsNormalizeDialectKey(d)
	if key == "" {
		return rsDefaultDialect, nil
	}
	for i := range rsDialectAliases {
		if rsDialectAliases[i].alias == key {
			return rsDialectAliases[i].canonical, nil
		}
	}
	return "", fmt.Errorf("rsfront: unsupported Rust dialect %q", d)
}

// ParseTokens reduces one file's token stream into a layered AIR tree.
func (s *rsSpec) ParseTokens(tokens []frontkit.Token, path string) (air.Node,
	[]air.Diagnostic, error) {
	return newRsParser(tokens, path, s.cfg).parse()
}

// The rsMapper lowering methods and the ordered-table lookups below live
// here so every source file stays within the 300-line limit; they belong
// to the same mapping layer as rsmap.go.

// binaryNodes lowers already-reduced operands to KindBinary.
func (m *rsMapper) binaryNodes(left air.Node, op frontkit.Token,
	right air.Node) air.Node {
	return m.boundOperator(air.KindBinary, spanToken(left, op),
		spanToken(right, op), op, left, right)
}

// TryNode lowers the postfix question-mark operator to KindUnary with a try-op note.
func (m *rsMapper) tryNode(left air.Node, op frontkit.Token) air.Node {
	node := m.boundOperator(air.KindUnary, spanToken(left, op), op, op, left)
	return m.annotate(node, air.SemNote, m.note("try-op"))
}

// RangeNode lowers open and closed range forms to KindBinary with a range note.
func (m *rsMapper) rangeNode(left air.Node, op frontkit.Token,
	right air.Node) air.Node {
	start := op
	if left != nil {
		start = spanToken(left, op)
	}
	end := op
	children := make([]air.Node, 0, 2)
	if left != nil {
		children = append(children, left)
	}
	if right != nil {
		children = append(children, right)
		end = spanToken(right, op)
	}
	node := m.boundOperator(air.KindBinary, start, end, op, children...)
	return m.annotate(node, air.SemNote, m.note("range"))
}

// cast lowers an as-conversion to KindCast.
func (m *rsMapper) cast(left air.Node, op frontkit.Token,
	right air.Node) air.Node {
	return m.boundOperator(air.KindCast, spanToken(left, op),
		spanToken(right, op), op, left, right)
}

// arrow lowers a match-arm fat arrow to KindArrow, binding the operator token
// text into the node payload. This is what lets a mutated arrow (say "=>" ->
// "=!=") change the AIR node identity, instead of both spellings being dropped
// as an unbound operator and silently yielding the same tree.
func (m *rsMapper) arrow(left air.Node, op frontkit.Token,
	right air.Node) air.Node {
	children := make([]air.Node, 0, 2)
	if left != nil {
		children = append(children, left)
	}
	if right != nil {
		children = append(children, right)
	}
	start := op
	if left != nil {
		start = spanToken(left, op)
	}
	end := op
	if right != nil {
		end = spanToken(right, op)
	}
	return m.boundOperator(air.KindArrow, start, end, op, children...)
}

// refNode lowers &expr to KindRef.
func (m *rsMapper) refNode(op frontkit.Token, operand air.Node) air.Node {
	node := m.boundOperator(air.KindRef, op, spanToken(operand, op), op, operand)
	return m.annotate(node, air.SemNote, m.note("ref"))
}

// mutRefNode lowers &mut expr to KindRef with a mutability marker.
func (m *rsMapper) mutRefNode(op frontkit.Token, operand air.Node) air.Node {
	node := m.boundOperator(air.KindRef, op, spanToken(operand, op), op, operand)
	node = m.annotate(node, air.SemMutability, "mutable")
	return m.annotate(node, air.SemNote, m.note("ref"))
}

// addrNode lowers &&expr to KindAddr.
func (m *rsMapper) addrNode(op frontkit.Token, operand air.Node) air.Node {
	return m.unary(op, operand, air.KindAddr, "ref")
}

// starNode lowers a dereference to KindStar.
func (m *rsMapper) starNode(op frontkit.Token, operand air.Node) air.Node {
	return m.unary(op, operand, air.KindStar, "")
}

// lifetime lowers 'a to KindGenericParam with a lifetime note.
func (m *rsMapper) lifetime(start, name frontkit.Token) air.Node {
	node := m.build(air.KindGenericParam, start, name, name.Text,
		m.leaf(name, air.KindIdent))
	return m.annotate(node, air.SemNote, m.note("lifetime"))
}

// call lowers a plain call expression to KindCall.
func (m *rsMapper) call(left air.Node, open, end frontkit.Token,
	args ...air.Node) air.Node {
	children := append([]air.Node{left}, args...)
	return m.build(air.KindCall, spanToken(left, open), end, "", children...)
}

// index lowers a subscript to KindIndex.
func (m *rsMapper) index(left air.Node, open, end frontkit.Token,
	parts ...air.Node) air.Node {
	children := append([]air.Node{left}, parts...)
	return m.build(air.KindIndex, spanToken(left, open), end, "", children...)
}

// field lowers an attribute access to KindField.
func (m *rsMapper) field(left air.Node, name frontkit.Token) air.Node {
	return m.build(air.KindField, spanToken(left, name), name, "",
		left, m.lexical(name))
}

// selector lowers a::b to KindSelector.
func (m *rsMapper) selector(left air.Node, op, name frontkit.Token) air.Node {
	return m.boundOperator(air.KindSelector, spanToken(left, op), name, op,
		left, m.lexical(name))
}

// awaitNode lowers .await to KindAwait with SemAsync.
func (m *rsMapper) awaitNode(left air.Node, name frontkit.Token) air.Node {
	node := m.build(air.KindAwait, spanToken(left, name), name, "", left)
	return m.annotate(node, air.SemAsync, "true")
}

// tuple lowers a parenthesized tuple to KindTuple, or the sole inner node.
func (m *rsMapper) tuple(open, end frontkit.Token,
	items ...air.Node) air.Node {
	if len(items) == 1 {
		return items[0]
	}
	return m.build(air.KindTuple, open, end, "", items...)
}

// array lowers an array literal to KindCompositeLit with an array note.
func (m *rsMapper) array(open, end frontkit.Token,
	items ...air.Node) air.Node {
	node := m.build(air.KindCompositeLit, open, end, "", items...)
	return m.annotate(node, air.SemNote, m.note("array"))
}

// block lowers a bare braced expression to KindBlock.
func (m *rsMapper) block(open, end frontkit.Token,
	children ...air.Node) air.Node {
	return m.build(air.KindBlock, open, end, "", children...)
}

// lambda lowers a closure to KindLambda.
func (m *rsMapper) lambda(start, end frontkit.Token, params []air.Node,
	body air.Node) air.Node {
	children := append([]air.Node{}, params...)
	finish := end
	if body != nil {
		children = append(children, body)
	} else if len(params) > 0 {
		finish = spanToken(params[len(params)-1], start)
	}
	return m.build(air.KindLambda, start, finish, "", children...)
}
