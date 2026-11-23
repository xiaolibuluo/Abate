/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import (
	"container/list"
	"testing"
)

func TestEvictLockedKeepsPoolUnderCap(t *testing.T) {
	c := &core{staged: make(map[[32]byte]*stagedEntry), stageLRU: list.New()}
	for i := 0; i < stagePoolCap+5; i++ {
		var fp [32]byte
		fp[0] = byte(i)
		entry := &stagedEntry{dir: "/nonexistent", compiles: true, refs: 1}
		entry.elem = c.stageLRU.PushFront(fp)
		c.staged[fp] = entry
	}
	// All entries have refs>0, so evictLocked must not remove any.
	c.evictLocked()
	if len(c.staged) <= stagePoolCap {
		t.Fatal("evictLocked must not evict entries with refs>0")
	}
	now := len(c.staged)

	// Drop all refs, then one eviction cycle must remove at least one entry.
	for _, entry := range c.staged {
		entry.refs = 0
	}
	c.evictLocked()
	if len(c.staged) >= now {
		t.Fatal("evictLocked must evict entries when refs==0")
	}
}

func TestEvictLockedLRUOrder(t *testing.T) {
	c := &core{staged: make(map[[32]byte]*stagedEntry), stageLRU: list.New()}
	// Insert 3 entries above cap.
	for i := 0; i < stagePoolCap+3; i++ {
		var fp [32]byte
		fp[0] = byte(i)
		entry := &stagedEntry{dir: "/nonexistent", compiles: true, refs: 0}
		entry.elem = c.stageLRU.PushFront(fp)
		c.staged[fp] = entry
	}
	c.evictLocked()
	if len(c.staged) != stagePoolCap {
		t.Fatalf("after eviction pool must be at cap %d, got %d", stagePoolCap, len(c.staged))
	}
	// The back (LRU) entry must be gone, the front (MRU) entries preserved.
	back, _ := c.stageLRU.Back().Value.([32]byte)
	if back[0] == byte(0) {
		t.Fatal("oldest entry (fp[0]=0) must have been evicted")
	}
}

func TestStagedReleaseDecrementsRefs(t *testing.T) {
	c := &core{staged: make(map[[32]byte]*stagedEntry)}
	var fp [32]byte
	fp[0] = 42
	entry := &stagedEntry{dir: "/tmp/abate-test", compiles: true, refs: 3}
	c.staged[fp] = entry
	c.stagedRelease(fp)
	if entry.refs != 2 {
		t.Fatalf("stagedRelease must decrement refs, got %d", entry.refs)
	}
	c.stagedRelease(fp)
	c.stagedRelease(fp)
	if entry.refs != 0 {
		t.Fatalf("stagedRelease must reach zero, got %d", entry.refs)
	}
	// Must not go below zero.
	c.stagedRelease(fp)
	if entry.refs != 0 {
		t.Fatalf("stagedRelease must not go below zero, got %d", entry.refs)
	}
}

func TestStagedReleaseMissingFingerprint(t *testing.T) {
	c := &core{staged: make(map[[32]byte]*stagedEntry)}
	var fp [32]byte
	fp[0] = 99
	// Must not panic on missing entry.
	c.stagedRelease(fp)
}
