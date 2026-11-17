/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"fmt"
	"strings"
)

// Validate checks one sandbox configuration against the granted capabilities.
//
// The checks run in a fixed order, so the same configuration always reports
// the same first violation. The posture is default deny: a capability the host
// did not grant, a mount outside the sandbox root and a memory ceiling above
// the hard limit are all rejected.
func Validate(config Config, granted Capability) error {
	if config.PluginID == "" {
		return fmt.Errorf("%w: the plugin id must not be empty", ErrInvalidManifest)
	}
	if !config.Platform.Known() {
		return fmt.Errorf("%w: %q", ErrUnknownPlatform, config.Platform)
	}
	if !config.Network.Known() {
		return fmt.Errorf("%w: %q", ErrUnknownNetwork, config.Network)
	}
	if config.MemoryLimitMB <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidMemory, config.MemoryLimitMB)
	}
	if config.MemoryLimitMB > MaxMemoryLimitMB {
		return fmt.Errorf("%w: %d > %d",
			ErrMemoryLimit, config.MemoryLimitMB, MaxMemoryLimitMB)
	}
	if missing := config.Capabilities.Missing(granted); len(missing) > 0 {
		return fmt.Errorf("%w: %s", ErrCapabilityDenied, strings.Join(missing, ","))
	}
	if !config.RootFS.ValidPath() {
		return fmt.Errorf("%w: rootfs path %q", ErrInvalidPath, config.RootFS.Path)
	}
	for _, mount := range config.Filesystem.Mounts() {
		if !isAbsoluteCleanPath(mount) {
			return fmt.Errorf("%w: mount %q", ErrInvalidPath, mount)
		}
		if !containsPath(config.RootFS.Path, mount) {
			return fmt.Errorf("%w: mount %q leaves %q",
				ErrPathEscape, mount, config.RootFS.Path)
		}
	}
	if config.Network != NetworkNone && !config.Capabilities.Has(CapNet) {
		return fmt.Errorf("%w: network %q needs the net capability",
			ErrPolicyMismatch, config.Network)
	}
	if config.Network == NetworkNone && len(config.NetworkAllowlist) > 0 {
		return fmt.Errorf("%w: the network allow list needs a network posture",
			ErrPolicyMismatch)
	}
	if config.Filesystem.NeedsWrite() && !config.Capabilities.Has(CapFS) {
		return fmt.Errorf("%w: writable mounts need the fs capability",
			ErrPolicyMismatch)
	}
	if !config.RootFS.ReadOnly && !config.Capabilities.Has(CapFS) {
		return fmt.Errorf("%w: a writable rootfs needs the fs capability",
			ErrPolicyMismatch)
	}
	if len(config.EnvAllowlist) > 0 && !config.Capabilities.Has(CapEnv) {
		return fmt.Errorf("%w: the environment allow list needs the env capability",
			ErrPolicyMismatch)
	}
	if config.Capabilities.Has(CapClock) && !config.ClockVirtual {
		return fmt.Errorf("%w: the clock capability needs the virtual clock",
			ErrPolicyMismatch)
	}
	return nil
}
