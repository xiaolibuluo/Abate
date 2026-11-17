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

	"abate/internal/corpus"
	e2ecoverage "abate/internal/e2e/coverage"
	"abate/internal/e2e/metrics"
	"abate/internal/e2e/mutate"
	"abate/internal/e2e/report"
	"abate/lang/air"
)

// errEmptySample reports a manifest whose byte window selected no files.
var errEmptySample = errors.New("e2e: empty stratified sample")

// fileState is the complete in-memory result of one sampled corpus file.
type fileState struct {
	entry         corpus.Entry
	content       []byte
	unit          *air.Unit
	rootID        string
	diagnostics   int
	outcome       Outcome
	shaOK         bool
	spdxClass     string
	intervalError bool
	changedRanges int
	consistent    bool
}

// Runner executes the deterministic end-to-end measurement in one goroutine.
type Runner struct {
	plan          Plan
	manifest      corpus.Manifest
	frontends     FrontendTable
	registry      mutate.Registry
	equivalents   mutate.EquivalentSet
	coverageIndex e2ecoverage.Index
	probe         ResourceProbe
	trace         *Trace
}

// NewRunner wires the seven frontends, the eight operators, the equivalent
// definitions, the construct inventory and the resource probe. Every failure
// is returned, so a partially built runner is never used.
func NewRunner(p Plan, m corpus.Manifest, trace *Trace) (*Runner, error) {
	frontends, err := NewFrontendTable()
	if err != nil {
		return nil, err
	}
	registry, err := mutate.DefaultRegistry()
	if err != nil {
		return nil, err
	}
	equivalents, err := mutate.LoadEquivalentSet(p.EquivalentPath)
	if err != nil {
		return nil, err
	}
	index, err := e2ecoverage.NewCoverageIndex(p.ConstructsPath)
	if err != nil {
		return nil, err
	}
	return &Runner{
		plan: p, manifest: m, frontends: frontends, registry: registry,
		equivalents: equivalents, coverageIndex: index,
		probe: NewResourceProbe(p.Budget()), trace: trace,
	}, nil
}

// Run parses the stratified sample, mutates it, damages it, recomputes the
// construct coverage and returns the report plus the full drilldown record
// set. Aggregates come only from metrics.Aggregate over those records.
func (r *Runner) Run() (report.Report, []metrics.DrilldownRecord, error) {
	entries := sortedEntries(corpus.StratifiedSample(
		r.manifest, r.plan.SamplePlan()))
	if len(entries) == 0 {
		return report.Report{}, nil, errEmptySample
	}
	states, err := r.parseFiles(entries)
	if err != nil {
		return report.Report{}, nil, err
	}
	mutationRecs, err := r.mutateFiles(states)
	if err != nil {
		return report.Report{}, nil, err
	}
	brokenRecs, precision, err := r.chaosPass(states)
	if err != nil {
		return report.Report{}, nil, err
	}
	records := sortRecords(append(mutationRecs, brokenRecs...))
	precision, coverageStatus, err := r.coveragePass(states, precision)
	if err != nil {
		return report.Report{}, nil, err
	}
	facts := manifestFacts(entries, states)
	findings, err := r.neutralityFindings()
	if err != nil {
		return report.Report{}, nil, err
	}
	summary := metrics.Aggregate(records)
	baseline := report.DeriveBaseline(records, precision, facts)
	counters := (counterFacts{
		entries: entries, states: states, records: records,
		summary: summary, findings: findings, facts: facts,
	}).Counters()
	thresholdsSHA, err := report.DigestFile(r.plan.ThresholdsPath)
	if err != nil {
		return report.Report{}, nil, err
	}
	rep := buildReport(reportInput{
		plan: r.plan, manifest: r.manifest, summary: summary,
		records: records, precision: precision, baseline: baseline,
		counters: counters, coverageStatus: coverageStatus, facts: facts,
		states: states, thresholdsSHA256: thresholdsSHA,
	})
	if err = r.trace.Flush(); err != nil {
		return report.Report{}, nil, err
	}
	if err = r.trace.Err(); err != nil {
		return report.Report{}, nil, err
	}
	return rep, records, nil
}

// reportInput bundles the buildReport arguments into one option struct.
type reportInput struct {
	plan             Plan
	manifest         corpus.Manifest
	summary          metrics.Summary
	records          []metrics.DrilldownRecord
	precision        []metrics.PrecisionRow
	baseline         []report.BaselineValue
	counters         []report.Counter
	coverageStatus   []string
	facts            report.ManifestFacts
	states           []fileState
	thresholdsSHA256 string
}

// buildReport assembles the deterministic report struct. No wall clock, no
// absolute path, no map and no runtime version string appears in any field.
func buildReport(in reportInput) report.Report {
	return report.Report{
		SchemaVersion:            ReportSchemaVersion,
		ToolVersion:              ToolVersion,
		PlanSHA256:               in.plan.PlanSHA256,
		ManifestSHA256:           in.manifest.Digest(),
		ThresholdsSHA256:         in.thresholdsSHA256,
		Seed:                     in.plan.Seed,
		Languages:                slices.Clone(in.plan.Languages),
		MaxFileBytes:             in.plan.MaxFileBytes,
		ParseTimeoutMillis:       in.plan.ParseTimeoutMillis,
		MemLimitBytes:            in.plan.MemLimitBytes,
		MaxRuns:                  in.plan.MaxRuns,
		Merge:                    MergeLabel,
		Columns:                  in.summary.Columns,
		Precision:                in.precision,
		Mutation:                 in.summary.Mutation,
		Robustness:               in.summary.Robustness,
		Equivalence:              in.summary.Equivalence,
		Localizers:               localizersOf(in.records),
		CompetitorAbsolute:       CompetitorNotMeasured,
		CoverageStatus:           slices.Clone(in.coverageStatus),
		Facts:                    in.facts,
		Counters:                 in.counters,
		BaselineValues:           in.baseline,
		ProcessedFiles:           len(in.states),
		RegisteredFiles:          in.facts.RegisteredFiles,
		IntervalFindings:         intervalCount(in.states),
		IncrementalChangedRanges: changedRanges(in.states),
		IncrementalConsistent:    consistent(in.states),
	}
}
