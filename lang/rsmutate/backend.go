/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package rsmutate implements the real rustc execution backend for the
// mutation kernel (KRN-MUT). It shares the mutcore session engine and
// supplies only the Rust toolchain knowledge: frontkit-based lexical
// scanning, assert-macro test discovery, source-based coverage through
// rustc -Cinstrument-coverage plus llvm-profdata/llvm-cov, and per-mutant
// judgement by compiling and running the mutated tree with rustc.
package rsmutate

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
	// RustcBin is the rustc executable used for compilation and runs.
	RustcBin string
	// ProfdataBin merges raw coverage profiles into an indexed one.
	ProfdataBin string
	// CovBin exports the indexed profile as a JSON line-coverage report.
	CovBin string
	// TestTimeout bounds one test invocation in seconds.
	TestTimeout int
}

// driver is the mutcore.Driver implementation for Rust workspaces.
type driver struct {
	cfg config
}

// New builds a real-workspace mutation session rooted at root. The driver
// probes rustc and the paired llvm coverage tools before any scanning
// starts, so a host without the toolchain fails fast with
// ErrToolchainMissing instead of a half-probed baseline. The llvm tools
// ship inside the rustup toolchain tree rather than on PATH, so they are
// resolved against rustc --print target-libdir.
func New(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	return mutcore.New(ctx, driver{cfg: cfg}, root)
}

// defaultConfig resolves rustc from PATH and the llvm coverage tools from
// the rustup toolchain bin directory that rustc reports. Every missing
// piece aborts with ErrToolchainMissing so the backend never starts a
// session it cannot measure or judge.
func defaultConfig() (config, error) {
	rustc, err := exec.LookPath("rustc")
	if err != nil {
		return config{}, fmt.Errorf("%w: rustc", mutcore.ErrToolchainMissing)
	}
	libdir, err := rustTargetLibdir(rustc)
	if err != nil {
		return config{}, err
	}
	binDir := filepath.Join(filepath.Dir(libdir), "bin")
	cfg := config{
		RustcBin:    rustc,
		ProfdataBin: filepath.Join(binDir, "llvm-profdata"),
		CovBin:      filepath.Join(binDir, "llvm-cov"),
		TestTimeout: 60,
	}
	for _, tool := range []struct {
		name string
		path string
	}{
		{"llvm-profdata", cfg.ProfdataBin},
		{"llvm-cov", cfg.CovBin},
	} {
		if info, statErr := os.Stat(tool.path); statErr != nil || info.IsDir() {
			return config{}, fmt.Errorf("%w: %s (install the llvm-tools rustup component)",
				mutcore.ErrToolchainMissing, tool.name)
		}
	}
	return cfg, nil
}

// rustTargetLibdir asks rustc for the target libdir of its own toolchain,
// which anchors the sibling bin directory holding the llvm tools.
func rustTargetLibdir(rustc string) (string, error) {
	out, err := exec.Command(rustc, "--print", "target-libdir").Output()
	if err != nil {
		return "", fmt.Errorf("%w: rustc --print target-libdir: %v",
			mutcore.ErrToolchainMissing, err)
	}
	libdir := strings.TrimSpace(string(out))
	if libdir == "" {
		return "", fmt.Errorf("%w: empty target-libdir", mutcore.ErrToolchainMissing)
	}
	return libdir, nil
}

// Language returns the canonical short tag of the backend.
func (d driver) Language() string { return "rs" }

// NativeCompiler returns the two-level gating accelerator for Rust.
func (d driver) NativeCompiler() mutcore.NativeCompiler {
	return newNativeCompiler(d.cfg)
}

// rsClassState enumerates the states of the Rust oracle-exclusion machine.
// The machine tracks fn keyword → test_* identifier → brace-delimited body
// so tokens inside a test_* function block are never classified as mutable.
type rsClassState int

const (
	rsNormal rsClassState = iota
	rsSawFn
	rsSawTestFn
	rsSkipping
)

// rsClassifier is the stateful per-stream oracle-exclusion classifier for
// Rust. Because the Rust lexer is not indent-sensitive, the classifier
// tracks brace depth to delimit test_* function bodies.
type rsClassifier struct {
	state rsClassState
	depth int
}

