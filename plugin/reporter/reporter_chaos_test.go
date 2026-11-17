/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package reporter

import (
	"bytes"
	"errors"
	"strconv"
	"sync"
	"testing"

	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/host"
)

// TestChaosUnknownFormat proves an unreachable format fails closed locally.
func TestChaosUnknownFormat(t *testing.T) {
	t.Parallel()
	bogus := newArtifactReporter(reporterManifest("format:sarif"), Format("bogus"))
	out, err := bogus.Build(reporterFindings(), reporterReport())
	if !errors.Is(err, ErrUnknownFormat) {
		t.Fatalf("unknown format error = %v, want ErrUnknownFormat", err)
	}
	if out != nil {
		t.Fatalf("unknown format produced %d bytes", len(out))
	}
}

// TestChaosNilFindings proves every format handles an empty finding batch.
func TestChaosNilFindings(t *testing.T) {
	t.Parallel()
	var rep mutation.Report
	for _, format := range AllFormats() {
		plugin, err := Factory(reporterManifest(string(format)))
		if err != nil {
			t.Fatalf("Factory(%s): %v", format, err)
		}
		builder, ok := plugin.(host.ReporterPlugin)
		if !ok {
			t.Fatalf("Factory did not return a ReporterPlugin: %T", plugin)
		}
		first, err := builder.Build(nil, rep)
		if err != nil {
			t.Fatalf("Build(%s, nil): %v", format, err)
		}
		second, err := builder.Build(nil, rep)
		if err != nil {
			t.Fatalf("second Build(%s, nil): %v", format, err)
		}
		if len(first) == 0 || !bytes.Equal(first, second) {
			t.Fatalf("format %s produced an empty or unstable nil-input artifact", format)
		}
	}
}

// TestChaosOversizedPayload proves a large finding batch never panics and
// stays byte-stable across two runs of every format.
func TestChaosOversizedPayload(t *testing.T) {
	t.Parallel()
	const total = 1200
	findings := make([]rule.Finding, 0, total)
	for i := 0; i < total; i++ {
		findings = append(findings, rule.Finding{
			RuleID: "GO-EG-001", Path: "pkg/file" + strconv.Itoa(i) + ".go",
			Line: i + 1, Column: 2, Severity: rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Message:   "large batch finding " + strconv.Itoa(i),
			Reference: `docs/reference/effective-go.md — Names`,
		})
	}
	rep := reporterReport()
	for _, format := range AllFormats() {
		plugin, err := Factory(reporterManifest(string(format)))
		if err != nil {
			t.Fatalf("Factory(%s): %v", format, err)
		}
		builder, ok := plugin.(host.ReporterPlugin)
		if !ok {
			t.Fatalf("Factory did not return a ReporterPlugin: %T", plugin)
		}
		first, err := builder.Build(findings, rep)
		if err != nil {
			t.Fatalf("Build(%s, large): %v", format, err)
		}
		second, err := builder.Build(findings, rep)
		if err != nil {
			t.Fatalf("second Build(%s, large): %v", format, err)
		}
		if !bytes.Equal(first, second) {
			t.Fatalf("format %s is unstable on a large batch", format)
		}
	}
}

// TestChaosConcurrentBuild proves concurrent read-only builds agree.
func TestChaosConcurrentBuild(t *testing.T) {
	t.Parallel()
	findings := reporterFindings()
	rep := reporterReport()
	plugin, err := Factory(reporterManifest("format:json"))
	if err != nil {
		t.Fatalf("Factory: %v", err)
	}
	builder, ok := plugin.(host.ReporterPlugin)
	if !ok {
		t.Fatalf("Factory did not return a ReporterPlugin, got %T", plugin)
	}
	want, err := builder.Build(findings, rep)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	var group sync.WaitGroup
	results := make(chan []byte, 24)
	for i := 0; i < 24; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			out, buildErr := builder.Build(findings, rep)
			if buildErr != nil {
				results <- nil
				return
			}
			results <- out
		}()
	}
	group.Wait()
	close(results)
	for out := range results {
		if !bytes.Equal(out, want) {
			t.Fatal("a concurrent build produced different bytes")
		}
	}
}
