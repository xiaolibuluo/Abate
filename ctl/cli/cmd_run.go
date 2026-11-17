/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/ces/validate"
	"abate/sim/record"
)

// runOptions holds the parsed run flags.
type runOptions struct {
	cesPath    string
	seed       uint64
	format     string
	out        string
	recordPath string
	dryRun     bool
	strict     bool
}

// runOutcome carries everything needed to render a run report.
type runOutcome struct {
	opts       runOptions
	exp        *schema.Experiment
	violations []validate.Violation
	execution  *execution
	exitCode   int
}

// runHelp documents the run command and its copyable examples.
const runHelp = `Usage: abate run <ces-file> [flags]

Parse, normalize and validate a CES document, then drive a deterministic
kernel experiment and emit a report. --dry-run stops after validation.

Flags:
  --seed uint        Deterministic seed override (0 uses the document seed).
  --format console|json  Report format (default console).
  --out path          Also write the report bytes to a file.
  --record path       Write the deterministic SIM-REC archive to a file.
  --dry-run           Parse and validate only; do not execute.
  --strict            Promote warning-severity violations to failures.
  --help, -h           Print this help text.

` + exitCodeHelp + `
Examples:
  abate run ctl/cli/testdata/ok.yaml
  abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
  abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
  abate replay /tmp/run.arch
  abate run ctl/cli/testdata/ok.yaml --dry-run
  abate run ctl/cli/testdata/ok.yaml --strict --format json
`

// cmdRun drives the parse, validate, execute and report chain.
func cmdRun(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeRunHelp(stdout)
		return exitOK
	}
	opts, code := parseRunOptions(args, stderr)
	if code != exitOK {
		return code
	}
	data, err := os.ReadFile(opts.cesPath) //nolint:gosec // the CES path is supplied explicitly by the caller.
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate run: cannot read CES document %q: %v\n", opts.cesPath, err)
		return exitUsageError
	}
	parsed, err := parse.Parse(data, parse.FormatAuto)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate run: %v\n", err)
		return exitUsageError
	}
	normalized, err := parse.Normalize(parsed)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate run: %v\n", err)
		return exitUsageError
	}
	if opts.seed != 0 {
		normalized.Spec.Seed = opts.seed
	}
	vs, err := validate.Validate(normalized)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate run: %v\n", err)
		return exitInternalError
	}
	out := runOutcome{opts: opts, exp: normalized, violations: vs}
	// PLAN-07 §4 maps validation-phase violations to exit 3 because an invalid
	// CES document is a usage error, whereas a steady-state VIOLATED verdict
	// raised during execution is the business failure mapped to exit 1. Strict
	// mode promotes warning violations into this same usage-error branch.
	var blocked bool
	if opts.strict {
		blocked = blockingStrict(vs)
	} else {
		blocked = blocking(vs)
	}
	if blocked {
		out.exitCode = exitUsageError
		return emitRunReport(out, stdout, stderr)
	}
	if opts.dryRun {
		out.exitCode = exitOK
		return emitRunReport(out, stdout, stderr)
	}
	exec, err := runExperiment(normalized)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate run: %v\n", err)
		return exitInternalError
	}
	out.execution = &exec
	out.exitCode = exitOK
	if steadyViolated(exec.verdicts) {
		out.exitCode = exitBusinessFailure
	}
	if opts.recordPath != "" {
		if err := writeArchive(exec.archive, opts.recordPath); err != nil {
			_, _ = fmt.Fprintf(stderr, "abate run: %v\n", err)
			return exitInternalError
		}
	}
	return emitRunReport(out, stdout, stderr)
}

// writeArchive seals the recording of one run into a deterministic file.
func writeArchive(arch *record.Archive, path string) error {
	if arch == nil {
		return errors.New("record: no archive was produced")
	}
	var buf bytes.Buffer
	if err := arch.Save(&buf); err != nil {
		return fmt.Errorf("record: encode archive: %w", err)
	}
	if err := os.WriteFile(path, buf.Bytes(), 0o600); err != nil { //nolint:gosec // the path is supplied explicitly by the caller.
		return fmt.Errorf("record: write %q: %w", path, err)
	}
	return nil
}

// parseRunOptions parses and validates all run arguments once.
func parseRunOptions(args []string, stderr io.Writer) (runOptions, int) {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	seed := fs.Uint64("seed", 0, "deterministic seed override")
	format := fs.String("format", formatConsole, "console or json")
	out := fs.String("out", "", "report output path")
	recordPath := fs.String("record", "", "archive output path")
	dryRun := fs.Bool("dry-run", false, "parse and validate only")
	strict := fs.Bool("strict", false, "promote warnings to failures")
	ordered := reorderArgs(args, runValueFlags())
	if err := fs.Parse(ordered); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate run: %v\n", err)
		writeRunHelp(stderr)
		return runOptions{}, exitUsageError
	}
	if fs.NArg() != 1 {
		writeRunHelp(stderr)
		return runOptions{}, exitUsageError
	}
	if !validFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate run: unsupported format %q\n", *format)
		return runOptions{}, exitUsageError
	}
	return runOptions{
		cesPath: fs.Arg(0), seed: *seed, format: *format,
		out: *out, recordPath: *recordPath, dryRun: *dryRun, strict: *strict,
	}, exitOK
}

// emitRunReport renders the report and maps output failures to exit code 2.
func emitRunReport(o runOutcome, stdout, stderr io.Writer) int {
	data := renderReport(buildRunReport(o), o.opts.format)
	if err := writeOutput(data, o.opts.out, stdout, stderr); err != nil {
		return exitInternalError
	}
	return o.exitCode
}

// writeRunHelp prints the deterministic run usage text.
func writeRunHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, runHelp)
}
