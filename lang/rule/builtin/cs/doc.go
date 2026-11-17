/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package csbuiltin ships the C#/.NET official-style rule set layered on top of
// the shared AIR structural templates.
//
// 参考 .NET Framework Design Guidelines 与 C# Coding Conventions 归档规范，
// 改进点=以共享 AIR 结构模式生成并叠加 C# 专有规则（LINQ、async 后缀、record、
// 可空引用、属性/索引器、宽异常捕获等）。每条规则经 csfront 真实前端端到端验证。
package csbuiltin
