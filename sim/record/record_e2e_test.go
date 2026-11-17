/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"testing"

	"abate/kernel/observe"
	"abate/kernel/scheduler"
	"abate/sim/clock"
	"abate/sim/fs"
	"abate/sim/net"
	"abate/sim/rand"
)

// runScenario executes a self-contained mixed experiment through the real
// scheduler, observation pipeline, virtual clock, random stream, network and
// file system, then seals the recording into an archive.
func runScenario(seed uint64) (*Archive, [32]byte) {
	clk := clock.NewVirtualClock()
	rng := rand.New(seed)
	virtualNet := net.NewNetworkWithSeed(clk, seed^0x5a5a)
	virtualNet.Link("a", "b", net.LinkConfig{BaseLatency: 1000, BandwidthBps: 1000})
	if err := virtualNet.InjectFault(net.Fault{Kind: net.FaultLatency, A: "a", B: "b", Value: 777}); err != nil {
		panic(err)
	}
	if err := virtualNet.RecoverFault(net.Fault{A: "a", B: "b"}); err != nil {
		panic(err)
	}
	disk := fs.NewFileSystem(clk)
	if err := disk.MkdirAll("/data"); err != nil {
		panic(err)
	}
	file, err := disk.Create("/data/state.bin")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write([]byte("abate")); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}

	fp := sha256.Sum256([]byte(fmt.Sprintf("e2e/%d", seed)))
	rec := NewRecorder(fp)
	obs := make([]observe.Observation, 0, 4)
	handler := rec.WrapHandler(func(e scheduler.Event) ([]scheduler.Event, error) {
		_, _ = clk.Advance(clock.Duration(1 + rng.IntN(8)))
		_, _ = virtualNet.Send("a", "b", []byte("p"))
		obs = append(obs, observe.Observation{
			At:      e.At.UnixNano(),
			Seq:     e.Seq,
			ProbeID: "e2e",
			Value:   observe.NewNumber(float64(rng.IntN(100))),
			Status:  observe.StatusOK,
		})
		if e.Kind == scheduler.EventApply {
			rec.Record(Frame{
				At:      clock.Instant(e.At.UnixNano()),
				Kind:    KindFault,
				Payload: []byte("latency/recover"),
			})
			return []scheduler.Event{
				{At: e.At.Add(1000), Kind: scheduler.EventProbe, Payload: []byte("probe")},
				{At: e.At.Add(2000), Kind: scheduler.EventRollback, Payload: []byte("rollback")},
			}, nil
		}
		return nil, nil
	})
	res := scheduler.New(handler, scheduler.DefaultBudget()).Run(context.Background(),
		scheduler.Experiment{Seed: seed, Initial: []scheduler.Event{
			{At: 1000, Kind: scheduler.EventApply, Payload: []byte("perturb")},
		}})
	if res.Err != nil {
		panic(res.Err)
	}
	at := clock.Instant(res.End.UnixNano())
	rec.Subscribe(
		NewRNGSource(at, rng),
		NewCLKSource(clk),
		NewOBSSource(at, obs),
		&staticSource{frames: []Frame{{At: at, Kind: KindNet, Payload: []byte("a-b")}}},
	)
	arch, err := rec.Stop()
	if err != nil {
		panic(err)
	}
	return arch, observe.StreamDigest(obs)
}

// TestE2ERecordSaveLoadReplayCompare proves the full chain is lossless.
func TestE2ERecordSaveLoadReplayCompare(t *testing.T) {
	t.Parallel()
	original, digest := runScenario(42)
	data := archiveBytes(t, original)
	loaded := loadBytes(t, data)
	if err := (Replayer{}).Verify(loaded); err != nil {
		t.Fatalf("replay verification: %v", err)
	}
	collector := &frameCollector{}
	if err := (Replayer{}).Replay(loaded, collector); err != nil {
		t.Fatalf("replay: %v", err)
	}
	if len(collector.frames) != loaded.Len() {
		t.Fatalf("replayed %d of %d frames", len(collector.frames), loaded.Len())
	}
	prev := loaded.RootHash()
	for i, f := range collector.frames {
		h := chainHash(prev, f)
		if h != loaded.Hash(i) {
			t.Fatalf("replayed frame %d hash mismatch", i)
		}
		prev = h
	}
	diff, err := Compare(original, loaded)
	if err != nil || !diff.Equal || diff.First != nil {
		t.Fatalf("e2e compare: equal=%v first=%v err=%v", diff.Equal, diff.First, err)
	}
	if diff.Summary.FramesCompared != loaded.Len() || diff.Summary.FieldsDifferent != 0 {
		t.Fatalf("e2e summary = %+v", diff.Summary)
	}
	repeated, repeatedDigest := runScenario(42)
	if digest != repeatedDigest || !bytes.Equal(data, archiveBytes(t, repeated)) {
		t.Fatal("same-seed e2e run did not reproduce byte-identical state and archive")
	}
}

// TestE2EStateRestoreContinuesIdentically proves RNG and clock restoration
// resumes the original trajectory without re-simulating the prefix.
func TestE2EStateRestoreContinuesIdentically(t *testing.T) {
	t.Parallel()
	origin := rand.New(2024)
	origin.Uint64()
	origin.IntN(97)
	origin.Float64()
	seed, draws := origin.State()
	want := make([]uint64, 0, 5)
	for i := 0; i < 5; i++ {
		want = append(want, origin.Uint64())
	}
	restored := rand.New(seed)
	if err := restored.Restore(seed, draws); err != nil {
		t.Fatalf("restore: %v", err)
	}
	for i := range want {
		if got := restored.Uint64(); got != want[i] {
			t.Fatalf("restored draw %d = %d, want %d", i, got, want[i])
		}
	}
	clk := clock.NewVirtualClock()
	if _, err := clk.Advance(clock.Duration(500)); err != nil {
		t.Fatalf("advance: %v", err)
	}
	clk.Sync(clk.Now())
	if clk.Now() != clock.Instant(500) {
		t.Fatalf("clock restore reading = %d", clk.Now())
	}
	if _, err := clk.Advance(clock.Duration(100)); err != nil {
		t.Fatalf("post-restore advance: %v", err)
	}
	if clk.Now() != clock.Instant(600) {
		t.Fatalf("post-restore trajectory diverged: %d", clk.Now())
	}
}
