/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command pitparity-report renders the PIT parity report of the KRN-MUT
// mutation kernel.
//
// The command replays the vendored official PIT report against the real
// mutation engine and writes both the machine readable JSON artifact and the
// human readable review document. It never starts a JVM.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"abate/kernel/mutation/pitparity"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "pitparity-report: %v\n", err)
		os.Exit(1)
	}
}

// run executes the report generation and returns the first failure.
func run() error {
	reportPath := flag.String("pit", "", "PIT JSON report path (default: auto locate)")
	xmlPath := flag.String("xml", "", "PIT XML report to convert into the JSON report")
	jsonOut := flag.String("json", ".artifacts/pit-parity.json", "JSON report output path")
	mdOut := flag.String("md", "docs/dev/pit-parity-report.md", "Markdown report output path")
	minRate := flag.String("min-rate", "100.00%", "minimum acceptable consistency rate")
	flag.Parse()

	source := *reportPath
	if source == "" {
		located, err := pitparity.LocateReport(".")
		if err != nil {
			return err
		}
		source = located
	}
	if *xmlPath != "" {
		converted, err := convertXML(*xmlPath)
		if err != nil {
			return err
		}
		if err := writeFile(source, converted); err != nil {
			return err
		}
		_, _ = fmt.Printf("PIT XML 报告已转换为 JSON: %s\n", source)
	}
	records, err := pitparity.LoadMutations(source)
	if err != nil {
		return err
	}
	report, err := pitparity.Compare(records)
	if err != nil {
		return err
	}
	payload, err := report.JSON()
	if err != nil {
		return err
	}
	if err := writeFile(*jsonOut, payload); err != nil {
		return err
	}
	if err := writeFile(*mdOut, report.Markdown()); err != nil {
		return err
	}
	_, _ = fmt.Printf("PIT 变异总数: %d\n", report.Total)
	_, _ = fmt.Printf("可比对点数:   %d\n", report.Comparable)
	_, _ = fmt.Printf("一致点数:     %d\n", report.Consistent)
	_, _ = fmt.Printf("一致率:       %s\n", report.Rate)
	_, _ = fmt.Printf("JSON 报告:    %s\n", *jsonOut)
	_, _ = fmt.Printf("Markdown 报告: %s\n", *mdOut)
	return checkRate(report.Rate, *minRate)
}

// convertXML converts one PIT XML report into the canonical JSON form.
func convertXML(path string) ([]byte, error) {
	//nolint:gosec // the XML path is supplied explicitly by the operator.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取 PIT XML 报告失败 %s: %w", path, err)
	}
	return pitparity.ConvertXML(data)
}

// writeFile writes one artifact, creating the parent directory when needed.
func writeFile(path string, data []byte) error {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o750); err != nil {
			return fmt.Errorf("创建目录失败 %s: %w", dir, err)
		}
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("写入文件失败 %s: %w", path, err)
	}
	return nil
}

// checkRate reports an error when the consistency rate is below the threshold.
func checkRate(rate, minimum string) error {
	actual, err := parseRate(rate)
	if err != nil {
		return err
	}
	want, err := parseRate(minimum)
	if err != nil {
		return err
	}
	if actual < want {
		return fmt.Errorf("一致率 %s 低于门槛 %s", rate, minimum)
	}
	return nil
}

// parseRate parses a percentage into hundredths of a percent.
func parseRate(rate string) (int, error) {
	trimmed := strings.TrimSuffix(strings.TrimSpace(rate), "%")
	whole, frac, found := strings.Cut(trimmed, ".")
	if !found {
		frac = "00"
	}
	if len(frac) > 2 {
		frac = frac[:2]
	}
	for len(frac) < 2 {
		frac += "0"
	}
	wholeValue, err := strconv.Atoi(whole)
	if err != nil {
		return 0, fmt.Errorf("无法解析一致率 %q: %w", rate, err)
	}
	fracValue, err := strconv.Atoi(frac)
	if err != nil {
		return 0, fmt.Errorf("无法解析一致率 %q: %w", rate, err)
	}
	return wholeValue*100 + fracValue, nil
}
