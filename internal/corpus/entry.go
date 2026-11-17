/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"path/filepath"
)

// Entry is one real upstream source file together with its provenance and the
// SHA-256 digest of its original bytes.
type Entry struct {
	Language      string   `json:"language"`
	Project       string   `json:"project"`
	UpstreamURL   string   `json:"upstreamURL"`
	Commit        string   `json:"commit"`
	SPDX          string   `json:"spdx"`
	SPDXEvidence  string   `json:"spdxEvidence"`
	RelPath       string   `json:"relPath"`
	Bytes         int64    `json:"bytes"`
	SHA256        string   `json:"sha256"`
	ConstructTags []string `json:"constructTags"`
	Source        Source   `json:"source"`
}

// hashBytes returns the lowercase hexadecimal SHA-256 of the original bytes.
func hashBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// fileDigest reads one read-only file and returns its digest and byte count.
// Using fs.ReadFile keeps the read inside the corpus root and avoids G304.
func fileDigest(fsys fs.FS, relPath string) (string, int64, error) {
	data, err := fs.ReadFile(fsys, relPath)
	if err != nil {
		return "", 0, err
	}
	return hashBytes(data), int64(len(data)), nil
}

// languageOfExt maps a file extension to its registration language, returning
// an empty string for files that are not one of the seven source languages.
func languageOfExt(relPath string) string {
	return extensionLanguage[filepath.Ext(relPath)]
}
