/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"slices"
	"strings"

	"abate/lang/targetgen"
)

// targetgenSchema identifies the deterministic target-generation report layout.
const targetgenSchema = "abate.targetgen.v1"

// targetgenHelp documents the targetgen command and its copyable examples.
const targetgenHelp = `Usage: abate targetgen --out <dir> [flags]
Generate a deterministic mutation-testing target project for one of the
seven abate language frontends (go/py/js/c/cs/rs/kt): a complete,
self-contained project whose functions carry the mutable structures of the
seven operator classes, each guarded by real tests. The same seed always
yields byte-identical files, and every delivered project passes a real
toolchain self-check before it lands on disk.
Flags:
  --out path            Target project directory (required, must not exist).
  --lang go|py|js|c|cs|rs|kt
                        Language of the generated project (default go).
  --seed uint           Deterministic generation seed (default 42).
  --module path         Module path for go.mod, Go only (default example.com/target).
  --package name        Go package name, Go only (default target).
  --funcs int           Number of target functions, 1..64 (default 8).
  --level easy|medium|hard
                        Assertion density steering the expected mutation score:
                        easy pins almost every edge case, hard leaves most
                        behavior unobserved (default medium).
  --mixed ratio         Fraction of functions rendered as mixed-structure
                        combos, in [0, 1] (default 0).
  --complex             Enable the complex structure families in the Go
                        rotation (slice, map, error, nested, method, iface).
  --skip-check          Skip the real toolchain self-check before delivery.
  --format console|json Report format (default console).
  --help, -h            Print this help text.
` + exitCodeHelp + `
Examples:
  abate targetgen --out /tmp/tgt --seed 42 --funcs 14 --level easy
  abate targetgen --out /tmp/tpy --lang py --funcs 10 --mixed 0.5
  abate mutate /tmp/tgt --parallel 4 --format json
`

// targetgenOptions holds the parsed targetgen flags.
type targetgenOptions struct {
	seed      uint64
	out       string
	lang      string
	module    string
	pkg       string
	funcs     int
	level     string
	mixed     float64
	complex   bool
	format    string
	skipCheck bool
}

// targetgenStructureCount is one entry of the deterministic structure
// census: how many generated functions carry one structure family.
type targetgenStructureCount struct {
	Structure string `json:"structure"`
	Count     int    `json:"count"`
}

// targetgenView is the deterministic wire layout of a targetgen report.
type targetgenView struct {
	Schema      string                    `json:"schema"`
	Command     string                    `json:"command"`
	Seed        uint64                    `json:"seed"`
	ExitCode    int                       `json:"exitCode"`
	Out         string                    `json:"out"`
	Lang        string                    `json:"lang"`
	Module      string                    `json:"module"`
	Package     string                    `json:"package"`
	Funcs       int                       `json:"funcs"`
	Level       string                    `json:"level"`
	Mixed       float64                   `json:"mixed"`
	Complex     bool                      `json:"complex"`
	Structures  []targetgenStructureCount `json:"structures"`
	Files       int                       `json:"files"`
	Bytes       int                       `json:"bytes"`
	Tokens      int                       `json:"tokens"`
	Pinned      int                       `json:"pinned"`
	SelfCheck   bool                      `json:"selfCheck"`
	Fingerprint string                    `json:"fingerprint"`
}

// cmdTargetgen runs the deterministic target generation end to end.
func cmdTargetgen(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		_, _ = fmt.Fprint(stdout, targetgenHelp)
		return exitOK
	}
	opts, code := parseTargetgenOptions(args, stderr)
	if code != exitOK {
		return code
	}
	target, err := targetgen.Generate(targetgen.Spec{
		Lang:    opts.lang,
		Module:  opts.module,
		Package: opts.pkg,
		Funcs:   opts.funcs,
		Level:   opts.level,
		Mixed:   opts.mixed,
		Complex: opts.complex,
	}, opts.seed)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate targetgen: %v\n", err)
		return exitUsageError
	}
	if err := targetgen.Write(opts.out, target, !opts.skipCheck); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate targetgen: %v\n", err)
		return exitInternalError
	}
	view := buildTargetgenView(opts, target)
	data := renderTargetgenReport(view, opts.format)
	if err := writeOutput(data, "", stdout, stderr); err != nil {
		return exitInternalError
	}
	return exitOK
}

