/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package signature implements a self-developed Ed25519 plugin signature
// verification chain for the Abate L1 plugin runtime (WAS-SIG module).
//
// Design note (frozen architectural decision, see docs/plan/00-MASTER-PLAN.md
// and docs/plan/02-wasm-runtime.md WAS-SIG section): this package replaces the
// originally planned openpgp/cosign dual-channel verification with a fully
// in-house Ed25519 verification chain built only on the Go standard library
// (crypto/ed25519, crypto/sha256, encoding/base64, encoding/json). No third
// party crypto or WASM library is imported or vendored.
//
// Iron law of this chain: "verification failure never falls back to unsigned".
// A plugin whose signature cannot be proven valid is rejected outright; there
// is no code path that treats a failed verification as "no signature needed".
//
// Verification performs three independent checks, each with its own sentinel:
//
//  1. Trust anchor resolution: the manifest TrustAnchor must resolve to a known
//     public key in the TrustStore (ErrTrustAnchorUnknown).
//  2. Ed25519 signature check over the signable digest (ErrSignatureInvalid).
//  3. SHA256 content-hash cross-check: the manifest SHA256 field must equal the
//     signable digest (ErrManifestHashMismatch).
//
// Determinism: CanonicalBytes serializes fields in a fixed order with
// length-prefixed strings and big-endian integers, so two equal manifests
// always produce byte-identical canonical bytes. The verification logic avoids
// map iteration order leaks, global mutable state, time.Now, and math/rand.
package signature
