/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package cmutate implements the real gcc execution backend for the
// mutation kernel (KRN-MUT). It shares the mutcore session engine and
// supplies only the C toolchain knowledge: frontkit-based lexical
// scanning, assert-macro test discovery, gcov-based coverage measurement
// and per-mutant judgement by compiling, linking and running the mutated
// tree with gcc.
package cmutate

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"abate/lang/frontkit"
	"abate/lang/mutcore"
)

// config carries the process-level execution knobs of the backend. They
// are inputs to the driver constructor, never global state, so two
// backends over two workspaces can run concurrently without interfering.
type config struct {
	// CcBin is the C compiler used for syntax checks, linking and
	// coverage instrumentation.
	CcBin string
	// GcovBin is the coverage decoder paired with the compiler.
	GcovBin string
	// TestTimeout bounds one test invocation in seconds.
	TestTimeout int
}

// DefaultConfig returns the production configuration: gcc and gcov
// resolved from PATH and a 60-second per-test bound.
func DefaultConfig() config {
	return config{CcBin: "gcc", GcovBin: "gcov", TestTimeout: 60}
}

// driver is the mutcore.Driver implementation for C workspaces.
type driver struct {
	cfg config
}

// New builds a real-workspace mutation session rooted at root. The driver
// probes gcc and gcov before any scanning starts, so a host without the
// toolchain fails fast with ErrToolchainMissing instead of a half-probed
// baseline.
func New(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg := DefaultConfig()
	for _, bin := range []string{cfg.CcBin, cfg.GcovBin} {
		if _, err := exec.LookPath(bin); err != nil {
			return nil, fmt.Errorf("%w: %s", mutcore.ErrToolchainMissing, bin)
		}
	}
	return mutcore.New(ctx, driver{cfg: cfg}, root)
}

// Language returns the canonical short tag of the backend.
func (d driver) Language() string { return "c" }

// NativeCompiler returns the two-level gating accelerator for C.
func (d driver) NativeCompiler() mutcore.NativeCompiler {
	return newNativeCompiler(d.cfg)
}

// ScanTokens discovers the mutable tokens of the workspace by lexing every
// non-test .c file through the shared C lexical tables, in deterministic
// (path, offset) order.
func (d driver) ScanTokens(root string) ([]mutcore.Token, error) {
	return mutcore.LexicalScan(root, map[string]bool{".c": true}, isTestFile,
		cLexSpec{}, mutcore.ClassifyLexical)
}

// isTestFile reports whether a workspace-relative path holds the assertion
// oracle rather than production code. Test files themselves never mutate:
// rewriting the oracle would make every verdict meaningless.
func isTestFile(rel string) bool {
	base := filepath.Base(rel)
	if strings.HasPrefix(base, "test-") || strings.HasPrefix(base, "test_") ||
		strings.HasSuffix(base, "_test.c") || strings.HasSuffix(base, ".test.c") {
		return true
	}
	dir := filepath.ToSlash(filepath.Dir(rel))
	return dir == "test" || dir == "tests" ||
		strings.HasPrefix(dir, "test/") || strings.HasPrefix(dir, "tests/")
}

// RequiredFile reports whether a file must exist in a staged workspace for
// gcc to compile, link and run the tree: C sources and their headers.
func (d driver) RequiredFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".c" || ext == ".h"
}

// sourceFiles enumerates the workspace .c files in deterministic
// lexicographic order of their workspace-relative slash paths, so a staged
// mutant links exactly the same translation units as the pristine tree.
func sourceFiles(dir string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".c") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

// Compile reports whether the staged mutant workspace still compiles,
// judged by gcc -fsyntax-only on the mutated translation unit. A mutant
// that fails the syntax check is killed by the toolchain itself, without
// spawning any doomed link or run.
func (d driver) Compile(ctx context.Context, dir string, tok mutcore.Token) bool {
	cmd := exec.CommandContext(ctx, d.cfg.CcBin, "-fsyntax-only",
		filepath.Join(dir, filepath.FromSlash(tok.Path)))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// DiscoverTests enumerates the executable test units of the workspace:
// one unit per .c file that declares at least one void test_* function
// behind the assert macro, in deterministic identifier order. The C
// harness is one main function that calls every test_*, so the file is
// the honest judging granularity for a framework-free target.
func (d driver) DiscoverTests(ctx context.Context, root string) ([]mutcore.TestRef, error) {
	detector := &testFileDetector{}
	_, err := mutcore.LexicalScan(root, map[string]bool{".c": true}, nil,
		cLexSpec{}, detector.classify)
	if err != nil && len(detector.files) == 0 {
		return nil, mutcore.ErrNoTests
	}
	if len(detector.files) == 0 {
		return nil, mutcore.ErrNoTests
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", mutcore.ErrInvalidRoot, err)
	}
	tests := make([]mutcore.TestRef, 0, len(detector.files))
	for _, file := range detector.files {
		tests = append(tests, mutcore.TestRef{
			ID:   file,
			Dir:  abs,
			Name: file,
		})
	}
	sort.Slice(tests, func(i, j int) bool { return tests[i].ID < tests[j].ID })
	return tests, nil
}

// testFileDetector is a lexical scan sink that records the relative path
// of every file declaring a void test_* function: a void keyword arming
// the name check, a test_* identifier completing it, and EOF markers
// resetting the per-file state.
type testFileDetector struct {
	files    []string
	sawVoid  bool
	reported bool
}

// classify implements the token-side detection. The path argument carries
// the workspace-relative slash path of the file owning the token, supplied
// by the scan driver at every call.
func (d *testFileDetector) classify(path string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		d.sawVoid, d.reported = false, false
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent,
		frontkit.TokenComment:
		return mutcore.ClassUnknown, false
	}
	if d.reported {
		return mutcore.ClassUnknown, false
	}
	if d.sawVoid {
		d.sawVoid = false
		if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, "test_") {
			d.reported = true
			d.files = append(d.files, path)
		}
		return mutcore.ClassUnknown, false
	}
	if lt.Text == "void" {
		d.sawVoid = true
	}
	return mutcore.ClassUnknown, false
}

