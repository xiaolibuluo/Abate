#!/usr/bin/env bash
# =============================================================================
# Abate 强制推送（Delivery Ops）
#
# 用途：把本地 main 分支与版本 tag 强制同步到 origin（git push --force）。
# 通道：走 SSH remote（git@github.com:...），依赖本机 ~/.ssh/ 下的部署密钥，
#       不把任何令牌/密钥写进命令行、日志或仓库（凭据由 ssh-agent/IdentityFile 在进程内读取）。
#
# 安全护栏：
#   - 拒绝空仓库、无 origin、工作树有未提交改动时默认中止（--allow-dirty 可覆盖）。
#   - 推送后用 git ls-remote 权威回读，断言远端 sha 与本地 sha 逐字一致，否则非零退出。
#   - 全程不回显任何凭据。
#
# 用法：
#   scripts/force-push.sh                 # 强制推送 main + 全部 tag
#   scripts/force-push.sh --tag v1.0.0    # 强制推送 main + 指定 tag
#   scripts/force-push.sh --dry-run       # 只做前置校验，不真正推送
#   scripts/force-push.sh --allow-dirty   # 允许工作树有未提交改动时推送
#
# 配套：若仓库启用了 pull_request / non_fast_forward Ruleset，直接强制推送会被
#       拒绝（"Changes must be made through a pull request"）。此时按顺序：
#         1) GITHUB_TOKEN_FILE=... scripts/ruleset-toggle.sh disable
#         2) scripts/force-push.sh --tag vX.Y.Z
#         3) GITHUB_TOKEN_FILE=... scripts/ruleset-toggle.sh enable   # 务必立即恢复
# =============================================================================
set -euo pipefail

cd "$(dirname "$0")/.." || exit 2

TAGARG=""
DRY_RUN=0
ALLOW_DIRTY=0
while [ $# -gt 0 ]; do
  case "$1" in
    --tag)
      [ $# -ge 2 ] || { echo "[FAIL] --tag 缺少参数" >&2; exit 2; }
      TAGARG="$2"; shift 2 ;;
    --dry-run) DRY_RUN=1; shift ;;
    --allow-dirty) ALLOW_DIRTY=1; shift ;;
    *) echo "usage: force-push.sh [--tag vX.Y.Z] [--dry-run] [--allow-dirty]" >&2; exit 2 ;;
  esac
done

echo "===== Abate 强制推送 $(date '+%F %T') ====="

# [1] origin 存在且为 SSH（避免令牌 URL 泄漏）。
ORIGIN="$(git remote get-url origin 2>/dev/null || true)"
[ -n "$ORIGIN" ] || { echo "[FAIL] 未配置 origin" >&2; exit 1; }
case "$ORIGIN" in
  git@*|ssh://*) echo "[PASS] origin 为 SSH 通道（脱敏）" ;;
  *) echo "[FAIL] origin 非 SSH 通道，拒绝推送（避免令牌 URL 泄漏）: ${ORIGIN%%@*}@***" >&2; exit 1 ;;
esac

# [2] 工作树洁净检查。
DIRTY="$(git status --porcelain | wc -l | tr -d ' ')"
if [ "$DIRTY" != "0" ]; then
  if [ "$ALLOW_DIRTY" = "1" ]; then
    echo "[WARN] 工作树有 ${DIRTY} 处未提交改动（--allow-dirty 已放行）"
  else
    echo "[FAIL] 工作树有 ${DIRTY} 处未提交改动，拒绝强制推送（可用 --allow-dirty 覆盖）" >&2
    exit 1
  fi
else
  echo "[PASS] 工作树洁净"
fi

# [3] 本地 HEAD 与 tag 基线。
LOCAL_HEAD="$(git rev-parse HEAD)"
BRANCH="$(git rev-parse --abbrev-ref HEAD)"
[ "$BRANCH" = "main" ] || { echo "[FAIL] 当前分支为 ${BRANCH}，强制推送仅允许 main" >&2; exit 1; }
echo "[PASS] 分支=main 本地HEAD=${LOCAL_HEAD:0:12}"

TAGS=()
if [ -n "$TAGARG" ]; then
  git rev-parse -q --verify "refs/tags/${TAGARG}" >/dev/null || { echo "[FAIL] tag 不存在: ${TAGARG}" >&2; exit 1; }
  TAGS=("$TAGARG")
else
  while IFS= read -r t; do [ -n "$t" ] && TAGS+=("$t"); done < <(git tag --list)
fi
echo "[PASS] 待推送 tag: ${TAGS[*]:-（无）}"

if [ "$DRY_RUN" = "1" ]; then
  echo "[DRY-RUN] 跳过实际推送"
  exit 0
fi

# [4] 强制推送 main。
echo "--- force push main ---"
GIT_TERMINAL_PROMPT=0 git push --force origin main
echo "[PASS] main 强制推送完成"

# [5] 强制推送 tag。
for t in "${TAGS[@]}"; do
  echo "--- force push tag ${t} ---"
  GIT_TERMINAL_PROMPT=0 git push --force origin "refs/tags/${t}:refs/tags/${t}"
  echo "[PASS] tag ${t} 强制推送完成"
done

# [6] 权威回读：git ls-remote 断言远端与本地一致。
echo "--- 回读校验 ---"
REMOTE_HEAD="$(timeout 40 git ls-remote origin refs/heads/main 2>/dev/null | awk '{print $1}')"
[ "${REMOTE_HEAD:-}" = "$LOCAL_HEAD" ] || { echo "[FAIL] 远端 main=${REMOTE_HEAD:-空} 与本地=${LOCAL_HEAD} 不一致" >&2; exit 1; }
echo "[PASS] 远端 main == 本地 HEAD (${LOCAL_HEAD:0:12})"
for t in "${TAGS[@]}"; do
  LOCAL_T="$(git rev-parse "refs/tags/${t}")"
  REMOTE_T="$(timeout 40 git ls-remote origin "refs/tags/${t}" 2>/dev/null | awk '{print $1}')"
  [ "${REMOTE_T:-}" = "$LOCAL_T" ] || { echo "[FAIL] 远端 tag ${t}=${REMOTE_T:-空} 与本地=${LOCAL_T} 不一致" >&2; exit 1; }
  echo "[PASS] 远端 tag ${t} == 本地 (${LOCAL_T:0:12})"
done

echo "===== FORCE PUSH: PASSED ====="
