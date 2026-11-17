/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"errors"
	"slices"
	"strings"

	"abate/internal/corpus"
	e2ecoverage "abate/internal/e2e/coverage"
	"abate/internal/e2e/metrics"
	"abate/internal/e2e/neutrality"
	"abate/internal/e2e/report"
	"abate/lang/air"
)

// coveragePass joins the real parsed units with the six embedded grammars and
// the self-built Go inventory. A NOT_MEASURED row is reported honestly instead
// of being replaced by a fabricated coverage ratio.
func (r *Runner) coveragePass(states []fileState,
	rows []metrics.PrecisionRow) ([]metrics.PrecisionRow, []string, error) {
	firstUnit := firstUnitPerLanguage(states)
	status := make([]string, 0, len(rows))
	for i := range rows {
		lang := rows[i].Language
		unit := firstUnit[lang]
		rows[i].DegradeNodePct = r.coverageIndex.Degradation(lang, unit)
		cov, err := r.coverageIndex.Coverage(lang, unit)
		if errors.Is(err, e2ecoverage.ErrNotMeasured) {
			rows[i].ConstructCoveragePct = notMeasuredPct
			status = append(status, lang+"="+CompetitorNotMeasured)
			continue
		}
		if err != nil {
			return nil, nil, err
		}
		rows[i].ConstructCoveragePct = cov.Percent()
		status = append(status, lang+"="+r.coverageIndex.Status(lang))
	}
	slices.Sort(status)
	return rows, status, nil
}

// firstUnitPerLanguage picks the first real parsed unit of each language in
// canonical order, so coverage is computed from an executed frontend.
func firstUnitPerLanguage(states []fileState) map[string]*air.Unit {
	out := make(map[string]*air.Unit)
	for i := range states {
		lang := states[i].entry.Language
		if _, found := out[lang]; !found && states[i].unit != nil {
			out[lang] = states[i].unit
		}
	}
	return out
}

// manifestFacts recomputes the SPDX classification and the SHA-256 match of
// every sampled entry from the bytes that were actually read.
func manifestFacts(entries []corpus.Entry, states []fileState) report.ManifestFacts {
	facts := report.ManifestFacts{RegisteredFiles: len(entries)}
	for i := range states {
		switch states[i].spdxClass {
		case "allowlisted":
			facts.SPDXAllowlisted++
		case "boundary":
			facts.SPDXBoundary++
		default:
			facts.SPDXViolations++
		}
		if !states[i].shaOK {
			facts.SHA256Mismatches++
		}
	}
	return facts
}

// counterFacts bundles the zero-tolerance counter inputs.
type counterFacts struct {
	entries  []corpus.Entry
	states   []fileState
	records  []metrics.DrilldownRecord
	summary  metrics.Summary
	findings []neutrality.Finding
	facts    report.ManifestFacts
}

// Counters derives all eleven zero-tolerance observations from real data.
func (i counterFacts) Counters() []report.Counter {
	byRule := make(map[string]int)
	for _, finding := range i.findings {
		byRule[finding.Rule]++
	}
	exclusion := len(i.entries) - len(i.states)
	if exclusion < 0 {
		exclusion = 0
	}
	panicEmpty := 0
	for _, rec := range i.records {
		if rec.OutcomeKind == metrics.OutcomePanic && rec.Localizer() == "" {
			panicEmpty++
		}
	}
	pairs := []report.Counter{
		{Key: "corpusExclusionCount", Value: exclusion},
		{Key: "determinismFailureCount", Value: i.summary.Robustness.DeterminismFailureCount},
		{Key: "equivalenceSilentDrops", Value: i.summary.Equivalence.SilentDrops},
		{Key: "oomCount", Value: i.summary.Robustness.OOMCount},
		{Key: "panicLocalizerEmpty", Value: panicEmpty},
		{Key: "recoverUsageCount", Value: byRule["recover-swallow"]},
		{Key: "sha256Mismatches", Value: i.facts.SHA256Mismatches},
		{Key: "sourceNeutralityHits", Value: len(i.findings)},
		{Key: "spdxViolations", Value: i.facts.SPDXViolations},
		{Key: "timeoutCount", Value: i.summary.Robustness.TimeoutCount},
		{Key: "typeCheckBypassCount", Value: byRule["typecheck-bypass"]},
	}
	slices.SortFunc(pairs, func(a, b report.Counter) int {
		return strings.Compare(a.Key, b.Key)
	})
	return pairs
}
