<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# CLI 参考

> 本文描述当前二进制（`abate 1.2.29`）的全部命令与退出码。每个命令的帮助可用 `abate <cmd> --help` 现场查看。

## 命令总览

| 命令 | 用途 |
|---|---|
| `run <ces-file>` | 执行 CES 实验并输出确定性报告 |
| `verify <a> <b>` | 逐字节比较两份报告 |
| `replay <archive> [other]` | 校验 SIM-REC 档案，可选比较两份 |
| `check <paths...>` | 七语言结构检查 |
| `derive` | 从 Go 规则派生变异算子并跑一次变异实验 |
| `mutate <path>` | 对真实 Go 工程跑确定性变异实验（真实 go test 判活） |
| `targetgen` | 生成 7 语言确定性变异测试目标项目 |
| `plugin <subcmd>` | 签名插件管理 |
| `web` | 本地控制 API + 浏览器控制台 |
| `version` | 版本与 provenance |
| `--help` | 帮助 |

---

## run

解析、规范化、校验 CES 文档，驱动确定性内核实验并输出报告。`--dry-run` 只校验不执行。

```
abate run <ces-file> [flags]
```

| Flag | 说明 |
|---|---|
| `--seed uint` | 确定性种子覆盖（0 用文档自带种子） |
| `--format console\|json` | 报告格式（默认 console） |
| `--out path` | 报告字节另写到文件 |
| `--record path` | 写 SIM-REC 确定性档案 |
| `--dry-run` | 只解析校验，不执行 |
| `--strict` | 把 warning 级违反升为失败 |
| `--help, -h` | 帮助 |

示例：
```bash
abate run ctl/cli/testdata/ok.yaml
abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
abate run ctl/cli/testdata/ok.yaml --dry-run
```

## verify

逐字节比较两份 run 报告，打印两者 SHA-256、大小、首个差异偏移与映射退出码。完全相同退出 0，任意差异退出 1，文件读不到退出 3。

```
abate verify <a> <b> [--format console|json] [--out path]
```

## replay

加载 SIM-REC 档案，校验其 hash 链、全序序列与虚拟时间单调性；可选比较第二份档案并输出结构化 Diff。档案只读，不伪造观测。

```
abate replay <archive> [compare-archive] [flags]
```
Flags：`--format console|json`、`--out path`、`--help`。

## check

按扩展名分派七语言 AIR 前端，运行对应规则集，输出确定性 console / JSON / SARIF 2.1.0 报告。路径可为文件或目录；`vendor`、`third_party`、`.git` 会被跳过。

```
abate check <paths...> [flags]
```

| Flag | 说明 |
|---|---|
| `--format console\|json\|sarif` | 报告格式（默认 console） |
| `--severity error\|warning\|info` | 最低严重度（默认 info） |
| `--rules prefix[,prefix...]` | 只跑指定规则 ID 前缀 |
| `--languages lang[,lang...]` | 只检查指定语言（go,py,js,c,cs,rs,kt） |
| `--out path` | 报告字节另写到文件 |
| `--help, -h` | 帮助 |

`check` 不修改源码；规则→算子派生属于独立的 `derive` 命令。

示例：
```bash
abate check ./ctl/cli
abate check app.py report.js --format sarif --out /tmp/check.sarif
abate check ./src --languages py,js --rules PY-PEP8
abate check ./lang/air --rules GO-EG --format json
```

## derive

从内置 Go 规则集派生 M 族变异算子，把内存中的 Go 样例解析成 AIR，注册算子后跑一次确定性变异实验。
**killing 是静态 AIR 级近似，不是真实 go test 执行。**

```
abate derive [flags]
```
Flags：`--seed uint`（默认 42）、`--format console|json`、`--out path`、`--source path`（可选 Go 源文件替代内置样例）、`--help`。

示例：`abate derive --seed 42 --format json`


## mutate

对真实 Go 工程目录做确定性变异实验：在私有暂存副本中重写可变 token，每个变异体的存活判定由真实的 `go build` + `go test -run` 执行结果给出（KILLED / SURVIVED），与 `derive` 的静态近似判定不同。

```
abate mutate <path> [flags]
```

| Flag | 说明 |
|---|---|
| `--seed uint` | 确定性种子（默认 42） |
| `--format console\|json` | 报告格式（默认 console） |
| `--out path` | 报告字节另写到文件 |
| `--parallel n` | 并行度；>1 时钳制到 CPU 数，报告与顺序执行字节一致 |
| `--chaos` | 注入并回滚 L6 注入内核的 12 个 C 族混沌故障，报告增加确定性 chaos 摘要块 |
| `--native` | vendor/node_modules/target 复制一次并符号链接到各变异体工作区 |
| `--help, -h` | 帮助 |

