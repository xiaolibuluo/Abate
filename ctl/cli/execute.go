/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"fmt"
	"io"
	"os"
)

// Execute runs the Abate command line and returns the process exit code.
//
// Every invocation is independent: the command tree holds no state between
// calls, so repeated runs with identical arguments produce identical bytes.
func Execute(args []string) int {
	return dispatch(args, os.Stdout, os.Stderr)
}

// dispatch dispatches a command with injectable output streams for tests.
func dispatch(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		writeRootHelp(stdout)
		return exitUsageError
	}
	switch args[0] {
	case "run":
		return cmdRun(args[1:], stdout, stderr)
	case "version":
		return cmdVersion(args[1:], stdout, stderr)
	case "verify":
		return cmdVerify(args[1:], stdout, stderr)
	case "replay":
		return cmdReplay(args[1:], stdout, stderr)
	case "check":
		return cmdCheck(args[1:], stdout, stderr)
	case "derive":
		return cmdDerive(args[1:], stdout, stderr)
	case "plugin":
		return cmdPlugin(args[1:], stdout, stderr)
	case "web":
		return cmdWeb(args[1:], stdout, stderr)
	case "help", "--help", "-h":
		writeRootHelp(stdout)
		return exitOK
	default:
		_, _ = fmt.Fprintf(stderr, "abate: unknown command %q\n", args[0])
		writeRootHelp(stderr)
		return exitUsageError
	}
}

// helpRequested reports whether an argument slice asks for command help.
func helpRequested(args []string) bool {
	for _, a := range args {
		if a == "-h" || a == "--help" || a == "-help" {
			return true
		}
	}
	return false
}

// writeRootHelp prints the deterministic root usage text.
func writeRootHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, rootHelp)
}

// rootHelp documents the registered command tree and its copyable examples.
const rootHelp = `abate - deterministic experiment control plane

Usage:
  abate <command> [flags]

Commands:
  run <ces-file>     Execute a CES experiment and emit a deterministic report.
                     --record <path> also writes a SIM-REC archive.
  verify <a> <b>     Compare two run reports byte for byte.
  replay <archive> [other]  Verify an archive, optionally compare two.
  check <paths...>   Run 7-language built-in rule sets by file extension
                     (.go/.py/.js/.c/.cs/.rs/.kt); emit console/JSON/SARIF.
  derive             Derive M-family operators from rules; run a mutation
                     experiment (static AIR-level killing approximation).
  plugin             Verify, load, call, list and hot-swap signed plugins.
  web                Serve the control API with a self-contained browser console.
  version            Print build version and provenance.

Global:
  --help, -h          Print this help text.

` + exitCodeHelp + `
Examples:
  abate run ctl/cli/testdata/ok.yaml
  abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
  abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
  abate replay /tmp/run.arch
  abate verify report-a.json report-b.json
  abate check ./ctl/cli
  abate check ./lang/rule --format sarif --out /tmp/check.sarif
  abate derive --seed 42 --format json
  abate plugin load --manifest plugin-manifest.json --type operator
  abate plugin call --id abate/operator --entry apply --args ZWJhdGU=
  abate plugin --help
  abate web --listen 127.0.0.1:8430
  abate version
`

// exitCodeHelp is the single source of truth for the PLAN-07 §4 exit codes.
const exitCodeHelp = `Exit codes (PLAN-07 §4):
  0  success (legal CES and steady state HELD).
  1  business failure (steady state VIOLATED during execution).
  2  internal error (scheduler or report output failure).
  3  usage error (invalid arguments, invalid CES document, or a CES-VALID
     error violation; --strict promotes warnings to this code too).
`
