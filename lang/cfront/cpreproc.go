/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"path"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// cPreproc 归约预处理行。keywords 决定哪些井号指令被识别为 KindDirective。
//
// 宏不展开：指令行只登记名字与语义，绝不改写后续 token。头文件只登记引用名，
// 内容由 air.SourceUnit.Files 以内存缓冲提供；本包从不打开真实文件系统。
type cPreproc struct {
	keywords []string
}

// newCPreproc 用一份独立副本构造预处理器，避免共享可变状态。
func newCPreproc(keywords []string) *cPreproc {
	return &cPreproc{keywords: append([]string(nil), keywords...)}
}

// isDirective 报告 name 是否是当前预处理表里的已识别指令。
func (p *cPreproc) isDirective(name string) bool {
	for _, keyword := range p.keywords {
		if keyword == name {
			return true
		}
	}
	return false
}

// lineEnd 返回当前预处理行的结束索引（换行或 EOF 或 hi，不含）。
func (*cPreproc) lineEnd(tokens []frontkit.Token, i, hi int) int {
	for j := i + 1; j < hi; j++ {
		if tokens[j].Kind == frontkit.TokenNewline || tokens[j].Kind == frontkit.TokenEOF {
			return j
		}
	}
	return hi
}

// parseLine 把一条井号指令行归约成一个节点并返回下一个待解析索引。
//
// 已识别指令映射到 KindDirective + SemDirective；#include 额外携带
// SemNote="header:<name>"，#define/#undef 携带 SemNote="macro:<name>"。
// 未识别指令降级为 KindIdent，绝不 panic。
func (p *cPreproc) parseLine(tokens []frontkit.Token, i, hi int) (air.Node, int) {
	end := p.lineEnd(tokens, i, hi)
	if end <= i {
		end = i + 1
	}
	if end > len(tokens) {
		end = len(tokens)
	}
	name := ""
	if i+1 < end {
		name = tokens[i+1].Text
	}
	kind := air.KindIdent
	if p.isDirective(name) {
		kind = air.KindDirective
	}
	start, last := cSpan(tokens, i, end)
	node := air.NewNodeAt(kind, start, last, name,
		cLeafNodes(tokens, i+1, end)...)
	if kind == air.KindDirective {
		_ = air.Annotate(node, air.SemDirective, name)
		_ = air.Annotate(node, air.SemTag, "preproc")
		switch name {
		case "include":
			header := p.headerName(tokens, i+2, end)
			if header != "" {
				_ = air.Annotate(node, air.SemNote, "header:"+header)
			}
		case "define", "undef":
			if i+2 < end {
				_ = air.Annotate(node, air.SemNote, "macro:"+tokens[i+2].Text)
			}
		default:
			_ = air.Annotate(node, air.SemNote, "directive:"+name)
		}
	} else {
		_ = air.Annotate(node, air.SemNote, "unknown-directive:"+name)
	}
	return node, end
}

// headerName 从内存 token 序列提取 #include 的头文件名，兼容 "x.h" 与 <x.h>。
func (*cPreproc) headerName(tokens []frontkit.Token, i, end int) string {
	if i >= end || i < 0 {
		return ""
	}
	first := tokens[i]
	if first.Kind == frontkit.TokenStr || first.Kind == frontkit.TokenChar {
		return strings.Trim(first.Text, "\"'")
	}
	if first.Text == "<" {
		var builder strings.Builder
		for j := i + 1; j < end; j++ {
			text := tokens[j].Text
			if text == ">" || text == ">>" {
				break
			}
			_, _ = builder.WriteString(text)
		}
		return builder.String()
	}
	return strings.Trim(first.Text, "\"'")
}

// ResolveInMemoryHeader 在 air.SourceUnit.Files 的内存缓冲里查找头文件。
//
// 这是「头文件改由 SourceUnit 内存提供」的落点：解析只依赖传入的缓冲，
// 不 import os、不读取磁盘、不调用任何外部 API。
func ResolveInMemoryHeader(name string, files []air.SourceFile) (air.SourceFile, bool) {
	if name == "" {
		return air.SourceFile{}, false
	}
	base := path.Base(name)
	for _, file := range files {
		if file.Path == name || path.Base(file.Path) == base {
			return file, true
		}
	}
	return air.SourceFile{}, false
}
