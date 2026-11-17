/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package jsfront 实现 LNG-JS 的 JavaScript/TypeScript 前端，把源码确定性归约为 AIR.
//
// 参考 swc_ecma_parser 的词法与 AST 归约模型 + 改进点.
//
// 改进点=以纯 Go 标准库自研自动分号插入（ASI）、正则与除法歧义判定、模板插值
// 二次扫描以及 TypeScript 类型保留解析器，类型注解、泛型、接口与枚举全量保留而
// 不擦除；剥离 Rust/C ABI、CGO 与 transform/lint 路径，不依赖任何第三方运行时、
// 子进程或外部 API.
//
// 六种扩展名 .js/.jsx/.mjs/.cjs/.ts/.tsx 在解析入口协商方言：JSX 归约、
// TypeScript 类型归约与模块制式分别受其控制；专有构造一律降级到 lang/air 中
// 真实存在的 NodeKind 与 SemKey.
//
// AIR itself is never modified.
package jsfront
