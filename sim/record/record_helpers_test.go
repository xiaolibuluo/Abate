/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"testing"

	"abate/kernel/observe"
	"abate/sim/clock"
)

// testFingerprint is the fixed experiment fingerprint of the test fixtures.
func testFingerprint() [32]byte {
	return sha256.Sum256([]byte("abate/record/test"))
}

// allKinds enumerates every declared frame kind in tag order.
func allKinds() []FrameKind {
	return []FrameKind{KindSched, KindObs, KindRNG, KindClk,
		KindFault, KindNet, KindFS, KindVTask}
}

// sampleFrames returns one frame per kind with distinct, non-empty payloads.
func sampleFrames() []Frame {
	kinds := allKinds()
	out := make([]Frame, 0, len(kinds))
	for i, k := range kinds {
		out = append(out, Frame{
			At:      clock.Instant(i + 1),
			Kind:    k,
			Payload: []byte{byte(0xA0 + i), byte(i), 0x5A},
		})
	}
	return out
}

// makeArchive seals frames through a real Recorder so the chain hashes are valid.
func makeArchive(t *testing.T, frames []Frame) *Archive {
	t.Helper()
	return makeArchiveWith(t, testFingerprint(), frames)
}

// makeArchiveWith seals frames under an explicit fingerprint.
func makeArchiveWith(t *testing.T, fp [32]byte, frames []Frame) *Archive {
	t.Helper()
	rec := NewRecorder(fp)
	for _, f := range frames {
		rec.Record(f)
	}
	arch, err := rec.Stop()
	if err != nil {
		t.Fatalf("stop recorder: %v", err)
	}
	return arch
}

// archiveBytes renders an archive through Save.
func archiveBytes(t *testing.T, a *Archive) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := a.Save(&buf); err != nil {
		t.Fatalf("save archive: %v", err)
	}
	return buf.Bytes()
}

// loadBytes loads an archive from raw bytes.
func loadBytes(t *testing.T, data []byte) *Archive {
	t.Helper()
	a, err := Load(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("load archive: %v", err)
	}
	return a
}

// chainHash recomputes the chain hash of one frame for tamper tests.
func chainHash(prev [32]byte, f Frame) [32]byte {
	body := encodeFrameBody(f)
	buf := make([]byte, 0, len(prev)+len(body))
	buf = append(buf, prev[:]...)
	buf = append(buf, body...)
	return sha256.Sum256(buf)
}

// errHookStop is the deterministic sentinel reported by a failing replay hook.
var errHookStop = errors.New("record test: hook stop")

// frameCollector gathers replayed frames and can fail on a chosen ordinal.
type frameCollector struct {
	frames []Frame
	failAt int
}

// OnFrame records one frame and optionally aborts the replay.
func (c *frameCollector) OnFrame(f Frame) error {
	c.frames = append(c.frames, f)
	if c.failAt > 0 && len(c.frames) == c.failAt {
		return errHookStop
	}
	return nil
}

// reseal recomputes the whole chain so Archive.Verify passes after a field edit,
// leaving Replayer.Verify free to detect sequence or time violations.
func reseal(a *Archive) {
	prev := a.headerHash()
	for i := range a.entries {
		a.entries[i].prev = prev
		a.entries[i].hash = chainHash(prev, a.entries[i].frame)
		prev = a.entries[i].hash
	}
}

// failingWriter always fails a Write call.
type failingWriter struct{}

// Write reports a deterministic write failure.
func (failingWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("record test: write failed")
}

// failingReader always fails a Read call.
type failingReader struct{}

// Read reports a deterministic read failure.
func (failingReader) Read(_ []byte) (int, error) {
	return 0, errors.New("record test: read failed")
}

// sampleObservations builds four deterministic observations for OBS frames.
func sampleObservations() []observe.Observation {
	out := make([]observe.Observation, 0, 4)
	for i := 0; i < 4; i++ {
		out = append(out, observe.Observation{
			At:      int64(i + 1),
			Seq:     uint64(i + 1), //nolint:gosec // the loop bound is a small constant.
			ProbeID: "p1",
			Value:   observe.NewNumber(float64(i)),
			Status:  observe.StatusOK,
		})
	}
	return out
}
