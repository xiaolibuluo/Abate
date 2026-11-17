/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package neutrality

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Reverse fixtures are assembled by concatenation so the repository's own
// textual scans never mistake a fixture for a production forbidden construct.
// Each fixture is still a syntactically complete Go source parsed by go/parser.
var (
	fixtureSourceBranch = "package fixture\n\ntype E struct { Source string }\n" +
		"\nfunc f(e E) {\n if e.Source == \"validation\" { return }\n}\n"
	fixtureCorpusSkip = "package fixture\n\nimport \"abate/internal/corpus\"\n" +
		"\nfunc f(m corpus.Manifest, p corpus.SamplePlan) {\n" +
		" for _, e := range corpus.StratifiedSample(m, p) {\n" +
		"  if e.Source == \"validation\" { continue }\n  _ = e\n }\n}\n"
	fixtureCorpusExclusion = "package fixture\n\nimport \"abate/internal/corpus\"\n" +
		"\nfunc f(m corpus.Manifest) int {\n" +
		" for _, e := range m.Entries {\n" +
		"  if e.Source == \"x\" { return 0 }\n }\n return 1\n}\n"
	fixtureTypeCheck = "package fixture\n\ntype G struct{}\n" +
		"\nfunc (G) WithTypeCheck(enabled bool) G { return G{} }\n" +
		"\nfunc f(g G) { _ = g.WithTypeCheck(" + "false) }\n"
	fixtureRecover = "package fixture\n\nfunc f() { _ = rec" + "over() }\n"
	fixtureNow     = "package fixture\n\nimport \"time\"\n" +
		"\nfunc f() { _ = time" + ".Now() }\n"
	fixtureHTTP = "package fixture\n\nimport \"net/http\"\n" +
		"\nfunc f(req *http.Request) { _, _ = http.DefaultClient.Do(req) }\n"
	fixtureDial = "package fixture\n\nfunc f() { _, _ = net" +
		".Dial(\"tcp\", \"x\") }\n"
	fixtureExit = "package fixture\n\nimport \"os\"\n" +
		"\nfunc f() { os.Exit(1) }\n"
	fixtureFlag  = "package fixture\n\nfunc f(enabled bool) int { return 1 }\n"
	fixtureClean = "package fixture\n\nimport \"abate/internal/corpus\"\n" +
		"\nfunc f(m corpus.Manifest, p corpus.SamplePlan) int {\n" +
		" n := 0\n for _, e := range corpus.StratifiedSample(m, p) {\n" +
		"  n += len(e.RelPath)\n }\n return n\n}\n"
)

// fixtureCase pairs one reverse fixture with the rule it must trigger.
type fixtureCase struct {
	name string
	rule string
	src  string
}

// reverseCases is the fixed table of reverse fixtures.
var reverseCases = []fixtureCase{
	{"source_branch", ruleSourceBranch, fixtureSourceBranch},
	{"corpus_loop_skip", ruleCorpusSkip, fixtureCorpusSkip},
	{"corpus_exclusion", ruleCorpusExclusion, fixtureCorpusExclusion},
	{"typecheck_bypass", ruleTypeCheck, fixtureTypeCheck},
	{"recover_swallow", ruleRecover, fixtureRecover},
	{"time_now", ruleDeterminism, fixtureNow},
	{"http_do", ruleDeterminism, fixtureHTTP},
	{"net_dial", ruleDeterminism, fixtureDial},
	{"deep_exit", ruleDeepExit, fixtureExit},
	{"flag_parameter", ruleFlagParameter, fixtureFlag},
}

// TestReverseFixturesCatchEveryInvariant proves all eight syntax-level rules
// reject a real go/ast counterexample, including provenance exclusion, corpus
// loop continue, type-check bypass, recover, wall clock and network calls.
func TestReverseFixturesCatchEveryInvariant(t *testing.T) {
	t.Parallel()
	analyzer := NewAnalyzer(".")
	for _, tc := range reverseCases {
		res, err := analyzer.CheckSource(tc.name+".go", []byte(tc.src))
		if err != nil {
			t.Fatalf("fixture %s failed to parse: %v", tc.name, err)
		}
		if !hasRule(res.Findings, tc.rule) {
			t.Fatalf("fixture %s did not trigger %s: %+v",
				tc.name, tc.rule, res.Findings)
		}
	}
}

// TestCleanFixtureHasNoFindings proves a corpus loop without skips or
// provenance branches passes the eight invariants with zero findings.
func TestCleanFixtureHasNoFindings(t *testing.T) {
	t.Parallel()
	res, err := NewAnalyzer(".").CheckSource("clean.go",
		[]byte(fixtureClean))
	if err != nil {
		t.Fatalf("clean fixture failed to parse: %v", err)
	}
	if len(res.Findings) != 0 {
		t.Fatalf("clean fixture produced findings: %+v", res.Findings)
	}
	if res.ProcessedFiles != res.RegisteredFiles {
		t.Fatalf("clean fixture processed %d of %d",
			res.ProcessedFiles, res.RegisteredFiles)
	}
}

// TestEightRulesRegistered proves the invariant table is complete.
func TestEightRulesRegistered(t *testing.T) {
	t.Parallel()
	if len(ruleTable) != 8 {
		t.Fatalf("rule table has %d rules, want 8", len(ruleTable))
	}
}

// TestRegisteredSourcesAreFullyProcessed scans the real completeness anchor
// with zero exemptions and asserts every registered .go file was parsed and
// no invariant was violated.
func TestRegisteredSourcesAreFullyProcessed(t *testing.T) {
	t.Parallel()
	root := filepath.Join("..", "..", "..")
	data, err := fs.ReadFile(os.DirFS(root), "ci/e2e_sources.txt")
	if err != nil {
		t.Fatalf("cannot read source anchor: %v", err)
	}
	paths := make([]string, 0)
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			paths = append(paths, trimmed)
		}
	}
	res, err := NewAnalyzer(root).Check(paths)
	if err != nil {
		t.Fatalf("registered source analysis failed: %v", err)
	}
	if res.ProcessedFiles != res.RegisteredFiles {
		t.Fatalf("ProcessedFiles=%d RegisteredFiles=%d",
			res.ProcessedFiles, res.RegisteredFiles)
	}
	if len(res.Findings) != 0 {
		t.Fatalf("zero-exemption findings: %+v", res.Findings)
	}
}

// hasRule reports whether a finding set contains one rule.
func hasRule(findings []Finding, rule string) bool {
	for _, finding := range findings {
		if finding.Rule == rule {
			return true
		}
	}
	return false
}
