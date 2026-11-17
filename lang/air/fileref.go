/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// FileRef describes one source file bound to a compilation unit.
//
// The file content itself is never stored in AIR; the content hash is enough
// to decide whether an incremental parse may reuse the existing subtree.
type FileRef struct {
	// Path is the source path, used as the canonical file ordering key.
	Path string
	// ContentHash is the SHA-256 digest of the original source bytes.
	ContentHash [32]byte
	// LineEnding is the dominant line terminator: "\n", "\r\n", or "\r".
	LineEnding string
	// Encoding is the detected source encoding label.
	Encoding string
	// Size is the byte length of the original source.
	Size int
}

// NewFileRef derives file metadata from the in-memory source bytes.
func NewFileRef(path string, content []byte) FileRef {
	ending := "\n"
	switch {
	case bytes.Contains(content, []byte("\r\n")):
		ending = "\r\n"
	case bytes.Contains(content, []byte("\r")):
		ending = "\r"
	}
	encoding := "utf-8"
	if bytes.HasPrefix(content, []byte{0xEF, 0xBB, 0xBF}) {
		encoding = "utf-8-bom"
	}
	return FileRef{
		Path:        path,
		ContentHash: sha256.Sum256(content),
		LineEnding:  ending,
		Encoding:    encoding,
		Size:        len(content),
	}
}

// Validate checks the invariants required before serialization.
func (f FileRef) Validate() error {
	if f.Path == "" || f.Size < 0 {
		return ErrMalformedUnit
	}
	return nil
}

// Digest returns the lowercase hexadecimal content hash.
func (f FileRef) Digest() string {
	return hex.EncodeToString(f.ContentHash[:])
}

// Canonical returns a stable comparison key for the file reference.
func (f FileRef) Canonical() string {
	return f.Path + "\x00" + f.Digest() + "\x00" + f.LineEnding +
		"\x00" + f.Encoding + "\x00" + strconv.Itoa(f.Size)
}
