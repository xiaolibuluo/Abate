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
)

// verifyReport is the deterministic byte-comparison result of two reports.
type verifyReport struct {
	Schema     string `json:"schema"`
	Command    string `json:"command"`
	Left       string `json:"left"`
	Right      string `json:"right"`
	LeftSHA    string `json:"leftSha256"`
	RightSHA   string `json:"rightSha256"`
	LeftBytes  int    `json:"leftBytes"`
	RightBytes int    `json:"rightBytes"`
	Equal      bool   `json:"equal"`
	FirstDiff  int    `json:"firstDiff"`
	ExitCode   int    `json:"exitCode"`
}

// verifyOptions holds the parsed verify flags and the two file names.
type verifyOptions struct {
	format string
	out    string
	left   string
	right  string
}

// verifyHelp documents the double-run comparison command and its examples.
const verifyHelp = `Usage: abate verify <a> <b> [--format console|json] [--out path]

Compare two run reports byte for byte and print their SHA-256 digests, sizes,
the first differing offset and the mapped exit code. Zero differences exit 0;
any difference exits 1. Files that cannot be read exit 3.

Examples:
  abate run ctl/cli/testdata/ok.yaml --format json --out report-a.json
  abate run ctl/cli/testdata/ok.yaml --format json --out report-b.json
  abate verify report-a.json report-b.json
  abate verify --format json report-a.json report-b.json
`

// cmdVerify compares two report files and maps the outcome to an exit code.
func cmdVerify(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeVerifyHelp(stdout)
		return exitOK
	}
	opts, code := parseVerifyOptions(args, stderr)
	if code != exitOK {
		return code
	}
	left, err := os.ReadFile(opts.left) //nolint:gosec // both paths are supplied explicitly by the caller.
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate verify: cannot read %q: %v\n", opts.left, err)
		return exitUsageError
	}
	right, err := os.ReadFile(opts.right) //nolint:gosec // both paths are supplied explicitly by the caller.
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate verify: cannot read %q: %v\n", opts.right, err)
		return exitUsageError
	}
	rep := compareReports(opts.left, opts.right, left, right)
	// PLAN-07 §4 requires verify to return 1 on any difference and 0 on equality.
	rep.ExitCode = exitOK
	if !rep.Equal {
		rep.ExitCode = exitBusinessFailure
	}
	data := renderVerifyReport(rep, opts.format)
	if err := writeOutput(data, opts.out, stdout, stderr); err != nil {
		return exitInternalError
	}
	return rep.ExitCode
}

// parseVerifyOptions parses and validates all verify arguments once.
func parseVerifyOptions(args []string, stderr io.Writer) (verifyOptions, int) {
	fs := flag.NewFlagSet("verify", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	format := fs.String("format", formatConsole, "console or json")
	out := fs.String("out", "", "comparison output path")
	if err := fs.Parse(reorderArgs(args, []string{"format", "out"})); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate verify: %v\n", err)
		writeVerifyHelp(stderr)
		return verifyOptions{}, exitUsageError
	}
	if fs.NArg() != 2 {
		writeVerifyHelp(stderr)
		return verifyOptions{}, exitUsageError
	}
	if !validFormat(*format) {
		_, _ = fmt.Fprintf(stderr, "abate verify: unsupported format %q\n", *format)
		return verifyOptions{}, exitUsageError
	}
	return verifyOptions{
		format: *format, out: *out,
		left: fs.Arg(0), right: fs.Arg(1),
	}, exitOK
}

// compareReports computes the deterministic comparison record.
func compareReports(leftName, rightName string, left, right []byte) verifyReport {
	ls := sha256.Sum256(left)
	rs := sha256.Sum256(right)
	return verifyReport{
		Schema:     reportSchema,
		Command:    "verify",
		Left:       leftName,
		Right:      rightName,
		LeftSHA:    hex.EncodeToString(ls[:]),
		RightSHA:   hex.EncodeToString(rs[:]),
		LeftBytes:  len(left),
		RightBytes: len(right),
		Equal:      bytes.Equal(left, right),
		FirstDiff:  firstDifference(left, right),
	}
}

// firstDifference returns the first differing byte offset, or -1 when equal.
func firstDifference(left, right []byte) int {
	n := len(left)
	if len(right) < n {
		n = len(right)
	}
	for i := 0; i < n; i++ {
		if left[i] != right[i] {
			return i
		}
	}
	if len(left) != len(right) {
		return n
	}
	return -1
}

// renderVerifyReport serializes the comparison in a fixed field order.
func renderVerifyReport(rep verifyReport, format string) []byte {
	if format == formatJSON {
		b, err := json.MarshalIndent(rep, "", "  ")
		if err == nil {
			return append(b, '\n')
		}
	}
	return renderVerifyConsole(rep)
}

// renderVerifyConsole renders the comparison as deterministic text lines.
func renderVerifyConsole(rep verifyReport) []byte {
	var b strings.Builder
	appendLine(&b, "abate verify report\n")
	appendLine(&b, "schema: %s\n", rep.Schema)
	appendLine(&b, "left: %s\n", rep.Left)
	appendLine(&b, "right: %s\n", rep.Right)
	appendLine(&b, "leftSha256: %s\n", rep.LeftSHA)
	appendLine(&b, "rightSha256: %s\n", rep.RightSHA)
	appendLine(&b, "leftBytes: %d\n", rep.LeftBytes)
	appendLine(&b, "rightBytes: %d\n", rep.RightBytes)
	appendLine(&b, "equal: %t\n", rep.Equal)
	appendLine(&b, "firstDiff: %d\n", rep.FirstDiff)
	appendLine(&b, "exitCode: %d\n", rep.ExitCode)
	return []byte(b.String())
}

// writeVerifyHelp prints the deterministic verify usage text.
func writeVerifyHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, verifyHelp)
}
