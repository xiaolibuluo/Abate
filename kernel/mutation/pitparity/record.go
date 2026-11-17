/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Sentinel errors of the golden regression harness.
var (
	// ErrGoldenNotFound reports that the self-developed golden cannot be located.
	ErrGoldenNotFound = errors.New("pitparity: self-developed golden not found")

	// ErrInvalidGolden reports a malformed self-developed golden.
	ErrInvalidGolden = errors.New("pitparity: invalid self-developed golden")
)

// GoldenRelativePath is the repository relative path of the frozen golden.
const GoldenRelativePath = "kernel/mutation/testdata/pit/golden/mutation-semantics.golden"

// The expected verdict classifications frozen in the self-developed golden.
//
// These are local observation labels of the Abate fixture, not status codes of
// any external mutation tool.
const (
	// ExpectKilled marks a case the suite distinguishes from the original.
	ExpectKilled = "KILLED"

	// ExpectSurvived marks a case the suite cannot distinguish.
	ExpectSurvived = "SURVIVED"

	// ExpectNoCoverage marks a case no test of the suite reaches.
	ExpectNoCoverage = "NO_COVERAGE"
)

// Case is one self-developed golden case of the regression harness.
//
// 本仓自研判例：方法/算子标识只作本地夹具用例编号，期望分类由本仓语义见证派生，
// 不是任何上游工具的产物字段。
type Case struct {
	// Method is the local fixture method identifier of the case.
	Method string `json:"method"`

	// Mutator is the local fixture operator identifier of the case.
	Mutator string `json:"mutator"`

	// Line is the local fixture line identifier of the case.
	Line int `json:"line"`

	// Expect is the frozen expected verdict classification of the case.
	Expect string `json:"expect"`
}

// goldenFields is the tab separated field count of one outcome line.
const goldenFields = 12

// goldenPrefix marks an outcome line inside the golden file.
const goldenPrefix = "outcome"

// ParseGolden decodes the frozen golden into self-developed cases.
//
// The decoder uses the standard library only, so the regression stays offline
// and free of third party dependencies. Comment and header lines are skipped.
func ParseGolden(data []byte) ([]Case, error) {
	var cases []Case
	for i, raw := range strings.Split(string(data), "\n") {
		line := strings.TrimSuffix(raw, "\r")
		if line == "" || strings.HasPrefix(line, "#") ||
			!strings.HasPrefix(line, goldenPrefix+"\t") {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) != goldenFields {
			return nil, fmt.Errorf("%w: line %d has %d fields, want %d",
				ErrInvalidGolden, i+1, len(fields), goldenFields)
		}
		number, err := strconv.Atoi(strings.TrimSpace(fields[3]))
		if err != nil {
			return nil, fmt.Errorf("%w: line %d line number: %w", ErrInvalidGolden, i+1, err)
		}
		entry := Case{
			Method:  fields[2],
			Mutator: fields[4],
			Line:    number,
			Expect:  fields[5],
		}
		if err := entry.validate(); err != nil {
			return nil, fmt.Errorf("%w: line %d: %w", ErrInvalidGolden, i+1, err)
		}
		cases = append(cases, entry)
	}
	if len(cases) == 0 {
		return nil, fmt.Errorf("%w: golden holds no case", ErrInvalidGolden)
	}
	return cases, nil
}

// LoadGolden reads and decodes the frozen golden stored at path.
func LoadGolden(path string) ([]Case, error) {
	//nolint:gosec // the golden path is supplied explicitly by the caller.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("pitparity: read %s: %w", path, err)
	}
	return ParseGolden(data)
}

// LocateGolden walks up from start and returns the path of the frozen golden.
//
// The golden lives in the parent package testdata directory, which the
// //go:embed directive cannot reach because it forbids parent directory
// patterns, so the harness resolves it on disk instead.
func LocateGolden(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGoldenNotFound, err)
	}
	for {
		candidate := filepath.Join(dir, filepath.FromSlash(GoldenRelativePath))
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("%w: %s", ErrGoldenNotFound, GoldenRelativePath)
		}
		dir = parent
	}
}

// validate reports whether one decoded case is well formed.
func (c Case) validate() error {
	switch c.Expect {
	case ExpectKilled, ExpectSurvived, ExpectNoCoverage:
	default:
		return fmt.Errorf("unknown expected classification %q", c.Expect)
	}
	if c.Method == "" {
		return errors.New("empty method")
	}
	if c.Mutator == "" {
		return errors.New("empty mutator")
	}
	if c.Line <= 0 {
		return fmt.Errorf("invalid line number %d", c.Line)
	}
	return nil
}

// RenderGolden renders one report as the canonical frozen golden bytes.
//
// The rendering holds no timestamp and no map iteration order, so two runs over
// the same self-developed cases are byte identical.
func RenderGolden(report Report) []byte {
	var b strings.Builder
	b.WriteString("# Abate 自研变异语义冻结 golden（历史观测值，非上游官方产物）\n")
	b.WriteString("# 由 kernel/mutation/pitparity/gen_golden.go 用本仓引擎与同构夹具生成，可复跑。\n")
	_, _ = fmt.Fprintf(&b, "total=%d\ncomparable=%d\nconsistent=%d\nrate=%s\n",
		report.Total, report.Comparable, report.Consistent, report.Rate)
	for i, outcome := range report.Outcomes {
		_, _ = fmt.Fprintf(&b, "outcome\t%d\t%s\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%v\n",
			i+1, outcome.Method, outcome.Line, outcome.Mutator,
			outcome.Expect, outcome.Node, outcome.Class,
			outcome.AbateVerdict, outcome.Witness, outcome.Disposition,
			outcome.Consistent)
	}
	return []byte(b.String())
}

// RegenerateGolden renders the golden from self-developed tables and the real
// mutation engine, with no external artifact as input.
func RegenerateGolden() ([]byte, error) {
	cases, err := buildCases()
	if err != nil {
		return nil, err
	}
	report, err := Compare(cases)
	if err != nil {
		return nil, err
	}
	return RenderGolden(report), nil
}
