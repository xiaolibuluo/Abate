/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package redact

// Placeholder is the fixed marker substituted for every detected secret value.
// It is deliberately free of value terminators so re-running Text or Path on an
// already redacted string is idempotent.
const Placeholder = "[REDACTED]"

// Text redacts known secret shapes in s and returns the sanitized string.
//
// It is intended for free-form log messages and recovered panic values. The
// scan is single-pass and deterministic: Bearer tokens, Authorization headers,
// and sensitive key/value pairs have their values replaced by Placeholder,
// while unrecognized text is preserved byte for byte.
func Text(s string) string {
	return scrub(s, false)
}

// Path redacts known secret shapes in a URL path or query string.
//
// It behaves like Text but additionally treats '/' and '?' as value
// terminators, so a redacted route keeps its surrounding path structure (for
// example "/v1/token=[REDACTED]/detail") instead of swallowing the rest of the
// route.
func Path(p string) string {
	return scrub(p, true)
}

// Scrub scans s once from left to right, replacing the value of every known
// secret marker with Placeholder. When urlMode is true, '/' and '?' also end a
// value. The output is never rescanned, which makes redaction non-recursive.
func scrub(s string, urlMode bool) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); {
		if i > 0 && isKeyByte(s[i-1]) {
			out = append(out, s[i])
			i++
			continue
		}
		if width, ok := bearerWidth(s, i); ok {
			out = append(out, s[i:i+width]...)
			out = append(out, Placeholder...)
			i += width
			i += skipValue(s[i:], urlMode)
			continue
		}
		if width, ok := authorizationWidth(s, i); ok {
			out = append(out, s[i:i+width]...)
			out = append(out, Placeholder...)
			i += width
			i += skipHeaderValue(s[i:])
			continue
		}
		if eq, ok := sensitiveKeyEnd(s, i); ok {
			out = append(out, s[i:eq+1]...)
			out = append(out, Placeholder...)
			i = eq + 1
			i += skipValue(s[i:], urlMode)
			continue
		}
		out = append(out, s[i])
		i++
	}
	return string(out)
}
