/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"abate/lang/air"
)

// Predicate is one compiled where clause evaluated against a match.
//
// The clause syntax is "<op>:<arg>", except for the argument-free "distinct"
// operator. Ops fall into five families: capture type constraints, value
// matching with regular expressions or numeric ranges, capture cardinality,
// semantic annotation checks, and source-path scope checks.
type Predicate struct {
	// Op names the predicate operation.
	Op string
	// Arg is the operation argument.
	Arg string

	regex *regexp.Regexp
}

// ParsePredicate compiles one where clause into a Predicate.
func ParsePredicate(clause string) (Predicate, error) {
	text := strings.TrimSpace(clause)
	if text == "" {
		return Predicate{}, fmt.Errorf("%w: empty clause", ErrInvalidPredicate)
	}
	op, arg, _ := strings.Cut(text, ":")
	if !predicateOpKnown(op) {
		return Predicate{}, fmt.Errorf("%w: unknown op %q", ErrInvalidPredicate, op)
	}
	if op != "distinct" && arg == "" {
		return Predicate{}, fmt.Errorf("%w: %s needs an argument", ErrInvalidPredicate, op)
	}
	out := Predicate{Op: op, Arg: arg}
	if op == "regex" || op == "text-not-regex" || op == "path-not-regex" {
		compiled, err := regexp.Compile(arg)
		if err != nil {
			return Predicate{}, fmt.Errorf("%w: %w", ErrInvalidPredicate, err)
		}
		out.regex = compiled
	}
	if op == "num-min" || op == "num-max" || op == "count-min" || op == "count-max" {
		if _, err := strconv.ParseFloat(arg, 64); err != nil {
			return Predicate{}, fmt.Errorf("%w: %s needs a number: %w",
				ErrInvalidPredicate, op, err)
		}
	}
	return out, nil
}

// Valid reports whether the predicate carries a supported operation.
func (p Predicate) Valid() bool {
	return predicateOpKnown(p.Op)
}

// String rebuilds the clause text in a stable form.
func (p Predicate) String() string {
	if p.Op == "distinct" {
		return p.Op
	}
	return p.Op + ":" + p.Arg
}

// Evaluate reports whether a match satisfies the predicate.
//
// Path-scoped predicates see an empty source path here; the rule engine
// supplies the real file path through evaluateWithPath.
func (p Predicate) Evaluate(node air.Node, caps []Capture) bool {
	return p.evaluateWithPath(node, caps, "")
}

// evaluateWithPath reports whether a match satisfies the predicate for one source path.
func (p Predicate) evaluateWithPath(node air.Node, caps []Capture, path string) bool {
	if node == nil || !p.Valid() {
		return false
	}
	switch p.Op {
	case "kind", "kind-not":
		return p.matchKind(node, caps)
	case "text", "text~", "text^", "text+", "text-not", "regex":
		return p.matchText(node, caps)
	case "text-not-regex":
		return p.matchTextNotRegex(node, caps)
	case "path-not-regex":
		return p.matchPathNotRegex(path)
	case "num-min", "num-max":
		return p.matchNumber(node, caps)
	case "count-min", "count-max":
		return p.matchCount(caps)
	case "sem":
		return hasSemantic(node, p.Arg)
	case "capture":
		return hasCapture(caps, p.Arg)
	case "distinct":
		return capturesDistinct(caps)
	default:
		return false
	}
}

// matchKind applies a capture type constraint.
func (p Predicate) matchKind(node air.Node, caps []Capture) bool {
	kinds := make([]string, 0, len(caps)+1)
	kinds = append(kinds, node.Kind().String())
	for _, c := range caps {
		kinds = append(kinds, c.Kind)
	}
	for _, k := range kinds {
		if p.Op == "kind" && k == p.Arg {
			return true
		}
		if p.Op == "kind-not" && k == p.Arg {
			return false
		}
	}
	return p.Op == "kind-not"
}

