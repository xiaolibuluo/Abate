/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"abate/lang/air"
)

// errUnknownEditPath reports an incremental edit aimed at an absent file.
var errUnknownEditPath = errors.New("frontkit: edit path not in source unit")

// ParseIncremental reparses after applying edits, reusing the AIR subtrees of
// files whose content hash is unchanged. This is honest file-level subtree
// reuse: no node-level incremental reduction inside a changed file is claimed.
// FullRebuild is set whenever nothing could be reused, and the unit annotations
// state the reuse mode explicitly, so a full rebuild is never presented as a
// node-incremental parse. The Diff is computed from identifier multisets.
func (f *FrontendBase) ParseIncremental(unit air.SourceUnit, old *air.Unit,
	edits []air.Edit) (*air.Unit, air.Diff, error) {
	if f == nil || f.spec == nil {
		return nil, air.Diff{}, fmt.Errorf("%w: frontend", ErrNilSpec)
	}
	if err := unit.Validate(); err != nil {
		return nil, air.Diff{}, err
	}
	applied, err := f.applyEdits(unit, edits)
	if err != nil {
		return nil, air.Diff{}, err
	}
	tag, ok := CanonicalTag(f.spec.Tag())
	if !ok {
		return nil, air.Diff{}, fmt.Errorf("%w: %q",
			ErrUnsupportedLanguage, f.spec.Tag())
	}
	dialect, err := f.spec.NormalizeDialect(unit.Dialect)
	if err != nil {
		return nil, air.Diff{}, fmt.Errorf("%w: %w", ErrBadDialect, err)
	}
	reusedRoots, reusedFiles := reusableRoots(applied, old,
		oldDialect(old, tag), dialect)
	changed := air.SourceUnit{Language: unit.Language, Dialect: unit.Dialect}
	for _, src := range applied.SortedFiles() {
		if _, exists := reusedRoots[src.Path]; !exists {
			changed.Files = append(changed.Files, src)
		}
	}
	fresh, _, err := f.Parse(changed)
	if err != nil {
		return nil, air.Diff{}, err
	}
	name := fresh.Package.Name
	imports := slices.Clone(fresh.Package.Imports)
	if old != nil {
		if name == "" {
			name = old.Package.Name
		}
		imports = append(slices.Clone(old.Package.Imports), imports...)
	}
	out := air.NewUnit(air.NewPackageRef(name, imports))
	for i := range fresh.Files {
		if err := out.AddFile(fresh.Files[i], fresh.Roots[i]); err != nil {
			return nil, air.Diff{}, err
		}
	}
	paths := make([]string, 0, len(reusedRoots))
	for path := range reusedRoots {
		paths = append(paths, path)
	}
	slices.Sort(paths)
	for _, path := range paths {
		if err := out.AddFile(reusedFiles[path], reusedRoots[path]); err != nil {
			return nil, air.Diff{}, err
		}
	}
	for _, key := range fresh.AnnotationKeys() {
		out.Annotations[key] = fresh.Annotations[key]
	}
	prefix := "lang." + tag + "."
	out.Annotations[prefix+"files"] = strconv.Itoa(len(out.Files))
	out.Annotations[prefix+"incremental"] = reuseMode(old, len(reusedRoots))
	out.Annotations[prefix+"reusedFiles"] = strconv.Itoa(len(reusedRoots))
	diff := air.ComputeDiff(old, out)
	diff.FullRebuild = old == nil || len(reusedRoots) == 0
	return out, diff, nil
}

// applyEdits groups edits by path in sorted order and applies each group once.
func (f *FrontendBase) applyEdits(unit air.SourceUnit,
	edits []air.Edit) (air.SourceUnit, error) {
	if f == nil || f.spec == nil {
		return air.SourceUnit{}, fmt.Errorf("%w: frontend", ErrNilSpec)
	}
	grouped := make(map[string][]air.Edit, len(edits))
	for _, edit := range edits {
		grouped[edit.Path] = append(grouped[edit.Path], edit)
	}
	paths := make([]string, 0, len(grouped))
	for path := range grouped {
		paths = append(paths, path)
	}
	slices.Sort(paths)
	applied := unit
	for _, path := range paths {
		src, ok := applied.Find(path)
		if !ok {
			return air.SourceUnit{}, fmt.Errorf("%w: %q", errUnknownEditPath, path)
		}
		next, err := air.ApplyEdits(src.Content, grouped[path])
		if err != nil {
			return air.SourceUnit{}, err
		}
		applied = applied.WithFile(air.NewSourceFile(path, next))
	}
	return applied, nil
}

// reusableRoots pairs unchanged content hashes with their prior root and file.
//
// The reuse is skipped when the dialect changed, and also when the previous
// unit does not record a dialect at all. A dialect switch can change how the
// same bytes are reduced, so reusing a tree that was produced under another
// dialect would report a stale result without any visible sign.
func reusableRoots(applied air.SourceUnit, old *air.Unit,
	oldDialect, dialect string) (map[string]air.Node, map[string]air.FileRef) {
	roots := make(map[string]air.Node)
	files := make(map[string]air.FileRef)
	if old == nil {
		return roots, files
	}
	if oldDialect == "" || oldDialect != dialect {
		return roots, files
	}
	for i, file := range old.Files {
		if i >= len(old.Roots) {
			continue
		}
		src, ok := applied.Find(file.Path)
		if !ok || src.Digest() != file.ContentHash {
			continue
		}
		roots[file.Path] = old.Roots[i]
		files[file.Path] = file
	}
	return roots, files
}

// oldDialect returns the dialect the previous unit was parsed under.
//
// The value comes from the unit annotation that every parse writes, so the
// comparison is made between two normalized dialect spellings.
func oldDialect(old *air.Unit, tag string) string {
	if old == nil {
		return ""
	}
	return old.Annotations["lang."+tag+".dialect"]
}

// reuseMode renders the honest reuse label stored in unit annotations.
func reuseMode(old *air.Unit, reused int) string {
	if old == nil || reused == 0 {
		return "full-rebuild"
	}
	return "file-level-reuse"
}
