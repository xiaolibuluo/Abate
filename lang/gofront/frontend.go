/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"

	"abate/lang/air"
)

// GoFrontend implements air.Frontend for the Go language.
//
// Parsing is deterministic: files are handled in path order, diagnostics carry
// exact byte positions, and the same SourceUnit always produces the same AIR.
// The zero value is not usable; construct one with NewGoFrontend.
type GoFrontend struct {
	maxDiagnostics int
	typeCheck      bool
}

// NewGoFrontend returns the default Go frontend with semantic annotation on.
func NewGoFrontend() GoFrontend {
	return GoFrontend{maxDiagnostics: air.MaxDiagnostics, typeCheck: true}
}

// WithTypeCheck enables or disables the in-memory go/types annotation pass.
func (f GoFrontend) WithTypeCheck(enabled bool) GoFrontend {
	f.typeCheck = enabled
	return f
}

// WithMaxDiagnostics sets the diagnostic cap, clamped to a positive value.
func (f GoFrontend) WithMaxDiagnostics(limit int) GoFrontend {
	if limit > 0 {
		f.maxDiagnostics = limit
	}
	return f
}

// compile-time proof that GoFrontend satisfies the seven-language contract.
var _ air.Frontend = GoFrontend{}

// parsedFile bundles one file's syntax, view, and AIR root.
type parsedFile struct {
	file   air.SourceFile
	view   *GoView
	root   air.Node
	parsed *ast.File
	fset   *token.FileSet
}

// Parse performs a full parse of every source buffer in a unit.
func (f GoFrontend) Parse(unit air.SourceUnit) (*air.Unit, []air.Diagnostic, error) {
	if err := unit.Validate(); err != nil {
		return nil, nil, err
	}
	if unit.Language != "" && unit.Language != "go" {
		return nil, nil, fmt.Errorf("%w: %q", ErrUnsupportedLanguage, unit.Language)
	}

	limit := f.maxDiagnostics
	if limit <= 0 {
		limit = air.MaxDiagnostics
	}
	diagnostics := make([]air.Diagnostic, 0)
	files := unit.SortedFiles()
	parsedFiles := make([]parsedFile, 0, len(files))
	packageName := ""
	imports := make([]string, 0)

	for _, src := range files {
		fset := token.NewFileSet()
		parsed, parseErr := parser.ParseFile(fset, src.Path, src.Content,
			parser.ParseComments|parser.AllErrors|parser.SkipObjectResolution)
		diagnostics = append(diagnostics,
			collectDiagnostics(src.Path, parseErr, limit-len(diagnostics))...)
		if len(diagnostics) > limit {
			diagnostics = diagnostics[:limit]
		}
		view := newGoView(parsed, fset, src.Content)
		root := air.Node(nil)
		if parsed == nil {
			root = view.synthetic(air.KindPackage, "")
		} else {
			root = mapFile(view)
			if packageName == "" {
				packageName = view.PackageName()
			}
			imports = append(imports, view.Imports()...)
			if f.typeCheck {
				annotateTypes(view, fset, []*ast.File{parsed}, view.PackageName())
			}
		}
		parsedFiles = append(parsedFiles, parsedFile{
			file: src, view: view, root: root, parsed: parsed, fset: fset,
		})
		if len(diagnostics) >= limit {
			break
		}
	}

	out := air.NewUnit(air.NewPackageRef(packageName, imports))
	for _, entry := range parsedFiles {
		if err := out.AddFile(air.NewFileRef(entry.file.Path, entry.file.Content),
			entry.root); err != nil {
			return nil, diagnostics, err
		}
	}
	return out, diagnostics, nil
}

// ParseIncremental reparses only the files changed by edits.
//
// Unchanged files are detected by comparing their new content hash with the
// hash recorded in old.Files, and their AIR subtrees and FileRefs are reused
// verbatim instead of being parsed again. Changed files go through Parse, and
// a real node-level Diff is computed from identifier multisets. This is
// file-level subtree reuse; node-level reuse inside a changed file is an M5
// enhancement. FullRebuild is set when nothing could be reused, so the
// semantics are reported honestly rather than claimed as node-incremental.
func (f GoFrontend) ParseIncremental(unit air.SourceUnit, old *air.Unit,
	edits []air.Edit) (*air.Unit, air.Diff, error) {
	if err := unit.Validate(); err != nil {
		return nil, air.Diff{}, err
	}
	applied, err := applyIncrementalEdits(unit, edits)
	if err != nil {
		return nil, air.Diff{}, err
	}
	reusedRoots, reusedFiles := collectReusableFiles(old, applied)
	changed := assembleChangedUnit(unit, applied, reusedRoots)
	fresh, diagnostics, err := f.Parse(changed)
	if err != nil {
		return nil, air.Diff{}, err
	}

	name, imports := mergePackageMetadata(old, fresh)
	out := air.NewUnit(air.NewPackageRef(name, imports))
	for i, file := range fresh.Files {
		if err := out.AddFile(file, fresh.Roots[i]); err != nil {
			return nil, air.Diff{}, err
		}
	}
	for _, src := range applied.SortedFiles() {
		root, ok := reusedRoots[src.Path]
		if !ok {
			continue
		}
		if err := out.AddFile(reusedFiles[src.Path], root); err != nil {
			return nil, air.Diff{}, err
		}
	}
	annotateDiagnostics(out, diagnostics)
	diff := air.ComputeDiff(old, out)
	diff.FullRebuild = old == nil || len(reusedRoots) == 0
	return out, diff, nil
}

// collectDiagnostics normalizes a go/parser failure into structured findings.
func collectDiagnostics(path string, parseErr error, limit int) []air.Diagnostic {
	if parseErr == nil || limit <= 0 {
		return nil
	}
	var list scanner.ErrorList
	if !errors.As(parseErr, &list) {
		return []air.Diagnostic{{
			Path: path, Severity: air.SeverityError,
			Code: "LANG-GO-E-001", Message: parseErr.Error(),
		}}
	}
	out := make([]air.Diagnostic, 0, min(len(list), limit))
	for i, item := range list {
		if i >= limit {
			break
		}
		out = append(out, air.Diagnostic{
			Path: path, Line: item.Pos.Line, Column: item.Pos.Column,
			Severity: air.SeverityError,
			Code:     fmt.Sprintf("LANG-GO-E-%03d", i+1),
			Message:  item.Msg,
		})
	}
	return out
}
