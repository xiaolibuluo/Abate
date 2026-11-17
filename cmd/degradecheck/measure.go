/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"io/fs"
	"os"

	"abate/internal/corpus"
	"abate/internal/e2e"
	e2ecov "abate/internal/e2e/coverage"
	"abate/lang/air"
	fkcov "abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

const (
	manifestFile   = "ci/corpus_manifest.json"
	reportSchema   = 1
	notMeasured    = "NOT_MEASURED"
	measured       = "MEASURED"
	goLang         = "go"
	degradeFormula = "degraded mappings with observed kind>0 / (exact+degraded mappings) * 100"
)

// reportLanguages is the canonical seven-language report order.
var reportLanguages = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// langRow is one language row of the full-corpus degrade measurement.
type langRow struct {
	Language         string  `json:"language"`
	Status           string  `json:"status"`
	Files            int     `json:"files"`
	ParseOK          int     `json:"parseOK"`
	ParseFail        int     `json:"parseFail"`
	EmptyUnit        int     `json:"emptyUnit"`
	ObservedKinds    int     `json:"observedKinds"`
	Mappings         int     `json:"mappings"`
	DegradedObserved int     `json:"degradedObserved"`
	DegradePct       float64 `json:"degradePct"`
}

// measureReport is the deterministic full-corpus degradation report.
type measureReport struct {
	SchemaVersion  int       `json:"schemaVersion"`
	ManifestSHA256 string    `json:"manifestSha256"`
	Formula        string    `json:"formula"`
	Results        []langRow `json:"results"`
}

// tallyState holds the per-language accumulators of one measurement run.
type tallyState struct {
	files     map[string]int
	parseOK   map[string]int
	parseFail map[string]int
	emptyUnit map[string]int
	observed  map[string]map[air.NodeKind]int
}

// measure parses every registered corpus entry with its real frontend, merges
// the observed AIR node kinds per language, and applies the frozen Degradation
// formula. Go has no embedded grammar mappings and stays NOT_MEASURED.
func measure() (measureReport, error) {
	m, err := corpus.Load(manifestFile)
	if err != nil {
		return measureReport{}, err
	}
	ft, err := e2e.NewFrontendTable()
	if err != nil {
		return measureReport{}, err
	}
	fsys := os.DirFS(corpus.DefaultCorpusRoot)
	state := newTally()
	for _, entry := range m.Entries {
		observeEntry(fsys, ft, state, entry)
	}
	return buildResult(m.Digest(), state), nil
}

// observeEntry parses one corpus file and folds its observed kinds into state.
// A read or parse failure is recorded, never skipped, so the result stays honest.
func observeEntry(fsys fs.FS, ft e2e.FrontendTable, state *tallyState, entry corpus.Entry) {
	state.files[entry.Language]++
	fe, dialect, ferr := ft.For(entry.Language)
	if ferr != nil {
		state.parseFail[entry.Language]++
		return
	}
	data, rerr := fs.ReadFile(fsys, entry.RelPath)
	if rerr != nil {
		state.parseFail[entry.Language]++
		return
	}
	unit, _, perr := fe.Parse(air.SourceUnit{
		Language: entry.Language,
		Dialect:  dialect,
		Files:    []air.SourceFile{air.NewSourceFile(entry.RelPath, data)},
	})
	if perr != nil || unit == nil {
		state.parseFail[entry.Language]++
		return
	}
	state.parseOK[entry.Language]++
	obs := fkcov.Observe(unit)
	if len(obs) == 0 {
		state.emptyUnit[entry.Language]++
	}
	state.observed[entry.Language] = fkcov.Merge(state.observed[entry.Language], obs)
}

// buildResult turns the merged tallies into the canonical ordered report.
func buildResult(manifestSHA string, state *tallyState) measureReport {
	rows := make([]langRow, 0, len(reportLanguages))
	for _, lang := range reportLanguages {
		rows = append(rows, buildRow(lang, state))
	}
	return measureReport{
		SchemaVersion:  reportSchema,
		ManifestSHA256: manifestSHA,
		Formula:        degradeFormula,
		Results:        rows,
	}
}

// buildRow computes one language row. Go keeps its honest NOT_MEASURED status.
func buildRow(lang string, state *tallyState) langRow {
	r := langRow{
		Language:      lang,
		Files:         state.files[lang],
		ParseOK:       state.parseOK[lang],
		ParseFail:     state.parseFail[lang],
		EmptyUnit:     state.emptyUnit[lang],
		ObservedKinds: len(state.observed[lang]),
	}
	if lang == goLang {
		r.Status = notMeasured
		return r
	}
	facts, err := grammar.Load(lang)
	if err != nil {
		r.Status = notMeasured
		return r
	}
	mappings := grammar.Map(facts, grammar.DefaultRules(lang))
	r.Mappings = len(mappings)
	for _, mapping := range mappings {
		if mapping.Status == grammar.StatusDegraded && state.observed[lang][mapping.Kind] > 0 {
			r.DegradedObserved++
		}
	}
	r.DegradePct = e2ecov.Degradation(mappings, state.observed[lang])
	r.Status = measured
	return r
}

// newTally allocates all accumulators with the canonical language keys.
func newTally() *tallyState {
	s := &tallyState{
		files:     make(map[string]int),
		parseOK:   make(map[string]int),
		parseFail: make(map[string]int),
		emptyUnit: make(map[string]int),
		observed:  make(map[string]map[air.NodeKind]int),
	}
	for _, lang := range reportLanguages {
		s.observed[lang] = make(map[air.NodeKind]int)
	}
	return s
}
