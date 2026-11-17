# Abate M4 插件生态 CI 证据（commit → publish ≤ 1h）

本文件是 M4 插件生态的验收证据索引，位于 `plugin/` 下，不修改 `docs/` 任何文件。
所有证据都可通过仓库内命令复现，不依赖外部服务、不联网、不引入第三方依赖。

## 1. 代理证据说明（诚实边界）

本沙箱没有真实的 GitHub Actions runner，因此「CI 证据」由两层组成：

- **声明层**：`.github/workflows/plugin-publish.yml` 是真实可执行的 GitHub Actions
  工作流（`push` 到 `main` 触发，先跑 9 段严格门禁，再 build + sign + 上传产物）。
- **执行层**：`scripts/plugin-publish.sh` 在本地完整模拟同一流水线的每一步
  （构建、竞态测试、CLI 测试、签名、产物加载/列举/调用验证、时间戳记录），
  并生成 `.artifacts/plugin-publish/PUBLISH-EVIDENCE.txt`。

工作流与脚本共用同一套步骤和同一个 3600 秒上界，所以本地执行层可覆盖
build / race 测试 / CLI 测试 / 签名 / 产物验证 / 耗时断言的全部行为。
**但有一处 CI 独有步骤无法在本沙箱实跑**：workflow 中
`Install golangci-lint v2` 是构建期联网拉取受版本固定（v2.1.6）工具二进制的
步骤，本沙箱既无 runner 也无该网络条件；该步骤仅做 YAML 解析与静态审阅
（`python3 -c "import yaml"` 可解析、局部 `env:` 覆盖已核对），**不宣称已在本地
真实执行**。gate / publish / upload 三个步骤仍保持 `GOPROXY=off` +
`GOSUMDB=off` + `GOFLAGS=-mod=vendor` 全离线，与本地脚本一致。
真实 runner 上的 `startedAt/finishedAt` 由 CI 每次运行重新写入证据文件。

## 2. commit → publish SLA

| 约束 | 位置 |
| --- | --- |
| 作业硬超时 60 分钟 | `.github/workflows/plugin-publish.yml` 的 `timeout-minutes: 60` |
| 门禁先行 | workflow 的 `Strict gate (9 segments)` 步骤调用 `bash scripts/lint-gate.sh` |
| 有界耗时断言 | `scripts/plugin-publish.sh` 计算 `elapsedSeconds` 并在超过 3600 秒时失败 |
| 时间戳证据 | 证据文件记录 `startedAt`、`finishedAt`、`elapsedSeconds`、`status=PASS` |
| 复现命令 | `bash scripts/plugin-publish.sh`（可重复、可传绝对路径输出目录） |

流水线步骤固定且全部本地执行：
`go build ./...` → `go test -race ./plugin/...` → `go test ./ctl/cli/...`
→ `go run ./internal/pluginpublish`（生成密钥、签名、自检 Verify、写 manifest 信封）
→ `go build -o abate .` → `abate plugin load/list/call` 验证签名产物
→ 记录耗时。私钥只在进程内存在，从不落盘；信封只携带公钥，CLI 加载时注册它并验证签名。
签名器每次运行 `GenerateKey()` 生成新公钥，因此 `manifestSha256` 逐次不同属预期语义
（每次发布独立密钥），不参与确定性回放比对。

CLI 的 `load/list/call/hotswap` 是独立进程，会话通过工作目录下固定的
`.abate/plugin-state.json` 状态文件显式重建：每条命令重新读取全部已签名信封并
重新执行 `Host.Load` 验签，因此跨进程可用且不会把未签名插件带进注册表。

### 2.1 M4 QA 第 1 轮 FAIL 的修复（会话状态重放 / 绝对路径 / 工具链自举）

QA 复验发现旧脚本二次运行必然失败（退出码 3）：发布脚本在产物目录运行 CLI，
其会话状态锚定在该目录的 `.abate/plugin-state.json`；二次运行时
`loadPluginSession` 重放上一次的旧信封，先注册旧公钥，而签名器本次生成的是
新公钥、锚点名固定 `plugin-release-anchor`，于是
`registerPluginAnchor` 检出「同锚点不同指纹」并拒绝加载。更严重的是失败发生在
写证据之前，旧的 `status=PASS` 证据会残留在目录里，被 workflow 整目录上传，
造成「失败却上传历史通过证据」的风险。

修复（均在 `scripts/plugin-publish.sh`，**未改动 `internal/pluginpublish` 的锚点名、
未持久化私钥、未破坏「每次发布生成独立密钥」语义**）：

