/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package targetgen

import (
	"fmt"
	"strings"

	"abate/sim/rand"
)

// goFunctionBody carries the rendered source of one generated Go function
// (plus optional helpers) and the rendered source of its guarding test.
type goFunctionBody struct {
	source string
	test   string
	pinned bool
}

// generateGo renders the whole Go module: one source file with every
// function, one test file with every guarding test, plus go.mod. A share of
// functions (Spec.Mixed) is rendered as mixed-structure combos.
func generateGo(spec Spec, stream *rand.Stream, seed uint64) Target {
	src := &goSourceBuilder{pkg: spec.Package}
	tst := &goTestBuilder{pkg: spec.Package}
	funcs := make([]FunctionInfo, 0, spec.Funcs)
	pinnedCount := 0
	for i := 0; i < spec.Funcs; i++ {
		name := fmt.Sprintf("Target%02d", i)
		structure := pickStructure(spec, i)
		var body goFunctionBody
		if spec.Mixed > 0 && stream.Float64() < spec.Mixed {
			body = buildGoMixed(stream, name, structure, pickStructure(spec, i+3), spec.Level)
			structure = structure + "+" + pickStructure(spec, i+3)
		} else {
			body = buildGoFunction(stream, name, structure, spec.Level)
		}
		src.addFunction(body.source)
		tst.addFunction(body.test)
		if body.pinned {
			pinnedCount++
		}
		funcs = append(funcs, FunctionInfo{
			Name:      name,
			Structure: structure,
			TestName:  "Test" + name,
		})
	}
	source := src.render()
	files := map[string][]byte{
		"target.go":      []byte(source),
		"target_test.go": []byte(tst.render()),
		"go.mod":         []byte(fmt.Sprintf("module %s\n\ngo 1.21\n", spec.Module)),
	}
	return Target{
		Spec:             spec,
		Seed:             seed,
		Files:            files,
		Funcs:            funcs,
		Tokens:           countGoTokens([]byte(source)),
		PinnedStructures: pinnedCount,
	}
}

// buildGoFunction dispatches one structure family to its builder.
func buildGoFunction(stream *rand.Stream, name, structure, level string) goFunctionBody {
	switch structure {
	case "arith":
		return goArith(stream, name, level)
	case "boundary":
		return goBoundary(stream, name, level)
	case "constant":
		return goConstant(stream, name, level)
	case "controlflow":
		return goControlFlow(stream, name, level)
	case "returnvalue":
		return goReturnValue(stream, name, level)
	case "negate":
		return goNegate(stream, name, level)
	case "callremove":
		return goCallRemove(stream, name, level)
	case "slice":
		return goSlice(stream, name, level)
	case "map":
		return goMap(stream, name, level)
	case "errorpath":
		return goErrorPath(stream, name, level)
	case "nested":
		return goNested(stream, name, level)
	case "method":
		return goMethod(stream, name, level)
	case "iface":
		return goIface(stream, name, level)
	}
	panic("targetgen: unknown go structure " + structure)
}

// pinned reports whether the current structure's edge case is asserted,
// drawn with a level-dependent probability: easy asserts almost every edge,
// medium about half, hard almost none.
func pinned(stream *rand.Stream, level string) bool {
	draw := stream.IntN(100)
	switch level {
	case LevelEasy:
		return draw < 90
	case LevelHard:
		return draw < 25
	default:
		return draw < 60
	}
}

// intOperand draws a small positive operand constant.
func intOperand(stream *rand.Stream) int {
	return 2 + stream.IntN(47)
}

