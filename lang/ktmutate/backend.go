/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package ktmutate implements the real kotlinc execution backend for the
// mutation kernel (KRN-MUT). It shares the mutcore session engine and
// supplies only the Kotlin toolchain knowledge: frontkit-based lexical
// scanning, check-based test discovery, line coverage through the JaCoCo
// agent and report tool, and per-mutant judgement by compiling the
// mutated tree with kotlinc and running it on the JVM.
package ktmutate

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
	// KotlincBin is the Kotlin compiler executable.
	KotlincBin string
	// JavaBin is the JVM launcher used for every run.
	JavaBin string
	// JavapBin disassembles compiled classes for entry-point detection.
	JavapBin string
	// StdlibJar is the Kotlin standard library jar every run links
	// against, resolved from the kotlinc installation tree.
	StdlibJar string
	// JacocoAgent is the JaCoCo agent jar used for coverage runs.
	JacocoAgent string
	// JacocoCli is the JaCoCo report jar that renders exec data as XML.
	JacocoCli string
	// TestTimeout bounds one test invocation in seconds.
	TestTimeout int
}

// driver is the mutcore.Driver implementation for Kotlin workspaces.
type driver struct {
	cfg config
}

// New builds a real-workspace mutation session rooted at root. The driver
// probes kotlinc, the JVM and the JaCoCo tools before any scanning
// starts, so a host without the toolchain fails fast with
// ErrToolchainMissing instead of a half-probed baseline.
func New(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	return mutcore.New(ctx, driver{cfg: cfg}, root)
}

// defaultConfig resolves every tool the backend needs: kotlinc, java and
// javap from PATH, the standard library jar from the kotlinc installation
// tree, and the JaCoCo agent/report jars from their well-known locations
// or operator overrides. Every missing piece aborts with
// ErrToolchainMissing so the backend never starts a session it cannot
// measure or judge.
func defaultConfig() (config, error) {
	kotlinc, err := exec.LookPath("kotlinc")
	if err != nil {
		return config{}, fmt.Errorf("%w: kotlinc", mutcore.ErrToolchainMissing)
	}
	javaBin, err := exec.LookPath("java")
	if err != nil {
		return config{}, fmt.Errorf("%w: java", mutcore.ErrToolchainMissing)
	}
	javap, err := exec.LookPath("javap")
	if err != nil {
		return config{}, fmt.Errorf("%w: javap", mutcore.ErrToolchainMissing)
	}
	stdlib, err := resolveStdlib(kotlinc)
	if err != nil {
		return config{}, err
	}
	agent, err := resolveJacoco("JACOCO_AGENT", "jacocoagent.jar", os.Getenv)
	if err != nil {
		return config{}, err
	}
	cli, err := resolveJacoco("JACOCO_CLI", "jacococli.jar", os.Getenv)
	if err != nil {
		return config{}, err
	}
	return config{
		KotlincBin:  kotlinc,
		JavaBin:     javaBin,
		JavapBin:    javap,
		StdlibJar:   stdlib,
		JacocoAgent: agent,
		JacocoCli:   cli,
		TestTimeout: 60,
	}, nil
}

// resolveStdlib locates kotlin-stdlib.jar inside the kotlinc installation
// tree by resolving the compiler symlink chain: <home>/bin/kotlinc
// implies <home>/lib/kotlin-stdlib.jar.
func resolveStdlib(kotlinc string) (string, error) {
	resolved, err := filepath.EvalSymlinks(kotlinc)
	if err != nil {
		resolved = kotlinc
	}
	home := filepath.Dir(filepath.Dir(resolved))
	candidate := filepath.Join(home, "lib", "kotlin-stdlib.jar")
	if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
		return candidate, nil
	}
	return "", fmt.Errorf("%w: kotlin-stdlib.jar under %s", mutcore.ErrToolchainMissing, home)
}

// resolveJacoco locates one JaCoCo jar: an operator-set environment
// variable wins, then the well-known /opt/jacoco install directory, then
// the user-local ~/.jacoco directory. The environment lookup is injected
// by the caller so this resolver stays free of direct process-state reads.
func resolveJacoco(envKey, name string, getenv func(string) string) (string, error) {
	if env := getenv(envKey); env != "" {
		if info, err := os.Stat(env); err == nil && !info.IsDir() {
			return env, nil
		}
		return "", fmt.Errorf("%w: %s points at %s which is not a file",
			mutcore.ErrToolchainMissing, envKey, env)
	}
	candidates := []string{filepath.Join("/opt/jacoco", name)}
	if home, err := os.UserHomeDir(); err == nil {
		candidates = append(candidates, filepath.Join(home, ".jacoco", name))
	}
	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("%w: %s (install JaCoCo or set %s)",
		mutcore.ErrToolchainMissing, name, envKey)
}

