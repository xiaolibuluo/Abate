#!/usr/bin/env bash
# =============================================================================
# Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
# Abate is dual-licensed under AGPLv3 and a commercial license.
# For commercial licensing, contact xiaolibuluo@163.com.
# Source: https://github.com/xiaolibuluo/Abate
# =============================================================================
# 第十段 E2E 闭环门禁。S1/S2/S3 完备性反查与零豁免静态检查真实可跑；
# S4 同 seed 双跑逐字节 cmp；S5 崩溃定位；S6 真跑 Recompute 零差 + Evaluate
# 判定 verdict.passed。无 skip、无白名单、无阈值放宽开关、无环境变量开关；
# 红灯即 exit 1。
# 归档目录名取 manifestSha256 前 12 位，确定且无时间戳。
# =============================================================================
set -uo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; NC='\033[0m'
FAILED=0
REPORT=""

cd "$(dirname "$0")/.." || exit 2

fail() { FAILED=1; REPORT="${REPORT}\n  [FAIL] $1"; echo -e "${RED}[FAIL] $1${NC}"; }
pass() { echo -e "${GREEN}[PASS] $1${NC}"; }
step() { echo -e "${YELLOW}>>> $1${NC}"; }

MANIFEST_SHA=""
OUT_DIR=""

step "S1 manifestSha256 复算 + corpusctl verify 全量重算 sha256 / SPDX"
if [ ! -f ci/corpus_manifest.json ] || [ ! -f ci/corpus_manifest.sha256 ]; then
  fail "S1: manifest 缺失，需先执行 go run ./cmd/corpusctl register"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
MANIFEST_SHA="$(sha256sum ci/corpus_manifest.json | cut -d' ' -f1)"
STORED_SHA="$(tr -d '\n' < ci/corpus_manifest.sha256)"
OUT_DIR=".artifacts/e2e/${MANIFEST_SHA:0:12}"
mkdir -p "$OUT_DIR"
if [ "$MANIFEST_SHA" != "$STORED_SHA" ]; then
  fail "S1: manifestSha256 不符 expect=${STORED_SHA} actual=${MANIFEST_SHA}"
else
  pass "S1: manifestSha256=${MANIFEST_SHA}"
fi
VERIFY_OUT="$(go run ./cmd/corpusctl verify 2>&1)"
VERIFY_RC=$?
printf '%s\n' "$VERIFY_OUT"
jq -n --arg sha "$MANIFEST_SHA" --arg raw "$VERIFY_OUT" \
  '{manifestSha256:$sha, verifyOutput:$raw}' > "$OUT_DIR/verify_result.json"
if [ "$VERIFY_RC" -ne 0 ]; then
  fail "S1: corpusctl verify exit=${VERIFY_RC}（全量重算或 SPDX 零容忍未通过）"
else
  pass "S1: SHA256Mismatches=0 / SPDXViolations=0，全量重算通过"
fi

step "S2 阈值只收紧单调校验（tightenOnly，零容忍键齐全恒 0）"
for cfg in ci/e2e_plan.json ci/e2e_thresholds.json ci/e2e_thresholds_baseline.json \
           ci/operator_kind_map.json ci/constructs_go.json ci/equivalent_mutations.json; do
  if ! jq -e . "$cfg" >/dev/null 2>&1; then
    fail "S2: 配置 JSON 无法解析 ${cfg}"
  fi
done
BASE_SHA="$(sha256sum ci/e2e_thresholds_baseline.json | cut -d' ' -f1)"
PLAN_SHA="$(jq -r '.thresholdsBaselineSha256' ci/e2e_plan.json)"
if [ "$BASE_SHA" != "$PLAN_SHA" ]; then
  fail "S2: plan.thresholdsBaselineSha256 不符 expect=${BASE_SHA} actual=${PLAN_SHA}"
fi
TIGHTEN="$(jq -r '.tightenOnly' ci/e2e_thresholds.json)"
if [ "$TIGHTEN" != "true" ]; then
  fail "S2: tightenOnly 必须为 true，实测 ${TIGHTEN}"
fi
Z_CUR="$(jq -r '.zeroTolerance | keys[]' ci/e2e_thresholds.json | LC_ALL=C sort)"
Z_BASE="$(jq -r '.zeroTolerance | keys[]' ci/e2e_thresholds_baseline.json | LC_ALL=C sort)"
if [ "$Z_CUR" != "$Z_BASE" ]; then
  fail "S2: 零容忍键集合与 baseline 不完全相同"
