/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package abi ports a minimal, deterministic subset of the Linux syscall ABI
// and host-architecture parameters from gVisor (Apache-2.0) for use by the
// Abate kernel and sandbox layers.
//
// Only the material actually exercised by the kernel/sandbox layers is kept:
// the amd64/arm64 syscall number-to-name tables, the Linux errno constants,
// and the page-size and address primitives. The code is rewritten in Abate
// style rather than imported from gVisor; upstream Apache-2.0 attribution is
// preserved inline at each ported fragment.
//
// Upstream source (snapshotted 2026-11-11, v0.0.0-20261111045955-a42ad63bdbc8):
//   - gvisor/pkg/sentry/syscalls/linux/linux64.go (syscall number-to-name tables)
//   - gvisor/pkg/abi/linux/errno/errno.go (errno constants)
//   - gvisor/pkg/hostarch/{hostarch,addr,sizes_util}_*.go (address/page primitives)
package abi
