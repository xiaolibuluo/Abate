/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"context"
	"encoding/hex"
	"sync"
	"testing"
)

func hexOf(sum [32]byte) string { return hex.EncodeToString(sum[:]) }

// runEventStream builds one seeded world, injects all twelve faults in order
// and returns the SHA-256 digest of the recorded injection event stream.
func runEventStream(t *testing.T, seed uint64) [32]byte {
	t.Helper()
	injector := mustInjector(t, mustWorld(t, seed))
	ctx := context.Background()
	for _, op := range injector.Registry().All() {
		if err := injector.Inject(ctx, op); err != nil {
			t.Fatalf("seed %d operator %s: %v", seed, op.ID(), err)
		}
		if err := injector.Rollback(ctx, op); err != nil {
			t.Fatalf("seed %d rollback %s: %v", seed, op.ID(), err)
		}
	}
	return injector.Digest()
}

// TestSameSeedEventStreamSHAUnique100 certifies one digest over 100 seeded runs.
func TestSameSeedEventStreamSHAUnique100(t *testing.T) {
	t.Parallel()
	unique := make(map[string]int)
	for i := 0; i < 100; i++ {
		unique[hexOf(runEventStream(t, 20261109))]++
	}
	if len(unique) != 1 {
		t.Fatalf("same-seed event streams produced %d distinct SHA-256 digests: %v",
			len(unique), unique)
	}
	for digest, count := range unique {
		t.Logf("INJ event-stream SHA-256=%s count=%d", digest, count)
	}
}

func TestConcurrentInjectorsSafe(t *testing.T) {
	t.Parallel()
	const workers = 24
	digests := make(chan [32]byte, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			injector, err := NewInjector(mustWorldConcurrent())
			if err != nil {
				return
			}
			ctx := context.Background()
			for _, op := range injector.Registry().All() {
				if err := injector.Inject(ctx, op); err != nil {
					return
				}
				if err := injector.Rollback(ctx, op); err != nil {
					return
				}
			}
			digests <- injector.Digest()
		}()
	}
	wg.Wait()
	close(digests)
	var first [32]byte
	set := false
	count := 0
	for sum := range digests {
		count++
		if !set {
			first = sum
			set = true
			continue
		}
		if sum != first {
			t.Fatal("concurrent injector digests disagree")
		}
	}
	if count != workers {
		t.Fatalf("only %d/%d concurrent injectors completed", count, workers)
	}
}

func mustWorldConcurrent() *World {
	world, err := NewWorld(20261109)
	if err != nil {
		return nil
	}
	return world
}

func TestSharedInjectorConcurrentInject(t *testing.T) {
	t.Parallel()
	injector := mustInjector(t, mustWorld(t, 5))
	op := injector.Registry().All()[0]
	const workers = 16
	errs := make(chan error, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs <- injector.Inject(context.Background(), op)
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("shared injector concurrent Inject failed: %v", err)
		}
	}
	if len(injector.Frames()) != workers*4 {
		t.Fatalf("shared injector recorded %d frames, want %d",
			len(injector.Frames()), workers*4)
	}
}

func TestFramesAreDefensiveCopies(t *testing.T) {
	t.Parallel()
	injector := mustInjector(t, mustWorld(t, 8))
	op := injector.Registry().All()[0]
	if err := injector.Inject(context.Background(), op); err != nil {
		t.Fatalf("inject: %v", err)
	}
	before := injector.Digest()
	frames := injector.Frames()
	frames[0].Payload[0] ^= 0xff
	frames[0].Kind = "TAMPERED"
	if injector.Digest() != before {
		t.Fatal("Frames() leaked the internal frame slice")
	}
	if injector.Frames()[0].Kind != "FAULT" {
		t.Fatal("Frames() returned a tampered kind")
	}
}
