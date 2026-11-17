# 自研语料全语料降级率测量（诚实口径）

> 文档编号：ABATE-FULLCORPUS-DEGRADE-001
> 测量日期：2026-11-18
> 可复现命令：`bash scripts/fullcorpus-degrade.sh`（或直接 `go run ./cmd/degradecheck`）
> 原始数据：`fullcorpus-degrade.json`，SHA-256 见 `fullcorpus-degrade.json.sha256`。

## 1. 为什么做这次测量

`README.md` / `README.en.md` 的诚实口径此前写明：E2E degrade「达标」**仅限冻结的单文件（首 unit）口径**，并注明"自研语料的全语料降级率尚未重测，为 `NOT_MEASURED`"。

E2E 门禁（`internal/e2e/runner_coverage.go:firstUnitPerLanguage`）确实用每语言**第一个真实 unit** 计算降级率。`ci/corpus_manifest.json` 登记的 20520 条自研语料虽然全部被 E2E 采样链路真实 parse，但降级率聚合只取了每语言首 unit。因此"全语料降级率"是一个真实存在、此前未测的口径缺口。本次测量把它闭环为可复现的真实数据。

## 2. 测量方法（与 E2E 同公式，全量聚合）

1. 读取 `ci/corpus_manifest.json` 全部 20520 条自研语料条目。
2. 对**每条**条目用其真实前端（`internal/e2e.FrontendTable`）parse，得到 AIR unit。
3. 用 `lang/frontkit/coverage.Observe` 统计真实 parse 产出的 AIR 节点 kind；`Merge` 按语言合并全语料观测。
4. 对每种有嵌入语法映射的语言，用 `internal/e2e/coverage.Degradation` 计算降级率：
   `degraded 且被观测到的映射数 / (exact + degraded 映射数) × 100`。
5. Go 未嵌入语法映射（`mappings=0`），诚实保留 `NOT_MEASURED`。

**反造假约束**：degraded 映射绝不重分类为 exact（`degrade_status_test.go` 有专门回归测试守护）；观测全部来自真实 parse；`parseFail` 与 `emptyUnit` 如实记录、不跳过；双跑逐字节 `cmp` 通过才归档；报告附 SHA-256。

## 3. 实测结果（原始数据见 JSON）

| 语言 | 状态 | 降级率 | 文件数 | parse 失败 | 空 unit | 观测 kind 数 | 映射数 | 被观测 degraded |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| go | NOT_MEASURED | — | 2932 | 0 | 0 | 38 | 0 | 0 |
| py | MEASURED | 5.04% | 2932 | 0 | 0 | 43 | 122 | 6 |
| js | MEASURED | 46.49% | 2932 | 0 | 0 | 32 | 114 | 53 |
| c | MEASURED | 10.40% | 2932 | 0 | 0 | 28 | 125 | 13 |
| cs | MEASURED | 5.12% | 2932 | 0 | 0 | 43 | 215 | 11 |
| rs | MEASURED | 6.13% | 2932 | 0 | 0 | 45 | 163 | 10 |
| kt | MEASURED | 5.22% | 2928 | 0 | 0 | 33 | 134 | 7 |

合计 20520 条，全部真实 parse 成功（parse 失败 0、空 unit 0）。

## 4. 诚实结论与边界

- 本表是**全语料口径**的真实测量，与 E2E 门禁的**单文件（首 unit）口径**不同，两者不得混用或互相外推。
- js 全语料降级率 46.49%，说明其降级映射在全语料覆盖下暴露更充分；这是客观现状，**不掩盖、不美化、不通过改标签降低**。
- 全语料口径数值**未达**单文件冻结阈值；本测量只作事实披露，**不构成**"高保真""达标""国际一线"等结论。
- Go 无嵌入语法映射，保持 `NOT_MEASURED`，不以 0 冒充达标。
- 历史第三方语料的六语数值（js 49.12% 等）见 `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`，与本表自研语料口径不同，禁止混用。

## 5. 复现

```bash
cd <repo>
bash scripts/fullcorpus-degrade.sh
# 产物：docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json
# 校验：sha256sum -c <(sed 's|$|  docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json|' \
#         docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json.sha256)
```

两份连续运行的报告逐字节一致，确定性成立。
