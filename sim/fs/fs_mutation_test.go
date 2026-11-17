/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo/163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"abate/sim/clock"
)

// TestMutationBadBlock confirms a bad-block fault makes the range unreadable and
// that recovery restores readability.
func TestMutationBadBlock(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("0123456789"))

	if err := f.InjectFault(Fault{Kind: FaultBadBlock, Path: "/f", Offset: 2, Value: 3}); err != nil {
		t.Fatalf("inject: %v", err)
	}
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	buf := make([]byte, 10)
	if _, err := h.Read(buf); !errors.Is(err, ErrBadBlock) {
		t.Fatalf("read over bad block want ErrBadBlock, got %v", err)
	}
	// A read that avoids the bad range still succeeds.
	if _, err := h.Seek(5, io.SeekStart); err != nil {
		t.Fatalf("seek: %v", err)
	}
	n, err := h.Read(buf)
	if err != nil || n != 5 {
		t.Fatalf("clean read = %v,%d", err, n)
	}
	_ = h.Close()

	if err := f.RecoverFault("/f"); err != nil {
		t.Fatalf("recover: %v", err)
	}
	h2, err := f.Open("/f")
	if err != nil {
		t.Fatalf("reopen: %v", err)
	}
	n, err = h2.Read(buf)
	if err != nil || n != 10 {
		t.Fatalf("post-recover read = %v,%d", err, n)
	}
}

// TestMutationQuotaPerFile confirms a single file cannot exceed 100 MiB.
func TestMutationQuotaPerFile(t *testing.T) {
	clk := clock.NewVirtualClock()
	f := NewFileSystem(clk)
	h, err := f.Create("/big")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	huge := make([]byte, MaxFileBytes+1)
	if _, err := h.Write(huge); !errors.Is(err, ErrQuotaExceeded) {
		t.Fatalf("oversize write want ErrQuotaExceeded, got %v", err)
	}
	// Exactly at the limit is allowed.
	ok := make([]byte, MaxFileBytes)
	if _, err := h.Write(ok); err != nil {
		t.Fatalf("at-limit write err: %v", err)
	}
}

// TestMutationQuotaTotal confirms the aggregate 1 GiB ceiling is enforced. The
// per-file cap is 100 MiB, so the aggregate is reached with ten such files.
func TestMutationQuotaTotal(t *testing.T) {
	clk := clock.NewVirtualClock()
	f := NewFileSystem(clk)
	for i := 0; i < 10; i++ {
		h, err := f.Create(fmt.Sprintf("/f%d", i))
		if err != nil {
			t.Fatalf("create f%d: %v", i, err)
		}
		if _, err := h.Write(make([]byte, MaxFileBytes)); err != nil {
			t.Fatalf("fill f%d: %v", i, err)
		}
	}
	// Ten 100 MiB files total 1048 MiB, still under 1 GiB. An eleventh full
	// 100 MiB file would push the aggregate over the 1 GiB ceiling.
	hExtra, _ := f.Create("/extra")
	if _, err := hExtra.Write(make([]byte, MaxFileBytes)); !errors.Is(err, ErrQuotaExceeded) {
		t.Fatalf("aggregate overflow want ErrQuotaExceeded, got %v", err)
	}
}

// TestMutationQuotaRecoversOnDelete confirms deleting a file frees capacity.
func TestMutationQuotaRecoversOnDelete(t *testing.T) {
	clk := clock.NewVirtualClock()
	f := NewFileSystem(clk)
	for i := 0; i < 10; i++ {
		h, err := f.Create(fmt.Sprintf("/f%d", i))
		if err != nil {
			t.Fatalf("create f%d: %v", i, err)
		}
		if _, err := h.Write(make([]byte, MaxFileBytes)); err != nil {
			t.Fatalf("fill f%d: %v", i, err)
		}
	}
	if err := f.Remove("/f0"); err != nil {
		t.Fatalf("delete f0: %v", err)
	}
	// 100 MiB freed; a replacement file of that size now fits.
	hRep, _ := f.Create("/repl")
	if _, err := hRep.Write(make([]byte, MaxFileBytes)); err != nil {
		t.Fatalf("post-delete write err: %v", err)
	}
}

// TestMutationLatency confirms the latency fault advances the clock per op.
func TestMutationLatency(t *testing.T) {
	clk := clock.NewVirtualClock()
	f := NewFileSystem(clk)
	if err := f.InjectFault(Fault{Kind: FaultLatency, Value: 5000}); err != nil {
		t.Fatalf("inject latency: %v", err)
	}
	before := clk.Now().UnixNano()
	h, err := f.Create("/x")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	_, _ = h.Write([]byte("y"))
	_ = h.Close()
	// Three operations (create, write, close-no-delay) charged latency.
	if clk.Now().UnixNano()-before < 5000 {
		t.Fatalf("latency not applied: delta=%d", clk.Now().UnixNano()-before)
	}
}

// TestMutationFaultOnUnknown reports faults targeting missing entries.
func TestMutationFaultOnUnknown(t *testing.T) {
	f, _ := newFS()
	if err := f.InjectFault(Fault{Kind: FaultBadBlock, Path: "/ghost", Offset: 0, Value: 1}); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("bad-block on missing want ErrNoEntry, got %v", err)
	}
	if err := f.InjectFault(Fault{Kind: FaultKind("bogus")}); !errors.Is(err, ErrUnknownFault) {
		t.Fatalf("unknown fault want ErrUnknownFault, got %v", err)
	}
}
