#!/usr/bin/env bash
# =============================================================================
# Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
# Abate is dual-licensed under AGPLv3 and a commercial license.
# For commercial licensing, contact xiaolibuluo@163.com.
# Source: https://github.com/xiaolibuluo/Abate
# =============================================================================
# 三层数字代码水印校验（源码常量 / ldflags 构建注入 / SHA256 完整性清单）。
# 用法: bash scripts/watermark-verify.sh
# 输出逐层 PASS/FAIL，任一层失败即退出 1；全部通过退出 0。
# =============================================================================
set -uo pipefail

GREEN='\033[0;32m'; RED='\033[0;31m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
SRC="internal/watermark/watermark.go"
CANON_ANCHOR='xiaolibuluo/Abate © 小黎的部落'
FP_ANCHOR='e92fa6a0a80b33498123d5545e07e5b2c8b1ce30f220e3a933644bd8118b81bc'
OWNER_ANCHOR='小黎的部落'
SOURCE_ANCHOR='https://github.com/xiaolibuluo/Abate'

cd "$(dirname "$0")/.." || exit 2

fail() { FAILED=1; echo -e "${RED}[FAIL] $*${NC}"; }
pass() { echo -e "${GREEN}[PASS] $*${NC}"; }
step() { echo -e "${YELLOW}>>> $*${NC}"; }

# extract_const 从 Go 源文件读取具名字符串常量。
extract_const() {
  sed -n "s/^[[:space:]]*$1[[:space:]]*=[[:space:]]*\"\([^\"]*\)\".*/\1/p" "$SRC" | head -n 1
}

step "L1 源码层：解析常量并与独立锚点比对"
if [ ! -f "$SRC" ]; then
  fail "L1: 水印源文件缺失 ${SRC}"
else
  owner="$(extract_const OwnerName)"
  source_url="$(extract_const SourceURL)"
  canonical="$(extract_const Canonical)"
  fingerprint="$(extract_const FingerprintSHA256)"
  for pair in "OwnerName:${owner}" "SourceURL:${source_url}" \
              "Canonical:${canonical}" "FingerprintSHA256:${fingerprint}"; do
    if [ -z "${pair#*:}" ]; then
      fail "L1: 无法解析常量 ${pair%%:*}"
    fi
  done
  if [ "$owner" = "$OWNER_ANCHOR" ]; then
    pass "L1: OwnerName=${owner}"
  else
    fail "L1: OwnerName 不符 expect=${OWNER_ANCHOR} actual=${owner}"
  fi
  if [ "$source_url" = "$SOURCE_ANCHOR" ]; then
    pass "L1: SourceURL=${source_url}"
  else
    fail "L1: SourceURL 不符 expect=${SOURCE_ANCHOR} actual=${source_url}"
  fi
  if [ "$canonical" = "$CANON_ANCHOR" ]; then
    pass "L1: Canonical=${canonical}"
  else
    fail "L1: Canonical 不符 expect=${CANON_ANCHOR} actual=${canonical}"
  fi
  recomputed="$(printf '%s' "$canonical" | sha256sum | cut -d' ' -f1)"
  if [ "$recomputed" = "$FP_ANCHOR" ] && [ "$fingerprint" = "$FP_ANCHOR" ]; then
    pass "L1: 规范串重算 sha256=${recomputed}，与常量一致"
  else
    fail "L1: 指纹不一致 anchor=${FP_ANCHOR} constant=${fingerprint} recomputed=${recomputed}"
  fi
fi

if bash scripts/header-check.sh >/dev/null 2>&1; then
  pass "L1: 全仓 Go 版权头水印齐全"
else
  fail "L1: header-check 未通过"
fi

step "L2 构建层：ldflags 注入与专用入口实跑"
if grep -Fq -- '-X $(WATERMARK_PKG).BuildTime=' Makefile &&
   grep -Fq -- '-X $(WATERMARK_PKG).GitCommit=' Makefile; then
  pass "L2: Makefile 含 BuildTime / GitCommit 的 -X 注入"
else
  fail "L2: Makefile 缺少 -X 构建注入"
fi
pkg="$(sed -n 's/^WATERMARK_PKG[[:space:]]*:=[[:space:]]*//p' Makefile | head -n 1)"
if [ -z "$pkg" ]; then
  fail "L2: 无法从 Makefile 解析 WATERMARK_PKG"
else
  probe="$(go run -ldflags "-X ${pkg}.BuildTime=verify-time -X ${pkg}.GitCommit=verify-commit" ./cmd/watermark 2>&1)"
  probe_rc=$?
  if [ "$probe_rc" -ne 0 ]; then
    fail "L2: 专用入口退出码 ${probe_rc}: ${probe}"
  else
    printf '%s\n' "$probe"
    for needle in "owner: ${OWNER_ANCHOR}" "fingerprint: ${FP_ANCHOR}" \
                  "build-time: verify-time" "git-commit: verify-commit" \
                  "verified: PASS"; do
      if printf '%s\n' "$probe" | grep -Fq -- "$needle"; then
        pass "L2: 入口输出包含 ${needle}"
      else
        fail "L2: 入口输出缺少 ${needle}"
      fi
    done
  fi
fi

step "L3 完整性层：SHA256 基线清单校验"
if bash scripts/hash-check.sh verify WATERMARK.sha256; then
  pass "L3: WATERMARK.sha256 清单校验通过"
else
  fail "L3: WATERMARK.sha256 清单校验失败"
fi

echo ""
if [ "$FAILED" -ne 0 ]; then
  echo -e "${RED}=== WATERMARK VERIFY: FAIL ===${NC}"
  exit 1
fi
echo -e "${GREEN}=== WATERMARK VERIFY: PASS ===${NC}"
exit 0
