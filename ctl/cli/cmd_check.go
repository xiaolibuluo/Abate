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

	"abate/lang/air"
	"abate/lang/cfront"
	"abate/lang/check"
	"abate/lang/csfront"
	"abate/lang/gofront"
	"abate/lang/jsfront"
	"abate/lang/ktfront"
	"abate/lang/pyfront"
	"abate/lang/rsfront"
	"abate/lang/rule"
	builtinc "abate/lang/rule/builtin/c"
	builtincs "abate/lang/rule/builtin/cs"
	builtingo "abate/lang/rule/builtin/go"
	builtinjs "abate/lang/rule/builtin/js"
	builtinkt "abate/lang/rule/builtin/kt"
	builtinpy "abate/lang/rule/builtin/py"
	builtinrs "abate/lang/rule/builtin/rs"
)

// checkOptions holds the parsed check flags and positional paths.
type checkOptions struct {
	format    string
	severity  string
	rules     string
	languages string
	out       string
	paths     []string
}

// checkHelp documents the check command and its copyable examples.
const checkHelp = `Usage: abate check <paths...> [flags]

Parse sources with the language-specific AIR frontend, run that language's
built-in rule set and emit a deterministic console, JSON or SARIF 2.1.0 report.
Seven languages are dispatched by file extension: .go .py .js .c .cs .rs .kt.
A path may be a file or a directory; vendor, third_party and .git are skipped.

Flags:
  --format console|json|sarif  Report format (default console).
  --severity error|warning|info  Minimum severity (default info).
  --rules prefix[,prefix...]  Restrict the run to rule identifier prefixes.
  --languages lang[,lang...]  Restrict to language tags (go,py,js,c,cs,rs,kt).
  --out path                  Also write the report bytes to a file.
  --help, -h                  Print this help text.

Rule-to-operator derivation is the separate abate derive command; check
never mutates sources.

Exit codes:
  0  no findings (clean).
  1  one or more findings — including warnings (business failure). Use
     --severity error to fail only on errors; --rules to narrow the set.
  2  internal error (frontend or report output failure).
  3  usage error (invalid arguments, missing paths, or no rules matched).

Examples:
  abate check ./ctl/cli
  abate check program.py report.js --format sarif --out /tmp/check.sarif
  abate check ./src --languages py,js --rules PY-PEP8
  abate check ./lang/air --rules GO-EG --format json
  abate check ./ctl/cli --help
`

// writeCheckHelp prints the deterministic check usage text.
func writeCheckHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, checkHelp)
}

// languageBinding holds one language's frontend constructor and rule loader.
type languageBinding struct {
	tag        string
	extensions []string
	frontend   func() (air.Frontend, error)
	loadRules  func(*rule.Registry) error
}

// languageBindings is the single dispatch table for the check command. Each
// language owns an isolated registry so rule identifiers never collide across
// languages (GO-*, PY-*, JS-*, C-*, CS-*, RS-*, KT-*).
var languageBindings = []languageBinding{
	{tag: "go", extensions: []string{".go"},
		frontend: goFrontendFactory, loadRules: builtingo.LoadBuiltins},
	{tag: "py", extensions: []string{".py"},
		frontend: func() (air.Frontend, error) { return pyfront.NewPyFrontend() }, loadRules: builtinpy.LoadBuiltins},
	{tag: "js", extensions: []string{".js", ".jsx", ".mjs", ".cjs"},
		frontend: func() (air.Frontend, error) { return jsfront.NewJsFrontend() }, loadRules: builtinjs.LoadBuiltins},
	{tag: "c", extensions: []string{".c", ".h"},
		frontend: func() (air.Frontend, error) { return cfront.NewCFrontend() }, loadRules: builtinc.LoadBuiltins},
	{tag: "cs", extensions: []string{".cs"},
		frontend: func() (air.Frontend, error) { return csfront.NewCsFrontend() }, loadRules: builtincs.LoadBuiltins},
	{tag: "rs", extensions: []string{".rs"},
		frontend: func() (air.Frontend, error) { return rsfront.NewRsFrontend() }, loadRules: builtinrs.LoadBuiltins},
	{tag: "kt", extensions: []string{".kt", ".kts"},
		frontend: func() (air.Frontend, error) { return ktfront.NewKtFrontend() }, loadRules: builtinkt.LoadBuiltins},
}

// goFrontendFactory adapts the concrete GoFrontend to the air.Frontend factory
// signature used by the other language bindings.
func goFrontendFactory() (air.Frontend, error) {
	return gofront.NewGoFrontend().WithTypeCheck(false), nil
}

// buildMultiChecker constructs a MultiChecker from the seven language
// bindings. Any frontend or rule-set construction failure aborts the command
// with an internal-error exit rather than silently checking fewer languages.
func buildMultiChecker() (check.MultiChecker, error) {
	specs := make([]check.LanguageSpec, 0, len(languageBindings))
	for _, binding := range languageBindings {
		frontend, err := binding.frontend()
		if err != nil {
			return check.MultiChecker{}, fmt.Errorf("init %s frontend: %w", binding.tag, err)
		}
		registry := rule.NewRegistry()
		if err := binding.loadRules(registry); err != nil {
			return check.MultiChecker{}, fmt.Errorf("load %s rules: %w", binding.tag, err)
		}
		rules := registry.For(binding.tag)
		if len(rules) == 0 {
			return check.MultiChecker{}, fmt.Errorf("load %s rules: empty rule set", binding.tag)
		}
		specs = append(specs, check.LanguageSpec{
			Language: binding.tag, Frontend: frontend,
			Rules: rules, Extensions: binding.extensions,
		})
	}
	return check.NewMultiChecker(specs)
}

// cmdCheck wires the multi-language frontends, per-language rule sets, the AIR
// checker and the SARIF/console renderer into one command, then maps the four
// CLI exit states.
func cmdCheck(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeCheckHelp(stdout)
		return exitOK
	}
	opts, code := parseCheckOptions(args, stderr)
	if code != exitOK {
		return code
	}
	checker, err := buildMultiChecker()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate check: %v\n", err)
		return exitInternalError
	}
	options := check.DefaultOptions()
	options.Format = opts.format
	options.MinSeverity = rule.Severity(opts.severity)
	if opts.rules != "" {
		options.Rules = strings.Split(opts.rules, ",")
	}
	if opts.languages != "" {
		options.Languages = strings.Split(opts.languages, ",")
	}
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
	languages := flags.String("languages", "", "language tags")
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
		languages: *languages, out: *out, paths: flags.Args(),
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
