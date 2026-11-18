#!/usr/bin/env bash
# =============================================================================
# Abate 能力证据门禁（Credibility Evidence Gate）
#
# 把 capabilities-and-limits.md 声称的每一项能力变成"可一键复现、带断言"的
# 仓库内证据，取代临时手工命令。现场从源码构建二进制，全部断言通过才退出 0：
#   1. 零依赖（go.mod 无第三方 require、无 vendor/）
#   2. 七语言按扩展名分派（混合目录 files=7、7 种规则前缀）
#   3. 真 AST 结构化解析（四层嵌套产出多类 PY-AIR-* 规则命中）
#   4. 容错解析边界（破损输入不拒绝，产出报告=非语法守门）
#   5. 确定性（同一输入两次 json 输出逐字节一致）
#   6. --languages 筛选
#   7. check --help 为 check 专用退出码语义（非 run 的"稳态违反"）
#
# 全流程离线。退出码：0 全通过；1 有断言失败。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/.." || exit 2
export PATH="/usr/local/go/bin:$PATH:$PATH" GOPROXY=off CGO_ENABLED=0
unset GOFLAGS

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT
BIN="$TMP/abate"

banner "Abate 能力证据门禁 - 现场构建 + 7 项可复现断言"

# --- 0. 现场构建 ---
banner "0. 从当前源码构建二进制"
if go build -trimpath -o "$BIN" . >/dev/null 2>&1; then
  pass "go build 成功"
else
  fail "go build 失败，无法继续"
  exit 1
fi

# --- 1. 零依赖 ---
banner "1. 零外部依赖"
REQ="$(grep -c '^require' go.mod || true)"
if [ "$REQ" -eq 0 ]; then pass "go.mod 无第三方 require 块"; else fail "go.mod 含 require 块（$REQ）"; fi
if [ ! -d vendor ]; then pass "无 vendor/ 目录"; else fail "存在 vendor/ 目录"; fi

# --- 2. 七语言分派 ---
banner "2. 七语言按扩展名分派"
mkdir -p "$TMP/mix"
printf 'package main\nfunc main() {}\n' > "$TMP/mix/a.go"
printf 'def f(x):\n  return x\n'        > "$TMP/mix/b.py"
printf 'function f(){return 1;}\n'      > "$TMP/mix/c.js"
printf 'int main(){return 0;}\n'        > "$TMP/mix/d.c"
printf 'class A { void M() {} }\n'      > "$TMP/mix/e.cs"
printf 'fn main() {}\n'                 > "$TMP/mix/f.rs"
printf 'class A {}\n'                   > "$TMP/mix/g.kt"
"$BIN" check "$TMP/mix" --format json --out "$TMP/mix.json" >/dev/null 2>&1
if [ -s "$TMP/mix.json" ]; then
  SUMMARY="$(python3 - "$TMP/mix.json" <<'PY'
import sys, json
d = json.load(open(sys.argv[1]))
s = d["summary"]
prefixes = sorted({k["key"].split("-")[0] for k in s["byRule"]})
print(s["files"], len(prefixes), ",".join(prefixes))
PY
)"
  FILES="$(echo "$SUMMARY" | awk '{print $1}')"
  NPREFIX="$(echo "$SUMMARY" | awk '{print $2}')"
  PLIST="$(echo "$SUMMARY" | awk '{print $3}')"
  if [ "$FILES" -eq 7 ]; then pass "混合目录 files=7"; else fail "混合目录 files=$FILES（期望 7）"; fi
  if [ "$NPREFIX" -ge 7 ]; then pass "规则前缀覆盖 7 种：$PLIST"; else fail "规则前缀仅 $NPREFIX 种：$PLIST"; fi
else
  fail "七语言混合检查未产出报告"
fi

# --- 3. 真 AST 结构化解析 ---
banner "3. 真 AST 结构化解析（嵌套结构命中多类规则）"
cat > "$TMP/nested.py" <<'PY'
def outer(x):
    if x:
        for i in range(3):
            y = i
            return y
    return x
PY
"$BIN" check "$TMP/nested.py" --format json --out "$TMP/nested.json" >/dev/null 2>&1
if [ -s "$TMP/nested.json" ]; then
  NSTAT="$(python3 - "$TMP/nested.json" <<'PY'
