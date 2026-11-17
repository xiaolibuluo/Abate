/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"abate/ces/validate"
	"abate/kernel/verify"
)

// statusValidated marks a dry-run report that stopped after validation.
const statusValidated = "VALIDATED"

// report is the deterministic run report; field order is fixed by the struct.
type report struct {
	Schema       string          `json:"schema"`
	Command      string          `json:"command"`
	APIVersion   string          `json:"apiVersion"`
	Name         string          `json:"name"`
	Namespace    string          `json:"namespace"`
	Seed         uint64          `json:"seed"`
	DryRun       bool            `json:"dryRun"`
	Strict       bool            `json:"strict"`
	Status       string          `json:"status"`
	ExitCode     int             `json:"exitCode"`
	Events       uint64          `json:"events"`
	Observations int             `json:"observations"`
	End          int64           `json:"end"`
	RunDigest    string          `json:"runDigest"`
	StreamDigest string          `json:"streamDigest"`
	Violations   []violationView `json:"violations"`
	Verdicts     []verdictView   `json:"verdicts"`
}

// violationView is one ordered validation finding in a report.
type violationView struct {
	RuleID    string `json:"ruleId"`
	FieldPath string `json:"fieldPath"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
}

// verdictView is one ordered kernel verifier decision in a report.
type verdictView struct {
	ProbeID  string   `json:"probeId"`
	Outcome  string   `json:"outcome"`
	Reason   string   `json:"reason"`
	FirstSeq uint64   `json:"firstSeq"`
	Timeline []uint64 `json:"timeline,omitempty"`
}

// buildRunReport projects a run outcome onto the wire report layout.
func buildRunReport(o runOutcome) report {
	rep := report{
		Schema:     reportSchema,
		Command:    "run",
		APIVersion: o.exp.APIVersion,
		Name:       o.exp.Metadata.Name,
		Namespace:  o.exp.Metadata.Namespace,
		Seed:       o.exp.Spec.Seed,
		DryRun:     o.opts.dryRun,
		Strict:     o.opts.strict,
		Status:     statusValidated,
		ExitCode:   o.exitCode,
		Violations: violationViews(o.violations),
		Verdicts:   verdictViews(nil),
	}
	if o.execution != nil {
		rep.Status = o.execution.result.Status
		rep.Events = o.execution.result.Events
		rep.Observations = o.execution.observations
		rep.End = o.execution.result.End.UnixNano()
		rep.RunDigest = hex.EncodeToString(o.execution.result.Digest[:])
		rep.StreamDigest = hex.EncodeToString(o.execution.streamDigest[:])
		rep.Verdicts = verdictViews(o.execution.verdicts)
	}
	return rep
}

// violationViews preserves the validator's deterministic ordering.
func violationViews(vs []validate.Violation) []violationView {
	out := make([]violationView, 0, len(vs))
	for _, v := range vs {
		out = append(out, violationView{
			RuleID: v.RuleID, FieldPath: v.FieldPath,
			Message: v.Message, Severity: string(v.Severity),
		})
	}
	return out
}

// verdictViews preserves the kernel verifier's deterministic ordering.
func verdictViews(vs []verify.Verdict) []verdictView {
	out := make([]verdictView, 0, len(vs))
	for _, v := range vs {
		out = append(out, verdictView{
			ProbeID: v.ProbeID, Outcome: string(v.Outcome),
			Reason: v.Reason, FirstSeq: v.FirstSeq,
			Timeline: append([]uint64(nil), v.Timeline...),
		})
	}
	return out
}

// renderReport serializes a report in the requested deterministic format.
func renderReport(rep report, format string) []byte {
	if format == formatJSON {
		b, err := json.MarshalIndent(rep, "", "  ")
		if err == nil {
			return append(b, '\n')
		}
	}
	return renderRunConsole(rep)
}

// renderRunConsole renders a report as fixed-order text lines.
func renderRunConsole(rep report) []byte {
	var b strings.Builder
	appendLine(&b, "abate %s report\n", rep.Command)
	appendLine(&b, "schema: %s\n", rep.Schema)
	appendLine(&b, "apiVersion: %s\n", rep.APIVersion)
	appendLine(&b, "name: %s\n", rep.Name)
	appendLine(&b, "namespace: %s\n", rep.Namespace)
	appendLine(&b, "seed: %d\n", rep.Seed)
	appendLine(&b, "dryRun: %t\n", rep.DryRun)
	appendLine(&b, "strict: %t\n", rep.Strict)
	appendLine(&b, "status: %s\n", rep.Status)
	appendLine(&b, "exitCode: %d\n", rep.ExitCode)
	appendLine(&b, "events: %d\n", rep.Events)
	appendLine(&b, "observations: %d\n", rep.Observations)
	appendLine(&b, "end: %d\n", rep.End)
	appendLine(&b, "runDigest: %s\n", rep.RunDigest)
	appendLine(&b, "streamDigest: %s\n", rep.StreamDigest)
	appendLine(&b, "violations: %d\n", len(rep.Violations))
	for _, v := range rep.Violations {
		appendLine(&b, "  %s | %s | %s | %s\n",
			v.RuleID, v.FieldPath, v.Severity, v.Message)
	}
	appendLine(&b, "verdicts: %d\n", len(rep.Verdicts))
	for _, v := range rep.Verdicts {
		appendLine(&b, "  %s | %s | %s\n", v.ProbeID, v.Outcome, v.Reason)
	}
	return []byte(b.String())
}

// appendLine writes one deterministic console line.
func appendLine(b *strings.Builder, format string, args ...any) {
	_, _ = fmt.Fprintf(b, format, args...)
}

// writeOutput writes the report to an optional file before stdout, so a failed
// file write returns an error without emitting a report whose exitCode disagrees
// with the process exit code.
func writeOutput(data []byte, out string, stdout, stderr io.Writer) error {
	if out != "" {
		if err := os.WriteFile(out, data, 0o600); err != nil { //nolint:gosec // the output path is supplied explicitly by the caller.
			_, _ = fmt.Fprintf(stderr, "abate: cannot write report %q: %v\n", out, err)
			return err
		}
	}
	if _, err := stdout.Write(data); err != nil {
		return err
	}
	return nil
}
