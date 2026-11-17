/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import "errors"

// ErrDrainTimeout reports that in-flight plugin calls did not drain within the
// wall-clock DrainTimeout. The new instance remains registered and serving; the
// old instance stays superseded and is deliberately not closed while calls may
// still be executing.
var ErrDrainTimeout = errors.New("plugin: hot-swap drain timeout")

// ErrSwapInvalidInstance reports a nil or identical pair of instances.
var ErrSwapInvalidInstance = errors.New("plugin: hot-swap requires two distinct live instances")

// ErrSwapIdentityMismatch reports that the two manifests do not describe the
// same plugin id and type.
var ErrSwapIdentityMismatch = errors.New("plugin: hot-swap manifest identity mismatch")

// ErrSwapNotAllowed reports a version step that violates the fail-closed
// Allowed policy, or a replacement instance that is already superseded.
var ErrSwapNotAllowed = errors.New("plugin: hot-swap version transition not allowed")

// ErrSwapNoRegistry reports that neither instance carries a host registry, so
// the atomic SWITCH cannot be performed.
var ErrSwapNoRegistry = errors.New("plugin: hot-swap instance has no registry")
