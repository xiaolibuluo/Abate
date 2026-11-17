/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command e2e runs the deterministic seven-language end-to-end measurement.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"abate/internal/corpus"
	"abate/internal/e2e"
	"abate/internal/e2e/report"
)

// Exit codes are fixed so the tenth gate can interpret them without parsing.
const (
	exitOK    = 0
	exitFail  = 1
	exitUsage = 2
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

// run accepts exactly the three flags fixed by ci/gate_e2e.sh: -plan,
// -manifest and -out. -out is the report file path; its directory also holds
// e2e_trace.jsonl and e2e_drilldown.jsonl.
func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("e2e", flag.ContinueOnError)
	flags.SetOutput(stderr)
	planPath := flags.String("plan", "", "path to ci/e2e_plan.json")
	manifestPath := flags.String("manifest", "", "path to ci/corpus_manifest.json")
	outPath := flags.String("out", "", "path to e2e_report.json")
	if err := flags.Parse(args[1:]); err != nil {
		return exitUsage
	}
	if *planPath == "" || *manifestPath == "" || *outPath == "" {
		emit(stderr, "usage: e2e -plan <file> -manifest <file> -out <file>\n")
		return exitUsage
	}
	plan, err := e2e.LoadPlan(*planPath)
	if err != nil {
		emit(stderr, "e2e: plan: %v\n", err)
		return exitFail
	}
	manifest, err := corpus.Load(*manifestPath)
	if err != nil {
		emit(stderr, "e2e: manifest: %v\n", err)
		return exitFail
	}
	dir := filepath.Dir(*outPath)
	if err = os.MkdirAll(dir, e2e.DirPerm); err != nil {
		emit(stderr, "e2e: output directory: %v\n", err)
		return exitFail
	}
	current, err := report.LoadThresholds(plan.ThresholdsPath)
	if err != nil {
		emit(stderr, "e2e: thresholds: %v\n", err)
		return exitFail
	}
	baseline, err := report.LoadThresholds(plan.BaselinePath)
	if err != nil {
		emit(stderr, "e2e: baseline thresholds: %v\n", err)
		return exitFail
	}
	trace := e2e.NewTrace(filepath.Join(dir, e2e.TraceFileName))
	runner, err := e2e.NewRunner(plan, manifest, trace)
	if err != nil {
		emit(stderr, "e2e: runner: %v\n", err)
		return exitFail
	}
	rep, records, err := runner.Run()
	if err != nil {
		emit(stderr, "e2e: run: %v\n", err)
		return exitFail
	}
	if err = report.Write(*outPath, rep, records); err != nil {
		emit(stderr, "e2e: write: %v\n", err)
		return exitFail
	}
	// Recompute rebuilds every aggregate and baseline value from the drilldown
	// record stream and asserts an exact, zero-difference match with the report.
	// Any discrepancy is an integrity failure, never an assumption.
	if _, err = report.Recompute(rep, records); err != nil {
		emit(stderr, "e2e: recompute: %v\n", err)
		return exitFail
	}
	// Evaluate applies the tighten-only invariant and the zero-tolerance gate;
	// its verdict is the single pass/fail signal the tenth gate consumes.
	verdict := report.Evaluate(current, baseline, rep)
	if err = report.WriteVerdict(*outPath, verdict); err != nil {
		emit(stderr, "e2e: verdict: %v\n", err)
		return exitFail
	}
	emit(stdout,
		"E2E_OK languages=%d records=%d processed=%d registered=%d planSha256=%s manifestSha256=%s thresholdsSha256=%s verdictPassed=%v failures=%d\n",
		len(rep.Languages), len(records), rep.ProcessedFiles,
		rep.RegisteredFiles, rep.PlanSHA256, rep.ManifestSHA256,
		rep.ThresholdsSHA256, verdict.Passed, len(verdict.Failures))
	if !verdict.Passed {
		for _, f := range verdict.Failures {
			emit(stderr, "E2E_VERDICT_FAIL %s\n", f.String())
		}
		return exitFail
	}
	return exitOK
}

// emit writes one line with an explicitly discarded write result.
func emit(w io.Writer, format string, args ...any) {
	_, _ = io.WriteString(w, fmt.Sprintf(format, args...))
}
