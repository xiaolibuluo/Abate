/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package pymutate implements the real pytest execution backend for the
// mutation kernel (KRN-MUT). It shares the mutcore session engine and
// supplies only the Python toolchain knowledge: frontkit-based lexical
// scanning, pytest discovery, per-test coverage via coverage.py and
// per-mutant judgement by executing the mutated tree.
package pymutate

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
	// PythonBin is the interpreter used for plain-script execution.
	PythonBin string
	// PytestBin is the pytest entry point (binary or -m module style).
	PytestBin string
	// TestTimeout bounds one test invocation in seconds.
	TestTimeout int
}

// DefaultConfig returns the production configuration: python3 and pytest
// resolved from PATH, and a 60-second per-test bound.
func DefaultConfig() config {
	return config{PythonBin: "python3", PytestBin: "pytest", TestTimeout: 60}
}

// driver is the mutcore.Driver implementation for Python workspaces.
type driver struct {
	cfg config
}

// New builds a real-workspace mutation session rooted at root. The driver
// probes both python3 and pytest before any scanning starts, so a host
// without the toolchain fails fast with ErrToolchainMissing instead of a
// half-probed baseline.
func New(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg := DefaultConfig()
	if _, err := exec.LookPath(cfg.PythonBin); err != nil {
		return nil, fmt.Errorf("%w: python3", mutcore.ErrToolchainMissing)
	}
	if _, err := exec.LookPath(cfg.PytestBin); err != nil {
		return nil, fmt.Errorf("%w: pytest", mutcore.ErrToolchainMissing)
	}
	return mutcore.New(ctx, driver{cfg: cfg}, root)
}

// Language returns the canonical short tag of the backend.
func (d driver) Language() string { return "py" }

// NativeCompiler returns the two-level gating accelerator for Python.
func (d driver) NativeCompiler() mutcore.NativeCompiler {
	return newNativeCompiler(d.cfg)
}

// isTestFile reports whether a workspace-relative path is a pytest
// discovery target. Test files themselves never mutate: rewriting the
// oracle would make every verdict meaningless.
func isTestFile(rel string) bool {
	base := filepath.Base(rel)
	return strings.HasPrefix(base, "test_") || strings.HasSuffix(base, "_test.py")
}

// RequiredFile reports whether a file must exist in a staged workspace for
// the interpreter and pytest to work: Python sources plus the pytest and
// coverage manifests.
func (d driver) RequiredFile(name string) bool {
	if strings.HasSuffix(name, ".py") {
		return true
	}
	switch name {
	case "pytest.ini", "pyproject.toml", "setup.cfg", "tox.ini", ".coveragerc":
		return true
	}
	return false
}

// Compile reports whether the staged mutant workspace still parses. Python
// has no separate build step: the honest compile verdict is a syntax check
// of the mutated file, which the interpreter performs at a fraction of the
// cost of a full import.
func (d driver) Compile(ctx context.Context, dir string, tok mutcore.Token) bool {
	return compileFile(ctx, d.cfg, filepath.Join(dir, filepath.FromSlash(tok.Path)))
}

