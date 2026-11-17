/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import "abate/internal/e2e/metrics"

type (
	// Report is the complete deterministic E2E result. Field order fixes the
	// JSON byte order; every slice is sorted; there is no map, no wall clock,
	// no absolute path and no runtime version string. The four SHA-256 chains
	// bind the plan, the manifest, the thresholds and the drilldown file.
	Report struct {
		SchemaVersion            int                    `json:"schemaVersion"`
		ToolVersion              string                 `json:"toolVersion"`
		PlanSHA256               string                 `json:"planSha256"`
		ManifestSHA256           string                 `json:"manifestSha256"`
		ThresholdsSHA256         string                 `json:"thresholdsSha256"`
		DrilldownSHA256          string                 `json:"drilldownSha256"`
		Seed                     uint64                 `json:"seed"`
		Languages                []string               `json:"languages"`
		MaxFileBytes             int64                  `json:"maxFileBytes"`
		ParseTimeoutMillis       int64                  `json:"parseTimeoutMillis"`
		MemLimitBytes            uint64                 `json:"memLimitBytes"`
		MaxRuns                  int                    `json:"maxRuns"`
		Merge                    string                 `json:"merge"`
		Columns                  metrics.Columns        `json:"columns"`
		Precision                []metrics.PrecisionRow `json:"precision"`
		Mutation                 []metrics.MutationRow  `json:"mutation"`
		Robustness               metrics.Robustness     `json:"robustness"`
		Equivalence              metrics.Equivalence    `json:"equivalence"`
		Localizers               []string               `json:"localizers"`
		CompetitorAbsolute       string                 `json:"competitorAbsolute"`
		CoverageStatus           []string               `json:"coverageStatus"`
		Facts                    ManifestFacts          `json:"facts"`
		Counters                 []Counter              `json:"counters"`
		BaselineValues           []BaselineValue        `json:"baselineValues"`
		ProcessedFiles           int                    `json:"processedFiles"`
		RegisteredFiles          int                    `json:"registeredFiles"`
		IntervalFindings         int                    `json:"intervalFindings"`
		IncrementalChangedRanges int                    `json:"incrementalChangedRanges"`
		IncrementalConsistent    bool                   `json:"incrementalConsistent"`
	}

	// ManifestFacts is the real SPDX and digest recomputation over the sampled
	// files. It is manifest-level evidence, not a drilldown-derived aggregate.
	ManifestFacts struct {
		RegisteredFiles  int `json:"registeredFiles"`
		SPDXAllowlisted  int `json:"spdxAllowlisted"`
		SPDXBoundary     int `json:"spdxBoundary"`
		SPDXViolations   int `json:"spdxViolations"`
		SHA256Mismatches int `json:"sha256Mismatches"`
	}
)
