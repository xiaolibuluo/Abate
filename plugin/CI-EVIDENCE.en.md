<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](CI-EVIDENCE.en.md) | [中文](CI-EVIDENCE.md)

# Abate M4 Plugin Ecosystem CI Evidence (commit → publish ≤ 1h)

English translation of `CI-EVIDENCE.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; the title `commit → publish ≤ 1h` is translated as-is
(a process/SLA description, not upgraded into a capability claim); evidence
statements are kept as-is and are consistent with the honest boundaries in
§6b/§6d of the handoff; no new conclusions are added; this document does not
constitute legal advice.

This file is the acceptance-evidence index for the M4 plugin ecosystem; it
lives under `plugin/` and does not modify anything under `docs/`. All evidence
is reproducible with in-repository commands: no external services, no network,
no third-party dependencies.

## 1. Agent-evidence statement (honest boundary)

This sandbox has no real GitHub Actions runner, so the "CI evidence" has two
layers:

- **Declaration layer**: `.github/workflows/plugin-publish.yml` is a real,
  executable GitHub Actions workflow (triggered on `push` to `main`; runs the
  9-segment strict gate first, then build + sign + artifact upload).
- **Execution layer**: `scripts/plugin-publish.sh` fully simulates every step
  of the same pipeline locally (build, race tests, CLI tests, signing,
  artifact load/list/call verification, timestamp recording) and generates
  `.artifacts/plugin-publish/PUBLISH-EVIDENCE.txt`.

The workflow and the script share the same step set and the same 3600-second
upper bound, so the local execution layer covers all behavior of build / race
tests / CLI tests / signing / artifact verification / elapsed-time assertion.
**But one CI-only step cannot really run in this sandbox**: the workflow's
`Install golangci-lint v2` step fetches a version-pinned (v2.1.6) tool binary
over the network at build time; this sandbox has neither a runner nor that
network condition. That step was only statically reviewed (YAML parses with
`python3 -c "import yaml"`; the step-local `env:` overrides were checked) and
**is not claimed to have really executed locally**. The gate / publish / upload
steps remain fully offline with `GOPROXY=off` + `GOSUMDB=off` +
`GOFLAGS=-mod=vendor`, matching the local script. On a real runner,
`startedAt/finishedAt` are rewritten into the evidence file on every CI run.

## 2. commit → publish SLA

| Constraint | Location |
| --- | --- |
| Job hard timeout 60 minutes | `timeout-minutes: 60` in `.github/workflows/plugin-publish.yml` |
| Gate first | the workflow's `Strict gate (9 segments)` step calls `bash scripts/lint-gate.sh` |
| Bounded elapsed-time assertion | `scripts/plugin-publish.sh` computes `elapsedSeconds` and fails beyond 3600 seconds |
| Timestamp evidence | the evidence file records `startedAt`, `finishedAt`, `elapsedSeconds`, `status=PASS` |
| Reproduction command | `bash scripts/plugin-publish.sh` (repeatable; accepts an absolute-path output directory) |

The pipeline steps are fixed and run entirely locally:
`go build ./...` → `go test -race ./plugin/...` → `go test ./ctl/cli/...`
→ `go run ./internal/pluginpublish` (generate keys, sign, self-verify, write
the manifest envelope)
→ `go build -o abate .` → `abate plugin load/list/call` to verify the signed
artifacts → record elapsed time. The private key exists only in-process and is
never written to disk; the envelope carries only the public key, which the CLI
registers at load time and verifies signatures against. The signer generates a
new public key per run via `GenerateKey()`, so a differing `manifestSha256`
per run is the expected semantics (independent key per release) and is not
part of deterministic-replay comparison.

The CLI's `load/list/call/hotswap` are separate processes; the session is
explicitly rebuilt from the fixed `.abate/plugin-state.json` state file under
the working directory: each command re-reads all signed envelopes and re-runs
`Host.Load` signature verification, so it works across processes and never
admits unsigned plugins into the registry.

### 2.1 Fix for M4 QA round 1 FAIL (session-state replay / absolute paths / toolchain bootstrap)

QA re-verification found that the old script always failed on a second run
(exit code 3): the publish script ran the CLI in the artifact directory, whose
session state was anchored to that directory's `.abate/plugin-state.json`;
on a second run, `loadPluginSession` replayed the previous run's old envelope
and registered the old public key first, while the signer generated a new key
and the anchor name is fixed (`plugin-release-anchor`), so
`registerPluginAnchor` detected "same anchor, different fingerprint" and
refused to load. Worse, the failure happened before the evidence file was
written, so a stale `status=PASS` evidence file lingered in the directory and
the workflow would upload the whole directory — the risk of "uploading
historical passing evidence after a failure".

Fixes (all in `scripts/plugin-publish.sh`; **the anchor name in
`internal/pluginpublish` was not changed, no private key persisted, and the
"independent key per release" semantics was not broken**):

1. `rm -rf "$OUT_DIR/.abate"` at the start of every run, making
   "second run == first run" and eliminating the replay conflict at its root;
2. Also delete the old `PUBLISH-EVIDENCE.txt` and `load.json` / `list.json` /
   `call.json`, and place the 3600-second SLA assertion before the evidence
   write, so any mid-way failure leaves no stale `status=PASS` evidence; the
   evidence file is written only after the full flow succeeds;
3. `OUT_DIR` is `mkdir -p`ed and then normalized via `cd` into a canonical
   absolute path; from then on only `"$OUT_DIR/abate"` /
   `"$OUT_DIR/plugin-manifest.json"` are used, so absolute-path arguments are
   no longer concatenated into `/workspace/abate//tmp/.../abate` (the old
   defect's exit code 127);
