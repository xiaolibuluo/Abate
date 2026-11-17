/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

// NetworkPolicy is the network posture of the sandbox.
//
// 模型来源=gVisor runsc 的 --network 配置项语义（none / sandbox / host），
// Abate 使用自己的命名空间承载该语义，不调用 runsc 进程。
type NetworkPolicy string

// The known network postures.
const (
	// NetworkNone keeps the loopback stack only, the default posture.
	NetworkNone NetworkPolicy = "none"

	// NetworkSandbox uses the isolated user space network stack.
	NetworkSandbox NetworkPolicy = "sandbox"

	// NetworkHost forwards the network calls to the host network.
	NetworkHost NetworkPolicy = "host"
)

// String returns the canonical name of the network posture.
func (n NetworkPolicy) String() string { return string(n) }

// Known reports whether the network posture is one of the known values.
func (n NetworkPolicy) Known() bool {
	switch n {
	case NetworkNone, NetworkSandbox, NetworkHost:
		return true
	default:
		return false
	}
}

// Isolated reports whether the posture keeps the guest off the host network.
func (n NetworkPolicy) Isolated() bool {
	return n == NetworkNone || n == NetworkSandbox
}
