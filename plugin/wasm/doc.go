/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package wasm provides a self-developed, deterministic WebAssembly bytecode
// (WAT) encoder/decoder toolchain.
//
// This package is NOT a plugin runtime: it does not load, instantiate or
// execute guest modules. It is retained solely as an offline tool that
// translates WAT text into deterministic WASM bytes, together with its
// byte-level unit tests. Abate plugins are native Go values hosted by
// plugin/host (the frozen WAS-HOST decision, see docs/plan/02-wasm-runtime.md);
// guest determinism is supplied by internal/determinism and by the sim/clock
// and sim/rand injectors.
//
// No third-party runtime is imported here. The toolchain is pure Go and
// deterministic: identical WAT input always yields identical bytes.
package wasm
