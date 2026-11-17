/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

// HTMLOptions controls the self-contained HTML renderer. It deliberately has no
// timestamp or network fields: the page never records wall-clock time and never
// loads a remote resource.
//
// 参考 PITest 的 HTML 报告配置项，改进点=去掉 CI 时间戳开关，仅保留纯展示选项。
// Every option has a deterministic default that never depends on wall-clock time.
type HTMLOptions struct {
	// EmbedSource includes the canonical finding excerpt in each violation row.
	EmbedSource bool
	// Title overrides the document title.
	Title string
	// MaxRows caps the rendered rows in each table; zero or negative means all.
	MaxRows int
}

// defaultHTMLTitle is the title used whenever the caller leaves it empty.
const defaultHTMLTitle = "Abate Report"

// DefaultHTMLOptions returns the deterministic default rendering options.
func DefaultHTMLOptions() HTMLOptions {
	return HTMLOptions{
		EmbedSource: false,
		Title:       defaultHTMLTitle,
		MaxRows:     0,
	}
}

// normalizeHTMLOptions fills missing values and clamps the row cap so a zero
// options value produces exactly the same bytes as the defaults.
func normalizeHTMLOptions(opts HTMLOptions) HTMLOptions {
	norm := DefaultHTMLOptions()
	if opts.Title != "" {
		norm.Title = opts.Title
	}
	norm.EmbedSource = opts.EmbedSource
	if opts.MaxRows > 0 {
		norm.MaxRows = opts.MaxRows
	}
	return norm
}
