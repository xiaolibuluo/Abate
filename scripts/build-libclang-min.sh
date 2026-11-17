#!/usr/bin/env bash
# =============================================================================
# Abate libclang-min 构建脚本（M5 语言前端蓝本：只保留解析 + AST 访问能力）
#
# 目标：从 third_party/_sources/llvm-project 构建只含「解析 + AST 访问」的最小
#       libclang 静态库 libclang-min.a（剥离 clang driver / CodeGen / LLVM 后端），
#       体积上限 50MB，打印实际体积与构建日志路径。
#
# 事实边界（如实声明，不伪造产物）：
#   * 入库的 llvm-project 是**裁剪快照**（只保留 clang/lib、clang/include、
#     llvm/include、llvm/lib/Support）。上游 LLVM 的 CMake 构建系统
#     （llvm/CMakeLists.txt、llvm/cmake/modules、llvm/utils/TableGen）与
#     libclang 的 C API 实现（clang/tools/libclang）**不在库内**，
#     tablegen 生成的 .inc 头文件也不在库内。
#   * 因此本脚本先做**构建前置依赖诊断**：前置齐备则真实执行 CMake 构建；
#     前置缺失则**如实报错并列出缺失项与补齐方式**，绝不产出假产物。
#   * 本脚本不联网、不调用任何外部二进制做解析，只做本地 CMake 构建。
# =============================================================================
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'

cd "$(dirname "$0")/.." || exit 2

