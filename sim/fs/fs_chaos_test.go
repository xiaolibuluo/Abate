/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

import (
	"encoding/json"
	"fmt"
	"testing"

	"abate/internal/determinism"
	"abate/sim/clock"
)

// fsSnapshot serializes the tree into a deterministic sorted map of path to
// signed size (negative marks a directory), which json.Marshal orders by key.
func fsSnapshot(f *FileSystem) []byte {
	m := make(map[string]int64)
	var walk func(d *Dir, prefix string)
	walk = func(d *Dir, prefix string) {
		for _, name := range d.Names() {
			p := prefix + "/" + name
			c := d.children[name]
			if fl, ok := c.(*File); ok {
				m[p] = fl.Size()
			} else {
				m[p] = -1
				walk(c.(*Dir), p)
			}
		}
	}
	walk(f.root, "")
	data, _ := json.Marshal(m)
	return data
}

// TestFsChaosDeterminism certifies the file system rebuilds identically across
// one hundred independent runs under the same seed.
func TestFsChaosDeterminism(t *testing.T) {
	determinism.MustMatch100(t, 1, func() []byte {
		clk := clock.NewVirtualClock()
		f := NewFileSystem(clk)
		if err := f.MkdirAll("/data"); err != nil {
			panic(err)
		}
		if err := f.MkdirAll("/data/shard"); err != nil {
			panic(err)
		}
		for i := 0; i < 24; i++ {
			path := fmt.Sprintf("/data/shard/f%02d", i)
			h, err := f.Create(path)
			if err != nil {
				panic(err)
			}
			payload := make([]byte, 32+i*7)
			for j := range payload {
				payload[j] = byte(i * 31)
			}
			if _, err := h.Write(payload); err != nil {
				panic(err)
			}
			_ = h.Close()
		}
		// Read a sample file back to exercise the read path.
		h, err := f.Open("/data/shard/f05")
		if err != nil {
			panic(err)
		}
		buf := make([]byte, 64)
		_, _ = h.Read(buf)
		_ = h.Close()
		// Remove one file to vary the tail state.
		if err := f.Remove("/data/shard/f10"); err != nil {
			panic(err)
		}
		return fsSnapshot(f)
	})
}

// TestFsRouteDeterminism confirms repeated resolution of the same path yields a
// byte-identical snapshot, a weaker but independent determinism check.
func TestFsRouteDeterminism(t *testing.T) {
	determinism.MustMatch100(t, 2, func() []byte {
		clk := clock.NewVirtualClock()
		f := NewFileSystem(clk)
		_ = f.MkdirAll("/a")
		_ = f.MkdirAll("/a/b")
		h, err := f.Create("/a/b/file")
		if err != nil {
			panic(err)
		}
		if _, err := h.Write([]byte("deterministic")); err != nil {
			panic(err)
		}
		_ = h.Close()
		st, err := f.Stat("/a/b/file")
		if err != nil {
			panic(err)
		}
		return []byte(fmt.Sprintf("%s:%d", st.Name(), st.Size()))
	})
}
