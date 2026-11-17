/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import "io"

// Handle is an open file or directory. For files it tracks a read/write offset;
// for directories it exposes a sorted listing. All I/O charges virtual latency
// and, for writes, enforces the file-system capacity limits. Handle implements
// io.Seeker for files so offsets can be moved relative to the start, current
// position or end of the file.
type Handle struct {
	fs     *FileSystem
	file   *File
	dir    *Dir
	offset int64
	closed bool
}

// Close marks the handle closed.
func (h *Handle) Close() error {
	h.closed = true
	return nil
}

// IsDir reports whether the handle targets a directory.
func (h *Handle) IsDir() bool { return h.dir != nil }

// Name returns the entry name.
func (h *Handle) Name() string {
	if h.file != nil {
		return h.file.Name()
	}
	return h.dir.Name()
}

// Size returns the byte size of the entry.
func (h *Handle) Size() int64 {
	if h.file != nil {
		return h.file.Size()
	}
	return h.dir.Size()
}

// ModTime returns the virtual modification time in nanoseconds.
func (h *Handle) ModTime() int64 {
	if h.file != nil {
		return h.file.ModTime()
	}
	return h.dir.ModTime()
}

// Seek moves the file offset, honoring the standard io.Seeker conventions. A
// resulting negative position is rejected. Directories do not support seeking.
func (h *Handle) Seek(offset int64, whence int) (int64, error) {
	if h.closed {
		return 0, ErrClosed
	}
	if h.file == nil {
		return 0, ErrNotFile
	}
	var base int64
	switch whence {
	case io.SeekStart:
		base = 0
	case io.SeekCurrent:
		base = h.offset
	case io.SeekEnd:
		base = h.file.Size()
	default:
		return 0, ErrOffset
	}
	target := base + offset
	if target < 0 {
		return 0, ErrOffset
	}
	h.offset = target
	return target, nil
}

// Read copies up to len(p) bytes from the current offset into p.
func (h *Handle) Read(p []byte) (int, error) {
	if h.closed {
		return 0, ErrClosed
	}
	if h.file == nil {
		return 0, ErrNotFile
	}
	h.fs.ioDelay()
	data, err := h.file.ReadAt(h.offset, len(p))
	if err != nil {
		return 0, err
	}
	n := copy(p, data)
	h.offset += int64(n)
	return n, nil
}

// Write appends p at the current offset, enforcing per-file and aggregate
// capacity limits before committing the bytes.
func (h *Handle) Write(p []byte) (int, error) {
	if h.closed {
		return 0, ErrClosed
	}
	if h.file == nil {
		return 0, ErrNotFile
	}
	h.fs.ioDelay()
	end := h.offset + int64(len(p))
	newActual := end
	cur := h.file.Size()
	if cur > newActual {
		newActual = cur
	}
	if newActual > MaxFileBytes {
		return 0, ErrQuotaExceeded
	}
	if h.fs.total+(newActual-cur) > MaxTotalBytes {
		return 0, ErrQuotaExceeded
	}
	h.fs.total += newActual - cur
	if _, err := h.file.WriteAt(h.offset, p, h.fs.clk.Now().UnixNano()); err != nil {
		return 0, err
	}
	h.offset += int64(len(p))
	return len(p), nil
}

// ReadDir returns the sorted child names of a directory handle.
func (h *Handle) ReadDir() ([]string, error) {
	if h.closed {
		return nil, ErrClosed
	}
	if h.dir == nil {
		return nil, ErrNotDir
	}
	h.fs.ioDelay()
	return h.dir.Names(), nil
}