import sys, json
d = json.load(open(sys.argv[1]))
s = d["summary"]
py = {k["key"] for k in s["byRule"] if k["key"].startswith("PY-AIR-")}
print(s["total"], len(py))
PY
)"
  TOTAL="$(echo "$NSTAT" | awk '{print $1}')"
  NPY="$(echo "$NSTAT" | awk '{print $2}')"
  if [ "$TOTAL" -gt 0 ] && [ "$NPY" -ge 3 ]; then
    pass "嵌套结构产出 total=$TOTAL、$NPY 类 PY-AIR-* 结构规则（树上匹配证据）"
  else
    fail "嵌套结构命中不足：total=$TOTAL、PY-AIR 类=$NPY"
  fi
else
  fail "嵌套结构检查未产出报告"
fi

# --- 4. 容错解析边界 ---
banner "4. 容错解析边界（破损输入不拒绝 = 非语法守门）"
printf 'def f(:\n  x =\n' > "$TMP/broken.py"
"$BIN" check "$TMP/broken.py" --format json --out "$TMP/broken.json" >/dev/null 2>&1
BRC=$?
if [ -s "$TMP/broken.json" ]; then
  pass "破损语法输入仍产出报告（rc=$BRC，证实容错、非语法守门）"
else
  fail "破损语法输入未产出报告（与容错声明不符）"
fi

# --- 5. 确定性 ---
banner "5. 确定性（同一输入两次 json 逐字节一致）"
"$BIN" check "$TMP/nested.py" --format json --out "$TMP/d1.json" >/dev/null 2>&1
"$BIN" check "$TMP/nested.py" --format json --out "$TMP/d2.json" >/dev/null 2>&1
if cmp -s "$TMP/d1.json" "$TMP/d2.json"; then
  FINGER="$(python3 -c 'import json;print(json.load(open("'"$TMP"'/d1.json"))["fingerprint"])')"
  pass "两次输出逐字节一致（fingerprint=$FINGER）"
else
  fail "两次输出不一致（确定性漂移）"
fi

# --- 6. --languages 筛选 ---
banner "6. --languages 筛选"
"$BIN" check "$TMP/mix" --languages py --format json --out "$TMP/onlypy.json" >/dev/null 2>&1
if [ -s "$TMP/onlypy.json" ]; then
  PF="$(python3 -c 'import json;print(",".join(sorted({k["key"].split("-")[0] for k in json.load(open("'"$TMP"'/onlypy.json"))["summary"]["byRule"]})))')"
  PF_FILES="$(python3 -c 'import json;print(json.load(open("'"$TMP"'/onlypy.json"))["summary"]["files"])')"
  if [ "$PF" = "PY" ] && [ "$PF_FILES" -eq 1 ]; then
    pass "--languages py 仅检查 py（files=$PF_FILES，前缀=$PF）"
  else
    fail "--languages 筛选异常（files=$PF_FILES，前缀=$PF）"
  fi
else
  fail "--languages 筛选未产出报告"
fi

# --- 7. check 专用退出码文案 ---
banner "7. check --help 为 check 专用退出码语义"
HELP="$("$BIN" check --help 2>&1)"
if echo "$HELP" | grep -q "Exit codes:" && \
   echo "$HELP" | grep -q "including warnings" && \
   echo "$HELP" | grep -q -- "--severity error" && \
   ! echo "$HELP" | grep -q "steady state VIOLATED"; then
  pass "check --help 为专用语义（0/1/2/3 + --severity error），无 run 的稳态违反文案"
else
  fail "check --help 文案不符合 check 专用语义"
fi

# --- 汇总 ---
echo ""
if [ "$FAILED" -eq 0 ]; then
  echo -e "${GREEN}============================================${NC}"
  echo -e "${GREEN} ALL EVIDENCE CHECKS PASSED — 能力声明可复现 ${NC}"
  echo -e "${GREEN}============================================${NC}"
  exit 0
else
  echo -e "${RED}============================================${NC}"
  echo -e "${RED} SOME EVIDENCE CHECKS FAILED — 详见上方 [FAIL] ${NC}"
  echo -e "${RED}============================================${NC}"
  exit 1
fi
