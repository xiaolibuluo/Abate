/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cli

import (
	"abate/kernel/mutation"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strings"
)

// mutateSchema identifies the deterministic real-mutation report layout.
const mutateSchema = "abate.mutate.v1"

// mutateKillingModels states the honest scope of every real backend:
// each verdict comes from an actual test execution of the language
// toolchain against a materialized mutant workspace, with no static
// approximation involved.
var mutateKillingModels = map[string]string{
	"go": "real-go-test-execution (go build + go test -run)",
	"py": "real-pytest-execution (py_compile + pytest node run)",
	"js": "real-node-test-execution (node --check + node --test per file)",
	"c":  "real-gcc-execution (gcc -fsyntax-only + per-source --coverage link/run + gcov)",
	"rs": "real-rustc-execution (rustc --emit=metadata + -Cinstrument-coverage run + llvm-cov)",
	"cs": "real-dotnet-execution (Roslyn csc + dotnet build/run + dotnet-coverage cobertura)",
	"kt": "real-kotlinc-execution (kotlinc compile + JVM run + JaCoCo agent/cli coverage)",
}

// mutateHelp documents the mutate command and its copyable examples.
const mutateHelp = `Usage: abate mutate <path> [flags]
Run a deterministic mutation experiment against a real Go workspace. Every
mutable token discovered by the official Go frontend is rewritten in place
inside a private staged copy, and every verdict (KILLED / SURVIVED) is the
outcome of a real ` + "`go build`" + ` plus ` + "`go test -run`" + ` execution. There is
no static killing approximation in this command.
Flags:
  --seed uint           Deterministic experiment seed (default 42).
  --parallel int        Concurrent mutant-judging workers (default 1 = sequential).
                        Values above one are clamped to the CPU count. The report
                        is byte-identical to the sequential run: worker results
                        are written back by scheduled mutation index, so the
                        mutation score and both digests never depend on timing.
  --chaos               Inject and roll back the twelve C-family chaos faults
                        of the L6 injection kernel before judging mutants. The
                        faults live in the fully virtual kernel world (zero
                        syscalls), so the mutation report stays byte-identical
                        to a chaos-free run over the same workspace; the report
                        gains a deterministic chaos summary block instead.
  --native              Use native sandbox mode: Once directories (vendor,
                        node_modules, target) are copied once and symlinked
                        into each per-mutant workspace, reducing copy cost.
  --format console|json Report format (default console).
  --out path            Also write the report bytes to a file.
  --help, -h            Print this help text.
` + exitCodeHelp + `
Examples:
  abate mutate ./lang/gomutate
  abate mutate ./mymodule --seed 42 --format json
  abate mutate ./internal/yamlmini --parallel 4 --format json
`

// mutateOptions holds the parsed mutate flags.
type mutateOptions struct {
	seed     uint64
	format   string
	out      string
	path     string
	lang     string
	parallel int
	chaos    bool
	native   bool
}

// mutateView is the deterministic wire layout of a real-mutation report.
type mutateView struct {
	Schema       string        `json:"schema"`
	Command      string        `json:"command"`
	Seed         uint64        `json:"seed"`
	ExitCode     int           `json:"exitCode"`
	Root         string        `json:"root"`
	Module       string        `json:"module"`
	Language     string        `json:"language"`
	Tokens       int           `json:"tokens"`
	Tests        int           `json:"tests"`
	Mutants      int           `json:"mutants"`
	Killed       uint64        `json:"killed"`
	Survived     uint64        `json:"survived"`
	Equivalent   uint64        `json:"equivalent"`
	Timeout      uint64        `json:"timeout"`
	Score        float64       `json:"score"`
	KillingModel string        `json:"killingModel"`
	Parallel     int           `json:"parallel"`
	Chaos        *ChaosSummary `json:"chaos,omitempty"`
	Fingerprint  string        `json:"fingerprint"`
	ReportDigest string        `json:"reportDigest"`
}

