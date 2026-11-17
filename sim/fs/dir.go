/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import "sort"

// Dir is a virtual directory holding named entries in a map. Names are returned
// sorted to keep directory listings deterministic regardless of map traversal.
type Dir struct {
	name     string
	mtime    int64
	children map[string]fsNode
}

// newDir builds an empty directory stamped with the given virtual time.
func newDir(name string, now int64) *Dir {
	return &Dir{name: name, mtime: now, children: make(map[string]fsNode)}
}

// insert adds or replaces a child entry and bumps the directory mtime.
func (d *Dir) insert(name string, n fsNode) {
	d.children[name] = n
	d.mtime++
}

// remove deletes a child entry by name.
func (d *Dir) remove(name string) { delete(d.children, name) }

// lookup returns the named child or nil when absent.
func (d *Dir) lookup(name string) fsNode { return d.children[name] }

// Names returns the child names in lexicographic order.
func (d *Dir) Names() []string {
	out := make([]string, 0, len(d.children))
	for k := range d.children {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// treeSize sums the bytes of every file beneath this directory.
func (d *Dir) treeSize() int64 {
	var s int64
	for _, c := range d.children {
		s += c.nodeSize()
	}
	return s
}

// Size returns the aggregated byte size of the subtree.
func (d *Dir) Size() int64 { return d.treeSize() }

// Name returns the directory name.
func (d *Dir) Name() string { return d.name }

// ModTime returns the virtual modification time in nanoseconds.
func (d *Dir) ModTime() int64 { return d.mtime }

// nodeSize satisfies the fsNode contract.
func (d *Dir) nodeSize() int64 { return d.treeSize() }
