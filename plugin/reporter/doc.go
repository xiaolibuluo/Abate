/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package reporter wires the deterministic report builders into the host
// plugin contract as a single manifest-driven artifact reporter.
//
// 参考 plugin/sample/reporter.go 的插件包装，改进点=单类型按 manifest 格式分派
// 五个自研确定性 Build 函数，复用 findings + mutation.Report。
// The package is pure computation: it performs no I/O, reads no clock and
// never contacts an external service.
package reporter
