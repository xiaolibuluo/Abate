#!/usr/bin/env bash
# =============================================================================
# Abate 供应链许可证门禁
# 铁律：禁止接入 GPL 系许可证的开源（项目自身双授权 AGPLv3+商业 不在此限）。
# 任何导入依赖若被判定为 GPL/LGPL/AGPL -> 直接 FAIL。
# 同时校验 third_party/LICENSE-MANIFEST.md 登记的组件许可证与 go.mod 一致。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

FAILED=0
REPORT=""

cd "$(dirname "$0")/.." || exit 2
export PATH="$PATH:$(go env GOPATH)/bin"

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

# 已知许可证映射（模块路径关键字 -> 许可证）。用于无法从 module cache 读取
# LICENSE 文本时的可信回退，全部为已审查的非 GPL 宽松许可。
declare -A KNOWN_LICENSE=(
  ["gopkg.in/yaml.v3"]="MIT"
  ["gvisor.dev/gvisor"]="Apache-2.0"
)

# GPL 系关键词（大小写不敏感）。命中即违规。
GPL_PATTERNS=("GNU GENERAL PUBLIC LICENSE" "GNU LESSER GENERAL PUBLIC" "GNU AFFERO" "GNU LIBRARY GENERAL PUBLIC")

# 允许的非 GPL 关键词（命中其一即视为合规，用于辅助判定）。
# BSD-2/3-Clause 标准法律文本不含 "BSD" 字面词，靠其特征短语判定。
OK_PATTERNS=("MIT License" "Apache License" "BSD" "ISC License" "CC0" "Unlicense" "public domain" "Public Domain" "Redistribution and use in source and binary forms")

classify_license() {
  # $1 = module path@version ; 返回判定字符串：ALLOW / GPL / UNKNOWN
  local mod="$1"
  local base="${mod%@*}"
  local cache
  cache="$(go env GOMODCACHE)"

  # 已知映射优先
  for key in "${!KNOWN_LICENSE[@]}"; do
    if [[ "$base" == "$key"* ]]; then
      echo "${KNOWN_LICENSE[$key]}"
      return
    fi
  done

  # 尝试从 module cache 读取 LICENSE 文本
  local dir="$cache/${mod}"
  local lic=""
  if [ -d "$dir" ]; then
    lic="$(cat "$dir"/LICENSE* 2>/dev/null || true)"
  fi

  if [ -z "$lic" ]; then
    echo "UNKNOWN"
    return
  fi

  local upper
  upper="$(printf '%s' "$lic" | tr '[:lower:]' '[:upper:]')"
  for p in "${GPL_PATTERNS[@]}"; do
    if printf '%s' "$upper" | grep -qi -- "$p"; then
      echo "GPL"
      return
    fi
  done

  for p in "${OK_PATTERNS[@]}"; do
    if printf '%s' "$upper" | grep -qi -- "$p"; then
      echo "ALLOW"
      return
    fi
  done

  echo "UNKNOWN"
}

banner "Abate License Gate - 扫描导入依赖许可证"

# vendor 模式下 go list -m all 只返回主模块；vendor/modules.txt 才是实际
# 接入源码的权威清单，优先从它读取，确保检查反映真实供应链。
MODS=""
if [ -f vendor/modules.txt ]; then
  MODS="$(grep -E '^# [^ ]+ v' vendor/modules.txt | sed 's/^# //' | tr ' ' '@' || true)"
fi
if [ -z "$MODS" ]; then
  MODS="$(go list -m all 2>/dev/null | tail -n +2 | tr ' ' '@' || true)"
fi
if [ -z "$MODS" ]; then
  pass "无导入依赖（仅主模块）"
else
  VIOLATIONS=0
  while IFS= read -r mod; do
    [ -z "$mod" ] && continue
    verdict="$(classify_license "$mod")"
    case "$verdict" in
      GPL)
        fail "GPL 系许可证命中（禁止接入）: $mod"
        VIOLATIONS=1
        ;;
      UNKNOWN)
        fail "无法判定许可证（默认拒绝，需人工评审并登记到 third_party/LICENSE-MANIFEST.md）: $mod"
        VIOLATIONS=1
        ;;
      *)
        echo "  [OK] $mod -> $verdict"
        ;;
    esac
  done <<< "$MODS"
  [ "$VIOLATIONS" -eq 0 ] && pass "全部导入依赖许可证合规（无 GPL）"
fi

# 零依赖断言：彻底自研后，产品编译依赖必须为 0。
# 用 -mod=mod 绕过 vendor 目录，直接读取 go.mod 的真实 require 面；vendor/ 在 T04
# 才会物理删除，在此之前不能以 vendor/modules.txt 为空来假装零依赖。
# 判据：go list -m all 输出只应含主模块自身（abate）；出现任何第三方模块即 FAIL。
banner "零依赖断言：go list -m all 仅允许主模块"
DEPS="$(GOFLAGS=-mod=mod go list -m all 2>/dev/null | tail -n +2 | tr ' ' '@' || true)"
if [ -z "$DEPS" ]; then
  pass "彻底自研：go.mod 第三方依赖 = 0（仅主模块 abate）"
else
  echo -e "${RED}  当前残留模块:${NC}"
  printf '%s\n' "$DEPS" | sed 's/^/    /'
  fail "零依赖断言未通过：仍有 $(printf '%s\n' "$DEPS" | grep -c .) 个第三方模块（T03 清 wazero / T04 清 yaml.v3 后归零）"
fi

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "LICENSE GATE: FAILED"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
banner "LICENSE GATE: PASSED - 无 GPL 系依赖"
exit 0
