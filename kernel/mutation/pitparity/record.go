/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Sentinel errors of the parity harness.
var (
	// ErrReportNotFound reports that the vendored PIT report cannot be located.
	ErrReportNotFound = errors.New("pitparity: PIT report not found")

	// ErrInvalidReport reports a malformed PIT mutation report.
	ErrInvalidReport = errors.New("pitparity: invalid PIT report")
)

// The mutation statuses reported by PITest 1.19.6.
const (
	// StatusKilled is the PIT status of a mutant killed by the suite.
	StatusKilled = "KILLED"

	// StatusSurvived is the PIT status of a mutant that survived the suite.
	StatusSurvived = "SURVIVED"

	// StatusNoCoverage is the PIT status of a mutant no test reaches.
	StatusNoCoverage = "NO_COVERAGE"
)

// ReportRelativePath is the repository relative path of the vendored report.
const ReportRelativePath = "kernel/mutation/testdata/pit/pit-mutations.json"

// Mutation is one PIT mutation record of the official PIT report.
//
// 参考 PITest 的 MutationDetails（字节码位置 + 算子 + 状态），改进点=解析为
// 语言无关的纯数据记录，保留行号与击杀测试，便于逐点对拍与确定性渲染。
type Mutation struct {
	// Status is the PIT verdict of the mutant.
	Status string `json:"status"`

	// Detected is the PIT detection flag rendered as "true" or "false".
	Detected string `json:"detected"`

	// SourceFile is the mutated source file of the baseline project.
	SourceFile string `json:"sourceFile"`

	// MutatedMethod is the mutated method of the baseline project.
	MutatedMethod string `json:"mutatedMethod"`

	// LineNumber is the mutated line inside the baseline source file.
	LineNumber int `json:"lineNumber"`

	// Mutator is the PIT mutator family that produced the mutant.
	Mutator string `json:"mutator"`

	// Description is the human readable description of the mutation.
	Description string `json:"description"`

	// KillingTest is the test that killed the mutant, empty when none did.
	KillingTest string `json:"killingTest"`
}

// ParseMutations decodes a PIT mutation report and validates every record.
//
// The decoder is the standard library JSON codec, so the harness stays offline
// and free of third party dependencies.
func ParseMutations(data []byte) ([]Mutation, error) {
	var records []Mutation
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidReport, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("%w: report holds no mutation", ErrInvalidReport)
	}
	for i := range records {
		if err := records[i].validate(); err != nil {
			return nil, fmt.Errorf("%w: record %d: %w", ErrInvalidReport, i, err)
		}
	}
	return records, nil
}

// LoadMutations reads and decodes the PIT mutation report stored at path.
func LoadMutations(path string) ([]Mutation, error) {
	//nolint:gosec // the report path is supplied explicitly by the caller.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("pitparity: read %s: %w", path, err)
	}
	return ParseMutations(data)
}

// LocateReport walks up from start and returns the path of the vendored report.
//
// The report lives in the parent package testdata directory, which the
// //go:embed directive cannot reach because it forbids parent directory
// patterns, so the harness resolves it on disk instead.
func LocateReport(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrReportNotFound, err)
	}
	for {
		candidate := filepath.Join(dir, filepath.FromSlash(ReportRelativePath))
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("%w: %s", ErrReportNotFound, ReportRelativePath)
		}
		dir = parent
	}
}

// validate reports whether one decoded record is well formed.
func (m Mutation) validate() error {
	switch m.Status {
	case StatusKilled, StatusSurvived, StatusNoCoverage:
	default:
		return fmt.Errorf("unknown status %q", m.Status)
	}
	if m.MutatedMethod == "" {
		return errors.New("empty mutated method")
	}
	if m.Mutator == "" {
		return errors.New("empty mutator")
	}
	if m.LineNumber <= 0 {
		return fmt.Errorf("invalid line number %d", m.LineNumber)
	}
	return nil
}
