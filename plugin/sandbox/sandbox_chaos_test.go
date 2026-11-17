/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"sync"
	"testing"
)

// TestBuildAndMarshalAreStableOver100Runs certifies that one hundred runs of
// the declaration path produce exactly one configuration document.
func TestBuildAndMarshalAreStableOver100Runs(t *testing.T) {
	t.Parallel()
	manifest := Manifest{
		PluginID:         "abate/demo",
		Capabilities:     CapFS | CapClock,
		NetworkAllowlist: []string{"b.example:443", "a.example:443"},
		EnvAllowlist:     []string{"ZONE", "ALPHA"},
		ReadOnlyPaths:    []string{"/var/lib/abate/sandbox/rootfs/etc"},
	}
	unique := make(map[string]int)
	for i := 0; i < 100; i++ {
		config, err := Build(manifest)
		if err != nil {
			t.Fatalf("iteration %d: build failed: %v", i, err)
		}
		payload, err := Marshal(config)
		if err != nil {
			t.Fatalf("iteration %d: marshal failed: %v", i, err)
		}
		unique[string(payload)]++
	}
	if len(unique) != 1 {
		t.Fatalf("the hundred iterations produced %d distinct documents", len(unique))
	}
	for payload, count := range unique {
		t.Logf("SANDBOX config bytes=%d count=%d", len(payload), count)
	}
}

// TestConcurrentSandboxUseIsSafe proves the layer holds no shared state.
func TestConcurrentSandboxUseIsSafe(t *testing.T) {
	t.Parallel()
	manifest := Manifest{
		PluginID:     "abate/demo",
		Capabilities: CapFS,
		WritablePaths: []string{
			"/var/lib/abate/sandbox/rootfs/tmp",
		},
	}
	const workers = 32
	results := make(chan string, workers)
	errs := make(chan error, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			config, err := Build(manifest)
			if err != nil {
				errs <- err
				return
			}
			if err := Validate(config, CapFS); err != nil {
				errs <- err
				return
			}
			payload, err := Marshal(config)
			if err != nil {
				errs <- err
				return
			}
			results <- string(payload)
		}()
	}
	wg.Wait()
	close(results)
	close(errs)
	for err := range errs {
		t.Fatalf("concurrent sandbox use failed: %v", err)
	}
	seen := ""
	for payload := range results {
		if seen == "" {
			seen = payload
			continue
		}
		if payload != seen {
			t.Fatal("concurrent runs disagree")
		}
	}
	if seen == "" {
		t.Fatal("no concurrent run produced a configuration")
	}
}

// TestCanonicalDoesNotMutateTheSource proves the canonical form is a copy.
func TestCanonicalDoesNotMutateTheSource(t *testing.T) {
	t.Parallel()
	config := DefaultConfig("abate/demo")
	config.NetworkAllowlist = []string{"b.example:443", "a.example:443"}
	canonical := config.Canonical()
	if config.NetworkAllowlist[0] != "b.example:443" {
		t.Fatal("Canonical() reordered the source configuration")
	}
	if canonical.NetworkAllowlist[0] != "a.example:443" {
		t.Fatalf("canonical allow list = %v", canonical.NetworkAllowlist)
	}
	canonical.NetworkAllowlist[0] = "tampered"
	if config.NetworkAllowlist[0] == "tampered" {
		t.Fatal("Canonical() leaked its backing array")
	}
}
