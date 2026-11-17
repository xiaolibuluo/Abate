#!/usr/bin/env bash
# =============================================================================
# Abate 工具链与开源源码采购引导脚本（幂等，镜像优先，对齐项目版本）
# 严禁降级：Go 版本由 go.mod 的 go 1.26.11 经 GOTOOLCHAIN=auto 自动拉取。
# 所有 Go 依赖走 GOPROXY 镜像（mirrors.tencent.com/go）。
# 不下载任何非 Go 源码（遵守「必须 Go」红线）。
# =============================================================================
set -uo pipefail

cd "$(dirname "$0")/.." || exit 2

# ---- 环境与镜像 ----
export GOPROXY="https://mirrors.tencent.com/go/,direct"
export GOTOOLCHAIN=auto
export GOFLAGS="-mod=mod"
export GOPATH="$(go env GOPATH)"
export PATH="$PATH:$GOPATH/bin"

log() { echo -e "\033[1;33m[BOOTSTRAP] $1\033[0m"; }
ok()  { echo -e "\033[0;32m[OK] $1\033[0m"; }

# ---- 1. Go toolchain（对齐 go.mod 的 go 1.26.11）----
log "确保 Go 工具链 = $(grep '^go ' go.mod | awk '{print $2}')"
go version

# ---- 2. golangci-lint v2（CI 指定 /v2）----
log "安装 golangci-lint v2.1.6（镜像）"
if command -v golangci-lint >/dev/null 2>&1 && golangci-lint version 2>/dev/null | grep -q 'v2'; then
  ok "golangci-lint 已是 v2: $(golangci-lint version 2>&1 | head -1)"
else
  go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6 && ok "golangci-lint 安装完成" || echo "[FAIL] golangci-lint 安装失败"
fi

# ---- 3. 辅助 Go 工具（均为 Go、非 GPL）----
for spec in "mvdan.cc/gofumpt@latest" "github.com/denis-tingajkin/goheader@latest" "github.com/google/go-licenses@latest"; do
  bin="$(basename "${spec%@*}")"
  log "安装辅助工具 $bin"
  go install "$spec" && ok "$bin 安装完成" || echo "[FAIL] $bin 安装失败"
done

# ---- 4. 源码接入的开源（Go、非 GPL、镜像）----
log "下载 yaml.v3 源码（MIT, M1 CES）"
go mod download gopkg.in/yaml.v3@v3.0.1 && ok "yaml.v3 源码已入 module cache" || echo "[FAIL] yaml.v3 下载失败"

log "下载 gVisor 源码（Apache-2.0, M4 沙箱，用户指定）"
go mod download gvisor.dev/gvisor@latest && ok "gVisor 源码已入 module cache" || echo "[FAIL] gVisor 下载失败（不影响 M1）"

# ---- 汇总 ----
log "已安装工具："
ls -1 "$GOPATH/bin" 2>/dev/null || true
log "bootstrap 完成（非 Go 源码按红线未下载，见 docs/procurement.md §5）"
