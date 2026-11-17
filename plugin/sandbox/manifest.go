/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"fmt"
)

// Manifest is the plugin declaration the sandbox configuration is built from.
//
// The declaration is what a plugin publishes; the configuration is what the
// host derives from it after applying the defaults and the policy checks.
type Manifest struct {
	// PluginID identifies the plugin.
	PluginID string `json:"pluginId"`

	// Platform requests a guest platform layer, empty means the default one.
	Platform Platform `json:"platform"`

	// Network requests a network posture, empty means the default one.
	Network NetworkPolicy `json:"network"`

	// RootFSPath requests a rootfs bundle path, empty means the default one.
	RootFSPath string `json:"rootfsPath"`

	// RootFSWritable asks for a writable rootfs bundle, the default is read only.
	RootFSWritable bool `json:"rootfsWritable"`

	// RootFSDigest is the expected content digest of the requested bundle.
	RootFSDigest string `json:"rootfsDigest"`

	// Capabilities are the capabilities the plugin asks for.
	Capabilities Capability `json:"capabilities"`

	// MemoryLimitMB requests a memory ceiling, zero means the default one.
	MemoryLimitMB int `json:"memoryLimitMb"`

	// NetworkAllowlist declares the permitted network endpoints.
	NetworkAllowlist []string `json:"networkAllowlist"`

	// EnvAllowlist declares the environment variables of the guest.
	EnvAllowlist []string `json:"envAllowlist"`

	// ReadOnlyPaths declares the read only mounts of the guest.
	ReadOnlyPaths []string `json:"readOnlyPaths"`

	// WritablePaths declares the read write mounts of the guest.
	WritablePaths []string `json:"writablePaths"`
}

// Build derives the sandbox configuration of one plugin declaration.
//
// The derivation is pure and deterministic: the same declaration always yields
// the same configuration, and every allow list is sorted so the serialized
// form is comparable.
func Build(manifest Manifest) (Config, error) {
	if manifest.PluginID == "" {
		return Config{}, fmt.Errorf("%w: the plugin id must not be empty",
			ErrInvalidManifest)
	}
	platform := manifest.Platform
	if platform == "" {
		platform = DefaultPlatform
	}
	if !platform.Known() {
		return Config{}, fmt.Errorf("%w: %q", ErrUnknownPlatform, platform)
	}
	network := manifest.Network
	if network == "" {
		network = DefaultNetwork
	}
	if !network.Known() {
		return Config{}, fmt.Errorf("%w: %q", ErrUnknownNetwork, network)
	}
	rootfsPath := manifest.RootFSPath
	if rootfsPath == "" {
		rootfsPath = DefaultRootFSPath
	}
	if !isAbsoluteCleanPath(rootfsPath) {
		return Config{}, fmt.Errorf("%w: rootfs path %q", ErrInvalidPath, rootfsPath)
	}
	memory := manifest.MemoryLimitMB
	if memory == 0 {
		memory = DefaultMemoryLimitMB
	}
	if memory < 0 {
		return Config{}, fmt.Errorf("%w: %d", ErrInvalidMemory, memory)
	}
	if memory > MaxMemoryLimitMB {
		return Config{}, fmt.Errorf("%w: %d > %d", ErrMemoryLimit, memory, MaxMemoryLimitMB)
	}
	config := Config{
		PluginID:      manifest.PluginID,
		Platform:      platform,
		Network:       network,
		RootFS:        RootFS{Path: rootfsPath, ReadOnly: !manifest.RootFSWritable},
		Capabilities:  manifest.Capabilities,
		MemoryLimitMB: memory,
		ClockVirtual:  true,
	}
	config.NetworkAllowlist = sortedCopy(manifest.NetworkAllowlist)
	config.EnvAllowlist = sortedCopy(manifest.EnvAllowlist)
	config.Filesystem = FilesystemPolicy{
		ReadOnlyPaths: sortedCopy(manifest.ReadOnlyPaths),
		WritablePaths: sortedCopy(manifest.WritablePaths),
	}
	if manifest.RootFSDigest != "" {
		config.RootFS.Digest = manifest.RootFSDigest
	}
	return config, nil
}
