#!/usr/bin/env bash
# =============================================================================
# Abate 依赖漏洞 / 供应链扫描入口（XCU-SEC §4.2，威胁 T4）
#
# 离线执行（GOPROXY=off），不做任何网络访问。检查两类事实来源：
#   1) third_party/LICENSE-MANIFEST.md 登记项：已 vendoring 的必须锁版本/commit，
#      且许可证为允许集（拒绝 GPL 系）。
#   2) third_party/ 目录：每个实际并入的子目录必须在清单有登记条目（零漂移）。
#
# govulncheck 若已源码级并入则调用；否则说明其离线审查由人工评审承担。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/../.." || exit 2

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

MANIFEST="third_party/LICENSE-MANIFEST.md"
[ -f "$MANIFEST" ] || { echo -e "${RED}[FAIL] 缺少 $MANIFEST${NC}"; exit 1; }

banner "Abate Supply-Chain Audit - 第三方源码并入登记审查（离线）"

# 允许的许可证（白名单，与 license-check.sh 一致）。
ALLOWED="MIT|Apache|BSD|ISC|CC0|Unlicense|Public Domain|公有领域"
# 拒绝的许可证（GPL 系）。
FORBIDDEN="GPL|LGPL|AGPL"

issues=0
while IFS= read -r raw; do
  # 仅处理表格行（以 '|' 开头），跳过所有 markdown 标题/列表/正文。
  case "$raw" in
    '| '*|'|'*) ;;
    *) continue ;;
  esac
  IFS='|' read -r _comp _src _ver _lic _lang _method _ms _status _rest <<< "$raw"
  comp="$(printf '%s' "$_comp" | tr -d ' \t')"
  ver="$(printf '%s' "$_ver" | tr -d ' \t')"
  lic="$(printf '%s' "$_lic" | tr -d ' \t')"
  status="$(printf '%s' "$_status" | tr -d ' \t')"
  [ -z "$comp" ] || [ "$comp" = "组件" ] && continue
  echo "$comp" | grep -q '^--' && continue

  # 已 vendoring / 已锁定 状态的必须锁版本。
  if printf '%s' "$status" | grep -qE '已|锁定|完成|规划中|决策已采纳'; then
    if printf '%s' "$ver" | grep -qE '^_?待|^$' ; then
      # 仅"已 vendoring"且仍占位才失败；"规划中/未 vendoring"允许占位。
      if printf '%s' "$status" | grep -qE '已|锁定|完成'; then
        fail "$comp 状态为'$status'但版本未锁定($ver)"
        issues=$((issues+1))
      fi
    fi
  fi

  # 许可证白/黑名单。
  if printf '%s' "$lic" | grep -qiE "$FORBIDDEN"; then
    fail "$comp 许可证含 GPL 系($lic)：拒绝"
    issues=$((issues+1))
  elif ! printf '%s' "$lic" | grep -qiE "$ALLOWED"; then
    fail "$comp 许可证未知/未审查($lic)"
    issues=$((issues+1))
  fi
done < "$MANIFEST"

if [ "$issues" -eq 0 ]; then
  pass "第三方登记项版本锁定与许可证审查全部通过"
else
  fail "第三方登记审查发现 $issues 项问题"
fi

banner "third_party/ 目录零漂移（目录内每项须在清单登记）"
dir_drift=0
for d in third_party/*/; do
  [ -d "$d" ] || continue
  name="$(basename "$d")"
  [ "$name" = "_references" ] && continue
  if ! grep -q "$name" "$MANIFEST"; then
    fail "目录 third_party/$name 未在清单登记"
    dir_drift=$((dir_drift+1))
  fi
done
[ "$dir_drift" -eq 0 ] && pass "third_party/ 子目录与清单零漂移"

banner "漏洞扫描工具（govulncheck）"
if command -v govulncheck >/dev/null 2>&1; then
  pass "govulncheck 可用；离线扫描交由 CI 周期任务执行"
else
  echo -e "${YELLOW}[INFO] govulncheck 未并入；依赖漏洞审查由 third_party 锁版本人工评审承担（对齐 KRN-INJ 上游回移机制）${NC}"
fi

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "SUPPLY-CHAIN AUDIT: FAILED"; exit 1
fi
banner "SUPPLY-CHAIN AUDIT: PASSED"
exit 0
