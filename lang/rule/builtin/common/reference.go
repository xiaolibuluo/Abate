/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"slices"
	"strings"
)

// Reference is one traceable citation into a repository-archived specification.
//
// Section always names an archived docs/reference file plus a chapter title, so
// every generated rule carries a non-empty, offline-verifiable source.
type Reference struct {
	// Tag is the lowercase short language tag the citation belongs to.
	Tag string
	// Family is the defect-class theme the citation is bound to.
	Family string
	// Section is the archived chapter name backing the rule.
	Section string
}

// Cite returns the archived chapter name of the citation.
func (r Reference) Cite() string {
	return r.Section
}

// chapter is one language-to-archive mapping row.
type chapter struct {
	tag   string
	doc   string
	title string
}

// section is one theme-to-chapter mapping row.
type section struct {
	family string
	title  string
}

// defaultReferenceTag is the language-neutral fallback citation source.
const defaultReferenceTag = "cert"

// languageChapters is the frozen six-language plus fallback chapter table.
var languageChapters = []chapter{
	{tag: "go", doc: "docs/reference/effective-go.md", title: "Effective Go / Names"},
	{tag: "py", doc: "docs/reference/python-pep8.md", title: "PEP 8 / Naming Conventions"},
	{tag: "js", doc: "docs/reference/js-airbnb.md", title: "Airbnb JavaScript Style Guide / Comparison Operators"},
	{tag: "c", doc: "docs/reference/c-cert.md", title: "CERT C / Expressions"},
	{tag: "rs", doc: "docs/reference/rust-api.md", title: "Rust API Guidelines / Naming"},
	{tag: "cs", doc: "docs/reference/csharp-design.md", title: ".NET Design Guidelines / Capitalization"},
	{tag: "kt", doc: "docs/reference/kotlin-conventions.md", title: "Kotlin Conventions / Naming Rules"},
	{tag: defaultReferenceTag, doc: "docs/reference/c-cert.md", title: "CERT C / General"},
}

// familySections maps each theme to an archived section title.
var familySections = []section{
	{family: "naming", title: "Naming"},
	{family: "empty-capture", title: "Declarations and Initialization"},
	{family: "magic-number", title: "Literals and Constants"},
	{family: "error-ignored", title: "Error Handling"},
	{family: "comparison-operator", title: "Operators"},
	{family: "bare-dereference", title: "Pointer Usage"},
	{family: "bool-literal-compare", title: "Boolean Expressions"},
	{family: "constant-naming", title: "Constant Naming"},
	{family: "long-parameters", title: "Method Signatures"},
	{family: "unused-binding", title: "Declarations and Initialization"},
	{family: "comment-placeholder", title: "Comments"},
	{family: "boundary-compare", title: "Boundary Conditions"},
}

// findChapter looks up one archived chapter by normalized short tag.
func findChapter(tag string) (chapter, bool) {
	for _, item := range languageChapters {
		if item.tag == tag {
			return item, true
		}
	}
	return chapter{}, false
}

// findSection looks up one archived section by theme name.
func findSection(familyName string) (section, bool) {
	for _, item := range familySections {
		if item.family == familyName {
			return item, true
		}
	}
	return section{}, false
}

// ReferenceFor resolves the citation for one language tag and theme family.
//
// An unknown tag falls back to the language-neutral CERT chapter, and an
// unknown family falls back to the General section, so the result is always
// non-empty and traceable instead of failing the whole generation batch.
func ReferenceFor(tag, familyName string) (Reference, bool) {
	normalized := strings.ToLower(strings.TrimSpace(tag))
	ch, ok := findChapter(normalized)
	if !ok {
		ch, ok = findChapter(defaultReferenceTag)
		if !ok {
			return Reference{}, false
		}
	}
	sec, found := findSection(familyName)
	if !found {
		sec = section{family: familyName, title: "General"}
	}
	return Reference{
		Tag:     ch.tag,
		Family:  familyName,
		Section: ch.doc + " — " + ch.title + " / " + sec.title,
	}, true
}

// ReferenceTable returns the ordered cross-product of chapters and sections.
//
// Iteration follows two frozen slices, never a map, so the table is byte-stable
// across repeated calls and across concurrent readers.
func ReferenceTable() []Reference {
	out := make([]Reference, 0, len(languageChapters)*len(familySections))
	for _, ch := range languageChapters {
		for _, sec := range familySections {
			ref, ok := ReferenceFor(ch.tag, sec.family)
			if !ok {
				continue
			}
			out = append(out, ref)
		}
	}
	slices.SortFunc(out, func(a, b Reference) int {
		if c := strings.Compare(a.Tag, b.Tag); c != 0 {
			return c
		}
		if c := strings.Compare(a.Family, b.Family); c != 0 {
			return c
		}
		return strings.Compare(a.Section, b.Section)
	})
	return out
}
