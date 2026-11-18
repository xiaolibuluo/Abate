<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T01-build.md) | [English](T01-build.en.md)

# T01 · 从零离线安装与构建

## 1 目标读者

第一次接触 Abate 的新用户：目标是把二进制从源码构建出来，并确认构建溯源。

## 2 前置条件

- 已安装 **Go 1.26.11**（唯一权威构建版本，以 `go.mod` 的 `go` 指令为准；
  本教程所有版本输出均以本机 `go version` 实测为准）；
- 已获得仓库源码（项目**零外部依赖**：`go.sum` 为空、无 `vendor/`，无需联网下载模块）；
- 本教程全部命令在仓库根目录执行。

## 3 可复制命令

```bash
go version
cd /path/to/abate
export GOPROXY=off
go build -o /tmp/abate_verify .
/tmp/abate_verify version
```

## 4 预期输出

`go version`：

```
go version go1.26.11 linux/amd64
```

（`linux/amd64` 因平台而异；关键是 `go1.26.11`。）

`/tmp/abate_verify version`（6 行，逐字节稳定）：

```
abate 1.2.28
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
```

**验证点**：命令退出码为 `0`；版本输出恰好 6 行且包含
`license: AGPLv3 or commercial` 与 `watermark: 小黎的部落`。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| `go: downloading ...` 卡住或报网络错误 | 未设置 `GOPROXY=off`，Go 试图联网 | 执行 `export GOPROXY=off` 后重试（本仓库零外部依赖，无需联网） |
| `go: go.mod requires go >= 1.26.11` 之类版本错误 | 本机 Go 版本低于 1.26.11 | 升级到 Go 1.26.11 |
| 构建产物运行报 `permission denied` | 目标目录不可写或无执行位 | 换可写目录输出，或 `chmod +x` |
| `abate version` 输出行数不符 | 运行了旧二进制 | 删除旧产物后按第 3 节重新构建 |
