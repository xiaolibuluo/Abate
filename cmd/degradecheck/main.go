/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command degradecheck recomputes the full-corpus degrade rate over every
// self-researched corpus entry using the exact frozen E2E Degradation formula.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	exitOK    = 0
	exitFail  = 1
	exitUsage = 2
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

// run produces one deterministic JSON report to stdout and leaves the corpus untouched.
func run(args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		emit(stderr, "usage: degradecheck (no arguments)\n")
		return exitUsage
	}
	report, err := measure()
	if err != nil {
		emit(stderr, "degradecheck: %v\n", err)
		return exitFail
	}
	enc := json.NewEncoder(stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(report); err != nil {
		emit(stderr, "degradecheck: encode: %v\n", err)
		return exitFail
	}
	return exitOK
}

// emit writes one line and explicitly discards the write result.
func emit(w io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(w, format, args...)
}
