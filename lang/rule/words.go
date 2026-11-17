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
	"slices"
	"strings"

	"abate/lang/air"
)

// maxPatternDepth bounds recursive pattern compilation so malformed input
// cannot exhaust the goroutine stack.
const maxPatternDepth = 32

// parseWord compiles one word into a pattern vertex.
func parseWord(text string) (*patternNode, error) {
	switch {
	case text == "$_":
		return &patternNode{wild: true}, nil
	case text == "$..":
		return &patternNode{ellipsis: true}, nil
	case text == "ANY":
		return &patternNode{kind: "ANY"}, nil
	case strings.HasPrefix(text, "!"):
		return parseNegation(text)
	case strings.HasPrefix(text, "@"):
		key, err := lookupSemKey(text[1:])
		if err != nil {
			return nil, err
		}
		return &patternNode{semKey: key}, nil
	case strings.HasPrefix(text, "#"):
		if len(text) < 2 {
			return nil, fmt.Errorf("%w: empty attribute", ErrBadPattern)
		}
		return &patternNode{attrOp: "=", attrVal: text[1:]}, nil
	case strings.HasPrefix(text, "$"):
		return parseCapture(text)
	default:
		return parseConstraint(text)
	}
}

// parseNegation compiles a negated vertex pattern.
func parseNegation(text string) (*patternNode, error) {
	if len(text) < 2 {
		return nil, fmt.Errorf("%w: empty negation", ErrBadPattern)
	}
	inner, err := parseWord(text[1:])
	if err != nil {
		return nil, err
	}
	return &patternNode{neg: true, inner: inner}, nil
}

// parseCapture compiles a metavariable, optionally with a repeat marker.
func parseCapture(text string) (*patternNode, error) {
	name := text[1:]
	repeat := strings.HasSuffix(name, "*")
	if repeat {
		name = name[:len(name)-1]
	}
	if !validCaptureName(name) {
		return nil, fmt.Errorf("%w: bad metavariable %q", ErrBadPattern, text)
	}
	return &patternNode{capture: name, repeat: repeat}, nil
}

// parseConstraint compiles a kind word with an optional attribute constraint.
func parseConstraint(text string) (*patternNode, error) {
	kind, op, value := splitConstraint(text)
	if kind == "" || !attrOpKnown(op) {
		return nil, fmt.Errorf("%w: bad constraint %q", ErrBadPattern, text)
	}
	node := &patternNode{kind: kind, attrOp: op, attrVal: value}
	if op == "" {
		return node, nil
	}
	if value == "" {
		return nil, fmt.Errorf("%w: empty value in %q", ErrBadPattern, text)
	}
	if op == "regex" {
		compiled, err := regexp.Compile(value)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrBadPattern, err)
		}
		node.regex = compiled
	}
	return node, nil
}

// splitConstraint separates a kind name from its attribute operator and value.
func splitConstraint(text string) (kind, op, value string) {
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '=', '~', '^', '+':
			return text[:i], string(text[i]), text[i+1:]
		case '!':
			if i+1 < len(text) && text[i+1] == '=' {
				return text[:i], "!=", text[i+2:]
			}
		case '/':
			rest := text[i+1:]
			if strings.HasSuffix(rest, "/") {
				return text[:i], "regex", rest[:len(rest)-1]
			}
			return text[:i], "unterminated", ""
		}
	}
	return text, "", ""
}

// attrOpKnown reports whether an attribute operator is implemented.
func attrOpKnown(op string) bool {
	switch op {
	case "", "=", "~", "^", "+", "!=", "regex":
		return true
	default:
		return false
	}
}

// validCaptureName validates a metavariable name.
func validCaptureName(name string) bool {
	if name == "" || !isFirstByte(name[0]) {
		return false
	}
	for i := 1; i < len(name); i++ {
		if !isRestByte(name[i]) {
			return false
		}
	}
	return true
}

// isFirstByte reports whether a byte may start a metavariable name.
func isFirstByte(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isRestByte reports whether a byte may continue a metavariable name.
func isRestByte(c byte) bool {
	return isFirstByte(c) || (c >= '0' && c <= '9')
}

// lookupSemKey resolves an AIR annotation name to its wire key.
func lookupSemKey(name string) (air.SemKey, error) {
	for key := air.SemTypeRef; key <= air.SemKeyLast; key++ {
		if strings.EqualFold(key.String(), name) {
			return key, nil
		}
	}
	return air.SemInvalid, fmt.Errorf("%w: unknown annotation %q", ErrBadPattern, name)
}

// validatePatternTree checks structural invariants of a compiled pattern.
func validatePatternTree(node *patternNode, depth int) error {
	if depth > maxPatternDepth {
		return fmt.Errorf("%w: pattern too deep", ErrBadPattern)
	}
	if (node.ellipsis || node.repeat) && len(node.children) > 0 {
		return fmt.Errorf("%w: marker cannot own children", ErrBadPattern)
	}
	if node.neg {
		if node.inner == nil {
			return fmt.Errorf("%w: empty negation", ErrBadPattern)
		}
		if err := validatePatternTree(node.inner, depth+1); err != nil {
			return err
		}
	}
	for _, child := range node.children {
		if err := validatePatternTree(child, depth+1); err != nil {
			return err
		}
	}
	return nil
}

// patternCaptureNames collects the metavariables of a pattern in sorted order.
func patternCaptureNames(node *patternNode) []string {
	seen := make(map[string]struct{}, 4)
	collectCaptureNames(node, seen)
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

// collectCaptureNames walks a pattern tree into a name set.
func collectCaptureNames(node *patternNode, seen map[string]struct{}) {
	if node == nil {
		return
	}
	if node.capture != "" {
		seen[node.capture] = struct{}{}
	}
	collectCaptureNames(node.inner, seen)
	for _, child := range node.children {
		collectCaptureNames(child, seen)
	}
}
