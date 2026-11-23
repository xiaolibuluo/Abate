/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package mutcore is the language-agnostic mutation session core of the
// real-test mutation backends (KRN-MUT).
//
// The package factors out everything the per-language backends share:
// the token/test bookkeeping, the outcome cache, the LRU pool of staged
// mutant workspaces and the two kernel port adapters. A language backend
// supplies only its toolchain knowledge through the Driver interface:
// how sources lex into mutable tokens, how tests are discovered, how
// per-test coverage is measured on the pristine tree, and how one staged
// mutant is compiled and executed.
//
// A mutant is killed by the real toolchain: its rewritten source fails to
// build, or at least one covering test fails against the mutated tree. A
// mutant survives only when every covering test passes. No outcome is
// ever approximated, sampled or inferred statically.
package mutcore
