#!/usr/bin/env bash
# =============================================================================
# Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
# Abate is dual-licensed under AGPLv3 and a commercial license.
# For commercial licensing, contact xiaolibuluo@163.com.
# Source: https://github.com/xiaolibuluo/Abate
# =============================================================================
# 通用文件哈希清单工具（完整性层）。
# 用法: bash scripts/hash-check.sh generate [清单路径]
#       bash scripts/hash-check.sh verify   [清单路径]
# 默认清单为仓库根目录 SHA256SUMS，默认校验仓库根关键交付文件。
# generate 覆盖写入清单，verify 逐项校验并输出 PASS/FAIL，退出码 0/1/2。
# =============================================================================
set -uo pipefail

GREEN='\033[0;32m'; RED='\033[0;31m'; YELLOW='\033[1;33m'; NC='\033[0m'
MODE="${1:-verify}"
MANIFEST="${2:-SHA256SUMS}"

CRITICAL_FILES=(
  "README.md"
  "LICENSE"
  "Makefile"
  "main.go"
  "internal/watermark/watermark.go"
  "internal/watermark/buildinfo.go"
  "internal/watermark/watermark_test.go"
  "internal/watermark/buildinfo_test.go"
  "cmd/watermark/main.go"
  "scripts/hash-check.sh"
  "scripts/watermark-verify.sh"
)

cd "$(dirname "$0")/.." || exit 2

usage() {
  echo "Usage: bash scripts/hash-check.sh [generate|verify] [manifest]" >&2
}

echo -e "${YELLOW}== HASH CHECK (${MODE}) ==${NC}"

case "$MODE" in
  generate)
    tmp="$(mktemp)" || exit 2
    for f in "${CRITICAL_FILES[@]}"; do
      if [ ! -f "$f" ]; then
        echo -e "${RED}[FAIL] 待登记文件不存在: ${f}${NC}" >&2
        rm -f "$tmp"
        exit 1
      fi
      if ! sha256sum "$f" >> "$tmp"; then
        echo -e "${RED}[FAIL] 计算哈希失败: ${f}${NC}" >&2
        rm -f "$tmp"
        exit 2
      fi
    done
    if ! mv "$tmp" "$MANIFEST"; then
      echo -e "${RED}[FAIL] 写入清单失败: ${MANIFEST}${NC}" >&2
      rm -f "$tmp"
      exit 2
    fi
    echo -e "${GREEN}[PASS] 已生成 ${MANIFEST}（${#CRITICAL_FILES[@]} 个关键文件）${NC}"
    exit 0
    ;;
  verify)
    if [ ! -f "$MANIFEST" ]; then
      echo -e "${RED}[FAIL] 清单不存在: ${MANIFEST}（先运行 generate）${NC}" >&2
      exit 1
    fi
    listed="$(awk '{print $2}' "$MANIFEST" | LC_ALL=C sort)"
    for f in "${CRITICAL_FILES[@]}"; do
      if ! printf '%s\n' "$listed" | grep -Fxq "$f"; then
        echo -e "${RED}[FAIL] 清单缺少关键文件: ${f}${NC}" >&2
        exit 1
      fi
    done
    if sha256sum -c --strict --quiet "$MANIFEST"; then
      echo -e "${GREEN}[PASS] ${MANIFEST} 校验通过（${#CRITICAL_FILES[@]} 个关键文件哈希一致）${NC}"
      exit 0
    fi
    echo -e "${RED}[FAIL] ${MANIFEST} 校验失败${NC}" >&2
    exit 1
    ;;
  *)
    usage
    exit 2
    ;;
esac
