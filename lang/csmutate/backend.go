/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package csmutate implements the real dotnet execution backend for the
// mutation kernel (KRN-MUT). It shares the mutcore session engine and
// supplies only the C# toolchain knowledge: frontkit-based lexical
// scanning, Check-helper test discovery, line coverage through
// dotnet-coverage, and per-mutant judgement by compiling and running the
// mutated tree through the dotnet SDK.
package csmutate

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

// harnessAssembly is the fixed assembly name of the temporary project the
// backend builds around the workspace sources.
const harnessAssembly = "AbateHarness"

// config carries the process-level execution knobs of the backend. They
// are inputs to the driver constructor, never global state, so two
// backends over two workspaces can run concurrently without interfering.
type config struct {
	// DotnetBin is the dotnet host executable.
	DotnetBin string
	// CoverageBin is the dotnet-coverage global tool used for line
	// coverage collection.
	CoverageBin string
	// CscDLL is the Roslyn compiler assembly inside the dotnet SDK, used
	// for the narrow per-mutant compile check.
	CscDLL string
	// RefDir is the reference-assemblies directory of the target
	// framework (ref/net8.0), resolved once from the dotnet packs
	// tree. The standalone csc.dll invocation needs explicit -r
	// references to the BCL assemblies because the Roslyn executor
	// (dotnet-csc) is not the build subsystem (MSBuild) and does
	// not auto-import framework references.
	RefDir string
	// DotnetRoot anchors the dotnet installation for child processes;
	// dotnet-coverage fails to resolve libhostfxr.so without it.
	DotnetRoot string
	// getenv is the injected process-environment lookup. It is wired at
	// the construction boundary so the driver never reads the process
	// environment directly (GO-CRC-028); tests can inject a fixture.
	getenv func(string) string
	// TestTimeout bounds one test invocation in seconds.
	TestTimeout int
}

// driver is the mutcore.Driver implementation for C# workspaces.
type driver struct {
	cfg config
}

// New builds a real-workspace mutation session rooted at root. The driver
// probes dotnet, the Roslyn compiler and the coverage tool before any
// scanning starts, so a host without the toolchain fails fast with
// ErrToolchainMissing instead of a half-probed baseline.
func New(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	return mutcore.New(ctx, driver{cfg: cfg}, root)
}

// defaultConfig resolves every tool the backend needs: dotnet from PATH,
// the Roslyn csc.dll from the newest installed SDK, dotnet-coverage from
// PATH or the default global-tools directory, and DOTNET_ROOT from the
// resolved dotnet host. Every missing piece aborts with
// ErrToolchainMissing so the backend never starts a session it cannot
// measure or judge.
func defaultConfig() (config, error) {
	dotnet, err := exec.LookPath("dotnet")
	if err != nil {
		return config{}, fmt.Errorf("%w: dotnet", mutcore.ErrToolchainMissing)
	}
	csc, err := resolveCsc(dotnet)
	if err != nil {
		return config{}, err
	}
	coverage, err := resolveCoverageTool()
	if err != nil {
		return config{}, err
	}
	refDir, err := resolveRefDir(csc)
	if err != nil {
		return config{}, err
	}
	return config{
		DotnetBin:   dotnet,
		CoverageBin: coverage,
		CscDLL:      csc,
		RefDir:      refDir,
		DotnetRoot:  resolveDotnetRoot(dotnet, os.Getenv),
		getenv:      os.Getenv,
		TestTimeout: 60,
	}, nil
}