// classify inspects one token in stream order. Inside a test_* function
// body the classifier returns false for every token; outside it delegates
// to the shared lexical classification sets. The path argument is accepted
// for the callback signature but never used by the classifier logic.
func (c *rsClassifier) classify(_ string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch c.state {
	case rsNormal:
		return c.classifyNormal(lt)
	case rsSawFn:
		return c.classifyFnName(lt)
	case rsSawTestFn:
		return c.classifyTestHeader(lt)
	default:
		return c.classifySkipping(lt)
	}
}

func (c *rsClassifier) classifyNormal(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		c.state = rsNormal
		c.depth = 0
		return mutcore.ClassUnknown, false
	}
	if lt.Kind == frontkit.TokenKeyword && lt.Text == "fn" {
		c.state = rsSawFn
		return mutcore.ClassUnknown, false
	}
	return mutcore.ClassifyLexical("", lt)
}

func (c *rsClassifier) classifyFnName(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		c.state = rsNormal
		c.depth = 0
		return mutcore.ClassUnknown, false
	}
	if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, "test_") {
		c.state = rsSawTestFn
		return mutcore.ClassUnknown, false
	}
	c.state = rsNormal
	return mutcore.ClassifyLexical("", lt)
}

func (c *rsClassifier) classifyTestHeader(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		c.state = rsNormal
		c.depth = 0
	case frontkit.TokenPunct:
		if lt.Text == "{" {
			c.state = rsSkipping
			c.depth = 1
		}
	}
	return mutcore.ClassUnknown, false
}

func (c *rsClassifier) classifySkipping(lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		c.state = rsNormal
		c.depth = 0
	case frontkit.TokenPunct:
		if lt.Text == "{" {
			c.depth++
		} else if lt.Text == "}" {
			c.depth--
			if c.depth <= 0 {
				c.state = rsNormal
			}
		}
	}
	return mutcore.ClassUnknown, false
}

// ScanTokens discovers the mutable tokens of the workspace by lexing every
// non-test .rs file through the shared Rust lexical tables. A per-token
// classifier skips tokens inside fn test_* bodies so source files that
// declare both production code and test harness in one crate root still
// expose every mutable token outside the oracle blocks.
func (d driver) ScanTokens(root string) ([]mutcore.Token, error) {
	cl := &rsClassifier{}
	// Pass nil for the file-level skip: the per-token rsClassifier
	// handles hybrid single-file crates where production code and
	// test harness coexist in one .rs file. Dedicated test-only
	// directories produce zero mutable tokens, which the session
	// layer handles through ErrNoSourceFiles.
	return mutcore.LexicalScan(root, map[string]bool{".rs": true}, nil,
		rsLexSpec{}, cl.classify)
}

// isTestFile reports whether a workspace-relative path is a dedicated test
// file whose every token belongs to the assertion oracle. The per-token
// classifier handles hybrid files (one .rs serving as both crate root and
// test harness); this file-level filter catches standard multi-file Cargo
// project layouts like tests/ and *_test.rs.
func isTestFile(rel string) bool {
	base := filepath.Base(rel)
	if strings.HasPrefix(base, "test_") || strings.HasSuffix(base, "_test.rs") ||
		strings.HasSuffix(base, ".test.rs") {
		return true
	}
	dir := filepath.ToSlash(filepath.Dir(rel))
	return dir == "test" || dir == "tests" ||
		strings.HasPrefix(dir, "test/") || strings.HasPrefix(dir, "tests/")
}

// RequiredFile reports whether a file must exist in a staged workspace for
// rustc to compile and run the tree: Rust sources plus the Cargo manifest.
func (d driver) RequiredFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".rs" || name == "Cargo.toml"
}

// sourceFiles enumerates the workspace .rs files in deterministic
// lexicographic order of their workspace-relative slash paths, so a staged
// mutant compiles exactly the same crate root as the pristine tree.
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
		if strings.EqualFold(filepath.Ext(entry.Name()), ".rs") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