// cmdMutate runs the real-mutation experiment end to end.
func cmdMutate(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		_, _ = fmt.Fprint(stdout, mutateHelp)
		return exitOK
	}
	opts, code := parseMutateOptions(args, stderr)
	if code != exitOK {
		return code
	}
	var (
		rep     mutation.Report
		backend mutationBackend
		chaos   *ChaosSummary
		err     error
	)
	if opts.chaos {
		var summary ChaosSummary
		summary, rep, backend, err = runChaosMutation(context.Background(),
			opts.seed, opts.path, opts.lang, opts.parallel, opts.native, defaultChaosMutationPorts())
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "abate mutate: %v\n", err)
			return exitInternalError
		}
		chaos = &summary
	} else {
		rep, backend, err = runMutateExperiment(opts.seed, opts.path, opts.lang, opts.parallel, opts.native)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "abate mutate: %v\n", err)
			return exitInternalError
		}
	}
	view := buildMutateView(opts.seed, opts.path, opts.parallel, rep, backend)
	view.Chaos = chaos
	data := renderMutateReport(view, opts.format)
	if err := writeOutput(data, opts.out, stdout, stderr); err != nil {
		return exitInternalError
	}
	return exitOK
}

// runMutateExperiment builds the real backend of the requested language
// over the target workspace and runs one deterministic mutation experiment
// through the kernel engine. The backend is closed before returning: the
// report already carries every verdict, so no staged workspace persists
// after the run.
func runMutateExperiment(seed uint64, path, lang string, workers int, native bool) (mutation.Report, mutationBackend, error) {
	ctx := context.Background()
	backend, err := newMutationBackend(ctx, lang, path, native)
	if err != nil {
		return mutation.Report{}, nil, err
	}
	defer func() { _ = backend.Close() }()
	engine := mutation.NewEngineParallel(seed, workers)
	rep, err := engine.Run(ctx, backend.Target(), backend.Suite())
	if err != nil {
		return mutation.Report{}, nil, err
	}
	return rep, backend, nil
}

// buildMutateView projects a mutation report onto the wire view, attaching
// the workspace metadata the backend exposes.
func buildMutateView(seed uint64, path string, workers int, rep mutation.Report, backend mutationBackend) mutateView {
	return mutateView{
		Schema: mutateSchema, Command: "mutate", Seed: seed,
		ExitCode: exitOK, Root: backend.Root(), Module: backend.ModulePath(),
		Language: backend.Language(),
		Tokens:   backend.TokenCount(), Tests: backend.TestCount(),
		Mutants: len(rep.Points), Killed: rep.Score.Killed,
		Survived: rep.Score.Survived, Equivalent: rep.Score.Equivalent,
		Timeout: rep.Score.Timeout, Score: rep.Score.Value(),
		KillingModel: mutateKillingModels[backend.Language()],
		Parallel:     workers,
		Fingerprint:  digestHexOfBytes(rep.Fingerprint[:]),
		ReportDigest: digestHex(rep.Bytes()),
	}
}

// digestHexOfBytes renders an already-computed digest as lowercase hex.
func digestHexOfBytes(sum []byte) string {
	const digits = "0123456789abcdef"
	out := make([]byte, 0, len(sum)*2)
	for _, b := range sum {
		out = append(out, digits[b>>4], digits[b&0x0f])
	}
	return string(out)
}

// renderMutateReport serializes the mutate view deterministically.
func renderMutateReport(view mutateView, format string) []byte {
	if format == formatJSON {
		data, err := json.MarshalIndent(view, "", "  ")
		if err == nil {
			return append(data, '\n')
		}
	}
	return renderMutateConsole(view)
}

