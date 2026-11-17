/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command corpusgen is the deterministic, zero-dependency entry point for
// Abate's fully self-researched E2E corpus. It never edits the frozen plan,
// thresholds or gate scripts.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"

	"abate/internal/corpusgen"
)

const (
	exitOK    = 0
	exitFail  = 1
	exitUsage = 2
	// The defaultCorpusFileCount constant is the design's total file count.
	defaultCorpusFileCount = 20520
	defaultStrataCount     = 8
)

// options is the complete CLI configuration. Configuration travels in this
// struct so no generation function takes a boolean flag parameter.
type options struct {
	planPath   string
	root       string
	brokenRoot string
	count      int
	strata     int
	verifyOnly bool
}

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

// run parses flags, dispatches generation or verification and converts any
// error into an exit code. Nothing is skipped on failure.
func run(args []string, stdout, stderr io.Writer) int {
	opts := options{}
	fs := flag.NewFlagSet("corpusgen", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.StringVar(&opts.planPath, "plan", "ci/e2e_plan.json", "read-only frozen E2E plan path")
	fs.StringVar(&opts.root, "root", "testdata/corpus", "self-researched corpus output root")
	fs.StringVar(&opts.brokenRoot, "broken-root", "testdata/broken", "QA-only broken fixture root")
	fs.IntVar(&opts.count, "count", defaultCorpusFileCount, "total generated source files")
	fs.IntVar(&opts.strata, "strata", defaultStrataCount, "byte strata count")
	fs.BoolVar(&opts.verifyOnly, "verify-only", false, "only re-check existing generated corpus")
	if err := fs.Parse(args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitOK
		}
		return exitUsage
	}
	if opts.count < 1 || opts.strata < 1 {
		_, _ = fmt.Fprintln(stderr, "corpusgen: count and strata must be positive")
		return exitUsage
	}
	genOpts := corpusgen.GenOptions{
		PlanPath: opts.planPath, Root: opts.root, BrokenRoot: opts.brokenRoot,
		Count: opts.count, Strata: opts.strata,
	}
	if opts.verifyOnly {
		stats, err := corpusgen.Verify(genOpts)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "corpusgen verify: %v\n", err)
			return exitFail
		}
		printVerify(stdout, stats)
		return exitOK
	}
	generated, err := corpusgen.Generate(genOpts)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "corpusgen generate: %v\n", err)
		return exitFail
	}
	stats, err := corpusgen.Verify(genOpts)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "corpusgen verify: %v\n", err)
		return exitFail
	}
	printGenerate(stdout, generated, stats)
	return exitOK
}

// printGenerate writes a deterministic, fully sorted summary.
func printGenerate(w io.Writer, generated corpusgen.GenerateStats,
	stats corpusgen.VerifyStats) {
	_, _ = fmt.Fprintf(w, "Generated=%d\n", generated.Total)
	_, _ = fmt.Fprintf(w, "Bytes=%d\n", generated.Bytes)
	_, _ = fmt.Fprintf(w, "Broken=%d\n", generated.Broken)
	_, _ = fmt.Fprintf(w, "VerifiedFiles=%d\n", stats.Files)
	_, _ = fmt.Fprintf(w, "DistinctSHA256=%d\n", stats.DistinctSHA)
	_, _ = fmt.Fprintf(w, "DedicationSHA256=%s\n", stats.Dedication)
	for _, key := range sortedKeys(generated.PerLanguage) {
		_, _ = fmt.Fprintf(w, "Language[%s]=%d\n", key, generated.PerLanguage[key])
	}
	for _, key := range sortedKeys(generated.PerStratum) {
		_, _ = fmt.Fprintf(w, "Stratum[%s]=%d\n", key, generated.PerStratum[key])
	}
}

// printVerify writes the verification-only summary.
func printVerify(w io.Writer, stats corpusgen.VerifyStats) {
	_, _ = fmt.Fprintf(w, "VerifiedFiles=%d\n", stats.Files)
	_, _ = fmt.Fprintf(w, "DistinctSHA256=%d\n", stats.DistinctSHA)
	_, _ = fmt.Fprintf(w, "DedicationSHA256=%s\n", stats.Dedication)
	for _, key := range sortedKeys(stats.PerLanguage) {
		_, _ = fmt.Fprintf(w, "Language[%s]=%d\n", key, stats.PerLanguage[key])
	}
	for _, key := range sortedKeys(stats.PerStratum) {
		_, _ = fmt.Fprintf(w, "Stratum[%s]=%d\n", key, stats.PerStratum[key])
	}
}

// sortedKeys returns map keys in deterministic lexical order.
func sortedKeys(values map[string]int) []string {
	out := make([]string, 0, len(values))
	for key := range values {
		out = append(out, key)
	}
	slices.Sort(out)
	return out
}
