# Abate 第三方资产 SSOT 台账（清除前冻结快照）

> 文档编号：ABATE-REFACTOR-INVENTORY-001　生成日期：2026-11-16
> 基线：`main` HEAD = `0b81c7b3`；工作树仅 `?? docs/benchmark/`（未跟踪残留）
> 生成：寇豆码（工程师）　依据：`docs/refactor/TODO.md` P0-01、`docs/refactor/DESIGN.md` §1.2 / §6 T01
>
> **性质**：本文件是**清除动作发生前**的单一事实来源。所有 sha256 均由本机实测命令产生，
> 未锁定来源 commit 的条目如实标注「未锁定」，不使用推测值。物理删除后本台账即不可复原的审计证据。
> ⚖️ 许可证与事实数据版权结论均标注「待权利人/法务确认」，本文不构成法律意见。

---

## 0. 采集方法（可复跑）

```bash
cd /workspace/abate
git rev-parse --short HEAD
du -sh third_party vendor docs/benchmark
# 目录聚合摘要：文件字节流按路径排序后逐文件 sha256，再对拼接结果取 sha256
find <dir> -type f -print0 | sort -z | xargs -0 sha256sum | sha256sum
```

实测体积：`third_party` = 528M；`vendor` = 13M；`docs/benchmark` = 105M。

---

## 1. 真正参与编译的第三方依赖（go.mod 权威）

| 组件 | 版本 | 引入路径 | 产品代码 import 命中 | 处置 | 责任包 |
|---|---|---|---|---|---|
| `github.com/tetratelabs/wazero` | v1.12.0 | 直接 require，由 `plugin/wasm/` 薄层引入 | 5 个文件（runtime/loader/hostfunc/instance/wat 薄层） | 废弃执行路径，收敛 `plugin/host` 原生宿主 | T03 |
| `gopkg.in/yaml.v3` | v3.0.1 | 直接 require，CES 解析 + TCK/生成器 | `ces/parse`、`ces/tck`、`cmd/gen-tck` 等 | 自研 `internal/yamlmini/` 严格子集替换 | T04 |
| `golang.org/x/sys` | v0.44.0 | indirect，由 wazero 引入 | 产品代码实测 **0 命中** | 不写替代代码，`go mod tidy` 随 wazero 清除 | T04 |

冻结文件摘要（清除前）：

| 文件 | 字节 sha256 |
|---|---|
| `go.mod` | `81c0c43a491035920ec304e1650ac3d7881b204214c7d405840af2d8b0c39fe0` |
| `go.sum` | `ff03af622cff01702d94e9697d34880461e7979067de39365d668b4c30713e8f` |
| `vendor/modules.txt` | `6ec47f5ff888e9095ff2ef270ac256ed7238ab282e3aa0862da84fb8e4c8d91f` |

`vendor/` 13M，含上述三份模块源码副本。**必须最后删**（当前构建依赖它；T04 代码清零后删）。

---

## 2. `third_party/_sources/`（纯参考源码副本，`_` 前缀不参与编译）

聚合摘要口径见 §0。全部条目处置 = **T02 物理全删**；均带 `_` 前缀，不进入 `go list ./...`。

