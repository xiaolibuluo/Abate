/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

import "testing"

// TestChunkedLedgerSnapshot proves a ledger spanning many chunks is flattened
// into the same contiguous, indexed and defensive slice as a single-slice ledger.
func TestChunkedLedgerSnapshot(t *testing.T) {
	t.Parallel()
	const total = 3*drawChunkMax + 7
	s := New(20260601)
	want := make([]uint64, total)
	for i := 0; i < total; i++ {
		want[i] = s.Uint64()
	}
	if s.acct.Len() != total {
		t.Fatalf("accounting len = %d, want %d", s.acct.Len(), total)
	}
	snap := s.Snapshot()
	if len(snap) != total {
		t.Fatalf("snapshot len = %d, want %d", len(snap), total)
	}
	if len(s.acct.chunks) < 4 {
		t.Fatalf("ledger used %d chunks, want at least 4", len(s.acct.chunks))
	}
	for i := range snap {
		if snap[i].Index != i || snap[i].Result != want[i] {
			t.Fatalf("record %d = (index %d, result %d), want (index %d, result %d)",
				i, snap[i].Index, snap[i].Result, i, want[i])
		}
		if snap[i].Kind != "uint64" || snap[i].Args != nil {
			t.Fatalf("record %d carries unexpected kind/args %q/%v", i, snap[i].Kind, snap[i].Args)
		}
	}
	snap[0].Result = 0
	if s.Snapshot()[0].Result == 0 {
		t.Fatal("chunked snapshot is not a defensive copy")
	}
}

// TestChunkedLedgerEmptyAfterRestore proves a restored stream starts from an
// empty ledger and remains fully accounted for subsequent draws.
func TestChunkedLedgerEmptyAfterRestore(t *testing.T) {
	t.Parallel()
	s := New(77)
	for i := 0; i < 2*drawChunkMax; i++ {
		s.Uint64()
	}
	if err := s.Restore(77, 2*drawChunkMax); err != nil {
		t.Fatalf("Restore err = %v", err)
	}
	if got := s.Snapshot(); len(got) != 0 {
		t.Fatalf("restored ledger len = %d, want 0", len(got))
	}
	for i := 0; i < drawChunkMax+1; i++ {
		s.IntN(13)
	}
	snap := s.Snapshot()
	if len(snap) != drawChunkMax+1 {
		t.Fatalf("post-restore ledger len = %d, want %d", len(snap), drawChunkMax+1)
	}
	for i := range snap {
		if snap[i].Index != i || snap[i].Kind != "intn" ||
			len(snap[i].Args) != 1 || snap[i].Args[0] != 13 {
			t.Fatalf("post-restore record %d malformed: %+v", i, snap[i])
		}
	}
}

// TestChunkSizingAndKindCodes locks the geometric chunk growth and the kind-code
// mapping so a future refactor cannot silently change the ledger layout.
func TestChunkSizingAndKindCodes(t *testing.T) {
	t.Parallel()
	wantSizes := []int{drawChunkMin, drawChunkMax, drawChunkMax,
		drawChunkMax, drawChunkMax, drawChunkMax}
	for n, want := range wantSizes {
		if got := nextChunkSize(n); got != want {
			t.Fatalf("nextChunkSize(%d) = %d, want %d", n, got, want)
		}
	}
	codes := []uint8{drawKindUint64, drawKindInt63, drawKindFloat64,
		drawKindIntN, drawKindBool, drawKindBytes, drawKindFork}
	wantKinds := []string{"uint64", "int63", "float64", "intn", "bool", "bytes", "fork"}
	for i, code := range codes {
		if got := drawKindName(code); got != wantKinds[i] {
			t.Fatalf("drawKindName(%d) = %q, want %q", code, got, wantKinds[i])
		}
	}
	if got := drawKindName(255); got != "" {
		t.Fatalf("drawKindName(255) = %q, want empty", got)
	}
}
