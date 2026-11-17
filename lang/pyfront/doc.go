/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package pyfront implements the Abate Python frontend using only the Go
// standard library, reducing Python 3.8-3.12 sources to the frozen AIR tree.
//
// 参考 ruff 的 rustpython-parser 与 tree-sitter-python 缩进/容错算法 + 改进点：
// 纯 Go 标准库自研缩进栈 tokenizer 与递归下降/Pratt 解析器，剔除 Rust/C ABI、
// CGO、子进程调用与双解析器降级，改为单遍容错恢复并以 SemPartial 标注降级节点。
//
// Token tables, Pratt binding tables and construct-to-AIR lowering rules are
// ordered data tables, so identical source bytes always yield identical AIR.
package pyfront
