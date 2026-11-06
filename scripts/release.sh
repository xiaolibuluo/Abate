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
# release.sh —— 构建发布产物并生成完整性凭据（任务三）：
#   1. 构建二进制 abate
#   2. 生成 SHA-256 校验值文件 abate-<version>.sha256
#   3. 生成 GPG 分离签名 abate-<version>.sig（armor）
# 验证方法见 docs/SIGNING.md《数字签名与验证指南》。
# 退出码：0 成功 / 1 失败 / 2 用法错误。
set -euo pipefail

VERSION="dev"
OUTDIR="dist"
KEY_ID=""

usage() {
  cat >&2 <<'USAGE'
用法: release.sh [-v version] [-o outdir] [-k key-id]
  -v version   版本号（默认 dev，注入 -ldflags -X main.version）
  -o outdir    输出目录（默认 dist）
  -k key-id    GPG 签名密钥 ID（默认取首个可用私钥）
  -h           帮助
USAGE
}

while getopts ":v:o:k:h" opt; do
  case "$opt" in
    v) VERSION="$OPTARG" ;;
    o) OUTDIR="$OPTARG" ;;
    k) KEY_ID="$OPTARG" ;;
    h) usage; exit 0 ;;
    *) usage; exit 2 ;;
  esac
done

# 前置检查：go 工具链。
if ! command -v go >/dev/null 2>&1; then
  echo "release: go toolchain not found in PATH" >&2
  exit 1
fi

# 前置检查：gpg 与签名私钥（缺失时明确提示，不静默跳过）。
if ! command -v gpg >/dev/null 2>&1; then
  echo "release: gpg not found — 无法生成 GPG 签名。请安装 gnupg 后重试。" >&2
  echo "release: gpg not found; cannot produce the GPG signature. Install gnupg and retry." >&2
  exit 1
fi
if [ -z "$KEY_ID" ]; then
  KEY_ID="$(gpg --list-secret-keys --with-colons 2>/dev/null | awk -F: '/^sec/ {print $5; exit}')" || true
fi
if [ -z "$KEY_ID" ]; then
  echo "release: 未找到可用 GPG 私钥 — 请用 -k 指定密钥或先导入私钥。" >&2
  echo "release: no usable GPG secret key found — pass -k <key-id> or import one first." >&2
  exit 1
fi

BIN="abate"
mkdir -p "$OUTDIR"

echo "==> 构建 Building $BIN (version=$VERSION)"
go build -ldflags "-X main.version=$VERSION" -o "$OUTDIR/$BIN" ./cmd/abate

echo "==> 生成 SHA-256 校验值 Generating SHA-256 checksum"
SHA_FILE="$OUTDIR/$BIN-$VERSION.sha256"
(cd "$OUTDIR" && sha256sum "$BIN" > "$BIN-$VERSION.sha256")
cat "$SHA_FILE"

echo "==> 生成 GPG 分离签名 Generating GPG detached signature (key=$KEY_ID)"
SIG_FILE="$OUTDIR/$BIN-$VERSION.sig"
gpg --batch --yes --local-user "$KEY_ID" --detach-sign --armor --output "$SIG_FILE" "$OUTDIR/$BIN"

echo "==> 完成 Done. 产物 Artifacts:"
ls -l "$OUTDIR/$BIN" "$SHA_FILE" "$SIG_FILE"
echo "验证方法见 docs/SIGNING.md — see docs/SIGNING.md for verification."