| 目录 | 实测体积 | 聚合 sha256 | 上游来源（据 `LICENSE-MANIFEST.md` / `MANIFEST.md`） | 来源 commit | 处置 |
|---|---|---|---|---|---|
| `gvisor/` | 50M | `8db5787072bd4d7d03dafc53505113d95e6ed8824b33a70567df25273305eccc` | gvisor.dev/gvisor，Apache-2.0 | 未锁定（ADR-001 记「待 M4 锁定」） | 全删；ABI 事实数据另以自研 golden 冻结（§6） |
| `kotlin-compiler/` | 56M | `f8c54e6cf434fd662e6ed320d9e5181f0c8423bc4b47e671bed89ad90650e94a` | JetBrains kotlin-compiler | 未锁定 | 全删 |
| `llvm-project/` | 150M | `23119938ff2bdecac46944ef97c5489e11dc98ba8179eadb78be4ca1edd49db6` | LLVM project（含 LLVM exception） | tag/tarball 锁定，4 项 `RATE_LIMIT_PENDING` | 全删 |
| `pit/` | 2.1M | `26c8c22fd4de19da732196ee930921f34a0a90e2abcbc2e15e7f220fe42febdc` | PITest（PIT 变异核参考） | 未锁定 | 全删；语义以自研 golden 承接 |
| `roslyn/` | 24M | `3b01d5861918e7609c129fe4fa90deac1e2f7669f569f69a269698b61f104e64` | .NET Roslyn (codeanalysis) | 未锁定 | 全删 |
| `swc_ecma_parser/` | 916K | `d6c6e2020f228ac349ce3f67ce04c6047c9bb832ae4bc9eab38e8a56e58daf2a` | swc_ecma_parser | 未锁定 | 全删 |
| `syn/` | 2.4M | `fbe0398e1ba220551624915f04c2871a8dc6d14812a596e2cd70dfd89b17c087` | rust syn / quote | 未锁定 | 全删 |
| `tree-sitter-c/` | 8.7M | `1f4072930f3c7d81b9077ab4d5a93612f8674421f9da284b6dd33de84fcdce47` | tree-sitter-c | 未锁定 | 全删 |
| `tree-sitter-c-sharp/` | 65M | `59115af4485af8635bf98371c390cc56627fe492e545d3467673a17938328d14` | tree-sitter-c-sharp | 未锁定 | 全删 |
| `tree-sitter-javascript/` | 6.3M | `43f4483dabc2d9e3006f78f2de0f6e58420192f4eccd6eae62aef15e4c870a90` | tree-sitter-javascript | 未锁定 | 全删 |
| `tree-sitter-kotlin/` | 23M | `58d4d4ee9d6e3a7b88711e4bdbe62c4952cc48bbe4b6ebde2a65b550a155a463` | tree-sitter-kotlin | 未锁定 | 全删 |
| `tree-sitter-python/` | 7.3M | `5275e4eafdc09a97a46370d9342d87aad9a721d61884d8aae2d5d3ca8fb50a23` | tree-sitter-python | 未锁定 | 全删 |
| `tree-sitter-runtime/` | 872K | `78e708b4e80164cfda2aab3e450ed4850e3b9695ca3b3b6cea799cb70457eede` | tree-sitter runtime | 未锁定 | 全删 |
| `tree-sitter-rust/` | 15M | `3ca305461541b7c9e8833b2115041c6c015f5f70e562f0554e8501e87f581928` | tree-sitter-rust | 未锁定 | 全删 |
| `_validation/` | 39M | `d33ddfbaf07f1c8f6b6d31f8ace86613ca55dd8b7fc14dac4a9f67499d8b613d` | E2E 语料（11 个 pinned 上游，详见 `ci/corpus_manifest.json`） | tag + tarball sha256；4 项 `RATE_LIMIT_PENDING` | 全删 |
| `wazero/` | 81M | `8138071461abdd9da8803b371a00b5833013b618f9b2a1db96afeb21401fa2bc` | tetratelabs/wazero | v1.12.0 | 全删（T03 同步删执行层） |
| `README.md` | 12K | 单文件，随目录删除 | 目录说明 | — | 全删 |

---

## 3. `third_party/_references/`（纯参考副本）

聚合 sha256（整个目录）：`80643201d2e02eeb44b45763bdd59e640bb432d5faa011fb1e5e797e332cc044`