// renderMutateConsole renders the mutate view as fixed-order text lines.
func renderMutateConsole(view mutateView) []byte {
	var b strings.Builder
	appendLine(&b, "abate %s report\n", view.Command)
	appendLine(&b, "schema: %s\n", view.Schema)
	appendLine(&b, "seed: %d\n", view.Seed)
	appendLine(&b, "exitCode: %d\n", view.ExitCode)
	appendLine(&b, "root: %s\n", view.Root)
	appendLine(&b, "module: %s\n", view.Module)
	appendLine(&b, "language: %s\n", view.Language)
	appendLine(&b, "tokens: %d\n", view.Tokens)
	appendLine(&b, "tests: %d\n", view.Tests)
	appendLine(&b, "mutants: %d\n", view.Mutants)
	appendLine(&b, "killed: %d\n", view.Killed)
	appendLine(&b, "survived: %d\n", view.Survived)
	appendLine(&b, "equivalent: %d\n", view.Equivalent)
	appendLine(&b, "timeout: %d\n", view.Timeout)
	appendLine(&b, "score: %.4f\n", view.Score)
	appendLine(&b, "killingModel: %s\n", view.KillingModel)
	appendLine(&b, "parallel: %d\n", view.Parallel)
	if view.Chaos != nil {
		appendLine(&b, "chaos.operators: %d\n", view.Chaos.Operators)
		appendLine(&b, "chaos.classes: %s\n", strings.Join(view.Chaos.Classes, ","))
		appendLine(&b, "chaos.frames: %d\n", view.Chaos.Frames)
		appendLine(&b, "chaos.digest: %s\n", view.Chaos.Digest)
	}
	appendLine(&b, "fingerprint: %s\n", view.Fingerprint)
	appendLine(&b, "reportDigest: %s\n", view.ReportDigest)
	return []byte(b.String())
}

// parseMutateOptions parses mutate flags once and validates them.
func parseMutateOptions(args []string, stderr io.Writer) (mutateOptions, int) {
	fs := flag.NewFlagSet("mutate", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	seed := fs.Uint64("seed", 42, "deterministic experiment seed")
	parallel := fs.Int("parallel", 1, "concurrent mutant-judging workers")
	chaos := fs.Bool("chaos", false, "inject and roll back the 12 C-family chaos faults before judging mutants")
	native := fs.Bool("native", false, "use native sandbox mode (symlink heavy dirs once)")
	lang := fs.String("lang", "", "mutation backend language (go, py, js, c, rs, cs, kt); empty auto-detects from the workspace")
	format := fs.String("format", formatConsole, "console or json")
	out := fs.String("out", "", "report output path")
	if err := fs.Parse(reorderArgs(args, mutateValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate mutate: %v\n", err)
		_, _ = fmt.Fprint(stderr, mutateHelp)
		return mutateOptions{}, exitUsageError
	}
	if fs.NArg() != 1 {
		_, _ = fmt.Fprint(stderr, mutateHelp)
		return mutateOptions{}, exitUsageError
	}
	if !validFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate mutate: unsupported format %q\n", *format)
		return mutateOptions{}, exitUsageError
	}
	switch *lang {
	case "", "go", "py", "js", "c", "rs", "cs", "kt":
	default:
		_, _ = fmt.Fprintf(stderr, "abate mutate: unsupported language %q\n", *lang)
		return mutateOptions{}, exitUsageError
	}
	return mutateOptions{
		seed: *seed, format: *format, out: *out, path: fs.Arg(0),
		lang: *lang, parallel: *parallel, chaos: *chaos, native: *native,
	}, exitOK
}

// mutateValueFlags lists the mutate flags that consume the following argument.
func mutateValueFlags() []string {
	return []string{"seed", "parallel", "lang", "format", "out"}
}

// chaosHelpNote documents the --chaos flag inside mutate help.
const chaosHelpNote = "  --chaos               Inject and roll back the twelve C-family chaos\n                        faults of the L6 injection kernel before judging\n                        mutants; the mutation report stays byte-identical\n                        to a chaos-free run over the same workspace.\n"
