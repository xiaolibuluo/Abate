<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 七语言源码检查保姆式教程（Multilang Check Guide）

> 本文面向第一次使用 `abate check` 的人。从构建二进制开始，到独立检查 7 种语言、
> 混合目录一次检查、输出机器可读报告、按语言筛选，每一步都给出可直接复制的命令
> 和预期输出。全部命令已在本机实测通过。

## 目录

- [0. 你能用它做什么](#0-你能用它做什么)
- [1. 准备：构建二进制](#1-准备构建二进制)
- [2. 第一个例子：检查一个 Go 文件](#2-第一个例子检查一个-go-文件)
- [3. 七语言逐一手把手](#3-七语言逐一手把手)
- [4. 一次检查整个项目（混合目录）](#4-一次检查整个项目混合目录)
- [5. 只看某几种语言 / 某几条规则](#5-只看某几种语言--某几条规则)
- [6. 输出 JSON / SARIF 报告](#6-输出-json--sarif-报告)
- [7. 理解退出码](#7-理解退出码)
- [8. 常见问题排查](#8-常见问题排查)
- [9. 扩展名与语言对照表](#9-扩展名与语言对照表)

---

## 0. 你能用它做什么

`abate check` 会**读取源码 → 用对应语言的解析器生成 AIR（统一中间表示）→ 跑该语言的自研规则集 → 输出确定性报告**。

- 支持 **7 种语言**：Go、Python、JavaScript、C、C#、Rust、Kotlin；
- 按**文件扩展名自动分派**，你不用手动指定语言；
- 报告格式有 console（人看）、json（机器读）、sarif 2.1.0（CI/代码平台集成）；
- 相同输入**逐字节一致**输出，可重复、可比对、可入 git。

它**不修改你的源码**（只检查，不格式化、不自动修复）。

---

## 1. 准备：构建二进制

项目零外部依赖，离线即可构建：

```bash
cd /path/to/Abate
export GOPROXY=off
go build -o abate .
./abate version
```

预期输出（第一行）：

```
abate 1.2.29
```

确认 `check` 帮助里有 7 语言说明：

```bash
./abate check --help
```

应看到 `Seven languages are dispatched by file extension: .go .py .js .c .cs .rs .kt`。

> 下面命令统一用 `./abate` 指代刚构建的二进制。你也可以把它放到 PATH 里，直接写 `abate`。

---

## 2. 第一个例子：检查一个 Go 文件

新建一个故意写得不好的 Go 文件（含下划线命名，触发命名规则）：

```bash
mkdir -p /tmp/check-demo
cat > /tmp/check-demo/demo.go <<'EOF'
package demo

var foo_bar int
EOF
```

运行检查：

```bash
./abate check /tmp/check-demo/demo.go
```

你会看到类似输出（行号/列号以实际为准）：

```
/tmp/check-demo/demo.go:3:5: GO-EG-001: identifier foo_bar should not contain underscores [info]
```

逐项解读：
- `/tmp/check-demo/demo.go:3:5` → 文件、行、列；
- `GO-EG-001` → 规则编号，`GO-` 前缀表示 Go 语言规则；
- `[info]` → 严重级别。

此时终端退出码是 `1`（检出违规），这是**正常的业务失败**，不是程序出错：

```bash
echo $?   # 输出 1
```

---

## 3. 七语言逐一手把手

下面为每种语言创建一个最小样例并单独检查。每条命令都可独立复制执行。

### 3.1 Go

```bash
cat > /tmp/check-demo/sample.go <<'EOF'
package p
var foo_bar int
EOF
./abate check /tmp/check-demo/sample.go
```

### 3.2 Python

```bash
cat > /tmp/check-demo/sample.py <<'EOF'
def greet(name):
    result = 1
    return result
EOF
./abate check /tmp/check-demo/sample.py
```

### 3.3 JavaScript

```bash
cat > /tmp/check-demo/sample.js <<'EOF'
function greet(name) {
  var result = 1;
  return result;
}
EOF
./abate check /tmp/check-demo/sample.js
```

### 3.4 C

```bash
cat > /tmp/check-demo/sample.c <<'EOF'
int add(int a, int b) {
    int foo_bar = a + b;
    return foo_bar;
}
EOF
./abate check /tmp/check-demo/sample.c
```

### 3.5 C#

```bash
cat > /tmp/check-demo/sample.cs <<'EOF'
public class Greeter {
    public int Add(int a, int b) {
        var foo_bar = a + b;
        return foo_bar;
    }
}
EOF
./abate check /tmp/check-demo/sample.cs
```

### 3.6 Rust

```bash
cat > /tmp/check-demo/sample.rs <<'EOF'
fn add(a: i32) -> i32 {
    let foo_bar = a;
    foo_bar
}
EOF
./abate check /tmp/check-demo/sample.rs
```

### 3.7 Kotlin

```bash
cat > /tmp/check-demo/sample.kt <<'EOF'
fun add(a: Int): Int {
    val foo_bar = a
    return foo_bar
}
EOF
./abate check /tmp/check-demo/sample.kt
```

每个命令都会输出对应语言前缀的规则编号（`PY-`、`JS-`、`C-`、`CS-`、`RS-`、`KT-`），
证明该文件被正确分派到了该语言的前端和规则集。

---

## 4. 一次检查整个项目（混合目录）

把上面 7 个样例都放进 `/tmp/check-demo/` 后，直接检查整个目录：

```bash
./abate check /tmp/check-demo --format json
```

`abate` 会：
1. 递归扫描目录（跳过 `vendor/`、`third_party/`、`.git/`）；
2. 按扩展名把每个文件分派给对应语言前端；
3. 汇总所有发现，按「文件路径 → 行 → 列 → 规则号」固定排序；
4. 输出一份统一报告。

可以从 JSON 里看到每种语言都贡献了发现（`ruleId` 的前缀即语言）：
`GO-` / `PY-` / `JS-` / `C-` / `CS-` / `RS-` / `KT-`。

一次运行同时处理多个目录或多个文件也可以：

```bash
./abate check ./src ./tests extra.py --format json
```

---

## 5. 只看某几种语言 / 某几条规则

**只看指定语言**（用 `--languages`，逗号分隔语言标签）：

```bash
# 只检查 Python 和 JavaScript，目录里其他语言文件被忽略
./abate check /tmp/check-demo --languages py,js
```

**只看指定规则**（用 `--rules`，逗号分隔规则 ID 前缀）：

```bash
# 只跑所有 PY-PEP8 开头的规则
./abate check /tmp/check-demo --rules PY-PEP8
```

**只看某语言 + 某严重级别**（`--severity` 是下限）：

```bash
# 只报 error 及以上
./abate check /tmp/check-demo --severity error
```

语言标签与规则前缀对照见[第 9 节](#9-扩展名与语言对照表)。

---

## 6. 输出 JSON / SARIF 报告

**JSON**（适合脚本处理）：

```bash
./abate check /tmp/check-demo --format json --out report.json
head -20 report.json
```

报告含 `schema: abate.check.v1`、`summary`（files/rules/total）、`findings`、`fingerprint`。
`--out` 会把报告字节同时写入文件，内容与 stdout 完全一致。

**SARIF 2.1.0**（适合 GitHub/GitLab 代码扫描集成）：

```bash
./abate check /tmp/check-demo --format sarif --out report.sarif
```

SARIF 报告不含时间戳，字段顺序固定，可安全入库做差分。

**验证确定性**（同输入两次运行逐字节一致）：

```bash
./abate check /tmp/check-demo --format json --out a.json
./abate check /tmp/check-demo --format json --out b.json
cmp a.json b.json && echo "一致"   # 应输出"一致"
```

---

## 7. 理解退出码

`check` 遵循 PLAN-07 §4 的四态退出码：

| 码 | 含义 | 你该怎么做 |
|---|---|---|
| `0` | 无违规 | 通过 |
| `1` | 检出违规（业务失败） | 看报告修代码；这是最常见的"有发现"状态 |
| `2` | 内部错误（前端/输出故障） | 报 bug，附命令和输出 |
| `3` | 用法错误（参数非法、路径不存在、无匹配规则） | 检查命令拼写、路径、`--rules`/`--languages` 值 |

在 CI 里通常把 `0` 视为通过，`1` 视为需要人工处理的发现，`2`/`3` 视为配置错误。

---

## 8. 常见问题排查

**Q：`check` 返回 1，是不是程序坏了？**
不是。1 表示"检出了违规"，是规则引擎正常工作的结果。看报告里的规则编号和位置即可。若想让它返回 0，用 `--rules` 限定到没有发现的子集，或修复代码。

**Q：为什么某个文件没被检查？**
只检查[第 9 节](#9-扩展名与语言对照表)列出的扩展名。`.ts`、`.cpp`、`.java` 等暂不支持；被 `vendor/`/`third_party/`/`.git/` 目录包含的文件会被跳过。

**Q：`--languages xx` 报用法错误？**
语言标签只能是 `go,py,js,c,cs,rs,kt`（小写）。用别的（如 `python`、`golang`）会没有匹配规则而报 3。

**Q：报告里没有某语言的发现？**
确认文件扩展名正确、该语言有规则集加载。可用 `./abate check <file> --languages <tag> --format json` 单独验证，看 `summary.rules` 是否大于 0。

**Q：路径里有空格怎么办？**
给路径加引号：`./abate check "my dir/sample.py"`。

**Q：`check` 会改我的源码吗？**
不会。它是纯只读检查，从不写源文件（`--out` 只写报告文件）。

---

## 9. 扩展名与语言对照表

| 扩展名 | 语言标签 | 规则前缀 |
|---|---|---|
| `.go` | `go` | `GO-` |
| `.py` | `py` | `PY-` |
| `.js` `.jsx` `.mjs` `.cjs` | `js` | `JS-` |
| `.c` `.h` | `c` | `C-` |
| `.cs` | `cs` | `CS-` |
| `.rs` | `rs` | `RS-` |
| `.kt` `.kts` | `kt` | `KT-` |

> 更多命令见 `docs/user/cli-reference.md`；规范原文见 `docs/spec/ces-v1.md`；
> 语言前端与保真度口径见 `docs/ARCHITECTURE.md` 与 `docs/user/language-coverage.md`。