1. 每次运行开始即 `rm -rf "$OUT_DIR/.abate"`，使「二次运行 == 首次运行」，
   从根上消除重放冲突；
2. 同时删除旧的 `PUBLISH-EVIDENCE.txt` 与 `load.json` / `list.json` / `call.json`，
   并把 3600 秒 SLA 断言前置于证据写入，保证任何中途失败都不留下陈旧
   `status=PASS` 证据，证据文件只在全流程成功后写；
3. `OUT_DIR` 先 `mkdir -p` 再 `cd` 归一为规范绝对路径，之后一律用
   `"$OUT_DIR/abate"` / `"$OUT_DIR/plugin-manifest.json"`，绝对路径入参不再被
   拼成 `/workspace/abate//tmp/.../abate`（旧缺陷退出码 127）；
4. 脚本内按 `go version` 探测工具链：只有确认为 `go1.26.11` 才保留调用方环境，
   否则显式切到受版本固定的 1.26.11 工具链并重置 `PATH`；因此 `env -u GOROOT`
   裸跑（旧脚本继承预置旧版 go 首步即报 `go.mod requires go >= 1.26.11`）也能全绿，
   调用方已正确 export 时则完全不改动其环境。

修复后的可复现命令与本地实测结果（离线 `-mod=vendor`，`GOPROXY=off`）：

| 场景 | 命令 | 退出码 | 证据 `status` |
| --- | --- | --- | --- |
| 默认目录首次运行 | `bash scripts/plugin-publish.sh` | 0 | PASS（`elapsedSeconds=4`） |
| 同目录二次运行 | `bash scripts/plugin-publish.sh` | 0 | PASS（`elapsedSeconds=3`） |
| 绝对路径全新目录 | `bash scripts/plugin-publish.sh /tmp/abate-pub-abs` | 0 | PASS（`elapsedSeconds=3`） |
| 裸跑自举 | `env -u GOROOT bash scripts/plugin-publish.sh` | 0 | PASS（`elapsedSeconds=4`） |
| 完全干净环境自举 | `env -u GOROOT -u GOFLAGS -u GOPROXY -u GOSUMDB PATH=/usr/local/go/bin:/usr/bin:/bin bash scripts/plugin-publish.sh` | 0 | PASS（`elapsedSeconds=3`） |

其余全仓验证：`go build ./...`、`go vet ./...`、
`go test -race -count=1 ./plugin/... ./ctl/cli/...` 均退出码 0；
报告确定性不变量
`go run . run ctl/cli/testdata/ok.yaml --seed 42 --format json | sha256sum`
仍为 `c00c280272c7611da4a79de07f7223f7ab76bfe6db5e97b0c44ea899ef9fea57`。
旧表述「为每次运行分配独立、可复现的会话状态」在旧脚本下并不成立（只 `mkdir -p`
不清理，二次运行必然重放冲突），现已由上表两种场景实测替代。

### 2.2 lint 工具安装：构建期局部联网，运行时离线

全局 `env:` 保持 `GOPROXY: "off"` / `GOSUMDB: "off"` / `GOFLAGS: "-mod=vendor"`。
旧 workflow 把 `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6`
放在该全局离线 env 下，实测 `cannot query module due to -mod=vendor` 退出码 1，
而 `scripts/lint-gate.sh:95` 依赖该二进制 → 声明层门禁在离线环境无法完成。

修复：把 lint 工具安装拆成**独立的联网步骤**，仅在该 step 内用 `env:` 局部覆盖
`GOPROXY=https://proxy.golang.org,direct`、`GOSUMDB=sum.golang.org`、`GOFLAGS=""`，
拉取版本固定的构建期工具二进制并把 `$GOPATH/bin` 加入 `PATH`；其余
gate / publish / upload 步骤继续继承全局离线 env。这是对**构建期工具链**的
局部联网，不是 Abate 运行时 API 接入：运行时零网络、零第三方 import、零 CGO、
零子进程；lint 二进制既不 vendor 进仓、也不进 `go.mod`，门禁不跳过 lint。
`timeout-minutes: 60` 的 SLA 硬边界与证据上传步骤均保留。

静态审阅结论：YAML 可被 `python3 -c "import yaml"` 解析；`timeout-minutes: 60`、
全局离线 env、安装步骤局部 `env:` 覆盖与版本固定、gate/publish 步骤无局部 `env:`、
`actions/upload-artifact@v4` 存在，全部核对通过。**本沙箱无法真跑 GitHub runner，
故此步骤只做静态正确性证明，不宣称已在真实 CI 执行。**

