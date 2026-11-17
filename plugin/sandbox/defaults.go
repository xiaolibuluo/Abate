/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

// The default posture of the Abate plugin sandbox.
//
// 模型来源=gVisor runsc 的默认配置语义（默认沙箱网络、默认 ptrace 平台层、
// 默认独占根文件系统访问），Abate 以自己的命名空间固化该默认值，默认全拒。
const (
	// DefaultPlatform is the default guest platform layer.
	DefaultPlatform = PlatformSystrap

	// DefaultNetwork is the default network posture, loopback only.
	DefaultNetwork = NetworkNone

	// DefaultMemoryLimitMB is the default memory ceiling of a guest.
	DefaultMemoryLimitMB = 256

	// MaxMemoryLimitMB is the hard memory ceiling of a guest.
	MaxMemoryLimitMB = 4096

	// DefaultRootFSPath is the default rootfs bundle path.
	DefaultRootFSPath = "/var/lib/abate/sandbox/rootfs"
)

// DefaultConfig returns the default sandbox configuration of one plugin.
//
// The default posture denies everything: no capability is requested, the guest
// stays on the loopback stack, the rootfs is read only and the memory ceiling
// is the default one.
func DefaultConfig(pluginID string) Config {
	return Config{
		PluginID:      pluginID,
		Platform:      DefaultPlatform,
		Network:       DefaultNetwork,
		RootFS:        RootFS{Path: DefaultRootFSPath, ReadOnly: true},
		Filesystem:    FilesystemPolicy{},
		Capabilities:  Capability(0),
		MemoryLimitMB: DefaultMemoryLimitMB,
		ClockVirtual:  true,
	}
}
