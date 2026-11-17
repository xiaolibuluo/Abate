//go:build ignore

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command gen_golden freezes the self-developed semantic golden produced by
// the repository's own parity engine over its isomorphic fixture. The ignore
// build tag keeps it out of normal builds; run it explicitly with:
//
//	go run kernel/mutation/pitparity/gen_golden.go [-out <path>]
//
// The output is a self-developed frozen golden. It is a historical observation
// of this repository's engine and witness, not an upstream official artifact.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"abate/kernel/mutation/pitparity"
)

const (
	dirMode  = 0o750
	fileMode = 0o600
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: %v\n", err)
		os.Exit(1)
	}
}

// run regenerates the frozen golden and writes it to the requested path.
func run() error {
	out := flag.String("out", pitparity.GoldenRelativePath, "golden output path")
	flag.Parse()
	payload, err := pitparity.RegenerateGolden()
	if err != nil {
		return fmt.Errorf("regenerate golden: %w", err)
	}
	if dir := filepath.Dir(*out); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirMode); err != nil {
			return fmt.Errorf("create dir: %w", err)
		}
	}
	if err := os.WriteFile(*out, payload, fileMode); err != nil {
		return fmt.Errorf("write golden: %w", err)
	}
	_, _ = fmt.Printf("自研 golden 已生成: %s\n", *out)
	return nil
}
