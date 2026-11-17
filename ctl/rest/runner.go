/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import "abate/lang/check"

// TaskRunner executes a check request to completion. The control plane depends
// only on this port; the production kernel scheduler and the in-memory check
// runner both satisfy it, so the real experiment engine can be plugged in
// later without touching the routing or storage layers.
type TaskRunner interface {
	Run(req multiCheckRequest) (checkJobResult, error)
}

// checkRunner is the default in-memory runner backed by the multi-language
// checker. It performs no network I/O and returns deterministic results.
type checkRunner struct{}

// Run delegates to the language-neutral checker.
func (checkRunner) Run(req multiCheckRequest) (checkJobResult, error) {
	return runMultiCheck(req)
}

// toSummaryView projects a check summary into its JSON-facing view, copying the
// severity map deterministically (the source map is never leaked by reference).
func toSummaryView(summary check.Summary) summaryView {
	bySeverity := make(map[string]int, len(summary.BySeverity))
	for k, v := range summary.BySeverity {
		bySeverity[k] = v
	}
	return summaryView{
		Files:      summary.Files,
		Rules:      summary.Rules,
		Total:      summary.Total,
		BySeverity: bySeverity,
	}
}
