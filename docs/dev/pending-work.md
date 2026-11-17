# Abate 项目 · 完工状态归档（原「全量待完成工作清单」）

> 核实时间：2026-11-14（GMT+8）
> 来源：2026-11-11 18:20 的待办快照（当时标注「工程实现约 87%」）。
> 本文按源码与工件实况逐项复核后重写；每项给出结论、证据路径与遗留边界。
> **禁止伪造**：未完成项如实保留，不以「已完成」措辞掩盖。

---

## 一、总览

| # | 项目 | 优先级 | 结论 | 关键证据 |
|---|------|--------|------|---------|
| 1 | libclang-min 真实构建 | P0 | ⚠️ 产物级达成、功能级未闭环 | `.artifacts/libclang-min/build/lib/libclang-18.so.18` |
| 2 | WASM 宿主（`plugin/wasm`） | P0 | ✅ 已完成 | 32 个 .go + wazero v1.12.0 vendoring |
| 3 | gVisor `pkg/abi` 裁剪移植 | P1 | ✅ 已完成 | `kernel/abi/` + `docs/dev/abi-port-parity.md` |
| 4 | `jsfront/jsmap.go` 拆分 | P1 | ✅ 已完成 | 3 个文件均 ≤300 行 |
| 5 | 终版九段门禁 + `-race` + govulncheck | P0 | ✅ 已完成 | `CI GATE: PASSED`，辅助门禁全绿 |
| 6 | 文档统一收口 | P2 | ✅ 本次完成 | 本文 + `CHANGELOG.md` |
| 7 | SBOM + Release 流水线实跑 | P2 | ⬜ 未执行（非门禁项） | 脚本就绪，无本轮工件 |
| 8 | 完工度终审报告落盘 | P0 | ⬜ 未落盘（判定见 §五） | 独立报告文件缺失 |

