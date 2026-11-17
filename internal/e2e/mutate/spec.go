/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"abate/internal/determinism"
)

// minValidPoints is the required number of concrete mutation points an
// operator must declare for every registered language.
const minValidPoints = 30

var (
	errInvalidSpec = errors.New("mutate: invalid operator spec")
	errNoPoints    = errors.New("mutate: no mutation points for language")
)

// languageTags is the canonical order of the seven registered languages.
var languageTags = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// forbiddenMarkers are the operator-family names excluded from the registry.
var forbiddenMarkers = []string{
	"comment", "whitespace", "white-space", "format", "indent",
	"newline", "spacing", "layout", "blank", "tab",
}

// forbiddenTokens are pure comment or comment-like tokens that must never
// become a mutation point.
var forbiddenTokens = []string{"//", "/*", "*/", "#", "<!--", "-->"}

type (
	// MutationPoint is one concrete before/after source rewrite. A point is a
	// real syntactic token pair, never a comment, whitespace or formatting-only
	// change.
	MutationPoint struct {
		Before string
		After  string
	}

	// OperatorSpec describes one mutation operator. Points holds at least
	// MinValidPoints concrete rewrites for every registered language, and Apply
	// is the single scanning entry point over an in-memory buffer.
	OperatorSpec struct {
		ID             string
		Class          MutationClass
		MinValidPoints int
		Points         map[string][]MutationPoint
		Apply          func(OpContext) ([]Edit, error)
	}
)

// Validate checks the operator contract, including the per-language point
// budget and the absence of comment, whitespace and formatting operators.
func (s OperatorSpec) Validate() error {
	if s.ID == "" || !s.Class.Valid() || s.Apply == nil ||
		s.MinValidPoints < minValidPoints || containsForbidden(s.ID) ||
		len(s.Points) != len(languageTags) {
		return fmt.Errorf("%w: %s", errInvalidSpec, s.ID)
	}
	for _, tag := range languageTags {
		points, found := s.Points[tag]
		if !found || len(points) < s.MinValidPoints {
			return fmt.Errorf("%w: %s/%s has %d points", errInvalidSpec,
				s.ID, tag, len(points))
		}
		for _, point := range points {
			if point.Before == "" || point.Before == point.After ||
				strings.TrimSpace(point.Before) == "" ||
				forbiddenToken(point.Before) {
				return fmt.Errorf("%w: %s/%s point %q", errInvalidSpec,
					s.ID, tag, point.Before)
			}
		}
	}
	return nil
}

// Languages returns the language tags of the operator in ascending order.
func (s OperatorSpec) Languages() []string {
	out := make([]string, 0, len(s.Points))
	for tag := range s.Points {
		out = append(out, tag)
	}
	slices.Sort(out)
	return out
}

// PointsFor returns an independent copy of one language's point list.
func (s OperatorSpec) PointsFor(language string) []MutationPoint {
	return slices.Clone(s.Points[language])
}