fi
Z_NONZERO="$(jq -r '.zeroTolerance | to_entries[] | select(.value != 0) | "\(.key)=\(.value)"' ci/e2e_thresholds.json)"
if [ -n "$Z_NONZERO" ]; then
  fail "S2: 零容忍项非 0 -> ${Z_NONZERO}"
fi
for section in minima maxima; do
  CUR_K="$(jq -r ".${section} | keys[]" ci/e2e_thresholds.json | LC_ALL=C sort)"
  BASE_K="$(jq -r ".${section} | keys[]" ci/e2e_thresholds_baseline.json | LC_ALL=C sort)"
  if [ "$CUR_K" != "$BASE_K" ]; then
    fail "S2: ${section} 键集合与 baseline 不一致（禁止缺键或新增未登记键）"
  fi
done
RELAXED="$(jq -n --slurpfile c ci/e2e_thresholds.json \
                  --slurpfile b ci/e2e_thresholds_baseline.json '
  [ $c[0].minima as $cur | $b[0].minima as $base
    | ($cur | to_entries[]) as $e | ($e.key) as $k | ($base[$k]) as $bv
    | if $bv == null then "UNREGISTERED_KEY:minima.\($k)"
      elif $e.value < $bv then "RELAXED:minima.\($k) current=\($e.value) baseline=\($bv) direction=must-increase"
      else empty end ]
  + [ $c[0].maxima as $cur | $b[0].maxima as $base
    | ($cur | to_entries[]) as $e | ($e.key) as $k | ($base[$k]) as $bv
    | if $bv == null then "UNREGISTERED_KEY:maxima.\($k)"
      elif $e.value > $bv then "RELAXED:maxima.\($k) current=\($e.value) baseline=\($bv) direction=must-decrease"
      else empty end ]
  | .[]' 2>&1)"
JQ_RC=$?
if [ "$JQ_RC" -ne 0 ]; then
  fail "S2: 单调校验 jq 执行失败 -> ${RELAXED}"
elif [ -n "$RELAXED" ]; then
  fail "S2: 阈值被放宽 -> ${RELAXED}"
else
  pass "S2: minima 只增 / maxima 只减，零容忍键齐全恒 0，baselineSha256=${BASE_SHA}"
fi
jq -n --arg base "$BASE_SHA" --arg plan "$PLAN_SHA" \
      --argjson tighten "$(jq -c '.tightenOnly' ci/e2e_thresholds.json)" \
  '{baselineSha256:$base, planBaselineSha256:$plan, tightenOnly:$tighten}' \
  > "$OUT_DIR/threshold_check.json"

step "S3 完备性反查（find 新增根 vs e2e_sources.txt）+ 零豁免静态检查"
FOUND="$(while IFS= read -r root; do
  [ -z "$root" ] && continue
  find "$root" -name '*.go' -type f
done < ci/e2e_new_roots.txt | LC_ALL=C sort)"
printf '%s\n' "$FOUND" > "$OUT_DIR/source_list.txt"
SOURCE_COUNT="$(wc -l < ci/e2e_sources.txt | tr -d ' ')"
if ! printf '%s\n' "$FOUND" | cmp -s - ci/e2e_sources.txt; then
  fail "S3: 完备性反查 FAIL —— find(新增根) 与 ci/e2e_sources.txt 不一致"
else
  pass "S3: 完备性反查通过（${SOURCE_COUNT} 个 .go 全部登记，无遗漏）"
fi
if [ ! -f cmd/neutralitycheck/main.go ]; then
  fail "S3: cmd/neutralitycheck/main.go 未就绪，零豁免 go/ast 静态检查不可执行"
else
  NEUTRALITY_OUT="$OUT_DIR/neutrality.json"
  go run ./cmd/neutralitycheck -sources ci/e2e_sources.txt -out "$NEUTRALITY_OUT"
  NC_RC=$?
  if [ "$NC_RC" -ne 0 ]; then
    fail "S3: neutralitycheck 退出码非零（${NC_RC}），零豁免 go/ast 静态检查未通过"
  else
    pass "S3: neutralitycheck 通过（registered==processed 且 findings=0）"
  fi
fi

