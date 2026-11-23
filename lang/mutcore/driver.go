/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import "context"

// Driver is the contract every language backend implements to plug its
// toolchain into the shared mutation session core. The core owns the
// staged-workspace pool, the outcome cache and the kernel port adapters;
// the driver owns everything that differs between languages: how source
// files are lexed into mutable tokens, how tests are discovered, how
// per-test coverage is measured on the pristine tree, and how one staged
// mutant is compiled and judged.
//
// Every method must be deterministic for a fixed workspace snapshot: the
// mutation kernel replays and parallelises evaluations, so any hidden
// randomness or host-state dependence would corrupt report fingerprints.
type Driver interface {
	// Language returns the canonical short tag of the backend (py, js,
	// c, cs, rs, kt, ...), used in diagnostics and CLI dispatch.
	Language() string

	// ScanTokens discovers the mutable tokens of the workspace rooted at
	// root, in deterministic (path, offset) order. It must fail with
	// ErrNoSourceFiles when the tree carries no mutable source file.
	ScanTokens(root string) ([]Token, error)

	// DiscoverTests enumerates the tests the runner can execute against
	// the pristine workspace, in deterministic identifier order. It must
	// fail with ErrNoTests when the tree declares no test at all, and
	// with ErrDiscoveryFailed when the toolchain probing itself failed.
	DiscoverTests(ctx context.Context, root string) ([]TestRef, error)

	// MeasureCoverage runs every discovered test once on the pristine
	// tree and records which lines of which files each test executed.
	// A failing pristine test is fatal: no mutant may be judged against
	// a red baseline. Failures wrap ErrCoverageFailed.
	MeasureCoverage(ctx context.Context, root string, tests []TestRef) (*CoverMatrix, error)

	// RequiredFile reports whether a file of the staged workspace must be
	// materialized for the toolchain to compile and test: sources plus
	// the manifests the build resolves. Everything else is skipped to
	// keep per-mutant staging cheap.
	RequiredFile(name string) bool

	// Compile reports whether the staged mutant workspace still builds,
	// judging by the narrowest honest build the language offers. A false
	// verdict is a kill delivered by the toolchain itself. The caller
	// passes the engine context so a hung toolchain process is bounded
	// by a deadline instead of blocking the worker forever.
	Compile(ctx context.Context, dir string, tok Token) bool

	// RunTest executes exactly one discovered test inside the staged
	// mutant workspace. A nil error with a false verdict means the test
	// ran and failed (kill); a non-nil error means the runner itself
	// failed and the verdict must be treated as inconclusive, never as
	// a fabricated kill.
	RunTest(ctx context.Context, dir string, test TestRef) (bool, error)
}
