/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package frontkit defines the language-neutral contract layer for Abate's
// multi-language frontends.
//
// 参考 ruff/tree-sitter、swc、libclang、syn+quote、Roslyn、kotlin-compiler 的
// 词法与语法分析算法模型，改进点=以纯 Go 标准库重写为确定性 token 流驱动的
// 递归下降与 Pratt 解析骨架，剥离全部上游运行时、CGO、进程调用与第三方类型。
//
// It exposes only deterministic data contracts: tokens, lexical tables,
// canonical language tags and two pure interfaces. No implementation type from
// a later batch is referenced, so this package builds and is tested on its own.
package frontkit
