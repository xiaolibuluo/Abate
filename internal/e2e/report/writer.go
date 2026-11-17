/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"abate/internal/e2e/metrics"
)

// Names and permissions of the deterministic output artifacts.
const (
	drilldownFileName = "e2e_drilldown.jsonl"
	markdownFileName  = "e2e_report.md"
	filePerm          = 0o600
	dirPerm           = 0o750
)

// drilldownLine is the canonical, deterministic projection of one drilldown
// record. AllocDeltaBytes is deliberately absent: TotalAlloc increments are
// runtime observations that live in e2e_trace.jsonl, and no runtime-derived
// byte may enter the artifacts compared by cmp. Aggregate never reads that
// field, so recomputation from these lines is exact.
type drilldownLine struct {
	Language          string `json:"language"`
	FileSHA256        string `json:"fileSha256"`
	RelPath           string `json:"relPath"`
	OperatorID        string `json:"operatorID"`
	Class             string `json:"class"`
	Offset            int    `json:"offset"`
	EndOffset         int    `json:"endOffset"`
	BeforeText        string `json:"beforeText"`
	AfterText         string `json:"afterText"`
	AIRIDBefore       string `json:"airIDBefore"`
	AIRIDAfter        string `json:"airIDAfter"`
	Captured          bool   `json:"captured"`
	Equivalent        bool   `json:"equivalent"`
	EquivalentProofID string `json:"equivalentProofID"`
	OutcomeKind       string `json:"outcomeKind"`
	SilentDrops       int    `json:"silentDrops"`
}

// toDrilldownLine projects one in-memory record onto its canonical line.
func toDrilldownLine(rec metrics.DrilldownRecord) drilldownLine {
	return drilldownLine{
		Language: rec.Language, FileSHA256: rec.FileSHA256,
		RelPath: rec.RelPath, OperatorID: rec.OperatorID,
		Class: rec.Class, Offset: rec.Offset, EndOffset: rec.EndOffset,
		BeforeText: rec.BeforeText, AfterText: rec.AfterText,
		AIRIDBefore: rec.AIRIDBefore, AIRIDAfter: rec.AIRIDAfter,
		Captured: rec.Captured, Equivalent: rec.Equivalent,
		EquivalentProofID: rec.EquivalentProofID,
		OutcomeKind:       rec.OutcomeKind, SilentDrops: rec.SilentDrops,
	}
}

// record reconstructs the aggregation input from one canonical line.
func (l drilldownLine) record() metrics.DrilldownRecord {
	return metrics.DrilldownRecord{
		Language: l.Language, FileSHA256: l.FileSHA256,
		RelPath: l.RelPath, OperatorID: l.OperatorID,
		Class: l.Class, Offset: l.Offset, EndOffset: l.EndOffset,
		BeforeText: l.BeforeText, AfterText: l.AfterText,
		AIRIDBefore: l.AIRIDBefore, AIRIDAfter: l.AIRIDAfter,
		Captured: l.Captured, Equivalent: l.Equivalent,
		EquivalentProofID: l.EquivalentProofID,
		OutcomeKind:       l.OutcomeKind, SilentDrops: l.SilentDrops,
	}
}

// Write emits e2e_report.json, e2e_drilldown.jsonl and e2e_report.md into the
// directory of out. The drilldown file is encoded first, so DrilldownSHA256 is
// the SHA-256 of exactly the bytes written beside the report.
func Write(out string, r Report, recs []metrics.DrilldownRecord) error {
	dir := filepath.Dir(out)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return err
	}
	drilldown, err := encodeDrilldown(recs)
	if err != nil {
		return err
	}
	r.DrilldownSHA256 = DigestBytes(drilldown)
	encoded, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	encoded = append(encoded, '\n')
	if err = os.WriteFile(out, encoded, filePerm); err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(dir, drilldownFileName),
		drilldown, filePerm); err != nil {
		return err
	}
	markdown := []byte(Markdown(r) + "\n")
	return os.WriteFile(filepath.Join(dir, markdownFileName), markdown, filePerm)
}

// verdictFileName is the deterministic verdict artifact written beside the
// report. It is the single source of truth for gate_e2e.sh S6; its passed
// field is the only signal the gate needs to raise or lower the red light.
const verdictFileName = "e2e_verdict.json"

// WriteVerdict emits e2e_verdict.json into the directory of out. The verdict
// is the authoritative tighten-only evaluation result: a failed verdict pins
// every violation with the fixed Failure{Key,Rule,Localizer,Actual,Expected}
// form, and its passed field drives the gate's exit code without parsing.
func WriteVerdict(out string, v Verdict) error {
	dir := filepath.Dir(out)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return err
	}
	encoded, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	encoded = append(encoded, '\n')
	return os.WriteFile(filepath.Join(dir, verdictFileName), encoded, filePerm)
}

// encodeDrilldown renders every record in the runner's deterministic order.
func encodeDrilldown(recs []metrics.DrilldownRecord) ([]byte, error) {
	var buf bytes.Buffer
	for i := range recs {
		data, err := json.Marshal(toDrilldownLine(recs[i]))
		if err != nil {
			return nil, err
		}
		_, _ = buf.Write(data)
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes(), nil
}
