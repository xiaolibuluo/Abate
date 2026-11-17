#!/usr/bin/env bash
# =============================================================================
# Abate 自研 golden 语义回归编排脚本。
#
# 单一离线模式：用本仓 kernel/mutation 引擎对同构夹具重生成自研 golden，与
# 冻结入库的 golden 逐字节比对；再用冻结 golden 驱动报告。全程零第三方依赖、
# 不启动 JVM、不读取任何上游产物，气隙可跑。
#
# 事实来源：kernel/mutation/testdata/pit/golden/mutation-semantics.golden 是本仓
# 引擎与语义见证的历史观测快照，非 PIT 官方产物；竞品对拍分数维持 NOT_MEASURED。
# 本脚本只做编排与如实报告，不修改任何判定口径、不放宽门禁阈值。
# =============================================================================
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'

cd "$(dirname "$0")/.." || exit 2
export PATH="$PATH:$(go env GOPATH)/bin"

REPORT_JSON=".artifacts/pit-parity.json"
REPORT_MD="docs/dev/pit-parity-report.md"
GOLDEN="kernel/mutation/testdata/pit/golden/mutation-semantics.golden"
MIN_RATE="100.00%"

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { echo -e "${RED}[FAIL] $1${NC}" >&2; exit 1; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

usage() {
  cat <<'EOF'
用法: scripts/pit-parity.sh

离线跑自研 golden 语义回归：本仓引擎重生成 golden → 与冻结 golden 逐字节比对 →
生成 JSON/Markdown 报告。无外部参数，零第三方依赖。

  -h, --help   显示本帮助
EOF
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    -h|--help) usage; exit 0 ;;
    *) usage; echo -e "${RED}[FAIL] 未知参数: $1${NC}" >&2; exit 2 ;;
  esac
done

banner "Abate 自研 golden 语义回归（非 PIT 官方对拍）"

command -v go >/dev/null 2>&1 || fail "未找到 go 工具链"
[ -f "$GOLDEN" ] || fail "缺少冻结自研 golden: $GOLDEN"

# ---------------------------------------------------------------- 逐字节回归
banner "重生成自研 golden 并与冻结 golden 逐字节比对"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT
FRESH="$WORK/mutation-semantics.golden"
GOFLAGS="${GOFLAGS:-}" GOPROXY=off \
  go run kernel/mutation/pitparity/gen_golden.go -out "$FRESH"

if cmp -s "$FRESH" "$GOLDEN"; then
  RATE="100.00%"
  pass "自研 golden 逐字节一致（一致率 $RATE，28 判例）"
else
  # 真实计算行级一致率并如实报出，绝不伪造 100%。
  TOTAL="$(wc -l < "$GOLDEN" | tr -d ' ')"
  MATCHED="$(awk 'NR==FNR{line[FNR]=$0; next} $0==line[FNR]{m++} END{print m+0}' "$FRESH" "$GOLDEN")"
  RATE="$(awk -v m="$MATCHED" -v t="$TOTAL" 'BEGIN{printf "%.2f%%", m*100/t}')"
  echo -e "${RED}[FAIL] 自研 golden 重生成结果与冻结 golden 不一致（真实一致率 $RATE，$MATCHED/$TOTAL 行）${NC}" >&2
  diff -u "$GOLDEN" "$FRESH" >&2 || true
  exit 1
fi

# ---------------------------------------------------------------- 报告生成
banner "生成自研 golden 回归报告（对比本仓引擎输出 vs 冻结 golden）"
GOFLAGS="${GOFLAGS:-}" GOPROXY=off \
  go run ./kernel/mutation/pitparity/cmd/pitparity-report \
  -golden "$GOLDEN" -json "$REPORT_JSON" -md "$REPORT_MD" -min-rate "$MIN_RATE"
pass "对拍报告已生成（竞品对拍分数 NOT_MEASURED）"

echo ""
echo "  JSON 报告:   $REPORT_JSON"
echo "  Markdown:    $REPORT_MD"
banner "GOLDEN REGRESSION: PASSED（一致率 $RATE，自研 golden 非 PIT 官方对拍）"
