/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"abate/lang/air"
)

// repoGoFiles walks the Abate repository and returns every first-party Go file.
func repoGoFiles(t *testing.T) []string {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	root := filepath.Clean(filepath.Join(cwd, "..", ".."))
	paths := make([]string, 0, 256)
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			// Other packages run in parallel and create/remove temporary
			// directories inside the repo (e.g. ctl/cli plugin-state-*). If one
			// disappears between WalkDir's readdir and our visit, skip it rather
			// than failing the whole dogfood scan — this is a transient race, not
			// a missing first-party source file.
			if os.IsNotExist(walkErr) {
				return nil
			}
			return walkErr
		}
		if entry.IsDir() {
			switch {
			case entry.Name() == "vendor" || entry.Name() == "third_party" ||
				entry.Name() == "testdata" || entry.Name() == ".git" ||
				entry.Name() == "bench":
				return filepath.SkipDir
			case strings.HasPrefix(entry.Name(), "plugin-state-"):
				// Transient test scratch dir created by ctl/cli tests; not a
				// first-party source tree.
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".go") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("WalkDir: %v", err)
	}
	slices.Sort(paths)
	return paths
}

// readGoFile reads a repository source file for parsing.
func readGoFile(t *testing.T, path string) []byte {
	t.Helper()
	content, err := os.ReadFile(path) //nolint:gosec // the path comes from a repository walk.
	if err != nil {
		t.Fatalf("ReadFile %s: %v", path, err)
	}
	return content
}

// parseRepoFile parses one repository file with the given frontend.
func parseRepoFile(t *testing.T, frontend GoFrontend, path string,
	content []byte) *air.Unit {
	t.Helper()
	unit, diagnostics, err := frontend.Parse(air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile(path, content)},
	})
	if err != nil {
		t.Fatalf("Parse %s: %v", path, err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.IsError() {
			t.Fatalf("syntax error in %s: %s", path, diagnostic.String())
		}
	}
	if unit == nil || len(unit.Files) != 1 || unit.Roots[0] == nil {
		t.Fatalf("incomplete AIR for %s", path)
	}
	return unit
}

// TestDogfoodParsesAbateRepo parses every first-party Go file with zero errors.
func TestDogfoodParsesAbateRepo(t *testing.T) {
	paths := repoGoFiles(t)
	if len(paths) < 100 {
		t.Fatalf("only %d Go files discovered", len(paths))
	}
	syntax := NewGoFrontend().WithTypeCheck(false)
	totalNodes := 0
	for _, path := range paths {
		unit := parseRepoFile(t, syntax, path, readGoFile(t, path))
		totalNodes += unit.NodeCount()
	}
	t.Logf("dogfood: %d Go files parsed with zero syntax errors, %d AIR nodes",
		len(paths), totalNodes)
}

// TestDogfoodTypeCheckSample runs the in-memory go/types pass on real files.
func TestDogfoodTypeCheckSample(t *testing.T) {
	paths := repoGoFiles(t)
	frontend := NewGoFrontend()
	limit := min(len(paths), 30)
	for _, path := range paths[:limit] {
		parseRepoFile(t, frontend, path, readGoFile(t, path))
	}
	t.Logf("dogfood: in-memory type check completed on %d real files", limit)
}

// TestDogfoodTokenFidelitySample checks token equivalence on real files.
func TestDogfoodTokenFidelitySample(t *testing.T) {
	paths := repoGoFiles(t)
	syntax := NewGoFrontend().WithTypeCheck(false)
	checked := 0
	for _, path := range paths {
		if checked >= 100 {
			break
		}
		content := readGoFile(t, path)
		unit := parseRepoFile(t, syntax, path, content)
		if err := VerifyTokenFidelity(content, unit); err != nil {
			t.Fatalf("token fidelity on %s: %v", path, err)
		}
		checked++
	}
	if checked < 100 {
		t.Fatalf("only %d files available for the token sample", checked)
	}
	t.Logf("dogfood: token-level equivalence verified on %d real files", checked)
}

// TestDogfoodAIRRoundTrip checks that AIR from real files survives the wire.
func TestDogfoodAIRRoundTrip(t *testing.T) {
	paths := repoGoFiles(t)
	syntax := NewGoFrontend().WithTypeCheck(false)
	for _, path := range paths[:min(len(paths), 20)] {
		unit := parseRepoFile(t, syntax, path, readGoFile(t, path))
		archive, err := air.NewEncoder().Encode(unit)
		if err != nil {
			t.Fatalf("Encode %s: %v", path, err)
		}
		back, err := air.NewDecoder().Decode(archive)
		if err != nil {
			t.Fatalf("Decode %s: %v", path, err)
		}
		if err := air.CompareUnits(unit, back); err != nil {
			t.Fatalf("round trip %s: %v", path, err)
		}
	}
}
