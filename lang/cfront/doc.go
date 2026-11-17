/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package cfront 实现 Abate 的 C/C++ 语言前端，把内存中的源码归约成语言无关的 AIR 节点树。
//
// 参考 libclang + 改进点：以 libclang 的 lexer/parser/AST 游标遍历模型为算法蓝本，改进点是纯 Go 标准库自研预处理指令归约与模板归约，剔除 LLVM 后端、CGO、真实文件系统与宏展开，头文件改由 air.SourceUnit 的内存文件表提供。整个包不 import os、不读取磁盘、不启动子进程，仅使用标准库与冻结的 air 枚举，解析结果逐字节确定.
package cfront