**结论**：工程实现与冻结门禁口径（含降级率单文件口径）已全部达成（1–6 项，其中第 1 项的功能级验证为已知边界）；
铁律口径（功能级验证）与全语料口径另行说明，见本文 §五及 `docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度——
降级率达标仅限冻结单文件口径，全语料同公式六语全部超阈；7、8 属流程性收尾，未进入任何 CI 门禁，如实列为剩余项。

---

## 二、逐项明细

### 1. libclang-min 真实构建（P0）— ⚠️ 产物级达成、功能级未闭环

**已达成**
- 从 LLVM 18.1.3 源码真实编译出 `libclang-18.so.18`（79,650,944 字节，2026-11-13 18:55），
  `ldd` 全部依赖解析到系统 `libstdc++ / libm / libgcc_s / libc`。
- 同目录产出 `libclang*.a` / `libLLVM*.a` 静态库（含 `libclangAST.a` 22MB、
  `libclangSema.a` 30MB 等）。
- 构建配置可追溯：`.artifacts/libclang-min/build/CMakeCache.txt` 中
  `CMAKE_HOME_DIRECTORY=/tmp/llvm-src/llvm-toolchain-18-18.1.3/llvm`。

**未闭环（如实标注）**
- `build-report.txt` 与 `build.log` 仍是 2026-11-11 的失败快照（第 9 轮
  `[563/2380]` 编译器 OOM），未随 11-13 的成功构建更新。
- 原验收标准中的 `libclang-min.a` 静态库产物、C API 功能级验证
  （最小 C 程序经 `clang_getCXCursor` 类调用）**未落盘**；源码中无 CGO 桥接，
  `lang/rule/builtin/c/doc.go` 明确 AIR 不暴露 clang/libclang 类型。
- 构建产物与平台绑定，本次已从版本控制移除并纳入 `.gitignore`（见 §四）。

**判定**：产物真实存在且可解析，但「功能级验证」铁律口径未满足，不计入 100% 铁律完成。

### 2. WASM 宿主实现（P0）— ✅ 已完成

- `plugin/wasm/` 共 **32 个 .go**（WAT 词法 / 解析 / 编码、模块加载、fuel/gas 限额、
  确定性时钟、宿主函数绑定、运行时），`go test ./...` 全绿。
- wazero v1.12.0 已完整 vendoring：`go.mod`、`vendor/modules.txt`、
  `third_party/_sources/wazero/`，经 `go mod verify` 与 `license-check` 非 GPL 判定。
- 原清单「0 个 .go 文件 / wazero 尚未 vendoring」为 11-11 旧状态，已失效。
- 依赖面门禁 `scripts/m5-reference-scan.sh` ③ 段已由「仅 yaml.v3」过时快照改为
  以 `go.mod` 为权威的双向一致性校验（提交 `cbedd1fc`），消除对 wazero / x/sys 的误报。

### 3. gVisor `pkg/abi` 裁剪移植（P1）— ✅ 已完成

- `kernel/abi/`：syscall 号→名表（amd64 / arm64）、errno、`Addr` / `AddrRange` /
  `PageSize`、架构常量与 `syscall_test.go`。按 Abate 风格重写，不 import 上游。
- `docs/dev/abi-port-parity.md` 记录逐项对拍：amd64 **352 条**、arm64 **295 条**、
  errno 131 + 2 + 4，与 `third_party/_sources/gvisor/` 上游逐条一致；测试在
  `go test` 时直接读取上游 `linux64.go`，上游变更会立即暴露偏差。

### 4. `jsfront/jsmap.go` 拆分（P1）— ✅ 已完成

- 原 384 行单文件已拆为 `jsmap.go` (106) / `jsmap_children.go` (118) /
  `jsmap_node.go` (193)，均 ≤300 行；`style-check.sh`、`header-check.sh` 通过。

### 5. 终版门禁全量实跑（P0）— ✅ 已完成

- `scripts/lint-gate.sh`：`CI GATE: PASSED`（header / style / license / gofmt /
  vet / golangci-lint **0 issues** / constant-lint / `go test -race` / coverage / build）。
- 辅助门禁：`determinism-check` PASSED、`drift-check` PASSED、
  `pit-parity` **100%**、`m5-reference-scan` PASSED、`vuln-scan` **PASS**、
  `go mod verify` all modules verified。
- TLA+ `scripts/tla/check.sh`：PitParity 44 / SandboxFuel 61 /
  StreamDeterminism 4800 状态，3 spec 全通过。
- **GO-2026-6859**（`net/http` MaxBytesReader header race，影响 go1.26.10、
  修复于 1.26.11）：工具链统一升级至 **Go 1.26.11**，`govulncheck` 真实结果
  `No vulnerabilities found`（提交 `cbedd1fc`）。原清单「govulncheck 待终跑」已闭环。
- 覆盖率维持既有水位（mutation 98.3% / pitparity 95.5% / sandbox 96.2%）。

### 6. 文档统一收口（P2）— ✅ 本次完成

- 本文按实况重写，消除「WASM 未开工」「wazero 未 vendoring」等陈旧陈述。
- `CHANGELOG.md` `[v1.0.0]` 段修正「gVisor 平台级沙箱与真 WASM 运行时仍为未交付边界」
  这一与代码实况矛盾的表述，并补记本轮增量收官（版本号不变，漂移门禁不受影响）。
- `docs/dev/frontend-coverage/` 六语言覆盖清单与 `docs/dev/pit-parity-report.md`
  均为已产出状态，交叉引用保持不变。

### 7. SBOM + Release 流水线实跑（P2）— ⬜ 未执行

- `scripts/sbom.sh`、`scripts/release.sh`、`scripts/plugin-publish.sh` 均就绪；
  `.artifacts/` 下**无**本轮 SBOM / release 工件，`plugin-publish/` 等目录为
  早前演练留存的验收证据。
- 该项不进入九段门禁，属可选发布流程收尾；如实保留「待执行」，不标完成。

### 8. 完工度终审报告落盘（P0）— ⬜ 未落盘

- 独立终审报告文件未生成；本文 §五 给出即时双口径判定作为替代结论。
- 原清单「终审通过后方可宣称 100% 完工」的流程条件未以独立文档形式闭环。

---

## 三、已闭环链路（原「依赖与顺序」）

```
[2 WASM 宿主] ──┐
[3 gVisor 移植] ├─→ [5 终版门禁] ─→ [6 文档收口] ─→ 工程 / 门禁口径 100%
[4 jsmap 拆分] ─┤        ↑
[1 libclang 产物]┘   [TLA+ / 5 项辅助门禁]
        │
        └── ⚠️ 功能级验证未闭环（铁律口径遗留边界）

