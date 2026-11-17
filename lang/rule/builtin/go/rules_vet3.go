/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goVET3Rules declares the vet structtag, unmarshal, httpresponse, timeformat
// and slog semantic subset. Struct tags and layout strings are inspected by
// attribute constraints rather than by re-parsing Go syntax.
func goVET3Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-VET-023", Reference: "cmd/vet structtag: verify json tag spelling",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `(Attribute~json)`,
			Message: "verify the json struct tag key and options",
		},
		{
			ID: "GO-VET-024", Reference: "cmd/vet unmarshal: pointer target",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Unmarshal) $..)`,
			Message: "json.Unmarshal requires a non-nil pointer target",
		},
		{
			ID: "GO-VET-025", Reference: "cmd/vet unmarshal: check the marshal error",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Marshal) $..)`,
			Message: "check the error returned by json.Marshal",
		},
		{
			ID: "GO-VET-026", Reference: "cmd/vet httpresponse: close the response body",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Get) $.. StrLit)`,
			Message: "close the body of an http.Get response",
		},
		{
			ID: "GO-VET-027", Reference: "cmd/vet timeformat: Format reference layout",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Format) $.. StrLit~2006)`,
			Message: "verify the Go reference layout used by time.Format",
		},
		{
			ID: "GO-VET-028", Reference: "cmd/vet timeformat: Parse reference layout",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Parse) $.. StrLit~2006)`,
			Message: "verify the Go reference layout used by time.Parse",
		},
		{
			ID: "GO-VET-029", Reference: "cmd/vet slog: With needs key-value pairs",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `(Call (Selector+With) $..)`,
			Message: "slog With arguments must alternate string keys and values",
		},
		{
			ID: "GO-VET-030", Reference: "cmd/vet slog: no printf directives in Log",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Log) $.. StrLit~%)`,
			Message: "structured logging messages must not contain printf directives",
		},
		{
			ID: "GO-VET-031", Reference: "cmd/vet slog: LogAttrs arguments",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `(Call (Selector+LogAttrs) $..)`,
			Message: "verify that LogAttrs receives well-formed slog.Attr values",
		},
		{
			ID: "GO-VET-032", Reference: "cmd/vet slog: no printf directives in Info",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Info) $.. StrLit~%)`,
			Message: "structured logging messages must not contain printf directives",
		},
		{
			ID: "GO-VET-033", Reference: "cmd/vet slog: no printf directives in Warn",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Warn) $.. StrLit~%)`,
			Message: "structured logging messages must not contain printf directives",
		},
	}
}
