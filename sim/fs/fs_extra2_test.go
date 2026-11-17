/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import (
	"errors"
	"testing"
)

// TestFileReadAtBounds covers the out-of-range and clamp branches of ReadAt.
func TestFileReadAtBounds(t *testing.T) {
	fl := newFile("f", 0)
	if _, err := fl.WriteAt(0, []byte("abcdef"), 0); err != nil {
		t.Fatalf("writeAt: %v", err)
	}
	// Reading past the end clamps to the file size.
	out, err := fl.ReadAt(4, 10)
	if err != nil || string(out) != "ef" {
		t.Fatalf("read past end = %q,%v", out, err)
	}
	// Reading exactly at the size returns empty (EOF), not an error.
	eof, err := fl.ReadAt(6, 1)
	if err != nil || len(eof) != 0 {
		t.Fatalf("read at size = %q,%v", eof, err)
	}
	// An offset beyond the size is an error.
	if _, err := fl.ReadAt(7, 1); !errors.Is(err, ErrOffset) {
		t.Fatalf("offset beyond size want ErrOffset, got %v", err)
	}
	if _, err := fl.ReadAt(-1, 1); !errors.Is(err, ErrOffset) {
		t.Fatalf("negative offset want ErrOffset, got %v", err)
	}
}

// TestFileWriteAtNegative covers the negative-offset guard of WriteAt.
func TestFileWriteAtNegative(t *testing.T) {
	fl := newFile("f", 0)
	if _, err := fl.WriteAt(-1, []byte("x"), 0); !errors.Is(err, ErrOffset) {
		t.Fatalf("negative writeAt want ErrOffset, got %v", err)
	}
}

// TestFileMarkBadNegative covers length clamping and clearing of bad blocks.
func TestFileMarkBadNegative(t *testing.T) {
	fl := newFile("f", 0)
	fl.markBad(0, -5)
	if v, ok := fl.bad[0]; !ok || v != 0 {
		t.Fatalf("negative length should clamp to 0, got %d (ok=%v)", v, ok)
	}
	fl.clearBad()
	if len(fl.bad) != 0 {
		t.Fatalf("clearBad should empty the map")
	}
}

// TestHandleReadDirOnFile covers the ErrNotDir branch of a file handle ReadDir.
func TestHandleReadDirOnFile(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("x"))
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if _, err := h.ReadDir(); !errors.Is(err, ErrNotDir) {
		t.Fatalf("readdir on file want ErrNotDir, got %v", err)
	}
}

// TestHandleReadAfterClose covers the closed-branch of a file handle Read.
func TestHandleReadAfterClose(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("x"))
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if err := h.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	if _, err := h.Read(make([]byte, 1)); !errors.Is(err, ErrClosed) {
		t.Fatalf("read after close want ErrClosed, got %v", err)
	}
}
