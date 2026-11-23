/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gotarget

import (
	"fmt"
	"strings"

	"abate/sim/rand"
)

// functionBody carries the rendered source of one generated function and
// the rendered source of its guarding test function.
type functionBody struct {
	source string
	test   string
}

// buildFunction renders one target function of the given structure family
// plus its test. All varying parts (operand constants, bounds, whether the
// edge case is asserted) draw from the stream in a fixed order.
func buildFunction(stream *rand.Stream, name, structure, level string) functionBody {
	switch structure {
	case "arith":
		return buildArith(stream, name, level)
	case "boundary":
		return buildBoundary(stream, name, level)
	case "constant":
		return buildConstant(stream, name, level)
	case "controlflow":
		return buildControlFlow(stream, name, level)
	case "returnvalue":
		return buildReturnValue(stream, name, level)
	case "negate":
		return buildNegate(stream, name, level)
	case "callremove":
		return buildCallRemove(stream, name, level)
	}
	// Unreachable: structureOrder is a closed set.
	panic("gotarget: unknown structure " + structure)
}

// pinned reports whether the current structure's edge case is asserted by
// its test, drawn from the stream with a level-dependent probability:
// easy asserts almost every edge, medium about half, hard almost none.
// A pinned edge kills the mutants that touch it; an unpinned edge lets them
// survive, which is how the levels shape the resulting mutation score
// without ever touching the judging itself.
func pinned(stream *rand.Stream, level string) bool {
	draw := stream.IntN(100)
	switch level {
	case LevelEasy:
		return draw < 90
	case LevelHard:
		return draw < 25
	default: // LevelMedium
		return draw < 60
	}
}

// intOperand draws a small positive operand constant.
func intOperand(stream *rand.Stream) int {
	return 2 + stream.IntN(47) // 2..48
}

// buildArith renders an int arithmetic function with one mutable operator
// between two operands, plus a test asserting the exact result on drawn
// inputs. A mutant flipping the operator is killed by that assertion.
func buildArith(stream *rand.Stream, name, level string) functionBody {
	a := intOperand(stream)
	b := intOperand(stream)
	x := intOperand(stream)
	y := intOperand(stream)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s scales x by %[2]d and adds y scaled by %[3]d.
func %[1]s(x, y int) int {
	return x*%[2]d + y*%[3]d
}
`, name, a, b)
	want := x*a + y*b
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if got := %[1]s(0, 1); got != %[2]d {
		t.Errorf("%[1]s(0,1) = %%d, want %[2]d", got)
	}
`, name, b)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(%[2]d, %[3]d); got != %[4]d {
		t.Errorf("%[1]s(%[2]d,%[3]d) = %%d, want %[4]d", got)
	}
%[5]s}
`, name, x, y, want, edgeCase)
	return functionBody{source: src, test: test}
}

// buildBoundary renders a comparison-based clamp with one boundary token;
// the test asserts interior behavior always and the boundary point only
// when pinned, so a boundary mutant survives exactly on unpinned targets.
func buildBoundary(stream *rand.Stream, name, level string) functionBody {
	limit := intOperand(stream)
	below := 1 + stream.IntN(limit-1)   // always below the limit
	above := limit + intOperand(stream) // always above the limit
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s reports whether n stays below %[2]d.
func %[1]s(n int) bool {
	return n < %[2]d
}
`, name, limit)
	var edgeCase string
	if edge {
		if level == LevelEasy {
			// The interior side of the boundary is asserted too, so a
			// boundary mutant shifting the comparison by one is exposed
			// from both sides, not just the boundary point itself.
			edgeCase = fmt.Sprintf(`
	if !%[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = false, want true just below the boundary")
	}
	if %[1]s(%[3]d) {
		t.Errorf("%[1]s(%[3]d) = true, want false at the boundary")
	}
`, name, limit-1, limit)
		} else {
			edgeCase = fmt.Sprintf(`
	if %[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = true, want false at the boundary")
	}
`, name, limit)
		}
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if !%[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = false, want true")
	}
	if %[1]s(%[3]d) {
		t.Errorf("%[1]s(%[3]d) = true, want false")
	}
%[4]s}
`, name, below, above, edgeCase)
	return functionBody{source: src, test: test}
}

// buildConstant renders a function returning a constant-derived value; the
// test pins the exact constant so a perturbation mutant is always killed.
func buildConstant(stream *rand.Stream, name, level string) functionBody {
	base := intOperand(stream)
	factor := intOperand(stream)
	offset := intOperand(stream)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s returns the seed-derived constant scaled by factor %[2]d.
func %[1]s() int {
	base := %[3]d
	return base*%[2]d + %[4]d
}
`, name, factor, base, offset)
	want := base*factor + offset
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if %[1]s() == 0 {
		t.Errorf("%[1]s() = 0, want a positive constant")
	}
`, name)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(); got != %[2]d {
		t.Errorf("%[1]s() = %%d, want %[2]d", got)
	}
%[3]s}
`, name, want, edgeCase)
	return functionBody{source: src, test: test}
}

