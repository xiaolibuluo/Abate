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
# mcp-demo.sh —— 演示 Abate MCP Server（JSON-RPC 2.0 over stdio）
# mcp-demo.sh —— Demo the Abate MCP Server (JSON-RPC 2.0 over stdio)
#
# 用法 Usage:
#   ./scripts/mcp-demo.sh
#   ABATE_BIN=/path/to/abate ./scripts/mcp-demo.sh
#
# 依赖 Dependencies: jq（可选 abate 二进制缺失时自动 go build）

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ABATE_BIN="${ABATE_BIN:-}"

for cmd in jq; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "mcp-demo: 缺少依赖 missing dependency: $cmd" >&2
    exit 1
  fi
done

if [ -z "$ABATE_BIN" ]; then
  if [ -x "$ROOT/abate" ]; then
    ABATE_BIN="$ROOT/abate"
  elif command -v abate >/dev/null 2>&1; then
    ABATE_BIN="$(command -v abate)"
  else
    echo "mcp-demo: 未找到 abate 二进制，尝试 go build …" >&2
    if ! command -v go >/dev/null 2>&1 && [ -x /usr/local/go/bin/go ]; then
      export PATH="$PATH:/usr/local/go/bin"
    fi
    ABATE_BIN="$(mktemp -u /tmp/abate-mcp-demo.XXXXXX)"
    (cd "$ROOT" && go build -o "$ABATE_BIN" ./cmd/abate)
    trap 'rm -f "$ABATE_BIN"' EXIT
  fi
fi

echo "== abate 二进制 binary: $ABATE_BIN =="
echo ""

send() {
  printf '%s\n' "$1" | "$ABATE_BIN" mcp
}

REQ_INIT='{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"mcp-demo","version":"0.1.0"}}}'
REQ_LIST='{"jsonrpc":"2.0","id":2,"method":"tools/list"}'
REQ_FAULTS='{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"list_fault_classes","arguments":{}}}'
REQ_MUT='{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"generate_mutants","arguments":{"filename":"calc.go","source":"package calc\nfunc Add(a, b int) int { return a + b }","seed":42}}}'

echo "== 1. initialize =="
send "$REQ_INIT" | jq .
echo ""
echo "== 2. tools/list =="
send "$REQ_LIST" | jq -r '.result.tools[] | "\(.name): \(.description)"'
echo ""
echo "== 3. tools/call list_fault_classes =="
send "$REQ_FAULTS" | jq -r '.result.content[0].text' | jq .
echo ""
echo "== 4. tools/call generate_mutants (seed=42) =="
send "$REQ_MUT" | jq -r '.result.content[0].text' | jq .
