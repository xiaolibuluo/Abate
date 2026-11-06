// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package watermark 提供代码级数字水印与运行时开机亮牌。
// 本包为纯内存计算：不读写文件、不发起任何网络调用（禁止联网心跳校验，
// 由 no_network_test.go 以 AST 静态卡死）。
//
// Package watermark provides the code-level digital watermark and the
// runtime boot banner. This package is pure in-memory computation: no file
// I/O and no network calls of any kind (no phone-home heartbeat checks,
// enforced statically by no_network_test.go via AST).
package watermark

// 代码级数字水印常量：随二进制编译分发，任何副本均携带以下标识。
// Code-level watermark constants: compiled into the binary; every copy
// carries the following identification.
const (
	// Project 项目名。
	Project = "Abate"
	// Tagline 项目口号。
	Tagline = "止息混沌 · where chaos abates"
	// Spec 规范版本。
	Spec = "CES v0.1"
	// License 许可证。
	License = "AGPL-3.0-only + commercial"
	// Vendor 实现方。
	Vendor = "小黎的部落"
	// Website 实现方站点。
	Website = "www.qingdujian.com"
	// Contact 联系邮箱。
	Contact = "xiaolibuluo@163.com"
	// Copyright 版权声明。
	Copyright = "Copyright (C) 2026 小黎的部落（www.qingdujian.com）"
)

// Banner 返回运行时开机亮牌横幅：多行、以 \n 分隔、末尾不带换行，
// 由调用方决定输出流（serve 走 stdout，mcp 走 stderr）与换行方式。
//
// Banner returns the runtime boot banner: multi-line, \n separated, no
// trailing newline; the caller decides the output stream (stdout for serve,
// stderr for mcp) and how to terminate the line.
func Banner(version string) string {
	return Project + " —— " + Tagline + "\n" +
		Copyright + "\n" +
		"License: " + License + " · 商业用途须事先获得书面授权 Commercial use requires prior written authorization\n" +
		"Contact: " + Contact + " · " + Website + "\n" +
		"Version: " + version + " · Spec: " + Spec
}
