#!/usr/bin/env bash
# =============================================================================
# Abate 七语言覆盖量化（Language Coverage Quantifier）
#
# 用真实数据回答"各语言规则深度未量化"的质疑，不编造数字：
#   A. 规则源码统计：每语言 builtin 目录的 ID: 声明条数（专属规则）；
#   B. 实测命中统计：对每语言一份含典型坏味道的样例跑 check，统计命中规则类；
#   C. 共享 AIR 模板数：summary.rules（对任一语言相同）。
#
# 输出为 Markdown 表格（stdout），可重定向生成 docs/user/language-coverage.md。
# 全流程离线，现场从源码构建二进制。退出码：0 成功。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/.." || exit 2
export PATH="/usr/local/go/bin:$PATH:$PATH" GOPROXY=off CGO_ENABLED=0
unset GOFLAGS

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT
BIN="$TMP/abate"

go build -trimpath -o "$BIN" . >/dev/null 2>&1 || { echo "build failed" >&2; exit 1; }

# A. 规则源码条数（ID: 声明）
rule_count() {
  local dir="$1"
  grep -rhoE 'ID:[[:space:]]*"' "lang/rule/builtin/$dir" 2>/dev/null | wc -l | tr -d ' '
}

# B. 每语言样例
mkdir -p "$TMP/s"
cat > "$TMP/s/a.go"  <<'EOF'
package p
var foo_bar int
func BadName() {}
EOF
cat > "$TMP/s/b.py"  <<'EOF'
import os
x = 12345
class my_class:
    def BadMethod(self):
        l = 1
        return l
EOF
cat > "$TMP/s/c.js"  <<'EOF'
var BadName = 12345;
function bad_name() { var l = 1; return l; }
EOF
cat > "$TMP/s/d.c"   <<'EOF'
int Bad_Name = 12345;
int main(void) { int l = 1; return l; }
EOF
cat > "$TMP/s/e.cs"  <<'EOF'
class bad_class { int BadField = 12345; void BadMethod() {} }
EOF
cat > "$TMP/s/f.rs"  <<'EOF'
fn BadFunc() { let BadVar = 12345; }
struct bad_struct { BadField: i32 }
EOF
cat > "$TMP/s/g.kt"  <<'EOF'
class bad_class { fun BadFun() { val BadVar = 12345 } }
EOF

echo "<!-- 本表由 scripts/language-coverage.sh 现场实测生成，请勿手工编辑数字 -->"
echo "| 语言 | 专属规则数(源码 ID) | 样例命中规则类(实测) | 样例命中条数 | 规则族 |"
echo "|------|--------------------|---------------------|-------------|--------|"

for pair in "Go:go:a.go" "Python:py:b.py" "JavaScript:js:c.js" "C:c:d.c" "C#:cs:e.cs" "Rust:rs:f.rs" "Kotlin:kt:g.kt"; do
  name="${pair%%:*}"; rest="${pair#*:}"; ext="${rest%%:*}"; file="${rest##*:}"
  count="$(rule_count "$ext")"
  "$BIN" check "$TMP/s/$file" --format json --out "$TMP/$ext.json" >/dev/null 2>&1
  TMPDIR="$TMP" python3 -c '
import json, os, sys
name, count, ext = sys.argv[1], sys.argv[2], sys.argv[3]
tmp = os.environ["TMPDIR"]
try:
    d = json.load(open(f"{tmp}/{ext}.json")); s = d["summary"]
    ncls = len(s["byRule"]); ntot = s["total"]
    fams = sorted({k["key"].split("-")[1] for k in s["byRule"] if "-" in k["key"]})
    famtxt = "/".join(fams) if fams else "-"
    print(f"| {name} | {count} | {ncls} | {ntot} | {famtxt} |")
except FileNotFoundError:
    print(f"| {name} | {count} | - | - | (无报告) |")
' "$name" "$count" "$ext"
done

echo ""
echo "共享 AIR 结构模板数（summary.rules，各语言相同）："
"$BIN" check "$TMP/s/b.py" --format json --out "$TMP/shared.json" >/dev/null 2>&1
TMPDIR="$TMP" python3 -c 'import json,os; d=json.load(open(os.environ["TMPDIR"]+"/shared.json")); print(d["summary"]["rules"])' | sed 's/^/  /'
