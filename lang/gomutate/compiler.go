/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gomutate

import (
	"context"
	"go/parser"
	"go/token"
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the Go implementation of mutcore.NativeCompiler. It
// reuses the session's toolchain configuration so every invocation honors
// the same GOFLAGS/GOCACHE pinning as the legacy pipeline, keeping the
// verdicts of the two modes comparable.
type nativeCompiler struct {
	cfg runConfig
}

// compile-time check: nativeCompiler must satisfy the mutcore interface.
var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

// newNativeCompiler returns the two-level gating accelerator bound to one
// runConfig. The value is stateless apart from the config, so concurrent
// judging workers can share it without synchronization.
func newNativeCompiler(cfg runConfig) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg}
}

// FastCompile reports whether the mutated file still parses, using the
// standard library parser only: no type checking, no imports resolution,
// no toolchain subprocess. A rewrite that breaks syntax is killed here in
// microseconds instead of paying for a `go build` invocation. The file is
// read from the staged workspace, never from the pristine tree.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	fset := token.NewFileSet()
	target := filepath.Join(dir, filepath.FromSlash(file))
	// parser.SkipObjectResolution keeps the pass cheap: mutation verdicts
	// never need the resolved object graph, only syntactic validity.
	_, err := parser.ParseFile(fset, target, nil, parser.SkipObjectResolution)
	return err == nil
}

// PackageBuild reports whether the package holding the mutated file still
// compiles, delegating to the same buildPackage helper the legacy
// compileTree uses. Sharing the helper guarantees a mutant killed by the
// native gate would receive the identical verdict on the legacy path.
func (c *nativeCompiler) PackageBuild(dir, file string) bool {
	return buildPackage(c.cfg, dir, file)
}

// PackageTest runs exactly one discovered test against the staged
// workspace and reports whether it passed. A non-nil error means the
// toolchain invocation itself failed, which the caller must treat as
// inconclusive rather than a kill.
func (c *nativeCompiler) PackageTest(dir, testDir, testName string) (bool, error) {
	passed, _, runErr := runGoTest(context.Background(), c.cfg,
		filepath.Join(dir, filepath.FromSlash(testDir)), "-run", "^"+testName+"$")
	if runErr != nil {
		return false, runErr
	}
	return passed, nil
}
