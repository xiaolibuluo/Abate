#!/usr/bin/env bash
#
# Abate TLA+ 形式化验证门禁
#
# 对 scripts/tla/specs 下的全部 .tla 规格运行 TLC 模型检验。
# 任一规格未通过（语义错误 / 不变量违反 / 活性性质不成立）即整体退出码非 0。
#
# 依赖：Java 运行时（java 在 PATH 中）。TLA+ 工具链 tla2tools.jar 优先使用
# scripts/tla/tla2tools.jar，缺失时自动从官方 release 下载缓存到本地。
#
# 用法：bash scripts/tla/check.sh

set -u

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SPECS_DIR="${SCRIPT_DIR}/specs"
JAR="${SCRIPT_DIR}/tla2tools.jar"
TLA_RELEASE="https://github.com/tlaplus/tlaplus/releases/download/v1.8.0/tla2tools.jar"

if ! command -v java >/dev/null 2>&1; then
  echo "[FAIL] 未找到 java，请先安装 Java 运行时" >&2
  exit 1
fi

# 确保工具链可用
if [[ ! -f "${JAR}" ]]; then
  echo "[INFO] 本地无 tla2tools.jar，尝试下载缓存到 ${JAR}"
  if command -v curl >/dev/null 2>&1; then
    curl -fsSL -o "${JAR}" "${TLA_RELEASE}" || {
      echo "[FAIL] 下载 tla2tools.jar 失败（需要网络或手动放置到 ${JAR}）" >&2
      exit 1
    }
  else
    echo "[FAIL] 无 curl 且无本地 tla2tools.jar，请手动放置" >&2
    exit 1
  fi
fi

rc=0
shopt -s nullglob
for spec in "${SPECS_DIR}"/*.tla; do
  name="$(basename "${spec}" .tla)"
  cfg="${SPECS_DIR}/${name}.cfg"
  if [[ ! -f "${cfg}" ]]; then
    echo "[SKIP] ${name}: 缺少对应 .cfg"
    continue
  fi
  echo "===== ${name} ====="
  out="$(java -jar "${JAR}" "${spec}" -config "${cfg}" 2>&1)"
  if echo "${out}" | grep -qiE 'Model checking completed\. No error has been found\.'; then
    states="$(echo "${out}" | grep -oE '[0-9]+ distinct states' | head -1)"
    echo "[PASS] ${name}  (${states})"
  else
    echo "[FAIL] ${name}"
    echo "${out}" | tail -20
    rc=1
  fi
done

if [[ ${rc} -eq 0 ]]; then
  echo "===== TLA+ 形式化验证全部通过 ====="
else
  echo "===== TLA+ 形式化验证存在未通过项 ====="
fi
exit ${rc}
