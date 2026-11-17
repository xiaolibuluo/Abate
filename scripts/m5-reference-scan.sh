#!/usr/bin/env bash
# =============================================================================
# Abate M5 静态门禁：蓝本隔离 / 依赖面冻结 / 确定性红线
#
# 四段检查（对应任务书 Part B 验收）：
#   ① 新增源码禁七项目包名 import / `import "C"` / `os/exec` / `net.Dial` /
#      `http.Get|Post|Do|NewRequest`；并对全仓 .go 的 import 声明行做兜底扫描。
#   ② `go list ./...` 输出不得含 `_references/*.go.txt` 或 `_m5_algorithms`，
#      且蓝本目录内不得出现可编译的 `.go` 文件。
#   ③ 依赖面一致性：`go.mod` 的 require 集合与 `vendor/modules.txt` 的模块集合
#      双向相等（无幽灵依赖 / 无未 vendor 声明），且全部通过 license-check 非 GPL
#      判定；新增依赖必须三方同步并经审查。
#   ④ 新增 .go 禁 `time.Now` / `time.Since` / 全局 `math/rand` / `os.Environ`。
# 命中即非零退出，并打印命中文件与行号。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0; REPORT=""

cd "$(dirname "$0")/.." || exit 2

# ---- 工具链自举：与 scripts/plugin-publish.sh 同一策略，裸跑亦可用 ----
REQUIRED_GO_VERSION="go1.26.11"
FALLBACK_GOROOT="/usr/local/go-packages/pkg/mod/golang.org/toolchain@v0.0.1-go1.26.11.linux-amd64"

detect_go_version() {
  if [ -n "${GOROOT:-}" ] && [ -x "$GOROOT/bin/go" ]; then
    "$GOROOT/bin/go" version 2>/dev/null || true
  elif command -v go >/dev/null 2>&1; then
    go version 2>/dev/null || true
  fi
}

CURRENT_GO_VERSION="$(detect_go_version)"
case "$CURRENT_GO_VERSION" in
  *"$REQUIRED_GO_VERSION"*) ;;
  *)
    export GOROOT="$FALLBACK_GOROOT"
    export PATH="$GOROOT/bin:$PATH"
    ;;
esac
export GOTOOLCHAIN=local
export GOSUMDB=off GOPROXY=off GOFLAGS=-mod=vendor

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

# 本批（B01）新增 Go 源码目录；后续批次可经 M5_SCAN_DIRS 覆盖，默认不变。
SCAN_DIRS=("lang/frontkit")
if [ -n "${M5_SCAN_DIRS:-}" ]; then
  read -r -a SCAN_DIRS <<< "$M5_SCAN_DIRS"
fi
NEW_GO="$(find "${SCAN_DIRS[@]}" -type f -name '*.go' 2>/dev/null | sort)"
ALL_GO="$(find . -type f -name '*.go' -not -path './vendor/*' -not -path './third_party/*' | sort)"

# ① 禁用 import：仅匹配真正的 import 语句——单行 `import "p"`，或 import
#    块内独立一行 `"p"` / `alias "p"` / `. "p"`。关键约束：行内不得含逗号且
#    行尾结束（可带行尾注释），从而排除 `forbidden := []string{ "math/rand", }`
#    这类测试夹具里的字符串切片数据（旧正则的误报来源）。
PATTERN_IMPORT='^[[:space:]]*(import[[:space:]]+)?(\.?[_A-Za-z][_A-Za-z0-9]*[[:space:]]+)?"([^"]*(os/exec|math/rand|ruff|tree[-_]sitter|swc|libclang|clang|llvm|roslyn|codeanalysis|kotlin[-_]compiler|jetbrains|opentelemetry|go\.opentelemetry\.io|/syn|/quote)[^"]*|C)"[[:space:]]*(//.*)?$'
# ① 禁用调用：仅对新增源码扫描（全仓既有夹具含字符串常量，不能全仓匹配）。
PATTERN_CALL='net\.Dial|http\.(Get|Post|PostForm|Do|NewRequest|NewRequestWithContext|Head)\('
# ④ 确定性红线。
PATTERN_TIME='time\.Now|time\.Since|math/rand|os\.Environ'

banner "Abate M5 Reference Gate - 蓝本隔离 / 依赖面 / 确定性"
echo "新增源码目录: ${SCAN_DIRS[*]}（共 $(printf '%s\n' "$NEW_GO" | grep -c .) 个 .go 文件）"

# ---- ① 禁用 import 与外部调用 ----
echo ">>> [1] 新增源码 + 全仓 import 行：禁七项目包名 / CGO / 子进程 / 客户端调用"
HITS_1=0
for f in $NEW_GO; do
  [ -z "$f" ] && continue
  hits="$(grep -nE "$PATTERN_IMPORT|$PATTERN_CALL" "$f" 2>/dev/null || true)"
  if [ -n "$hits" ]; then
    fail "新增源码命中禁用 import 或外部调用: $f"
    printf '  %s:%s\n' "$f" "$hits"
    HITS_1=1
  fi
