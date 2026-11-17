/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strconv"

	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/rule"
)

// otelServiceName is the fixed resource identifier of the air-gapped producer.
const otelServiceName = "abate-rpt"

// otelTraceID derives a 32-character trace identifier from the report digest.
func otelTraceID(rep mutation.Report) string {
	return hex.EncodeToString(rep.Fingerprint[:16])
}

// otelSpanID derives a 16-character span identifier from the report digest and
// a stable name, with no random source involved.
func otelSpanID(rep mutation.Report, name string) string {
	h := sha256.New()
	_, _ = h.Write(rep.Fingerprint[:])
	_, _ = h.Write([]byte(name))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum[:8])
}

// otelVerdictSeverity maps a mutant verdict to a log severity level.
func otelVerdictSeverity(verdict string) string {
	switch verify.MutationVerdict(verdict) {
	case verify.MutantSurvived:
		return "ERROR"
	case verify.MutantEquivalent, verify.MutantTimeout:
		return "WARN"
	case verify.MutantKilled:
		return "INFO"
	default:
		return "WARN"
	}
}

// otelFindingSeverity maps a rule severity to a log severity level.
func otelFindingSeverity(severity rule.Severity) string {
	switch severity {
	case rule.SeverityError:
		return "ERROR"
	case rule.SeverityWarning:
		return "WARN"
	case rule.SeverityInfo:
		return "INFO"
	default:
		return "WARN"
	}
}

// otelSpanLess is the canonical order over spans.
func otelSpanLess(a, b otelSpan) bool {
	return a.SpanID < b.SpanID
}

// otelLogLess is the canonical order over log records.
func otelLogLess(a, b otelLogRecord) bool {
	if a.Body != b.Body {
		return a.Body < b.Body
	}
	if a.VirtualTime != b.VirtualTime {
		return a.VirtualTime < b.VirtualTime
	}
	if a.SeverityText != b.SeverityText {
		return a.SeverityText < b.SeverityText
	}
	return otelLogKind(a) < otelLogKind(b)
}

// otelLogKind returns the leading attribute value, or an empty sentinel.
func otelLogKind(r otelLogRecord) string {
	if len(r.Attributes) == 0 {
		return ""
	}
	return r.Attributes[0].Value
}

// buildOTelResourceSpans maps the HPOV phase tree into a deterministic span tree.
func buildOTelResourceSpans(rep mutation.Report) otelResourceSpans {
	traceID := otelTraceID(rep)
	rootID := otelSpanID(rep, "abate.run")
	verifyID := otelSpanID(rep, "verify")
	spans := []otelSpan{{
		TraceID:     traceID,
		SpanID:      rootID,
		Name:        "abate.run",
		VirtualTime: strconv.FormatUint(rep.Seed, 10),
		Attributes: []otelAttribute{
			{Key: "service.name", Value: otelServiceName},
			{Key: "report.fingerprint", Value: hex.EncodeToString(rep.Fingerprint[:])},
			{Key: "report.seed", Value: strconv.FormatUint(rep.Seed, 10)},
		},
	}}
	phases := [4]string{"discover", "schedule", "verify", "score"}
	for i, phase := range phases {
		spans = append(spans, otelSpan{
			TraceID:      traceID,
			SpanID:       otelSpanID(rep, phase),
			ParentSpanID: rootID,
			Name:         phase,
			VirtualTime:  strconv.Itoa(i + 1),
			Attributes: []otelAttribute{
				{Key: "phase.index", Value: strconv.Itoa(i + 1)},
			},
		})
	}
	for _, row := range mutantRows(rep) {
		name := "mutant." + strconv.Itoa(row.Index)
		spans = append(spans, otelSpan{
			TraceID:      traceID,
			SpanID:       otelSpanID(rep, name+"."+row.Fingerprint),
			ParentSpanID: verifyID,
			Name:         name,
			VirtualTime:  strconv.Itoa(row.Index),
			Attributes: []otelAttribute{
				{Key: "mutant.verdict", Value: row.Verdict},
				{Key: "mutant.fingerprint", Value: row.Fingerprint},
				{Key: "mutant.node", Value: row.Node},
				{Key: "mutant.original", Value: row.Original},
				{Key: "mutant.mutated", Value: row.Mutated},
				{Key: "mutant.killingTest", Value: row.KillingTest},
			},
		})
	}
	sort.Slice(spans, func(i, j int) bool {
		return otelSpanLess(spans[i], spans[j])
	})
	return otelResourceSpans{
		Resource: otelResource{Attributes: []otelAttribute{
			{Key: "service.name", Value: otelServiceName},
			{Key: "deployment.mode", Value: "airgap"},
			{Key: "report.seed", Value: strconv.FormatUint(rep.Seed, 10)},
		}},
		Spans: spans,
	}
}

// buildOTelMetrics maps score counters and finding counts into offline metrics.
func buildOTelMetrics(findings []rule.Finding, rep mutation.Report) []otelMetric {
	counters := [6]struct {
		name  string
		value uint64
	}{
		{"abate.mutants.total", rep.Score.Total},
		{"abate.mutants.killed", rep.Score.Killed},
		{"abate.mutants.survived", rep.Score.Survived},
		{"abate.mutants.equivalent", rep.Score.Equivalent},
		{"abate.mutants.timeout", rep.Score.Timeout},
		{"abate.findings.total", uint64(len(findings))}, //nolint:gosec // slice length is non-negative.
	}
	metrics := make([]otelMetric, 0, len(counters))
	for i, c := range counters {
		metrics = append(metrics, otelMetric{
			Name:        c.name,
			Value:       c.value,
			VirtualTime: strconv.Itoa(i),
			Attributes: []otelAttribute{
				{Key: "metric.index", Value: strconv.Itoa(i)},
				{Key: "report.seed", Value: strconv.FormatUint(rep.Seed, 10)},
			},
		})
	}
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Name < metrics[j].Name
	})
	return metrics
}

// buildOTelLogs maps verdicts and findings into deterministic log records.
func buildOTelLogs(findings []rule.Finding, rep mutation.Report) []otelLogRecord {
	rows := mutantRows(rep)
	sorted := sortedFindings(findings)
	logs := make([]otelLogRecord, 0, len(rows)+len(sorted))
	for _, row := range rows {
		logs = append(logs, otelLogRecord{
			SeverityText: otelVerdictSeverity(row.Verdict),
			Body:         "mutant " + row.Fingerprint + " " + row.Verdict,
			VirtualTime:  strconv.Itoa(row.Index),
			Attributes: []otelAttribute{
				{Key: "log.kind", Value: "verdict"},
				{Key: "mutant.verdict", Value: row.Verdict},
				{Key: "mutant.killingTest", Value: row.KillingTest},
			},
		})
	}
	for i, f := range sorted {
		logs = append(logs, otelLogRecord{
			SeverityText: otelFindingSeverity(f.Severity),
			Body:         f.Message,
			VirtualTime:  strconv.Itoa(i + 1),
			Attributes: []otelAttribute{
				{Key: "log.kind", Value: "finding"},
				{Key: "rule.id", Value: string(f.RuleID)},
				{Key: "rule.path", Value: f.Path},
				{Key: "rule.reference", Value: f.Reference},
			},
		})
	}
	sort.Slice(logs, func(i, j int) bool {
		return otelLogLess(logs[i], logs[j])
	})
	return logs
}
