#!/usr/bin/env bash
# =============================================================================
# Abate 分支保护规则集临时开关（Delivery Ops）
#
# 用途：配合 force-push.sh 使用。GitHub Ruleset（pull_request / non_fast_forward）
#       会拒绝直接/强制推送到 main。强制推送前临时 disable，推送后立即恢复 active。
#
# 凭据：从环境变量 GITHUB_TOKEN_FILE 指向的文件读取令牌（默认
#       /root/uploads/GitHub令牌.md 需按实际路径传），进程内读取，绝不回显或入仓。
#
# 用法：
#   GITHUB_TOKEN_FILE=/path/to/token.md scripts/ruleset-toggle.sh disable
#   GITHUB_TOKEN_FILE=/path/to/token.md scripts/ruleset-toggle.sh enable
#
# 注意：本脚本只切换 enforcement 状态，不改规则内容；操作会回读确认。
#       强制推送属高风险运维操作，请仅在 owner 授权下临时暂停保护并在推送后立即恢复。
# =============================================================================
set -euo pipefail

cd "$(dirname "$0")/.." || exit 2

ACTION="${1:-}"
case "$ACTION" in
  disable|enable) ;;
  *) echo "usage: ruleset-toggle.sh disable|enable" >&2; exit 2 ;;
esac

TOKEN_FILE="${GITHUB_TOKEN_FILE:-/root/uploads/GitHub令牌.md}"
[ -f "$TOKEN_FILE" ] || { echo "[FAIL] 令牌文件不存在: $TOKEN_FILE" >&2; exit 1; }
TOKEN="$(tr -d '[:space:]' < "$TOKEN_FILE")"
[ -n "$TOKEN" ] || { echo "[FAIL] 令牌为空" >&2; exit 1; }

RULESET_ID="${RULESET_ID:-27377536}"
API="https://api.github.com/repos/xiaolibuluo/Abate/rulesets/${RULESET_ID}"

echo "===== Ruleset 临时开关: ${ACTION} ====="

# 读取当前规则集完整对象（去掉只读字段后回写）。
curl -s --max-time 20 -H "Authorization: Bearer $TOKEN" -H "Accept: application/vnd.github+json" "$API" -o /tmp/_rs.json
python3 - "$ACTION" <<'PY'
import json,sys
action=sys.argv[1]
r=json.load(open('/tmp/_rs.json'))
if r.get('message'):
    sys.exit('[FAIL] 读取规则集失败: '+str(r.get('message')))
for k in ['id','node_id','created_at','updated_at','html_url','url']:
    r.pop(k,None)
r['enforcement']='disabled' if action=='disable' else 'active'
json.dump(r,open('/tmp/_rs_payload.json','w'))
print('rules:',[x['type'] for x in r.get('rules',[])])
PY

HTTP=$(curl -s --max-time 20 -X PUT \
  -H "Authorization: Bearer $TOKEN" -H "Accept: application/vnd.github+json" \
  -H "Content-Type: application/json" --data-binary @/tmp/_rs_payload.json \
  "$API" -o /tmp/_rs_resp.json -w '%{http_code}')
[ "$HTTP" = "200" ] || { echo "[FAIL] 更新规则集 HTTP=$HTTP" >&2; cat /tmp/_rs_resp.json >&2; exit 1; }

# 回读确认。
curl -s --max-time 20 -H "Authorization: Bearer $TOKEN" -H "Accept: application/vnd.github+json" "$API" -o /tmp/_rs_check.json
NOW=$(python3 -c "import json; print(json.load(open('/tmp/_rs_check.json')).get('enforcement'))")
EXPECT=$([ "$ACTION" = disable ] && echo disabled || echo active)
[ "$NOW" = "$EXPECT" ] || { echo "[FAIL] 回读 enforcement=${NOW}，期望=${EXPECT}" >&2; exit 1; }

rm -f /tmp/_rs.json /tmp/_rs_payload.json /tmp/_rs_resp.json /tmp/_rs_check.json
echo "[PASS] enforcement=${NOW}"
echo "===== RULESET TOGGLE: PASSED ====="
