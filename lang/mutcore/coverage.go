/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import "sort"

// CoverMatrix records which source lines each discovered test executed on
// the pristine workspace. Language drivers populate it from their
// toolchain-specific coverage format; the session core consumes it for
// coverage-guided pruning and suite fingerprinting.
type CoverMatrix struct {
	byTest map[string]map[string]map[int]bool
}

// NewCoverMatrix returns an empty coverage matrix.
func NewCoverMatrix() *CoverMatrix {
	return &CoverMatrix{byTest: make(map[string]map[string]map[int]bool)}
}

// Add records that the given test executed line of file. File paths are
// stored in workspace-relative slash form so fingerprints stay stable
// across hosts.
func (m *CoverMatrix) Add(testID, file string, line int) {
	files, ok := m.byTest[testID]
	if !ok {
		files = make(map[string]map[int]bool)
		m.byTest[testID] = files
	}
	lines, ok := files[file]
	if !ok {
		lines = make(map[int]bool)
		files[file] = lines
	}
	lines[line] = true
}

// TestCovers reports whether the given test executed any line of the
// given file in its own measured run.
func (m *CoverMatrix) TestCovers(testID, file string, line int) bool {
	lines, ok := m.byTest[testID][file]
	if !ok {
		return false
	}
	return lines[line]
}

// Covering returns how many discovered tests executed the given line of
// the given file.
func (m *CoverMatrix) Covering(file string, line int) uint64 {
	var n uint64
	for _, files := range m.byTest {
		if lines, ok := files[file]; ok && lines[line] {
			n++
		}
	}
	return n
}

// Files returns the sorted set of files any test covered, used for the
// deterministic suite fingerprint.
func (m *CoverMatrix) Files() []string {
	set := make(map[string]bool)
	for _, files := range m.byTest {
		for file := range files {
			set[file] = true
		}
	}
	out := make([]string, 0, len(set))
	for file := range set {
		out = append(out, file)
	}
	sort.Strings(out)
	return out
}
