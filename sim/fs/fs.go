/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package fs is a deterministic in-memory virtual file system used by the
// simulator to model durable storage without touching the host.
//
// Every mutation, quota check and fault is resolved through the injected
// virtual clock and explicit fault records, so the visible state is fully
// reproducible. Paths are absolute; any ".." component that would leave the
// synthetic root is rejected to prevent escape.
package fs

import (
	"errors"
	"strings"

	"abate/sim/clock"
)

// Capacity limits enforced by the virtual file system.
const (
	// MaxFileBytes caps a single file at 100 MiB.
	MaxFileBytes = int64(100 << 20)

	// MaxTotalBytes caps the aggregate of all files at 1 GiB.
	MaxTotalBytes = int64(1 << 30)

	// DefaultIOLatency is the virtual nanoseconds charged per operation.
	DefaultIOLatency = clock.Duration(1000)
)

// Sentinel errors of the virtual file system.
var (
	// ErrNoEntry reports a missing path.
	ErrNoEntry = errors.New("fs: no such entry")

	// ErrExist reports that a create target already exists.
	ErrExist = errors.New("fs: entry already exists")

	// ErrNotDir reports an operation that needs a directory on a file.
	ErrNotDir = errors.New("fs: not a directory")

	// ErrNotFile reports an operation that needs a file on a directory.
	ErrNotFile = errors.New("fs: not a file")

	// ErrIsDir reports that a path resolves to a directory where a file is needed.
	ErrIsDir = errors.New("fs: is a directory")

	// ErrEscapeRoot reports a path component attempting to leave the root.
	ErrEscapeRoot = errors.New("fs: path escapes root")

	// ErrInvalidPath reports a malformed (non-absolute or empty) path.
	ErrInvalidPath = errors.New("fs: invalid path")

	// ErrQuotaExceeded reports a write that would breach a capacity limit.
	ErrQuotaExceeded = errors.New("fs: quota exceeded")

	// ErrBadBlock reports a read that intersects a faulted region.
	ErrBadBlock = errors.New("fs: bad block")

	// ErrClosed reports I/O on a closed handle.
	ErrClosed = errors.New("fs: handle closed")

	// ErrOffset reports a negative or out-of-range seek offset.
	ErrOffset = errors.New("fs: invalid offset")

	// ErrUnknownFault reports an unrecognized fault kind.
	ErrUnknownFault = errors.New("fs: unknown fault kind")
)

// fsNode is the private common contract for directory entries. Concrete entries
// (*File, *Dir) also expose exported accessors used by handles.
type fsNode interface {
	nodeSize() int64
}

// resolved is the outcome of a path resolution: the terminal node, its parent
// directory, the terminal basename, and any resolution error.
type resolved struct {
	node   fsNode
	parent *Dir
	base   string
	err    error
}

// FileSystem is a deterministic in-memory tree rooted at "/".
//
// It enforces per-file and aggregate capacity limits, charges a virtual IO
// latency for every operation, and supports link-level bad-block faults. No
// real file descriptor, network socket or wall-clock read is ever used.
type FileSystem struct {
	clk       *clock.VirtualClock
	root      *Dir
	total     int64
	ioLatency clock.Duration
}

// NewFileSystem builds an empty file system driven by the given clock.
func NewFileSystem(clk *clock.VirtualClock) *FileSystem {
	return &FileSystem{
		clk:       clk,
		root:      newDir("/", clk.Now().UnixNano()),
		ioLatency: DefaultIOLatency,
	}
}

// normalize splits an absolute path into ordered components, collapsing "." and
// rejecting any ".." that would escape the synthetic root.
func (*FileSystem) normalize(path string) ([]string, error) {
	if path == "" || path[0] != '/' {
		return nil, ErrInvalidPath
	}
	parts := strings.Split(path, "/")
	out := make([]string, 0, len(parts))
	for _, p := range parts[1:] {
		if p == "" {
			continue
		}
		if p == "." {
			continue
		}
		if p == ".." {
			return nil, ErrEscapeRoot
		}
		out = append(out, p)
	}
	return out, nil
}

