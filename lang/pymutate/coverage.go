/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package pymutate

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

// coverageReport mirrors the JSON document `coverage json` emits: the
// files map carries one entry per measured source with the executed line
// numbers under executed_lines.
type coverageReport struct {
	Files map[string]struct {
		ExecutedLines []int `json:"executed_lines"`
	} `json:"files"`
}

// MeasureCoverage runs every discovered test once under coverage.py on the
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

// measureOne executes one collected node id under coverage.py with a
// private data file, then folds the executed line set of the pristine run
// into the coverage matrix under workspace-relative slash paths.
func (d driver) measureOne(ctx context.Context, root string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	dataFile := filepath.Join(os.TempDir(), fmt.Sprintf("abate-pymutate-cov-%d", os.Getpid()))
	defer func() { _ = os.Remove(dataFile) }()
	cmd := exec.CommandContext(ctx, d.cfg.PythonBin, "-m", "coverage", "run",
		"--data-file", dataFile, "-m", "pytest",
		"-q", "-x", "--disable-warnings", "-p", "no:cacheprovider", test.ID)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: pristine test %s failed: %s",
			mutcore.ErrCoverageFailed, test.ID, tail(buf.String()))
	}
	report, err := d.exportCoverage(ctx, root, dataFile)
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(report) }()
	data, err := os.ReadFile(report) //nolint:gosec // path is our own temp artifact.
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	var decoded coverageReport
	if err := json.Unmarshal(data, &decoded); err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	for file, entry := range decoded.Files {
		relFile := workspaceRelative(root, file)
		if relFile == "" {
			continue
		}
		for _, line := range entry.ExecutedLines {
			matrix.Add(test.ID, relFile, line)
		}
	}
	return nil
}

// exportCoverage renders the private coverage data file of one run into a
// JSON report inside the system temporary tree and returns its path.
func (d driver) exportCoverage(ctx context.Context, root, dataFile string) (string, error) {
	report := filepath.Join(os.TempDir(), fmt.Sprintf("abate-pymutate-cov-%d.json", os.Getpid()))
	cmd := exec.CommandContext(ctx, d.cfg.PythonBin, "-m", "coverage", "json",
		"--data-file", dataFile, "-o", report)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "PYTHONDONTWRITEBYTECODE=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w: coverage json: %s",
			mutcore.ErrCoverageFailed, tail(buf.String()))
	}
	return report, nil
}

// workspaceRelative converts one coverage path into the workspace-relative
// slash form the fingerprints expect. coverage.py emits relative paths
// anchored at the directory its run executed in, which the driver pins to
// the workspace root, so relative entries resolve against root rather than
// the Go process working directory. Files outside the workspace
// (interpreter stdlib, site-packages) return the empty string and are
// dropped.
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
