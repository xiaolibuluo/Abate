/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestFrameKindRoundTrip(t *testing.T) {
	t.Parallel()
	for i, k := range allKinds() {
		tag := k.Marshal()
		if len(tag) != 1 || tag[0] != byte(i) { //nolint:gosec // the tag table is a fixed eight-element slice.
			t.Fatalf("kind %s tag = %v", k, tag)
		}
		got, ok := FrameKindFromTag(tag[0])
		if !ok || got != k || k.String() != string(k) {
			t.Fatalf("kind %s round trip = %s,%v", k, got, ok)
		}
	}
	if _, ok := FrameKindFromTag(0xFF); ok {
		t.Fatal("an unknown tag must be rejected")
	}
	if got := FrameKind("UNKNOWN").Marshal(); got[0] != 0xFF {
		t.Fatalf("unknown kind tag = %v", got)
	}
}

func TestArchiveSaveLoadAllKinds(t *testing.T) {
	t.Parallel()
	want := sampleFrames()
	arch := makeArchive(t, want)
	got := loadBytes(t, archiveBytes(t, arch))
	if got.Version != arch.Version || got.Fingerprint != arch.Fingerprint {
		t.Fatalf("header mismatch: %+v", got)
	}
	if got.Len() != len(want) || got.RootHash() != arch.RootHash() {
		t.Fatalf("length or root hash mismatch: %d", got.Len())
	}
	if err := got.Verify(); err != nil {
		t.Fatalf("loaded archive failed verification: %v", err)
	}
	for i := range want {
		f := got.Frame(i)
		if f.Seq != uint64(i+1) || f.At != want[i].At || f.Kind != want[i].Kind || //nolint:gosec // loop index is a small constant.
			!bytes.Equal(f.Payload, want[i].Payload) {
			t.Fatalf("frame %d = %+v, want %+v", i, f, want[i])
		}
		if got.Hash(i) != arch.Hash(i) {
			t.Fatalf("frame %d chain hash mismatch", i)
		}
	}
}

func TestArchiveByteFlipRejected(t *testing.T) {
	t.Parallel()
	data := archiveBytes(t, makeArchive(t, sampleFrames()))
	for i := range data {
		mut := append([]byte(nil), data...)
		mut[i] ^= 0xFF
		if _, err := Load(bytes.NewReader(mut)); !errors.Is(err, ErrInvalidFormat) {
			t.Fatalf("byte %d flip was accepted: %v", i, err)
		}
	}
}

func TestReplayFrameHashes(t *testing.T) {
	t.Parallel()
	arch := makeArchive(t, sampleFrames())
	c := &frameCollector{}
	if err := (Replayer{}).Replay(arch, c); err != nil {
		t.Fatalf("replay: %v", err)
	}
	if len(c.frames) != arch.Len() {
		t.Fatalf("replayed %d frames, want %d", len(c.frames), arch.Len())
	}
	prev := arch.RootHash()
	for i, f := range c.frames {
		if f.Seq != arch.Frame(i).Seq || f.Kind != arch.Frame(i).Kind ||
			!bytes.Equal(f.Payload, arch.Frame(i).Payload) {
			t.Fatalf("replayed frame %d = %+v", i, f)
		}
		h := chainHash(prev, f)
		if h != arch.Hash(i) {
			t.Fatalf("frame %d hash chain mismatch", i)
		}
		prev = h
	}
}

func TestReplayHookErrorAndNilHooks(t *testing.T) {
	t.Parallel()
	arch := makeArchive(t, sampleFrames())
	c := &frameCollector{failAt: 3}
	if err := (Replayer{}).Replay(arch, c); !errors.Is(err, errHookStop) {
		t.Fatalf("hook error = %v", err)
	}
	if len(c.frames) != 3 {
		t.Fatalf("replay did not stop at the failing hook: %d", len(c.frames))
	}
	if err := (Replayer{}).Replay(arch, nil); err != nil {
		t.Fatalf("nil hooks replay: %v", err)
	}
}

func TestCompareEqualAndFirstDifference(t *testing.T) {
	t.Parallel()
	a := makeArchive(t, sampleFrames())
	b := makeArchive(t, sampleFrames())
	d, err := Compare(a, b)
	if err != nil || !d.Equal || d.First != nil {
		t.Fatalf("identical archives: equal=%v first=%v err=%v", d.Equal, d.First, err)
	}
	if d.Summary.FramesCompared != a.Len() || d.Summary.FieldsDifferent != 0 {
		t.Fatalf("equality summary = %+v", d.Summary)
	}
	frames := sampleFrames()
	frames[2].Kind = KindNet
	c := makeArchive(t, frames)
	diff, err := Compare(a, c)
	if err != nil || diff.Equal || diff.First == nil {
		t.Fatalf("one-frame difference: equal=%v first=%v err=%v", diff.Equal, diff.First, err)
	}
	if diff.First.Path != "Frames[2].Kind" ||
		diff.First.Kind != allKinds()[2] ||
		diff.First.Want != string(allKinds()[2]) || diff.First.Got != "NET" {
		t.Fatalf("first difference is not the frame kind: %+v", diff.First)
	}
	if diff.Summary.FirstDivergentFrame != 2 || diff.Summary.FieldsDifferent == 0 {
		t.Fatalf("difference summary = %+v", diff.Summary)
	}
	t.Logf("REC Compare first difference: %s", diff.First.String())
}

