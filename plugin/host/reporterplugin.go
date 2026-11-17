/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"abate/kernel/mutation"
	"abate/lang/rule"
)

// ReporterPlugin emits a report artifact from static findings and a mutation
// report. Host.Call dispatches EntryBuild to this interface, gated on CapObserve.
type ReporterPlugin interface {
	Plugin

	// Build renders findings and a mutation report into an opaque artifact
	// (for example a SARIF document) consumable by downstream reporting.
	Build(findings []rule.Finding, report mutation.Report) ([]byte, error)
}
