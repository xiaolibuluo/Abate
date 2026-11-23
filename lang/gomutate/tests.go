/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// testRef is one discovered Go test: its fully qualified identifier in the
// "packagePath.Name" form that becomes the kernel-facing test key, plus the
// working directory the runner must enter to execute it with -run.
type testRef struct {
	// ID is the stable kernel-facing identifier, "pkg/path.TestName".
	ID string
	// Dir is the absolute directory of the package holding the test.
	Dir string
	// Name is the bare test function name accepted by -run.
	Name string
}

// packageDir pairs one package import path with its absolute directory, as
// reported by a single deterministic `go list -f` query.
type packageDir struct {
	path string
	dir  string
}

// discoverTests enumerates every TestXxx function below the user-selected
// scanRoot subtree through `go list` plus per-package `go test -list`
// invocations. Listing still runs at modRoot so module context stays intact,
// but packages outside scanRoot are dropped: a mutate session judges the
// tokens of the target subtree only, and pulling in unrelated packages would
// both waste coverage runs and let an already-red test elsewhere in the
// module abort an otherwise judgeable target. The fully qualified
// identifiers are sorted, so the mutation kernel always observes one
// deterministic test order for a fixed workspace snapshot.
func discoverTests(ctx context.Context, cfg runConfig, modRoot, scanRoot string) ([]testRef, error) {
	pkgs, err := listPackages(ctx, cfg, modRoot)
	if err != nil {
		return nil, err
	}
	tests := make([]testRef, 0, 16)
	for _, pkg := range pkgs {
		if !withinDir(scanRoot, pkg.dir) {
			continue
		}
		names, err := listPackageTests(ctx, cfg, pkg.dir)
		if err != nil {
			return nil, err
		}
		for _, name := range names {
			tests = append(tests, testRef{ID: pkg.path + "." + name, Dir: pkg.dir, Name: name})
		}
	}
	if len(tests) == 0 {
		return nil, ErrNoTests
	}
	slices.SortFunc(tests, func(a, b testRef) int {
		return strings.Compare(a.ID, b.ID)
	})
	return tests, nil
}

// withinDir reports whether dir equals root or lies inside it. The inputs
// are absolute directories reported by the toolchain, so a lexical check on
// the path separator boundary is exact.
func withinDir(root, dir string) bool {
	if dir == root {
		return true
	}
	return strings.HasPrefix(dir, root+string(filepath.Separator))
}

// listPackages runs `go list -f {{.ImportPath}}\t{{.Dir}} ./...` once inside
// root and parses the tab-separated rows. One invocation answers both the
// import path and the directory, so discovery never issues one subprocess
// per package and never guesses a directory from string arithmetic.
func listPackages(ctx context.Context, cfg runConfig, root string) ([]packageDir, error) {
	cmd := exec.CommandContext(ctx, cfg.GoBin, "list", "-f", "{{.ImportPath}}\t{{.Dir}}", "./...")
	cmd.Dir = root
	cmd.Env = cfg.goEnv()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%w: go list: %s: %v", ErrListTests, tail(buf.String()), err)
	}
	out := make([]packageDir, 0, 8)
	scan := bufio.NewScanner(&buf)
	for scan.Scan() {
		path, dir, ok := strings.Cut(scan.Text(), "\t")
		if ok && path != "" && dir != "" {
			out = append(out, packageDir{path: path, dir: dir})
		}
	}
	return out, nil
}

// listPackageTests runs `go test -list ^Test` inside one package directory
// and returns every matched top-level test name in output order. The -list
// flag only compiles the test binary and prints matches, so no test body
// executes during discovery.
func listPackageTests(ctx context.Context, cfg runConfig, dir string) ([]string, error) {
	cmd := exec.CommandContext(ctx, cfg.GoBin, "test", "-count=1", "-list", "^Test")
	cmd.Dir = dir
	cmd.Env = cfg.goEnv()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%w: go test -list in %s: %s: %v", ErrListTests, dir, tail(buf.String()), err)
	}
	out := make([]string, 0, 8)
	scan := bufio.NewScanner(&buf)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if strings.HasPrefix(line, "Test") {
			out = append(out, line)
		}
	}
	return out, nil
}

// modulePath resolves the module path of the workspace through `go list -m`,
// which coverage profiles use as the file name prefix.
func modulePath(ctx context.Context, cfg runConfig, root string) (string, error) {
	cmd := exec.CommandContext(ctx, cfg.GoBin, "list", "-m")
	cmd.Dir = root
	cmd.Env = cfg.goEnv()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w: go list -m: %s: %v", ErrListTests, tail(buf.String()), err)
	}
	path := strings.TrimSpace(buf.String())
	if path == "" {
		return "", fmt.Errorf("%w: go list -m returned an empty module path", ErrListTests)
	}
	return path, nil
}

// tail truncates toolchain output so an error message stays readable while
// still carrying the failing compiler or test runner diagnostics.
func tail(output string) string {
	text := strings.TrimSpace(output)
	if len(text) > 512 {
		return text[:512]
	}
	return text
}
