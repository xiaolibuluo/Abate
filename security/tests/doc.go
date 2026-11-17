/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package tests is the executable security attack-case suite for Abate.
//
// It is the T1-T10 verification surface required by XCU-SEC: every threat in the
// threat model must have at least one negative test proving the mitigation
// intercepts the attack. Each test is named for the threat it exercises
// (TestThreatT1SignatureRejection, TestThreatT6ParseInjection, ...) and asserts
// the attack is refused, escaped, or otherwise neutralized.
//
// Threat coverage matrix (mitigation owner -> test):
//
//	T1  malicious plugin load        -> plugin/signature.Verify (TestThreatT1*)
//	T2  plugin privilege escalation  -> capability gate + audit (TestThreatT2CapabilityEscalation)
//	T3  plugin resource exhaustion   -> host result-size cap (TestThreatT3ResourceExhaustion)
//	T4  supply-chain poisoning       -> source-level merge + commit pinning [docs]
//	T5  signing-key leakage          -> offline key custody (process) [docs]
//	T6  CES document injection       -> parse depth/size caps (TestThreatT6*)
//	T7  report XSS                   -> html/template auto-escaping (TestThreatT7*)
//	T8  REST unauthorized access     -> Bearer auth + remote-bind guard (TestThreatT8*)
//	T9  secret leak into logs        -> redact middleware + rest log hooks (TestThreatT9SecretRedaction)
//	T10 archive tampering            -> frame hash chain + full SHA-256 (TestThreatT10*)
//
// The suite is deterministic and offline: it imports only in-module packages and
// the self-researched internal/yamlmini (no third-party dependencies), and it never
// reads the clock, the environment or the network (aligning with the project's
// determinism iron law).
package tests