// TestCompareIgnoresConstructionOrder proves archives built through different
// recorder paths but with identical frame content still compare equal.
func TestCompareIgnoresConstructionOrder(t *testing.T) {
	t.Parallel()
	frames := sampleFrames()
	byRecord := makeArchive(t, frames)
	rec := NewRecorder(testFingerprint())
	rec.Subscribe(&staticSource{frames: frames})
	bySource, err := rec.Stop()
	if err != nil {
		t.Fatalf("stop source-based recorder: %v", err)
	}
	if !bytes.Equal(archiveBytes(t, byRecord), archiveBytes(t, bySource)) {
		t.Fatal("different construction paths produced different archive bytes")
	}
	d, err := Compare(byRecord, bySource)
	if err != nil || !d.Equal || d.First != nil {
		t.Fatalf("construction order changed the comparison: equal=%v first=%v err=%v",
			d.Equal, d.First, err)
	}
}

func TestCompareStructuralPayloadPath(t *testing.T) {
	t.Parallel()
	frames := sampleFrames()
	frames[4].Payload = append([]byte(nil), frames[4].Payload...)
	frames[4].Payload[1] ^= 0xFF
	a := makeArchive(t, sampleFrames())
	b := makeArchive(t, frames)
	d, err := Compare(a, b)
	if err != nil || d.Equal || d.First == nil {
		t.Fatalf("payload mutation: equal=%v first=%v err=%v", d.Equal, d.First, err)
	}
	if d.First.Path != "Frames[4].Payload[1]" || d.First.ByteOffset != 1 ||
		d.First.Field != "Payload" || d.First.FrameIndex != 4 {
		t.Fatalf("payload divergence was not localized: %+v", d.First)
	}
	if !strings.Contains(d.First.String(), "path=Frames[4].Payload[1]") {
		t.Fatalf("diff point rendering = %s", d.First.String())
	}
	t.Logf("REC Compare payload mutation: %s", d.First.String())
}

func TestCompareDifferentLengthsAndErrors(t *testing.T) {
	t.Parallel()
	a := makeArchive(t, sampleFrames())
	short := makeArchive(t, sampleFrames()[:3])
	d, err := Compare(a, short)
	if err != nil || d.Equal || d.First == nil ||
		!strings.HasPrefix(d.First.Path, "Frames[") {
		t.Fatalf("length difference: equal=%v first=%v err=%v", d.Equal, d.First, err)
	}
	longer := makeArchive(t, append(sampleFrames(), Frame{
		At: 9, Kind: KindFS, Payload: []byte("extra"),
	}))
	longDiff, err := Compare(a, longer)
	if err != nil || longDiff.Equal || longDiff.First == nil {
		t.Fatalf("extra-frame difference: equal=%v first=%v err=%v",
			longDiff.Equal, longDiff.First, err)
	}
	if _, err := Compare(nil, a); !errors.Is(err, ErrNilArchive) {
		t.Fatalf("nil left: %v", err)
	}
	if _, err := Compare(a, nil); !errors.Is(err, ErrNilArchive) {
		t.Fatalf("nil right: %v", err)
	}
	bad := makeArchive(t, sampleFrames())
	bad.entries[1].frame.Payload[0] ^= 0xFF
	if _, err := Compare(a, bad); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("corrupt left: %v", err)
	}
	if _, err := Compare(bad, a); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("corrupt right: %v", err)
	}
}

func TestVerifySequenceAndTimeMonotonicity(t *testing.T) {
	t.Parallel()
	badSeq := makeArchive(t, sampleFrames())
	badSeq.entries[3].frame.Seq = 99
	reseal(badSeq)
	if err := (Replayer{}).Verify(badSeq); err == nil ||
		!strings.Contains(err.Error(), "sequence discontinuity") {
		t.Fatalf("out-of-order sequence was accepted: %v", err)
	}
	badTime := makeArchive(t, []Frame{
		{At: 10, Kind: KindSched, Payload: []byte("a")},
		{At: 5, Kind: KindSched, Payload: []byte("b")},
	})
	if err := (Replayer{}).Verify(badTime); err == nil ||
		!strings.Contains(err.Error(), "not monotonic") {
		t.Fatalf("time regression was accepted: %v", err)
	}
	if err := (Replayer{}).Verify(makeArchive(t, sampleFrames())); err != nil {
		t.Fatalf("valid archive rejected: %v", err)
	}
	if err := (Replayer{}).Verify(makeArchive(t, nil)); err != nil {
		t.Fatalf("empty archive rejected: %v", err)
	}
}

func TestDiffSummaryMapOrderAndRendering(t *testing.T) {
	t.Parallel()
	one := DiffSummary{KindCounts: map[FrameKind]int{}}
	one.KindCounts[KindSched] = 1
	one.KindCounts[KindObs] = 2
	one.KindCounts[KindRNG] = 3
	two := DiffSummary{KindCounts: map[FrameKind]int{}}
	two.KindCounts[KindRNG] = 3
	two.KindCounts[KindObs] = 2
	two.KindCounts[KindSched] = 1
	if one.String() != two.String() {
		t.Fatalf("map insertion order leaked: %s vs %s", one.String(), two.String())
	}
	if !strings.Contains(one.String(), "kindCounts=[OBS=2 RNG=3 SCHED=1]") {
		t.Fatalf("kind counts are not sorted: %s", one.String())
	}
	equal, err := Compare(makeArchive(t, sampleFrames()), makeArchive(t, sampleFrames()))
	if err != nil || !strings.Contains(equal.String(), "equal: true") ||
		!strings.Contains(equal.String(), "summary: leftFrames=") {
		t.Fatalf("equal diff rendering = %q, err=%v", equal.String(), err)
	}
	changed := sampleFrames()
	changed[0].Payload[0] ^= 0xFF
	differ, err := Compare(makeArchive(t, sampleFrames()), makeArchive(t, changed))
	if err != nil || !strings.Contains(differ.String(), "equal: false") ||
		!strings.Contains(differ.String(), "firstDiff: path=") {
		t.Fatalf("differing diff rendering = %q, err=%v", differ.String(), err)
	}
}
