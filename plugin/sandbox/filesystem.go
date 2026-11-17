/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

// FilesystemPolicy is the filesystem allow list of the sandbox.
//
// 模型来源=gVisor runsc 的 bind mount 与 file-access 语义，Abate 以纯数据表达：
// 挂载点必须落在 rootfs 之内，写权限必须由 fs 能力显式授予。
type FilesystemPolicy struct {
	// Writable enables write access to the whole sandbox root.
	Writable bool `json:"writable"`

	// ReadOnlyPaths are the absolute paths mounted read only.
	ReadOnlyPaths []string `json:"readOnlyPaths"`

	// WritablePaths are the absolute paths mounted read write.
	WritablePaths []string `json:"writablePaths"`
}

// Mounts returns every declared mount in canonical order.
func (f FilesystemPolicy) Mounts() []string {
	out := make([]string, 0, len(f.ReadOnlyPaths)+len(f.WritablePaths))
	out = append(out, f.ReadOnlyPaths...)
	out = append(out, f.WritablePaths...)
	return out
}

// NeedsWrite reports whether the policy asks for any write access.
func (f FilesystemPolicy) NeedsWrite() bool {
	return f.Writable || len(f.WritablePaths) > 0
}
