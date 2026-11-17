#!/usr/bin/env bash
# =============================================================================
# Abate M4 插件发布流水线（本地模拟 / CI 复用）
#
# 证明 "commit -> publish <= 1h" 的可执行证据：
#   build -> race 测试 -> CLI 测试 -> 签名 manifest -> 构建 CLI 二进制
#   -> load/list/call 验证产物 -> 记录真实墙钟起止时间戳与有界耗时。
#
# 全流程离线（-mod=vendor，GOPROXY=off），仅 stdlib + vendored yaml.v3。
# 时间戳只用于证据记录，不参与任何插件输出，因此不破坏确定性红线。
#
# 修复记录（M4 QA 第 1 轮判定 FAIL｜Engineer）：
#   A1 每次运行先清 OUT_DIR/.abate 会话状态与旧证据文件，使二次运行等价于
#      首次运行；证据只在全流程（含 SLA 断言）成功后写入。
#   A2 OUT_DIR 归一为绝对路径后再拼 BIN / MANIFEST，绝对路径入参不再拼错。
#   A3 脚本内版本探测自举 Go 1.26.11 工具链，不依赖调用方预先 export。
# =============================================================================
set -euo pipefail

cd "$(dirname "$0")/.."

# ---- A3：工具链自举（版本探测，不盲从外部继承的 GOROOT）----
# 沙箱预置工具链可能与项目声明版本不一致；旧脚本用
# GOROOT="${GOROOT:-...}" 继承该错误值，裸跑首步即失败。这里先探测当前
# 可用 go 的版本，只有确认为 1.26.11 才保留调用方环境，否则显式切到受版本
# 固定的 1.26.11 工具链并重置 PATH。调用方已正确 export 时不做任何改动。
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
  *"$REQUIRED_GO_VERSION"*)
    # 调用方已提供正确工具链：保持 GOROOT / PATH 不变，不破坏既有环境。
    ;;
  *)
    # 裸跑（env -u GOROOT）或继承到旧工具链：显式切到 Go 1.26.11。
    export GOROOT="$FALLBACK_GOROOT"
    export PATH="$GOROOT/bin:$PATH"
    ;;
esac
export GOTOOLCHAIN=local
export GOSUMDB=off GOPROXY=off GOFLAGS=-mod=vendor

# ---- A2：OUT_DIR 归一为绝对路径 ----
# 旧脚本用 "$(pwd)/$OUT_DIR/abate" 拼接，OUT_DIR 为绝对路径时会得到
# "/workspace/abate//tmp/.../abate" 并以 127 退出。统一先 mkdir 再 cd 取回
# 规范绝对路径，之后一律使用 "$OUT_DIR/..." 拼接。
OUT_DIR="${1:-.artifacts/plugin-publish}"
SLA_SECONDS=3600
mkdir -p "$OUT_DIR"
OUT_DIR="$(cd "$OUT_DIR" && pwd)"
BIN="$OUT_DIR/abate"
MANIFEST="$OUT_DIR/plugin-manifest.json"

# ---- A1：清会话状态与陈旧证据，保证「二次运行 == 首次运行」----
# CLI 会话状态锚定在工作目录的 .abate/plugin-state.json；第 6 步 cd 到 OUT_DIR
# 后 load 会写入该状态。旧脚本只 mkdir -p 不清理，二次运行时 loadPluginSession
# 重放上一次的旧信封，先注册旧公钥；而签名器每次 GenerateKey() 生成新公钥、
# 锚点名固定，随后 registerPluginAnchor 检出「同锚点不同指纹」直接失败。且失败
# 发生在写证据之前，会残留历史 status=PASS 证据被 workflow 整目录上传。
# 因此运行开始即删除会话状态与全部旧证据/中间产物，证据只在全流程成功后写。
rm -rf "$OUT_DIR/.abate"
rm -f "$OUT_DIR/PUBLISH-EVIDENCE.txt" \
      "$OUT_DIR/load.json" "$OUT_DIR/list.json" "$OUT_DIR/call.json"

START_EPOCH="$(date +%s)"
START_UTC="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
COMMIT="${GITHUB_SHA:-$(git rev-parse --short HEAD 2>/dev/null || echo local)}"

echo "==> [1/7] go build ./... (offline vendor)"
go build ./...

echo "==> [2/7] go test -race ./plugin/..."
go test -race -count=1 ./plugin/...

echo "==> [3/7] go test ./ctl/cli/..."
go test -count=1 ./ctl/cli/...

echo "==> [4/7] sign manifest (offline ed25519, private key never persisted)"
go run ./internal/pluginpublish \
  -out "$MANIFEST" \
  -id abate/operator \
  -version 1.0.0 \
  -build-ref "$COMMIT"

echo "==> [5/7] build the CLI binary"
go build -o "$BIN" .

echo "==> [6/7] load / list / call the published artifact"
cd "$OUT_DIR"
"$BIN" plugin load --manifest "$MANIFEST" --type operator | tee load.json
"$BIN" plugin list | tee list.json
"$BIN" plugin call --id abate/operator --entry apply --args ZWJhdGU= | tee call.json
grep -q '"id":"abate/operator"' load.json
grep -q 'abate/operator' list.json
grep -q '"resultSha256"' call.json

echo "==> [7/7] record bounded elapsed time"
END_EPOCH="$(date +%s)"
END_UTC="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
ELAPSED=$((END_EPOCH - START_EPOCH))
MANIFEST_SHA="$(sha256sum "$MANIFEST" | cut -d' ' -f1)"

# SLA 断言前置于证据写入：SLA 违规同样不能留下 status=PASS 证据。
if [ "$ELAPSED" -gt "$SLA_SECONDS" ]; then
  echo "plugin-publish: SLA violated: ${ELAPSED}s > ${SLA_SECONDS}s" >&2
  exit 1
fi

# 证据只在全流程成功后写入（A1）：此前任一步骤失败都会因 set -e 退出，
# 且运行开始时已删除旧证据，绝不出现「失败却上传历史 PASS 证据」。
cat > PUBLISH-EVIDENCE.txt <<EOF
pipeline=abate-plugin-publish
commit=$COMMIT
startedAt=$START_UTC
finishedAt=$END_UTC
elapsedSeconds=$ELAPSED
slaSeconds=$SLA_SECONDS
sla=commit-to-publish-must-be-at-most-3600-seconds
status=PASS
manifestSha256=$MANIFEST_SHA
steps=build,race-test,cli-test,sign,binary,load,list,call
offline=true
EOF

echo "plugin-publish: PASS in ${ELAPSED}s (SLA ${SLA_SECONDS}s), evidence: $OUT_DIR/PUBLISH-EVIDENCE.txt"
