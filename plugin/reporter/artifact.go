/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package reporter

import (
	"fmt"

	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/host"
	"abate/plugin/report"
)

// Format enumerates the five deterministic artifact encodings.
type Format string

// The five manifest-resolvable output formats.
const (
	// FormatHTML selects the self-contained HTML renderer.
	FormatHTML Format = "html"
	// FormatJUnit selects the JUnit XML renderer.
	FormatJUnit Format = "junit"
	// FormatOTel selects the air-gapped OTLP semantic payload renderer.
	FormatOTel Format = "otel"
	// FormatSARIF selects the SARIF 2.1.0 renderer.
	FormatSARIF Format = "sarif"
	// FormatJSON selects the compact JSON summary renderer.
	FormatJSON Format = "json"
)

// AllFormats returns the five formats in a frozen, deterministic order.
func AllFormats() []Format {
	return []Format{FormatHTML, FormatJUnit, FormatOTel, FormatSARIF, FormatJSON}
}

// artifactReporter is the single ReporterPlugin implementation in this package.
//
// 参考 sample.reporterPlugin 的单格式包装，改进点=按 manifest 解析出的格式分派
// 到五个纯函数构建器，使一个类型覆盖全部报告格式且双跑字节完全一致。
// It holds no mutable state and performs no I/O.
type artifactReporter struct {
	manifest host.Manifest
	format   Format
}

// Compile-time proof that artifactReporter satisfies the host contract.
var _ host.ReporterPlugin = (*artifactReporter)(nil)

// newArtifactReporter binds one verified manifest and its resolved format.
func newArtifactReporter(m host.Manifest, f Format) *artifactReporter {
	return &artifactReporter{manifest: m, format: f}
}

// Manifest returns the signed metadata of the plugin.
func (a *artifactReporter) Manifest() host.Manifest {
	return a.manifest
}

// Type returns the plugin kind.
func (*artifactReporter) Type() host.PluginType {
	return host.PluginTypeReporter
}

// Close releases no resources for this in-memory reporter.
func (*artifactReporter) Close() error {
	return nil
}

// Build renders findings and a mutation report into the format resolved from
// the manifest. Each branch delegates to exactly one pure report.Build*
// function, so the artifact bytes equal a direct call with the same inputs.
func (a *artifactReporter) Build(findings []rule.Finding, rep mutation.Report) ([]byte, error) {
	switch a.format {
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
		return nil, fmt.Errorf("%w: %q", ErrUnknownFormat, a.format)
	}
}
