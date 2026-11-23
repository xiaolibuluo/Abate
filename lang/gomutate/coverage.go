/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// coverBlock is one parsed row of a coverprofile: the file it describes in
// workspace-relative form, its inclusive start and end line and the number
// of times the block executed during one test run.
type coverBlock struct {
	file      string
	startLine int
	endLine   int
	count     int
}

// coverMatrix stores, for every discovered test, the exact set of executed
// coverage blocks of that test's own run. The per-test association is the
// measured fact the backend relies on: a test covers a token only when the
// block list of that very test contains the token's line.
type coverMatrix struct {
	// byTest maps a test identifier to the blocks its run executed.
	byTest map[string][]coverBlock
}

// measureCoverage runs every discovered test with -coverprofile against the
// pristine workspace and folds each profile under its own test identifier.
// The profiles land in a private temporary directory removed afterwards, so
// a run never leaves coverage artifacts inside the user's workspace.
func measureCoverage(ctx context.Context, cfg runConfig, root, module string,
	tests []testRef,
) (*coverMatrix, error) {
	dir, err := os.MkdirTemp("", "abate-gomutate-cov-")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCoverage, err)
	}
	defer func() { _ = os.RemoveAll(dir) }()
	m := &coverMatrix{byTest: make(map[string][]coverBlock, len(tests))}
	for i, test := range tests {
		profile := filepath.Join(dir, "cov-"+strconv.Itoa(i)+".out")
		if err := runCoveringTest(ctx, cfg, test, profile); err != nil {
			return nil, err
		}
		blocks, err := parseProfile(profile, module)
		if err != nil {
			return nil, err
		}
		m.byTest[test.ID] = blocks
	}
	return m, nil
}

// runCoveringTest executes exactly one test with coverage instrumentation
// and writes its profile to the given path. A pristine-workspace failure is
// fatal: judging mutants against an already-red baseline would fabricate
// kills, which this backend never does.
func runCoveringTest(ctx context.Context, cfg runConfig, test testRef, profile string) error {
	passed, output, err := runGoTest(ctx, cfg, test.Dir,
		"-run", "^"+test.Name+"$", "-coverprofile", profile)
	if err != nil {
		return fmt.Errorf("%w: %s: %v", ErrCoverage, test.ID, err)
	}
	if !passed {
		return fmt.Errorf("%w: %s fails on the pristine workspace: %s", ErrCoverage, test.ID, tail(output))
	}
	return nil
}

// parseProfile reads one coverprofile and converts every executed block
// into a coverBlock with a workspace-relative file path. Blocks with zero
// execution count are dropped: they cannot witness a kill.
func parseProfile(path, module string) ([]coverBlock, error) {
	data, err := os.ReadFile(path) //nolint:gosec // the path is created inside our own temp directory.
	if err != nil {
		return nil, fmt.Errorf("%w: read %s: %v", ErrCoverage, path, err)
	}
	out := make([]coverBlock, 0, 32)
	for i, line := range strings.Split(string(data), "\n") {
		if i == 0 || line == "" {
			continue
		}
		block, ok := parseBlockLine(line, module)
		if !ok {
			return nil, fmt.Errorf("%w: malformed profile line %d of %s", ErrCoverage, i+1, path)
		}
		if block.count > 0 {
			out = append(out, block)
		}
	}
	return out, nil
}

// parseBlockLine decodes one "file.go:startLine.col,endLine.col stmts count"
// coverprofile row and rewrites the file field from module-path form into
// the workspace-relative slash form the scanner reports.
func parseBlockLine(line, module string) (coverBlock, bool) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return coverBlock{}, false
	}
	count, err := strconv.Atoi(fields[2])
	if err != nil {
		return coverBlock{}, false
	}
	colon := strings.LastIndexByte(fields[0], ':')
	if colon < 0 {
		return coverBlock{}, false
	}
	file := strings.TrimPrefix(fields[0][:colon], module+"/")
	startPos, endPos, ok := strings.Cut(fields[0][colon+1:], ",")
	if !ok {
		return coverBlock{}, false
	}
	start, ok1 := lineOf(startPos)
	end, ok2 := lineOf(endPos)
	if !ok1 || !ok2 {
		return coverBlock{}, false
	}
	return coverBlock{file: file, startLine: start, endLine: end, count: count}, true
}

// lineOf extracts the 1-based line number from a "line.column" position.
func lineOf(pos string) (int, bool) {
	dot := strings.IndexByte(pos, '.')
	if dot <= 0 {
		return 0, false
	}
	line, err := strconv.Atoi(pos[:dot])
	return line, err == nil && line > 0
}

// testCovers reports whether the given test's own measured blocks contain
// the given line of the given file.
func (m *coverMatrix) testCovers(testID, file string, line int) bool {
	for _, block := range m.byTest[testID] {
		if block.file == file && block.startLine <= line && line <= block.endLine {
			return true
		}
	}
	return false
}

// covering returns how many discovered tests execute the given line of the
// given file, counted over the per-test measured block sets.
func (m *coverMatrix) covering(file string, line int) uint64 {
	var n uint64
	for _, blocks := range m.byTest {
		for _, block := range blocks {
			if block.file == file && block.startLine <= line && line <= block.endLine {
				n++
				break
			}
		}
	}
	return n
}

// files returns the covered file paths in ascending order, used for the
// deterministic fingerprint of the suite.
func (m *coverMatrix) files() []string {
	seen := make(map[string]bool, 32)
	for _, blocks := range m.byTest {
		for _, block := range blocks {
			seen[block.file] = true
		}
	}
	out := make([]string, 0, len(seen))
	for file := range seen {
		out = append(out, file)
	}
	slices.Sort(out)
	return out
}
