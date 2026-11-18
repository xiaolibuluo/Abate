/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"flag"
	"fmt"
	"io"
)

// versionText is the fixed build provenance emitted by the version command.
const versionText = `abate 1.2.26
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
`

// versionHelp documents the version command and its examples.
const versionHelp = `Usage: abate version

Print the Abate build version, schema, license and provenance watermark.
This command accepts no positional arguments and reads no input.

Examples:
  abate version
  abate version --help
`

// cmdVersion prints build provenance and accepts no positional arguments.
func cmdVersion(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeVersionHelp(stdout)
		return exitOK
	}
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate version: %v\n", err)
		writeVersionHelp(stderr)
		return exitUsageError
	}
	if fs.NArg() != 0 {
		writeVersionHelp(stderr)
		return exitUsageError
	}
	_, _ = fmt.Fprint(stdout, versionText)
	return exitOK
}

// writeVersionHelp prints the deterministic version usage text.
func writeVersionHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, versionHelp)
}
