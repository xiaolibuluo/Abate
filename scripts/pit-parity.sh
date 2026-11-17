#!/usr/bin/env bash
# =============================================================================
# Abate PIT 对拍编排脚本（M2 验收：四类变异判定与分数口径与 PIT 参考实现对拍一致，
# 对拍报告入库；PIT 对拍一致率 100%，差异逐条说明）
#
# 两种模式：
#   默认（离线）  : 用入库的官方 PIT 产物 kernel/mutation/testdata/pit/pit-mutations.json
#                    跑 Go 侧逐点对拍，生成 docs/dev/pit-parity-report.md 与
#                    .artifacts/pit-parity.json；零外部依赖，气隙可跑。
#   --rerun-pit   : 真实重跑 PITest 1.19.6（需要 java 与 mvn，且 Maven 镜像可达），
#                    把新产物写回 testdata 后再对拍；依赖缺失时如实报错并提示用默认模式，
#                    绝不伪造产物。
#
# 事实来源：PIT 参考产物由 PITest 1.19.6 官方实跑生成（28 条变异：Killed 23、
# SURVIVED 3、NO_COVERAGE 2），本脚本只做编排与复跑，不修改任何判定口径。
# =============================================================================
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'

cd "$(dirname "$0")/.." || exit 2
export PATH="$PATH:$(go env GOPATH)/bin"

REPORT_JSON=".artifacts/pit-parity.json"
REPORT_MD="docs/dev/pit-parity-report.md"
PIT_JSON="kernel/mutation/testdata/pit/pit-mutations.json"
PIT_XML="kernel/mutation/testdata/pit/pit-mutations.xml"
PIT_BASELINE="kernel/mutation/testdata/pit/baseline"
MIN_RATE="100.00%"
RERUN=0
MAVEN_SETTINGS="${MAVEN_SETTINGS:-}"

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { echo -e "${RED}[FAIL] $1${NC}" >&2; exit 1; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

usage() {
  cat <<'EOF'
用法: scripts/pit-parity.sh [--rerun-pit] [--settings <maven-settings.xml>]

  --rerun-pit            真实重跑 PITest 1.19.6 并把产物写回 testdata
  --settings <path>      Maven settings.xml（--rerun-pit 模式使用）
  -h, --help             显示本帮助

默认模式完全离线，只使用入库的官方 PIT 产物。
EOF
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    --rerun-pit) RERUN=1; shift ;;
    --settings) MAVEN_SETTINGS="${2:-}"; shift 2 ;;
    -h|--help) usage; exit 0 ;;
    *) usage; echo -e "${RED}[FAIL] 未知参数: $1${NC}" >&2; exit 2 ;;
  esac
done

banner "Abate PIT Parity - KRN-MUT 变异核与 PITest 1.19.6 逐点对拍"

command -v go >/dev/null 2>&1 || fail "未找到 go 工具链"

if [ "$RERUN" -eq 1 ]; then
  # ---------------------------------------------------------------- 重跑模式
  banner "重跑 PITest 1.19.6"
  command -v java >/dev/null 2>&1 || {
    echo -e "${RED}[FAIL] --rerun-pit 需要 java（JDK 21），当前环境未安装${NC}" >&2
    echo "提示: 请改用默认模式 scripts/pit-parity.sh，它使用入库的官方 PIT 产物。" >&2
    exit 2
  }
  command -v mvn >/dev/null 2>&1 || {
    echo -e "${RED}[FAIL] --rerun-pit 需要 mvn，当前环境未安装${NC}" >&2
    echo "提示: 请改用默认模式 scripts/pit-parity.sh，它使用入库的官方 PIT 产物。" >&2
    exit 2
  }
  WORK="$(mktemp -d)"
  trap 'rm -rf "$WORK"' EXIT
  # 入库工件把源码目录改名为 src-main/src-test 以便归档，重跑时还原 Maven 标准布局。
  mkdir -p "$WORK/src/main" "$WORK/src/test"
  cp -R "$PIT_BASELINE/src-main/java" "$WORK/src/main/java"
  cp -R "$PIT_BASELINE/src-test/java" "$WORK/src/test/java"
  cp "$PIT_BASELINE/pom.xml" "$WORK/pom.xml"
  echo "重跑工作目录: $WORK"
  MVN_ARGS=(-q -B org.pitest:pitest-maven:1.19.6:mutationCoverage)
  if [ -n "$MAVEN_SETTINGS" ]; then
    MVN_ARGS=(-s "$MAVEN_SETTINGS" "${MVN_ARGS[@]}")
  fi
  ( cd "$WORK" && mvn "${MVN_ARGS[@]}" ) || {
    echo -e "${RED}[FAIL] PITest 重跑失败（Maven 仓库不可达或工程编译失败）${NC}" >&2
    echo "提示: 请改用默认模式 scripts/pit-parity.sh，它使用入库的官方 PIT 产物。" >&2
    exit 2
  }
  NEW_XML="$WORK/target/pit-reports/mutations.xml"
  [ -f "$NEW_XML" ] || {
    echo -e "${RED}[FAIL] PITest 未产出 mutations.xml${NC}" >&2
    exit 2
  }
  echo "新产物: $NEW_XML"
  cp "$NEW_XML" "$PIT_XML"
  GOFLAGS="${GOFLAGS:-} -mod=vendor" GOPROXY=off \
    go run ./kernel/mutation/pitparity/cmd/pitparity-report \
    -pit "$PIT_JSON" -xml "$PIT_XML" -json "$REPORT_JSON" -md "$REPORT_MD" \
    -min-rate "$MIN_RATE"
  pass "PITest 重跑产物已写回 testdata 并完成对拍"
else
  # ---------------------------------------------------------------- 默认模式
  banner "离线对拍（使用入库的官方 PIT 产物）"
  [ -f "$PIT_JSON" ] || fail "缺少入库产物: $PIT_JSON"
  GOFLAGS="${GOFLAGS:-} -mod=vendor" GOPROXY=off \
    go run ./kernel/mutation/pitparity/cmd/pitparity-report \
    -pit "$PIT_JSON" -json "$REPORT_JSON" -md "$REPORT_MD" -min-rate "$MIN_RATE"
  pass "对拍报告已生成"
fi

echo ""
echo "  JSON 报告:   $REPORT_JSON"
echo "  Markdown:    $REPORT_MD"
banner "PIT PARITY: PASSED（一致率门槛 $MIN_RATE，真实数字见上方输出）"
