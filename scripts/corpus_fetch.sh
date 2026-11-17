#!/usr/bin/env bash
# =============================================================================
# Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
# Abate is dual-licensed under AGPLv3 and a commercial license.
# For commercial licensing, contact xiaolibuluo@163.com.
# Source: https://github.com/xiaolibuluo/Abate
# =============================================================================
# 离线语料准备工具（T01）。门禁绝不调用本脚本，Go 侧完全不联网。
# 每个 tarball 以 tag + tarball_sha256 为权威溯源：4 项 commit 因 GitHub API
# 限流仍为 RATE_LIMIT_PENDING，此处不伪造任何 commit 哈希。
# 逐包重算 tarball sha256，不符立即 exit 1；已落地目录幂等跳过重复解压。
# =============================================================================
set -euo pipefail

CACHE_DIR=".artifacts/corpus_cache"
DEST_ROOT="third_party/_sources/_validation"

# 数据表字段：tarball|lang|project|tag|spdx|tarball_sha256|目标子目录
# sha256 逐字来自 abate.e2e.upstream/v1 溯源清单（11 个 pinned tarball）。
PINNED=(
  "go.tgz|go|gin-gonic/gin|v1.9.1|MIT|704d6924b821c60a9ef3ad373d86aa51cc316259cb1c417f7652b23c6bda7ce3|_go/_gin_v1_9_1"
  "py.tgz|py|psf/requests|v2.31.0|Apache-2.0|24bc1acfdf14f8b94cb2b6a8c38ef9a99bdd146a2976b5a2db7706fba0bcb7cd|_py/_requests_v2_31_0"
  "py2.tgz|py|pallets/click|8.1.7|BSD-3-Clause|89251974dba8552b4e22990ca34adfb93a47ba7deb27fe7358a6661a09ca8793|_py/_click_8_1_7"
  "js.tgz|js|lodash/lodash|4.17.21|MIT|b06f51220fe2a6e66af19f22d787cc8d4fcb75353d8da8a30c7376d3eac58890|_js/_lodash_4_17_21"
  "js_prettier.tgz|js|prettier/prettier|3.1.1|MIT|95d9d6a95ff3e8eec383d2a7c23b0a497bface687e90e6a280823a942657f61a|_js/_prettier"
  "js_express.tgz|js|expressjs/express|4.18.2|MIT|a9ace1a92a2b0a7e3fcd7a7612553d6208a464561c15ad38f349fc3ab0d7cda0|_js/_express"
  "js_axios.tgz|js|axios/axios|v1.6.2|MIT|f4d6848137dd08215d64390fbdacdc9ab325ce3e26d3075f52e09e01d238ef0c|_js/_axios"
  "c.tgz|c|DaveGamble/cJSON|v1.7.16|MIT|451131a92c55efc5457276807fc0c4c2c2707c9ee96ef90c47d68852d5384c6c|_c/_cjson_v1_7_16"
  "cs.tgz|cs|JamesNK/Newtonsoft.Json|13.0.3|MIT|3dc1eeeb8f781122e0713481bad993329c669feb9928cc4ec86e0ab24d9eff9e|_cs/_newtonsoft_13_0_3"
  "rs.tgz|rs|serde-rs/serde|v1.0.193|MIT OR Apache-2.0|0449bb10c414d8614c16ac9b3b501bb50876b36d5eaafdc4708ecc1d2fcef8ba|_rs/_serde_v1_0_193"
  "kt.tgz|kt|square/okhttp|parent-4.12.0|Apache-2.0|9241fa2365e401391713977463e19c3ce9f79afb160e962eafd45f41cd37a5b5|_kt/_okhttp_4_12_0"
)

# tarballURL 由 project + tag 按 GitHub codeload 的 tag 归档规则构造。
tarballURL() {
  printf 'https://codeload.github.com/%s/tar.gz/refs/tags/%s' "$1" "$2"
}

# fetchOne 校验一个包的真实字节，并在目标目录缺失时解压一次。
fetchOne() {
  local tarball="$1" project="$2" tag="$3" want="$4" dest="$5"
  local cached="${CACHE_DIR}/${tarball}"
  local url
  url="$(tarballURL "$project" "$tag")"
  if [ ! -f "$cached" ]; then
    echo ">>> 下载 ${tarball} <- ${url}"
    curl -fL --retry 3 --connect-timeout 15 -o "$cached" "$url"
  fi
  local actual
  actual="$(sha256sum "$cached" | cut -d' ' -f1)"
  if [ "$actual" != "$want" ]; then
    echo "[FAIL] tarball sha256 不符: ${tarball} expect=${want} actual=${actual}" >&2
    return 1
  fi
  echo "[OK] ${tarball} sha256=${actual} (${project} @ ${tag})"
  if [ -d "${DEST_ROOT}/${dest}" ]; then
    echo "     幂等：${DEST_ROOT}/${dest} 已落地，跳过重复解压"
    return 0
  fi
  mkdir -p "${DEST_ROOT}/${dest}"
  tar -xzf "$cached" -C "${DEST_ROOT}/${dest}" --strip-components=1
  echo "     解压完成：${DEST_ROOT}/${dest}"
}

echo "=== Abate E2E 语料准备：${#PINNED[@]} 个 pinned tarball ==="
mkdir -p "$CACHE_DIR"
for row in "${PINNED[@]}"; do
  IFS='|' read -r tarball lang project tag spdx want dest <<< "$row"
  echo "--- lang=${lang} project=${project} tag=${tag} spdx=${spdx}"
  fetchOne "$tarball" "$project" "$tag" "$want" "$dest"
done
echo "=== 全部 ${#PINNED[@]} 个 tarball 的 sha256 重算通过；目录均位于 _ 前缀内 ==="
