#!/usr/bin/env bash
# =============================================================================
# Abate 零第三方源码审计（DESIGN §1.1 口径 C2）
#
# 对以下对象做真实扫描，不读硬编码结论、不放宽判定：
#   1. third_party/_sources/、third_party/_references/、vendor/ 是否仍存在；
#   2. docs/benchmark/ 下是否仍有开源源码文件（扩展名 + 目录）；
#   3. kernel/mutation/testdata/pit/ 上游 baseline 源码与官方产物；
#   4. go.mod 是否仍有第三方 require；
#   5. 产品代码 .go import 是否仍引用第三方模块；
#   6. 产品代码 .go 是否仍有上游来源归属注释。
#
# 退出码：0 = 第三方源码副本 = 0；1 = 仍存在（列出每一项）。
# 清除前运行应报 N>0（证明非空跑）；清除后运行必须报 0。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FINDINGS=0

cd "$(dirname "$0")/.." || exit 2

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
finding() { FINDINGS=$((FINDINGS + 1)); echo -e "${RED}[FOUND ${FINDINGS}] $1${NC}"; }
ok() { echo -e "${GREEN}[OK] $1${NC}"; }

banner "Abate Zero-Third-Party Audit - 第三方源码副本扫描"

# ---- 1. 第三方目录 ----
echo ">>> [1] 第三方源码目录"
for d in third_party/_sources third_party/_references vendor; do
  if [ -d "$d" ]; then
    n="$(find "$d" -type f 2>/dev/null | wc -l | tr -d ' ')"
    sz="$(du -sh "$d" 2>/dev/null | cut -f1)"
    finding "目录仍存在: $d（${n} 个文件，${sz}）"
  else
    ok "不存在: $d"
  fi
done

# ---- 2. benchmark 源码残留 ----
echo ">>> [2] docs/benchmark 开源源码"
if [ -d docs/benchmark ]; then
  for d in docs/benchmark/cases docs/benchmark/config; do
    if [ -d "$d" ]; then
      n="$(find "$d" -type f 2>/dev/null | wc -l | tr -d ' ')"
      finding "benchmark 源码/依赖目录仍存在: $d（${n} 个文件）"
    fi
  done
  # 只扫描 cases/ 下的上游语料副本；scripts/ 中的 .py 是本仓自有的测量脚本，
  # 按 DESIGN §1.2 保留，不属于第三方源码。
  SRC_HITS="$(find docs/benchmark/cases -type f \
    \( -name '*.js' -o -name '*.mjs' -o -name '*.cjs' -o -name '*.ts' -o -name '*.tsx' \
       -o -name '*.c' -o -name '*.h' -o -name '*.cc' -o -name '*.cpp' -o -name '*.java' \
       -o -name '*.py' -o -name '*.rs' -o -name '*.go' -o -name '*.cs' -o -name '*.kt' \) \
    2>/dev/null | head -20)"
  if [ -n "$SRC_HITS" ]; then
    finding "benchmark 下仍有源码文件（前 20 条）:"
    printf '%s\n' "$SRC_HITS"
  fi
  if [ -d docs/benchmark/config/node_modules ]; then
    finding "node_modules 仍存在: docs/benchmark/config/node_modules"
  fi
else
  ok "docs/benchmark 不存在"
fi

# ---- 3. PIT 上游 baseline 源码 / 官方产物 ----
echo ">>> [3] PIT 上游 baseline 源码与官方产物"
for f in \
  kernel/mutation/testdata/pit/baseline/pom.xml \
  kernel/mutation/testdata/pit/baseline/src-main/java/com/abate/parity/ParityTarget.java \
  kernel/mutation/testdata/pit/baseline/src-test/java/com/abate/parity/ParityTargetTest.java \
  kernel/mutation/testdata/pit/pit-mutations.json \
  kernel/mutation/testdata/pit/pit-mutations.xml \
  kernel/mutation/testdata/pit/pit-console-report.txt; do
  if [ -e "$f" ]; then finding "PIT 上游文件仍存在: $f"; fi
done
if [ -d kernel/mutation/testdata/pit/baseline ]; then
  finding "PIT baseline 目录仍存在: kernel/mutation/testdata/pit/baseline"
fi

# ---- 4. go.mod 第三方 require ----
echo ">>> [4] go.mod 第三方 require"
REQ="$({
  sed -n -E '/^require \(/,/^\)/ s/^[[:space:]]*([^ )]+) v[0-9].*/\1/p' go.mod
  sed -n -E 's/^require[[:space:]]+([^ ]+) v[0-9].*/\1/p' go.mod
} | sort -u)"
if [ -n "$REQ" ]; then
  finding "go.mod 仍有第三方 require: $(echo "$REQ" | tr '\n' ' ')"
else
  ok "go.mod 第三方 require = 0"
fi

# ---- 5. 产品代码 .go 第三方 import ----
echo ">>> [5] 产品代码 .go import"
IMPORT_HITS="$(grep -rEn '"(github\.com|gopkg\.in|golang\.org/x)/' --include='*.go' \
  cmd ctl ces kernel lang plugin sim internal security 2>/dev/null | head -20 || true)"
if [ -n "$IMPORT_HITS" ]; then
  finding "产品代码仍 import 第三方模块（前 20 条）:"
  printf '%s\n' "$IMPORT_HITS"
else
  ok "产品代码第三方 import = 0"
fi

# ---- 6. 上游来源归属注释 ----
echo ">>> [6] 产品代码上游来源归属注释"
MARK_HITS="$(grep -rEni \
  'Adapted from (gvisor|PIT|chaosblade)|参考 (PITest|PIT |chaosblade|gvisor|wazero|tree-sitter|swc|roslyn|libclang)|upstream source \(snapshotted|PITest [0-9]+\.[0-9]+' \
  --include='*.go' cmd ctl ces kernel lang plugin sim internal security 2>/dev/null | head -20 || true)"
if [ -n "$MARK_HITS" ]; then
  finding "产品代码仍有上游来源归属注释（前 20 条）:"
  printf '%s\n' "$MARK_HITS"
else
  ok "产品代码上游来源归属注释 = 0"
fi

echo ""
if [ "$FINDINGS" -eq 0 ]; then
  echo -e "${GREEN}[VERDICT] 第三方源码副本 = 0${NC}"
  exit 0
fi
echo -e "${RED}[VERDICT] 第三方源码副本 = ${FINDINGS} 项（未清空）${NC}"
exit 1
