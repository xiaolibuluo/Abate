/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"crypto/sha256"
	"slices"
)

// SourceFile is one in-memory source buffer handed to a frontend.
type SourceFile struct {
	// Path is the virtual path used for ordering and diagnostics.
	Path string
	// Content is the exact source byte stream.
	Content []byte
}

// NewSourceFile copies content so later caller mutations cannot change a parse.
func NewSourceFile(path string, content []byte) SourceFile {
	return SourceFile{Path: path, Content: slices.Clone(content)}
}

// Clone returns an independent copy of the source buffer.
func (f SourceFile) Clone() SourceFile {
	return SourceFile{Path: f.Path, Content: slices.Clone(f.Content)}
}

// Digest returns the SHA-256 digest of the source bytes.
func (f SourceFile) Digest() [32]byte {
	return sha256.Sum256(f.Content)
}

// Size returns the byte length of the source buffer.
func (f SourceFile) Size() int {
	return len(f.Content)
}
