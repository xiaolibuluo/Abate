//go:build ignore

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command gen_golden freezes the current ABI fact tables into self-developed
// golden files. The ignore build tag keeps it out of normal builds; run it
// explicitly with: go run kernel/abi/gen_golden.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"abate/kernel/abi"
)

const (
	amd64Path = "kernel/abi/testdata/amd64.golden"
	arm64Path = "kernel/abi/testdata/arm64.golden"
	header    = "# Abate 自研冻结 golden：OS ABI 事实数据（syscall number name）\n" +
		"# 由 kernel/abi/gen_golden.go 从本仓实现表生成，可复跑；非上游源码副本。\n"
	dirMode  = 0o750
	fileMode = 0o600
)

// main regenerates both architecture goldens.
func main() {
	if err := writeGolden(amd64Path, abi.AMD64Names); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: %v\n", err)
		os.Exit(1)
	}
	if err := writeGolden(arm64Path, abi.ARM64Names); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "gen_golden: %v\n", err)
		os.Exit(1)
	}
	_, _ = fmt.Printf("golden written: %s, %s\n", amd64Path, arm64Path)
}

// writeGolden renders one table in ascending number order for byte stability.
func writeGolden(path string, table map[uintptr]string) error {
	numbers := make([]uintptr, 0, len(table))
	for no := range table {
		numbers = append(numbers, no)
	}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	var b strings.Builder
	b.WriteString(header)
	for _, no := range numbers {
		_, _ = fmt.Fprintf(&b, "%d %s\n", no, table[no])
	}
	if err := os.MkdirAll("kernel/abi/testdata", dirMode); err != nil {
		return fmt.Errorf("create golden dir: %w", err)
	}
	if err := os.WriteFile(path, []byte(b.String()), fileMode); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