done
for f in $ALL_GO; do
  [ -z "$f" ] && continue
  hits="$(grep -nE "$PATTERN_IMPORT" "$f" 2>/dev/null || true)"
  if [ -n "$hits" ]; then
    fail "全仓 import 行命中禁用依赖: $f"
    printf '  %s:%s\n' "$f" "$hits"
    HITS_1=1
  fi
done
[ "$HITS_1" -eq 0 ] && pass "[1] 零禁用 import / CGO / 子进程 / 客户端调用"

# ---- ② 蓝本不被 go list 收录 ----
echo ">>> [2] 蓝本隔离：_references/*.go.txt 不进入 go list ./..."
BLUEPRINT="third_party/_references/_m5_algorithms.go.txt"
if [ ! -f "$BLUEPRINT" ]; then
  fail "M5 蓝本文件缺失: $BLUEPRINT"
fi
GO_IN_REF="$(find third_party/_references -type f -name '*.go' 2>/dev/null || true)"
if [ -n "$GO_IN_REF" ]; then
  fail "_references 目录出现可编译 .go 文件: $GO_IN_REF"
fi
PKGS="$(go list ./... 2>&1)"; PKG_RC=$?
if [ "$PKG_RC" -ne 0 ]; then
  fail "go list ./... 执行失败: $PKGS"
else
  hits="$(printf '%s\n' "$PKGS" | grep -nE '_references|\.go\.txt|_m5_algorithms' || true)"
  if [ -n "$hits" ]; then
    fail "go list ./... 收录了只读蓝本路径"
    printf '  %s\n' "$hits"
  else
    pass "[2] go list ./... 共 $(printf '%s\n' "$PKGS" | grep -c .) 个包，零蓝本路径"
  fi
fi

# ---- ③ 依赖面一致性（M6 精度提升）----
# M5 阶段把冻结面写死为「仅 yaml.v3」，用于隔离七项目蓝本源码。M6 按
# docs/plan/08-crosscutting.md 正式引入 wazero（Apache-2.0）及其间接依赖
# x/sys，二者均由 go.mod 显式 require、vendor/modules.txt 同步登记，并经
# license-check 判定非 GPL。规则据此改为以 go.mod 为权威的双向一致性校验：
#   - vendor 中的每个模块必须在 go.mod require 中（无幽灵依赖）；
#   - go.mod 的每个 require 必须出现在 vendor 中（离线可构建）；
#   - 所有模块必须通过 license-check 的非 GPL 判定。
# 新增依赖仍需 go.mod / vendor / 许可证三方同步并经审查，冻结语义不放宽。
echo ">>> [3] 依赖面：go.mod require 与 vendor/modules.txt 双向一致 + 非 GPL"
HITS_3=0
REQ_MODS="$({
  sed -n -E '/^require \(/,/^\)/ s/^[[:space:]]*([^ )]+) v[0-9].*/\1/p' go.mod
  sed -n -E 's/^require[[:space:]]+([^ ]+) v[0-9].*/\1/p' go.mod
} | sort -u)"
VEND_MODS="$(grep -E '^# [^ ]+ v' vendor/modules.txt 2>/dev/null | awk '{print $2}' | sort -u)"
if [ -z "$REQ_MODS" ] || [ -z "$VEND_MODS" ]; then
  fail "无法取得 go.mod require 或 vendor/modules.txt 模块清单"
  HITS_3=1
else
  MISSING="$(comm -23 <(printf '%s\n' "$REQ_MODS") <(printf '%s\n' "$VEND_MODS"))"
  EXTRA="$(comm -13 <(printf '%s\n' "$REQ_MODS") <(printf '%s\n' "$VEND_MODS"))"
  if [ -n "$MISSING" ]; then
    fail "go.mod require 未同步到 vendor/modules.txt: $(echo $MISSING)"
    HITS_3=1
  fi
  if [ -n "$EXTRA" ]; then
    fail "vendor/modules.txt 出现未在 go.mod 声明的幽灵依赖: $(echo $EXTRA)"
    HITS_3=1
  fi
  if ! bash scripts/license-check.sh >/dev/null 2>&1; then
    fail "依赖许可证校验未通过（scripts/license-check.sh）"
    HITS_3=1
  fi
  printf '%s\n' "$VEND_MODS" | sed 's/^/  [OK] /'
fi
[ "$HITS_3" -eq 0 ] && pass "[3] 依赖面三方一致（go.mod / vendor / 许可证）"

# ---- ④ 确定性红线 ----
echo ">>> [4] 新增 .go：禁 time.Now / time.Since / 全局 math/rand / os.Environ"
HITS_4=0
for f in $NEW_GO; do
  [ -z "$f" ] && continue
  hits="$(grep -nE "$PATTERN_TIME" "$f" 2>/dev/null || true)"
  if [ -n "$hits" ]; then
    fail "新增源码命中确定性红线: $f"
    printf '  %s:%s\n' "$f" "$hits"
    HITS_4=1
  fi
done
[ "$HITS_4" -eq 0 ] && pass "[4] 零 time.Now / 全局 math/rand / os.Environ"

echo ""
if [ "$FAILED" -ne 0 ]; then
  banner "M5 REFERENCE GATE: FAILED"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
banner "M5 REFERENCE GATE: PASSED"
exit 0