// compileFile runs py_compile against one file. A mutant that fails to
// parse is killed by the toolchain itself, without spawning any doomed
// test run.
func compileFile(ctx context.Context, cfg config, file string) bool {
	if ctx == nil {
		ctx = context.Background()
	}
	cmd := exec.CommandContext(ctx, cfg.PythonBin, "-m", "py_compile", file)
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// DiscoverTests enumerates the tests pytest can execute against the
// pristine workspace in two deterministic stages: a lexical pass lists the
// source files that declare test functions, then one --collect-only run
// against those explicit paths returns the node ids. Explicit paths bypass
// the python_files name filter, so a workspace needs no pytest.ini for its
// tests to be found, and collection never executes test bodies. A workspace
// whose collection is empty fails with ErrNoTests; a collection error fails
// with ErrDiscoveryFailed.
func (d driver) DiscoverTests(ctx context.Context, root string) ([]mutcore.TestRef, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	files, err := d.testFiles(root)
	if err != nil {
		return nil, err
	}
	args := append([]string{"--collect-only", "-q", "--disable-warnings",
		"-p", "no:cacheprovider"}, files...)
	cmd := exec.CommandContext(ctx, d.cfg.PytestBin, args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%w: %v", mutcore.ErrDiscoveryFailed, tail(buf.String()))
	}
	tests := parseCollected(buf.String(), root)
	if len(tests) == 0 {
		return nil, mutcore.ErrNoTests
	}
	sort.Slice(tests, func(i, j int) bool { return tests[i].ID < tests[j].ID })
	return tests, nil
}

// testFiles lists the workspace .py files declaring at least one test
// function, in deterministic order. The detection is lexical: a def token
// followed by a test_* identifier, regardless of the file name, so the
// backend never depends on pytest's python_files pattern to find the
// oracle of a single-file target.
func (d driver) testFiles(root string) ([]string, error) {
	detector := &testFileDetector{}
	_, err := mutcore.LexicalScan(root, map[string]bool{".py": true}, nil,
		pyLexSpec{}, detector.classify)
	if err != nil && len(detector.files) == 0 {
		return nil, mutcore.ErrNoTests
	}
	return detector.files, nil
}

// testFileDetector is a lexical scan sink that records the relative path
// of every file declaring a test_* function: a def token arms the name
// check, a test_* identifier in the next position records the file once,
// and EOF markers reset the per-file state so one detector instance can
// walk the whole workspace stream.
type testFileDetector struct {
	files    []string
	sawDef   bool
	reported bool
}

// classify implements the token-side detection. The path argument carries
// the workspace-relative slash path of the file owning the token, supplied
// by the scan driver at every call.
func (d *testFileDetector) classify(path string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		d.sawDef = false
		d.reported = false
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent:
		return mutcore.ClassUnknown, false
	}
	if d.sawDef {
		d.sawDef = false
		if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, "test_") && !d.reported {
			d.reported = true
			d.files = append(d.files, path)
		}
		return mutcore.ClassUnknown, false
	}
	if lt.Text == "def" {
		d.sawDef = true
	}
	return mutcore.ClassUnknown, false
}

// parseCollected turns pytest --collect-only -q output into the sorted
// test reference list. Each line of the quiet collector is one node id in
// path::Class::test or path::test form; the identifier is deterministic
// because pytest enumerates files in directory order.
func parseCollected(out, root string) []mutcore.TestRef {
	var tests []mutcore.TestRef
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || !strings.Contains(line, "::") {
			continue
		}
		file := strings.SplitN(line, "::", 2)[0]
		tests = append(tests, mutcore.TestRef{
			ID:   line,
			Dir:  filepath.Join(root, filepath.FromSlash(filepath.Dir(file))),
			Name: line,
		})
	}
	return tests
}

// RunTest executes exactly one collected node id inside the staged mutant
// workspace. Node ids are workspace-relative by construction (discovery
// anchors collection at the pristine root), and the staged tree mirrors
// the same layout, so the id is passed through unchanged with the staged
// root as the working directory. A nil error with a false verdict means
// the test ran and failed (kill); a non-nil error means pytest itself
// failed to start and the verdict must be treated as inconclusive.
func (d driver) RunTest(ctx context.Context, dir string, test mutcore.TestRef) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx,
		timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	cmd := exec.CommandContext(ctx, d.cfg.PytestBin,
		"-q", "-x", "--disable-warnings", "-p", "no:cacheprovider", test.ID)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
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

// tail truncates one toolchain output blob to the last 512 bytes so error
// values stay readable without flooding the report.
func tail(out string) string {
	const max = 512
	if len(out) <= max {
		return out
	}
	return out[len(out)-max:]
}
