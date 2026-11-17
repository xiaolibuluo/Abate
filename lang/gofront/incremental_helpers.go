/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"fmt"
	"slices"
	"strconv"

	"abate/lang/air"
)

// Applies edits grouped by path and returns the resulting source unit.
func applyIncrementalEdits(unit air.SourceUnit,
	edits []air.Edit) (air.SourceUnit, error) {
	byPath := make(map[string][]air.Edit)
	for _, edit := range edits {
		byPath[edit.Path] = append(byPath[edit.Path], edit)
	}
	applied := unit
	for _, src := range unit.SortedFiles() {
		group := byPath[src.Path]
		if len(group) == 0 {
			continue
		}
		next, err := air.ApplyEdits(src.Content, group)
		if err != nil {
			return air.SourceUnit{}, err
		}
		applied = applied.WithFile(air.NewSourceFile(src.Path, next))
	}
	editPaths := make([]string, 0, len(byPath))
	for path := range byPath {
		editPaths = append(editPaths, path)
	}
	slices.Sort(editPaths)
	for _, path := range editPaths {
		if _, ok := applied.Find(path); !ok {
			return air.SourceUnit{}, fmt.Errorf("%w: %q", ErrUnknownEditPath, path)
		}
	}
	return applied, nil
}

// Collects roots and file references that may be reused from the old unit.
func collectReusableFiles(old *air.Unit,
	applied air.SourceUnit) (map[string]air.Node, map[string]air.FileRef) {
	reusedRoots := make(map[string]air.Node)
	reusedFiles := make(map[string]air.FileRef)
	if old != nil {
		for i, file := range old.Files {
			src, ok := applied.Find(file.Path)
			if !ok || src.Digest() != file.ContentHash || i >= len(old.Roots) {
				continue
			}
			reusedRoots[file.Path] = old.Roots[i]
			reusedFiles[file.Path] = file
		}
	}
	return reusedRoots, reusedFiles
}

// Builds the unit of files that must be reparsed after subtree reuse.
func assembleChangedUnit(unit air.SourceUnit, applied air.SourceUnit,
	reusedRoots map[string]air.Node) air.SourceUnit {
	changed := air.SourceUnit{Language: unit.Language, Dialect: unit.Dialect}
	for _, src := range applied.SortedFiles() {
		if _, ok := reusedRoots[src.Path]; !ok {
			changed.Files = append(changed.Files, src)
		}
	}
	return changed
}

// Merges fresh package metadata with the metadata of the previous unit.
func mergePackageMetadata(old *air.Unit, fresh *air.Unit) (string, []string) {
	name, imports := fresh.Package.Name, fresh.Package.Imports
	if old != nil {
		if name == "" {
			name = old.Package.Name
		}
		imports = append(append([]string(nil), old.Package.Imports...), imports...)
	}
	return name, imports
}

// Records diagnostic count and the first finding as unit annotations.
func annotateDiagnostics(out *air.Unit, diagnostics []air.Diagnostic) {
	if len(diagnostics) > 0 {
		out.Annotations["lang.go.diagnostics"] = strconv.Itoa(len(diagnostics))
		out.Annotations["lang.go.firstDiagnostic"] = diagnostics[0].String()
	}
}
