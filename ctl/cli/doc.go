/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package cli is the deterministic command-line orchestration layer of Abate.
//
// It owns no experiment state: every invocation parses its arguments, hands
// every semantic decision to the CES validator and the kernel verifier, and
// renders a report whose bytes depend only on the document, seed and flags.
// The run command can seal a SIM-REC archive with --record, and replay verifies
// that archive and compares two of them through the structured Diff. Commands
// that cannot be implemented honestly are not registered.
package cli
