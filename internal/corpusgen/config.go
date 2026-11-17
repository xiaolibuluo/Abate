/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

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
)

// errInvalidPlanRef reports a plan document that cannot drive generation.
var errInvalidPlanRef = errors.New("corpusgen: invalid plan reference")

// canonicalLanguages is the exact language order declared by the frozen plan.
var canonicalLanguages = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// planDocument is the read-only wire format of ci/e2e_plan.json. Only fields
// used by the generator are decoded; unknown fields are deliberately ignored.
type planDocument struct {
	Seed              uint64   `json:"seed"`
	Languages         []string `json:"languages"`
	FilesPerLanguage  int      `json:"filesPerLanguage"`
	MinFileBytes      int64    `json:"minFileBytes"`
	MaxFileBytes      int64    `json:"maxFileBytes"`
	StrataCount       int      `json:"strataCount"`
	BrokenPerLanguage int      `json:"brokenPerLanguage"`
	Operators         []string `json:"operators"`
}

// PlanRef is the frozen, digest-bound subset of the E2E plan used by the
// generator. It is a value type with no mutable global state.
type PlanRef struct {
	Seed              uint64
	Languages         []string
	FilesPerLanguage  int
	MinFileBytes      int
	MaxFileBytes      int
	StrataCount       int
	BrokenPerLanguage int
	Operators         []string
	PlanSHA256        string
}

// LoadPlanRef reads the plan as an immutable file and binds its raw-byte
// digest to the returned value. No environment variable or wall clock is used.
func LoadPlanRef(path string) (PlanRef, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return PlanRef{}, fmt.Errorf("%w: read %s: %w", errInvalidPlanRef, path, err)
	}
	var raw planDocument
	if err = json.Unmarshal(data, &raw); err != nil {
		return PlanRef{}, fmt.Errorf("%w: decode %s: %w", errInvalidPlanRef, path, err)
	}
	sum := sha256.Sum256(data)
	ref := PlanRef{
		Seed:              raw.Seed,
		Languages:         slices.Clone(raw.Languages),
		FilesPerLanguage:  raw.FilesPerLanguage,
		MinFileBytes:      int(raw.MinFileBytes),
		MaxFileBytes:      int(raw.MaxFileBytes),
		StrataCount:       raw.StrataCount,
		BrokenPerLanguage: raw.BrokenPerLanguage,
		Operators:         slices.Clone(raw.Operators),
		PlanSHA256:        hex.EncodeToString(sum[:]),
	}
	slices.Sort(ref.Operators)
	if err = ref.validate(); err != nil {
		return PlanRef{}, err
	}
	return ref, nil
}

// validate enforces the frozen plan contract without consulting any source
// metadata. A mismatched language order is rejected, never skipped.
func (p PlanRef) validate() error {
	if p.Seed == 0 || p.FilesPerLanguage < 1 || p.StrataCount < 1 ||
		p.MinFileBytes < 1 || p.MaxFileBytes < p.MinFileBytes ||
		p.BrokenPerLanguage < 1 {
		return fmt.Errorf("%w: missing or invalid budget", errInvalidPlanRef)
	}
	if !slices.Equal(p.Languages, canonicalLanguages) {
		return fmt.Errorf("%w: language order is not canonical", errInvalidPlanRef)
	}
	if len(p.Operators) == 0 {
		return fmt.Errorf("%w: operator set must be non-empty", errInvalidPlanRef)
	}
	return nil
}
