/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// runConfig carries the process-level execution knobs of one backend. They
// are inputs to the constructor, never global state, so two backends over
// two workspaces can run concurrently without interfering.
type runConfig struct {
	// GoBin is the go toolchain binary used for every invocation.
	GoBin string
	// TestTimeout bounds one `go test -run` invocation in seconds.
	TestTimeout int
	// BuildCache is the GOCACHE directory shared across invocations.
	BuildCache string
}

// defaultConfig returns the production configuration: the go binary on PATH,
// a 120-second per-test bound and a private build cache inside the system
// temporary tree so the host cache is never polluted by mutant builds.
func defaultConfig() runConfig {
	return runConfig{
		GoBin:       "go",
		TestTimeout: 120,
		BuildCache:  filepath.Join(os.TempDir(), "abate-gomutate-gocache"),
	}
}

// goEnv builds the environment for one toolchain invocation. GOFLAGS is
// pinned to -mod=mod so a workspace without a populated vendor tree resolves
// exactly like the original tree, and GOTMPDIR/GOCACHE stay inside the
// private cache directory to keep every artifact of a run contained.
func (c runConfig) goEnv() []string {
	env := os.Environ()
	env = append(env,
		"GOFLAGS=-mod=mod",
		"GOCACHE="+c.BuildCache,
	)
	return dedupeEnv(env)
}

// dedupeEnv keeps the last occurrence of every KEY so the pinned values win
// over inherited ones regardless of the host environment.
func dedupeEnv(env []string) []string {
	index := make(map[string]int, len(env))
	for i, kv := range env {
		if j := strings.IndexByte(kv, '='); j > 0 {
			index[kv[:j]] = i
		}
	}
	keys := make([]string, 0, len(index))
	for k := range index {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	out := make([]string, 0, len(keys))
	for _, k := range keys {
		out = append(out, env[index[k]])
	}
	return out
}

// goRequiredFile reports whether a file must exist in a staged workspace
// for `go build` and `go test` to work: Go sources and the module
// definition files. Everything else (docs, scripts, archives, CI data) is
// irrelevant to the toolchain and copying it would only multiply the cost
// of staging one mutant, so it is skipped deliberately.
func goRequiredFile(name string) bool {
	if strings.HasSuffix(name, ".go") {
		return true
	}
	switch name {
	case "go.mod", "go.sum", "go.work", "go.work.sum":
		return true
	}
	return false
}

// copyWorkspace materializes the toolchain-visible subset of the workspace
// tree at src into dst: every Go source file and every module file, with
// the vendor and .git directories skipped. Permissions are preserved and
// the copy is fully synchronous, so the caller never observes a partially
// materialized tree.
func copyWorkspace(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspace, err)
	}
	walk := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			if d.Name() == "vendor" || d.Name() == ".git" {
				return filepath.SkipDir
			}
			return os.MkdirAll(filepath.Join(dst, rel), 0o755)
		}
		if !goRequiredFile(d.Name()) {
			return nil
		}
		return copyFile(path, filepath.Join(dst, rel), d)
	}
	if err := filepath.WalkDir(src, walk); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspace, err)
	}
	return nil
}

// copyFile copies one regular file preserving its permission bits. The
// destination always sits inside a freshly staged directory, so the file
// never exists yet and WriteFile applies the source permissions exactly.
func copyFile(src, dst string, d fs.DirEntry) error {
	info, err := d.Info()
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return nil
	}
	data, err := os.ReadFile(src) //nolint:gosec // the path is enumerated from the workspace tree.
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, info.Mode().Perm()) //nolint:gosec // dst is derived from the caller-supplied workspace root.
}

// rewriteToken returns the file content with exactly one token replaced by
// the mutated text. The token interval must fall inside the file, which the
// scanner guarantees because every token was discovered on this snapshot.
func rewriteToken(content []byte, tok Token, mutated string) ([]byte, error) {
	if tok.Offset < 0 || tok.End > len(content) || tok.Offset >= tok.End {
		return nil, fmt.Errorf("%w: token interval [%d,%d) outside %s (%d bytes)",
			ErrWorkspace, tok.Offset, tok.End, tok.Path, len(content))
	}
	if got := string(content[tok.Offset:tok.End]); got != tok.Text {
		return nil, fmt.Errorf("%w: token drift in %s: expected %q, found %q",
			ErrWorkspace, tok.Path, tok.Text, got)
	}
	out := make([]byte, 0, len(content)-(tok.End-tok.Offset)+len(mutated))
	out = append(out, content[:tok.Offset]...)
	out = append(out, mutated...)
	out = append(out, content[tok.End:]...)
	return out, nil
}

// runGoTest executes `go test` against dir with the given extra arguments
// and reports whether the run passed. A non-nil error means the invocation
// itself failed (toolchain missing, context canceled); a nil error with a
// false verdict means the tests ran and failed, which is the kill signal.
func runGoTest(ctx context.Context, cfg runConfig, dir string, args ...string) (bool, string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	full := append([]string{"test", "-count=1",
		fmt.Sprintf("-timeout=%ds", cfg.TestTimeout)}, args...)
	cmd := exec.CommandContext(ctx, cfg.GoBin, full...)
	cmd.Dir = dir
	cmd.Env = cfg.goEnv()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	runErr := cmd.Run()
	output := buf.String()
	if runErr == nil {
		return true, output, nil
	}
	var exitErr *exec.ExitError
	if errors.As(runErr, &exitErr) {
		return false, output, nil
	}
	return false, output, fmt.Errorf("%w: %v", ErrWorkspace, runErr)
}