// resolveRefDir derives the BCL reference-assemblies directory for the
// target framework from the resolved csc.dll path. The Roslyn compiler
// invoked standalone (as a managed DLL via dotnet <csc.dll>) does not
// carry the implicit framework references bundled by the MSBuild C#
// target, so every invocation must pass explicit -r flags to the BCL
// assemblies needed by the mutated source.
func resolveRefDir(csc string) (string, error) {
	// csc is <dotnet-root>/sdk/<version>/Roslyn/bincore/csc.dll (5 segments below root).
	// The ref pack lives at <dotnet-root>/packs/Microsoft.NETCore.App.Ref/<version>/ref/net8.0.
	// Walk up until the parent directory contains a "packs" sibling.
	dotnetRoot := filepath.Dir(csc) // bincore
	for i := 0; i < 6; i++ {
		packsProbe := filepath.Join(dotnetRoot, "packs")
		if info, probeErr := os.Stat(packsProbe); probeErr == nil && info.IsDir() {
			break
		}
		dotnetRoot = filepath.Dir(dotnetRoot)
	}
	candidate := filepath.Join(dotnetRoot, "packs", "Microsoft.NETCore.App.Ref")
	entries, err := os.ReadDir(candidate)
	if err != nil {
		return "", fmt.Errorf("%w: Microsoft.NETCore.App.Ref pack not found at %s",
			mutcore.ErrToolchainMissing, candidate)
	}
	var refDir string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		p := filepath.Join(candidate, e.Name(), "ref", "net8.0", "System.Runtime.dll")
		if info, statErr := os.Stat(p); statErr == nil && !info.IsDir() {
			refDir = filepath.Join(candidate, e.Name(), "ref", "net8.0")
		}
	}
	if refDir == "" {
		return "", fmt.Errorf("%w: BCL reference assemblies not found in %s",
			mutcore.ErrToolchainMissing, candidate)
	}
	return refDir, nil
}

// resolveCsc locates the Roslyn compiler assembly of the newest installed
// SDK by parsing dotnet --list-sdks, whose lines carry the version and the
// SDK root in bracketed form: 8.0.428 [/opt/dotnet/sdk].
func resolveCsc(dotnet string) (string, error) {
	out, err := exec.Command(dotnet, "--list-sdks").Output()
	if err != nil {
		return "", fmt.Errorf("%w: dotnet --list-sdks: %v", mutcore.ErrToolchainMissing, err)
	}
	var csc string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		open := strings.Index(line, "[")
		closeIdx := strings.Index(line, "]")
		if open <= 0 || closeIdx <= open {
			continue
		}
		version := strings.TrimSpace(line[:open])
		root := line[open+1 : closeIdx]
		candidate := filepath.Join(root, version, "Roslyn", "bincore", "csc.dll")
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			csc = candidate // last listed SDK wins: --list-sdks prints ascending versions.
		}
	}
	if csc == "" {
		return "", fmt.Errorf("%w: Roslyn csc.dll inside the dotnet SDK", mutcore.ErrToolchainMissing)
	}
	return csc, nil
}

