/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package reporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/report"
)

// reporterManifest builds a reporter manifest carrying one format token.
func reporterManifest(buildRef string) host.Manifest {
	return host.Manifest{
		ID:       host.PluginID("abate/rpt-m5"),
		Type:     host.PluginTypeReporter,
		BuildRef: buildRef,
	}
}

// reporterFindings returns a fixed, deliberately unordered finding batch.
func reporterFindings() []rule.Finding {
	return []rule.Finding{
		{
			RuleID: "GO-EG-002", Path: "pkg/b.go", Line: 9, Column: 4,
			Severity: rule.SeverityError, Category: rule.CategoryNaming,
			Message: "avoid ALL_CAPS identifiers", Reference: "Effective Go, Names",
		},
		{
			RuleID: "GO-EG-001", Path: "pkg/a.go", Line: 2, Column: 1,
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Message: `mixed <tag> & "quote"`, Reference: "Effective Go, Names",
		},
		{
			RuleID: "GO-VET-003", Path: "pkg/c.go", Line: 17, Column: 6,
			Severity: rule.SeverityInfo, Category: rule.CategoryPerformance,
			Message: "unnamed numeric literal", Reference: "Go Code Review Comments",
		},
	}
}

// reporterReport returns a deterministic mutation report.
func reporterReport() mutation.Report {
	return mutation.NewReport(7, [32]byte{}, [32]byte{}, nil, nil, nil)
}

// pureBuild calls the matching pure report.Build* function directly.
func pureBuild(format Format, findings []rule.Finding, rep mutation.Report) ([]byte, error) {
	switch format {
	case FormatHTML:
		return report.BuildHTML(findings, rep, report.DefaultHTMLOptions())
	case FormatJUnit:
		return report.BuildJUnit(findings, rep)
	case FormatOTel:
		return report.BuildOTel(findings, rep)
	case FormatSARIF:
		return report.BuildSARIF(findings)
	case FormatJSON:
		return report.BuildJSON(findings, rep)
	default:
		return nil, ErrUnknownFormat
	}
}

// TestFiveFormatsMatchPureFunctions proves dispatch and byte equality.
func TestFiveFormatsMatchPureFunctions(t *testing.T) {
	t.Parallel()
	findings := reporterFindings()
	rep := reporterReport()
	for _, format := range AllFormats() {
		t.Run(string(format), func(t *testing.T) {
			plugin, err := Factory(reporterManifest("format:" + string(format)))
			if err != nil {
				t.Fatalf("Factory: %v", err)
			}
			builder, ok := plugin.(host.ReporterPlugin)
			if !ok {
				t.Fatal("Factory did not return a ReporterPlugin")
			}
			want, err := pureBuild(format, findings, rep)
			if err != nil {
				t.Fatalf("pure builder: %v", err)
			}
			got, err := builder.Build(findings, rep)
			if err != nil {
				t.Fatalf("Build: %v", err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("plugin output differs from the pure %s builder", format)
			}
			again, err := builder.Build(findings, rep)
			if err != nil {
				t.Fatalf("second Build: %v", err)
			}
			if !bytes.Equal(got, again) {
				t.Fatalf("plugin %s output is not byte-stable", format)
			}
			if len(got) == 0 {
				t.Fatalf("plugin %s produced an empty artifact", format)
			}
		})
	}
}

// TestHostCallEntryBuildMatchesPureFunctions proves the real host dispatch
// path produces the very same bytes as a direct pure builder call.
func TestHostCallEntryBuildMatchesPureFunctions(t *testing.T) {
	t.Parallel()
	findings := reporterFindings()
	rep := reporterReport()
	args, err := host.EncodeBuildArgs(host.BuildArgs{Findings: findings, Report: rep})
	if err != nil {
		t.Fatalf("EncodeBuildArgs: %v", err)
	}
	for _, format := range AllFormats() {
		plugin, err := Factory(reporterManifest(string(format)))
		if err != nil {
			t.Fatalf("Factory(%s): %v", format, err)
		}
		var caps host.Capabilities
		caps.Grant(capability.CapObserve)
		inst := &host.Instance{
			Plugin:  plugin,
			Caps:    caps,
			Entries: []string{host.EntryBuild},
		}
		out, err := (&host.Host{}).Call(inst, host.EntryBuild, args, nil)
		if err != nil {
			t.Fatalf("Host.Call(%s): %v", format, err)
		}
		want, err := pureBuild(format, findings, rep)
		if err != nil {
			t.Fatalf("pure builder(%s): %v", format, err)
		}
		if !bytes.Equal(out, want) {
			t.Fatalf("Host.Call(%s) is not byte-identical to the pure builder", format)
		}
	}
}

// TestFactoryFailClosed proves wrong types and unknown formats are rejected.
func TestFactoryFailClosed(t *testing.T) {
	t.Parallel()
	badType := reporterManifest("format:json")
	badType.Type = host.PluginTypeAdapter
	if _, err := Factory(badType); !errors.Is(err, ErrWrongPluginType) {
		t.Fatalf("wrong type error = %v, want ErrWrongPluginType", err)
	}
	for _, ref := range []string{"", "   ", "xml", "format:pdf", "format="} {
		if _, err := Factory(reporterManifest(ref)); !errors.Is(err, ErrUnknownFormat) {
			t.Fatalf("Factory(%q) error = %v, want ErrUnknownFormat", ref, err)
		}
	}
	for _, ref := range []string{"JSON", "format=HTML", " format:otel ", "sarif"} {
		if _, err := Factory(reporterManifest(ref)); err != nil {
			t.Fatalf("Factory(%q) rejected a valid format: %v", ref, err)
		}
	}
}

// TestReporterMetadata proves the plugin satisfies the base host contract.
func TestReporterMetadata(t *testing.T) {
	t.Parallel()
	m := reporterManifest("format:junit")
	plugin, err := Factory(m)
	if err != nil {
		t.Fatalf("Factory: %v", err)
	}
	if plugin.Manifest().ID != m.ID || plugin.Type() != host.PluginTypeReporter {
		t.Fatalf("metadata mismatch: %+v", plugin.Manifest())
	}
	if err := plugin.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if len(AllFormats()) != 5 {
		t.Fatalf("format count = %d, want 5", len(AllFormats()))
	}
}

// TestArtifactsAreWellFormed proves every dispatch branch emits parseable data.
func TestArtifactsAreWellFormed(t *testing.T) {
	t.Parallel()
	findings := reporterFindings()
	rep := reporterReport()
	for _, format := range []Format{FormatJSON, FormatOTel, FormatSARIF} {
		out, err := pureBuild(format, findings, rep)
		if err != nil {
			t.Fatalf("pure builder %s: %v", format, err)
		}
		var doc any
		if err := json.Unmarshal(out, &doc); err != nil {
			t.Fatalf("format %s is not valid JSON: %v", format, err)
		}
	}
}