// buildGoMixed renders one function whose body combines two structure
// families: a leading computation of family A feeds a boundary dispatch of
// family B, so one body carries the mutable surface of both. The test
// asserts each face separately, and the pinned draw governs both edges.
func buildGoMixed(stream *rand.Stream, name, primary, secondary, level string) goFunctionBody {
	a := intOperand(stream)
	b := intOperand(stream)
	limit := intOperand(stream)
	x := intOperand(stream)
	y := intOperand(stream)
	edge := pinned(stream, level)
	scale := x*a + y*b
	verdict := scale < limit
	lowIn, highIn := 1+stream.IntN(limit-1), limit+intOperand(stream)
	src := fmt.Sprintf(`
// %[1]s scales x by %[2]d and y by %[3]d, then compares the sum against %[4]d.
// The body mixes an arithmetic structure with a boundary structure, so a
// mutant in either face is judged by the same real test.
func %[1]s(x, y int) bool {
	total := x*%[2]d + y*%[3]d
	return total < %[4]d
}
`, name, a, b, limit)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if got := %[1]s(%[2]d, 0); got != %[3]t {
		t.Errorf("%[1]s(%[2]d,0) = %%v, want %[3]t", got)
	}
`, name, limit/a, (limit/a)*a < limit)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(%[2]d, %[3]d); got != %[4]t {
		t.Errorf("%[1]s(%[2]d,%[3]d) = %%v, want %[4]t", got)
	}
	if %[1]s(%[5]d, 0) == %[1]s(%[6]d, 0) {
		t.Log("boundary inputs agree")
	}
%[7]s}
`, name, x, y, verdict, lowIn, highIn, edgeCase)
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goArith renders an int arithmetic function with one mutable operator.
func goArith(stream *rand.Stream, name, level string) goFunctionBody {
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goBoundary renders a comparison-based boundary predicate.
func goBoundary(stream *rand.Stream, name, level string) goFunctionBody {
	limit := intOperand(stream)
	below := 1 + stream.IntN(limit-1)
	above := limit + intOperand(stream)
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goConstant renders a function returning a constant-derived value.
func goConstant(stream *rand.Stream, name, level string) goFunctionBody {
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goControlFlow renders a loop with break/continue.
func goControlFlow(stream *rand.Stream, name, level string) goFunctionBody {
	stop := 3 + stream.IntN(9)
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goReturnValue renders a boolean classifier on disjoint branches.
func goReturnValue(stream *rand.Stream, name, level string) goFunctionBody {
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goNegate renders an equality predicate.
func goNegate(stream *rand.Stream, name, level string) goFunctionBody {
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goCallRemove renders a function whose result depends on a helper call.
func goCallRemove(stream *rand.Stream, name, level string) goFunctionBody {
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goSlice renders a slice-processing function: a loop over a seed-built
// slice with a filter condition, an accumulation and a boundary. It carries
// the mutable surface of loop, comparison and arithmetic at once.
func goSlice(stream *rand.Stream, name, level string) goFunctionBody {
	size := 3 + stream.IntN(5)
	limit := intOperand(stream)
	factor := intOperand(stream)
	edge := pinned(stream, level)
	values := make([]int, size)
	for i := range values {
		values[i] = stream.IntN(limit * 3)
	}
	lit := make([]string, size)
	total := 0
	for i, v := range values {
		lit[i] = fmt.Sprintf("%d", v)
		if v < limit {
			total += v * factor
		}
	}
	src := fmt.Sprintf(`
// %[1]s accumulates the elements below %[2]d, each scaled by %[3]d.
func %[1]s(values []int) int {
	total := 0
	for _, v := range values {
		if v < %[2]d {
			total += v * %[3]d
		}
	}
	return total
}
`, name, limit, factor)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if got := %[1]s([]int{%[2]d}); got != %[3]d {
		t.Errorf("%[1]s(edge) = %%d, want %[3]d", got)
	}
`, name, limit, 0)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s([]int{%[2]s}); got != %[3]d {
		t.Errorf("%[1]s(input) = %%d, want %[3]d", got)
	}
%[4]s}
`, name, strings.Join(lit, ", "), total, edgeCase)
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goMap renders a map-lookup function with an existence branch.
func goMap(stream *rand.Stream, name, level string) goFunctionBody {
	base := intOperand(stream)
	bonus := intOperand(stream)
	present := intOperand(stream)
	absent := present + intOperand(stream)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s prices item n: catalog entries earn a bonus of %[2]d over %[3]d.
func %[1]s(n int) int {
	catalog := map[int]int{%[4]d: %[3]d + %[2]d}
	if price, ok := catalog[n]; ok {
		return price
	}
	return %[3]d
}
`, name, bonus, base, present)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if got := %[1]s(%[2]d); got != %[3]d {
		t.Errorf("%[1]s(%[2]d) = %%d, want the fallback %[3]d", got)
	}
`, name, absent, base)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(%[2]d); got != %[3]d {
		t.Errorf("%[1]s(%[2]d) = %%d, want %[3]d", got)
	}
%[4]s}
`, name, present, base+bonus, edgeCase)
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goErrorPath renders a function with an error return path: the mutable
// surface is the comparison deciding the error branch.
func goErrorPath(stream *rand.Stream, name, level string) goFunctionBody {
	limit := intOperand(stream)
	factor := intOperand(stream)
	good := intOperand(stream) % limit
	bad := limit + intOperand(stream)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]s doubles n when n stays below %[2]d and reports an error otherwise.
func %[1]s(n int) (int, error) {
	if n >= %[2]d {
		return 0, err%[1]sRange
	}
	return n * %[3]d, nil
}

// err%[1]sRange is the sentinel returned when n leaves the accepted range.
var err%[1]sRange = errors.New("%[4]s out of range")
`, name, limit, factor, strings.ToLower(name))
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if _, err := %[1]s(%[2]d); err == nil {
		t.Errorf("%[1]s(%[2]d) error = nil, want the range error")
	}
`, name, limit)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	got, err := %[1]s(%[2]d)
	if err != nil {
		t.Fatalf("%[1]s(%[2]d) error = %%v, want nil", err)
	}
	if got != %[3]d {
		t.Errorf("%[1]s(%[2]d) = %%d, want %[3]d", got)
	}
	if _, err := %[1]s(%[4]d); err == nil {
		t.Errorf("%[1]s(%[4]d) error = nil, want the range error")
	}
%[5]s}
`, name, good, good*factor, bad, edgeCase)
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goNested renders a doubly nested loop with a deep break/continue surface.
func goNested(stream *rand.Stream, name, level string) goFunctionBody {
	rows := 2 + stream.IntN(3)
	cols := 2 + stream.IntN(3)
	skipCol := stream.IntN(cols)
	edge := pinned(stream, level)
	want := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if c == skipCol {
				continue
			}
			want += r*cols + c
		}
	}
	src := fmt.Sprintf(`
// %[1]s sums the grid cells of %[2]dx%[3]d, skipping column %[4]d.
func %[1]s() int {
	total := 0
	for r := 0; r < %[2]d; r++ {
		for c := 0; c < %[3]d; c++ {
			if c == %[4]d {
				continue
			}
			total += r*%[3]d + c
		}
	}
	return total
}
`, name, rows, cols, skipCol)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if %[1]s() <= 0 {
		t.Errorf("%[1]s() = non-positive, want a positive grid sum")
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
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goMethod renders a type with a pointer-receiver method whose behavior the
// test asserts; the mutable surface is the method's arithmetic and branch.
func goMethod(stream *rand.Stream, name, level string) goFunctionBody {
	step := intOperand(stream)
	limit := intOperand(stream)
	start := stream.IntN(limit)
	moves := 1 + stream.IntN(4)
	edge := pinned(stream, level)
	want := start
	for i := 0; i < moves; i++ {
		if want < limit {
			want += step
		}
	}
	src := fmt.Sprintf(`
// %[1]sCounter accumulates steps of %[2]d but never crosses %[3]d.
type %[1]sCounter struct {
	value int
}

// advance adds one step while the counter stays below %[3]d.
func (c *%[1]sCounter) advance() {
	if c.value < %[3]d {
		c.value += %[2]d
	}
}

// valueOf reports the current counter value.
func (c *%[1]sCounter) valueOf() int {
	return c.value
}
`, name, step, limit)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	c2 := &%[1]sCounter{value: %[2]d}
	c2.advance()
	if c2.valueOf() != %[3]d {
		t.Errorf("advance from %[2]d = %%d, want %[3]d", c2.valueOf())
	}
`, name, limit, limit)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	c := &%[1]sCounter{value: %[2]d}
	for i := 0; i < %[3]d; i++ {
		c.advance()
	}
	if c.valueOf() != %[4]d {
		t.Errorf("%[1]s after %[3]d moves = %%d, want %[4]d", c.valueOf())
	}
%[5]s}
`, name, start, moves, want, edgeCase)
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goIface renders an interface dispatch: two implementations of one
// interface, selected by a boundary, so a mutant flipping the selection
// swaps the dispatched behavior.
func goIface(stream *rand.Stream, name, level string) goFunctionBody {
	limit := intOperand(stream)
	small := intOperand(stream)
	big := small + intOperand(stream)
	below := 1 + stream.IntN(limit-1)
	above := limit + intOperand(stream)
	edge := pinned(stream, level)
	src := fmt.Sprintf(`
// %[1]sPricer prices one unit.
type %[1]sPricer interface {
	price() int
}

type %[1]sRetail struct{}

func (%[1]sRetail) price() int { return %[2]d }

type %[1]sWholesale struct{}

func (%[1]sWholesale) price() int { return %[3]d }

// %[1]s picks the pricer for quantity n: wholesale at %[4]d and above.
func %[1]s(n int) int {
	var pricer %[1]sPricer = %[1]sRetail{}
	if n >= %[4]d {
		pricer = %[1]sWholesale{}
	}
	return pricer.price()
}
`, name, small, big, limit)
	var edgeCase string
	if edge {
		edgeCase = fmt.Sprintf(`
	if got := %[1]s(%[2]d); got != %[3]d {
		t.Errorf("%[1]s(%[2]d) = %%d, want the wholesale price %[3]d", got)
	}
`, name, limit, big)
	}
	test := fmt.Sprintf(`
func Test%[1]s(t *testing.T) {
	if got := %[1]s(%[2]d); got != %[3]d {
		t.Errorf("%[1]s(%[2]d) = %%d, want %[3]d", got)
	}
%[4]s}
`, name, below, small, edgeCase)
	_ = above
	return goFunctionBody{source: src, test: test, pinned: edge}
}

// goSourceBuilder accumulates the declarations of target.go.
type goSourceBuilder struct {
	pkg   string
	parts []string
}

func (b *goSourceBuilder) addFunction(src string) { b.parts = append(b.parts, src) }

// render produces the final target.go. The errors import is included
// unconditionally: the errorpath family needs it and a fixed import block
// keeps the rendering deterministic without tracking which families fired.
func (b *goSourceBuilder) render() string {
	var sb strings.Builder
	sb.WriteString("// Code generated by abate targetgen. DO NOT EDIT.\n")
	sb.WriteString("// The module is a deterministic mutation-testing target.\n\n")
	fmt.Fprintf(&sb, "package %s\n\n", b.pkg)
	sb.WriteString("import \"errors\"\n\n")
	sb.WriteString("var _ = errors.New\n")
	for _, part := range b.parts {
		sb.WriteString(part)
	}
	return sb.String()
}

// goTestBuilder accumulates the test functions of target_test.go.
type goTestBuilder struct {
	pkg   string
	parts []string
}

func (b *goTestBuilder) addFunction(src string) { b.parts = append(b.parts, src) }

// render produces the final target_test.go with a deterministic header.
func (b *goTestBuilder) render() string {
	var sb strings.Builder
	sb.WriteString("// Code generated by abate targetgen. DO NOT EDIT.\n\n")
	fmt.Fprintf(&sb, "package %s\n\n", b.pkg)
	sb.WriteString("import \"testing\"\n")
	for _, part := range b.parts {
		sb.WriteString(part)
	}
	return sb.String()
}

// countGoTokens approximates the mutable token count of the generated
// source for reporting; it is a report convenience, never a judging input.
func countGoTokens(source []byte) int {
	text := string(source)
	count := 0
	for _, tok := range []string{" + ", " - ", " * ", " / ", " < ", " <= ",
		" > ", " >= ", " == ", " != ", "break", "continue", "true", "false"} {
		count += strings.Count(text, tok)
	}
	return count
}
