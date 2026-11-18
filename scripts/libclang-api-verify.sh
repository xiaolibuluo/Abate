#!/usr/bin/env bash
# =============================================================================
# Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
# Abate is dual-licensed under AGPLv3 and a commercial license.
# For commercial licensing, contact xiaolibuluo@163.com.
# Source: https://github.com/xiaolibuluo/Abate
# =============================================================================
# Abate libclang C API 功能级验证（对应 pending-work #1 铁律口径）
#
# 真实调用自构建的 libclang-18.so.18 的 CXCursor C API：解析一个真实 C 源文件、
# 遍历 AST、断言具体声明被恢复（struct Point / 函数 add、main）。三重防造假：
#   1) ldd 断言二进制链接的是自研 .so（.artifacts 路径），不是系统 libclang；
#   2) 断言 diagnostics=0、cursor 数与具体声明符号真实回收，不用空跑充数；
#   3) C 源码由本脚本内联生成到临时目录，脚本自包含、任何环境可复现。
#
# 全程离线，不联网、不调用 clang driver、不读第三方语料。
# 用法：bash scripts/libclang-api-verify.sh
# =============================================================================
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
cd "$(dirname "$0")/.." || exit 2

banner() { echo -e "${YELLOW}== $1 ==${NC}"; }
fail() { echo -e "${RED}[FAIL] $1${NC}" >&2; exit "${2:-1}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }

banner "Abate libclang C API 功能级验证（CXCursor 真实遍历）"

# ---------------------------------------------------------- 1. 前置资产检查
banner "1/4 前置资产检查"
ARTIFACT_DIR=".artifacts/libclang-min"
SO="${ARTIFACT_DIR}/build/lib/libclang-18.so.18"
CACHE="${ARTIFACT_DIR}/build/CMakeCache.txt"
RESOURCE_DIR="${ARTIFACT_DIR}/build/lib/clang/18"
[ -f "$SO" ] || fail "自研 libclang .so 不存在：${SO}（先跑构建）" 3
command -v gcc >/dev/null 2>&1 || fail "缺少 gcc，无法编译 C API 探针" 3
command -v ldd >/dev/null 2>&1 || fail "缺少 ldd，无法验证链接来源" 3
[ -d "$RESOURCE_DIR/include" ] || fail "自研 clang 资源目录缺失：${RESOURCE_DIR}/include" 3

# 头文件从 CMakeCache 的源码路径动态推导，不硬编码 /tmp 绝对路径。
[ -f "$CACHE" ] || fail "缺少 CMakeCache.txt，无法推导 clang 头文件路径" 3
LLVM_HOME="$(sed -n 's/^CMAKE_HOME_DIRECTORY:INTERNAL=//p' "$CACHE" | head -n1)"
[ -n "$LLVM_HOME" ] || fail "CMakeCache 中无 CMAKE_HOME_DIRECTORY" 3
CLANG_INC="$(dirname "$LLVM_HOME")/clang/include"
[ -f "$CLANG_INC/clang-c/Index.h" ] || \
  fail "clang-c/Index.h 不存在：${CLANG_INC}（LLVM 源码树可能已清理）" 3
echo "  自研库:   ${SO}"
echo "  头文件:   ${CLANG_INC}/clang-c/Index.h"
echo "  资源目录: ${RESOURCE_DIR}"

# ---------------------------------------------------------- 2. 内联探针源码
banner "2/4 生成自包含探针（临时目录）"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

cat > "${WORK}/capi_probe.c" <<'CPROBE'
/* libclang C API functional probe: parse a real C file and walk its AST. */
#include <clang-c/Index.h>
#include <stdio.h>
#include <string.h>

#ifndef ABATE_RESOURCE_DIR
#define ABATE_RESOURCE_DIR "."
#endif

struct Probe { int total; int struct_point; int funcs; };

static enum CXChildVisitResult visit(CXCursor c, CXCursor parent, CXClientData d) {
	struct Probe *p = (struct Probe *)d;
	enum CXCursorKind kind = clang_getCursorKind(c);
	CXString name = clang_getCursorSpelling(c);
	const char *s = clang_getCString(name);
	(void)parent;
	if (kind == CXCursor_StructDecl && s && strcmp(s, "Point") == 0) p->struct_point++;
	if (kind == CXCursor_FunctionDecl && s &&
	    (strcmp(s, "add") == 0 || strcmp(s, "main") == 0)) p->funcs++;
	p->total++;
	clang_disposeString(name);
	return CXChildVisit_Recurse;
}

int main(int argc, char **argv) {
	const char *src = (argc > 1) ? argv[1] : "capi_sample.c";
	struct Probe probe = {0, 0, 0};
	CXIndex index; CXTranslationUnit tu;
	const char *args[] = {"-x", "c", "-resource-dir", ABATE_RESOURCE_DIR};
	enum CXErrorCode err;
	index = clang_createIndex(0, 0);
	if (!index) { fprintf(stderr, "FAIL: clang_createIndex\n"); return 2; }
	err = clang_parseTranslationUnit2(index, src, args, 4, NULL, 0,
	                                  CXTranslationUnit_None, &tu);
	if (err != CXError_Success || !tu) {
		fprintf(stderr, "FAIL: clang_parseTranslationUnit2 err=%d\n", (int)err);
		clang_disposeIndex(index); return 3;
	}
	printf("diagnostics=%u\n", clang_getNumDiagnostics(tu));
	clang_visitChildren(clang_getTranslationUnitCursor(tu), visit, &probe);
	printf("total_cursors=%d struct_Point=%d funcs_add_main=%d\n",
	       probe.total, probe.struct_point, probe.funcs);
	if (!(probe.total > 0 && probe.struct_point == 1 && probe.funcs == 2)) {
		fprintf(stderr, "FAIL: unexpected AST recovery\n");
		clang_disposeTranslationUnit(tu); clang_disposeIndex(index); return 4;
	}
	printf("RESULT: PASS\n");
	clang_disposeTranslationUnit(tu); clang_disposeIndex(index);
	return 0;
}
CPROBE

cat > "${WORK}/capi_sample.c" <<'CSAMPLE'
struct Point { int x; int y; };
int add(int a, int b) { return a + b; }
int main(void) { struct Point p; p.x = 1; return add(p.x, 2); }
CSAMPLE
pass "探针源码已生成（${WORK}）"

# ---------------------------------------------------------- 3. 编译 + 链接来源断言
banner "3/4 编译并断言链接自研库"
SO_ABS="$(cd "$(dirname "$SO")" && pwd)/$(basename "$SO")"
gcc -DABATE_RESOURCE_DIR="\"$(pwd)/${RESOURCE_DIR}\"" \
    -I"$CLANG_INC" "${WORK}/capi_probe.c" "$SO_ABS" \
    -Wl,-rpath,"$(dirname "$SO_ABS")" -o "${WORK}/capi_probe" || \
  fail "探针编译失败" 4

LINKED="$(ldd "${WORK}/capi_probe" | grep 'libclang-18' | awk '{print $3}')"
echo "  ldd libclang-18 => ${LINKED}"
case "$LINKED" in
  "$SO_ABS") pass "链接来源确认：自研 .so（非系统库）" ;;
  *) fail "链接到非自研 libclang：${LINKED}（期望 ${SO_ABS}）" 5 ;;
esac

# ---------------------------------------------------------- 4. 真实执行断言
banner "4/4 真实执行 C API 功能验证"
OUTPUT="$(cd "$WORK" && ./capi_probe capi_sample.c)"; RC=$?
echo "$OUTPUT" | sed 's/^/  /'
[ "$RC" -eq 0 ] || fail "C API 探针执行失败（rc=${RC}）" "$RC"
echo "$OUTPUT" | grep -q '^diagnostics=0$' || fail "diagnostics 非 0" 6
echo "$OUTPUT" | grep -q '^RESULT: PASS$' || fail "探针未返回 RESULT: PASS" 6

banner "LIBCLANG C API VERIFY: PASSED（自研库真实解析 + AST 遍历 + 声明回收）"
