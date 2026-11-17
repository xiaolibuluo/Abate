/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"abate/sim/record"
)

// replayOptions holds parsed replay flags and up to two archive paths.
type replayOptions struct {
	format string
	out    string
	left   string
	right  string
}

// replayReport is the deterministic replay verification record.
type replayReport struct {
	Schema    string `json:"schema"`
	Command   string `json:"command"`
	Left      string `json:"left"`
	Right     string `json:"right"`
	LeftSHA   string `json:"leftSha256"`
	RightSHA  string `json:"rightSha256,omitempty"`
	Frames    int    `json:"frames"`
	Verified  bool   `json:"verified"`
	Compared  bool   `json:"compared"`
	Equal     bool   `json:"equal"`
	FirstDiff string `json:"firstDiff,omitempty"`
	Summary   string `json:"summary,omitempty"`
	ExitCode  int    `json:"exitCode"`
}

// replayHelp documents the archive verification and comparison command.
const replayHelp = `Usage: abate replay <archive> [compare-archive] [flags]

Load a SIM-REC archive, verify its hash chain, total-order sequence and
virtual-time monotonicity, then optionally compare a second archive and print
the structured Diff. Archives are read-only; no observation is fabricated.

Flags:
  --format console|json  Report format (default console).
  --out path             Also write the report bytes to a file.
  --help, -h             Print this help text.

` + exitCodeHelp + `
Replay exit codes (PLAN-07 §4):
  0  verified (and equal when two archives are supplied).
  1  comparison difference (the structured Diff is printed).
  2  internal error (archive parse, integrity or report output failure).
  3  usage error (invalid arguments or an unreadable archive path).

Examples:
  abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
  abate replay /tmp/run.arch
  abate replay /tmp/run-a.arch /tmp/run-b.arch
`

// cmdReplay verifies one archive and optionally compares it with a second one.
func cmdReplay(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeReplayHelp(stdout)
		return exitOK
	}
	opts, code := parseReplayOptions(args, stderr)
	if code != exitOK {
		return code
	}
	leftBytes, err := os.ReadFile(opts.left) //nolint:gosec // the archive path is supplied explicitly by the caller.
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: cannot read %q: %v\n", opts.left, err)
		return exitUsageError
	}
	left, err := record.Load(bytes.NewReader(leftBytes))
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: cannot parse %q: %v\n", opts.left, err)
		return exitInternalError
	}
	if err := (record.Replayer{}).Verify(left); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: verify %q: %v\n", opts.left, err)
		return exitInternalError
	}
	rep := replayReport{
		Schema: reportSchema, Command: "replay", Left: opts.left,
		LeftSHA: hex.EncodeToString(archiveDigest(leftBytes)),
		Frames:  left.Len(), Verified: true, Equal: true, ExitCode: exitOK,
	}
	if opts.right != "" {
		if code := compareArchives(&rep, left, opts.right, stderr); code != exitOK {
			return code
		}
	}
	data := renderReplayReport(rep, opts.format)
	if err := writeOutput(data, opts.out, stdout, stderr); err != nil {
		return exitInternalError
	}
	return rep.ExitCode
}

// compareArchives loads and compares a second archive, filling the diff fields.
func compareArchives(rep *replayReport, left *record.Archive, rightPath string,
	stderr io.Writer,
) int {
	rightBytes, err := os.ReadFile(rightPath) //nolint:gosec // the archive path is supplied explicitly by the caller.
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: cannot read %q: %v\n", rightPath, err)
		return exitUsageError
	}
	right, err := record.Load(bytes.NewReader(rightBytes))
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: cannot parse %q: %v\n", rightPath, err)
		return exitInternalError
	}
	if err := (record.Replayer{}).Verify(right); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: verify %q: %v\n", rightPath, err)
		return exitInternalError
	}
	diff, err := record.Compare(left, right)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: compare: %v\n", err)
		return exitInternalError
	}
	rep.Right = rightPath
	rep.RightSHA = hex.EncodeToString(archiveDigest(rightBytes))
	rep.Compared = true
	rep.Equal = diff.Equal
	if !diff.Equal {
		rep.ExitCode = exitBusinessFailure
		if diff.First != nil {
			rep.FirstDiff = diff.First.String()
		}
		rep.Summary = diff.Summary.String()
	}
	return exitOK
}

// parseReplayOptions parses and validates all replay arguments once.
func parseReplayOptions(args []string, stderr io.Writer) (replayOptions, int) {
	fs := flag.NewFlagSet("replay", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	format := fs.String("format", formatConsole, "console or json")
	out := fs.String("out", "", "report output path")
	if err := fs.Parse(reorderArgs(args, replayValueFlags())); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate replay: %v\n", err)
		writeReplayHelp(stderr)
		return replayOptions{}, exitUsageError
	}
	if fs.NArg() < 1 || fs.NArg() > 2 {
		writeReplayHelp(stderr)
		return replayOptions{}, exitUsageError
	}
	if !validFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate replay: unsupported format %q\n", *format)
		return replayOptions{}, exitUsageError
	}
	opts := replayOptions{format: *format, out: *out, left: fs.Arg(0)}
	if fs.NArg() == 2 {
		opts.right = fs.Arg(1)
	}
	return opts, exitOK
}

// archiveDigest returns the SHA-256 digest of one archive's raw bytes.
func archiveDigest(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

// renderReplayReport serializes the replay report in a fixed field order.
func renderReplayReport(rep replayReport, format string) []byte {
	if format == formatJSON {
		b, err := json.MarshalIndent(rep, "", "  ")
		if err == nil {
			return append(b, '\n')
		}
	}
	return renderReplayConsole(rep)
}

// renderReplayConsole renders the verification as deterministic text lines.
func renderReplayConsole(rep replayReport) []byte {
	var b strings.Builder
	appendLine(&b, "abate replay report\n")
	appendLine(&b, "schema: %s\n", rep.Schema)
	appendLine(&b, "left: %s\n", rep.Left)
	appendLine(&b, "leftSha256: %s\n", rep.LeftSHA)
	appendLine(&b, "right: %s\n", rep.Right)
	appendLine(&b, "rightSha256: %s\n", rep.RightSHA)
	appendLine(&b, "frames: %d\n", rep.Frames)
	appendLine(&b, "verified: %t\n", rep.Verified)
	appendLine(&b, "compared: %t\n", rep.Compared)
	appendLine(&b, "equal: %t\n", rep.Equal)
	if rep.FirstDiff != "" {
		appendLine(&b, "firstDiff: %s\n", rep.FirstDiff)
	}
	if rep.Summary != "" {
		appendLine(&b, "summary: %s\n", rep.Summary)
	}
	appendLine(&b, "exitCode: %d\n", rep.ExitCode)
	return []byte(b.String())
}

// writeReplayHelp prints the deterministic replay usage text.
func writeReplayHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, replayHelp)
}
