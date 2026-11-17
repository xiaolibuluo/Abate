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
	"encoding/binary"
	"errors"
	"testing"

	"abate/kernel/scheduler"
	"abate/sim/clock"
	"abate/sim/rand"
)

// resealBytes rewrites the tail checksum so Load reaches deeper parse branches.
func resealBytes(data []byte) []byte {
	out := append([]byte(nil), data...)
	end := len(out) - sha256.Size
	sum := sha256.Sum256(out[:end])
	copy(out[end:], sum[:])
	return out
}

func TestMutationEveryFieldDetected(t *testing.T) {
	t.Parallel()
	base := sampleFrames()
	for _, field := range []string{"At", "Kind", "Payload[0]", "Payload[1]", "Payload[2]"} {
		frames := sampleFrames()
		switch field {
		case "At":
			frames[1].At += 100
		case "Kind":
			frames[1].Kind = KindVTask
		default:
			idx := int(field[len("Payload[")] - '0')
			frames[1].Payload = append([]byte(nil), frames[1].Payload...)
			frames[1].Payload[idx] ^= 0xFF
		}
		d, err := Compare(makeArchive(t, base), makeArchive(t, frames))
		if err != nil || d.Equal || d.First == nil || d.First.FrameIndex != 1 {
			t.Fatalf("field %s mutation not localized: equal=%v first=%v err=%v",
				field, d.Equal, d.First, err)
		}
	}
}

func TestMutationEveryKindPayloadDetected(t *testing.T) {
	t.Parallel()
	for i := range allKinds() {
		frames := sampleFrames()
		frames[i].Payload = append([]byte(nil), frames[i].Payload...)
		frames[i].Payload[0] ^= 0xFF
		d, err := Compare(makeArchive(t, sampleFrames()), makeArchive(t, frames))
		if err != nil || d.Equal || d.First == nil || d.First.FrameIndex != i {
			t.Fatalf("kind %s payload mutation undetected: equal=%v first=%v err=%v",
				allKinds()[i], d.Equal, d.First, err)
		}
		if d.First.Kind != allKinds()[i] || d.First.Field != "Payload" {
			t.Fatalf("kind %s divergence misreported: %+v", allKinds()[i], d.First)
		}
	}
}

func TestMutationEmptyAndMalformedArchives(t *testing.T) {
	t.Parallel()
	for _, data := range [][]byte{
		nil,
		{},
		[]byte("ABATEREC"),
		make([]byte, archiveHeaderLen+sha256.Size-1),
	} {
		if _, err := Load(bytes.NewReader(data)); !errors.Is(err, ErrInvalidFormat) {
			t.Fatalf("malformed input of %d bytes accepted: %v", len(data), err)
		}
	}
	base := archiveBytes(t, makeArchive(t, sampleFrames()))
	variants := make(map[string][]byte)
	badMagic := append([]byte(nil), base...)
	badMagic[0] ^= 0xFF
	variants["magic"] = resealBytes(badMagic)
	badVersion := append([]byte(nil), base...)
	badVersion[9] = 0x7F
	variants["version"] = resealBytes(badVersion)
	badChecksum := append([]byte(nil), base...)
	badChecksum[len(badChecksum)-1] ^= 0xFF
	variants["checksum"] = badChecksum
	hugeCount := append([]byte(nil), base...)
	binary.BigEndian.PutUint64(hugeCount[42:50], 4096)
	variants["frame-count"] = resealBytes(hugeCount)
	shortIndex := append([]byte(nil), base...)
	binary.BigEndian.PutUint64(shortIndex[42:50], 1)
	variants["index-length"] = resealBytes(shortIndex)
	badTag := append([]byte(nil), base...)
	badTag[archiveHeaderLen+32+8+8] = 0xFF
	variants["frame-kind"] = resealBytes(badTag)
	badLength := append([]byte(nil), base...)
	binary.BigEndian.PutUint64(badLength[archiveHeaderLen+32+8+8+1:], ^uint64(0))
	variants["payload-length"] = resealBytes(badLength)
	brokenChain := append([]byte(nil), base...)
	brokenChain[archiveHeaderLen+32+8+8+1+8] ^= 0xFF
	variants["broken-chain"] = resealBytes(brokenChain)
	for name, data := range variants {
		if _, err := Load(bytes.NewReader(data)); !errors.Is(err, ErrInvalidFormat) {
			t.Fatalf("%s variant was accepted: %v", name, err)
		}
	}
}

