<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 命令行参考（CLI Reference）

> 由 `abate --help` 与 `abate <cmd> --help` 的真实输出整理。命令树无状态，
> 相同参数多次运行产生相同字节（确定性）。

## 命令树

| 命令 | 作用 |
|---|---|
| `run <ces-file>` | 解析/归一化/校验 CES 并执行确定性实验，输出报告。`--record` 还可写出 SIM-REC 档案。 |
| `verify <a> <b>` | 两份运行报告逐字节比对。 |
| `replay <archive> [other]` | 校验档案（帧哈希链 + 全档 SHA-256），可选比对两份。 |
| `check <paths...>` | 按扩展名分派到 7 语言前端（`.go/.py/.js/.c/.cs/.rs/.kt`），运行对应语言规则集；输出 console/JSON/SARIF。 |
| `derive` | 从规则派生 M 族算子并跑变异实验（静态 AIR 级 killing 近似）。 |
| `plugin` | 校验/加载/调用/列举/热换签名的插件。 |
| `web` | 启动控制 API 并内嵌浏览器控制台。 |
| `version` | 打印构建版本与溯源。 |
| `help` / `--help` / `-h` | 打印帮助。 |

## abate check

七语言源码检查，按文件扩展名自动分派到对应前端与规则集：

```
Usage: abate check <paths...> [flags]
  --format console|json|sarif  报告格式（默认 console）。
  --severity error|warning|info  最低严重级别（默认 info）。
  --rules prefix[,prefix...]   只跑指定规则 ID 前缀。
  --languages lang[,lang...]   只检查指定语言（go,py,js,c,cs,rs,kt）。
  --out path                   同时把报告字节写入文件。
```

| 扩展名 | 语言 | 规则前缀 |
|---|---|---|
| `.go` | go | `GO-` |
| `.py` | py | `PY-` |
| `.js` `.jsx` `.mjs` `.cjs` | js | `JS-` |
| `.c` `.h` | c | `C-` |
| `.cs` | cs | `CS-` |
| `.rs` | rs | `RS-` |
| `.kt` `.kts` | kt | `KT-` |

四种退出码：`0` 无违规；`1` 检出违规（业务失败）；`2` 内部错误；`3` 用法错误。

完整保姆式教程见 `docs/user/multilang-check-guide.md`。

## abate run

```
Usage: abate run <ces-file> [flags]
Flags:
  --seed uint        确定性种子覆盖（0 用文档种子）。
  --format console|json  报告格式（默认 console）。
  --out path          同时把报告字节写入文件。
  --record path       写出确定性 SIM-REC 档案。
  --dry-run           仅解析+校验，不执行。
  --strict            将 warning 级违规提升为失败。
  --help, -h          打印帮助。
```

退出码（PLAN-07 §4）：

| 码 | 含义 |
|---|---|
| 0 | 成功（合法 CES 且稳态 HELD） |
| 1 | 业务失败（执行中稳态 VIOLATED） |
| 2 | 内部错误（调度或报告输出失败） |
| 3 | 用法错误（参数非法、CES 文档非法、CES-VALID 违规；`--strict` 将 warning 也提升为此码） |

## abate version

```
abate 1.2.27
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
```

> 引擎内部版本为 `0.1.0`（与发布版本 `vX.Y.Z` 分属构建溯源与发布节拍两个维度）；
> 发布资产由 `scripts/release.sh` 注入 `ldflags` 版本/commit 元数据。

## 复制即用示例

```bash
abate run ctl/cli/testdata/ok.yaml
abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
abate replay /tmp/run.arch
abate verify report-a.json report-b.json
abate check ./ctl/cli
abate check ./lang/rule --format sarif --out /tmp/check.sarif
abate derive --seed 42 --format json
abate plugin load --manifest plugin-manifest.json --type operator
abate plugin call --id abate/operator --entry apply --args ZWJhdGU=
abate web --listen 127.0.0.1:8430
abate version
```
