/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package rsmutate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"abate/lang/mutcore"
)

// covReport mirrors the llvm-cov export --format=text document: the data
// list carries per-file entries whose segments encode line/column spans
// with their execution counts. Only the fields the matrix consumes are
// decoded.
type covReport struct {
	Data []struct {
		Files []struct {
			Filename string  `json:"filename"`
			Segments [][]any `json:"segments"`
		} `json:"files"`
	} `json:"data"`
}

// MeasureCoverage runs every discovered test once under rustc coverage
// instrumentation on the pristine tree and records which lines of which
// files each test executed. A failing pristine test is fatal: no mutant
// may be judged against a red baseline. Failures wrap ErrCoverageFailed.
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

// measureOne executes one test binary under rustc -Cinstrument-coverage
// instrumentation inside a private build directory, merges its raw
// profile with llvm-profdata, then folds the executed line set decoded by
// llvm-cov into the coverage matrix under workspace-relative slash paths.
func (d driver) measureOne(ctx context.Context, root string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	covDir, err := os.MkdirTemp("", "abate-rsmutate-cov-")
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	defer func() { _ = os.RemoveAll(covDir) }()
	buildDir, cleanup, buildErr := d.buildTree(ctx, root, true, covDir)
	if buildErr != nil {
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, buildErr)
	}
	defer cleanup()
	profRaw := filepath.Join(covDir, "coverage.profraw")
	passed, runErr := d.runBinary(ctx, filepath.Join(buildDir, "abate-rs-harness"), root, profRaw)
	if runErr != nil {
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, runErr)
	}
	if !passed {
		return fmt.Errorf("%w: pristine test %s failed",
			mutcore.ErrCoverageFailed, test.ID)
	}
	return d.foldCoverage(ctx, root, buildDir, profRaw, test, matrix)
}

// foldCoverage merges the raw profile of one instrumented run and decodes
// the resulting JSON document, recording every executed line of every
// workspace source into the matrix. Sources outside the workspace (std
// library internals) are dropped by the workspace-relative path check.
func (d driver) foldCoverage(ctx context.Context, root, buildDir, profRaw string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	profData := filepath.Join(buildDir, "coverage.profdata")
	merge := exec.CommandContext(ctx, d.cfg.ProfdataBin, "merge", "-sparse", profRaw, "-o", profData)
	var mergeBuf strings.Builder
	merge.Stdout = &mergeBuf
	merge.Stderr = &mergeBuf
	if err := merge.Run(); err != nil {
		return fmt.Errorf("%w: llvm-profdata merge: %s",
			mutcore.ErrCoverageFailed, tail(mergeBuf.String()))
	}
	binary := filepath.Join(buildDir, "abate-rs-harness")
	export := exec.CommandContext(ctx, d.cfg.CovBin, "export", binary,
		"--instr-profile="+profData, "--format=text")
	out, err := export.Output()
	if err != nil {
		return fmt.Errorf("%w: llvm-cov export: %v", mutcore.ErrCoverageFailed, err)
	}
	var decoded covReport
	if err := json.Unmarshal(out, &decoded); err != nil {
		return fmt.Errorf("%w: llvm-cov json: %v", mutcore.ErrCoverageFailed, err)
	}
	for _, data := range decoded.Data {
		for _, file := range data.Files {
			relFile := workspaceRelative(root, file.Filename)
			if relFile == "" {
				continue
			}
			for _, line := range coveredLines(file.Segments) {
				matrix.Add(test.ID, relFile, line)
			}
		}
	}
	return nil
}

// coveredLines decodes llvm-cov segments into the set of one-based line
// numbers the run executed. Each segment is [line, column, count,
// hasCount, isRegionEntry, ...]; a segment with a positive count covers
// from its own line up to (but excluding) the next segment's line. Only
// the line numbers are materialized, which is the granularity the
// coverage matrix consumes.
func coveredLines(segments [][]any) []int {
	if len(segments) == 0 {
		return nil
	}
	lines := make(map[int]bool)
	for i, seg := range segments {
		if len(seg) < 3 {
			continue
		}
		count, ok := asFloat(seg[2])
		if !ok || count <= 0 {
			continue
		}
		start, ok := asFloat(seg[0])
		if !ok {
			continue
		}
		startLine := int(start)
		endLine := startLine
		if i+1 < len(segments) && len(segments[i+1]) >= 1 {
			if next, nok := asFloat(segments[i+1][0]); nok && int(next) > endLine {
				endLine = int(next) - 1
			}
		}
		for line := startLine; line <= endLine; line++ {
			lines[line] = true
		}
	}
	if len(lines) == 0 {
		return nil
	}
	out := make([]int, 0, len(lines))
	for line := range lines {
		out = append(out, line)
	}
	return out
}

// asFloat coerces one decoded JSON scalar into a float64. The llvm-cov
// document carries integers in the segment cells; encoding/json decodes
// them as float64, which is the only shape this helper accepts.
func asFloat(v any) (float64, bool) {
	f, ok := v.(float64)
	return f, ok
}

// workspaceRelative converts one llvm-cov source path into the
// workspace-relative slash form the fingerprints expect. rustc records
// sources by absolute path; files outside the workspace (std library
// internals) return the empty string and are dropped.
func workspaceRelative(root, file string) string {
	abs := file
	if !filepath.IsAbs(abs) {
		abs = filepath.Join(root, file)
	}
	rel, err := filepath.Rel(root, abs)
	if err != nil || strings.HasPrefix(rel, "..") {
		return ""
	}
	return filepath.ToSlash(rel)
}
