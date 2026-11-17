#!/usr/bin/env bash
# =============================================================================
# Abate 版权头门禁（XCU-GOV）
# 所有源文件必须携带统一版权头（见架构 §10 强制版权头 + 数字水印）。
# 缺失即 FAIL。等价于 goheader 的约束，但以脚本实现（goheader 未在镜像可用）。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0; REPORT=""

cd "$(dirname "$0")/.." || exit 2

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

banner "Abate Header Gate - 校验统一版权头"

VIOLATIONS=0
while IFS= read -r f; do
  [ -z "$f" ] && continue
  # 取文件前 8 行，检查版权头关键标记。
  head8="$(head -n 8 "$f" 2>/dev/null || true)"
  if ! printf '%s' "$head8" | grep -q 'Copyright'; then
    fail "缺少版权头 Copyright: $f"; VIOLATIONS=1; continue
  fi
  if ! printf '%s' "$head8" | grep -q 'xiaolibuluo/Abate'; then
    fail "版权头缺少仓库来源(xiaolibuluo/Abate): $f"; VIOLATIONS=1; continue
  fi
  if ! printf '%s' "$head8" | grep -q '小黎的部落'; then
    fail "版权头缺少水印(小黎的部落): $f"; VIOLATIONS=1; continue
  fi
done < <(find . -name '*.go' -not -path './vendor/*' -not -path './third_party/*' -not -name '*.pb.go' | sort)

[ "$VIOLATIONS" -eq 0 ] && pass "全部源文件版权头齐全（含小黎的部落水印）"

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "HEADER GATE: FAILED"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
banner "HEADER GATE: PASSED"
exit 0
