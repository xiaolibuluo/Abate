/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package capability implements the WAS-CAP capability sandbox for the Abate L1
// plugin runtime.
//
// Security model (see docs/plan/02-wasm-runtime.md WAS-CAP section):
//
//   - Default deny: a freshly loaded plugin has zero capabilities.
//   - Explicit grant: capabilities are granted exactly once at load time from the
//     plugin manifest declaration; plugins can never escalate at runtime.
//   - Convergence: only the five known capability names are accepted; an unknown
//     name rejects the whole declaration (load denied).
//   - Interception: every host-function call is gated by Check/CheckMany before
//     execution, returning ErrCapabilityDenied on missing authority.
//   - Audit: every grant, denial, and use is recorded in an AuditLog with a
//     virtual monotonic counter (never time.Now) for deterministic forensics.
//
// The package depends only on the Go standard library and introduces no global
// mutable state. Per the project determinism law, audit recording is provided by
// the *Audit variants of the gate functions; the frozen Check/CheckMany are pure
// gates so they can be called without any shared log.
package capability
