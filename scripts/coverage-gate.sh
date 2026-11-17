#!/usr/bin/env bash
# =============================================================================
# Abate 覆盖率门禁（XCU-TST / 用户规则：混沌测试 & 变异测试必须 100% 覆盖）
#
# 策略（事实导向，不伪造）：
#   - 对每个「已有测试文件」的包，跑 go test -coverprofile 取覆盖率。
#   - 内核包（kernel/ sim/ ces/）语句覆盖率门槛 ≥ 95%；其余 ≥ 85%（架构 DoD 总则）。
#   - 内核包一旦有测试，必须同时存在 *_chaos_test.go 与 *_mutation_test.go
#     （混沌测试 / 变异测试必须覆盖），缺失即 FAIL。
#   - 尚无测试文件的包跳过（不阻断采购期空仓库），但会列出提示。
# 报告真实覆盖率数字，绝不虚报 100%。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0; REPORT=""

cd "$(dirname "$0")/.." || exit 2
export PATH="$PATH:$(go env GOPATH)/bin"

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

KERNEL_THRESHOLD=95
ECO_THRESHOLD=85

# 收集所有含 .go 的包（排除 vendor/third_party/cmd 下的 main 包不强制覆盖率）。
PKGS="$(go list ./... 2>/dev/null | grep -v '/vendor/' | grep -v '/third_party/' || true)"

banner "Abate Coverage Gate - 覆盖率 + 混沌/变异测试覆盖"

if [ -z "$PKGS" ]; then
  pass "无 Go 包（采购期空仓库），覆盖率门禁跳过"
  exit 0
fi

for pkg in $PKGS; do
  # 该包目录下的测试文件
  dir="${pkg#abate/}"
  [ "$dir" = "$pkg" ] && dir="."
  dir="./$dir"
  tests="$(find "$dir" -maxdepth 1 -name '*_test.go' 2>/dev/null | sort)"

  if [ -z "$tests" ]; then
    # 精准判定：main 包与生成代码包天然不可被单元测试覆盖，明确排除且不阻断；
    # 其余库包必须自带测试，缺失即 FAIL（不再以「尚无测试」笼统豁免）。
    pname="$(go list -f '{{.Name}}' "$pkg" 2>/dev/null || true)"
    gen=0
    case "$pkg" in
      */gen|*/gen/*) gen=1 ;;
    esac
    if [ "$pname" = "main" ] || [ "$gen" -eq 1 ]; then
      echo "  [N/A] $pkg 为 main/生成包，覆盖率不适用"
      continue
    fi
    fail "库包缺少测试文件（必须覆盖，禁止豁免）: $pkg"
    continue
  fi

  is_kernel=0
  case "$pkg" in
    abate/kernel/*|abate/sim/*|abate/ces/*) is_kernel=1 ;;
  esac

  # 混沌 / 变异测试文件存在性（仅内核包强制）
  if [ "$is_kernel" -eq 1 ]; then
    echo "$tests" | grep -q '_chaos_test.go' || fail "内核包缺少混沌测试文件(*_chaos_test.go): $pkg"
    echo "$tests" | grep -q '_mutation_test.go' || fail "内核包缺少变异测试文件(*_mutation_test.go): $pkg"
  fi

  # 覆盖率
  prof="$(mktemp)"
  out="$(go test -coverprofile="$prof" -count=1 "$pkg" 2>&1)"
  cov="$(printf '%s' "$out" | grep -oE 'coverage: [0-9.]+%' | grep -oE '[0-9.]+' | head -1)"
  rm -f "$prof"
  # 零语句纯测试包（仅 *_test.go + 无生产语句的 doc.go）：go test -cover 会
  # 报告 "coverage: [no statements]"。此类包没有可覆盖的生产语句，覆盖率
  # 概念不适用，按 N/A 跳过而非误判为「测试失败」。真实的测试失败（输出含
  # FAIL / build failed）不含该标记，仍走下方 fail 分支，门槛不因此降低。
  if [ -z "$cov" ] && printf '%s' "$out" | grep -q 'no statements' &&
     ! printf '%s' "$out" | grep -q 'FAIL'; then
    echo "  [N/A] $pkg 无生产语句（纯测试包），覆盖率不适用"
    continue
  fi
  if [ -z "$cov" ]; then
    fail "无法取得覆盖率（测试可能失败）: $pkg"; continue
  fi
  thr=$ECO_THRESHOLD; [ "$is_kernel" -eq 1 ] && thr=$KERNEL_THRESHOLD
  # 用 awk 比较浮点
  ok="$(awk -v c="$cov" -v t="$thr" 'BEGIN { if (c + 0 >= t + 0) print 1; else print 0 }')"
  if [ "$ok" -eq 1 ]; then
    echo "  [OK] $pkg coverage=${cov}% (门槛 ${thr}%)"
  else
    fail "$pkg 覆盖率 ${cov}% 低于门槛 ${thr}%"
  fi
done

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "COVERAGE GATE: FAILED"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
banner "COVERAGE GATE: PASSED（仅统计已有测试的包；真实数字见上方）"
exit 0
