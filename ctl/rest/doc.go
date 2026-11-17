/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package rest implements the Abate local control API: experiment and check
// submission, report retrieval, plugin management and system probes.
//
// 参考 PLAN-07 CTL-REST 模块（17 端点 + OpenAPI 3.1 + 强制鉴权 + 限流 + 恢复）。
// 改进点=以 Go 标准库 net/http 自研轻量方法路由与中间件链，零第三方路由框架，
// 全链路离线、确定性（无 time.Now / 无全局随机 / 无 map 迭代序依赖）。
//
// 控制面为纯编排层：检查资源复用 lang/check + 各语言前端 + 内置规则；
// 实验资源以同一检查后端异步执行；插件资源经可注入 PluginManager 端口暴露，
// 默认提供静态内存实现，便于真实插件宿主后续插桩而不改动路由。
package rest
