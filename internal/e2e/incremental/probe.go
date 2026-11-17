/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package incremental

import (
	"errors"

	"abate/lang/air"
)

var (
	errNilFrontend = errors.New("incremental: nil frontend")
	errNilUnit     = errors.New("incremental: nil prior unit")
)

type (
	// Result summarises one incremental-parse consistency check.
	// ChangedRanges counts distinct edited byte intervals.
	// Consistent reports whether every unedited file kept its AIR root identifier.
	Result struct {
		ChangedRanges int
		Consistent    bool
	}

	// Probe checks edit-region consistency through ParseIncremental.
	Probe struct {
		frontend air.Frontend
	}
)

// NewProbe builds a probe around one language frontend.
func NewProbe(frontend air.Frontend) Probe {
	return Probe{frontend: frontend}
}

// Check reparses unit with edits and verifies the fresh unit validates, that a
// non-empty edit produced a node-level diff, and that every file outside the
// edited ranges keeps its AIR root identifier.
func (p Probe) Check(unit air.SourceUnit, old *air.Unit,
	edits []air.Edit) (Result, error) {
	if p.frontend == nil {
		return Result{}, errNilFrontend
	}
	if old == nil {
		return Result{}, errNilUnit
	}
	if err := unit.Validate(); err != nil {
		return Result{}, err
	}
	fresh, diff, err := p.frontend.ParseIncremental(unit, old, edits)
	if err != nil {
		return Result{}, err
	}
	result := Result{
		ChangedRanges: countRanges(edits),
		Consistent:    true,
	}
	if fresh == nil || fresh.Validate() != nil {
		result.Consistent = false
	}
	if len(edits) > 0 && diff.Empty() {
		result.Consistent = false
	}
	edited := pathsOf(edits)
	for _, file := range unit.Files {
		if !edited[file.Path] {
			oldRoot, oldOK := rootOf(old, file.Path)
			newRoot, newOK := rootOf(fresh, file.Path)
			if !oldOK || !newOK ||
				air.IDOf(oldRoot).String() != air.IDOf(newRoot).String() {
				result.Consistent = false
			}
		}
	}
	return result, nil
}

// pathsOf collects the paths touched by the edits.
func pathsOf(edits []air.Edit) map[string]bool {
	out := make(map[string]bool, len(edits))
	for _, edit := range edits {
		out[edit.Path] = true
	}
	return out
}

// countRanges counts distinct edited byte intervals.
func countRanges(edits []air.Edit) int {
	seen := make(map[[2]int]struct{}, len(edits))
	for _, edit := range edits {
		seen[[2]int{edit.Offset, edit.EndOffset}] = struct{}{}
	}
	return len(seen)
}

// rootOf returns the AIR root registered under one path.
func rootOf(u *air.Unit, path string) (air.Node, bool) {
	if u == nil {
		return nil, false
	}
	for i, file := range u.Files {
		if file.Path == path && i < len(u.Roots) {
			return u.Roots[i], true
		}
	}
	return nil, false
}
