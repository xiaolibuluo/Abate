/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "slices"

// SourceUnit is one language-specific compilation request held entirely in
// memory; frontends must never read the real filesystem.
type SourceUnit struct {
	// Files is the set of source buffers to parse together.
	Files []SourceFile
	// Language is the lowercase language tag, for example "go".
	Language string
	// Dialect carries options such as "go1.26" or "python3.12".
	Dialect string
}

// Validate checks path uniqueness and non-empty source paths.
func (s SourceUnit) Validate() error {
	seen := make(map[string]struct{}, len(s.Files))
	for _, file := range s.Files {
		if file.Path == "" {
			return ErrMalformedUnit
		}
		if _, exists := seen[file.Path]; exists {
			return ErrDuplicatePath
		}
		seen[file.Path] = struct{}{}
	}
	return nil
}

// SortedFiles returns a path-ordered copy of the file list.
//
// Sorting here is what makes a frontend deterministic even when a caller
// supplies the same buffers in a different order.
func (s SourceUnit) SortedFiles() []SourceFile {
	out := slices.Clone(s.Files)
	slices.SortFunc(out, func(a, b SourceFile) int {
		if a.Path < b.Path {
			return -1
		}
		if a.Path > b.Path {
			return 1
		}
		return 0
	})
	return out
}

// Find returns the source buffer registered under path.
func (s SourceUnit) Find(path string) (SourceFile, bool) {
	for _, file := range s.Files {
		if file.Path == path {
			return file, true
		}
	}
	return SourceFile{}, false
}

// WithFile returns a copy of the unit with one source buffer replaced or added.
func (s SourceUnit) WithFile(file SourceFile) SourceUnit {
	out := SourceUnit{
		Files:    slices.Clone(s.Files),
		Language: s.Language,
		Dialect:  s.Dialect,
	}
	for i := range out.Files {
		if out.Files[i].Path == file.Path {
			out.Files[i] = file.Clone()
			return out
		}
	}
	out.Files = append(out.Files, file.Clone())
	return out
}
