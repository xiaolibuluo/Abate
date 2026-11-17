/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"bytes"
	"fmt"
	"html/template"

	"abate/kernel/mutation"
	"abate/lang/rule"
)

// htmlDocumentTemplate is the fixed self-contained page skeleton. All assets are
// compile-time constants and all user values pass through the template escaper.
const htmlDocumentTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Title}}</title>
<style>
` + htmlInlineCSS + `
</style>
</head>
<body>
<main>
<h1>{{.Title}}</h1>
<p class="digest">findings digest {{.FindingsDigest}}</p>
<section>
<h2>Violations ({{.ViolationTotal}})</h2>
<table>
<thead><tr><th>#</th><th>Rule</th><th>Path</th><th>Line</th><th>Column</th><th>Severity</th><th>Category</th><th>Message</th><th>Reference</th>{{if .ShowSource}}<th>Evidence</th>{{end}}</tr></thead>
<tbody>
{{range .Violations}}<tr data-severity="{{.Severity}}">
<td>{{.Index}}</td>
<td>{{.RuleID}}</td>
<td>{{.Path}}</td>
<td>{{.Line}}</td>
<td>{{.Column}}</td>
<td>{{.Severity}}</td>
<td>{{.Category}}</td>
<td>{{.Message}}</td>
<td>{{.Reference}}</td>
{{if $.ShowSource}}<td><code>{{.Evidence}}</code></td>{{end}}</tr>
{{end}}</tbody>
</table>
</section>
<section>
<h2>Mutation matrix ({{.MutantTotal}})</h2>
<table>
<thead><tr><th>#</th><th>Node</th><th>Original</th><th>Mutated</th><th>Priority</th><th>Selector</th><th>Verdict</th><th>Killing test</th><th>Fingerprint</th></tr></thead>
<tbody>
{{range .Mutants}}<tr data-verdict="{{.Verdict}}"><td>{{.Index}}</td><td>{{.Node}}</td><td>{{.Original}}</td><td>{{.Mutated}}</td><td>{{.Priority}}</td><td>{{.Selector}}</td><td>{{.Verdict}}</td><td>{{.KillingTest}}</td><td><code>{{.Fingerprint}}</code></td></tr>
{{end}}</tbody>
</table>
</section>
<section>
<h2>Timeline (virtual time)</h2>
<table>
<thead><tr><th>Virtual time</th><th>Phase</th><th>Value</th></tr></thead>
<tbody>
{{range .Timeline}}<tr data-phase="{{.Phase}}"><td>{{.VirtualTime}}</td><td>{{.Phase}}</td><td><code>{{.Value}}</code></td></tr>
{{end}}</tbody>
</table>
</section>
<section>
<h2>Score gauge</h2>
<div class="gauge">
<div>Seed<b>{{.Gauge.Seed}}</b></div>
<div>Total<b>{{.Gauge.Total}}</b></div>
<div>Killed<b>{{.Gauge.Killed}}</b></div>
<div>Survived<b>{{.Gauge.Survived}}</b></div>
<div>Equivalent<b>{{.Gauge.Equivalent}}</b></div>
<div>Timeout<b>{{.Gauge.Timeout}}</b></div>
<div>Kill rate per mille<b>{{.Gauge.RatePermille}}</b></div>
<div>Killing tests<b>{{.Gauge.KillingTests}}</b></div>
</div>
</section>
{{if .Truncated}}<p class="digest">rows truncated by MaxRows</p>{{end}}
<footer>report fingerprint {{.Footer}} / findings digest {{.FindingsDigest}}</footer>
</main>
<script>
` + htmlInlineScript + `
</script>
</body>
</html>
`

// htmlPage is parsed once at package initialization. A parsed template is safe
// for concurrent execution as long as it is never modified afterwards.
var htmlPage = template.Must(template.New("abate-report").Parse(htmlDocumentTemplate))

// BuildHTML renders findings and a mutation report into one self-contained HTML
// document. Output is a pure function of the inputs: findings are canonically
// sorted, every dynamic value is escaped by html/template and the only footer
// markers are the report fingerprint and the findings digest.
//
// 参考 PITest 的单页 HTML 报告，改进点=资源全内联、页脚无时间戳、零外部请求。
// The renderer never reads the clock, the environment or the network.
func BuildHTML(findings []rule.Finding, rep mutation.Report, opts HTMLOptions) ([]byte, error) {
	model := buildHTMLModel(findings, rep, opts)
	var buf bytes.Buffer
	if err := htmlPage.Execute(&buf, model); err != nil {
		return nil, fmt.Errorf("report: execute html template: %w", err)
	}
	return buf.Bytes(), nil
}
