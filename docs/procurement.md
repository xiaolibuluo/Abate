# Abate 采购清单（Procurement Manifest）

> 状态：采购阶段（禁止编码，待全部齐备后进入 M1 实现）。
> 红线：工具版本对齐项目（go 1.26.11 / golangci-lint v2），**禁止降级**；
> 接入开源必须 **Go 语言 + 源码接入 + 非 GPL**；一律走镜像下载。

## 1. 主工具链（对齐项目版本）

| 组件 | 项目声明版本 | 来源（镜像） | 状态 |
|---|---|---|---|
| Go toolchain | go 1.26.11 | 官方镜像 `golang.google.cn/dl` 下载 + SHA256 校验，`GOTOOLCHAIN=local` 固定 | ✅ 已安装（go1.26.11 linux/amd64，校验通过） |
| golangci-lint | v2（CI 使用 `/v2/cmd/golangci-lint`） | `go install .../golangci-lint/v2/cmd/golangci-lint@v2.1.6`（镜像） | ✅ 已安装 v2.1.6（既有二进制构建自 go1.26.10，兼容 1.26.11 工具链） |

## 2. 辅助工具链（均为 Go、走镜像、非 GPL）

| 组件 | 用途 | 安装命令 | 状态 |
|---|---|---|---|
| gofumpt | 比 gofmt 更严格的格式化（门禁可选增强） | `go install mvdan.cc/gofumpt@latest` | ✅ 已安装（v0.12.0，镜像） |
| goheader | 版权头门禁（XCU-GOV） | `go install github.com/denis-tingajkin/goheader@latest` | ⚠️ 镜像 404 不可用 → **以 `scripts/header-check.sh` 脚本等价替代**（已接入门禁） |
| go-licenses | 依赖许可证分类（LICENSE 门禁增强，GPL 检测） | `go install github.com/google/go-licenses@latest` | ✅ 已安装（v1.6.0，镜像） |

## 3. 需源码接入的开源（Go / 非 GPL / 镜像）

| 组件 | 许可证 | 语言 | 里程碑 | 接入方式 | 状态 |
|---|---|---|---|---|---|
| yaml.v3 | MIT | Go | M1（CES YAML 解析） | Go module 源码可读（module cache） | ✅ 源码已入 module cache（v3.0.1） |
| gVisor（隔离内核子集） | Apache-2.0 | Go | M4（WAS-CAP 沙箱，用户指定） | `go mod download gvisor.dev/gvisor` 取源码，M4 落 `third_party/gvisor` 五步并入 | ✅ 源码已入 module cache（latest；注意 gVisor@latest 要求 Go≥1.27.3 方可构建，M4 将锁定兼容版本或升 Go） |
| chaosblade（注入核） | Apache-2.0 | Go | M2 | M2 落 `third_party/chaosblade` | 📅 M2 |
| SplitMix64 | Public Domain | Go（自实现） | M1（SIM-RNG） | 算法常量，sim/rand 自实现，零外部依赖 | 📅 M1 实现期 |

## 采购完成记录（2026-11-08）

- 全部门禁脚本已就绪并 `chmod +x`：`header-check.sh` / `style-check.sh` / `license-check.sh` / `coverage-gate.sh` / `lint-gate.sh`。
- `.golangci.yml` 已重写为 **v2 语法**并实测通过（0 issues）；`Makefile` / `.github/workflows/ci.yml` 已对齐 Go 1.26 + golangci-lint v2.1.6。
- 七语言官方 + 社区规范规则已归档 `docs/reference/`（go-effective/go-uber/python-pep8/python-google/js-airbnb/js-ts-official/c-cert/cpp-core/rust-api/csharp-design/kotlin-conventions/kotlin-ktlint）。
- **全部门禁实测通过**（header/style/license/gofmt/vet/golangci-lint/test-race/coverage/build 全绿）。
- ⚠️ 待用户决策：见 §5 非 Go 底层库与「必须 Go」红线冲突（仅影响 M3/M5，不阻塞 M1）。

