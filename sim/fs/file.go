/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

// File is a virtual file holding an in-memory byte slice and a set of faulted
// (bad-block) byte ranges used by the fault injector.
type File struct {
	name  string
	data  []byte
	mtime int64
	bad   map[int64]int64
}

// newFile builds an empty file stamped with the given virtual modification time.
func newFile(name string, now int64) *File {
	return &File{name: name, mtime: now, bad: make(map[int64]int64)}
}

// Size returns the current byte length of the file.
func (fl *File) Size() int64 { return int64(len(fl.data)) }

// Name returns the file name.
func (fl *File) Name() string { return fl.name }

// ModTime returns the virtual modification time in nanoseconds.
func (fl *File) ModTime() int64 { return fl.mtime }

// markBad records a faulted byte range starting at off of the given length.
func (fl *File) markBad(off, length int64) {
	if length < 0 {
		length = 0
	}
	fl.bad[off] = length
}

// clearBad removes all recorded bad-block ranges.
func (fl *File) clearBad() { fl.bad = make(map[int64]int64) }

// ReadAt returns the bytes in [off, off+n), failing if the range intersects a
// faulted block or starts out of bounds.
func (fl *File) ReadAt(off int64, n int) ([]byte, error) {
	if off < 0 || off > fl.Size() {
		return nil, ErrOffset
	}
	end := off + int64(n)
	if end > fl.Size() {
		end = fl.Size()
	}
	for b, blen := range fl.bad {
		if b < end && b+blen > off {
			return nil, ErrBadBlock
		}
	}
	out := make([]byte, end-off)
	copy(out, fl.data[off:end])
	return out, nil
}

// WriteAt writes p at off, growing the file as needed. The mutation timestamp is
// updated to now on a non-empty write.
func (fl *File) WriteAt(off int64, p []byte, now int64) (int64, error) {
	if off < 0 {
		return 0, ErrOffset
	}
	end := off + int64(len(p))
	if end > fl.Size() {
		grown := make([]byte, end)
		copy(grown, fl.data)
		fl.data = grown
	}
	copy(fl.data[off:end], p)
	if len(p) > 0 {
		fl.mtime = now
	}
	return int64(len(p)), nil
}

// nodeSize satisfies the fsNode contract.
func (fl *File) nodeSize() int64 { return fl.Size() }
