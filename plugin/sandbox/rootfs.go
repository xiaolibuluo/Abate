/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"path/filepath"
	"strings"
)

// RootFS is the OCI style root filesystem of the sandbox.
//
// 模型来源=gVisor runsc 的 rootfs bundle 语义，Abate 以纯数据表达，不挂载真实目录。
type RootFS struct {
	// Path is the absolute path of the rootfs bundle.
	Path string `json:"path"`

	// ReadOnly marks the rootfs bundle as immutable.
	ReadOnly bool `json:"readOnly"`

	// Digest is the expected content digest of the bundle.
	Digest string `json:"digest"`
}

// Mounts returns the absolute paths the rootfs exposes to the guest.
func (r RootFS) Mounts() []string {
	if r.Path == "" {
		return nil
	}
	return []string{r.Path}
}

// ValidPath reports whether the rootfs path is an absolute clean path.
func (r RootFS) ValidPath() bool {
	return isAbsoluteCleanPath(r.Path)
}

// isAbsoluteCleanPath reports whether one path is absolute and already clean.
//
// The sandbox rejects relative paths and paths holding a parent element,
// because both let a plugin leave the declared sandbox root.
func isAbsoluteCleanPath(target string) bool {
	if target == "" || !filepath.IsAbs(target) {
		return false
	}
	return filepath.Clean(target) == target
}

// containsPath reports whether one absolute path is inside the other one.
//
// The comparison is done on the relative path components, so a sibling
// directory sharing the same prefix is not treated as a child.
func containsPath(root, target string) bool {
	if root == "" || target == "" {
		return false
	}
	relative, err := filepath.Rel(root, target)
	if err != nil {
		return false
	}
	if relative == "." {
		return true
	}
	return relative != ".." &&
		!strings.HasPrefix(relative, ".."+string(filepath.Separator))
}