4. The script probes the toolchain via `go version`: the caller's environment
   is kept only if it is exactly `go1.26.11`; otherwise it explicitly switches
   to the version-pinned 1.26.11 toolchain and resets `PATH`. As a result,
   even `env -u GOROOT` bare runs (where the old script inherited a stale
   preinstalled go and failed at step one with
   `go.mod requires go >= 1.26.11`) pass fully green; and when the caller has
   already exported the correct environment, it is left untouched.

Reproducible commands and locally measured results after the fix (offline
`-mod=vendor`, `GOPROXY=off`):

| Scenario | Command | Exit code | Evidence `status` |
| --- | --- | --- | --- |
| First run, default directory | `bash scripts/plugin-publish.sh` | 0 | PASS (`elapsedSeconds=4`) |
| Second run, same directory | `bash scripts/plugin-publish.sh` | 0 | PASS (`elapsedSeconds=3`) |
| Fresh absolute-path directory | `bash scripts/plugin-publish.sh /tmp/abate-pub-abs` | 0 | PASS (`elapsedSeconds=3`) |
| Bare bootstrap | `env -u GOROOT bash scripts/plugin-publish.sh` | 0 | PASS (`elapsedSeconds=4`) |
| Fully clean environment bootstrap | `env -u GOROOT -u GOFLAGS -u GOPROXY -u GOSUMDB PATH=/usr/local/go/bin:/usr/bin:/bin bash scripts/plugin-publish.sh` | 0 | PASS (`elapsedSeconds=3`) |

Other repository-wide verification: `go build ./...`, `go vet ./...`,
`go test -race -count=1 ./plugin/... ./ctl/cli/...` all exit 0; the report
determinism invariant
`go run . run ctl/cli/testdata/ok.yaml --seed 42 --format json | sha256sum`
remains
`c00c280272c7611da4a79de07f7223f7ab76bfe6db5e97b0c44ea899ef9fea57`.
The old wording "an independent, reproducible session state allocated per
run" did not hold under the old script (it only `mkdir -p`ed without
cleaning, so a second run necessarily hit the replay conflict); it has been
replaced by the two measured scenarios in the table above.

### 2.2 Lint tool installation: build-time local network, runtime offline

The global `env:` keeps `GOPROXY: "off"` / `GOSUMDB: "off"` /
`GOFLAGS: "-mod=vendor"`. The old workflow placed
`go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6`
under that global offline env and really failed with exit code 1
(`cannot query module due to -mod=vendor`), while `scripts/lint-gate.sh:95`
depends on that binary → the declaration-layer gate could not complete
offline.

Fix: the lint tool install was split into an **independent networked step**
that only within that step uses step-local `env:` overrides
`GOPROXY=https://proxy.golang.org,direct`, `GOSUMDB=sum.golang.org`,
`GOFLAGS=""` to fetch the version-pinned build-time tool binary and add
`$GOPATH/bin` to `PATH`; the remaining gate / publish / upload steps continue
to inherit the global offline env. This is local networking for the
**build-time toolchain**, not an Abate runtime API integration: at runtime
there is zero network, zero third-party imports, zero CGO, zero
subprocesses; the lint binary is neither vendored into the repository nor
added to `go.mod`, and the gate does not skip lint. The `timeout-minutes: 60`
SLA hard boundary and the evidence-upload step are retained.