// matchText applies an attribute value constraint.
func (p Predicate) matchText(node air.Node, caps []Capture) bool {
	texts := captureTexts(node, caps)
	for _, text := range texts {
		if p.textMatches(text) {
			return true
		}
	}
	return false
}

// textMatches reports whether one attribute text satisfies the op.
func (p Predicate) textMatches(text string) bool {
	switch p.Op {
	case "text":
		return text == p.Arg
	case "text~":
		return strings.Contains(text, p.Arg)
	case "text^":
		return strings.HasPrefix(text, p.Arg)
	case "text+":
		return strings.HasSuffix(text, p.Arg)
	case "text-not":
		return text != p.Arg
	case "regex":
		return p.regex != nil && p.regex.MatchString(text)
	default:
		return false
	}
}

// matchTextNotRegex keeps a match only when no captured attribute text matches
// the predicate regex.
func (p Predicate) matchTextNotRegex(node air.Node, caps []Capture) bool {
	if p.regex == nil {
		return false
	}
	for _, text := range captureTexts(node, caps) {
		if p.regex.MatchString(text) {
			return false
		}
	}
	return true
}

// matchPathNotRegex keeps a match only when the source path falls outside the
// scope named by the predicate regex.
func (p Predicate) matchPathNotRegex(path string) bool {
	return p.regex != nil && !p.regex.MatchString(path)
}

// matchNumber applies a numeric range constraint to literal text.
func (p Predicate) matchNumber(node air.Node, caps []Capture) bool {
	limit, err := strconv.ParseFloat(p.Arg, 64)
	if err != nil {
		return false
	}
	for _, text := range captureTexts(node, caps) {
		value, parseErr := strconv.ParseFloat(strings.ReplaceAll(text, "_", ""), 64)
		if parseErr != nil {
			continue
		}
		if p.Op == "num-min" && value >= limit {
			return true
		}
		if p.Op == "num-max" && value <= limit {
			return true
		}
	}
	return false
}

// matchCount applies a capture cardinality constraint.
func (p Predicate) matchCount(caps []Capture) bool {
	limit, err := strconv.ParseFloat(p.Arg, 64)
	if err != nil {
		return false
	}
	if p.Op == "count-min" {
		return float64(len(caps)) >= limit
	}
	return float64(len(caps)) <= limit
}

// captureTexts returns the matched attribute texts, falling back to the node.
func captureTexts(node air.Node, caps []Capture) []string {
	out := make([]string, 0, len(caps)+1)
	for _, c := range caps {
		out = append(out, c.Text)
	}
	if len(out) == 0 {
		out = append(out, air.AttributeOf(node))
	}
	return out
}

// hasCapture reports whether a named metavariable was bound.
func hasCapture(caps []Capture, name string) bool {
	for _, c := range caps {
		if c.Name == name {
			return true
		}
	}
	return false
}

// capturesDistinct reports whether all captured texts differ.
func capturesDistinct(caps []Capture) bool {
	seen := make(map[string]struct{}, len(caps))
	for _, c := range caps {
		if _, ok := seen[c.Text]; ok {
			return false
		}
		seen[c.Text] = struct{}{}
	}
	return true
}

// hasSemantic reports whether the node carries the named AIR annotation.
func hasSemantic(node air.Node, name string) bool {
	for key := air.SemTypeRef; key <= air.SemKeyLast; key++ {
		if strings.EqualFold(key.String(), name) {
			return node.Sem(key) != nil
		}
	}
	return false
}

// predicateOpKnown reports whether an op name is implemented.
func predicateOpKnown(op string) bool {
	switch op {
	case "kind", "kind-not", "text", "text~", "text^", "text+",
		"text-not", "regex", "text-not-regex", "path-not-regex",
		"num-min", "num-max", "count-min", "count-max", "sem",
		"capture", "distinct":
		return true
	default:
		return false
	}
}
