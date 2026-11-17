//go:build ignore

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command gen_golden freezes the mutation-semantics verdicts produced by the
// repository's own parity engine over its isomorphic fixture. The ignore build
// tag keeps it out of normal builds; run it explicitly with:
//
//	go run kernel/mutation/pitparity/gen_golden.go
//
// The output is a self-developed frozen golden. It is a historical observation
// of this repository's engine, not an upstream official artifact.
package main

import (
	"fmt"
	"os"
	"strings"

	"abate/kernel/mutation/pitparity"
)

const (
	sourcePath = "kernel/mutation/testdata/pit/pit-mutations.json"
	outPath    = "kernel/mutation/testdata/pit/golden/mutation-semantics.golden"
	dirMode    = 0o750
	fileMode   = 0o600
)

// main regenerates the frozen semantics golden.
func main() {
	records, err := pitparity.LoadMutations(sourcePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: load source: %v\n", err)
		os.Exit(1)
	}
	report, err := pitparity.Compare(records)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: compare: %v\n", err)
		os.Exit(1)
	}

	var b strings.Builder
	b.WriteString("# Abate 自研变异语义冻结 golden（历史观测值，非上游官方产物）\n")
	b.WriteString("# 由 kernel/mutation/pitparity/gen_golden.go 用本仓引擎与同构夹具生成，可复跑。\n")
	fmt.Fprintf(&b, "total=%d\ncomparable=%d\nconsistent=%d\nrate=%s\n",
		report.Total, report.Comparable, report.Consistent, report.Rate)
	for i, outcome := range report.Outcomes {
		fmt.Fprintf(&b, "outcome\t%d\t%s\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%v\n",
			i+1, outcome.Method, outcome.Line, outcome.Mutator,
			outcome.PITStatus, outcome.Node, outcome.Class,
			outcome.AbateVerdict, outcome.Witness, outcome.Disposition,
			outcome.Consistent)
	}

	if err := os.MkdirAll("kernel/mutation/testdata/pit/golden", dirMode); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: create dir: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(outPath, []byte(b.String()), fileMode); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: write: %v\n", err)
		os.Exit(1)
	}
	_, _ = fmt.Printf("golden written: %s\n", outPath)
}
