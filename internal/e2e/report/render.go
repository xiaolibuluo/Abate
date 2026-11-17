/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"fmt"
	"strings"

	"abate/internal/e2e/metrics"
)

// Markdown renders the deterministic human-readable report. The four columns
// are printed separately, timeouts and OOMs are independent rows, and the
// competitor absolute score is always labelled NOT_MEASURED.
func Markdown(r Report) string {
	var b strings.Builder
	writeHeader(&b, r)
	writeColumns(&b, r)
	writeRobustness(&b, r)
	writePrecision(&b, r.Precision)
	writeMutation(&b, r.Mutation)
	writeCounters(&b, r.Counters)
	writeBaseline(&b, r.BaselineValues)
	writeFooter(&b, r)
	return b.String()
}

// writeHeader emits the title and the four SHA-256 integrity chains.
func writeHeader(b *strings.Builder, r Report) {
	b.WriteString("# Abate E2E deterministic report\n\n")
	_, _ = fmt.Fprintf(b, "- toolVersion: `%s`\n", r.ToolVersion)
	_, _ = fmt.Fprintf(b, "- seed: `%d`\n", r.Seed)
	_, _ = fmt.Fprintf(b, "- languages: `%s`\n", strings.Join(r.Languages, ","))
	_, _ = fmt.Fprintf(b, "- planSha256: `%s`\n", r.PlanSHA256)
	_, _ = fmt.Fprintf(b, "- manifestSha256: `%s`\n", r.ManifestSHA256)
	_, _ = fmt.Fprintf(b, "- thresholdsSha256: `%s`\n", r.ThresholdsSHA256)
	_, _ = fmt.Fprintf(b, "- drilldownSha256: `%s`\n", r.DrilldownSHA256)
	_, _ = fmt.Fprintf(b, "- budgets: maxFileBytes=%d timeoutMillis=%d memLimitBytes=%d maxRuns=%d merge=%s\n\n",
		r.MaxFileBytes, r.ParseTimeoutMillis, r.MemLimitBytes, r.MaxRuns, r.Merge)
}

// writeColumns prints the four distinct mutation columns.
func writeColumns(b *strings.Builder, r Report) {
	c := r.Columns
	b.WriteString("## Four separate columns\n\n")
	b.WriteString("| column | gross | effective | captured | capturePct |\n")
	b.WriteString("|---|---:|---:|---:|---:|\n")
	rows := []struct {
		name  string
		stats metrics.ColumnStats
	}{
		{"text", c.Text}, {"structure", c.Structure},
		{"broken", c.Broken}, {"equivalent", c.Equivalent},
	}
	for _, row := range rows {
		_, _ = fmt.Fprintf(b, "| %s | %d | %d | %d | %.2f |\n",
			row.name, row.stats.Gross, row.stats.Effective,
			row.stats.Captured, row.stats.CapturePct)
	}
	_, _ = fmt.Fprintf(b, "\n- equivalence: gross=%d effective=%d equivalentCount=%d silentDrops=%d\n\n",
		r.Equivalence.Gross, r.Equivalence.Effective,
		r.Equivalence.EquivalentCount, r.Equivalence.SilentDrops)
}

// writeRobustness prints timeout and OOM as two independent columns.
func writeRobustness(b *strings.Builder, r Report) {
	b.WriteString("## Robustness\n\n")
	b.WriteString("| metric | value |\n|---|---:|\n")
	_, _ = fmt.Fprintf(b, "| timeoutCount | %d |\n", r.Robustness.TimeoutCount)
	_, _ = fmt.Fprintf(b, "| oomCount | %d |\n", r.Robustness.OOMCount)
	_, _ = fmt.Fprintf(b, "| determinismFailureCount | %d |\n",
		r.Robustness.DeterminismFailureCount)
	_, _ = fmt.Fprintf(b, "| processedFiles | %d |\n", r.ProcessedFiles)
	_, _ = fmt.Fprintf(b, "| registeredFiles | %d |\n", r.RegisteredFiles)
	_, _ = fmt.Fprintf(b, "| panicLocalizer | `%s` |\n\n",
		r.Robustness.PanicLocalizer)
}

// writePrecision prints one row per language.
func writePrecision(b *strings.Builder, rows []metrics.PrecisionRow) {
	b.WriteString("## Precision by language\n\n")
	b.WriteString("| language | clean | diagnostics | parseErrors | degradeNodePct | outsideRetentionPct | constructCoveragePct |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|\n")
	for _, row := range rows {
		_, _ = fmt.Fprintf(b, "| %s | %d | %d | %d | %.2f | %.2f | %.2f |\n",
			row.Language, row.ZeroDiagnosticZeroError, row.Diagnostics,
			row.ParseErrors, row.DegradeNodePct,
			row.OutsideDamageRetentionPct, row.ConstructCoveragePct)
	}
	b.WriteString("\n")
}

// writeMutation prints every deterministic mutation row.
func writeMutation(b *strings.Builder, rows []metrics.MutationRow) {
	b.WriteString("## Mutation rows\n\n")
	b.WriteString("| language | operator | class | gross | captured | missed | timeout | unviable | capturePct |\n")
	b.WriteString("|---|---|---|---:|---:|---:|---:|---:|---:|\n")
	for _, row := range rows {
		_, _ = fmt.Fprintf(b, "| %s | %s | %s | %d | %d | %d | %d | %d | %.2f |\n",
			row.Language, row.OperatorID, row.Class, row.Gross,
			row.Captured, row.Missed, row.Timeout, row.Unviable,
			row.CapturePct)
	}
	b.WriteString("\n")
}

// writeCounters prints the zero-tolerance observations.
func writeCounters(b *strings.Builder, counters []Counter) {
	b.WriteString("## Zero tolerance\n\n")
	b.WriteString("| key | value |\n|---|---:|\n")
	for _, counter := range counters {
		_, _ = fmt.Fprintf(b, "| %s | %d |\n", counter.Key, counter.Value)
	}
	b.WriteString("\n")
}

// writeBaseline prints the recomputable baseline current values.
func writeBaseline(b *strings.Builder, values []BaselineValue) {
	b.WriteString("## Baseline current values\n\n")
	b.WriteString("| key | value |\n|---|---:|\n")
	for _, value := range values {
		_, _ = fmt.Fprintf(b, "| %s | %.6f |\n", value.Key, value.Value)
	}
	b.WriteString("\n")
}

// writeFooter prints the P2 honesty label and the measurement statuses.
func writeFooter(b *strings.Builder, r Report) {
	_, _ = fmt.Fprintf(b, "## Competitor and coverage status\n\n")
	_, _ = fmt.Fprintf(b, "- competitorAbsolute: `%s`\n", r.CompetitorAbsolute)
	_, _ = fmt.Fprintf(b, "- coverageStatus: `%s`\n",
		strings.Join(r.CoverageStatus, ","))
	_, _ = fmt.Fprintf(b, "- incrementalChangedRanges: `%d`\n",
		r.IncrementalChangedRanges)
	_, _ = fmt.Fprintf(b, "- incrementalConsistent: `%t`\n",
		r.IncrementalConsistent)
	_, _ = fmt.Fprintf(b, "- intervalFindings: `%d`\n", r.IntervalFindings)
}
