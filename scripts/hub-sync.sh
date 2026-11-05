#!/usr/bin/env bash
# Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
# Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
# Copyright (C) 2026 小黎的部落（www.qingdujian.com）
#
# 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
# 小黎的部落书面授权，详见 COMMERCIAL.md。
# This file is published under the GNU Affero General Public License v3 (AGPLv3).
# Commercial use requires prior written authorization from 小黎的部落.
# See COMMERCIAL.md for details.
#
# hub-sync.sh —— 同步 ChaosHub 子编码注册表（CES 附录 B）
# hub-sync.sh —— Sync the ChaosHub sub-code registry (CES Appendix B)
#
# 用法 Usage:
#   ./scripts/hub-sync.sh -u https://hub.example.com [-k API_KEY]
#   ABATE_HUB_URL=https://hub.example.com ./scripts/hub-sync.sh
#
# 依赖 Dependencies: curl, jq

set -euo pipefail

HUB_URL="${ABATE_HUB_URL:-}"
API_KEY="${ABATE_HUB_API_KEY:-}"

usage() {
  cat >&2 <<'USAGE'
用法 Usage: hub-sync.sh -u <ChaosHub URL> [-k <API Key>]
  -u, --url       ChaosHub 数据源 URL（或设 ABATE_HUB_URL）
  -k, --api-key   API Key（或设 ABATE_HUB_API_KEY；不会写入任何日志）
  -h, --help      显示本帮助 Show this help
USAGE
}

while [ $# -gt 0 ]; do
  case "$1" in
    -u|--url)     HUB_URL="${2:?missing url value}"; shift 2 ;;
    -k|--api-key) API_KEY="${2:?missing api-key value}"; shift 2 ;;
    -h|--help)    usage; exit 0 ;;
    *) echo "hub-sync: unknown argument: $1" >&2; usage; exit 2 ;;
  esac
done

if [ -z "$HUB_URL" ]; then
  echo "hub-sync: 未指定 ChaosHub URL（用 -u 或 ABATE_HUB_URL）" >&2
  echo "hub-sync: ChaosHub URL required (use -u or ABATE_HUB_URL)" >&2
  usage
  exit 2
fi

for cmd in curl jq; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "hub-sync: 缺少依赖 missing dependency: $cmd" >&2
    exit 1
  fi
done

CURL_ARGS=(-fsSL --max-time 30)
if [ -n "$API_KEY" ]; then
  CURL_ARGS+=(-H "Authorization: Bearer ${API_KEY}")
fi

RESPONSE="$(curl "${CURL_ARGS[@]}" "${HUB_URL%/}/api/v1/hub/subcodes")"

VERSION="$(printf '%s' "$RESPONSE" | jq -r '.version')"
COUNT="$(printf '%s' "$RESPONSE" | jq -r '.count')"

echo "ChaosHub 子编码注册表 Sub-Code Registry"
echo "  数据源 Source : $HUB_URL"
echo "  版本 Version  : $VERSION"
echo "  数量 Count    : $COUNT"
echo ""
echo "CODE      NAME  PARAMS"
printf '%s' "$RESPONSE" | jq -r '.codes[] | "\(.code)  \(.name)  [\(.params)]"'
