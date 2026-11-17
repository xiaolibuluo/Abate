/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package ktfront 实现 Abate 的 Kotlin 语言前端，把内存源码归约成语言无关的 AIR 节点树。
//
// 参考 kotlin-compiler 的 lexer/parser/PSI 前端模型 + 改进点：纯 Go 标准库自研
// 字符串模板 $x 与 ${} 插值、when 与空安全 ?. / !! / ?: 归约，剥离 JVM/JS/Native
// 后端、IDE 集成、Kotlin/Native 静态库与 kotlinc 进程调用；词法表、Pratt 绑定表与
// 构造映射表均为有序数据，不读取磁盘、不启动子进程、不使用 CGO，杜绝 map 迭代序、
// 时钟与随机数参与归约，保证相同源码字节恒产出相同的 AIR 与指纹。
package ktfront
