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

// TestValidationMatrix pins every rejection path of the policy gate.
//
// The matrix is the mutation style coverage of this package: one row per
// boundary of the policy, each row asserting the sentinel error that the row
// must raise.
func TestValidationMatrix(t *testing.T) {
	t.Parallel()
	base := DefaultConfig("abate/demo")
	cases := []struct {
		name    string
		config  Config
		granted Capability
		want    error
	}{
		{"default", base, Capability(0), nil},
		{"no-plugin-id", Config{}, Capability(0), ErrInvalidManifest},
		{
			name:    "unknown-platform",
			config:  withPlatform(base, "qemu"),
			granted: Capability(0),
			want:    ErrUnknownPlatform,
		},
		{
			name:    "unknown-network",
			config:  withNetwork(base, "bridge"),
			granted: Capability(0),
			want:    ErrUnknownNetwork,
		},
		{
			name:    "zero-memory",
			config:  withMemory(base, 0),
			granted: Capability(0),
			want:    ErrInvalidMemory,
		},
		{
			name:    "negative-memory",
			config:  withMemory(base, -1),
			granted: Capability(0),
			want:    ErrInvalidMemory,
		},
		{
			name:    "memory-above-ceiling",
			config:  withMemory(base, MaxMemoryLimitMB+1),
			granted: Capability(0),
			want:    ErrMemoryLimit,
		},
		{
			name:    "memory-at-ceiling",
			config:  withMemory(base, MaxMemoryLimitMB),
			granted: Capability(0),
			want:    nil,
		},
		{
			name:    "capability-not-granted",
			config:  withCapabilities(base, CapNet),
			granted: Capability(0),
			want:    ErrCapabilityDenied,
		},
		{
			name:    "capability-granted",
			config:  withCapabilities(base, CapNet),
			granted: CapNet,
			want:    nil,
		},
		{
			name:    "relative-rootfs",
			config:  withRootFS(base, "rootfs"),
			granted: Capability(0),
			want:    ErrInvalidPath,
		},
		{
			name:    "mount-outside-root",
			config:  withReadOnly(base, "/etc/passwd"),
			granted: Capability(0),
			want:    ErrPathEscape,
		},
		{
			name:    "mount-inside-root",
			config:  withReadOnly(base, "/var/lib/abate/sandbox/rootfs/etc"),
			granted: Capability(0),
			want:    nil,
		},
		{
			name:    "mount-equal-to-root",
			config:  withReadOnly(base, "/var/lib/abate/sandbox/rootfs"),
			granted: Capability(0),
			want:    nil,
		},
	}
	for _, tc := range cases {
		err := Validate(tc.config, tc.granted)
		if tc.want == nil {
			if err != nil {
				t.Fatalf("%s: unexpected error %v", tc.name, err)
			}
			continue
		}
		if !errors.Is(err, tc.want) {
			t.Fatalf("%s: error = %v, want %v", tc.name, err, tc.want)
		}
	}
}

// withPlatform returns a copy of one configuration with another platform.
func withPlatform(config Config, platform Platform) Config {
	config.Platform = platform
	return config
}

// withNetwork returns a copy of one configuration with another posture.
func withNetwork(config Config, network NetworkPolicy) Config {
	config.Network = network
	return config
}

// withMemory returns a copy of one configuration with another memory ceiling.
func withMemory(config Config, limit int) Config {
	config.MemoryLimitMB = limit
	return config
}

// withCapabilities returns a copy of one configuration with other capabilities.
func withCapabilities(config Config, caps Capability) Config {
	config.Capabilities = caps
	return config
}

// withRootFS returns a copy of one configuration with another rootfs path.
func withRootFS(config Config, path string) Config {
	config.RootFS.Path = path
	return config
}

// withReadOnly returns a copy of one configuration with another read only mount.
func withReadOnly(config Config, path string) Config {
	config.Filesystem.ReadOnlyPaths = []string{path}
	return config
}

// TestBuildKeepsTheDeclaredResources proves the derivation keeps the inputs.
func TestBuildKeepsTheDeclaredResources(t *testing.T) {
	t.Parallel()
	config, err := Build(Manifest{
		PluginID:         "abate/demo",
		Platform:         PlatformKVM,
		Network:          NetworkSandbox,
		RootFSPath:       "/srv/rootfs",
		RootFSWritable:   true,
		RootFSDigest:     "sha256:abc",
		Capabilities:     CapFS | CapNet | CapEnv,
		MemoryLimitMB:    512,
		NetworkAllowlist: []string{"a.example:443"},
		EnvAllowlist:     []string{"ZONE"},
		WritablePaths:    []string{"/srv/rootfs/tmp"},
	})
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if config.Platform != PlatformKVM || config.Network != NetworkSandbox {
		t.Fatalf("config = %+v", config)
	}
	if config.RootFS.Path != "/srv/rootfs" || config.RootFS.ReadOnly {
		t.Fatalf("rootfs = %+v", config.RootFS)
	}
	if config.RootFS.Digest != "sha256:abc" {
		t.Fatalf("rootfs digest = %q", config.RootFS.Digest)
	}
	if config.MemoryLimitMB != 512 {
		t.Fatalf("memory = %d", config.MemoryLimitMB)
	}
	if err := Validate(config, CapFS|CapNet|CapEnv); err != nil {
		t.Fatalf("a granted configuration was rejected: %v", err)
	}
}
