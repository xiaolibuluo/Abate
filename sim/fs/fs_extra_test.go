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

// TestRootPath exercises resolution of the synthetic root itself.
func TestRootPath(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("mkdir root want ErrInvalidPath, got %v", err)
	}
	if _, err := f.Create("/"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("create root want ErrInvalidPath, got %v", err)
	}
	if err := f.Remove("/"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("delete root want ErrInvalidPath, got %v", err)
	}
	h, err := f.Open("/")
	if err != nil {
		t.Fatalf("open root: %v", err)
	}
	if !h.IsDir() {
		t.Fatalf("root should be a directory")
	}
	if h.Name() != "/" {
		t.Fatalf("root name = %q", h.Name())
	}
}

// TestIntermediateNotDir confirms walking through a file is rejected.
func TestIntermediateNotDir(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("x"))
	if _, err := f.Open("/f/x"); !errors.Is(err, ErrNotDir) {
		t.Fatalf("walk through file want ErrNotDir, got %v", err)
	}
}

// TestReaddirErrors covers the missing-entry and directory-handle branches.
func TestReaddirErrors(t *testing.T) {
	f, _ := newFS()
	if _, err := f.Readdir("/missing"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("readdir missing want ErrNoEntry, got %v", err)
	}
	// Root readdir on an empty file system yields no entries.
	names, err := f.Readdir("/")
	if err != nil || len(names) != 0 {
		t.Fatalf("root readdir = %v,%v", names, err)
	}
}

// TestHandleDirAccessors covers the directory-side branches of handle methods.
func TestHandleDirAccessors(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	h, err := f.Open("/d")
	if err != nil {
		t.Fatalf("open dir: %v", err)
	}
	if h.Size() != 0 {
		t.Fatalf("empty dir size = %d", h.Size())
	}
	if h.ModTime() <= 0 {
		t.Fatalf("dir mtime = %d", h.ModTime())
	}
	buf := make([]byte, 4)
	if _, err := h.Read(buf); !errors.Is(err, ErrNotFile) {
		t.Fatalf("read on dir want ErrNotFile, got %v", err)
	}
	if _, err := h.Write([]byte("x")); !errors.Is(err, ErrNotFile) {
		t.Fatalf("write on dir want ErrNotFile, got %v", err)
	}
	// ReadDir after close reports ErrClosed.
	if err := h.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	if _, err := h.ReadDir(); !errors.Is(err, ErrClosed) {
		t.Fatalf("readdir after close want ErrClosed, got %v", err)
	}
}

// TestHandleFileModTime covers the file-side ModTime accessor.
func TestHandleFileModTime(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("abc"))
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if h.ModTime() <= 0 {
		t.Fatalf("file mtime = %d", h.ModTime())
	}
}

// TestFaultRecoverErrors covers fault recovery on missing and directory entries.
func TestFaultRecoverErrors(t *testing.T) {
	f, _ := newFS()
	if err := f.RecoverFault("/ghost"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("recover missing want ErrNoEntry, got %v", err)
	}
	if err := f.RecoverFault("/"); err != nil {
		t.Fatalf("recover on dir should be a no-op: %v", err)
	}
}

// TestFaultLatencyNegative rejects a negative latency value.
func TestFaultLatencyNegative(t *testing.T) {
	f, _ := newFS()
	if err := f.InjectFault(Fault{Kind: FaultLatency, Value: -1}); !errors.Is(err, ErrOffset) {
		t.Fatalf("negative latency want ErrOffset, got %v", err)
	}
}

// TestDeleteMissing covers the missing-entry branch of Delete.
func TestDeleteMissing(t *testing.T) {
	f, _ := newFS()
	if err := f.Remove("/ghost"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("delete missing want ErrNoEntry, got %v", err)
	}
}
