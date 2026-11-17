#!/usr/bin/env bash
# =============================================================================
# Abate 全语料降级率测量（ABATE-FULLCORPUS-DEGRADE-001）
#
# 对自研测试语料全部条目逐条以真实前端 parse，合并观测到的 AIR 节点 kind，
# 应用与冻结 E2E 完全相同的 Degradation 公式，产出确定性的 JSON 报告。
# 与 E2E 的单文件（首 unit）口径互补：本脚本给出全语料口径的真实降级率。
#
# 反造假约束：degraded 映射绝不重分类为 exact；观测由真实 parse 产出；
# 双跑逐字节 cmp 通过才归档；报告附 SHA-256。红灯即 exit 1。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/.." || exit 2

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

REPORT_DIR="docs/refactor/fullcorpus-degrade"
REPORT="$REPORT_DIR/fullcorpus-degrade.json"
DIGEST="$REPORT_DIR/fullcorpus-degrade.json.sha256"

banner "Abate Full-Corpus Degrade - 全语料真实降级率测量"
mkdir -p "$REPORT_DIR"

RUN1="$(mktemp)"; RUN2="$(mktemp)"
trap 'rm -f "$RUN1" "$RUN2"' EXIT

go run ./cmd/degradecheck > "$RUN1" 2>/dev/null
RC1=$?
go run ./cmd/degradecheck > "$RUN2" 2>/dev/null
RC2=$?
if [ "$RC1" -ne 0 ] || [ "$RC2" -ne 0 ]; then
  fail "degradecheck 退出非零 (run1=$RC1 run2=$RC2)"
elif ! cmp -s "$RUN1" "$RUN2"; then
  fail "双跑报告逐字节不一致（确定性破缺）"
else
  pass "双跑逐字节一致（确定性成立）"
fi

if [ "$FAILED" -eq 0 ]; then
  cp "$RUN1" "$REPORT"
  SHA="$(sha256sum "$REPORT" | cut -d' ' -f1)"
  printf '%s\n' "$SHA" > "$DIGEST"
  pass "报告已归档: $REPORT"
  pass "SHA-256: $SHA"
  echo ""
  jq -r '.results[] | "  \(.language): \(.status) degrade=\(.degradePct)% (files=\(.files), parseFail=\(.parseFail))"' "$REPORT"
fi

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "FULL-CORPUS DEGRADE: FAILED"; exit 1
fi
banner "FULL-CORPUS DEGRADE: PASSED"
exit 0
