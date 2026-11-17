#!/usr/bin/env bash
# =============================================================================
# Abate 代码风格门禁（XCU-CI 自定义门禁段）
# 规则：
#   1. 单文件上限 300 行（硬 FAIL）。
#   2. 一类一个文件：每个非测试 .go 文件至多 1 个导出类型（未导出辅助类型不限）。
#      所有文件（含历史基石文件）一视同仁，无任何白名单豁免；命中即按规则拆分。
#   3. 禁止出现 TODO/FIXME/BUG/XXX/HACK 占位注释（与 godox 双重约束，脚本兜底扫描）。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0; REPORT=""

cd "$(dirname "$0")/.." || exit 2

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

MAX_LINES=300

banner "Abate Style Gate - 文件长度 / 一类一个文件 / 占位注释"

VIOLATIONS=0
while IFS= read -r f; do
  [ -z "$f" ] && continue

  # ---- 规则 1：单文件 ≤ 300 行 ----
  n=$(wc -l < "$f" | tr -d ' ')
  if [ "$n" -gt "$MAX_LINES" ]; then
    fail "文件超过 ${MAX_LINES} 行 (${n} 行): $f"; VIOLATIONS=1
  fi

  # ---- 规则 3：禁止占位注释 ----
  if grep -nE '\b(TODO|FIXME|BUG|XXX|HACK)\b' "$f" >/dev/null 2>&1; then
    fail "存在占位注释(TODO/FIXME/BUG/XXX/HACK): $f"; VIOLATIONS=1
  fi

  # ---- 规则 2：一类一个文件（仅非测试）----
  case "$f" in
    *_test.go) continue ;;
  esac

  # 统计导出类型（行首 `type Xxx` 且首字母大写）。
  exp=$(grep -cE '^type[[:space:]]+[A-Z][A-Za-z0-9_]*' "$f" 2>/dev/null || true)
  if [ "$exp" -gt 1 ]; then
    fail "单文件导出类型超过 1 个 (${exp} 个，违反「一类一个文件」): $f"; VIOLATIONS=1
  fi
done < <(find . -name '*.go' -not -path './vendor/*' -not -path './third_party/*' | sort)

[ "$VIOLATIONS" -eq 0 ] && pass "文件长度≤${MAX_LINES} / 一类一个文件 / 无占位注释 全部通过"

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "STYLE GATE: FAILED"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
banner "STYLE GATE: PASSED"
exit 0