func TestMutationSingleFrameAndInputSeq(t *testing.T) {
	t.Parallel()
	single := makeArchive(t, []Frame{{At: 2, Kind: KindFault, Payload: []byte("z")}})
	if single.Len() != 1 {
		t.Fatalf("single frame archive length = %d", single.Len())
	}
	loaded := loadBytes(t, archiveBytes(t, single))
	if err := (Replayer{}).Verify(loaded); err != nil {
		t.Fatalf("single frame verification: %v", err)
	}
	d, err := Compare(single, loaded)
	if err != nil || !d.Equal {
		t.Fatalf("single frame self comparison: equal=%v err=%v", d.Equal, err)
	}
	c := &frameCollector{}
	if err := (Replayer{}).Replay(loaded, c); err != nil || len(c.frames) != 1 {
		t.Fatalf("single frame replay: %d frames, err=%v", len(c.frames), err)
	}
	outOfOrder := makeArchive(t, []Frame{
		{Seq: 99, At: 1, Kind: KindSched, Payload: []byte("a")},
		{Seq: 1, At: 2, Kind: KindSched, Payload: []byte("b")},
		{Seq: 50, At: 3, Kind: KindSched, Payload: []byte("c")},
	})
	for i := 0; i < outOfOrder.Len(); i++ {
		if outOfOrder.Frame(i).Seq != uint64(i+1) { //nolint:gosec // loop index is a small constant.
			t.Fatalf("Stop did not normalize sequence: frame %d seq=%d", i, outOfOrder.Frame(i).Seq)
		}
	}
	if err := (Replayer{}).Verify(outOfOrder); err != nil {
		t.Fatalf("normalized archive rejected: %v", err)
	}
}

func TestCodecRoundTrips(t *testing.T) {
	t.Parallel()
	body := encodeFrameBody(Frame{Seq: 7, At: 1234, Kind: KindSched, Payload: []byte("abc")})
	if len(body) != 28 || binary.BigEndian.Uint64(body[0:8]) != 7 ||
		binary.BigEndian.Uint64(body[8:16]) != 1234 || body[16] != 0 ||
		binary.BigEndian.Uint64(body[17:25]) != 3 || string(body[25:]) != "abc" {
		t.Fatalf("frame body layout = %x", body)
	}
	var buf bytes.Buffer
	writeUint64(&buf, 42)
	writeUint16(&buf, 7)
	if !bytes.Equal(buf.Bytes(), []byte{0, 0, 0, 0, 0, 0, 0, 42, 0, 7}) {
		t.Fatalf("buffer helpers = %x", buf.Bytes())
	}
	tmp := make([]byte, 8)
	putUint64(tmp, 99)
	if binary.BigEndian.Uint64(tmp) != 99 {
		t.Fatalf("putUint64 = %x", tmp)
	}
	event := scheduler.Event{At: 9, Seq: 3, Kind: scheduler.EventApply, Payload: []byte("p")}
	if len(encodeEvent(event)) == 0 {
		t.Fatal("encodeEvent produced no bytes")
	}
	draw := rand.DrawRecord{Index: 2, Kind: "intn", Args: []uint64{10, 20}, Result: 42}
	if len(encodeDraw(draw)) == 0 {
		t.Fatal("encodeDraw produced no bytes")
	}
	if clockStateString(clock.Instant(5), clock.StateRunning) != "5 RUNNING" {
		t.Fatal("clock state string is not deterministic")
	}
	if string(encodeClock(clock.Instant(5), clock.StateFast)) != "5 FAST" {
		t.Fatal("encodeClock lost the clock state")
	}
	if len(encodeObservation(sampleObservations()[0])) == 0 {
		t.Fatal("encodeObservation produced no bytes")
	}
}

func TestMutationIOErrorBranches(t *testing.T) {
	t.Parallel()
	if err := makeArchive(t, sampleFrames()).Save(failingWriter{}); err == nil {
		t.Fatal("Save accepted a failing writer")
	}
	if _, err := Load(failingReader{}); err == nil {
		t.Fatal("Load accepted a failing reader")
	}
}

