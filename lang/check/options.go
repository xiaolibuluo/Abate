/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"fmt"
	"slices"
	"strings"

	"abate/lang/rule"
)

// Report formats supported by the checker.
const (
	// FormatConsole is the human-readable line-oriented rendering.
	FormatConsole = "console"
	// FormatJSON is the deterministic intermediate JSON rendering.
	FormatJSON = "json"
	// FormatSARIF is the minimal deterministic SARIF 2.1.0 rendering.
	FormatSARIF = "sarif"
)

// Options filters and configures one check invocation.
type Options struct {
	// Languages restricts the run to lowercase language tags.
	Languages []string
	// Rules restricts the run to rule identifier prefixes.
	Rules []string
	// MinSeverity drops findings below this severity rank.
	MinSeverity rule.Severity
	// Format selects the report rendering.
	Format string
}

// DefaultOptions returns the console-format, unfiltered options.
func DefaultOptions() Options {
	return Options{Format: FormatConsole, MinSeverity: rule.SeverityInfo}
}

// Normalize returns a canonical copy with sorted, deduplicated filters.
func (o Options) Normalize() Options {
	out := o
	out.Languages = normalizeList(o.Languages)
	out.Rules = normalizeList(o.Rules)
	out.Format = strings.ToLower(strings.TrimSpace(o.Format))
	if out.Format == "" {
		out.Format = FormatConsole
	}
	if !out.MinSeverity.Valid() {
		out.MinSeverity = rule.SeverityInfo
	}
	return out
}

// Validate checks the option invariants before a run starts.
func (o Options) Validate() error {
	switch o.Format {
	case FormatConsole, FormatJSON, FormatSARIF:
	default:
		return fmt.Errorf("%w: %q", ErrInvalidFormat, o.Format)
	}
	if !o.MinSeverity.Valid() {
		return fmt.Errorf("%w: bad severity %q", ErrInvalidOptions, o.MinSeverity)
	}
	for _, language := range o.Languages {
		if strings.TrimSpace(language) == "" {
			return fmt.Errorf("%w: empty language filter", ErrInvalidOptions)
		}
	}
	for _, prefix := range o.Rules {
		if strings.TrimSpace(prefix) == "" {
			return fmt.Errorf("%w: empty rule filter", ErrInvalidOptions)
		}
	}
	return nil
}

// Accepts reports whether a rule and a finding survive the option filters.
func (o Options) Accepts(spec rule.Spec, severity rule.Severity) bool {
	if severity.Rank() < o.MinSeverity.Rank() {
		return false
	}
	if len(o.Languages) > 0 && !slices.Contains(o.Languages, spec.Language()) {
		return false
	}
	if len(o.Rules) == 0 {
		return true
	}
	identifier := strings.ToLower(string(spec.ID))
	for _, prefix := range o.Rules {
		if strings.HasPrefix(identifier, prefix) {
			return true
		}
	}
	return false
}

// normalizeList trims, lowercases, sorts and deduplicates a filter list.
func normalizeList(items []string) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		clean := strings.ToLower(strings.TrimSpace(item))
		if clean != "" {
			out = append(out, clean)
		}
	}
	slices.Sort(out)
	return slices.Compact(out)
}