剩余（非门禁项）：[7 SBOM / Release 实跑] · [8 独立终审报告落盘]
```

---

## 四、已完成基线（供终审对照）

| 事项 | 状态 |
|------|------|
| 六类上游源码 vendoring（gVisor / PIT / LLVM / Roslyn / Kotlin / tree-sitter 系 + syn / swc_ecma_parser） | ✅ `third_party/_sources/` |
| tree-sitter grammar 数据 + 运行时功能级验证（gcc 编译解析 Python） | ✅ sha256 与上游一致 |
| PIT 1.19.6 官方对拍 | ✅ 28 变异 → 可比对 21 → 一致 21 → **100%**（`.artifacts/pit-parity.json`） |
| `kernel/mutation/pitparity`（15 .go）+ `plugin/sandbox`（14 .go） | ✅ 已交付并复验 |
| 六语言前端覆盖清单（`docs/dev/frontend-coverage/`） | ✅ 已产出 |
| `lang/pyfront` pyAsPower 前置符号 | ✅ 已修复（`pytokens.go:26`） |
| 代码库 TODO/FIXME 残留 | ✅ 0 处 |
| GitHub 全量上传（xiaolibuluo/Abate） | ✅ 11-11 快照已上传；本地后续提交未推送（无 remote） |
| **wazero v1.12.0 vendoring + `plugin/wasm/` 真 WASM 运行时（32 .go）** | ✅ 已交付 |
| **`kernel/abi/` gVisor 裁剪移植 + 上游逐项对拍** | ✅ amd64 352 / arm64 295 / errno 137 |
| **`lang/jsfront/jsmap.go` 拆分（≤300 行）** | ✅ 3 文件全达标 |
| **清零全部 lint 命中（7 高复杂度 + 8 连带 + 6 超限文件拆分）** | ✅ `golangci-lint` 0 issues（`b69f2c1c`） |
| **九段门禁 + TLA+ 3 spec + 5 项辅助门禁全绿** | ✅ `CI GATE: PASSED` |
| **Go 1.26.11 升级修复 GO-2026-6859** | ✅ `vuln-scan` PASS（`cbedd1fc`） |
| **构建产物清理**：1013 个 `libclang-min/build/` 文件移出版本控制并纳入 `.gitignore` | ✅ 本次（工作区文件保留，不影响本地构建） |

---

## 五、完工度判定（双口径，如实标注）

| 口径 | 结论 |
|------|------|
| **工程实现 + 门禁验证** | **100%**：1–6 项闭环，九段 CI 门禁、TLA+ 3 spec、determinism / drift / pit-parity(100%) / m5 / vuln-scan / `go mod verify` 全绿，`golangci-lint` 0 issues，59 包 `go test -race` 0 FAIL |
| **铁律口径（真实 vendoring + 功能级验证）** | **未宣称 100%**：`libclang` C API 功能级验证与 `build-report.txt` 更新未闭环；六类上游 vendoring + wazero + gVisor abi 的功能级验证均已闭环 |

**遗留边界（禁止伪造）**
1. `libclang-18.so.18` 已真实编译，但 CXCursor 遍历 / CGO 桥接 / 报告更新未落盘。
2. `SBOM` + `Release` 流水线本轮未实跑（脚本就绪，无工件）。
3. 独立终审报告未落盘（本文 §五 为即时判定）。
4. `.artifacts/libclang-min/build/` 已从索引移除，但历史 commit 中仍含这些产物；
   如需彻底瘦身历史，需独立的 `git filter-repo` 操作（改写全部历史 hash），本次未执行。

**若要补齐铁律口径 100%**，建议顺序：
① libclang C API 最小功能级验证 + 重写 `build-report.txt`；
② `sbom.sh` / `release.sh` 实跑并留存工件；
③ 落盘独立终审报告。
