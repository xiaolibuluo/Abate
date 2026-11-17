/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"context"
	"sync"
	"testing"
)

// TestRecorderStartIsInert proves Start leaves the recorder untouched: it adds
// no frames or sources, does not advance the time bound, ignores the context
// state, and a following Record/Stop still yields a valid one-frame archive.
func TestRecorderStartIsInert(t *testing.T) {
	t.Parallel()
	rec := NewRecorder(testFingerprint())
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	rec.Start(ctx)
	rec.Start(context.Background())
	if len(rec.frames) != 0 || len(rec.sources) != 0 || rec.maxAt != 0 {
		t.Fatalf("Start mutated recorder state: frames=%d sources=%d maxAt=%d",
			len(rec.frames), len(rec.sources), rec.maxAt)
	}
	rec.Record(Frame{At: 3, Kind: KindSched, Payload: []byte("evt")})
	arch, err := rec.Stop()
	if err != nil {
		t.Fatalf("Stop: %v", err)
	}
	if arch.Len() != 1 {
		t.Fatalf("archive has %d frames, want 1", arch.Len())
	}
	if got := arch.Frame(0); got.Kind != KindSched || string(got.Payload) != "evt" {
		t.Fatalf("frame 0 = %+v", got)
	}
	if err := arch.Verify(); err != nil {
		t.Fatalf("archive failed verification: %v", err)
	}
}

// TestRecorderStartConcurrentWithRecord proves Start is safe to call while
// frames are being recorded, and that it never injects an extra frame.
func TestRecorderStartConcurrentWithRecord(t *testing.T) {
	t.Parallel()
	rec := NewRecorder(testFingerprint())
	const workers = 8
	var wg sync.WaitGroup
	wg.Add(2 * workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			rec.Start(context.Background())
		}()
		go func() {
			defer wg.Done()
			rec.Record(Frame{At: 1, Kind: KindObs, Payload: []byte("x")})
		}()
	}
	wg.Wait()
	arch, err := rec.Stop()
	if err != nil {
		t.Fatalf("Stop: %v", err)
	}
	if arch.Len() != workers {
		t.Fatalf("archive has %d frames, want %d", arch.Len(), workers)
	}
	if err := arch.Verify(); err != nil {
		t.Fatalf("archive failed verification: %v", err)
	}
}