func TestMutationChainLinkageAndPayloadLength(t *testing.T) {
	t.Parallel()
	linked := makeArchive(t, sampleFrames())
	linked.entries[2].prev[0] ^= 0xFF
	linked.entries[2].hash = chainHash(linked.entries[2].prev, linked.entries[2].frame)
	if err := linked.Verify(); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("broken prev linkage accepted: %v", err)
	}
	shorter := sampleFrames()
	shorter[1].Payload = shorter[1].Payload[:1]
	d, err := Compare(makeArchive(t, sampleFrames()), makeArchive(t, shorter))
	if err != nil || d.Equal || d.First == nil ||
		d.First.Path != "Frames[1].Payload.len" {
		t.Fatalf("payload length difference: equal=%v first=%v err=%v",
			d.Equal, d.First, err)
	}
}

func TestMutationLoadTruncatedFrames(t *testing.T) {
	t.Parallel()
	for _, extra := range []int{8, 40, 60} {
		raw := make([]byte, archiveHeaderLen+extra)
		copy(raw[0:8], archiveMagic[:])
		binary.BigEndian.PutUint16(raw[8:10], archiveVersion)
		binary.BigEndian.PutUint64(raw[42:50], 1)
		raw = append(raw, make([]byte, sha256.Size)...)
		sum := sha256.Sum256(raw[:len(raw)-sha256.Size])
		copy(raw[len(raw)-sha256.Size:], sum[:])
		if _, err := Load(bytes.NewReader(raw)); !errors.Is(err, ErrInvalidFormat) {
			t.Fatalf("truncated frame of %d body bytes accepted: %v", extra, err)
		}
	}
}

func TestSourcesCollectAndRecorderBranches(t *testing.T) {
	t.Parallel()
	clk := clock.NewVirtualClock()
	if _, err := clk.Advance(clock.Duration(7)); err != nil {
		t.Fatalf("advance: %v", err)
	}
	clkFrames := NewCLKSource(clk).Collect()
	if len(clkFrames) != 1 || clkFrames[0].Kind != KindClk ||
		clkFrames[0].At != clk.Now() ||
		string(clkFrames[0].Payload) != clockStateString(clk.Now(), clk.State()) {
		t.Fatalf("CLK source = %+v", clkFrames)
	}
	rng := rand.New(42)
	rng.Uint64()
	rng.IntN(10)
	rngFrames := NewRNGSource(clock.Instant(3), rng).Collect()
	if len(rngFrames) != 2 {
		t.Fatalf("RNG source produced %d frames", len(rngFrames))
	}
	for _, f := range rngFrames {
		if f.Kind != KindRNG || f.At != 3 || len(f.Payload) == 0 {
			t.Fatalf("RNG frame = %+v", f)
		}
	}
	for _, f := range NewOBSSource(clock.Instant(4), sampleObservations()).Collect() {
		if f.Kind != KindObs || f.At != 4 || len(f.Payload) == 0 {
			t.Fatalf("OBS frame = %+v", f)
		}
	}

	rec := NewRecorder(testFingerprint())
	rec.Start(context.Background())
	inner := func(e scheduler.Event) ([]scheduler.Event, error) {
		return []scheduler.Event{{At: e.At, Kind: scheduler.EventProbe, Payload: []byte("child")}}, nil
	}
	children, err := rec.WrapHandler(inner)(scheduler.Event{
		At: 12, Kind: scheduler.EventApply, Payload: []byte("x"),
	})
	if err != nil || len(children) != 1 || children[0].Kind != scheduler.EventProbe {
		t.Fatalf("WrapHandler altered the event stream: %+v, %v", children, err)
	}
	rec.Record(Frame{Kind: KindNet, Payload: []byte("net")})
	rec.Subscribe(&staticSource{frames: []Frame{{At: 30, Kind: KindFS, Payload: []byte("fs")}}})
	arch, err := rec.Stop()
	if err != nil {
		t.Fatalf("stop: %v", err)
	}
	if arch.Len() != 3 || arch.Frame(1).At != 30 {
		t.Fatalf("Stop did not stamp zero-time frames: %+v", arch.Frame(1))
	}
	if err := (Replayer{}).Verify(arch); err != nil {
		t.Fatalf("recorder branches produced an invalid archive: %v", err)
	}
}
