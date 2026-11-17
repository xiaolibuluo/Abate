/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package ktbuiltin ships the Kotlin official-style rule set layered on top of
// the shared AIR structural templates.
//
// 参考 Kotlin Coding Conventions 与 ktlint 标准规则集归档规范，
// 改进点=以共享 AIR 结构模式生成并叠加 Kotlin 专有规则（suspend、when、
// 空安全、data/sealed class、companion object、lateinit、inline 等）。
// 每条规则经 ktfront 真实前端端到端验证。
package ktbuiltin
