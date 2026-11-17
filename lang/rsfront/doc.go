/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package rsfront implements the Abate Rust frontend using only the Go standard
// library, reducing Rust 2015-2024 sources to the frozen AIR tree.
//
// 参考 syn 的 item/expr/stmt/pat/ty 归约与 quote 的 token 树思想 + 改进点：
// 纯 Go 标准库自研宏调用、生命周期与模式匹配归约，剥离 proc-macro 支撑、clippy
// 逻辑、Rust/C ABI 与 cargo 调用，改为有序数据表驱动的单遍确定性词法与
// 递归下降/Pratt 解析，专有构造只用 lang/air 冻结枚举并降级标注，绝不扩展 AIR，
// 绝不启用子进程、CGO 或外部 API。
//
// Token tables, Pratt binding tables and construct-to-AIR lowering rules are
// ordered data tables, so identical source bytes always yield identical AIR.
package rsfront
