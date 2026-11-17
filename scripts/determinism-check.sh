#!/usr/bin/env bash
# =============================================================================
# Abate 确定性门禁（XCU-CI §3.3）
#
# 核心包双跑（同种子）→ 哈希零差异。命中包清单：kernel/*、sim/*、ces/*。
# 这些包的 Chaos/Mutation/Record/Replay 测试内部以同种子双跑并断言逐字节一致
# （sim/record.Compare 对帧哈希链做零差异比对），本门禁执行它们并汇总。
#
# 全流程离线（-mod=vendor）。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/.." || exit 2
export PATH="/usr/local/go/bin:$PATH:$PATH" GOFLAGS="-mod=vendor" GOPROXY=off

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

banner "Abate Determinism Gate - 核心包同种子双跑哈希零差异"

# 确定性敏感的测试族：Chaos（混沌自测）/ Mutation（变异自测）/ Record（回放哈希链）
# / Replay / Determinism。这些测试断言同种子运行结果逐字节一致。
RUN_FILTER='Chaos|Mutation|Replay|Record|Determinism'

run_pkg() {
  local label="$1"; shift
  if go test -count=1 -run "$RUN_FILTER" "$@" >/dev/null 2>&1; then
    pass "$label 确定性测试全绿"
  else
    fail "$label 确定性测试存在漂移/失败"
  fi
}

banner "kernel/*（L0 内核：代数/调度/观测/验证/变异/注入）"
run_pkg "kernel" ./kernel/...

banner "sim/*（L3 模拟器：时钟/网络/文件系统/虚拟调度/录制回放/PRNG）"
run_pkg "sim/record" ./sim/record/...
run_pkg "sim/rand" ./sim/rand/...
run_pkg "sim" ./sim/clock/... ./sim/net/... ./sim/fs/... ./sim/vsched/...

banner "ces/*（L2 规范：解析/校验/模式）"
run_pkg "ces" ./ces/...

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "DETERMINISM GATE: FAILED"; exit 1
fi
banner "DETERMINISM GATE: PASSED"
exit 0