step "S4 同 seed + 同 manifest 双跑 e2e_report.json 逐字节 cmp"
if [ ! -f cmd/e2e/main.go ]; then
  fail "S4: cmd/e2e/main.go 未就绪（T04），无法执行双跑；闭环红灯起点"
else
  RUN1="$OUT_DIR/run1"; RUN2="$OUT_DIR/run2"
  mkdir -p "$RUN1" "$RUN2"
  go run ./cmd/e2e -plan ci/e2e_plan.json -manifest ci/corpus_manifest.json -out "$RUN1/e2e_report.json"
  R1=$?
  go run ./cmd/e2e -plan ci/e2e_plan.json -manifest ci/corpus_manifest.json -out "$RUN2/e2e_report.json"
  R2=$?
  # cmd/e2e 在 verdict 判定失败时也会退出非零，但报告与 verdict 已落盘；
  # 仅当报告未产出时才视为基础设施故障，否则交由 S6 判定 verdict。
  if [ "$R1" -ne 0 ] && [ ! -f "$RUN1/e2e_report.json" ]; then
    fail "S4: e2e run1 退出非零且未产出报告（基础设施故障）run1=${R1}"
  fi
  if [ "$R2" -ne 0 ] && [ ! -f "$RUN2/e2e_report.json" ]; then
    fail "S4: e2e run2 退出非零且未产出报告（基础设施故障）run2=${R2}"
  fi
  if [ ! -f "$RUN1/e2e_report.json" ] || [ ! -f "$RUN2/e2e_report.json" ]; then
    : # 已按基础设施故障处理
  elif ! cmp -s "$RUN1/e2e_report.json" "$RUN2/e2e_report.json"; then
    OFFSET="$(cmp "$RUN1/e2e_report.json" "$RUN2/e2e_report.json" 2>&1 || true)"
    fail "S4: 双跑报告逐字节不一致 -> ${OFFSET}"
  else
    pass "S4: 同 seed 双跑逐字节一致（verdict 由 S6 判定）"
  fi
fi

step "S5 崩溃定位（零 recover，靠退出码 + 报告完整性 + Trace 末行）"
if [ ! -f "$OUT_DIR/run1/e2e_report.json" ] || [ ! -f "$OUT_DIR/run1/e2e_trace.jsonl" ]; then
  fail "S5: e2e_report.json / e2e_trace.jsonl 未产出（T04 未就绪），无法给出崩溃定位"
else
  tail -n 1 "$OUT_DIR/run1/e2e_trace.jsonl" > "$OUT_DIR/trace.tail"
  pass "S5: Trace 末行已归档 ${OUT_DIR}/trace.tail"
fi

step "S6 Recompute 重算零差 + Evaluate 生成 Verdict（passed 必须 true）"
VERDICT="$OUT_DIR/run1/e2e_verdict.json"
if [ ! -f "$VERDICT" ]; then
  fail "S6: e2e_verdict.json 未产出（cmd/e2e 未调用 Evaluate），无法判定"
elif ! jq -e '.passed' "$VERDICT" >/dev/null 2>&1; then
  fail "S6: Evaluate 判定 FAIL —— 见 e2e_verdict.json 的 failures 列表"
  jq -r '.failures[]? | "    [FAIL] \(.key) actual=\(.actual) expected=\(.expected) rule=\(.rule)"' \
    "$VERDICT" >&2 || true
else
  pass "S6: Evaluate 判定 PASSED（零容忍 + 单调 + 阈值全绿）"
fi

step "S7 归档（manifestSha256 前 12 位，目录名确定、无时间戳）"
if ! cp ci/corpus_manifest.json ci/corpus_manifest.sha256 ci/e2e_thresholds.json \
        ci/e2e_thresholds_baseline.json ci/e2e_plan.json "$OUT_DIR/"; then
  fail "S7: 归档 manifest / thresholds / plan 失败"
fi
printf '%s\n' "$FAILED" > "$OUT_DIR/exit_code"
pass "S7: 归档目录 ${OUT_DIR}（exit_code=${FAILED}）"

echo ""
if [ "$FAILED" -ne 0 ]; then
  echo -e "${RED}=== E2E GATE: FAILED（红灯起点，逐项定位如下）===${NC}"
  echo -e "${RED}${REPORT}${NC}"
  exit 1
fi
echo -e "${GREEN}=== E2E GATE: PASSED ===${NC}"
exit 0