// buildControlFlow renders a loop with break/continue; the test asserts the
// aggregate so a control-flow mutant changing the loop exit is killed.
func buildControlFlow(stream *rand.Stream, name, level string) functionBody {
	stop := 3 + stream.IntN(9) // 3..11
	skip := 1 + stream.IntN(stop-1)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s sums the values below %[2]d that are not skipped.
func %[1]s() int {
	total := 0
	for i := 0; i < %[2]d; i++ {
		if i == %[3]d {
			continue
		}
		if total > %[2]d*%[3]d {
			break
		}
		total += i
	}
	return total
}
`, name, stop, skip)
	want := 0
	for i := 0; i < stop; i++ {
		if i == skip {
			continue
		}
		if want > stop*skip {
			break
		}
		want += i
	}
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if %[1]s() < 0 {
		t.Errorf("%[1]s() negative, want a non-negative sum")
	}
`, name)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(); got != %[2]d {
		t.Errorf("%[1]s() = %%d, want %[2]d", got)
	}
%[3]s}
`, name, want, edgeCase)
	return functionBody{source: src, test: test}
}

// buildReturnValue renders a boolean classifier returning true/false on
// disjoint branches; the test asserts both branches when pinned.
func buildReturnValue(stream *rand.Stream, name, level string) functionBody {
	threshold := intOperand(stream)
	low := intOperand(stream) % threshold
	high := threshold + intOperand(stream)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s reports whether n qualifies against %[2]d.
func %[1]s(n int) bool {
	if n >= %[2]d {
		return true
	}
	return false
}
`, name, threshold)
	var edgeCase string
	if edge {
		if level == LevelEasy {
			edgeCase = fmt.Sprintf(`
	if %[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = true, want false below the threshold")
	}
	if !%[1]s(%[3]d) {
		t.Errorf("%[1]s(%[3]d) = false, want true just at the threshold")
	}
`, name, low, threshold)
		} else {
			edgeCase = fmt.Sprintf(`
	if %[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = true, want false below the threshold")
	}
`, name, low)
		}
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if !%[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = false, want true")
	}
%[3]s}
`, name, high, edgeCase)
	return functionBody{source: src, test: test}
}

// buildNegate renders a predicate with an equality token; a negation mutant
// flips the outcome, which the drawn assertions catch when pinned.
func buildNegate(stream *rand.Stream, name, level string) functionBody {
	want := intOperand(stream)
	hit := want
	miss := want + 1 + stream.IntN(5)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s reports whether n equals %[2]d.
func %[1]s(n int) bool {
	return n == %[2]d
}
`, name, want)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if %[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = true, want false")
	}
`, name, miss)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if !%[1]s(%[2]d) {
		t.Errorf("%[1]s(%[2]d) = false, want true")
	}
%[3]s}
`, name, hit, edgeCase)
	return functionBody{source: src, test: test}
}

// buildCallRemove renders a function whose result depends on a helper call;
// the test asserts the helper's contribution so a call-removal mutant is
// killed, and pins the zero-input edge only when pinned.
func buildCallRemove(stream *rand.Stream, name, level string) functionBody {
	base := intOperand(stream)
	input := intOperand(stream)
	edge := pinned(stream, level)
	helper := name + "Add"
	src := fmt.Sprintf(`
// %[1]s adds its argument to the base %[2]d through %[3]s.
func %[1]s(n int) int {
	return %[3]s(n)
}

// %[3]s returns n shifted by %[2]d.
func %[3]s(n int) int {
	return n + %[2]d
}
`, name, base, helper)
	want := input + base
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if got := %[1]s(0); got != %[2]d {
		t.Errorf("%[1]s(0) = %%d, want %[2]d", got)
	}
`, name, base)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(%[2]d); got != %[3]d {
		t.Errorf("%[1]s(%[2]d) = %%d, want %[3]d", got)
	}
%[4]s}
`, name, input, want, edgeCase)
	return functionBody{source: src, test: test}
}

// sourceBuilder accumulates the functions of target.go.
type sourceBuilder struct {
	pkg   string
	parts []string
}

func (b *sourceBuilder) addFunction(src string) { b.parts = append(b.parts, src) }

// render produces the final target.go with a deterministic header.
func (b *sourceBuilder) render() string {
	var sb strings.Builder
	sb.WriteString("// Code generated by abate targetgen. DO NOT EDIT.\n")
	sb.WriteString("// The module is a deterministic mutation-testing target.\n\n")
	fmt.Fprintf(&sb, "package %s\n", b.pkg)
	for _, part := range b.parts {
		sb.WriteString(part)
	}
	return sb.String()
}

// testBuilder accumulates the test functions of target_test.go.
type testBuilder struct {
	pkg   string
	parts []string
}

func (b *testBuilder) addFunction(src string) { b.parts = append(b.parts, src) }

// render produces the final target_test.go with a deterministic header.
func (b *testBuilder) render() string {
	var sb strings.Builder
	sb.WriteString("// Code generated by abate targetgen. DO NOT EDIT.\n\n")
	fmt.Fprintf(&sb, "package %s\n\n", b.pkg)
	sb.WriteString("import \"testing\"\n")
	for _, part := range b.parts {
		sb.WriteString(part)
	}
	return sb.String()
}

// renderGoMod produces the deterministic go.mod of the target module. The
// language version is pinned to the generator's own floor so the module
// builds on any toolchain that can build abate itself.
func renderGoMod(module string) string {
	return fmt.Sprintf("module %s\n\ngo 1.21\n", module)
}

// countMutableTokens approximates the mutable token count of the generated
// source for reporting: it counts the token classes the mutation scanner
// looks for, using plain substring matching over the rendered source. This
// is a report convenience, never a judging input.
func countMutableTokens(source []byte) int {
	text := string(source)
	count := 0
	for _, tok := range []string{" + ", " - ", " * ", " / ", " < ", " <= ",
		" > ", " >= ", " == ", " != ", "break", "continue", "true", "false"} {
		count += strings.Count(text, tok)
	}
	return count
}
