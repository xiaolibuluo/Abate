#!/usr/bin/env bash
# =============================================================================
# Abate 依赖漏洞扫描门禁（M6 验收标准 2：依赖漏洞扫描全绿）
#
# 策略（事实导向，不伪造、不豁免）：
#   - 使用 Go 官方 govulncheck 扫描**可达调用路径**上的已知漏洞。
#   - 工具缺失时如实报错并给出安装命令（退出码 1），绝不静默跳过。
#   - 扫描需要漏洞数据库网络可达；数据库不可达时**如实判定为环境受限并硬 FAIL**
#     （退出码 1），明确标注「本轮未完成真实可达性判定，须联网恢复后重跑」。
#     严禁以 exit 0 / SKIP 伪造 PASS：未能完成真实扫描者，绝不得宣称安全。
#   - 仅在「基于漏洞数据库完成真实扫描且未发现可达漏洞」时方判定 PASS。
#   - 报告原文落盘 .artifacts/vuln-scan.txt，含每条漏洞的可达调用栈。
# =============================================================================
set -uo pipefail

RED='\033[0.31m'; GREEN='\033[0.32m'; YELLOW='\033[1;33m'; NC='\033[0m'
REPORT_DIR=".artifacts"
OUT="$REPORT_DIR/vuln-scan.txt"

cd "$(dirname "$0")/.." || exit 2
export PATH="$PATH:$(go env GOPATH)/bin"

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }

mkdir -p "$REPORT_DIR"

banner "Abate Vulnerability Gate - govulncheck 依赖漏洞扫描"

GOVULN=""
for cand in "$(command -v govulncheck 2>/dev/null)" "$(go env GOPATH)/bin/govulncheck"; do
  if [ -n "$cand" ] && [ -x "$cand" ]; then GOVULN="$cand"; break; fi
done

if [ -z "$GOVULN" ]; then
  echo -e "${RED}[FAIL] 未找到 govulncheck${NC}"
  echo "安装命令："
  echo "  go install golang.org/x/vuln/cmd/govulncheck@latest"
  echo "然后重跑：bash scripts/vuln-scan.sh"
  exit 1
fi

echo "工具：$GOVULN ($("$GOVULN" --version 2>/dev/null | head -1))"
echo "扫描目标：./... （仅可达调用路径）"

GOFLAGS=-mod=vendor "$GOVULN" ./... >"$OUT" 2>&1
RC=$?

echo "报告：$OUT"
cat "$OUT"

# 判定（事实导向，不伪造、不豁免）：
#   - 非零退出优先区分「数据库不可达（环境受限）」与「发现可达漏洞」与「其它异常」，
#     三者均不得伪造绿灯；数据库不可达明确标注「须联网重跑」，不得 exit 0。
#   - 仅当基于漏洞数据库完成真实扫描、且未发现可达漏洞时，方判定 PASS。
if [ $RC -ne 0 ]; then
  if grep -qE "Your code is affected" "$OUT"; then
    echo -e "${RED}[FAIL] 检测到可达漏洞（govulncheck 退出码 $RC）${NC}"
    exit 1
  fi
  if grep -qE "vuln.go.dev|connection refused|no such host|timeout|dial tcp|no such file or directory|context deadline exceeded" "$OUT"; then
    echo -e "${RED}[FAIL] 漏洞数据库不可达，本轮未完成真实扫描（退出码 $RC）${NC}"
    echo "（如实标注：未做真实可达性判定；须联网恢复后重跑本脚本，绝不伪造 PASS）"
    exit 1
  fi
  echo -e "${RED}[FAIL] govulncheck 异常退出（退出码 $RC），详见报告${NC}"
  exit 1
fi

if grep -q "Your code is affected" "$OUT"; then
  echo -e "${RED}[FAIL] 检测到可达漏洞${NC}"
  exit 1
fi

echo -e "${GREEN}[PASS] 未发现可达漏洞（已基于漏洞数据库完成真实扫描）${NC}"
exit 0
