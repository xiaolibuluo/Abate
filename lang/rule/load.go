/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"slices"
	"strings"

	"abate/lang/air"
)

// compiledRule is the immutable, AIR-facing implementation of Rule.
type compiledRule struct {
	spec       Spec
	matcher    Matcher
	predicates []Predicate
}

// Compile validates and compiles a Specification into a Rule.
//
// Compilation performs pattern compilation and predicate parsing once, so the
// matcher can be reused across every file without recompiling anything.
func Compile(spec Spec) (Rule, error) {
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	pattern, err := CompilePattern(spec.Pattern)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrBadPattern, spec.ID, err)
	}
	predicates := make([]Predicate, 0, len(spec.Where))
	for _, clause := range spec.Where {
		predicate, parseErr := ParsePredicate(clause)
		if parseErr != nil {
			return nil, fmt.Errorf("%w: %s: %w", ErrInvalidPredicate,
				spec.ID, parseErr)
		}
		predicates = append(predicates, predicate)
	}
	return compiledRule{
		spec: spec, matcher: NewMatcher(pattern), predicates: predicates,
	}, nil
}

// Spec returns the declaration the rule was compiled from.
func (r compiledRule) Spec() Spec {
	return r.spec
}

// Check reports every match of the rule inside one AIR compilation unit.
func (r compiledRule) Check(unit *air.Unit) []Finding {
	if unit == nil {
		return nil
	}
	out := make([]Finding, 0)
	for i, root := range unit.Roots {
		if root == nil {
			continue
		}
		path := ""
		if i < len(unit.Files) {
			path = unit.Files[i].Path
		}
		r.walk(root, path, &out)
	}
	sortFindings(out)
	return out
}

// walk visits one file root and appends filtered findings.
func (r compiledRule) walk(root air.Node, path string, out *[]Finding) {
	air.Walk(root, func(node air.Node) bool {
		caps := r.matcher.Match(node)
		if len(caps) == 0 {
			return true
		}
		for _, predicate := range r.predicates {
			if !predicate.evaluateWithPath(node, caps, path) {
				return true
			}
		}
		pos := node.Pos()
		*out = append(*out, Finding{
			RuleID: r.spec.ID, Path: path,
			Line: pos.Line, Column: pos.Column,
			Severity: r.spec.Severity, Category: r.spec.Category,
			Message:   expandMessage(r.spec.Message, caps),
			Reference: r.spec.Reference, Captures: caps,
		})
		return true
	})
}

// expandMessage substitutes metavariable references in a message template.
func expandMessage(template string, caps []Capture) string {
	ordered := slices.Clone(caps)
	slices.SortStableFunc(ordered, func(a, b Capture) int {
		return len(b.Name) - len(a.Name)
	})
	out := template
	for _, c := range ordered {
		out = strings.ReplaceAll(out, "$"+c.Name, c.Text)
	}
	return out
}

// sortFindings orders findings by file, line, column, then rule identifier.
func sortFindings(findings []Finding) {
	slices.SortStableFunc(findings, func(a, b Finding) int {
		return compareFindings(a, b)
	})
}

// compareFindings implements the canonical finding ordering.
func compareFindings(a, b Finding) int {
	if c := strings.Compare(a.Path, b.Path); c != 0 {
		return c
	}
	if a.Line != b.Line {
		return a.Line - b.Line
	}
	if a.Column != b.Column {
		return a.Column - b.Column
	}
	return strings.Compare(string(a.RuleID), string(b.RuleID))
}
