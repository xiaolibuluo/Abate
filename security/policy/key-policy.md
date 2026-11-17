<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 密钥管理策略（XCU-SEC / T5）

> 对应 PLAN-08 XCU-SEC §4.3 与威胁 T5（签名密钥泄露）。本策略为流程控制，
> 由 `scripts/release.sh` 的 GPG 签名段执行，GA 前完成一次轮换演练。

## 1. 密钥分类

| 密钥 | 用途 | 保管 |
|---|---|---|
| 发布签名主密钥 | 对发布资产 + 清单 + tag 签名 | 离线硬件介质（HSM / 离线密钥卡） |
| 应急轮换密钥 | 主密钥失效时临时签名 | 独立离线介质，与主密钥物理隔离 |
| 信任锚公钥 | 用户验签（公示指纹） | 仓库 `SECURITY.md` + 发布公告 |

## 2. 离线保管

- 私钥**永不**进入 CI runner、镜像或任何联网主机；仅发布作业在隔离环境使用。
- 构建与签名分离：CI 产出资产，签名在离线/受控环境完成。
- 私钥口令长度 ≥ 30 字符，分段交由不同保管人（shamir 可选）。

## 3. 轮换流程

1. 新密钥在离线环境生成，信任锚公钥指纹公示于 `SECURITY.md` 与发布公告。
2. **重叠期 2 个版本**：新旧密钥并存验签，旧资产仍可被新旧密钥验证。
3. 重叠期结束后，旧密钥从信任锚列表移除并销毁私钥介质。
4. 轮换事件记入 `CHANGELOG.md` 与审计记录。

## 4. 泄露响应

- 主密钥疑似泄露 → 立即用应急轮换密钥重签在役资产并公告。
- 收回泄露密钥信任（从 `SECURITY.md` 信任锚移除）。
- 触发 GA 前外部安全审计复核。

## 5. 演练记录

- GA 前完成一次完整密钥生成 → 签名 → 验签 → 轮换重叠 → 销毁演练，记录入库。
- `scripts/release.sh --demo-sign` 提供一次性密钥的签名链演练（本地证据，不对外）。