// buildTargetgenView projects a generated target onto the wire view. The
// structure census is sorted by family name so the report never depends on
// map iteration order.
func buildTargetgenView(opts targetgenOptions, target targetgen.Target) targetgenView {
	counts := make(map[string]int, len(target.Funcs))
	for _, fn := range target.Funcs {
		counts[fn.Structure]++
	}
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	slices.Sort(names)
	structures := make([]targetgenStructureCount, 0, len(names))
	for _, name := range names {
		structures = append(structures, targetgenStructureCount{Structure: name, Count: counts[name]})
	}
	fp := target.Fingerprint()
	return targetgenView{
		Schema: targetgenSchema, Command: "targetgen", Seed: opts.seed,
		ExitCode: exitOK, Out: opts.out, Lang: target.Spec.Lang,
		Module: target.Spec.Module, Package: target.Spec.Package,
		Funcs: len(target.Funcs), Level: target.Spec.Level,
		Mixed: target.Spec.Mixed, Complex: target.Spec.Complex,
		Structures: structures, Files: len(target.Files), Bytes: target.ByteSize(),
		Tokens: target.Tokens, Pinned: target.PinnedStructures, SelfCheck: !opts.skipCheck,
		Fingerprint: digestHexOfBytes(fp[:]),
	}
}

// renderTargetgenReport serializes the targetgen view deterministically.
func renderTargetgenReport(view targetgenView, format string) []byte {
	if format == formatJSON {
		data, err := json.MarshalIndent(view, "", "  ")
		if err == nil {
			return append(data, '\n')
		}
	}
	return renderTargetgenConsole(view)
}

// renderTargetgenConsole renders the targetgen view as fixed-order text lines.
func renderTargetgenConsole(view targetgenView) []byte {
	var b strings.Builder
	appendLine(&b, "abate %s report\n", view.Command)
	appendLine(&b, "schema: %s\n", view.Schema)
	appendLine(&b, "seed: %d\n", view.Seed)
	appendLine(&b, "exitCode: %d\n", view.ExitCode)
	appendLine(&b, "out: %s\n", view.Out)
	appendLine(&b, "lang: %s\n", view.Lang)
	appendLine(&b, "module: %s\n", view.Module)
	appendLine(&b, "package: %s\n", view.Package)
	appendLine(&b, "funcs: %d\n", view.Funcs)
	appendLine(&b, "level: %s\n", view.Level)
	appendLine(&b, "mixed: %g\n", view.Mixed)
	appendLine(&b, "complex: %t\n", view.Complex)
	var census strings.Builder
	for i, entry := range view.Structures {
		if i > 0 {
			census.WriteString(",")
		}
		fmt.Fprintf(&census, "%s=%d", entry.Structure, entry.Count)
	}
	appendLine(&b, "structures: %s\n", census.String())
	appendLine(&b, "files: %d\n", view.Files)
	appendLine(&b, "bytes: %d\n", view.Bytes)
	appendLine(&b, "tokens: %d\n", view.Tokens)
	appendLine(&b, "pinned: %d\n", view.Pinned)
	appendLine(&b, "selfCheck: %t\n", view.SelfCheck)
	appendLine(&b, "fingerprint: %s\n", view.Fingerprint)
	return []byte(b.String())
}

// parseTargetgenOptions parses targetgen flags once and validates them.
func parseTargetgenOptions(args []string, stderr io.Writer) (targetgenOptions, int) {
	fs := flag.NewFlagSet("targetgen", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	seed := fs.Uint64("seed", 42, "deterministic generation seed")
	out := fs.String("out", "", "target project directory")
	lang := fs.String("lang", targetgen.LangGo, "language of the generated project")
	module := fs.String("module", "example.com/target", "module path for go.mod")
	pkg := fs.String("package", "target", "Go package name")
	funcs := fs.Int("funcs", 8, "number of target functions")
	level := fs.String("level", targetgen.LevelMedium, "easy, medium or hard")
	mixed := fs.Float64("mixed", 0, "fraction of mixed-structure functions, in [0, 1]")
	complex := fs.Bool("complex", false, "enable complex Go structure families")
	skipCheck := fs.Bool("skip-check", false, "skip the toolchain self-check")
	format := fs.String("format", formatConsole, "console or json")
	if err := fs.Parse(reorderArgs(args, targetgenValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate targetgen: %v\n", err)
		_, _ = fmt.Fprint(stderr, targetgenHelp)
		return targetgenOptions{}, exitUsageError
	}
	if fs.NArg() != 0 {
		_, _ = fmt.Fprint(stderr, targetgenHelp)
		return targetgenOptions{}, exitUsageError
	}
	if *out == "" {
		_, _ = fmt.Fprintln(stderr, "abate targetgen: --out is required")
		_, _ = fmt.Fprint(stderr, targetgenHelp)
		return targetgenOptions{}, exitUsageError
	}
	if !validFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate targetgen: unsupported format %q\n", *format)
		return targetgenOptions{}, exitUsageError
	}
	return targetgenOptions{
		seed: *seed, out: *out, lang: *lang, module: *module, pkg: *pkg,
		funcs: *funcs, level: *level, mixed: *mixed, complex: *complex,
		format: *format, skipCheck: *skipCheck,
	}, exitOK
}

// targetgenValueFlags lists the targetgen flags that consume the following argument.
func targetgenValueFlags() []string {
	return []string{"seed", "out", "lang", "module", "package", "funcs", "level", "mixed", "format"}
}
