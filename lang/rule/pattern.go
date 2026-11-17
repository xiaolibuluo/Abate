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
)

// Pattern is a compiled structural pattern ready to build a Matcher.
//
// The pattern language supports all five structural features: "$_" matches any
// single vertex, "$x" captures one vertex, "$.." skips descendants at any
// depth, "!P" negates a vertex pattern, and "$x*" repeats a capture zero or
// more times. A kind word may carry an attribute constraint using "=", "!=",
// "~", "^", "+", or "/regex/".
type Pattern struct {
	source string
	root   *patternNode
}

// CompilePattern validates and compiles a structural pattern source string.
func CompilePattern(source string) (Pattern, error) {
	root, err := parsePatternSource(source)
	if err != nil {
		return Pattern{}, err
	}
	return Pattern{source: source, root: root}, nil
}

// Source returns the original pattern text.
func (p Pattern) Source() string {
	return p.source
}

// Valid reports whether the pattern has a compiled root vertex.
func (p Pattern) Valid() bool {
	return p.root != nil
}

// Metavariables returns the capture names in ascending order.
func (p Pattern) Metavariables() []string {
	if !p.Valid() {
		return nil
	}
	return patternCaptureNames(p.root)
}

// Features returns the distinct structural features used by the pattern in
// ascending order.
func (p Pattern) Features() []string {
	seen := make(map[string]struct{}, 5)
	collectFeatures(p.root, seen)
	out := make([]string, 0, len(seen))
	for name := range seen {
		out = append(out, name)
	}
	slices.Sort(out)
	return out
}

// collectFeatures walks a compiled pattern into a feature set.
func collectFeatures(node *patternNode, seen map[string]struct{}) {
	if node == nil {
		return
	}
	markFeature(node, seen)
	collectFeatures(node.inner, seen)
	for _, child := range node.children {
		collectFeatures(child, seen)
	}
}

// markFeature records the features present on one vertex.
func markFeature(node *patternNode, seen map[string]struct{}) {
	if node.wild {
		seen["wildcard"] = struct{}{}
	}
	if node.capture != "" && !node.repeat {
		seen["capture"] = struct{}{}
	}
	if node.repeat {
		seen["repeat"] = struct{}{}
	}
	if node.ellipsis {
		seen["deep"] = struct{}{}
	}
	if node.neg {
		seen["negative"] = struct{}{}
	}
}

// String renders the pattern source for diagnostics.
func (p Pattern) String() string {
	return fmt.Sprintf("pattern(%q)", p.source)
}
