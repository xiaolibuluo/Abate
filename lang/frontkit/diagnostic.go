/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"fmt"
	"slices"
	"strings"

	"abate/lang/air"
)

// sinkEntry keeps the byte offset used by the deterministic sort key; the
// public air.Diagnostic type carries only line and column coordinates.
type sinkEntry struct {
	diagnostic air.Diagnostic
	offset     int
}

// DiagnosticSink collects structured, language-tagged parsing messages for one
// virtual source path. Entries are emitted in a canonical (Path, Offset, Code)
// order and never exceed air.MaxDiagnostics, so damaged input cannot flood a
// downstream report. The sink is a single-owner value, not a shared component.
type DiagnosticSink struct {
	path      string
	tag       string
	entries   []sinkEntry
	truncated bool
}

// NewDiagnosticSink returns an empty sink whose codes use the canonical short
// tag in upper case. Unknown tags fall back to a sanitized ASCII upper-case
// spelling, and an unusable tag falls back to LANG.
func NewDiagnosticSink(path, tag string) *DiagnosticSink {
	return &DiagnosticSink{
		path: path, tag: diagnosticTag(tag),
		entries: make([]sinkEntry, 0),
	}
}

// Error records an error-severity message with a LANG-<TAG>-E-xxx code.
func (s *DiagnosticSink) Error(code, offset, line, column int, message string) error {
	return s.add(code, offset, line, column, air.SeverityError, "E", message)
}

// Warning records a warning-severity message with a LANG-<TAG>-W-xxx code.
func (s *DiagnosticSink) Warning(code, offset, line, column int, message string) error {
	return s.add(code, offset, line, column, air.SeverityWarning, "W", message)
}

// add normalizes one message and enforces the fixed diagnostic ceiling.
func (s *DiagnosticSink) add(code, offset, line, column int,
	severity, letter, message string) error {
	if s == nil {
		return fmt.Errorf("%w: nil sink", ErrDiagnosticLimit)
	}
	if len(s.entries) >= air.MaxDiagnostics {
		s.truncated = true
		return fmt.Errorf("%w: %d diagnostics", ErrDiagnosticLimit, air.MaxDiagnostics)
	}
	code = clampInt(code, 0, 999)
	offset = clampInt(offset, 0, 1<<30)
	line = clampInt(line, 1, 1<<30)
	column = clampInt(column, 1, 1<<30)
	s.entries = append(s.entries, sinkEntry{
		diagnostic: air.Diagnostic{
			Path: s.path, Line: line, Column: column,
			Severity: severity,
			Code:     fmt.Sprintf("LANG-%s-%s-%03d", s.tag, letter, code),
			Message:  message,
		},
		offset: offset,
	})
	return nil
}

// Diagnostics returns the canonical, capped snapshot of collected messages.
func (s *DiagnosticSink) Diagnostics() []air.Diagnostic {
	if s == nil {
		return []air.Diagnostic{}
	}
	entries := slices.Clone(s.entries)
	slices.SortStableFunc(entries, func(a, b sinkEntry) int {
		if order := strings.Compare(a.diagnostic.Path, b.diagnostic.Path); order != 0 {
			return order
		}
		if a.offset != b.offset {
			if a.offset < b.offset {
				return -1
			}
			return 1
		}
		return strings.Compare(a.diagnostic.Code, b.diagnostic.Code)
	})
	if len(entries) > air.MaxDiagnostics {
		entries = entries[:air.MaxDiagnostics]
	}
	out := make([]air.Diagnostic, 0, len(entries))
	for _, entry := range entries {
		out = append(out, entry.diagnostic)
	}
	return out
}

// IsError reports whether the sink holds at least one error-severity message.
func (s *DiagnosticSink) IsError() bool {
	if s == nil {
		return false
	}
	for _, entry := range s.entries {
		if entry.diagnostic.IsError() {
			return true
		}
	}
	return false
}

// Len returns the number of retained messages.
func (s *DiagnosticSink) Len() int {
	if s == nil {
		return 0
	}
	return len(s.entries)
}

// Truncated reports whether the diagnostic ceiling rejected a message.
func (s *DiagnosticSink) Truncated() bool {
	if s == nil {
		return false
	}
	return s.truncated
}

// Tag returns the upper-case tag segment used in every generated code.
func (s *DiagnosticSink) Tag() string {
	if s == nil {
		return "LANG"
	}
	return s.tag
}

// diagnosticTag normalizes one tag spelling into a safe code segment.
func diagnosticTag(tag string) string {
	if canonical, ok := CanonicalTag(tag); ok {
		return strings.ToUpper(canonical)
	}
	out := make([]byte, 0, len(tag))
	for i := 0; i < len(tag); i++ {
		ch := tag[i]
		if ch >= 'a' && ch <= 'z' {
			ch -= 'a' - 'A'
		}
		if (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			out = append(out, ch)
		}
	}
	if len(out) == 0 {
		return "LANG"
	}
	return string(out)
}
