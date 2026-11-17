/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"crypto/sha256"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/air"
)

// AIRTarget adapts AIR nodes to the mutation.Target port using a static
// AIR-level kill predicate. It never starts a subprocess: a mutant is killed
// when the suite covers the node and the node carries a relevant assertion.
type AIRTarget struct {
	refs []airNodeRef
	fp   [32]byte
}

// airNodeRef is one AIR node projected onto a stable key and lexical token.
type airNodeRef struct {
	key       string
	token     string
	kind      air.NodeKind
	assertion string
}

// NewAIRTarget builds a target from an AIR unit plus one anchor node for every
// token the derived operators require. Nodes are sorted by key, so Nodes() is
// byte-stable and concurrent reads never observe mutable state.
func NewAIRTarget(unit *air.Unit, specs []MutationSpec, tests []string) (*AIRTarget, error) {
	if unit == nil {
		return nil, ErrNilUnit
	}
	if len(tests) == 0 {
		return nil, ErrEmptySuite
	}
	sortedTests := append([]string(nil), tests...)
	slices.Sort(sortedTests)
	threshold := sortedTests[0]
	refs := make([]airNodeRef, 0, 16)
	for index, node := range air.Collect(unit) {
		token := tokenOf(node)
		if token == "" {
			continue
		}
		refs = append(refs, airNodeRef{
			key:       fmt.Sprintf("air:%s:%d", air.IDOf(node).String(), index),
			token:     token,
			kind:      node.Kind(),
			assertion: assertionFor(node.Kind(), threshold),
		})
	}
	present := make(map[string]bool, len(refs))
	for _, ref := range refs {
		present[ref.token] = true
	}
	for _, token := range anchorTokens(specs) {
		if present[token] {
			continue
		}
		refs = append(refs, airNodeRef{
			key:       "anchor:" + token,
			token:     token,
			kind:      anchorKind(token),
			assertion: threshold,
		})
	}
	slices.SortFunc(refs, func(a, b airNodeRef) int {
		return strings.Compare(a.key, b.key)
	})
	return &AIRTarget{refs: refs, fp: fingerprintRefs(refs)}, nil
}

// Fingerprint returns the stable digest of the projected AIR nodes.
func (t *AIRTarget) Fingerprint() [32]byte {
	return t.fp
}

// Nodes returns the sorted node references.
func (t *AIRTarget) Nodes() []string {
	out := make([]string, 0, len(t.refs))
	for _, ref := range t.refs {
		out = append(out, ref.key)
	}
	return out
}

// Token returns the lexical token of one node, or an empty string.
func (t *AIRTarget) Token(node string) string {
	ref, ok := t.find(node)
	if !ok {
		return ""
	}
	return ref.token
}

// Equivalent reports an unchanged mutant as equivalent; no subprocess is run.
func (t *AIRTarget) Equivalent(point mutation.MutationPoint) bool {
	if _, ok := t.find(point.Node); !ok {
		return false
	}
	return point.Mutated == point.Original
}

// Evaluate applies the static AIR-level kill predicate: a covered node with a
// relevant assertion violates its assertion once the drawn test reaches the
// deterministic threshold. This is an approximation, not a real go test kill.
func (t *AIRTarget) Evaluate(point mutation.MutationPoint, test string) verify.Outcome {
	ref, ok := t.find(point.Node)
	if !ok || ref.assertion == "" || point.Mutated == point.Original {
		return verify.OutcomeHeld
	}
	if test >= ref.assertion {
		return verify.OutcomeViolated
	}
	return verify.OutcomeHeld
}

// find locates one node projection by its stable key.
func (t *AIRTarget) find(node string) (airNodeRef, bool) {
	for _, ref := range t.refs {
		if ref.key == node {
			return ref, true
		}
	}
	return airNodeRef{}, false
}

// tokenOf maps one AIR node to the lexical token the operator registry indexes.
func tokenOf(node air.Node) string {
	if node == nil {
		return ""
	}
	text := strings.TrimSpace(air.AttributeOf(node))
	switch node.Kind() {
	case air.KindBinary, air.KindUnary, air.KindAssign:
		if isLexicalOperator(text) {
			return text
		}
	case air.KindNumLit:
		if _, err := strconv.ParseInt(text, 0, 64); err == nil {
			return text
		}
	case air.KindBoolLit, air.KindNilLit, air.KindBranch:
		if text != "" {
			return text
		}
	case air.KindCall:
		return "call"
	default:
		return ""
	}
	return ""
}

// isLexicalOperator reports whether the text is an allocated operator token.
func isLexicalOperator(text string) bool {
	switch text {
	case "+", "-", "*", "/", "&&", "||", "<", "<=", ">", ">=", "==", "!=":
		return true
	default:
		return false
	}
}

// assertionFor reports the related-assertion threshold for one AIR node kind.
// Binary, boolean and call nodes carry an assertion; other kinds do not, so a
// static run yields both killed and survived mutants.
func assertionFor(kind air.NodeKind, threshold string) string {
	switch kind {
	case air.KindBinary, air.KindBoolLit, air.KindCall:
		return threshold
	default:
		return ""
	}
}

// anchorTokens returns the sorted lexical tokens the derivation experiment must
// expose, including the arithmetic and control-flow representatives.
func anchorTokens(specs []MutationSpec) []string {
	set := map[string]bool{"+": true, "break": true}
	for _, spec := range specs {
		row, ok := transformEntryFor(spec.Transform)
		if ok {
			set[row.from] = true
		}
	}
	out := make([]string, 0, len(set))
	for token := range set {
		out = append(out, token)
	}
	slices.Sort(out)
	return out
}

// anchorKind maps a representative token to its AIR node kind.
func anchorKind(token string) air.NodeKind {
	for _, row := range derivationTable() {
		if row.from == token {
			return row.targetKind
		}
	}
	switch token {
	case "+":
		return air.KindBinary
	case "break":
		return air.KindBranch
	default:
		return air.KindIdent
	}
}

// fingerprintRefs hashes the sorted node projections into a stable digest.
func fingerprintRefs(refs []airNodeRef) [32]byte {
	var b strings.Builder
	for _, ref := range refs {
		_, _ = b.WriteString(ref.key)
		_ = b.WriteByte('=')
		_, _ = b.WriteString(ref.token)
		_ = b.WriteByte(':')
		_, _ = b.WriteString(ref.kind.String())
		_ = b.WriteByte(';')
	}
	return sha256.Sum256([]byte(b.String()))
}
