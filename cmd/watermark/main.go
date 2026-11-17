/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command watermark prints the three-layer provenance watermark and verifies
// the source-layer fingerprint before it reports success.
package main

import (
	"fmt"
	"os"

	"abate/internal/watermark"
)

func main() {
	os.Exit(run(os.Stdout))
}

// run writes the provenance report and returns the process exit code.
func run(stdout *os.File) int {
	info := watermark.Build()
	_, _ = fmt.Fprintf(stdout, "owner: %s\n", info.Owner)
	_, _ = fmt.Fprintf(stdout, "source: %s\n", watermark.Source())
	_, _ = fmt.Fprintf(stdout, "version: %s\n", info.Version)
	_, _ = fmt.Fprintf(stdout, "build-time: %s\n", info.BuildTime)
	_, _ = fmt.Fprintf(stdout, "git-commit: %s\n", info.GitCommit)
	_, _ = fmt.Fprintf(stdout, "fingerprint: %s\n", info.Fingerprint)
	if !watermark.Verify() {
		_, _ = fmt.Fprintln(stdout, "verified: FAIL")
		return 1
	}
	_, _ = fmt.Fprintln(stdout, "verified: PASS")
	return 0
}
