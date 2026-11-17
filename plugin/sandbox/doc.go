/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package sandbox models the Abate plugin sandbox as a pure, deterministic
// configuration layer.
//
// 模型来源=gVisor runsc 配置语义（runsc/config 的 --network / --platform /
// --file-access 等配置项语义），Abate 自研纯函数实现，不调用 runsc 进程。
// The package never starts a process, never loads an external binary and never
// touches the network: it turns a plugin declaration into a runsc shaped
// configuration document, validates the configuration against the granted
// capability set and renders it as canonical JSON.
//
// 参考 gVisor runsc 的平台/网络/文件系统配置项，改进点=去进程化（纯数据模型）、
// 默认全拒的能力校验、确定性可比较的 JSON 序列化。
package sandbox
