/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cmutate

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

// gcovReport mirrors the gcov --json-format document: the files list
// carries one entry per measured source with the per-line execution
// counts under lines. Only the fields the matrix consumes are decoded.
type gcovReport struct {
	Files []struct {
		File  string `json:"file"`
		Lines []struct {
			LineNumber int  `json:"line_number"`
			Count      int  `json:"count"`
			Unexecuted bool `json:"unexecuted_block"`
		} `json:"lines"`
	} `json:"files"`
}

// MeasureCoverage runs every discovered test once under gcc coverage
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

// measureOne executes one test binary under gcc --coverage instrumentation
// with a private build directory, then folds the executed line set of the
// pristine run into the coverage matrix under workspace-relative slash
// paths. Line numbers come straight from the gcov JSON document.
func (d driver) measureOne(ctx context.Context, root string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	covDir, err := os.MkdirTemp("", "abate-cmutate-cov-")
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
	passed, runErr := d.runBinary(ctx, buildDir, root)
	if runErr != nil {
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, runErr)
	}
	if !passed {
		return fmt.Errorf("%w: pristine test %s failed",
			mutcore.ErrCoverageFailed, test.ID)
	}
	return d.foldGcov(ctx, root, buildDir, test, matrix)
}

// foldGcov decodes the gcov JSON document of one instrumented run and
// records every executed line of every workspace source into the matrix.
// Sources outside the workspace (system headers) are dropped by the
// workspace-relative path check. gcov needs the matching .gcno and .gcda
// pair; both live inside the private coverage directory.
func (d driver) foldGcov(ctx context.Context, root, covDir string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	sources, err := sourceFiles(root)
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	for _, src := range sources {
		if err := d.foldOneSource(ctx, root, covDir, src, test, matrix); err != nil {
			return err
		}
	}
	return nil
}

// foldOneSource decodes the gcov JSON of one instrumented source file and
// folds its executed lines into the matrix. The gcov invocation runs
// inside the private coverage directory where the .gcno/.gcda pair lives,
// and the source is passed by its bare relative name: gcov resolves the
// notes and data files against the working directory, so an absolute
// source path would send it looking for them next to the source instead
// of inside the coverage directory and fail. The JSON document still
// records the source by its absolute path, which workspaceRelative
// accepts.
func (d driver) foldOneSource(ctx context.Context, root, covDir, src string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	cmd := exec.CommandContext(ctx, d.cfg.GcovBin, "--json-format", "--stdout", src)
	cmd.Dir = covDir
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%w: gcov %s: %v", mutcore.ErrCoverageFailed, src, err)
	}
	var decoded gcovReport
	if err := json.Unmarshal(out, &decoded); err != nil {
		return fmt.Errorf("%w: gcov json %s: %v", mutcore.ErrCoverageFailed, src, err)
	}
	for _, file := range decoded.Files {
		relFile := workspaceRelative(root, file.File)
		if relFile == "" {
			continue
		}
		for _, line := range file.Lines {
			if line.Count == 0 {
				continue
			}
			matrix.Add(test.ID, relFile, line.LineNumber)
		}
	}
	return nil
}

// workspaceRelative converts one gcov source path into the
// workspace-relative slash form the fingerprints expect. gcov resolves
// the source argument against the directory the compiler ran in, which
// the driver pins to the workspace root, so entries resolve against root
// rather than the coverage build directory. Files outside the workspace
// (system headers) return the empty string and are dropped.
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