// Compile reports whether the staged mutant workspace still compiles,
// judged by rustc --emit=metadata on the mutated crate root. A mutant
// that fails the metadata check is killed by the toolchain itself,
// without spawning any doomed link or run. --emit=metadata performs the
// full parse, name resolution and type check without producing object
// code, which is the narrowest honest build Rust offers.
func (d driver) Compile(ctx context.Context, dir string, tok mutcore.Token) bool {
	cmd := exec.CommandContext(ctx, d.cfg.RustcBin, "--emit=metadata", "-o", os.DevNull,
		filepath.Join(dir, filepath.FromSlash(tok.Path)))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// DiscoverTests enumerates the executable test units of the workspace:
// one unit per .rs file that declares at least one fn test_* function
// behind the assert macros, in deterministic identifier order. The Rust
// harness is one fn main that calls every test_*, so the file is the
// honest judging granularity for a framework-free target.
func (d driver) DiscoverTests(ctx context.Context, root string) ([]mutcore.TestRef, error) {
	detector := &testFileDetector{}
	_, err := mutcore.LexicalScan(root, map[string]bool{".rs": true}, nil,
		rsLexSpec{}, detector.classify)
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
// of every file declaring an fn test_* function: an fn keyword arming the
// name check, a test_* identifier completing it, and EOF markers
// resetting the per-file state.
type testFileDetector struct {
	files    []string
	sawFn    bool
	reported bool
}

// classify implements the token-side detection. The path argument carries
// the workspace-relative slash path of the file owning the token, supplied
// by the scan driver at every call.
func (d *testFileDetector) classify(path string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		d.sawFn, d.reported = false, false
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent,
		frontkit.TokenComment:
		return mutcore.ClassUnknown, false
	}
	if d.reported {
		return mutcore.ClassUnknown, false
	}
	if d.sawFn {
		d.sawFn = false
		if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, "test_") {
			d.reported = true
			d.files = append(d.files, path)
		}
		return mutcore.ClassUnknown, false
	}
	if lt.Text == "fn" {
		d.sawFn = true
	}
	return mutcore.ClassUnknown, false
}

// RunTest executes the staged mutant workspace by compiling its crate root
// into one binary and running it. A nil error with a false verdict means
// the assertions failed or the program exited non-zero (kill); a non-nil
// error means the toolchain itself failed to build or start and the
// verdict must be treated as inconclusive, never as a fabricated kill.
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
	passed, err := d.runBinary(ctx, filepath.Join(buildDir, "abate-rs-harness"), dir)
	if err != nil {
		return false, err
	}
	return passed, nil
}

// buildTree compiles the crate root of dir into one harness binary inside
// a scratch build directory so the staged tree never collects build
// artifacts. When coverage is enabled the binary carries rustc
// -Cinstrument-coverage instrumentation and the run emits a raw profile
// for llvm-profdata/llvm-cov to decode; the coverage directory is then
// supplied by the caller. It returns the build directory holding the
// harness and a cleanup that releases the scratch space.
func (d driver) buildTree(ctx context.Context, dir string, coverage bool, covDir string) (string, func(), error) {
	sources, err := sourceFiles(dir)
	if err != nil {
		return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	if len(sources) == 0 {
		return "", func() {}, fmt.Errorf("%w: no .rs sources in %s", mutcore.ErrWorkspaceFailed, dir)
	}
	buildDir := covDir
	if !coverage {
		tmp, mkErr := os.MkdirTemp("", "abate-rsmutate-build-")
		if mkErr != nil {
			return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, mkErr)
		}
		buildDir = tmp
	}
	cleanup := func() {}
	if !coverage {
		cleanup = func() { _ = os.RemoveAll(buildDir) }
	}
	binary := filepath.Join(buildDir, "abate-rs-harness")
	args := []string{"-O", "-o", binary}
	if coverage {
		args = append([]string{"-Cinstrument-coverage"}, args...)
	}
	// The generated target is a single-file crate; the crate root is the
	// lexicographically first .rs source, which rustc compiles as the
	// whole program.
	args = append(args, filepath.Join(dir, sources[0]))
	cmd := exec.CommandContext(ctx, d.cfg.RustcBin, args...)
	cmd.Dir = dir
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		cleanup()
		return "", func() {}, fmt.Errorf("%w: rustc: %s",
			mutcore.ErrWorkspaceFailed, tail(buf.String()))
	}
	return buildDir, cleanup, nil
}

// runBinary executes the harness with the working directory pinned to the
// source tree. The boolean result is the program verdict: true for a
// clean zero exit, false for any non-zero exit, assertion abort or
// timeout. A non-nil error means the process itself could not start,
// which is inconclusive. When profRaw is non-empty the run records its
// raw coverage profile at that path for llvm-profdata to merge.
func (d driver) runBinary(ctx context.Context, binary, dir string, profRaw ...string) (bool, error) {
	run := exec.CommandContext(ctx, binary)
	run.Dir = dir
	if len(profRaw) > 0 && profRaw[0] != "" {
		run.Env = append(os.Environ(), "LLVM_PROFILE_FILE="+profRaw[0])
	}
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
