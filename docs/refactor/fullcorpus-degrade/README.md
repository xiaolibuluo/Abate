<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 自研语料全语料降级率（实测数据）

> 可复现命令：`bash scripts/fullcorpus-degrade.sh`（或 `go run ./cmd/degradecheck`）。
> 原始数据：`fullcorpus-degrade.json`，SHA-256 见 `fullcorpus-degrade.json.sha256`。
> 数据由脚本现场逐条真实 parse 产出；重跑即覆盖，以下数值以脚本输出为准。

## 测量方法

1. 读取 `ci/corpus_manifest.json` 全部 20520 条自研语料。
2. 对每条用其真实前端（`internal/e2e.FrontendTable`）parse 得到 AIR unit。
3. 用 `lang/frontkit/coverage.Observe` 统计真实 parse 产出的 AIR 节点 kind，`Merge` 按语言合并。
4. 对有嵌入语法映射的语言，用 `internal/e2e/coverage.Degradation` 计算降级率：
   `degraded 且被观测到的映射数 / (exact + degraded 映射数) × 100`。
5. Go 无嵌入语法映射（`mappings=0`），诚实保留 `NOT_MEASURED`。

反造假约束：degraded 映射绝不重分类为 exact（有专门回归测试守护）；观测全部来自真实 parse；
parse 失败与空 unit 如实记录；双跑逐字节 `cmp` 通过才归档；报告附 SHA-256。

## 当前实测结果

| 语言 | 状态 | 降级率 | 文件数 | parse 失败 | 空 unit | 观测 kind 数 | 映射数 | 被观测 degraded |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| go | NOT_MEASURED | — | 2932 | 0 | 0 | 38 | 0 | 0 |
| py | MEASURED | 3.36% | 2932 | 0 | 0 | 43 | 122 | 4 |
| js | MEASURED | 3.51% | 2932 | 0 | 0 | 58 | 114 | 4 |
| c | MEASURED | 4.80% | 2932 | 0 | 0 | 34 | 125 | 6 |
| cs | MEASURED | 4.65% | 2932 | 0 | 0 | 43 | 215 | 10 |
| rs | MEASURED | 4.91% | 2932 | 0 | 0 | 47 | 163 | 8 |
| kt | MEASURED | 4.48% | 2928 | 0 | 0 | 34 | 134 | 6 |

合计 20520 条，全部真实 parse 成功。六种有嵌入语法映射的语言均低于 5%。

## 诚实边界

- 本表是**全语料口径**的真实测量，与 E2E 门禁的**单文件（首 unit）口径**不同，不得混用或互相外推。
- 全语料口径数值**未达**单文件冻结阈值；本测量只作事实披露，**不构成**"高保真""达标"结论。
- Go 无嵌入语法映射，保持 `NOT_MEASURED`，不以 0 冒充达标。
- 降级率下降均来自前端在真实语法归约点产出专用节点（有最小样例 + 真实语料双重守护测试），
  而非改标签凑数；映射到 AIR 尚无独立节点的构造如实保留 degraded，并由 pin 测试锁定。

## 复现

```bash
bash scripts/fullcorpus-degrade.sh
sha256sum -c <(sed 's|$|  docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json|' \
        docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json.sha256)
```
