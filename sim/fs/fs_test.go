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
	"strings"
	"testing"

	"abate/sim/clock"
)

// newFS builds a file system on a fresh virtual clock.
func newFS() (*FileSystem, *clock.VirtualClock) {
	clk := clock.NewVirtualClock()
	return NewFileSystem(clk), clk
}

// mkdirParents creates every missing parent directory in path order.
func mkdirParents(t *testing.T, f *FileSystem, dir string) {
	t.Helper()
	cur := ""
	for _, part := range strings.Split(dir, "/") {
		if part == "" {
			continue
		}
		cur += "/" + part
		if err := f.MkdirAll(cur); err != nil && !errors.Is(err, ErrExist) {
			t.Fatalf("mkdir %s: %v", cur, err)
		}
	}
}

// writeAll creates a file at path and writes the full payload.
func writeAll(t *testing.T, f *FileSystem, path string, data []byte) {
	t.Helper()
	if i := strings.LastIndexByte(path, '/'); i > 0 {
		mkdirParents(t, f, path[:i])
	}
	h, err := f.Create(path)
	if err != nil {
		t.Fatalf("create %s: %v", path, err)
	}
	if _, err := h.Write(data); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
	if err := h.Close(); err != nil {
		t.Fatalf("close %s: %v", path, err)
	}
}

func TestCreateAndRead(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/hello.txt", []byte("world"))

	h, err := f.Open("/hello.txt")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if h.IsDir() {
		t.Fatalf("file reported as dir")
	}
	buf := make([]byte, 16)
	n, err := h.Read(buf)
	if err != nil || n != 5 || string(buf[:n]) != "world" {
		t.Fatalf("read = %q,%v (n=%d)", buf[:n], err, n)
	}
	if h.Size() != 5 {
		t.Fatalf("size = %d", h.Size())
	}
}

func TestPathNormalizationRejectsEscape(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/a"); err != nil {
		t.Fatalf("mkdir /a: %v", err)
	}
	writeAll(t, f, "/a/b.txt", []byte("x"))

	// A ".." that would leave the synthetic root is rejected.
	if _, err := f.Open("/a/../../etc/passwd"); !errors.Is(err, ErrEscapeRoot) {
		t.Fatalf("want ErrEscapeRoot, got %v", err)
	}
	// Non-absolute and empty paths are invalid.
	if _, err := f.Open("relative"); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("want ErrInvalidPath, got %v", err)
	}
	if _, err := f.Open(""); !errors.Is(err, ErrInvalidPath) {
		t.Fatalf("want ErrInvalidPath for empty, got %v", err)
	}
	// Collapsed slashes and "." are tolerated.
	if _, err := f.Open("/a/./b.txt"); err != nil {
		t.Fatalf("normalized path should open: %v", err)
	}
}

func TestMkdirAndReaddirOrdered(t *testing.T) {
	f, _ := newFS()
	for _, name := range []string{"zeta", "alpha", "mid"} {
		if err := f.MkdirAll("/d/" + name); err != nil && !errors.Is(err, ErrNoEntry) {
			if err := f.MkdirAll("/d"); err != nil {
				t.Fatalf("mkdir d: %v", err)
			}
		}
	}
	_ = f.MkdirAll("/d")
	for _, name := range []string{"zeta", "alpha", "mid"} {
		if err := f.MkdirAll("/d/" + name); err != nil {
			t.Fatalf("mkdir %s: %v", name, err)
		}
	}
	names, err := f.Readdir("/d")
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}
	want := []string{"alpha", "mid", "zeta"}
	if len(names) != len(want) {
		t.Fatalf("readdir = %v", names)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("readdir order = %v, want %v", names, want)
		}
	}
}

func TestCreateErrors(t *testing.T) {
	f, _ := newFS()
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if _, err := f.Create("/d"); !errors.Is(err, ErrExist) {
		t.Fatalf("create over dir want ErrExist, got %v", err)
	}
	if _, err := f.Create("/nope/f.txt"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("create missing parent want ErrNoEntry, got %v", err)
	}
	if _, err := f.Open("/missing"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("open missing want ErrNoEntry, got %v", err)
	}
}

func TestDeleteFileAndDir(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/a.txt", make([]byte, 100))
	if err := f.MkdirAll("/d"); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeAll(t, f, "/d/b.txt", make([]byte, 50))

	if err := f.Remove("/a.txt"); err != nil {
		t.Fatalf("delete file: %v", err)
	}
	if _, err := f.Open("/a.txt"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("file should be gone, got %v", err)
	}
	if err := f.Remove("/d"); err != nil {
		t.Fatalf("delete dir: %v", err)
	}
	if _, err := f.Open("/d/b.txt"); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("nested file should be gone, got %v", err)
	}
}

func TestOffsetAndSeek(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/f", []byte("abcdef"))
	h, err := f.Open("/f")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if _, err := h.Seek(2, io.SeekStart); err != nil {
		t.Fatalf("seek: %v", err)
	}
	buf := make([]byte, 2)
	n, err := h.Read(buf)
	if err != nil || n != 2 || string(buf) != "cd" {
		t.Fatalf("read at offset = %q,%v", buf[:n], err)
	}
	if _, err := h.Seek(-1, io.SeekStart); !errors.Is(err, ErrOffset) {
		t.Fatalf("negative seek want ErrOffset, got %v", err)
	}
}

func TestWriteAfterClose(t *testing.T) {
	f, _ := newFS()
	h, err := f.Create("/x")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := h.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	if _, err := h.Write([]byte("y")); !errors.Is(err, ErrClosed) {
		t.Fatalf("write after close want ErrClosed, got %v", err)
	}
}

func TestReadDirOnFileFails(t *testing.T) {
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

func TestWalkDeterministic(t *testing.T) {
	f, _ := newFS()
	writeAll(t, f, "/a/1.txt", []byte("1"))
	writeAll(t, f, "/a/2.txt", []byte("2"))
	writeAll(t, f, "/b.txt", []byte("3"))
	writeAll(t, f, "/c/d/3.txt", []byte("4"))

	var got []string
	err := f.Walk("/", func(path string) error {
		got = append(got, path)
		return nil
	})
	if err != nil {
		t.Fatalf("walk: %v", err)
	}
	// Lexicographic depth-first order is fully determined by child names.
	want := []string{"/", "/a", "/a/1.txt", "/a/2.txt", "/b.txt", "/c", "/c/d", "/c/d/3.txt"}
	if len(got) != len(want) {
		t.Fatalf("walk = %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("walk[%d] = %q, want %q (full=%v)", i, got[i], want[i], got)
		}
	}

	// A missing root is reported as ErrNoEntry.
	if err := f.Walk("/nope", func(string) error { return nil }); !errors.Is(err, ErrNoEntry) {
		t.Fatalf("walk missing root want ErrNoEntry, got %v", err)
	}

	// The first error from fn aborts the walk deterministically.
	called := 0
	stop := errors.New("stop")
	werr := f.Walk("/", func(string) error {
		called++
		return stop
	})
	if !errors.Is(werr, stop) || called != 1 {
		t.Fatalf("walk did not stop at first fn error: err=%v called=%d", werr, called)
	}
}
