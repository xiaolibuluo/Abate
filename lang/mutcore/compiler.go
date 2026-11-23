/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutcore

// NativeCompiler is an optional language-specific accelerator for the
// native-sandbox mutation path. A backend that implements it unlocks the
// two-level gating pipeline: every mutant first passes a cheap per-file
// syntax check that rejects broken rewrites without spawning the full
// toolchain, and only survivors pay for a package-level build plus the
// targeted test run.
//
// All methods run inside a staged mutant workspace: dir is the materialized
// mutant tree root, and file is the workspace-relative path of the mutated
// source file. Implementations must be safe for concurrent use: the kernel
// judges mutants in parallel and never serializes toolchain invocations.
type NativeCompiler interface {
	// FastCompile reports whether the mutated file is still syntactically
	// valid, judged by the cheapest language-level check available (for Go
	// this is `gofmt -e`-style parsing, for Python `py_compile`, for JS a
	// parser pass). A false verdict kills the mutant immediately: a file
	// that does not parse can never build or pass a test, so skipping the
	// expensive stages saves one full toolchain invocation per broken
	// rewrite.
	//
	// The check must not read or write outside dir, and must not depend
	// on network access. Implementations should finish in tens of
	// milliseconds; anything slower defeats the purpose of the gate.
	FastCompile(dir, file string) bool

	// PackageBuild reports whether the package holding the mutated file
	// still compiles, judged by the language build tool against the
	// staged workspace. The verdict must be identical to the build step
	// of the non-native path so a mutant killed here would also have been
	// killed by the legacy pipeline.
	PackageBuild(dir, file string) bool

	// PackageTest runs exactly one test against the mutated workspace and
	// reports whether it passed. testName is the runner-local name of the
	// test (the Name field of the discovered TestRef), and testDir is the
	// workspace-relative directory the runner must execute in. A non-nil
	// error means the toolchain itself failed (binary missing, sandbox
	// broken), which the caller must treat as inconclusive rather than a
	// kill.
	PackageTest(dir, testDir, testName string) (bool, error)
}

// NativeCompilerOf returns the NativeCompiler a backend exposes, or nil
// when the backend only supports the default build-and-test pipeline. The
// session layer type-asserts its driver against this interface once at
// construction and caches the result.
type NativeCompilerProvider interface {
	// NativeCompiler returns the compiler implementation for this driver,
	// or nil when the language has no two-level gating support yet.
	NativeCompiler() NativeCompiler
}