// resolveCoverageTool locates the dotnet-coverage global tool on PATH or
// in the default .NET global-tools directory under the user home.
func resolveCoverageTool() (string, error) {
	if path, err := exec.LookPath("dotnet-coverage"); err == nil {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err == nil {
		candidate := filepath.Join(home, ".dotnet", "tools", "dotnet-coverage")
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("%w: dotnet-coverage (install the dotnet-coverage global tool)",
		mutcore.ErrToolchainMissing)
}

// resolveDotnetRoot derives the installation root of the dotnet host by
// resolving the symlink chain of the binary found on PATH. An existing
// DOTNET_ROOT environment value wins, so an operator override is never
// clobbered. The environment lookup is injected by the caller so this
// resolver stays free of direct process-state reads.
func resolveDotnetRoot(dotnet string, getenv func(string) string) string {
	if env := getenv("DOTNET_ROOT"); env != "" {
		return env
	}
	resolved, err := filepath.EvalSymlinks(dotnet)
	if err != nil {
		resolved = dotnet
	}
	return filepath.Dir(resolved)
}

// env returns the process environment for toolchain invocations,
// injecting DOTNET_ROOT when it is known so child processes always
// resolve the same runtime the driver probed.
func (d driver) env() []string {
	env := os.Environ()
	if d.cfg.DotnetRoot != "" && d.cfg.getenv("DOTNET_ROOT") == "" {
		env = append(env, "DOTNET_ROOT="+d.cfg.DotnetRoot)
	}
	return env
}

// Language returns the canonical short tag of the backend.
func (d driver) Language() string { return "cs" }

// NativeCompiler returns the two-level gating accelerator for C#.
func (d driver) NativeCompiler() mutcore.NativeCompiler {
	return newNativeCompiler(d.cfg)
}

// ScanTokens discovers the mutable tokens of the workspace by lexing every
// non-test .cs file through the shared C# lexical tables, in deterministic
// (path, offset) order.
func (d driver) ScanTokens(root string) ([]mutcore.Token, error) {
	return mutcore.LexicalScan(root, map[string]bool{".cs": true}, isTestFile,
		csLexSpec{}, mutcore.ClassifyLexical)
}

// isTestFile reports whether a workspace-relative path holds the assertion
// oracle rather than production code. Test files themselves never mutate:
// rewriting the oracle would make every verdict meaningless.
func isTestFile(rel string) bool {
	base := filepath.Base(rel)
	if strings.HasPrefix(base, "test_") || strings.HasPrefix(base, "test-") ||
		strings.HasSuffix(base, "_test.cs") || strings.HasSuffix(base, ".test.cs") ||
		strings.HasSuffix(base, "tests.cs") {
		return true
	}
	dir := filepath.ToSlash(filepath.Dir(rel))
	return dir == "test" || dir == "tests" ||
		strings.HasPrefix(dir, "test/") || strings.HasPrefix(dir, "tests/")
}

// RequiredFile reports whether a file must exist in a staged workspace for
// the dotnet toolchain to compile and run the tree: C# sources plus any
// project manifest the workspace already carries.
func (d driver) RequiredFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".cs" || ext == ".csproj"
}

// sourceFiles enumerates the workspace .cs files in deterministic
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
		if strings.EqualFold(filepath.Ext(entry.Name()), ".cs") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

// Compile reports whether the staged mutant workspace still compiles,
// judged by the Roslyn compiler on the mutated file. A mutant that fails
// the compile is killed by the toolchain itself, without spawning any
// doomed build or run.
func (d driver) Compile(ctx context.Context, dir string, tok mutcore.Token) bool {
	outDir, err := os.MkdirTemp("", "abate-csmutate-csc-")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(outDir) }()
	args := []string{
		d.cfg.CscDLL, "/nologo", "/t:exe",
		"/out:" + filepath.Join(outDir, harnessAssembly+".dll"),
	}
	if d.cfg.RefDir != "" {
		args = append(args,
			"-r:"+filepath.Join(d.cfg.RefDir, "System.Runtime.dll"),
			"-r:"+filepath.Join(d.cfg.RefDir, "System.Console.dll"),
			"-r:"+filepath.Join(d.cfg.RefDir, "System.Linq.dll"),
			"-r:"+filepath.Join(d.cfg.RefDir, "System.Collections.dll"),
		)
	}
	args = append(args, filepath.Join(dir, filepath.FromSlash(tok.Path)))
	cmd := exec.CommandContext(ctx, d.cfg.DotnetBin, args...)
	cmd.Env = d.env()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// DiscoverTests enumerates the executable test units of the workspace:
// one unit per .cs file that declares at least one void test_* function
// behind the Check assertion helper, in deterministic identifier order.
// The C# harness runs every test_* from its top-level statements, so the
// file is the honest judging granularity for a framework-free target.
func (d driver) DiscoverTests(ctx context.Context, root string) ([]mutcore.TestRef, error) {
	detector := &testFileDetector{}
	_, err := mutcore.LexicalScan(root, map[string]bool{".cs": true}, nil,
		csLexSpec{}, detector.classify)
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

// RunTest executes the staged mutant workspace by building it into one
// assembly and running it. A nil error with a false verdict means the
// assertions threw or the program exited non-zero (kill); a non-nil error
// means the toolchain itself failed to build or start and the verdict
// must be treated as inconclusive, never as a fabricated kill.
func (d driver) RunTest(ctx context.Context, dir string, test mutcore.TestRef) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	buildDir, cleanup, err := d.buildProject(ctx, dir)
	if err != nil {
		return false, err
	}
	defer cleanup()
	passed, err := d.runAssembly(ctx, filepath.Join(buildDir, "out", harnessAssembly+".dll"), dir)
	if err != nil {
		return false, err
	}
	return passed, nil
}

// buildProject writes the temporary harness project around the workspace
// sources and builds it with dotnet build. The project lives in a scratch
// directory so the tree under test never collects build artifacts; its
// Compile items reference the sources by absolute path, which keeps the
// PDB filenames anchored at the tree the coverage report must map back
// to. It returns the scratch directory holding the built assembly and a
// cleanup that releases it.
func (d driver) buildProject(ctx context.Context, dir string) (string, func(), error) {
	sources, err := sourceFiles(dir)
	if err != nil {
		return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	if len(sources) == 0 {
		return "", func() {}, fmt.Errorf("%w: no .cs sources in %s", mutcore.ErrWorkspaceFailed, dir)
	}
	buildDir, err := os.MkdirTemp("", "abate-csmutate-build-")
	if err != nil {
		return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	cleanup := func() { _ = os.RemoveAll(buildDir) }
	var includes strings.Builder
	for _, src := range sources {
		includes.WriteString("    <Compile Include=\"")
		includes.WriteString(filepath.Join(dir, src))
		includes.WriteString("\" />\n")
	}
	project := "<Project Sdk=\"Microsoft.NET.Sdk\">\n" +
		"  <PropertyGroup>\n" +
		"    <OutputType>Exe</OutputType>\n" +
		"    <TargetFramework>net8.0</TargetFramework>\n" +
		"    <Nullable>disable</Nullable>\n" +
		"    <ImplicitUsings>disable</ImplicitUsings>\n" +
		"    <EnableDefaultCompileItems>false</EnableDefaultCompileItems>\n" +
		"    <AssemblyName>" + harnessAssembly + "</AssemblyName>\n" +
		"  </PropertyGroup>\n" +
		"  <ItemGroup>\n" + includes.String() + "  </ItemGroup>\n" +
		"</Project>\n"
	projectPath := filepath.Join(buildDir, harnessAssembly+".csproj")
	if err := os.WriteFile(projectPath, []byte(project), 0o644); err != nil {
		cleanup()
		return "", func() {}, fmt.Errorf("%w: %v", mutcore.ErrWorkspaceFailed, err)
	}
	cmd := exec.CommandContext(ctx, d.cfg.DotnetBin, "build", projectPath,
		"-c", "Release", "--nologo", "-v", "q", "-o", filepath.Join(buildDir, "out"))
	cmd.Dir = buildDir
	cmd.Env = d.env()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		cleanup()
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// A mutant that fails the project build was never actually
			// judged: the per-file compile check precedes it, so reaching
			// here means the tree is structurally broken beyond a
			// single-file edit. Treat the runner as failed rather than
			// fabricating a kill.
			return "", func() {}, fmt.Errorf("%w: dotnet build: %s",
				mutcore.ErrWorkspaceFailed, tail(buf.String()))
		}
		return "", func() {}, fmt.Errorf("%w: dotnet build start: %v", mutcore.ErrWorkspaceFailed, err)
	}
	return buildDir, cleanup, nil
}

// runAssembly executes the built harness assembly through the dotnet
// host. The boolean result is the program verdict: true for a clean zero
// exit, false for any non-zero exit, assertion throw or timeout. A
// non-nil error means the process itself could not start, which is
// inconclusive.
func (d driver) runAssembly(ctx context.Context, assembly, dir string) (bool, error) {
	run := exec.CommandContext(ctx, d.cfg.DotnetBin, assembly)
	run.Dir = dir
	run.Env = d.env()
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
