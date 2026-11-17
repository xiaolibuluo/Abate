/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package csfront 实现 Abate 的 C# 语言前端，把内存中的源码归约成语言无关的 AIR 节点树。
//
// 参考 Roslyn SyntaxTree 的 trivia 与节点区间模型 + 改进点：纯 Go 标准库自研
// LINQ 查询子句、特性标记与可空引用类型的归约，剥离 IDE 工作区、分析器、代码修复、
// NativeAOT 编译以及 .NET 运行时进程调用；词法表、Pratt 绑定表与构造映射表均为有序
// 数据，不读取磁盘、不启动子进程、不使用 CGO，杜绝 map 迭代序、时钟与随机数参与归约，
// 保证相同源码字节恒产出相同的 AIR 与指纹。
package csfront