// Language returns the canonical short tag of the backend.
func (d driver) Language() string { return "kt" }

// NativeCompiler returns the two-level gating accelerator for Kotlin.
func (d driver) NativeCompiler() mutcore.NativeCompiler {
	return newNativeCompiler(d.cfg)
}

// ScanTokens discovers the mutable tokens of the workspace by lexing every
// non-test .kt file through the shared Kotlin lexical tables, in
// deterministic (path, offset) order.
func (d driver) ScanTokens(root string) ([]mutcore.Token, error) {
	return mutcore.LexicalScan(root, map[string]bool{".kt": true}, isTestFile,
		ktLexSpec{}, mutcore.ClassifyLexical)
}

// isTestFile reports whether a workspace-relative path holds the assertion
// oracle rather than production code. Test files themselves never mutate:
// rewriting the oracle would make every verdict meaningless.
func isTestFile(rel string) bool {
	base := filepath.Base(rel)
	if strings.HasPrefix(base, "test_") || strings.HasPrefix(base, "test-") ||
		strings.HasSuffix(base, "_test.kt") || strings.HasSuffix(base, ".test.kt") ||
		strings.HasSuffix(base, "tests.kt") {
		return true
	}
	dir := filepath.ToSlash(filepath.Dir(rel))
	return dir == "test" || dir == "tests" ||
		strings.HasPrefix(dir, "test/") || strings.HasPrefix(dir, "tests/")
}

// RequiredFile reports whether a file must exist in a staged workspace for
// kotlinc to compile and run the tree: Kotlin sources only.
func (d driver) RequiredFile(name string) bool {
	return strings.EqualFold(filepath.Ext(name), ".kt")
}

// sourceFiles enumerates the workspace .kt files in deterministic
// lexicographic order, so a staged mutant compiles exactly the same
// compilation unit set as the pristine tree.
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
		if strings.EqualFold(filepath.Ext(entry.Name()), ".kt") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

// Compile reports whether the staged mutant workspace still compiles,
// judged by kotlinc on the mutated file into a scratch class directory.
// A mutant that fails the compile is killed by the toolchain itself,
// without spawning any doomed run. Every invocation is a fresh compiler
// process, so parallel judging workers never share compiler state.
func (d driver) Compile(ctx context.Context, dir string, tok mutcore.Token) bool {
	classDir, err := os.MkdirTemp("", "abate-ktmutate-csc-")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(classDir) }()
	cmd := exec.CommandContext(ctx, d.cfg.KotlincBin,
		filepath.Join(dir, filepath.FromSlash(tok.Path)), "-d", classDir)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// DiscoverTests enumerates the executable test units of the workspace:
// one unit per .kt file that declares at least one fun test_* function
// behind the check assertion builtin, in deterministic identifier order.
// The Kotlin harness is one fun main that calls every test_*, so the file
// is the honest judging granularity for a framework-free target.
func (d driver) DiscoverTests(ctx context.Context, root string) ([]mutcore.TestRef, error) {
	detector := &testFileDetector{}
	_, err := mutcore.LexicalScan(root, map[string]bool{".kt": true}, nil,
		ktLexSpec{}, detector.classify)
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
// of every file declaring a fun test_* function: a fun keyword arming the
// name check, a test_* identifier completing it, and EOF markers
// resetting the per-file state.
type testFileDetector struct {
	files    []string
	sawFun   bool
	reported bool
}

// classify implements the token-side detection. The path argument carries
// the workspace-relative slash path of the file owning the token, supplied
// by the scan driver at every call.
func (d *testFileDetector) classify(path string, lt frontkit.Token) (mutcore.TokenClass, bool) {
	switch lt.Kind {
	case frontkit.TokenEOF:
		d.sawFun, d.reported = false, false
		return mutcore.ClassUnknown, false
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent,
		frontkit.TokenComment:
		return mutcore.ClassUnknown, false
	}
	if d.reported {
		return mutcore.ClassUnknown, false
	}
	if d.sawFun {
		d.sawFun = false
		if lt.Kind == frontkit.TokenIdent && strings.HasPrefix(lt.Text, "test_") {
			d.reported = true
			d.files = append(d.files, path)
		}
		return mutcore.ClassUnknown, false
	}
	if lt.Text == "fun" {
		d.sawFun = true
	}
	return mutcore.ClassUnknown, false
}

// RunTest executes the staged mutant workspace by compiling every .kt
// source into one class directory and running its entry point. A nil
// error with a false verdict means the checks threw or the program exited
// non-zero (kill); a non-nil error means the toolchain itself failed to
// build or start and the verdict must be treated as inconclusive, never
// as a fabricated kill.
func (d driver) RunTest(ctx context.Context, dir string, test mutcore.TestRef) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	build, cleanup, err := d.buildClasses(ctx, dir)
	if err != nil {
		return false, err
	}
	defer cleanup()
	passed, err := d.runMain(ctx, build, dir, "")
	if err != nil {
		return false, err
	}
	return passed, nil
}

