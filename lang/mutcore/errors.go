/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import "errors"

// Sentinel errors shared by every language backend built on mutcore. Each
// backend wraps exactly one of these values so callers classify failures
// with errors.Is regardless of the underlying toolchain.
var (
	// ErrInvalidNode reports a kernel-facing node reference that does not
	// parse into (path, offset).
	ErrInvalidNode = errors.New("mutcore: invalid node reference")
	// ErrNoSourceFiles reports a workspace without any mutable source file.
	ErrNoSourceFiles = errors.New("mutcore: no source files discovered")
	// ErrNoTests reports a workspace whose runner discovers no test.
	ErrNoTests = errors.New("mutcore: no tests discovered")
	// ErrDiscoveryFailed reports a failure of the test discovery step.
	ErrDiscoveryFailed = errors.New("mutcore: test discovery failed")
	// ErrCoverageFailed reports a failure of the per-test coverage step.
	ErrCoverageFailed = errors.New("mutcore: coverage collection failed")
	// ErrWorkspaceFailed reports a failure while materializing a mutated
	// workspace.
	ErrWorkspaceFailed = errors.New("mutcore: workspace materialization failed")
	// ErrInvalidRoot reports an unreadable or empty workspace root.
	ErrInvalidRoot = errors.New("mutcore: invalid workspace root")
	// ErrToolchainMissing reports a backend whose required toolchain binary
	// was not found on the host.
	ErrToolchainMissing = errors.New("mutcore: toolchain binary not found")
)