// resolve walks the path and returns the terminal node, its parent directory,
// and the terminal basename. The parent is nil for the root path.
func (f *FileSystem) resolve(path string) resolved {
	parts, err := f.normalize(path)
	if err != nil {
		return resolved{err: err}
	}
	if len(parts) == 0 {
		return resolved{node: f.root}
	}
	cur := f.root
	for i := 0; i < len(parts)-1; i++ {
		child := cur.lookup(parts[i])
		if child == nil {
			return resolved{err: ErrNoEntry}
		}
		next, ok := child.(*Dir)
		if !ok {
			return resolved{err: ErrNotDir}
		}
		cur = next
	}
	base := parts[len(parts)-1]
	return resolved{node: cur.lookup(base), parent: cur, base: base}
}

// ioDelay charges the configured virtual latency to the clock.
func (f *FileSystem) ioDelay() {
	if f.ioLatency > 0 {
		_, _ = f.clk.Advance(f.ioLatency)
	}
}

// Create makes an empty file at path and returns a writable handle.
func (f *FileSystem) Create(path string) (*Handle, error) {
	f.ioDelay()
	r := f.resolve(path)
	if r.err != nil {
		return nil, r.err
	}
	if r.parent == nil {
		return nil, ErrInvalidPath
	}
	if r.node != nil {
		return nil, ErrExist
	}
	fl := newFile(r.base, f.clk.Now().UnixNano())
	r.parent.insert(r.base, fl)
	return &Handle{fs: f, file: fl}, nil
}

// Open opens an existing file or directory for reading.
func (f *FileSystem) Open(path string) (*Handle, error) {
	f.ioDelay()
	r := f.resolve(path)
	if r.err != nil {
		return nil, r.err
	}
	if r.node == nil {
		return nil, ErrNoEntry
	}
	if d, ok := r.node.(*Dir); ok {
		return &Handle{fs: f, dir: d}, nil
	}
	return &Handle{fs: f, file: r.node.(*File)}, nil
}

// Stat returns a read-only handle for inspecting the entry at path.
func (f *FileSystem) Stat(path string) (*Handle, error) {
	return f.Open(path)
}

// MkdirAll creates a directory at path, creating any missing parents.
func (f *FileSystem) MkdirAll(path string) error {
	f.ioDelay()
	r := f.resolve(path)
	if r.err != nil {
		return r.err
	}
	if r.parent == nil {
		return ErrInvalidPath
	}
	if r.node != nil {
		return ErrExist
	}
	r.parent.insert(r.base, newDir(r.base, f.clk.Now().UnixNano()))
	return nil
}

// Remove removes a file or recursively deletes a directory at path.
func (f *FileSystem) Remove(path string) error {
	f.ioDelay()
	r := f.resolve(path)
	if r.err != nil {
		return r.err
	}
	if r.node == nil {
		return ErrNoEntry
	}
	if r.parent == nil {
		return ErrInvalidPath
	}
	f.total -= r.node.nodeSize()
	r.parent.remove(r.base)
	return nil
}

// Readdir lists the names of a directory in deterministic (sorted) order.
func (f *FileSystem) Readdir(path string) ([]string, error) {
	f.ioDelay()
	r := f.resolve(path)
	if r.err != nil {
		return nil, r.err
	}
	if r.node == nil {
		return nil, ErrNoEntry
	}
	d, ok := r.node.(*Dir)
	if !ok {
		return nil, ErrNotDir
	}
	return d.Names(), nil
}

// Walk traverses the directory tree rooted at path in deterministic lexicographic
// order, invoking fn once for the root and once for every descendant. The walk
// stops and returns the first error fn reports, matching the contract of the
// standard library filepath.Walk without touching the host filesystem.
func (f *FileSystem) Walk(root string, fn func(path string) error) error {
	r := f.resolve(root)
	if r.err != nil {
		return r.err
	}
	if r.node == nil {
		return ErrNoEntry
	}
	return f.walkNode(root, r.node, fn)
}

// walkNode is the recursive worker for Walk; it reports path then descends into
// directories in sorted child order so the emitted sequence is fully determined.
func (f *FileSystem) walkNode(path string, n fsNode, fn func(path string) error) error {
	if err := fn(path); err != nil {
		return err
	}
	d, ok := n.(*Dir)
	if !ok {
		return nil
	}
	for _, name := range d.Names() {
		child := d.lookup(name)
		childPath := strings.TrimSuffix(path, "/") + "/" + name
		if err := f.walkNode(childPath, child, fn); err != nil {
			return err
		}
	}
	return nil
}
