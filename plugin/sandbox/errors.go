/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import "errors"

// Sentinel errors of the sandbox configuration layer.
//
// Callers must test them with errors.Is; the wrapped errors carry the offending
// value so a rejection can be traced without re-parsing the configuration.
var (
	// ErrInvalidManifest reports a malformed plugin declaration.
	ErrInvalidManifest = errors.New("sandbox: invalid plugin manifest")

	// ErrUnknownPlatform reports a platform outside the known set.
	ErrUnknownPlatform = errors.New("sandbox: unknown sandbox platform")

	// ErrUnknownNetwork reports a network posture outside the known set.
	ErrUnknownNetwork = errors.New("sandbox: unknown network policy")

	// ErrCapabilityDenied reports a capability the host did not grant.
	ErrCapabilityDenied = errors.New("sandbox: capability denied")

	// ErrInvalidPath reports a path that is not an absolute clean path.
	ErrInvalidPath = errors.New("sandbox: invalid sandbox path")

	// ErrPathEscape reports a mount that leaves the sandbox root.
	ErrPathEscape = errors.New("sandbox: path escapes the sandbox root")

	// ErrInvalidMemory reports a memory limit outside the valid range.
	ErrInvalidMemory = errors.New("sandbox: invalid memory limit")

	// ErrMemoryLimit reports a memory limit above the hard ceiling.
	ErrMemoryLimit = errors.New("sandbox: memory limit exceeds the ceiling")

	// ErrPolicyMismatch reports a configuration that contradicts its own
	// capability set.
	ErrPolicyMismatch = errors.New("sandbox: policy contradicts the capability set")
)