## 4. 七语言官方规范 + 社区规范规则（参考规则库，本阶段下载归档）

> 这些规则是「自研规则集」的权威来源（PLAN-05 §3.3）。本阶段把官方/社区规范文本
> 抓取到 `docs/reference/`，供 M3/M5 规则引擎落地时逐项映射到 AIR 模式。

| 语言 | 官方规范 | 社区规范 | 归档路径 |
|---|---|---|---|
| Go | Effective Go + CodeReviewComments + go vet | Uber Go Style Guide | docs/reference/go-* |
| Python | PEP 8 / PEP 257 / PEP 484 | Google Python Style Guide | docs/reference/python-* |
| JS/TS | ECMAScript 规范 + TS Handbook/Do's & Don'ts | Airbnb JavaScript Style / StandardJS | docs/reference/js-* |
| C/C++ | CERT C/C++ Secure Coding | C++ Core Guidelines (isocpp) | docs/reference/c-* |
| Rust | Rust API Guidelines + Rust Reference | clippy lints | docs/reference/rust-* |
| C# | .NET Framework Design Guidelines + C# Coding Conventions | IDesign / StyleCop | docs/reference/csharp-* |
| Kotlin | Kotlin Coding Conventions | ktlint / Android Kotlin Guide | docs/reference/kotlin-* |

## 5. ⚠️ 必须决策：七语言前端「底层库」与「必须 Go」红线的冲突

PLAN-05 §3.1 规划的七语言解析底层库里，**有 5 个不是 Go 实现**：

| 语言 | 架构规划的底层库 | 语言 | 与「必须 Go」红线 |
|---|---|---|---|
| Python | ruff（rustpython-parser 系）/ tree-sitter-python | **Rust / C** | ❌ 冲突 |
| JS/TS | swc | **Rust** | ❌ 冲突 |
| C/C++ | libclang（LLVM） | **C++** | ❌ 冲突 |
| Rust | syn + quote | Rust（语言自身，无法替代） | ⚠️ 必然非 Go |
| C# | Roslyn | **C#** | ❌ 冲突 |
| Kotlin | kotlin-compiler | **Kotlin/JVM** | ❌ 冲突 |
| Go | go/ast（标准库） | Go | ✅ 无冲突 |

此外 PLAN-00 R1 规划的变异核 **PIT 是 Java 实现**，同样与「必须 Go」红线冲突。

**影响范围**：仅 M3（Go 前端）/ M5（其余六语言前端）触及；**M1 不受影响**（M1 仅 Go 内核 + CES + CLI，全部 Go 标准库 + yaml.v3）。

**待用户决策（不阻塞 M1）**：
- 方案 A：七语言前端解析器一律改用 **Go 实现**（如 Go 版 tree-sitter 绑定、自研 Go 解析器、或 `go/ast` 仅覆盖 Go），彻底满足「必须 Go」；工作量上移到 M5。
- 方案 B：对「语言自身实现」的底层库（Rust 的 syn、Kotlin 编译器）做例外豁免，仅允许其作为解析产物来源，但**不链接进引擎**，解析结果经端口注入 AIR。
- 方案 C：在 M3/M5 阶段正式修订 PLAN-05，把非 Go 底层库替换为 Go 生态等价物。

本采购阶段**不下载任何非 Go 源码**（遵守红线），上述冲突在 M3/M5 启动前由用户拍板。

## 6. 辅助校验工具（可选，非引擎依赖）

| 组件 | 用途 | 说明 |
|---|---|---|
| PractRand | RNG 统计质量（PLAN-04 要求 128MB 档） | C++ 工具，**非 Go**、非引擎依赖；与「必须 Go」红线无关（仅测试辅助）。M1 先用 Go 内置卡方/频数检验自验，PractRand 作为后续可选增强。 |
