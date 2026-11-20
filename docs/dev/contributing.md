<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 贡献指南（Contributing）

## 1. 协议

- **DCO 强制**：每次提交须 `Signed-off-by`（`git commit -s`），声明你有权提交
  并同意 DCO 1.1。
- **CLA 可选**：签署后明确授予项目使用与再许可权利。
- 贡献即授予**双授权**权利（AGPLv3 + 商业），见 `governance/DUAL-LICENSE.md`。

## 2. 门禁（合入前必须全绿）

```bash
# 五段门禁
./scripts/lint-gate.sh
# 版权头门禁
./scripts/header-check.sh
# 许可证门禁
./scripts/license-check.sh
# 确定性门禁（核心包同种子双跑）
./scripts/determinism-check.sh
# 漂移检查门禁
./scripts/drift-check.sh
```

本地推荐一条龙：`go test -race -count=1 ./...` 后跑上述门禁脚本。

## 3. 红线（不可触碰）

- **零外部依赖**：产品代码只用 Go 标准库与自研实现（`go.mod` 无 `require`、无 `vendor/`、
  无 `third_party/`）；**禁止引入第三方模块**（API / 网络接入同样禁止，除工具链下载）。
- **确定性铁律**：禁止 `time.Now`（真实时钟）、全局 `math/rand`、map 迭代序依赖、
  `os.Environ` 输出进行为；内核 L0 对外部网络零依赖。
- **版权头**：逐字节复用（见 `governance/HEADERS.md`），新文件必带。
- **≤300 行/文件**；**禁止 TODO/FIXME**（godox 拦截）。
- **禁止遗留 scratch/debug 测试文件**。

## 4. 测试要求

- 表驱动，覆盖正常/边界/错误三路径。
- 新增/修改须附确定性回放测试（同种子双跑哈希一致）。
- 验证只跑自己包：`go test -race -count=1 ./<pkg>`（离线 `-mod=vendor`）。
- 覆盖率：内核 ≥95%/85%（语句/分支），生态层 ≥85%/75%。

## 5. 提交流程

本项目**不接受外部 Pull Request**；仓库变更由**权利人直接签名提交并推送**至 `main`
（规则集仅保留强制 GPG 签名 + 禁 force push，不含 PR 审批要求）。

1. 在本地跑全部门禁全绿（`bash scripts/lint-gate.sh`）。
2. `git commit -S -s`（GPG 签名 + DCO `Signed-off-by`）。
3. `git push origin main` 直接推送；CI 对推送做全量复核（夜间另跑模糊 + 对拍 + 规模档）。