示例：
```bash
abate targetgen --out /tmp/tgt --seed 42
abate mutate /tmp/tgt --seed 42
abate mutate /tmp/tgt --seed 42 --parallel 4
abate mutate ./lang/gomutate --seed 42 --format json --out /tmp/mut.json
```

常见错误：
- `path` 不是 Go 工程（无 go.mod / 无测试）→ 退出码 3，先用 `targetgen` 生成工程跑通流程。
- 机器 CPU 较少时 `--parallel` 过大 → 自动钳制，不报错，报告内容不受影响。

## targetgen

生成确定性变异测试目标项目：7 语言（go/py/js/c/cs/rs/kt）可选，同一种子产物字节一致；每个生成结构都有真实测试守护，交付前经真实工具链自检。

```
abate targetgen --out <dir> [flags]
```

| Flag | 说明 |
|---|---|
| `--out dir` | 输出目录，**必须不存在**（防覆盖） |
| `--lang` | 目标语言：go/py/js/c/cs/rs/kt（默认 go） |
| `--seed uint` | 确定性种子（默认 42） |
| `--module path` | Go module 路径（默认 example.com/target） |
| `--package name` | 包名（默认 target） |
| `--funcs n` | 生成函数数量 1..64（默认 8） |
| `--level` | 难度：easy/medium/hard（默认 medium） |
| `--mixed f` | 混合比例 [0,1]（默认 0） |
| `--complex` | 启用 slice/map/error/nested/method/iface 族结构 |
| `--skip-check` | 跳过交付前的工具链自检（不推荐） |
| `--help, -h` | 帮助 |

示例：
```bash
abate targetgen --out /tmp/tgt --seed 42 --funcs 14 --level easy
abate targetgen --out /tmp/tgt-py --lang py --seed 42
abate targetgen --out /tmp/tgt2 --seed 42
```

常见错误：
- `--out` 目录已存在 → 退出码 3，删除或换路径后重试。
- `--lang cs/rs` 时本机无对应工具链 → 自检阶段失败；加 `--skip-check` 可跳过自检仅生成源码。

## plugin

原生插件 host 中的签名插件管理。每次加载都会用 manifest 文件中的公钥校验 Ed25519 签名；未签名或被篡改的 manifest 会被拒绝、绝不注册。

```
abate plugin <subcommand> [flags]
```

| 子命令 | 说明 |
|---|---|
| `load --manifest <file> [--type operator\|reporter\|adapter]` | 校验并加载签名 manifest |
| `list` | 按确定性顺序列出已注册插件 id |
| `call --id <id> --entry <entry> [--args <base64>]` | 调用入口（apply / build / detect / emit） |
| `hotswap --id <id> --new-manifest <file>` | 原子替换实例，不中断在途调用 |

## web

单进程同时提供本地控制 API 与内嵌浏览器控制台，同源提供，无需 CORS，API 不单独暴露。

```
abate web [flags]
```

| Flag | 说明 |
|---|---|
| `--listen host:port` | 绑定地址（默认 127.0.0.1:8430） |
| `--token string` | API Bearer token；非 loopback 绑定**必填**，拒绝无鉴权暴露 |
| `--rest string` | 外部 REST base URL（设置后 /api 反向代理，控制台仍本地提供） |
| `--static dir` | 从目录提供静态资源（仅开发） |
| `--help, -h` | 帮助 |

## version

打印版本与 provenance：`abate 1.2.29` / `schema: ces.tencent.com/v1` / `build: deterministic` / `license` / `source` / `watermark`。

---

## 退出码

顶层 / run / derive 的通用语义（PLAN-07 §4）：

| 码 | 含义 |
|---|---|
| 0 | 成功（CES 合法且稳态 HELD） |
| 1 | 业务失败（执行期稳态 VIOLATED） |
| 2 | 内部错误（调度器或报告输出失败） |
| 3 | 用法错误（参数非法、CES 非法、CES-VALID error；`--strict` 把 warning 升到此码） |

<a id="退出码"></a>

**check 专用**：0 无发现；1 有发现（含 warning，`--severity error` 只让 error 失败）；2 内部错误（前端或报告输出）；3 用法错误（参数非法、缺路径、无规则匹配）。

**replay 专用**：0 验证通过（两份则相等）；1 比较差异（打印结构化 Diff）；2 内部错误（档案解析/完整性/报告输出）；3 用法错误（参数非法或档案路径不可读）。

**plugin**：0 成功；1 业务失败（签名/能力/hotswap 拒绝）；2 内部错误（输出/编码）；3 用法错误（flag 非法、manifest 不可读或格式错误）。

**verify**：0 完全一致；1 有差异；3 文件不可读。
