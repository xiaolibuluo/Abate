/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import "abate/internal/corpus"

// Fixed strings and file names. ToolVersion is a build-time constant, never a
// runtime version string, so two runs of the same seed emit identical bytes.
const (
	ToolVersion           = "abate-e2e/1.0.0-fixed"
	CompetitorNotMeasured = "NOT_MEASURED"
	MergeLabel            = "merge=1"
	DefaultCorpusRoot     = corpus.DefaultCorpusRoot
	ThresholdsName        = "e2e_thresholds.json"
	BaselineName          = "e2e_thresholds_baseline.json"
	ConstructsName        = "constructs_go.json"
	EquivalentName        = "equivalent_mutations.json"
	SourcesName           = "e2e_sources.txt"
	TraceFileName         = "e2e_trace.jsonl"
	ReportSchemaVersion   = 1
	FilePerm              = 0o600
	DirPerm               = 0o750
)

// libFuzzer-aligned budgets. Every large literal lives in a named constant so
// the constant gate never sees a bare production magic number.
const (
	defaultFilesPerLanguage   = 200
	defaultStrataCount        = 8
	defaultMinFileBytes       = 256
	defaultMaxFileBytes       = 180 * corpus.Kib
	defaultParseTimeoutMillis = 5000
	defaultMemLimitBytes      = 256 * corpus.Mib
	defaultMaxRuns            = 200000
	defaultBrokenPerLanguage  = 200
	maxEditsPerOperator       = 1
	damagedBudgetExtra        = 8
	traceFlushInterval        = 64
	notMeasuredPct            = 0
	fullCoveragePct           = 100
)

// languageOrder is the canonical report order of the seven frontends.
var languageOrder = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// dialectTable maps each language tag onto the dialect handed to its frontend.
var dialectTable = map[string]string{
	"go": "", "py": "", "rs": "", "js": "", "c": "c11", "cs": "cs12", "kt": "kt2.1",
}

// defaultOperators is the sorted identifier set declared by ci/e2e_plan.json.
var defaultOperators = []string{
	"block", "call", "cmp", "cond", "digit", "literal", "op", "return",
}
