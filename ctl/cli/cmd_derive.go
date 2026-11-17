/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"abate/kernel/mutation"
	"abate/lang/derive"
	"abate/lang/rule"
	gobuiltin "abate/lang/rule/builtin/go"
)

// deriveSchema identifies the deterministic derivation report layout.
const deriveSchema = "abate.derive.v1"

// killingModel states the honest scope of the static AIR-level kill predicate.
const killingModel = "static-air-approximation (not a real go test kill)"

// deriveSource is the in-memory Go source parsed into AIR for the derivation
// experiment. It carries every representative lexical token the operator
// classes dispatch on.
const deriveSource = `package sample

func helper(value int) int { return value }

func compare(left int, right int) bool {
	if left < right {
		return true
	}
	if left == right {
		return false
	}
	return left >= 1
}

func loop(limit int) int {
	total := 0
	for i := 0; i < limit; i++ {
		if i == limit {
			break
		}
		total = total + 1
	}
	return helper(total)
}
`

// deriveOptions holds the parsed derive flags.
type deriveOptions struct {
	seed   uint64
	format string
	out    string
	source string
}

// deriveHelp documents the derive command and its copyable examples.
const deriveHelp = `Usage: abate derive [flags]

Derive M-family mutation operators from the built-in Go rule set, parse an
in-memory Go sample into AIR, register the derived operators with KRN-MUT and
run one deterministic mutation experiment. Killing is a static AIR-level
approximation, not a real go test execution.

Flags:
  --seed uint           Deterministic experiment seed (default 42).
  --format console|json Report format (default console).
  --out path            Also write the report bytes to a file.
  --source path         Optional Go source file to parse instead of the sample.
  --help, -h            Print this help text.

` + exitCodeHelp + `
Examples:
  abate derive
  abate derive --seed 42 --format json
  abate derive --seed 7 --format json --out /tmp/derive.json
`

// deriveView is the deterministic wire layout of a derivation report.
type deriveView struct {
	Schema       string   `json:"schema"`
	Command      string   `json:"command"`
	Seed         uint64   `json:"seed"`
	ExitCode     int      `json:"exitCode"`
	Rules        int      `json:"rules"`
	DerivedIDs   []string `json:"derivedIds"`
	Mutants      int      `json:"mutants"`
	Killed       uint64   `json:"killed"`
	Survived     uint64   `json:"survived"`
	Equivalent   uint64   `json:"equivalent"`
	Timeout      uint64   `json:"timeout"`
	KillingModel string   `json:"killingModel"`
	Fingerprint  string   `json:"fingerprint"`
	ReportDigest string   `json:"reportDigest"`
}

// writeDeriveHelp prints the deterministic derive usage text.
func writeDeriveHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, deriveHelp)
}

// cmdDerive runs the rule-to-operator derivation experiment end to end.
func cmdDerive(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeDeriveHelp(stdout)
		return exitOK
	}
	opts, code := parseDeriveOptions(args, stderr)
	if code != exitOK {
		return code
	}
	rep, specs, err := runDeriveExperiment(opts.seed, opts.source)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate derive: %v\n", err)
		return exitInternalError
	}
	view := buildDeriveView(opts.seed, rep, specs)
	data := renderDeriveReport(view, opts.format)
	if err := writeOutput(data, opts.out, stdout, stderr); err != nil {
		return exitInternalError
	}
	return exitOK
}

// runDeriveExperiment loads the built-in rules, parses a Go source into AIR
// and runs the deterministic derivation experiment.
func runDeriveExperiment(seed uint64, sourcePath string) (mutation.Report, []derive.MutationSpec, error) {
	source := []byte(deriveSource)
	if sourcePath != "" {
		data, err := os.ReadFile(sourcePath) //nolint:gosec // the source path is supplied explicitly by the caller.
		if err != nil {
			return mutation.Report{}, nil, err
		}
		source = data
	}
	registry := rule.NewRegistry()
	if err := gobuiltin.LoadBuiltins(registry); err != nil {
		return mutation.Report{}, nil, err
	}
	unit, err := gobuiltin.ParseUnit("derive/sample.go", source)
	if err != nil {
		return mutation.Report{}, nil, err
	}
	return derive.Experiment(seed, unit, registry.For("go"))
}

// buildDeriveView projects a mutation report onto the deterministic wire view.
func buildDeriveView(seed uint64, rep mutation.Report, specs []derive.MutationSpec) deriveView {
	return deriveView{
		Schema: deriveSchema, Command: "derive", Seed: seed,
		ExitCode: exitOK, Rules: len(specs), DerivedIDs: derive.DerivedIDs(rep),
		Mutants: len(rep.Points), Killed: rep.Score.Killed,
		Survived: rep.Score.Survived, Equivalent: rep.Score.Equivalent,
		Timeout: rep.Score.Timeout, KillingModel: killingModel,
		Fingerprint:  hex.EncodeToString(rep.Fingerprint[:]),
		ReportDigest: digestHex(rep.Bytes()),
	}
}

// digestHex renders the SHA-256 digest of report bytes as lowercase hex.
func digestHex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// renderDeriveReport serializes the derivation view deterministically.
func renderDeriveReport(view deriveView, format string) []byte {
	if format == formatJSON {
		data, err := json.MarshalIndent(view, "", "  ")
		if err == nil {
			return append(data, '\n')
		}
	}
	return renderDeriveConsole(view)
}

// renderDeriveConsole renders the derivation view as fixed-order text lines.
func renderDeriveConsole(view deriveView) []byte {
	var b strings.Builder
	appendLine(&b, "abate %s report\n", view.Command)
	appendLine(&b, "schema: %s\n", view.Schema)
	appendLine(&b, "seed: %d\n", view.Seed)
	appendLine(&b, "exitCode: %d\n", view.ExitCode)
	appendLine(&b, "rules: %d\n", view.Rules)
	appendLine(&b, "derivedIds: %d\n", len(view.DerivedIDs))
	for _, id := range view.DerivedIDs {
		appendLine(&b, "  %s\n", id)
	}
	appendLine(&b, "mutants: %d\n", view.Mutants)
	appendLine(&b, "killed: %d\n", view.Killed)
	appendLine(&b, "survived: %d\n", view.Survived)
	appendLine(&b, "equivalent: %d\n", view.Equivalent)
	appendLine(&b, "timeout: %d\n", view.Timeout)
	appendLine(&b, "killingModel: %s\n", view.KillingModel)
	appendLine(&b, "fingerprint: %s\n", view.Fingerprint)
	appendLine(&b, "reportDigest: %s\n", view.ReportDigest)
	return []byte(b.String())
}

// parseDeriveOptions parses derive flags once and validates the format.
func parseDeriveOptions(args []string, stderr io.Writer) (deriveOptions, int) {
	fs := flag.NewFlagSet("derive", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	seed := fs.Uint64("seed", 42, "deterministic experiment seed")
	format := fs.String("format", formatConsole, "console or json")
	out := fs.String("out", "", "report output path")
	source := fs.String("source", "", "optional Go source file")
	if err := fs.Parse(reorderArgs(args, deriveValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate derive: %v\n", err)
		writeDeriveHelp(stderr)
		return deriveOptions{}, exitUsageError
	}
	if fs.NArg() != 0 {
		writeDeriveHelp(stderr)
		return deriveOptions{}, exitUsageError
	}
	if !validFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate derive: unsupported format %q\n", *format)
		return deriveOptions{}, exitUsageError
	}
	return deriveOptions{
		seed: *seed, format: *format, out: *out, source: *source,
	}, exitOK
}
