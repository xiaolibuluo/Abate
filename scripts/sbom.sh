#!/usr/bin/env bash
# =============================================================================
# Abate SBOM 生成（XCU-REL）
#
# 产出 SPDX 2.3 JSON：合并两部分事实来源（SSOT）：
#   1) go list -m all 的全部 Go 模块依赖（离线，GOPROXY=off）。
#   2) third_party/LICENSE-MANIFEST.md 登记的开源并入项（含许可证审查结论）。
#
# 全程离线，不访问网络。输出 ${OUT}/sbom.spdx.json。
# =============================================================================
set -euo pipefail

cd "$(dirname "$0")/.." || exit 2

VERSION="${1:-unknown}"
OUT="${2:-dist/${VERSION}}"
mkdir -p "$OUT"
SBOM="${OUT}/sbom.spdx.json"
export GOFLAGS="-mod=vendor" GOPROXY=off PATH="/usr/local/go/bin:$PATH:$PATH"

DOC_NS="https://github.com/xiaolibuluo/Abate/releases/${VERSION}/sbom.spdx.json"
NOW="$(date -u +%Y-%m-%dT%H:%M:%SZ)"

# 收集 Go 模块依赖为 JSON 数组片段（直接解析 go.mod，避免 vendor 模式下列表失败）。
modules_json="$(awk '
  /^module /{mod=$2; next}
  /^require \(/{inblock=1; next}
  /^\)/{inblock=0; next}
  /^require /{name=$2; ver=$3; printf "%s|%s\n", name, ver; next}
  inblock && /^[ \t]*[A-Za-z]/{name=$1; ver=$2; printf "%s|%s\n", name, ver}
' go.mod | while IFS='|' read -r name ver; do
  [ -z "$name" ] && continue
  sid="SPDXRef-Go-$(printf '%s' "$name" | tr -c 'A-Za-z0-9' '-')"
  jq -n --arg sid "$sid" --arg name "$name" --arg ver "$ver" \
    '{SPDXID:$sid, name:$name, versionInfo:$ver, downloadLocation:"NOASSERTION", licenseConcluded:"NOASSERTION", supplier:"NOASSERTION"}'
done | jq -s '.')"
# 主模块自身作为根包纳入 SBOM。
root_json="$(jq -n --arg ver "$VERSION" '{SPDXID:"SPDXRef-Go-abate", name:"abate", versionInfo:$ver, downloadLocation:"https://github.com/xiaolibuluo/Abate", licenseConcluded:"NOASSERTION", supplier:"Organization: 小黎的部落"}')"

# 收集 third_party 登记项（从 LICENSE-MANIFEST.md 表格行提取组件 + 许可证）。
tp_json="$(awk -F'|' '/^\| /{
  gsub(/[ \t]+/,"",$2); comp=$2;
  if (comp=="组件"||comp==""||comp=="---") next;
  lic=$5; gsub(/[ \t]+/,"",lic);
  sid="SPDXRef-TP-"comp; gsub(/[^A-Za-z0-9]/,"-",sid);
  printf "{\"SPDXID\":\"%s\",\"name\":\"%s\",\"versionInfo\":\"\",\"downloadLocation\":\"NOASSERTION\",\"licenseConcluded\":\"%s\",\"supplier\":\"NOASSERTION\"}\n", sid, comp, (lic==""?"NOASSERTION":lic);
}' third_party/LICENSE-MANIFEST.md | jq -s '.')"

# 合并并写出 SPDX 2.3 文档。
jq -n \
  --arg ver "$VERSION" \
  --arg ns "$DOC_NS" \
  --arg now "$NOW" \
  --argjson go "$modules_json" \
  --argjson tp "$tp_json" \
  --argjson root_json "$root_json" \
  '{
    spdxVersion: "SPDX-2.3",
    dataLicense: "CC0-1.0",
    SPDXID: "SPDXRef-DOCUMENT",
    name: ("Abate-" + $ver + "-SBOM"),
    documentNamespace: $ns,
    creationInfo: {
      created: $now,
      creators: ["Tool: abate-sbom/1.0", "Organization: 小黎的部落"]
    },
    packages: ([$root_json] + $go + $tp)
  }' > "$SBOM"

count="$(jq '.packages | length' "$SBOM")"
echo "[PASS] SBOM 生成: ${SBOM}（${count} 个包：Go 模块 + third_party 登记）"
