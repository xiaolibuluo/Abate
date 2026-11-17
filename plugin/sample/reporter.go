/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sample

import (
	"fmt"

	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/host"
	"abate/plugin/report"
)

// ReporterPlugin renders findings and a mutation report into portable artifacts.
// It delegates to the report package so the emitted SARIF is the same OASIS-valid
// document the acceptance tests validate.
type ReporterPlugin struct {
	manifest host.Manifest
}

// NewReporterPlugin constructs a sample reporter plugin from its manifest.
func NewReporterPlugin(m host.Manifest) *ReporterPlugin {
	return &ReporterPlugin{manifest: m}
}

// Manifest returns the signed metadata of the plugin.
func (p *ReporterPlugin) Manifest() host.Manifest { return p.manifest }

// Type returns the plugin kind.
func (*ReporterPlugin) Type() host.PluginType { return host.PluginTypeReporter }

// Close releases no resources for the in-memory sample reporter.
func (*ReporterPlugin) Close() error { return nil }

// Build renders the findings and mutation report into a SARIF 2.1.0 artifact that
// is validated against the OASIS schema by the report package.
func (*ReporterPlugin) Build(findings []rule.Finding, _ mutation.Report) ([]byte, error) {
	return report.BuildSARIF(findings)
}

// ReporterFactory instantiates the sample reporter plugin from a verified manifest.
// It fails closed if the manifest is not a reporter.
func ReporterFactory(m host.Manifest) (host.Plugin, error) {
	if m.Type != host.PluginTypeReporter {
		return nil, fmt.Errorf("sample: manifest %q is not a reporter (type %q)", m.ID, m.Type)
	}
	return NewReporterPlugin(m), nil
}
