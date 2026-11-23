/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package ktmutate

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"abate/lang/mutcore"
)

// jacocoReport mirrors the jacococli report XML document: every
// sourcefile entry carries the per-line execution counters (ci =
// covered instructions) of its class. Only the fields the matrix
// consumes are decoded.
type jacocoReport struct {
	XMLName  xml.Name `xml:"report"`
	Packages []struct {
		SourceFiles []struct {
			Name  string `xml:"name,attr"`
			Lines []struct {
				Nr int `xml:"nr,attr"`
				Mi int `xml:"mi,attr"`
				Ci int `xml:"ci,attr"`
			} `xml:"line"`
		} `xml:"sourcefile"`
	} `xml:"package"`
}

// MeasureCoverage runs every discovered test once under the JaCoCo agent
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

// measureOne builds the pristine tree into one class directory, runs its
// entry point under the JaCoCo agent with a private exec file, then folds
// the executed line set rendered by jacococli into the coverage matrix
// under workspace-relative slash paths. Line numbers come straight from
// the report document.
func (d driver) measureOne(ctx context.Context, root string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutSeconds(d.cfg.TestTimeout))
	defer cancel()
	build, cleanup, buildErr := d.buildClasses(ctx, root)
	if buildErr != nil {
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, buildErr)
	}
	defer cleanup()
	execFile := filepath.Join(build.classDir, "coverage.exec")
	passed, runErr := d.runMain(ctx, build, root, execFile)
	if runErr != nil {
		return fmt.Errorf("%w: pristine test %s toolchain: %v",
			mutcore.ErrCoverageFailed, test.ID, runErr)
	}
	if !passed {
		return fmt.Errorf("%w: pristine test %s failed",
			mutcore.ErrCoverageFailed, test.ID)
	}
	return d.foldCoverage(ctx, root, build, execFile, test, matrix)
}

// foldCoverage renders the exec data of one instrumented run as XML and
// records every executed line of every workspace source into the matrix.
// Sources outside the workspace (Kotlin runtime internals) are dropped
// because the report only analyzes the application class files, and any
// residual foreign path fails the workspace-relative check.
func (d driver) foldCoverage(ctx context.Context, root string, build build, execFile string, test mutcore.TestRef, matrix *mutcore.CoverMatrix) error {
	report := filepath.Join(build.classDir, "coverage.xml")
	args := []string{"-jar", d.cfg.JacocoCli, "report", execFile, "--xml", report}
	for _, classFile := range build.classes {
		args = append(args, "--classfiles", classFile)
	}
	cmd := exec.CommandContext(ctx, d.cfg.JavaBin, args...)
	cmd.Dir = build.classDir
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w: jacococli report: %v: %s",
			mutcore.ErrCoverageFailed, err, tail(string(out)))
	}
	data, err := os.ReadFile(report) //nolint:gosec // path is our own build artifact.
	if err != nil {
		return fmt.Errorf("%w: %v", mutcore.ErrCoverageFailed, err)
	}
	var decoded jacocoReport
	if err := xml.Unmarshal(data, &decoded); err != nil {
		return fmt.Errorf("%w: jacoco xml: %v", mutcore.ErrCoverageFailed, err)
	}
	for _, pkg := range decoded.Packages {
		for _, source := range pkg.SourceFiles {
			relFile := workspaceRelative(root, source.Name)
			if relFile == "" {
				continue
			}
			for _, line := range source.Lines {
				if line.Ci <= 0 {
					continue
				}
				matrix.Add(test.ID, relFile, line.Nr)
			}
		}
	}
	return nil
}

// workspaceRelative converts one JaCoCo source name into the
// workspace-relative slash form the fingerprints expect. jacococli
// records the source by its bare file name as embedded in the class
// debug info (t.kt for a root-level file), which resolves against the
// workspace root; names that would escape the workspace return the empty
// string and are dropped.
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
