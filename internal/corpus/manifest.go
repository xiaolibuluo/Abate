/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

// Manifest is the complete registry of upstream source files, ordered by
// RelPath ascending and free of timestamps and absolute paths.
type Manifest struct {
	SchemaVersion int     `json:"schemaVersion"`
	Entries       []Entry `json:"entries"`
	digest        string
}

// Load decodes a manifest from disk and validates its ordering. The digest is
// computed from the raw file bytes so gate_e2e.sh can recompute it offline.
func Load(path string) (Manifest, error) {
	var m Manifest
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return m, fmt.Errorf("%w: %s", errManifestUnreadable, path)
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return m, fmt.Errorf("%w: %s", errManifestMalformed, path)
	}
	for i := 1; i < len(m.Entries); i++ {
		if m.Entries[i-1].RelPath >= m.Entries[i].RelPath {
			return m, fmt.Errorf("%w: %s then %s", errUnsortedManifest,
				m.Entries[i-1].RelPath, m.Entries[i].RelPath)
		}
	}
	m.digest = hashBytes(data)
	return m, nil
}

// Sorted returns a copy of the entries ordered by RelPath ascending.
func (m Manifest) Sorted() []Entry {
	out := make([]Entry, len(m.Entries))
	copy(out, m.Entries)
	sort.Slice(out, func(i, j int) bool {
		return out[i].RelPath < out[j].RelPath
	})
	return out
}

// Digest returns the lowercase SHA-256 of the canonical manifest encoding.
func (m Manifest) Digest() string {
	if m.digest != "" {
		return m.digest
	}
	return hashBytes(encodeManifest(m))
}

// encodeManifest produces the exact bytes written by corpusctl register.
func encodeManifest(m Manifest) []byte {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return []byte{}
	}
	return data
}

// withDigest attaches the digest of the canonical encoding to a manifest.
func withDigest(m Manifest) Manifest {
	m.digest = hashBytes(encodeManifest(m))
	return m
}
