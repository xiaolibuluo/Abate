/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"errors"
	"fmt"
	"strings"

	"abate/internal/determinism"
	"abate/internal/e2e/metrics"
	"abate/internal/e2e/mutate"
	"abate/lang/air"
)

// errMissingOperator reports an operator declared by the plan but absent from
// the registry, which is a contract failure rather than a reason to skip.
var errMissingOperator = errors.New("e2e: unregistered operator")

// mutateFiles applies every declared operator to every sampled file in the
// fixed order and returns one drilldown record per real in-memory rewrite.
func (r *Runner) mutateFiles(states []fileState) ([]metrics.DrilldownRecord, error) {
	out := make([]metrics.DrilldownRecord, 0,
		len(states)*len(r.plan.Operators))
	for i := range states {
		recs, err := r.mutateOne(&states[i])
		if err != nil {
			return nil, err
		}
		out = append(out, recs...)
	}
	return out, nil
}

// mutateOne scans one file for every operator and reparses each selected edit.
func (r *Runner) mutateOne(st *fileState) ([]metrics.DrilldownRecord, error) {
	out := make([]metrics.DrilldownRecord, 0, len(r.plan.Operators))
	for _, id := range r.plan.Operators {
		spec, found := r.registry.Lookup(id)
		if !found {
			return nil, fmt.Errorf("%w: %s", errMissingOperator, id)
		}
		edits, err := spec.Apply(mutate.OpContext{
			Content: st.content, Language: st.entry.Language,
			ManifestSHA256: r.manifest.Digest(), Seed: r.plan.Seed,
		})
		if err != nil {
			return nil, err
		}
		for _, edit := range takeEdits(edits, maxEditsPerOperator) {
			rec, err := r.applyEdit(st, spec, edit)
			if err != nil {
				return nil, err
			}
			out = append(out, rec)
		}
	}
	return out, nil
}

// takeEdits selects a deterministic, equally spaced subset of the real edits.
func takeEdits(edits []mutate.Edit, want int) []mutate.Edit {
	indices := determinism.Stride(len(edits), want)
	out := make([]mutate.Edit, 0, len(indices))
	for _, index := range indices {
		out = append(out, edits[index])
	}
	return out
}

// applyEdit performs one in-memory replacement, reparses the mutated buffer,
// and returns the AIR-ID capture drilldown record for that mutation.
func (r *Runner) applyEdit(st *fileState, spec mutate.OperatorSpec,
	edit mutate.Edit) (metrics.DrilldownRecord, error) {
	mutated, err := mutate.Apply([]mutate.Edit{edit}, st.content)
	if err != nil {
		return metrics.DrilldownRecord{}, err
	}
	frontend, dialect, err := r.frontends.For(st.entry.Language)
	if err != nil {
		return metrics.DrilldownRecord{}, err
	}
	unit := air.SourceUnit{
		Language: st.entry.Language,
		Dialect:  dialect,
		Files: []air.SourceFile{
			air.NewSourceFile(st.entry.RelPath, mutated),
		},
	}
	r.trace.Mark(st.entry, "mutate.before."+edit.OperatorID, 0)
	var parsed *air.Unit
	var diags []air.Diagnostic
	outcome := r.probe.Measure(func() (*air.Unit, []air.Diagnostic, error) {
		var parseErr error
		parsed, diags, parseErr = frontend.Parse(unit)
		return parsed, diags, parseErr
	})
	r.trace.Mark(st.entry, "mutate.after."+edit.OperatorID,
		outcome.AllocDeltaBytes)
	afterID := rootIDOf(parsed)
	column := spec.Class.Column()
	return metrics.DrilldownRecord{
		Language: st.entry.Language, FileSHA256: st.entry.SHA256,
		RelPath: st.entry.RelPath, OperatorID: edit.OperatorID,
		Class: column, Offset: edit.Offset, EndOffset: edit.EndOffset,
		BeforeText: edit.Before, AfterText: edit.After,
		AIRIDBefore: st.rootID, AIRIDAfter: afterID,
		Captured:          st.rootID != afterID,
		Equivalent:        column == metrics.ColumnEquivalent,
		EquivalentProofID: proofFor(r.equivalents, edit.OperatorID),
		OutcomeKind:       outcome.Kind, AllocDeltaBytes: outcome.AllocDeltaBytes,
		SilentDrops: 0,
	}, nil
}

// proofFor resolves the explicit equivalent-mutation proof of an operator.
func proofFor(set mutate.EquivalentSet, operatorID string) string {
	for _, def := range set.All() {
		if strings.HasPrefix(def.OperatorID, operatorID+".") {
			return def.ProofID
		}
	}
	return ""
}