| 条目 | 实测体积 | 处置 | 备注 |
|---|---|---|---|
| `google__go-cmp/` | 24K | T02 全删 | — |
| `Shopify__toxiproxy/` | 12K | T02 全删 | — |
| `spf13__cobra/` | 60K | T02 全删 | — |
| `go-gremlins__gremlins/` | 8.0K | T02 全删 | — |
| `_m5_algorithms.go.txt` | 12K | T02 全删 | `m5-reference-scan.sh` ② 段硬依赖其存在，T05 改「零蓝本 = PASS」 |
| `sarif-2.1.0.schema.json` | 116K | T02 全删 | — |
| `MANIFEST.md` | 12K | T02 全删 | — |

---

## 4. `docs/benchmark/`（未跟踪残留，含开源源码）

| 目录 | 实测体积 | 聚合 sha256 | 处置 |
|---|---|---|---|
| `cases/js_lodash_4_17_21/src/` | 4.9M | `60f3384f07398da12944ccef71bac2cded97b770b00ab05e0257ef3117dccf0c` | T02 删源码 |
| `cases/js_lodash_4_17_21/config/` | 74M | `3a61261f5b4bbf94574e08d1cc97e021f1700fe4f4c15446e91cb81d856e5937` | T02 删 `node_modules` |
| `cases/c_cjson_v1_7_16/src/` | 1.8M | `5d07cd3c8514ce4518fe288ae275761a389f41ab985218b202d8603ac1de7af2` | T02 删源码 |
| `cases/c_cjson_v1_7_16/config/` | 22M | `b133574f6e7f4ce99137eaf9caaafec30059584da2b07c9a4d84c8526560b366` | T02 删依赖副本 |
| `raw/` | 424K | `381cd3c6d709b42c003d6a498d41d910b9a6ce19e476210f9bcff24b3f564891` | 保留测量数据；逐文件确认无源码片段，加「历史·未对拍」标注 |
| `audit/` | 12K | `894939b5da41f94e47b8a44fd0cb54905d404ee766b063303067e84d0e890458` | 保留，加标注 |
| `scripts/` | 68K | `8c819877f216e6e1a4a5c1ee6e10f4516642ff584db230f13648597febcdc5f8` | 保留脚本，`corpus_fetch.sh` 停用标注 |

---

## 5. `kernel/mutation/testdata/pit/`（PIT 官方产物 + Java baseline 源码）

| 文件 | 字节数 | sha256 | 处置 |
|---|---|---|---|
| `baseline/pom.xml` | 1988 | `e588385dd9e2c878ac3e847bd4c9bdfd68430316b7b48a6b67e7fa2b17689231` | T02/T05 删除（上游工程文件） |
| `baseline/src-main/java/com/abate/parity/ParityTarget.java` | 2116 | `61f4b41c3816365659e9314dd3712e6c1d636a53d64a7f1a97461110198f6cad` | 删除（上游 baseline 源码） |
| `baseline/src-test/java/com/abate/parity/ParityTargetTest.java` | 1553 | `4f72f7b8b6dda4e1fac896fac7885a1f5637e0a1738b952e4ecef76e9bda51f0` | 删除（上游 baseline 源码） |
| `pit-console-report.txt` | 369 | `192f335d5b7b6c3cfeebaea817fb7ae2b76e75cb530193fc921feb09e2a0c81d` | T05 改自研 golden 口径或删 |
| `pit-mutations.json` | 8690 | `dc5d5b1703dc5f4218c0bab55028d1e45d5fc7298dc36a7fc13e0d24e240b95a` | T05 改造为自研 frozen golden（28 条，非官方产物） |
| `pit-mutations.xml` | 18527 | `6122359af199039b601f87e9808f896be857ee7353feea5c70de09b64d4500f1` | T05 改自研 golden 或删 |

T01 已在本目录新建 `golden/`：值由本仓 `kernel/mutation/pitparity` 引擎对同构夹具的真实判定生成，
生成器 `kernel/mutation/pitparity/gen_golden.go`（`//go:build ignore`）可复跑，标注「非 PIT 官方产物」。

---

## 6. 融入产品代码的「二次并入」逐项甄别（DESIGN §4 映射）