// build holds one compiled class directory of a workspace: its location,
// the application class files inside it and the detected entry point.
type build struct {
	classDir string
	classes  []string
	main     string
}

// buildClasses compiles every .kt source of dir into one scratch class
// directory and detects its JVM entry point. The class directory lives
// outside the tree under test so it never collects build artifacts. It
// returns the build descriptor and a cleanup that releases the scratch
// space.
func (d driver) buildClasses(ctx context.Context, dir string) (build, func(), error) {
	sources, err := sourceFiles(dir)
	if err != nil {
		return build{}, func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	if len(sources) == 0 {
		return build{}, func() {}, fmt.Errorf("%w: no .kt sources in %s", mutcore.ErrWorkspaceFailed, dir)
	}
	classDir, err := os.MkdirTemp("", "abate-ktmutate-build-")
	if err != nil {
		return build{}, func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	cleanup := func() { _ = os.RemoveAll(classDir) }
	args := make([]string, 0, len(sources)+2)
	for _, src := range sources {
		args = append(args, filepath.Join(dir, src))
	}
	args = append(args, "-d", classDir)
	cmd := exec.CommandContext(ctx, d.cfg.KotlincBin, args...)
	cmd.Dir = dir
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		cleanup()
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// A mutant that fails the batch compile was never actually
			// judged: the per-file compile check precedes it, so reaching
			// here means the tree is structurally broken beyond a
			// single-file edit. Treat the runner as failed rather than
			// fabricating a kill.
			return build{}, func() {}, fmt.Errorf("%w: kotlinc: %s",
				mutcore.ErrWorkspaceFailed, tail(buf.String()))
		}
		return build{}, func() {}, fmt.Errorf("%w: kotlinc start: %v", mutcore.ErrWorkspaceFailed, err)
	}
	classes, err := appClasses(classDir)
	if err != nil {
		cleanup()
		return build{}, func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	main, err := d.detectMain(ctx, classes)
	if err != nil {
		cleanup()
		return build{}, func() {}, err
	}
	return build{classDir: classDir, classes: classes, main: main}, cleanup, nil
}

// appClasses enumerates the top-level .class files of one class directory
// in deterministic order. Nested-package directories never appear in the
// framework-free targets the backend runs, and the Kotlin runtime is
// linked from the standard library jar instead of being embedded, so the
// root level holds exactly the application classes.
func appClasses(classDir string) ([]string, error) {
	entries, err := os.ReadDir(classDir)
	if err != nil {
		return nil, err
	}
	var classes []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".class") {
			classes = append(classes, filepath.Join(classDir, entry.Name()))
		}
	}
	sort.Strings(classes)
	if len(classes) == 0 {
		return nil, fmt.Errorf("no classes compiled into %s", classDir)
	}
	return classes, nil
}

// detectMain finds the JVM entry point among the compiled classes by
// disassembling each one with javap in deterministic order and picking
// the first that declares the standard main signature. kotlinc mangles
// top-level file functions into a <File>Kt holder class, so a name guess
// would break on any file whose name is not a plain identifier; the
// bytecode is the honest source.
func (d driver) detectMain(ctx context.Context, classes []string) (string, error) {
	for _, classFile := range classes {
		cmd := exec.CommandContext(ctx, d.cfg.JavapBin, "-p", classFile)
		out, err := cmd.Output()
		if err != nil {
			continue
		}
		if !strings.Contains(string(out), "public static void main(java.lang.String[])") {
			continue
		}
		base := filepath.Base(classFile)
		return strings.TrimSuffix(base, ".class"), nil
	}
	return "", fmt.Errorf("%w: no JVM entry point among compiled classes",
		mutcore.ErrWorkspaceFailed)
}

// runMain executes the detected entry point on the JVM with the class
// directory and the standard library jar on the classpath. The boolean
// result is the program verdict: true for a clean zero exit, false for
// any non-zero exit, failed check or timeout. A non-nil error means the
// process itself could not start, which is inconclusive. When execFile
// is non-empty the run carries the JaCoCo agent writing its execution
// data there.
func (d driver) runMain(ctx context.Context, b build, dir, execFile string) (bool, error) {
	classpath := b.classDir + string(os.PathListSeparator) + d.cfg.StdlibJar
	args := []string{}
	if execFile != "" {
		args = append(args, "-javaagent:"+d.cfg.JacocoAgent+"=destfile="+execFile)
	}
	args = append(args, "-cp", classpath, b.main)
	run := exec.CommandContext(ctx, d.cfg.JavaBin, args...)
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
	return false, fmt.Errorf("%w: java start: %v", mutcore.ErrWorkspaceFailed, runErr)
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
