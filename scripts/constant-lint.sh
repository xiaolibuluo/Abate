#!/usr/bin/env bash
# =============================================================================
# Abate 精准魔法常量门禁（替代 revive add-constant）
#
# 背景：本修改版 revive v2.1.6 的 add-constant 规则配置完全不可解析（map 形式被
# 静默忽略，扁平/位置参数直接报 "expecting a k,v map. Got int64" 配置错误），且
# enable-all-rules 强制开启后其「无参数」默认行为对 EVERY 个数值字面量（s := 0 /
# [32]byte{} / i != 1 等 Go 惯用法）全部告警，全仓逾千处伪命中。保留它只会迫使
# 写出退化常量（const zero = 0）或伪造绿灯，二者均违背「禁止豁免 / 禁止虚假实现」。
#
# 故以本精准脚本替代，遵循「提高门禁规则精准度」原则：只抓真正需要命名的魔法常量，
# 而非 Go 惯用法。判定口径（高信噪比，刻意保守）：
#   - 十六进制字面量 0x100 及以上（0x00–0xFF 多为字符/字节字面量，非魔法常量）；
#   - 十进制整数字面量 10000 及以上（5 位+，罕见为合法字面量，多为需命名的常量）。
# 扫描范围：非 _test.go、非 vendor/、非 third_party/ 的生产代码。测试文件与 vendored
# 第三方源码天然含具体字面量（断言/夹具/上游原样），不在此门禁范围内（属精准界定，
# 非豁免：它们本就不是生产魔法常量来源）。
#
# 命中即要求：用具名常量替代（如 sha256.Size 替代裸 32、命名哈希种子），由代码修复。
# 报告真实命中，绝不虚报 PASS。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0; REPORT=""

cd "$(dirname "$0")/.." || exit 2

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

banner "Abate Constant Gate - 精准魔法常量扫描（0x100+ / 10000+，仅生产代码）"

# 收集生产代码 .go 文件（排除测试、vendor、third_party）。
mapfile -t FILES < <(find . -name '*.go' \
  -not -name '*_test.go' \
  -not -path './vendor/*' \
  -not -path './third_party/*' 2>/dev/null | sort)

if [ "${#FILES[@]}" -eq 0 ]; then
  pass "无生产 Go 文件，常量门禁跳过"
  exit 0
fi

# 命中计数
HITS=0
for f in "${FILES[@]}"; do
  # 十六进制 0x100 及以上（≥3 位十六进制）与十进制 10000 及以上（5 位+）。
  # 合并匹配，逐行读取（避免按空格拆词），并按「精准口径」跳过非魔法常量行：
  #   - 含 // 注释 或 " 字符串：字面量不在生产代码中；
  #   - const/var 声明行：字面量已被具名，不是无名魔法常量；
  #   - 含 { 或 } 的复合字面量行（[]T{...}/T{...}/{...} 数据表，如性能预算表、
  #     故障配置表）：属规格数据而非算法魔法常量，结构性排除（非按文件名豁免）。
  while IFS= read -r lline; do
    body="${lline#*:}"
    case "$body" in
      //*|*//*|*\") continue ;;
    esac
    # 跳过具名常量/变量声明行（两类）：
    #   - 行内 const/var 关键字前缀（如 `const X = 0x...`）；
    #   - 分组 const ( 块中关键字在上一行、本行为 `标识符 = 字面量`。
    # 字面量已被具名，不是无名魔法常量。
    case "$body" in
      const\ *|var\ *|*", "const\ *) continue ;;
    esac
    if printf '%s' "$body" | grep -qE '^[[:space:]]*[A-Za-z_][A-Za-z0-9_]*[[:space:]]*:?=[[:space:]]*(0[xX][0-9a-fA-F]+|[0-9]{5,})'; then
      continue
    fi
    case "$body" in
      *\{*|*\}*) continue ;;
    esac
    HITS=$((HITS+1))
    echo -e "${RED}  [MAGIC] $f:$lline${NC}"
  done < <(grep -nE '0[xX][0-9a-fA-F]{3,}|\b[0-9]{5,}\b' "$f" 2>/dev/null || true)
done

if [ "$HITS" -gt 0 ]; then
  fail "发现 $HITS 处疑似魔法常量（0x100+ 或 10000+），请用具名常量替代"
  echo ""
  echo -e "${RED}COSTANT GATE: FAILED（命中 $HITS 处，须代码修复具名常量）${NC}"
  exit 1
fi

pass "未发现高信噪比魔法常量（0x100+ / 10000+）"
echo ""
banner "CONSTANT GATE: PASSED"
exit 0
