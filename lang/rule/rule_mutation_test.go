/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"errors"
	"testing"

	"abate/lang/air"
)

// mutationLeaf builds a tiny identifier node for boundary matching.
func mutationLeaf(text string) air.Node {
	return air.NewNode(air.KindIdent,
		air.NewPosition(1, 1, 0, len(text)), text)
}

// TestMutationBadRules covers every rejection path of Compile and Registry.
func TestMutationBadRules(t *testing.T) {
	t.Parallel()
	badSeverity := validSpec("GO-EG-001")
	badSeverity.Severity = "bogus"
	if _, err := Compile(badSeverity); !errors.Is(err, ErrInvalidSpec) {
		t.Fatalf("Compile with a bad severity = %v", err)
	}
	badPattern := validSpec("GO-EG-001")
	badPattern.Pattern = "("
	if _, err := Compile(badPattern); !errors.Is(err, ErrBadPattern) {
		t.Fatalf("Compile with a bad pattern = %v", err)
	}
	badWhere := validSpec("GO-EG-001")
	badWhere.Where = []string{"bogus:x"}
	if _, err := Compile(badWhere); !errors.Is(err, ErrInvalidPredicate) {
		t.Fatalf("Compile with a bad where clause = %v", err)
	}
	zero := &Registry{}
	if err := zero.Register(validSpec("GO-EG-001")); err != nil || zero.Len() != 1 {
		t.Fatalf("zero-value Register = %v, size %d", err, zero.Len())
	}
	zeroBatch := &Registry{}
	if err := zeroBatch.Load([]Spec{validSpec("GO-EG-002")}); err != nil ||
		zeroBatch.Len() != 1 {
		t.Fatalf("zero-value Load = %v, size %d", err, zeroBatch.Len())
	}
	if RuleID("GO-EG-abc").Sequence() != 0 {
		t.Fatal("unparseable sequence must return zero")
	}
}

// TestMutationPatternBoundaries covers malformed trees, negated roots and
// predicates that the parser cannot produce directly.
func TestMutationPatternBoundaries(t *testing.T) {
	t.Parallel()
	for _, text := range []string{"!", "!Ident"} {
		if _, err := CompilePattern(text); !errors.Is(err, ErrBadPattern) {
			t.Errorf("CompilePattern(%q) = %v, want ErrBadPattern", text, err)
		}
	}
	cases := []*patternNode{
		{ellipsis: true, children: []*patternNode{{}}},
		{repeat: true, children: []*patternNode{{}}},
		{neg: true},
		{neg: true, inner: &patternNode{ellipsis: true,
			children: []*patternNode{{}}}},
	}
	for i, node := range cases {
		if err := validatePatternTree(node, 0); err == nil {
			t.Errorf("boundary tree %d was accepted", i)
		}
	}
	if err := validatePatternTree(&patternNode{}, maxPatternDepth+1); err == nil {
		t.Fatal("over-deep tree was accepted")
	}
	if (Predicate{Op: "bogus"}).textMatches("x") {
		t.Fatal("unknown text operator must never match")
	}
	_, _, call := testTree()
	if hasSemantic(call, "NotAKey") {
		t.Fatal("unknown annotation name must not be found")
	}
	if got := expandMessage("$short $longer", []Capture{
		{Name: "short", Text: "A"}, {Name: "longer", Text: "B"},
	}); got != "A B" {
		t.Fatalf("expandMessage = %q, want %q", got, "A B")
	}
}

// TestMutationMatcherEdges covers matchAt and matchRepeat boundary branches.
func TestMutationMatcherEdges(t *testing.T) {
	t.Parallel()
	_, _, call := testTree()
	empty := Matcher{}
	if _, ok := empty.matchAt(nil, &patternNode{}, nil); ok {
		t.Fatal("nil node must not match")
	}
	if _, ok := empty.matchAt(call, nil, nil); ok {
		t.Fatal("nil pattern vertex must not match")
	}
	if _, ok := empty.matchAt(call, &patternNode{ellipsis: true}, nil); ok {
		t.Fatal("ellipsis cannot be a root vertex")
	}
	if _, ok := empty.matchAt(call, &patternNode{repeat: true}, nil); ok {
		t.Fatal("repeat cannot be a root vertex")
	}
	if _, ok := empty.matchAt(call,
		&patternNode{neg: true, inner: &patternNode{kind: "Call"}}, nil); ok {
		t.Fatal("a satisfied negation must reject")
	}
	if _, ok := empty.matchAt(call,
		&patternNode{neg: true, inner: &patternNode{kind: "Return"}}, nil); !ok {
		t.Fatal("an unsatisfied negation must accept")
	}
	if _, ok := empty.matchAt(call, &patternNode{attrOp: "bogus"}, nil); ok {
		t.Fatal("an unknown attribute operator must reject")
	}
	caps, ok := empty.matchAt(call, &patternNode{capture: "x"}, nil)
	if !ok || len(caps) != 1 || caps[0].Name != "x" {
		t.Fatalf("capture match = %v, %v", caps, ok)
	}
	literal := air.NewNode(air.KindStrLit,
		air.NewPosition(1, 1, 0, 3), `"s"`)
	candidates := []air.Node{mutationLeaf("a"), literal, mutationLeaf("c")}
	constrained := &patternNode{repeat: true, capture: "x", kind: "Ident"}
	if _, ok := matchRepeat(empty, candidates, 0,
		[]*patternNode{constrained, {kind: "NumLit"}}, 0, nil); ok {
		t.Fatal("repeat with a constrained atom must eventually fail")
	}
	for _, text := range []string{
		"(Call $x*)", "(Call $x* Ident=c)", "(Call $x* Ident=zzz)",
		"(Call Ident=a Ident=b)", "(Call $.. !ANY)",
	} {
		pattern, err := CompilePattern(text)
		if err != nil {
			t.Fatalf("compile %q: %v", text, err)
		}
		tree := air.NewNode(air.KindCall, air.NewPosition(1, 1, 0, 30), "",
			mutationLeaf("a"), mutationLeaf("b"), mutationLeaf("c"))
		_ = NewMatcher(pattern).Match(tree)
	}
}

// TestMutationCheckEdges covers nil roots and roots without file metadata.
func TestMutationCheckEdges(t *testing.T) {
	t.Parallel()
	compiled, err := Compile(validSpec("GO-EG-001"))
	if err != nil {
		t.Fatalf("compile: %v", err)
	}
	unit, root, _ := testTree()
	unit.Roots = append(unit.Roots, root)
	findings := compiled.Check(unit)
	if len(findings) == 0 {
		t.Fatal("an extra root without file metadata must still be checked")
	}
	if findings[0].Path != "" || findings[len(findings)-1].Path != "p.go" {
		t.Fatalf("root path mapping is wrong: %v", findings)
	}
	unit.Roots[0] = nil
	if got := compiled.Check(unit); len(got) != 1 {
		t.Fatalf("a nil root must leave exactly one finding, got %d", len(got))
	}
}
