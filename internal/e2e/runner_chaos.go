/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"abate/internal/e2e/chaos"
	"abate/internal/e2e/metrics"
	"abate/lang/air"
)

// chaosPass builds one deterministic, sha256-deduplicated damaged corpus per
// language, reparses every damaged input and returns both the broken-column
// drilldown records and the per-language precision observation.
func (r *Runner) chaosPass(states []fileState) (
	[]metrics.DrilldownRecord, []metrics.PrecisionRow, error) {
	byLang := make(map[string][]int)
	for i := range states {
		byLang[states[i].entry.Language] = append(
			byLang[states[i].entry.Language], i)
	}
	records := make([]metrics.DrilldownRecord, 0)
	rows := make([]metrics.PrecisionRow, 0, len(languageOrder))
	for _, lang := range languageOrder {
		langRecs, row, err := r.chaosLanguage(lang, byLang[lang], states)
		if err != nil {
			return nil, nil, err
		}
		records = append(records, langRecs...)
		rows = append(rows, row)
	}
	return records, rows, nil
}

// chaosLanguage builds and measures the damaged corpus of one language.
func (r *Runner) chaosLanguage(lang string, idxs []int, states []fileState) (
	[]metrics.DrilldownRecord, metrics.PrecisionRow, error) {
	row := metrics.PrecisionRow{Language: lang}
	if len(idxs) == 0 {
		return []metrics.DrilldownRecord{}, row, nil
	}
	frontend, dialect, err := r.frontends.For(lang)
	if err != nil {
		return nil, row, err
	}
	base := &states[idxs[0]]
	damaged := r.damagedInputs(lang, idxs, states, base)
	limit := min(len(damaged), r.plan.BrokenPerLanguage)
	records := make([]metrics.DrilldownRecord, 0, limit)
	retentionTotal := 0.0
	diagnostics := base.diagnostics
	parseErrors := 0
	cleanBase := 0
	if base.diagnostics == 0 && base.outcome.Errors == 0 {
		cleanBase = 1
	}
	for i := 0; i < limit; i++ {
		unit := air.SourceUnit{
			Language: lang,
			Dialect:  dialect,
			Files: []air.SourceFile{
				air.NewSourceFile(base.entry.RelPath, damaged[i]),
			},
		}
		r.trace.Mark(base.entry, "chaos.before", 0)
		var parsed *air.Unit
		var diags []air.Diagnostic
		outcome := r.probe.Measure(
			func() (*air.Unit, []air.Diagnostic, error) {
				var parseErr error
				parsed, diags, parseErr = frontend.Parse(unit)
				return parsed, diags, parseErr
			})
		r.trace.Mark(base.entry, "chaos.after", outcome.AllocDeltaBytes)
		afterID := rootIDOf(parsed)
		records = append(records, metrics.DrilldownRecord{
			Language: lang, FileSHA256: base.entry.SHA256,
			RelPath: base.entry.RelPath, OperatorID: "chaos",
			Class: metrics.ColumnBroken, Offset: i, EndOffset: i + 1,
			BeforeText: "clean", AfterText: "damaged",
			AIRIDBefore: base.rootID, AIRIDAfter: afterID,
			Captured:        base.rootID != afterID,
			OutcomeKind:     outcome.Kind,
			AllocDeltaBytes: outcome.AllocDeltaBytes,
			SilentDrops:     0,
		})
		retentionTotal += retentionPct(base.unit, parsed)
		diagnostics += len(diags)
		if outcome.Kind == metrics.OutcomeContractError {
			parseErrors++
		}
	}
	row.ZeroDiagnosticZeroError = cleanBase
	row.Diagnostics = diagnostics
	row.ParseErrors = parseErrors
	if limit > 0 {
		row.OutsideDamageRetentionPct = retentionTotal / float64(limit)
	}
	return records, row, nil
}

// damagedInputs assembles the deterministic damaged corpus: half truncations
// of every sampled file, the two G1 reproductions for Go, and bounded byte
// splices selected by the SplitMix64 generator, all merged by sha256.
func (r *Runner) damagedInputs(lang string, idxs []int,
	states []fileState, base *fileState) [][]byte {
	inputs := make([][]byte, 0, len(idxs)+r.plan.BrokenPerLanguage)
	for _, i := range idxs {
		inputs = append(inputs, chaos.TruncateHalf(states[i].content))
	}
	if lang == "go" {
		inputs = append(inputs, chaos.G1Reproductions()...)
	}
	generator := chaos.NewGenerator(r.plan.Seed, r.manifest.Digest(),
		r.plan.BrokenPerLanguage+damagedBudgetExtra)
	inputs = append(inputs, generator.Generate(base.content, lang)...)
	return chaos.Merge(inputs)
}

// retentionPct measures the share of AIR node identifiers of the clean parse
// still present after a damage, using a multiset intersection.
func retentionPct(before, after *air.Unit) float64 {
	if before == nil || after == nil {
		return 0
	}
	nodes := air.Collect(before)
	if len(nodes) == 0 {
		return 0
	}
	counts := make(map[air.NodeID]int, len(nodes))
	for _, node := range nodes {
		counts[air.IDOf(node)]++
	}
	retained := 0
	for _, node := range air.Collect(after) {
		id := air.IDOf(node)
		if counts[id] > 0 {
			counts[id]--
			retained++
		}
	}
	return float64(retained) * fullCoveragePct / float64(len(nodes))
}