| 位置 | 判定 | 处置 | ⚖️ 法务状态 |
|---|---|---|---|
| `kernel/mutation/` | 开源改写（已 Go 重写，无上游类型） | 代码保留，T05 改来源注释 + 独立 parity 测试 | L04：历史并入无遗留义务待确认 |
| `kernel/injection/` | 开源改写（已全虚拟化重写） | 代码保留，T05 改来源注释，跑 chaos/e2e | L04：待确认 |
| `kernel/abi/` | 开源改写（OS ABI 事实数据移植） | T01 测试脱钩自研 golden；T05 改注释 | **L05：syscall 表是否属不受版权保护的事实数据，待法务认定** |
| `lang/{py,js,rs,kt,cs,c}front/` | 算法参考、代码自研 | 代码保留，T05 改来源注释为「依据官方规范自研」 | 「参考算法」表述边界待确认 |
| `sim/rand/splitmix.go` | 公有领域算法 | 代码保留，T05 固化出处标注 | L05 类：公有领域认定待书面确认 |
| `sim/{clock,net,fs,vsched,record}/` | 自研 | 保留 | — |
| `plugin/{host,capability,signature,hotswap,sandbox,report,adapter}/` | 自研 | 保留 | — |
| `plugin/wasm/wat_*.go` | 自研 WAT 编解码（约 1400 行，无第三方 import） | T03 保留为离线工具或迁 `internal/wat/`，不得称插件运行时 | — |

---

## 7. 台账/声明类文件（条目改「已清除」，不写「已合规」）

| 文件 | sha256 | 处置 | 责任包 |
|---|---|---|---|
| `third_party/LICENSE-MANIFEST.md` | `f72fc405924e22352cc3bced630087a260164e17c85e9a83f6ef17e04375d1f7` | 保留文件，改「历史台账：全部组件已清除（日期+commit）」；`license-check.sh` 依赖其存在 | T02 |
| `ci/corpus_manifest.json` | `53ef136be8f8b21faf8cf376396c4782da0261b1b550f8aa586b0fd761289c84` | 条目标「已清除」 | T02 |
| `NOTICE` | `a64ee464f9491a3c4fceb0008268c3f458cabb3ff826ce1c36b55a73c3dfa350` | 同步改写 | T02/T05 |
| `THIRD-PARTY-NOTICES.md` | `dbc135463cc74b54c3e7bb8fa3d1284ff9feeeebbb53e486341a62d73e13f852` | 同步改写 | T02/T05 |
| `THIRD-PARTY-NOTICES.en.md` | `95eb7c6ed874638a254dda9348b7fe099e346ef6d8f50d4974fcc0426d6fbde6` | 同步改写 | T02/T05 |
| `LICENSE-COMPATIBILITY.md` | `232cdf7b97768f5aabad5f1d3d7e2624c9290ed75d891912b0e06813182c6343` | ⚖️ 争议点标「待权利人/法务确认」 | T05 |
| `LICENSE-COMPATIBILITY.en.md` | `a3e6713501893efcf47a3dfc82470974e1383d8d55d52dbf2518552ba134fa60` | 同上 | T05 |
| `scripts/corpus_fetch.sh` | 随脚本目录 | 停用 + 标注「已废弃」，保留审计痕迹 | T02 |

---

## 8. 清除后不可复原性声明

本台账覆盖 TODO §2.3（参考副本）、§2.4（benchmark 残留）、§2.5（二次并入）全部条目，
每项均有处置决策与责任任务包。物理删除后，任何「是否清干净」的争辩以本台账 +
`scripts/audit-zero-third-party.sh` 原始输出 + `docs/refactor/zero-third-party-audit.md` 为准，
不以声明为准。

**⚖️ 本台账不含任何法律结论**：AGPL §13 / DCO-CLA / SPDX / chaosblade-PIT 兼容性 / syscall 事实数据认定，
一律为「待权利人/法务确认」，见 TODO §8 L01–L05。
