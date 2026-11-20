<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 七语言覆盖量化（实测）

> 数据由 `scripts/language-coverage.sh` **现场从源码构建二进制后实测生成**，不是估算。
> 重跑：`bash scripts/language-coverage.sh`。任何数字以脚本输出为准。

## 量化表（实测）

| 语言 | 专属规则数（源码 `ID:` 声明） | 样例命中规则类 | 样例命中条数 | 规则族 |
|------|------------------------------|---------------|-------------|--------|
| Go | 116 | 2 | 2 | EG / VET |
| Python | 12 | 16 | 34 | AIR / PEP8 |
| JavaScript | 12 | 15 | 32 | AIR / ES |
| C | 14 | 11 | 35 | AIR / CERT |
| C# | 12 | 15 | 26 | AIR / DGN |
| Rust | 14 | 13 | 34 | AIR |
| Kotlin | 12 | 15 | 30 | AIR / DGN |

共享 AIR 结构模板数：**72**（`summary.rules`，各语言相同）。

> 说明：样例是每语言一份"含典型坏味道"的小文件（命名违规、大字面量、可疑结构等）。
> "命中规则类"是该样例实际触发的不同规则 ID 数；不同样例会得到不同数字。此表用于说明
> "每语言确实有可触发的规则"，不代表规则全集覆盖率。

## 如何读懂这组数据

**Go 的 116 条专属规则 vs 样例只命中 2 类**——不是"Go 弱"，而是两种不同深度：
Go 侧规则是 `rules_vet*` / `rules_eg*` / `rules_crc*` 系列，对标 go vet / Effective Go /
CodeReviewComments，是**针对具体 bug 模式的精准分析器**。本量化样例只是命名违规
（`foo_bar`、`BadName`），触发不了需要特定语义上下文的 vet 规则，只命中 `GO-EG-001` 与 `GO-VET-043`。
Go 走"**少而精、需特定上下文**"路线；其余六语言走"**共享结构模板、易广泛命中**"路线。

**其余六语言命中 11–16 类**——主要来自 72 条共享 AIR 结构模板（AIR-001/002/.../050 等），
叠加每语言 12–14 条专属规则。这些是**结构坏味道**（命名、字面量、可疑模式），命中多说明
共享模板覆盖广，但不等于"语义分析深"。

## 各语言规则族含义

| 规则族 | 语言 | 来源侧重 |
|--------|------|----------|
| GO-VET / GO-EG / GO-CRC | Go | go vet 分析器 / Effective Go / CodeReviewComments |
| PY-AIR / PY-PEP8 | Python | 共享 AIR 结构 / PEP 8 命名与风格 |
| JS-AIR / JS-ES | JavaScript | 共享 AIR 结构 / ES 惯用法 |
| C-AIR / C-CERT | C | 共享 AIR 结构 / CERT C 安全编码 |
| CS-AIR / CS-DGN | C# | 共享 AIR 结构 / .NET 设计指南 |
| RS-AIR | Rust | 共享 AIR 结构（专属规则偏命名/风格） |
| KT-AIR / KT-DGN | Kotlin | 共享 AIR 结构 / Kotlin 设计约定 |

## 覆盖范围诚实声明

1. 这是"能触发多少规则"，**不是"语言规范覆盖率"**：AIR 规则聚焦命名/风格/常见结构坏味道，
   不覆盖类型系统、并发安全、框架特有 API 误用。
2. Go 规则深度藏在上下文里：量化样例触发少不等于能力弱，需用专门样例才能触发 vet 系列。
3. 数字随样例而变：本表是固定样例的一次实测，重跑脚本可复现当前值。
4. 不做跨语言能力均等承诺：各语言规则数量与族不同是为贴合各自生态惯例。

## 全语料保真度（另一口径，勿与上表混用）

除上面的"规则触发"量化外，`scripts/fullcorpus-degrade.sh` 对全部 20520 条自研语料逐条真实 parse，
按 E2E 同公式测得全语料降级率：py 3.36% / js 3.51% / c 4.80% / cs 4.65% / rs 4.91% / kt 4.48%，
go 无嵌入语法映射为 `NOT_MEASURED`。六语均降到 5% 以下；该口径**未达**单文件冻结阈值，
只作事实披露，不构成"高保真/达标"结论。

## 复现

```bash
bash scripts/language-coverage.sh                      # 打印上表（现场构建 + 实测）
bash scripts/fullcorpus-degrade.sh                     # 全语料降级率
for d in go py js c cs rs kt common; do printf "%-8s " "$d"; \
  grep -rhoE 'ID:[[:space:]]*"' lang/rule/builtin/$d | wc -l; done   # 手工核对规则数
```
