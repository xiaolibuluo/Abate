/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// corpusCase is one conformance sample from testdata/manifest.json.
type corpusCase struct {
	File   string `json:"file"`
	Accept bool   `json:"accept"`
	Rule   string `json:"rule"`
	Format string `json:"format"`
}

// loadCorpus reads and validates the sample manifest.
func loadCorpus(t *testing.T) []corpusCase {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("testdata", "manifest.json"))
	requireOK(t, err)
	var cases []corpusCase
	requireOK(t, json.Unmarshal(raw, &cases))
	if len(cases) != 100 {
		t.Fatalf("corpus size = %d, want 100", len(cases))
	}
	valid, invalid := 0, 0
	for _, c := range cases {
		if c.Accept {
			valid++
		} else {
			invalid++
			if c.Rule == "" {
				t.Fatalf("invalid sample %s has no expected rule", c.File)
			}
		}
	}
	if valid != 60 || invalid != 40 {
		t.Fatalf("corpus split = %d valid / %d invalid, want 60/40", valid, invalid)
	}
	return cases
}

// TestCorpusConformance asserts every sample's decision and exact rule id.
func TestCorpusConformance(t *testing.T) {
	t.Parallel()
	cases := loadCorpus(t)
	for _, c := range cases {
		data, err := os.ReadFile(filepath.Join("testdata", c.File))
		requireOK(t, err)
		fmt := FormatAuto
		if c.Format != "" && c.Format != string(FormatAuto) {
			fmt = Format(c.Format)
		}
		exp, err := runChain(data, fmt)
		if c.Accept {
			if err != nil {
				t.Fatalf("valid sample %s rejected: %v", c.File, err)
			}
			if exp == nil {
				t.Fatalf("valid sample %s produced a nil experiment", c.File)
			}
			continue
		}
		requireRule(t, err, c.Rule)
	}
}

// TestCorpusRuleCoverage proves the corpus exercises every failure rule.
func TestCorpusRuleCoverage(t *testing.T) {
	t.Parallel()
	cases := loadCorpus(t)
	seen := make(map[string]bool)
	for _, c := range cases {
		if !c.Accept {
			seen[c.Rule] = true
		}
	}
	want := []string{
		RuleUnknownField, RuleUnsupportedVersion, RuleInvalidFormat,
		RuleEmptyDocument, RuleMalformedDocument, RuleMissingField,
		RuleDanglingProbeRef, RuleScheduleExclusive, RuleExprSyntax,
		RuleExprDepth, RuleExprUnknownOp, RuleExprNotInvertible,
	}
	for _, r := range want {
		if !seen[r] {
			t.Fatalf("corpus never exercises rule %s", r)
		}
	}
	keys := make([]string, 0, len(seen))
	for k := range seen {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	if len(keys) != len(want) {
		t.Fatalf("unexpected rule coverage set: %v", keys)
	}
}

// TestCorpusDeterministicReplay replays all valid samples twice identically.
func TestCorpusDeterministicReplay(t *testing.T) {
	t.Parallel()
	cases := loadCorpus(t)
	for _, c := range cases {
		if !c.Accept {
			continue
		}
		data, err := os.ReadFile(filepath.Join("testdata", c.File))
		requireOK(t, err)
		first, err := runChain(data, FormatAuto)
		requireOK(t, err)
		second, err := runChain(data, FormatAuto)
		requireOK(t, err)
		if digestOf(t, first) != digestOf(t, second) {
			t.Fatalf("sample %s is not deterministic", c.File)
		}
	}
}
