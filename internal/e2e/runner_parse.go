/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"cmp"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"abate/internal/corpus"
	"abate/internal/e2e/incremental"
	"abate/internal/e2e/metrics"
	"abate/internal/e2e/neutrality"
	"abate/lang/air"
)

// errCorpusRead reports a corpus file that could not be read from the root.
var errCorpusRead = errors.New("e2e: corpus read failed")

// parseFiles reads and parses every sampled entry in order. A read or parse
// contract failure is returned instead of skipping the entry, so the
// ProcessedFiles == RegisteredFiles invariant is never weakened.
func (r *Runner) parseFiles(entries []corpus.Entry) ([]fileState, error) {
	fsys := os.DirFS(r.plan.CorpusRoot)
	states := make([]fileState, 0, len(entries))
	for _, entry := range entries {
		state, err := r.parseOne(fsys, entry)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	if err := r.probeIncremental(states); err != nil {
		return nil, err
	}
	return states, nil
}

// parseOne reads one corpus file, parses it with its real frontend and records
// the AIR root identifier, the digest match, the SPDX class and the interval
// observation. No recover is used and no file is excluded by provenance.
func (r *Runner) parseOne(fsys fs.FS, entry corpus.Entry) (fileState, error) {
	data, err := fs.ReadFile(fsys, entry.RelPath)
	if err != nil {
		return fileState{}, fmt.Errorf("%w: %s: %w", errCorpusRead,
			entry.RelPath, err)
	}
	frontend, dialect, err := r.frontends.For(entry.Language)
	if err != nil {
		return fileState{}, err
	}
	unit := air.SourceUnit{
		Language: entry.Language,
		Dialect:  dialect,
		Files:    []air.SourceFile{air.NewSourceFile(entry.RelPath, data)},
	}
	r.trace.Mark(entry, "parse.before", 0)
	var parsed *air.Unit
	var diags []air.Diagnostic
	outcome := r.probe.Measure(func() (*air.Unit, []air.Diagnostic, error) {
		var parseErr error
		parsed, diags, parseErr = frontend.Parse(unit)
		return parsed, diags, parseErr
	})
	r.trace.Mark(entry, "parse.after", outcome.AllocDeltaBytes)
	sum := sha256.Sum256(data)
	state := fileState{
		entry: entry, content: data, unit: parsed,
		rootID: rootIDOf(parsed), diagnostics: len(diags), outcome: outcome,
		shaOK:     hex.EncodeToString(sum[:]) == entry.SHA256,
		spdxClass: corpus.ClassifySPDX(entry.SPDX),
	}
	if parsed != nil {
		for _, root := range parsed.Roots {
			if air.CheckIntervals(root, len(data)) != nil {
				state.intervalError = true
			}
		}
	}
	return state, nil
}

// probeIncremental runs the P1 ParseIncremental consistency check on the first
// real unit of each language and records its edited-range count.
func (r *Runner) probeIncremental(states []fileState) error {
	seen := make(map[string]bool)
	for i := range states {
		st := &states[i]
		if st.unit == nil || seen[st.entry.Language] {
			continue
		}
		seen[st.entry.Language] = true
		frontend, dialect, err := r.frontends.For(st.entry.Language)
		if err != nil {
			return err
		}
		unit := air.SourceUnit{
			Language: st.entry.Language,
			Dialect:  dialect,
			Files: []air.SourceFile{
				air.NewSourceFile(st.entry.RelPath, st.content),
			},
		}
		edits := []air.Edit{{
			Path: st.entry.RelPath, Offset: 0, EndOffset: 1,
			Text: []byte{' '},
		}}
		result, err := incremental.NewProbe(frontend).Check(unit, st.unit, edits)
		if err != nil {
			return err
		}
		st.changedRanges = result.ChangedRanges
		st.consistent = result.Consistent
	}
	return nil
}

// rootIDOf returns the AIR identifier of a unit's first root, or the empty
// string when no root was produced.
func rootIDOf(u *air.Unit) string {
	if u == nil || len(u.Roots) == 0 {
		return ""
	}
	return air.IDOf(u.Roots[0]).String()
}

// neutralityFindings runs the zero-exemption go/ast check over every .go path
// registered in ci/e2e_sources.txt. The list is an anchor, not a whitelist.
func (r *Runner) neutralityFindings() ([]neutrality.Finding, error) {
	paths, err := readSourceList(r.plan.SourcesPath)
	if err != nil {
		return nil, err
	}
	res, err := neutrality.NewAnalyzer(r.plan.RepoRoot).Check(paths)
	if err != nil {
		return nil, err
	}
	return res.Findings, nil
}

// readSourceList reads the newline-separated source paths relative to the
// repository root through fs.ReadFile, avoiding variable-path os.Open.
func readSourceList(path string) ([]string, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return nil, fmt.Errorf("%w: read source list %s: %w",
			errCorpusRead, path, err)
	}
	out := make([]string, 0)
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("%w: empty source list", errCorpusRead)
	}
	return out, nil
}

// sortedEntries orders the sample by (Language, RelPath) ascending.
func sortedEntries(in []corpus.Entry) []corpus.Entry {
	out := slices.Clone(in)
	slices.SortFunc(out, func(a, b corpus.Entry) int {
		if c := strings.Compare(a.Language, b.Language); c != 0 {
			return c
		}
		return strings.Compare(a.RelPath, b.RelPath)
	})
	return out
}

// sortRecords orders the drilldown by (Language, RelPath, OperatorID, Offset).
func sortRecords(in []metrics.DrilldownRecord) []metrics.DrilldownRecord {
	out := slices.Clone(in)
	slices.SortFunc(out, func(a, b metrics.DrilldownRecord) int {
		if c := strings.Compare(a.Language, b.Language); c != 0 {
			return c
		}
		if c := strings.Compare(a.RelPath, b.RelPath); c != 0 {
			return c
		}
		if c := strings.Compare(a.OperatorID, b.OperatorID); c != 0 {
			return c
		}
		return cmp.Compare(a.Offset, b.Offset)
	})
	return out
}

// localizersOf keeps a bounded, sorted set of real drilldown localizers.
func localizersOf(recs []metrics.DrilldownRecord) []string {
	seen := make(map[string]bool)
	out := make([]string, 0)
	add := func(value string) {
		if value != "" && !seen[value] && len(out) < 16 {
			seen[value] = true
			out = append(out, value)
		}
	}
	if len(recs) > 0 {
		add(recs[0].Localizer())
	}
	for _, rec := range recs {
		if rec.OutcomeKind != metrics.OutcomeOK {
			add(rec.Localizer())
		}
	}
	slices.Sort(out)
	return out
}

// intervalCount totals real CheckIntervals observations across sampled files.
func intervalCount(states []fileState) int {
	count := 0
	for _, st := range states {
		if st.intervalError {
			count++
		}
	}
	return count
}

// changedRanges totals the P1 incremental probe's edited ranges.
func changedRanges(states []fileState) int {
	count := 0
	for _, st := range states {
		count += st.changedRanges
	}
	return count
}

// consistent reports whether every P1 incremental probe stayed consistent.
func consistent(states []fileState) bool {
	for _, st := range states {
		if st.changedRanges > 0 && !st.consistent {
			return false
		}
	}
	return true
}
