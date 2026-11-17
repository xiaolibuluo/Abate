/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"abate/lang/check"
	"abate/lang/gofront"
	"abate/lang/rule"
	gobuiltin "abate/lang/rule/builtin/go"
)

// checkOptions holds the parsed check flags and positional paths.
type checkOptions struct {
	format   string
	severity string
	rules    string
	out      string
	paths    []string
}

// checkHelp documents the check command and its copyable examples.
const checkHelp = `Usage: abate check <paths...> [flags]

Parse Go sources with the AIR frontend, run the built-in Go rule set and emit
a deterministic console, JSON or SARIF 2.1.0 report. A path may be a file or
a directory; vendor and third_party directories are skipped.

Flags:
  --format console|json|sarif  Report format (default console).
  --severity error|warning|info  Minimum severity (default info).
  --rules prefix[,prefix...]  Restrict the run to rule identifier prefixes.
  --out path                  Also write the report bytes to a file.
  --help, -h                  Print this help text.

Rule-to-operator derivation is the separate abate derive command; check
never mutates sources.

` + exitCodeHelp + `
Examples:
  abate check ./ctl/cli
  abate check ./lang/rule --format sarif --out /tmp/check.sarif
  abate check ./lang/air --rules GO-EG --format json
  abate check ./ctl/cli --help
`

// writeCheckHelp prints the deterministic check usage text.
func writeCheckHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, checkHelp)
}

// cmdCheck wires gofront.Parse, the rule registry, the AIR checker and the
// SARIF/console renderer into one command, then maps the four CLI exit states.
func cmdCheck(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeCheckHelp(stdout)
		return exitOK
	}
	opts, code := parseCheckOptions(args, stderr)
	if code != exitOK {
		return code
	}
	registry := rule.NewRegistry()
	if err := gobuiltin.LoadBuiltins(registry); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate check: load rules: %v\n", err)
		return exitInternalError
	}
	options := check.DefaultOptions()
	options.Format = opts.format
	options.MinSeverity = rule.Severity(opts.severity)
	if opts.rules != "" {
		options.Rules = strings.Split(opts.rules, ",")
	}
	frontend := gofront.NewGoFrontend().WithTypeCheck(false)
	checker := check.NewChecker(frontend, registry.For("go"))
	result, err := checker.Check(opts.paths, options)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate check: %v\n", err)
		return checkErrorExit(err)
	}
	report := result.Render(options.Format)
	if _, writeErr := stdout.Write(report); writeErr != nil {
		return exitInternalError
	}
	if opts.out != "" {
		if writeErr := os.WriteFile(opts.out, report, 0o600); writeErr != nil {
			_, _ = fmt.Fprintf(stderr, "abate check: write report: %v\n", writeErr)
			return exitInternalError
		}
	}
	if check.ExitCode(result, nil) == check.ExitViolations {
		return exitBusinessFailure
	}
	return exitOK
}

// parseCheckOptions parses flags, allowing flags before or after paths.
func parseCheckOptions(args []string, stderr io.Writer) (checkOptions, int) {
	flags := flag.NewFlagSet("check", flag.ContinueOnError)
	flags.SetOutput(stderr)
	format := flags.String("format", check.FormatConsole, "report format")
	severity := flags.String("severity", string(rule.SeverityInfo), "minimum severity")
	rules := flags.String("rules", "", "rule identifier prefixes")
	out := flags.String("out", "", "also write the report to a file")
	if err := flags.Parse(reorderArgs(args, checkValueFlags())); err != nil {
		return checkOptions{}, exitUsageError
	}
	if !validCheckFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate check: unsupported format %q\n", *format)
		return checkOptions{}, exitUsageError
	}
	if !validCheckSeverity(*severity) {
		_, _ = fmt.Fprintf(stderr, "abate check: unsupported severity %q\n", *severity)
		return checkOptions{}, exitUsageError
	}
	if flags.NArg() == 0 {
		_, _ = fmt.Fprintln(stderr, "abate check: no source paths")
		return checkOptions{}, exitUsageError
	}
	return checkOptions{
		format: *format, severity: *severity, rules: *rules,
		out: *out, paths: flags.Args(),
	}, exitOK
}

// validCheckFormat reports whether a format is implemented by the checker.
func validCheckFormat(format string) bool {
	switch format {
	case check.FormatConsole, check.FormatJSON, check.FormatSARIF:
		return true
	default:
		return false
	}
}

// validCheckSeverity reports whether a severity threshold is allocated.
func validCheckSeverity(value string) bool {
	return rule.Severity(value).Valid()
}

// checkErrorExit separates user-facing path/option failures from internal
// frontend failures, keeping the four PLAN-07 section 4 exit states intact.
func checkErrorExit(err error) int {
	switch {
	case errors.Is(err, check.ErrInvalidOptions),
		errors.Is(err, check.ErrInvalidFormat),
		errors.Is(err, check.ErrNoPaths),
		errors.Is(err, check.ErrNoRules),
		errors.Is(err, check.ErrReadSource):
		return exitUsageError
	default:
		return exitInternalError
	}
}
