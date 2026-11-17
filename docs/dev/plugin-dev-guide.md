<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 插件开发指南（Plugin Dev Guide）

> 对应 PLAN-02（WASM 运行时）/ PLAN-07（控制面）。插件是生态小时级扩展的载体：
> 从提交到发布全流程 ≤1 小时（见 `scripts/plugin-publish.sh` 演练）。

## 1. 插件类型

| 类型 | 作用 |
|---|---|
| `operator` | 扰动算子（注入/变异原语） |
| `reporter` | 报告器（报告格式/导出） |
| `adapter` | 适配器（CI / issue / 消息通道） |

## 2. 签名链（威胁 T1）

- 所有插件**必须签名**；未签名加载被拒（`plugin/signature.Verify`）。
- 信任锚公钥指纹公示于 `SECURITY.md`；签名密钥离线保管（见
  `security/policy/key-policy.md`）。
- 发布：`scripts/release.sh` 的 GPG 段对资产 + 清单签名（可 `--demo-sign` 演练）。

## 3. 能力沙箱（威胁 T2/T3）

- 默认**全拒**：host function 调用未经显式授予即拒绝（`plugin/capability`）。
- 运行时门控 + 审计；资源三重上限：燃料 + 内存 + 结果长度（`plugin/host`）。
- 越权调用被拦截（安全测试用例覆盖，T2/T3 待运行时宿主交付后补齐可执行用例）。

## 4. 开发流程

1. 用 WIT 定义接口（宿主层隔离 WASI 细节，WIT 版本化 + 兼容矩阵）。
2. 实现算子/reporter/adapter，编译为 WASI Preview2 组件。
3. 本地签名：`scripts/plugin-publish.sh`（build → race → CLI 测试 → 签名 manifest）。
4. `abate plugin load --manifest <m> --type <operator|reporter|adapter>` 校验加载。
5. `abate plugin call --id <id> --entry <entry> --args <base64>` 调用。
6. 热换：`abate plugin` 热加载新版本，旧实例排空后卸载（不中断在跑实验）。

## 5. 安全自检

- 未签名插件加载必须被拒（T1）。
- 能力越权必须被沙箱拦截（T2）。
- 资源耗尽必须被燃料/内存上限拦截（T3）。
- 签名密钥泄露由离线保管 + 轮换流程缓解（T5）。

## 6. 发布节奏

- 插件仓库独立 semver + 兼容性矩阵（宿主版本区间）。
- CI 按小时发布窗口（生态小时级）；与内核年轮版本解耦。
