/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"strconv"
	"strings"
)

// Finding is one rule match reported against an exact source location.
type Finding struct {
	// RuleID names the rule that produced the finding.
	RuleID RuleID
	// Path is the virtual source path of the matched unit.
	Path string
	// Line is the one-based start line of the matched node.
	Line int
	// Column is the one-based start column of the matched node.
	Column int
	// Severity ranks the finding.
	Severity Severity
	// Category groups the finding by defect class.
	Category Category
	// Message explains the violation.
	Message string
	// Reference cites the official specification entry.
	Reference string
	// Captures holds the metavariable bindings of the match.
	Captures []Capture
}

// Key returns the deterministic deduplication key of the finding.
func (f Finding) Key() string {
	return string(f.RuleID) + "\x00" + f.Path + "\x00" +
		strconv.Itoa(f.Line) + "\x00" + strconv.Itoa(f.Column)
}

// String renders the finding in the stable console form.
func (f Finding) String() string {
	return fmt.Sprintf("%s:%d:%d: %s: %s [%s]", f.Path, f.Line,
		f.Column, f.RuleID, f.Message, f.Severity)
}

// Canonical renders the finding in a byte-stable form used by fingerprints.
func (f Finding) Canonical() string {
	parts := make([]string, 0, 9+len(f.Captures))
	parts = append(parts, string(f.RuleID), f.Path,
		strconv.Itoa(f.Line), strconv.Itoa(f.Column),
		f.Severity.String(), f.Category.String(), f.Message, f.Reference)
	for _, c := range f.Captures {
		parts = append(parts, c.Canonical())
	}
	return strings.Join(parts, "\x00")
}