Static-review conclusion: the YAML parses with
`python3 -c "import yaml"`; `timeout-minutes: 60`, the global offline env,
the install step's local `env:` override and version pinning, no local `env:`
on the gate/publish steps, and the presence of `actions/upload-artifact@v4`
were all checked. **This sandbox cannot really run a GitHub runner, so this
step is only a static correctness proof and is not claimed to have executed
in real CI.**

## 3. Evidence for the four M4 acceptance criteria

### (1) Unsigned plugins rejected / capability escalation intercepted

- `plugin/host/host_test.go::TestLoadUnsignedManifestRejected`: `Host.Load`
  of an unsigned manifest returns `ErrSignatureInvalid` and does not reach
  the factory.
- `plugin/host/host_test.go::TestLoadUnknownCapabilityRejected`,
  `TestCallMissingCapabilityDenied`: unknown capability declarations fail
  closed; calls missing a capability return `ErrCapabilityDenied`.
- `plugin/host/host_entries_test.go`: reporter/adapter entries are
  mandatory-`CapObserve` gated; even an empty want list cannot bypass.
- `ctl/cli/cmd_plugin_test.go::TestPluginRejectsUnsignedManifestWithoutRegistration`:
  the CLI end-to-end rejects an unsigned manifest (exit code 1) and
  `plugin list` confirms nothing was registered.
- `plugin/sample/e2e_demo_test.go`: with all three plugins loaded, signature
  rejection and side-effect absence are verified again.

### (2) Hot-swap switching without interrupting experiments

- `plugin/hotswap/` implements the four-phase state machine
  `PREPARE → SWITCH → DRAIN → UNLOAD`; `Allowed` is fail-closed: major must
  match and versions must never regress.
- `plugin/hotswap/hotswap_test.go::TestInFlightCallDrainsAndNewInstanceServes`:
  switching while the old instance holds an in-flight `BeginCall`; verifies
  DRAIN waits for it, the new instance serves during DRAIN, and the old
  instance rejects new calls (`ErrInstanceStale`).
- `plugin/hotswap/hotswap_chaos_test.go`: 32 independent Hosts switching
  concurrently, 24 concurrent in-flight calls, and 8 chained consecutive
  upgrades, all passing under `go test -race`.
- `ctl/cli/cmd_plugin_test.go::TestPluginHotswapEndToEnd`: the CLI loads
  v1.0.0 → hotswaps to v1.0.1 → state `UNLOAD` → `list` shows the new version
  → the new version continues to `call apply`; rollback to v1.0.0 is refused.

### (3) commit → publish ≤ 1h

See section 2: the workflow hard timeout, the script's elapsed-time
assertion, the timestamp evidence file, and the SLA noted in comments in
`.github/workflows/plugin-publish.yml`; section 2.1 gives the measured
reproduction commands and exit codes for the second-run and absolute-path
scenarios.

### (4) RPT-SARIF passes the official validator (OASIS schema)

- The official OASIS SARIF 2.1.0 schema is stored as a data file at
  `third_party/_references/sarif-2.1.0.schema.json` (not compiled).
- `plugin/report/schema/` is an in-house, offline draft-07 subset
  JSON-Schema validator implementing `type`, `properties`, `required`,
  `enum`, `items`, `additionalProperties`, `$ref`, `oneOf`/`anyOf`/`allOf`,
  `pattern`, `minimum`/`maximum`, `format`, with no third-party dependencies.
- `plugin/report/report_test.go::TestBuildSARIFValidatesAgainstOASIS` and
  `TestBuildSARIFRequiredFields`: the generated SARIF validates against the
  OASIS schema with zero errors and carries `version:"2.1.0"`, an absolute
  schema URI, and at least one run.
- `plugin/report/report_mutation_test.go::TestOASISRejectsInvalidDocument`
  and `plugin/report/schema/schema_mutation_test.go`: mutation tests prove
  the validator is not a no-op (missing required, extra properties, invalid
  enums, out-of-range numbers are all rejected).
- `ctl/cli/cmd_plugin_test.go::TestPluginReporterAndAdapterEntries` and
  `plugin/sample/e2e_demo_test.go`: SARIF produced by the CLI / end-to-end
  chain validates as well.

## 4. Determinism and red lines

The plugin runtime uses only the injected `sim/clock.VirtualClock` and
`sim/rand.Stream`; the `DrainTimeout` in `plugin/hotswap/timeout.go` is the
only real wall-clock boundary, used solely to wait for in-flight goroutines
to drain, and participates in no plugin output. Registry enumeration,
findings sorting, and schema-validation traversal are explicitly sorted; no
map iteration order leaks. The publish script's `date` timestamps go only
into the evidence file and participate in no plugin output or hashes, so the
`c00c2802…9fea57` report invariant is unaffected.
