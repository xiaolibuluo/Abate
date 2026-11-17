/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package reporter

import (
	"bytes"
	"testing"

	"abate/lang/rule"
	"abate/plugin/host"
)

// TestMutationFormatChangesBytes proves that mutating the manifest format
// changes the emitted artifact for one fixed input pair.
func TestMutationFormatChangesBytes(t *testing.T) {
	t.Parallel()
	findings := reporterFindings()
	rep := reporterReport()
	seen := make(map[string][]byte, len(AllFormats()))
	for _, format := range AllFormats() {
		plugin, err := Factory(reporterManifest(string(format)))
		if err != nil {
			t.Fatalf("Factory(%s): %v", format, err)
		}
		out, err := plugin.(host.ReporterPlugin).Build(findings, rep)
		if err != nil {
			t.Fatalf("Build(%s): %v", format, err)
		}
		for other, prior := range seen {
			if bytes.Equal(out, prior) {
				t.Fatalf("formats %s and %s produced identical artifacts", format, other)
			}
		}
		seen[string(format)] = out
	}
}

// TestMutationBuildRefChangesFactory proves mutating the format carrier
// changes the fail-closed decision.
func TestMutationBuildRefChangesFactory(t *testing.T) {
	t.Parallel()
	base := reporterManifest("format:sarif")
	good, err := Factory(base)
	if err != nil {
		t.Fatalf("baseline Factory: %v", err)
	}
	mutated := base
	mutated.BuildRef = "format:unsupported"
	if _, err := Factory(mutated); err == nil {
		t.Fatal("mutating BuildRef to an unknown format was accepted")
	}
	mutated.BuildRef = "format:otel"
	other, err := Factory(mutated)
	if err != nil {
		t.Fatalf("mutated Factory: %v", err)
	}
	findings := reporterFindings()
	rep := reporterReport()
	sarif, err := good.(host.ReporterPlugin).Build(findings, rep)
	if err != nil {
		t.Fatalf("baseline Build: %v", err)
	}
	otel, err := other.(host.ReporterPlugin).Build(findings, rep)
	if err != nil {
		t.Fatalf("mutated Build: %v", err)
	}
	if bytes.Equal(sarif, otel) {
		t.Fatal("mutating the format carrier did not change the artifact")
	}
}

// TestMutationFindingsChangeBytes proves a single finding mutation is visible
// in every format, so the dispatch is not a constant-output stub.
func TestMutationFindingsChangeBytes(t *testing.T) {
	t.Parallel()
	findings := reporterFindings()
	rep := reporterReport()
	mutated := make([]rule.Finding, len(findings))
	copy(mutated, findings)
	mutated[0] = rule.Finding{
		RuleID: "GO-EG-999", Path: "pkg/mutated.go", Line: 99, Column: 8,
		Severity: rule.SeverityError, Category: rule.CategoryBoundary,
		Message: "mutated finding", Reference: "Mutated Reference",
	}
	for _, format := range AllFormats() {
		plugin, err := Factory(reporterManifest(string(format)))
		if err != nil {
			t.Fatalf("Factory(%s): %v", format, err)
		}
		builder, ok := plugin.(host.ReporterPlugin)
		if !ok {
			t.Fatalf("Factory(%s) did not return a ReporterPlugin: %T", format, plugin)
		}
		before, err := builder.Build(findings, rep)
		if err != nil {
			t.Fatalf("Build(%s): %v", format, err)
		}
		after, err := builder.Build(mutated, rep)
		if err != nil {
			t.Fatalf("mutated Build(%s): %v", format, err)
		}
		if bytes.Equal(before, after) {
			t.Fatalf("format %s ignored a mutated finding", format)
		}
	}
}
