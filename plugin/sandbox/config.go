/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import "slices"

// Config is the deterministic sandbox configuration of one plugin.
//
// The configuration is pure data: it carries no handle, no file descriptor and
// no process reference, so it can be compared, hashed and replayed byte for
// byte.
type Config struct {
	// PluginID identifies the plugin the configuration belongs to.
	PluginID string `json:"pluginId"`

	// Platform is the guest platform layer.
	Platform Platform `json:"platform"`

	// Network is the network posture of the guest.
	Network NetworkPolicy `json:"network"`

	// NetworkAllowlist holds the permitted network endpoints.
	NetworkAllowlist []string `json:"networkAllowlist"`

	// RootFS is the root filesystem bundle of the guest.
	RootFS RootFS `json:"rootfs"`

	// Filesystem is the filesystem allow list of the guest.
	Filesystem FilesystemPolicy `json:"filesystem"`

	// Capabilities holds the capabilities the sandbox requests.
	Capabilities Capability `json:"capabilities"`

	// MemoryLimitMB is the memory ceiling of the guest in MiB.
	MemoryLimitMB int `json:"memoryLimitMb"`

	// EnvAllowlist holds the environment variables passed to the guest.
	EnvAllowlist []string `json:"envAllowlist"`

	// ClockVirtual marks the clock as served by the virtual clock.
	ClockVirtual bool `json:"clockVirtual"`
}

// Canonical returns the configuration with every allow list sorted.
//
// Two configurations that declare the same resources in a different order
// therefore share one canonical form, which keeps the serialized comparison
// meaningful.
func (c Config) Canonical() Config {
	out := c
	out.NetworkAllowlist = sortedCopy(c.NetworkAllowlist)
	out.EnvAllowlist = sortedCopy(c.EnvAllowlist)
	out.Filesystem.ReadOnlyPaths = sortedCopy(c.Filesystem.ReadOnlyPaths)
	out.Filesystem.WritablePaths = sortedCopy(c.Filesystem.WritablePaths)
	return out
}

// sortedCopy returns a sorted copy of one path list.
func sortedCopy(in []string) []string {
	if in == nil {
		return nil
	}
	out := append([]string(nil), in...)
	slices.Sort(out)
	return out
}
