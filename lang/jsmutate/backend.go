/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package jsmutate implements the real node:test execution backend for
// the mutation kernel (KRN-MUT). It shares the mutcore session engine and
// supplies only the JavaScript toolchain knowledge: frontkit-based
// lexical scanning, per-file test discovery and per-mutant judgement by
// executing the mutated tree through node --test.
package jsmutate

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	// NodeBin is the node executable used for every invocation.
	NodeBin string
	// TestTimeout bounds one test invocation in seconds.
	TestTimeout int
}

// DefaultConfig returns the production configuration: node resolved from
// PATH and a 60-second per-test bound.
func DefaultConfig() config {
	return config{NodeBin: "node", TestTimeout: 60}
}

// driver is the mutcore.Driver implementation for JavaScript workspaces.
type driver struct {
	cfg config
}

// New builds a real-workspace mutation session rooted at root. The driver
// probes node before any scanning starts, so a host without the runtime
// fails fast with ErrToolchainMissing instead of a half-probed baseline.
func New(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg := DefaultConfig()
	if _, err := exec.LookPath(cfg.NodeBin); err != nil {
		return nil, fmt.Errorf("%w: node", mutcore.ErrToolchainMissing)
	}
	return mutcore.New(ctx, driver{cfg: cfg}, root)
}

// Language returns the canonical short tag of the backend.
func (d driver) Language() string { return "js" }

// NativeCompiler returns the two-level gating accelerator for JavaScript.
func (d driver) NativeCompiler() mutcore.NativeCompiler {
	return newNativeCompiler(d.cfg)
}

// ScanTokens discovers the mutable tokens of the workspace by lexing every
// non-test .js file through the shared JavaScript lexical tables, in
// deterministic (path, offset) order.
func (d driver) ScanTokens(root string) ([]mutcore.Token, error) {
	return mutcore.LexicalScan(root, map[string]bool{".js": true}, isTestFile,
		jsLexSpec{}, mutcore.ClassifyLexical)
}

// isTestFile reports whether a workspace-relative path holds the assertion
// oracle rather than production code. Test files themselves never mutate:
// rewriting the oracle would make every verdict meaningless.
func isTestFile(rel string) bool {
	base := filepath.Base(rel)
	if strings.HasPrefix(base, "test-") || strings.HasSuffix(base, ".test.js") ||
		strings.HasSuffix(base, "_test.js") {
		return true
	}
	dir := filepath.ToSlash(filepath.Dir(rel))
	return dir == "test" || strings.HasPrefix(dir, "test/")
}

// RequiredFile reports whether a file must exist in a staged workspace for
// node to resolve and execute the tree: JavaScript sources plus the
// package manifest.
func (d driver) RequiredFile(name string) bool {
	return strings.HasSuffix(name, ".js") || name == "package.json"
}

// Compile reports whether the staged mutant workspace still parses,
// judged by node --check on the mutated file. A mutant that fails the
// syntax check is killed by the toolchain itself, without spawning any
// doomed test run.
func (d driver) Compile(ctx context.Context, dir string, tok mutcore.Token) bool {
	cmd := exec.CommandContext(ctx, d.cfg.NodeBin, "--check",
		filepath.Join(dir, filepath.FromSlash(tok.Path)))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// DiscoverTests enumerates the executable test units of the workspace:
// one unit per JavaScript file that declares test functions or imports
// node:test, in deterministic identifier order. node --test executes an
// explicitly passed file as one subtest, so the file is the honest judging
// granularity for a framework-free target.
func (d driver) DiscoverTests(ctx context.Context, root string) ([]mutcore.TestRef, error) {
	detector := &testFileDetector{}
	_, err := mutcore.LexicalScan(root, map[string]bool{".js": true}, nil,
		jsLexSpec{}, detector.classify)
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
// of every file declaring a test function or importing node:test: a
// function keyword arming the name check, a test_* identifier completing
// it, a node:test string in a require/import completing it too, and EOF
// markers resetting the per-file state.
type testFileDetector struct {
	files      []string
	sawFunc    bool
	sawRequire bool
	reported   bool
}

// classify implements the token-side detection. The path argument carries
// the workspace-relative slash path of the file owning the token, supplied
// by the scan driver at every call.
func (d *testFileDetector) classify(path string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		d.sawFunc, d.sawRequire, d.reported = false, false, false
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent:
		return mutcore.ClassUnknown, false
	}
	if d.reported {
		return mutcore.ClassUnknown, false
	}
	if d.sawFunc {
		d.sawFunc = false
		if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, "test") {
			d.reported = true
			d.files = append(d.files, path)
		}
		return mutcore.ClassUnknown, false
	}
	if d.sawRequire {
		d.sawRequire = false
		if lt.Kind == frontkit.TokenStr && strings.Contains(lt.Text, "node:test") {
			d.reported = true
			d.files = append(d.files, path)
		}
		return mutcore.ClassUnknown, false
	}
	switch lt.Text {
	case "function":
		d.sawFunc = true
	case "require", "import":
		d.sawRequire = true
	}
	return mutcore.ClassUnknown, false
}

// RunTest executes one test file inside the staged mutant workspace
// through node --test. A nil error with a false verdict means the file's
// assertions failed (kill); a non-nil error means node itself failed to
// start and the verdict must be treated as inconclusive.
func (d driver) RunTest(ctx context.Context, dir string, test mutcore.TestRef) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	cmd := exec.CommandContext(ctx, d.cfg.NodeBin, "--test", test.Name)
	cmd.Dir = dir
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	runErr := cmd.Run()
	if runErr == nil {
		return true, nil
	}
	var exitErr *exec.ExitError
	if errors.As(runErr, &exitErr) {
		return false, nil
	}
	return false, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, runErr)
}