// Tokens returns the distinct before-tokens of one language, longest first, so
// an operator such as <= is never shadowed by its < prefix.
func (s OperatorSpec) Tokens(language string) []string {
	seen := make(map[string]struct{})
	out := make([]string, 0)
	for _, point := range s.Points[language] {
		if _, exists := seen[point.Before]; !exists {
			seen[point.Before] = struct{}{}
			out = append(out, point.Before)
		}
	}
	slices.SortFunc(out, func(a, b string) int {
		if len(a) != len(b) {
			return len(b) - len(a)
		}
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	return out
}

// containsForbidden reports whether an identifier names a formatting family.
func containsForbidden(id string) bool {
	lower := strings.ToLower(id)
	for _, marker := range forbiddenMarkers {
		if strings.Contains(lower, marker) {
			return true
		}
	}
	return false
}

// forbiddenToken reports whether a token is a pure comment marker.
func forbiddenToken(token string) bool {
	for _, marker := range forbiddenTokens {
		if token == marker {
			return true
		}
	}
	return false
}

// crossPoints builds every distinct before/after combination of two token
// families. Equal pairs are excluded because they are not mutations.
func crossPoints(before, after []string) []MutationPoint {
	out := make([]MutationPoint, 0, len(before)*len(after))
	for _, left := range before {
		for _, right := range after {
			if left != right {
				out = append(out, MutationPoint{Before: left, After: right})
			}
		}
	}
	return out
}

// crossByLanguage applies crossPoints to every registered language.
func crossByLanguage(before, after map[string][]string) map[string][]MutationPoint {
	out := make(map[string][]MutationPoint, len(languageTags))
	for _, tag := range languageTags {
		out[tag] = crossPoints(before[tag], after[tag])
	}
	return out
}

// spreadPoints assigns one shared point family to every registered language.
func spreadPoints(points []MutationPoint) map[string][]MutationPoint {
	out := make(map[string][]MutationPoint, len(languageTags))
	for _, tag := range languageTags {
		out[tag] = slices.Clone(points)
	}
	return out
}

// radixPoints builds count equivalent integer-radix rewrites, for example 16
// to 0x10. Every produced pair has the same integer value, so the operator is
// a genuine equivalent mutation rather than a synthetic padding entry.
func radixPoints(count int) []MutationPoint {
	out := make([]MutationPoint, 0, count)
	for value := 0; value < count; value++ {
		out = append(out, MutationPoint{
			Before: strconv.Itoa(value),
			After:  "0x" + fmt.Sprintf("%X", value),
		})
	}
	return out
}

// newSpec fills the shared defaults and binds the scanning entry point.
func newSpec(base OperatorSpec) OperatorSpec {
	base.MinValidPoints = minValidPoints
	base.Apply = func(ctx OpContext) ([]Edit, error) {
		return scanEdits(ctx, base)
	}
	return base
}

// scanEdits finds every non-overlapping mutation point of one language. The
// SplitMix64 generator only rotates the point order (mutation-point choice);
// it never selects a corpus file.
func scanEdits(ctx OpContext, spec OperatorSpec) ([]Edit, error) {
	if err := ctx.Validate(); err != nil {
		return nil, err
	}
	points := spec.PointsFor(ctx.Language)
	if len(points) == 0 {
		return nil, fmt.Errorf("%w: %s/%s", errNoPoints, spec.ID, ctx.Language)
	}
	rng := determinism.NewRNG(ctx.Seed, ctx.ManifestSHA256,
		spec.ID+":"+ctx.Language)
	start := rng.Intn(len(points))
	found := make([]Edit, 0)
	for i := 0; i < len(points); i++ {
		point := points[(start+i)%len(points)]
		found = append(found, scanPoint(ctx.Content, spec.ID, point)...)
	}
	slices.SortFunc(found, compareEdits)
	return dedupeOverlaps(found), nil
}

// scanPoint collects the bounded occurrences of one point in a byte buffer.
func scanPoint(content []byte, operatorID string, point MutationPoint) []Edit {
	out := make([]Edit, 0)
	if point.Before == "" {
		return out
	}
	needle := []byte(point.Before)
	first, last := needle[0], needle[len(needle)-1]
	for at := 0; at <= len(content); {
		index := bytes.Index(content[at:], needle)
		if index < 0 {
			break
		}
		start := at + index
		end := start + len(needle)
		leftOK := !isWordByte(first) || start == 0 || !isWordByte(content[start-1])
		rightOK := !isWordByte(last) || end == len(content) || !isWordByte(content[end])
		if leftOK && rightOK {
			out = append(out, Edit{OperatorID: operatorID,
				Offset: start, EndOffset: end,
				Before: point.Before, After: point.After})
		}
		at = end
	}
	return out
}

// isWordByte reports whether b can extend an identifier or a number.
func isWordByte(b byte) bool {
	return b == '_' || (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// compareEdits orders edits by source interval and then by replacement text.
func compareEdits(a, b Edit) int {
	if a.Offset != b.Offset {
		return a.Offset - b.Offset
	}
	if a.EndOffset != b.EndOffset {
		return a.EndOffset - b.EndOffset
	}
	if a.After < b.After {
		return -1
	}
	if a.After > b.After {
		return 1
	}
	return 0
}

// dedupeOverlaps keeps the first edit of each byte interval in sorted order.
func dedupeOverlaps(edits []Edit) []Edit {
	out := make([]Edit, 0, len(edits))
	last := 0
	for _, edit := range edits {
		if edit.Offset >= last {
			out = append(out, edit)
			last = edit.EndOffset
		}
	}
	return out
}
