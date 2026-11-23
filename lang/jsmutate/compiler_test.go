/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package jsmutate

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// jsToolchain returns the production config or skips the test when node is
// absent from the host.
func jsToolchain(t *testing.T) config {
	t.Helper()
	cfg := DefaultConfig()
	if _, err := exec.LookPath(cfg.NodeBin); err != nil {
		t.Skipf("Node toolchain unavailable: %v", err)
	}
	return cfg
}

// writeJSModule creates a temporary .js file and returns its directory.
func writeJSModule(t *testing.T, src string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "lib.js"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestJSFastCompileRejectsSyntaxError(t *testing.T) {
	cfg := jsToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeJSModule(t, "function broken( { let")
	if nc.FastCompile(dir, "lib.js") {
		t.Fatal("node --check must reject a syntax error")
	}
	// JS: PackageBuild is the same node --check gate.
	if nc.PackageBuild(dir, "lib.js") {
		t.Fatal("node --check must reject a syntax error (PackageBuild)")
	}
}

func TestJSBothGatesAcceptValidModule(t *testing.T) {
	cfg := jsToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeJSModule(t, "function add(a, b) { return a + b; }\nmodule.exports = { add };\n")
	if !nc.FastCompile(dir, "lib.js") {
		t.Fatal("node --check must accept a valid module")
	}
	if !nc.PackageBuild(dir, "lib.js") {
		t.Fatal("node --check must accept a valid module (PackageBuild)")
	}
}
