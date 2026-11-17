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
	"encoding/hex"
	"testing"

	"abate/internal/determinism"
	"abate/sim/rand"
)

// contractSequenceSeed fixes the origin of the golden RNG draw stream.
const contractSequenceSeed uint64 = 0xC0FFEE

// contractSnapshotSHA locks the canonical JSON digest of the accounted draw
// ledger. It was captured from the pre-optimization implementation, so a change
// to Kind, Args (including nil versus empty), Index or Result breaks this test
// instead of silently rewriting recorded RNG frames.
const contractSnapshotSHA = "aca02b2c9fde8ee89aa5452189bcc52e4fb50335df8681064c182672329b52c5"

// contractEncodedSHA locks the SHA-256 of the concatenated encodeDraw bytes of
// the same ledger, i.e. the exact RNG frame payloads written to an archive. It
// was captured from the pre-optimization implementation.
const contractEncodedSHA = "31c634fb04efdfcae4a3a7d2c218917840ffa5d50001c072c110e878490c624b"

// Golden ledger shape of the fixed contract sequence.
const (
	contractRecordCount  = 312
	contractEncodedBytes = 14395
)

// contractStream produces a fixed, fully accounted draw sequence exercising
// every public draw method. It crosses several ledger chunks and mixes
// argument-free and single-argument draws so the golden digest covers both.
func contractStream() *rand.Stream {
	s := rand.New(contractSequenceSeed)
	s.Uint64()
	s.Int63()
	s.Float64()
	s.IntN(17)
	s.Bool()
	s.Bytes(3)
	s.Bytes(11)
	_ = s.Fork()
	for i := 0; i < 300; i++ {
		s.Uint64()
	}
	s.IntN(1000)
	s.Bytes(9)
	return s
}

// TestDrawSnapshotContentGolden proves the accounted draw ledger keeps exactly
// the same element values as the pre-optimization implementation. The digest is
// the canonical JSON used by the determinism certification, so a nil-to-empty
// Args change or a reordered Index would be caught here.
func TestDrawSnapshotContentGolden(t *testing.T) {
	t.Parallel()
	snap := contractStream().Snapshot()
	if len(snap) != contractRecordCount {
		t.Fatalf("ledger len = %d, want %d", len(snap), contractRecordCount)
	}
	if snap[0].Args != nil {
		t.Fatal("argument-free draw must carry a nil Args slice")
	}
	if snap[3].Kind != "intn" || len(snap[3].Args) != 1 || snap[3].Args[0] != 17 {
		t.Fatalf("intn draw args = %v, want [17]", snap[3].Args)
	}
	if snap[6].Kind != "bytes" || len(snap[6].Args) != 1 || snap[6].Args[0] != 11 {
		t.Fatalf("bytes draw args = %v, want [11]", snap[6].Args)
	}
	sum := determinism.HashState(snap)
	if got := hex.EncodeToString(sum[:]); got != contractSnapshotSHA {
		t.Fatalf("snapshot digest = %s, want %s", got, contractSnapshotSHA)
	}
}

// TestDrawFrameEncodingGolden proves encodeDraw consumes the accounted ledger
// into byte-identical RNG frame payloads. Any future attempt to inline Args,
// reorder fields or alter the versioned encoding breaks this digest before it
// can change the hash of a recorded SIM-REC archive.
func TestDrawFrameEncodingGolden(t *testing.T) {
	t.Parallel()
	snap := contractStream().Snapshot()
	var buf bytes.Buffer
	for _, d := range snap {
		_, _ = buf.Write(encodeDraw(d))
	}
	if buf.Len() != contractEncodedBytes {
		t.Fatalf("encoded byte length = %d, want %d", buf.Len(), contractEncodedBytes)
	}
	sum := sha256.Sum256(buf.Bytes())
	if got := hex.EncodeToString(sum[:]); got != contractEncodedSHA {
		t.Fatalf("encoded draw digest = %s, want %s", got, contractEncodedSHA)
	}
}
