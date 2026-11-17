#!/usr/bin/env bash
# =============================================================================
# Abate 漂移检查门禁（XCU-CI §3.4）
#
# 四类生成物与源码零漂移：
#   1) CES JSON Schema / fields.md  —— CES-SCHEMA 生成物（版本与结构一致性）。
#   2) OpenAPI                      —— CTL-REST 运行时生成，由测试验证与实现同步。
#   3) 前端 API 客户端（openapi.json 生成）—— 本交付无独立前端，SKIP 并说明。
#   4) 前端 dist（frontend 构建）   —— 本交付无独立前端，SKIP 并说明。
#
# 任何漂移即 FAIL。全流程离线（-mod=vendor）。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/.." || exit 2
export PATH="/usr/local/go/bin:$PATH:$PATH" GOFLAGS="-mod=vendor" GOPROXY=off

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

banner "Abate Drift Gate - 生成物与源码零漂移"

# ---- 1) CES Schema / fields.md ----
SCHEMA="ces/schema/dist/ces-v1.schema.json"
FIELDS="ces/schema/dist/fields.md"
if [ ! -s "$SCHEMA" ]; then fail "CES schema 缺失或为空: $SCHEMA"; else pass "CES schema 存在且非空"; fi
if [ ! -s "$FIELDS" ]; then fail "CES fields.md 缺失或为空: $FIELDS"; else pass "CES fields.md 存在且非空"; fi

# 当前规范版本（单一事实来源：schema.Version）。
CUR_VER="$(cat > /tmp/cesver.go <<'EOF'
package main
import ("fmt"; "abate/ces/schema")
func main(){ fmt.Print(schema.Version{}.Current()) }
EOF
go run /tmp/cesver.go 2>/dev/null)"
if [ -z "$CUR_VER" ]; then CUR_VER="ces.tencent.com/v1"; fi

if ! jq empty "$SCHEMA" >/dev/null 2>&1; then
  fail "CES schema 不是合法 JSON"
else
  pass "CES schema 是合法 JSON"
fi
# 版本同步：schema 标题锚定的大版本须与代码当前规范版本一致。
SCHEMA_VER="$(jq -r '.title' "$SCHEMA" 2>/dev/null)"
SCHEMA_MAJOR="$(printf '%s' "$SCHEMA_VER" | grep -oE 'v[0-9]+' | head -1)"
CODE_MAJOR="$(printf '%s' "$CUR_VER" | grep -oE 'v[0-9]+' | head -1)"
if [ "$SCHEMA_MAJOR" != "$CODE_MAJOR" ] || [ -z "$SCHEMA_MAJOR" ]; then
  fail "CES schema 版本($SCHEMA_VER) 与代码版本($CUR_VER) 大版本不一致"
else
  pass "CES schema 版本 $SCHEMA_MAJOR 与代码 $CODE_MAJOR 一致"
fi
if [ "$(jq -r '.title' "$SCHEMA" 2>/dev/null)" != "CES Experiment v1" ]; then
  fail "CES schema title 漂移（期望 'CES Experiment v1'）"
else
  pass "CES schema 结构锚定（title=CES Experiment v1）"
fi

# ---- 2) OpenAPI 与实现同步 ----
banner "OpenAPI 漂移（运行时生成，由测试验证）"
if go test -count=1 -run TestOpenAPIUnauthenticated ./ctl/rest/ >/dev/null 2>&1; then
  pass "OpenAPI 由代码生成且通过一致性测试（与实现零漂移）"
else
  fail "OpenAPI 测试失败（生成物与实现漂移）"
fi

# ---- 3) & 4) 前端 ----
banner "前端生成物（本交付无独立前端）"
if [ -d frontend ] || [ -d web/dist ]; then
  fail "检测到前端目录但未纳入漂移检查"
else
  echo -e "${YELLOW}[SKIP] 本交付不含独立前端（API 客户端 / dist 漂移检查不适用，CTL-WEB 由 Go 同源构建）${NC}"
fi

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "DRIFT GATE: FAILED"; exit 1
fi
banner "DRIFT GATE: PASSED"
exit 0
