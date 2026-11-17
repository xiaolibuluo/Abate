/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"errors"
	"reflect"
	"testing"
)

// TestDefaultConfigDeniesEverything pins the default deny posture.
func TestDefaultConfigDeniesEverything(t *testing.T) {
	t.Parallel()
	config := DefaultConfig("abate/demo")
	if config.PluginID != "abate/demo" {
		t.Fatalf("plugin id = %q", config.PluginID)
	}
	if config.Platform != DefaultPlatform {
		t.Fatalf("platform = %q, want %q", config.Platform, DefaultPlatform)
	}
	if config.Network != DefaultNetwork {
		t.Fatalf("network = %q, want %q", config.Network, DefaultNetwork)
	}
	if !config.Capabilities.IsZero() {
		t.Fatalf("the default configuration requests %q", config.Capabilities)
	}
	if config.MemoryLimitMB != DefaultMemoryLimitMB {
		t.Fatalf("memory = %d, want %d", config.MemoryLimitMB, DefaultMemoryLimitMB)
	}
	if !config.RootFS.ReadOnly {
		t.Fatal("the default rootfs must be read only")
	}
	if !config.ClockVirtual {
		t.Fatal("the default clock must be virtual")
	}
	if err := Validate(config, Capability(0)); err != nil {
		t.Fatalf("the default configuration was rejected: %v", err)
	}
}

// TestBuildAppliesTheDefaults covers the declaration to configuration path.
func TestBuildAppliesTheDefaults(t *testing.T) {
	t.Parallel()
	config, err := Build(Manifest{PluginID: "abate/demo"})
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	want := DefaultConfig("abate/demo")
	if !reflect.DeepEqual(config, want) {
		t.Fatalf("config = %+v, want %+v", config, want)
	}
}

// TestBuildSortsTheAllowLists proves the derived form is comparable.
func TestBuildSortsTheAllowLists(t *testing.T) {
	t.Parallel()
	config, err := Build(Manifest{
		PluginID:         "abate/demo",
		NetworkAllowlist: []string{"b.example:443", "a.example:443"},
		EnvAllowlist:     []string{"ZONE", "ALPHA"},
		ReadOnlyPaths:    []string{"/var/lib/abate/sandbox/rootfs/etc"},
	})
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	if config.NetworkAllowlist[0] != "a.example:443" {
		t.Fatalf("network allow list = %v", config.NetworkAllowlist)
	}
	if config.EnvAllowlist[0] != "ALPHA" {
		t.Fatalf("environment allow list = %v", config.EnvAllowlist)
	}
	if config.Filesystem.ReadOnlyPaths[0] != "/var/lib/abate/sandbox/rootfs/etc" {
		t.Fatalf("read only paths = %v", config.Filesystem.ReadOnlyPaths)
	}
}

// TestBuildRejectsInvalidDeclarations covers the build failure branches.
func TestBuildRejectsInvalidDeclarations(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		manifest Manifest
		want     error
	}{
		{"no-plugin", Manifest{}, ErrInvalidManifest},
		{"unknown-platform", Manifest{PluginID: "p", Platform: "qemu"}, ErrUnknownPlatform},
		{"unknown-network", Manifest{PluginID: "p", Network: "bridge"}, ErrUnknownNetwork},
		{"relative-rootfs", Manifest{PluginID: "p", RootFSPath: "rootfs"}, ErrInvalidPath},
		{"dirty-rootfs", Manifest{PluginID: "p", RootFSPath: "/a/../b"}, ErrInvalidPath},
		{"negative-memory", Manifest{PluginID: "p", MemoryLimitMB: -1}, ErrInvalidMemory},
		{"huge-memory", Manifest{PluginID: "p", MemoryLimitMB: MaxMemoryLimitMB + 1}, ErrMemoryLimit},
	}
	for _, tc := range cases {
		if _, err := Build(tc.manifest); !errors.Is(err, tc.want) {
			t.Fatalf("%s: error = %v, want %v", tc.name, err, tc.want)
		}
	}
}

