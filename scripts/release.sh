#!/usr/bin/env bash
# =============================================================================
# Abate 发布工程编排（XCU-REL）
#
# 全流程离线：构建（GOFLAGS=-trimpath -mod=vendor，GOPROXY=off，CGO_ENABLED=0）
# -> 校验和清单（含小黎的部落水印行）-> SBOM（SPDX 2.3）-> GPG 签名（可选）
# -> 注解签名 tag（可选，需 git + 密钥）。
#
# 可复现性：二进制仅注入 version + commit（固定 per tag），不注入构建时间，
# 因此双构建字节一致。清单时间戳只用于人类审阅，不参与二进制。
#
# 用法：
#   release.sh [--dry-run] [--demo-sign] vX.Y.Z
#   --dry-run    只做前置校验 + 构建 + 清单，不签名不建 tag（默认行为若无可签密钥）。
#   --demo-sign  在临时 GNUPGHOME 生成一次性密钥完成签名演练（仅本地证据，不对外）。
# =============================================================================
set -euo pipefail

cd "$(dirname "$0")/.." || exit 2

DRY_RUN=0
DEMO_SIGN=0
VERSION=""
WATERMARK="小黎的部落"
while [ $# -gt 0 ]; do
  case "$1" in
    --dry-run) DRY_RUN=1 ;;
    --demo-sign) DEMO_SIGN=1 ;;
    v[0-9]*|*[0-9].[0-9]*.[0-9]*) VERSION="$1" ;;
    *) echo "usage: release.sh [--dry-run] [--demo-sign] vX.Y.Z" >&2; exit 2 ;;
  esac
  shift
done
[ -z "$VERSION" ] && { echo "error: missing version (e.g. v1.0.0)" >&2; exit 2; }
echo "$VERSION" | grep -Eq '^v[0-9]+\.[0-9]+\.[0-9]+$' || { echo "error: version must be vMAJOR.MINOR.PATCH" >&2; exit 2; }

DIST="dist/${VERSION}"
mkdir -p "$DIST"
OUT="$DIST"
BIN_NAME="abate"

# 前置校验：CHANGELOG 含本版本条目（三处版本一致性之一）。
if ! grep -q "## \[${VERSION}\]" CHANGELOG.md 2>/dev/null; then
  echo "[FAIL] CHANGELOG.md 缺少本版本条目 ## [${VERSION}]" >&2; exit 1
fi
echo "[PASS] CHANGELOG 含 ${VERSION} 条目"

# 前置校验：全量构建冒烟（不触发完整门禁，门禁由 CI 常驻执行）。
export GOFLAGS="-mod=vendor" GOPROXY=off CGO_ENABLED=0 PATH="/usr/local/go/bin:$PATH:$PATH"
go build ./... || { echo "[FAIL] go build ./... 失败" >&2; exit 1; }
echo "[PASS] go build ./... 冒烟通过"

# 目标矩阵。
TARGETS="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64"
COMMIT="$(git rev-parse --short HEAD 2>/dev/null || echo unknown)"
LDFLAGS="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}"

ASSETS=()
for t in $TARGETS; do
  os="${t%/*}"; arch="${t#*/}"
  ext=""; [ "$os" = "windows" ] && ext=".exe"
  out="${OUT}/${BIN_NAME}-${VERSION}-${os}-${arch}${ext}"
  echo "== build ${t} =="
  GOOS="$os" GOARCH="$arch" go build -trimpath -ldflags "$LDFLAGS" -o "$out" . || {
    echo "[FAIL] build ${t} 失败" >&2; exit 1; }
  ASSETS+=("$out")
done
echo "[PASS] 构建矩阵完成（${#ASSETS[@]} 个资产）"

# 校验和清单（含水印行）。
MANIFEST="${OUT}/manifest.txt"
{
  echo "# Abate ${VERSION} 发布清单 — 水印: ${WATERMARK}"
  echo "# 生成时间(仅审阅): $(date -u +%Y-%m-%dT%H:%M:%SZ)"
  echo "# 工具链: go $(go version | awk '{print $3}')   commit: ${COMMIT}"
  for a in "${ASSETS[@]}"; do
    sha="$(sha256sum "$a" | awk '{print $1}')"
    echo "${sha}  $(basename "$a")"
  done
} > "$MANIFEST"
echo "[PASS] 校验和清单 ${MANIFEST}"

# SBOM（SPDX 2.3）。
if [ -x scripts/sbom.sh ]; then
  scripts/sbom.sh "$VERSION" "$OUT" || echo "[WARN] SBOM 生成跳过"
else
  echo "[WARN] scripts/sbom.sh 缺失，跳过 SBOM"
fi

# 签名：需要密钥。无密钥或 --dry-run 时仅演练校验和。
SIGN_KEY="${ABATE_SIGN_KEY:-}"
GNUPGHOME_TMP=""
if [ "$DEMO_SIGN" -eq 1 ]; then
  GNUPGHOME_TMP="$(mktemp -d)"
  export GNUPGHOME="$GNUPGHOME_TMP"
  gpg --batch --pinentry-mode loopback --passphrase '' --quick-generate-key "Abate Demo <demo@abate.local>" default default never >/dev/null 2>&1
  SIGN_KEY="$(gpg --list-keys --with-colons 2>/dev/null | awk -F: '/^pub:/{f=1} f&&$1=="fpr"{print $10; exit}')"
fi

if [ "$DRY_RUN" -eq 1 ]; then
  echo "[INFO] --dry-run：跳过签名与 tag（校验和 + 清单 + SBOM 已产出）"
elif [ -n "$SIGN_KEY" ] && command -v gpg >/dev/null 2>&1; then
  for a in "${ASSETS[@]}" "$MANIFEST"; do
    gpg --batch --yes --local-user "$SIGN_KEY" --detach-sign --armor -o "$a.sig" "$a" || echo "[WARN] 签名失败: $a"
  done
  echo "[PASS] 资产 + 清单 GPG 签名完成（密钥 ${SIGN_KEY:0:16}…）"
  if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    git tag -s "${VERSION}" -m "Abate ${VERSION}" && echo "[PASS] 注解签名 tag ${VERSION}"
  else
    echo "[INFO] 非 git 仓库，跳过 tag（CI 中由发布作业执行）"
  fi
else
  echo "[INFO] 无可签密钥（ABATE_SIGN_KEY 未设置且无 --demo-sign）：仅产出校验和清单"
fi

[ -n "$GNUPGHOME_TMP" ] && rm -rf "$GNUPGHOME_TMP"
echo ""
echo "== 发布产物目录: ${OUT} =="
ls -1 "$OUT"
echo "== release.sh 完成（${VERSION}）=="