## 3. M4 四条验收标准的达成证据

### (1) 未签名插件被拒 / 能力越权被拦截

- `plugin/host/host_test.go::TestLoadUnsignedManifestRejected`：未签名 manifest 的
  `Host.Load` 返回 `ErrSignatureInvalid`，且不会进入 factory。
- `plugin/host/host_test.go::TestLoadUnknownCapabilityRejected`、
  `TestCallMissingCapabilityDenied`：未知能力声明 fail-closed，缺能力的调用返回
  `ErrCapabilityDenied`。
- `plugin/host/host_entries_test.go`：reporter/adapter 入口强制 `CapObserve` 门控，
  即使调用方传入空 want 列表也无法绕过。
- `ctl/cli/cmd_plugin_test.go::TestPluginRejectsUnsignedManifestWithoutRegistration`：
  CLI 端到端拒绝未签名 manifest（退出码 1）且 `plugin list` 确认没有任何注册。
- `plugin/sample/e2e_demo_test.go`：三插件全载入后再次验证签名拒绝与无副作用。

### (2) 热加载切换不中断实验

- `plugin/hotswap/` 实现 `PREPARE → SWITCH → DRAIN → UNLOAD` 四阶段状态机；
  `Allowed` 采用 major 必须相同、版本不得回退的 fail-closed 策略。
- `plugin/hotswap/hotswap_test.go::TestInFlightCallDrainsAndNewInstanceServes`：
  旧实例持有在途 `BeginCall` 时执行切换，验证 DRAIN 等待其完成、新实例在
  DRAIN 期间正常服务、旧实例拒绝新调用（`ErrInstanceStale`）。
- `plugin/hotswap/hotswap_chaos_test.go`：32 个独立 Host 并发切换、24 个并发在途调用、
  链式 8 次连续升级，全部在 `go test -race` 下通过。
- `ctl/cli/cmd_plugin_test.go::TestPluginHotswapEndToEnd`：CLI 载入 v1.0.0 →
  hotswap 到 v1.0.1 → 状态 `UNLOAD` → `list` 显示新版本 → 新版本继续 `call apply`；
  v1.0.0 回退被拒。

### (3) commit → publish ≤ 1h

见本文件第 2 节：workflow 硬超时、脚本耗时断言、时间戳证据文件，以及
`.github/workflows/plugin-publish.yml` 中注释写明的 SLA；第 2.1 节给出二次运行与
绝对路径两种场景的实测复现命令与退出码。

### (4) RPT-SARIF 通过官方 validator（OASIS schema）

- 官方 OASIS SARIF 2.1.0 schema 作为数据文件保存在
  `third_party/_references/sarif-2.1.0.schema.json`（不编译）。
- `plugin/report/schema/` 是自研、离线的 draft-07 子集 JSON-Schema 校验器，
  实现 `type`、`properties`、`required`、`enum`、`items`、`additionalProperties`、
  `$ref`、`oneOf`/`anyOf`/`allOf`、`pattern`、`minimum`/`maximum`、`format`，
  不引入任何第三方依赖。
- `plugin/report/report_test.go::TestBuildSARIFValidatesAgainstOASIS` 与
  `TestBuildSARIFRequiredFields`：生成的 SARIF 通过 OASIS schema 校验零错误，
  且具备 `version:"2.1.0"`、绝对 schema URI、至少一个 run。
- `plugin/report/report_mutation_test.go::TestOASISRejectsInvalidDocument` 与
  `plugin/report/schema/schema_mutation_test.go`：变异测试证明校验器不是空判定
  （缺 required、额外属性、非法枚举、越界数值都会被拒）。
- `ctl/cli/cmd_plugin_test.go::TestPluginReporterAndAdapterEntries` 与
  `plugin/sample/e2e_demo_test.go`：CLI / 端到端链路产出的 SARIF 同样通过校验。

## 4. 确定性与红线

插件运行时只使用注入的 `sim/clock.VirtualClock` 与 `sim/rand.Stream`；
`plugin/hotswap/timeout.go` 的 `DrainTimeout` 是唯一真实墙钟边界，仅用于
等待在途 goroutine 排空，不参与任何插件输出。注册表列举、findings 排序、
schema 校验遍历均显式排序，不泄漏 map 迭代序。发布脚本的 `date` 时间戳只写入
证据文件，不参与任何插件输出或哈希，因此 `c00c2802…9fea57` 报告不变量不受影响。