// RunTest executes the staged mutant workspace by compiling and linking
// every .c source into one binary and running it. A nil error with a
// false verdict means the assertions failed or the program exited
// non-zero (kill); a non-nil error means the toolchain itself failed to
// build or start and the verdict must be treated as inconclusive, never
// as a fabricated kill.
func (d driver) RunTest(ctx context.Context, dir string, test mutcore.TestRef) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	buildDir, cleanup, err := d.buildTree(ctx, dir, false, "")
	if err != nil {
		return false, err
	}
	defer cleanup()
	passed, err := d.runBinary(ctx, buildDir, dir)
	if err != nil {
		return false, err
	}
	return passed, nil
}

// buildTree compiles every .c source of dir into its own object file and
// links them into one harness binary. The build always runs per-source
// (-c) so gcc names the coverage notes after the source (target.gcno),
// which gcov requires; a single-shot link would name them after the
// binary instead and leave gcov nothing to find. The object files land in
// buildDir: the staged tree itself for a plain run, or the caller-owned
// coverage directory for an instrumented run so the pristine tree never
// collects build artifacts. It returns the build directory holding the
// harness and a cleanup that releases any scratch space.
func (d driver) buildTree(ctx context.Context, dir string, coverage bool, covDir string) (string, func(), error) {
	sources, err := sourceFiles(dir)
	if err != nil {
		return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	if len(sources) == 0 {
		return "", func() {}, fmt.Errorf("%w: no .c sources in %s", mutcore.ErrWorkspaceFailed, dir)
	}
	buildDir := dir
	cleanup := func() {}
	if coverage {
		buildDir = covDir
	} else {
		tmp, err := os.MkdirTemp("", "abate-cmutate-build-")
		if err != nil {
			return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
		}
		buildDir = tmp
		cleanup = func() { _ = os.RemoveAll(tmp) }
	}
	objects := make([]string, 0, len(sources))
	for _, src := range sources {
		obj := filepath.Join(buildDir, strings.TrimSuffix(src, ".c")+".o")
		args := []string{"-O0", "-c", filepath.Join(dir, src), "-o", obj}
		if coverage {
			args = append([]string{"--coverage"}, args...)
		}
		cmd := exec.CommandContext(ctx, d.cfg.CcBin, args...)
		cmd.Dir = dir
		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		if err := cmd.Run(); err != nil {
			cleanup()
			return "", func() {}, fmt.Errorf("%w: compile %s: %s",
				mutcore.ErrWorkspaceFailed, src, tail(buf.String()))
		}
		objects = append(objects, obj)
	}
	binary := filepath.Join(buildDir, "abate-c-harness")
	linkArgs := []string{"-O0", "-o", binary}
	if coverage {
		linkArgs = append([]string{"--coverage"}, linkArgs...)
	}
	linkArgs = append(linkArgs, objects...)
	link := exec.CommandContext(ctx, d.cfg.CcBin, linkArgs...)
	link.Dir = dir
	var linkBuf bytes.Buffer
	link.Stdout = &linkBuf
	link.Stderr = &linkBuf
	if err := link.Run(); err != nil {
		cleanup()
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// A mutant that fails to link was never actually judged: the
			// syntax check precedes it, so reaching here means the tree
			// is structurally broken beyond a single-file edit. Treat the
			// runner as failed rather than fabricating a kill.
			return "", func() {}, fmt.Errorf("%w: link failed: %s",
				mutcore.ErrWorkspaceFailed, tail(linkBuf.String()))
		}
		return "", func() {}, fmt.Errorf("%w: link start: %v", mutcore.ErrWorkspaceFailed, err)
	}
	return buildDir, cleanup, nil
}

// runBinary executes the harness inside buildDir with the working
// directory pinned to the source tree. The boolean result is the program
// verdict: true for a clean zero exit, false for any non-zero exit,
// assertion abort or timeout. A non-nil error means the process itself
// could not start, which is inconclusive.
func (d driver) runBinary(ctx context.Context, buildDir, dir string) (bool, error) {
	run := exec.CommandContext(ctx, filepath.Join(buildDir, "abate-c-harness"))
	run.Dir = dir
	var runBuf bytes.Buffer
	run.Stdout = &runBuf
	run.Stderr = &runBuf
	runErr := run.Run()
	if runErr == nil {
		return true, nil
	}
	var exitErr *exec.ExitError
	if errors.As(runErr, &exitErr) {
		return false, nil
	}
	return false, fmt.Errorf("%w: run start: %v", mutcore.ErrWorkspaceFailed, runErr)
}

// tail truncates one toolchain output blob to the last 512 bytes so error
// values stay readable without flooding the report.
func tail(out string) string {
	const max = 512
	if len(out) <= max {
		return out
	}
	return out[len(out)-max:]
}
