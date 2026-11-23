/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package csmutate

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"abate/lang/mutcore"
)

// coberturaReport mirrors the dotnet-coverage cobertura document: every
// class entry carries its source filename and the per-line hit counts of
// its methods. Only the fields the matrix consumes are decoded.
type coberturaReport struct {
	XMLName  xml.Name `xml:"coverage"`
	Packages []struct {
		Classes []struct {
			Filename string `xml:"filename,attr"`
			Methods  []struct {
				Lines []struct {
					Number int `xml:"number,attr"`
					Hits   int `xml:"hits,attr"`
				} `xml:"lines>line"`
			} `xml:"methods>method"`
		} `xml:"classes>class"`
	} `xml:"packages>package"`
}

// MeasureCoverage runs every discovered test once under dotnet-coverage
// on the pristine tree and records which lines of which files each test
// executed. A failing pristine test is fatal: no mutant may be judged
// against a red baseline. Failures wrap ErrCoverageFailed.
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

// measureOne builds the pristine tree into one assembly, runs it under
// dotnet-coverage with a private report file, then folds the executed
// line set of the run into the coverage matrix under workspace-relative
// slash paths. Line numbers come straight from the cobertura document.
func (d driver) measureOne(ctx context.Context, root string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	buildDir, cleanup, buildErr := d.buildProject(ctx, root)
	if buildErr != nil {
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, buildErr)
	}
	defer cleanup()
	report := filepath.Join(buildDir, "coverage.xml")
	assembly := filepath.Join(buildDir, "out", harnessAssembly+".dll")
	collect := exec.CommandContext(ctx, d.cfg.CoverageBin, "collect",
		"-f", "cobertura", "-o", report, d.cfg.DotnetBin, assembly)
	collect.Dir = root
	collect.Env = append(d.env(), "DOTNET_COVERAGE_NOLOGO=1")
	var buf bytes.Buffer
	collect.Stdout = &buf
	collect.Stderr = &buf
	if err := collect.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// dotnet-coverage forwards the exit code of the profiled
			// process: a non-zero exit here is a red baseline, never a
			// fabricated green one.
			return fmt.Errorf("%w: pristine test %s failed: %s",
				mutcore.ErrCoverageFailed, test.ID, tail(buf.String()))
		}
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, err)
	}
	return d.foldReport(root, report, test, matrix)
}

// foldReport decodes one cobertura document and records every executed
// line of every workspace source into the matrix. Sources outside the
// workspace (framework assemblies) are dropped by the workspace-relative
// path check.
func (d driver) foldReport(root, report string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	data, err := os.ReadFile(report) //nolint:gosec // path is our own build artifact.
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	var decoded coberturaReport
	if err := xml.Unmarshal(data, &decoded); err != nil {
		return fmt.Errorf("%w: cobertura xml: %v", mutcore.ErrCoverageFailed, err)
	}
	for _, pkg := range decoded.Packages {
		for _, class := range pkg.Classes {
			relFile := workspaceRelative(root, class.Filename)
			if relFile == "" {
				continue
			}
			for _, method := range class.Methods {
				for _, line := range method.Lines {
					if line.Hits <= 0 {
						continue
					}
					matrix.Add(test.ID, relFile, line.Number)
				}
			}
		}
	}
	return nil
}

// workspaceRelative converts one cobertura source path into the
// workspace-relative slash form the fingerprints expect. dotnet-coverage
// records sources by the absolute path the compiler embedded in the PDB,
// which the build anchored at the workspace root; files outside the
// workspace (framework assemblies) return the empty string and are
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
