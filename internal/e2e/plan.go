/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"abate/internal/corpus"
)

// errInvalidPlan reports a plan that cannot drive a deterministic run.
var errInvalidPlan = errors.New("e2e: invalid plan")

// Plan is the complete configuration of one end-to-end run. The first block is
// decoded verbatim from ci/e2e_plan.json; the second block is derived from the
// plan file location and is never serialized into the compared report.
type Plan struct {
	Seed                     uint64   `json:"seed"`
	Languages                []string `json:"languages"`
	FilesPerLanguage         int      `json:"filesPerLanguage"`
	MinFileBytes             int64    `json:"minFileBytes"`
	MaxFileBytes             int64    `json:"maxFileBytes"`
	StrataCount              int      `json:"strataCount"`
	ParseTimeoutMillis       int64    `json:"parseTimeoutMillis"`
	MemLimitBytes            uint64   `json:"memLimitBytes"`
	MaxRuns                  int      `json:"maxRuns"`
	BrokenPerLanguage        int      `json:"brokenPerLanguage"`
	Operators                []string `json:"operators"`
	ReportDir                string   `json:"reportDir"`
	ThresholdsBaselineSHA256 string   `json:"thresholdsBaselineSha256"`
	PlanSHA256               string   `json:"-"`
	CorpusRoot               string   `json:"-"`
	RepoRoot                 string   `json:"-"`
	ThresholdsPath           string   `json:"-"`
	BaselinePath             string   `json:"-"`
	ConstructsPath           string   `json:"-"`
	EquivalentPath           string   `json:"-"`
	SourcesPath              string   `json:"-"`
}

// LoadPlan decodes a plan, validates it, and derives every sibling path from
// the plan directory. The digest covers the raw file bytes, so gate_e2e.sh can
// recompute it offline without any Go code.
func LoadPlan(path string) (Plan, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return Plan{}, fmt.Errorf("%w: read %s: %w", errInvalidPlan, path, err)
	}
	var decoded Plan
	if err = json.Unmarshal(data, &decoded); err != nil {
		return Plan{}, fmt.Errorf("%w: decode %s: %w", errInvalidPlan, path, err)
	}
	out := normalizePlan(decoded)
	sum := sha256.Sum256(data)
	out.PlanSHA256 = hex.EncodeToString(sum[:])
	dir := filepath.Dir(path)
	out.CorpusRoot = DefaultCorpusRoot
	out.RepoRoot = "."
	out.ThresholdsPath = filepath.Join(dir, ThresholdsName)
	out.BaselinePath = filepath.Join(dir, BaselineName)
	out.ConstructsPath = filepath.Join(dir, ConstructsName)
	out.EquivalentPath = filepath.Join(dir, EquivalentName)
	out.SourcesPath = filepath.Join(dir, SourcesName)
	if err = out.Validate(); err != nil {
		return Plan{}, err
	}
	return out, nil
}

// Validate checks that the fixed seed, the seven canonical languages and a
// non-empty sorted operator set are present. No environment-derived value is
// consulted, so the same plan file always yields the same configuration.
func (p Plan) Validate() error {
	if p.Seed == 0 || p.FilesPerLanguage < 1 || p.StrataCount < 1 ||
		p.MaxFileBytes < p.MinFileBytes || p.ParseTimeoutMillis < 1 ||
		p.MemLimitBytes < 1 || p.MaxRuns < 1 || p.BrokenPerLanguage < 1 {
		return fmt.Errorf("%w: missing budget", errInvalidPlan)
	}
	if !slices.Equal(p.Languages, languageOrder) {
		return fmt.Errorf("%w: language order is not canonical", errInvalidPlan)
	}
	if len(p.Operators) == 0 || !slices.IsSorted(p.Operators) {
		return fmt.Errorf("%w: operators must be sorted and non-empty", errInvalidPlan)
	}
	for _, id := range p.Operators {
		if !slices.Contains(defaultOperators, id) {
			return fmt.Errorf("%w: unregistered operator %s", errInvalidPlan, id)
		}
	}
	return nil
}

// SamplePlan projects the run configuration onto the corpus sampler.
func (p Plan) SamplePlan() corpus.SamplePlan {
	return corpus.SamplePlan{
		FilesPerLanguage: p.FilesPerLanguage,
		StrataCount:      p.StrataCount,
		MinFileBytes:     p.MinFileBytes,
		MaxFileBytes:     p.MaxFileBytes,
	}
}

// Budget projects the libFuzzer-aligned timeout and memory limits.
func (p Plan) Budget() ProbeBudget {
	return ProbeBudget{
		TimeoutMillis: p.ParseTimeoutMillis,
		MemLimitBytes: p.MemLimitBytes,
	}
}

// normalizePlan fills unset fields from the named package constants and sorts
// the operator identifiers so the report order never depends on JSON order.
func normalizePlan(in Plan) Plan {
	out := in
	if out.FilesPerLanguage < 1 {
		out.FilesPerLanguage = defaultFilesPerLanguage
	}
	if out.StrataCount < 1 {
		out.StrataCount = defaultStrataCount
	}
	if out.MinFileBytes < 1 {
		out.MinFileBytes = defaultMinFileBytes
	}
	if out.MaxFileBytes < out.MinFileBytes {
		out.MaxFileBytes = defaultMaxFileBytes
	}
	if out.ParseTimeoutMillis < 1 {
		out.ParseTimeoutMillis = defaultParseTimeoutMillis
	}
	if out.MemLimitBytes < 1 {
		out.MemLimitBytes = defaultMemLimitBytes
	}
	if out.MaxRuns < 1 {
		out.MaxRuns = defaultMaxRuns
	}
	if out.BrokenPerLanguage < 1 {
		out.BrokenPerLanguage = defaultBrokenPerLanguage
	}
	if len(out.Languages) == 0 {
		out.Languages = slices.Clone(languageOrder)
	}
	if len(out.Operators) == 0 {
		out.Operators = slices.Clone(defaultOperators)
	} else {
		out.Operators = slices.Clone(out.Operators)
		slices.Sort(out.Operators)
	}
	return out
}
