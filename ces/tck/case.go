/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"slices"
	"strings"

	"abate/internal/yamlmini"
)

//go:embed cases
var caseFS embed.FS

// Case is one data-driven TCK case.
//
// It is serialized as YAML under cases/ and never references any implementation
// type, so a third party can read the same file to learn what compliance means.
type Case struct {
	// ID is the case identifier, e.g. TCK-PARSE-001.
	ID string `yaml:"id"`
	// Spec cites the specification clause the case derives from.
	Spec string `yaml:"spec"`
	// Kind selects the driver method and the expected shape.
	Kind CaseKind `yaml:"kind"`
	// Input is the CES document bytes (YAML or JSON) under test.
	Input string `yaml:"input"`
	// Expect is the neutral expected outcome.
	Expect CaseExpect `yaml:"expect"`
}

// LoadCases reads every YAML case under the embedded library.
//
// Cases are sorted by ID (then kind) so the running order is deterministic and
// the report is stable across runs and platforms.
func LoadCases() ([]Case, error) { return loadCasesFrom(caseFS) }

// loadCasesFrom reads the case library from fsys. Injecting the filesystem keeps the
// embedded library read-only while making the defensive branches (missing root,
// unreadable groups, malformed YAML, invalid cases, tie-break ordering) directly
// testable with an in-memory fs.FS instead of corrupting real case data.
func loadCasesFrom(fsys fs.FS) ([]Case, error) {
	entries, err := fs.ReadDir(fsys, "cases")
	if err != nil {
		return nil, fmt.Errorf("tck: read case library: %w", err)
	}
	var out []Case
	for _, dir := range entries {
		if !dir.IsDir() {
			continue
		}
		group := "cases/" + dir.Name()
		files, err := fs.ReadDir(fsys, group)
		if err != nil {
			return nil, fmt.Errorf("tck: read %s: %w", group, err)
		}
		for _, f := range files {
			if f.IsDir() || !strings.HasSuffix(f.Name(), ".yaml") {
				continue
			}
			raw, err := fs.ReadFile(fsys, group+"/"+f.Name())
			if err != nil {
				return nil, fmt.Errorf("tck: read %s/%s: %w", group, f.Name(), err)
			}
			var c Case
			if err := yamlmini.Unmarshal(raw, &c); err != nil {
				return nil, fmt.Errorf("tck: decode %s/%s: %w", group, f.Name(), err)
			}
			if err := validateCase(c); err != nil {
				return nil, fmt.Errorf("tck: %s/%s: %w", group, f.Name(), err)
			}
			out = append(out, c)
		}
	}
	slices.SortFunc(out, func(a, b Case) int {
		if c := strings.Compare(a.ID, b.ID); c != 0 {
			return c
		}
		return strings.Compare(string(a.Kind), string(b.Kind))
	})
	return out, nil
}

// validateCase enforces the data-quality invariants required by the
// specification: every case has an ID, a spec citation, a known kind and a
// non-empty input.
func validateCase(c Case) error {
	if c.ID == "" {
		return errors.New("case missing id")
	}
	if c.Spec == "" {
		return fmt.Errorf("case %s missing spec citation", c.ID)
	}
	switch c.Kind {
	case KindParse, KindValidate, KindRoundtrip:
	default:
		return fmt.Errorf("case %s has unknown kind %q", c.ID, c.Kind)
	}
	if strings.TrimSpace(c.Input) == "" {
		return fmt.Errorf("case %s missing input", c.ID)
	}
	return nil
}
