/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

// Platform is the guest platform layer of the sandbox.
//
// 模型来源=gVisor runsc 的 --platform 配置项语义（ptrace / systrap / kvm），
// Abate 使用自己的命名空间承载该语义，不调用 runsc 进程。
type Platform string

// The known platform layers.
const (
	// PlatformPtrace emulates the guest with the ptrace platform.
	PlatformPtrace Platform = "ptrace"

	// PlatformSystrap emulates the guest with the systrap platform.
	PlatformSystrap Platform = "systrap"

	// PlatformKVM runs the guest on hardware virtualization.
	PlatformKVM Platform = "kvm"
)

// String returns the canonical name of the platform layer.
func (p Platform) String() string { return string(p) }

// Known reports whether the platform layer is one of the known values.
func (p Platform) Known() bool {
	switch p {
	case PlatformPtrace, PlatformSystrap, PlatformKVM:
		return true
	default:
		return false
	}
}

// HardwareAccelerated reports whether the platform layer needs a hypervisor.
func (p Platform) HardwareAccelerated() bool {
	return p == PlatformKVM
}
