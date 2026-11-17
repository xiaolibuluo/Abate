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

	"abate/sim/clock"
)

// TestMkdirAllErrorCodes locks every rejection path of MkdirAll.
func TestMkdirAllErrorCodes(t *testing.T) {
	t.Parallel()
	f, _ := newFS()
	for _, path := range []string{"/a/../../etc", "/../outside"} {
		if err := f.MkdirAll(path); !errors.Is(err, ErrEscapeRoot) {
			t.Fatalf("MkdirAll(%q) = %v, want ErrEscapeRoot", path, err)
		}
	}
	for _, path := range []string{"relative", ""} {
		if err := f.MkdirAll(path); !errors.Is(err, ErrInvalidPath) {
			t.Fatalf("MkdirAll(%q) = %v, want ErrInvalidPath", path, err)
		}
	}
	if err := f.MkdirAll("/"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("MkdirAll(root) = %v", err)
	}
	if err := f.MkdirAll("/x/y"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("MkdirAll(missing parent) = %v", err)
	}
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("MkdirAll(/d): %v", err)
	}
	if err := f.MkdirAll("/d"); !errors.Is(err, ErrExist) {
		t.Fatalf("duplicate MkdirAll = %v", err)
	}
	if err := f.MkdirAll("/d/sub"); err != nil {
		t.Fatalf("nested MkdirAll: %v", err)
	}
	writeAll(t, f, "/f", []byte("x"))
	if err := f.MkdirAll("/f"); !errors.Is(err, ErrExist) {
		t.Fatalf("MkdirAll over file = %v", err)
	}
	if err := f.MkdirAll("/f/x"); !errors.Is(err, ErrNotDir) {
		t.Fatalf("MkdirAll through file = %v", err)
	}
}

// TestRemoveErrorCodes covers the resolution and existence branches of Remove.
func TestRemoveErrorCodes(t *testing.T) {
	t.Parallel()
	f, _ := newFS()
	if err := f.Remove("/../x"); !errors.Is(err, ErrEscapeRoot) {
		t.Fatalf("Remove escape = %v", err)
	}
	if err := f.Remove("relative"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("Remove relative = %v", err)
	}
	if err := f.Remove("/"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("Remove root = %v", err)
	}
	if err := f.Remove("/x/y"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("Remove missing parent = %v", err)
	}
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeAll(t, f, "/d/f", []byte("abc"))
	if err := f.Remove("/d"); err != nil {
		t.Fatalf("remove subtree: %v", err)
	}
	if err := f.Remove("/d"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("repeated remove = %v", err)
	}
	if _, err := f.Open("/d/f"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("child survived parent removal: %v", err)
	}
}

// TestWalkOrderAndErrors proves lexicographic traversal and error propagation.
func TestWalkOrderAndErrors(t *testing.T) {
	t.Parallel()
	f, _ := newFS()
	writeAll(t, f, "/z/2.txt", []byte("2"))
	writeAll(t, f, "/a/1.txt", []byte("1"))
	writeAll(t, f, "/a/0.txt", []byte("0"))
	var got []string
	if err := f.Walk("/", func(path string) error {
		got = append(got, path)
		return nil
	}); err != nil {
		t.Fatalf("walk: %v", err)
	}
	want := []string{"/", "/a", "/a/0.txt", "/a/1.txt", "/z", "/z/2.txt"}
	if len(got) != len(want) {
		t.Fatalf("walk = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("walk order = %v, want %v", got, want)
		}
	}
	if err := f.Walk("/../x", func(string) error { return nil }); !errors.Is(err, ErrEscapeRoot) {
		t.Fatalf("walk escape = %v", err)
	}
	if err := f.Walk("relative", func(string) error { return nil }); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("walk relative = %v", err)
	}
	// A file is a valid walk root with no descendants.
	called := 0
	if err := f.Walk("/a/0.txt", func(string) error {
		called++
		return nil
	}); err != nil || called != 1 {
		t.Fatalf("file-root walk = %d calls, err=%v", called, err)
	}
	// An error raised deep in the tree propagates back through every recursion level.
	stop := errors.New("stop deep")
	visited := 0
	err := f.Walk("/", func(string) error {
		visited++
		if visited == 4 {
			return stop
		}
		return nil
	})
	if !errors.Is(err, stop) || visited != 4 {
		t.Fatalf("deep walk error = %v after %d visits", err, visited)
	}
}

// TestReaddirAndFaultErrorCodes covers the remaining resolution branches.
func TestReaddirAndFaultErrorCodes(t *testing.T) {
	t.Parallel()
	f, _ := newFS()
	if _, err := f.Readdir("/../x"); !errors.Is(err, ErrEscapeRoot) {
		t.Fatalf("Readdir escape = %v", err)
	}
	if _, err := f.Readdir("relative"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("Readdir relative = %v", err)
	}
	if err := f.InjectFault(Fault{Kind: FaultBadBlock, Path: "relative"}); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("InjectFault relative = %v", err)
	}
	if err := f.RecoverFault("relative"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("RecoverFault relative = %v", err)
	}
	if err := f.InjectFault(Fault{Kind: FaultLatency, Value: 0}); err != nil {
		t.Fatalf("zero latency: %v", err)
	}
	if err := f.InjectFault(Fault{Kind: FaultBadBlock, Path: "/"}); !errors.Is(err, ErrNotFile) {
		t.Fatalf("bad block on root = %v", err)
	}
}

// TestHandleSeekBranches covers every whence and the invalid-seek branches.
func TestHandleSeekBranches(t *testing.T) {
	t.Parallel()
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("abcdef"))
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if at, err := h.Seek(2, io.SeekCurrent); err != nil || at != 2 {
		t.Fatalf("SeekCurrent = %d, %v", at, err)
	}
	if at, err := h.Seek(-1, io.SeekEnd); err != nil || at != 5 {
		t.Fatalf("SeekEnd = %d, %v", at, err)
	}
	if _, err := h.Seek(0, 99); !errors.Is(err, ErrOffset) {
		t.Fatalf("invalid whence = %v", err)
	}
	if err := h.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	if _, err := h.Seek(0, io.SeekStart); !errors.Is(err, ErrClosed) {
		t.Fatalf("seek after close = %v", err)
	}
	dir, err := f.Open("/")
	if err != nil {
		t.Fatalf("open root: %v", err)
	}
	if _, err := dir.Seek(0, io.SeekStart); !errors.Is(err, ErrNotFile) {
		t.Fatalf("seek on dir = %v", err)
	}
}

// TestStatAndIODelayBranches covers the Stat alias and latency charging.
func TestStatAndIODelayBranches(t *testing.T) {
	t.Parallel()
	clk := clock.NewVirtualClock()
	f := NewFileSystem(clk)
	writeAll(t, f, "/f", []byte("abc"))
	st, err := f.Stat("/f")
	if err != nil || st.Size() != 3 {
		t.Fatalf("Stat = %v, %v", st, err)
	}
	if _, err := f.Stat("/../x"); !errors.Is(err, ErrEscapeRoot) {
		t.Fatalf("Stat escape = %v", err)
	}
	before := clk.Now().UnixNano()
	if err := f.InjectFault(Fault{Kind: FaultLatency, Value: 777}); err != nil {
		t.Fatalf("inject: %v", err)
	}
	if _, err := f.Open("/f"); err != nil {
		t.Fatalf("open: %v", err)
	}
	if clk.Now().UnixNano()-before != 777 {
		t.Fatalf("IO latency not charged: %d", clk.Now().UnixNano()-before)
	}
}