SRC="third_party/_sources/llvm-project"
ARTIFACT_DIR=".artifacts/libclang-min"
BUILD_DIR="${ARTIFACT_DIR}/build"
OUT_LIB="${ARTIFACT_DIR}/libclang-min.a"
LOG_FILE="${ARTIFACT_DIR}/build.log"
REPORT_FILE="${ARTIFACT_DIR}/build-report.txt"
MAX_SIZE_MB=50

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { echo -e "${RED}[FAIL] $1${NC}" >&2; exit "${2:-1}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

banner "Abate libclang-min 构建（解析 + AST 访问，最小静态库）"

# ---------------------------------------------------------------- 1. 工具链
banner "1/4 工具链检查"
command -v cmake >/dev/null 2>&1 || fail "缺少 cmake，请安装 cmake ≥ 3.20 后重试" 3
CMAKE_VERSION="$(cmake --version | head -n1)"
echo "  cmake: ${CMAKE_VERSION}"
if command -v ninja >/dev/null 2>&1; then
  GENERATOR="Ninja"
  echo "  ninja: $(ninja --version)"
elif command -v make >/dev/null 2>&1; then
  GENERATOR="Unix Makefiles"
  echo "  make: $(make --version | head -n1)"
else
  fail "缺少构建后端，请安装 ninja 或 make 后重试" 3
fi
if command -v c++ >/dev/null 2>&1; then
  echo "  c++:   $(c++ --version | head -n1)"
elif command -v g++ >/dev/null 2>&1; then
  echo "  g++:   $(g++ --version | head -n1)"
else
  fail "缺少 C++ 编译器，请安装 g++ 或 clang++ 后重试" 3
fi

# ------------------------------------------------------- 2. 构建前置依赖诊断
banner "2/4 构建前置依赖诊断"
[ -d "$SRC" ] || fail "缺少源码目录：${SRC}" 3
mkdir -p "${ARTIFACT_DIR}"
MISSING=()
for required in \
  "llvm/CMakeLists.txt:上游 LLVM 顶层 CMake 入口" \
  "llvm/cmake/modules:上游 LLVM CMake 宏模块（add_llvm_library 等）" \
  "llvm/utils/TableGen:tablegen 源码（生成 .inc 头的唯一来源）" \
  "llvm/lib/CMakeLists.txt:LLVM 支持库源码与 CMake 入口（裁剪快照中已扁平化到 llvm/lib/）" \
  "clang/lib/CMakeLists.txt:clang 组件库 CMake 入口" \
  "clang/tools/libclang:libclang C API 实现"
do
  relative="${required%%:*}"
  reason="${required#*:}"
  if [ ! -e "${SRC}/${relative}" ]; then
    MISSING+=("${relative}（${reason}）")
  fi
done
if [ "${#MISSING[@]}" -gt 0 ]; then
  echo -e "${RED}[FAIL] 构建前置依赖缺失，无法产出真实 libclang-min.a：${NC}" >&2
  for item in "${MISSING[@]}"; do
    echo "  - 缺失：${item}" >&2
  done
  echo "" >&2
  echo "  原因：入库的 llvm-project 是裁剪快照，只保留 clang/lib、clang/include、" >&2
  echo "        llvm/include、llvm/lib/Support；上游构建系统与 tablegen 未随库入库，" >&2
  echo "        clang 的 AST/Sema 又强依赖 tablegen 生成的 .inc 头文件。" >&2
  echo "" >&2
  echo "  补齐方式（任选其一）：" >&2
  echo "   1) 放入上游 llvm-project 18.1.3 完整源码树（需含 llvm/cmake、llvm/utils/TableGen、" >&2
  echo "      clang/tools/libclang），再重新执行本脚本；" >&2
  echo "   2) 改用发行版 libclang（apt install libclang-dev），并在本目录登记其头文件与库路径。" >&2
  echo "" >&2
  echo "  本脚本不会产出占位文件或伪造体积数字。" >&2
  {
    echo "libclang-min 构建报告"
    echo "结论：未构建（构建前置依赖缺失）"
    echo "缺失项："
    for item in "${MISSING[@]}"; do echo "  - ${item}"; done
    echo "源码目录：${SRC}"
  } >"${REPORT_FILE}"
  echo "  诊断报告：${REPORT_FILE}" >&2
  exit 3
fi
pass "构建前置依赖齐备"

# ------------------------------------------------------------- 3. 真实构建
banner "3/4 CMake 构建（剥离 driver/CodeGen/后端）"
mkdir -p "${BUILD_DIR}" "${ARTIFACT_DIR}"
# LLVM_SRC 环境变量允许指向完整上游源码树（如 apt deb-src 取得的
# llvm-toolchain-18 完整源码）；默认使用仓库内裁剪树。裁剪树若因缺少
# Target 后端 / tools / runtimes 目录而无法 configure，本脚本会如实
# 打印原因并退出，绝不伪造构建产物。
BUILD_SRC="${LLVM_SRC:+${LLVM_SRC}/llvm}"
BUILD_SRC="${BUILD_SRC:-${SRC}/llvm}"
if [ -n "${LLVM_SRC:-}" ]; then
  echo "  源码树：${BUILD_SRC}（LLVM_SRC 覆盖）"
else
  echo "  源码树：${BUILD_SRC}（仓库内裁剪树；可用 LLVM_SRC=<完整树路径> 覆盖）"
fi
{
  echo "=== configure ${GENERATOR} / Release ==="
  cmake -G "${GENERATOR}" -S "${BUILD_SRC}" -B "${BUILD_DIR}" \
    -DCMAKE_BUILD_TYPE=Release \
    -DLLVM_ENABLE_PROJECTS=clang \
    -DLLVM_TARGETS_TO_BUILD=Native \
    -DLLVM_ENABLE_ASSERTIONS=OFF \
    -DLLVM_INCLUDE_TESTS=OFF \
    -DLLVM_INCLUDE_EXAMPLES=OFF \
    -DLLVM_INCLUDE_BENCHMARKS=OFF \
    -DLLVM_INCLUDE_DOCS=OFF \
    -DLLVM_BUILD_TOOLS=OFF \
    -DLLVM_BUILD_UTILS=OFF \
    -DLLVM_ENABLE_ZLIB=OFF \
    -DLLVM_ENABLE_ZSTD=OFF \
    -DLLVM_ENABLE_TERMINFO=OFF \
    -DLLVM_ENABLE_LIBXML2=OFF \
    -DCLANG_ENABLE_STATIC_ANALYZER=OFF \
    -DCLANG_ENABLE_ARCMT=OFF \
    -DLLVM_ENABLE_RUNTIMES="" \
    -DBUILD_SHARED_LIBS=OFF
  echo "=== build libclang ==="
  cmake --build "${BUILD_DIR}" --target libclang
} 2>&1 | tee "${LOG_FILE}"

BUILT_LIB="$(find "${BUILD_DIR}" -name 'libclang.a' -type f | head -n1)"
[ -n "$BUILT_LIB" ] || fail "构建结束但未找到 libclang.a，详见构建日志 ${LOG_FILE}" 1
cp "$BUILT_LIB" "$OUT_LIB"
pass "静态库已生成：${OUT_LIB}"

# ------------------------------------------------------- 4. 体积核算与报告
banner "4/4 体积核算"
SIZE_BYTES="$(wc -c <"${OUT_LIB}")"
SIZE_MB="$(awk -v bytes="$SIZE_BYTES" 'BEGIN { printf "%.2f", bytes / 1048576 }')"
echo "  目标：libclang-min.a ≤ ${MAX_SIZE_MB}MB"
echo "  实际：${SIZE_MB}MB（${SIZE_BYTES} 字节）"
{
  echo "libclang-min 构建报告"
  echo "源码目录：${SRC}"
  echo "生成器：${GENERATOR}"
  echo "产物：${OUT_LIB}"
  echo "体积：${SIZE_MB}MB（${SIZE_BYTES} 字节）"
  echo "上限：${MAX_SIZE_MB}MB"
  echo "构建日志：${LOG_FILE}"
} >"${REPORT_FILE}"
if awk -v size="$SIZE_MB" -v max="$MAX_SIZE_MB" 'BEGIN { exit !(size > max) }'; then
  fail "体积 ${SIZE_MB}MB 超过上限 ${MAX_SIZE_MB}MB，详见 ${REPORT_FILE}" 1
fi
pass "体积达标（${SIZE_MB}MB ≤ ${MAX_SIZE_MB}MB）"
echo "  构建日志：${LOG_FILE}"
echo "  构建报告：${REPORT_FILE}"
banner "LIBCLANG-MIN BUILD: PASSED"
