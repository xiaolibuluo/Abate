/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// errDigestRead reports a configuration file whose digest could not be taken.
var errDigestRead = errors.New("report: digest source unreadable")

// DigestBytes returns the lowercase hexadecimal SHA-256 of one byte buffer.
func DigestBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// DigestFile returns the SHA-256 of the raw bytes of one relative file. The
// read is rooted in the file's directory, so no absolute path or variable
// os.Open is involved.
func DigestFile(path string) (string, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return "", fmt.Errorf("%w: %s: %w", errDigestRead, path, err)
	}
	return DigestBytes(data), nil
}
