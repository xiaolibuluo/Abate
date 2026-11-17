/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

// Syscall identifies a Linux syscall number within a given architecture.
type Syscall uintptr

// AMD64Names maps a Linux amd64 syscall number to its canonical name.
//
// Adapted from gvisor/pkg/sentry/syscalls/linux/linux64.go (AMD64 table),
// Apache-2.0.
var AMD64Names = mergeSyscallMaps(amd64NamesLow, amd64NamesHigh)

// ARM64Names maps a Linux arm64 syscall number to its canonical name.
//
// Adapted from gvisor/pkg/sentry/syscalls/linux/linux64.go (ARM64 table),
// Apache-2.0.
var ARM64Names = mergeSyscallMaps(arm64NamesLow, arm64NamesHigh)

// mergeSyscallMaps concatenates the per-range upstream fragment maps into a
// single authoritative table. It panics on a duplicate number so that any
// divergence from the upstream source is caught at package init.
func mergeSyscallMaps(maps ...map[uintptr]string) map[uintptr]string {
	merged := make(map[uintptr]string)
	for _, m := range maps {
		for no, name := range m {
			if _, dup := merged[no]; dup {
				panic("abi: duplicate syscall number in upstream fragment")
			}
			merged[no] = name
		}
	}
	return merged
}

// AMD64Name returns the canonical name for an amd64 syscall number and whether
// the number is present in the ported table.
func AMD64Name(no Syscall) (string, bool) {
	name, ok := AMD64Names[uintptr(no)]
	return name, ok
}

// ARM64Name returns the canonical name for an arm64 syscall number and whether
// the number is present in the ported table.
func ARM64Name(no Syscall) (string, bool) {
	name, ok := ARM64Names[uintptr(no)]
	return name, ok
}
