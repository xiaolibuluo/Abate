# 融入代码甄别台账（骨架 · T05 定稿）

> 文档编号：ABATE-REFACTOR-INTEGRATION-001　生成：2026-11-16（T01 骨架）
> 依据：`docs/refactor/DESIGN.md` §4 甄别矩阵；完整逐文件结论与扫描命令原始输出由 **T05** 补齐。
> 证据底稿：`docs/refactor/assets-inventory.md` §6。

---

## 1. 判定口径

- **开源改写** = 源码/结构曾参考上游，需重写或改写来源表述；
- **自研** = 纯自研实现；
- **公有领域算法** = 数学/事实数据，⚖️ 需法务确认保留边界。

AI 只给事实与处置建议，**不出法律结论**。

---

## 2. 逐项甄别状态（T01 快照）

| 模块 | 判定（DESIGN §4） | 当前状态 | 定稿动作（T05） |
|---|---|---|---|
| `kernel/mutation/` | 开源改写（已 Go 重写） | 代码无上游类型，来源注释仍在 | 改注释 + 独立 parity 测试 |
| `kernel/mutation/testdata/pit/` | 上游产物 + 源码副本 | T01 已生成自研 frozen golden；baseline 尚未删 | T02 删 baseline；T05 改脚本口径 |
| `kernel/injection/` | 开源改写（已全虚拟化） | 来源注释仍在 | 改注释 + 跑 chaos/e2e |
| `kernel/abi/` | 开源改写（事实数据移植） | T01 测试已脱钩自研 golden；注释仍在 | 改注释；⚖️ L05 待法务认定 |
| `lang/pyfront/` | 算法参考、代码自研 | 来源注释仍在 | 改注释为「依据 Python 官方规范自研」 |
| `lang/jsfront/` | 同上 | 同上 | 改注释为「依据 ECMAScript 规范自研」 |
| `lang/rsfront/` | 同上 | 同上 | 改注释为「依据 Rust Reference 自研」 |
| `lang/ktfront/` | 同上 | 同上 | 改注释为「依据 Kotlin 语法规范自研」 |
| `lang/csfront/` | 同上 | 同上 | 改注释为「依据 C# 语言规范自研」 |
| `lang/cfront/` | 同上（libclang 产物已出索引） | 同上 | 改注释；历史产物瘦身本次不做（R4） |
| `sim/rand/` | 公有领域算法 | SplitMix64 常量与出处注释在 | 固化出处标注；⚖️ 待书面确认 |
| `sim/{clock,net,fs,vsched,record}/` | 自研 | 保留 | 无需改 |
| `plugin/{host,capability,signature,hotswap,sandbox,report,adapter}/` | 自研 | 保留 | 无需改 |
| `plugin/wasm/wat_*.go` | 自研 WAT 编解码 | 保留（非插件运行时） | T03 决定原包保留或迁 `internal/wat/` 并记录 |
| `internal/standards/sarif/sarif-2.1.0.schema.json` | 公开标准事实数据（OASIS SARIF 2.1.0 官方 schema，非专有开源实现源码） | T02 从 `third_party/_references/` 原样迁入，sha256=`ad6db49878699b091f3eeb765b6e29e92a34bad4da88664d000c923b549c3a25`；`plugin/report/schema` 校验器 100% 自研 | 保留为事实数据；⚖️ 标准数据文件的分发/许可条款待权利人/法务确认，仅标注「待确认」，不出法律结论 |

---

## 3. 待补证据（T05）

1. 全仓 `find . -name '*.go'` 排除 `vendor/`、`third_party/` 的上游版权头扫描原始输出；
2. 逐文件判定覆盖表（路径 → 判定 → 处置 → 复跑命令）；
3. `header-check.sh` 全绿输出，作为「无上游版权头」的机器证据；
4. ⚖️ L01–L05 状态表（每项「待权利人/法务确认」+ 责任人，无 AI 法律结论）。

> 本文件在 T05 定稿前不作为完成证据。
