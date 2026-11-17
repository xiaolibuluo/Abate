/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import (
	"errors"
	"io"
	"testing"
)

// TestHandleReadDirNormal covers the success path of a directory handle.
func TestHandleReadDirNormal(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeAll(t, f, "/d/a", []byte("x"))
	writeAll(t, f, "/d/b", []byte("y"))
	h, err := f.Open("/d")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	names, err := h.ReadDir()
	if err != nil || len(names) != 2 {
		t.Fatalf("readdir = %v,%v", names, err)
	}
}

// TestMkdirMissingParent covers the resolve-error branch of Mkdir.
func TestMkdirMissingParent(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/x/y"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("mkdir missing parent want ErrNoEntry, got %v", err)
	}
}

// TestDeleteMissingParent covers the resolve-error branch of Delete.
func TestDeleteMissingParent(t *testing.T) {
	f, _ := newFS()
	if err := f.Remove("/x/y"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("delete missing parent want ErrNoEntry, got %v", err)
	}
}

// TestReaddirOnFile covers the ErrNotDir branch of the FileSystem.Readdir method.
func TestReaddirOnFile(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("x"))
	if _, err := f.Readdir("/f"); !errors.Is(err, ErrNotDir) {
		t.Fatalf("readdir on file want ErrNotDir, got %v", err)
	}
}

// TestInjectBadBlockOnDir covers the ErrNotFile branch of InjectFault.
func TestInjectBadBlockOnDir(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := f.InjectFault(Fault{Kind: FaultBadBlock, Path: "/d", Offset: 0, Value: 1}); !errors.Is(err, ErrNotFile) {
		t.Fatalf("bad-block on dir want ErrNotFile, got %v", err)
	}
}

// TestRecoverMissingParent covers the resolve-error branch of RecoverFault.
func TestRecoverMissingParent(t *testing.T) {
	f, _ := newFS()
	if err := f.RecoverFault("/x/y"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("recover missing parent want ErrNoEntry, got %v", err)
	}
}

// TestWriteOverwriteMiddle covers the no-growth branch when writing inside a file.
func TestWriteOverwriteMiddle(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("abcdefghij"))
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if _, err := h.Seek(0, io.SeekStart); err != nil {
		t.Fatalf("seek: %v", err)
	}
	// Overwriting the first three bytes; the file must not shrink.
	if _, err := h.Write([]byte("XYZ")); err != nil {
		t.Fatalf("write: %v", err)
	}
	if h.Size() != 10 {
		t.Fatalf("overwrite shrank file: size=%d", h.Size())
	}
	if _, err := h.Seek(0, io.SeekStart); err != nil {
		t.Fatalf("seek: %v", err)
	}
	buf := make([]byte, 10)
	n, err := h.Read(buf)
	if err != nil || string(buf[:n]) != "XYZdefghij" {
		t.Fatalf("overwrite result = %q,%v", buf[:n], err)
	}
}
