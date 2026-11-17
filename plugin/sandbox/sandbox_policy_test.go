/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"errors"
	"testing"
)

// TestValidateRejectsUngrantedCapabilities covers the capability gate.
func TestValidateRejectsUngrantedCapabilities(t *testing.T) {
	t.Parallel()
	config := DefaultConfig("abate/demo")
	config.Capabilities = CapFS | CapNet
	err := Validate(config, CapFS)
	if !errors.Is(err, ErrCapabilityDenied) {
		t.Fatalf("error = %v, want ErrCapabilityDenied", err)
	}
	if err := Validate(config, CapFS|CapNet); err != nil {
		t.Fatalf("granted capabilities were rejected: %v", err)
	}
}

// TestValidateRejectsEscapingMounts covers the path gate.
func TestValidateRejectsEscapingMounts(t *testing.T) {
	t.Parallel()
	config := DefaultConfig("abate/demo")
	config.RootFS.Path = "/var/lib/abate/sandbox/rootfs"
	config.Filesystem.ReadOnlyPaths = []string{"/var/lib/abate/sandbox/rootfs/etc"}
	if err := Validate(config, Capability(0)); err != nil {
		t.Fatalf("a mount inside the root was rejected: %v", err)
	}
	config.Filesystem.ReadOnlyPaths = []string{"/var/lib/abate/sandbox/other/etc"}
	if err := Validate(config, Capability(0)); !errors.Is(err, ErrPathEscape) {
		t.Fatalf("error = %v, want ErrPathEscape", err)
	}
	config.Filesystem.ReadOnlyPaths = []string{"etc"}
	if err := Validate(config, Capability(0)); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("error = %v, want ErrInvalidPath", err)
	}
	config.Filesystem.ReadOnlyPaths = nil
	config.RootFS.Path = "relative/rootfs"
	if err := Validate(config, Capability(0)); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("error = %v, want ErrInvalidPath", err)
	}
}

// TestValidateRejectsInvalidShapes covers the structural gates.
func TestValidateRejectsInvalidShapes(t *testing.T) {
	t.Parallel()
	base := DefaultConfig("abate/demo")
	cases := []struct {
		name   string
		mutate func(*Config)
		want   error
	}{
		{"no-plugin", func(c *Config) { c.PluginID = "" }, ErrInvalidManifest},
		{"unknown-platform", func(c *Config) { c.Platform = "qemu" }, ErrUnknownPlatform},
		{"unknown-network", func(c *Config) { c.Network = "bridge" }, ErrUnknownNetwork},
		{"zero-memory", func(c *Config) { c.MemoryLimitMB = 0 }, ErrInvalidMemory},
		{"huge-memory", func(c *Config) { c.MemoryLimitMB = MaxMemoryLimitMB + 1 }, ErrMemoryLimit},
	}
	for _, tc := range cases {
		config := base
		tc.mutate(&config)
		if err := Validate(config, Capability(0)); !errors.Is(err, tc.want) {
			t.Fatalf("%s: error = %v, want %v", tc.name, err, tc.want)
		}
	}
}

// TestValidateRejectsPolicyContradictions covers the posture gates.
func TestValidateRejectsPolicyContradictions(t *testing.T) {
	t.Parallel()
	base := DefaultConfig("abate/demo")
	base.RootFS.Path = "/var/lib/abate/sandbox/rootfs"
	cases := []struct {
		name    string
		mutate  func(*Config)
		granted Capability
		want    error
	}{
		{
			name: "network-without-capability",
			mutate: func(c *Config) {
				c.Network = NetworkSandbox
				c.NetworkAllowlist = []string{"a.example:443"}
			},
			want: ErrPolicyMismatch,
		},
		{
			name: "allowlist-without-posture",
			mutate: func(c *Config) {
				c.NetworkAllowlist = []string{"a.example:443"}
			},
			want: ErrPolicyMismatch,
		},
		{
			name: "write-without-capability",
			mutate: func(c *Config) {
				c.Filesystem.WritablePaths = []string{"/var/lib/abate/sandbox/rootfs/tmp"}
			},
			want: ErrPolicyMismatch,
		},
		{
			name: "writable-rootfs-without-capability",
			mutate: func(c *Config) {
				c.RootFS.ReadOnly = false
			},
			want: ErrPolicyMismatch,
		},
		{
			name: "env-without-capability",
			mutate: func(c *Config) {
				c.EnvAllowlist = []string{"ZONE"}
			},
			want: ErrPolicyMismatch,
		},
		{
			name: "clock-without-virtual-clock",
			mutate: func(c *Config) {
				c.Capabilities = CapClock
				c.ClockVirtual = false
			},
			granted: CapClock,
			want:    ErrPolicyMismatch,
		},
	}
	for _, tc := range cases {
		config := base
		tc.mutate(&config)
		if err := Validate(config, tc.granted); !errors.Is(err, tc.want) {
			t.Fatalf("%s: error = %v, want %v", tc.name, err, tc.want)
		}
	}
}

// TestValidateAcceptsAFullyGrantedConfiguration covers the happy path.
func TestValidateAcceptsAFullyGrantedConfiguration(t *testing.T) {
	t.Parallel()
	granted := CapFS | CapNet | CapClock | CapEnv
	config := DefaultConfig("abate/demo")
	config.Capabilities = granted
	config.Network = NetworkSandbox
	config.NetworkAllowlist = []string{"a.example:443"}
	config.Filesystem.ReadOnlyPaths = []string{"/var/lib/abate/sandbox/rootfs/etc"}
	config.Filesystem.WritablePaths = []string{"/var/lib/abate/sandbox/rootfs/tmp"}
	config.EnvAllowlist = []string{"ZONE"}
	if err := Validate(config, granted); err != nil {
		t.Fatalf("a fully granted configuration was rejected: %v", err)
	}
}
