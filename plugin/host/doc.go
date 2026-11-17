/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package host is the native Go plugin runtime for the Abate L1 plugin layer.
//
// Architecture deviation (frozen decision, see docs/plan/02-wasm-runtime.md
// WAS-HOST): this package does NOT vendor or embed a third-party WASM runtime.
// Instead it provides a native plugin host where plugins are ordinary Go values
// that satisfy one of the Plugin interfaces (OperatorPlugin, ReporterPlugin,
// AdapterPlugin). The original plan called for a third-party WASM runtime; the
// decision was to self-develop the host on the Go standard library so the determinism, air-gap,
// and "no third-party WASM/crypto" iron laws are guaranteed by construction.
//
// The host-function backend (virtual clock reads and deterministic randomness)
// that the plan attributed to WASM host functions is realized directly with
// SIM-CLK (sim/clock) and SIM-RNG (sim/rand), injected into the call context by
// Host.Call. Verification, capability granting, instantiation, and registry
// bookkeeping follow the Load chain documented on Host.Load.
package host
