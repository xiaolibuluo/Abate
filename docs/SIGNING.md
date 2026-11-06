<!--
Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
Copyright (C) 2026 小黎的部落（www.qingdujian.com）
本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
小黎的部落书面授权，详见 COMMERCIAL.md。
This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See COMMERCIAL.md for details.
-->

# 数字签名与验证指南 | Digital Signing & Verification Guide

> Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
> Copyright (C) 2026 小黎的部落（www.qingdujian.com）· AGPLv3 + 商业书面授权

本指南说明 Abate 发布产物的归属与完整性凭据，以及用户侧验证方法。

## 1. 凭据总览

| 层级 | 凭据 | 形式 | 所在位置 |
|---|---|---|---|
| 代码级 | 数字水印 | 编译进二进制的常量 | `watermark` 包 |
| 运行时 | 开机亮牌 | serve/mcp 启动横幅 | stdout / stderr |
| 发布物 | SHA-256 校验值 | `abate-<version>.sha256` | `dist/` |
| 发布物 | GPG 分离签名 | `abate-<version>.sig` | `dist/` |

## 2. 代码级数字水印

`watermark` 包将以下标识编译进二进制，任何副本均携带：

- 项目：Abate · 止息混沌 · where chaos abates
- 规范：CES v0.1
- 许可证：AGPL-3.0-only + commercial
- 实现方：小黎的部落（www.qingdujian.com）
- 联系邮箱：xiaolibuluo@163.com

由 `watermark_test.go` 锁死常量完整性；`no_network_test.go` 以 AST 静态检查保证水印计算零网络调用、零文件 I/O。

## 3. 运行时开机亮牌

- `abate serve`：横幅输出到 **stdout**，随后输出监听地址。
- `abate mcp`：横幅输出到 **stderr**——stdout 完整保留给 JSON-RPC 2.0 流，不污染协议（CES-8.4.3）。

横幅包含项目名、版权声明、许可证、联系邮箱、站点与版本号。

## 4. SHA-256 校验值验证

发布包附带 `abate-<version>.sha256`（`sha256sum` 格式）。验证：

```bash
sha256sum -c abate-<version>.sha256
```

期望输出：`abate: OK`。校验失败说明文件损坏或被篡改，请勿使用，并联系第 8 节邮箱。

## 5. GPG 签名验证

发布包附带 ASCII Armor 分离签名 `abate-<version>.sig`。验证：

```bash
gpg --verify abate-<version>.sig abate
```

首次验证需导入发布方公钥（获取渠道见第 8 节联系方式）。输出 `Good signature` 即通过。

## 6. 维护者发布流程

```bash
scripts/release.sh -v <version> [-o outdir] [-k key-id]
```

产出三件套：`dist/abate`（二进制）、`dist/abate-<version>.sha256`、`dist/abate-<version>.sig`。

退出码：0 成功 / 1 失败（缺 go、gpg 或签名私钥）/ 2 用法错误。gpg 或私钥缺失时脚本明确提示，不静默跳过签名。

## 7. 隐私承诺：禁止联网心跳校验

Abate **不做任何形式的联网心跳 / 电话回家（phone-home）校验**：

- 水印与亮牌为纯内存计算，不读写文件、不发起网络请求；
- `watermark` 包由 AST 测试静态卡死零网络调用；
- 许可证执法依赖法律途径（AGPLv3 传染条款 + 商业书面授权），不依赖技术遥测。

## 8. 联系方式

- 邮箱：xiaolibuluo@163.com
- 站点：www.qingdujian.com

商业授权、公钥获取与安全问题报告均通过上述邮箱联系。