// TestMarshalIsByteStable proves the deterministic serialization.
func TestMarshalIsByteStable(t *testing.T) {
	t.Parallel()
	config := DefaultConfig("abate/demo")
	config.NetworkAllowlist = []string{"b.example:443", "a.example:443"}
	first, err := Marshal(config)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	second, err := Marshal(config)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(first) != string(second) {
		t.Fatal("two encodings of one configuration differ")
	}
	other := config
	other.NetworkAllowlist = []string{"a.example:443", "b.example:443"}
	third, err := Marshal(other)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(first) != string(third) {
		t.Fatal("the canonical form depends on the allow list order")
	}
	indented, err := MarshalIndent(config)
	if err != nil {
		t.Fatalf("marshal indent: %v", err)
	}
	if len(indented) <= len(first) {
		t.Fatal("the indented form is not larger than the compact form")
	}
}

// TestJSONRoundTrip keeps every field of the configuration.
func TestJSONRoundTrip(t *testing.T) {
	t.Parallel()
	config := DefaultConfig("abate/demo")
	config.Capabilities = CapFS | CapEnv
	config.Network = NetworkSandbox
	config.NetworkAllowlist = []string{"a.example:443"}
	config.Filesystem.WritablePaths = []string{"/var/lib/abate/sandbox/rootfs/tmp"}
	config.EnvAllowlist = []string{"ZONE"}
	data, err := Marshal(config)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	decoded, err := Unmarshal(data)
	if err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !reflect.DeepEqual(decoded, config.Canonical()) {
		t.Fatalf("decoded = %+v, want %+v", decoded, config.Canonical())
	}
	if _, err := Unmarshal([]byte("{")); err == nil {
		t.Fatal("unmarshal accepted a broken document")
	}
}

// TestCapabilityHelpers covers the capability bit operations.
func TestCapabilityHelpers(t *testing.T) {
	t.Parallel()
	var caps Capability
	if !caps.IsZero() {
		t.Fatal("the zero value grants a capability")
	}
	caps.Grant(CapFS)
	caps.Grant(CapFS)
	if !caps.Has(CapFS) || caps.Has(CapNet) {
		t.Fatalf("capability mask = %q", caps)
	}
	if caps.String() != "fs" {
		t.Fatalf("capability string = %q", caps.String())
	}
	caps.Grant(CapNet)
	if got := caps.Names(); len(got) != 2 || got[0] != "fs" || got[1] != "net" {
		t.Fatalf("capability names = %v", got)
	}
	caps.Revoke(CapFS)
	caps.Revoke(CapFS)
	if caps.Has(CapFS) {
		t.Fatal("revoke did not remove the capability")
	}
	missing := (CapFS | CapClock).Missing(CapFS)
	if len(missing) != 1 || missing[0] != "clock" {
		t.Fatalf("missing capabilities = %v", missing)
	}
}

// TestPlatformAndNetworkHelpers covers the enum helpers.
func TestPlatformAndNetworkHelpers(t *testing.T) {
	t.Parallel()
	if !PlatformSystrap.Known() || Platform("qemu").Known() {
		t.Fatal("the platform knowledge check is wrong")
	}
	if PlatformSystrap.String() != "systrap" {
		t.Fatalf("platform name = %q", PlatformSystrap.String())
	}
	if !PlatformKVM.HardwareAccelerated() || PlatformPtrace.HardwareAccelerated() {
		t.Fatal("the hardware acceleration check is wrong")
	}
	if !NetworkNone.Known() || NetworkPolicy("bridge").Known() {
		t.Fatal("the network knowledge check is wrong")
	}
	if !NetworkNone.Isolated() || NetworkHost.Isolated() {
		t.Fatal("the isolation check is wrong")
	}
	if NetworkSandbox.String() != "sandbox" {
		t.Fatalf("network name = %q", NetworkSandbox.String())
	}
}

// TestFilesystemHelpers covers the filesystem policy helpers.
func TestFilesystemHelpers(t *testing.T) {
	t.Parallel()
	policy := FilesystemPolicy{
		ReadOnlyPaths: []string{"/root/etc"},
		WritablePaths: []string{"/root/tmp"},
	}
	if !policy.NeedsWrite() {
		t.Fatal("a writable mount must request write access")
	}
	if (FilesystemPolicy{}).NeedsWrite() {
		t.Fatal("an empty policy must not request write access")
	}
	mounts := policy.Mounts()
	if len(mounts) != 2 || mounts[0] != "/root/etc" || mounts[1] != "/root/tmp" {
		t.Fatalf("mounts = %v", mounts)
	}
	if len((RootFS{}).Mounts()) != 0 {
		t.Fatal("an empty rootfs must not expose a mount")
	}
}
