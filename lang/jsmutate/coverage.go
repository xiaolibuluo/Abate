/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package jsmutate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"abate/lang/mutcore"
)

// v8Coverage mirrors the V8 coverage document NODE_V8_COVERAGE emits: a
// result list of per-script entries with function ranges and execution
// counts. Only the fields the matrix consumes are decoded.
type v8Coverage struct {
	Result []struct {
		URL       string `json:"url"`
		Functions []struct {
			FunctionName string `json:"functionName"`
			Ranges       []struct {
				StartOffset int `json:"startOffset"`
				EndOffset   int `json:"endOffset"`
				Count       int `json:"count"`
			} `json:"ranges"`
		} `json:"functions"`
	} `json:"result"`
}

// MeasureCoverage runs every discovered test once under V8 coverage on the
// pristine tree and records which lines of which files each test executed.
// A failing pristine test is fatal: no mutant may be judged against a red
// baseline. Failures wrap ErrCoverageFailed.
func (d driver) MeasureCoverage(ctx context.Context, root string, tests []mutcore.TestRef) (*mutcore.CoverMatrix, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	matrix := mutcore.NewCoverMatrix()
	for _, test := range tests {
		if err := d.measureOne(ctx, root, test, matrix); err != nil {
			return nil, err
		}
	}
	return matrix, nil
}

// measureOne executes one test file with NODE_V8_COVERAGE pointed at a
// private directory, then folds the executed line set of the pristine run
// into the coverage matrix under workspace-relative slash paths. Line
// numbers come from mapping each covered function range's start offset
// back through the source text of the staged file.
func (d driver) measureOne(ctx context.Context, root string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	covDir, err := os.MkdirTemp("", "abate-jsmutate-cov-")
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	defer func() { _ = os.RemoveAll(covDir) }()
	cmd := exec.CommandContext(ctx, d.cfg.NodeBin, "--test", test.Name)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "NODE_V8_COVERAGE="+covDir)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: pristine test %s failed: %s",
			mutcore.ErrCoverageFailed, test.ID, tail(buf.String()))
	}
	reports, err := filepath.Glob(filepath.Join(covDir, "coverage-*.json"))
	if err != nil || len(reports) == 0 {
		return fmt.Errorf("%w: no v8 coverage report for %s",
			mutcore.ErrCoverageFailed, test.ID)
	}
	for _, report := range reports {
		if err := d.foldReport(root, test, report, matrix); err != nil {
			return err
		}
	}
	return nil
}

// foldReport decodes one V8 coverage document and records every executed
// line of every workspace file into the matrix. Scripts outside the
// workspace (node internals, node:test harness) are dropped by the
// workspace-relative path check.
func (d driver) foldReport(root string, test mutcore.TestRef, report string, matrix *mutcore.CoverMatrix) error {
	data, err := os.ReadFile(report) //nolint:gosec // path is our own temp artifact.
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	var decoded v8Coverage
	if err := json.Unmarshal(data, &decoded); err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	for _, script := range decoded.Result {
		file := workspaceRelative(root, script.URL)
		if file == "" {
			continue
		}
		source, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(file))) //nolint:gosec // workspace file.
		if err != nil {
			return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
		}
		lineIndex := buildLineIndex(source)
		for _, fn := range script.Functions {
			for _, r := range fn.Ranges {
				if r.Count == 0 {
					continue
				}
				start := offsetToLine(lineIndex, r.StartOffset)
				end := offsetToLine(lineIndex, r.EndOffset)
				for line := start; line <= end; line++ {
					matrix.Add(test.ID, file, line)
				}
			}
		}
	}
	return nil
}

// buildLineIndex returns the byte offsets at which every line of source
// starts, so an absolute byte offset maps to its one-based line with one
// binary search.
func buildLineIndex(source []byte) []int {
	starts := []int{0}
	for i, b := range source {
		if b == '\n' {
			starts = append(starts, i+1)
		}
	}
	return starts
}

// offsetToLine maps one absolute byte offset onto its one-based line using
// the precomputed line-start index.
func offsetToLine(starts []int, offset int) int {
	lo, hi := 0, len(starts)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if starts[mid] <= offset {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo == 0 {
		return 1
	}
	return lo
}

// workspaceRelative converts one V8 script URL into the workspace-relative
// slash form the fingerprints expect. file:// URLs resolve to their path;
// plain paths resolve against the workspace root. Scripts outside the
// workspace return the empty string and are dropped.
func workspaceRelative(root, url string) string {
	path := strings.TrimPrefix(url, "file://")
	if strings.Contains(path, "://") || strings.HasPrefix(path, "node:") {
		return ""
	}
	if idx := strings.IndexByte(path, ':'); idx >= 0 && !filepath.IsAbs(path) {
		return ""
	}
	abs := path
	if !filepath.IsAbs(abs) {
		abs = filepath.Join(root, path)
	}
	rel, err := filepath.Rel(root, abs)
	if err != nil || strings.HasPrefix(rel, "..") {
		return ""
	}
	return filepath.ToSlash(rel)
}

// tail truncates one toolchain output blob to the last 512 bytes so error
// values stay readable without flooding the report.
func tail(out string) string {
	const max = 512
	if len(out) <= max {
		return out
	}
	return out[len(out)-max:]
}
