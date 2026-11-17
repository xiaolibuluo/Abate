/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"abate/lang/rule"
)

// DefaultTag is the language-neutral namespace of the standalone template set.
const DefaultTag = "COM"

// sourceAIR is the frozen SOURCE segment of every generated rule identifier.
const sourceAIR = "AIR"

// maxTemplateOrdinal bounds the three-digit sequence segment of a rule id.
const maxTemplateOrdinal = 999

// ErrInvalidTag reports a language tag that cannot form a rule-id namespace.
var ErrInvalidTag = errors.New("common: invalid language tag")

// ErrTooManyTemplates reports a batch that overflows the three-digit id space.
var ErrTooManyTemplates = errors.New("common: too many templates")

// Generate builds the common template family plus language-local templates.
//
// The declarations are sorted by identifier, deduplicated by declaration key,
// and each candidate passes Spec.Validate, CompilePattern and ParsePredicate
// before it is returned. A batch containing one invalid template yields nil as
// an explicit failure signal rather than silently dropping that template; use
// GenerateChecked to obtain the precise validation error.
func Generate(tag string, local []Template) []rule.Spec {
	specs, err := GenerateChecked(tag, local)
	if err != nil {
		return nil
	}
	return specs
}

// GenerateChecked is the error-reporting form of Generate.
//
// It is the single entry point used by LoadBuiltins, so a malformed template
// can never be silently swallowed on its way into a rule.Registry.
func GenerateChecked(tag string, local []Template) ([]rule.Spec, error) {
	namespace, err := normalizeTag(tag)
	if err != nil {
		return nil, err
	}
	merged := make([]Template, 0, templateCount+len(local))
	merged = append(merged, builtinTemplates()...)
	merged = append(merged, local...)
	templates := dedupeTemplates(merged)
	if len(templates) > maxTemplateOrdinal {
		return nil, fmt.Errorf("%w: %d templates", ErrTooManyTemplates, len(templates))
	}
	out := make([]rule.Spec, 0, len(templates))
	for i, t := range templates {
		ref, ok := ReferenceFor(tag, t.Family)
		if !ok {
			return nil, fmt.Errorf("common: no reference for family %q", t.Family)
		}
		spec := t.toSpec(namespace, i+1, ref.Section)
		if err := spec.Validate(); err != nil {
			return nil, fmt.Errorf("common: template %d (%s) fails schema: %w",
				i+1, t.Family, err)
		}
		if _, err := rule.CompilePattern(spec.Pattern); err != nil {
			return nil, fmt.Errorf("common: template %d (%s) has a bad pattern: %w",
				i+1, t.Family, err)
		}
		for _, clause := range spec.Where {
			if _, err := rule.ParsePredicate(clause); err != nil {
				return nil, fmt.Errorf("common: template %d (%s) has a bad clause %q: %w",
					i+1, t.Family, clause, err)
			}
		}
		out = append(out, spec)
	}
	slices.SortFunc(out, func(a, b rule.Spec) int {
		return strings.Compare(string(a.ID), string(b.ID))
	})
	return out, nil
}

// normalizeTag uppercases a short language tag and checks its id namespace.
func normalizeTag(tag string) (string, error) {
	text := strings.ToUpper(strings.TrimSpace(tag))
	if text == "" {
		text = DefaultTag
	}
	for i := 0; i < len(text); i++ {
		c := text[i]
		if (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			return "", fmt.Errorf("%w: %q", ErrInvalidTag, tag)
		}
	}
	return text, nil
}

// dedupeTemplates keeps the first occurrence of each declaration key.
//
// The seen map is only used for point lookups; output order comes exclusively
// from the input slice, so no map iteration order can leak into a rule id.
func dedupeTemplates(templates []Template) []Template {
	seen := make(map[string]struct{}, len(templates))
	out := make([]Template, 0, len(templates))
	for _, t := range templates {
		key := t.Key()
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, t)
	}
	return out
}
